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

package bigqueryreservation

import (
	pb "cloud.google.com/go/bigquery/reservation/apiv1/reservationpb"
	krm "github.com/GoogleCloudPlatform/k8s-config-connector/apis/bigqueryreservation/v1alpha1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct"
)

func BigQueryReservationReservationGroupSpec_v1alpha1_FromProto(mapCtx *direct.MapContext, in *pb.ReservationGroup) *krm.BigQueryReservationReservationGroupSpec {
	if in == nil {
		return nil
	}
	out := &krm.BigQueryReservationReservationGroupSpec{}
	if in.GetParentGroup() != "" {
		out.ReservationGroupRef = &krm.BigQueryReservationReservationGroupRef{External: in.GetParentGroup()}
	}
	return out
}

func BigQueryReservationReservationGroupSpec_v1alpha1_ToProto(mapCtx *direct.MapContext, in *krm.BigQueryReservationReservationGroupSpec) *pb.ReservationGroup {
	if in == nil {
		return nil
	}
	out := &pb.ReservationGroup{}
	if in.ReservationGroupRef != nil {
		out.ParentGroup = in.ReservationGroupRef.External
	}
	return out
}
