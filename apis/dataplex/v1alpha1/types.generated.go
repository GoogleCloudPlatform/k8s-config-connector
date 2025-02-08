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


// +kcc:proto=google.cloud.dataplex.v1.Entity
type Entity struct {

	// Optional. Display name must be shorter than or equal to 256 characters.
	// +kcc:proto:field=google.cloud.dataplex.v1.Entity.display_name
	DisplayName *string `json:"displayName,omitempty"`

	// Optional. User friendly longer description text. Must be shorter than or
	//  equal to 1024 characters.
	// +kcc:proto:field=google.cloud.dataplex.v1.Entity.description
	Description *string `json:"description,omitempty"`

	// Required. A user-provided entity ID. It is mutable, and will be used as the
	//  published table name. Specifying a new ID in an update entity
	//  request will override the existing value.
	//  The ID must contain only letters (a-z, A-Z), numbers (0-9), and
	//  underscores, and consist of 256 or fewer characters.
	// +kcc:proto:field=google.cloud.dataplex.v1.Entity.id
	ID *string `json:"id,omitempty"`

	// Optional. The etag associated with the entity, which can be retrieved with
	//  a [GetEntity][] request. Required for update and delete requests.
	// +kcc:proto:field=google.cloud.dataplex.v1.Entity.etag
	Etag *string `json:"etag,omitempty"`

	// Required. Immutable. The type of entity.
	// +kcc:proto:field=google.cloud.dataplex.v1.Entity.type
	Type *string `json:"type,omitempty"`

	// Required. Immutable. The ID of the asset associated with the storage
	//  location containing the entity data. The entity must be with in the same
	//  zone with the asset.
	// +kcc:proto:field=google.cloud.dataplex.v1.Entity.asset
	Asset *string `json:"asset,omitempty"`

	// Required. Immutable. The storage path of the entity data.
	//  For Cloud Storage data, this is the fully-qualified path to the entity,
	//  such as `gs://bucket/path/to/data`. For BigQuery data, this is the name of
	//  the table resource, such as
	//  `projects/project_id/datasets/dataset_id/tables/table_id`.
	// +kcc:proto:field=google.cloud.dataplex.v1.Entity.data_path
	DataPath *string `json:"dataPath,omitempty"`

	// Optional. The set of items within the data path constituting the data in
	//  the entity, represented as a glob path. Example:
	//  `gs://bucket/path/to/data/**/*.csv`.
	// +kcc:proto:field=google.cloud.dataplex.v1.Entity.data_path_pattern
	DataPathPattern *string `json:"dataPathPattern,omitempty"`

	// Required. Immutable. Identifies the storage system of the entity data.
	// +kcc:proto:field=google.cloud.dataplex.v1.Entity.system
	System *string `json:"system,omitempty"`

	// Required. Identifies the storage format of the entity data.
	//  It does not apply to entities with data stored in BigQuery.
	// +kcc:proto:field=google.cloud.dataplex.v1.Entity.format
	Format *StorageFormat `json:"format,omitempty"`

	// Required. The description of the data structure and layout.
	//  The schema is not included in list responses. It is only included in
	//  `SCHEMA` and `FULL` entity views of a `GetEntity` response.
	// +kcc:proto:field=google.cloud.dataplex.v1.Entity.schema
	Schema *Schema `json:"schema,omitempty"`
}

// +kcc:proto=google.cloud.dataplex.v1.Entity.CompatibilityStatus
type Entity_CompatibilityStatus struct {
}

// +kcc:proto=google.cloud.dataplex.v1.Entity.CompatibilityStatus.Compatibility
type Entity_CompatibilityStatus_Compatibility struct {
}

// +kcc:proto=google.cloud.dataplex.v1.Schema
type Schema struct {
	// Required. Set to `true` if user-managed or `false` if managed by Dataplex.
	//  The default is `false` (managed by Dataplex).
	//
	//  - Set to `false`to enable Dataplex discovery to update the schema.
	//    including new data discovery, schema inference, and schema evolution.
	//    Users retain the ability to input and edit the schema. Dataplex
	//    treats schema input by the user as though produced
	//    by a previous Dataplex discovery operation, and it will
	//    evolve the schema and take action based on that treatment.
	//
	//  - Set to `true` to fully manage the entity
	//    schema. This setting guarantees that Dataplex will not
	//    change schema fields.
	// +kcc:proto:field=google.cloud.dataplex.v1.Schema.user_managed
	UserManaged *bool `json:"userManaged,omitempty"`

	// Optional. The sequence of fields describing data in table entities.
	//  **Note:** BigQuery SchemaFields are immutable.
	// +kcc:proto:field=google.cloud.dataplex.v1.Schema.fields
	Fields []Schema_SchemaField `json:"fields,omitempty"`

	// Optional. The sequence of fields describing the partition structure in
	//  entities. If this field is empty, there are no partitions within the data.
	// +kcc:proto:field=google.cloud.dataplex.v1.Schema.partition_fields
	PartitionFields []Schema_PartitionField `json:"partitionFields,omitempty"`

	// Optional. The structure of paths containing partition data within the
	//  entity.
	// +kcc:proto:field=google.cloud.dataplex.v1.Schema.partition_style
	PartitionStyle *string `json:"partitionStyle,omitempty"`
}

