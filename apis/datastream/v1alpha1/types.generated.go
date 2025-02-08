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


// +kcc:proto=google.cloud.datastream.v1.AvroFileFormat
type AvroFileFormat struct {
}

// +kcc:proto=google.cloud.datastream.v1.BigQueryDestinationConfig
type BigQueryDestinationConfig struct {
	// Single destination dataset.
	// +kcc:proto:field=google.cloud.datastream.v1.BigQueryDestinationConfig.single_target_dataset
	SingleTargetDataset *BigQueryDestinationConfig_SingleTargetDataset `json:"singleTargetDataset,omitempty"`

	// Source hierarchy datasets.
	// +kcc:proto:field=google.cloud.datastream.v1.BigQueryDestinationConfig.source_hierarchy_datasets
	SourceHierarchyDatasets *BigQueryDestinationConfig_SourceHierarchyDatasets `json:"sourceHierarchyDatasets,omitempty"`

	// The guaranteed data freshness (in seconds) when querying tables created by
	//  the stream. Editing this field will only affect new tables created in the
	//  future, but existing tables will not be impacted. Lower values mean that
	//  queries will return fresher data, but may result in higher cost.
	// +kcc:proto:field=google.cloud.datastream.v1.BigQueryDestinationConfig.data_freshness
	DataFreshness *string `json:"dataFreshness,omitempty"`

	// The standard mode
	// +kcc:proto:field=google.cloud.datastream.v1.BigQueryDestinationConfig.merge
	Merge *BigQueryDestinationConfig_Merge `json:"merge,omitempty"`

	// Append only mode
	// +kcc:proto:field=google.cloud.datastream.v1.BigQueryDestinationConfig.append_only
	AppendOnly *BigQueryDestinationConfig_AppendOnly `json:"appendOnly,omitempty"`
}

// +kcc:proto=google.cloud.datastream.v1.BigQueryDestinationConfig.AppendOnly
type BigQueryDestinationConfig_AppendOnly struct {
}

// +kcc:proto=google.cloud.datastream.v1.BigQueryDestinationConfig.Merge
type BigQueryDestinationConfig_Merge struct {
}

// +kcc:proto=google.cloud.datastream.v1.BigQueryDestinationConfig.SingleTargetDataset
type BigQueryDestinationConfig_SingleTargetDataset struct {
	// The dataset ID of the target dataset.
	//  DatasetIds allowed characters:
	//  https://cloud.google.com/bigquery/docs/reference/rest/v2/datasets#datasetreference.
	// +kcc:proto:field=google.cloud.datastream.v1.BigQueryDestinationConfig.SingleTargetDataset.dataset_id
	DatasetID *string `json:"datasetID,omitempty"`
}

// +kcc:proto=google.cloud.datastream.v1.BigQueryDestinationConfig.SourceHierarchyDatasets
type BigQueryDestinationConfig_SourceHierarchyDatasets struct {
	// The dataset template to use for dynamic dataset creation.
	// +kcc:proto:field=google.cloud.datastream.v1.BigQueryDestinationConfig.SourceHierarchyDatasets.dataset_template
	DatasetTemplate *BigQueryDestinationConfig_SourceHierarchyDatasets_DatasetTemplate `json:"datasetTemplate,omitempty"`
}

// +kcc:proto=google.cloud.datastream.v1.BigQueryDestinationConfig.SourceHierarchyDatasets.DatasetTemplate
type BigQueryDestinationConfig_SourceHierarchyDatasets_DatasetTemplate struct {
	// Required. The geographic location where the dataset should reside. See
	//  https://cloud.google.com/bigquery/docs/locations for supported
	//  locations.
	// +kcc:proto:field=google.cloud.datastream.v1.BigQueryDestinationConfig.SourceHierarchyDatasets.DatasetTemplate.location
	Location *string `json:"location,omitempty"`

	// If supplied, every created dataset will have its name prefixed by the
	//  provided value. The prefix and name will be separated by an underscore.
	//  i.e. <prefix>_<dataset_name>.
	// +kcc:proto:field=google.cloud.datastream.v1.BigQueryDestinationConfig.SourceHierarchyDatasets.DatasetTemplate.dataset_id_prefix
	DatasetIDPrefix *string `json:"datasetIDPrefix,omitempty"`

	// Describes the Cloud KMS encryption key that will be used to
	//  protect destination BigQuery table. The BigQuery Service Account
	//  associated with your project requires access to this encryption key.
	//  i.e.
	//  projects/{project}/locations/{location}/keyRings/{key_ring}/cryptoKeys/{cryptoKey}.
	//  See https://cloud.google.com/bigquery/docs/customer-managed-encryption
	//  for more information.
	// +kcc:proto:field=google.cloud.datastream.v1.BigQueryDestinationConfig.SourceHierarchyDatasets.DatasetTemplate.kms_key_name
	KMSKeyName *string `json:"kmsKeyName,omitempty"`
}

