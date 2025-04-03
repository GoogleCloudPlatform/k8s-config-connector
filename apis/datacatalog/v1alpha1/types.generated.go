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
// krm.group: datacatalog.cnrm.cloud.google.com
// krm.version: v1alpha1
// proto.service: google.cloud.datacatalog.v1
// resource: DataCatalogEntry:Entry

package v1alpha1

// +kcc:proto=google.cloud.datacatalog.v1.BigQueryConnectionSpec
type BigQueryConnectionSpec struct {
	// The type of the BigQuery connection.
	// +kcc:proto:field=google.cloud.datacatalog.v1.BigQueryConnectionSpec.connection_type
	ConnectionType *string `json:"connectionType,omitempty"`

	// Specification for the BigQuery connection to a Cloud SQL instance.
	// +kcc:proto:field=google.cloud.datacatalog.v1.BigQueryConnectionSpec.cloud_sql
	CloudSQL *CloudSQLBigQueryConnectionSpec `json:"cloudSQL,omitempty"`

	// True if there are credentials attached to the BigQuery connection; false
	//  otherwise.
	// +kcc:proto:field=google.cloud.datacatalog.v1.BigQueryConnectionSpec.has_credential
	HasCredential *bool `json:"hasCredential,omitempty"`
}

// +kcc:proto=google.cloud.datacatalog.v1.BigQueryDateShardedSpec
type BigQueryDateShardedSpec struct {
}

// +kcc:proto=google.cloud.datacatalog.v1.BigQueryRoutineSpec
type BigQueryRoutineSpec struct {
	// Paths of the imported libraries.
	// +kcc:proto:field=google.cloud.datacatalog.v1.BigQueryRoutineSpec.imported_libraries
	ImportedLibraries []string `json:"importedLibraries,omitempty"`
}

// +kcc:proto=google.cloud.datacatalog.v1.BigQueryTableSpec
type BigQueryTableSpec struct {

	// Table view specification. Populated only if
	//  the `table_source_type` is `BIGQUERY_VIEW`.
	// +kcc:proto:field=google.cloud.datacatalog.v1.BigQueryTableSpec.view_spec
	ViewSpec *ViewSpec `json:"viewSpec,omitempty"`

	// Specification of a BigQuery table. Populated only if
	//  the `table_source_type` is `BIGQUERY_TABLE`.
	// +kcc:proto:field=google.cloud.datacatalog.v1.BigQueryTableSpec.table_spec
	TableSpec *TableSpec `json:"tableSpec,omitempty"`
}

// +kcc:proto=google.cloud.datacatalog.v1.BusinessContext
type BusinessContext struct {
	// Entry overview fields for rich text descriptions of entries.
	// +kcc:proto:field=google.cloud.datacatalog.v1.BusinessContext.entry_overview
	EntryOverview *EntryOverview `json:"entryOverview,omitempty"`

	// Contact people for the entry.
	// +kcc:proto:field=google.cloud.datacatalog.v1.BusinessContext.contacts
	Contacts *Contacts `json:"contacts,omitempty"`
}

// +kcc:proto=google.cloud.datacatalog.v1.CloudBigtableInstanceSpec
type CloudBigtableInstanceSpec struct {
	// The list of clusters for the Instance.
	// +kcc:proto:field=google.cloud.datacatalog.v1.CloudBigtableInstanceSpec.cloud_bigtable_cluster_specs
	CloudBigtableClusterSpecs []CloudBigtableInstanceSpec_CloudBigtableClusterSpec `json:"cloudBigtableClusterSpecs,omitempty"`
}

// +kcc:proto=google.cloud.datacatalog.v1.CloudBigtableInstanceSpec.CloudBigtableClusterSpec
type CloudBigtableInstanceSpec_CloudBigtableClusterSpec struct {
	// Name of the cluster.
	// +kcc:proto:field=google.cloud.datacatalog.v1.CloudBigtableInstanceSpec.CloudBigtableClusterSpec.display_name
	DisplayName *string `json:"displayName,omitempty"`

	// Location of the cluster, typically a Cloud zone.
	// +kcc:proto:field=google.cloud.datacatalog.v1.CloudBigtableInstanceSpec.CloudBigtableClusterSpec.location
	Location *string `json:"location,omitempty"`

	// Type of the resource. For a cluster this would be "CLUSTER".
	// +kcc:proto:field=google.cloud.datacatalog.v1.CloudBigtableInstanceSpec.CloudBigtableClusterSpec.type
	Type *string `json:"type,omitempty"`

	// A link back to the parent resource, in this case Instance.
	// +kcc:proto:field=google.cloud.datacatalog.v1.CloudBigtableInstanceSpec.CloudBigtableClusterSpec.linked_resource
	LinkedResource *string `json:"linkedResource,omitempty"`
}

