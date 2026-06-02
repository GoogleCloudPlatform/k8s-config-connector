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

package computesecuritypolicy

import (
	pb "cloud.google.com/go/compute/apiv1/computepb"
	krm "github.com/GoogleCloudPlatform/k8s-config-connector/apis/compute/v1beta1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct"
)

// Mapping functions for SecurityPolicyRecaptchaOptionsConfig
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
	if in.RedirectSiteKey != nil {
		out.RedirectSiteKeyRef = &krm.RecaptchaEnterpriseKeyRef{External: *in.RedirectSiteKey}
	}
	return out
}

// Mapping functions for SecurityPolicyAdaptiveProtectionConfig
func SecurityPolicyAdaptiveProtectionConfig_v1beta1_ToProto(mapCtx *direct.MapContext, in *krm.SecurityPolicyAdaptiveProtectionConfig) *pb.SecurityPolicyAdaptiveProtectionConfig {
	if in == nil {
		return nil
	}
	out := &pb.SecurityPolicyAdaptiveProtectionConfig{}
	out.Layer7DdosDefenseConfig = SecurityPolicyAdaptiveProtectionConfigLayer7DdosDefenseConfig_v1beta1_ToProto(mapCtx, in.Layer7DdosDefenseConfig)
	return out
}

func SecurityPolicyAdaptiveProtectionConfig_v1beta1_FromProto(mapCtx *direct.MapContext, in *pb.SecurityPolicyAdaptiveProtectionConfig) *krm.SecurityPolicyAdaptiveProtectionConfig {
	if in == nil {
		return nil
	}
	out := &krm.SecurityPolicyAdaptiveProtectionConfig{}
	out.Layer7DdosDefenseConfig = SecurityPolicyAdaptiveProtectionConfigLayer7DdosDefenseConfig_v1beta1_FromProto(mapCtx, in.Layer7DdosDefenseConfig)
	return out
}

// Mapping functions for Layer7DdosDefenseConfig
func SecurityPolicyAdaptiveProtectionConfigLayer7DdosDefenseConfig_v1beta1_ToProto(mapCtx *direct.MapContext, in *krm.SecurityPolicyAdaptiveProtectionConfigLayer7DdosDefenseConfig) *pb.SecurityPolicyAdaptiveProtectionConfigLayer7DdosDefenseConfig {
	if in == nil {
		return nil
	}
	out := &pb.SecurityPolicyAdaptiveProtectionConfigLayer7DdosDefenseConfig{}
	out.Enable = in.Enable
	out.RuleVisibility = in.RuleVisibility
	return out
}

func SecurityPolicyAdaptiveProtectionConfigLayer7DdosDefenseConfig_v1beta1_FromProto(mapCtx *direct.MapContext, in *pb.SecurityPolicyAdaptiveProtectionConfigLayer7DdosDefenseConfig) *krm.SecurityPolicyAdaptiveProtectionConfigLayer7DdosDefenseConfig {
	if in == nil {
		return nil
	}
	out := &krm.SecurityPolicyAdaptiveProtectionConfigLayer7DdosDefenseConfig{}
	out.Enable = in.Enable
	out.RuleVisibility = in.RuleVisibility
	return out
}

// Mapping functions for AdvancedOptionsConfig
func SecurityPolicyAdvancedOptionsConfig_v1beta1_ToProto(mapCtx *direct.MapContext, in *krm.SecurityPolicyAdvancedOptionsConfig) *pb.SecurityPolicyAdvancedOptionsConfig {
	if in == nil {
		return nil
	}
	out := &pb.SecurityPolicyAdvancedOptionsConfig{}
	out.JsonCustomConfig = SecurityPolicyAdvancedOptionsConfigJsonCustomConfig_v1beta1_ToProto(mapCtx, in.JsonCustomConfig)
	out.JsonParsing = in.JsonParsing
	out.LogLevel = in.LogLevel
	out.UserIpRequestHeaders = in.UserIPRequestHeaders
	return out
}

