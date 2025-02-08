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


// +kcc:proto=google.cloud.networkconnectivity.v1.PolicyBasedRoute
type PolicyBasedRoute struct {
	// Optional. VM instances to which this policy based route applies to.
	// +kcc:proto:field=google.cloud.networkconnectivity.v1.PolicyBasedRoute.virtual_machine
	VirtualMachine *PolicyBasedRoute_VirtualMachine `json:"virtualMachine,omitempty"`

	// Optional. The interconnect attachments to which this route applies to.
	// +kcc:proto:field=google.cloud.networkconnectivity.v1.PolicyBasedRoute.interconnect_attachment
	InterconnectAttachment *PolicyBasedRoute_InterconnectAttachment `json:"interconnectAttachment,omitempty"`

	// Optional. The IP of a global access enabled L4 ILB that should be the
	//  next hop to handle matching packets. For this version, only
	//  next_hop_ilb_ip is supported.
	// +kcc:proto:field=google.cloud.networkconnectivity.v1.PolicyBasedRoute.next_hop_ilb_ip
	NextHopIlbIP *string `json:"nextHopIlbIP,omitempty"`

	// Optional. Other routes that will be referenced to determine the next hop
	//  of the packet.
	// +kcc:proto:field=google.cloud.networkconnectivity.v1.PolicyBasedRoute.next_hop_other_routes
	NextHopOtherRoutes *string `json:"nextHopOtherRoutes,omitempty"`

	// Immutable. A unique name of the resource in the form of
	//  `projects/{project_number}/locations/global/PolicyBasedRoutes/{policy_based_route_id}`
	// +kcc:proto:field=google.cloud.networkconnectivity.v1.PolicyBasedRoute.name
	Name *string `json:"name,omitempty"`

	// User-defined labels.
	// +kcc:proto:field=google.cloud.networkconnectivity.v1.PolicyBasedRoute.labels
	Labels map[string]string `json:"labels,omitempty"`

	// Optional. An optional description of this resource. Provide this field when
	//  you create the resource.
	// +kcc:proto:field=google.cloud.networkconnectivity.v1.PolicyBasedRoute.description
	Description *string `json:"description,omitempty"`

	// Required. Fully-qualified URL of the network that this route applies to.
	//  e.g. projects/my-project/global/networks/my-network.
	// +kcc:proto:field=google.cloud.networkconnectivity.v1.PolicyBasedRoute.network
	Network *string `json:"network,omitempty"`

	// Required. The filter to match L4 traffic.
	// +kcc:proto:field=google.cloud.networkconnectivity.v1.PolicyBasedRoute.filter
	Filter *PolicyBasedRoute_Filter `json:"filter,omitempty"`

	// Optional. The priority of this policy based route. Priority is used to
	//  break ties in cases where there are more than one matching policy based
	//  routes found. In cases where multiple policy based routes are matched, the
	//  one with the lowest-numbered priority value wins. The default value is
	//  1000. The priority value must be from 1 to 65535, inclusive.
	// +kcc:proto:field=google.cloud.networkconnectivity.v1.PolicyBasedRoute.priority
	Priority *int32 `json:"priority,omitempty"`
}

// +kcc:proto=google.cloud.networkconnectivity.v1.PolicyBasedRoute.Filter
type PolicyBasedRoute_Filter struct {
	// Optional. The IP protocol that this policy based route applies to. Valid
	//  values are 'TCP', 'UDP', and 'ALL'. Default is 'ALL'.
	// +kcc:proto:field=google.cloud.networkconnectivity.v1.PolicyBasedRoute.Filter.ip_protocol
	IPProtocol *string `json:"ipProtocol,omitempty"`

	// Optional. The source IP range of outgoing packets that this policy based
	//  route applies to. Default is "0.0.0.0/0" if protocol version is IPv4.
	// +kcc:proto:field=google.cloud.networkconnectivity.v1.PolicyBasedRoute.Filter.src_range
	SrcRange *string `json:"srcRange,omitempty"`

	// Optional. The destination IP range of outgoing packets that this policy
	//  based route applies to. Default is "0.0.0.0/0" if protocol version is
	//  IPv4.
	// +kcc:proto:field=google.cloud.networkconnectivity.v1.PolicyBasedRoute.Filter.dest_range
	DestRange *string `json:"destRange,omitempty"`

	// Required. Internet protocol versions this policy based route applies to.
	//  For this version, only IPV4 is supported.
	// +kcc:proto:field=google.cloud.networkconnectivity.v1.PolicyBasedRoute.Filter.protocol_version
	ProtocolVersion *string `json:"protocolVersion,omitempty"`
}

