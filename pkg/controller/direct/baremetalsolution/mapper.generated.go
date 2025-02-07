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
func Instance_FromProto(mapCtx *direct.MapContext, in *pb.Instance) *krm.Instance {
	if in == nil {
		return nil
	}
	out := &krm.Instance{}
	out.Name = direct.LazyPtr(in.GetName())
	// MISSING: ID
	// MISSING: CreateTime
	// MISSING: UpdateTime
	out.MachineType = direct.LazyPtr(in.GetMachineType())
	// MISSING: State
	out.HyperthreadingEnabled = direct.LazyPtr(in.GetHyperthreadingEnabled())
	out.Labels = in.Labels
	out.Luns = direct.Slice_FromProto(mapCtx, in.Luns, Lun_FromProto)
	out.Volumes = direct.Slice_FromProto(mapCtx, in.Volumes, Volume_FromProto)
	// MISSING: Networks
	// MISSING: InteractiveSerialConsoleEnabled
	out.OsImage = direct.LazyPtr(in.GetOsImage())
	out.Pod = direct.LazyPtr(in.GetPod())
	out.NetworkTemplate = direct.LazyPtr(in.GetNetworkTemplate())
	out.LogicalInterfaces = direct.Slice_FromProto(mapCtx, in.LogicalInterfaces, LogicalInterface_FromProto)
	// MISSING: LoginInfo
	out.WorkloadProfile = direct.Enum_FromProto(mapCtx, in.GetWorkloadProfile())
	// MISSING: FirmwareVersion
	return out
}
func Instance_ToProto(mapCtx *direct.MapContext, in *krm.Instance) *pb.Instance {
	if in == nil {
		return nil
	}
	out := &pb.Instance{}
	out.Name = direct.ValueOf(in.Name)
	// MISSING: ID
	// MISSING: CreateTime
	// MISSING: UpdateTime
	out.MachineType = direct.ValueOf(in.MachineType)
	// MISSING: State
	out.HyperthreadingEnabled = direct.ValueOf(in.HyperthreadingEnabled)
	out.Labels = in.Labels
	out.Luns = direct.Slice_ToProto(mapCtx, in.Luns, Lun_ToProto)
	out.Volumes = direct.Slice_ToProto(mapCtx, in.Volumes, Volume_ToProto)
	// MISSING: Networks
	// MISSING: InteractiveSerialConsoleEnabled
	out.OsImage = direct.ValueOf(in.OsImage)
	out.Pod = direct.ValueOf(in.Pod)
	out.NetworkTemplate = direct.ValueOf(in.NetworkTemplate)
	out.LogicalInterfaces = direct.Slice_ToProto(mapCtx, in.LogicalInterfaces, LogicalInterface_ToProto)
	// MISSING: LoginInfo
	out.WorkloadProfile = direct.Enum_ToProto[pb.WorkloadProfile](mapCtx, in.WorkloadProfile)
	// MISSING: FirmwareVersion
	return out
}
func InstanceObservedState_FromProto(mapCtx *direct.MapContext, in *pb.Instance) *krm.InstanceObservedState {
	if in == nil {
		return nil
	}
	out := &krm.InstanceObservedState{}
	// MISSING: Name
	out.ID = direct.LazyPtr(in.GetId())
	out.CreateTime = direct.StringTimestamp_FromProto(mapCtx, in.GetCreateTime())
	out.UpdateTime = direct.StringTimestamp_FromProto(mapCtx, in.GetUpdateTime())
	// MISSING: MachineType
	out.State = direct.Enum_FromProto(mapCtx, in.GetState())
	// MISSING: HyperthreadingEnabled
	// MISSING: Labels
	out.Luns = direct.Slice_FromProto(mapCtx, in.Luns, LunObservedState_FromProto)
	out.Volumes = direct.Slice_FromProto(mapCtx, in.Volumes, VolumeObservedState_FromProto)
	out.Networks = direct.Slice_FromProto(mapCtx, in.Networks, Network_FromProto)
	out.InteractiveSerialConsoleEnabled = direct.LazyPtr(in.GetInteractiveSerialConsoleEnabled())
	// MISSING: OsImage
	// MISSING: Pod
	// MISSING: NetworkTemplate
	// MISSING: LogicalInterfaces
	out.LoginInfo = direct.LazyPtr(in.GetLoginInfo())
	// MISSING: WorkloadProfile
	out.FirmwareVersion = direct.LazyPtr(in.GetFirmwareVersion())
	return out
}
func InstanceObservedState_ToProto(mapCtx *direct.MapContext, in *krm.InstanceObservedState) *pb.Instance {
	if in == nil {
		return nil
	}
	out := &pb.Instance{}
	// MISSING: Name
	out.Id = direct.ValueOf(in.ID)
	out.CreateTime = direct.StringTimestamp_ToProto(mapCtx, in.CreateTime)
	out.UpdateTime = direct.StringTimestamp_ToProto(mapCtx, in.UpdateTime)
	// MISSING: MachineType
	out.State = direct.Enum_ToProto[pb.Instance_State](mapCtx, in.State)
	// MISSING: HyperthreadingEnabled
	// MISSING: Labels
	out.Luns = direct.Slice_ToProto(mapCtx, in.Luns, LunObservedState_ToProto)
	out.Volumes = direct.Slice_ToProto(mapCtx, in.Volumes, VolumeObservedState_ToProto)
	out.Networks = direct.Slice_ToProto(mapCtx, in.Networks, Network_ToProto)
	out.InteractiveSerialConsoleEnabled = direct.ValueOf(in.InteractiveSerialConsoleEnabled)
	// MISSING: OsImage
	// MISSING: Pod
	// MISSING: NetworkTemplate
	// MISSING: LogicalInterfaces
	out.LoginInfo = direct.ValueOf(in.LoginInfo)
	// MISSING: WorkloadProfile
	out.FirmwareVersion = direct.ValueOf(in.FirmwareVersion)
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
func Lun_FromProto(mapCtx *direct.MapContext, in *pb.Lun) *krm.Lun {
	if in == nil {
		return nil
	}
	out := &krm.Lun{}
	// MISSING: Name
	out.ID = direct.LazyPtr(in.GetId())
	out.State = direct.Enum_FromProto(mapCtx, in.GetState())
	out.SizeGB = direct.LazyPtr(in.GetSizeGb())
	out.MultiprotocolType = direct.Enum_FromProto(mapCtx, in.GetMultiprotocolType())
	out.StorageVolume = direct.LazyPtr(in.GetStorageVolume())
	out.Shareable = direct.LazyPtr(in.GetShareable())
	out.BootLun = direct.LazyPtr(in.GetBootLun())
	out.StorageType = direct.Enum_FromProto(mapCtx, in.GetStorageType())
	out.Wwid = direct.LazyPtr(in.GetWwid())
	// MISSING: ExpireTime
	// MISSING: Instances
	return out
}
func Lun_ToProto(mapCtx *direct.MapContext, in *krm.Lun) *pb.Lun {
	if in == nil {
		return nil
	}
	out := &pb.Lun{}
	// MISSING: Name
	out.Id = direct.ValueOf(in.ID)
	out.State = direct.Enum_ToProto[pb.Lun_State](mapCtx, in.State)
	out.SizeGb = direct.ValueOf(in.SizeGB)
	out.MultiprotocolType = direct.Enum_ToProto[pb.Lun_MultiprotocolType](mapCtx, in.MultiprotocolType)
	out.StorageVolume = direct.ValueOf(in.StorageVolume)
	out.Shareable = direct.ValueOf(in.Shareable)
	out.BootLun = direct.ValueOf(in.BootLun)
	out.StorageType = direct.Enum_ToProto[pb.Lun_StorageType](mapCtx, in.StorageType)
	out.Wwid = direct.ValueOf(in.Wwid)
	// MISSING: ExpireTime
	// MISSING: Instances
	return out
}
func LunObservedState_FromProto(mapCtx *direct.MapContext, in *pb.Lun) *krm.LunObservedState {
	if in == nil {
		return nil
	}
	out := &krm.LunObservedState{}
	out.Name = direct.LazyPtr(in.GetName())
	// MISSING: ID
	// MISSING: State
	// MISSING: SizeGB
	// MISSING: MultiprotocolType
	// MISSING: StorageVolume
	// MISSING: Shareable
	// MISSING: BootLun
	// MISSING: StorageType
	// MISSING: Wwid
	out.ExpireTime = direct.StringTimestamp_FromProto(mapCtx, in.GetExpireTime())
	out.Instances = in.Instances
	return out
}
func LunObservedState_ToProto(mapCtx *direct.MapContext, in *krm.LunObservedState) *pb.Lun {
	if in == nil {
		return nil
	}
	out := &pb.Lun{}
	out.Name = direct.ValueOf(in.Name)
	// MISSING: ID
	// MISSING: State
	// MISSING: SizeGB
	// MISSING: MultiprotocolType
	// MISSING: StorageVolume
	// MISSING: Shareable
	// MISSING: BootLun
	// MISSING: StorageType
	// MISSING: Wwid
	out.ExpireTime = direct.StringTimestamp_ToProto(mapCtx, in.ExpireTime)
	out.Instances = in.Instances
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