func SecurityPolicyAdvancedOptionsConfig_v1beta1_FromProto(mapCtx *direct.MapContext, in *pb.SecurityPolicyAdvancedOptionsConfig) *krm.SecurityPolicyAdvancedOptionsConfig {
	if in == nil {
		return nil
	}
	out := &krm.SecurityPolicyAdvancedOptionsConfig{}
	out.JsonCustomConfig = SecurityPolicyAdvancedOptionsConfigJsonCustomConfig_v1beta1_FromProto(mapCtx, in.JsonCustomConfig)
	out.JsonParsing = in.JsonParsing
	out.LogLevel = in.LogLevel
	out.UserIPRequestHeaders = in.UserIpRequestHeaders
	return out
}

// Mapping functions for JsonCustomConfig
func SecurityPolicyAdvancedOptionsConfigJsonCustomConfig_v1beta1_ToProto(mapCtx *direct.MapContext, in *krm.SecurityPolicyAdvancedOptionsConfigJsonCustomConfig) *pb.SecurityPolicyAdvancedOptionsConfigJsonCustomConfig {
	if in == nil {
		return nil
	}
	out := &pb.SecurityPolicyAdvancedOptionsConfigJsonCustomConfig{}
	out.ContentTypes = in.ContentTypes
	return out
}

func SecurityPolicyAdvancedOptionsConfigJsonCustomConfig_v1beta1_FromProto(mapCtx *direct.MapContext, in *pb.SecurityPolicyAdvancedOptionsConfigJsonCustomConfig) *krm.SecurityPolicyAdvancedOptionsConfigJsonCustomConfig {
	if in == nil {
		return nil
	}
	out := &krm.SecurityPolicyAdvancedOptionsConfigJsonCustomConfig{}
	out.ContentTypes = in.ContentTypes
	return out
}

// Mapping functions for SecurityPolicyRulePreconfiguredWafConfig
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

// Mapping functions for Exclusion
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

// Mapping functions for ExclusionFieldParams
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

// Mapping functions for SecurityPolicyRuleMatcher
func SecurityPolicyRuleMatcher_v1beta1_ToProto(mapCtx *direct.MapContext, in *krm.SecurityPolicyRuleMatcher) *pb.SecurityPolicyRuleMatcher {
	if in == nil {
		return nil
	}
	out := &pb.SecurityPolicyRuleMatcher{}
	out.Config = SecurityPolicyRuleMatcherConfig_v1beta1_ToProto(mapCtx, in.Config)
	out.Expr = SecurityPolicyRuleMatcherExpr_v1beta1_ToProto(mapCtx, in.Expr)
	out.VersionedExpr = in.VersionedExpr
	return out
}

func SecurityPolicyRuleMatcher_v1beta1_FromProto(mapCtx *direct.MapContext, in *pb.SecurityPolicyRuleMatcher) *krm.SecurityPolicyRuleMatcher {
	if in == nil {
		return nil
	}
	out := &krm.SecurityPolicyRuleMatcher{}
	out.Config = SecurityPolicyRuleMatcherConfig_v1beta1_FromProto(mapCtx, in.Config)
	out.Expr = SecurityPolicyRuleMatcherExpr_v1beta1_FromProto(mapCtx, in.Expr)
	out.VersionedExpr = in.VersionedExpr
	return out
}

// Mapping functions for SecurityPolicyRuleMatcherExpr
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

// Mapping functions for SecurityPolicyRuleMatcherConfig
func SecurityPolicyRuleMatcherConfig_v1beta1_ToProto(mapCtx *direct.MapContext, in *krm.SecurityPolicyRuleMatcherConfig) *pb.SecurityPolicyRuleMatcherConfig {
	if in == nil {
		return nil
	}
	out := &pb.SecurityPolicyRuleMatcherConfig{}
	out.SrcIpRanges = in.SrcIPRanges
	return out
}

func SecurityPolicyRuleMatcherConfig_v1beta1_FromProto(mapCtx *direct.MapContext, in *pb.SecurityPolicyRuleMatcherConfig) *krm.SecurityPolicyRuleMatcherConfig {
	if in == nil {
		return nil
	}
	out := &krm.SecurityPolicyRuleMatcherConfig{}
	out.SrcIPRanges = in.SrcIpRanges
	return out
}

// Mapping functions for SecurityPolicyRuleRedirectOptions
func SecurityPolicyRuleRedirectOptions_v1beta1_ToProto(mapCtx *direct.MapContext, in *krm.SecurityPolicyRuleRedirectOptions) *pb.SecurityPolicyRuleRedirectOptions {
	if in == nil {
		return nil
	}
	out := &pb.SecurityPolicyRuleRedirectOptions{}
	out.Target = in.Target
	out.Type = in.Type
	return out
}

