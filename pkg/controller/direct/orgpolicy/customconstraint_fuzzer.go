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
// proto.message: google.cloud.orgpolicy.v2.CustomConstraint
// api.group: orgpolicy.cnrm.cloud.google.com

package orgpolicy

import (
	pb "cloud.google.com/go/orgpolicy/apiv2/orgpolicypb"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/fuzztesting"
)

func init() {
	fuzztesting.RegisterKRMFuzzer(orgPolicyCustomConstraintFuzzer())
}

func orgPolicyCustomConstraintFuzzer() fuzztesting.KRMFuzzer {
	f := fuzztesting.NewKRMTypedFuzzer(&pb.CustomConstraint{},
		OrgPolicyCustomConstraintSpec_FromProto, OrgPolicyCustomConstraintSpec_ToProto,
		OrgPolicyCustomConstraintObservedState_FromProto, OrgPolicyCustomConstraintObservedState_ToProto,
	)

	f.SpecFields.Insert(".resource_types")
	f.SpecFields.Insert(".method_types")
	f.SpecFields.Insert(".condition")
	f.SpecFields.Insert(".action_type")
	f.SpecFields.Insert(".display_name")
	f.SpecFields.Insert(".description")

	f.StatusFields.Insert(".update_time")

	f.UnimplementedFields.Insert(".name") // special field

	return f
}
