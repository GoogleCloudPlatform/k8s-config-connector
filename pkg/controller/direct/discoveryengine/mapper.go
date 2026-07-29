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

package discoveryengine

import (
	pb "cloud.google.com/go/discoveryengine/apiv1/discoveryenginepb"
	discoveryenginepb "cloud.google.com/go/discoveryengine/apiv1beta/discoveryenginepb"
	krm "github.com/GoogleCloudPlatform/k8s-config-connector/apis/discoveryengine/v1alpha1"
	modelarmorv1alpha1 "github.com/GoogleCloudPlatform/k8s-config-connector/apis/modelarmor/v1alpha1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct"
	datepb "google.golang.org/genproto/googleapis/type/date"
	"google.golang.org/protobuf/types/known/structpb"
	apiextensionsv1 "k8s.io/apiextensions-apiserver/pkg/apis/apiextensions/v1"
)

// Override but should be unreachable.
// Would not be needed if we did a reachability analysis in our proto generation.
func StructSchema_FromProto(mapCtx *direct.MapContext, in *structpb.Struct) apiextensionsv1.JSON {
	mapCtx.NotImplemented()
	return apiextensionsv1.JSON{}
}

// Override but should be unreachable
// Would not be needed if we did a reachability analysis in our proto generation.
func StructSchema_ToProto(mapCtx *direct.MapContext, in apiextensionsv1.JSON) *structpb.Struct {
	mapCtx.NotImplemented()
	return nil
}

// Override but should be unreachable
// Would not be needed if we did a reachability analysis in our proto generation.
func Schema_JsonSchema_FromProto(mapCtx *direct.MapContext, in *structpb.Struct) *string {
	mapCtx.NotImplemented()
	return nil
}

// Override but should be unreachable
// Would not be needed if we did a reachability analysis in our proto generation.
func Schema_JsonSchema_ToProto(mapCtx *direct.MapContext, in *string) *pb.Schema_StructSchema {
	mapCtx.NotImplemented()
	return nil
}

// We have to override because of DataStoreRefs
func DiscoveryEngineEngineSpec_FromProto(mapCtx *direct.MapContext, in *pb.Engine) *krm.DiscoveryEngineEngineSpec {
	if in == nil {
		return nil
	}
	out := &krm.DiscoveryEngineEngineSpec{}
	out.ChatEngineConfig = Engine_ChatEngineConfig_v1alpha1_FromProto(mapCtx, in.GetChatEngineConfig())
	out.SearchEngineConfig = Engine_SearchEngineConfig_v1alpha1_FromProto(mapCtx, in.GetSearchEngineConfig())
	// MISSING: ChatEngineMetadata
	out.DisplayName = direct.LazyPtr(in.GetDisplayName())
	out.SolutionType = direct.Enum_FromProto(mapCtx, in.GetSolutionType())
	out.IndustryVertical = direct.Enum_FromProto(mapCtx, in.GetIndustryVertical())
	out.CommonConfig = Engine_CommonConfig_v1alpha1_FromProto(mapCtx, in.GetCommonConfig())
	out.DisableAnalytics = direct.LazyPtr(in.GetDisableAnalytics())

	for _, dataStoreID := range in.DataStoreIds {
		out.DataStoreRefs = append(out.DataStoreRefs, &krm.DiscoveryEngineDataStoreRef{External: dataStoreID})
	}

	return out
}

// We have to override because of DataStoreRefs
func DiscoveryEngineEngineSpec_ToProto(mapCtx *direct.MapContext, in *krm.DiscoveryEngineEngineSpec) *pb.Engine {
	if in == nil {
		return nil
	}
	out := &pb.Engine{}
	if oneof := Engine_ChatEngineConfig_v1alpha1_ToProto(mapCtx, in.ChatEngineConfig); oneof != nil {
		out.EngineConfig = &pb.Engine_ChatEngineConfig_{ChatEngineConfig: oneof}
	}
	if oneof := Engine_SearchEngineConfig_v1alpha1_ToProto(mapCtx, in.SearchEngineConfig); oneof != nil {
		out.EngineConfig = &pb.Engine_SearchEngineConfig_{SearchEngineConfig: oneof}
	}
	// MISSING: ChatEngineMetadata
	out.DisplayName = direct.ValueOf(in.DisplayName)
	out.SolutionType = direct.Enum_ToProto[pb.SolutionType](mapCtx, in.SolutionType)
	out.IndustryVertical = direct.Enum_ToProto[pb.IndustryVertical](mapCtx, in.IndustryVertical)
	out.CommonConfig = Engine_CommonConfig_v1alpha1_ToProto(mapCtx, in.CommonConfig)
	out.DisableAnalytics = direct.ValueOf(in.DisableAnalytics)

	for _, dataStoreRef := range in.DataStoreRefs {
		out.DataStoreIds = append(out.DataStoreIds, dataStoreRef.External)
	}

	return out
}

