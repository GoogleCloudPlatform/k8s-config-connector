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

var APIHubExternalAPIGVK = GroupVersion.WithKind("APIHubExternalAPI")

// APIHubExternalAPISpec defines the desired state of APIHubExternalAPI
// +kcc:spec:proto=google.cloud.apihub.v1.ExternalApi
type APIHubExternalAPISpec struct {
	// The project that this resource belongs to.
	ProjectRef *refsv1beta1.ProjectRef `json:"projectRef"`

	// The location of this resource.
	Location *string `json:"location"`

	// The APIHubExternalAPI name. If not given, the metadata.name will be used.
	ResourceID *string `json:"resourceID,omitempty"`

	// Required. Display name of the external API. Max length is 63 characters (Unicode Code Points).
	// +kubebuilder:validation:Required
	DisplayName *string `json:"displayName,omitempty"`

	// Optional. Description of the external API. Max length is 2000 characters (Unicode Code Points).
	// +kubebuilder:validation:Optional
	Description *string `json:"description,omitempty"`

	// Optional. List of endpoints on which this API is accessible.
	// +kubebuilder:validation:Optional
	Endpoints []string `json:"endpoints,omitempty"`

	// Optional. List of paths served by this API.
	// +kubebuilder:validation:Optional
	Paths []string `json:"paths,omitempty"`

	// Optional. Documentation of the external API.
	// +kubebuilder:validation:Optional
	Documentation *Documentation `json:"documentation,omitempty"`

	// Optional. The list of user defined attributes associated with the Version resource.
	// The key is the attribute name. It will be of the format: `projects/{project}/locations/{location}/attributes/{attribute}`.
	// The value is the attribute values associated with the resource.
	// +kcc:proto:field=google.cloud.apihub.v1.ExternalApi.attributes
	// +kubebuilder:validation:Optional
	AttributeRefs []APIHubExternalAPIAttribute `json:"attributeRefs,omitempty"`
}

// APIHubExternalAPIStatus defines the config connector machine state of APIHubExternalAPI
type APIHubExternalAPIStatus struct {
	/* Conditions represent the latest available observations of the
	   object's current state. */
	Conditions []v1alpha1.Condition `json:"conditions,omitempty"`

	// ObservedGeneration is the generation of the resource that was most recently observed by the Config Connector controller. If this is equal to metadata.generation, then that means that the current reported status reflects the most recent desired state of the resource.
	ObservedGeneration *int64 `json:"observedGeneration,omitempty"`

	// A unique specifier for the APIHubExternalAPI resource in Google Cloud.
	ExternalRef *string `json:"externalRef,omitempty"`

	// ObservedState is the state of the resource as most recently observed in Google Cloud.
	ObservedState *APIHubExternalAPIObservedState `json:"observedState,omitempty"`
}

// APIHubExternalAPIObservedState is the state of the APIHubExternalAPI resource as most recently observed in Google Cloud.
// +kcc:observedstate:proto=google.cloud.apihub.v1.ExternalApi
type APIHubExternalAPIObservedState struct {
	// Output only. Creation timestamp.
	// +kubebuilder:validation:Optional
	CreateTime *string `json:"createTime,omitempty"`

	// Output only. Last update timestamp.
	// +kubebuilder:validation:Optional
	UpdateTime *string `json:"updateTime,omitempty"`
}

// +genclient
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:resource:categories=gcp,shortName=gcpapihubexternalapi;gcpapihubexternalapis
// +kubebuilder:subresource:status
// +kubebuilder:metadata:labels="cnrm.cloud.google.com/managed-by-kcc=true"
// +kubebuilder:metadata:labels="cnrm.cloud.google.com/system=true"
// +kubebuilder:metadata:labels="cnrm.cloud.google.com/stability-level=alpha"
// +kubebuilder:printcolumn:name="Age",JSONPath=".metadata.creationTimestamp",type="date"
// +kubebuilder:printcolumn:name="Ready",JSONPath=".status.conditions[?(@.type=='Ready')].status",type="string",description="When 'True', the most recent reconcile of the resource succeeded"
// +kubebuilder:printcolumn:name="Status",JSONPath=".status.conditions[?(@.type=='Ready')].reason",type="string",description="The reason for the value in 'Ready'"
// +kubebuilder:printcolumn:name="Status Age",JSONPath=".status.conditions[?(@.type=='Ready')].lastTransitionTime",type="date",description="The last transition time for the value in 'Status'"

// APIHubExternalAPI is the Schema for the APIHubExternalAPI API
// +k8s:openapi-gen=true
type APIHubExternalAPI struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	// +required
	Spec   APIHubExternalAPISpec   `json:"spec,omitempty"`
	Status APIHubExternalAPIStatus `json:"status,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// APIHubExternalAPIList contains a list of APIHubExternalAPI
type APIHubExternalAPIList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []APIHubExternalAPI `json:"items"`
}

type APIHubExternalAPIAttribute struct {
	// Reference to the attribute.
	// +kubebuilder:validation:Required
	AttributeRef *APIHubAttributeRef `json:"attributeRef"`
	// The value of the attribute.
	// +kubebuilder:validation:Required
	Values *AttributeValues `json:"values"`
}

func init() {
	SchemeBuilder.Register(&APIHubExternalAPI{}, &APIHubExternalAPIList{})
}
