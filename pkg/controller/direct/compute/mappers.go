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
)

func RouternatAction_v1beta1_ToProto(mapCtx *direct.MapContext, in *krm.RouternatAction) *pb.RouterNatRuleAction {
	if in == nil {
		return nil
	}
	out := &pb.RouterNatRuleAction{}
	for _, ref := range in.SourceNatActiveIpsRefs {
		if ref.External != "" {
			out.SourceNatActiveIps = append(out.SourceNatActiveIps, ref.External)
		}
	}
	for _, ref := range in.SourceNatDrainIpsRefs {
		if ref.External != "" {
			out.SourceNatDrainIps = append(out.SourceNatDrainIps, ref.External)
		}
	}
	return out
}

func RouternatAction_v1beta1_FromProto(mapCtx *direct.MapContext, in *pb.RouterNatRuleAction) *krm.RouternatAction {
	if in == nil {
		return nil
	}
	out := &krm.RouternatAction{}
	for _, ip := range in.SourceNatActiveIps {
		out.SourceNatActiveIpsRefs = append(out.SourceNatActiveIpsRefs, krm.ResourceRef{External: ip})
	}
	for _, ip := range in.SourceNatDrainIps {
		out.SourceNatDrainIpsRefs = append(out.SourceNatDrainIpsRefs, krm.ResourceRef{External: ip})
	}
	return out
}

func RouternatLogConfig_v1beta1_ToProto(mapCtx *direct.MapContext, in *krm.RouternatLogConfig) *pb.RouterNatLogConfig {
	if in == nil {
		return nil
	}
	out := &pb.RouterNatLogConfig{}
	out.Enable = &in.Enable
	out.Filter = &in.Filter
	return out
}

func RouternatLogConfig_v1beta1_FromProto(mapCtx *direct.MapContext, in *pb.RouterNatLogConfig) *krm.RouternatLogConfig {
	if in == nil {
		return nil
	}
	out := &krm.RouternatLogConfig{}
	if in.Enable != nil {
		out.Enable = *in.Enable
	}
	if in.Filter != nil {
		out.Filter = *in.Filter
	}
	return out
}

func RouternatRules_v1beta1_ToProto(mapCtx *direct.MapContext, in *krm.RouternatRules) *pb.RouterNatRule {
	if in == nil {
		return nil
	}
	out := &pb.RouterNatRule{}
	out.Action = RouternatAction_v1beta1_ToProto(mapCtx, in.Action)
	out.Description = in.Description
	out.Match = &in.Match
	val := uint32(in.RuleNumber)
	out.RuleNumber = &val
	return out
}

func RouternatRules_v1beta1_FromProto(mapCtx *direct.MapContext, in *pb.RouterNatRule) *krm.RouternatRules {
	if in == nil {
		return nil
	}
	out := &krm.RouternatRules{}
	out.Action = RouternatAction_v1beta1_FromProto(mapCtx, in.GetAction())
	out.Description = in.Description
	if in.Match != nil {
		out.Match = *in.Match
	}
	if in.RuleNumber != nil {
		out.RuleNumber = int64(*in.RuleNumber)
	}
	return out
}

func RouternatSubnetwork_v1beta1_ToProto(mapCtx *direct.MapContext, in *krm.RouternatSubnetwork) *pb.RouterNatSubnetworkToNat {
	if in == nil {
		return nil
	}
	out := &pb.RouterNatSubnetworkToNat{}
	if in.SubnetworkRef.External != "" {
		out.Name = &in.SubnetworkRef.External
	}
	out.SecondaryIpRangeNames = in.SecondaryIpRangeNames
	out.SourceIpRangesToNat = in.SourceIpRangesToNat
	return out
}

func RouternatSubnetwork_v1beta1_FromProto(mapCtx *direct.MapContext, in *pb.RouterNatSubnetworkToNat) *krm.RouternatSubnetwork {
	if in == nil {
		return nil
	}
	out := &krm.RouternatSubnetwork{}
	if in.Name != nil {
		out.SubnetworkRef.External = *in.Name
	}
	out.SecondaryIpRangeNames = in.SecondaryIpRangeNames
	out.SourceIpRangesToNat = in.SourceIpRangesToNat
	return out
}

