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


// +kcc:proto=google.cloud.dataplex.v1.Asset
type Asset struct {

	// Optional. User friendly display name.
	// +kcc:proto:field=google.cloud.dataplex.v1.Asset.display_name
	DisplayName *string `json:"displayName,omitempty"`

	// Optional. User defined labels for the asset.
	// +kcc:proto:field=google.cloud.dataplex.v1.Asset.labels
	Labels map[string]string `json:"labels,omitempty"`

	// Optional. Description of the asset.
	// +kcc:proto:field=google.cloud.dataplex.v1.Asset.description
	Description *string `json:"description,omitempty"`

	// Required. Specification of the resource that is referenced by this asset.
	// +kcc:proto:field=google.cloud.dataplex.v1.Asset.resource_spec
	ResourceSpec *Asset_ResourceSpec `json:"resourceSpec,omitempty"`

	// Optional. Specification of the discovery feature applied to data referenced
	//  by this asset. When this spec is left unset, the asset will use the spec
	//  set on the parent zone.
	// +kcc:proto:field=google.cloud.dataplex.v1.Asset.discovery_spec
	DiscoverySpec *Asset_DiscoverySpec `json:"discoverySpec,omitempty"`
}

// +kcc:proto=google.cloud.dataplex.v1.Asset.DiscoverySpec
type Asset_DiscoverySpec struct {
	// Optional. Whether discovery is enabled.
	// +kcc:proto:field=google.cloud.dataplex.v1.Asset.DiscoverySpec.enabled
	Enabled *bool `json:"enabled,omitempty"`

	// Optional. The list of patterns to apply for selecting data to include
	//  during discovery if only a subset of the data should considered.  For
	//  Cloud Storage bucket assets, these are interpreted as glob patterns used
	//  to match object names. For BigQuery dataset assets, these are interpreted
	//  as patterns to match table names.
	// +kcc:proto:field=google.cloud.dataplex.v1.Asset.DiscoverySpec.include_patterns
	IncludePatterns []string `json:"includePatterns,omitempty"`

	// Optional. The list of patterns to apply for selecting data to exclude
	//  during discovery.  For Cloud Storage bucket assets, these are interpreted
	//  as glob patterns used to match object names. For BigQuery dataset assets,
	//  these are interpreted as patterns to match table names.
	// +kcc:proto:field=google.cloud.dataplex.v1.Asset.DiscoverySpec.exclude_patterns
	ExcludePatterns []string `json:"excludePatterns,omitempty"`

	// Optional. Configuration for CSV data.
	// +kcc:proto:field=google.cloud.dataplex.v1.Asset.DiscoverySpec.csv_options
	CsvOptions *Asset_DiscoverySpec_CsvOptions `json:"csvOptions,omitempty"`

	// Optional. Configuration for Json data.
	// +kcc:proto:field=google.cloud.dataplex.v1.Asset.DiscoverySpec.json_options
	JsonOptions *Asset_DiscoverySpec_JsonOptions `json:"jsonOptions,omitempty"`

	// Optional. Cron schedule (https://en.wikipedia.org/wiki/Cron) for
	//  running discovery periodically. Successive discovery runs must be
	//  scheduled at least 60 minutes apart. The default value is to run
	//  discovery every 60 minutes. To explicitly set a timezone to the cron
	//  tab, apply a prefix in the cron tab: "CRON_TZ=${IANA_TIME_ZONE}" or
	//  TZ=${IANA_TIME_ZONE}". The ${IANA_TIME_ZONE} may only be a valid string
	//  from IANA time zone database. For example, `CRON_TZ=America/New_York 1
	//  * * * *`, or `TZ=America/New_York 1 * * * *`.
	// +kcc:proto:field=google.cloud.dataplex.v1.Asset.DiscoverySpec.schedule
	Schedule *string `json:"schedule,omitempty"`
}

