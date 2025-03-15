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

// +kcc:proto=google.cloud.bigquery.v2.AggregationThresholdPolicy
type AggregationThresholdPolicy struct {
	// Optional. The threshold for the "aggregation threshold" policy.
	// +kcc:proto:field=google.cloud.bigquery.v2.AggregationThresholdPolicy.threshold
	Threshold *int64 `json:"threshold,omitempty"`

	// Optional. The privacy unit column(s) associated with this policy.
	//  For now, only one column per data source object (table, view) is allowed as
	//  a privacy unit column.
	//  Representing as a repeated field in metadata for extensibility to
	//  multiple columns in future.
	//  Duplicates and Repeated struct fields are not allowed.
	//  For nested fields, use dot notation ("outer.inner")
	// +kcc:proto:field=google.cloud.bigquery.v2.AggregationThresholdPolicy.privacy_unit_columns
	PrivacyUnitColumns []string `json:"privacyUnitColumns,omitempty"`
}

// +kcc:proto=google.cloud.bigquery.v2.AvroOptions
type AvroOptions struct {
	// Optional. If sourceFormat is set to "AVRO", indicates whether to interpret
	//  logical types as the corresponding BigQuery data type (for example,
	//  TIMESTAMP), instead of using the raw type (for example, INTEGER).
	// +kcc:proto:field=google.cloud.bigquery.v2.AvroOptions.use_avro_logical_types
	UseAvroLogicalTypes *bool `json:"useAvroLogicalTypes,omitempty"`
}

// +kcc:proto=google.cloud.bigquery.v2.BigLakeConfiguration
type BigLakeConfiguration struct {
	// Optional. The connection specifying the credentials to be used to read and
	//  write to external storage, such as Cloud Storage. The connection_id can
	//  have the form `{project}.{location}.{connection_id}` or
	//  `projects/{project}/locations/{location}/connections/{connection_id}".
	// +kcc:proto:field=google.cloud.bigquery.v2.BigLakeConfiguration.connection_id
	ConnectionID *string `json:"connectionID,omitempty"`

	// Optional. The fully qualified location prefix of the external folder where
	//  table data is stored. The '*' wildcard character is not allowed. The URI
	//  should be in the format `gs://bucket/path_to_table/`
	// +kcc:proto:field=google.cloud.bigquery.v2.BigLakeConfiguration.storage_uri
	StorageURI *string `json:"storageURI,omitempty"`

	// Optional. The file format the table data is stored in.
	// +kcc:proto:field=google.cloud.bigquery.v2.BigLakeConfiguration.file_format
	FileFormat *string `json:"fileFormat,omitempty"`

	// Optional. The table format the metadata only snapshots are stored in.
	// +kcc:proto:field=google.cloud.bigquery.v2.BigLakeConfiguration.table_format
	TableFormat *string `json:"tableFormat,omitempty"`
}

// +kcc:proto=google.cloud.bigquery.v2.BigtableColumn
type BigtableColumn struct {
	// [Required] Qualifier of the column.
	//  Columns in the parent column family that has this exact qualifier are
	//  exposed as `<family field name>.<column field name>` field.
	//  If the qualifier is valid UTF-8 string, it can be specified in the
	//  qualifier_string field.  Otherwise, a base-64 encoded value must be set to
	//  qualifier_encoded.
	//  The column field name is the same as the column qualifier. However, if the
	//  qualifier is not a valid BigQuery field identifier i.e. does not match
	//  [a-zA-Z][a-zA-Z0-9_]*, a valid identifier must be provided as field_name.
	// +kcc:proto:field=google.cloud.bigquery.v2.BigtableColumn.qualifier_encoded
	QualifierEncoded *BytesValue `json:"qualifierEncoded,omitempty"`

	// Qualifier string.
	// +kcc:proto:field=google.cloud.bigquery.v2.BigtableColumn.qualifier_string
	QualifierString *string `json:"qualifierString,omitempty"`

	// Optional. If the qualifier is not a valid BigQuery field identifier i.e.
	//  does not match [a-zA-Z][a-zA-Z0-9_]*,  a valid identifier must be provided
	//  as the column field name and is used as field name in queries.
	// +kcc:proto:field=google.cloud.bigquery.v2.BigtableColumn.field_name
	FieldName *string `json:"fieldName,omitempty"`

	// Optional. The type to convert the value in cells of this column.
	//  The values are expected to be encoded using HBase Bytes.toBytes function
	//  when using the BINARY encoding value.
	//  Following BigQuery types are allowed (case-sensitive):
	//
	//  * BYTES
	//  * STRING
	//  * INTEGER
	//  * FLOAT
	//  * BOOLEAN
	//  * JSON
	//
	//  Default type is BYTES.
	//  'type' can also be set at the column family level. However, the setting at
	//  this level takes precedence if 'type' is set at both levels.
	// +kcc:proto:field=google.cloud.bigquery.v2.BigtableColumn.type
	Type *string `json:"type,omitempty"`

	// Optional. The encoding of the values when the type is not STRING.
	//  Acceptable encoding values are:
	//    TEXT - indicates values are alphanumeric text strings.
	//    BINARY - indicates values are encoded using HBase Bytes.toBytes family of
	//             functions.
	//  'encoding' can also be set at the column family level. However, the setting
	//  at this level takes precedence if 'encoding' is set at both levels.
	// +kcc:proto:field=google.cloud.bigquery.v2.BigtableColumn.encoding
	Encoding *string `json:"encoding,omitempty"`

	// Optional. If this is set, only the latest version of value in this column
	//              are exposed.
	//  'onlyReadLatest' can also be set at the column family level. However, the
	//  setting at this level takes precedence if 'onlyReadLatest' is set at both
	//  levels.
	// +kcc:proto:field=google.cloud.bigquery.v2.BigtableColumn.only_read_latest
	OnlyReadLatest *bool `json:"onlyReadLatest,omitempty"`
}

// +kcc:proto=google.cloud.bigquery.v2.BigtableColumnFamily
type BigtableColumnFamily struct {
	// Identifier of the column family.
	// +kcc:proto:field=google.cloud.bigquery.v2.BigtableColumnFamily.family_id
	FamilyID *string `json:"familyID,omitempty"`

	// Optional. The type to convert the value in cells of this column family.
	//  The values are expected to be encoded using HBase Bytes.toBytes function
	//  when using the BINARY encoding value.
	//  Following BigQuery types are allowed (case-sensitive):
	//
	//  * BYTES
	//  * STRING
	//  * INTEGER
	//  * FLOAT
	//  * BOOLEAN
	//  * JSON
	//
	//  Default type is BYTES.
	//  This can be overridden for a specific column by listing that column in
	//  'columns' and specifying a type for it.
	// +kcc:proto:field=google.cloud.bigquery.v2.BigtableColumnFamily.type
	Type *string `json:"type,omitempty"`

	// Optional. The encoding of the values when the type is not STRING.
	//  Acceptable encoding values are:
	//    TEXT - indicates values are alphanumeric text strings.
	//    BINARY - indicates values are encoded using HBase Bytes.toBytes family of
	//             functions.
	//  This can be overridden for a specific column by listing that column in
	//  'columns' and specifying an encoding for it.
	// +kcc:proto:field=google.cloud.bigquery.v2.BigtableColumnFamily.encoding
	Encoding *string `json:"encoding,omitempty"`

	// Optional. Lists of columns that should be exposed as individual fields as
	//  opposed to a list of (column name, value) pairs.
	//  All columns whose qualifier matches a qualifier in this list can be
	//  accessed as `<family field name>.<column field name>`.
	//  Other columns can be accessed as a list through
	//  the `<family field name>.Column` field.
	// +kcc:proto:field=google.cloud.bigquery.v2.BigtableColumnFamily.columns
	Columns []BigtableColumn `json:"columns,omitempty"`

	// Optional. If this is set only the latest version of value are exposed for
	//  all columns in this column family.
	//  This can be overridden for a specific column by listing that column in
	//  'columns' and specifying a different setting
	//  for that column.
	// +kcc:proto:field=google.cloud.bigquery.v2.BigtableColumnFamily.only_read_latest
	OnlyReadLatest *bool `json:"onlyReadLatest,omitempty"`
}

