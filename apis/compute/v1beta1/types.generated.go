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
// krm.group: compute.cnrm.cloud.google.com
// krm.version: v1beta1
// proto.service: google.cloud.compute.v1
// resource: ComputeFirewallPolicyRule:FirewallPolicyRule
// resource: ComputeForwardingRule:ForwardingRule
// resource: ComputeNetwork:Network
// resource: ComputeSubnetwork:Subnetwork
// resource: ComputeTargetTcpProxy:TargetTcpProxy

package v1beta1

// +kcc:proto=google.cloud.compute.v1.FirewallPolicyRuleSecureTag
type FirewallPolicyRuleSecureTag struct {
	// Name of the secure tag, created with TagManager's TagValue API.
	// +kcc:proto:field=google.cloud.compute.v1.FirewallPolicyRuleSecureTag.name
	Name *string `json:"name,omitempty"`

	// [Output Only] State of the secure tag, either `EFFECTIVE` or `INEFFECTIVE`. A secure tag is `INEFFECTIVE` when it is deleted or its network is deleted.
	//  Check the State enum for the list of possible values.
	// +kcc:proto:field=google.cloud.compute.v1.FirewallPolicyRuleSecureTag.state
	State *string `json:"state,omitempty"`
}

// +kcc:proto=google.cloud.compute.v1.NetworkParams
type NetworkParams struct {
	// Tag keys/values directly bound to this resource. Tag keys and values have the same definition as resource manager tags. The field is allowed for INSERT only. The keys/values to set on the resource should be specified in either ID { : } or Namespaced format { : }. For example the following are valid inputs: * {"tagKeys/333" : "tagValues/444", "tagKeys/123" : "tagValues/456"} * {"123/environment" : "production", "345/abc" : "xyz"} Note: * Invalid combinations of ID & namespaced format is not supported. For instance: {"123/environment" : "tagValues/444"} is invalid.
	// +kcc:proto:field=google.cloud.compute.v1.NetworkParams.resource_manager_tags
	ResourceManagerTags map[string]string `json:"resourceManagerTags,omitempty"`
}

// +kcc:proto=google.cloud.compute.v1.NetworkPeering
type NetworkPeering struct {
	// This field will be deprecated soon. Use the exchange_subnet_routes field instead. Indicates whether full mesh connectivity is created and managed automatically between peered networks. Currently this field should always be true since Google Compute Engine will automatically create and manage subnetwork routes between two networks when peering state is ACTIVE.
	// +kcc:proto:field=google.cloud.compute.v1.NetworkPeering.auto_create_routes
	AutoCreateRoutes *bool `json:"autoCreateRoutes,omitempty"`

	// [Output Only] The effective state of the peering connection as a whole.
	// +kcc:proto:field=google.cloud.compute.v1.NetworkPeering.connection_status
	ConnectionStatus *NetworkPeeringConnectionStatus `json:"connectionStatus,omitempty"`

	// Indicates whether full mesh connectivity is created and managed automatically between peered networks. Currently this field should always be true since Google Compute Engine will automatically create and manage subnetwork routes between two networks when peering state is ACTIVE.
	// +kcc:proto:field=google.cloud.compute.v1.NetworkPeering.exchange_subnet_routes
	ExchangeSubnetRoutes *bool `json:"exchangeSubnetRoutes,omitempty"`

	// Whether to export the custom routes to peer network. The default value is false.
	// +kcc:proto:field=google.cloud.compute.v1.NetworkPeering.export_custom_routes
	ExportCustomRoutes *bool `json:"exportCustomRoutes,omitempty"`

	// Whether subnet routes with public IP range are exported. The default value is true, all subnet routes are exported. IPv4 special-use ranges are always exported to peers and are not controlled by this field.
	// +kcc:proto:field=google.cloud.compute.v1.NetworkPeering.export_subnet_routes_with_public_ip
	ExportSubnetRoutesWithPublicIP *bool `json:"exportSubnetRoutesWithPublicIP,omitempty"`

	// Whether to import the custom routes from peer network. The default value is false.
	// +kcc:proto:field=google.cloud.compute.v1.NetworkPeering.import_custom_routes
	ImportCustomRoutes *bool `json:"importCustomRoutes,omitempty"`

	// Whether subnet routes with public IP range are imported. The default value is false. IPv4 special-use ranges are always imported from peers and are not controlled by this field.
	// +kcc:proto:field=google.cloud.compute.v1.NetworkPeering.import_subnet_routes_with_public_ip
	ImportSubnetRoutesWithPublicIP *bool `json:"importSubnetRoutesWithPublicIP,omitempty"`

	// Name of this peering. Provided by the client when the peering is created. The name must comply with RFC1035. Specifically, the name must be 1-63 characters long and match regular expression `[a-z]([-a-z0-9]*[a-z0-9])?`. The first character must be a lowercase letter, and all the following characters must be a dash, lowercase letter, or digit, except the last character, which cannot be a dash.
	// +kcc:proto:field=google.cloud.compute.v1.NetworkPeering.name
	Name *string `json:"name,omitempty"`

	// The URL of the peer network. It can be either full URL or partial URL. The peer network may belong to a different project. If the partial URL does not contain project, it is assumed that the peer network is in the same project as the current network.
	// +kcc:proto:field=google.cloud.compute.v1.NetworkPeering.network
	Network *string `json:"network,omitempty"`

	// [Output Only] Maximum Transmission Unit in bytes of the peer network.
	// +kcc:proto:field=google.cloud.compute.v1.NetworkPeering.peer_mtu
	PeerMtu *int32 `json:"peerMtu,omitempty"`

	// Which IP version(s) of traffic and routes are allowed to be imported or exported between peer networks. The default value is IPV4_ONLY.
	//  Check the StackType enum for the list of possible values.
	// +kcc:proto:field=google.cloud.compute.v1.NetworkPeering.stack_type
	StackType *string `json:"stackType,omitempty"`

	// [Output Only] State for the peering, either `ACTIVE` or `INACTIVE`. The peering is `ACTIVE` when there's a matching configuration in the peer network.
	//  Check the State enum for the list of possible values.
	// +kcc:proto:field=google.cloud.compute.v1.NetworkPeering.state
	State *string `json:"state,omitempty"`

	// [Output Only] Details about the current state of the peering.
	// +kcc:proto:field=google.cloud.compute.v1.NetworkPeering.state_details
	StateDetails *string `json:"stateDetails,omitempty"`

	// The update strategy determines the semantics for updates and deletes to the peering connection configuration.
	//  Check the UpdateStrategy enum for the list of possible values.
	// +kcc:proto:field=google.cloud.compute.v1.NetworkPeering.update_strategy
	UpdateStrategy *string `json:"updateStrategy,omitempty"`
}

