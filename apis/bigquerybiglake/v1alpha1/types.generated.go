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
