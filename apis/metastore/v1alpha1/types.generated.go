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

// +kcc:proto=google.cloud.metastore.v1.BackendMetastore
type BackendMetastore struct {
	// The relative resource name of the metastore that is being federated.
	//  The formats of the relative resource names for the currently supported
	//  metastores are listed below:
	//
	//  * BigQuery
	//      * `projects/{project_id}`
	//  * Dataproc Metastore
	//      * `projects/{project_id}/locations/{location}/services/{service_id}`
	// +kcc:proto:field=google.cloud.metastore.v1.BackendMetastore.name
	Name *string `json:"name,omitempty"`

	// The type of the backend metastore.
	// +kcc:proto:field=google.cloud.metastore.v1.BackendMetastore.metastore_type
	MetastoreType *string `json:"metastoreType,omitempty"`
}
