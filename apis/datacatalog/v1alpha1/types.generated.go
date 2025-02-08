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


// +kcc:proto=google.cloud.datacatalog.v1beta1.BigQueryDateShardedSpec
type BigQueryDateShardedSpec struct {
}

// +kcc:proto=google.cloud.datacatalog.v1beta1.BigQueryTableSpec
type BigQueryTableSpec struct {

	// Table view specification. This field should only be populated if
	//  `table_source_type` is `BIGQUERY_VIEW`.
	// +kcc:proto:field=google.cloud.datacatalog.v1beta1.BigQueryTableSpec.view_spec
	ViewSpec *ViewSpec `json:"viewSpec,omitempty"`

	// Spec of a BigQuery table. This field should only be populated if
	//  `table_source_type` is `BIGQUERY_TABLE`.
	// +kcc:proto:field=google.cloud.datacatalog.v1beta1.BigQueryTableSpec.table_spec
	TableSpec *TableSpec `json:"tableSpec,omitempty"`
}

// +kcc:proto=google.cloud.datacatalog.v1beta1.ColumnSchema
type ColumnSchema struct {
	// Required. Name of the column.
	// +kcc:proto:field=google.cloud.datacatalog.v1beta1.ColumnSchema.column
	Column *string `json:"column,omitempty"`

	// Required. Type of the column.
	// +kcc:proto:field=google.cloud.datacatalog.v1beta1.ColumnSchema.type
	Type *string `json:"type,omitempty"`

	// Optional. Description of the column. Default value is an empty string.
	// +kcc:proto:field=google.cloud.datacatalog.v1beta1.ColumnSchema.description
	Description *string `json:"description,omitempty"`

	// Optional. A column's mode indicates whether the values in this column are
	//  required, nullable, etc. Only `NULLABLE`, `REQUIRED` and `REPEATED` are
	//  supported. Default mode is `NULLABLE`.
	// +kcc:proto:field=google.cloud.datacatalog.v1beta1.ColumnSchema.mode
	Mode *string `json:"mode,omitempty"`

	// Optional. Schema of sub-columns. A column can have zero or more
	//  sub-columns.
	// +kcc:proto:field=google.cloud.datacatalog.v1beta1.ColumnSchema.subcolumns
	Subcolumns []ColumnSchema `json:"subcolumns,omitempty"`
}

// +kcc:proto=google.cloud.datacatalog.v1beta1.Entry
type Entry struct {

	// The resource this metadata entry refers to.
	//
	//  For Google Cloud Platform resources, `linked_resource` is the [full name of
	//  the
	//  resource](https://cloud.google.com/apis/design/resource_names#full_resource_name).
	//  For example, the `linked_resource` for a table resource from BigQuery is:
	//
	//  * //bigquery.googleapis.com/projects/projectId/datasets/datasetId/tables/tableId
	//
	//  Output only when Entry is of type in the EntryType enum. For entries with
	//  user_specified_type, this field is optional and defaults to an empty
	//  string.
	// +kcc:proto:field=google.cloud.datacatalog.v1beta1.Entry.linked_resource
	LinkedResource *string `json:"linkedResource,omitempty"`

	// The type of the entry.
	//  Only used for Entries with types in the EntryType enum.
	// +kcc:proto:field=google.cloud.datacatalog.v1beta1.Entry.type
	Type *string `json:"type,omitempty"`

	// Entry type if it does not fit any of the input-allowed values listed in
	//  `EntryType` enum above. When creating an entry, users should check the
	//  enum values first, if nothing matches the entry to be created, then
	//  provide a custom value, for example "my_special_type".
	//  `user_specified_type` strings must begin with a letter or underscore and
	//  can only contain letters, numbers, and underscores; are case insensitive;
	//  must be at least 1 character and at most 64 characters long.
	//
	//  Currently, only FILESET enum value is allowed. All other entries created
	//  through Data Catalog must use `user_specified_type`.
	// +kcc:proto:field=google.cloud.datacatalog.v1beta1.Entry.user_specified_type
	UserSpecifiedType *string `json:"userSpecifiedType,omitempty"`

	// This field indicates the entry's source system that Data Catalog does not
	//  integrate with. `user_specified_system` strings must begin with a letter
	//  or underscore and can only contain letters, numbers, and underscores; are
	//  case insensitive; must be at least 1 character and at most 64 characters
	//  long.
	// +kcc:proto:field=google.cloud.datacatalog.v1beta1.Entry.user_specified_system
	UserSpecifiedSystem *string `json:"userSpecifiedSystem,omitempty"`

	// Specification that applies to a Cloud Storage fileset. This is only valid
	//  on entries of type FILESET.
	// +kcc:proto:field=google.cloud.datacatalog.v1beta1.Entry.gcs_fileset_spec
	GcsFilesetSpec *GcsFilesetSpec `json:"gcsFilesetSpec,omitempty"`

	// Specification that applies to a BigQuery table. This is only valid on
	//  entries of type `TABLE`.
	// +kcc:proto:field=google.cloud.datacatalog.v1beta1.Entry.bigquery_table_spec
	BigqueryTableSpec *BigQueryTableSpec `json:"bigqueryTableSpec,omitempty"`

	// Specification for a group of BigQuery tables with name pattern
	//  `[prefix]YYYYMMDD`. Context:
	//  https://cloud.google.com/bigquery/docs/partitioned-tables#partitioning_versus_sharding.
	// +kcc:proto:field=google.cloud.datacatalog.v1beta1.Entry.bigquery_date_sharded_spec
	BigqueryDateShardedSpec *BigQueryDateShardedSpec `json:"bigqueryDateShardedSpec,omitempty"`

	// Display information such as title and description. A short name to identify
	//  the entry, for example, "Analytics Data - Jan 2011". Default value is an
	//  empty string.
	// +kcc:proto:field=google.cloud.datacatalog.v1beta1.Entry.display_name
	DisplayName *string `json:"displayName,omitempty"`

	// Entry description, which can consist of several sentences or paragraphs
	//  that describe entry contents. Default value is an empty string.
	// +kcc:proto:field=google.cloud.datacatalog.v1beta1.Entry.description
	Description *string `json:"description,omitempty"`

	// Schema of the entry. An entry might not have any schema attached to it.
	// +kcc:proto:field=google.cloud.datacatalog.v1beta1.Entry.schema
	Schema *Schema `json:"schema,omitempty"`
}

