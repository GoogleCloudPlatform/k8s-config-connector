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


// +kcc:proto=google.api.Distribution
type Distribution struct {
	// The number of values in the population. Must be non-negative. This value
	//  must equal the sum of the values in `bucket_counts` if a histogram is
	//  provided.
	// +kcc:proto:field=google.api.Distribution.count
	Count *int64 `json:"count,omitempty"`

	// The arithmetic mean of the values in the population. If `count` is zero
	//  then this field must be zero.
	// +kcc:proto:field=google.api.Distribution.mean
	Mean *float64 `json:"mean,omitempty"`

	// The sum of squared deviations from the mean of the values in the
	//  population. For values x_i this is:
	//
	//      Sum[i=1..n]((x_i - mean)^2)
	//
	//  Knuth, "The Art of Computer Programming", Vol. 2, page 232, 3rd edition
	//  describes Welford's method for accumulating this sum in one pass.
	//
	//  If `count` is zero then this field must be zero.
	// +kcc:proto:field=google.api.Distribution.sum_of_squared_deviation
	SumOfSquaredDeviation *float64 `json:"sumOfSquaredDeviation,omitempty"`

	// If specified, contains the range of the population values. The field
	//  must not be present if the `count` is zero.
	// +kcc:proto:field=google.api.Distribution.range
	Range *Distribution_Range `json:"range,omitempty"`

	// Defines the histogram bucket boundaries. If the distribution does not
	//  contain a histogram, then omit this field.
	// +kcc:proto:field=google.api.Distribution.bucket_options
	BucketOptions *Distribution_BucketOptions `json:"bucketOptions,omitempty"`

	// The number of values in each bucket of the histogram, as described in
	//  `bucket_options`. If the distribution does not have a histogram, then omit
	//  this field. If there is a histogram, then the sum of the values in
	//  `bucket_counts` must equal the value in the `count` field of the
	//  distribution.
	//
	//  If present, `bucket_counts` should contain N values, where N is the number
	//  of buckets specified in `bucket_options`. If you supply fewer than N
	//  values, the remaining values are assumed to be 0.
	//
	//  The order of the values in `bucket_counts` follows the bucket numbering
	//  schemes described for the three bucket types. The first value must be the
	//  count for the underflow bucket (number 0). The next N-2 values are the
	//  counts for the finite buckets (number 1 through N-2). The N'th value in
	//  `bucket_counts` is the count for the overflow bucket (number N-1).
	// +kcc:proto:field=google.api.Distribution.bucket_counts
	BucketCounts []int64 `json:"bucketCounts,omitempty"`

	// Must be in increasing order of `value` field.
	// +kcc:proto:field=google.api.Distribution.exemplars
	Exemplars []Distribution_Exemplar `json:"exemplars,omitempty"`
}

// +kcc:proto=google.api.Distribution.BucketOptions
type Distribution_BucketOptions struct {
	// The linear bucket.
	// +kcc:proto:field=google.api.Distribution.BucketOptions.linear_buckets
	LinearBuckets *Distribution_BucketOptions_Linear `json:"linearBuckets,omitempty"`

	// The exponential buckets.
	// +kcc:proto:field=google.api.Distribution.BucketOptions.exponential_buckets
	ExponentialBuckets *Distribution_BucketOptions_Exponential `json:"exponentialBuckets,omitempty"`

	// The explicit buckets.
	// +kcc:proto:field=google.api.Distribution.BucketOptions.explicit_buckets
	ExplicitBuckets *Distribution_BucketOptions_Explicit `json:"explicitBuckets,omitempty"`
}

// +kcc:proto=google.api.Distribution.BucketOptions.Explicit
type Distribution_BucketOptions_Explicit struct {
	// The values must be monotonically increasing.
	// +kcc:proto:field=google.api.Distribution.BucketOptions.Explicit.bounds
	Bounds []float64 `json:"bounds,omitempty"`
}

// +kcc:proto=google.api.Distribution.BucketOptions.Exponential
type Distribution_BucketOptions_Exponential struct {
	// Must be greater than 0.
	// +kcc:proto:field=google.api.Distribution.BucketOptions.Exponential.num_finite_buckets
	NumFiniteBuckets *int32 `json:"numFiniteBuckets,omitempty"`

	// Must be greater than 1.
	// +kcc:proto:field=google.api.Distribution.BucketOptions.Exponential.growth_factor
	GrowthFactor *float64 `json:"growthFactor,omitempty"`

	// Must be greater than 0.
	// +kcc:proto:field=google.api.Distribution.BucketOptions.Exponential.scale
	Scale *float64 `json:"scale,omitempty"`
}

// +kcc:proto=google.api.Distribution.BucketOptions.Linear
type Distribution_BucketOptions_Linear struct {
	// Must be greater than 0.
	// +kcc:proto:field=google.api.Distribution.BucketOptions.Linear.num_finite_buckets
	NumFiniteBuckets *int32 `json:"numFiniteBuckets,omitempty"`

	// Must be greater than 0.
	// +kcc:proto:field=google.api.Distribution.BucketOptions.Linear.width
	Width *float64 `json:"width,omitempty"`

	// Lower bound of the first bucket.
	// +kcc:proto:field=google.api.Distribution.BucketOptions.Linear.offset
	Offset *float64 `json:"offset,omitempty"`
}