// +kcc:proto=google.cloud.datacatalog.v1.CloudBigtableSystemSpec
type CloudBigtableSystemSpec struct {
	// Display name of the Instance. This is user specified and different from
	//  the resource name.
	// +kcc:proto:field=google.cloud.datacatalog.v1.CloudBigtableSystemSpec.instance_display_name
	InstanceDisplayName *string `json:"instanceDisplayName,omitempty"`
}

// +kcc:proto=google.cloud.datacatalog.v1.CloudSqlBigQueryConnectionSpec
type CloudSQLBigQueryConnectionSpec struct {
	// Cloud SQL instance ID in the format of `project:location:instance`.
	// +kcc:proto:field=google.cloud.datacatalog.v1.CloudSqlBigQueryConnectionSpec.instance_id
	InstanceID *string `json:"instanceID,omitempty"`

	// Database name.
	// +kcc:proto:field=google.cloud.datacatalog.v1.CloudSqlBigQueryConnectionSpec.database
	Database *string `json:"database,omitempty"`

	// Type of the Cloud SQL database.
	// +kcc:proto:field=google.cloud.datacatalog.v1.CloudSqlBigQueryConnectionSpec.type
	Type *string `json:"type,omitempty"`
}

// +kcc:proto=google.cloud.datacatalog.v1.ColumnSchema.FieldElementType
type ColumnSchema_FieldElementType struct {
	// Required. The type of a field element. See
	//  [ColumnSchema.type][google.cloud.datacatalog.v1.ColumnSchema.type].
	// +required
	// +kcc:proto:field=google.cloud.datacatalog.v1.ColumnSchema.FieldElementType.type
	Type *string `json:"type,omitempty"`
}

// +kcc:proto=google.cloud.datacatalog.v1.ColumnSchema.LookerColumnSpec
type ColumnSchema_LookerColumnSpec struct {
	// Looker specific column type of this column.
	// +kcc:proto:field=google.cloud.datacatalog.v1.ColumnSchema.LookerColumnSpec.type
	Type *string `json:"type,omitempty"`
}

// +kcc:proto=google.cloud.datacatalog.v1.CommonUsageStats
type CommonUsageStats struct {
	// View count in source system.
	// +kcc:proto:field=google.cloud.datacatalog.v1.CommonUsageStats.view_count
	ViewCount *int64 `json:"viewCount,omitempty"`
}

// +kcc:proto=google.cloud.datacatalog.v1.Contacts
type Contacts struct {
	// The list of contact people for the entry.
	// +kcc:proto:field=google.cloud.datacatalog.v1.Contacts.people
	People []Contacts_Person `json:"people,omitempty"`
}

// +kcc:proto=google.cloud.datacatalog.v1.Contacts.Person
type Contacts_Person struct {
	// Designation of the person, for example, Data Steward.
	// +kcc:proto:field=google.cloud.datacatalog.v1.Contacts.Person.designation
	Designation *string `json:"designation,omitempty"`

	// Email of the person in the format of `john.doe@xyz`,
	//  `<john.doe@xyz>`, or `John Doe<john.doe@xyz>`.
	// +kcc:proto:field=google.cloud.datacatalog.v1.Contacts.Person.email
	Email *string `json:"email,omitempty"`
}

// +kcc:proto=google.cloud.datacatalog.v1.DataSourceConnectionSpec
type DataSourceConnectionSpec struct {
	// Output only. Fields specific to BigQuery connections.
	// +kcc:proto:field=google.cloud.datacatalog.v1.DataSourceConnectionSpec.bigquery_connection_spec
	BigqueryConnectionSpec *BigQueryConnectionSpec `json:"bigqueryConnectionSpec,omitempty"`
}

// +kcc:proto=google.cloud.datacatalog.v1.DatabaseTableSpec
type DatabaseTableSpec struct {
	// Type of this table.
	// +kcc:proto:field=google.cloud.datacatalog.v1.DatabaseTableSpec.type
	Type *string `json:"type,omitempty"`

	// Spec what aplies to tables that are actually views.
	//  Not set for "real" tables.
	// +kcc:proto:field=google.cloud.datacatalog.v1.DatabaseTableSpec.database_view_spec
	DatabaseViewSpec *DatabaseTableSpec_DatabaseViewSpec `json:"databaseViewSpec,omitempty"`
}

// +kcc:proto=google.cloud.datacatalog.v1.DatabaseTableSpec.DatabaseViewSpec
type DatabaseTableSpec_DatabaseViewSpec struct {
	// Type of this view.
	// +kcc:proto:field=google.cloud.datacatalog.v1.DatabaseTableSpec.DatabaseViewSpec.view_type
	ViewType *string `json:"viewType,omitempty"`

	// Name of a singular table this view reflects one to one.
	// +kcc:proto:field=google.cloud.datacatalog.v1.DatabaseTableSpec.DatabaseViewSpec.base_table
	BaseTable *string `json:"baseTable,omitempty"`

	// SQL query used to generate this view.
	// +kcc:proto:field=google.cloud.datacatalog.v1.DatabaseTableSpec.DatabaseViewSpec.sql_query
	SQLQuery *string `json:"sqlQuery,omitempty"`
}

