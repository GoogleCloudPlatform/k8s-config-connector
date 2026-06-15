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
// proto.message: google.cloud.compute.v1.FirewallPolicyRule
// api.group: compute.cnrm.cloud.google.com

package compute

import (
	pb "cloud.google.com/go/compute/apiv1/computepb"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/fuzztesting"
)

func init() {
	fuzztesting.RegisterKRMFuzzer(computeFirewallPolicyRuleFuzzer())
}

func computeFirewallPolicyRuleFuzzer() fuzztesting.KRMFuzzer {
	f := fuzztesting.NewKRMTypedFuzzer(&pb.FirewallPolicyRule{},
		ComputeFirewallPolicyRuleSpec_v1beta1_FromProto, ComputeFirewallPolicyRuleSpec_v1beta1_ToProto,
		ComputeFirewallPolicyRuleStatus_v1beta1_FromProto, ComputeFirewallPolicyRuleStatus_v1beta1_ToProto,
	)

	f.SpecField(".action")
	f.SpecField(".description")
	f.SpecField(".direction")
	f.SpecField(".disabled")
	f.SpecField(".enable_logging")
	f.SpecField(".match")
	f.SpecField(".priority")
	f.SpecField(".target_resources")
	f.SpecField(".target_service_accounts")

	f.StatusField(".kind")
	f.StatusField(".rule_tuple_count")

	f.Unimplemented_NotYetTriaged(".rule_name")
	f.Unimplemented_NotYetTriaged(".security_profile_group")
	f.Unimplemented_NotYetTriaged(".target_secure_tags")
	f.Unimplemented_NotYetTriaged(".tls_inspect")
	f.Unimplemented_NotYetTriaged(".match.src_secure_tags")
	f.Unimplemented_NotYetTriaged(".match.src_networks")
	f.Unimplemented_NotYetTriaged(".match.src_network_type")
	f.Unimplemented_NotYetTriaged(".match.dest_network_type")
	f.Unimplemented_NotYetTriaged(".match.dest_network_context")
	f.Unimplemented_NotYetTriaged(".match.src_network_context")

	return f
}
