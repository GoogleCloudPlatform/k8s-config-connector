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
// proto.message: google.cloud.compute.v1.FirewallPolicyRule
// api.group: compute.cnrm.cloud.google.com

package compute

import (
	pb "cloud.google.com/go/compute/apiv1/computepb"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/fuzztesting"
)

func init() {
	fuzztesting.RegisterKRMFuzzer(computeOrganizationSecurityPolicyRuleFuzzer())
}

func computeOrganizationSecurityPolicyRuleFuzzer() fuzztesting.KRMFuzzer {
	f := fuzztesting.NewKRMTypedFuzzer(&pb.FirewallPolicyRule{},
		ComputeOrganizationSecurityPolicyRuleSpec_v1alpha1_FromProto, ComputeOrganizationSecurityPolicyRuleSpec_v1alpha1_ToProto,
		ComputeOrganizationSecurityPolicyRuleObservedState_v1alpha1_FromProto, ComputeOrganizationSecurityPolicyRuleObservedState_v1alpha1_ToProto,
	)

	// Spec fields
	f.SpecField(".action")
	f.SpecField(".description")
	f.SpecField(".direction")
	f.SpecField(".enable_logging")
	f.SpecField(".match")
	f.SpecField(".match.dest_ip_ranges")
	f.SpecField(".match.src_ip_ranges")
	f.SpecField(".match.layer4_configs")
	f.SpecField(".priority") // mapped to ResourceID
	f.SpecField(".target_resources")
	f.SpecField(".target_service_accounts")

	// Unimplemented / Internal fields
	f.Unimplemented_Internal(".kind")
	f.Unimplemented_Internal(".rule_name")
	f.Unimplemented_Internal(".rule_tuple_count")
	f.Unimplemented_Internal(".security_profile_group")
	f.Unimplemented_Internal(".target_secure_tags")
	f.Unimplemented_Internal(".tls_inspect")
	f.Unimplemented_Internal(".disabled")

	// Unimplemented / Unmapped Matcher fields
	f.Unimplemented_Internal(".match.dest_address_groups")
	f.Unimplemented_Internal(".match.dest_fqdns")
	f.Unimplemented_Internal(".match.dest_network_context")
	f.Unimplemented_Internal(".match.dest_network_type")
	f.Unimplemented_Internal(".match.dest_region_codes")
	f.Unimplemented_Internal(".match.dest_threat_intelligences")
	f.Unimplemented_Internal(".match.src_address_groups")
	f.Unimplemented_Internal(".match.src_fqdns")
	f.Unimplemented_Internal(".match.src_network_context")
	f.Unimplemented_Internal(".match.src_network_type")
	f.Unimplemented_Internal(".match.src_networks")
	f.Unimplemented_Internal(".match.src_region_codes")
	f.Unimplemented_Internal(".match.src_secure_tags")
	f.Unimplemented_Internal(".match.src_threat_intelligences")

	// Normalize Match struct to nil if all mapped fields are empty, after zeroing out unmapped fields.
	f.FilterSpec = func(in *pb.FirewallPolicyRule) {
		if in.Match != nil {
			in.Match.DestAddressGroups = nil
			in.Match.DestFqdns = nil
			in.Match.DestNetworkContext = nil
			in.Match.DestNetworkType = nil
			in.Match.DestRegionCodes = nil
			in.Match.DestThreatIntelligences = nil
			in.Match.SrcAddressGroups = nil
			in.Match.SrcFqdns = nil
			in.Match.SrcNetworkContext = nil
			in.Match.SrcNetworkType = nil
			in.Match.SrcNetworks = nil
			in.Match.SrcRegionCodes = nil
			in.Match.SrcSecureTags = nil
			in.Match.SrcThreatIntelligences = nil

			if len(in.Match.DestIpRanges) == 0 && len(in.Match.SrcIpRanges) == 0 && len(in.Match.Layer4Configs) == 0 {
				in.Match = nil
			}
		}
	}

	return f
}
