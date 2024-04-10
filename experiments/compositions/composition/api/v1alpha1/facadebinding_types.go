/*
Copyright 2024.

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

package v1alpha1

import (
	apiextensionsv1 "k8s.io/apiextensions-apiserver/pkg/apis/apiextensions/v1"
	"k8s.io/apimachinery/pkg/api/meta"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// FacadeBindingSpec defines the desired state of FacadeBinding
type FacadeBindingSpec struct {
	// Use existing KRM API created offline
	FacadeAPI string `json:"facadeAPI,omitempty"`
	// Bring your own open API spec
	// +kubebuilder:validation:Schemaless
	// +kubebuilder:validation:Type=object
	// +kubebuilder:pruning:PreserveUnknownFields
	// +k8s:conversion-gen=false
	OpenAPIV3Schema      *apiextensionsv1.JSONSchemaProps `json:"openAPIV3Schema,omitempty"`
	FacadeKind           string                           `json:"facadeKind,omitempty"`
	CompositionName      string                           `json:"compositionName"`
	CompositionNamespace string                           `json:"compositionNamespace"`
	Description          string                           `json:"description,omitempty"`
}

// FacadeBindingStatus defines the observed state of FacadeBinding
type FacadeBindingStatus struct {
	Conditions []metav1.Condition `json:"conditions,omitempty"`
}

// Status helpers
func (s *FacadeBindingStatus) ClearCondition(condition ConditionType) {
	meta.RemoveStatusCondition(&s.Conditions, string(condition))
}

//+kubebuilder:object:root=true
//+kubebuilder:subresource:status

// FacadeBinding is the Schema for the facadebindings API
type FacadeBinding struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   FacadeBindingSpec   `json:"spec,omitempty"`
	Status FacadeBindingStatus `json:"status,omitempty"`
}

// Validation
func (f *FacadeBinding) Validate() bool {
	return true
}

//+kubebuilder:object:root=true

// FacadeBindingList contains a list of FacadeBinding
type FacadeBindingList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []FacadeBinding `json:"items"`
}

func init() {
	SchemeBuilder.Register(&FacadeBinding{}, &FacadeBindingList{})
}
