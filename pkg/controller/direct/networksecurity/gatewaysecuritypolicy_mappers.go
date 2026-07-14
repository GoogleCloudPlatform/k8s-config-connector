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

package networksecurity

import (
	pb "cloud.google.com/go/networksecurity/apiv1/networksecuritypb"
	krm "github.com/GoogleCloudPlatform/k8s-config-connector/apis/networksecurity/v1alpha1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct"
)

func NetworkSecurityGatewaySecurityPolicySpec_FromProto(mapCtx *direct.MapContext, in *pb.GatewaySecurityPolicy) *krm.NetworkSecurityGatewaySecurityPolicySpec {
	if in == nil {
		return nil
	}
	out := &krm.NetworkSecurityGatewaySecurityPolicySpec{}
	out.Description = direct.LazyPtr(in.GetDescription())
	if in.GetTlsInspectionPolicy() != "" {
		out.TLSInspectionPolicyRef = &krm.TLSInspectionPolicyRef{
			External: in.GetTlsInspectionPolicy(),
		}
	}
	return out
}

func NetworkSecurityGatewaySecurityPolicySpec_ToProto(mapCtx *direct.MapContext, in *krm.NetworkSecurityGatewaySecurityPolicySpec) *pb.GatewaySecurityPolicy {
	if in == nil {
		return nil
	}
	out := &pb.GatewaySecurityPolicy{}
	out.Description = direct.ValueOf(in.Description)
	if in.TLSInspectionPolicyRef != nil {
		out.TlsInspectionPolicy = in.TLSInspectionPolicyRef.External
	}
	return out
}

func NetworkSecurityGatewaySecurityPolicyObservedState_FromProto(mapCtx *direct.MapContext, in *pb.GatewaySecurityPolicy) *krm.NetworkSecurityGatewaySecurityPolicyObservedState {
	if in == nil {
		return nil
	}
	out := &krm.NetworkSecurityGatewaySecurityPolicyObservedState{}
	out.CreateTime = direct.StringTimestamp_FromProto(mapCtx, in.GetCreateTime())
	out.UpdateTime = direct.StringTimestamp_FromProto(mapCtx, in.GetUpdateTime())
	return out
}

func NetworkSecurityGatewaySecurityPolicyObservedState_ToProto(mapCtx *direct.MapContext, in *krm.NetworkSecurityGatewaySecurityPolicyObservedState) *pb.GatewaySecurityPolicy {
	if in == nil {
		return nil
	}
	out := &pb.GatewaySecurityPolicy{}
	out.CreateTime = direct.StringTimestamp_ToProto(mapCtx, in.CreateTime)
	out.UpdateTime = direct.StringTimestamp_ToProto(mapCtx, in.UpdateTime)
	return out
}