// +kcc:proto=google.cloud.datastream.v1.DestinationConfig
type DestinationConfig struct {
	// Required. Destination connection profile resource.
	//  Format: `projects/{project}/locations/{location}/connectionProfiles/{name}`
	// +kcc:proto:field=google.cloud.datastream.v1.DestinationConfig.destination_connection_profile
	DestinationConnectionProfile *string `json:"destinationConnectionProfile,omitempty"`

	// A configuration for how data should be loaded to Cloud Storage.
	// +kcc:proto:field=google.cloud.datastream.v1.DestinationConfig.gcs_destination_config
	GcsDestinationConfig *GcsDestinationConfig `json:"gcsDestinationConfig,omitempty"`

	// BigQuery destination configuration.
	// +kcc:proto:field=google.cloud.datastream.v1.DestinationConfig.bigquery_destination_config
	BigqueryDestinationConfig *BigQueryDestinationConfig `json:"bigqueryDestinationConfig,omitempty"`
}

// +kcc:proto=google.cloud.datastream.v1.Error
type Error struct {
	// A title that explains the reason for the error.
	// +kcc:proto:field=google.cloud.datastream.v1.Error.reason
	Reason *string `json:"reason,omitempty"`

	// A unique identifier for this specific error,
	//  allowing it to be traced throughout the system in logs and API responses.
	// +kcc:proto:field=google.cloud.datastream.v1.Error.error_uuid
	ErrorUuid *string `json:"errorUuid,omitempty"`

	// A message containing more information about the error that occurred.
	// +kcc:proto:field=google.cloud.datastream.v1.Error.message
	Message *string `json:"message,omitempty"`

	// The time when the error occurred.
	// +kcc:proto:field=google.cloud.datastream.v1.Error.error_time
	ErrorTime *string `json:"errorTime,omitempty"`

	// Additional information about the error.
	// +kcc:proto:field=google.cloud.datastream.v1.Error.details
	Details map[string]string `json:"details,omitempty"`
}

// +kcc:proto=google.cloud.datastream.v1.GcsDestinationConfig
type GcsDestinationConfig struct {
	// Path inside the Cloud Storage bucket to write data to.
	// +kcc:proto:field=google.cloud.datastream.v1.GcsDestinationConfig.path
	Path *string `json:"path,omitempty"`

	// The maximum file size to be saved in the bucket.
	// +kcc:proto:field=google.cloud.datastream.v1.GcsDestinationConfig.file_rotation_mb
	FileRotationMb *int32 `json:"fileRotationMb,omitempty"`

	// The maximum duration for which new events are added before a file is
	//  closed and a new file is created. Values within the range of 15-60 seconds
	//  are allowed.
	// +kcc:proto:field=google.cloud.datastream.v1.GcsDestinationConfig.file_rotation_interval
	FileRotationInterval *string `json:"fileRotationInterval,omitempty"`

	// AVRO file format configuration.
	// +kcc:proto:field=google.cloud.datastream.v1.GcsDestinationConfig.avro_file_format
	AvroFileFormat *AvroFileFormat `json:"avroFileFormat,omitempty"`

	// JSON file format configuration.
	// +kcc:proto:field=google.cloud.datastream.v1.GcsDestinationConfig.json_file_format
	JsonFileFormat *JsonFileFormat `json:"jsonFileFormat,omitempty"`
}

// +kcc:proto=google.cloud.datastream.v1.JsonFileFormat
type JsonFileFormat struct {
	// The schema file format along JSON data files.
	// +kcc:proto:field=google.cloud.datastream.v1.JsonFileFormat.schema_file_format
	SchemaFileFormat *string `json:"schemaFileFormat,omitempty"`

	// Compression of the loaded JSON file.
	// +kcc:proto:field=google.cloud.datastream.v1.JsonFileFormat.compression
	Compression *string `json:"compression,omitempty"`
}

