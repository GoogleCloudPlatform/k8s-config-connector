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

import (
	refs "github.com/GoogleCloudPlatform/k8s-config-connector/apis/refs/v1beta1"
)

// +kcc:proto=google.cloud.bigquery.v2.Access
type Access struct {
	// An IAM role ID that should be granted to the user, group,
	//  or domain specified in this access entry.
	//  The following legacy mappings will be applied:
	//
	//  * `OWNER`: `roles/bigquery.dataOwner`
	//  * `WRITER`: `roles/bigquery.dataEditor`
	//  * `READER`: `roles/bigquery.dataViewer`
	//
	//  This field will accept any of the above formats, but will return only
	//  the legacy format. For example, if you set this field to
	//  "roles/bigquery.dataOwner", it will be returned back as "OWNER".
	Role *string `json:"role,omitempty"`

	// [Pick one] An email address of a user to grant access to. For example:
	//  fred@example.com. Maps to IAM policy member "user:EMAIL" or
	//  "serviceAccount:EMAIL".
	UserByEmail *string `json:"userByEmail,omitempty"`

	// [Pick one] An email address of a Google Group to grant access to.
	//  Maps to IAM policy member "group:GROUP".
	GroupByEmail *string `json:"groupByEmail,omitempty"`

	// [Pick one] A domain to grant access to. Any users signed in with the domain
	//  specified will be granted the specified access. Example: "example.com".
	//  Maps to IAM policy member "domain:DOMAIN".
	Domain *string `json:"domain,omitempty"`

	// [Pick one] A special group to grant access to. Possible values include:
	//
	//    * projectOwners: Owners of the enclosing project.
	//    * projectReaders: Readers of the enclosing project.
	//    * projectWriters: Writers of the enclosing project.
	//    * allAuthenticatedUsers: All authenticated BigQuery users.
	//
	//  Maps to similarly-named IAM members.
	SpecialGroup *string `json:"specialGroup,omitempty"`

	// [Pick one] Some other type of member that appears in the IAM Policy but
	//  isn't a user, group, domain, or special group.
	IamMember *string `json:"iamMember,omitempty"`

	// [Pick one] A view from a different dataset to grant access to. Queries
	//  executed against that view will have read access to views/tables/routines
	//  in this dataset.
	//  The role field is not required when this field is set. If that view is
	//  updated by any user, access to the view needs to be granted again via an
	//  update operation.
	View *TableReference `json:"view,omitempty"`

	// [Pick one] A routine from a different dataset to grant access to. Queries
	//  executed against that routine will have read access to
	//  views/tables/routines in this dataset. Only UDF is supported for now.
	//  The role field is not required when this field is set. If that routine is
	//  updated by any user, access to the routine needs to be granted again via
	//  an update operation.
	Routine *RoutineReference `json:"routine,omitempty"`

	// [Pick one] A grant authorizing all resources of a particular type in a
	//  particular dataset access to this dataset. Only views are supported for
	//  now. The role field is not required when this field is set. If that dataset
	//  is deleted and re-created, its access needs to be granted again via an
	//  update operation.
	Dataset *DatasetAccessEntry `json:"dataset,omitempty"`
}

// +kcc:proto=google.cloud.bigquery.v2.AggregationThresholdPolicy
type AggregationThresholdPolicy struct {
	// Optional. The threshold for the "aggregation threshold" policy.
	Threshold *int64 `json:"threshold,omitempty"`

	// Optional. The privacy unit column(s) associated with this policy.
	//  For now, only one column per data source object (table, view) is allowed as
	//  a privacy unit column.
	//  Representing as a repeated field in metadata for extensibility to
	//  multiple columns in future.
	//  Duplicates and Repeated struct fields are not allowed.
	//  For nested fields, use dot notation ("outer.inner")
	PrivacyUnitColumns []string `json:"privacyUnitColumns,omitempty"`
}

// +kcc:proto=google.cloud.bigquery.v2.AvroOptions
type AvroOptions struct {
	// Optional. If sourceFormat is set to "AVRO", indicates whether to interpret
	//  logical types as the corresponding BigQuery data type (for example,
	//  TIMESTAMP), instead of using the raw type (for example, INTEGER).
	UseAvroLogicalTypes *bool `json:"useAvroLogicalTypes,omitempty"`
}

// +kcc:proto=google.cloud.bigquery.v2.BiEngineReason
type BiEngineReason struct {
	// Output only. High-level BI Engine reason for partial or disabled
	//  acceleration
	Code *string `json:"code,omitempty"`

	// Output only. Free form human-readable reason for partial or disabled
	//  acceleration.
	Message *string `json:"message,omitempty"`
}

// +kcc:proto=google.cloud.bigquery.v2.BiEngineStatistics
type BiEngineStatistics struct {
	// Output only. Specifies which mode of BI Engine acceleration was performed
	//  (if any).
	BiEngineMode *string `json:"biEngineMode,omitempty"`

	// Output only. Specifies which mode of BI Engine acceleration was performed
	//  (if any).
	AccelerationMode *string `json:"accelerationMode,omitempty"`

	// In case of DISABLED or PARTIAL bi_engine_mode, these contain the
	//  explanatory reasons as to why BI Engine could not accelerate.
	//  In case the full query was accelerated, this field is not populated.
	BiEngineReasons []BiEngineReason `json:"biEngineReasons,omitempty"`
}

// +kcc:proto=google.cloud.bigquery.v2.BigLakeConfiguration
type BigLakeConfiguration struct {
	// Required. The connection specifying the credentials to be used to read and
	//  write to external storage, such as Cloud Storage. The connection_id can
	//  have the form `{project}.{location}.{connection_id}` or
	//  `projects/{project}/locations/{location}/connections/{connection_id}".
	ConnectionID *string `json:"connectionID,omitempty"`

	// Required. The fully qualified location prefix of the external folder where
	//  table data is stored. The '*' wildcard character is not allowed. The URI
	//  should be in the format `gs://bucket/path_to_table/`
	StorageUri *string `json:"storageUri,omitempty"`

	// Required. The file format the table data is stored in.
	FileFormat *string `json:"fileFormat,omitempty"`

	// Required. The table format the metadata only snapshots are stored in.
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
	QualifierEncoded *byte `json:"qualifierEncoded,omitempty"`

	// Qualifier string.
	QualifierString *string `json:"qualifierString,omitempty"`

	// Optional. If the qualifier is not a valid BigQuery field identifier i.e.
	//  does not match [a-zA-Z][a-zA-Z0-9_]*,  a valid identifier must be provided
	//  as the column field name and is used as field name in queries.
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
	Type *string `json:"type,omitempty"`

	// Optional. The encoding of the values when the type is not STRING.
	//  Acceptable encoding values are:
	//    TEXT - indicates values are alphanumeric text strings.
	//    BINARY - indicates values are encoded using HBase Bytes.toBytes family of
	//             functions.
	//  'encoding' can also be set at the column family level. However, the setting
	//  at this level takes precedence if 'encoding' is set at both levels.
	Encoding *string `json:"encoding,omitempty"`

	// Optional. If this is set, only the latest version of value in this column
	//              are exposed.
	//  'onlyReadLatest' can also be set at the column family level. However, the
	//  setting at this level takes precedence if 'onlyReadLatest' is set at both
	//  levels.
	OnlyReadLatest *bool `json:"onlyReadLatest,omitempty"`
}

// +kcc:proto=google.cloud.bigquery.v2.BigtableColumnFamily
type BigtableColumnFamily struct {
	// Identifier of the column family.
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
	Type *string `json:"type,omitempty"`

	// Optional. The encoding of the values when the type is not STRING.
	//  Acceptable encoding values are:
	//    TEXT - indicates values are alphanumeric text strings.
	//    BINARY - indicates values are encoded using HBase Bytes.toBytes family of
	//             functions.
	//  This can be overridden for a specific column by listing that column in
	//  'columns' and specifying an encoding for it.
	Encoding *string `json:"encoding,omitempty"`

	// Optional. Lists of columns that should be exposed as individual fields as
	//  opposed to a list of (column name, value) pairs.
	//  All columns whose qualifier matches a qualifier in this list can be
	//  accessed as `<family field name>.<column field name>`.
	//  Other columns can be accessed as a list through
	//  the `<family field name>.Column` field.
	Columns []BigtableColumn `json:"columns,omitempty"`

	// Optional. If this is set only the latest version of value are exposed for
	//  all columns in this column family.
	//  This can be overridden for a specific column by listing that column in
	//  'columns' and specifying a different setting
	//  for that column.
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
	ColumnFamilies []BigtableColumnFamily `json:"columnFamilies,omitempty"`

	// Optional. If field is true, then the column families that are not
	//  specified in columnFamilies list are not exposed in the table schema.
	//  Otherwise, they are read with BYTES type values.
	//  The default value is false.
	IgnoreUnspecifiedColumnFamilies *bool `json:"ignoreUnspecifiedColumnFamilies,omitempty"`

	// Optional. If field is true, then the rowkey column families will be read
	//  and converted to string. Otherwise they are read with BYTES type values and
	//  users need to manually cast them with CAST if necessary.
	//  The default value is false.
	ReadRowkeyAsString *bool `json:"readRowkeyAsString,omitempty"`

	// Optional. If field is true, then each column family will be read as a
	//  single JSON column. Otherwise they are read as a repeated cell structure
	//  containing timestamp/value tuples. The default value is false.
	OutputColumnFamiliesAsJson *bool `json:"outputColumnFamiliesAsJson,omitempty"`
}

// +kcc:proto=google.cloud.bigquery.v2.CloneDefinition
type CloneDefinition struct {
	// Required. Reference describing the ID of the table that was cloned.
	BaseTableReference *TableReference `json:"baseTableReference,omitempty"`

	// Required. The time at which the base table was cloned. This value is
	//  reported in the JSON response using RFC3339 format.
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
	Fields []string `json:"fields,omitempty"`
}

// +kcc:proto=google.cloud.bigquery.v2.ColumnReference
type ColumnReference struct {
	// Required. The column that composes the foreign key.
	ReferencingColumn *string `json:"referencingColumn,omitempty"`

	// Required. The column in the primary key that are referenced by the
	//  referencing_column.
	ReferencedColumn *string `json:"referencedColumn,omitempty"`
}

// +kcc:proto=google.cloud.bigquery.v2.ConnectionProperty
type ConnectionProperty struct {
	// The key of the property to set.
	Key *string `json:"key,omitempty"`

	// The value of the property to set.
	Value *string `json:"value,omitempty"`
}

// +kcc:proto=google.cloud.bigquery.v2.CopyJobStatistics
type CopyJobStatistics struct {
	// Output only. Number of rows copied to the destination table.
	CopiedRows *int64 `json:"copiedRows,omitempty"`

	// Output only. Number of logical bytes copied to the destination table.
	CopiedLogicalBytes *int64 `json:"copiedLogicalBytes,omitempty"`
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
	SkipLeadingRows *int64 `json:"skipLeadingRows,omitempty"`

	// Optional. The value that is used to quote data sections in a CSV file.
	//  BigQuery converts the string to ISO-8859-1 encoding, and then uses the
	//  first byte of the encoded string to split the data in its raw, binary
	//  state.
	//  The default value is a float64-quote (").
	//  If your data does not contain quoted sections,
	//  set the property value to an empty string.
	//  If your data contains quoted newline characters, you must also set the
	//  allowQuotedNewlines property to true.
	//  To include the specific quote character within a quoted value, precede it
	//  with an additional matching quote character. For example, if you want to
	//  escape the default character  ' " ', use ' "" '.
	Quote *string `json:"quote,omitempty"`

	// Optional. Indicates if BigQuery should allow quoted data sections that
	//  contain newline characters in a CSV file. The default value is false.
	AllowQuotedNewlines *bool `json:"allowQuotedNewlines,omitempty"`

	// Optional. Indicates if BigQuery should accept rows that are missing
	//  trailing optional columns. If true, BigQuery treats missing trailing
	//  columns as null values.
	//  If false, records with missing trailing columns are treated as bad records,
	//  and if there are too many bad records, an invalid error is returned in the
	//  job result. The default value is false.
	AllowJaggedRows *bool `json:"allowJaggedRows,omitempty"`

	// Optional. The character encoding of the data.
	//  The supported values are UTF-8, ISO-8859-1, UTF-16BE, UTF-16LE, UTF-32BE,
	//  and UTF-32LE.  The default value is UTF-8.
	//  BigQuery decodes the data after the raw, binary data has been split using
	//  the values of the quote and fieldDelimiter properties.
	Encoding *string `json:"encoding,omitempty"`

	// Optional. Indicates if the embedded ASCII control characters (the first 32
	//  characters in the ASCII-table, from '\x00' to '\x1F') are preserved.
	PreserveAsciiControlCharacters *bool `json:"preserveAsciiControlCharacters,omitempty"`

	// Optional. Specifies a string that represents a null value in a CSV file.
	//  For example, if you specify "\N", BigQuery interprets "\N" as a null value
	//  when querying a CSV file.
	//  The default value is the empty string. If you set this property to a custom
	//  value, BigQuery throws an error if an empty string is present for all data
	//  types except for STRING and BYTE. For STRING and BYTE columns, BigQuery
	//  interprets the empty string as an empty value.
	NullMarker *string `json:"nullMarker,omitempty"`
}

// +kcc:proto=google.cloud.bigquery.v2.DataFormatOptions
type DataFormatOptions struct {
	// Optional. Output timestamp as usec int64. Default is false.
	UseInt64Timestamp *bool `json:"useInt64Timestamp,omitempty"`
}

// +kcc:proto=google.cloud.bigquery.v2.DataMaskingStatistics
type DataMaskingStatistics struct {
	// Whether any accessed data was protected by the data masking.
	DataMaskingApplied *bool `json:"dataMaskingApplied,omitempty"`
}

// +kcc:proto=google.cloud.bigquery.v2.DataPolicyOption
type DataPolicyOption struct {
	// Data policy resource name in the form of
	//  projects/project_id/locations/location_id/dataPolicies/data_policy_id.
	Name *string `json:"name,omitempty"`
}

// +kcc:proto=google.cloud.bigquery.v2.Dataset
type Dataset struct {
	// Output only. The resource type.
	Kind *string `json:"kind,omitempty"`

	// Output only. A hash of the resource.
	Etag *string `json:"etag,omitempty"`

	// Output only. The fully-qualified unique name of the dataset in the format
	//  projectId:datasetId. The dataset name without the project name is given in
	//  the datasetId field. When creating a new dataset, leave this field blank,
	//  and instead specify the datasetId field.
	ID *string `json:"id,omitempty"`

	// Output only. A URL that can be used to access the resource again. You can
	//  use this URL in Get or Update requests to the resource.
	SelfLink *string `json:"selfLink,omitempty"`

	// Required. A reference that identifies the dataset.
	DatasetReference *DatasetReference `json:"datasetReference,omitempty"`

	// Optional. A descriptive name for the dataset.
	FriendlyName *string `json:"friendlyName,omitempty"`

	// Optional. A user-friendly description of the dataset.
	Description *string `json:"description,omitempty"`

	// Optional. The default lifetime of all tables in the dataset, in
	//  milliseconds. The minimum lifetime value is 3600000 milliseconds (one
	//  hour). To clear an existing default expiration with a PATCH request, set to
	//  0. Once this property is set, all newly-created tables in the dataset will
	//  have an expirationTime property set to the creation time plus the value in
	//  this property, and changing the value will only affect new tables, not
	//  existing ones. When the expirationTime for a given table is reached, that
	//  table will be deleted automatically.
	//  If a table's expirationTime is modified or removed before the table
	//  expires, or if you provide an explicit expirationTime when creating a
	//  table, that value takes precedence over the default expiration time
	//  indicated by this property.
	DefaultTableExpirationMs *int64 `json:"defaultTableExpirationMs,omitempty"`

	// This default partition expiration, expressed in milliseconds.
	//
	//  When new time-partitioned tables are created in a dataset where this
	//  property is set, the table will inherit this value, propagated as the
	//  `TimePartitioning.expirationMs` property on the new table.  If you set
	//  `TimePartitioning.expirationMs` explicitly when creating a table,
	//  the `defaultPartitionExpirationMs` of the containing dataset is ignored.
	//
	//  When creating a partitioned table, if `defaultPartitionExpirationMs`
	//  is set, the `defaultTableExpirationMs` value is ignored and the table
	//  will not be inherit a table expiration deadline.
	DefaultPartitionExpirationMs *int64 `json:"defaultPartitionExpirationMs,omitempty"`

	// The labels associated with this dataset. You can use these
	//  to organize and group your datasets.
	//  You can set this property when inserting or updating a dataset.
	//  See [Creating and Updating Dataset
	//  Labels](https://cloud.google.com/bigquery/docs/creating-managing-labels#creating_and_updating_dataset_labels)
	//  for more information.
	Labels map[string]string `json:"labels,omitempty"`

	// Optional. An array of objects that define dataset access for one or more
	//  entities. You can set this property when inserting or updating a dataset in
	//  order to control who is allowed to access the data. If unspecified at
	//  dataset creation time, BigQuery adds default dataset access for the
	//  following entities: access.specialGroup: projectReaders; access.role:
	//  READER; access.specialGroup: projectWriters; access.role: WRITER;
	//  access.specialGroup: projectOwners; access.role: OWNER;
	//  access.userByEmail: [dataset creator email]; access.role: OWNER;
	//  If you patch a dataset, then this field is overwritten by the patched
	//  dataset's access field. To add entities, you must supply the entire
	//  existing access array in addition to any new entities that you want to add.
	Access []Access `json:"access,omitempty"`

	// Output only. The time when this dataset was created, in milliseconds since
	//  the epoch.
	CreationTime *int64 `json:"creationTime,omitempty"`

	// Output only. The date when this dataset was last modified, in milliseconds
	//  since the epoch.
	LastModifiedTime *int64 `json:"lastModifiedTime,omitempty"`

	// The geographic location where the dataset should reside. See
	//  https://cloud.google.com/bigquery/docs/locations for supported
	//  locations.
	Location *string `json:"location,omitempty"`

	// The default encryption key for all tables in the dataset.
	//  After this property is set, the encryption key of all newly-created tables
	//  in the dataset is set to this value unless the table creation request or
	//  query explicitly overrides the key.
	DefaultEncryptionConfiguration *EncryptionConfiguration `json:"defaultEncryptionConfiguration,omitempty"`

	// Output only. Reserved for future use.
	SatisfiesPzs *bool `json:"satisfiesPzs,omitempty"`

	// Output only. Reserved for future use.
	SatisfiesPzi *bool `json:"satisfiesPzi,omitempty"`

	// Output only. Same as `type` in `ListFormatDataset`.
	//  The type of the dataset, one of:
	//
	//  * DEFAULT - only accessible by owner and authorized accounts,
	//  * PUBLIC - accessible by everyone,
	//  * LINKED - linked dataset,
	//  * EXTERNAL - dataset with definition in external metadata catalog.
	Type *string `json:"type,omitempty"`

	// Optional. The source dataset reference when the dataset is of type LINKED.
	//  For all other dataset types it is not set. This field cannot be updated
	//  once it is set. Any attempt to update this field using Update and Patch API
	//  Operations will be ignored.
	LinkedDatasetSource *LinkedDatasetSource `json:"linkedDatasetSource,omitempty"`

	// Output only. Metadata about the LinkedDataset. Filled out when the dataset
	//  type is LINKED.
	LinkedDatasetMetadata *LinkedDatasetMetadata `json:"linkedDatasetMetadata,omitempty"`

	// Optional. Reference to a read-only external dataset defined in data
	//  catalogs outside of BigQuery. Filled out when the dataset type is EXTERNAL.
	ExternalDatasetReference *ExternalDatasetReference `json:"externalDatasetReference,omitempty"`

	// Optional. Options defining open source compatible datasets living in the
	//  BigQuery catalog. Contains metadata of open source database, schema or
	//  namespace represented by the current dataset.
	ExternalCatalogDatasetOptions *ExternalCatalogDatasetOptions `json:"externalCatalogDatasetOptions,omitempty"`

	// Optional. TRUE if the dataset and its table names are case-insensitive,
	//  otherwise FALSE. By default, this is FALSE, which means the dataset and its
	//  table names are case-sensitive. This field does not affect routine
	//  references.
	IsCaseInsensitive *bool `json:"isCaseInsensitive,omitempty"`

	// Optional. Defines the default collation specification of future tables
	//  created in the dataset. If a table is created in this dataset without
	//  table-level default collation, then the table inherits the dataset default
	//  collation, which is applied to the string fields that do not have explicit
	//  collation specified. A change to this field affects only tables created
	//  afterwards, and does not alter the existing tables.
	//  The following values are supported:
	//
	//  * 'und:ci': undetermined locale, case insensitive.
	//  * '': empty string. Default to case-sensitive behavior.
	DefaultCollation *string `json:"defaultCollation,omitempty"`

	// Optional. Defines the default rounding mode specification of new tables
	//  created within this dataset. During table creation, if this field is
	//  specified, the table within this dataset will inherit the default rounding
	//  mode of the dataset. Setting the default rounding mode on a table overrides
	//  this option. Existing tables in the dataset are unaffected.
	//  If columns are defined during that table creation,
	//  they will immediately inherit the table's default rounding mode,
	//  unless otherwise specified.
	DefaultRoundingMode *string `json:"defaultRoundingMode,omitempty"`

	// Optional. Defines the time travel window in hours. The value can be from 48
	//  to 168 hours (2 to 7 days). The default value is 168 hours if this is not
	//  set.
	MaxTimeTravelHours *int64 `json:"maxTimeTravelHours,omitempty"`

	// Output only. Tags for the dataset. To provide tags as inputs, use the
	//  `resourceTags` field.
	Tags []GcpTag `json:"tags,omitempty"`

	// Optional. Updates storage_billing_model for the dataset.
	StorageBillingModel *string `json:"storageBillingModel,omitempty"`

	// Optional. Output only. Restriction config for all tables and dataset. If
	//  set, restrict certain accesses on the dataset and all its tables based on
	//  the config. See [Data
	//  egress](https://cloud.google.com/bigquery/docs/analytics-hub-introduction#data_egress)
	//  for more details.
	Restrictions *RestrictionConfig `json:"restrictions,omitempty"`

	// Optional. The [tags](https://cloud.google.com/bigquery/docs/tags) attached
	//  to this dataset. Tag keys are globally unique. Tag key is expected to be in
	//  the namespaced format, for example "123456789012/environment" where
	//  123456789012 is the ID of the parent organization or project resource for
	//  this tag key. Tag value is expected to be the short name, for example
	//  "Production". See [Tag
	//  definitions](https://cloud.google.com/iam/docs/tags-access-control#definitions)
	//  for more details.
	ResourceTags map[string]string `json:"resourceTags,omitempty"`
}

// +kcc:proto=google.cloud.bigquery.v2.DatasetAccessEntry
type DatasetAccessEntry struct {
	// The dataset this entry applies to.
	Dataset *DatasetReference `json:"dataset"`

	// Which resources in the dataset this entry applies to. Currently, only
	//  views are supported, but additional target types may be added in the
	//  future.
	TargetTypes []string `json:"targetTypes"`
}

// +kcc:proto=google.cloud.bigquery.v2.DatasetList
type DatasetList struct {
	// Output only. The resource type.
	//  This property always returns the value "bigquery#datasetList"
	Kind *string `json:"kind,omitempty"`

	// Output only. A hash value of the results page. You can use this property to
	//  determine if the page has changed since the last request.
	Etag *string `json:"etag,omitempty"`

	// A token that can be used to request the next results page. This property is
	//  omitted on the final results page.
	NextPageToken *string `json:"nextPageToken,omitempty"`

	// An array of the dataset resources in the project.
	//  Each resource contains basic information.
	//  For full information about a particular dataset resource, use the Datasets:
	//  get method. This property is omitted when there are no datasets in the
	//  project.
	Datasets []ListFormatDataset `json:"datasets,omitempty"`

	// A list of skipped locations that were unreachable. For more information
	//  about BigQuery locations, see:
	//  https://cloud.google.com/bigquery/docs/locations. Example: "europe-west5"
	Unreachable []string `json:"unreachable,omitempty"`
}

// +kcc:proto=google.cloud.bigquery.v2.DatasetReference
type DatasetReference struct {
	// Required. A unique ID for this dataset, without the project name. The ID
	//  must contain only letters (a-z, A-Z), numbers (0-9), or underscores (_).
	//  The maximum length is 1,024 characters.
	DatasetId *string `json:"datasetId"`

	// Required. The ID of the project containing this dataset.
	ProjectId *string `json:"projectId"`
}

// +kcc:proto=google.cloud.bigquery.v2.DestinationTableProperties
type DestinationTableProperties struct {
	// Optional. Friendly name for the destination table. If the table already
	//  exists, it should be same as the existing friendly name.
	FriendlyName *string `json:"friendlyName,omitempty"`

	// Optional. The description for the destination table.
	//  This will only be used if the destination table is newly created.
	//  If the table already exists and a value different than the current
	//  description is provided, the job will fail.
	Description *string `json:"description,omitempty"`

	// Optional. The labels associated with this table. You can use these to
	//  organize and group your tables. This will only be used if the destination
	//  table is newly created. If the table already exists and labels are
	//  different than the current labels are provided, the job will fail.
	Labels map[string]string `json:"labels,omitempty"`
}

// +kcc:proto=google.cloud.bigquery.v2.DifferentialPrivacyPolicy
type DifferentialPrivacyPolicy struct {
	// Optional. The maximum epsilon value that a query can consume. If the
	//  subscriber specifies epsilon as a parameter in a SELECT query, it must be
	//  less than or equal to this value. The epsilon parameter controls the amount
	//  of noise that is added to the groups — a higher epsilon means less noise.
	MaxEpsilonPerQuery *float64 `json:"maxEpsilonPerQuery,omitempty"`

	// Optional. The delta value that is used per query. Delta represents the
	//  probability that any row will fail to be epsilon differentially private.
	//  Indicates the risk associated with exposing aggregate rows in the result of
	//  a query.
	DeltaPerQuery *float64 `json:"deltaPerQuery,omitempty"`

	// Optional. The maximum groups contributed value that is used per query.
	//  Represents the maximum number of groups to which each protected entity can
	//  contribute. Changing this value does not improve or worsen privacy. The
	//  best value for accuracy and utility depends on the query and data.
	MaxGroupsContributed *int64 `json:"maxGroupsContributed,omitempty"`

	// Optional. The privacy unit column associated with this policy. Differential
	//  privacy policies can only have one privacy unit column per data source
	//  object (table, view).
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
	DeltaBudget *float64 `json:"deltaBudget,omitempty"`

	// Output only. The epsilon budget remaining. If budget is exhausted, no more
	//  queries are allowed. Note that the budget for queries that are in progress
	//  is deducted before the query executes. If the query fails or is cancelled
	//  then the budget is refunded. In this case the amount of budget remaining
	//  can increase.
	EpsilonBudgetRemaining *float64 `json:"epsilonBudgetRemaining,omitempty"`

	// Output only. The delta budget remaining. If budget is exhausted, no more
	//  queries are allowed. Note that the budget for queries that are in progress
	//  is deducted before the query executes. If the query fails or is cancelled
	//  then the budget is refunded. In this case the amount of budget remaining
	//  can increase.
	DeltaBudgetRemaining *float64 `json:"deltaBudgetRemaining,omitempty"`
}

