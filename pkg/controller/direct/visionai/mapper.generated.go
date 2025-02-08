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

package visionai

import (
	refs "github.com/GoogleCloudPlatform/k8s-config-connector/apis/refs/v1beta1"
	pb "cloud.google.com/go/visionai/apiv1/visionaipb"
	krm "github.com/GoogleCloudPlatform/k8s-config-connector/apis/visionai/v1alpha1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct"
)
func DataSchema_FromProto(mapCtx *direct.MapContext, in *pb.DataSchema) *krm.DataSchema {
	if in == nil {
		return nil
	}
	out := &krm.DataSchema{}
	out.Name = direct.LazyPtr(in.GetName())
	out.Key = direct.LazyPtr(in.GetKey())
	out.SchemaDetails = DataSchemaDetails_FromProto(mapCtx, in.GetSchemaDetails())
	return out
}
func DataSchema_ToProto(mapCtx *direct.MapContext, in *krm.DataSchema) *pb.DataSchema {
	if in == nil {
		return nil
	}
	out := &pb.DataSchema{}
	out.Name = direct.ValueOf(in.Name)
	out.Key = direct.ValueOf(in.Key)
	out.SchemaDetails = DataSchemaDetails_ToProto(mapCtx, in.SchemaDetails)
	return out
}
func DataSchemaDetails_FromProto(mapCtx *direct.MapContext, in *pb.DataSchemaDetails) *krm.DataSchemaDetails {
	if in == nil {
		return nil
	}
	out := &krm.DataSchemaDetails{}
	out.Type = direct.Enum_FromProto(mapCtx, in.GetType())
	out.ProtoAnyConfig = DataSchemaDetails_ProtoAnyConfig_FromProto(mapCtx, in.GetProtoAnyConfig())
	out.ListConfig = DataSchemaDetails_ListConfig_FromProto(mapCtx, in.GetListConfig())
	out.CustomizedStructConfig = DataSchemaDetails_CustomizedStructConfig_FromProto(mapCtx, in.GetCustomizedStructConfig())
	out.Granularity = direct.Enum_FromProto(mapCtx, in.GetGranularity())
	out.SearchStrategy = DataSchemaDetails_SearchStrategy_FromProto(mapCtx, in.GetSearchStrategy())
	return out
}
func DataSchemaDetails_ToProto(mapCtx *direct.MapContext, in *krm.DataSchemaDetails) *pb.DataSchemaDetails {
	if in == nil {
		return nil
	}
	out := &pb.DataSchemaDetails{}
	if oneof := DataSchemaDetails_Type_ToProto(mapCtx, in.Type); oneof != nil {
		out.Type = oneof
	}
	out.ProtoAnyConfig = DataSchemaDetails_ProtoAnyConfig_ToProto(mapCtx, in.ProtoAnyConfig)
	out.ListConfig = DataSchemaDetails_ListConfig_ToProto(mapCtx, in.ListConfig)
	out.CustomizedStructConfig = DataSchemaDetails_CustomizedStructConfig_ToProto(mapCtx, in.CustomizedStructConfig)
	if oneof := DataSchemaDetails_Granularity_ToProto(mapCtx, in.Granularity); oneof != nil {
		out.Granularity = oneof
	}
	out.SearchStrategy = DataSchemaDetails_SearchStrategy_ToProto(mapCtx, in.SearchStrategy)
	return out
}
func DataSchemaDetails_CustomizedStructConfig_FromProto(mapCtx *direct.MapContext, in *pb.DataSchemaDetails_CustomizedStructConfig) *krm.DataSchemaDetails_CustomizedStructConfig {
	if in == nil {
		return nil
	}
	out := &krm.DataSchemaDetails_CustomizedStructConfig{}
	// MISSING: FieldSchemas
	return out
}
func DataSchemaDetails_CustomizedStructConfig_ToProto(mapCtx *direct.MapContext, in *krm.DataSchemaDetails_CustomizedStructConfig) *pb.DataSchemaDetails_CustomizedStructConfig {
	if in == nil {
		return nil
	}
	out := &pb.DataSchemaDetails_CustomizedStructConfig{}
	// MISSING: FieldSchemas
	return out
}
func DataSchemaDetails_ListConfig_FromProto(mapCtx *direct.MapContext, in *pb.DataSchemaDetails_ListConfig) *krm.DataSchemaDetails_ListConfig {
	if in == nil {
		return nil
	}
	out := &krm.DataSchemaDetails_ListConfig{}
	out.ValueSchema = DataSchemaDetails_FromProto(mapCtx, in.GetValueSchema())
	return out
}
func DataSchemaDetails_ListConfig_ToProto(mapCtx *direct.MapContext, in *krm.DataSchemaDetails_ListConfig) *pb.DataSchemaDetails_ListConfig {
	if in == nil {
		return nil
	}
	out := &pb.DataSchemaDetails_ListConfig{}
	out.ValueSchema = DataSchemaDetails_ToProto(mapCtx, in.ValueSchema)
	return out
}
func DataSchemaDetails_ProtoAnyConfig_FromProto(mapCtx *direct.MapContext, in *pb.DataSchemaDetails_ProtoAnyConfig) *krm.DataSchemaDetails_ProtoAnyConfig {
	if in == nil {
		return nil
	}
	out := &krm.DataSchemaDetails_ProtoAnyConfig{}
	out.TypeURI = direct.LazyPtr(in.GetTypeUri())
	return out
}
func DataSchemaDetails_ProtoAnyConfig_ToProto(mapCtx *direct.MapContext, in *krm.DataSchemaDetails_ProtoAnyConfig) *pb.DataSchemaDetails_ProtoAnyConfig {
	if in == nil {
		return nil
	}
	out := &pb.DataSchemaDetails_ProtoAnyConfig{}
	out.TypeUri = direct.ValueOf(in.TypeURI)
	return out
}
func DataSchemaDetails_SearchStrategy_FromProto(mapCtx *direct.MapContext, in *pb.DataSchemaDetails_SearchStrategy) *krm.DataSchemaDetails_SearchStrategy {
	if in == nil {
		return nil
	}
	out := &krm.DataSchemaDetails_SearchStrategy{}
	out.SearchStrategyType = direct.Enum_FromProto(mapCtx, in.GetSearchStrategyType())
	out.ConfidenceScoreIndexConfig = DataSchemaDetails_SearchStrategy_ConfidenceScoreIndexConfig_FromProto(mapCtx, in.GetConfidenceScoreIndexConfig())
	return out
}
func DataSchemaDetails_SearchStrategy_ToProto(mapCtx *direct.MapContext, in *krm.DataSchemaDetails_SearchStrategy) *pb.DataSchemaDetails_SearchStrategy {
	if in == nil {
		return nil
	}
	out := &pb.DataSchemaDetails_SearchStrategy{}
	if oneof := DataSchemaDetails_SearchStrategy_SearchStrategyType_ToProto(mapCtx, in.SearchStrategyType); oneof != nil {
		out.SearchStrategyType = oneof
	}
	out.ConfidenceScoreIndexConfig = DataSchemaDetails_SearchStrategy_ConfidenceScoreIndexConfig_ToProto(mapCtx, in.ConfidenceScoreIndexConfig)
	return out
}
func DataSchemaDetails_SearchStrategy_ConfidenceScoreIndexConfig_FromProto(mapCtx *direct.MapContext, in *pb.DataSchemaDetails_SearchStrategy_ConfidenceScoreIndexConfig) *krm.DataSchemaDetails_SearchStrategy_ConfidenceScoreIndexConfig {
	if in == nil {
		return nil
	}
	out := &krm.DataSchemaDetails_SearchStrategy_ConfidenceScoreIndexConfig{}
	out.FieldPath = direct.LazyPtr(in.GetFieldPath())
	out.Threshold = direct.LazyPtr(in.GetThreshold())
	return out
}
func DataSchemaDetails_SearchStrategy_ConfidenceScoreIndexConfig_ToProto(mapCtx *direct.MapContext, in *krm.DataSchemaDetails_SearchStrategy_ConfidenceScoreIndexConfig) *pb.DataSchemaDetails_SearchStrategy_ConfidenceScoreIndexConfig {
	if in == nil {
		return nil
	}
	out := &pb.DataSchemaDetails_SearchStrategy_ConfidenceScoreIndexConfig{}
	out.FieldPath = direct.ValueOf(in.FieldPath)
	out.Threshold = direct.ValueOf(in.Threshold)
	return out
}
func VisionaiDataSchemaObservedState_FromProto(mapCtx *direct.MapContext, in *pb.DataSchema) *krm.VisionaiDataSchemaObservedState {
	if in == nil {
		return nil
	}
	out := &krm.VisionaiDataSchemaObservedState{}
	// MISSING: Name
	// MISSING: Key
	// MISSING: SchemaDetails
	return out
}
func VisionaiDataSchemaObservedState_ToProto(mapCtx *direct.MapContext, in *krm.VisionaiDataSchemaObservedState) *pb.DataSchema {
	if in == nil {
		return nil
	}
	out := &pb.DataSchema{}
	// MISSING: Name
	// MISSING: Key
	// MISSING: SchemaDetails
	return out
}
func VisionaiDataSchemaSpec_FromProto(mapCtx *direct.MapContext, in *pb.DataSchema) *krm.VisionaiDataSchemaSpec {
	if in == nil {
		return nil
	}
	out := &krm.VisionaiDataSchemaSpec{}
	// MISSING: Name
	// MISSING: Key
	// MISSING: SchemaDetails
	return out
}
func VisionaiDataSchemaSpec_ToProto(mapCtx *direct.MapContext, in *krm.VisionaiDataSchemaSpec) *pb.DataSchema {
	if in == nil {
		return nil
	}
	out := &pb.DataSchema{}
	// MISSING: Name
	// MISSING: Key
	// MISSING: SchemaDetails
	return out
}
