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

package datastream

import (
	refs "github.com/GoogleCloudPlatform/k8s-config-connector/apis/refs/v1beta1"
	pb "cloud.google.com/go/datastream/apiv1/datastreampb"
	krm "github.com/GoogleCloudPlatform/k8s-config-connector/apis/datastream/v1alpha1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct"
)
func DatastreamRouteObservedState_FromProto(mapCtx *direct.MapContext, in *pb.Route) *krm.DatastreamRouteObservedState {
	if in == nil {
		return nil
	}
	out := &krm.DatastreamRouteObservedState{}
	// MISSING: Name
	// MISSING: CreateTime
	// MISSING: UpdateTime
	// MISSING: Labels
	// MISSING: DisplayName
	// MISSING: DestinationAddress
	// MISSING: DestinationPort
	return out
}
func DatastreamRouteObservedState_ToProto(mapCtx *direct.MapContext, in *krm.DatastreamRouteObservedState) *pb.Route {
	if in == nil {
		return nil
	}
	out := &pb.Route{}
	// MISSING: Name
	// MISSING: CreateTime
	// MISSING: UpdateTime
	// MISSING: Labels
	// MISSING: DisplayName
	// MISSING: DestinationAddress
	// MISSING: DestinationPort
	return out
}
func DatastreamRouteSpec_FromProto(mapCtx *direct.MapContext, in *pb.Route) *krm.DatastreamRouteSpec {
	if in == nil {
		return nil
	}
	out := &krm.DatastreamRouteSpec{}
	// MISSING: Name
	// MISSING: CreateTime
	// MISSING: UpdateTime
	// MISSING: Labels
	// MISSING: DisplayName
	// MISSING: DestinationAddress
	// MISSING: DestinationPort
	return out
}
func DatastreamRouteSpec_ToProto(mapCtx *direct.MapContext, in *krm.DatastreamRouteSpec) *pb.Route {
	if in == nil {
		return nil
	}
	out := &pb.Route{}
	// MISSING: Name
	// MISSING: CreateTime
	// MISSING: UpdateTime
	// MISSING: Labels
	// MISSING: DisplayName
	// MISSING: DestinationAddress
	// MISSING: DestinationPort
	return out
}
func Route_FromProto(mapCtx *direct.MapContext, in *pb.Route) *krm.Route {
	if in == nil {
		return nil
	}
	out := &krm.Route{}
	// MISSING: Name
	// MISSING: CreateTime
	// MISSING: UpdateTime
	out.Labels = in.Labels
	out.DisplayName = direct.LazyPtr(in.GetDisplayName())
	out.DestinationAddress = direct.LazyPtr(in.GetDestinationAddress())
	out.DestinationPort = direct.LazyPtr(in.GetDestinationPort())
	return out
}
func Route_ToProto(mapCtx *direct.MapContext, in *krm.Route) *pb.Route {
	if in == nil {
		return nil
	}
	out := &pb.Route{}
	// MISSING: Name
	// MISSING: CreateTime
	// MISSING: UpdateTime
	out.Labels = in.Labels
	out.DisplayName = direct.ValueOf(in.DisplayName)
	out.DestinationAddress = direct.ValueOf(in.DestinationAddress)
	out.DestinationPort = direct.ValueOf(in.DestinationPort)
	return out
}
func RouteObservedState_FromProto(mapCtx *direct.MapContext, in *pb.Route) *krm.RouteObservedState {
	if in == nil {
		return nil
	}
	out := &krm.RouteObservedState{}
	out.Name = direct.LazyPtr(in.GetName())
	out.CreateTime = direct.StringTimestamp_FromProto(mapCtx, in.GetCreateTime())
	out.UpdateTime = direct.StringTimestamp_FromProto(mapCtx, in.GetUpdateTime())
	// MISSING: Labels
	// MISSING: DisplayName
	// MISSING: DestinationAddress
	// MISSING: DestinationPort
	return out
}
func RouteObservedState_ToProto(mapCtx *direct.MapContext, in *krm.RouteObservedState) *pb.Route {
	if in == nil {
		return nil
	}
	out := &pb.Route{}
	out.Name = direct.ValueOf(in.Name)
	out.CreateTime = direct.StringTimestamp_ToProto(mapCtx, in.CreateTime)
	out.UpdateTime = direct.StringTimestamp_ToProto(mapCtx, in.UpdateTime)
	// MISSING: Labels
	// MISSING: DisplayName
	// MISSING: DestinationAddress
	// MISSING: DestinationPort
	return out
}
