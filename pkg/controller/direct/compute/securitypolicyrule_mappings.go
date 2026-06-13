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
	"strconv"

	pb "cloud.google.com/go/compute/apiv1/computepb"
	krm "github.com/GoogleCloudPlatform/k8s-config-connector/apis/compute/v1alpha1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct"
)

// ComputeOrganizationSecurityPolicyRuleSpec_v1alpha1_FromProto maps FirewallPolicyRule proto to ComputeOrganizationSecurityPolicyRuleSpec KRM.
// Handwritten because of string/pointer conversions (e.g. Action) and custom nested matching configurations.
func ComputeOrganizationSecurityPolicyRuleSpec_v1alpha1_FromProto(mapCtx *direct.MapContext, in *pb.FirewallPolicyRule) *krm.ComputeOrganizationSecurityPolicyRuleSpec {
	if in == nil {
		return nil
	}
	out := &krm.ComputeOrganizationSecurityPolicyRuleSpec{}
	out.Action = in.GetAction()
	out.Description = in.Description
	out.Direction = direct.LazyPtr(in.GetDirection())
	out.EnableLogging = in.EnableLogging
	if in.Match != nil {
		out.Match = *OrganizationsecuritypolicyruleMatch_v1alpha1_FromProto(mapCtx, in.Match)
	}
	// PolicyId is passed as a path parameter rather than a field in FirewallPolicyRule itself.
	// It is left to be populated by the reconciler logic.
	// Preview is not supported on FirewallPolicyRule proto, so it is ignored.
	if in.Priority != nil {
		out.ResourceID = direct.LazyPtr(strconv.FormatInt(int64(*in.Priority), 10))
	}
	out.TargetResources = in.TargetResources
	out.TargetServiceAccounts = in.TargetServiceAccounts
	return out
}

// ComputeOrganizationSecurityPolicyRuleSpec_v1alpha1_ToProto maps ComputeOrganizationSecurityPolicyRuleSpec KRM to FirewallPolicyRule proto.
// Handwritten because of string/pointer conversions (e.g. Action) and custom nested matching configurations.
func ComputeOrganizationSecurityPolicyRuleSpec_v1alpha1_ToProto(mapCtx *direct.MapContext, in *krm.ComputeOrganizationSecurityPolicyRuleSpec) *pb.FirewallPolicyRule {
	if in == nil {
		return nil
	}
	out := &pb.FirewallPolicyRule{}
	out.Action = direct.LazyPtr(in.Action)
	out.Description = in.Description
	if in.Direction != nil {
		out.Direction = direct.LazyPtr(*in.Direction)
	}
	out.EnableLogging = in.EnableLogging
	out.Match = OrganizationsecuritypolicyruleMatch_v1alpha1_ToProto(mapCtx, &in.Match)
	// Preview is not supported on FirewallPolicyRule proto, so it is ignored.
	if in.ResourceID != nil {
		priorityVal, err := strconv.ParseInt(*in.ResourceID, 10, 32)
		if err != nil {
			mapCtx.Errorf("invalid priority / resourceID: %v", err)
		} else {
			out.Priority = direct.LazyPtr(int32(priorityVal))
		}
	}
	out.TargetResources = in.TargetResources
	out.TargetServiceAccounts = in.TargetServiceAccounts
	return out
}

func OrganizationsecuritypolicyruleMatch_v1alpha1_FromProto(mapCtx *direct.MapContext, in *pb.FirewallPolicyRuleMatcher) *krm.OrganizationsecuritypolicyruleMatch {
	if in == nil {
		return nil
	}
	out := &krm.OrganizationsecuritypolicyruleMatch{}
	// Description and VersionedExpr are not present on pb.FirewallPolicyRuleMatcher, so they are skipped.
	config := OrganizationsecuritypolicyruleConfig_v1alpha1_FromProto(mapCtx, in)
	if config != nil {
		out.Config = *config
	}
	return out
}

