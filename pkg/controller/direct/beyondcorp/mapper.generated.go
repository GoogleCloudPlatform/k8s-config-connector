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
	pb "cloud.google.com/go/beyondcorp/clientconnectorservices/apiv1/clientconnectorservicespb"
	krm "github.com/GoogleCloudPlatform/k8s-config-connector/apis/beyondcorp/v1alpha1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct"
	refs "github.com/GoogleCloudPlatform/k8s-config-connector/apis/refs/v1beta1"
)
func BeyondcorpClientConnectorServiceObservedState_FromProto(mapCtx *direct.MapContext, in *pb.ClientConnectorService) *krm.BeyondcorpClientConnectorServiceObservedState {
	if in == nil {
		return nil
	}
	out := &krm.BeyondcorpClientConnectorServiceObservedState{}
	// MISSING: Name
	// MISSING: CreateTime
	// MISSING: UpdateTime
	// MISSING: DisplayName
	// MISSING: Ingress
	// MISSING: Egress
	// MISSING: State
	return out
}
func BeyondcorpClientConnectorServiceObservedState_ToProto(mapCtx *direct.MapContext, in *krm.BeyondcorpClientConnectorServiceObservedState) *pb.ClientConnectorService {
	if in == nil {
		return nil
	}
	out := &pb.ClientConnectorService{}
	// MISSING: Name
	// MISSING: CreateTime
	// MISSING: UpdateTime
	// MISSING: DisplayName
	// MISSING: Ingress
	// MISSING: Egress
	// MISSING: State
	return out
}
func BeyondcorpClientConnectorServiceSpec_FromProto(mapCtx *direct.MapContext, in *pb.ClientConnectorService) *krm.BeyondcorpClientConnectorServiceSpec {
	if in == nil {
		return nil
	}
	out := &krm.BeyondcorpClientConnectorServiceSpec{}
	// MISSING: Name
	// MISSING: CreateTime
	// MISSING: UpdateTime
	// MISSING: DisplayName
	// MISSING: Ingress
	// MISSING: Egress
	// MISSING: State
	return out
}
func BeyondcorpClientConnectorServiceSpec_ToProto(mapCtx *direct.MapContext, in *krm.BeyondcorpClientConnectorServiceSpec) *pb.ClientConnectorService {
	if in == nil {
		return nil
	}
	out := &pb.ClientConnectorService{}
	// MISSING: Name
	// MISSING: CreateTime
	// MISSING: UpdateTime
	// MISSING: DisplayName
	// MISSING: Ingress
	// MISSING: Egress
	// MISSING: State
	return out
}
func ClientConnectorService_FromProto(mapCtx *direct.MapContext, in *pb.ClientConnectorService) *krm.ClientConnectorService {
	if in == nil {
		return nil
	}
	out := &krm.ClientConnectorService{}
	out.Name = direct.LazyPtr(in.GetName())
	// MISSING: CreateTime
	// MISSING: UpdateTime
	out.DisplayName = direct.LazyPtr(in.GetDisplayName())
	out.Ingress = ClientConnectorService_Ingress_FromProto(mapCtx, in.GetIngress())
	out.Egress = ClientConnectorService_Egress_FromProto(mapCtx, in.GetEgress())
	// MISSING: State
	return out
}
func ClientConnectorService_ToProto(mapCtx *direct.MapContext, in *krm.ClientConnectorService) *pb.ClientConnectorService {
	if in == nil {
		return nil
	}
	out := &pb.ClientConnectorService{}
	out.Name = direct.ValueOf(in.Name)
	// MISSING: CreateTime
	// MISSING: UpdateTime
	out.DisplayName = direct.ValueOf(in.DisplayName)
	out.Ingress = ClientConnectorService_Ingress_ToProto(mapCtx, in.Ingress)
	out.Egress = ClientConnectorService_Egress_ToProto(mapCtx, in.Egress)
	// MISSING: State
	return out
}
func ClientConnectorServiceObservedState_FromProto(mapCtx *direct.MapContext, in *pb.ClientConnectorService) *krm.ClientConnectorServiceObservedState {
	if in == nil {
		return nil
	}
	out := &krm.ClientConnectorServiceObservedState{}
	// MISSING: Name
	out.CreateTime = direct.StringTimestamp_FromProto(mapCtx, in.GetCreateTime())
	out.UpdateTime = direct.StringTimestamp_FromProto(mapCtx, in.GetUpdateTime())
	// MISSING: DisplayName
	// MISSING: Ingress
	// MISSING: Egress
	out.State = direct.Enum_FromProto(mapCtx, in.GetState())
	return out
}
func ClientConnectorServiceObservedState_ToProto(mapCtx *direct.MapContext, in *krm.ClientConnectorServiceObservedState) *pb.ClientConnectorService {
	if in == nil {
		return nil
	}
	out := &pb.ClientConnectorService{}
	// MISSING: Name
	out.CreateTime = direct.StringTimestamp_ToProto(mapCtx, in.CreateTime)
	out.UpdateTime = direct.StringTimestamp_ToProto(mapCtx, in.UpdateTime)
	// MISSING: DisplayName
	// MISSING: Ingress
	// MISSING: Egress
	out.State = direct.Enum_ToProto[pb.ClientConnectorService_State](mapCtx, in.State)
	return out
}
func ClientConnectorService_Egress_FromProto(mapCtx *direct.MapContext, in *pb.ClientConnectorService_Egress) *krm.ClientConnectorService_Egress {
	if in == nil {
		return nil
	}
	out := &krm.ClientConnectorService_Egress{}
	out.PeeredVpc = ClientConnectorService_Egress_PeeredVpc_FromProto(mapCtx, in.GetPeeredVpc())
	return out
}
func ClientConnectorService_Egress_ToProto(mapCtx *direct.MapContext, in *krm.ClientConnectorService_Egress) *pb.ClientConnectorService_Egress {
	if in == nil {
		return nil
	}
	out := &pb.ClientConnectorService_Egress{}
	if oneof := ClientConnectorService_Egress_PeeredVpc_ToProto(mapCtx, in.PeeredVpc); oneof != nil {
		out.DestinationType = &pb.ClientConnectorService_Egress_PeeredVpc_{PeeredVpc: oneof}
	}
	return out
}
func ClientConnectorService_Egress_PeeredVpc_FromProto(mapCtx *direct.MapContext, in *pb.ClientConnectorService_Egress_PeeredVpc) *krm.ClientConnectorService_Egress_PeeredVpc {
	if in == nil {
		return nil
	}
	out := &krm.ClientConnectorService_Egress_PeeredVpc{}
	out.NetworkVpc = direct.LazyPtr(in.GetNetworkVpc())
	return out
}
func ClientConnectorService_Egress_PeeredVpc_ToProto(mapCtx *direct.MapContext, in *krm.ClientConnectorService_Egress_PeeredVpc) *pb.ClientConnectorService_Egress_PeeredVpc {
	if in == nil {
		return nil
	}
	out := &pb.ClientConnectorService_Egress_PeeredVpc{}
	out.NetworkVpc = direct.ValueOf(in.NetworkVpc)
	return out
}
func ClientConnectorService_Ingress_FromProto(mapCtx *direct.MapContext, in *pb.ClientConnectorService_Ingress) *krm.ClientConnectorService_Ingress {
	if in == nil {
		return nil
	}
	out := &krm.ClientConnectorService_Ingress{}
	out.Config = ClientConnectorService_Ingress_Config_FromProto(mapCtx, in.GetConfig())
	return out
}
func ClientConnectorService_Ingress_ToProto(mapCtx *direct.MapContext, in *krm.ClientConnectorService_Ingress) *pb.ClientConnectorService_Ingress {
	if in == nil {
		return nil
	}
	out := &pb.ClientConnectorService_Ingress{}
	if oneof := ClientConnectorService_Ingress_Config_ToProto(mapCtx, in.Config); oneof != nil {
		out.IngressConfig = &pb.ClientConnectorService_Ingress_Config_{Config: oneof}
	}
	return out
}
func ClientConnectorService_Ingress_Config_FromProto(mapCtx *direct.MapContext, in *pb.ClientConnectorService_Ingress_Config) *krm.ClientConnectorService_Ingress_Config {
	if in == nil {
		return nil
	}
	out := &krm.ClientConnectorService_Ingress_Config{}
	out.TransportProtocol = direct.Enum_FromProto(mapCtx, in.GetTransportProtocol())
	out.DestinationRoutes = direct.Slice_FromProto(mapCtx, in.DestinationRoutes, ClientConnectorService_Ingress_Config_DestinationRoute_FromProto)
	return out
}
func ClientConnectorService_Ingress_Config_ToProto(mapCtx *direct.MapContext, in *krm.ClientConnectorService_Ingress_Config) *pb.ClientConnectorService_Ingress_Config {
	if in == nil {
		return nil
	}
	out := &pb.ClientConnectorService_Ingress_Config{}
	out.TransportProtocol = direct.Enum_ToProto[pb.ClientConnectorService_Ingress_Config_TransportProtocol](mapCtx, in.TransportProtocol)
	out.DestinationRoutes = direct.Slice_ToProto(mapCtx, in.DestinationRoutes, ClientConnectorService_Ingress_Config_DestinationRoute_ToProto)
	return out
}
func ClientConnectorService_Ingress_Config_DestinationRoute_FromProto(mapCtx *direct.MapContext, in *pb.ClientConnectorService_Ingress_Config_DestinationRoute) *krm.ClientConnectorService_Ingress_Config_DestinationRoute {
	if in == nil {
		return nil
	}
	out := &krm.ClientConnectorService_Ingress_Config_DestinationRoute{}
	out.Address = direct.LazyPtr(in.GetAddress())
	out.Netmask = direct.LazyPtr(in.GetNetmask())
	return out
}
func ClientConnectorService_Ingress_Config_DestinationRoute_ToProto(mapCtx *direct.MapContext, in *krm.ClientConnectorService_Ingress_Config_DestinationRoute) *pb.ClientConnectorService_Ingress_Config_DestinationRoute {
	if in == nil {
		return nil
	}
	out := &pb.ClientConnectorService_Ingress_Config_DestinationRoute{}
	out.Address = direct.ValueOf(in.Address)
	out.Netmask = direct.ValueOf(in.Netmask)
	return out
}
