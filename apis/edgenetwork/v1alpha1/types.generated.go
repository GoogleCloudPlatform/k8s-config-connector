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


// +kcc:proto=google.cloud.edgenetwork.v1.Router
type Router struct {
	// Required. The canonical resource name of the router.
	// +kcc:proto:field=google.cloud.edgenetwork.v1.Router.name
	Name *string `json:"name,omitempty"`

	// Labels associated with this resource.
	// +kcc:proto:field=google.cloud.edgenetwork.v1.Router.labels
	Labels map[string]string `json:"labels,omitempty"`

	// Optional. A free-text description of the resource. Max length 1024
	//  characters.
	// +kcc:proto:field=google.cloud.edgenetwork.v1.Router.description
	Description *string `json:"description,omitempty"`

	// Required. The canonical name of the network to which this router belongs.
	//  The name is in the form of
	//  `projects/{project}/locations/{location}/zones/{zone}/networks/{network}`.
	// +kcc:proto:field=google.cloud.edgenetwork.v1.Router.network
	Network *string `json:"network,omitempty"`

	// Router interfaces.
	// +kcc:proto:field=google.cloud.edgenetwork.v1.Router.interface
	Interface []Router_Interface `json:"interface,omitempty"`

	// BGP peers.
	// +kcc:proto:field=google.cloud.edgenetwork.v1.Router.bgp_peer
	BgpPeer []Router_BgpPeer `json:"bgpPeer,omitempty"`

	// BGP information specific to this router.
	// +kcc:proto:field=google.cloud.edgenetwork.v1.Router.bgp
	Bgp *Router_Bgp `json:"bgp,omitempty"`

	// Optional. A list of CIDRs in IP/Length format to advertise northbound as
	//  static routes from this router.
	// +kcc:proto:field=google.cloud.edgenetwork.v1.Router.route_advertisements
	RouteAdvertisements []string `json:"routeAdvertisements,omitempty"`
}

// +kcc:proto=google.cloud.edgenetwork.v1.Router.Bgp
type Router_Bgp struct {
	// Locally assigned BGP ASN.
	// +kcc:proto:field=google.cloud.edgenetwork.v1.Router.Bgp.asn
	Asn *uint32 `json:"asn,omitempty"`

	// The interval in seconds between BGP keepalive messages that are
	//  sent to the peer. Default is 20 with value between 20 and 60.
	// +kcc:proto:field=google.cloud.edgenetwork.v1.Router.Bgp.keepalive_interval_in_seconds
	KeepaliveIntervalInSeconds *uint32 `json:"keepaliveIntervalInSeconds,omitempty"`
}

// +kcc:proto=google.cloud.edgenetwork.v1.Router.BgpPeer
type Router_BgpPeer struct {
	// Name of this BGP peer. Unique within the Zones resource.
	// +kcc:proto:field=google.cloud.edgenetwork.v1.Router.BgpPeer.name
	Name *string `json:"name,omitempty"`

	// Name of the RouterInterface the BGP peer is associated with.
	// +kcc:proto:field=google.cloud.edgenetwork.v1.Router.BgpPeer.interface
	Interface *string `json:"interface,omitempty"`

	// IP range of the interface within Google.
	// +kcc:proto:field=google.cloud.edgenetwork.v1.Router.BgpPeer.interface_ipv4_cidr
	InterfaceIpv4Cidr *string `json:"interfaceIpv4Cidr,omitempty"`

	// IPv6 range of the interface within Google.
	// +kcc:proto:field=google.cloud.edgenetwork.v1.Router.BgpPeer.interface_ipv6_cidr
	InterfaceIpv6Cidr *string `json:"interfaceIpv6Cidr,omitempty"`

	// IP range of the BGP interface outside Google.
	// +kcc:proto:field=google.cloud.edgenetwork.v1.Router.BgpPeer.peer_ipv4_cidr
	PeerIpv4Cidr *string `json:"peerIpv4Cidr,omitempty"`

	// IPv6 range of the BGP interface outside Google.
	// +kcc:proto:field=google.cloud.edgenetwork.v1.Router.BgpPeer.peer_ipv6_cidr
	PeerIpv6Cidr *string `json:"peerIpv6Cidr,omitempty"`

	// Peer BGP Autonomous System Number (ASN). Each BGP interface may use
	//  a different value.
	// +kcc:proto:field=google.cloud.edgenetwork.v1.Router.BgpPeer.peer_asn
	PeerAsn *uint32 `json:"peerAsn,omitempty"`
}

// +kcc:proto=google.cloud.edgenetwork.v1.Router.Interface
type Router_Interface struct {
	// Name of this interface entry. Unique within the Zones resource.
	// +kcc:proto:field=google.cloud.edgenetwork.v1.Router.Interface.name
	Name *string `json:"name,omitempty"`

	// IP address and range of the interface.
	// +kcc:proto:field=google.cloud.edgenetwork.v1.Router.Interface.ipv4_cidr
	Ipv4Cidr *string `json:"ipv4Cidr,omitempty"`

	// IPv6 address and range of the interface.
	// +kcc:proto:field=google.cloud.edgenetwork.v1.Router.Interface.ipv6_cidr
	Ipv6Cidr *string `json:"ipv6Cidr,omitempty"`

	// The canonical name of the linked Interconnect attachment.
	// +kcc:proto:field=google.cloud.edgenetwork.v1.Router.Interface.linked_interconnect_attachment
	LinkedInterconnectAttachment *string `json:"linkedInterconnectAttachment,omitempty"`

	// The canonical name of the subnetwork resource that this interface
	//  belongs to.
	// +kcc:proto:field=google.cloud.edgenetwork.v1.Router.Interface.subnetwork
	Subnetwork *string `json:"subnetwork,omitempty"`

	// Create loopback interface in the router when specified.
	//  The number of IP addresses must match the number of TOR devices.
	// +kcc:proto:field=google.cloud.edgenetwork.v1.Router.Interface.loopback_ip_addresses
	LoopbackIPAddresses []string `json:"loopbackIPAddresses,omitempty"`
}

// +kcc:proto=google.cloud.edgenetwork.v1.Router
type RouterObservedState struct {
	// Output only. The time when the router was created.
	// +kcc:proto:field=google.cloud.edgenetwork.v1.Router.create_time
	CreateTime *string `json:"createTime,omitempty"`

	// Output only. The time when the router was last updated.
	// +kcc:proto:field=google.cloud.edgenetwork.v1.Router.update_time
	UpdateTime *string `json:"updateTime,omitempty"`

	// BGP peers.
	// +kcc:proto:field=google.cloud.edgenetwork.v1.Router.bgp_peer
	BgpPeer []Router_BgpPeerObservedState `json:"bgpPeer,omitempty"`

	// Output only. Current stage of the resource to the device by config push.
	// +kcc:proto:field=google.cloud.edgenetwork.v1.Router.state
	State *string `json:"state,omitempty"`
}

// +kcc:proto=google.cloud.edgenetwork.v1.Router.BgpPeer
type Router_BgpPeerObservedState struct {
	// Output only. Local BGP Autonomous System Number (ASN).
	//  This field is ST_NOT_REQUIRED because it stores private ASNs, which are
	//  meaningless outside the zone in which they are being used.
	// +kcc:proto:field=google.cloud.edgenetwork.v1.Router.BgpPeer.local_asn
	LocalAsn *uint32 `json:"localAsn,omitempty"`
}