// +kcc:proto=google.cloud.dataplex.v1.Schema.PartitionField
type Schema_PartitionField struct {
	// Required. Partition field name must consist of letters, numbers, and
	//  underscores only, with a maximum of length of 256 characters, and must
	//  begin with a letter or underscore..
	// +kcc:proto:field=google.cloud.dataplex.v1.Schema.PartitionField.name
	Name *string `json:"name,omitempty"`

	// Required. Immutable. The type of field.
	// +kcc:proto:field=google.cloud.dataplex.v1.Schema.PartitionField.type
	Type *string `json:"type,omitempty"`
}

// +kcc:proto=google.cloud.dataplex.v1.Schema.SchemaField
type Schema_SchemaField struct {
	// Required. The name of the field. Must contain only letters, numbers and
	//  underscores, with a maximum length of 767 characters,
	//  and must begin with a letter or underscore.
	// +kcc:proto:field=google.cloud.dataplex.v1.Schema.SchemaField.name
	Name *string `json:"name,omitempty"`

	// Optional. User friendly field description. Must be less than or equal to
	//  1024 characters.
	// +kcc:proto:field=google.cloud.dataplex.v1.Schema.SchemaField.description
	Description *string `json:"description,omitempty"`

	// Required. The type of field.
	// +kcc:proto:field=google.cloud.dataplex.v1.Schema.SchemaField.type
	Type *string `json:"type,omitempty"`

	// Required. Additional field semantics.
	// +kcc:proto:field=google.cloud.dataplex.v1.Schema.SchemaField.mode
	Mode *string `json:"mode,omitempty"`

	// Optional. Any nested field for complex types.
	// +kcc:proto:field=google.cloud.dataplex.v1.Schema.SchemaField.fields
	Fields []Schema_SchemaField `json:"fields,omitempty"`
}

// +kcc:proto=google.cloud.dataplex.v1.StorageAccess
type StorageAccess struct {
}

// +kcc:proto=google.cloud.dataplex.v1.StorageFormat
type StorageFormat struct {

	// Optional. The compression type associated with the stored data.
	//  If unspecified, the data is uncompressed.
	// +kcc:proto:field=google.cloud.dataplex.v1.StorageFormat.compression_format
	CompressionFormat *string `json:"compressionFormat,omitempty"`

	// Required. The mime type descriptor for the data. Must match the pattern
	//  {type}/{subtype}. Supported values:
	//
	//  - application/x-parquet
	//  - application/x-avro
	//  - application/x-orc
	//  - application/x-tfrecord
	//  - application/x-parquet+iceberg
	//  - application/x-avro+iceberg
	//  - application/x-orc+iceberg
	//  - application/json
	//  - application/{subtypes}
	//  - text/csv
	//  - text/<subtypes>
	//  - image/{image subtype}
	//  - video/{video subtype}
	//  - audio/{audio subtype}
	// +kcc:proto:field=google.cloud.dataplex.v1.StorageFormat.mime_type
	MimeType *string `json:"mimeType,omitempty"`

	// Optional. Additional information about CSV formatted data.
	// +kcc:proto:field=google.cloud.dataplex.v1.StorageFormat.csv
	Csv *StorageFormat_CsvOptions `json:"csv,omitempty"`

	// Optional. Additional information about CSV formatted data.
	// +kcc:proto:field=google.cloud.dataplex.v1.StorageFormat.json
	Json *StorageFormat_JsonOptions `json:"json,omitempty"`

	// Optional. Additional information about iceberg tables.
	// +kcc:proto:field=google.cloud.dataplex.v1.StorageFormat.iceberg
	Iceberg *StorageFormat_IcebergOptions `json:"iceberg,omitempty"`
}

// +kcc:proto=google.cloud.dataplex.v1.StorageFormat.CsvOptions
type StorageFormat_CsvOptions struct {
	// Optional. The character encoding of the data. Accepts "US-ASCII",
	//  "UTF-8", and "ISO-8859-1". Defaults to UTF-8 if unspecified.
	// +kcc:proto:field=google.cloud.dataplex.v1.StorageFormat.CsvOptions.encoding
	Encoding *string `json:"encoding,omitempty"`

	// Optional. The number of rows to interpret as header rows that should be
	//  skipped when reading data rows. Defaults to 0.
	// +kcc:proto:field=google.cloud.dataplex.v1.StorageFormat.CsvOptions.header_rows
	HeaderRows *int32 `json:"headerRows,omitempty"`

	// Optional. The delimiter used to separate values. Defaults to ','.
	// +kcc:proto:field=google.cloud.dataplex.v1.StorageFormat.CsvOptions.delimiter
	Delimiter *string `json:"delimiter,omitempty"`

	// Optional. The character used to quote column values. Accepts '"'
	//  (double quotation mark) or ''' (single quotation mark). Defaults to
	//  '"' (double quotation mark) if unspecified.
	// +kcc:proto:field=google.cloud.dataplex.v1.StorageFormat.CsvOptions.quote
	Quote *string `json:"quote,omitempty"`
}

