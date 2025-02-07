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


// +kcc:proto=google.cloud.baremetalsolution.v2.NetworkConfig
type NetworkConfig struct {

	// A transient unique identifier to identify a volume within an
	//  ProvisioningConfig request.
	// +kcc:proto:field=google.cloud.baremetalsolution.v2.NetworkConfig.id
	ID *string `json:"id,omitempty"`

	// The type of this network, either Client or Private.
	// +kcc:proto:field=google.cloud.baremetalsolution.v2.NetworkConfig.type
	Type *string `json:"type,omitempty"`

	// Interconnect bandwidth. Set only when type is CLIENT.
	// +kcc:proto:field=google.cloud.baremetalsolution.v2.NetworkConfig.bandwidth
	Bandwidth *string `json:"bandwidth,omitempty"`

	// List of VLAN attachments. As of now there are always 2 attachments, but it
	//  is going to change in  the future (multi vlan).
	// +kcc:proto:field=google.cloud.baremetalsolution.v2.NetworkConfig.vlan_attachments
	VlanAttachments []NetworkConfig_IntakeVlanAttachment `json:"vlanAttachments,omitempty"`

	// CIDR range of the network.
	// +kcc:proto:field=google.cloud.baremetalsolution.v2.NetworkConfig.cidr
	Cidr *string `json:"cidr,omitempty"`

	// Service CIDR, if any.
	// +kcc:proto:field=google.cloud.baremetalsolution.v2.NetworkConfig.service_cidr
	ServiceCidr *string `json:"serviceCidr,omitempty"`

	// User note field, it can be used by customers to add additional information
	//  for the BMS Ops team .
	// +kcc:proto:field=google.cloud.baremetalsolution.v2.NetworkConfig.user_note
	UserNote *string `json:"userNote,omitempty"`

	// The GCP service of the network. Available gcp_service are in
	//  https://cloud.google.com/bare-metal/docs/bms-planning.
	// +kcc:proto:field=google.cloud.baremetalsolution.v2.NetworkConfig.gcp_service
	GcpService *string `json:"gcpService,omitempty"`

	// Whether the VLAN attachment pair is located in the same project.
	// +kcc:proto:field=google.cloud.baremetalsolution.v2.NetworkConfig.vlan_same_project
	VlanSameProject *bool `json:"vlanSameProject,omitempty"`

	// The JumboFramesEnabled option for customer to set.
	// +kcc:proto:field=google.cloud.baremetalsolution.v2.NetworkConfig.jumbo_frames_enabled
	JumboFramesEnabled *bool `json:"jumboFramesEnabled,omitempty"`
}

// +kcc:proto=google.cloud.baremetalsolution.v2.NetworkConfig.IntakeVlanAttachment
type NetworkConfig_IntakeVlanAttachment struct {
	// Identifier of the VLAN attachment.
	// +kcc:proto:field=google.cloud.baremetalsolution.v2.NetworkConfig.IntakeVlanAttachment.id
	ID *string `json:"id,omitempty"`

	// Attachment pairing key.
	// +kcc:proto:field=google.cloud.baremetalsolution.v2.NetworkConfig.IntakeVlanAttachment.pairing_key
	PairingKey *string `json:"pairingKey,omitempty"`
}

// +kcc:proto=google.cloud.baremetalsolution.v2.NetworkConfig
type NetworkConfigObservedState struct {
	// Output only. The name of the network config.
	// +kcc:proto:field=google.cloud.baremetalsolution.v2.NetworkConfig.name
	Name *string `json:"name,omitempty"`
}