func SearchResponse_Summary_v1alpha1_FromProto(mapCtx *direct.MapContext, in *pb.SearchResponse_Summary) *krm.SearchResponse_Summary {
	if in == nil {
		return nil
	}
	out := &krm.SearchResponse_Summary{}
	out.SummaryText = direct.LazyPtr(in.GetSummaryText())
	out.SummarySkippedReasons = direct.EnumSlice_FromProto(mapCtx, in.SummarySkippedReasons)
	return out
}

func SearchResponse_Summary_v1alpha1_ToProto(mapCtx *direct.MapContext, in *krm.SearchResponse_Summary) *pb.SearchResponse_Summary {
	if in == nil {
		return nil
	}
	out := &pb.SearchResponse_Summary{}
	out.SummaryText = direct.ValueOf(in.SummaryText)
	out.SummarySkippedReasons = direct.EnumSlice_ToProto[pb.SearchResponse_Summary_SummarySkippedReason](mapCtx, in.SummarySkippedReasons)
	return out
}

func Date_v1alpha1_FromProto(mapCtx *direct.MapContext, in *datepb.Date) *krm.Date {
	if in == nil {
		return nil
	}
	out := &krm.Date{}
	out.Year = direct.LazyPtr(in.GetYear())
	out.Month = direct.LazyPtr(in.GetMonth())
	out.Day = direct.LazyPtr(in.GetDay())
	return out
}

func Date_v1alpha1_ToProto(mapCtx *direct.MapContext, in *krm.Date) *datepb.Date {
	if in == nil {
		return nil
	}
	out := &datepb.Date{}
	out.Year = direct.ValueOf(in.Year)
	out.Month = direct.ValueOf(in.Month)
	out.Day = direct.ValueOf(in.Day)
	return out
}

func AssistantGroundedContent_TextGroundingMetadata_v1alpha1_FromProto(mapCtx *direct.MapContext, in *pb.AssistantGroundedContent_TextGroundingMetadata) *krm.AssistantGroundedContent_TextGroundingMetadata {
	if in == nil {
		return nil
	}
	out := &krm.AssistantGroundedContent_TextGroundingMetadata{}
	out.Segments = direct.Slice_FromProto(mapCtx, in.Segments, AssistantGroundedContent_TextGroundingMetadata_Segment_v1alpha1_FromProto)
	out.References = direct.Slice_FromProto(mapCtx, in.References, AssistantGroundedContent_TextGroundingMetadata_Reference_v1alpha1_FromProto)
	return out
}

func AssistantGroundedContent_TextGroundingMetadata_v1alpha1_ToProto(mapCtx *direct.MapContext, in *krm.AssistantGroundedContent_TextGroundingMetadata) *pb.AssistantGroundedContent_TextGroundingMetadata {
	if in == nil {
		return nil
	}
	out := &pb.AssistantGroundedContent_TextGroundingMetadata{}
	out.Segments = direct.Slice_ToProto(mapCtx, in.Segments, AssistantGroundedContent_TextGroundingMetadata_Segment_v1alpha1_ToProto)
	out.References = direct.Slice_ToProto(mapCtx, in.References, AssistantGroundedContent_TextGroundingMetadata_Reference_v1alpha1_ToProto)
	return out
}

func AssistantGroundedContent_TextGroundingMetadata_Segment_v1alpha1_FromProto(mapCtx *direct.MapContext, in *pb.AssistantGroundedContent_TextGroundingMetadata_Segment) *krm.AssistantGroundedContent_TextGroundingMetadata_Segment {
	if in == nil {
		return nil
	}
	out := &krm.AssistantGroundedContent_TextGroundingMetadata_Segment{}
	out.StartIndex = direct.LazyPtr(in.GetStartIndex())
	out.EndIndex = direct.LazyPtr(in.GetEndIndex())
	out.ReferenceIndices = in.GetReferenceIndices()
	out.GroundingScore = direct.LazyPtr(in.GetGroundingScore())
	out.Text = direct.LazyPtr(in.GetText())
	return out
}

