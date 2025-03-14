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

// +kcc:proto=google.cloud.vmwareengine.v1.NetworkPeering
type NetworkPeering struct {

	// Required. The relative resource name of the network to peer with
	//  a standard VMware Engine network. The provided network can be a
	//  consumer VPC network or another standard VMware Engine network. If the
	//  `peer_network_type` is VMWARE_ENGINE_NETWORK, specify the name in the form:
	//  `projects/{project}/locations/global/vmwareEngineNetworks/{vmware_engine_network_id}`.
	//  Otherwise specify the name in the form:
	//  `projects/{project}/global/networks/{network_id}`, where
	//  `{project}` can either be a project number or a project ID.
	// +kcc:proto:field=google.cloud.vmwareengine.v1.NetworkPeering.peer_network
	PeerNetwork *string `json:"peerNetwork,omitempty"`

	// Optional. True if custom routes are exported to the peered network;
	//  false otherwise. The default value is true.
	// +kcc:proto:field=google.cloud.vmwareengine.v1.NetworkPeering.export_custom_routes
	ExportCustomRoutes *bool `json:"exportCustomRoutes,omitempty"`

	// Optional. True if custom routes are imported from the peered network;
	//  false otherwise. The default value is true.
	// +kcc:proto:field=google.cloud.vmwareengine.v1.NetworkPeering.import_custom_routes
	ImportCustomRoutes *bool `json:"importCustomRoutes,omitempty"`

	// Optional. True if full mesh connectivity is created and managed
	//  automatically between peered networks; false otherwise. Currently this
	//  field is always true because Google Compute Engine automatically creates
	//  and manages subnetwork routes between two VPC networks when peering state
	//  is 'ACTIVE'.
	// +kcc:proto:field=google.cloud.vmwareengine.v1.NetworkPeering.exchange_subnet_routes
	ExchangeSubnetRoutes *bool `json:"exchangeSubnetRoutes,omitempty"`

	// Optional. True if all subnet routes with a public IP address range are
	//  exported; false otherwise. The default value is true. IPv4 special-use
	//  ranges (https://en.wikipedia.org/wiki/IPv4#Special_addresses) are always
	//  exported to peers and are not controlled by this field.
	// +kcc:proto:field=google.cloud.vmwareengine.v1.NetworkPeering.export_custom_routes_with_public_ip
	ExportCustomRoutesWithPublicIP *bool `json:"exportCustomRoutesWithPublicIP,omitempty"`

	// Optional. True if all subnet routes with public IP address range are
	//  imported; false otherwise. The default value is true. IPv4 special-use
	//  ranges (https://en.wikipedia.org/wiki/IPv4#Special_addresses) are always
	//  imported to peers and are not controlled by this field.
	// +kcc:proto:field=google.cloud.vmwareengine.v1.NetworkPeering.import_custom_routes_with_public_ip
	ImportCustomRoutesWithPublicIP *bool `json:"importCustomRoutesWithPublicIP,omitempty"`

	// Optional. Maximum transmission unit (MTU) in bytes.
	//  The default value is `1500`. If a value of `0` is provided for this field,
	//  VMware Engine uses the default value instead.
	// +kcc:proto:field=google.cloud.vmwareengine.v1.NetworkPeering.peer_mtu
	PeerMtu *int32 `json:"peerMtu,omitempty"`

	// Required. The type of the network to peer with the VMware Engine network.
	// +kcc:proto:field=google.cloud.vmwareengine.v1.NetworkPeering.peer_network_type
	PeerNetworkType *string `json:"peerNetworkType,omitempty"`

	// Required. The relative resource name of the VMware Engine network.
	//  Specify the name in the following form:
	//  `projects/{project}/locations/{location}/vmwareEngineNetworks/{vmware_engine_network_id}`
	//  where `{project}` can either be a project number or a project ID.
	// +kcc:proto:field=google.cloud.vmwareengine.v1.NetworkPeering.vmware_engine_network
	VmwareEngineNetwork *string `json:"vmwareEngineNetwork,omitempty"`

	// Optional. User-provided description for this network peering.
	// +kcc:proto:field=google.cloud.vmwareengine.v1.NetworkPeering.description
	Description *string `json:"description,omitempty"`
}

// +kcc:proto=google.cloud.vmwareengine.v1.VmwareEngineNetwork.VpcNetwork
type VmwareEngineNetwork_VpcNetwork struct {
}

// +kcc:proto=google.cloud.vmwareengine.v1.NetworkPeering
type NetworkPeeringObservedState struct {
	// Output only. The resource name of the network peering. NetworkPeering is a
	//  global resource and location can only be global. Resource names are
	//  scheme-less URIs that follow the conventions in
	//  https://cloud.google.com/apis/design/resource_names.
	//  For example:
	//  `projects/my-project/locations/global/networkPeerings/my-peering`
	// +kcc:proto:field=google.cloud.vmwareengine.v1.NetworkPeering.name
	Name *string `json:"name,omitempty"`

	// Output only. Creation time of this resource.
	// +kcc:proto:field=google.cloud.vmwareengine.v1.NetworkPeering.create_time
	CreateTime *string `json:"createTime,omitempty"`

	// Output only. Last update time of this resource.
	// +kcc:proto:field=google.cloud.vmwareengine.v1.NetworkPeering.update_time
	UpdateTime *string `json:"updateTime,omitempty"`

	// Output only. State of the network peering. This field
	//  has a value of 'ACTIVE' when there's a matching configuration in the peer
	//  network. New values may be added to this enum when appropriate.
	// +kcc:proto:field=google.cloud.vmwareengine.v1.NetworkPeering.state
	State *string `json:"state,omitempty"`

	// Output only. Output Only. Details about the current state of the network
	//  peering.
	// +kcc:proto:field=google.cloud.vmwareengine.v1.NetworkPeering.state_details
	StateDetails *string `json:"stateDetails,omitempty"`

	// Output only. System-generated unique identifier for the resource.
	// +kcc:proto:field=google.cloud.vmwareengine.v1.NetworkPeering.uid
	Uid *string `json:"uid,omitempty"`
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
