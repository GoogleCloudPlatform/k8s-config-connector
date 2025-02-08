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
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct"
	refs "github.com/GoogleCloudPlatform/k8s-config-connector/apis/refs/v1beta1"
	pb "cloud.google.com/go/discoveryengine/apiv1beta/discoveryenginepb"
	krm "github.com/GoogleCloudPlatform/k8s-config-connector/apis/discoveryengine/v1alpha1"
)
func DataStore_FromProto(mapCtx *direct.MapContext, in *pb.DataStore) *krm.DataStore {
	if in == nil {
		return nil
	}
	out := &krm.DataStore{}
	out.Name = direct.LazyPtr(in.GetName())
	out.DisplayName = direct.LazyPtr(in.GetDisplayName())
	out.IndustryVertical = direct.Enum_FromProto(mapCtx, in.GetIndustryVertical())
	out.SolutionTypes = direct.EnumSlice_FromProto(mapCtx, in.SolutionTypes)
	// MISSING: DefaultSchemaID
	out.ContentConfig = direct.Enum_FromProto(mapCtx, in.GetContentConfig())
	// MISSING: CreateTime
	out.LanguageInfo = LanguageInfo_FromProto(mapCtx, in.GetLanguageInfo())
	out.NaturalLanguageQueryUnderstandingConfig = NaturalLanguageQueryUnderstandingConfig_FromProto(mapCtx, in.GetNaturalLanguageQueryUnderstandingConfig())
	// MISSING: BillingEstimation
	out.WorkspaceConfig = WorkspaceConfig_FromProto(mapCtx, in.GetWorkspaceConfig())
	out.DocumentProcessingConfig = DocumentProcessingConfig_FromProto(mapCtx, in.GetDocumentProcessingConfig())
	out.StartingSchema = Schema_FromProto(mapCtx, in.GetStartingSchema())
	out.ServingConfigDataStore = DataStore_ServingConfigDataStore_FromProto(mapCtx, in.GetServingConfigDataStore())
	return out
}
func DataStore_ToProto(mapCtx *direct.MapContext, in *krm.DataStore) *pb.DataStore {
	if in == nil {
		return nil
	}
	out := &pb.DataStore{}
	out.Name = direct.ValueOf(in.Name)
	out.DisplayName = direct.ValueOf(in.DisplayName)
	out.IndustryVertical = direct.Enum_ToProto[pb.IndustryVertical](mapCtx, in.IndustryVertical)
	out.SolutionTypes = direct.EnumSlice_ToProto[pb.SolutionType](mapCtx, in.SolutionTypes)
	// MISSING: DefaultSchemaID
	out.ContentConfig = direct.Enum_ToProto[pb.DataStore_ContentConfig](mapCtx, in.ContentConfig)
	// MISSING: CreateTime
	out.LanguageInfo = LanguageInfo_ToProto(mapCtx, in.LanguageInfo)
	out.NaturalLanguageQueryUnderstandingConfig = NaturalLanguageQueryUnderstandingConfig_ToProto(mapCtx, in.NaturalLanguageQueryUnderstandingConfig)
	// MISSING: BillingEstimation
	out.WorkspaceConfig = WorkspaceConfig_ToProto(mapCtx, in.WorkspaceConfig)
	out.DocumentProcessingConfig = DocumentProcessingConfig_ToProto(mapCtx, in.DocumentProcessingConfig)
	out.StartingSchema = Schema_ToProto(mapCtx, in.StartingSchema)
	out.ServingConfigDataStore = DataStore_ServingConfigDataStore_ToProto(mapCtx, in.ServingConfigDataStore)
	return out
}
func DataStoreObservedState_FromProto(mapCtx *direct.MapContext, in *pb.DataStore) *krm.DataStoreObservedState {
	if in == nil {
		return nil
	}
	out := &krm.DataStoreObservedState{}
	// MISSING: Name
	// MISSING: DisplayName
	// MISSING: IndustryVertical
	// MISSING: SolutionTypes
	out.DefaultSchemaID = direct.LazyPtr(in.GetDefaultSchemaId())
	// MISSING: ContentConfig
	out.CreateTime = direct.StringTimestamp_FromProto(mapCtx, in.GetCreateTime())
	out.LanguageInfo = LanguageInfoObservedState_FromProto(mapCtx, in.GetLanguageInfo())
	// MISSING: NaturalLanguageQueryUnderstandingConfig
	out.BillingEstimation = DataStore_BillingEstimation_FromProto(mapCtx, in.GetBillingEstimation())
	// MISSING: WorkspaceConfig
	// MISSING: DocumentProcessingConfig
	// MISSING: StartingSchema
	// MISSING: ServingConfigDataStore
	return out
}
func DataStoreObservedState_ToProto(mapCtx *direct.MapContext, in *krm.DataStoreObservedState) *pb.DataStore {
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
	// MISSING: NaturalLanguageQueryUnderstandingConfig
	out.BillingEstimation = DataStore_BillingEstimation_ToProto(mapCtx, in.BillingEstimation)
	// MISSING: WorkspaceConfig
	// MISSING: DocumentProcessingConfig
	// MISSING: StartingSchema
	// MISSING: ServingConfigDataStore
	return out
}
func DataStore_BillingEstimation_FromProto(mapCtx *direct.MapContext, in *pb.DataStore_BillingEstimation) *krm.DataStore_BillingEstimation {
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
func DataStore_BillingEstimation_ToProto(mapCtx *direct.MapContext, in *krm.DataStore_BillingEstimation) *pb.DataStore_BillingEstimation {
	if in == nil {
		return nil
	}
	out := &pb.DataStore_BillingEstimation{}
	out.StructuredDataSize = direct.ValueOf(in.StructuredDataSize)
	out.UnstructuredDataSize = direct.ValueOf(in.UnstructuredDataSize)
	out.WebsiteDataSize = direct.ValueOf(in.WebsiteDataSize)
	out.StructuredDataUpdateTime = direct.StringTimestamp_ToProto(mapCtx, in.StructuredDataUpdateTime)
	out.UnstructuredDataUpdateTime = direct.StringTimestamp_ToProto(mapCtx, in.UnstructuredDataUpdateTime)
	out.WebsiteDataUpdateTime = direct.StringTimestamp_ToProto(mapCtx, in.WebsiteDataUpdateTime)
	return out
}
func DataStore_ServingConfigDataStore_FromProto(mapCtx *direct.MapContext, in *pb.DataStore_ServingConfigDataStore) *krm.DataStore_ServingConfigDataStore {
	if in == nil {
		return nil
	}
	out := &krm.DataStore_ServingConfigDataStore{}
	out.DisabledForServing = direct.LazyPtr(in.GetDisabledForServing())
	return out
}
func DataStore_ServingConfigDataStore_ToProto(mapCtx *direct.MapContext, in *krm.DataStore_ServingConfigDataStore) *pb.DataStore_ServingConfigDataStore {
	if in == nil {
		return nil
	}
	out := &pb.DataStore_ServingConfigDataStore{}
	out.DisabledForServing = direct.ValueOf(in.DisabledForServing)
	return out
}
func DocumentProcessingConfig_FromProto(mapCtx *direct.MapContext, in *pb.DocumentProcessingConfig) *krm.DocumentProcessingConfig {
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
func DocumentProcessingConfig_ToProto(mapCtx *direct.MapContext, in *krm.DocumentProcessingConfig) *pb.DocumentProcessingConfig {
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
func DocumentProcessingConfig_ChunkingConfig_FromProto(mapCtx *direct.MapContext, in *pb.DocumentProcessingConfig_ChunkingConfig) *krm.DocumentProcessingConfig_ChunkingConfig {
	if in == nil {
		return nil
	}
	out := &krm.DocumentProcessingConfig_ChunkingConfig{}
	out.LayoutBasedChunkingConfig = DocumentProcessingConfig_ChunkingConfig_LayoutBasedChunkingConfig_FromProto(mapCtx, in.GetLayoutBasedChunkingConfig())
	return out
}
func DocumentProcessingConfig_ChunkingConfig_ToProto(mapCtx *direct.MapContext, in *krm.DocumentProcessingConfig_ChunkingConfig) *pb.DocumentProcessingConfig_ChunkingConfig {
	if in == nil {
		return nil
	}
	out := &pb.DocumentProcessingConfig_ChunkingConfig{}
	if oneof := DocumentProcessingConfig_ChunkingConfig_LayoutBasedChunkingConfig_ToProto(mapCtx, in.LayoutBasedChunkingConfig); oneof != nil {
		out.ChunkMode = &pb.DocumentProcessingConfig_ChunkingConfig_LayoutBasedChunkingConfig_{LayoutBasedChunkingConfig: oneof}
	}
	return out
}
func DocumentProcessingConfig_ChunkingConfig_LayoutBasedChunkingConfig_FromProto(mapCtx *direct.MapContext, in *pb.DocumentProcessingConfig_ChunkingConfig_LayoutBasedChunkingConfig) *krm.DocumentProcessingConfig_ChunkingConfig_LayoutBasedChunkingConfig {
	if in == nil {
		return nil
	}
	out := &krm.DocumentProcessingConfig_ChunkingConfig_LayoutBasedChunkingConfig{}
	out.ChunkSize = direct.LazyPtr(in.GetChunkSize())
	out.IncludeAncestorHeadings = direct.LazyPtr(in.GetIncludeAncestorHeadings())
	return out
}
func DocumentProcessingConfig_ChunkingConfig_LayoutBasedChunkingConfig_ToProto(mapCtx *direct.MapContext, in *krm.DocumentProcessingConfig_ChunkingConfig_LayoutBasedChunkingConfig) *pb.DocumentProcessingConfig_ChunkingConfig_LayoutBasedChunkingConfig {
	if in == nil {
		return nil
	}
	out := &pb.DocumentProcessingConfig_ChunkingConfig_LayoutBasedChunkingConfig{}
	out.ChunkSize = direct.ValueOf(in.ChunkSize)
	out.IncludeAncestorHeadings = direct.ValueOf(in.IncludeAncestorHeadings)
	return out
}
func DocumentProcessingConfig_ParsingConfig_FromProto(mapCtx *direct.MapContext, in *pb.DocumentProcessingConfig_ParsingConfig) *krm.DocumentProcessingConfig_ParsingConfig {
	if in == nil {
		return nil
	}
	out := &krm.DocumentProcessingConfig_ParsingConfig{}
	out.DigitalParsingConfig = DocumentProcessingConfig_ParsingConfig_DigitalParsingConfig_FromProto(mapCtx, in.GetDigitalParsingConfig())
	out.OcrParsingConfig = DocumentProcessingConfig_ParsingConfig_OcrParsingConfig_FromProto(mapCtx, in.GetOcrParsingConfig())
	out.LayoutParsingConfig = DocumentProcessingConfig_ParsingConfig_LayoutParsingConfig_FromProto(mapCtx, in.GetLayoutParsingConfig())
	return out
}
func DocumentProcessingConfig_ParsingConfig_ToProto(mapCtx *direct.MapContext, in *krm.DocumentProcessingConfig_ParsingConfig) *pb.DocumentProcessingConfig_ParsingConfig {
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
func DocumentProcessingConfig_ParsingConfig_DigitalParsingConfig_FromProto(mapCtx *direct.MapContext, in *pb.DocumentProcessingConfig_ParsingConfig_DigitalParsingConfig) *krm.DocumentProcessingConfig_ParsingConfig_DigitalParsingConfig {
	if in == nil {
		return nil
	}
	out := &krm.DocumentProcessingConfig_ParsingConfig_DigitalParsingConfig{}
	return out
}
func DocumentProcessingConfig_ParsingConfig_DigitalParsingConfig_ToProto(mapCtx *direct.MapContext, in *krm.DocumentProcessingConfig_ParsingConfig_DigitalParsingConfig) *pb.DocumentProcessingConfig_ParsingConfig_DigitalParsingConfig {
	if in == nil {
		return nil
	}
	out := &pb.DocumentProcessingConfig_ParsingConfig_DigitalParsingConfig{}
	return out
}
func DocumentProcessingConfig_ParsingConfig_LayoutParsingConfig_FromProto(mapCtx *direct.MapContext, in *pb.DocumentProcessingConfig_ParsingConfig_LayoutParsingConfig) *krm.DocumentProcessingConfig_ParsingConfig_LayoutParsingConfig {
	if in == nil {
		return nil
	}
	out := &krm.DocumentProcessingConfig_ParsingConfig_LayoutParsingConfig{}
	return out
}
func DocumentProcessingConfig_ParsingConfig_LayoutParsingConfig_ToProto(mapCtx *direct.MapContext, in *krm.DocumentProcessingConfig_ParsingConfig_LayoutParsingConfig) *pb.DocumentProcessingConfig_ParsingConfig_LayoutParsingConfig {
	if in == nil {
		return nil
	}
	out := &pb.DocumentProcessingConfig_ParsingConfig_LayoutParsingConfig{}
	return out
}
func DocumentProcessingConfig_ParsingConfig_OcrParsingConfig_FromProto(mapCtx *direct.MapContext, in *pb.DocumentProcessingConfig_ParsingConfig_OcrParsingConfig) *krm.DocumentProcessingConfig_ParsingConfig_OcrParsingConfig {
	if in == nil {
		return nil
	}
	out := &krm.DocumentProcessingConfig_ParsingConfig_OcrParsingConfig{}
	out.EnhancedDocumentElements = in.EnhancedDocumentElements
	out.UseNativeText = direct.LazyPtr(in.GetUseNativeText())
	return out
}
func DocumentProcessingConfig_ParsingConfig_OcrParsingConfig_ToProto(mapCtx *direct.MapContext, in *krm.DocumentProcessingConfig_ParsingConfig_OcrParsingConfig) *pb.DocumentProcessingConfig_ParsingConfig_OcrParsingConfig {
	if in == nil {
		return nil
	}
	out := &pb.DocumentProcessingConfig_ParsingConfig_OcrParsingConfig{}
	out.EnhancedDocumentElements = in.EnhancedDocumentElements
	out.UseNativeText = direct.ValueOf(in.UseNativeText)
	return out
}
func LanguageInfo_FromProto(mapCtx *direct.MapContext, in *pb.LanguageInfo) *krm.LanguageInfo {
	if in == nil {
		return nil
	}
	out := &krm.LanguageInfo{}
	out.LanguageCode = direct.LazyPtr(in.GetLanguageCode())
	// MISSING: NormalizedLanguageCode
	// MISSING: Language
	// MISSING: Region
	return out
}
func LanguageInfo_ToProto(mapCtx *direct.MapContext, in *krm.LanguageInfo) *pb.LanguageInfo {
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
func LanguageInfoObservedState_FromProto(mapCtx *direct.MapContext, in *pb.LanguageInfo) *krm.LanguageInfoObservedState {
	if in == nil {
		return nil
	}
	out := &krm.LanguageInfoObservedState{}
	// MISSING: LanguageCode
	out.NormalizedLanguageCode = direct.LazyPtr(in.GetNormalizedLanguageCode())
	out.Language = direct.LazyPtr(in.GetLanguage())
	out.Region = direct.LazyPtr(in.GetRegion())
	return out
}
func LanguageInfoObservedState_ToProto(mapCtx *direct.MapContext, in *krm.LanguageInfoObservedState) *pb.LanguageInfo {
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
func NaturalLanguageQueryUnderstandingConfig_FromProto(mapCtx *direct.MapContext, in *pb.NaturalLanguageQueryUnderstandingConfig) *krm.NaturalLanguageQueryUnderstandingConfig {
	if in == nil {
		return nil
	}
	out := &krm.NaturalLanguageQueryUnderstandingConfig{}
	out.Mode = direct.Enum_FromProto(mapCtx, in.GetMode())
	return out
}
func NaturalLanguageQueryUnderstandingConfig_ToProto(mapCtx *direct.MapContext, in *krm.NaturalLanguageQueryUnderstandingConfig) *pb.NaturalLanguageQueryUnderstandingConfig {
	if in == nil {
		return nil
	}
	out := &pb.NaturalLanguageQueryUnderstandingConfig{}
	out.Mode = direct.Enum_ToProto[pb.NaturalLanguageQueryUnderstandingConfig_Mode](mapCtx, in.Mode)
	return out
}
func Schema_FromProto(mapCtx *direct.MapContext, in *pb.Schema) *krm.Schema {
	if in == nil {
		return nil
	}
	out := &krm.Schema{}
	out.StructSchema = StructSchema_FromProto(mapCtx, in.GetStructSchema())
	out.JsonSchema = direct.LazyPtr(in.GetJsonSchema())
	out.Name = direct.LazyPtr(in.GetName())
	return out
}
func Schema_ToProto(mapCtx *direct.MapContext, in *krm.Schema) *pb.Schema {
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
	return out
}
func WorkspaceConfig_FromProto(mapCtx *direct.MapContext, in *pb.WorkspaceConfig) *krm.WorkspaceConfig {
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
func WorkspaceConfig_ToProto(mapCtx *direct.MapContext, in *krm.WorkspaceConfig) *pb.WorkspaceConfig {
	if in == nil {
		return nil
	}
	out := &pb.WorkspaceConfig{}
	out.Type = direct.Enum_ToProto[pb.WorkspaceConfig_Type](mapCtx, in.Type)
	out.DasherCustomerId = direct.ValueOf(in.DasherCustomerID)
	out.SuperAdminServiceAccount = direct.ValueOf(in.SuperAdminServiceAccount)
	out.SuperAdminEmailAddress = direct.ValueOf(in.SuperAdminEmailAddress)
	return out
}
