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
	krmdiscoveryenginev1alpha1 "github.com/GoogleCloudPlatform/k8s-config-connector/apis/discoveryengine/v1alpha1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct"
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
	out.ChatEngineConfig = Engine_ChatEngineConfig_FromProto(mapCtx, in.GetChatEngineConfig())
	out.SearchEngineConfig = Engine_SearchEngineConfig_FromProto(mapCtx, in.GetSearchEngineConfig())
	// MISSING: ChatEngineMetadata
	out.DisplayName = direct.LazyPtr(in.GetDisplayName())
	out.SolutionType = direct.Enum_FromProto(mapCtx, in.GetSolutionType())
	out.IndustryVertical = direct.Enum_FromProto(mapCtx, in.GetIndustryVertical())
	out.CommonConfig = Engine_CommonConfig_FromProto(mapCtx, in.GetCommonConfig())
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
	out.DisableAnalytics = direct.ValueOf(in.DisableAnalytics)

	for _, dataStoreRef := range in.DataStoreRefs {
		out.DataStoreIds = append(out.DataStoreIds, dataStoreRef.External)
	}

	return out
}

func SearchResponse_Summary_FromProto(mapCtx *direct.MapContext, in *pb.SearchResponse_Summary) *krm.SearchResponse_Summary {
	if in == nil {
		return nil
	}
	out := &krm.SearchResponse_Summary{}
	out.SummaryText = direct.LazyPtr(in.GetSummaryText())
	out.SummarySkippedReasons = direct.EnumSlice_FromProto(mapCtx, in.SummarySkippedReasons)
	return out
}

func SearchResponse_Summary_ToProto(mapCtx *direct.MapContext, in *krm.SearchResponse_Summary) *pb.SearchResponse_Summary {
	if in == nil {
		return nil
	}
	out := &pb.SearchResponse_Summary{}
	out.SummaryText = direct.ValueOf(in.SummaryText)
	out.SummarySkippedReasons = direct.EnumSlice_ToProto[pb.SearchResponse_Summary_SummarySkippedReason](mapCtx, in.SummarySkippedReasons)
	return out
}

// --- Unversioned delegating forwarders / manual overrides version wrappers ---

func DiscoveryEngineControlObservedState_FromProto(mapCtx *direct.MapContext, in *pb.Control) *krmdiscoveryenginev1alpha1.DiscoveryEngineControlObservedState {
	return DiscoveryEngineControlObservedState_v1alpha1_FromProto(mapCtx, in)
}

func DiscoveryEngineDataStoreSpec_FromProto(mapCtx *direct.MapContext, in *pb.DataStore) *krmdiscoveryenginev1alpha1.DiscoveryEngineDataStoreSpec {
	return DiscoveryEngineDataStoreSpec_v1alpha1_FromProto(mapCtx, in)
}

func DiscoveryEngineDataStoreTargetSiteSpec_ToProto(mapCtx *direct.MapContext, in *krmdiscoveryenginev1alpha1.DiscoveryEngineDataStoreTargetSiteSpec) *pb.TargetSite {
	return DiscoveryEngineDataStoreTargetSiteSpec_v1alpha1_ToProto(mapCtx, in)
}

func DiscoveryEngineControlSpec_FromProto(mapCtx *direct.MapContext, in *pb.Control) *krmdiscoveryenginev1alpha1.DiscoveryEngineControlSpec {
	return DiscoveryEngineControlSpec_v1alpha1_FromProto(mapCtx, in)
}

func DiscoveryEngineDataStoreTargetSiteSpec_FromProto(mapCtx *direct.MapContext, in *pb.TargetSite) *krmdiscoveryenginev1alpha1.DiscoveryEngineDataStoreTargetSiteSpec {
	return DiscoveryEngineDataStoreTargetSiteSpec_v1alpha1_FromProto(mapCtx, in)
}

func DiscoveryEngineDataStoreSpec_ToProto(mapCtx *direct.MapContext, in *krmdiscoveryenginev1alpha1.DiscoveryEngineDataStoreSpec) *pb.DataStore {
	return DiscoveryEngineDataStoreSpec_v1alpha1_ToProto(mapCtx, in)
}

func DiscoveryEngineDataStoreObservedState_FromProto(mapCtx *direct.MapContext, in *pb.DataStore) *krmdiscoveryenginev1alpha1.DiscoveryEngineDataStoreObservedState {
	return DiscoveryEngineDataStoreObservedState_v1alpha1_FromProto(mapCtx, in)
}

func DiscoveryEngineControlSpec_ToProto(mapCtx *direct.MapContext, in *krmdiscoveryenginev1alpha1.DiscoveryEngineControlSpec) *pb.Control {
	return DiscoveryEngineControlSpec_v1alpha1_ToProto(mapCtx, in)
}

func DiscoveryEngineControlObservedState_ToProto(mapCtx *direct.MapContext, in *krmdiscoveryenginev1alpha1.DiscoveryEngineControlObservedState) *pb.Control {
	return DiscoveryEngineControlObservedState_v1alpha1_ToProto(mapCtx, in)
}

// --- Unversioned delegating forwarders / manual overrides version wrappers ---

