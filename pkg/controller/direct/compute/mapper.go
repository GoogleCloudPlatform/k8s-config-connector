// Copyright 2026 Google LLC
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

package compute

import (
	pb "cloud.google.com/go/compute/apiv1/computepb"
	krm "github.com/GoogleCloudPlatform/k8s-config-connector/apis/compute/v1beta1"
	refsv1beta1 "github.com/GoogleCloudPlatform/k8s-config-connector/apis/refs/v1beta1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct"
)

func ComputeResourcePolicySpec_v1beta1_FromProto(mapCtx *direct.MapContext, in *pb.ResourcePolicy) *krm.ComputeResourcePolicySpec {
	if in == nil {
		return nil
	}
	out := &krm.ComputeResourcePolicySpec{}
	out.Region = direct.LazyPtr(in.GetRegion())
	out.ResourceID = direct.LazyPtr(in.GetName())
	out.Description = direct.LazyPtr(in.GetDescription())
	out.DiskConsistencyGroupPolicy = ResourcePolicyDiskConsistencyGroupPolicy_v1beta1_FromProto(mapCtx, in.GetDiskConsistencyGroupPolicy())
	out.GroupPlacementPolicy = ResourcePolicyGroupPlacementPolicy_v1beta1_FromProto(mapCtx, in.GetGroupPlacementPolicy())
	out.InstanceSchedulePolicy = ResourcePolicyInstanceSchedulePolicy_v1beta1_FromProto(mapCtx, in.GetInstanceSchedulePolicy())
	out.SnapshotSchedulePolicy = ResourcePolicySnapshotSchedulePolicy_v1beta1_FromProto(mapCtx, in.GetSnapshotSchedulePolicy())
	return out
}

func ComputeResourcePolicySpec_v1beta1_ToProto(mapCtx *direct.MapContext, in *krm.ComputeResourcePolicySpec) *pb.ResourcePolicy {
	if in == nil {
		return nil
	}
	out := &pb.ResourcePolicy{}
	out.Region = in.Region
	out.Name = in.ResourceID
	out.Description = in.Description
	out.DiskConsistencyGroupPolicy = ResourcePolicyDiskConsistencyGroupPolicy_v1beta1_ToProto(mapCtx, in.DiskConsistencyGroupPolicy)
	out.GroupPlacementPolicy = ResourcePolicyGroupPlacementPolicy_v1beta1_ToProto(mapCtx, in.GroupPlacementPolicy)
	out.InstanceSchedulePolicy = ResourcePolicyInstanceSchedulePolicy_v1beta1_ToProto(mapCtx, in.InstanceSchedulePolicy)
	out.SnapshotSchedulePolicy = ResourcePolicySnapshotSchedulePolicy_v1beta1_ToProto(mapCtx, in.SnapshotSchedulePolicy)
	return out
}

func ResourcePolicyDiskConsistencyGroupPolicy_v1beta1_FromProto(mapCtx *direct.MapContext, in *pb.ResourcePolicyDiskConsistencyGroupPolicy) *krm.ResourcePolicyDiskConsistencyGroupPolicy {
	if in == nil {
		return nil
	}
	out := &krm.ResourcePolicyDiskConsistencyGroupPolicy{}
	out.Enabled = direct.LazyPtr(true)
	return out
}

func ResourcePolicyDiskConsistencyGroupPolicy_v1beta1_ToProto(mapCtx *direct.MapContext, in *krm.ResourcePolicyDiskConsistencyGroupPolicy) *pb.ResourcePolicyDiskConsistencyGroupPolicy {
	if in == nil {
		return nil
	}
	if in.Enabled != nil && !*in.Enabled {
		return nil
	}
	out := &pb.ResourcePolicyDiskConsistencyGroupPolicy{}
	return out
}

func ResourcePolicyGroupPlacementPolicy_v1beta1_FromProto(mapCtx *direct.MapContext, in *pb.ResourcePolicyGroupPlacementPolicy) *krm.ResourcePolicyGroupPlacementPolicy {
	if in == nil {
		return nil
	}
	out := &krm.ResourcePolicyGroupPlacementPolicy{}
	out.AvailabilityDomainCount = ConvertInt32ToInt(in.AvailabilityDomainCount)
	out.Collocation = direct.LazyPtr(in.GetCollocation())
	// maxDistance is missing from v1 proto but in CRD.
	out.VmCount = ConvertInt32ToInt(in.VmCount)
	return out
}

