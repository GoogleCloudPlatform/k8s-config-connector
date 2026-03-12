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

package compute

import (
	computepbv1beta "cloud.google.com/go/compute/apiv1beta/computepb"
	krm "github.com/GoogleCloudPlatform/k8s-config-connector/apis/compute/v1beta1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct"
)

func ComputeFutureReservationObservedState_v1beta1_FromProto(mapCtx *direct.MapContext, in *computepbv1beta.FutureReservation) *krm.ComputeFutureReservationObservedState {
	if in == nil {
		return nil
	}
	out := &krm.ComputeFutureReservationObservedState{}
	out.CreationTimestamp = in.CreationTimestamp
	out.ID = in.Id
	out.Kind = in.Kind
	out.SelfLink = in.SelfLink
	out.SelfLinkWithID = in.SelfLinkWithId
	out.Zone = in.Zone

	status := in.GetStatus()
	if status != nil {
		out.AmendmentStatus = status.AmendmentStatus
		out.AutoCreatedReservations = status.AutoCreatedReservations
		out.ExistingMatchingUsageInfo = FutureReservationStatusExistingMatchingUsageInfo_v1beta1_FromProto(mapCtx, status.GetExistingMatchingUsageInfo())
		out.FulfilledCount = status.FulfilledCount
		out.LastKnownGoodState = FutureReservationStatusLastKnownGoodState_v1beta1_FromProto(mapCtx, status.GetLastKnownGoodState())
		out.LockTime = status.LockTime
		out.ProcurementStatus = status.ProcurementStatus
		out.SpecificSkuProperties = FutureReservationStatusSpecificSkuProperties_v1beta1_FromProto(mapCtx, status.GetSpecificSkuProperties())
	}

	return out
}

func ComputeFutureReservationObservedState_v1beta1_ToProto(mapCtx *direct.MapContext, in *krm.ComputeFutureReservationObservedState) *computepbv1beta.FutureReservation {
	if in == nil {
		return nil
	}
	out := &computepbv1beta.FutureReservation{}
	out.CreationTimestamp = in.CreationTimestamp
	out.Id = in.ID
	out.Kind = in.Kind
	out.SelfLink = in.SelfLink
	out.SelfLinkWithId = in.SelfLinkWithID
	out.Zone = in.Zone

	status := &computepbv1beta.FutureReservationStatus{}
	status.AmendmentStatus = in.AmendmentStatus
	status.AutoCreatedReservations = in.AutoCreatedReservations
	status.ExistingMatchingUsageInfo = FutureReservationStatusExistingMatchingUsageInfo_v1beta1_ToProto(mapCtx, in.ExistingMatchingUsageInfo)
	status.FulfilledCount = in.FulfilledCount
	status.LastKnownGoodState = FutureReservationStatusLastKnownGoodState_v1beta1_ToProto(mapCtx, in.LastKnownGoodState)
	status.LockTime = in.LockTime
	status.ProcurementStatus = in.ProcurementStatus
	status.SpecificSkuProperties = FutureReservationStatusSpecificSkuProperties_v1beta1_ToProto(mapCtx, in.SpecificSkuProperties)
	out.Status = status

	return out
}
