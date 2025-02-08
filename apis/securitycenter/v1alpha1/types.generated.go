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


// +kcc:proto=google.cloud.securitycenter.v1.MuteConfig
type MuteConfig struct {
	// This field will be ignored if provided on config creation. Format
	//  `organizations/{organization}/muteConfigs/{mute_config}`
	//  `folders/{folder}/muteConfigs/{mute_config}`
	//  `projects/{project}/muteConfigs/{mute_config}`
	//  `organizations/{organization}/locations/global/muteConfigs/{mute_config}`
	//  `folders/{folder}/locations/global/muteConfigs/{mute_config}`
	//  `projects/{project}/locations/global/muteConfigs/{mute_config}`
	// +kcc:proto:field=google.cloud.securitycenter.v1.MuteConfig.name
	Name *string `json:"name,omitempty"`

	// The human readable name to be displayed for the mute config.
	// +kcc:proto:field=google.cloud.securitycenter.v1.MuteConfig.display_name
	DisplayName *string `json:"displayName,omitempty"`

	// A description of the mute config.
	// +kcc:proto:field=google.cloud.securitycenter.v1.MuteConfig.description
	Description *string `json:"description,omitempty"`

	// Required. An expression that defines the filter to apply across
	//  create/update events of findings. While creating a filter string, be
	//  mindful of the scope in which the mute configuration is being created.
	//  E.g., If a filter contains project = X but is created under the project = Y
	//  scope, it might not match any findings.
	//
	//  The following field and operator combinations are supported:
	//
	//  * severity: `=`, `:`
	//  * category: `=`, `:`
	//  * resource.name: `=`, `:`
	//  * resource.project_name: `=`, `:`
	//  * resource.project_display_name: `=`, `:`
	//  * resource.folders.resource_folder: `=`, `:`
	//  * resource.parent_name: `=`, `:`
	//  * resource.parent_display_name: `=`, `:`
	//  * resource.type: `=`, `:`
	//  * finding_class: `=`, `:`
	//  * indicator.ip_addresses: `=`, `:`
	//  * indicator.domains: `=`, `:`
	// +kcc:proto:field=google.cloud.securitycenter.v1.MuteConfig.filter
	Filter *string `json:"filter,omitempty"`

	// Optional. The type of the mute config, which determines what type of mute
	//  state the config affects. The static mute state takes precedence over the
	//  dynamic mute state. Immutable after creation. STATIC by default if not set
	//  during creation.
	// +kcc:proto:field=google.cloud.securitycenter.v1.MuteConfig.type
	Type *string `json:"type,omitempty"`

	// Optional. The expiry of the mute config. Only applicable for dynamic
	//  configs. If the expiry is set, when the config expires, it is removed from
	//  all findings.
	// +kcc:proto:field=google.cloud.securitycenter.v1.MuteConfig.expiry_time
	ExpiryTime *string `json:"expiryTime,omitempty"`
}

// +kcc:proto=google.cloud.securitycenter.v1.MuteConfig
type MuteConfigObservedState struct {
	// Output only. The time at which the mute config was created.
	//  This field is set by the server and will be ignored if provided on config
	//  creation.
	// +kcc:proto:field=google.cloud.securitycenter.v1.MuteConfig.create_time
	CreateTime *string `json:"createTime,omitempty"`

	// Output only. The most recent time at which the mute config was updated.
	//  This field is set by the server and will be ignored if provided on config
	//  creation or update.
	// +kcc:proto:field=google.cloud.securitycenter.v1.MuteConfig.update_time
	UpdateTime *string `json:"updateTime,omitempty"`

	// Output only. Email address of the user who last edited the mute config.
	//  This field is set by the server and will be ignored if provided on config
	//  creation or update.
	// +kcc:proto:field=google.cloud.securitycenter.v1.MuteConfig.most_recent_editor
	MostRecentEditor *string `json:"mostRecentEditor,omitempty"`
}