// +kcc:proto=google.cloud.datacatalog.v1.DataplexExternalTable
type DataplexExternalTable struct {
	// Service in which the external table is registered.
	// +kcc:proto:field=google.cloud.datacatalog.v1.DataplexExternalTable.system
	System *string `json:"system,omitempty"`

	// Fully qualified name (FQN) of the external table.
	// +kcc:proto:field=google.cloud.datacatalog.v1.DataplexExternalTable.fully_qualified_name
	FullyQualifiedName *string `json:"fullyQualifiedName,omitempty"`

	// Google Cloud resource name of the external table.
	// +kcc:proto:field=google.cloud.datacatalog.v1.DataplexExternalTable.google_cloud_resource
	GoogleCloudResource *string `json:"googleCloudResource,omitempty"`

	// Name of the Data Catalog entry representing the external table.
	// +kcc:proto:field=google.cloud.datacatalog.v1.DataplexExternalTable.data_catalog_entry
	DataCatalogEntry *string `json:"dataCatalogEntry,omitempty"`
}

// +kcc:proto=google.cloud.datacatalog.v1.DataplexFilesetSpec
type DataplexFilesetSpec struct {
	// Common Dataplex fields.
	// +kcc:proto:field=google.cloud.datacatalog.v1.DataplexFilesetSpec.dataplex_spec
	DataplexSpec *DataplexSpec `json:"dataplexSpec,omitempty"`
}

// +kcc:proto=google.cloud.datacatalog.v1.DataplexSpec
type DataplexSpec struct {
	// Fully qualified resource name of an asset in Dataplex, to which the
	//  underlying data source (Cloud Storage bucket or BigQuery dataset) of the
	//  entity is attached.
	// +kcc:proto:field=google.cloud.datacatalog.v1.DataplexSpec.asset
	Asset *string `json:"asset,omitempty"`

	// Format of the data.
	// +kcc:proto:field=google.cloud.datacatalog.v1.DataplexSpec.data_format
	DataFormat *PhysicalSchema `json:"dataFormat,omitempty"`

	// Compression format of the data, e.g., zip, gzip etc.
	// +kcc:proto:field=google.cloud.datacatalog.v1.DataplexSpec.compression_format
	CompressionFormat *string `json:"compressionFormat,omitempty"`

	// Project ID of the underlying Cloud Storage or BigQuery data. Note that
	//  this may not be the same project as the correspondingly Dataplex lake /
	//  zone / asset.
	// +kcc:proto:field=google.cloud.datacatalog.v1.DataplexSpec.project_id
	ProjectID *string `json:"projectID,omitempty"`
}

// +kcc:proto=google.cloud.datacatalog.v1.DataplexTableSpec
type DataplexTableSpec struct {
	// List of external tables registered by Dataplex in other systems based on
	//  the same underlying data.
	//
	//  External tables allow to query this data in those systems.
	// +kcc:proto:field=google.cloud.datacatalog.v1.DataplexTableSpec.external_tables
	ExternalTables []DataplexExternalTable `json:"externalTables,omitempty"`

	// Common Dataplex fields.
	// +kcc:proto:field=google.cloud.datacatalog.v1.DataplexTableSpec.dataplex_spec
	DataplexSpec *DataplexSpec `json:"dataplexSpec,omitempty"`

	// Indicates if the table schema is managed by the user or not.
	// +kcc:proto:field=google.cloud.datacatalog.v1.DataplexTableSpec.user_managed
	UserManaged *bool `json:"userManaged,omitempty"`
}

// +kcc:proto=google.cloud.datacatalog.v1.DatasetSpec
type DatasetSpec struct {
	// Vertex AI Dataset specific fields
	// +kcc:proto:field=google.cloud.datacatalog.v1.DatasetSpec.vertex_dataset_spec
	VertexDatasetSpec *VertexDatasetSpec `json:"vertexDatasetSpec,omitempty"`
}

// +kcc:proto=google.cloud.datacatalog.v1.EntryOverview
type EntryOverview struct {
	// Entry overview with support for rich text.
	//
	//  The overview must only contain Unicode characters, and should be
	//  formatted using HTML.
	//  The maximum length is 10 MiB as this value holds HTML descriptions
	//  including encoded images. The maximum length of the text without images
	//  is 100 KiB.
	// +kcc:proto:field=google.cloud.datacatalog.v1.EntryOverview.overview
	Overview *string `json:"overview,omitempty"`
}

// +kcc:proto=google.cloud.datacatalog.v1.FeatureOnlineStoreSpec
type FeatureOnlineStoreSpec struct {
}

