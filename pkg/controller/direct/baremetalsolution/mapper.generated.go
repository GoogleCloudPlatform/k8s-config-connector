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

package baremetalsolution

import (
	refs "github.com/GoogleCloudPlatform/k8s-config-connector/apis/refs/v1beta1"
	pb "cloud.google.com/go/baremetalsolution/apiv2/baremetalsolutionpb"
	krm "github.com/GoogleCloudPlatform/k8s-config-connector/apis/baremetalsolution/v1alpha1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct"
)
func BaremetalsolutionInstanceObservedState_FromProto(mapCtx *direct.MapContext, in *pb.Instance) *krm.BaremetalsolutionInstanceObservedState {
	if in == nil {
		return nil
	}
	out := &krm.BaremetalsolutionInstanceObservedState{}
	// MISSING: Name
	// MISSING: ID
	// MISSING: CreateTime
	// MISSING: UpdateTime
	// MISSING: MachineType
	// MISSING: State
	// MISSING: HyperthreadingEnabled
	// MISSING: Labels
	// MISSING: Luns
	// MISSING: Volumes
	// MISSING: Networks
	// MISSING: InteractiveSerialConsoleEnabled
	// MISSING: OsImage
	// MISSING: Pod
	// MISSING: NetworkTemplate
	// MISSING: LogicalInterfaces
	// MISSING: LoginInfo
	// MISSING: WorkloadProfile
	// MISSING: FirmwareVersion
	return out
}
func BaremetalsolutionInstanceObservedState_ToProto(mapCtx *direct.MapContext, in *krm.BaremetalsolutionInstanceObservedState) *pb.Instance {
	if in == nil {
		return nil
	}
	out := &pb.Instance{}
	// MISSING: Name
	// MISSING: ID
	// MISSING: CreateTime
	// MISSING: UpdateTime
	// MISSING: MachineType
	// MISSING: State
	// MISSING: HyperthreadingEnabled
	// MISSING: Labels
	// MISSING: Luns
	// MISSING: Volumes
	// MISSING: Networks
	// MISSING: InteractiveSerialConsoleEnabled
	// MISSING: OsImage
	// MISSING: Pod
	// MISSING: NetworkTemplate
	// MISSING: LogicalInterfaces
	// MISSING: LoginInfo
	// MISSING: WorkloadProfile
	// MISSING: FirmwareVersion
	return out
}
func BaremetalsolutionInstanceSpec_FromProto(mapCtx *direct.MapContext, in *pb.Instance) *krm.BaremetalsolutionInstanceSpec {
	if in == nil {
		return nil
	}
	out := &krm.BaremetalsolutionInstanceSpec{}
	// MISSING: Name
	// MISSING: ID
	// MISSING: CreateTime
	// MISSING: UpdateTime
	// MISSING: MachineType
	// MISSING: State
	// MISSING: HyperthreadingEnabled
	// MISSING: Labels
	// MISSING: Luns
	// MISSING: Volumes
	// MISSING: Networks
	// MISSING: InteractiveSerialConsoleEnabled
	// MISSING: OsImage
	// MISSING: Pod
	// MISSING: NetworkTemplate
	// MISSING: LogicalInterfaces
	// MISSING: LoginInfo
	// MISSING: WorkloadProfile
	// MISSING: FirmwareVersion
	return out
}
func BaremetalsolutionInstanceSpec_ToProto(mapCtx *direct.MapContext, in *krm.BaremetalsolutionInstanceSpec) *pb.Instance {
	if in == nil {
		return nil
	}
	out := &pb.Instance{}
	// MISSING: Name
	// MISSING: ID
	// MISSING: CreateTime
	// MISSING: UpdateTime
	// MISSING: MachineType
	// MISSING: State
	// MISSING: HyperthreadingEnabled
	// MISSING: Labels
	// MISSING: Luns
	// MISSING: Volumes
	// MISSING: Networks
	// MISSING: InteractiveSerialConsoleEnabled
	// MISSING: OsImage
	// MISSING: Pod
	// MISSING: NetworkTemplate
	// MISSING: LogicalInterfaces
	// MISSING: LoginInfo
	// MISSING: WorkloadProfile
	// MISSING: FirmwareVersion
	return out
}
func BaremetalsolutionLunObservedState_FromProto(mapCtx *direct.MapContext, in *pb.Lun) *krm.BaremetalsolutionLunObservedState {
	if in == nil {
		return nil
	}
	out := &krm.BaremetalsolutionLunObservedState{}
	// MISSING: Name
	// MISSING: ID
	// MISSING: State
	// MISSING: SizeGB
	// MISSING: MultiprotocolType
	// MISSING: StorageVolume
	// MISSING: Shareable
	// MISSING: BootLun
	// MISSING: StorageType
	// MISSING: Wwid
	// MISSING: ExpireTime
	// MISSING: Instances
	return out
}
func BaremetalsolutionLunObservedState_ToProto(mapCtx *direct.MapContext, in *krm.BaremetalsolutionLunObservedState) *pb.Lun {
	if in == nil {
		return nil
	}
	out := &pb.Lun{}
	// MISSING: Name
	// MISSING: ID
	// MISSING: State
	// MISSING: SizeGB
	// MISSING: MultiprotocolType
	// MISSING: StorageVolume
	// MISSING: Shareable
	// MISSING: BootLun
	// MISSING: StorageType
	// MISSING: Wwid
	// MISSING: ExpireTime
	// MISSING: Instances
	return out
}
func BaremetalsolutionLunSpec_FromProto(mapCtx *direct.MapContext, in *pb.Lun) *krm.BaremetalsolutionLunSpec {
	if in == nil {
		return nil
	}
	out := &krm.BaremetalsolutionLunSpec{}
	// MISSING: Name
	// MISSING: ID
	// MISSING: State
	// MISSING: SizeGB
	// MISSING: MultiprotocolType
	// MISSING: StorageVolume
	// MISSING: Shareable
	// MISSING: BootLun
	// MISSING: StorageType
	// MISSING: Wwid
	// MISSING: ExpireTime
	// MISSING: Instances
	return out
}
func BaremetalsolutionLunSpec_ToProto(mapCtx *direct.MapContext, in *krm.BaremetalsolutionLunSpec) *pb.Lun {
	if in == nil {
		return nil
	}
	out := &pb.Lun{}
	// MISSING: Name
	// MISSING: ID
	// MISSING: State
	// MISSING: SizeGB
	// MISSING: MultiprotocolType
	// MISSING: StorageVolume
	// MISSING: Shareable
	// MISSING: BootLun
	// MISSING: StorageType
	// MISSING: Wwid
	// MISSING: ExpireTime
	// MISSING: Instances
	return out
}
func BaremetalsolutionNetworkObservedState_FromProto(mapCtx *direct.MapContext, in *pb.Network) *krm.BaremetalsolutionNetworkObservedState {
	if in == nil {
		return nil
	}
	out := &krm.BaremetalsolutionNetworkObservedState{}
	// MISSING: Name
	// MISSING: ID
	// MISSING: Type
	// MISSING: IPAddress
	// MISSING: MacAddress
	// MISSING: State
	// MISSING: VlanID
	// MISSING: Cidr
	// MISSING: Vrf
	// MISSING: Labels
	// MISSING: ServicesCidr
	// MISSING: Reservations
	// MISSING: Pod
	// MISSING: MountPoints
	// MISSING: JumboFramesEnabled
	// MISSING: GatewayIP
	return out
}
func BaremetalsolutionNetworkObservedState_ToProto(mapCtx *direct.MapContext, in *krm.BaremetalsolutionNetworkObservedState) *pb.Network {
	if in == nil {
		return nil
	}
	out := &pb.Network{}
	// MISSING: Name
	// MISSING: ID
	// MISSING: Type
	// MISSING: IPAddress
	// MISSING: MacAddress
	// MISSING: State
	// MISSING: VlanID
	// MISSING: Cidr
	// MISSING: Vrf
	// MISSING: Labels
	// MISSING: ServicesCidr
	// MISSING: Reservations
	// MISSING: Pod
	// MISSING: MountPoints
	// MISSING: JumboFramesEnabled
	// MISSING: GatewayIP
	return out
}
func BaremetalsolutionNetworkSpec_FromProto(mapCtx *direct.MapContext, in *pb.Network) *krm.BaremetalsolutionNetworkSpec {
	if in == nil {
		return nil
	}
	out := &krm.BaremetalsolutionNetworkSpec{}
	// MISSING: Name
	// MISSING: ID
	// MISSING: Type
	// MISSING: IPAddress
	// MISSING: MacAddress
	// MISSING: State
	// MISSING: VlanID
	// MISSING: Cidr
	// MISSING: Vrf
	// MISSING: Labels
	// MISSING: ServicesCidr
	// MISSING: Reservations
	// MISSING: Pod
	// MISSING: MountPoints
	// MISSING: JumboFramesEnabled
	// MISSING: GatewayIP
	return out
}
func BaremetalsolutionNetworkSpec_ToProto(mapCtx *direct.MapContext, in *krm.BaremetalsolutionNetworkSpec) *pb.Network {
	if in == nil {
		return nil
	}
	out := &pb.Network{}
	// MISSING: Name
	// MISSING: ID
	// MISSING: Type
	// MISSING: IPAddress
	// MISSING: MacAddress
	// MISSING: State
	// MISSING: VlanID
	// MISSING: Cidr
	// MISSING: Vrf
	// MISSING: Labels
	// MISSING: ServicesCidr
	// MISSING: Reservations
	// MISSING: Pod
	// MISSING: MountPoints
	// MISSING: JumboFramesEnabled
	// MISSING: GatewayIP
	return out
}
func BaremetalsolutionNfsShareObservedState_FromProto(mapCtx *direct.MapContext, in *pb.NfsShare) *krm.BaremetalsolutionNfsShareObservedState {
	if in == nil {
		return nil
	}
	out := &krm.BaremetalsolutionNfsShareObservedState{}
	// MISSING: Name
	// MISSING: NfsShareID
	// MISSING: ID
	// MISSING: State
	// MISSING: Volume
	// MISSING: AllowedClients
	// MISSING: Labels
	// MISSING: RequestedSizeGib
	// MISSING: StorageType
	return out
}
func BaremetalsolutionNfsShareObservedState_ToProto(mapCtx *direct.MapContext, in *krm.BaremetalsolutionNfsShareObservedState) *pb.NfsShare {
	if in == nil {
		return nil
	}
	out := &pb.NfsShare{}
	// MISSING: Name
	// MISSING: NfsShareID
	// MISSING: ID
	// MISSING: State
	// MISSING: Volume
	// MISSING: AllowedClients
	// MISSING: Labels
	// MISSING: RequestedSizeGib
	// MISSING: StorageType
	return out
}
func BaremetalsolutionNfsShareSpec_FromProto(mapCtx *direct.MapContext, in *pb.NfsShare) *krm.BaremetalsolutionNfsShareSpec {
	if in == nil {
		return nil
	}
	out := &krm.BaremetalsolutionNfsShareSpec{}
	// MISSING: Name
	// MISSING: NfsShareID
	// MISSING: ID
	// MISSING: State
	// MISSING: Volume
	// MISSING: AllowedClients
	// MISSING: Labels
	// MISSING: RequestedSizeGib
	// MISSING: StorageType
	return out
}
func BaremetalsolutionNfsShareSpec_ToProto(mapCtx *direct.MapContext, in *krm.BaremetalsolutionNfsShareSpec) *pb.NfsShare {
	if in == nil {
		return nil
	}
	out := &pb.NfsShare{}
	// MISSING: Name
	// MISSING: NfsShareID
	// MISSING: ID
	// MISSING: State
	// MISSING: Volume
	// MISSING: AllowedClients
	// MISSING: Labels
	// MISSING: RequestedSizeGib
	// MISSING: StorageType
	return out
}
func BaremetalsolutionOSImageObservedState_FromProto(mapCtx *direct.MapContext, in *pb.OSImage) *krm.BaremetalsolutionOSImageObservedState {
	if in == nil {
		return nil
	}
	out := &krm.BaremetalsolutionOSImageObservedState{}
	// MISSING: Name
	// MISSING: Code
	// MISSING: Description
	// MISSING: ApplicableInstanceTypes
	// MISSING: SupportedNetworkTemplates
	return out
}
func BaremetalsolutionOSImageObservedState_ToProto(mapCtx *direct.MapContext, in *krm.BaremetalsolutionOSImageObservedState) *pb.OSImage {
	if in == nil {
		return nil
	}
	out := &pb.OSImage{}
	// MISSING: Name
	// MISSING: Code
	// MISSING: Description
	// MISSING: ApplicableInstanceTypes
	// MISSING: SupportedNetworkTemplates
	return out
}
func BaremetalsolutionOSImageSpec_FromProto(mapCtx *direct.MapContext, in *pb.OSImage) *krm.BaremetalsolutionOSImageSpec {
	if in == nil {
		return nil
	}
	out := &krm.BaremetalsolutionOSImageSpec{}
	// MISSING: Name
	// MISSING: Code
	// MISSING: Description
	// MISSING: ApplicableInstanceTypes
	// MISSING: SupportedNetworkTemplates
	return out
}
func BaremetalsolutionOSImageSpec_ToProto(mapCtx *direct.MapContext, in *krm.BaremetalsolutionOSImageSpec) *pb.OSImage {
	if in == nil {
		return nil
	}
	out := &pb.OSImage{}
	// MISSING: Name
	// MISSING: Code
	// MISSING: Description
	// MISSING: ApplicableInstanceTypes
	// MISSING: SupportedNetworkTemplates
	return out
}
func BaremetalsolutionProvisioningConfigObservedState_FromProto(mapCtx *direct.MapContext, in *pb.ProvisioningConfig) *krm.BaremetalsolutionProvisioningConfigObservedState {
	if in == nil {
		return nil
	}
	out := &krm.BaremetalsolutionProvisioningConfigObservedState{}
	// MISSING: Name
	// MISSING: Instances
	// MISSING: Networks
	// MISSING: Volumes
	// MISSING: TicketID
	// MISSING: HandoverServiceAccount
	// MISSING: Email
	// MISSING: State
	// MISSING: Location
	// MISSING: UpdateTime
	// MISSING: CloudConsoleURI
	// MISSING: VpcScEnabled
	// MISSING: StatusMessage
	// MISSING: CustomID
	return out
}
func BaremetalsolutionProvisioningConfigObservedState_ToProto(mapCtx *direct.MapContext, in *krm.BaremetalsolutionProvisioningConfigObservedState) *pb.ProvisioningConfig {
	if in == nil {
		return nil
	}
	out := &pb.ProvisioningConfig{}
	// MISSING: Name
	// MISSING: Instances
	// MISSING: Networks
	// MISSING: Volumes
	// MISSING: TicketID
	// MISSING: HandoverServiceAccount
	// MISSING: Email
	// MISSING: State
	// MISSING: Location
	// MISSING: UpdateTime
	// MISSING: CloudConsoleURI
	// MISSING: VpcScEnabled
	// MISSING: StatusMessage
	// MISSING: CustomID
	return out
}
func BaremetalsolutionProvisioningConfigSpec_FromProto(mapCtx *direct.MapContext, in *pb.ProvisioningConfig) *krm.BaremetalsolutionProvisioningConfigSpec {
	if in == nil {
		return nil
	}
	out := &krm.BaremetalsolutionProvisioningConfigSpec{}
	// MISSING: Name
	// MISSING: Instances
	// MISSING: Networks
	// MISSING: Volumes
	// MISSING: TicketID
	// MISSING: HandoverServiceAccount
	// MISSING: Email
	// MISSING: State
	// MISSING: Location
	// MISSING: UpdateTime
	// MISSING: CloudConsoleURI
	// MISSING: VpcScEnabled
	// MISSING: StatusMessage
	// MISSING: CustomID
	return out
}
func BaremetalsolutionProvisioningConfigSpec_ToProto(mapCtx *direct.MapContext, in *krm.BaremetalsolutionProvisioningConfigSpec) *pb.ProvisioningConfig {
	if in == nil {
		return nil
	}
	out := &pb.ProvisioningConfig{}
	// MISSING: Name
	// MISSING: Instances
	// MISSING: Networks
	// MISSING: Volumes
	// MISSING: TicketID
	// MISSING: HandoverServiceAccount
	// MISSING: Email
	// MISSING: State
	// MISSING: Location
	// MISSING: UpdateTime
	// MISSING: CloudConsoleURI
	// MISSING: VpcScEnabled
	// MISSING: StatusMessage
	// MISSING: CustomID
	return out
}
func BaremetalsolutionServerNetworkTemplateObservedState_FromProto(mapCtx *direct.MapContext, in *pb.ServerNetworkTemplate) *krm.BaremetalsolutionServerNetworkTemplateObservedState {
	if in == nil {
		return nil
	}
	out := &krm.BaremetalsolutionServerNetworkTemplateObservedState{}
	// MISSING: Name
	// MISSING: ApplicableInstanceTypes
	// MISSING: LogicalInterfaces
	return out
}
func BaremetalsolutionServerNetworkTemplateObservedState_ToProto(mapCtx *direct.MapContext, in *krm.BaremetalsolutionServerNetworkTemplateObservedState) *pb.ServerNetworkTemplate {
	if in == nil {
		return nil
	}
	out := &pb.ServerNetworkTemplate{}
	// MISSING: Name
	// MISSING: ApplicableInstanceTypes
	// MISSING: LogicalInterfaces
	return out
}
func BaremetalsolutionServerNetworkTemplateSpec_FromProto(mapCtx *direct.MapContext, in *pb.ServerNetworkTemplate) *krm.BaremetalsolutionServerNetworkTemplateSpec {
	if in == nil {
		return nil
	}
	out := &krm.BaremetalsolutionServerNetworkTemplateSpec{}
	// MISSING: Name
	// MISSING: ApplicableInstanceTypes
	// MISSING: LogicalInterfaces
	return out
}
func BaremetalsolutionServerNetworkTemplateSpec_ToProto(mapCtx *direct.MapContext, in *krm.BaremetalsolutionServerNetworkTemplateSpec) *pb.ServerNetworkTemplate {
	if in == nil {
		return nil
	}
	out := &pb.ServerNetworkTemplate{}
	// MISSING: Name
	// MISSING: ApplicableInstanceTypes
	// MISSING: LogicalInterfaces
	return out
}
func InstanceConfig_FromProto(mapCtx *direct.MapContext, in *pb.InstanceConfig) *krm.InstanceConfig {
	if in == nil {
		return nil
	}
	out := &krm.InstanceConfig{}
	// MISSING: Name
	out.ID = direct.LazyPtr(in.GetId())
	out.InstanceType = direct.LazyPtr(in.GetInstanceType())
	out.Hyperthreading = direct.LazyPtr(in.GetHyperthreading())
	out.OsImage = direct.LazyPtr(in.GetOsImage())
	out.ClientNetwork = InstanceConfig_NetworkAddress_FromProto(mapCtx, in.GetClientNetwork())
	out.PrivateNetwork = InstanceConfig_NetworkAddress_FromProto(mapCtx, in.GetPrivateNetwork())
	out.UserNote = direct.LazyPtr(in.GetUserNote())
	out.AccountNetworksEnabled = direct.LazyPtr(in.GetAccountNetworksEnabled())
	out.NetworkConfig = direct.Enum_FromProto(mapCtx, in.GetNetworkConfig())
	out.NetworkTemplate = direct.LazyPtr(in.GetNetworkTemplate())
	out.LogicalInterfaces = direct.Slice_FromProto(mapCtx, in.LogicalInterfaces, LogicalInterface_FromProto)
	out.SSHKeyNames = in.SshKeyNames
	return out
}
func InstanceConfig_ToProto(mapCtx *direct.MapContext, in *krm.InstanceConfig) *pb.InstanceConfig {
	if in == nil {
		return nil
	}
	out := &pb.InstanceConfig{}
	// MISSING: Name
	out.Id = direct.ValueOf(in.ID)
	out.InstanceType = direct.ValueOf(in.InstanceType)
	out.Hyperthreading = direct.ValueOf(in.Hyperthreading)
	out.OsImage = direct.ValueOf(in.OsImage)
	out.ClientNetwork = InstanceConfig_NetworkAddress_ToProto(mapCtx, in.ClientNetwork)
	out.PrivateNetwork = InstanceConfig_NetworkAddress_ToProto(mapCtx, in.PrivateNetwork)
	out.UserNote = direct.ValueOf(in.UserNote)
	out.AccountNetworksEnabled = direct.ValueOf(in.AccountNetworksEnabled)
	out.NetworkConfig = direct.Enum_ToProto[pb.InstanceConfig_NetworkConfig](mapCtx, in.NetworkConfig)
	out.NetworkTemplate = direct.ValueOf(in.NetworkTemplate)
	out.LogicalInterfaces = direct.Slice_ToProto(mapCtx, in.LogicalInterfaces, LogicalInterface_ToProto)
	out.SshKeyNames = in.SSHKeyNames
	return out
}
func InstanceConfigObservedState_FromProto(mapCtx *direct.MapContext, in *pb.InstanceConfig) *krm.InstanceConfigObservedState {
	if in == nil {
		return nil
	}
	out := &krm.InstanceConfigObservedState{}
	out.Name = direct.LazyPtr(in.GetName())
	// MISSING: ID
	// MISSING: InstanceType
	// MISSING: Hyperthreading
	// MISSING: OsImage
	// MISSING: ClientNetwork
	// MISSING: PrivateNetwork
	// MISSING: UserNote
	// MISSING: AccountNetworksEnabled
	// MISSING: NetworkConfig
	// MISSING: NetworkTemplate
	// MISSING: LogicalInterfaces
	// MISSING: SSHKeyNames
	return out
}
func InstanceConfigObservedState_ToProto(mapCtx *direct.MapContext, in *krm.InstanceConfigObservedState) *pb.InstanceConfig {
	if in == nil {
		return nil
	}
	out := &pb.InstanceConfig{}
	out.Name = direct.ValueOf(in.Name)
	// MISSING: ID
	// MISSING: InstanceType
	// MISSING: Hyperthreading
	// MISSING: OsImage
	// MISSING: ClientNetwork
	// MISSING: PrivateNetwork
	// MISSING: UserNote
	// MISSING: AccountNetworksEnabled
	// MISSING: NetworkConfig
	// MISSING: NetworkTemplate
	// MISSING: LogicalInterfaces
	// MISSING: SSHKeyNames
	return out
}
func InstanceConfig_NetworkAddress_FromProto(mapCtx *direct.MapContext, in *pb.InstanceConfig_NetworkAddress) *krm.InstanceConfig_NetworkAddress {
	if in == nil {
		return nil
	}
	out := &krm.InstanceConfig_NetworkAddress{}
	out.NetworkID = direct.LazyPtr(in.GetNetworkId())
	out.Address = direct.LazyPtr(in.GetAddress())
	out.ExistingNetworkID = direct.LazyPtr(in.GetExistingNetworkId())
	return out
}
func InstanceConfig_NetworkAddress_ToProto(mapCtx *direct.MapContext, in *krm.InstanceConfig_NetworkAddress) *pb.InstanceConfig_NetworkAddress {
	if in == nil {
		return nil
	}
	out := &pb.InstanceConfig_NetworkAddress{}
	out.NetworkId = direct.ValueOf(in.NetworkID)
	out.Address = direct.ValueOf(in.Address)
	out.ExistingNetworkId = direct.ValueOf(in.ExistingNetworkID)
	return out
}
func LogicalInterface_FromProto(mapCtx *direct.MapContext, in *pb.LogicalInterface) *krm.LogicalInterface {
	if in == nil {
		return nil
	}
	out := &krm.LogicalInterface{}
	out.LogicalNetworkInterfaces = direct.Slice_FromProto(mapCtx, in.LogicalNetworkInterfaces, LogicalInterface_LogicalNetworkInterface_FromProto)
	out.Name = direct.LazyPtr(in.GetName())
	out.InterfaceIndex = direct.LazyPtr(in.GetInterfaceIndex())
	return out
}
func LogicalInterface_ToProto(mapCtx *direct.MapContext, in *krm.LogicalInterface) *pb.LogicalInterface {
	if in == nil {
		return nil
	}
	out := &pb.LogicalInterface{}
	out.LogicalNetworkInterfaces = direct.Slice_ToProto(mapCtx, in.LogicalNetworkInterfaces, LogicalInterface_LogicalNetworkInterface_ToProto)
	out.Name = direct.ValueOf(in.Name)
	out.InterfaceIndex = direct.ValueOf(in.InterfaceIndex)
	return out
}
func LogicalInterface_LogicalNetworkInterface_FromProto(mapCtx *direct.MapContext, in *pb.LogicalInterface_LogicalNetworkInterface) *krm.LogicalInterface_LogicalNetworkInterface {
	if in == nil {
		return nil
	}
	out := &krm.LogicalInterface_LogicalNetworkInterface{}
	out.Network = direct.LazyPtr(in.GetNetwork())
	out.IPAddress = direct.LazyPtr(in.GetIpAddress())
	out.DefaultGateway = direct.LazyPtr(in.GetDefaultGateway())
	out.NetworkType = direct.Enum_FromProto(mapCtx, in.GetNetworkType())
	out.ID = direct.LazyPtr(in.GetId())
	return out
}
func LogicalInterface_LogicalNetworkInterface_ToProto(mapCtx *direct.MapContext, in *krm.LogicalInterface_LogicalNetworkInterface) *pb.LogicalInterface_LogicalNetworkInterface {
	if in == nil {
		return nil
	}
	out := &pb.LogicalInterface_LogicalNetworkInterface{}
	out.Network = direct.ValueOf(in.Network)
	out.IpAddress = direct.ValueOf(in.IPAddress)
	out.DefaultGateway = direct.ValueOf(in.DefaultGateway)
	out.NetworkType = direct.Enum_ToProto[pb.Network_Type](mapCtx, in.NetworkType)
	out.Id = direct.ValueOf(in.ID)
	return out
}
func NetworkConfig_FromProto(mapCtx *direct.MapContext, in *pb.NetworkConfig) *krm.NetworkConfig {
	if in == nil {
		return nil
	}
	out := &krm.NetworkConfig{}
	// MISSING: Name
	out.ID = direct.LazyPtr(in.GetId())
	out.Type = direct.Enum_FromProto(mapCtx, in.GetType())
	out.Bandwidth = direct.Enum_FromProto(mapCtx, in.GetBandwidth())
	out.VlanAttachments = direct.Slice_FromProto(mapCtx, in.VlanAttachments, NetworkConfig_IntakeVlanAttachment_FromProto)
	out.Cidr = direct.LazyPtr(in.GetCidr())
	out.ServiceCidr = direct.Enum_FromProto(mapCtx, in.GetServiceCidr())
	out.UserNote = direct.LazyPtr(in.GetUserNote())
	out.GcpService = direct.LazyPtr(in.GetGcpService())
	out.VlanSameProject = direct.LazyPtr(in.GetVlanSameProject())
	out.JumboFramesEnabled = direct.LazyPtr(in.GetJumboFramesEnabled())
	return out
}
func NetworkConfig_ToProto(mapCtx *direct.MapContext, in *krm.NetworkConfig) *pb.NetworkConfig {
	if in == nil {
		return nil
	}
	out := &pb.NetworkConfig{}
	// MISSING: Name
	out.Id = direct.ValueOf(in.ID)
	out.Type = direct.Enum_ToProto[pb.NetworkConfig_Type](mapCtx, in.Type)
	out.Bandwidth = direct.Enum_ToProto[pb.NetworkConfig_Bandwidth](mapCtx, in.Bandwidth)
	out.VlanAttachments = direct.Slice_ToProto(mapCtx, in.VlanAttachments, NetworkConfig_IntakeVlanAttachment_ToProto)
	out.Cidr = direct.ValueOf(in.Cidr)
	out.ServiceCidr = direct.Enum_ToProto[pb.NetworkConfig_ServiceCidr](mapCtx, in.ServiceCidr)
	out.UserNote = direct.ValueOf(in.UserNote)
	out.GcpService = direct.ValueOf(in.GcpService)
	out.VlanSameProject = direct.ValueOf(in.VlanSameProject)
	out.JumboFramesEnabled = direct.ValueOf(in.JumboFramesEnabled)
	return out
}
func NetworkConfigObservedState_FromProto(mapCtx *direct.MapContext, in *pb.NetworkConfig) *krm.NetworkConfigObservedState {
	if in == nil {
		return nil
	}
	out := &krm.NetworkConfigObservedState{}
	out.Name = direct.LazyPtr(in.GetName())
	// MISSING: ID
	// MISSING: Type
	// MISSING: Bandwidth
	// MISSING: VlanAttachments
	// MISSING: Cidr
	// MISSING: ServiceCidr
	// MISSING: UserNote
	// MISSING: GcpService
	// MISSING: VlanSameProject
	// MISSING: JumboFramesEnabled
	return out
}
func NetworkConfigObservedState_ToProto(mapCtx *direct.MapContext, in *krm.NetworkConfigObservedState) *pb.NetworkConfig {
	if in == nil {
		return nil
	}
	out := &pb.NetworkConfig{}
	out.Name = direct.ValueOf(in.Name)
	// MISSING: ID
	// MISSING: Type
	// MISSING: Bandwidth
	// MISSING: VlanAttachments
	// MISSING: Cidr
	// MISSING: ServiceCidr
	// MISSING: UserNote
	// MISSING: GcpService
	// MISSING: VlanSameProject
	// MISSING: JumboFramesEnabled
	return out
}
func NetworkConfig_IntakeVlanAttachment_FromProto(mapCtx *direct.MapContext, in *pb.NetworkConfig_IntakeVlanAttachment) *krm.NetworkConfig_IntakeVlanAttachment {
	if in == nil {
		return nil
	}
	out := &krm.NetworkConfig_IntakeVlanAttachment{}
	out.ID = direct.LazyPtr(in.GetId())
	out.PairingKey = direct.LazyPtr(in.GetPairingKey())
	return out
}
func NetworkConfig_IntakeVlanAttachment_ToProto(mapCtx *direct.MapContext, in *krm.NetworkConfig_IntakeVlanAttachment) *pb.NetworkConfig_IntakeVlanAttachment {
	if in == nil {
		return nil
	}
	out := &pb.NetworkConfig_IntakeVlanAttachment{}
	out.Id = direct.ValueOf(in.ID)
	out.PairingKey = direct.ValueOf(in.PairingKey)
	return out
}
func ProvisioningConfig_FromProto(mapCtx *direct.MapContext, in *pb.ProvisioningConfig) *krm.ProvisioningConfig {
	if in == nil {
		return nil
	}
	out := &krm.ProvisioningConfig{}
	// MISSING: Name
	out.Instances = direct.Slice_FromProto(mapCtx, in.Instances, InstanceConfig_FromProto)
	out.Networks = direct.Slice_FromProto(mapCtx, in.Networks, NetworkConfig_FromProto)
	out.Volumes = direct.Slice_FromProto(mapCtx, in.Volumes, VolumeConfig_FromProto)
	out.TicketID = direct.LazyPtr(in.GetTicketId())
	out.HandoverServiceAccount = direct.LazyPtr(in.GetHandoverServiceAccount())
	out.Email = direct.LazyPtr(in.GetEmail())
	// MISSING: State
	out.Location = direct.LazyPtr(in.GetLocation())
	// MISSING: UpdateTime
	// MISSING: CloudConsoleURI
	out.VpcScEnabled = direct.LazyPtr(in.GetVpcScEnabled())
	out.StatusMessage = direct.LazyPtr(in.GetStatusMessage())
	out.CustomID = direct.LazyPtr(in.GetCustomId())
	return out
}
func ProvisioningConfig_ToProto(mapCtx *direct.MapContext, in *krm.ProvisioningConfig) *pb.ProvisioningConfig {
	if in == nil {
		return nil
	}
	out := &pb.ProvisioningConfig{}
	// MISSING: Name
	out.Instances = direct.Slice_ToProto(mapCtx, in.Instances, InstanceConfig_ToProto)
	out.Networks = direct.Slice_ToProto(mapCtx, in.Networks, NetworkConfig_ToProto)
	out.Volumes = direct.Slice_ToProto(mapCtx, in.Volumes, VolumeConfig_ToProto)
	out.TicketId = direct.ValueOf(in.TicketID)
	out.HandoverServiceAccount = direct.ValueOf(in.HandoverServiceAccount)
	out.Email = direct.ValueOf(in.Email)
	// MISSING: State
	out.Location = direct.ValueOf(in.Location)
	// MISSING: UpdateTime
	// MISSING: CloudConsoleURI
	out.VpcScEnabled = direct.ValueOf(in.VpcScEnabled)
	out.StatusMessage = direct.ValueOf(in.StatusMessage)
	out.CustomId = direct.ValueOf(in.CustomID)
	return out
}
func ProvisioningConfigObservedState_FromProto(mapCtx *direct.MapContext, in *pb.ProvisioningConfig) *krm.ProvisioningConfigObservedState {
	if in == nil {
		return nil
	}
	out := &krm.ProvisioningConfigObservedState{}
	out.Name = direct.LazyPtr(in.GetName())
	out.Instances = direct.Slice_FromProto(mapCtx, in.Instances, InstanceConfigObservedState_FromProto)
	out.Networks = direct.Slice_FromProto(mapCtx, in.Networks, NetworkConfigObservedState_FromProto)
	out.Volumes = direct.Slice_FromProto(mapCtx, in.Volumes, VolumeConfigObservedState_FromProto)
	// MISSING: TicketID
	// MISSING: HandoverServiceAccount
	// MISSING: Email
	out.State = direct.Enum_FromProto(mapCtx, in.GetState())
	// MISSING: Location
	out.UpdateTime = direct.StringTimestamp_FromProto(mapCtx, in.GetUpdateTime())
	out.CloudConsoleURI = direct.LazyPtr(in.GetCloudConsoleUri())
	// MISSING: VpcScEnabled
	// MISSING: StatusMessage
	// MISSING: CustomID
	return out
}
func ProvisioningConfigObservedState_ToProto(mapCtx *direct.MapContext, in *krm.ProvisioningConfigObservedState) *pb.ProvisioningConfig {
	if in == nil {
		return nil
	}
	out := &pb.ProvisioningConfig{}
	out.Name = direct.ValueOf(in.Name)
	out.Instances = direct.Slice_ToProto(mapCtx, in.Instances, InstanceConfigObservedState_ToProto)
	out.Networks = direct.Slice_ToProto(mapCtx, in.Networks, NetworkConfigObservedState_ToProto)
	out.Volumes = direct.Slice_ToProto(mapCtx, in.Volumes, VolumeConfigObservedState_ToProto)
	// MISSING: TicketID
	// MISSING: HandoverServiceAccount
	// MISSING: Email
	out.State = direct.Enum_ToProto[pb.ProvisioningConfig_State](mapCtx, in.State)
	// MISSING: Location
	out.UpdateTime = direct.StringTimestamp_ToProto(mapCtx, in.UpdateTime)
	out.CloudConsoleUri = direct.ValueOf(in.CloudConsoleURI)
	// MISSING: VpcScEnabled
	// MISSING: StatusMessage
	// MISSING: CustomID
	return out
}
func VolumeConfig_FromProto(mapCtx *direct.MapContext, in *pb.VolumeConfig) *krm.VolumeConfig {
	if in == nil {
		return nil
	}
	out := &krm.VolumeConfig{}
	// MISSING: Name
	out.ID = direct.LazyPtr(in.GetId())
	out.SnapshotsEnabled = direct.LazyPtr(in.GetSnapshotsEnabled())
	out.Type = direct.Enum_FromProto(mapCtx, in.GetType())
	out.Protocol = direct.Enum_FromProto(mapCtx, in.GetProtocol())
	out.SizeGB = direct.LazyPtr(in.GetSizeGb())
	out.LunRanges = direct.Slice_FromProto(mapCtx, in.LunRanges, VolumeConfig_LunRange_FromProto)
	out.MachineIds = in.MachineIds
	out.NfsExports = direct.Slice_FromProto(mapCtx, in.NfsExports, VolumeConfig_NfsExport_FromProto)
	out.UserNote = direct.LazyPtr(in.GetUserNote())
	out.GcpService = direct.LazyPtr(in.GetGcpService())
	out.PerformanceTier = direct.Enum_FromProto(mapCtx, in.GetPerformanceTier())
	return out
}
func VolumeConfig_ToProto(mapCtx *direct.MapContext, in *krm.VolumeConfig) *pb.VolumeConfig {
	if in == nil {
		return nil
	}
	out := &pb.VolumeConfig{}
	// MISSING: Name
	out.Id = direct.ValueOf(in.ID)
	out.SnapshotsEnabled = direct.ValueOf(in.SnapshotsEnabled)
	out.Type = direct.Enum_ToProto[pb.VolumeConfig_Type](mapCtx, in.Type)
	out.Protocol = direct.Enum_ToProto[pb.VolumeConfig_Protocol](mapCtx, in.Protocol)
	out.SizeGb = direct.ValueOf(in.SizeGB)
	out.LunRanges = direct.Slice_ToProto(mapCtx, in.LunRanges, VolumeConfig_LunRange_ToProto)
	out.MachineIds = in.MachineIds
	out.NfsExports = direct.Slice_ToProto(mapCtx, in.NfsExports, VolumeConfig_NfsExport_ToProto)
	out.UserNote = direct.ValueOf(in.UserNote)
	out.GcpService = direct.ValueOf(in.GcpService)
	out.PerformanceTier = direct.Enum_ToProto[pb.VolumePerformanceTier](mapCtx, in.PerformanceTier)
	return out
}
func VolumeConfigObservedState_FromProto(mapCtx *direct.MapContext, in *pb.VolumeConfig) *krm.VolumeConfigObservedState {
	if in == nil {
		return nil
	}
	out := &krm.VolumeConfigObservedState{}
	out.Name = direct.LazyPtr(in.GetName())
	// MISSING: ID
	// MISSING: SnapshotsEnabled
	// MISSING: Type
	// MISSING: Protocol
	// MISSING: SizeGB
	// MISSING: LunRanges
	// MISSING: MachineIds
	// MISSING: NfsExports
	// MISSING: UserNote
	// MISSING: GcpService
	// MISSING: PerformanceTier
	return out
}
func VolumeConfigObservedState_ToProto(mapCtx *direct.MapContext, in *krm.VolumeConfigObservedState) *pb.VolumeConfig {
	if in == nil {
		return nil
	}
	out := &pb.VolumeConfig{}
	out.Name = direct.ValueOf(in.Name)
	// MISSING: ID
	// MISSING: SnapshotsEnabled
	// MISSING: Type
	// MISSING: Protocol
	// MISSING: SizeGB
	// MISSING: LunRanges
	// MISSING: MachineIds
	// MISSING: NfsExports
	// MISSING: UserNote
	// MISSING: GcpService
	// MISSING: PerformanceTier
	return out
}
func VolumeConfig_LunRange_FromProto(mapCtx *direct.MapContext, in *pb.VolumeConfig_LunRange) *krm.VolumeConfig_LunRange {
	if in == nil {
		return nil
	}
	out := &krm.VolumeConfig_LunRange{}
	out.Quantity = direct.LazyPtr(in.GetQuantity())
	out.SizeGB = direct.LazyPtr(in.GetSizeGb())
	return out
}
func VolumeConfig_LunRange_ToProto(mapCtx *direct.MapContext, in *krm.VolumeConfig_LunRange) *pb.VolumeConfig_LunRange {
	if in == nil {
		return nil
	}
	out := &pb.VolumeConfig_LunRange{}
	out.Quantity = direct.ValueOf(in.Quantity)
	out.SizeGb = direct.ValueOf(in.SizeGB)
	return out
}
func VolumeConfig_NfsExport_FromProto(mapCtx *direct.MapContext, in *pb.VolumeConfig_NfsExport) *krm.VolumeConfig_NfsExport {
	if in == nil {
		return nil
	}
	out := &krm.VolumeConfig_NfsExport{}
	out.NetworkID = direct.LazyPtr(in.GetNetworkId())
	out.MachineID = direct.LazyPtr(in.GetMachineId())
	out.Cidr = direct.LazyPtr(in.GetCidr())
	out.Permissions = direct.Enum_FromProto(mapCtx, in.GetPermissions())
	out.NoRootSquash = direct.LazyPtr(in.GetNoRootSquash())
	out.AllowSuid = direct.LazyPtr(in.GetAllowSuid())
	out.AllowDev = direct.LazyPtr(in.GetAllowDev())
	return out
}
func VolumeConfig_NfsExport_ToProto(mapCtx *direct.MapContext, in *krm.VolumeConfig_NfsExport) *pb.VolumeConfig_NfsExport {
	if in == nil {
		return nil
	}
	out := &pb.VolumeConfig_NfsExport{}
	out.NetworkId = direct.ValueOf(in.NetworkID)
	if oneof := VolumeConfig_NfsExport_MachineId_ToProto(mapCtx, in.MachineID); oneof != nil {
		out.Client = oneof
	}
	if oneof := VolumeConfig_NfsExport_Cidr_ToProto(mapCtx, in.Cidr); oneof != nil {
		out.Client = oneof
	}
	out.Permissions = direct.Enum_ToProto[pb.VolumeConfig_NfsExport_Permissions](mapCtx, in.Permissions)
	out.NoRootSquash = direct.ValueOf(in.NoRootSquash)
	out.AllowSuid = direct.ValueOf(in.AllowSuid)
	out.AllowDev = direct.ValueOf(in.AllowDev)
	return out
}
