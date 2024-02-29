// Package v1beta1 contains API Schema definitions for the securesourcemanager v1alpha1 API group
package v1alpha1

import (
	"reflect"

	"k8s.io/apimachinery/pkg/runtime/schema"
	"sigs.k8s.io/controller-runtime/pkg/scheme"
)

// +kubebuilder:object:generate=true
// +groupName=securesourcemanager.cnrm.cloud.google.com

//go:generate go run sigs.k8s.io/controller-tools/cmd/controller-gen@v0.14.0 output:artifacts:code=.,config=../config/crds object crd:crdVersions=v1 paths=.

var (
	GroupVersion  = schema.GroupVersion{Group: "securesourcemanager.cnrm.cloud.google.com", Version: "v1alpha1"}
	SchemeBuilder = &scheme.Builder{GroupVersion: GroupVersion}
	AddToScheme   = SchemeBuilder.AddToScheme

	SecureSourceManagerInstanceGVK = schema.GroupVersionKind{
		Group:   GroupVersion.Group,
		Version: GroupVersion.Version,
		Kind:    reflect.TypeOf(SecureSourceManagerInstance{}).Name(),
	}
)