// +kcc:proto=google.api.Distribution.Exemplar
type Distribution_Exemplar struct {
	// Value of the exemplar point. This value determines to which bucket the
	//  exemplar belongs.
	// +kcc:proto:field=google.api.Distribution.Exemplar.value
	Value *float64 `json:"value,omitempty"`

	// The observation (sampling) time of the above value.
	// +kcc:proto:field=google.api.Distribution.Exemplar.timestamp
	Timestamp *string `json:"timestamp,omitempty"`

	// Contextual information about the example value. Examples are:
	//
	//    Trace: type.googleapis.com/google.monitoring.v3.SpanContext
	//
	//    Literal string: type.googleapis.com/google.protobuf.StringValue
	//
	//    Labels dropped during aggregation:
	//      type.googleapis.com/google.monitoring.v3.DroppedLabels
	//
	//  There may be only a single attachment of any given message type in a
	//  single exemplar, and this is enforced by the system.
	// +kcc:proto:field=google.api.Distribution.Exemplar.attachments
	Attachments []Any `json:"attachments,omitempty"`
}

// +kcc:proto=google.api.Distribution.Range
type Distribution_Range struct {
	// The minimum of the population values.
	// +kcc:proto:field=google.api.Distribution.Range.min
	Min *float64 `json:"min,omitempty"`

	// The maximum of the population values.
	// +kcc:proto:field=google.api.Distribution.Range.max
	Max *float64 `json:"max,omitempty"`
}

// +kcc:proto=google.cloud.bigquery.migration.v2.AzureSynapseDialect
type AzureSynapseDialect struct {
}

// +kcc:proto=google.cloud.bigquery.migration.v2.BigQueryDialect
type BigQueryDialect struct {
}

// +kcc:proto=google.cloud.bigquery.migration.v2.DB2Dialect
type DB2Dialect struct {
}

// +kcc:proto=google.cloud.bigquery.migration.v2.Dialect
type Dialect struct {
	// The BigQuery dialect
	// +kcc:proto:field=google.cloud.bigquery.migration.v2.Dialect.bigquery_dialect
	BigqueryDialect *BigQueryDialect `json:"bigqueryDialect,omitempty"`

	// The HiveQL dialect
	// +kcc:proto:field=google.cloud.bigquery.migration.v2.Dialect.hiveql_dialect
	HiveqlDialect *HiveQLDialect `json:"hiveqlDialect,omitempty"`

	// The Redshift dialect
	// +kcc:proto:field=google.cloud.bigquery.migration.v2.Dialect.redshift_dialect
	RedshiftDialect *RedshiftDialect `json:"redshiftDialect,omitempty"`

	// The Teradata dialect
	// +kcc:proto:field=google.cloud.bigquery.migration.v2.Dialect.teradata_dialect
	TeradataDialect *TeradataDialect `json:"teradataDialect,omitempty"`

	// The Oracle dialect
	// +kcc:proto:field=google.cloud.bigquery.migration.v2.Dialect.oracle_dialect
	OracleDialect *OracleDialect `json:"oracleDialect,omitempty"`

	// The SparkSQL dialect
	// +kcc:proto:field=google.cloud.bigquery.migration.v2.Dialect.sparksql_dialect
	SparksqlDialect *SparkSQLDialect `json:"sparksqlDialect,omitempty"`

	// The Snowflake dialect
	// +kcc:proto:field=google.cloud.bigquery.migration.v2.Dialect.snowflake_dialect
	SnowflakeDialect *SnowflakeDialect `json:"snowflakeDialect,omitempty"`

	// The Netezza dialect
	// +kcc:proto:field=google.cloud.bigquery.migration.v2.Dialect.netezza_dialect
	NetezzaDialect *NetezzaDialect `json:"netezzaDialect,omitempty"`

	// The Azure Synapse dialect
	// +kcc:proto:field=google.cloud.bigquery.migration.v2.Dialect.azure_synapse_dialect
	AzureSynapseDialect *AzureSynapseDialect `json:"azureSynapseDialect,omitempty"`

	// The Vertica dialect
	// +kcc:proto:field=google.cloud.bigquery.migration.v2.Dialect.vertica_dialect
	VerticaDialect *VerticaDialect `json:"verticaDialect,omitempty"`

	// The SQL Server dialect
	// +kcc:proto:field=google.cloud.bigquery.migration.v2.Dialect.sql_server_dialect
	SqlServerDialect *SQLServerDialect `json:"sqlServerDialect,omitempty"`

	// The Postgresql dialect
	// +kcc:proto:field=google.cloud.bigquery.migration.v2.Dialect.postgresql_dialect
	PostgresqlDialect *PostgresqlDialect `json:"postgresqlDialect,omitempty"`

	// The Presto dialect
	// +kcc:proto:field=google.cloud.bigquery.migration.v2.Dialect.presto_dialect
	PrestoDialect *PrestoDialect `json:"prestoDialect,omitempty"`

	// The MySQL dialect
	// +kcc:proto:field=google.cloud.bigquery.migration.v2.Dialect.mysql_dialect
	MysqlDialect *MySQLDialect `json:"mysqlDialect,omitempty"`

	// DB2 dialect
	// +kcc:proto:field=google.cloud.bigquery.migration.v2.Dialect.db2_dialect
	Db2Dialect *DB2Dialect `json:"db2Dialect,omitempty"`

	// SQLite dialect
	// +kcc:proto:field=google.cloud.bigquery.migration.v2.Dialect.sqlite_dialect
	SqliteDialect *SQLiteDialect `json:"sqliteDialect,omitempty"`

	// Greenplum dialect
	// +kcc:proto:field=google.cloud.bigquery.migration.v2.Dialect.greenplum_dialect
	GreenplumDialect *GreenplumDialect `json:"greenplumDialect,omitempty"`
}

