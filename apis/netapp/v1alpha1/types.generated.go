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


// +kcc:proto=google.cloud.netapp.v1.KmsConfig
type KmsConfig struct {
	// Identifier. Name of the KmsConfig.
	// +kcc:proto:field=google.cloud.netapp.v1.KmsConfig.name
	Name *string `json:"name,omitempty"`

	// Required. Customer managed crypto key resource full name. Format:
	//  projects/{project}/locations/{location}/keyRings/{key_ring}/cryptoKeys/{key}.
	// +kcc:proto:field=google.cloud.netapp.v1.KmsConfig.crypto_key_name
	CryptoKeyName *string `json:"cryptoKeyName,omitempty"`

	// Description of the KmsConfig.
	// +kcc:proto:field=google.cloud.netapp.v1.KmsConfig.description
	Description *string `json:"description,omitempty"`

	// Labels as key value pairs
	// +kcc:proto:field=google.cloud.netapp.v1.KmsConfig.labels
	Labels map[string]string `json:"labels,omitempty"`
}

// +kcc:proto=google.cloud.netapp.v1.KmsConfig
type KmsConfigObservedState struct {
	// Output only. State of the KmsConfig.
	// +kcc:proto:field=google.cloud.netapp.v1.KmsConfig.state
	State *string `json:"state,omitempty"`

	// Output only. State details of the KmsConfig.
	// +kcc:proto:field=google.cloud.netapp.v1.KmsConfig.state_details
	StateDetails *string `json:"stateDetails,omitempty"`

	// Output only. Create time of the KmsConfig.
	// +kcc:proto:field=google.cloud.netapp.v1.KmsConfig.create_time
	CreateTime *string `json:"createTime,omitempty"`

	// Output only. Instructions to provide the access to the customer provided
	//  encryption key.
	// +kcc:proto:field=google.cloud.netapp.v1.KmsConfig.instructions
	Instructions *string `json:"instructions,omitempty"`

	// Output only. The Service account which will have access to the customer
	//  provided encryption key.
	// +kcc:proto:field=google.cloud.netapp.v1.KmsConfig.service_account
	ServiceAccount *string `json:"serviceAccount,omitempty"`
}
