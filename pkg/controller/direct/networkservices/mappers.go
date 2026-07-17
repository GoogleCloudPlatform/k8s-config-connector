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

package networkservices

import (
	pb "cloud.google.com/go/networkservices/apiv1/networkservicespb"
	networksecurityv1beta1 "github.com/GoogleCloudPlatform/k8s-config-connector/apis/networksecurity/v1beta1"
	krm "github.com/GoogleCloudPlatform/k8s-config-connector/apis/networkservices/v1beta1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct"
)

// NetworkServicesGatewaySpec_FromProto converts a Proto Gateway message to a KRM NetworkServicesGatewaySpec.
// Handcoded because of type mismatch in Ports field ([]int32 vs []int64) and custom reference types.
func NetworkServicesGatewaySpec_FromProto(mapCtx *direct.MapContext, in *pb.Gateway) *krm.NetworkServicesGatewaySpec {
	if in == nil {
		return nil
	}
	out := &krm.NetworkServicesGatewaySpec{}
	out.Addresses = in.Addresses
	out.Description = direct.LazyPtr(in.GetDescription())
	if in.Ports != nil {
		out.Ports = make([]int64, len(in.Ports))
		for i, p := range in.Ports {
			out.Ports[i] = int64(p)
		}
	}
	out.Scope = in.GetScope()
	if in.GetServerTlsPolicy() != "" {
		out.ServerTlsPolicyRef = &networksecurityv1beta1.NetworkSecurityServerTLSPolicyRef{
			External: in.GetServerTlsPolicy(),
		}
	}
	out.Type = direct.Enum_FromProto(mapCtx, in.GetType())
	return out
}

// NetworkServicesGatewaySpec_ToProto converts a KRM NetworkServicesGatewaySpec to a Proto Gateway message.
// Handcoded because of type mismatch in Ports field ([]int64 vs []int32) and custom reference types.
func NetworkServicesGatewaySpec_ToProto(mapCtx *direct.MapContext, in *krm.NetworkServicesGatewaySpec) *pb.Gateway {
	if in == nil {
		return nil
	}
	out := &pb.Gateway{}
	out.Addresses = in.Addresses
	out.Description = direct.ValueOf(in.Description)
	if in.Ports != nil {
		out.Ports = make([]int32, len(in.Ports))
		for i, p := range in.Ports {
			out.Ports[i] = int32(p)
		}
	}
	out.Scope = in.Scope
	if in.ServerTlsPolicyRef != nil {
		out.ServerTlsPolicy = in.ServerTlsPolicyRef.External
	}
	out.Type = direct.Enum_ToProto[pb.Gateway_Type](mapCtx, in.Type)
	return out
}

// NetworkServicesGatewayStatus_FromProto converts a Proto Gateway message to a KRM NetworkServicesGatewayStatus.
func NetworkServicesGatewayStatus_FromProto(mapCtx *direct.MapContext, in *pb.Gateway) *krm.NetworkServicesGatewayStatus {
	if in == nil {
		return nil
	}
	out := &krm.NetworkServicesGatewayStatus{}
	out.SelfLink = direct.LazyPtr(in.GetSelfLink())
	out.CreateTime = direct.StringTimestamp_FromProto(mapCtx, in.GetCreateTime())
	out.UpdateTime = direct.StringTimestamp_FromProto(mapCtx, in.GetUpdateTime())
	return out
}

// NetworkServicesGatewayStatus_ToProto converts a KRM NetworkServicesGatewayStatus to a Proto Gateway message.
func NetworkServicesGatewayStatus_ToProto(mapCtx *direct.MapContext, in *krm.NetworkServicesGatewayStatus) *pb.Gateway {
	if in == nil {
		return nil
	}
	out := &pb.Gateway{}
	out.SelfLink = direct.ValueOf(in.SelfLink)
	out.CreateTime = direct.StringTimestamp_ToProto(mapCtx, in.CreateTime)
	out.UpdateTime = direct.StringTimestamp_ToProto(mapCtx, in.UpdateTime)
	return out
}
