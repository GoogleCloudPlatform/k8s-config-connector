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

func ComputeHealthCheckSpec_v1beta1_FromProto(mapCtx *direct.MapContext, in *pb.HealthCheck) *krm.ComputeHealthCheckSpec {
	if in == nil {
		return nil
	}
	out := &krm.ComputeHealthCheckSpec{}
	out.CheckIntervalSec = Int64_FromProto(in.CheckIntervalSec)
	out.Description = in.Description
	out.GRPCHealthCheck = HealthCheckGRPCHealthCheck_v1beta1_FromProto(mapCtx, in.GrpcHealthCheck)
	out.HealthyThreshold = Int64_FromProto(in.HealthyThreshold)
	out.HTTP2HealthCheck = HealthCheckHTTP2HealthCheck_v1beta1_FromProto(mapCtx, in.Http2HealthCheck)
	out.HTTPHealthCheck = HealthCheckHTTPHealthCheck_v1beta1_FromProto(mapCtx, in.HttpHealthCheck)
	out.HTTPSHealthCheck = HealthCheckHTTPSHealthCheck_v1beta1_FromProto(mapCtx, in.HttpsHealthCheck)
	out.LogConfig = HealthCheckLogConfig_v1beta1_FromProto(mapCtx, in.LogConfig)
	out.SSLHealthCheck = HealthCheckSSLHealthCheck_v1beta1_FromProto(mapCtx, in.SslHealthCheck)
	out.TCPHealthCheck = HealthCheckTCPHealthCheck_v1beta1_FromProto(mapCtx, in.TcpHealthCheck)
	out.TimeoutSec = Int64_FromProto(in.TimeoutSec)
	out.UnhealthyThreshold = Int64_FromProto(in.UnhealthyThreshold)
	return out
}

func ComputeHealthCheckSpec_v1beta1_ToProto(mapCtx *direct.MapContext, in *krm.ComputeHealthCheckSpec) *pb.HealthCheck {
	if in == nil {
		return nil
	}
	out := &pb.HealthCheck{}
	out.CheckIntervalSec = Int64_ToProto(in.CheckIntervalSec)
	out.Description = in.Description
	out.GrpcHealthCheck = HealthCheckGRPCHealthCheck_v1beta1_ToProto(mapCtx, in.GRPCHealthCheck)
	out.HealthyThreshold = Int64_ToProto(in.HealthyThreshold)
	out.Http2HealthCheck = HealthCheckHTTP2HealthCheck_v1beta1_ToProto(mapCtx, in.HTTP2HealthCheck)
	out.HttpHealthCheck = HealthCheckHTTPHealthCheck_v1beta1_ToProto(mapCtx, in.HTTPHealthCheck)
	out.HttpsHealthCheck = HealthCheckHTTPSHealthCheck_v1beta1_ToProto(mapCtx, in.HTTPSHealthCheck)
	out.LogConfig = HealthCheckLogConfig_v1beta1_ToProto(mapCtx, in.LogConfig)
	out.SslHealthCheck = HealthCheckSSLHealthCheck_v1beta1_ToProto(mapCtx, in.SSLHealthCheck)
	out.TcpHealthCheck = HealthCheckTCPHealthCheck_v1beta1_ToProto(mapCtx, in.TCPHealthCheck)
	out.TimeoutSec = Int64_ToProto(in.TimeoutSec)
	out.UnhealthyThreshold = Int64_ToProto(in.UnhealthyThreshold)
	return out
}

func HealthCheckGRPCHealthCheck_v1beta1_FromProto(mapCtx *direct.MapContext, in *pb.GRPCHealthCheck) *krm.HealthCheckGRPCHealthCheck {
	if in == nil {
		return nil
	}
	out := &krm.HealthCheckGRPCHealthCheck{}
	out.GRPCServiceName = in.GrpcServiceName
	out.Port = Int64_FromProto(in.Port)
	out.PortName = in.PortName
	out.PortSpecification = in.PortSpecification
	return out
}

func HealthCheckGRPCHealthCheck_v1beta1_ToProto(mapCtx *direct.MapContext, in *krm.HealthCheckGRPCHealthCheck) *pb.GRPCHealthCheck {
	if in == nil {
		return nil
	}
	out := &pb.GRPCHealthCheck{}
	out.GrpcServiceName = in.GRPCServiceName
	out.Port = Int64_ToProto(in.Port)
	out.PortName = in.PortName
	out.PortSpecification = in.PortSpecification
	return out
}

func HealthCheckHTTP2HealthCheck_v1beta1_FromProto(mapCtx *direct.MapContext, in *pb.HTTP2HealthCheck) *krm.HealthCheckHTTP2HealthCheck {
	if in == nil {
		return nil
	}
	out := &krm.HealthCheckHTTP2HealthCheck{}
	out.Host = in.Host
	out.Port = Int64_FromProto(in.Port)
	out.PortName = in.PortName
	out.PortSpecification = in.PortSpecification
	out.ProxyHeader = in.ProxyHeader
	out.RequestPath = in.RequestPath
	out.Response = in.Response
	return out
}

func HealthCheckHTTP2HealthCheck_v1beta1_ToProto(mapCtx *direct.MapContext, in *krm.HealthCheckHTTP2HealthCheck) *pb.HTTP2HealthCheck {
	if in == nil {
		return nil
	}
	out := &pb.HTTP2HealthCheck{}
	out.Host = in.Host
	out.Port = Int64_ToProto(in.Port)
	out.PortName = in.PortName
	out.PortSpecification = in.PortSpecification
	out.ProxyHeader = in.ProxyHeader
	out.RequestPath = in.RequestPath
	out.Response = in.Response
	return out
}

