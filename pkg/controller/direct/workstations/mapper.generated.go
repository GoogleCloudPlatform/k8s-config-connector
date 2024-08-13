// Copyright 2024 Google LLC
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

package workstations

import (
	pb "cloud.google.com/go/workstations/apiv1/workstationspb"
	krm "github.com/GoogleCloudPlatform/k8s-config-connector/apis/workstations/v1alpha1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct"
)

func WorkstationCluster_PrivateClusterConfig_FromProto(mapCtx *direct.MapContext, in *pb.WorkstationCluster_PrivateClusterConfig) *krm.WorkstationCluster_PrivateClusterConfig {
	if in == nil {
		return nil
	}
	out := &krm.WorkstationCluster_PrivateClusterConfig{}
	out.EnablePrivateEndpoint = direct.LazyPtr(in.GetEnablePrivateEndpoint())
	// MISSING: ClusterHostname
	// MISSING: ServiceAttachmentUri
	out.AllowedProjects = in.AllowedProjects
	return out
}
func WorkstationCluster_PrivateClusterConfig_ToProto(mapCtx *direct.MapContext, in *krm.WorkstationCluster_PrivateClusterConfig) *pb.WorkstationCluster_PrivateClusterConfig {
	if in == nil {
		return nil
	}
	out := &pb.WorkstationCluster_PrivateClusterConfig{}
	out.EnablePrivateEndpoint = direct.ValueOf(in.EnablePrivateEndpoint)
	// MISSING: ClusterHostname
	// MISSING: ServiceAttachmentUri
	out.AllowedProjects = in.AllowedProjects
	return out
}
func WorkstationConfig_Container_FromProto(mapCtx *direct.MapContext, in *pb.WorkstationConfig_Container) *krm.WorkstationConfig_Container {
	if in == nil {
		return nil
	}
	out := &krm.WorkstationConfig_Container{}
	out.Image = direct.LazyPtr(in.GetImage())
	out.Command = in.Command
	out.Args = in.Args
	out.Env = in.Env
	out.WorkingDir = direct.LazyPtr(in.GetWorkingDir())
	out.RunAsUser = direct.LazyPtr(in.GetRunAsUser())
	return out
}
func WorkstationConfig_Container_ToProto(mapCtx *direct.MapContext, in *krm.WorkstationConfig_Container) *pb.WorkstationConfig_Container {
	if in == nil {
		return nil
	}
	out := &pb.WorkstationConfig_Container{}
	out.Image = direct.ValueOf(in.Image)
	out.Command = in.Command
	out.Args = in.Args
	out.Env = in.Env
	out.WorkingDir = direct.ValueOf(in.WorkingDir)
	out.RunAsUser = direct.ValueOf(in.RunAsUser)
	return out
}
func WorkstationConfig_CustomerEncryptionKey_FromProto(mapCtx *direct.MapContext, in *pb.WorkstationConfig_CustomerEncryptionKey) *krm.WorkstationConfig_CustomerEncryptionKey {
	if in == nil {
		return nil
	}
	out := &krm.WorkstationConfig_CustomerEncryptionKey{}
	out.KmsKey = direct.LazyPtr(in.GetKmsKey())
	out.KmsKeyServiceAccount = direct.LazyPtr(in.GetKmsKeyServiceAccount())
	return out
}
func WorkstationConfig_CustomerEncryptionKey_ToProto(mapCtx *direct.MapContext, in *krm.WorkstationConfig_CustomerEncryptionKey) *pb.WorkstationConfig_CustomerEncryptionKey {
	if in == nil {
		return nil
	}
	out := &pb.WorkstationConfig_CustomerEncryptionKey{}
	out.KmsKey = direct.ValueOf(in.KmsKey)
	out.KmsKeyServiceAccount = direct.ValueOf(in.KmsKeyServiceAccount)
	return out
}
func WorkstationConfig_Host_FromProto(mapCtx *direct.MapContext, in *pb.WorkstationConfig_Host) *krm.WorkstationConfig_Host {
	if in == nil {
		return nil
	}
	out := &krm.WorkstationConfig_Host{}
	out.GceInstance = WorkstationConfig_Host_GceInstance_FromProto(mapCtx, in.GetGceInstance())
	return out
}
func WorkstationConfig_Host_ToProto(mapCtx *direct.MapContext, in *krm.WorkstationConfig_Host) *pb.WorkstationConfig_Host {
	if in == nil {
		return nil
	}
	out := &pb.WorkstationConfig_Host{}
	if oneof := WorkstationConfig_Host_GceInstance_ToProto(mapCtx, in.GceInstance); oneof != nil {
		out.Config = &pb.WorkstationConfig_Host_GceInstance_{GceInstance: oneof}
	}
	return out
}
func WorkstationConfig_Host_GceInstance_FromProto(mapCtx *direct.MapContext, in *pb.WorkstationConfig_Host_GceInstance) *krm.WorkstationConfig_Host_GceInstance {
	if in == nil {
		return nil
	}
	out := &krm.WorkstationConfig_Host_GceInstance{}
	out.MachineType = direct.LazyPtr(in.GetMachineType())
	out.ServiceAccount = direct.LazyPtr(in.GetServiceAccount())
	out.ServiceAccountScopes = in.ServiceAccountScopes
	out.Tags = in.Tags
	out.PoolSize = direct.LazyPtr(in.GetPoolSize())
	out.PooledInstances = direct.LazyPtr(in.GetPooledInstances())
	out.DisablePublicIPAddresses = direct.LazyPtr(in.GetDisablePublicIpAddresses())
	out.EnableNestedVirtualization = direct.LazyPtr(in.GetEnableNestedVirtualization())
	out.ShieldedInstanceConfig = WorkstationConfig_Host_GceInstance_GceShieldedInstanceConfig_FromProto(mapCtx, in.GetShieldedInstanceConfig())
	out.ConfidentialInstanceConfig = WorkstationConfig_Host_GceInstance_GceConfidentialInstanceConfig_FromProto(mapCtx, in.GetConfidentialInstanceConfig())
	out.BootDiskSizeGb = direct.LazyPtr(in.GetBootDiskSizeGb())
	return out
}
func WorkstationConfig_Host_GceInstance_ToProto(mapCtx *direct.MapContext, in *krm.WorkstationConfig_Host_GceInstance) *pb.WorkstationConfig_Host_GceInstance {
	if in == nil {
		return nil
	}
	out := &pb.WorkstationConfig_Host_GceInstance{}
	out.MachineType = direct.ValueOf(in.MachineType)
	out.ServiceAccount = direct.ValueOf(in.ServiceAccount)
	out.ServiceAccountScopes = in.ServiceAccountScopes
	out.Tags = in.Tags
	out.PoolSize = direct.ValueOf(in.PoolSize)
	out.PooledInstances = direct.ValueOf(in.PooledInstances)
	out.DisablePublicIpAddresses = direct.ValueOf(in.DisablePublicIPAddresses)
	out.EnableNestedVirtualization = direct.ValueOf(in.EnableNestedVirtualization)
	out.ShieldedInstanceConfig = WorkstationConfig_Host_GceInstance_GceShieldedInstanceConfig_ToProto(mapCtx, in.ShieldedInstanceConfig)
	out.ConfidentialInstanceConfig = WorkstationConfig_Host_GceInstance_GceConfidentialInstanceConfig_ToProto(mapCtx, in.ConfidentialInstanceConfig)
	out.BootDiskSizeGb = direct.ValueOf(in.BootDiskSizeGb)
	return out
}
func WorkstationConfig_Host_GceInstance_GceConfidentialInstanceConfig_FromProto(mapCtx *direct.MapContext, in *pb.WorkstationConfig_Host_GceInstance_GceConfidentialInstanceConfig) *krm.WorkstationConfig_Host_GceInstance_GceConfidentialInstanceConfig {
	if in == nil {
		return nil
	}
	out := &krm.WorkstationConfig_Host_GceInstance_GceConfidentialInstanceConfig{}
	out.EnableConfidentialCompute = direct.LazyPtr(in.GetEnableConfidentialCompute())
	return out
}
func WorkstationConfig_Host_GceInstance_GceConfidentialInstanceConfig_ToProto(mapCtx *direct.MapContext, in *krm.WorkstationConfig_Host_GceInstance_GceConfidentialInstanceConfig) *pb.WorkstationConfig_Host_GceInstance_GceConfidentialInstanceConfig {
	if in == nil {
		return nil
	}
	out := &pb.WorkstationConfig_Host_GceInstance_GceConfidentialInstanceConfig{}
	out.EnableConfidentialCompute = direct.ValueOf(in.EnableConfidentialCompute)
	return out
}
func WorkstationConfig_Host_GceInstance_GceShieldedInstanceConfig_FromProto(mapCtx *direct.MapContext, in *pb.WorkstationConfig_Host_GceInstance_GceShieldedInstanceConfig) *krm.WorkstationConfig_Host_GceInstance_GceShieldedInstanceConfig {
	if in == nil {
		return nil
	}
	out := &krm.WorkstationConfig_Host_GceInstance_GceShieldedInstanceConfig{}
	out.EnableSecureBoot = direct.LazyPtr(in.GetEnableSecureBoot())
	out.EnableVtpm = direct.LazyPtr(in.GetEnableVtpm())
	out.EnableIntegrityMonitoring = direct.LazyPtr(in.GetEnableIntegrityMonitoring())
	return out
}
func WorkstationConfig_Host_GceInstance_GceShieldedInstanceConfig_ToProto(mapCtx *direct.MapContext, in *krm.WorkstationConfig_Host_GceInstance_GceShieldedInstanceConfig) *pb.WorkstationConfig_Host_GceInstance_GceShieldedInstanceConfig {
	if in == nil {
		return nil
	}
	out := &pb.WorkstationConfig_Host_GceInstance_GceShieldedInstanceConfig{}
	out.EnableSecureBoot = direct.ValueOf(in.EnableSecureBoot)
	out.EnableVtpm = direct.ValueOf(in.EnableVtpm)
	out.EnableIntegrityMonitoring = direct.ValueOf(in.EnableIntegrityMonitoring)
	return out
}
func WorkstationConfig_PersistentDirectory_FromProto(mapCtx *direct.MapContext, in *pb.WorkstationConfig_PersistentDirectory) *krm.WorkstationConfig_PersistentDirectory {
	if in == nil {
		return nil
	}
	out := &krm.WorkstationConfig_PersistentDirectory{}
	out.GcePd = WorkstationConfig_PersistentDirectory_GceRegionalPersistentDisk_FromProto(mapCtx, in.GetGcePd())
	out.MountPath = direct.LazyPtr(in.GetMountPath())
	return out
}
func WorkstationConfig_PersistentDirectory_ToProto(mapCtx *direct.MapContext, in *krm.WorkstationConfig_PersistentDirectory) *pb.WorkstationConfig_PersistentDirectory {
	if in == nil {
		return nil
	}
	out := &pb.WorkstationConfig_PersistentDirectory{}
	if oneof := WorkstationConfig_PersistentDirectory_GceRegionalPersistentDisk_ToProto(mapCtx, in.GcePd); oneof != nil {
		out.DirectoryType = &pb.WorkstationConfig_PersistentDirectory_GcePd{GcePd: oneof}
	}
	out.MountPath = direct.ValueOf(in.MountPath)
	return out
}
func WorkstationConfig_PersistentDirectory_GceRegionalPersistentDisk_FromProto(mapCtx *direct.MapContext, in *pb.WorkstationConfig_PersistentDirectory_GceRegionalPersistentDisk) *krm.WorkstationConfig_PersistentDirectory_GceRegionalPersistentDisk {
	if in == nil {
		return nil
	}
	out := &krm.WorkstationConfig_PersistentDirectory_GceRegionalPersistentDisk{}
	out.SizeGb = direct.LazyPtr(in.GetSizeGb())
	out.FsType = direct.LazyPtr(in.GetFsType())
	out.DiskType = direct.LazyPtr(in.GetDiskType())
	out.SourceSnapshot = direct.LazyPtr(in.GetSourceSnapshot())
	out.ReclaimPolicy = direct.Enum_FromProto(mapCtx, in.GetReclaimPolicy())
	return out
}
func WorkstationConfig_PersistentDirectory_GceRegionalPersistentDisk_ToProto(mapCtx *direct.MapContext, in *krm.WorkstationConfig_PersistentDirectory_GceRegionalPersistentDisk) *pb.WorkstationConfig_PersistentDirectory_GceRegionalPersistentDisk {
	if in == nil {
		return nil
	}
	out := &pb.WorkstationConfig_PersistentDirectory_GceRegionalPersistentDisk{}
	out.SizeGb = direct.ValueOf(in.SizeGb)
	out.FsType = direct.ValueOf(in.FsType)
	out.DiskType = direct.ValueOf(in.DiskType)
	out.SourceSnapshot = direct.ValueOf(in.SourceSnapshot)
	out.ReclaimPolicy = direct.Enum_ToProto[pb.WorkstationConfig_PersistentDirectory_GceRegionalPersistentDisk_ReclaimPolicy](mapCtx, in.ReclaimPolicy)
	return out
}
func WorkstationConfig_ReadinessCheck_FromProto(mapCtx *direct.MapContext, in *pb.WorkstationConfig_ReadinessCheck) *krm.WorkstationConfig_ReadinessCheck {
	if in == nil {
		return nil
	}
	out := &krm.WorkstationConfig_ReadinessCheck{}
	out.Path = direct.LazyPtr(in.GetPath())
	out.Port = direct.LazyPtr(in.GetPort())
	return out
}
func WorkstationConfig_ReadinessCheck_ToProto(mapCtx *direct.MapContext, in *krm.WorkstationConfig_ReadinessCheck) *pb.WorkstationConfig_ReadinessCheck {
	if in == nil {
		return nil
	}
	out := &pb.WorkstationConfig_ReadinessCheck{}
	out.Path = direct.ValueOf(in.Path)
	out.Port = direct.ValueOf(in.Port)
	return out
}
func WorkstationsWorkstationClusterObservedState_FromProto(mapCtx *direct.MapContext, in *pb.WorkstationCluster) *krm.WorkstationsWorkstationClusterObservedState {
	if in == nil {
		return nil
	}
	out := &krm.WorkstationsWorkstationClusterObservedState{}
	// MISSING: Name
	// MISSING: DisplayName
	out.Uid = direct.LazyPtr(in.GetUid())
	// MISSING: Reconciling
	// MISSING: Annotations
	// MISSING: Labels
	out.CreateTime = WorkstationCluster_CreateTime_FromProto(mapCtx, in.GetCreateTime())
	out.UpdateTime = WorkstationCluster_UpdateTime_FromProto(mapCtx, in.GetUpdateTime())
	out.DeleteTime = WorkstationCluster_DeleteTime_FromProto(mapCtx, in.GetDeleteTime())
	out.Etag = direct.LazyPtr(in.GetEtag())
	// MISSING: Network
	// MISSING: Subnetwork
	out.ControlPlaneIP = direct.LazyPtr(in.GetControlPlaneIp())
	// MISSING: PrivateClusterConfig
	// MISSING: Degraded
	// MISSING: Conditions
	return out
}
func WorkstationsWorkstationClusterObservedState_ToProto(mapCtx *direct.MapContext, in *krm.WorkstationsWorkstationClusterObservedState) *pb.WorkstationCluster {
	if in == nil {
		return nil
	}
	out := &pb.WorkstationCluster{}
	// MISSING: Name
	// MISSING: DisplayName
	out.Uid = direct.ValueOf(in.Uid)
	// MISSING: Reconciling
	// MISSING: Annotations
	// MISSING: Labels
	out.CreateTime = WorkstationCluster_CreateTime_ToProto(mapCtx, in.CreateTime)
	out.UpdateTime = WorkstationCluster_UpdateTime_ToProto(mapCtx, in.UpdateTime)
	out.DeleteTime = WorkstationCluster_DeleteTime_ToProto(mapCtx, in.DeleteTime)
	out.Etag = direct.ValueOf(in.Etag)
	// MISSING: Network
	// MISSING: Subnetwork
	out.ControlPlaneIp = direct.ValueOf(in.ControlPlaneIP)
	// MISSING: PrivateClusterConfig
	// MISSING: Degraded
	// MISSING: Conditions
	return out
}
func WorkstationsWorkstationClusterSpec_FromProto(mapCtx *direct.MapContext, in *pb.WorkstationCluster) *krm.WorkstationsWorkstationClusterSpec {
	if in == nil {
		return nil
	}
	out := &krm.WorkstationsWorkstationClusterSpec{}
	// MISSING: Name
	out.DisplayName = direct.LazyPtr(in.GetDisplayName())
	// MISSING: Uid
	// MISSING: Reconciling
	out.Annotations = in.Annotations
	// MISSING: Labels
	// MISSING: CreateTime
	// MISSING: UpdateTime
	// MISSING: DeleteTime
	// MISSING: Etag
	out.Network = direct.LazyPtr(in.GetNetwork())
	out.Subnetwork = direct.LazyPtr(in.GetSubnetwork())
	// MISSING: ControlPlaneIP
	out.PrivateClusterConfig = WorkstationCluster_PrivateClusterConfig_FromProto(mapCtx, in.GetPrivateClusterConfig())
	// MISSING: Degraded
	// MISSING: Conditions
	return out
}
func WorkstationsWorkstationClusterSpec_ToProto(mapCtx *direct.MapContext, in *krm.WorkstationsWorkstationClusterSpec) *pb.WorkstationCluster {
	if in == nil {
		return nil
	}
	out := &pb.WorkstationCluster{}
	// MISSING: Name
	out.DisplayName = direct.ValueOf(in.DisplayName)
	// MISSING: Uid
	// MISSING: Reconciling
	out.Annotations = in.Annotations
	// MISSING: Labels
	// MISSING: CreateTime
	// MISSING: UpdateTime
	// MISSING: DeleteTime
	// MISSING: Etag
	out.Network = direct.ValueOf(in.Network)
	out.Subnetwork = direct.ValueOf(in.Subnetwork)
	// MISSING: ControlPlaneIP
	out.PrivateClusterConfig = WorkstationCluster_PrivateClusterConfig_ToProto(mapCtx, in.PrivateClusterConfig)
	// MISSING: Degraded
	// MISSING: Conditions
	return out
}
func WorkstationsWorkstationConfigObservedState_FromProto(mapCtx *direct.MapContext, in *pb.WorkstationConfig) *krm.WorkstationsWorkstationConfigObservedState {
	if in == nil {
		return nil
	}
	out := &krm.WorkstationsWorkstationConfigObservedState{}
	// MISSING: Name
	// MISSING: DisplayName
	out.Uid = direct.LazyPtr(in.GetUid())
	// MISSING: Reconciling
	// MISSING: Annotations
	// MISSING: Labels
	out.CreateTime = WorkstationConfig_CreateTime_FromProto(mapCtx, in.GetCreateTime())
	out.UpdateTime = WorkstationConfig_UpdateTime_FromProto(mapCtx, in.GetUpdateTime())
	out.DeleteTime = WorkstationConfig_DeleteTime_FromProto(mapCtx, in.GetDeleteTime())
	out.Etag = direct.LazyPtr(in.GetEtag())
	// MISSING: IdleTimeout
	// MISSING: RunningTimeout
	// MISSING: Host
	// MISSING: PersistentDirectories
	// MISSING: Container
	// MISSING: EncryptionKey
	// MISSING: ReadinessChecks
	// MISSING: ReplicaZones
	// MISSING: Degraded
	// MISSING: Conditions
	return out
}
func WorkstationsWorkstationConfigObservedState_ToProto(mapCtx *direct.MapContext, in *krm.WorkstationsWorkstationConfigObservedState) *pb.WorkstationConfig {
	if in == nil {
		return nil
	}
	out := &pb.WorkstationConfig{}
	// MISSING: Name
	// MISSING: DisplayName
	out.Uid = direct.ValueOf(in.Uid)
	// MISSING: Reconciling
	// MISSING: Annotations
	// MISSING: Labels
	out.CreateTime = WorkstationConfig_CreateTime_ToProto(mapCtx, in.CreateTime)
	out.UpdateTime = WorkstationConfig_UpdateTime_ToProto(mapCtx, in.UpdateTime)
	out.DeleteTime = WorkstationConfig_DeleteTime_ToProto(mapCtx, in.DeleteTime)
	out.Etag = direct.ValueOf(in.Etag)
	// MISSING: IdleTimeout
	// MISSING: RunningTimeout
	// MISSING: Host
	// MISSING: PersistentDirectories
	// MISSING: Container
	// MISSING: EncryptionKey
	// MISSING: ReadinessChecks
	// MISSING: ReplicaZones
	// MISSING: Degraded
	// MISSING: Conditions
	return out
}
func WorkstationsWorkstationConfigSpec_FromProto(mapCtx *direct.MapContext, in *pb.WorkstationConfig) *krm.WorkstationsWorkstationConfigSpec {
	if in == nil {
		return nil
	}
	out := &krm.WorkstationsWorkstationConfigSpec{}
	out.Name = direct.LazyPtr(in.GetName())
	out.DisplayName = direct.LazyPtr(in.GetDisplayName())
	// MISSING: Uid
	// MISSING: Reconciling
	out.Annotations = in.Annotations
	// MISSING: Labels
	// MISSING: CreateTime
	// MISSING: UpdateTime
	// MISSING: DeleteTime
	// MISSING: Etag
	out.IdleTimeout = WorkstationConfig_IdleTimeout_FromProto(mapCtx, in.GetIdleTimeout())
	out.RunningTimeout = WorkstationConfig_RunningTimeout_FromProto(mapCtx, in.GetRunningTimeout())
	out.Host = WorkstationConfig_Host_FromProto(mapCtx, in.GetHost())
	out.PersistentDirectories = direct.Slice_FromProto(mapCtx, in.PersistentDirectories, WorkstationConfig_PersistentDirectory_FromProto)
	out.Container = WorkstationConfig_Container_FromProto(mapCtx, in.GetContainer())
	out.EncryptionKey = WorkstationConfig_CustomerEncryptionKey_FromProto(mapCtx, in.GetEncryptionKey())
	out.ReadinessChecks = direct.Slice_FromProto(mapCtx, in.ReadinessChecks, WorkstationConfig_ReadinessCheck_FromProto)
	out.ReplicaZones = in.ReplicaZones
	// MISSING: Degraded
	// MISSING: Conditions
	return out
}
func WorkstationsWorkstationConfigSpec_ToProto(mapCtx *direct.MapContext, in *krm.WorkstationsWorkstationConfigSpec) *pb.WorkstationConfig {
	if in == nil {
		return nil
	}
	out := &pb.WorkstationConfig{}
	out.Name = direct.ValueOf(in.Name)
	out.DisplayName = direct.ValueOf(in.DisplayName)
	// MISSING: Uid
	// MISSING: Reconciling
	out.Annotations = in.Annotations
	// MISSING: Labels
	// MISSING: CreateTime
	// MISSING: UpdateTime
	// MISSING: DeleteTime
	// MISSING: Etag
	out.IdleTimeout = WorkstationConfig_IdleTimeout_ToProto(mapCtx, in.IdleTimeout)
	out.RunningTimeout = WorkstationConfig_RunningTimeout_ToProto(mapCtx, in.RunningTimeout)
	out.Host = WorkstationConfig_Host_ToProto(mapCtx, in.Host)
	out.PersistentDirectories = direct.Slice_ToProto(mapCtx, in.PersistentDirectories, WorkstationConfig_PersistentDirectory_ToProto)
	out.Container = WorkstationConfig_Container_ToProto(mapCtx, in.Container)
	out.EncryptionKey = WorkstationConfig_CustomerEncryptionKey_ToProto(mapCtx, in.EncryptionKey)
	out.ReadinessChecks = direct.Slice_ToProto(mapCtx, in.ReadinessChecks, WorkstationConfig_ReadinessCheck_ToProto)
	out.ReplicaZones = in.ReplicaZones
	// MISSING: Degraded
	// MISSING: Conditions
	return out
}
func WorkstationsWorkstationObservedState_FromProto(mapCtx *direct.MapContext, in *pb.Workstation) *krm.WorkstationsWorkstationObservedState {
	if in == nil {
		return nil
	}
	out := &krm.WorkstationsWorkstationObservedState{}
	// MISSING: Name
	// MISSING: DisplayName
	out.Uid = direct.LazyPtr(in.GetUid())
	// MISSING: Reconciling
	// MISSING: Annotations
	// MISSING: Labels
	out.CreateTime = Workstation_CreateTime_FromProto(mapCtx, in.GetCreateTime())
	out.UpdateTime = Workstation_UpdateTime_FromProto(mapCtx, in.GetUpdateTime())
	out.StartTime = Workstation_StartTime_FromProto(mapCtx, in.GetStartTime())
	out.DeleteTime = Workstation_DeleteTime_FromProto(mapCtx, in.GetDeleteTime())
	out.Etag = direct.LazyPtr(in.GetEtag())
	out.State = direct.Enum_FromProto(mapCtx, in.GetState())
	out.Host = direct.LazyPtr(in.GetHost())
	return out
}
func WorkstationsWorkstationObservedState_ToProto(mapCtx *direct.MapContext, in *krm.WorkstationsWorkstationObservedState) *pb.Workstation {
	if in == nil {
		return nil
	}
	out := &pb.Workstation{}
	// MISSING: Name
	// MISSING: DisplayName
	out.Uid = direct.ValueOf(in.Uid)
	// MISSING: Reconciling
	// MISSING: Annotations
	// MISSING: Labels
	out.CreateTime = Workstation_CreateTime_ToProto(mapCtx, in.CreateTime)
	out.UpdateTime = Workstation_UpdateTime_ToProto(mapCtx, in.UpdateTime)
	out.StartTime = Workstation_StartTime_ToProto(mapCtx, in.StartTime)
	out.DeleteTime = Workstation_DeleteTime_ToProto(mapCtx, in.DeleteTime)
	out.Etag = direct.ValueOf(in.Etag)
	out.State = direct.Enum_ToProto[pb.Workstation_State](mapCtx, in.State)
	out.Host = direct.ValueOf(in.Host)
	return out
}
func WorkstationsWorkstationSpec_FromProto(mapCtx *direct.MapContext, in *pb.Workstation) *krm.WorkstationsWorkstationSpec {
	if in == nil {
		return nil
	}
	out := &krm.WorkstationsWorkstationSpec{}
	// MISSING: Name
	out.DisplayName = direct.LazyPtr(in.GetDisplayName())
	// MISSING: Uid
	// MISSING: Reconciling
	out.Annotations = in.Annotations
	// MISSING: Labels
	// MISSING: CreateTime
	// MISSING: UpdateTime
	// MISSING: StartTime
	// MISSING: DeleteTime
	// MISSING: Etag
	// MISSING: State
	// MISSING: Host
	return out
}
func WorkstationsWorkstationSpec_ToProto(mapCtx *direct.MapContext, in *krm.WorkstationsWorkstationSpec) *pb.Workstation {
	if in == nil {
		return nil
	}
	out := &pb.Workstation{}
	// MISSING: Name
	out.DisplayName = direct.ValueOf(in.DisplayName)
	// MISSING: Uid
	// MISSING: Reconciling
	out.Annotations = in.Annotations
	// MISSING: Labels
	// MISSING: CreateTime
	// MISSING: UpdateTime
	// MISSING: StartTime
	// MISSING: DeleteTime
	// MISSING: Etag
	// MISSING: State
	// MISSING: Host
	return out
}