// +kcc:proto=google.cloud.datastream.v1.MysqlColumn
type MysqlColumn struct {
	// Column name.
	// +kcc:proto:field=google.cloud.datastream.v1.MysqlColumn.column
	Column *string `json:"column,omitempty"`

	// The MySQL data type. Full data types list can be found here:
	//  https://dev.mysql.com/doc/refman/8.0/en/data-types.html
	// +kcc:proto:field=google.cloud.datastream.v1.MysqlColumn.data_type
	DataType *string `json:"dataType,omitempty"`

	// Column length.
	// +kcc:proto:field=google.cloud.datastream.v1.MysqlColumn.length
	Length *int32 `json:"length,omitempty"`

	// Column collation.
	// +kcc:proto:field=google.cloud.datastream.v1.MysqlColumn.collation
	Collation *string `json:"collation,omitempty"`

	// Whether or not the column represents a primary key.
	// +kcc:proto:field=google.cloud.datastream.v1.MysqlColumn.primary_key
	PrimaryKey *bool `json:"primaryKey,omitempty"`

	// Whether or not the column can accept a null value.
	// +kcc:proto:field=google.cloud.datastream.v1.MysqlColumn.nullable
	Nullable *bool `json:"nullable,omitempty"`

	// The ordinal position of the column in the table.
	// +kcc:proto:field=google.cloud.datastream.v1.MysqlColumn.ordinal_position
	OrdinalPosition *int32 `json:"ordinalPosition,omitempty"`

	// Column precision.
	// +kcc:proto:field=google.cloud.datastream.v1.MysqlColumn.precision
	Precision *int32 `json:"precision,omitempty"`

	// Column scale.
	// +kcc:proto:field=google.cloud.datastream.v1.MysqlColumn.scale
	Scale *int32 `json:"scale,omitempty"`
}

// +kcc:proto=google.cloud.datastream.v1.MysqlDatabase
type MysqlDatabase struct {
	// Database name.
	// +kcc:proto:field=google.cloud.datastream.v1.MysqlDatabase.database
	Database *string `json:"database,omitempty"`

	// Tables in the database.
	// +kcc:proto:field=google.cloud.datastream.v1.MysqlDatabase.mysql_tables
	MysqlTables []MysqlTable `json:"mysqlTables,omitempty"`
}

// +kcc:proto=google.cloud.datastream.v1.MysqlRdbms
type MysqlRdbms struct {
	// Mysql databases on the server
	// +kcc:proto:field=google.cloud.datastream.v1.MysqlRdbms.mysql_databases
	MysqlDatabases []MysqlDatabase `json:"mysqlDatabases,omitempty"`
}

// +kcc:proto=google.cloud.datastream.v1.MysqlSourceConfig
type MysqlSourceConfig struct {
	// MySQL objects to retrieve from the source.
	// +kcc:proto:field=google.cloud.datastream.v1.MysqlSourceConfig.include_objects
	IncludeObjects *MysqlRdbms `json:"includeObjects,omitempty"`

	// MySQL objects to exclude from the stream.
	// +kcc:proto:field=google.cloud.datastream.v1.MysqlSourceConfig.exclude_objects
	ExcludeObjects *MysqlRdbms `json:"excludeObjects,omitempty"`

	// Maximum number of concurrent CDC tasks. The number should be non negative.
	//  If not set (or set to 0), the system's default value will be used.
	// +kcc:proto:field=google.cloud.datastream.v1.MysqlSourceConfig.max_concurrent_cdc_tasks
	MaxConcurrentCdcTasks *int32 `json:"maxConcurrentCdcTasks,omitempty"`

	// Maximum number of concurrent backfill tasks. The number should be non
	//  negative. If not set (or set to 0), the system's default value will be
	//  used.
	// +kcc:proto:field=google.cloud.datastream.v1.MysqlSourceConfig.max_concurrent_backfill_tasks
	MaxConcurrentBackfillTasks *int32 `json:"maxConcurrentBackfillTasks,omitempty"`

	// Use Binary log position based replication.
	// +kcc:proto:field=google.cloud.datastream.v1.MysqlSourceConfig.binary_log_position
	BinaryLogPosition *MysqlSourceConfig_BinaryLogPosition `json:"binaryLogPosition,omitempty"`

	// Use GTID based replication.
	// +kcc:proto:field=google.cloud.datastream.v1.MysqlSourceConfig.gtid
	Gtid *MysqlSourceConfig_Gtid `json:"gtid,omitempty"`
}

// +kcc:proto=google.cloud.datastream.v1.MysqlSourceConfig.BinaryLogPosition
type MysqlSourceConfig_BinaryLogPosition struct {
}

// +kcc:proto=google.cloud.datastream.v1.MysqlSourceConfig.Gtid
type MysqlSourceConfig_Gtid struct {
}

// +kcc:proto=google.cloud.datastream.v1.MysqlTable
type MysqlTable struct {
	// Table name.
	// +kcc:proto:field=google.cloud.datastream.v1.MysqlTable.table
	Table *string `json:"table,omitempty"`

	// MySQL columns in the database.
	//  When unspecified as part of include/exclude objects, includes/excludes
	//  everything.
	// +kcc:proto:field=google.cloud.datastream.v1.MysqlTable.mysql_columns
	MysqlColumns []MysqlColumn `json:"mysqlColumns,omitempty"`
}

