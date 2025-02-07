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


// +kcc:proto=google.cloud.bigquery.storage.v1beta1.ArrowSchema
type ArrowSchema struct {
	// IPC serialized Arrow schema.
	// +kcc:proto:field=google.cloud.bigquery.storage.v1beta1.ArrowSchema.serialized_schema
	SerializedSchema []byte `json:"serializedSchema,omitempty"`
}

// +kcc:proto=google.cloud.bigquery.storage.v1beta1.AvroSchema
type AvroSchema struct {
	// Json serialized schema, as described at
	//  https://avro.apache.org/docs/1.8.1/spec.html
	// +kcc:proto:field=google.cloud.bigquery.storage.v1beta1.AvroSchema.schema
	Schema *string `json:"schema,omitempty"`
}

// +kcc:proto=google.cloud.bigquery.storage.v1beta1.ReadSession
type ReadSession struct {
	// Unique identifier for the session, in the form
	//  `projects/{project_id}/locations/{location}/sessions/{session_id}`.
	// +kcc:proto:field=google.cloud.bigquery.storage.v1beta1.ReadSession.name
	Name *string `json:"name,omitempty"`

	// Time at which the session becomes invalid. After this time, subsequent
	//  requests to read this Session will return errors.
	// +kcc:proto:field=google.cloud.bigquery.storage.v1beta1.ReadSession.expire_time
	ExpireTime *string `json:"expireTime,omitempty"`

	// Avro schema.
	// +kcc:proto:field=google.cloud.bigquery.storage.v1beta1.ReadSession.avro_schema
	AvroSchema *AvroSchema `json:"avroSchema,omitempty"`

	// Arrow schema.
	// +kcc:proto:field=google.cloud.bigquery.storage.v1beta1.ReadSession.arrow_schema
	ArrowSchema *ArrowSchema `json:"arrowSchema,omitempty"`

	// Streams associated with this session.
	// +kcc:proto:field=google.cloud.bigquery.storage.v1beta1.ReadSession.streams
	Streams []Stream `json:"streams,omitempty"`

	// Table that this ReadSession is reading from.
	// +kcc:proto:field=google.cloud.bigquery.storage.v1beta1.ReadSession.table_reference
	TableReference *TableReference `json:"tableReference,omitempty"`

	// Any modifiers which are applied when reading from the specified table.
	// +kcc:proto:field=google.cloud.bigquery.storage.v1beta1.ReadSession.table_modifiers
	TableModifiers *TableModifiers `json:"tableModifiers,omitempty"`

	// The strategy to use for distributing data among the streams.
	// +kcc:proto:field=google.cloud.bigquery.storage.v1beta1.ReadSession.sharding_strategy
	ShardingStrategy *string `json:"shardingStrategy,omitempty"`
}

// +kcc:proto=google.cloud.bigquery.storage.v1beta1.Stream
type Stream struct {
	// Name of the stream, in the form
	//  `projects/{project_id}/locations/{location}/streams/{stream_id}`.
	// +kcc:proto:field=google.cloud.bigquery.storage.v1beta1.Stream.name
	Name *string `json:"name,omitempty"`
}

// +kcc:proto=google.cloud.bigquery.storage.v1beta1.TableModifiers
type TableModifiers struct {
	// The snapshot time of the table. If not set, interpreted as now.
	// +kcc:proto:field=google.cloud.bigquery.storage.v1beta1.TableModifiers.snapshot_time
	SnapshotTime *string `json:"snapshotTime,omitempty"`
}

// +kcc:proto=google.cloud.bigquery.storage.v1beta1.TableReference
type TableReference struct {
	// The assigned project ID of the project.
	// +kcc:proto:field=google.cloud.bigquery.storage.v1beta1.TableReference.project_id
	ProjectID *string `json:"projectID,omitempty"`

	// The ID of the dataset in the above project.
	// +kcc:proto:field=google.cloud.bigquery.storage.v1beta1.TableReference.dataset_id
	DatasetID *string `json:"datasetID,omitempty"`

	// The ID of the table in the above dataset.
	// +kcc:proto:field=google.cloud.bigquery.storage.v1beta1.TableReference.table_id
	TableID *string `json:"tableID,omitempty"`
}