// +kcc:proto=google.cloud.dataplex.v1.Asset.DiscoverySpec.CsvOptions
type Asset_DiscoverySpec_CsvOptions struct {
	// Optional. The number of rows to interpret as header rows that should be
	//  skipped when reading data rows.
	// +kcc:proto:field=google.cloud.dataplex.v1.Asset.DiscoverySpec.CsvOptions.header_rows
	HeaderRows *int32 `json:"headerRows,omitempty"`

	// Optional. The delimiter being used to separate values. This defaults to
	//  ','.
	// +kcc:proto:field=google.cloud.dataplex.v1.Asset.DiscoverySpec.CsvOptions.delimiter
	Delimiter *string `json:"delimiter,omitempty"`

	// Optional. The character encoding of the data. The default is UTF-8.
	// +kcc:proto:field=google.cloud.dataplex.v1.Asset.DiscoverySpec.CsvOptions.encoding
	Encoding *string `json:"encoding,omitempty"`

	// Optional. Whether to disable the inference of data type for CSV data.
	//  If true, all columns will be registered as strings.
	// +kcc:proto:field=google.cloud.dataplex.v1.Asset.DiscoverySpec.CsvOptions.disable_type_inference
	DisableTypeInference *bool `json:"disableTypeInference,omitempty"`
}

// +kcc:proto=google.cloud.dataplex.v1.Asset.DiscoverySpec.JsonOptions
type Asset_DiscoverySpec_JsonOptions struct {
	// Optional. The character encoding of the data. The default is UTF-8.
	// +kcc:proto:field=google.cloud.dataplex.v1.Asset.DiscoverySpec.JsonOptions.encoding
	Encoding *string `json:"encoding,omitempty"`

	// Optional. Whether to disable the inference of data type for Json data.
	//  If true, all columns will be registered as their primitive types
	//  (strings, number or boolean).
	// +kcc:proto:field=google.cloud.dataplex.v1.Asset.DiscoverySpec.JsonOptions.disable_type_inference
	DisableTypeInference *bool `json:"disableTypeInference,omitempty"`
}

// +kcc:proto=google.cloud.dataplex.v1.Asset.DiscoveryStatus
type Asset_DiscoveryStatus struct {
	// The current status of the discovery feature.
	// +kcc:proto:field=google.cloud.dataplex.v1.Asset.DiscoveryStatus.state
	State *string `json:"state,omitempty"`

	// Additional information about the current state.
	// +kcc:proto:field=google.cloud.dataplex.v1.Asset.DiscoveryStatus.message
	Message *string `json:"message,omitempty"`

	// Last update time of the status.
	// +kcc:proto:field=google.cloud.dataplex.v1.Asset.DiscoveryStatus.update_time
	UpdateTime *string `json:"updateTime,omitempty"`

	// The start time of the last discovery run.
	// +kcc:proto:field=google.cloud.dataplex.v1.Asset.DiscoveryStatus.last_run_time
	LastRunTime *string `json:"lastRunTime,omitempty"`

	// Data Stats of the asset reported by discovery.
	// +kcc:proto:field=google.cloud.dataplex.v1.Asset.DiscoveryStatus.stats
	Stats *Asset_DiscoveryStatus_Stats `json:"stats,omitempty"`

	// The duration of the last discovery run.
	// +kcc:proto:field=google.cloud.dataplex.v1.Asset.DiscoveryStatus.last_run_duration
	LastRunDuration *string `json:"lastRunDuration,omitempty"`
}

// +kcc:proto=google.cloud.dataplex.v1.Asset.DiscoveryStatus.Stats
type Asset_DiscoveryStatus_Stats struct {
	// The count of data items within the referenced resource.
	// +kcc:proto:field=google.cloud.dataplex.v1.Asset.DiscoveryStatus.Stats.data_items
	DataItems *int64 `json:"dataItems,omitempty"`

	// The number of stored data bytes within the referenced resource.
	// +kcc:proto:field=google.cloud.dataplex.v1.Asset.DiscoveryStatus.Stats.data_size
	DataSize *int64 `json:"dataSize,omitempty"`

	// The count of table entities within the referenced resource.
	// +kcc:proto:field=google.cloud.dataplex.v1.Asset.DiscoveryStatus.Stats.tables
	Tables *int64 `json:"tables,omitempty"`

	// The count of fileset entities within the referenced resource.
	// +kcc:proto:field=google.cloud.dataplex.v1.Asset.DiscoveryStatus.Stats.filesets
	Filesets *int64 `json:"filesets,omitempty"`
}

