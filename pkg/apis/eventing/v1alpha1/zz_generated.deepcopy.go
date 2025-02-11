//go:build !ignore_autogenerated
// +build !ignore_autogenerated

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

// Code generated by deepcopy-gen. DO NOT EDIT.

package v1alpha1

import (
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	runtime "k8s.io/apimachinery/pkg/runtime"
	apisduckv1 "knative.dev/eventing/pkg/apis/duck/v1"
	eventingv1 "knative.dev/eventing/pkg/apis/eventing/v1"
	duckv1 "knative.dev/pkg/apis/duck/v1"
)

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *EventPolicy) DeepCopyInto(out *EventPolicy) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new EventPolicy.
func (in *EventPolicy) DeepCopy() *EventPolicy {
	if in == nil {
		return nil
	}
	out := new(EventPolicy)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *EventPolicy) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *EventPolicyFromReference) DeepCopyInto(out *EventPolicyFromReference) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new EventPolicyFromReference.
func (in *EventPolicyFromReference) DeepCopy() *EventPolicyFromReference {
	if in == nil {
		return nil
	}
	out := new(EventPolicyFromReference)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *EventPolicyList) DeepCopyInto(out *EventPolicyList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]EventPolicy, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new EventPolicyList.
func (in *EventPolicyList) DeepCopy() *EventPolicyList {
	if in == nil {
		return nil
	}
	out := new(EventPolicyList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *EventPolicyList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *EventPolicySelector) DeepCopyInto(out *EventPolicySelector) {
	*out = *in
	if in.LabelSelector != nil {
		in, out := &in.LabelSelector, &out.LabelSelector
		*out = new(v1.LabelSelector)
		(*in).DeepCopyInto(*out)
	}
	if in.TypeMeta != nil {
		in, out := &in.TypeMeta, &out.TypeMeta
		*out = new(v1.TypeMeta)
		**out = **in
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new EventPolicySelector.
func (in *EventPolicySelector) DeepCopy() *EventPolicySelector {
	if in == nil {
		return nil
	}
	out := new(EventPolicySelector)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *EventPolicySpec) DeepCopyInto(out *EventPolicySpec) {
	*out = *in
	if in.To != nil {
		in, out := &in.To, &out.To
		*out = make([]EventPolicySpecTo, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.From != nil {
		in, out := &in.From, &out.From
		*out = make([]EventPolicySpecFrom, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.Filters != nil {
		in, out := &in.Filters, &out.Filters
		*out = make([]eventingv1.SubscriptionsAPIFilter, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new EventPolicySpec.
func (in *EventPolicySpec) DeepCopy() *EventPolicySpec {
	if in == nil {
		return nil
	}
	out := new(EventPolicySpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *EventPolicySpecFrom) DeepCopyInto(out *EventPolicySpecFrom) {
	*out = *in
	if in.Ref != nil {
		in, out := &in.Ref, &out.Ref
		*out = new(EventPolicyFromReference)
		**out = **in
	}
	if in.Sub != nil {
		in, out := &in.Sub, &out.Sub
		*out = new(string)
		**out = **in
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new EventPolicySpecFrom.
func (in *EventPolicySpecFrom) DeepCopy() *EventPolicySpecFrom {
	if in == nil {
		return nil
	}
	out := new(EventPolicySpecFrom)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *EventPolicySpecTo) DeepCopyInto(out *EventPolicySpecTo) {
	*out = *in
	if in.Ref != nil {
		in, out := &in.Ref, &out.Ref
		*out = new(EventPolicyToReference)
		**out = **in
	}
	if in.Selector != nil {
		in, out := &in.Selector, &out.Selector
		*out = new(EventPolicySelector)
		(*in).DeepCopyInto(*out)
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new EventPolicySpecTo.
func (in *EventPolicySpecTo) DeepCopy() *EventPolicySpecTo {
	if in == nil {
		return nil
	}
	out := new(EventPolicySpecTo)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *EventPolicyStatus) DeepCopyInto(out *EventPolicyStatus) {
	*out = *in
	in.Status.DeepCopyInto(&out.Status)
	if in.From != nil {
		in, out := &in.From, &out.From
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new EventPolicyStatus.
func (in *EventPolicyStatus) DeepCopy() *EventPolicyStatus {
	if in == nil {
		return nil
	}
	out := new(EventPolicyStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *EventPolicyToReference) DeepCopyInto(out *EventPolicyToReference) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new EventPolicyToReference.
func (in *EventPolicyToReference) DeepCopy() *EventPolicyToReference {
	if in == nil {
		return nil
	}
	out := new(EventPolicyToReference)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *EventTransform) DeepCopyInto(out *EventTransform) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new EventTransform.
func (in *EventTransform) DeepCopy() *EventTransform {
	if in == nil {
		return nil
	}
	out := new(EventTransform)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *EventTransform) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *EventTransformList) DeepCopyInto(out *EventTransformList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]EventTransform, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new EventTransformList.
func (in *EventTransformList) DeepCopy() *EventTransformList {
	if in == nil {
		return nil
	}
	out := new(EventTransformList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *EventTransformList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *EventTransformSpec) DeepCopyInto(out *EventTransformSpec) {
	*out = *in
	if in.Sink != nil {
		in, out := &in.Sink, &out.Sink
		*out = new(duckv1.Destination)
		(*in).DeepCopyInto(*out)
	}
	in.Transformations.DeepCopyInto(&out.Transformations)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new EventTransformSpec.
func (in *EventTransformSpec) DeepCopy() *EventTransformSpec {
	if in == nil {
		return nil
	}
	out := new(EventTransformSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *EventTransformStatus) DeepCopyInto(out *EventTransformStatus) {
	*out = *in
	in.SourceStatus.DeepCopyInto(&out.SourceStatus)
	in.AddressStatus.DeepCopyInto(&out.AddressStatus)
	if in.JsonataTransformationStatus != nil {
		in, out := &in.JsonataTransformationStatus, &out.JsonataTransformationStatus
		*out = new(JsonataEventTransformationStatus)
		(*in).DeepCopyInto(*out)
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new EventTransformStatus.
func (in *EventTransformStatus) DeepCopy() *EventTransformStatus {
	if in == nil {
		return nil
	}
	out := new(EventTransformStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *EventTransformations) DeepCopyInto(out *EventTransformations) {
	*out = *in
	if in.Jsonata != nil {
		in, out := &in.Jsonata, &out.Jsonata
		*out = new(JsonataEventTransformationSpec)
		**out = **in
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new EventTransformations.
func (in *EventTransformations) DeepCopy() *EventTransformations {
	if in == nil {
		return nil
	}
	out := new(EventTransformations)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *JsonataEventTransformationSpec) DeepCopyInto(out *JsonataEventTransformationSpec) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new JsonataEventTransformationSpec.
func (in *JsonataEventTransformationSpec) DeepCopy() *JsonataEventTransformationSpec {
	if in == nil {
		return nil
	}
	out := new(JsonataEventTransformationSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *JsonataEventTransformationStatus) DeepCopyInto(out *JsonataEventTransformationStatus) {
	*out = *in
	in.Deployment.DeepCopyInto(&out.Deployment)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new JsonataEventTransformationStatus.
func (in *JsonataEventTransformationStatus) DeepCopy() *JsonataEventTransformationStatus {
	if in == nil {
		return nil
	}
	out := new(JsonataEventTransformationStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *RequestReply) DeepCopyInto(out *RequestReply) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new RequestReply.
func (in *RequestReply) DeepCopy() *RequestReply {
	if in == nil {
		return nil
	}
	out := new(RequestReply)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *RequestReply) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *RequestReplyList) DeepCopyInto(out *RequestReplyList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]RequestReply, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new RequestReplyList.
func (in *RequestReplyList) DeepCopy() *RequestReplyList {
	if in == nil {
		return nil
	}
	out := new(RequestReplyList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *RequestReplyList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *RequestReplySpec) DeepCopyInto(out *RequestReplySpec) {
	*out = *in
	in.BrokerRef.DeepCopyInto(&out.BrokerRef)
	if in.Timeout != nil {
		in, out := &in.Timeout, &out.Timeout
		*out = new(string)
		**out = **in
	}
	if in.Delivery != nil {
		in, out := &in.Delivery, &out.Delivery
		*out = new(apisduckv1.DeliverySpec)
		(*in).DeepCopyInto(*out)
	}
	if in.Secrets != nil {
		in, out := &in.Secrets, &out.Secrets
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new RequestReplySpec.
func (in *RequestReplySpec) DeepCopy() *RequestReplySpec {
	if in == nil {
		return nil
	}
	out := new(RequestReplySpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *RequestReplyStatus) DeepCopyInto(out *RequestReplyStatus) {
	*out = *in
	in.Status.DeepCopyInto(&out.Status)
	in.AddressStatus.DeepCopyInto(&out.AddressStatus)
	in.AppliedEventPoliciesStatus.DeepCopyInto(&out.AppliedEventPoliciesStatus)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new RequestReplyStatus.
func (in *RequestReplyStatus) DeepCopy() *RequestReplyStatus {
	if in == nil {
		return nil
	}
	out := new(RequestReplyStatus)
	in.DeepCopyInto(out)
	return out
}
