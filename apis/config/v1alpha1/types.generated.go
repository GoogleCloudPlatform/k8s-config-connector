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


// +kcc:proto=google.cloud.config.v1.TerraformVersion
type TerraformVersion struct {
	// Identifier. The version name is in the format:
	//  'projects/{project_id}/locations/{location}/terraformVersions/{terraform_version}'.
	// +kcc:proto:field=google.cloud.config.v1.TerraformVersion.name
	Name *string `json:"name,omitempty"`
}

// +kcc:proto=google.cloud.config.v1.TerraformVersion
type TerraformVersionObservedState struct {
	// Output only. The state of the version, ACTIVE, DEPRECATED or OBSOLETE.
	// +kcc:proto:field=google.cloud.config.v1.TerraformVersion.state
	State *string `json:"state,omitempty"`

	// Output only. When the version is supported.
	// +kcc:proto:field=google.cloud.config.v1.TerraformVersion.support_time
	SupportTime *string `json:"supportTime,omitempty"`

	// Output only. When the version is deprecated.
	// +kcc:proto:field=google.cloud.config.v1.TerraformVersion.deprecate_time
	DeprecateTime *string `json:"deprecateTime,omitempty"`

	// Output only. When the version is obsolete.
	// +kcc:proto:field=google.cloud.config.v1.TerraformVersion.obsolete_time
	ObsoleteTime *string `json:"obsoleteTime,omitempty"`
}
