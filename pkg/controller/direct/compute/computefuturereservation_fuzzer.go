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

// +tool:fuzz-gen
// proto.message: google.cloud.compute.v1.FutureReservation
// api.group: compute.cnrm.cloud.google.com

package compute

import (
	pb "cloud.google.com/go/compute/apiv1/computepb"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/fuzztesting"
)

func init() {
	fuzztesting.RegisterKRMFuzzer(computeFutureReservationFuzzer())
}

func computeFutureReservationFuzzer() fuzztesting.KRMFuzzer {
	f := fuzztesting.NewKRMTypedFuzzer(&pb.FutureReservation{},
		ComputeFutureReservationSpec_v1alpha1_FromProto, ComputeFutureReservationSpec_v1alpha1_ToProto,
		ComputeFutureReservationObservedState_v1alpha1_FromProto, ComputeFutureReservationObservedState_v1alpha1_ToProto,
	)

	// Spec fields
	f.SpecField(".aggregate_reservation")
	f.SpecField(".auto_created_reservations_delete_time")
	f.SpecField(".auto_created_reservations_duration")
	f.SpecField(".auto_delete_auto_created_reservations")
	f.SpecField(".commitment_info")
	f.SpecField(".deployment_type")
	f.SpecField(".description")
	f.SpecField(".enable_emergent_maintenance")
	f.SpecField(".name_prefix")
	f.SpecField(".planning_status")
	f.SpecField(".reservation_mode")
	f.SpecField(".reservation_name")
	f.SpecField(".scheduling_type")
	f.SpecField(".share_settings")
	f.SpecField(".share_settings.project_map")
	f.SpecField(".specific_reservation_required")
	f.SpecField(".specific_sku_properties")
	f.SpecField(".time_window")

	// Status fields
	f.StatusField(".aggregate_reservation")
	f.StatusField(".creation_timestamp")
	f.StatusField(".id")
	f.StatusField(".kind")
	f.StatusField(".self_link")
	f.StatusField(".self_link_with_id")
	f.StatusField(".status")
	f.StatusField(".status.last_known_good_state.future_reservation_specs.share_settings.project_map")
	f.StatusField(".zone")

	// Identity field
	f.Unimplemented_Identity(".name")

	// Unimplemented / Not Yet Triaged fields
	f.Unimplemented_NotYetTriaged(".confidential_compute_type")
	f.Unimplemented_NotYetTriaged(".params")

	return f
}
