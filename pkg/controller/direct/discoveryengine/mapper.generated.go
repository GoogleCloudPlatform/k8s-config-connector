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
	krm "github.com/GoogleCloudPlatform/k8s-config-connector/apis/discoveryengine/v1alpha1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct"
)

func DiscoveryEngineEngineObservedState_FromProto(mapCtx *direct.MapContext, in *pb.Engine) *krm.DiscoveryEngineEngineObservedState {
	if in == nil {
		return nil
	}
	out := &krm.DiscoveryEngineEngineObservedState{}
	// MISSING: ChatEngineMetadata
	// MISSING: Name
	// MISSING: CreateTime
	// MISSING: UpdateTime
	// MISSING: DisableAnalytics
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
	// MISSING: DisableAnalytics
	return out
}
func DiscoveryEngineEngineSpec_FromProto(mapCtx *direct.MapContext, in *pb.Engine) *krm.DiscoveryEngineEngineSpec {
	if in == nil {
		return nil
	}
	out := &krm.DiscoveryEngineEngineSpec{}
	out.ChatEngineConfig = Engine_ChatEngineConfig_FromProto(mapCtx, in.GetChatEngineConfig())
	out.SearchEngineConfig = Engine_SearchEngineConfig_FromProto(mapCtx, in.GetSearchEngineConfig())
	// MISSING: ChatEngineMetadata
	// MISSING: Name
	out.DisplayName = direct.LazyPtr(in.GetDisplayName())
	// MISSING: CreateTime
	// MISSING: UpdateTime
	out.DataStoreIds = in.DataStoreIds
	out.SolutionType = direct.Enum_FromProto(mapCtx, in.GetSolutionType())
	out.IndustryVertical = direct.Enum_FromProto(mapCtx, in.GetIndustryVertical())
	out.CommonConfig = Engine_CommonConfig_FromProto(mapCtx, in.GetCommonConfig())
	// MISSING: DisableAnalytics
	return out
}
func DiscoveryEngineEngineSpec_ToProto(mapCtx *direct.MapContext, in *krm.DiscoveryEngineEngineSpec) *pb.Engine {
	if in == nil {
		return nil
	}
	out := &pb.Engine{}
	if oneof := Engine_ChatEngineConfig_ToProto(mapCtx, in.ChatEngineConfig); oneof != nil {
		out.EngineConfig = &pb.Engine_ChatEngineConfig_{ChatEngineConfig: oneof}
	}
	if oneof := Engine_SearchEngineConfig_ToProto(mapCtx, in.SearchEngineConfig); oneof != nil {
		out.EngineConfig = &pb.Engine_SearchEngineConfig_{SearchEngineConfig: oneof}
	}
	// MISSING: ChatEngineMetadata
	// MISSING: Name
	out.DisplayName = direct.ValueOf(in.DisplayName)
	// MISSING: CreateTime
	// MISSING: UpdateTime
	out.DataStoreIds = in.DataStoreIds
	out.SolutionType = direct.Enum_ToProto[pb.SolutionType](mapCtx, in.SolutionType)
	out.IndustryVertical = direct.Enum_ToProto[pb.IndustryVertical](mapCtx, in.IndustryVertical)
	out.CommonConfig = Engine_CommonConfig_ToProto(mapCtx, in.CommonConfig)
	// MISSING: DisableAnalytics
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
