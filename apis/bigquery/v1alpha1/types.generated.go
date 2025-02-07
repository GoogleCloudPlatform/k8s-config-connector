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


// +kcc:proto=google.cloud.bigquery.biglake.v1.Database
type Database struct {
	// Options of a Hive database.
	// +kcc:proto:field=google.cloud.bigquery.biglake.v1.Database.hive_options
	HiveOptions *HiveDatabaseOptions `json:"hiveOptions,omitempty"`

	// The database type.
	// +kcc:proto:field=google.cloud.bigquery.biglake.v1.Database.type
	Type *string `json:"type,omitempty"`
}

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

// +kcc:proto=google.cloud.bigquery.biglake.v1.Database
type DatabaseObservedState struct {
	// Output only. The resource name.
	//  Format:
	//  projects/{project_id_or_number}/locations/{location_id}/catalogs/{catalog_id}/databases/{database_id}
	// +kcc:proto:field=google.cloud.bigquery.biglake.v1.Database.name
	Name *string `json:"name,omitempty"`

	// Output only. The creation time of the database.
	// +kcc:proto:field=google.cloud.bigquery.biglake.v1.Database.create_time
	CreateTime *string `json:"createTime,omitempty"`

	// Output only. The last modification time of the database.
	// +kcc:proto:field=google.cloud.bigquery.biglake.v1.Database.update_time
	UpdateTime *string `json:"updateTime,omitempty"`

	// Output only. The deletion time of the database. Only set after the database
	//  is deleted.
	// +kcc:proto:field=google.cloud.bigquery.biglake.v1.Database.delete_time
	DeleteTime *string `json:"deleteTime,omitempty"`

	// Output only. The time when this database is considered expired. Only set
	//  after the database is deleted.
	// +kcc:proto:field=google.cloud.bigquery.biglake.v1.Database.expire_time
	ExpireTime *string `json:"expireTime,omitempty"`
}
