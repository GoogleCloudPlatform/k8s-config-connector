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
	pb "cloud.google.com/go/beyondcorp/clientgateways/apiv1/clientgatewayspb"
	krm "github.com/GoogleCloudPlatform/k8s-config-connector/apis/beyondcorp/v1alpha1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct"
)
func BeyondcorpClientGatewayObservedState_FromProto(mapCtx *direct.MapContext, in *pb.ClientGateway) *krm.BeyondcorpClientGatewayObservedState {
	if in == nil {
		return nil
	}
	out := &krm.BeyondcorpClientGatewayObservedState{}
	// MISSING: Name
	// MISSING: CreateTime
	// MISSING: UpdateTime
	// MISSING: State
	// MISSING: ID
	// MISSING: ClientConnectorService
	return out
}
func BeyondcorpClientGatewayObservedState_ToProto(mapCtx *direct.MapContext, in *krm.BeyondcorpClientGatewayObservedState) *pb.ClientGateway {
	if in == nil {
		return nil
	}
	out := &pb.ClientGateway{}
	// MISSING: Name
	// MISSING: CreateTime
	// MISSING: UpdateTime
	// MISSING: State
	// MISSING: ID
	// MISSING: ClientConnectorService
	return out
}
func BeyondcorpClientGatewaySpec_FromProto(mapCtx *direct.MapContext, in *pb.ClientGateway) *krm.BeyondcorpClientGatewaySpec {
	if in == nil {
		return nil
	}
	out := &krm.BeyondcorpClientGatewaySpec{}
	// MISSING: Name
	// MISSING: CreateTime
	// MISSING: UpdateTime
	// MISSING: State
	// MISSING: ID
	// MISSING: ClientConnectorService
	return out
}
func BeyondcorpClientGatewaySpec_ToProto(mapCtx *direct.MapContext, in *krm.BeyondcorpClientGatewaySpec) *pb.ClientGateway {
	if in == nil {
		return nil
	}
	out := &pb.ClientGateway{}
	// MISSING: Name
	// MISSING: CreateTime
	// MISSING: UpdateTime
	// MISSING: State
	// MISSING: ID
	// MISSING: ClientConnectorService
	return out
}
func ClientGateway_FromProto(mapCtx *direct.MapContext, in *pb.ClientGateway) *krm.ClientGateway {
	if in == nil {
		return nil
	}
	out := &krm.ClientGateway{}
	out.Name = direct.LazyPtr(in.GetName())
	// MISSING: CreateTime
	// MISSING: UpdateTime
	// MISSING: State
	// MISSING: ID
	// MISSING: ClientConnectorService
	return out
}
func ClientGateway_ToProto(mapCtx *direct.MapContext, in *krm.ClientGateway) *pb.ClientGateway {
	if in == nil {
		return nil
	}
	out := &pb.ClientGateway{}
	out.Name = direct.ValueOf(in.Name)
	// MISSING: CreateTime
	// MISSING: UpdateTime
	// MISSING: State
	// MISSING: ID
	// MISSING: ClientConnectorService
	return out
}
func ClientGatewayObservedState_FromProto(mapCtx *direct.MapContext, in *pb.ClientGateway) *krm.ClientGatewayObservedState {
	if in == nil {
		return nil
	}
	out := &krm.ClientGatewayObservedState{}
	// MISSING: Name
	out.CreateTime = direct.StringTimestamp_FromProto(mapCtx, in.GetCreateTime())
	out.UpdateTime = direct.StringTimestamp_FromProto(mapCtx, in.GetUpdateTime())
	out.State = direct.Enum_FromProto(mapCtx, in.GetState())
	out.ID = direct.LazyPtr(in.GetId())
	out.ClientConnectorService = direct.LazyPtr(in.GetClientConnectorService())
	return out
}
func ClientGatewayObservedState_ToProto(mapCtx *direct.MapContext, in *krm.ClientGatewayObservedState) *pb.ClientGateway {
	if in == nil {
		return nil
	}
	out := &pb.ClientGateway{}
	// MISSING: Name
	out.CreateTime = direct.StringTimestamp_ToProto(mapCtx, in.CreateTime)
	out.UpdateTime = direct.StringTimestamp_ToProto(mapCtx, in.UpdateTime)
	out.State = direct.Enum_ToProto[pb.ClientGateway_State](mapCtx, in.State)
	out.Id = direct.ValueOf(in.ID)
	out.ClientConnectorService = direct.ValueOf(in.ClientConnectorService)
	return out
}
