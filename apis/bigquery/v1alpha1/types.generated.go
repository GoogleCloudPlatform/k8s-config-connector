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


// +kcc:proto=google.cloud.bigquery.analyticshub.v1.DataProvider
type DataProvider struct {
	// Optional. Name of the data provider.
	// +kcc:proto:field=google.cloud.bigquery.analyticshub.v1.DataProvider.name
	Name *string `json:"name,omitempty"`

	// Optional. Email or URL of the data provider.
	//  Max Length: 1000 bytes.
	// +kcc:proto:field=google.cloud.bigquery.analyticshub.v1.DataProvider.primary_contact
	PrimaryContact *string `json:"primaryContact,omitempty"`
}

// +kcc:proto=google.cloud.bigquery.analyticshub.v1.Listing
type Listing struct {
	// Required. Shared dataset i.e. BigQuery dataset source.
	// +kcc:proto:field=google.cloud.bigquery.analyticshub.v1.Listing.bigquery_dataset
	BigqueryDataset *Listing_BigQueryDatasetSource `json:"bigqueryDataset,omitempty"`

	// Required. Human-readable display name of the listing. The display name must
	//  contain only Unicode letters, numbers (0-9), underscores (_), dashes (-),
	//  spaces ( ), ampersands (&) and can't start or end with spaces. Default
	//  value is an empty string. Max length: 63 bytes.
	// +kcc:proto:field=google.cloud.bigquery.analyticshub.v1.Listing.display_name
	DisplayName *string `json:"displayName,omitempty"`

	// Optional. Short description of the listing. The description must not
	//  contain Unicode non-characters and C0 and C1 control codes except tabs
	//  (HT), new lines (LF), carriage returns (CR), and page breaks (FF). Default
	//  value is an empty string. Max length: 2000 bytes.
	// +kcc:proto:field=google.cloud.bigquery.analyticshub.v1.Listing.description
	Description *string `json:"description,omitempty"`

	// Optional. Email or URL of the primary point of contact of the listing.
	//  Max Length: 1000 bytes.
	// +kcc:proto:field=google.cloud.bigquery.analyticshub.v1.Listing.primary_contact
	PrimaryContact *string `json:"primaryContact,omitempty"`

	// Optional. Documentation describing the listing.
	// +kcc:proto:field=google.cloud.bigquery.analyticshub.v1.Listing.documentation
	Documentation *string `json:"documentation,omitempty"`

	// Optional. Base64 encoded image representing the listing. Max Size: 3.0MiB
	//  Expected image dimensions are 512x512 pixels, however the API only
	//  performs validation on size of the encoded data.
	//  Note: For byte fields, the contents of the field are base64-encoded (which
	//  increases the size of the data by 33-36%) when using JSON on the wire.
	// +kcc:proto:field=google.cloud.bigquery.analyticshub.v1.Listing.icon
	Icon []byte `json:"icon,omitempty"`

	// Optional. Details of the data provider who owns the source data.
	// +kcc:proto:field=google.cloud.bigquery.analyticshub.v1.Listing.data_provider
	DataProvider *DataProvider `json:"dataProvider,omitempty"`

	// Optional. Categories of the listing. Up to two categories are allowed.
	// +kcc:proto:field=google.cloud.bigquery.analyticshub.v1.Listing.categories
	Categories []string `json:"categories,omitempty"`

	// Optional. Details of the publisher who owns the listing and who can share
	//  the source data.
	// +kcc:proto:field=google.cloud.bigquery.analyticshub.v1.Listing.publisher
	Publisher *Publisher `json:"publisher,omitempty"`

	// Optional. Email or URL of the request access of the listing.
	//  Subscribers can use this reference to request access.
	//  Max Length: 1000 bytes.
	// +kcc:proto:field=google.cloud.bigquery.analyticshub.v1.Listing.request_access
	RequestAccess *string `json:"requestAccess,omitempty"`

	// Optional. If set, restricted export configuration will be propagated and
	//  enforced on the linked dataset.
	// +kcc:proto:field=google.cloud.bigquery.analyticshub.v1.Listing.restricted_export_config
	RestrictedExportConfig *Listing_RestrictedExportConfig `json:"restrictedExportConfig,omitempty"`

	// Optional. Type of discovery of the listing on the discovery page.
	// +kcc:proto:field=google.cloud.bigquery.analyticshub.v1.Listing.discovery_type
	DiscoveryType *string `json:"discoveryType,omitempty"`
}

// +kcc:proto=google.cloud.bigquery.analyticshub.v1.Listing.BigQueryDatasetSource
type Listing_BigQueryDatasetSource struct {
	// Resource name of the dataset source for this listing.
	//  e.g. `projects/myproject/datasets/123`
	// +kcc:proto:field=google.cloud.bigquery.analyticshub.v1.Listing.BigQueryDatasetSource.dataset
	Dataset *string `json:"dataset,omitempty"`

	// Optional. Resources in this dataset that are selectively shared.
	//  If this field is empty, then the entire dataset (all resources) are
	//  shared. This field is only valid for data clean room exchanges.
	// +kcc:proto:field=google.cloud.bigquery.analyticshub.v1.Listing.BigQueryDatasetSource.selected_resources
	SelectedResources []Listing_BigQueryDatasetSource_SelectedResource `json:"selectedResources,omitempty"`

	// Optional. If set, restricted export policy will be propagated and
	//  enforced on the linked dataset.
	// +kcc:proto:field=google.cloud.bigquery.analyticshub.v1.Listing.BigQueryDatasetSource.restricted_export_policy
	RestrictedExportPolicy *Listing_BigQueryDatasetSource_RestrictedExportPolicy `json:"restrictedExportPolicy,omitempty"`
}

