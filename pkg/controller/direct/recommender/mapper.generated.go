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
	refs "github.com/GoogleCloudPlatform/k8s-config-connector/apis/refs/v1beta1"
	pb "cloud.google.com/go/recommender/apiv1/recommenderpb"
	krm "github.com/GoogleCloudPlatform/k8s-config-connector/apis/recommender/v1alpha1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct"
)
func InsightTypeConfig_FromProto(mapCtx *direct.MapContext, in *pb.InsightTypeConfig) *krm.InsightTypeConfig {
	if in == nil {
		return nil
	}
	out := &krm.InsightTypeConfig{}
	out.Name = direct.LazyPtr(in.GetName())
	out.InsightTypeGenerationConfig = InsightTypeGenerationConfig_FromProto(mapCtx, in.GetInsightTypeGenerationConfig())
	out.Etag = direct.LazyPtr(in.GetEtag())
	out.UpdateTime = direct.StringTimestamp_FromProto(mapCtx, in.GetUpdateTime())
	// MISSING: RevisionID
	out.Annotations = in.Annotations
	out.DisplayName = direct.LazyPtr(in.GetDisplayName())
	return out
}
func InsightTypeConfig_ToProto(mapCtx *direct.MapContext, in *krm.InsightTypeConfig) *pb.InsightTypeConfig {
	if in == nil {
		return nil
	}
	out := &pb.InsightTypeConfig{}
	out.Name = direct.ValueOf(in.Name)
	out.InsightTypeGenerationConfig = InsightTypeGenerationConfig_ToProto(mapCtx, in.InsightTypeGenerationConfig)
	out.Etag = direct.ValueOf(in.Etag)
	out.UpdateTime = direct.StringTimestamp_ToProto(mapCtx, in.UpdateTime)
	// MISSING: RevisionID
	out.Annotations = in.Annotations
	out.DisplayName = direct.ValueOf(in.DisplayName)
	return out
}
func InsightTypeConfigObservedState_FromProto(mapCtx *direct.MapContext, in *pb.InsightTypeConfig) *krm.InsightTypeConfigObservedState {
	if in == nil {
		return nil
	}
	out := &krm.InsightTypeConfigObservedState{}
	// MISSING: Name
	// MISSING: InsightTypeGenerationConfig
	// MISSING: Etag
	// MISSING: UpdateTime
	out.RevisionID = direct.LazyPtr(in.GetRevisionId())
	// MISSING: Annotations
	// MISSING: DisplayName
	return out
}
func InsightTypeConfigObservedState_ToProto(mapCtx *direct.MapContext, in *krm.InsightTypeConfigObservedState) *pb.InsightTypeConfig {
	if in == nil {
		return nil
	}
	out := &pb.InsightTypeConfig{}
	// MISSING: Name
	// MISSING: InsightTypeGenerationConfig
	// MISSING: Etag
	// MISSING: UpdateTime
	out.RevisionId = direct.ValueOf(in.RevisionID)
	// MISSING: Annotations
	// MISSING: DisplayName
	return out
}
func InsightTypeGenerationConfig_FromProto(mapCtx *direct.MapContext, in *pb.InsightTypeGenerationConfig) *krm.InsightTypeGenerationConfig {
	if in == nil {
		return nil
	}
	out := &krm.InsightTypeGenerationConfig{}
	out.Params = Params_FromProto(mapCtx, in.GetParams())
	return out
}
func InsightTypeGenerationConfig_ToProto(mapCtx *direct.MapContext, in *krm.InsightTypeGenerationConfig) *pb.InsightTypeGenerationConfig {
	if in == nil {
		return nil
	}
	out := &pb.InsightTypeGenerationConfig{}
	out.Params = Params_ToProto(mapCtx, in.Params)
	return out
}
func RecommenderInsightTypeConfigObservedState_FromProto(mapCtx *direct.MapContext, in *pb.InsightTypeConfig) *krm.RecommenderInsightTypeConfigObservedState {
	if in == nil {
		return nil
	}
	out := &krm.RecommenderInsightTypeConfigObservedState{}
	// MISSING: Name
	// MISSING: InsightTypeGenerationConfig
	// MISSING: Etag
	// MISSING: UpdateTime
	// MISSING: RevisionID
	// MISSING: Annotations
	// MISSING: DisplayName
	return out
}
func RecommenderInsightTypeConfigObservedState_ToProto(mapCtx *direct.MapContext, in *krm.RecommenderInsightTypeConfigObservedState) *pb.InsightTypeConfig {
	if in == nil {
		return nil
	}
	out := &pb.InsightTypeConfig{}
	// MISSING: Name
	// MISSING: InsightTypeGenerationConfig
	// MISSING: Etag
	// MISSING: UpdateTime
	// MISSING: RevisionID
	// MISSING: Annotations
	// MISSING: DisplayName
	return out
}
func RecommenderInsightTypeConfigSpec_FromProto(mapCtx *direct.MapContext, in *pb.InsightTypeConfig) *krm.RecommenderInsightTypeConfigSpec {
	if in == nil {
		return nil
	}
	out := &krm.RecommenderInsightTypeConfigSpec{}
	// MISSING: Name
	// MISSING: InsightTypeGenerationConfig
	// MISSING: Etag
	// MISSING: UpdateTime
	// MISSING: RevisionID
	// MISSING: Annotations
	// MISSING: DisplayName
	return out
}
func RecommenderInsightTypeConfigSpec_ToProto(mapCtx *direct.MapContext, in *krm.RecommenderInsightTypeConfigSpec) *pb.InsightTypeConfig {
	if in == nil {
		return nil
	}
	out := &pb.InsightTypeConfig{}
	// MISSING: Name
	// MISSING: InsightTypeGenerationConfig
	// MISSING: Etag
	// MISSING: UpdateTime
	// MISSING: RevisionID
	// MISSING: Annotations
	// MISSING: DisplayName
	return out
}
