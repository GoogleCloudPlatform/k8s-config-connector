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

package v1alpha1

// +kcc:proto=google.firestore.admin.v1.Backup
type Backup struct {
	// Output only. The unique resource name of the Backup.
	//
	//  Format is `projects/{project}/locations/{location}/backups/{backup}`.
	Name *string `json:"name,omitempty"`

	// Output only. Name of the Firestore database that the backup is from.
	//
	//  Format is `projects/{project}/databases/{database}`.
	Database *string `json:"database,omitempty"`

	// Output only. The system-generated UUID4 for the Firestore database that the
	//  backup is from.
	DatabaseUid *string `json:"databaseUid,omitempty"`

	// Output only. The backup contains an externally consistent copy of the
	//  database at this time.
	SnapshotTime *string `json:"snapshotTime,omitempty"`

	// Output only. The timestamp at which this backup expires.
	ExpireTime *string `json:"expireTime,omitempty"`

	// Output only. Statistics about the backup.
	//
	//  This data only becomes available after the backup is fully materialized to
	//  secondary storage. This field will be empty till then.
	Stats *Backup_Stats `json:"stats,omitempty"`

	// Output only. The current state of the backup.
	State *string `json:"state,omitempty"`
}

// +kcc:proto=google.firestore.admin.v1.Backup.Stats
type Backup_Stats struct {
	// Output only. Summation of the size of all documents and index entries in
	//  the backup, measured in bytes.
	SizeBytes *int64 `json:"sizeBytes,omitempty"`

	// Output only. The total number of documents contained in the backup.
	DocumentCount *int64 `json:"documentCount,omitempty"`

	// Output only. The total number of index entries contained in the backup.
	IndexCount *int64 `json:"indexCount,omitempty"`
}

// +kcc:proto=google.firestore.admin.v1.BackupSchedule
type BackupSchedule struct {
	// Output only. The unique backup schedule identifier across all locations and
	//  databases for the given project.
	//
	//  This will be auto-assigned.
	//
	//  Format is
	//  `projects/{project}/databases/{database}/backupSchedules/{backup_schedule}`
	Name *string `json:"name,omitempty"`

	// Output only. The timestamp at which this backup schedule was created and
	//  effective since.
	//
	//  No backups will be created for this schedule before this time.
	CreateTime *string `json:"createTime,omitempty"`

	// Output only. The timestamp at which this backup schedule was most recently
	//  updated. When a backup schedule is first created, this is the same as
	//  create_time.
	UpdateTime *string `json:"updateTime,omitempty"`

	// At what relative time in the future, compared to its creation time,
	//  the backup should be deleted, e.g. keep backups for 7 days.
	Retention *string `json:"retention,omitempty"`

	// For a schedule that runs daily.
	DailyRecurrence *DailyRecurrence `json:"dailyRecurrence,omitempty"`

	// For a schedule that runs weekly on a specific day.
	WeeklyRecurrence *WeeklyRecurrence `json:"weeklyRecurrence,omitempty"`
}

// +kcc:proto=google.firestore.admin.v1.DailyRecurrence
type DailyRecurrence struct {
}

