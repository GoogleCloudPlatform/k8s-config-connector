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


// +kcc:proto=google.cloud.dataplex.v1.AssetStatus
type AssetStatus struct {
	// Last update time of the status.
	// +kcc:proto:field=google.cloud.dataplex.v1.AssetStatus.update_time
	UpdateTime *string `json:"updateTime,omitempty"`

	// Number of active assets.
	// +kcc:proto:field=google.cloud.dataplex.v1.AssetStatus.active_assets
	ActiveAssets *int32 `json:"activeAssets,omitempty"`

	// Number of assets that are in process of updating the security policy on
	//  attached resources.
	// +kcc:proto:field=google.cloud.dataplex.v1.AssetStatus.security_policy_applying_assets
	SecurityPolicyApplyingAssets *int32 `json:"securityPolicyApplyingAssets,omitempty"`
}

// +kcc:proto=google.cloud.dataplex.v1.Zone
type Zone struct {

	// Optional. User friendly display name.
	// +kcc:proto:field=google.cloud.dataplex.v1.Zone.display_name
	DisplayName *string `json:"displayName,omitempty"`

	// Optional. User defined labels for the zone.
	// +kcc:proto:field=google.cloud.dataplex.v1.Zone.labels
	Labels map[string]string `json:"labels,omitempty"`

	// Optional. Description of the zone.
	// +kcc:proto:field=google.cloud.dataplex.v1.Zone.description
	Description *string `json:"description,omitempty"`

	// Required. Immutable. The type of the zone.
	// +kcc:proto:field=google.cloud.dataplex.v1.Zone.type
	Type *string `json:"type,omitempty"`

	// Optional. Specification of the discovery feature applied to data in this
	//  zone.
	// +kcc:proto:field=google.cloud.dataplex.v1.Zone.discovery_spec
	DiscoverySpec *Zone_DiscoverySpec `json:"discoverySpec,omitempty"`

	// Required. Specification of the resources that are referenced by the assets
	//  within this zone.
	// +kcc:proto:field=google.cloud.dataplex.v1.Zone.resource_spec
	ResourceSpec *Zone_ResourceSpec `json:"resourceSpec,omitempty"`
}

// +kcc:proto=google.cloud.dataplex.v1.Zone.DiscoverySpec
type Zone_DiscoverySpec struct {
	// Required. Whether discovery is enabled.
	// +kcc:proto:field=google.cloud.dataplex.v1.Zone.DiscoverySpec.enabled
	Enabled *bool `json:"enabled,omitempty"`

	// Optional. The list of patterns to apply for selecting data to include
	//  during discovery if only a subset of the data should considered. For
	//  Cloud Storage bucket assets, these are interpreted as glob patterns used
	//  to match object names. For BigQuery dataset assets, these are interpreted
	//  as patterns to match table names.
	// +kcc:proto:field=google.cloud.dataplex.v1.Zone.DiscoverySpec.include_patterns
	IncludePatterns []string `json:"includePatterns,omitempty"`

	// Optional. The list of patterns to apply for selecting data to exclude
	//  during discovery.  For Cloud Storage bucket assets, these are interpreted
	//  as glob patterns used to match object names. For BigQuery dataset assets,
	//  these are interpreted as patterns to match table names.
	// +kcc:proto:field=google.cloud.dataplex.v1.Zone.DiscoverySpec.exclude_patterns
	ExcludePatterns []string `json:"excludePatterns,omitempty"`

	// Optional. Configuration for CSV data.
	// +kcc:proto:field=google.cloud.dataplex.v1.Zone.DiscoverySpec.csv_options
	CsvOptions *Zone_DiscoverySpec_CsvOptions `json:"csvOptions,omitempty"`

	// Optional. Configuration for Json data.
	// +kcc:proto:field=google.cloud.dataplex.v1.Zone.DiscoverySpec.json_options
	JsonOptions *Zone_DiscoverySpec_JsonOptions `json:"jsonOptions,omitempty"`

	// Optional. Cron schedule (https://en.wikipedia.org/wiki/Cron) for
	//  running discovery periodically. Successive discovery runs must be
	//  scheduled at least 60 minutes apart. The default value is to run
	//  discovery every 60 minutes. To explicitly set a timezone to the cron
	//  tab, apply a prefix in the cron tab: "CRON_TZ=${IANA_TIME_ZONE}" or
	//  TZ=${IANA_TIME_ZONE}". The ${IANA_TIME_ZONE} may only be a valid string
	//  from IANA time zone database. For example, `CRON_TZ=America/New_York 1
	//  * * * *`, or `TZ=America/New_York 1 * * * *`.
	// +kcc:proto:field=google.cloud.dataplex.v1.Zone.DiscoverySpec.schedule
	Schedule *string `json:"schedule,omitempty"`
}

