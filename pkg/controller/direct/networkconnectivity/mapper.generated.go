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

package networkconnectivity

import (
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct"
	refs "github.com/GoogleCloudPlatform/k8s-config-connector/apis/refs/v1beta1"
	pb "github.com/GoogleCloudPlatform/k8s-config-connector/mockgcp/generated/mockgcp/cloud/networkconnectivity/v1"
	krm "github.com/GoogleCloudPlatform/k8s-config-connector/apis/networkconnectivity/v1alpha1"
)
func NetworkconnectivityRouteTableObservedState_FromProto(mapCtx *direct.MapContext, in *pb.RouteTable) *krm.NetworkconnectivityRouteTableObservedState {
	if in == nil {
		return nil
	}
	out := &krm.NetworkconnectivityRouteTableObservedState{}
	// MISSING: Name
	// MISSING: CreateTime
	// MISSING: UpdateTime
	// MISSING: Labels
	// MISSING: Description
	// MISSING: Uid
	// MISSING: State
	return out
}
func NetworkconnectivityRouteTableObservedState_ToProto(mapCtx *direct.MapContext, in *krm.NetworkconnectivityRouteTableObservedState) *pb.RouteTable {
	if in == nil {
		return nil
	}
	out := &pb.RouteTable{}
	// MISSING: Name
	// MISSING: CreateTime
	// MISSING: UpdateTime
	// MISSING: Labels
	// MISSING: Description
	// MISSING: Uid
	// MISSING: State
	return out
}
func NetworkconnectivityRouteTableSpec_FromProto(mapCtx *direct.MapContext, in *pb.RouteTable) *krm.NetworkconnectivityRouteTableSpec {
	if in == nil {
		return nil
	}
	out := &krm.NetworkconnectivityRouteTableSpec{}
	// MISSING: Name
	// MISSING: CreateTime
	// MISSING: UpdateTime
	// MISSING: Labels
	// MISSING: Description
	// MISSING: Uid
	// MISSING: State
	return out
}
func NetworkconnectivityRouteTableSpec_ToProto(mapCtx *direct.MapContext, in *krm.NetworkconnectivityRouteTableSpec) *pb.RouteTable {
	if in == nil {
		return nil
	}
	out := &pb.RouteTable{}
	// MISSING: Name
	// MISSING: CreateTime
	// MISSING: UpdateTime
	// MISSING: Labels
	// MISSING: Description
	// MISSING: Uid
	// MISSING: State
	return out
}
func RouteTable_FromProto(mapCtx *direct.MapContext, in *pb.RouteTable) *krm.RouteTable {
	if in == nil {
		return nil
	}
	out := &krm.RouteTable{}
	out.Name = direct.LazyPtr(in.GetName())
	// MISSING: CreateTime
	// MISSING: UpdateTime
	out.Labels = in.Labels
	out.Description = direct.LazyPtr(in.GetDescription())
	// MISSING: Uid
	// MISSING: State
	return out
}
func RouteTable_ToProto(mapCtx *direct.MapContext, in *krm.RouteTable) *pb.RouteTable {
	if in == nil {
		return nil
	}
	out := &pb.RouteTable{}
	out.Name = direct.ValueOf(in.Name)
	// MISSING: CreateTime
	// MISSING: UpdateTime
	out.Labels = in.Labels
	out.Description = direct.ValueOf(in.Description)
	// MISSING: Uid
	// MISSING: State
	return out
}
func RouteTableObservedState_FromProto(mapCtx *direct.MapContext, in *pb.RouteTable) *krm.RouteTableObservedState {
	if in == nil {
		return nil
	}
	out := &krm.RouteTableObservedState{}
	// MISSING: Name
	out.CreateTime = direct.StringTimestamp_FromProto(mapCtx, in.GetCreateTime())
	out.UpdateTime = direct.StringTimestamp_FromProto(mapCtx, in.GetUpdateTime())
	// MISSING: Labels
	// MISSING: Description
	out.Uid = direct.LazyPtr(in.GetUid())
	out.State = direct.Enum_FromProto(mapCtx, in.GetState())
	return out
}
func RouteTableObservedState_ToProto(mapCtx *direct.MapContext, in *krm.RouteTableObservedState) *pb.RouteTable {
	if in == nil {
		return nil
	}
	out := &pb.RouteTable{}
	// MISSING: Name
	out.CreateTime = direct.StringTimestamp_ToProto(mapCtx, in.CreateTime)
	out.UpdateTime = direct.StringTimestamp_ToProto(mapCtx, in.UpdateTime)
	// MISSING: Labels
	// MISSING: Description
	out.Uid = direct.ValueOf(in.Uid)
	out.State = direct.Enum_ToProto[pb.State](mapCtx, in.State)
	return out
}
