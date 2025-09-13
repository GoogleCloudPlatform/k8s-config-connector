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
// krm.group: compute.cnrm.cloud.google.com
// krm.version: v1beta1
// proto.service: google.cloud.compute.v1beta

package compute

import (
	pb "cloud.google.com/go/compute/apiv1beta/computepb"
	krmv1beta1 "github.com/GoogleCloudPlatform/k8s-config-connector/apis/compute/v1beta1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct"
)

func AcceleratorConfig_FromProto(mapCtx *direct.MapContext, in *pb.AcceleratorConfig) *krmv1beta1.AcceleratorConfig {
	if in == nil {
		return nil
	}
	out := &krmv1beta1.AcceleratorConfig{}
	out.AcceleratorCount = in.AcceleratorCount
	out.AcceleratorType = in.AcceleratorType
	return out
}
func AcceleratorConfig_ToProto(mapCtx *direct.MapContext, in *krmv1beta1.AcceleratorConfig) *pb.AcceleratorConfig {
	if in == nil {
		return nil
	}
	out := &pb.AcceleratorConfig{}
	out.AcceleratorCount = in.AcceleratorCount
	out.AcceleratorType = in.AcceleratorType
	return out
}
func AllocationAggregateReservation_FromProto(mapCtx *direct.MapContext, in *pb.AllocationAggregateReservation) *krmv1beta1.AllocationAggregateReservation {
	if in == nil {
		return nil
	}
	out := &krmv1beta1.AllocationAggregateReservation{}
	out.InUseResources = direct.Slice_FromProto(mapCtx, in.InUseResources, AllocationAggregateReservationReservedResourceInfo_FromProto)
	out.ReservedResources = direct.Slice_FromProto(mapCtx, in.ReservedResources, AllocationAggregateReservationReservedResourceInfo_FromProto)
	out.VMFamily = in.VmFamily
	out.WorkloadType = in.WorkloadType
	return out
}
func AllocationAggregateReservation_ToProto(mapCtx *direct.MapContext, in *krmv1beta1.AllocationAggregateReservation) *pb.AllocationAggregateReservation {
	if in == nil {
		return nil
	}
	out := &pb.AllocationAggregateReservation{}
	out.InUseResources = direct.Slice_ToProto(mapCtx, in.InUseResources, AllocationAggregateReservationReservedResourceInfo_ToProto)
	out.ReservedResources = direct.Slice_ToProto(mapCtx, in.ReservedResources, AllocationAggregateReservationReservedResourceInfo_ToProto)
	out.VmFamily = in.VMFamily
	out.WorkloadType = in.WorkloadType
	return out
}
func AllocationAggregateReservationReservedResourceInfo_FromProto(mapCtx *direct.MapContext, in *pb.AllocationAggregateReservationReservedResourceInfo) *krmv1beta1.AllocationAggregateReservationReservedResourceInfo {
	if in == nil {
		return nil
	}
	out := &krmv1beta1.AllocationAggregateReservationReservedResourceInfo{}
	out.Accelerator = AllocationAggregateReservationReservedResourceInfoAccelerator_FromProto(mapCtx, in.GetAccelerator())
	return out
}
func AllocationAggregateReservationReservedResourceInfo_ToProto(mapCtx *direct.MapContext, in *krmv1beta1.AllocationAggregateReservationReservedResourceInfo) *pb.AllocationAggregateReservationReservedResourceInfo {
	if in == nil {
		return nil
	}
	out := &pb.AllocationAggregateReservationReservedResourceInfo{}
	if oneof := AllocationAggregateReservationReservedResourceInfoAccelerator_ToProto(mapCtx, in.Accelerator); oneof != nil {
		out.Accelerator = &pb.AllocationAggregateReservationReservedResourceInfoAccelerator{
			AcceleratorCount: oneof.AcceleratorCount,
			AcceleratorType:  oneof.AcceleratorType,
		}
	}
	return out
}
func AllocationAggregateReservationReservedResourceInfoAccelerator_FromProto(mapCtx *direct.MapContext, in *pb.AllocationAggregateReservationReservedResourceInfoAccelerator) *krmv1beta1.AllocationAggregateReservationReservedResourceInfoAccelerator {
	if in == nil {
		return nil
	}
	out := &krmv1beta1.AllocationAggregateReservationReservedResourceInfoAccelerator{}
	out.AcceleratorCount = in.AcceleratorCount
	out.AcceleratorType = in.AcceleratorType
	return out
}
func AllocationAggregateReservationReservedResourceInfoAccelerator_ToProto(mapCtx *direct.MapContext, in *krmv1beta1.AllocationAggregateReservationReservedResourceInfoAccelerator) *pb.AllocationAggregateReservationReservedResourceInfoAccelerator {
	if in == nil {
		return nil
	}
	out := &pb.AllocationAggregateReservationReservedResourceInfoAccelerator{}
	out.AcceleratorCount = in.AcceleratorCount
	out.AcceleratorType = in.AcceleratorType
	return out
}
func AllocationSpecificSkuAllocationAllocatedInstancePropertiesReservedDisk_FromProto(mapCtx *direct.MapContext, in *pb.AllocationSpecificSKUAllocationAllocatedInstancePropertiesReservedDisk) *krmv1beta1.AllocationSpecificSkuAllocationAllocatedInstancePropertiesReservedDisk {
	if in == nil {
		return nil
	}
	out := &krmv1beta1.AllocationSpecificSkuAllocationAllocatedInstancePropertiesReservedDisk{}
	out.DiskSizeGB = in.DiskSizeGb
	out.Interface = in.Interface
	return out
}
func AllocationSpecificSkuAllocationAllocatedInstancePropertiesReservedDisk_ToProto(mapCtx *direct.MapContext, in *krmv1beta1.AllocationSpecificSkuAllocationAllocatedInstancePropertiesReservedDisk) *pb.AllocationSpecificSKUAllocationAllocatedInstancePropertiesReservedDisk {
	if in == nil {
		return nil
	}
	out := &pb.AllocationSpecificSKUAllocationAllocatedInstancePropertiesReservedDisk{}
	out.DiskSizeGb = in.DiskSizeGB
	out.Interface = in.Interface
	return out
}
func AllocationSpecificSkuAllocationReservedInstanceProperties_FromProto(mapCtx *direct.MapContext, in *pb.AllocationSpecificSKUAllocationReservedInstanceProperties) *krmv1beta1.AllocationSpecificSkuAllocationReservedInstanceProperties {
	if in == nil {
		return nil
	}
	out := &krmv1beta1.AllocationSpecificSkuAllocationReservedInstanceProperties{}
	out.GuestAccelerators = direct.Slice_FromProto(mapCtx, in.GuestAccelerators, AcceleratorConfig_FromProto)
	out.LocalSsds = direct.Slice_FromProto(mapCtx, in.LocalSsds, AllocationSpecificSkuAllocationAllocatedInstancePropertiesReservedDisk_FromProto)
	out.LocationHint = in.LocationHint
	out.MachineType = in.MachineType
	out.MaintenanceFreezeDurationHours = in.MaintenanceFreezeDurationHours
	out.MaintenanceInterval = in.MaintenanceInterval
	out.MinCPUPlatform = in.MinCpuPlatform
	return out
}
func AllocationSpecificSkuAllocationReservedInstanceProperties_ToProto(mapCtx *direct.MapContext, in *krmv1beta1.AllocationSpecificSkuAllocationReservedInstanceProperties) *pb.AllocationSpecificSKUAllocationReservedInstanceProperties {
	if in == nil {
		return nil
	}
	out := &pb.AllocationSpecificSKUAllocationReservedInstanceProperties{}
	out.GuestAccelerators = direct.Slice_ToProto(mapCtx, in.GuestAccelerators, AcceleratorConfig_ToProto)
	out.LocalSsds = direct.Slice_ToProto(mapCtx, in.LocalSsds, AllocationSpecificSkuAllocationAllocatedInstancePropertiesReservedDisk_ToProto)
	out.LocationHint = in.LocationHint
	out.MachineType = in.MachineType
	out.MaintenanceFreezeDurationHours = in.MaintenanceFreezeDurationHours
	out.MaintenanceInterval = in.MaintenanceInterval
	out.MinCpuPlatform = in.MinCPUPlatform
	return out
}
func Duration_FromProto(mapCtx *direct.MapContext, in *pb.Duration) *krmv1beta1.Duration {
	if in == nil {
		return nil
	}
	out := &krmv1beta1.Duration{}
	out.Nanos = in.Nanos
	out.Seconds = in.Seconds
	return out
}
func Duration_ToProto(mapCtx *direct.MapContext, in *krmv1beta1.Duration) *pb.Duration {
	if in == nil {
		return nil
	}
	out := &pb.Duration{}
	out.Nanos = in.Nanos
	out.Seconds = in.Seconds
	return out
}
func FutureReservationCommitmentInfo_FromProto(mapCtx *direct.MapContext, in *pb.FutureReservationCommitmentInfo) *krmv1beta1.FutureReservationCommitmentInfo {
	if in == nil {
		return nil
	}
	out := &krmv1beta1.FutureReservationCommitmentInfo{}
	out.CommitmentName = in.CommitmentName
	out.CommitmentPlan = in.CommitmentPlan
	out.PreviousCommitmentTerms = in.PreviousCommitmentTerms
	return out
}
func FutureReservationCommitmentInfo_ToProto(mapCtx *direct.MapContext, in *krmv1beta1.FutureReservationCommitmentInfo) *pb.FutureReservationCommitmentInfo {
	if in == nil {
		return nil
	}
	out := &pb.FutureReservationCommitmentInfo{}
	out.CommitmentName = in.CommitmentName
	out.CommitmentPlan = in.CommitmentPlan
	out.PreviousCommitmentTerms = in.PreviousCommitmentTerms
	return out
}
func ComputeFutureReservationObservedState_FromProto(mapCtx *direct.MapContext, in *pb.FutureReservation) *krmv1beta1.ComputeFutureReservationObservedState {
	if in == nil {
		return nil
	}
	out := &krmv1beta1.ComputeFutureReservationObservedState{}
	out.CreationTimestamp = in.CreationTimestamp
	out.ID = in.Id
	out.Kind = in.Kind
	out.SelfLink = in.SelfLink
	out.SelfLinkWithID = in.SelfLinkWithId
	out.SpecificSkuProperties = FutureReservationStatusSpecificSkuProperties_FromProto(mapCtx, in.GetStatus().GetSpecificSkuProperties())
	// MISSING: Status
	out.Zone = in.Zone
	return out
}
func ComputeFutureReservationObservedState_ToProto(mapCtx *direct.MapContext, in *krmv1beta1.ComputeFutureReservationObservedState) *pb.FutureReservation {
	if in == nil {
		return nil
	}
	out := &pb.FutureReservation{}
	out.CreationTimestamp = in.CreationTimestamp
	out.Id = in.ID
	out.Kind = in.Kind
	out.SelfLink = in.SelfLink
	out.SelfLinkWithId = in.SelfLinkWithID
	if oneof := FutureReservationStatusSpecificSkuProperties_ToProto(mapCtx, in.SpecificSkuProperties); oneof != nil {
		out.SpecificSkuProperties = &pb.FutureReservationSpecificSKUProperties{SourceInstanceTemplate: oneof.SourceInstanceTemplateId}
	}
	// MISSING: Status
	out.Zone = in.Zone
	return out
}
func ComputeFutureReservationSpec_FromProto(mapCtx *direct.MapContext, in *pb.FutureReservation) *krmv1beta1.ComputeFutureReservationSpec {
	if in == nil {
		return nil
	}
	out := &krmv1beta1.ComputeFutureReservationSpec{}
	out.AggregateReservation = AllocationAggregateReservation_FromProto(mapCtx, in.GetAggregateReservation())
	out.AutoCreatedReservationsDeleteTime = in.AutoCreatedReservationsDeleteTime
	out.AutoCreatedReservationsDuration = Duration_FromProto(mapCtx, in.GetAutoCreatedReservationsDuration())
	out.AutoDeleteAutoCreatedReservations = in.AutoDeleteAutoCreatedReservations
	out.CommitmentInfo = FutureReservationCommitmentInfo_FromProto(mapCtx, in.GetCommitmentInfo())
	out.DeploymentType = in.DeploymentType
	out.Description = in.Description
	out.EnableEmergentMaintenance = in.EnableEmergentMaintenance
	out.Name = in.Name
	out.NamePrefix = in.NamePrefix
	out.PlanningStatus = in.PlanningStatus
	out.ReservationMode = in.ReservationMode
	out.ReservationName = in.ReservationName
	out.SchedulingType = in.SchedulingType
	out.ShareSettings = ShareSettings_FromProto(mapCtx, in.GetShareSettings())
	out.SpecificReservationRequired = in.SpecificReservationRequired
	out.SpecificSkuProperties = FutureReservationSpecificSkuProperties_FromProto(mapCtx, in.GetSpecificSkuProperties())
	// MISSING: Status
	out.TimeWindow = FutureReservationTimeWindow_FromProto(mapCtx, in.GetTimeWindow())
	return out
}
func ComputeFutureReservationSpec_ToProto(mapCtx *direct.MapContext, in *krmv1beta1.ComputeFutureReservationSpec) *pb.FutureReservation {
	if in == nil {
		return nil
	}
	out := &pb.FutureReservation{}
	if oneof := AllocationAggregateReservation_ToProto(mapCtx, in.AggregateReservation); oneof != nil {
		out.AggregateReservation = &pb.AllocationAggregateReservation{
			InUseResources:    oneof.InUseResources,
			ReservedResources: oneof.ReservedResources,
			VmFamily:          oneof.VmFamily,
			WorkloadType:      oneof.WorkloadType,
		}
	}
	out.AutoCreatedReservationsDeleteTime = in.AutoCreatedReservationsDeleteTime
	if oneof := Duration_ToProto(mapCtx, in.AutoCreatedReservationsDuration); oneof != nil {
		out.AutoCreatedReservationsDuration = &pb.Duration{
			Nanos:   oneof.Nanos,
			Seconds: oneof.Seconds,
		}
	}
	out.AutoDeleteAutoCreatedReservations = in.AutoDeleteAutoCreatedReservations
	if oneof := FutureReservationCommitmentInfo_ToProto(mapCtx, in.CommitmentInfo); oneof != nil {
		out.CommitmentInfo = &pb.FutureReservationCommitmentInfo{
			CommitmentName:          oneof.CommitmentName,
			CommitmentPlan:          oneof.CommitmentPlan,
			PreviousCommitmentTerms: oneof.PreviousCommitmentTerms,
		}
	}
	out.DeploymentType = in.DeploymentType
	out.Description = in.Description
	out.EnableEmergentMaintenance = in.EnableEmergentMaintenance
	out.Name = in.Name
	out.NamePrefix = in.NamePrefix
	out.PlanningStatus = in.PlanningStatus
	out.ReservationMode = in.ReservationMode
	out.ReservationName = in.ReservationName
	out.SchedulingType = in.SchedulingType
	if oneof := ShareSettings_ToProto(mapCtx, in.ShareSettings); oneof != nil {
		out.ShareSettings = &pb.ShareSettings{
			ProjectMap: oneof.ProjectMap,
			Projects:   oneof.Projects,
			ShareType:  oneof.ShareType,
		}
	}
	out.SpecificReservationRequired = in.SpecificReservationRequired
	if oneof := FutureReservationSpecificSkuProperties_ToProto(mapCtx, in.SpecificSkuProperties); oneof != nil {
		out.SpecificSkuProperties = &pb.FutureReservationSpecificSKUProperties{
			InstanceProperties:     oneof.InstanceProperties,
			SourceInstanceTemplate: oneof.SourceInstanceTemplate,
			TotalCount:             oneof.TotalCount,
		}
	}
	// MISSING: Status
	if oneof := FutureReservationTimeWindow_ToProto(mapCtx, in.TimeWindow); oneof != nil {
		out.TimeWindow = &pb.FutureReservationTimeWindow{
			Duration:  oneof.Duration,
			EndTime:   oneof.EndTime,
			StartTime: oneof.StartTime,
		}
	}
	return out
}
func FutureReservationSpecificSkuProperties_FromProto(mapCtx *direct.MapContext, in *pb.FutureReservationSpecificSKUProperties) *krmv1beta1.FutureReservationSpecificSkuProperties {
	if in == nil {
		return nil
	}
	out := &krmv1beta1.FutureReservationSpecificSkuProperties{}
	out.InstanceProperties = AllocationSpecificSkuAllocationReservedInstanceProperties_FromProto(mapCtx, in.GetInstanceProperties())
	out.SourceInstanceTemplate = in.SourceInstanceTemplate
	out.TotalCount = in.TotalCount
	return out
}
func FutureReservationSpecificSkuProperties_ToProto(mapCtx *direct.MapContext, in *krmv1beta1.FutureReservationSpecificSkuProperties) *pb.FutureReservationSpecificSKUProperties {
	if in == nil {
		return nil
	}
	out := &pb.FutureReservationSpecificSKUProperties{}
	if oneof := AllocationSpecificSkuAllocationReservedInstanceProperties_ToProto(mapCtx, in.InstanceProperties); oneof != nil {
		out.InstanceProperties = &pb.AllocationSpecificSKUAllocationReservedInstanceProperties{
			GuestAccelerators:              oneof.GuestAccelerators,
			LocalSsds:                      oneof.LocalSsds,
			LocationHint:                   oneof.LocationHint,
			MachineType:                    oneof.MachineType,
			MaintenanceFreezeDurationHours: oneof.MaintenanceFreezeDurationHours,
			MaintenanceInterval:            oneof.MaintenanceInterval,
			MinCpuPlatform:                 oneof.MinCpuPlatform,
		}
	}
	out.SourceInstanceTemplate = in.SourceInstanceTemplate
	out.TotalCount = in.TotalCount
	return out
}
func FutureReservationStatusExistingMatchingUsageInfo_FromProto(mapCtx *direct.MapContext, in *pb.FutureReservationStatusExistingMatchingUsageInfo) *krmv1beta1.FutureReservationStatusExistingMatchingUsageInfo {
	if in == nil {
		return nil
	}
	out := &krmv1beta1.FutureReservationStatusExistingMatchingUsageInfo{}
	out.Count = in.Count
	out.Timestamp = in.Timestamp
	return out
}
func FutureReservationStatusExistingMatchingUsageInfo_ToProto(mapCtx *direct.MapContext, in *krmv1beta1.FutureReservationStatusExistingMatchingUsageInfo) *pb.FutureReservationStatusExistingMatchingUsageInfo {
	if in == nil {
		return nil
	}
	out := &pb.FutureReservationStatusExistingMatchingUsageInfo{}
	out.Count = in.Count
	out.Timestamp = in.Timestamp
	return out
}
func FutureReservationStatusLastKnownGoodState_FromProto(mapCtx *direct.MapContext, in *pb.FutureReservationStatusLastKnownGoodState) *krmv1beta1.FutureReservationStatusLastKnownGoodState {
	if in == nil {
		return nil
	}
	out := &krmv1beta1.FutureReservationStatusLastKnownGoodState{}
	out.Description = in.Description
	out.ExistingMatchingUsageInfo = FutureReservationStatusExistingMatchingUsageInfo_FromProto(mapCtx, in.GetExistingMatchingUsageInfo())
	out.FutureReservationSpecs = FutureReservationStatusLastKnownGoodStateFutureReservationSpecs_FromProto(mapCtx, in.GetFutureReservationSpecs())
	out.LockTime = in.LockTime
	out.NamePrefix = in.NamePrefix
	out.ProcurementStatus = in.ProcurementStatus
	return out
}
func FutureReservationStatusLastKnownGoodState_ToProto(mapCtx *direct.MapContext, in *krmv1beta1.FutureReservationStatusLastKnownGoodState) *pb.FutureReservationStatusLastKnownGoodState {
	if in == nil {
		return nil
	}
	out := &pb.FutureReservationStatusLastKnownGoodState{}
	out.Description = in.Description
	if oneof := FutureReservationStatusExistingMatchingUsageInfo_ToProto(mapCtx, in.ExistingMatchingUsageInfo); oneof != nil {
		out.ExistingMatchingUsageInfo = &pb.FutureReservationStatusExistingMatchingUsageInfo{
			Count:     oneof.Count,
			Timestamp: oneof.Timestamp,
		}
	}
	if oneof := FutureReservationStatusLastKnownGoodStateFutureReservationSpecs_ToProto(mapCtx, in.FutureReservationSpecs); oneof != nil {
		out.FutureReservationSpecs = &pb.FutureReservationStatusLastKnownGoodStateFutureReservationSpecs{
			ShareSettings:         oneof.ShareSettings,
			SpecificSkuProperties: oneof.SpecificSkuProperties,
			TimeWindow:            oneof.TimeWindow,
		}
	}
	out.LockTime = in.LockTime
	out.NamePrefix = in.NamePrefix
	out.ProcurementStatus = in.ProcurementStatus
	return out
}
func FutureReservationStatusLastKnownGoodStateFutureReservationSpecs_FromProto(mapCtx *direct.MapContext, in *pb.FutureReservationStatusLastKnownGoodStateFutureReservationSpecs) *krmv1beta1.FutureReservationStatusLastKnownGoodStateFutureReservationSpecs {
	if in == nil {
		return nil
	}
	out := &krmv1beta1.FutureReservationStatusLastKnownGoodStateFutureReservationSpecs{}
	out.ShareSettings = ShareSettings_FromProto(mapCtx, in.GetShareSettings())
	out.SpecificSkuProperties = FutureReservationSpecificSkuProperties_FromProto(mapCtx, in.GetSpecificSkuProperties())
	out.TimeWindow = FutureReservationTimeWindow_FromProto(mapCtx, in.GetTimeWindow())
	return out
}
func FutureReservationStatusLastKnownGoodStateFutureReservationSpecs_ToProto(mapCtx *direct.MapContext, in *krmv1beta1.FutureReservationStatusLastKnownGoodStateFutureReservationSpecs) *pb.FutureReservationStatusLastKnownGoodStateFutureReservationSpecs {
	if in == nil {
		return nil
	}
	out := &pb.FutureReservationStatusLastKnownGoodStateFutureReservationSpecs{}
	if oneof := ShareSettings_ToProto(mapCtx, in.ShareSettings); oneof != nil {
		out.ShareSettings = &pb.ShareSettings{
			ProjectMap: oneof.ProjectMap,
			Projects:   oneof.Projects,
			ShareType:  oneof.ShareType,
		}
	}
	if oneof := FutureReservationSpecificSkuProperties_ToProto(mapCtx, in.SpecificSkuProperties); oneof != nil {
		out.SpecificSkuProperties = &pb.FutureReservationSpecificSKUProperties{
			InstanceProperties:     oneof.InstanceProperties,
			SourceInstanceTemplate: oneof.SourceInstanceTemplate,
			TotalCount:             oneof.TotalCount,
		}
	}
	if oneof := FutureReservationTimeWindow_ToProto(mapCtx, in.TimeWindow); oneof != nil {
		out.TimeWindow = &pb.FutureReservationTimeWindow{
			Duration:  oneof.Duration,
			EndTime:   oneof.EndTime,
			StartTime: oneof.StartTime,
		}
	}
	return out
}
func FutureReservationStatusSpecificSkuProperties_FromProto(mapCtx *direct.MapContext, in *pb.FutureReservationStatusSpecificSKUProperties) *krmv1beta1.FutureReservationStatusSpecificSkuProperties {
	if in == nil {
		return nil
	}
	out := &krmv1beta1.FutureReservationStatusSpecificSkuProperties{}
	out.SourceInstanceTemplateID = in.SourceInstanceTemplateId
	return out
}
func FutureReservationStatusSpecificSkuProperties_ToProto(mapCtx *direct.MapContext, in *krmv1beta1.FutureReservationStatusSpecificSkuProperties) *pb.FutureReservationStatusSpecificSKUProperties {
	if in == nil {
		return nil
	}
	out := &pb.FutureReservationStatusSpecificSKUProperties{}
	out.SourceInstanceTemplateId = in.SourceInstanceTemplateID
	return out
}
func FutureReservationTimeWindow_FromProto(mapCtx *direct.MapContext, in *pb.FutureReservationTimeWindow) *krmv1beta1.FutureReservationTimeWindow {
	if in == nil {
		return nil
	}
	out := &krmv1beta1.FutureReservationTimeWindow{}
	out.Duration = Duration_FromProto(mapCtx, in.GetDuration())
	out.EndTime = in.EndTime
	out.StartTime = in.StartTime
	return out
}
func FutureReservationTimeWindow_ToProto(mapCtx *direct.MapContext, in *krmv1beta1.FutureReservationTimeWindow) *pb.FutureReservationTimeWindow {
	if in == nil {
		return nil
	}
	out := &pb.FutureReservationTimeWindow{}
	if oneof := Duration_ToProto(mapCtx, in.Duration); oneof != nil {
		out.Duration = &pb.Duration{
			Nanos:   oneof.Nanos,
			Seconds: oneof.Seconds,
		}
	}
	out.EndTime = in.EndTime
	out.StartTime = in.StartTime
	return out
}
func ShareSettings_FromProto(mapCtx *direct.MapContext, in *pb.ShareSettings) *krmv1beta1.ShareSettings {
	if in == nil {
		return nil
	}
	out := &krmv1beta1.ShareSettings{}
	// MISSING: ProjectMap
	out.Projects = in.Projects
	out.ShareType = in.ShareType
	return out
}
func ShareSettings_ToProto(mapCtx *direct.MapContext, in *krmv1beta1.ShareSettings) *pb.ShareSettings {
	if in == nil {
		return nil
	}
	out := &pb.ShareSettings{}
	// MISSING: ProjectMap
	out.Projects = in.Projects
	out.ShareType = in.ShareType
	return out
}
func ShareSettingsProjectConfig_FromProto(mapCtx *direct.MapContext, in *pb.ShareSettingsProjectConfig) *krmv1beta1.ShareSettingsProjectConfig {
	if in == nil {
		return nil
	}
	out := &krmv1beta1.ShareSettingsProjectConfig{}
	out.ProjectID = in.ProjectId
	return out
}
func ShareSettingsProjectConfig_ToProto(mapCtx *direct.MapContext, in *krmv1beta1.ShareSettingsProjectConfig) *pb.ShareSettingsProjectConfig {
	if in == nil {
		return nil
	}
	out := &pb.ShareSettingsProjectConfig{}
	out.ProjectId = in.ProjectID
	return out
}
