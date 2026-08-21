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

var APIHubAPIGVK = GroupVersion.WithKind("APIHubAPI")

// APIHubAPISpec defines the desired state of APIHubAPI
// +kcc:spec:proto=google.cloud.apihub.v1.Api
type APIHubAPISpec struct {
	// The project that this resource belongs to.
	ProjectRef *refsv1beta1.ProjectRef `json:"projectRef"`

	// The location of this resource.
	Location *string `json:"location"`

	// The APIHubAPI name. If not given, the metadata.name will be used.
	ResourceID *string `json:"resourceID,omitempty"`

	// Required. The display name of the API resource.
	// +kubebuilder:validation:Required
	DisplayName *string `json:"displayName,omitempty"`

	// Optional. The description of the API resource.
	// +kubebuilder:validation:Optional
	Description *string `json:"description,omitempty"`

	// Optional. The documentation for the API resource.
	// +kubebuilder:validation:Optional
	Documentation *Documentation `json:"documentation,omitempty"`

	// Optional. Owner details for the API resource.
	// +kubebuilder:validation:Optional
	Owner *Owner `json:"owner,omitempty"`

	// Optional. The target users for the API.
	//  This maps to the following system defined attribute:
	//  `projects/{project}/locations/{location}/attributes/system-target-user`
	// +kubebuilder:validation:Optional
	TargetUserRef *APIHubAttributeValueRef `json:"targetUserRef,omitempty"`

	// Optional. The team owning the API.
	//  This maps to the following system defined attribute:
	//  `projects/{project}/locations/{location}/attributes/system-team`
	// +kubebuilder:validation:Optional
	TeamRef *APIHubAttributeValueRef `json:"teamRef,omitempty"`

	// Optional. The business unit owning the API.
	//  This maps to the following system defined attribute:
	//  `projects/{project}/locations/{location}/attributes/system-business-unit`
	// +kubebuilder:validation:Optional
	BusinessUnitRef *APIHubAttributeValueRef `json:"businessUnitRef,omitempty"`

	// Optional. The maturity level of the API.
	//  This maps to the following system defined attribute:
	//  `projects/{project}/locations/{location}/attributes/system-maturity-level`
	// +kubebuilder:validation:Optional
	MaturityLevelRef *APIHubAttributeValueRef `json:"maturityLevelRef,omitempty"`

	// Optional. The style of the API.
	//  This maps to the following system defined attribute:
	//  `projects/{project}/locations/{location}/attributes/system-api-style`
	// +kubebuilder:validation:Optional
	APIStyleRef *APIHubAttributeValueRef `json:"apiStyleRef,omitempty"`

	// Optional. The selected version for an API resource.
	//  This can be used when special handling is needed on client side for
	//  particular version of the API. Format is
	//  `projects/{project}/locations/{location}/apis/{api}/versions/{version}`
	// +kubebuilder:validation:Optional
	SelectedVersionRef *APIHubVersionRef `json:"selectedVersionRef,omitempty"`

	// Optional. The list of user defined attributes associated with the API
	//  resource. The key is the attribute name. It will be of the format:
	//  `projects/{project}/locations/{location}/attributes/{attribute}`.
	// +kcc:proto:field=google.cloud.apihub.v1.Api.attributes
	// +kubebuilder:validation:Optional
	AttributeRefs []APIHubAPIAttribute `json:"attributeRefs,omitempty"`

	// Optional. The API requirements of the API.
	// +kcc:proto:field=google.cloud.apihub.v1.Api.api_requirements
	// +kubebuilder:validation:Optional
	APIRequirements *AttributeValues `json:"apiRequirements,omitempty"`

	// Optional. The API functional requirements of the API.
	// +kcc:proto:field=google.cloud.apihub.v1.Api.api_functional_requirements
	// +kubebuilder:validation:Optional
	APIFunctionalRequirements *AttributeValues `json:"apiFunctionalRequirements,omitempty"`

	// Optional. The API technical requirements of the API.
	// +kcc:proto:field=google.cloud.apihub.v1.Api.api_technical_requirements
	// +kubebuilder:validation:Optional
	APITechnicalRequirements *AttributeValues `json:"apiTechnicalRequirements,omitempty"`

	// Optional. The base64-encoded fingerprint of the API resource.
	// +kcc:proto:field=google.cloud.apihub.v1.Api.fingerprint
	// +kubebuilder:validation:Optional
	Fingerprint *string `json:"fingerprint,omitempty"`
}

// APIHubAPIStatus defines the config connector machine state of APIHubAPI
type APIHubAPIStatus struct {
	/* Conditions represent the latest available observations of the
	   object's current state. */
	Conditions []v1alpha1.Condition `json:"conditions,omitempty"`

	// ObservedGeneration is the generation of the resource that was most recently observed by the Config Connector controller. If this is equal to metadata.generation, then that means that the current reported status reflects the most recent desired state of the resource.
	ObservedGeneration *int64 `json:"observedGeneration,omitempty"`

	// A unique specifier for the APIHubAPI resource in GCP.
	ExternalRef *string `json:"externalRef,omitempty"`

	// ObservedState is the state of the resource as most recently observed in GCP.
	ObservedState *APIHubAPIObservedState `json:"observedState,omitempty"`
}

