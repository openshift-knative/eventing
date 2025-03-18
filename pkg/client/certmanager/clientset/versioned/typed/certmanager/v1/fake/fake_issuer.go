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
	"context"

	v1 "github.com/cert-manager/cert-manager/pkg/apis/certmanager/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	labels "k8s.io/apimachinery/pkg/labels"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	testing "k8s.io/client-go/testing"
)

// FakeIssuers implements IssuerInterface
type FakeIssuers struct {
	Fake *FakeCertmanagerV1
	ns   string
}

var issuersResource = v1.SchemeGroupVersion.WithResource("issuers")

var issuersKind = v1.SchemeGroupVersion.WithKind("Issuer")

// Get takes name of the issuer, and returns the corresponding issuer object, and an error if there is any.
func (c *FakeIssuers) Get(ctx context.Context, name string, options metav1.GetOptions) (result *v1.Issuer, err error) {
	emptyResult := &v1.Issuer{}
	obj, err := c.Fake.
		Invokes(testing.NewGetActionWithOptions(issuersResource, c.ns, name, options), emptyResult)

	if obj == nil {
		return emptyResult, err
	}
	return obj.(*v1.Issuer), err
}

// List takes label and field selectors, and returns the list of Issuers that match those selectors.
func (c *FakeIssuers) List(ctx context.Context, opts metav1.ListOptions) (result *v1.IssuerList, err error) {
	emptyResult := &v1.IssuerList{}
	obj, err := c.Fake.
		Invokes(testing.NewListActionWithOptions(issuersResource, issuersKind, c.ns, opts), emptyResult)

	if obj == nil {
		return emptyResult, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &v1.IssuerList{ListMeta: obj.(*v1.IssuerList).ListMeta}
	for _, item := range obj.(*v1.IssuerList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested issuers.
func (c *FakeIssuers) Watch(ctx context.Context, opts metav1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewWatchActionWithOptions(issuersResource, c.ns, opts))

}

// Create takes the representation of a issuer and creates it.  Returns the server's representation of the issuer, and an error, if there is any.
func (c *FakeIssuers) Create(ctx context.Context, issuer *v1.Issuer, opts metav1.CreateOptions) (result *v1.Issuer, err error) {
	emptyResult := &v1.Issuer{}
	obj, err := c.Fake.
		Invokes(testing.NewCreateActionWithOptions(issuersResource, c.ns, issuer, opts), emptyResult)

	if obj == nil {
		return emptyResult, err
	}
	return obj.(*v1.Issuer), err
}

// Update takes the representation of a issuer and updates it. Returns the server's representation of the issuer, and an error, if there is any.
func (c *FakeIssuers) Update(ctx context.Context, issuer *v1.Issuer, opts metav1.UpdateOptions) (result *v1.Issuer, err error) {
	emptyResult := &v1.Issuer{}
	obj, err := c.Fake.
		Invokes(testing.NewUpdateActionWithOptions(issuersResource, c.ns, issuer, opts), emptyResult)

	if obj == nil {
		return emptyResult, err
	}
	return obj.(*v1.Issuer), err
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *FakeIssuers) UpdateStatus(ctx context.Context, issuer *v1.Issuer, opts metav1.UpdateOptions) (result *v1.Issuer, err error) {
	emptyResult := &v1.Issuer{}
	obj, err := c.Fake.
		Invokes(testing.NewUpdateSubresourceActionWithOptions(issuersResource, "status", c.ns, issuer, opts), emptyResult)

	if obj == nil {
		return emptyResult, err
	}
	return obj.(*v1.Issuer), err
}

// Delete takes name of the issuer and deletes it. Returns an error if one occurs.
func (c *FakeIssuers) Delete(ctx context.Context, name string, opts metav1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewDeleteActionWithOptions(issuersResource, c.ns, name, opts), &v1.Issuer{})

	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeIssuers) DeleteCollection(ctx context.Context, opts metav1.DeleteOptions, listOpts metav1.ListOptions) error {
	action := testing.NewDeleteCollectionActionWithOptions(issuersResource, c.ns, opts, listOpts)

	_, err := c.Fake.Invokes(action, &v1.IssuerList{})
	return err
}

// Patch applies the patch and returns the patched issuer.
func (c *FakeIssuers) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts metav1.PatchOptions, subresources ...string) (result *v1.Issuer, err error) {
	emptyResult := &v1.Issuer{}
	obj, err := c.Fake.
		Invokes(testing.NewPatchSubresourceActionWithOptions(issuersResource, c.ns, name, pt, data, opts, subresources...), emptyResult)

	if obj == nil {
		return emptyResult, err
	}
	return obj.(*v1.Issuer), err
}