// +kcc:proto=google.firestore.admin.v1.Database
type Database struct {
	// The resource name of the Database.
	//  Format: `projects/{project}/databases/{database}`
	Name *string `json:"name,omitempty"`

	// Output only. The system-generated UUID4 for this Database.
	Uid *string `json:"uid,omitempty"`

	// Output only. The timestamp at which this database was created. Databases
	//  created before 2016 do not populate create_time.
	CreateTime *string `json:"createTime,omitempty"`

	// Output only. The timestamp at which this database was most recently
	//  updated. Note this only includes updates to the database resource and not
	//  data contained by the database.
	UpdateTime *string `json:"updateTime,omitempty"`

	// The location of the database. Available locations are listed at
	//  https://cloud.google.com/firestore/docs/locations.
	LocationID *string `json:"locationID,omitempty"`

	// The type of the database.
	//  See https://cloud.google.com/datastore/docs/firestore-or-datastore for
	//  information about how to choose.
	Type *string `json:"type,omitempty"`

	// The concurrency control mode to use for this database.
	ConcurrencyMode *string `json:"concurrencyMode,omitempty"`

	// Output only. The period during which past versions of data are retained in
	//  the database.
	//
	//  Any [read][google.firestore.v1.GetDocumentRequest.read_time]
	//  or [query][google.firestore.v1.ListDocumentsRequest.read_time] can specify
	//  a `read_time` within this window, and will read the state of the database
	//  at that time.
	//
	//  If the PITR feature is enabled, the retention period is 7 days. Otherwise,
	//  the retention period is 1 hour.
	VersionRetentionPeriod *string `json:"versionRetentionPeriod,omitempty"`

	// Output only. The earliest timestamp at which older versions of the data can
	//  be read from the database. See [version_retention_period] above; this field
	//  is populated with `now - version_retention_period`.
	//
	//  This value is continuously updated, and becomes stale the moment it is
	//  queried. If you are using this value to recover data, make sure to account
	//  for the time from the moment when the value is queried to the moment when
	//  you initiate the recovery.
	EarliestVersionTime *string `json:"earliestVersionTime,omitempty"`

	// Whether to enable the PITR feature on this database.
	PointInTimeRecoveryEnablement *string `json:"pointInTimeRecoveryEnablement,omitempty"`

	// The App Engine integration mode to use for this database.
	AppEngineIntegrationMode *string `json:"appEngineIntegrationMode,omitempty"`

	// Output only. The key_prefix for this database. This key_prefix is used, in
	//  combination with the project id ("<key prefix>~<project id>") to construct
	//  the application id that is returned from the Cloud Datastore APIs in Google
	//  App Engine first generation runtimes.
	//
	//  This value may be empty in which case the appid to use for URL-encoded keys
	//  is the project_id (eg: foo instead of v~foo).
	KeyPrefix *string `json:"keyPrefix,omitempty"`

	// State of delete protection for the database.
	DeleteProtectionState *string `json:"deleteProtectionState,omitempty"`

	// This checksum is computed by the server based on the value of other
	//  fields, and may be sent on update and delete requests to ensure the
	//  client has an up-to-date value before proceeding.
	Etag *string `json:"etag,omitempty"`
}

// +kcc:proto=google.firestore.admin.v1.Field
type Field struct {
	// Required. A field name of the form:
	//  `projects/{project_id}/databases/{database_id}/collectionGroups/{collection_id}/fields/{field_path}`
	//
	//  A field path can be a simple field name, e.g. `address` or a path to fields
	//  within `map_value` , e.g. `address.city`,
	//  or a special field path. The only valid special field is `*`, which
	//  represents any field.
	//
	//  Field paths can be quoted using `` ` `` (backtick). The only character that
	//  must be escaped within a quoted field path is the backtick character
	//  itself, escaped using a backslash. Special characters in field paths that
	//  must be quoted include: `*`, `.`,
	//  `` ` `` (backtick), `[`, `]`, as well as any ascii symbolic characters.
	//
	//  Examples:
	//  `` `address.city` `` represents a field named `address.city`, not the map
	//  key `city` in the field `address`. `` `*` `` represents a field named `*`,
	//  not any field.
	//
	//  A special `Field` contains the default indexing settings for all fields.
	//  This field's resource name is:
	//  `projects/{project_id}/databases/{database_id}/collectionGroups/__default__/fields/*`
	//  Indexes defined on this `Field` will be applied to all fields which do not
	//  have their own `Field` index configuration.
	Name *string `json:"name,omitempty"`

	// The index configuration for this field. If unset, field indexing will
	//  revert to the configuration defined by the `ancestor_field`. To
	//  explicitly remove all indexes for this field, specify an index config
	//  with an empty list of indexes.
	IndexConfig *Field_IndexConfig `json:"indexConfig,omitempty"`

	// The TTL configuration for this `Field`.
	//  Setting or unsetting this will enable or disable the TTL for
	//  documents that have this `Field`.
	TtlConfig *Field_TtlConfig `json:"ttlConfig,omitempty"`
}