func ResourcePolicyGroupPlacementPolicy_v1beta1_ToProto(mapCtx *direct.MapContext, in *krm.ResourcePolicyGroupPlacementPolicy) *pb.ResourcePolicyGroupPlacementPolicy {
	if in == nil {
		return nil
	}
	out := &pb.ResourcePolicyGroupPlacementPolicy{}
	out.AvailabilityDomainCount = ConvertIntToInt32(in.AvailabilityDomainCount)
	out.Collocation = in.Collocation
	out.VmCount = ConvertIntToInt32(in.VmCount)
	return out
}

func ResourcePolicyInstanceSchedulePolicy_v1beta1_FromProto(mapCtx *direct.MapContext, in *pb.ResourcePolicyInstanceSchedulePolicy) *krm.ResourcePolicyInstanceSchedulePolicy {
	if in == nil {
		return nil
	}
	out := &krm.ResourcePolicyInstanceSchedulePolicy{}
	out.ExpirationTime = direct.LazyPtr(in.GetExpirationTime())
	out.StartTime = direct.LazyPtr(in.GetStartTime())
	out.TimeZone = direct.LazyPtr(in.GetTimeZone())
	out.VmStartSchedule = ResourcePolicyInstanceSchedulePolicySchedule_v1beta1_FromProto(mapCtx, in.GetVmStartSchedule())
	out.VmStopSchedule = ResourcePolicyInstanceSchedulePolicySchedule_v1beta1_FromProto(mapCtx, in.GetVmStopSchedule())
	return out
}

func ResourcePolicyInstanceSchedulePolicy_v1beta1_ToProto(mapCtx *direct.MapContext, in *krm.ResourcePolicyInstanceSchedulePolicy) *pb.ResourcePolicyInstanceSchedulePolicy {
	if in == nil {
		return nil
	}
	out := &pb.ResourcePolicyInstanceSchedulePolicy{}
	out.ExpirationTime = in.ExpirationTime
	out.StartTime = in.StartTime
	out.TimeZone = in.TimeZone
	out.VmStartSchedule = ResourcePolicyInstanceSchedulePolicySchedule_v1beta1_ToProto(mapCtx, in.VmStartSchedule)
	out.VmStopSchedule = ResourcePolicyInstanceSchedulePolicySchedule_v1beta1_ToProto(mapCtx, in.VmStopSchedule)
	return out
}

func ResourcePolicyInstanceSchedulePolicySchedule_v1beta1_FromProto(mapCtx *direct.MapContext, in *pb.ResourcePolicyInstanceSchedulePolicySchedule) *krm.ResourcePolicyInstanceSchedulePolicySchedule {
	if in == nil {
		return nil
	}
	out := &krm.ResourcePolicyInstanceSchedulePolicySchedule{}
	out.Schedule = direct.LazyPtr(in.GetSchedule())
	return out
}

func ResourcePolicyInstanceSchedulePolicySchedule_v1beta1_ToProto(mapCtx *direct.MapContext, in *krm.ResourcePolicyInstanceSchedulePolicySchedule) *pb.ResourcePolicyInstanceSchedulePolicySchedule {
	if in == nil {
		return nil
	}
	out := &pb.ResourcePolicyInstanceSchedulePolicySchedule{}
	out.Schedule = in.Schedule
	return out
}

func ResourcePolicySnapshotSchedulePolicy_v1beta1_FromProto(mapCtx *direct.MapContext, in *pb.ResourcePolicySnapshotSchedulePolicy) *krm.ResourcePolicySnapshotSchedulePolicy {
	if in == nil {
		return nil
	}
	out := &krm.ResourcePolicySnapshotSchedulePolicy{}
	out.RetentionPolicy = ResourcePolicySnapshotSchedulePolicyRetentionPolicy_v1beta1_FromProto(mapCtx, in.GetRetentionPolicy())
	out.Schedule = ResourcePolicySnapshotSchedulePolicySchedule_v1beta1_FromProto(mapCtx, in.GetSchedule())
	out.SnapshotProperties = ResourcePolicySnapshotSchedulePolicySnapshotProperties_v1beta1_FromProto(mapCtx, in.GetSnapshotProperties())
	return out
}

