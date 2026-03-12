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

// +tool:fuzz-gen
// proto.message: google.cloud.bigquery.reservation.v1.Reservation
// api.group: bigqueryreservation.cnrm.cloud.google.com

package bigqueryreservation

import (
	pb "cloud.google.com/go/bigquery/reservation/apiv1/reservationpb"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/fuzztesting"
)

func init() {
	fuzztesting.RegisterKRMSpecFuzzer(BigQueryReservationReservationFuzzer())
}

func BigQueryReservationReservationFuzzer() fuzztesting.KRMFuzzer {
	f := fuzztesting.NewKRMTypedFuzzer(&pb.Reservation{},
		BigQueryReservationReservationSpec_FromProto, BigQueryReservationReservationSpec_ToProto,
		BigQueryReservationReservationObservedState_FromProto, BigQueryReservationReservationObservedState_ToProto,
	)

	f.SpecFields.Insert(".slot_capacity")
	f.SpecFields.Insert(".ignore_idle_slots")
	f.SpecFields.Insert(".autoscale")
	f.SpecFields.Insert(".concurrency")
	f.SpecFields.Insert(".edition")
	f.SpecFields.Insert(".secondary_location")

	f.StatusFields.Insert(".primary_location")
	f.StatusFields.Insert(".original_primary_location")

	f.IdentityField(".name")

	f.UnimplementedFields.Insert(".creation_time")
	f.UnimplementedFields.Insert(".update_time")
	f.UnimplementedFields.Insert(".multi_region_auxiliary")
	f.UnimplementedFields.Insert(".scaling_mode")
	f.UnimplementedFields.Insert(".autoscale.current_slots")
	f.UnimplementedFields.Insert(".replication_status")
	f.UnimplementedFields.Insert(".max_slots")

	f.Unimplemented_LabelsAnnotations(".labels")
	f.Unimplemented_NotYetTriaged(".reservation_group")
	f.Unimplemented_NotYetTriaged(".scheduling_policy")

	return f
}
