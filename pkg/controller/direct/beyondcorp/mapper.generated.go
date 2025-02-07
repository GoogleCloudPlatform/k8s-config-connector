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
	pb "cloud.google.com/go/beyondcorp/appgateways/apiv1/appgatewayspb"
	krm "github.com/GoogleCloudPlatform/k8s-config-connector/apis/beyondcorp/v1alpha1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct"
)
func AppGateway_FromProto(mapCtx *direct.MapContext, in *pb.AppGateway) *krm.AppGateway {
	if in == nil {
		return nil
	}
	out := &krm.AppGateway{}
	out.Name = direct.LazyPtr(in.GetName())
	// MISSING: CreateTime
	// MISSING: UpdateTime
	out.Labels = in.Labels
	out.DisplayName = direct.LazyPtr(in.GetDisplayName())
	// MISSING: Uid
	out.Type = direct.Enum_FromProto(mapCtx, in.GetType())
	// MISSING: State
	// MISSING: URI
	// MISSING: AllocatedConnections
	out.HostType = direct.Enum_FromProto(mapCtx, in.GetHostType())
	return out
}
func AppGateway_ToProto(mapCtx *direct.MapContext, in *krm.AppGateway) *pb.AppGateway {
	if in == nil {
		return nil
	}
	out := &pb.AppGateway{}
	out.Name = direct.ValueOf(in.Name)
	// MISSING: CreateTime
	// MISSING: UpdateTime
	out.Labels = in.Labels
	out.DisplayName = direct.ValueOf(in.DisplayName)
	// MISSING: Uid
	out.Type = direct.Enum_ToProto[pb.AppGateway_Type](mapCtx, in.Type)
	// MISSING: State
	// MISSING: URI
	// MISSING: AllocatedConnections
	out.HostType = direct.Enum_ToProto[pb.AppGateway_HostType](mapCtx, in.HostType)
	return out
}
func AppGatewayObservedState_FromProto(mapCtx *direct.MapContext, in *pb.AppGateway) *krm.AppGatewayObservedState {
	if in == nil {
		return nil
	}
	out := &krm.AppGatewayObservedState{}
	// MISSING: Name
	out.CreateTime = direct.StringTimestamp_FromProto(mapCtx, in.GetCreateTime())
	out.UpdateTime = direct.StringTimestamp_FromProto(mapCtx, in.GetUpdateTime())
	// MISSING: Labels
	// MISSING: DisplayName
	out.Uid = direct.LazyPtr(in.GetUid())
	// MISSING: Type
	out.State = direct.Enum_FromProto(mapCtx, in.GetState())
	out.URI = direct.LazyPtr(in.GetUri())
	out.AllocatedConnections = direct.Slice_FromProto(mapCtx, in.AllocatedConnections, AppGateway_AllocatedConnection_FromProto)
	// MISSING: HostType
	return out
}
func AppGatewayObservedState_ToProto(mapCtx *direct.MapContext, in *krm.AppGatewayObservedState) *pb.AppGateway {
	if in == nil {
		return nil
	}
	out := &pb.AppGateway{}
	// MISSING: Name
	out.CreateTime = direct.StringTimestamp_ToProto(mapCtx, in.CreateTime)
	out.UpdateTime = direct.StringTimestamp_ToProto(mapCtx, in.UpdateTime)
	// MISSING: Labels
	// MISSING: DisplayName
	out.Uid = direct.ValueOf(in.Uid)
	// MISSING: Type
	out.State = direct.Enum_ToProto[pb.AppGateway_State](mapCtx, in.State)
	out.Uri = direct.ValueOf(in.URI)
	out.AllocatedConnections = direct.Slice_ToProto(mapCtx, in.AllocatedConnections, AppGateway_AllocatedConnection_ToProto)
	// MISSING: HostType
	return out
}
func AppGateway_AllocatedConnection_FromProto(mapCtx *direct.MapContext, in *pb.AppGateway_AllocatedConnection) *krm.AppGateway_AllocatedConnection {
	if in == nil {
		return nil
	}
	out := &krm.AppGateway_AllocatedConnection{}
	out.PscURI = direct.LazyPtr(in.GetPscUri())
	out.IngressPort = direct.LazyPtr(in.GetIngressPort())
	return out
}
func AppGateway_AllocatedConnection_ToProto(mapCtx *direct.MapContext, in *krm.AppGateway_AllocatedConnection) *pb.AppGateway_AllocatedConnection {
	if in == nil {
		return nil
	}
	out := &pb.AppGateway_AllocatedConnection{}
	out.PscUri = direct.ValueOf(in.PscURI)
	out.IngressPort = direct.ValueOf(in.IngressPort)
	return out
}
func BeyondcorpAppGatewayObservedState_FromProto(mapCtx *direct.MapContext, in *pb.AppGateway) *krm.BeyondcorpAppGatewayObservedState {
	if in == nil {
		return nil
	}
	out := &krm.BeyondcorpAppGatewayObservedState{}
	// MISSING: Name
	// MISSING: CreateTime
	// MISSING: UpdateTime
	// MISSING: Labels
	// MISSING: DisplayName
	// MISSING: Uid
	// MISSING: Type
	// MISSING: State
	// MISSING: URI
	// MISSING: AllocatedConnections
	// MISSING: HostType
	return out
}
func BeyondcorpAppGatewayObservedState_ToProto(mapCtx *direct.MapContext, in *krm.BeyondcorpAppGatewayObservedState) *pb.AppGateway {
	if in == nil {
		return nil
	}
	out := &pb.AppGateway{}
	// MISSING: Name
	// MISSING: CreateTime
	// MISSING: UpdateTime
	// MISSING: Labels
	// MISSING: DisplayName
	// MISSING: Uid
	// MISSING: Type
	// MISSING: State
	// MISSING: URI
	// MISSING: AllocatedConnections
	// MISSING: HostType
	return out
}
func BeyondcorpAppGatewaySpec_FromProto(mapCtx *direct.MapContext, in *pb.AppGateway) *krm.BeyondcorpAppGatewaySpec {
	if in == nil {
		return nil
	}
	out := &krm.BeyondcorpAppGatewaySpec{}
	// MISSING: Name
	// MISSING: CreateTime
	// MISSING: UpdateTime
	// MISSING: Labels
	// MISSING: DisplayName
	// MISSING: Uid
	// MISSING: Type
	// MISSING: State
	// MISSING: URI
	// MISSING: AllocatedConnections
	// MISSING: HostType
	return out
}
func BeyondcorpAppGatewaySpec_ToProto(mapCtx *direct.MapContext, in *krm.BeyondcorpAppGatewaySpec) *pb.AppGateway {
	if in == nil {
		return nil
	}
	out := &pb.AppGateway{}
	// MISSING: Name
	// MISSING: CreateTime
	// MISSING: UpdateTime
	// MISSING: Labels
	// MISSING: DisplayName
	// MISSING: Uid
	// MISSING: Type
	// MISSING: State
	// MISSING: URI
	// MISSING: AllocatedConnections
	// MISSING: HostType
	return out
}
