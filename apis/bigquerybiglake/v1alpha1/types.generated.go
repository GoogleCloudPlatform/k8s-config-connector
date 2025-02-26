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


// +kcc:proto=google.cloud.bigquery.biglake.v1.HiveTableOptions
type HiveTableOptions struct {
	// Stores user supplied Hive table parameters.
	// +kcc:proto:field=google.cloud.bigquery.biglake.v1.HiveTableOptions.parameters
	Parameters map[string]string `json:"parameters,omitempty"`

	// Hive table type. For example, MANAGED_TABLE, EXTERNAL_TABLE.
	// +kcc:proto:field=google.cloud.bigquery.biglake.v1.HiveTableOptions.table_type
	TableType *string `json:"tableType,omitempty"`

	// Stores physical storage information of the data.
	// +kcc:proto:field=google.cloud.bigquery.biglake.v1.HiveTableOptions.storage_descriptor
	StorageDescriptor *HiveTableOptions_StorageDescriptor `json:"storageDescriptor,omitempty"`
}

// +kcc:proto=google.cloud.bigquery.biglake.v1.HiveTableOptions.SerDeInfo
type HiveTableOptions_SerDeInfo struct {
	// The fully qualified Java class name of the serialization library.
	// +kcc:proto:field=google.cloud.bigquery.biglake.v1.HiveTableOptions.SerDeInfo.serialization_lib
	SerializationLib *string `json:"serializationLib,omitempty"`
}

// +kcc:proto=google.cloud.bigquery.biglake.v1.HiveTableOptions.StorageDescriptor
type HiveTableOptions_StorageDescriptor struct {
	// Cloud Storage folder URI where the table data is stored, starting with
	//  "gs://".
	// +kcc:proto:field=google.cloud.bigquery.biglake.v1.HiveTableOptions.StorageDescriptor.location_uri
	LocationURI *string `json:"locationURI,omitempty"`

	// The fully qualified Java class name of the input format.
	// +kcc:proto:field=google.cloud.bigquery.biglake.v1.HiveTableOptions.StorageDescriptor.input_format
	InputFormat *string `json:"inputFormat,omitempty"`

	// The fully qualified Java class name of the output format.
	// +kcc:proto:field=google.cloud.bigquery.biglake.v1.HiveTableOptions.StorageDescriptor.output_format
	OutputFormat *string `json:"outputFormat,omitempty"`

	// Serializer and deserializer information.
	// +kcc:proto:field=google.cloud.bigquery.biglake.v1.HiveTableOptions.StorageDescriptor.serde_info
	SerdeInfo *HiveTableOptions_SerDeInfo `json:"serdeInfo,omitempty"`
}

// +kcc:proto=google.cloud.bigquery.biglake.v1.Table
type Table struct {
	// Options of a Hive table.
	// +kcc:proto:field=google.cloud.bigquery.biglake.v1.Table.hive_options
	HiveOptions *HiveTableOptions `json:"hiveOptions,omitempty"`

	// The table type.
	// +kcc:proto:field=google.cloud.bigquery.biglake.v1.Table.type
	Type *string `json:"type,omitempty"`

	// The checksum of a table object computed by the server based on the value of
	//  other fields. It may be sent on update requests to ensure the client has an
	//  up-to-date value before proceeding. It is only checked for update table
	//  operations.
	// +kcc:proto:field=google.cloud.bigquery.biglake.v1.Table.etag
	Etag *string `json:"etag,omitempty"`
}

// +kcc:proto=google.cloud.bigquery.biglake.v1.Table
type TableObservedState struct {
	// Output only. The resource name.
	//  Format:
	//  projects/{project_id_or_number}/locations/{location_id}/catalogs/{catalog_id}/databases/{database_id}/tables/{table_id}
	// +kcc:proto:field=google.cloud.bigquery.biglake.v1.Table.name
	Name *string `json:"name,omitempty"`

	// Output only. The creation time of the table.
	// +kcc:proto:field=google.cloud.bigquery.biglake.v1.Table.create_time
	CreateTime *string `json:"createTime,omitempty"`

	// Output only. The last modification time of the table.
	// +kcc:proto:field=google.cloud.bigquery.biglake.v1.Table.update_time
	UpdateTime *string `json:"updateTime,omitempty"`

	// Output only. The deletion time of the table. Only set after the table is
	//  deleted.
	// +kcc:proto:field=google.cloud.bigquery.biglake.v1.Table.delete_time
	DeleteTime *string `json:"deleteTime,omitempty"`

	// Output only. The time when this table is considered expired. Only set after
	//  the table is deleted.
	// +kcc:proto:field=google.cloud.bigquery.biglake.v1.Table.expire_time
	ExpireTime *string `json:"expireTime,omitempty"`
}
