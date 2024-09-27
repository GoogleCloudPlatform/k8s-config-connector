// Copyright 2024 Google LLC
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

// +kcc:proto=google.cloud.bigquery.analyticshub.v1.DataProvider
type DataProvider struct {
	// Optional. Name of the data provider.
	Name *string `json:"name,omitempty"`

	// Optional. Email or URL of the data provider.
	//  Max Length: 1000 bytes.
	PrimaryContact *string `json:"primaryContact,omitempty"`
}

// +kcc:proto=google.cloud.bigquery.analyticshub.v1.Listing.BigQueryDatasetSource
type Listing_BigQueryDatasetSource struct {
	// Resource name of the dataset source for this listing.
	//  e.g. `projects/myproject/datasets/123`
	Dataset *string `json:"dataset,omitempty"`

	// Optional. Resources in this dataset that are selectively shared.
	//  If this field is empty, then the entire dataset (all resources) are
	//  shared. This field is only valid for data clean room exchanges.
	SelectedResources []Listing_BigQueryDatasetSource_SelectedResource `json:"selectedResources,omitempty"`

	// Optional. If set, restricted export policy will be propagated and
	//  enforced on the linked dataset.
	RestrictedExportPolicy *Listing_BigQueryDatasetSource_RestrictedExportPolicy `json:"restrictedExportPolicy,omitempty"`
}

// +kcc:proto=google.cloud.bigquery.analyticshub.v1.Listing.BigQueryDatasetSource.RestrictedExportPolicy
type Listing_BigQueryDatasetSource_RestrictedExportPolicy struct {
	// Optional. If true, enable restricted export.
	Enabled *BoolValue `json:"enabled,omitempty"`

	// Optional. If true, restrict direct table access (read
	//  api/tabledata.list) on linked table.
	RestrictDirectTableAccess *BoolValue `json:"restrictDirectTableAccess,omitempty"`

	// Optional. If true, restrict export of query result derived from
	//  restricted linked dataset table.
	RestrictQueryResult *BoolValue `json:"restrictQueryResult,omitempty"`
}

// +kcc:proto=google.cloud.bigquery.analyticshub.v1.Listing.BigQueryDatasetSource.SelectedResource
type Listing_BigQueryDatasetSource_SelectedResource struct {
	// Optional. Format:
	//  For table:
	//  `projects/{projectId}/datasets/{datasetId}/tables/{tableId}`
	//  Example:"projects/test_project/datasets/test_dataset/tables/test_table"
	Table *string `json:"table,omitempty"`
}

// +kcc:proto=google.cloud.bigquery.analyticshub.v1.Listing.RestrictedExportConfig
type Listing_RestrictedExportConfig struct {
	// Optional. If true, enable restricted export.
	Enabled *bool `json:"enabled,omitempty"`

	// Output only. If true, restrict direct table access(read
	//  api/tabledata.list) on linked table.
	RestrictDirectTableAccess *bool `json:"restrictDirectTableAccess,omitempty"`

	// Optional. If true, restrict export of query result derived from
	//  restricted linked dataset table.
	RestrictQueryResult *bool `json:"restrictQueryResult,omitempty"`
}

// +kcc:proto=google.cloud.bigquery.analyticshub.v1.Publisher
type Publisher struct {
	// Optional. Name of the listing publisher.
	Name *string `json:"name,omitempty"`

	// Optional. Email or URL of the listing publisher.
	//  Max Length: 1000 bytes.
	PrimaryContact *string `json:"primaryContact,omitempty"`
}

// +kcc:proto=google.protobuf.BoolValue
type BoolValue struct {
	// The bool value.
	Value *bool `json:"value,omitempty"`
}