// +kcc:proto=google.cloud.datastream.v1.OracleColumn
type OracleColumn struct {
	// Column name.
	// +kcc:proto:field=google.cloud.datastream.v1.OracleColumn.column
	Column *string `json:"column,omitempty"`

	// The Oracle data type.
	// +kcc:proto:field=google.cloud.datastream.v1.OracleColumn.data_type
	DataType *string `json:"dataType,omitempty"`

	// Column length.
	// +kcc:proto:field=google.cloud.datastream.v1.OracleColumn.length
	Length *int32 `json:"length,omitempty"`

	// Column precision.
	// +kcc:proto:field=google.cloud.datastream.v1.OracleColumn.precision
	Precision *int32 `json:"precision,omitempty"`

	// Column scale.
	// +kcc:proto:field=google.cloud.datastream.v1.OracleColumn.scale
	Scale *int32 `json:"scale,omitempty"`

	// Column encoding.
	// +kcc:proto:field=google.cloud.datastream.v1.OracleColumn.encoding
	Encoding *string `json:"encoding,omitempty"`

	// Whether or not the column represents a primary key.
	// +kcc:proto:field=google.cloud.datastream.v1.OracleColumn.primary_key
	PrimaryKey *bool `json:"primaryKey,omitempty"`

	// Whether or not the column can accept a null value.
	// +kcc:proto:field=google.cloud.datastream.v1.OracleColumn.nullable
	Nullable *bool `json:"nullable,omitempty"`

	// The ordinal position of the column in the table.
	// +kcc:proto:field=google.cloud.datastream.v1.OracleColumn.ordinal_position
	OrdinalPosition *int32 `json:"ordinalPosition,omitempty"`
}

// +kcc:proto=google.cloud.datastream.v1.OracleRdbms
type OracleRdbms struct {
	// Oracle schemas/databases in the database server.
	// +kcc:proto:field=google.cloud.datastream.v1.OracleRdbms.oracle_schemas
	OracleSchemas []OracleSchema `json:"oracleSchemas,omitempty"`
}

// +kcc:proto=google.cloud.datastream.v1.OracleSchema
type OracleSchema struct {
	// Schema name.
	// +kcc:proto:field=google.cloud.datastream.v1.OracleSchema.schema
	Schema *string `json:"schema,omitempty"`

	// Tables in the schema.
	// +kcc:proto:field=google.cloud.datastream.v1.OracleSchema.oracle_tables
	OracleTables []OracleTable `json:"oracleTables,omitempty"`
}

// +kcc:proto=google.cloud.datastream.v1.OracleSourceConfig
type OracleSourceConfig struct {
	// Oracle objects to include in the stream.
	// +kcc:proto:field=google.cloud.datastream.v1.OracleSourceConfig.include_objects
	IncludeObjects *OracleRdbms `json:"includeObjects,omitempty"`

	// Oracle objects to exclude from the stream.
	// +kcc:proto:field=google.cloud.datastream.v1.OracleSourceConfig.exclude_objects
	ExcludeObjects *OracleRdbms `json:"excludeObjects,omitempty"`

	// Maximum number of concurrent CDC tasks. The number should be non-negative.
	//  If not set (or set to 0), the system's default value is used.
	// +kcc:proto:field=google.cloud.datastream.v1.OracleSourceConfig.max_concurrent_cdc_tasks
	MaxConcurrentCdcTasks *int32 `json:"maxConcurrentCdcTasks,omitempty"`

	// Maximum number of concurrent backfill tasks. The number should be
	//  non-negative. If not set (or set to 0), the system's default value is used.
	// +kcc:proto:field=google.cloud.datastream.v1.OracleSourceConfig.max_concurrent_backfill_tasks
	MaxConcurrentBackfillTasks *int32 `json:"maxConcurrentBackfillTasks,omitempty"`

	// Drop large object values.
	// +kcc:proto:field=google.cloud.datastream.v1.OracleSourceConfig.drop_large_objects
	DropLargeObjects *OracleSourceConfig_DropLargeObjects `json:"dropLargeObjects,omitempty"`

	// Stream large object values.
	// +kcc:proto:field=google.cloud.datastream.v1.OracleSourceConfig.stream_large_objects
	StreamLargeObjects *OracleSourceConfig_StreamLargeObjects `json:"streamLargeObjects,omitempty"`

	// Use LogMiner.
	// +kcc:proto:field=google.cloud.datastream.v1.OracleSourceConfig.log_miner
	LogMiner *OracleSourceConfig_LogMiner `json:"logMiner,omitempty"`

	// Use Binary Log Parser.
	// +kcc:proto:field=google.cloud.datastream.v1.OracleSourceConfig.binary_log_parser
	BinaryLogParser *OracleSourceConfig_BinaryLogParser `json:"binaryLogParser,omitempty"`
}

