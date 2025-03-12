// Copyright 2024 Google LLC
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

// +tool:fuzz-gen
// proto.message: google.cloud.discoveryengine.v1beta.ServingConfig
// api.group: discoveryengine.cnrm.cloud.google.com

package discoveryengine

import (
	pb "cloud.google.com/go/discoveryengine/apiv1beta/discoveryenginepb"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/fuzztesting"
)

func init() {
	fuzztesting.RegisterKRMFuzzer(discoveryEngineServingConfigFuzzer())
}

func discoveryEngineServingConfigFuzzer() fuzztesting.KRMFuzzer {
	f := fuzztesting.NewKRMTypedFuzzer(&pb.ServingConfig{},
		DiscoveryEngineServingConfigSpec_FromProto, DiscoveryEngineServingConfigSpec_ToProto,
		DiscoveryEngineServingConfigObservedState_FromProto, DiscoveryEngineServingConfigObservedState_ToProto,
	)

	f.UnimplementedFields.Insert(".name") // special field

	f.SpecFields.Insert(".media_config")
	f.SpecFields.Insert(".generic_config")
	f.SpecFields.Insert(".display_name")
	f.SpecFields.Insert(".solution_type")
	f.SpecFields.Insert(".model_id")
	f.SpecFields.Insert(".diversity_level")
	f.SpecFields.Insert(".embedding_config")
	f.SpecFields.Insert(".ranking_expression")
	f.SpecFields.Insert(".filter_control_ids")
	f.SpecFields.Insert(".boost_control_ids")
	f.SpecFields.Insert(".redirect_control_ids")
	f.SpecFields.Insert(".synonyms_control_ids")
	f.SpecFields.Insert(".oneway_synonyms_control_ids")
	f.SpecFields.Insert(".dissociate_control_ids")
	f.SpecFields.Insert(".replacement_control_ids")
	f.SpecFields.Insert(".ignore_control_ids")
	f.SpecFields.Insert(".personalization_spec")

	// Fields we don't want to implement (yet) because they are very volatile
	f.UnimplementedFields.Insert(".create_time")
	f.UnimplementedFields.Insert(".update_time")

	return f
}

func DiscoveryEngineServingConfigSpec_FromProto(mapCtx *direct.MapContext, in *pb.ServingConfig) *krm.DiscoveryEngineServingConfigSpec {
	if in == nil {
		return nil
	}
	out := &krm.DiscoveryEngineServingConfigSpec{}
	out.MediaConfig = ServingConfig_MediaConfig_FromProto(mapCtx, in.GetMediaConfig())
	out.GenericConfig = ServingConfig_GenericConfig_FromProto(mapCtx, in.GetGenericConfig())
	out.Name = direct.LazyPtr(in.GetName())
	out.DisplayName = direct.LazyPtr(in.GetDisplayName())
	out.SolutionType = direct.Enum_FromProto(mapCtx, in.GetSolutionType())
	out.ModelID = direct.LazyPtr(in.GetModelId())
	out.DiversityLevel = direct.LazyPtr(in.GetDiversityLevel())
	out.EmbeddingConfig = EmbeddingConfig_FromProto(mapCtx, in.GetEmbeddingConfig())
	out.RankingExpression = direct.LazyPtr(in.GetRankingExpression())
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
func DiscoveryEngineServingConfigSpec_ToProto(mapCtx *direct.MapContext, in *krm.DiscoveryEngineServingConfigSpec) *pb.ServingConfig {
	if in == nil {
		return nil
	}
	out := &pb.ServingConfig{}
	if oneof := ServingConfig_MediaConfig_ToProto(mapCtx, in.MediaConfig); oneof != nil {
		out.VerticalConfig = &pb.ServingConfig_MediaConfig{MediaConfig: oneof}
	}
	if oneof := ServingConfig_GenericConfig_ToProto(mapCtx, in.GenericConfig); oneof != nil {
		out.VerticalConfig = &pb.ServingConfig_GenericConfig{GenericConfig: oneof}
	}
	out.Name = direct.ValueOf(in.Name)
	out.DisplayName = direct.ValueOf(in.DisplayName)
	out.SolutionType = direct.Enum_ToProto[pb.SolutionType](mapCtx, in.SolutionType)
	out.ModelId = direct.ValueOf(in.ModelID)
	out.DiversityLevel = direct.ValueOf(in.DiversityLevel)
	out.EmbeddingConfig = EmbeddingConfig_ToProto(mapCtx, in.EmbeddingConfig)
	out.RankingExpression = direct.ValueOf(in.RankingExpression)
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
func DiscoveryEngineServingConfigObservedState_FromProto(mapCtx *direct.MapContext, in *pb.ServingConfig) *krm.DiscoveryEngineServingConfigObservedState {
	if in == nil {
		return nil
	}
	out := &krm.DiscoveryEngineServingConfigObservedState{}
	out.CreateTime = direct.StringTimestamp_FromProto(mapCtx, in.GetCreateTime())
	out.UpdateTime = direct.StringTimestamp_FromProto(mapCtx, in.GetUpdateTime())
	return out
}
func DiscoveryEngineServingConfigObservedState_ToProto(mapCtx *direct.MapContext, in *krm.DiscoveryEngineServingConfigObservedState) *pb.ServingConfig {
	if in == nil {
		return nil
	}
	out := &pb.ServingConfig{}
	out.CreateTime = direct.StringTimestamp_ToProto(mapCtx, in.CreateTime)
	out.UpdateTime = direct.StringTimestamp_ToProto(mapCtx, in.UpdateTime)
	return out
}
