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
	pb "cloud.google.com/go/compute/apiv1beta/computepb"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/fuzztesting"
)

func init() {
	fuzztesting.RegisterKRMFuzzer(ComputeFutureReservationFuzzer())
}

func ComputeFutureReservationFuzzer() fuzztesting.KRMFuzzer {
	f := fuzztesting.NewKRMTypedFuzzer(&pb.FutureReservation{},
		FutureReservationSpec_FromProto, FutureReservationSpec_ToProto,
		FutureReservationObservedState_FromProto, FutureReservationObservedState_ToProto,
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

	f.UnimplementedFields.Insert(".status")
	return f
}
