// Copyright 2025 Google LLC
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package v1alpha1


// +kcc:proto=google.cloud.gkebackup.v1.NamespacedName
type NamespacedName struct {
	// Optional. The Namespace of the Kubernetes resource.
	// +kcc:proto:field=google.cloud.gkebackup.v1.NamespacedName.namespace
	Namespace *string `json:"namespace,omitempty"`

	// Optional. The name of the Kubernetes resource.
	// +kcc:proto:field=google.cloud.gkebackup.v1.NamespacedName.name
	Name *string `json:"name,omitempty"`
}

// +kcc:proto=google.cloud.gkebackup.v1.VolumeBackup
type VolumeBackup struct {
}

// +kcc:proto=google.cloud.gkebackup.v1.VolumeBackup
type VolumeBackupObservedState struct {
	// Output only. The full name of the VolumeBackup resource.
	//  Format: `projects/*/locations/*/backupPlans/*/backups/*/volumeBackups/*`.
	// +kcc:proto:field=google.cloud.gkebackup.v1.VolumeBackup.name
	Name *string `json:"name,omitempty"`

	// Output only. Server generated global unique identifier of
	//  [UUID](https://en.wikipedia.org/wiki/Universally_unique_identifier) format.
	// +kcc:proto:field=google.cloud.gkebackup.v1.VolumeBackup.uid
	Uid *string `json:"uid,omitempty"`

	// Output only. The timestamp when this VolumeBackup resource was
	//  created.
	// +kcc:proto:field=google.cloud.gkebackup.v1.VolumeBackup.create_time
	CreateTime *string `json:"createTime,omitempty"`

	// Output only. The timestamp when this VolumeBackup resource was last
	//  updated.
	// +kcc:proto:field=google.cloud.gkebackup.v1.VolumeBackup.update_time
	UpdateTime *string `json:"updateTime,omitempty"`

	// Output only. A reference to the source Kubernetes PVC from which this
	//  VolumeBackup was created.
	// +kcc:proto:field=google.cloud.gkebackup.v1.VolumeBackup.source_pvc
	SourcePvc *NamespacedName `json:"sourcePvc,omitempty"`

	// Output only. A storage system-specific opaque handle to the underlying
	//  volume backup.
	// +kcc:proto:field=google.cloud.gkebackup.v1.VolumeBackup.volume_backup_handle
	VolumeBackupHandle *string `json:"volumeBackupHandle,omitempty"`

	// Output only. The format used for the volume backup.
	// +kcc:proto:field=google.cloud.gkebackup.v1.VolumeBackup.format
	Format *string `json:"format,omitempty"`

	// Output only. The aggregate size of the underlying artifacts associated with
	//  this VolumeBackup in the backup storage. This may change over time when
	//  multiple backups of the same volume share the same backup storage
	//  location. In particular, this is likely to increase in size when
	//  the immediately preceding backup of the same volume is deleted.
	// +kcc:proto:field=google.cloud.gkebackup.v1.VolumeBackup.storage_bytes
	StorageBytes *int64 `json:"storageBytes,omitempty"`

	// Output only. The minimum size of the disk to which this VolumeBackup can be
	//  restored.
	// +kcc:proto:field=google.cloud.gkebackup.v1.VolumeBackup.disk_size_bytes
	DiskSizeBytes *int64 `json:"diskSizeBytes,omitempty"`

	// Output only. The timestamp when the associated underlying volume backup
	//  operation completed.
	// +kcc:proto:field=google.cloud.gkebackup.v1.VolumeBackup.complete_time
	CompleteTime *string `json:"completeTime,omitempty"`

	// Output only. The current state of this VolumeBackup.
	// +kcc:proto:field=google.cloud.gkebackup.v1.VolumeBackup.state
	State *string `json:"state,omitempty"`

	// Output only. A human readable message explaining why the VolumeBackup is in
	//  its current state.
	// +kcc:proto:field=google.cloud.gkebackup.v1.VolumeBackup.state_message
	StateMessage *string `json:"stateMessage,omitempty"`

	// Output only. `etag` is used for optimistic concurrency control as a way to
	//  help prevent simultaneous updates of a volume backup from overwriting each
	//  other. It is strongly suggested that systems make use of the `etag` in the
	//  read-modify-write cycle to perform volume backup updates in order to avoid
	//  race conditions.
	// +kcc:proto:field=google.cloud.gkebackup.v1.VolumeBackup.etag
	Etag *string `json:"etag,omitempty"`
}
