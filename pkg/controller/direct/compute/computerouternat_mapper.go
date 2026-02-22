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
	refs "github.com/GoogleCloudPlatform/k8s-config-connector/apis/refs/v1beta1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct"
)

func ComputeRouterNATSpec_v1beta1_FromProto(mapCtx *direct.MapContext, in *pb.RouterNat) *krm.ComputeRouterNATSpec {
	if in == nil {
		return nil
	}
	out := &krm.ComputeRouterNATSpec{}
	out.AutoNetworkTier = in.AutoNetworkTier
	out.DrainNatIps = make([]refs.ComputeAddressRef, len(in.DrainNatIps))
	for i, v := range in.DrainNatIps {
		out.DrainNatIps[i] = ComputeAddressRef_v1beta1_FromProto(mapCtx, v)
	}
	out.EnableDynamicPortAllocation = in.EnableDynamicPortAllocation
	out.EnableEndpointIndependentMapping = in.EnableEndpointIndependentMapping
	out.EndpointTypes = in.EndpointTypes
	out.IcmpIdleTimeoutSec = in.IcmpIdleTimeoutSec
	out.LogConfig = RouterNatLogConfig_v1beta1_FromProto(mapCtx, in.GetLogConfig())
	out.MaxPortsPerVm = in.MaxPortsPerVm
	out.MinPortsPerVm = in.MinPortsPerVm
	out.Nat64Subnetworks = direct.Slice_FromProto(mapCtx, in.Nat64Subnetworks, RouterNat64Subnetwork_v1beta1_FromProto)
	out.NatIpAllocateOption = in.NatIpAllocateOption
	out.NatIps = make([]refs.ComputeAddressRef, len(in.NatIps))
	for i, v := range in.NatIps {
		out.NatIps[i] = ComputeAddressRef_v1beta1_FromProto(mapCtx, v)
	}
	out.Rules = direct.Slice_FromProto(mapCtx, in.Rules, RouterNatRule_v1beta1_FromProto)
	out.SourceSubnetworkIpRangesToNat = in.GetSourceSubnetworkIpRangesToNat()
	out.SourceSubnetworkIpRangesToNat64 = in.SourceSubnetworkIpRangesToNat64
	out.Subnetwork = direct.Slice_FromProto(mapCtx, in.Subnetworks, RouterNatSubnetwork_v1beta1_FromProto)
	out.TcpEstablishedIdleTimeoutSec = in.TcpEstablishedIdleTimeoutSec
	out.TcpTimeWaitTimeoutSec = in.TcpTimeWaitTimeoutSec
	out.TcpTransitoryIdleTimeoutSec = in.TcpTransitoryIdleTimeoutSec
	out.Type = in.Type
	out.UdpIdleTimeoutSec = in.UdpIdleTimeoutSec
	return out
}

func ComputeRouterNATSpec_v1beta1_ToProto(mapCtx *direct.MapContext, in *krm.ComputeRouterNATSpec) *pb.RouterNat {
	if in == nil {
		return nil
	}
	out := &pb.RouterNat{}
	out.AutoNetworkTier = in.AutoNetworkTier
	out.DrainNatIps = make([]string, len(in.DrainNatIps))
	for i, v := range in.DrainNatIps {
		out.DrainNatIps[i] = ComputeAddressRef_v1beta1_ToProto(mapCtx, v)
	}
	out.EnableDynamicPortAllocation = in.EnableDynamicPortAllocation
	out.EnableEndpointIndependentMapping = in.EnableEndpointIndependentMapping
	out.EndpointTypes = in.EndpointTypes
	out.IcmpIdleTimeoutSec = in.IcmpIdleTimeoutSec
	out.LogConfig = RouterNatLogConfig_v1beta1_ToProto(mapCtx, in.LogConfig)
	out.MaxPortsPerVm = in.MaxPortsPerVm
	out.MinPortsPerVm = in.MinPortsPerVm
	out.Nat64Subnetworks = direct.Slice_ToProto(mapCtx, in.Nat64Subnetworks, RouterNat64Subnetwork_v1beta1_ToProto)
	out.NatIpAllocateOption = in.NatIpAllocateOption
	out.NatIps = make([]string, len(in.NatIps))
	for i, v := range in.NatIps {
		out.NatIps[i] = ComputeAddressRef_v1beta1_ToProto(mapCtx, v)
	}
	out.Rules = direct.Slice_ToProto(mapCtx, in.Rules, RouterNatRule_v1beta1_ToProto)
	out.SourceSubnetworkIpRangesToNat = direct.PtrTo(in.SourceSubnetworkIpRangesToNat)
	out.SourceSubnetworkIpRangesToNat64 = in.SourceSubnetworkIpRangesToNat64
	out.Subnetworks = direct.Slice_ToProto(mapCtx, in.Subnetwork, RouterNatSubnetwork_v1beta1_ToProto)
	out.TcpEstablishedIdleTimeoutSec = in.TcpEstablishedIdleTimeoutSec
	out.TcpTimeWaitTimeoutSec = in.TcpTimeWaitTimeoutSec
	out.TcpTransitoryIdleTimeoutSec = in.TcpTransitoryIdleTimeoutSec
	out.Type = in.Type
	out.UdpIdleTimeoutSec = in.UdpIdleTimeoutSec
	return out
}