// +kcc:proto=google.cloud.bigquery.v2.DmlStats
type DmlStats struct {
	// Output only. Number of inserted Rows. Populated by DML INSERT and MERGE
	//  statements
	InsertedRowCount *int64 `json:"insertedRowCount,omitempty"`

	// Output only. Number of deleted Rows. populated by DML DELETE, MERGE and
	//  TRUNCATE statements.
	DeletedRowCount *int64 `json:"deletedRowCount,omitempty"`

	// Output only. Number of updated Rows. Populated by DML UPDATE and MERGE
	//  statements.
	UpdatedRowCount *int64 `json:"updatedRowCount,omitempty"`
}

// +kcc:proto=google.cloud.bigquery.v2.EncryptionConfiguration
type EncryptionConfiguration struct {
	// Optional. Describes the Cloud KMS encryption key that will be used to
	//  protect destination BigQuery table. The BigQuery Service Account associated
	//  with your project requires access to this encryption key.
	KmsKeyRef *refs.KMSCryptoKeyRef `json:"kmsKeyRef,omitempty"`
}

// +kcc:proto=google.cloud.bigquery.v2.ErrorProto
type ErrorProto struct {
	// A short error code that summarizes the error.
	Reason *string `json:"reason,omitempty"`

	// Specifies where the error occurred, if present.
	Location *string `json:"location,omitempty"`

	// Debugging information. This property is internal to Google and should not
	//  be used.
	DebugInfo *string `json:"debugInfo,omitempty"`

	// A human-readable description of the error.
	Message *string `json:"message,omitempty"`
}

// +kcc:proto=google.cloud.bigquery.v2.ExplainQueryStage
type ExplainQueryStage struct {
	// Human-readable name for the stage.
	Name *string `json:"name,omitempty"`

	// Unique ID for the stage within the plan.
	ID *int64 `json:"id,omitempty"`

	// Stage start time represented as milliseconds since the epoch.
	StartMs *int64 `json:"startMs,omitempty"`

	// Stage end time represented as milliseconds since the epoch.
	EndMs *int64 `json:"endMs,omitempty"`

	// IDs for stages that are inputs to this stage.
	InputStages []int64 `json:"inputStages,omitempty"`

	// Relative amount of time the average shard spent waiting to be
	//  scheduled.
	WaitRatioAvg *float64 `json:"waitRatioAvg,omitempty"`

	// Milliseconds the average shard spent waiting to be scheduled.
	WaitMsAvg *int64 `json:"waitMsAvg,omitempty"`

	// Relative amount of time the slowest shard spent waiting to be
	//  scheduled.
	WaitRatioMax *float64 `json:"waitRatioMax,omitempty"`

	// Milliseconds the slowest shard spent waiting to be scheduled.
	WaitMsMax *int64 `json:"waitMsMax,omitempty"`

	// Relative amount of time the average shard spent reading input.
	ReadRatioAvg *float64 `json:"readRatioAvg,omitempty"`

	// Milliseconds the average shard spent reading input.
	ReadMsAvg *int64 `json:"readMsAvg,omitempty"`

	// Relative amount of time the slowest shard spent reading input.
	ReadRatioMax *float64 `json:"readRatioMax,omitempty"`

	// Milliseconds the slowest shard spent reading input.
	ReadMsMax *int64 `json:"readMsMax,omitempty"`

	// Relative amount of time the average shard spent on CPU-bound tasks.
	ComputeRatioAvg *float64 `json:"computeRatioAvg,omitempty"`

	// Milliseconds the average shard spent on CPU-bound tasks.
	ComputeMsAvg *int64 `json:"computeMsAvg,omitempty"`

	// Relative amount of time the slowest shard spent on CPU-bound tasks.
	ComputeRatioMax *float64 `json:"computeRatioMax,omitempty"`

	// Milliseconds the slowest shard spent on CPU-bound tasks.
	ComputeMsMax *int64 `json:"computeMsMax,omitempty"`

	// Relative amount of time the average shard spent on writing output.
	WriteRatioAvg *float64 `json:"writeRatioAvg,omitempty"`

	// Milliseconds the average shard spent on writing output.
	WriteMsAvg *int64 `json:"writeMsAvg,omitempty"`

	// Relative amount of time the slowest shard spent on writing output.
	WriteRatioMax *float64 `json:"writeRatioMax,omitempty"`

	// Milliseconds the slowest shard spent on writing output.
	WriteMsMax *int64 `json:"writeMsMax,omitempty"`

	// Total number of bytes written to shuffle.
	ShuffleOutputBytes *int64 `json:"shuffleOutputBytes,omitempty"`

	// Total number of bytes written to shuffle and spilled to disk.
	ShuffleOutputBytesSpilled *int64 `json:"shuffleOutputBytesSpilled,omitempty"`

	// Number of records read into the stage.
	RecordsRead *int64 `json:"recordsRead,omitempty"`

	// Number of records written by the stage.
	RecordsWritten *int64 `json:"recordsWritten,omitempty"`

	// Number of parallel input segments to be processed
	ParallelInputs *int64 `json:"parallelInputs,omitempty"`

	// Number of parallel input segments completed.
	CompletedParallelInputs *int64 `json:"completedParallelInputs,omitempty"`

	// Current status for this stage.
	Status *string `json:"status,omitempty"`

	// List of operations within the stage in dependency order (approximately
	//  chronological).
	Steps []ExplainQueryStep `json:"steps,omitempty"`

	// Slot-milliseconds used by the stage.
	SlotMs *int64 `json:"slotMs,omitempty"`

	// Output only. Compute mode for this stage.
	ComputeMode *string `json:"computeMode,omitempty"`
}

// +kcc:proto=google.cloud.bigquery.v2.ExplainQueryStep
type ExplainQueryStep struct {
	// Machine-readable operation type.
	Kind *string `json:"kind,omitempty"`

	// Human-readable description of the step(s).
	Substeps []string `json:"substeps,omitempty"`
}

// +kcc:proto=google.cloud.bigquery.v2.ExportDataStatistics
type ExportDataStatistics struct {
	// Number of destination files generated in case of EXPORT DATA
	//  statement only.
	FileCount *int64 `json:"fileCount,omitempty"`

	// [Alpha] Number of destination rows generated in case of EXPORT DATA
	//  statement only.
	RowCount *int64 `json:"rowCount,omitempty"`
}

// +kcc:proto=google.cloud.bigquery.v2.ExternalCatalogDatasetOptions
type ExternalCatalogDatasetOptions struct {
	// Optional. A map of key value pairs defining the parameters and properties
	//  of the open source schema. Maximum size of 2Mib.
	Parameters map[string]string `json:"parameters,omitempty"`

	// Optional. The storage location URI for all tables in the dataset.
	//  Equivalent to hive metastore's database locationUri. Maximum length of 1024
	//  characters.
	DefaultStorageLocationUri *string `json:"defaultStorageLocationUri,omitempty"`
}

// +kcc:proto=google.cloud.bigquery.v2.ExternalCatalogTableOptions
type ExternalCatalogTableOptions struct {
	// Optional. A map of key value pairs defining the parameters and properties
	//  of the open source table. Corresponds with hive meta store table
	//  parameters. Maximum size of 4Mib.
	Parameters map[string]string `json:"parameters,omitempty"`

	// Optional. A storage descriptor containing information about the physical
	//  storage of this table.
	StorageDescriptor *StorageDescriptor `json:"storageDescriptor,omitempty"`

	// Optional. The connection specifying the credentials to be used to read
	//  external storage, such as Azure Blob, Cloud Storage, or S3. The connection
	//  is needed to read the open source table from BigQuery Engine. The
	//  connection_id can have the form
	//  `<project_id>.<location_id>.<connection_id>` or
	//  `projects/<project_id>/locations/<location_id>/connections/<connection_id>`.
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
	SourceUris []string `json:"sourceUris,omitempty"`

	// Optional. Specifies how source URIs are interpreted for constructing the
	//  file set to load.  By default source URIs are expanded against the
	//  underlying storage.  Other options include specifying manifest files. Only
	//  applicable to object storage systems.
	FileSetSpecType *string `json:"fileSetSpecType,omitempty"`

	// Optional. The schema for the data.
	//  Schema is required for CSV and JSON formats if autodetect is not on.
	//  Schema is disallowed for Google Cloud Bigtable, Cloud Datastore backups,
	//  Avro, ORC and Parquet formats.
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
	SourceFormat *string `json:"sourceFormat,omitempty"`

	// Optional. The maximum number of bad records that BigQuery can ignore when
	//  reading data. If the number of bad records exceeds this value, an invalid
	//  error is returned in the job result. The default value is 0, which requires
	//  that all records are valid. This setting is ignored for Google Cloud
	//  Bigtable, Google Cloud Datastore backups, Avro, ORC and Parquet formats.
	MaxBadRecords *int32 `json:"maxBadRecords,omitempty"`

	// Try to detect schema and format options automatically.
	//  Any option specified explicitly will be honored.
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
	IgnoreUnknownValues *bool `json:"ignoreUnknownValues,omitempty"`

	// Optional. The compression type of the data source.
	//  Possible values include GZIP and NONE. The default value is NONE.
	//  This setting is ignored for Google Cloud Bigtable, Google Cloud Datastore
	//  backups, Avro, ORC and Parquet
	//  formats. An empty string is an invalid value.
	Compression *string `json:"compression,omitempty"`

	// Optional. Additional properties to set if sourceFormat is set to CSV.
	CsvOptions *CsvOptions `json:"csvOptions,omitempty"`

	// Optional. Additional properties to set if sourceFormat is set to JSON.
	JsonOptions *JsonOptions `json:"jsonOptions,omitempty"`

	// Optional. Additional options if sourceFormat is set to BIGTABLE.
	BigtableOptions *BigtableOptions `json:"bigtableOptions,omitempty"`

	// Optional. Additional options if sourceFormat is set to GOOGLE_SHEETS.
	GoogleSheetsOptions *GoogleSheetsOptions `json:"googleSheetsOptions,omitempty"`

	// Optional. When set, configures hive partitioning support. Not all storage
	//  formats support hive partitioning -- requesting hive partitioning on an
	//  unsupported format will lead to an error, as will providing an invalid
	//  specification.
	HivePartitioningOptions *HivePartitioningOptions `json:"hivePartitioningOptions,omitempty"`

	// Optional. The connection specifying the credentials to be used to read
	//  external storage, such as Azure Blob, Cloud Storage, or S3. The
	//  connection_id can have the form
	//  `{project_id}.{location_id};{connection_id}` or
	//  `projects/{project_id}/locations/{location_id}/connections/{connection_id}`.
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
	DecimalTargetTypes []string `json:"decimalTargetTypes,omitempty"`

	// Optional. Additional properties to set if sourceFormat is set to AVRO.
	AvroOptions *AvroOptions `json:"avroOptions,omitempty"`

	// Optional. Load option to be used together with source_format
	//  newline-delimited JSON to indicate that a variant of JSON is being loaded.
	//  To load newline-delimited GeoJSON, specify GEOJSON (and source_format must
	//  be set to NEWLINE_DELIMITED_JSON).
	JsonExtension *string `json:"jsonExtension,omitempty"`

	// Optional. Additional properties to set if sourceFormat is set to PARQUET.
	ParquetOptions *ParquetOptions `json:"parquetOptions,omitempty"`

	// Optional. ObjectMetadata is used to create Object Tables. Object Tables
	//  contain a listing of objects (with their metadata) found at the
	//  source_uris. If ObjectMetadata is set, source_format should be omitted.
	//
	//  Currently SIMPLE is the only supported Object Metadata type.
	ObjectMetadata *string `json:"objectMetadata,omitempty"`

	// Optional. When creating an external table, the user can provide a reference
	//  file with the table schema. This is enabled for the following formats:
	//  AVRO, PARQUET, ORC.
	ReferenceFileSchemaUri *string `json:"referenceFileSchemaUri,omitempty"`

	// Optional. Metadata Cache Mode for the table. Set this to enable caching of
	//  metadata from external data source.
	MetadataCacheMode *string `json:"metadataCacheMode,omitempty"`
}

// +kcc:proto=google.cloud.bigquery.v2.ExternalDatasetReference
type ExternalDatasetReference struct {
	// Required. External source that backs this dataset.
	ExternalSource *string `json:"externalSource"`

	// Required. The connection id that is used to access the external_source.
	//
	//  Format:
	//    projects/{project_id}/locations/{location_id}/connections/{connection_id}
	Connection *string `json:"connection"`
}

// +kcc:proto=google.cloud.bigquery.v2.ExternalServiceCost
type ExternalServiceCost struct {
	// External service name.
	ExternalService *string `json:"externalService,omitempty"`

	// External service cost in terms of bigquery bytes processed.
	BytesProcessed *int64 `json:"bytesProcessed,omitempty"`

	// External service cost in terms of bigquery bytes billed.
	BytesBilled *int64 `json:"bytesBilled,omitempty"`

	// External service cost in terms of bigquery slot milliseconds.
	SlotMs *int64 `json:"slotMs,omitempty"`

	// Non-preemptable reserved slots used for external job.
	//  For example, reserved slots for Cloua AI Platform job are the VM usages
	//  converted to BigQuery slot with equivalent mount of price.
	ReservedSlotCount *int64 `json:"reservedSlotCount,omitempty"`
}

// +kcc:proto=google.cloud.bigquery.v2.ForeignKey
type ForeignKey struct {
	// Optional. Set only if the foreign key constraint is named.
	Name *string `json:"name,omitempty"`

	// Required. The table that holds the primary key and is referenced by this
	//  foreign key.
	ReferencedTable *TableReference `json:"referencedTable,omitempty"`

	// Required. The columns that compose the foreign key.
	ColumnReferences []ColumnReference `json:"columnReferences,omitempty"`
}

// +kcc:proto=google.cloud.bigquery.v2.ForeignTypeInfo
type ForeignTypeInfo struct {
	// Required. Specifies the system which defines the foreign data type.
	TypeSystem *string `json:"typeSystem,omitempty"`
}

// +kcc:proto=google.cloud.bigquery.v2.ForeignViewDefinition
type ForeignViewDefinition struct {
	// Required. The query that defines the view.
	Query *string `json:"query,omitempty"`

	// Optional. Represents the dialect of the query.
	Dialect *string `json:"dialect,omitempty"`
}

// +kcc:proto=google.cloud.bigquery.v2.GcpTag
type GcpTag struct {
	// Required. The namespaced friendly name of the tag key, e.g.
	//  "12345/environment" where 12345 is org id.
	TagKey *string `json:"tagKey,omitempty"`

	// Required. The friendly short name of the tag value, e.g. "production".
	TagValue *string `json:"tagValue,omitempty"`
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
	SkipLeadingRows *int64 `json:"skipLeadingRows,omitempty"`

	// Optional. Range of a sheet to query from. Only used when non-empty.
	//  Typical format: sheet_name!top_left_cell_id:bottom_right_cell_id
	//  For example: sheet1!A1:B20
	Range *string `json:"range,omitempty"`
}

// +kcc:proto=google.cloud.bigquery.v2.HighCardinalityJoin
type HighCardinalityJoin struct {
	// Output only. Count of left input rows.
	LeftRows *int64 `json:"leftRows,omitempty"`

	// Output only. Count of right input rows.
	RightRows *int64 `json:"rightRows,omitempty"`

	// Output only. Count of the output rows.
	OutputRows *int64 `json:"outputRows,omitempty"`

	// Output only. The index of the join operator in the ExplainQueryStep lists.
	StepIndex *int32 `json:"stepIndex,omitempty"`
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
	SourceUriPrefix *string `json:"sourceUriPrefix,omitempty"`

	// Optional. If set to true, queries over this table require a partition
	//  filter that can be used for partition elimination to be specified.
	//
	//  Note that this field should only be true when creating a permanent
	//  external table or querying a temporary external table.
	//
	//  Hive-partitioned loads with require_partition_filter explicitly set to
	//  true will fail.
	RequirePartitionFilter *bool `json:"requirePartitionFilter,omitempty"`

	// Output only. For permanent external tables, this field is populated with
	//  the hive partition keys in the order they were inferred. The types of the
	//  partition keys can be deduced by checking the table schema (which will
	//  include the partition keys). Not every API will populate this field in the
	//  output. For example, Tables.Get will populate it, but Tables.List will not
	//  contain this field.
	Fields []string `json:"fields,omitempty"`
}

// +kcc:proto=google.cloud.bigquery.v2.IndexUnusedReason
type IndexUnusedReason struct {
	// Specifies the high-level reason for the scenario when no search index was
	//  used.
	Code *string `json:"code,omitempty"`

	// Free form human-readable reason for the scenario when no search index was
	//  used.
	Message *string `json:"message,omitempty"`

	// Specifies the base table involved in the reason that no search index was
	//  used.
	BaseTable *TableReference `json:"baseTable,omitempty"`

	// Specifies the name of the unused search index, if available.
	IndexName *string `json:"indexName,omitempty"`
}

// +kcc:proto=google.cloud.bigquery.v2.InputDataChange
type InputDataChange struct {
	// Output only. Records read difference percentage compared to a previous run.
	RecordsReadDiffPercentage *float64 `json:"recordsReadDiffPercentage,omitempty"`
}

// +kcc:proto=google.cloud.bigquery.v2.Job
type Job struct {
	// Output only. The type of the resource.
	Kind *string `json:"kind,omitempty"`

	// Output only. A hash of this resource.
	Etag *string `json:"etag,omitempty"`

	// Output only. Opaque ID field of the job.
	ID *string `json:"id,omitempty"`

	// Output only. A URL that can be used to access the resource again.
	SelfLink *string `json:"selfLink,omitempty"`

	// Output only. Email address of the user who ran the job.
	UserEmail *string `json:"userEmail,omitempty"`

	// Required. Describes the job configuration.
	Configuration *JobConfiguration `json:"configuration,omitempty"`

	// Optional. Reference describing the unique-per-user name of the job.
	JobReference *JobReference `json:"jobReference,omitempty"`

	// Output only. Information about the job, including starting time and ending
	//  time of the job.
	Statistics *JobStatistics `json:"statistics,omitempty"`

	// Output only. The status of this job. Examine this value when polling an
	//  asynchronous job to see if the job is complete.
	Status *JobStatus `json:"status,omitempty"`

	// Output only. [Full-projection-only] String representation of identity of
	//  requesting party. Populated for both first- and third-party identities.
	//  Only present for APIs that support third-party identities.
	PrincipalSubject *string `json:"principalSubject,omitempty"`

	// Output only. The reason why a Job was created.
	//  [Preview](https://cloud.google.com/products/#product-launch-stages)
	JobCreationReason *JobCreationReason `json:"jobCreationReason,omitempty"`
}

// +kcc:proto=google.cloud.bigquery.v2.JobConfiguration
type JobConfiguration struct {
	// Output only. The type of the job. Can be QUERY, LOAD, EXTRACT, COPY or
	//  UNKNOWN.
	JobType *string `json:"jobType,omitempty"`

	// [Pick one] Configures a query job.
	Query *JobConfigurationQuery `json:"query,omitempty"`

	// [Pick one] Configures a load job.
	Load *JobConfigurationLoad `json:"load,omitempty"`

	// [Pick one] Copies a table.
	Copy *JobConfigurationTableCopy `json:"copy,omitempty"`

	// [Pick one] Configures an extract job.
	Extract *JobConfigurationExtract `json:"extract,omitempty"`

	// Optional. If set, don't actually run this job. A valid query will return
	//  a mostly empty response with some processing statistics, while an invalid
	//  query will return the same error it would if it wasn't a dry run. Behavior
	//  of non-query jobs is undefined.
	DryRun *bool `json:"dryRun,omitempty"`

	// Optional. Job timeout in milliseconds. If this time limit is exceeded,
	//  BigQuery will attempt to stop a longer job, but may not always succeed in
	//  canceling it before the job completes. For example, a job that takes more
	//  than 60 seconds to complete has a better chance of being stopped than a job
	//  that takes 10 seconds to complete.
	JobTimeoutMs *int64 `json:"jobTimeoutMs,omitempty"`

	// The labels associated with this job. You can use these to organize and
	//  group your jobs.
	//  Label keys and values can be no longer than 63 characters, can only contain
	//  lowercase letters, numeric characters, underscores and dashes.
	//  International characters are allowed. Label values are optional.  Label
	//  keys must start with a letter and each label in the list must have a
	//  different key.
	Labels map[string]string `json:"labels,omitempty"`
}

// +kcc:proto=google.cloud.bigquery.v2.JobConfigurationExtract
type JobConfigurationExtract struct {
	// A reference to the table being exported.
	SourceTable *TableReference `json:"sourceTable,omitempty"`

	// A reference to the model being exported.
	SourceModel *ModelReference `json:"sourceModel,omitempty"`

	// [Pick one] A list of fully-qualified Google Cloud Storage URIs where the
	//  extracted table should be written.
	DestinationUris []string `json:"destinationUris,omitempty"`

	// Optional. Whether to print out a header row in the results.
	//  Default is true. Not applicable when extracting models.
	PrintHeader *bool `json:"printHeader,omitempty"`

	// Optional. When extracting data in CSV format, this defines the
	//  delimiter to use between fields in the exported data.
	//  Default is ','. Not applicable when extracting models.
	FieldDelimiter *string `json:"fieldDelimiter,omitempty"`

	// Optional. The exported file format. Possible values include CSV,
	//  NEWLINE_DELIMITED_JSON, PARQUET, or AVRO for tables and ML_TF_SAVED_MODEL
	//  or ML_XGBOOST_BOOSTER for models. The default value for tables is CSV.
	//  Tables with nested or repeated fields cannot be exported as CSV. The
	//  default value for models is ML_TF_SAVED_MODEL.
	DestinationFormat *string `json:"destinationFormat,omitempty"`

	// Optional. The compression type to use for exported files. Possible values
	//  include DEFLATE, GZIP, NONE, SNAPPY, and ZSTD. The default value is NONE.
	//  Not all compression formats are support for all file formats. DEFLATE is
	//  only supported for Avro. ZSTD is only supported for Parquet. Not applicable
	//  when extracting models.
	Compression *string `json:"compression,omitempty"`

	// Whether to use logical types when extracting to AVRO format. Not applicable
	//  when extracting models.
	UseAvroLogicalTypes *bool `json:"useAvroLogicalTypes,omitempty"`

	// Optional. Model extract options only applicable when extracting models.
	ModelExtractOptions *JobConfigurationExtract_ModelExtractOptions `json:"modelExtractOptions,omitempty"`
}

// +kcc:proto=google.cloud.bigquery.v2.JobConfigurationExtract.ModelExtractOptions
type JobConfigurationExtract_ModelExtractOptions struct {
	// The 1-based ID of the trial to be exported from a hyperparameter tuning
	//  model. If not specified, the trial with id =
	//  [Model](https://cloud.google.com/bigquery/docs/reference/rest/v2/models#resource:-model).defaultTrialId
	//  is exported. This field is ignored for models not trained with
	//  hyperparameter tuning.
	TrialID *int64 `json:"trialID,omitempty"`
}

