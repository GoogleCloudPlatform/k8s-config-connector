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
func InstanceQuota_FromProto(mapCtx *direct.MapContext, in *pb.InstanceQuota) *krm.InstanceQuota {
	if in == nil {
		return nil
	}
	out := &krm.InstanceQuota{}
	// MISSING: Name
	out.InstanceType = direct.LazyPtr(in.GetInstanceType())
	out.GcpService = direct.LazyPtr(in.GetGcpService())
	out.Location = direct.LazyPtr(in.GetLocation())
	out.AvailableMachineCount = direct.LazyPtr(in.GetAvailableMachineCount())
	return out
}
func InstanceQuota_ToProto(mapCtx *direct.MapContext, in *krm.InstanceQuota) *pb.InstanceQuota {
	if in == nil {
		return nil
	}
	out := &pb.InstanceQuota{}
	// MISSING: Name
	out.InstanceType = direct.ValueOf(in.InstanceType)
	out.GcpService = direct.ValueOf(in.GcpService)
	out.Location = direct.ValueOf(in.Location)
	out.AvailableMachineCount = direct.ValueOf(in.AvailableMachineCount)
	return out
}
func InstanceQuotaObservedState_FromProto(mapCtx *direct.MapContext, in *pb.InstanceQuota) *krm.InstanceQuotaObservedState {
	if in == nil {
		return nil
	}
	out := &krm.InstanceQuotaObservedState{}
	out.Name = direct.LazyPtr(in.GetName())
	// MISSING: InstanceType
	// MISSING: GcpService
	// MISSING: Location
	// MISSING: AvailableMachineCount
	return out
}
func InstanceQuotaObservedState_ToProto(mapCtx *direct.MapContext, in *krm.InstanceQuotaObservedState) *pb.InstanceQuota {
	if in == nil {
		return nil
	}
	out := &pb.InstanceQuota{}
	out.Name = direct.ValueOf(in.Name)
	// MISSING: InstanceType
	// MISSING: GcpService
	// MISSING: Location
	// MISSING: AvailableMachineCount
	return out
}
func ProvisioningQuota_FromProto(mapCtx *direct.MapContext, in *pb.ProvisioningQuota) *krm.ProvisioningQuota {
	if in == nil {
		return nil
	}
	out := &krm.ProvisioningQuota{}
	// MISSING: Name
	out.AssetType = direct.Enum_FromProto(mapCtx, in.GetAssetType())
	out.GcpService = direct.LazyPtr(in.GetGcpService())
	out.Location = direct.LazyPtr(in.GetLocation())
	out.AvailableCount = direct.LazyPtr(in.GetAvailableCount())
	out.InstanceQuota = InstanceQuota_FromProto(mapCtx, in.GetInstanceQuota())
	out.ServerCount = direct.LazyPtr(in.GetServerCount())
	out.NetworkBandwidth = direct.LazyPtr(in.GetNetworkBandwidth())
	out.StorageGib = direct.LazyPtr(in.GetStorageGib())
	return out
}
func ProvisioningQuota_ToProto(mapCtx *direct.MapContext, in *krm.ProvisioningQuota) *pb.ProvisioningQuota {
	if in == nil {
		return nil
	}
	out := &pb.ProvisioningQuota{}
	// MISSING: Name
	out.AssetType = direct.Enum_ToProto[pb.ProvisioningQuota_AssetType](mapCtx, in.AssetType)
	out.GcpService = direct.ValueOf(in.GcpService)
	out.Location = direct.ValueOf(in.Location)
	out.AvailableCount = direct.ValueOf(in.AvailableCount)
	if oneof := InstanceQuota_ToProto(mapCtx, in.InstanceQuota); oneof != nil {
		out.Quota = &pb.ProvisioningQuota_InstanceQuota{InstanceQuota: oneof}
	}
	if oneof := ProvisioningQuota_ServerCount_ToProto(mapCtx, in.ServerCount); oneof != nil {
		out.Availability = oneof
	}
	if oneof := ProvisioningQuota_NetworkBandwidth_ToProto(mapCtx, in.NetworkBandwidth); oneof != nil {
		out.Availability = oneof
	}
	if oneof := ProvisioningQuota_StorageGib_ToProto(mapCtx, in.StorageGib); oneof != nil {
		out.Availability = oneof
	}
	return out
}
func ProvisioningQuotaObservedState_FromProto(mapCtx *direct.MapContext, in *pb.ProvisioningQuota) *krm.ProvisioningQuotaObservedState {
	if in == nil {
		return nil
	}
	out := &krm.ProvisioningQuotaObservedState{}
	out.Name = direct.LazyPtr(in.GetName())
	// MISSING: AssetType
	// MISSING: GcpService
	// MISSING: Location
	// MISSING: AvailableCount
	out.InstanceQuota = InstanceQuotaObservedState_FromProto(mapCtx, in.GetInstanceQuota())
	// MISSING: ServerCount
	// MISSING: NetworkBandwidth
	// MISSING: StorageGib
	return out
}
func ProvisioningQuotaObservedState_ToProto(mapCtx *direct.MapContext, in *krm.ProvisioningQuotaObservedState) *pb.ProvisioningQuota {
	if in == nil {
		return nil
	}
	out := &pb.ProvisioningQuota{}
	out.Name = direct.ValueOf(in.Name)
	// MISSING: AssetType
	// MISSING: GcpService
	// MISSING: Location
	// MISSING: AvailableCount
	if oneof := InstanceQuotaObservedState_ToProto(mapCtx, in.InstanceQuota); oneof != nil {
		out.Quota = &pb.ProvisioningQuota_InstanceQuota{InstanceQuota: oneof}
	}
	// MISSING: ServerCount
	// MISSING: NetworkBandwidth
	// MISSING: StorageGib
	return out
}
