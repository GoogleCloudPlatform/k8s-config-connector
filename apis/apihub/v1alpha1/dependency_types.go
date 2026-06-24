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

var APIHubDependencyGVK = GroupVersion.WithKind("APIHubDependency")

// +kcc:proto=google.cloud.apihub.v1.DependencyEntityReference
type DependencyEntityReference struct {
	// The resource name of an operation in the API Hub.
	//
	//  Format:
	//  `projects/{project}/locations/{location}/apis/{api}/versions/{version}/operations/{operation}`
	OperationRef *APIHubOperationRef `json:"operationRef,omitempty"`

	// The resource name of an external API in the API Hub.
	//
	//  Format:
	//  `projects/{project}/locations/{location}/externalApis/{external_api}`
	ExternalAPIRef *APIHubExternalAPIRef `json:"externalAPIRef,omitempty"`
}

type APIHubOperationRef struct {
	// A reference to an externally managed API Hub Operation resource.
	// Should be of the format `projects/{project}/locations/{location}/apis/{api}/versions/{version}/operations/{operation}`.
	External string `json:"external,omitempty"`
}

type APIHubExternalAPIRef struct {
	// A reference to an externally managed API Hub External API resource.
	// Should be of the format `projects/{project}/locations/{location}/externalApis/{external_api}`.
	External string `json:"external,omitempty"`
}

type DependencyAttribute struct {
	// The attribute reference.
	AttributeRef APIHubAttributeRef `json:"attributeRef"`

	// The attribute values associated with a resource.
	// +kubebuilder:validation:Required
	Values *AttributeValues `json:"values,omitempty"`
}

// +kcc:proto=google.cloud.apihub.v1.AttributeValues
type AttributeValues struct {
	// The attribute values associated with a resource in case attribute data
	//  type is enum.
	EnumValues *AttributeValues_EnumAttributeValues `json:"enumValues,omitempty"`

	// The attribute values associated with a resource in case attribute data
	//  type is string.
	StringValues *AttributeValues_StringAttributeValues `json:"stringValues,omitempty"`

	// The attribute values associated with a resource in case attribute data
	//  type is JSON.
	// +kcc:proto:field=google.cloud.apihub.v1.AttributeValues.json_values
	JSONValues *AttributeValues_StringAttributeValues `json:"jsonValues,omitempty"`
}

// +kcc:observedstate:proto=google.cloud.apihub.v1.AttributeValues
type AttributeValuesObservedState struct {
	// Output only. The name of the attribute.
	//  Format: projects/{project}/locations/{location}/attributes/{attribute}
	// +kcc:proto:field=google.cloud.apihub.v1.AttributeValues.attribute
	Attribute *string `json:"attribute,omitempty"`
}

// APIHubDependencySpec defines the desired state of APIHubDependency
// +kcc:spec:proto=google.cloud.apihub.v1.Dependency
type APIHubDependencySpec struct {
	// The project that this resource belongs to.
	ProjectRef *refsv1beta1.ProjectRef `json:"projectRef"`

	// The location of this resource.
	// +kubebuilder:validation:Required
	Location *string `json:"location"`

	// The APIHubDependency name. If not given, the metadata.name will be used.
	ResourceID *string `json:"resourceID,omitempty"`

	// Required. Immutable. The entity acting as the consumer in the dependency.
	// +kubebuilder:validation:Required
	Consumer *DependencyEntityReference `json:"consumer,omitempty"`

	// Required. Immutable. The entity acting as the supplier in the dependency.
	// +kubebuilder:validation:Required
	Supplier *DependencyEntityReference `json:"supplier,omitempty"`

	// Optional. Human readable description corresponding of the dependency.
	Description *string `json:"description,omitempty"`

	// Optional. The list of user defined attributes associated with the
	//  dependency resource.
	Attributes []DependencyAttribute `json:"attributes,omitempty"`
}

// +kcc:proto=google.cloud.apihub.v1.DependencyErrorDetail
type DependencyErrorDetail struct {
	// Optional. Error in the dependency.
	Error *string `json:"error,omitempty"`

	// Optional. Timestamp at which the error was found.
	ErrorTime *string `json:"errorTime,omitempty"`
}

// +kcc:observedstate:proto=google.cloud.apihub.v1.DependencyEntityReference
type DependencyEntityReferenceObservedState struct {
	// Output only. Display name of the entity.
	DisplayName *string `json:"displayName,omitempty"`
}

// APIHubDependencyObservedState is the state of the APIHubDependency resource as most recently observed in GCP.
// +kcc:observedstate:proto=google.cloud.apihub.v1.Dependency
type APIHubDependencyObservedState struct {
	// Required. Immutable. The entity acting as the consumer in the dependency.
	Consumer *DependencyEntityReferenceObservedState `json:"consumer,omitempty"`

	// Output only. State of the dependency.
	State *string `json:"state,omitempty"`

	// Output only. Discovery mode of the dependency.
	DiscoveryMode *string `json:"discoveryMode,omitempty"`

	// Output only. Error details of a dependency if the system has detected it
	//  internally.
	ErrorDetail *DependencyErrorDetail `json:"errorDetail,omitempty"`

	// Output only. The time at which the dependency was created.
	CreateTime *string `json:"createTime,omitempty"`

	// Output only. The time at which the dependency was last updated.
	UpdateTime *string `json:"updateTime,omitempty"`
}

// APIHubDependencyStatus defines the config connector machine state of APIHubDependency
type APIHubDependencyStatus struct {
	/* Conditions represent the latest available observations of the
	   object's current state. */
	Conditions []v1alpha1.Condition `json:"conditions,omitempty"`

	// ObservedGeneration is the generation of the resource that was most recently observed by the Config Connector controller. If this is equal to metadata.generation, then that means that the current reported status reflects the most recent desired state of the resource.
	ObservedGeneration *int64 `json:"observedGeneration,omitempty"`

	// A unique specifier for the APIHubDependency resource in GCP.
	ExternalRef *string `json:"externalRef,omitempty"`

	// ObservedState is the state of the resource as most recently observed in GCP.
	ObservedState *APIHubDependencyObservedState `json:"observedState,omitempty"`
}

// +genclient
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:resource:categories=gcp,shortName=gcpapihubdependency;gcpapihubdependencies
// +kubebuilder:subresource:status
// +kubebuilder:metadata:labels="cnrm.cloud.google.com/managed-by-kcc=true"
// +kubebuilder:metadata:labels="cnrm.cloud.google.com/system=true"
// +kubebuilder:printcolumn:name="Age",JSONPath=".metadata.creationTimestamp",type="date"
// +kubebuilder:printcolumn:name="Ready",JSONPath=".status.conditions[?(@.type=='Ready')].status",type="string",description="When 'True', the most recent reconcile of the resource succeeded"
// +kubebuilder:printcolumn:name="Status",JSONPath=".status.conditions[?(@.type=='Ready')].reason",type="string",description="The reason for the value in 'Ready'"
// +kubebuilder:printcolumn:name="Status Age",JSONPath=".status.conditions[?(@.type=='Ready')].lastTransitionTime",type="date",description="The last transition time for the value in 'Status'"

// APIHubDependency is the Schema for the APIHubDependency API
// +k8s:openapi-gen=true
type APIHubDependency struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	// +required
	Spec   APIHubDependencySpec   `json:"spec,omitempty"`
	Status APIHubDependencyStatus `json:"status,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// APIHubDependencyList contains a list of APIHubDependency
type APIHubDependencyList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []APIHubDependency `json:"items"`
}

func init() {
	SchemeBuilder.Register(&APIHubDependency{}, &APIHubDependencyList{})
}