// +kcc:proto=google.cloud.bigquery.v2.BigtableOptions
type BigtableOptions struct {
	// Optional. List of column families to expose in the table schema along with
	//  their types.
	//  This list restricts the column families that can be referenced in queries
	//  and specifies their value types.
	//  You can use this list to do type conversions - see the 'type' field for
	//  more details.
	//  If you leave this list empty, all column families are present in the table
	//  schema and their values are read as BYTES.
	//  During a query only the column families referenced in that query are read
	//  from Bigtable.
	// +kcc:proto:field=google.cloud.bigquery.v2.BigtableOptions.column_families
	ColumnFamilies []BigtableColumnFamily `json:"columnFamilies,omitempty"`

	// Optional. If field is true, then the column families that are not
	//  specified in columnFamilies list are not exposed in the table schema.
	//  Otherwise, they are read with BYTES type values.
	//  The default value is false.
	// +kcc:proto:field=google.cloud.bigquery.v2.BigtableOptions.ignore_unspecified_column_families
	IgnoreUnspecifiedColumnFamilies *bool `json:"ignoreUnspecifiedColumnFamilies,omitempty"`

	// Optional. If field is true, then the rowkey column families will be read
	//  and converted to string. Otherwise they are read with BYTES type values and
	//  users need to manually cast them with CAST if necessary.
	//  The default value is false.
	// +kcc:proto:field=google.cloud.bigquery.v2.BigtableOptions.read_rowkey_as_string
	ReadRowkeyAsString *bool `json:"readRowkeyAsString,omitempty"`

	// Optional. If field is true, then each column family will be read as a
	//  single JSON column. Otherwise they are read as a repeated cell structure
	//  containing timestamp/value tuples. The default value is false.
	// +kcc:proto:field=google.cloud.bigquery.v2.BigtableOptions.output_column_families_as_json
	OutputColumnFamiliesAsJson *bool `json:"outputColumnFamiliesAsJson,omitempty"`
}

// +kcc:proto=google.cloud.bigquery.v2.CloneDefinition
type CloneDefinition struct {
	// Required. Reference describing the ID of the table that was cloned.
	// +kcc:proto:field=google.cloud.bigquery.v2.CloneDefinition.base_table_reference
	BaseTableReference *TableReference `json:"baseTableReference,omitempty"`

	// Required. The time at which the base table was cloned. This value is
	//  reported in the JSON response using RFC3339 format.
	// +kcc:proto:field=google.cloud.bigquery.v2.CloneDefinition.clone_time
	CloneTime *string `json:"cloneTime,omitempty"`
}

// +kcc:proto=google.cloud.bigquery.v2.Clustering
type Clustering struct {
	// One or more fields on which data should be clustered. Only top-level,
	//  non-repeated, simple-type fields are supported. The ordering of the
	//  clustering fields should be prioritized from most to least important
	//  for filtering purposes.
	//
	//  Additional information on limitations can be found here:
	//  https://cloud.google.com/bigquery/docs/creating-clustered-tables#limitations
	// +kcc:proto:field=google.cloud.bigquery.v2.Clustering.fields
	Fields []string `json:"fields,omitempty"`
}

// +kcc:proto=google.cloud.bigquery.v2.ColumnReference
type ColumnReference struct {
	// Required. The column that composes the foreign key.
	// +kcc:proto:field=google.cloud.bigquery.v2.ColumnReference.referencing_column
	ReferencingColumn *string `json:"referencingColumn,omitempty"`

	// Required. The column in the primary key that are referenced by the
	//  referencing_column.
	// +kcc:proto:field=google.cloud.bigquery.v2.ColumnReference.referenced_column
	ReferencedColumn *string `json:"referencedColumn,omitempty"`
}

// +kcc:proto=google.cloud.bigquery.v2.CsvOptions
type CsvOptions struct {
	// Optional. The separator character for fields in a CSV file. The separator
	//  is interpreted as a single byte. For files encoded in ISO-8859-1, any
	//  single character can be used as a separator. For files encoded in UTF-8,
	//  characters represented in decimal range 1-127 (U+0001-U+007F) can be used
	//  without any modification. UTF-8 characters encoded with multiple bytes
	//  (i.e. U+0080 and above) will have only the first byte used for separating
	//  fields. The remaining bytes will be treated as a part of the field.
	//  BigQuery also supports the escape sequence "\t" (U+0009) to specify a tab
	//  separator. The default value is comma (",", U+002C).
	// +kcc:proto:field=google.cloud.bigquery.v2.CsvOptions.field_delimiter
	FieldDelimiter *string `json:"fieldDelimiter,omitempty"`

	// Optional. The number of rows at the top of a CSV file that BigQuery will
	//  skip when reading the data. The default value is 0. This property is
	//  useful if you have header rows in the file that should be skipped.
	//  When autodetect is on, the behavior is the following:
	//
	//  * skipLeadingRows unspecified - Autodetect tries to detect headers in the
	//    first row. If they are not detected, the row is read as data. Otherwise
	//    data is read starting from the second row.
	//  * skipLeadingRows is 0 - Instructs autodetect that there are no headers and
	//    data should be read starting from the first row.
	//  * skipLeadingRows = N > 0 - Autodetect skips N-1 rows and tries to detect
	//    headers in row N. If headers are not detected, row N is just skipped.
	//    Otherwise row N is used to extract column names for the detected schema.
	// +kcc:proto:field=google.cloud.bigquery.v2.CsvOptions.skip_leading_rows
	SkipLeadingRows *int64 `json:"skipLeadingRows,omitempty"`

	// Optional. The value that is used to quote data sections in a CSV file.
	//  BigQuery converts the string to ISO-8859-1 encoding, and then uses the
	//  first byte of the encoded string to split the data in its raw, binary
	//  state.
	//  The default value is a double-quote (").
	//  If your data does not contain quoted sections,
	//  set the property value to an empty string.
	//  If your data contains quoted newline characters, you must also set the
	//  allowQuotedNewlines property to true.
	//  To include the specific quote character within a quoted value, precede it
	//  with an additional matching quote character. For example, if you want to
	//  escape the default character  ' " ', use ' "" '.
	// +kcc:proto:field=google.cloud.bigquery.v2.CsvOptions.quote
	Quote *string `json:"quote,omitempty"`

	// Optional. Indicates if BigQuery should allow quoted data sections that
	//  contain newline characters in a CSV file. The default value is false.
	// +kcc:proto:field=google.cloud.bigquery.v2.CsvOptions.allow_quoted_newlines
	AllowQuotedNewlines *bool `json:"allowQuotedNewlines,omitempty"`

	// Optional. Indicates if BigQuery should accept rows that are missing
	//  trailing optional columns. If true, BigQuery treats missing trailing
	//  columns as null values.
	//  If false, records with missing trailing columns are treated as bad records,
	//  and if there are too many bad records, an invalid error is returned in the
	//  job result. The default value is false.
	// +kcc:proto:field=google.cloud.bigquery.v2.CsvOptions.allow_jagged_rows
	AllowJaggedRows *bool `json:"allowJaggedRows,omitempty"`

	// Optional. The character encoding of the data.
	//  The supported values are UTF-8, ISO-8859-1, UTF-16BE, UTF-16LE, UTF-32BE,
	//  and UTF-32LE.  The default value is UTF-8.
	//  BigQuery decodes the data after the raw, binary data has been split using
	//  the values of the quote and fieldDelimiter properties.
	// +kcc:proto:field=google.cloud.bigquery.v2.CsvOptions.encoding
	Encoding *string `json:"encoding,omitempty"`

	// Optional. Indicates if the embedded ASCII control characters (the first 32
	//  characters in the ASCII-table, from '\x00' to '\x1F') are preserved.
	// +kcc:proto:field=google.cloud.bigquery.v2.CsvOptions.preserve_ascii_control_characters
	PreserveAsciiControlCharacters *bool `json:"preserveAsciiControlCharacters,omitempty"`

	// Optional. Specifies a string that represents a null value in a CSV file.
	//  For example, if you specify "\N", BigQuery interprets "\N" as a null value
	//  when querying a CSV file.
	//  The default value is the empty string. If you set this property to a custom
	//  value, BigQuery throws an error if an empty string is present for all data
	//  types except for STRING and BYTE. For STRING and BYTE columns, BigQuery
	//  interprets the empty string as an empty value.
	// +kcc:proto:field=google.cloud.bigquery.v2.CsvOptions.null_marker
	NullMarker *string `json:"nullMarker,omitempty"`
}

// +kcc:proto=google.cloud.bigquery.v2.DataPolicyOption
type DataPolicyOption struct {
	// Data policy resource name in the form of
	//  projects/project_id/locations/location_id/dataPolicies/data_policy_id.
	// +kcc:proto:field=google.cloud.bigquery.v2.DataPolicyOption.name
	Name *string `json:"name,omitempty"`
}