// +kcc:proto=google.cloud.datacatalog.v1.FilesetSpec
type FilesetSpec struct {
	// Fields specific to a Dataplex fileset and present only in the Dataplex
	//  fileset entries.
	// +kcc:proto:field=google.cloud.datacatalog.v1.FilesetSpec.dataplex_fileset
	DataplexFileset *DataplexFilesetSpec `json:"dataplexFileset,omitempty"`
}

// +kcc:proto=google.cloud.datacatalog.v1.GcsFileSpec
type GCSFileSpec struct {
	// Required. Full file path. Example: `gs://bucket_name/a/b.txt`.
	// +required
	// +kcc:proto:field=google.cloud.datacatalog.v1.GcsFileSpec.file_path
	FilePath *string `json:"filePath,omitempty"`
}

// +kcc:proto=google.cloud.datacatalog.v1.GcsFilesetSpec
type GCSFilesetSpec struct {
	// Required. Patterns to identify a set of files in Google Cloud Storage.
	//
	//  For more information, see [Wildcard Names]
	//  (https://cloud.google.com/storage/docs/gsutil/addlhelp/WildcardNames).
	//
	//  Note: Currently, bucket wildcards are not supported.
	//
	//  Examples of valid `file_patterns`:
	//
	//   * `gs://bucket_name/dir/*`: matches all files in `bucket_name/dir`
	//                               directory
	//   * `gs://bucket_name/dir/**`: matches all files in `bucket_name/dir`
	//                                and all subdirectories
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
	//                               the `a/*/b` pattern, such as `a/c/b`, `a/d/b`
	//   * `gs://another_bucket/a.txt`: matches `gs://another_bucket/a.txt`
	//
	//  You can combine wildcards to match complex sets of files, for example:
	//
	//  `gs://bucket_name/[a-m]??.j*g`
	// +required
	// +kcc:proto:field=google.cloud.datacatalog.v1.GcsFilesetSpec.file_patterns
	FilePatterns []string `json:"filePatterns,omitempty"`
}

// +kcc:proto=google.cloud.datacatalog.v1.LookerSystemSpec
type LookerSystemSpec struct {
	// ID of the parent Looker Instance. Empty if it does not exist.
	//  Example value: `someinstance.looker.com`
	// +kcc:proto:field=google.cloud.datacatalog.v1.LookerSystemSpec.parent_instance_id
	ParentInstanceID *string `json:"parentInstanceID,omitempty"`

	// Name of the parent Looker Instance. Empty if it does not exist.
	// +kcc:proto:field=google.cloud.datacatalog.v1.LookerSystemSpec.parent_instance_display_name
	ParentInstanceDisplayName *string `json:"parentInstanceDisplayName,omitempty"`

	// ID of the parent Model. Empty if it does not exist.
	// +kcc:proto:field=google.cloud.datacatalog.v1.LookerSystemSpec.parent_model_id
	ParentModelID *string `json:"parentModelID,omitempty"`

	// Name of the parent Model. Empty if it does not exist.
	// +kcc:proto:field=google.cloud.datacatalog.v1.LookerSystemSpec.parent_model_display_name
	ParentModelDisplayName *string `json:"parentModelDisplayName,omitempty"`

	// ID of the parent View. Empty if it does not exist.
	// +kcc:proto:field=google.cloud.datacatalog.v1.LookerSystemSpec.parent_view_id
	ParentViewID *string `json:"parentViewID,omitempty"`

	// Name of the parent View. Empty if it does not exist.
	// +kcc:proto:field=google.cloud.datacatalog.v1.LookerSystemSpec.parent_view_display_name
	ParentViewDisplayName *string `json:"parentViewDisplayName,omitempty"`
}

// +kcc:proto=google.cloud.datacatalog.v1.ModelSpec
type ModelSpec struct {
	// Specification for vertex model resources.
	// +kcc:proto:field=google.cloud.datacatalog.v1.ModelSpec.vertex_model_spec
	VertexModelSpec *VertexModelSpec `json:"vertexModelSpec,omitempty"`
}

// +kcc:proto=google.cloud.datacatalog.v1.PersonalDetails
type PersonalDetails struct {
	// True if the entry is starred by the user; false otherwise.
	// +kcc:proto:field=google.cloud.datacatalog.v1.PersonalDetails.starred
	Starred *bool `json:"starred,omitempty"`

	// Set if the entry is starred; unset otherwise.
	// +kcc:proto:field=google.cloud.datacatalog.v1.PersonalDetails.star_time
	StarTime *string `json:"starTime,omitempty"`
}

