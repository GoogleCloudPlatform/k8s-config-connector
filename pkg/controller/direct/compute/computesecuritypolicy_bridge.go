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

package compute

import (
	pb "cloud.google.com/go/compute/apiv1/computepb"
	krm "github.com/GoogleCloudPlatform/k8s-config-connector/apis/compute/v1beta1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct/compute/computesecuritypolicy"
)

func SecurityPolicyAdaptiveProtectionConfig_v1beta1_FromProto(mapCtx *direct.MapContext, in *pb.SecurityPolicyAdaptiveProtectionConfig) *krm.SecurityPolicyAdaptiveProtectionConfig {
	return computesecuritypolicy.SecurityPolicyAdaptiveProtectionConfig_v1beta1_FromProto(mapCtx, in)
}

func SecurityPolicyRecaptchaOptionsConfig_v1beta1_FromProto(mapCtx *direct.MapContext, in *pb.SecurityPolicyRecaptchaOptionsConfig) *krm.SecurityPolicyRecaptchaOptionsConfig {
	return computesecuritypolicy.SecurityPolicyRecaptchaOptionsConfig_v1beta1_FromProto(mapCtx, in)
}

func SecurityPolicyAdaptiveProtectionConfig_v1beta1_ToProto(mapCtx *direct.MapContext, in *krm.SecurityPolicyAdaptiveProtectionConfig) *pb.SecurityPolicyAdaptiveProtectionConfig {
	return computesecuritypolicy.SecurityPolicyAdaptiveProtectionConfig_v1beta1_ToProto(mapCtx, in)
}

func SecurityPolicyRecaptchaOptionsConfig_v1beta1_ToProto(mapCtx *direct.MapContext, in *krm.SecurityPolicyRecaptchaOptionsConfig) *pb.SecurityPolicyRecaptchaOptionsConfig {
	return computesecuritypolicy.SecurityPolicyRecaptchaOptionsConfig_v1beta1_ToProto(mapCtx, in)
}

func SecurityPolicyRulePreconfiguredWafConfig_v1beta1_FromProto(mapCtx *direct.MapContext, in *pb.SecurityPolicyRulePreconfiguredWafConfig) *krm.SecurityPolicyRulePreconfiguredWafConfig {
	return computesecuritypolicy.SecurityPolicyRulePreconfiguredWafConfig_v1beta1_FromProto(mapCtx, in)
}

func SecurityPolicyRulePreconfiguredWafConfig_v1beta1_ToProto(mapCtx *direct.MapContext, in *krm.SecurityPolicyRulePreconfiguredWafConfig) *pb.SecurityPolicyRulePreconfiguredWafConfig {
	return computesecuritypolicy.SecurityPolicyRulePreconfiguredWafConfig_v1beta1_ToProto(mapCtx, in)
}

func SecurityPolicyRuleMatcherExpr_v1beta1_FromProto(mapCtx *direct.MapContext, in *pb.Expr) *krm.SecurityPolicyRuleMatcherExpr {
	return computesecuritypolicy.SecurityPolicyRuleMatcherExpr_v1beta1_FromProto(mapCtx, in)
}

func SecurityPolicyRuleMatcherExpr_v1beta1_ToProto(mapCtx *direct.MapContext, in *krm.SecurityPolicyRuleMatcherExpr) *pb.Expr {
	return computesecuritypolicy.SecurityPolicyRuleMatcherExpr_v1beta1_ToProto(mapCtx, in)
}