// +kcc:proto=google.cloud.bigquery.v2.JobConfigurationLoad
type JobConfigurationLoad struct {
	// [Required] The fully-qualified URIs that point to your data in Google
	//  Cloud.
	//  For Google Cloud Storage URIs:
	//    Each URI can contain one '*' wildcard character and it must come after
	//    the 'bucket' name. Size limits related to load jobs apply to external
	//    data sources.
	//  For Google Cloud Bigtable URIs:
	//    Exactly one URI can be specified and it has be a fully specified and
	//    valid HTTPS URL for a Google Cloud Bigtable table.
	//  For Google Cloud Datastore backups:
	//   Exactly one URI can be specified. Also, the '*' wildcard character is not
	//   allowed.
	SourceUris []string `json:"sourceUris,omitempty"`

	// Optional. Specifies how source URIs are interpreted for constructing the
	//  file set to load. By default, source URIs are expanded against the
	//  underlying storage. You can also specify manifest files to control how the
	//  file set is constructed. This option is only applicable to object storage
	//  systems.
	FileSetSpecType *string `json:"fileSetSpecType,omitempty"`

	// Optional. The schema for the destination table. The schema can be
	//  omitted if the destination table already exists, or if you're loading data
	//  from Google Cloud Datastore.
	Schema *TableSchema `json:"schema,omitempty"`

	// [Required] The destination table to load the data into.
	DestinationTable *TableReference `json:"destinationTable,omitempty"`

	// Optional. [Experimental] Properties with which to create the destination
	//  table if it is new.
	DestinationTableProperties *DestinationTableProperties `json:"destinationTableProperties,omitempty"`

	// Optional. Specifies whether the job is allowed to create new tables.
	//  The following values are supported:
	//
	//  * CREATE_IF_NEEDED: If the table does not exist, BigQuery creates the
	//  table.
	//  * CREATE_NEVER: The table must already exist. If it does not,
	//  a 'notFound' error is returned in the job result.
	//  The default value is CREATE_IF_NEEDED.
	//  Creation, truncation and append actions occur as one atomic update
	//  upon job completion.
	CreateDisposition *string `json:"createDisposition,omitempty"`

	// Optional. Specifies the action that occurs if the destination table
	//  already exists. The following values are supported:
	//
	//  * WRITE_TRUNCATE:  If the table already exists, BigQuery overwrites the
	//  data, removes the constraints and uses the schema from the load job.
	//  * WRITE_APPEND: If the table already exists, BigQuery appends the data to
	//  the table.
	//  * WRITE_EMPTY: If the table already exists and contains data, a 'duplicate'
	//  error is returned in the job result.
	//
	//  The default value is WRITE_APPEND.
	//  Each action is atomic and only occurs if BigQuery is able to complete the
	//  job successfully.
	//  Creation, truncation and append actions occur as one atomic update
	//  upon job completion.
	WriteDisposition *string `json:"writeDisposition,omitempty"`

	// Optional. Specifies a string that represents a null value in a CSV file.
	//  For example, if you specify "\N", BigQuery interprets "\N" as a null value
	//  when loading a CSV file.
	//  The default value is the empty string. If you set this property to a custom
	//  value, BigQuery throws an error if an empty string is present for all data
	//  types except for STRING and BYTE. For STRING and BYTE columns, BigQuery
	//  interprets the empty string as an empty value.
	NullMarker *string `json:"nullMarker,omitempty"`

	// Optional. The separator character for fields in a CSV file. The separator
	//  is interpreted as a single byte. For files encoded in ISO-8859-1, any
	//  single character can be used as a separator. For files encoded in UTF-8,
	//  characters represented in decimal range 1-127 (U+0001-U+007F) can be used
	//  without any modification. UTF-8 characters encoded with multiple bytes
	//  (i.e. U+0080 and above) will have only the first byte used for separating
	//  fields. The remaining bytes will be treated as a part of the field.
	//  BigQuery also supports the escape sequence "\t" (U+0009) to specify a tab
	//  separator. The default value is comma (",", U+002C).
	FieldDelimiter *string `json:"fieldDelimiter,omitempty"`

	// Optional. The number of rows at the top of a CSV file that BigQuery will
	//  skip when loading the data. The default value is 0. This property is useful
	//  if you have header rows in the file that should be skipped. When autodetect
	//  is on, the behavior is the following:
	//
	//  * skipLeadingRows unspecified - Autodetect tries to detect headers in the
	//    first row. If they are not detected, the row is read as data. Otherwise
	//    data is read starting from the second row.
	//  * skipLeadingRows is 0 - Instructs autodetect that there are no headers and
	//    data should be read starting from the first row.
	//  * skipLeadingRows = N > 0 - Autodetect skips N-1 rows and tries to detect
	//    headers in row N. If headers are not detected, row N is just skipped.
	//    Otherwise row N is used to extract column names for the detected schema.
	SkipLeadingRows *int32 `json:"skipLeadingRows,omitempty"`

	// Optional. The character encoding of the data.
	//  The supported values are UTF-8, ISO-8859-1, UTF-16BE, UTF-16LE, UTF-32BE,
	//  and UTF-32LE. The default value is UTF-8. BigQuery decodes the data after
	//  the raw, binary data has been split using the values of the `quote` and
	//  `fieldDelimiter` properties.
	//
	//  If you don't specify an encoding, or if you specify a UTF-8 encoding when
	//  the CSV file is not UTF-8 encoded, BigQuery attempts to convert the data to
	//  UTF-8. Generally, your data loads successfully, but it may not match
	//  byte-for-byte what you expect. To avoid this, specify the correct encoding
	//  by using the `--encoding` flag.
	//
	//  If BigQuery can't convert a character other than the ASCII `0` character,
	//  BigQuery converts the character to the standard Unicode replacement
	//  character: &#65533;.
	Encoding *string `json:"encoding,omitempty"`

	// Optional. The value that is used to quote data sections in a CSV file.
	//  BigQuery converts the string to ISO-8859-1 encoding, and then uses the
	//  first byte of the encoded string to split the data in its raw, binary
	//  state.
	//  The default value is a float64-quote ('"').
	//  If your data does not contain quoted sections, set the property value to an
	//  empty string.
	//  If your data contains quoted newline characters, you must also set the
	//  allowQuotedNewlines property to true.
	//  To include the specific quote character within a quoted value, precede it
	//  with an additional matching quote character. For example, if you want to
	//  escape the default character  ' " ', use ' "" '.
	//  @default "
	Quote *string `json:"quote,omitempty"`

	// Optional. The maximum number of bad records that BigQuery can ignore when
	//  running the job. If the number of bad records exceeds this value, an
	//  invalid error is returned in the job result.
	//  The default value is 0, which requires that all records are valid.
	//  This is only supported for CSV and NEWLINE_DELIMITED_JSON file formats.
	MaxBadRecords *int32 `json:"maxBadRecords,omitempty"`

	// Indicates if BigQuery should allow quoted data sections that contain
	//  newline characters in a CSV file. The default value is false.
	AllowQuotedNewlines *bool `json:"allowQuotedNewlines,omitempty"`

	// Optional. The format of the data files.
	//  For CSV files, specify "CSV". For datastore backups,
	//  specify "DATASTORE_BACKUP". For newline-delimited JSON,
	//  specify "NEWLINE_DELIMITED_JSON". For Avro, specify "AVRO".
	//  For parquet, specify "PARQUET". For orc, specify "ORC".
	//  The default value is CSV.
	SourceFormat *string `json:"sourceFormat,omitempty"`

	// Optional. Accept rows that are missing trailing optional columns.
	//  The missing values are treated as nulls.
	//  If false, records with missing trailing columns are treated as bad records,
	//  and if there are too many bad records, an invalid error is returned in the
	//  job result.
	//  The default value is false.
	//  Only applicable to CSV, ignored for other formats.
	AllowJaggedRows *bool `json:"allowJaggedRows,omitempty"`

	// Optional. Indicates if BigQuery should allow extra values that are not
	//  represented in the table schema.
	//  If true, the extra values are ignored.
	//  If false, records with extra columns are treated as bad records, and if
	//  there are too many bad records, an invalid error is returned in the job
	//  result. The default value is false.
	//  The sourceFormat property determines what BigQuery treats as an extra
	//  value:
	//    CSV: Trailing columns
	//    JSON: Named values that don't match any column names in the table schema
	//    Avro, Parquet, ORC: Fields in the file schema that don't exist in the
	//    table schema.
	IgnoreUnknownValues *bool `json:"ignoreUnknownValues,omitempty"`

	// If sourceFormat is set to "DATASTORE_BACKUP", indicates which entity
	//  properties to load into BigQuery from a Cloud Datastore backup. Property
	//  names are case sensitive and must be top-level properties. If no properties
	//  are specified, BigQuery loads all properties. If any named property isn't
	//  found in the Cloud Datastore backup, an invalid error is returned in the
	//  job result.
	ProjectionFields []string `json:"projectionFields,omitempty"`

	// Optional. Indicates if we should automatically infer the options and
	//  schema for CSV and JSON sources.
	Autodetect *bool `json:"autodetect,omitempty"`

	// Allows the schema of the destination table to be updated as a side effect
	//  of the load job if a schema is autodetected or supplied in the job
	//  configuration.
	//  Schema update options are supported in two cases:
	//  when writeDisposition is WRITE_APPEND;
	//  when writeDisposition is WRITE_TRUNCATE and the destination table is a
	//  partition of a table, specified by partition decorators. For normal tables,
	//  WRITE_TRUNCATE will always overwrite the schema.
	//  One or more of the following values are specified:
	//
	//  * ALLOW_FIELD_ADDITION: allow adding a nullable field to the schema.
	//  * ALLOW_FIELD_RELAXATION: allow relaxing a required field in the original
	//  schema to nullable.
	SchemaUpdateOptions []string `json:"schemaUpdateOptions,omitempty"`

	// Time-based partitioning specification for the destination table. Only one
	//  of timePartitioning and rangePartitioning should be specified.
	TimePartitioning *TimePartitioning `json:"timePartitioning,omitempty"`

	// Range partitioning specification for the destination table.
	//  Only one of timePartitioning and rangePartitioning should be specified.
	RangePartitioning *RangePartitioning `json:"rangePartitioning,omitempty"`

	// Clustering specification for the destination table.
	Clustering *Clustering `json:"clustering,omitempty"`

	// Custom encryption configuration (e.g., Cloud KMS keys)
	DestinationEncryptionConfiguration *EncryptionConfiguration `json:"destinationEncryptionConfiguration,omitempty"`

	// Optional. If sourceFormat is set to "AVRO", indicates whether to interpret
	//  logical types as the corresponding BigQuery data type (for example,
	//  TIMESTAMP), instead of using the raw type (for example, INTEGER).
	UseAvroLogicalTypes *bool `json:"useAvroLogicalTypes,omitempty"`

	// Optional. The user can provide a reference file with the reader schema.
	//  This file is only loaded if it is part of source URIs, but is not loaded
	//  otherwise. It is enabled for the following formats: AVRO, PARQUET, ORC.
	ReferenceFileSchemaUri *string `json:"referenceFileSchemaUri,omitempty"`

	// Optional. When set, configures hive partitioning support.
	//  Not all storage formats support hive partitioning -- requesting hive
	//  partitioning on an unsupported format will lead to an error, as will
	//  providing an invalid specification.
	HivePartitioningOptions *HivePartitioningOptions `json:"hivePartitioningOptions,omitempty"`

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
	DecimalTargetTypes []string `json:"decimalTargetTypes,omitempty"`

	// Optional. Load option to be used together with source_format
	//  newline-delimited JSON to indicate that a variant of JSON is being loaded.
	//  To load newline-delimited GeoJSON, specify GEOJSON (and source_format must
	//  be set to NEWLINE_DELIMITED_JSON).
	JsonExtension *string `json:"jsonExtension,omitempty"`

	// Optional. Additional properties to set if sourceFormat is set to PARQUET.
	ParquetOptions *ParquetOptions `json:"parquetOptions,omitempty"`

	// Optional. When sourceFormat is set to "CSV", this indicates whether the
	//  embedded ASCII control characters (the first 32 characters in the
	//  ASCII-table, from
	//  '\x00' to '\x1F') are preserved.
	PreserveAsciiControlCharacters *bool `json:"preserveAsciiControlCharacters,omitempty"`

	// Optional. Connection properties which can modify the load job behavior.
	//  Currently, only the 'session_id' connection property is supported, and is
	//  used to resolve _SESSION appearing as the dataset id.
	ConnectionProperties []ConnectionProperty `json:"connectionProperties,omitempty"`

	// Optional. If this property is true, the job creates a new session using a
	//  randomly generated session_id.  To continue using a created session with
	//  subsequent queries, pass the existing session identifier as a
	//  `ConnectionProperty` value.  The session identifier is returned as part of
	//  the `SessionInfo` message within the query statistics.
	//
	//  The new session's location will be set to `Job.JobReference.location` if it
	//  is present, otherwise it's set to the default location based on existing
	//  routing logic.
	CreateSession *bool `json:"createSession,omitempty"`

	// Optional. Character map supported for column names in CSV/Parquet loads.
	//  Defaults to STRICT and can be overridden by Project Config Service. Using
	//  this option with unsupporting load formats will result in an error.
	ColumnNameCharacterMap *string `json:"columnNameCharacterMap,omitempty"`

	// Optional. [Experimental] Configures the load job to copy files directly to
	//  the destination BigLake managed table, bypassing file content reading and
	//  rewriting.
	//
	//  Copying files only is supported when all the following are true:
	//
	//  * `source_uris` are located in the same Cloud Storage location as the
	//    destination table's `storage_uri` location.
	//  * `source_format` is `PARQUET`.
	//  * `destination_table` is an existing BigLake managed table. The table's
	//    schema does not have flexible column names. The table's columns do not
	//    have type parameters other than precision and scale.
	//  * No options other than the above are specified.
	CopyFilesOnly *bool `json:"copyFilesOnly,omitempty"`
}

// +kcc:proto=google.cloud.bigquery.v2.JobConfigurationQuery
type JobConfigurationQuery struct {
	// [Required] SQL query text to execute. The useLegacySql field can be used
	//  to indicate whether the query uses legacy SQL or GoogleSQL.
	Query *string `json:"query,omitempty"`

	// Optional. Describes the table where the query results should be stored.
	//  This property must be set for large results that exceed the maximum
	//  response size.  For queries that produce anonymous (cached) results, this
	//  field will be populated by BigQuery.
	DestinationTable *TableReference `json:"destinationTable,omitempty"`

	// TODO: map type string message for external_table_definitions

	// Describes user-defined function resources used in the query.
	UserDefinedFunctionResources []UserDefinedFunctionResource `json:"userDefinedFunctionResources,omitempty"`

	// Optional. Specifies whether the job is allowed to create new tables.
	//  The following values are supported:
	//
	//  * CREATE_IF_NEEDED: If the table does not exist, BigQuery creates the
	//  table.
	//  * CREATE_NEVER: The table must already exist. If it does not,
	//  a 'notFound' error is returned in the job result.
	//
	//  The default value is CREATE_IF_NEEDED.
	//  Creation, truncation and append actions occur as one atomic update
	//  upon job completion.
	CreateDisposition *string `json:"createDisposition,omitempty"`

	// Optional. Specifies the action that occurs if the destination table
	//  already exists. The following values are supported:
	//
	//  * WRITE_TRUNCATE: If the table already exists, BigQuery overwrites the
	//  data, removes the constraints, and uses the schema from the query result.
	//  * WRITE_APPEND: If the table already exists, BigQuery appends the data to
	//  the table.
	//  * WRITE_EMPTY: If the table already exists and contains data, a 'duplicate'
	//  error is returned in the job result.
	//
	//  The default value is WRITE_EMPTY. Each action is atomic and only occurs if
	//  BigQuery is able to complete the job successfully. Creation, truncation and
	//  append actions occur as one atomic update upon job completion.
	WriteDisposition *string `json:"writeDisposition,omitempty"`

	// Optional. Specifies the default dataset to use for unqualified
	//  table names in the query. This setting does not alter behavior of
	//  unqualified dataset names. Setting the system variable
	//  `@@dataset_id` achieves the same behavior.  See
	//  https://cloud.google.com/bigquery/docs/reference/system-variables for more
	//  information on system variables.
	DefaultDataset *DatasetReference `json:"defaultDataset,omitempty"`

	// Optional. Specifies a priority for the query. Possible values include
	//  INTERACTIVE and BATCH. The default value is INTERACTIVE.
	Priority *string `json:"priority,omitempty"`

	// Optional. If true and query uses legacy SQL dialect, allows the query
	//  to produce arbitrarily large result tables at a slight cost in performance.
	//  Requires destinationTable to be set.
	//  For GoogleSQL queries, this flag is ignored and large results are
	//  always allowed.  However, you must still set destinationTable when result
	//  size exceeds the allowed maximum response size.
	AllowLargeResults *bool `json:"allowLargeResults,omitempty"`

	// Optional. Whether to look for the result in the query cache. The query
	//  cache is a best-effort cache that will be flushed whenever tables in the
	//  query are modified. Moreover, the query cache is only available when a
	//  query does not have a destination table specified. The default value is
	//  true.
	UseQueryCache *bool `json:"useQueryCache,omitempty"`

	// Optional. If true and query uses legacy SQL dialect, flattens all nested
	//  and repeated fields in the query results.
	//  allowLargeResults must be true if this is set to false.
	//  For GoogleSQL queries, this flag is ignored and results are never
	//  flattened.
	FlattenResults *bool `json:"flattenResults,omitempty"`

	// Limits the bytes billed for this job. Queries that will have
	//  bytes billed beyond this limit will fail (without incurring a charge).
	//  If unspecified, this will be set to your project default.
	MaximumBytesBilled *int64 `json:"maximumBytesBilled,omitempty"`

	// Optional. Specifies whether to use BigQuery's legacy SQL dialect for this
	//  query. The default value is true. If set to false, the query will use
	//  BigQuery's GoogleSQL:
	//  https://cloud.google.com/bigquery/sql-reference/
	//
	//  When useLegacySql is set to false, the value of flattenResults is ignored;
	//  query will be run as if flattenResults is false.
	UseLegacySql *bool `json:"useLegacySql,omitempty"`

	// GoogleSQL only. Set to POSITIONAL to use positional (?) query parameters
	//  or to NAMED to use named (@myparam) query parameters in this query.
	ParameterMode *string `json:"parameterMode,omitempty"`

	// Query parameters for GoogleSQL queries.
	QueryParameters []QueryParameter `json:"queryParameters,omitempty"`

	// Output only. System variables for GoogleSQL queries. A system variable is
	//  output if the variable is settable and its value differs from the system
	//  default.
	//  "@@" prefix is not included in the name of the System variables.
	SystemVariables *SystemVariables `json:"systemVariables,omitempty"`

	// Allows the schema of the destination table to be updated as a side effect
	//  of the query job. Schema update options are supported in two cases:
	//  when writeDisposition is WRITE_APPEND;
	//  when writeDisposition is WRITE_TRUNCATE and the destination table is a
	//  partition of a table, specified by partition decorators. For normal tables,
	//  WRITE_TRUNCATE will always overwrite the schema.
	//  One or more of the following values are specified:
	//
	//  * ALLOW_FIELD_ADDITION: allow adding a nullable field to the schema.
	//  * ALLOW_FIELD_RELAXATION: allow relaxing a required field in the original
	//  schema to nullable.
	SchemaUpdateOptions []string `json:"schemaUpdateOptions,omitempty"`

	// Time-based partitioning specification for the destination table. Only one
	//  of timePartitioning and rangePartitioning should be specified.
	TimePartitioning *TimePartitioning `json:"timePartitioning,omitempty"`

	// Range partitioning specification for the destination table.
	//  Only one of timePartitioning and rangePartitioning should be specified.
	RangePartitioning *RangePartitioning `json:"rangePartitioning,omitempty"`

	// Clustering specification for the destination table.
	Clustering *Clustering `json:"clustering,omitempty"`

	// Custom encryption configuration (e.g., Cloud KMS keys)
	DestinationEncryptionConfiguration *EncryptionConfiguration `json:"destinationEncryptionConfiguration,omitempty"`

	// Options controlling the execution of scripts.
	ScriptOptions *ScriptOptions `json:"scriptOptions,omitempty"`

	// Connection properties which can modify the query behavior.
	ConnectionProperties []ConnectionProperty `json:"connectionProperties,omitempty"`

	// If this property is true, the job creates a new session using a randomly
	//  generated session_id.  To continue using a created session with
	//  subsequent queries, pass the existing session identifier as a
	//  `ConnectionProperty` value.  The session identifier is returned as part of
	//  the `SessionInfo` message within the query statistics.
	//
	//  The new session's location will be set to `Job.JobReference.location` if it
	//  is present, otherwise it's set to the default location based on existing
	//  routing logic.
	CreateSession *bool `json:"createSession,omitempty"`

	// Optional. Whether to run the query as continuous or a regular query.
	//  Continuous query is currently in experimental stage and not ready for
	//  general usage.
	Continuous *bool `json:"continuous,omitempty"`
}

// +kcc:proto=google.cloud.bigquery.v2.JobConfigurationTableCopy
type JobConfigurationTableCopy struct {
	// [Pick one] Source table to copy.
	SourceTable *TableReference `json:"sourceTable,omitempty"`

	// [Pick one] Source tables to copy.
	SourceTables []TableReference `json:"sourceTables,omitempty"`

	// [Required] The destination table.
	DestinationTable *TableReference `json:"destinationTable,omitempty"`

	// Optional. Specifies whether the job is allowed to create new tables.
	//  The following values are supported:
	//
	//  * CREATE_IF_NEEDED: If the table does not exist, BigQuery creates the
	//  table.
	//  * CREATE_NEVER: The table must already exist. If it does not,
	//  a 'notFound' error is returned in the job result.
	//
	//  The default value is CREATE_IF_NEEDED.
	//  Creation, truncation and append actions occur as one atomic update
	//  upon job completion.
	CreateDisposition *string `json:"createDisposition,omitempty"`

	// Optional. Specifies the action that occurs if the destination table
	//  already exists. The following values are supported:
	//
	//  * WRITE_TRUNCATE: If the table already exists, BigQuery overwrites the
	//  table data and uses the schema and table constraints from the source table.
	//  * WRITE_APPEND: If the table already exists, BigQuery appends the data to
	//  the table.
	//  * WRITE_EMPTY: If the table already exists and contains data, a 'duplicate'
	//  error is returned in the job result.
	//
	//  The default value is WRITE_EMPTY. Each action is atomic and only occurs if
	//  BigQuery is able to complete the job successfully. Creation, truncation and
	//  append actions occur as one atomic update upon job completion.
	WriteDisposition *string `json:"writeDisposition,omitempty"`

	// Custom encryption configuration (e.g., Cloud KMS keys).
	DestinationEncryptionConfiguration *EncryptionConfiguration `json:"destinationEncryptionConfiguration,omitempty"`

	// Optional. Supported operation types in table copy job.
	OperationType *string `json:"operationType,omitempty"`

	// Optional. The time when the destination table expires. Expired tables will
	//  be deleted and their storage reclaimed.
	DestinationExpirationTime *string `json:"destinationExpirationTime,omitempty"`
}

// +kcc:proto=google.cloud.bigquery.v2.JobCreationReason
type JobCreationReason struct {
	// Output only. Specifies the high level reason why a Job was created.
	Code *string `json:"code,omitempty"`
}

// +kcc:proto=google.cloud.bigquery.v2.JobList
type JobList struct {
	// A hash of this page of results.
	Etag *string `json:"etag,omitempty"`

	// The resource type of the response.
	Kind *string `json:"kind,omitempty"`

	// A token to request the next page of results.
	NextPageToken *string `json:"nextPageToken,omitempty"`

	// List of jobs that were requested.
	Jobs []ListFormatJob `json:"jobs,omitempty"`

	// A list of skipped locations that were unreachable. For more information
	//  about BigQuery locations, see:
	//  https://cloud.google.com/bigquery/docs/locations. Example: "europe-west5"
	Unreachable []string `json:"unreachable,omitempty"`
}

// +kcc:proto=google.cloud.bigquery.v2.JobReference
type JobReference struct {
	// Required. The ID of the project containing this job.
	ProjectID *string `json:"projectID,omitempty"`

	// Required. The ID of the job. The ID must contain only letters (a-z, A-Z),
	//  numbers (0-9), underscores (_), or dashes (-). The maximum length is 1,024
	//  characters.
	JobID *string `json:"jobID,omitempty"`

	// Optional. The geographic location of the job. The default value is US.
	//
	//  For more information about BigQuery locations, see:
	//  https://cloud.google.com/bigquery/docs/locations
	Location *string `json:"location,omitempty"`

	// This field should not be used.
	LocationAlternative []string `json:"locationAlternative,omitempty"`
}

// +kcc:proto=google.cloud.bigquery.v2.JobStatistics
type JobStatistics struct {
	// Output only. Creation time of this job, in milliseconds since the epoch.
	//  This field will be present on all jobs.
	CreationTime *int64 `json:"creationTime,omitempty"`

	// Output only. Start time of this job, in milliseconds since the epoch.
	//  This field will be present when the job transitions from the PENDING state
	//  to either RUNNING or DONE.
	StartTime *int64 `json:"startTime,omitempty"`

	// Output only. End time of this job, in milliseconds since the epoch. This
	//  field will be present whenever a job is in the DONE state.
	EndTime *int64 `json:"endTime,omitempty"`

	// Output only. Total bytes processed for the job.
	TotalBytesProcessed *int64 `json:"totalBytesProcessed,omitempty"`

	// Output only. [TrustedTester] Job progress (0.0 -> 1.0) for LOAD and
	//  EXTRACT jobs.
	CompletionRatio *float64 `json:"completionRatio,omitempty"`

	// Output only. Quotas which delayed this job's start time.
	QuotaDeferments []string `json:"quotaDeferments,omitempty"`

	// Output only. Statistics for a query job.
	Query *JobStatistics2 `json:"query,omitempty"`

	// Output only. Statistics for a load job.
	Load *JobStatistics3 `json:"load,omitempty"`

	// Output only. Statistics for an extract job.
	Extract *JobStatistics4 `json:"extract,omitempty"`

	// Output only. Statistics for a copy job.
	Copy *CopyJobStatistics `json:"copy,omitempty"`

	// Output only. Slot-milliseconds for the job.
	TotalSlotMs *int64 `json:"totalSlotMs,omitempty"`

	// Output only. Name of the primary reservation assigned to this job. Note
	//  that this could be different than reservations reported in the reservation
	//  usage field if parent reservations were used to execute this job.
	ReservationID *string `json:"reservationID,omitempty"`

	// Output only. Number of child jobs executed.
	NumChildJobs *int64 `json:"numChildJobs,omitempty"`

	// Output only. If this is a child job, specifies the job ID of the parent.
	ParentJobID *string `json:"parentJobID,omitempty"`

	// Output only. If this a child job of a script, specifies information about
	//  the context of this job within the script.
	ScriptStatistics *ScriptStatistics `json:"scriptStatistics,omitempty"`

	// Output only. Statistics for row-level security. Present only for query and
	//  extract jobs.
	RowLevelSecurityStatistics *RowLevelSecurityStatistics `json:"rowLevelSecurityStatistics,omitempty"`

	// Output only. Statistics for data-masking. Present only for query and
	//  extract jobs.
	DataMaskingStatistics *DataMaskingStatistics `json:"dataMaskingStatistics,omitempty"`

	// Output only. [Alpha] Information of the multi-statement transaction if this
	//  job is part of one.
	//
	//  This property is only expected on a child job or a job that is in a
	//  session. A script parent job is not part of the transaction started in the
	//  script.
	TransactionInfo *JobStatistics_TransactionInfo `json:"transactionInfo,omitempty"`

	// Output only. Information of the session if this job is part of one.
	SessionInfo *SessionInfo `json:"sessionInfo,omitempty"`

	// Output only. The duration in milliseconds of the execution of the final
	//  attempt of this job, as BigQuery may internally re-attempt to execute the
	//  job.
	FinalExecutionDurationMs *int64 `json:"finalExecutionDurationMs,omitempty"`

	// Output only. Name of edition corresponding to the reservation for this job
	//  at the time of this update.
	Edition *string `json:"edition,omitempty"`
}

// +kcc:proto=google.cloud.bigquery.v2.JobStatistics.TransactionInfo
type JobStatistics_TransactionInfo struct {
	// Output only. [Alpha] Id of the transaction.
	TransactionID *string `json:"transactionID,omitempty"`
}

