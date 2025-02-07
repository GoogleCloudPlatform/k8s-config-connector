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

package apigateway

import (
	refs "github.com/GoogleCloudPlatform/k8s-config-connector/apis/refs/v1beta1"
	pb "cloud.google.com/go/apigateway/apiv1/apigatewaypb"
	krm "github.com/GoogleCloudPlatform/k8s-config-connector/apis/apigateway/v1alpha1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct"
)
func ApigatewayApiConfigObservedState_FromProto(mapCtx *direct.MapContext, in *pb.ApiConfig) *krm.ApigatewayApiConfigObservedState {
	if in == nil {
		return nil
	}
	out := &krm.ApigatewayApiConfigObservedState{}
	// MISSING: Name
	// MISSING: CreateTime
	// MISSING: UpdateTime
	// MISSING: Labels
	// MISSING: DisplayName
	// MISSING: GatewayServiceAccount
	// MISSING: ServiceConfigID
	// MISSING: State
	// MISSING: OpenapiDocuments
	// MISSING: GrpcServices
	// MISSING: ManagedServiceConfigs
	return out
}
func ApigatewayApiConfigObservedState_ToProto(mapCtx *direct.MapContext, in *krm.ApigatewayApiConfigObservedState) *pb.ApiConfig {
	if in == nil {
		return nil
	}
	out := &pb.ApiConfig{}
	// MISSING: Name
	// MISSING: CreateTime
	// MISSING: UpdateTime
	// MISSING: Labels
	// MISSING: DisplayName
	// MISSING: GatewayServiceAccount
	// MISSING: ServiceConfigID
	// MISSING: State
	// MISSING: OpenapiDocuments
	// MISSING: GrpcServices
	// MISSING: ManagedServiceConfigs
	return out
}
func ApigatewayApiConfigSpec_FromProto(mapCtx *direct.MapContext, in *pb.ApiConfig) *krm.ApigatewayApiConfigSpec {
	if in == nil {
		return nil
	}
	out := &krm.ApigatewayApiConfigSpec{}
	// MISSING: Name
	// MISSING: CreateTime
	// MISSING: UpdateTime
	// MISSING: Labels
	// MISSING: DisplayName
	// MISSING: GatewayServiceAccount
	// MISSING: ServiceConfigID
	// MISSING: State
	// MISSING: OpenapiDocuments
	// MISSING: GrpcServices
	// MISSING: ManagedServiceConfigs
	return out
}
func ApigatewayApiConfigSpec_ToProto(mapCtx *direct.MapContext, in *krm.ApigatewayApiConfigSpec) *pb.ApiConfig {
	if in == nil {
		return nil
	}
	out := &pb.ApiConfig{}
	// MISSING: Name
	// MISSING: CreateTime
	// MISSING: UpdateTime
	// MISSING: Labels
	// MISSING: DisplayName
	// MISSING: GatewayServiceAccount
	// MISSING: ServiceConfigID
	// MISSING: State
	// MISSING: OpenapiDocuments
	// MISSING: GrpcServices
	// MISSING: ManagedServiceConfigs
	return out
}
func ApigatewayApiObservedState_FromProto(mapCtx *direct.MapContext, in *pb.Api) *krm.ApigatewayApiObservedState {
	if in == nil {
		return nil
	}
	out := &krm.ApigatewayApiObservedState{}
	// MISSING: Name
	// MISSING: CreateTime
	// MISSING: UpdateTime
	// MISSING: Labels
	// MISSING: DisplayName
	// MISSING: ManagedService
	// MISSING: State
	return out
}
func ApigatewayApiObservedState_ToProto(mapCtx *direct.MapContext, in *krm.ApigatewayApiObservedState) *pb.Api {
	if in == nil {
		return nil
	}
	out := &pb.Api{}
	// MISSING: Name
	// MISSING: CreateTime
	// MISSING: UpdateTime
	// MISSING: Labels
	// MISSING: DisplayName
	// MISSING: ManagedService
	// MISSING: State
	return out
}
func ApigatewayApiSpec_FromProto(mapCtx *direct.MapContext, in *pb.Api) *krm.ApigatewayApiSpec {
	if in == nil {
		return nil
	}
	out := &krm.ApigatewayApiSpec{}
	// MISSING: Name
	// MISSING: CreateTime
	// MISSING: UpdateTime
	// MISSING: Labels
	// MISSING: DisplayName
	// MISSING: ManagedService
	// MISSING: State
	return out
}
func ApigatewayApiSpec_ToProto(mapCtx *direct.MapContext, in *krm.ApigatewayApiSpec) *pb.Api {
	if in == nil {
		return nil
	}
	out := &pb.Api{}
	// MISSING: Name
	// MISSING: CreateTime
	// MISSING: UpdateTime
	// MISSING: Labels
	// MISSING: DisplayName
	// MISSING: ManagedService
	// MISSING: State
	return out
}
func ApigatewayGatewayObservedState_FromProto(mapCtx *direct.MapContext, in *pb.Gateway) *krm.ApigatewayGatewayObservedState {
	if in == nil {
		return nil
	}
	out := &krm.ApigatewayGatewayObservedState{}
	// MISSING: Name
	// MISSING: CreateTime
	// MISSING: UpdateTime
	// MISSING: Labels
	// MISSING: DisplayName
	// MISSING: ApiConfig
	// MISSING: State
	// MISSING: DefaultHostname
	return out
}
func ApigatewayGatewayObservedState_ToProto(mapCtx *direct.MapContext, in *krm.ApigatewayGatewayObservedState) *pb.Gateway {
	if in == nil {
		return nil
	}
	out := &pb.Gateway{}
	// MISSING: Name
	// MISSING: CreateTime
	// MISSING: UpdateTime
	// MISSING: Labels
	// MISSING: DisplayName
	// MISSING: ApiConfig
	// MISSING: State
	// MISSING: DefaultHostname
	return out
}
func ApigatewayGatewaySpec_FromProto(mapCtx *direct.MapContext, in *pb.Gateway) *krm.ApigatewayGatewaySpec {
	if in == nil {
		return nil
	}
	out := &krm.ApigatewayGatewaySpec{}
	// MISSING: Name
	// MISSING: CreateTime
	// MISSING: UpdateTime
	// MISSING: Labels
	// MISSING: DisplayName
	// MISSING: ApiConfig
	// MISSING: State
	// MISSING: DefaultHostname
	return out
}
func ApigatewayGatewaySpec_ToProto(mapCtx *direct.MapContext, in *krm.ApigatewayGatewaySpec) *pb.Gateway {
	if in == nil {
		return nil
	}
	out := &pb.Gateway{}
	// MISSING: Name
	// MISSING: CreateTime
	// MISSING: UpdateTime
	// MISSING: Labels
	// MISSING: DisplayName
	// MISSING: ApiConfig
	// MISSING: State
	// MISSING: DefaultHostname
	return out
}
func Gateway_FromProto(mapCtx *direct.MapContext, in *pb.Gateway) *krm.Gateway {
	if in == nil {
		return nil
	}
	out := &krm.Gateway{}
	// MISSING: Name
	// MISSING: CreateTime
	// MISSING: UpdateTime
	out.Labels = in.Labels
	out.DisplayName = direct.LazyPtr(in.GetDisplayName())
	out.ApiConfig = direct.LazyPtr(in.GetApiConfig())
	// MISSING: State
	// MISSING: DefaultHostname
	return out
}
func Gateway_ToProto(mapCtx *direct.MapContext, in *krm.Gateway) *pb.Gateway {
	if in == nil {
		return nil
	}
	out := &pb.Gateway{}
	// MISSING: Name
	// MISSING: CreateTime
	// MISSING: UpdateTime
	out.Labels = in.Labels
	out.DisplayName = direct.ValueOf(in.DisplayName)
	out.ApiConfig = direct.ValueOf(in.ApiConfig)
	// MISSING: State
	// MISSING: DefaultHostname
	return out
}
func GatewayObservedState_FromProto(mapCtx *direct.MapContext, in *pb.Gateway) *krm.GatewayObservedState {
	if in == nil {
		return nil
	}
	out := &krm.GatewayObservedState{}
	out.Name = direct.LazyPtr(in.GetName())
	out.CreateTime = direct.StringTimestamp_FromProto(mapCtx, in.GetCreateTime())
	out.UpdateTime = direct.StringTimestamp_FromProto(mapCtx, in.GetUpdateTime())
	// MISSING: Labels
	// MISSING: DisplayName
	// MISSING: ApiConfig
	out.State = direct.Enum_FromProto(mapCtx, in.GetState())
	out.DefaultHostname = direct.LazyPtr(in.GetDefaultHostname())
	return out
}
func GatewayObservedState_ToProto(mapCtx *direct.MapContext, in *krm.GatewayObservedState) *pb.Gateway {
	if in == nil {
		return nil
	}
	out := &pb.Gateway{}
	out.Name = direct.ValueOf(in.Name)
	out.CreateTime = direct.StringTimestamp_ToProto(mapCtx, in.CreateTime)
	out.UpdateTime = direct.StringTimestamp_ToProto(mapCtx, in.UpdateTime)
	// MISSING: Labels
	// MISSING: DisplayName
	// MISSING: ApiConfig
	out.State = direct.Enum_ToProto[pb.Gateway_State](mapCtx, in.State)
	out.DefaultHostname = direct.ValueOf(in.DefaultHostname)
	return out
}
