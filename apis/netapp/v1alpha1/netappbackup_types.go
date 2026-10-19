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

// API sources for NetappBackup, recorded by generate-types:
// +kcc:source:proto=https://github.com/googleapis/googleapis/blob/e09e85d32ca349e1b205514a412817a7692595c6/google/cloud/netapp/v1/backup.proto
// +kcc:source:service-docs=https://cloud.google.com/netapp/
// +kcc:source:resource-docs=https://docs.cloud.google.com/netapp/volumes/docs/reference/rest/v1/projects.locations.backupVaults.backups

package v1alpha1

import (
	refsv1beta1 "github.com/GoogleCloudPlatform/k8s-config-connector/apis/refs/v1beta1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/apis/k8s/v1alpha1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

var NetappBackupGVK = GroupVersion.WithKind("NetappBackup")

// NetappBackupSpec defines the desired state of NetappBackup
// +kcc:spec:proto=google.cloud.netapp.v1.Backup
// +kcc:required-from-proto
type NetappBackupSpec struct {
	// The project that this resource belongs to.
	ProjectRef *refsv1beta1.ProjectRef `json:"projectRef"`

	// The location of this resource.
	Location *string `json:"location"`

	// A reference to the projects/{project}/locations/{location}/backupVaults/{backup_vault} this resource belongs to.
	// +kcc:guess
	// BackupVaultRef *BackupVaultRef `json:"backupVaultRef,omitempty"`

	// The NetappBackup name. If not given, the metadata.name will be used.
	ResourceID *string `json:"resourceID,omitempty"`
	// A description of the backup with 2048 characters or less.
	//  Requests with longer descriptions will be rejected.
	// +kcc:proto:field=google.cloud.netapp.v1.Backup.description
	Description *string `json:"description,omitempty"`

	// Volume full name of this backup belongs to.
	//  Either source_volume or ontap_source should be provided.
	//  Format:
	//  `projects/{projects_id}/locations/{location}/volumes/{volume_id}`
	// +kcc:proto:field=google.cloud.netapp.v1.Backup.source_volume
	SourceVolume *string `json:"sourceVolume,omitempty"`

	// If specified, backup will be created from the given snapshot.
	//  If not specified, there will be a new snapshot taken to initiate the backup
	//  creation. Format:
	//  `projects/{project_id}/locations/{location}/volumes/{volume_id}/snapshots/{snapshot_id}`
	// +kcc:proto:field=google.cloud.netapp.v1.Backup.source_snapshot
	SourceSnapshot *string `json:"sourceSnapshot,omitempty"`

	// Resource labels to represent user provided metadata.
	// +kcc:proto:field=google.cloud.netapp.v1.Backup.labels
	Labels map[string]string `json:"labels,omitempty"`
}

// NetappBackupStatus defines the config connector machine state of NetappBackup
type NetappBackupStatus struct {
	/* Conditions represent the latest available observations of the
	   object's current state. */
	Conditions []v1alpha1.Condition `json:"conditions,omitempty"`

	// ObservedGeneration is the generation of the resource that was most recently observed by the Config Connector controller. If this is equal to metadata.generation, then that means that the current reported status reflects the most recent desired state of the resource.
	ObservedGeneration *int64 `json:"observedGeneration,omitempty"`

	// A unique specifier for the NetappBackup resource in GCP.
	ExternalRef *string `json:"externalRef,omitempty"`

	// ObservedState is the state of the resource as most recently observed in GCP.
	ObservedState *NetappBackupObservedState `json:"observedState,omitempty"`
}

// NetappBackupObservedState is the state of the NetappBackup resource as most recently observed in GCP.
// +kcc:observedstate:proto=google.cloud.netapp.v1.Backup
type NetappBackupObservedState struct {
	// Output only. The backup state.
	// +kcc:proto:field=google.cloud.netapp.v1.Backup.state
	State *string `json:"state,omitempty"`

	// Output only. Size of the file system when the backup was created. When
	//  creating a new volume from the backup, the volume capacity will have to be
	//  at least as big.
	// +kcc:proto:field=google.cloud.netapp.v1.Backup.volume_usage_bytes
	VolumeUsageBytes *int64 `json:"volumeUsageBytes,omitempty"`

	// Output only. Type of backup, manually created or created by a backup
	//  policy.
	// +kcc:proto:field=google.cloud.netapp.v1.Backup.backup_type
	BackupType *string `json:"backupType,omitempty"`

	// Output only. The time when the backup was created.
	// +kcc:proto:field=google.cloud.netapp.v1.Backup.create_time
	CreateTime *string `json:"createTime,omitempty"`

	// Output only. Total size of all backups in a chain in bytes = baseline
	//  backup size + sum(incremental backup size)
	// +kcc:proto:field=google.cloud.netapp.v1.Backup.chain_storage_bytes
	ChainStorageBytes *int64 `json:"chainStorageBytes,omitempty"`

	// Output only. Reserved for future use
	// +kcc:proto:field=google.cloud.netapp.v1.Backup.satisfies_pzs
	SatisfiesPzs *bool `json:"satisfiesPzs,omitempty"`

	// Output only. Reserved for future use
	// +kcc:proto:field=google.cloud.netapp.v1.Backup.satisfies_pzi
	SatisfiesPzi *bool `json:"satisfiesPzi,omitempty"`

	// Output only. Region of the volume from which the backup was created.
	//  Format: `projects/{project_id}/locations/{location}`
	// +kcc:proto:field=google.cloud.netapp.v1.Backup.volume_region
	VolumeRegion *string `json:"volumeRegion,omitempty"`

	// Output only. Region in which backup is stored.
	//  Format: `projects/{project_id}/locations/{location}`
	// +kcc:proto:field=google.cloud.netapp.v1.Backup.backup_region
	BackupRegion *string `json:"backupRegion,omitempty"`

	// Output only. The time until which the backup is not deletable.
	// +kcc:proto:field=google.cloud.netapp.v1.Backup.enforced_retention_end_time
	EnforcedRetentionEndTime *string `json:"enforcedRetentionEndTime,omitempty"`
}

// +genclient
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:resource:categories=gcp,shortName=gcpnetappbackup;gcpnetappbackups
// +kubebuilder:subresource:status
// +kubebuilder:metadata:labels="cnrm.cloud.google.com/managed-by-kcc=true"
// +kubebuilder:metadata:labels="cnrm.cloud.google.com/system=true"
// +kubebuilder:printcolumn:name="Age",JSONPath=".metadata.creationTimestamp",type="date"
// +kubebuilder:printcolumn:name="Ready",JSONPath=".status.conditions[?(@.type=='Ready')].status",type="string",description="When 'True', the most recent reconcile of the resource succeeded"
// +kubebuilder:printcolumn:name="Status",JSONPath=".status.conditions[?(@.type=='Ready')].reason",type="string",description="The reason for the value in 'Ready'"
// +kubebuilder:printcolumn:name="Status Age",JSONPath=".status.conditions[?(@.type=='Ready')].lastTransitionTime",type="date",description="The last transition time for the value in 'Status'"

// NetappBackup is the Schema for the NetappBackup API
// +k8s:openapi-gen=true
type NetappBackup struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	// +required
	Spec   NetappBackupSpec   `json:"spec,omitempty"`
	Status NetappBackupStatus `json:"status,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// NetappBackupList contains a list of NetappBackup
type NetappBackupList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []NetappBackup `json:"items"`
}

func init() {
	SchemeBuilder.Register(&NetappBackup{}, &NetappBackupList{})
}
