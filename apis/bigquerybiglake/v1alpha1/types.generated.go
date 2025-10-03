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

// +generated:types
// krm.group: bigquerybiglake.cnrm.cloud.google.com
// krm.version: v1alpha1
// proto.service: google.cloud.bigquery.biglake.v1
// resource: BigLakeTable:Table
// resource: BigLakeCatalog:Catalog
// resource: BigLakeDatabase:Database

package v1alpha1

// +kcc:proto=google.cloud.bigquery.biglake.v1.HiveDatabaseOptions
type HiveDatabaseOptions struct {
	// Cloud Storage folder URI where the database data is stored, starting with
	//  "gs://".
	// +kcc:proto:field=google.cloud.bigquery.biglake.v1.HiveDatabaseOptions.location_uri
	LocationURI *string `json:"locationURI,omitempty"`

	// Stores user supplied Hive database parameters.
	// +kcc:proto:field=google.cloud.bigquery.biglake.v1.HiveDatabaseOptions.parameters
	Parameters map[string]string `json:"parameters,omitempty"`
}

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
