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
func BaremetalsolutionInstanceConfigObservedState_FromProto(mapCtx *direct.MapContext, in *pb.InstanceConfig) *krm.BaremetalsolutionInstanceConfigObservedState {
	if in == nil {
		return nil
	}
	out := &krm.BaremetalsolutionInstanceConfigObservedState{}
	// MISSING: Name
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
func BaremetalsolutionInstanceConfigObservedState_ToProto(mapCtx *direct.MapContext, in *krm.BaremetalsolutionInstanceConfigObservedState) *pb.InstanceConfig {
	if in == nil {
		return nil
	}
	out := &pb.InstanceConfig{}
	// MISSING: Name
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
func BaremetalsolutionInstanceConfigSpec_FromProto(mapCtx *direct.MapContext, in *pb.InstanceConfig) *krm.BaremetalsolutionInstanceConfigSpec {
	if in == nil {
		return nil
	}
	out := &krm.BaremetalsolutionInstanceConfigSpec{}
	// MISSING: Name
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
func BaremetalsolutionInstanceConfigSpec_ToProto(mapCtx *direct.MapContext, in *krm.BaremetalsolutionInstanceConfigSpec) *pb.InstanceConfig {
	if in == nil {
		return nil
	}
	out := &pb.InstanceConfig{}
	// MISSING: Name
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
func BaremetalsolutionProvisioningQuotaObservedState_FromProto(mapCtx *direct.MapContext, in *pb.ProvisioningQuota) *krm.BaremetalsolutionProvisioningQuotaObservedState {
	if in == nil {
		return nil
	}
	out := &krm.BaremetalsolutionProvisioningQuotaObservedState{}
	// MISSING: Name
	// MISSING: AssetType
	// MISSING: GcpService
	// MISSING: Location
	// MISSING: AvailableCount
	// MISSING: InstanceQuota
	// MISSING: ServerCount
	// MISSING: NetworkBandwidth
	// MISSING: StorageGib
	return out
}
func BaremetalsolutionProvisioningQuotaObservedState_ToProto(mapCtx *direct.MapContext, in *krm.BaremetalsolutionProvisioningQuotaObservedState) *pb.ProvisioningQuota {
	if in == nil {
		return nil
	}
	out := &pb.ProvisioningQuota{}
	// MISSING: Name
	// MISSING: AssetType
	// MISSING: GcpService
	// MISSING: Location
	// MISSING: AvailableCount
	// MISSING: InstanceQuota
	// MISSING: ServerCount
	// MISSING: NetworkBandwidth
	// MISSING: StorageGib
	return out
}
func BaremetalsolutionProvisioningQuotaSpec_FromProto(mapCtx *direct.MapContext, in *pb.ProvisioningQuota) *krm.BaremetalsolutionProvisioningQuotaSpec {
	if in == nil {
		return nil
	}
	out := &krm.BaremetalsolutionProvisioningQuotaSpec{}
	// MISSING: Name
	// MISSING: AssetType
	// MISSING: GcpService
	// MISSING: Location
	// MISSING: AvailableCount
	// MISSING: InstanceQuota
	// MISSING: ServerCount
	// MISSING: NetworkBandwidth
	// MISSING: StorageGib
	return out
}
func BaremetalsolutionProvisioningQuotaSpec_ToProto(mapCtx *direct.MapContext, in *krm.BaremetalsolutionProvisioningQuotaSpec) *pb.ProvisioningQuota {
	if in == nil {
		return nil
	}
	out := &pb.ProvisioningQuota{}
	// MISSING: Name
	// MISSING: AssetType
	// MISSING: GcpService
	// MISSING: Location
	// MISSING: AvailableCount
	// MISSING: InstanceQuota
	// MISSING: ServerCount
	// MISSING: NetworkBandwidth
	// MISSING: StorageGib
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
