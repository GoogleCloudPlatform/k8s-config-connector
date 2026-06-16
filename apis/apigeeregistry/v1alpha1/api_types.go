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
	refsv1beta1 "github.com/GoogleCloudPlatform/k8s-config-connector/apis/refs/v1beta1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/apis/k8s/v1alpha1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

var ApigeeRegistryAPIGVK = GroupVersion.WithKind("ApigeeRegistryAPI")

// ApigeeRegistryAPISpec defines the desired state of ApigeeRegistryAPI
// +kcc:spec:proto=google.cloud.apigeeregistry.v1.Api
type ApigeeRegistryAPISpec struct {
	// The project that this resource belongs to.
	ProjectRef *refsv1beta1.ProjectRef `json:"projectRef"`

	// The location of this resource.
	Location *string `json:"location"`

	// The ApigeeRegistryAPI name. If not given, the metadata.name will be used.
	ResourceID *string `json:"resourceID,omitempty"`

	// Human-meaningful name.
	// +kcc:proto:field=google.cloud.apigeeregistry.v1.Api.display_name
	DisplayName *string `json:"displayName,omitempty"`

	// A detailed description.
	// +kcc:proto:field=google.cloud.apigeeregistry.v1.Api.description
	Description *string `json:"description,omitempty"`

	// A user-definable description of the availability of this service.
	//  Format: free-form, but we expect single words that describe availability,
	//  e.g., "NONE", "TESTING", "PREVIEW", "GENERAL", "DEPRECATED", "SHUTDOWN".
	// +kcc:proto:field=google.cloud.apigeeregistry.v1.Api.availability
	Availability *string `json:"availability,omitempty"`

	// The recommended version of the API.
	//  Format: `apis/{api}/versions/{version}`
	// +kcc:proto:field=google.cloud.apigeeregistry.v1.Api.recommended_version
	RecommendedVersion *string `json:"recommendedVersion,omitempty"`

	// The recommended deployment of the API.
	//  Format: `apis/{api}/deployments/{deployment}`
	// +kcc:proto:field=google.cloud.apigeeregistry.v1.Api.recommended_deployment
	RecommendedDeployment *string `json:"recommendedDeployment,omitempty"`

	// Labels attach identifying metadata to resources. Identifying metadata can
	//  be used to filter list operations.
	//
	//  Label keys and values can be no longer than 64 characters
	//  (Unicode codepoints), can only contain lowercase letters, numeric
	//  characters, underscores, and dashes. International characters are allowed.
	//  No more than 64 user labels can be associated with one resource (System
	//  labels are excluded).
	//
	//  See https://goo.gl/xmQnxf for more information and examples of labels.
	//  System reserved label keys are prefixed with
	//  `apigeeregistry.googleapis.com/` and cannot be changed.
	// +kcc:proto:field=google.cloud.apigeeregistry.v1.Api.labels
	Labels map[string]string `json:"labels,omitempty"`

	// Annotations attach non-identifying metadata to resources.
	//
	//  Annotation keys and values are less restricted than those of labels, but
	//  should be generally used for small values of broad interest. Larger, topic-
	//  specific metadata should be stored in Artifacts.
	// +kcc:proto:field=google.cloud.apigeeregistry.v1.Api.annotations
	Annotations map[string]string `json:"annotations,omitempty"`
}

// ApigeeRegistryAPIStatus defines the config connector machine state of ApigeeRegistryAPI
type ApigeeRegistryAPIStatus struct {
	/* Conditions represent the latest available observations of the
	   object's current state. */
	Conditions []v1alpha1.Condition `json:"conditions,omitempty"`

	// ObservedGeneration is the generation of the resource that was most recently observed by the Config Connector controller. If this is equal to metadata.generation, then that means that the current reported status reflects the most recent desired state of the resource.
	ObservedGeneration *int64 `json:"observedGeneration,omitempty"`

	// A unique specifier for the ApigeeRegistryAPI resource in GCP.
	ExternalRef *string `json:"externalRef,omitempty"`

	// ObservedState is the state of the resource as most recently observed in GCP.
	ObservedState *ApigeeRegistryAPIObservedState `json:"observedState,omitempty"`
}

// ApigeeRegistryAPIObservedState is the state of the ApigeeRegistryAPI resource as most recently observed in GCP.
// +kcc:observedstate:proto=google.cloud.apigeeregistry.v1.Api
type ApigeeRegistryAPIObservedState struct {
	// Output only. Creation timestamp.
	// +kcc:proto:field=google.cloud.apigeeregistry.v1.Api.create_time
	CreateTime *string `json:"createTime,omitempty"`

	// Output only. Last update timestamp.
	// +kcc:proto:field=google.cloud.apigeeregistry.v1.Api.update_time
	UpdateTime *string `json:"updateTime,omitempty"`
}

// +genclient
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:resource:categories=gcp,shortName=gcpapigeeregistryapi;gcpapigeeregistryapis
// +kubebuilder:subresource:status
// +kubebuilder:metadata:labels="cnrm.cloud.google.com/managed-by-kcc=true"
// +kubebuilder:metadata:labels="cnrm.cloud.google.com/system=true"
// +kubebuilder:printcolumn:name="Age",JSONPath=".metadata.creationTimestamp",type="date"
// +kubebuilder:printcolumn:name="Ready",JSONPath=".status.conditions[?(@.type=='Ready')].status",type="string",description="When 'True', the most recent reconcile of the resource succeeded"
// +kubebuilder:printcolumn:name="Status",JSONPath=".status.conditions[?(@.type=='Ready')].reason",type="string",description="The reason for the value in 'Ready'"
// +kubebuilder:printcolumn:name="Status Age",JSONPath=".status.conditions[?(@.type=='Ready')].lastTransitionTime",type="date",description="The last transition time for the value in 'Status'"

// ApigeeRegistryAPI is the Schema for the ApigeeRegistryAPI API
// +k8s:openapi-gen=true
type ApigeeRegistryAPI struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	// +required
	Spec   ApigeeRegistryAPISpec   `json:"spec,omitempty"`
	Status ApigeeRegistryAPIStatus `json:"status,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// ApigeeRegistryAPIList contains a list of ApigeeRegistryAPI
type ApigeeRegistryAPIList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []ApigeeRegistryAPI `json:"items"`
}

func init() {
	SchemeBuilder.Register(&ApigeeRegistryAPI{}, &ApigeeRegistryAPIList{})
}
