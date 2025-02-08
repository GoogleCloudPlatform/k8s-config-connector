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


// +kcc:proto=google.cloud.documentai.v1beta3.Dataset
type Dataset struct {
	// Optional. User-managed Cloud Storage dataset configuration. Use this
	//  configuration if the dataset documents are stored under a user-managed
	//  Cloud Storage location.
	// +kcc:proto:field=google.cloud.documentai.v1beta3.Dataset.gcs_managed_config
	GcsManagedConfig *Dataset_GCSManagedConfig `json:"gcsManagedConfig,omitempty"`

	// Optional. Deprecated. Warehouse-based dataset configuration is not
	//  supported.
	// +kcc:proto:field=google.cloud.documentai.v1beta3.Dataset.document_warehouse_config
	DocumentWarehouseConfig *Dataset_DocumentWarehouseConfig `json:"documentWarehouseConfig,omitempty"`

	// Optional. Unmanaged dataset configuration. Use this configuration if the
	//  dataset documents are managed by the document service internally (not
	//  user-managed).
	// +kcc:proto:field=google.cloud.documentai.v1beta3.Dataset.unmanaged_dataset_config
	UnmanagedDatasetConfig *Dataset_UnmanagedDatasetConfig `json:"unmanagedDatasetConfig,omitempty"`

	// Optional. A lightweight indexing source with low latency and high
	//  reliability, but lacking advanced features like CMEK and content-based
	//  search.
	// +kcc:proto:field=google.cloud.documentai.v1beta3.Dataset.spanner_indexing_config
	SpannerIndexingConfig *Dataset_SpannerIndexingConfig `json:"spannerIndexingConfig,omitempty"`

	// Dataset resource name.
	//  Format:
	//  `projects/{project}/locations/{location}/processors/{processor}/dataset`
	// +kcc:proto:field=google.cloud.documentai.v1beta3.Dataset.name
	Name *string `json:"name,omitempty"`

	// Required. State of the dataset. Ignored when updating dataset.
	// +kcc:proto:field=google.cloud.documentai.v1beta3.Dataset.state
	State *string `json:"state,omitempty"`
}

// +kcc:proto=google.cloud.documentai.v1beta3.Dataset.DocumentWarehouseConfig
type Dataset_DocumentWarehouseConfig struct {
}

// +kcc:proto=google.cloud.documentai.v1beta3.Dataset.GCSManagedConfig
type Dataset_GCSManagedConfig struct {
	// Required. The Cloud Storage URI (a directory) where the documents
	//  belonging to the dataset must be stored.
	// +kcc:proto:field=google.cloud.documentai.v1beta3.Dataset.GCSManagedConfig.gcs_prefix
	GcsPrefix *GcsPrefix `json:"gcsPrefix,omitempty"`
}

// +kcc:proto=google.cloud.documentai.v1beta3.Dataset.SpannerIndexingConfig
type Dataset_SpannerIndexingConfig struct {
}

// +kcc:proto=google.cloud.documentai.v1beta3.Dataset.UnmanagedDatasetConfig
type Dataset_UnmanagedDatasetConfig struct {
}

// +kcc:proto=google.cloud.documentai.v1beta3.GcsPrefix
type GcsPrefix struct {
	// The URI prefix.
	// +kcc:proto:field=google.cloud.documentai.v1beta3.GcsPrefix.gcs_uri_prefix
	GcsURIPrefix *string `json:"gcsURIPrefix,omitempty"`
}

// +kcc:proto=google.cloud.documentai.v1beta3.Dataset
type DatasetObservedState struct {
	// Optional. Deprecated. Warehouse-based dataset configuration is not
	//  supported.
	// +kcc:proto:field=google.cloud.documentai.v1beta3.Dataset.document_warehouse_config
	DocumentWarehouseConfig *Dataset_DocumentWarehouseConfigObservedState `json:"documentWarehouseConfig,omitempty"`

	// Output only. Reserved for future use.
	// +kcc:proto:field=google.cloud.documentai.v1beta3.Dataset.satisfies_pzs
	SatisfiesPzs *bool `json:"satisfiesPzs,omitempty"`

	// Output only. Reserved for future use.
	// +kcc:proto:field=google.cloud.documentai.v1beta3.Dataset.satisfies_pzi
	SatisfiesPzi *bool `json:"satisfiesPzi,omitempty"`
}

// +kcc:proto=google.cloud.documentai.v1beta3.Dataset.DocumentWarehouseConfig
type Dataset_DocumentWarehouseConfigObservedState struct {
	// Output only. The collection in Document AI Warehouse associated with the
	//  dataset.
	// +kcc:proto:field=google.cloud.documentai.v1beta3.Dataset.DocumentWarehouseConfig.collection
	Collection *string `json:"collection,omitempty"`

	// Output only. The schema in Document AI Warehouse associated with the
	//  dataset.
	// +kcc:proto:field=google.cloud.documentai.v1beta3.Dataset.DocumentWarehouseConfig.schema
	Schema *string `json:"schema,omitempty"`
}
