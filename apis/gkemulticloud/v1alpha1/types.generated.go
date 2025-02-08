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


// +kcc:proto=google.cloud.gkemulticloud.v1.AttachedPlatformVersionInfo
type AttachedPlatformVersionInfo struct {
	// Platform version name.
	// +kcc:proto:field=google.cloud.gkemulticloud.v1.AttachedPlatformVersionInfo.version
	Version *string `json:"version,omitempty"`
}

// +kcc:proto=google.cloud.gkemulticloud.v1.AttachedServerConfig
type AttachedServerConfig struct {
	// The resource name of the config.
	// +kcc:proto:field=google.cloud.gkemulticloud.v1.AttachedServerConfig.name
	Name *string `json:"name,omitempty"`

	// List of valid platform versions.
	// +kcc:proto:field=google.cloud.gkemulticloud.v1.AttachedServerConfig.valid_versions
	ValidVersions []AttachedPlatformVersionInfo `json:"validVersions,omitempty"`
}
