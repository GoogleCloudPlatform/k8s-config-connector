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


// +kcc:proto=google.cloud.redis.cluster.v1beta1.Backup
type Backup struct {
	// Identifier. Full resource path of the backup. the last part of the name is
	//  the backup id with the following format: [YYYYMMDDHHMMSS]_[Shorted Cluster
	//  UID] OR customer specified while backup cluster. Example:
	//  20240515123000_1234
	// +kcc:proto:field=google.cloud.redis.cluster.v1beta1.Backup.name
	Name *string `json:"name,omitempty"`
}

// +kcc:proto=google.cloud.redis.cluster.v1beta1.BackupFile
type BackupFile struct {
}

// +kcc:proto=google.cloud.redis.cluster.v1beta1.EncryptionInfo
type EncryptionInfo struct {
}

// +kcc:proto=google.cloud.redis.cluster.v1beta1.Backup
type BackupObservedState struct {
	// Output only. The time when the backup was created.
	// +kcc:proto:field=google.cloud.redis.cluster.v1beta1.Backup.create_time
	CreateTime *string `json:"createTime,omitempty"`

	// Output only. Cluster resource path of this backup.
	// +kcc:proto:field=google.cloud.redis.cluster.v1beta1.Backup.cluster
	Cluster *string `json:"cluster,omitempty"`

	// Output only. Cluster uid of this backup.
	// +kcc:proto:field=google.cloud.redis.cluster.v1beta1.Backup.cluster_uid
	ClusterUid *string `json:"clusterUid,omitempty"`

	// Output only. Total size of the backup in bytes.
	// +kcc:proto:field=google.cloud.redis.cluster.v1beta1.Backup.total_size_bytes
	TotalSizeBytes *int64 `json:"totalSizeBytes,omitempty"`

	// Output only. The time when the backup will expire.
	// +kcc:proto:field=google.cloud.redis.cluster.v1beta1.Backup.expire_time
	ExpireTime *string `json:"expireTime,omitempty"`

	// Output only. redis-7.2, valkey-7.5
	// +kcc:proto:field=google.cloud.redis.cluster.v1beta1.Backup.engine_version
	EngineVersion *string `json:"engineVersion,omitempty"`

	// Output only. List of backup files of the backup.
	// +kcc:proto:field=google.cloud.redis.cluster.v1beta1.Backup.backup_files
	BackupFiles []BackupFile `json:"backupFiles,omitempty"`

	// Output only. Node type of the cluster.
	// +kcc:proto:field=google.cloud.redis.cluster.v1beta1.Backup.node_type
	NodeType *string `json:"nodeType,omitempty"`

	// Output only. Number of replicas for the cluster.
	// +kcc:proto:field=google.cloud.redis.cluster.v1beta1.Backup.replica_count
	ReplicaCount *int32 `json:"replicaCount,omitempty"`

	// Output only. Number of shards for the cluster.
	// +kcc:proto:field=google.cloud.redis.cluster.v1beta1.Backup.shard_count
	ShardCount *int32 `json:"shardCount,omitempty"`

	// Output only. Type of the backup.
	// +kcc:proto:field=google.cloud.redis.cluster.v1beta1.Backup.backup_type
	BackupType *string `json:"backupType,omitempty"`

	// Output only. State of the backup.
	// +kcc:proto:field=google.cloud.redis.cluster.v1beta1.Backup.state
	State *string `json:"state,omitempty"`

	// Output only. Encryption information of the backup.
	// +kcc:proto:field=google.cloud.redis.cluster.v1beta1.Backup.encryption_info
	EncryptionInfo *EncryptionInfo `json:"encryptionInfo,omitempty"`

	// Output only. System assigned unique identifier of the backup.
	// +kcc:proto:field=google.cloud.redis.cluster.v1beta1.Backup.uid
	Uid *string `json:"uid,omitempty"`
}

// +kcc:proto=google.cloud.redis.cluster.v1beta1.BackupFile
type BackupFileObservedState struct {
	// Output only. e.g: <shard-id>.rdb
	// +kcc:proto:field=google.cloud.redis.cluster.v1beta1.BackupFile.file_name
	FileName *string `json:"fileName,omitempty"`

	// Output only. Size of the backup file in bytes.
	// +kcc:proto:field=google.cloud.redis.cluster.v1beta1.BackupFile.size_bytes
	SizeBytes *int64 `json:"sizeBytes,omitempty"`

	// Output only. The time when the backup file was created.
	// +kcc:proto:field=google.cloud.redis.cluster.v1beta1.BackupFile.create_time
	CreateTime *string `json:"createTime,omitempty"`
}

// +kcc:proto=google.cloud.redis.cluster.v1beta1.EncryptionInfo
type EncryptionInfoObservedState struct {
	// Output only. Type of encryption.
	// +kcc:proto:field=google.cloud.redis.cluster.v1beta1.EncryptionInfo.encryption_type
	EncryptionType *string `json:"encryptionType,omitempty"`

	// Output only. KMS key versions that are being used to protect the data
	//  at-rest.
	// +kcc:proto:field=google.cloud.redis.cluster.v1beta1.EncryptionInfo.kms_key_versions
	KMSKeyVersions []string `json:"kmsKeyVersions,omitempty"`

	// Output only. The state of the primary version of the KMS key perceived by
	//  the system. This field is not populated in backups.
	// +kcc:proto:field=google.cloud.redis.cluster.v1beta1.EncryptionInfo.kms_key_primary_state
	KMSKeyPrimaryState *string `json:"kmsKeyPrimaryState,omitempty"`

	// Output only. The most recent time when the encryption info was updated.
	// +kcc:proto:field=google.cloud.redis.cluster.v1beta1.EncryptionInfo.last_update_time
	LastUpdateTime *string `json:"lastUpdateTime,omitempty"`
}