func ComputeAddressRef_v1beta1_FromProto(mapCtx *direct.MapContext, in string) refs.ComputeAddressRef {
	return refs.ComputeAddressRef{
		External: in,
	}
}

func ComputeAddressRef_v1beta1_ToProto(mapCtx *direct.MapContext, in refs.ComputeAddressRef) string {
	if in.External == "" && in.Name != "" {
		mapCtx.Errorf("reference %s was not pre-resolved", in.Name)
	}
	return in.External
}

func ComputeSubnetworkRef_v1beta1_FromProto(mapCtx *direct.MapContext, in string) refs.ComputeSubnetworkRef {
	return refs.ComputeSubnetworkRef{
		External: in,
	}
}

func ComputeSubnetworkRef_v1beta1_ToProto(mapCtx *direct.MapContext, in refs.ComputeSubnetworkRef) string {
	if in.External == "" && in.Name != "" {
		mapCtx.Errorf("reference %s was not pre-resolved", in.Name)
	}
	return in.External
}

func RouterNatLogConfig_v1beta1_FromProto(mapCtx *direct.MapContext, in *pb.RouterNatLogConfig) *krm.RouterNatLogConfig {
	if in == nil {
		return nil
	}
	out := &krm.RouterNatLogConfig{}
	out.Enable = in.Enable
	out.Filter = in.Filter
	return out
}

func RouterNatLogConfig_v1beta1_ToProto(mapCtx *direct.MapContext, in *krm.RouterNatLogConfig) *pb.RouterNatLogConfig {
	if in == nil {
		return nil
	}
	out := &pb.RouterNatLogConfig{}
	out.Enable = in.Enable
	out.Filter = in.Filter
	return out
}

func RouterNatRule_v1beta1_FromProto(mapCtx *direct.MapContext, in *pb.RouterNatRule) *krm.RouterNatRule {
	if in == nil {
		return nil
	}
	out := &krm.RouterNatRule{}
	out.Action = RouterNatRuleAction_v1beta1_FromProto(mapCtx, in.Action)
	out.Description = in.Description
	out.Match = in.Match
	out.RuleNumber = in.RuleNumber
	return out
}

func RouterNatRule_v1beta1_ToProto(mapCtx *direct.MapContext, in *krm.RouterNatRule) *pb.RouterNatRule {
	if in == nil {
		return nil
	}
	out := &pb.RouterNatRule{}
	out.Action = RouterNatRuleAction_v1beta1_ToProto(mapCtx, in.Action)
	out.Description = in.Description
	out.Match = in.Match
	out.RuleNumber = in.RuleNumber
	return out
}

