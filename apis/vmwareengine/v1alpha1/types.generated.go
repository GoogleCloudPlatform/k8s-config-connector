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


// +kcc:proto=google.cloud.vmwareengine.v1.HcxActivationKey
type HcxActivationKey struct {
}

// +kcc:proto=google.cloud.vmwareengine.v1.HcxActivationKey
type HcxActivationKeyObservedState struct {
	// Output only. The resource name of this HcxActivationKey.
	//  Resource names are schemeless URIs that follow the conventions in
	//  https://cloud.google.com/apis/design/resource_names.
	//  For example:
	//  `projects/my-project/locations/us-central1/privateClouds/my-cloud/hcxActivationKeys/my-key`
	// +kcc:proto:field=google.cloud.vmwareengine.v1.HcxActivationKey.name
	Name *string `json:"name,omitempty"`

	// Output only. Creation time of HCX activation key.
	// +kcc:proto:field=google.cloud.vmwareengine.v1.HcxActivationKey.create_time
	CreateTime *string `json:"createTime,omitempty"`

	// Output only. State of HCX activation key.
	// +kcc:proto:field=google.cloud.vmwareengine.v1.HcxActivationKey.state
	State *string `json:"state,omitempty"`

	// Output only. HCX activation key.
	// +kcc:proto:field=google.cloud.vmwareengine.v1.HcxActivationKey.activation_key
	ActivationKey *string `json:"activationKey,omitempty"`

	// Output only. System-generated unique identifier for the resource.
	// +kcc:proto:field=google.cloud.vmwareengine.v1.HcxActivationKey.uid
	Uid *string `json:"uid,omitempty"`
}
