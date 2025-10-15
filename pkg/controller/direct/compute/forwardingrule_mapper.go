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
	"strconv"
	"strings"

	pb "cloud.google.com/go/compute/apiv1/computepb"
	krm "github.com/GoogleCloudPlatform/k8s-config-connector/apis/compute/v1beta1"
	refs "github.com/GoogleCloudPlatform/k8s-config-connector/apis/refs/v1beta1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct"
)

func ComputeForwardingRuleSpec_FromProto(mapCtx *direct.MapContext, in *pb.ForwardingRule) *krm.ComputeForwardingRuleSpec {
	if in == nil {
		return nil
	}
	out := &krm.ComputeForwardingRuleSpec{}
	out.IpAddress = ComputeForwardingRuleSpec_IpAddress_FromProto(mapCtx, in.GetIPAddress())
	out.IpProtocol = in.IPProtocol
	out.AllPorts = in.AllPorts
	out.AllowGlobalAccess = in.AllowGlobalAccess
	out.AllowPscGlobalAccess = in.AllowPscGlobalAccess
	out.BackendServiceRef = ComputeForwardingRuleSpec_BackendSeriviceRef_FromProto(mapCtx, in.GetBackendService())
	// MISSING: BaseForwardingRule
	// MISSING: CreationTimestamp
	out.Description = in.Description
	// MISSING: Fingerprint
	// MISSING: Id
	// MISSING: IpCollection
	out.IpVersion = in.IpVersion
	out.IsMirroringCollector = in.IsMirroringCollector
	// MISSING: Kind
	// MISSING: LabelFingerprint
	// MISSING: Labels
	out.LoadBalancingScheme = in.LoadBalancingScheme
	out.MetadataFilters = direct.Slice_FromProto(mapCtx, in.MetadataFilters, MetadataFilter_FromProto)
	// MISSING: Name
	out.NetworkRef = ComputeForwardingRuleSpec_NetworkRef_FromProto(mapCtx, in.GetNetwork())
	out.NetworkTier = in.NetworkTier
	out.NoAutomateDnsZone = in.NoAutomateDnsZone
	out.PortRange = in.PortRange
	out.Ports = in.Ports
	// MISSING: PscConnectionId
	// MISSING: PscConnectionStatus
	// MISSING: Region
	// MISSING: SelfLink
	out.ServiceDirectoryRegistrations = direct.Slice_FromProto(mapCtx, in.ServiceDirectoryRegistrations, ForwardingruleServiceDirectoryRegistrations_FromProto)
	out.ServiceLabel = in.ServiceLabel
	// MISSING: ServiceName
	out.SourceIpRanges = in.SourceIpRanges
	out.SubnetworkRef = ComputeForwardingRuleSpec_SubnetworkRef_FromProto(mapCtx, in.GetSubnetwork())
	out.Target = ComputeForwardingRuleSpec_Target_FromProto(mapCtx, in.GetTarget())
	return out
}
func ComputeForwardingRuleSpec_ToProto(mapCtx *direct.MapContext, in *krm.ComputeForwardingRuleSpec) *pb.ForwardingRule {
	if in == nil {
		return nil
	}
	out := &pb.ForwardingRule{}
	out.IPAddress = ComputeForwardingRuleSpec_IpAddress_ToProto(mapCtx, in.IpAddress)
	out.IPProtocol = in.IpProtocol
	out.AllPorts = in.AllPorts
	out.AllowGlobalAccess = in.AllowGlobalAccess
	out.AllowPscGlobalAccess = in.AllowPscGlobalAccess
	out.BackendService = ComputeForwardingRuleSpec_BackendSeriviceRef_ToProto(mapCtx, in.BackendServiceRef)
	// MISSING: BaseForwardingRule
	// MISSING: CreationTimestamp
	out.Description = in.Description
	// MISSING: Fingerprint
	// MISSING: Id
	// MISSING: IpCollection
	out.IpVersion = in.IpVersion
	out.IsMirroringCollector = in.IsMirroringCollector
	// MISSING: Kind
	// MISSING: LabelFingerprint
	// MISSING: Labels
	out.LoadBalancingScheme = in.LoadBalancingScheme
	out.MetadataFilters = direct.Slice_ToProto(mapCtx, in.MetadataFilters, MetadataFilter_ToProto)
	// MISSING: Name
	out.Network = ComputeForwardingRuleSpec_NetworkRef_ToProto(mapCtx, in.NetworkRef)
	out.NetworkTier = in.NetworkTier
	out.NoAutomateDnsZone = in.NoAutomateDnsZone
	out.PortRange = in.PortRange
	out.Ports = in.Ports
	// MISSING: PscConnectionId
	// MISSING: PscConnectionStatus
	// MISSING: Region
	// MISSING: SelfLink
	out.ServiceDirectoryRegistrations = direct.Slice_ToProto(mapCtx, in.ServiceDirectoryRegistrations, ForwardingruleServiceDirectoryRegistrations_ToProto)
	out.ServiceLabel = in.ServiceLabel
	// MISSING: ServiceName
	out.SourceIpRanges = in.SourceIpRanges
	out.Subnetwork = ComputeForwardingRuleSpec_SubnetworkRef_ToProto(mapCtx, in.SubnetworkRef)
	out.Target = ComputeForwardingRuleSpec_Target_ToProto(mapCtx, in.Target)
	return out
}
func ComputeForwardingRuleStatus_FromProto(mapCtx *direct.MapContext, in *pb.ForwardingRule) *krm.ComputeForwardingRuleStatus {
	if in == nil {
		return nil
	}
	out := &krm.ComputeForwardingRuleStatus{}
	out.BaseForwardingRule = in.BaseForwardingRule
	out.CreationTimestamp = in.CreationTimestamp
	out.LabelFingerprint = in.LabelFingerprint
	out.PscConnectionId = ComputeForwardingRuleStatus_PscConnectionId_FromProto(mapCtx, in.GetPscConnectionId())
	out.PscConnectionStatus = in.PscConnectionStatus
	out.SelfLink = in.SelfLink
	out.ServiceName = in.ServiceName
	return out
}
func ComputeForwardingRuleStatus_ToProto(mapCtx *direct.MapContext, in *krm.ComputeForwardingRuleStatus) *pb.ForwardingRule {
	if in == nil {
		return nil
	}
	out := &pb.ForwardingRule{}
	out.BaseForwardingRule = in.BaseForwardingRule
	out.CreationTimestamp = in.CreationTimestamp
	out.LabelFingerprint = in.LabelFingerprint
	out.PscConnectionId = ComputeForwardingRuleStatus_PscConnectionId_ToProto(mapCtx, in.PscConnectionId)
	out.PscConnectionStatus = in.PscConnectionStatus
	out.SelfLink = in.SelfLink
	out.ServiceName = in.ServiceName
	return out
}

