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
	refs "github.com/GoogleCloudPlatform/k8s-config-connector/apis/refs/v1beta1"
	krm "github.com/GoogleCloudPlatform/k8s-config-connector/apis/workstations/v1alpha1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct"
)

func WorkstationConfigObservedState_FromProto(mapCtx *direct.MapContext, in *pb.WorkstationConfig) *krm.WorkstationConfigObservedState {
	if in == nil {
		return nil
	}
	out := &krm.WorkstationConfigObservedState{}
	// MISSING: Name
	// MISSING: Uid
	// MISSING: Reconciling
	out.CreateTime = direct.StringTimestamp_FromProto(mapCtx, in.GetCreateTime())
	out.UpdateTime = direct.StringTimestamp_FromProto(mapCtx, in.GetUpdateTime())
	out.DeleteTime = direct.StringTimestamp_FromProto(mapCtx, in.GetDeleteTime())
	out.Etag = direct.LazyPtr(in.GetEtag())
	out.Degraded = direct.LazyPtr(in.GetDegraded())
	// MISSING: Conditions
	return out
}

func WorkstationConfigObservedState_ToProto(mapCtx *direct.MapContext, in *krm.WorkstationConfigObservedState) *pb.WorkstationConfig {
	if in == nil {
		return nil
	}
	out := &pb.WorkstationConfig{}
	// MISSING: Name
	// MISSING: Uid
	// MISSING: Reconciling
	out.CreateTime = direct.StringTimestamp_ToProto(mapCtx, in.CreateTime)
	out.UpdateTime = direct.StringTimestamp_ToProto(mapCtx, in.UpdateTime)
	out.DeleteTime = direct.StringTimestamp_ToProto(mapCtx, in.DeleteTime)
	out.Etag = direct.ValueOf(in.Etag)
	out.Degraded = direct.ValueOf(in.Degraded)
	// MISSING: Conditions
	return out
}

func WorkstationConfigSpec_FromProto(mapCtx *direct.MapContext, in *pb.WorkstationConfig) *krm.WorkstationConfigSpec {
	if in == nil {
		return nil
	}
	out := &krm.WorkstationConfigSpec{}
	// MISSING: Name
	out.DisplayName = direct.LazyPtr(in.GetDisplayName())
	// MISSING: Uid
	// MISSING: Reconciling
	out.Annotations = WorkstationAnnotations_FromProto_Alpha(mapCtx, in.Annotations)
	out.Labels = WorkstationLabels_FromProto_Alpha(mapCtx, in.Labels)
	out.IdleTimeout = direct.StringDuration_FromProto(mapCtx, in.GetIdleTimeout())
	out.RunningTimeout = direct.StringDuration_FromProto(mapCtx, in.GetRunningTimeout())
	out.Host = WorkstationConfig_Host_FromProto(mapCtx, in.GetHost())
	out.PersistentDirectories = direct.Slice_FromProto(mapCtx, in.PersistentDirectories, WorkstationConfig_PersistentDirectory_FromProto)
	out.Container = WorkstationConfig_Container_FromProto(mapCtx, in.GetContainer())
	out.EncryptionKey = WorkstationConfig_CustomerEncryptionKey_FromProto(mapCtx, in.GetEncryptionKey())
	out.ReadinessChecks = direct.Slice_FromProto(mapCtx, in.ReadinessChecks, WorkstationConfig_ReadinessCheck_FromProto)
	out.ReplicaZones = in.ReplicaZones
	// MISSING: Conditions
	return out
}