// +kcc:proto=google.cloud.datacatalog.v1.PhysicalSchema
type PhysicalSchema struct {
	// Schema in Avro JSON format.
	// +kcc:proto:field=google.cloud.datacatalog.v1.PhysicalSchema.avro
	Avro *PhysicalSchema_AvroSchema `json:"avro,omitempty"`

	// Schema in Thrift format.
	// +kcc:proto:field=google.cloud.datacatalog.v1.PhysicalSchema.thrift
	Thrift *PhysicalSchema_ThriftSchema `json:"thrift,omitempty"`

	// Schema in protocol buffer format.
	// +kcc:proto:field=google.cloud.datacatalog.v1.PhysicalSchema.protobuf
	Protobuf *PhysicalSchema_ProtobufSchema `json:"protobuf,omitempty"`

	// Marks a Parquet-encoded data source.
	// +kcc:proto:field=google.cloud.datacatalog.v1.PhysicalSchema.parquet
	Parquet *PhysicalSchema_ParquetSchema `json:"parquet,omitempty"`

	// Marks an ORC-encoded data source.
	// +kcc:proto:field=google.cloud.datacatalog.v1.PhysicalSchema.orc
	Orc *PhysicalSchema_OrcSchema `json:"orc,omitempty"`

	// Marks a CSV-encoded data source.
	// +kcc:proto:field=google.cloud.datacatalog.v1.PhysicalSchema.csv
	Csv *PhysicalSchema_CsvSchema `json:"csv,omitempty"`
}

// +kcc:proto=google.cloud.datacatalog.v1.PhysicalSchema.AvroSchema
type PhysicalSchema_AvroSchema struct {
	// JSON source of the Avro schema.
	// +kcc:proto:field=google.cloud.datacatalog.v1.PhysicalSchema.AvroSchema.text
	Text *string `json:"text,omitempty"`
}

// +kcc:proto=google.cloud.datacatalog.v1.PhysicalSchema.CsvSchema
type PhysicalSchema_CsvSchema struct {
}

// +kcc:proto=google.cloud.datacatalog.v1.PhysicalSchema.OrcSchema
type PhysicalSchema_OrcSchema struct {
}

// +kcc:proto=google.cloud.datacatalog.v1.PhysicalSchema.ParquetSchema
type PhysicalSchema_ParquetSchema struct {
}

// +kcc:proto=google.cloud.datacatalog.v1.PhysicalSchema.ProtobufSchema
type PhysicalSchema_ProtobufSchema struct {
	// Protocol buffer source of the schema.
	// +kcc:proto:field=google.cloud.datacatalog.v1.PhysicalSchema.ProtobufSchema.text
	Text *string `json:"text,omitempty"`
}

// +kcc:proto=google.cloud.datacatalog.v1.PhysicalSchema.ThriftSchema
type PhysicalSchema_ThriftSchema struct {
	// Thrift IDL source of the schema.
	// +kcc:proto:field=google.cloud.datacatalog.v1.PhysicalSchema.ThriftSchema.text
	Text *string `json:"text,omitempty"`
}

// +kcc:proto=google.cloud.datacatalog.v1.RoutineSpec
type RoutineSpec struct {
	// The type of the routine.
	// +kcc:proto:field=google.cloud.datacatalog.v1.RoutineSpec.routine_type
	RoutineType *string `json:"routineType,omitempty"`

	// The language the routine is written in. The exact value depends on the
	//  source system. For BigQuery routines, possible values are:
	//
	//  * `SQL`
	//  * `JAVASCRIPT`
	// +kcc:proto:field=google.cloud.datacatalog.v1.RoutineSpec.language
	Language *string `json:"language,omitempty"`

	// Arguments of the routine.
	// +kcc:proto:field=google.cloud.datacatalog.v1.RoutineSpec.routine_arguments
	RoutineArguments []RoutineSpec_Argument `json:"routineArguments,omitempty"`

	// Return type of the argument. The exact value depends on the source system
	//  and the language.
	// +kcc:proto:field=google.cloud.datacatalog.v1.RoutineSpec.return_type
	ReturnType *string `json:"returnType,omitempty"`

	// The body of the routine.
	// +kcc:proto:field=google.cloud.datacatalog.v1.RoutineSpec.definition_body
	DefinitionBody *string `json:"definitionBody,omitempty"`

	// Fields specific for BigQuery routines.
	// +kcc:proto:field=google.cloud.datacatalog.v1.RoutineSpec.bigquery_routine_spec
	BigqueryRoutineSpec *BigQueryRoutineSpec `json:"bigqueryRoutineSpec,omitempty"`
}

// +kcc:proto=google.cloud.datacatalog.v1.RoutineSpec.Argument
type RoutineSpec_Argument struct {
	// The name of the argument. A return argument of a function might not have
	//  a name.
	// +kcc:proto:field=google.cloud.datacatalog.v1.RoutineSpec.Argument.name
	Name *string `json:"name,omitempty"`

	// Specifies whether the argument is input or output.
	// +kcc:proto:field=google.cloud.datacatalog.v1.RoutineSpec.Argument.mode
	Mode *string `json:"mode,omitempty"`

	// Type of the argument. The exact value depends on the source system and
	//  the language.
	// +kcc:proto:field=google.cloud.datacatalog.v1.RoutineSpec.Argument.type
	Type *string `json:"type,omitempty"`
}

