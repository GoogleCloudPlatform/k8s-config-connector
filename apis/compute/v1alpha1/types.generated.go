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

// +kcc:proto=google.cloud.compute.v1.NetworkAttachment
type NetworkAttachment struct {
	// [Output Only] An array of connections for all the producers connected to this network attachment.
	// +kcc:proto:field=google.cloud.compute.v1.NetworkAttachment.connection_endpoints
	ConnectionEndpoints []NetworkAttachmentConnectedEndpoint `json:"connectionEndpoints,omitempty"`

	// Check the ConnectionPreference enum for the list of possible values.
	// +kcc:proto:field=google.cloud.compute.v1.NetworkAttachment.connection_preference
	ConnectionPreference *string `json:"connectionPreference,omitempty"`

	// [Output Only] Creation timestamp in RFC3339 text format.
	// +kcc:proto:field=google.cloud.compute.v1.NetworkAttachment.creation_timestamp
	CreationTimestamp *string `json:"creationTimestamp,omitempty"`

	// An optional description of this resource. Provide this property when you create the resource.
	// +kcc:proto:field=google.cloud.compute.v1.NetworkAttachment.description
	Description *string `json:"description,omitempty"`

	// Fingerprint of this resource. A hash of the contents stored in this object. This field is used in optimistic locking. An up-to-date fingerprint must be provided in order to patch.
	// +kcc:proto:field=google.cloud.compute.v1.NetworkAttachment.fingerprint
	Fingerprint *string `json:"fingerprint,omitempty"`

	// [Output Only] The unique identifier for the resource type. The server generates this identifier.
	// +kcc:proto:field=google.cloud.compute.v1.NetworkAttachment.id
	ID *uint64 `json:"id,omitempty"`

	// [Output Only] Type of the resource.
	// +kcc:proto:field=google.cloud.compute.v1.NetworkAttachment.kind
	Kind *string `json:"kind,omitempty"`

	// Name of the resource. Provided by the client when the resource is created. The name must be 1-63 characters long, and comply with RFC1035. Specifically, the name must be 1-63 characters long and match the regular expression `[a-z]([-a-z0-9]*[a-z0-9])?` which means the first character must be a lowercase letter, and all following characters must be a dash, lowercase letter, or digit, except the last character, which cannot be a dash.
	// +kcc:proto:field=google.cloud.compute.v1.NetworkAttachment.name
	Name *string `json:"name,omitempty"`

	// [Output Only] The URL of the network which the Network Attachment belongs to. Practically it is inferred by fetching the network of the first subnetwork associated. Because it is required that all the subnetworks must be from the same network, it is assured that the Network Attachment belongs to the same network as all the subnetworks.
	// +kcc:proto:field=google.cloud.compute.v1.NetworkAttachment.network
	Network *string `json:"network,omitempty"`

	// Projects that are allowed to connect to this network attachment. The project can be specified using its id or number.
	// +kcc:proto:field=google.cloud.compute.v1.NetworkAttachment.producer_accept_lists
	ProducerAcceptLists []string `json:"producerAcceptLists,omitempty"`

	// Projects that are not allowed to connect to this network attachment. The project can be specified using its id or number.
	// +kcc:proto:field=google.cloud.compute.v1.NetworkAttachment.producer_reject_lists
	ProducerRejectLists []string `json:"producerRejectLists,omitempty"`

	// [Output Only] URL of the region where the network attachment resides. This field applies only to the region resource. You must specify this field as part of the HTTP request URL. It is not settable as a field in the request body.
	// +kcc:proto:field=google.cloud.compute.v1.NetworkAttachment.region
	Region *string `json:"region,omitempty"`

	// [Output Only] Server-defined URL for the resource.
	// +kcc:proto:field=google.cloud.compute.v1.NetworkAttachment.self_link
	SelfLink *string `json:"selfLink,omitempty"`

	// [Output Only] Server-defined URL for this resource's resource id.
	// +kcc:proto:field=google.cloud.compute.v1.NetworkAttachment.self_link_with_id
	SelfLinkWithID *string `json:"selfLinkWithID,omitempty"`

	// An array of URLs where each entry is the URL of a subnet provided by the service consumer to use for endpoints in the producers that connect to this network attachment.
	// +kcc:proto:field=google.cloud.compute.v1.NetworkAttachment.subnetworks
	Subnetworks []string `json:"subnetworks,omitempty"`
}

// +kcc:proto=google.cloud.compute.v1.NetworkAttachmentConnectedEndpoint
type NetworkAttachmentConnectedEndpoint struct {
	// The IPv4 address assigned to the producer instance network interface. This value will be a range in case of Serverless.
	// +kcc:proto:field=google.cloud.compute.v1.NetworkAttachmentConnectedEndpoint.ip_address
	IPAddress *string `json:"ipAddress,omitempty"`

	// The IPv6 address assigned to the producer instance network interface. This is only assigned when the stack types of both the instance network interface and the consumer subnet are IPv4_IPv6.
	// +kcc:proto:field=google.cloud.compute.v1.NetworkAttachmentConnectedEndpoint.ipv6_address
	IPV6Address *string `json:"ipv6Address,omitempty"`

	// The project id or number of the interface to which the IP was assigned.
	// +kcc:proto:field=google.cloud.compute.v1.NetworkAttachmentConnectedEndpoint.project_id_or_num
	ProjectIDOrNum *string `json:"projectIDOrNum,omitempty"`

	// Alias IP ranges from the same subnetwork.
	// +kcc:proto:field=google.cloud.compute.v1.NetworkAttachmentConnectedEndpoint.secondary_ip_cidr_ranges
	SecondaryIPCIDRRanges []string `json:"secondaryIPCIDRRanges,omitempty"`

	// The status of a connected endpoint to this network attachment.
	//  Check the Status enum for the list of possible values.
	// +kcc:proto:field=google.cloud.compute.v1.NetworkAttachmentConnectedEndpoint.status
	Status *string `json:"status,omitempty"`

	// The subnetwork used to assign the IP to the producer instance network interface.
	// +kcc:proto:field=google.cloud.compute.v1.NetworkAttachmentConnectedEndpoint.subnetwork
	Subnetwork *string `json:"subnetwork,omitempty"`

	// [Output Only] The CIDR range of the subnet from which the IPv4 internal IP was allocated from.
	// +kcc:proto:field=google.cloud.compute.v1.NetworkAttachmentConnectedEndpoint.subnetwork_cidr_range
	SubnetworkCIDRRange *string `json:"subnetworkCIDRRange,omitempty"`
}
