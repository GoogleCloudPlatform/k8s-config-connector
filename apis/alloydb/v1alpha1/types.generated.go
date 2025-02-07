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


// +kcc:proto=google.cloud.alloydb.v1.Backup
type Backup struct {

	// User-settable and human-readable display name for the Backup.
	// +kcc:proto:field=google.cloud.alloydb.v1.Backup.display_name
	DisplayName *string `json:"displayName,omitempty"`

	// Labels as key value pairs
	// +kcc:proto:field=google.cloud.alloydb.v1.Backup.labels
	Labels map[string]string `json:"labels,omitempty"`

	// The backup type, which suggests the trigger for the backup.
	// +kcc:proto:field=google.cloud.alloydb.v1.Backup.type
	Type *string `json:"type,omitempty"`

	// User-provided description of the backup.
	// +kcc:proto:field=google.cloud.alloydb.v1.Backup.description
	Description *string `json:"description,omitempty"`

	// Required. The full resource name of the backup source cluster
	//  (e.g., projects/{project}/locations/{region}/clusters/{cluster_id}).
	// +kcc:proto:field=google.cloud.alloydb.v1.Backup.cluster_name
	ClusterName *string `json:"clusterName,omitempty"`

	// Optional. The encryption config can be specified to encrypt the
	//  backup with a customer-managed encryption key (CMEK). When this field is
	//  not specified, the backup will then use default encryption scheme to
	//  protect the user data.
	// +kcc:proto:field=google.cloud.alloydb.v1.Backup.encryption_config
	EncryptionConfig *EncryptionConfig `json:"encryptionConfig,omitempty"`

	// For Resource freshness validation (https://google.aip.dev/154)
	// +kcc:proto:field=google.cloud.alloydb.v1.Backup.etag
	Etag *string `json:"etag,omitempty"`

	// Annotations to allow client tools to store small amount of arbitrary data.
	//  This is distinct from labels.
	//  https://google.aip.dev/128
	// +kcc:proto:field=google.cloud.alloydb.v1.Backup.annotations
	Annotations map[string]string `json:"annotations,omitempty"`

	// Optional. Input only. Immutable. Tag keys/values directly bound to this
	//  resource. For example:
	//  ```
	//  "123/environment": "production",
	//  "123/costCenter": "marketing"
	//  ```
	// +kcc:proto:field=google.cloud.alloydb.v1.Backup.tags
	Tags map[string]string `json:"tags,omitempty"`
}

// +kcc:proto=google.cloud.alloydb.v1.Backup.QuantityBasedExpiry
type Backup_QuantityBasedExpiry struct {
}

// +kcc:proto=google.cloud.alloydb.v1.EncryptionConfig
type EncryptionConfig struct {
	// The fully-qualified resource name of the KMS key.
	//  Each Cloud KMS key is regionalized and has the following format:
	//  projects/[PROJECT]/locations/[REGION]/keyRings/[RING]/cryptoKeys/[KEY_NAME]
	// +kcc:proto:field=google.cloud.alloydb.v1.EncryptionConfig.kms_key_name
	KMSKeyName *string `json:"kmsKeyName,omitempty"`
}

// +kcc:proto=google.cloud.alloydb.v1.EncryptionInfo
type EncryptionInfo struct {
}