func SecurityPolicyRuleRedirectOptions_v1beta1_FromProto(mapCtx *direct.MapContext, in *pb.SecurityPolicyRuleRedirectOptions) *krm.SecurityPolicyRuleRedirectOptions {
	if in == nil {
		return nil
	}
	out := &krm.SecurityPolicyRuleRedirectOptions{}
	out.Target = in.Target
	out.Type = in.Type
	return out
}

// Mapping functions for SecurityPolicyRuleRateLimitOptionsThreshold
func SecurityPolicyRuleRateLimitOptionsThreshold_v1beta1_ToProto(mapCtx *direct.MapContext, in *krm.SecurityPolicyRuleRateLimitOptionsThreshold) *pb.SecurityPolicyRuleRateLimitOptionsThreshold {
	if in == nil {
		return nil
	}
	out := &pb.SecurityPolicyRuleRateLimitOptionsThreshold{}
	out.Count = in.Count
	out.IntervalSec = in.IntervalSec
	return out
}

func SecurityPolicyRuleRateLimitOptionsThreshold_v1beta1_FromProto(mapCtx *direct.MapContext, in *pb.SecurityPolicyRuleRateLimitOptionsThreshold) *krm.SecurityPolicyRuleRateLimitOptionsThreshold {
	if in == nil {
		return nil
	}
	out := &krm.SecurityPolicyRuleRateLimitOptionsThreshold{}
	out.Count = in.Count
	out.IntervalSec = in.IntervalSec
	return out
}

// Mapping functions for SecurityPolicyRuleRateLimitOptionsEnforceOnKeyConfig
func SecurityPolicyRuleRateLimitOptionsEnforceOnKeyConfig_v1beta1_ToProto(mapCtx *direct.MapContext, in *krm.SecurityPolicyRuleRateLimitOptionsEnforceOnKeyConfig) *pb.SecurityPolicyRuleRateLimitOptionsEnforceOnKeyConfig {
	if in == nil {
		return nil
	}
	out := &pb.SecurityPolicyRuleRateLimitOptionsEnforceOnKeyConfig{}
	out.EnforceOnKeyName = in.EnforceOnKeyName
	out.EnforceOnKeyType = in.EnforceOnKeyType
	return out
}

func SecurityPolicyRuleRateLimitOptionsEnforceOnKeyConfig_v1beta1_FromProto(mapCtx *direct.MapContext, in *pb.SecurityPolicyRuleRateLimitOptionsEnforceOnKeyConfig) *krm.SecurityPolicyRuleRateLimitOptionsEnforceOnKeyConfig {
	if in == nil {
		return nil
	}
	out := &krm.SecurityPolicyRuleRateLimitOptionsEnforceOnKeyConfig{}
	out.EnforceOnKeyName = in.EnforceOnKeyName
	out.EnforceOnKeyType = in.EnforceOnKeyType
	return out
}

// Mapping functions for SecurityPolicyRuleRateLimitOptions
func SecurityPolicyRuleRateLimitOptions_v1beta1_ToProto(mapCtx *direct.MapContext, in *krm.SecurityPolicyRuleRateLimitOptions) *pb.SecurityPolicyRuleRateLimitOptions {
	if in == nil {
		return nil
	}
	out := &pb.SecurityPolicyRuleRateLimitOptions{}
	out.BanDurationSec = in.BanDurationSec
	out.BanThreshold = SecurityPolicyRuleRateLimitOptionsThreshold_v1beta1_ToProto(mapCtx, in.BanThreshold)
	out.ConformAction = in.ConformAction
	out.EnforceOnKey = in.EnforceOnKey
	out.EnforceOnKeyConfigs = direct.Slice_ToProto(mapCtx, in.EnforceOnKeyConfigs, SecurityPolicyRuleRateLimitOptionsEnforceOnKeyConfig_v1beta1_ToProto)
	out.EnforceOnKeyName = in.EnforceOnKeyName
	out.ExceedAction = in.ExceedAction
	out.ExceedRedirectOptions = SecurityPolicyRuleRedirectOptions_v1beta1_ToProto(mapCtx, in.ExceedRedirectOptions)
	out.RateLimitThreshold = SecurityPolicyRuleRateLimitOptionsThreshold_v1beta1_ToProto(mapCtx, in.RateLimitThreshold)
	return out
}

