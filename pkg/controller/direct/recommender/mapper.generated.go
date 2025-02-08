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

package recommender

import (
	pb "cloud.google.com/go/recommender/apiv1/recommenderpb"
	krm "github.com/GoogleCloudPlatform/k8s-config-connector/apis/recommender/v1alpha1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct"
	refs "github.com/GoogleCloudPlatform/k8s-config-connector/apis/refs/v1beta1"
)
func RecommenderConfig_FromProto(mapCtx *direct.MapContext, in *pb.RecommenderConfig) *krm.RecommenderConfig {
	if in == nil {
		return nil
	}
	out := &krm.RecommenderConfig{}
	out.Name = direct.LazyPtr(in.GetName())
	out.RecommenderGenerationConfig = RecommenderGenerationConfig_FromProto(mapCtx, in.GetRecommenderGenerationConfig())
	out.Etag = direct.LazyPtr(in.GetEtag())
	out.UpdateTime = direct.StringTimestamp_FromProto(mapCtx, in.GetUpdateTime())
	// MISSING: RevisionID
	out.Annotations = in.Annotations
	out.DisplayName = direct.LazyPtr(in.GetDisplayName())
	return out
}
func RecommenderConfig_ToProto(mapCtx *direct.MapContext, in *krm.RecommenderConfig) *pb.RecommenderConfig {
	if in == nil {
		return nil
	}
	out := &pb.RecommenderConfig{}
	out.Name = direct.ValueOf(in.Name)
	out.RecommenderGenerationConfig = RecommenderGenerationConfig_ToProto(mapCtx, in.RecommenderGenerationConfig)
	out.Etag = direct.ValueOf(in.Etag)
	out.UpdateTime = direct.StringTimestamp_ToProto(mapCtx, in.UpdateTime)
	// MISSING: RevisionID
	out.Annotations = in.Annotations
	out.DisplayName = direct.ValueOf(in.DisplayName)
	return out
}
func RecommenderConfigObservedState_FromProto(mapCtx *direct.MapContext, in *pb.RecommenderConfig) *krm.RecommenderConfigObservedState {
	if in == nil {
		return nil
	}
	out := &krm.RecommenderConfigObservedState{}
	// MISSING: Name
	// MISSING: RecommenderGenerationConfig
	// MISSING: Etag
	// MISSING: UpdateTime
	out.RevisionID = direct.LazyPtr(in.GetRevisionId())
	// MISSING: Annotations
	// MISSING: DisplayName
	return out
}
func RecommenderConfigObservedState_ToProto(mapCtx *direct.MapContext, in *krm.RecommenderConfigObservedState) *pb.RecommenderConfig {
	if in == nil {
		return nil
	}
	out := &pb.RecommenderConfig{}
	// MISSING: Name
	// MISSING: RecommenderGenerationConfig
	// MISSING: Etag
	// MISSING: UpdateTime
	out.RevisionId = direct.ValueOf(in.RevisionID)
	// MISSING: Annotations
	// MISSING: DisplayName
	return out
}
func RecommenderGenerationConfig_FromProto(mapCtx *direct.MapContext, in *pb.RecommenderGenerationConfig) *krm.RecommenderGenerationConfig {
	if in == nil {
		return nil
	}
	out := &krm.RecommenderGenerationConfig{}
	out.Params = Params_FromProto(mapCtx, in.GetParams())
	return out
}
func RecommenderGenerationConfig_ToProto(mapCtx *direct.MapContext, in *krm.RecommenderGenerationConfig) *pb.RecommenderGenerationConfig {
	if in == nil {
		return nil
	}
	out := &pb.RecommenderGenerationConfig{}
	out.Params = Params_ToProto(mapCtx, in.Params)
	return out
}
func RecommenderRecommenderConfigObservedState_FromProto(mapCtx *direct.MapContext, in *pb.RecommenderConfig) *krm.RecommenderRecommenderConfigObservedState {
	if in == nil {
		return nil
	}
	out := &krm.RecommenderRecommenderConfigObservedState{}
	// MISSING: Name
	// MISSING: RecommenderGenerationConfig
	// MISSING: Etag
	// MISSING: UpdateTime
	// MISSING: RevisionID
	// MISSING: Annotations
	// MISSING: DisplayName
	return out
}
func RecommenderRecommenderConfigObservedState_ToProto(mapCtx *direct.MapContext, in *krm.RecommenderRecommenderConfigObservedState) *pb.RecommenderConfig {
	if in == nil {
		return nil
	}
	out := &pb.RecommenderConfig{}
	// MISSING: Name
	// MISSING: RecommenderGenerationConfig
	// MISSING: Etag
	// MISSING: UpdateTime
	// MISSING: RevisionID
	// MISSING: Annotations
	// MISSING: DisplayName
	return out
}
func RecommenderRecommenderConfigSpec_FromProto(mapCtx *direct.MapContext, in *pb.RecommenderConfig) *krm.RecommenderRecommenderConfigSpec {
	if in == nil {
		return nil
	}
	out := &krm.RecommenderRecommenderConfigSpec{}
	// MISSING: Name
	// MISSING: RecommenderGenerationConfig
	// MISSING: Etag
	// MISSING: UpdateTime
	// MISSING: RevisionID
	// MISSING: Annotations
	// MISSING: DisplayName
	return out
}
func RecommenderRecommenderConfigSpec_ToProto(mapCtx *direct.MapContext, in *krm.RecommenderRecommenderConfigSpec) *pb.RecommenderConfig {
	if in == nil {
		return nil
	}
	out := &pb.RecommenderConfig{}
	// MISSING: Name
	// MISSING: RecommenderGenerationConfig
	// MISSING: Etag
	// MISSING: UpdateTime
	// MISSING: RevisionID
	// MISSING: Annotations
	// MISSING: DisplayName
	return out
}
