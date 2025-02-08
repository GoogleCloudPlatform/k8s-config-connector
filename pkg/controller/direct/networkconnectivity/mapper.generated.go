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
	krm "github.com/GoogleCloudPlatform/k8s-config-connector/apis/networkconnectivity/v1alpha1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct"
	refs "github.com/GoogleCloudPlatform/k8s-config-connector/apis/refs/v1beta1"
	pb "github.com/GoogleCloudPlatform/k8s-config-connector/mockgcp/generated/mockgcp/cloud/networkconnectivity/v1"
)
func AutoAccept_FromProto(mapCtx *direct.MapContext, in *pb.AutoAccept) *krm.AutoAccept {
	if in == nil {
		return nil
	}
	out := &krm.AutoAccept{}
	out.AutoAcceptProjects = in.AutoAcceptProjects
	return out
}
func AutoAccept_ToProto(mapCtx *direct.MapContext, in *krm.AutoAccept) *pb.AutoAccept {
	if in == nil {
		return nil
	}
	out := &pb.AutoAccept{}
	out.AutoAcceptProjects = in.AutoAcceptProjects
	return out
}
func Group_FromProto(mapCtx *direct.MapContext, in *pb.Group) *krm.Group {
	if in == nil {
		return nil
	}
	out := &krm.Group{}
	out.Name = direct.LazyPtr(in.GetName())
	// MISSING: CreateTime
	// MISSING: UpdateTime
	out.Labels = in.Labels
	out.Description = direct.LazyPtr(in.GetDescription())
	// MISSING: Uid
	// MISSING: State
	out.AutoAccept = AutoAccept_FromProto(mapCtx, in.GetAutoAccept())
	// MISSING: RouteTable
	return out
}
func Group_ToProto(mapCtx *direct.MapContext, in *krm.Group) *pb.Group {
	if in == nil {
		return nil
	}
	out := &pb.Group{}
	out.Name = direct.ValueOf(in.Name)
	// MISSING: CreateTime
	// MISSING: UpdateTime
	out.Labels = in.Labels
	out.Description = direct.ValueOf(in.Description)
	// MISSING: Uid
	// MISSING: State
	out.AutoAccept = AutoAccept_ToProto(mapCtx, in.AutoAccept)
	// MISSING: RouteTable
	return out
}
func GroupObservedState_FromProto(mapCtx *direct.MapContext, in *pb.Group) *krm.GroupObservedState {
	if in == nil {
		return nil
	}
	out := &krm.GroupObservedState{}
	// MISSING: Name
	out.CreateTime = direct.StringTimestamp_FromProto(mapCtx, in.GetCreateTime())
	out.UpdateTime = direct.StringTimestamp_FromProto(mapCtx, in.GetUpdateTime())
	// MISSING: Labels
	// MISSING: Description
	out.Uid = direct.LazyPtr(in.GetUid())
	out.State = direct.Enum_FromProto(mapCtx, in.GetState())
	// MISSING: AutoAccept
	out.RouteTable = direct.LazyPtr(in.GetRouteTable())
	return out
}
func GroupObservedState_ToProto(mapCtx *direct.MapContext, in *krm.GroupObservedState) *pb.Group {
	if in == nil {
		return nil
	}
	out := &pb.Group{}
	// MISSING: Name
	out.CreateTime = direct.StringTimestamp_ToProto(mapCtx, in.CreateTime)
	out.UpdateTime = direct.StringTimestamp_ToProto(mapCtx, in.UpdateTime)
	// MISSING: Labels
	// MISSING: Description
	out.Uid = direct.ValueOf(in.Uid)
	out.State = direct.Enum_ToProto[pb.State](mapCtx, in.State)
	// MISSING: AutoAccept
	out.RouteTable = direct.ValueOf(in.RouteTable)
	return out
}
func NetworkconnectivityGroupObservedState_FromProto(mapCtx *direct.MapContext, in *pb.Group) *krm.NetworkconnectivityGroupObservedState {
	if in == nil {
		return nil
	}
	out := &krm.NetworkconnectivityGroupObservedState{}
	// MISSING: Name
	// MISSING: CreateTime
	// MISSING: UpdateTime
	// MISSING: Labels
	// MISSING: Description
	// MISSING: Uid
	// MISSING: State
	// MISSING: AutoAccept
	// MISSING: RouteTable
	return out
}
func NetworkconnectivityGroupObservedState_ToProto(mapCtx *direct.MapContext, in *krm.NetworkconnectivityGroupObservedState) *pb.Group {
	if in == nil {
		return nil
	}
	out := &pb.Group{}
	// MISSING: Name
	// MISSING: CreateTime
	// MISSING: UpdateTime
	// MISSING: Labels
	// MISSING: Description
	// MISSING: Uid
	// MISSING: State
	// MISSING: AutoAccept
	// MISSING: RouteTable
	return out
}
func NetworkconnectivityGroupSpec_FromProto(mapCtx *direct.MapContext, in *pb.Group) *krm.NetworkconnectivityGroupSpec {
	if in == nil {
		return nil
	}
	out := &krm.NetworkconnectivityGroupSpec{}
	// MISSING: Name
	// MISSING: CreateTime
	// MISSING: UpdateTime
	// MISSING: Labels
	// MISSING: Description
	// MISSING: Uid
	// MISSING: State
	// MISSING: AutoAccept
	// MISSING: RouteTable
	return out
}
func NetworkconnectivityGroupSpec_ToProto(mapCtx *direct.MapContext, in *krm.NetworkconnectivityGroupSpec) *pb.Group {
	if in == nil {
		return nil
	}
	out := &pb.Group{}
	// MISSING: Name
	// MISSING: CreateTime
	// MISSING: UpdateTime
	// MISSING: Labels
	// MISSING: Description
	// MISSING: Uid
	// MISSING: State
	// MISSING: AutoAccept
	// MISSING: RouteTable
	return out
}
