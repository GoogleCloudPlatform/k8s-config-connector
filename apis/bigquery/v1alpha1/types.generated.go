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


// +kcc:proto=google.cloud.bigquery.storage.v1beta2.ArrowSchema
type ArrowSchema struct {
	// IPC serialized Arrow schema.
	// +kcc:proto:field=google.cloud.bigquery.storage.v1beta2.ArrowSchema.serialized_schema
	SerializedSchema []byte `json:"serializedSchema,omitempty"`
}

// +kcc:proto=google.cloud.bigquery.storage.v1beta2.ArrowSerializationOptions
type ArrowSerializationOptions struct {
	// The Arrow IPC format to use.
	// +kcc:proto:field=google.cloud.bigquery.storage.v1beta2.ArrowSerializationOptions.format
	Format *string `json:"format,omitempty"`
}

// +kcc:proto=google.cloud.bigquery.storage.v1beta2.AvroSchema
type AvroSchema struct {
	// Json serialized schema, as described at
	//  https://avro.apache.org/docs/1.8.1/spec.html.
	// +kcc:proto:field=google.cloud.bigquery.storage.v1beta2.AvroSchema.schema
	Schema *string `json:"schema,omitempty"`
}

// +kcc:proto=google.cloud.bigquery.storage.v1beta2.ReadSession
type ReadSession struct {

	// Immutable. Data format of the output data.
	// +kcc:proto:field=google.cloud.bigquery.storage.v1beta2.ReadSession.data_format
	DataFormat *string `json:"dataFormat,omitempty"`

	// Immutable. Table that this ReadSession is reading from, in the form
	//  `projects/{project_id}/datasets/{dataset_id}/tables/{table_id}
	// +kcc:proto:field=google.cloud.bigquery.storage.v1beta2.ReadSession.table
	Table *string `json:"table,omitempty"`

	// Optional. Any modifiers which are applied when reading from the specified table.
	// +kcc:proto:field=google.cloud.bigquery.storage.v1beta2.ReadSession.table_modifiers
	TableModifiers *ReadSession_TableModifiers `json:"tableModifiers,omitempty"`

	// Optional. Read options for this session (e.g. column selection, filters).
	// +kcc:proto:field=google.cloud.bigquery.storage.v1beta2.ReadSession.read_options
	ReadOptions *ReadSession_TableReadOptions `json:"readOptions,omitempty"`
}

// +kcc:proto=google.cloud.bigquery.storage.v1beta2.ReadSession.TableModifiers
type ReadSession_TableModifiers struct {
	// The snapshot time of the table. If not set, interpreted as now.
	// +kcc:proto:field=google.cloud.bigquery.storage.v1beta2.ReadSession.TableModifiers.snapshot_time
	SnapshotTime *string `json:"snapshotTime,omitempty"`
}

// +kcc:proto=google.cloud.bigquery.storage.v1beta2.ReadSession.TableReadOptions
type ReadSession_TableReadOptions struct {
	// Names of the fields in the table that should be read. If empty, all
	//  fields will be read. If the specified field is a nested field, all
	//  the sub-fields in the field will be selected. The output field order is
	//  unrelated to the order of fields in selected_fields.
	// +kcc:proto:field=google.cloud.bigquery.storage.v1beta2.ReadSession.TableReadOptions.selected_fields
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
	// +kcc:proto:field=google.cloud.bigquery.storage.v1beta2.ReadSession.TableReadOptions.row_restriction
	RowRestriction *string `json:"rowRestriction,omitempty"`

	// Optional. Options specific to the Apache Arrow output format.
	// +kcc:proto:field=google.cloud.bigquery.storage.v1beta2.ReadSession.TableReadOptions.arrow_serialization_options
	ArrowSerializationOptions *ArrowSerializationOptions `json:"arrowSerializationOptions,omitempty"`
}

// +kcc:proto=google.cloud.bigquery.storage.v1beta2.ReadStream
type ReadStream struct {
}

// +kcc:proto=google.cloud.bigquery.storage.v1beta2.ReadSession
type ReadSessionObservedState struct {
	// Output only. Unique identifier for the session, in the form
	//  `projects/{project_id}/locations/{location}/sessions/{session_id}`.
	// +kcc:proto:field=google.cloud.bigquery.storage.v1beta2.ReadSession.name
	Name *string `json:"name,omitempty"`

	// Output only. Time at which the session becomes invalid. After this time, subsequent
	//  requests to read this Session will return errors. The expire_time is
	//  automatically assigned and currently cannot be specified or updated.
	// +kcc:proto:field=google.cloud.bigquery.storage.v1beta2.ReadSession.expire_time
	ExpireTime *string `json:"expireTime,omitempty"`

	// Output only. Avro schema.
	// +kcc:proto:field=google.cloud.bigquery.storage.v1beta2.ReadSession.avro_schema
	AvroSchema *AvroSchema `json:"avroSchema,omitempty"`

	// Output only. Arrow schema.
	// +kcc:proto:field=google.cloud.bigquery.storage.v1beta2.ReadSession.arrow_schema
	ArrowSchema *ArrowSchema `json:"arrowSchema,omitempty"`

	// Output only. A list of streams created with the session.
	//
	//  At least one stream is created with the session. In the future, larger
	//  request_stream_count values *may* result in this list being unpopulated,
	//  in that case, the user will need to use a List method to get the streams
	//  instead, which is not yet available.
	// +kcc:proto:field=google.cloud.bigquery.storage.v1beta2.ReadSession.streams
	Streams []ReadStream `json:"streams,omitempty"`
}

// +kcc:proto=google.cloud.bigquery.storage.v1beta2.ReadStream
type ReadStreamObservedState struct {
	// Output only. Name of the stream, in the form
	//  `projects/{project_id}/locations/{location}/sessions/{session_id}/streams/{stream_id}`.
	// +kcc:proto:field=google.cloud.bigquery.storage.v1beta2.ReadStream.name
	Name *string `json:"name,omitempty"`
}
