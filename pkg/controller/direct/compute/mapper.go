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

	krm "github.com/GoogleCloudPlatform/k8s-config-connector/apis/compute/v1beta1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/clients/generated/apis/k8s/v1alpha1"
)

func ComputeForwardingRuleSpec_IpAddress_ToProto(mapCtx *MapContext, in *krm.ForwardingruleIpAddress) *string {
	if in == nil {
		return nil
	}

	var out *string
	if oneof := ResourceRef_ToProto(mapCtx, in.AddressRef); oneof != nil {
		out = oneof
	}
	if in.Ip != nil {
		out = in.Ip
	}
	return out
}

func ComputeForwardingRuleSpec_IpAddress_FromProto(mapCtx *MapContext, in string) *krm.ForwardingruleIpAddress {
	if in == "" {
		return nil
	}
	out := &krm.ForwardingruleIpAddress{}
	out.AddressRef = &v1alpha1.ResourceRef{
		External: in,
	}
	return out
}

func ComputeForwardingRuleSpec_Target_ToProto(mapCtx *MapContext, in *krm.ForwardingruleTarget) *string {
	if in == nil {
		return nil
	}

	var out *string
	if oneof := ResourceRef_ToProto(mapCtx, in.ServiceAttachmentRef); oneof != nil {
		out = oneof
	}
	if oneof := ResourceRef_ToProto(mapCtx, in.TargetGRPCProxyRef); oneof != nil {
		out = oneof
	}
	if oneof := ResourceRef_ToProto(mapCtx, in.TargetHTTPProxyRef); oneof != nil {
		out = oneof
	}
	if oneof := ResourceRef_ToProto(mapCtx, in.TargetHTTPSProxyRef); oneof != nil {
		out = oneof
	}
	if oneof := ResourceRef_ToProto(mapCtx, in.TargetSSLProxyRef); oneof != nil {
		out = oneof
	}
	if oneof := ResourceRef_ToProto(mapCtx, in.TargetTCPProxyRef); oneof != nil {
		out = oneof
	}
	if oneof := ResourceRef_ToProto(mapCtx, in.TargetVPNGatewayRef); oneof != nil {
		out = oneof
	}
	return out
}

func ComputeForwardingRuleSpec_Target_FromProto(mapCtx *MapContext, in string) *krm.ForwardingruleTarget {
	if in == "" {
		return nil
	}
	out := &krm.ForwardingruleTarget{}
	// TODO(yuhou): ForwardingRuleTarget can be one of multiple target objects. We need to determine which one to assign the value to.
	// Assign to TargetHTTPProxy temporarily
	out.TargetHTTPProxyRef = &v1alpha1.ResourceRef{
		External: in,
	}
	return out
}

func ComputeForwardingRuleStatus_PscConnectionId_FromProto(mapCtx *MapContext, in uint64) *string {
	if in == 0 {
		return nil
	}
	strValue := strconv.FormatUint(in, 10)
	return &strValue

}

func ComputeForwardingRuleStatus_PscConnectionId_ToProto(mapCtx *MapContext, in *string) *uint64 {
	if in == nil {
		return nil
	}

	num, err := strconv.ParseUint(*in, 10, 64)
	if err != nil {
		mapCtx.Errorf("Error converting string to uint64")
		return nil
	}

	return &num
}
