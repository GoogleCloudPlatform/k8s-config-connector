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
	krm "github.com/GoogleCloudPlatform/k8s-config-connector/apis/visionai/v1alpha1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct"
	refs "github.com/GoogleCloudPlatform/k8s-config-connector/apis/refs/v1beta1"
	pb "cloud.google.com/go/visionai/apiv1/visionaipb"
)
func DeployedIndexReference_FromProto(mapCtx *direct.MapContext, in *pb.DeployedIndexReference) *krm.DeployedIndexReference {
	if in == nil {
		return nil
	}
	out := &krm.DeployedIndexReference{}
	out.IndexEndpoint = direct.LazyPtr(in.GetIndexEndpoint())
	return out
}
func DeployedIndexReference_ToProto(mapCtx *direct.MapContext, in *krm.DeployedIndexReference) *pb.DeployedIndexReference {
	if in == nil {
		return nil
	}
	out := &pb.DeployedIndexReference{}
	out.IndexEndpoint = direct.ValueOf(in.IndexEndpoint)
	return out
}
func Index_FromProto(mapCtx *direct.MapContext, in *pb.Index) *krm.Index {
	if in == nil {
		return nil
	}
	out := &krm.Index{}
	out.EntireCorpus = direct.LazyPtr(in.GetEntireCorpus())
	// MISSING: Name
	out.DisplayName = direct.LazyPtr(in.GetDisplayName())
	out.Description = direct.LazyPtr(in.GetDescription())
	// MISSING: State
	// MISSING: CreateTime
	// MISSING: UpdateTime
	// MISSING: DeployedIndexes
	// MISSING: SatisfiesPzs
	// MISSING: SatisfiesPzi
	return out
}
func Index_ToProto(mapCtx *direct.MapContext, in *krm.Index) *pb.Index {
	if in == nil {
		return nil
	}
	out := &pb.Index{}
	if oneof := Index_EntireCorpus_ToProto(mapCtx, in.EntireCorpus); oneof != nil {
		out.AssetFilter = oneof
	}
	// MISSING: Name
	out.DisplayName = direct.ValueOf(in.DisplayName)
	out.Description = direct.ValueOf(in.Description)
	// MISSING: State
	// MISSING: CreateTime
	// MISSING: UpdateTime
	// MISSING: DeployedIndexes
	// MISSING: SatisfiesPzs
	// MISSING: SatisfiesPzi
	return out
}
func IndexObservedState_FromProto(mapCtx *direct.MapContext, in *pb.Index) *krm.IndexObservedState {
	if in == nil {
		return nil
	}
	out := &krm.IndexObservedState{}
	// MISSING: EntireCorpus
	out.Name = direct.LazyPtr(in.GetName())
	// MISSING: DisplayName
	// MISSING: Description
	out.State = direct.Enum_FromProto(mapCtx, in.GetState())
	out.CreateTime = direct.StringTimestamp_FromProto(mapCtx, in.GetCreateTime())
	out.UpdateTime = direct.StringTimestamp_FromProto(mapCtx, in.GetUpdateTime())
	out.DeployedIndexes = direct.Slice_FromProto(mapCtx, in.DeployedIndexes, DeployedIndexReference_FromProto)
	out.SatisfiesPzs = in.SatisfiesPzs
	out.SatisfiesPzi = in.SatisfiesPzi
	return out
}
func IndexObservedState_ToProto(mapCtx *direct.MapContext, in *krm.IndexObservedState) *pb.Index {
	if in == nil {
		return nil
	}
	out := &pb.Index{}
	// MISSING: EntireCorpus
	out.Name = direct.ValueOf(in.Name)
	// MISSING: DisplayName
	// MISSING: Description
	out.State = direct.Enum_ToProto[pb.Index_State](mapCtx, in.State)
	out.CreateTime = direct.StringTimestamp_ToProto(mapCtx, in.CreateTime)
	out.UpdateTime = direct.StringTimestamp_ToProto(mapCtx, in.UpdateTime)
	out.DeployedIndexes = direct.Slice_ToProto(mapCtx, in.DeployedIndexes, DeployedIndexReference_ToProto)
	out.SatisfiesPzs = in.SatisfiesPzs
	out.SatisfiesPzi = in.SatisfiesPzi
	return out
}
func VisionaiIndexObservedState_FromProto(mapCtx *direct.MapContext, in *pb.Index) *krm.VisionaiIndexObservedState {
	if in == nil {
		return nil
	}
	out := &krm.VisionaiIndexObservedState{}
	// MISSING: EntireCorpus
	// MISSING: Name
	// MISSING: DisplayName
	// MISSING: Description
	// MISSING: State
	// MISSING: CreateTime
	// MISSING: UpdateTime
	// MISSING: DeployedIndexes
	// MISSING: SatisfiesPzs
	// MISSING: SatisfiesPzi
	return out
}
func VisionaiIndexObservedState_ToProto(mapCtx *direct.MapContext, in *krm.VisionaiIndexObservedState) *pb.Index {
	if in == nil {
		return nil
	}
	out := &pb.Index{}
	// MISSING: EntireCorpus
	// MISSING: Name
	// MISSING: DisplayName
	// MISSING: Description
	// MISSING: State
	// MISSING: CreateTime
	// MISSING: UpdateTime
	// MISSING: DeployedIndexes
	// MISSING: SatisfiesPzs
	// MISSING: SatisfiesPzi
	return out
}
func VisionaiIndexSpec_FromProto(mapCtx *direct.MapContext, in *pb.Index) *krm.VisionaiIndexSpec {
	if in == nil {
		return nil
	}
	out := &krm.VisionaiIndexSpec{}
	// MISSING: EntireCorpus
	// MISSING: Name
	// MISSING: DisplayName
	// MISSING: Description
	// MISSING: State
	// MISSING: CreateTime
	// MISSING: UpdateTime
	// MISSING: DeployedIndexes
	// MISSING: SatisfiesPzs
	// MISSING: SatisfiesPzi
	return out
}
func VisionaiIndexSpec_ToProto(mapCtx *direct.MapContext, in *krm.VisionaiIndexSpec) *pb.Index {
	if in == nil {
		return nil
	}
	out := &pb.Index{}
	// MISSING: EntireCorpus
	// MISSING: Name
	// MISSING: DisplayName
	// MISSING: Description
	// MISSING: State
	// MISSING: CreateTime
	// MISSING: UpdateTime
	// MISSING: DeployedIndexes
	// MISSING: SatisfiesPzs
	// MISSING: SatisfiesPzi
	return out
}
