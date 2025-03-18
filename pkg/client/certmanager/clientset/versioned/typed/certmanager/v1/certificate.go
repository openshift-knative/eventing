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
	"context"

	v1 "github.com/cert-manager/cert-manager/pkg/apis/certmanager/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	gentype "k8s.io/client-go/gentype"
	scheme "knative.dev/eventing/pkg/client/certmanager/clientset/versioned/scheme"
)

// CertificatesGetter has a method to return a CertificateInterface.
// A group's client should implement this interface.
type CertificatesGetter interface {
	Certificates(namespace string) CertificateInterface
}

// CertificateInterface has methods to work with Certificate resources.
type CertificateInterface interface {
	Create(ctx context.Context, certificate *v1.Certificate, opts metav1.CreateOptions) (*v1.Certificate, error)
	Update(ctx context.Context, certificate *v1.Certificate, opts metav1.UpdateOptions) (*v1.Certificate, error)
	// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
	UpdateStatus(ctx context.Context, certificate *v1.Certificate, opts metav1.UpdateOptions) (*v1.Certificate, error)
	Delete(ctx context.Context, name string, opts metav1.DeleteOptions) error
	DeleteCollection(ctx context.Context, opts metav1.DeleteOptions, listOpts metav1.ListOptions) error
	Get(ctx context.Context, name string, opts metav1.GetOptions) (*v1.Certificate, error)
	List(ctx context.Context, opts metav1.ListOptions) (*v1.CertificateList, error)
	Watch(ctx context.Context, opts metav1.ListOptions) (watch.Interface, error)
	Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts metav1.PatchOptions, subresources ...string) (result *v1.Certificate, err error)
	CertificateExpansion
}

// certificates implements CertificateInterface
type certificates struct {
	*gentype.ClientWithList[*v1.Certificate, *v1.CertificateList]
}

// newCertificates returns a Certificates
func newCertificates(c *CertmanagerV1Client, namespace string) *certificates {
	return &certificates{
		gentype.NewClientWithList[*v1.Certificate, *v1.CertificateList](
			"certificates",
			c.RESTClient(),
			scheme.ParameterCodec,
			namespace,
			func() *v1.Certificate { return &v1.Certificate{} },
			func() *v1.CertificateList { return &v1.CertificateList{} }),
	}
}
