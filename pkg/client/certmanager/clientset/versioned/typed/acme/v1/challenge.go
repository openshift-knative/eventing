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

package v1

import (
	context "context"

	acmev1 "github.com/cert-manager/cert-manager/pkg/apis/acme/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	gentype "k8s.io/client-go/gentype"
	scheme "knative.dev/eventing/pkg/client/certmanager/clientset/versioned/scheme"
)

// ChallengesGetter has a method to return a ChallengeInterface.
// A group's client should implement this interface.
type ChallengesGetter interface {
	Challenges(namespace string) ChallengeInterface
}

// ChallengeInterface has methods to work with Challenge resources.
type ChallengeInterface interface {
	Create(ctx context.Context, challenge *acmev1.Challenge, opts metav1.CreateOptions) (*acmev1.Challenge, error)
	Update(ctx context.Context, challenge *acmev1.Challenge, opts metav1.UpdateOptions) (*acmev1.Challenge, error)
	// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
	UpdateStatus(ctx context.Context, challenge *acmev1.Challenge, opts metav1.UpdateOptions) (*acmev1.Challenge, error)
	Delete(ctx context.Context, name string, opts metav1.DeleteOptions) error
	DeleteCollection(ctx context.Context, opts metav1.DeleteOptions, listOpts metav1.ListOptions) error
	Get(ctx context.Context, name string, opts metav1.GetOptions) (*acmev1.Challenge, error)
	List(ctx context.Context, opts metav1.ListOptions) (*acmev1.ChallengeList, error)
	Watch(ctx context.Context, opts metav1.ListOptions) (watch.Interface, error)
	Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts metav1.PatchOptions, subresources ...string) (result *acmev1.Challenge, err error)
	ChallengeExpansion
}

// challenges implements ChallengeInterface
type challenges struct {
	*gentype.ClientWithList[*acmev1.Challenge, *acmev1.ChallengeList]
}

// newChallenges returns a Challenges
func newChallenges(c *AcmeV1Client, namespace string) *challenges {
	return &challenges{
		gentype.NewClientWithList[*acmev1.Challenge, *acmev1.ChallengeList](
			"challenges",
			c.RESTClient(),
			scheme.ParameterCodec,
			namespace,
			func() *acmev1.Challenge { return &acmev1.Challenge{} },
			func() *acmev1.ChallengeList { return &acmev1.ChallengeList{} },
		),
	}
}