func RouterNatRuleAction_v1beta1_FromProto(mapCtx *direct.MapContext, in *pb.RouterNatRuleAction) *krm.RouterNatRuleAction {
	if in == nil {
		return nil
	}
	out := &krm.RouterNatRuleAction{}
	out.SourceNatActiveIpsRefs = make([]refs.ComputeAddressRef, len(in.SourceNatActiveIps))
	for i, v := range in.SourceNatActiveIps {
		out.SourceNatActiveIpsRefs[i] = ComputeAddressRef_v1beta1_FromProto(mapCtx, v)
	}
	out.SourceNatActiveRangesRefs = make([]refs.ComputeSubnetworkRef, len(in.SourceNatActiveRanges))
	for i, v := range in.SourceNatActiveRanges {
		out.SourceNatActiveRangesRefs[i] = ComputeSubnetworkRef_v1beta1_FromProto(mapCtx, v)
	}
	out.SourceNatDrainIpsRefs = make([]refs.ComputeAddressRef, len(in.SourceNatDrainIps))
	for i, v := range in.SourceNatDrainIps {
		out.SourceNatDrainIpsRefs[i] = ComputeAddressRef_v1beta1_FromProto(mapCtx, v)
	}
	out.SourceNatDrainRangesRefs = make([]refs.ComputeSubnetworkRef, len(in.SourceNatDrainRanges))
	for i, v := range in.SourceNatDrainRanges {
		out.SourceNatDrainRangesRefs[i] = ComputeSubnetworkRef_v1beta1_FromProto(mapCtx, v)
	}
	return out
}

func RouterNatRuleAction_v1beta1_ToProto(mapCtx *direct.MapContext, in *krm.RouterNatRuleAction) *pb.RouterNatRuleAction {
	if in == nil {
		return nil
	}
	out := &pb.RouterNatRuleAction{}
	out.SourceNatActiveIps = make([]string, len(in.SourceNatActiveIpsRefs))
	for i, v := range in.SourceNatActiveIpsRefs {
		out.SourceNatActiveIps[i] = ComputeAddressRef_v1beta1_ToProto(mapCtx, v)
	}
	out.SourceNatActiveRanges = make([]string, len(in.SourceNatActiveRangesRefs))
	for i, v := range in.SourceNatActiveRangesRefs {
		out.SourceNatActiveRanges[i] = ComputeSubnetworkRef_v1beta1_ToProto(mapCtx, v)
	}
	out.SourceNatDrainIps = make([]string, len(in.SourceNatDrainIpsRefs))
	for i, v := range in.SourceNatDrainIpsRefs {
		out.SourceNatDrainIps[i] = ComputeAddressRef_v1beta1_ToProto(mapCtx, v)
	}
	out.SourceNatDrainRanges = make([]string, len(in.SourceNatDrainRangesRefs))
	for i, v := range in.SourceNatDrainRangesRefs {
		out.SourceNatDrainRanges[i] = ComputeSubnetworkRef_v1beta1_ToProto(mapCtx, v)
	}
	return out
}

func RouterNatSubnetwork_v1beta1_FromProto(mapCtx *direct.MapContext, in *pb.RouterNatSubnetworkToNat) *krm.RouterNatSubnetwork {
	if in == nil {
		return nil
	}
	out := &krm.RouterNatSubnetwork{}
	out.SubnetworkRef = ComputeSubnetworkRef_v1beta1_FromProto(mapCtx, in.GetName())
	out.SecondaryIpRangeNames = in.SecondaryIpRangeNames
	out.SourceIpRangesToNat = in.SourceIpRangesToNat
	return out
}

func RouterNatSubnetwork_v1beta1_ToProto(mapCtx *direct.MapContext, in *krm.RouterNatSubnetwork) *pb.RouterNatSubnetworkToNat {
	if in == nil {
		return nil
	}
	out := &pb.RouterNatSubnetworkToNat{}
	out.Name = direct.PtrTo(ComputeSubnetworkRef_v1beta1_ToProto(mapCtx, in.SubnetworkRef))
	out.SecondaryIpRangeNames = in.SecondaryIpRangeNames
	out.SourceIpRangesToNat = in.SourceIpRangesToNat
	return out
}

func RouterNat64Subnetwork_v1beta1_FromProto(mapCtx *direct.MapContext, in *pb.RouterNatSubnetworkToNat64) *krm.RouterNat64Subnetwork {
	if in == nil {
		return nil
	}
	out := &krm.RouterNat64Subnetwork{}
	out.SubnetworkRef = ComputeSubnetworkRef_v1beta1_FromProto(mapCtx, in.GetName())
	return out
}

func RouterNat64Subnetwork_v1beta1_ToProto(mapCtx *direct.MapContext, in *krm.RouterNat64Subnetwork) *pb.RouterNatSubnetworkToNat64 {
	if in == nil {
		return nil
	}
	out := &pb.RouterNatSubnetworkToNat64{}
	out.Name = direct.PtrTo(ComputeSubnetworkRef_v1beta1_ToProto(mapCtx, in.SubnetworkRef))
	return out
}
