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

// +generated:types
// krm.group: datastream.cnrm.cloud.google.com
// krm.version: v1alpha1
// proto.service: google.cloud.datastream.v1
// resource: DatastreamPrivateConnection:PrivateConnection
// resource: DatastreamConnectionProfile:ConnectionProfile
// resource: DatastreamRoute:Route
// resource: DatastreamStream:Stream

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

// +kcc:proto=google.cloud.datastream.v1.BigQueryDestinationConfig.SourceHierarchyDatasets
type BigQueryDestinationConfig_SourceHierarchyDatasets struct {
	// The dataset template to use for dynamic dataset creation.
	// +kcc:proto:field=google.cloud.datastream.v1.BigQueryDestinationConfig.SourceHierarchyDatasets.dataset_template
	DatasetTemplate *BigQueryDestinationConfig_SourceHierarchyDatasets_DatasetTemplate `json:"datasetTemplate,omitempty"`
}

// +kcc:proto=google.cloud.datastream.v1.BigQueryProfile
type BigQueryProfile struct {
}

// +kcc:proto=google.cloud.datastream.v1.GcsDestinationConfig
type GCSDestinationConfig struct {
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

// +kcc:proto=google.cloud.datastream.v1.MysqlSslConfig
type MysqlSSLConfig struct {
	// Input only. PEM-encoded private key associated with the Client Certificate.
	//  If this field is used then the 'client_certificate' and the
	//  'ca_certificate' fields are mandatory.
	// +kcc:proto:field=google.cloud.datastream.v1.MysqlSslConfig.client_key
	ClientKey *string `json:"clientKey,omitempty"`

	// Input only. PEM-encoded certificate that will be used by the replica to
	//  authenticate against the source database server. If this field is used
	//  then the 'client_key' and the 'ca_certificate' fields are mandatory.
	// +kcc:proto:field=google.cloud.datastream.v1.MysqlSslConfig.client_certificate
	ClientCertificate *string `json:"clientCertificate,omitempty"`

	// Input only. PEM-encoded certificate of the CA that signed the source
	//  database server's certificate.
	// +kcc:proto:field=google.cloud.datastream.v1.MysqlSslConfig.ca_certificate
	CACertificate *string `json:"caCertificate,omitempty"`
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

// +kcc:proto=google.cloud.datastream.v1.OracleSslConfig
type OracleSSLConfig struct {
	// Input only. PEM-encoded certificate of the CA that signed the source
	//  database server's certificate.
	// +kcc:proto:field=google.cloud.datastream.v1.OracleSslConfig.ca_certificate
	CACertificate *string `json:"caCertificate,omitempty"`
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

// +kcc:proto=google.cloud.datastream.v1.SqlServerChangeTables
type SQLServerChangeTables struct {
}

// +kcc:proto=google.cloud.datastream.v1.SqlServerColumn
type SQLServerColumn struct {
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
type SQLServerRdbms struct {
	// SQLServer schemas in the database server.
	// +kcc:proto:field=google.cloud.datastream.v1.SqlServerRdbms.schemas
	Schemas []SQLServerSchema `json:"schemas,omitempty"`
}

// +kcc:proto=google.cloud.datastream.v1.SqlServerSchema
type SQLServerSchema struct {
	// Schema name.
	// +kcc:proto:field=google.cloud.datastream.v1.SqlServerSchema.schema
	Schema *string `json:"schema,omitempty"`

	// Tables in the schema.
	// +kcc:proto:field=google.cloud.datastream.v1.SqlServerSchema.tables
	Tables []SQLServerTable `json:"tables,omitempty"`
}

// +kcc:proto=google.cloud.datastream.v1.SqlServerSourceConfig
type SQLServerSourceConfig struct {
	// SQLServer objects to include in the stream.
	// +kcc:proto:field=google.cloud.datastream.v1.SqlServerSourceConfig.include_objects
	IncludeObjects *SQLServerRdbms `json:"includeObjects,omitempty"`

	// SQLServer objects to exclude from the stream.
	// +kcc:proto:field=google.cloud.datastream.v1.SqlServerSourceConfig.exclude_objects
	ExcludeObjects *SQLServerRdbms `json:"excludeObjects,omitempty"`

	// Max concurrent CDC tasks.
	// +kcc:proto:field=google.cloud.datastream.v1.SqlServerSourceConfig.max_concurrent_cdc_tasks
	MaxConcurrentCdcTasks *int32 `json:"maxConcurrentCdcTasks,omitempty"`

	// Max concurrent backfill tasks.
	// +kcc:proto:field=google.cloud.datastream.v1.SqlServerSourceConfig.max_concurrent_backfill_tasks
	MaxConcurrentBackfillTasks *int32 `json:"maxConcurrentBackfillTasks,omitempty"`

	// CDC reader reads from transaction logs.
	// +kcc:proto:field=google.cloud.datastream.v1.SqlServerSourceConfig.transaction_logs
	TransactionLogs *SQLServerTransactionLogs `json:"transactionLogs,omitempty"`

	// CDC reader reads from change tables.
	// +kcc:proto:field=google.cloud.datastream.v1.SqlServerSourceConfig.change_tables
	ChangeTables *SQLServerChangeTables `json:"changeTables,omitempty"`
}

// +kcc:proto=google.cloud.datastream.v1.SqlServerTable
type SQLServerTable struct {
	// Table name.
	// +kcc:proto:field=google.cloud.datastream.v1.SqlServerTable.table
	Table *string `json:"table,omitempty"`

	// SQLServer columns in the schema.
	//  When unspecified as part of include/exclude objects,
	//  includes/excludes everything.
	// +kcc:proto:field=google.cloud.datastream.v1.SqlServerTable.columns
	Columns []SQLServerColumn `json:"columns,omitempty"`
}

// +kcc:proto=google.cloud.datastream.v1.SqlServerTransactionLogs
type SQLServerTransactionLogs struct {
}

// +kcc:proto=google.cloud.datastream.v1.StaticServiceIpConnectivity
type StaticServiceIPConnectivity struct {
}

// +kcc:proto=google.cloud.datastream.v1.Stream.BackfillNoneStrategy
type Stream_BackfillNoneStrategy struct {
}

// +kcc:proto=google.cloud.datastream.v1.MysqlSslConfig
type MysqlSSLConfigObservedState struct {
	// Output only. Indicates whether the client_key field is set.
	// +kcc:proto:field=google.cloud.datastream.v1.MysqlSslConfig.client_key_set
	ClientKeySet *bool `json:"clientKeySet,omitempty"`

	// Output only. Indicates whether the client_certificate field is set.
	// +kcc:proto:field=google.cloud.datastream.v1.MysqlSslConfig.client_certificate_set
	ClientCertificateSet *bool `json:"clientCertificateSet,omitempty"`

	// Output only. Indicates whether the ca_certificate field is set.
	// +kcc:proto:field=google.cloud.datastream.v1.MysqlSslConfig.ca_certificate_set
	CACertificateSet *bool `json:"caCertificateSet,omitempty"`
}

// +kcc:proto=google.cloud.datastream.v1.OracleSslConfig
type OracleSSLConfigObservedState struct {
	// Output only. Indicates whether the ca_certificate field has been set for
	//  this Connection-Profile.
	// +kcc:proto:field=google.cloud.datastream.v1.OracleSslConfig.ca_certificate_set
	CACertificateSet *bool `json:"caCertificateSet,omitempty"`
}
