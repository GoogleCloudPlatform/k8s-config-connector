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

// +kcc:proto=google.logging.v2.BigQueryOptions
type BigQueryOptions struct {
	// Optional. Whether to use [BigQuery's partition
	//  tables](https://cloud.google.com/bigquery/docs/partitioned-tables). By
	//  default, Cloud Logging creates dated tables based on the log entries'
	//  timestamps, e.g. syslog_20170523. With partitioned tables the date suffix
	//  is no longer present and [special query
	//  syntax](https://cloud.google.com/bigquery/docs/querying-partitioned-tables)
	//  has to be used instead. In both cases, tables are sharded based on UTC
	//  timezone.
	// +kcc:proto:field=google.logging.v2.BigQueryOptions.use_partitioned_tables
	UsePartitionedTables *bool `json:"usePartitionedTables,omitempty"`
}

// +kcc:proto=google.logging.v2.LogExclusion
type LogExclusion struct {
	// Required. A client-assigned identifier, such as
	//  `"load-balancer-exclusion"`. Identifiers are limited to 100 characters and
	//  can include only letters, digits, underscores, hyphens, and periods. First
	//  character has to be alphanumeric.
	// +kcc:proto:field=google.logging.v2.LogExclusion.name
	Name *string `json:"name,omitempty"`

	// Optional. A description of this exclusion.
	// +kcc:proto:field=google.logging.v2.LogExclusion.description
	Description *string `json:"description,omitempty"`

	// Required. An [advanced logs
	//  filter](https://cloud.google.com/logging/docs/view/advanced-queries) that
	//  matches the log entries to be excluded. By using the [sample
	//  function](https://cloud.google.com/logging/docs/view/advanced-queries#sample),
	//  you can exclude less than 100% of the matching log entries.
	//
	//  For example, the following query matches 99% of low-severity log entries
	//  from Google Cloud Storage buckets:
	//
	//    `resource.type=gcs_bucket severity<ERROR sample(insertId, 0.99)`
	// +kcc:proto:field=google.logging.v2.LogExclusion.filter
	Filter *string `json:"filter,omitempty"`

	// Optional. If set to True, then this exclusion is disabled and it does not
	//  exclude any log entries. You can [update an
	//  exclusion][google.logging.v2.ConfigServiceV2.UpdateExclusion] to change the
	//  value of this field.
	// +kcc:proto:field=google.logging.v2.LogExclusion.disabled
	Disabled *bool `json:"disabled,omitempty"`
}

// +kcc:proto=google.logging.v2.LogSink
type LogSink struct {
	// Required. The client-assigned sink identifier, unique within the project.
	//
	//  For example: `"my-syslog-errors-to-pubsub"`. Sink identifiers are limited
	//  to 100 characters and can include only the following characters: upper and
	//  lower-case alphanumeric characters, underscores, hyphens, and periods.
	//  First character has to be alphanumeric.
	// +kcc:proto:field=google.logging.v2.LogSink.name
	Name *string `json:"name,omitempty"`

	// Required. The export destination:
	//
	//      "storage.googleapis.com/[GCS_BUCKET]"
	//      "bigquery.googleapis.com/projects/[PROJECT_ID]/datasets/[DATASET]"
	//      "pubsub.googleapis.com/projects/[PROJECT_ID]/topics/[TOPIC_ID]"
	//
	//  The sink's `writer_identity`, set when the sink is created, must have
	//  permission to write to the destination or else the log entries are not
	//  exported. For more information, see
	//  [Exporting Logs with
	//  Sinks](https://cloud.google.com/logging/docs/api/tasks/exporting-logs).
	// +kcc:proto:field=google.logging.v2.LogSink.destination
	Destination *string `json:"destination,omitempty"`

	// Optional. An [advanced logs
	//  filter](https://cloud.google.com/logging/docs/view/advanced-queries). The
	//  only exported log entries are those that are in the resource owning the
	//  sink and that match the filter.
	//
	//  For example:
	//
	//    `logName="projects/[PROJECT_ID]/logs/[LOG_ID]" AND severity>=ERROR`
	// +kcc:proto:field=google.logging.v2.LogSink.filter
	Filter *string `json:"filter,omitempty"`

	// Optional. A description of this sink.
	//
	//  The maximum length of the description is 8000 characters.
	// +kcc:proto:field=google.logging.v2.LogSink.description
	Description *string `json:"description,omitempty"`

	// Optional. If set to true, then this sink is disabled and it does not export
	//  any log entries.
	// +kcc:proto:field=google.logging.v2.LogSink.disabled
	Disabled *bool `json:"disabled,omitempty"`

	// Optional. Log entries that match any of these exclusion filters will not be
	//  exported.
	//
	//  If a log entry is matched by both `filter` and one of `exclusion_filters`
	//  it will not be exported.
	// +kcc:proto:field=google.logging.v2.LogSink.exclusions
	Exclusions []LogExclusion `json:"exclusions,omitempty"`

	// Deprecated. This field is unused.
	// +kcc:proto:field=google.logging.v2.LogSink.output_version_format
	OutputVersionFormat *string `json:"outputVersionFormat,omitempty"`

	// Optional. This field applies only to sinks owned by organizations and
	//  folders. If the field is false, the default, only the logs owned by the
	//  sink's parent resource are available for export. If the field is true, then
	//  log entries from all the projects, folders, and billing accounts contained
	//  in the sink's parent resource are also available for export. Whether a
	//  particular log entry from the children is exported depends on the sink's
	//  filter expression.
	//
	//  For example, if this field is true, then the filter
	//  `resource.type=gce_instance` would export all Compute Engine VM instance
	//  log entries from all projects in the sink's parent.
	//
	//  To only export entries from certain child projects, filter on the project
	//  part of the log name:
	//
	//    logName:("projects/test-project1/" OR "projects/test-project2/") AND
	//    resource.type=gce_instance
	// +kcc:proto:field=google.logging.v2.LogSink.include_children
	IncludeChildren *bool `json:"includeChildren,omitempty"`

	// Optional. Options that affect sinks exporting data to BigQuery.
	// +kcc:proto:field=google.logging.v2.LogSink.bigquery_options
	BigqueryOptions *BigQueryOptions `json:"bigqueryOptions,omitempty"`
}

