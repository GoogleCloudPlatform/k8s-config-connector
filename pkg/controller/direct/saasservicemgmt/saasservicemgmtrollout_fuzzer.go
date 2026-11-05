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
// proto.message: google.cloud.saasplatform.saasservicemgmt.v1beta1.Rollout
// api.group: saasservicemgmt.cnrm.cloud.google.com

package saasservicemgmt

import (
	pb "cloud.google.com/go/saasplatform/saasservicemgmt/apiv1beta1/saasservicemgmtpb"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/fuzztesting"
)

func init() {
	fuzztesting.RegisterKRMFuzzer(saasServiceMgmtRolloutFuzzer())
}

func saasServiceMgmtRolloutFuzzer() fuzztesting.KRMFuzzer {
	f := fuzztesting.NewKRMTypedFuzzer(&pb.Rollout{},
		SaaSServiceMgmtRolloutSpec_FromProto, SaaSServiceMgmtRolloutSpec_ToProto,
		SaaSServiceMgmtRolloutObservedState_FromProto, SaaSServiceMgmtRolloutObservedState_ToProto,
	)

	// Identity fields that are not in KRM fields
	f.Unimplemented_Identity(".name")

	// Spec fields to fuzz
	f.SpecField(".release")
	f.SpecField(".rollout_orchestration_strategy")
	f.SpecField(".unit_filter")
	f.SpecField(".rollout_kind")
	f.SpecField(".control")

	// Status fields (ObservedState) to fuzz
	f.StatusField(".start_time")
	f.StatusField(".end_time")
	f.StatusField(".state")
	f.StatusField(".state_message")
	f.StatusField(".state_transition_time")
	f.StatusField(".root_rollout")
	f.StatusField(".parent_rollout")
	f.StatusField(".stats")
	f.StatusField(".uid")
	f.StatusField(".create_time")
	f.StatusField(".update_time")

	// Unimplemented or ignored fields
	f.Unimplemented_LabelsAnnotations(".labels")
	f.Unimplemented_LabelsAnnotations(".annotations")
	f.Unimplemented_Etag()
	f.Unimplemented_NotYetTriaged(".delete_time")
	f.Unimplemented_NotYetTriaged(".effective_unit_filter")
	f.Unimplemented_NotYetTriaged(".stats.estimated_total_unit_count")

	return f
}
