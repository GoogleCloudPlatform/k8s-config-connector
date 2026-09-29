// Copyright 2026 Google LLC
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

package aiplatform

import (
	pb "cloud.google.com/go/aiplatform/apiv1/aiplatformpb"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct"
)

func ComputationBasedMetricSpec_Type_ToProto(mapCtx *direct.MapContext, in *string) *pb.ComputationBasedMetricSpec_ComputationBasedMetricType {
	if in == nil {
		return nil
	}
	val := direct.Enum_ToProto[pb.ComputationBasedMetricSpec_ComputationBasedMetricType](mapCtx, in)
	return &val
}

func GenerationConfig_MediaResolution_ToProto(mapCtx *direct.MapContext, in *string) *pb.GenerationConfig_MediaResolution {
	if in == nil {
		return nil
	}
	val := direct.Enum_ToProto[pb.GenerationConfig_MediaResolution](mapCtx, in)
	return &val
}

func GenerationConfig_RoutingConfig_AutoRoutingMode_ModelRoutingPreference_ToProto(mapCtx *direct.MapContext, in *string) *pb.GenerationConfig_RoutingConfig_AutoRoutingMode_ModelRoutingPreference {
	if in == nil {
		return nil
	}
	val := direct.Enum_ToProto[pb.GenerationConfig_RoutingConfig_AutoRoutingMode_ModelRoutingPreference](mapCtx, in)
	return &val
}

func GenerationConfig_ThinkingConfig_ThinkingLevel_ToProto(mapCtx *direct.MapContext, in *string) *pb.GenerationConfig_ThinkingConfig_ThinkingLevel {
	if in == nil {
		return nil
	}
	val := direct.Enum_ToProto[pb.GenerationConfig_ThinkingConfig_ThinkingLevel](mapCtx, in)
	return &val
}

func ImageConfig_PersonGeneration_ToProto(mapCtx *direct.MapContext, in *string) *pb.ImageConfig_PersonGeneration {
	if in == nil {
		return nil
	}
	val := direct.Enum_ToProto[pb.ImageConfig_PersonGeneration](mapCtx, in)
	return &val
}
