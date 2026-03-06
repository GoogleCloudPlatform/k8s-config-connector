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
	computepb "cloud.google.com/go/compute/apiv1beta/computepb"
	krm "github.com/GoogleCloudPlatform/k8s-config-connector/apis/compute/v1beta1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct"
)

func ComputeFutureReservationObservedState_v1beta1_FromProto(mapCtx *direct.MapContext, in *computepb.FutureReservation) *krm.ComputeFutureReservationObservedState {
	if in == nil {
		return nil
	}
	out := &krm.ComputeFutureReservationObservedState{}
	out.CreationTimestamp = in.CreationTimestamp
	out.ID = in.Id
	out.Kind = in.Kind
	out.SelfLink = in.SelfLink
	out.SelfLinkWithID = in.SelfLinkWithId
	out.SpecificSkuProperties = FutureReservationStatusSpecificSkuProperties_v1beta1_FromProto(mapCtx, in.GetStatus().GetSpecificSkuProperties())
	out.Zone = in.Zone
	out.AmendmentStatus = in.GetStatus().AmendmentStatus
	out.AutoCreatedReservations = in.GetStatus().AutoCreatedReservations
	out.ExistingMatchingUsageInfo = FutureReservationStatusExistingMatchingUsageInfo_v1beta1_FromProto(mapCtx, in.GetStatus().GetExistingMatchingUsageInfo())
	out.FulfilledCount = in.GetStatus().FulfilledCount
	out.LastKnownGoodState = FutureReservationStatusLastKnownGoodState_v1beta1_FromProto(mapCtx, in.GetStatus().GetLastKnownGoodState())
	out.LockTime = in.GetStatus().LockTime
	out.ProcurementStatus = in.GetStatus().ProcurementStatus
	return out
}

func ComputeFutureReservationObservedState_v1beta1_ToProto(mapCtx *direct.MapContext, in *krm.ComputeFutureReservationObservedState) *computepb.FutureReservation {
	if in == nil {
		return nil
	}
	out := &computepb.FutureReservation{}
	out.CreationTimestamp = in.CreationTimestamp
	out.Id = in.ID
	out.Kind = in.Kind
	out.SelfLink = in.SelfLink
	out.SelfLinkWithId = in.SelfLinkWithID

	out.Status = &computepb.FutureReservationStatus{}
	out.Status.SpecificSkuProperties = FutureReservationStatusSpecificSkuProperties_v1beta1_ToProto(mapCtx, in.SpecificSkuProperties)
	out.Zone = in.Zone
	out.Status.AmendmentStatus = in.AmendmentStatus
	out.Status.AutoCreatedReservations = in.AutoCreatedReservations
	out.Status.ExistingMatchingUsageInfo = FutureReservationStatusExistingMatchingUsageInfo_v1beta1_ToProto(mapCtx, in.ExistingMatchingUsageInfo)
	out.Status.FulfilledCount = in.FulfilledCount
	out.Status.LastKnownGoodState = FutureReservationStatusLastKnownGoodState_v1beta1_ToProto(mapCtx, in.LastKnownGoodState)
	out.Status.LockTime = in.LockTime
	out.Status.ProcurementStatus = in.ProcurementStatus
	return out
}
