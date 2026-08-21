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
// proto.message: google.cloud.compute.v1.SecurityPolicy
// api.group: compute.cnrm.cloud.google.com

package compute

import (
	pb "cloud.google.com/go/compute/apiv1/computepb"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/fuzztesting"
)

func init() {
	fuzztesting.RegisterKRMFuzzer(computeSecurityPolicyFuzzer())
}

func computeSecurityPolicyFuzzer() fuzztesting.KRMFuzzer {
	f := fuzztesting.NewKRMTypedFuzzer(&pb.SecurityPolicy{},
		ComputeSecurityPolicySpec_v1beta1_FromProto, ComputeSecurityPolicySpec_v1beta1_ToProto,
		ComputeSecurityPolicyObservedState_v1beta1_FromProto, ComputeSecurityPolicyObservedState_v1beta1_ToProto,
	)

	// Spec fields
	f.SpecField(".adaptive_protection_config")
	f.SpecField(".advanced_options_config")
	f.SpecField(".description")
	f.SpecField(".recaptcha_options_config")
	f.SpecField(".rules")
	f.SpecField(".type")
	f.SpecField(".region")

	// Unimplemented / Identity fields
	f.Unimplemented_Identity(".name")

	// Unimplemented / Not Yet Triaged fields
	f.Unimplemented_NotYetTriaged(".creation_timestamp")
	f.Unimplemented_NotYetTriaged(".ddos_protection_config")
	f.Unimplemented_NotYetTriaged(".ddos_protection_config.ddos_protection")
	f.Unimplemented_NotYetTriaged(".fingerprint")
	f.Unimplemented_NotYetTriaged(".id")
	f.Unimplemented_NotYetTriaged(".kind")
	f.Unimplemented_NotYetTriaged(".label_fingerprint")
	f.Unimplemented_NotYetTriaged(".labels")
	f.Unimplemented_NotYetTriaged(".self_link")
	f.Unimplemented_NotYetTriaged(".user_defined_fields")
	f.Unimplemented_NotYetTriaged(".user_defined_fields[].base")
	f.Unimplemented_NotYetTriaged(".user_defined_fields[].mask")
	f.Unimplemented_NotYetTriaged(".user_defined_fields[].name")
	f.Unimplemented_NotYetTriaged(".user_defined_fields[].offset")
	f.Unimplemented_NotYetTriaged(".user_defined_fields[].size")
	f.Unimplemented_NotYetTriaged(".associations")
	f.Unimplemented_NotYetTriaged(".associations[].display_name")
	f.Unimplemented_NotYetTriaged(".associations[].security_policy_id")
	f.Unimplemented_NotYetTriaged(".associations[].name")
	f.Unimplemented_NotYetTriaged(".associations[].attachment_id")
	f.Unimplemented_NotYetTriaged(".associations[].excluded_folders")
	f.Unimplemented_NotYetTriaged(".associations[].excluded_projects")
	f.Unimplemented_NotYetTriaged(".associations[].short_name")
	f.Unimplemented_NotYetTriaged(".parent")
	f.Unimplemented_NotYetTriaged(".short_name")
	f.Unimplemented_NotYetTriaged(".advanced_options_config.request_body_inspection_size")

	// Rules subfields that are not mapped
	f.Unimplemented_NotYetTriaged(".rules[].kind")
	f.Unimplemented_NotYetTriaged(".rules[].network_match")
	f.Unimplemented_NotYetTriaged(".rules[].match.expr_options")
	f.Unimplemented_NotYetTriaged(".rules[].match.expr.description")
	f.Unimplemented_NotYetTriaged(".rules[].match.expr.location")
	f.Unimplemented_NotYetTriaged(".rules[].match.expr.title")

	// ThresholdConfig subfields that are not mapped
	f.Unimplemented_NotYetTriaged(".adaptive_protection_config.layer7_ddos_defense_config.threshold_configs[].name")
	f.Unimplemented_NotYetTriaged(".adaptive_protection_config.layer7_ddos_defense_config.threshold_configs[].detection_absolute_qps")
	f.Unimplemented_NotYetTriaged(".adaptive_protection_config.layer7_ddos_defense_config.threshold_configs[].detection_load_threshold")
	f.Unimplemented_NotYetTriaged(".adaptive_protection_config.layer7_ddos_defense_config.threshold_configs[].detection_relative_to_baseline_qps")
	f.Unimplemented_NotYetTriaged(".adaptive_protection_config.layer7_ddos_defense_config.threshold_configs[].traffic_granularity_configs")
	f.Unimplemented_NotYetTriaged(".adaptive_protection_config.layer7_ddos_defense_config.threshold_configs[].traffic_granularity_configs[].enable_each_unique_value")
	f.Unimplemented_NotYetTriaged(".adaptive_protection_config.layer7_ddos_defense_config.threshold_configs[].traffic_granularity_configs[].type")
	f.Unimplemented_NotYetTriaged(".adaptive_protection_config.layer7_ddos_defense_config.threshold_configs[].traffic_granularity_configs[].value")

	// FilterSpec to handle lossy mapping of adaptive_protection_config.layer7_ddos_defense_config.threshold_configs
	f.FilterSpec = func(in *pb.SecurityPolicy) {
		if in.AdaptiveProtectionConfig != nil && in.AdaptiveProtectionConfig.Layer7DdosDefenseConfig != nil {
			configs := in.AdaptiveProtectionConfig.Layer7DdosDefenseConfig.ThresholdConfigs
			var autoDeployConfig *pb.SecurityPolicyAdaptiveProtectionConfigLayer7DdosDefenseConfigThresholdConfig
			for _, tc := range configs {
				if tc.AutoDeployLoadThreshold != nil || tc.AutoDeployConfidenceThreshold != nil || tc.AutoDeployImpactedBaselineThreshold != nil || tc.AutoDeployExpirationSec != nil {
					autoDeployConfig = &pb.SecurityPolicyAdaptiveProtectionConfigLayer7DdosDefenseConfigThresholdConfig{
						AutoDeployLoadThreshold:             tc.AutoDeployLoadThreshold,
						AutoDeployConfidenceThreshold:       tc.AutoDeployConfidenceThreshold,
						AutoDeployImpactedBaselineThreshold: tc.AutoDeployImpactedBaselineThreshold,
						AutoDeployExpirationSec:             tc.AutoDeployExpirationSec,
					}
					break
				}
			}
			if autoDeployConfig != nil {
				in.AdaptiveProtectionConfig.Layer7DdosDefenseConfig.ThresholdConfigs = []*pb.SecurityPolicyAdaptiveProtectionConfigLayer7DdosDefenseConfigThresholdConfig{autoDeployConfig}
			} else {
				in.AdaptiveProtectionConfig.Layer7DdosDefenseConfig.ThresholdConfigs = nil
			}
		}
	}

	return f
}
