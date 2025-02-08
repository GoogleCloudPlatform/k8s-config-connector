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


// +kcc:proto=google.cloud.vmwareengine.v1.ExternalAddress
type ExternalAddress struct {

	// The internal IP address of a workload VM.
	// +kcc:proto:field=google.cloud.vmwareengine.v1.ExternalAddress.internal_ip
	InternalIP *string `json:"internalIP,omitempty"`

	// User-provided description for this resource.
	// +kcc:proto:field=google.cloud.vmwareengine.v1.ExternalAddress.description
	Description *string `json:"description,omitempty"`
}

// +kcc:proto=google.cloud.vmwareengine.v1.ExternalAddress
type ExternalAddressObservedState struct {
	// Output only. The resource name of this external IP address.
	//  Resource names are schemeless URIs that follow the conventions in
	//  https://cloud.google.com/apis/design/resource_names.
	//  For example:
	//  `projects/my-project/locations/us-central1-a/privateClouds/my-cloud/externalAddresses/my-address`
	// +kcc:proto:field=google.cloud.vmwareengine.v1.ExternalAddress.name
	Name *string `json:"name,omitempty"`

	// Output only. Creation time of this resource.
	// +kcc:proto:field=google.cloud.vmwareengine.v1.ExternalAddress.create_time
	CreateTime *string `json:"createTime,omitempty"`

	// Output only. Last update time of this resource.
	// +kcc:proto:field=google.cloud.vmwareengine.v1.ExternalAddress.update_time
	UpdateTime *string `json:"updateTime,omitempty"`

	// Output only. The external IP address of a workload VM.
	// +kcc:proto:field=google.cloud.vmwareengine.v1.ExternalAddress.external_ip
	ExternalIP *string `json:"externalIP,omitempty"`

	// Output only. The state of the resource.
	// +kcc:proto:field=google.cloud.vmwareengine.v1.ExternalAddress.state
	State *string `json:"state,omitempty"`

	// Output only. System-generated unique identifier for the resource.
	// +kcc:proto:field=google.cloud.vmwareengine.v1.ExternalAddress.uid
	Uid *string `json:"uid,omitempty"`
}
