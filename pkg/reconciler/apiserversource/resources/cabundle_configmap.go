package resources

import (
	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"knative.dev/pkg/kmeta"
)

const (
	// user-provided and system CA certificates
	TrustedCAConfigMapName   = "config-openshift-trusted-cabundle"
	TrustedCAConfigMapVolume = TrustedCAConfigMapName + "-volume"
	TrustedCAKey             = "ca-bundle.crt"
)

func MakeTrustedCABundleConfigMap(args *ReceiveAdapterArgs) *corev1.ConfigMap {
	return &corev1.ConfigMap{
		ObjectMeta: metav1.ObjectMeta{
			Name:      TrustedCAConfigMapName,
			Namespace: args.Source.Namespace,
			Labels: map[string]string{
				"app.kubernetes.io/name": "knative-eventing",
				// user-provided and system CA certificates
				"config.openshift.io/inject-trusted-cabundle": "true",
			},
			OwnerReferences: []metav1.OwnerReference{
				*kmeta.NewControllerRef(args.Source),
			},
		},
	}
}
