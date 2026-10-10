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

var MigrationCenterSourceGVK = GroupVersion.WithKind("MigrationCenterSource")

// MigrationCenterSourceSpec defines the desired state of MigrationCenterSource
// +kcc:spec:proto=google.cloud.migrationcenter.v1.Source
type MigrationCenterSourceSpec struct {
	// The project that this resource belongs to.
	// +required
	ProjectRef *refsv1beta1.ProjectRef `json:"projectRef"`

	// +kubebuilder:validation:XValidation:rule="self == oldSelf",message="Location field is immutable"
	// Immutable. The location of this resource.
	// +required
	Location *string `json:"location"`

	// The MigrationCenterSource name. If not given, the metadata.name will be used.
	// +optional
	ResourceID *string `json:"resourceID,omitempty"`

	// User-friendly display name.
	// +optional
	// +kcc:proto:field=google.cloud.migrationcenter.v1.Source.display_name
	DisplayName *string `json:"displayName,omitempty"`

	// Free-text description.
	// +optional
	// +kcc:proto:field=google.cloud.migrationcenter.v1.Source.description
	Description *string `json:"description,omitempty"`

	// Data source type.
	// +kubebuilder:validation:Enum=SOURCE_TYPE_UNKNOWN;SOURCE_TYPE_UPLOAD;SOURCE_TYPE_GUEST_OS_SCAN;SOURCE_TYPE_INVENTORY_SCAN;SOURCE_TYPE_CUSTOM
	// +optional
	// +kcc:proto:field=google.cloud.migrationcenter.v1.Source.type
	Type *string `json:"type,omitempty"`

	// The information confidence of the source.
	// The higher the value, the higher the confidence.
	// +optional
	// +kcc:proto:field=google.cloud.migrationcenter.v1.Source.priority
	Priority *int32 `json:"priority,omitempty"`

	// If `true`, the source is managed by other service(s).
	// +optional
	// +kcc:proto:field=google.cloud.migrationcenter.v1.Source.managed
	Managed *bool `json:"managed,omitempty"`
}

// MigrationCenterSourceStatus defines the config connector machine state of MigrationCenterSource
type MigrationCenterSourceStatus struct {
	/* Conditions represent the latest available observations of the
	   object's current state. */
	Conditions []v1alpha1.Condition `json:"conditions,omitempty"`

	// ObservedGeneration is the generation of the resource that was most recently observed by the Config Connector controller. If this is equal to metadata.generation, then that means that the current reported status reflects the most recent desired state of the resource.
	ObservedGeneration *int64 `json:"observedGeneration,omitempty"`

	// A unique specifier for the MigrationCenterSource resource in GCP.
	ExternalRef *string `json:"externalRef,omitempty"`

	// ObservedState is the state of the resource as most recently observed in GCP.
	ObservedState *MigrationCenterSourceObservedState `json:"observedState,omitempty"`
}

// MigrationCenterSourceObservedState is the state of the MigrationCenterSource resource as most recently observed in GCP.
// +kcc:observedstate:proto=google.cloud.migrationcenter.v1.Source
type MigrationCenterSourceObservedState struct {
	// Output only. The timestamp when the source was created.
	// +optional
	// +kcc:proto:field=google.cloud.migrationcenter.v1.Source.create_time
	CreateTime *string `json:"createTime,omitempty"`

	// Output only. The timestamp when the source was last updated.
	// +optional
	// +kcc:proto:field=google.cloud.migrationcenter.v1.Source.update_time
	UpdateTime *string `json:"updateTime,omitempty"`

	// Output only. Number of frames that are still being processed.
	// +optional
	// +kcc:proto:field=google.cloud.migrationcenter.v1.Source.pending_frame_count
	PendingFrameCount *int32 `json:"pendingFrameCount,omitempty"`

	// Output only. The number of frames that were reported by the source and contained errors.
	// +optional
	// +kcc:proto:field=google.cloud.migrationcenter.v1.Source.error_frame_count
	ErrorFrameCount *int32 `json:"errorFrameCount,omitempty"`

	// Output only. The state of the source.
	// +kubebuilder:validation:Enum=STATE_UNSPECIFIED;ACTIVE;DELETING;INVALID
	// +optional
	// +kcc:proto:field=google.cloud.migrationcenter.v1.Source.state
	State *string `json:"state,omitempty"`
}

// +genclient
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:resource:categories=gcp,shortName=gcpmigrationcentersource;gcpmigrationcentersources
// +kubebuilder:subresource:status
// +kubebuilder:metadata:labels="cnrm.cloud.google.com/managed-by-kcc=true"
// +kubebuilder:metadata:labels="cnrm.cloud.google.com/system=true"
// +kubebuilder:metadata:labels="cnrm.cloud.google.com/stability-level=alpha"
// +kubebuilder:printcolumn:name="Age",JSONPath=".metadata.creationTimestamp",type="date"
// +kubebuilder:printcolumn:name="Ready",JSONPath=".status.conditions[?(@.type=='Ready')].status",type="string",description="When 'True', the most recent reconcile of the resource succeeded"
// +kubebuilder:printcolumn:name="Status",JSONPath=".status.conditions[?(@.type=='Ready')].reason",type="string",description="The reason for the value in 'Ready'"
// +kubebuilder:printcolumn:name="Status Age",JSONPath=".status.conditions[?(@.type=='Ready')].lastTransitionTime",type="date",description="The last transition time for the value in 'Status'"

// MigrationCenterSource is the Schema for the MigrationCenterSource API
// +k8s:openapi-gen=true
type MigrationCenterSource struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	// +required
	Spec   MigrationCenterSourceSpec   `json:"spec,omitempty"`
	Status MigrationCenterSourceStatus `json:"status,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// MigrationCenterSourceList contains a list of MigrationCenterSource
type MigrationCenterSourceList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []MigrationCenterSource `json:"items"`
}

func init() {
	SchemeBuilder.Register(&MigrationCenterSource{}, &MigrationCenterSourceList{})
}
