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


// +kcc:proto=google.cloud.edgenetwork.v1.Subnet
type Subnet struct {
	// Required. The canonical resource name of the subnet.
	// +kcc:proto:field=google.cloud.edgenetwork.v1.Subnet.name
	Name *string `json:"name,omitempty"`

	// Labels associated with this resource.
	// +kcc:proto:field=google.cloud.edgenetwork.v1.Subnet.labels
	Labels map[string]string `json:"labels,omitempty"`

	// Optional. A free-text description of the resource. Max length 1024
	//  characters.
	// +kcc:proto:field=google.cloud.edgenetwork.v1.Subnet.description
	Description *string `json:"description,omitempty"`

	// Required. The network that this subnetwork belongs to.
	// +kcc:proto:field=google.cloud.edgenetwork.v1.Subnet.network
	Network *string `json:"network,omitempty"`

	// The ranges of ipv4 addresses that are owned by this subnetwork.
	// +kcc:proto:field=google.cloud.edgenetwork.v1.Subnet.ipv4_cidr
	Ipv4Cidr []string `json:"ipv4Cidr,omitempty"`

	// The ranges of ipv6 addresses that are owned by this subnetwork.
	// +kcc:proto:field=google.cloud.edgenetwork.v1.Subnet.ipv6_cidr
	Ipv6Cidr []string `json:"ipv6Cidr,omitempty"`

	// Optional. VLAN id provided by user. If not specified we assign one
	//  automatically.
	// +kcc:proto:field=google.cloud.edgenetwork.v1.Subnet.vlan_id
	VlanID *int32 `json:"vlanID,omitempty"`

	// Optional. A bonding type in the subnet creation specifies whether a VLAN
	//  being created will be present on Bonded or Non-Bonded or Both port types.
	//  In addition, this flag is to be used to set the specific network
	//  configuration which clusters can then use for their workloads based on the
	//  bonding choice.
	// +kcc:proto:field=google.cloud.edgenetwork.v1.Subnet.bonding_type
	BondingType *string `json:"bondingType,omitempty"`
}

// +kcc:proto=google.cloud.edgenetwork.v1.Subnet
type SubnetObservedState struct {
	// Output only. The time when the subnet was created.
	// +kcc:proto:field=google.cloud.edgenetwork.v1.Subnet.create_time
	CreateTime *string `json:"createTime,omitempty"`

	// Output only. The time when the subnet was last updated.
	// +kcc:proto:field=google.cloud.edgenetwork.v1.Subnet.update_time
	UpdateTime *string `json:"updateTime,omitempty"`

	// Output only. Current stage of the resource to the device by config push.
	// +kcc:proto:field=google.cloud.edgenetwork.v1.Subnet.state
	State *string `json:"state,omitempty"`
}