// +kcc:proto=google.cloud.bigquery.migration.v2.ErrorDetail
type ErrorDetail struct {
	// Optional. The exact location within the resource (if applicable).
	// +kcc:proto:field=google.cloud.bigquery.migration.v2.ErrorDetail.location
	Location *ErrorLocation `json:"location,omitempty"`

	// Required. Describes the cause of the error with structured detail.
	// +kcc:proto:field=google.cloud.bigquery.migration.v2.ErrorDetail.error_info
	ErrorInfo *ErrorInfo `json:"errorInfo,omitempty"`
}

// +kcc:proto=google.cloud.bigquery.migration.v2.ErrorLocation
type ErrorLocation struct {
	// Optional. If applicable, denotes the line where the error occurred. A zero
	//  value means that there is no line information.
	// +kcc:proto:field=google.cloud.bigquery.migration.v2.ErrorLocation.line
	Line *int32 `json:"line,omitempty"`

	// Optional. If applicable, denotes the column where the error occurred. A
	//  zero value means that there is no columns information.
	// +kcc:proto:field=google.cloud.bigquery.migration.v2.ErrorLocation.column
	Column *int32 `json:"column,omitempty"`
}

// +kcc:proto=google.cloud.bigquery.migration.v2.GcsReportLogMessage
type GcsReportLogMessage struct {
	// Severity of the translation record.
	// +kcc:proto:field=google.cloud.bigquery.migration.v2.GcsReportLogMessage.severity
	Severity *string `json:"severity,omitempty"`

	// Category of the error/warning. Example: SyntaxError
	// +kcc:proto:field=google.cloud.bigquery.migration.v2.GcsReportLogMessage.category
	Category *string `json:"category,omitempty"`

	// The file path in which the error occurred
	// +kcc:proto:field=google.cloud.bigquery.migration.v2.GcsReportLogMessage.file_path
	FilePath *string `json:"filePath,omitempty"`

	// The file name in which the error occurred
	// +kcc:proto:field=google.cloud.bigquery.migration.v2.GcsReportLogMessage.filename
	Filename *string `json:"filename,omitempty"`

	// Specifies the row from the source text where the error occurred (0 based,
	//  -1 for messages without line location). Example: 2
	// +kcc:proto:field=google.cloud.bigquery.migration.v2.GcsReportLogMessage.source_script_line
	SourceScriptLine *int32 `json:"sourceScriptLine,omitempty"`

	// Specifies the column from the source texts where the error occurred. (0
	//  based, -1 for messages without column location) example: 6
	// +kcc:proto:field=google.cloud.bigquery.migration.v2.GcsReportLogMessage.source_script_column
	SourceScriptColumn *int32 `json:"sourceScriptColumn,omitempty"`

	// Detailed message of the record.
	// +kcc:proto:field=google.cloud.bigquery.migration.v2.GcsReportLogMessage.message
	Message *string `json:"message,omitempty"`

	// The script context (obfuscated) in which the error occurred
	// +kcc:proto:field=google.cloud.bigquery.migration.v2.GcsReportLogMessage.script_context
	ScriptContext *string `json:"scriptContext,omitempty"`

	// Category of the error/warning. Example: SyntaxError
	// +kcc:proto:field=google.cloud.bigquery.migration.v2.GcsReportLogMessage.action
	Action *string `json:"action,omitempty"`

	// Effect of the error/warning. Example: COMPATIBILITY
	// +kcc:proto:field=google.cloud.bigquery.migration.v2.GcsReportLogMessage.effect
	Effect *string `json:"effect,omitempty"`

	// Name of the affected object in the log message.
	// +kcc:proto:field=google.cloud.bigquery.migration.v2.GcsReportLogMessage.object_name
	ObjectName *string `json:"objectName,omitempty"`
}

// +kcc:proto=google.cloud.bigquery.migration.v2.GreenplumDialect
type GreenplumDialect struct {
}

// +kcc:proto=google.cloud.bigquery.migration.v2.HiveQLDialect
type HiveQLDialect struct {
}

// +kcc:proto=google.cloud.bigquery.migration.v2.Literal
type Literal struct {
	// Literal string data.
	// +kcc:proto:field=google.cloud.bigquery.migration.v2.Literal.literal_string
	LiteralString *string `json:"literalString,omitempty"`

	// Literal byte data.
	// +kcc:proto:field=google.cloud.bigquery.migration.v2.Literal.literal_bytes
	LiteralBytes []byte `json:"literalBytes,omitempty"`

	// Required. The identifier of the literal entry.
	// +kcc:proto:field=google.cloud.bigquery.migration.v2.Literal.relative_path
	RelativePath *string `json:"relativePath,omitempty"`
}