func AssistantGroundedContent_TextGroundingMetadata_Segment_v1alpha1_ToProto(mapCtx *direct.MapContext, in *krm.AssistantGroundedContent_TextGroundingMetadata_Segment) *pb.AssistantGroundedContent_TextGroundingMetadata_Segment {
	if in == nil {
		return nil
	}
	out := &pb.AssistantGroundedContent_TextGroundingMetadata_Segment{}
	out.StartIndex = direct.ValueOf(in.StartIndex)
	out.EndIndex = direct.ValueOf(in.EndIndex)
	out.ReferenceIndices = in.ReferenceIndices
	out.GroundingScore = direct.ValueOf(in.GroundingScore)
	out.Text = direct.ValueOf(in.Text)
	return out
}

func AssistantGroundedContent_TextGroundingMetadata_Reference_v1alpha1_FromProto(mapCtx *direct.MapContext, in *pb.AssistantGroundedContent_TextGroundingMetadata_Reference) *krm.AssistantGroundedContent_TextGroundingMetadata_Reference {
	if in == nil {
		return nil
	}
	out := &krm.AssistantGroundedContent_TextGroundingMetadata_Reference{}
	out.Content = direct.LazyPtr(in.GetContent())
	out.DocumentMetadata = AssistantGroundedContent_TextGroundingMetadata_Reference_DocumentMetadata_v1alpha1_FromProto(mapCtx, in.GetDocumentMetadata())
	return out
}

func AssistantGroundedContent_TextGroundingMetadata_Reference_v1alpha1_ToProto(mapCtx *direct.MapContext, in *krm.AssistantGroundedContent_TextGroundingMetadata_Reference) *pb.AssistantGroundedContent_TextGroundingMetadata_Reference {
	if in == nil {
		return nil
	}
	out := &pb.AssistantGroundedContent_TextGroundingMetadata_Reference{}
	out.Content = direct.ValueOf(in.Content)
	out.DocumentMetadata = AssistantGroundedContent_TextGroundingMetadata_Reference_DocumentMetadata_v1alpha1_ToProto(mapCtx, in.DocumentMetadata)
	return out
}

func AssistantGroundedContent_TextGroundingMetadata_Reference_DocumentMetadata_v1alpha1_FromProto(mapCtx *direct.MapContext, in *pb.AssistantGroundedContent_TextGroundingMetadata_Reference_DocumentMetadata) *krm.AssistantGroundedContent_TextGroundingMetadata_Reference_DocumentMetadata {
	if in == nil {
		return nil
	}
	out := &krm.AssistantGroundedContent_TextGroundingMetadata_Reference_DocumentMetadata{}
	out.Document = direct.LazyPtr(in.GetDocument())
	out.URI = direct.LazyPtr(in.GetUri())
	out.Title = direct.LazyPtr(in.GetTitle())
	out.PageIdentifier = direct.LazyPtr(in.GetPageIdentifier())
	out.Domain = direct.LazyPtr(in.GetDomain())
	return out
}

func AssistantGroundedContent_TextGroundingMetadata_Reference_DocumentMetadata_v1alpha1_ToProto(mapCtx *direct.MapContext, in *krm.AssistantGroundedContent_TextGroundingMetadata_Reference_DocumentMetadata) *pb.AssistantGroundedContent_TextGroundingMetadata_Reference_DocumentMetadata {
	if in == nil {
		return nil
	}
	out := &pb.AssistantGroundedContent_TextGroundingMetadata_Reference_DocumentMetadata{}
	out.Document = in.Document
	out.Uri = in.URI
	out.Title = in.Title
	out.PageIdentifier = in.PageIdentifier
	out.Domain = in.Domain
	return out
}

func SearchRequest_PersonalizationSpec_v1alpha1_FromProto(mapCtx *direct.MapContext, in *discoveryenginepb.SearchRequest_PersonalizationSpec) *krm.SearchRequest_PersonalizationSpec {
	if in == nil {
		return nil
	}
	out := &krm.SearchRequest_PersonalizationSpec{}
	out.Mode = direct.Enum_FromProto(mapCtx, in.GetMode())
	return out
}

