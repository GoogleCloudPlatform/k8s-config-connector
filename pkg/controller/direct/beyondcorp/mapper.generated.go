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

package beyondcorp

import (
	refs "github.com/GoogleCloudPlatform/k8s-config-connector/apis/refs/v1beta1"
	pb "cloud.google.com/go/beyondcorp/appconnections/apiv1/appconnectionspb"
	krm "github.com/GoogleCloudPlatform/k8s-config-connector/apis/beyondcorp/v1alpha1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct"
)
func AppConnection_FromProto(mapCtx *direct.MapContext, in *pb.AppConnection) *krm.AppConnection {
	if in == nil {
		return nil
	}
	out := &krm.AppConnection{}
	out.Name = direct.LazyPtr(in.GetName())
	// MISSING: CreateTime
	// MISSING: UpdateTime
	out.Labels = in.Labels
	out.DisplayName = direct.LazyPtr(in.GetDisplayName())
	// MISSING: Uid
	out.Type = direct.Enum_FromProto(mapCtx, in.GetType())
	out.ApplicationEndpoint = AppConnection_ApplicationEndpoint_FromProto(mapCtx, in.GetApplicationEndpoint())
	out.Connectors = in.Connectors
	// MISSING: State
	out.Gateway = AppConnection_Gateway_FromProto(mapCtx, in.GetGateway())
	return out
}
func AppConnection_ToProto(mapCtx *direct.MapContext, in *krm.AppConnection) *pb.AppConnection {
	if in == nil {
		return nil
	}
	out := &pb.AppConnection{}
	out.Name = direct.ValueOf(in.Name)
	// MISSING: CreateTime
	// MISSING: UpdateTime
	out.Labels = in.Labels
	out.DisplayName = direct.ValueOf(in.DisplayName)
	// MISSING: Uid
	out.Type = direct.Enum_ToProto[pb.AppConnection_Type](mapCtx, in.Type)
	out.ApplicationEndpoint = AppConnection_ApplicationEndpoint_ToProto(mapCtx, in.ApplicationEndpoint)
	out.Connectors = in.Connectors
	// MISSING: State
	out.Gateway = AppConnection_Gateway_ToProto(mapCtx, in.Gateway)
	return out
}
func AppConnectionObservedState_FromProto(mapCtx *direct.MapContext, in *pb.AppConnection) *krm.AppConnectionObservedState {
	if in == nil {
		return nil
	}
	out := &krm.AppConnectionObservedState{}
	// MISSING: Name
	out.CreateTime = direct.StringTimestamp_FromProto(mapCtx, in.GetCreateTime())
	out.UpdateTime = direct.StringTimestamp_FromProto(mapCtx, in.GetUpdateTime())
	// MISSING: Labels
	// MISSING: DisplayName
	out.Uid = direct.LazyPtr(in.GetUid())
	// MISSING: Type
	// MISSING: ApplicationEndpoint
	// MISSING: Connectors
	out.State = direct.Enum_FromProto(mapCtx, in.GetState())
	out.Gateway = AppConnection_GatewayObservedState_FromProto(mapCtx, in.GetGateway())
	return out
}
func AppConnectionObservedState_ToProto(mapCtx *direct.MapContext, in *krm.AppConnectionObservedState) *pb.AppConnection {
	if in == nil {
		return nil
	}
	out := &pb.AppConnection{}
	// MISSING: Name
	out.CreateTime = direct.StringTimestamp_ToProto(mapCtx, in.CreateTime)
	out.UpdateTime = direct.StringTimestamp_ToProto(mapCtx, in.UpdateTime)
	// MISSING: Labels
	// MISSING: DisplayName
	out.Uid = direct.ValueOf(in.Uid)
	// MISSING: Type
	// MISSING: ApplicationEndpoint
	// MISSING: Connectors
	out.State = direct.Enum_ToProto[pb.AppConnection_State](mapCtx, in.State)
	out.Gateway = AppConnection_GatewayObservedState_ToProto(mapCtx, in.Gateway)
	return out
}
func AppConnection_ApplicationEndpoint_FromProto(mapCtx *direct.MapContext, in *pb.AppConnection_ApplicationEndpoint) *krm.AppConnection_ApplicationEndpoint {
	if in == nil {
		return nil
	}
	out := &krm.AppConnection_ApplicationEndpoint{}
	out.Host = direct.LazyPtr(in.GetHost())
	out.Port = direct.LazyPtr(in.GetPort())
	return out
}
func AppConnection_ApplicationEndpoint_ToProto(mapCtx *direct.MapContext, in *krm.AppConnection_ApplicationEndpoint) *pb.AppConnection_ApplicationEndpoint {
	if in == nil {
		return nil
	}
	out := &pb.AppConnection_ApplicationEndpoint{}
	out.Host = direct.ValueOf(in.Host)
	out.Port = direct.ValueOf(in.Port)
	return out
}
func AppConnection_Gateway_FromProto(mapCtx *direct.MapContext, in *pb.AppConnection_Gateway) *krm.AppConnection_Gateway {
	if in == nil {
		return nil
	}
	out := &krm.AppConnection_Gateway{}
	out.Type = direct.Enum_FromProto(mapCtx, in.GetType())
	// MISSING: URI
	// MISSING: IngressPort
	out.AppGateway = direct.LazyPtr(in.GetAppGateway())
	return out
}
func AppConnection_Gateway_ToProto(mapCtx *direct.MapContext, in *krm.AppConnection_Gateway) *pb.AppConnection_Gateway {
	if in == nil {
		return nil
	}
	out := &pb.AppConnection_Gateway{}
	out.Type = direct.Enum_ToProto[pb.AppConnection_Gateway_Type](mapCtx, in.Type)
	// MISSING: URI
	// MISSING: IngressPort
	out.AppGateway = direct.ValueOf(in.AppGateway)
	return out
}
func AppConnection_GatewayObservedState_FromProto(mapCtx *direct.MapContext, in *pb.AppConnection_Gateway) *krm.AppConnection_GatewayObservedState {
	if in == nil {
		return nil
	}
	out := &krm.AppConnection_GatewayObservedState{}
	// MISSING: Type
	out.URI = direct.LazyPtr(in.GetUri())
	out.IngressPort = direct.LazyPtr(in.GetIngressPort())
	// MISSING: AppGateway
	return out
}
func AppConnection_GatewayObservedState_ToProto(mapCtx *direct.MapContext, in *krm.AppConnection_GatewayObservedState) *pb.AppConnection_Gateway {
	if in == nil {
		return nil
	}
	out := &pb.AppConnection_Gateway{}
	// MISSING: Type
	out.Uri = direct.ValueOf(in.URI)
	out.IngressPort = direct.ValueOf(in.IngressPort)
	// MISSING: AppGateway
	return out
}
func BeyondcorpAppConnectionObservedState_FromProto(mapCtx *direct.MapContext, in *pb.AppConnection) *krm.BeyondcorpAppConnectionObservedState {
	if in == nil {
		return nil
	}
	out := &krm.BeyondcorpAppConnectionObservedState{}
	// MISSING: Name
	// MISSING: CreateTime
	// MISSING: UpdateTime
	// MISSING: Labels
	// MISSING: DisplayName
	// MISSING: Uid
	// MISSING: Type
	// MISSING: ApplicationEndpoint
	// MISSING: Connectors
	// MISSING: State
	// MISSING: Gateway
	return out
}
func BeyondcorpAppConnectionObservedState_ToProto(mapCtx *direct.MapContext, in *krm.BeyondcorpAppConnectionObservedState) *pb.AppConnection {
	if in == nil {
		return nil
	}
	out := &pb.AppConnection{}
	// MISSING: Name
	// MISSING: CreateTime
	// MISSING: UpdateTime
	// MISSING: Labels
	// MISSING: DisplayName
	// MISSING: Uid
	// MISSING: Type
	// MISSING: ApplicationEndpoint
	// MISSING: Connectors
	// MISSING: State
	// MISSING: Gateway
	return out
}
func BeyondcorpAppConnectionSpec_FromProto(mapCtx *direct.MapContext, in *pb.AppConnection) *krm.BeyondcorpAppConnectionSpec {
	if in == nil {
		return nil
	}
	out := &krm.BeyondcorpAppConnectionSpec{}
	// MISSING: Name
	// MISSING: CreateTime
	// MISSING: UpdateTime
	// MISSING: Labels
	// MISSING: DisplayName
	// MISSING: Uid
	// MISSING: Type
	// MISSING: ApplicationEndpoint
	// MISSING: Connectors
	// MISSING: State
	// MISSING: Gateway
	return out
}
func BeyondcorpAppConnectionSpec_ToProto(mapCtx *direct.MapContext, in *krm.BeyondcorpAppConnectionSpec) *pb.AppConnection {
	if in == nil {
		return nil
	}
	out := &pb.AppConnection{}
	// MISSING: Name
	// MISSING: CreateTime
	// MISSING: UpdateTime
	// MISSING: Labels
	// MISSING: DisplayName
	// MISSING: Uid
	// MISSING: Type
	// MISSING: ApplicationEndpoint
	// MISSING: Connectors
	// MISSING: State
	// MISSING: Gateway
	return out
}
