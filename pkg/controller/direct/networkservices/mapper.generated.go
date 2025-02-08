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
	refs "github.com/GoogleCloudPlatform/k8s-config-connector/apis/refs/v1beta1"
	pb "cloud.google.com/go/networkservices/apiv1/networkservicespb"
	krm "github.com/GoogleCloudPlatform/k8s-config-connector/apis/networkservices/v1alpha1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct"
)
func NetworkservicesTcpRouteObservedState_FromProto(mapCtx *direct.MapContext, in *pb.TcpRoute) *krm.NetworkservicesTcpRouteObservedState {
	if in == nil {
		return nil
	}
	out := &krm.NetworkservicesTcpRouteObservedState{}
	// MISSING: Name
	// MISSING: SelfLink
	// MISSING: CreateTime
	// MISSING: UpdateTime
	// MISSING: Description
	// MISSING: Rules
	// MISSING: Meshes
	// MISSING: Gateways
	// MISSING: Labels
	return out
}
func NetworkservicesTcpRouteObservedState_ToProto(mapCtx *direct.MapContext, in *krm.NetworkservicesTcpRouteObservedState) *pb.TcpRoute {
	if in == nil {
		return nil
	}
	out := &pb.TcpRoute{}
	// MISSING: Name
	// MISSING: SelfLink
	// MISSING: CreateTime
	// MISSING: UpdateTime
	// MISSING: Description
	// MISSING: Rules
	// MISSING: Meshes
	// MISSING: Gateways
	// MISSING: Labels
	return out
}
func NetworkservicesTcpRouteSpec_FromProto(mapCtx *direct.MapContext, in *pb.TcpRoute) *krm.NetworkservicesTcpRouteSpec {
	if in == nil {
		return nil
	}
	out := &krm.NetworkservicesTcpRouteSpec{}
	// MISSING: Name
	// MISSING: SelfLink
	// MISSING: CreateTime
	// MISSING: UpdateTime
	// MISSING: Description
	// MISSING: Rules
	// MISSING: Meshes
	// MISSING: Gateways
	// MISSING: Labels
	return out
}
func NetworkservicesTcpRouteSpec_ToProto(mapCtx *direct.MapContext, in *krm.NetworkservicesTcpRouteSpec) *pb.TcpRoute {
	if in == nil {
		return nil
	}
	out := &pb.TcpRoute{}
	// MISSING: Name
	// MISSING: SelfLink
	// MISSING: CreateTime
	// MISSING: UpdateTime
	// MISSING: Description
	// MISSING: Rules
	// MISSING: Meshes
	// MISSING: Gateways
	// MISSING: Labels
	return out
}
func TcpRoute_FromProto(mapCtx *direct.MapContext, in *pb.TcpRoute) *krm.TcpRoute {
	if in == nil {
		return nil
	}
	out := &krm.TcpRoute{}
	out.Name = direct.LazyPtr(in.GetName())
	// MISSING: SelfLink
	// MISSING: CreateTime
	// MISSING: UpdateTime
	out.Description = direct.LazyPtr(in.GetDescription())
	out.Rules = direct.Slice_FromProto(mapCtx, in.Rules, TcpRoute_RouteRule_FromProto)
	out.Meshes = in.Meshes
	out.Gateways = in.Gateways
	out.Labels = in.Labels
	return out
}
func TcpRoute_ToProto(mapCtx *direct.MapContext, in *krm.TcpRoute) *pb.TcpRoute {
	if in == nil {
		return nil
	}
	out := &pb.TcpRoute{}
	out.Name = direct.ValueOf(in.Name)
	// MISSING: SelfLink
	// MISSING: CreateTime
	// MISSING: UpdateTime
	out.Description = direct.ValueOf(in.Description)
	out.Rules = direct.Slice_ToProto(mapCtx, in.Rules, TcpRoute_RouteRule_ToProto)
	out.Meshes = in.Meshes
	out.Gateways = in.Gateways
	out.Labels = in.Labels
	return out
}
func TcpRouteObservedState_FromProto(mapCtx *direct.MapContext, in *pb.TcpRoute) *krm.TcpRouteObservedState {
	if in == nil {
		return nil
	}
	out := &krm.TcpRouteObservedState{}
	// MISSING: Name
	out.SelfLink = direct.LazyPtr(in.GetSelfLink())
	out.CreateTime = direct.StringTimestamp_FromProto(mapCtx, in.GetCreateTime())
	out.UpdateTime = direct.StringTimestamp_FromProto(mapCtx, in.GetUpdateTime())
	// MISSING: Description
	// MISSING: Rules
	// MISSING: Meshes
	// MISSING: Gateways
	// MISSING: Labels
	return out
}
func TcpRouteObservedState_ToProto(mapCtx *direct.MapContext, in *krm.TcpRouteObservedState) *pb.TcpRoute {
	if in == nil {
		return nil
	}
	out := &pb.TcpRoute{}
	// MISSING: Name
	out.SelfLink = direct.ValueOf(in.SelfLink)
	out.CreateTime = direct.StringTimestamp_ToProto(mapCtx, in.CreateTime)
	out.UpdateTime = direct.StringTimestamp_ToProto(mapCtx, in.UpdateTime)
	// MISSING: Description
	// MISSING: Rules
	// MISSING: Meshes
	// MISSING: Gateways
	// MISSING: Labels
	return out
}
func TcpRoute_RouteAction_FromProto(mapCtx *direct.MapContext, in *pb.TcpRoute_RouteAction) *krm.TcpRoute_RouteAction {
	if in == nil {
		return nil
	}
	out := &krm.TcpRoute_RouteAction{}
	out.Destinations = direct.Slice_FromProto(mapCtx, in.Destinations, TcpRoute_RouteDestination_FromProto)
	out.OriginalDestination = direct.LazyPtr(in.GetOriginalDestination())
	return out
}
func TcpRoute_RouteAction_ToProto(mapCtx *direct.MapContext, in *krm.TcpRoute_RouteAction) *pb.TcpRoute_RouteAction {
	if in == nil {
		return nil
	}
	out := &pb.TcpRoute_RouteAction{}
	out.Destinations = direct.Slice_ToProto(mapCtx, in.Destinations, TcpRoute_RouteDestination_ToProto)
	out.OriginalDestination = direct.ValueOf(in.OriginalDestination)
	return out
}
func TcpRoute_RouteDestination_FromProto(mapCtx *direct.MapContext, in *pb.TcpRoute_RouteDestination) *krm.TcpRoute_RouteDestination {
	if in == nil {
		return nil
	}
	out := &krm.TcpRoute_RouteDestination{}
	out.ServiceName = direct.LazyPtr(in.GetServiceName())
	out.Weight = direct.LazyPtr(in.GetWeight())
	return out
}
func TcpRoute_RouteDestination_ToProto(mapCtx *direct.MapContext, in *krm.TcpRoute_RouteDestination) *pb.TcpRoute_RouteDestination {
	if in == nil {
		return nil
	}
	out := &pb.TcpRoute_RouteDestination{}
	out.ServiceName = direct.ValueOf(in.ServiceName)
	out.Weight = direct.ValueOf(in.Weight)
	return out
}
func TcpRoute_RouteMatch_FromProto(mapCtx *direct.MapContext, in *pb.TcpRoute_RouteMatch) *krm.TcpRoute_RouteMatch {
	if in == nil {
		return nil
	}
	out := &krm.TcpRoute_RouteMatch{}
	out.Address = direct.LazyPtr(in.GetAddress())
	out.Port = direct.LazyPtr(in.GetPort())
	return out
}
func TcpRoute_RouteMatch_ToProto(mapCtx *direct.MapContext, in *krm.TcpRoute_RouteMatch) *pb.TcpRoute_RouteMatch {
	if in == nil {
		return nil
	}
	out := &pb.TcpRoute_RouteMatch{}
	out.Address = direct.ValueOf(in.Address)
	out.Port = direct.ValueOf(in.Port)
	return out
}
func TcpRoute_RouteRule_FromProto(mapCtx *direct.MapContext, in *pb.TcpRoute_RouteRule) *krm.TcpRoute_RouteRule {
	if in == nil {
		return nil
	}
	out := &krm.TcpRoute_RouteRule{}
	out.Matches = direct.Slice_FromProto(mapCtx, in.Matches, TcpRoute_RouteMatch_FromProto)
	out.Action = TcpRoute_RouteAction_FromProto(mapCtx, in.GetAction())
	return out
}
func TcpRoute_RouteRule_ToProto(mapCtx *direct.MapContext, in *krm.TcpRoute_RouteRule) *pb.TcpRoute_RouteRule {
	if in == nil {
		return nil
	}
	out := &pb.TcpRoute_RouteRule{}
	out.Matches = direct.Slice_ToProto(mapCtx, in.Matches, TcpRoute_RouteMatch_ToProto)
	out.Action = TcpRoute_RouteAction_ToProto(mapCtx, in.Action)
	return out
}
