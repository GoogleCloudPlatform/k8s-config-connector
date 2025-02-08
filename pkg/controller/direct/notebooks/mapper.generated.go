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
	pb "cloud.google.com/go/notebooks/apiv1/notebookspb"
	krm "github.com/GoogleCloudPlatform/k8s-config-connector/apis/notebooks/v1alpha1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct"
	refs "github.com/GoogleCloudPlatform/k8s-config-connector/apis/refs/v1beta1"
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
func EncryptionConfig_FromProto(mapCtx *direct.MapContext, in *pb.EncryptionConfig) *krm.EncryptionConfig {
	if in == nil {
		return nil
	}
	out := &krm.EncryptionConfig{}
	out.KMSKey = direct.LazyPtr(in.GetKmsKey())
	return out
}
func EncryptionConfig_ToProto(mapCtx *direct.MapContext, in *krm.EncryptionConfig) *pb.EncryptionConfig {
	if in == nil {
		return nil
	}
	out := &pb.EncryptionConfig{}
	out.KmsKey = direct.ValueOf(in.KMSKey)
	return out
}
func LocalDisk_FromProto(mapCtx *direct.MapContext, in *pb.LocalDisk) *krm.LocalDisk {
	if in == nil {
		return nil
	}
	out := &krm.LocalDisk{}
	// MISSING: AutoDelete
	// MISSING: Boot
	// MISSING: DeviceName
	// MISSING: GuestOsFeatures
	// MISSING: Index
	out.InitializeParams = LocalDiskInitializeParams_FromProto(mapCtx, in.GetInitializeParams())
	out.Interface = direct.LazyPtr(in.GetInterface())
	// MISSING: Kind
	// MISSING: Licenses
	out.Mode = direct.LazyPtr(in.GetMode())
	out.Source = direct.LazyPtr(in.GetSource())
	out.Type = direct.LazyPtr(in.GetType())
	return out
}
func LocalDisk_ToProto(mapCtx *direct.MapContext, in *krm.LocalDisk) *pb.LocalDisk {
	if in == nil {
		return nil
	}
	out := &pb.LocalDisk{}
	// MISSING: AutoDelete
	// MISSING: Boot
	// MISSING: DeviceName
	// MISSING: GuestOsFeatures
	// MISSING: Index
	out.InitializeParams = LocalDiskInitializeParams_ToProto(mapCtx, in.InitializeParams)
	out.Interface = direct.ValueOf(in.Interface)
	// MISSING: Kind
	// MISSING: Licenses
	out.Mode = direct.ValueOf(in.Mode)
	out.Source = direct.ValueOf(in.Source)
	out.Type = direct.ValueOf(in.Type)
	return out
}
func LocalDiskInitializeParams_FromProto(mapCtx *direct.MapContext, in *pb.LocalDiskInitializeParams) *krm.LocalDiskInitializeParams {
	if in == nil {
		return nil
	}
	out := &krm.LocalDiskInitializeParams{}
	out.Description = direct.LazyPtr(in.GetDescription())
	out.DiskName = direct.LazyPtr(in.GetDiskName())
	out.DiskSizeGB = direct.LazyPtr(in.GetDiskSizeGb())
	out.DiskType = direct.Enum_FromProto(mapCtx, in.GetDiskType())
	out.Labels = in.Labels
	return out
}
func LocalDiskInitializeParams_ToProto(mapCtx *direct.MapContext, in *krm.LocalDiskInitializeParams) *pb.LocalDiskInitializeParams {
	if in == nil {
		return nil
	}
	out := &pb.LocalDiskInitializeParams{}
	out.Description = direct.ValueOf(in.Description)
	out.DiskName = direct.ValueOf(in.DiskName)
	out.DiskSizeGb = direct.ValueOf(in.DiskSizeGB)
	out.DiskType = direct.Enum_ToProto[pb.LocalDiskInitializeParams_DiskType](mapCtx, in.DiskType)
	out.Labels = in.Labels
	return out
}
func LocalDiskObservedState_FromProto(mapCtx *direct.MapContext, in *pb.LocalDisk) *krm.LocalDiskObservedState {
	if in == nil {
		return nil
	}
	out := &krm.LocalDiskObservedState{}
	out.AutoDelete = direct.LazyPtr(in.GetAutoDelete())
	out.Boot = direct.LazyPtr(in.GetBoot())
	out.DeviceName = direct.LazyPtr(in.GetDeviceName())
	out.GuestOsFeatures = direct.Slice_FromProto(mapCtx, in.GuestOsFeatures, LocalDisk_RuntimeGuestOsFeature_FromProto)
	out.Index = direct.LazyPtr(in.GetIndex())
	// MISSING: InitializeParams
	// MISSING: Interface
	out.Kind = direct.LazyPtr(in.GetKind())
	out.Licenses = in.Licenses
	// MISSING: Mode
	// MISSING: Source
	// MISSING: Type
	return out
}
func LocalDiskObservedState_ToProto(mapCtx *direct.MapContext, in *krm.LocalDiskObservedState) *pb.LocalDisk {
	if in == nil {
		return nil
	}
	out := &pb.LocalDisk{}
	out.AutoDelete = direct.ValueOf(in.AutoDelete)
	out.Boot = direct.ValueOf(in.Boot)
	out.DeviceName = direct.ValueOf(in.DeviceName)
	out.GuestOsFeatures = direct.Slice_ToProto(mapCtx, in.GuestOsFeatures, LocalDisk_RuntimeGuestOsFeature_ToProto)
	out.Index = direct.ValueOf(in.Index)
	// MISSING: InitializeParams
	// MISSING: Interface
	out.Kind = direct.ValueOf(in.Kind)
	out.Licenses = in.Licenses
	// MISSING: Mode
	// MISSING: Source
	// MISSING: Type
	return out
}
func LocalDisk_RuntimeGuestOsFeature_FromProto(mapCtx *direct.MapContext, in *pb.LocalDisk_RuntimeGuestOsFeature) *krm.LocalDisk_RuntimeGuestOsFeature {
	if in == nil {
		return nil
	}
	out := &krm.LocalDisk_RuntimeGuestOsFeature{}
	out.Type = direct.LazyPtr(in.GetType())
	return out
}
func LocalDisk_RuntimeGuestOsFeature_ToProto(mapCtx *direct.MapContext, in *krm.LocalDisk_RuntimeGuestOsFeature) *pb.LocalDisk_RuntimeGuestOsFeature {
	if in == nil {
		return nil
	}
	out := &pb.LocalDisk_RuntimeGuestOsFeature{}
	out.Type = direct.ValueOf(in.Type)
	return out
}
func NotebooksRuntimeObservedState_FromProto(mapCtx *direct.MapContext, in *pb.Runtime) *krm.NotebooksRuntimeObservedState {
	if in == nil {
		return nil
	}
	out := &krm.NotebooksRuntimeObservedState{}
	// MISSING: Name
	// MISSING: VirtualMachine
	// MISSING: State
	// MISSING: HealthState
	// MISSING: AccessConfig
	// MISSING: SoftwareConfig
	// MISSING: Metrics
	// MISSING: CreateTime
	// MISSING: UpdateTime
	return out
}
func NotebooksRuntimeObservedState_ToProto(mapCtx *direct.MapContext, in *krm.NotebooksRuntimeObservedState) *pb.Runtime {
	if in == nil {
		return nil
	}
	out := &pb.Runtime{}
	// MISSING: Name
	// MISSING: VirtualMachine
	// MISSING: State
	// MISSING: HealthState
	// MISSING: AccessConfig
	// MISSING: SoftwareConfig
	// MISSING: Metrics
	// MISSING: CreateTime
	// MISSING: UpdateTime
	return out
}
func NotebooksRuntimeSpec_FromProto(mapCtx *direct.MapContext, in *pb.Runtime) *krm.NotebooksRuntimeSpec {
	if in == nil {
		return nil
	}
	out := &krm.NotebooksRuntimeSpec{}
	// MISSING: Name
	// MISSING: VirtualMachine
	// MISSING: State
	// MISSING: HealthState
	// MISSING: AccessConfig
	// MISSING: SoftwareConfig
	// MISSING: Metrics
	// MISSING: CreateTime
	// MISSING: UpdateTime
	return out
}
func NotebooksRuntimeSpec_ToProto(mapCtx *direct.MapContext, in *krm.NotebooksRuntimeSpec) *pb.Runtime {
	if in == nil {
		return nil
	}
	out := &pb.Runtime{}
	// MISSING: Name
	// MISSING: VirtualMachine
	// MISSING: State
	// MISSING: HealthState
	// MISSING: AccessConfig
	// MISSING: SoftwareConfig
	// MISSING: Metrics
	// MISSING: CreateTime
	// MISSING: UpdateTime
	return out
}
func Runtime_FromProto(mapCtx *direct.MapContext, in *pb.Runtime) *krm.Runtime {
	if in == nil {
		return nil
	}
	out := &krm.Runtime{}
	// MISSING: Name
	out.VirtualMachine = VirtualMachine_FromProto(mapCtx, in.GetVirtualMachine())
	// MISSING: State
	// MISSING: HealthState
	out.AccessConfig = RuntimeAccessConfig_FromProto(mapCtx, in.GetAccessConfig())
	out.SoftwareConfig = RuntimeSoftwareConfig_FromProto(mapCtx, in.GetSoftwareConfig())
	// MISSING: Metrics
	// MISSING: CreateTime
	// MISSING: UpdateTime
	return out
}
func Runtime_ToProto(mapCtx *direct.MapContext, in *krm.Runtime) *pb.Runtime {
	if in == nil {
		return nil
	}
	out := &pb.Runtime{}
	// MISSING: Name
	if oneof := VirtualMachine_ToProto(mapCtx, in.VirtualMachine); oneof != nil {
		out.RuntimeType = &pb.Runtime_VirtualMachine{VirtualMachine: oneof}
	}
	// MISSING: State
	// MISSING: HealthState
	out.AccessConfig = RuntimeAccessConfig_ToProto(mapCtx, in.AccessConfig)
	out.SoftwareConfig = RuntimeSoftwareConfig_ToProto(mapCtx, in.SoftwareConfig)
	// MISSING: Metrics
	// MISSING: CreateTime
	// MISSING: UpdateTime
	return out
}
func RuntimeAcceleratorConfig_FromProto(mapCtx *direct.MapContext, in *pb.RuntimeAcceleratorConfig) *krm.RuntimeAcceleratorConfig {
	if in == nil {
		return nil
	}
	out := &krm.RuntimeAcceleratorConfig{}
	out.Type = direct.Enum_FromProto(mapCtx, in.GetType())
	out.CoreCount = direct.LazyPtr(in.GetCoreCount())
	return out
}
func RuntimeAcceleratorConfig_ToProto(mapCtx *direct.MapContext, in *krm.RuntimeAcceleratorConfig) *pb.RuntimeAcceleratorConfig {
	if in == nil {
		return nil
	}
	out := &pb.RuntimeAcceleratorConfig{}
	out.Type = direct.Enum_ToProto[pb.RuntimeAcceleratorConfig_AcceleratorType](mapCtx, in.Type)
	out.CoreCount = direct.ValueOf(in.CoreCount)
	return out
}
func RuntimeAccessConfig_FromProto(mapCtx *direct.MapContext, in *pb.RuntimeAccessConfig) *krm.RuntimeAccessConfig {
	if in == nil {
		return nil
	}
	out := &krm.RuntimeAccessConfig{}
	out.AccessType = direct.Enum_FromProto(mapCtx, in.GetAccessType())
	out.RuntimeOwner = direct.LazyPtr(in.GetRuntimeOwner())
	// MISSING: ProxyURI
	return out
}
func RuntimeAccessConfig_ToProto(mapCtx *direct.MapContext, in *krm.RuntimeAccessConfig) *pb.RuntimeAccessConfig {
	if in == nil {
		return nil
	}
	out := &pb.RuntimeAccessConfig{}
	out.AccessType = direct.Enum_ToProto[pb.RuntimeAccessConfig_RuntimeAccessType](mapCtx, in.AccessType)
	out.RuntimeOwner = direct.ValueOf(in.RuntimeOwner)
	// MISSING: ProxyURI
	return out
}
func RuntimeAccessConfigObservedState_FromProto(mapCtx *direct.MapContext, in *pb.RuntimeAccessConfig) *krm.RuntimeAccessConfigObservedState {
	if in == nil {
		return nil
	}
	out := &krm.RuntimeAccessConfigObservedState{}
	// MISSING: AccessType
	// MISSING: RuntimeOwner
	out.ProxyURI = direct.LazyPtr(in.GetProxyUri())
	return out
}
func RuntimeAccessConfigObservedState_ToProto(mapCtx *direct.MapContext, in *krm.RuntimeAccessConfigObservedState) *pb.RuntimeAccessConfig {
	if in == nil {
		return nil
	}
	out := &pb.RuntimeAccessConfig{}
	// MISSING: AccessType
	// MISSING: RuntimeOwner
	out.ProxyUri = direct.ValueOf(in.ProxyURI)
	return out
}
func RuntimeMetrics_FromProto(mapCtx *direct.MapContext, in *pb.RuntimeMetrics) *krm.RuntimeMetrics {
	if in == nil {
		return nil
	}
	out := &krm.RuntimeMetrics{}
	// MISSING: SystemMetrics
	return out
}
func RuntimeMetrics_ToProto(mapCtx *direct.MapContext, in *krm.RuntimeMetrics) *pb.RuntimeMetrics {
	if in == nil {
		return nil
	}
	out := &pb.RuntimeMetrics{}
	// MISSING: SystemMetrics
	return out
}
func RuntimeMetricsObservedState_FromProto(mapCtx *direct.MapContext, in *pb.RuntimeMetrics) *krm.RuntimeMetricsObservedState {
	if in == nil {
		return nil
	}
	out := &krm.RuntimeMetricsObservedState{}
	out.SystemMetrics = in.SystemMetrics
	return out
}
func RuntimeMetricsObservedState_ToProto(mapCtx *direct.MapContext, in *krm.RuntimeMetricsObservedState) *pb.RuntimeMetrics {
	if in == nil {
		return nil
	}
	out := &pb.RuntimeMetrics{}
	out.SystemMetrics = in.SystemMetrics
	return out
}
func RuntimeObservedState_FromProto(mapCtx *direct.MapContext, in *pb.Runtime) *krm.RuntimeObservedState {
	if in == nil {
		return nil
	}
	out := &krm.RuntimeObservedState{}
	out.Name = direct.LazyPtr(in.GetName())
	out.VirtualMachine = VirtualMachineObservedState_FromProto(mapCtx, in.GetVirtualMachine())
	out.State = direct.Enum_FromProto(mapCtx, in.GetState())
	out.HealthState = direct.Enum_FromProto(mapCtx, in.GetHealthState())
	out.AccessConfig = RuntimeAccessConfigObservedState_FromProto(mapCtx, in.GetAccessConfig())
	out.SoftwareConfig = RuntimeSoftwareConfigObservedState_FromProto(mapCtx, in.GetSoftwareConfig())
	out.Metrics = RuntimeMetrics_FromProto(mapCtx, in.GetMetrics())
	out.CreateTime = direct.StringTimestamp_FromProto(mapCtx, in.GetCreateTime())
	out.UpdateTime = direct.StringTimestamp_FromProto(mapCtx, in.GetUpdateTime())
	return out
}
func RuntimeObservedState_ToProto(mapCtx *direct.MapContext, in *krm.RuntimeObservedState) *pb.Runtime {
	if in == nil {
		return nil
	}
	out := &pb.Runtime{}
	out.Name = direct.ValueOf(in.Name)
	if oneof := VirtualMachineObservedState_ToProto(mapCtx, in.VirtualMachine); oneof != nil {
		out.RuntimeType = &pb.Runtime_VirtualMachine{VirtualMachine: oneof}
	}
	out.State = direct.Enum_ToProto[pb.Runtime_State](mapCtx, in.State)
	out.HealthState = direct.Enum_ToProto[pb.Runtime_HealthState](mapCtx, in.HealthState)
	out.AccessConfig = RuntimeAccessConfigObservedState_ToProto(mapCtx, in.AccessConfig)
	out.SoftwareConfig = RuntimeSoftwareConfigObservedState_ToProto(mapCtx, in.SoftwareConfig)
	out.Metrics = RuntimeMetrics_ToProto(mapCtx, in.Metrics)
	out.CreateTime = direct.StringTimestamp_ToProto(mapCtx, in.CreateTime)
	out.UpdateTime = direct.StringTimestamp_ToProto(mapCtx, in.UpdateTime)
	return out
}
func RuntimeShieldedInstanceConfig_FromProto(mapCtx *direct.MapContext, in *pb.RuntimeShieldedInstanceConfig) *krm.RuntimeShieldedInstanceConfig {
	if in == nil {
		return nil
	}
	out := &krm.RuntimeShieldedInstanceConfig{}
	out.EnableSecureBoot = direct.LazyPtr(in.GetEnableSecureBoot())
	out.EnableVTPM = direct.LazyPtr(in.GetEnableVtpm())
	out.EnableIntegrityMonitoring = direct.LazyPtr(in.GetEnableIntegrityMonitoring())
	return out
}
func RuntimeShieldedInstanceConfig_ToProto(mapCtx *direct.MapContext, in *krm.RuntimeShieldedInstanceConfig) *pb.RuntimeShieldedInstanceConfig {
	if in == nil {
		return nil
	}
	out := &pb.RuntimeShieldedInstanceConfig{}
	out.EnableSecureBoot = direct.ValueOf(in.EnableSecureBoot)
	out.EnableVtpm = direct.ValueOf(in.EnableVTPM)
	out.EnableIntegrityMonitoring = direct.ValueOf(in.EnableIntegrityMonitoring)
	return out
}
func RuntimeSoftwareConfig_FromProto(mapCtx *direct.MapContext, in *pb.RuntimeSoftwareConfig) *krm.RuntimeSoftwareConfig {
	if in == nil {
		return nil
	}
	out := &krm.RuntimeSoftwareConfig{}
	out.NotebookUpgradeSchedule = direct.LazyPtr(in.GetNotebookUpgradeSchedule())
	out.EnableHealthMonitoring = in.EnableHealthMonitoring
	out.IdleShutdown = in.IdleShutdown
	out.IdleShutdownTimeout = direct.LazyPtr(in.GetIdleShutdownTimeout())
	out.InstallGpuDriver = direct.LazyPtr(in.GetInstallGpuDriver())
	out.CustomGpuDriverPath = direct.LazyPtr(in.GetCustomGpuDriverPath())
	out.PostStartupScript = direct.LazyPtr(in.GetPostStartupScript())
	out.Kernels = direct.Slice_FromProto(mapCtx, in.Kernels, ContainerImage_FromProto)
	// MISSING: Upgradeable
	out.PostStartupScriptBehavior = direct.Enum_FromProto(mapCtx, in.GetPostStartupScriptBehavior())
	out.DisableTerminal = in.DisableTerminal
	// MISSING: Version
	return out
}
func RuntimeSoftwareConfig_ToProto(mapCtx *direct.MapContext, in *krm.RuntimeSoftwareConfig) *pb.RuntimeSoftwareConfig {
	if in == nil {
		return nil
	}
	out := &pb.RuntimeSoftwareConfig{}
	out.NotebookUpgradeSchedule = direct.ValueOf(in.NotebookUpgradeSchedule)
	out.EnableHealthMonitoring = in.EnableHealthMonitoring
	out.IdleShutdown = in.IdleShutdown
	out.IdleShutdownTimeout = direct.ValueOf(in.IdleShutdownTimeout)
	out.InstallGpuDriver = direct.ValueOf(in.InstallGpuDriver)
	out.CustomGpuDriverPath = direct.ValueOf(in.CustomGpuDriverPath)
	out.PostStartupScript = direct.ValueOf(in.PostStartupScript)
	out.Kernels = direct.Slice_ToProto(mapCtx, in.Kernels, ContainerImage_ToProto)
	// MISSING: Upgradeable
	out.PostStartupScriptBehavior = direct.Enum_ToProto[pb.RuntimeSoftwareConfig_PostStartupScriptBehavior](mapCtx, in.PostStartupScriptBehavior)
	out.DisableTerminal = in.DisableTerminal
	// MISSING: Version
	return out
}
func RuntimeSoftwareConfigObservedState_FromProto(mapCtx *direct.MapContext, in *pb.RuntimeSoftwareConfig) *krm.RuntimeSoftwareConfigObservedState {
	if in == nil {
		return nil
	}
	out := &krm.RuntimeSoftwareConfigObservedState{}
	// MISSING: NotebookUpgradeSchedule
	// MISSING: EnableHealthMonitoring
	// MISSING: IdleShutdown
	// MISSING: IdleShutdownTimeout
	// MISSING: InstallGpuDriver
	// MISSING: CustomGpuDriverPath
	// MISSING: PostStartupScript
	// MISSING: Kernels
	out.Upgradeable = in.Upgradeable
	// MISSING: PostStartupScriptBehavior
	// MISSING: DisableTerminal
	out.Version = in.Version
	return out
}
func RuntimeSoftwareConfigObservedState_ToProto(mapCtx *direct.MapContext, in *krm.RuntimeSoftwareConfigObservedState) *pb.RuntimeSoftwareConfig {
	if in == nil {
		return nil
	}
	out := &pb.RuntimeSoftwareConfig{}
	// MISSING: NotebookUpgradeSchedule
	// MISSING: EnableHealthMonitoring
	// MISSING: IdleShutdown
	// MISSING: IdleShutdownTimeout
	// MISSING: InstallGpuDriver
	// MISSING: CustomGpuDriverPath
	// MISSING: PostStartupScript
	// MISSING: Kernels
	out.Upgradeable = in.Upgradeable
	// MISSING: PostStartupScriptBehavior
	// MISSING: DisableTerminal
	out.Version = in.Version
	return out
}
func VirtualMachine_FromProto(mapCtx *direct.MapContext, in *pb.VirtualMachine) *krm.VirtualMachine {
	if in == nil {
		return nil
	}
	out := &krm.VirtualMachine{}
	// MISSING: InstanceName
	// MISSING: InstanceID
	out.VirtualMachineConfig = VirtualMachineConfig_FromProto(mapCtx, in.GetVirtualMachineConfig())
	return out
}
func VirtualMachine_ToProto(mapCtx *direct.MapContext, in *krm.VirtualMachine) *pb.VirtualMachine {
	if in == nil {
		return nil
	}
	out := &pb.VirtualMachine{}
	// MISSING: InstanceName
	// MISSING: InstanceID
	out.VirtualMachineConfig = VirtualMachineConfig_ToProto(mapCtx, in.VirtualMachineConfig)
	return out
}
func VirtualMachineConfig_FromProto(mapCtx *direct.MapContext, in *pb.VirtualMachineConfig) *krm.VirtualMachineConfig {
	if in == nil {
		return nil
	}
	out := &krm.VirtualMachineConfig{}
	// MISSING: Zone
	out.MachineType = direct.LazyPtr(in.GetMachineType())
	out.ContainerImages = direct.Slice_FromProto(mapCtx, in.ContainerImages, ContainerImage_FromProto)
	out.DataDisk = LocalDisk_FromProto(mapCtx, in.GetDataDisk())
	out.EncryptionConfig = EncryptionConfig_FromProto(mapCtx, in.GetEncryptionConfig())
	out.ShieldedInstanceConfig = RuntimeShieldedInstanceConfig_FromProto(mapCtx, in.GetShieldedInstanceConfig())
	out.AcceleratorConfig = RuntimeAcceleratorConfig_FromProto(mapCtx, in.GetAcceleratorConfig())
	out.Network = direct.LazyPtr(in.GetNetwork())
	out.Subnet = direct.LazyPtr(in.GetSubnet())
	out.InternalIPOnly = direct.LazyPtr(in.GetInternalIpOnly())
	out.Tags = in.Tags
	// MISSING: GuestAttributes
	out.Metadata = in.Metadata
	out.Labels = in.Labels
	out.NicType = direct.Enum_FromProto(mapCtx, in.GetNicType())
	out.ReservedIPRange = direct.LazyPtr(in.GetReservedIpRange())
	out.BootImage = VirtualMachineConfig_BootImage_FromProto(mapCtx, in.GetBootImage())
	return out
}
func VirtualMachineConfig_ToProto(mapCtx *direct.MapContext, in *krm.VirtualMachineConfig) *pb.VirtualMachineConfig {
	if in == nil {
		return nil
	}
	out := &pb.VirtualMachineConfig{}
	// MISSING: Zone
	out.MachineType = direct.ValueOf(in.MachineType)
	out.ContainerImages = direct.Slice_ToProto(mapCtx, in.ContainerImages, ContainerImage_ToProto)
	out.DataDisk = LocalDisk_ToProto(mapCtx, in.DataDisk)
	out.EncryptionConfig = EncryptionConfig_ToProto(mapCtx, in.EncryptionConfig)
	out.ShieldedInstanceConfig = RuntimeShieldedInstanceConfig_ToProto(mapCtx, in.ShieldedInstanceConfig)
	out.AcceleratorConfig = RuntimeAcceleratorConfig_ToProto(mapCtx, in.AcceleratorConfig)
	out.Network = direct.ValueOf(in.Network)
	out.Subnet = direct.ValueOf(in.Subnet)
	out.InternalIpOnly = direct.ValueOf(in.InternalIPOnly)
	out.Tags = in.Tags
	// MISSING: GuestAttributes
	out.Metadata = in.Metadata
	out.Labels = in.Labels
	out.NicType = direct.Enum_ToProto[pb.VirtualMachineConfig_NicType](mapCtx, in.NicType)
	out.ReservedIpRange = direct.ValueOf(in.ReservedIPRange)
	out.BootImage = VirtualMachineConfig_BootImage_ToProto(mapCtx, in.BootImage)
	return out
}
func VirtualMachineConfigObservedState_FromProto(mapCtx *direct.MapContext, in *pb.VirtualMachineConfig) *krm.VirtualMachineConfigObservedState {
	if in == nil {
		return nil
	}
	out := &krm.VirtualMachineConfigObservedState{}
	out.Zone = direct.LazyPtr(in.GetZone())
	// MISSING: MachineType
	// MISSING: ContainerImages
	out.DataDisk = LocalDiskObservedState_FromProto(mapCtx, in.GetDataDisk())
	// MISSING: EncryptionConfig
	// MISSING: ShieldedInstanceConfig
	// MISSING: AcceleratorConfig
	// MISSING: Network
	// MISSING: Subnet
	// MISSING: InternalIPOnly
	// MISSING: Tags
	out.GuestAttributes = in.GuestAttributes
	// MISSING: Metadata
	// MISSING: Labels
	// MISSING: NicType
	// MISSING: ReservedIPRange
	// MISSING: BootImage
	return out
}
func VirtualMachineConfigObservedState_ToProto(mapCtx *direct.MapContext, in *krm.VirtualMachineConfigObservedState) *pb.VirtualMachineConfig {
	if in == nil {
		return nil
	}
	out := &pb.VirtualMachineConfig{}
	out.Zone = direct.ValueOf(in.Zone)
	// MISSING: MachineType
	// MISSING: ContainerImages
	out.DataDisk = LocalDiskObservedState_ToProto(mapCtx, in.DataDisk)
	// MISSING: EncryptionConfig
	// MISSING: ShieldedInstanceConfig
	// MISSING: AcceleratorConfig
	// MISSING: Network
	// MISSING: Subnet
	// MISSING: InternalIPOnly
	// MISSING: Tags
	out.GuestAttributes = in.GuestAttributes
	// MISSING: Metadata
	// MISSING: Labels
	// MISSING: NicType
	// MISSING: ReservedIPRange
	// MISSING: BootImage
	return out
}
func VirtualMachineConfig_BootImage_FromProto(mapCtx *direct.MapContext, in *pb.VirtualMachineConfig_BootImage) *krm.VirtualMachineConfig_BootImage {
	if in == nil {
		return nil
	}
	out := &krm.VirtualMachineConfig_BootImage{}
	return out
}
func VirtualMachineConfig_BootImage_ToProto(mapCtx *direct.MapContext, in *krm.VirtualMachineConfig_BootImage) *pb.VirtualMachineConfig_BootImage {
	if in == nil {
		return nil
	}
	out := &pb.VirtualMachineConfig_BootImage{}
	return out
}
func VirtualMachineObservedState_FromProto(mapCtx *direct.MapContext, in *pb.VirtualMachine) *krm.VirtualMachineObservedState {
	if in == nil {
		return nil
	}
	out := &krm.VirtualMachineObservedState{}
	out.InstanceName = direct.LazyPtr(in.GetInstanceName())
	out.InstanceID = direct.LazyPtr(in.GetInstanceId())
	out.VirtualMachineConfig = VirtualMachineConfigObservedState_FromProto(mapCtx, in.GetVirtualMachineConfig())
	return out
}
func VirtualMachineObservedState_ToProto(mapCtx *direct.MapContext, in *krm.VirtualMachineObservedState) *pb.VirtualMachine {
	if in == nil {
		return nil
	}
	out := &pb.VirtualMachine{}
	out.InstanceName = direct.ValueOf(in.InstanceName)
	out.InstanceId = direct.ValueOf(in.InstanceID)
	out.VirtualMachineConfig = VirtualMachineConfigObservedState_ToProto(mapCtx, in.VirtualMachineConfig)
	return out
}
