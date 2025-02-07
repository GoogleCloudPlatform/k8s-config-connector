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

// +kcc:proto=google.cloud.baremetalsolution.v2.ProvisioningConfig
type ProvisioningConfig struct {

	// Instances to be created.
	// +kcc:proto:field=google.cloud.baremetalsolution.v2.ProvisioningConfig.instances
	Instances []InstanceConfig `json:"instances,omitempty"`

	// Networks to be created.
	// +kcc:proto:field=google.cloud.baremetalsolution.v2.ProvisioningConfig.networks
	Networks []NetworkConfig `json:"networks,omitempty"`

	// Volumes to be created.
	// +kcc:proto:field=google.cloud.baremetalsolution.v2.ProvisioningConfig.volumes
	Volumes []VolumeConfig `json:"volumes,omitempty"`

	// A generated ticket id to track provisioning request.
	// +kcc:proto:field=google.cloud.baremetalsolution.v2.ProvisioningConfig.ticket_id
	TicketID *string `json:"ticketID,omitempty"`

	// A service account to enable customers to access instance credentials upon
	//  handover.
	// +kcc:proto:field=google.cloud.baremetalsolution.v2.ProvisioningConfig.handover_service_account
	HandoverServiceAccount *string `json:"handoverServiceAccount,omitempty"`

	// Email provided to send a confirmation with provisioning config to.
	//  Deprecated in favour of email field in request messages.
	// +kcc:proto:field=google.cloud.baremetalsolution.v2.ProvisioningConfig.email
	Email *string `json:"email,omitempty"`

	// Optional. Location name of this ProvisioningConfig.
	//  It is optional only for Intake UI transition period.
	// +kcc:proto:field=google.cloud.baremetalsolution.v2.ProvisioningConfig.location
	Location *string `json:"location,omitempty"`

	// If true, VPC SC is enabled for the cluster.
	// +kcc:proto:field=google.cloud.baremetalsolution.v2.ProvisioningConfig.vpc_sc_enabled
	VpcScEnabled *bool `json:"vpcScEnabled,omitempty"`

	// Optional status messages associated with the FAILED state.
	// +kcc:proto:field=google.cloud.baremetalsolution.v2.ProvisioningConfig.status_message
	StatusMessage *string `json:"statusMessage,omitempty"`

	// Optional. The user-defined identifier of the provisioning config.
	// +kcc:proto:field=google.cloud.baremetalsolution.v2.ProvisioningConfig.custom_id
	CustomID *string `json:"customID,omitempty"`
}

// +kcc:proto=google.cloud.baremetalsolution.v2.VolumeConfig
type VolumeConfig struct {

	// A transient unique identifier to identify a volume within an
	//  ProvisioningConfig request.
	// +kcc:proto:field=google.cloud.baremetalsolution.v2.VolumeConfig.id
	ID *string `json:"id,omitempty"`

	// Whether snapshots should be enabled.
	// +kcc:proto:field=google.cloud.baremetalsolution.v2.VolumeConfig.snapshots_enabled
	SnapshotsEnabled *bool `json:"snapshotsEnabled,omitempty"`

	// The type of this Volume.
	// +kcc:proto:field=google.cloud.baremetalsolution.v2.VolumeConfig.type
	Type *string `json:"type,omitempty"`

	// Volume protocol.
	// +kcc:proto:field=google.cloud.baremetalsolution.v2.VolumeConfig.protocol
	Protocol *string `json:"protocol,omitempty"`

	// The requested size of this volume, in GB.
	// +kcc:proto:field=google.cloud.baremetalsolution.v2.VolumeConfig.size_gb
	SizeGB *int32 `json:"sizeGB,omitempty"`

	// LUN ranges to be configured. Set only when protocol is PROTOCOL_FC.
	// +kcc:proto:field=google.cloud.baremetalsolution.v2.VolumeConfig.lun_ranges
	LunRanges []VolumeConfig_LunRange `json:"lunRanges,omitempty"`

	// Machine ids connected to this volume. Set only when protocol is
	//  PROTOCOL_FC.
	// +kcc:proto:field=google.cloud.baremetalsolution.v2.VolumeConfig.machine_ids
	MachineIds []string `json:"machineIds,omitempty"`

	// NFS exports. Set only when protocol is PROTOCOL_NFS.
	// +kcc:proto:field=google.cloud.baremetalsolution.v2.VolumeConfig.nfs_exports
	NfsExports []VolumeConfig_NfsExport `json:"nfsExports,omitempty"`

	// User note field, it can be used by customers to add additional information
	//  for the BMS Ops team .
	// +kcc:proto:field=google.cloud.baremetalsolution.v2.VolumeConfig.user_note
	UserNote *string `json:"userNote,omitempty"`

	// The GCP service of the storage volume. Available gcp_service are in
	//  https://cloud.google.com/bare-metal/docs/bms-planning.
	// +kcc:proto:field=google.cloud.baremetalsolution.v2.VolumeConfig.gcp_service
	GcpService *string `json:"gcpService,omitempty"`

	// Performance tier of the Volume.
	//  Default is SHARED.
	// +kcc:proto:field=google.cloud.baremetalsolution.v2.VolumeConfig.performance_tier
	PerformanceTier *string `json:"performanceTier,omitempty"`
}

