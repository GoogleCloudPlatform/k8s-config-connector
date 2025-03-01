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

var ComposerEnvironmentGVK = GroupVersion.WithKind("ComposerEnvironment")

// ComposerEnvironmentSpec defines the desired state of ComposerEnvironment
// +kcc:proto=google.cloud.orchestration.airflow.service.v1.Environment
type ComposerEnvironmentSpec struct {
	// The ComposerEnvironment name. If not given, the metadata.name will be used.
	ResourceID *string `json:"resourceID,omitempty"`
}

// ComposerEnvironmentStatus defines the config connector machine state of ComposerEnvironment
type ComposerEnvironmentStatus struct {
	/* Conditions represent the latest available observations of the
	   object's current state. */
	Conditions []v1alpha1.Condition `json:"conditions,omitempty"`

	// ObservedGeneration is the generation of the resource that was most recently observed by the Config Connector controller. If this is equal to metadata.generation, then that means that the current reported status reflects the most recent desired state of the resource.
	ObservedGeneration *int64 `json:"observedGeneration,omitempty"`

	// A unique specifier for the ComposerEnvironment resource in GCP.
	ExternalRef *string `json:"externalRef,omitempty"`

	// ObservedState is the state of the resource as most recently observed in GCP.
	ObservedState *ComposerEnvironmentObservedState `json:"observedState,omitempty"`
}

// ComposerEnvironmentObservedState is the state of the ComposerEnvironment resource as most recently observed in GCP.
// +kcc:proto=google.cloud.orchestration.airflow.service.v1.Environment
type ComposerEnvironmentObservedState struct {
}

// +genclient
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// TODO(user): make sure the pluralizaiton below is correct
// +kubebuilder:resource:categories=gcp,shortName=gcpcomposerenvironment;gcpcomposerenvironments
// +kubebuilder:subresource:status
// +kubebuilder:metadata:labels="cnrm.cloud.google.com/managed-by-kcc=true";"cnrm.cloud.google.com/system=true"
// +kubebuilder:printcolumn:name="Age",JSONPath=".metadata.creationTimestamp",type="date"
// +kubebuilder:printcolumn:name="Ready",JSONPath=".status.conditions[?(@.type=='Ready')].status",type="string",description="When 'True', the most recent reconcile of the resource succeeded"
// +kubebuilder:printcolumn:name="Status",JSONPath=".status.conditions[?(@.type=='Ready')].reason",type="string",description="The reason for the value in 'Ready'"
// +kubebuilder:printcolumn:name="Status Age",JSONPath=".status.conditions[?(@.type=='Ready')].lastTransitionTime",type="date",description="The last transition time for the value in 'Status'"

// ComposerEnvironment is the Schema for the ComposerEnvironment API
// +k8s:openapi-gen=true
type ComposerEnvironment struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	// +required
	Spec   ComposerEnvironmentSpec   `json:"spec,omitempty"`
	Status ComposerEnvironmentStatus `json:"status,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// ComposerEnvironmentList contains a list of ComposerEnvironment
type ComposerEnvironmentList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []ComposerEnvironment `json:"items"`
}

func init() {
	SchemeBuilder.Register(&ComposerEnvironment{}, &ComposerEnvironmentList{})
}
 type ComposerEnvironmentSpec struct {
	// The ComposerEnvironment name. If not given, the metadata.name will be used.
	ResourceID *string `json:"resourceID,omitempty"`

	// Optional. Configuration parameters for this environment.
	// +kcc:proto:field=google.cloud.orchestration.airflow.service.v1.Environment.config
	Config *EnvironmentConfig `json:"config,omitempty"`

	// Optional. User-defined labels for this environment.
	//  The labels map can contain no more than 64 entries. Entries of the labels
	//  map are UTF8 strings that comply with the following restrictions:
	//
	//  * Keys must conform to regexp: [\p{Ll}\p{Lo}][\p{Ll}\p{Lo}\p{N}_-]{0,62}
	//  * Values must conform to regexp:  [\p{Ll}\p{Lo}\p{N}_-]{0,63}
	//  * Both keys and values are additionally constrained to be <= 128 bytes in
	//  size.
	// +kcc:proto:field=google.cloud.orchestration.airflow.service.v1.Environment.labels
	Labels map[string]string `json:"labels,omitempty"`

	/* Immutable. The Project that this resource belongs to. */
	ProjectRef *refs.ProjectRef `json:"projectRef"`

	// Immutable. The name of the region where the Environment will be created.
	// +kubebuilder:validation:XValidation:rule="self == oldSelf",message="Region is immutable."
	// Required.
	Region string `json:"region"`

	// Optional. Storage configuration for this environment.
	// +kcc:proto:field=google.cloud.orchestration.airflow.service.v1.Environment.storage_config
	StorageConfig *StorageConfig `json:"storageConfig,omitempty"`
}

 type ComposerEnvironmentSpec struct {
	// The ComposerEnvironment name. If not given, the metadata.name will be used.
	ResourceID *string `json:"resourceID,omitempty"`

	// Immutable. The Project that this resource belongs to.
	// +required
	ProjectRef refs.ProjectRef `json:"projectRef"`

	// Immutable. The location for the resource.
	// +required
	Location *string `json:"location,omitempty"`

	// Optional. Configuration parameters for this environment.
	// +kcc:proto:field=google.cloud.orchestration.airflow.service.v1.Environment.config
	Config *EnvironmentConfig `json:"config,omitempty"`

	// Optional. User-defined labels for this environment.
	//  The labels map can contain no more than 64 entries. Entries of the labels
	//  map are UTF8 strings that comply with the following restrictions:
	//
	//  * Keys must conform to regexp: [\p{Ll}\p{Lo}][\p{Ll}\p{Lo}\p{N}_-]{0,62}
	//  * Values must conform to regexp:  [\p{Ll}\p{Lo}\p{N}_-]{0,63}
	//  * Both keys and values are additionally constrained to be <= 128 bytes in
	//  size.
	// +kcc:proto:field=google.cloud.orchestration.airflow.service.v1.Environment.labels
	Labels map[string]string `json:"labels,omitempty"`

	// Optional. Storage configuration for this environment.
	// +kcc:proto:field=google.cloud.orchestration.airflow.service.v1.Environment.storage_config
	StorageConfig *StorageConfig `json:"storageConfig,omitempty"`
}