func SecurityPolicyRuleRateLimitOptions_v1beta1_FromProto(mapCtx *direct.MapContext, in *pb.SecurityPolicyRuleRateLimitOptions) *krm.SecurityPolicyRuleRateLimitOptions {
	if in == nil {
		return nil
	}
	out := &krm.SecurityPolicyRuleRateLimitOptions{}
	out.BanDurationSec = in.BanDurationSec
	out.BanThreshold = SecurityPolicyRuleRateLimitOptionsThreshold_v1beta1_FromProto(mapCtx, in.BanThreshold)
	out.ConformAction = in.ConformAction
	out.EnforceOnKey = in.EnforceOnKey
	out.EnforceOnKeyConfigs = direct.Slice_FromProto(mapCtx, in.EnforceOnKeyConfigs, SecurityPolicyRuleRateLimitOptionsEnforceOnKeyConfig_v1beta1_FromProto)
	out.EnforceOnKeyName = in.EnforceOnKeyName
	out.ExceedAction = in.ExceedAction
	out.ExceedRedirectOptions = SecurityPolicyRuleRedirectOptions_v1beta1_FromProto(mapCtx, in.ExceedRedirectOptions)
	out.RateLimitThreshold = SecurityPolicyRuleRateLimitOptionsThreshold_v1beta1_FromProto(mapCtx, in.RateLimitThreshold)
	return out
}

// Mapping functions for SecurityPolicyRuleHTTPHeaderActionHTTPHeaderOption
func SecurityPolicyRuleHTTPHeaderActionHTTPHeaderOption_v1beta1_ToProto(mapCtx *direct.MapContext, in *krm.SecurityPolicyRuleHTTPHeaderActionHTTPHeaderOption) *pb.SecurityPolicyRuleHttpHeaderActionHttpHeaderOption {
	if in == nil {
		return nil
	}
	out := &pb.SecurityPolicyRuleHttpHeaderActionHttpHeaderOption{}
	out.HeaderName = in.HeaderName
	out.HeaderValue = in.HeaderValue
	return out
}

func SecurityPolicyRuleHTTPHeaderActionHTTPHeaderOption_v1beta1_FromProto(mapCtx *direct.MapContext, in *pb.SecurityPolicyRuleHttpHeaderActionHttpHeaderOption) *krm.SecurityPolicyRuleHTTPHeaderActionHTTPHeaderOption {
	if in == nil {
		return nil
	}
	out := &krm.SecurityPolicyRuleHTTPHeaderActionHTTPHeaderOption{}
	out.HeaderName = in.HeaderName
	out.HeaderValue = in.HeaderValue
	return out
}

// Mapping functions for SecurityPolicyRuleHTTPHeaderAction
func SecurityPolicyRuleHTTPHeaderAction_v1beta1_ToProto(mapCtx *direct.MapContext, in *krm.SecurityPolicyRuleHTTPHeaderAction) *pb.SecurityPolicyRuleHttpHeaderAction {
	if in == nil {
		return nil
	}
	out := &pb.SecurityPolicyRuleHttpHeaderAction{}
	out.RequestHeadersToAdds = direct.Slice_ToProto(mapCtx, in.RequestHeadersToAdds, SecurityPolicyRuleHTTPHeaderActionHTTPHeaderOption_v1beta1_ToProto)
	return out
}

func SecurityPolicyRuleHTTPHeaderAction_v1beta1_FromProto(mapCtx *direct.MapContext, in *pb.SecurityPolicyRuleHttpHeaderAction) *krm.SecurityPolicyRuleHTTPHeaderAction {
	if in == nil {
		return nil
	}
	out := &krm.SecurityPolicyRuleHTTPHeaderAction{}
	out.RequestHeadersToAdds = direct.Slice_FromProto(mapCtx, in.RequestHeadersToAdds, SecurityPolicyRuleHTTPHeaderActionHTTPHeaderOption_v1beta1_FromProto)
	return out
}

