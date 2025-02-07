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


// +kcc:proto=google.cloud.baremetalsolution.v2.ServerNetworkTemplate
type ServerNetworkTemplate struct {

	// Instance types this template is applicable to.
	// +kcc:proto:field=google.cloud.baremetalsolution.v2.ServerNetworkTemplate.applicable_instance_types
	ApplicableInstanceTypes []string `json:"applicableInstanceTypes,omitempty"`

	// Logical interfaces.
	// +kcc:proto:field=google.cloud.baremetalsolution.v2.ServerNetworkTemplate.logical_interfaces
	LogicalInterfaces []ServerNetworkTemplate_LogicalInterface `json:"logicalInterfaces,omitempty"`
}

// +kcc:proto=google.cloud.baremetalsolution.v2.ServerNetworkTemplate.LogicalInterface
type ServerNetworkTemplate_LogicalInterface struct {
	// Interface name.
	//  This is not a globally unique identifier.
	//  Name is unique only inside the ServerNetworkTemplate. This is of syntax
	//  <bond><interface_type_index><bond_mode> or <nic><interface_type_index>
	//  and forms part of the network template name.
	// +kcc:proto:field=google.cloud.baremetalsolution.v2.ServerNetworkTemplate.LogicalInterface.name
	Name *string `json:"name,omitempty"`

	// Interface type.
	// +kcc:proto:field=google.cloud.baremetalsolution.v2.ServerNetworkTemplate.LogicalInterface.type
	Type *string `json:"type,omitempty"`

	// If true, interface must have network connected.
	// +kcc:proto:field=google.cloud.baremetalsolution.v2.ServerNetworkTemplate.LogicalInterface.required
	Required *bool `json:"required,omitempty"`
}

// +kcc:proto=google.cloud.baremetalsolution.v2.ServerNetworkTemplate
type ServerNetworkTemplateObservedState struct {
	// Output only. Template's unique name. The full resource name follows the
	//  pattern:
	//  `projects/{project}/locations/{location}/serverNetworkTemplate/{server_network_template}`
	//  Generally, the {server_network_template} follows the syntax of
	//  "bond<interface_type_index><bond_mode>" or "nic<interface_type_index>".
	// +kcc:proto:field=google.cloud.baremetalsolution.v2.ServerNetworkTemplate.name
	Name *string `json:"name,omitempty"`
}
