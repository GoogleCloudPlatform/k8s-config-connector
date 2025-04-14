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
	pb "cloud.google.com/go/discoveryengine/apiv1alpha/discoveryenginepb"
	krm "github.com/GoogleCloudPlatform/k8s-config-connector/apis/discoveryengine/v1alpha1"
	krmv1alpha1 "github.com/GoogleCloudPlatform/k8s-config-connector/apis/discoveryengine/v1alpha1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct"
	"google.golang.org/protobuf/types/known/structpb"
)

// Override but should be unreachable.
// Would not be needed if we did a reachability analysis in our proto generation.
func StructSchema_FromProto(mapCtx *direct.MapContext, in *structpb.Struct) map[string]string {
	mapCtx.NotImplemented()
	return nil
}

// Override but should be unreachable
// Would not be needed if we did a reachability analysis in our proto generation.
func StructSchema_ToProto(mapCtx *direct.MapContext, in map[string]string) *structpb.Struct {
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
	out.ChatEngineConfig = Engine_ChatEngineConfig_FromProto(mapCtx, in.GetChatEngineConfig())
	out.SearchEngineConfig = Engine_SearchEngineConfig_FromProto(mapCtx, in.GetSearchEngineConfig())
	// MISSING: ChatEngineMetadata
	out.DisplayName = direct.LazyPtr(in.GetDisplayName())
	out.SolutionType = direct.Enum_FromProto(mapCtx, in.GetSolutionType())
	out.IndustryVertical = direct.Enum_FromProto(mapCtx, in.GetIndustryVertical())
	out.CommonConfig = Engine_CommonConfig_FromProto(mapCtx, in.GetCommonConfig())
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
	if oneof := Engine_ChatEngineConfig_ToProto(mapCtx, in.ChatEngineConfig); oneof != nil {
		out.EngineConfig = &pb.Engine_ChatEngineConfig_{ChatEngineConfig: oneof}
	}
	if oneof := Engine_SearchEngineConfig_ToProto(mapCtx, in.SearchEngineConfig); oneof != nil {
		out.EngineConfig = &pb.Engine_SearchEngineConfig_{SearchEngineConfig: oneof}
	}
	// MISSING: ChatEngineMetadata
	out.DisplayName = direct.ValueOf(in.DisplayName)
	out.SolutionType = direct.Enum_ToProto[pb.SolutionType](mapCtx, in.SolutionType)
	out.IndustryVertical = direct.Enum_ToProto[pb.IndustryVertical](mapCtx, in.IndustryVertical)
	out.CommonConfig = Engine_CommonConfig_ToProto(mapCtx, in.CommonConfig)

	for _, dataStoreRef := range in.DataStoreRefs {
		out.DataStoreIds = append(out.DataStoreIds, dataStoreRef.External)
	}

	return out
}

func DiscoveryEngineDataStoreSpec_FromProto(mapCtx *direct.MapContext, in *pb.DataStore) *krmv1alpha1.DiscoveryEngineDataStoreSpec {
	if in == nil {
		return nil
	}
	out := &krmv1alpha1.DiscoveryEngineDataStoreSpec{}
	// MISSING: Name
	out.DisplayName = direct.LazyPtr(in.GetDisplayName())
	out.IndustryVertical = direct.Enum_FromProto(mapCtx, in.GetIndustryVertical())
	out.SolutionTypes = direct.EnumSlice_FromProto(mapCtx, in.SolutionTypes)
	// MISSING: DefaultSchemaID
	out.ContentConfig = direct.Enum_FromProto(mapCtx, in.GetContentConfig())
	// MISSING: CreateTime
	out.LanguageInfo = LanguageInfo_FromProto(mapCtx, in.GetLanguageInfo())
	// MISSING: IdpConfig
	out.AclEnabled = direct.LazyPtr(in.GetAclEnabled())
	out.WorkspaceConfig = WorkspaceConfig_FromProto(mapCtx, in.GetWorkspaceConfig())
	// MISSING: DocumentProcessingConfig
	out.StartingSchema = Schema_FromProto(mapCtx, in.GetStartingSchema())
	return out
}
func DiscoveryEngineDataStoreSpec_ToProto(mapCtx *direct.MapContext, in *krmv1alpha1.DiscoveryEngineDataStoreSpec) *pb.DataStore {
	if in == nil {
		return nil
	}
	out := &pb.DataStore{}
	// MISSING: Name
	out.DisplayName = direct.ValueOf(in.DisplayName)
	out.IndustryVertical = direct.Enum_ToProto[pb.IndustryVertical](mapCtx, in.IndustryVertical)
	out.SolutionTypes = direct.EnumSlice_ToProto[pb.SolutionType](mapCtx, in.SolutionTypes)
	// MISSING: DefaultSchemaID
	out.ContentConfig = direct.Enum_ToProto[pb.DataStore_ContentConfig](mapCtx, in.ContentConfig)
	// MISSING: CreateTime
	out.LanguageInfo = LanguageInfo_ToProto(mapCtx, in.LanguageInfo)
	// MISSING: IdpConfig
	out.AclEnabled = direct.ValueOf(in.AclEnabled)
	out.WorkspaceConfig = WorkspaceConfig_ToProto(mapCtx, in.WorkspaceConfig)
	// MISSING: DocumentProcessingConfig
	out.StartingSchema = Schema_ToProto(mapCtx, in.StartingSchema)
	return out
}

func DataStoreTargetSiteSpec_FromProto(mapCtx *direct.MapContext, in *pb.TargetSite) *krm.DiscoveryEngineDataStoreTargetSiteSpec {
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
func DataStoreTargetSiteSpec_ToProto(mapCtx *direct.MapContext, in *krm.DiscoveryEngineDataStoreTargetSiteSpec) *pb.TargetSite {
	if in == nil {
		return nil
	}
	out := &pb.TargetSite{}
	// MISSING: Name
	out.ProvidedUriPattern = direct.ValueOf(in.ProvidedURIPattern)
	out.Type = direct.Enum_ToProto[pb.TargetSite_Type](mapCtx, in.Type)
	out.ExactMatch = direct.ValueOf(in.ExactMatch)
	return out
}