func OrganizationsecuritypolicyruleMatch_v1alpha1_ToProto(mapCtx *direct.MapContext, in *krm.OrganizationsecuritypolicyruleMatch) *pb.FirewallPolicyRuleMatcher {
	if in == nil {
		return nil
	}
	// If the config is completely empty, treat the Match as empty/nil in Proto to avoid roundtrip mismatches.
	if len(in.Config.DestIpRanges) == 0 && len(in.Config.SrcIpRanges) == 0 && len(in.Config.Layer4Config) == 0 {
		return nil
	}
	out := &pb.FirewallPolicyRuleMatcher{}
	// Description and VersionedExpr are not present on pb.FirewallPolicyRuleMatcher, so they are skipped.
	// Flatten the config fields into the FirewallPolicyRuleMatcher:
	out.DestIpRanges = in.Config.DestIpRanges
	out.SrcIpRanges = in.Config.SrcIpRanges
	out.Layer4Configs = OrganizationsecuritypolicyruleLayer4Config_v1alpha1_SliceToProto(mapCtx, in.Config.Layer4Config)
	return out
}

func OrganizationsecuritypolicyruleConfig_v1alpha1_FromProto(mapCtx *direct.MapContext, in *pb.FirewallPolicyRuleMatcher) *krm.OrganizationsecuritypolicyruleConfig {
	if in == nil {
		return nil
	}
	out := &krm.OrganizationsecuritypolicyruleConfig{}
	out.DestIpRanges = in.DestIpRanges
	out.SrcIpRanges = in.SrcIpRanges
	out.Layer4Config = OrganizationsecuritypolicyruleLayer4Config_v1alpha1_SliceFromProto(mapCtx, in.Layer4Configs)
	return out
}

func OrganizationsecuritypolicyruleLayer4Config_v1alpha1_FromProto(mapCtx *direct.MapContext, in *pb.FirewallPolicyRuleMatcherLayer4Config) *krm.OrganizationsecuritypolicyruleLayer4Config {
	if in == nil {
		return nil
	}
	out := &krm.OrganizationsecuritypolicyruleLayer4Config{}
	out.IpProtocol = in.GetIpProtocol()
	out.Ports = in.Ports
	return out
}

func OrganizationsecuritypolicyruleLayer4Config_v1alpha1_ToProto(mapCtx *direct.MapContext, in *krm.OrganizationsecuritypolicyruleLayer4Config) *pb.FirewallPolicyRuleMatcherLayer4Config {
	if in == nil {
		return nil
	}
	out := &pb.FirewallPolicyRuleMatcherLayer4Config{}
	out.IpProtocol = direct.LazyPtr(in.IpProtocol)
	out.Ports = in.Ports
	return out
}

func OrganizationsecuritypolicyruleLayer4Config_v1alpha1_SliceFromProto(mapCtx *direct.MapContext, in []*pb.FirewallPolicyRuleMatcherLayer4Config) []krm.OrganizationsecuritypolicyruleLayer4Config {
	if in == nil {
		return nil
	}
	out := make([]krm.OrganizationsecuritypolicyruleLayer4Config, len(in))
	for i, v := range in {
		if v != nil {
			out[i] = *OrganizationsecuritypolicyruleLayer4Config_v1alpha1_FromProto(mapCtx, v)
		}
	}
	return out
}

func OrganizationsecuritypolicyruleLayer4Config_v1alpha1_SliceToProto(mapCtx *direct.MapContext, in []krm.OrganizationsecuritypolicyruleLayer4Config) []*pb.FirewallPolicyRuleMatcherLayer4Config {
	if in == nil {
		return nil
	}
	out := make([]*pb.FirewallPolicyRuleMatcherLayer4Config, len(in))
	for i, v := range in {
		out[i] = OrganizationsecuritypolicyruleLayer4Config_v1alpha1_ToProto(mapCtx, &v)
	}
	return out
}

// ComputeOrganizationSecurityPolicyRuleObservedState_v1alpha1_FromProto is a dummy mapper for fuzztesting.
func ComputeOrganizationSecurityPolicyRuleObservedState_v1alpha1_FromProto(mapCtx *direct.MapContext, in *pb.FirewallPolicyRule) *any {
	return nil
}

// ComputeOrganizationSecurityPolicyRuleObservedState_v1alpha1_ToProto is a dummy mapper for fuzztesting.
func ComputeOrganizationSecurityPolicyRuleObservedState_v1alpha1_ToProto(mapCtx *direct.MapContext, in *any) *pb.FirewallPolicyRule {
	return nil
}