// +kcc:proto=google.cloud.bigquery.v2.JobStatistics2
type JobStatistics2 struct {
	// Output only. Describes execution plan for the query.
	QueryPlan []ExplainQueryStage `json:"queryPlan,omitempty"`

	// Output only. The original estimate of bytes processed for the job.
	EstimatedBytesProcessed *int64 `json:"estimatedBytesProcessed,omitempty"`

	// Output only. Describes a timeline of job execution.
	Timeline []QueryTimelineSample `json:"timeline,omitempty"`

	// Output only. Total number of partitions processed from all partitioned
	//  tables referenced in the job.
	TotalPartitionsProcessed *int64 `json:"totalPartitionsProcessed,omitempty"`

	// Output only. Total bytes processed for the job.
	TotalBytesProcessed *int64 `json:"totalBytesProcessed,omitempty"`

	// Output only. For dry-run jobs, totalBytesProcessed is an estimate and this
	//  field specifies the accuracy of the estimate. Possible values can be:
	//  UNKNOWN: accuracy of the estimate is unknown.
	//  PRECISE: estimate is precise.
	//  LOWER_BOUND: estimate is lower bound of what the query would cost.
	//  UPPER_BOUND: estimate is upper bound of what the query would cost.
	TotalBytesProcessedAccuracy *string `json:"totalBytesProcessedAccuracy,omitempty"`

	// Output only. If the project is configured to use on-demand pricing,
	//  then this field contains the total bytes billed for the job.
	//  If the project is configured to use flat-rate pricing, then you are
	//  not billed for bytes and this field is informational only.
	TotalBytesBilled *int64 `json:"totalBytesBilled,omitempty"`

	// Output only. Billing tier for the job. This is a BigQuery-specific concept
	//  which is not related to the Google Cloud notion of "free tier". The value
	//  here is a measure of the query's resource consumption relative to the
	//  amount of data scanned. For on-demand queries, the limit is 100, and all
	//  queries within this limit are billed at the standard on-demand rates.
	//  On-demand queries that exceed this limit will fail with a
	//  billingTierLimitExceeded error.
	BillingTier *int32 `json:"billingTier,omitempty"`

	// Output only. Slot-milliseconds for the job.
	TotalSlotMs *int64 `json:"totalSlotMs,omitempty"`

	// Output only. Whether the query result was fetched from the query cache.
	CacheHit *bool `json:"cacheHit,omitempty"`

	// Output only. Referenced tables for the job. Queries that reference more
	//  than 50 tables will not have a complete list.
	ReferencedTables []TableReference `json:"referencedTables,omitempty"`

	// Output only. Referenced routines for the job.
	ReferencedRoutines []RoutineReference `json:"referencedRoutines,omitempty"`

	// Output only. The schema of the results. Present only for successful dry
	//  run of non-legacy SQL queries.
	Schema *TableSchema `json:"schema,omitempty"`

	// Output only. The number of rows affected by a DML statement. Present
	//  only for DML statements INSERT, UPDATE or DELETE.
	NumDmlAffectedRows *int64 `json:"numDmlAffectedRows,omitempty"`

	// Output only. Detailed statistics for DML statements INSERT, UPDATE, DELETE,
	//  MERGE or TRUNCATE.
	DmlStats *DmlStats `json:"dmlStats,omitempty"`

	// Output only. GoogleSQL only: list of undeclared query
	//  parameters detected during a dry run validation.
	UndeclaredQueryParameters []QueryParameter `json:"undeclaredQueryParameters,omitempty"`

	// Output only. The type of query statement, if valid.
	//  Possible values:
	//
	//  * `SELECT`:
	//  [`SELECT`](https://cloud.google.com/bigquery/docs/reference/standard-sql/query-syntax#select_list)
	//  statement.
	//  * `ASSERT`:
	//  [`ASSERT`](https://cloud.google.com/bigquery/docs/reference/standard-sql/debugging-statements#assert)
	//  statement.
	//  * `INSERT`:
	//  [`INSERT`](https://cloud.google.com/bigquery/docs/reference/standard-sql/dml-syntax#insert_statement)
	//  statement.
	//  * `UPDATE`:
	//  [`UPDATE`](https://cloud.google.com/bigquery/docs/reference/standard-sql/query-syntax#update_statement)
	//  statement.
	//  * `DELETE`:
	//  [`DELETE`](https://cloud.google.com/bigquery/docs/reference/standard-sql/data-manipulation-language)
	//  statement.
	//  * `MERGE`:
	//  [`MERGE`](https://cloud.google.com/bigquery/docs/reference/standard-sql/data-manipulation-language)
	//  statement.
	//  * `CREATE_TABLE`: [`CREATE
	//  TABLE`](https://cloud.google.com/bigquery/docs/reference/standard-sql/data-definition-language#create_table_statement)
	//  statement, without `AS SELECT`.
	//  * `CREATE_TABLE_AS_SELECT`: [`CREATE TABLE AS
	//  SELECT`](https://cloud.google.com/bigquery/docs/reference/standard-sql/data-definition-language#query_statement)
	//  statement.
	//  * `CREATE_VIEW`: [`CREATE
	//  VIEW`](https://cloud.google.com/bigquery/docs/reference/standard-sql/data-definition-language#create_view_statement)
	//  statement.
	//  * `CREATE_MODEL`: [`CREATE
	//  MODEL`](https://cloud.google.com/bigquery-ml/docs/reference/standard-sql/bigqueryml-syntax-create#create_model_statement)
	//  statement.
	//  * `CREATE_MATERIALIZED_VIEW`: [`CREATE MATERIALIZED
	//  VIEW`](https://cloud.google.com/bigquery/docs/reference/standard-sql/data-definition-language#create_materialized_view_statement)
	//  statement.
	//  * `CREATE_FUNCTION`: [`CREATE
	//  FUNCTION`](https://cloud.google.com/bigquery/docs/reference/standard-sql/data-definition-language#create_function_statement)
	//  statement.
	//  * `CREATE_TABLE_FUNCTION`: [`CREATE TABLE
	//  FUNCTION`](https://cloud.google.com/bigquery/docs/reference/standard-sql/data-definition-language#create_table_function_statement)
	//  statement.
	//  * `CREATE_PROCEDURE`: [`CREATE
	//  PROCEDURE`](https://cloud.google.com/bigquery/docs/reference/standard-sql/data-definition-language#create_procedure)
	//  statement.
	//  * `CREATE_ROW_ACCESS_POLICY`: [`CREATE ROW ACCESS
	//  POLICY`](https://cloud.google.com/bigquery/docs/reference/standard-sql/data-definition-language#create_row_access_policy_statement)
	//  statement.
	//  * `CREATE_SCHEMA`: [`CREATE
	//  SCHEMA`](https://cloud.google.com/bigquery/docs/reference/standard-sql/data-definition-language#create_schema_statement)
	//  statement.
	//  * `CREATE_SNAPSHOT_TABLE`: [`CREATE SNAPSHOT
	//  TABLE`](https://cloud.google.com/bigquery/docs/reference/standard-sql/data-definition-language#create_snapshot_table_statement)
	//  statement.
	//  * `CREATE_SEARCH_INDEX`: [`CREATE SEARCH
	//  INDEX`](https://cloud.google.com/bigquery/docs/reference/standard-sql/data-definition-language#create_search_index_statement)
	//  statement.
	//  * `DROP_TABLE`: [`DROP
	//  TABLE`](https://cloud.google.com/bigquery/docs/reference/standard-sql/data-definition-language#drop_table_statement)
	//  statement.
	//  * `DROP_EXTERNAL_TABLE`: [`DROP EXTERNAL
	//  TABLE`](https://cloud.google.com/bigquery/docs/reference/standard-sql/data-definition-language#drop_external_table_statement)
	//  statement.
	//  * `DROP_VIEW`: [`DROP
	//  VIEW`](https://cloud.google.com/bigquery/docs/reference/standard-sql/data-definition-language#drop_view_statement)
	//  statement.
	//  * `DROP_MODEL`: [`DROP
	//  MODEL`](https://cloud.google.com/bigquery-ml/docs/reference/standard-sql/bigqueryml-syntax-drop-model)
	//  statement.
	//  * `DROP_MATERIALIZED_VIEW`: [`DROP MATERIALIZED
	//   VIEW`](https://cloud.google.com/bigquery/docs/reference/standard-sql/data-definition-language#drop_materialized_view_statement)
	//  statement.
	//  * `DROP_FUNCTION` : [`DROP
	//  FUNCTION`](https://cloud.google.com/bigquery/docs/reference/standard-sql/data-definition-language#drop_function_statement)
	//  statement.
	//  * `DROP_TABLE_FUNCTION` : [`DROP TABLE
	//  FUNCTION`](https://cloud.google.com/bigquery/docs/reference/standard-sql/data-definition-language#drop_table_function)
	//  statement.
	//  * `DROP_PROCEDURE`: [`DROP
	//  PROCEDURE`](https://cloud.google.com/bigquery/docs/reference/standard-sql/data-definition-language#drop_procedure_statement)
	//  statement.
	//  * `DROP_SEARCH_INDEX`: [`DROP SEARCH
	//  INDEX`](https://cloud.google.com/bigquery/docs/reference/standard-sql/data-definition-language#drop_search_index)
	//  statement.
	//  * `DROP_SCHEMA`: [`DROP
	//  SCHEMA`](https://cloud.google.com/bigquery/docs/reference/standard-sql/data-definition-language#drop_schema_statement)
	//  statement.
	//  * `DROP_SNAPSHOT_TABLE`: [`DROP SNAPSHOT
	//  TABLE`](https://cloud.google.com/bigquery/docs/reference/standard-sql/data-definition-language#drop_snapshot_table_statement)
	//  statement.
	//  * `DROP_ROW_ACCESS_POLICY`: [`DROP [ALL] ROW ACCESS
	//  POLICY|POLICIES`](https://cloud.google.com/bigquery/docs/reference/standard-sql/data-definition-language#drop_row_access_policy_statement)
	//  statement.
	//  * `ALTER_TABLE`: [`ALTER
	//  TABLE`](https://cloud.google.com/bigquery/docs/reference/standard-sql/data-definition-language#alter_table_set_options_statement)
	//  statement.
	//  * `ALTER_VIEW`: [`ALTER
	//  VIEW`](https://cloud.google.com/bigquery/docs/reference/standard-sql/data-definition-language#alter_view_set_options_statement)
	//  statement.
	//  * `ALTER_MATERIALIZED_VIEW`: [`ALTER MATERIALIZED
	//  VIEW`](https://cloud.google.com/bigquery/docs/reference/standard-sql/data-definition-language#alter_materialized_view_set_options_statement)
	//  statement.
	//  * `ALTER_SCHEMA`: [`ALTER
	//  SCHEMA`](https://cloud.google.com/bigquery/docs/reference/standard-sql/data-definition-language#aalter_schema_set_options_statement)
	//  statement.
	//  * `SCRIPT`:
	//  [`SCRIPT`](https://cloud.google.com/bigquery/docs/reference/standard-sql/procedural-language).
	//  * `TRUNCATE_TABLE`: [`TRUNCATE
	//  TABLE`](https://cloud.google.com/bigquery/docs/reference/standard-sql/dml-syntax#truncate_table_statement)
	//  statement.
	//  * `CREATE_EXTERNAL_TABLE`: [`CREATE EXTERNAL
	//  TABLE`](https://cloud.google.com/bigquery/docs/reference/standard-sql/data-definition-language#create_external_table_statement)
	//  statement.
	//  * `EXPORT_DATA`: [`EXPORT
	//  DATA`](https://cloud.google.com/bigquery/docs/reference/standard-sql/other-statements#export_data_statement)
	//  statement.
	//  * `EXPORT_MODEL`: [`EXPORT
	//  MODEL`](https://cloud.google.com/bigquery-ml/docs/reference/standard-sql/bigqueryml-syntax-export-model)
	//  statement.
	//  * `LOAD_DATA`: [`LOAD
	//  DATA`](https://cloud.google.com/bigquery/docs/reference/standard-sql/other-statements#load_data_statement)
	//  statement.
	//  * `CALL`:
	//  [`CALL`](https://cloud.google.com/bigquery/docs/reference/standard-sql/procedural-language#call)
	//  statement.
	StatementType *string `json:"statementType,omitempty"`

	// Output only. The DDL operation performed, possibly
	//  dependent on the pre-existence of the DDL target.
	DdlOperationPerformed *string `json:"ddlOperationPerformed,omitempty"`

	// Output only. The DDL target table. Present only for
	//  CREATE/DROP TABLE/VIEW and DROP ALL ROW ACCESS POLICIES queries.
	DdlTargetTable *TableReference `json:"ddlTargetTable,omitempty"`

	// Output only. The table after rename. Present only for ALTER TABLE RENAME TO
	//  query.
	DdlDestinationTable *TableReference `json:"ddlDestinationTable,omitempty"`

	// Output only. The DDL target row access policy. Present only for
	//  CREATE/DROP ROW ACCESS POLICY queries.
	DdlTargetRowAccessPolicy *RowAccessPolicyReference `json:"ddlTargetRowAccessPolicy,omitempty"`

	// Output only. The number of row access policies affected by a DDL statement.
	//  Present only for DROP ALL ROW ACCESS POLICIES queries.
	DdlAffectedRowAccessPolicyCount *int64 `json:"ddlAffectedRowAccessPolicyCount,omitempty"`

	// Output only. [Beta] The DDL target routine. Present only for
	//  CREATE/DROP FUNCTION/PROCEDURE queries.
	DdlTargetRoutine *RoutineReference `json:"ddlTargetRoutine,omitempty"`

	// Output only. The DDL target dataset. Present only for CREATE/ALTER/DROP
	//  SCHEMA(dataset) queries.
	DdlTargetDataset *DatasetReference `json:"ddlTargetDataset,omitempty"`

	// Output only. Statistics of a BigQuery ML training job.
	MlStatistics *MlStatistics `json:"mlStatistics,omitempty"`

	// Output only. Stats for EXPORT DATA statement.
	ExportDataStatistics *ExportDataStatistics `json:"exportDataStatistics,omitempty"`

	// Output only. Job cost breakdown as bigquery internal cost and external
	//  service costs.
	ExternalServiceCosts []ExternalServiceCost `json:"externalServiceCosts,omitempty"`

	// Output only. BI Engine specific Statistics.
	BiEngineStatistics *BiEngineStatistics `json:"biEngineStatistics,omitempty"`

	// Output only. Statistics for a LOAD query.
	LoadQueryStatistics *LoadQueryStatistics `json:"loadQueryStatistics,omitempty"`

	// Output only. Referenced table for DCL statement.
	DclTargetTable *TableReference `json:"dclTargetTable,omitempty"`

	// Output only. Referenced view for DCL statement.
	DclTargetView *TableReference `json:"dclTargetView,omitempty"`

	// Output only. Referenced dataset for DCL statement.
	DclTargetDataset *DatasetReference `json:"dclTargetDataset,omitempty"`

	// Output only. Search query specific statistics.
	SearchStatistics *SearchStatistics `json:"searchStatistics,omitempty"`

	// Output only. Vector Search query specific statistics.
	VectorSearchStatistics *VectorSearchStatistics `json:"vectorSearchStatistics,omitempty"`

	// Output only. Performance insights.
	PerformanceInsights *PerformanceInsights `json:"performanceInsights,omitempty"`

	// Output only. Query optimization information for a QUERY job.
	QueryInfo *QueryInfo `json:"queryInfo,omitempty"`

	// Output only. Statistics of a Spark procedure job.
	SparkStatistics *SparkStatistics `json:"sparkStatistics,omitempty"`

	// Output only. Total bytes transferred for cross-cloud queries such as Cross
	//  Cloud Transfer and CREATE TABLE AS SELECT (CTAS).
	TransferredBytes *int64 `json:"transferredBytes,omitempty"`

	// Output only. Statistics of materialized views of a query job.
	MaterializedViewStatistics *MaterializedViewStatistics `json:"materializedViewStatistics,omitempty"`

	// Output only. Statistics of metadata cache usage in a query for BigLake
	//  tables.
	MetadataCacheStatistics *MetadataCacheStatistics `json:"metadataCacheStatistics,omitempty"`
}

// +kcc:proto=google.cloud.bigquery.v2.JobStatistics3
type JobStatistics3 struct {
	// Output only. Number of source files in a load job.
	InputFiles *int64 `json:"inputFiles,omitempty"`

	// Output only. Number of bytes of source data in a load job.
	InputFileBytes *int64 `json:"inputFileBytes,omitempty"`

	// Output only. Number of rows imported in a load job.
	//  Note that while an import job is in the running state, this
	//  value may change.
	OutputRows *int64 `json:"outputRows,omitempty"`

	// Output only. Size of the loaded data in bytes. Note
	//  that while a load job is in the running state, this value may change.
	OutputBytes *int64 `json:"outputBytes,omitempty"`

	// Output only. The number of bad records encountered. Note that if the job
	//  has failed because of more bad records encountered than the maximum
	//  allowed in the load job configuration, then this number can be less than
	//  the total number of bad records present in the input data.
	BadRecords *int64 `json:"badRecords,omitempty"`

	// Output only. Describes a timeline of job execution.
	Timeline []QueryTimelineSample `json:"timeline,omitempty"`
}

// +kcc:proto=google.cloud.bigquery.v2.JobStatistics4
type JobStatistics4 struct {
	// Output only. Number of files per destination URI or URI pattern
	//  specified in the extract configuration. These values will be in the same
	//  order as the URIs specified in the 'destinationUris' field.
	DestinationUriFileCounts []int64 `json:"destinationUriFileCounts,omitempty"`

	// Output only. Number of user bytes extracted into the result. This is the
	//  byte count as computed by BigQuery for billing purposes
	//  and doesn't have any relationship with the number of actual
	//  result bytes extracted in the desired format.
	InputBytes *int64 `json:"inputBytes,omitempty"`

	// Output only. Describes a timeline of job execution.
	Timeline []QueryTimelineSample `json:"timeline,omitempty"`
}

// +kcc:proto=google.cloud.bigquery.v2.JobStatus
type JobStatus struct {
	// Output only. Final error result of the job. If present, indicates that the
	//  job has completed and was unsuccessful.
	ErrorResult *ErrorProto `json:"errorResult,omitempty"`

	// Output only. The first errors encountered during the running of the job.
	//  The final message includes the number of errors that caused the process to
	//  stop. Errors here do not necessarily mean that the job has not completed or
	//  was unsuccessful.
	Errors []ErrorProto `json:"errors,omitempty"`

	// Output only. Running state of the job.  Valid states include 'PENDING',
	//  'RUNNING', and 'DONE'.
	State *string `json:"state,omitempty"`
}

// +kcc:proto=google.cloud.bigquery.v2.JoinRestrictionPolicy
type JoinRestrictionPolicy struct {
	// Optional. Specifies if a join is required or not on queries for the view.
	//  Default is JOIN_CONDITION_UNSPECIFIED.
	JoinCondition *string `json:"joinCondition,omitempty"`

	// Optional. The only columns that joins are allowed on.
	//  This field is must be specified for join_conditions JOIN_ANY and JOIN_ALL
	//  and it cannot be set for JOIN_BLOCKED.
	JoinAllowedColumns []string `json:"joinAllowedColumns,omitempty"`
}

// +kcc:proto=google.cloud.bigquery.v2.JsonOptions
type JsonOptions struct {
	// Optional. The character encoding of the data.
	//  The supported values are UTF-8, UTF-16BE, UTF-16LE, UTF-32BE,
	//  and UTF-32LE.  The default value is UTF-8.
	Encoding *string `json:"encoding,omitempty"`
}

// +kcc:proto=google.cloud.bigquery.v2.LinkedDatasetSource
type LinkedDatasetSource struct {
	// The source dataset reference contains project numbers and not project ids.
	SourceDataset *DatasetReference `json:"sourceDataset,omitempty"`
}

// +kcc:proto=google.cloud.bigquery.v2.ListFormatDataset
type ListFormatDataset struct {
	// The resource type.
	//  This property always returns the value "bigquery#dataset"
	Kind *string `json:"kind,omitempty"`

	// The fully-qualified, unique, opaque ID of the dataset.
	ID *string `json:"id,omitempty"`

	// The dataset reference.
	//  Use this property to access specific parts of the dataset's ID, such as
	//  project ID or dataset ID.
	DatasetReference *DatasetReference `json:"datasetReference,omitempty"`

	// The labels associated with this dataset.
	//  You can use these to organize and group your datasets.
	Labels map[string]string `json:"labels,omitempty"`

	// An alternate name for the dataset.  The friendly name is purely
	//  decorative in nature.
	FriendlyName *string `json:"friendlyName,omitempty"`

	// The geographic location where the dataset resides.
	Location *string `json:"location,omitempty"`
}

// +kcc:proto=google.cloud.bigquery.v2.ListFormatJob
type ListFormatJob struct {
	// Unique opaque ID of the job.
	ID *string `json:"id,omitempty"`

	// The resource type.
	Kind *string `json:"kind,omitempty"`

	// Unique opaque ID of the job.
	JobReference *JobReference `json:"jobReference,omitempty"`

	// Running state of the job. When the state is DONE, errorResult can be
	//  checked to determine whether the job succeeded or failed.
	State *string `json:"state,omitempty"`

	// A result object that will be present only if the job has failed.
	ErrorResult *ErrorProto `json:"errorResult,omitempty"`

	// Output only. Information about the job, including starting time and ending
	//  time of the job.
	Statistics *JobStatistics `json:"statistics,omitempty"`

	// Required. Describes the job configuration.
	Configuration *JobConfiguration `json:"configuration,omitempty"`

	// [Full-projection-only] Describes the status of this job.
	Status *JobStatus `json:"status,omitempty"`

	// [Full-projection-only] Email address of the user who ran the job.
	UserEmail *string `json:"userEmail,omitempty"`

	// [Full-projection-only] String representation of identity of requesting
	//  party. Populated for both first- and third-party identities. Only present
	//  for APIs that support third-party identities.
	PrincipalSubject *string `json:"principalSubject,omitempty"`
}

// +kcc:proto=google.cloud.bigquery.v2.ListFormatTable
type ListFormatTable struct {
	// The resource type.
	Kind *string `json:"kind,omitempty"`

	// An opaque ID of the table.
	ID *string `json:"id,omitempty"`

	// A reference uniquely identifying table.
	TableReference *TableReference `json:"tableReference,omitempty"`

	// The user-friendly name for this table.
	FriendlyName *string `json:"friendlyName,omitempty"`

	// The type of table.
	Type *string `json:"type,omitempty"`

	// The time-based partitioning for this table.
	TimePartitioning *TimePartitioning `json:"timePartitioning,omitempty"`

	// The range partitioning for this table.
	RangePartitioning *RangePartitioning `json:"rangePartitioning,omitempty"`

	// Clustering specification for this table, if configured.
	Clustering *Clustering `json:"clustering,omitempty"`

	// The labels associated with this table. You can use these to organize
	//  and group your tables.
	Labels map[string]string `json:"labels,omitempty"`

	// Additional details for a view.
	View *ListFormatView `json:"view,omitempty"`

	// Output only. The time when this table was created, in milliseconds since
	//  the epoch.
	CreationTime *int64 `json:"creationTime,omitempty"`

	// The time when this table expires, in milliseconds since the
	//  epoch. If not present, the table will persist indefinitely. Expired tables
	//  will be deleted and their storage reclaimed.
	ExpirationTime *int64 `json:"expirationTime,omitempty"`

	// Optional. If set to true, queries including this table must specify a
	//  partition filter. This filter is used for partition elimination.
	RequirePartitionFilter *bool `json:"requirePartitionFilter,omitempty"`
}

// +kcc:proto=google.cloud.bigquery.v2.ListFormatView
type ListFormatView struct {
	// True if view is defined in legacy SQL dialect,
	//  false if in GoogleSQL.
	UseLegacySql *bool `json:"useLegacySql,omitempty"`

	// Specifics the privacy policy for the view.
	PrivacyPolicy *PrivacyPolicy `json:"privacyPolicy,omitempty"`
}

// +kcc:proto=google.cloud.bigquery.v2.LoadQueryStatistics
type LoadQueryStatistics struct {
	// Output only. Number of source files in a LOAD query.
	InputFiles *int64 `json:"inputFiles,omitempty"`

	// Output only. Number of bytes of source data in a LOAD query.
	InputFileBytes *int64 `json:"inputFileBytes,omitempty"`

	// Output only. Number of rows imported in a LOAD query.
	//  Note that while a LOAD query is in the running state, this value may
	//  change.
	OutputRows *int64 `json:"outputRows,omitempty"`

	// Output only. Size of the loaded data in bytes. Note that while a LOAD query
	//  is in the running state, this value may change.
	OutputBytes *int64 `json:"outputBytes,omitempty"`

	// Output only. The number of bad records encountered while processing a LOAD
	//  query. Note that if the job has failed because of more bad records
	//  encountered than the maximum allowed in the load job configuration, then
	//  this number can be less than the total number of bad records present in the
	//  input data.
	BadRecords *int64 `json:"badRecords,omitempty"`
}

// +kcc:proto=google.cloud.bigquery.v2.MaterializedView
type MaterializedView struct {
	// The candidate materialized view.
	TableReference *TableReference `json:"tableReference,omitempty"`

	// Whether the materialized view is chosen for the query.
	//
	//  A materialized view can be chosen to rewrite multiple parts of the same
	//  query. If a materialized view is chosen to rewrite any part of the query,
	//  then this field is true, even if the materialized view was not chosen to
	//  rewrite others parts.
	Chosen *bool `json:"chosen,omitempty"`

	// If present, specifies a best-effort estimation of the bytes saved by using
	//  the materialized view rather than its base tables.
	EstimatedBytesSaved *int64 `json:"estimatedBytesSaved,omitempty"`

	// If present, specifies the reason why the materialized view was not chosen
	//  for the query.
	RejectedReason *string `json:"rejectedReason,omitempty"`
}

// +kcc:proto=google.cloud.bigquery.v2.MaterializedViewDefinition
type MaterializedViewDefinition struct {
	// Required. A query whose results are persisted.
	Query *string `json:"query,omitempty"`

	// Output only. The time when this materialized view was last refreshed, in
	//  milliseconds since the epoch.
	LastRefreshTime *int64 `json:"lastRefreshTime,omitempty"`

	// Optional. Enable automatic refresh of the materialized view when the base
	//  table is updated. The default value is "true".
	EnableRefresh *bool `json:"enableRefresh,omitempty"`

	// Optional. The maximum frequency at which this materialized view will be
	//  refreshed. The default value is "1800000" (30 minutes).
	RefreshIntervalMs *uint64 `json:"refreshIntervalMs,omitempty"`

	// Optional. This option declares the intention to construct a materialized
	//  view that isn't refreshed incrementally.
	AllowNonIncrementalDefinition *bool `json:"allowNonIncrementalDefinition,omitempty"`
}

// +kcc:proto=google.cloud.bigquery.v2.MaterializedViewStatistics
type MaterializedViewStatistics struct {
	// Materialized views considered for the query job. Only certain materialized
	//  views are used. For a detailed list, see the child message.
	//
	//  If many materialized views are considered, then the list might be
	//  incomplete.
	MaterializedView []MaterializedView `json:"materializedView,omitempty"`
}

