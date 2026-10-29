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

// API sources for VMMigrationSource, recorded by generate-types:
// +kcc:source:proto=https://github.com/googleapis/googleapis/blob/1765b559c42386788ff0c6412491277b4791107a/google/cloud/vmmigration/v1/vmmigration.proto
// +kcc:source:service-docs=https://cloud.google.com/migrate/virtual-machines
// +kcc:source:resource-docs=https://cloud.google.com/migrate/virtual-machines/docs/5.0/reference/rest/v1/projects.locations.sources

package v1alpha1

import (
	"github.com/GoogleCloudPlatform/k8s-config-connector/apis/common/parent"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/apis/k8s/v1alpha1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

var VMMigrationSourceGVK = GroupVersion.WithKind("VMMigrationSource")

// VMMigrationSourceSpec defines the desired state of VMMigrationSource
// +kcc:spec:proto=google.cloud.vmmigration.v1.Source
// +kcc:required-from-proto
type VMMigrationSourceSpec struct {
	*parent.ProjectAndLocationRef `json:",inline"`

	// The VMMigrationSource name. If not given, the metadata.name will be used.
	ResourceID *string `json:"resourceID,omitempty"`
	// Vmware type source details.
	// +kcc:proto:field=google.cloud.vmmigration.v1.Source.vmware
	Vmware *VmwareSourceDetails `json:"vmware,omitempty"`

	// AWS type source details.
	// +kcc:proto:field=google.cloud.vmmigration.v1.Source.aws
	Aws *AwsSourceDetails `json:"aws,omitempty"`

	// User-provided description of the source.
	// +kcc:proto:field=google.cloud.vmmigration.v1.Source.description
	Description *string `json:"description,omitempty"`
}

// VMMigrationSourceStatus defines the config connector machine state of VMMigrationSource
type VMMigrationSourceStatus struct {
	/* Conditions represent the latest available observations of the
	   object's current state. */
	Conditions []v1alpha1.Condition `json:"conditions,omitempty"`

	// ObservedGeneration is the generation of the resource that was most recently observed by the Config Connector controller. If this is equal to metadata.generation, then that means that the current reported status reflects the most recent desired state of the resource.
	ObservedGeneration *int64 `json:"observedGeneration,omitempty"`

	// A unique specifier for the VMMigrationSource resource in GCP.
	ExternalRef *string `json:"externalRef,omitempty"`

	// ObservedState is the state of the resource as most recently observed in GCP.
	ObservedState *VMMigrationSourceObservedState `json:"observedState,omitempty"`
}

// VMMigrationSourceObservedState is the state of the VMMigrationSource resource as most recently observed in GCP.
// +kcc:observedstate:proto=google.cloud.vmmigration.v1.Source
type VMMigrationSourceObservedState struct {
	// AWS type source details.
	// +kcc:proto:field=google.cloud.vmmigration.v1.Source.aws
	Aws *AwsSourceDetailsObservedState `json:"aws,omitempty"`

	// Output only. The create time timestamp.
	// +kcc:proto:field=google.cloud.vmmigration.v1.Source.create_time
	CreateTime *string `json:"createTime,omitempty"`

	// Output only. The update time timestamp.
	// +kcc:proto:field=google.cloud.vmmigration.v1.Source.update_time
	UpdateTime *string `json:"updateTime,omitempty"`
}

// +genclient
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:resource:categories=gcp,shortName=gcpvmmigrationsource;gcpvmmigrationsources
// +kubebuilder:subresource:status
// +kubebuilder:metadata:labels="cnrm.cloud.google.com/managed-by-kcc=true"
// +kubebuilder:metadata:labels="cnrm.cloud.google.com/system=true"
// +kubebuilder:printcolumn:name="Age",JSONPath=".metadata.creationTimestamp",type="date"
// +kubebuilder:printcolumn:name="Ready",JSONPath=".status.conditions[?(@.type=='Ready')].status",type="string",description="When 'True', the most recent reconcile of the resource succeeded"
// +kubebuilder:printcolumn:name="Status",JSONPath=".status.conditions[?(@.type=='Ready')].reason",type="string",description="The reason for the value in 'Ready'"
// +kubebuilder:printcolumn:name="Status Age",JSONPath=".status.conditions[?(@.type=='Ready')].lastTransitionTime",type="date",description="The last transition time for the value in 'Status'"

// VMMigrationSource is the Schema for the VMMigrationSource API
// +k8s:openapi-gen=true
type VMMigrationSource struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	// +required
	Spec   VMMigrationSourceSpec   `json:"spec,omitempty"`
	Status VMMigrationSourceStatus `json:"status,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// VMMigrationSourceList contains a list of VMMigrationSource
type VMMigrationSourceList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []VMMigrationSource `json:"items"`
}

func init() {
	SchemeBuilder.Register(&VMMigrationSource{}, &VMMigrationSourceList{})
}
