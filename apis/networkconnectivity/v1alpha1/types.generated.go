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


// +kcc:proto=google.cloud.networkconnectivity.v1.NextHopInterconnectAttachment
type NextHopInterconnectAttachment struct {
	// The URI of the interconnect attachment resource.
	// +kcc:proto:field=google.cloud.networkconnectivity.v1.NextHopInterconnectAttachment.uri
	URI *string `json:"uri,omitempty"`

	// The VPC network where this interconnect attachment is located.
	// +kcc:proto:field=google.cloud.networkconnectivity.v1.NextHopInterconnectAttachment.vpc_network
	VpcNetwork *string `json:"vpcNetwork,omitempty"`

	// Indicates whether site-to-site data transfer is allowed for this
	//  interconnect attachment resource. Data transfer is available only in
	//  [supported
	//  locations](https://cloud.google.com/network-connectivity/docs/network-connectivity-center/concepts/locations).
	// +kcc:proto:field=google.cloud.networkconnectivity.v1.NextHopInterconnectAttachment.site_to_site_data_transfer
	SiteToSiteDataTransfer *bool `json:"siteToSiteDataTransfer,omitempty"`
}

// +kcc:proto=google.cloud.networkconnectivity.v1.NextHopRouterApplianceInstance
type NextHopRouterApplianceInstance struct {
	// The URI of the Router appliance instance.
	// +kcc:proto:field=google.cloud.networkconnectivity.v1.NextHopRouterApplianceInstance.uri
	URI *string `json:"uri,omitempty"`

	// The VPC network where this VM is located.
	// +kcc:proto:field=google.cloud.networkconnectivity.v1.NextHopRouterApplianceInstance.vpc_network
	VpcNetwork *string `json:"vpcNetwork,omitempty"`

	// Indicates whether site-to-site data transfer is allowed for this Router
	//  appliance instance resource. Data transfer is available only in [supported
	//  locations](https://cloud.google.com/network-connectivity/docs/network-connectivity-center/concepts/locations).
	// +kcc:proto:field=google.cloud.networkconnectivity.v1.NextHopRouterApplianceInstance.site_to_site_data_transfer
	SiteToSiteDataTransfer *bool `json:"siteToSiteDataTransfer,omitempty"`
}

// +kcc:proto=google.cloud.networkconnectivity.v1.NextHopVPNTunnel
type NextHopVPNTunnel struct {
	// The URI of the VPN tunnel resource.
	// +kcc:proto:field=google.cloud.networkconnectivity.v1.NextHopVPNTunnel.uri
	URI *string `json:"uri,omitempty"`

	// The VPC network where this VPN tunnel is located.
	// +kcc:proto:field=google.cloud.networkconnectivity.v1.NextHopVPNTunnel.vpc_network
	VpcNetwork *string `json:"vpcNetwork,omitempty"`

	// Indicates whether site-to-site data transfer is allowed for this VPN tunnel
	//  resource. Data transfer is available only in [supported
	//  locations](https://cloud.google.com/network-connectivity/docs/network-connectivity-center/concepts/locations).
	// +kcc:proto:field=google.cloud.networkconnectivity.v1.NextHopVPNTunnel.site_to_site_data_transfer
	SiteToSiteDataTransfer *bool `json:"siteToSiteDataTransfer,omitempty"`
}

// +kcc:proto=google.cloud.networkconnectivity.v1.NextHopVpcNetwork
type NextHopVpcNetwork struct {
	// The URI of the VPC network resource
	// +kcc:proto:field=google.cloud.networkconnectivity.v1.NextHopVpcNetwork.uri
	URI *string `json:"uri,omitempty"`
}