// +kcc:proto=google.cloud.bigquery.v2.MaterializedViewStatus
type MaterializedViewStatus struct {
	// Output only. Refresh watermark of materialized view. The base tables' data
	//  were collected into the materialized view cache until this time.
	RefreshWatermark *string `json:"refreshWatermark,omitempty"`

	// Output only. Error result of the last automatic refresh. If present,
	//  indicates that the last automatic refresh was unsuccessful.
	LastRefreshStatus *ErrorProto `json:"lastRefreshStatus,omitempty"`
}

// +kcc:proto=google.cloud.bigquery.v2.MetadataCacheStatistics
type MetadataCacheStatistics struct {
	// Set for the Metadata caching eligible tables referenced in the query.
	TableMetadataCacheUsage []TableMetadataCacheUsage `json:"tableMetadataCacheUsage,omitempty"`
}

// +kcc:proto=google.cloud.bigquery.v2.MlStatistics
type MlStatistics struct {
	// Output only. Maximum number of iterations specified as max_iterations in
	//  the 'CREATE MODEL' query. The actual number of iterations may be less than
	//  this number due to early stop.
	MaxIterations *int64 `json:"maxIterations,omitempty"`

	// Results for all completed iterations.
	//  Empty for [hyperparameter tuning
	//  jobs](https://cloud.google.com/bigquery-ml/docs/reference/standard-sql/bigqueryml-syntax-hp-tuning-overview).
	IterationResults []Model_TrainingRun_IterationResult `json:"iterationResults,omitempty"`

	// Output only. The type of the model that is being trained.
	ModelType *string `json:"modelType,omitempty"`

	// Output only. Training type of the job.
	TrainingType *string `json:"trainingType,omitempty"`

	// Output only. Trials of a [hyperparameter tuning
	//  job](https://cloud.google.com/bigquery-ml/docs/reference/standard-sql/bigqueryml-syntax-hp-tuning-overview)
	//  sorted by trial_id.
	HparamTrials []Model_HparamTuningTrial `json:"hparamTrials,omitempty"`
}

// +kcc:proto=google.cloud.bigquery.v2.Model
type Model struct {
	// Output only. A hash of this resource.
	Etag *string `json:"etag,omitempty"`

	// Required. Unique identifier for this model.
	ModelReference *ModelReference `json:"modelReference,omitempty"`

	// Output only. The time when this model was created, in millisecs since the
	//  epoch.
	CreationTime *int64 `json:"creationTime,omitempty"`

	// Output only. The time when this model was last modified, in millisecs since
	//  the epoch.
	LastModifiedTime *int64 `json:"lastModifiedTime,omitempty"`

	// Optional. A user-friendly description of this model.
	Description *string `json:"description,omitempty"`

	// Optional. A descriptive name for this model.
	FriendlyName *string `json:"friendlyName,omitempty"`

	// The labels associated with this model. You can use these to organize
	//  and group your models. Label keys and values can be no longer
	//  than 63 characters, can only contain lowercase letters, numeric
	//  characters, underscores and dashes. International characters are allowed.
	//  Label values are optional. Label keys must start with a letter and each
	//  label in the list must have a different key.
	Labels map[string]string `json:"labels,omitempty"`

	// Optional. The time when this model expires, in milliseconds since the
	//  epoch. If not present, the model will persist indefinitely. Expired models
	//  will be deleted and their storage reclaimed.  The defaultTableExpirationMs
	//  property of the encapsulating dataset can be used to set a default
	//  expirationTime on newly created models.
	ExpirationTime *int64 `json:"expirationTime,omitempty"`

	// Output only. The geographic location where the model resides. This value
	//  is inherited from the dataset.
	Location *string `json:"location,omitempty"`

	// Custom encryption configuration (e.g., Cloud KMS keys). This shows the
	//  encryption configuration of the model data while stored in BigQuery
	//  storage. This field can be used with PatchModel to update encryption key
	//  for an already encrypted model.
	EncryptionConfiguration *EncryptionConfiguration `json:"encryptionConfiguration,omitempty"`

	// Output only. Type of the model resource.
	ModelType *string `json:"modelType,omitempty"`

	// Information for all training runs in increasing order of start_time.
	TrainingRuns []Model_TrainingRun `json:"trainingRuns,omitempty"`

	// Output only. Input feature columns for the model inference. If the model is
	//  trained with TRANSFORM clause, these are the input of the TRANSFORM clause.
	FeatureColumns []StandardSqlField `json:"featureColumns,omitempty"`

	// Output only. Label columns that were used to train this model.
	//  The output of the model will have a "predicted_" prefix to these columns.
	LabelColumns []StandardSqlField `json:"labelColumns,omitempty"`

	// Output only. This field will be populated if a TRANSFORM clause was used to
	//  train a model. TRANSFORM clause (if used) takes feature_columns as input
	//  and outputs transform_columns. transform_columns then are used to train the
	//  model.
	TransformColumns []TransformColumn `json:"transformColumns,omitempty"`

	// Output only. All hyperparameter search spaces in this model.
	HparamSearchSpaces *Model_HparamSearchSpaces `json:"hparamSearchSpaces,omitempty"`

	// Output only. The default trial_id to use in TVFs when the trial_id is not
	//  passed in. For single-objective [hyperparameter
	//  tuning](https://cloud.google.com/bigquery-ml/docs/reference/standard-sql/bigqueryml-syntax-hp-tuning-overview)
	//  models, this is the best trial ID. For multi-objective [hyperparameter
	//  tuning](https://cloud.google.com/bigquery-ml/docs/reference/standard-sql/bigqueryml-syntax-hp-tuning-overview)
	//  models, this is the smallest trial ID among all Pareto optimal trials.
	DefaultTrialID *int64 `json:"defaultTrialID,omitempty"`

	// Output only. Trials of a [hyperparameter
	//  tuning](https://cloud.google.com/bigquery-ml/docs/reference/standard-sql/bigqueryml-syntax-hp-tuning-overview)
	//  model sorted by trial_id.
	HparamTrials []Model_HparamTuningTrial `json:"hparamTrials,omitempty"`

	// Output only. For single-objective [hyperparameter
	//  tuning](https://cloud.google.com/bigquery-ml/docs/reference/standard-sql/bigqueryml-syntax-hp-tuning-overview)
	//  models, it only contains the best trial. For multi-objective
	//  [hyperparameter
	//  tuning](https://cloud.google.com/bigquery-ml/docs/reference/standard-sql/bigqueryml-syntax-hp-tuning-overview)
	//  models, it contains all Pareto optimal trials sorted by trial_id.
	OptimalTrialIds []int64 `json:"optimalTrialIds,omitempty"`

	// Output only. Remote model info
	RemoteModelInfo *RemoteModelInfo `json:"remoteModelInfo,omitempty"`
}

// +kcc:proto=google.cloud.bigquery.v2.Model.AggregateClassificationMetrics
type Model_AggregateClassificationMetrics struct {
	// Precision is the fraction of actual positive predictions that had
	//  positive actual labels. For multiclass this is a macro-averaged
	//  metric treating each class as a binary classifier.
	Precision *float64 `json:"precision,omitempty"`

	// Recall is the fraction of actual positive labels that were given a
	//  positive prediction. For multiclass this is a macro-averaged metric.
	Recall *float64 `json:"recall,omitempty"`

	// Accuracy is the fraction of predictions given the correct label. For
	//  multiclass this is a micro-averaged metric.
	Accuracy *float64 `json:"accuracy,omitempty"`

	// Threshold at which the metrics are computed. For binary
	//  classification models this is the positive class threshold.
	//  For multi-class classfication models this is the confidence
	//  threshold.
	Threshold *float64 `json:"threshold,omitempty"`

	// The F1 score is an average of recall and precision. For multiclass
	//  this is a macro-averaged metric.
	F1Score *float64 `json:"f1Score,omitempty"`

	// Logarithmic Loss. For multiclass this is a macro-averaged metric.
	LogLoss *float64 `json:"logLoss,omitempty"`

	// Area Under a ROC Curve. For multiclass this is a macro-averaged
	//  metric.
	RocAuc *float64 `json:"rocAuc,omitempty"`
}

// +kcc:proto=google.cloud.bigquery.v2.Model.ArimaFittingMetrics
type Model_ArimaFittingMetrics struct {
	// Log-likelihood.
	LogLikelihood *float64 `json:"logLikelihood,omitempty"`

	// AIC.
	Aic *float64 `json:"aic,omitempty"`

	// Variance.
	Variance *float64 `json:"variance,omitempty"`
}

// +kcc:proto=google.cloud.bigquery.v2.Model.ArimaForecastingMetrics
type Model_ArimaForecastingMetrics struct {
	// Repeated as there can be many metric sets (one for each model) in
	//  auto-arima and the large-scale case.
	ArimaSingleModelForecastingMetrics []Model_ArimaForecastingMetrics_ArimaSingleModelForecastingMetrics `json:"arimaSingleModelForecastingMetrics,omitempty"`
}

// +kcc:proto=google.cloud.bigquery.v2.Model.ArimaForecastingMetrics.ArimaSingleModelForecastingMetrics
type Model_ArimaForecastingMetrics_ArimaSingleModelForecastingMetrics struct {
	// Non-seasonal order.
	NonSeasonalOrder *Model_ArimaOrder `json:"nonSeasonalOrder,omitempty"`

	// Arima fitting metrics.
	ArimaFittingMetrics *Model_ArimaFittingMetrics `json:"arimaFittingMetrics,omitempty"`

	// Is arima model fitted with drift or not. It is always false when d
	//  is not 1.
	HasDrift *bool `json:"hasDrift,omitempty"`

	// The time_series_id value for this time series. It will be one of
	//  the unique values from the time_series_id_column specified during
	//  ARIMA model training. Only present when time_series_id_column
	//  training option was used.
	TimeSeriesID *string `json:"timeSeriesID,omitempty"`

	// The tuple of time_series_ids identifying this time series. It will
	//  be one of the unique tuples of values present in the
	//  time_series_id_columns specified during ARIMA model training. Only
	//  present when time_series_id_columns training option was used and
	//  the order of values here are same as the order of
	//  time_series_id_columns.
	TimeSeriesIds []string `json:"timeSeriesIds,omitempty"`

	// Seasonal periods. Repeated because multiple periods are supported
	//  for one time series.
	SeasonalPeriods []string `json:"seasonalPeriods,omitempty"`

	// If true, holiday_effect is a part of time series decomposition result.
	HasHolidayEffect *bool `json:"hasHolidayEffect,omitempty"`

	// If true, spikes_and_dips is a part of time series decomposition result.
	HasSpikesAndDips *bool `json:"hasSpikesAndDips,omitempty"`

	// If true, step_changes is a part of time series decomposition result.
	HasStepChanges *bool `json:"hasStepChanges,omitempty"`
}

// +kcc:proto=google.cloud.bigquery.v2.Model.ArimaOrder
type Model_ArimaOrder struct {
	// Order of the autoregressive part.
	P *int64 `json:"p,omitempty"`

	// Order of the differencing part.
	D *int64 `json:"d,omitempty"`

	// Order of the moving-average part.
	Q *int64 `json:"q,omitempty"`
}

// +kcc:proto=google.cloud.bigquery.v2.Model.BinaryClassificationMetrics
type Model_BinaryClassificationMetrics struct {
	// Aggregate classification metrics.
	AggregateClassificationMetrics *Model_AggregateClassificationMetrics `json:"aggregateClassificationMetrics,omitempty"`

	// Binary confusion matrix at multiple thresholds.
	BinaryConfusionMatrixList []Model_BinaryClassificationMetrics_BinaryConfusionMatrix `json:"binaryConfusionMatrixList,omitempty"`

	// Label representing the positive class.
	PositiveLabel *string `json:"positiveLabel,omitempty"`

	// Label representing the negative class.
	NegativeLabel *string `json:"negativeLabel,omitempty"`
}

// +kcc:proto=google.cloud.bigquery.v2.Model.BinaryClassificationMetrics.BinaryConfusionMatrix
type Model_BinaryClassificationMetrics_BinaryConfusionMatrix struct {
	// Threshold value used when computing each of the following metric.
	PositiveClassThreshold *float64 `json:"positiveClassThreshold,omitempty"`

	// Number of true samples predicted as true.
	TruePositives *int64 `json:"truePositives,omitempty"`

	// Number of false samples predicted as true.
	FalsePositives *int64 `json:"falsePositives,omitempty"`

	// Number of true samples predicted as false.
	TrueNegatives *int64 `json:"trueNegatives,omitempty"`

	// Number of false samples predicted as false.
	FalseNegatives *int64 `json:"falseNegatives,omitempty"`

	// The fraction of actual positive predictions that had positive actual
	//  labels.
	Precision *float64 `json:"precision,omitempty"`

	// The fraction of actual positive labels that were given a positive
	//  prediction.
	Recall *float64 `json:"recall,omitempty"`

	// The equally weighted average of recall and precision.
	F1Score *float64 `json:"f1Score,omitempty"`

	// The fraction of predictions given the correct label.
	Accuracy *float64 `json:"accuracy,omitempty"`
}

// +kcc:proto=google.cloud.bigquery.v2.Model.BoostedTreeOptionEnums
type Model_BoostedTreeOptionEnums struct {
}

// +kcc:proto=google.cloud.bigquery.v2.Model.CategoryEncodingMethod
type Model_CategoryEncodingMethod struct {
}

// +kcc:proto=google.cloud.bigquery.v2.Model.ClusteringMetrics
type Model_ClusteringMetrics struct {
	// Davies-Bouldin index.
	DaviesBouldinIndex *float64 `json:"daviesBouldinIndex,omitempty"`

	// Mean of squared distances between each sample to its cluster centroid.
	MeanSquaredDistance *float64 `json:"meanSquaredDistance,omitempty"`

	// Information for all clusters.
	Clusters []Model_ClusteringMetrics_Cluster `json:"clusters,omitempty"`
}

// +kcc:proto=google.cloud.bigquery.v2.Model.ClusteringMetrics.Cluster
type Model_ClusteringMetrics_Cluster struct {
	// Centroid id.
	CentroidID *int64 `json:"centroidID,omitempty"`

	// Values of highly variant features for this cluster.
	FeatureValues []Model_ClusteringMetrics_Cluster_FeatureValue `json:"featureValues,omitempty"`

	// Count of training data rows that were assigned to this cluster.
	Count *int64 `json:"count,omitempty"`
}

// +kcc:proto=google.cloud.bigquery.v2.Model.ClusteringMetrics.Cluster.FeatureValue
type Model_ClusteringMetrics_Cluster_FeatureValue struct {
	// The feature column name.
	FeatureColumn *string `json:"featureColumn,omitempty"`

	// The numerical feature value. This is the centroid value for this
	//  feature.
	NumericalValue *float64 `json:"numericalValue,omitempty"`

	// The categorical feature value.
	CategoricalValue *Model_ClusteringMetrics_Cluster_FeatureValue_CategoricalValue `json:"categoricalValue,omitempty"`
}

// +kcc:proto=google.cloud.bigquery.v2.Model.ClusteringMetrics.Cluster.FeatureValue.CategoricalValue
type Model_ClusteringMetrics_Cluster_FeatureValue_CategoricalValue struct {
	// Counts of all categories for the categorical feature. If there are
	//  more than ten categories, we return top ten (by count) and return
	//  one more CategoryCount with category "_OTHER_" and count as
	//  aggregate counts of remaining categories.
	CategoryCounts []Model_ClusteringMetrics_Cluster_FeatureValue_CategoricalValue_CategoryCount `json:"categoryCounts,omitempty"`
}

// +kcc:proto=google.cloud.bigquery.v2.Model.ClusteringMetrics.Cluster.FeatureValue.CategoricalValue.CategoryCount
type Model_ClusteringMetrics_Cluster_FeatureValue_CategoricalValue_CategoryCount struct {
	// The name of category.
	Category *string `json:"category,omitempty"`

	// The count of training samples matching the category within the
	//  cluster.
	Count *int64 `json:"count,omitempty"`
}

// +kcc:proto=google.cloud.bigquery.v2.Model.DataSplitResult
type Model_DataSplitResult struct {
	// Table reference of the training data after split.
	TrainingTable *TableReference `json:"trainingTable,omitempty"`

	// Table reference of the evaluation data after split.
	EvaluationTable *TableReference `json:"evaluationTable,omitempty"`

	// Table reference of the test data after split.
	TestTable *TableReference `json:"testTable,omitempty"`
}

// +kcc:proto=google.cloud.bigquery.v2.Model.DimensionalityReductionMetrics
type Model_DimensionalityReductionMetrics struct {
	// Total percentage of variance explained by the selected principal
	//  components.
	TotalExplainedVarianceRatio *float64 `json:"totalExplainedVarianceRatio,omitempty"`
}

// +kcc:proto=google.cloud.bigquery.v2.Model.float64HparamSearchSpace
type Model_float64HparamSearchSpace struct {
	// Range of the float64 hyperparameter.
	Range *Model_float64HparamSearchSpace_float64Range `json:"range,omitempty"`

	// Candidates of the float64 hyperparameter.
	Candidates *Model_float64HparamSearchSpace_float64Candidates `json:"candidates,omitempty"`
}

// +kcc:proto=google.cloud.bigquery.v2.Model.float64HparamSearchSpace.float64Candidates
type Model_float64HparamSearchSpace_float64Candidates struct {
	// Candidates for the float64 parameter in increasing order.
	Candidates []float64 `json:"candidates,omitempty"`
}

// +kcc:proto=google.cloud.bigquery.v2.Model.float64HparamSearchSpace.float64Range
type Model_float64HparamSearchSpace_float64Range struct {
	// Min value of the float64 parameter.
	Min *float64 `json:"min,omitempty"`

	// Max value of the float64 parameter.
	Max *float64 `json:"max,omitempty"`
}

// +kcc:proto=google.cloud.bigquery.v2.Model.EvaluationMetrics
type Model_EvaluationMetrics struct {
	// Populated for regression models and explicit feedback type matrix
	//  factorization models.
	RegressionMetrics *Model_RegressionMetrics `json:"regressionMetrics,omitempty"`

	// Populated for binary classification/classifier models.
	BinaryClassificationMetrics *Model_BinaryClassificationMetrics `json:"binaryClassificationMetrics,omitempty"`

	// Populated for multi-class classification/classifier models.
	MultiClassClassificationMetrics *Model_MultiClassClassificationMetrics `json:"multiClassClassificationMetrics,omitempty"`

	// Populated for clustering models.
	ClusteringMetrics *Model_ClusteringMetrics `json:"clusteringMetrics,omitempty"`

	// Populated for implicit feedback type matrix factorization models.
	RankingMetrics *Model_RankingMetrics `json:"rankingMetrics,omitempty"`

	// Populated for ARIMA models.
	ArimaForecastingMetrics *Model_ArimaForecastingMetrics `json:"arimaForecastingMetrics,omitempty"`

	// Evaluation metrics when the model is a dimensionality reduction model,
	//  which currently includes PCA.
	DimensionalityReductionMetrics *Model_DimensionalityReductionMetrics `json:"dimensionalityReductionMetrics,omitempty"`
}

// +kcc:proto=google.cloud.bigquery.v2.Model.GlobalExplanation
type Model_GlobalExplanation struct {
	// A list of the top global explanations. Sorted by absolute value of
	//  attribution in descending order.
	Explanations []Model_GlobalExplanation_Explanation `json:"explanations,omitempty"`

	// Class label for this set of global explanations. Will be empty/null for
	//  binary logistic and linear regression models. Sorted alphabetically in
	//  descending order.
	ClassLabel *string `json:"classLabel,omitempty"`
}

// +kcc:proto=google.cloud.bigquery.v2.Model.GlobalExplanation.Explanation
type Model_GlobalExplanation_Explanation struct {
	// The full feature name. For non-numerical features, will be formatted
	//  like `<column_name>.<encoded_feature_name>`. Overall size of feature
	//  name will always be truncated to first 120 characters.
	FeatureName *string `json:"featureName,omitempty"`

	// Attribution of feature.
	Attribution *float64 `json:"attribution,omitempty"`
}

// +kcc:proto=google.cloud.bigquery.v2.Model.HparamSearchSpaces
type Model_HparamSearchSpaces struct {
	// Learning rate of training jobs.
	LearnRate *Model_float64HparamSearchSpace `json:"learnRate,omitempty"`

	// L1 regularization coefficient.
	L1Reg *Model_float64HparamSearchSpace `json:"l1Reg,omitempty"`

	// L2 regularization coefficient.
	L2Reg *Model_float64HparamSearchSpace `json:"l2Reg,omitempty"`

	// Number of clusters for k-means.
	NumClusters *Model_IntHparamSearchSpace `json:"numClusters,omitempty"`

	// Number of latent factors to train on.
	NumFactors *Model_IntHparamSearchSpace `json:"numFactors,omitempty"`

	// Hidden units for neural network models.
	HiddenUnits *Model_IntArrayHparamSearchSpace `json:"hiddenUnits,omitempty"`

	// Mini batch sample size.
	BatchSize *Model_IntHparamSearchSpace `json:"batchSize,omitempty"`

	// Dropout probability for dnn model training and boosted tree models
	//  using dart booster.
	Dropout *Model_float64HparamSearchSpace `json:"dropout,omitempty"`

	// Maximum depth of a tree for boosted tree models.
	MaxTreeDepth *Model_IntHparamSearchSpace `json:"maxTreeDepth,omitempty"`

	// Subsample the training data to grow tree to prevent overfitting for
	//  boosted tree models.
	Subsample *Model_float64HparamSearchSpace `json:"subsample,omitempty"`

	// Minimum split loss for boosted tree models.
	MinSplitLoss *Model_float64HparamSearchSpace `json:"minSplitLoss,omitempty"`

	// Hyperparameter for matrix factoration when implicit feedback type is
	//  specified.
	WalsAlpha *Model_float64HparamSearchSpace `json:"walsAlpha,omitempty"`

	// Booster type for boosted tree models.
	BoosterType *Model_StringHparamSearchSpace `json:"boosterType,omitempty"`

	// Number of parallel trees for boosted tree models.
	NumParallelTree *Model_IntHparamSearchSpace `json:"numParallelTree,omitempty"`

	// Dart normalization type for boosted tree models.
	DartNormalizeType *Model_StringHparamSearchSpace `json:"dartNormalizeType,omitempty"`

	// Tree construction algorithm for boosted tree models.
	TreeMethod *Model_StringHparamSearchSpace `json:"treeMethod,omitempty"`

	// Minimum sum of instance weight needed in a child for boosted tree models.
	MinTreeChildWeight *Model_IntHparamSearchSpace `json:"minTreeChildWeight,omitempty"`

	// Subsample ratio of columns when constructing each tree for boosted tree
	//  models.
	ColsampleBytree *Model_float64HparamSearchSpace `json:"colsampleBytree,omitempty"`

	// Subsample ratio of columns for each level for boosted tree models.
	ColsampleBylevel *Model_float64HparamSearchSpace `json:"colsampleBylevel,omitempty"`

	// Subsample ratio of columns for each node(split) for boosted tree models.
	ColsampleBynode *Model_float64HparamSearchSpace `json:"colsampleBynode,omitempty"`

	// Activation functions of neural network models.
	ActivationFn *Model_StringHparamSearchSpace `json:"activationFn,omitempty"`

	// Optimizer of TF models.
	Optimizer *Model_StringHparamSearchSpace `json:"optimizer,omitempty"`
}

// +kcc:proto=google.cloud.bigquery.v2.Model.HparamTuningEnums
type Model_HparamTuningEnums struct {
}

// +kcc:proto=google.cloud.bigquery.v2.Model.HparamTuningTrial
type Model_HparamTuningTrial struct {
	// 1-based index of the trial.
	TrialID *int64 `json:"trialID,omitempty"`

	// Starting time of the trial.
	StartTimeMs *int64 `json:"startTimeMs,omitempty"`

	// Ending time of the trial.
	EndTimeMs *int64 `json:"endTimeMs,omitempty"`

	// The hyperprameters selected for this trial.
	Hparams *Model_TrainingRun_TrainingOptions `json:"hparams,omitempty"`

	// Evaluation metrics of this trial calculated on the test data.
	//  Empty in Job API.
	EvaluationMetrics *Model_EvaluationMetrics `json:"evaluationMetrics,omitempty"`

	// The status of the trial.
	Status *string `json:"status,omitempty"`

	// Error message for FAILED and INFEASIBLE trial.
	ErrorMessage *string `json:"errorMessage,omitempty"`

	// Loss computed on the training data at the end of trial.
	TrainingLoss *float64 `json:"trainingLoss,omitempty"`

	// Loss computed on the eval data at the end of trial.
	EvalLoss *float64 `json:"evalLoss,omitempty"`

	// Hyperparameter tuning evaluation metrics of this trial calculated on the
	//  eval data. Unlike evaluation_metrics, only the fields corresponding to
	//  the hparam_tuning_objectives are set.
	HparamTuningEvaluationMetrics *Model_EvaluationMetrics `json:"hparamTuningEvaluationMetrics,omitempty"`
}

// +kcc:proto=google.cloud.bigquery.v2.Model.IntArrayHparamSearchSpace
type Model_IntArrayHparamSearchSpace struct {
	// Candidates for the int array parameter.
	Candidates []Model_IntArrayHparamSearchSpace_IntArray `json:"candidates,omitempty"`
}

// +kcc:proto=google.cloud.bigquery.v2.Model.IntArrayHparamSearchSpace.IntArray
type Model_IntArrayHparamSearchSpace_IntArray struct {
	// Elements in the int array.
	Elements []int64 `json:"elements,omitempty"`
}

// +kcc:proto=google.cloud.bigquery.v2.Model.IntHparamSearchSpace
type Model_IntHparamSearchSpace struct {
	// Range of the int hyperparameter.
	Range *Model_IntHparamSearchSpace_IntRange `json:"range,omitempty"`

	// Candidates of the int hyperparameter.
	Candidates *Model_IntHparamSearchSpace_IntCandidates `json:"candidates,omitempty"`
}

// +kcc:proto=google.cloud.bigquery.v2.Model.IntHparamSearchSpace.IntCandidates
type Model_IntHparamSearchSpace_IntCandidates struct {
	// Candidates for the int parameter in increasing order.
	Candidates []int64 `json:"candidates,omitempty"`
}

// +kcc:proto=google.cloud.bigquery.v2.Model.IntHparamSearchSpace.IntRange
type Model_IntHparamSearchSpace_IntRange struct {
	// Min value of the int parameter.
	Min *int64 `json:"min,omitempty"`

	// Max value of the int parameter.
	Max *int64 `json:"max,omitempty"`
}

// +kcc:proto=google.cloud.bigquery.v2.Model.KmeansEnums
type Model_KmeansEnums struct {
}

// +kcc:proto=google.cloud.bigquery.v2.Model.ModelRegistryOptionEnums
type Model_ModelRegistryOptionEnums struct {
}