// +kcc:proto=google.cloud.alloydb.v1.Backup
type BackupObservedState struct {
	// Output only. The name of the backup resource with the format:
	//   * projects/{project}/locations/{region}/backups/{backup_id}
	//  where the cluster and backup ID segments should satisfy the regex
	//  expression `[a-z]([a-z0-9-]{0,61}[a-z0-9])?`, e.g. 1-63 characters of
	//  lowercase letters, numbers, and dashes, starting with a letter, and ending
	//  with a letter or number. For more details see https://google.aip.dev/122.
	//  The prefix of the backup resource name is the name of the parent
	//  resource:
	//   * projects/{project}/locations/{region}
	// +kcc:proto:field=google.cloud.alloydb.v1.Backup.name
	Name *string `json:"name,omitempty"`

	// Output only. The system-generated UID of the resource. The UID is assigned
	//  when the resource is created, and it is retained until it is deleted.
	// +kcc:proto:field=google.cloud.alloydb.v1.Backup.uid
	Uid *string `json:"uid,omitempty"`

	// Output only. Create time stamp
	// +kcc:proto:field=google.cloud.alloydb.v1.Backup.create_time
	CreateTime *string `json:"createTime,omitempty"`

	// Output only. Update time stamp
	// +kcc:proto:field=google.cloud.alloydb.v1.Backup.update_time
	UpdateTime *string `json:"updateTime,omitempty"`

	// Output only. Delete time stamp
	// +kcc:proto:field=google.cloud.alloydb.v1.Backup.delete_time
	DeleteTime *string `json:"deleteTime,omitempty"`

	// Output only. The current state of the backup.
	// +kcc:proto:field=google.cloud.alloydb.v1.Backup.state
	State *string `json:"state,omitempty"`

	// Output only. The system-generated UID of the cluster which was used to
	//  create this resource.
	// +kcc:proto:field=google.cloud.alloydb.v1.Backup.cluster_uid
	ClusterUid *string `json:"clusterUid,omitempty"`

	// Output only. Reconciling (https://google.aip.dev/128#reconciliation), if
	//  true, indicates that the service is actively updating the resource. This
	//  can happen due to user-triggered updates or system actions like failover or
	//  maintenance.
	// +kcc:proto:field=google.cloud.alloydb.v1.Backup.reconciling
	Reconciling *bool `json:"reconciling,omitempty"`

	// Output only. The encryption information for the backup.
	// +kcc:proto:field=google.cloud.alloydb.v1.Backup.encryption_info
	EncryptionInfo *EncryptionInfo `json:"encryptionInfo,omitempty"`

	// Output only. The size of the backup in bytes.
	// +kcc:proto:field=google.cloud.alloydb.v1.Backup.size_bytes
	SizeBytes *int64 `json:"sizeBytes,omitempty"`

	// Output only. The time at which after the backup is eligible to be garbage
	//  collected. It is the duration specified by the backup's retention policy,
	//  added to the backup's create_time.
	// +kcc:proto:field=google.cloud.alloydb.v1.Backup.expiry_time
	ExpiryTime *string `json:"expiryTime,omitempty"`

	// Output only. The QuantityBasedExpiry of the backup, specified by the
	//  backup's retention policy. Once the expiry quantity is over retention, the
	//  backup is eligible to be garbage collected.
	// +kcc:proto:field=google.cloud.alloydb.v1.Backup.expiry_quantity
	ExpiryQuantity *Backup_QuantityBasedExpiry `json:"expiryQuantity,omitempty"`

	// Output only. Reserved for future use.
	// +kcc:proto:field=google.cloud.alloydb.v1.Backup.satisfies_pzs
	SatisfiesPzs *bool `json:"satisfiesPzs,omitempty"`

	// Output only. The database engine major version of the cluster this backup
	//  was created from. Any restored cluster created from this backup will have
	//  the same database version.
	// +kcc:proto:field=google.cloud.alloydb.v1.Backup.database_version
	DatabaseVersion *string `json:"databaseVersion,omitempty"`
}

// +kcc:proto=google.cloud.alloydb.v1.Backup.QuantityBasedExpiry
type Backup_QuantityBasedExpiryObservedState struct {
	// Output only. The backup's position among its backups with the same source
	//  cluster and type, by descending chronological order create time(i.e.
	//  newest first).
	// +kcc:proto:field=google.cloud.alloydb.v1.Backup.QuantityBasedExpiry.retention_count
	RetentionCount *int32 `json:"retentionCount,omitempty"`

	// Output only. The length of the quantity-based queue, specified by the
	//  backup's retention policy.
	// +kcc:proto:field=google.cloud.alloydb.v1.Backup.QuantityBasedExpiry.total_retention_count
	TotalRetentionCount *int32 `json:"totalRetentionCount,omitempty"`
}

// +kcc:proto=google.cloud.alloydb.v1.EncryptionInfo
type EncryptionInfoObservedState struct {
	// Output only. Type of encryption.
	// +kcc:proto:field=google.cloud.alloydb.v1.EncryptionInfo.encryption_type
	EncryptionType *string `json:"encryptionType,omitempty"`

	// Output only. Cloud KMS key versions that are being used to protect the
	//  database or the backup.
	// +kcc:proto:field=google.cloud.alloydb.v1.EncryptionInfo.kms_key_versions
	KMSKeyVersions []string `json:"kmsKeyVersions,omitempty"`
}