// +kcc:proto=google.cloud.bigquery.analyticshub.v1.Listing.BigQueryDatasetSource.RestrictedExportPolicy
type Listing_BigQueryDatasetSource_RestrictedExportPolicy struct {
	// Optional. If true, enable restricted export.
	// +kcc:proto:field=google.cloud.bigquery.analyticshub.v1.Listing.BigQueryDatasetSource.RestrictedExportPolicy.enabled
	Enabled *bool `json:"enabled,omitempty"`

	// Optional. If true, restrict direct table access (read
	//  api/tabledata.list) on linked table.
	// +kcc:proto:field=google.cloud.bigquery.analyticshub.v1.Listing.BigQueryDatasetSource.RestrictedExportPolicy.restrict_direct_table_access
	RestrictDirectTableAccess *bool `json:"restrictDirectTableAccess,omitempty"`

	// Optional. If true, restrict export of query result derived from
	//  restricted linked dataset table.
	// +kcc:proto:field=google.cloud.bigquery.analyticshub.v1.Listing.BigQueryDatasetSource.RestrictedExportPolicy.restrict_query_result
	RestrictQueryResult *bool `json:"restrictQueryResult,omitempty"`
}

// +kcc:proto=google.cloud.bigquery.analyticshub.v1.Listing.BigQueryDatasetSource.SelectedResource
type Listing_BigQueryDatasetSource_SelectedResource struct {
	// Optional. Format:
	//  For table:
	//  `projects/{projectId}/datasets/{datasetId}/tables/{tableId}`
	//  Example:"projects/test_project/datasets/test_dataset/tables/test_table"
	// +kcc:proto:field=google.cloud.bigquery.analyticshub.v1.Listing.BigQueryDatasetSource.SelectedResource.table
	Table *string `json:"table,omitempty"`
}

// +kcc:proto=google.cloud.bigquery.analyticshub.v1.Listing.RestrictedExportConfig
type Listing_RestrictedExportConfig struct {
	// Optional. If true, enable restricted export.
	// +kcc:proto:field=google.cloud.bigquery.analyticshub.v1.Listing.RestrictedExportConfig.enabled
	Enabled *bool `json:"enabled,omitempty"`

	// Optional. If true, restrict export of query result derived from
	//  restricted linked dataset table.
	// +kcc:proto:field=google.cloud.bigquery.analyticshub.v1.Listing.RestrictedExportConfig.restrict_query_result
	RestrictQueryResult *bool `json:"restrictQueryResult,omitempty"`
}

// +kcc:proto=google.cloud.bigquery.analyticshub.v1.Publisher
type Publisher struct {
	// Optional. Name of the listing publisher.
	// +kcc:proto:field=google.cloud.bigquery.analyticshub.v1.Publisher.name
	Name *string `json:"name,omitempty"`

	// Optional. Email or URL of the listing publisher.
	//  Max Length: 1000 bytes.
	// +kcc:proto:field=google.cloud.bigquery.analyticshub.v1.Publisher.primary_contact
	PrimaryContact *string `json:"primaryContact,omitempty"`
}

// +kcc:proto=google.cloud.bigquery.analyticshub.v1.Listing
type ListingObservedState struct {
	// Output only. The resource name of the listing.
	//  e.g. `projects/myproject/locations/US/dataExchanges/123/listings/456`
	// +kcc:proto:field=google.cloud.bigquery.analyticshub.v1.Listing.name
	Name *string `json:"name,omitempty"`

	// Output only. Current state of the listing.
	// +kcc:proto:field=google.cloud.bigquery.analyticshub.v1.Listing.state
	State *string `json:"state,omitempty"`

	// Optional. If set, restricted export configuration will be propagated and
	//  enforced on the linked dataset.
	// +kcc:proto:field=google.cloud.bigquery.analyticshub.v1.Listing.restricted_export_config
	RestrictedExportConfig *Listing_RestrictedExportConfigObservedState `json:"restrictedExportConfig,omitempty"`
}

// +kcc:proto=google.cloud.bigquery.analyticshub.v1.Listing.RestrictedExportConfig
type Listing_RestrictedExportConfigObservedState struct {
	// Output only. If true, restrict direct table access(read
	//  api/tabledata.list) on linked table.
	// +kcc:proto:field=google.cloud.bigquery.analyticshub.v1.Listing.RestrictedExportConfig.restrict_direct_table_access
	RestrictDirectTableAccess *bool `json:"restrictDirectTableAccess,omitempty"`
}