// +kcc:proto=google.cloud.bigquery.v2.DifferentialPrivacyPolicy
type DifferentialPrivacyPolicy struct {
	// Optional. The maximum epsilon value that a query can consume. If the
	//  subscriber specifies epsilon as a parameter in a SELECT query, it must be
	//  less than or equal to this value. The epsilon parameter controls the amount
	//  of noise that is added to the groups — a higher epsilon means less noise.
	// +kcc:proto:field=google.cloud.bigquery.v2.DifferentialPrivacyPolicy.max_epsilon_per_query
	MaxEpsilonPerQuery *float64 `json:"maxEpsilonPerQuery,omitempty"`

	// Optional. The delta value that is used per query. Delta represents the
	//  probability that any row will fail to be epsilon differentially private.
	//  Indicates the risk associated with exposing aggregate rows in the result of
	//  a query.
	// +kcc:proto:field=google.cloud.bigquery.v2.DifferentialPrivacyPolicy.delta_per_query
	DeltaPerQuery *float64 `json:"deltaPerQuery,omitempty"`

	// Optional. The maximum groups contributed value that is used per query.
	//  Represents the maximum number of groups to which each protected entity can
	//  contribute. Changing this value does not improve or worsen privacy. The
	//  best value for accuracy and utility depends on the query and data.
	// +kcc:proto:field=google.cloud.bigquery.v2.DifferentialPrivacyPolicy.max_groups_contributed
	MaxGroupsContributed *int64 `json:"maxGroupsContributed,omitempty"`

	// Optional. The privacy unit column associated with this policy. Differential
	//  privacy policies can only have one privacy unit column per data source
	//  object (table, view).
	// +kcc:proto:field=google.cloud.bigquery.v2.DifferentialPrivacyPolicy.privacy_unit_column
	PrivacyUnitColumn *string `json:"privacyUnitColumn,omitempty"`

	// Optional. The total epsilon budget for all queries against the
	//  privacy-protected view. Each subscriber query against this view charges the
	//  amount of epsilon they request in their query. If there is sufficient
	//  budget, then the subscriber query attempts to complete. It might still fail
	//  due to other reasons, in which case the charge is refunded. If there is
	//  insufficient budget the query is rejected. There might be multiple charge
	//  attempts if a single query references multiple views. In this case there
	//  must be sufficient budget for all charges or the query is rejected and
	//  charges are refunded in best effort. The budget does not have a refresh
	//  policy and can only be updated via ALTER VIEW or circumvented by creating a
	//  new view that can be queried with a fresh budget.
	// +kcc:proto:field=google.cloud.bigquery.v2.DifferentialPrivacyPolicy.epsilon_budget
	EpsilonBudget *float64 `json:"epsilonBudget,omitempty"`

	// Optional. The total delta budget for all queries against the
	//  privacy-protected view. Each subscriber query against this view charges the
	//  amount of delta that is pre-defined by the contributor through the privacy
	//  policy delta_per_query field. If there is sufficient budget, then the
	//  subscriber query attempts to complete. It might still fail due to other
	//  reasons, in which case the charge is refunded. If there is insufficient
	//  budget the query is rejected. There might be multiple charge attempts if a
	//  single query references multiple views. In this case there must be
	//  sufficient budget for all charges or the query is rejected and charges are
	//  refunded in best effort. The budget does not have a refresh policy and can
	//  only be updated via ALTER VIEW or circumvented by creating a new view that
	//  can be queried with a fresh budget.
	// +kcc:proto:field=google.cloud.bigquery.v2.DifferentialPrivacyPolicy.delta_budget
	DeltaBudget *float64 `json:"deltaBudget,omitempty"`
}

// +kcc:proto=google.cloud.bigquery.v2.EncryptionConfiguration
type EncryptionConfiguration struct {
	// Optional. Describes the Cloud KMS encryption key that will be used to
	//  protect destination BigQuery table. The BigQuery Service Account associated
	//  with your project requires access to this encryption key.
	// +kcc:proto:field=google.cloud.bigquery.v2.EncryptionConfiguration.kms_key_name
	KMSKeyName *string `json:"kmsKeyName,omitempty"`
}

// +kcc:proto=google.cloud.bigquery.v2.ErrorProto
type ErrorProto struct {
	// A short error code that summarizes the error.
	// +kcc:proto:field=google.cloud.bigquery.v2.ErrorProto.reason
	Reason *string `json:"reason,omitempty"`

	// Specifies where the error occurred, if present.
	// +kcc:proto:field=google.cloud.bigquery.v2.ErrorProto.location
	Location *string `json:"location,omitempty"`

	// Debugging information. This property is internal to Google and should not
	//  be used.
	// +kcc:proto:field=google.cloud.bigquery.v2.ErrorProto.debug_info
	DebugInfo *string `json:"debugInfo,omitempty"`

	// A human-readable description of the error.
	// +kcc:proto:field=google.cloud.bigquery.v2.ErrorProto.message
	Message *string `json:"message,omitempty"`
}

// +kcc:proto=google.cloud.bigquery.v2.ExternalCatalogTableOptions
type ExternalCatalogTableOptions struct {
	// Optional. A map of key value pairs defining the parameters and properties
	//  of the open source table. Corresponds with hive meta store table
	//  parameters. Maximum size of 4Mib.
	// +kcc:proto:field=google.cloud.bigquery.v2.ExternalCatalogTableOptions.parameters
	Parameters map[string]string `json:"parameters,omitempty"`

	// Optional. A storage descriptor containing information about the physical
	//  storage of this table.
	// +kcc:proto:field=google.cloud.bigquery.v2.ExternalCatalogTableOptions.storage_descriptor
	StorageDescriptor *StorageDescriptor `json:"storageDescriptor,omitempty"`

	// Optional. The connection specifying the credentials to be used to read
	//  external storage, such as Azure Blob, Cloud Storage, or S3. The connection
	//  is needed to read the open source table from BigQuery Engine. The
	//  connection_id can have the form
	//  `<project_id>.<location_id>.<connection_id>` or
	//  `projects/<project_id>/locations/<location_id>/connections/<connection_id>`.
	// +kcc:proto:field=google.cloud.bigquery.v2.ExternalCatalogTableOptions.connection_id
	ConnectionID *string `json:"connectionID,omitempty"`
}

