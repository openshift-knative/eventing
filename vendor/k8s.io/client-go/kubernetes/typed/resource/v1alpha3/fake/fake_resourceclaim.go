/*
Copyright The Kubernetes Authors.

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
	v1alpha3 "k8s.io/api/resource/v1alpha3"
	resourcev1alpha3 "k8s.io/client-go/applyconfigurations/resource/v1alpha3"
	gentype "k8s.io/client-go/gentype"
	typedresourcev1alpha3 "k8s.io/client-go/kubernetes/typed/resource/v1alpha3"
)

// fakeResourceClaims implements ResourceClaimInterface
type fakeResourceClaims struct {
	*gentype.FakeClientWithListAndApply[*v1alpha3.ResourceClaim, *v1alpha3.ResourceClaimList, *resourcev1alpha3.ResourceClaimApplyConfiguration]
	Fake *FakeResourceV1alpha3
}

func newFakeResourceClaims(fake *FakeResourceV1alpha3, namespace string) typedresourcev1alpha3.ResourceClaimInterface {
	return &fakeResourceClaims{
		gentype.NewFakeClientWithListAndApply[*v1alpha3.ResourceClaim, *v1alpha3.ResourceClaimList, *resourcev1alpha3.ResourceClaimApplyConfiguration](
			fake.Fake,
			namespace,
			v1alpha3.SchemeGroupVersion.WithResource("resourceclaims"),
			v1alpha3.SchemeGroupVersion.WithKind("ResourceClaim"),
			func() *v1alpha3.ResourceClaim { return &v1alpha3.ResourceClaim{} },
			func() *v1alpha3.ResourceClaimList { return &v1alpha3.ResourceClaimList{} },
			func(dst, src *v1alpha3.ResourceClaimList) { dst.ListMeta = src.ListMeta },
			func(list *v1alpha3.ResourceClaimList) []*v1alpha3.ResourceClaim {
				return gentype.ToPointerSlice(list.Items)
			},
			func(list *v1alpha3.ResourceClaimList, items []*v1alpha3.ResourceClaim) {
				list.Items = gentype.FromPointerSlice(items)
			},
		),
		fake,
	}
}
