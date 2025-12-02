// Copyright 2025 Google LLC
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
	"github.com/GoogleCloudPlatform/k8s-config-connector/apis/refs/v1beta1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/apis/k8s/v1alpha1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

var APIHubAPIGVK = GroupVersion.WithKind("APIHubAPI")

type APIParent struct {
	// Required. The location of the API.
	// +required
	Location string `json:"location,omitempty"`

	// Required. The host project of the API.
	// +required
	ProjectRef *v1beta1.ProjectRef `json:"projectRef,omitempty"`
}

// APIHubAPISpec defines the desired state of APIHubAPI
// +kcc:spec:proto=google.cloud.apihub.v1.Api
type APIHubAPISpec struct {
	APIParent `json:",inline"`

	// The APIHubAPI name. If not given, the metadata.name will be used.
	ResourceID *string `json:"resourceID,omitempty"`

	// Required. The display name of the API resource.
	// +kcc:proto:field=google.cloud.apihub.v1.Api.display_name
	// +required
	DisplayName *string `json:"displayName,omitempty"`

	// Optional. The description of the API resource.
	// +kcc:proto:field=google.cloud.apihub.v1.Api.description
	Description *string `json:"description,omitempty"`

	// Optional. The documentation for the API resource.
	// +kcc:proto:field=google.cloud.apihub.v1.Api.documentation
	Documentation *Documentation `json:"documentation,omitempty"`

	// Optional. Owner details for the API resource.
	// +kcc:proto:field=google.cloud.apihub.v1.Api.owner
	Owner *Owner `json:"owner,omitempty"`

	// Optional. The target users for the API.
	//  This maps to the following system defined attribute:
	//  `projects/{project}/locations/{location}/attributes/system-target-user`
	//  attribute.
	//  The number of values for this attribute will be based on the
	//  cardinality of the attribute. The same can be retrieved via GetAttribute
	//  API. All values should be from the list of allowed values defined for the
	//  attribute.
	// +kcc:proto:field=google.cloud.apihub.v1.Api.target_user
	TargetUser *AttributeValues `json:"targetUser,omitempty"`

	// Optional. The team owning the API.
	//  This maps to the following system defined attribute:
	//  `projects/{project}/locations/{location}/attributes/system-team`
	//  attribute.
	//  The number of values for this attribute will be based on the
	//  cardinality of the attribute. The same can be retrieved via GetAttribute
	//  API. All values should be from the list of allowed values defined for the
	//  attribute.
	// +kcc:proto:field=google.cloud.apihub.v1.Api.team
	Team *AttributeValues `json:"team,omitempty"`

	// Optional. The business unit owning the API.
	//  This maps to the following system defined attribute:
	//  `projects/{project}/locations/{location}/attributes/system-business-unit`
	//  attribute.
	//  The number of values for this attribute will be based on the
	//  cardinality of the attribute. The same can be retrieved via GetAttribute
	//  API. All values should be from the list of allowed values defined for the
	//  attribute.
	// +kcc:proto:field=google.cloud.apihub.v1.Api.business_unit
	BusinessUnit *AttributeValues `json:"businessUnit,omitempty"`

	// Optional. The maturity level of the API.
	//  This maps to the following system defined attribute:
	//  `projects/{project}/locations/{location}/attributes/system-maturity-level`
	//  attribute.
	//  The number of values for this attribute will be based on the
	//  cardinality of the attribute. The same can be retrieved via GetAttribute
	//  API. All values should be from the list of allowed values defined for the
	//  attribute.
	// +kcc:proto:field=google.cloud.apihub.v1.Api.maturity_level
	MaturityLevel *AttributeValues `json:"maturityLevel,omitempty"`

	// TODO: unsupported map type with key string and value message

	// Optional. The style of the API.
	//  This maps to the following system defined attribute:
	//  `projects/{project}/locations/{location}/attributes/system-api-style`
	//  attribute.
	//  The number of values for this attribute will be based on the
	//  cardinality of the attribute. The same can be retrieved via GetAttribute
	//  API. All values should be from the list of allowed values defined for the
	//  attribute.
	// +kcc:proto:field=google.cloud.apihub.v1.Api.api_style
	APIStyle *AttributeValues `json:"apiStyle,omitempty"`

	// Optional. The selected version for an API resource.
	//  This can be used when special handling is needed on client side for
	//  particular version of the API. Format is
	//  `projects/{project}/locations/{location}/apis/{api}/versions/{version}`
	// +kcc:proto:field=google.cloud.apihub.v1.Api.selected_version
	SelectedVersion *string `json:"selectedVersion,omitempty"`
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
	// +kcc:proto:field=google.cloud.apihub.v1.Api.versions
	Versions []string `json:"versions,omitempty"`

	// Output only. The time at which the API resource was created.
	// +kcc:proto:field=google.cloud.apihub.v1.Api.create_time
	CreateTime *string `json:"createTime,omitempty"`

	// Output only. The time at which the API resource was last updated.
	// +kcc:proto:field=google.cloud.apihub.v1.Api.update_time
	UpdateTime *string `json:"updateTime,omitempty"`

	// Optional. The target users for the API.
	//  This maps to the following system defined attribute:
	//  `projects/{project}/locations/{location}/attributes/system-target-user`
	//  attribute.
	//  The number of values for this attribute will be based on the
	//  cardinality of the attribute. The same can be retrieved via GetAttribute
	//  API. All values should be from the list of allowed values defined for the
	//  attribute.
	// +kcc:proto:field=google.cloud.apihub.v1.Api.target_user
	TargetUser *AttributeValuesObservedState `json:"targetUser,omitempty"`
}

// +genclient
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:resource:categories=gcp,shortName=gcpapihubapi;gcpapihubapis
// +kubebuilder:subresource:status
// +kubebuilder:metadata:labels="cnrm.cloud.google.com/managed-by-kcc=true";"cnrm.cloud.google.com/system=true"
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

func init() {
	SchemeBuilder.Register(&APIHubAPI{}, &APIHubAPIList{})
}
