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
