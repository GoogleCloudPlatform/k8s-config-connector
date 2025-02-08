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


// +kcc:proto=google.cloud.edgenetwork.v1.InterconnectAttachment
type InterconnectAttachment struct {
	// Required. The canonical resource name of the interconnect attachment.
	// +kcc:proto:field=google.cloud.edgenetwork.v1.InterconnectAttachment.name
	Name *string `json:"name,omitempty"`

	// Labels associated with this resource.
	// +kcc:proto:field=google.cloud.edgenetwork.v1.InterconnectAttachment.labels
	Labels map[string]string `json:"labels,omitempty"`

	// Optional. A free-text description of the resource. Max length 1024
	//  characters.
	// +kcc:proto:field=google.cloud.edgenetwork.v1.InterconnectAttachment.description
	Description *string `json:"description,omitempty"`

	// Required. The canonical name of underlying Interconnect object that this
	//  attachment's traffic will traverse through. The name is in the form of
	//  `projects/{project}/locations/{location}/zones/{zone}/interconnects/{interconnect}`.
	// +kcc:proto:field=google.cloud.edgenetwork.v1.InterconnectAttachment.interconnect
	Interconnect *string `json:"interconnect,omitempty"`

	// Optional. The canonical Network name in the form of
	//  `projects/{project}/locations/{location}/zones/{zone}/networks/{network}`.
	// +kcc:proto:field=google.cloud.edgenetwork.v1.InterconnectAttachment.network
	Network *string `json:"network,omitempty"`

	// Required. VLAN id provided by user. Must be site-wise unique.
	// +kcc:proto:field=google.cloud.edgenetwork.v1.InterconnectAttachment.vlan_id
	VlanID *int32 `json:"vlanID,omitempty"`

	// IP (L3) MTU value of the virtual edge cloud.
	//  Valid values are: 1500 and 9000.
	//  Default to 1500 if not set.
	// +kcc:proto:field=google.cloud.edgenetwork.v1.InterconnectAttachment.mtu
	Mtu *int32 `json:"mtu,omitempty"`
}

// +kcc:proto=google.cloud.edgenetwork.v1.InterconnectAttachment
type InterconnectAttachmentObservedState struct {
	// Output only. The time when the interconnect attachment was created.
	// +kcc:proto:field=google.cloud.edgenetwork.v1.InterconnectAttachment.create_time
	CreateTime *string `json:"createTime,omitempty"`

	// Output only. The time when the interconnect attachment was last updated.
	// +kcc:proto:field=google.cloud.edgenetwork.v1.InterconnectAttachment.update_time
	UpdateTime *string `json:"updateTime,omitempty"`

	// Output only. Current stage of the resource to the device by config push.
	// +kcc:proto:field=google.cloud.edgenetwork.v1.InterconnectAttachment.state
	State *string `json:"state,omitempty"`
}