// +kcc:proto=google.cloud.bigquery.migration.v2.MigrationTask
type MigrationTask struct {
	// Task configuration for CW Batch/Offline SQL Translation.
	// +kcc:proto:field=google.cloud.bigquery.migration.v2.MigrationTask.translation_config_details
	TranslationConfigDetails *TranslationConfigDetails `json:"translationConfigDetails,omitempty"`

	// Task details for unified SQL Translation.
	// +kcc:proto:field=google.cloud.bigquery.migration.v2.MigrationTask.translation_details
	TranslationDetails *TranslationDetails `json:"translationDetails,omitempty"`

	// The type of the task. This must be one of the supported task types:
	//  Translation_Teradata2BQ, Translation_Redshift2BQ, Translation_Bteq2BQ,
	//  Translation_Oracle2BQ, Translation_HiveQL2BQ, Translation_SparkSQL2BQ,
	//  Translation_Snowflake2BQ, Translation_Netezza2BQ,
	//  Translation_AzureSynapse2BQ, Translation_Vertica2BQ,
	//  Translation_SQLServer2BQ, Translation_Presto2BQ, Translation_MySQL2BQ,
	//  Translation_Postgresql2BQ, Translation_SQLite2BQ, Translation_Greenplum2BQ.
	// +kcc:proto:field=google.cloud.bigquery.migration.v2.MigrationTask.type
	Type *string `json:"type,omitempty"`

	// Time when the task was created.
	// +kcc:proto:field=google.cloud.bigquery.migration.v2.MigrationTask.create_time
	CreateTime *string `json:"createTime,omitempty"`

	// Time when the task was last updated.
	// +kcc:proto:field=google.cloud.bigquery.migration.v2.MigrationTask.last_update_time
	LastUpdateTime *string `json:"lastUpdateTime,omitempty"`

	// The number or resources with errors. Note: This is not the total
	//  number of errors as each resource can have more than one error.
	//  This is used to indicate truncation by having a `resource_error_count`
	//  that is higher than the size of `resource_error_details`.
	// +kcc:proto:field=google.cloud.bigquery.migration.v2.MigrationTask.resource_error_count
	ResourceErrorCount *int32 `json:"resourceErrorCount,omitempty"`

	// The metrics for the task.
	// +kcc:proto:field=google.cloud.bigquery.migration.v2.MigrationTask.metrics
	Metrics []TimeSeries `json:"metrics,omitempty"`

	// Count of all the processing errors in this task and its subtasks.
	// +kcc:proto:field=google.cloud.bigquery.migration.v2.MigrationTask.total_processing_error_count
	TotalProcessingErrorCount *int32 `json:"totalProcessingErrorCount,omitempty"`

	// Count of all the resource errors in this task and its subtasks.
	// +kcc:proto:field=google.cloud.bigquery.migration.v2.MigrationTask.total_resource_error_count
	TotalResourceErrorCount *int32 `json:"totalResourceErrorCount,omitempty"`
}

// +kcc:proto=google.cloud.bigquery.migration.v2.MigrationTaskResult
type MigrationTaskResult struct {
	// Details specific to translation task types.
	// +kcc:proto:field=google.cloud.bigquery.migration.v2.MigrationTaskResult.translation_task_result
	TranslationTaskResult *TranslationTaskResult `json:"translationTaskResult,omitempty"`
}

// +kcc:proto=google.cloud.bigquery.migration.v2.MigrationWorkflow
type MigrationWorkflow struct {

	// The display name of the workflow. This can be set to give a workflow
	//  a descriptive name. There is no guarantee or enforcement of uniqueness.
	// +kcc:proto:field=google.cloud.bigquery.migration.v2.MigrationWorkflow.display_name
	DisplayName *string `json:"displayName,omitempty"`

	// TODO: unsupported map type with key string and value message


	// Time when the workflow was created.
	// +kcc:proto:field=google.cloud.bigquery.migration.v2.MigrationWorkflow.create_time
	CreateTime *string `json:"createTime,omitempty"`

	// Time when the workflow was last updated.
	// +kcc:proto:field=google.cloud.bigquery.migration.v2.MigrationWorkflow.last_update_time
	LastUpdateTime *string `json:"lastUpdateTime,omitempty"`
}

// +kcc:proto=google.cloud.bigquery.migration.v2.MySQLDialect
type MySQLDialect struct {
}

// +kcc:proto=google.cloud.bigquery.migration.v2.NameMappingKey
type NameMappingKey struct {
	// The type of object that is being mapped.
	// +kcc:proto:field=google.cloud.bigquery.migration.v2.NameMappingKey.type
	Type *string `json:"type,omitempty"`

	// The database name (BigQuery project ID equivalent in the source data
	//  warehouse).
	// +kcc:proto:field=google.cloud.bigquery.migration.v2.NameMappingKey.database
	Database *string `json:"database,omitempty"`

	// The schema name (BigQuery dataset equivalent in the source data warehouse).
	// +kcc:proto:field=google.cloud.bigquery.migration.v2.NameMappingKey.schema
	Schema *string `json:"schema,omitempty"`

	// The relation name (BigQuery table or view equivalent in the source data
	//  warehouse).
	// +kcc:proto:field=google.cloud.bigquery.migration.v2.NameMappingKey.relation
	Relation *string `json:"relation,omitempty"`

	// The attribute name (BigQuery column equivalent in the source data
	//  warehouse).
	// +kcc:proto:field=google.cloud.bigquery.migration.v2.NameMappingKey.attribute
	Attribute *string `json:"attribute,omitempty"`
}

