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
	fuzztesting.RegisterKRMFuzzer(computeNetworkFirewallPolicyFuzzer())
}

func computeNetworkFirewallPolicyFuzzer() fuzztesting.KRMFuzzer {
	f := fuzztesting.NewKRMTypedFuzzer(&pb.FirewallPolicy{},
		ComputeNetworkFirewallPolicySpec_v1beta1_FromProto, ComputeNetworkFirewallPolicySpec_v1beta1_ToProto,
		ComputeNetworkFirewallPolicyStatus_v1beta1_FromProto, ComputeNetworkFirewallPolicyStatus_v1beta1_ToProto,
	)

	// Detailed Comparison of KRM Fields with Proto Fields:
	// +-------------------------+---------------------------+----------------------------------------------+
	// | KRM Spec Field          | Proto Field               | Fuzzer/Mapping Status                        |
	// +-------------------------+---------------------------+----------------------------------------------+
	// | description             | description               | SpecField (".description")                   |
	// | projectRef              | N/A                       | Part of resource URI/identity                |
	// | resourceID              | name                      | Identity/name field (Unimplemented_Identity) |
	// +-------------------------+---------------------------+----------------------------------------------+
	//
	// +-------------------------+---------------------------+----------------------------------------------+
	// | KRM Status Field        | Proto Field               | Fuzzer/Mapping Status                        |
	// +-------------------------+---------------------------+----------------------------------------------+
	// | conditions              | N/A                       | KCC status conditions                        |
	// | creationTimestamp       | creation_timestamp        | StatusField (".creation_timestamp")          |
	// | fingerprint             | fingerprint               | StatusField (".fingerprint")                 |
	// | networkFirewallPolicyId | id                        | StatusField (".id")                          |
	// | observedGeneration      | N/A                       | KCC observed generation                      |
	// | ruleTupleCount          | rule_tuple_count          | StatusField (".rule_tuple_count")            |
	// | selfLink                | self_link                 | StatusField (".self_link")                   |
	// | selfLinkWithId          | self_link_with_id         | StatusField (".self_link_with_id")          |
	// +-------------------------+---------------------------+----------------------------------------------+

	// Spec fields
	f.SpecField(".description")

	// Status fields
	f.StatusField(".creation_timestamp")
	f.StatusField(".fingerprint")
	f.StatusField(".id")
	f.StatusField(".rule_tuple_count")
	f.StatusField(".self_link")
	f.StatusField(".self_link_with_id")

	// Identity / Special fields
	f.Unimplemented_Identity(".name")

	// Internal fields
	f.Unimplemented_Internal(".kind")
	f.Unimplemented_Internal(".parent")

	// Unimplemented / Not Yet Triaged fields
	f.Unimplemented_NotYetTriaged(".associations")
	f.Unimplemented_NotYetTriaged(".display_name")
	f.Unimplemented_NotYetTriaged(".packet_mirroring_rules")
	f.Unimplemented_NotYetTriaged(".policy_type")
	f.Unimplemented_NotYetTriaged(".region")
	f.Unimplemented_NotYetTriaged(".rules")
	f.Unimplemented_NotYetTriaged(".short_name")

	return f
}
