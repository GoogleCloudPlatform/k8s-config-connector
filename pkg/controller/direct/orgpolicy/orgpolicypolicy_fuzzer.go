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
// proto.message: google.cloud.orgpolicy.v2.Policy
// api.group: orgpolicy.cnrm.cloud.google.com

package orgpolicy

import (
	pb "cloud.google.com/go/orgpolicy/apiv2/orgpolicypb"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/fuzztesting"
)

func init() {
	fuzztesting.RegisterKRMFuzzer(orgPolicyPolicyFuzzer())
}

func orgPolicyPolicyFuzzer() fuzztesting.KRMFuzzer {
	f := fuzztesting.NewKRMTypedFuzzer(&pb.Policy{},
		OrgPolicyPolicySpec_FromProto, OrgPolicyPolicySpec_ToProto,
		OrgPolicyPolicyObservedState_FromProto, OrgPolicyPolicyObservedState_ToProto,
	)

	f.SpecFields.Insert(".spec")
	f.SpecFields.Insert(".dry_run_spec")

	f.StatusFields.Insert(".spec.update_time")
	f.StatusFields.Insert(".dry_run_spec.update_time")

	f.UnimplementedFields.Insert(".name")
	f.UnimplementedFields.Insert(".alternate") // deprecated field
	f.UnimplementedFields.Insert(".dry_run_spec.etag")

	f.Unimplemented_Etag()
	f.UnimplementedFields.Insert(".spec.etag")

	// PolicyRule.parameters is a typed KRM struct that only models the
	// 23 known managed-constraint parameter keys; random structpb.Struct
	// values produced by the fuzzer will lose unknown keys on the
	// FromProto -> ToProto round-trip. Skip it here. The "[]" suffix
	// matches the repeated rules field; see pkg/test/fuzz/generate.go.
	f.UnimplementedFields.Insert(".spec.rules[].parameters")
	f.UnimplementedFields.Insert(".dry_run_spec.rules[].parameters")

	return f
}
