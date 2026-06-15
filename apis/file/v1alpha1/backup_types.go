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

var FilestoreBackupGVK = GroupVersion.WithKind("FilestoreBackup")

// FilestoreBackupSpec defines the desired state of FilestoreBackup
// +kcc:spec:proto=google.cloud.filestore.v1.Backup
type FilestoreBackupSpec struct {
	// The project that this resource belongs to.
	// +required
	ProjectRef *refsv1beta1.ProjectRef `json:"projectRef"`

	// Immutable. The location of this resource.
	// +required
	Location string `json:"location"`

	// The FilestoreBackup name. If not given, the metadata.name will be used.
	ResourceID *string `json:"resourceID,omitempty"`

	// A description of the backup with 2048 characters or less.
	// Requests with longer descriptions will be rejected.
	// +kcc:proto:field=google.cloud.filestore.v1.Backup.description
	// +optional
	Description *string `json:"description,omitempty"`

	// Resource labels to represent user provided metadata.
	// +kcc:proto:field=google.cloud.filestore.v1.Backup.labels
	// +optional
	Labels map[string]string `json:"labels,omitempty"`

	// The resource name of the source Filestore instance, used to create this backup.
	// +kcc:proto:field=google.cloud.filestore.v1.Backup.source_instance
	// +required
	SourceInstanceRef *refsv1beta1.FilestoreInstanceRef `json:"sourceInstanceRef"`

	// Name of the file share in the source Filestore instance that the
	// backup is created from.
	// +kcc:proto:field=google.cloud.filestore.v1.Backup.source_file_share
	// +required
	SourceFileShare *string `json:"sourceFileShare"`

	// Immutable. KMS key name used for data encryption.
	// +kcc:proto:field=google.cloud.filestore.v1.Backup.kms_key
	// +optional
	KmsKeyRef *refsv1beta1.KMSCryptoKeyRef `json:"kmsKeyRef,omitempty"`

	// Optional. Input only. Immutable. Tag key-value pairs bound to this
	// resource. Each key must be a namespaced name and each value a short name.
	// +kcc:proto:field=google.cloud.filestore.v1.Backup.tags
	// +optional
	Tags map[string]string `json:"tags,omitempty"`
}

// FilestoreBackupStatus defines the config connector machine state of FilestoreBackup
type FilestoreBackupStatus struct {
	/* Conditions represent the latest available observations of the
	   object's current state. */
	Conditions []v1alpha1.Condition `json:"conditions,omitempty"`

	// ObservedGeneration is the generation of the resource that was most recently observed by the Config Connector controller. If this is equal to metadata.generation, then that means that the current reported status reflects the most recent desired state of the resource.
	ObservedGeneration *int64 `json:"observedGeneration,omitempty"`

	// A unique specifier for the FilestoreBackup resource in GCP.
	ExternalRef *string `json:"externalRef,omitempty"`

	// ObservedState is the state of the resource as most recently observed in GCP.
	ObservedState *FilestoreBackupObservedState `json:"observedState,omitempty"`
}

// FilestoreBackupObservedState is the state of the FilestoreBackup resource as most recently observed in GCP.
// +kcc:observedstate:proto=google.cloud.filestore.v1.Backup
type FilestoreBackupObservedState struct {
	// Output only. The resource name of the backup, in the format
	// `projects/{project_number}/locations/{location_id}/backups/{backup_id}`.
	// +kcc:proto:field=google.cloud.filestore.v1.Backup.name
	Name *string `json:"name,omitempty"`

	// Output only. The backup state.
	// +kcc:proto:field=google.cloud.filestore.v1.Backup.state
	State *string `json:"state,omitempty"`

	// Output only. The time when the backup was created.
	// +kcc:proto:field=google.cloud.filestore.v1.Backup.create_time
	CreateTime *string `json:"createTime,omitempty"`

	// Output only. Capacity of the source file share when the backup was created.
	// +kcc:proto:field=google.cloud.filestore.v1.Backup.capacity_gb
	CapacityGB *int64 `json:"capacityGB,omitempty"`

	// Output only. The size of the storage used by the backup. As backups share
	// storage, this number is expected to change with backup creation/deletion.
	// +kcc:proto:field=google.cloud.filestore.v1.Backup.storage_bytes
	StorageBytes *int64 `json:"storageBytes,omitempty"`

	// Output only. The service tier of the source Filestore instance that this
	// backup is created from.
	// +kcc:proto:field=google.cloud.filestore.v1.Backup.source_instance_tier
	SourceInstanceTier *string `json:"sourceInstanceTier,omitempty"`

	// Output only. Amount of bytes that will be downloaded if the backup is
	// restored. This may be different than storage bytes, since sequential
	// backups of the same disk will share storage.
	// +kcc:proto:field=google.cloud.filestore.v1.Backup.download_bytes
	DownloadBytes *int64 `json:"downloadBytes,omitempty"`

	// Output only. Reserved for future use.
	// +kcc:proto:field=google.cloud.filestore.v1.Backup.satisfies_pzs
	SatisfiesPzs *bool `json:"satisfiesPzs,omitempty"`

	// Output only. Reserved for future use.
	// +kcc:proto:field=google.cloud.filestore.v1.Backup.satisfies_pzi
	SatisfiesPzi *bool `json:"satisfiesPzi,omitempty"`

	// Output only. The file system protocol of the source Filestore instance that
	// this backup is created from.
	// +kcc:proto:field=google.cloud.filestore.v1.Backup.file_system_protocol
	FileSystemProtocol *string `json:"fileSystemProtocol,omitempty"`
}

// +genclient
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:resource:categories=gcp,shortName=gcpfilestorebackup;gcpfilestorebackups
// +kubebuilder:subresource:status
// +kubebuilder:metadata:labels="cnrm.cloud.google.com/managed-by-kcc=true"
// +kubebuilder:metadata:labels="cnrm.cloud.google.com/stability-level=alpha"
// +kubebuilder:metadata:labels="cnrm.cloud.google.com/system=true"
// +kubebuilder:printcolumn:name="Age",JSONPath=".metadata.creationTimestamp",type="date"
// +kubebuilder:printcolumn:name="Ready",JSONPath=".status.conditions[?(@.type=='Ready')].status",type="string",description="When 'True', the most recent reconcile of the resource succeeded"
// +kubebuilder:printcolumn:name="Status",JSONPath=".status.conditions[?(@.type=='Ready')].reason",type="string",description="The reason for the value in 'Ready'"
// +kubebuilder:printcolumn:name="Status Age",JSONPath=".status.conditions[?(@.type=='Ready')].lastTransitionTime",type="date",description="The last transition time for the value in 'Status'"

// FilestoreBackup is the Schema for the FilestoreBackup API
// +k8s:openapi-gen=true
type FilestoreBackup struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	// +required
	Spec   FilestoreBackupSpec   `json:"spec,omitempty"`
	Status FilestoreBackupStatus `json:"status,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// FilestoreBackupList contains a list of FilestoreBackup
type FilestoreBackupList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []FilestoreBackup `json:"items"`
}

func init() {
	SchemeBuilder.Register(&FilestoreBackup{}, &FilestoreBackupList{})
}
