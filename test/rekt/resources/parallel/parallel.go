/*
Copyright 2021 The Knative Authors

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

package parallel

import (
	"context"
	"embed"
	"strings"
	"time"

	"k8s.io/apimachinery/pkg/runtime/schema"
	duckv1 "knative.dev/pkg/apis/duck/v1"
	"knative.dev/reconciler-test/pkg/feature"
	"knative.dev/reconciler-test/pkg/k8s"
	"knative.dev/reconciler-test/pkg/manifest"

	"knative.dev/eventing/test/rekt/resources/addressable"
	"knative.dev/eventing/test/rekt/resources/channel_template"
)

//go:embed *.yaml
var yaml embed.FS

func GVR() schema.GroupVersionResource {
	return schema.GroupVersionResource{Group: "flows.knative.dev", Version: "v1", Resource: "parallels"}
}

func GVK() schema.GroupVersionKind {
	return schema.GroupVersionKind{Group: "flows.knative.dev", Version: "v1", Kind: "Parallel"}
}

// Install will create a Parallel resource, augmented with the config fn options.
func Install(name string, opts ...manifest.CfgFn) feature.StepFn {
	cfg := map[string]interface{}{
		"name": name,
	}
	for _, fn := range opts {
		fn(cfg)
	}
	return func(ctx context.Context, t feature.T) {
		if _, err := manifest.InstallYamlFS(ctx, yaml, cfg); err != nil {
			t.Fatal(err)
		}
	}
}

// IsReady tests to see if a Parallel becomes ready within the time given.
func IsReady(name string, timing ...time.Duration) feature.StepFn {
	return k8s.IsReady(GVR(), name, timing...)
}

// ValidateAddress validates the address retured by Address
func ValidateAddress(name string, validate addressable.ValidateAddressFn, timings ...time.Duration) feature.StepFn {
	return addressable.ValidateAddress(GVR(), name, validate, timings...)
}

// IsAddressable tests to see if a Parallel becomes addressable within the  time
// given.
func IsAddressable(name string, timing ...time.Duration) feature.StepFn {
	return k8s.IsAddressable(GVR(), name, timing...)
}

// Address returns a Parallel's address.
func Address(ctx context.Context, name string, timings ...time.Duration) (*duckv1.Addressable, error) {
	return addressable.Address(ctx, GVR(), name, timings...)
}

// AsRef returns a KRef for a Parallel without namespace.
func AsRef(name string) *duckv1.KReference {
	apiVersion, kind := GVK().ToAPIVersionAndKind()
	return &duckv1.KReference{
		Kind:       kind,
		APIVersion: apiVersion,
		Name:       name,
	}
}

// WithSubscriberAt adds the subscriber related config to a Parallel spec at branches[`index`].
func WithSubscriberAt(index int, d *duckv1.Destination) manifest.CfgFn {
	return func(cfg map[string]interface{}) {
		if _, set := cfg["branches"]; !set {
			cfg["branches"] = []map[string]interface{}{}
		}

		branches := cfg["branches"].([]map[string]interface{})
		// Grow the array.
		for cap(branches) <= index {
			branches = append(branches, map[string]interface{}{})
		}

		branch := branches[index]
		if _, set := branch["subscriber"]; !set {
			branch["subscriber"] = map[string]interface{}{}
		}
		subscriber := branch["subscriber"].(map[string]interface{})

		if d.URI != nil {
			subscriber["uri"] = d.URI.String()
		}
		if d.Ref != nil {
			if _, set := subscriber["ref"]; !set {
				subscriber["ref"] = map[string]interface{}{}
			}
			sref := subscriber["ref"].(map[string]interface{})
			sref["apiVersion"] = d.Ref.APIVersion
			sref["kind"] = d.Ref.Kind
			sref["namespace"] = d.Ref.Namespace
			sref["name"] = d.Ref.Name
		}
		if d.CACerts != nil {
			// This is a multi-line string and should be indented accordingly.
			// Replace "new line" with "new line + spaces".
			subscriber["CACerts"] = strings.ReplaceAll(*d.CACerts, "\n", "\n          ")
		}

		cfg["branches"] = branches
	}
}

// WithFilterAt adds the filter related config to a Parallel spec at branches[`index`].
func WithFilterAt(index int, d *duckv1.Destination) manifest.CfgFn {
	return func(cfg map[string]interface{}) {
		if _, set := cfg["branches"]; !set {
			cfg["branches"] = []map[string]interface{}{}
		}

		branches := cfg["branches"].([]map[string]interface{})
		// Grow the array.
		for cap(branches) <= index {
			branches = append(branches, map[string]interface{}{})
		}

		branch := branches[index]
		if _, set := branch["filter"]; !set {
			branch["filter"] = map[string]interface{}{}
		}
		filter := branch["filter"].(map[string]interface{})

		if d == nil {
			return
		}

		if d.URI != nil {
			filter["uri"] = d.URI.String()
		}
		if d.Ref != nil {
			if _, set := filter["ref"]; !set {
				filter["ref"] = map[string]interface{}{}
			}
			fref := filter["ref"].(map[string]interface{})
			fref["apiVersion"] = d.Ref.APIVersion
			fref["kind"] = d.Ref.Kind
			fref["namespace"] = d.Ref.Namespace
			fref["name"] = d.Ref.Name
		}
		if d.CACerts != nil {
			// This is a multi-line string and should be indented accordingly.
			// Replace "new line" with "new line + spaces".
			filter["CACerts"] = strings.ReplaceAll(*d.CACerts, "\n", "\n          ")
		}

		cfg["branches"] = branches
	}
}

// WithReplyAt adds the reply related config to a Parallel spec at branches[`index`].
func WithReplyAt(index int, d *duckv1.Destination) manifest.CfgFn {
	return func(cfg map[string]interface{}) {
		if _, set := cfg["branches"]; !set {
			cfg["branches"] = []map[string]interface{}{}
		}

		branches := cfg["branches"].([]map[string]interface{})
		// Grow the array.
		for cap(branches) <= index {
			branches = append(branches, map[string]interface{}{})
		}

		branch := branches[index]
		if _, set := branch["reply"]; !set {
			branch["reply"] = map[string]interface{}{}
		}
		reply := branch["reply"].(map[string]interface{})

		if d == nil {
			return
		}

		if d.URI != nil {
			reply["uri"] = d.URI.String()
		}
		if d.Ref != nil {
			if _, set := reply["ref"]; !set {
				reply["ref"] = map[string]interface{}{}
			}
			rref := reply["ref"].(map[string]interface{})
			rref["apiVersion"] = d.Ref.APIVersion
			rref["kind"] = d.Ref.Kind
			rref["namespace"] = d.Ref.Namespace
			rref["name"] = d.Ref.Name
		}
		if d.CACerts != nil {
			// This is a multi-line string and should be indented accordingly.
			// Replace "new line" with "new line + spaces".
			reply["CACerts"] = strings.ReplaceAll(*d.CACerts, "\n", "\n          ")
		}

		cfg["branches"] = branches
	}
}

// WithReply adds the top level reply config to a Parallel spec.
func WithReply(d *duckv1.Destination) manifest.CfgFn {
	return func(cfg map[string]interface{}) {
		if _, set := cfg["reply"]; !set {
			cfg["reply"] = map[string]interface{}{}
		}
		reply := cfg["reply"].(map[string]interface{})

		if d == nil {
			return
		}

		if d.URI != nil {
			reply["uri"] = d.URI.String()
		}
		if d.Ref != nil {
			if _, set := reply["ref"]; !set {
				reply["ref"] = map[string]interface{}{}
			}
			rref := reply["ref"].(map[string]interface{})
			rref["apiVersion"] = d.Ref.APIVersion
			rref["kind"] = d.Ref.Kind
			rref["namespace"] = d.Ref.Namespace
			rref["name"] = d.Ref.Name

		}
		if d.CACerts != nil {
			// This is a multi-line string and should be indented accordingly.
			// Replace "new line" with "new line + spaces".
			reply["CACerts"] = strings.ReplaceAll(*d.CACerts, "\n", "\n      ")
		}
	}
}

// WithChannelTemplate adds the top level channel references.
func WithChannelTemplate(template channel_template.ChannelTemplate) manifest.CfgFn {
	return func(cfg map[string]interface{}) {
		if _, set := cfg["channelTemplate"]; !set {
			cfg["channelTemplate"] = map[string]interface{}{}
		}
		channelTemplate := cfg["channelTemplate"].(map[string]interface{})

		channelTemplate["kind"] = template.Kind
		channelTemplate["apiVersion"] = template.APIVersion

		channelTemplate["spec"] = template.Spec
	}
}

// ValidateAddress validates the address retured by Address
func ValidateAddress(name string, validate addressable.ValidateAddressFn, timings ...time.Duration) feature.StepFn {
	return addressable.ValidateAddress(GVR(), name, validate, timings...)
}
