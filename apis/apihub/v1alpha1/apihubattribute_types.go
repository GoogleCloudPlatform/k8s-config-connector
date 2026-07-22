// Copyright 2026 Google LLC
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//    http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package v1alpha1

import (
	refsv1beta1 "github.com/GoogleCloudPlatform/k8s-config-connector/apis/refs/v1beta1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/apis/k8s/v1alpha1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

var APIHubAttributeGVK = GroupVersion.WithKind("APIHubAttribute")

// APIHubAttributeSpec defines the desired state of APIHubAttribute
// +kcc:spec:proto=google.cloud.apihub.v1.Attribute
type APIHubAttributeSpec struct {
	// The project that this resource belongs to.
	ProjectRef *refsv1beta1.ProjectRef `json:"projectRef"`

	// The location of this resource.
	Location *string `json:"location"`

	// The APIHubAttribute name. If not given, the metadata.name will be used.
	ResourceID *string `json:"resourceID,omitempty"`

	// Required. The display name of the attribute.
	// +kubebuilder:validation:Required
	DisplayName *string `json:"displayName,omitempty"`

	// Optional. The description of the attribute.
	// +kubebuilder:validation:Optional
	Description *string `json:"description,omitempty"`

	// Required. The scope of the attribute. It represents the resource in the API Hub to which the attribute can be linked.
	// +kubebuilder:validation:Required
	// +kubebuilder:validation:Enum=API;VERSION;SPEC;API_OPERATION;DEPLOYMENT;DEPENDENCY;DEFINITION;EXTERNAL_API;PLUGIN
	Scope *string `json:"scope,omitempty"`

	// Required. The type of the data of the attribute.
	// +kubebuilder:validation:Required
	// +kubebuilder:validation:Enum=ENUM;JSON;STRING
	DataType *string `json:"dataType,omitempty"`

	// Optional. The list of allowed values when the attribute value is of type enum. This is required when the dataType of the attribute is ENUM. The maximum number of allowed values of an attribute will be 1000.
	// +kubebuilder:validation:Optional
	AllowedValues []Attribute_AllowedValue `json:"allowedValues,omitempty"`

	// Optional. The maximum number of values that the attribute can have when associated with an API Hub resource. Cardinality 1 would represent a single-valued attribute. It must not be less than 1 or greater than 20. If not specified, the cardinality would be set to 1 by default and represent a single-valued attribute.
	// +kubebuilder:validation:Optional
	Cardinality *int32 `json:"cardinality,omitempty"`
}

// APIHubAttributeStatus defines the config connector machine state of APIHubAttribute
type APIHubAttributeStatus struct {
	/* Conditions represent the latest available observations of the
	   object's current state. */
	Conditions []v1alpha1.Condition `json:"conditions,omitempty"`

	// ObservedGeneration is the generation of the resource that was most recently observed by the Config Connector controller. If this is equal to metadata.generation, then that means that the current reported status reflects the most recent desired state of the resource.
	ObservedGeneration *int64 `json:"observedGeneration,omitempty"`

	// A unique specifier for the APIHubAttribute resource in GCP.
	ExternalRef *string `json:"externalRef,omitempty"`

	// ObservedState is the state of the resource as most recently observed in GCP.
	ObservedState *APIHubAttributeObservedState `json:"observedState,omitempty"`
}

// APIHubAttributeObservedState is the state of the APIHubAttribute resource as most recently observed in GCP.
// +kcc:observedstate:proto=google.cloud.apihub.v1.Attribute
type APIHubAttributeObservedState struct {
	// Output only. The definition type of the attribute.
	DefinitionType *string `json:"definitionType,omitempty"`

	// Output only. When mandatory is true, the attribute is mandatory for the resource specified in the scope. Only System defined attributes can be mandatory.
	Mandatory *bool `json:"mandatory,omitempty"`

	// Output only. The time at which the attribute was created.
	CreateTime *string `json:"createTime,omitempty"`

	// Output only. The time at which the attribute was last updated.
	UpdateTime *string `json:"updateTime,omitempty"`
}

// +genclient
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:resource:categories=gcp,shortName=gcpapihubattribute;gcpapihubattributes
// +kubebuilder:subresource:status
// +kubebuilder:metadata:labels="cnrm.cloud.google.com/managed-by-kcc=true"
// +kubebuilder:metadata:labels="cnrm.cloud.google.com/system=true"
// +kubebuilder:metadata:labels="cnrm.cloud.google.com/stability-level=alpha"
// +kubebuilder:printcolumn:name="Age",JSONPath=".metadata.creationTimestamp",type="date"
// +kubebuilder:printcolumn:name="Ready",JSONPath=".status.conditions[?(@.type=='Ready')].status",type="string",description="When 'True', the most recent reconcile of the resource succeeded"
// +kubebuilder:printcolumn:name="Status",JSONPath=".status.conditions[?(@.type=='Ready')].reason",type="string",description="The reason for the value in 'Ready'"
// +kubebuilder:printcolumn:name="Status Age",JSONPath=".status.conditions[?(@.type=='Ready')].lastTransitionTime",type="date",description="The last transition time for the value in 'Status'"

// APIHubAttribute is the Schema for the APIHubAttribute API
// +k8s:openapi-gen=true
type APIHubAttribute struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	// +required
	Spec   APIHubAttributeSpec   `json:"spec,omitempty"`
	Status APIHubAttributeStatus `json:"status,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// APIHubAttributeList contains a list of APIHubAttribute
type APIHubAttributeList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []APIHubAttribute `json:"items"`
}

func init() {
	SchemeBuilder.Register(&APIHubAttribute{}, &APIHubAttributeList{})
}
