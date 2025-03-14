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
	pb_v1beta "cloud.google.com/go/discoveryengine/apiv1beta/discoveryenginepb"
	krm "github.com/GoogleCloudPlatform/k8s-config-connector/apis/discoveryengine/v1alpha1"
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

// +function-gen:special-mapper
// proto.field_definition: float content_watched_percentage_threshold = 2;
// krm.field: ContentWatchedPercentageThreshold *float32 `json:"contentWatchedPercentageThreshold,omitempty"`
// proto.field: google.cloud.discoveryengine.v1beta.ServingConfig.MediaConfig.content_watched_percentage_threshold
func ServingConfig_MediaConfig_ContentWatchedPercentageThreshold_ToProto(mapCtx *direct.MapContext, in *float32) *pb_v1beta.ServingConfig_MediaConfig_ContentWatchedPercentageThreshold {
	if in == nil {
		return nil
	}
	return &pb_v1beta.ServingConfig_MediaConfig_ContentWatchedPercentageThreshold{ContentWatchedPercentageThreshold: *in}
}

// +function-gen:special-mapper
// proto.field_definition: float content_watched_seconds_threshold = 5;
// krm.field: ContentWatchedSecondsThreshold *float32 `json:"contentWatchedSecondsThreshold,omitempty"`
// proto.field: google.cloud.discoveryengine.v1beta.ServingConfig.MediaConfig.content_watched_seconds_threshold
func ServingConfig_MediaConfig_ContentWatchedSecondsThreshold_ToProto(mapCtx *direct.MapContext, in *float32) *pb_v1beta.ServingConfig_MediaConfig_ContentWatchedSecondsThreshold {
	if in == nil {
		return nil
	}
	return &pb_v1beta.ServingConfig_MediaConfig_ContentWatchedSecondsThreshold{ContentWatchedSecondsThreshold: *in}
}