// +kcc:proto=google.cloud.datacatalog.v1.Schema
type Schema struct {
	// The unified GoogleSQL-like schema of columns.
	//
	//  The overall maximum number of columns and nested columns is 10,000.
	//  The maximum nested depth is 15 levels.
	// +kcc:proto:field=google.cloud.datacatalog.v1.Schema.columns
	Columns []ColumnSchema `json:"columns,omitempty"`
}

// +kcc:proto=google.cloud.datacatalog.v1.ServiceSpec
type ServiceSpec struct {
	// Specification that applies to Instance entries of `CLOUD_BIGTABLE`
	//  system.
	// +kcc:proto:field=google.cloud.datacatalog.v1.ServiceSpec.cloud_bigtable_instance_spec
	CloudBigtableInstanceSpec *CloudBigtableInstanceSpec `json:"cloudBigtableInstanceSpec,omitempty"`
}

// +kcc:proto=google.cloud.datacatalog.v1.SqlDatabaseSystemSpec
type SQLDatabaseSystemSpec struct {
	// SQL Database Engine.
	//  enum SqlEngine {
	//   UNDEFINED = 0;
	//   MY_SQL = 1;
	//   POSTGRE_SQL = 2;
	//   SQL_SERVER = 3;
	//  }
	//  Engine of the enclosing database instance.
	// +kcc:proto:field=google.cloud.datacatalog.v1.SqlDatabaseSystemSpec.sql_engine
	SQLEngine *string `json:"sqlEngine,omitempty"`

	// Version of the database engine.
	// +kcc:proto:field=google.cloud.datacatalog.v1.SqlDatabaseSystemSpec.database_version
	DatabaseVersion *string `json:"databaseVersion,omitempty"`

	// Host of the SQL database
	//  enum InstanceHost {
	//   UNDEFINED = 0;
	//   SELF_HOSTED = 1;
	//   CLOUD_SQL = 2;
	//   AMAZON_RDS = 3;
	//   AZURE_SQL = 4;
	//  }
	//  Host of the enclousing database instance.
	// +kcc:proto:field=google.cloud.datacatalog.v1.SqlDatabaseSystemSpec.instance_host
	InstanceHost *string `json:"instanceHost,omitempty"`
}

// +kcc:proto=google.cloud.datacatalog.v1.StorageProperties
type StorageProperties struct {
	// Patterns to identify a set of files for this fileset.
	//
	//  Examples of a valid `file_pattern`:
	//
	//   * `gs://bucket_name/dir/*`: matches all files in the `bucket_name/dir`
	//                               directory
	//   * `gs://bucket_name/dir/**`: matches all files in the `bucket_name/dir`
	//                                and all subdirectories recursively
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
	//                               the `a/*/b` pattern, such as `a/c/b`, `a/d/b`
	//   * `gs://another_bucket/a.txt`: matches `gs://another_bucket/a.txt`
	// +kcc:proto:field=google.cloud.datacatalog.v1.StorageProperties.file_pattern
	FilePattern []string `json:"filePattern,omitempty"`

	// File type in MIME format, for example, `text/plain`.
	// +kcc:proto:field=google.cloud.datacatalog.v1.StorageProperties.file_type
	FileType *string `json:"fileType,omitempty"`
}

// +kcc:proto=google.cloud.datacatalog.v1.SystemTimestamps
type SystemTimestamps struct {
	// Creation timestamp of the resource within the given system.
	// +kcc:proto:field=google.cloud.datacatalog.v1.SystemTimestamps.create_time
	CreateTime *string `json:"createTime,omitempty"`

	// Timestamp of the last modification of the resource or its metadata within
	//  a given system.
	//
	//  Note: Depending on the source system, not every modification updates this
	//  timestamp.
	//  For example, BigQuery timestamps every metadata modification but not data
	//  or permission changes.
	// +kcc:proto:field=google.cloud.datacatalog.v1.SystemTimestamps.update_time
	UpdateTime *string `json:"updateTime,omitempty"`
}

// +kcc:proto=google.cloud.datacatalog.v1.TableSpec
type TableSpec struct {
}

// +kcc:proto=google.cloud.datacatalog.v1.UsageSignal
type UsageSignal struct {
	// The end timestamp of the duration of usage statistics.
	// +kcc:proto:field=google.cloud.datacatalog.v1.UsageSignal.update_time
	UpdateTime *string `json:"updateTime,omitempty"`

	// TODO: unsupported map type with key string and value message

	// Favorite count in the source system.
	// +kcc:proto:field=google.cloud.datacatalog.v1.UsageSignal.favorite_count
	FavoriteCount *int64 `json:"favoriteCount,omitempty"`
}