// Mapping functions for SecurityPolicyRule
func SecurityPolicyRule_v1beta1_ToProto(mapCtx *direct.MapContext, in *krm.SecurityPolicyRule) *pb.SecurityPolicyRule {
	if in == nil {
		return nil
	}
	out := &pb.SecurityPolicyRule{}
	out.Action = in.Action
	out.Description = in.Description
	out.HeaderAction = SecurityPolicyRuleHTTPHeaderAction_v1beta1_ToProto(mapCtx, in.HeaderAction)
	out.Match = SecurityPolicyRuleMatcher_v1beta1_ToProto(mapCtx, in.Match)
	out.PreconfiguredWafConfig = SecurityPolicyRulePreconfiguredWafConfig_v1beta1_ToProto(mapCtx, in.PreconfiguredWafConfig)
	out.Preview = in.Preview
	out.Priority = in.Priority
	out.RateLimitOptions = SecurityPolicyRuleRateLimitOptions_v1beta1_ToProto(mapCtx, in.RateLimitOptions)
	out.RedirectOptions = SecurityPolicyRuleRedirectOptions_v1beta1_ToProto(mapCtx, in.RedirectOptions)
	return out
}

func SecurityPolicyRule_v1beta1_FromProto(mapCtx *direct.MapContext, in *pb.SecurityPolicyRule) *krm.SecurityPolicyRule {
	if in == nil {
		return nil
	}
	out := &krm.SecurityPolicyRule{}
	out.Action = in.Action
	out.Description = in.Description
	out.HeaderAction = SecurityPolicyRuleHTTPHeaderAction_v1beta1_FromProto(mapCtx, in.HeaderAction)
	out.Match = SecurityPolicyRuleMatcher_v1beta1_FromProto(mapCtx, in.Match)
	out.PreconfiguredWafConfig = SecurityPolicyRulePreconfiguredWafConfig_v1beta1_FromProto(mapCtx, in.PreconfiguredWafConfig)
	out.Preview = in.Preview
	out.Priority = in.Priority
	out.RateLimitOptions = SecurityPolicyRuleRateLimitOptions_v1beta1_FromProto(mapCtx, in.RateLimitOptions)
	out.RedirectOptions = SecurityPolicyRuleRedirectOptions_v1beta1_FromProto(mapCtx, in.RedirectOptions)
	return out
}

// Mapping functions for ComputeSecurityPolicySpec
func ComputeSecurityPolicySpec_v1beta1_ToProto(mapCtx *direct.MapContext, in *krm.ComputeSecurityPolicySpec) *pb.SecurityPolicy {
	if in == nil {
		return nil
	}
	out := &pb.SecurityPolicy{}
	out.AdaptiveProtectionConfig = SecurityPolicyAdaptiveProtectionConfig_v1beta1_ToProto(mapCtx, in.AdaptiveProtectionConfig)
	out.AdvancedOptionsConfig = SecurityPolicyAdvancedOptionsConfig_v1beta1_ToProto(mapCtx, in.AdvancedOptionsConfig)
	out.Description = in.Description
	out.RecaptchaOptionsConfig = SecurityPolicyRecaptchaOptionsConfig_v1beta1_ToProto(mapCtx, in.RecaptchaOptionsConfig)
	out.Rules = direct.Slice_ToProto(mapCtx, in.Rules, SecurityPolicyRule_v1beta1_ToProto)
	out.Type = in.Type
	return out
}

func ComputeSecurityPolicySpec_v1beta1_FromProto(mapCtx *direct.MapContext, in *pb.SecurityPolicy) *krm.ComputeSecurityPolicySpec {
	if in == nil {
		return nil
	}
	out := &krm.ComputeSecurityPolicySpec{}
	out.AdaptiveProtectionConfig = SecurityPolicyAdaptiveProtectionConfig_v1beta1_FromProto(mapCtx, in.AdaptiveProtectionConfig)
	out.AdvancedOptionsConfig = SecurityPolicyAdvancedOptionsConfig_v1beta1_FromProto(mapCtx, in.AdvancedOptionsConfig)
	out.Description = in.Description
	out.RecaptchaOptionsConfig = SecurityPolicyRecaptchaOptionsConfig_v1beta1_FromProto(mapCtx, in.RecaptchaOptionsConfig)
	out.Rules = direct.Slice_FromProto(mapCtx, in.Rules, SecurityPolicyRule_v1beta1_FromProto)
	out.Type = in.Type
	return out
}
