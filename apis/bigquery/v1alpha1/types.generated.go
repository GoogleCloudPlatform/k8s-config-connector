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
// krm.group: bigquery.cnrm.cloud.google.com
// krm.version: v1alpha1
// proto.service: google.cloud.bigquery.biglake.v1
// resource: BigqueryCatalog:Catalog

package v1alpha1

// +kcc:proto=google.cloud.bigquery.biglake.v1.Catalog
type Catalog struct {
}

// +kcc:proto=google.cloud.bigquery.biglake.v1.Catalog
type CatalogObservedState struct {
	// Output only. The resource name.
	//  Format:
	//  projects/{project_id_or_number}/locations/{location_id}/catalogs/{catalog_id}
	// +kcc:proto:field=google.cloud.bigquery.biglake.v1.Catalog.name
	Name *string `json:"name,omitempty"`

	// Output only. The creation time of the catalog.
	// +kcc:proto:field=google.cloud.bigquery.biglake.v1.Catalog.create_time
	CreateTime *string `json:"createTime,omitempty"`

	// Output only. The last modification time of the catalog.
	// +kcc:proto:field=google.cloud.bigquery.biglake.v1.Catalog.update_time
	UpdateTime *string `json:"updateTime,omitempty"`

	// Output only. The deletion time of the catalog. Only set after the catalog
	//  is deleted.
	// +kcc:proto:field=google.cloud.bigquery.biglake.v1.Catalog.delete_time
	DeleteTime *string `json:"deleteTime,omitempty"`

	// Output only. The time when this catalog is considered expired. Only set
	//  after the catalog is deleted.
	// +kcc:proto:field=google.cloud.bigquery.biglake.v1.Catalog.expire_time
	ExpireTime *string `json:"expireTime,omitempty"`
}
