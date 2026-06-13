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
	f.Unimplemented_NotYetTriaged(".delete_after_duration")
	f.Unimplemented_NotYetTriaged(".delete_at_time")
	f.Unimplemented_NotYetTriaged(".deployment_type")
	f.Unimplemented_NotYetTriaged(".enable_emergent_maintenance")
	f.Unimplemented_NotYetTriaged(".linked_commitments")
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

	// FilterSpec is required because although the KRM schema defines Count, InUseCount, and DiskSizeGb as int32,
	// the GCP Go client library (computepb) represents these fields as int64.
	// As a result, the fuzzer (which operates on the pb.Reservation struct) generates random int64 values
	// that can exceed the int32 range and truncate during round-trip translation. We normalize them to the int32
	// range here to ensure a lossless round-trip check.
	f.FilterSpec = func(in *pb.Reservation) {
		if in.SpecificReservation != nil {
			if in.SpecificReservation.Count != nil {
				val := int64(int32(*in.SpecificReservation.Count))
				in.SpecificReservation.Count = &val
			}
			if in.SpecificReservation.InUseCount != nil {
				val := int64(int32(*in.SpecificReservation.InUseCount))
				in.SpecificReservation.InUseCount = &val
			}
			if in.SpecificReservation.InstanceProperties != nil {
				for _, disk := range in.SpecificReservation.InstanceProperties.LocalSsds {
					if disk.DiskSizeGb != nil {
						val := int64(int32(*disk.DiskSizeGb))
						disk.DiskSizeGb = &val
					}
				}
			}
		}
	}

	f.FilterStatus = func(in *pb.Reservation) {
		// Nothing additional needed for status
	}

	return f
}
