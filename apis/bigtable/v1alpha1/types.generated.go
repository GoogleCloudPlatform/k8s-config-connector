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


// +kcc:proto=google.bigtable.admin.v2.Backup
type Backup struct {
	// A globally unique identifier for the backup which cannot be
	//  changed. Values are of the form
	//  `projects/{project}/instances/{instance}/clusters/{cluster}/
	//     backups/[_a-zA-Z0-9][-_.a-zA-Z0-9]*`
	//  The final segment of the name must be between 1 and 50 characters
	//  in length.
	//
	//  The backup is stored in the cluster identified by the prefix of the backup
	//  name of the form
	//  `projects/{project}/instances/{instance}/clusters/{cluster}`.
	// +kcc:proto:field=google.bigtable.admin.v2.Backup.name
	Name *string `json:"name,omitempty"`

	// Required. Immutable. Name of the table from which this backup was created.
	//  This needs to be in the same instance as the backup. Values are of the form
	//  `projects/{project}/instances/{instance}/tables/{source_table}`.
	// +kcc:proto:field=google.bigtable.admin.v2.Backup.source_table
	SourceTable *string `json:"sourceTable,omitempty"`

	// Required. The expiration time of the backup.
	//  When creating a backup or updating its `expire_time`, the value must be
	//  greater than the backup creation time by:
	//  - At least 6 hours
	//  - At most 90 days
	//
	//  Once the `expire_time` has passed, Cloud Bigtable will delete the backup.
	// +kcc:proto:field=google.bigtable.admin.v2.Backup.expire_time
	ExpireTime *string `json:"expireTime,omitempty"`

	// Indicates the backup type of the backup.
	// +kcc:proto:field=google.bigtable.admin.v2.Backup.backup_type
	BackupType *string `json:"backupType,omitempty"`

	// The time at which the hot backup will be converted to a standard backup.
	//  Once the `hot_to_standard_time` has passed, Cloud Bigtable will convert the
	//  hot backup to a standard backup. This value must be greater than the backup
	//  creation time by:
	//  - At least 24 hours
	//
	//  This field only applies for hot backups. When creating or updating a
	//  standard backup, attempting to set this field will fail the request.
	// +kcc:proto:field=google.bigtable.admin.v2.Backup.hot_to_standard_time
	HotToStandardTime *string `json:"hotToStandardTime,omitempty"`
}

// +kcc:proto=google.bigtable.admin.v2.EncryptionInfo
type EncryptionInfo struct {
}

// +kcc:proto=google.protobuf.Any
type Any struct {
	// A URL/resource name that uniquely identifies the type of the serialized
	//  protocol buffer message. This string must contain at least
	//  one "/" character. The last segment of the URL's path must represent
	//  the fully qualified name of the type (as in
	//  `path/google.protobuf.Duration`). The name should be in a canonical form
	//  (e.g., leading "." is not accepted).
	//
	//  In practice, teams usually precompile into the binary all types that they
	//  expect it to use in the context of Any. However, for URLs which use the
	//  scheme `http`, `https`, or no scheme, one can optionally set up a type
	//  server that maps type URLs to message definitions as follows:
	//
	//  * If no scheme is provided, `https` is assumed.
	//  * An HTTP GET on the URL must yield a [google.protobuf.Type][]
	//    value in binary format, or produce an error.
	//  * Applications are allowed to cache lookup results based on the
	//    URL, or have them precompiled into a binary to avoid any
	//    lookup. Therefore, binary compatibility needs to be preserved
	//    on changes to types. (Use versioned type names to manage
	//    breaking changes.)
	//
	//  Note: this functionality is not currently available in the official
	//  protobuf release, and it is not used for type URLs beginning with
	//  type.googleapis.com.
	//
	//  Schemes other than `http`, `https` (or the empty scheme) might be
	//  used with implementation specific semantics.
	// +kcc:proto:field=google.protobuf.Any.type_url
	TypeURL *string `json:"typeURL,omitempty"`

	// Must be a valid serialized protocol buffer of the above specified type.
	// +kcc:proto:field=google.protobuf.Any.value
	Value []byte `json:"value,omitempty"`
}

// +kcc:proto=google.rpc.Status
type Status struct {
	// The status code, which should be an enum value of
	//  [google.rpc.Code][google.rpc.Code].
	// +kcc:proto:field=google.rpc.Status.code
	Code *int32 `json:"code,omitempty"`

	// A developer-facing error message, which should be in English. Any
	//  user-facing error message should be localized and sent in the
	//  [google.rpc.Status.details][google.rpc.Status.details] field, or localized
	//  by the client.
	// +kcc:proto:field=google.rpc.Status.message
	Message *string `json:"message,omitempty"`

	// A list of messages that carry the error details.  There is a common set of
	//  message types for APIs to use.
	// +kcc:proto:field=google.rpc.Status.details
	Details []Any `json:"details,omitempty"`
}

// +kcc:proto=google.bigtable.admin.v2.Backup
type BackupObservedState struct {
	// Output only. Name of the backup from which this backup was copied. If a
	//  backup is not created by copying a backup, this field will be empty. Values
	//  are of the form:
	//  projects/<project>/instances/<instance>/clusters/<cluster>/backups/<backup>
	// +kcc:proto:field=google.bigtable.admin.v2.Backup.source_backup
	SourceBackup *string `json:"sourceBackup,omitempty"`

	// Output only. `start_time` is the time that the backup was started
	//  (i.e. approximately the time the
	//  [CreateBackup][google.bigtable.admin.v2.BigtableTableAdmin.CreateBackup]
	//  request is received).  The row data in this backup will be no older than
	//  this timestamp.
	// +kcc:proto:field=google.bigtable.admin.v2.Backup.start_time
	StartTime *string `json:"startTime,omitempty"`

	// Output only. `end_time` is the time that the backup was finished. The row
	//  data in the backup will be no newer than this timestamp.
	// +kcc:proto:field=google.bigtable.admin.v2.Backup.end_time
	EndTime *string `json:"endTime,omitempty"`

	// Output only. Size of the backup in bytes.
	// +kcc:proto:field=google.bigtable.admin.v2.Backup.size_bytes
	SizeBytes *int64 `json:"sizeBytes,omitempty"`

	// Output only. The current state of the backup.
	// +kcc:proto:field=google.bigtable.admin.v2.Backup.state
	State *string `json:"state,omitempty"`

	// Output only. The encryption information for the backup.
	// +kcc:proto:field=google.bigtable.admin.v2.Backup.encryption_info
	EncryptionInfo *EncryptionInfo `json:"encryptionInfo,omitempty"`
}

// +kcc:proto=google.bigtable.admin.v2.EncryptionInfo
type EncryptionInfoObservedState struct {
	// Output only. The type of encryption used to protect this resource.
	// +kcc:proto:field=google.bigtable.admin.v2.EncryptionInfo.encryption_type
	EncryptionType *string `json:"encryptionType,omitempty"`

	// Output only. The status of encrypt/decrypt calls on underlying data for
	//  this resource. Regardless of status, the existing data is always encrypted
	//  at rest.
	// +kcc:proto:field=google.bigtable.admin.v2.EncryptionInfo.encryption_status
	EncryptionStatus *Status `json:"encryptionStatus,omitempty"`

	// Output only. The version of the Cloud KMS key specified in the parent
	//  cluster that is in use for the data underlying this table.
	// +kcc:proto:field=google.bigtable.admin.v2.EncryptionInfo.kms_key_version
	KMSKeyVersion *string `json:"kmsKeyVersion,omitempty"`
}
