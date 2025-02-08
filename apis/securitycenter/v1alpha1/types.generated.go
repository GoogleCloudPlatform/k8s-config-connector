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


// +kcc:proto=google.cloud.securitycenter.v1p1beta1.OrganizationSettings
type OrganizationSettings struct {
	// The relative resource name of the settings. See:
	//  https://cloud.google.com/apis/design/resource_names#relative_resource_name
	//  Example:
	//  "organizations/{organization_id}/organizationSettings".
	// +kcc:proto:field=google.cloud.securitycenter.v1p1beta1.OrganizationSettings.name
	Name *string `json:"name,omitempty"`

	// A flag that indicates if Asset Discovery should be enabled. If the flag is
	//  set to `true`, then discovery of assets will occur. If it is set to `false,
	//  all historical assets will remain, but discovery of future assets will not
	//  occur.
	// +kcc:proto:field=google.cloud.securitycenter.v1p1beta1.OrganizationSettings.enable_asset_discovery
	EnableAssetDiscovery *bool `json:"enableAssetDiscovery,omitempty"`

	// The configuration used for Asset Discovery runs.
	// +kcc:proto:field=google.cloud.securitycenter.v1p1beta1.OrganizationSettings.asset_discovery_config
	AssetDiscoveryConfig *OrganizationSettings_AssetDiscoveryConfig `json:"assetDiscoveryConfig,omitempty"`
}

// +kcc:proto=google.cloud.securitycenter.v1p1beta1.OrganizationSettings.AssetDiscoveryConfig
type OrganizationSettings_AssetDiscoveryConfig struct {
	// The project ids to use for filtering asset discovery.
	// +kcc:proto:field=google.cloud.securitycenter.v1p1beta1.OrganizationSettings.AssetDiscoveryConfig.project_ids
	ProjectIds []string `json:"projectIds,omitempty"`

	// The mode to use for filtering asset discovery.
	// +kcc:proto:field=google.cloud.securitycenter.v1p1beta1.OrganizationSettings.AssetDiscoveryConfig.inclusion_mode
	InclusionMode *string `json:"inclusionMode,omitempty"`

	// The folder ids to use for filtering asset discovery.
	//  It consists of only digits, e.g., 756619654966.
	// +kcc:proto:field=google.cloud.securitycenter.v1p1beta1.OrganizationSettings.AssetDiscoveryConfig.folder_ids
	FolderIds []string `json:"folderIds,omitempty"`
}
