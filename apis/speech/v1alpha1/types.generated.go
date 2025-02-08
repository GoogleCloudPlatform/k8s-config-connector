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


// +kcc:proto=google.cloud.speech.v2.Config
type Config struct {

	// Optional. An optional [KMS key
	//  name](https://cloud.google.com/kms/docs/resource-hierarchy#keys) that if
	//  present, will be used to encrypt Speech-to-Text resources at-rest. Updating
	//  this key will not encrypt existing resources using this key; only new
	//  resources will be encrypted using this key. The expected format is
	//  `projects/{project}/locations/{location}/keyRings/{key_ring}/cryptoKeys/{crypto_key}`.
	// +kcc:proto:field=google.cloud.speech.v2.Config.kms_key_name
	KMSKeyName *string `json:"kmsKeyName,omitempty"`
}

// +kcc:proto=google.cloud.speech.v2.Config
type ConfigObservedState struct {
	// Output only. Identifier. The name of the config resource. There is exactly
	//  one config resource per project per location. The expected format is
	//  `projects/{project}/locations/{location}/config`.
	// +kcc:proto:field=google.cloud.speech.v2.Config.name
	Name *string `json:"name,omitempty"`

	// Output only. The most recent time this resource was modified.
	// +kcc:proto:field=google.cloud.speech.v2.Config.update_time
	UpdateTime *string `json:"updateTime,omitempty"`
}
