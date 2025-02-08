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
func NetworkservicesTlsRouteObservedState_FromProto(mapCtx *direct.MapContext, in *pb.TlsRoute) *krm.NetworkservicesTlsRouteObservedState {
	if in == nil {
		return nil
	}
	out := &krm.NetworkservicesTlsRouteObservedState{}
	// MISSING: Name
	// MISSING: SelfLink
	// MISSING: CreateTime
	// MISSING: UpdateTime
	// MISSING: Description
	// MISSING: Rules
	// MISSING: Meshes
	// MISSING: Gateways
	return out
}
func NetworkservicesTlsRouteObservedState_ToProto(mapCtx *direct.MapContext, in *krm.NetworkservicesTlsRouteObservedState) *pb.TlsRoute {
	if in == nil {
		return nil
	}
	out := &pb.TlsRoute{}
	// MISSING: Name
	// MISSING: SelfLink
	// MISSING: CreateTime
	// MISSING: UpdateTime
	// MISSING: Description
	// MISSING: Rules
	// MISSING: Meshes
	// MISSING: Gateways
	return out
}
func NetworkservicesTlsRouteSpec_FromProto(mapCtx *direct.MapContext, in *pb.TlsRoute) *krm.NetworkservicesTlsRouteSpec {
	if in == nil {
		return nil
	}
	out := &krm.NetworkservicesTlsRouteSpec{}
	// MISSING: Name
	// MISSING: SelfLink
	// MISSING: CreateTime
	// MISSING: UpdateTime
	// MISSING: Description
	// MISSING: Rules
	// MISSING: Meshes
	// MISSING: Gateways
	return out
}
func NetworkservicesTlsRouteSpec_ToProto(mapCtx *direct.MapContext, in *krm.NetworkservicesTlsRouteSpec) *pb.TlsRoute {
	if in == nil {
		return nil
	}
	out := &pb.TlsRoute{}
	// MISSING: Name
	// MISSING: SelfLink
	// MISSING: CreateTime
	// MISSING: UpdateTime
	// MISSING: Description
	// MISSING: Rules
	// MISSING: Meshes
	// MISSING: Gateways
	return out
}
func TlsRoute_FromProto(mapCtx *direct.MapContext, in *pb.TlsRoute) *krm.TlsRoute {
	if in == nil {
		return nil
	}
	out := &krm.TlsRoute{}
	out.Name = direct.LazyPtr(in.GetName())
	// MISSING: SelfLink
	// MISSING: CreateTime
	// MISSING: UpdateTime
	out.Description = direct.LazyPtr(in.GetDescription())
	out.Rules = direct.Slice_FromProto(mapCtx, in.Rules, TlsRoute_RouteRule_FromProto)
	out.Meshes = in.Meshes
	out.Gateways = in.Gateways
	return out
}
func TlsRoute_ToProto(mapCtx *direct.MapContext, in *krm.TlsRoute) *pb.TlsRoute {
	if in == nil {
		return nil
	}
	out := &pb.TlsRoute{}
	out.Name = direct.ValueOf(in.Name)
	// MISSING: SelfLink
	// MISSING: CreateTime
	// MISSING: UpdateTime
	out.Description = direct.ValueOf(in.Description)
	out.Rules = direct.Slice_ToProto(mapCtx, in.Rules, TlsRoute_RouteRule_ToProto)
	out.Meshes = in.Meshes
	out.Gateways = in.Gateways
	return out
}
func TlsRouteObservedState_FromProto(mapCtx *direct.MapContext, in *pb.TlsRoute) *krm.TlsRouteObservedState {
	if in == nil {
		return nil
	}
	out := &krm.TlsRouteObservedState{}
	// MISSING: Name
	out.SelfLink = direct.LazyPtr(in.GetSelfLink())
	out.CreateTime = direct.StringTimestamp_FromProto(mapCtx, in.GetCreateTime())
	out.UpdateTime = direct.StringTimestamp_FromProto(mapCtx, in.GetUpdateTime())
	// MISSING: Description
	// MISSING: Rules
	// MISSING: Meshes
	// MISSING: Gateways
	return out
}
func TlsRouteObservedState_ToProto(mapCtx *direct.MapContext, in *krm.TlsRouteObservedState) *pb.TlsRoute {
	if in == nil {
		return nil
	}
	out := &pb.TlsRoute{}
	// MISSING: Name
	out.SelfLink = direct.ValueOf(in.SelfLink)
	out.CreateTime = direct.StringTimestamp_ToProto(mapCtx, in.CreateTime)
	out.UpdateTime = direct.StringTimestamp_ToProto(mapCtx, in.UpdateTime)
	// MISSING: Description
	// MISSING: Rules
	// MISSING: Meshes
	// MISSING: Gateways
	return out
}
func TlsRoute_RouteAction_FromProto(mapCtx *direct.MapContext, in *pb.TlsRoute_RouteAction) *krm.TlsRoute_RouteAction {
	if in == nil {
		return nil
	}
	out := &krm.TlsRoute_RouteAction{}
	out.Destinations = direct.Slice_FromProto(mapCtx, in.Destinations, TlsRoute_RouteDestination_FromProto)
	return out
}
func TlsRoute_RouteAction_ToProto(mapCtx *direct.MapContext, in *krm.TlsRoute_RouteAction) *pb.TlsRoute_RouteAction {
	if in == nil {
		return nil
	}
	out := &pb.TlsRoute_RouteAction{}
	out.Destinations = direct.Slice_ToProto(mapCtx, in.Destinations, TlsRoute_RouteDestination_ToProto)
	return out
}
func TlsRoute_RouteDestination_FromProto(mapCtx *direct.MapContext, in *pb.TlsRoute_RouteDestination) *krm.TlsRoute_RouteDestination {
	if in == nil {
		return nil
	}
	out := &krm.TlsRoute_RouteDestination{}
	out.ServiceName = direct.LazyPtr(in.GetServiceName())
	out.Weight = direct.LazyPtr(in.GetWeight())
	return out
}
func TlsRoute_RouteDestination_ToProto(mapCtx *direct.MapContext, in *krm.TlsRoute_RouteDestination) *pb.TlsRoute_RouteDestination {
	if in == nil {
		return nil
	}
	out := &pb.TlsRoute_RouteDestination{}
	out.ServiceName = direct.ValueOf(in.ServiceName)
	out.Weight = direct.ValueOf(in.Weight)
	return out
}
func TlsRoute_RouteMatch_FromProto(mapCtx *direct.MapContext, in *pb.TlsRoute_RouteMatch) *krm.TlsRoute_RouteMatch {
	if in == nil {
		return nil
	}
	out := &krm.TlsRoute_RouteMatch{}
	out.SniHost = in.SniHost
	out.Alpn = in.Alpn
	return out
}
func TlsRoute_RouteMatch_ToProto(mapCtx *direct.MapContext, in *krm.TlsRoute_RouteMatch) *pb.TlsRoute_RouteMatch {
	if in == nil {
		return nil
	}
	out := &pb.TlsRoute_RouteMatch{}
	out.SniHost = in.SniHost
	out.Alpn = in.Alpn
	return out
}
func TlsRoute_RouteRule_FromProto(mapCtx *direct.MapContext, in *pb.TlsRoute_RouteRule) *krm.TlsRoute_RouteRule {
	if in == nil {
		return nil
	}
	out := &krm.TlsRoute_RouteRule{}
	out.Matches = direct.Slice_FromProto(mapCtx, in.Matches, TlsRoute_RouteMatch_FromProto)
	out.Action = TlsRoute_RouteAction_FromProto(mapCtx, in.GetAction())
	return out
}
func TlsRoute_RouteRule_ToProto(mapCtx *direct.MapContext, in *krm.TlsRoute_RouteRule) *pb.TlsRoute_RouteRule {
	if in == nil {
		return nil
	}
	out := &pb.TlsRoute_RouteRule{}
	out.Matches = direct.Slice_ToProto(mapCtx, in.Matches, TlsRoute_RouteMatch_ToProto)
	out.Action = TlsRoute_RouteAction_ToProto(mapCtx, in.Action)
	return out
}