// +kcc:proto=google.cloud.dataplex.v1.Asset.ResourceSpec
type Asset_ResourceSpec struct {
	// Immutable. Relative name of the cloud resource that contains the data
	//  that is being managed within a lake. For example:
	//    `projects/{project_number}/buckets/{bucket_id}`
	//    `projects/{project_number}/datasets/{dataset_id}`
	// +kcc:proto:field=google.cloud.dataplex.v1.Asset.ResourceSpec.name
	Name *string `json:"name,omitempty"`

	// Required. Immutable. Type of resource.
	// +kcc:proto:field=google.cloud.dataplex.v1.Asset.ResourceSpec.type
	Type *string `json:"type,omitempty"`

	// Optional. Determines how read permissions are handled for each asset and
	//  their associated tables. Only available to storage buckets assets.
	// +kcc:proto:field=google.cloud.dataplex.v1.Asset.ResourceSpec.read_access_mode
	ReadAccessMode *string `json:"readAccessMode,omitempty"`
}

// +kcc:proto=google.cloud.dataplex.v1.Asset.ResourceStatus
type Asset_ResourceStatus struct {
	// The current state of the managed resource.
	// +kcc:proto:field=google.cloud.dataplex.v1.Asset.ResourceStatus.state
	State *string `json:"state,omitempty"`

	// Additional information about the current state.
	// +kcc:proto:field=google.cloud.dataplex.v1.Asset.ResourceStatus.message
	Message *string `json:"message,omitempty"`

	// Last update time of the status.
	// +kcc:proto:field=google.cloud.dataplex.v1.Asset.ResourceStatus.update_time
	UpdateTime *string `json:"updateTime,omitempty"`
}

// +kcc:proto=google.cloud.dataplex.v1.Asset.SecurityStatus
type Asset_SecurityStatus struct {
	// The current state of the security policy applied to the attached
	//  resource.
	// +kcc:proto:field=google.cloud.dataplex.v1.Asset.SecurityStatus.state
	State *string `json:"state,omitempty"`

	// Additional information about the current state.
	// +kcc:proto:field=google.cloud.dataplex.v1.Asset.SecurityStatus.message
	Message *string `json:"message,omitempty"`

	// Last update time of the status.
	// +kcc:proto:field=google.cloud.dataplex.v1.Asset.SecurityStatus.update_time
	UpdateTime *string `json:"updateTime,omitempty"`
}

// +kcc:proto=google.cloud.dataplex.v1.Asset
type AssetObservedState struct {
	// Output only. The relative resource name of the asset, of the form:
	//  `projects/{project_number}/locations/{location_id}/lakes/{lake_id}/zones/{zone_id}/assets/{asset_id}`.
	// +kcc:proto:field=google.cloud.dataplex.v1.Asset.name
	Name *string `json:"name,omitempty"`

	// Output only. System generated globally unique ID for the asset. This ID
	//  will be different if the asset is deleted and re-created with the same
	//  name.
	// +kcc:proto:field=google.cloud.dataplex.v1.Asset.uid
	Uid *string `json:"uid,omitempty"`

	// Output only. The time when the asset was created.
	// +kcc:proto:field=google.cloud.dataplex.v1.Asset.create_time
	CreateTime *string `json:"createTime,omitempty"`

	// Output only. The time when the asset was last updated.
	// +kcc:proto:field=google.cloud.dataplex.v1.Asset.update_time
	UpdateTime *string `json:"updateTime,omitempty"`

	// Output only. Current state of the asset.
	// +kcc:proto:field=google.cloud.dataplex.v1.Asset.state
	State *string `json:"state,omitempty"`

	// Output only. Status of the resource referenced by this asset.
	// +kcc:proto:field=google.cloud.dataplex.v1.Asset.resource_status
	ResourceStatus *Asset_ResourceStatus `json:"resourceStatus,omitempty"`

	// Output only. Status of the security policy applied to resource referenced
	//  by this asset.
	// +kcc:proto:field=google.cloud.dataplex.v1.Asset.security_status
	SecurityStatus *Asset_SecurityStatus `json:"securityStatus,omitempty"`

	// Output only. Status of the discovery feature applied to data referenced by
	//  this asset.
	// +kcc:proto:field=google.cloud.dataplex.v1.Asset.discovery_status
	DiscoveryStatus *Asset_DiscoveryStatus `json:"discoveryStatus,omitempty"`
}

// +kcc:proto=google.cloud.dataplex.v1.Asset.ResourceStatus
type Asset_ResourceStatusObservedState struct {
	// Output only. Service account associated with the BigQuery Connection.
	// +kcc:proto:field=google.cloud.dataplex.v1.Asset.ResourceStatus.managed_access_identity
	ManagedAccessIdentity *string `json:"managedAccessIdentity,omitempty"`
}
