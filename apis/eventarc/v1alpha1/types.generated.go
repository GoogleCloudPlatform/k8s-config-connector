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

// +generated:types
// krm.group: eventarc.cnrm.cloud.google.com
// krm.version: v1alpha1
// proto.service: google.cloud.eventarc.v1
// resource: EventarcGoogleChannelConfig:GoogleChannelConfig

package v1alpha1

// +kcc:proto=google.cloud.eventarc.v1.GoogleChannelConfig
type GoogleChannelConfig struct {
	// +required
	// Required. The resource name of the config. Must be in the format of,
	//  `projects/{project}/locations/{location}/googleChannelConfig`.
	// +kcc:proto:field=google.cloud.eventarc.v1.GoogleChannelConfig.name
	Name *string `json:"name,omitempty"`

	// Optional. Resource name of a KMS crypto key (managed by the user) used to
	//  encrypt/decrypt their event data.
	//
	//  It must match the pattern
	//  `projects/*/locations/*/keyRings/*/cryptoKeys/*`.
	// +kcc:proto:field=google.cloud.eventarc.v1.GoogleChannelConfig.crypto_key_name
	CryptoKeyName *string `json:"cryptoKeyName,omitempty"`
}

// +kcc:proto=google.cloud.eventarc.v1.GoogleChannelConfig
type GoogleChannelConfigObservedState struct {
	// Output only. The last-modified time.
	// +kcc:proto:field=google.cloud.eventarc.v1.GoogleChannelConfig.update_time
	UpdateTime *string `json:"updateTime,omitempty"`
}
