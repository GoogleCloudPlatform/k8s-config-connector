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
	pb "cloud.google.com/go/notebooks/apiv2/notebookspb"
	krm "github.com/GoogleCloudPlatform/k8s-config-connector/apis/notebooks/v1alpha1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct"
	refs "github.com/GoogleCloudPlatform/k8s-config-connector/apis/refs/v1beta1"
)
func AcceleratorConfig_FromProto(mapCtx *direct.MapContext, in *pb.AcceleratorConfig) *krm.AcceleratorConfig {
	if in == nil {
		return nil
	}
	out := &krm.AcceleratorConfig{}
	out.Type = direct.Enum_FromProto(mapCtx, in.GetType())
	out.CoreCount = direct.LazyPtr(in.GetCoreCount())
	return out
}
func AcceleratorConfig_ToProto(mapCtx *direct.MapContext, in *krm.AcceleratorConfig) *pb.AcceleratorConfig {
	if in == nil {
		return nil
	}
	out := &pb.AcceleratorConfig{}
	out.Type = direct.Enum_ToProto[pb.AcceleratorConfig_AcceleratorType](mapCtx, in.Type)
	out.CoreCount = direct.ValueOf(in.CoreCount)
	return out
}
func BootDisk_FromProto(mapCtx *direct.MapContext, in *pb.BootDisk) *krm.BootDisk {
	if in == nil {
		return nil
	}
	out := &krm.BootDisk{}
	out.DiskSizeGB = direct.LazyPtr(in.GetDiskSizeGb())
	out.DiskType = direct.Enum_FromProto(mapCtx, in.GetDiskType())
	out.DiskEncryption = direct.Enum_FromProto(mapCtx, in.GetDiskEncryption())
	out.KMSKey = direct.LazyPtr(in.GetKmsKey())
	return out
}
func BootDisk_ToProto(mapCtx *direct.MapContext, in *krm.BootDisk) *pb.BootDisk {
	if in == nil {
		return nil
	}
	out := &pb.BootDisk{}
	out.DiskSizeGb = direct.ValueOf(in.DiskSizeGB)
	out.DiskType = direct.Enum_ToProto[pb.DiskType](mapCtx, in.DiskType)
	out.DiskEncryption = direct.Enum_ToProto[pb.DiskEncryption](mapCtx, in.DiskEncryption)
	out.KmsKey = direct.ValueOf(in.KMSKey)
	return out
}
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
func DataDisk_FromProto(mapCtx *direct.MapContext, in *pb.DataDisk) *krm.DataDisk {
	if in == nil {
		return nil
	}
	out := &krm.DataDisk{}
	out.DiskSizeGB = direct.LazyPtr(in.GetDiskSizeGb())
	out.DiskType = direct.Enum_FromProto(mapCtx, in.GetDiskType())
	out.DiskEncryption = direct.Enum_FromProto(mapCtx, in.GetDiskEncryption())
	out.KMSKey = direct.LazyPtr(in.GetKmsKey())
	return out
}
func DataDisk_ToProto(mapCtx *direct.MapContext, in *krm.DataDisk) *pb.DataDisk {
	if in == nil {
		return nil
	}
	out := &pb.DataDisk{}
	out.DiskSizeGb = direct.ValueOf(in.DiskSizeGB)
	out.DiskType = direct.Enum_ToProto[pb.DiskType](mapCtx, in.DiskType)
	out.DiskEncryption = direct.Enum_ToProto[pb.DiskEncryption](mapCtx, in.DiskEncryption)
	out.KmsKey = direct.ValueOf(in.KMSKey)
	return out
}
func GPUDriverConfig_FromProto(mapCtx *direct.MapContext, in *pb.GPUDriverConfig) *krm.GPUDriverConfig {
	if in == nil {
		return nil
	}
	out := &krm.GPUDriverConfig{}
	out.EnableGpuDriver = direct.LazyPtr(in.GetEnableGpuDriver())
	out.CustomGpuDriverPath = direct.LazyPtr(in.GetCustomGpuDriverPath())
	return out
}
func GPUDriverConfig_ToProto(mapCtx *direct.MapContext, in *krm.GPUDriverConfig) *pb.GPUDriverConfig {
	if in == nil {
		return nil
	}
	out := &pb.GPUDriverConfig{}
	out.EnableGpuDriver = direct.ValueOf(in.EnableGpuDriver)
	out.CustomGpuDriverPath = direct.ValueOf(in.CustomGpuDriverPath)
	return out
}
func GceSetup_FromProto(mapCtx *direct.MapContext, in *pb.GceSetup) *krm.GceSetup {
	if in == nil {
		return nil
	}
	out := &krm.GceSetup{}
	out.MachineType = direct.LazyPtr(in.GetMachineType())
	out.AcceleratorConfigs = direct.Slice_FromProto(mapCtx, in.AcceleratorConfigs, AcceleratorConfig_FromProto)
	out.ServiceAccounts = direct.Slice_FromProto(mapCtx, in.ServiceAccounts, ServiceAccount_FromProto)
	out.VmImage = VmImage_FromProto(mapCtx, in.GetVmImage())
	out.ContainerImage = ContainerImage_FromProto(mapCtx, in.GetContainerImage())
	out.BootDisk = BootDisk_FromProto(mapCtx, in.GetBootDisk())
	out.DataDisks = direct.Slice_FromProto(mapCtx, in.DataDisks, DataDisk_FromProto)
	out.ShieldedInstanceConfig = ShieldedInstanceConfig_FromProto(mapCtx, in.GetShieldedInstanceConfig())
	out.NetworkInterfaces = direct.Slice_FromProto(mapCtx, in.NetworkInterfaces, NetworkInterface_FromProto)
	out.DisablePublicIP = direct.LazyPtr(in.GetDisablePublicIp())
	out.Tags = in.Tags
	out.Metadata = in.Metadata
	out.EnableIPForwarding = direct.LazyPtr(in.GetEnableIpForwarding())
	out.GpuDriverConfig = GPUDriverConfig_FromProto(mapCtx, in.GetGpuDriverConfig())
	return out
}
func GceSetup_ToProto(mapCtx *direct.MapContext, in *krm.GceSetup) *pb.GceSetup {
	if in == nil {
		return nil
	}
	out := &pb.GceSetup{}
	out.MachineType = direct.ValueOf(in.MachineType)
	out.AcceleratorConfigs = direct.Slice_ToProto(mapCtx, in.AcceleratorConfigs, AcceleratorConfig_ToProto)
	out.ServiceAccounts = direct.Slice_ToProto(mapCtx, in.ServiceAccounts, ServiceAccount_ToProto)
	if oneof := VmImage_ToProto(mapCtx, in.VmImage); oneof != nil {
		out.Image = &pb.GceSetup_VmImage{VmImage: oneof}
	}
	if oneof := ContainerImage_ToProto(mapCtx, in.ContainerImage); oneof != nil {
		out.Image = &pb.GceSetup_ContainerImage{ContainerImage: oneof}
	}
	out.BootDisk = BootDisk_ToProto(mapCtx, in.BootDisk)
	out.DataDisks = direct.Slice_ToProto(mapCtx, in.DataDisks, DataDisk_ToProto)
	out.ShieldedInstanceConfig = ShieldedInstanceConfig_ToProto(mapCtx, in.ShieldedInstanceConfig)
	out.NetworkInterfaces = direct.Slice_ToProto(mapCtx, in.NetworkInterfaces, NetworkInterface_ToProto)
	out.DisablePublicIp = direct.ValueOf(in.DisablePublicIP)
	out.Tags = in.Tags
	out.Metadata = in.Metadata
	out.EnableIpForwarding = direct.ValueOf(in.EnableIPForwarding)
	out.GpuDriverConfig = GPUDriverConfig_ToProto(mapCtx, in.GpuDriverConfig)
	return out
}
func GceSetupObservedState_FromProto(mapCtx *direct.MapContext, in *pb.GceSetup) *krm.GceSetupObservedState {
	if in == nil {
		return nil
	}
	out := &krm.GceSetupObservedState{}
	// MISSING: MachineType
	// MISSING: AcceleratorConfigs
	out.ServiceAccounts = direct.Slice_FromProto(mapCtx, in.ServiceAccounts, ServiceAccountObservedState_FromProto)
	// MISSING: VmImage
	// MISSING: ContainerImage
	// MISSING: BootDisk
	// MISSING: DataDisks
	// MISSING: ShieldedInstanceConfig
	// MISSING: NetworkInterfaces
	// MISSING: DisablePublicIP
	// MISSING: Tags
	// MISSING: Metadata
	// MISSING: EnableIPForwarding
	// MISSING: GpuDriverConfig
	return out
}
func GceSetupObservedState_ToProto(mapCtx *direct.MapContext, in *krm.GceSetupObservedState) *pb.GceSetup {
	if in == nil {
		return nil
	}
	out := &pb.GceSetup{}
	// MISSING: MachineType
	// MISSING: AcceleratorConfigs
	out.ServiceAccounts = direct.Slice_ToProto(mapCtx, in.ServiceAccounts, ServiceAccountObservedState_ToProto)
	// MISSING: VmImage
	// MISSING: ContainerImage
	// MISSING: BootDisk
	// MISSING: DataDisks
	// MISSING: ShieldedInstanceConfig
	// MISSING: NetworkInterfaces
	// MISSING: DisablePublicIP
	// MISSING: Tags
	// MISSING: Metadata
	// MISSING: EnableIPForwarding
	// MISSING: GpuDriverConfig
	return out
}
func Instance_FromProto(mapCtx *direct.MapContext, in *pb.Instance) *krm.Instance {
	if in == nil {
		return nil
	}
	out := &krm.Instance{}
	// MISSING: Name
	out.GCESetup = GceSetup_FromProto(mapCtx, in.GetGceSetup())
	// MISSING: ProxyURI
	out.InstanceOwners = in.InstanceOwners
	// MISSING: Creator
	// MISSING: State
	// MISSING: UpgradeHistory
	// MISSING: ID
	// MISSING: HealthState
	// MISSING: HealthInfo
	// MISSING: CreateTime
	// MISSING: UpdateTime
	out.DisableProxyAccess = direct.LazyPtr(in.GetDisableProxyAccess())
	out.Labels = in.Labels
	return out
}
func Instance_ToProto(mapCtx *direct.MapContext, in *krm.Instance) *pb.Instance {
	if in == nil {
		return nil
	}
	out := &pb.Instance{}
	// MISSING: Name
	if oneof := GceSetup_ToProto(mapCtx, in.GCESetup); oneof != nil {
		out.Infrastructure = &pb.Instance_GceSetup{GceSetup: oneof}
	}
	// MISSING: ProxyURI
	out.InstanceOwners = in.InstanceOwners
	// MISSING: Creator
	// MISSING: State
	// MISSING: UpgradeHistory
	// MISSING: ID
	// MISSING: HealthState
	// MISSING: HealthInfo
	// MISSING: CreateTime
	// MISSING: UpdateTime
	out.DisableProxyAccess = direct.ValueOf(in.DisableProxyAccess)
	out.Labels = in.Labels
	return out
}
func InstanceObservedState_FromProto(mapCtx *direct.MapContext, in *pb.Instance) *krm.InstanceObservedState {
	if in == nil {
		return nil
	}
	out := &krm.InstanceObservedState{}
	out.Name = direct.LazyPtr(in.GetName())
	out.GCESetup = GceSetupObservedState_FromProto(mapCtx, in.GetGceSetup())
	out.ProxyURI = direct.LazyPtr(in.GetProxyUri())
	// MISSING: InstanceOwners
	out.Creator = direct.LazyPtr(in.GetCreator())
	out.State = direct.Enum_FromProto(mapCtx, in.GetState())
	out.UpgradeHistory = direct.Slice_FromProto(mapCtx, in.UpgradeHistory, UpgradeHistoryEntry_FromProto)
	out.ID = direct.LazyPtr(in.GetId())
	out.HealthState = direct.Enum_FromProto(mapCtx, in.GetHealthState())
	out.HealthInfo = in.HealthInfo
	out.CreateTime = direct.StringTimestamp_FromProto(mapCtx, in.GetCreateTime())
	out.UpdateTime = direct.StringTimestamp_FromProto(mapCtx, in.GetUpdateTime())
	// MISSING: DisableProxyAccess
	// MISSING: Labels
	return out
}
func InstanceObservedState_ToProto(mapCtx *direct.MapContext, in *krm.InstanceObservedState) *pb.Instance {
	if in == nil {
		return nil
	}
	out := &pb.Instance{}
	out.Name = direct.ValueOf(in.Name)
	if oneof := GceSetupObservedState_ToProto(mapCtx, in.GCESetup); oneof != nil {
		out.Infrastructure = &pb.Instance_GceSetup{GceSetup: oneof}
	}
	out.ProxyUri = direct.ValueOf(in.ProxyURI)
	// MISSING: InstanceOwners
	out.Creator = direct.ValueOf(in.Creator)
	out.State = direct.Enum_ToProto[pb.State](mapCtx, in.State)
	out.UpgradeHistory = direct.Slice_ToProto(mapCtx, in.UpgradeHistory, UpgradeHistoryEntry_ToProto)
	out.Id = direct.ValueOf(in.ID)
	out.HealthState = direct.Enum_ToProto[pb.HealthState](mapCtx, in.HealthState)
	out.HealthInfo = in.HealthInfo
	out.CreateTime = direct.StringTimestamp_ToProto(mapCtx, in.CreateTime)
	out.UpdateTime = direct.StringTimestamp_ToProto(mapCtx, in.UpdateTime)
	// MISSING: DisableProxyAccess
	// MISSING: Labels
	return out
}
func NetworkInterface_FromProto(mapCtx *direct.MapContext, in *pb.NetworkInterface) *krm.NetworkInterface {
	if in == nil {
		return nil
	}
	out := &krm.NetworkInterface{}
	out.Network = direct.LazyPtr(in.GetNetwork())
	out.Subnet = direct.LazyPtr(in.GetSubnet())
	out.NicType = direct.Enum_FromProto(mapCtx, in.GetNicType())
	return out
}
func NetworkInterface_ToProto(mapCtx *direct.MapContext, in *krm.NetworkInterface) *pb.NetworkInterface {
	if in == nil {
		return nil
	}
	out := &pb.NetworkInterface{}
	out.Network = direct.ValueOf(in.Network)
	out.Subnet = direct.ValueOf(in.Subnet)
	out.NicType = direct.Enum_ToProto[pb.NetworkInterface_NicType](mapCtx, in.NicType)
	return out
}
func ServiceAccount_FromProto(mapCtx *direct.MapContext, in *pb.ServiceAccount) *krm.ServiceAccount {
	if in == nil {
		return nil
	}
	out := &krm.ServiceAccount{}
	out.Email = direct.LazyPtr(in.GetEmail())
	// MISSING: Scopes
	return out
}
func ServiceAccount_ToProto(mapCtx *direct.MapContext, in *krm.ServiceAccount) *pb.ServiceAccount {
	if in == nil {
		return nil
	}
	out := &pb.ServiceAccount{}
	out.Email = direct.ValueOf(in.Email)
	// MISSING: Scopes
	return out
}
func ServiceAccountObservedState_FromProto(mapCtx *direct.MapContext, in *pb.ServiceAccount) *krm.ServiceAccountObservedState {
	if in == nil {
		return nil
	}
	out := &krm.ServiceAccountObservedState{}
	// MISSING: Email
	out.Scopes = in.Scopes
	return out
}
func ServiceAccountObservedState_ToProto(mapCtx *direct.MapContext, in *krm.ServiceAccountObservedState) *pb.ServiceAccount {
	if in == nil {
		return nil
	}
	out := &pb.ServiceAccount{}
	// MISSING: Email
	out.Scopes = in.Scopes
	return out
}
func ShieldedInstanceConfig_FromProto(mapCtx *direct.MapContext, in *pb.ShieldedInstanceConfig) *krm.ShieldedInstanceConfig {
	if in == nil {
		return nil
	}
	out := &krm.ShieldedInstanceConfig{}
	out.EnableSecureBoot = direct.LazyPtr(in.GetEnableSecureBoot())
	out.EnableVTPM = direct.LazyPtr(in.GetEnableVtpm())
	out.EnableIntegrityMonitoring = direct.LazyPtr(in.GetEnableIntegrityMonitoring())
	return out
}
func ShieldedInstanceConfig_ToProto(mapCtx *direct.MapContext, in *krm.ShieldedInstanceConfig) *pb.ShieldedInstanceConfig {
	if in == nil {
		return nil
	}
	out := &pb.ShieldedInstanceConfig{}
	out.EnableSecureBoot = direct.ValueOf(in.EnableSecureBoot)
	out.EnableVtpm = direct.ValueOf(in.EnableVTPM)
	out.EnableIntegrityMonitoring = direct.ValueOf(in.EnableIntegrityMonitoring)
	return out
}
func UpgradeHistoryEntry_FromProto(mapCtx *direct.MapContext, in *pb.UpgradeHistoryEntry) *krm.UpgradeHistoryEntry {
	if in == nil {
		return nil
	}
	out := &krm.UpgradeHistoryEntry{}
	out.Snapshot = direct.LazyPtr(in.GetSnapshot())
	out.VmImage = direct.LazyPtr(in.GetVmImage())
	out.ContainerImage = direct.LazyPtr(in.GetContainerImage())
	out.Framework = direct.LazyPtr(in.GetFramework())
	out.Version = direct.LazyPtr(in.GetVersion())
	// MISSING: State
	out.CreateTime = direct.StringTimestamp_FromProto(mapCtx, in.GetCreateTime())
	out.Action = direct.Enum_FromProto(mapCtx, in.GetAction())
	out.TargetVersion = direct.LazyPtr(in.GetTargetVersion())
	return out
}
func UpgradeHistoryEntry_ToProto(mapCtx *direct.MapContext, in *krm.UpgradeHistoryEntry) *pb.UpgradeHistoryEntry {
	if in == nil {
		return nil
	}
	out := &pb.UpgradeHistoryEntry{}
	out.Snapshot = direct.ValueOf(in.Snapshot)
	out.VmImage = direct.ValueOf(in.VmImage)
	out.ContainerImage = direct.ValueOf(in.ContainerImage)
	out.Framework = direct.ValueOf(in.Framework)
	out.Version = direct.ValueOf(in.Version)
	// MISSING: State
	out.CreateTime = direct.StringTimestamp_ToProto(mapCtx, in.CreateTime)
	out.Action = direct.Enum_ToProto[pb.UpgradeHistoryEntry_Action](mapCtx, in.Action)
	out.TargetVersion = direct.ValueOf(in.TargetVersion)
	return out
}
func UpgradeHistoryEntryObservedState_FromProto(mapCtx *direct.MapContext, in *pb.UpgradeHistoryEntry) *krm.UpgradeHistoryEntryObservedState {
	if in == nil {
		return nil
	}
	out := &krm.UpgradeHistoryEntryObservedState{}
	// MISSING: Snapshot
	// MISSING: VmImage
	// MISSING: ContainerImage
	// MISSING: Framework
	// MISSING: Version
	out.State = direct.Enum_FromProto(mapCtx, in.GetState())
	// MISSING: CreateTime
	// MISSING: Action
	// MISSING: TargetVersion
	return out
}
func UpgradeHistoryEntryObservedState_ToProto(mapCtx *direct.MapContext, in *krm.UpgradeHistoryEntryObservedState) *pb.UpgradeHistoryEntry {
	if in == nil {
		return nil
	}
	out := &pb.UpgradeHistoryEntry{}
	// MISSING: Snapshot
	// MISSING: VmImage
	// MISSING: ContainerImage
	// MISSING: Framework
	// MISSING: Version
	out.State = direct.Enum_ToProto[pb.UpgradeHistoryEntry_State](mapCtx, in.State)
	// MISSING: CreateTime
	// MISSING: Action
	// MISSING: TargetVersion
	return out
}
func VmImage_FromProto(mapCtx *direct.MapContext, in *pb.VmImage) *krm.VmImage {
	if in == nil {
		return nil
	}
	out := &krm.VmImage{}
	out.Project = direct.LazyPtr(in.GetProject())
	out.Name = direct.LazyPtr(in.GetName())
	out.Family = direct.LazyPtr(in.GetFamily())
	return out
}
func VmImage_ToProto(mapCtx *direct.MapContext, in *krm.VmImage) *pb.VmImage {
	if in == nil {
		return nil
	}
	out := &pb.VmImage{}
	out.Project = direct.ValueOf(in.Project)
	if oneof := VmImage_Name_ToProto(mapCtx, in.Name); oneof != nil {
		out.Image = oneof
	}
	if oneof := VmImage_Family_ToProto(mapCtx, in.Family); oneof != nil {
		out.Image = oneof
	}
	return out
}