func DiscoveryEngineDataStoreTargetSiteObservedState_ToProto(mapCtx *direct.MapContext, in *krmdiscoveryenginev1alpha1.DiscoveryEngineDataStoreTargetSiteObservedState) *pb.TargetSite {
	return DiscoveryEngineDataStoreTargetSiteObservedState_v1alpha1_ToProto(mapCtx, in)
}

func DiscoveryEngineDataStoreTargetSiteObservedState_FromProto(mapCtx *direct.MapContext, in *pb.TargetSite) *krmdiscoveryenginev1alpha1.DiscoveryEngineDataStoreTargetSiteObservedState {
	return DiscoveryEngineDataStoreTargetSiteObservedState_v1alpha1_FromProto(mapCtx, in)
}

func DiscoveryEngineSessionSpec_FromProto(mapCtx *direct.MapContext, in *pb.Session) *krmdiscoveryenginev1alpha1.DiscoveryEngineSessionSpec {
	return DiscoveryEngineSessionSpec_v1alpha1_FromProto(mapCtx, in)
}

func DiscoveryEngineSessionObservedState_FromProto(mapCtx *direct.MapContext, in *pb.Session) *krmdiscoveryenginev1alpha1.DiscoveryEngineSessionObservedState {
	return DiscoveryEngineSessionObservedState_v1alpha1_FromProto(mapCtx, in)
}

func DiscoveryEngineEngineObservedState_FromProto(mapCtx *direct.MapContext, in *pb.Engine) *krmdiscoveryenginev1alpha1.DiscoveryEngineEngineObservedState {
	return DiscoveryEngineEngineObservedState_v1alpha1_FromProto(mapCtx, in)
}

func DiscoveryEngineEngineObservedState_ToProto(mapCtx *direct.MapContext, in *krmdiscoveryenginev1alpha1.DiscoveryEngineEngineObservedState) *pb.Engine {
	return DiscoveryEngineEngineObservedState_v1alpha1_ToProto(mapCtx, in)
}

func DiscoveryEngineSessionSpec_ToProto(mapCtx *direct.MapContext, in *krmdiscoveryenginev1alpha1.DiscoveryEngineSessionSpec) *pb.Session {
	return DiscoveryEngineSessionSpec_v1alpha1_ToProto(mapCtx, in)
}

// --- Unversioned delegating forwarders / manual overrides version wrappers ---

func Engine_SearchEngineConfig_FromProto(mapCtx *direct.MapContext, in *pb.Engine_SearchEngineConfig) *krmdiscoveryenginev1alpha1.Engine_SearchEngineConfig {
	return Engine_SearchEngineConfig_v1alpha1_FromProto(mapCtx, in)
}

func DiscoveryEngineSessionObservedState_ToProto(mapCtx *direct.MapContext, in *krmdiscoveryenginev1alpha1.DiscoveryEngineSessionObservedState) *pb.Session {
	return DiscoveryEngineSessionObservedState_v1alpha1_ToProto(mapCtx, in)
}

func Engine_SearchEngineConfig_ToProto(mapCtx *direct.MapContext, in *krmdiscoveryenginev1alpha1.Engine_SearchEngineConfig) *pb.Engine_SearchEngineConfig {
	return Engine_SearchEngineConfig_v1alpha1_ToProto(mapCtx, in)
}

func Engine_ChatEngineConfig_FromProto(mapCtx *direct.MapContext, in *pb.Engine_ChatEngineConfig) *krmdiscoveryenginev1alpha1.Engine_ChatEngineConfig {
	return Engine_ChatEngineConfig_v1alpha1_FromProto(mapCtx, in)
}

func Engine_CommonConfig_FromProto(mapCtx *direct.MapContext, in *pb.Engine_CommonConfig) *krmdiscoveryenginev1alpha1.Engine_CommonConfig {
	return Engine_CommonConfig_v1alpha1_FromProto(mapCtx, in)
}

func DiscoveryEngineDataStoreObservedState_ToProto(mapCtx *direct.MapContext, in *krmdiscoveryenginev1alpha1.DiscoveryEngineDataStoreObservedState) *pb.DataStore {
	return DiscoveryEngineDataStoreObservedState_v1alpha1_ToProto(mapCtx, in)
}

func Engine_CommonConfig_ToProto(mapCtx *direct.MapContext, in *krmdiscoveryenginev1alpha1.Engine_CommonConfig) *pb.Engine_CommonConfig {
	return Engine_CommonConfig_v1alpha1_ToProto(mapCtx, in)
}

func Engine_ChatEngineConfig_ToProto(mapCtx *direct.MapContext, in *krmdiscoveryenginev1alpha1.Engine_ChatEngineConfig) *pb.Engine_ChatEngineConfig {
	return Engine_ChatEngineConfig_v1alpha1_ToProto(mapCtx, in)
}

// --- Unversioned delegating forwarders / manual overrides version wrappers ---

func SearchResponse_Summary_v1alpha1_FromProto(mapCtx *direct.MapContext, in *pb.SearchResponse_Summary) *krm.SearchResponse_Summary {
	return SearchResponse_Summary_FromProto(mapCtx, in)
}

func SearchResponse_Summary_v1alpha1_ToProto(mapCtx *direct.MapContext, in *krm.SearchResponse_Summary) *pb.SearchResponse_Summary {
	return SearchResponse_Summary_ToProto(mapCtx, in)
}
