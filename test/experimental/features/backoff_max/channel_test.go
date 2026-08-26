/*
Copyright 2026 The Knative Authors

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

package backoff_max

import (
	"testing"
	"time"

	cetest "github.com/cloudevents/sdk-go/v2/test"
	"github.com/stretchr/testify/require"
	"knative.dev/reconciler-test/pkg/eventshub"
	"knative.dev/reconciler-test/pkg/feature"
)

func TestChannelToSinkRunsSenderAfterReadiness(t *testing.T) {
	timings := make(map[string]feature.Timing)
	for _, step := range ChannelToSink().Steps {
		timings[step.Name] = step.T
	}

	for _, name := range []string{"channel is ready", "subscription is ready"} {
		timing, ok := timings[name]
		require.Truef(t, ok, "step %q not found", name)
		require.Equal(t, feature.Requirement, timing)
	}

	timing, ok := timings["send event"]
	require.True(t, ok, "step %q not found", "send event")
	require.Equal(t, feature.Assert, timing)
}

func TestDeliveriesFollowBackoff(t *testing.T) {
	event := cetest.FullEvent()
	expected := []time.Duration{time.Second, 2 * time.Second, 2 * time.Second, 2 * time.Second}
	tests := []struct {
		name      string
		intervals []time.Duration
		wantErr   bool
	}{
		{
			name:      "capped exponential backoff",
			intervals: []time.Duration{time.Second, 2 * time.Second, 2 * time.Second, 2 * time.Second},
		},
		{
			name:      "uncapped exponential backoff",
			intervals: []time.Duration{time.Second, 2 * time.Second, 4 * time.Second, 8 * time.Second},
			wantErr:   true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			matcher := deliveriesFollowBackoff(event.ID(), expected)
			receivedAt := time.Unix(0, 0)
			var gotErr error
			for i := 0; i <= len(tt.intervals); i++ {
				if i > 0 {
					receivedAt = receivedAt.Add(tt.intervals[i-1])
				}

				kind := eventshub.EventRejected
				sequence := uint64(i + 1)
				if i == len(tt.intervals) {
					kind = eventshub.EventReceived
					sequence = 1
				}
				gotErr = matcher(eventshub.EventInfo{
					Event:    &event,
					Kind:     kind,
					Time:     receivedAt,
					Sequence: sequence,
				})
				if i < len(tt.intervals) {
					require.NoError(t, gotErr)
				}
			}

			if tt.wantErr {
				require.Error(t, gotErr)
			} else {
				require.NoError(t, gotErr)
			}
		})
	}
}
