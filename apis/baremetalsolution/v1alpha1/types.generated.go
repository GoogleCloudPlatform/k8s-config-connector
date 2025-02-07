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


// +kcc:proto=google.cloud.baremetalsolution.v2.OSImage
type OSImage struct {

	// OS Image code.
	// +kcc:proto:field=google.cloud.baremetalsolution.v2.OSImage.code
	Code *string `json:"code,omitempty"`

	// OS Image description.
	// +kcc:proto:field=google.cloud.baremetalsolution.v2.OSImage.description
	Description *string `json:"description,omitempty"`

	// Instance types this image is applicable to.
	//  [Available
	//  types](https://cloud.google.com/bare-metal/docs/bms-planning#server_configurations)
	// +kcc:proto:field=google.cloud.baremetalsolution.v2.OSImage.applicable_instance_types
	ApplicableInstanceTypes []string `json:"applicableInstanceTypes,omitempty"`

	// Network templates that can be used with this OS Image.
	// +kcc:proto:field=google.cloud.baremetalsolution.v2.OSImage.supported_network_templates
	SupportedNetworkTemplates []string `json:"supportedNetworkTemplates,omitempty"`
}

// +kcc:proto=google.cloud.baremetalsolution.v2.OSImage
type OSImageObservedState struct {
	// Output only. OS Image's unique name.
	// +kcc:proto:field=google.cloud.baremetalsolution.v2.OSImage.name
	Name *string `json:"name,omitempty"`
}
