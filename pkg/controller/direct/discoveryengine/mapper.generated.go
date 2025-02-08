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

package discoveryengine

import (
	pb "cloud.google.com/go/discoveryengine/apiv1beta/discoveryenginepb"
	krm "github.com/GoogleCloudPlatform/k8s-config-connector/apis/discoveryengine/v1alpha1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct"
	refs "github.com/GoogleCloudPlatform/k8s-config-connector/apis/refs/v1beta1"
)
func DiscoveryengineServingConfigObservedState_FromProto(mapCtx *direct.MapContext, in *pb.ServingConfig) *krm.DiscoveryengineServingConfigObservedState {
	if in == nil {
		return nil
	}
	out := &krm.DiscoveryengineServingConfigObservedState{}
	// MISSING: MediaConfig
	// MISSING: GenericConfig
	// MISSING: Name
	// MISSING: DisplayName
	// MISSING: SolutionType
	// MISSING: ModelID
	// MISSING: DiversityLevel
	// MISSING: EmbeddingConfig
	// MISSING: RankingExpression
	// MISSING: CreateTime
	// MISSING: UpdateTime
	// MISSING: FilterControlIds
	// MISSING: BoostControlIds
	// MISSING: RedirectControlIds
	// MISSING: SynonymsControlIds
	// MISSING: OnewaySynonymsControlIds
	// MISSING: DissociateControlIds
	// MISSING: ReplacementControlIds
	// MISSING: IgnoreControlIds
	// MISSING: PersonalizationSpec
	return out
}
func DiscoveryengineServingConfigObservedState_ToProto(mapCtx *direct.MapContext, in *krm.DiscoveryengineServingConfigObservedState) *pb.ServingConfig {
	if in == nil {
		return nil
	}
	out := &pb.ServingConfig{}
	// MISSING: MediaConfig
	// MISSING: GenericConfig
	// MISSING: Name
	// MISSING: DisplayName
	// MISSING: SolutionType
	// MISSING: ModelID
	// MISSING: DiversityLevel
	// MISSING: EmbeddingConfig
	// MISSING: RankingExpression
	// MISSING: CreateTime
	// MISSING: UpdateTime
	// MISSING: FilterControlIds
	// MISSING: BoostControlIds
	// MISSING: RedirectControlIds
	// MISSING: SynonymsControlIds
	// MISSING: OnewaySynonymsControlIds
	// MISSING: DissociateControlIds
	// MISSING: ReplacementControlIds
	// MISSING: IgnoreControlIds
	// MISSING: PersonalizationSpec
	return out
}
func DiscoveryengineServingConfigSpec_FromProto(mapCtx *direct.MapContext, in *pb.ServingConfig) *krm.DiscoveryengineServingConfigSpec {
	if in == nil {
		return nil
	}
	out := &krm.DiscoveryengineServingConfigSpec{}
	// MISSING: MediaConfig
	// MISSING: GenericConfig
	// MISSING: Name
	// MISSING: DisplayName
	// MISSING: SolutionType
	// MISSING: ModelID
	// MISSING: DiversityLevel
	// MISSING: EmbeddingConfig
	// MISSING: RankingExpression
	// MISSING: CreateTime
	// MISSING: UpdateTime
	// MISSING: FilterControlIds
	// MISSING: BoostControlIds
	// MISSING: RedirectControlIds
	// MISSING: SynonymsControlIds
	// MISSING: OnewaySynonymsControlIds
	// MISSING: DissociateControlIds
	// MISSING: ReplacementControlIds
	// MISSING: IgnoreControlIds
	// MISSING: PersonalizationSpec
	return out
}
func DiscoveryengineServingConfigSpec_ToProto(mapCtx *direct.MapContext, in *krm.DiscoveryengineServingConfigSpec) *pb.ServingConfig {
	if in == nil {
		return nil
	}
	out := &pb.ServingConfig{}
	// MISSING: MediaConfig
	// MISSING: GenericConfig
	// MISSING: Name
	// MISSING: DisplayName
	// MISSING: SolutionType
	// MISSING: ModelID
	// MISSING: DiversityLevel
	// MISSING: EmbeddingConfig
	// MISSING: RankingExpression
	// MISSING: CreateTime
	// MISSING: UpdateTime
	// MISSING: FilterControlIds
	// MISSING: BoostControlIds
	// MISSING: RedirectControlIds
	// MISSING: SynonymsControlIds
	// MISSING: OnewaySynonymsControlIds
	// MISSING: DissociateControlIds
	// MISSING: ReplacementControlIds
	// MISSING: IgnoreControlIds
	// MISSING: PersonalizationSpec
	return out
}
func EmbeddingConfig_FromProto(mapCtx *direct.MapContext, in *pb.EmbeddingConfig) *krm.EmbeddingConfig {
	if in == nil {
		return nil
	}
	out := &krm.EmbeddingConfig{}
	out.FieldPath = direct.LazyPtr(in.GetFieldPath())
	return out
}
func EmbeddingConfig_ToProto(mapCtx *direct.MapContext, in *krm.EmbeddingConfig) *pb.EmbeddingConfig {
	if in == nil {
		return nil
	}
	out := &pb.EmbeddingConfig{}
	out.FieldPath = direct.ValueOf(in.FieldPath)
	return out
}
func ServingConfig_FromProto(mapCtx *direct.MapContext, in *pb.ServingConfig) *krm.ServingConfig {
	if in == nil {
		return nil
	}
	out := &krm.ServingConfig{}
	out.MediaConfig = ServingConfig_MediaConfig_FromProto(mapCtx, in.GetMediaConfig())
	out.GenericConfig = ServingConfig_GenericConfig_FromProto(mapCtx, in.GetGenericConfig())
	out.Name = direct.LazyPtr(in.GetName())
	out.DisplayName = direct.LazyPtr(in.GetDisplayName())
	out.SolutionType = direct.Enum_FromProto(mapCtx, in.GetSolutionType())
	out.ModelID = direct.LazyPtr(in.GetModelId())
	out.DiversityLevel = direct.LazyPtr(in.GetDiversityLevel())
	out.EmbeddingConfig = EmbeddingConfig_FromProto(mapCtx, in.GetEmbeddingConfig())
	out.RankingExpression = direct.LazyPtr(in.GetRankingExpression())
	// MISSING: CreateTime
	// MISSING: UpdateTime
	out.FilterControlIds = in.FilterControlIds
	out.BoostControlIds = in.BoostControlIds
	out.RedirectControlIds = in.RedirectControlIds
	out.SynonymsControlIds = in.SynonymsControlIds
	out.OnewaySynonymsControlIds = in.OnewaySynonymsControlIds
	out.DissociateControlIds = in.DissociateControlIds
	out.ReplacementControlIds = in.ReplacementControlIds
	out.IgnoreControlIds = in.IgnoreControlIds
	out.PersonalizationSpec = SearchRequest_PersonalizationSpec_FromProto(mapCtx, in.GetPersonalizationSpec())
	return out
}
func ServingConfig_ToProto(mapCtx *direct.MapContext, in *krm.ServingConfig) *pb.ServingConfig {
	if in == nil {
		return nil
	}
	out := &pb.ServingConfig{}
	if oneof := ServingConfig_MediaConfig_ToProto(mapCtx, in.MediaConfig); oneof != nil {
		out.VerticalConfig = &pb.ServingConfig_MediaConfig_{MediaConfig: oneof}
	}
	if oneof := ServingConfig_GenericConfig_ToProto(mapCtx, in.GenericConfig); oneof != nil {
		out.VerticalConfig = &pb.ServingConfig_GenericConfig_{GenericConfig: oneof}
	}
	out.Name = direct.ValueOf(in.Name)
	out.DisplayName = direct.ValueOf(in.DisplayName)
	out.SolutionType = direct.Enum_ToProto[pb.SolutionType](mapCtx, in.SolutionType)
	out.ModelId = direct.ValueOf(in.ModelID)
	out.DiversityLevel = direct.ValueOf(in.DiversityLevel)
	out.EmbeddingConfig = EmbeddingConfig_ToProto(mapCtx, in.EmbeddingConfig)
	out.RankingExpression = direct.ValueOf(in.RankingExpression)
	// MISSING: CreateTime
	// MISSING: UpdateTime
	out.FilterControlIds = in.FilterControlIds
	out.BoostControlIds = in.BoostControlIds
	out.RedirectControlIds = in.RedirectControlIds
	out.SynonymsControlIds = in.SynonymsControlIds
	out.OnewaySynonymsControlIds = in.OnewaySynonymsControlIds
	out.DissociateControlIds = in.DissociateControlIds
	out.ReplacementControlIds = in.ReplacementControlIds
	out.IgnoreControlIds = in.IgnoreControlIds
	out.PersonalizationSpec = SearchRequest_PersonalizationSpec_ToProto(mapCtx, in.PersonalizationSpec)
	return out
}
func ServingConfigObservedState_FromProto(mapCtx *direct.MapContext, in *pb.ServingConfig) *krm.ServingConfigObservedState {
	if in == nil {
		return nil
	}
	out := &krm.ServingConfigObservedState{}
	// MISSING: MediaConfig
	// MISSING: GenericConfig
	// MISSING: Name
	// MISSING: DisplayName
	// MISSING: SolutionType
	// MISSING: ModelID
	// MISSING: DiversityLevel
	// MISSING: EmbeddingConfig
	// MISSING: RankingExpression
	out.CreateTime = direct.StringTimestamp_FromProto(mapCtx, in.GetCreateTime())
	out.UpdateTime = direct.StringTimestamp_FromProto(mapCtx, in.GetUpdateTime())
	// MISSING: FilterControlIds
	// MISSING: BoostControlIds
	// MISSING: RedirectControlIds
	// MISSING: SynonymsControlIds
	// MISSING: OnewaySynonymsControlIds
	// MISSING: DissociateControlIds
	// MISSING: ReplacementControlIds
	// MISSING: IgnoreControlIds
	// MISSING: PersonalizationSpec
	return out
}
func ServingConfigObservedState_ToProto(mapCtx *direct.MapContext, in *krm.ServingConfigObservedState) *pb.ServingConfig {
	if in == nil {
		return nil
	}
	out := &pb.ServingConfig{}
	// MISSING: MediaConfig
	// MISSING: GenericConfig
	// MISSING: Name
	// MISSING: DisplayName
	// MISSING: SolutionType
	// MISSING: ModelID
	// MISSING: DiversityLevel
	// MISSING: EmbeddingConfig
	// MISSING: RankingExpression
	out.CreateTime = direct.StringTimestamp_ToProto(mapCtx, in.CreateTime)
	out.UpdateTime = direct.StringTimestamp_ToProto(mapCtx, in.UpdateTime)
	// MISSING: FilterControlIds
	// MISSING: BoostControlIds
	// MISSING: RedirectControlIds
	// MISSING: SynonymsControlIds
	// MISSING: OnewaySynonymsControlIds
	// MISSING: DissociateControlIds
	// MISSING: ReplacementControlIds
	// MISSING: IgnoreControlIds
	// MISSING: PersonalizationSpec
	return out
}
func ServingConfig_GenericConfig_FromProto(mapCtx *direct.MapContext, in *pb.ServingConfig_GenericConfig) *krm.ServingConfig_GenericConfig {
	if in == nil {
		return nil
	}
	out := &krm.ServingConfig_GenericConfig{}
	out.ContentSearchSpec = SearchRequest_ContentSearchSpec_FromProto(mapCtx, in.GetContentSearchSpec())
	return out
}
func ServingConfig_GenericConfig_ToProto(mapCtx *direct.MapContext, in *krm.ServingConfig_GenericConfig) *pb.ServingConfig_GenericConfig {
	if in == nil {
		return nil
	}
	out := &pb.ServingConfig_GenericConfig{}
	out.ContentSearchSpec = SearchRequest_ContentSearchSpec_ToProto(mapCtx, in.ContentSearchSpec)
	return out
}
func ServingConfig_MediaConfig_FromProto(mapCtx *direct.MapContext, in *pb.ServingConfig_MediaConfig) *krm.ServingConfig_MediaConfig {
	if in == nil {
		return nil
	}
	out := &krm.ServingConfig_MediaConfig{}
	out.ContentWatchedPercentageThreshold = direct.LazyPtr(in.GetContentWatchedPercentageThreshold())
	out.ContentWatchedSecondsThreshold = direct.LazyPtr(in.GetContentWatchedSecondsThreshold())
	out.DemotionEventType = direct.LazyPtr(in.GetDemotionEventType())
	out.DemoteContentWatchedPastDays = direct.LazyPtr(in.GetDemoteContentWatchedPastDays())
	out.ContentFreshnessCutoffDays = direct.LazyPtr(in.GetContentFreshnessCutoffDays())
	return out
}
func ServingConfig_MediaConfig_ToProto(mapCtx *direct.MapContext, in *krm.ServingConfig_MediaConfig) *pb.ServingConfig_MediaConfig {
	if in == nil {
		return nil
	}
	out := &pb.ServingConfig_MediaConfig{}
	if oneof := ServingConfig_MediaConfig_ContentWatchedPercentageThreshold_ToProto(mapCtx, in.ContentWatchedPercentageThreshold); oneof != nil {
		out.DemoteContentWatched = oneof
	}
	if oneof := ServingConfig_MediaConfig_ContentWatchedSecondsThreshold_ToProto(mapCtx, in.ContentWatchedSecondsThreshold); oneof != nil {
		out.DemoteContentWatched = oneof
	}
	out.DemotionEventType = direct.ValueOf(in.DemotionEventType)
	out.DemoteContentWatchedPastDays = direct.ValueOf(in.DemoteContentWatchedPastDays)
	out.ContentFreshnessCutoffDays = direct.ValueOf(in.ContentFreshnessCutoffDays)
	return out
}
