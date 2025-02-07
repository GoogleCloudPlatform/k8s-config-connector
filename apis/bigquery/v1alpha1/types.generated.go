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


// +kcc:proto=google.cloud.bigquery.storage.v1.TableFieldSchema
type TableFieldSchema struct {
	// Required. The field name. The name must contain only letters (a-z, A-Z),
	//  numbers (0-9), or underscores (_), and must start with a letter or
	//  underscore. The maximum length is 128 characters.
	// +kcc:proto:field=google.cloud.bigquery.storage.v1.TableFieldSchema.name
	Name *string `json:"name,omitempty"`

	// Required. The field data type.
	// +kcc:proto:field=google.cloud.bigquery.storage.v1.TableFieldSchema.type
	Type *string `json:"type,omitempty"`

	// Optional. The field mode. The default value is NULLABLE.
	// +kcc:proto:field=google.cloud.bigquery.storage.v1.TableFieldSchema.mode
	Mode *string `json:"mode,omitempty"`

	// Optional. Describes the nested schema fields if the type property is set to
	//  STRUCT.
	// +kcc:proto:field=google.cloud.bigquery.storage.v1.TableFieldSchema.fields
	Fields []TableFieldSchema `json:"fields,omitempty"`

	// Optional. The field description. The maximum length is 1,024 characters.
	// +kcc:proto:field=google.cloud.bigquery.storage.v1.TableFieldSchema.description
	Description *string `json:"description,omitempty"`

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
	//  It is invalid to set this field if type is not "STRING" or "BYTES".
	// +kcc:proto:field=google.cloud.bigquery.storage.v1.TableFieldSchema.max_length
	MaxLength *int64 `json:"maxLength,omitempty"`

	// Optional. Precision (maximum number of total digits in base 10) and scale
	//  (maximum number of digits in the fractional part in base 10) constraints
	//  for values of this field for NUMERIC or BIGNUMERIC.
	//
	//  It is invalid to set precision or scale if type is not "NUMERIC" or
	//  "BIGNUMERIC".
	//
	//  If precision and scale are not specified, no value range constraint is
	//  imposed on this field insofar as values are permitted by the type.
	//
	//  Values of this NUMERIC or BIGNUMERIC field must be in this range when:
	//
	//  * Precision (P) and scale (S) are specified:
	//    [-10^(P-S) + 10^(-S), 10^(P-S) - 10^(-S)]
	//  * Precision (P) is specified but not scale (and thus scale is
	//    interpreted to be equal to zero):
	//    [-10^P + 1, 10^P - 1].
	//
	//  Acceptable values for precision and scale if both are specified:
	//
	//  * If type = "NUMERIC":
	//    1 <= precision - scale <= 29 and 0 <= scale <= 9.
	//  * If type = "BIGNUMERIC":
	//    1 <= precision - scale <= 38 and 0 <= scale <= 38.
	//
	//  Acceptable values for precision if only precision is specified but not
	//  scale (and thus scale is interpreted to be equal to zero):
	//
	//  * If type = "NUMERIC": 1 <= precision <= 29.
	//  * If type = "BIGNUMERIC": 1 <= precision <= 38.
	//
	//  If scale is specified but not precision, then it is invalid.
	// +kcc:proto:field=google.cloud.bigquery.storage.v1.TableFieldSchema.precision
	Precision *int64 `json:"precision,omitempty"`

	// Optional. See documentation for precision.
	// +kcc:proto:field=google.cloud.bigquery.storage.v1.TableFieldSchema.scale
	Scale *int64 `json:"scale,omitempty"`

	// Optional. A SQL expression to specify the [default value]
	//  (https://cloud.google.com/bigquery/docs/default-values) for this field.
	// +kcc:proto:field=google.cloud.bigquery.storage.v1.TableFieldSchema.default_value_expression
	DefaultValueExpression *string `json:"defaultValueExpression,omitempty"`

	// Optional. The subtype of the RANGE, if the type of this field is RANGE. If
	//  the type is RANGE, this field is required. Possible values for the field
	//  element type of a RANGE include:
	//  * DATE
	//  * DATETIME
	//  * TIMESTAMP
	// +kcc:proto:field=google.cloud.bigquery.storage.v1.TableFieldSchema.range_element_type
	RangeElementType *TableFieldSchema_FieldElementType `json:"rangeElementType,omitempty"`
}

// +kcc:proto=google.cloud.bigquery.storage.v1.TableFieldSchema.FieldElementType
type TableFieldSchema_FieldElementType struct {
	// Required. The type of a field element.
	// +kcc:proto:field=google.cloud.bigquery.storage.v1.TableFieldSchema.FieldElementType.type
	Type *string `json:"type,omitempty"`
}

// +kcc:proto=google.cloud.bigquery.storage.v1.TableSchema
type TableSchema struct {
	// Describes the fields in a table.
	// +kcc:proto:field=google.cloud.bigquery.storage.v1.TableSchema.fields
	Fields []TableFieldSchema `json:"fields,omitempty"`
}

// +kcc:proto=google.cloud.bigquery.storage.v1.WriteStream
type WriteStream struct {

	// Immutable. Type of the stream.
	// +kcc:proto:field=google.cloud.bigquery.storage.v1.WriteStream.type
	Type *string `json:"type,omitempty"`

	// Immutable. Mode of the stream.
	// +kcc:proto:field=google.cloud.bigquery.storage.v1.WriteStream.write_mode
	WriteMode *string `json:"writeMode,omitempty"`

	// Immutable. The geographic location where the stream's dataset resides. See
	//  https://cloud.google.com/bigquery/docs/locations for supported
	//  locations.
	// +kcc:proto:field=google.cloud.bigquery.storage.v1.WriteStream.location
	Location *string `json:"location,omitempty"`
}

// +kcc:proto=google.cloud.bigquery.storage.v1.WriteStream
type WriteStreamObservedState struct {
	// Output only. Name of the stream, in the form
	//  `projects/{project}/datasets/{dataset}/tables/{table}/streams/{stream}`.
	// +kcc:proto:field=google.cloud.bigquery.storage.v1.WriteStream.name
	Name *string `json:"name,omitempty"`

	// Output only. Create time of the stream. For the _default stream, this is
	//  the creation_time of the table.
	// +kcc:proto:field=google.cloud.bigquery.storage.v1.WriteStream.create_time
	CreateTime *string `json:"createTime,omitempty"`

	// Output only. Commit time of the stream.
	//  If a stream is of `COMMITTED` type, then it will have a commit_time same as
	//  `create_time`. If the stream is of `PENDING` type, empty commit_time
	//  means it is not committed.
	// +kcc:proto:field=google.cloud.bigquery.storage.v1.WriteStream.commit_time
	CommitTime *string `json:"commitTime,omitempty"`

	// Output only. The schema of the destination table. It is only returned in
	//  `CreateWriteStream` response. Caller should generate data that's
	//  compatible with this schema to send in initial `AppendRowsRequest`.
	//  The table schema could go out of date during the life time of the stream.
	// +kcc:proto:field=google.cloud.bigquery.storage.v1.WriteStream.table_schema
	TableSchema *TableSchema `json:"tableSchema,omitempty"`
}