// +kcc:proto=google.cloud.bigquery.v2.ExternalDataConfiguration
type ExternalDataConfiguration struct {
	// [Required] The fully-qualified URIs that point to your data in Google
	//  Cloud. For Google Cloud Storage URIs:
	//    Each URI can contain one '*' wildcard character and it must come after
	//    the 'bucket' name.
	//    Size limits related to load jobs apply to external data sources.
	//  For Google Cloud Bigtable URIs:
	//    Exactly one URI can be specified and it has be a fully specified and
	//    valid HTTPS URL for a Google Cloud Bigtable table.
	//  For Google Cloud Datastore backups, exactly one URI can be specified. Also,
	//  the '*' wildcard character is not allowed.
	// +kcc:proto:field=google.cloud.bigquery.v2.ExternalDataConfiguration.source_uris
	SourceUris []string `json:"sourceUris,omitempty"`

	// Optional. Specifies how source URIs are interpreted for constructing the
	//  file set to load.  By default source URIs are expanded against the
	//  underlying storage.  Other options include specifying manifest files. Only
	//  applicable to object storage systems.
	// +kcc:proto:field=google.cloud.bigquery.v2.ExternalDataConfiguration.file_set_spec_type
	FileSetSpecType *string `json:"fileSetSpecType,omitempty"`

	// Optional. The schema for the data.
	//  Schema is required for CSV and JSON formats if autodetect is not on.
	//  Schema is disallowed for Google Cloud Bigtable, Cloud Datastore backups,
	//  Avro, ORC and Parquet formats.
	// +kcc:proto:field=google.cloud.bigquery.v2.ExternalDataConfiguration.schema
	Schema *TableSchema `json:"schema,omitempty"`

	// [Required] The data format.
	//  For CSV files, specify "CSV".
	//  For Google sheets, specify "GOOGLE_SHEETS".
	//  For newline-delimited JSON, specify "NEWLINE_DELIMITED_JSON".
	//  For Avro files, specify "AVRO".
	//  For Google Cloud Datastore backups, specify "DATASTORE_BACKUP".
	//  For Apache Iceberg tables, specify "ICEBERG".
	//  For ORC files, specify "ORC".
	//  For Parquet files, specify "PARQUET".
	//  [Beta] For Google Cloud Bigtable, specify "BIGTABLE".
	// +kcc:proto:field=google.cloud.bigquery.v2.ExternalDataConfiguration.source_format
	SourceFormat *string `json:"sourceFormat,omitempty"`

	// Optional. The maximum number of bad records that BigQuery can ignore when
	//  reading data. If the number of bad records exceeds this value, an invalid
	//  error is returned in the job result. The default value is 0, which requires
	//  that all records are valid. This setting is ignored for Google Cloud
	//  Bigtable, Google Cloud Datastore backups, Avro, ORC and Parquet formats.
	// +kcc:proto:field=google.cloud.bigquery.v2.ExternalDataConfiguration.max_bad_records
	MaxBadRecords *Int32Value `json:"maxBadRecords,omitempty"`

	// Try to detect schema and format options automatically.
	//  Any option specified explicitly will be honored.
	// +kcc:proto:field=google.cloud.bigquery.v2.ExternalDataConfiguration.autodetect
	Autodetect *bool `json:"autodetect,omitempty"`

	// Optional. Indicates if BigQuery should allow extra values that are not
	//  represented in the table schema.
	//  If true, the extra values are ignored.
	//  If false, records with extra columns are treated as bad records, and if
	//  there are too many bad records, an invalid error is returned in the job
	//  result.
	//  The default value is false.
	//  The sourceFormat property determines what BigQuery treats as an extra
	//  value:
	//    CSV: Trailing columns
	//    JSON: Named values that don't match any column names
	//    Google Cloud Bigtable: This setting is ignored.
	//    Google Cloud Datastore backups: This setting is ignored.
	//    Avro: This setting is ignored.
	//    ORC: This setting is ignored.
	//    Parquet: This setting is ignored.
	// +kcc:proto:field=google.cloud.bigquery.v2.ExternalDataConfiguration.ignore_unknown_values
	IgnoreUnknownValues *bool `json:"ignoreUnknownValues,omitempty"`

	// Optional. The compression type of the data source.
	//  Possible values include GZIP and NONE. The default value is NONE.
	//  This setting is ignored for Google Cloud Bigtable, Google Cloud Datastore
	//  backups, Avro, ORC and Parquet
	//  formats. An empty string is an invalid value.
	// +kcc:proto:field=google.cloud.bigquery.v2.ExternalDataConfiguration.compression
	Compression *string `json:"compression,omitempty"`

	// Optional. Additional properties to set if sourceFormat is set to CSV.
	// +kcc:proto:field=google.cloud.bigquery.v2.ExternalDataConfiguration.csv_options
	CsvOptions *CsvOptions `json:"csvOptions,omitempty"`

	// Optional. Additional properties to set if sourceFormat is set to JSON.
	// +kcc:proto:field=google.cloud.bigquery.v2.ExternalDataConfiguration.json_options
	JsonOptions *JsonOptions `json:"jsonOptions,omitempty"`

	// Optional. Additional options if sourceFormat is set to BIGTABLE.
	// +kcc:proto:field=google.cloud.bigquery.v2.ExternalDataConfiguration.bigtable_options
	BigtableOptions *BigtableOptions `json:"bigtableOptions,omitempty"`

	// Optional. Additional options if sourceFormat is set to GOOGLE_SHEETS.
	// +kcc:proto:field=google.cloud.bigquery.v2.ExternalDataConfiguration.google_sheets_options
	GoogleSheetsOptions *GoogleSheetsOptions `json:"googleSheetsOptions,omitempty"`

	// Optional. When set, configures hive partitioning support. Not all storage
	//  formats support hive partitioning -- requesting hive partitioning on an
	//  unsupported format will lead to an error, as will providing an invalid
	//  specification.
	// +kcc:proto:field=google.cloud.bigquery.v2.ExternalDataConfiguration.hive_partitioning_options
	HivePartitioningOptions *HivePartitioningOptions `json:"hivePartitioningOptions,omitempty"`

	// Optional. The connection specifying the credentials to be used to read
	//  external storage, such as Azure Blob, Cloud Storage, or S3. The
	//  connection_id can have the form
	//  `{project_id}.{location_id};{connection_id}` or
	//  `projects/{project_id}/locations/{location_id}/connections/{connection_id}`.
	// +kcc:proto:field=google.cloud.bigquery.v2.ExternalDataConfiguration.connection_id
	ConnectionID *string `json:"connectionID,omitempty"`

	// Defines the list of possible SQL data types to which the source decimal
	//  values are converted. This list and the precision and the scale parameters
	//  of the decimal field determine the target type. In the order of NUMERIC,
	//  BIGNUMERIC, and STRING, a
	//  type is picked if it is in the specified list and if it supports the
	//  precision and the scale. STRING supports all precision and scale values.
	//  If none of the listed types supports the precision and the scale, the type
	//  supporting the widest range in the specified list is picked, and if a value
	//  exceeds the supported range when reading the data, an error will be thrown.
	//
	//  Example: Suppose the value of this field is ["NUMERIC", "BIGNUMERIC"].
	//  If (precision,scale) is:
	//
	//  * (38,9) -> NUMERIC;
	//  * (39,9) -> BIGNUMERIC (NUMERIC cannot hold 30 integer digits);
	//  * (38,10) -> BIGNUMERIC (NUMERIC cannot hold 10 fractional digits);
	//  * (76,38) -> BIGNUMERIC;
	//  * (77,38) -> BIGNUMERIC (error if value exeeds supported range).
	//
	//  This field cannot contain duplicate types. The order of the types in this
	//  field is ignored. For example, ["BIGNUMERIC", "NUMERIC"] is the same as
	//  ["NUMERIC", "BIGNUMERIC"] and NUMERIC always takes precedence over
	//  BIGNUMERIC.
	//
	//  Defaults to ["NUMERIC", "STRING"] for ORC and ["NUMERIC"] for the other
	//  file formats.
	// +kcc:proto:field=google.cloud.bigquery.v2.ExternalDataConfiguration.decimal_target_types
	DecimalTargetTypes []string `json:"decimalTargetTypes,omitempty"`

	// Optional. Additional properties to set if sourceFormat is set to AVRO.
	// +kcc:proto:field=google.cloud.bigquery.v2.ExternalDataConfiguration.avro_options
	AvroOptions *AvroOptions `json:"avroOptions,omitempty"`

	// Optional. Load option to be used together with source_format
	//  newline-delimited JSON to indicate that a variant of JSON is being loaded.
	//  To load newline-delimited GeoJSON, specify GEOJSON (and source_format must
	//  be set to NEWLINE_DELIMITED_JSON).
	// +kcc:proto:field=google.cloud.bigquery.v2.ExternalDataConfiguration.json_extension
	JsonExtension *string `json:"jsonExtension,omitempty"`

	// Optional. Additional properties to set if sourceFormat is set to PARQUET.
	// +kcc:proto:field=google.cloud.bigquery.v2.ExternalDataConfiguration.parquet_options
	ParquetOptions *ParquetOptions `json:"parquetOptions,omitempty"`

	// Optional. ObjectMetadata is used to create Object Tables. Object Tables
	//  contain a listing of objects (with their metadata) found at the
	//  source_uris. If ObjectMetadata is set, source_format should be omitted.
	//
	//  Currently SIMPLE is the only supported Object Metadata type.
	// +kcc:proto:field=google.cloud.bigquery.v2.ExternalDataConfiguration.object_metadata
	ObjectMetadata *string `json:"objectMetadata,omitempty"`

	// Optional. When creating an external table, the user can provide a reference
	//  file with the table schema. This is enabled for the following formats:
	//  AVRO, PARQUET, ORC.
	// +kcc:proto:field=google.cloud.bigquery.v2.ExternalDataConfiguration.reference_file_schema_uri
	ReferenceFileSchemaURI *string `json:"referenceFileSchemaURI,omitempty"`

	// Optional. Metadata Cache Mode for the table. Set this to enable caching of
	//  metadata from external data source.
	// +kcc:proto:field=google.cloud.bigquery.v2.ExternalDataConfiguration.metadata_cache_mode
	MetadataCacheMode *string `json:"metadataCacheMode,omitempty"`
}

// +kcc:proto=google.cloud.bigquery.v2.ForeignKey
type ForeignKey struct {
	// Optional. Set only if the foreign key constraint is named.
	// +kcc:proto:field=google.cloud.bigquery.v2.ForeignKey.name
	Name *string `json:"name,omitempty"`

	// Required. The table that holds the primary key and is referenced by this
	//  foreign key.
	// +kcc:proto:field=google.cloud.bigquery.v2.ForeignKey.referenced_table
	ReferencedTable *TableReference `json:"referencedTable,omitempty"`

	// Required. The columns that compose the foreign key.
	// +kcc:proto:field=google.cloud.bigquery.v2.ForeignKey.column_references
	ColumnReferences []ColumnReference `json:"columnReferences,omitempty"`
}

// +kcc:proto=google.cloud.bigquery.v2.ForeignTypeInfo
type ForeignTypeInfo struct {
	// Required. Specifies the system which defines the foreign data type.
	// +kcc:proto:field=google.cloud.bigquery.v2.ForeignTypeInfo.type_system
	TypeSystem *string `json:"typeSystem,omitempty"`
}

// +kcc:proto=google.cloud.bigquery.v2.ForeignViewDefinition
type ForeignViewDefinition struct {
	// Required. The query that defines the view.
	// +kcc:proto:field=google.cloud.bigquery.v2.ForeignViewDefinition.query
	Query *string `json:"query,omitempty"`

	// Optional. Represents the dialect of the query.
	// +kcc:proto:field=google.cloud.bigquery.v2.ForeignViewDefinition.dialect
	Dialect *string `json:"dialect,omitempty"`
}

// +kcc:proto=google.cloud.bigquery.v2.GoogleSheetsOptions
type GoogleSheetsOptions struct {
	// Optional. The number of rows at the top of a sheet that BigQuery will skip
	//  when reading the data. The default value is 0. This property is useful if
	//  you have header rows that should be skipped. When autodetect is on,
	//  the behavior is the following:
	//  * skipLeadingRows unspecified - Autodetect tries to detect headers in the
	//    first row. If they are not detected, the row is read as data. Otherwise
	//    data is read starting from the second row.
	//  * skipLeadingRows is 0 - Instructs autodetect that there are no headers and
	//    data should be read starting from the first row.
	//  * skipLeadingRows = N > 0 - Autodetect skips N-1 rows and tries to detect
	//    headers in row N. If headers are not detected, row N is just skipped.
	//    Otherwise row N is used to extract column names for the detected schema.
	// +kcc:proto:field=google.cloud.bigquery.v2.GoogleSheetsOptions.skip_leading_rows
	SkipLeadingRows *int64 `json:"skipLeadingRows,omitempty"`

	// Optional. Range of a sheet to query from. Only used when non-empty.
	//  Typical format: sheet_name!top_left_cell_id:bottom_right_cell_id
	//  For example: sheet1!A1:B20
	// +kcc:proto:field=google.cloud.bigquery.v2.GoogleSheetsOptions.range
	Range *string `json:"range,omitempty"`
}

