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
func Network_FromProto(mapCtx *direct.MapContext, in *pb.Network) *krm.Network {
	if in == nil {
		return nil
	}
	out := &krm.Network{}
	// MISSING: Name
	out.ID = direct.LazyPtr(in.GetId())
	out.Type = direct.Enum_FromProto(mapCtx, in.GetType())
	out.IPAddress = direct.LazyPtr(in.GetIpAddress())
	out.MacAddress = in.MacAddress
	out.State = direct.Enum_FromProto(mapCtx, in.GetState())
	out.VlanID = direct.LazyPtr(in.GetVlanId())
	out.Cidr = direct.LazyPtr(in.GetCidr())
	out.Vrf = VRF_FromProto(mapCtx, in.GetVrf())
	out.Labels = in.Labels
	out.ServicesCidr = direct.LazyPtr(in.GetServicesCidr())
	out.Reservations = direct.Slice_FromProto(mapCtx, in.Reservations, NetworkAddressReservation_FromProto)
	// MISSING: Pod
	out.MountPoints = direct.Slice_FromProto(mapCtx, in.MountPoints, NetworkMountPoint_FromProto)
	out.JumboFramesEnabled = direct.LazyPtr(in.GetJumboFramesEnabled())
	// MISSING: GatewayIP
	return out
}
func Network_ToProto(mapCtx *direct.MapContext, in *krm.Network) *pb.Network {
	if in == nil {
		return nil
	}
	out := &pb.Network{}
	// MISSING: Name
	out.Id = direct.ValueOf(in.ID)
	out.Type = direct.Enum_ToProto[pb.Network_Type](mapCtx, in.Type)
	out.IpAddress = direct.ValueOf(in.IPAddress)
	out.MacAddress = in.MacAddress
	out.State = direct.Enum_ToProto[pb.Network_State](mapCtx, in.State)
	out.VlanId = direct.ValueOf(in.VlanID)
	out.Cidr = direct.ValueOf(in.Cidr)
	out.Vrf = VRF_ToProto(mapCtx, in.Vrf)
	out.Labels = in.Labels
	out.ServicesCidr = direct.ValueOf(in.ServicesCidr)
	out.Reservations = direct.Slice_ToProto(mapCtx, in.Reservations, NetworkAddressReservation_ToProto)
	// MISSING: Pod
	out.MountPoints = direct.Slice_ToProto(mapCtx, in.MountPoints, NetworkMountPoint_ToProto)
	out.JumboFramesEnabled = direct.ValueOf(in.JumboFramesEnabled)
	// MISSING: GatewayIP
	return out
}
func NetworkAddressReservation_FromProto(mapCtx *direct.MapContext, in *pb.NetworkAddressReservation) *krm.NetworkAddressReservation {
	if in == nil {
		return nil
	}
	out := &krm.NetworkAddressReservation{}
	out.StartAddress = direct.LazyPtr(in.GetStartAddress())
	out.EndAddress = direct.LazyPtr(in.GetEndAddress())
	out.Note = direct.LazyPtr(in.GetNote())
	return out
}
func NetworkAddressReservation_ToProto(mapCtx *direct.MapContext, in *krm.NetworkAddressReservation) *pb.NetworkAddressReservation {
	if in == nil {
		return nil
	}
	out := &pb.NetworkAddressReservation{}
	out.StartAddress = direct.ValueOf(in.StartAddress)
	out.EndAddress = direct.ValueOf(in.EndAddress)
	out.Note = direct.ValueOf(in.Note)
	return out
}
func NetworkMountPoint_FromProto(mapCtx *direct.MapContext, in *pb.NetworkMountPoint) *krm.NetworkMountPoint {
	if in == nil {
		return nil
	}
	out := &krm.NetworkMountPoint{}
	out.Instance = direct.LazyPtr(in.GetInstance())
	out.LogicalInterface = direct.LazyPtr(in.GetLogicalInterface())
	out.DefaultGateway = direct.LazyPtr(in.GetDefaultGateway())
	out.IPAddress = direct.LazyPtr(in.GetIpAddress())
	return out
}
func NetworkMountPoint_ToProto(mapCtx *direct.MapContext, in *krm.NetworkMountPoint) *pb.NetworkMountPoint {
	if in == nil {
		return nil
	}
	out := &pb.NetworkMountPoint{}
	out.Instance = direct.ValueOf(in.Instance)
	out.LogicalInterface = direct.ValueOf(in.LogicalInterface)
	out.DefaultGateway = direct.ValueOf(in.DefaultGateway)
	out.IpAddress = direct.ValueOf(in.IPAddress)
	return out
}
func NetworkObservedState_FromProto(mapCtx *direct.MapContext, in *pb.Network) *krm.NetworkObservedState {
	if in == nil {
		return nil
	}
	out := &krm.NetworkObservedState{}
	out.Name = direct.LazyPtr(in.GetName())
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
	out.Pod = direct.LazyPtr(in.GetPod())
	// MISSING: MountPoints
	// MISSING: JumboFramesEnabled
	out.GatewayIP = direct.LazyPtr(in.GetGatewayIp())
	return out
}
func NetworkObservedState_ToProto(mapCtx *direct.MapContext, in *krm.NetworkObservedState) *pb.Network {
	if in == nil {
		return nil
	}
	out := &pb.Network{}
	out.Name = direct.ValueOf(in.Name)
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
	out.Pod = direct.ValueOf(in.Pod)
	// MISSING: MountPoints
	// MISSING: JumboFramesEnabled
	out.GatewayIp = direct.ValueOf(in.GatewayIP)
	return out
}
func VRF_FromProto(mapCtx *direct.MapContext, in *pb.VRF) *krm.VRF {
	if in == nil {
		return nil
	}
	out := &krm.VRF{}
	out.Name = direct.LazyPtr(in.GetName())
	out.State = direct.Enum_FromProto(mapCtx, in.GetState())
	out.QosPolicy = VRF_QosPolicy_FromProto(mapCtx, in.GetQosPolicy())
	out.VlanAttachments = direct.Slice_FromProto(mapCtx, in.VlanAttachments, VRF_VlanAttachment_FromProto)
	return out
}
func VRF_ToProto(mapCtx *direct.MapContext, in *krm.VRF) *pb.VRF {
	if in == nil {
		return nil
	}
	out := &pb.VRF{}
	out.Name = direct.ValueOf(in.Name)
	out.State = direct.Enum_ToProto[pb.VRF_State](mapCtx, in.State)
	out.QosPolicy = VRF_QosPolicy_ToProto(mapCtx, in.QosPolicy)
	out.VlanAttachments = direct.Slice_ToProto(mapCtx, in.VlanAttachments, VRF_VlanAttachment_ToProto)
	return out
}
func VRF_QosPolicy_FromProto(mapCtx *direct.MapContext, in *pb.VRF_QosPolicy) *krm.VRF_QosPolicy {
	if in == nil {
		return nil
	}
	out := &krm.VRF_QosPolicy{}
	out.BandwidthGbps = direct.LazyPtr(in.GetBandwidthGbps())
	return out
}
func VRF_QosPolicy_ToProto(mapCtx *direct.MapContext, in *krm.VRF_QosPolicy) *pb.VRF_QosPolicy {
	if in == nil {
		return nil
	}
	out := &pb.VRF_QosPolicy{}
	out.BandwidthGbps = direct.ValueOf(in.BandwidthGbps)
	return out
}
func VRF_VlanAttachment_FromProto(mapCtx *direct.MapContext, in *pb.VRF_VlanAttachment) *krm.VRF_VlanAttachment {
	if in == nil {
		return nil
	}
	out := &krm.VRF_VlanAttachment{}
	out.PeerVlanID = direct.LazyPtr(in.GetPeerVlanId())
	out.PeerIP = direct.LazyPtr(in.GetPeerIp())
	out.RouterIP = direct.LazyPtr(in.GetRouterIp())
	out.PairingKey = direct.LazyPtr(in.GetPairingKey())
	out.QosPolicy = VRF_QosPolicy_FromProto(mapCtx, in.GetQosPolicy())
	out.ID = direct.LazyPtr(in.GetId())
	out.InterconnectAttachment = direct.LazyPtr(in.GetInterconnectAttachment())
	return out
}
func VRF_VlanAttachment_ToProto(mapCtx *direct.MapContext, in *krm.VRF_VlanAttachment) *pb.VRF_VlanAttachment {
	if in == nil {
		return nil
	}
	out := &pb.VRF_VlanAttachment{}
	out.PeerVlanId = direct.ValueOf(in.PeerVlanID)
	out.PeerIp = direct.ValueOf(in.PeerIP)
	out.RouterIp = direct.ValueOf(in.RouterIP)
	out.PairingKey = direct.ValueOf(in.PairingKey)
	out.QosPolicy = VRF_QosPolicy_ToProto(mapCtx, in.QosPolicy)
	out.Id = direct.ValueOf(in.ID)
	out.InterconnectAttachment = direct.ValueOf(in.InterconnectAttachment)
	return out
}
