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
	v1alpha1 "knative.dev/eventing/pkg/apis/sinks/v1alpha1"
	sinksv1alpha1 "knative.dev/eventing/pkg/client/clientset/versioned/typed/sinks/v1alpha1"
)

// fakeIntegrationSinks implements IntegrationSinkInterface
type fakeIntegrationSinks struct {
	*gentype.FakeClientWithList[*v1alpha1.IntegrationSink, *v1alpha1.IntegrationSinkList]
	Fake *FakeSinksV1alpha1
}

func newFakeIntegrationSinks(fake *FakeSinksV1alpha1, namespace string) sinksv1alpha1.IntegrationSinkInterface {
	return &fakeIntegrationSinks{
		gentype.NewFakeClientWithList[*v1alpha1.IntegrationSink, *v1alpha1.IntegrationSinkList](
			fake.Fake,
			namespace,
			v1alpha1.SchemeGroupVersion.WithResource("integrationsinks"),
			v1alpha1.SchemeGroupVersion.WithKind("IntegrationSink"),
			func() *v1alpha1.IntegrationSink { return &v1alpha1.IntegrationSink{} },
			func() *v1alpha1.IntegrationSinkList { return &v1alpha1.IntegrationSinkList{} },
			func(dst, src *v1alpha1.IntegrationSinkList) { dst.ListMeta = src.ListMeta },
			func(list *v1alpha1.IntegrationSinkList) []*v1alpha1.IntegrationSink {
				return gentype.ToPointerSlice(list.Items)
			},
			func(list *v1alpha1.IntegrationSinkList, items []*v1alpha1.IntegrationSink) {
				list.Items = gentype.FromPointerSlice(items)
			},
		),
		fake,
	}
}
