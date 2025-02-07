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

package batch

import (
	pb "cloud.google.com/go/batch/apiv1/batchpb"
	krm "github.com/GoogleCloudPlatform/k8s-config-connector/apis/batch/v1alpha1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct"
	refs "github.com/GoogleCloudPlatform/k8s-config-connector/apis/refs/v1beta1"
)
func BatchJobObservedState_FromProto(mapCtx *direct.MapContext, in *pb.Job) *krm.BatchJobObservedState {
	if in == nil {
		return nil
	}
	out := &krm.BatchJobObservedState{}
	// MISSING: Name
	// MISSING: Uid
	// MISSING: Priority
	// MISSING: TaskGroups
	// MISSING: AllocationPolicy
	// MISSING: Labels
	// MISSING: Status
	// MISSING: CreateTime
	// MISSING: UpdateTime
	// MISSING: LogsPolicy
	// MISSING: Notifications
	return out
}
func BatchJobObservedState_ToProto(mapCtx *direct.MapContext, in *krm.BatchJobObservedState) *pb.Job {
	if in == nil {
		return nil
	}
	out := &pb.Job{}
	// MISSING: Name
	// MISSING: Uid
	// MISSING: Priority
	// MISSING: TaskGroups
	// MISSING: AllocationPolicy
	// MISSING: Labels
	// MISSING: Status
	// MISSING: CreateTime
	// MISSING: UpdateTime
	// MISSING: LogsPolicy
	// MISSING: Notifications
	return out
}
func BatchJobSpec_FromProto(mapCtx *direct.MapContext, in *pb.Job) *krm.BatchJobSpec {
	if in == nil {
		return nil
	}
	out := &krm.BatchJobSpec{}
	// MISSING: Name
	// MISSING: Uid
	// MISSING: Priority
	// MISSING: TaskGroups
	// MISSING: AllocationPolicy
	// MISSING: Labels
	// MISSING: Status
	// MISSING: CreateTime
	// MISSING: UpdateTime
	// MISSING: LogsPolicy
	// MISSING: Notifications
	return out
}
func BatchJobSpec_ToProto(mapCtx *direct.MapContext, in *krm.BatchJobSpec) *pb.Job {
	if in == nil {
		return nil
	}
	out := &pb.Job{}
	// MISSING: Name
	// MISSING: Uid
	// MISSING: Priority
	// MISSING: TaskGroups
	// MISSING: AllocationPolicy
	// MISSING: Labels
	// MISSING: Status
	// MISSING: CreateTime
	// MISSING: UpdateTime
	// MISSING: LogsPolicy
	// MISSING: Notifications
	return out
}
func BatchTaskGroupObservedState_FromProto(mapCtx *direct.MapContext, in *pb.TaskGroup) *krm.BatchTaskGroupObservedState {
	if in == nil {
		return nil
	}
	out := &krm.BatchTaskGroupObservedState{}
	// MISSING: Name
	// MISSING: TaskSpec
	// MISSING: TaskCount
	// MISSING: Parallelism
	// MISSING: SchedulingPolicy
	// MISSING: TaskEnvironments
	// MISSING: TaskCountPerNode
	// MISSING: RequireHostsFile
	// MISSING: PermissiveSSH
	// MISSING: RunAsNonRoot
	return out
}
func BatchTaskGroupObservedState_ToProto(mapCtx *direct.MapContext, in *krm.BatchTaskGroupObservedState) *pb.TaskGroup {
	if in == nil {
		return nil
	}
	out := &pb.TaskGroup{}
	// MISSING: Name
	// MISSING: TaskSpec
	// MISSING: TaskCount
	// MISSING: Parallelism
	// MISSING: SchedulingPolicy
	// MISSING: TaskEnvironments
	// MISSING: TaskCountPerNode
	// MISSING: RequireHostsFile
	// MISSING: PermissiveSSH
	// MISSING: RunAsNonRoot
	return out
}
func BatchTaskGroupSpec_FromProto(mapCtx *direct.MapContext, in *pb.TaskGroup) *krm.BatchTaskGroupSpec {
	if in == nil {
		return nil
	}
	out := &krm.BatchTaskGroupSpec{}
	// MISSING: Name
	// MISSING: TaskSpec
	// MISSING: TaskCount
	// MISSING: Parallelism
	// MISSING: SchedulingPolicy
	// MISSING: TaskEnvironments
	// MISSING: TaskCountPerNode
	// MISSING: RequireHostsFile
	// MISSING: PermissiveSSH
	// MISSING: RunAsNonRoot
	return out
}
func BatchTaskGroupSpec_ToProto(mapCtx *direct.MapContext, in *krm.BatchTaskGroupSpec) *pb.TaskGroup {
	if in == nil {
		return nil
	}
	out := &pb.TaskGroup{}
	// MISSING: Name
	// MISSING: TaskSpec
	// MISSING: TaskCount
	// MISSING: Parallelism
	// MISSING: SchedulingPolicy
	// MISSING: TaskEnvironments
	// MISSING: TaskCountPerNode
	// MISSING: RequireHostsFile
	// MISSING: PermissiveSSH
	// MISSING: RunAsNonRoot
	return out
}
func ComputeResource_FromProto(mapCtx *direct.MapContext, in *pb.ComputeResource) *krm.ComputeResource {
	if in == nil {
		return nil
	}
	out := &krm.ComputeResource{}
	out.CpuMilli = direct.LazyPtr(in.GetCpuMilli())
	out.MemoryMib = direct.LazyPtr(in.GetMemoryMib())
	out.BootDiskMib = direct.LazyPtr(in.GetBootDiskMib())
	return out
}
func ComputeResource_ToProto(mapCtx *direct.MapContext, in *krm.ComputeResource) *pb.ComputeResource {
	if in == nil {
		return nil
	}
	out := &pb.ComputeResource{}
	out.CpuMilli = direct.ValueOf(in.CpuMilli)
	out.MemoryMib = direct.ValueOf(in.MemoryMib)
	out.BootDiskMib = direct.ValueOf(in.BootDiskMib)
	return out
}
func Environment_FromProto(mapCtx *direct.MapContext, in *pb.Environment) *krm.Environment {
	if in == nil {
		return nil
	}
	out := &krm.Environment{}
	out.Variables = in.Variables
	out.SecretVariables = in.SecretVariables
	out.EncryptedVariables = Environment_KMSEnvMap_FromProto(mapCtx, in.GetEncryptedVariables())
	return out
}
func Environment_ToProto(mapCtx *direct.MapContext, in *krm.Environment) *pb.Environment {
	if in == nil {
		return nil
	}
	out := &pb.Environment{}
	out.Variables = in.Variables
	out.SecretVariables = in.SecretVariables
	out.EncryptedVariables = Environment_KMSEnvMap_ToProto(mapCtx, in.EncryptedVariables)
	return out
}
func Environment_KMSEnvMap_FromProto(mapCtx *direct.MapContext, in *pb.Environment_KMSEnvMap) *krm.Environment_KMSEnvMap {
	if in == nil {
		return nil
	}
	out := &krm.Environment_KMSEnvMap{}
	out.KeyName = direct.LazyPtr(in.GetKeyName())
	out.CipherText = direct.LazyPtr(in.GetCipherText())
	return out
}
func Environment_KMSEnvMap_ToProto(mapCtx *direct.MapContext, in *krm.Environment_KMSEnvMap) *pb.Environment_KMSEnvMap {
	if in == nil {
		return nil
	}
	out := &pb.Environment_KMSEnvMap{}
	out.KeyName = direct.ValueOf(in.KeyName)
	out.CipherText = direct.ValueOf(in.CipherText)
	return out
}
func GCS_FromProto(mapCtx *direct.MapContext, in *pb.GCS) *krm.GCS {
	if in == nil {
		return nil
	}
	out := &krm.GCS{}
	out.RemotePath = direct.LazyPtr(in.GetRemotePath())
	return out
}
func GCS_ToProto(mapCtx *direct.MapContext, in *krm.GCS) *pb.GCS {
	if in == nil {
		return nil
	}
	out := &pb.GCS{}
	out.RemotePath = direct.ValueOf(in.RemotePath)
	return out
}
func LifecyclePolicy_FromProto(mapCtx *direct.MapContext, in *pb.LifecyclePolicy) *krm.LifecyclePolicy {
	if in == nil {
		return nil
	}
	out := &krm.LifecyclePolicy{}
	out.Action = direct.Enum_FromProto(mapCtx, in.GetAction())
	out.ActionCondition = LifecyclePolicy_ActionCondition_FromProto(mapCtx, in.GetActionCondition())
	return out
}
func LifecyclePolicy_ToProto(mapCtx *direct.MapContext, in *krm.LifecyclePolicy) *pb.LifecyclePolicy {
	if in == nil {
		return nil
	}
	out := &pb.LifecyclePolicy{}
	out.Action = direct.Enum_ToProto[pb.LifecyclePolicy_Action](mapCtx, in.Action)
	out.ActionCondition = LifecyclePolicy_ActionCondition_ToProto(mapCtx, in.ActionCondition)
	return out
}
func LifecyclePolicy_ActionCondition_FromProto(mapCtx *direct.MapContext, in *pb.LifecyclePolicy_ActionCondition) *krm.LifecyclePolicy_ActionCondition {
	if in == nil {
		return nil
	}
	out := &krm.LifecyclePolicy_ActionCondition{}
	out.ExitCodes = in.ExitCodes
	return out
}
func LifecyclePolicy_ActionCondition_ToProto(mapCtx *direct.MapContext, in *krm.LifecyclePolicy_ActionCondition) *pb.LifecyclePolicy_ActionCondition {
	if in == nil {
		return nil
	}
	out := &pb.LifecyclePolicy_ActionCondition{}
	out.ExitCodes = in.ExitCodes
	return out
}
func NFS_FromProto(mapCtx *direct.MapContext, in *pb.NFS) *krm.NFS {
	if in == nil {
		return nil
	}
	out := &krm.NFS{}
	out.Server = direct.LazyPtr(in.GetServer())
	out.RemotePath = direct.LazyPtr(in.GetRemotePath())
	return out
}
func NFS_ToProto(mapCtx *direct.MapContext, in *krm.NFS) *pb.NFS {
	if in == nil {
		return nil
	}
	out := &pb.NFS{}
	out.Server = direct.ValueOf(in.Server)
	out.RemotePath = direct.ValueOf(in.RemotePath)
	return out
}
func Runnable_FromProto(mapCtx *direct.MapContext, in *pb.Runnable) *krm.Runnable {
	if in == nil {
		return nil
	}
	out := &krm.Runnable{}
	out.Container = Runnable_Container_FromProto(mapCtx, in.GetContainer())
	out.Script = Runnable_Script_FromProto(mapCtx, in.GetScript())
	out.Barrier = Runnable_Barrier_FromProto(mapCtx, in.GetBarrier())
	out.DisplayName = direct.LazyPtr(in.GetDisplayName())
	out.IgnoreExitStatus = direct.LazyPtr(in.GetIgnoreExitStatus())
	out.Background = direct.LazyPtr(in.GetBackground())
	out.AlwaysRun = direct.LazyPtr(in.GetAlwaysRun())
	out.Environment = Environment_FromProto(mapCtx, in.GetEnvironment())
	out.Timeout = direct.StringDuration_FromProto(mapCtx, in.GetTimeout())
	out.Labels = in.Labels
	return out
}
func Runnable_ToProto(mapCtx *direct.MapContext, in *krm.Runnable) *pb.Runnable {
	if in == nil {
		return nil
	}
	out := &pb.Runnable{}
	if oneof := Runnable_Container_ToProto(mapCtx, in.Container); oneof != nil {
		out.Executable = &pb.Runnable_Container_{Container: oneof}
	}
	if oneof := Runnable_Script_ToProto(mapCtx, in.Script); oneof != nil {
		out.Executable = &pb.Runnable_Script_{Script: oneof}
	}
	if oneof := Runnable_Barrier_ToProto(mapCtx, in.Barrier); oneof != nil {
		out.Executable = &pb.Runnable_Barrier_{Barrier: oneof}
	}
	out.DisplayName = direct.ValueOf(in.DisplayName)
	out.IgnoreExitStatus = direct.ValueOf(in.IgnoreExitStatus)
	out.Background = direct.ValueOf(in.Background)
	out.AlwaysRun = direct.ValueOf(in.AlwaysRun)
	out.Environment = Environment_ToProto(mapCtx, in.Environment)
	out.Timeout = direct.StringDuration_ToProto(mapCtx, in.Timeout)
	out.Labels = in.Labels
	return out
}
func Runnable_Barrier_FromProto(mapCtx *direct.MapContext, in *pb.Runnable_Barrier) *krm.Runnable_Barrier {
	if in == nil {
		return nil
	}
	out := &krm.Runnable_Barrier{}
	out.Name = direct.LazyPtr(in.GetName())
	return out
}
func Runnable_Barrier_ToProto(mapCtx *direct.MapContext, in *krm.Runnable_Barrier) *pb.Runnable_Barrier {
	if in == nil {
		return nil
	}
	out := &pb.Runnable_Barrier{}
	out.Name = direct.ValueOf(in.Name)
	return out
}
func Runnable_Container_FromProto(mapCtx *direct.MapContext, in *pb.Runnable_Container) *krm.Runnable_Container {
	if in == nil {
		return nil
	}
	out := &krm.Runnable_Container{}
	out.ImageURI = direct.LazyPtr(in.GetImageUri())
	out.Commands = in.Commands
	out.Entrypoint = direct.LazyPtr(in.GetEntrypoint())
	out.Volumes = in.Volumes
	out.Options = direct.LazyPtr(in.GetOptions())
	out.BlockExternalNetwork = direct.LazyPtr(in.GetBlockExternalNetwork())
	out.Username = direct.LazyPtr(in.GetUsername())
	out.Password = direct.LazyPtr(in.GetPassword())
	out.EnableImageStreaming = direct.LazyPtr(in.GetEnableImageStreaming())
	return out
}
func Runnable_Container_ToProto(mapCtx *direct.MapContext, in *krm.Runnable_Container) *pb.Runnable_Container {
	if in == nil {
		return nil
	}
	out := &pb.Runnable_Container{}
	out.ImageUri = direct.ValueOf(in.ImageURI)
	out.Commands = in.Commands
	out.Entrypoint = direct.ValueOf(in.Entrypoint)
	out.Volumes = in.Volumes
	out.Options = direct.ValueOf(in.Options)
	out.BlockExternalNetwork = direct.ValueOf(in.BlockExternalNetwork)
	out.Username = direct.ValueOf(in.Username)
	out.Password = direct.ValueOf(in.Password)
	out.EnableImageStreaming = direct.ValueOf(in.EnableImageStreaming)
	return out
}
func Runnable_Script_FromProto(mapCtx *direct.MapContext, in *pb.Runnable_Script) *krm.Runnable_Script {
	if in == nil {
		return nil
	}
	out := &krm.Runnable_Script{}
	out.Path = direct.LazyPtr(in.GetPath())
	out.Text = direct.LazyPtr(in.GetText())
	return out
}
func Runnable_Script_ToProto(mapCtx *direct.MapContext, in *krm.Runnable_Script) *pb.Runnable_Script {
	if in == nil {
		return nil
	}
	out := &pb.Runnable_Script{}
	if oneof := Runnable_Script_Path_ToProto(mapCtx, in.Path); oneof != nil {
		out.Command = oneof
	}
	if oneof := Runnable_Script_Text_ToProto(mapCtx, in.Text); oneof != nil {
		out.Command = oneof
	}
	return out
}
func TaskGroup_FromProto(mapCtx *direct.MapContext, in *pb.TaskGroup) *krm.TaskGroup {
	if in == nil {
		return nil
	}
	out := &krm.TaskGroup{}
	// MISSING: Name
	out.TaskSpec = TaskSpec_FromProto(mapCtx, in.GetTaskSpec())
	out.TaskCount = direct.LazyPtr(in.GetTaskCount())
	out.Parallelism = direct.LazyPtr(in.GetParallelism())
	out.SchedulingPolicy = direct.Enum_FromProto(mapCtx, in.GetSchedulingPolicy())
	out.TaskEnvironments = direct.Slice_FromProto(mapCtx, in.TaskEnvironments, Environment_FromProto)
	out.TaskCountPerNode = direct.LazyPtr(in.GetTaskCountPerNode())
	out.RequireHostsFile = direct.LazyPtr(in.GetRequireHostsFile())
	out.PermissiveSSH = direct.LazyPtr(in.GetPermissiveSsh())
	out.RunAsNonRoot = direct.LazyPtr(in.GetRunAsNonRoot())
	return out
}
func TaskGroup_ToProto(mapCtx *direct.MapContext, in *krm.TaskGroup) *pb.TaskGroup {
	if in == nil {
		return nil
	}
	out := &pb.TaskGroup{}
	// MISSING: Name
	out.TaskSpec = TaskSpec_ToProto(mapCtx, in.TaskSpec)
	out.TaskCount = direct.ValueOf(in.TaskCount)
	out.Parallelism = direct.ValueOf(in.Parallelism)
	out.SchedulingPolicy = direct.Enum_ToProto[pb.TaskGroup_SchedulingPolicy](mapCtx, in.SchedulingPolicy)
	out.TaskEnvironments = direct.Slice_ToProto(mapCtx, in.TaskEnvironments, Environment_ToProto)
	out.TaskCountPerNode = direct.ValueOf(in.TaskCountPerNode)
	out.RequireHostsFile = direct.ValueOf(in.RequireHostsFile)
	out.PermissiveSsh = direct.ValueOf(in.PermissiveSSH)
	out.RunAsNonRoot = direct.ValueOf(in.RunAsNonRoot)
	return out
}
func TaskGroupObservedState_FromProto(mapCtx *direct.MapContext, in *pb.TaskGroup) *krm.TaskGroupObservedState {
	if in == nil {
		return nil
	}
	out := &krm.TaskGroupObservedState{}
	out.Name = direct.LazyPtr(in.GetName())
	// MISSING: TaskSpec
	// MISSING: TaskCount
	// MISSING: Parallelism
	// MISSING: SchedulingPolicy
	// MISSING: TaskEnvironments
	// MISSING: TaskCountPerNode
	// MISSING: RequireHostsFile
	// MISSING: PermissiveSSH
	// MISSING: RunAsNonRoot
	return out
}
func TaskGroupObservedState_ToProto(mapCtx *direct.MapContext, in *krm.TaskGroupObservedState) *pb.TaskGroup {
	if in == nil {
		return nil
	}
	out := &pb.TaskGroup{}
	out.Name = direct.ValueOf(in.Name)
	// MISSING: TaskSpec
	// MISSING: TaskCount
	// MISSING: Parallelism
	// MISSING: SchedulingPolicy
	// MISSING: TaskEnvironments
	// MISSING: TaskCountPerNode
	// MISSING: RequireHostsFile
	// MISSING: PermissiveSSH
	// MISSING: RunAsNonRoot
	return out
}
func TaskSpec_FromProto(mapCtx *direct.MapContext, in *pb.TaskSpec) *krm.TaskSpec {
	if in == nil {
		return nil
	}
	out := &krm.TaskSpec{}
	out.Runnables = direct.Slice_FromProto(mapCtx, in.Runnables, Runnable_FromProto)
	out.ComputeResource = ComputeResource_FromProto(mapCtx, in.GetComputeResource())
	out.MaxRunDuration = direct.StringDuration_FromProto(mapCtx, in.GetMaxRunDuration())
	out.MaxRetryCount = direct.LazyPtr(in.GetMaxRetryCount())
	out.LifecyclePolicies = direct.Slice_FromProto(mapCtx, in.LifecyclePolicies, LifecyclePolicy_FromProto)
	out.Environments = in.Environments
	out.Volumes = direct.Slice_FromProto(mapCtx, in.Volumes, Volume_FromProto)
	out.Environment = Environment_FromProto(mapCtx, in.GetEnvironment())
	return out
}
func TaskSpec_ToProto(mapCtx *direct.MapContext, in *krm.TaskSpec) *pb.TaskSpec {
	if in == nil {
		return nil
	}
	out := &pb.TaskSpec{}
	out.Runnables = direct.Slice_ToProto(mapCtx, in.Runnables, Runnable_ToProto)
	out.ComputeResource = ComputeResource_ToProto(mapCtx, in.ComputeResource)
	out.MaxRunDuration = direct.StringDuration_ToProto(mapCtx, in.MaxRunDuration)
	out.MaxRetryCount = direct.ValueOf(in.MaxRetryCount)
	out.LifecyclePolicies = direct.Slice_ToProto(mapCtx, in.LifecyclePolicies, LifecyclePolicy_ToProto)
	out.Environments = in.Environments
	out.Volumes = direct.Slice_ToProto(mapCtx, in.Volumes, Volume_ToProto)
	out.Environment = Environment_ToProto(mapCtx, in.Environment)
	return out
}
func Volume_FromProto(mapCtx *direct.MapContext, in *pb.Volume) *krm.Volume {
	if in == nil {
		return nil
	}
	out := &krm.Volume{}
	out.Nfs = NFS_FromProto(mapCtx, in.GetNfs())
	out.Gcs = GCS_FromProto(mapCtx, in.GetGcs())
	out.DeviceName = direct.LazyPtr(in.GetDeviceName())
	out.MountPath = direct.LazyPtr(in.GetMountPath())
	out.MountOptions = in.MountOptions
	return out
}
func Volume_ToProto(mapCtx *direct.MapContext, in *krm.Volume) *pb.Volume {
	if in == nil {
		return nil
	}
	out := &pb.Volume{}
	if oneof := NFS_ToProto(mapCtx, in.Nfs); oneof != nil {
		out.Source = &pb.Volume_Nfs{Nfs: oneof}
	}
	if oneof := GCS_ToProto(mapCtx, in.Gcs); oneof != nil {
		out.Source = &pb.Volume_Gcs{Gcs: oneof}
	}
	if oneof := Volume_DeviceName_ToProto(mapCtx, in.DeviceName); oneof != nil {
		out.Source = oneof
	}
	out.MountPath = direct.ValueOf(in.MountPath)
	out.MountOptions = in.MountOptions
	return out
}
