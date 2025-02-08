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
func AdaptingOSStep_FromProto(mapCtx *direct.MapContext, in *pb.AdaptingOSStep) *krm.AdaptingOSStep {
	if in == nil {
		return nil
	}
	out := &krm.AdaptingOSStep{}
	return out
}
func AdaptingOSStep_ToProto(mapCtx *direct.MapContext, in *krm.AdaptingOSStep) *pb.AdaptingOSStep {
	if in == nil {
		return nil
	}
	out := &pb.AdaptingOSStep{}
	return out
}
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
func CloneJob_FromProto(mapCtx *direct.MapContext, in *pb.CloneJob) *krm.CloneJob {
	if in == nil {
		return nil
	}
	out := &krm.CloneJob{}
	// MISSING: ComputeEngineTargetDetails
	// MISSING: CreateTime
	// MISSING: EndTime
	// MISSING: Name
	// MISSING: State
	// MISSING: StateTime
	// MISSING: Error
	// MISSING: Steps
	return out
}
func CloneJob_ToProto(mapCtx *direct.MapContext, in *krm.CloneJob) *pb.CloneJob {
	if in == nil {
		return nil
	}
	out := &pb.CloneJob{}
	// MISSING: ComputeEngineTargetDetails
	// MISSING: CreateTime
	// MISSING: EndTime
	// MISSING: Name
	// MISSING: State
	// MISSING: StateTime
	// MISSING: Error
	// MISSING: Steps
	return out
}
func CloneJobObservedState_FromProto(mapCtx *direct.MapContext, in *pb.CloneJob) *krm.CloneJobObservedState {
	if in == nil {
		return nil
	}
	out := &krm.CloneJobObservedState{}
	out.ComputeEngineTargetDetails = ComputeEngineTargetDetails_FromProto(mapCtx, in.GetComputeEngineTargetDetails())
	out.CreateTime = direct.StringTimestamp_FromProto(mapCtx, in.GetCreateTime())
	out.EndTime = direct.StringTimestamp_FromProto(mapCtx, in.GetEndTime())
	out.Name = direct.LazyPtr(in.GetName())
	out.State = direct.Enum_FromProto(mapCtx, in.GetState())
	out.StateTime = direct.StringTimestamp_FromProto(mapCtx, in.GetStateTime())
	out.Error = Status_FromProto(mapCtx, in.GetError())
	out.Steps = direct.Slice_FromProto(mapCtx, in.Steps, CloneStep_FromProto)
	return out
}
func CloneJobObservedState_ToProto(mapCtx *direct.MapContext, in *krm.CloneJobObservedState) *pb.CloneJob {
	if in == nil {
		return nil
	}
	out := &pb.CloneJob{}
	if oneof := ComputeEngineTargetDetails_ToProto(mapCtx, in.ComputeEngineTargetDetails); oneof != nil {
		out.TargetVmDetails = &pb.CloneJob_ComputeEngineTargetDetails{ComputeEngineTargetDetails: oneof}
	}
	out.CreateTime = direct.StringTimestamp_ToProto(mapCtx, in.CreateTime)
	out.EndTime = direct.StringTimestamp_ToProto(mapCtx, in.EndTime)
	out.Name = direct.ValueOf(in.Name)
	out.State = direct.Enum_ToProto[pb.CloneJob_State](mapCtx, in.State)
	out.StateTime = direct.StringTimestamp_ToProto(mapCtx, in.StateTime)
	out.Error = Status_ToProto(mapCtx, in.Error)
	out.Steps = direct.Slice_ToProto(mapCtx, in.Steps, CloneStep_ToProto)
	return out
}
func CloneStep_FromProto(mapCtx *direct.MapContext, in *pb.CloneStep) *krm.CloneStep {
	if in == nil {
		return nil
	}
	out := &krm.CloneStep{}
	out.AdaptingOs = AdaptingOSStep_FromProto(mapCtx, in.GetAdaptingOs())
	out.PreparingVmDisks = PreparingVMDisksStep_FromProto(mapCtx, in.GetPreparingVmDisks())
	out.InstantiatingMigratedVm = InstantiatingMigratedVMStep_FromProto(mapCtx, in.GetInstantiatingMigratedVm())
	out.StartTime = direct.StringTimestamp_FromProto(mapCtx, in.GetStartTime())
	out.EndTime = direct.StringTimestamp_FromProto(mapCtx, in.GetEndTime())
	return out
}
func CloneStep_ToProto(mapCtx *direct.MapContext, in *krm.CloneStep) *pb.CloneStep {
	if in == nil {
		return nil
	}
	out := &pb.CloneStep{}
	if oneof := AdaptingOSStep_ToProto(mapCtx, in.AdaptingOs); oneof != nil {
		out.Step = &pb.CloneStep_AdaptingOs{AdaptingOs: oneof}
	}
	if oneof := PreparingVMDisksStep_ToProto(mapCtx, in.PreparingVmDisks); oneof != nil {
		out.Step = &pb.CloneStep_PreparingVmDisks{PreparingVmDisks: oneof}
	}
	if oneof := InstantiatingMigratedVMStep_ToProto(mapCtx, in.InstantiatingMigratedVm); oneof != nil {
		out.Step = &pb.CloneStep_InstantiatingMigratedVm{InstantiatingMigratedVm: oneof}
	}
	out.StartTime = direct.StringTimestamp_ToProto(mapCtx, in.StartTime)
	out.EndTime = direct.StringTimestamp_ToProto(mapCtx, in.EndTime)
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
func VmmigrationCloneJobObservedState_FromProto(mapCtx *direct.MapContext, in *pb.CloneJob) *krm.VmmigrationCloneJobObservedState {
	if in == nil {
		return nil
	}
	out := &krm.VmmigrationCloneJobObservedState{}
	// MISSING: ComputeEngineTargetDetails
	// MISSING: CreateTime
	// MISSING: EndTime
	// MISSING: Name
	// MISSING: State
	// MISSING: StateTime
	// MISSING: Error
	// MISSING: Steps
	return out
}
func VmmigrationCloneJobObservedState_ToProto(mapCtx *direct.MapContext, in *krm.VmmigrationCloneJobObservedState) *pb.CloneJob {
	if in == nil {
		return nil
	}
	out := &pb.CloneJob{}
	// MISSING: ComputeEngineTargetDetails
	// MISSING: CreateTime
	// MISSING: EndTime
	// MISSING: Name
	// MISSING: State
	// MISSING: StateTime
	// MISSING: Error
	// MISSING: Steps
	return out
}
func VmmigrationCloneJobSpec_FromProto(mapCtx *direct.MapContext, in *pb.CloneJob) *krm.VmmigrationCloneJobSpec {
	if in == nil {
		return nil
	}
	out := &krm.VmmigrationCloneJobSpec{}
	// MISSING: ComputeEngineTargetDetails
	// MISSING: CreateTime
	// MISSING: EndTime
	// MISSING: Name
	// MISSING: State
	// MISSING: StateTime
	// MISSING: Error
	// MISSING: Steps
	return out
}
func VmmigrationCloneJobSpec_ToProto(mapCtx *direct.MapContext, in *krm.VmmigrationCloneJobSpec) *pb.CloneJob {
	if in == nil {
		return nil
	}
	out := &pb.CloneJob{}
	// MISSING: ComputeEngineTargetDetails
	// MISSING: CreateTime
	// MISSING: EndTime
	// MISSING: Name
	// MISSING: State
	// MISSING: StateTime
	// MISSING: Error
	// MISSING: Steps
	return out
}
