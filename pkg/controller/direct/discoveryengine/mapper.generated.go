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
	pb_v1 "cloud.google.com/go/discoveryengine/apiv1/discoveryenginepb"
	pb_v1beta "cloud.google.com/go/discoveryengine/apiv1beta/discoveryenginepb"
	krm "github.com/GoogleCloudPlatform/k8s-config-connector/apis/discoveryengine/v1alpha1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct"
)

func DataStore_BillingEstimation_FromProto(mapCtx *direct.MapContext, in *pb_v1.DataStore_BillingEstimation) *krm.DataStore_BillingEstimation {
	if in == nil {
		return nil
	}
	out := &krm.DataStore_BillingEstimation{}
	out.StructuredDataSize = direct.LazyPtr(in.GetStructuredDataSize())
	out.UnstructuredDataSize = direct.LazyPtr(in.GetUnstructuredDataSize())
	out.WebsiteDataSize = direct.LazyPtr(in.GetWebsiteDataSize())
	out.StructuredDataUpdateTime = direct.StringTimestamp_FromProto(mapCtx, in.GetStructuredDataUpdateTime())
	out.UnstructuredDataUpdateTime = direct.StringTimestamp_FromProto(mapCtx, in.GetUnstructuredDataUpdateTime())
	out.WebsiteDataUpdateTime = direct.StringTimestamp_FromProto(mapCtx, in.GetWebsiteDataUpdateTime())
	return out
}
func DataStore_BillingEstimation_ToProto(mapCtx *direct.MapContext, in *krm.DataStore_BillingEstimation) *pb_v1.DataStore_BillingEstimation {
	if in == nil {
		return nil
	}
	out := &pb_v1.DataStore_BillingEstimation{}
	out.StructuredDataSize = direct.ValueOf(in.StructuredDataSize)
	out.UnstructuredDataSize = direct.ValueOf(in.UnstructuredDataSize)
	out.WebsiteDataSize = direct.ValueOf(in.WebsiteDataSize)
	out.StructuredDataUpdateTime = direct.StringTimestamp_ToProto(mapCtx, in.StructuredDataUpdateTime)
	out.UnstructuredDataUpdateTime = direct.StringTimestamp_ToProto(mapCtx, in.UnstructuredDataUpdateTime)
	out.WebsiteDataUpdateTime = direct.StringTimestamp_ToProto(mapCtx, in.WebsiteDataUpdateTime)
	return out
}
func DiscoveryEngineDataStoreObservedState_FromProto(mapCtx *direct.MapContext, in *pb_v1.DataStore) *krm.DiscoveryEngineDataStoreObservedState {
	if in == nil {
		return nil
	}
	out := &krm.DiscoveryEngineDataStoreObservedState{}
	// MISSING: Name
	out.DefaultSchemaID = direct.LazyPtr(in.GetDefaultSchemaId())
	out.CreateTime = direct.StringTimestamp_FromProto(mapCtx, in.GetCreateTime())
	out.BillingEstimation = DataStore_BillingEstimation_FromProto(mapCtx, in.GetBillingEstimation())
	// MISSING: DocumentProcessingConfig
	// MISSING: StartingSchema
	return out
}
func DiscoveryEngineDataStoreObservedState_ToProto(mapCtx *direct.MapContext, in *krm.DiscoveryEngineDataStoreObservedState) *pb_v1.DataStore {
	if in == nil {
		return nil
	}
	out := &pb_v1.DataStore{}
	// MISSING: Name
	out.DefaultSchemaId = direct.ValueOf(in.DefaultSchemaID)
	out.CreateTime = direct.StringTimestamp_ToProto(mapCtx, in.CreateTime)
	out.BillingEstimation = DataStore_BillingEstimation_ToProto(mapCtx, in.BillingEstimation)
	// MISSING: DocumentProcessingConfig
	// MISSING: StartingSchema
	return out
}
func DiscoveryEngineDataStoreSpec_FromProto(mapCtx *direct.MapContext, in *pb_v1.DataStore) *krm.DiscoveryEngineDataStoreSpec {
	if in == nil {
		return nil
	}
	out := &krm.DiscoveryEngineDataStoreSpec{}
	// MISSING: Name
	out.DisplayName = direct.LazyPtr(in.GetDisplayName())
	out.IndustryVertical = direct.Enum_FromProto(mapCtx, in.GetIndustryVertical())
	out.SolutionTypes = direct.EnumSlice_FromProto(mapCtx, in.SolutionTypes)
	out.ContentConfig = direct.Enum_FromProto(mapCtx, in.GetContentConfig())
	out.WorkspaceConfig = WorkspaceConfig_FromProto(mapCtx, in.GetWorkspaceConfig())
	// MISSING: DocumentProcessingConfig
	// MISSING: StartingSchema
	return out
}
func DiscoveryEngineDataStoreSpec_ToProto(mapCtx *direct.MapContext, in *krm.DiscoveryEngineDataStoreSpec) *pb_v1.DataStore {
	if in == nil {
		return nil
	}
	out := &pb_v1.DataStore{}
	// MISSING: Name
	out.DisplayName = direct.ValueOf(in.DisplayName)
	out.IndustryVertical = direct.Enum_ToProto[pb_v1.IndustryVertical](mapCtx, in.IndustryVertical)
	out.SolutionTypes = direct.EnumSlice_ToProto[pb_v1.SolutionType](mapCtx, in.SolutionTypes)
	out.ContentConfig = direct.Enum_ToProto[pb_v1.DataStore_ContentConfig](mapCtx, in.ContentConfig)
	out.WorkspaceConfig = WorkspaceConfig_ToProto(mapCtx, in.WorkspaceConfig)
	// MISSING: DocumentProcessingConfig
	// MISSING: StartingSchema
	return out
}
func DiscoveryEngineDataStoreTargetSiteObservedState_FromProto(mapCtx *direct.MapContext, in *pb_v1.TargetSite) *krm.DiscoveryEngineDataStoreTargetSiteObservedState {
	if in == nil {
		return nil
	}
	out := &krm.DiscoveryEngineDataStoreTargetSiteObservedState{}
	// MISSING: Name
	out.GeneratedURIPattern = direct.LazyPtr(in.GetGeneratedUriPattern())
	out.RootDomainURI = direct.LazyPtr(in.GetRootDomainUri())
	out.SiteVerificationInfo = SiteVerificationInfo_FromProto(mapCtx, in.GetSiteVerificationInfo())
	out.IndexingStatus = direct.Enum_FromProto(mapCtx, in.GetIndexingStatus())
	out.UpdateTime = direct.StringTimestamp_FromProto(mapCtx, in.GetUpdateTime())
	out.FailureReason = TargetSite_FailureReason_FromProto(mapCtx, in.GetFailureReason())
	return out
}
func DiscoveryEngineDataStoreTargetSiteObservedState_ToProto(mapCtx *direct.MapContext, in *krm.DiscoveryEngineDataStoreTargetSiteObservedState) *pb_v1.TargetSite {
	if in == nil {
		return nil
	}
	out := &pb_v1.TargetSite{}
	// MISSING: Name
	out.GeneratedUriPattern = direct.ValueOf(in.GeneratedURIPattern)
	out.RootDomainUri = direct.ValueOf(in.RootDomainURI)
	out.SiteVerificationInfo = SiteVerificationInfo_ToProto(mapCtx, in.SiteVerificationInfo)
	out.IndexingStatus = direct.Enum_ToProto[pb_v1.TargetSite_IndexingStatus](mapCtx, in.IndexingStatus)
	out.UpdateTime = direct.StringTimestamp_ToProto(mapCtx, in.UpdateTime)
	out.FailureReason = TargetSite_FailureReason_ToProto(mapCtx, in.FailureReason)
	return out
}
func DiscoveryEngineDataStoreTargetSiteSpec_FromProto(mapCtx *direct.MapContext, in *pb_v1.TargetSite) *krm.DiscoveryEngineDataStoreTargetSiteSpec {
	if in == nil {
		return nil
	}
	out := &krm.DiscoveryEngineDataStoreTargetSiteSpec{}
	// MISSING: Name
	out.ProvidedURIPattern = direct.LazyPtr(in.GetProvidedUriPattern())
	out.Type = direct.Enum_FromProto(mapCtx, in.GetType())
	out.ExactMatch = direct.LazyPtr(in.GetExactMatch())
	return out
}
func DiscoveryEngineDataStoreTargetSiteSpec_ToProto(mapCtx *direct.MapContext, in *krm.DiscoveryEngineDataStoreTargetSiteSpec) *pb_v1.TargetSite {
	if in == nil {
		return nil
	}
	out := &pb_v1.TargetSite{}
	// MISSING: Name
	out.ProvidedUriPattern = direct.ValueOf(in.ProvidedURIPattern)
	out.Type = direct.Enum_ToProto[pb_v1.TargetSite_Type](mapCtx, in.Type)
	out.ExactMatch = direct.ValueOf(in.ExactMatch)
	return out
}
func DiscoveryEngineEngineObservedState_FromProto(mapCtx *direct.MapContext, in *pb_v1.Engine) *krm.DiscoveryEngineEngineObservedState {
	if in == nil {
		return nil
	}
	out := &krm.DiscoveryEngineEngineObservedState{}
	// MISSING: ChatEngineMetadata
	// MISSING: Name
	// MISSING: CreateTime
	// MISSING: UpdateTime
	// MISSING: DataStoreIds
	return out
}
func DiscoveryEngineEngineObservedState_ToProto(mapCtx *direct.MapContext, in *krm.DiscoveryEngineEngineObservedState) *pb_v1.Engine {
	if in == nil {
		return nil
	}
	out := &pb_v1.Engine{}
	// MISSING: ChatEngineMetadata
	// MISSING: Name
	// MISSING: CreateTime
	// MISSING: UpdateTime
	// MISSING: DataStoreIds
	return out
}
func DiscoveryEngineServingConfigObservedState_FromProto(mapCtx *direct.MapContext, in *pb_v1beta.ServingConfig) *krm.DiscoveryEngineServingConfigObservedState {
	if in == nil {
		return nil
	}
	out := &krm.DiscoveryEngineServingConfigObservedState{}
	// MISSING: CreateTime
	// MISSING: UpdateTime
	return out
}
func DiscoveryEngineServingConfigObservedState_ToProto(mapCtx *direct.MapContext, in *krm.DiscoveryEngineServingConfigObservedState) *pb_v1beta.ServingConfig {
	if in == nil {
		return nil
	}
	out := &pb_v1beta.ServingConfig{}
	// MISSING: CreateTime
	// MISSING: UpdateTime
	return out
}
func DiscoveryEngineServingConfigSpec_FromProto(mapCtx *direct.MapContext, in *pb_v1beta.ServingConfig) *krm.DiscoveryEngineServingConfigSpec {
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
func DiscoveryEngineServingConfigSpec_ToProto(mapCtx *direct.MapContext, in *krm.DiscoveryEngineServingConfigSpec) *pb_v1beta.ServingConfig {
	if in == nil {
		return nil
	}
	out := &pb_v1beta.ServingConfig{}
	if oneof := ServingConfig_MediaConfig_ToProto(mapCtx, in.MediaConfig); oneof != nil {
		out.VerticalConfig = &pb_v1beta.ServingConfig_MediaConfig_{MediaConfig: oneof}
	}
	if oneof := ServingConfig_GenericConfig_ToProto(mapCtx, in.GenericConfig); oneof != nil {
		out.VerticalConfig = &pb_v1beta.ServingConfig_GenericConfig_{GenericConfig: oneof}
	}
	out.Name = direct.ValueOf(in.Name)
	out.DisplayName = direct.ValueOf(in.DisplayName)
	out.SolutionType = direct.Enum_ToProto[pb_v1beta.SolutionType](mapCtx, in.SolutionType)
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
func DocumentProcessingConfig_FromProto(mapCtx *direct.MapContext, in *pb_v1.DocumentProcessingConfig) *krm.DocumentProcessingConfig {
	if in == nil {
		return nil
	}
	out := &krm.DocumentProcessingConfig{}
	out.Name = direct.LazyPtr(in.GetName())
	out.ChunkingConfig = DocumentProcessingConfig_ChunkingConfig_FromProto(mapCtx, in.GetChunkingConfig())
	out.DefaultParsingConfig = DocumentProcessingConfig_ParsingConfig_FromProto(mapCtx, in.GetDefaultParsingConfig())
	// MISSING: ParsingConfigOverrides
	return out
}
func DocumentProcessingConfig_ToProto(mapCtx *direct.MapContext, in *krm.DocumentProcessingConfig) *pb_v1.DocumentProcessingConfig {
	if in == nil {
		return nil
	}
	out := &pb_v1.DocumentProcessingConfig{}
	out.Name = direct.ValueOf(in.Name)
	out.ChunkingConfig = DocumentProcessingConfig_ChunkingConfig_ToProto(mapCtx, in.ChunkingConfig)
	out.DefaultParsingConfig = DocumentProcessingConfig_ParsingConfig_ToProto(mapCtx, in.DefaultParsingConfig)
	// MISSING: ParsingConfigOverrides
	return out
}
func DocumentProcessingConfig_ChunkingConfig_FromProto(mapCtx *direct.MapContext, in *pb_v1.DocumentProcessingConfig_ChunkingConfig) *krm.DocumentProcessingConfig_ChunkingConfig {
	if in == nil {
		return nil
	}
	out := &krm.DocumentProcessingConfig_ChunkingConfig{}
	out.LayoutBasedChunkingConfig = DocumentProcessingConfig_ChunkingConfig_LayoutBasedChunkingConfig_FromProto(mapCtx, in.GetLayoutBasedChunkingConfig())
	return out
}
func DocumentProcessingConfig_ChunkingConfig_ToProto(mapCtx *direct.MapContext, in *krm.DocumentProcessingConfig_ChunkingConfig) *pb_v1.DocumentProcessingConfig_ChunkingConfig {
	if in == nil {
		return nil
	}
	out := &pb_v1.DocumentProcessingConfig_ChunkingConfig{}
	if oneof := DocumentProcessingConfig_ChunkingConfig_LayoutBasedChunkingConfig_ToProto(mapCtx, in.LayoutBasedChunkingConfig); oneof != nil {
		out.ChunkMode = &pb_v1.DocumentProcessingConfig_ChunkingConfig_LayoutBasedChunkingConfig_{LayoutBasedChunkingConfig: oneof}
	}
	return out
}
func DocumentProcessingConfig_ChunkingConfig_LayoutBasedChunkingConfig_FromProto(mapCtx *direct.MapContext, in *pb_v1.DocumentProcessingConfig_ChunkingConfig_LayoutBasedChunkingConfig) *krm.DocumentProcessingConfig_ChunkingConfig_LayoutBasedChunkingConfig {
	if in == nil {
		return nil
	}
	out := &krm.DocumentProcessingConfig_ChunkingConfig_LayoutBasedChunkingConfig{}
	out.ChunkSize = direct.LazyPtr(in.GetChunkSize())
	out.IncludeAncestorHeadings = direct.LazyPtr(in.GetIncludeAncestorHeadings())
	return out
}
func DocumentProcessingConfig_ChunkingConfig_LayoutBasedChunkingConfig_ToProto(mapCtx *direct.MapContext, in *krm.DocumentProcessingConfig_ChunkingConfig_LayoutBasedChunkingConfig) *pb_v1.DocumentProcessingConfig_ChunkingConfig_LayoutBasedChunkingConfig {
	if in == nil {
		return nil
	}
	out := &pb_v1.DocumentProcessingConfig_ChunkingConfig_LayoutBasedChunkingConfig{}
	out.ChunkSize = direct.ValueOf(in.ChunkSize)
	out.IncludeAncestorHeadings = direct.ValueOf(in.IncludeAncestorHeadings)
	return out
}
func DocumentProcessingConfig_ParsingConfig_FromProto(mapCtx *direct.MapContext, in *pb_v1.DocumentProcessingConfig_ParsingConfig) *krm.DocumentProcessingConfig_ParsingConfig {
	if in == nil {
		return nil
	}
	out := &krm.DocumentProcessingConfig_ParsingConfig{}
	out.DigitalParsingConfig = DocumentProcessingConfig_ParsingConfig_DigitalParsingConfig_FromProto(mapCtx, in.GetDigitalParsingConfig())
	out.OcrParsingConfig = DocumentProcessingConfig_ParsingConfig_OcrParsingConfig_FromProto(mapCtx, in.GetOcrParsingConfig())
	out.LayoutParsingConfig = DocumentProcessingConfig_ParsingConfig_LayoutParsingConfig_FromProto(mapCtx, in.GetLayoutParsingConfig())
	return out
}
func DocumentProcessingConfig_ParsingConfig_ToProto(mapCtx *direct.MapContext, in *krm.DocumentProcessingConfig_ParsingConfig) *pb_v1.DocumentProcessingConfig_ParsingConfig {
	if in == nil {
		return nil
	}
	out := &pb_v1.DocumentProcessingConfig_ParsingConfig{}
	if oneof := DocumentProcessingConfig_ParsingConfig_DigitalParsingConfig_ToProto(mapCtx, in.DigitalParsingConfig); oneof != nil {
		out.TypeDedicatedConfig = &pb_v1.DocumentProcessingConfig_ParsingConfig_DigitalParsingConfig_{DigitalParsingConfig: oneof}
	}
	if oneof := DocumentProcessingConfig_ParsingConfig_OcrParsingConfig_ToProto(mapCtx, in.OcrParsingConfig); oneof != nil {
		out.TypeDedicatedConfig = &pb_v1.DocumentProcessingConfig_ParsingConfig_OcrParsingConfig_{OcrParsingConfig: oneof}
	}
	if oneof := DocumentProcessingConfig_ParsingConfig_LayoutParsingConfig_ToProto(mapCtx, in.LayoutParsingConfig); oneof != nil {
		out.TypeDedicatedConfig = &pb_v1.DocumentProcessingConfig_ParsingConfig_LayoutParsingConfig_{LayoutParsingConfig: oneof}
	}
	return out
}
func DocumentProcessingConfig_ParsingConfig_DigitalParsingConfig_FromProto(mapCtx *direct.MapContext, in *pb_v1.DocumentProcessingConfig_ParsingConfig_DigitalParsingConfig) *krm.DocumentProcessingConfig_ParsingConfig_DigitalParsingConfig {
	if in == nil {
		return nil
	}
	out := &krm.DocumentProcessingConfig_ParsingConfig_DigitalParsingConfig{}
	return out
}
func DocumentProcessingConfig_ParsingConfig_DigitalParsingConfig_ToProto(mapCtx *direct.MapContext, in *krm.DocumentProcessingConfig_ParsingConfig_DigitalParsingConfig) *pb_v1.DocumentProcessingConfig_ParsingConfig_DigitalParsingConfig {
	if in == nil {
		return nil
	}
	out := &pb_v1.DocumentProcessingConfig_ParsingConfig_DigitalParsingConfig{}
	return out
}
func DocumentProcessingConfig_ParsingConfig_LayoutParsingConfig_FromProto(mapCtx *direct.MapContext, in *pb_v1.DocumentProcessingConfig_ParsingConfig_LayoutParsingConfig) *krm.DocumentProcessingConfig_ParsingConfig_LayoutParsingConfig {
	if in == nil {
		return nil
	}
	out := &krm.DocumentProcessingConfig_ParsingConfig_LayoutParsingConfig{}
	return out
}
func DocumentProcessingConfig_ParsingConfig_LayoutParsingConfig_ToProto(mapCtx *direct.MapContext, in *krm.DocumentProcessingConfig_ParsingConfig_LayoutParsingConfig) *pb_v1.DocumentProcessingConfig_ParsingConfig_LayoutParsingConfig {
	if in == nil {
		return nil
	}
	out := &pb_v1.DocumentProcessingConfig_ParsingConfig_LayoutParsingConfig{}
	return out
}
func DocumentProcessingConfig_ParsingConfig_OcrParsingConfig_FromProto(mapCtx *direct.MapContext, in *pb_v1.DocumentProcessingConfig_ParsingConfig_OcrParsingConfig) *krm.DocumentProcessingConfig_ParsingConfig_OcrParsingConfig {
	if in == nil {
		return nil
	}
	out := &krm.DocumentProcessingConfig_ParsingConfig_OcrParsingConfig{}
	out.EnhancedDocumentElements = in.EnhancedDocumentElements
	out.UseNativeText = direct.LazyPtr(in.GetUseNativeText())
	return out
}
func DocumentProcessingConfig_ParsingConfig_OcrParsingConfig_ToProto(mapCtx *direct.MapContext, in *krm.DocumentProcessingConfig_ParsingConfig_OcrParsingConfig) *pb_v1.DocumentProcessingConfig_ParsingConfig_OcrParsingConfig {
	if in == nil {
		return nil
	}
	out := &pb_v1.DocumentProcessingConfig_ParsingConfig_OcrParsingConfig{}
	out.EnhancedDocumentElements = in.EnhancedDocumentElements
	out.UseNativeText = direct.ValueOf(in.UseNativeText)
	return out
}
func EmbeddingConfig_FromProto(mapCtx *direct.MapContext, in *pb_v1beta.EmbeddingConfig) *krm.EmbeddingConfig {
	if in == nil {
		return nil
	}
	out := &krm.EmbeddingConfig{}
	out.FieldPath = direct.LazyPtr(in.GetFieldPath())
	return out
}
func EmbeddingConfig_ToProto(mapCtx *direct.MapContext, in *krm.EmbeddingConfig) *pb_v1beta.EmbeddingConfig {
	if in == nil {
		return nil
	}
	out := &pb_v1beta.EmbeddingConfig{}
	out.FieldPath = direct.ValueOf(in.FieldPath)
	return out
}
func Engine_ChatEngineConfig_FromProto(mapCtx *direct.MapContext, in *pb_v1.Engine_ChatEngineConfig) *krm.Engine_ChatEngineConfig {
	if in == nil {
		return nil
	}
	out := &krm.Engine_ChatEngineConfig{}
	out.AgentCreationConfig = Engine_ChatEngineConfig_AgentCreationConfig_FromProto(mapCtx, in.GetAgentCreationConfig())
	out.DialogflowAgentToLink = direct.LazyPtr(in.GetDialogflowAgentToLink())
	return out
}
func Engine_ChatEngineConfig_ToProto(mapCtx *direct.MapContext, in *krm.Engine_ChatEngineConfig) *pb_v1.Engine_ChatEngineConfig {
	if in == nil {
		return nil
	}
	out := &pb_v1.Engine_ChatEngineConfig{}
	out.AgentCreationConfig = Engine_ChatEngineConfig_AgentCreationConfig_ToProto(mapCtx, in.AgentCreationConfig)
	out.DialogflowAgentToLink = direct.ValueOf(in.DialogflowAgentToLink)
	return out
}
func Engine_ChatEngineConfig_AgentCreationConfig_FromProto(mapCtx *direct.MapContext, in *pb_v1.Engine_ChatEngineConfig_AgentCreationConfig) *krm.Engine_ChatEngineConfig_AgentCreationConfig {
	if in == nil {
		return nil
	}
	out := &krm.Engine_ChatEngineConfig_AgentCreationConfig{}
	out.Business = direct.LazyPtr(in.GetBusiness())
	out.DefaultLanguageCode = direct.LazyPtr(in.GetDefaultLanguageCode())
	out.TimeZone = direct.LazyPtr(in.GetTimeZone())
	out.Location = direct.LazyPtr(in.GetLocation())
	return out
}
func Engine_ChatEngineConfig_AgentCreationConfig_ToProto(mapCtx *direct.MapContext, in *krm.Engine_ChatEngineConfig_AgentCreationConfig) *pb_v1.Engine_ChatEngineConfig_AgentCreationConfig {
	if in == nil {
		return nil
	}
	out := &pb_v1.Engine_ChatEngineConfig_AgentCreationConfig{}
	out.Business = direct.ValueOf(in.Business)
	out.DefaultLanguageCode = direct.ValueOf(in.DefaultLanguageCode)
	out.TimeZone = direct.ValueOf(in.TimeZone)
	out.Location = direct.ValueOf(in.Location)
	return out
}
func Engine_CommonConfig_FromProto(mapCtx *direct.MapContext, in *pb_v1.Engine_CommonConfig) *krm.Engine_CommonConfig {
	if in == nil {
		return nil
	}
	out := &krm.Engine_CommonConfig{}
	out.CompanyName = direct.LazyPtr(in.GetCompanyName())
	return out
}
func Engine_CommonConfig_ToProto(mapCtx *direct.MapContext, in *krm.Engine_CommonConfig) *pb_v1.Engine_CommonConfig {
	if in == nil {
		return nil
	}
	out := &pb_v1.Engine_CommonConfig{}
	out.CompanyName = direct.ValueOf(in.CompanyName)
	return out
}
func Engine_SearchEngineConfig_FromProto(mapCtx *direct.MapContext, in *pb_v1.Engine_SearchEngineConfig) *krm.Engine_SearchEngineConfig {
	if in == nil {
		return nil
	}
	out := &krm.Engine_SearchEngineConfig{}
	out.SearchTier = direct.Enum_FromProto(mapCtx, in.GetSearchTier())
	out.SearchAddOns = direct.EnumSlice_FromProto(mapCtx, in.SearchAddOns)
	return out
}
func Engine_SearchEngineConfig_ToProto(mapCtx *direct.MapContext, in *krm.Engine_SearchEngineConfig) *pb_v1.Engine_SearchEngineConfig {
	if in == nil {
		return nil
	}
	out := &pb_v1.Engine_SearchEngineConfig{}
	out.SearchTier = direct.Enum_ToProto[pb_v1.SearchTier](mapCtx, in.SearchTier)
	out.SearchAddOns = direct.EnumSlice_ToProto[pb_v1.SearchAddOn](mapCtx, in.SearchAddOns)
	return out
}
func Schema_FromProto(mapCtx *direct.MapContext, in *pb_v1.Schema) *krm.Schema {
	if in == nil {
		return nil
	}
	out := &krm.Schema{}
	out.StructSchema = StructSchema_FromProto(mapCtx, in.GetStructSchema())
	out.JsonSchema = direct.LazyPtr(in.GetJsonSchema())
	out.Name = direct.LazyPtr(in.GetName())
	return out
}
func Schema_ToProto(mapCtx *direct.MapContext, in *krm.Schema) *pb_v1.Schema {
	if in == nil {
		return nil
	}
	out := &pb_v1.Schema{}
	if oneof := StructSchema_ToProto(mapCtx, in.StructSchema); oneof != nil {
		out.Schema = &pb_v1.Schema_StructSchema{StructSchema: oneof}
	}
	if oneof := Schema_JsonSchema_ToProto(mapCtx, in.JsonSchema); oneof != nil {
		out.Schema = oneof
	}
	out.Name = direct.ValueOf(in.Name)
	return out
}
func SearchRequest_ContentSearchSpec_FromProto(mapCtx *direct.MapContext, in *pb_v1beta.SearchRequest_ContentSearchSpec) *krm.SearchRequest_ContentSearchSpec {
	if in == nil {
		return nil
	}
	out := &krm.SearchRequest_ContentSearchSpec{}
	out.SnippetSpec = SearchRequest_ContentSearchSpec_SnippetSpec_FromProto(mapCtx, in.GetSnippetSpec())
	out.SummarySpec = SearchRequest_ContentSearchSpec_SummarySpec_FromProto(mapCtx, in.GetSummarySpec())
	out.ExtractiveContentSpec = SearchRequest_ContentSearchSpec_ExtractiveContentSpec_FromProto(mapCtx, in.GetExtractiveContentSpec())
	out.SearchResultMode = direct.Enum_FromProto(mapCtx, in.GetSearchResultMode())
	out.ChunkSpec = SearchRequest_ContentSearchSpec_ChunkSpec_FromProto(mapCtx, in.GetChunkSpec())
	return out
}
func SearchRequest_ContentSearchSpec_ToProto(mapCtx *direct.MapContext, in *krm.SearchRequest_ContentSearchSpec) *pb_v1beta.SearchRequest_ContentSearchSpec {
	if in == nil {
		return nil
	}
	out := &pb_v1beta.SearchRequest_ContentSearchSpec{}
	out.SnippetSpec = SearchRequest_ContentSearchSpec_SnippetSpec_ToProto(mapCtx, in.SnippetSpec)
	out.SummarySpec = SearchRequest_ContentSearchSpec_SummarySpec_ToProto(mapCtx, in.SummarySpec)
	out.ExtractiveContentSpec = SearchRequest_ContentSearchSpec_ExtractiveContentSpec_ToProto(mapCtx, in.ExtractiveContentSpec)
	out.SearchResultMode = direct.Enum_ToProto[pb_v1beta.SearchRequest_ContentSearchSpec_SearchResultMode](mapCtx, in.SearchResultMode)
	out.ChunkSpec = SearchRequest_ContentSearchSpec_ChunkSpec_ToProto(mapCtx, in.ChunkSpec)
	return out
}
func SearchRequest_ContentSearchSpec_ChunkSpec_FromProto(mapCtx *direct.MapContext, in *pb_v1beta.SearchRequest_ContentSearchSpec_ChunkSpec) *krm.SearchRequest_ContentSearchSpec_ChunkSpec {
	if in == nil {
		return nil
	}
	out := &krm.SearchRequest_ContentSearchSpec_ChunkSpec{}
	out.NumPreviousChunks = direct.LazyPtr(in.GetNumPreviousChunks())
	out.NumNextChunks = direct.LazyPtr(in.GetNumNextChunks())
	return out
}
func SearchRequest_ContentSearchSpec_ChunkSpec_ToProto(mapCtx *direct.MapContext, in *krm.SearchRequest_ContentSearchSpec_ChunkSpec) *pb_v1beta.SearchRequest_ContentSearchSpec_ChunkSpec {
	if in == nil {
		return nil
	}
	out := &pb_v1beta.SearchRequest_ContentSearchSpec_ChunkSpec{}
	out.NumPreviousChunks = direct.ValueOf(in.NumPreviousChunks)
	out.NumNextChunks = direct.ValueOf(in.NumNextChunks)
	return out
}
func SearchRequest_ContentSearchSpec_ExtractiveContentSpec_FromProto(mapCtx *direct.MapContext, in *pb_v1beta.SearchRequest_ContentSearchSpec_ExtractiveContentSpec) *krm.SearchRequest_ContentSearchSpec_ExtractiveContentSpec {
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
func SearchRequest_ContentSearchSpec_ExtractiveContentSpec_ToProto(mapCtx *direct.MapContext, in *krm.SearchRequest_ContentSearchSpec_ExtractiveContentSpec) *pb_v1beta.SearchRequest_ContentSearchSpec_ExtractiveContentSpec {
	if in == nil {
		return nil
	}
	out := &pb_v1beta.SearchRequest_ContentSearchSpec_ExtractiveContentSpec{}
	out.MaxExtractiveAnswerCount = direct.ValueOf(in.MaxExtractiveAnswerCount)
	out.MaxExtractiveSegmentCount = direct.ValueOf(in.MaxExtractiveSegmentCount)
	out.ReturnExtractiveSegmentScore = direct.ValueOf(in.ReturnExtractiveSegmentScore)
	out.NumPreviousSegments = direct.ValueOf(in.NumPreviousSegments)
	out.NumNextSegments = direct.ValueOf(in.NumNextSegments)
	return out
}
func SearchRequest_ContentSearchSpec_SnippetSpec_FromProto(mapCtx *direct.MapContext, in *pb_v1beta.SearchRequest_ContentSearchSpec_SnippetSpec) *krm.SearchRequest_ContentSearchSpec_SnippetSpec {
	if in == nil {
		return nil
	}
	out := &krm.SearchRequest_ContentSearchSpec_SnippetSpec{}
	out.MaxSnippetCount = direct.LazyPtr(in.GetMaxSnippetCount())
	out.ReferenceOnly = direct.LazyPtr(in.GetReferenceOnly())
	out.ReturnSnippet = direct.LazyPtr(in.GetReturnSnippet())
	return out
}
func SearchRequest_ContentSearchSpec_SnippetSpec_ToProto(mapCtx *direct.MapContext, in *krm.SearchRequest_ContentSearchSpec_SnippetSpec) *pb_v1beta.SearchRequest_ContentSearchSpec_SnippetSpec {
	if in == nil {
		return nil
	}
	out := &pb_v1beta.SearchRequest_ContentSearchSpec_SnippetSpec{}
	out.MaxSnippetCount = direct.ValueOf(in.MaxSnippetCount)
	out.ReferenceOnly = direct.ValueOf(in.ReferenceOnly)
	out.ReturnSnippet = direct.ValueOf(in.ReturnSnippet)
	return out
}
func SearchRequest_ContentSearchSpec_SummarySpec_FromProto(mapCtx *direct.MapContext, in *pb_v1beta.SearchRequest_ContentSearchSpec_SummarySpec) *krm.SearchRequest_ContentSearchSpec_SummarySpec {
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
	out.ModelPromptSpec = SearchRequest_ContentSearchSpec_SummarySpec_ModelPromptSpec_FromProto(mapCtx, in.GetModelPromptSpec())
	out.LanguageCode = direct.LazyPtr(in.GetLanguageCode())
	out.ModelSpec = SearchRequest_ContentSearchSpec_SummarySpec_ModelSpec_FromProto(mapCtx, in.GetModelSpec())
	out.UseSemanticChunks = direct.LazyPtr(in.GetUseSemanticChunks())
	return out
}
func SearchRequest_ContentSearchSpec_SummarySpec_ToProto(mapCtx *direct.MapContext, in *krm.SearchRequest_ContentSearchSpec_SummarySpec) *pb_v1beta.SearchRequest_ContentSearchSpec_SummarySpec {
	if in == nil {
		return nil
	}
	out := &pb_v1beta.SearchRequest_ContentSearchSpec_SummarySpec{}
	out.SummaryResultCount = direct.ValueOf(in.SummaryResultCount)
	out.IncludeCitations = direct.ValueOf(in.IncludeCitations)
	out.IgnoreAdversarialQuery = direct.ValueOf(in.IgnoreAdversarialQuery)
	out.IgnoreNonSummarySeekingQuery = direct.ValueOf(in.IgnoreNonSummarySeekingQuery)
	out.IgnoreLowRelevantContent = direct.ValueOf(in.IgnoreLowRelevantContent)
	out.IgnoreJailBreakingQuery = direct.ValueOf(in.IgnoreJailBreakingQuery)
	out.ModelPromptSpec = SearchRequest_ContentSearchSpec_SummarySpec_ModelPromptSpec_ToProto(mapCtx, in.ModelPromptSpec)
	out.LanguageCode = direct.ValueOf(in.LanguageCode)
	out.ModelSpec = SearchRequest_ContentSearchSpec_SummarySpec_ModelSpec_ToProto(mapCtx, in.ModelSpec)
	out.UseSemanticChunks = direct.ValueOf(in.UseSemanticChunks)
	return out
}
func SearchRequest_ContentSearchSpec_SummarySpec_ModelPromptSpec_FromProto(mapCtx *direct.MapContext, in *pb_v1beta.SearchRequest_ContentSearchSpec_SummarySpec_ModelPromptSpec) *krm.SearchRequest_ContentSearchSpec_SummarySpec_ModelPromptSpec {
	if in == nil {
		return nil
	}
	out := &krm.SearchRequest_ContentSearchSpec_SummarySpec_ModelPromptSpec{}
	out.Preamble = direct.LazyPtr(in.GetPreamble())
	return out
}
func SearchRequest_ContentSearchSpec_SummarySpec_ModelPromptSpec_ToProto(mapCtx *direct.MapContext, in *krm.SearchRequest_ContentSearchSpec_SummarySpec_ModelPromptSpec) *pb_v1beta.SearchRequest_ContentSearchSpec_SummarySpec_ModelPromptSpec {
	if in == nil {
		return nil
	}
	out := &pb_v1beta.SearchRequest_ContentSearchSpec_SummarySpec_ModelPromptSpec{}
	out.Preamble = direct.ValueOf(in.Preamble)
	return out
}
func SearchRequest_ContentSearchSpec_SummarySpec_ModelSpec_FromProto(mapCtx *direct.MapContext, in *pb_v1beta.SearchRequest_ContentSearchSpec_SummarySpec_ModelSpec) *krm.SearchRequest_ContentSearchSpec_SummarySpec_ModelSpec {
	if in == nil {
		return nil
	}
	out := &krm.SearchRequest_ContentSearchSpec_SummarySpec_ModelSpec{}
	out.Version = direct.LazyPtr(in.GetVersion())
	return out
}
func SearchRequest_ContentSearchSpec_SummarySpec_ModelSpec_ToProto(mapCtx *direct.MapContext, in *krm.SearchRequest_ContentSearchSpec_SummarySpec_ModelSpec) *pb_v1beta.SearchRequest_ContentSearchSpec_SummarySpec_ModelSpec {
	if in == nil {
		return nil
	}
	out := &pb_v1beta.SearchRequest_ContentSearchSpec_SummarySpec_ModelSpec{}
	out.Version = direct.ValueOf(in.Version)
	return out
}
func SearchRequest_PersonalizationSpec_FromProto(mapCtx *direct.MapContext, in *pb_v1beta.SearchRequest_PersonalizationSpec) *krm.SearchRequest_PersonalizationSpec {
	if in == nil {
		return nil
	}
	out := &krm.SearchRequest_PersonalizationSpec{}
	out.Mode = direct.Enum_FromProto(mapCtx, in.GetMode())
	return out
}
func SearchRequest_PersonalizationSpec_ToProto(mapCtx *direct.MapContext, in *krm.SearchRequest_PersonalizationSpec) *pb_v1beta.SearchRequest_PersonalizationSpec {
	if in == nil {
		return nil
	}
	out := &pb_v1beta.SearchRequest_PersonalizationSpec{}
	out.Mode = direct.Enum_ToProto[pb_v1beta.SearchRequest_PersonalizationSpec_Mode](mapCtx, in.Mode)
	return out
}
func ServingConfig_GenericConfig_FromProto(mapCtx *direct.MapContext, in *pb_v1beta.ServingConfig_GenericConfig) *krm.ServingConfig_GenericConfig {
	if in == nil {
		return nil
	}
	out := &krm.ServingConfig_GenericConfig{}
	out.ContentSearchSpec = SearchRequest_ContentSearchSpec_FromProto(mapCtx, in.GetContentSearchSpec())
	return out
}
func ServingConfig_GenericConfig_ToProto(mapCtx *direct.MapContext, in *krm.ServingConfig_GenericConfig) *pb_v1beta.ServingConfig_GenericConfig {
	if in == nil {
		return nil
	}
	out := &pb_v1beta.ServingConfig_GenericConfig{}
	out.ContentSearchSpec = SearchRequest_ContentSearchSpec_ToProto(mapCtx, in.ContentSearchSpec)
	return out
}
func ServingConfig_MediaConfig_FromProto(mapCtx *direct.MapContext, in *pb_v1beta.ServingConfig_MediaConfig) *krm.ServingConfig_MediaConfig {
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
func ServingConfig_MediaConfig_ToProto(mapCtx *direct.MapContext, in *krm.ServingConfig_MediaConfig) *pb_v1beta.ServingConfig_MediaConfig {
	if in == nil {
		return nil
	}
	out := &pb_v1beta.ServingConfig_MediaConfig{}
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
func SiteVerificationInfo_FromProto(mapCtx *direct.MapContext, in *pb_v1.SiteVerificationInfo) *krm.SiteVerificationInfo {
	if in == nil {
		return nil
	}
	out := &krm.SiteVerificationInfo{}
	out.SiteVerificationState = direct.Enum_FromProto(mapCtx, in.GetSiteVerificationState())
	out.VerifyTime = direct.StringTimestamp_FromProto(mapCtx, in.GetVerifyTime())
	return out
}
func SiteVerificationInfo_ToProto(mapCtx *direct.MapContext, in *krm.SiteVerificationInfo) *pb_v1.SiteVerificationInfo {
	if in == nil {
		return nil
	}
	out := &pb_v1.SiteVerificationInfo{}
	out.SiteVerificationState = direct.Enum_ToProto[pb_v1.SiteVerificationInfo_SiteVerificationState](mapCtx, in.SiteVerificationState)
	out.VerifyTime = direct.StringTimestamp_ToProto(mapCtx, in.VerifyTime)
	return out
}
func TargetSite_FailureReason_FromProto(mapCtx *direct.MapContext, in *pb_v1.TargetSite_FailureReason) *krm.TargetSite_FailureReason {
	if in == nil {
		return nil
	}
	out := &krm.TargetSite_FailureReason{}
	out.QuotaFailure = TargetSite_FailureReason_QuotaFailure_FromProto(mapCtx, in.GetQuotaFailure())
	return out
}
func TargetSite_FailureReason_ToProto(mapCtx *direct.MapContext, in *krm.TargetSite_FailureReason) *pb_v1.TargetSite_FailureReason {
	if in == nil {
		return nil
	}
	out := &pb_v1.TargetSite_FailureReason{}
	if oneof := TargetSite_FailureReason_QuotaFailure_ToProto(mapCtx, in.QuotaFailure); oneof != nil {
		out.Failure = &pb_v1.TargetSite_FailureReason_QuotaFailure_{QuotaFailure: oneof}
	}
	return out
}
func TargetSite_FailureReason_QuotaFailure_FromProto(mapCtx *direct.MapContext, in *pb_v1.TargetSite_FailureReason_QuotaFailure) *krm.TargetSite_FailureReason_QuotaFailure {
	if in == nil {
		return nil
	}
	out := &krm.TargetSite_FailureReason_QuotaFailure{}
	out.TotalRequiredQuota = direct.LazyPtr(in.GetTotalRequiredQuota())
	return out
}
func TargetSite_FailureReason_QuotaFailure_ToProto(mapCtx *direct.MapContext, in *krm.TargetSite_FailureReason_QuotaFailure) *pb_v1.TargetSite_FailureReason_QuotaFailure {
	if in == nil {
		return nil
	}
	out := &pb_v1.TargetSite_FailureReason_QuotaFailure{}
	out.TotalRequiredQuota = direct.ValueOf(in.TotalRequiredQuota)
	return out
}
func WorkspaceConfig_FromProto(mapCtx *direct.MapContext, in *pb_v1.WorkspaceConfig) *krm.WorkspaceConfig {
	if in == nil {
		return nil
	}
	out := &krm.WorkspaceConfig{}
	out.Type = direct.Enum_FromProto(mapCtx, in.GetType())
	out.DasherCustomerID = direct.LazyPtr(in.GetDasherCustomerId())
	out.SuperAdminServiceAccount = direct.LazyPtr(in.GetSuperAdminServiceAccount())
	out.SuperAdminEmailAddress = direct.LazyPtr(in.GetSuperAdminEmailAddress())
	return out
}
func WorkspaceConfig_ToProto(mapCtx *direct.MapContext, in *krm.WorkspaceConfig) *pb_v1.WorkspaceConfig {
	if in == nil {
		return nil
	}
	out := &pb_v1.WorkspaceConfig{}
	out.Type = direct.Enum_ToProto[pb_v1.WorkspaceConfig_Type](mapCtx, in.Type)
	out.DasherCustomerId = direct.ValueOf(in.DasherCustomerID)
	out.SuperAdminServiceAccount = direct.ValueOf(in.SuperAdminServiceAccount)
	out.SuperAdminEmailAddress = direct.ValueOf(in.SuperAdminEmailAddress)
	return out
}
