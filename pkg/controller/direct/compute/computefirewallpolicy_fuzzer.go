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
// proto.message: google.cloud.compute.v1.FirewallPolicy
// api.group: compute.cnrm.cloud.google.com

package compute

import (
	pb "cloud.google.com/go/compute/apiv1/computepb"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/fuzztesting"
)

func init() {
	fuzztesting.RegisterKRMFuzzer(computeFirewallPolicyFuzzer())
}

func computeFirewallPolicyFuzzer() fuzztesting.KRMFuzzer {
	f := fuzztesting.NewKRMTypedFuzzer(&pb.FirewallPolicy{},
		ComputeFirewallPolicySpec_v1beta1_FromProto, ComputeFirewallPolicySpec_v1beta1_ToProto,
		ComputeFirewallPolicyStatus_v1beta1_FromProto, ComputeFirewallPolicyStatus_v1beta1_ToProto,
	)

	// Spec fields
	f.SpecField(".description")
	f.SpecField(".short_name")

	// Status fields
	f.StatusField(".creation_timestamp")
	f.StatusField(".fingerprint")
	f.StatusField(".id")
	f.StatusField(".rule_tuple_count")
	f.StatusField(".self_link")
	f.StatusField(".self_link_with_id")

	// Unimplemented / identity / internal / other fields
	f.Unimplemented_Internal(".parent")
	f.Unimplemented_Identity(".name")
	f.Unimplemented_Internal(".kind")

	// Managed as separate resources or deprecated/unused
	f.Unimplemented_NotYetTriaged(".associations")
	f.Unimplemented_NotYetTriaged(".display_name")
	f.Unimplemented_NotYetTriaged(".packet_mirroring_rules")
	f.Unimplemented_NotYetTriaged(".rules")
	f.Unimplemented_NotYetTriaged(".region")
	f.Unimplemented_NotYetTriaged(".policy_type")

	return f
}