// +kcc:proto=google.cloud.bigquery.v2.Model.MultiClassClassificationMetrics
type Model_MultiClassClassificationMetrics struct {
	// Aggregate classification metrics.
	AggregateClassificationMetrics *Model_AggregateClassificationMetrics `json:"aggregateClassificationMetrics,omitempty"`

	// Confusion matrix at different thresholds.
	ConfusionMatrixList []Model_MultiClassClassificationMetrics_ConfusionMatrix `json:"confusionMatrixList,omitempty"`
}

// +kcc:proto=google.cloud.bigquery.v2.Model.MultiClassClassificationMetrics.ConfusionMatrix
type Model_MultiClassClassificationMetrics_ConfusionMatrix struct {
	// Confidence threshold used when computing the entries of the
	//  confusion matrix.
	ConfidenceThreshold *float64 `json:"confidenceThreshold,omitempty"`

	// One row per actual label.
	Rows []Model_MultiClassClassificationMetrics_ConfusionMatrix_Row `json:"rows,omitempty"`
}

// +kcc:proto=google.cloud.bigquery.v2.Model.MultiClassClassificationMetrics.ConfusionMatrix.Entry
type Model_MultiClassClassificationMetrics_ConfusionMatrix_Entry struct {
	// The predicted label. For confidence_threshold > 0, we will
	//  also add an entry indicating the number of items under the
	//  confidence threshold.
	PredictedLabel *string `json:"predictedLabel,omitempty"`

	// Number of items being predicted as this label.
	ItemCount *int64 `json:"itemCount,omitempty"`
}

// +kcc:proto=google.cloud.bigquery.v2.Model.MultiClassClassificationMetrics.ConfusionMatrix.Row
type Model_MultiClassClassificationMetrics_ConfusionMatrix_Row struct {
	// The original label of this row.
	ActualLabel *string `json:"actualLabel,omitempty"`

	// Info describing predicted label distribution.
	Entries []Model_MultiClassClassificationMetrics_ConfusionMatrix_Entry `json:"entries,omitempty"`
}

// +kcc:proto=google.cloud.bigquery.v2.Model.PcaSolverOptionEnums
type Model_PcaSolverOptionEnums struct {
}

// +kcc:proto=google.cloud.bigquery.v2.Model.RankingMetrics
type Model_RankingMetrics struct {
	// Calculates a precision per user for all the items by ranking them and
	//  then averages all the precisions across all the users.
	MeanAveragePrecision *float64 `json:"meanAveragePrecision,omitempty"`

	// Similar to the mean squared error computed in regression and explicit
	//  recommendation models except instead of computing the rating directly,
	//  the output from evaluate is computed against a preference which is 1 or 0
	//  depending on if the rating exists or not.
	MeanSquaredError *float64 `json:"meanSquaredError,omitempty"`

	// A metric to determine the goodness of a ranking calculated from the
	//  predicted confidence by comparing it to an ideal rank measured by the
	//  original ratings.
	NormalizedDiscountedCumulativeGain *float64 `json:"normalizedDiscountedCumulativeGain,omitempty"`

	// Determines the goodness of a ranking by computing the percentile rank
	//  from the predicted confidence and dividing it by the original rank.
	AverageRank *float64 `json:"averageRank,omitempty"`
}

// +kcc:proto=google.cloud.bigquery.v2.Model.RegressionMetrics
type Model_RegressionMetrics struct {
	// Mean absolute error.
	MeanAbsoluteError *float64 `json:"meanAbsoluteError,omitempty"`

	// Mean squared error.
	MeanSquaredError *float64 `json:"meanSquaredError,omitempty"`

	// Mean squared log error.
	MeanSquaredLogError *float64 `json:"meanSquaredLogError,omitempty"`

	// Median absolute error.
	MedianAbsoluteError *float64 `json:"medianAbsoluteError,omitempty"`

	// R^2 score. This corresponds to r2_score in ML.EVALUATE.
	RSquared *float64 `json:"rSquared,omitempty"`
}

// +kcc:proto=google.cloud.bigquery.v2.Model.SeasonalPeriod
type Model_SeasonalPeriod struct {
}

// +kcc:proto=google.cloud.bigquery.v2.Model.StringHparamSearchSpace
type Model_StringHparamSearchSpace struct {
	// Canididates for the string or enum parameter in lower case.
	Candidates []string `json:"candidates,omitempty"`
}

// +kcc:proto=google.cloud.bigquery.v2.Model.TrainingRun
type Model_TrainingRun struct {
	// Output only. Options that were used for this training run, includes
	//  user specified and default options that were used.
	TrainingOptions *Model_TrainingRun_TrainingOptions `json:"trainingOptions,omitempty"`

	// Output only. The start time of this training run.
	StartTime *string `json:"startTime,omitempty"`

	// Output only. Output of each iteration run, results.size() <=
	//  max_iterations.
	Results []Model_TrainingRun_IterationResult `json:"results,omitempty"`

	// Output only. The evaluation metrics over training/eval data that were
	//  computed at the end of training.
	EvaluationMetrics *Model_EvaluationMetrics `json:"evaluationMetrics,omitempty"`

	// Output only. Data split result of the training run. Only set when the
	//  input data is actually split.
	DataSplitResult *Model_DataSplitResult `json:"dataSplitResult,omitempty"`

	// Output only. Global explanation contains the explanation of top features
	//  on the model level. Applies to both regression and classification models.
	ModelLevelGlobalExplanation *Model_GlobalExplanation `json:"modelLevelGlobalExplanation,omitempty"`

	// Output only. Global explanation contains the explanation of top features
	//  on the class level. Applies to classification models only.
	ClassLevelGlobalExplanations []Model_GlobalExplanation `json:"classLevelGlobalExplanations,omitempty"`

	// The model id in the [Vertex AI Model
	//  Registry](https://cloud.google.com/vertex-ai/docs/model-registry/introduction)
	//  for this training run.
	VertexAiModelID *string `json:"vertexAiModelID,omitempty"`

	// Output only. The model version in the [Vertex AI Model
	//  Registry](https://cloud.google.com/vertex-ai/docs/model-registry/introduction)
	//  for this training run.
	VertexAiModelVersion *string `json:"vertexAiModelVersion,omitempty"`
}

// +kcc:proto=google.cloud.bigquery.v2.Model.TrainingRun.IterationResult
type Model_TrainingRun_IterationResult struct {
	// Index of the iteration, 0 based.
	Index *int32 `json:"index,omitempty"`

	// Time taken to run the iteration in milliseconds.
	DurationMs *int64 `json:"durationMs,omitempty"`

	// Loss computed on the training data at the end of iteration.
	TrainingLoss *float64 `json:"trainingLoss,omitempty"`

	// Loss computed on the eval data at the end of iteration.
	EvalLoss *float64 `json:"evalLoss,omitempty"`

	// Learn rate used for this iteration.
	LearnRate *float64 `json:"learnRate,omitempty"`

	// Information about top clusters for clustering models.
	ClusterInfos []Model_TrainingRun_IterationResult_ClusterInfo `json:"clusterInfos,omitempty"`

	// Arima result.
	ArimaResult *Model_TrainingRun_IterationResult_ArimaResult `json:"arimaResult,omitempty"`

	// The information of the principal components.
	PrincipalComponentInfos []Model_TrainingRun_IterationResult_PrincipalComponentInfo `json:"principalComponentInfos,omitempty"`
}

// +kcc:proto=google.cloud.bigquery.v2.Model.TrainingRun.IterationResult.ArimaResult
type Model_TrainingRun_IterationResult_ArimaResult struct {
	// This message is repeated because there are multiple arima models
	//  fitted in auto-arima. For non-auto-arima model, its size is one.
	ArimaModelInfo []Model_TrainingRun_IterationResult_ArimaResult_ArimaModelInfo `json:"arimaModelInfo,omitempty"`

	// Seasonal periods. Repeated because multiple periods are supported for
	//  one time series.
	SeasonalPeriods []string `json:"seasonalPeriods,omitempty"`
}

// +kcc:proto=google.cloud.bigquery.v2.Model.TrainingRun.IterationResult.ArimaResult.ArimaCoefficients
type Model_TrainingRun_IterationResult_ArimaResult_ArimaCoefficients struct {
	// Auto-regressive coefficients, an array of float64.
	AutoRegressiveCoefficients []float64 `json:"autoRegressiveCoefficients,omitempty"`

	// Moving-average coefficients, an array of float64.
	MovingAverageCoefficients []float64 `json:"movingAverageCoefficients,omitempty"`

	// Intercept coefficient, just a float64 not an array.
	InterceptCoefficient *float64 `json:"interceptCoefficient,omitempty"`
}

// +kcc:proto=google.cloud.bigquery.v2.Model.TrainingRun.IterationResult.ArimaResult.ArimaModelInfo
type Model_TrainingRun_IterationResult_ArimaResult_ArimaModelInfo struct {
	// Non-seasonal order.
	NonSeasonalOrder *Model_ArimaOrder `json:"nonSeasonalOrder,omitempty"`

	// Arima coefficients.
	ArimaCoefficients *Model_TrainingRun_IterationResult_ArimaResult_ArimaCoefficients `json:"arimaCoefficients,omitempty"`

	// Arima fitting metrics.
	ArimaFittingMetrics *Model_ArimaFittingMetrics `json:"arimaFittingMetrics,omitempty"`

	// Whether Arima model fitted with drift or not. It is always false
	//  when d is not 1.
	HasDrift *bool `json:"hasDrift,omitempty"`

	// The time_series_id value for this time series. It will be one of
	//  the unique values from the time_series_id_column specified during
	//  ARIMA model training. Only present when time_series_id_column
	//  training option was used.
	TimeSeriesID *string `json:"timeSeriesID,omitempty"`

	// The tuple of time_series_ids identifying this time series. It will
	//  be one of the unique tuples of values present in the
	//  time_series_id_columns specified during ARIMA model training. Only
	//  present when time_series_id_columns training option was used and
	//  the order of values here are same as the order of
	//  time_series_id_columns.
	TimeSeriesIds []string `json:"timeSeriesIds,omitempty"`

	// Seasonal periods. Repeated because multiple periods are supported
	//  for one time series.
	SeasonalPeriods []string `json:"seasonalPeriods,omitempty"`

	// If true, holiday_effect is a part of time series decomposition
	//  result.
	HasHolidayEffect *bool `json:"hasHolidayEffect,omitempty"`

	// If true, spikes_and_dips is a part of time series decomposition
	//  result.
	HasSpikesAndDips *bool `json:"hasSpikesAndDips,omitempty"`

	// If true, step_changes is a part of time series decomposition
	//  result.
	HasStepChanges *bool `json:"hasStepChanges,omitempty"`
}

// +kcc:proto=google.cloud.bigquery.v2.Model.TrainingRun.IterationResult.ClusterInfo
type Model_TrainingRun_IterationResult_ClusterInfo struct {
	// Centroid id.
	CentroidID *int64 `json:"centroidID,omitempty"`

	// Cluster radius, the average distance from centroid
	//  to each point assigned to the cluster.
	ClusterRadius *float64 `json:"clusterRadius,omitempty"`

	// Cluster size, the total number of points assigned to the cluster.
	ClusterSize *int64 `json:"clusterSize,omitempty"`
}

// +kcc:proto=google.cloud.bigquery.v2.Model.TrainingRun.IterationResult.PrincipalComponentInfo
type Model_TrainingRun_IterationResult_PrincipalComponentInfo struct {
	// Id of the principal component.
	PrincipalComponentID *int64 `json:"principalComponentID,omitempty"`

	// Explained variance by this principal component, which is simply the
	//  eigenvalue.
	ExplainedVariance *float64 `json:"explainedVariance,omitempty"`

	// Explained_variance over the total explained variance.
	ExplainedVarianceRatio *float64 `json:"explainedVarianceRatio,omitempty"`

	// The explained_variance is pre-ordered in the descending order to
	//  compute the cumulative explained variance ratio.
	CumulativeExplainedVarianceRatio *float64 `json:"cumulativeExplainedVarianceRatio,omitempty"`
}

// +kcc:proto=google.cloud.bigquery.v2.Model.TrainingRun.TrainingOptions
type Model_TrainingRun_TrainingOptions struct {
	// The maximum number of iterations in training. Used only for iterative
	//  training algorithms.
	MaxIterations *int64 `json:"maxIterations,omitempty"`

	// Type of loss function used during training run.
	LossType *string `json:"lossType,omitempty"`

	// Learning rate in training. Used only for iterative training algorithms.
	LearnRate *float64 `json:"learnRate,omitempty"`

	// L1 regularization coefficient.
	L1Regularization *float64 `json:"l1Regularization,omitempty"`

	// L2 regularization coefficient.
	L2Regularization *float64 `json:"l2Regularization,omitempty"`

	// When early_stop is true, stops training when accuracy improvement is
	//  less than 'min_relative_progress'. Used only for iterative training
	//  algorithms.
	MinRelativeProgress *float64 `json:"minRelativeProgress,omitempty"`

	// Whether to train a model from the last checkpoint.
	WarmStart *bool `json:"warmStart,omitempty"`

	// Whether to stop early when the loss doesn't improve significantly
	//  any more (compared to min_relative_progress). Used only for iterative
	//  training algorithms.
	EarlyStop *bool `json:"earlyStop,omitempty"`

	// Name of input label columns in training data.
	InputLabelColumns []string `json:"inputLabelColumns,omitempty"`

	// The data split type for training and evaluation, e.g. RANDOM.
	DataSplitMethod *string `json:"dataSplitMethod,omitempty"`

	// The fraction of evaluation data over the whole input data. The rest
	//  of data will be used as training data. The format should be float64.
	//  Accurate to two decimal places.
	//  Default value is 0.2.
	DataSplitEvalFraction *float64 `json:"dataSplitEvalFraction,omitempty"`

	// The column to split data with. This column won't be used as a
	//  feature.
	//  1. When data_split_method is CUSTOM, the corresponding column should
	//  be boolean. The rows with true value tag are eval data, and the false
	//  are training data.
	//  2. When data_split_method is SEQ, the first DATA_SPLIT_EVAL_FRACTION
	//  rows (from smallest to largest) in the corresponding column are used
	//  as training data, and the rest are eval data. It respects the order
	//  in Orderable data types:
	//  https://cloud.google.com/bigquery/docs/reference/standard-sql/data-types#data-type-properties
	DataSplitColumn *string `json:"dataSplitColumn,omitempty"`

	// The strategy to determine learn rate for the current iteration.
	LearnRateStrategy *string `json:"learnRateStrategy,omitempty"`

	// Specifies the initial learning rate for the line search learn rate
	//  strategy.
	InitialLearnRate *float64 `json:"initialLearnRate,omitempty"`

	// TODO: map type string float64 for label_class_weights

	// User column specified for matrix factorization models.
	UserColumn *string `json:"userColumn,omitempty"`

	// Item column specified for matrix factorization models.
	ItemColumn *string `json:"itemColumn,omitempty"`

	// Distance type for clustering models.
	DistanceType *string `json:"distanceType,omitempty"`

	// Number of clusters for clustering models.
	NumClusters *int64 `json:"numClusters,omitempty"`

	// Google Cloud Storage URI from which the model was imported. Only
	//  applicable for imported models.
	ModelUri *string `json:"modelUri,omitempty"`

	// Optimization strategy for training linear regression models.
	OptimizationStrategy *string `json:"optimizationStrategy,omitempty"`

	// Hidden units for dnn models.
	HiddenUnits []int64 `json:"hiddenUnits,omitempty"`

	// Batch size for dnn models.
	BatchSize *int64 `json:"batchSize,omitempty"`

	// Dropout probability for dnn models.
	Dropout *float64 `json:"dropout,omitempty"`

	// Maximum depth of a tree for boosted tree models.
	MaxTreeDepth *int64 `json:"maxTreeDepth,omitempty"`

	// Subsample fraction of the training data to grow tree to prevent
	//  overfitting for boosted tree models.
	Subsample *float64 `json:"subsample,omitempty"`

	// Minimum split loss for boosted tree models.
	MinSplitLoss *float64 `json:"minSplitLoss,omitempty"`

	// Booster type for boosted tree models.
	BoosterType *string `json:"boosterType,omitempty"`

	// Number of parallel trees constructed during each iteration for boosted
	//  tree models.
	NumParallelTree *int64 `json:"numParallelTree,omitempty"`

	// Type of normalization algorithm for boosted tree models using
	//  dart booster.
	DartNormalizeType *string `json:"dartNormalizeType,omitempty"`

	// Tree construction algorithm for boosted tree models.
	TreeMethod *string `json:"treeMethod,omitempty"`

	// Minimum sum of instance weight needed in a child for boosted tree
	//  models.
	MinTreeChildWeight *int64 `json:"minTreeChildWeight,omitempty"`

	// Subsample ratio of columns when constructing each tree for boosted tree
	//  models.
	ColsampleBytree *float64 `json:"colsampleBytree,omitempty"`

	// Subsample ratio of columns for each level for boosted tree models.
	ColsampleBylevel *float64 `json:"colsampleBylevel,omitempty"`

	// Subsample ratio of columns for each node(split) for boosted tree
	//  models.
	ColsampleBynode *float64 `json:"colsampleBynode,omitempty"`

	// Num factors specified for matrix factorization models.
	NumFactors *int64 `json:"numFactors,omitempty"`

	// Feedback type that specifies which algorithm to run for matrix
	//  factorization.
	FeedbackType *string `json:"feedbackType,omitempty"`

	// Hyperparameter for matrix factoration when implicit feedback type is
	//  specified.
	WalsAlpha *float64 `json:"walsAlpha,omitempty"`

	// The method used to initialize the centroids for kmeans algorithm.
	KmeansInitializationMethod *string `json:"kmeansInitializationMethod,omitempty"`

	// The column used to provide the initial centroids for kmeans algorithm
	//  when kmeans_initialization_method is CUSTOM.
	KmeansInitializationColumn *string `json:"kmeansInitializationColumn,omitempty"`

	// Column to be designated as time series timestamp for ARIMA model.
	TimeSeriesTimestampColumn *string `json:"timeSeriesTimestampColumn,omitempty"`

	// Column to be designated as time series data for ARIMA model.
	TimeSeriesDataColumn *string `json:"timeSeriesDataColumn,omitempty"`

	// Whether to enable auto ARIMA or not.
	AutoArima *bool `json:"autoArima,omitempty"`

	// A specification of the non-seasonal part of the ARIMA model: the three
	//  components (p, d, q) are the AR order, the degree of differencing, and
	//  the MA order.
	NonSeasonalOrder *Model_ArimaOrder `json:"nonSeasonalOrder,omitempty"`

	// The data frequency of a time series.
	DataFrequency *string `json:"dataFrequency,omitempty"`

	// Whether or not p-value test should be computed for this model. Only
	//  available for linear and logistic regression models.
	CalculatePValues *bool `json:"calculatePValues,omitempty"`

	// Include drift when fitting an ARIMA model.
	IncludeDrift *bool `json:"includeDrift,omitempty"`

	// The geographical region based on which the holidays are considered in
	//  time series modeling. If a valid value is specified, then holiday
	//  effects modeling is enabled.
	HolidayRegion *string `json:"holidayRegion,omitempty"`

	// A list of geographical regions that are used for time series modeling.
	HolidayRegions []string `json:"holidayRegions,omitempty"`

	// The time series id column that was used during ARIMA model training.
	TimeSeriesIDColumn *string `json:"timeSeriesIDColumn,omitempty"`

	// The time series id columns that were used during ARIMA model training.
	TimeSeriesIDColumns []string `json:"timeSeriesIDColumns,omitempty"`

	// The number of periods ahead that need to be forecasted.
	Horizon *int64 `json:"horizon,omitempty"`

	// The max value of the sum of non-seasonal p and q.
	AutoArimaMaxOrder *int64 `json:"autoArimaMaxOrder,omitempty"`

	// The min value of the sum of non-seasonal p and q.
	AutoArimaMinOrder *int64 `json:"autoArimaMinOrder,omitempty"`

	// Number of trials to run this hyperparameter tuning job.
	NumTrials *int64 `json:"numTrials,omitempty"`

	// Maximum number of trials to run in parallel.
	MaxParallelTrials *int64 `json:"maxParallelTrials,omitempty"`

	// The target evaluation metrics to optimize the hyperparameters for.
	HparamTuningObjectives []string `json:"hparamTuningObjectives,omitempty"`

	// If true, perform decompose time series and save the results.
	DecomposeTimeSeries *bool `json:"decomposeTimeSeries,omitempty"`

	// If true, clean spikes and dips in the input time series.
	CleanSpikesAndDips *bool `json:"cleanSpikesAndDips,omitempty"`

	// If true, detect step changes and make data adjustment in the input time
	//  series.
	AdjustStepChanges *bool `json:"adjustStepChanges,omitempty"`

	// If true, enable global explanation during training.
	EnableGlobalExplain *bool `json:"enableGlobalExplain,omitempty"`

	// Number of paths for the sampled Shapley explain method.
	SampledShapleyNumPaths *int64 `json:"sampledShapleyNumPaths,omitempty"`

	// Number of integral steps for the integrated gradients explain method.
	IntegratedGradientsNumSteps *int64 `json:"integratedGradientsNumSteps,omitempty"`

	// Categorical feature encoding method.
	CategoryEncodingMethod *string `json:"categoryEncodingMethod,omitempty"`

	// Based on the selected TF version, the corresponding docker image is
	//  used to train external models.
	TfVersion *string `json:"tfVersion,omitempty"`

	// Enums for color space, used for processing images in Object Table.
	//  See more details at
	//  https://www.tensorflow.org/io/tutorials/colorspace.
	ColorSpace *string `json:"colorSpace,omitempty"`

	// Name of the instance weight column for training data.
	//  This column isn't be used as a feature.
	InstanceWeightColumn *string `json:"instanceWeightColumn,omitempty"`

	// Smoothing window size for the trend component. When a positive value is
	//  specified, a center moving average smoothing is applied on the history
	//  trend. When the smoothing window is out of the boundary at the
	//  beginning or the end of the trend, the first element or the last
	//  element is padded to fill the smoothing window before the average is
	//  applied.
	TrendSmoothingWindowSize *int64 `json:"trendSmoothingWindowSize,omitempty"`

	// The fraction of the interpolated length of the time series that's used
	//  to model the time series trend component. All of the time points of the
	//  time series are used to model the non-trend component. This training
	//  option accelerates modeling training without sacrificing much
	//  forecasting accuracy. You can use this option with
	//  `minTimeSeriesLength` but not with `maxTimeSeriesLength`.
	TimeSeriesLengthFraction *float64 `json:"timeSeriesLengthFraction,omitempty"`

	// The minimum number of time points in a time series that are used in
	//  modeling the trend component of the time series. If you use this option
	//  you must also set the `timeSeriesLengthFraction` option. This training
	//  option ensures that enough time points are available when you use
	//  `timeSeriesLengthFraction` in trend modeling. This is particularly
	//  important when forecasting multiple time series in a single query using
	//  `timeSeriesIdColumn`. If the total number of time points is less than
	//  the `minTimeSeriesLength` value, then the query uses all available time
	//  points.
	MinTimeSeriesLength *int64 `json:"minTimeSeriesLength,omitempty"`

	// The maximum number of time points in a time series that can be used in
	//  modeling the trend component of the time series. Don't use this option
	//  with the `timeSeriesLengthFraction` or `minTimeSeriesLength` options.
	MaxTimeSeriesLength *int64 `json:"maxTimeSeriesLength,omitempty"`

	// User-selected XGBoost versions for training of XGBoost models.
	XgboostVersion *string `json:"xgboostVersion,omitempty"`

	// Whether to use approximate feature contribution method in XGBoost model
	//  explanation for global explain.
	ApproxGlobalFeatureContrib *bool `json:"approxGlobalFeatureContrib,omitempty"`

	// Whether the model should include intercept during model training.
	FitIntercept *bool `json:"fitIntercept,omitempty"`

	// Number of principal components to keep in the PCA model. Must be <= the
	//  number of features.
	NumPrincipalComponents *int64 `json:"numPrincipalComponents,omitempty"`

	// The minimum ratio of cumulative explained variance that needs to be
	//  given by the PCA model.
	PcaExplainedVarianceRatio *float64 `json:"pcaExplainedVarianceRatio,omitempty"`

	// If true, scale the feature values by dividing the feature standard
	//  deviation. Currently only apply to PCA.
	ScaleFeatures *bool `json:"scaleFeatures,omitempty"`

	// The solver for PCA.
	PcaSolver *string `json:"pcaSolver,omitempty"`

	// Whether to calculate class weights automatically based on the
	//  popularity of each label.
	AutoClassWeights *bool `json:"autoClassWeights,omitempty"`

	// Activation function of the neural nets.
	ActivationFn *string `json:"activationFn,omitempty"`

	// Optimizer used for training the neural nets.
	Optimizer *string `json:"optimizer,omitempty"`

	// Budget in hours for AutoML training.
	BudgetHours *float64 `json:"budgetHours,omitempty"`

	// Whether to standardize numerical features. Default to true.
	StandardizeFeatures *bool `json:"standardizeFeatures,omitempty"`

	// L1 regularization coefficient to activations.
	L1RegActivation *float64 `json:"l1RegActivation,omitempty"`

	// The model registry.
	ModelRegistry *string `json:"modelRegistry,omitempty"`

	// The version aliases to apply in Vertex AI model registry. Always
	//  overwrite if the version aliases exists in a existing model.
	VertexAiModelVersionAliases []string `json:"vertexAiModelVersionAliases,omitempty"`
}

// +kcc:proto=google.cloud.bigquery.v2.ModelReference
type ModelReference struct {
	// Required. The ID of the project containing this model.
	ProjectID *string `json:"projectID,omitempty"`

	// Required. The ID of the dataset containing this model.
	DatasetID *string `json:"datasetID,omitempty"`

	// Required. The ID of the model. The ID must contain only
	//  letters (a-z, A-Z), numbers (0-9), or underscores (_). The maximum
	//  length is 1,024 characters.
	ModelID *string `json:"modelID,omitempty"`
}

// +kcc:proto=google.cloud.bigquery.v2.ParquetOptions
type ParquetOptions struct {
	// Optional. Indicates whether to infer Parquet ENUM logical type as STRING
	//  instead of BYTES by default.
	EnumAsString *bool `json:"enumAsString,omitempty"`

	// Optional. Indicates whether to use schema inference specifically for
	//  Parquet LIST logical type.
	EnableListInference *bool `json:"enableListInference,omitempty"`

	// Optional. Indicates how to represent a Parquet map if present.
	MapTargetType *string `json:"mapTargetType,omitempty"`
}

