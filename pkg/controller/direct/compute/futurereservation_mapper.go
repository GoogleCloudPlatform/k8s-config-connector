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
	krmcomputev1alpha1 "github.com/GoogleCloudPlatform/k8s-config-connector/apis/compute/v1alpha1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct"
)

func ComputeFutureReservationObservedState_v1alpha1_FromProto(mapCtx *direct.MapContext, in *pb.FutureReservation) *krmcomputev1alpha1.ComputeFutureReservationObservedState {
	if in == nil {
		return nil
	}
	out := &krmcomputev1alpha1.ComputeFutureReservationObservedState{}
	out.AggregateReservation = AllocationAggregateReservationObservedState_v1alpha1_FromProto(mapCtx, in.GetAggregateReservation())
	out.CreationTimestamp = in.CreationTimestamp
	out.ID = in.Id
	out.Kind = in.Kind
	// MISSING: Name
	out.SelfLink = in.SelfLink
	out.SelfLinkWithID = in.SelfLinkWithId
	out.Status = FutureReservationStatus_v1alpha1_FromProto(mapCtx, in.GetStatus())
	out.Zone = in.Zone
	return out
}
func ComputeFutureReservationObservedState_v1alpha1_ToProto(mapCtx *direct.MapContext, in *krmcomputev1alpha1.ComputeFutureReservationObservedState) *pb.FutureReservation {
	if in == nil {
		return nil
	}
	out := &pb.FutureReservation{}
	out.AggregateReservation = AllocationAggregateReservationObservedState_v1alpha1_ToProto(mapCtx, in.AggregateReservation)
	out.CreationTimestamp = in.CreationTimestamp
	out.Id = in.ID
	out.Kind = in.Kind
	// MISSING: Name
	out.SelfLink = in.SelfLink
	out.SelfLinkWithId = in.SelfLinkWithID
	out.Status = FutureReservationStatus_v1alpha1_ToProto(mapCtx, in.Status)
	out.Zone = in.Zone
	return out
}
func ComputeFutureReservationSpec_v1alpha1_FromProto(mapCtx *direct.MapContext, in *pb.FutureReservation) *krmcomputev1alpha1.ComputeFutureReservationSpec {
	if in == nil {
		return nil
	}
	out := &krmcomputev1alpha1.ComputeFutureReservationSpec{}
	out.AggregateReservation = AllocationAggregateReservation_v1alpha1_FromProto(mapCtx, in.GetAggregateReservation())
	out.AutoCreatedReservationsDeleteTime = in.AutoCreatedReservationsDeleteTime
	out.AutoCreatedReservationsDuration = Duration_v1alpha1_FromProto(mapCtx, in.GetAutoCreatedReservationsDuration())
	out.AutoDeleteAutoCreatedReservations = in.AutoDeleteAutoCreatedReservations
	out.CommitmentInfo = FutureReservationCommitmentInfo_v1alpha1_FromProto(mapCtx, in.GetCommitmentInfo())
	out.DeploymentType = in.DeploymentType
	out.Description = in.Description
	out.EnableEmergentMaintenance = in.EnableEmergentMaintenance
	// MISSING: Name
	out.NamePrefix = in.NamePrefix
	out.PlanningStatus = in.PlanningStatus
	out.ReservationMode = in.ReservationMode
	out.ReservationName = in.ReservationName
	out.SchedulingType = in.SchedulingType
	out.ShareSettings = ShareSettings_v1alpha1_FromProto(mapCtx, in.GetShareSettings())
	out.SpecificReservationRequired = in.SpecificReservationRequired
	out.SpecificSkuProperties = FutureReservationSpecificSkuProperties_v1alpha1_FromProto(mapCtx, in.GetSpecificSkuProperties())
	out.TimeWindow = FutureReservationTimeWindow_v1alpha1_FromProto(mapCtx, in.GetTimeWindow())
	return out
}
func ComputeFutureReservationSpec_v1alpha1_ToProto(mapCtx *direct.MapContext, in *krmcomputev1alpha1.ComputeFutureReservationSpec) *pb.FutureReservation {
	if in == nil {
		return nil
	}
	out := &pb.FutureReservation{}
	out.AggregateReservation = AllocationAggregateReservation_v1alpha1_ToProto(mapCtx, in.AggregateReservation)
	out.AutoCreatedReservationsDeleteTime = in.AutoCreatedReservationsDeleteTime
	out.AutoCreatedReservationsDuration = Duration_v1alpha1_ToProto(mapCtx, in.AutoCreatedReservationsDuration)
	out.AutoDeleteAutoCreatedReservations = in.AutoDeleteAutoCreatedReservations
	out.CommitmentInfo = FutureReservationCommitmentInfo_v1alpha1_ToProto(mapCtx, in.CommitmentInfo)
	out.DeploymentType = in.DeploymentType
	out.Description = in.Description
	out.EnableEmergentMaintenance = in.EnableEmergentMaintenance
	// MISSING: Name
	out.NamePrefix = in.NamePrefix
	out.PlanningStatus = in.PlanningStatus
	out.ReservationMode = in.ReservationMode
	out.ReservationName = in.ReservationName
	out.SchedulingType = in.SchedulingType
	out.ShareSettings = ShareSettings_v1alpha1_ToProto(mapCtx, in.ShareSettings)
	out.SpecificReservationRequired = in.SpecificReservationRequired
	out.SpecificSkuProperties = FutureReservationSpecificSkuProperties_v1alpha1_ToProto(mapCtx, in.SpecificSkuProperties)
	out.TimeWindow = FutureReservationTimeWindow_v1alpha1_ToProto(mapCtx, in.TimeWindow)
	return out
}
func FutureReservation_v1alpha1_FromProto(mapCtx *direct.MapContext, in *pb.FutureReservation) *krmcomputev1alpha1.FutureReservation {
	if in == nil {
		return nil
	}
	out := &krmcomputev1alpha1.FutureReservation{}
	out.AggregateReservation = AllocationAggregateReservation_v1alpha1_FromProto(mapCtx, in.GetAggregateReservation())
	out.AutoCreatedReservationsDeleteTime = in.AutoCreatedReservationsDeleteTime
	out.AutoCreatedReservationsDuration = Duration_v1alpha1_FromProto(mapCtx, in.GetAutoCreatedReservationsDuration())
	out.AutoDeleteAutoCreatedReservations = in.AutoDeleteAutoCreatedReservations
	out.CommitmentInfo = FutureReservationCommitmentInfo_v1alpha1_FromProto(mapCtx, in.GetCommitmentInfo())
	out.CreationTimestamp = in.CreationTimestamp
	out.DeploymentType = in.DeploymentType
	out.Description = in.Description
	out.EnableEmergentMaintenance = in.EnableEmergentMaintenance
	out.ID = in.Id
	out.Kind = in.Kind
	out.Name = in.Name
	out.NamePrefix = in.NamePrefix
	out.PlanningStatus = in.PlanningStatus
	out.ReservationMode = in.ReservationMode
	out.ReservationName = in.ReservationName
	out.SchedulingType = in.SchedulingType
	out.SelfLink = in.SelfLink
	out.SelfLinkWithID = in.SelfLinkWithId
	out.ShareSettings = ShareSettings_v1alpha1_FromProto(mapCtx, in.GetShareSettings())
	out.SpecificReservationRequired = in.SpecificReservationRequired
	out.SpecificSkuProperties = FutureReservationSpecificSkuProperties_v1alpha1_FromProto(mapCtx, in.GetSpecificSkuProperties())
	out.Status = FutureReservationStatus_v1alpha1_FromProto(mapCtx, in.GetStatus())
	out.TimeWindow = FutureReservationTimeWindow_v1alpha1_FromProto(mapCtx, in.GetTimeWindow())
	out.Zone = in.Zone
	return out
}
func FutureReservation_v1alpha1_ToProto(mapCtx *direct.MapContext, in *krmcomputev1alpha1.FutureReservation) *pb.FutureReservation {
	if in == nil {
		return nil
	}
	out := &pb.FutureReservation{}
	out.AggregateReservation = AllocationAggregateReservation_v1alpha1_ToProto(mapCtx, in.AggregateReservation)
	out.AutoCreatedReservationsDeleteTime = in.AutoCreatedReservationsDeleteTime
	out.AutoCreatedReservationsDuration = Duration_v1alpha1_ToProto(mapCtx, in.AutoCreatedReservationsDuration)
	out.AutoDeleteAutoCreatedReservations = in.AutoDeleteAutoCreatedReservations
	out.CommitmentInfo = FutureReservationCommitmentInfo_v1alpha1_ToProto(mapCtx, in.CommitmentInfo)
	out.CreationTimestamp = in.CreationTimestamp
	out.DeploymentType = in.DeploymentType
	out.Description = in.Description
	out.EnableEmergentMaintenance = in.EnableEmergentMaintenance
	out.Id = in.ID
	out.Kind = in.Kind
	out.Name = in.Name
	out.NamePrefix = in.NamePrefix
	out.PlanningStatus = in.PlanningStatus
	out.ReservationMode = in.ReservationMode
	out.ReservationName = in.ReservationName
	out.SchedulingType = in.SchedulingType
	out.SelfLink = in.SelfLink
	out.SelfLinkWithId = in.SelfLinkWithID
	out.ShareSettings = ShareSettings_v1alpha1_ToProto(mapCtx, in.ShareSettings)
	out.SpecificReservationRequired = in.SpecificReservationRequired
	out.SpecificSkuProperties = FutureReservationSpecificSkuProperties_v1alpha1_ToProto(mapCtx, in.SpecificSkuProperties)
	out.Status = FutureReservationStatus_v1alpha1_ToProto(mapCtx, in.Status)
	out.TimeWindow = FutureReservationTimeWindow_v1alpha1_ToProto(mapCtx, in.TimeWindow)
	out.Zone = in.Zone
	return out
}
func FutureReservationCommitmentInfo_v1alpha1_FromProto(mapCtx *direct.MapContext, in *pb.FutureReservationCommitmentInfo) *krmcomputev1alpha1.FutureReservationCommitmentInfo {
	if in == nil {
		return nil
	}
	out := &krmcomputev1alpha1.FutureReservationCommitmentInfo{}
	out.CommitmentName = in.CommitmentName
	out.CommitmentPlan = in.CommitmentPlan
	out.PreviousCommitmentTerms = in.PreviousCommitmentTerms
	return out
}
func FutureReservationCommitmentInfo_v1alpha1_ToProto(mapCtx *direct.MapContext, in *krmcomputev1alpha1.FutureReservationCommitmentInfo) *pb.FutureReservationCommitmentInfo {
	if in == nil {
		return nil
	}
	out := &pb.FutureReservationCommitmentInfo{}
	out.CommitmentName = in.CommitmentName
	out.CommitmentPlan = in.CommitmentPlan
	out.PreviousCommitmentTerms = in.PreviousCommitmentTerms
	return out
}
func FutureReservationSpecificSkuProperties_v1alpha1_FromProto(mapCtx *direct.MapContext, in *pb.FutureReservationSpecificSKUProperties) *krmcomputev1alpha1.FutureReservationSpecificSkuProperties {
	if in == nil {
		return nil
	}
	out := &krmcomputev1alpha1.FutureReservationSpecificSkuProperties{}
	out.InstanceProperties = AllocationSpecificSkuAllocationReservedInstanceProperties_v1alpha1_FromProto(mapCtx, in.GetInstanceProperties())
	out.SourceInstanceTemplate = in.SourceInstanceTemplate
	out.TotalCount = in.TotalCount
	return out
}
func FutureReservationSpecificSkuProperties_v1alpha1_ToProto(mapCtx *direct.MapContext, in *krmcomputev1alpha1.FutureReservationSpecificSkuProperties) *pb.FutureReservationSpecificSKUProperties {
	if in == nil {
		return nil
	}
	out := &pb.FutureReservationSpecificSKUProperties{}
	out.InstanceProperties = AllocationSpecificSkuAllocationReservedInstanceProperties_v1alpha1_ToProto(mapCtx, in.InstanceProperties)
	out.SourceInstanceTemplate = in.SourceInstanceTemplate
	out.TotalCount = in.TotalCount
	return out
}
func FutureReservationStatus_v1alpha1_FromProto(mapCtx *direct.MapContext, in *pb.FutureReservationStatus) *krmcomputev1alpha1.FutureReservationStatus {
	if in == nil {
		return nil
	}
	out := &krmcomputev1alpha1.FutureReservationStatus{}
	out.AmendmentStatus = in.AmendmentStatus
	out.AutoCreatedReservations = in.AutoCreatedReservations
	out.ExistingMatchingUsageInfo = FutureReservationStatusExistingMatchingUsageInfo_v1alpha1_FromProto(mapCtx, in.GetExistingMatchingUsageInfo())
	out.FulfilledCount = in.FulfilledCount
	out.LastKnownGoodState = FutureReservationStatusLastKnownGoodState_v1alpha1_FromProto(mapCtx, in.GetLastKnownGoodState())
	out.LockTime = in.LockTime
	out.ProcurementStatus = in.ProcurementStatus
	out.SpecificSkuProperties = FutureReservationStatusSpecificSkuProperties_v1alpha1_FromProto(mapCtx, in.GetSpecificSkuProperties())
	return out
}
func FutureReservationStatus_v1alpha1_ToProto(mapCtx *direct.MapContext, in *krmcomputev1alpha1.FutureReservationStatus) *pb.FutureReservationStatus {
	if in == nil {
		return nil
	}
	out := &pb.FutureReservationStatus{}
	out.AmendmentStatus = in.AmendmentStatus
	out.AutoCreatedReservations = in.AutoCreatedReservations
	out.ExistingMatchingUsageInfo = FutureReservationStatusExistingMatchingUsageInfo_v1alpha1_ToProto(mapCtx, in.ExistingMatchingUsageInfo)
	out.FulfilledCount = in.FulfilledCount
	out.LastKnownGoodState = FutureReservationStatusLastKnownGoodState_v1alpha1_ToProto(mapCtx, in.LastKnownGoodState)
	out.LockTime = in.LockTime
	out.ProcurementStatus = in.ProcurementStatus
	out.SpecificSkuProperties = FutureReservationStatusSpecificSkuProperties_v1alpha1_ToProto(mapCtx, in.SpecificSkuProperties)
	return out
}
func FutureReservationStatusExistingMatchingUsageInfo_v1alpha1_FromProto(mapCtx *direct.MapContext, in *pb.FutureReservationStatusExistingMatchingUsageInfo) *krmcomputev1alpha1.FutureReservationStatusExistingMatchingUsageInfo {
	if in == nil {
		return nil
	}
	out := &krmcomputev1alpha1.FutureReservationStatusExistingMatchingUsageInfo{}
	out.Count = in.Count
	out.Timestamp = in.Timestamp
	return out
}
func FutureReservationStatusExistingMatchingUsageInfo_v1alpha1_ToProto(mapCtx *direct.MapContext, in *krmcomputev1alpha1.FutureReservationStatusExistingMatchingUsageInfo) *pb.FutureReservationStatusExistingMatchingUsageInfo {
	if in == nil {
		return nil
	}
	out := &pb.FutureReservationStatusExistingMatchingUsageInfo{}
	out.Count = in.Count
	out.Timestamp = in.Timestamp
	return out
}
func FutureReservationStatusLastKnownGoodState_v1alpha1_FromProto(mapCtx *direct.MapContext, in *pb.FutureReservationStatusLastKnownGoodState) *krmcomputev1alpha1.FutureReservationStatusLastKnownGoodState {
	if in == nil {
		return nil
	}
	out := &krmcomputev1alpha1.FutureReservationStatusLastKnownGoodState{}
	out.Description = in.Description
	out.ExistingMatchingUsageInfo = FutureReservationStatusExistingMatchingUsageInfo_v1alpha1_FromProto(mapCtx, in.GetExistingMatchingUsageInfo())
	out.FutureReservationSpecs = FutureReservationStatusLastKnownGoodStateFutureReservationSpecs_v1alpha1_FromProto(mapCtx, in.GetFutureReservationSpecs())
	out.LockTime = in.LockTime
	out.NamePrefix = in.NamePrefix
	out.ProcurementStatus = in.ProcurementStatus
	return out
}
func FutureReservationStatusLastKnownGoodState_v1alpha1_ToProto(mapCtx *direct.MapContext, in *krmcomputev1alpha1.FutureReservationStatusLastKnownGoodState) *pb.FutureReservationStatusLastKnownGoodState {
	if in == nil {
		return nil
	}
	out := &pb.FutureReservationStatusLastKnownGoodState{}
	out.Description = in.Description
	out.ExistingMatchingUsageInfo = FutureReservationStatusExistingMatchingUsageInfo_v1alpha1_ToProto(mapCtx, in.ExistingMatchingUsageInfo)
	out.FutureReservationSpecs = FutureReservationStatusLastKnownGoodStateFutureReservationSpecs_v1alpha1_ToProto(mapCtx, in.FutureReservationSpecs)
	out.LockTime = in.LockTime
	out.NamePrefix = in.NamePrefix
	out.ProcurementStatus = in.ProcurementStatus
	return out
}
func FutureReservationStatusLastKnownGoodStateFutureReservationSpecs_v1alpha1_FromProto(mapCtx *direct.MapContext, in *pb.FutureReservationStatusLastKnownGoodStateFutureReservationSpecs) *krmcomputev1alpha1.FutureReservationStatusLastKnownGoodStateFutureReservationSpecs {
	if in == nil {
		return nil
	}
	out := &krmcomputev1alpha1.FutureReservationStatusLastKnownGoodStateFutureReservationSpecs{}
	out.ShareSettings = ShareSettings_v1alpha1_FromProto(mapCtx, in.GetShareSettings())
	out.SpecificSkuProperties = FutureReservationSpecificSkuProperties_v1alpha1_FromProto(mapCtx, in.GetSpecificSkuProperties())
	out.TimeWindow = FutureReservationTimeWindow_v1alpha1_FromProto(mapCtx, in.GetTimeWindow())
	return out
}
func FutureReservationStatusLastKnownGoodStateFutureReservationSpecs_v1alpha1_ToProto(mapCtx *direct.MapContext, in *krmcomputev1alpha1.FutureReservationStatusLastKnownGoodStateFutureReservationSpecs) *pb.FutureReservationStatusLastKnownGoodStateFutureReservationSpecs {
	if in == nil {
		return nil
	}
	out := &pb.FutureReservationStatusLastKnownGoodStateFutureReservationSpecs{}
	out.ShareSettings = ShareSettings_v1alpha1_ToProto(mapCtx, in.ShareSettings)
	out.SpecificSkuProperties = FutureReservationSpecificSkuProperties_v1alpha1_ToProto(mapCtx, in.SpecificSkuProperties)
	out.TimeWindow = FutureReservationTimeWindow_v1alpha1_ToProto(mapCtx, in.TimeWindow)
	return out
}
func FutureReservationStatusSpecificSkuProperties_v1alpha1_FromProto(mapCtx *direct.MapContext, in *pb.FutureReservationStatusSpecificSKUProperties) *krmcomputev1alpha1.FutureReservationStatusSpecificSkuProperties {
	if in == nil {
		return nil
	}
	out := &krmcomputev1alpha1.FutureReservationStatusSpecificSkuProperties{}
	out.SourceInstanceTemplateID = in.SourceInstanceTemplateId
	return out
}
func FutureReservationStatusSpecificSkuProperties_v1alpha1_ToProto(mapCtx *direct.MapContext, in *krmcomputev1alpha1.FutureReservationStatusSpecificSkuProperties) *pb.FutureReservationStatusSpecificSKUProperties {
	if in == nil {
		return nil
	}
	out := &pb.FutureReservationStatusSpecificSKUProperties{}
	out.SourceInstanceTemplateId = in.SourceInstanceTemplateID
	return out
}
func FutureReservationTimeWindow_v1alpha1_FromProto(mapCtx *direct.MapContext, in *pb.FutureReservationTimeWindow) *krmcomputev1alpha1.FutureReservationTimeWindow {
	if in == nil {
		return nil
	}
	out := &krmcomputev1alpha1.FutureReservationTimeWindow{}
	out.Duration = Duration_v1alpha1_FromProto(mapCtx, in.GetDuration())
	out.EndTime = in.EndTime
	out.StartTime = in.StartTime
	return out
}
func FutureReservationTimeWindow_v1alpha1_ToProto(mapCtx *direct.MapContext, in *krmcomputev1alpha1.FutureReservationTimeWindow) *pb.FutureReservationTimeWindow {
	if in == nil {
		return nil
	}
	out := &pb.FutureReservationTimeWindow{}
	out.Duration = Duration_v1alpha1_ToProto(mapCtx, in.Duration)
	out.EndTime = in.EndTime
	out.StartTime = in.StartTime
	return out
}
