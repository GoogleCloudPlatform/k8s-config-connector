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
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/apis/k8s/v1alpha1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

var DataplexEnvironmentGVK = GroupVersion.WithKind("DataplexEnvironment")

// DataplexEnvironmentSpec defines the desired state of DataplexEnvironment
// +kcc:proto=google.cloud.dataplex.v1.Environment
type DataplexEnvironmentSpec struct {
	// The DataplexEnvironment name. If not given, the metadata.name will be used.
	ResourceID *string `json:"resourceID,omitempty"`

	// Optional. User friendly display name.
	// +kcc:proto:field=google.cloud.dataplex.v1.Environment.display_name
	DisplayName *string `json:"displayName,omitempty"`

	// Optional. User defined labels for the environment.
	// +kcc:proto:field=google.cloud.dataplex.v1.Environment.labels
	Labels map[string]string `json:"labels,omitempty"`

	// Optional. Description of the environment.
	// +kcc:proto:field=google.cloud.dataplex.v1.Environment.description
	Description *string `json:"description,omitempty"`

	// Required. Infrastructure specification for the Environment.
	// +kcc:proto:field=google.cloud.dataplex.v1.Environment.infrastructure_spec
	InfrastructureSpec *Environment_InfrastructureSpec `json:"infrastructureSpec,omitempty"`

	// Optional. Configuration for sessions created for this environment.
	// +kcc:proto:field=google.cloud.dataplex.v1.Environment.session_spec
	SessionSpec *Environment_SessionSpec `json:"sessionSpec,omitempty"`
}

// DataplexEnvironmentStatus defines the config connector machine state of DataplexEnvironment
type DataplexEnvironmentStatus struct {
	/* Conditions represent the latest available observations of the
	   object's current state. */
	Conditions []v1alpha1.Condition `json:"conditions,omitempty"`

	// ObservedGeneration is the generation of the resource that was most recently observed by the Config Connector controller. If this is equal to metadata.generation, then that means that the current reported status reflects the most recent desired state of the resource.
	ObservedGeneration *int64 `json:"observedGeneration,omitempty"`

	// A unique specifier for the DataplexEnvironment resource in GCP.
	ExternalRef *string `json:"externalRef,omitempty"`

	// ObservedState is the state of the resource as most recently observed in GCP.
	ObservedState *DataplexEnvironmentObservedState `json:"observedState,omitempty"`
}

// DataplexEnvironmentObservedState is the state of the DataplexEnvironment resource as most recently observed in GCP.
// +kcc:proto=google.cloud.dataplex.v1.Environment
type DataplexEnvironmentObservedState struct {
	// Output only. System generated globally unique ID for the environment. This
	//  ID will be different if the environment is deleted and re-created with the
	//  same name.
	// +kcc:proto:field=google.cloud.dataplex.v1.Environment.uid
	Uid *string `json:"uid,omitempty"`

	// Output only. Environment creation time.
	// +kcc:proto:field=google.cloud.dataplex.v1.Environment.create_time
	CreateTime *string `json:"createTime,omitempty"`

	// Output only. The time when the environment was last updated.
	// +kcc:proto:field=google.cloud.dataplex.v1.Environment.update_time
	UpdateTime *string `json:"updateTime,omitempty"`

	// Output only. Current state of the environment.
	// +kcc:proto:field=google.cloud.dataplex.v1.Environment.state
	State *string `json:"state,omitempty"`

	// Output only. Status of sessions created for this environment.
	// +kcc:proto:field=google.cloud.dataplex.v1.Environment.session_status
	SessionStatus *Environment_SessionStatus `json:"sessionStatus,omitempty"`

	// Output only. URI Endpoints to access sessions associated with the
	//  Environment.
	// +kcc:proto:field=google.cloud.dataplex.v1.Environment.endpoints
	Endpoints *Environment_Endpoints `json:"endpoints,omitempty"`
}

// +genclient
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// TODO(user): make sure the pluralizaiton below is correct
// +kubebuilder:resource:categories=gcp,shortName=gcpdataplexenvironment;gcpdataplexenvironments
// +kubebuilder:subresource:status
// +kubebuilder:metadata:labels="cnrm.cloud.google.com/managed-by-kcc=true";"cnrm.cloud.google.com/system=true"
// +kubebuilder:printcolumn:name="Age",JSONPath=".metadata.creationTimestamp",type="date"
// +kubebuilder:printcolumn:name="Ready",JSONPath=".status.conditions[?(@.type=='Ready')].status",type="string",description="When 'True', the most recent reconcile of the resource succeeded"
// +kubebuilder:printcolumn:name="Status",JSONPath=".status.conditions[?(@.type=='Ready')].reason",type="string",description="The reason for the value in 'Ready'"
// +kubebuilder:printcolumn:name="Status Age",JSONPath=".status.conditions[?(@.type=='Ready')].lastTransitionTime",type="date",description="The last transition time for the value in 'Status'"

// DataplexEnvironment is the Schema for the DataplexEnvironment API
// +k8s:openapi-gen=true
type DataplexEnvironment struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	// +required
	Spec   DataplexEnvironmentSpec   `json:"spec,omitempty"`
	Status DataplexEnvironmentStatus `json:"status,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// DataplexEnvironmentList contains a list of DataplexEnvironment
type DataplexEnvironmentList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []DataplexEnvironment `json:"items"`
}

func init() {
	SchemeBuilder.Register(&DataplexEnvironment{}, &DataplexEnvironmentList{})
}