func SearchRequest_PersonalizationSpec_v1alpha1_ToProto(mapCtx *direct.MapContext, in *krm.SearchRequest_PersonalizationSpec) *discoveryenginepb.SearchRequest_PersonalizationSpec {
	if in == nil {
		return nil
	}
	out := &discoveryenginepb.SearchRequest_PersonalizationSpec{}
	out.Mode = direct.Enum_ToProto[discoveryenginepb.SearchRequest_PersonalizationSpec_Mode](mapCtx, in.Mode)
	return out
}

func SearchRequest_ContentSearchSpec_v1alpha1_FromProto(mapCtx *direct.MapContext, in *discoveryenginepb.SearchRequest_ContentSearchSpec) *krm.SearchRequest_ContentSearchSpec {
	if in == nil {
		return nil
	}
	out := &krm.SearchRequest_ContentSearchSpec{}
	out.SnippetSpec = SearchRequest_ContentSearchSpec_SnippetSpec_v1alpha1_FromProto(mapCtx, in.GetSnippetSpec())
	out.SummarySpec = SearchRequest_ContentSearchSpec_SummarySpec_v1alpha1_FromProto(mapCtx, in.GetSummarySpec())
	out.ExtractiveContentSpec = SearchRequest_ContentSearchSpec_ExtractiveContentSpec_v1alpha1_FromProto(mapCtx, in.GetExtractiveContentSpec())
	out.SearchResultMode = direct.Enum_FromProto(mapCtx, in.GetSearchResultMode())
	out.ChunkSpec = SearchRequest_ContentSearchSpec_ChunkSpec_v1alpha1_FromProto(mapCtx, in.GetChunkSpec())
	return out
}

func SearchRequest_ContentSearchSpec_v1alpha1_ToProto(mapCtx *direct.MapContext, in *krm.SearchRequest_ContentSearchSpec) *discoveryenginepb.SearchRequest_ContentSearchSpec {
	if in == nil {
		return nil
	}
	out := &discoveryenginepb.SearchRequest_ContentSearchSpec{}
	out.SnippetSpec = SearchRequest_ContentSearchSpec_SnippetSpec_v1alpha1_ToProto(mapCtx, in.SnippetSpec)
	out.SummarySpec = SearchRequest_ContentSearchSpec_SummarySpec_v1alpha1_ToProto(mapCtx, in.SummarySpec)
	out.ExtractiveContentSpec = SearchRequest_ContentSearchSpec_ExtractiveContentSpec_v1alpha1_ToProto(mapCtx, in.ExtractiveContentSpec)
	out.SearchResultMode = direct.Enum_ToProto[discoveryenginepb.SearchRequest_ContentSearchSpec_SearchResultMode](mapCtx, in.SearchResultMode)
	out.ChunkSpec = SearchRequest_ContentSearchSpec_ChunkSpec_v1alpha1_ToProto(mapCtx, in.ChunkSpec)
	return out
}

func SearchRequest_ContentSearchSpec_SnippetSpec_v1alpha1_FromProto(mapCtx *direct.MapContext, in *discoveryenginepb.SearchRequest_ContentSearchSpec_SnippetSpec) *krm.SearchRequest_ContentSearchSpec_SnippetSpec {
	if in == nil {
		return nil
	}
	out := &krm.SearchRequest_ContentSearchSpec_SnippetSpec{}
	out.MaxSnippetCount = direct.LazyPtr(in.GetMaxSnippetCount())
	out.ReferenceOnly = direct.LazyPtr(in.GetReferenceOnly())
	out.ReturnSnippet = direct.LazyPtr(in.GetReturnSnippet())
	return out
}

func SearchRequest_ContentSearchSpec_SnippetSpec_v1alpha1_ToProto(mapCtx *direct.MapContext, in *krm.SearchRequest_ContentSearchSpec_SnippetSpec) *discoveryenginepb.SearchRequest_ContentSearchSpec_SnippetSpec {
	if in == nil {
		return nil
	}
	out := &discoveryenginepb.SearchRequest_ContentSearchSpec_SnippetSpec{}
	out.MaxSnippetCount = direct.ValueOf(in.MaxSnippetCount)
	out.ReferenceOnly = direct.ValueOf(in.ReferenceOnly)
	out.ReturnSnippet = direct.ValueOf(in.ReturnSnippet)
	return out
}

