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

package v1alpha1

import (
	"context"
	"time"

	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	rest "k8s.io/client-go/rest"
	v1alpha1 "knative.dev/eventing/pkg/apis/eventing/v1alpha1"
	scheme "knative.dev/eventing/pkg/client/clientset/versioned/scheme"
)

// EventTransformsGetter has a method to return a EventTransformInterface.
// A group's client should implement this interface.
type EventTransformsGetter interface {
	EventTransforms(namespace string) EventTransformInterface
}

// EventTransformInterface has methods to work with EventTransform resources.
type EventTransformInterface interface {
	Create(ctx context.Context, eventTransform *v1alpha1.EventTransform, opts v1.CreateOptions) (*v1alpha1.EventTransform, error)
	Update(ctx context.Context, eventTransform *v1alpha1.EventTransform, opts v1.UpdateOptions) (*v1alpha1.EventTransform, error)
	UpdateStatus(ctx context.Context, eventTransform *v1alpha1.EventTransform, opts v1.UpdateOptions) (*v1alpha1.EventTransform, error)
	Delete(ctx context.Context, name string, opts v1.DeleteOptions) error
	DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error
	Get(ctx context.Context, name string, opts v1.GetOptions) (*v1alpha1.EventTransform, error)
	List(ctx context.Context, opts v1.ListOptions) (*v1alpha1.EventTransformList, error)
	Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error)
	Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *v1alpha1.EventTransform, err error)
	EventTransformExpansion
}

// eventTransforms implements EventTransformInterface
type eventTransforms struct {
	client rest.Interface
	ns     string
}

// newEventTransforms returns a EventTransforms
func newEventTransforms(c *EventingV1alpha1Client, namespace string) *eventTransforms {
	return &eventTransforms{
		client: c.RESTClient(),
		ns:     namespace,
	}
}

// Get takes name of the eventTransform, and returns the corresponding eventTransform object, and an error if there is any.
func (c *eventTransforms) Get(ctx context.Context, name string, options v1.GetOptions) (result *v1alpha1.EventTransform, err error) {
	result = &v1alpha1.EventTransform{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("eventtransforms").
		Name(name).
		VersionedParams(&options, scheme.ParameterCodec).
		Do(ctx).
		Into(result)
	return
}

// List takes label and field selectors, and returns the list of EventTransforms that match those selectors.
func (c *eventTransforms) List(ctx context.Context, opts v1.ListOptions) (result *v1alpha1.EventTransformList, err error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	result = &v1alpha1.EventTransformList{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("eventtransforms").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Do(ctx).
		Into(result)
	return
}

// Watch returns a watch.Interface that watches the requested eventTransforms.
func (c *eventTransforms) Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	opts.Watch = true
	return c.client.Get().
		Namespace(c.ns).
		Resource("eventtransforms").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Watch(ctx)
}

// Create takes the representation of a eventTransform and creates it.  Returns the server's representation of the eventTransform, and an error, if there is any.
func (c *eventTransforms) Create(ctx context.Context, eventTransform *v1alpha1.EventTransform, opts v1.CreateOptions) (result *v1alpha1.EventTransform, err error) {
	result = &v1alpha1.EventTransform{}
	err = c.client.Post().
		Namespace(c.ns).
		Resource("eventtransforms").
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(eventTransform).
		Do(ctx).
		Into(result)
	return
}

// Update takes the representation of a eventTransform and updates it. Returns the server's representation of the eventTransform, and an error, if there is any.
func (c *eventTransforms) Update(ctx context.Context, eventTransform *v1alpha1.EventTransform, opts v1.UpdateOptions) (result *v1alpha1.EventTransform, err error) {
	result = &v1alpha1.EventTransform{}
	err = c.client.Put().
		Namespace(c.ns).
		Resource("eventtransforms").
		Name(eventTransform.Name).
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(eventTransform).
		Do(ctx).
		Into(result)
	return
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *eventTransforms) UpdateStatus(ctx context.Context, eventTransform *v1alpha1.EventTransform, opts v1.UpdateOptions) (result *v1alpha1.EventTransform, err error) {
	result = &v1alpha1.EventTransform{}
	err = c.client.Put().
		Namespace(c.ns).
		Resource("eventtransforms").
		Name(eventTransform.Name).
		SubResource("status").
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(eventTransform).
		Do(ctx).
		Into(result)
	return
}

// Delete takes name of the eventTransform and deletes it. Returns an error if one occurs.
func (c *eventTransforms) Delete(ctx context.Context, name string, opts v1.DeleteOptions) error {
	return c.client.Delete().
		Namespace(c.ns).
		Resource("eventtransforms").
		Name(name).
		Body(&opts).
		Do(ctx).
		Error()
}

// DeleteCollection deletes a collection of objects.
func (c *eventTransforms) DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error {
	var timeout time.Duration
	if listOpts.TimeoutSeconds != nil {
		timeout = time.Duration(*listOpts.TimeoutSeconds) * time.Second
	}
	return c.client.Delete().
		Namespace(c.ns).
		Resource("eventtransforms").
		VersionedParams(&listOpts, scheme.ParameterCodec).
		Timeout(timeout).
		Body(&opts).
		Do(ctx).
		Error()
}

// Patch applies the patch and returns the patched eventTransform.
func (c *eventTransforms) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *v1alpha1.EventTransform, err error) {
	result = &v1alpha1.EventTransform{}
	err = c.client.Patch(pt).
		Namespace(c.ns).
		Resource("eventtransforms").
		Name(name).
		SubResource(subresources...).
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(data).
		Do(ctx).
		Into(result)
	return
}
