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

var ApiHubApiGVK = GroupVersion.WithKind("ApiHubApi")

// ApiHubApiSpec defines the desired state of ApiHubApi
// +kcc:spec:proto=google.cloud.apihub.v1.Api
type ApiHubApiSpec struct {
	// The project that this resource belongs to.
	ProjectRef *refsv1beta1.ProjectRef `json:"projectRef"`

	// The location of this resource.
	Location *string `json:"location"`

	// The ApiHubApi name. If not given, the metadata.name will be used.
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
	SelectedVersion *string `json:"selectedVersion,omitempty"`
}

// ApiHubApiStatus defines the config connector machine state of ApiHubApi
type ApiHubApiStatus struct {
	/* Conditions represent the latest available observations of the
	   object's current state. */
	Conditions []v1alpha1.Condition `json:"conditions,omitempty"`

	// ObservedGeneration is the generation of the resource that was most recently observed by the Config Connector controller. If this is equal to metadata.generation, then that means that the current reported status reflects the most recent desired state of the resource.
	ObservedGeneration *int64 `json:"observedGeneration,omitempty"`

	// A unique specifier for the ApiHubApi resource in GCP.
	ExternalRef *string `json:"externalRef,omitempty"`

	// ObservedState is the state of the resource as most recently observed in GCP.
	ObservedState *ApiHubApiObservedState `json:"observedState,omitempty"`
}

// ApiHubApiObservedState is the state of the ApiHubApi resource as most recently observed in GCP.
// +kcc:observedstate:proto=google.cloud.apihub.v1.Api
type ApiHubApiObservedState struct {
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

// ApiHubApi is the Schema for the ApiHubApi API
// +k8s:openapi-gen=true
type ApiHubApi struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	// +required
	Spec   ApiHubApiSpec   `json:"spec,omitempty"`
	Status ApiHubApiStatus `json:"status,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// ApiHubApiList contains a list of ApiHubApi
type ApiHubApiList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []ApiHubApi `json:"items"`
}

func init() {
	SchemeBuilder.Register(&ApiHubApi{}, &ApiHubApiList{})
}