// +kcc:proto=google.logging.v2.BigQueryOptions
type BigQueryOptionsObservedState struct {
	// Output only. True if new timestamp column based partitioning is in use,
	//  false if legacy ingestion-time partitioning is in use.
	//
	//  All new sinks will have this field set true and will use timestamp column
	//  based partitioning. If use_partitioned_tables is false, this value has no
	//  meaning and will be false. Legacy sinks using partitioned tables will have
	//  this field set to false.
	// +kcc:proto:field=google.logging.v2.BigQueryOptions.uses_timestamp_column_partitioning
	UsesTimestampColumnPartitioning *bool `json:"usesTimestampColumnPartitioning,omitempty"`
}

// +kcc:proto=google.logging.v2.LogExclusion
type LogExclusionObservedState struct {
	// Output only. The creation timestamp of the exclusion.
	//
	//  This field may not be present for older exclusions.
	// +kcc:proto:field=google.logging.v2.LogExclusion.create_time
	CreateTime *string `json:"createTime,omitempty"`

	// Output only. The last update timestamp of the exclusion.
	//
	//  This field may not be present for older exclusions.
	// +kcc:proto:field=google.logging.v2.LogExclusion.update_time
	UpdateTime *string `json:"updateTime,omitempty"`
}

// +kcc:proto=google.logging.v2.LogSink
type LogSinkObservedState struct {
	// Optional. Log entries that match any of these exclusion filters will not be
	//  exported.
	//
	//  If a log entry is matched by both `filter` and one of `exclusion_filters`
	//  it will not be exported.
	// +kcc:proto:field=google.logging.v2.LogSink.exclusions
	Exclusions []LogExclusionObservedState `json:"exclusions,omitempty"`

	// Output only. An IAM identity&mdash;a service account or group&mdash;under
	//  which Cloud Logging writes the exported log entries to the sink's
	//  destination. This field is either set by specifying
	//  `custom_writer_identity` or set automatically by
	//  [sinks.create][google.logging.v2.ConfigServiceV2.CreateSink] and
	//  [sinks.update][google.logging.v2.ConfigServiceV2.UpdateSink] based on the
	//  value of `unique_writer_identity` in those methods.
	//
	//  Until you grant this identity write-access to the destination, log entry
	//  exports from this sink will fail. For more information, see [Granting
	//  Access for a
	//  Resource](https://cloud.google.com/iam/docs/granting-roles-to-service-accounts#granting_access_to_a_service_account_for_a_resource).
	//  Consult the destination service's documentation to determine the
	//  appropriate IAM roles to assign to the identity.
	//
	//  Sinks that have a destination that is a log bucket in the same project as
	//  the sink cannot have a writer_identity and no additional permissions are
	//  required.
	// +kcc:proto:field=google.logging.v2.LogSink.writer_identity
	WriterIdentity *string `json:"writerIdentity,omitempty"`

	// Optional. Options that affect sinks exporting data to BigQuery.
	// +kcc:proto:field=google.logging.v2.LogSink.bigquery_options
	BigqueryOptions *BigQueryOptionsObservedState `json:"bigqueryOptions,omitempty"`

	// Output only. The creation timestamp of the sink.
	//
	//  This field may not be present for older sinks.
	// +kcc:proto:field=google.logging.v2.LogSink.create_time
	CreateTime *string `json:"createTime,omitempty"`

	// Output only. The last update timestamp of the sink.
	//
	//  This field may not be present for older sinks.
	// +kcc:proto:field=google.logging.v2.LogSink.update_time
	UpdateTime *string `json:"updateTime,omitempty"`
}