// +kcc:proto=google.cloud.bigquery.migration.v2.NameMappingValue
type NameMappingValue struct {
	// The database name (BigQuery project ID equivalent in the target data
	//  warehouse).
	// +kcc:proto:field=google.cloud.bigquery.migration.v2.NameMappingValue.database
	Database *string `json:"database,omitempty"`

	// The schema name (BigQuery dataset equivalent in the target data warehouse).
	// +kcc:proto:field=google.cloud.bigquery.migration.v2.NameMappingValue.schema
	Schema *string `json:"schema,omitempty"`

	// The relation name (BigQuery table or view equivalent in the target data
	//  warehouse).
	// +kcc:proto:field=google.cloud.bigquery.migration.v2.NameMappingValue.relation
	Relation *string `json:"relation,omitempty"`

	// The attribute name (BigQuery column equivalent in the target data
	//  warehouse).
	// +kcc:proto:field=google.cloud.bigquery.migration.v2.NameMappingValue.attribute
	Attribute *string `json:"attribute,omitempty"`
}

// +kcc:proto=google.cloud.bigquery.migration.v2.NetezzaDialect
type NetezzaDialect struct {
}

// +kcc:proto=google.cloud.bigquery.migration.v2.ObjectNameMapping
type ObjectNameMapping struct {
	// The name of the object in source that is being mapped.
	// +kcc:proto:field=google.cloud.bigquery.migration.v2.ObjectNameMapping.source
	Source *NameMappingKey `json:"source,omitempty"`

	// The desired target name of the object that is being mapped.
	// +kcc:proto:field=google.cloud.bigquery.migration.v2.ObjectNameMapping.target
	Target *NameMappingValue `json:"target,omitempty"`
}

// +kcc:proto=google.cloud.bigquery.migration.v2.ObjectNameMappingList
type ObjectNameMappingList struct {
	// The elements of the object name map.
	// +kcc:proto:field=google.cloud.bigquery.migration.v2.ObjectNameMappingList.name_map
	NameMap []ObjectNameMapping `json:"nameMap,omitempty"`
}

// +kcc:proto=google.cloud.bigquery.migration.v2.OracleDialect
type OracleDialect struct {
}

// +kcc:proto=google.cloud.bigquery.migration.v2.Point
type Point struct {
	// The time interval to which the data point applies.  For `GAUGE` metrics,
	//  the start time does not need to be supplied, but if it is supplied, it must
	//  equal the end time.  For `DELTA` metrics, the start and end time should
	//  specify a non-zero interval, with subsequent points specifying contiguous
	//  and non-overlapping intervals.  For `CUMULATIVE` metrics, the start and end
	//  time should specify a non-zero interval, with subsequent points specifying
	//  the same start time and increasing end times, until an event resets the
	//  cumulative value to zero and sets a new start time for the following
	//  points.
	// +kcc:proto:field=google.cloud.bigquery.migration.v2.Point.interval
	Interval *TimeInterval `json:"interval,omitempty"`

	// The value of the data point.
	// +kcc:proto:field=google.cloud.bigquery.migration.v2.Point.value
	Value *TypedValue `json:"value,omitempty"`
}

// +kcc:proto=google.cloud.bigquery.migration.v2.PostgresqlDialect
type PostgresqlDialect struct {
}

// +kcc:proto=google.cloud.bigquery.migration.v2.PrestoDialect
type PrestoDialect struct {
}

// +kcc:proto=google.cloud.bigquery.migration.v2.RedshiftDialect
type RedshiftDialect struct {
}

// +kcc:proto=google.cloud.bigquery.migration.v2.ResourceErrorDetail
type ResourceErrorDetail struct {
	// Required. Information about the resource where the error is located.
	// +kcc:proto:field=google.cloud.bigquery.migration.v2.ResourceErrorDetail.resource_info
	ResourceInfo *ResourceInfo `json:"resourceInfo,omitempty"`

	// Required. The error details for the resource.
	// +kcc:proto:field=google.cloud.bigquery.migration.v2.ResourceErrorDetail.error_details
	ErrorDetails []ErrorDetail `json:"errorDetails,omitempty"`

	// Required. How many errors there are in total for the resource. Truncation
	//  can be indicated by having an `error_count` that is higher than the size of
	//  `error_details`.
	// +kcc:proto:field=google.cloud.bigquery.migration.v2.ResourceErrorDetail.error_count
	ErrorCount *int32 `json:"errorCount,omitempty"`
}

// +kcc:proto=google.cloud.bigquery.migration.v2.SQLServerDialect
type SQLServerDialect struct {
}

// +kcc:proto=google.cloud.bigquery.migration.v2.SQLiteDialect
type SQLiteDialect struct {
}

// +kcc:proto=google.cloud.bigquery.migration.v2.SnowflakeDialect
type SnowflakeDialect struct {
}

