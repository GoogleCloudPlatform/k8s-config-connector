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
func UtilizationReport_FromProto(mapCtx *direct.MapContext, in *pb.UtilizationReport) *krm.UtilizationReport {
	if in == nil {
		return nil
	}
	out := &krm.UtilizationReport{}
	// MISSING: Name
	out.DisplayName = direct.LazyPtr(in.GetDisplayName())
	// MISSING: State
	// MISSING: StateTime
	// MISSING: Error
	// MISSING: CreateTime
	out.TimeFrame = direct.Enum_FromProto(mapCtx, in.GetTimeFrame())
	// MISSING: FrameEndTime
	// MISSING: VmCount
	out.Vms = direct.Slice_FromProto(mapCtx, in.Vms, VmUtilizationInfo_FromProto)
	return out
}
func UtilizationReport_ToProto(mapCtx *direct.MapContext, in *krm.UtilizationReport) *pb.UtilizationReport {
	if in == nil {
		return nil
	}
	out := &pb.UtilizationReport{}
	// MISSING: Name
	out.DisplayName = direct.ValueOf(in.DisplayName)
	// MISSING: State
	// MISSING: StateTime
	// MISSING: Error
	// MISSING: CreateTime
	out.TimeFrame = direct.Enum_ToProto[pb.UtilizationReport_TimeFrame](mapCtx, in.TimeFrame)
	// MISSING: FrameEndTime
	// MISSING: VmCount
	out.Vms = direct.Slice_ToProto(mapCtx, in.Vms, VmUtilizationInfo_ToProto)
	return out
}
func UtilizationReportObservedState_FromProto(mapCtx *direct.MapContext, in *pb.UtilizationReport) *krm.UtilizationReportObservedState {
	if in == nil {
		return nil
	}
	out := &krm.UtilizationReportObservedState{}
	out.Name = direct.LazyPtr(in.GetName())
	// MISSING: DisplayName
	out.State = direct.Enum_FromProto(mapCtx, in.GetState())
	out.StateTime = direct.StringTimestamp_FromProto(mapCtx, in.GetStateTime())
	out.Error = Status_FromProto(mapCtx, in.GetError())
	out.CreateTime = direct.StringTimestamp_FromProto(mapCtx, in.GetCreateTime())
	// MISSING: TimeFrame
	out.FrameEndTime = direct.StringTimestamp_FromProto(mapCtx, in.GetFrameEndTime())
	out.VmCount = direct.LazyPtr(in.GetVmCount())
	out.Vms = direct.Slice_FromProto(mapCtx, in.Vms, VmUtilizationInfoObservedState_FromProto)
	return out
}
func UtilizationReportObservedState_ToProto(mapCtx *direct.MapContext, in *krm.UtilizationReportObservedState) *pb.UtilizationReport {
	if in == nil {
		return nil
	}
	out := &pb.UtilizationReport{}
	out.Name = direct.ValueOf(in.Name)
	// MISSING: DisplayName
	out.State = direct.Enum_ToProto[pb.UtilizationReport_State](mapCtx, in.State)
	out.StateTime = direct.StringTimestamp_ToProto(mapCtx, in.StateTime)
	out.Error = Status_ToProto(mapCtx, in.Error)
	out.CreateTime = direct.StringTimestamp_ToProto(mapCtx, in.CreateTime)
	// MISSING: TimeFrame
	out.FrameEndTime = direct.StringTimestamp_ToProto(mapCtx, in.FrameEndTime)
	out.VmCount = direct.ValueOf(in.VmCount)
	out.Vms = direct.Slice_ToProto(mapCtx, in.Vms, VmUtilizationInfoObservedState_ToProto)
	return out
}
func VmUtilizationInfo_FromProto(mapCtx *direct.MapContext, in *pb.VmUtilizationInfo) *krm.VmUtilizationInfo {
	if in == nil {
		return nil
	}
	out := &krm.VmUtilizationInfo{}
	out.VmwareVmDetails = VmwareVmDetails_FromProto(mapCtx, in.GetVmwareVmDetails())
	out.VmID = direct.LazyPtr(in.GetVmId())
	out.Utilization = VmUtilizationMetrics_FromProto(mapCtx, in.GetUtilization())
	return out
}
func VmUtilizationInfo_ToProto(mapCtx *direct.MapContext, in *krm.VmUtilizationInfo) *pb.VmUtilizationInfo {
	if in == nil {
		return nil
	}
	out := &pb.VmUtilizationInfo{}
	if oneof := VmwareVmDetails_ToProto(mapCtx, in.VmwareVmDetails); oneof != nil {
		out.VmDetails = &pb.VmUtilizationInfo_VmwareVmDetails{VmwareVmDetails: oneof}
	}
	out.VmId = direct.ValueOf(in.VmID)
	out.Utilization = VmUtilizationMetrics_ToProto(mapCtx, in.Utilization)
	return out
}
func VmUtilizationInfoObservedState_FromProto(mapCtx *direct.MapContext, in *pb.VmUtilizationInfo) *krm.VmUtilizationInfoObservedState {
	if in == nil {
		return nil
	}
	out := &krm.VmUtilizationInfoObservedState{}
	out.VmwareVmDetails = VmwareVmDetailsObservedState_FromProto(mapCtx, in.GetVmwareVmDetails())
	// MISSING: VmID
	// MISSING: Utilization
	return out
}
func VmUtilizationInfoObservedState_ToProto(mapCtx *direct.MapContext, in *krm.VmUtilizationInfoObservedState) *pb.VmUtilizationInfo {
	if in == nil {
		return nil
	}
	out := &pb.VmUtilizationInfo{}
	if oneof := VmwareVmDetailsObservedState_ToProto(mapCtx, in.VmwareVmDetails); oneof != nil {
		out.VmDetails = &pb.VmUtilizationInfo_VmwareVmDetails{VmwareVmDetails: oneof}
	}
	// MISSING: VmID
	// MISSING: Utilization
	return out
}
func VmUtilizationMetrics_FromProto(mapCtx *direct.MapContext, in *pb.VmUtilizationMetrics) *krm.VmUtilizationMetrics {
	if in == nil {
		return nil
	}
	out := &krm.VmUtilizationMetrics{}
	out.CpuMaxPercent = direct.LazyPtr(in.GetCpuMaxPercent())
	out.CpuAveragePercent = direct.LazyPtr(in.GetCpuAveragePercent())
	out.MemoryMaxPercent = direct.LazyPtr(in.GetMemoryMaxPercent())
	out.MemoryAveragePercent = direct.LazyPtr(in.GetMemoryAveragePercent())
	out.DiskIoRateMaxKbps = direct.LazyPtr(in.GetDiskIoRateMaxKbps())
	out.DiskIoRateAverageKbps = direct.LazyPtr(in.GetDiskIoRateAverageKbps())
	out.NetworkThroughputMaxKbps = direct.LazyPtr(in.GetNetworkThroughputMaxKbps())
	out.NetworkThroughputAverageKbps = direct.LazyPtr(in.GetNetworkThroughputAverageKbps())
	return out
}
func VmUtilizationMetrics_ToProto(mapCtx *direct.MapContext, in *krm.VmUtilizationMetrics) *pb.VmUtilizationMetrics {
	if in == nil {
		return nil
	}
	out := &pb.VmUtilizationMetrics{}
	out.CpuMaxPercent = direct.ValueOf(in.CpuMaxPercent)
	out.CpuAveragePercent = direct.ValueOf(in.CpuAveragePercent)
	out.MemoryMaxPercent = direct.ValueOf(in.MemoryMaxPercent)
	out.MemoryAveragePercent = direct.ValueOf(in.MemoryAveragePercent)
	out.DiskIoRateMaxKbps = direct.ValueOf(in.DiskIoRateMaxKbps)
	out.DiskIoRateAverageKbps = direct.ValueOf(in.DiskIoRateAverageKbps)
	out.NetworkThroughputMaxKbps = direct.ValueOf(in.NetworkThroughputMaxKbps)
	out.NetworkThroughputAverageKbps = direct.ValueOf(in.NetworkThroughputAverageKbps)
	return out
}
func VmmigrationUtilizationReportObservedState_FromProto(mapCtx *direct.MapContext, in *pb.UtilizationReport) *krm.VmmigrationUtilizationReportObservedState {
	if in == nil {
		return nil
	}
	out := &krm.VmmigrationUtilizationReportObservedState{}
	// MISSING: Name
	// MISSING: DisplayName
	// MISSING: State
	// MISSING: StateTime
	// MISSING: Error
	// MISSING: CreateTime
	// MISSING: TimeFrame
	// MISSING: FrameEndTime
	// MISSING: VmCount
	// MISSING: Vms
	return out
}
func VmmigrationUtilizationReportObservedState_ToProto(mapCtx *direct.MapContext, in *krm.VmmigrationUtilizationReportObservedState) *pb.UtilizationReport {
	if in == nil {
		return nil
	}
	out := &pb.UtilizationReport{}
	// MISSING: Name
	// MISSING: DisplayName
	// MISSING: State
	// MISSING: StateTime
	// MISSING: Error
	// MISSING: CreateTime
	// MISSING: TimeFrame
	// MISSING: FrameEndTime
	// MISSING: VmCount
	// MISSING: Vms
	return out
}
func VmmigrationUtilizationReportSpec_FromProto(mapCtx *direct.MapContext, in *pb.UtilizationReport) *krm.VmmigrationUtilizationReportSpec {
	if in == nil {
		return nil
	}
	out := &krm.VmmigrationUtilizationReportSpec{}
	// MISSING: Name
	// MISSING: DisplayName
	// MISSING: State
	// MISSING: StateTime
	// MISSING: Error
	// MISSING: CreateTime
	// MISSING: TimeFrame
	// MISSING: FrameEndTime
	// MISSING: VmCount
	// MISSING: Vms
	return out
}
func VmmigrationUtilizationReportSpec_ToProto(mapCtx *direct.MapContext, in *krm.VmmigrationUtilizationReportSpec) *pb.UtilizationReport {
	if in == nil {
		return nil
	}
	out := &pb.UtilizationReport{}
	// MISSING: Name
	// MISSING: DisplayName
	// MISSING: State
	// MISSING: StateTime
	// MISSING: Error
	// MISSING: CreateTime
	// MISSING: TimeFrame
	// MISSING: FrameEndTime
	// MISSING: VmCount
	// MISSING: Vms
	return out
}
func VmwareVmDetails_FromProto(mapCtx *direct.MapContext, in *pb.VmwareVmDetails) *krm.VmwareVmDetails {
	if in == nil {
		return nil
	}
	out := &krm.VmwareVmDetails{}
	out.VmID = direct.LazyPtr(in.GetVmId())
	out.DatacenterID = direct.LazyPtr(in.GetDatacenterId())
	out.DatacenterDescription = direct.LazyPtr(in.GetDatacenterDescription())
	out.Uuid = direct.LazyPtr(in.GetUuid())
	out.DisplayName = direct.LazyPtr(in.GetDisplayName())
	out.PowerState = direct.Enum_FromProto(mapCtx, in.GetPowerState())
	out.CpuCount = direct.LazyPtr(in.GetCpuCount())
	out.MemoryMb = direct.LazyPtr(in.GetMemoryMb())
	out.DiskCount = direct.LazyPtr(in.GetDiskCount())
	out.CommittedStorageMb = direct.LazyPtr(in.GetCommittedStorageMb())
	out.GuestDescription = direct.LazyPtr(in.GetGuestDescription())
	// MISSING: BootOption
	return out
}
func VmwareVmDetails_ToProto(mapCtx *direct.MapContext, in *krm.VmwareVmDetails) *pb.VmwareVmDetails {
	if in == nil {
		return nil
	}
	out := &pb.VmwareVmDetails{}
	out.VmId = direct.ValueOf(in.VmID)
	out.DatacenterId = direct.ValueOf(in.DatacenterID)
	out.DatacenterDescription = direct.ValueOf(in.DatacenterDescription)
	out.Uuid = direct.ValueOf(in.Uuid)
	out.DisplayName = direct.ValueOf(in.DisplayName)
	out.PowerState = direct.Enum_ToProto[pb.VmwareVmDetails_PowerState](mapCtx, in.PowerState)
	out.CpuCount = direct.ValueOf(in.CpuCount)
	out.MemoryMb = direct.ValueOf(in.MemoryMb)
	out.DiskCount = direct.ValueOf(in.DiskCount)
	out.CommittedStorageMb = direct.ValueOf(in.CommittedStorageMb)
	out.GuestDescription = direct.ValueOf(in.GuestDescription)
	// MISSING: BootOption
	return out
}
func VmwareVmDetailsObservedState_FromProto(mapCtx *direct.MapContext, in *pb.VmwareVmDetails) *krm.VmwareVmDetailsObservedState {
	if in == nil {
		return nil
	}
	out := &krm.VmwareVmDetailsObservedState{}
	// MISSING: VmID
	// MISSING: DatacenterID
	// MISSING: DatacenterDescription
	// MISSING: Uuid
	// MISSING: DisplayName
	// MISSING: PowerState
	// MISSING: CpuCount
	// MISSING: MemoryMb
	// MISSING: DiskCount
	// MISSING: CommittedStorageMb
	// MISSING: GuestDescription
	out.BootOption = direct.Enum_FromProto(mapCtx, in.GetBootOption())
	return out
}
func VmwareVmDetailsObservedState_ToProto(mapCtx *direct.MapContext, in *krm.VmwareVmDetailsObservedState) *pb.VmwareVmDetails {
	if in == nil {
		return nil
	}
	out := &pb.VmwareVmDetails{}
	// MISSING: VmID
	// MISSING: DatacenterID
	// MISSING: DatacenterDescription
	// MISSING: Uuid
	// MISSING: DisplayName
	// MISSING: PowerState
	// MISSING: CpuCount
	// MISSING: MemoryMb
	// MISSING: DiskCount
	// MISSING: CommittedStorageMb
	// MISSING: GuestDescription
	out.BootOption = direct.Enum_ToProto[pb.VmwareVmDetails_BootOption](mapCtx, in.BootOption)
	return out
}
