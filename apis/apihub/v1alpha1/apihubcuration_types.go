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

var APIHubCurationGVK = GroupVersion.WithKind("APIHubCuration")

// APIHubCurationSpec defines the desired state of APIHubCuration
// +kcc:spec:proto=google.cloud.apihub.v1.Curation
type APIHubCurationSpec struct {
	// The project that this resource belongs to.
	ProjectRef *refsv1beta1.ProjectRef `json:"projectRef"`

	// The location of this resource.
	Location string `json:"location"`

	// The APIHubCuration name. If not given, the metadata.name will be used.
	ResourceID *string `json:"resourceID,omitempty"`

	// Required. The display name of the curation.
	// +kubebuilder:validation:Required
	DisplayName *string `json:"displayName,omitempty"`

	// Optional. The description of the curation.
	// +kubebuilder:validation:Optional
	Description *string `json:"description,omitempty"`

	// Required. The endpoint to be triggered for curation.
	// +kubebuilder:validation:Required
	Endpoint *Endpoint `json:"endpoint,omitempty"`
}

// APIHubCurationStatus defines the config connector machine state of APIHubCuration
type APIHubCurationStatus struct {
	/* Conditions represent the latest available observations of the
	   object's current state. */
	Conditions []v1alpha1.Condition `json:"conditions,omitempty"`

	// ObservedGeneration is the generation of the resource that was most recently observed by the Config Connector controller. If this is equal to metadata.generation, then that means that the current reported status reflects the most recent desired state of the resource.
	ObservedGeneration *int64 `json:"observedGeneration,omitempty"`

	// A unique specifier for the APIHubCuration resource in GCP.
	ExternalRef *string `json:"externalRef,omitempty"`

	// ObservedState is the state of the resource as most recently observed in GCP.
	ObservedState *APIHubCurationObservedState `json:"observedState,omitempty"`
}

// APIHubCurationObservedState is the state of the APIHubCuration resource as most recently observed in GCP.
// +kcc:observedstate:proto=google.cloud.apihub.v1.Curation
type APIHubCurationObservedState struct {
	// Output only. The plugin instances and associated actions that are using the curation.
	// +kubebuilder:validation:Optional
	PluginInstanceActions []Curation_PluginInstanceActionIDObservedState `json:"pluginInstanceActions,omitempty"`

	// Output only. The last execution state of the curation.
	// +kubebuilder:validation:Optional
	LastExecutionState *string `json:"lastExecutionState,omitempty"`

	// Output only. The error code of the last execution of the curation.
	// +kubebuilder:validation:Optional
	LastExecutionErrorCode *string `json:"lastExecutionErrorCode,omitempty"`

	// Output only. Error message describing the failure, if any, during the last execution of the curation.
	// +kubebuilder:validation:Optional
	LastExecutionErrorMessage *string `json:"lastExecutionErrorMessage,omitempty"`

	// Output only. The time at which the curation was created.
	// +kubebuilder:validation:Optional
	CreateTime *string `json:"createTime,omitempty"`

	// Output only. The time at which the curation was last updated.
	// +kubebuilder:validation:Optional
	UpdateTime *string `json:"updateTime,omitempty"`
}

// +genclient
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:resource:categories=gcp,shortName=gcpapihubcuration;gcpapihubcurations
// +kubebuilder:subresource:status
// +kubebuilder:metadata:labels="cnrm.cloud.google.com/managed-by-kcc=true"
// +kubebuilder:metadata:labels="cnrm.cloud.google.com/system=true"
// +kubebuilder:metadata:labels="cnrm.cloud.google.com/stability-level=alpha"
// +kubebuilder:printcolumn:name="Age",JSONPath=".metadata.creationTimestamp",type="date"
// +kubebuilder:printcolumn:name="Ready",JSONPath=".status.conditions[?(@.type=='Ready')].status",type="string",description="When 'True', the most recent reconcile of the resource succeeded"
// +kubebuilder:printcolumn:name="Status",JSONPath=".status.conditions[?(@.type=='Ready')].reason",type="string",description="The reason for the value in 'Ready'"
// +kubebuilder:printcolumn:name="Status Age",JSONPath=".status.conditions[?(@.type=='Ready')].lastTransitionTime",type="date",description="The last transition time for the value in 'Status'"

// APIHubCuration is the Schema for the APIHubCuration API
// +k8s:openapi-gen=true
type APIHubCuration struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	// +required
	Spec   APIHubCurationSpec   `json:"spec,omitempty"`
	Status APIHubCurationStatus `json:"status,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// APIHubCurationList contains a list of APIHubCuration
type APIHubCurationList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []APIHubCuration `json:"items"`
}

func init() {
	SchemeBuilder.Register(&APIHubCuration{}, &APIHubCurationList{})
}
