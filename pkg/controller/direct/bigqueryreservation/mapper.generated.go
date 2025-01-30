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

func Reservation_FromProto(mapCtx *direct.MapContext, in *pb.Reservation) *krm.Reservation {
	if in == nil {
		return nil
	}
	out := &krm.Reservation{}
	out.Name = direct.LazyPtr(in.GetName())
	out.SlotCapacity = direct.LazyPtr(in.GetSlotCapacity())
	out.IgnoreIdleSlots = direct.LazyPtr(in.GetIgnoreIdleSlots())
	out.Autoscale = Reservation_Autoscale_FromProto(mapCtx, in.GetAutoscale())
	out.Concurrency = direct.LazyPtr(in.GetConcurrency())
	// MISSING: CreationTime
	// MISSING: UpdateTime
	out.MultiRegionAuxiliary = direct.LazyPtr(in.GetMultiRegionAuxiliary())
	out.Edition = direct.Enum_FromProto(mapCtx, in.GetEdition())
	out.PrimaryLocation = direct.LazyPtr(in.GetPrimaryLocation())
	out.SecondaryLocation = direct.LazyPtr(in.GetSecondaryLocation())
	out.OriginalPrimaryLocation = direct.LazyPtr(in.GetOriginalPrimaryLocation())
	return out
}
func Reservation_ToProto(mapCtx *direct.MapContext, in *krm.Reservation) *pb.Reservation {
	if in == nil {
		return nil
	}
	out := &pb.Reservation{}
	out.Name = direct.ValueOf(in.Name)
	out.SlotCapacity = direct.ValueOf(in.SlotCapacity)
	out.IgnoreIdleSlots = direct.ValueOf(in.IgnoreIdleSlots)
	out.Autoscale = Reservation_Autoscale_ToProto(mapCtx, in.Autoscale)
	out.Concurrency = direct.ValueOf(in.Concurrency)
	// MISSING: CreationTime
	// MISSING: UpdateTime
	out.MultiRegionAuxiliary = direct.ValueOf(in.MultiRegionAuxiliary)
	out.Edition = direct.Enum_ToProto[pb.Edition](mapCtx, in.Edition)
	out.PrimaryLocation = direct.ValueOf(in.PrimaryLocation)
	out.SecondaryLocation = direct.ValueOf(in.SecondaryLocation)
	out.OriginalPrimaryLocation = direct.ValueOf(in.OriginalPrimaryLocation)
	return out
}
func ReservationObservedState_FromProto(mapCtx *direct.MapContext, in *pb.Reservation) *krm.ReservationObservedState {
	if in == nil {
		return nil
	}
	out := &krm.ReservationObservedState{}
	// MISSING: Name
	// MISSING: SlotCapacity
	// MISSING: IgnoreIdleSlots
	out.Autoscale = Reservation_AutoscaleObservedState_FromProto(mapCtx, in.GetAutoscale())
	// MISSING: Concurrency
	out.CreationTime = direct.StringTimestamp_FromProto(mapCtx, in.GetCreationTime())
	out.UpdateTime = direct.StringTimestamp_FromProto(mapCtx, in.GetUpdateTime())
	// MISSING: MultiRegionAuxiliary
	// MISSING: Edition
	// MISSING: PrimaryLocation
	// MISSING: SecondaryLocation
	// MISSING: OriginalPrimaryLocation
	return out
}
func ReservationObservedState_ToProto(mapCtx *direct.MapContext, in *krm.ReservationObservedState) *pb.Reservation {
	if in == nil {
		return nil
	}
	out := &pb.Reservation{}
	// MISSING: Name
	// MISSING: SlotCapacity
	// MISSING: IgnoreIdleSlots
	out.Autoscale = Reservation_AutoscaleObservedState_ToProto(mapCtx, in.Autoscale)
	// MISSING: Concurrency
	out.CreationTime = direct.StringTimestamp_ToProto(mapCtx, in.CreationTime)
	out.UpdateTime = direct.StringTimestamp_ToProto(mapCtx, in.UpdateTime)
	// MISSING: MultiRegionAuxiliary
	// MISSING: Edition
	// MISSING: PrimaryLocation
	// MISSING: SecondaryLocation
	// MISSING: OriginalPrimaryLocation
	return out
}
func Reservation_Autoscale_FromProto(mapCtx *direct.MapContext, in *pb.Reservation_Autoscale) *krm.Reservation_Autoscale {
	if in == nil {
		return nil
	}
	out := &krm.Reservation_Autoscale{}
	// MISSING: CurrentSlots
	out.MaxSlots = direct.LazyPtr(in.GetMaxSlots())
	return out
}
func Reservation_Autoscale_ToProto(mapCtx *direct.MapContext, in *krm.Reservation_Autoscale) *pb.Reservation_Autoscale {
	if in == nil {
		return nil
	}
	out := &pb.Reservation_Autoscale{}
	// MISSING: CurrentSlots
	out.MaxSlots = direct.ValueOf(in.MaxSlots)
	return out
}
func Reservation_AutoscaleObservedState_FromProto(mapCtx *direct.MapContext, in *pb.Reservation_Autoscale) *krm.Reservation_AutoscaleObservedState {
	if in == nil {
		return nil
	}
	out := &krm.Reservation_AutoscaleObservedState{}
	out.CurrentSlots = direct.LazyPtr(in.GetCurrentSlots())
	// MISSING: MaxSlots
	return out
}
func Reservation_AutoscaleObservedState_ToProto(mapCtx *direct.MapContext, in *krm.Reservation_AutoscaleObservedState) *pb.Reservation_Autoscale {
	if in == nil {
		return nil
	}
	out := &pb.Reservation_Autoscale{}
	out.CurrentSlots = direct.ValueOf(in.CurrentSlots)
	// MISSING: MaxSlots
	return out
}
