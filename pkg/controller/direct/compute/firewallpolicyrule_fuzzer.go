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
		ComputeFirewallPolicyRuleSpec_FromProto, ComputeFirewallPolicyRuleSpec_ToProto,
		ComputeFirewallPolicyRuleStatus_FromProto, ComputeFirewallPolicyRuleStatus_ToProto,
	)

	f.SpecFields.Insert(".action")
	f.SpecFields.Insert(".description")
	f.SpecFields.Insert(".direction")
	f.SpecFields.Insert(".disabled")
	f.SpecFields.Insert(".enable_logging")
	f.SpecFields.Insert(".match")
	f.SpecFields.Insert(".priority")
	f.SpecFields.Insert(".target_resources")
	f.SpecFields.Insert(".target_service_accounts")

	f.StatusFields.Insert(".kind")
	f.StatusFields.Insert(".rule_tuple_count")

	f.UnimplementedFields.Insert(".rule_name")
	f.UnimplementedFields.Insert(".security_profile_group")
	f.UnimplementedFields.Insert(".target_secure_tags")
	f.UnimplementedFields.Insert(".tls_inspect")
	f.UnimplementedFields.Insert(".match.src_secure_tags")
	f.UnimplementedFields.Insert(".match.src_networks")
	f.UnimplementedFields.Insert(".match.src_network_type")
	f.UnimplementedFields.Insert(".match.dest_network_type")

	return f
}
