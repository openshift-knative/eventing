/*
Copyright 2020 The Knative Authors

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package shell

import (
	"bytes"
	"fmt"
	"os"
	"os/exec"
	"strings"
	"time"
)

const (
	defaultLabelOut = "[OUT]"
	defaultLabelErr = "[ERR]"
	executeMode     = 0o700
)

// NewExecutor creates a new executor.
func NewExecutor(t TestingT, loc ProjectLocation, opts ...Option) Executor {
	config := &ExecutorConfig{
		ProjectLocation: loc,
		Streams:         testingTStreams(t),
	}
	for _, opt := range opts {
		opt(config)
	}
	configureDefaultValues(config)
	return &streamingExecutor{
		ExecutorConfig: *config,
	}
}

// testingTStreams returns Streams which writes to test log.
func testingTStreams(t TestingT) Streams {
	tWriter := testingWriter{t: t}
	return Streams{
		Out: tWriter,
		Err: tWriter,
	}
}

// RunScript executes a shell script with args.
func (s *streamingExecutor) RunScript(script Script, args ...string) error {
	cnt := script.scriptContent(s.ProjectLocation, args)
	return withTempScript(cnt, func(bin string) error {
		return stream(bin, s.ExecutorConfig, script.Label)
	})
}

// RunFunction executes a shell function with args.
func (s *streamingExecutor) RunFunction(fn Function, args ...string) error {
	cnt := fn.scriptContent(s.ProjectLocation, args)
	return withTempScript(cnt, func(bin string) error {
		return stream(bin, s.ExecutorConfig, fn.Label)
	})
}

type streamingExecutor struct {
	ExecutorConfig
}

func configureDefaultValues(config *ExecutorConfig) {
	if config.LabelOut == "" {
		config.LabelOut = defaultLabelOut
	}
	if config.LabelErr == "" {
		config.LabelErr = defaultLabelErr
	}
	if config.Environ == nil {
		config.Environ = os.Environ()
	}
	if !config.SkipDate && config.DateFormat == "" {
		config.DateFormat = time.StampMilli
	}
	if config.PrefixFunc == nil {
		config.PrefixFunc = defaultPrefixFunc
	}
}

func stream(bin string, cfg ExecutorConfig, label string) error {
	c := exec.Command(bin)
	c.Env = cfg.Environ
	c.Stdout = NewPrefixer(cfg.Out, prefixFunc(StreamTypeOut, label, cfg))
	c.Stderr = NewPrefixer(cfg.Err, prefixFunc(StreamTypeErr, label, cfg))
	return c.Run()
}

func prefixFunc(st StreamType, label string, cfg ExecutorConfig) func() string {
	return func() string {
		return cfg.PrefixFunc(st, label, cfg)
	}
}

func defaultPrefixFunc(st StreamType, label string, cfg ExecutorConfig) string {
	sep := " "
	var buf []string
	if !cfg.SkipDate {
		dt := time.Now().Format(cfg.DateFormat)
		buf = append(buf, dt)
	}
	buf = append(buf, label)
	switch st {
	case StreamTypeOut:
		buf = append(buf, cfg.LabelOut)
	case StreamTypeErr:
		buf = append(buf, cfg.LabelErr)
	}
	return strings.Join(buf, sep) + sep
}

func withTempScript(contents string, fn func(bin string) error) error {
	tmpfile, err := os.CreateTemp("", "shellout-*.sh")
	if err != nil {
		return err
	}
	_, err = tmpfile.WriteString(contents)
	if err != nil {
		return err
	}
	err = tmpfile.Chmod(executeMode)
	if err != nil {
		return err
	}
	err = tmpfile.Close()
	if err != nil {
		return err
	}
	defer func() {
		// clean up
		_ = os.Remove(tmpfile.Name())
	}()

	return fn(tmpfile.Name())
}

func (fn *Function) scriptContent(location ProjectLocation, args []string) string {
	return fmt.Sprintf(`#!/usr/bin/env bash

set -Eeuo pipefail

cd "%s"
source %s

%s %s
`, location.RootPath(), fn.ScriptPath, fn.FunctionName, quoteArgs(args))
}

func (sc *Script) scriptContent(location ProjectLocation, args []string) string {
	return fmt.Sprintf(`#!/usr/bin/env bash

set -Eeuo pipefail

cd "%s"
%s %s
`, location.RootPath(), sc.ScriptPath, quoteArgs(args))
}

func quoteArgs(args []string) string {
	quoted := make([]string, len(args))
	for i, arg := range args {
		quoted[i] = "\"" + strings.ReplaceAll(arg, "\"", "\\\"") + "\""
	}
	return strings.Join(quoted, " ")
}

func (w testingWriter) Write(p []byte) (n int, err error) {
	n = len(p)

	// Strip trailing newline because t.Log always adds one.
	p = bytes.TrimRight(p, "\n")

	for _, line := range strings.Split(string(p), "\n") {
		w.t.Logf(line)
	}

	return n, nil
}