func ComputeForwardingRuleSpec_IpAddress_ToProto(mapCtx *direct.MapContext, in *krm.IpAddress) *string {
	if in == nil {
		return nil
	}

	var out *string
	if oneof := in.AddressRef; oneof != nil {
		if oneof.External == "" {
			mapCtx.Errorf("reference %s was not pre-resolved", oneof.Name)
		}
		out = direct.LazyPtr(oneof.External)
	}
	if oneof := in.Ip; oneof != nil {
		out = in.Ip
	}
	return out
}

func ComputeForwardingRuleSpec_IpAddress_FromProto(mapCtx *direct.MapContext, in string) *krm.IpAddress {
	if in == "" {
		return nil
	}
	out := &krm.IpAddress{}
	out.Ip = direct.LazyPtr(in)
	return out
}

func ComputeForwardingRuleSpec_BackendSeriviceRef_FromProto(mapCtx *direct.MapContext, in string) *krm.ComputeBackendServiceRef {
	if in == "" {
		return nil
	}
	return &krm.ComputeBackendServiceRef{
		External: in,
	}
}

func ComputeForwardingRuleSpec_BackendSeriviceRef_ToProto(mapCtx *direct.MapContext, in *krm.ComputeBackendServiceRef) *string {
	if in == nil {
		return nil
	}
	if in.External == "" {
		mapCtx.Errorf("reference %s was not pre-resolved", in.Name)
	}
	return direct.LazyPtr(in.External)
}

func ComputeForwardingRuleSpec_NetworkRef_FromProto(mapCtx *direct.MapContext, in string) *refs.ComputeNetworkRef {
	if in == "" {
		return nil
	}
	return &refs.ComputeNetworkRef{
		External: in,
	}
}

func ComputeForwardingRuleSpec_NetworkRef_ToProto(mapCtx *direct.MapContext, in *refs.ComputeNetworkRef) *string {
	if in == nil {
		return nil
	}
	if in.External == "" {
		mapCtx.Errorf("reference %s was not pre-resolved", in.Name)
	}
	return direct.LazyPtr(in.External)
}

func ComputeForwardingRuleSpec_SubnetworkRef_FromProto(mapCtx *direct.MapContext, in string) *refs.ComputeSubnetworkRef {
	if in == "" {
		return nil
	}
	return &refs.ComputeSubnetworkRef{
		External: in,
	}
}