func HealthCheckHTTPHealthCheck_v1beta1_FromProto(mapCtx *direct.MapContext, in *pb.HTTPHealthCheck) *krm.HealthCheckHTTPHealthCheck {
	if in == nil {
		return nil
	}
	out := &krm.HealthCheckHTTPHealthCheck{}
	out.Host = in.Host
	out.Port = Int64_FromProto(in.Port)
	out.PortName = in.PortName
	out.PortSpecification = in.PortSpecification
	out.ProxyHeader = in.ProxyHeader
	out.RequestPath = in.RequestPath
	out.Response = in.Response
	return out
}

func HealthCheckHTTPHealthCheck_v1beta1_ToProto(mapCtx *direct.MapContext, in *krm.HealthCheckHTTPHealthCheck) *pb.HTTPHealthCheck {
	if in == nil {
		return nil
	}
	out := &pb.HTTPHealthCheck{}
	out.Host = in.Host
	out.Port = Int64_ToProto(in.Port)
	out.PortName = in.PortName
	out.PortSpecification = in.PortSpecification
	out.ProxyHeader = in.ProxyHeader
	out.RequestPath = in.RequestPath
	out.Response = in.Response
	return out
}

func HealthCheckHTTPSHealthCheck_v1beta1_FromProto(mapCtx *direct.MapContext, in *pb.HTTPSHealthCheck) *krm.HealthCheckHTTPSHealthCheck {
	if in == nil {
		return nil
	}
	out := &krm.HealthCheckHTTPSHealthCheck{}
	out.Host = in.Host
	out.Port = Int64_FromProto(in.Port)
	out.PortName = in.PortName
	out.PortSpecification = in.PortSpecification
	out.ProxyHeader = in.ProxyHeader
	out.RequestPath = in.RequestPath
	out.Response = in.Response
	return out
}

func HealthCheckHTTPSHealthCheck_v1beta1_ToProto(mapCtx *direct.MapContext, in *krm.HealthCheckHTTPSHealthCheck) *pb.HTTPSHealthCheck {
	if in == nil {
		return nil
	}
	out := &pb.HTTPSHealthCheck{}
	out.Host = in.Host
	out.Port = Int64_ToProto(in.Port)
	out.PortName = in.PortName
	out.PortSpecification = in.PortSpecification
	out.ProxyHeader = in.ProxyHeader
	out.RequestPath = in.RequestPath
	out.Response = in.Response
	return out
}

func HealthCheckSSLHealthCheck_v1beta1_FromProto(mapCtx *direct.MapContext, in *pb.SSLHealthCheck) *krm.HealthCheckSSLHealthCheck {
	if in == nil {
		return nil
	}
	out := &krm.HealthCheckSSLHealthCheck{}
	out.Port = Int64_FromProto(in.Port)
	out.PortName = in.PortName
	out.PortSpecification = in.PortSpecification
	out.ProxyHeader = in.ProxyHeader
	out.Request = in.Request
	out.Response = in.Response
	return out
}

func HealthCheckSSLHealthCheck_v1beta1_ToProto(mapCtx *direct.MapContext, in *krm.HealthCheckSSLHealthCheck) *pb.SSLHealthCheck {
	if in == nil {
		return nil
	}
	out := &pb.SSLHealthCheck{}
	out.Port = Int64_ToProto(in.Port)
	out.PortName = in.PortName
	out.PortSpecification = in.PortSpecification
	out.ProxyHeader = in.ProxyHeader
	out.Request = in.Request
	out.Response = in.Response
	return out
}

func HealthCheckTCPHealthCheck_v1beta1_FromProto(mapCtx *direct.MapContext, in *pb.TCPHealthCheck) *krm.HealthCheckTCPHealthCheck {
	if in == nil {
		return nil
	}
	out := &krm.HealthCheckTCPHealthCheck{}
	out.Port = Int64_FromProto(in.Port)
	out.PortName = in.PortName
	out.PortSpecification = in.PortSpecification
	out.ProxyHeader = in.ProxyHeader
	out.Request = in.Request
	out.Response = in.Response
	return out
}

func HealthCheckTCPHealthCheck_v1beta1_ToProto(mapCtx *direct.MapContext, in *krm.HealthCheckTCPHealthCheck) *pb.TCPHealthCheck {
	if in == nil {
		return nil
	}
	out := &pb.TCPHealthCheck{}
	out.Port = Int64_ToProto(in.Port)
	out.PortName = in.PortName
	out.PortSpecification = in.PortSpecification
	out.ProxyHeader = in.ProxyHeader
	out.Request = in.Request
	out.Response = in.Response
	return out
}

func ComputeHealthCheckStatus_v1beta1_FromProto(mapCtx *direct.MapContext, in *pb.HealthCheck) *krm.ComputeHealthCheckStatus {
	if in == nil {
		return nil
	}
	out := &krm.ComputeHealthCheckStatus{}
	out.CreationTimestamp = in.CreationTimestamp
	out.SelfLink = in.SelfLink
	out.Type = in.Type
	return out
}

func ComputeHealthCheckStatus_v1beta1_ToProto(mapCtx *direct.MapContext, in *krm.ComputeHealthCheckStatus) *pb.HealthCheck {
	if in == nil {
		return nil
	}
	out := &pb.HealthCheck{}
	out.CreationTimestamp = in.CreationTimestamp
	out.SelfLink = in.SelfLink
	out.Type = in.Type
	return out
}

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

func Int64_FromProto(in *int32) *int64 {
	if in == nil {
		return nil
	}
	out := int64(*in)
	return &out
}

func Int64_ToProto(in *int64) *int32 {
	if in == nil {
		return nil
	}
	out := int32(*in)
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
