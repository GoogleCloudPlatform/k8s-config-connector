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


// +kcc:proto=google.cloud.baremetalsolution.v2.InstanceConfig
type InstanceConfig struct {

	// A transient unique identifier to idenfity an instance within an
	//  ProvisioningConfig request.
	// +kcc:proto:field=google.cloud.baremetalsolution.v2.InstanceConfig.id
	ID *string `json:"id,omitempty"`

	// Instance type.
	//  [Available
	//  types](https://cloud.google.com/bare-metal/docs/bms-planning#server_configurations)
	// +kcc:proto:field=google.cloud.baremetalsolution.v2.InstanceConfig.instance_type
	InstanceType *string `json:"instanceType,omitempty"`

	// Whether the instance should be provisioned with Hyperthreading enabled.
	// +kcc:proto:field=google.cloud.baremetalsolution.v2.InstanceConfig.hyperthreading
	Hyperthreading *bool `json:"hyperthreading,omitempty"`

	// OS image to initialize the instance.
	//  [Available
	//  images](https://cloud.google.com/bare-metal/docs/bms-planning#server_configurations)
	// +kcc:proto:field=google.cloud.baremetalsolution.v2.InstanceConfig.os_image
	OsImage *string `json:"osImage,omitempty"`

	// Client network address. Filled if InstanceConfig.multivlan_config is false.
	// +kcc:proto:field=google.cloud.baremetalsolution.v2.InstanceConfig.client_network
	ClientNetwork *InstanceConfig_NetworkAddress `json:"clientNetwork,omitempty"`

	// Private network address, if any. Filled if InstanceConfig.multivlan_config
	//  is false.
	// +kcc:proto:field=google.cloud.baremetalsolution.v2.InstanceConfig.private_network
	PrivateNetwork *InstanceConfig_NetworkAddress `json:"privateNetwork,omitempty"`

	// User note field, it can be used by customers to add additional information
	//  for the BMS Ops team .
	// +kcc:proto:field=google.cloud.baremetalsolution.v2.InstanceConfig.user_note
	UserNote *string `json:"userNote,omitempty"`

	// If true networks can be from different projects of the same vendor account.
	// +kcc:proto:field=google.cloud.baremetalsolution.v2.InstanceConfig.account_networks_enabled
	AccountNetworksEnabled *bool `json:"accountNetworksEnabled,omitempty"`

	// The type of network configuration on the instance.
	// +kcc:proto:field=google.cloud.baremetalsolution.v2.InstanceConfig.network_config
	NetworkConfig *string `json:"networkConfig,omitempty"`

	// Server network template name. Filled if InstanceConfig.multivlan_config is
	//  true.
	// +kcc:proto:field=google.cloud.baremetalsolution.v2.InstanceConfig.network_template
	NetworkTemplate *string `json:"networkTemplate,omitempty"`

	// List of logical interfaces for the instance. The number of logical
	//  interfaces will be the same as number of hardware bond/nic on the chosen
	//  network template. Filled if InstanceConfig.multivlan_config is true.
	// +kcc:proto:field=google.cloud.baremetalsolution.v2.InstanceConfig.logical_interfaces
	LogicalInterfaces []LogicalInterface `json:"logicalInterfaces,omitempty"`

	// List of names of ssh keys used to provision the instance.
	// +kcc:proto:field=google.cloud.baremetalsolution.v2.InstanceConfig.ssh_key_names
	SSHKeyNames []string `json:"sshKeyNames,omitempty"`
}

// +kcc:proto=google.cloud.baremetalsolution.v2.InstanceConfig.NetworkAddress
type InstanceConfig_NetworkAddress struct {
	// Id of the network to use, within the same ProvisioningConfig request.
	// +kcc:proto:field=google.cloud.baremetalsolution.v2.InstanceConfig.NetworkAddress.network_id
	NetworkID *string `json:"networkID,omitempty"`

	// IPv4 address to be assigned to the server.
	// +kcc:proto:field=google.cloud.baremetalsolution.v2.InstanceConfig.NetworkAddress.address
	Address *string `json:"address,omitempty"`

	// Name of the existing network to use.
	// +kcc:proto:field=google.cloud.baremetalsolution.v2.InstanceConfig.NetworkAddress.existing_network_id
	ExistingNetworkID *string `json:"existingNetworkID,omitempty"`
}

// +kcc:proto=google.cloud.baremetalsolution.v2.LogicalInterface
type LogicalInterface struct {
	// List of logical network interfaces within a logical interface.
	// +kcc:proto:field=google.cloud.baremetalsolution.v2.LogicalInterface.logical_network_interfaces
	LogicalNetworkInterfaces []LogicalInterface_LogicalNetworkInterface `json:"logicalNetworkInterfaces,omitempty"`

	// Interface name. This is of syntax <bond><bond_mode> or <nic> and
	//  forms part of the network template name.
	// +kcc:proto:field=google.cloud.baremetalsolution.v2.LogicalInterface.name
	Name *string `json:"name,omitempty"`

	// The index of the logical interface mapping to the index of the hardware
	//  bond or nic on the chosen network template. This field is deprecated.
	// +kcc:proto:field=google.cloud.baremetalsolution.v2.LogicalInterface.interface_index
	InterfaceIndex *int32 `json:"interfaceIndex,omitempty"`
}

// +kcc:proto=google.cloud.baremetalsolution.v2.LogicalInterface.LogicalNetworkInterface
type LogicalInterface_LogicalNetworkInterface struct {
	// Name of the network
	// +kcc:proto:field=google.cloud.baremetalsolution.v2.LogicalInterface.LogicalNetworkInterface.network
	Network *string `json:"network,omitempty"`

	// IP address in the network
	// +kcc:proto:field=google.cloud.baremetalsolution.v2.LogicalInterface.LogicalNetworkInterface.ip_address
	IPAddress *string `json:"ipAddress,omitempty"`

	// Whether this interface is the default gateway for the instance. Only
	//  one interface can be the default gateway for the instance.
	// +kcc:proto:field=google.cloud.baremetalsolution.v2.LogicalInterface.LogicalNetworkInterface.default_gateway
	DefaultGateway *bool `json:"defaultGateway,omitempty"`

	// Type of network.
	// +kcc:proto:field=google.cloud.baremetalsolution.v2.LogicalInterface.LogicalNetworkInterface.network_type
	NetworkType *string `json:"networkType,omitempty"`

	// An identifier for the `Network`, generated by the backend.
	// +kcc:proto:field=google.cloud.baremetalsolution.v2.LogicalInterface.LogicalNetworkInterface.id
	ID *string `json:"id,omitempty"`
}

// +kcc:proto=google.cloud.baremetalsolution.v2.InstanceConfig
type InstanceConfigObservedState struct {
	// Output only. The name of the instance config.
	// +kcc:proto:field=google.cloud.baremetalsolution.v2.InstanceConfig.name
	Name *string `json:"name,omitempty"`
}