func SearchRequest_ContentSearchSpec_SummarySpec_v1alpha1_FromProto(mapCtx *direct.MapContext, in *discoveryenginepb.SearchRequest_ContentSearchSpec_SummarySpec) *krm.SearchRequest_ContentSearchSpec_SummarySpec {
	if in == nil {
		return nil
	}
	out := &krm.SearchRequest_ContentSearchSpec_SummarySpec{}
	out.SummaryResultCount = direct.LazyPtr(in.GetSummaryResultCount())
	out.IncludeCitations = direct.LazyPtr(in.GetIncludeCitations())
	out.IgnoreAdversarialQuery = direct.LazyPtr(in.GetIgnoreAdversarialQuery())
	out.IgnoreNonSummarySeekingQuery = direct.LazyPtr(in.GetIgnoreNonSummarySeekingQuery())
	out.IgnoreLowRelevantContent = direct.LazyPtr(in.GetIgnoreLowRelevantContent())
	out.IgnoreJailBreakingQuery = direct.LazyPtr(in.GetIgnoreJailBreakingQuery())
	out.ModelPromptSpec = SearchRequest_ContentSearchSpec_SummarySpec_ModelPromptSpec_v1alpha1_FromProto(mapCtx, in.GetModelPromptSpec())
	out.LanguageCode = direct.LazyPtr(in.GetLanguageCode())
	out.ModelSpec = SearchRequest_ContentSearchSpec_SummarySpec_ModelSpec_v1alpha1_FromProto(mapCtx, in.GetModelSpec())
	out.UseSemanticChunks = direct.LazyPtr(in.GetUseSemanticChunks())
	return out
}

func SearchRequest_ContentSearchSpec_SummarySpec_v1alpha1_ToProto(mapCtx *direct.MapContext, in *krm.SearchRequest_ContentSearchSpec_SummarySpec) *discoveryenginepb.SearchRequest_ContentSearchSpec_SummarySpec {
	if in == nil {
		return nil
	}
	out := &discoveryenginepb.SearchRequest_ContentSearchSpec_SummarySpec{}
	out.SummaryResultCount = direct.ValueOf(in.SummaryResultCount)
	out.IncludeCitations = direct.ValueOf(in.IncludeCitations)
	out.IgnoreAdversarialQuery = direct.ValueOf(in.IgnoreAdversarialQuery)
	out.IgnoreNonSummarySeekingQuery = direct.ValueOf(in.IgnoreNonSummarySeekingQuery)
	out.IgnoreLowRelevantContent = direct.ValueOf(in.IgnoreLowRelevantContent)
	out.IgnoreJailBreakingQuery = direct.ValueOf(in.IgnoreJailBreakingQuery)
	out.ModelPromptSpec = SearchRequest_ContentSearchSpec_SummarySpec_ModelPromptSpec_v1alpha1_ToProto(mapCtx, in.ModelPromptSpec)
	out.LanguageCode = direct.ValueOf(in.LanguageCode)
	out.ModelSpec = SearchRequest_ContentSearchSpec_SummarySpec_ModelSpec_v1alpha1_ToProto(mapCtx, in.ModelSpec)
	out.UseSemanticChunks = direct.ValueOf(in.UseSemanticChunks)
	return out
}

func SearchRequest_ContentSearchSpec_SummarySpec_ModelPromptSpec_v1alpha1_FromProto(mapCtx *direct.MapContext, in *discoveryenginepb.SearchRequest_ContentSearchSpec_SummarySpec_ModelPromptSpec) *krm.SearchRequest_ContentSearchSpec_SummarySpec_ModelPromptSpec {
	if in == nil {
		return nil
	}
	out := &krm.SearchRequest_ContentSearchSpec_SummarySpec_ModelPromptSpec{}
	out.Preamble = direct.LazyPtr(in.GetPreamble())
	return out
}