func ComputeRouterNATSpec_v1beta1_ToProto(mapCtx *direct.MapContext, in *krm.ComputeRouterNATSpec) *pb.RouterNat {
	if in == nil {
		return nil
	}
	out := &pb.RouterNat{}
	for _, ref := range in.DrainNatIps {
		if ref.External != "" {
			out.DrainNatIps = append(out.DrainNatIps, ref.External)
		}
	}
	out.EnableDynamicPortAllocation = in.EnableDynamicPortAllocation
	out.EnableEndpointIndependentMapping = in.EnableEndpointIndependentMapping
	if in.IcmpIdleTimeoutSec != nil {
		val := int32(*in.IcmpIdleTimeoutSec)
		out.IcmpIdleTimeoutSec = &val
	}
	out.LogConfig = RouternatLogConfig_v1beta1_ToProto(mapCtx, in.LogConfig)
	if in.MaxPortsPerVm != nil {
		val := int32(*in.MaxPortsPerVm)
		out.MaxPortsPerVm = &val
	}
	if in.MinPortsPerVm != nil {
		val := int32(*in.MinPortsPerVm)
		out.MinPortsPerVm = &val
	}
	out.NatIpAllocateOption = &in.NatIpAllocateOption
	for _, ref := range in.NatIps {
		if ref.External != "" {
			out.NatIps = append(out.NatIps, ref.External)
		}
	}
	// Rules mapping
	out.Rules = direct.Slice_ToProto(mapCtx, in.Rules, RouternatRules_v1beta1_ToProto)
	out.SourceSubnetworkIpRangesToNat = &in.SourceSubnetworkIpRangesToNat
	out.Subnetworks = direct.Slice_ToProto(mapCtx, in.Subnetwork, RouternatSubnetwork_v1beta1_ToProto)
	if in.TcpEstablishedIdleTimeoutSec != nil {
		val := int32(*in.TcpEstablishedIdleTimeoutSec)
		out.TcpEstablishedIdleTimeoutSec = &val
	}
	if in.TcpTimeWaitTimeoutSec != nil {
		val := int32(*in.TcpTimeWaitTimeoutSec)
		out.TcpTimeWaitTimeoutSec = &val
	}
	if in.TcpTransitoryIdleTimeoutSec != nil {
		val := int32(*in.TcpTransitoryIdleTimeoutSec)
		out.TcpTransitoryIdleTimeoutSec = &val
	}
	if in.UdpIdleTimeoutSec != nil {
		val := int32(*in.UdpIdleTimeoutSec)
		out.UdpIdleTimeoutSec = &val
	}
	return out
}

func ComputeRouterNATSpec_v1beta1_FromProto(mapCtx *direct.MapContext, in *pb.RouterNat) *krm.ComputeRouterNATSpec {
	if in == nil {
		return nil
	}
	out := &krm.ComputeRouterNATSpec{}
	for _, ip := range in.DrainNatIps {
		out.DrainNatIps = append(out.DrainNatIps, krm.ResourceRef{External: ip})
	}
	out.EnableDynamicPortAllocation = in.EnableDynamicPortAllocation
	out.EnableEndpointIndependentMapping = in.EnableEndpointIndependentMapping
	if in.IcmpIdleTimeoutSec != nil {
		val := int64(*in.IcmpIdleTimeoutSec)
		out.IcmpIdleTimeoutSec = &val
	}
	out.LogConfig = RouternatLogConfig_v1beta1_FromProto(mapCtx, in.GetLogConfig())
	if in.MaxPortsPerVm != nil {
		val := int64(*in.MaxPortsPerVm)
		out.MaxPortsPerVm = &val
	}
	if in.MinPortsPerVm != nil {
		val := int64(*in.MinPortsPerVm)
		out.MinPortsPerVm = &val
	}
	if in.NatIpAllocateOption != nil {
		out.NatIpAllocateOption = *in.NatIpAllocateOption
	}
	for _, ip := range in.NatIps {
		out.NatIps = append(out.NatIps, krm.ResourceRef{External: ip})
	}
	out.Rules = direct.Slice_FromProto(mapCtx, in.Rules, RouternatRules_v1beta1_FromProto)
	if in.SourceSubnetworkIpRangesToNat != nil {
		out.SourceSubnetworkIpRangesToNat = *in.SourceSubnetworkIpRangesToNat
	}
	out.Subnetwork = direct.Slice_FromProto(mapCtx, in.Subnetworks, RouternatSubnetwork_v1beta1_FromProto)
	if in.TcpEstablishedIdleTimeoutSec != nil {
		val := int64(*in.TcpEstablishedIdleTimeoutSec)
		out.TcpEstablishedIdleTimeoutSec = &val
	}
	if in.TcpTimeWaitTimeoutSec != nil {
		val := int64(*in.TcpTimeWaitTimeoutSec)
		out.TcpTimeWaitTimeoutSec = &val
	}
	if in.TcpTransitoryIdleTimeoutSec != nil {
		val := int64(*in.TcpTransitoryIdleTimeoutSec)
		out.TcpTransitoryIdleTimeoutSec = &val
	}
	if in.UdpIdleTimeoutSec != nil {
		val := int64(*in.UdpIdleTimeoutSec)
		out.UdpIdleTimeoutSec = &val
	}
	return out
}
