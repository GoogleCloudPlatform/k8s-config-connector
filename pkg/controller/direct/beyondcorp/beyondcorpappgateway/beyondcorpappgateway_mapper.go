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

package beyondcorpappgateway

import (
	pb "cloud.google.com/go/beyondcorp/appgateways/apiv1/appgatewayspb"
	krm "github.com/GoogleCloudPlatform/k8s-config-connector/apis/beyondcorp/v1alpha1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct"
)

func BeyondCorpAppGatewaySpec_FromProto(mapCtx *direct.MapContext, in *pb.AppGateway) *krm.BeyondCorpAppGatewaySpec {
	if in == nil {
		return nil
	}
	out := &krm.BeyondCorpAppGatewaySpec{}
	out.DisplayName = direct.LazyPtr(in.GetDisplayName())
	out.Type = direct.Enum_FromProto(mapCtx, in.GetType())
	out.HostType = direct.Enum_FromProto(mapCtx, in.GetHostType())
	return out
}

func BeyondCorpAppGatewaySpec_ToProto(mapCtx *direct.MapContext, in *krm.BeyondCorpAppGatewaySpec) *pb.AppGateway {
	if in == nil {
		return nil
	}
	out := &pb.AppGateway{}
	out.DisplayName = direct.ValueOf(in.DisplayName)
	out.Type = direct.Enum_ToProto[pb.AppGateway_Type](mapCtx, in.Type)
	out.HostType = direct.Enum_ToProto[pb.AppGateway_HostType](mapCtx, in.HostType)
	return out
}

func AppGatewayAllocatedConnectionStatus_FromProto(mapCtx *direct.MapContext, in *pb.AppGateway_AllocatedConnection) *krm.AppGatewayAllocatedConnectionStatus {
	if in == nil {
		return nil
	}
	out := &krm.AppGatewayAllocatedConnectionStatus{}
	out.IngressPort = direct.LazyPtr(in.GetIngressPort())
	out.PscUri = direct.LazyPtr(in.GetPscUri())
	return out
}

func AppGatewayAllocatedConnectionStatus_ToProto(mapCtx *direct.MapContext, in *krm.AppGatewayAllocatedConnectionStatus) *pb.AppGateway_AllocatedConnection {
	if in == nil {
		return nil
	}
	out := &pb.AppGateway_AllocatedConnection{}
	out.IngressPort = direct.ValueOf(in.IngressPort)
	out.PscUri = direct.ValueOf(in.PscUri)
	return out
}

func BeyondCorpAppGatewayStatus_FromProto(mapCtx *direct.MapContext, in *pb.AppGateway) *krm.BeyondCorpAppGatewayStatus {
	if in == nil {
		return nil
	}
	out := &krm.BeyondCorpAppGatewayStatus{}
	out.State = direct.Enum_FromProto(mapCtx, in.GetState())
	out.Uri = direct.LazyPtr(in.GetUri())
	out.AllocatedConnections = direct.Slice_FromProto(mapCtx, in.AllocatedConnections, AppGatewayAllocatedConnectionStatus_FromProto)
	return out
}

func BeyondCorpAppGatewayStatus_ToProto(mapCtx *direct.MapContext, in *krm.BeyondCorpAppGatewayStatus) *pb.AppGateway {
	if in == nil {
		return nil
	}
	out := &pb.AppGateway{}
	out.State = direct.Enum_ToProto[pb.AppGateway_State](mapCtx, in.State)
	out.Uri = direct.ValueOf(in.Uri)
	out.AllocatedConnections = direct.Slice_ToProto(mapCtx, in.AllocatedConnections, AppGatewayAllocatedConnectionStatus_ToProto)
	return out
}

func BeyondCorpAppGatewayObservedState_FromProto(mapCtx *direct.MapContext, in *pb.AppGateway) *krm.BeyondCorpAppGatewayObservedState {
	if in == nil {
		return nil
	}
	out := &krm.BeyondCorpAppGatewayObservedState{}
	return out
}

func BeyondCorpAppGatewayObservedState_ToProto(mapCtx *direct.MapContext, in *krm.BeyondCorpAppGatewayObservedState) *pb.AppGateway {
	if in == nil {
		return nil
	}
	out := &pb.AppGateway{}
	return out
}