// +kcc:proto=google.firestore.admin.v1.Field.IndexConfig
type Field_IndexConfig struct {
	// The indexes supported for this field.
	Indexes []Index `json:"indexes,omitempty"`

	// Output only. When true, the `Field`'s index configuration is set from the
	//  configuration specified by the `ancestor_field`.
	//  When false, the `Field`'s index configuration is defined explicitly.
	UsesAncestorConfig *bool `json:"usesAncestorConfig,omitempty"`

	// Output only. Specifies the resource name of the `Field` from which this
	//  field's index configuration is set (when `uses_ancestor_config` is true),
	//  or from which it *would* be set if this field had no index configuration
	//  (when `uses_ancestor_config` is false).
	AncestorField *string `json:"ancestorField,omitempty"`

	// Output only
	//  When true, the `Field`'s index configuration is in the process of being
	//  reverted. Once complete, the index config will transition to the same
	//  state as the field specified by `ancestor_field`, at which point
	//  `uses_ancestor_config` will be `true` and `reverting` will be `false`.
	Reverting *bool `json:"reverting,omitempty"`
}

// +kcc:proto=google.firestore.admin.v1.Field.TtlConfig
type Field_TtlConfig struct {
	// Output only. The state of the TTL configuration.
	State *string `json:"state,omitempty"`
}

// +kcc:proto=google.firestore.admin.v1.Index
type Index struct {
	// Output only. A server defined name for this index.
	//  The form of this name for composite indexes will be:
	//  `projects/{project_id}/databases/{database_id}/collectionGroups/{collection_id}/indexes/{composite_index_id}`
	//  For single field indexes, this field will be empty.
	Name *string `json:"name,omitempty"`

	// Indexes with a collection query scope specified allow queries
	//  against a collection that is the child of a specific document, specified at
	//  query time, and that has the same collection id.
	//
	//  Indexes with a collection group query scope specified allow queries against
	//  all collections descended from a specific document, specified at query
	//  time, and that have the same collection id as this index.
	QueryScope *string `json:"queryScope,omitempty"`

	// The API scope supported by this index.
	ApiScope *string `json:"apiScope,omitempty"`

	// The fields supported by this index.
	//
	//  For composite indexes, this requires a minimum of 2 and a maximum of 100
	//  fields. The last field entry is always for the field path `__name__`. If,
	//  on creation, `__name__` was not specified as the last field, it will be
	//  added automatically with the same direction as that of the last field
	//  defined. If the final field in a composite index is not directional, the
	//  `__name__` will be ordered ASCENDING (unless explicitly specified).
	//
	//  For single field indexes, this will always be exactly one entry with a
	//  field path equal to the field path of the associated field.
	Fields []Index_IndexField `json:"fields,omitempty"`

	// Output only. The serving state of the index.
	State *string `json:"state,omitempty"`
}

// +kcc:proto=google.firestore.admin.v1.Index.IndexField
type Index_IndexField struct {
	// Can be __name__.
	//  For single field indexes, this must match the name of the field or may
	//  be omitted.
	FieldPath *string `json:"fieldPath,omitempty"`

	// Indicates that this field supports ordering by the specified order or
	//  comparing using =, !=, <, <=, >, >=.
	Order *string `json:"order,omitempty"`

	// Indicates that this field supports operations on `array_value`s.
	ArrayConfig *string `json:"arrayConfig,omitempty"`

	// Indicates that this field supports nearest neighbor and distance
	//  operations on vector.
	VectorConfig *Index_IndexField_VectorConfig `json:"vectorConfig,omitempty"`
}

// +kcc:proto=google.firestore.admin.v1.Index.IndexField.VectorConfig
type Index_IndexField_VectorConfig struct {
	// Required. The vector dimension this configuration applies to.
	//
	//  The resulting index will only include vectors of this dimension, and
	//  can be used for vector search with the same dimension.
	Dimension *int32 `json:"dimension,omitempty"`

	// Indicates the vector index is a flat index.
	Flat *Index_IndexField_VectorConfig_FlatIndex `json:"flat,omitempty"`
}

// +kcc:proto=google.firestore.admin.v1.Index.IndexField.VectorConfig.FlatIndex
type Index_IndexField_VectorConfig_FlatIndex struct {
}

// +kcc:proto=google.firestore.admin.v1.Progress
type Progress struct {
	// The amount of work estimated.
	EstimatedWork *int64 `json:"estimatedWork,omitempty"`

	// The amount of work completed.
	CompletedWork *int64 `json:"completedWork,omitempty"`
}

// +kcc:proto=google.firestore.admin.v1.WeeklyRecurrence
type WeeklyRecurrence struct {
	// The day of week to run.
	//
	//  DAY_OF_WEEK_UNSPECIFIED is not allowed.
	Day *string `json:"day,omitempty"`
}
