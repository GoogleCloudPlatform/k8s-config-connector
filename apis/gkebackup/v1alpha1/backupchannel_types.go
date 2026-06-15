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

var GKEBackupBackupChannelGVK = GroupVersion.WithKind("GKEBackupBackupChannel")

// GKEBackupBackupChannelSpec defines the desired state of GKEBackupBackupChannel
// +kcc:spec:proto=google.cloud.gkebackup.v1.BackupChannel
type GKEBackupBackupChannelSpec struct {
	// The project that this resource belongs to.
	ProjectRef *refsv1beta1.ProjectRef `json:"projectRef"`

	// The location of this resource.
	Location string `json:"location"`

	// The GKEBackupBackupChannel name. If not given, the metadata.name will be used.
	ResourceID *string `json:"resourceID,omitempty"`

	// Required. Immutable. The project where Backups are allowed to be stored.
	// +required
	// +kcc:proto:field=google.cloud.gkebackup.v1.BackupChannel.destination_project
	DestinationProjectRef *refsv1beta1.ProjectRef `json:"destinationProjectRef"`

	// Optional. A set of custom labels supplied by user.
	// +optional
	// +kcc:proto:field=google.cloud.gkebackup.v1.BackupChannel.labels
	Labels map[string]string `json:"labels,omitempty"`

	// Optional. User specified descriptive string for this BackupChannel.
	// +optional
	// +kcc:proto:field=google.cloud.gkebackup.v1.BackupChannel.description
	Description *string `json:"description,omitempty"`
}

// GKEBackupBackupChannelStatus defines the config connector machine state of GKEBackupBackupChannel
type GKEBackupBackupChannelStatus struct {
	/* Conditions represent the latest available observations of the
	   object's current state. */
	Conditions []v1alpha1.Condition `json:"conditions,omitempty"`

	// ObservedGeneration is the generation of the resource that was most recently observed by the Config Connector controller. If this is equal to metadata.generation, then that means that the current reported status reflects the most recent desired state of the resource.
	ObservedGeneration *int64 `json:"observedGeneration,omitempty"`

	// A unique specifier for the GKEBackupBackupChannel resource in GCP.
	ExternalRef *string `json:"externalRef,omitempty"`

	// ObservedState is the state of the resource as most recently observed in GCP.
	ObservedState *GKEBackupBackupChannelObservedState `json:"observedState,omitempty"`
}

// GKEBackupBackupChannelObservedState is the state of the GKEBackupBackupChannel resource as most recently observed in GCP.
// +kcc:observedstate:proto=google.cloud.gkebackup.v1.BackupChannel
type GKEBackupBackupChannelObservedState struct {
	// Output only. Server generated global unique identifier of
	//  [UUID](https://en.wikipedia.org/wiki/Universally_unique_identifier) format.
	// +kcc:proto:field=google.cloud.gkebackup.v1.BackupChannel.uid
	UID *string `json:"uid,omitempty"`

	// Output only. The timestamp when this BackupChannel resource was created.
	// +kcc:proto:field=google.cloud.gkebackup.v1.BackupChannel.create_time
	CreateTime *string `json:"createTime,omitempty"`

	// Output only. The timestamp when this BackupChannel resource was last
	//  updated.
	// +kcc:proto:field=google.cloud.gkebackup.v1.BackupChannel.update_time
	UpdateTime *string `json:"updateTime,omitempty"`

	// Output only. `etag` is used for optimistic concurrency control as a way to
	//  help prevent simultaneous updates of a BackupChannel from overwriting each
	//  other. It is strongly suggested that systems make use of the 'etag' in the
	//  read-modify-write cycle to perform BackupChannel updates in order to
	//  avoid race conditions: An `etag` is returned in the response to
	//  `GetBackupChannel`, and systems are expected to put that etag in the
	//  request to `UpdateBackupChannel` or `DeleteBackupChannel` to
	//  ensure that their change will be applied to the same version of the
	//  resource.
	// +kcc:proto:field=google.cloud.gkebackup.v1.BackupChannel.etag
	Etag *string `json:"etag,omitempty"`

	// Output only. The project_id where Backups are allowed to be stored.
	//  Example Project ID: "my-project-id".
	//  This will be an OUTPUT_ONLY field to return the project_id of the
	//  destination project.
	// +kcc:proto:field=google.cloud.gkebackup.v1.BackupChannel.destination_project_id
	DestinationProjectID *string `json:"destinationProjectID,omitempty"`
}

// +genclient
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:resource:categories=gcp,shortName=gcpgkebackupbackupchannel;gcpgkebackupbackupchannels
// +kubebuilder:subresource:status
// +kubebuilder:metadata:labels="cnrm.cloud.google.com/managed-by-kcc=true"
// +kubebuilder:metadata:labels="cnrm.cloud.google.com/system=true"
// +kubebuilder:metadata:labels="cnrm.cloud.google.com/stability-level=alpha"
// +kubebuilder:printcolumn:name="Age",JSONPath=".metadata.creationTimestamp",type="date"
// +kubebuilder:printcolumn:name="Ready",JSONPath=".status.conditions[?(@.type=='Ready')].status",type="string",description="When 'True', the most recent reconcile of the resource succeeded"
// +kubebuilder:printcolumn:name="Status",JSONPath=".status.conditions[?(@.type=='Ready')].reason",type="string",description="The reason for the value in 'Ready'"
// +kubebuilder:printcolumn:name="Status Age",JSONPath=".status.conditions[?(@.type=='Ready')].lastTransitionTime",type="date",description="The last transition time for the value in 'Status'"

// GKEBackupBackupChannel is the Schema for the GKEBackupBackupChannel API
// +k8s:openapi-gen=true
type GKEBackupBackupChannel struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	// +required
	Spec   GKEBackupBackupChannelSpec   `json:"spec,omitempty"`
	Status GKEBackupBackupChannelStatus `json:"status,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// GKEBackupBackupChannelList contains a list of GKEBackupBackupChannel
type GKEBackupBackupChannelList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []GKEBackupBackupChannel `json:"items"`
}

func init() {
	SchemeBuilder.Register(&GKEBackupBackupChannel{}, &GKEBackupBackupChannelList{})
}
