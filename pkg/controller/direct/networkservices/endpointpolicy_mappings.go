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

package networkservices

import (
	pb "cloud.google.com/go/networkservices/apiv1/networkservicespb"
	krm "github.com/GoogleCloudPlatform/k8s-config-connector/apis/networkservices/v1beta1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct"
)

func NetworkServicesEndpointPolicySpec_FromProto(ctx *direct.MapContext, in *pb.EndpointPolicy) *krm.NetworkServicesEndpointPolicySpec {
	if in == nil {
		return nil
	}
	out := &krm.NetworkServicesEndpointPolicySpec{}
	out.Labels = in.Labels
	out.Type = direct.Enum_FromProto(ctx, in.GetType())
	out.AuthorizationPolicy = direct.LazyPtr(in.GetAuthorizationPolicy())
	out.EndpointMatcher = EndpointMatcher_FromProto(ctx, in.GetEndpointMatcher())             // Assuming nested mapper exists
	out.TrafficPortSelector = TrafficPortSelector_FromProto(ctx, in.GetTrafficPortSelector()) // Assuming nested mapper exists
	out.Description = direct.LazyPtr(in.GetDescription())
	out.ServerTLSPolicy = direct.LazyPtr(in.GetServerTlsPolicy())
	out.ClientTLSPolicy = direct.LazyPtr(in.GetClientTlsPolicy())
	return out
}
func NetworkServicesEndpointPolicySpec_ToProto(ctx *direct.MapContext, in *krm.NetworkServicesEndpointPolicySpec) *pb.EndpointPolicy {
	if in == nil {
		return nil
	}
	out := &pb.EndpointPolicy{}
	out.Labels = in.Labels
	out.Type = direct.Enum_ToProto[pb.EndpointPolicy_EndpointPolicyType](ctx, in.Type)
	out.AuthorizationPolicy = direct.ValueOf(in.AuthorizationPolicy)
	out.EndpointMatcher = EndpointMatcher_ToProto(ctx, in.EndpointMatcher)             // Assuming nested mapper exists
	out.TrafficPortSelector = TrafficPortSelector_ToProto(ctx, in.TrafficPortSelector) // Assuming nested mapper exists
	out.Description = direct.ValueOf(in.Description)
	out.ServerTlsPolicy = direct.ValueOf(in.ServerTLSPolicy)
	out.ClientTlsPolicy = direct.ValueOf(in.ClientTLSPolicy)
	return out
}
func NetworkServicesEndpointPolicyObservedState_FromProto(mapCtx *direct.MapContext, in *pb.EndpointPolicy) *krm.NetworkServicesEndpointPolicyObservedState {
	if in == nil {
		return nil
	}
	out := &krm.NetworkServicesEndpointPolicyObservedState{}
	out.CreateTime = direct.StringTimestamp_FromProto(mapCtx, in.GetCreateTime())
	out.UpdateTime = direct.StringTimestamp_FromProto(mapCtx, in.GetUpdateTime())
	return out
}
func NetworkServicesEndpointPolicyObservedState_ToProto(mapCtx *direct.MapContext, in *krm.NetworkServicesEndpointPolicyObservedState) *pb.EndpointPolicy {
	if in == nil {
		return nil
	}
	out := &pb.EndpointPolicy{}
	out.CreateTime = direct.StringTimestamp_ToProto(mapCtx, in.CreateTime)
	out.UpdateTime = direct.StringTimestamp_ToProto(mapCtx, in.UpdateTime)
	return out
}