// +kcc:proto=google.cloud.datastream.v1.OracleSourceConfig.BinaryLogParser
type OracleSourceConfig_BinaryLogParser struct {
	// Use Oracle ASM.
	// +kcc:proto:field=google.cloud.datastream.v1.OracleSourceConfig.BinaryLogParser.oracle_asm_log_file_access
	OracleAsmLogFileAccess *OracleSourceConfig_BinaryLogParser_OracleAsmLogFileAccess `json:"oracleAsmLogFileAccess,omitempty"`

	// Use Oracle directories.
	// +kcc:proto:field=google.cloud.datastream.v1.OracleSourceConfig.BinaryLogParser.log_file_directories
	LogFileDirectories *OracleSourceConfig_BinaryLogParser_LogFileDirectories `json:"logFileDirectories,omitempty"`
}

// +kcc:proto=google.cloud.datastream.v1.OracleSourceConfig.BinaryLogParser.LogFileDirectories
type OracleSourceConfig_BinaryLogParser_LogFileDirectories struct {
	// Required. Oracle directory for online logs.
	// +kcc:proto:field=google.cloud.datastream.v1.OracleSourceConfig.BinaryLogParser.LogFileDirectories.online_log_directory
	OnlineLogDirectory *string `json:"onlineLogDirectory,omitempty"`

	// Required. Oracle directory for archived logs.
	// +kcc:proto:field=google.cloud.datastream.v1.OracleSourceConfig.BinaryLogParser.LogFileDirectories.archived_log_directory
	ArchivedLogDirectory *string `json:"archivedLogDirectory,omitempty"`
}

// +kcc:proto=google.cloud.datastream.v1.OracleSourceConfig.BinaryLogParser.OracleAsmLogFileAccess
type OracleSourceConfig_BinaryLogParser_OracleAsmLogFileAccess struct {
}

// +kcc:proto=google.cloud.datastream.v1.OracleSourceConfig.DropLargeObjects
type OracleSourceConfig_DropLargeObjects struct {
}

// +kcc:proto=google.cloud.datastream.v1.OracleSourceConfig.LogMiner
type OracleSourceConfig_LogMiner struct {
}

// +kcc:proto=google.cloud.datastream.v1.OracleSourceConfig.StreamLargeObjects
type OracleSourceConfig_StreamLargeObjects struct {
}

// +kcc:proto=google.cloud.datastream.v1.OracleTable
type OracleTable struct {
	// Table name.
	// +kcc:proto:field=google.cloud.datastream.v1.OracleTable.table
	Table *string `json:"table,omitempty"`

	// Oracle columns in the schema.
	//  When unspecified as part of include/exclude objects, includes/excludes
	//  everything.
	// +kcc:proto:field=google.cloud.datastream.v1.OracleTable.oracle_columns
	OracleColumns []OracleColumn `json:"oracleColumns,omitempty"`
}

// +kcc:proto=google.cloud.datastream.v1.PostgresqlColumn
type PostgresqlColumn struct {
	// Column name.
	// +kcc:proto:field=google.cloud.datastream.v1.PostgresqlColumn.column
	Column *string `json:"column,omitempty"`

	// The PostgreSQL data type.
	// +kcc:proto:field=google.cloud.datastream.v1.PostgresqlColumn.data_type
	DataType *string `json:"dataType,omitempty"`

	// Column length.
	// +kcc:proto:field=google.cloud.datastream.v1.PostgresqlColumn.length
	Length *int32 `json:"length,omitempty"`

	// Column precision.
	// +kcc:proto:field=google.cloud.datastream.v1.PostgresqlColumn.precision
	Precision *int32 `json:"precision,omitempty"`

	// Column scale.
	// +kcc:proto:field=google.cloud.datastream.v1.PostgresqlColumn.scale
	Scale *int32 `json:"scale,omitempty"`

	// Whether or not the column represents a primary key.
	// +kcc:proto:field=google.cloud.datastream.v1.PostgresqlColumn.primary_key
	PrimaryKey *bool `json:"primaryKey,omitempty"`

	// Whether or not the column can accept a null value.
	// +kcc:proto:field=google.cloud.datastream.v1.PostgresqlColumn.nullable
	Nullable *bool `json:"nullable,omitempty"`

	// The ordinal position of the column in the table.
	// +kcc:proto:field=google.cloud.datastream.v1.PostgresqlColumn.ordinal_position
	OrdinalPosition *int32 `json:"ordinalPosition,omitempty"`
}

