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

package compute

import (
	pb "cloud.google.com/go/compute/apiv1/computepb"
	krm "github.com/GoogleCloudPlatform/k8s-config-connector/apis/compute/v1beta1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct"
)

func SecurityPolicyRecaptchaOptionsConfig_v1beta1_ToProto(mapCtx *direct.MapContext, in *krm.SecurityPolicyRecaptchaOptionsConfig) *pb.SecurityPolicyRecaptchaOptionsConfig {
	if in == nil {
		return nil
	}
	out := &pb.SecurityPolicyRecaptchaOptionsConfig{}
	if in.RedirectSiteKeyRef != nil {
		out.RedirectSiteKey = &in.RedirectSiteKeyRef.External
	}
	return out
}

func SecurityPolicyRecaptchaOptionsConfig_v1beta1_FromProto(mapCtx *direct.MapContext, in *pb.SecurityPolicyRecaptchaOptionsConfig) *krm.SecurityPolicyRecaptchaOptionsConfig {
	if in == nil {
		return nil
	}
	out := &krm.SecurityPolicyRecaptchaOptionsConfig{}
	if in.GetRedirectSiteKey() != "" {
		out.RedirectSiteKeyRef = &krm.RecaptchaEnterpriseKeyRef{External: in.GetRedirectSiteKey()}
	}
	return out
}

func SecurityPolicyAdaptiveProtectionConfig_v1beta1_ToProto(mapCtx *direct.MapContext, in *krm.SecurityPolicyAdaptiveProtectionConfig) *pb.SecurityPolicyAdaptiveProtectionConfig {
	if in == nil {
		return nil
	}
	out := &pb.SecurityPolicyAdaptiveProtectionConfig{}
	if in.Layer7DdosDefenseConfig != nil {
		out.Layer7DdosDefenseConfig = &pb.SecurityPolicyAdaptiveProtectionConfigLayer7DdosDefenseConfig{
			Enable:         in.Layer7DdosDefenseConfig.Enable,
			RuleVisibility: in.Layer7DdosDefenseConfig.RuleVisibility,
		}
	}

	if in.AutoDeployConfig != nil {
		// Map AutoDeployConfig to a specific ThresholdConfig in the list.
		// We append a new ThresholdConfig with the values.
		// Note: The Name field is required for ThresholdConfig in proto but AutoDeployConfig in Terraform doesn't have it.
		// The API might expect a specific name or allow empty?
		// We will try leaving it empty or using a sentinel if needed.
		thresholdConfig := &pb.SecurityPolicyAdaptiveProtectionConfigLayer7DdosDefenseConfigThresholdConfig{
			AutoDeployLoadThreshold:             in.AutoDeployConfig.LoadThreshold,
			AutoDeployConfidenceThreshold:       in.AutoDeployConfig.ConfidenceThreshold,
			AutoDeployImpactedBaselineThreshold: in.AutoDeployConfig.ImpactedBaselineThreshold,
			AutoDeployExpirationSec:             in.AutoDeployConfig.ExpirationSec,
		}

		if out.Layer7DdosDefenseConfig == nil {
			out.Layer7DdosDefenseConfig = &pb.SecurityPolicyAdaptiveProtectionConfigLayer7DdosDefenseConfig{}
		}
		out.Layer7DdosDefenseConfig.ThresholdConfigs = append(out.Layer7DdosDefenseConfig.ThresholdConfigs, thresholdConfig)
	}

	return out
}

func SecurityPolicyAdaptiveProtectionConfig_v1beta1_FromProto(mapCtx *direct.MapContext, in *pb.SecurityPolicyAdaptiveProtectionConfig) *krm.SecurityPolicyAdaptiveProtectionConfig {
	if in == nil {
		return nil
	}
	out := &krm.SecurityPolicyAdaptiveProtectionConfig{}
	if in.Layer7DdosDefenseConfig != nil {
		out.Layer7DdosDefenseConfig = &krm.SecurityPolicyAdaptiveProtectionConfigLayer7DdosDefenseConfig{
			Enable:         in.Layer7DdosDefenseConfig.Enable,
			RuleVisibility: in.Layer7DdosDefenseConfig.RuleVisibility,
		}

		// Try to find the ThresholdConfig that corresponds to AutoDeployConfig.
		// We assume it's the one with AutoDeploy* fields set.
		for _, tc := range in.Layer7DdosDefenseConfig.ThresholdConfigs {
			if tc.AutoDeployLoadThreshold != nil || tc.AutoDeployConfidenceThreshold != nil || tc.AutoDeployImpactedBaselineThreshold != nil || tc.AutoDeployExpirationSec != nil {
				out.AutoDeployConfig = &krm.SecurityPolicyAdaptiveProtectionConfigAutoDeployConfig{
					LoadThreshold:             tc.AutoDeployLoadThreshold,
					ConfidenceThreshold:       tc.AutoDeployConfidenceThreshold,
					ImpactedBaselineThreshold: tc.AutoDeployImpactedBaselineThreshold,
					ExpirationSec:             tc.AutoDeployExpirationSec,
				}
				// We assume only one such config exists or we take the first one.
				break
			}
		}
	}
	return out
}

func SecurityPolicyRulePreconfiguredWafConfig_v1beta1_ToProto(mapCtx *direct.MapContext, in *krm.SecurityPolicyRulePreconfiguredWafConfig) *pb.SecurityPolicyRulePreconfiguredWafConfig {
	if in == nil {
		return nil
	}
	out := &pb.SecurityPolicyRulePreconfiguredWafConfig{}
	out.Exclusions = direct.Slice_ToProto(mapCtx, in.Exclusion, SecurityPolicyRulePreconfiguredWafConfigExclusion_v1beta1_ToProto)
	return out
}

