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
// krm.group: bigqueryreservation.cnrm.cloud.google.com
// krm.version: v1beta1
// proto.service: google.cloud.bigquery.reservation.v1

package bigqueryreservation

import (
	pb "cloud.google.com/go/bigquery/reservation/apiv1/reservationpb"
	krm "github.com/GoogleCloudPlatform/k8s-config-connector/apis/bigqueryreservation/v1beta1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct"
)

func AutoscaleObservedState_ToProto(mapCtx *direct.MapContext, in *krm.AutoscaleObservedState) *pb.Reservation_Autoscale {
	if in == nil {
		return nil
	}
	out := &pb.Reservation_Autoscale{}
	out.CurrentSlots = direct.ValueOf(in.CurrentSlots)
	// MISSING: MaxSlots
	return out
}
func BigQueryReservationAssignmentObservedState_FromProto(mapCtx *direct.MapContext, in *pb.Assignment) *krm.BigQueryReservationAssignmentObservedState {
	if in == nil {
		return nil
	}
	out := &krm.BigQueryReservationAssignmentObservedState{}
	// MISSING: Name
	out.State = direct.Enum_FromProto(mapCtx, in.GetState())
	// MISSING: EnableGeminiInBigquery
	return out
}
func BigQueryReservationAssignmentObservedState_ToProto(mapCtx *direct.MapContext, in *krm.BigQueryReservationAssignmentObservedState) *pb.Assignment {
	if in == nil {
		return nil
	}
	out := &pb.Assignment{}
	// MISSING: Name
	out.State = direct.Enum_ToProto[pb.Assignment_State](mapCtx, in.State)
	// MISSING: EnableGeminiInBigquery
	return out
}
func Reservation_ReplicationStatus_FromProto(mapCtx *direct.MapContext, in *pb.Reservation_ReplicationStatus) *krm.Reservation_ReplicationStatus {
	if in == nil {
		return nil
	}
	out := &krm.Reservation_ReplicationStatus{}
	// MISSING: Error
	// MISSING: LastErrorTime
	// MISSING: LastReplicationTime
	return out
}
func Reservation_ReplicationStatus_ToProto(mapCtx *direct.MapContext, in *krm.Reservation_ReplicationStatus) *pb.Reservation_ReplicationStatus {
	if in == nil {
		return nil
	}
	out := &pb.Reservation_ReplicationStatus{}
	// MISSING: Error
	// MISSING: LastErrorTime
	// MISSING: LastReplicationTime
	return out
}
func Reservation_ReplicationStatusObservedState_FromProto(mapCtx *direct.MapContext, in *pb.Reservation_ReplicationStatus) *krm.Reservation_ReplicationStatusObservedState {
	if in == nil {
		return nil
	}
	out := &krm.Reservation_ReplicationStatusObservedState{}
	out.Error = Status_FromProto(mapCtx, in.GetError())
	out.LastErrorTime = direct.StringTimestamp_FromProto(mapCtx, in.GetLastErrorTime())
	out.LastReplicationTime = direct.StringTimestamp_FromProto(mapCtx, in.GetLastReplicationTime())
	return out
}
func Reservation_ReplicationStatusObservedState_ToProto(mapCtx *direct.MapContext, in *krm.Reservation_ReplicationStatusObservedState) *pb.Reservation_ReplicationStatus {
	if in == nil {
		return nil
	}
	out := &pb.Reservation_ReplicationStatus{}
	out.Error = Status_ToProto(mapCtx, in.Error)
	out.LastErrorTime = direct.StringTimestamp_ToProto(mapCtx, in.LastErrorTime)
	out.LastReplicationTime = direct.StringTimestamp_ToProto(mapCtx, in.LastReplicationTime)
	return out
}
