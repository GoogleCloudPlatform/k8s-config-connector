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
// proto.message: google.cloud.saasplatform.saasservicemgmt.v1beta1.RolloutKind
// api.group: saasservicemgmt.cnrm.cloud.google.com

package saasservicemgmt

import (
	pb "cloud.google.com/go/saasplatform/saasservicemgmt/apiv1beta1/saasservicemgmtpb"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/fuzztesting"
)

func init() {
	fuzztesting.RegisterKRMFuzzer(saasServiceMgmtRolloutKindFuzzer())
}

func saasServiceMgmtRolloutKindFuzzer() fuzztesting.KRMFuzzer {
	f := fuzztesting.NewKRMTypedFuzzer(&pb.RolloutKind{},
		SaaSServiceMgmtRolloutKindSpec_FromProto, SaaSServiceMgmtRolloutKindSpec_ToProto,
		SaaSServiceMgmtRolloutKindObservedState_FromProto, SaaSServiceMgmtRolloutKindObservedState_ToProto,
	)

	// Identity fields that are not in KRM fields
	f.Unimplemented_Identity(".name")

	// Spec fields to fuzz
	f.SpecField(".unit_kind")
	f.SpecField(".rollout_orchestration_strategy")
	f.SpecField(".unit_filter")
	f.SpecField(".update_unit_kind_strategy")
	f.SpecField(".error_budget")

	// Status fields (ObservedState) to fuzz
	f.StatusField(".uid")
	f.StatusField(".etag")
	f.StatusField(".create_time")
	f.StatusField(".update_time")

	// Unimplemented or ignored fields
	f.Unimplemented_LabelsAnnotations(".labels")
	f.Unimplemented_LabelsAnnotations(".annotations")

	return f
}
