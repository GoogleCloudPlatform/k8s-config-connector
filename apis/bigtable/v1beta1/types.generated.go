// Copyright 2024 Google LLC
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

package v1beta1

// +kcc:proto=google.bigtable.admin.v2.AppProfile
type AppProfile struct {
	// The unique name of the app profile. Values are of the form
	//  `projects/{project}/instances/{instance}/appProfiles/[_a-zA-Z0-9][-_.a-zA-Z0-9]*`.
	Name *string `json:"name,omitempty"`

	// Strongly validated etag for optimistic concurrency control. Preserve the
	//  value returned from `GetAppProfile` when calling `UpdateAppProfile` to
	//  fail the request if there has been a modification in the mean time. The
	//  `update_mask` of the request need not include `etag` for this protection
	//  to apply.
	//  See [Wikipedia](https://en.wikipedia.org/wiki/HTTP_ETag) and
	//  [RFC 7232](https://tools.ietf.org/html/rfc7232#section-2.3) for more
	//  details.
	Etag *string `json:"etag,omitempty"`

	// Long form description of the use case for this AppProfile.
	Description *string `json:"description,omitempty"`

	// Use a multi-cluster routing policy.
	MultiClusterRoutingUseAny *AppProfile_MultiClusterRoutingUseAny `json:"multiClusterRoutingUseAny,omitempty"`

	// Use a single-cluster routing policy.
	SingleClusterRouting *AppProfile_SingleClusterRouting `json:"singleClusterRouting,omitempty"`

	// This field has been deprecated in favor of `standard_isolation.priority`.
	//  If you set this field, `standard_isolation.priority` will be set instead.
	//
	//  The priority of requests sent using this app profile.
	Priority *string `json:"priority,omitempty"`

	// The standard options used for isolating this app profile's traffic from
	//  other use cases.
	StandardIsolation *AppProfile_StandardIsolation `json:"standardIsolation,omitempty"`

	// Specifies that this app profile is intended for read-only usage via the
	//  Data Boost feature.
	DataBoostIsolationReadOnly *AppProfile_DataBoostIsolationReadOnly `json:"dataBoostIsolationReadOnly,omitempty"`
}

// +kcc:proto=google.bigtable.admin.v2.AppProfile.DataBoostIsolationReadOnly
type AppProfile_DataBoostIsolationReadOnly struct {
	// The Compute Billing Owner for this Data Boost App Profile.
	ComputeBillingOwner *string `json:"computeBillingOwner,omitempty"`
}

// +kcc:proto=google.bigtable.admin.v2.AppProfile.MultiClusterRoutingUseAny
type AppProfile_MultiClusterRoutingUseAny struct {
	// The set of clusters to route to. The order is ignored; clusters will be
	//  tried in order of distance. If left empty, all clusters are eligible.
	ClusterIds []string `json:"clusterIds,omitempty"`
}

// +kcc:proto=google.bigtable.admin.v2.AppProfile.SingleClusterRouting
type AppProfile_SingleClusterRouting struct {
	// The cluster to which read/write requests should be routed.
	ClusterID *string `json:"clusterID,omitempty"`

	// Whether or not `CheckAndMutateRow` and `ReadModifyWriteRow` requests are
	//  allowed by this app profile. It is unsafe to send these requests to
	//  the same table/row/column in multiple clusters.
	AllowTransactionalWrites *bool `json:"allowTransactionalWrites,omitempty"`
}

// +kcc:proto=google.bigtable.admin.v2.AppProfile.StandardIsolation
type AppProfile_StandardIsolation struct {
	// The priority of requests sent using this app profile.
	Priority *string `json:"priority,omitempty"`
}

// +kcc:proto=google.bigtable.admin.v2.AuthorizedView
type AuthorizedView struct {
	// Identifier. The name of this AuthorizedView.
	//  Values are of the form
	//  `projects/{project}/instances/{instance}/tables/{table}/authorizedViews/{authorized_view}`
	Name *string `json:"name,omitempty"`

	// An AuthorizedView permitting access to an explicit subset of a Table.
	SubsetView *AuthorizedView_SubsetView `json:"subsetView,omitempty"`

	// The etag for this AuthorizedView.
	//  If this is provided on update, it must match the server's etag. The server
	//  returns ABORTED error on a mismatched etag.
	Etag *string `json:"etag,omitempty"`

	// Set to true to make the AuthorizedView protected against deletion.
	//  The parent Table and containing Instance cannot be deleted if an
	//  AuthorizedView has this bit set.
	DeletionProtection *bool `json:"deletionProtection,omitempty"`
}

