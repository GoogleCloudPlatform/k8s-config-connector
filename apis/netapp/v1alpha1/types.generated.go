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


// +kcc:proto=google.cloud.netapp.v1.Backup
type Backup struct {
	// Identifier. The resource name of the backup.
	//  Format:
	//  `projects/{project_id}/locations/{location}/backupVaults/{backup_vault_id}/backups/{backup_id}`.
	// +kcc:proto:field=google.cloud.netapp.v1.Backup.name
	Name *string `json:"name,omitempty"`

	// A description of the backup with 2048 characters or less.
	//  Requests with longer descriptions will be rejected.
	// +kcc:proto:field=google.cloud.netapp.v1.Backup.description
	Description *string `json:"description,omitempty"`

	// Volume full name of this backup belongs to.
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

// +kcc:proto=google.cloud.netapp.v1.Backup
type BackupObservedState struct {
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
}