// +kcc:proto=google.cloud.dataplex.v1.StorageFormat.IcebergOptions
type StorageFormat_IcebergOptions struct {
	// Optional. The location of where the iceberg metadata is present, must be
	//  within the table path
	// +kcc:proto:field=google.cloud.dataplex.v1.StorageFormat.IcebergOptions.metadata_location
	MetadataLocation *string `json:"metadataLocation,omitempty"`
}

// +kcc:proto=google.cloud.dataplex.v1.StorageFormat.JsonOptions
type StorageFormat_JsonOptions struct {
	// Optional. The character encoding of the data. Accepts "US-ASCII", "UTF-8"
	//  and "ISO-8859-1". Defaults to UTF-8 if not specified.
	// +kcc:proto:field=google.cloud.dataplex.v1.StorageFormat.JsonOptions.encoding
	Encoding *string `json:"encoding,omitempty"`
}

// +kcc:proto=google.cloud.dataplex.v1.Entity
type EntityObservedState struct {
	// Output only. The resource name of the entity, of the form:
	//  `projects/{project_number}/locations/{location_id}/lakes/{lake_id}/zones/{zone_id}/entities/{id}`.
	// +kcc:proto:field=google.cloud.dataplex.v1.Entity.name
	Name *string `json:"name,omitempty"`

	// Output only. The time when the entity was created.
	// +kcc:proto:field=google.cloud.dataplex.v1.Entity.create_time
	CreateTime *string `json:"createTime,omitempty"`

	// Output only. The time when the entity was last updated.
	// +kcc:proto:field=google.cloud.dataplex.v1.Entity.update_time
	UpdateTime *string `json:"updateTime,omitempty"`

	// Output only. The name of the associated Data Catalog entry.
	// +kcc:proto:field=google.cloud.dataplex.v1.Entity.catalog_entry
	CatalogEntry *string `json:"catalogEntry,omitempty"`

	// Required. Identifies the storage format of the entity data.
	//  It does not apply to entities with data stored in BigQuery.
	// +kcc:proto:field=google.cloud.dataplex.v1.Entity.format
	Format *StorageFormatObservedState `json:"format,omitempty"`

	// Output only. Metadata stores that the entity is compatible with.
	// +kcc:proto:field=google.cloud.dataplex.v1.Entity.compatibility
	Compatibility *Entity_CompatibilityStatus `json:"compatibility,omitempty"`

	// Output only. Identifies the access mechanism to the entity. Not user
	//  settable.
	// +kcc:proto:field=google.cloud.dataplex.v1.Entity.access
	Access *StorageAccess `json:"access,omitempty"`

	// Output only. System generated unique ID for the Entity. This ID will be
	//  different if the Entity is deleted and re-created with the same name.
	// +kcc:proto:field=google.cloud.dataplex.v1.Entity.uid
	Uid *string `json:"uid,omitempty"`
}

// +kcc:proto=google.cloud.dataplex.v1.Entity.CompatibilityStatus
type Entity_CompatibilityStatusObservedState struct {
	// Output only. Whether this entity is compatible with Hive Metastore.
	// +kcc:proto:field=google.cloud.dataplex.v1.Entity.CompatibilityStatus.hive_metastore
	HiveMetastore *Entity_CompatibilityStatus_Compatibility `json:"hiveMetastore,omitempty"`

	// Output only. Whether this entity is compatible with BigQuery.
	// +kcc:proto:field=google.cloud.dataplex.v1.Entity.CompatibilityStatus.bigquery
	Bigquery *Entity_CompatibilityStatus_Compatibility `json:"bigquery,omitempty"`
}

// +kcc:proto=google.cloud.dataplex.v1.Entity.CompatibilityStatus.Compatibility
type Entity_CompatibilityStatus_CompatibilityObservedState struct {
	// Output only. Whether the entity is compatible and can be represented in
	//  the metadata store.
	// +kcc:proto:field=google.cloud.dataplex.v1.Entity.CompatibilityStatus.Compatibility.compatible
	Compatible *bool `json:"compatible,omitempty"`

	// Output only. Provides additional detail if the entity is incompatible
	//  with the metadata store.
	// +kcc:proto:field=google.cloud.dataplex.v1.Entity.CompatibilityStatus.Compatibility.reason
	Reason *string `json:"reason,omitempty"`
}

// +kcc:proto=google.cloud.dataplex.v1.StorageAccess
type StorageAccessObservedState struct {
	// Output only. Describes the read access mechanism of the data. Not user
	//  settable.
	// +kcc:proto:field=google.cloud.dataplex.v1.StorageAccess.read
	Read *string `json:"read,omitempty"`
}

// +kcc:proto=google.cloud.dataplex.v1.StorageFormat
type StorageFormatObservedState struct {
	// Output only. The data format associated with the stored data, which
	//  represents content type values. The value is inferred from mime type.
	// +kcc:proto:field=google.cloud.dataplex.v1.StorageFormat.format
	Format *string `json:"format,omitempty"`
}
