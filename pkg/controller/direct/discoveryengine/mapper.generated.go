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
	pb "cloud.google.com/go/discoveryengine/apiv1/discoveryenginepb"
	krm "github.com/GoogleCloudPlatform/k8s-config-connector/apis/discoveryengine/v1alpha1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct"
)

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
func DiscoveryEngineDataStoreObservedState_FromProto(mapCtx *direct.MapContext, in *pb.DataStore) *krm.DiscoveryEngineDataStoreObservedState {
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
func DiscoveryEngineDataStoreObservedState_ToProto(mapCtx *direct.MapContext, in *krm.DiscoveryEngineDataStoreObservedState) *pb.DataStore {
	if in == nil {
		return nil
	}
	out := &pb.DataStore{}
	// MISSING: Name
	out.DefaultSchemaId = direct.ValueOf(in.DefaultSchemaID)
	out.CreateTime = direct.StringTimestamp_ToProto(mapCtx, in.CreateTime)
	out.BillingEstimation = DataStore_BillingEstimation_ToProto(mapCtx, in.BillingEstimation)
	// MISSING: DocumentProcessingConfig
	// MISSING: StartingSchema
	return out
}
func DiscoveryEngineDataStoreSpec_FromProto(mapCtx *direct.MapContext, in *pb.DataStore) *krm.DiscoveryEngineDataStoreSpec {
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
func DiscoveryEngineDataStoreSpec_ToProto(mapCtx *direct.MapContext, in *krm.DiscoveryEngineDataStoreSpec) *pb.DataStore {
	if in == nil {
		return nil
	}
	out := &pb.DataStore{}
	// MISSING: Name
	out.DisplayName = direct.ValueOf(in.DisplayName)
	out.IndustryVertical = direct.Enum_ToProto[pb.IndustryVertical](mapCtx, in.IndustryVertical)
	out.SolutionTypes = direct.EnumSlice_ToProto[pb.SolutionType](mapCtx, in.SolutionTypes)
	out.ContentConfig = direct.Enum_ToProto[pb.DataStore_ContentConfig](mapCtx, in.ContentConfig)
	out.WorkspaceConfig = WorkspaceConfig_ToProto(mapCtx, in.WorkspaceConfig)
	// MISSING: DocumentProcessingConfig
	// MISSING: StartingSchema
	return out
}
func DiscoveryEngineDataStoreTargetSiteObservedState_FromProto(mapCtx *direct.MapContext, in *pb.TargetSite) *krm.DiscoveryEngineDataStoreTargetSiteObservedState {
	if in == nil {
		return nil
	}
	out := &krm.DiscoveryEngineDataStoreTargetSiteObservedState{}
	// MISSING: Name
	// MISSING: ProvidedURIPattern
	// MISSING: GeneratedURIPattern
	// (near miss): "GeneratedURIPattern" vs "GeneratedUriPattern"
	// MISSING: RootDomainURI
	// (near miss): "RootDomainURI" vs "RootDomainUri"
	out.SiteVerificationInfo = SiteVerificationInfo_FromProto(mapCtx, in.GetSiteVerificationInfo())
	out.IndexingStatus = direct.Enum_FromProto(mapCtx, in.GetIndexingStatus())
	out.UpdateTime = direct.StringTimestamp_FromProto(mapCtx, in.GetUpdateTime())
	out.FailureReason = TargetSite_FailureReason_FromProto(mapCtx, in.GetFailureReason())
	return out
}
func DiscoveryEngineDataStoreTargetSiteObservedState_ToProto(mapCtx *direct.MapContext, in *krm.DiscoveryEngineDataStoreTargetSiteObservedState) *pb.TargetSite {
	if in == nil {
		return nil
	}
	out := &pb.TargetSite{}
	// MISSING: Name
	// MISSING: ProvidedURIPattern
	// MISSING: GeneratedURIPattern
	// (near miss): "GeneratedURIPattern" vs "GeneratedUriPattern"
	// MISSING: RootDomainURI
	// (near miss): "RootDomainURI" vs "RootDomainUri"
	out.SiteVerificationInfo = SiteVerificationInfo_ToProto(mapCtx, in.SiteVerificationInfo)
	out.IndexingStatus = direct.Enum_ToProto[pb.TargetSite_IndexingStatus](mapCtx, in.IndexingStatus)
	out.UpdateTime = direct.StringTimestamp_ToProto(mapCtx, in.UpdateTime)
	out.FailureReason = TargetSite_FailureReason_ToProto(mapCtx, in.FailureReason)
	return out
}
func DiscoveryEngineDataStoreTargetSiteSpec_FromProto(mapCtx *direct.MapContext, in *pb.TargetSite) *krm.DiscoveryEngineDataStoreTargetSiteSpec {
	if in == nil {
		return nil
	}
	out := &krm.DiscoveryEngineDataStoreTargetSiteSpec{}
	// MISSING: Name
	// MISSING: ProvidedURIPattern
	// (near miss): "ProvidedURIPattern" vs "ProvidedUriPattern"
	out.Type = direct.Enum_FromProto(mapCtx, in.GetType())
	out.ExactMatch = direct.LazyPtr(in.GetExactMatch())
	// MISSING: GeneratedURIPattern
	// MISSING: RootDomainURI
	return out
}
func DiscoveryEngineDataStoreTargetSiteSpec_ToProto(mapCtx *direct.MapContext, in *krm.DiscoveryEngineDataStoreTargetSiteSpec) *pb.TargetSite {
	if in == nil {
		return nil
	}
	out := &pb.TargetSite{}
	// MISSING: Name
	// MISSING: ProvidedURIPattern
	// (near miss): "ProvidedURIPattern" vs "ProvidedUriPattern"
	out.Type = direct.Enum_ToProto[pb.TargetSite_Type](mapCtx, in.Type)
	out.ExactMatch = direct.ValueOf(in.ExactMatch)
	// MISSING: GeneratedURIPattern
	// MISSING: RootDomainURI
	return out
}
func DiscoveryEngineEngineObservedState_FromProto(mapCtx *direct.MapContext, in *pb.Engine) *krm.DiscoveryEngineEngineObservedState {
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
func DiscoveryEngineEngineObservedState_ToProto(mapCtx *direct.MapContext, in *krm.DiscoveryEngineEngineObservedState) *pb.Engine {
	if in == nil {
		return nil
	}
	out := &pb.Engine{}
	// MISSING: ChatEngineMetadata
	// MISSING: Name
	// MISSING: CreateTime
	// MISSING: UpdateTime
	// MISSING: DataStoreIds
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
func Engine_ChatEngineConfig_FromProto(mapCtx *direct.MapContext, in *pb.Engine_ChatEngineConfig) *krm.Engine_ChatEngineConfig {
	if in == nil {
		return nil
	}
	out := &krm.Engine_ChatEngineConfig{}
	out.AgentCreationConfig = Engine_ChatEngineConfig_AgentCreationConfig_FromProto(mapCtx, in.GetAgentCreationConfig())
	out.DialogflowAgentToLink = direct.LazyPtr(in.GetDialogflowAgentToLink())
	return out
}
func Engine_ChatEngineConfig_ToProto(mapCtx *direct.MapContext, in *krm.Engine_ChatEngineConfig) *pb.Engine_ChatEngineConfig {
	if in == nil {
		return nil
	}
	out := &pb.Engine_ChatEngineConfig{}
	out.AgentCreationConfig = Engine_ChatEngineConfig_AgentCreationConfig_ToProto(mapCtx, in.AgentCreationConfig)
	out.DialogflowAgentToLink = direct.ValueOf(in.DialogflowAgentToLink)
	return out
}
func Engine_ChatEngineConfig_AgentCreationConfig_FromProto(mapCtx *direct.MapContext, in *pb.Engine_ChatEngineConfig_AgentCreationConfig) *krm.Engine_ChatEngineConfig_AgentCreationConfig {
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
func Engine_ChatEngineConfig_AgentCreationConfig_ToProto(mapCtx *direct.MapContext, in *krm.Engine_ChatEngineConfig_AgentCreationConfig) *pb.Engine_ChatEngineConfig_AgentCreationConfig {
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
func Engine_CommonConfig_FromProto(mapCtx *direct.MapContext, in *pb.Engine_CommonConfig) *krm.Engine_CommonConfig {
	if in == nil {
		return nil
	}
	out := &krm.Engine_CommonConfig{}
	out.CompanyName = direct.LazyPtr(in.GetCompanyName())
	return out
}
func Engine_CommonConfig_ToProto(mapCtx *direct.MapContext, in *krm.Engine_CommonConfig) *pb.Engine_CommonConfig {
	if in == nil {
		return nil
	}
	out := &pb.Engine_CommonConfig{}
	out.CompanyName = direct.ValueOf(in.CompanyName)
	return out
}
func Engine_SearchEngineConfig_FromProto(mapCtx *direct.MapContext, in *pb.Engine_SearchEngineConfig) *krm.Engine_SearchEngineConfig {
	if in == nil {
		return nil
	}
	out := &krm.Engine_SearchEngineConfig{}
	out.SearchTier = direct.Enum_FromProto(mapCtx, in.GetSearchTier())
	out.SearchAddOns = direct.EnumSlice_FromProto(mapCtx, in.SearchAddOns)
	return out
}
func Engine_SearchEngineConfig_ToProto(mapCtx *direct.MapContext, in *krm.Engine_SearchEngineConfig) *pb.Engine_SearchEngineConfig {
	if in == nil {
		return nil
	}
	out := &pb.Engine_SearchEngineConfig{}
	out.SearchTier = direct.Enum_ToProto[pb.SearchTier](mapCtx, in.SearchTier)
	out.SearchAddOns = direct.EnumSlice_ToProto[pb.SearchAddOn](mapCtx, in.SearchAddOns)
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
func SiteVerificationInfo_FromProto(mapCtx *direct.MapContext, in *pb.SiteVerificationInfo) *krm.SiteVerificationInfo {
	if in == nil {
		return nil
	}
	out := &krm.SiteVerificationInfo{}
	out.SiteVerificationState = direct.Enum_FromProto(mapCtx, in.GetSiteVerificationState())
	out.VerifyTime = direct.StringTimestamp_FromProto(mapCtx, in.GetVerifyTime())
	return out
}
func SiteVerificationInfo_ToProto(mapCtx *direct.MapContext, in *krm.SiteVerificationInfo) *pb.SiteVerificationInfo {
	if in == nil {
		return nil
	}
	out := &pb.SiteVerificationInfo{}
	out.SiteVerificationState = direct.Enum_ToProto[pb.SiteVerificationInfo_SiteVerificationState](mapCtx, in.SiteVerificationState)
	out.VerifyTime = direct.StringTimestamp_ToProto(mapCtx, in.VerifyTime)
	return out
}
func TargetSite_FailureReason_FromProto(mapCtx *direct.MapContext, in *pb.TargetSite_FailureReason) *krm.TargetSite_FailureReason {
	if in == nil {
		return nil
	}
	out := &krm.TargetSite_FailureReason{}
	out.QuotaFailure = TargetSite_FailureReason_QuotaFailure_FromProto(mapCtx, in.GetQuotaFailure())
	return out
}
func TargetSite_FailureReason_ToProto(mapCtx *direct.MapContext, in *krm.TargetSite_FailureReason) *pb.TargetSite_FailureReason {
	if in == nil {
		return nil
	}
	out := &pb.TargetSite_FailureReason{}
	if oneof := TargetSite_FailureReason_QuotaFailure_ToProto(mapCtx, in.QuotaFailure); oneof != nil {
		out.Failure = &pb.TargetSite_FailureReason_QuotaFailure_{QuotaFailure: oneof}
	}
	return out
}
func TargetSite_FailureReason_QuotaFailure_FromProto(mapCtx *direct.MapContext, in *pb.TargetSite_FailureReason_QuotaFailure) *krm.TargetSite_FailureReason_QuotaFailure {
	if in == nil {
		return nil
	}
	out := &krm.TargetSite_FailureReason_QuotaFailure{}
	out.TotalRequiredQuota = direct.LazyPtr(in.GetTotalRequiredQuota())
	return out
}
func TargetSite_FailureReason_QuotaFailure_ToProto(mapCtx *direct.MapContext, in *krm.TargetSite_FailureReason_QuotaFailure) *pb.TargetSite_FailureReason_QuotaFailure {
	if in == nil {
		return nil
	}
	out := &pb.TargetSite_FailureReason_QuotaFailure{}
	out.TotalRequiredQuota = direct.ValueOf(in.TotalRequiredQuota)
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
