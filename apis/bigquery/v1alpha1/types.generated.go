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


// +kcc:proto=google.cloud.bigquery.storage.v1beta2.TableFieldSchema
type TableFieldSchema struct {
	// Required. The field name. The name must contain only letters (a-z, A-Z),
	//  numbers (0-9), or underscores (_), and must start with a letter or
	//  underscore. The maximum length is 128 characters.
	// +kcc:proto:field=google.cloud.bigquery.storage.v1beta2.TableFieldSchema.name
	Name *string `json:"name,omitempty"`

	// Required. The field data type.
	// +kcc:proto:field=google.cloud.bigquery.storage.v1beta2.TableFieldSchema.type
	Type *string `json:"type,omitempty"`

	// Optional. The field mode. The default value is NULLABLE.
	// +kcc:proto:field=google.cloud.bigquery.storage.v1beta2.TableFieldSchema.mode
	Mode *string `json:"mode,omitempty"`

	// Optional. Describes the nested schema fields if the type property is set to STRUCT.
	// +kcc:proto:field=google.cloud.bigquery.storage.v1beta2.TableFieldSchema.fields
	Fields []TableFieldSchema `json:"fields,omitempty"`

	// Optional. The field description. The maximum length is 1,024 characters.
	// +kcc:proto:field=google.cloud.bigquery.storage.v1beta2.TableFieldSchema.description
	Description *string `json:"description,omitempty"`
}

// +kcc:proto=google.cloud.bigquery.storage.v1beta2.TableSchema
type TableSchema struct {
	// Describes the fields in a table.
	// +kcc:proto:field=google.cloud.bigquery.storage.v1beta2.TableSchema.fields
	Fields []TableFieldSchema `json:"fields,omitempty"`
}

// +kcc:proto=google.cloud.bigquery.storage.v1beta2.WriteStream
type WriteStream struct {

	// Immutable. Type of the stream.
	// +kcc:proto:field=google.cloud.bigquery.storage.v1beta2.WriteStream.type
	Type *string `json:"type,omitempty"`
}

// +kcc:proto=google.cloud.bigquery.storage.v1beta2.WriteStream
type WriteStreamObservedState struct {
	// Output only. Name of the stream, in the form
	//  `projects/{project}/datasets/{dataset}/tables/{table}/streams/{stream}`.
	// +kcc:proto:field=google.cloud.bigquery.storage.v1beta2.WriteStream.name
	Name *string `json:"name,omitempty"`

	// Output only. Create time of the stream. For the _default stream, this is the
	//  creation_time of the table.
	// +kcc:proto:field=google.cloud.bigquery.storage.v1beta2.WriteStream.create_time
	CreateTime *string `json:"createTime,omitempty"`

	// Output only. Commit time of the stream.
	//  If a stream is of `COMMITTED` type, then it will have a commit_time same as
	//  `create_time`. If the stream is of `PENDING` type, commit_time being empty
	//  means it is not committed.
	// +kcc:proto:field=google.cloud.bigquery.storage.v1beta2.WriteStream.commit_time
	CommitTime *string `json:"commitTime,omitempty"`

	// Output only. The schema of the destination table. It is only returned in
	//  `CreateWriteStream` response. Caller should generate data that's
	//  compatible with this schema to send in initial `AppendRowsRequest`.
	//  The table schema could go out of date during the life time of the stream.
	// +kcc:proto:field=google.cloud.bigquery.storage.v1beta2.WriteStream.table_schema
	TableSchema *TableSchema `json:"tableSchema,omitempty"`
}
