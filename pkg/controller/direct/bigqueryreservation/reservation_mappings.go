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
	krm "github.com/GoogleCloudPlatform/k8s-config-connector/apis/bigqueryreservation/v1beta1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct"
	statuspb "google.golang.org/genproto/googleapis/rpc/status"
)

func BigQueryReservationReservationSpec_ToProto(mapCtx *direct.MapContext, in *krm.BigQueryReservationReservationSpec) *pb.Reservation {
	if in == nil {
		return nil
	}
	out := &pb.Reservation{}
	// MISSING: Name
	out.SlotCapacity = direct.ValueOf(in.SlotCapacity)
	out.IgnoreIdleSlots = direct.ValueOf(in.IgnoreIdleSlots)
	out.Autoscale = AutoscaleSpec_ToProto(mapCtx, in.Autoscale)
	out.Concurrency = direct.ValueOf(in.Concurrency)
	// MISSING: CreationTime
	// MISSING: UpdateTime
	// MISSING: MultiRegionAuxiliary
	out.Edition = direct.Enum_ToProto[pb.Edition](mapCtx, in.Edition)
	// MISSING: PrimaryLocation
	// MISSING: SecondaryLocation
	// MISSING: OriginalPrimaryLocation
	// MISSING: MaxSlots
	// MISSING: ScalingMode
	// MISSING: ReplicationStatus
	out.SecondaryLocation = FailoverSpec_ToProto(mapCtx, in.FailOver)
	return out
}

func BigQueryReservationReservationObservedState_ToProto(mapCtx *direct.MapContext, in *krm.BigQueryReservationReservationObservedState) *pb.Reservation {
	if in == nil {
		return nil
	}
	out := &pb.Reservation{}
	// MISSING: Name
	out.Autoscale = AutoscaleObservedState_ToProto(mapCtx, in.Autoscale)
	out.PrimaryLocation, out.SecondaryLocation, out.OriginalPrimaryLocation = FailoverObservedState_ToProto(mapCtx, in.FailOver)
	// MISSING: CreationTime
	// MISSING: UpdateTime
	// MISSING: MultiRegionAuxiliary
	// MISSING: PrimaryLocation
	// MISSING: SecondaryLocation
	// MISSING: OriginalPrimaryLocation
	// MISSING: MaxSlots
	// MISSING: ScalingMode
	// MISSING: ReplicationStatus
	return out
}

func BigQueryReservationReservationObservedState_FromProto(mapCtx *direct.MapContext, in *pb.Reservation) *krm.BigQueryReservationReservationObservedState {
	if in == nil {
		return nil
	}
	out := &krm.BigQueryReservationReservationObservedState{}
	// MISSING: Name
	out.Autoscale = AutoscaleObservedState_FromProto(mapCtx, in.GetAutoscale())
	// MISSING: CreationTime
	// MISSING: CreationTime
	// MISSING: UpdateTime
	out.FailOver = FailoverObservedState_FromProto(mapCtx, in)
	return out
}

func FailoverSpec_ToProto(mapCtx *direct.MapContext, in *krm.FailoverSpec) string {
	if in == nil {
		return ""
	}
	return direct.ValueOf(in.SecondaryLocation)
}

func Status_ToProto(mapCtx *direct.MapContext, in *krm.Status) *statuspb.Status {
	if in == nil {
		return nil
	}
	out := &statuspb.Status{}
	out.Code = direct.ValueOf(in.Code)
	out.Message = direct.ValueOf(in.Message)
	// NOTYET
	// out.Details
	return out
}
func FailoverObservedState_FromProto(mapCtx *direct.MapContext, in *pb.Reservation) *krm.FailoverObservedState {
	if in == nil {
		return nil
	}
	if in.PrimaryLocation == "" && in.SecondaryLocation == "" && in.OriginalPrimaryLocation == "" {
		return nil
	}
	out := &krm.FailoverObservedState{}
	out.PrimaryLocation = direct.LazyPtr(in.PrimaryLocation)
	out.SecondaryLocation = direct.LazyPtr(in.SecondaryLocation)
	out.OriginalPrimaryLocation = direct.LazyPtr(in.OriginalPrimaryLocation)
	return out
}

func FailoverObservedState_ToProto(mapCtx *direct.MapContext, in *krm.FailoverObservedState) (string, string, string) {
	if in == nil {
		return "", "", ""
	}
	return direct.ValueOf(in.PrimaryLocation), direct.ValueOf(in.SecondaryLocation), direct.ValueOf(in.OriginalPrimaryLocation)
}

func AutoscaleSpec_ToProto(mapCtx *direct.MapContext, in *krm.AutoscaleSpec) *pb.Reservation_Autoscale {
	if in == nil {
		return nil
	}
	out := &pb.Reservation_Autoscale{}
	out.MaxSlots = direct.ValueOf(in.MaxSlots)
	return out
}

func AutoscaleSpec_FromProto(mapctx *direct.MapContext, in *pb.Reservation_Autoscale) *krm.AutoscaleSpec {
	if in == nil {
		return nil
	}
	out := &krm.AutoscaleSpec{}
	out.MaxSlots = direct.LazyPtr(in.MaxSlots)
	return out
}

func AutoscaleObservedState_FromProto(mapCtx *direct.MapContext, in *pb.Reservation_Autoscale) *krm.AutoscaleObservedState {
	if in == nil {
		return nil
	}
	out := &krm.AutoscaleObservedState{}
	if out.CurrentSlots = direct.LazyPtr(in.GetCurrentSlots()); out.CurrentSlots == nil {
		return nil
	}
	return out
}

func Status_FromProto(mapCtx *direct.MapContext, in *statuspb.Status) *krm.Status {
	if in == nil {
		return nil
	}
	out := &krm.Status{}
	out.Code = direct.LazyPtr(in.GetCode())
	out.Message = direct.LazyPtr(in.GetMessage())
	// NOTYET
	// out.Details
	return out
}