func ComputeForwardingRuleSpec_SubnetworkRef_ToProto(mapCtx *direct.MapContext, in *refs.ComputeSubnetworkRef) *string {
	if in == nil {
		return nil
	}
	if in.External == "" {
		mapCtx.Errorf("reference %s was not pre-resolved", in.Name)
	}
	return direct.LazyPtr(in.External)
}

func ComputeForwardingRuleSpec_Target_ToProto(mapCtx *direct.MapContext, in *krm.Target) *string {
	if in == nil {
		return nil
	}

	var out *string
	if oneof := in.GoogleAPIsBundle; oneof != nil {
		out = in.GoogleAPIsBundle
	}
	if oneof := in.ServiceAttachmentRef; oneof != nil {
		if oneof.External == "" {
			mapCtx.Errorf("reference %s was not pre-resolved", oneof.Name)
		}
		out = direct.LazyPtr(oneof.External)
	}
	if oneof := in.TargetGRPCProxyRef; oneof != nil {
		if oneof.External == "" {
			mapCtx.Errorf("reference %s was not pre-resolved", oneof.Name)
		}
		out = direct.LazyPtr(oneof.External)
	}
	if oneof := in.TargetHTTPSProxyRef; oneof != nil {
		if oneof.External == "" {
			mapCtx.Errorf("reference %s was not pre-resolved", oneof.Name)
		}
		out = direct.LazyPtr(oneof.External)
	}
	if oneof := in.TargetHTTPProxyRef; oneof != nil {
		if oneof.External == "" {
			mapCtx.Errorf("reference %s was not pre-resolved", oneof.Name)
		}
		out = direct.LazyPtr(oneof.External)
	}
	if oneof := in.TargetSSLProxyRef; oneof != nil {
		if oneof.External == "" {
			mapCtx.Errorf("reference %s was not pre-resolved", oneof.Name)
		}
		out = direct.LazyPtr(oneof.External)
	}
	if oneof := in.TargetTCPProxyRef; oneof != nil {
		if oneof.External == "" {
			mapCtx.Errorf("reference %s was not pre-resolved", oneof.Name)
		}
		out = direct.LazyPtr(oneof.External)
	}
	if oneof := in.TargetVPNGatewayRef; oneof != nil {
		if oneof.External == "" {
			mapCtx.Errorf("reference %s was not pre-resolved", oneof.Name)
		}
		out = direct.LazyPtr(oneof.External)
	}
	return out
}

func ComputeForwardingRuleSpec_Target_FromProto(mapCtx *direct.MapContext, in string) *krm.Target {
	if in == "" {
		return nil
	}
	out := &krm.Target{}
	if strings.Contains(in, "serviceAttachments") {
		out.ServiceAttachmentRef = &refs.ComputeServiceAttachmentRef{
			External: in,
		}
	} else if strings.Contains(in, "targetGrpcProxies") {
		out.TargetGRPCProxyRef = &refs.ComputeTargetGrpcProxyRef{
			External: in,
		}
	} else if strings.Contains(in, "targetHttpProxies") {
		out.TargetHTTPProxyRef = &refs.ComputeTargetHTTPProxyRef{
			External: in,
		}
	} else if strings.Contains(in, "targetHttpsProxies") {
		out.TargetHTTPSProxyRef = &refs.ComputeTargetHTTPSProxyRef{
			External: in,
		}
	} else if strings.Contains(in, "targetSslProxies") {
		out.TargetSSLProxyRef = &refs.ComputeTargetSSLProxyRef{
			External: in,
		}
	} else if strings.Contains(in, "targetTcpProxies") {
		out.TargetTCPProxyRef = &refs.ComputeTargetTCPProxyRef{
			External: in,
		}
	} else if strings.Contains(in, "targetVpnGateways") {
		out.TargetVPNGatewayRef = &refs.ComputeTargetVPNGatewayRef{
			External: in,
		}
	}
	return out
}

func ComputeForwardingRuleStatus_PscConnectionId_FromProto(mapCtx *direct.MapContext, in uint64) *string {
	if in == 0 {
		return nil
	}
	strValue := strconv.FormatUint(in, 10)
	return &strValue

}

func ComputeForwardingRuleStatus_PscConnectionId_ToProto(mapCtx *direct.MapContext, in *string) *uint64 {
	if in == nil {
		return nil
	}

	num, err := strconv.ParseUint(*in, 10, 64)
	if err != nil {
		mapCtx.Errorf("Error converting string %s to uint64", direct.ValueOf(in))
		return nil
	}

	return &num
}
