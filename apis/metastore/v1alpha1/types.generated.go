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


// +kcc:proto=google.cloud.metastore.v1.MetadataImport
type MetadataImport struct {
	// Immutable. A database dump from a pre-existing metastore's database.
	// +kcc:proto:field=google.cloud.metastore.v1.MetadataImport.database_dump
	DatabaseDump *MetadataImport_DatabaseDump `json:"databaseDump,omitempty"`

	// Immutable. The relative resource name of the metadata import, of the form:
	//
	//  `projects/{project_number}/locations/{location_id}/services/{service_id}/metadataImports/{metadata_import_id}`.
	// +kcc:proto:field=google.cloud.metastore.v1.MetadataImport.name
	Name *string `json:"name,omitempty"`

	// The description of the metadata import.
	// +kcc:proto:field=google.cloud.metastore.v1.MetadataImport.description
	Description *string `json:"description,omitempty"`
}

// +kcc:proto=google.cloud.metastore.v1.MetadataImport.DatabaseDump
type MetadataImport_DatabaseDump struct {
	// The type of the database.
	// +kcc:proto:field=google.cloud.metastore.v1.MetadataImport.DatabaseDump.database_type
	DatabaseType *string `json:"databaseType,omitempty"`

	// A Cloud Storage object or folder URI that specifies the source from which
	//  to import metadata. It must begin with `gs://`.
	// +kcc:proto:field=google.cloud.metastore.v1.MetadataImport.DatabaseDump.gcs_uri
	GcsURI *string `json:"gcsURI,omitempty"`

	// The name of the source database.
	// +kcc:proto:field=google.cloud.metastore.v1.MetadataImport.DatabaseDump.source_database
	SourceDatabase *string `json:"sourceDatabase,omitempty"`

	// Optional. The type of the database dump. If unspecified, defaults to
	//  `MYSQL`.
	// +kcc:proto:field=google.cloud.metastore.v1.MetadataImport.DatabaseDump.type
	Type *string `json:"type,omitempty"`
}

// +kcc:proto=google.cloud.metastore.v1.MetadataImport
type MetadataImportObservedState struct {
	// Output only. The time when the metadata import was started.
	// +kcc:proto:field=google.cloud.metastore.v1.MetadataImport.create_time
	CreateTime *string `json:"createTime,omitempty"`

	// Output only. The time when the metadata import was last updated.
	// +kcc:proto:field=google.cloud.metastore.v1.MetadataImport.update_time
	UpdateTime *string `json:"updateTime,omitempty"`

	// Output only. The time when the metadata import finished.
	// +kcc:proto:field=google.cloud.metastore.v1.MetadataImport.end_time
	EndTime *string `json:"endTime,omitempty"`

	// Output only. The current state of the metadata import.
	// +kcc:proto:field=google.cloud.metastore.v1.MetadataImport.state
	State *string `json:"state,omitempty"`
}
