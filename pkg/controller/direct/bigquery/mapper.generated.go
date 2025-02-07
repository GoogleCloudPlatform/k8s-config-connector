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
func BiReservation_FromProto(mapCtx *direct.MapContext, in *pb.BiReservation) *krm.BiReservation {
	if in == nil {
		return nil
	}
	out := &krm.BiReservation{}
	out.Name = direct.LazyPtr(in.GetName())
	// MISSING: UpdateTime
	out.Size = direct.LazyPtr(in.GetSize())
	out.PreferredTables = direct.Slice_FromProto(mapCtx, in.PreferredTables, TableReference_FromProto)
	return out
}
func BiReservation_ToProto(mapCtx *direct.MapContext, in *krm.BiReservation) *pb.BiReservation {
	if in == nil {
		return nil
	}
	out := &pb.BiReservation{}
	out.Name = direct.ValueOf(in.Name)
	// MISSING: UpdateTime
	out.Size = direct.ValueOf(in.Size)
	out.PreferredTables = direct.Slice_ToProto(mapCtx, in.PreferredTables, TableReference_ToProto)
	return out
}
func BiReservationObservedState_FromProto(mapCtx *direct.MapContext, in *pb.BiReservation) *krm.BiReservationObservedState {
	if in == nil {
		return nil
	}
	out := &krm.BiReservationObservedState{}
	// MISSING: Name
	out.UpdateTime = direct.StringTimestamp_FromProto(mapCtx, in.GetUpdateTime())
	// MISSING: Size
	// MISSING: PreferredTables
	return out
}
func BiReservationObservedState_ToProto(mapCtx *direct.MapContext, in *krm.BiReservationObservedState) *pb.BiReservation {
	if in == nil {
		return nil
	}
	out := &pb.BiReservation{}
	// MISSING: Name
	out.UpdateTime = direct.StringTimestamp_ToProto(mapCtx, in.UpdateTime)
	// MISSING: Size
	// MISSING: PreferredTables
	return out
}
func BigqueryAssignmentObservedState_FromProto(mapCtx *direct.MapContext, in *pb.Assignment) *krm.BigqueryAssignmentObservedState {
	if in == nil {
		return nil
	}
	out := &krm.BigqueryAssignmentObservedState{}
	// MISSING: Name
	// MISSING: Assignee
	// MISSING: JobType
	// MISSING: State
	return out
}
func BigqueryAssignmentObservedState_ToProto(mapCtx *direct.MapContext, in *krm.BigqueryAssignmentObservedState) *pb.Assignment {
	if in == nil {
		return nil
	}
	out := &pb.Assignment{}
	// MISSING: Name
	// MISSING: Assignee
	// MISSING: JobType
	// MISSING: State
	return out
}
func BigqueryAssignmentSpec_FromProto(mapCtx *direct.MapContext, in *pb.Assignment) *krm.BigqueryAssignmentSpec {
	if in == nil {
		return nil
	}
	out := &krm.BigqueryAssignmentSpec{}
	// MISSING: Name
	// MISSING: Assignee
	// MISSING: JobType
	// MISSING: State
	return out
}
func BigqueryAssignmentSpec_ToProto(mapCtx *direct.MapContext, in *krm.BigqueryAssignmentSpec) *pb.Assignment {
	if in == nil {
		return nil
	}
	out := &pb.Assignment{}
	// MISSING: Name
	// MISSING: Assignee
	// MISSING: JobType
	// MISSING: State
	return out
}
func BigqueryBiReservationObservedState_FromProto(mapCtx *direct.MapContext, in *pb.BiReservation) *krm.BigqueryBiReservationObservedState {
	if in == nil {
		return nil
	}
	out := &krm.BigqueryBiReservationObservedState{}
	// MISSING: Name
	// MISSING: UpdateTime
	// MISSING: Size
	// MISSING: PreferredTables
	return out
}
func BigqueryBiReservationObservedState_ToProto(mapCtx *direct.MapContext, in *krm.BigqueryBiReservationObservedState) *pb.BiReservation {
	if in == nil {
		return nil
	}
	out := &pb.BiReservation{}
	// MISSING: Name
	// MISSING: UpdateTime
	// MISSING: Size
	// MISSING: PreferredTables
	return out
}
func BigqueryBiReservationSpec_FromProto(mapCtx *direct.MapContext, in *pb.BiReservation) *krm.BigqueryBiReservationSpec {
	if in == nil {
		return nil
	}
	out := &krm.BigqueryBiReservationSpec{}
	// MISSING: Name
	// MISSING: UpdateTime
	// MISSING: Size
	// MISSING: PreferredTables
	return out
}
func BigqueryBiReservationSpec_ToProto(mapCtx *direct.MapContext, in *krm.BigqueryBiReservationSpec) *pb.BiReservation {
	if in == nil {
		return nil
	}
	out := &pb.BiReservation{}
	// MISSING: Name
	// MISSING: UpdateTime
	// MISSING: Size
	// MISSING: PreferredTables
	return out
}
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
func TableReference_FromProto(mapCtx *direct.MapContext, in *pb.TableReference) *krm.TableReference {
	if in == nil {
		return nil
	}
	out := &krm.TableReference{}
	out.ProjectID = direct.LazyPtr(in.GetProjectId())
	out.DatasetID = direct.LazyPtr(in.GetDatasetId())
	out.TableID = direct.LazyPtr(in.GetTableId())
	return out
}
func TableReference_ToProto(mapCtx *direct.MapContext, in *krm.TableReference) *pb.TableReference {
	if in == nil {
		return nil
	}
	out := &pb.TableReference{}
	out.ProjectId = direct.ValueOf(in.ProjectID)
	out.DatasetId = direct.ValueOf(in.DatasetID)
	out.TableId = direct.ValueOf(in.TableID)
	return out
}
