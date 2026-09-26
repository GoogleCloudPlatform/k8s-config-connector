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

var GKEBackupRestoreChannelGVK = GroupVersion.WithKind("GKEBackupRestoreChannel")

// GKEBackupRestoreChannelSpec defines the desired state of GKEBackupRestoreChannel
// +kcc:spec:proto=google.cloud.gkebackup.v1.RestoreChannel
type GKEBackupRestoreChannelSpec struct {
	// The project that this resource belongs to.
	ProjectRef *refsv1beta1.ProjectRef `json:"projectRef"`

	// The location of this resource.
	Location *string `json:"location,omitempty"`

	// The GKEBackupRestoreChannel name. If not given, the metadata.name will be used.
	ResourceID *string `json:"resourceID,omitempty"`

	// Required. Immutable. The project into which the backups will be restored.
	// The format is `projects/{projectId}` or `projects/{projectNumber}`.
	// +required
	// +kcc:proto:field=google.cloud.gkebackup.v1.RestoreChannel.destination_project
	DestinationProjectRef *refsv1beta1.ProjectRef `json:"destinationProjectRef"`

	// Optional. A set of custom labels supplied by user.
	// +kcc:proto:field=google.cloud.gkebackup.v1.RestoreChannel.labels
	Labels map[string]string `json:"labels,omitempty"`

	// Optional. User specified descriptive string for this RestoreChannel.
	// +kcc:proto:field=google.cloud.gkebackup.v1.RestoreChannel.description
	Description *string `json:"description,omitempty"`
}

// GKEBackupRestoreChannelStatus defines the config connector machine state of GKEBackupRestoreChannel
type GKEBackupRestoreChannelStatus struct {
	/* Conditions represent the latest available observations of the
	   object's current state. */
	Conditions []v1alpha1.Condition `json:"conditions,omitempty"`

	// ObservedGeneration is the generation of the resource that was most recently observed by the Config Connector controller. If this is equal to metadata.generation, then that means that the current reported status reflects the most recent desired state of the resource.
	ObservedGeneration *int64 `json:"observedGeneration,omitempty"`

	// A unique specifier for the GKEBackupRestoreChannel resource in GCP.
	ExternalRef *string `json:"externalRef,omitempty"`

	// ObservedState is the state of the resource as most recently observed in GCP.
	ObservedState *GKEBackupRestoreChannelObservedState `json:"observedState,omitempty"`
}

// GKEBackupRestoreChannelObservedState is the state of the GKEBackupRestoreChannel resource as most recently observed in GCP.
// +kcc:observedstate:proto=google.cloud.gkebackup.v1.RestoreChannel
type GKEBackupRestoreChannelObservedState struct {
	// Output only. Server generated global unique identifier of [UUID](https://en.wikipedia.org/wiki/Universally_unique_identifier) format.
	// +kcc:proto:field=google.cloud.gkebackup.v1.RestoreChannel.uid
	UID *string `json:"uid,omitempty"`

	// Output only. The timestamp when this RestoreChannel resource was created.
	// +kcc:proto:field=google.cloud.gkebackup.v1.RestoreChannel.create_time
	CreateTime *string `json:"createTime,omitempty"`

	// Output only. The timestamp when this RestoreChannel resource was last updated.
	// +kcc:proto:field=google.cloud.gkebackup.v1.RestoreChannel.update_time
	UpdateTime *string `json:"updateTime,omitempty"`

	// Output only. `etag` is used for optimistic concurrency control as a way to help prevent simultaneous updates of a RestoreChannel from overwriting each other.
	// +kcc:proto:field=google.cloud.gkebackup.v1.RestoreChannel.etag
	Etag *string `json:"etag,omitempty"`

	// Output only. The project_id where backups will be restored.
	// +kcc:proto:field=google.cloud.gkebackup.v1.RestoreChannel.destination_project_id
	DestinationProjectID *string `json:"destinationProjectID,omitempty"`
}

// +genclient
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:resource:categories=gcp,shortName=gcpgkebackuprestorechannel;gcpgkebackuprestorechannels
// +kubebuilder:subresource:status
// +kubebuilder:metadata:labels="cnrm.cloud.google.com/managed-by-kcc=true"
// +kubebuilder:metadata:labels="cnrm.cloud.google.com/system=true"
// +kubebuilder:metadata:labels="cnrm.cloud.google.com/stability-level=alpha"
// +kubebuilder:printcolumn:name="Age",JSONPath=".metadata.creationTimestamp",type="date"
// +kubebuilder:printcolumn:name="Ready",JSONPath=".status.conditions[?(@.type=='Ready')].status",type="string",description="When 'True', the most recent reconcile of the resource succeeded"
// +kubebuilder:printcolumn:name="Status",JSONPath=".status.conditions[?(@.type=='Ready')].reason",type="string",description="The reason for the value in 'Ready'"
// +kubebuilder:printcolumn:name="Status Age",JSONPath=".status.conditions[?(@.type=='Ready')].lastTransitionTime",type="date",description="The last transition time for the value in 'Status'"

// GKEBackupRestoreChannel is the Schema for the GKEBackupRestoreChannel API
// +k8s:openapi-gen=true
type GKEBackupRestoreChannel struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	// +required
	Spec   GKEBackupRestoreChannelSpec   `json:"spec,omitempty"`
	Status GKEBackupRestoreChannelStatus `json:"status,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// GKEBackupRestoreChannelList contains a list of GKEBackupRestoreChannel
type GKEBackupRestoreChannelList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []GKEBackupRestoreChannel `json:"items"`
}

func init() {
	SchemeBuilder.Register(&GKEBackupRestoreChannel{}, &GKEBackupRestoreChannelList{})
}
