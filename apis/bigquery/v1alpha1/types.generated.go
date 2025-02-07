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


// +kcc:proto=google.cloud.bigquery.analyticshub.v1.DataExchange
type DataExchange struct {

	// Required. Human-readable display name of the data exchange. The display
	//  name must contain only Unicode letters, numbers (0-9), underscores (_),
	//  dashes (-), spaces ( ), ampersands (&) and must not start or end with
	//  spaces. Default value is an empty string. Max length: 63 bytes.
	// +kcc:proto:field=google.cloud.bigquery.analyticshub.v1.DataExchange.display_name
	DisplayName *string `json:"displayName,omitempty"`

	// Optional. Description of the data exchange. The description must not
	//  contain Unicode non-characters as well as C0 and C1 control codes except
	//  tabs (HT), new lines (LF), carriage returns (CR), and page breaks (FF).
	//  Default value is an empty string.
	//  Max length: 2000 bytes.
	// +kcc:proto:field=google.cloud.bigquery.analyticshub.v1.DataExchange.description
	Description *string `json:"description,omitempty"`

	// Optional. Email or URL of the primary point of contact of the data
	//  exchange. Max Length: 1000 bytes.
	// +kcc:proto:field=google.cloud.bigquery.analyticshub.v1.DataExchange.primary_contact
	PrimaryContact *string `json:"primaryContact,omitempty"`

	// Optional. Documentation describing the data exchange.
	// +kcc:proto:field=google.cloud.bigquery.analyticshub.v1.DataExchange.documentation
	Documentation *string `json:"documentation,omitempty"`

	// Optional. Base64 encoded image representing the data exchange. Max
	//  Size: 3.0MiB Expected image dimensions are 512x512 pixels, however the API
	//  only performs validation on size of the encoded data. Note: For byte
	//  fields, the content of the fields are base64-encoded (which increases the
	//  size of the data by 33-36%) when using JSON on the wire.
	// +kcc:proto:field=google.cloud.bigquery.analyticshub.v1.DataExchange.icon
	Icon []byte `json:"icon,omitempty"`

	// Optional. Configurable data sharing environment option for a data exchange.
	// +kcc:proto:field=google.cloud.bigquery.analyticshub.v1.DataExchange.sharing_environment_config
	SharingEnvironmentConfig *SharingEnvironmentConfig `json:"sharingEnvironmentConfig,omitempty"`

	// Optional. Type of discovery on the discovery page for all the listings
	//  under this exchange. Updating this field also updates (overwrites) the
	//  discovery_type field for all the listings under this exchange.
	// +kcc:proto:field=google.cloud.bigquery.analyticshub.v1.DataExchange.discovery_type
	DiscoveryType *string `json:"discoveryType,omitempty"`
}

// +kcc:proto=google.cloud.bigquery.analyticshub.v1.SharingEnvironmentConfig
type SharingEnvironmentConfig struct {
	// Default Analytics Hub data exchange, used for secured data sharing.
	// +kcc:proto:field=google.cloud.bigquery.analyticshub.v1.SharingEnvironmentConfig.default_exchange_config
	DefaultExchangeConfig *SharingEnvironmentConfig_DefaultExchangeConfig `json:"defaultExchangeConfig,omitempty"`

	// Data Clean Room (DCR), used for privacy-safe and secured data sharing.
	// +kcc:proto:field=google.cloud.bigquery.analyticshub.v1.SharingEnvironmentConfig.dcr_exchange_config
	DcrExchangeConfig *SharingEnvironmentConfig_DcrExchangeConfig `json:"dcrExchangeConfig,omitempty"`
}

// +kcc:proto=google.cloud.bigquery.analyticshub.v1.SharingEnvironmentConfig.DcrExchangeConfig
type SharingEnvironmentConfig_DcrExchangeConfig struct {
}

// +kcc:proto=google.cloud.bigquery.analyticshub.v1.SharingEnvironmentConfig.DefaultExchangeConfig
type SharingEnvironmentConfig_DefaultExchangeConfig struct {
}

// +kcc:proto=google.cloud.bigquery.analyticshub.v1.DataExchange
type DataExchangeObservedState struct {
	// Output only. The resource name of the data exchange.
	//  e.g. `projects/myproject/locations/US/dataExchanges/123`.
	// +kcc:proto:field=google.cloud.bigquery.analyticshub.v1.DataExchange.name
	Name *string `json:"name,omitempty"`

	// Output only. Number of listings contained in the data exchange.
	// +kcc:proto:field=google.cloud.bigquery.analyticshub.v1.DataExchange.listing_count
	ListingCount *int32 `json:"listingCount,omitempty"`

	// Optional. Configurable data sharing environment option for a data exchange.
	// +kcc:proto:field=google.cloud.bigquery.analyticshub.v1.DataExchange.sharing_environment_config
	SharingEnvironmentConfig *SharingEnvironmentConfigObservedState `json:"sharingEnvironmentConfig,omitempty"`
}

// +kcc:proto=google.cloud.bigquery.analyticshub.v1.SharingEnvironmentConfig
type SharingEnvironmentConfigObservedState struct {
	// Data Clean Room (DCR), used for privacy-safe and secured data sharing.
	// +kcc:proto:field=google.cloud.bigquery.analyticshub.v1.SharingEnvironmentConfig.dcr_exchange_config
	DcrExchangeConfig *SharingEnvironmentConfig_DcrExchangeConfigObservedState `json:"dcrExchangeConfig,omitempty"`
}

// +kcc:proto=google.cloud.bigquery.analyticshub.v1.SharingEnvironmentConfig.DcrExchangeConfig
type SharingEnvironmentConfig_DcrExchangeConfigObservedState struct {
	// Output only. If True, this DCR restricts the contributors to sharing
	//  only a single resource in a Listing. And no two resources should have the
	//  same IDs. So if a contributor adds a view with a conflicting name, the
	//  CreateListing API will reject the request. if False, the data contributor
	//  can publish an entire dataset (as before). This is not configurable, and
	//  by default, all new DCRs will have the restriction set to True.
	// +kcc:proto:field=google.cloud.bigquery.analyticshub.v1.SharingEnvironmentConfig.DcrExchangeConfig.single_selected_resource_sharing_restriction
	SingleSelectedResourceSharingRestriction *bool `json:"singleSelectedResourceSharingRestriction,omitempty"`

	// Output only. If True, when subscribing to this DCR, it will create only
	//  one linked dataset containing all resources shared within the
	//  cleanroom. If False, when subscribing to this DCR, it will
	//  create 1 linked dataset per listing. This is not configurable, and by
	//  default, all new DCRs will have the restriction set to True.
	// +kcc:proto:field=google.cloud.bigquery.analyticshub.v1.SharingEnvironmentConfig.DcrExchangeConfig.single_linked_dataset_per_cleanroom
	SingleLinkedDatasetPerCleanroom *bool `json:"singleLinkedDatasetPerCleanroom,omitempty"`
}
