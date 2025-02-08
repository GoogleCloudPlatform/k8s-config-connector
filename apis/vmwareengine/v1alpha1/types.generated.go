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


// +kcc:proto=google.cloud.vmwareengine.v1.ManagementDnsZoneBinding
type ManagementDnsZoneBinding struct {

	// User-provided description for this resource.
	// +kcc:proto:field=google.cloud.vmwareengine.v1.ManagementDnsZoneBinding.description
	Description *string `json:"description,omitempty"`

	// Network to bind is a standard consumer VPC.
	//  Specify the name in the following form for consumer
	//  VPC network: `projects/{project}/global/networks/{network_id}`.
	//  `{project}` can either be a project number or a project ID.
	// +kcc:proto:field=google.cloud.vmwareengine.v1.ManagementDnsZoneBinding.vpc_network
	VpcNetwork *string `json:"vpcNetwork,omitempty"`

	// Network to bind is a VMware Engine network.
	//  Specify the name in the following form for VMware engine network:
	//  `projects/{project}/locations/global/vmwareEngineNetworks/{vmware_engine_network_id}`.
	//  `{project}` can either be a project number or a project ID.
	// +kcc:proto:field=google.cloud.vmwareengine.v1.ManagementDnsZoneBinding.vmware_engine_network
	VmwareEngineNetwork *string `json:"vmwareEngineNetwork,omitempty"`
}

// +kcc:proto=google.cloud.vmwareengine.v1.ManagementDnsZoneBinding
type ManagementDnsZoneBindingObservedState struct {
	// Output only. The resource name of this binding.
	//  Resource names are schemeless URIs that follow the conventions in
	//  https://cloud.google.com/apis/design/resource_names.
	//  For example:
	//  `projects/my-project/locations/us-central1-a/privateClouds/my-cloud/managementDnsZoneBindings/my-management-dns-zone-binding`
	// +kcc:proto:field=google.cloud.vmwareengine.v1.ManagementDnsZoneBinding.name
	Name *string `json:"name,omitempty"`

	// Output only. Creation time of this resource.
	// +kcc:proto:field=google.cloud.vmwareengine.v1.ManagementDnsZoneBinding.create_time
	CreateTime *string `json:"createTime,omitempty"`

	// Output only. Last update time of this resource.
	// +kcc:proto:field=google.cloud.vmwareengine.v1.ManagementDnsZoneBinding.update_time
	UpdateTime *string `json:"updateTime,omitempty"`

	// Output only. The state of the resource.
	// +kcc:proto:field=google.cloud.vmwareengine.v1.ManagementDnsZoneBinding.state
	State *string `json:"state,omitempty"`

	// Output only. System-generated unique identifier for the resource.
	// +kcc:proto:field=google.cloud.vmwareengine.v1.ManagementDnsZoneBinding.uid
	Uid *string `json:"uid,omitempty"`
}
