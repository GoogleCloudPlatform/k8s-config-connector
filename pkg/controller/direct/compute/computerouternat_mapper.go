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

// RouternatLogConfig_v1beta1_FromProto maps a pb.RouterNatLogConfig to a krm.RouternatLogConfig
func RouternatLogConfig_v1beta1_FromProto(mapCtx *direct.MapContext, in *pb.RouterNatLogConfig) *krm.RouternatLogConfig {
	if in == nil {
		return nil
	}
	out := &krm.RouternatLogConfig{}
	out.Enable = direct.ValueOf(in.Enable)
	out.Filter = direct.ValueOf(in.Filter)
	return out
}

// RouternatLogConfig_v1beta1_ToProto maps a krm.RouternatLogConfig to a pb.RouterNatLogConfig
func RouternatLogConfig_v1beta1_ToProto(mapCtx *direct.MapContext, in *krm.RouternatLogConfig) *pb.RouterNatLogConfig {
	if in == nil {
		return nil
	}
	out := &pb.RouterNatLogConfig{}
	out.Enable = &in.Enable
	if in.Filter != "" {
		out.Filter = direct.LazyPtr(in.Filter)
	}
	return out
}

// RouternatAction_v1beta1_FromProto maps a pb.RouterNatRuleAction to a krm.RouternatAction
func RouternatAction_v1beta1_FromProto(mapCtx *direct.MapContext, in *pb.RouterNatRuleAction) *krm.RouternatAction {
	if in == nil {
		return nil
	}
	out := &krm.RouternatAction{}
	for _, ip := range in.SourceNatActiveIps {
		out.SourceNatActiveIpsRefs = append(out.SourceNatActiveIpsRefs, krm.ComputeAddressRef{External: ip})
	}
	for _, ip := range in.SourceNatDrainIps {
		out.SourceNatDrainIpsRefs = append(out.SourceNatDrainIpsRefs, krm.ComputeAddressRef{External: ip})
	}
	return out
}

// RouternatAction_v1beta1_ToProto maps a krm.RouternatAction to a pb.RouterNatRuleAction
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

// RouternatRules_v1beta1_FromProto maps a pb.RouterNatRule to a krm.RouternatRules
func RouternatRules_v1beta1_FromProto(mapCtx *direct.MapContext, in *pb.RouterNatRule) *krm.RouternatRules {
	if in == nil {
		return nil
	}
	out := &krm.RouternatRules{}
	out.Action = RouternatAction_v1beta1_FromProto(mapCtx, in.Action)
	out.Description = in.Description
	out.Match = direct.ValueOf(in.Match)
	if in.RuleNumber != nil {
		out.RuleNumber = int64(*in.RuleNumber)
	}
	return out
}

// RouternatRules_v1beta1_ToProto maps a krm.RouternatRules to a pb.RouterNatRule
func RouternatRules_v1beta1_ToProto(mapCtx *direct.MapContext, in *krm.RouternatRules) *pb.RouterNatRule {
	if in == nil {
		return nil
	}
	out := &pb.RouterNatRule{}
	out.Action = RouternatAction_v1beta1_ToProto(mapCtx, in.Action)
	out.Description = in.Description
	if in.Match != "" {
		out.Match = direct.LazyPtr(in.Match)
	}
	if in.RuleNumber != 0 {
		val32 := uint32(in.RuleNumber)
		out.RuleNumber = &val32
	}
	return out
}

// RouternatSubnetwork_v1beta1_FromProto maps a pb.RouterNatSubnetworkToNat to a krm.RouternatSubnetwork
func RouternatSubnetwork_v1beta1_FromProto(mapCtx *direct.MapContext, in *pb.RouterNatSubnetworkToNat) *krm.RouternatSubnetwork {
	if in == nil {
		return nil
	}
	out := &krm.RouternatSubnetwork{}
	out.SecondaryIpRangeNames = in.SecondaryIpRangeNames
	out.SourceIpRangesToNat = in.SourceIpRangesToNat
	if in.GetName() != "" {
		out.SubnetworkRef = krm.ComputeSubnetworkRef{External: in.GetName()}
	}
	return out
}

