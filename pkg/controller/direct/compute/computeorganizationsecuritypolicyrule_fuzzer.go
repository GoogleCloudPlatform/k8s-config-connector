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
// proto.message: google.cloud.compute.v1.SecurityPolicyRule
// api.group: compute.cnrm.cloud.google.com

package compute

import (
	pb "cloud.google.com/go/compute/apiv1/computepb"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/fuzztesting"
)

func init() {
	fuzztesting.RegisterKRMSpecFuzzer(computeOrganizationSecurityPolicyRuleFuzzer())
}

func computeOrganizationSecurityPolicyRuleFuzzer() fuzztesting.KRMFuzzer {
	f := fuzztesting.NewKRMTypedSpecFuzzer(&pb.SecurityPolicyRule{},
		ComputeOrganizationSecurityPolicyRuleSpec_v1alpha1_FromProto,
		ComputeOrganizationSecurityPolicyRuleSpec_v1alpha1_ToProto,
	)

	// Field Comparison & Mapping Documentation:
	//
	// KRM Spec Fields:
	// - action                -> f.SpecField(".action")
	// - description           -> f.SpecField(".description")
	// - direction             -> (not present in pb.SecurityPolicyRule; handled by GCP API directly)
	// - enableLogging         -> (not present in pb.SecurityPolicyRule; handled by GCP API directly)
	// - match                 -> f.SpecField(".match")
	//   - config              -> f.SpecField(".match.config")
	//     - destIpRanges      -> (not present in pb.SecurityPolicyRuleMatcherConfig)
	//     - layer4Config      -> (not present in pb.SecurityPolicyRuleMatcherConfig)
	//     - srcIpRanges       -> f.SpecField(".match.config.src_ip_ranges")
	//   - description         -> (not present in pb.SecurityPolicyRuleMatcher)
	//   - versionedExpr       -> f.SpecField(".match.versioned_expr")
	// - policyId              -> (parent URL identifier, not part of rule proto)
	// - preview               -> f.SpecField(".preview")
	// - resourceID            -> f.SpecField(".priority") (maps to rule priority)
	// - targetResources       -> (not present in pb.SecurityPolicyRule; handled by GCP API directly)
	// - targetServiceAccounts -> (not present in pb.SecurityPolicyRule; handled by GCP API directly)
	//
	// Proto Fields Unimplemented / Not Yet Triaged:
	// - .header_action
	// - .kind
	// - .match.expr
	// - .match.expr_options
	// - .network_match
	// - .preconfigured_waf_config
	// - .rate_limit_options
	// - .redirect_options

	// Spec fields
	f.SpecField(".action")
	f.SpecField(".description")
	f.SpecField(".match")
	f.SpecField(".match.config")
	f.SpecField(".match.config.src_ip_ranges")
	f.SpecField(".match.versioned_expr")
	f.SpecField(".preview")
	f.SpecField(".priority")

	// Unimplemented / Not Yet Triaged fields
	f.Unimplemented_NotYetTriaged(".header_action")
	f.Unimplemented_NotYetTriaged(".kind")
	f.Unimplemented_NotYetTriaged(".match.expr")
	f.Unimplemented_NotYetTriaged(".match.expr_options")
	f.Unimplemented_NotYetTriaged(".network_match")
	f.Unimplemented_NotYetTriaged(".preconfigured_waf_config")
	f.Unimplemented_NotYetTriaged(".rate_limit_options")
	f.Unimplemented_NotYetTriaged(".redirect_options")

	f.FilterSpec = func(in *pb.SecurityPolicyRule) {
		if in.Action != nil && *in.Action == "" {
			in.Action = nil
		}
	}

	return f
}