func SecurityPolicyRulePreconfiguredWafConfig_v1beta1_FromProto(mapCtx *direct.MapContext, in *pb.SecurityPolicyRulePreconfiguredWafConfig) *krm.SecurityPolicyRulePreconfiguredWafConfig {
	if in == nil {
		return nil
	}
	out := &krm.SecurityPolicyRulePreconfiguredWafConfig{}
	out.Exclusion = direct.Slice_FromProto(mapCtx, in.Exclusions, SecurityPolicyRulePreconfiguredWafConfigExclusion_v1beta1_FromProto)
	return out
}

func SecurityPolicyRulePreconfiguredWafConfigExclusion_v1beta1_ToProto(mapCtx *direct.MapContext, in *krm.SecurityPolicyRulePreconfiguredWafConfigExclusion) *pb.SecurityPolicyRulePreconfiguredWafConfigExclusion {
	if in == nil {
		return nil
	}
	out := &pb.SecurityPolicyRulePreconfiguredWafConfigExclusion{}
	out.RequestCookiesToExclude = direct.Slice_ToProto(mapCtx, in.RequestCookie, SecurityPolicyRulePreconfiguredWafConfigExclusionFieldParams_v1beta1_ToProto)
	out.RequestHeadersToExclude = direct.Slice_ToProto(mapCtx, in.RequestHeader, SecurityPolicyRulePreconfiguredWafConfigExclusionFieldParams_v1beta1_ToProto)
	out.RequestQueryParamsToExclude = direct.Slice_ToProto(mapCtx, in.RequestQueryParam, SecurityPolicyRulePreconfiguredWafConfigExclusionFieldParams_v1beta1_ToProto)
	out.RequestUrisToExclude = direct.Slice_ToProto(mapCtx, in.RequestUri, SecurityPolicyRulePreconfiguredWafConfigExclusionFieldParams_v1beta1_ToProto)
	out.TargetRuleIds = in.TargetRuleIds
	out.TargetRuleSet = in.TargetRuleSet
	return out
}

func SecurityPolicyRulePreconfiguredWafConfigExclusion_v1beta1_FromProto(mapCtx *direct.MapContext, in *pb.SecurityPolicyRulePreconfiguredWafConfigExclusion) *krm.SecurityPolicyRulePreconfiguredWafConfigExclusion {
	if in == nil {
		return nil
	}
	out := &krm.SecurityPolicyRulePreconfiguredWafConfigExclusion{}
	out.RequestCookie = direct.Slice_FromProto(mapCtx, in.RequestCookiesToExclude, SecurityPolicyRulePreconfiguredWafConfigExclusionFieldParams_v1beta1_FromProto)
	out.RequestHeader = direct.Slice_FromProto(mapCtx, in.RequestHeadersToExclude, SecurityPolicyRulePreconfiguredWafConfigExclusionFieldParams_v1beta1_FromProto)
	out.RequestQueryParam = direct.Slice_FromProto(mapCtx, in.RequestQueryParamsToExclude, SecurityPolicyRulePreconfiguredWafConfigExclusionFieldParams_v1beta1_FromProto)
	out.RequestUri = direct.Slice_FromProto(mapCtx, in.RequestUrisToExclude, SecurityPolicyRulePreconfiguredWafConfigExclusionFieldParams_v1beta1_FromProto)
	out.TargetRuleIds = in.TargetRuleIds
	out.TargetRuleSet = in.TargetRuleSet
	return out
}

func SecurityPolicyRulePreconfiguredWafConfigExclusionFieldParams_v1beta1_ToProto(mapCtx *direct.MapContext, in *krm.SecurityPolicyRulePreconfiguredWafConfigExclusionFieldParams) *pb.SecurityPolicyRulePreconfiguredWafConfigExclusionFieldParams {
	if in == nil {
		return nil
	}
	out := &pb.SecurityPolicyRulePreconfiguredWafConfigExclusionFieldParams{}
	out.Op = in.Operator
	out.Val = in.Value
	return out
}

func SecurityPolicyRulePreconfiguredWafConfigExclusionFieldParams_v1beta1_FromProto(mapCtx *direct.MapContext, in *pb.SecurityPolicyRulePreconfiguredWafConfigExclusionFieldParams) *krm.SecurityPolicyRulePreconfiguredWafConfigExclusionFieldParams {
	if in == nil {
		return nil
	}
	out := &krm.SecurityPolicyRulePreconfiguredWafConfigExclusionFieldParams{}
	out.Operator = in.Op
	out.Value = in.Val
	return out
}

func SecurityPolicyRuleMatcherExpr_v1beta1_ToProto(mapCtx *direct.MapContext, in *krm.SecurityPolicyRuleMatcherExpr) *pb.Expr {
	if in == nil {
		return nil
	}
	out := &pb.Expr{}
	out.Expression = in.Expression
	return out
}

func SecurityPolicyRuleMatcherExpr_v1beta1_FromProto(mapCtx *direct.MapContext, in *pb.Expr) *krm.SecurityPolicyRuleMatcherExpr {
	if in == nil {
		return nil
	}
	out := &krm.SecurityPolicyRuleMatcherExpr{}
	out.Expression = in.Expression
	return out
}