// +kcc:proto=google.cloud.bigquery.v2.PartitionSkew
type PartitionSkew struct {
	// Output only. Source stages which produce skewed data.
	SkewSources []PartitionSkew_SkewSource `json:"skewSources,omitempty"`
}

// +kcc:proto=google.cloud.bigquery.v2.PartitionSkew.SkewSource
type PartitionSkew_SkewSource struct {
	// Output only. Stage id of the skew source stage.
	StageID *int64 `json:"stageID,omitempty"`
}

// +kcc:proto=google.cloud.bigquery.v2.PartitionedColumn
type PartitionedColumn struct {
	// Required. The name of the partition column.
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
	PartitionedColumn []PartitionedColumn `json:"partitionedColumn,omitempty"`
}

// +kcc:proto=google.cloud.bigquery.v2.PerformanceInsights
type PerformanceInsights struct {
	// Output only. Average execution ms of previous runs. Indicates the job ran
	//  slow compared to previous executions. To find previous executions, use
	//  INFORMATION_SCHEMA tables and filter jobs with same query hash.
	AvgPreviousExecutionMs *int64 `json:"avgPreviousExecutionMs,omitempty"`

	// Output only. Standalone query stage performance insights, for exploring
	//  potential improvements.
	StagePerformanceStandaloneInsights []StagePerformanceStandaloneInsight `json:"stagePerformanceStandaloneInsights,omitempty"`

	// Output only. Query stage performance insights compared to previous runs,
	//  for diagnosing performance regression.
	StagePerformanceChangeInsights []StagePerformanceChangeInsight `json:"stagePerformanceChangeInsights,omitempty"`
}

// +kcc:proto=google.cloud.bigquery.v2.PrimaryKey
type PrimaryKey struct {
	// Required. The columns that are composed of the primary key constraint.
	Columns []string `json:"columns,omitempty"`
}

// +kcc:proto=google.cloud.bigquery.v2.PrivacyPolicy
type PrivacyPolicy struct {
	// Optional. Policy used for aggregation thresholds.
	AggregationThresholdPolicy *AggregationThresholdPolicy `json:"aggregationThresholdPolicy,omitempty"`

	// Optional. Policy used for differential privacy.
	DifferentialPrivacyPolicy *DifferentialPrivacyPolicy `json:"differentialPrivacyPolicy,omitempty"`

	// Optional. Join restriction policy is outside of the one of policies, since
	//  this policy can be set along with other policies. This policy gives data
	//  providers the ability to enforce joins on the 'join_allowed_columns' when
	//  data is queried from a privacy protected view.
	JoinRestrictionPolicy *JoinRestrictionPolicy `json:"joinRestrictionPolicy,omitempty"`
}

// +kcc:proto=google.cloud.bigquery.v2.QueryInfo
type QueryInfo struct {
	// Output only. Information about query optimizations.
	OptimizationDetails *google_protobuf_Struct `json:"optimizationDetails,omitempty"`
}

// +kcc:proto=google.cloud.bigquery.v2.QueryParameter
type QueryParameter struct {
	// Optional. If unset, this is a positional parameter. Otherwise, should be
	//  unique within a query.
	Name *string `json:"name,omitempty"`

	// Required. The type of this parameter.
	ParameterType *QueryParameterType `json:"parameterType,omitempty"`

	// Required. The value of this parameter.
	ParameterValue *QueryParameterValue `json:"parameterValue,omitempty"`
}

// +kcc:proto=google.cloud.bigquery.v2.QueryParameterStructType
type QueryParameterStructType struct {
	// Optional. The name of this field.
	Name *string `json:"name,omitempty"`

	// Required. The type of this field.
	Type *QueryParameterType `json:"type,omitempty"`

	// Optional. Human-oriented description of the field.
	Description *string `json:"description,omitempty"`
}

// +kcc:proto=google.cloud.bigquery.v2.QueryParameterType
type QueryParameterType struct {
	// Required. The top level type of this field.
	Type *string `json:"type,omitempty"`

	// Optional. The type of the array's elements, if this is an array.
	ArrayType *QueryParameterType `json:"arrayType,omitempty"`

	// Optional. The types of the fields of this struct, in order, if this is a
	//  struct.
	StructTypes []QueryParameterStructType `json:"structTypes,omitempty"`

	// Optional. The element type of the range, if this is a range.
	RangeElementType *QueryParameterType `json:"rangeElementType,omitempty"`
}

// +kcc:proto=google.cloud.bigquery.v2.QueryParameterValue
type QueryParameterValue struct {
	// Optional. The value of this value, if a simple scalar type.
	Value *string `json:"value,omitempty"`

	// Optional. The array values, if this is an array type.
	ArrayValues []QueryParameterValue `json:"arrayValues,omitempty"`

	// TODO: map type string message for struct_values

	// Optional. The range value, if this is a range type.
	RangeValue *RangeValue `json:"rangeValue,omitempty"`

	// This field should not be used.
	AltStructValues []google_protobuf_Value `json:"altStructValues,omitempty"`
}

// +kcc:proto=google.cloud.bigquery.v2.QueryTimelineSample
type QueryTimelineSample struct {
	// Milliseconds elapsed since the start of query execution.
	ElapsedMs *int64 `json:"elapsedMs,omitempty"`

	// Cumulative slot-ms consumed by the query.
	TotalSlotMs *int64 `json:"totalSlotMs,omitempty"`

	// Total units of work remaining for the query. This number can be revised
	//  (increased or decreased) while the query is running.
	PendingUnits *int64 `json:"pendingUnits,omitempty"`

	// Total parallel units of work completed by this query.
	CompletedUnits *int64 `json:"completedUnits,omitempty"`

	// Total number of active workers. This does not correspond directly to
	//  slot usage. This is the largest value observed since the last sample.
	ActiveUnits *int64 `json:"activeUnits,omitempty"`

	// Units of work that can be scheduled immediately. Providing additional slots
	//  for these units of work will accelerate the query, if no other query in
	//  the reservation needs additional slots.
	EstimatedRunnableUnits *int64 `json:"estimatedRunnableUnits,omitempty"`
}

// +kcc:proto=google.cloud.bigquery.v2.RangePartitioning
type RangePartitioning struct {
	// Required. The name of the column to partition the table on. It must be a
	//  top-level, INT64 column whose mode is NULLABLE or REQUIRED.
	Field *string `json:"field,omitempty"`

	// Defines the ranges for range partitioning.
	Range *RangePartitioning_Range `json:"range,omitempty"`
}

// +kcc:proto=google.cloud.bigquery.v2.RangePartitioning.Range
type RangePartitioning_Range struct {
	// Required. The start of range partitioning, inclusive. This field is an
	//  INT64 value represented as a string.
	Start *string `json:"start,omitempty"`

	// Required. The end of range partitioning, exclusive. This field is an
	//  INT64 value represented as a string.
	End *string `json:"end,omitempty"`

	// Required. The width of each interval. This field is an INT64 value
	//  represented as a string.
	Interval *string `json:"interval,omitempty"`
}

// +kcc:proto=google.cloud.bigquery.v2.RangeValue
type RangeValue struct {
	// Optional. The start value of the range. A missing value represents an
	//  unbounded start.
	Start *QueryParameterValue `json:"start,omitempty"`

	// Optional. The end value of the range. A missing value represents an
	//  unbounded end.
	End *QueryParameterValue `json:"end,omitempty"`
}

// +kcc:proto=google.cloud.bigquery.v2.RemoteModelInfo
type RemoteModelInfo struct {
	// Output only. The endpoint for remote model.
	Endpoint *string `json:"endpoint,omitempty"`

	// Output only. The remote service type for remote model.
	RemoteServiceType *string `json:"remoteServiceType,omitempty"`

	// Output only. Fully qualified name of the user-provided connection object of
	//  the remote model. Format:
	//  ```"projects/{project_id}/locations/{location_id}/connections/{connection_id}"```
	Connection *string `json:"connection,omitempty"`

	// Output only. Max number of rows in each batch sent to the remote service.
	//  If unset, the number of rows in each batch is set dynamically.
	MaxBatchingRows *int64 `json:"maxBatchingRows,omitempty"`

	// Output only. The model version for LLM.
	RemoteModelVersion *string `json:"remoteModelVersion,omitempty"`

	// Output only. The name of the speech recognizer to use for speech
	//  recognition. The expected format is
	//  `projects/{project}/locations/{location}/recognizers/{recognizer}`.
	//  Customers can specify this field at model creation. If not specified, a
	//  default recognizer `projects/{model
	//  project}/locations/global/recognizers/_` will be used. See more details at
	//  [recognizers](https://cloud.google.com/speech-to-text/v2/docs/reference/rest/v2/projects.locations.recognizers)
	SpeechRecognizer *string `json:"speechRecognizer,omitempty"`
}

// +kcc:proto=google.cloud.bigquery.v2.RestrictionConfig
type RestrictionConfig struct {
	// Output only. Specifies the type of dataset/table restriction.
	Type *string `json:"type,omitempty"`
}

// +kcc:proto=google.cloud.bigquery.v2.Routine
type Routine struct {
	// Output only. A hash of this resource.
	Etag *string `json:"etag,omitempty"`

	// Required. Reference describing the ID of this routine.
	RoutineReference *RoutineReference `json:"routineReference,omitempty"`

	// Required. The type of routine.
	RoutineType *string `json:"routineType,omitempty"`

	// Output only. The time when this routine was created, in milliseconds since
	//  the epoch.
	CreationTime *int64 `json:"creationTime,omitempty"`

	// Output only. The time when this routine was last modified, in milliseconds
	//  since the epoch.
	LastModifiedTime *int64 `json:"lastModifiedTime,omitempty"`

	// Optional. Defaults to "SQL" if remote_function_options field is absent, not
	//  set otherwise.
	Language *string `json:"language,omitempty"`

	// Optional.
	Arguments []Routine_Argument `json:"arguments,omitempty"`

	// Optional if language = "SQL"; required otherwise.
	//  Cannot be set if routine_type = "TABLE_VALUED_FUNCTION".
	//
	//  If absent, the return type is inferred from definition_body at query time
	//  in each query that references this routine. If present, then the evaluated
	//  result will be cast to the specified returned type at query time.
	//
	//  For example, for the functions created with the following statements:
	//
	//  * `CREATE FUNCTION Add(x FLOAT64, y FLOAT64) RETURNS FLOAT64 AS (x + y);`
	//
	//  * `CREATE FUNCTION Increment(x FLOAT64) AS (Add(x, 1));`
	//
	//  * `CREATE FUNCTION Decrement(x FLOAT64) RETURNS FLOAT64 AS (Add(x, -1));`
	//
	//  The return_type is `{type_kind: "FLOAT64"}` for `Add` and `Decrement`, and
	//  is absent for `Increment` (inferred as FLOAT64 at query time).
	//
	//  Suppose the function `Add` is replaced by
	//    `CREATE OR REPLACE FUNCTION Add(x INT64, y INT64) AS (x + y);`
	//
	//  Then the inferred return type of `Increment` is automatically changed to
	//  INT64 at query time, while the return type of `Decrement` remains FLOAT64.
	ReturnType *StandardSqlDataType `json:"returnType,omitempty"`

	// Optional. Can be set only if routine_type = "TABLE_VALUED_FUNCTION".
	//
	//  If absent, the return table type is inferred from definition_body at query
	//  time in each query that references this routine. If present, then the
	//  columns in the evaluated table result will be cast to match the column
	//  types specified in return table type, at query time.
	ReturnTableType *StandardSqlTableType `json:"returnTableType,omitempty"`

	// Optional. If language = "JAVASCRIPT", this field stores the path of the
	//  imported JAVASCRIPT libraries.
	ImportedLibraries []string `json:"importedLibraries,omitempty"`

	// Required. The body of the routine.
	//
	//  For functions, this is the expression in the AS clause.
	//
	//  If language=SQL, it is the substring inside (but excluding) the
	//  parentheses. For example, for the function created with the following
	//  statement:
	//
	//  `CREATE FUNCTION JoinLines(x string, y string) as (concat(x, "\n", y))`
	//
	//  The definition_body is `concat(x, "\n", y)` (\n is not replaced with
	//  linebreak).
	//
	//  If language=JAVASCRIPT, it is the evaluated string in the AS clause.
	//  For example, for the function created with the following statement:
	//
	//  `CREATE FUNCTION f() RETURNS STRING LANGUAGE js AS 'return "\n";\n'`
	//
	//  The definition_body is
	//
	//  `return "\n";\n`
	//
	//  Note that both \n are replaced with linebreaks.
	DefinitionBody *string `json:"definitionBody,omitempty"`

	// Optional. The description of the routine, if defined.
	Description *string `json:"description,omitempty"`

	// Optional. The determinism level of the JavaScript UDF, if defined.
	DeterminismLevel *string `json:"determinismLevel,omitempty"`

	// Optional. The security mode of the routine, if defined. If not defined, the
	//  security mode is automatically determined from the routine's configuration.
	SecurityMode *string `json:"securityMode,omitempty"`

	// Optional. Use this option to catch many common errors. Error checking is
	//  not exhaustive, and successfully creating a procedure doesn't guarantee
	//  that the procedure will successfully execute at runtime. If `strictMode` is
	//  set to `TRUE`, the procedure body is further checked for errors such as
	//  non-existent tables or columns. The `CREATE PROCEDURE` statement fails if
	//  the body fails any of these checks.
	//
	//  If `strictMode` is set to `FALSE`, the procedure body is checked only for
	//  syntax. For procedures that invoke themselves recursively, specify
	//  `strictMode=FALSE` to avoid non-existent procedure errors during
	//  validation.
	//
	//  Default value is `TRUE`.
	StrictMode *bool `json:"strictMode,omitempty"`

	// Optional. Remote function specific options.
	RemoteFunctionOptions *Routine_RemoteFunctionOptions `json:"remoteFunctionOptions,omitempty"`

	// Optional. Spark specific options.
	SparkOptions *SparkOptions `json:"sparkOptions,omitempty"`

	// Optional. If set to `DATA_MASKING`, the function is validated and made
	//  available as a masking function. For more information, see [Create custom
	//  masking
	//  routines](https://cloud.google.com/bigquery/docs/user-defined-functions#custom-mask).
	DataGovernanceType *string `json:"dataGovernanceType,omitempty"`
}

// +kcc:proto=google.cloud.bigquery.v2.Routine.Argument
type Routine_Argument struct {
	// Optional. The name of this argument. Can be absent for function return
	//  argument.
	Name *string `json:"name,omitempty"`

	// Optional. Defaults to FIXED_TYPE.
	ArgumentKind *string `json:"argumentKind,omitempty"`

	// Optional. Specifies whether the argument is input or output.
	//  Can be set for procedures only.
	Mode *string `json:"mode,omitempty"`

	// Required unless argument_kind = ANY_TYPE.
	DataType *StandardSqlDataType `json:"dataType,omitempty"`

	// Optional. Whether the argument is an aggregate function parameter.
	//  Must be Unset for routine types other than AGGREGATE_FUNCTION.
	//  For AGGREGATE_FUNCTION, if set to false, it is equivalent to adding "NOT
	//  AGGREGATE" clause in DDL; Otherwise, it is equivalent to omitting "NOT
	//  AGGREGATE" clause in DDL.
	IsAggregate *bool `json:"isAggregate,omitempty"`
}

// +kcc:proto=google.cloud.bigquery.v2.Routine.RemoteFunctionOptions
type Routine_RemoteFunctionOptions struct {
	// Endpoint of the user-provided remote service, e.g.
	//  ```https://us-east1-my_gcf_project.cloudfunctions.net/remote_add```
	Endpoint *string `json:"endpoint,omitempty"`

	// Fully qualified name of the user-provided connection object which holds
	//  the authentication information to send requests to the remote service.
	//  Format:
	//  ```"projects/{projectId}/locations/{locationId}/connections/{connectionId}"```
	Connection *string `json:"connection,omitempty"`

	// User-defined context as a set of key/value pairs, which will be sent as
	//  function invocation context together with batched arguments in the
	//  requests to the remote service. The total number of bytes of keys and
	//  values must be less than 8KB.
	UserDefinedContext map[string]string `json:"userDefinedContext,omitempty"`

	// Max number of rows in each batch sent to the remote service.
	//  If absent or if 0, BigQuery dynamically decides the number of rows in a
	//  batch.
	MaxBatchingRows *int64 `json:"maxBatchingRows,omitempty"`
}

// +kcc:proto=google.cloud.bigquery.v2.RoutineReference
type RoutineReference struct {
	// Required. The ID of the project containing this routine.
	ProjectId *string `json:"projectId"`

	// Required. The ID of the dataset containing this routine.
	DatasetId *string `json:"datasetId"`

	// Required. The ID of the routine. The ID must contain only
	//  letters (a-z, A-Z), numbers (0-9), or underscores (_). The maximum
	//  length is 256 characters.
	RoutineId *string `json:"routineId"`
}

// +kcc:proto=google.cloud.bigquery.v2.RowAccessPolicy
type RowAccessPolicy struct {
	// Output only. A hash of this resource.
	Etag *string `json:"etag,omitempty"`

	// Required. Reference describing the ID of this row access policy.
	RowAccessPolicyReference *RowAccessPolicyReference `json:"rowAccessPolicyReference,omitempty"`

	// Required. A SQL boolean expression that represents the rows defined by this
	//  row access policy, similar to the boolean expression in a WHERE clause of a
	//  SELECT query on a table.
	//  References to other tables, routines, and temporary functions are not
	//  supported.
	//
	//  Examples: region="EU"
	//            date_field = CAST('2019-9-27' as DATE)
	//            nullable_field is not NULL
	//            numeric_field BETWEEN 1.0 AND 5.0
	FilterPredicate *string `json:"filterPredicate,omitempty"`

	// Output only. The time when this row access policy was created, in
	//  milliseconds since the epoch.
	CreationTime *string `json:"creationTime,omitempty"`

	// Output only. The time when this row access policy was last modified, in
	//  milliseconds since the epoch.
	LastModifiedTime *string `json:"lastModifiedTime,omitempty"`
}

// +kcc:proto=google.cloud.bigquery.v2.RowAccessPolicyReference
type RowAccessPolicyReference struct {
	// Required. The ID of the project containing this row access policy.
	ProjectID *string `json:"projectID,omitempty"`

	// Required. The ID of the dataset containing this row access policy.
	DatasetID *string `json:"datasetID,omitempty"`

	// Required. The ID of the table containing this row access policy.
	TableID *string `json:"tableID,omitempty"`

	// Required. The ID of the row access policy. The ID must contain only
	//  letters (a-z, A-Z), numbers (0-9), or underscores (_). The maximum
	//  length is 256 characters.
	PolicyID *string `json:"policyID,omitempty"`
}

// +kcc:proto=google.cloud.bigquery.v2.RowLevelSecurityStatistics
type RowLevelSecurityStatistics struct {
	// Whether any accessed data was protected by row access policies.
	RowLevelSecurityApplied *bool `json:"rowLevelSecurityApplied,omitempty"`
}

// +kcc:proto=google.cloud.bigquery.v2.ScriptOptions
type ScriptOptions struct {
	// Timeout period for each statement in a script.
	StatementTimeoutMs *int64 `json:"statementTimeoutMs,omitempty"`

	// Limit on the number of bytes billed per statement. Exceeding this budget
	//  results in an error.
	StatementByteBudget *int64 `json:"statementByteBudget,omitempty"`

	// Determines which statement in the script represents the "key result",
	//  used to populate the schema and query results of the script job.
	//  Default is LAST.
	KeyResultStatement *string `json:"keyResultStatement,omitempty"`
}

// +kcc:proto=google.cloud.bigquery.v2.ScriptStatistics
type ScriptStatistics struct {
	// Whether this child job was a statement or expression.
	EvaluationKind *string `json:"evaluationKind,omitempty"`

	// Stack trace showing the line/column/procedure name of each frame on the
	//  stack at the point where the current evaluation happened. The leaf frame
	//  is first, the primary script is last. Never empty.
	StackFrames []ScriptStatistics_ScriptStackFrame `json:"stackFrames,omitempty"`
}

// +kcc:proto=google.cloud.bigquery.v2.ScriptStatistics.ScriptStackFrame
type ScriptStatistics_ScriptStackFrame struct {
	// Output only. One-based start line.
	StartLine *int32 `json:"startLine,omitempty"`

	// Output only. One-based start column.
	StartColumn *int32 `json:"startColumn,omitempty"`

	// Output only. One-based end line.
	EndLine *int32 `json:"endLine,omitempty"`

	// Output only. One-based end column.
	EndColumn *int32 `json:"endColumn,omitempty"`

	// Output only. Name of the active procedure, empty if in a top-level
	//  script.
	ProcedureID *string `json:"procedureID,omitempty"`

	// Output only. Text of the current statement/expression.
	Text *string `json:"text,omitempty"`
}

// +kcc:proto=google.cloud.bigquery.v2.SearchStatistics
type SearchStatistics struct {
	// Specifies the index usage mode for the query.
	IndexUsageMode *string `json:"indexUsageMode,omitempty"`

	// When `indexUsageMode` is `UNUSED` or `PARTIALLY_USED`, this field explains
	//  why indexes were not used in all or part of the search query. If
	//  `indexUsageMode` is `FULLY_USED`, this field is not populated.
	IndexUnusedReasons []IndexUnusedReason `json:"indexUnusedReasons,omitempty"`
}

// +kcc:proto=google.cloud.bigquery.v2.SerDeInfo
type SerDeInfo struct {
	// Optional. Name of the SerDe.
	//  The maximum length is 256 characters.
	Name *string `json:"name,omitempty"`

	// Required. Specifies a fully-qualified class name of the serialization
	//  library that is responsible for the translation of data between table
	//  representation and the underlying low-level input and output format
	//  structures. The maximum length is 256 characters.
	SerializationLibrary *string `json:"serializationLibrary,omitempty"`

	// Optional. Key-value pairs that define the initialization parameters for the
	//  serialization library.
	//  Maximum size 10 Kib.
	Parameters map[string]string `json:"parameters,omitempty"`
}

// +kcc:proto=google.cloud.bigquery.v2.SessionInfo
type SessionInfo struct {
	// Output only. The id of the session.
	SessionID *string `json:"sessionID,omitempty"`
}

// +kcc:proto=google.cloud.bigquery.v2.SnapshotDefinition
type SnapshotDefinition struct {
	// Required. Reference describing the ID of the table that was snapshot.
	BaseTableReference *TableReference `json:"baseTableReference,omitempty"`

	// Required. The time at which the base table was snapshot. This value is
	//  reported in the JSON response using RFC3339 format.
	SnapshotTime *string `json:"snapshotTime,omitempty"`
}

// +kcc:proto=google.cloud.bigquery.v2.SparkOptions
type SparkOptions struct {
	// Fully qualified name of the user-provided Spark connection object. Format:
	//  ```"projects/{project_id}/locations/{location_id}/connections/{connection_id}"```
	Connection *string `json:"connection,omitempty"`

	// Runtime version. If not specified, the default runtime version is used.
	RuntimeVersion *string `json:"runtimeVersion,omitempty"`

	// Custom container image for the runtime environment.
	ContainerImage *string `json:"containerImage,omitempty"`

	// Configuration properties as a set of key/value pairs, which will be passed
	//  on to the Spark application. For more information, see
	//  [Apache Spark](https://spark.apache.org/docs/latest/index.html) and the
	//  [procedure option
	//  list](https://cloud.google.com/bigquery/docs/reference/standard-sql/data-definition-language#procedure_option_list).
	Properties map[string]string `json:"properties,omitempty"`

	// The main file/jar URI of the Spark application. Exactly one of the
	//  definition_body field and the main_file_uri field must be set for Python.
	//  Exactly one of main_class and main_file_uri field
	//  should be set for Java/Scala language type.
	MainFileUri *string `json:"mainFileUri,omitempty"`

	// Python files to be placed on the PYTHONPATH for PySpark application.
	//  Supported file types: `.py`, `.egg`, and `.zip`. For more information
	//  about Apache Spark, see
	//  [Apache Spark](https://spark.apache.org/docs/latest/index.html).
	PyFileUris []string `json:"pyFileUris,omitempty"`

	// JARs to include on the driver and executor CLASSPATH.
	//  For more information about Apache Spark, see
	//  [Apache Spark](https://spark.apache.org/docs/latest/index.html).
	JarUris []string `json:"jarUris,omitempty"`

	// Files to be placed in the working directory of each executor.
	//  For more information about Apache Spark, see
	//  [Apache Spark](https://spark.apache.org/docs/latest/index.html).
	FileUris []string `json:"fileUris,omitempty"`

	// Archive files to be extracted into the working directory of each executor.
	//  For more information about Apache Spark, see
	//  [Apache Spark](https://spark.apache.org/docs/latest/index.html).
	ArchiveUris []string `json:"archiveUris,omitempty"`

	// The fully qualified name of a class in jar_uris, for example,
	//  com.example.wordcount. Exactly one of main_class and main_jar_uri field
	//   should be set for Java/Scala language type.
	MainClass *string `json:"mainClass,omitempty"`
}

// +kcc:proto=google.cloud.bigquery.v2.SparkStatistics
type SparkStatistics struct {
	// Output only. Spark job ID if a Spark job is created successfully.
	SparkJobID *string `json:"sparkJobID,omitempty"`

	// Output only. Location where the Spark job is executed.
	//  A location is selected by BigQueury for jobs configured to run in a
	//  multi-region.
	SparkJobLocation *string `json:"sparkJobLocation,omitempty"`

	// Output only. Endpoints returned from Dataproc.
	//  Key list:
	//   - history_server_endpoint: A link to Spark job UI.
	Endpoints map[string]string `json:"endpoints,omitempty"`

	// Output only. Logging info is used to generate a link to Cloud Logging.
	LoggingInfo *SparkStatistics_LoggingInfo `json:"loggingInfo,omitempty"`

	// Output only. The Cloud KMS encryption key that is used to protect the
	//  resources created by the Spark job. If the Spark procedure uses the invoker
	//  security mode, the Cloud KMS encryption key is either inferred from the
	//  provided system variable,
	//  `@@spark_proc_properties.kms_key_name`, or the default key of the BigQuery
	//  job's project (if the CMEK organization policy is enforced). Otherwise, the
	//  Cloud KMS key is either inferred from the Spark connection associated with
	//  the procedure (if it is provided), or from the default key of the Spark
	//  connection's project if the CMEK organization policy is enforced.
	//
	//  Example:
	//
	//  * `projects/[kms_project_id]/locations/[region]/keyRings/[key_region]/cryptoKeys/[key]`
	KmsKeyName *string `json:"kmsKeyName,omitempty"`

	// Output only. The Google Cloud Storage bucket that is used as the default
	//  file system by the Spark application. This field is only filled when the
	//  Spark procedure uses the invoker security mode. The `gcsStagingBucket`
	//  bucket is inferred from the `@@spark_proc_properties.staging_bucket` system
	//  variable (if it is provided). Otherwise, BigQuery creates a default staging
	//  bucket for the job and returns the bucket name in this field.
	//
	//  Example:
	//
	//  * `gs://[bucket_name]`
	GcsStagingBucket *string `json:"gcsStagingBucket,omitempty"`
}