func SearchRequest_ContentSearchSpec_SummarySpec_ModelPromptSpec_v1alpha1_ToProto(mapCtx *direct.MapContext, in *krm.SearchRequest_ContentSearchSpec_SummarySpec_ModelPromptSpec) *discoveryenginepb.SearchRequest_ContentSearchSpec_SummarySpec_ModelPromptSpec {
	if in == nil {
		return nil
	}
	out := &discoveryenginepb.SearchRequest_ContentSearchSpec_SummarySpec_ModelPromptSpec{}
	out.Preamble = direct.ValueOf(in.Preamble)
	return out
}

func SearchRequest_ContentSearchSpec_SummarySpec_ModelSpec_v1alpha1_FromProto(mapCtx *direct.MapContext, in *discoveryenginepb.SearchRequest_ContentSearchSpec_SummarySpec_ModelSpec) *krm.SearchRequest_ContentSearchSpec_SummarySpec_ModelSpec {
	if in == nil {
		return nil
	}
	out := &krm.SearchRequest_ContentSearchSpec_SummarySpec_ModelSpec{}
	out.Version = direct.LazyPtr(in.GetVersion())
	return out
}

func SearchRequest_ContentSearchSpec_SummarySpec_ModelSpec_v1alpha1_ToProto(mapCtx *direct.MapContext, in *krm.SearchRequest_ContentSearchSpec_SummarySpec_ModelSpec) *discoveryenginepb.SearchRequest_ContentSearchSpec_SummarySpec_ModelSpec {
	if in == nil {
		return nil
	}
	out := &discoveryenginepb.SearchRequest_ContentSearchSpec_SummarySpec_ModelSpec{}
	out.Version = direct.ValueOf(in.Version)
	return out
}

func SearchRequest_ContentSearchSpec_ExtractiveContentSpec_v1alpha1_FromProto(mapCtx *direct.MapContext, in *discoveryenginepb.SearchRequest_ContentSearchSpec_ExtractiveContentSpec) *krm.SearchRequest_ContentSearchSpec_ExtractiveContentSpec {
	if in == nil {
		return nil
	}
	out := &krm.SearchRequest_ContentSearchSpec_ExtractiveContentSpec{}
	out.MaxExtractiveAnswerCount = direct.LazyPtr(in.GetMaxExtractiveAnswerCount())
	out.MaxExtractiveSegmentCount = direct.LazyPtr(in.GetMaxExtractiveSegmentCount())
	out.ReturnExtractiveSegmentScore = direct.LazyPtr(in.GetReturnExtractiveSegmentScore())
	out.NumPreviousSegments = direct.LazyPtr(in.GetNumPreviousSegments())
	out.NumNextSegments = direct.LazyPtr(in.GetNumNextSegments())
	return out
}

func SearchRequest_ContentSearchSpec_ExtractiveContentSpec_v1alpha1_ToProto(mapCtx *direct.MapContext, in *krm.SearchRequest_ContentSearchSpec_ExtractiveContentSpec) *discoveryenginepb.SearchRequest_ContentSearchSpec_ExtractiveContentSpec {
	if in == nil {
		return nil
	}
	out := &discoveryenginepb.SearchRequest_ContentSearchSpec_ExtractiveContentSpec{}
	out.MaxExtractiveAnswerCount = direct.ValueOf(in.MaxExtractiveAnswerCount)
	out.MaxExtractiveSegmentCount = direct.ValueOf(in.MaxExtractiveSegmentCount)
	out.ReturnExtractiveSegmentScore = direct.ValueOf(in.ReturnExtractiveSegmentScore)
	out.NumPreviousSegments = direct.ValueOf(in.NumPreviousSegments)
	out.NumNextSegments = direct.ValueOf(in.NumNextSegments)
	return out
}

func SearchRequest_ContentSearchSpec_ChunkSpec_v1alpha1_FromProto(mapCtx *direct.MapContext, in *discoveryenginepb.SearchRequest_ContentSearchSpec_ChunkSpec) *krm.SearchRequest_ContentSearchSpec_ChunkSpec {
	if in == nil {
		return nil
	}
	out := &krm.SearchRequest_ContentSearchSpec_ChunkSpec{}
	out.NumPreviousChunks = direct.LazyPtr(in.GetNumPreviousChunks())
	out.NumNextChunks = direct.LazyPtr(in.GetNumNextChunks())
	return out
}