// +kcc:proto=google.cloud.bigquery.v2.HivePartitioningOptions
type HivePartitioningOptions struct {
	// Optional. When set, what mode of hive partitioning to use when reading
	//  data.  The following modes are supported:
	//
	//  * AUTO: automatically infer partition key name(s) and type(s).
	//
	//  * STRINGS: automatically infer partition key name(s).  All types are
	//  strings.
	//
	//  * CUSTOM: partition key schema is encoded in the source URI prefix.
	//
	//  Not all storage formats support hive partitioning. Requesting hive
	//  partitioning on an unsupported format will lead to an error.
	//  Currently supported formats are: JSON, CSV, ORC, Avro and Parquet.
	// +kcc:proto:field=google.cloud.bigquery.v2.HivePartitioningOptions.mode
	Mode *string `json:"mode,omitempty"`

	// Optional. When hive partition detection is requested, a common prefix for
	//  all source uris must be required.  The prefix must end immediately before
	//  the partition key encoding begins. For example, consider files following
	//  this data layout:
	//
	//  gs://bucket/path_to_table/dt=2019-06-01/country=USA/id=7/file.avro
	//
	//  gs://bucket/path_to_table/dt=2019-05-31/country=CA/id=3/file.avro
	//
	//  When hive partitioning is requested with either AUTO or STRINGS detection,
	//  the common prefix can be either of gs://bucket/path_to_table or
	//  gs://bucket/path_to_table/.
	//
	//  CUSTOM detection requires encoding the partitioning schema immediately
	//  after the common prefix.  For CUSTOM, any of
	//
	//  * gs://bucket/path_to_table/{dt:DATE}/{country:STRING}/{id:INTEGER}
	//
	//  * gs://bucket/path_to_table/{dt:STRING}/{country:STRING}/{id:INTEGER}
	//
	//  * gs://bucket/path_to_table/{dt:DATE}/{country:STRING}/{id:STRING}
	//
	//  would all be valid source URI prefixes.
	// +kcc:proto:field=google.cloud.bigquery.v2.HivePartitioningOptions.source_uri_prefix
	SourceURIPrefix *string `json:"sourceURIPrefix,omitempty"`

	// Optional. If set to true, queries over this table require a partition
	//  filter that can be used for partition elimination to be specified.
	//
	//  Note that this field should only be true when creating a permanent
	//  external table or querying a temporary external table.
	//
	//  Hive-partitioned loads with require_partition_filter explicitly set to
	//  true will fail.
	// +kcc:proto:field=google.cloud.bigquery.v2.HivePartitioningOptions.require_partition_filter
	RequirePartitionFilter *bool `json:"requirePartitionFilter,omitempty"`
}

// +kcc:proto=google.cloud.bigquery.v2.JoinRestrictionPolicy
type JoinRestrictionPolicy struct {
	// Optional. Specifies if a join is required or not on queries for the view.
	//  Default is JOIN_CONDITION_UNSPECIFIED.
	// +kcc:proto:field=google.cloud.bigquery.v2.JoinRestrictionPolicy.join_condition
	JoinCondition *string `json:"joinCondition,omitempty"`

	// Optional. The only columns that joins are allowed on.
	//  This field is must be specified for join_conditions JOIN_ANY and JOIN_ALL
	//  and it cannot be set for JOIN_BLOCKED.
	// +kcc:proto:field=google.cloud.bigquery.v2.JoinRestrictionPolicy.join_allowed_columns
	JoinAllowedColumns []string `json:"joinAllowedColumns,omitempty"`
}

// +kcc:proto=google.cloud.bigquery.v2.JsonOptions
type JsonOptions struct {
	// Optional. The character encoding of the data.
	//  The supported values are UTF-8, UTF-16BE, UTF-16LE, UTF-32BE,
	//  and UTF-32LE.  The default value is UTF-8.
	// +kcc:proto:field=google.cloud.bigquery.v2.JsonOptions.encoding
	Encoding *string `json:"encoding,omitempty"`
}

// +kcc:proto=google.cloud.bigquery.v2.MaterializedViewDefinition
type MaterializedViewDefinition struct {
	// Required. A query whose results are persisted.
	// +kcc:proto:field=google.cloud.bigquery.v2.MaterializedViewDefinition.query
	Query *string `json:"query,omitempty"`

	// Optional. Enable automatic refresh of the materialized view when the base
	//  table is updated. The default value is "true".
	// +kcc:proto:field=google.cloud.bigquery.v2.MaterializedViewDefinition.enable_refresh
	EnableRefresh *bool `json:"enableRefresh,omitempty"`

	// Optional. The maximum frequency at which this materialized view will be
	//  refreshed. The default value is "1800000" (30 minutes).
	// +kcc:proto:field=google.cloud.bigquery.v2.MaterializedViewDefinition.refresh_interval_ms
	RefreshIntervalMs *UInt64Value `json:"refreshIntervalMs,omitempty"`

	// Optional. This option declares the intention to construct a materialized
	//  view that isn't refreshed incrementally.
	// +kcc:proto:field=google.cloud.bigquery.v2.MaterializedViewDefinition.allow_non_incremental_definition
	AllowNonIncrementalDefinition *bool `json:"allowNonIncrementalDefinition,omitempty"`
}

// +kcc:proto=google.cloud.bigquery.v2.MaterializedViewStatus
type MaterializedViewStatus struct {
}

// +kcc:proto=google.cloud.bigquery.v2.ParquetOptions
type ParquetOptions struct {
	// Optional. Indicates whether to infer Parquet ENUM logical type as STRING
	//  instead of BYTES by default.
	// +kcc:proto:field=google.cloud.bigquery.v2.ParquetOptions.enum_as_string
	EnumAsString *bool `json:"enumAsString,omitempty"`

	// Optional. Indicates whether to use schema inference specifically for
	//  Parquet LIST logical type.
	// +kcc:proto:field=google.cloud.bigquery.v2.ParquetOptions.enable_list_inference
	EnableListInference *bool `json:"enableListInference,omitempty"`

	// Optional. Indicates how to represent a Parquet map if present.
	// +kcc:proto:field=google.cloud.bigquery.v2.ParquetOptions.map_target_type
	MapTargetType *string `json:"mapTargetType,omitempty"`
}

// +kcc:proto=google.cloud.bigquery.v2.PartitionedColumn
type PartitionedColumn struct {
	// Required. The name of the partition column.
	// +kcc:proto:field=google.cloud.bigquery.v2.PartitionedColumn.field
	Field *string `json:"field,omitempty"`
}

// +kcc:proto=google.cloud.bigquery.v2.PartitioningDefinition
type PartitioningDefinition struct {
	// Optional. Details about each partitioning column. This field is output only
	//  for all partitioning types other than metastore partitioned tables.
	//  BigQuery native tables only support 1 partitioning column. Other table
	//  types may support 0, 1 or more partitioning columns.
	//  For metastore partitioned tables, the order must match the definition order
	//  in the Hive Metastore, where it must match the physical layout of the
	//  table. For example,
	//
	//  CREATE TABLE a_table(id BIGINT, name STRING)
	//  PARTITIONED BY (city STRING, state STRING).
	//
	//  In this case the values must be ['city', 'state'] in that order.
	// +kcc:proto:field=google.cloud.bigquery.v2.PartitioningDefinition.partitioned_column
	PartitionedColumn []PartitionedColumn `json:"partitionedColumn,omitempty"`
}

// +kcc:proto=google.cloud.bigquery.v2.PrimaryKey
type PrimaryKey struct {
	// Required. The columns that are composed of the primary key constraint.
	// +kcc:proto:field=google.cloud.bigquery.v2.PrimaryKey.columns
	Columns []string `json:"columns,omitempty"`
}

// +kcc:proto=google.cloud.bigquery.v2.PrivacyPolicy
type PrivacyPolicy struct {
	// Optional. Policy used for aggregation thresholds.
	// +kcc:proto:field=google.cloud.bigquery.v2.PrivacyPolicy.aggregation_threshold_policy
	AggregationThresholdPolicy *AggregationThresholdPolicy `json:"aggregationThresholdPolicy,omitempty"`

	// Optional. Policy used for differential privacy.
	// +kcc:proto:field=google.cloud.bigquery.v2.PrivacyPolicy.differential_privacy_policy
	DifferentialPrivacyPolicy *DifferentialPrivacyPolicy `json:"differentialPrivacyPolicy,omitempty"`

	// Optional. Join restriction policy is outside of the one of policies, since
	//  this policy can be set along with other policies. This policy gives data
	//  providers the ability to enforce joins on the 'join_allowed_columns' when
	//  data is queried from a privacy protected view.
	// +kcc:proto:field=google.cloud.bigquery.v2.PrivacyPolicy.join_restriction_policy
	JoinRestrictionPolicy *JoinRestrictionPolicy `json:"joinRestrictionPolicy,omitempty"`
}

