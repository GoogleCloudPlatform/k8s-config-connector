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


// +kcc:proto=google.cloud.kms.v1.EkmConfig
type EkmConfig struct {

	// Optional. Resource name of the default
	//  [EkmConnection][google.cloud.kms.v1.EkmConnection]. Setting this field to
	//  the empty string removes the default.
	// +kcc:proto:field=google.cloud.kms.v1.EkmConfig.default_ekm_connection
	DefaultEkmConnection *string `json:"defaultEkmConnection,omitempty"`
}

// +kcc:proto=google.cloud.kms.v1.EkmConfig
type EkmConfigObservedState struct {
	// Output only. The resource name for the
	//  [EkmConfig][google.cloud.kms.v1.EkmConfig] in the format
	//  `projects/*/locations/*/ekmConfig`.
	// +kcc:proto:field=google.cloud.kms.v1.EkmConfig.name
	Name *string `json:"name,omitempty"`
}