// +kcc:proto=google.cloud.networkconnectivity.v1.PolicyBasedRoute.InterconnectAttachment
type PolicyBasedRoute_InterconnectAttachment struct {
	// Optional. Cloud region to install this policy based route on interconnect
	//  attachment. Use `all` to install it on all interconnect attachments.
	// +kcc:proto:field=google.cloud.networkconnectivity.v1.PolicyBasedRoute.InterconnectAttachment.region
	Region *string `json:"region,omitempty"`
}

// +kcc:proto=google.cloud.networkconnectivity.v1.PolicyBasedRoute.VirtualMachine
type PolicyBasedRoute_VirtualMachine struct {
	// Optional. A list of VM instance tags to which this policy based route
	//  applies to. VM instances that have ANY of tags specified here will
	//  install this PBR.
	// +kcc:proto:field=google.cloud.networkconnectivity.v1.PolicyBasedRoute.VirtualMachine.tags
	Tags []string `json:"tags,omitempty"`
}

// +kcc:proto=google.cloud.networkconnectivity.v1.PolicyBasedRoute.Warnings
type PolicyBasedRoute_Warnings struct {
}

// +kcc:proto=google.cloud.networkconnectivity.v1.PolicyBasedRoute
type PolicyBasedRouteObservedState struct {
	// Output only. Time when the PolicyBasedRoute was created.
	// +kcc:proto:field=google.cloud.networkconnectivity.v1.PolicyBasedRoute.create_time
	CreateTime *string `json:"createTime,omitempty"`

	// Output only. Time when the PolicyBasedRoute was updated.
	// +kcc:proto:field=google.cloud.networkconnectivity.v1.PolicyBasedRoute.update_time
	UpdateTime *string `json:"updateTime,omitempty"`

	// Output only. If potential misconfigurations are detected for this route,
	//  this field will be populated with warning messages.
	// +kcc:proto:field=google.cloud.networkconnectivity.v1.PolicyBasedRoute.warnings
	Warnings []PolicyBasedRoute_Warnings `json:"warnings,omitempty"`

	// Output only. Server-defined fully-qualified URL for this resource.
	// +kcc:proto:field=google.cloud.networkconnectivity.v1.PolicyBasedRoute.self_link
	SelfLink *string `json:"selfLink,omitempty"`

	// Output only. Type of this resource. Always
	//  networkconnectivity#policyBasedRoute for Policy Based Route resources.
	// +kcc:proto:field=google.cloud.networkconnectivity.v1.PolicyBasedRoute.kind
	Kind *string `json:"kind,omitempty"`
}

// +kcc:proto=google.cloud.networkconnectivity.v1.PolicyBasedRoute.Warnings
type PolicyBasedRoute_WarningsObservedState struct {
	// Output only. A warning code, if applicable.
	// +kcc:proto:field=google.cloud.networkconnectivity.v1.PolicyBasedRoute.Warnings.code
	Code *string `json:"code,omitempty"`

	// Output only. Metadata about this warning in key: value format. The key
	//  should provides more detail on the warning being returned. For example,
	//  for warnings where there are no results in a list request for a
	//  particular zone, this key might be scope and the key value might be the
	//  zone name. Other examples might be a key indicating a deprecated resource
	//  and a suggested replacement.
	// +kcc:proto:field=google.cloud.networkconnectivity.v1.PolicyBasedRoute.Warnings.data
	Data map[string]string `json:"data,omitempty"`

	// Output only. A human-readable description of the warning code.
	// +kcc:proto:field=google.cloud.networkconnectivity.v1.PolicyBasedRoute.Warnings.warning_message
	WarningMessage *string `json:"warningMessage,omitempty"`
}
