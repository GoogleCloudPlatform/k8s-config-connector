// Copyright 2024 Google LLC
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

package firewallpolicyrule

import (
	pb "cloud.google.com/go/compute/apiv1/computepb"
	krm "github.com/GoogleCloudPlatform/k8s-config-connector/apis/compute/v1beta1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct"
)

func ComputeFirewallPolicyRuleSpec_FromProto(mapCtx *direct.MapContext, in *pb.FirewallPolicyRule) *krm.ComputeFirewallPolicyRuleSpec {
	if in == nil {
		return nil
	}
	out := &krm.ComputeFirewallPolicyRuleSpec{}
	out.Action = in.GetAction()
	out.Description = in.Description
	out.Direction = in.GetDirection()
	out.Disabled = in.Disabled
	out.EnableLogging = in.EnableLogging
	out.Match = FirewallpolicyruleMatch_FromProto(mapCtx, in.Match)
	out.Priority = int64(in.GetPriority())
	// MISSING: RuleName
	// MISSING: SecurityProfileGroup
	out.TargetResources = ComputeFirewallPolicyRuleSpec_TargetResources_FromProto(mapCtx, in.TargetResources)
	// MISSING: TargetSecureTags
	out.TargetServiceAccounts = ComputeFirewallPolicyRuleSpec_TargetServiceAccounts_FromProto(mapCtx, in.TargetServiceAccounts)
	// MISSING: TlsInspect
	return out
}
func ComputeFirewallPolicyRuleSpec_ToProto(mapCtx *direct.MapContext, in *krm.ComputeFirewallPolicyRuleSpec) *pb.FirewallPolicyRule {
	if in == nil {
		return nil
	}
	out := &pb.FirewallPolicyRule{}
	out.Action = direct.LazyPtr(in.Action)
	out.Description = in.Description
	out.Direction = direct.LazyPtr(in.Direction)
	out.Disabled = in.Disabled
	out.EnableLogging = in.EnableLogging
	out.Match = FirewallpolicyruleMatch_ToProto(mapCtx, in.Match)
	out.Priority = direct.LazyPtr(int32(in.Priority))
	// MISSING: RuleName
	// MISSING: SecurityProfileGroup
	out.TargetResources = ComputeFirewallPolicyRuleSpec_TargetResources_ToProto(mapCtx, in.TargetResources)
	out.TargetServiceAccounts = ComputeFirewallPolicyRuleSpec_TargetServiceAccounts_ToProto(mapCtx, in.TargetServiceAccounts)
	// MISSING: TlsInspect
	return out
}
func ComputeFirewallPolicyRuleStatus_FromProto(mapCtx *direct.MapContext, in *pb.FirewallPolicyRule) *krm.ComputeFirewallPolicyRuleStatus {
	if in == nil {
		return nil
	}
	out := &krm.ComputeFirewallPolicyRuleStatus{}
	out.Kind = in.Kind
	out.RuleTupleCount = direct.LazyPtr(int64(in.GetRuleTupleCount()))
	// MISSING: TargetSecureTags
	return out
}
func ComputeFirewallPolicyRuleStatus_ToProto(mapCtx *direct.MapContext, in *krm.ComputeFirewallPolicyRuleStatus) *pb.FirewallPolicyRule {
	if in == nil {
		return nil
	}
	out := &pb.FirewallPolicyRule{}
	out.Kind = in.Kind
	out.RuleTupleCount = direct.LazyPtr(int32(*in.RuleTupleCount))
	// MISSING: TargetSecureTags
	return out
}

func FirewallpolicyruleLayer4Configs_FromProto(mapCtx *direct.MapContext, in *pb.FirewallPolicyRuleMatcherLayer4Config) *krm.FirewallPolicyRuleLayer4Configs {
	if in == nil {
		return nil
	}
	out := &krm.FirewallPolicyRuleLayer4Configs{}
	out.IPProtocol = in.GetIpProtocol()
	out.Ports = in.Ports
	return out
}
func FirewallpolicyruleLayer4Configs_ToProto(mapCtx *direct.MapContext, in *krm.FirewallPolicyRuleLayer4Configs) *pb.FirewallPolicyRuleMatcherLayer4Config {
	if in == nil {
		return nil
	}
	out := &pb.FirewallPolicyRuleMatcherLayer4Config{}
	out.IpProtocol = direct.LazyPtr(in.IPProtocol)
	out.Ports = in.Ports
	return out
}
func FirewallpolicyruleMatch_FromProto(mapCtx *direct.MapContext, in *pb.FirewallPolicyRuleMatcher) *krm.FirewallPolicyRuleMatch {
	if in == nil {
		return nil
	}
	out := &krm.FirewallPolicyRuleMatch{}
	out.DestAddressGroups = in.DestAddressGroups
	out.DestFqdns = in.DestFqdns
	out.DestIPRanges = in.DestIpRanges
	out.DestRegionCodes = in.DestRegionCodes
	out.DestThreatIntelligences = in.DestThreatIntelligences
	out.Layer4Configs = direct.Slice_FromProto(mapCtx, in.Layer4Configs, FirewallpolicyruleLayer4Configs_FromProto)
	out.SrcAddressGroups = in.SrcAddressGroups
	out.SrcFqdns = in.SrcFqdns
	out.SrcIPRanges = in.SrcIpRanges
	out.SrcRegionCodes = in.SrcRegionCodes
	// MISSING: SrcSecureTags
	out.SrcThreatIntelligences = in.SrcThreatIntelligences
	return out
}
func FirewallpolicyruleMatch_ToProto(mapCtx *direct.MapContext, in *krm.FirewallPolicyRuleMatch) *pb.FirewallPolicyRuleMatcher {
	if in == nil {
		return nil
	}
	out := &pb.FirewallPolicyRuleMatcher{}
	out.DestAddressGroups = in.DestAddressGroups
	out.DestFqdns = in.DestFqdns
	out.DestIpRanges = in.DestIPRanges
	out.DestRegionCodes = in.DestRegionCodes
	out.DestThreatIntelligences = in.DestThreatIntelligences
	out.Layer4Configs = direct.Slice_ToProto(mapCtx, in.Layer4Configs, FirewallpolicyruleLayer4Configs_ToProto)
	out.SrcAddressGroups = in.SrcAddressGroups
	out.SrcFqdns = in.SrcFqdns
	out.SrcIpRanges = in.SrcIPRanges
	out.SrcRegionCodes = in.SrcRegionCodes
	// MISSING: SrcSecureTags
	out.SrcThreatIntelligences = in.SrcThreatIntelligences
	return out
}
