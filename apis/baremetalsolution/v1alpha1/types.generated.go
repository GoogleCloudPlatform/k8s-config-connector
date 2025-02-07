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


// +kcc:proto=google.cloud.baremetalsolution.v2.Instance
type Instance struct {
	// Immutable. The resource name of this `Instance`.
	//  Resource names are schemeless URIs that follow the conventions in
	//  https://cloud.google.com/apis/design/resource_names.
	//  Format:
	//  `projects/{project}/locations/{location}/instances/{instance}`
	// +kcc:proto:field=google.cloud.baremetalsolution.v2.Instance.name
	Name *string `json:"name,omitempty"`

	// Immutable. The server type.
	//  [Available server
	//  types](https://cloud.google.com/bare-metal/docs/bms-planning#server_configurations)
	// +kcc:proto:field=google.cloud.baremetalsolution.v2.Instance.machine_type
	MachineType *string `json:"machineType,omitempty"`

	// True if you enable hyperthreading for the server, otherwise false.
	//  The default value is false.
	// +kcc:proto:field=google.cloud.baremetalsolution.v2.Instance.hyperthreading_enabled
	HyperthreadingEnabled *bool `json:"hyperthreadingEnabled,omitempty"`

	// Labels as key value pairs.
	// +kcc:proto:field=google.cloud.baremetalsolution.v2.Instance.labels
	Labels map[string]string `json:"labels,omitempty"`

	// Immutable. List of LUNs associated with this server.
	// +kcc:proto:field=google.cloud.baremetalsolution.v2.Instance.luns
	Luns []Lun `json:"luns,omitempty"`

	// Input only. List of Volumes to attach to this Instance on creation.
	//  This field won't be populated in Get/List responses.
	// +kcc:proto:field=google.cloud.baremetalsolution.v2.Instance.volumes
	Volumes []Volume `json:"volumes,omitempty"`

	// The OS image currently installed on the server.
	// +kcc:proto:field=google.cloud.baremetalsolution.v2.Instance.os_image
	OsImage *string `json:"osImage,omitempty"`

	// Immutable. Pod name.
	//  Pod is an independent part of infrastructure.
	//  Instance can be connected to the assets (networks, volumes) allocated
	//  in the same pod only.
	// +kcc:proto:field=google.cloud.baremetalsolution.v2.Instance.pod
	Pod *string `json:"pod,omitempty"`

	// Instance network template name. For eg, bondaa-bondaa, bondab-nic, etc.
	//  Generally, the template name follows the syntax of
	//  "bond<bond_mode>" or "nic".
	// +kcc:proto:field=google.cloud.baremetalsolution.v2.Instance.network_template
	NetworkTemplate *string `json:"networkTemplate,omitempty"`

	// List of logical interfaces for the instance. The number of logical
	//  interfaces will be the same as number of hardware bond/nic on the chosen
	//  network template. For the non-multivlan configurations (for eg, existing
	//  servers) that use existing default network template (bondaa-bondaa), both
	//  the Instance.networks field and the Instance.logical_interfaces fields will
	//  be filled to ensure backward compatibility. For the others, only
	//  Instance.logical_interfaces will be filled.
	// +kcc:proto:field=google.cloud.baremetalsolution.v2.Instance.logical_interfaces
	LogicalInterfaces []LogicalInterface `json:"logicalInterfaces,omitempty"`

	// The workload profile for the instance.
	// +kcc:proto:field=google.cloud.baremetalsolution.v2.Instance.workload_profile
	WorkloadProfile *string `json:"workloadProfile,omitempty"`
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

// +kcc:proto=google.cloud.baremetalsolution.v2.Lun
type Lun struct {

	// An identifier for the LUN, generated by the backend.
	// +kcc:proto:field=google.cloud.baremetalsolution.v2.Lun.id
	ID *string `json:"id,omitempty"`

	// The state of this storage volume.
	// +kcc:proto:field=google.cloud.baremetalsolution.v2.Lun.state
	State *string `json:"state,omitempty"`

	// The size of this LUN, in gigabytes.
	// +kcc:proto:field=google.cloud.baremetalsolution.v2.Lun.size_gb
	SizeGB *int64 `json:"sizeGB,omitempty"`

	// The LUN multiprotocol type ensures the characteristics of the LUN are
	//  optimized for each operating system.
	// +kcc:proto:field=google.cloud.baremetalsolution.v2.Lun.multiprotocol_type
	MultiprotocolType *string `json:"multiprotocolType,omitempty"`

	// Display the storage volume for this LUN.
	// +kcc:proto:field=google.cloud.baremetalsolution.v2.Lun.storage_volume
	StorageVolume *string `json:"storageVolume,omitempty"`

	// Display if this LUN can be shared between multiple physical servers.
	// +kcc:proto:field=google.cloud.baremetalsolution.v2.Lun.shareable
	Shareable *bool `json:"shareable,omitempty"`

	// Display if this LUN is a boot LUN.
	// +kcc:proto:field=google.cloud.baremetalsolution.v2.Lun.boot_lun
	BootLun *bool `json:"bootLun,omitempty"`

	// The storage type for this LUN.
	// +kcc:proto:field=google.cloud.baremetalsolution.v2.Lun.storage_type
	StorageType *string `json:"storageType,omitempty"`

	// The WWID for this LUN.
	// +kcc:proto:field=google.cloud.baremetalsolution.v2.Lun.wwid
	Wwid *string `json:"wwid,omitempty"`
}

// +kcc:proto=google.cloud.baremetalsolution.v2.Network
type Network struct {

	// An identifier for the `Network`, generated by the backend.
	// +kcc:proto:field=google.cloud.baremetalsolution.v2.Network.id
	ID *string `json:"id,omitempty"`

	// The type of this network.
	// +kcc:proto:field=google.cloud.baremetalsolution.v2.Network.type
	Type *string `json:"type,omitempty"`

	// IP address configured.
	// +kcc:proto:field=google.cloud.baremetalsolution.v2.Network.ip_address
	IPAddress *string `json:"ipAddress,omitempty"`

	// List of physical interfaces.
	// +kcc:proto:field=google.cloud.baremetalsolution.v2.Network.mac_address
	MacAddress []string `json:"macAddress,omitempty"`

	// The Network state.
	// +kcc:proto:field=google.cloud.baremetalsolution.v2.Network.state
	State *string `json:"state,omitempty"`

	// The vlan id of the Network.
	// +kcc:proto:field=google.cloud.baremetalsolution.v2.Network.vlan_id
	VlanID *string `json:"vlanID,omitempty"`

	// The cidr of the Network.
	// +kcc:proto:field=google.cloud.baremetalsolution.v2.Network.cidr
	Cidr *string `json:"cidr,omitempty"`

	// The vrf for the Network.
	// +kcc:proto:field=google.cloud.baremetalsolution.v2.Network.vrf
	Vrf *VRF `json:"vrf,omitempty"`

	// Labels as key value pairs.
	// +kcc:proto:field=google.cloud.baremetalsolution.v2.Network.labels
	Labels map[string]string `json:"labels,omitempty"`

	// IP range for reserved for services (e.g. NFS).
	// +kcc:proto:field=google.cloud.baremetalsolution.v2.Network.services_cidr
	ServicesCidr *string `json:"servicesCidr,omitempty"`

	// List of IP address reservations in this network.
	//  When updating this field, an error will be generated if a reservation
	//  conflicts with an IP address already allocated to a physical server.
	// +kcc:proto:field=google.cloud.baremetalsolution.v2.Network.reservations
	Reservations []NetworkAddressReservation `json:"reservations,omitempty"`

	// Input only. List of mount points to attach the network to.
	// +kcc:proto:field=google.cloud.baremetalsolution.v2.Network.mount_points
	MountPoints []NetworkMountPoint `json:"mountPoints,omitempty"`

	// Whether network uses standard frames or jumbo ones.
	// +kcc:proto:field=google.cloud.baremetalsolution.v2.Network.jumbo_frames_enabled
	JumboFramesEnabled *bool `json:"jumboFramesEnabled,omitempty"`
}

// +kcc:proto=google.cloud.baremetalsolution.v2.NetworkAddressReservation
type NetworkAddressReservation struct {
	// The first address of this reservation block.
	//  Must be specified as a single IPv4 address, e.g. 10.1.2.2.
	// +kcc:proto:field=google.cloud.baremetalsolution.v2.NetworkAddressReservation.start_address
	StartAddress *string `json:"startAddress,omitempty"`

	// The last address of this reservation block, inclusive. I.e., for cases when
	//  reservations are only single addresses, end_address and start_address will
	//  be the same.
	//  Must be specified as a single IPv4 address, e.g. 10.1.2.2.
	// +kcc:proto:field=google.cloud.baremetalsolution.v2.NetworkAddressReservation.end_address
	EndAddress *string `json:"endAddress,omitempty"`

	// A note about this reservation, intended for human consumption.
	// +kcc:proto:field=google.cloud.baremetalsolution.v2.NetworkAddressReservation.note
	Note *string `json:"note,omitempty"`
}

// +kcc:proto=google.cloud.baremetalsolution.v2.NetworkMountPoint
type NetworkMountPoint struct {
	// Instance to attach network to.
	// +kcc:proto:field=google.cloud.baremetalsolution.v2.NetworkMountPoint.instance
	Instance *string `json:"instance,omitempty"`

	// Logical interface to detach from.
	// +kcc:proto:field=google.cloud.baremetalsolution.v2.NetworkMountPoint.logical_interface
	LogicalInterface *string `json:"logicalInterface,omitempty"`

	// Network should be a default gateway.
	// +kcc:proto:field=google.cloud.baremetalsolution.v2.NetworkMountPoint.default_gateway
	DefaultGateway *bool `json:"defaultGateway,omitempty"`

	// Ip address of the server.
	// +kcc:proto:field=google.cloud.baremetalsolution.v2.NetworkMountPoint.ip_address
	IPAddress *string `json:"ipAddress,omitempty"`
}

// +kcc:proto=google.cloud.baremetalsolution.v2.VRF
type VRF struct {
	// The name of the VRF.
	// +kcc:proto:field=google.cloud.baremetalsolution.v2.VRF.name
	Name *string `json:"name,omitempty"`

	// The possible state of VRF.
	// +kcc:proto:field=google.cloud.baremetalsolution.v2.VRF.state
	State *string `json:"state,omitempty"`

	// The QOS policy applied to this VRF.
	//  The value is only meaningful when all the vlan attachments have the same
	//  QoS. This field should not be used for new integrations, use vlan
	//  attachment level qos instead. The field is left for backward-compatibility.
	// +kcc:proto:field=google.cloud.baremetalsolution.v2.VRF.qos_policy
	QosPolicy *VRF_QosPolicy `json:"qosPolicy,omitempty"`

	// The list of VLAN attachments for the VRF.
	// +kcc:proto:field=google.cloud.baremetalsolution.v2.VRF.vlan_attachments
	VlanAttachments []VRF_VlanAttachment `json:"vlanAttachments,omitempty"`
}

// +kcc:proto=google.cloud.baremetalsolution.v2.VRF.QosPolicy
type VRF_QosPolicy struct {
	// The bandwidth permitted by the QOS policy, in gbps.
	// +kcc:proto:field=google.cloud.baremetalsolution.v2.VRF.QosPolicy.bandwidth_gbps
	BandwidthGbps *float64 `json:"bandwidthGbps,omitempty"`
}

// +kcc:proto=google.cloud.baremetalsolution.v2.VRF.VlanAttachment
type VRF_VlanAttachment struct {
	// The peer vlan ID of the attachment.
	// +kcc:proto:field=google.cloud.baremetalsolution.v2.VRF.VlanAttachment.peer_vlan_id
	PeerVlanID *int64 `json:"peerVlanID,omitempty"`

	// The peer IP of the attachment.
	// +kcc:proto:field=google.cloud.baremetalsolution.v2.VRF.VlanAttachment.peer_ip
	PeerIP *string `json:"peerIP,omitempty"`

	// The router IP of the attachment.
	// +kcc:proto:field=google.cloud.baremetalsolution.v2.VRF.VlanAttachment.router_ip
	RouterIP *string `json:"routerIP,omitempty"`

	// Input only. Pairing key.
	// +kcc:proto:field=google.cloud.baremetalsolution.v2.VRF.VlanAttachment.pairing_key
	PairingKey *string `json:"pairingKey,omitempty"`

	// The QOS policy applied to this VLAN attachment.
	//  This value should be preferred to using qos at vrf level.
	// +kcc:proto:field=google.cloud.baremetalsolution.v2.VRF.VlanAttachment.qos_policy
	QosPolicy *VRF_QosPolicy `json:"qosPolicy,omitempty"`

	// Immutable. The identifier of the attachment within vrf.
	// +kcc:proto:field=google.cloud.baremetalsolution.v2.VRF.VlanAttachment.id
	ID *string `json:"id,omitempty"`

	// Optional. The name of the vlan attachment within vrf. This is of the form
	//  projects/{project_number}/regions/{region}/interconnectAttachments/{interconnect_attachment}
	// +kcc:proto:field=google.cloud.baremetalsolution.v2.VRF.VlanAttachment.interconnect_attachment
	InterconnectAttachment *string `json:"interconnectAttachment,omitempty"`
}

// +kcc:proto=google.cloud.baremetalsolution.v2.Volume
type Volume struct {

	// An identifier for the `Volume`, generated by the backend.
	// +kcc:proto:field=google.cloud.baremetalsolution.v2.Volume.id
	ID *string `json:"id,omitempty"`

	// The storage type for this volume.
	// +kcc:proto:field=google.cloud.baremetalsolution.v2.Volume.storage_type
	StorageType *string `json:"storageType,omitempty"`

	// The state of this storage volume.
	// +kcc:proto:field=google.cloud.baremetalsolution.v2.Volume.state
	State *string `json:"state,omitempty"`

	// The requested size of this storage volume, in GiB.
	// +kcc:proto:field=google.cloud.baremetalsolution.v2.Volume.requested_size_gib
	RequestedSizeGib *int64 `json:"requestedSizeGib,omitempty"`

	// Originally requested size, in GiB.
	// +kcc:proto:field=google.cloud.baremetalsolution.v2.Volume.originally_requested_size_gib
	OriginallyRequestedSizeGib *int64 `json:"originallyRequestedSizeGib,omitempty"`

	// The current size of this storage volume, in GiB, including space reserved
	//  for snapshots. This size might be different than the requested size if the
	//  storage volume has been configured with auto grow or auto shrink.
	// +kcc:proto:field=google.cloud.baremetalsolution.v2.Volume.current_size_gib
	CurrentSizeGib *int64 `json:"currentSizeGib,omitempty"`

	// Additional emergency size that was requested for this Volume, in GiB.
	//  current_size_gib includes this value.
	// +kcc:proto:field=google.cloud.baremetalsolution.v2.Volume.emergency_size_gib
	EmergencySizeGib *int64 `json:"emergencySizeGib,omitempty"`

	// Maximum size volume can be expanded to in case of evergency, in GiB.
	// +kcc:proto:field=google.cloud.baremetalsolution.v2.Volume.max_size_gib
	MaxSizeGib *int64 `json:"maxSizeGib,omitempty"`

	// The size, in GiB, that this storage volume has expanded as a result of an
	//  auto grow policy. In the absence of auto-grow, the value is 0.
	// +kcc:proto:field=google.cloud.baremetalsolution.v2.Volume.auto_grown_size_gib
	AutoGrownSizeGib *int64 `json:"autoGrownSizeGib,omitempty"`

	// The space remaining in the storage volume for new LUNs, in GiB, excluding
	//  space reserved for snapshots.
	// +kcc:proto:field=google.cloud.baremetalsolution.v2.Volume.remaining_space_gib
	RemainingSpaceGib *int64 `json:"remainingSpaceGib,omitempty"`

	// Details about snapshot space reservation and usage on the storage volume.
	// +kcc:proto:field=google.cloud.baremetalsolution.v2.Volume.snapshot_reservation_detail
	SnapshotReservationDetail *Volume_SnapshotReservationDetail `json:"snapshotReservationDetail,omitempty"`

	// The behavior to use when snapshot reserved space is full.
	// +kcc:proto:field=google.cloud.baremetalsolution.v2.Volume.snapshot_auto_delete_behavior
	SnapshotAutoDeleteBehavior *string `json:"snapshotAutoDeleteBehavior,omitempty"`

	// Labels as key value pairs.
	// +kcc:proto:field=google.cloud.baremetalsolution.v2.Volume.labels
	Labels map[string]string `json:"labels,omitempty"`

	// Whether snapshots are enabled.
	// +kcc:proto:field=google.cloud.baremetalsolution.v2.Volume.snapshot_enabled
	SnapshotEnabled *bool `json:"snapshotEnabled,omitempty"`

	// Immutable. Pod name.
	// +kcc:proto:field=google.cloud.baremetalsolution.v2.Volume.pod
	Pod *string `json:"pod,omitempty"`

	// Immutable. Performance tier of the Volume.
	//  Default is SHARED.
	// +kcc:proto:field=google.cloud.baremetalsolution.v2.Volume.performance_tier
	PerformanceTier *string `json:"performanceTier,omitempty"`

	// Input only. User-specified notes for new Volume.
	//  Used to provision Volumes that require manual intervention.
	// +kcc:proto:field=google.cloud.baremetalsolution.v2.Volume.notes
	Notes *string `json:"notes,omitempty"`

	// The workload profile for the volume.
	// +kcc:proto:field=google.cloud.baremetalsolution.v2.Volume.workload_profile
	WorkloadProfile *string `json:"workloadProfile,omitempty"`
}

// +kcc:proto=google.cloud.baremetalsolution.v2.Volume.SnapshotReservationDetail
type Volume_SnapshotReservationDetail struct {
	// The space on this storage volume reserved for snapshots, shown in GiB.
	// +kcc:proto:field=google.cloud.baremetalsolution.v2.Volume.SnapshotReservationDetail.reserved_space_gib
	ReservedSpaceGib *int64 `json:"reservedSpaceGib,omitempty"`

	// The percent of snapshot space on this storage volume actually being used
	//  by the snapshot copies. This value might be higher than 100% if the
	//  snapshot copies have overflowed into the data portion of the storage
	//  volume.
	// +kcc:proto:field=google.cloud.baremetalsolution.v2.Volume.SnapshotReservationDetail.reserved_space_used_percent
	ReservedSpaceUsedPercent *int32 `json:"reservedSpaceUsedPercent,omitempty"`

	// The amount, in GiB, of available space in this storage volume's reserved
	//  snapshot space.
	// +kcc:proto:field=google.cloud.baremetalsolution.v2.Volume.SnapshotReservationDetail.reserved_space_remaining_gib
	ReservedSpaceRemainingGib *int64 `json:"reservedSpaceRemainingGib,omitempty"`

	// Percent of the total Volume size reserved for snapshot copies.
	//  Enabling snapshots requires reserving 20% or more of
	//  the storage volume space for snapshots. Maximum reserved space for
	//  snapshots is 40%.
	//  Setting this field will effectively set snapshot_enabled to true.
	// +kcc:proto:field=google.cloud.baremetalsolution.v2.Volume.SnapshotReservationDetail.reserved_space_percent
	ReservedSpacePercent *int32 `json:"reservedSpacePercent,omitempty"`
}

// +kcc:proto=google.cloud.baremetalsolution.v2.Instance
type InstanceObservedState struct {
	// Output only. An identifier for the `Instance`, generated by the backend.
	// +kcc:proto:field=google.cloud.baremetalsolution.v2.Instance.id
	ID *string `json:"id,omitempty"`

	// Output only. Create a time stamp.
	// +kcc:proto:field=google.cloud.baremetalsolution.v2.Instance.create_time
	CreateTime *string `json:"createTime,omitempty"`

	// Output only. Update a time stamp.
	// +kcc:proto:field=google.cloud.baremetalsolution.v2.Instance.update_time
	UpdateTime *string `json:"updateTime,omitempty"`

	// Output only. The state of the server.
	// +kcc:proto:field=google.cloud.baremetalsolution.v2.Instance.state
	State *string `json:"state,omitempty"`

	// Immutable. List of LUNs associated with this server.
	// +kcc:proto:field=google.cloud.baremetalsolution.v2.Instance.luns
	Luns []LunObservedState `json:"luns,omitempty"`

	// Input only. List of Volumes to attach to this Instance on creation.
	//  This field won't be populated in Get/List responses.
	// +kcc:proto:field=google.cloud.baremetalsolution.v2.Instance.volumes
	Volumes []VolumeObservedState `json:"volumes,omitempty"`

	// Output only. List of networks associated with this server.
	// +kcc:proto:field=google.cloud.baremetalsolution.v2.Instance.networks
	Networks []Network `json:"networks,omitempty"`

	// Output only. True if the interactive serial console feature is enabled for
	//  the instance, false otherwise. The default value is false.
	// +kcc:proto:field=google.cloud.baremetalsolution.v2.Instance.interactive_serial_console_enabled
	InteractiveSerialConsoleEnabled *bool `json:"interactiveSerialConsoleEnabled,omitempty"`

	// Output only. Text field about info for logging in.
	// +kcc:proto:field=google.cloud.baremetalsolution.v2.Instance.login_info
	LoginInfo *string `json:"loginInfo,omitempty"`

	// Output only. The firmware version for the instance.
	// +kcc:proto:field=google.cloud.baremetalsolution.v2.Instance.firmware_version
	FirmwareVersion *string `json:"firmwareVersion,omitempty"`
}

// +kcc:proto=google.cloud.baremetalsolution.v2.Lun
type LunObservedState struct {
	// Output only. The name of the LUN.
	// +kcc:proto:field=google.cloud.baremetalsolution.v2.Lun.name
	Name *string `json:"name,omitempty"`

	// Output only. Time after which LUN will be fully deleted.
	//  It is filled only for LUNs in COOL_OFF state.
	// +kcc:proto:field=google.cloud.baremetalsolution.v2.Lun.expire_time
	ExpireTime *string `json:"expireTime,omitempty"`

	// Output only. Instances this Lun is attached to.
	// +kcc:proto:field=google.cloud.baremetalsolution.v2.Lun.instances
	Instances []string `json:"instances,omitempty"`
}

// +kcc:proto=google.cloud.baremetalsolution.v2.Network
type NetworkObservedState struct {
	// Output only. The resource name of this `Network`.
	//  Resource names are schemeless URIs that follow the conventions in
	//  https://cloud.google.com/apis/design/resource_names.
	//  Format:
	//  `projects/{project}/locations/{location}/networks/{network}`
	// +kcc:proto:field=google.cloud.baremetalsolution.v2.Network.name
	Name *string `json:"name,omitempty"`

	// Output only. Pod name.
	// +kcc:proto:field=google.cloud.baremetalsolution.v2.Network.pod
	Pod *string `json:"pod,omitempty"`

	// Output only. Gateway ip address.
	// +kcc:proto:field=google.cloud.baremetalsolution.v2.Network.gateway_ip
	GatewayIP *string `json:"gatewayIP,omitempty"`
}

// +kcc:proto=google.cloud.baremetalsolution.v2.Volume
type VolumeObservedState struct {
	// Output only. The resource name of this `Volume`.
	//  Resource names are schemeless URIs that follow the conventions in
	//  https://cloud.google.com/apis/design/resource_names.
	//  Format:
	//  `projects/{project}/locations/{location}/volumes/{volume}`
	// +kcc:proto:field=google.cloud.baremetalsolution.v2.Volume.name
	Name *string `json:"name,omitempty"`

	// Output only. Storage protocol for the Volume.
	// +kcc:proto:field=google.cloud.baremetalsolution.v2.Volume.protocol
	Protocol *string `json:"protocol,omitempty"`

	// Output only. Whether this volume is a boot volume. A boot volume is one
	//  which contains a boot LUN.
	// +kcc:proto:field=google.cloud.baremetalsolution.v2.Volume.boot_volume
	BootVolume *bool `json:"bootVolume,omitempty"`

	// Output only. Time after which volume will be fully deleted.
	//  It is filled only for volumes in COOLOFF state.
	// +kcc:proto:field=google.cloud.baremetalsolution.v2.Volume.expire_time
	ExpireTime *string `json:"expireTime,omitempty"`

	// Output only. Instances this Volume is attached to.
	//  This field is set only in Get requests.
	// +kcc:proto:field=google.cloud.baremetalsolution.v2.Volume.instances
	Instances []string `json:"instances,omitempty"`

	// Output only. Is the Volume attached at at least one instance.
	//  This field is a lightweight counterpart of `instances` field.
	//  It is filled in List responses as well.
	// +kcc:proto:field=google.cloud.baremetalsolution.v2.Volume.attached
	Attached *bool `json:"attached,omitempty"`
}