func WorkstationConfigSpec_ToProto(mapCtx *direct.MapContext, in *krm.WorkstationConfigSpec) *pb.WorkstationConfig {
	if in == nil {
		return nil
	}
	out := &pb.WorkstationConfig{}
	// MISSING: Name
	out.DisplayName = direct.ValueOf(in.DisplayName)
	// MISSING: Uid
	// MISSING: Reconciling
	out.Annotations = WorkstationAnnotations_ToProto_Alpha(mapCtx, in.Annotations)
	out.Labels = WorkstationLabels_ToProto_Alpha(mapCtx, in.Labels)
	out.IdleTimeout = direct.StringDuration_ToProto(mapCtx, in.IdleTimeout)
	out.RunningTimeout = direct.StringDuration_ToProto(mapCtx, in.RunningTimeout)
	out.Host = WorkstationConfig_Host_ToProto(mapCtx, in.Host)
	out.PersistentDirectories = direct.Slice_ToProto(mapCtx, in.PersistentDirectories, WorkstationConfig_PersistentDirectory_ToProto)
	out.Container = WorkstationConfig_Container_ToProto(mapCtx, in.Container)
	out.EncryptionKey = WorkstationConfig_CustomerEncryptionKey_ToProto(mapCtx, in.EncryptionKey)
	out.ReadinessChecks = direct.Slice_ToProto(mapCtx, in.ReadinessChecks, WorkstationConfig_ReadinessCheck_ToProto)
	out.ReplicaZones = in.ReplicaZones
	// MISSING: Conditions
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
	out.Env = WorkstationConfig_Container_EnvVar_FromProto(mapCtx, in.Env)
	out.WorkingDir = direct.LazyPtr(in.GetWorkingDir())
	out.RunAsUser = direct.LazyPtr(in.GetRunAsUser())
	return out
}

func WorkstationConfig_Container_EnvVar_FromProto(mapCtx *direct.MapContext, in map[string]string) []krm.WorkstationConfig_Container_EnvVar {
	if in == nil {
		return nil
	}
	out := []krm.WorkstationConfig_Container_EnvVar{}
	for name, value := range in {
		out = append(out, krm.WorkstationConfig_Container_EnvVar{
			Name:  name,
			Value: value,
		})
	}
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
	out.Env = WorkstationConfig_Container_EnvVar_ToProto(mapCtx, in.Env)
	out.WorkingDir = direct.ValueOf(in.WorkingDir)
	out.RunAsUser = direct.ValueOf(in.RunAsUser)
	return out
}

func WorkstationConfig_Container_EnvVar_ToProto(mapCtx *direct.MapContext, in []krm.WorkstationConfig_Container_EnvVar) map[string]string {
	if in == nil {
		return nil
	}
	out := make(map[string]string)
	for _, envVar := range in {
		out[envVar.Name] = envVar.Value
	}
	return out
}

func WorkstationConfig_CustomerEncryptionKey_FromProto(mapCtx *direct.MapContext, in *pb.WorkstationConfig_CustomerEncryptionKey) *krm.WorkstationConfig_CustomerEncryptionKey {
	if in == nil {
		return nil
	}
	out := &krm.WorkstationConfig_CustomerEncryptionKey{}
	out.KmsCryptoKeyRef = &refs.KMSCryptoKeyRef{External: in.KmsKey}
	out.ServiceAccountRef = &refs.IAMServiceAccountRef{External: in.KmsKeyServiceAccount}
	return out
}

func WorkstationConfig_CustomerEncryptionKey_ToProto(mapCtx *direct.MapContext, in *krm.WorkstationConfig_CustomerEncryptionKey) *pb.WorkstationConfig_CustomerEncryptionKey {
	if in == nil {
		return nil
	}
	out := &pb.WorkstationConfig_CustomerEncryptionKey{}
	out.KmsKey = in.KmsCryptoKeyRef.External
	out.KmsKeyServiceAccount = in.ServiceAccountRef.External
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
	if in.GetServiceAccount() != "" {
		out.ServiceAccountRef = &refs.IAMServiceAccountRef{External: in.GetServiceAccount()}
	}
	out.ServiceAccountScopes = in.ServiceAccountScopes
	out.Tags = in.Tags
	out.PoolSize = direct.LazyPtr(in.GetPoolSize())
	// MISSING: PooledInstances
	out.DisablePublicIPAddresses = direct.LazyPtr(in.DisablePublicIpAddresses)
	out.EnableNestedVirtualization = direct.LazyPtr(in.GetEnableNestedVirtualization())
	out.ShieldedInstanceConfig = WorkstationConfig_Host_GceInstance_GceShieldedInstanceConfig_FromProto(mapCtx, in.GetShieldedInstanceConfig())
	out.ConfidentialInstanceConfig = WorkstationConfig_Host_GceInstance_GceConfidentialInstanceConfig_FromProto(mapCtx, in.GetConfidentialInstanceConfig())
	out.BootDiskSizeGB = direct.LazyPtr(in.BootDiskSizeGb)
	return out
}

