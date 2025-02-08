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

package vmmigration

import (
	refs "github.com/GoogleCloudPlatform/k8s-config-connector/apis/refs/v1beta1"
	pb "cloud.google.com/go/vmmigration/apiv1/vmmigrationpb"
	krm "github.com/GoogleCloudPlatform/k8s-config-connector/apis/vmmigration/v1alpha1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct"
)
func AppliedLicense_FromProto(mapCtx *direct.MapContext, in *pb.AppliedLicense) *krm.AppliedLicense {
	if in == nil {
		return nil
	}
	out := &krm.AppliedLicense{}
	out.Type = direct.Enum_FromProto(mapCtx, in.GetType())
	out.OsLicense = direct.LazyPtr(in.GetOsLicense())
	return out
}
func AppliedLicense_ToProto(mapCtx *direct.MapContext, in *krm.AppliedLicense) *pb.AppliedLicense {
	if in == nil {
		return nil
	}
	out := &pb.AppliedLicense{}
	out.Type = direct.Enum_ToProto[pb.AppliedLicense_Type](mapCtx, in.Type)
	out.OsLicense = direct.ValueOf(in.OsLicense)
	return out
}
func ComputeEngineTargetDetails_FromProto(mapCtx *direct.MapContext, in *pb.ComputeEngineTargetDetails) *krm.ComputeEngineTargetDetails {
	if in == nil {
		return nil
	}
	out := &krm.ComputeEngineTargetDetails{}
	out.VmName = direct.LazyPtr(in.GetVmName())
	out.Project = direct.LazyPtr(in.GetProject())
	out.Zone = direct.LazyPtr(in.GetZone())
	out.MachineTypeSeries = direct.LazyPtr(in.GetMachineTypeSeries())
	out.MachineType = direct.LazyPtr(in.GetMachineType())
	out.NetworkTags = in.NetworkTags
	out.NetworkInterfaces = direct.Slice_FromProto(mapCtx, in.NetworkInterfaces, NetworkInterface_FromProto)
	out.ServiceAccount = direct.LazyPtr(in.GetServiceAccount())
	out.DiskType = direct.Enum_FromProto(mapCtx, in.GetDiskType())
	out.Labels = in.Labels
	out.LicenseType = direct.Enum_FromProto(mapCtx, in.GetLicenseType())
	out.AppliedLicense = AppliedLicense_FromProto(mapCtx, in.GetAppliedLicense())
	out.ComputeScheduling = ComputeScheduling_FromProto(mapCtx, in.GetComputeScheduling())
	out.SecureBoot = direct.LazyPtr(in.GetSecureBoot())
	out.BootOption = direct.Enum_FromProto(mapCtx, in.GetBootOption())
	out.Metadata = in.Metadata
	out.AdditionalLicenses = in.AdditionalLicenses
	out.Hostname = direct.LazyPtr(in.GetHostname())
	return out
}
func ComputeEngineTargetDetails_ToProto(mapCtx *direct.MapContext, in *krm.ComputeEngineTargetDetails) *pb.ComputeEngineTargetDetails {
	if in == nil {
		return nil
	}
	out := &pb.ComputeEngineTargetDetails{}
	out.VmName = direct.ValueOf(in.VmName)
	out.Project = direct.ValueOf(in.Project)
	out.Zone = direct.ValueOf(in.Zone)
	out.MachineTypeSeries = direct.ValueOf(in.MachineTypeSeries)
	out.MachineType = direct.ValueOf(in.MachineType)
	out.NetworkTags = in.NetworkTags
	out.NetworkInterfaces = direct.Slice_ToProto(mapCtx, in.NetworkInterfaces, NetworkInterface_ToProto)
	out.ServiceAccount = direct.ValueOf(in.ServiceAccount)
	out.DiskType = direct.Enum_ToProto[pb.ComputeEngineDiskType](mapCtx, in.DiskType)
	out.Labels = in.Labels
	out.LicenseType = direct.Enum_ToProto[pb.ComputeEngineLicenseType](mapCtx, in.LicenseType)
	out.AppliedLicense = AppliedLicense_ToProto(mapCtx, in.AppliedLicense)
	out.ComputeScheduling = ComputeScheduling_ToProto(mapCtx, in.ComputeScheduling)
	out.SecureBoot = direct.ValueOf(in.SecureBoot)
	out.BootOption = direct.Enum_ToProto[pb.ComputeEngineBootOption](mapCtx, in.BootOption)
	out.Metadata = in.Metadata
	out.AdditionalLicenses = in.AdditionalLicenses
	out.Hostname = direct.ValueOf(in.Hostname)
	return out
}
func ComputeScheduling_FromProto(mapCtx *direct.MapContext, in *pb.ComputeScheduling) *krm.ComputeScheduling {
	if in == nil {
		return nil
	}
	out := &krm.ComputeScheduling{}
	out.OnHostMaintenance = direct.Enum_FromProto(mapCtx, in.GetOnHostMaintenance())
	out.RestartType = direct.Enum_FromProto(mapCtx, in.GetRestartType())
	out.NodeAffinities = direct.Slice_FromProto(mapCtx, in.NodeAffinities, SchedulingNodeAffinity_FromProto)
	out.MinNodeCpus = direct.LazyPtr(in.GetMinNodeCpus())
	return out
}
func ComputeScheduling_ToProto(mapCtx *direct.MapContext, in *krm.ComputeScheduling) *pb.ComputeScheduling {
	if in == nil {
		return nil
	}
	out := &pb.ComputeScheduling{}
	out.OnHostMaintenance = direct.Enum_ToProto[pb.ComputeScheduling_OnHostMaintenance](mapCtx, in.OnHostMaintenance)
	out.RestartType = direct.Enum_ToProto[pb.ComputeScheduling_RestartType](mapCtx, in.RestartType)
	out.NodeAffinities = direct.Slice_ToProto(mapCtx, in.NodeAffinities, SchedulingNodeAffinity_ToProto)
	out.MinNodeCpus = direct.ValueOf(in.MinNodeCpus)
	return out
}
func CutoverJob_FromProto(mapCtx *direct.MapContext, in *pb.CutoverJob) *krm.CutoverJob {
	if in == nil {
		return nil
	}
	out := &krm.CutoverJob{}
	// MISSING: ComputeEngineTargetDetails
	// MISSING: CreateTime
	// MISSING: EndTime
	// MISSING: Name
	// MISSING: State
	// MISSING: StateTime
	// MISSING: ProgressPercent
	// MISSING: Error
	// MISSING: StateMessage
	// MISSING: Steps
	return out
}
func CutoverJob_ToProto(mapCtx *direct.MapContext, in *krm.CutoverJob) *pb.CutoverJob {
	if in == nil {
		return nil
	}
	out := &pb.CutoverJob{}
	// MISSING: ComputeEngineTargetDetails
	// MISSING: CreateTime
	// MISSING: EndTime
	// MISSING: Name
	// MISSING: State
	// MISSING: StateTime
	// MISSING: ProgressPercent
	// MISSING: Error
	// MISSING: StateMessage
	// MISSING: Steps
	return out
}
func CutoverJobObservedState_FromProto(mapCtx *direct.MapContext, in *pb.CutoverJob) *krm.CutoverJobObservedState {
	if in == nil {
		return nil
	}
	out := &krm.CutoverJobObservedState{}
	out.ComputeEngineTargetDetails = ComputeEngineTargetDetails_FromProto(mapCtx, in.GetComputeEngineTargetDetails())
	out.CreateTime = direct.StringTimestamp_FromProto(mapCtx, in.GetCreateTime())
	out.EndTime = direct.StringTimestamp_FromProto(mapCtx, in.GetEndTime())
	out.Name = direct.LazyPtr(in.GetName())
	out.State = direct.Enum_FromProto(mapCtx, in.GetState())
	out.StateTime = direct.StringTimestamp_FromProto(mapCtx, in.GetStateTime())
	out.ProgressPercent = direct.LazyPtr(in.GetProgressPercent())
	out.Error = Status_FromProto(mapCtx, in.GetError())
	out.StateMessage = direct.LazyPtr(in.GetStateMessage())
	out.Steps = direct.Slice_FromProto(mapCtx, in.Steps, CutoverStep_FromProto)
	return out
}
func CutoverJobObservedState_ToProto(mapCtx *direct.MapContext, in *krm.CutoverJobObservedState) *pb.CutoverJob {
	if in == nil {
		return nil
	}
	out := &pb.CutoverJob{}
	if oneof := ComputeEngineTargetDetails_ToProto(mapCtx, in.ComputeEngineTargetDetails); oneof != nil {
		out.TargetVmDetails = &pb.CutoverJob_ComputeEngineTargetDetails{ComputeEngineTargetDetails: oneof}
	}
	out.CreateTime = direct.StringTimestamp_ToProto(mapCtx, in.CreateTime)
	out.EndTime = direct.StringTimestamp_ToProto(mapCtx, in.EndTime)
	out.Name = direct.ValueOf(in.Name)
	out.State = direct.Enum_ToProto[pb.CutoverJob_State](mapCtx, in.State)
	out.StateTime = direct.StringTimestamp_ToProto(mapCtx, in.StateTime)
	out.ProgressPercent = direct.ValueOf(in.ProgressPercent)
	out.Error = Status_ToProto(mapCtx, in.Error)
	out.StateMessage = direct.ValueOf(in.StateMessage)
	out.Steps = direct.Slice_ToProto(mapCtx, in.Steps, CutoverStep_ToProto)
	return out
}
func CutoverStep_FromProto(mapCtx *direct.MapContext, in *pb.CutoverStep) *krm.CutoverStep {
	if in == nil {
		return nil
	}
	out := &krm.CutoverStep{}
	out.PreviousReplicationCycle = ReplicationCycle_FromProto(mapCtx, in.GetPreviousReplicationCycle())
	out.ShuttingDownSourceVm = ShuttingDownSourceVMStep_FromProto(mapCtx, in.GetShuttingDownSourceVm())
	out.FinalSync = ReplicationCycle_FromProto(mapCtx, in.GetFinalSync())
	out.PreparingVmDisks = PreparingVMDisksStep_FromProto(mapCtx, in.GetPreparingVmDisks())
	out.InstantiatingMigratedVm = InstantiatingMigratedVMStep_FromProto(mapCtx, in.GetInstantiatingMigratedVm())
	out.StartTime = direct.StringTimestamp_FromProto(mapCtx, in.GetStartTime())
	out.EndTime = direct.StringTimestamp_FromProto(mapCtx, in.GetEndTime())
	return out
}
func CutoverStep_ToProto(mapCtx *direct.MapContext, in *krm.CutoverStep) *pb.CutoverStep {
	if in == nil {
		return nil
	}
	out := &pb.CutoverStep{}
	if oneof := ReplicationCycle_ToProto(mapCtx, in.PreviousReplicationCycle); oneof != nil {
		out.Step = &pb.CutoverStep_PreviousReplicationCycle{PreviousReplicationCycle: oneof}
	}
	if oneof := ShuttingDownSourceVMStep_ToProto(mapCtx, in.ShuttingDownSourceVm); oneof != nil {
		out.Step = &pb.CutoverStep_ShuttingDownSourceVm{ShuttingDownSourceVm: oneof}
	}
	if oneof := ReplicationCycle_ToProto(mapCtx, in.FinalSync); oneof != nil {
		out.Step = &pb.CutoverStep_FinalSync{FinalSync: oneof}
	}
	if oneof := PreparingVMDisksStep_ToProto(mapCtx, in.PreparingVmDisks); oneof != nil {
		out.Step = &pb.CutoverStep_PreparingVmDisks{PreparingVmDisks: oneof}
	}
	if oneof := InstantiatingMigratedVMStep_ToProto(mapCtx, in.InstantiatingMigratedVm); oneof != nil {
		out.Step = &pb.CutoverStep_InstantiatingMigratedVm{InstantiatingMigratedVm: oneof}
	}
	out.StartTime = direct.StringTimestamp_ToProto(mapCtx, in.StartTime)
	out.EndTime = direct.StringTimestamp_ToProto(mapCtx, in.EndTime)
	return out
}
func CycleStep_FromProto(mapCtx *direct.MapContext, in *pb.CycleStep) *krm.CycleStep {
	if in == nil {
		return nil
	}
	out := &krm.CycleStep{}
	out.InitializingReplication = InitializingReplicationStep_FromProto(mapCtx, in.GetInitializingReplication())
	out.Replicating = ReplicatingStep_FromProto(mapCtx, in.GetReplicating())
	out.PostProcessing = PostProcessingStep_FromProto(mapCtx, in.GetPostProcessing())
	out.StartTime = direct.StringTimestamp_FromProto(mapCtx, in.GetStartTime())
	out.EndTime = direct.StringTimestamp_FromProto(mapCtx, in.GetEndTime())
	return out
}
func CycleStep_ToProto(mapCtx *direct.MapContext, in *krm.CycleStep) *pb.CycleStep {
	if in == nil {
		return nil
	}
	out := &pb.CycleStep{}
	if oneof := InitializingReplicationStep_ToProto(mapCtx, in.InitializingReplication); oneof != nil {
		out.Step = &pb.CycleStep_InitializingReplication{InitializingReplication: oneof}
	}
	if oneof := ReplicatingStep_ToProto(mapCtx, in.Replicating); oneof != nil {
		out.Step = &pb.CycleStep_Replicating{Replicating: oneof}
	}
	if oneof := PostProcessingStep_ToProto(mapCtx, in.PostProcessing); oneof != nil {
		out.Step = &pb.CycleStep_PostProcessing{PostProcessing: oneof}
	}
	out.StartTime = direct.StringTimestamp_ToProto(mapCtx, in.StartTime)
	out.EndTime = direct.StringTimestamp_ToProto(mapCtx, in.EndTime)
	return out
}
func InitializingReplicationStep_FromProto(mapCtx *direct.MapContext, in *pb.InitializingReplicationStep) *krm.InitializingReplicationStep {
	if in == nil {
		return nil
	}
	out := &krm.InitializingReplicationStep{}
	return out
}
func InitializingReplicationStep_ToProto(mapCtx *direct.MapContext, in *krm.InitializingReplicationStep) *pb.InitializingReplicationStep {
	if in == nil {
		return nil
	}
	out := &pb.InitializingReplicationStep{}
	return out
}
func InstantiatingMigratedVMStep_FromProto(mapCtx *direct.MapContext, in *pb.InstantiatingMigratedVMStep) *krm.InstantiatingMigratedVMStep {
	if in == nil {
		return nil
	}
	out := &krm.InstantiatingMigratedVMStep{}
	return out
}
func InstantiatingMigratedVMStep_ToProto(mapCtx *direct.MapContext, in *krm.InstantiatingMigratedVMStep) *pb.InstantiatingMigratedVMStep {
	if in == nil {
		return nil
	}
	out := &pb.InstantiatingMigratedVMStep{}
	return out
}
func NetworkInterface_FromProto(mapCtx *direct.MapContext, in *pb.NetworkInterface) *krm.NetworkInterface {
	if in == nil {
		return nil
	}
	out := &krm.NetworkInterface{}
	out.Network = direct.LazyPtr(in.GetNetwork())
	out.Subnetwork = direct.LazyPtr(in.GetSubnetwork())
	out.InternalIP = direct.LazyPtr(in.GetInternalIp())
	out.ExternalIP = direct.LazyPtr(in.GetExternalIp())
	return out
}
func NetworkInterface_ToProto(mapCtx *direct.MapContext, in *krm.NetworkInterface) *pb.NetworkInterface {
	if in == nil {
		return nil
	}
	out := &pb.NetworkInterface{}
	out.Network = direct.ValueOf(in.Network)
	out.Subnetwork = direct.ValueOf(in.Subnetwork)
	out.InternalIp = direct.ValueOf(in.InternalIP)
	out.ExternalIp = direct.ValueOf(in.ExternalIP)
	return out
}
func PostProcessingStep_FromProto(mapCtx *direct.MapContext, in *pb.PostProcessingStep) *krm.PostProcessingStep {
	if in == nil {
		return nil
	}
	out := &krm.PostProcessingStep{}
	return out
}
func PostProcessingStep_ToProto(mapCtx *direct.MapContext, in *krm.PostProcessingStep) *pb.PostProcessingStep {
	if in == nil {
		return nil
	}
	out := &pb.PostProcessingStep{}
	return out
}
func PreparingVMDisksStep_FromProto(mapCtx *direct.MapContext, in *pb.PreparingVMDisksStep) *krm.PreparingVMDisksStep {
	if in == nil {
		return nil
	}
	out := &krm.PreparingVMDisksStep{}
	return out
}
func PreparingVMDisksStep_ToProto(mapCtx *direct.MapContext, in *krm.PreparingVMDisksStep) *pb.PreparingVMDisksStep {
	if in == nil {
		return nil
	}
	out := &pb.PreparingVMDisksStep{}
	return out
}
func ReplicatingStep_FromProto(mapCtx *direct.MapContext, in *pb.ReplicatingStep) *krm.ReplicatingStep {
	if in == nil {
		return nil
	}
	out := &krm.ReplicatingStep{}
	out.TotalBytes = direct.LazyPtr(in.GetTotalBytes())
	out.ReplicatedBytes = direct.LazyPtr(in.GetReplicatedBytes())
	out.LastTwoMinutesAverageBytesPerSecond = direct.LazyPtr(in.GetLastTwoMinutesAverageBytesPerSecond())
	out.LastThirtyMinutesAverageBytesPerSecond = direct.LazyPtr(in.GetLastThirtyMinutesAverageBytesPerSecond())
	return out
}
func ReplicatingStep_ToProto(mapCtx *direct.MapContext, in *krm.ReplicatingStep) *pb.ReplicatingStep {
	if in == nil {
		return nil
	}
	out := &pb.ReplicatingStep{}
	out.TotalBytes = direct.ValueOf(in.TotalBytes)
	out.ReplicatedBytes = direct.ValueOf(in.ReplicatedBytes)
	out.LastTwoMinutesAverageBytesPerSecond = direct.ValueOf(in.LastTwoMinutesAverageBytesPerSecond)
	out.LastThirtyMinutesAverageBytesPerSecond = direct.ValueOf(in.LastThirtyMinutesAverageBytesPerSecond)
	return out
}
func ReplicationCycle_FromProto(mapCtx *direct.MapContext, in *pb.ReplicationCycle) *krm.ReplicationCycle {
	if in == nil {
		return nil
	}
	out := &krm.ReplicationCycle{}
	out.Name = direct.LazyPtr(in.GetName())
	out.CycleNumber = direct.LazyPtr(in.GetCycleNumber())
	out.StartTime = direct.StringTimestamp_FromProto(mapCtx, in.GetStartTime())
	out.EndTime = direct.StringTimestamp_FromProto(mapCtx, in.GetEndTime())
	out.TotalPauseDuration = direct.StringDuration_FromProto(mapCtx, in.GetTotalPauseDuration())
	out.ProgressPercent = direct.LazyPtr(in.GetProgressPercent())
	out.Steps = direct.Slice_FromProto(mapCtx, in.Steps, CycleStep_FromProto)
	out.State = direct.Enum_FromProto(mapCtx, in.GetState())
	out.Error = Status_FromProto(mapCtx, in.GetError())
	return out
}
func ReplicationCycle_ToProto(mapCtx *direct.MapContext, in *krm.ReplicationCycle) *pb.ReplicationCycle {
	if in == nil {
		return nil
	}
	out := &pb.ReplicationCycle{}
	out.Name = direct.ValueOf(in.Name)
	out.CycleNumber = direct.ValueOf(in.CycleNumber)
	out.StartTime = direct.StringTimestamp_ToProto(mapCtx, in.StartTime)
	out.EndTime = direct.StringTimestamp_ToProto(mapCtx, in.EndTime)
	out.TotalPauseDuration = direct.StringDuration_ToProto(mapCtx, in.TotalPauseDuration)
	out.ProgressPercent = direct.ValueOf(in.ProgressPercent)
	out.Steps = direct.Slice_ToProto(mapCtx, in.Steps, CycleStep_ToProto)
	out.State = direct.Enum_ToProto[pb.ReplicationCycle_State](mapCtx, in.State)
	out.Error = Status_ToProto(mapCtx, in.Error)
	return out
}
func SchedulingNodeAffinity_FromProto(mapCtx *direct.MapContext, in *pb.SchedulingNodeAffinity) *krm.SchedulingNodeAffinity {
	if in == nil {
		return nil
	}
	out := &krm.SchedulingNodeAffinity{}
	out.Key = direct.LazyPtr(in.GetKey())
	out.Operator = direct.Enum_FromProto(mapCtx, in.GetOperator())
	out.Values = in.Values
	return out
}
func SchedulingNodeAffinity_ToProto(mapCtx *direct.MapContext, in *krm.SchedulingNodeAffinity) *pb.SchedulingNodeAffinity {
	if in == nil {
		return nil
	}
	out := &pb.SchedulingNodeAffinity{}
	out.Key = direct.ValueOf(in.Key)
	out.Operator = direct.Enum_ToProto[pb.SchedulingNodeAffinity_Operator](mapCtx, in.Operator)
	out.Values = in.Values
	return out
}
func ShuttingDownSourceVMStep_FromProto(mapCtx *direct.MapContext, in *pb.ShuttingDownSourceVMStep) *krm.ShuttingDownSourceVMStep {
	if in == nil {
		return nil
	}
	out := &krm.ShuttingDownSourceVMStep{}
	return out
}
func ShuttingDownSourceVMStep_ToProto(mapCtx *direct.MapContext, in *krm.ShuttingDownSourceVMStep) *pb.ShuttingDownSourceVMStep {
	if in == nil {
		return nil
	}
	out := &pb.ShuttingDownSourceVMStep{}
	return out
}
func VmmigrationCutoverJobObservedState_FromProto(mapCtx *direct.MapContext, in *pb.CutoverJob) *krm.VmmigrationCutoverJobObservedState {
	if in == nil {
		return nil
	}
	out := &krm.VmmigrationCutoverJobObservedState{}
	// MISSING: ComputeEngineTargetDetails
	// MISSING: CreateTime
	// MISSING: EndTime
	// MISSING: Name
	// MISSING: State
	// MISSING: StateTime
	// MISSING: ProgressPercent
	// MISSING: Error
	// MISSING: StateMessage
	// MISSING: Steps
	return out
}
func VmmigrationCutoverJobObservedState_ToProto(mapCtx *direct.MapContext, in *krm.VmmigrationCutoverJobObservedState) *pb.CutoverJob {
	if in == nil {
		return nil
	}
	out := &pb.CutoverJob{}
	// MISSING: ComputeEngineTargetDetails
	// MISSING: CreateTime
	// MISSING: EndTime
	// MISSING: Name
	// MISSING: State
	// MISSING: StateTime
	// MISSING: ProgressPercent
	// MISSING: Error
	// MISSING: StateMessage
	// MISSING: Steps
	return out
}
func VmmigrationCutoverJobSpec_FromProto(mapCtx *direct.MapContext, in *pb.CutoverJob) *krm.VmmigrationCutoverJobSpec {
	if in == nil {
		return nil
	}
	out := &krm.VmmigrationCutoverJobSpec{}
	// MISSING: ComputeEngineTargetDetails
	// MISSING: CreateTime
	// MISSING: EndTime
	// MISSING: Name
	// MISSING: State
	// MISSING: StateTime
	// MISSING: ProgressPercent
	// MISSING: Error
	// MISSING: StateMessage
	// MISSING: Steps
	return out
}
func VmmigrationCutoverJobSpec_ToProto(mapCtx *direct.MapContext, in *krm.VmmigrationCutoverJobSpec) *pb.CutoverJob {
	if in == nil {
		return nil
	}
	out := &pb.CutoverJob{}
	// MISSING: ComputeEngineTargetDetails
	// MISSING: CreateTime
	// MISSING: EndTime
	// MISSING: Name
	// MISSING: State
	// MISSING: StateTime
	// MISSING: ProgressPercent
	// MISSING: Error
	// MISSING: StateMessage
	// MISSING: Steps
	return out
}
