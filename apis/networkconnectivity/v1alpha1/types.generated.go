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


// +kcc:proto=google.cloud.networkconnectivity.v1.LinkedInterconnectAttachments
type LinkedInterconnectAttachments struct {
	// The URIs of linked interconnect attachment resources
	// +kcc:proto:field=google.cloud.networkconnectivity.v1.LinkedInterconnectAttachments.uris
	Uris []string `json:"uris,omitempty"`

	// A value that controls whether site-to-site data transfer is enabled for
	//  these resources. Data transfer is available only in [supported
	//  locations](https://cloud.google.com/network-connectivity/docs/network-connectivity-center/concepts/locations).
	// +kcc:proto:field=google.cloud.networkconnectivity.v1.LinkedInterconnectAttachments.site_to_site_data_transfer
	SiteToSiteDataTransfer *bool `json:"siteToSiteDataTransfer,omitempty"`

	// Optional. IP ranges allowed to be included during import from hub (does not
	//  control transit connectivity). The only allowed value for now is
	//  "ALL_IPV4_RANGES".
	// +kcc:proto:field=google.cloud.networkconnectivity.v1.LinkedInterconnectAttachments.include_import_ranges
	IncludeImportRanges []string `json:"includeImportRanges,omitempty"`
}

// +kcc:proto=google.cloud.networkconnectivity.v1.LinkedProducerVpcNetwork
type LinkedProducerVpcNetwork struct {
	// Immutable. The URI of the Service Consumer VPC that the Producer VPC is
	//  peered with.
	// +kcc:proto:field=google.cloud.networkconnectivity.v1.LinkedProducerVpcNetwork.network
	Network *string `json:"network,omitempty"`

	// Immutable. The name of the VPC peering between the Service Consumer VPC and
	//  the Producer VPC (defined in the Tenant project) which is added to the NCC
	//  hub. This peering must be in ACTIVE state.
	// +kcc:proto:field=google.cloud.networkconnectivity.v1.LinkedProducerVpcNetwork.peering
	Peering *string `json:"peering,omitempty"`

	// Optional. IP ranges encompassing the subnets to be excluded from peering.
	// +kcc:proto:field=google.cloud.networkconnectivity.v1.LinkedProducerVpcNetwork.exclude_export_ranges
	ExcludeExportRanges []string `json:"excludeExportRanges,omitempty"`

	// Optional. IP ranges allowed to be included from peering.
	// +kcc:proto:field=google.cloud.networkconnectivity.v1.LinkedProducerVpcNetwork.include_export_ranges
	IncludeExportRanges []string `json:"includeExportRanges,omitempty"`
}

// +kcc:proto=google.cloud.networkconnectivity.v1.LinkedRouterApplianceInstances
type LinkedRouterApplianceInstances struct {
	// The list of router appliance instances.
	// +kcc:proto:field=google.cloud.networkconnectivity.v1.LinkedRouterApplianceInstances.instances
	Instances []RouterApplianceInstance `json:"instances,omitempty"`

	// A value that controls whether site-to-site data transfer is enabled for
	//  these resources. Data transfer is available only in [supported
	//  locations](https://cloud.google.com/network-connectivity/docs/network-connectivity-center/concepts/locations).
	// +kcc:proto:field=google.cloud.networkconnectivity.v1.LinkedRouterApplianceInstances.site_to_site_data_transfer
	SiteToSiteDataTransfer *bool `json:"siteToSiteDataTransfer,omitempty"`

	// Optional. IP ranges allowed to be included during import from hub (does not
	//  control transit connectivity). The only allowed value for now is
	//  "ALL_IPV4_RANGES".
	// +kcc:proto:field=google.cloud.networkconnectivity.v1.LinkedRouterApplianceInstances.include_import_ranges
	IncludeImportRanges []string `json:"includeImportRanges,omitempty"`
}