// +kcc:proto=google.bigtable.admin.v2.AuthorizedView.FamilySubsets
type AuthorizedView_FamilySubsets struct {
	// Individual exact column qualifiers to be included in the AuthorizedView.
	Qualifiers [][]byte `json:"qualifiers,omitempty"`

	// Prefixes for qualifiers to be included in the AuthorizedView. Every
	//  qualifier starting with one of these prefixes is included in the
	//  AuthorizedView. To provide access to all qualifiers, include the empty
	//  string as a prefix
	//  ("").
	QualifierPrefixes [][]byte `json:"qualifierPrefixes,omitempty"`
}

// +kcc:proto=google.bigtable.admin.v2.AuthorizedView.SubsetView
type AuthorizedView_SubsetView struct {
	// Row prefixes to be included in the AuthorizedView.
	//  To provide access to all rows, include the empty string as a prefix ("").
	RowPrefixes [][]byte `json:"rowPrefixes,omitempty"`

	// TODO: map type string message for family_subsets

}

// +kcc:proto=google.bigtable.admin.v2.AutoscalingLimits
type AutoscalingLimits struct {
	// Required. Minimum number of nodes to scale down to.
	MinServeNodes *int32 `json:"minServeNodes,omitempty"`

	// Required. Maximum number of nodes to scale up to.
	MaxServeNodes *int32 `json:"maxServeNodes,omitempty"`
}

// +kcc:proto=google.bigtable.admin.v2.AutoscalingTargets
type AutoscalingTargets struct {
	// The cpu utilization that the Autoscaler should be trying to achieve.
	//  This number is on a scale from 0 (no utilization) to
	//  100 (total utilization), and is limited between 10 and 80, otherwise it
	//  will return INVALID_ARGUMENT error.
	CpuUtilizationPercent *int32 `json:"cpuUtilizationPercent,omitempty"`

	// The storage utilization that the Autoscaler should be trying to achieve.
	//  This number is limited between 2560 (2.5TiB) and 5120 (5TiB) for a SSD
	//  cluster and between 8192 (8TiB) and 16384 (16TiB) for an HDD cluster,
	//  otherwise it will return INVALID_ARGUMENT error. If this value is set to 0,
	//  it will be treated as if it were set to the default value: 2560 for SSD,
	//  8192 for HDD.
	StorageUtilizationGibPerNode *int32 `json:"storageUtilizationGibPerNode,omitempty"`
}

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
	Name *string `json:"name,omitempty"`

	// Required. Immutable. Name of the table from which this backup was created.
	//  This needs to be in the same instance as the backup. Values are of the form
	//  `projects/{project}/instances/{instance}/tables/{source_table}`.
	SourceTable *string `json:"sourceTable,omitempty"`

	// Output only. Name of the backup from which this backup was copied. If a
	//  backup is not created by copying a backup, this field will be empty. Values
	//  are of the form: projects/<project>/instances/<instance>/backups/<backup>.
	SourceBackup *string `json:"sourceBackup,omitempty"`

	// Required. The expiration time of the backup, with microseconds
	//  granularity that must be at least 6 hours and at most 90 days
	//  from the time the request is received. Once the `expire_time`
	//  has passed, Cloud Bigtable will delete the backup and free the
	//  resources used by the backup.
	ExpireTime *string `json:"expireTime,omitempty"`

	// Output only. `start_time` is the time that the backup was started
	//  (i.e. approximately the time the
	//  [CreateBackup][google.bigtable.admin.v2.BigtableTableAdmin.CreateBackup]
	//  request is received).  The row data in this backup will be no older than
	//  this timestamp.
	StartTime *string `json:"startTime,omitempty"`

	// Output only. `end_time` is the time that the backup was finished. The row
	//  data in the backup will be no newer than this timestamp.
	EndTime *string `json:"endTime,omitempty"`

	// Output only. Size of the backup in bytes.
	SizeBytes *int64 `json:"sizeBytes,omitempty"`

	// Output only. The current state of the backup.
	State *string `json:"state,omitempty"`

	// Output only. The encryption information for the backup.
	EncryptionInfo *EncryptionInfo `json:"encryptionInfo,omitempty"`
}