// +kcc:proto=google.cloud.bigquery.migration.v2.SourceEnv
type SourceEnv struct {
	// The default database name to fully qualify SQL objects when their database
	//  name is missing.
	// +kcc:proto:field=google.cloud.bigquery.migration.v2.SourceEnv.default_database
	DefaultDatabase *string `json:"defaultDatabase,omitempty"`

	// The schema search path. When SQL objects are missing schema name,
	//  translation engine will search through this list to find the value.
	// +kcc:proto:field=google.cloud.bigquery.migration.v2.SourceEnv.schema_search_path
	SchemaSearchPath []string `json:"schemaSearchPath,omitempty"`

	// Optional. Expects a valid BigQuery dataset ID that exists, e.g.,
	//  project-123.metadata_store_123.  If specified, translation will search and
	//  read the required schema information from a metadata store in this dataset.
	//  If metadata store doesn't exist, translation will parse the metadata file
	//  and upload the schema info to a temp table in the dataset to speed up
	//  future translation jobs.
	// +kcc:proto:field=google.cloud.bigquery.migration.v2.SourceEnv.metadata_store_dataset
	MetadataStoreDataset *string `json:"metadataStoreDataset,omitempty"`
}

// +kcc:proto=google.cloud.bigquery.migration.v2.SourceEnvironment
type SourceEnvironment struct {
	// The default database name to fully qualify SQL objects when their database
	//  name is missing.
	// +kcc:proto:field=google.cloud.bigquery.migration.v2.SourceEnvironment.default_database
	DefaultDatabase *string `json:"defaultDatabase,omitempty"`

	// The schema search path. When SQL objects are missing schema name,
	//  translation engine will search through this list to find the value.
	// +kcc:proto:field=google.cloud.bigquery.migration.v2.SourceEnvironment.schema_search_path
	SchemaSearchPath []string `json:"schemaSearchPath,omitempty"`

	// Optional. Expects a validQ BigQuery dataset ID that exists, e.g.,
	//  project-123.metadata_store_123.  If specified, translation will search and
	//  read the required schema information from a metadata store in this dataset.
	//  If metadata store doesn't exist, translation will parse the metadata file
	//  and upload the schema info to a temp table in the dataset to speed up
	//  future translation jobs.
	// +kcc:proto:field=google.cloud.bigquery.migration.v2.SourceEnvironment.metadata_store_dataset
	MetadataStoreDataset *string `json:"metadataStoreDataset,omitempty"`
}

// +kcc:proto=google.cloud.bigquery.migration.v2.SourceSpec
type SourceSpec struct {
	// The base URI for all files to be read in as sources for translation.
	// +kcc:proto:field=google.cloud.bigquery.migration.v2.SourceSpec.base_uri
	BaseURI *string `json:"baseURI,omitempty"`

	// Source literal.
	// +kcc:proto:field=google.cloud.bigquery.migration.v2.SourceSpec.literal
	Literal *Literal `json:"literal,omitempty"`

	// Optional. The optional field to specify the encoding of the sql bytes.
	// +kcc:proto:field=google.cloud.bigquery.migration.v2.SourceSpec.encoding
	Encoding *string `json:"encoding,omitempty"`
}

// +kcc:proto=google.cloud.bigquery.migration.v2.SourceTargetMapping
type SourceTargetMapping struct {
	// The source SQL or the path to it.
	// +kcc:proto:field=google.cloud.bigquery.migration.v2.SourceTargetMapping.source_spec
	SourceSpec *SourceSpec `json:"sourceSpec,omitempty"`

	// The target SQL or the path for it.
	// +kcc:proto:field=google.cloud.bigquery.migration.v2.SourceTargetMapping.target_spec
	TargetSpec *TargetSpec `json:"targetSpec,omitempty"`
}

// +kcc:proto=google.cloud.bigquery.migration.v2.SparkSQLDialect
type SparkSQLDialect struct {
}

// +kcc:proto=google.cloud.bigquery.migration.v2.TargetSpec
type TargetSpec struct {
	// The relative path for the target data. Given source file
	//  `base_uri/input/sql`, the output would be
	//  `target_base_uri/sql/relative_path/input.sql`.
	// +kcc:proto:field=google.cloud.bigquery.migration.v2.TargetSpec.relative_path
	RelativePath *string `json:"relativePath,omitempty"`
}

// +kcc:proto=google.cloud.bigquery.migration.v2.TeradataDialect
type TeradataDialect struct {
	// Which Teradata sub-dialect mode the user specifies.
	// +kcc:proto:field=google.cloud.bigquery.migration.v2.TeradataDialect.mode
	Mode *string `json:"mode,omitempty"`
}

// +kcc:proto=google.cloud.bigquery.migration.v2.TimeInterval
type TimeInterval struct {
	// Optional. The beginning of the time interval.  The default value
	//  for the start time is the end time. The start time must not be
	//  later than the end time.
	// +kcc:proto:field=google.cloud.bigquery.migration.v2.TimeInterval.start_time
	StartTime *string `json:"startTime,omitempty"`

	// Required. The end of the time interval.
	// +kcc:proto:field=google.cloud.bigquery.migration.v2.TimeInterval.end_time
	EndTime *string `json:"endTime,omitempty"`
}

