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

package bigquery

import (
	refs "github.com/GoogleCloudPlatform/k8s-config-connector/apis/refs/v1beta1"
	pb "cloud.google.com/go/bigquery/reservation/apiv1/reservationpb"
	krm "github.com/GoogleCloudPlatform/k8s-config-connector/apis/bigquery/v1alpha1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct"
)
func BigqueryCapacityCommitmentObservedState_FromProto(mapCtx *direct.MapContext, in *pb.CapacityCommitment) *krm.BigqueryCapacityCommitmentObservedState {
	if in == nil {
		return nil
	}
	out := &krm.BigqueryCapacityCommitmentObservedState{}
	// MISSING: Name
	// MISSING: SlotCount
	// MISSING: Plan
	// MISSING: State
	// MISSING: CommitmentStartTime
	// MISSING: CommitmentEndTime
	// MISSING: FailureStatus
	// MISSING: RenewalPlan
	// MISSING: MultiRegionAuxiliary
	// MISSING: Edition
	// MISSING: IsFlatRate
	return out
}
func BigqueryCapacityCommitmentObservedState_ToProto(mapCtx *direct.MapContext, in *krm.BigqueryCapacityCommitmentObservedState) *pb.CapacityCommitment {
	if in == nil {
		return nil
	}
	out := &pb.CapacityCommitment{}
	// MISSING: Name
	// MISSING: SlotCount
	// MISSING: Plan
	// MISSING: State
	// MISSING: CommitmentStartTime
	// MISSING: CommitmentEndTime
	// MISSING: FailureStatus
	// MISSING: RenewalPlan
	// MISSING: MultiRegionAuxiliary
	// MISSING: Edition
	// MISSING: IsFlatRate
	return out
}
func BigqueryCapacityCommitmentSpec_FromProto(mapCtx *direct.MapContext, in *pb.CapacityCommitment) *krm.BigqueryCapacityCommitmentSpec {
	if in == nil {
		return nil
	}
	out := &krm.BigqueryCapacityCommitmentSpec{}
	// MISSING: Name
	// MISSING: SlotCount
	// MISSING: Plan
	// MISSING: State
	// MISSING: CommitmentStartTime
	// MISSING: CommitmentEndTime
	// MISSING: FailureStatus
	// MISSING: RenewalPlan
	// MISSING: MultiRegionAuxiliary
	// MISSING: Edition
	// MISSING: IsFlatRate
	return out
}
func BigqueryCapacityCommitmentSpec_ToProto(mapCtx *direct.MapContext, in *krm.BigqueryCapacityCommitmentSpec) *pb.CapacityCommitment {
	if in == nil {
		return nil
	}
	out := &pb.CapacityCommitment{}
	// MISSING: Name
	// MISSING: SlotCount
	// MISSING: Plan
	// MISSING: State
	// MISSING: CommitmentStartTime
	// MISSING: CommitmentEndTime
	// MISSING: FailureStatus
	// MISSING: RenewalPlan
	// MISSING: MultiRegionAuxiliary
	// MISSING: Edition
	// MISSING: IsFlatRate
	return out
}
func BigqueryReservationObservedState_FromProto(mapCtx *direct.MapContext, in *pb.Reservation) *krm.BigqueryReservationObservedState {
	if in == nil {
		return nil
	}
	out := &krm.BigqueryReservationObservedState{}
	// MISSING: Name
	// MISSING: SlotCapacity
	// MISSING: IgnoreIdleSlots
	// MISSING: Autoscale
	// MISSING: Concurrency
	// MISSING: CreationTime
	// MISSING: UpdateTime
	// MISSING: MultiRegionAuxiliary
	// MISSING: Edition
	// MISSING: PrimaryLocation
	// MISSING: SecondaryLocation
	// MISSING: OriginalPrimaryLocation
	return out
}
func BigqueryReservationObservedState_ToProto(mapCtx *direct.MapContext, in *krm.BigqueryReservationObservedState) *pb.Reservation {
	if in == nil {
		return nil
	}
	out := &pb.Reservation{}
	// MISSING: Name
	// MISSING: SlotCapacity
	// MISSING: IgnoreIdleSlots
	// MISSING: Autoscale
	// MISSING: Concurrency
	// MISSING: CreationTime
	// MISSING: UpdateTime
	// MISSING: MultiRegionAuxiliary
	// MISSING: Edition
	// MISSING: PrimaryLocation
	// MISSING: SecondaryLocation
	// MISSING: OriginalPrimaryLocation
	return out
}
func BigqueryReservationSpec_FromProto(mapCtx *direct.MapContext, in *pb.Reservation) *krm.BigqueryReservationSpec {
	if in == nil {
		return nil
	}
	out := &krm.BigqueryReservationSpec{}
	// MISSING: Name
	// MISSING: SlotCapacity
	// MISSING: IgnoreIdleSlots
	// MISSING: Autoscale
	// MISSING: Concurrency
	// MISSING: CreationTime
	// MISSING: UpdateTime
	// MISSING: MultiRegionAuxiliary
	// MISSING: Edition
	// MISSING: PrimaryLocation
	// MISSING: SecondaryLocation
	// MISSING: OriginalPrimaryLocation
	return out
}
func BigqueryReservationSpec_ToProto(mapCtx *direct.MapContext, in *krm.BigqueryReservationSpec) *pb.Reservation {
	if in == nil {
		return nil
	}
	out := &pb.Reservation{}
	// MISSING: Name
	// MISSING: SlotCapacity
	// MISSING: IgnoreIdleSlots
	// MISSING: Autoscale
	// MISSING: Concurrency
	// MISSING: CreationTime
	// MISSING: UpdateTime
	// MISSING: MultiRegionAuxiliary
	// MISSING: Edition
	// MISSING: PrimaryLocation
	// MISSING: SecondaryLocation
	// MISSING: OriginalPrimaryLocation
	return out
}
func CapacityCommitment_FromProto(mapCtx *direct.MapContext, in *pb.CapacityCommitment) *krm.CapacityCommitment {
	if in == nil {
		return nil
	}
	out := &krm.CapacityCommitment{}
	// MISSING: Name
	out.SlotCount = direct.LazyPtr(in.GetSlotCount())
	out.Plan = direct.Enum_FromProto(mapCtx, in.GetPlan())
	// MISSING: State
	// MISSING: CommitmentStartTime
	// MISSING: CommitmentEndTime
	// MISSING: FailureStatus
	out.RenewalPlan = direct.Enum_FromProto(mapCtx, in.GetRenewalPlan())
	out.MultiRegionAuxiliary = direct.LazyPtr(in.GetMultiRegionAuxiliary())
	out.Edition = direct.Enum_FromProto(mapCtx, in.GetEdition())
	// MISSING: IsFlatRate
	return out
}
func CapacityCommitment_ToProto(mapCtx *direct.MapContext, in *krm.CapacityCommitment) *pb.CapacityCommitment {
	if in == nil {
		return nil
	}
	out := &pb.CapacityCommitment{}
	// MISSING: Name
	out.SlotCount = direct.ValueOf(in.SlotCount)
	out.Plan = direct.Enum_ToProto[pb.CapacityCommitment_CommitmentPlan](mapCtx, in.Plan)
	// MISSING: State
	// MISSING: CommitmentStartTime
	// MISSING: CommitmentEndTime
	// MISSING: FailureStatus
	out.RenewalPlan = direct.Enum_ToProto[pb.CapacityCommitment_CommitmentPlan](mapCtx, in.RenewalPlan)
	out.MultiRegionAuxiliary = direct.ValueOf(in.MultiRegionAuxiliary)
	out.Edition = direct.Enum_ToProto[pb.Edition](mapCtx, in.Edition)
	// MISSING: IsFlatRate
	return out
}
func CapacityCommitmentObservedState_FromProto(mapCtx *direct.MapContext, in *pb.CapacityCommitment) *krm.CapacityCommitmentObservedState {
	if in == nil {
		return nil
	}
	out := &krm.CapacityCommitmentObservedState{}
	out.Name = direct.LazyPtr(in.GetName())
	// MISSING: SlotCount
	// MISSING: Plan
	out.State = direct.Enum_FromProto(mapCtx, in.GetState())
	out.CommitmentStartTime = direct.StringTimestamp_FromProto(mapCtx, in.GetCommitmentStartTime())
	out.CommitmentEndTime = direct.StringTimestamp_FromProto(mapCtx, in.GetCommitmentEndTime())
	out.FailureStatus = Status_FromProto(mapCtx, in.GetFailureStatus())
	// MISSING: RenewalPlan
	// MISSING: MultiRegionAuxiliary
	// MISSING: Edition
	out.IsFlatRate = direct.LazyPtr(in.GetIsFlatRate())
	return out
}
func CapacityCommitmentObservedState_ToProto(mapCtx *direct.MapContext, in *krm.CapacityCommitmentObservedState) *pb.CapacityCommitment {
	if in == nil {
		return nil
	}
	out := &pb.CapacityCommitment{}
	out.Name = direct.ValueOf(in.Name)
	// MISSING: SlotCount
	// MISSING: Plan
	out.State = direct.Enum_ToProto[pb.CapacityCommitment_State](mapCtx, in.State)
	out.CommitmentStartTime = direct.StringTimestamp_ToProto(mapCtx, in.CommitmentStartTime)
	out.CommitmentEndTime = direct.StringTimestamp_ToProto(mapCtx, in.CommitmentEndTime)
	out.FailureStatus = Status_ToProto(mapCtx, in.FailureStatus)
	// MISSING: RenewalPlan
	// MISSING: MultiRegionAuxiliary
	// MISSING: Edition
	out.IsFlatRate = direct.ValueOf(in.IsFlatRate)
	return out
}