// +kcc:proto=google.cloud.datastream.v1.PostgresqlRdbms
type PostgresqlRdbms struct {
	// PostgreSQL schemas in the database server.
	// +kcc:proto:field=google.cloud.datastream.v1.PostgresqlRdbms.postgresql_schemas
	PostgresqlSchemas []PostgresqlSchema `json:"postgresqlSchemas,omitempty"`
}

// +kcc:proto=google.cloud.datastream.v1.PostgresqlSchema
type PostgresqlSchema struct {
	// Schema name.
	// +kcc:proto:field=google.cloud.datastream.v1.PostgresqlSchema.schema
	Schema *string `json:"schema,omitempty"`

	// Tables in the schema.
	// +kcc:proto:field=google.cloud.datastream.v1.PostgresqlSchema.postgresql_tables
	PostgresqlTables []PostgresqlTable `json:"postgresqlTables,omitempty"`
}

// +kcc:proto=google.cloud.datastream.v1.PostgresqlSourceConfig
type PostgresqlSourceConfig struct {
	// PostgreSQL objects to include in the stream.
	// +kcc:proto:field=google.cloud.datastream.v1.PostgresqlSourceConfig.include_objects
	IncludeObjects *PostgresqlRdbms `json:"includeObjects,omitempty"`

	// PostgreSQL objects to exclude from the stream.
	// +kcc:proto:field=google.cloud.datastream.v1.PostgresqlSourceConfig.exclude_objects
	ExcludeObjects *PostgresqlRdbms `json:"excludeObjects,omitempty"`

	// Required. Immutable. The name of the logical replication slot that's
	//  configured with the pgoutput plugin.
	// +kcc:proto:field=google.cloud.datastream.v1.PostgresqlSourceConfig.replication_slot
	ReplicationSlot *string `json:"replicationSlot,omitempty"`

	// Required. The name of the publication that includes the set of all tables
	//  that are defined in the stream's include_objects.
	// +kcc:proto:field=google.cloud.datastream.v1.PostgresqlSourceConfig.publication
	Publication *string `json:"publication,omitempty"`

	// Maximum number of concurrent backfill tasks. The number should be non
	//  negative. If not set (or set to 0), the system's default value will be
	//  used.
	// +kcc:proto:field=google.cloud.datastream.v1.PostgresqlSourceConfig.max_concurrent_backfill_tasks
	MaxConcurrentBackfillTasks *int32 `json:"maxConcurrentBackfillTasks,omitempty"`
}

// +kcc:proto=google.cloud.datastream.v1.PostgresqlTable
type PostgresqlTable struct {
	// Table name.
	// +kcc:proto:field=google.cloud.datastream.v1.PostgresqlTable.table
	Table *string `json:"table,omitempty"`

	// PostgreSQL columns in the schema.
	//  When unspecified as part of include/exclude objects,
	//  includes/excludes everything.
	// +kcc:proto:field=google.cloud.datastream.v1.PostgresqlTable.postgresql_columns
	PostgresqlColumns []PostgresqlColumn `json:"postgresqlColumns,omitempty"`
}

// +kcc:proto=google.cloud.datastream.v1.SourceConfig
type SourceConfig struct {
	// Required. Source connection profile resoource.
	//  Format: `projects/{project}/locations/{location}/connectionProfiles/{name}`
	// +kcc:proto:field=google.cloud.datastream.v1.SourceConfig.source_connection_profile
	SourceConnectionProfile *string `json:"sourceConnectionProfile,omitempty"`

	// Oracle data source configuration.
	// +kcc:proto:field=google.cloud.datastream.v1.SourceConfig.oracle_source_config
	OracleSourceConfig *OracleSourceConfig `json:"oracleSourceConfig,omitempty"`

	// MySQL data source configuration.
	// +kcc:proto:field=google.cloud.datastream.v1.SourceConfig.mysql_source_config
	MysqlSourceConfig *MysqlSourceConfig `json:"mysqlSourceConfig,omitempty"`

	// PostgreSQL data source configuration.
	// +kcc:proto:field=google.cloud.datastream.v1.SourceConfig.postgresql_source_config
	PostgresqlSourceConfig *PostgresqlSourceConfig `json:"postgresqlSourceConfig,omitempty"`

	// SQLServer data source configuration.
	// +kcc:proto:field=google.cloud.datastream.v1.SourceConfig.sql_server_source_config
	SqlServerSourceConfig *SqlServerSourceConfig `json:"sqlServerSourceConfig,omitempty"`
}

