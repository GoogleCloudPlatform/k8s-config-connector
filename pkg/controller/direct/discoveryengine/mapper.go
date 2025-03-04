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
	"fmt"
	"strings"

	pb "cloud.google.com/go/discoveryengine/apiv1/discoveryenginepb"
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
		tokens := strings.Split(in.Name, "/")
		if len(tokens) != 8 || tokens[0] != "projects" || tokens[2] != "locations" || tokens[4] != "collections" || tokens[6] != "engines" {
			mapCtx.Errorf("unexpected name %q in DiscoveryEngineEngineSpec_FromProto", in.Name)
			continue
		}
		projectID := tokens[1]
		location := tokens[3]
		collection := tokens[5]
		external := fmt.Sprintf("projects/%s/locations/%s/collections/%s/dataStores/%s", projectID, location, collection, dataStoreID)
		out.DataStoreRefs = append(out.DataStoreRefs, krm.DiscoveryEngineDataStoreRef{External: external})
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
		id, err := krm.ParseDiscoveryEngineDataStoreExternal(dataStoreRef.External)
		if err != nil {
			mapCtx.Errorf("parsing ref: %w", err)
			continue
		}
		out.DataStoreIds = append(out.DataStoreIds, id.DataStore)
	}

	return out
}