func ResourcePolicySnapshotSchedulePolicy_v1beta1_ToProto(mapCtx *direct.MapContext, in *krm.ResourcePolicySnapshotSchedulePolicy) *pb.ResourcePolicySnapshotSchedulePolicy {
	if in == nil {
		return nil
	}
	out := &pb.ResourcePolicySnapshotSchedulePolicy{}
	out.RetentionPolicy = ResourcePolicySnapshotSchedulePolicyRetentionPolicy_v1beta1_ToProto(mapCtx, in.RetentionPolicy)
	out.Schedule = ResourcePolicySnapshotSchedulePolicySchedule_v1beta1_ToProto(mapCtx, in.Schedule)
	out.SnapshotProperties = ResourcePolicySnapshotSchedulePolicySnapshotProperties_v1beta1_ToProto(mapCtx, in.SnapshotProperties)
	return out
}

func ResourcePolicySnapshotSchedulePolicyRetentionPolicy_v1beta1_FromProto(mapCtx *direct.MapContext, in *pb.ResourcePolicySnapshotSchedulePolicyRetentionPolicy) *krm.ResourcePolicySnapshotSchedulePolicyRetentionPolicy {
	if in == nil {
		return nil
	}
	out := &krm.ResourcePolicySnapshotSchedulePolicyRetentionPolicy{}
	out.MaxRetentionDays = ConvertInt32ToInt(in.MaxRetentionDays)
	out.OnSourceDiskDelete = direct.LazyPtr(in.GetOnSourceDiskDelete())
	return out
}

func ResourcePolicySnapshotSchedulePolicyRetentionPolicy_v1beta1_ToProto(mapCtx *direct.MapContext, in *krm.ResourcePolicySnapshotSchedulePolicyRetentionPolicy) *pb.ResourcePolicySnapshotSchedulePolicyRetentionPolicy {
	if in == nil {
		return nil
	}
	out := &pb.ResourcePolicySnapshotSchedulePolicyRetentionPolicy{}
	out.MaxRetentionDays = ConvertIntToInt32(in.MaxRetentionDays)
	out.OnSourceDiskDelete = in.OnSourceDiskDelete
	return out
}

func ResourcePolicySnapshotSchedulePolicySchedule_v1beta1_FromProto(mapCtx *direct.MapContext, in *pb.ResourcePolicySnapshotSchedulePolicySchedule) *krm.ResourcePolicySnapshotSchedulePolicySchedule {
	if in == nil {
		return nil
	}
	out := &krm.ResourcePolicySnapshotSchedulePolicySchedule{}
	out.DailySchedule = ResourcePolicySnapshotSchedulePolicyScheduleDailySchedule_v1beta1_FromProto(mapCtx, in.GetDailySchedule())
	out.HourlySchedule = ResourcePolicySnapshotSchedulePolicyScheduleHourlySchedule_v1beta1_FromProto(mapCtx, in.GetHourlySchedule())
	out.WeeklySchedule = ResourcePolicySnapshotSchedulePolicyScheduleWeeklySchedule_v1beta1_FromProto(mapCtx, in.GetWeeklySchedule())
	return out
}

func ResourcePolicySnapshotSchedulePolicySchedule_v1beta1_ToProto(mapCtx *direct.MapContext, in *krm.ResourcePolicySnapshotSchedulePolicySchedule) *pb.ResourcePolicySnapshotSchedulePolicySchedule {
	if in == nil {
		return nil
	}
	out := &pb.ResourcePolicySnapshotSchedulePolicySchedule{}
	out.DailySchedule = ResourcePolicySnapshotSchedulePolicyScheduleDailySchedule_v1beta1_ToProto(mapCtx, in.DailySchedule)
	out.HourlySchedule = ResourcePolicySnapshotSchedulePolicyScheduleHourlySchedule_v1beta1_ToProto(mapCtx, in.HourlySchedule)
	out.WeeklySchedule = ResourcePolicySnapshotSchedulePolicyScheduleWeeklySchedule_v1beta1_ToProto(mapCtx, in.WeeklySchedule)
	return out
}