// +kcc:proto=google.cloud.bigquery.v2.RangePartitioning
type RangePartitioning struct {
	// Required. The name of the column to partition the table on. It must be a
	//  top-level, INT64 column whose mode is NULLABLE or REQUIRED.
	// +kcc:proto:field=google.cloud.bigquery.v2.RangePartitioning.field
	Field *string `json:"field,omitempty"`

	// Defines the ranges for range partitioning.
	// +kcc:proto:field=google.cloud.bigquery.v2.RangePartitioning.range
	Range *RangePartitioning_Range `json:"range,omitempty"`
}

// +kcc:proto=google.cloud.bigquery.v2.RangePartitioning.Range
type RangePartitioning_Range struct {
	// Required. The start of range partitioning, inclusive. This field is an
	//  INT64 value represented as a string.
	// +kcc:proto:field=google.cloud.bigquery.v2.RangePartitioning.Range.start
	Start *string `json:"start,omitempty"`

	// Required. The end of range partitioning, exclusive. This field is an
	//  INT64 value represented as a string.
	// +kcc:proto:field=google.cloud.bigquery.v2.RangePartitioning.Range.end
	End *string `json:"end,omitempty"`

	// Required. The width of each interval. This field is an INT64 value
	//  represented as a string.
	// +kcc:proto:field=google.cloud.bigquery.v2.RangePartitioning.Range.interval
	Interval *string `json:"interval,omitempty"`
}

// +kcc:proto=google.cloud.bigquery.v2.RestrictionConfig
type RestrictionConfig struct {
}

// +kcc:proto=google.cloud.bigquery.v2.SerDeInfo
type SerDeInfo struct {
	// Optional. Name of the SerDe.
	//  The maximum length is 256 characters.
	// +kcc:proto:field=google.cloud.bigquery.v2.SerDeInfo.name
	Name *string `json:"name,omitempty"`

	// Required. Specifies a fully-qualified class name of the serialization
	//  library that is responsible for the translation of data between table
	//  representation and the underlying low-level input and output format
	//  structures. The maximum length is 256 characters.
	// +kcc:proto:field=google.cloud.bigquery.v2.SerDeInfo.serialization_library
	SerializationLibrary *string `json:"serializationLibrary,omitempty"`

	// Optional. Key-value pairs that define the initialization parameters for the
	//  serialization library.
	//  Maximum size 10 Kib.
	// +kcc:proto:field=google.cloud.bigquery.v2.SerDeInfo.parameters
	Parameters map[string]string `json:"parameters,omitempty"`
}

// +kcc:proto=google.cloud.bigquery.v2.SnapshotDefinition
type SnapshotDefinition struct {
	// Required. Reference describing the ID of the table that was snapshot.
	// +kcc:proto:field=google.cloud.bigquery.v2.SnapshotDefinition.base_table_reference
	BaseTableReference *TableReference `json:"baseTableReference,omitempty"`

	// Required. The time at which the base table was snapshot. This value is
	//  reported in the JSON response using RFC3339 format.
	// +kcc:proto:field=google.cloud.bigquery.v2.SnapshotDefinition.snapshot_time
	SnapshotTime *string `json:"snapshotTime,omitempty"`
}

// +kcc:proto=google.cloud.bigquery.v2.StorageDescriptor
type StorageDescriptor struct {
	// Optional. The physical location of the table
	//  (e.g. `gs://spark-dataproc-data/pangea-data/case_sensitive/` or
	//  `gs://spark-dataproc-data/pangea-data/*`).
	//  The maximum length is 2056 bytes.
	// +kcc:proto:field=google.cloud.bigquery.v2.StorageDescriptor.location_uri
	LocationURI *string `json:"locationURI,omitempty"`

	// Optional. Specifies the fully qualified class name of the InputFormat
	//  (e.g. "org.apache.hadoop.hive.ql.io.orc.OrcInputFormat").
	//  The maximum length is 128 characters.
	// +kcc:proto:field=google.cloud.bigquery.v2.StorageDescriptor.input_format
	InputFormat *string `json:"inputFormat,omitempty"`

	// Optional. Specifies the fully qualified class name of the OutputFormat
	//  (e.g. "org.apache.hadoop.hive.ql.io.orc.OrcOutputFormat").
	//  The maximum length is 128 characters.
	// +kcc:proto:field=google.cloud.bigquery.v2.StorageDescriptor.output_format
	OutputFormat *string `json:"outputFormat,omitempty"`

	// Optional. Serializer and deserializer information.
	// +kcc:proto:field=google.cloud.bigquery.v2.StorageDescriptor.serde_info
	SerdeInfo *SerDeInfo `json:"serdeInfo,omitempty"`
}

// +kcc:proto=google.cloud.bigquery.v2.Streamingbuffer
type Streamingbuffer struct {
}

// +kcc:proto=google.cloud.bigquery.v2.TableConstraints
type TableConstraints struct {
	// Optional. Represents a primary key constraint on a table's columns.
	//  Present only if the table has a primary key.
	//  The primary key is not enforced.
	// +kcc:proto:field=google.cloud.bigquery.v2.TableConstraints.primary_key
	PrimaryKey *PrimaryKey `json:"primaryKey,omitempty"`

	// Optional. Present only if the table has a foreign key.
	//  The foreign key is not enforced.
	// +kcc:proto:field=google.cloud.bigquery.v2.TableConstraints.foreign_keys
	ForeignKeys []ForeignKey `json:"foreignKeys,omitempty"`
}