// APIHubAPIObservedState is the state of the APIHubAPI resource as most recently observed in GCP.
// +kcc:observedstate:proto=google.cloud.apihub.v1.Api
type APIHubAPIObservedState struct {
	// Output only. The list of versions present in an API resource.
	//  Note: An API resource can be associated with more than 1 version.
	//  Format is
	//  `projects/{project}/locations/{location}/apis/{api}/versions/{version}`
	// +kubebuilder:validation:Optional
	Versions []string `json:"versions,omitempty"`

	// Output only. The time at which the API resource was created.
	// +kubebuilder:validation:Optional
	CreateTime *string `json:"createTime,omitempty"`

	// Output only. The time at which the API resource was last updated.
	// +kubebuilder:validation:Optional
	UpdateTime *string `json:"updateTime,omitempty"`

	// Output only. The metadata describing the source of the API resource.
	// +kcc:proto:field=google.cloud.apihub.v1.Api.source_metadata
	// +kubebuilder:validation:Optional
	SourceMetadata []SourceMetadataObservedState `json:"sourceMetadata,omitempty"`
}

type SourceMetadataObservedState struct {
	// Output only. The type of the source.
	// +kubebuilder:validation:Optional
	SourceType *string `json:"sourceType,omitempty"`

	// Output only. The unique identifier of the resource at the source.
	// +kubebuilder:validation:Optional
	OriginalResourceID *string `json:"originalResourceID,omitempty"`

	// Output only. The time at which the resource was created at the source.
	// +kubebuilder:validation:Optional
	OriginalResourceCreateTime *string `json:"originalResourceCreateTime,omitempty"`

	// Output only. The time at which the resource was last updated at the source.
	// +kubebuilder:validation:Optional
	OriginalResourceUpdateTime *string `json:"originalResourceUpdateTime,omitempty"`
}

// +genclient
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:resource:categories=gcp,shortName=gcpapihubapi;gcpapihubapis
// +kubebuilder:subresource:status
// +kubebuilder:metadata:labels="cnrm.cloud.google.com/managed-by-kcc=true"
// +kubebuilder:metadata:labels="cnrm.cloud.google.com/system=true"
// +kubebuilder:metadata:labels="cnrm.cloud.google.com/stability-level=alpha"
// +kubebuilder:printcolumn:name="Age",JSONPath=".metadata.creationTimestamp",type="date"
// +kubebuilder:printcolumn:name="Ready",JSONPath=".status.conditions[?(@.type=='Ready')].status",type="string",description="When 'True', the most recent reconcile of the resource succeeded"
// +kubebuilder:printcolumn:name="Status",JSONPath=".status.conditions[?(@.type=='Ready')].reason",type="string",description="The reason for the value in 'Ready'"
// +kubebuilder:printcolumn:name="Status Age",JSONPath=".status.conditions[?(@.type=='Ready')].lastTransitionTime",type="date",description="The last transition time for the value in 'Status'"

// APIHubAPI is the Schema for the APIHubAPI API
// +k8s:openapi-gen=true
type APIHubAPI struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	// +required
	Spec   APIHubAPISpec   `json:"spec,omitempty"`
	Status APIHubAPIStatus `json:"status,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// APIHubAPIList contains a list of APIHubAPI
type APIHubAPIList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []APIHubAPI `json:"items"`
}

type APIHubVersionRef struct {
	/* The `id` of a `APIHubVersion` resource, when not managed by Config Connector. */
	External string `json:"external,omitempty"`
	/* The `name` field of a `APIHubVersion` resource. */
	Name string `json:"name,omitempty"`
	/* The `namespace` field of a `APIHubVersion` resource. */
	Namespace string `json:"namespace,omitempty"`
}

type APIHubAttributeRef struct {
	/* The `id` of a `APIHubAttribute` resource, when not managed by Config Connector. */
	External string `json:"external,omitempty"`
	/* The `name` field of a `APIHubAttribute` resource. */
	Name string `json:"name,omitempty"`
	/* The `namespace` field of a `APIHubAttribute` resource. */
	Namespace string `json:"namespace,omitempty"`
}

type APIHubAPIAttribute struct {
	// Reference to the attribute.
	// +kubebuilder:validation:Required
	AttributeRef *APIHubAttributeRef `json:"attributeRef"`
	// The value of the attribute.
	// +kubebuilder:validation:Required
	Values *AttributeValues `json:"values"`
}

func init() {
	SchemeBuilder.Register(&APIHubAPI{}, &APIHubAPIList{})
}