// +kcc:proto=google.cloud.datastream.v1.SqlServerChangeTables
type SqlServerChangeTables struct {
}

// +kcc:proto=google.cloud.datastream.v1.SqlServerColumn
type SqlServerColumn struct {
	// Column name.
	// +kcc:proto:field=google.cloud.datastream.v1.SqlServerColumn.column
	Column *string `json:"column,omitempty"`

	// The SQLServer data type.
	// +kcc:proto:field=google.cloud.datastream.v1.SqlServerColumn.data_type
	DataType *string `json:"dataType,omitempty"`

	// Column length.
	// +kcc:proto:field=google.cloud.datastream.v1.SqlServerColumn.length
	Length *int32 `json:"length,omitempty"`

	// Column precision.
	// +kcc:proto:field=google.cloud.datastream.v1.SqlServerColumn.precision
	Precision *int32 `json:"precision,omitempty"`

	// Column scale.
	// +kcc:proto:field=google.cloud.datastream.v1.SqlServerColumn.scale
	Scale *int32 `json:"scale,omitempty"`

	// Whether or not the column represents a primary key.
	// +kcc:proto:field=google.cloud.datastream.v1.SqlServerColumn.primary_key
	PrimaryKey *bool `json:"primaryKey,omitempty"`

	// Whether or not the column can accept a null value.
	// +kcc:proto:field=google.cloud.datastream.v1.SqlServerColumn.nullable
	Nullable *bool `json:"nullable,omitempty"`

	// The ordinal position of the column in the table.
	// +kcc:proto:field=google.cloud.datastream.v1.SqlServerColumn.ordinal_position
	OrdinalPosition *int32 `json:"ordinalPosition,omitempty"`
}

// +kcc:proto=google.cloud.datastream.v1.SqlServerRdbms
type SqlServerRdbms struct {
	// SQLServer schemas in the database server.
	// +kcc:proto:field=google.cloud.datastream.v1.SqlServerRdbms.schemas
	Schemas []SqlServerSchema `json:"schemas,omitempty"`
}

// +kcc:proto=google.cloud.datastream.v1.SqlServerSchema
type SqlServerSchema struct {
	// Schema name.
	// +kcc:proto:field=google.cloud.datastream.v1.SqlServerSchema.schema
	Schema *string `json:"schema,omitempty"`

	// Tables in the schema.
	// +kcc:proto:field=google.cloud.datastream.v1.SqlServerSchema.tables
	Tables []SqlServerTable `json:"tables,omitempty"`
}

// +kcc:proto=google.cloud.datastream.v1.SqlServerSourceConfig
type SqlServerSourceConfig struct {
	// SQLServer objects to include in the stream.
	// +kcc:proto:field=google.cloud.datastream.v1.SqlServerSourceConfig.include_objects
	IncludeObjects *SqlServerRdbms `json:"includeObjects,omitempty"`

	// SQLServer objects to exclude from the stream.
	// +kcc:proto:field=google.cloud.datastream.v1.SqlServerSourceConfig.exclude_objects
	ExcludeObjects *SqlServerRdbms `json:"excludeObjects,omitempty"`

	// Max concurrent CDC tasks.
	// +kcc:proto:field=google.cloud.datastream.v1.SqlServerSourceConfig.max_concurrent_cdc_tasks
	MaxConcurrentCdcTasks *int32 `json:"maxConcurrentCdcTasks,omitempty"`

	// Max concurrent backfill tasks.
	// +kcc:proto:field=google.cloud.datastream.v1.SqlServerSourceConfig.max_concurrent_backfill_tasks
	MaxConcurrentBackfillTasks *int32 `json:"maxConcurrentBackfillTasks,omitempty"`

	// CDC reader reads from transaction logs.
	// +kcc:proto:field=google.cloud.datastream.v1.SqlServerSourceConfig.transaction_logs
	TransactionLogs *SqlServerTransactionLogs `json:"transactionLogs,omitempty"`

	// CDC reader reads from change tables.
	// +kcc:proto:field=google.cloud.datastream.v1.SqlServerSourceConfig.change_tables
	ChangeTables *SqlServerChangeTables `json:"changeTables,omitempty"`
}

// +kcc:proto=google.cloud.datastream.v1.SqlServerTable
type SqlServerTable struct {
	// Table name.
	// +kcc:proto:field=google.cloud.datastream.v1.SqlServerTable.table
	Table *string `json:"table,omitempty"`

	// SQLServer columns in the schema.
	//  When unspecified as part of include/exclude objects,
	//  includes/excludes everything.
	// +kcc:proto:field=google.cloud.datastream.v1.SqlServerTable.columns
	Columns []SqlServerColumn `json:"columns,omitempty"`
}