// +kcc:proto=google.cloud.bigquery.v2.TableFieldSchema
type TableFieldSchema struct {
	// Required. The field name. The name must contain only letters (a-z, A-Z),
	//  numbers (0-9), or underscores (_), and must start with a letter or
	//  underscore. The maximum length is 300 characters.
	// +kcc:proto:field=google.cloud.bigquery.v2.TableFieldSchema.name
	Name *string `json:"name,omitempty"`

	// Required. The field data type. Possible values include:
	//
	//  * STRING
	//  * BYTES
	//  * INTEGER (or INT64)
	//  * FLOAT (or FLOAT64)
	//  * BOOLEAN (or BOOL)
	//  * TIMESTAMP
	//  * DATE
	//  * TIME
	//  * DATETIME
	//  * GEOGRAPHY
	//  * NUMERIC
	//  * BIGNUMERIC
	//  * JSON
	//  * RECORD (or STRUCT)
	//  * RANGE
	//
	//  Use of RECORD/STRUCT indicates that the field contains a nested schema.
	// +kcc:proto:field=google.cloud.bigquery.v2.TableFieldSchema.type
	Type *string `json:"type,omitempty"`

	// Optional. The field mode. Possible values include NULLABLE, REQUIRED and
	//  REPEATED. The default value is NULLABLE.
	// +kcc:proto:field=google.cloud.bigquery.v2.TableFieldSchema.mode
	Mode *string `json:"mode,omitempty"`

	// Optional. Describes the nested schema fields if the type property is set
	//  to RECORD.
	// +kcc:proto:field=google.cloud.bigquery.v2.TableFieldSchema.fields
	Fields []TableFieldSchema `json:"fields,omitempty"`

	// Optional. The field description. The maximum length is 1,024 characters.
	// +kcc:proto:field=google.cloud.bigquery.v2.TableFieldSchema.description
	Description *string `json:"description,omitempty"`

	// Optional. The policy tags attached to this field, used for field-level
	//  access control. If not set, defaults to empty policy_tags.
	// +kcc:proto:field=google.cloud.bigquery.v2.TableFieldSchema.policy_tags
	PolicyTags *TableFieldSchema_PolicyTagList `json:"policyTags,omitempty"`

	// Optional. Data policy options, will replace the data_policies.
	// +kcc:proto:field=google.cloud.bigquery.v2.TableFieldSchema.data_policies
	DataPolicies []DataPolicyOption `json:"dataPolicies,omitempty"`

	// Optional. Maximum length of values of this field for STRINGS or BYTES.
	//
	//  If max_length is not specified, no maximum length constraint is imposed
	//  on this field.
	//
	//  If type = "STRING", then max_length represents the maximum UTF-8
	//  length of strings in this field.
	//
	//  If type = "BYTES", then max_length represents the maximum number of
	//  bytes in this field.
	//
	//  It is invalid to set this field if type &ne; "STRING" and &ne; "BYTES".
	// +kcc:proto:field=google.cloud.bigquery.v2.TableFieldSchema.max_length
	MaxLength *int64 `json:"maxLength,omitempty"`

	// Optional. Precision (maximum number of total digits in base 10) and scale
	//  (maximum number of digits in the fractional part in base 10) constraints
	//  for values of this field for NUMERIC or BIGNUMERIC.
	//
	//  It is invalid to set precision or scale if type &ne; "NUMERIC" and &ne;
	//  "BIGNUMERIC".
	//
	//  If precision and scale are not specified, no value range constraint is
	//  imposed on this field insofar as values are permitted by the type.
	//
	//  Values of this NUMERIC or BIGNUMERIC field must be in this range when:
	//
	//  * Precision (<var>P</var>) and scale (<var>S</var>) are specified:
	//    [-10<sup><var>P</var>-<var>S</var></sup> + 10<sup>-<var>S</var></sup>,
	//     10<sup><var>P</var>-<var>S</var></sup> - 10<sup>-<var>S</var></sup>]
	//  * Precision (<var>P</var>) is specified but not scale (and thus scale is
	//    interpreted to be equal to zero):
	//    [-10<sup><var>P</var></sup> + 1, 10<sup><var>P</var></sup> - 1].
	//
	//  Acceptable values for precision and scale if both are specified:
	//
	//  * If type = "NUMERIC":
	//    1 &le; precision - scale &le; 29 and 0 &le; scale &le; 9.
	//  * If type = "BIGNUMERIC":
	//    1 &le; precision - scale &le; 38 and 0 &le; scale &le; 38.
	//
	//  Acceptable values for precision if only precision is specified but not
	//  scale (and thus scale is interpreted to be equal to zero):
	//
	//  * If type = "NUMERIC": 1 &le; precision &le; 29.
	//  * If type = "BIGNUMERIC": 1 &le; precision &le; 38.
	//
	//  If scale is specified but not precision, then it is invalid.
	// +kcc:proto:field=google.cloud.bigquery.v2.TableFieldSchema.precision
	Precision *int64 `json:"precision,omitempty"`

	// Optional. See documentation for precision.
	// +kcc:proto:field=google.cloud.bigquery.v2.TableFieldSchema.scale
	Scale *int64 `json:"scale,omitempty"`

	// Optional. Specifies the rounding mode to be used when storing values of
	//  NUMERIC and BIGNUMERIC type.
	// +kcc:proto:field=google.cloud.bigquery.v2.TableFieldSchema.rounding_mode
	RoundingMode *string `json:"roundingMode,omitempty"`

	// Optional. Field collation can be set only when the type of field is STRING.
	//  The following values are supported:
	//
	//  * 'und:ci': undetermined locale, case insensitive.
	//  * '': empty string. Default to case-sensitive behavior.
	// +kcc:proto:field=google.cloud.bigquery.v2.TableFieldSchema.collation
	Collation *string `json:"collation,omitempty"`

	// Optional. A SQL expression to specify the [default value]
	//  (https://cloud.google.com/bigquery/docs/default-values) for this field.
	// +kcc:proto:field=google.cloud.bigquery.v2.TableFieldSchema.default_value_expression
	DefaultValueExpression *string `json:"defaultValueExpression,omitempty"`

	// Optional. The subtype of the RANGE, if the type of this field is RANGE. If
	//  the type is RANGE, this field is required. Values for the field element
	//  type can be the following:
	//
	//  * DATE
	//  * DATETIME
	//  * TIMESTAMP
	// +kcc:proto:field=google.cloud.bigquery.v2.TableFieldSchema.range_element_type
	RangeElementType *TableFieldSchema_FieldElementType `json:"rangeElementType,omitempty"`

	// Optional. Definition of the foreign data type.
	//  Only valid for top-level schema fields (not nested fields).
	//  If the type is FOREIGN, this field is required.
	// +kcc:proto:field=google.cloud.bigquery.v2.TableFieldSchema.foreign_type_definition
	ForeignTypeDefinition *string `json:"foreignTypeDefinition,omitempty"`
}

// +kcc:proto=google.cloud.bigquery.v2.TableFieldSchema.FieldElementType
type TableFieldSchema_FieldElementType struct {
	// Required. The type of a field element. For more information, see
	//  [TableFieldSchema.type][google.cloud.bigquery.v2.TableFieldSchema.type].
	// +kcc:proto:field=google.cloud.bigquery.v2.TableFieldSchema.FieldElementType.type
	Type *string `json:"type,omitempty"`
}

// +kcc:proto=google.cloud.bigquery.v2.TableFieldSchema.PolicyTagList
type TableFieldSchema_PolicyTagList struct {
	// A list of policy tag resource names. For example,
	//  "projects/1/locations/eu/taxonomies/2/policyTags/3". At most 1 policy tag
	//  is currently allowed.
	// +kcc:proto:field=google.cloud.bigquery.v2.TableFieldSchema.PolicyTagList.names
	Names []string `json:"names,omitempty"`
}

// +kcc:proto=google.cloud.bigquery.v2.TableReference
type TableReference struct {
	// Required. The ID of the project containing this table.
	// +kcc:proto:field=google.cloud.bigquery.v2.TableReference.project_id
	ProjectID *string `json:"projectID,omitempty"`

	// Required. The ID of the dataset containing this table.
	// +kcc:proto:field=google.cloud.bigquery.v2.TableReference.dataset_id
	DatasetID *string `json:"datasetID,omitempty"`

	// Required. The ID of the table. The ID can contain Unicode characters in
	//  category L (letter), M (mark), N (number), Pc (connector, including
	//  underscore), Pd (dash), and Zs (space). For more information, see [General
	//  Category](https://wikipedia.org/wiki/Unicode_character_property#General_Category).
	//  The maximum length is 1,024 characters.  Certain operations allow suffixing
	//  of the table ID with a partition decorator, such as
	//  `sample_table$20190123`.
	// +kcc:proto:field=google.cloud.bigquery.v2.TableReference.table_id
	TableID *string `json:"tableID,omitempty"`
}

// +kcc:proto=google.cloud.bigquery.v2.TableReplicationInfo
type TableReplicationInfo struct {
	// Required. Source table reference that is replicated.
	// +kcc:proto:field=google.cloud.bigquery.v2.TableReplicationInfo.source_table
	SourceTable *TableReference `json:"sourceTable,omitempty"`

	// Optional. Specifies the interval at which the source table is polled for
	//  updates.
	//  It's Optional. If not specified, default replication interval would be
	//  applied.
	// +kcc:proto:field=google.cloud.bigquery.v2.TableReplicationInfo.replication_interval_ms
	ReplicationIntervalMs *int64 `json:"replicationIntervalMs,omitempty"`
}

// +kcc:proto=google.cloud.bigquery.v2.TableSchema
type TableSchema struct {
	// Describes the fields in a table.
	// +kcc:proto:field=google.cloud.bigquery.v2.TableSchema.fields
	Fields []TableFieldSchema `json:"fields,omitempty"`

	// Optional. Specifies metadata of the foreign data type definition in field
	//  schema
	//  ([TableFieldSchema.foreign_type_definition][google.cloud.bigquery.v2.TableFieldSchema.foreign_type_definition]).
	// +kcc:proto:field=google.cloud.bigquery.v2.TableSchema.foreign_type_info
	ForeignTypeInfo *ForeignTypeInfo `json:"foreignTypeInfo,omitempty"`
}

// +kcc:proto=google.cloud.bigquery.v2.TimePartitioning
type TimePartitioning struct {
	// Required. The supported types are DAY, HOUR, MONTH, and YEAR, which will
	//  generate one partition per day, hour, month, and year, respectively.
	// +kcc:proto:field=google.cloud.bigquery.v2.TimePartitioning.type
	Type *string `json:"type,omitempty"`

	// Optional. Number of milliseconds for which to keep the storage for a
	//  partition.
	//  A wrapper is used here because 0 is an invalid value.
	// +kcc:proto:field=google.cloud.bigquery.v2.TimePartitioning.expiration_ms
	ExpirationMs *int64 `json:"expirationMs,omitempty"`

	// Optional. If not set, the table is partitioned by pseudo
	//  column '_PARTITIONTIME'; if set, the table is partitioned by this field.
	//  The field must be a top-level TIMESTAMP or DATE field. Its mode must be
	//  NULLABLE or REQUIRED.
	//  A wrapper is used here because an empty string is an invalid value.
	// +kcc:proto:field=google.cloud.bigquery.v2.TimePartitioning.field
	Field *string `json:"field,omitempty"`
}

// +kcc:proto=google.cloud.bigquery.v2.UserDefinedFunctionResource
type UserDefinedFunctionResource struct {
	// [Pick one] A code resource to load from a Google Cloud Storage URI
	//  (gs://bucket/path).
	// +kcc:proto:field=google.cloud.bigquery.v2.UserDefinedFunctionResource.resource_uri
	ResourceURI *string `json:"resourceURI,omitempty"`

	// [Pick one] An inline resource that contains code for a user-defined
	//  function (UDF). Providing a inline code resource is equivalent to providing
	//  a URI for a file containing the same code.
	// +kcc:proto:field=google.cloud.bigquery.v2.UserDefinedFunctionResource.inline_code
	InlineCode *string `json:"inlineCode,omitempty"`
}