func SearchRequest_ContentSearchSpec_ChunkSpec_v1alpha1_ToProto(mapCtx *direct.MapContext, in *krm.SearchRequest_ContentSearchSpec_ChunkSpec) *discoveryenginepb.SearchRequest_ContentSearchSpec_ChunkSpec {
	if in == nil {
		return nil
	}
	out := &discoveryenginepb.SearchRequest_ContentSearchSpec_ChunkSpec{}
	out.NumPreviousChunks = direct.ValueOf(in.NumPreviousChunks)
	out.NumNextChunks = direct.ValueOf(in.NumNextChunks)
	return out
}

func DiscoveryEngineServingConfigSpec_v1alpha1_FromProto(mapCtx *direct.MapContext, in *discoveryenginepb.ServingConfig) *krm.DiscoveryEngineServingConfigSpec {
	if in == nil {
		return nil
	}
	out := &krm.DiscoveryEngineServingConfigSpec{}
	out.MediaConfig = ServingConfig_MediaConfig_v1alpha1_FromProto(mapCtx, in.GetMediaConfig())
	out.GenericConfig = ServingConfig_GenericConfig_v1alpha1_FromProto(mapCtx, in.GetGenericConfig())
	// MISSING: Name
	out.DisplayName = direct.LazyPtr(in.GetDisplayName())
	out.SolutionType = direct.Enum_FromProto(mapCtx, in.GetSolutionType())
	out.ModelID = direct.LazyPtr(in.GetModelId())
	out.DiversityLevel = direct.LazyPtr(in.GetDiversityLevel())
	out.EmbeddingConfig = EmbeddingConfig_v1alpha1_FromProto(mapCtx, in.GetEmbeddingConfig())
	out.RankingExpression = direct.LazyPtr(in.GetRankingExpression())
	out.FilterControlIDs = in.GetFilterControlIds()
	out.BoostControlIDs = in.GetBoostControlIds()
	out.RedirectControlIDs = in.GetRedirectControlIds()
	out.SynonymsControlIDs = in.GetSynonymsControlIds()
	out.OnewaySynonymsControlIDs = in.GetOnewaySynonymsControlIds()
	out.DissociateControlIDs = in.GetDissociateControlIds()
	out.ReplacementControlIDs = in.GetReplacementControlIds()
	out.IgnoreControlIDs = in.GetIgnoreControlIds()
	out.PersonalizationSpec = SearchRequest_PersonalizationSpec_v1alpha1_FromProto(mapCtx, in.GetPersonalizationSpec())
	return out
}

func DiscoveryEngineServingConfigSpec_v1alpha1_ToProto(mapCtx *direct.MapContext, in *krm.DiscoveryEngineServingConfigSpec) *discoveryenginepb.ServingConfig {
	if in == nil {
		return nil
	}
	out := &discoveryenginepb.ServingConfig{}
	if oneof := ServingConfig_MediaConfig_v1alpha1_ToProto(mapCtx, in.MediaConfig); oneof != nil {
		out.VerticalConfig = &discoveryenginepb.ServingConfig_MediaConfig_{MediaConfig: oneof}
	}
	if oneof := ServingConfig_GenericConfig_v1alpha1_ToProto(mapCtx, in.GenericConfig); oneof != nil {
		out.VerticalConfig = &discoveryenginepb.ServingConfig_GenericConfig_{GenericConfig: oneof}
	}
	// MISSING: Name
	out.DisplayName = direct.ValueOf(in.DisplayName)
	out.SolutionType = direct.Enum_ToProto[discoveryenginepb.SolutionType](mapCtx, in.SolutionType)
	out.ModelId = direct.ValueOf(in.ModelID)
	out.DiversityLevel = direct.ValueOf(in.DiversityLevel)
	out.EmbeddingConfig = EmbeddingConfig_v1alpha1_ToProto(mapCtx, in.EmbeddingConfig)
	out.RankingExpression = direct.ValueOf(in.RankingExpression)
	out.FilterControlIds = in.FilterControlIDs
	out.BoostControlIds = in.BoostControlIDs
	out.RedirectControlIds = in.RedirectControlIDs
	out.SynonymsControlIds = in.SynonymsControlIDs
	out.OnewaySynonymsControlIds = in.OnewaySynonymsControlIDs
	out.DissociateControlIds = in.DissociateControlIDs
	out.ReplacementControlIds = in.ReplacementControlIDs
	out.IgnoreControlIds = in.IgnoreControlIDs
	out.PersonalizationSpec = SearchRequest_PersonalizationSpec_v1alpha1_ToProto(mapCtx, in.PersonalizationSpec)
	return out
}