// +kcc:proto=google.bigtable.admin.v2.BackupInfo
type BackupInfo struct {
	// Output only. Name of the backup.
	Backup *string `json:"backup,omitempty"`

	// Output only. The time that the backup was started. Row data in the backup
	//  will be no older than this timestamp.
	StartTime *string `json:"startTime,omitempty"`

	// Output only. This time that the backup was finished. Row data in the
	//  backup will be no newer than this timestamp.
	EndTime *string `json:"endTime,omitempty"`

	// Output only. Name of the table the backup was created from.
	SourceTable *string `json:"sourceTable,omitempty"`

	// Output only. Name of the backup from which this backup was copied. If a
	//  backup is not created by copying a backup, this field will be empty. Values
	//  are of the form: projects/<project>/instances/<instance>/backups/<backup>.
	SourceBackup *string `json:"sourceBackup,omitempty"`
}

// +kcc:proto=google.bigtable.admin.v2.ChangeStreamConfig
type ChangeStreamConfig struct {
	// How long the change stream should be retained. Change stream data older
	//  than the retention period will not be returned when reading the change
	//  stream from the table.
	//  Values must be at least 1 day and at most 7 days, and will be truncated to
	//  microsecond granularity.
	RetentionPeriod *string `json:"retentionPeriod,omitempty"`
}

// +kcc:proto=google.bigtable.admin.v2.Cluster
type Cluster struct {
	// The unique name of the cluster. Values are of the form
	//  `projects/{project}/instances/{instance}/clusters/[a-z][-a-z0-9]*`.
	Name *string `json:"name,omitempty"`

	// Immutable. The location where this cluster's nodes and storage reside. For
	//  best performance, clients should be located as close as possible to this
	//  cluster. Currently only zones are supported, so values should be of the
	//  form `projects/{project}/locations/{zone}`.
	Location *string `json:"location,omitempty"`

	// Output only. The current state of the cluster.
	State *string `json:"state,omitempty"`

	// The number of nodes allocated to this cluster. More nodes enable higher
	//  throughput and more consistent performance.
	ServeNodes *int32 `json:"serveNodes,omitempty"`

	// Configuration for this cluster.
	ClusterConfig *Cluster_ClusterConfig `json:"clusterConfig,omitempty"`

	// Immutable. The type of storage used by this cluster to serve its
	//  parent instance's tables, unless explicitly overridden.
	DefaultStorageType *string `json:"defaultStorageType,omitempty"`

	// Immutable. The encryption configuration for CMEK-protected clusters.
	EncryptionConfig *Cluster_EncryptionConfig `json:"encryptionConfig,omitempty"`
}

// +kcc:proto=google.bigtable.admin.v2.Cluster.ClusterAutoscalingConfig
type Cluster_ClusterAutoscalingConfig struct {
	// Required. Autoscaling limits for this cluster.
	AutoscalingLimits *AutoscalingLimits `json:"autoscalingLimits,omitempty"`

	// Required. Autoscaling targets for this cluster.
	AutoscalingTargets *AutoscalingTargets `json:"autoscalingTargets,omitempty"`
}

// +kcc:proto=google.bigtable.admin.v2.Cluster.ClusterConfig
type Cluster_ClusterConfig struct {
	// Autoscaling configuration for this cluster.
	ClusterAutoscalingConfig *Cluster_ClusterAutoscalingConfig `json:"clusterAutoscalingConfig,omitempty"`
}

// +kcc:proto=google.bigtable.admin.v2.Cluster.EncryptionConfig
type Cluster_EncryptionConfig struct {
	// Describes the Cloud KMS encryption key that will be used to protect the
	//  destination Bigtable cluster. The requirements for this key are:
	//   1) The Cloud Bigtable service account associated with the project that
	//   contains this cluster must be granted the
	//   `cloudkms.cryptoKeyEncrypterDecrypter` role on the CMEK key.
	//   2) Only regional keys can be used and the region of the CMEK key must
	//   match the region of the cluster.
	//   3) All clusters within an instance must use the same CMEK key.
	//  Values are of the form
	//  `projects/{project}/locations/{location}/keyRings/{keyring}/cryptoKeys/{key}`
	KmsKeyName *string `json:"kmsKeyName,omitempty"`
}