// +kcc:proto=google.cloud.networkconnectivity.v1.LinkedVpcNetwork
type LinkedVpcNetwork struct {
	// Required. The URI of the VPC network resource.
	// +kcc:proto:field=google.cloud.networkconnectivity.v1.LinkedVpcNetwork.uri
	URI *string `json:"uri,omitempty"`

	// Optional. IP ranges encompassing the subnets to be excluded from peering.
	// +kcc:proto:field=google.cloud.networkconnectivity.v1.LinkedVpcNetwork.exclude_export_ranges
	ExcludeExportRanges []string `json:"excludeExportRanges,omitempty"`

	// Optional. IP ranges allowed to be included from peering.
	// +kcc:proto:field=google.cloud.networkconnectivity.v1.LinkedVpcNetwork.include_export_ranges
	IncludeExportRanges []string `json:"includeExportRanges,omitempty"`
}

// +kcc:proto=google.cloud.networkconnectivity.v1.LinkedVpnTunnels
type LinkedVpnTunnels struct {
	// The URIs of linked VPN tunnel resources.
	// +kcc:proto:field=google.cloud.networkconnectivity.v1.LinkedVpnTunnels.uris
	Uris []string `json:"uris,omitempty"`

	// A value that controls whether site-to-site data transfer is enabled for
	//  these resources. Data transfer is available only in [supported
	//  locations](https://cloud.google.com/network-connectivity/docs/network-connectivity-center/concepts/locations).
	// +kcc:proto:field=google.cloud.networkconnectivity.v1.LinkedVpnTunnels.site_to_site_data_transfer
	SiteToSiteDataTransfer *bool `json:"siteToSiteDataTransfer,omitempty"`

	// Optional. IP ranges allowed to be included during import from hub (does not
	//  control transit connectivity). The only allowed value for now is
	//  "ALL_IPV4_RANGES".
	// +kcc:proto:field=google.cloud.networkconnectivity.v1.LinkedVpnTunnels.include_import_ranges
	IncludeImportRanges []string `json:"includeImportRanges,omitempty"`
}

// +kcc:proto=google.cloud.networkconnectivity.v1.RouterApplianceInstance
type RouterApplianceInstance struct {
	// The URI of the VM.
	// +kcc:proto:field=google.cloud.networkconnectivity.v1.RouterApplianceInstance.virtual_machine
	VirtualMachine *string `json:"virtualMachine,omitempty"`

	// The IP address on the VM to use for peering.
	// +kcc:proto:field=google.cloud.networkconnectivity.v1.RouterApplianceInstance.ip_address
	IPAddress *string `json:"ipAddress,omitempty"`
}

