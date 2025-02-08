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

package notebooks

import (
	refs "github.com/GoogleCloudPlatform/k8s-config-connector/apis/refs/v1beta1"
	pb "cloud.google.com/go/notebooks/apiv1beta1/notebookspb"
	krm "github.com/GoogleCloudPlatform/k8s-config-connector/apis/notebooks/v1alpha1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct"
)
func ContainerImage_FromProto(mapCtx *direct.MapContext, in *pb.ContainerImage) *krm.ContainerImage {
	if in == nil {
		return nil
	}
	out := &krm.ContainerImage{}
	out.Repository = direct.LazyPtr(in.GetRepository())
	out.Tag = direct.LazyPtr(in.GetTag())
	return out
}
func ContainerImage_ToProto(mapCtx *direct.MapContext, in *krm.ContainerImage) *pb.ContainerImage {
	if in == nil {
		return nil
	}
	out := &pb.ContainerImage{}
	out.Repository = direct.ValueOf(in.Repository)
	out.Tag = direct.ValueOf(in.Tag)
	return out
}
func Instance_FromProto(mapCtx *direct.MapContext, in *pb.Instance) *krm.Instance {
	if in == nil {
		return nil
	}
	out := &krm.Instance{}
	// MISSING: Name
	out.VmImage = VmImage_FromProto(mapCtx, in.GetVmImage())
	out.ContainerImage = ContainerImage_FromProto(mapCtx, in.GetContainerImage())
	out.PostStartupScript = direct.LazyPtr(in.GetPostStartupScript())
	// MISSING: ProxyURI
	out.InstanceOwners = in.InstanceOwners
	out.ServiceAccount = direct.LazyPtr(in.GetServiceAccount())
	out.MachineType = direct.LazyPtr(in.GetMachineType())
	out.AcceleratorConfig = Instance_AcceleratorConfig_FromProto(mapCtx, in.GetAcceleratorConfig())
	// MISSING: State
	out.InstallGpuDriver = direct.LazyPtr(in.GetInstallGpuDriver())
	out.CustomGpuDriverPath = direct.LazyPtr(in.GetCustomGpuDriverPath())
	out.BootDiskType = direct.Enum_FromProto(mapCtx, in.GetBootDiskType())
	out.BootDiskSizeGB = direct.LazyPtr(in.GetBootDiskSizeGb())
	out.DataDiskType = direct.Enum_FromProto(mapCtx, in.GetDataDiskType())
	out.DataDiskSizeGB = direct.LazyPtr(in.GetDataDiskSizeGb())
	out.NoRemoveDataDisk = direct.LazyPtr(in.GetNoRemoveDataDisk())
	out.DiskEncryption = direct.Enum_FromProto(mapCtx, in.GetDiskEncryption())
	out.KMSKey = direct.LazyPtr(in.GetKmsKey())
	out.NoPublicIP = direct.LazyPtr(in.GetNoPublicIp())
	out.NoProxyAccess = direct.LazyPtr(in.GetNoProxyAccess())
	out.Network = direct.LazyPtr(in.GetNetwork())
	out.Subnet = direct.LazyPtr(in.GetSubnet())
	out.Labels = in.Labels
	out.Metadata = in.Metadata
	out.NicType = direct.Enum_FromProto(mapCtx, in.GetNicType())
	out.ReservationAffinity = ReservationAffinity_FromProto(mapCtx, in.GetReservationAffinity())
	out.CanIPForward = direct.LazyPtr(in.GetCanIpForward())
	// MISSING: CreateTime
	// MISSING: UpdateTime
	return out
}
func Instance_ToProto(mapCtx *direct.MapContext, in *krm.Instance) *pb.Instance {
	if in == nil {
		return nil
	}
	out := &pb.Instance{}
	// MISSING: Name
	if oneof := VmImage_ToProto(mapCtx, in.VmImage); oneof != nil {
		out.Environment = &pb.Instance_VmImage{VmImage: oneof}
	}
	if oneof := ContainerImage_ToProto(mapCtx, in.ContainerImage); oneof != nil {
		out.Environment = &pb.Instance_ContainerImage{ContainerImage: oneof}
	}
	out.PostStartupScript = direct.ValueOf(in.PostStartupScript)
	// MISSING: ProxyURI
	out.InstanceOwners = in.InstanceOwners
	out.ServiceAccount = direct.ValueOf(in.ServiceAccount)
	out.MachineType = direct.ValueOf(in.MachineType)
	out.AcceleratorConfig = Instance_AcceleratorConfig_ToProto(mapCtx, in.AcceleratorConfig)
	// MISSING: State
	out.InstallGpuDriver = direct.ValueOf(in.InstallGpuDriver)
	out.CustomGpuDriverPath = direct.ValueOf(in.CustomGpuDriverPath)
	out.BootDiskType = direct.Enum_ToProto[pb.Instance_DiskType](mapCtx, in.BootDiskType)
	out.BootDiskSizeGb = direct.ValueOf(in.BootDiskSizeGB)
	out.DataDiskType = direct.Enum_ToProto[pb.Instance_DiskType](mapCtx, in.DataDiskType)
	out.DataDiskSizeGb = direct.ValueOf(in.DataDiskSizeGB)
	out.NoRemoveDataDisk = direct.ValueOf(in.NoRemoveDataDisk)
	out.DiskEncryption = direct.Enum_ToProto[pb.Instance_DiskEncryption](mapCtx, in.DiskEncryption)
	out.KmsKey = direct.ValueOf(in.KMSKey)
	out.NoPublicIp = direct.ValueOf(in.NoPublicIP)
	out.NoProxyAccess = direct.ValueOf(in.NoProxyAccess)
	out.Network = direct.ValueOf(in.Network)
	out.Subnet = direct.ValueOf(in.Subnet)
	out.Labels = in.Labels
	out.Metadata = in.Metadata
	out.NicType = direct.Enum_ToProto[pb.Instance_NicType](mapCtx, in.NicType)
	out.ReservationAffinity = ReservationAffinity_ToProto(mapCtx, in.ReservationAffinity)
	out.CanIpForward = direct.ValueOf(in.CanIPForward)
	// MISSING: CreateTime
	// MISSING: UpdateTime
	return out
}
func InstanceObservedState_FromProto(mapCtx *direct.MapContext, in *pb.Instance) *krm.InstanceObservedState {
	if in == nil {
		return nil
	}
	out := &krm.InstanceObservedState{}
	out.Name = direct.LazyPtr(in.GetName())
	// MISSING: VmImage
	// MISSING: ContainerImage
	// MISSING: PostStartupScript
	out.ProxyURI = direct.LazyPtr(in.GetProxyUri())
	// MISSING: InstanceOwners
	// MISSING: ServiceAccount
	// MISSING: MachineType
	// MISSING: AcceleratorConfig
	out.State = direct.Enum_FromProto(mapCtx, in.GetState())
	// MISSING: InstallGpuDriver
	// MISSING: CustomGpuDriverPath
	// MISSING: BootDiskType
	// MISSING: BootDiskSizeGB
	// MISSING: DataDiskType
	// MISSING: DataDiskSizeGB
	// MISSING: NoRemoveDataDisk
	// MISSING: DiskEncryption
	// MISSING: KMSKey
	// MISSING: NoPublicIP
	// MISSING: NoProxyAccess
	// MISSING: Network
	// MISSING: Subnet
	// MISSING: Labels
	// MISSING: Metadata
	// MISSING: NicType
	// MISSING: ReservationAffinity
	// MISSING: CanIPForward
	out.CreateTime = direct.StringTimestamp_FromProto(mapCtx, in.GetCreateTime())
	out.UpdateTime = direct.StringTimestamp_FromProto(mapCtx, in.GetUpdateTime())
	return out
}
func InstanceObservedState_ToProto(mapCtx *direct.MapContext, in *krm.InstanceObservedState) *pb.Instance {
	if in == nil {
		return nil
	}
	out := &pb.Instance{}
	out.Name = direct.ValueOf(in.Name)
	// MISSING: VmImage
	// MISSING: ContainerImage
	// MISSING: PostStartupScript
	out.ProxyUri = direct.ValueOf(in.ProxyURI)
	// MISSING: InstanceOwners
	// MISSING: ServiceAccount
	// MISSING: MachineType
	// MISSING: AcceleratorConfig
	out.State = direct.Enum_ToProto[pb.Instance_State](mapCtx, in.State)
	// MISSING: InstallGpuDriver
	// MISSING: CustomGpuDriverPath
	// MISSING: BootDiskType
	// MISSING: BootDiskSizeGB
	// MISSING: DataDiskType
	// MISSING: DataDiskSizeGB
	// MISSING: NoRemoveDataDisk
	// MISSING: DiskEncryption
	// MISSING: KMSKey
	// MISSING: NoPublicIP
	// MISSING: NoProxyAccess
	// MISSING: Network
	// MISSING: Subnet
	// MISSING: Labels
	// MISSING: Metadata
	// MISSING: NicType
	// MISSING: ReservationAffinity
	// MISSING: CanIPForward
	out.CreateTime = direct.StringTimestamp_ToProto(mapCtx, in.CreateTime)
	out.UpdateTime = direct.StringTimestamp_ToProto(mapCtx, in.UpdateTime)
	return out
}
func Instance_AcceleratorConfig_FromProto(mapCtx *direct.MapContext, in *pb.Instance_AcceleratorConfig) *krm.Instance_AcceleratorConfig {
	if in == nil {
		return nil
	}
	out := &krm.Instance_AcceleratorConfig{}
	out.Type = direct.Enum_FromProto(mapCtx, in.GetType())
	out.CoreCount = direct.LazyPtr(in.GetCoreCount())
	return out
}
func Instance_AcceleratorConfig_ToProto(mapCtx *direct.MapContext, in *krm.Instance_AcceleratorConfig) *pb.Instance_AcceleratorConfig {
	if in == nil {
		return nil
	}
	out := &pb.Instance_AcceleratorConfig{}
	out.Type = direct.Enum_ToProto[pb.Instance_AcceleratorType](mapCtx, in.Type)
	out.CoreCount = direct.ValueOf(in.CoreCount)
	return out
}
func ReservationAffinity_FromProto(mapCtx *direct.MapContext, in *pb.ReservationAffinity) *krm.ReservationAffinity {
	if in == nil {
		return nil
	}
	out := &krm.ReservationAffinity{}
	out.ConsumeReservationType = direct.Enum_FromProto(mapCtx, in.GetConsumeReservationType())
	out.Key = direct.LazyPtr(in.GetKey())
	out.Values = in.Values
	return out
}
func ReservationAffinity_ToProto(mapCtx *direct.MapContext, in *krm.ReservationAffinity) *pb.ReservationAffinity {
	if in == nil {
		return nil
	}
	out := &pb.ReservationAffinity{}
	out.ConsumeReservationType = direct.Enum_ToProto[pb.ReservationAffinity_Type](mapCtx, in.ConsumeReservationType)
	out.Key = direct.ValueOf(in.Key)
	out.Values = in.Values
	return out
}
func VmImage_FromProto(mapCtx *direct.MapContext, in *pb.VmImage) *krm.VmImage {
	if in == nil {
		return nil
	}
	out := &krm.VmImage{}
	out.Project = direct.LazyPtr(in.GetProject())
	out.ImageName = direct.LazyPtr(in.GetImageName())
	out.ImageFamily = direct.LazyPtr(in.GetImageFamily())
	return out
}
func VmImage_ToProto(mapCtx *direct.MapContext, in *krm.VmImage) *pb.VmImage {
	if in == nil {
		return nil
	}
	out := &pb.VmImage{}
	out.Project = direct.ValueOf(in.Project)
	if oneof := VmImage_ImageName_ToProto(mapCtx, in.ImageName); oneof != nil {
		out.Image = oneof
	}
	if oneof := VmImage_ImageFamily_ToProto(mapCtx, in.ImageFamily); oneof != nil {
		out.Image = oneof
	}
	return out
}
