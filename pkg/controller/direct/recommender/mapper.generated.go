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
func Insight_FromProto(mapCtx *direct.MapContext, in *pb.Insight) *krm.Insight {
	if in == nil {
		return nil
	}
	out := &krm.Insight{}
	out.Name = direct.LazyPtr(in.GetName())
	out.Description = direct.LazyPtr(in.GetDescription())
	out.TargetResources = in.TargetResources
	out.InsightSubtype = direct.LazyPtr(in.GetInsightSubtype())
	out.Content = Content_FromProto(mapCtx, in.GetContent())
	out.LastRefreshTime = direct.StringTimestamp_FromProto(mapCtx, in.GetLastRefreshTime())
	out.ObservationPeriod = direct.StringDuration_FromProto(mapCtx, in.GetObservationPeriod())
	out.StateInfo = InsightStateInfo_FromProto(mapCtx, in.GetStateInfo())
	out.Category = direct.Enum_FromProto(mapCtx, in.GetCategory())
	out.Severity = direct.Enum_FromProto(mapCtx, in.GetSeverity())
	out.Etag = direct.LazyPtr(in.GetEtag())
	out.AssociatedRecommendations = direct.Slice_FromProto(mapCtx, in.AssociatedRecommendations, Insight_RecommendationReference_FromProto)
	return out
}
func Insight_ToProto(mapCtx *direct.MapContext, in *krm.Insight) *pb.Insight {
	if in == nil {
		return nil
	}
	out := &pb.Insight{}
	out.Name = direct.ValueOf(in.Name)
	out.Description = direct.ValueOf(in.Description)
	out.TargetResources = in.TargetResources
	out.InsightSubtype = direct.ValueOf(in.InsightSubtype)
	out.Content = Content_ToProto(mapCtx, in.Content)
	out.LastRefreshTime = direct.StringTimestamp_ToProto(mapCtx, in.LastRefreshTime)
	out.ObservationPeriod = direct.StringDuration_ToProto(mapCtx, in.ObservationPeriod)
	out.StateInfo = InsightStateInfo_ToProto(mapCtx, in.StateInfo)
	out.Category = direct.Enum_ToProto[pb.Insight_Category](mapCtx, in.Category)
	out.Severity = direct.Enum_ToProto[pb.Insight_Severity](mapCtx, in.Severity)
	out.Etag = direct.ValueOf(in.Etag)
	out.AssociatedRecommendations = direct.Slice_ToProto(mapCtx, in.AssociatedRecommendations, Insight_RecommendationReference_ToProto)
	return out
}
func InsightStateInfo_FromProto(mapCtx *direct.MapContext, in *pb.InsightStateInfo) *krm.InsightStateInfo {
	if in == nil {
		return nil
	}
	out := &krm.InsightStateInfo{}
	out.State = direct.Enum_FromProto(mapCtx, in.GetState())
	out.StateMetadata = in.StateMetadata
	return out
}
func InsightStateInfo_ToProto(mapCtx *direct.MapContext, in *krm.InsightStateInfo) *pb.InsightStateInfo {
	if in == nil {
		return nil
	}
	out := &pb.InsightStateInfo{}
	out.State = direct.Enum_ToProto[pb.InsightStateInfo_State](mapCtx, in.State)
	out.StateMetadata = in.StateMetadata
	return out
}
func Insight_RecommendationReference_FromProto(mapCtx *direct.MapContext, in *pb.Insight_RecommendationReference) *krm.Insight_RecommendationReference {
	if in == nil {
		return nil
	}
	out := &krm.Insight_RecommendationReference{}
	out.Recommendation = direct.LazyPtr(in.GetRecommendation())
	return out
}
func Insight_RecommendationReference_ToProto(mapCtx *direct.MapContext, in *krm.Insight_RecommendationReference) *pb.Insight_RecommendationReference {
	if in == nil {
		return nil
	}
	out := &pb.Insight_RecommendationReference{}
	out.Recommendation = direct.ValueOf(in.Recommendation)
	return out
}
func RecommenderInsightObservedState_FromProto(mapCtx *direct.MapContext, in *pb.Insight) *krm.RecommenderInsightObservedState {
	if in == nil {
		return nil
	}
	out := &krm.RecommenderInsightObservedState{}
	// MISSING: Name
	// MISSING: Description
	// MISSING: TargetResources
	// MISSING: InsightSubtype
	// MISSING: Content
	// MISSING: LastRefreshTime
	// MISSING: ObservationPeriod
	// MISSING: StateInfo
	// MISSING: Category
	// MISSING: Severity
	// MISSING: Etag
	// MISSING: AssociatedRecommendations
	return out
}
func RecommenderInsightObservedState_ToProto(mapCtx *direct.MapContext, in *krm.RecommenderInsightObservedState) *pb.Insight {
	if in == nil {
		return nil
	}
	out := &pb.Insight{}
	// MISSING: Name
	// MISSING: Description
	// MISSING: TargetResources
	// MISSING: InsightSubtype
	// MISSING: Content
	// MISSING: LastRefreshTime
	// MISSING: ObservationPeriod
	// MISSING: StateInfo
	// MISSING: Category
	// MISSING: Severity
	// MISSING: Etag
	// MISSING: AssociatedRecommendations
	return out
}
func RecommenderInsightSpec_FromProto(mapCtx *direct.MapContext, in *pb.Insight) *krm.RecommenderInsightSpec {
	if in == nil {
		return nil
	}
	out := &krm.RecommenderInsightSpec{}
	// MISSING: Name
	// MISSING: Description
	// MISSING: TargetResources
	// MISSING: InsightSubtype
	// MISSING: Content
	// MISSING: LastRefreshTime
	// MISSING: ObservationPeriod
	// MISSING: StateInfo
	// MISSING: Category
	// MISSING: Severity
	// MISSING: Etag
	// MISSING: AssociatedRecommendations
	return out
}
func RecommenderInsightSpec_ToProto(mapCtx *direct.MapContext, in *krm.RecommenderInsightSpec) *pb.Insight {
	if in == nil {
		return nil
	}
	out := &pb.Insight{}
	// MISSING: Name
	// MISSING: Description
	// MISSING: TargetResources
	// MISSING: InsightSubtype
	// MISSING: Content
	// MISSING: LastRefreshTime
	// MISSING: ObservationPeriod
	// MISSING: StateInfo
	// MISSING: Category
	// MISSING: Severity
	// MISSING: Etag
	// MISSING: AssociatedRecommendations
	return out
}
