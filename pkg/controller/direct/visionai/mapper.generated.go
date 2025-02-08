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

package visionai

import (
	refs "github.com/GoogleCloudPlatform/k8s-config-connector/apis/refs/v1beta1"
	pb "cloud.google.com/go/visionai/apiv1/visionaipb"
	krm "github.com/GoogleCloudPlatform/k8s-config-connector/apis/visionai/v1alpha1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct"
)
func DeployedIndex_FromProto(mapCtx *direct.MapContext, in *pb.DeployedIndex) *krm.DeployedIndex {
	if in == nil {
		return nil
	}
	out := &krm.DeployedIndex{}
	out.Index = direct.LazyPtr(in.GetIndex())
	return out
}
func DeployedIndex_ToProto(mapCtx *direct.MapContext, in *krm.DeployedIndex) *pb.DeployedIndex {
	if in == nil {
		return nil
	}
	out := &pb.DeployedIndex{}
	out.Index = direct.ValueOf(in.Index)
	return out
}
func IndexEndpoint_FromProto(mapCtx *direct.MapContext, in *pb.IndexEndpoint) *krm.IndexEndpoint {
	if in == nil {
		return nil
	}
	out := &krm.IndexEndpoint{}
	// MISSING: Name
	out.DisplayName = direct.LazyPtr(in.GetDisplayName())
	out.Description = direct.LazyPtr(in.GetDescription())
	// MISSING: DeployedIndex
	// MISSING: State
	out.Labels = in.Labels
	// MISSING: CreateTime
	// MISSING: UpdateTime
	// MISSING: SatisfiesPzs
	// MISSING: SatisfiesPzi
	return out
}
func IndexEndpoint_ToProto(mapCtx *direct.MapContext, in *krm.IndexEndpoint) *pb.IndexEndpoint {
	if in == nil {
		return nil
	}
	out := &pb.IndexEndpoint{}
	// MISSING: Name
	out.DisplayName = direct.ValueOf(in.DisplayName)
	out.Description = direct.ValueOf(in.Description)
	// MISSING: DeployedIndex
	// MISSING: State
	out.Labels = in.Labels
	// MISSING: CreateTime
	// MISSING: UpdateTime
	// MISSING: SatisfiesPzs
	// MISSING: SatisfiesPzi
	return out
}
func IndexEndpointObservedState_FromProto(mapCtx *direct.MapContext, in *pb.IndexEndpoint) *krm.IndexEndpointObservedState {
	if in == nil {
		return nil
	}
	out := &krm.IndexEndpointObservedState{}
	out.Name = direct.LazyPtr(in.GetName())
	// MISSING: DisplayName
	// MISSING: Description
	out.DeployedIndex = DeployedIndex_FromProto(mapCtx, in.GetDeployedIndex())
	out.State = direct.Enum_FromProto(mapCtx, in.GetState())
	// MISSING: Labels
	out.CreateTime = direct.StringTimestamp_FromProto(mapCtx, in.GetCreateTime())
	out.UpdateTime = direct.StringTimestamp_FromProto(mapCtx, in.GetUpdateTime())
	out.SatisfiesPzs = in.SatisfiesPzs
	out.SatisfiesPzi = in.SatisfiesPzi
	return out
}
func IndexEndpointObservedState_ToProto(mapCtx *direct.MapContext, in *krm.IndexEndpointObservedState) *pb.IndexEndpoint {
	if in == nil {
		return nil
	}
	out := &pb.IndexEndpoint{}
	out.Name = direct.ValueOf(in.Name)
	// MISSING: DisplayName
	// MISSING: Description
	out.DeployedIndex = DeployedIndex_ToProto(mapCtx, in.DeployedIndex)
	out.State = direct.Enum_ToProto[pb.IndexEndpoint_State](mapCtx, in.State)
	// MISSING: Labels
	out.CreateTime = direct.StringTimestamp_ToProto(mapCtx, in.CreateTime)
	out.UpdateTime = direct.StringTimestamp_ToProto(mapCtx, in.UpdateTime)
	out.SatisfiesPzs = in.SatisfiesPzs
	out.SatisfiesPzi = in.SatisfiesPzi
	return out
}
func VisionaiIndexEndpointObservedState_FromProto(mapCtx *direct.MapContext, in *pb.IndexEndpoint) *krm.VisionaiIndexEndpointObservedState {
	if in == nil {
		return nil
	}
	out := &krm.VisionaiIndexEndpointObservedState{}
	// MISSING: Name
	// MISSING: DisplayName
	// MISSING: Description
	// MISSING: DeployedIndex
	// MISSING: State
	// MISSING: Labels
	// MISSING: CreateTime
	// MISSING: UpdateTime
	// MISSING: SatisfiesPzs
	// MISSING: SatisfiesPzi
	return out
}
func VisionaiIndexEndpointObservedState_ToProto(mapCtx *direct.MapContext, in *krm.VisionaiIndexEndpointObservedState) *pb.IndexEndpoint {
	if in == nil {
		return nil
	}
	out := &pb.IndexEndpoint{}
	// MISSING: Name
	// MISSING: DisplayName
	// MISSING: Description
	// MISSING: DeployedIndex
	// MISSING: State
	// MISSING: Labels
	// MISSING: CreateTime
	// MISSING: UpdateTime
	// MISSING: SatisfiesPzs
	// MISSING: SatisfiesPzi
	return out
}
func VisionaiIndexEndpointSpec_FromProto(mapCtx *direct.MapContext, in *pb.IndexEndpoint) *krm.VisionaiIndexEndpointSpec {
	if in == nil {
		return nil
	}
	out := &krm.VisionaiIndexEndpointSpec{}
	// MISSING: Name
	// MISSING: DisplayName
	// MISSING: Description
	// MISSING: DeployedIndex
	// MISSING: State
	// MISSING: Labels
	// MISSING: CreateTime
	// MISSING: UpdateTime
	// MISSING: SatisfiesPzs
	// MISSING: SatisfiesPzi
	return out
}
func VisionaiIndexEndpointSpec_ToProto(mapCtx *direct.MapContext, in *krm.VisionaiIndexEndpointSpec) *pb.IndexEndpoint {
	if in == nil {
		return nil
	}
	out := &pb.IndexEndpoint{}
	// MISSING: Name
	// MISSING: DisplayName
	// MISSING: Description
	// MISSING: DeployedIndex
	// MISSING: State
	// MISSING: Labels
	// MISSING: CreateTime
	// MISSING: UpdateTime
	// MISSING: SatisfiesPzs
	// MISSING: SatisfiesPzi
	return out
}
