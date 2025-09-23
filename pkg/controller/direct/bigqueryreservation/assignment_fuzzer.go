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
// proto.message: google.cloud.bigquery.reservation.v1.Assignment
// api.group: bigqueryreservation.cnrm.cloud.google.com

package bigqueryreservation

import (
	pb "cloud.google.com/go/bigquery/reservation/apiv1/reservationpb"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/fuzztesting"
)

func init() {
	fuzztesting.RegisterKRMSpecFuzzer(BigQueryReservationAssignmentFuzzer())
}

func BigQueryReservationAssignmentFuzzer() fuzztesting.KRMFuzzer {
	f := fuzztesting.NewKRMTypedFuzzer(&pb.Assignment{},
		BigQueryReservationAssignmentSpec_FromProto, BigQueryReservationAssignmentSpec_ToProto,
		BigQueryReservationAssignmentObservedState_FromProto, BigQueryReservationAssignmentObservedState_ToProto,
	)

	f.SpecFields.Insert(".job_type")

	f.StatusFields.Insert(".state")

	f.UnimplementedFields.Insert(".name")
	f.UnimplementedFields.Insert(".assignee") // Assignee has to be in a specific format, not a random string.
	f.UnimplementedFields.Insert(".enable_gemini_in_bigquery")

	return f
}
