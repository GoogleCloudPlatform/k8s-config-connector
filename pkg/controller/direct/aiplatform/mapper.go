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
	krm "github.com/GoogleCloudPlatform/k8s-config-connector/apis/aiplatform/v1alpha1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/types/known/wrapperspb"
	apiextensionsv1 "k8s.io/apiextensions-apiserver/pkg/apis/apiextensions/v1"
)

func Int32Value_FromProto(mapCtx *direct.MapContext, in *wrapperspb.Int32Value) *krm.Int32Value {
	if in == nil {
		return nil
	}
	out := in.Value
	return &krm.Int32Value{
		Value: &out,
	}
}

func Int32Value_ToProto(mapCtx *direct.MapContext, in *krm.Int32Value) *wrapperspb.Int32Value {
	if in == nil || in.Value == nil {
		return nil
	}
	return wrapperspb.Int32(*in.Value)
}

func ComputationBasedMetricSpec_Type_ToProto(mapCtx *direct.MapContext, in *string) *pb.ComputationBasedMetricSpec_ComputationBasedMetricType {
	if in == nil {
		return nil
	}
	out := direct.Enum_ToProto[pb.ComputationBasedMetricSpec_ComputationBasedMetricType](mapCtx, in)
	return &out
}

func GenerationConfig_MediaResolution_ToProto(mapCtx *direct.MapContext, in *string) *pb.GenerationConfig_MediaResolution {
	if in == nil {
		return nil
	}
	out := direct.Enum_ToProto[pb.GenerationConfig_MediaResolution](mapCtx, in)
	return &out
}

func GenerationConfig_RoutingConfig_AutoRoutingMode_ModelRoutingPreference_ToProto(mapCtx *direct.MapContext, in *string) *pb.GenerationConfig_RoutingConfig_AutoRoutingMode_ModelRoutingPreference {
	if in == nil {
		return nil
	}
	out := direct.Enum_ToProto[pb.GenerationConfig_RoutingConfig_AutoRoutingMode_ModelRoutingPreference](mapCtx, in)
	return &out
}

func GenerationConfig_ThinkingConfig_ThinkingLevel_ToProto(mapCtx *direct.MapContext, in *string) *pb.GenerationConfig_ThinkingConfig_ThinkingLevel {
	if in == nil {
		return nil
	}
	out := direct.Enum_ToProto[pb.GenerationConfig_ThinkingConfig_ThinkingLevel](mapCtx, in)
	return &out
}

func ImageConfig_PersonGeneration_ToProto(mapCtx *direct.MapContext, in *string) *pb.ImageConfig_PersonGeneration {
	if in == nil {
		return nil
	}
	out := direct.Enum_ToProto[pb.ImageConfig_PersonGeneration](mapCtx, in)
	return &out
}

func Schema_FromProto(mapCtx *direct.MapContext, in *pb.Schema) apiextensionsv1.JSON {
	if in == nil {
		return apiextensionsv1.JSON{}
	}
	b, err := protojson.Marshal(in)
	if err != nil {
		mapCtx.Errorf("error marshalling Schema to JSON: %v", err)
		return apiextensionsv1.JSON{}
	}
	return apiextensionsv1.JSON{Raw: b}
}

func Schema_ToProto(mapCtx *direct.MapContext, in apiextensionsv1.JSON) *pb.Schema {
	if len(in.Raw) == 0 {
		return nil
	}
	out := &pb.Schema{}
	if err := protojson.Unmarshal(in.Raw, out); err != nil {
		mapCtx.Errorf("error unmarshalling JSON to Schema: %v", err)
		return nil
	}
	return out
}
