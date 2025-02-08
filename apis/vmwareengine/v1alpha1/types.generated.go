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


// +kcc:proto=google.cloud.vmwareengine.v1.Subnet
type Subnet struct {

	// The IP address range of the subnet in CIDR format '10.0.0.0/24'.
	// +kcc:proto:field=google.cloud.vmwareengine.v1.Subnet.ip_cidr_range
	IPCidrRange *string `json:"ipCidrRange,omitempty"`

	// The IP address of the gateway of this subnet.
	//  Must fall within the IP prefix defined above.
	// +kcc:proto:field=google.cloud.vmwareengine.v1.Subnet.gateway_ip
	GatewayIP *string `json:"gatewayIP,omitempty"`
}

// +kcc:proto=google.cloud.vmwareengine.v1.Subnet
type SubnetObservedState struct {
	// Output only. The resource name of this subnet.
	//  Resource names are schemeless URIs that follow the conventions in
	//  https://cloud.google.com/apis/design/resource_names.
	//  For example:
	//  `projects/my-project/locations/us-central1-a/privateClouds/my-cloud/subnets/my-subnet`
	// +kcc:proto:field=google.cloud.vmwareengine.v1.Subnet.name
	Name *string `json:"name,omitempty"`

	// Output only. The type of the subnet. For example "management" or
	//  "userDefined".
	// +kcc:proto:field=google.cloud.vmwareengine.v1.Subnet.type
	Type *string `json:"type,omitempty"`

	// Output only. The state of the resource.
	// +kcc:proto:field=google.cloud.vmwareengine.v1.Subnet.state
	State *string `json:"state,omitempty"`

	// Output only. VLAN ID of the VLAN on which the subnet is configured
	// +kcc:proto:field=google.cloud.vmwareengine.v1.Subnet.vlan_id
	VlanID *int32 `json:"vlanID,omitempty"`
}
