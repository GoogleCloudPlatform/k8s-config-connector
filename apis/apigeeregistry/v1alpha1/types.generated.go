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


// +kcc:proto=google.cloud.apigeeregistry.v1.Instance
type Instance struct {
	// Format: `projects/*/locations/*/instance`.
	//  Currently only `locations/global` is supported.
	// +kcc:proto:field=google.cloud.apigeeregistry.v1.Instance.name
	Name *string `json:"name,omitempty"`

	// Required. Config of the Instance.
	// +kcc:proto:field=google.cloud.apigeeregistry.v1.Instance.config
	Config *Instance_Config `json:"config,omitempty"`
}

// +kcc:proto=google.cloud.apigeeregistry.v1.Instance.Config
type Instance_Config struct {

	// Required. The Customer Managed Encryption Key (CMEK) used for data encryption.
	//  The CMEK name should follow the format of
	//  `projects/([^/]+)/locations/([^/]+)/keyRings/([^/]+)/cryptoKeys/([^/]+)`,
	//  where the `location` must match InstanceConfig.location.
	// +kcc:proto:field=google.cloud.apigeeregistry.v1.Instance.Config.cmek_key_name
	CmekKeyName *string `json:"cmekKeyName,omitempty"`
}

// +kcc:proto=google.cloud.apigeeregistry.v1.Instance
type InstanceObservedState struct {
	// Output only. Creation timestamp.
	// +kcc:proto:field=google.cloud.apigeeregistry.v1.Instance.create_time
	CreateTime *string `json:"createTime,omitempty"`

	// Output only. Last update timestamp.
	// +kcc:proto:field=google.cloud.apigeeregistry.v1.Instance.update_time
	UpdateTime *string `json:"updateTime,omitempty"`

	// Output only. The current state of the Instance.
	// +kcc:proto:field=google.cloud.apigeeregistry.v1.Instance.state
	State *string `json:"state,omitempty"`

	// Output only. Extra information of Instance.State if the state is `FAILED`.
	// +kcc:proto:field=google.cloud.apigeeregistry.v1.Instance.state_message
	StateMessage *string `json:"stateMessage,omitempty"`

	// Required. Config of the Instance.
	// +kcc:proto:field=google.cloud.apigeeregistry.v1.Instance.config
	Config *Instance_ConfigObservedState `json:"config,omitempty"`
}

// +kcc:proto=google.cloud.apigeeregistry.v1.Instance.Config
type Instance_ConfigObservedState struct {
	// Output only. The GCP location where the Instance resides.
	// +kcc:proto:field=google.cloud.apigeeregistry.v1.Instance.Config.location
	Location *string `json:"location,omitempty"`
}