// +kcc:proto=google.cloud.networkconnectivity.v1.Spoke
type Spoke struct {
	// Immutable. The name of the spoke. Spoke names must be unique. They use the
	//  following form:
	//      `projects/{project_number}/locations/{region}/spokes/{spoke_id}`
	// +kcc:proto:field=google.cloud.networkconnectivity.v1.Spoke.name
	Name *string `json:"name,omitempty"`

	// Optional labels in key-value pair format. For more information about
	//  labels, see [Requirements for
	//  labels](https://cloud.google.com/resource-manager/docs/creating-managing-labels#requirements).
	// +kcc:proto:field=google.cloud.networkconnectivity.v1.Spoke.labels
	Labels map[string]string `json:"labels,omitempty"`

	// An optional description of the spoke.
	// +kcc:proto:field=google.cloud.networkconnectivity.v1.Spoke.description
	Description *string `json:"description,omitempty"`

	// Immutable. The name of the hub that this spoke is attached to.
	// +kcc:proto:field=google.cloud.networkconnectivity.v1.Spoke.hub
	Hub *string `json:"hub,omitempty"`

	// Optional. The name of the group that this spoke is associated with.
	// +kcc:proto:field=google.cloud.networkconnectivity.v1.Spoke.group
	Group *string `json:"group,omitempty"`

	// VPN tunnels that are associated with the spoke.
	// +kcc:proto:field=google.cloud.networkconnectivity.v1.Spoke.linked_vpn_tunnels
	LinkedVpnTunnels *LinkedVpnTunnels `json:"linkedVpnTunnels,omitempty"`

	// VLAN attachments that are associated with the spoke.
	// +kcc:proto:field=google.cloud.networkconnectivity.v1.Spoke.linked_interconnect_attachments
	LinkedInterconnectAttachments *LinkedInterconnectAttachments `json:"linkedInterconnectAttachments,omitempty"`

	// Router appliance instances that are associated with the spoke.
	// +kcc:proto:field=google.cloud.networkconnectivity.v1.Spoke.linked_router_appliance_instances
	LinkedRouterApplianceInstances *LinkedRouterApplianceInstances `json:"linkedRouterApplianceInstances,omitempty"`

	// Optional. VPC network that is associated with the spoke.
	// +kcc:proto:field=google.cloud.networkconnectivity.v1.Spoke.linked_vpc_network
	LinkedVpcNetwork *LinkedVpcNetwork `json:"linkedVpcNetwork,omitempty"`

	// Optional. The linked producer VPC that is associated with the spoke.
	// +kcc:proto:field=google.cloud.networkconnectivity.v1.Spoke.linked_producer_vpc_network
	LinkedProducerVpcNetwork *LinkedProducerVpcNetwork `json:"linkedProducerVpcNetwork,omitempty"`
}

// +kcc:proto=google.cloud.networkconnectivity.v1.Spoke.StateReason
type Spoke_StateReason struct {
	// The code associated with this reason.
	// +kcc:proto:field=google.cloud.networkconnectivity.v1.Spoke.StateReason.code
	Code *string `json:"code,omitempty"`

	// Human-readable details about this reason.
	// +kcc:proto:field=google.cloud.networkconnectivity.v1.Spoke.StateReason.message
	Message *string `json:"message,omitempty"`

	// Additional information provided by the user in the RejectSpoke call.
	// +kcc:proto:field=google.cloud.networkconnectivity.v1.Spoke.StateReason.user_details
	UserDetails *string `json:"userDetails,omitempty"`
}

// +kcc:proto=google.cloud.networkconnectivity.v1.LinkedInterconnectAttachments
type LinkedInterconnectAttachmentsObservedState struct {
	// Output only. The VPC network where these VLAN attachments are located.
	// +kcc:proto:field=google.cloud.networkconnectivity.v1.LinkedInterconnectAttachments.vpc_network
	VpcNetwork *string `json:"vpcNetwork,omitempty"`
}

// +kcc:proto=google.cloud.networkconnectivity.v1.LinkedProducerVpcNetwork
type LinkedProducerVpcNetworkObservedState struct {
	// Output only. The Service Consumer Network spoke.
	// +kcc:proto:field=google.cloud.networkconnectivity.v1.LinkedProducerVpcNetwork.service_consumer_vpc_spoke
	ServiceConsumerVpcSpoke *string `json:"serviceConsumerVpcSpoke,omitempty"`

	// Output only. The URI of the Producer VPC.
	// +kcc:proto:field=google.cloud.networkconnectivity.v1.LinkedProducerVpcNetwork.producer_network
	ProducerNetwork *string `json:"producerNetwork,omitempty"`
}

// +kcc:proto=google.cloud.networkconnectivity.v1.LinkedRouterApplianceInstances
type LinkedRouterApplianceInstancesObservedState struct {
	// Output only. The VPC network where these router appliance instances are
	//  located.
	// +kcc:proto:field=google.cloud.networkconnectivity.v1.LinkedRouterApplianceInstances.vpc_network
	VpcNetwork *string `json:"vpcNetwork,omitempty"`
}

