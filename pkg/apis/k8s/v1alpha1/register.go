// Package v1alpha1 contains API Schema definitions for the k8s v1alpha1 API group
// +k8s:openapi-gen=false
// +k8s:deepcopy-gen=
// +k8s:conversion-gen=cnrm.googlesource.com/cnrm/pkg/apis/k8s
// +k8s:defaulter-gen=TypeMeta
// +groupName=k8s.cnrm.cloud.google.com
package v1alpha1

import (
	"k8s.io/apimachinery/pkg/runtime/schema"
	"sigs.k8s.io/controller-runtime/pkg/scheme"
)

var (
	// SchemeGroupVersion is group version used to register these objects
	SchemeGroupVersion = schema.GroupVersion{Group: "k8s.cnrm.cloud.google.com", Version: "v1alpha1"}

	// SchemeBuilder is used to add go types to the GroupVersionKind scheme
	SchemeBuilder = &scheme.Builder{GroupVersion: SchemeGroupVersion}

	// AddToScheme is a global function that registers this API group & version to a scheme
	AddToScheme = SchemeBuilder.AddToScheme
)
