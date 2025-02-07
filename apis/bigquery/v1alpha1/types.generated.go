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


// +kcc:proto=google.cloud.bigquery.storage.v1.ArrowSchema
type ArrowSchema struct {
	// IPC serialized Arrow schema.
	// +kcc:proto:field=google.cloud.bigquery.storage.v1.ArrowSchema.serialized_schema
	SerializedSchema []byte `json:"serializedSchema,omitempty"`
}

// +kcc:proto=google.cloud.bigquery.storage.v1.ArrowSerializationOptions
type ArrowSerializationOptions struct {
	// The compression codec to use for Arrow buffers in serialized record
	//  batches.
	// +kcc:proto:field=google.cloud.bigquery.storage.v1.ArrowSerializationOptions.buffer_compression
	BufferCompression *string `json:"bufferCompression,omitempty"`
}

// +kcc:proto=google.cloud.bigquery.storage.v1.AvroSchema
type AvroSchema struct {
	// Json serialized schema, as described at
	//  https://avro.apache.org/docs/1.8.1/spec.html.
	// +kcc:proto:field=google.cloud.bigquery.storage.v1.AvroSchema.schema
	Schema *string `json:"schema,omitempty"`
}

// +kcc:proto=google.cloud.bigquery.storage.v1.AvroSerializationOptions
type AvroSerializationOptions struct {
	// Enable displayName attribute in Avro schema.
	//
	//  The Avro specification requires field names to be alphanumeric.  By
	//  default, in cases when column names do not conform to these requirements
	//  (e.g. non-ascii unicode codepoints) and Avro is requested as an output
	//  format, the CreateReadSession call will fail.
	//
	//  Setting this field to true, populates avro field names with a placeholder
	//  value and populates a "displayName" attribute for every avro field with the
	//  original column name.
	// +kcc:proto:field=google.cloud.bigquery.storage.v1.AvroSerializationOptions.enable_display_name_attribute
	EnableDisplayNameAttribute *bool `json:"enableDisplayNameAttribute,omitempty"`
}

// +kcc:proto=google.cloud.bigquery.storage.v1.ReadSession
type ReadSession struct {

	// Immutable. Data format of the output data. DATA_FORMAT_UNSPECIFIED not
	//  supported.
	// +kcc:proto:field=google.cloud.bigquery.storage.v1.ReadSession.data_format
	DataFormat *string `json:"dataFormat,omitempty"`

	// Immutable. Table that this ReadSession is reading from, in the form
	//  `projects/{project_id}/datasets/{dataset_id}/tables/{table_id}`
	// +kcc:proto:field=google.cloud.bigquery.storage.v1.ReadSession.table
	Table *string `json:"table,omitempty"`

	// Optional. Any modifiers which are applied when reading from the specified
	//  table.
	// +kcc:proto:field=google.cloud.bigquery.storage.v1.ReadSession.table_modifiers
	TableModifiers *ReadSession_TableModifiers `json:"tableModifiers,omitempty"`

	// Optional. Read options for this session (e.g. column selection, filters).
	// +kcc:proto:field=google.cloud.bigquery.storage.v1.ReadSession.read_options
	ReadOptions *ReadSession_TableReadOptions `json:"readOptions,omitempty"`

	// Optional. ID set by client to annotate a session identity.  This does not
	//  need to be strictly unique, but instead the same ID should be used to group
	//  logically connected sessions (e.g. All using the same ID for all sessions
	//  needed to complete a Spark SQL query is reasonable).
	//
	//  Maximum length is 256 bytes.
	// +kcc:proto:field=google.cloud.bigquery.storage.v1.ReadSession.trace_id
	TraceID *string `json:"traceID,omitempty"`
}

// +kcc:proto=google.cloud.bigquery.storage.v1.ReadSession.TableModifiers
type ReadSession_TableModifiers struct {
	// The snapshot time of the table. If not set, interpreted as now.
	// +kcc:proto:field=google.cloud.bigquery.storage.v1.ReadSession.TableModifiers.snapshot_time
	SnapshotTime *string `json:"snapshotTime,omitempty"`
}