// +kcc:proto=google.cloud.datacatalog.v1beta1.GcsFileSpec
type GcsFileSpec struct {
	// Required. The full file path. Example: `gs://bucket_name/a/b.txt`.
	// +kcc:proto:field=google.cloud.datacatalog.v1beta1.GcsFileSpec.file_path
	FilePath *string `json:"filePath,omitempty"`
}

// +kcc:proto=google.cloud.datacatalog.v1beta1.GcsFilesetSpec
type GcsFilesetSpec struct {
	// Required. Patterns to identify a set of files in Google Cloud Storage.
	//  See [Cloud Storage
	//  documentation](https://cloud.google.com/storage/docs/gsutil/addlhelp/WildcardNames)
	//  for more information. Note that bucket wildcards are currently not
	//  supported.
	//
	//  Examples of valid file_patterns:
	//
	//   * `gs://bucket_name/dir/*`: matches all files within `bucket_name/dir`
	//                               directory.
	//   * `gs://bucket_name/dir/**`: matches all files in `bucket_name/dir`
	//                                spanning all subdirectories.
	//   * `gs://bucket_name/file*`: matches files prefixed by `file` in
	//                               `bucket_name`
	//   * `gs://bucket_name/??.txt`: matches files with two characters followed by
	//                                `.txt` in `bucket_name`
	//   * `gs://bucket_name/[aeiou].txt`: matches files that contain a single
	//                                     vowel character followed by `.txt` in
	//                                     `bucket_name`
	//   * `gs://bucket_name/[a-m].txt`: matches files that contain `a`, `b`, ...
	//                                   or `m` followed by `.txt` in `bucket_name`
	//   * `gs://bucket_name/a/*/b`: matches all files in `bucket_name` that match
	//                               `a/*/b` pattern, such as `a/c/b`, `a/d/b`
	//   * `gs://another_bucket/a.txt`: matches `gs://another_bucket/a.txt`
	//
	//  You can combine wildcards to provide more powerful matches, for example:
	//
	//   * `gs://bucket_name/[a-m]??.j*g`
	// +kcc:proto:field=google.cloud.datacatalog.v1beta1.GcsFilesetSpec.file_patterns
	FilePatterns []string `json:"filePatterns,omitempty"`
}

// +kcc:proto=google.cloud.datacatalog.v1beta1.Schema
type Schema struct {
	// Required. Schema of columns. A maximum of 10,000 columns and sub-columns
	//  can be specified.
	// +kcc:proto:field=google.cloud.datacatalog.v1beta1.Schema.columns
	Columns []ColumnSchema `json:"columns,omitempty"`
}

// +kcc:proto=google.cloud.datacatalog.v1beta1.SystemTimestamps
type SystemTimestamps struct {
	// The creation time of the resource within the given system.
	// +kcc:proto:field=google.cloud.datacatalog.v1beta1.SystemTimestamps.create_time
	CreateTime *string `json:"createTime,omitempty"`

	// The last-modified time of the resource within the given system.
	// +kcc:proto:field=google.cloud.datacatalog.v1beta1.SystemTimestamps.update_time
	UpdateTime *string `json:"updateTime,omitempty"`
}

