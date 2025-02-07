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


// +kcc:proto=google.cloud.apihub.v1.ApiHubInstance
type ApiHubInstance struct {
	// Identifier. Format:
	//  `projects/{project}/locations/{location}/apiHubInstances/{apiHubInstance}`.
	// +kcc:proto:field=google.cloud.apihub.v1.ApiHubInstance.name
	Name *string `json:"name,omitempty"`

	// Required. Config of the ApiHub instance.
	// +kcc:proto:field=google.cloud.apihub.v1.ApiHubInstance.config
	Config *ApiHubInstance_Config `json:"config,omitempty"`

	// Optional. Instance labels to represent user-provided metadata.
	//  Refer to cloud documentation on labels for more details.
	//  https://cloud.google.com/compute/docs/labeling-resources
	// +kcc:proto:field=google.cloud.apihub.v1.ApiHubInstance.labels
	Labels map[string]string `json:"labels,omitempty"`

	// Optional. Description of the ApiHub instance.
	// +kcc:proto:field=google.cloud.apihub.v1.ApiHubInstance.description
	Description *string `json:"description,omitempty"`
}

// +kcc:proto=google.cloud.apihub.v1.ApiHubInstance.Config
type ApiHubInstance_Config struct {
	// Required. The Customer Managed Encryption Key (CMEK) used for data
	//  encryption. The CMEK name should follow the format of
	//  `projects/([^/]+)/locations/([^/]+)/keyRings/([^/]+)/cryptoKeys/([^/]+)`,
	//  where the location must match the instance location.
	// +kcc:proto:field=google.cloud.apihub.v1.ApiHubInstance.Config.cmek_key_name
	CmekKeyName *string `json:"cmekKeyName,omitempty"`
}

// +kcc:proto=google.cloud.apihub.v1.ApiHubInstance
type ApiHubInstanceObservedState struct {
	// Output only. Creation timestamp.
	// +kcc:proto:field=google.cloud.apihub.v1.ApiHubInstance.create_time
	CreateTime *string `json:"createTime,omitempty"`

	// Output only. Last update timestamp.
	// +kcc:proto:field=google.cloud.apihub.v1.ApiHubInstance.update_time
	UpdateTime *string `json:"updateTime,omitempty"`

	// Output only. The current state of the ApiHub instance.
	// +kcc:proto:field=google.cloud.apihub.v1.ApiHubInstance.state
	State *string `json:"state,omitempty"`

	// Output only. Extra information about ApiHub instance state. Currently the
	//  message would be populated when state is `FAILED`.
	// +kcc:proto:field=google.cloud.apihub.v1.ApiHubInstance.state_message
	StateMessage *string `json:"stateMessage,omitempty"`
}
