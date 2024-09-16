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

package compute

import (
	"strconv"
	"strings"

	computev1beta1 "github.com/GoogleCloudPlatform/k8s-config-connector/apis/compute/v1beta1"

	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct"

	krm "github.com/GoogleCloudPlatform/k8s-config-connector/apis/compute/v1beta1"
)

func ComputeForwardingRuleSpec_IpAddress_ToProto(mapCtx *direct.MapContext, in *krm.ForwardingruleIpAddress) *string {
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
	if in.Ip != nil {
		out = in.Ip
	}
	return out
}

func ComputeForwardingRuleSpec_IpAddress_FromProto(mapCtx *direct.MapContext, in string) *krm.ForwardingruleIpAddress {
	if in == "" {
		return nil
	}
	out := &krm.ForwardingruleIpAddress{}
	out.Ip = direct.LazyPtr(in)
	return out
}

func ComputeForwardingRuleSpec_BackendSeriviceRef_FromProto(mapCtx *direct.MapContext, in string) *computev1beta1.ComputeBackendServiceRef {
	if in == "" {
		return nil
	}
	return &computev1beta1.ComputeBackendServiceRef{
		External: in,
	}
}

func ComputeForwardingRuleSpec_BackendSeriviceRef_ToProto(mapCtx *direct.MapContext, in *computev1beta1.ComputeBackendServiceRef) *string {
	if in == nil {
		return nil
	}
	if in.External == "" {
		mapCtx.Errorf("reference %s was not pre-resolved", in.Name)
	}
	return direct.LazyPtr(in.External)
}

func ComputeForwardingRuleSpec_NetworkRef_FromProto(mapCtx *direct.MapContext, in string) *computev1beta1.ComputeNetworkRef {
	if in == "" {
		return nil
	}
	return &computev1beta1.ComputeNetworkRef{
		External: in,
	}
}

func ComputeForwardingRuleSpec_NetworkRef_ToProto(mapCtx *direct.MapContext, in *computev1beta1.ComputeNetworkRef) *string {
	if in == nil {
		return nil
	}
	if in.External == "" {
		mapCtx.Errorf("reference %s was not pre-resolved", in.Name)
	}
	return direct.LazyPtr(in.External)
}

func ComputeForwardingRuleSpec_SubnetworkRef_FromProto(mapCtx *direct.MapContext, in string) *computev1beta1.ComputeSubnetworkRef {
	if in == "" {
		return nil
	}
	return &computev1beta1.ComputeSubnetworkRef{
		External: in,
	}
}

func ComputeForwardingRuleSpec_SubnetworkRef_ToProto(mapCtx *direct.MapContext, in *computev1beta1.ComputeSubnetworkRef) *string {
	if in == nil {
		return nil
	}
	if in.External == "" {
		mapCtx.Errorf("reference %s was not pre-resolved", in.Name)
	}
	return direct.LazyPtr(in.External)
}

func ComputeForwardingRuleSpec_Target_ToProto(mapCtx *direct.MapContext, in *krm.ForwardingruleTarget) *string {
	if in == nil {
		return nil
	}

	var out *string
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

func ComputeForwardingRuleSpec_Target_FromProto(mapCtx *direct.MapContext, in string) *krm.ForwardingruleTarget {
	if in == "" {
		return nil
	}
	out := &krm.ForwardingruleTarget{}
	if strings.Contains(in, "serviceAttachments") {
		out.ServiceAttachmentRef = &computev1beta1.ComputeServiceAttachmentRef{
			External: in,
		}
	} else if strings.Contains(in, "targetGrpcProxies") {
		out.TargetGRPCProxyRef = &computev1beta1.ComputeTargetGrpcProxyRef{
			External: in,
		}
	} else if strings.Contains(in, "targetHttpProxies") {
		out.TargetHTTPProxyRef = &computev1beta1.ComputeTargetHTTPProxyRef{
			External: in,
		}
	} else if strings.Contains(in, "targetHttpsProxies") {
		out.TargetHTTPSProxyRef = &computev1beta1.ComputeTargetHTTPSProxyRef{
			External: in,
		}
	} else if strings.Contains(in, "targetSslProxies") {
		out.TargetSSLProxyRef = &computev1beta1.ComputeTargetSSLProxyRef{
			External: in,
		}
	} else if strings.Contains(in, "targetTcpProxies") {
		out.TargetTCPProxyRef = &computev1beta1.ComputeTargetTCPProxyRef{
			External: in,
		}
	} else if strings.Contains(in, "targetVpnGateways") {
		out.TargetVPNGatewayRef = &computev1beta1.ComputeTargetVPNGatewayRef{
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
