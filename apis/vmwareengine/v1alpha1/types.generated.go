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

// +kcc:proto=google.cloud.vmwareengine.v1.NetworkPolicy
type NetworkPolicy struct {

	// Network service that allows VMware workloads to access the internet.
	// +kcc:proto:field=google.cloud.vmwareengine.v1.NetworkPolicy.internet_access
	InternetAccess *NetworkPolicy_NetworkService `json:"internetAccess,omitempty"`

	// Network service that allows External IP addresses to be assigned to VMware
	//  workloads. This service can only be enabled when `internet_access` is also
	//  enabled.
	// +kcc:proto:field=google.cloud.vmwareengine.v1.NetworkPolicy.external_ip
	ExternalIP *NetworkPolicy_NetworkService `json:"externalIP,omitempty"`

	// Required. IP address range in CIDR notation used to create internet access
	//  and external IP access. An RFC 1918 CIDR block, with a "/26" prefix, is
	//  required. The range cannot overlap with any prefixes either in the consumer
	//  VPC network or in use by the private clouds attached to that VPC network.
	// +kcc:proto:field=google.cloud.vmwareengine.v1.NetworkPolicy.edge_services_cidr
	EdgeServicesCIDR *string `json:"edgeServicesCIDR,omitempty"`

	// Optional. The relative resource name of the VMware Engine network.
	//  Specify the name in the following form:
	//  `projects/{project}/locations/{location}/vmwareEngineNetworks/{vmware_engine_network_id}`
	//  where `{project}` can either be a project number or a project ID.
	// +kcc:proto:field=google.cloud.vmwareengine.v1.NetworkPolicy.vmware_engine_network
	VmwareEngineNetwork *string `json:"vmwareEngineNetwork,omitempty"`

	// Optional. User-provided description for this network policy.
	// +kcc:proto:field=google.cloud.vmwareengine.v1.NetworkPolicy.description
	Description *string `json:"description,omitempty"`
}

// +kcc:proto=google.cloud.vmwareengine.v1.NetworkPolicy.NetworkService
type NetworkPolicy_NetworkService struct {
	// True if the service is enabled; false otherwise.
	// +kcc:proto:field=google.cloud.vmwareengine.v1.NetworkPolicy.NetworkService.enabled
	Enabled *bool `json:"enabled,omitempty"`
}

// +kcc:proto=google.cloud.vmwareengine.v1.VmwareEngineNetwork.VpcNetwork
type VmwareEngineNetwork_VpcNetwork struct {
}

// +kcc:proto=google.cloud.vmwareengine.v1.NetworkPolicy
type NetworkPolicyObservedState struct {
	// Output only. The resource name of this network policy.
	//  Resource names are schemeless URIs that follow the conventions in
	//  https://cloud.google.com/apis/design/resource_names.
	//  For example:
	//  `projects/my-project/locations/us-central1/networkPolicies/my-network-policy`
	// +kcc:proto:field=google.cloud.vmwareengine.v1.NetworkPolicy.name
	Name *string `json:"name,omitempty"`

	// Output only. Creation time of this resource.
	// +kcc:proto:field=google.cloud.vmwareengine.v1.NetworkPolicy.create_time
	CreateTime *string `json:"createTime,omitempty"`

	// Output only. Last update time of this resource.
	// +kcc:proto:field=google.cloud.vmwareengine.v1.NetworkPolicy.update_time
	UpdateTime *string `json:"updateTime,omitempty"`

	// Network service that allows VMware workloads to access the internet.
	// +kcc:proto:field=google.cloud.vmwareengine.v1.NetworkPolicy.internet_access
	InternetAccess *NetworkPolicy_NetworkServiceObservedState `json:"internetAccess,omitempty"`

	// Output only. System-generated unique identifier for the resource.
	// +kcc:proto:field=google.cloud.vmwareengine.v1.NetworkPolicy.uid
	Uid *string `json:"uid,omitempty"`

	// Output only. The canonical name of the VMware Engine network in the form:
	//  `projects/{project_number}/locations/{location}/vmwareEngineNetworks/{vmware_engine_network_id}`
	// +kcc:proto:field=google.cloud.vmwareengine.v1.NetworkPolicy.vmware_engine_network_canonical
	VmwareEngineNetworkCanonical *string `json:"vmwareEngineNetworkCanonical,omitempty"`
}

// +kcc:proto=google.cloud.vmwareengine.v1.NetworkPolicy.NetworkService
type NetworkPolicy_NetworkServiceObservedState struct {
	// Output only. State of the service. New values may be added to this enum
	//  when appropriate.
	// +kcc:proto:field=google.cloud.vmwareengine.v1.NetworkPolicy.NetworkService.state
	State *string `json:"state,omitempty"`
}

// +kcc:proto=google.cloud.vmwareengine.v1.VmwareEngineNetwork.VpcNetwork
type VmwareEngineNetwork_VpcNetworkObservedState struct {
	// Output only. Type of VPC network (INTRANET, INTERNET, or
	//  GOOGLE_CLOUD)
	// +kcc:proto:field=google.cloud.vmwareengine.v1.VmwareEngineNetwork.VpcNetwork.type
	Type *string `json:"type,omitempty"`

	// Output only. The relative resource name of the service VPC network this
	//  VMware Engine network is attached to. For example:
	//  `projects/123123/global/networks/my-network`
	// +kcc:proto:field=google.cloud.vmwareengine.v1.VmwareEngineNetwork.VpcNetwork.network
	Network *string `json:"network,omitempty"`
}
