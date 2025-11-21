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

// +generated:mapper
// krm.group: compute.cnrm.cloud.google.com
// krm.version: v1beta1
// proto.service: google.cloud.compute.v1

package compute

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
	out.Action = in.Action
	out.Description = in.Description
	out.Direction = in.Direction
	out.Disabled = in.Disabled
	out.EnableLogging = in.EnableLogging
	// MISSING: Kind
	out.Match = FirewallPolicyRuleMatcher_FromProto(mapCtx, in.GetMatch())
	out.Priority = in.Priority
	// MISSING: RuleName
	// MISSING: RuleTupleCount
	// MISSING: SecurityProfileGroup
	out.TargetResources = ComputeFirewallPolicyRuleSpec_TargetResources_FromProto(mapCtx, in.TargetResources)
	// MISSING: TargetSecureTags
	out.TargetServiceAccounts = ComputeFirewallPolicyRuleSpec_TargetServiceAccounts_FromProto(mapCtx, in.TargetServiceAccounts)
	// MISSING: TLSInspect
	return out
}
func ComputeFirewallPolicyRuleSpec_ToProto(mapCtx *direct.MapContext, in *krm.ComputeFirewallPolicyRuleSpec) *pb.FirewallPolicyRule {
	if in == nil {
		return nil
	}
	out := &pb.FirewallPolicyRule{}
	out.Action = in.Action
	out.Description = in.Description
	out.Direction = in.Direction
	out.Disabled = in.Disabled
	out.EnableLogging = in.EnableLogging
	// MISSING: Kind
	out.Match = FirewallPolicyRuleMatcher_ToProto(mapCtx, in.Match)
	out.Priority = in.Priority
	// MISSING: RuleName
	// MISSING: RuleTupleCount
	// MISSING: SecurityProfileGroup
	out.TargetResources = ComputeFirewallPolicyRuleSpec_TargetResources_ToProto(mapCtx, in.TargetResources)
	// MISSING: TargetSecureTags
	out.TargetServiceAccounts = ComputeFirewallPolicyRuleSpec_TargetServiceAccounts_ToProto(mapCtx, in.TargetServiceAccounts)
	// MISSING: TLSInspect
	return out
}
func ComputeFirewallPolicyRuleStatus_FromProto(mapCtx *direct.MapContext, in *pb.FirewallPolicyRule) *krm.ComputeFirewallPolicyRuleStatus {
	if in == nil {
		return nil
	}
	out := &krm.ComputeFirewallPolicyRuleStatus{}
	// MISSING: Action
	// MISSING: Description
	// MISSING: Direction
	// MISSING: Disabled
	// MISSING: EnableLogging
	out.Kind = in.Kind
	// MISSING: Match
	// MISSING: Priority
	// MISSING: RuleName
	out.RuleTupleCount = in.RuleTupleCount
	// MISSING: SecurityProfileGroup
	// MISSING: TargetResources
	// MISSING: TargetSecureTags
	// MISSING: TargetServiceAccounts
	// MISSING: TLSInspect
	return out
}
func ComputeFirewallPolicyRuleStatus_ToProto(mapCtx *direct.MapContext, in *krm.ComputeFirewallPolicyRuleStatus) *pb.FirewallPolicyRule {
	if in == nil {
		return nil
	}
	out := &pb.FirewallPolicyRule{}
	// MISSING: Action
	// MISSING: Description
	// MISSING: Direction
	// MISSING: Disabled
	// MISSING: EnableLogging
	out.Kind = in.Kind
	// MISSING: Match
	// MISSING: Priority
	// MISSING: RuleName
	out.RuleTupleCount = in.RuleTupleCount
	// MISSING: SecurityProfileGroup
	// MISSING: TargetResources
	// MISSING: TargetSecureTags
	// MISSING: TargetServiceAccounts
	// MISSING: TLSInspect
	return out
}
func FirewallPolicyRuleMatcher_FromProto(mapCtx *direct.MapContext, in *pb.FirewallPolicyRuleMatcher) *krm.FirewallPolicyRuleMatcher {
	if in == nil {
		return nil
	}
	out := &krm.FirewallPolicyRuleMatcher{}
	out.DestAddressGroups = in.DestAddressGroups
	out.DestFqdns = in.DestFqdns
	out.DestIPRanges = in.DestIpRanges
	// MISSING: DestNetworkType
	out.DestRegionCodes = in.DestRegionCodes
	out.DestThreatIntelligences = in.DestThreatIntelligences
	out.Layer4Configs = direct.Slice_FromProto(mapCtx, in.Layer4Configs, FirewallPolicyRuleMatcherLayer4Config_FromProto)
	out.SrcAddressGroups = in.SrcAddressGroups
	out.SrcFqdns = in.SrcFqdns
	out.SrcIPRanges = in.SrcIpRanges
	// MISSING: SrcNetworkType
	// MISSING: SrcNetworks
	out.SrcRegionCodes = in.SrcRegionCodes
	// MISSING: SrcSecureTags
	out.SrcThreatIntelligences = in.SrcThreatIntelligences
	return out
}
func FirewallPolicyRuleMatcher_ToProto(mapCtx *direct.MapContext, in *krm.FirewallPolicyRuleMatcher) *pb.FirewallPolicyRuleMatcher {
	if in == nil {
		return nil
	}
	out := &pb.FirewallPolicyRuleMatcher{}
	out.DestAddressGroups = in.DestAddressGroups
	out.DestFqdns = in.DestFqdns
	out.DestIpRanges = in.DestIPRanges
	// MISSING: DestNetworkType
	out.DestRegionCodes = in.DestRegionCodes
	out.DestThreatIntelligences = in.DestThreatIntelligences
	out.Layer4Configs = direct.Slice_ToProto(mapCtx, in.Layer4Configs, FirewallPolicyRuleMatcherLayer4Config_ToProto)
	out.SrcAddressGroups = in.SrcAddressGroups
	out.SrcFqdns = in.SrcFqdns
	out.SrcIpRanges = in.SrcIPRanges
	// MISSING: SrcNetworkType
	// MISSING: SrcNetworks
	out.SrcRegionCodes = in.SrcRegionCodes
	// MISSING: SrcSecureTags
	out.SrcThreatIntelligences = in.SrcThreatIntelligences
	return out
}
func FirewallPolicyRuleMatcherLayer4Config_FromProto(mapCtx *direct.MapContext, in *pb.FirewallPolicyRuleMatcherLayer4Config) *krm.FirewallPolicyRuleMatcherLayer4Config {
	if in == nil {
		return nil
	}
	out := &krm.FirewallPolicyRuleMatcherLayer4Config{}
	out.IPProtocol = in.IpProtocol
	out.Ports = in.Ports
	return out
}
func FirewallPolicyRuleMatcherLayer4Config_ToProto(mapCtx *direct.MapContext, in *krm.FirewallPolicyRuleMatcherLayer4Config) *pb.FirewallPolicyRuleMatcherLayer4Config {
	if in == nil {
		return nil
	}
	out := &pb.FirewallPolicyRuleMatcherLayer4Config{}
	out.IpProtocol = in.IPProtocol
	out.Ports = in.Ports
	return out
}
func FirewallPolicyRuleSecureTag_FromProto(mapCtx *direct.MapContext, in *pb.FirewallPolicyRuleSecureTag) *krm.FirewallPolicyRuleSecureTag {
	if in == nil {
		return nil
	}
	out := &krm.FirewallPolicyRuleSecureTag{}
	out.Name = in.Name
	out.State = in.State
	return out
}
func FirewallPolicyRuleSecureTag_ToProto(mapCtx *direct.MapContext, in *krm.FirewallPolicyRuleSecureTag) *pb.FirewallPolicyRuleSecureTag {
	if in == nil {
		return nil
	}
	out := &pb.FirewallPolicyRuleSecureTag{}
	out.Name = in.Name
	out.State = in.State
	return out
}