// +kcc:proto=google.cloud.datastream.v1.SqlServerTransactionLogs
type SqlServerTransactionLogs struct {
}

// +kcc:proto=google.cloud.datastream.v1.Stream
type Stream struct {

	// Labels.
	// +kcc:proto:field=google.cloud.datastream.v1.Stream.labels
	Labels map[string]string `json:"labels,omitempty"`

	// Required. Display name.
	// +kcc:proto:field=google.cloud.datastream.v1.Stream.display_name
	DisplayName *string `json:"displayName,omitempty"`

	// Required. Source connection profile configuration.
	// +kcc:proto:field=google.cloud.datastream.v1.Stream.source_config
	SourceConfig *SourceConfig `json:"sourceConfig,omitempty"`

	// Required. Destination connection profile configuration.
	// +kcc:proto:field=google.cloud.datastream.v1.Stream.destination_config
	DestinationConfig *DestinationConfig `json:"destinationConfig,omitempty"`

	// The state of the stream.
	// +kcc:proto:field=google.cloud.datastream.v1.Stream.state
	State *string `json:"state,omitempty"`

	// Automatically backfill objects included in the stream source
	//  configuration. Specific objects can be excluded.
	// +kcc:proto:field=google.cloud.datastream.v1.Stream.backfill_all
	BackfillAll *Stream_BackfillAllStrategy `json:"backfillAll,omitempty"`

	// Do not automatically backfill any objects.
	// +kcc:proto:field=google.cloud.datastream.v1.Stream.backfill_none
	BackfillNone *Stream_BackfillNoneStrategy `json:"backfillNone,omitempty"`

	// Immutable. A reference to a KMS encryption key.
	//  If provided, it will be used to encrypt the data.
	//  If left blank, data will be encrypted using an internal Stream-specific
	//  encryption key provisioned through KMS.
	// +kcc:proto:field=google.cloud.datastream.v1.Stream.customer_managed_encryption_key
	CustomerManagedEncryptionKey *string `json:"customerManagedEncryptionKey,omitempty"`
}

// +kcc:proto=google.cloud.datastream.v1.Stream.BackfillAllStrategy
type Stream_BackfillAllStrategy struct {
	// Oracle data source objects to avoid backfilling.
	// +kcc:proto:field=google.cloud.datastream.v1.Stream.BackfillAllStrategy.oracle_excluded_objects
	OracleExcludedObjects *OracleRdbms `json:"oracleExcludedObjects,omitempty"`

	// MySQL data source objects to avoid backfilling.
	// +kcc:proto:field=google.cloud.datastream.v1.Stream.BackfillAllStrategy.mysql_excluded_objects
	MysqlExcludedObjects *MysqlRdbms `json:"mysqlExcludedObjects,omitempty"`

	// PostgreSQL data source objects to avoid backfilling.
	// +kcc:proto:field=google.cloud.datastream.v1.Stream.BackfillAllStrategy.postgresql_excluded_objects
	PostgresqlExcludedObjects *PostgresqlRdbms `json:"postgresqlExcludedObjects,omitempty"`

	// SQLServer data source objects to avoid backfilling
	// +kcc:proto:field=google.cloud.datastream.v1.Stream.BackfillAllStrategy.sql_server_excluded_objects
	SqlServerExcludedObjects *SqlServerRdbms `json:"sqlServerExcludedObjects,omitempty"`
}

// +kcc:proto=google.cloud.datastream.v1.Stream.BackfillNoneStrategy
type Stream_BackfillNoneStrategy struct {
}

// +kcc:proto=google.cloud.datastream.v1.Stream
type StreamObservedState struct {
	// Output only. The stream's name.
	// +kcc:proto:field=google.cloud.datastream.v1.Stream.name
	Name *string `json:"name,omitempty"`

	// Output only. The creation time of the stream.
	// +kcc:proto:field=google.cloud.datastream.v1.Stream.create_time
	CreateTime *string `json:"createTime,omitempty"`

	// Output only. The last update time of the stream.
	// +kcc:proto:field=google.cloud.datastream.v1.Stream.update_time
	UpdateTime *string `json:"updateTime,omitempty"`

	// Output only. Errors on the Stream.
	// +kcc:proto:field=google.cloud.datastream.v1.Stream.errors
	Errors []Error `json:"errors,omitempty"`

	// Output only. If the stream was recovered, the time of the last recovery.
	//  Note: This field is currently experimental.
	// +kcc:proto:field=google.cloud.datastream.v1.Stream.last_recovery_time
	LastRecoveryTime *string `json:"lastRecoveryTime,omitempty"`
}
