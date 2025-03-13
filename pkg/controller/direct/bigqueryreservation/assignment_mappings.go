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

package bigqueryreservation

import (
	pb "cloud.google.com/go/bigquery/reservation/apiv1/reservationpb"
	krm "github.com/GoogleCloudPlatform/k8s-config-connector/apis/bigqueryreservation/v1alpha1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct"
)

func BigqueryReservationAssignmentSpec_FromProto(mapCtx *direct.MapContext, in *pb.Assignment) *krm.BigQueryReservationAssignmentSpec {
	if in == nil {
		return nil
	}
	out := &krm.BigQueryReservationAssignmentSpec{}
	out.Assignee = direct.LazyPtr(in.GetAssignee())
	out.JobType = direct.Enum_FromProto(mapCtx, in.GetJobType())
	return out
}

func BigqueryReservationAssignmentSpec_ToProto(mapCtx *direct.MapContext, in *krm.BigQueryReservationAssignmentSpec) *pb.Assignment {
	if in == nil {
		return nil
	}
	out := &pb.Assignment{}
	out.Assignee = direct.ValueOf(in.Assignee)
	out.JobType = direct.Enum_ToProto[pb.Assignment_JobType](mapCtx, in.JobType)
	return out
}

func BigqueryReservationAssignmentObservedState_FromProto(mapCtx *direct.MapContext, in *pb.Assignment) *krm.BigQueryReservationAssignmentObservedState {
	if in == nil {
		return nil
	}
	out := &krm.BigQueryReservationAssignmentObservedState{}
	out.Name = direct.LazyPtr(in.GetName())
	out.State = direct.Enum_FromProto(mapCtx, in.GetState())
	return out
}

func BigqueryReservationObservedState_ToProto(mapCtx *direct.MapContext, in *krm.BigQueryReservationAssignmentObservedState) *pb.Assignment {
	if in == nil {
		return nil
	}
	out := &pb.Assignment{}
	out.Name = direct.ValueOf(in.Name)
	out.State = direct.Enum_ToProto[pb.Assignment_State](mapCtx, in.State)
	return out
}