// +kcc:proto=google.cloud.networkconnectivity.v1.Route
type Route struct {
	// Immutable. The name of the route. Route names must be unique. Route names
	//  use the following form:
	//       `projects/{project_number}/locations/global/hubs/{hub}/routeTables/{route_table_id}/routes/{route_id}`
	// +kcc:proto:field=google.cloud.networkconnectivity.v1.Route.name
	Name *string `json:"name,omitempty"`

	// The destination IP address range.
	// +kcc:proto:field=google.cloud.networkconnectivity.v1.Route.ip_cidr_range
	IPCidrRange *string `json:"ipCidrRange,omitempty"`

	// Immutable. The destination VPC network for packets on this route.
	// +kcc:proto:field=google.cloud.networkconnectivity.v1.Route.next_hop_vpc_network
	NextHopVpcNetwork *NextHopVpcNetwork `json:"nextHopVpcNetwork,omitempty"`

	// Optional labels in key-value pair format. For more information about
	//  labels, see [Requirements for
	//  labels](https://cloud.google.com/resource-manager/docs/creating-managing-labels#requirements).
	// +kcc:proto:field=google.cloud.networkconnectivity.v1.Route.labels
	Labels map[string]string `json:"labels,omitempty"`

	// An optional description of the route.
	// +kcc:proto:field=google.cloud.networkconnectivity.v1.Route.description
	Description *string `json:"description,omitempty"`

	// Immutable. The spoke that this route leads to.
	//  Example: projects/12345/locations/global/spokes/SPOKE
	// +kcc:proto:field=google.cloud.networkconnectivity.v1.Route.spoke
	Spoke *string `json:"spoke,omitempty"`

	// Immutable. The next-hop VPN tunnel for packets on this route.
	// +kcc:proto:field=google.cloud.networkconnectivity.v1.Route.next_hop_vpn_tunnel
	NextHopVpnTunnel *NextHopVPNTunnel `json:"nextHopVpnTunnel,omitempty"`

	// Immutable. The next-hop Router appliance instance for packets on this
	//  route.
	// +kcc:proto:field=google.cloud.networkconnectivity.v1.Route.next_hop_router_appliance_instance
	NextHopRouterApplianceInstance *NextHopRouterApplianceInstance `json:"nextHopRouterApplianceInstance,omitempty"`

	// Immutable. The next-hop VLAN attachment for packets on this route.
	// +kcc:proto:field=google.cloud.networkconnectivity.v1.Route.next_hop_interconnect_attachment
	NextHopInterconnectAttachment *NextHopInterconnectAttachment `json:"nextHopInterconnectAttachment,omitempty"`
}

// +kcc:proto=google.cloud.networkconnectivity.v1.Route
type RouteObservedState struct {
	// Output only. The time the route was created.
	// +kcc:proto:field=google.cloud.networkconnectivity.v1.Route.create_time
	CreateTime *string `json:"createTime,omitempty"`

	// Output only. The time the route was last updated.
	// +kcc:proto:field=google.cloud.networkconnectivity.v1.Route.update_time
	UpdateTime *string `json:"updateTime,omitempty"`

	// Output only. The route's type. Its type is determined by the properties of
	//  its IP address range.
	// +kcc:proto:field=google.cloud.networkconnectivity.v1.Route.type
	Type *string `json:"type,omitempty"`

	// Output only. The Google-generated UUID for the route. This value is unique
	//  across all Network Connectivity Center route resources. If a
	//  route is deleted and another with the same name is created,
	//  the new route is assigned a different `uid`.
	// +kcc:proto:field=google.cloud.networkconnectivity.v1.Route.uid
	Uid *string `json:"uid,omitempty"`

	// Output only. The current lifecycle state of the route.
	// +kcc:proto:field=google.cloud.networkconnectivity.v1.Route.state
	State *string `json:"state,omitempty"`

	// Output only. The origin location of the route.
	//  Uses the following form: "projects/{project}/locations/{location}"
	//  Example: projects/1234/locations/us-central1
	// +kcc:proto:field=google.cloud.networkconnectivity.v1.Route.location
	Location *string `json:"location,omitempty"`

	// Output only. The priority of this route. Priority is used to break ties in
	//  cases where a destination matches more than one route. In these cases the
	//  route with the lowest-numbered priority value wins.
	// +kcc:proto:field=google.cloud.networkconnectivity.v1.Route.priority
	Priority *int64 `json:"priority,omitempty"`
}