// +kcc:proto=google.cloud.compute.v1.NetworkPeeringConnectionStatus
type NetworkPeeringConnectionStatus struct {
	// The consensus state contains information about the status of update and delete for a consensus peering connection.
	// +kcc:proto:field=google.cloud.compute.v1.NetworkPeeringConnectionStatus.consensus_state
	ConsensusState *NetworkPeeringConnectionStatusConsensusState `json:"consensusState,omitempty"`

	// The active connectivity settings for the peering connection based on the settings of the network peerings.
	// +kcc:proto:field=google.cloud.compute.v1.NetworkPeeringConnectionStatus.traffic_configuration
	TrafficConfiguration *NetworkPeeringConnectionStatusTrafficConfiguration `json:"trafficConfiguration,omitempty"`

	// The update strategy determines the update/delete semantics for this peering connection.
	//  Check the UpdateStrategy enum for the list of possible values.
	// +kcc:proto:field=google.cloud.compute.v1.NetworkPeeringConnectionStatus.update_strategy
	UpdateStrategy *string `json:"updateStrategy,omitempty"`
}

// +kcc:proto=google.cloud.compute.v1.NetworkPeeringConnectionStatusConsensusState
type NetworkPeeringConnectionStatusConsensusState struct {
	// The status of the delete request.
	//  Check the DeleteStatus enum for the list of possible values.
	// +kcc:proto:field=google.cloud.compute.v1.NetworkPeeringConnectionStatusConsensusState.delete_status
	DeleteStatus *string `json:"deleteStatus,omitempty"`

	// The status of the update request.
	//  Check the UpdateStatus enum for the list of possible values.
	// +kcc:proto:field=google.cloud.compute.v1.NetworkPeeringConnectionStatusConsensusState.update_status
	UpdateStatus *string `json:"updateStatus,omitempty"`
}

