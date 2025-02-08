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
func CostProjection_FromProto(mapCtx *direct.MapContext, in *pb.CostProjection) *krm.CostProjection {
	if in == nil {
		return nil
	}
	out := &krm.CostProjection{}
	out.Cost = Money_FromProto(mapCtx, in.GetCost())
	out.Duration = direct.StringDuration_FromProto(mapCtx, in.GetDuration())
	out.CostInLocalCurrency = Money_FromProto(mapCtx, in.GetCostInLocalCurrency())
	return out
}
func CostProjection_ToProto(mapCtx *direct.MapContext, in *krm.CostProjection) *pb.CostProjection {
	if in == nil {
		return nil
	}
	out := &pb.CostProjection{}
	out.Cost = Money_ToProto(mapCtx, in.Cost)
	out.Duration = direct.StringDuration_ToProto(mapCtx, in.Duration)
	out.CostInLocalCurrency = Money_ToProto(mapCtx, in.CostInLocalCurrency)
	return out
}
func Impact_FromProto(mapCtx *direct.MapContext, in *pb.Impact) *krm.Impact {
	if in == nil {
		return nil
	}
	out := &krm.Impact{}
	out.Category = direct.Enum_FromProto(mapCtx, in.GetCategory())
	out.CostProjection = CostProjection_FromProto(mapCtx, in.GetCostProjection())
	out.SecurityProjection = SecurityProjection_FromProto(mapCtx, in.GetSecurityProjection())
	out.SustainabilityProjection = SustainabilityProjection_FromProto(mapCtx, in.GetSustainabilityProjection())
	out.ReliabilityProjection = ReliabilityProjection_FromProto(mapCtx, in.GetReliabilityProjection())
	return out
}
func Impact_ToProto(mapCtx *direct.MapContext, in *krm.Impact) *pb.Impact {
	if in == nil {
		return nil
	}
	out := &pb.Impact{}
	out.Category = direct.Enum_ToProto[pb.Impact_Category](mapCtx, in.Category)
	if oneof := CostProjection_ToProto(mapCtx, in.CostProjection); oneof != nil {
		out.Projection = &pb.Impact_CostProjection{CostProjection: oneof}
	}
	if oneof := SecurityProjection_ToProto(mapCtx, in.SecurityProjection); oneof != nil {
		out.Projection = &pb.Impact_SecurityProjection{SecurityProjection: oneof}
	}
	if oneof := SustainabilityProjection_ToProto(mapCtx, in.SustainabilityProjection); oneof != nil {
		out.Projection = &pb.Impact_SustainabilityProjection{SustainabilityProjection: oneof}
	}
	if oneof := ReliabilityProjection_ToProto(mapCtx, in.ReliabilityProjection); oneof != nil {
		out.Projection = &pb.Impact_ReliabilityProjection{ReliabilityProjection: oneof}
	}
	return out
}
func Operation_FromProto(mapCtx *direct.MapContext, in *pb.Operation) *krm.Operation {
	if in == nil {
		return nil
	}
	out := &krm.Operation{}
	out.Action = direct.LazyPtr(in.GetAction())
	out.ResourceType = direct.LazyPtr(in.GetResourceType())
	out.Resource = direct.LazyPtr(in.GetResource())
	out.Path = direct.LazyPtr(in.GetPath())
	out.SourceResource = direct.LazyPtr(in.GetSourceResource())
	out.SourcePath = direct.LazyPtr(in.GetSourcePath())
	out.Value = Value_FromProto(mapCtx, in.GetValue())
	out.ValueMatcher = ValueMatcher_FromProto(mapCtx, in.GetValueMatcher())
	// MISSING: PathFilters
	// MISSING: PathValueMatchers
	return out
}
func Operation_ToProto(mapCtx *direct.MapContext, in *krm.Operation) *pb.Operation {
	if in == nil {
		return nil
	}
	out := &pb.Operation{}
	out.Action = direct.ValueOf(in.Action)
	out.ResourceType = direct.ValueOf(in.ResourceType)
	out.Resource = direct.ValueOf(in.Resource)
	out.Path = direct.ValueOf(in.Path)
	out.SourceResource = direct.ValueOf(in.SourceResource)
	out.SourcePath = direct.ValueOf(in.SourcePath)
	if oneof := Value_ToProto(mapCtx, in.Value); oneof != nil {
		out.PathValue = &pb.Operation_Value{Value: oneof}
	}
	if oneof := ValueMatcher_ToProto(mapCtx, in.ValueMatcher); oneof != nil {
		out.PathValue = &pb.Operation_ValueMatcher{ValueMatcher: oneof}
	}
	// MISSING: PathFilters
	// MISSING: PathValueMatchers
	return out
}
func OperationGroup_FromProto(mapCtx *direct.MapContext, in *pb.OperationGroup) *krm.OperationGroup {
	if in == nil {
		return nil
	}
	out := &krm.OperationGroup{}
	out.Operations = direct.Slice_FromProto(mapCtx, in.Operations, Operation_FromProto)
	return out
}
func OperationGroup_ToProto(mapCtx *direct.MapContext, in *krm.OperationGroup) *pb.OperationGroup {
	if in == nil {
		return nil
	}
	out := &pb.OperationGroup{}
	out.Operations = direct.Slice_ToProto(mapCtx, in.Operations, Operation_ToProto)
	return out
}
func Recommendation_FromProto(mapCtx *direct.MapContext, in *pb.Recommendation) *krm.Recommendation {
	if in == nil {
		return nil
	}
	out := &krm.Recommendation{}
	out.Name = direct.LazyPtr(in.GetName())
	out.Description = direct.LazyPtr(in.GetDescription())
	out.RecommenderSubtype = direct.LazyPtr(in.GetRecommenderSubtype())
	out.LastRefreshTime = direct.StringTimestamp_FromProto(mapCtx, in.GetLastRefreshTime())
	out.PrimaryImpact = Impact_FromProto(mapCtx, in.GetPrimaryImpact())
	out.AdditionalImpact = direct.Slice_FromProto(mapCtx, in.AdditionalImpact, Impact_FromProto)
	out.Priority = direct.Enum_FromProto(mapCtx, in.GetPriority())
	out.Content = RecommendationContent_FromProto(mapCtx, in.GetContent())
	out.StateInfo = RecommendationStateInfo_FromProto(mapCtx, in.GetStateInfo())
	out.Etag = direct.LazyPtr(in.GetEtag())
	out.AssociatedInsights = direct.Slice_FromProto(mapCtx, in.AssociatedInsights, Recommendation_InsightReference_FromProto)
	out.XorGroupID = direct.LazyPtr(in.GetXorGroupId())
	return out
}
func Recommendation_ToProto(mapCtx *direct.MapContext, in *krm.Recommendation) *pb.Recommendation {
	if in == nil {
		return nil
	}
	out := &pb.Recommendation{}
	out.Name = direct.ValueOf(in.Name)
	out.Description = direct.ValueOf(in.Description)
	out.RecommenderSubtype = direct.ValueOf(in.RecommenderSubtype)
	out.LastRefreshTime = direct.StringTimestamp_ToProto(mapCtx, in.LastRefreshTime)
	out.PrimaryImpact = Impact_ToProto(mapCtx, in.PrimaryImpact)
	out.AdditionalImpact = direct.Slice_ToProto(mapCtx, in.AdditionalImpact, Impact_ToProto)
	out.Priority = direct.Enum_ToProto[pb.Recommendation_Priority](mapCtx, in.Priority)
	out.Content = RecommendationContent_ToProto(mapCtx, in.Content)
	out.StateInfo = RecommendationStateInfo_ToProto(mapCtx, in.StateInfo)
	out.Etag = direct.ValueOf(in.Etag)
	out.AssociatedInsights = direct.Slice_ToProto(mapCtx, in.AssociatedInsights, Recommendation_InsightReference_ToProto)
	out.XorGroupId = direct.ValueOf(in.XorGroupID)
	return out
}
func RecommendationContent_FromProto(mapCtx *direct.MapContext, in *pb.RecommendationContent) *krm.RecommendationContent {
	if in == nil {
		return nil
	}
	out := &krm.RecommendationContent{}
	out.OperationGroups = direct.Slice_FromProto(mapCtx, in.OperationGroups, OperationGroup_FromProto)
	out.Overview = Overview_FromProto(mapCtx, in.GetOverview())
	return out
}
func RecommendationContent_ToProto(mapCtx *direct.MapContext, in *krm.RecommendationContent) *pb.RecommendationContent {
	if in == nil {
		return nil
	}
	out := &pb.RecommendationContent{}
	out.OperationGroups = direct.Slice_ToProto(mapCtx, in.OperationGroups, OperationGroup_ToProto)
	out.Overview = Overview_ToProto(mapCtx, in.Overview)
	return out
}
func RecommendationStateInfo_FromProto(mapCtx *direct.MapContext, in *pb.RecommendationStateInfo) *krm.RecommendationStateInfo {
	if in == nil {
		return nil
	}
	out := &krm.RecommendationStateInfo{}
	out.State = direct.Enum_FromProto(mapCtx, in.GetState())
	out.StateMetadata = in.StateMetadata
	return out
}
func RecommendationStateInfo_ToProto(mapCtx *direct.MapContext, in *krm.RecommendationStateInfo) *pb.RecommendationStateInfo {
	if in == nil {
		return nil
	}
	out := &pb.RecommendationStateInfo{}
	out.State = direct.Enum_ToProto[pb.RecommendationStateInfo_State](mapCtx, in.State)
	out.StateMetadata = in.StateMetadata
	return out
}
func Recommendation_InsightReference_FromProto(mapCtx *direct.MapContext, in *pb.Recommendation_InsightReference) *krm.Recommendation_InsightReference {
	if in == nil {
		return nil
	}
	out := &krm.Recommendation_InsightReference{}
	out.Insight = direct.LazyPtr(in.GetInsight())
	return out
}
func Recommendation_InsightReference_ToProto(mapCtx *direct.MapContext, in *krm.Recommendation_InsightReference) *pb.Recommendation_InsightReference {
	if in == nil {
		return nil
	}
	out := &pb.Recommendation_InsightReference{}
	out.Insight = direct.ValueOf(in.Insight)
	return out
}
func RecommenderRecommendationObservedState_FromProto(mapCtx *direct.MapContext, in *pb.Recommendation) *krm.RecommenderRecommendationObservedState {
	if in == nil {
		return nil
	}
	out := &krm.RecommenderRecommendationObservedState{}
	// MISSING: Name
	// MISSING: Description
	// MISSING: RecommenderSubtype
	// MISSING: LastRefreshTime
	// MISSING: PrimaryImpact
	// MISSING: AdditionalImpact
	// MISSING: Priority
	// MISSING: Content
	// MISSING: StateInfo
	// MISSING: Etag
	// MISSING: AssociatedInsights
	// MISSING: XorGroupID
	return out
}
func RecommenderRecommendationObservedState_ToProto(mapCtx *direct.MapContext, in *krm.RecommenderRecommendationObservedState) *pb.Recommendation {
	if in == nil {
		return nil
	}
	out := &pb.Recommendation{}
	// MISSING: Name
	// MISSING: Description
	// MISSING: RecommenderSubtype
	// MISSING: LastRefreshTime
	// MISSING: PrimaryImpact
	// MISSING: AdditionalImpact
	// MISSING: Priority
	// MISSING: Content
	// MISSING: StateInfo
	// MISSING: Etag
	// MISSING: AssociatedInsights
	// MISSING: XorGroupID
	return out
}
func RecommenderRecommendationSpec_FromProto(mapCtx *direct.MapContext, in *pb.Recommendation) *krm.RecommenderRecommendationSpec {
	if in == nil {
		return nil
	}
	out := &krm.RecommenderRecommendationSpec{}
	// MISSING: Name
	// MISSING: Description
	// MISSING: RecommenderSubtype
	// MISSING: LastRefreshTime
	// MISSING: PrimaryImpact
	// MISSING: AdditionalImpact
	// MISSING: Priority
	// MISSING: Content
	// MISSING: StateInfo
	// MISSING: Etag
	// MISSING: AssociatedInsights
	// MISSING: XorGroupID
	return out
}
func RecommenderRecommendationSpec_ToProto(mapCtx *direct.MapContext, in *krm.RecommenderRecommendationSpec) *pb.Recommendation {
	if in == nil {
		return nil
	}
	out := &pb.Recommendation{}
	// MISSING: Name
	// MISSING: Description
	// MISSING: RecommenderSubtype
	// MISSING: LastRefreshTime
	// MISSING: PrimaryImpact
	// MISSING: AdditionalImpact
	// MISSING: Priority
	// MISSING: Content
	// MISSING: StateInfo
	// MISSING: Etag
	// MISSING: AssociatedInsights
	// MISSING: XorGroupID
	return out
}
func ReliabilityProjection_FromProto(mapCtx *direct.MapContext, in *pb.ReliabilityProjection) *krm.ReliabilityProjection {
	if in == nil {
		return nil
	}
	out := &krm.ReliabilityProjection{}
	out.Risks = direct.EnumSlice_FromProto(mapCtx, in.Risks)
	out.Details = Details_FromProto(mapCtx, in.GetDetails())
	return out
}
func ReliabilityProjection_ToProto(mapCtx *direct.MapContext, in *krm.ReliabilityProjection) *pb.ReliabilityProjection {
	if in == nil {
		return nil
	}
	out := &pb.ReliabilityProjection{}
	out.Risks = direct.EnumSlice_ToProto[pb.ReliabilityProjection_RiskType](mapCtx, in.Risks)
	out.Details = Details_ToProto(mapCtx, in.Details)
	return out
}
func SecurityProjection_FromProto(mapCtx *direct.MapContext, in *pb.SecurityProjection) *krm.SecurityProjection {
	if in == nil {
		return nil
	}
	out := &krm.SecurityProjection{}
	out.Details = Details_FromProto(mapCtx, in.GetDetails())
	return out
}
func SecurityProjection_ToProto(mapCtx *direct.MapContext, in *krm.SecurityProjection) *pb.SecurityProjection {
	if in == nil {
		return nil
	}
	out := &pb.SecurityProjection{}
	out.Details = Details_ToProto(mapCtx, in.Details)
	return out
}
func SustainabilityProjection_FromProto(mapCtx *direct.MapContext, in *pb.SustainabilityProjection) *krm.SustainabilityProjection {
	if in == nil {
		return nil
	}
	out := &krm.SustainabilityProjection{}
	out.KgCO2e = direct.LazyPtr(in.GetKgCO2e())
	out.Duration = direct.StringDuration_FromProto(mapCtx, in.GetDuration())
	return out
}
func SustainabilityProjection_ToProto(mapCtx *direct.MapContext, in *krm.SustainabilityProjection) *pb.SustainabilityProjection {
	if in == nil {
		return nil
	}
	out := &pb.SustainabilityProjection{}
	out.KgCO2e = direct.ValueOf(in.KgCO2e)
	out.Duration = direct.StringDuration_ToProto(mapCtx, in.Duration)
	return out
}
func ValueMatcher_FromProto(mapCtx *direct.MapContext, in *pb.ValueMatcher) *krm.ValueMatcher {
	if in == nil {
		return nil
	}
	out := &krm.ValueMatcher{}
	out.MatchesPattern = direct.LazyPtr(in.GetMatchesPattern())
	return out
}
func ValueMatcher_ToProto(mapCtx *direct.MapContext, in *krm.ValueMatcher) *pb.ValueMatcher {
	if in == nil {
		return nil
	}
	out := &pb.ValueMatcher{}
	if oneof := ValueMatcher_MatchesPattern_ToProto(mapCtx, in.MatchesPattern); oneof != nil {
		out.MatchVariant = oneof
	}
	return out
}