func ResourcePolicySnapshotSchedulePolicyScheduleDailySchedule_v1beta1_FromProto(mapCtx *direct.MapContext, in *pb.ResourcePolicyDailyCycle) *krm.ResourcePolicySnapshotSchedulePolicyScheduleDailySchedule {
	if in == nil {
		return nil
	}
	out := &krm.ResourcePolicySnapshotSchedulePolicyScheduleDailySchedule{}
	out.DaysInCycle = ConvertInt32ToInt(in.DaysInCycle)
	out.StartTime = direct.LazyPtr(in.GetStartTime())
	return out
}

func ResourcePolicySnapshotSchedulePolicyScheduleDailySchedule_v1beta1_ToProto(mapCtx *direct.MapContext, in *krm.ResourcePolicySnapshotSchedulePolicyScheduleDailySchedule) *pb.ResourcePolicyDailyCycle {
	if in == nil {
		return nil
	}
	out := &pb.ResourcePolicyDailyCycle{}
	out.DaysInCycle = ConvertIntToInt32(in.DaysInCycle)
	out.StartTime = in.StartTime
	return out
}

func ResourcePolicySnapshotSchedulePolicyScheduleHourlySchedule_v1beta1_FromProto(mapCtx *direct.MapContext, in *pb.ResourcePolicyHourlyCycle) *krm.ResourcePolicySnapshotSchedulePolicyScheduleHourlySchedule {
	if in == nil {
		return nil
	}
	out := &krm.ResourcePolicySnapshotSchedulePolicyScheduleHourlySchedule{}
	out.HoursInCycle = ConvertInt32ToInt(in.HoursInCycle)
	out.StartTime = direct.LazyPtr(in.GetStartTime())
	return out
}

func ResourcePolicySnapshotSchedulePolicyScheduleHourlySchedule_v1beta1_ToProto(mapCtx *direct.MapContext, in *krm.ResourcePolicySnapshotSchedulePolicyScheduleHourlySchedule) *pb.ResourcePolicyHourlyCycle {
	if in == nil {
		return nil
	}
	out := &pb.ResourcePolicyHourlyCycle{}
	out.HoursInCycle = ConvertIntToInt32(in.HoursInCycle)
	out.StartTime = in.StartTime
	return out
}

func ResourcePolicySnapshotSchedulePolicyScheduleWeeklySchedule_v1beta1_FromProto(mapCtx *direct.MapContext, in *pb.ResourcePolicyWeeklyCycle) *krm.ResourcePolicySnapshotSchedulePolicyScheduleWeeklySchedule {
	if in == nil {
		return nil
	}
	out := &krm.ResourcePolicySnapshotSchedulePolicyScheduleWeeklySchedule{}
	out.DayOfWeeks = direct.Slice_FromProto(mapCtx, in.DayOfWeeks, ResourcePolicyWeeklyCycleDayOfWeek_v1beta1_FromProto)
	return out
}

func ResourcePolicySnapshotSchedulePolicyScheduleWeeklySchedule_v1beta1_ToProto(mapCtx *direct.MapContext, in *krm.ResourcePolicySnapshotSchedulePolicyScheduleWeeklySchedule) *pb.ResourcePolicyWeeklyCycle {
	if in == nil {
		return nil
	}
	out := &pb.ResourcePolicyWeeklyCycle{}
	out.DayOfWeeks = direct.Slice_ToProto(mapCtx, in.DayOfWeeks, ResourcePolicyWeeklyCycleDayOfWeek_v1beta1_ToProto)
	return out
}

func ResourcePolicyWeeklyCycleDayOfWeek_v1beta1_FromProto(mapCtx *direct.MapContext, in *pb.ResourcePolicyWeeklyCycleDayOfWeek) *krm.ResourcePolicySnapshotSchedulePolicyScheduleWeeklyScheduleDayOfWeek {
	if in == nil {
		return nil
	}
	out := &krm.ResourcePolicySnapshotSchedulePolicyScheduleWeeklyScheduleDayOfWeek{}
	out.Day = direct.LazyPtr(in.GetDay())
	out.StartTime = direct.LazyPtr(in.GetStartTime())
	return out
}