// +kcc:proto=google.cloud.bigquery.v2.ViewDefinition
type ViewDefinition struct {
	// Required. A query that BigQuery executes when the view is referenced.
	// +kcc:proto:field=google.cloud.bigquery.v2.ViewDefinition.query
	Query *string `json:"query,omitempty"`

	// Describes user-defined function resources used in the query.
	// +kcc:proto:field=google.cloud.bigquery.v2.ViewDefinition.user_defined_function_resources
	UserDefinedFunctionResources []UserDefinedFunctionResource `json:"userDefinedFunctionResources,omitempty"`

	// Specifies whether to use BigQuery's legacy SQL for this view.
	//  The default value is true. If set to false, the view will use
	//  BigQuery's GoogleSQL:
	//  https://cloud.google.com/bigquery/sql-reference/
	//
	//  Queries and views that reference this view must use the same flag value.
	//  A wrapper is used here because the default value is True.
	// +kcc:proto:field=google.cloud.bigquery.v2.ViewDefinition.use_legacy_sql
	UseLegacySQL *bool `json:"useLegacySQL,omitempty"`

	// True if the column names are explicitly specified. For example by using the
	//  'CREATE VIEW v(c1, c2) AS ...' syntax.
	//  Can only be set for GoogleSQL views.
	// +kcc:proto:field=google.cloud.bigquery.v2.ViewDefinition.use_explicit_column_names
	UseExplicitColumnNames *bool `json:"useExplicitColumnNames,omitempty"`

	// Optional. Specifies the privacy policy for the view.
	// +kcc:proto:field=google.cloud.bigquery.v2.ViewDefinition.privacy_policy
	PrivacyPolicy *PrivacyPolicy `json:"privacyPolicy,omitempty"`

	// Optional. Foreign view representations.
	// +kcc:proto:field=google.cloud.bigquery.v2.ViewDefinition.foreign_definitions
	ForeignDefinitions []ForeignViewDefinition `json:"foreignDefinitions,omitempty"`
}

// +kcc:proto=google.protobuf.BytesValue
type BytesValue struct {
	// The bytes value.
	// +kcc:proto:field=google.protobuf.BytesValue.value
	Value []byte `json:"value,omitempty"`
}

// +kcc:proto=google.protobuf.Int32Value
type Int32Value struct {
	// The int32 value.
	// +kcc:proto:field=google.protobuf.Int32Value.value
	Value *int32 `json:"value,omitempty"`
}

// +kcc:proto=google.protobuf.UInt64Value
type UInt64Value struct {
	// The uint64 value.
	// +kcc:proto:field=google.protobuf.UInt64Value.value
	Value *uint64 `json:"value,omitempty"`
}

// +kcc:proto=google.cloud.bigquery.v2.DifferentialPrivacyPolicy
type DifferentialPrivacyPolicyObservedState struct {
	// Output only. The epsilon budget remaining. If budget is exhausted, no more
	//  queries are allowed. Note that the budget for queries that are in progress
	//  is deducted before the query executes. If the query fails or is cancelled
	//  then the budget is refunded. In this case the amount of budget remaining
	//  can increase.
	// +kcc:proto:field=google.cloud.bigquery.v2.DifferentialPrivacyPolicy.epsilon_budget_remaining
	EpsilonBudgetRemaining *float64 `json:"epsilonBudgetRemaining,omitempty"`

	// Output only. The delta budget remaining. If budget is exhausted, no more
	//  queries are allowed. Note that the budget for queries that are in progress
	//  is deducted before the query executes. If the query fails or is cancelled
	//  then the budget is refunded. In this case the amount of budget remaining
	//  can increase.
	// +kcc:proto:field=google.cloud.bigquery.v2.DifferentialPrivacyPolicy.delta_budget_remaining
	DeltaBudgetRemaining *float64 `json:"deltaBudgetRemaining,omitempty"`
}

// +kcc:proto=google.cloud.bigquery.v2.ExternalDataConfiguration
type ExternalDataConfigurationObservedState struct {
	// Optional. When set, configures hive partitioning support. Not all storage
	//  formats support hive partitioning -- requesting hive partitioning on an
	//  unsupported format will lead to an error, as will providing an invalid
	//  specification.
	// +kcc:proto:field=google.cloud.bigquery.v2.ExternalDataConfiguration.hive_partitioning_options
	HivePartitioningOptions *HivePartitioningOptionsObservedState `json:"hivePartitioningOptions,omitempty"`
}

// +kcc:proto=google.cloud.bigquery.v2.HivePartitioningOptions
type HivePartitioningOptionsObservedState struct {
	// Output only. For permanent external tables, this field is populated with
	//  the hive partition keys in the order they were inferred. The types of the
	//  partition keys can be deduced by checking the table schema (which will
	//  include the partition keys). Not every API will populate this field in the
	//  output. For example, Tables.Get will populate it, but Tables.List will not
	//  contain this field.
	// +kcc:proto:field=google.cloud.bigquery.v2.HivePartitioningOptions.fields
	Fields []string `json:"fields,omitempty"`
}

// +kcc:proto=google.cloud.bigquery.v2.MaterializedViewDefinition
type MaterializedViewDefinitionObservedState struct {
	// Output only. The time when this materialized view was last refreshed, in
	//  milliseconds since the epoch.
	// +kcc:proto:field=google.cloud.bigquery.v2.MaterializedViewDefinition.last_refresh_time
	LastRefreshTime *int64 `json:"lastRefreshTime,omitempty"`
}

// +kcc:proto=google.cloud.bigquery.v2.MaterializedViewStatus
type MaterializedViewStatusObservedState struct {
	// Output only. Refresh watermark of materialized view. The base tables' data
	//  were collected into the materialized view cache until this time.
	// +kcc:proto:field=google.cloud.bigquery.v2.MaterializedViewStatus.refresh_watermark
	RefreshWatermark *string `json:"refreshWatermark,omitempty"`

	// Output only. Error result of the last automatic refresh. If present,
	//  indicates that the last automatic refresh was unsuccessful.
	// +kcc:proto:field=google.cloud.bigquery.v2.MaterializedViewStatus.last_refresh_status
	LastRefreshStatus *ErrorProto `json:"lastRefreshStatus,omitempty"`
}

// +kcc:proto=google.cloud.bigquery.v2.PrivacyPolicy
type PrivacyPolicyObservedState struct {
	// Optional. Policy used for differential privacy.
	// +kcc:proto:field=google.cloud.bigquery.v2.PrivacyPolicy.differential_privacy_policy
	DifferentialPrivacyPolicy *DifferentialPrivacyPolicyObservedState `json:"differentialPrivacyPolicy,omitempty"`
}

// +kcc:proto=google.cloud.bigquery.v2.RestrictionConfig
type RestrictionConfigObservedState struct {
	// Output only. Specifies the type of dataset/table restriction.
	// +kcc:proto:field=google.cloud.bigquery.v2.RestrictionConfig.type
	Type *string `json:"type,omitempty"`
}

// +kcc:proto=google.cloud.bigquery.v2.Streamingbuffer
type StreamingbufferObservedState struct {
	// Output only. A lower-bound estimate of the number of bytes currently in
	//  the streaming buffer.
	// +kcc:proto:field=google.cloud.bigquery.v2.Streamingbuffer.estimated_bytes
	EstimatedBytes *uint64 `json:"estimatedBytes,omitempty"`

	// Output only. A lower-bound estimate of the number of rows currently in the
	//  streaming buffer.
	// +kcc:proto:field=google.cloud.bigquery.v2.Streamingbuffer.estimated_rows
	EstimatedRows *uint64 `json:"estimatedRows,omitempty"`

	// Output only. Contains the timestamp of the oldest entry in the streaming
	//  buffer, in milliseconds since the epoch, if the streaming buffer is
	//  available.
	// +kcc:proto:field=google.cloud.bigquery.v2.Streamingbuffer.oldest_entry_time
	OldestEntryTime *uint64 `json:"oldestEntryTime,omitempty"`
}

// +kcc:proto=google.cloud.bigquery.v2.TableReplicationInfo
type TableReplicationInfoObservedState struct {
	// Optional. Output only. If source is a materialized view, this field
	//  signifies the last refresh time of the source.
	// +kcc:proto:field=google.cloud.bigquery.v2.TableReplicationInfo.replicated_source_last_refresh_time
	ReplicatedSourceLastRefreshTime *int64 `json:"replicatedSourceLastRefreshTime,omitempty"`

	// Optional. Output only. Replication status of configured replication.
	// +kcc:proto:field=google.cloud.bigquery.v2.TableReplicationInfo.replication_status
	ReplicationStatus *string `json:"replicationStatus,omitempty"`

	// Optional. Output only. Replication error that will permanently stopped
	//  table replication.
	// +kcc:proto:field=google.cloud.bigquery.v2.TableReplicationInfo.replication_error
	ReplicationError *ErrorProto `json:"replicationError,omitempty"`
}

// +kcc:proto=google.cloud.bigquery.v2.ViewDefinition
type ViewDefinitionObservedState struct {
	// Optional. Specifies the privacy policy for the view.
	// +kcc:proto:field=google.cloud.bigquery.v2.ViewDefinition.privacy_policy
	PrivacyPolicy *PrivacyPolicyObservedState `json:"privacyPolicy,omitempty"`
}