// +kcc:proto=google.bigtable.admin.v2.ColumnFamily
type ColumnFamily struct {
	// Garbage collection rule specified as a protobuf.
	//  Must serialize to at most 500 bytes.
	//
	//  NOTE: Garbage collection executes opportunistically in the background, and
	//  so it's possible for reads to return a cell even if it matches the active
	//  GC expression for its family.
	GcRule *GcRule `json:"gcRule,omitempty"`

	// The type of data stored in each of this family's cell values, including its
	//  full encoding. If omitted, the family only serves raw untyped bytes.
	//
	//  For now, only the `Aggregate` type is supported.
	//
	//  `Aggregate` can only be set at family creation and is immutable afterwards.
	//
	//
	//  If `value_type` is `Aggregate`, written data must be compatible with:
	//   * `value_type.input_type` for `AddInput` mutations
	ValueType *Type `json:"valueType,omitempty"`
}

// +kcc:proto=google.bigtable.admin.v2.DataBoostReadLocalWrites
type DataBoostReadLocalWrites struct {
}

// +kcc:proto=google.bigtable.admin.v2.GcRule
type GcRule struct {
	// Delete all cells in a column except the most recent N.
	MaxNumVersions *int32 `json:"maxNumVersions,omitempty"`

	// Delete cells in a column older than the given age.
	//  Values must be at least one millisecond, and will be truncated to
	//  microsecond granularity.
	MaxAge *string `json:"maxAge,omitempty"`

	// Delete cells that would be deleted by every nested rule.
	Intersection *GcRule_Intersection `json:"intersection,omitempty"`

	// Delete cells that would be deleted by any nested rule.
	Union *GcRule_Union `json:"union,omitempty"`
}

// +kcc:proto=google.bigtable.admin.v2.GcRule.Intersection
type GcRule_Intersection struct {
	// Only delete cells which would be deleted by every element of `rules`.
	Rules []GcRule `json:"rules,omitempty"`
}

// +kcc:proto=google.bigtable.admin.v2.GcRule.Union
type GcRule_Union struct {
	// Delete cells which would be deleted by any element of `rules`.
	Rules []GcRule `json:"rules,omitempty"`
}

// +kcc:proto=google.bigtable.admin.v2.HotTablet
type HotTablet struct {
	// The unique name of the hot tablet. Values are of the form
	//  `projects/{project}/instances/{instance}/clusters/{cluster}/hotTablets/[a-zA-Z0-9_-]*`.
	Name *string `json:"name,omitempty"`

	// Name of the table that contains the tablet. Values are of the form
	//  `projects/{project}/instances/{instance}/tables/[_a-zA-Z0-9][-_.a-zA-Z0-9]*`.
	TableName *string `json:"tableName,omitempty"`

	// Output only. The start time of the hot tablet.
	StartTime *string `json:"startTime,omitempty"`

	// Output only. The end time of the hot tablet.
	EndTime *string `json:"endTime,omitempty"`

	// Tablet Start Key (inclusive).
	StartKey *string `json:"startKey,omitempty"`

	// Tablet End Key (inclusive).
	EndKey *string `json:"endKey,omitempty"`

	// Output only. The average CPU usage spent by a node on this tablet over the
	//  start_time to end_time time range. The percentage is the amount of CPU used
	//  by the node to serve the tablet, from 0% (tablet was not interacted with)
	//  to 100% (the node spent all cycles serving the hot tablet).
	NodeCpuUsagePercent *float32 `json:"nodeCpuUsagePercent,omitempty"`
}

// +kcc:proto=google.bigtable.admin.v2.Instance
type Instance struct {
	// The unique name of the instance. Values are of the form
	//  `projects/{project}/instances/[a-z][a-z0-9\\-]+[a-z0-9]`.
	Name *string `json:"name,omitempty"`

	// Required. The descriptive name for this instance as it appears in UIs.
	//  Can be changed at any time, but should be kept globally unique
	//  to avoid confusion.
	DisplayName *string `json:"displayName,omitempty"`

	// (`OutputOnly`)
	//  The current state of the instance.
	State *string `json:"state,omitempty"`

	// The type of the instance. Defaults to `PRODUCTION`.
	Type *string `json:"type,omitempty"`

	// Labels are a flexible and lightweight mechanism for organizing cloud
	//  resources into groups that reflect a customer's organizational needs and
	//  deployment strategies. They can be used to filter resources and aggregate
	//  metrics.
	//
	//  * Label keys must be between 1 and 63 characters long and must conform to
	//    the regular expression: `[\p{Ll}\p{Lo}][\p{Ll}\p{Lo}\p{N}_-]{0,62}`.
	//  * Label values must be between 0 and 63 characters long and must conform to
	//    the regular expression: `[\p{Ll}\p{Lo}\p{N}_-]{0,63}`.
	//  * No more than 64 labels can be associated with a given resource.
	//  * Keys and values must both be under 128 bytes.
	Labels map[string]string `json:"labels,omitempty"`

	// Output only. A server-assigned timestamp representing when this Instance
	//  was created. For instances created before this field was added (August
	//  2021), this value is `seconds: 0, nanos: 1`.
	CreateTime *string `json:"createTime,omitempty"`

	// Output only. Reserved for future use.
	SatisfiesPzs *bool `json:"satisfiesPzs,omitempty"`
}