// +kcc:proto=google.cloud.datacatalog.v1.UsageStats
type UsageStats struct {
	// The number of successful uses of the underlying entry.
	// +kcc:proto:field=google.cloud.datacatalog.v1.UsageStats.total_completions
	TotalCompletions *float32 `json:"totalCompletions,omitempty"`

	// The number of failed attempts to use the underlying entry.
	// +kcc:proto:field=google.cloud.datacatalog.v1.UsageStats.total_failures
	TotalFailures *float32 `json:"totalFailures,omitempty"`

	// The number of cancelled attempts to use the underlying entry.
	// +kcc:proto:field=google.cloud.datacatalog.v1.UsageStats.total_cancellations
	TotalCancellations *float32 `json:"totalCancellations,omitempty"`

	// Total time spent only on successful uses, in milliseconds.
	// +kcc:proto:field=google.cloud.datacatalog.v1.UsageStats.total_execution_time_for_completions_millis
	TotalExecutionTimeForCompletionsMillis *float32 `json:"totalExecutionTimeForCompletionsMillis,omitempty"`
}

// +kcc:proto=google.cloud.datacatalog.v1.VertexDatasetSpec
type VertexDatasetSpec struct {
	// The number of DataItems in this Dataset. Only apply for non-structured
	//  Dataset.
	// +kcc:proto:field=google.cloud.datacatalog.v1.VertexDatasetSpec.data_item_count
	DataItemCount *int64 `json:"dataItemCount,omitempty"`

	// Type of the dataset.
	// +kcc:proto:field=google.cloud.datacatalog.v1.VertexDatasetSpec.data_type
	DataType *string `json:"dataType,omitempty"`
}

// +kcc:proto=google.cloud.datacatalog.v1.VertexModelSourceInfo
type VertexModelSourceInfo struct {
	// Type of the model source.
	// +kcc:proto:field=google.cloud.datacatalog.v1.VertexModelSourceInfo.source_type
	SourceType *string `json:"sourceType,omitempty"`

	// If this Model is copy of another Model. If true then
	//  [source_type][google.cloud.datacatalog.v1.VertexModelSourceInfo.source_type]
	//  pertains to the original.
	// +kcc:proto:field=google.cloud.datacatalog.v1.VertexModelSourceInfo.copy
	Copy *bool `json:"copy,omitempty"`
}

// +kcc:proto=google.cloud.datacatalog.v1.VertexModelSpec
type VertexModelSpec struct {
	// The version ID of the model.
	// +kcc:proto:field=google.cloud.datacatalog.v1.VertexModelSpec.version_id
	VersionID *string `json:"versionID,omitempty"`

	// User provided version aliases so that a model version can be referenced via
	//  alias
	// +kcc:proto:field=google.cloud.datacatalog.v1.VertexModelSpec.version_aliases
	VersionAliases []string `json:"versionAliases,omitempty"`

	// The description of this version.
	// +kcc:proto:field=google.cloud.datacatalog.v1.VertexModelSpec.version_description
	VersionDescription *string `json:"versionDescription,omitempty"`

	// Source of a Vertex model.
	// +kcc:proto:field=google.cloud.datacatalog.v1.VertexModelSpec.vertex_model_source_info
	VertexModelSourceInfo *VertexModelSourceInfo `json:"vertexModelSourceInfo,omitempty"`

	// URI of the Docker image to be used as the custom container for serving
	//  predictions.
	// +kcc:proto:field=google.cloud.datacatalog.v1.VertexModelSpec.container_image_uri
	ContainerImageURI *string `json:"containerImageURI,omitempty"`
}

// +kcc:proto=google.cloud.datacatalog.v1.ViewSpec
type ViewSpec struct {
}

// +kcc:proto=google.cloud.datacatalog.v1.BigQueryDateShardedSpec
type BigQueryDateShardedSpecObservedState struct {
	// Output only. The Data Catalog resource name of the dataset entry the
	//  current table belongs to. For example:
	//
	//  `projects/{PROJECT_ID}/locations/{LOCATION}/entrygroups/{ENTRY_GROUP_ID}/entries/{ENTRY_ID}`.
	// +kcc:proto:field=google.cloud.datacatalog.v1.BigQueryDateShardedSpec.dataset
	Dataset *string `json:"dataset,omitempty"`

	// Output only. The table name prefix of the shards.
	//
	//  The name of any given shard is `[table_prefix]YYYYMMDD`.
	//  For example, for the `MyTable20180101` shard, the
	//  `table_prefix` is `MyTable`.
	// +kcc:proto:field=google.cloud.datacatalog.v1.BigQueryDateShardedSpec.table_prefix
	TablePrefix *string `json:"tablePrefix,omitempty"`

	// Output only. Total number of shards.
	// +kcc:proto:field=google.cloud.datacatalog.v1.BigQueryDateShardedSpec.shard_count
	ShardCount *int64 `json:"shardCount,omitempty"`

	// Output only. BigQuery resource name of the latest shard.
	// +kcc:proto:field=google.cloud.datacatalog.v1.BigQueryDateShardedSpec.latest_shard_resource
	LatestShardResource *string `json:"latestShardResource,omitempty"`
}