func ResourcePolicyWeeklyCycleDayOfWeek_v1beta1_ToProto(mapCtx *direct.MapContext, in *krm.ResourcePolicySnapshotSchedulePolicyScheduleWeeklyScheduleDayOfWeek) *pb.ResourcePolicyWeeklyCycleDayOfWeek {
	if in == nil {
		return nil
	}
	out := &pb.ResourcePolicyWeeklyCycleDayOfWeek{}
	out.Day = in.Day
	out.StartTime = in.StartTime
	return out
}

func ResourcePolicySnapshotSchedulePolicySnapshotProperties_v1beta1_FromProto(mapCtx *direct.MapContext, in *pb.ResourcePolicySnapshotSchedulePolicySnapshotProperties) *krm.ResourcePolicySnapshotSchedulePolicySnapshotProperties {
	if in == nil {
		return nil
	}
	out := &krm.ResourcePolicySnapshotSchedulePolicySnapshotProperties{}
	out.ChainName = direct.LazyPtr(in.GetChainName())
	out.GuestFlush = direct.LazyPtr(in.GetGuestFlush())
	out.Labels = in.Labels
	out.StorageLocations = in.StorageLocations
	return out
}

func ResourcePolicySnapshotSchedulePolicySnapshotProperties_v1beta1_ToProto(mapCtx *direct.MapContext, in *krm.ResourcePolicySnapshotSchedulePolicySnapshotProperties) *pb.ResourcePolicySnapshotSchedulePolicySnapshotProperties {
	if in == nil {
		return nil
	}
	out := &pb.ResourcePolicySnapshotSchedulePolicySnapshotProperties{}
	out.ChainName = in.ChainName
	out.GuestFlush = in.GuestFlush
	out.Labels = in.Labels
	out.StorageLocations = in.StorageLocations
	return out
}

func ComputeResourcePolicyObservedState_v1beta1_FromProto(mapCtx *direct.MapContext, in *pb.ResourcePolicy) *krm.ComputeResourcePolicyObservedState {
	if in == nil {
		return nil
	}
	out := &krm.ComputeResourcePolicyObservedState{}
	out.CreationTimestamp = direct.LazyPtr(in.GetCreationTimestamp())
	out.ID = in.Id
	out.Status = direct.LazyPtr(in.GetStatus())
	return out
}

func ComputeResourcePolicyObservedState_v1beta1_ToProto(mapCtx *direct.MapContext, in *krm.ComputeResourcePolicyObservedState) *pb.ResourcePolicy {
	if in == nil {
		return nil
	}
	out := &pb.ResourcePolicy{}
	out.CreationTimestamp = in.CreationTimestamp
	out.Id = in.ID
	out.Status = in.Status
	return out
}

func ConvertInt32ToInt(in *int32) *int {
	if in == nil {
		return nil
	}
	v := int(*in)
	return &v
}

func ConvertIntToInt32(in *int) *int32 {
	if in == nil {
		return nil
	}
	v := int32(*in)
	return &v
}

func ReservationSpecificReservation_v1beta1_FromProto(mapCtx *direct.MapContext, in *pb.AllocationSpecificSKUReservation) *krm.ReservationSpecificReservation {
	if in == nil {
		return nil
	}
	out := &krm.ReservationSpecificReservation{}
	out.Count = Int32_FromProto(in.Count)
	out.InUseCount = Int32_FromProto(in.InUseCount)
	out.InstanceProperties = ReservationInstanceProperties_v1beta1_FromProto(mapCtx, in.GetInstanceProperties())
	return out
}

func ReservationSpecificReservation_v1beta1_ToProto(mapCtx *direct.MapContext, in *krm.ReservationSpecificReservation) *pb.AllocationSpecificSKUReservation {
	if in == nil {
		return nil
	}
	out := &pb.AllocationSpecificSKUReservation{}
	out.Count = Int32_ToProto(in.Count)
	out.InUseCount = Int32_ToProto(in.InUseCount)
	out.InstanceProperties = ReservationInstanceProperties_v1beta1_ToProto(mapCtx, in.InstanceProperties)
	return out
}

