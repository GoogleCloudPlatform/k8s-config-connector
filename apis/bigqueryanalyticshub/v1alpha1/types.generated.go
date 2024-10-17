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

// +kcc:proto=google.cloud.bigquery.analyticshub.v1.DataExchange
type DataExchange struct {
	// Output only. The resource name of the data exchange.
	//  e.g. `projects/myproject/locations/US/dataExchanges/123`.
	Name *string `json:"name,omitempty"`

	// Required. Human-readable display name of the data exchange. The display
	//  name must contain only Unicode letters, numbers (0-9), underscores (_),
	//  dashes (-), spaces ( ), ampersands (&) and must not start or end with
	//  spaces. Default value is an empty string. Max length: 63 bytes.
	DisplayName *string `json:"displayName,omitempty"`

	// Optional. Description of the data exchange. The description must not
	//  contain Unicode non-characters as well as C0 and C1 control codes except
	//  tabs (HT), new lines (LF), carriage returns (CR), and page breaks (FF).
	//  Default value is an empty string.
	//  Max length: 2000 bytes.
	Description *string `json:"description,omitempty"`

	// Optional. Email or URL of the primary point of contact of the data
	//  exchange. Max Length: 1000 bytes.
	PrimaryContact *string `json:"primaryContact,omitempty"`

	// Optional. Documentation describing the data exchange.
	Documentation *string `json:"documentation,omitempty"`

	// Optional. Base64 encoded image representing the data exchange. Max
	//  Size: 3.0MiB Expected image dimensions are 512x512 pixels, however the API
	//  only performs validation on size of the encoded data. Note: For byte
	//  fields, the content of the fields are base64-encoded (which increases the
	//  size of the data by 33-36%) when using JSON on the wire.
	Icon []byte `json:"icon,omitempty"`

	// Optional. Configurable data sharing environment option for a data exchange.
	SharingEnvironmentConfig *SharingEnvironmentConfig `json:"sharingEnvironmentConfig,omitempty"`

	// Optional. Type of discovery on the discovery page for all the listings
	//  under this exchange. Updating this field also updates (overwrites) the
	//  discovery_type field for all the listings under this exchange.
	DiscoveryType *string `json:"discoveryType,omitempty"`
}

// +kcc:proto=google.cloud.bigquery.analyticshub.v1.SharingEnvironmentConfig
type SharingEnvironmentConfig struct {
	// Default Analytics Hub data exchange, used for secured data sharing.
	DefaultExchangeConfig *SharingEnvironmentConfig_DefaultExchangeConfig `json:"defaultExchangeConfig,omitempty"`

	// Data Clean Room (DCR), used for privacy-safe and secured data sharing.
	DcrExchangeConfig *SharingEnvironmentConfig_DcrExchangeConfig `json:"dcrExchangeConfig,omitempty"`
}

// +kcc:proto=google.cloud.bigquery.analyticshub.v1.SharingEnvironmentConfig.DcrExchangeConfig
type SharingEnvironmentConfig_DcrExchangeConfig struct {
	// Output only. If True, this DCR restricts the contributors to sharing
	//  only a single resource in a Listing. And no two resources should have the
	//  same IDs. So if a contributor adds a view with a conflicting name, the
	//  CreateListing API will reject the request. if False, the data contributor
	//  can publish an entire dataset (as before). This is not configurable, and
	//  by default, all new DCRs will have the restriction set to True.
	SingleSelectedResourceSharingRestriction *bool `json:"singleSelectedResourceSharingRestriction,omitempty"`

	// Output only. If True, when subscribing to this DCR, it will create only
	//  one linked dataset containing all resources shared within the
	//  cleanroom. If False, when subscribing to this DCR, it will
	//  create 1 linked dataset per listing. This is not configurable, and by
	//  default, all new DCRs will have the restriction set to True.
	SingleLinkedDatasetPerCleanroom *bool `json:"singleLinkedDatasetPerCleanroom,omitempty"`
}

// +kcc:proto=google.cloud.bigquery.analyticshub.v1.SharingEnvironmentConfig.DefaultExchangeConfig
type SharingEnvironmentConfig_DefaultExchangeConfig struct {
}

// +kcc:proto=google.cloud.bigquery.analyticshub.v1.DataProvider
type DataProvider struct {
	// Optional. Name of the data provider.
	Name *string `json:"name,omitempty"`

	// Optional. Email or URL of the data provider.
	//  Max Length: 1000 bytes.
	PrimaryContact *string `json:"primaryContact,omitempty"`
}

// +kcc:proto=google.cloud.bigquery.analyticshub.v1.Listing
type Listing struct {
	// Required. Shared dataset i.e. BigQuery dataset source.
	BigqueryDataset *Listing_BigQueryDatasetSource `json:"bigqueryDataset,omitempty"`

	// Output only. The resource name of the listing.
	//  e.g. `projects/myproject/locations/US/dataExchanges/123/listings/456`
	Name *string `json:"name,omitempty"`

	// Required. Human-readable display name of the listing. The display name must
	//  contain only Unicode letters, numbers (0-9), underscores (_), dashes (-),
	//  spaces ( ), ampersands (&) and can't start or end with spaces. Default
	//  value is an empty string. Max length: 63 bytes.
	DisplayName *string `json:"displayName,omitempty"`

	// Optional. Short description of the listing. The description must not
	//  contain Unicode non-characters and C0 and C1 control codes except tabs
	//  (HT), new lines (LF), carriage returns (CR), and page breaks (FF). Default
	//  value is an empty string. Max length: 2000 bytes.
	Description *string `json:"description,omitempty"`

	// Optional. Email or URL of the primary point of contact of the listing.
	//  Max Length: 1000 bytes.
	PrimaryContact *string `json:"primaryContact,omitempty"`

	// Optional. Documentation describing the listing.
	Documentation *string `json:"documentation,omitempty"`

	// Output only. Current state of the listing.
	State *string `json:"state,omitempty"`

	// Optional. Base64 encoded image representing the listing. Max Size: 3.0MiB
	//  Expected image dimensions are 512x512 pixels, however the API only
	//  performs validation on size of the encoded data.
	//  Note: For byte fields, the contents of the field are base64-encoded (which
	//  increases the size of the data by 33-36%) when using JSON on the wire.
	Icon []byte `json:"icon,omitempty"`

	// Optional. Details of the data provider who owns the source data.
	DataProvider *DataProvider `json:"dataProvider,omitempty"`

	// Optional. Categories of the listing. Up to two categories are allowed.
	Categories []string `json:"categories,omitempty"`

	// Optional. Details of the publisher who owns the listing and who can share
	//  the source data.
	Publisher *Publisher `json:"publisher,omitempty"`

	// Optional. Email or URL of the request access of the listing.
	//  Subscribers can use this reference to request access.
	//  Max Length: 1000 bytes.
	RequestAccess *string `json:"requestAccess,omitempty"`

	// Optional. If set, restricted export configuration will be propagated and
	//  enforced on the linked dataset.
	RestrictedExportConfig *Listing_RestrictedExportConfig `json:"restrictedExportConfig,omitempty"`

	// Optional. Type of discovery of the listing on the discovery page.
	DiscoveryType *string `json:"discoveryType,omitempty"`
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