// +kcc:proto=google.bigtable.admin.v2.OperationProgress
type OperationProgress struct {
	// Percent completion of the operation.
	//  Values are between 0 and 100 inclusive.
	ProgressPercent *int32 `json:"progressPercent,omitempty"`

	// Time the request was received.
	StartTime *string `json:"startTime,omitempty"`

	// If set, the time at which this operation failed or was completed
	//  successfully.
	EndTime *string `json:"endTime,omitempty"`
}

// +kcc:proto=google.bigtable.admin.v2.RestoreInfo
type RestoreInfo struct {
	// The type of the restore source.
	SourceType *string `json:"sourceType,omitempty"`

	// Information about the backup used to restore the table. The backup
	//  may no longer exist.
	BackupInfo *BackupInfo `json:"backupInfo,omitempty"`
}

// +kcc:proto=google.bigtable.admin.v2.Snapshot
type Snapshot struct {
	// The unique name of the snapshot.
	//  Values are of the form
	//  `projects/{project}/instances/{instance}/clusters/{cluster}/snapshots/{snapshot}`.
	Name *string `json:"name,omitempty"`

	// Output only. The source table at the time the snapshot was taken.
	SourceTable *Table `json:"sourceTable,omitempty"`

	// Output only. The size of the data in the source table at the time the
	//  snapshot was taken. In some cases, this value may be computed
	//  asynchronously via a background process and a placeholder of 0 will be used
	//  in the meantime.
	DataSizeBytes *int64 `json:"dataSizeBytes,omitempty"`

	// Output only. The time when the snapshot is created.
	CreateTime *string `json:"createTime,omitempty"`

	// The time when the snapshot will be deleted. The maximum amount of time a
	//  snapshot can stay active is 365 days. If 'ttl' is not specified,
	//  the default maximum of 365 days will be used.
	DeleteTime *string `json:"deleteTime,omitempty"`

	// Output only. The current state of the snapshot.
	State *string `json:"state,omitempty"`

	// Description of the snapshot.
	Description *string `json:"description,omitempty"`
}

// +kcc:proto=google.bigtable.admin.v2.StandardReadRemoteWrites
type StandardReadRemoteWrites struct {
}

// +kcc:proto=google.bigtable.admin.v2.Table
type Table struct {
	// The unique name of the table. Values are of the form
	//  `projects/{project}/instances/{instance}/tables/[_a-zA-Z0-9][-_.a-zA-Z0-9]*`.
	//  Views: `NAME_ONLY`, `SCHEMA_VIEW`, `REPLICATION_VIEW`, `FULL`
	Name *string `json:"name,omitempty"`

	// TODO: map type string message for cluster_states

	// TODO: map type string message for column_families

	// Immutable. The granularity (i.e. `MILLIS`) at which timestamps are stored
	//  in this table. Timestamps not matching the granularity will be rejected. If
	//  unspecified at creation time, the value will be set to `MILLIS`. Views:
	//  `SCHEMA_VIEW`, `FULL`.
	Granularity *string `json:"granularity,omitempty"`

	// Output only. If this table was restored from another data source (e.g. a
	//  backup), this field will be populated with information about the restore.
	RestoreInfo *RestoreInfo `json:"restoreInfo,omitempty"`

	// If specified, enable the change stream on this table.
	//  Otherwise, the change stream is disabled and the change stream is not
	//  retained.
	ChangeStreamConfig *ChangeStreamConfig `json:"changeStreamConfig,omitempty"`

	// Set to true to make the table protected against data loss. i.e. deleting
	//  the following resources through Admin APIs are prohibited:
	//
	//  * The table.
	//  * The column families in the table.
	//  * The instance containing the table.
	//
	//  Note one can still delete the data stored in the table through Data APIs.
	DeletionProtection *bool `json:"deletionProtection,omitempty"`

	// If specified, automated backups are enabled for this table.
	//  Otherwise, automated backups are disabled.
	AutomatedBackupPolicy *Table_AutomatedBackupPolicy `json:"automatedBackupPolicy,omitempty"`
}

