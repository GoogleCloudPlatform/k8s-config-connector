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

// +generated:mapper
// krm.group: discoveryengine.cnrm.cloud.google.com
// krm.version: v1alpha1
// proto.service: google.cloud.discoveryengine.v1alpha

package discoveryengine

import (
	pb "cloud.google.com/go/discoveryengine/apiv1alpha/discoveryenginepb"
	krmv1alpha1 "github.com/GoogleCloudPlatform/k8s-config-connector/apis/discoveryengine/v1alpha1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct"
)

func DataStoreObservedState_FromProto(mapCtx *direct.MapContext, in *pb.DataStore) *krmv1alpha1.DataStoreObservedState {
	if in == nil {
		return nil
	}
	out := &krmv1alpha1.DataStoreObservedState{}
	// MISSING: Name
	// MISSING: DisplayName
	// MISSING: IndustryVertical
	// MISSING: SolutionTypes
	out.DefaultSchemaID = direct.LazyPtr(in.GetDefaultSchemaId())
	// MISSING: ContentConfig
	out.CreateTime = direct.StringTimestamp_FromProto(mapCtx, in.GetCreateTime())
	out.LanguageInfo = LanguageInfoObservedState_FromProto(mapCtx, in.GetLanguageInfo())
	// MISSING: IdpConfig
	// MISSING: AclEnabled
	// MISSING: WorkspaceConfig
	// MISSING: DocumentProcessingConfig
	out.StartingSchema = SchemaObservedState_FromProto(mapCtx, in.GetStartingSchema())
	return out
}
func DataStoreObservedState_ToProto(mapCtx *direct.MapContext, in *krmv1alpha1.DataStoreObservedState) *pb.DataStore {
	if in == nil {
		return nil
	}
	out := &pb.DataStore{}
	// MISSING: Name
	// MISSING: DisplayName
	// MISSING: IndustryVertical
	// MISSING: SolutionTypes
	out.DefaultSchemaId = direct.ValueOf(in.DefaultSchemaID)
	// MISSING: ContentConfig
	out.CreateTime = direct.StringTimestamp_ToProto(mapCtx, in.CreateTime)
	out.LanguageInfo = LanguageInfoObservedState_ToProto(mapCtx, in.LanguageInfo)
	// MISSING: IdpConfig
	// MISSING: AclEnabled
	// MISSING: WorkspaceConfig
	// MISSING: DocumentProcessingConfig
	out.StartingSchema = SchemaObservedState_ToProto(mapCtx, in.StartingSchema)
	return out
}
func DocumentProcessingConfig_FromProto(mapCtx *direct.MapContext, in *pb.DocumentProcessingConfig) *krmv1alpha1.DocumentProcessingConfig {
	if in == nil {
		return nil
	}
	out := &krmv1alpha1.DocumentProcessingConfig{}
	out.Name = direct.LazyPtr(in.GetName())
	out.ChunkingConfig = DocumentProcessingConfig_ChunkingConfig_FromProto(mapCtx, in.GetChunkingConfig())
	out.DefaultParsingConfig = DocumentProcessingConfig_ParsingConfig_FromProto(mapCtx, in.GetDefaultParsingConfig())
	// MISSING: ParsingConfigOverrides
	return out
}
func DocumentProcessingConfig_ToProto(mapCtx *direct.MapContext, in *krmv1alpha1.DocumentProcessingConfig) *pb.DocumentProcessingConfig {
	if in == nil {
		return nil
	}
	out := &pb.DocumentProcessingConfig{}
	out.Name = direct.ValueOf(in.Name)
	out.ChunkingConfig = DocumentProcessingConfig_ChunkingConfig_ToProto(mapCtx, in.ChunkingConfig)
	out.DefaultParsingConfig = DocumentProcessingConfig_ParsingConfig_ToProto(mapCtx, in.DefaultParsingConfig)
	// MISSING: ParsingConfigOverrides
	return out
}
func DocumentProcessingConfig_ChunkingConfig_FromProto(mapCtx *direct.MapContext, in *pb.DocumentProcessingConfig_ChunkingConfig) *krmv1alpha1.DocumentProcessingConfig_ChunkingConfig {
	if in == nil {
		return nil
	}
	out := &krmv1alpha1.DocumentProcessingConfig_ChunkingConfig{}
	out.LayoutBasedChunkingConfig = DocumentProcessingConfig_ChunkingConfig_LayoutBasedChunkingConfig_FromProto(mapCtx, in.GetLayoutBasedChunkingConfig())
	return out
}
func DocumentProcessingConfig_ChunkingConfig_ToProto(mapCtx *direct.MapContext, in *krmv1alpha1.DocumentProcessingConfig_ChunkingConfig) *pb.DocumentProcessingConfig_ChunkingConfig {
	if in == nil {
		return nil
	}
	out := &pb.DocumentProcessingConfig_ChunkingConfig{}
	if oneof := DocumentProcessingConfig_ChunkingConfig_LayoutBasedChunkingConfig_ToProto(mapCtx, in.LayoutBasedChunkingConfig); oneof != nil {
		out.ChunkMode = &pb.DocumentProcessingConfig_ChunkingConfig_LayoutBasedChunkingConfig_{LayoutBasedChunkingConfig: oneof}
	}
	return out
}
func DocumentProcessingConfig_ChunkingConfig_LayoutBasedChunkingConfig_FromProto(mapCtx *direct.MapContext, in *pb.DocumentProcessingConfig_ChunkingConfig_LayoutBasedChunkingConfig) *krmv1alpha1.DocumentProcessingConfig_ChunkingConfig_LayoutBasedChunkingConfig {
	if in == nil {
		return nil
	}
	out := &krmv1alpha1.DocumentProcessingConfig_ChunkingConfig_LayoutBasedChunkingConfig{}
	out.ChunkSize = direct.LazyPtr(in.GetChunkSize())
	out.IncludeAncestorHeadings = direct.LazyPtr(in.GetIncludeAncestorHeadings())
	return out
}
func DocumentProcessingConfig_ChunkingConfig_LayoutBasedChunkingConfig_ToProto(mapCtx *direct.MapContext, in *krmv1alpha1.DocumentProcessingConfig_ChunkingConfig_LayoutBasedChunkingConfig) *pb.DocumentProcessingConfig_ChunkingConfig_LayoutBasedChunkingConfig {
	if in == nil {
		return nil
	}
	out := &pb.DocumentProcessingConfig_ChunkingConfig_LayoutBasedChunkingConfig{}
	out.ChunkSize = direct.ValueOf(in.ChunkSize)
	out.IncludeAncestorHeadings = direct.ValueOf(in.IncludeAncestorHeadings)
	return out
}
func DocumentProcessingConfig_ParsingConfig_FromProto(mapCtx *direct.MapContext, in *pb.DocumentProcessingConfig_ParsingConfig) *krmv1alpha1.DocumentProcessingConfig_ParsingConfig {
	if in == nil {
		return nil
	}
	out := &krmv1alpha1.DocumentProcessingConfig_ParsingConfig{}
	out.DigitalParsingConfig = DocumentProcessingConfig_ParsingConfig_DigitalParsingConfig_FromProto(mapCtx, in.GetDigitalParsingConfig())
	out.OcrParsingConfig = DocumentProcessingConfig_ParsingConfig_OcrParsingConfig_FromProto(mapCtx, in.GetOcrParsingConfig())
	out.LayoutParsingConfig = DocumentProcessingConfig_ParsingConfig_LayoutParsingConfig_FromProto(mapCtx, in.GetLayoutParsingConfig())
	return out
}
func DocumentProcessingConfig_ParsingConfig_ToProto(mapCtx *direct.MapContext, in *krmv1alpha1.DocumentProcessingConfig_ParsingConfig) *pb.DocumentProcessingConfig_ParsingConfig {
	if in == nil {
		return nil
	}
	out := &pb.DocumentProcessingConfig_ParsingConfig{}
	if oneof := DocumentProcessingConfig_ParsingConfig_DigitalParsingConfig_ToProto(mapCtx, in.DigitalParsingConfig); oneof != nil {
		out.TypeDedicatedConfig = &pb.DocumentProcessingConfig_ParsingConfig_DigitalParsingConfig_{DigitalParsingConfig: oneof}
	}
	if oneof := DocumentProcessingConfig_ParsingConfig_OcrParsingConfig_ToProto(mapCtx, in.OcrParsingConfig); oneof != nil {
		out.TypeDedicatedConfig = &pb.DocumentProcessingConfig_ParsingConfig_OcrParsingConfig_{OcrParsingConfig: oneof}
	}
	if oneof := DocumentProcessingConfig_ParsingConfig_LayoutParsingConfig_ToProto(mapCtx, in.LayoutParsingConfig); oneof != nil {
		out.TypeDedicatedConfig = &pb.DocumentProcessingConfig_ParsingConfig_LayoutParsingConfig_{LayoutParsingConfig: oneof}
	}
	return out
}
func DocumentProcessingConfig_ParsingConfig_DigitalParsingConfig_FromProto(mapCtx *direct.MapContext, in *pb.DocumentProcessingConfig_ParsingConfig_DigitalParsingConfig) *krmv1alpha1.DocumentProcessingConfig_ParsingConfig_DigitalParsingConfig {
	if in == nil {
		return nil
	}
	out := &krmv1alpha1.DocumentProcessingConfig_ParsingConfig_DigitalParsingConfig{}
	return out
}
func DocumentProcessingConfig_ParsingConfig_DigitalParsingConfig_ToProto(mapCtx *direct.MapContext, in *krmv1alpha1.DocumentProcessingConfig_ParsingConfig_DigitalParsingConfig) *pb.DocumentProcessingConfig_ParsingConfig_DigitalParsingConfig {
	if in == nil {
		return nil
	}
	out := &pb.DocumentProcessingConfig_ParsingConfig_DigitalParsingConfig{}
	return out
}
func DocumentProcessingConfig_ParsingConfig_LayoutParsingConfig_FromProto(mapCtx *direct.MapContext, in *pb.DocumentProcessingConfig_ParsingConfig_LayoutParsingConfig) *krmv1alpha1.DocumentProcessingConfig_ParsingConfig_LayoutParsingConfig {
	if in == nil {
		return nil
	}
	out := &krmv1alpha1.DocumentProcessingConfig_ParsingConfig_LayoutParsingConfig{}
	return out
}
func DocumentProcessingConfig_ParsingConfig_LayoutParsingConfig_ToProto(mapCtx *direct.MapContext, in *krmv1alpha1.DocumentProcessingConfig_ParsingConfig_LayoutParsingConfig) *pb.DocumentProcessingConfig_ParsingConfig_LayoutParsingConfig {
	if in == nil {
		return nil
	}
	out := &pb.DocumentProcessingConfig_ParsingConfig_LayoutParsingConfig{}
	return out
}
func DocumentProcessingConfig_ParsingConfig_OcrParsingConfig_FromProto(mapCtx *direct.MapContext, in *pb.DocumentProcessingConfig_ParsingConfig_OcrParsingConfig) *krmv1alpha1.DocumentProcessingConfig_ParsingConfig_OcrParsingConfig {
	if in == nil {
		return nil
	}
	out := &krmv1alpha1.DocumentProcessingConfig_ParsingConfig_OcrParsingConfig{}
	out.EnhancedDocumentElements = in.EnhancedDocumentElements
	out.UseNativeText = direct.LazyPtr(in.GetUseNativeText())
	return out
}
func DocumentProcessingConfig_ParsingConfig_OcrParsingConfig_ToProto(mapCtx *direct.MapContext, in *krmv1alpha1.DocumentProcessingConfig_ParsingConfig_OcrParsingConfig) *pb.DocumentProcessingConfig_ParsingConfig_OcrParsingConfig {
	if in == nil {
		return nil
	}
	out := &pb.DocumentProcessingConfig_ParsingConfig_OcrParsingConfig{}
	out.EnhancedDocumentElements = in.EnhancedDocumentElements
	out.UseNativeText = direct.ValueOf(in.UseNativeText)
	return out
}
func Engine_FromProto(mapCtx *direct.MapContext, in *pb.Engine) *krmv1alpha1.Engine {
	if in == nil {
		return nil
	}
	out := &krmv1alpha1.Engine{}
	out.SimilarDocumentsConfig = Engine_SimilarDocumentsEngineConfig_FromProto(mapCtx, in.GetSimilarDocumentsConfig())
	out.ChatEngineConfig = Engine_ChatEngineConfig_FromProto(mapCtx, in.GetChatEngineConfig())
	out.SearchEngineConfig = Engine_SearchEngineConfig_FromProto(mapCtx, in.GetSearchEngineConfig())
	out.MediaRecommendationEngineConfig = Engine_MediaRecommendationEngineConfig_FromProto(mapCtx, in.GetMediaRecommendationEngineConfig())
	// MISSING: RecommendationMetadata
	// MISSING: ChatEngineMetadata
	out.Name = direct.LazyPtr(in.GetName())
	out.DisplayName = direct.LazyPtr(in.GetDisplayName())
	// MISSING: CreateTime
	// MISSING: UpdateTime
	out.DataStoreIds = in.DataStoreIds
	out.SolutionType = direct.Enum_FromProto(mapCtx, in.GetSolutionType())
	out.IndustryVertical = direct.Enum_FromProto(mapCtx, in.GetIndustryVertical())
	out.CommonConfig = Engine_CommonConfig_FromProto(mapCtx, in.GetCommonConfig())
	return out
}
func Engine_ToProto(mapCtx *direct.MapContext, in *krmv1alpha1.Engine) *pb.Engine {
	if in == nil {
		return nil
	}
	out := &pb.Engine{}
	if oneof := Engine_SimilarDocumentsEngineConfig_ToProto(mapCtx, in.SimilarDocumentsConfig); oneof != nil {
		out.EngineConfig = &pb.Engine_SimilarDocumentsConfig{SimilarDocumentsConfig: oneof}
	}
	if oneof := Engine_ChatEngineConfig_ToProto(mapCtx, in.ChatEngineConfig); oneof != nil {
		out.EngineConfig = &pb.Engine_ChatEngineConfig_{ChatEngineConfig: oneof}
	}
	if oneof := Engine_SearchEngineConfig_ToProto(mapCtx, in.SearchEngineConfig); oneof != nil {
		out.EngineConfig = &pb.Engine_SearchEngineConfig_{SearchEngineConfig: oneof}
	}
	if oneof := Engine_MediaRecommendationEngineConfig_ToProto(mapCtx, in.MediaRecommendationEngineConfig); oneof != nil {
		out.EngineConfig = &pb.Engine_MediaRecommendationEngineConfig_{MediaRecommendationEngineConfig: oneof}
	}
	// MISSING: RecommendationMetadata
	// MISSING: ChatEngineMetadata
	out.Name = direct.ValueOf(in.Name)
	out.DisplayName = direct.ValueOf(in.DisplayName)
	// MISSING: CreateTime
	// MISSING: UpdateTime
	out.DataStoreIds = in.DataStoreIds
	out.SolutionType = direct.Enum_ToProto[pb.SolutionType](mapCtx, in.SolutionType)
	out.IndustryVertical = direct.Enum_ToProto[pb.IndustryVertical](mapCtx, in.IndustryVertical)
	out.CommonConfig = Engine_CommonConfig_ToProto(mapCtx, in.CommonConfig)
	return out
}
func Engine_ChatEngineConfig_FromProto(mapCtx *direct.MapContext, in *pb.Engine_ChatEngineConfig) *krmv1alpha1.Engine_ChatEngineConfig {
	if in == nil {
		return nil
	}
	out := &krmv1alpha1.Engine_ChatEngineConfig{}
	out.AgentCreationConfig = Engine_ChatEngineConfig_AgentCreationConfig_FromProto(mapCtx, in.GetAgentCreationConfig())
	out.DialogflowAgentToLink = direct.LazyPtr(in.GetDialogflowAgentToLink())
	return out
}
func Engine_ChatEngineConfig_ToProto(mapCtx *direct.MapContext, in *krmv1alpha1.Engine_ChatEngineConfig) *pb.Engine_ChatEngineConfig {
	if in == nil {
		return nil
	}
	out := &pb.Engine_ChatEngineConfig{}
	out.AgentCreationConfig = Engine_ChatEngineConfig_AgentCreationConfig_ToProto(mapCtx, in.AgentCreationConfig)
	out.DialogflowAgentToLink = direct.ValueOf(in.DialogflowAgentToLink)
	return out
}
func Engine_ChatEngineConfig_AgentCreationConfig_FromProto(mapCtx *direct.MapContext, in *pb.Engine_ChatEngineConfig_AgentCreationConfig) *krmv1alpha1.Engine_ChatEngineConfig_AgentCreationConfig {
	if in == nil {
		return nil
	}
	out := &krmv1alpha1.Engine_ChatEngineConfig_AgentCreationConfig{}
	out.Business = direct.LazyPtr(in.GetBusiness())
	out.DefaultLanguageCode = direct.LazyPtr(in.GetDefaultLanguageCode())
	out.TimeZone = direct.LazyPtr(in.GetTimeZone())
	out.Location = direct.LazyPtr(in.GetLocation())
	return out
}
func Engine_ChatEngineConfig_AgentCreationConfig_ToProto(mapCtx *direct.MapContext, in *krmv1alpha1.Engine_ChatEngineConfig_AgentCreationConfig) *pb.Engine_ChatEngineConfig_AgentCreationConfig {
	if in == nil {
		return nil
	}
	out := &pb.Engine_ChatEngineConfig_AgentCreationConfig{}
	out.Business = direct.ValueOf(in.Business)
	out.DefaultLanguageCode = direct.ValueOf(in.DefaultLanguageCode)
	out.TimeZone = direct.ValueOf(in.TimeZone)
	out.Location = direct.ValueOf(in.Location)
	return out
}
func Engine_CommonConfig_FromProto(mapCtx *direct.MapContext, in *pb.Engine_CommonConfig) *krmv1alpha1.Engine_CommonConfig {
	if in == nil {
		return nil
	}
	out := &krmv1alpha1.Engine_CommonConfig{}
	out.CompanyName = direct.LazyPtr(in.GetCompanyName())
	return out
}
func Engine_CommonConfig_ToProto(mapCtx *direct.MapContext, in *krmv1alpha1.Engine_CommonConfig) *pb.Engine_CommonConfig {
	if in == nil {
		return nil
	}
	out := &pb.Engine_CommonConfig{}
	out.CompanyName = direct.ValueOf(in.CompanyName)
	return out
}
func Engine_MediaRecommendationEngineConfig_FromProto(mapCtx *direct.MapContext, in *pb.Engine_MediaRecommendationEngineConfig) *krmv1alpha1.Engine_MediaRecommendationEngineConfig {
	if in == nil {
		return nil
	}
	out := &krmv1alpha1.Engine_MediaRecommendationEngineConfig{}
	out.Type = direct.LazyPtr(in.GetType())
	out.OptimizationObjective = direct.LazyPtr(in.GetOptimizationObjective())
	out.OptimizationObjectiveConfig = Engine_MediaRecommendationEngineConfig_OptimizationObjectiveConfig_FromProto(mapCtx, in.GetOptimizationObjectiveConfig())
	out.TrainingState = direct.Enum_FromProto(mapCtx, in.GetTrainingState())
	return out
}
func Engine_MediaRecommendationEngineConfig_ToProto(mapCtx *direct.MapContext, in *krmv1alpha1.Engine_MediaRecommendationEngineConfig) *pb.Engine_MediaRecommendationEngineConfig {
	if in == nil {
		return nil
	}
	out := &pb.Engine_MediaRecommendationEngineConfig{}
	out.Type = direct.ValueOf(in.Type)
	out.OptimizationObjective = direct.ValueOf(in.OptimizationObjective)
	out.OptimizationObjectiveConfig = Engine_MediaRecommendationEngineConfig_OptimizationObjectiveConfig_ToProto(mapCtx, in.OptimizationObjectiveConfig)
	out.TrainingState = direct.Enum_ToProto[pb.Engine_MediaRecommendationEngineConfig_TrainingState](mapCtx, in.TrainingState)
	return out
}
func Engine_MediaRecommendationEngineConfig_OptimizationObjectiveConfig_FromProto(mapCtx *direct.MapContext, in *pb.Engine_MediaRecommendationEngineConfig_OptimizationObjectiveConfig) *krmv1alpha1.Engine_MediaRecommendationEngineConfig_OptimizationObjectiveConfig {
	if in == nil {
		return nil
	}
	out := &krmv1alpha1.Engine_MediaRecommendationEngineConfig_OptimizationObjectiveConfig{}
	out.TargetField = direct.LazyPtr(in.GetTargetField())
	out.TargetFieldValueFloat = direct.LazyPtr(in.GetTargetFieldValueFloat())
	return out
}
func Engine_MediaRecommendationEngineConfig_OptimizationObjectiveConfig_ToProto(mapCtx *direct.MapContext, in *krmv1alpha1.Engine_MediaRecommendationEngineConfig_OptimizationObjectiveConfig) *pb.Engine_MediaRecommendationEngineConfig_OptimizationObjectiveConfig {
	if in == nil {
		return nil
	}
	out := &pb.Engine_MediaRecommendationEngineConfig_OptimizationObjectiveConfig{}
	out.TargetField = direct.ValueOf(in.TargetField)
	out.TargetFieldValueFloat = direct.ValueOf(in.TargetFieldValueFloat)
	return out
}
func Engine_SearchEngineConfig_FromProto(mapCtx *direct.MapContext, in *pb.Engine_SearchEngineConfig) *krmv1alpha1.Engine_SearchEngineConfig {
	if in == nil {
		return nil
	}
	out := &krmv1alpha1.Engine_SearchEngineConfig{}
	out.SearchTier = direct.Enum_FromProto(mapCtx, in.GetSearchTier())
	out.SearchAddOns = direct.EnumSlice_FromProto(mapCtx, in.SearchAddOns)
	return out
}
func Engine_SearchEngineConfig_ToProto(mapCtx *direct.MapContext, in *krmv1alpha1.Engine_SearchEngineConfig) *pb.Engine_SearchEngineConfig {
	if in == nil {
		return nil
	}
	out := &pb.Engine_SearchEngineConfig{}
	out.SearchTier = direct.Enum_ToProto[pb.SearchTier](mapCtx, in.SearchTier)
	out.SearchAddOns = direct.EnumSlice_ToProto[pb.SearchAddOn](mapCtx, in.SearchAddOns)
	return out
}
func Engine_SimilarDocumentsEngineConfig_FromProto(mapCtx *direct.MapContext, in *pb.Engine_SimilarDocumentsEngineConfig) *krmv1alpha1.Engine_SimilarDocumentsEngineConfig {
	if in == nil {
		return nil
	}
	out := &krmv1alpha1.Engine_SimilarDocumentsEngineConfig{}
	return out
}
func Engine_SimilarDocumentsEngineConfig_ToProto(mapCtx *direct.MapContext, in *krmv1alpha1.Engine_SimilarDocumentsEngineConfig) *pb.Engine_SimilarDocumentsEngineConfig {
	if in == nil {
		return nil
	}
	out := &pb.Engine_SimilarDocumentsEngineConfig{}
	return out
}
func FieldConfig_FromProto(mapCtx *direct.MapContext, in *pb.FieldConfig) *krmv1alpha1.FieldConfig {
	if in == nil {
		return nil
	}
	out := &krmv1alpha1.FieldConfig{}
	out.FieldPath = direct.LazyPtr(in.GetFieldPath())
	// MISSING: FieldType
	out.IndexableOption = direct.Enum_FromProto(mapCtx, in.GetIndexableOption())
	out.DynamicFacetableOption = direct.Enum_FromProto(mapCtx, in.GetDynamicFacetableOption())
	out.SearchableOption = direct.Enum_FromProto(mapCtx, in.GetSearchableOption())
	out.RetrievableOption = direct.Enum_FromProto(mapCtx, in.GetRetrievableOption())
	out.CompletableOption = direct.Enum_FromProto(mapCtx, in.GetCompletableOption())
	out.RecsFilterableOption = direct.Enum_FromProto(mapCtx, in.GetRecsFilterableOption())
	// MISSING: KeyPropertyType
	out.AdvancedSiteSearchDataSources = direct.EnumSlice_FromProto(mapCtx, in.AdvancedSiteSearchDataSources)
	out.SchemaOrgPaths = in.SchemaOrgPaths
	return out
}
func FieldConfig_ToProto(mapCtx *direct.MapContext, in *krmv1alpha1.FieldConfig) *pb.FieldConfig {
	if in == nil {
		return nil
	}
	out := &pb.FieldConfig{}
	out.FieldPath = direct.ValueOf(in.FieldPath)
	// MISSING: FieldType
	out.IndexableOption = direct.Enum_ToProto[pb.FieldConfig_IndexableOption](mapCtx, in.IndexableOption)
	out.DynamicFacetableOption = direct.Enum_ToProto[pb.FieldConfig_DynamicFacetableOption](mapCtx, in.DynamicFacetableOption)
	out.SearchableOption = direct.Enum_ToProto[pb.FieldConfig_SearchableOption](mapCtx, in.SearchableOption)
	out.RetrievableOption = direct.Enum_ToProto[pb.FieldConfig_RetrievableOption](mapCtx, in.RetrievableOption)
	out.CompletableOption = direct.Enum_ToProto[pb.FieldConfig_CompletableOption](mapCtx, in.CompletableOption)
	out.RecsFilterableOption = direct.Enum_ToProto[pb.FieldConfig_FilterableOption](mapCtx, in.RecsFilterableOption)
	// MISSING: KeyPropertyType
	out.AdvancedSiteSearchDataSources = direct.EnumSlice_ToProto[pb.FieldConfig_AdvancedSiteSearchDataSource](mapCtx, in.AdvancedSiteSearchDataSources)
	out.SchemaOrgPaths = in.SchemaOrgPaths
	return out
}
func FieldConfigObservedState_FromProto(mapCtx *direct.MapContext, in *pb.FieldConfig) *krmv1alpha1.FieldConfigObservedState {
	if in == nil {
		return nil
	}
	out := &krmv1alpha1.FieldConfigObservedState{}
	// MISSING: FieldPath
	out.FieldType = direct.Enum_FromProto(mapCtx, in.GetFieldType())
	// MISSING: IndexableOption
	// MISSING: DynamicFacetableOption
	// MISSING: SearchableOption
	// MISSING: RetrievableOption
	// MISSING: CompletableOption
	// MISSING: RecsFilterableOption
	out.KeyPropertyType = direct.LazyPtr(in.GetKeyPropertyType())
	// MISSING: AdvancedSiteSearchDataSources
	// MISSING: SchemaOrgPaths
	return out
}
func FieldConfigObservedState_ToProto(mapCtx *direct.MapContext, in *krmv1alpha1.FieldConfigObservedState) *pb.FieldConfig {
	if in == nil {
		return nil
	}
	out := &pb.FieldConfig{}
	// MISSING: FieldPath
	out.FieldType = direct.Enum_ToProto[pb.FieldConfig_FieldType](mapCtx, in.FieldType)
	// MISSING: IndexableOption
	// MISSING: DynamicFacetableOption
	// MISSING: SearchableOption
	// MISSING: RetrievableOption
	// MISSING: CompletableOption
	// MISSING: RecsFilterableOption
	out.KeyPropertyType = direct.ValueOf(in.KeyPropertyType)
	// MISSING: AdvancedSiteSearchDataSources
	// MISSING: SchemaOrgPaths
	return out
}
func IdpConfig_FromProto(mapCtx *direct.MapContext, in *pb.IdpConfig) *krmv1alpha1.IdpConfig {
	if in == nil {
		return nil
	}
	out := &krmv1alpha1.IdpConfig{}
	out.IdpType = direct.Enum_FromProto(mapCtx, in.GetIdpType())
	out.ExternalIdpConfig = IdpConfig_ExternalIdpConfig_FromProto(mapCtx, in.GetExternalIdpConfig())
	return out
}
func IdpConfig_ToProto(mapCtx *direct.MapContext, in *krmv1alpha1.IdpConfig) *pb.IdpConfig {
	if in == nil {
		return nil
	}
	out := &pb.IdpConfig{}
	out.IdpType = direct.Enum_ToProto[pb.IdpConfig_IdpType](mapCtx, in.IdpType)
	out.ExternalIdpConfig = IdpConfig_ExternalIdpConfig_ToProto(mapCtx, in.ExternalIdpConfig)
	return out
}
func IdpConfig_ExternalIdpConfig_FromProto(mapCtx *direct.MapContext, in *pb.IdpConfig_ExternalIdpConfig) *krmv1alpha1.IdpConfig_ExternalIdpConfig {
	if in == nil {
		return nil
	}
	out := &krmv1alpha1.IdpConfig_ExternalIdpConfig{}
	out.WorkforcePoolName = direct.LazyPtr(in.GetWorkforcePoolName())
	return out
}
func IdpConfig_ExternalIdpConfig_ToProto(mapCtx *direct.MapContext, in *krmv1alpha1.IdpConfig_ExternalIdpConfig) *pb.IdpConfig_ExternalIdpConfig {
	if in == nil {
		return nil
	}
	out := &pb.IdpConfig_ExternalIdpConfig{}
	out.WorkforcePoolName = direct.ValueOf(in.WorkforcePoolName)
	return out
}
func LanguageInfo_FromProto(mapCtx *direct.MapContext, in *pb.LanguageInfo) *krmv1alpha1.LanguageInfo {
	if in == nil {
		return nil
	}
	out := &krmv1alpha1.LanguageInfo{}
	out.LanguageCode = direct.LazyPtr(in.GetLanguageCode())
	// MISSING: NormalizedLanguageCode
	// MISSING: Language
	// MISSING: Region
	return out
}
func LanguageInfo_ToProto(mapCtx *direct.MapContext, in *krmv1alpha1.LanguageInfo) *pb.LanguageInfo {
	if in == nil {
		return nil
	}
	out := &pb.LanguageInfo{}
	out.LanguageCode = direct.ValueOf(in.LanguageCode)
	// MISSING: NormalizedLanguageCode
	// MISSING: Language
	// MISSING: Region
	return out
}
func LanguageInfoObservedState_FromProto(mapCtx *direct.MapContext, in *pb.LanguageInfo) *krmv1alpha1.LanguageInfoObservedState {
	if in == nil {
		return nil
	}
	out := &krmv1alpha1.LanguageInfoObservedState{}
	// MISSING: LanguageCode
	out.NormalizedLanguageCode = direct.LazyPtr(in.GetNormalizedLanguageCode())
	out.Language = direct.LazyPtr(in.GetLanguage())
	out.Region = direct.LazyPtr(in.GetRegion())
	return out
}
func LanguageInfoObservedState_ToProto(mapCtx *direct.MapContext, in *krmv1alpha1.LanguageInfoObservedState) *pb.LanguageInfo {
	if in == nil {
		return nil
	}
	out := &pb.LanguageInfo{}
	// MISSING: LanguageCode
	out.NormalizedLanguageCode = direct.ValueOf(in.NormalizedLanguageCode)
	out.Language = direct.ValueOf(in.Language)
	out.Region = direct.ValueOf(in.Region)
	return out
}
func Schema_FromProto(mapCtx *direct.MapContext, in *pb.Schema) *krmv1alpha1.Schema {
	if in == nil {
		return nil
	}
	out := &krmv1alpha1.Schema{}
	out.StructSchema = StructSchema_FromProto(mapCtx, in.GetStructSchema())
	out.JsonSchema = direct.LazyPtr(in.GetJsonSchema())
	out.Name = direct.LazyPtr(in.GetName())
	// MISSING: FieldConfigs
	return out
}
func Schema_ToProto(mapCtx *direct.MapContext, in *krmv1alpha1.Schema) *pb.Schema {
	if in == nil {
		return nil
	}
	out := &pb.Schema{}
	if oneof := StructSchema_ToProto(mapCtx, in.StructSchema); oneof != nil {
		out.Schema = &pb.Schema_StructSchema{StructSchema: oneof}
	}
	if oneof := Schema_JsonSchema_ToProto(mapCtx, in.JsonSchema); oneof != nil {
		out.Schema = oneof
	}
	out.Name = direct.ValueOf(in.Name)
	// MISSING: FieldConfigs
	return out
}
func SchemaObservedState_FromProto(mapCtx *direct.MapContext, in *pb.Schema) *krmv1alpha1.SchemaObservedState {
	if in == nil {
		return nil
	}
	out := &krmv1alpha1.SchemaObservedState{}
	// MISSING: StructSchema
	// MISSING: JsonSchema
	// MISSING: Name
	out.FieldConfigs = direct.Slice_FromProto(mapCtx, in.FieldConfigs, FieldConfig_FromProto)
	return out
}
func SchemaObservedState_ToProto(mapCtx *direct.MapContext, in *krmv1alpha1.SchemaObservedState) *pb.Schema {
	if in == nil {
		return nil
	}
	out := &pb.Schema{}
	// MISSING: StructSchema
	// MISSING: JsonSchema
	// MISSING: Name
	out.FieldConfigs = direct.Slice_ToProto(mapCtx, in.FieldConfigs, FieldConfig_ToProto)
	return out
}
func SiteVerificationInfo_FromProto(mapCtx *direct.MapContext, in *pb.SiteVerificationInfo) *krmv1alpha1.SiteVerificationInfo {
	if in == nil {
		return nil
	}
	out := &krmv1alpha1.SiteVerificationInfo{}
	out.SiteVerificationState = direct.Enum_FromProto(mapCtx, in.GetSiteVerificationState())
	out.VerifyTime = direct.StringTimestamp_FromProto(mapCtx, in.GetVerifyTime())
	return out
}
func SiteVerificationInfo_ToProto(mapCtx *direct.MapContext, in *krmv1alpha1.SiteVerificationInfo) *pb.SiteVerificationInfo {
	if in == nil {
		return nil
	}
	out := &pb.SiteVerificationInfo{}
	out.SiteVerificationState = direct.Enum_ToProto[pb.SiteVerificationInfo_SiteVerificationState](mapCtx, in.SiteVerificationState)
	out.VerifyTime = direct.StringTimestamp_ToProto(mapCtx, in.VerifyTime)
	return out
}
func TargetSite_FromProto(mapCtx *direct.MapContext, in *pb.TargetSite) *krmv1alpha1.TargetSite {
	if in == nil {
		return nil
	}
	out := &krmv1alpha1.TargetSite{}
	// MISSING: Name
	out.ProvidedURIPattern = direct.LazyPtr(in.GetProvidedUriPattern())
	out.Type = direct.Enum_FromProto(mapCtx, in.GetType())
	out.ExactMatch = direct.LazyPtr(in.GetExactMatch())
	// MISSING: GeneratedURIPattern
	// MISSING: RootDomainURI
	// MISSING: SiteVerificationInfo
	// MISSING: IndexingStatus
	// MISSING: UpdateTime
	// MISSING: FailureReason
	return out
}
func TargetSite_ToProto(mapCtx *direct.MapContext, in *krmv1alpha1.TargetSite) *pb.TargetSite {
	if in == nil {
		return nil
	}
	out := &pb.TargetSite{}
	// MISSING: Name
	out.ProvidedUriPattern = direct.ValueOf(in.ProvidedURIPattern)
	out.Type = direct.Enum_ToProto[pb.TargetSite_Type](mapCtx, in.Type)
	out.ExactMatch = direct.ValueOf(in.ExactMatch)
	// MISSING: GeneratedURIPattern
	// MISSING: RootDomainURI
	// MISSING: SiteVerificationInfo
	// MISSING: IndexingStatus
	// MISSING: UpdateTime
	// MISSING: FailureReason
	return out
}
func TargetSiteObservedState_FromProto(mapCtx *direct.MapContext, in *pb.TargetSite) *krmv1alpha1.TargetSiteObservedState {
	if in == nil {
		return nil
	}
	out := &krmv1alpha1.TargetSiteObservedState{}
	out.Name = direct.LazyPtr(in.GetName())
	// MISSING: ProvidedURIPattern
	// MISSING: Type
	// MISSING: ExactMatch
	out.GeneratedURIPattern = direct.LazyPtr(in.GetGeneratedUriPattern())
	out.RootDomainURI = direct.LazyPtr(in.GetRootDomainUri())
	out.SiteVerificationInfo = SiteVerificationInfo_FromProto(mapCtx, in.GetSiteVerificationInfo())
	out.IndexingStatus = direct.Enum_FromProto(mapCtx, in.GetIndexingStatus())
	out.UpdateTime = direct.StringTimestamp_FromProto(mapCtx, in.GetUpdateTime())
	out.FailureReason = TargetSite_FailureReason_FromProto(mapCtx, in.GetFailureReason())
	return out
}
func TargetSiteObservedState_ToProto(mapCtx *direct.MapContext, in *krmv1alpha1.TargetSiteObservedState) *pb.TargetSite {
	if in == nil {
		return nil
	}
	out := &pb.TargetSite{}
	out.Name = direct.ValueOf(in.Name)
	// MISSING: ProvidedURIPattern
	// MISSING: Type
	// MISSING: ExactMatch
	out.GeneratedUriPattern = direct.ValueOf(in.GeneratedURIPattern)
	out.RootDomainUri = direct.ValueOf(in.RootDomainURI)
	out.SiteVerificationInfo = SiteVerificationInfo_ToProto(mapCtx, in.SiteVerificationInfo)
	out.IndexingStatus = direct.Enum_ToProto[pb.TargetSite_IndexingStatus](mapCtx, in.IndexingStatus)
	out.UpdateTime = direct.StringTimestamp_ToProto(mapCtx, in.UpdateTime)
	out.FailureReason = TargetSite_FailureReason_ToProto(mapCtx, in.FailureReason)
	return out
}
func TargetSite_FailureReason_FromProto(mapCtx *direct.MapContext, in *pb.TargetSite_FailureReason) *krmv1alpha1.TargetSite_FailureReason {
	if in == nil {
		return nil
	}
	out := &krmv1alpha1.TargetSite_FailureReason{}
	out.QuotaFailure = TargetSite_FailureReason_QuotaFailure_FromProto(mapCtx, in.GetQuotaFailure())
	return out
}
func TargetSite_FailureReason_ToProto(mapCtx *direct.MapContext, in *krmv1alpha1.TargetSite_FailureReason) *pb.TargetSite_FailureReason {
	if in == nil {
		return nil
	}
	out := &pb.TargetSite_FailureReason{}
	if oneof := TargetSite_FailureReason_QuotaFailure_ToProto(mapCtx, in.QuotaFailure); oneof != nil {
		out.Failure = &pb.TargetSite_FailureReason_QuotaFailure_{QuotaFailure: oneof}
	}
	return out
}
func TargetSite_FailureReason_QuotaFailure_FromProto(mapCtx *direct.MapContext, in *pb.TargetSite_FailureReason_QuotaFailure) *krmv1alpha1.TargetSite_FailureReason_QuotaFailure {
	if in == nil {
		return nil
	}
	out := &krmv1alpha1.TargetSite_FailureReason_QuotaFailure{}
	out.TotalRequiredQuota = direct.LazyPtr(in.GetTotalRequiredQuota())
	return out
}
func TargetSite_FailureReason_QuotaFailure_ToProto(mapCtx *direct.MapContext, in *krmv1alpha1.TargetSite_FailureReason_QuotaFailure) *pb.TargetSite_FailureReason_QuotaFailure {
	if in == nil {
		return nil
	}
	out := &pb.TargetSite_FailureReason_QuotaFailure{}
	out.TotalRequiredQuota = direct.ValueOf(in.TotalRequiredQuota)
	return out
}
func WorkspaceConfig_FromProto(mapCtx *direct.MapContext, in *pb.WorkspaceConfig) *krmv1alpha1.WorkspaceConfig {
	if in == nil {
		return nil
	}
	out := &krmv1alpha1.WorkspaceConfig{}
	out.Type = direct.Enum_FromProto(mapCtx, in.GetType())
	out.DasherCustomerID = direct.LazyPtr(in.GetDasherCustomerId())
	return out
}
func WorkspaceConfig_ToProto(mapCtx *direct.MapContext, in *krmv1alpha1.WorkspaceConfig) *pb.WorkspaceConfig {
	if in == nil {
		return nil
	}
	out := &pb.WorkspaceConfig{}
	out.Type = direct.Enum_ToProto[pb.WorkspaceConfig_Type](mapCtx, in.Type)
	out.DasherCustomerId = direct.ValueOf(in.DasherCustomerID)
	return out
}
