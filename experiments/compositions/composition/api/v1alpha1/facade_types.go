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
	runtime "k8s.io/apimachinery/pkg/runtime"
)

type SimpleSchema struct {
	Spec     runtime.RawExtension `json:"spec"`
	Status   runtime.RawExtension `json:"status"`
	Required []string             `json:"required,omitempty"`
}

// FacadeSpec defines the desired state of Facade
type FacadeSpec struct {
	// Bring your own open API spec
	// +kubebuilder:validation:Schemaless
	// +kubebuilder:validation:Type=object
	// +kubebuilder:pruning:PreserveUnknownFields
	// +k8s:conversion-gen=false
	OpenAPIV3Schema *apiextensionsv1.JSONSchemaProps `json:"openAPIV3Schema,omitempty"`

	// Simple Schema
	SimpleSchema *SimpleSchema `json:"simpleSchema,omitempty"`

	FacadeKind string `json:"facadeKind"`
}

// FacadeStatus defines the observed state of Facade
type FacadeStatus struct {
	Conditions []metav1.Condition `json:"conditions,omitempty"`
}

//+kubebuilder:object:root=true
//+kubebuilder:subresource:status

// Facade is the Schema for the facades API
type Facade struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   FacadeSpec   `json:"spec,omitempty"`
	Status FacadeStatus `json:"status,omitempty"`
}

//+kubebuilder:object:root=true

// FacadeList contains a list of Facade
type FacadeList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []Facade `json:"items"`
}

// Status helpers
func (s *FacadeStatus) ClearCondition(condition ConditionType) {
	meta.RemoveStatusCondition(&s.Conditions, string(condition))
}

// Validation
func (f *Facade) Validate() bool {
	return true
}

func init() {
	SchemeBuilder.Register(&Facade{}, &FacadeList{})
}