// +kcc:proto=google.cloud.bigquery.migration.v2.TimeSeries
type TimeSeries struct {
	// Required. The name of the metric.
	//
	//  If the metric is not known by the service yet, it will be auto-created.
	// +kcc:proto:field=google.cloud.bigquery.migration.v2.TimeSeries.metric
	Metric *string `json:"metric,omitempty"`

	// Required. The value type of the time series.
	// +kcc:proto:field=google.cloud.bigquery.migration.v2.TimeSeries.value_type
	ValueType *string `json:"valueType,omitempty"`

	// Optional. The metric kind of the time series.
	//
	//  If present, it must be the same as the metric kind of the associated
	//  metric. If the associated metric's descriptor must be auto-created, then
	//  this field specifies the metric kind of the new descriptor and must be
	//  either `GAUGE` (the default) or `CUMULATIVE`.
	// +kcc:proto:field=google.cloud.bigquery.migration.v2.TimeSeries.metric_kind
	MetricKind *string `json:"metricKind,omitempty"`

	// Required. The data points of this time series. When listing time series,
	//  points are returned in reverse time order.
	//
	//  When creating a time series, this field must contain exactly one point and
	//  the point's type must be the same as the value type of the associated
	//  metric. If the associated metric's descriptor must be auto-created, then
	//  the value type of the descriptor is determined by the point's type, which
	//  must be `BOOL`, `INT64`, `DOUBLE`, or `DISTRIBUTION`.
	// +kcc:proto:field=google.cloud.bigquery.migration.v2.TimeSeries.points
	Points []Point `json:"points,omitempty"`
}

// +kcc:proto=google.cloud.bigquery.migration.v2.TranslationConfigDetails
type TranslationConfigDetails struct {
	// The Cloud Storage path for a directory of files to translate in a task.
	// +kcc:proto:field=google.cloud.bigquery.migration.v2.TranslationConfigDetails.gcs_source_path
	GcsSourcePath *string `json:"gcsSourcePath,omitempty"`

	// The Cloud Storage path to write back the corresponding input files to.
	// +kcc:proto:field=google.cloud.bigquery.migration.v2.TranslationConfigDetails.gcs_target_path
	GcsTargetPath *string `json:"gcsTargetPath,omitempty"`

	// The mapping of objects to their desired output names in list form.
	// +kcc:proto:field=google.cloud.bigquery.migration.v2.TranslationConfigDetails.name_mapping_list
	NameMappingList *ObjectNameMappingList `json:"nameMappingList,omitempty"`

	// The dialect of the input files.
	// +kcc:proto:field=google.cloud.bigquery.migration.v2.TranslationConfigDetails.source_dialect
	SourceDialect *Dialect `json:"sourceDialect,omitempty"`

	// The target dialect for the engine to translate the input to.
	// +kcc:proto:field=google.cloud.bigquery.migration.v2.TranslationConfigDetails.target_dialect
	TargetDialect *Dialect `json:"targetDialect,omitempty"`

	// The default source environment values for the translation.
	// +kcc:proto:field=google.cloud.bigquery.migration.v2.TranslationConfigDetails.source_env
	SourceEnv *SourceEnv `json:"sourceEnv,omitempty"`

	// The indicator to show translation request initiator.
	// +kcc:proto:field=google.cloud.bigquery.migration.v2.TranslationConfigDetails.request_source
	RequestSource *string `json:"requestSource,omitempty"`

	// The types of output to generate, e.g. sql, metadata etc. If not specified,
	//  a default set of targets will be generated. Some additional target types
	//  may be slower to generate. See the documentation for the set of available
	//  target types.
	// +kcc:proto:field=google.cloud.bigquery.migration.v2.TranslationConfigDetails.target_types
	TargetTypes []string `json:"targetTypes,omitempty"`
}

// +kcc:proto=google.cloud.bigquery.migration.v2.TranslationDetails
type TranslationDetails struct {
	// The mapping from source to target SQL.
	// +kcc:proto:field=google.cloud.bigquery.migration.v2.TranslationDetails.source_target_mapping
	SourceTargetMapping []SourceTargetMapping `json:"sourceTargetMapping,omitempty"`

	// The base URI for all writes to persistent storage.
	// +kcc:proto:field=google.cloud.bigquery.migration.v2.TranslationDetails.target_base_uri
	TargetBaseURI *string `json:"targetBaseURI,omitempty"`

	// The default source environment values for the translation.
	// +kcc:proto:field=google.cloud.bigquery.migration.v2.TranslationDetails.source_environment
	SourceEnvironment *SourceEnvironment `json:"sourceEnvironment,omitempty"`

	// The list of literal targets that will be directly returned to the response.
	//  Each entry consists of the constructed path, EXCLUDING the base path. Not
	//  providing a target_base_uri will prevent writing to persistent storage.
	// +kcc:proto:field=google.cloud.bigquery.migration.v2.TranslationDetails.target_return_literals
	TargetReturnLiterals []string `json:"targetReturnLiterals,omitempty"`

	// The types of output to generate, e.g. sql, metadata,
	//  lineage_from_sql_scripts, etc. If not specified, a default set of
	//  targets will be generated. Some additional target types may be slower to
	//  generate. See the documentation for the set of available target types.
	// +kcc:proto:field=google.cloud.bigquery.migration.v2.TranslationDetails.target_types
	TargetTypes []string `json:"targetTypes,omitempty"`
}

// +kcc:proto=google.cloud.bigquery.migration.v2.TranslationTaskResult
type TranslationTaskResult struct {
	// The list of the translated literals.
	// +kcc:proto:field=google.cloud.bigquery.migration.v2.TranslationTaskResult.translated_literals
	TranslatedLiterals []Literal `json:"translatedLiterals,omitempty"`

	// The records from the aggregate CSV report for a migration workflow.
	// +kcc:proto:field=google.cloud.bigquery.migration.v2.TranslationTaskResult.report_log_messages
	ReportLogMessages []GcsReportLogMessage `json:"reportLogMessages,omitempty"`
}

