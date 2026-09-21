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
	cetest "github.com/cloudevents/sdk-go/v2/test"
	"knative.dev/reconciler-test/pkg/eventshub/assert"
	"knative.dev/reconciler-test/pkg/feature"
)

func assertRetryDelivery(f *feature.Feature, receiverName, eventID string) {
	f.Assert("receiver rejects the first four deliveries", assert.OnStore(receiverName).
		MatchRejectedEvent(cetest.HasId(eventID)).Exact(4))
	f.Assert("receiver accepts the fifth delivery", assert.OnStore(receiverName).
		MatchReceivedEvent(cetest.HasId(eventID)).Exact(1))
}
