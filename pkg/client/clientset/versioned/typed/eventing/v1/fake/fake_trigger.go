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

// Code generated by client-gen. DO NOT EDIT.

package fake

import (
	gentype "k8s.io/client-go/gentype"
	v1 "knative.dev/eventing/pkg/apis/eventing/v1"
	eventingv1 "knative.dev/eventing/pkg/client/clientset/versioned/typed/eventing/v1"
)

// fakeTriggers implements TriggerInterface
type fakeTriggers struct {
	*gentype.FakeClientWithList[*v1.Trigger, *v1.TriggerList]
	Fake *FakeEventingV1
}

func newFakeTriggers(fake *FakeEventingV1, namespace string) eventingv1.TriggerInterface {
	return &fakeTriggers{
		gentype.NewFakeClientWithList[*v1.Trigger, *v1.TriggerList](
			fake.Fake,
			namespace,
			v1.SchemeGroupVersion.WithResource("triggers"),
			v1.SchemeGroupVersion.WithKind("Trigger"),
			func() *v1.Trigger { return &v1.Trigger{} },
			func() *v1.TriggerList { return &v1.TriggerList{} },
			func(dst, src *v1.TriggerList) { dst.ListMeta = src.ListMeta },
			func(list *v1.TriggerList) []*v1.Trigger { return gentype.ToPointerSlice(list.Items) },
			func(list *v1.TriggerList, items []*v1.Trigger) { list.Items = gentype.FromPointerSlice(items) },
		),
		fake,
	}
}