// +kcc:proto=google.cloud.bigquery.migration.v2.TypedValue
type TypedValue struct {
	// A Boolean value: `true` or `false`.
	// +kcc:proto:field=google.cloud.bigquery.migration.v2.TypedValue.bool_value
	BoolValue *bool `json:"boolValue,omitempty"`

	// A 64-bit integer. Its range is approximately `+/-9.2x10^18`.
	// +kcc:proto:field=google.cloud.bigquery.migration.v2.TypedValue.int64_value
	Int64Value *int64 `json:"int64Value,omitempty"`

	// A 64-bit double-precision floating-point number. Its magnitude
	//  is approximately `+/-10^(+/-300)` and it has 16 significant digits of
	//  precision.
	// +kcc:proto:field=google.cloud.bigquery.migration.v2.TypedValue.double_value
	DoubleValue *float64 `json:"doubleValue,omitempty"`

	// A variable-length string value.
	// +kcc:proto:field=google.cloud.bigquery.migration.v2.TypedValue.string_value
	StringValue *string `json:"stringValue,omitempty"`

	// A distribution value.
	// +kcc:proto:field=google.cloud.bigquery.migration.v2.TypedValue.distribution_value
	DistributionValue *Distribution `json:"distributionValue,omitempty"`
}

// +kcc:proto=google.cloud.bigquery.migration.v2.VerticaDialect
type VerticaDialect struct {
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

// +kcc:proto=google.rpc.ErrorInfo
type ErrorInfo struct {
	// The reason of the error. This is a constant value that identifies the
	//  proximate cause of the error. Error reasons are unique within a particular
	//  domain of errors. This should be at most 63 characters and match a
	//  regular expression of `[A-Z][A-Z0-9_]+[A-Z0-9]`, which represents
	//  UPPER_SNAKE_CASE.
	// +kcc:proto:field=google.rpc.ErrorInfo.reason
	Reason *string `json:"reason,omitempty"`

	// The logical grouping to which the "reason" belongs. The error domain
	//  is typically the registered service name of the tool or product that
	//  generates the error. Example: "pubsub.googleapis.com". If the error is
	//  generated by some common infrastructure, the error domain must be a
	//  globally unique value that identifies the infrastructure. For Google API
	//  infrastructure, the error domain is "googleapis.com".
	// +kcc:proto:field=google.rpc.ErrorInfo.domain
	Domain *string `json:"domain,omitempty"`

	// Additional structured details about this error.
	//
	//  Keys must match a regular expression of `[a-z][a-zA-Z0-9-_]+` but should
	//  ideally be lowerCamelCase. Also, they must be limited to 64 characters in
	//  length. When identifying the current value of an exceeded limit, the units
	//  should be contained in the key, not the value.  For example, rather than
	//  `{"instanceLimit": "100/request"}`, should be returned as,
	//  `{"instanceLimitPerRequest": "100"}`, if the client exceeds the number of
	//  instances that can be created in a single (batch) request.
	// +kcc:proto:field=google.rpc.ErrorInfo.metadata
	Metadata map[string]string `json:"metadata,omitempty"`
}

// +kcc:proto=google.rpc.ResourceInfo
type ResourceInfo struct {
	// A name for the type of resource being accessed, e.g. "sql table",
	//  "cloud storage bucket", "file", "Google calendar"; or the type URL
	//  of the resource: e.g. "type.googleapis.com/google.pubsub.v1.Topic".
	// +kcc:proto:field=google.rpc.ResourceInfo.resource_type
	ResourceType *string `json:"resourceType,omitempty"`

	// The name of the resource being accessed.  For example, a shared calendar
	//  name: "example.com_4fghdhgsrgh@group.calendar.google.com", if the current
	//  error is
	//  [google.rpc.Code.PERMISSION_DENIED][google.rpc.Code.PERMISSION_DENIED].
	// +kcc:proto:field=google.rpc.ResourceInfo.resource_name
	ResourceName *string `json:"resourceName,omitempty"`

	// The owner of the resource (optional).
	//  For example, "user:<owner email>" or "project:<Google developer project
	//  id>".
	// +kcc:proto:field=google.rpc.ResourceInfo.owner
	Owner *string `json:"owner,omitempty"`

	// Describes what error is encountered when accessing this resource.
	//  For example, updating a cloud project may require the `writer` permission
	//  on the developer console project.
	// +kcc:proto:field=google.rpc.ResourceInfo.description
	Description *string `json:"description,omitempty"`
}

// +kcc:proto=google.cloud.bigquery.migration.v2.MigrationWorkflow
type MigrationWorkflowObservedState struct {
	// Output only. Immutable. Identifier. The unique identifier for the migration
	//  workflow. The ID is server-generated.
	//
	//  Example: `projects/123/locations/us/workflows/345`
	// +kcc:proto:field=google.cloud.bigquery.migration.v2.MigrationWorkflow.name
	Name *string `json:"name,omitempty"`

	// Output only. That status of the workflow.
	// +kcc:proto:field=google.cloud.bigquery.migration.v2.MigrationWorkflow.state
	State *string `json:"state,omitempty"`
}
