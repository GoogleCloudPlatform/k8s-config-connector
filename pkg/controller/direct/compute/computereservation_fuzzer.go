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
// proto.message: google.cloud.compute.v1.Reservation
// api.group: compute.cnrm.cloud.google.com

package compute

import (
	pb "cloud.google.com/go/compute/apiv1/computepb"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/fuzztesting"
)

func init() {
	fuzztesting.RegisterKRMFuzzer(computeReservationFuzzer())
}

func computeReservationFuzzer() fuzztesting.KRMFuzzer {
	f := fuzztesting.NewKRMTypedFuzzer(&pb.Reservation{},
		ComputeReservationSpec_v1beta1_FromProto, ComputeReservationSpec_v1beta1_ToProto,
		ComputeReservationStatus_v1beta1_FromProto, ComputeReservationStatus_v1beta1_ToProto,
	)

	// Spec fields
	f.SpecField(".description")
	f.SpecField(".share_settings")
	f.SpecField(".specific_reservation")
	f.SpecField(".specific_reservation_required")
	f.SpecField(".zone")

	// Status fields
	f.StatusField(".commitment")
	f.StatusField(".creation_timestamp")
	f.StatusField(".self_link")
	f.StatusField(".status")

	// Unimplemented fields / URL/identity
	f.Unimplemented_Identity(".name")
	f.Unimplemented_Internal(".kind")

	// Unimplemented / other fields
	f.Unimplemented_NotYetTriaged(".id")
	f.Unimplemented_NotYetTriaged(".advanced_deployment_control")
	f.Unimplemented_NotYetTriaged(".aggregate_reservation")
	f.Unimplemented_NotYetTriaged(".confidential_compute_type")
	f.Unimplemented_NotYetTriaged(".delete_after_duration")
	f.Unimplemented_NotYetTriaged(".delete_at_time")
	f.Unimplemented_NotYetTriaged(".deployment_type")
	f.Unimplemented_NotYetTriaged(".early_access_maintenance")
	f.Unimplemented_NotYetTriaged(".enable_emergent_maintenance")
	f.Unimplemented_NotYetTriaged(".linked_commitments")
	f.Unimplemented_NotYetTriaged(".params")
	f.Unimplemented_NotYetTriaged(".protection_tier")
	f.Unimplemented_NotYetTriaged(".reservation_sharing_policy")
	f.Unimplemented_NotYetTriaged(".resource_policies")
	f.Unimplemented_NotYetTriaged(".resource_status")
	f.Unimplemented_NotYetTriaged(".satisfies_pzs")
	f.Unimplemented_NotYetTriaged(".scheduling_type")

	// Nested unimplemented fields in Spec structures
	f.Unimplemented_NotYetTriaged(".specific_reservation.assured_count")
	f.Unimplemented_NotYetTriaged(".specific_reservation.source_instance_template")
	f.Unimplemented_NotYetTriaged(".specific_reservation.instance_properties.location_hint")

	f.FilterSpec = func(in *pb.Reservation) {
		// Nothing additional needed for spec
	}

	f.FilterStatus = func(in *pb.Reservation) {
		// Nothing additional needed for status
	}

	return f
}