func ReservationLocalSsds_v1beta1_FromProto(mapCtx *direct.MapContext, in *pb.AllocationSpecificSKUAllocationAllocatedInstancePropertiesReservedDisk) *krm.ReservationLocalSsds {
	if in == nil {
		return nil
	}
	out := &krm.ReservationLocalSsds{}
	out.DiskSizeGb = Int32_FromProto(in.DiskSizeGb)
	out.Interface = in.Interface
	return out
}

func ReservationLocalSsds_v1beta1_ToProto(mapCtx *direct.MapContext, in *krm.ReservationLocalSsds) *pb.AllocationSpecificSKUAllocationAllocatedInstancePropertiesReservedDisk {
	if in == nil {
		return nil
	}
	out := &pb.AllocationSpecificSKUAllocationAllocatedInstancePropertiesReservedDisk{}
	out.DiskSizeGb = Int32_ToProto(in.DiskSizeGb)
	out.Interface = in.Interface
	return out
}

/*
func ComputeReservationObservedState_v1beta1_FromProto(mapCtx *direct.MapContext, in *pb.Reservation) *krm.ComputeReservationObservedState {
	if in == nil {
		return nil
	}
	out := &krm.ComputeReservationObservedState{}
	out.ID = Uint64_FromProto(in.Id)
	out.ResourceStatus = AllocationResourceStatus_v1beta1_FromProto(mapCtx, in.GetResourceStatus())
	return out
}

func ComputeReservationObservedState_v1beta1_ToProto(mapCtx *direct.MapContext, in *krm.ComputeReservationObservedState) *pb.Reservation {
	if in == nil {
		return nil
	}
	out := &pb.Reservation{}
	out.Id = Uint64_ToProto(in.ID)
	out.ResourceStatus = AllocationResourceStatus_v1beta1_ToProto(mapCtx, in.ResourceStatus)
	return out
}
*/

func Int32_FromProto(in *int64) *int32 {
	if in == nil {
		return nil
	}
	out := int32(*in)
	return &out
}

func Int32_ToProto(in *int32) *int64 {
	if in == nil {
		return nil
	}
	out := int64(*in)
	return &out
}

func Uint64_FromProto(in *uint64) *int64 {
	if in == nil {
		return nil
	}
	out := int64(*in)
	return &out
}

func Uint64_ToProto(in *int64) *uint64 {
	if in == nil {
		return nil
	}
	out := uint64(*in)
	return &out
}

func ReservationShareSettings_v1beta1_FromProto(mapCtx *direct.MapContext, in *pb.ShareSettings) *krm.ReservationShareSettings {
	if in == nil {
		return nil
	}
	out := &krm.ReservationShareSettings{}
	out.ShareType = in.ShareType
	if in.ProjectMap != nil {
		for k, v := range in.ProjectMap {
			out.ProjectMap = append(out.ProjectMap, krm.ReservationProjectMap{
				KeyRef: &refsv1beta1.ProjectRef{
					External: k,
				},
				ProjectIDRef: &refsv1beta1.ProjectRef{
					External: v.GetProjectId(),
				},
			})
		}
	}
	return out
}

func ReservationShareSettings_v1beta1_ToProto(mapCtx *direct.MapContext, in *krm.ReservationShareSettings) *pb.ShareSettings {
	if in == nil {
		return nil
	}
	out := &pb.ShareSettings{}
	out.ShareType = in.ShareType
	if in.ProjectMap != nil {
		out.ProjectMap = make(map[string]*pb.ShareSettingsProjectConfig)
		for _, entry := range in.ProjectMap {
			var key string
			if entry.KeyRef != nil {
				key = entry.KeyRef.External
			}
			if key == "" {
				mapCtx.Errorf("project keyRef.external was not pre-resolved")
				continue
			}
			projectConfig := &pb.ShareSettingsProjectConfig{}
			if entry.ProjectIDRef != nil {
				projectConfig.ProjectId = &entry.ProjectIDRef.External
			}
			out.ProjectMap[key] = projectConfig
		}
	}
	return out
}
