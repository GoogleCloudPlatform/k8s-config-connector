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

// +generated:mapper
// krm.group: notebooks.cnrm.cloud.google.com
// krm.version: v1alpha1
// proto.service: google.cloud.notebooks.v1

package notebooks

import (
	pb "cloud.google.com/go/notebooks/apiv1/notebookspb"
	kmsv1beta1 "github.com/GoogleCloudPlatform/k8s-config-connector/apis/kms/v1beta1"
	krmv1beta1 "github.com/GoogleCloudPlatform/k8s-config-connector/apis/notebooks/v1beta1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/apis/refs/v1beta1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct"
)

func ContainerImage_FromProto(mapCtx *direct.MapContext, in *pb.ContainerImage) *krmv1beta1.ContainerImage {
	if in == nil {
		return nil
	}
	out := &krmv1beta1.ContainerImage{}
	out.Repository = direct.LazyPtr(in.GetRepository())
	out.Tag = direct.LazyPtr(in.GetTag())
	return out
}
func ContainerImage_ToProto(mapCtx *direct.MapContext, in *krmv1beta1.ContainerImage) *pb.ContainerImage {
	if in == nil {
		return nil
	}
	out := &pb.ContainerImage{}
	out.Repository = direct.ValueOf(in.Repository)
	out.Tag = direct.ValueOf(in.Tag)
	return out
}
func Instance_AcceleratorConfig_FromProto(mapCtx *direct.MapContext, in *pb.Instance_AcceleratorConfig) *krmv1beta1.Instance_AcceleratorConfig {
	if in == nil {
		return nil
	}
	out := &krmv1beta1.Instance_AcceleratorConfig{}
	out.Type = direct.Enum_FromProto(mapCtx, in.GetType())
	out.CoreCount = direct.LazyPtr(in.GetCoreCount())
	return out
}
func Instance_AcceleratorConfig_ToProto(mapCtx *direct.MapContext, in *krmv1beta1.Instance_AcceleratorConfig) *pb.Instance_AcceleratorConfig {
	if in == nil {
		return nil
	}
	out := &pb.Instance_AcceleratorConfig{}
	out.Type = direct.Enum_ToProto[pb.Instance_AcceleratorType](mapCtx, in.Type)
	out.CoreCount = direct.ValueOf(in.CoreCount)
	return out
}
func Instance_Disk_FromProto(mapCtx *direct.MapContext, in *pb.Instance_Disk) *krmv1beta1.Instance_Disk {
	if in == nil {
		return nil
	}
	out := &krmv1beta1.Instance_Disk{}
	out.AutoDelete = direct.LazyPtr(in.GetAutoDelete())
	out.Boot = direct.LazyPtr(in.GetBoot())
	out.DeviceName = direct.LazyPtr(in.GetDeviceName())
	out.DiskSizeGB = direct.LazyPtr(in.GetDiskSizeGb())
	out.GuestOSFeatures = direct.Slice_FromProto(mapCtx, in.GuestOsFeatures, Instance_Disk_GuestOSFeature_FromProto)
	out.Index = direct.LazyPtr(in.GetIndex())
	out.Interface = direct.LazyPtr(in.GetInterface())
	out.Kind = direct.LazyPtr(in.GetKind())
	out.Licenses = in.Licenses
	out.Mode = direct.LazyPtr(in.GetMode())
	out.Source = direct.LazyPtr(in.GetSource())
	out.Type = direct.LazyPtr(in.GetType())
	return out
}
func Instance_Disk_ToProto(mapCtx *direct.MapContext, in *krmv1beta1.Instance_Disk) *pb.Instance_Disk {
	if in == nil {
		return nil
	}
	out := &pb.Instance_Disk{}
	out.AutoDelete = direct.ValueOf(in.AutoDelete)
	out.Boot = direct.ValueOf(in.Boot)
	out.DeviceName = direct.ValueOf(in.DeviceName)
	out.DiskSizeGb = direct.ValueOf(in.DiskSizeGB)
	out.GuestOsFeatures = direct.Slice_ToProto(mapCtx, in.GuestOSFeatures, Instance_Disk_GuestOSFeature_ToProto)
	out.Index = direct.ValueOf(in.Index)
	out.Interface = direct.ValueOf(in.Interface)
	out.Kind = direct.ValueOf(in.Kind)
	out.Licenses = in.Licenses
	out.Mode = direct.ValueOf(in.Mode)
	out.Source = direct.ValueOf(in.Source)
	out.Type = direct.ValueOf(in.Type)
	return out
}
func Instance_Disk_GuestOSFeature_FromProto(mapCtx *direct.MapContext, in *pb.Instance_Disk_GuestOsFeature) *krmv1beta1.Instance_Disk_GuestOSFeature {
	if in == nil {
		return nil
	}
	out := &krmv1beta1.Instance_Disk_GuestOSFeature{}
	out.Type = direct.LazyPtr(in.GetType())
	return out
}
func Instance_Disk_GuestOSFeature_ToProto(mapCtx *direct.MapContext, in *krmv1beta1.Instance_Disk_GuestOSFeature) *pb.Instance_Disk_GuestOsFeature {
	if in == nil {
		return nil
	}
	out := &pb.Instance_Disk_GuestOsFeature{}
	out.Type = direct.ValueOf(in.Type)
	return out
}
func Instance_ShieldedInstanceConfig_FromProto(mapCtx *direct.MapContext, in *pb.Instance_ShieldedInstanceConfig) *krmv1beta1.Instance_ShieldedInstanceConfig {
	if in == nil {
		return nil
	}
	out := &krmv1beta1.Instance_ShieldedInstanceConfig{}
	out.EnableSecureBoot = direct.LazyPtr(in.GetEnableSecureBoot())
	out.EnableVTPM = direct.LazyPtr(in.GetEnableVtpm())
	out.EnableIntegrityMonitoring = direct.LazyPtr(in.GetEnableIntegrityMonitoring())
	return out
}
func Instance_ShieldedInstanceConfig_ToProto(mapCtx *direct.MapContext, in *krmv1beta1.Instance_ShieldedInstanceConfig) *pb.Instance_ShieldedInstanceConfig {
	if in == nil {
		return nil
	}
	out := &pb.Instance_ShieldedInstanceConfig{}
	out.EnableSecureBoot = direct.ValueOf(in.EnableSecureBoot)
	out.EnableVtpm = direct.ValueOf(in.EnableVTPM)
	out.EnableIntegrityMonitoring = direct.ValueOf(in.EnableIntegrityMonitoring)
	return out
}
func Instance_UpgradeHistoryEntry_FromProto(mapCtx *direct.MapContext, in *pb.Instance_UpgradeHistoryEntry) *krmv1beta1.Instance_UpgradeHistoryEntry {
	if in == nil {
		return nil
	}
	out := &krmv1beta1.Instance_UpgradeHistoryEntry{}
	out.Snapshot = direct.LazyPtr(in.GetSnapshot())
	out.VMImage = direct.LazyPtr(in.GetVmImage())
	out.ContainerImage = direct.LazyPtr(in.GetContainerImage())
	out.Framework = direct.LazyPtr(in.GetFramework())
	out.Version = direct.LazyPtr(in.GetVersion())
	out.State = direct.Enum_FromProto(mapCtx, in.GetState())
	out.CreateTime = direct.StringTimestamp_FromProto(mapCtx, in.GetCreateTime())
	out.TargetImage = direct.LazyPtr(in.GetTargetImage())
	out.Action = direct.Enum_FromProto(mapCtx, in.GetAction())
	out.TargetVersion = direct.LazyPtr(in.GetTargetVersion())
	return out
}
func Instance_UpgradeHistoryEntry_ToProto(mapCtx *direct.MapContext, in *krmv1beta1.Instance_UpgradeHistoryEntry) *pb.Instance_UpgradeHistoryEntry {
	if in == nil {
		return nil
	}
	out := &pb.Instance_UpgradeHistoryEntry{}
	out.Snapshot = direct.ValueOf(in.Snapshot)
	out.VmImage = direct.ValueOf(in.VMImage)
	out.ContainerImage = direct.ValueOf(in.ContainerImage)
	out.Framework = direct.ValueOf(in.Framework)
	out.Version = direct.ValueOf(in.Version)
	out.State = direct.Enum_ToProto[pb.Instance_UpgradeHistoryEntry_State](mapCtx, in.State)
	out.CreateTime = direct.StringTimestamp_ToProto(mapCtx, in.CreateTime)
	out.TargetImage = direct.ValueOf(in.TargetImage)
	out.Action = direct.Enum_ToProto[pb.Instance_UpgradeHistoryEntry_Action](mapCtx, in.Action)
	out.TargetVersion = direct.ValueOf(in.TargetVersion)
	return out
}
func NotebookInstanceObservedState_FromProto(mapCtx *direct.MapContext, in *pb.Instance) *krmv1beta1.NotebookInstanceObservedState {
	if in == nil {
		return nil
	}
	out := &krmv1beta1.NotebookInstanceObservedState{}
	out.ProxyURI = direct.LazyPtr(in.GetProxyUri())
	out.State = direct.Enum_FromProto(mapCtx, in.GetState())
	out.Disks = direct.Slice_FromProto(mapCtx, in.Disks, Instance_Disk_FromProto)
	out.Creator = direct.LazyPtr(in.GetCreator())
	out.CreateTime = direct.StringTimestamp_FromProto(mapCtx, in.GetCreateTime())
	out.UpdateTime = direct.StringTimestamp_FromProto(mapCtx, in.GetUpdateTime())
	return out
}
func NotebookInstanceObservedState_ToProto(mapCtx *direct.MapContext, in *krmv1beta1.NotebookInstanceObservedState) *pb.Instance {
	if in == nil {
		return nil
	}
	out := &pb.Instance{}
	out.ProxyUri = direct.ValueOf(in.ProxyURI)
	out.State = direct.Enum_ToProto[pb.Instance_State](mapCtx, in.State)
	out.Disks = direct.Slice_ToProto(mapCtx, in.Disks, Instance_Disk_ToProto)
	out.Creator = direct.ValueOf(in.Creator)
	out.CreateTime = direct.StringTimestamp_ToProto(mapCtx, in.CreateTime)
	out.UpdateTime = direct.StringTimestamp_ToProto(mapCtx, in.UpdateTime)
	return out
}
func NotebookInstanceSpec_FromProto(mapCtx *direct.MapContext, in *pb.Instance) *krmv1beta1.NotebookInstanceSpec {
	if in == nil {
		return nil
	}
	out := &krmv1beta1.NotebookInstanceSpec{}
	out.VMImage = VMImage_FromProto(mapCtx, in.GetVmImage())
	out.ContainerImage = ContainerImage_FromProto(mapCtx, in.GetContainerImage())
	out.PostStartupScript = direct.LazyPtr(in.GetPostStartupScript())
	out.InstanceOwners = in.InstanceOwners
	if in.GetServiceAccount() != "" {
		out.ServiceAccountRef = &v1beta1.IAMServiceAccountRef{External: in.GetServiceAccount()}
	}
	out.ServiceAccountScopes = in.ServiceAccountScopes
	out.MachineType = direct.LazyPtr(in.GetMachineType())
	out.AcceleratorConfig = Instance_AcceleratorConfig_FromProto(mapCtx, in.GetAcceleratorConfig())
	out.InstallGpuDriver = direct.LazyPtr(in.GetInstallGpuDriver())
	out.CustomGpuDriverPath = direct.LazyPtr(in.GetCustomGpuDriverPath())
	out.BootDiskType = direct.Enum_FromProto(mapCtx, in.GetBootDiskType())
	out.BootDiskSizeGB = direct.LazyPtr(in.GetBootDiskSizeGb())
	out.DataDiskType = direct.Enum_FromProto(mapCtx, in.GetDataDiskType())
	out.DataDiskSizeGB = direct.LazyPtr(in.GetDataDiskSizeGb())
	out.NoRemoveDataDisk = direct.LazyPtr(in.GetNoRemoveDataDisk())
	out.DiskEncryption = direct.Enum_FromProto(mapCtx, in.GetDiskEncryption())
	if in.GetKmsKey() != "" {
		out.KMSKeyRef = &kmsv1beta1.KMSKeyRef_OneOf{External: in.GetKmsKey()}
	}
	out.ShieldedInstanceConfig = Instance_ShieldedInstanceConfig_FromProto(mapCtx, in.GetShieldedInstanceConfig())
	out.NoPublicIP = direct.LazyPtr(in.GetNoPublicIp())
	out.NoProxyAccess = direct.LazyPtr(in.GetNoProxyAccess())
	if in.GetNetwork() != "" {
		out.NetworkRef = &v1beta1.ComputeNetworkRef{External: in.GetNetwork()}
	}
	if in.GetSubnet() != "" {
		out.SubnetRef = &v1beta1.ComputeSubnetworkRef{External: in.GetSubnet()}
	}
	out.Labels = in.Labels
	out.Metadata = in.Metadata
	out.Tags = in.Tags
	out.UpgradeHistory = direct.Slice_FromProto(mapCtx, in.UpgradeHistory, Instance_UpgradeHistoryEntry_FromProto)
	out.NicType = direct.Enum_FromProto(mapCtx, in.GetNicType())
	out.ReservationAffinity = ReservationAffinity_FromProto(mapCtx, in.GetReservationAffinity())
	out.CanIPForward = direct.LazyPtr(in.GetCanIpForward())
	return out
}
func NotebookInstanceSpec_ToProto(mapCtx *direct.MapContext, in *krmv1beta1.NotebookInstanceSpec) *pb.Instance {
	if in == nil {
		return nil
	}
	out := &pb.Instance{}
	if oneof := VMImage_ToProto(mapCtx, in.VMImage); oneof != nil {
		out.Environment = &pb.Instance_VmImage{VmImage: oneof}
	}
	if oneof := ContainerImage_ToProto(mapCtx, in.ContainerImage); oneof != nil {
		out.Environment = &pb.Instance_ContainerImage{ContainerImage: oneof}
	}
	out.PostStartupScript = direct.ValueOf(in.PostStartupScript)
	out.InstanceOwners = in.InstanceOwners
	if in.ServiceAccountRef != nil {
		out.ServiceAccount = in.ServiceAccountRef.External
	}
	out.ServiceAccountScopes = in.ServiceAccountScopes
	out.MachineType = direct.ValueOf(in.MachineType)
	out.AcceleratorConfig = Instance_AcceleratorConfig_ToProto(mapCtx, in.AcceleratorConfig)
	out.InstallGpuDriver = direct.ValueOf(in.InstallGpuDriver)
	out.CustomGpuDriverPath = direct.ValueOf(in.CustomGpuDriverPath)
	out.BootDiskType = direct.Enum_ToProto[pb.Instance_DiskType](mapCtx, in.BootDiskType)
	out.BootDiskSizeGb = direct.ValueOf(in.BootDiskSizeGB)
	out.DataDiskType = direct.Enum_ToProto[pb.Instance_DiskType](mapCtx, in.DataDiskType)
	out.DataDiskSizeGb = direct.ValueOf(in.DataDiskSizeGB)
	out.NoRemoveDataDisk = direct.ValueOf(in.NoRemoveDataDisk)
	out.DiskEncryption = direct.Enum_ToProto[pb.Instance_DiskEncryption](mapCtx, in.DiskEncryption)
	if in.KMSKeyRef != nil {
		out.KmsKey = in.KMSKeyRef.External
	}
	out.ShieldedInstanceConfig = Instance_ShieldedInstanceConfig_ToProto(mapCtx, in.ShieldedInstanceConfig)
	out.NoPublicIp = direct.ValueOf(in.NoPublicIP)
	out.NoProxyAccess = direct.ValueOf(in.NoProxyAccess)
	if in.NetworkRef != nil {
		out.Network = in.NetworkRef.External
	}
	if in.SubnetRef != nil {
		out.Subnet = in.SubnetRef.External
	}
	out.Labels = in.Labels
	out.Metadata = in.Metadata
	out.Tags = in.Tags
	out.UpgradeHistory = direct.Slice_ToProto(mapCtx, in.UpgradeHistory, Instance_UpgradeHistoryEntry_ToProto)
	out.NicType = direct.Enum_ToProto[pb.Instance_NicType](mapCtx, in.NicType)
	out.ReservationAffinity = ReservationAffinity_ToProto(mapCtx, in.ReservationAffinity)
	out.CanIpForward = direct.ValueOf(in.CanIPForward)
	return out
}
func ReservationAffinity_FromProto(mapCtx *direct.MapContext, in *pb.ReservationAffinity) *krmv1beta1.ReservationAffinity {
	if in == nil {
		return nil
	}
	out := &krmv1beta1.ReservationAffinity{}
	out.ConsumeReservationType = direct.Enum_FromProto(mapCtx, in.GetConsumeReservationType())
	out.Key = direct.LazyPtr(in.GetKey())
	out.Values = in.Values
	return out
}
func ReservationAffinity_ToProto(mapCtx *direct.MapContext, in *krmv1beta1.ReservationAffinity) *pb.ReservationAffinity {
	if in == nil {
		return nil
	}
	out := &pb.ReservationAffinity{}
	out.ConsumeReservationType = direct.Enum_ToProto[pb.ReservationAffinity_Type](mapCtx, in.ConsumeReservationType)
	out.Key = direct.ValueOf(in.Key)
	out.Values = in.Values
	return out
}
func VMImage_FromProto(mapCtx *direct.MapContext, in *pb.VmImage) *krmv1beta1.VMImage {
	if in == nil {
		return nil
	}
	out := &krmv1beta1.VMImage{}
	out.Project = direct.LazyPtr(in.GetProject())
	out.ImageName = direct.LazyPtr(in.GetImageName())
	out.ImageFamily = direct.LazyPtr(in.GetImageFamily())
	return out
}
func VMImage_ToProto(mapCtx *direct.MapContext, in *krmv1beta1.VMImage) *pb.VmImage {
	if in == nil {
		return nil
	}
	out := &pb.VmImage{}
	out.Project = direct.ValueOf(in.Project)
	if oneof := VMImage_ImageName_ToProto(mapCtx, in.ImageName); oneof != nil {
		out.Image = oneof
	}
	if oneof := VMImage_ImageFamily_ToProto(mapCtx, in.ImageFamily); oneof != nil {
		out.Image = oneof
	}
	return out
}