// +kcc:proto=google.cloud.compute.v1.NetworkPeeringConnectionStatusTrafficConfiguration
type NetworkPeeringConnectionStatusTrafficConfiguration struct {
	// Whether custom routes are being exported to the peer network.
	// +kcc:proto:field=google.cloud.compute.v1.NetworkPeeringConnectionStatusTrafficConfiguration.export_custom_routes_to_peer
	ExportCustomRoutesToPeer *bool `json:"exportCustomRoutesToPeer,omitempty"`

	// Whether subnet routes with public IP ranges are being exported to the peer network.
	// +kcc:proto:field=google.cloud.compute.v1.NetworkPeeringConnectionStatusTrafficConfiguration.export_subnet_routes_with_public_ip_to_peer
	ExportSubnetRoutesWithPublicIPToPeer *bool `json:"exportSubnetRoutesWithPublicIPToPeer,omitempty"`

	// Whether custom routes are being imported from the peer network.
	// +kcc:proto:field=google.cloud.compute.v1.NetworkPeeringConnectionStatusTrafficConfiguration.import_custom_routes_from_peer
	ImportCustomRoutesFromPeer *bool `json:"importCustomRoutesFromPeer,omitempty"`

	// Whether subnet routes with public IP ranges are being imported from the peer network.
	// +kcc:proto:field=google.cloud.compute.v1.NetworkPeeringConnectionStatusTrafficConfiguration.import_subnet_routes_with_public_ip_from_peer
	ImportSubnetRoutesWithPublicIPFromPeer *bool `json:"importSubnetRoutesWithPublicIPFromPeer,omitempty"`

	// Which IP version(s) of traffic and routes are being imported or exported between peer networks.
	//  Check the StackType enum for the list of possible values.
	// +kcc:proto:field=google.cloud.compute.v1.NetworkPeeringConnectionStatusTrafficConfiguration.stack_type
	StackType *string `json:"stackType,omitempty"`
}

// +kcc:proto=google.cloud.compute.v1.NetworkRoutingConfig
type NetworkRoutingConfig struct {
	// Enable comparison of Multi-Exit Discriminators (MED) across routes with different neighbor ASNs when using the STANDARD BGP best path selection algorithm.
	// +kcc:proto:field=google.cloud.compute.v1.NetworkRoutingConfig.bgp_always_compare_med
	BGPAlwaysCompareMed *bool `json:"bgpAlwaysCompareMed,omitempty"`

	// The BGP best path selection algorithm to be employed within this network for dynamic routes learned by Cloud Routers. Can be LEGACY (default) or STANDARD.
	//  Check the BgpBestPathSelectionMode enum for the list of possible values.
	// +kcc:proto:field=google.cloud.compute.v1.NetworkRoutingConfig.bgp_best_path_selection_mode
	BGPBestPathSelectionMode *string `json:"bgpBestPathSelectionMode,omitempty"`

	// Allows to define a preferred approach for handling inter-region cost in the selection process when using the STANDARD BGP best path selection algorithm. Can be DEFAULT or ADD_COST_TO_MED.
	//  Check the BgpInterRegionCost enum for the list of possible values.
	// +kcc:proto:field=google.cloud.compute.v1.NetworkRoutingConfig.bgp_inter_region_cost
	BGPInterRegionCost *string `json:"bgpInterRegionCost,omitempty"`

	// [Output Only] Effective value of the bgp_always_compare_med field.
	// +kcc:proto:field=google.cloud.compute.v1.NetworkRoutingConfig.effective_bgp_always_compare_med
	EffectiveBGPAlwaysCompareMed *bool `json:"effectiveBGPAlwaysCompareMed,omitempty"`

	// [Output Only] Effective value of the bgp_inter_region_cost field.
	//  Check the EffectiveBgpInterRegionCost enum for the list of possible values.
	// +kcc:proto:field=google.cloud.compute.v1.NetworkRoutingConfig.effective_bgp_inter_region_cost
	EffectiveBGPInterRegionCost *string `json:"effectiveBGPInterRegionCost,omitempty"`

	// The network-wide routing mode to use. If set to REGIONAL, this network's Cloud Routers will only advertise routes with subnets of this network in the same region as the router. If set to GLOBAL, this network's Cloud Routers will advertise routes with all subnets of this network, across regions.
	//  Check the RoutingMode enum for the list of possible values.
	// +kcc:proto:field=google.cloud.compute.v1.NetworkRoutingConfig.routing_mode
	RoutingMode *string `json:"routingMode,omitempty"`
}

// +kcc:proto=google.cloud.compute.v1.SubnetworkParams
type SubnetworkParams struct {
	// Tag keys/values directly bound to this resource. Tag keys and values have the same definition as resource manager tags. The field is allowed for INSERT only. The keys/values to set on the resource should be specified in either ID { : } or Namespaced format { : }. For example the following are valid inputs: * {"tagKeys/333" : "tagValues/444", "tagKeys/123" : "tagValues/456"} * {"123/environment" : "production", "345/abc" : "xyz"} Note: * Invalid combinations of ID & namespaced format is not supported. For instance: {"123/environment" : "tagValues/444"} is invalid.
	// +kcc:proto:field=google.cloud.compute.v1.SubnetworkParams.resource_manager_tags
	ResourceManagerTags map[string]string `json:"resourceManagerTags,omitempty"`
}
