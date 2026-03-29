// Copyright 2024 Google LLC
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//	http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

// +tool:fuzz-gen
// proto.message: google.cloud.compute.v1beta1.FutureReservation
// api.group: compute.cnrm.cloud.google.com

package compute

import (
	computepb "cloud.google.com/go/compute/apiv1beta/computepb"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/fuzztesting"
	"google.golang.org/protobuf/proto"
)

func init() {
	fuzztesting.RegisterKRMFuzzer(computeFutureReservationFuzzer())
}

func computeFutureReservationFuzzer() fuzztesting.KRMFuzzer {
	f := fuzztesting.NewKRMTypedFuzzer(&computepb.FutureReservation{},
		ComputeFutureReservationSpec_computepb_v1beta1_FromProto, ComputeFutureReservationSpec_computepb_v1beta1_ToProto,
		ComputeFutureReservationObservedState_computepb_v1beta1_FromProto, ComputeFutureReservationObservedState_computepb_v1beta1_ToProto,
	)

	// Spec fields
	f.SpecFields.Insert(".aggregate_reservation")
	f.SpecFields.Insert(".auto_created_reservations_delete_time")
	f.SpecFields.Insert(".auto_created_reservations_duration")
	f.SpecFields.Insert(".auto_delete_auto_created_reservations")
	f.SpecFields.Insert(".commitment_info")
	f.SpecFields.Insert(".deployment_type")
	f.SpecFields.Insert(".description")
	f.SpecFields.Insert(".enable_emergent_maintenance")
	f.SpecFields.Insert(".name")
	f.SpecFields.Insert(".name_prefix")
	f.SpecFields.Insert(".planning_status")
	f.SpecFields.Insert(".reservation_mode")
	f.SpecFields.Insert(".reservation_name")
	f.SpecFields.Insert(".scheduling_type")
	f.SpecFields.Insert(".share_settings")
	f.SpecFields.Insert(".specific_reservation_required")
	f.SpecFields.Insert(".specific_sku_properties")
	f.SpecFields.Insert(".time_window")

	// Status/ObservedState fields
	f.StatusFields.Insert(".creation_timestamp")
	f.StatusFields.Insert(".id")
	f.StatusFields.Insert(".kind")
	f.StatusFields.Insert(".self_link")
	f.StatusFields.Insert(".self_link_with_id")
	f.StatusFields.Insert(".specific_sku_properties")
	f.StatusFields.Insert(".zone")

	f.FilterSpec = func(in *computepb.FutureReservation) {
		// Clear empty messages to avoid round-trip failures
		if in.AggregateReservation != nil {
			in.AggregateReservation.HostCount = nil
			in.AggregateReservation.InUseHostCount = nil
			in.AggregateReservation.InUseInstanceCount = nil
			if proto.Equal(in.AggregateReservation, &computepb.AllocationAggregateReservation{}) {
				in.AggregateReservation = nil
			}
		}
		if in.AutoCreatedReservationsDuration != nil && proto.Equal(in.AutoCreatedReservationsDuration, &computepb.Duration{}) {
			in.AutoCreatedReservationsDuration = nil
		}
		if in.CommitmentInfo != nil && proto.Equal(in.CommitmentInfo, &computepb.FutureReservationCommitmentInfo{}) {
			in.CommitmentInfo = nil
		}
		if in.ShareSettings != nil && proto.Equal(in.ShareSettings, &computepb.ShareSettings{}) {
			in.ShareSettings = nil
		}
		if in.SpecificSkuProperties != nil && proto.Equal(in.SpecificSkuProperties, &computepb.FutureReservationSpecificSKUProperties{}) {
			in.SpecificSkuProperties = nil
		}
		if in.TimeWindow != nil && proto.Equal(in.TimeWindow, &computepb.FutureReservationTimeWindow{}) {
			in.TimeWindow = nil
		}

		in.ProtectionTier = nil
	}

	f.FilterStatus = func(in *computepb.FutureReservation) {
		// Clear empty messages to avoid round-trip failures
		if in.Status != nil {
			if in.Status.ExistingMatchingUsageInfo != nil && proto.Equal(in.Status.ExistingMatchingUsageInfo, &computepb.FutureReservationStatusExistingMatchingUsageInfo{}) {
				in.Status.ExistingMatchingUsageInfo = nil
			}
			if in.Status.LastKnownGoodState != nil && proto.Equal(in.Status.LastKnownGoodState, &computepb.FutureReservationStatusLastKnownGoodState{}) {
				in.Status.LastKnownGoodState = nil
			}
			if in.Status.SpecificSkuProperties != nil && proto.Equal(in.Status.SpecificSkuProperties, &computepb.FutureReservationStatusSpecificSKUProperties{}) {
				in.Status.SpecificSkuProperties = nil
			}
			if proto.Equal(in.Status, &computepb.FutureReservationStatus{}) {
				in.Status = nil
			}
		}

		if in.AggregateReservation != nil {
			in.AggregateReservation.HostCount = nil
			in.AggregateReservation.InUseHostCount = nil
			in.AggregateReservation.InUseInstanceCount = nil
			if proto.Equal(in.AggregateReservation, &computepb.AllocationAggregateReservation{}) {
				in.AggregateReservation = nil
			}
		}

		in.ProtectionTier = nil
	}

	f.UnimplementedFields.Insert(".name")
	f.UnimplementedFields.Insert(".status")
	f.UnimplementedFields.Insert(".share_settings.project_map")
	f.UnimplementedFields.Insert(".protection_tier")
	f.UnimplementedFields.Insert(".aggregate_reservation.host_count")
	f.UnimplementedFields.Insert(".aggregate_reservation.in_use_host_count")
	f.UnimplementedFields.Insert(".aggregate_reservation.in_use_instance_count")

	return f
}