// +kcc:proto=google.cloud.datacatalog.v1beta1.TableSpec
type TableSpec struct {
}

// +kcc:proto=google.cloud.datacatalog.v1beta1.UsageSignal
type UsageSignal struct {
	// The timestamp of the end of the usage statistics duration.
	// +kcc:proto:field=google.cloud.datacatalog.v1beta1.UsageSignal.update_time
	UpdateTime *string `json:"updateTime,omitempty"`

	// TODO: unsupported map type with key string and value message

}

// +kcc:proto=google.cloud.datacatalog.v1beta1.UsageStats
type UsageStats struct {
	// The number of times that the underlying entry was successfully used.
	// +kcc:proto:field=google.cloud.datacatalog.v1beta1.UsageStats.total_completions
	TotalCompletions *float32 `json:"totalCompletions,omitempty"`

	// The number of times that the underlying entry was attempted to be used
	//  but failed.
	// +kcc:proto:field=google.cloud.datacatalog.v1beta1.UsageStats.total_failures
	TotalFailures *float32 `json:"totalFailures,omitempty"`

	// The number of times that the underlying entry was attempted to be used
	//  but was cancelled by the user.
	// +kcc:proto:field=google.cloud.datacatalog.v1beta1.UsageStats.total_cancellations
	TotalCancellations *float32 `json:"totalCancellations,omitempty"`

	// Total time spent (in milliseconds) during uses the resulted in completions.
	// +kcc:proto:field=google.cloud.datacatalog.v1beta1.UsageStats.total_execution_time_for_completions_millis
	TotalExecutionTimeForCompletionsMillis *float32 `json:"totalExecutionTimeForCompletionsMillis,omitempty"`
}

// +kcc:proto=google.cloud.datacatalog.v1beta1.ViewSpec
type ViewSpec struct {
}

// +kcc:proto=google.cloud.datacatalog.v1beta1.BigQueryDateShardedSpec
type BigQueryDateShardedSpecObservedState struct {
	// Output only. The Data Catalog resource name of the dataset entry the
	//  current table belongs to, for example,
	//  `projects/{project_id}/locations/{location}/entrygroups/{entry_group_id}/entries/{entry_id}`.
	// +kcc:proto:field=google.cloud.datacatalog.v1beta1.BigQueryDateShardedSpec.dataset
	Dataset *string `json:"dataset,omitempty"`

	// Output only. The table name prefix of the shards. The name of any given
	//  shard is
	//  `[table_prefix]YYYYMMDD`, for example, for shard `MyTable20180101`, the
	//  `table_prefix` is `MyTable`.
	// +kcc:proto:field=google.cloud.datacatalog.v1beta1.BigQueryDateShardedSpec.table_prefix
	TablePrefix *string `json:"tablePrefix,omitempty"`

	// Output only. Total number of shards.
	// +kcc:proto:field=google.cloud.datacatalog.v1beta1.BigQueryDateShardedSpec.shard_count
	ShardCount *int64 `json:"shardCount,omitempty"`
}

// +kcc:proto=google.cloud.datacatalog.v1beta1.BigQueryTableSpec
type BigQueryTableSpecObservedState struct {
	// Output only. The table source type.
	// +kcc:proto:field=google.cloud.datacatalog.v1beta1.BigQueryTableSpec.table_source_type
	TableSourceType *string `json:"tableSourceType,omitempty"`

	// Table view specification. This field should only be populated if
	//  `table_source_type` is `BIGQUERY_VIEW`.
	// +kcc:proto:field=google.cloud.datacatalog.v1beta1.BigQueryTableSpec.view_spec
	ViewSpec *ViewSpecObservedState `json:"viewSpec,omitempty"`

	// Spec of a BigQuery table. This field should only be populated if
	//  `table_source_type` is `BIGQUERY_TABLE`.
	// +kcc:proto:field=google.cloud.datacatalog.v1beta1.BigQueryTableSpec.table_spec
	TableSpec *TableSpecObservedState `json:"tableSpec,omitempty"`
}

