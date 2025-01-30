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

func BigqueryReservationReservationSpec_ToProto(mapCtx *direct.MapContext, in *krm.BigqueryReservationReservationSpec) *pb.Reservation {
	if in == nil {
		return nil
	}
	out := &pb.Reservation{}
	out.SlotCapacity = direct.ValueOf(in.SlotCapacity)
	out.IgnoreIdleSlots = direct.ValueOf(in.IgnoreIdleSlots)
	out.Autoscale = AutoscaleSpec_ToProto(mapCtx, in.Autoscale)
	out.Concurrency = direct.ValueOf(in.Concurrency)
	// MISSING: MultiRegionAuxiliary. Deprecated in v1beta1
	out.Edition = direct.Enum_ToProto[pb.Edition](mapCtx, in.Edition)
	out.SecondaryLocation = direct.ValueOf(in.SecondaryLocation)
	return out
}

func BigqueryReservationReservationSpec_FromProto(mapCtx *direct.MapContext, in *pb.Reservation) *krm.BigqueryReservationReservationSpec {
	if in == nil {
		return nil
	}
	out := &krm.BigqueryReservationReservationSpec{}
	out.SlotCapacity = direct.LazyPtr(in.GetSlotCapacity())
	out.IgnoreIdleSlots = direct.LazyPtr(in.GetIgnoreIdleSlots())
	out.Autoscale = AutoscaleSpec_FromProto(mapCtx, in.GetAutoscale())
	out.Concurrency = direct.LazyPtr(in.GetConcurrency())
	out.Edition = direct.Enum_FromProto(mapCtx, in.GetEdition())
	out.SecondaryLocation = direct.LazyPtr(in.GetSecondaryLocation())
	return out
}

func BigqueryReservationReservationObservedState_ToProto(mapCtx *direct.MapContext, in *krm.BigqueryReservationReservationObservedState) *pb.Reservation {
	if in == nil {
		return nil
	}
	out := &pb.Reservation{}
	out.Autoscale = AutoscaleObservedState_ToProto(mapCtx, in.CurrentSlots)
	out.PrimaryLocation = direct.ValueOf(in.PrimaryLocation)
	out.SecondaryLocation = direct.ValueOf(in.SecondaryLocation)
	out.OriginalPrimaryLocation = direct.ValueOf(in.OriginalPrimaryLocation)
	return out
}

func BigqueryReservationReservationObservedState_FromProto(mapCtx *direct.MapContext, in *pb.Reservation) *krm.BigqueryReservationReservationObservedState {
	if in == nil {
		return nil
	}
	out := &krm.BigqueryReservationReservationObservedState{}
	out.PrimaryLocation = direct.LazyPtr(in.GetPrimaryLocation())
	out.SecondaryLocation = direct.LazyPtr(in.GetSecondaryLocation())
	out.OriginalPrimaryLocation = direct.LazyPtr(in.GetOriginalPrimaryLocation())
	out.CurrentSlots = AutoscaleObservedState_FromProto(mapCtx, in.Autoscale)
	return out
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
	out.CurrentSlots = direct.LazyPtr(in.GetCurrentSlots())
	return out
}
func AutoscaleObservedState_ToProto(mapCtx *direct.MapContext, in *krm.AutoscaleObservedState) *pb.Reservation_Autoscale {
	if in == nil {
		return nil
	}
	out := &pb.Reservation_Autoscale{}
	out.CurrentSlots = direct.ValueOf(in.CurrentSlots)
	return out
}
