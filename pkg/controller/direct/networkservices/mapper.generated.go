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
	krm "github.com/GoogleCloudPlatform/k8s-config-connector/apis/networkservices/v1alpha1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct"
	refs "github.com/GoogleCloudPlatform/k8s-config-connector/apis/refs/v1beta1"
)
func Gateway_FromProto(mapCtx *direct.MapContext, in *pb.Gateway) *krm.Gateway {
	if in == nil {
		return nil
	}
	out := &krm.Gateway{}
	out.Name = direct.LazyPtr(in.GetName())
	// MISSING: SelfLink
	// MISSING: CreateTime
	// MISSING: UpdateTime
	out.Labels = in.Labels
	out.Description = direct.LazyPtr(in.GetDescription())
	out.Type = direct.Enum_FromProto(mapCtx, in.GetType())
	out.Ports = in.Ports
	out.Scope = direct.LazyPtr(in.GetScope())
	out.ServerTlsPolicy = direct.LazyPtr(in.GetServerTlsPolicy())
	return out
}
func Gateway_ToProto(mapCtx *direct.MapContext, in *krm.Gateway) *pb.Gateway {
	if in == nil {
		return nil
	}
	out := &pb.Gateway{}
	out.Name = direct.ValueOf(in.Name)
	// MISSING: SelfLink
	// MISSING: CreateTime
	// MISSING: UpdateTime
	out.Labels = in.Labels
	out.Description = direct.ValueOf(in.Description)
	out.Type = direct.Enum_ToProto[pb.Gateway_Type](mapCtx, in.Type)
	out.Ports = in.Ports
	out.Scope = direct.ValueOf(in.Scope)
	out.ServerTlsPolicy = direct.ValueOf(in.ServerTlsPolicy)
	return out
}
func GatewayObservedState_FromProto(mapCtx *direct.MapContext, in *pb.Gateway) *krm.GatewayObservedState {
	if in == nil {
		return nil
	}
	out := &krm.GatewayObservedState{}
	// MISSING: Name
	out.SelfLink = direct.LazyPtr(in.GetSelfLink())
	out.CreateTime = direct.StringTimestamp_FromProto(mapCtx, in.GetCreateTime())
	out.UpdateTime = direct.StringTimestamp_FromProto(mapCtx, in.GetUpdateTime())
	// MISSING: Labels
	// MISSING: Description
	// MISSING: Type
	// MISSING: Ports
	// MISSING: Scope
	// MISSING: ServerTlsPolicy
	return out
}
func GatewayObservedState_ToProto(mapCtx *direct.MapContext, in *krm.GatewayObservedState) *pb.Gateway {
	if in == nil {
		return nil
	}
	out := &pb.Gateway{}
	// MISSING: Name
	out.SelfLink = direct.ValueOf(in.SelfLink)
	out.CreateTime = direct.StringTimestamp_ToProto(mapCtx, in.CreateTime)
	out.UpdateTime = direct.StringTimestamp_ToProto(mapCtx, in.UpdateTime)
	// MISSING: Labels
	// MISSING: Description
	// MISSING: Type
	// MISSING: Ports
	// MISSING: Scope
	// MISSING: ServerTlsPolicy
	return out
}
func NetworkservicesGatewayObservedState_FromProto(mapCtx *direct.MapContext, in *pb.Gateway) *krm.NetworkservicesGatewayObservedState {
	if in == nil {
		return nil
	}
	out := &krm.NetworkservicesGatewayObservedState{}
	// MISSING: Name
	// MISSING: SelfLink
	// MISSING: CreateTime
	// MISSING: UpdateTime
	// MISSING: Labels
	// MISSING: Description
	// MISSING: Type
	// MISSING: Ports
	// MISSING: Scope
	// MISSING: ServerTlsPolicy
	return out
}
func NetworkservicesGatewayObservedState_ToProto(mapCtx *direct.MapContext, in *krm.NetworkservicesGatewayObservedState) *pb.Gateway {
	if in == nil {
		return nil
	}
	out := &pb.Gateway{}
	// MISSING: Name
	// MISSING: SelfLink
	// MISSING: CreateTime
	// MISSING: UpdateTime
	// MISSING: Labels
	// MISSING: Description
	// MISSING: Type
	// MISSING: Ports
	// MISSING: Scope
	// MISSING: ServerTlsPolicy
	return out
}
func NetworkservicesGatewaySpec_FromProto(mapCtx *direct.MapContext, in *pb.Gateway) *krm.NetworkservicesGatewaySpec {
	if in == nil {
		return nil
	}
	out := &krm.NetworkservicesGatewaySpec{}
	// MISSING: Name
	// MISSING: SelfLink
	// MISSING: CreateTime
	// MISSING: UpdateTime
	// MISSING: Labels
	// MISSING: Description
	// MISSING: Type
	// MISSING: Ports
	// MISSING: Scope
	// MISSING: ServerTlsPolicy
	return out
}
func NetworkservicesGatewaySpec_ToProto(mapCtx *direct.MapContext, in *krm.NetworkservicesGatewaySpec) *pb.Gateway {
	if in == nil {
		return nil
	}
	out := &pb.Gateway{}
	// MISSING: Name
	// MISSING: SelfLink
	// MISSING: CreateTime
	// MISSING: UpdateTime
	// MISSING: Labels
	// MISSING: Description
	// MISSING: Type
	// MISSING: Ports
	// MISSING: Scope
	// MISSING: ServerTlsPolicy
	return out
}