// +kcc:proto=google.cloud.dataplex.v1.Zone.DiscoverySpec.CsvOptions
type Zone_DiscoverySpec_CsvOptions struct {
	// Optional. The number of rows to interpret as header rows that should be
	//  skipped when reading data rows.
	// +kcc:proto:field=google.cloud.dataplex.v1.Zone.DiscoverySpec.CsvOptions.header_rows
	HeaderRows *int32 `json:"headerRows,omitempty"`

	// Optional. The delimiter being used to separate values. This defaults to
	//  ','.
	// +kcc:proto:field=google.cloud.dataplex.v1.Zone.DiscoverySpec.CsvOptions.delimiter
	Delimiter *string `json:"delimiter,omitempty"`

	// Optional. The character encoding of the data. The default is UTF-8.
	// +kcc:proto:field=google.cloud.dataplex.v1.Zone.DiscoverySpec.CsvOptions.encoding
	Encoding *string `json:"encoding,omitempty"`

	// Optional. Whether to disable the inference of data type for CSV data.
	//  If true, all columns will be registered as strings.
	// +kcc:proto:field=google.cloud.dataplex.v1.Zone.DiscoverySpec.CsvOptions.disable_type_inference
	DisableTypeInference *bool `json:"disableTypeInference,omitempty"`
}

// +kcc:proto=google.cloud.dataplex.v1.Zone.DiscoverySpec.JsonOptions
type Zone_DiscoverySpec_JsonOptions struct {
	// Optional. The character encoding of the data. The default is UTF-8.
	// +kcc:proto:field=google.cloud.dataplex.v1.Zone.DiscoverySpec.JsonOptions.encoding
	Encoding *string `json:"encoding,omitempty"`

	// Optional. Whether to disable the inference of data type for Json data.
	//  If true, all columns will be registered as their primitive types
	//  (strings, number or boolean).
	// +kcc:proto:field=google.cloud.dataplex.v1.Zone.DiscoverySpec.JsonOptions.disable_type_inference
	DisableTypeInference *bool `json:"disableTypeInference,omitempty"`
}

// +kcc:proto=google.cloud.dataplex.v1.Zone.ResourceSpec
type Zone_ResourceSpec struct {
	// Required. Immutable. The location type of the resources that are allowed
	//  to be attached to the assets within this zone.
	// +kcc:proto:field=google.cloud.dataplex.v1.Zone.ResourceSpec.location_type
	LocationType *string `json:"locationType,omitempty"`
}

// +kcc:proto=google.cloud.dataplex.v1.Zone
type ZoneObservedState struct {
	// Output only. The relative resource name of the zone, of the form:
	//  `projects/{project_number}/locations/{location_id}/lakes/{lake_id}/zones/{zone_id}`.
	// +kcc:proto:field=google.cloud.dataplex.v1.Zone.name
	Name *string `json:"name,omitempty"`

	// Output only. System generated globally unique ID for the zone. This ID will
	//  be different if the zone is deleted and re-created with the same name.
	// +kcc:proto:field=google.cloud.dataplex.v1.Zone.uid
	Uid *string `json:"uid,omitempty"`

	// Output only. The time when the zone was created.
	// +kcc:proto:field=google.cloud.dataplex.v1.Zone.create_time
	CreateTime *string `json:"createTime,omitempty"`

	// Output only. The time when the zone was last updated.
	// +kcc:proto:field=google.cloud.dataplex.v1.Zone.update_time
	UpdateTime *string `json:"updateTime,omitempty"`

	// Output only. Current state of the zone.
	// +kcc:proto:field=google.cloud.dataplex.v1.Zone.state
	State *string `json:"state,omitempty"`

	// Output only. Aggregated status of the underlying assets of the zone.
	// +kcc:proto:field=google.cloud.dataplex.v1.Zone.asset_status
	AssetStatus *AssetStatus `json:"assetStatus,omitempty"`
}