func EnabledTools_FromProto(mapCtx *direct.MapContext, in map[string]*discoveryenginepb.Assistant_ToolList) map[string]krm.Assistant_ToolList {
	if in == nil {
		return nil
	}
	out := make(map[string]krm.Assistant_ToolList)
	for k, v := range in {
		val := Assistant_ToolList_v1alpha1_FromProto(mapCtx, v)
		if val != nil {
			out[k] = *val
		}
	}
	return out
}

func EnabledTools_ToProto(mapCtx *direct.MapContext, in map[string]krm.Assistant_ToolList) map[string]*discoveryenginepb.Assistant_ToolList {
	if in == nil {
		return nil
	}
	out := make(map[string]*discoveryenginepb.Assistant_ToolList)
	for k, v := range in {
		out[k] = Assistant_ToolList_v1alpha1_ToProto(mapCtx, &v)
	}
	return out
}

func Assistant_CustomerPolicy_ModelArmorConfig_v1alpha1_FromProto(mapCtx *direct.MapContext, in *discoveryenginepb.Assistant_CustomerPolicy_ModelArmorConfig) *krm.Assistant_CustomerPolicy_ModelArmorConfig {
	if in == nil {
		return nil
	}
	out := &krm.Assistant_CustomerPolicy_ModelArmorConfig{}
	if in.GetUserPromptTemplate() != "" {
		out.UserPromptTemplateRef = &modelarmorv1alpha1.ModelArmorTemplateRef{External: in.GetUserPromptTemplate()}
	}
	if in.GetResponseTemplate() != "" {
		out.ResponseTemplateRef = &modelarmorv1alpha1.ModelArmorTemplateRef{External: in.GetResponseTemplate()}
	}
	out.FailureMode = direct.Enum_FromProto(mapCtx, in.GetFailureMode())
	return out
}

func Assistant_CustomerPolicy_ModelArmorConfig_v1alpha1_ToProto(mapCtx *direct.MapContext, in *krm.Assistant_CustomerPolicy_ModelArmorConfig) *discoveryenginepb.Assistant_CustomerPolicy_ModelArmorConfig {
	if in == nil {
		return nil
	}
	out := &discoveryenginepb.Assistant_CustomerPolicy_ModelArmorConfig{}
	if in.UserPromptTemplateRef != nil {
		out.UserPromptTemplate = in.UserPromptTemplateRef.External
	}
	if in.ResponseTemplateRef != nil {
		out.ResponseTemplate = in.ResponseTemplateRef.External
	}
	out.FailureMode = direct.Enum_ToProto[discoveryenginepb.Assistant_CustomerPolicy_ModelArmorConfig_FailureMode](mapCtx, in.FailureMode)
	return out
}

func Assistant_GenerationConfig_v1alpha1_FromProto(mapCtx *direct.MapContext, in *discoveryenginepb.Assistant_GenerationConfig) *krm.Assistant_GenerationConfig {
	if in == nil {
		return nil
	}
	out := &krm.Assistant_GenerationConfig{}
	out.DefaultModelID = direct.LazyPtr(in.GetDefaultModelId())
	out.AllowedModelIDs = in.GetAllowedModelIds()
	out.SystemInstruction = Assistant_GenerationConfig_SystemInstruction_v1alpha1_FromProto(mapCtx, in.GetSystemInstruction())
	out.DefaultLanguage = direct.LazyPtr(in.GetDefaultLanguage())
	return out
}

func Assistant_GenerationConfig_v1alpha1_ToProto(mapCtx *direct.MapContext, in *krm.Assistant_GenerationConfig) *discoveryenginepb.Assistant_GenerationConfig {
	if in == nil {
		return nil
	}
	out := &discoveryenginepb.Assistant_GenerationConfig{}
	out.DefaultModelId = direct.ValueOf(in.DefaultModelID)
	out.AllowedModelIds = in.AllowedModelIDs
	out.SystemInstruction = Assistant_GenerationConfig_SystemInstruction_v1alpha1_ToProto(mapCtx, in.SystemInstruction)
	out.DefaultLanguage = direct.ValueOf(in.DefaultLanguage)
	return out
}