// RouternatSubnetwork_v1beta1_ToProto maps a krm.RouternatSubnetwork to a pb.RouterNatSubnetworkToNat
func RouternatSubnetwork_v1beta1_ToProto(mapCtx *direct.MapContext, in *krm.RouternatSubnetwork) *pb.RouterNatSubnetworkToNat {
	if in == nil {
		return nil
	}
	out := &pb.RouterNatSubnetworkToNat{}
	out.SecondaryIpRangeNames = in.SecondaryIpRangeNames
	out.SourceIpRangesToNat = in.SourceIpRangesToNat
	if in.SubnetworkRef.External != "" {
		out.Name = direct.LazyPtr(in.SubnetworkRef.External)
	}
	return out
}

// ComputeRouterNATSpec_v1beta1_FromProto maps a pb.RouterNat to a krm.ComputeRouterNATSpec
func ComputeRouterNATSpec_v1beta1_FromProto(mapCtx *direct.MapContext, in *pb.RouterNat) *krm.ComputeRouterNATSpec {
	if in == nil {
		return nil
	}
	out := &krm.ComputeRouterNATSpec{}
	for _, ip := range in.DrainNatIps {
		out.DrainNatIps = append(out.DrainNatIps, krm.ComputeAddressRef{External: ip})
	}
	out.EnableDynamicPortAllocation = in.EnableDynamicPortAllocation
	out.EnableEndpointIndependentMapping = in.EnableEndpointIndependentMapping
	if in.IcmpIdleTimeoutSec != nil {
		val := int64(*in.IcmpIdleTimeoutSec)
		out.IcmpIdleTimeoutSec = &val
	}
	out.LogConfig = RouternatLogConfig_v1beta1_FromProto(mapCtx, in.LogConfig)
	if in.MaxPortsPerVm != nil {
		val := int64(*in.MaxPortsPerVm)
		out.MaxPortsPerVm = &val
	}
	if in.MinPortsPerVm != nil {
		val := int64(*in.MinPortsPerVm)
		out.MinPortsPerVm = &val
	}
	out.NatIpAllocateOption = direct.ValueOf(in.NatIpAllocateOption)
	for _, ip := range in.NatIps {
		out.NatIps = append(out.NatIps, krm.ComputeAddressRef{External: ip})
	}
	out.Rules = direct.Slice_FromProto(mapCtx, in.Rules, RouternatRules_v1beta1_FromProto)
	out.SourceSubnetworkIpRangesToNat = direct.ValueOf(in.SourceSubnetworkIpRangesToNat)
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

// ComputeRouterNATSpec_v1beta1_ToProto maps a krm.ComputeRouterNATSpec to a pb.RouterNat
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
		val32 := int32(*in.IcmpIdleTimeoutSec)
		out.IcmpIdleTimeoutSec = &val32
	}
	out.LogConfig = RouternatLogConfig_v1beta1_ToProto(mapCtx, in.LogConfig)
	if in.MaxPortsPerVm != nil {
		val32 := int32(*in.MaxPortsPerVm)
		out.MaxPortsPerVm = &val32
	}
	if in.MinPortsPerVm != nil {
		val32 := int32(*in.MinPortsPerVm)
		out.MinPortsPerVm = &val32
	}
	if in.NatIpAllocateOption != "" {
		out.NatIpAllocateOption = direct.LazyPtr(in.NatIpAllocateOption)
	}
	for _, ref := range in.NatIps {
		if ref.External != "" {
			out.NatIps = append(out.NatIps, ref.External)
		}
	}
	out.Rules = direct.Slice_ToProto(mapCtx, in.Rules, RouternatRules_v1beta1_ToProto)
	if in.SourceSubnetworkIpRangesToNat != "" {
		out.SourceSubnetworkIpRangesToNat = direct.LazyPtr(in.SourceSubnetworkIpRangesToNat)
	}
	out.Subnetworks = direct.Slice_ToProto(mapCtx, in.Subnetwork, RouternatSubnetwork_v1beta1_ToProto)
	if in.TcpEstablishedIdleTimeoutSec != nil {
		val32 := int32(*in.TcpEstablishedIdleTimeoutSec)
		out.TcpEstablishedIdleTimeoutSec = &val32
	}
	if in.TcpTimeWaitTimeoutSec != nil {
		val32 := int32(*in.TcpTimeWaitTimeoutSec)
		out.TcpTimeWaitTimeoutSec = &val32
	}
	if in.TcpTransitoryIdleTimeoutSec != nil {
		val32 := int32(*in.TcpTransitoryIdleTimeoutSec)
		out.TcpTransitoryIdleTimeoutSec = &val32
	}
	if in.UdpIdleTimeoutSec != nil {
		val32 := int32(*in.UdpIdleTimeoutSec)
		out.UdpIdleTimeoutSec = &val32
	}
	return out
}
