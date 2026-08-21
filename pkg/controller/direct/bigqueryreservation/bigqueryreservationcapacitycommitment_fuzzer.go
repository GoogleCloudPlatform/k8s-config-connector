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
// proto.message: google.cloud.bigquery.reservation.v1.CapacityCommitment
// api.group: bigqueryreservation.cnrm.cloud.google.com

package bigqueryreservation

import (
	pb "cloud.google.com/go/bigquery/reservation/apiv1/reservationpb"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/fuzztesting"
)

func init() {
	fuzztesting.RegisterKRMSpecFuzzer(BigQueryReservationCapacityCommitmentFuzzer())
}

func BigQueryReservationCapacityCommitmentFuzzer() fuzztesting.KRMFuzzer {
	f := fuzztesting.NewKRMTypedFuzzer(&pb.CapacityCommitment{},
		BigQueryReservationCapacityCommitmentSpec_v1alpha1_FromProto, BigQueryReservationCapacityCommitmentSpec_v1alpha1_ToProto,
		BigQueryReservationCapacityCommitmentStatus_v1alpha1_FromProto, BigQueryReservationCapacityCommitmentStatus_v1alpha1_ToProto,
	)

	f.SpecField(".slot_count")
	f.SpecField(".plan")
	f.SpecField(".renewal_plan")
	f.SpecField(".edition")

	f.StatusField(".state")
	f.StatusField(".commitment_start_time")
	f.StatusField(".commitment_end_time")

	f.IdentityField(".name")

	f.Unimplemented_NotYetTriaged(".failure_status")
	f.Unimplemented_NotYetTriaged(".multi_region_auxiliary")
	f.Unimplemented_NotYetTriaged(".is_flat_rate")

	return f
}
