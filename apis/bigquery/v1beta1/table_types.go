// Copyright 2025 Google LLC
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//    http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package v1beta1

import (
	refs "github.com/GoogleCloudPlatform/k8s-config-connector/apis/refs/v1beta1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/apis/k8s/v1alpha1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

var BigQueryTableGVK = GroupVersion.WithKind("BigQueryTable")

// +kcc:proto=google.cloud.bigquery.v2.AvroOptions
type AvroOptions struct {
	// If sourceFormat is set to "AVRO", indicates whether to interpret
	//  logical types as the corresponding BigQuery data type (for example,
	//  TIMESTAMP), instead of using the raw type (for example, INTEGER).
	// +required
	// +kcc:proto:field=google.cloud.bigquery.v2.AvroOptions.use_avro_logical_types
	UseAvroLogicalTypes *bool `json:"useAvroLogicalTypes,omitempty"`
}

// +kcc:proto=google.cloud.bigquery.v2.MaterializedViewDefinition
type MaterializedViewDefinition struct {
	// Required. A query whose results are persisted.
	// +required
	// +kcc:proto:field=google.cloud.bigquery.v2.MaterializedViewDefinition.query
	Query *string `json:"query,omitempty"`

	// Enable automatic refresh of the materialized view when the base
	//  table is updated. The default value is "true".
	// +optional
	// +kcc:proto:field=google.cloud.bigquery.v2.MaterializedViewDefinition.enable_refresh
	EnableRefresh *bool `json:"enableRefresh,omitempty"`

	// The maximum frequency at which this materialized view will be
	//  refreshed. The default value is "1800000" (30 minutes).
	// +optional
	// +kcc:proto:field=google.cloud.bigquery.v2.MaterializedViewDefinition.refresh_interval_ms
	RefreshIntervalMs *int64 `json:"refreshIntervalMs,omitempty"`

	// This option declares the intention to construct a materialized
	//  view that isn't refreshed incrementally.
	// +optional
	// +kcc:proto:field=google.cloud.bigquery.v2.MaterializedViewDefinition.allow_non_incremental_definition
	AllowNonIncrementalDefinition *bool `json:"allowNonIncrementalDefinition,omitempty"`
}

// +kcc:proto=google.cloud.bigquery.v2.CsvOptions
type CsvOptions struct {
	// The separator character for fields in a CSV file. The separator
	//  is interpreted as a single byte. For files encoded in ISO-8859-1, any
	//  single character can be used as a separator. For files encoded in UTF-8,
	//  characters represented in decimal range 1-127 (U+0001-U+007F) can be used
	//  without any modification. UTF-8 characters encoded with multiple bytes
	//  (i.e. U+0080 and above) will have only the first byte used for separating
	//  fields. The remaining bytes will be treated as a part of the field.
	//  BigQuery also supports the escape sequence "\t" (U+0009) to specify a tab
	//  separator. The default value is comma (",", U+002C).
	// +kcc:proto:field=google.cloud.bigquery.v2.CsvOptions.field_delimiter
	// +optional
	FieldDelimiter *string `json:"fieldDelimiter,omitempty"`

	// The number of rows at the top of a CSV file that BigQuery will
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
	// +optional
	// +kcc:proto:field=google.cloud.bigquery.v2.CsvOptions.skip_leading_rows
	SkipLeadingRows *int64 `json:"skipLeadingRows,omitempty"`

	// The value that is used to quote data sections in a CSV file.
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
	// +required
	// +kcc:proto:field=google.cloud.bigquery.v2.CsvOptions.quote
	Quote *string `json:"quote,omitempty"`

	// Indicates if BigQuery should allow quoted data sections that
	//  contain newline characters in a CSV file. The default value is false.
	// +optional
	// +kcc:proto:field=google.cloud.bigquery.v2.CsvOptions.allow_quoted_newlines
	AllowQuotedNewlines *bool `json:"allowQuotedNewlines,omitempty"`

	// Indicates if BigQuery should accept rows that are missing
	//  trailing optional columns. If true, BigQuery treats missing trailing
	//  columns as null values.
	//  If false, records with missing trailing columns are treated as bad records,
	//  and if there are too many bad records, an invalid error is returned in the
	//  job result. The default value is false.
	// +optional
	// +kcc:proto:field=google.cloud.bigquery.v2.CsvOptions.allow_jagged_rows
	AllowJaggedRows *bool `json:"allowJaggedRows,omitempty"`

	// The character encoding of the data.
	//  The supported values are UTF-8, ISO-8859-1, UTF-16BE, UTF-16LE, UTF-32BE,
	//  and UTF-32LE.  The default value is UTF-8.
	//  BigQuery decodes the data after the raw, binary data has been split using
	//  the values of the quote and fieldDelimiter properties.
	// +optional
	// +kcc:proto:field=google.cloud.bigquery.v2.CsvOptions.encoding
	Encoding *string `json:"encoding,omitempty"`

	// NOT YET
	// Optional. Indicates if the embedded ASCII control characters (the first 32
	//  characters in the ASCII-table, from '\x00' to '\x1F') are preserved.
	// +kcc:proto:field=google.cloud.bigquery.v2.CsvOptions.preserve_ascii_control_characters
	// PreserveAsciiControlCharacters *bool `json:"preserveAsciiControlCharacters,omitempty"`

	// NOT YET
	// Optional. Specifies a string that represents a null value in a CSV file.
	//  For example, if you specify "\N", BigQuery interprets "\N" as a null value
	//  when querying a CSV file.
	//  The default value is the empty string. If you set this property to a custom
	//  value, BigQuery throws an error if an empty string is present for all data
	//  types except for STRING and BYTE. For STRING and BYTE columns, BigQuery
	//  interprets the empty string as an empty value.
	// +kcc:proto:field=google.cloud.bigquery.v2.CsvOptions.null_marker
	// NullMarker *string `json:"nullMarker,omitempty"`
}

// +kcc:proto=google.cloud.bigquery.v2.GoogleSheetsOptions
type GoogleSheetsOptions struct {
	// The number of rows at the top of a sheet that BigQuery will skip
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
	// +optional
	// +kcc:proto:field=google.cloud.bigquery.v2.GoogleSheetsOptions.skip_leading_rows
	SkipLeadingRows *int64 `json:"skipLeadingRows,omitempty"`

	// Range of a sheet to query from. Only used when non-empty.
	//  Typical format: sheet_name!top_left_cell_id:bottom_right_cell_id
	//  For example: sheet1!A1:B20
	// +optional
	// +kcc:proto:field=google.cloud.bigquery.v2.GoogleSheetsOptions.range
	Range *string `json:"range,omitempty"`
}

// +kcc:proto=google.cloud.bigquery.v2.JsonOptions
type JsonOptions struct {
	// The character encoding of the data.
	//  The supported values are UTF-8, UTF-16BE, UTF-16LE, UTF-32BE,
	//  and UTF-32LE.  The default value is UTF-8.
	// +optional
	// +kcc:proto:field=google.cloud.bigquery.v2.JsonOptions.encoding
	Encoding *string `json:"encoding,omitempty"`
}

// +kcc:proto=google.cloud.bigquery.v2.HivePartitioningOptions
type HivePartitioningOptions struct {
	// When set, what mode of hive partitioning to use when reading
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
	// +optional
	// +kcc:proto:field=google.cloud.bigquery.v2.HivePartitioningOptions.mode
	Mode *string `json:"mode,omitempty"`

	// When hive partition detection is requested, a common prefix for
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
	// +optional
	// +kcc:proto:field=google.cloud.bigquery.v2.HivePartitioningOptions.source_uri_prefix
	SourceUriPrefix *string `json:"sourceUriPrefix,omitempty"`

	// If set to true, queries over this table require a partition
	//  filter that can be used for partition elimination to be specified.
	//
	//  Note that this field should only be true when creating a permanent
	//  external table or querying a temporary external table.
	//
	//  Hive-partitioned loads with require_partition_filter explicitly set to
	//  true will fail.
	// +optional
	// +kcc:proto:field=google.cloud.bigquery.v2.HivePartitioningOptions.require_partition_filter
	RequirePartitionFilter *bool `json:"requirePartitionFilter,omitempty"`
}

// +kcc:proto=google.cloud.bigquery.v2.ParquetOptions
type ParquetOptions struct {
	// Indicates whether to infer Parquet ENUM logical type as STRING
	//  instead of BYTES by default.
	// +optional
	// +kcc:proto:field=google.cloud.bigquery.v2.ParquetOptions.enum_as_string
	EnumAsString *bool `json:"enumAsString,omitempty"`

	// Indicates whether to use schema inference specifically for
	//  Parquet LIST logical type.
	// +optional
	// +kcc:proto:field=google.cloud.bigquery.v2.ParquetOptions.enable_list_inference
	EnableListInference *bool `json:"enableListInference,omitempty"`

	// NOT YET
	// Optional. Indicates how to represent a Parquet map if present.
	// +kcc:proto:field=google.cloud.bigquery.v2.ParquetOptions.map_target_type
	// MapTargetType *string `json:"mapTargetType,omitempty"`
}

// +kcc:proto=google.cloud.bigquery.v2.RangePartitioning.Range
type RangePartitioning_Range struct {
	// Immutable. The start of range partitioning, inclusive.
	// +required
	// +kcc:proto:field=google.cloud.bigquery.v2.RangePartitioning.Range.start
	Start *int64 `json:"start,omitempty"`

	// Required. The end of range partitioning, exclusive.
	// +required
	// +kcc:proto:field=google.cloud.bigquery.v2.RangePartitioning.Range.end
	End *int64 `json:"end,omitempty"`

	// Required. The width of each interval.
	// +required
	// +kcc:proto:field=google.cloud.bigquery.v2.RangePartitioning.Range.interval
	Interval *int64 `json:"interval,omitempty"`
}

// +kcc:proto=google.cloud.bigquery.v2.RangePartitioning
type RangePartitioning struct {
	// Required. The name of the column to partition the table on. It must be a
	//  top-level, INT64 column whose mode is NULLABLE or REQUIRED.
	// +required
	// +kcc:proto:field=google.cloud.bigquery.v2.RangePartitioning.field
	Field *string `json:"field,omitempty"`

	// Defines the ranges for range partitioning.
	// +required
	// +kcc:proto:field=google.cloud.bigquery.v2.RangePartitioning.range
	Range RangePartitioning_Range `json:"range,omitempty"`
}

// +kcc:proto=google.cloud.bigquery.v2.ExternalDataConfiguration
type ExternalDataConfiguration struct {
	// Try to detect schema and format options automatically.
	//  Any option specified explicitly will be honored.
	// +required
	// +kcc:proto:field=google.cloud.bigquery.v2.ExternalDataConfiguration.autodetect
	Autodetect *bool `json:"autodetect,omitempty"`

	// Additional properties to set if sourceFormat is set to AVRO.
	// +kcc:proto:field=google.cloud.bigquery.v2.ExternalDataConfiguration.avro_options
	// +optional
	AvroOptions *AvroOptions `json:"avroOptions,omitempty"`

	// The compression type of the data source.
	//  Possible values include GZIP and NONE. The default value is NONE.
	//  This setting is ignored for Google Cloud Bigtable, Google Cloud Datastore
	//  backups, Avro, ORC and Parquet
	//  formats. An empty string is an invalid value.
	// +kcc:proto:field=google.cloud.bigquery.v2.ExternalDataConfiguration.compression
	// +optional
	Compression *string `json:"compression,omitempty"`

	// The connection specifying the credentials to be used to read
	//  external storage, such as Azure Blob, Cloud Storage, or S3. The
	//  connection_id can have the form
	//  `{project_id}.{location_id};{connection_id}` or
	//  `projects/{project_id}/locations/{location_id}/connections/{connection_id}`.
	// +optional
	// +kcc:proto:field=google.cloud.bigquery.v2.ExternalDataConfiguration.connection_id
	ConnectionId *string `json:"connectionId,omitempty"`

	// Additional properties to set if sourceFormat is set to CSV.
	// +kcc:proto:field=google.cloud.bigquery.v2.ExternalDataConfiguration.csv_options
	// +optional
	CsvOptions *CsvOptions `json:"csvOptions,omitempty"`

	// Specifies how source URIs are interpreted for constructing the
	//  file set to load.  By default source URIs are expanded against the
	//  underlying storage.  Other options include specifying manifest files. Only
	//  applicable to object storage systems.
	// +optional
	// +kcc:proto:field=google.cloud.bigquery.v2.ExternalDataConfiguration.file_set_spec_type
	FileSetSpecType *string `json:"fileSetSpecType,omitempty"`

	// Additional options if sourceFormat is set to GOOGLE_SHEETS.
	// +optional
	// +kcc:proto:field=google.cloud.bigquery.v2.ExternalDataConfiguration.google_sheets_options
	GoogleSheetsOptions *GoogleSheetsOptions `json:"googleSheetsOptions,omitempty"`

	// When set, configures hive partitioning support. Not all storage
	//  formats support hive partitioning -- requesting hive partitioning on an
	//  unsupported format will lead to an error, as will providing an invalid
	//  specification.
	// +optional
	// +kcc:proto:field=google.cloud.bigquery.v2.ExternalDataConfiguration.hive_partitioning_options
	HivePartitioningOptions *HivePartitioningOptions `json:"hivePartitioningOptions,omitempty"`

	// Indicates if BigQuery should allow extra values that are not
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
	// +optional
	// +kcc:proto:field=google.cloud.bigquery.v2.ExternalDataConfiguration.ignore_unknown_values
	IgnoreUnknownValues *bool `json:"ignoreUnknownValues,omitempty"`

	// Additional properties to set if sourceFormat is set to JSON.
	// +optional
	// +kcc:proto:field=google.cloud.bigquery.v2.ExternalDataConfiguration.json_options
	JsonOptions *JsonOptions `json:"jsonOptions,omitempty"`

	// The maximum number of bad records that BigQuery can ignore when
	//  reading data. If the number of bad records exceeds this value, an invalid
	//  error is returned in the job result. The default value is 0, which requires
	//  that all records are valid. This setting is ignored for Google Cloud
	//  Bigtable, Google Cloud Datastore backups, Avro, ORC and Parquet formats.
	// +optional
	// +kcc:proto:field=google.cloud.bigquery.v2.ExternalDataConfiguration.max_bad_records
	MaxBadRecords *int64 `json:"maxBadRecords,omitempty"`

	// Metadata Cache Mode for the table. Set this to enable caching of
	//  metadata from external data source.
	// +optional
	// +kcc:proto:field=google.cloud.bigquery.v2.ExternalDataConfiguration.metadata_cache_mode
	MetadataCacheMode *string `json:"metadataCacheMode,omitempty"`

	// ObjectMetadata is used to create Object Tables. Object Tables
	//  contain a listing of objects (with their metadata) found at the
	//  source_uris. If ObjectMetadata is set, source_format should be omitted.
	//
	//  Currently SIMPLE is the only supported Object Metadata type.
	// +optional
	// +kcc:proto:field=google.cloud.bigquery.v2.ExternalDataConfiguration.object_metadata
	ObjectMetadata *string `json:"objectMetadata,omitempty"`

	// Additional properties to set if sourceFormat is set to PARQUET.
	// +optional
	// +kcc:proto:field=google.cloud.bigquery.v2.ExternalDataConfiguration.parquet_options
	ParquetOptions *ParquetOptions `json:"parquetOptions,omitempty"`

	// When creating an external table, the user can provide a reference
	//  file with the table schema. This is enabled for the following formats:
	//  AVRO, PARQUET, ORC.
	// +optional
	// +kcc:proto:field=google.cloud.bigquery.v2.ExternalDataConfiguration.reference_file_schema_uri
	ReferenceFileSchemaUri *string `json:"referenceFileSchemaUri,omitempty"`

	// The schema for the data.
	//  Schema is required for CSV and JSON formats if autodetect is not on.
	//  Schema is disallowed for Google Cloud Bigtable, Cloud Datastore backups,
	//  Avro, ORC and Parquet formats.
	// +optional
	// +kcc:proto:field=google.cloud.bigquery.v2.ExternalDataConfiguration.schema
	Schema *string `json:"schema,omitempty"`

	/* Please see sourceFormat under ExternalDataConfiguration in Bigquery's public API documentation (https://cloud.google.com/bigquery/docs/reference/rest/v2/tables#externaldataconfiguration) for supported formats. To use "GOOGLE_SHEETS" the scopes must include "googleapis.com/auth/drive.readonly". */
	// +optional
	// +kcc:proto:field=google.cloud.bigquery.v2.ExternalDataConfiguration.source_format
	SourceFormat *string `json:"sourceFormat,omitempty"`

	// A list of the fully-qualified URIs that point to	your data in Google Cloud.
	// +required
	// +kcc:proto:field=google.cloud.bigquery.v2.ExternalDataConfiguration.source_uris
	SourceUris []string `json:"sourceUris,omitempty"`

	// NOT YET
	// Optional. Additional options if sourceFormat is set to BIGTABLE.
	// +kcc:proto:field=google.cloud.bigquery.v2.ExternalDataConfiguration.bigtable_options
	// BigtableOptions *BigtableOptions `json:"bigtableOptions,omitempty"`

	// NOT YET
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
	// DecimalTargetTypes []string `json:"decimalTargetTypes,omitempty"`

	// NOT YET
	// Optional. Load option to be used together with source_format
	//  newline-delimited JSON to indicate that a variant of JSON is being loaded.
	//  To load newline-delimited GeoJSON, specify GEOJSON (and source_format must
	//  be set to NEWLINE_DELIMITED_JSON).
	// +kcc:proto:field=google.cloud.bigquery.v2.ExternalDataConfiguration.json_extension
	// JsonExtension *string `json:"jsonExtension,omitempty"`
}

// +kcc:proto=google.cloud.bigquery.v2.PrimaryKey
type PrimaryKey struct {
	// Required. The columns that are composed of the primary key constraint.
	// +required
	// +kcc:proto:field=google.cloud.bigquery.v2.PrimaryKey.columns
	Columns []string `json:"columns,omitempty"`
}

// +kcc:proto=google.cloud.bigquery.v2.ForeignKey
type ForeignKey struct {
	// Optional. Set only if the foreign key constraint is named.
	// +kcc:proto:field=google.cloud.bigquery.v2.ForeignKey.name
	// +optional
	Name *string `json:"name,omitempty"`

	// Required. The table that holds the primary key and is referenced by this
	//  foreign key.
	// +required
	// +kcc:proto:field=google.cloud.bigquery.v2.ForeignKey.referenced_table
	ReferencedTable TableReference `json:"referencedTable"`

	// Required. The columns that compose the foreign key.
	// +required
	// +kcc:proto:field=google.cloud.bigquery.v2.ForeignKey.column_references
	ColumnReferences ColumnReference `json:"columnReferences"`
}

// +kcc:proto=google.cloud.bigquery.v2.TableConstraints
type TableConstraints struct {
	// Represents a primary key constraint on a table's columns.
	//  Present only if the table has a primary key.
	//  The primary key is not enforced.
	// +optional
	// +kcc:proto:field=google.cloud.bigquery.v2.TableConstraints.primary_key
	PrimaryKey *PrimaryKey `json:"primaryKey,omitempty"`

	// Present only if the table has a foreign key.
	//  The foreign key is not enforced.
	// +optional
	// +kcc:proto:field=google.cloud.bigquery.v2.TableConstraints.foreign_keys
	ForeignKeys []ForeignKey `json:"foreignKeys,omitempty"`
}

// +kcc:proto=google.cloud.bigquery.v2.ColumnReference
type ColumnReference struct {
	// Required. The column that composes the foreign key.
	// +required
	// +kcc:proto:field=google.cloud.bigquery.v2.ColumnReference.referencing_column
	ReferencingColumn *string `json:"referencingColumn,omitempty"`

	// Required. The column in the primary key that are referenced by the
	//  referencing_column.
	// +required
	// +kcc:proto:field=google.cloud.bigquery.v2.ColumnReference.referenced_column
	ReferencedColumn *string `json:"referencedColumn,omitempty"`
}

// +kcc:proto=google.cloud.bigquery.v2.TimePartitioning
type TimePartitioning struct {
	// Required. The supported types are DAY, HOUR, MONTH, and YEAR, which will
	//  generate one partition per day, hour, month, and year, respectively.
	// +required
	// +kcc:proto:field=google.cloud.bigquery.v2.TimePartitioning.type
	Type string `json:"type,omitempty"`

	// Number of milliseconds for which to keep the storage for a
	//  partition.
	//  A wrapper is used here because 0 is an invalid value.
	// +optional
	// +kcc:proto:field=google.cloud.bigquery.v2.TimePartitioning.expiration_ms
	ExpirationMs *int64 `json:"expirationMs,omitempty"`

	/* DEPRECATED. This field is deprecated; please use the top level field with the same name instead. If set to true, queries over this table require a partition filter that can be used for partition elimination to be specified. */
	// +optional
	RequirePartitionFilter *bool `json:"requirePartitionFilter,omitempty"`

	// Optional. Immutable. If not set, the table is partitioned by pseudo
	//  column '_PARTITIONTIME'; if set, the table is partitioned by this field.
	//  The field must be a top-level TIMESTAMP or DATE field. Its mode must be
	//  NULLABLE or REQUIRED.
	//  A wrapper is used here because an empty string is an invalid value.
	// +optional
	// +kcc:proto:field=google.cloud.bigquery.v2.TimePartitioning.field
	Field *string `json:"field,omitempty"`
}

// +kcc:proto=google.cloud.bigquery.v2.ViewDefinition
type ViewDefinition struct {
	// Required. A query that BigQuery executes when the view is referenced.
	// +required
	// +kcc:proto:field=google.cloud.bigquery.v2.ViewDefinition.query
	Query *string `json:"query"`

	// Specifies whether to use BigQuery's legacy SQL for this view.
	//  The default value is true. If set to false, the view will use
	//  BigQuery's GoogleSQL:
	//  https://cloud.google.com/bigquery/sql-reference/
	//
	//  Queries and views that reference this view must use the same flag value.
	//  A wrapper is used here because the default value is True.
	// +kcc:proto:field=google.cloud.bigquery.v2.ViewDefinition.use_legacy_sql
	// +optional
	UseLegacySql *bool `json:"useLegacySql,omitempty"`

	// NOT YET
	// Describes user-defined function resources used in the query.
	// +kcc:proto:field=google.cloud.bigquery.v2.ViewDefinition.user_defined_function_resources
	// UserDefinedFunctionResources []UserDefinedFunctionResource `json:"userDefinedFunctionResources,omitempty"`

	// NOT YET
	// True if the column names are explicitly specified. For example by using the
	//  'CREATE VIEW v(c1, c2) AS ...' syntax.
	//  Can only be set for GoogleSQL views.
	// +kcc:proto:field=google.cloud.bigquery.v2.ViewDefinition.use_explicit_column_names
	// UseExplicitColumnNames *bool `json:"useExplicitColumnNames,omitempty"`

	// NOT YET
	// Optional. Specify the privacy policy for the view.
	// +kcc:proto:field=google.cloud.bigquery.v2.ViewDefinition.privacy_policy
	// PrivacyPolicy *PrivacyPolicy `json:"privacyPolicy,omitempty"`

	// NOT YET
	// Optional. Foreign view representations.
	// +kcc:proto:field=google.cloud.bigquery.v2.ViewDefinition.foreign_definitions
	// ForeignDefinitions []ForeignViewDefinition `json:"foreignDefinitions,omitempty"`
}

type TableEncryptionConfiguration struct {
	// Describes the Cloud KMS encryption key that will be used to
	//  protect destination BigQuery table. The BigQuery Service Account associated
	//  with your project requires access to this encryption key.
	// +required
	KmsKeyRef *refs.KMSCryptoKeyRef `json:"kmsKeyRef,omitempty"`

	/* DEPRECATED.
	// The table will be encrypted with the primary version of Cloud KMS encryption key.
	// The self link or full name of the kms key version used to encrypt this table. */
	// +optional
	KmsKeyVersion *string `json:"kmsKeyVersion,omitempty"`
}

// BigQueryTableSpec defines the desired state of BigQueryTable
// +kcc:proto=google.cloud.bigquery.v2.Table
type BigQueryTableSpec struct {
	// Clustering specification for the table. Must be specified with time-based
	//  partitioning, data in the table will be first partitioned and subsequently
	//  clustered.
	// +kcc:proto:field=google.cloud.bigquery.v2.Table.clustering
	Clustering []string `json:"clustering,omitempty"`

	// The BigQueryTable name. If not given, the metadata.name will be used.
	ResourceID *string `json:"resourceID,omitempty"`

	// +required
	DatasetRef *DatasetRef `json:"datasetRef"`

	// A user-friendly description of this table.
	// +optional
	// +kcc:proto:field=google.cloud.bigquery.v2.Table.description
	Description *string `json:"description,omitempty"`

	// Custom encryption configuration (e.g., Cloud KMS keys).
	// +kcc:proto:field=google.cloud.bigquery.v2.Table.encryption_configuration
	EncryptionConfiguration *TableEncryptionConfiguration `json:"encryptionConfiguration,omitempty"`

	// The time when this table expires, in milliseconds since the
	//  epoch. If not present, the table will persist indefinitely. Expired tables
	//  will be deleted and their storage reclaimed.  The defaultTableExpirationMs
	//  property of the encapsulating dataset can be used to set a default
	//  expirationTime on newly created tables.
	// +optional
	// +kcc:proto:field=google.cloud.bigquery.v2.Table.expiration_time
	ExpirationTime *int64 `json:"expirationTime,omitempty"`

	// Describes the data format, location, and other properties of
	//  a table stored outside of BigQuery. By defining these properties, the data
	//  source can then be queried as if it were a standard BigQuery table.
	// +optional
	// +kcc:proto:field=google.cloud.bigquery.v2.Table.external_data_configuration
	ExternalDataConfiguration *ExternalDataConfiguration `json:"externalDataConfiguration,omitempty"`

	// A descriptive name for this table.
	// +kcc:proto:field=google.cloud.bigquery.v2.Table.friendly_name
	// +optional
	FriendlyName *string `json:"friendlyName,omitempty"`

	// The materialized view definition.
	// +optional
	// +kcc:proto:field=google.cloud.bigquery.v2.Table.materialized_view
	MaterializedView *MaterializedViewDefinition `json:"materializedView,omitempty"`

	// The maximum staleness of data that could be returned when the
	//  table (or stale MV) is queried. Staleness encoded as a string encoding
	//  of sql IntervalValue type.
	// +optional
	// +kcc:proto:field=google.cloud.bigquery.v2.Table.max_staleness
	MaxStaleness *string `json:"maxStaleness,omitempty"`

	// If specified, configures range partitioning for this table.
	// +kcc:proto:field=google.cloud.bigquery.v2.Table.range_partitioning
	RangePartitioning *RangePartitioning `json:"rangePartitioning,omitempty"`

	// If set to true, queries over this table require
	//  a partition filter that can be used for partition elimination to be
	//  specified.
	// +optional
	// +kcc:proto:field=google.cloud.bigquery.v2.Table.require_partition_filter
	RequirePartitionFilter *bool `json:"requirePartitionFilter,omitempty"`

	// Describes the schema of this table.
	// +optional
	// +kcc:proto:field=google.cloud.bigquery.v2.Table.schema
	Schema *string `json:"schema,omitempty"`

	// Tables Primary Key and Foreign Key information
	// +optional
	// +kcc:proto:field=google.cloud.bigquery.v2.Table.table_constraints
	TableConstraints *TableConstraints `json:"tableConstraints,omitempty"`

	// If specified, configures time-based partitioning for this table.
	// +kcc:proto:field=google.cloud.bigquery.v2.Table.time_partitioning
	TimePartitioning *TimePartitioning `json:"timePartitioning,omitempty"`

	// The view definition.
	// +optional
	// +kcc:proto:field=google.cloud.bigquery.v2.Table.view
	View *ViewDefinition `json:"view,omitempty"`

	// When using `alpha.cnrm.cloud.google.com/reconciler:direct` annotion, use labels field
	// to set the labels for this resource on GCP.
	// Otherwise, use .metadata.labels.
	// Please refer to https://github.com/GoogleCloudPlatform/k8s-config-connector/issues/4274 for context.
	// The labels associated with this table. You can use these to organize and
	//  group your tables. Label keys and values can be no longer than 63
	//  characters, can only contain lowercase letters, numeric characters,
	//  underscores and dashes. International characters are allowed. Label values
	//  are optional. Label keys must start with a letter and each label in the
	//  list must have a different key.
	// +kcc:proto:field=google.cloud.bigquery.v2.Table.labels
	Labels map[string]string `json:"labels,omitempty"`

	// Required. Reference describing the ID of this table.
	// +kcc:proto:field=google.cloud.bigquery.v2.Table.table_reference
	// TableReference *TableReference `json:"tableReference,omitempty"`

	// NOT YET
	// Optional. The partition information for all table formats, including
	//  managed partitioned tables, hive partitioned tables, iceberg partitioned,
	//  and metastore partitioned tables. This field is only populated for
	//  metastore partitioned tables. For other table formats, this is an output
	//  only field.
	// +kcc:proto:field=google.cloud.bigquery.v2.Table.partition_definition
	// PartitionDefinition *PartitioningDefinition `json:"partitionDefinition,omitempty"`

	// NOT YET
	// Optional. Specifies the configuration of a BigLake managed table.
	// +kcc:proto:field=google.cloud.bigquery.v2.Table.biglake_configuration
	// BiglakeConfiguration *BigLakeConfiguration `json:"biglakeConfiguration,omitempty"`

	// NOT YET
	// Optional. Defines the default collation specification of new STRING fields
	//  in the table. During table creation or update, if a STRING field is added
	//  to this table without explicit collation specified, then the table inherits
	//  the table default collation. A change to this field affects only fields
	//  added afterwards, and does not alter the existing fields.
	//  The following values are supported:
	//
	//  * 'und:ci': undetermined locale, case insensitive.
	//  * '': empty string. Default to case-sensitive behavior.
	// +kcc:proto:field=google.cloud.bigquery.v2.Table.default_collation
	// DefaultCollation *string `json:"defaultCollation,omitempty"`

	// NOT YET
	// Optional. Defines the default rounding mode specification of new decimal
	//  fields (NUMERIC OR BIGNUMERIC) in the table. During table creation or
	//  update, if a decimal field is added to this table without an explicit
	//  rounding mode specified, then the field inherits the table default
	//  rounding mode. Changing this field doesn't affect existing fields.
	// +kcc:proto:field=google.cloud.bigquery.v2.Table.default_rounding_mode
	// DefaultRoundingMode *string `json:"defaultRoundingMode,omitempty"`

	// NOT YET
	// Optional. The [tags](https://cloud.google.com/bigquery/docs/tags) attached
	//  to this table. Tag keys are globally unique. Tag key is expected to be in
	//  the namespaced format, for example "123456789012/environment" where
	//  123456789012 is the ID of the parent organization or project resource for
	//  this tag key. Tag value is expected to be the short name, for example
	//  "Production". See [Tag
	//  definitions](https://cloud.google.com/iam/docs/tags-access-control#definitions)
	//  for more details.
	// +kcc:proto:field=google.cloud.bigquery.v2.Table.resource_tags
	// ResourceTags map[string]string `json:"resourceTags,omitempty"`

	// NOT YET
	// Optional. Table replication info for table created `AS REPLICA` DDL like:
	//  `CREATE MATERIALIZED VIEW mv1 AS REPLICA OF src_mv`
	// +kcc:proto:field=google.cloud.bigquery.v2.Table.table_replication_info
	// TableReplicationInfo *TableReplicationInfo `json:"tableReplicationInfo,omitempty"`

	// NOT YET
	// Optional. Options defining open source compatible table.
	// +kcc:proto:field=google.cloud.bigquery.v2.Table.external_catalog_table_options
	// ExternalCatalogTableOptions *ExternalCatalogTableOptions `json:"externalCatalogTableOptions,omitempty"`
}

// BigQueryTableStatus defines the config connector machine state of BigQueryTable
type BigQueryTableStatus struct {
	/* Conditions represent the latest available observations of the
	   object's current state. */
	Conditions []v1alpha1.Condition `json:"conditions,omitempty"`

	// ObservedGeneration is the generation of the resource that was most recently observed by the Config Connector controller. If this is equal to metadata.generation, then that means that the current reported status reflects the most recent desired state of the resource.
	ObservedGeneration *int64 `json:"observedGeneration,omitempty"`

	// A unique specifier for the BigQueryTable resource in GCP.
	ExternalRef *string `json:"externalRef,omitempty"`

	// ObservedState is the state of the resource as most recently observed in GCP.
	ObservedState *BigQueryTableObservedState `json:"observedState,omitempty"`

	// Output only. The time when this table was created, in milliseconds since
	//  the epoch.
	// +kcc:proto:field=google.cloud.bigquery.v2.Table.creation_time
	CreationTime *int64 `json:"creationTime,omitempty"`

	// Output only. A hash of this resource.
	// +kcc:proto:field=google.cloud.bigquery.v2.Table.etag
	Etag *string `json:"etag,omitempty"`

	// Output only. The time when this table was last modified, in milliseconds
	//  since the epoch.
	// +kcc:proto:field=google.cloud.bigquery.v2.Table.last_modified_time
	LastModifiedTime *int64 `json:"lastModifiedTime,omitempty"`

	// Output only. The geographic location where the table resides. This value
	//  is inherited from the dataset.
	// +kcc:proto:field=google.cloud.bigquery.v2.Table.location
	Location *string `json:"location,omitempty"`
	// Output only. The size of this table in logical bytes, excluding any data in
	//  the streaming buffer.
	// +kcc:proto:field=google.cloud.bigquery.v2.Table.num_bytes
	NumBytes *int64 `json:"numBytes,omitempty"`

	// Output only. The number of logical bytes in the table that are considered
	//  "long-term storage".
	// +kcc:proto:field=google.cloud.bigquery.v2.Table.num_long_term_bytes
	NumLongTermBytes *int64 `json:"numLongTermBytes,omitempty"`

	// Output only. The number of rows of data in this table, excluding any data
	//  in the streaming buffer.
	// +kcc:proto:field=google.cloud.bigquery.v2.Table.num_rows
	NumRows *int64 `json:"numRows,omitempty"`

	// Output only. A URL that can be used to access this resource again.
	// +kcc:proto:field=google.cloud.bigquery.v2.Table.self_link
	SelfLink *string `json:"selfLink,omitempty"`

	// Output only. Describes the table type. The following values are supported:
	//
	//  * `TABLE`: A normal BigQuery table.
	//  * `VIEW`: A virtual table defined by a SQL query.
	//  * `EXTERNAL`: A table that references data stored in an external storage
	//    system, such as Google Cloud Storage.
	//  * `MATERIALIZED_VIEW`: A precomputed view defined by a SQL query.
	//  * `SNAPSHOT`: An immutable BigQuery table that preserves the contents of a
	//    base table at a particular time. See additional information on
	//    [table
	//    snapshots](https://cloud.google.com/bigquery/docs/table-snapshots-intro).
	//
	//  The default value is `TABLE`.
	// +kcc:proto:field=google.cloud.bigquery.v2.Table.type
	Type *string `json:"type,omitempty"`
}

// BigQueryTableObservedState is the state of the BigQueryTable resource as most recently observed in GCP.
// +kcc:proto=google.cloud.bigquery.v2.Table
type BigQueryTableObservedState struct {
	// NOT YET
	// Output only. The physical size of this table in bytes. This includes
	//  storage used for time travel.
	// +kcc:proto:field=google.cloud.bigquery.v2.Table.num_physical_bytes
	// NumPhysicalBytes *int64 `json:"numPhysicalBytes,omitempty"`

	// NOT YET
	// Optional. The view definition.
	// +kcc:proto:field=google.cloud.bigquery.v2.Table.view
	// View *ViewDefinitionObservedState `json:"view,omitempty"`

	// NOT YET
	// Optional. The materialized view definition.
	// +kcc:proto:field=google.cloud.bigquery.v2.Table.materialized_view
	// MaterializedView *MaterializedViewDefinitionObservedState `json:"materializedView,omitempty"`

	// NOT YET
	// Output only. The materialized view status.
	// +kcc:proto:field=google.cloud.bigquery.v2.Table.materialized_view_status
	// MaterializedViewStatus *MaterializedViewStatus `json:"materializedViewStatus,omitempty"`

	// NOT YET
	// Optional. Describes the data format, location, and other properties of
	//  a table stored outside of BigQuery. By defining these properties, the data
	//  source can then be queried as if it were a standard BigQuery table.
	// +kcc:proto:field=google.cloud.bigquery.v2.Table.external_data_configuration
	// ExternalDataConfiguration *ExternalDataConfigurationObservedState `json:"externalDataConfiguration,omitempty"`

	// NOT YET
	// Output only. Contains information regarding this table's streaming buffer,
	//  if one is present. This field will be absent if the table is not being
	//  streamed to or if there is no data in the streaming buffer.
	// +kcc:proto:field=google.cloud.bigquery.v2.Table.streaming_buffer
	// StreamingBuffer *Streamingbuffer `json:"streamingBuffer,omitempty"`

	// NOT YET
	// Output only. Contains information about the snapshot. This value is set via
	//  snapshot creation.
	// +kcc:proto:field=google.cloud.bigquery.v2.Table.snapshot_definition
	// SnapshotDefinition *SnapshotDefinition `json:"snapshotDefinition,omitempty"`

	// NOT YET
	// Output only. Contains information about the clone. This value is set via
	//  the clone operation.
	// +kcc:proto:field=google.cloud.bigquery.v2.Table.clone_definition
	// CloneDefinition *CloneDefinition `json:"cloneDefinition,omitempty"`

	// NOT YET
	// Output only. Number of physical bytes used by time travel storage (deleted
	//  or changed data). This data is not kept in real time, and might be delayed
	//  by a few seconds to a few minutes.
	// +kcc:proto:field=google.cloud.bigquery.v2.Table.num_time_travel_physical_bytes
	// NumTimeTravelPhysicalBytes *int64 `json:"numTimeTravelPhysicalBytes,omitempty"`

	// NOT YET
	// Output only. Total number of logical bytes in the table or materialized
	//  view.
	// +kcc:proto:field=google.cloud.bigquery.v2.Table.num_total_logical_bytes
	// NumTotalLogicalBytes *int64 `json:"numTotalLogicalBytes,omitempty"`

	// NOT YET
	// Output only. Number of logical bytes that are less than 90 days old.
	// +kcc:proto:field=google.cloud.bigquery.v2.Table.num_active_logical_bytes
	// NumActiveLogicalBytes *int64 `json:"numActiveLogicalBytes,omitempty"`

	// NOT YET
	// Output only. Number of logical bytes that are more than 90 days old.
	// +kcc:proto:field=google.cloud.bigquery.v2.Table.num_long_term_logical_bytes
	// NumLongTermLogicalBytes *int64 `json:"numLongTermLogicalBytes,omitempty"`

	// NOT YET
	// Output only. Number of physical bytes used by current live data storage.
	//  This data is not kept in real time, and might be delayed by a few seconds
	//  to a few minutes.
	// +kcc:proto:field=google.cloud.bigquery.v2.Table.num_current_physical_bytes
	// NumCurrentPhysicalBytes *int64 `json:"numCurrentPhysicalBytes,omitempty"`

	// NOT YET
	// Output only. The physical size of this table in bytes. This also includes
	//  storage used for time travel. This data is not kept in real time, and might
	//  be delayed by a few seconds to a few minutes.
	// +kcc:proto:field=google.cloud.bigquery.v2.Table.num_total_physical_bytes
	// NumTotalPhysicalBytes *int64 `json:"numTotalPhysicalBytes,omitempty"`

	// NOT YET
	// Output only. Number of physical bytes less than 90 days old. This data is
	//  not kept in real time, and might be delayed by a few seconds to a few
	//  minutes.
	// +kcc:proto:field=google.cloud.bigquery.v2.Table.num_active_physical_bytes
	// NumActivePhysicalBytes *int64 `json:"numActivePhysicalBytes,omitempty"`

	// NOT YET
	// Output only. Number of physical bytes more than 90 days old.
	//  This data is not kept in real time, and might be delayed by a few seconds
	//  to a few minutes.
	// +kcc:proto:field=google.cloud.bigquery.v2.Table.num_long_term_physical_bytes
	// NumLongTermPhysicalBytes *int64 `json:"numLongTermPhysicalBytes,omitempty"`

	// NOT YET
	// Output only. The number of partitions present in the table or materialized
	//  view. This data is not kept in real time, and might be delayed by a few
	//  seconds to a few minutes.
	// +kcc:proto:field=google.cloud.bigquery.v2.Table.num_partitions
	// NumPartitions *int64 `json:"numPartitions,omitempty"`

	// NOT YET
	// Optional. Output only. Restriction config for table. If set, restrict
	//  certain accesses on the table based on the config. See [Data
	//  egress](https://cloud.google.com/bigquery/docs/analytics-hub-introduction#data_egress)
	//  for more details.
	// +kcc:proto:field=google.cloud.bigquery.v2.Table.restrictions
	// Restrictions *RestrictionConfig `json:"restrictions,omitempty"`

	// NOT YET
	// Optional. Table replication info for table created `AS REPLICA` DDL like:
	//  `CREATE MATERIALIZED VIEW mv1 AS REPLICA OF src_mv`
	// +kcc:proto:field=google.cloud.bigquery.v2.Table.table_replication_info
	// TableReplicationInfo *TableReplicationInfoObservedState `json:"tableReplicationInfo,omitempty"`

	// NOT YET
	// Optional. Output only. Table references of all replicas currently active on
	//  the table.
	// +kcc:proto:field=google.cloud.bigquery.v2.Table.replicas
	// Replicas []TableReference `json:"replicas,omitempty"`
}

// +genclient
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:resource:categories=gcp,shortName=gcpbigquerytable;gcpbigquerytables
// +kubebuilder:subresource:status
// +kubebuilder:metadata:labels="cnrm.cloud.google.com/managed-by-kcc=true";"cnrm.cloud.google.com/system=true";"cnrm.cloud.google.com/tf2crd=true"
// +kubebuilder:printcolumn:name="Age",JSONPath=".metadata.creationTimestamp",type="date"
// +kubebuilder:printcolumn:name="Ready",JSONPath=".status.conditions[?(@.type=='Ready')].status",type="string",description="When 'True', the most recent reconcile of the resource succeeded"
// +kubebuilder:printcolumn:name="Status",JSONPath=".status.conditions[?(@.type=='Ready')].reason",type="string",description="The reason for the value in 'Ready'"
// +kubebuilder:printcolumn:name="Status Age",JSONPath=".status.conditions[?(@.type=='Ready')].lastTransitionTime",type="date",description="The last transition time for the value in 'Status'"

// BigQueryTable is the Schema for the BigQueryTable API
// +k8s:openapi-gen=true
type BigQueryTable struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	// +required
	Spec   BigQueryTableSpec   `json:"spec,omitempty"`
	Status BigQueryTableStatus `json:"status,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// BigQueryTableList contains a list of BigQueryTable
type BigQueryTableList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []BigQueryTable `json:"items"`
}

func init() {
	SchemeBuilder.Register(&BigQueryTable{}, &BigQueryTableList{})
}
