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


// +kcc:proto=google.cloud.filestore.v1beta1.Backup
type Backup struct {

	// A description of the backup with 2048 characters or less.
	//  Requests with longer descriptions will be rejected.
	// +kcc:proto:field=google.cloud.filestore.v1beta1.Backup.description
	Description *string `json:"description,omitempty"`

	// Resource labels to represent user provided metadata.
	// +kcc:proto:field=google.cloud.filestore.v1beta1.Backup.labels
	Labels map[string]string `json:"labels,omitempty"`

	// The resource name of the source Filestore instance, in the format
	//  `projects/{project_id}/locations/{location_id}/instances/{instance_id}`,
	//  used to create this backup.
	// +kcc:proto:field=google.cloud.filestore.v1beta1.Backup.source_instance
	SourceInstance *string `json:"sourceInstance,omitempty"`

	// Name of the file share in the source Filestore instance that the
	//  backup is created from.
	// +kcc:proto:field=google.cloud.filestore.v1beta1.Backup.source_file_share
	SourceFileShare *string `json:"sourceFileShare,omitempty"`

	// Immutable. KMS key name used for data encryption.
	// +kcc:proto:field=google.cloud.filestore.v1beta1.Backup.kms_key_name
	KMSKeyName *string `json:"kmsKeyName,omitempty"`
}

// +kcc:proto=google.cloud.filestore.v1beta1.Backup
type BackupObservedState struct {
	// Output only. The resource name of the backup, in the format
	//  `projects/{project_id}/locations/{location_id}/backups/{backup_id}`.
	// +kcc:proto:field=google.cloud.filestore.v1beta1.Backup.name
	Name *string `json:"name,omitempty"`

	// Output only. The backup state.
	// +kcc:proto:field=google.cloud.filestore.v1beta1.Backup.state
	State *string `json:"state,omitempty"`

	// Output only. The time when the backup was created.
	// +kcc:proto:field=google.cloud.filestore.v1beta1.Backup.create_time
	CreateTime *string `json:"createTime,omitempty"`

	// Output only. Capacity of the source file share when the backup was created.
	// +kcc:proto:field=google.cloud.filestore.v1beta1.Backup.capacity_gb
	CapacityGB *int64 `json:"capacityGB,omitempty"`

	// Output only. The size of the storage used by the backup. As backups share
	//  storage, this number is expected to change with backup creation/deletion.
	// +kcc:proto:field=google.cloud.filestore.v1beta1.Backup.storage_bytes
	StorageBytes *int64 `json:"storageBytes,omitempty"`

	// Output only. The service tier of the source Filestore instance that this
	//  backup is created from.
	// +kcc:proto:field=google.cloud.filestore.v1beta1.Backup.source_instance_tier
	SourceInstanceTier *string `json:"sourceInstanceTier,omitempty"`

	// Output only. Amount of bytes that will be downloaded if the backup is
	//  restored
	// +kcc:proto:field=google.cloud.filestore.v1beta1.Backup.download_bytes
	DownloadBytes *int64 `json:"downloadBytes,omitempty"`

	// Output only. Reserved for future use.
	// +kcc:proto:field=google.cloud.filestore.v1beta1.Backup.satisfies_pzs
	SatisfiesPzs *bool `json:"satisfiesPzs,omitempty"`

	// Output only. Reserved for future use.
	// +kcc:proto:field=google.cloud.filestore.v1beta1.Backup.satisfies_pzi
	SatisfiesPzi *bool `json:"satisfiesPzi,omitempty"`
}