// +kcc:proto=google.bigtable.admin.v2.Table.AutomatedBackupPolicy
type Table_AutomatedBackupPolicy struct {
	// Required. How long the automated backups should be retained. The only
	//  supported value at this time is 3 days.
	RetentionPeriod *string `json:"retentionPeriod,omitempty"`

	// Required. How frequently automated backups should occur. The only
	//  supported value at this time is 24 hours.
	Frequency *string `json:"frequency,omitempty"`
}

// +kcc:proto=google.bigtable.admin.v2.Table.ClusterState
type Table_ClusterState struct {
	// Output only. The state of replication for the table in this cluster.
	ReplicationState *string `json:"replicationState,omitempty"`

	// Output only. The encryption information for the table in this cluster.
	//  If the encryption key protecting this resource is customer managed, then
	//  its version can be rotated in Cloud Key Management Service (Cloud KMS).
	//  The primary version of the key and its status will be reflected here when
	//  changes propagate from Cloud KMS.
	EncryptionInfo []EncryptionInfo `json:"encryptionInfo,omitempty"`
}

// +kcc:proto=google.bigtable.admin.v2.Type
type Type struct {
	// Bytes
	BytesType *Type_Bytes `json:"bytesType,omitempty"`

	// String
	StringType *Type_String `json:"stringType,omitempty"`

	// Int64
	Int64Type *Type_Int64 `json:"int64Type,omitempty"`

	// Aggregate
	AggregateType *Type_Aggregate `json:"aggregateType,omitempty"`
}

// +kcc:proto=google.bigtable.admin.v2.Type.Aggregate
type Type_Aggregate struct {
	// Type of the inputs that are accumulated by this `Aggregate`, which must
	//  specify a full encoding.
	//  Use `AddInput` mutations to accumulate new inputs.
	InputType *Type `json:"inputType,omitempty"`

	// Output only. Type that holds the internal accumulator state for the
	//  `Aggregate`. This is a function of the `input_type` and `aggregator`
	//  chosen, and will always specify a full encoding.
	StateType *Type `json:"stateType,omitempty"`

	// Sum aggregator.
	Sum *Type_Aggregate_Sum `json:"sum,omitempty"`
}

// +kcc:proto=google.bigtable.admin.v2.Type.Aggregate.Sum
type Type_Aggregate_Sum struct {
}

// +kcc:proto=google.bigtable.admin.v2.Type.Bytes
type Type_Bytes struct {
	// The encoding to use when converting to/from lower level types.
	Encoding *Type_Bytes_Encoding `json:"encoding,omitempty"`
}

// +kcc:proto=google.bigtable.admin.v2.Type.Bytes.Encoding
type Type_Bytes_Encoding struct {
	// Use `Raw` encoding.
	Raw *Type_Bytes_Encoding_Raw `json:"raw,omitempty"`
}

// +kcc:proto=google.bigtable.admin.v2.Type.Bytes.Encoding.Raw
type Type_Bytes_Encoding_Raw struct {
}

// +kcc:proto=google.bigtable.admin.v2.Type.Int64
type Type_Int64 struct {
	// The encoding to use when converting to/from lower level types.
	Encoding *Type_Int64_Encoding `json:"encoding,omitempty"`
}

// +kcc:proto=google.bigtable.admin.v2.Type.Int64.Encoding
type Type_Int64_Encoding struct {
	// Use `BigEndianBytes` encoding.
	BigEndianBytes *Type_Int64_Encoding_BigEndianBytes `json:"bigEndianBytes,omitempty"`
}

// +kcc:proto=google.bigtable.admin.v2.Type.Int64.Encoding.BigEndianBytes
type Type_Int64_Encoding_BigEndianBytes struct {
	// The underlying `Bytes` type, which may be able to encode further.
	BytesType *Type_Bytes `json:"bytesType,omitempty"`
}

// +kcc:proto=google.bigtable.admin.v2.Type.String
type Type_String struct {
	// The encoding to use when converting to/from lower level types.
	Encoding *Type_String_Encoding `json:"encoding,omitempty"`
}

// +kcc:proto=google.bigtable.admin.v2.Type.String.Encoding
type Type_String_Encoding struct {
	// Use `Utf8Raw` encoding.
	Utf8Raw *Type_String_Encoding_Utf8Raw `json:"utf8Raw,omitempty"`
}

// +kcc:proto=google.bigtable.admin.v2.Type.String.Encoding.Utf8Raw
type Type_String_Encoding_Utf8Raw struct {
}
