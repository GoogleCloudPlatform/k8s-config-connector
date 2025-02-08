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
	refs "github.com/GoogleCloudPlatform/k8s-config-connector/apis/refs/v1beta1"
	pb "cloud.google.com/go/discoveryengine/apiv1/discoveryenginepb"
	krm "github.com/GoogleCloudPlatform/k8s-config-connector/apis/discoveryengine/v1alpha1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct"
)
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