// +kcc:proto=google.cloud.bigquery.storage.v1.ReadSession.TableReadOptions
type ReadSession_TableReadOptions struct {
	// Optional. The names of the fields in the table to be returned. If no
	//  field names are specified, then all fields in the table are returned.
	//
	//  Nested fields -- the child elements of a STRUCT field -- can be selected
	//  individually using their fully-qualified names, and will be returned as
	//  record fields containing only the selected nested fields. If a STRUCT
	//  field is specified in the selected fields list, all of the child elements
	//  will be returned.
	//
	//  As an example, consider a table with the following schema:
	//
	//    {
	//        "name": "struct_field",
	//        "type": "RECORD",
	//        "mode": "NULLABLE",
	//        "fields": [
	//            {
	//                "name": "string_field1",
	//                "type": "STRING",
	//  .              "mode": "NULLABLE"
	//            },
	//            {
	//                "name": "string_field2",
	//                "type": "STRING",
	//                "mode": "NULLABLE"
	//            }
	//        ]
	//    }
	//
	//  Specifying "struct_field" in the selected fields list will result in a
	//  read session schema with the following logical structure:
	//
	//    struct_field {
	//        string_field1
	//        string_field2
	//    }
	//
	//  Specifying "struct_field.string_field1" in the selected fields list will
	//  result in a read session schema with the following logical structure:
	//
	//    struct_field {
	//        string_field1
	//    }
	//
	//  The order of the fields in the read session schema is derived from the
	//  table schema and does not correspond to the order in which the fields are
	//  specified in this list.
	// +kcc:proto:field=google.cloud.bigquery.storage.v1.ReadSession.TableReadOptions.selected_fields
	SelectedFields []string `json:"selectedFields,omitempty"`

	// SQL text filtering statement, similar to a WHERE clause in a query.
	//  Aggregates are not supported.
	//
	//  Examples: "int_field > 5"
	//            "date_field = CAST('2014-9-27' as DATE)"
	//            "nullable_field is not NULL"
	//            "st_equals(geo_field, st_geofromtext("POINT(2, 2)"))"
	//            "numeric_field BETWEEN 1.0 AND 5.0"
	//
	//  Restricted to a maximum length for 1 MB.
	// +kcc:proto:field=google.cloud.bigquery.storage.v1.ReadSession.TableReadOptions.row_restriction
	RowRestriction *string `json:"rowRestriction,omitempty"`

	// Optional. Options specific to the Apache Arrow output format.
	// +kcc:proto:field=google.cloud.bigquery.storage.v1.ReadSession.TableReadOptions.arrow_serialization_options
	ArrowSerializationOptions *ArrowSerializationOptions `json:"arrowSerializationOptions,omitempty"`

	// Optional. Options specific to the Apache Avro output format
	// +kcc:proto:field=google.cloud.bigquery.storage.v1.ReadSession.TableReadOptions.avro_serialization_options
	AvroSerializationOptions *AvroSerializationOptions `json:"avroSerializationOptions,omitempty"`

	// Optional. Specifies a table sampling percentage. Specifically, the query
	//  planner will use TABLESAMPLE SYSTEM (sample_percentage PERCENT). The
	//  sampling percentage is applied at the data block granularity. It will
	//  randomly choose for each data block whether to read the rows in that data
	//  block. For more details, see
	//  https://cloud.google.com/bigquery/docs/table-sampling)
	// +kcc:proto:field=google.cloud.bigquery.storage.v1.ReadSession.TableReadOptions.sample_percentage
	SamplePercentage *float64 `json:"samplePercentage,omitempty"`

	// Optional. Set response_compression_codec when creating a read session to
	//  enable application-level compression of ReadRows responses.
	// +kcc:proto:field=google.cloud.bigquery.storage.v1.ReadSession.TableReadOptions.response_compression_codec
	ResponseCompressionCodec *string `json:"responseCompressionCodec,omitempty"`
}

// +kcc:proto=google.cloud.bigquery.storage.v1.ReadStream
type ReadStream struct {
}

// +kcc:proto=google.cloud.bigquery.storage.v1.ReadSession
type ReadSessionObservedState struct {
	// Output only. Unique identifier for the session, in the form
	//  `projects/{project_id}/locations/{location}/sessions/{session_id}`.
	// +kcc:proto:field=google.cloud.bigquery.storage.v1.ReadSession.name
	Name *string `json:"name,omitempty"`

	// Output only. Time at which the session becomes invalid. After this time,
	//  subsequent requests to read this Session will return errors. The
	//  expire_time is automatically assigned and currently cannot be specified or
	//  updated.
	// +kcc:proto:field=google.cloud.bigquery.storage.v1.ReadSession.expire_time
	ExpireTime *string `json:"expireTime,omitempty"`

	// Output only. Avro schema.
	// +kcc:proto:field=google.cloud.bigquery.storage.v1.ReadSession.avro_schema
	AvroSchema *AvroSchema `json:"avroSchema,omitempty"`

	// Output only. Arrow schema.
	// +kcc:proto:field=google.cloud.bigquery.storage.v1.ReadSession.arrow_schema
	ArrowSchema *ArrowSchema `json:"arrowSchema,omitempty"`

	// Output only. A list of streams created with the session.
	//
	//  At least one stream is created with the session. In the future, larger
	//  request_stream_count values *may* result in this list being unpopulated,
	//  in that case, the user will need to use a List method to get the streams
	//  instead, which is not yet available.
	// +kcc:proto:field=google.cloud.bigquery.storage.v1.ReadSession.streams
	Streams []ReadStream `json:"streams,omitempty"`

	// Output only. An estimate on the number of bytes this session will scan when
	//  all streams are completely consumed. This estimate is based on
	//  metadata from the table which might be incomplete or stale.
	// +kcc:proto:field=google.cloud.bigquery.storage.v1.ReadSession.estimated_total_bytes_scanned
	EstimatedTotalBytesScanned *int64 `json:"estimatedTotalBytesScanned,omitempty"`

	// Output only. A pre-projected estimate of the total physical size of files
	//  (in bytes) that this session will scan when all streams are consumed. This
	//  estimate is independent of the selected columns and can be based on
	//  incomplete or stale metadata from the table.  This field is only set for
	//  BigLake tables.
	// +kcc:proto:field=google.cloud.bigquery.storage.v1.ReadSession.estimated_total_physical_file_size
	EstimatedTotalPhysicalFileSize *int64 `json:"estimatedTotalPhysicalFileSize,omitempty"`

	// Output only. An estimate on the number of rows present in this session's
	//  streams. This estimate is based on metadata from the table which might be
	//  incomplete or stale.
	// +kcc:proto:field=google.cloud.bigquery.storage.v1.ReadSession.estimated_row_count
	EstimatedRowCount *int64 `json:"estimatedRowCount,omitempty"`
}

// +kcc:proto=google.cloud.bigquery.storage.v1.ReadStream
type ReadStreamObservedState struct {
	// Output only. Name of the stream, in the form
	//  `projects/{project_id}/locations/{location}/sessions/{session_id}/streams/{stream_id}`.
	// +kcc:proto:field=google.cloud.bigquery.storage.v1.ReadStream.name
	Name *string `json:"name,omitempty"`
}