// +kcc:proto=google.cloud.datacatalog.v1beta1.Entry
type EntryObservedState struct {
	// Output only. Identifier. The Data Catalog resource name of the entry in URL
	//  format. Example:
	//
	//  * projects/{project_id}/locations/{location}/entryGroups/{entry_group_id}/entries/{entry_id}
	//
	//  Note that this Entry and its child resources may not actually be stored in
	//  the location in this name.
	// +kcc:proto:field=google.cloud.datacatalog.v1beta1.Entry.name
	Name *string `json:"name,omitempty"`

	// Output only. This field indicates the entry's source system that Data
	//  Catalog integrates with, such as BigQuery or Pub/Sub.
	// +kcc:proto:field=google.cloud.datacatalog.v1beta1.Entry.integrated_system
	IntegratedSystem *string `json:"integratedSystem,omitempty"`

	// Specification that applies to a Cloud Storage fileset. This is only valid
	//  on entries of type FILESET.
	// +kcc:proto:field=google.cloud.datacatalog.v1beta1.Entry.gcs_fileset_spec
	GcsFilesetSpec *GcsFilesetSpecObservedState `json:"gcsFilesetSpec,omitempty"`

	// Specification that applies to a BigQuery table. This is only valid on
	//  entries of type `TABLE`.
	// +kcc:proto:field=google.cloud.datacatalog.v1beta1.Entry.bigquery_table_spec
	BigqueryTableSpec *BigQueryTableSpecObservedState `json:"bigqueryTableSpec,omitempty"`

	// Specification for a group of BigQuery tables with name pattern
	//  `[prefix]YYYYMMDD`. Context:
	//  https://cloud.google.com/bigquery/docs/partitioned-tables#partitioning_versus_sharding.
	// +kcc:proto:field=google.cloud.datacatalog.v1beta1.Entry.bigquery_date_sharded_spec
	BigqueryDateShardedSpec *BigQueryDateShardedSpecObservedState `json:"bigqueryDateShardedSpec,omitempty"`

	// Output only. Timestamps about the underlying resource, not about this Data
	//  Catalog entry. Output only when Entry is of type in the EntryType enum. For
	//  entries with user_specified_type, this field is optional and defaults to an
	//  empty timestamp.
	// +kcc:proto:field=google.cloud.datacatalog.v1beta1.Entry.source_system_timestamps
	SourceSystemTimestamps *SystemTimestamps `json:"sourceSystemTimestamps,omitempty"`

	// Output only. Statistics on the usage level of the resource.
	// +kcc:proto:field=google.cloud.datacatalog.v1beta1.Entry.usage_signal
	UsageSignal *UsageSignal `json:"usageSignal,omitempty"`
}

// +kcc:proto=google.cloud.datacatalog.v1beta1.GcsFileSpec
type GcsFileSpecObservedState struct {
	// Output only. Timestamps about the Cloud Storage file.
	// +kcc:proto:field=google.cloud.datacatalog.v1beta1.GcsFileSpec.gcs_timestamps
	GcsTimestamps *SystemTimestamps `json:"gcsTimestamps,omitempty"`

	// Output only. The size of the file, in bytes.
	// +kcc:proto:field=google.cloud.datacatalog.v1beta1.GcsFileSpec.size_bytes
	SizeBytes *int64 `json:"sizeBytes,omitempty"`
}

// +kcc:proto=google.cloud.datacatalog.v1beta1.GcsFilesetSpec
type GcsFilesetSpecObservedState struct {
	// Output only. Sample files contained in this fileset, not all files
	//  contained in this fileset are represented here.
	// +kcc:proto:field=google.cloud.datacatalog.v1beta1.GcsFilesetSpec.sample_gcs_file_specs
	SampleGcsFileSpecs []GcsFileSpec `json:"sampleGcsFileSpecs,omitempty"`
}

// +kcc:proto=google.cloud.datacatalog.v1beta1.SystemTimestamps
type SystemTimestampsObservedState struct {
	// Output only. The expiration time of the resource within the given system.
	//  Currently only apllicable to BigQuery resources.
	// +kcc:proto:field=google.cloud.datacatalog.v1beta1.SystemTimestamps.expire_time
	ExpireTime *string `json:"expireTime,omitempty"`
}

// +kcc:proto=google.cloud.datacatalog.v1beta1.TableSpec
type TableSpecObservedState struct {
	// Output only. If the table is a dated shard, i.e., with name pattern
	//  `[prefix]YYYYMMDD`, `grouped_entry` is the Data Catalog resource name of
	//  the date sharded grouped entry, for example,
	//  `projects/{project_id}/locations/{location}/entrygroups/{entry_group_id}/entries/{entry_id}`.
	//  Otherwise, `grouped_entry` is empty.
	// +kcc:proto:field=google.cloud.datacatalog.v1beta1.TableSpec.grouped_entry
	GroupedEntry *string `json:"groupedEntry,omitempty"`
}

// +kcc:proto=google.cloud.datacatalog.v1beta1.ViewSpec
type ViewSpecObservedState struct {
	// Output only. The query that defines the table view.
	// +kcc:proto:field=google.cloud.datacatalog.v1beta1.ViewSpec.view_query
	ViewQuery *string `json:"viewQuery,omitempty"`
}