// +kcc:proto=google.cloud.networkconnectivity.v1.LinkedVpcNetwork
type LinkedVpcNetworkObservedState struct {
	// Output only. The list of Producer VPC spokes that this VPC spoke is a
	//  service consumer VPC spoke for. These producer VPCs are connected through
	//  VPC peering to this spoke's backing VPC network.
	// +kcc:proto:field=google.cloud.networkconnectivity.v1.LinkedVpcNetwork.producer_vpc_spokes
	ProducerVpcSpokes []string `json:"producerVpcSpokes,omitempty"`
}

// +kcc:proto=google.cloud.networkconnectivity.v1.LinkedVpnTunnels
type LinkedVpnTunnelsObservedState struct {
	// Output only. The VPC network where these VPN tunnels are located.
	// +kcc:proto:field=google.cloud.networkconnectivity.v1.LinkedVpnTunnels.vpc_network
	VpcNetwork *string `json:"vpcNetwork,omitempty"`
}

// +kcc:proto=google.cloud.networkconnectivity.v1.Spoke
type SpokeObservedState struct {
	// Output only. The time the spoke was created.
	// +kcc:proto:field=google.cloud.networkconnectivity.v1.Spoke.create_time
	CreateTime *string `json:"createTime,omitempty"`

	// Output only. The time the spoke was last updated.
	// +kcc:proto:field=google.cloud.networkconnectivity.v1.Spoke.update_time
	UpdateTime *string `json:"updateTime,omitempty"`

	// VPN tunnels that are associated with the spoke.
	// +kcc:proto:field=google.cloud.networkconnectivity.v1.Spoke.linked_vpn_tunnels
	LinkedVpnTunnels *LinkedVpnTunnelsObservedState `json:"linkedVpnTunnels,omitempty"`

	// VLAN attachments that are associated with the spoke.
	// +kcc:proto:field=google.cloud.networkconnectivity.v1.Spoke.linked_interconnect_attachments
	LinkedInterconnectAttachments *LinkedInterconnectAttachmentsObservedState `json:"linkedInterconnectAttachments,omitempty"`

	// Router appliance instances that are associated with the spoke.
	// +kcc:proto:field=google.cloud.networkconnectivity.v1.Spoke.linked_router_appliance_instances
	LinkedRouterApplianceInstances *LinkedRouterApplianceInstancesObservedState `json:"linkedRouterApplianceInstances,omitempty"`

	// Optional. VPC network that is associated with the spoke.
	// +kcc:proto:field=google.cloud.networkconnectivity.v1.Spoke.linked_vpc_network
	LinkedVpcNetwork *LinkedVpcNetworkObservedState `json:"linkedVpcNetwork,omitempty"`

	// Optional. The linked producer VPC that is associated with the spoke.
	// +kcc:proto:field=google.cloud.networkconnectivity.v1.Spoke.linked_producer_vpc_network
	LinkedProducerVpcNetwork *LinkedProducerVpcNetworkObservedState `json:"linkedProducerVpcNetwork,omitempty"`

	// Output only. The Google-generated UUID for the spoke. This value is unique
	//  across all spoke resources. If a spoke is deleted and another with the same
	//  name is created, the new spoke is assigned a different `unique_id`.
	// +kcc:proto:field=google.cloud.networkconnectivity.v1.Spoke.unique_id
	UniqueID *string `json:"uniqueID,omitempty"`

	// Output only. The current lifecycle state of this spoke.
	// +kcc:proto:field=google.cloud.networkconnectivity.v1.Spoke.state
	State *string `json:"state,omitempty"`

	// Output only. The reasons for current state of the spoke. Only present when
	//  the spoke is in the `INACTIVE` state.
	// +kcc:proto:field=google.cloud.networkconnectivity.v1.Spoke.reasons
	Reasons []Spoke_StateReason `json:"reasons,omitempty"`

	// Output only. The type of resource associated with the spoke.
	// +kcc:proto:field=google.cloud.networkconnectivity.v1.Spoke.spoke_type
	SpokeType *string `json:"spokeType,omitempty"`
}