// +kcc:proto=google.cloud.baremetalsolution.v2.VolumeConfig.LunRange
type VolumeConfig_LunRange struct {
	// Number of LUNs to create.
	// +kcc:proto:field=google.cloud.baremetalsolution.v2.VolumeConfig.LunRange.quantity
	Quantity *int32 `json:"quantity,omitempty"`

	// The requested size of each LUN, in GB.
	// +kcc:proto:field=google.cloud.baremetalsolution.v2.VolumeConfig.LunRange.size_gb
	SizeGB *int32 `json:"sizeGB,omitempty"`
}

// +kcc:proto=google.cloud.baremetalsolution.v2.VolumeConfig.NfsExport
type VolumeConfig_NfsExport struct {
	// Network to use to publish the export.
	// +kcc:proto:field=google.cloud.baremetalsolution.v2.VolumeConfig.NfsExport.network_id
	NetworkID *string `json:"networkID,omitempty"`

	// Either a single machine, identified by an ID, or a comma-separated
	//  list of machine IDs.
	// +kcc:proto:field=google.cloud.baremetalsolution.v2.VolumeConfig.NfsExport.machine_id
	MachineID *string `json:"machineID,omitempty"`

	// A CIDR range.
	// +kcc:proto:field=google.cloud.baremetalsolution.v2.VolumeConfig.NfsExport.cidr
	Cidr *string `json:"cidr,omitempty"`

	// Export permissions.
	// +kcc:proto:field=google.cloud.baremetalsolution.v2.VolumeConfig.NfsExport.permissions
	Permissions *string `json:"permissions,omitempty"`

	// Disable root squashing, which is a feature of NFS.
	//  Root squash is a special mapping of the remote superuser (root) identity
	//  when using identity authentication.
	// +kcc:proto:field=google.cloud.baremetalsolution.v2.VolumeConfig.NfsExport.no_root_squash
	NoRootSquash *bool `json:"noRootSquash,omitempty"`

	// Allow the setuid flag.
	// +kcc:proto:field=google.cloud.baremetalsolution.v2.VolumeConfig.NfsExport.allow_suid
	AllowSuid *bool `json:"allowSuid,omitempty"`

	// Allow dev flag in NfsShare AllowedClientsRequest.
	// +kcc:proto:field=google.cloud.baremetalsolution.v2.VolumeConfig.NfsExport.allow_dev
	AllowDev *bool `json:"allowDev,omitempty"`
}

// +kcc:proto=google.cloud.baremetalsolution.v2.InstanceConfig
type InstanceConfigObservedState struct {
	// Output only. The name of the instance config.
	// +kcc:proto:field=google.cloud.baremetalsolution.v2.InstanceConfig.name
	Name *string `json:"name,omitempty"`
}

// +kcc:proto=google.cloud.baremetalsolution.v2.NetworkConfig
type NetworkConfigObservedState struct {
	// Output only. The name of the network config.
	// +kcc:proto:field=google.cloud.baremetalsolution.v2.NetworkConfig.name
	Name *string `json:"name,omitempty"`
}

// +kcc:proto=google.cloud.baremetalsolution.v2.ProvisioningConfig
type ProvisioningConfigObservedState struct {
	// Output only. The system-generated name of the provisioning config. This
	//  follows the UUID format.
	// +kcc:proto:field=google.cloud.baremetalsolution.v2.ProvisioningConfig.name
	Name *string `json:"name,omitempty"`

	// Instances to be created.
	// +kcc:proto:field=google.cloud.baremetalsolution.v2.ProvisioningConfig.instances
	Instances []InstanceConfigObservedState `json:"instances,omitempty"`

	// Networks to be created.
	// +kcc:proto:field=google.cloud.baremetalsolution.v2.ProvisioningConfig.networks
	Networks []NetworkConfigObservedState `json:"networks,omitempty"`

	// Volumes to be created.
	// +kcc:proto:field=google.cloud.baremetalsolution.v2.ProvisioningConfig.volumes
	Volumes []VolumeConfigObservedState `json:"volumes,omitempty"`

	// Output only. State of ProvisioningConfig.
	// +kcc:proto:field=google.cloud.baremetalsolution.v2.ProvisioningConfig.state
	State *string `json:"state,omitempty"`

	// Output only. Last update timestamp.
	// +kcc:proto:field=google.cloud.baremetalsolution.v2.ProvisioningConfig.update_time
	UpdateTime *string `json:"updateTime,omitempty"`

	// Output only. URI to Cloud Console UI view of this provisioning config.
	// +kcc:proto:field=google.cloud.baremetalsolution.v2.ProvisioningConfig.cloud_console_uri
	CloudConsoleURI *string `json:"cloudConsoleURI,omitempty"`
}

// +kcc:proto=google.cloud.baremetalsolution.v2.VolumeConfig
type VolumeConfigObservedState struct {
	// Output only. The name of the volume config.
	// +kcc:proto:field=google.cloud.baremetalsolution.v2.VolumeConfig.name
	Name *string `json:"name,omitempty"`
}
