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
// proto.message: google.cloud.bigquery.reservation.v1.ReservationGroup
// api.group: bigqueryreservation.cnrm.cloud.google.com

package bigqueryreservation

import (
	pb "cloud.google.com/go/bigquery/reservation/apiv1/reservationpb"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/fuzztesting"
)

func init() {
	fuzztesting.RegisterKRMSpecFuzzer(reservationGroupFuzzer())
}

func reservationGroupFuzzer() fuzztesting.KRMFuzzer {
	f := fuzztesting.NewKRMTypedFuzzer(&pb.ReservationGroup{},
		BigQueryReservationReservationGroupSpec_v1alpha1_FromProto, BigQueryReservationReservationGroupSpec_v1alpha1_ToProto,
		BigQueryReservationReservationGroupObservedState_v1alpha1_FromProto, BigQueryReservationReservationGroupObservedState_v1alpha1_ToProto,
	)
	f.IdentityField(".name")
	return f
}
