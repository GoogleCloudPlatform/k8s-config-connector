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
func BaremetalsolutionInstanceQuotaObservedState_FromProto(mapCtx *direct.MapContext, in *pb.InstanceQuota) *krm.BaremetalsolutionInstanceQuotaObservedState {
	if in == nil {
		return nil
	}
	out := &krm.BaremetalsolutionInstanceQuotaObservedState{}
	// MISSING: Name
	// MISSING: InstanceType
	// MISSING: GcpService
	// MISSING: Location
	// MISSING: AvailableMachineCount
	return out
}
func BaremetalsolutionInstanceQuotaObservedState_ToProto(mapCtx *direct.MapContext, in *krm.BaremetalsolutionInstanceQuotaObservedState) *pb.InstanceQuota {
	if in == nil {
		return nil
	}
	out := &pb.InstanceQuota{}
	// MISSING: Name
	// MISSING: InstanceType
	// MISSING: GcpService
	// MISSING: Location
	// MISSING: AvailableMachineCount
	return out
}
func BaremetalsolutionInstanceQuotaSpec_FromProto(mapCtx *direct.MapContext, in *pb.InstanceQuota) *krm.BaremetalsolutionInstanceQuotaSpec {
	if in == nil {
		return nil
	}
	out := &krm.BaremetalsolutionInstanceQuotaSpec{}
	// MISSING: Name
	// MISSING: InstanceType
	// MISSING: GcpService
	// MISSING: Location
	// MISSING: AvailableMachineCount
	return out
}
func BaremetalsolutionInstanceQuotaSpec_ToProto(mapCtx *direct.MapContext, in *krm.BaremetalsolutionInstanceQuotaSpec) *pb.InstanceQuota {
	if in == nil {
		return nil
	}
	out := &pb.InstanceQuota{}
	// MISSING: Name
	// MISSING: InstanceType
	// MISSING: GcpService
	// MISSING: Location
	// MISSING: AvailableMachineCount
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
func BaremetalsolutionNetworkConfigObservedState_FromProto(mapCtx *direct.MapContext, in *pb.NetworkConfig) *krm.BaremetalsolutionNetworkConfigObservedState {
	if in == nil {
		return nil
	}
	out := &krm.BaremetalsolutionNetworkConfigObservedState{}
	// MISSING: Name
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
func BaremetalsolutionNetworkConfigObservedState_ToProto(mapCtx *direct.MapContext, in *krm.BaremetalsolutionNetworkConfigObservedState) *pb.NetworkConfig {
	if in == nil {
		return nil
	}
	out := &pb.NetworkConfig{}
	// MISSING: Name
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
func BaremetalsolutionNetworkConfigSpec_FromProto(mapCtx *direct.MapContext, in *pb.NetworkConfig) *krm.BaremetalsolutionNetworkConfigSpec {
	if in == nil {
		return nil
	}
	out := &krm.BaremetalsolutionNetworkConfigSpec{}
	// MISSING: Name
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
func BaremetalsolutionNetworkConfigSpec_ToProto(mapCtx *direct.MapContext, in *krm.BaremetalsolutionNetworkConfigSpec) *pb.NetworkConfig {
	if in == nil {
		return nil
	}
	out := &pb.NetworkConfig{}
	// MISSING: Name
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
func BaremetalsolutionSSHKeyObservedState_FromProto(mapCtx *direct.MapContext, in *pb.SSHKey) *krm.BaremetalsolutionSSHKeyObservedState {
	if in == nil {
		return nil
	}
	out := &krm.BaremetalsolutionSSHKeyObservedState{}
	// MISSING: Name
	// MISSING: PublicKey
	return out
}
func BaremetalsolutionSSHKeyObservedState_ToProto(mapCtx *direct.MapContext, in *krm.BaremetalsolutionSSHKeyObservedState) *pb.SSHKey {
	if in == nil {
		return nil
	}
	out := &pb.SSHKey{}
	// MISSING: Name
	// MISSING: PublicKey
	return out
}
func BaremetalsolutionSSHKeySpec_FromProto(mapCtx *direct.MapContext, in *pb.SSHKey) *krm.BaremetalsolutionSSHKeySpec {
	if in == nil {
		return nil
	}
	out := &krm.BaremetalsolutionSSHKeySpec{}
	// MISSING: Name
	// MISSING: PublicKey
	return out
}
func BaremetalsolutionSSHKeySpec_ToProto(mapCtx *direct.MapContext, in *krm.BaremetalsolutionSSHKeySpec) *pb.SSHKey {
	if in == nil {
		return nil
	}
	out := &pb.SSHKey{}
	// MISSING: Name
	// MISSING: PublicKey
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
func BaremetalsolutionVolumeObservedState_FromProto(mapCtx *direct.MapContext, in *pb.Volume) *krm.BaremetalsolutionVolumeObservedState {
	if in == nil {
		return nil
	}
	out := &krm.BaremetalsolutionVolumeObservedState{}
	// MISSING: Name
	// MISSING: ID
	// MISSING: StorageType
	// MISSING: State
	// MISSING: RequestedSizeGib
	// MISSING: OriginallyRequestedSizeGib
	// MISSING: CurrentSizeGib
	// MISSING: EmergencySizeGib
	// MISSING: MaxSizeGib
	// MISSING: AutoGrownSizeGib
	// MISSING: RemainingSpaceGib
	// MISSING: SnapshotReservationDetail
	// MISSING: SnapshotAutoDeleteBehavior
	// MISSING: Labels
	// MISSING: SnapshotEnabled
	// MISSING: Pod
	// MISSING: Protocol
	// MISSING: BootVolume
	// MISSING: PerformanceTier
	// MISSING: Notes
	// MISSING: WorkloadProfile
	// MISSING: ExpireTime
	// MISSING: Instances
	// MISSING: Attached
	return out
}
func BaremetalsolutionVolumeObservedState_ToProto(mapCtx *direct.MapContext, in *krm.BaremetalsolutionVolumeObservedState) *pb.Volume {
	if in == nil {
		return nil
	}
	out := &pb.Volume{}
	// MISSING: Name
	// MISSING: ID
	// MISSING: StorageType
	// MISSING: State
	// MISSING: RequestedSizeGib
	// MISSING: OriginallyRequestedSizeGib
	// MISSING: CurrentSizeGib
	// MISSING: EmergencySizeGib
	// MISSING: MaxSizeGib
	// MISSING: AutoGrownSizeGib
	// MISSING: RemainingSpaceGib
	// MISSING: SnapshotReservationDetail
	// MISSING: SnapshotAutoDeleteBehavior
	// MISSING: Labels
	// MISSING: SnapshotEnabled
	// MISSING: Pod
	// MISSING: Protocol
	// MISSING: BootVolume
	// MISSING: PerformanceTier
	// MISSING: Notes
	// MISSING: WorkloadProfile
	// MISSING: ExpireTime
	// MISSING: Instances
	// MISSING: Attached
	return out
}
func BaremetalsolutionVolumeSpec_FromProto(mapCtx *direct.MapContext, in *pb.Volume) *krm.BaremetalsolutionVolumeSpec {
	if in == nil {
		return nil
	}
	out := &krm.BaremetalsolutionVolumeSpec{}
	// MISSING: Name
	// MISSING: ID
	// MISSING: StorageType
	// MISSING: State
	// MISSING: RequestedSizeGib
	// MISSING: OriginallyRequestedSizeGib
	// MISSING: CurrentSizeGib
	// MISSING: EmergencySizeGib
	// MISSING: MaxSizeGib
	// MISSING: AutoGrownSizeGib
	// MISSING: RemainingSpaceGib
	// MISSING: SnapshotReservationDetail
	// MISSING: SnapshotAutoDeleteBehavior
	// MISSING: Labels
	// MISSING: SnapshotEnabled
	// MISSING: Pod
	// MISSING: Protocol
	// MISSING: BootVolume
	// MISSING: PerformanceTier
	// MISSING: Notes
	// MISSING: WorkloadProfile
	// MISSING: ExpireTime
	// MISSING: Instances
	// MISSING: Attached
	return out
}
func BaremetalsolutionVolumeSpec_ToProto(mapCtx *direct.MapContext, in *krm.BaremetalsolutionVolumeSpec) *pb.Volume {
	if in == nil {
		return nil
	}
	out := &pb.Volume{}
	// MISSING: Name
	// MISSING: ID
	// MISSING: StorageType
	// MISSING: State
	// MISSING: RequestedSizeGib
	// MISSING: OriginallyRequestedSizeGib
	// MISSING: CurrentSizeGib
	// MISSING: EmergencySizeGib
	// MISSING: MaxSizeGib
	// MISSING: AutoGrownSizeGib
	// MISSING: RemainingSpaceGib
	// MISSING: SnapshotReservationDetail
	// MISSING: SnapshotAutoDeleteBehavior
	// MISSING: Labels
	// MISSING: SnapshotEnabled
	// MISSING: Pod
	// MISSING: Protocol
	// MISSING: BootVolume
	// MISSING: PerformanceTier
	// MISSING: Notes
	// MISSING: WorkloadProfile
	// MISSING: ExpireTime
	// MISSING: Instances
	// MISSING: Attached
	return out
}
func Volume_FromProto(mapCtx *direct.MapContext, in *pb.Volume) *krm.Volume {
	if in == nil {
		return nil
	}
	out := &krm.Volume{}
	// MISSING: Name
	out.ID = direct.LazyPtr(in.GetId())
	out.StorageType = direct.Enum_FromProto(mapCtx, in.GetStorageType())
	out.State = direct.Enum_FromProto(mapCtx, in.GetState())
	out.RequestedSizeGib = direct.LazyPtr(in.GetRequestedSizeGib())
	out.OriginallyRequestedSizeGib = direct.LazyPtr(in.GetOriginallyRequestedSizeGib())
	out.CurrentSizeGib = direct.LazyPtr(in.GetCurrentSizeGib())
	out.EmergencySizeGib = direct.LazyPtr(in.GetEmergencySizeGib())
	out.MaxSizeGib = direct.LazyPtr(in.GetMaxSizeGib())
	out.AutoGrownSizeGib = direct.LazyPtr(in.GetAutoGrownSizeGib())
	out.RemainingSpaceGib = direct.LazyPtr(in.GetRemainingSpaceGib())
	out.SnapshotReservationDetail = Volume_SnapshotReservationDetail_FromProto(mapCtx, in.GetSnapshotReservationDetail())
	out.SnapshotAutoDeleteBehavior = direct.Enum_FromProto(mapCtx, in.GetSnapshotAutoDeleteBehavior())
	out.Labels = in.Labels
	out.SnapshotEnabled = direct.LazyPtr(in.GetSnapshotEnabled())
	out.Pod = direct.LazyPtr(in.GetPod())
	// MISSING: Protocol
	// MISSING: BootVolume
	out.PerformanceTier = direct.Enum_FromProto(mapCtx, in.GetPerformanceTier())
	out.Notes = direct.LazyPtr(in.GetNotes())
	out.WorkloadProfile = direct.Enum_FromProto(mapCtx, in.GetWorkloadProfile())
	// MISSING: ExpireTime
	// MISSING: Instances
	// MISSING: Attached
	return out
}
func Volume_ToProto(mapCtx *direct.MapContext, in *krm.Volume) *pb.Volume {
	if in == nil {
		return nil
	}
	out := &pb.Volume{}
	// MISSING: Name
	out.Id = direct.ValueOf(in.ID)
	out.StorageType = direct.Enum_ToProto[pb.Volume_StorageType](mapCtx, in.StorageType)
	out.State = direct.Enum_ToProto[pb.Volume_State](mapCtx, in.State)
	out.RequestedSizeGib = direct.ValueOf(in.RequestedSizeGib)
	out.OriginallyRequestedSizeGib = direct.ValueOf(in.OriginallyRequestedSizeGib)
	out.CurrentSizeGib = direct.ValueOf(in.CurrentSizeGib)
	out.EmergencySizeGib = direct.ValueOf(in.EmergencySizeGib)
	out.MaxSizeGib = direct.ValueOf(in.MaxSizeGib)
	out.AutoGrownSizeGib = direct.ValueOf(in.AutoGrownSizeGib)
	out.RemainingSpaceGib = direct.ValueOf(in.RemainingSpaceGib)
	out.SnapshotReservationDetail = Volume_SnapshotReservationDetail_ToProto(mapCtx, in.SnapshotReservationDetail)
	out.SnapshotAutoDeleteBehavior = direct.Enum_ToProto[pb.Volume_SnapshotAutoDeleteBehavior](mapCtx, in.SnapshotAutoDeleteBehavior)
	out.Labels = in.Labels
	out.SnapshotEnabled = direct.ValueOf(in.SnapshotEnabled)
	out.Pod = direct.ValueOf(in.Pod)
	// MISSING: Protocol
	// MISSING: BootVolume
	out.PerformanceTier = direct.Enum_ToProto[pb.VolumePerformanceTier](mapCtx, in.PerformanceTier)
	out.Notes = direct.ValueOf(in.Notes)
	out.WorkloadProfile = direct.Enum_ToProto[pb.Volume_WorkloadProfile](mapCtx, in.WorkloadProfile)
	// MISSING: ExpireTime
	// MISSING: Instances
	// MISSING: Attached
	return out
}
func VolumeObservedState_FromProto(mapCtx *direct.MapContext, in *pb.Volume) *krm.VolumeObservedState {
	if in == nil {
		return nil
	}
	out := &krm.VolumeObservedState{}
	out.Name = direct.LazyPtr(in.GetName())
	// MISSING: ID
	// MISSING: StorageType
	// MISSING: State
	// MISSING: RequestedSizeGib
	// MISSING: OriginallyRequestedSizeGib
	// MISSING: CurrentSizeGib
	// MISSING: EmergencySizeGib
	// MISSING: MaxSizeGib
	// MISSING: AutoGrownSizeGib
	// MISSING: RemainingSpaceGib
	// MISSING: SnapshotReservationDetail
	// MISSING: SnapshotAutoDeleteBehavior
	// MISSING: Labels
	// MISSING: SnapshotEnabled
	// MISSING: Pod
	out.Protocol = direct.Enum_FromProto(mapCtx, in.GetProtocol())
	out.BootVolume = direct.LazyPtr(in.GetBootVolume())
	// MISSING: PerformanceTier
	// MISSING: Notes
	// MISSING: WorkloadProfile
	out.ExpireTime = direct.StringTimestamp_FromProto(mapCtx, in.GetExpireTime())
	out.Instances = in.Instances
	out.Attached = direct.LazyPtr(in.GetAttached())
	return out
}
func VolumeObservedState_ToProto(mapCtx *direct.MapContext, in *krm.VolumeObservedState) *pb.Volume {
	if in == nil {
		return nil
	}
	out := &pb.Volume{}
	out.Name = direct.ValueOf(in.Name)
	// MISSING: ID
	// MISSING: StorageType
	// MISSING: State
	// MISSING: RequestedSizeGib
	// MISSING: OriginallyRequestedSizeGib
	// MISSING: CurrentSizeGib
	// MISSING: EmergencySizeGib
	// MISSING: MaxSizeGib
	// MISSING: AutoGrownSizeGib
	// MISSING: RemainingSpaceGib
	// MISSING: SnapshotReservationDetail
	// MISSING: SnapshotAutoDeleteBehavior
	// MISSING: Labels
	// MISSING: SnapshotEnabled
	// MISSING: Pod
	out.Protocol = direct.Enum_ToProto[pb.Volume_Protocol](mapCtx, in.Protocol)
	out.BootVolume = direct.ValueOf(in.BootVolume)
	// MISSING: PerformanceTier
	// MISSING: Notes
	// MISSING: WorkloadProfile
	out.ExpireTime = direct.StringTimestamp_ToProto(mapCtx, in.ExpireTime)
	out.Instances = in.Instances
	out.Attached = direct.ValueOf(in.Attached)
	return out
}
func Volume_SnapshotReservationDetail_FromProto(mapCtx *direct.MapContext, in *pb.Volume_SnapshotReservationDetail) *krm.Volume_SnapshotReservationDetail {
	if in == nil {
		return nil
	}
	out := &krm.Volume_SnapshotReservationDetail{}
	out.ReservedSpaceGib = direct.LazyPtr(in.GetReservedSpaceGib())
	out.ReservedSpaceUsedPercent = direct.LazyPtr(in.GetReservedSpaceUsedPercent())
	out.ReservedSpaceRemainingGib = direct.LazyPtr(in.GetReservedSpaceRemainingGib())
	out.ReservedSpacePercent = direct.LazyPtr(in.GetReservedSpacePercent())
	return out
}
func Volume_SnapshotReservationDetail_ToProto(mapCtx *direct.MapContext, in *krm.Volume_SnapshotReservationDetail) *pb.Volume_SnapshotReservationDetail {
	if in == nil {
		return nil
	}
	out := &pb.Volume_SnapshotReservationDetail{}
	out.ReservedSpaceGib = direct.ValueOf(in.ReservedSpaceGib)
	out.ReservedSpaceUsedPercent = direct.ValueOf(in.ReservedSpaceUsedPercent)
	out.ReservedSpaceRemainingGib = direct.ValueOf(in.ReservedSpaceRemainingGib)
	out.ReservedSpacePercent = direct.ValueOf(in.ReservedSpacePercent)
	return out
}