// +kcc:proto=google.cloud.bigquery.v2.SparkStatistics.LoggingInfo
type SparkStatistics_LoggingInfo struct {
	// Output only. Resource type used for logging.
	ResourceType *string `json:"resourceType,omitempty"`

	// Output only. Project ID where the Spark logs were written.
	ProjectID *string `json:"projectID,omitempty"`
}

// +kcc:proto=google.cloud.bigquery.v2.StagePerformanceChangeInsight
type StagePerformanceChangeInsight struct {
	// Output only. The stage id that the insight mapped to.
	StageID *int64 `json:"stageID,omitempty"`

	// Output only. Input data change insight of the query stage.
	InputDataChange *InputDataChange `json:"inputDataChange,omitempty"`
}

// +kcc:proto=google.cloud.bigquery.v2.StagePerformanceStandaloneInsight
type StagePerformanceStandaloneInsight struct {
	// Output only. The stage id that the insight mapped to.
	StageID *int64 `json:"stageID,omitempty"`

	// Output only. True if the stage has a slot contention issue.
	SlotContention *bool `json:"slotContention,omitempty"`

	// Output only. True if the stage has insufficient shuffle quota.
	InsufficientShuffleQuota *bool `json:"insufficientShuffleQuota,omitempty"`

	// Output only. If present, the stage had the following reasons for being
	//  disqualified from BI Engine execution.
	BiEngineReasons []BiEngineReason `json:"biEngineReasons,omitempty"`

	// Output only. High cardinality joins in the stage.
	HighCardinalityJoins []HighCardinalityJoin `json:"highCardinalityJoins,omitempty"`

	// Output only. Partition skew in the stage.
	PartitionSkew *PartitionSkew `json:"partitionSkew,omitempty"`
}

// +kcc:proto=google.cloud.bigquery.v2.StandardSqlDataType
type StandardSqlDataType struct {
	// Required. The top level type of this field.
	//  Can be any GoogleSQL data type (e.g., "INT64", "DATE", "ARRAY").
	TypeKind *string `json:"typeKind,omitempty"`

	// The type of the array's elements, if type_kind = "ARRAY".
	ArrayElementType *StandardSqlDataType `json:"arrayElementType,omitempty"`

	// The fields of this struct, in order, if type_kind = "STRUCT".
	StructType *StandardSqlStructType `json:"structType,omitempty"`

	// The type of the range's elements, if type_kind = "RANGE".
	RangeElementType *StandardSqlDataType `json:"rangeElementType,omitempty"`
}

// +kcc:proto=google.cloud.bigquery.v2.StandardSqlField
type StandardSqlField struct {
	// Optional. The name of this field. Can be absent for struct fields.
	Name *string `json:"name,omitempty"`

	// Optional. The type of this parameter. Absent if not explicitly
	//  specified (e.g., CREATE FUNCTION statement can omit the return type;
	//  in this case the output parameter does not have this "type" field).
	Type *StandardSqlDataType `json:"type,omitempty"`
}

// +kcc:proto=google.cloud.bigquery.v2.StandardSqlStructType
type StandardSqlStructType struct {
	// Fields within the struct.
	Fields []StandardSqlField `json:"fields,omitempty"`
}

// +kcc:proto=google.cloud.bigquery.v2.StandardSqlTableType
type StandardSqlTableType struct {
	// The columns in this table type
	Columns []StandardSqlField `json:"columns,omitempty"`
}

// +kcc:proto=google.cloud.bigquery.v2.StorageDescriptor
type StorageDescriptor struct {
	// Optional. The physical location of the table
	//  (e.g. 'gs://spark-dataproc-data/pangea-data/case_sensitive/' or
	//  'gs://spark-dataproc-data/pangea-data/*').
	//  The maximum length is 2056 bytes.
	LocationUri *string `json:"locationUri,omitempty"`

	// Optional. Specifies the fully qualified class name of the InputFormat
	//  (e.g. "org.apache.hadoop.hive.ql.io.orc.OrcInputFormat").
	//  The maximum length is 128 characters.
	InputFormat *string `json:"inputFormat,omitempty"`

	// Optional. Specifies the fully qualified class name of the OutputFormat
	//  (e.g. "org.apache.hadoop.hive.ql.io.orc.OrcOutputFormat").
	//  The maximum length is 128 characters.
	OutputFormat *string `json:"outputFormat,omitempty"`

	// Optional. Serializer and deserializer information.
	SerdeInfo *SerDeInfo `json:"serdeInfo,omitempty"`
}

// +kcc:proto=google.cloud.bigquery.v2.Streamingbuffer
type Streamingbuffer struct {
	// Output only. A lower-bound estimate of the number of bytes currently in
	//  the streaming buffer.
	EstimatedBytes *uint64 `json:"estimatedBytes,omitempty"`

	// Output only. A lower-bound estimate of the number of rows currently in the
	//  streaming buffer.
	EstimatedRows *uint64 `json:"estimatedRows,omitempty"`

	// Output only. Contains the timestamp of the oldest entry in the streaming
	//  buffer, in milliseconds since the epoch, if the streaming buffer is
	//  available.
	OldestEntryTime *uint64 `json:"oldestEntryTime,omitempty"`
}

// +kcc:proto=google.cloud.bigquery.v2.SystemVariables
type SystemVariables struct {

	// TODO: map type string message for types

	// Output only. Value for each system variable.
	Values *google_protobuf_Struct `json:"values,omitempty"`
}

// +kcc:proto=google.cloud.bigquery.v2.Table
type Table struct {
	// The type of resource ID.
	Kind *string `json:"kind,omitempty"`

	// Output only. A hash of this resource.
	Etag *string `json:"etag,omitempty"`

	// Output only. An opaque ID uniquely identifying the table.
	ID *string `json:"id,omitempty"`

	// Output only. A URL that can be used to access this resource again.
	SelfLink *string `json:"selfLink,omitempty"`

	// Required. Reference describing the ID of this table.
	TableReference *TableReference `json:"tableReference,omitempty"`

	// Optional. A descriptive name for this table.
	FriendlyName *string `json:"friendlyName,omitempty"`

	// Optional. A user-friendly description of this table.
	Description *string `json:"description,omitempty"`

	// The labels associated with this table. You can use these to organize and
	//  group your tables. Label keys and values can be no longer than 63
	//  characters, can only contain lowercase letters, numeric characters,
	//  underscores and dashes. International characters are allowed. Label values
	//  are optional. Label keys must start with a letter and each label in the
	//  list must have a different key.
	Labels map[string]string `json:"labels,omitempty"`

	// Optional. Describes the schema of this table.
	Schema *TableSchema `json:"schema,omitempty"`

	// If specified, configures time-based partitioning for this table.
	TimePartitioning *TimePartitioning `json:"timePartitioning,omitempty"`

	// If specified, configures range partitioning for this table.
	RangePartitioning *RangePartitioning `json:"rangePartitioning,omitempty"`

	// Clustering specification for the table. Must be specified with time-based
	//  partitioning, data in the table will be first partitioned and subsequently
	//  clustered.
	Clustering *Clustering `json:"clustering,omitempty"`

	// Optional. If set to true, queries over this table require
	//  a partition filter that can be used for partition elimination to be
	//  specified.
	RequirePartitionFilter *bool `json:"requirePartitionFilter,omitempty"`

	// Optional. The partition information for all table formats, including
	//  managed partitioned tables, hive partitioned tables, iceberg partitioned,
	//  and metastore partitioned tables. This field is only populated for
	//  metastore partitioned tables. For other table formats, this is an output
	//  only field.
	PartitionDefinition *PartitioningDefinition `json:"partitionDefinition,omitempty"`

	// Output only. The size of this table in logical bytes, excluding any data in
	//  the streaming buffer.
	NumBytes *int64 `json:"numBytes,omitempty"`

	// Output only. The physical size of this table in bytes. This includes
	//  storage used for time travel.
	NumPhysicalBytes *int64 `json:"numPhysicalBytes,omitempty"`

	// Output only. The number of logical bytes in the table that are considered
	//  "long-term storage".
	NumLongTermBytes *int64 `json:"numLongTermBytes,omitempty"`

	// Output only. The number of rows of data in this table, excluding any data
	//  in the streaming buffer.
	NumRows *uint64 `json:"numRows,omitempty"`

	// Output only. The time when this table was created, in milliseconds since
	//  the epoch.
	CreationTime *int64 `json:"creationTime,omitempty"`

	// Optional. The time when this table expires, in milliseconds since the
	//  epoch. If not present, the table will persist indefinitely. Expired tables
	//  will be deleted and their storage reclaimed.  The defaultTableExpirationMs
	//  property of the encapsulating dataset can be used to set a default
	//  expirationTime on newly created tables.
	ExpirationTime *int64 `json:"expirationTime,omitempty"`

	// Output only. The time when this table was last modified, in milliseconds
	//  since the epoch.
	LastModifiedTime *uint64 `json:"lastModifiedTime,omitempty"`

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
	Type *string `json:"type,omitempty"`

	// Optional. The view definition.
	View *ViewDefinition `json:"view,omitempty"`

	// Optional. The materialized view definition.
	MaterializedView *MaterializedViewDefinition `json:"materializedView,omitempty"`

	// Output only. The materialized view status.
	MaterializedViewStatus *MaterializedViewStatus `json:"materializedViewStatus,omitempty"`

	// Optional. Describes the data format, location, and other properties of
	//  a table stored outside of BigQuery. By defining these properties, the data
	//  source can then be queried as if it were a standard BigQuery table.
	ExternalDataConfiguration *ExternalDataConfiguration `json:"externalDataConfiguration,omitempty"`

	// Optional. Specifies the configuration of a BigLake managed table.
	BiglakeConfiguration *BigLakeConfiguration `json:"biglakeConfiguration,omitempty"`

	// Output only. The geographic location where the table resides. This value
	//  is inherited from the dataset.
	Location *string `json:"location,omitempty"`

	// Output only. Contains information regarding this table's streaming buffer,
	//  if one is present. This field will be absent if the table is not being
	//  streamed to or if there is no data in the streaming buffer.
	StreamingBuffer *Streamingbuffer `json:"streamingBuffer,omitempty"`

	// Custom encryption configuration (e.g., Cloud KMS keys).
	EncryptionConfiguration *EncryptionConfiguration `json:"encryptionConfiguration,omitempty"`

	// Output only. Contains information about the snapshot. This value is set via
	//  snapshot creation.
	SnapshotDefinition *SnapshotDefinition `json:"snapshotDefinition,omitempty"`

	// Optional. Defines the default collation specification of new STRING fields
	//  in the table. During table creation or update, if a STRING field is added
	//  to this table without explicit collation specified, then the table inherits
	//  the table default collation. A change to this field affects only fields
	//  added afterwards, and does not alter the existing fields.
	//  The following values are supported:
	//
	//  * 'und:ci': undetermined locale, case insensitive.
	//  * '': empty string. Default to case-sensitive behavior.
	DefaultCollation *string `json:"defaultCollation,omitempty"`

	// Optional. Defines the default rounding mode specification of new decimal
	//  fields (NUMERIC OR BIGNUMERIC) in the table. During table creation or
	//  update, if a decimal field is added to this table without an explicit
	//  rounding mode specified, then the field inherits the table default
	//  rounding mode. Changing this field doesn't affect existing fields.
	DefaultRoundingMode *string `json:"defaultRoundingMode,omitempty"`

	// Output only. Contains information about the clone. This value is set via
	//  the clone operation.
	CloneDefinition *CloneDefinition `json:"cloneDefinition,omitempty"`

	// Output only. Number of physical bytes used by time travel storage (deleted
	//  or changed data). This data is not kept in real time, and might be delayed
	//  by a few seconds to a few minutes.
	NumTimeTravelPhysicalBytes *int64 `json:"numTimeTravelPhysicalBytes,omitempty"`

	// Output only. Total number of logical bytes in the table or materialized
	//  view.
	NumTotalLogicalBytes *int64 `json:"numTotalLogicalBytes,omitempty"`

	// Output only. Number of logical bytes that are less than 90 days old.
	NumActiveLogicalBytes *int64 `json:"numActiveLogicalBytes,omitempty"`

	// Output only. Number of logical bytes that are more than 90 days old.
	NumLongTermLogicalBytes *int64 `json:"numLongTermLogicalBytes,omitempty"`

	// Output only. Number of physical bytes used by current live data storage.
	//  This data is not kept in real time, and might be delayed by a few seconds
	//  to a few minutes.
	NumCurrentPhysicalBytes *int64 `json:"numCurrentPhysicalBytes,omitempty"`

	// Output only. The physical size of this table in bytes. This also includes
	//  storage used for time travel. This data is not kept in real time, and might
	//  be delayed by a few seconds to a few minutes.
	NumTotalPhysicalBytes *int64 `json:"numTotalPhysicalBytes,omitempty"`

	// Output only. Number of physical bytes less than 90 days old. This data is
	//  not kept in real time, and might be delayed by a few seconds to a few
	//  minutes.
	NumActivePhysicalBytes *int64 `json:"numActivePhysicalBytes,omitempty"`

	// Output only. Number of physical bytes more than 90 days old.
	//  This data is not kept in real time, and might be delayed by a few seconds
	//  to a few minutes.
	NumLongTermPhysicalBytes *int64 `json:"numLongTermPhysicalBytes,omitempty"`

	// Output only. The number of partitions present in the table or materialized
	//  view. This data is not kept in real time, and might be delayed by a few
	//  seconds to a few minutes.
	NumPartitions *int64 `json:"numPartitions,omitempty"`

	// Optional. The maximum staleness of data that could be returned when the
	//  table (or stale MV) is queried. Staleness encoded as a string encoding
	//  of sql IntervalValue type.
	MaxStaleness *string `json:"maxStaleness,omitempty"`

	// Optional. Output only. Restriction config for table. If set, restrict
	//  certain accesses on the table based on the config. See [Data
	//  egress](https://cloud.google.com/bigquery/docs/analytics-hub-introduction#data_egress)
	//  for more details.
	Restrictions *RestrictionConfig `json:"restrictions,omitempty"`

	// Optional. Tables Primary Key and Foreign Key information
	TableConstraints *TableConstraints `json:"tableConstraints,omitempty"`

	// Optional. The [tags](https://cloud.google.com/bigquery/docs/tags) attached
	//  to this table. Tag keys are globally unique. Tag key is expected to be in
	//  the namespaced format, for example "123456789012/environment" where
	//  123456789012 is the ID of the parent organization or project resource for
	//  this tag key. Tag value is expected to be the short name, for example
	//  "Production". See [Tag
	//  definitions](https://cloud.google.com/iam/docs/tags-access-control#definitions)
	//  for more details.
	ResourceTags map[string]string `json:"resourceTags,omitempty"`

	// Optional. Table replication info for table created `AS REPLICA` DDL like:
	//  `CREATE MATERIALIZED VIEW mv1 AS REPLICA OF src_mv`
	TableReplicationInfo *TableReplicationInfo `json:"tableReplicationInfo,omitempty"`

	// Optional. Output only. Table references of all replicas currently active on
	//  the table.
	Replicas []TableReference `json:"replicas,omitempty"`

	// Optional. Options defining open source compatible table.
	ExternalCatalogTableOptions *ExternalCatalogTableOptions `json:"externalCatalogTableOptions,omitempty"`
}

// +kcc:proto=google.cloud.bigquery.v2.TableConstraints
type TableConstraints struct {
	// Optional. Represents a primary key constraint on a table's columns.
	//  Present only if the table has a primary key.
	//  The primary key is not enforced.
	PrimaryKey *PrimaryKey `json:"primaryKey,omitempty"`

	// Optional. Present only if the table has a foreign key.
	//  The foreign key is not enforced.
	ForeignKeys []ForeignKey `json:"foreignKeys,omitempty"`
}

// +kcc:proto=google.cloud.bigquery.v2.TableFieldSchema
type TableFieldSchema struct {
	// Required. The field name. The name must contain only letters (a-z, A-Z),
	//  numbers (0-9), or underscores (_), and must start with a letter or
	//  underscore. The maximum length is 300 characters.
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
	Type *string `json:"type,omitempty"`

	// Optional. The field mode. Possible values include NULLABLE, REQUIRED and
	//  REPEATED. The default value is NULLABLE.
	Mode *string `json:"mode,omitempty"`

	// Optional. Describes the nested schema fields if the type property is set
	//  to RECORD.
	Fields []TableFieldSchema `json:"fields,omitempty"`

	// Optional. The field description. The maximum length is 1,024 characters.
	Description *string `json:"description,omitempty"`

	// Optional. The policy tags attached to this field, used for field-level
	//  access control. If not set, defaults to empty policy_tags.
	PolicyTags *TableFieldSchema_PolicyTagList `json:"policyTags,omitempty"`

	// Optional. Data policy options, will replace the data_policies.
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
	Precision *int64 `json:"precision,omitempty"`

	// Optional. See documentation for precision.
	Scale *int64 `json:"scale,omitempty"`

	// Optional. Specifies the rounding mode to be used when storing values of
	//  NUMERIC and BIGNUMERIC type.
	RoundingMode *string `json:"roundingMode,omitempty"`

	// Optional. Field collation can be set only when the type of field is STRING.
	//  The following values are supported:
	//
	//  * 'und:ci': undetermined locale, case insensitive.
	//  * '': empty string. Default to case-sensitive behavior.
	Collation *string `json:"collation,omitempty"`

	// Optional. A SQL expression to specify the [default value]
	//  (https://cloud.google.com/bigquery/docs/default-values) for this field.
	DefaultValueExpression *string `json:"defaultValueExpression,omitempty"`

	// Optional. The subtype of the RANGE, if the type of this field is RANGE. If
	//  the type is RANGE, this field is required. Values for the field element
	//  type can be the following:
	//
	//  * DATE
	//  * DATETIME
	//  * TIMESTAMP
	RangeElementType *TableFieldSchema_FieldElementType `json:"rangeElementType,omitempty"`

	// Optional. Definition of the foreign data type.
	//  Only valid for top-level schema fields (not nested fields).
	//  If the type is FOREIGN, this field is required.
	ForeignTypeDefinition *string `json:"foreignTypeDefinition,omitempty"`
}

// +kcc:proto=google.cloud.bigquery.v2.TableFieldSchema.FieldElementType
type TableFieldSchema_FieldElementType struct {
	// Required. The type of a field element. For more information, see
	//  [TableFieldSchema.type][google.cloud.bigquery.v2.TableFieldSchema.type].
	Type *string `json:"type,omitempty"`
}

// +kcc:proto=google.cloud.bigquery.v2.TableFieldSchema.PolicyTagList
type TableFieldSchema_PolicyTagList struct {
	// A list of policy tag resource names. For example,
	//  "projects/1/locations/eu/taxonomies/2/policyTags/3". At most 1 policy tag
	//  is currently allowed.
	Names []string `json:"names,omitempty"`
}

// +kcc:proto=google.cloud.bigquery.v2.TableList
type TableList struct {
	// The type of list.
	Kind *string `json:"kind,omitempty"`

	// A hash of this page of results.
	Etag *string `json:"etag,omitempty"`

	// A token to request the next page of results.
	NextPageToken *string `json:"nextPageToken,omitempty"`

	// Tables in the requested dataset.
	Tables []ListFormatTable `json:"tables,omitempty"`

	// The total number of tables in the dataset.
	TotalItems *int32 `json:"totalItems,omitempty"`
}

// +kcc:proto=google.cloud.bigquery.v2.TableMetadataCacheUsage
type TableMetadataCacheUsage struct {
	// Metadata caching eligible table referenced in the query.
	TableReference *TableReference `json:"tableReference,omitempty"`

	// Reason for not using metadata caching for the table.
	UnusedReason *string `json:"unusedReason,omitempty"`

	// Free form human-readable reason metadata caching was unused for
	//  the job.
	Explanation *string `json:"explanation,omitempty"`

	// Duration since last refresh as of this job for managed tables (indicates
	//  metadata cache staleness as seen by this job).
	Staleness *string `json:"staleness,omitempty"`

	// [Table
	//  type](https://cloud.google.com/bigquery/docs/reference/rest/v2/tables#Table.FIELDS.type).
	TableType *string `json:"tableType,omitempty"`
}

// +kcc:proto=google.cloud.bigquery.v2.TableReference
type TableReference struct {
	// Required. The ID of the project containing this table.
	ProjectId *string `json:"projectId"`

	// Required. The ID of the dataset containing this table.
	DatasetId *string `json:"datasetId"`

	// Required. The ID of the table. The ID can contain Unicode characters in
	//  category L (letter), M (mark), N (number), Pc (connector, including
	//  underscore), Pd (dash), and Zs (space). For more information, see [General
	//  Category](https://wikipedia.org/wiki/Unicode_character_property#General_Category).
	//  The maximum length is 1,024 characters.  Certain operations allow suffixing
	//  of the table ID with a partition decorator, such as
	//  `sample_table$20190123`.
	TableId *string `json:"tableId"`
}

// +kcc:proto=google.cloud.bigquery.v2.TableReplicationInfo
type TableReplicationInfo struct {
	// Required. Source table reference that is replicated.
	SourceTable *TableReference `json:"sourceTable,omitempty"`

	// Optional. Specifies the interval at which the source table is polled for
	//  updates.
	//  It's Optional. If not specified, default replication interval would be
	//  applied.
	ReplicationIntervalMs *int64 `json:"replicationIntervalMs,omitempty"`

	// Optional. Output only. If source is a materialized view, this field
	//  signifies the last refresh time of the source.
	ReplicatedSourceLastRefreshTime *int64 `json:"replicatedSourceLastRefreshTime,omitempty"`

	// Optional. Output only. Replication status of configured replication.
	ReplicationStatus *string `json:"replicationStatus,omitempty"`

	// Optional. Output only. Replication error that will permanently stopped
	//  table replication.
	ReplicationError *ErrorProto `json:"replicationError,omitempty"`
}

// +kcc:proto=google.cloud.bigquery.v2.TableSchema
type TableSchema struct {
	// Describes the fields in a table.
	Fields []TableFieldSchema `json:"fields,omitempty"`

	// Optional. Specifies metadata of the foreign data type definition in field
	//  schema
	//  ([TableFieldSchema.foreign_type_definition][google.cloud.bigquery.v2.TableFieldSchema.foreign_type_definition]).
	ForeignTypeInfo *ForeignTypeInfo `json:"foreignTypeInfo,omitempty"`
}

// +kcc:proto=google.cloud.bigquery.v2.TimePartitioning
type TimePartitioning struct {
	// Required. The supported types are DAY, HOUR, MONTH, and YEAR, which will
	//  generate one partition per day, hour, month, and year, respectively.
	Type *string `json:"type,omitempty"`

	// Optional. Number of milliseconds for which to keep the storage for a
	//  partition.
	//  A wrapper is used here because 0 is an invalid value.
	ExpirationMs *int64 `json:"expirationMs,omitempty"`

	// Optional. If not set, the table is partitioned by pseudo
	//  column '_PARTITIONTIME'; if set, the table is partitioned by this field.
	//  The field must be a top-level TIMESTAMP or DATE field. Its mode must be
	//  NULLABLE or REQUIRED.
	//  A wrapper is used here because an empty string is an invalid value.
	Field *string `json:"field,omitempty"`
}

// +kcc:proto=google.cloud.bigquery.v2.TransformColumn
type TransformColumn struct {
	// Output only. Name of the column.
	Name *string `json:"name,omitempty"`

	// Output only. Data type of the column after the transform.
	Type *StandardSqlDataType `json:"type,omitempty"`

	// Output only. The SQL expression used in the column transform.
	TransformSql *string `json:"transformSql,omitempty"`
}

// +kcc:proto=google.cloud.bigquery.v2.UserDefinedFunctionResource
type UserDefinedFunctionResource struct {
	// [Pick one] A code resource to load from a Google Cloud Storage URI
	//  (gs://bucket/path).
	ResourceUri *string `json:"resourceUri,omitempty"`

	// [Pick one] An inline resource that contains code for a user-defined
	//  function (UDF). Providing a inline code resource is equivalent to providing
	//  a URI for a file containing the same code.
	InlineCode *string `json:"inlineCode,omitempty"`
}

// +kcc:proto=google.cloud.bigquery.v2.VectorSearchStatistics
type VectorSearchStatistics struct {
	// Specifies the index usage mode for the query.
	IndexUsageMode *string `json:"indexUsageMode,omitempty"`

	// When `indexUsageMode` is `UNUSED` or `PARTIALLY_USED`, this field explains
	//  why indexes were not used in all or part of the vector search query. If
	//  `indexUsageMode` is `FULLY_USED`, this field is not populated.
	IndexUnusedReasons []IndexUnusedReason `json:"indexUnusedReasons,omitempty"`
}

// +kcc:proto=google.cloud.bigquery.v2.ViewDefinition
type ViewDefinition struct {
	// Required. A query that BigQuery executes when the view is referenced.
	Query *string `json:"query,omitempty"`

	// Describes user-defined function resources used in the query.
	UserDefinedFunctionResources []UserDefinedFunctionResource `json:"userDefinedFunctionResources,omitempty"`

	// Specifies whether to use BigQuery's legacy SQL for this view.
	//  The default value is true. If set to false, the view will use
	//  BigQuery's GoogleSQL:
	//  https://cloud.google.com/bigquery/sql-reference/
	//
	//  Queries and views that reference this view must use the same flag value.
	//  A wrapper is used here because the default value is True.
	UseLegacySql *bool `json:"useLegacySql,omitempty"`

	// True if the column names are explicitly specified. For example by using the
	//  'CREATE VIEW v(c1, c2) AS ...' syntax.
	//  Can only be set for GoogleSQL views.
	UseExplicitColumnNames *bool `json:"useExplicitColumnNames,omitempty"`

	// Optional. Specifics the privacy policy for the view.
	PrivacyPolicy *PrivacyPolicy `json:"privacyPolicy,omitempty"`

	// Optional. Foreign view representations.
	ForeignDefinitions []ForeignViewDefinition `json:"foreignDefinitions,omitempty"`
}