// +kcc:proto=google.cloud.datacatalog.v1.BigQueryTableSpec
type BigQueryTableSpecObservedState struct {
	// Output only. The table source type.
	// +kcc:proto:field=google.cloud.datacatalog.v1.BigQueryTableSpec.table_source_type
	TableSourceType *string `json:"tableSourceType,omitempty"`

	// Table view specification. Populated only if
	//  the `table_source_type` is `BIGQUERY_VIEW`.
	// +kcc:proto:field=google.cloud.datacatalog.v1.BigQueryTableSpec.view_spec
	ViewSpec *ViewSpecObservedState `json:"viewSpec,omitempty"`

	// Specification of a BigQuery table. Populated only if
	//  the `table_source_type` is `BIGQUERY_TABLE`.
	// +kcc:proto:field=google.cloud.datacatalog.v1.BigQueryTableSpec.table_spec
	TableSpec *TableSpecObservedState `json:"tableSpec,omitempty"`
}

// +kcc:proto=google.cloud.datacatalog.v1.DataSource
type DataSource struct {
}

// +kcc:proto=google.cloud.datacatalog.v1.DatabaseTableSpec
type DatabaseTableSpecObservedState struct {
	// Output only. Fields specific to a Dataplex table and present only in the
	//  Dataplex table entries.
	// +kcc:proto:field=google.cloud.datacatalog.v1.DatabaseTableSpec.dataplex_table
	DataplexTable *DataplexTableSpec `json:"dataplexTable,omitempty"`
}

// +kcc:proto=google.cloud.datacatalog.v1.FeatureOnlineStoreSpec
type FeatureOnlineStoreSpecObservedState struct {
	// Output only. Type of underelaying storage for the FeatureOnlineStore.
	// +kcc:proto:field=google.cloud.datacatalog.v1.FeatureOnlineStoreSpec.storage_type
	StorageType *string `json:"storageType,omitempty"`
}

// +kcc:proto=google.cloud.datacatalog.v1.GcsFileSpec
type GCSFileSpecObservedState struct {
	// Output only. Creation, modification, and expiration timestamps of a Cloud
	//  Storage file.
	// +kcc:proto:field=google.cloud.datacatalog.v1.GcsFileSpec.gcs_timestamps
	GCSTimestamps *SystemTimestamps `json:"gcsTimestamps,omitempty"`

	// Output only. File size in bytes.
	// +kcc:proto:field=google.cloud.datacatalog.v1.GcsFileSpec.size_bytes
	SizeBytes *int64 `json:"sizeBytes,omitempty"`
}

// +kcc:proto=google.cloud.datacatalog.v1.GcsFilesetSpec
type GCSFilesetSpecObservedState struct {
	// Output only. Sample files contained in this fileset, not all files
	//  contained in this fileset are represented here.
	// +kcc:proto:field=google.cloud.datacatalog.v1.GcsFilesetSpec.sample_gcs_file_specs
	SampleGCSFileSpecs []GCSFileSpec `json:"sampleGCSFileSpecs,omitempty"`
}

// +kcc:proto=google.cloud.datacatalog.v1.SystemTimestamps
type SystemTimestampsObservedState struct {
	// Output only. Expiration timestamp of the resource within the given system.
	//
	//  Currently only applicable to BigQuery resources.
	// +kcc:proto:field=google.cloud.datacatalog.v1.SystemTimestamps.expire_time
	ExpireTime *string `json:"expireTime,omitempty"`
}

// +kcc:proto=google.cloud.datacatalog.v1.TableSpec
type TableSpecObservedState struct {
	// Output only. If the table is date-sharded, that is, it matches the
	//  `[prefix]YYYYMMDD` name pattern, this field is the Data Catalog resource
	//  name of the date-sharded grouped entry. For example:
	//
	//  `projects/{PROJECT_ID}/locations/{LOCATION}/entrygroups/{ENTRY_GROUP_ID}/entries/{ENTRY_ID}`.
	//
	//  Otherwise, `grouped_entry` is empty.
	// +kcc:proto:field=google.cloud.datacatalog.v1.TableSpec.grouped_entry
	GroupedEntry *string `json:"groupedEntry,omitempty"`
}

// +kcc:proto=google.cloud.datacatalog.v1.UsageSignal
type UsageSignalObservedState struct {

	// TODO: unsupported map type with key string and value message

}

// +kcc:proto=google.cloud.datacatalog.v1.ViewSpec
type ViewSpecObservedState struct {
	// Output only. The query that defines the table view.
	// +kcc:proto:field=google.cloud.datacatalog.v1.ViewSpec.view_query
	ViewQuery *string `json:"viewQuery,omitempty"`
}