func WorkstationConfig_Host_GceInstance_ToProto(mapCtx *direct.MapContext, in *krm.WorkstationConfig_Host_GceInstance) *pb.WorkstationConfig_Host_GceInstance {
	if in == nil {
		return nil
	}
	out := &pb.WorkstationConfig_Host_GceInstance{}
	out.MachineType = direct.ValueOf(in.MachineType)
	if in.ServiceAccountRef != nil {
		out.ServiceAccount = in.ServiceAccountRef.External
	}
	out.ServiceAccountScopes = in.ServiceAccountScopes
	out.Tags = in.Tags
	out.PoolSize = direct.ValueOf(in.PoolSize)
	// MISSING: PooledInstances
	out.DisablePublicIpAddresses = direct.ValueOf(in.DisablePublicIPAddresses)
	out.EnableNestedVirtualization = direct.ValueOf(in.EnableNestedVirtualization)
	out.ShieldedInstanceConfig = WorkstationConfig_Host_GceInstance_GceShieldedInstanceConfig_ToProto(mapCtx, in.ShieldedInstanceConfig)
	out.ConfidentialInstanceConfig = WorkstationConfig_Host_GceInstance_GceConfidentialInstanceConfig_ToProto(mapCtx, in.ConfidentialInstanceConfig)
	out.BootDiskSizeGb = direct.ValueOf(in.BootDiskSizeGB)
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
	out.EnableVTPM = direct.LazyPtr(in.EnableVtpm)
	out.EnableIntegrityMonitoring = direct.LazyPtr(in.GetEnableIntegrityMonitoring())
	return out
}

func WorkstationConfig_Host_GceInstance_GceShieldedInstanceConfig_ToProto(mapCtx *direct.MapContext, in *krm.WorkstationConfig_Host_GceInstance_GceShieldedInstanceConfig) *pb.WorkstationConfig_Host_GceInstance_GceShieldedInstanceConfig {
	if in == nil {
		return nil
	}
	out := &pb.WorkstationConfig_Host_GceInstance_GceShieldedInstanceConfig{}
	out.EnableSecureBoot = direct.ValueOf(in.EnableSecureBoot)
	out.EnableVtpm = direct.ValueOf(in.EnableVTPM)
	out.EnableIntegrityMonitoring = direct.ValueOf(in.EnableIntegrityMonitoring)
	return out
}

func WorkstationConfig_PersistentDirectory_FromProto(mapCtx *direct.MapContext, in *pb.WorkstationConfig_PersistentDirectory) *krm.WorkstationConfig_PersistentDirectory {
	if in == nil {
		return nil
	}
	out := &krm.WorkstationConfig_PersistentDirectory{}
	out.GcePD = WorkstationConfig_PersistentDirectory_GceRegionalPersistentDisk_FromProto(mapCtx, in.GetGcePd())
	out.MountPath = direct.LazyPtr(in.GetMountPath())
	return out
}

func WorkstationConfig_PersistentDirectory_ToProto(mapCtx *direct.MapContext, in *krm.WorkstationConfig_PersistentDirectory) *pb.WorkstationConfig_PersistentDirectory {
	if in == nil {
		return nil
	}
	out := &pb.WorkstationConfig_PersistentDirectory{}
	if oneof := WorkstationConfig_PersistentDirectory_GceRegionalPersistentDisk_ToProto(mapCtx, in.GcePD); oneof != nil {
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
	out.SizeGB = direct.LazyPtr(in.GetSizeGb())
	out.FSType = direct.LazyPtr(in.GetFsType())
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
	out.SizeGb = direct.ValueOf(in.SizeGB)
	out.FsType = direct.ValueOf(in.FSType)
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
