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
	pb "cloud.google.com/go/baremetalsolution/apiv2/baremetalsolutionpb"
	krm "github.com/GoogleCloudPlatform/k8s-config-connector/apis/baremetalsolution/v1alpha1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct"
	refs "github.com/GoogleCloudPlatform/k8s-config-connector/apis/refs/v1beta1"
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
func BaremetalsolutionVolumeConfigObservedState_FromProto(mapCtx *direct.MapContext, in *pb.VolumeConfig) *krm.BaremetalsolutionVolumeConfigObservedState {
	if in == nil {
		return nil
	}
	out := &krm.BaremetalsolutionVolumeConfigObservedState{}
	// MISSING: Name
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
func BaremetalsolutionVolumeConfigObservedState_ToProto(mapCtx *direct.MapContext, in *krm.BaremetalsolutionVolumeConfigObservedState) *pb.VolumeConfig {
	if in == nil {
		return nil
	}
	out := &pb.VolumeConfig{}
	// MISSING: Name
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
func BaremetalsolutionVolumeConfigSpec_FromProto(mapCtx *direct.MapContext, in *pb.VolumeConfig) *krm.BaremetalsolutionVolumeConfigSpec {
	if in == nil {
		return nil
	}
	out := &krm.BaremetalsolutionVolumeConfigSpec{}
	// MISSING: Name
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
func BaremetalsolutionVolumeConfigSpec_ToProto(mapCtx *direct.MapContext, in *krm.BaremetalsolutionVolumeConfigSpec) *pb.VolumeConfig {
	if in == nil {
		return nil
	}
	out := &pb.VolumeConfig{}
	// MISSING: Name
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
