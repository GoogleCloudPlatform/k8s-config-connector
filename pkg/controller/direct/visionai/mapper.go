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

package visionai

import (
	pb "cloud.google.com/go/visionai/apiv1/visionaipb"
	krm "github.com/GoogleCloudPlatform/k8s-config-connector/apis/visionai/v1alpha1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct"
)

func AnnotationValue_ToProto(mapCtx *direct.MapContext, in *krm.AnnotationValue) *pb.AnnotationValue {
	if in == nil {
		return nil
	}
	out := &pb.AnnotationValue{}
	if oneof := AnnotationValue_IntValue_ToProto(mapCtx, in.IntValue); oneof != nil {
		out.Value = oneof
	}
	if oneof := AnnotationValue_FloatValue_ToProto(mapCtx, in.FloatValue); oneof != nil {
		out.Value = oneof
	}
	if oneof := AnnotationValue_StrValue_ToProto(mapCtx, in.StrValue); oneof != nil {
		out.Value = oneof
	}
	if oneof := AnnotationValue_DatetimeValue_ToProto(mapCtx, in.DatetimeValue); oneof != nil {
		out.Value = oneof
	}
	if oneof := GeoCoordinate_ToProto(mapCtx, in.GeoCoordinate); oneof != nil {
		out.Value = &pb.AnnotationValue_GeoCoordinate{GeoCoordinate: oneof}
	}
	if oneof := Any_ToProto(mapCtx, in.ProtoAnyValue); oneof != nil {
		out.Value = &pb.AnnotationValue_ProtoAnyValue{ProtoAnyValue: oneof}
	}
	if oneof := AnnotationValue_BoolValue_ToProto(mapCtx, in.BoolValue); oneof != nil {
		out.Value = oneof
	}
	if oneof := direct.Struct_ToProto(mapCtx, &in.CustomizedStructDataValue); oneof != nil {
		out.Value = &pb.AnnotationValue_CustomizedStructDataValue{CustomizedStructDataValue: oneof}
	}
	if oneof := AnnotationList_ToProto(mapCtx, in.ListValue); oneof != nil {
		out.Value = &pb.AnnotationValue_ListValue{ListValue: oneof}
	}
	if oneof := AnnotationCustomizedStruct_ToProto(mapCtx, in.CustomizedStructValue); oneof != nil {
		out.Value = &pb.AnnotationValue_CustomizedStructValue{CustomizedStructValue: oneof}
	}
	return out
}

func AttributeValue_S_ToProto(mapCtx *direct.MapContext, in []byte) *pb.AttributeValue_S {
	if in == nil {
		return nil
	}
	return &pb.AttributeValue_S{S: in}
}

func DataSchemaDetails_Type_ToProto(mapCtx *direct.MapContext, in *string) *pb.DataSchemaDetails_DataType {
	if in == nil {
		return nil
	}
	v := direct.Enum_ToProto[pb.DataSchemaDetails_DataType](mapCtx, in)
	return &v
}

func DataSchemaDetails_Granularity_ToProto(mapCtx *direct.MapContext, in *string) *pb.DataSchemaDetails_Granularity {
	if in == nil {
		return nil
	}
	v := direct.Enum_ToProto[pb.DataSchemaDetails_Granularity](mapCtx, in)
	return &v
}

func DataSchemaDetails_SearchStrategy_SearchStrategyType_ToProto(mapCtx *direct.MapContext, in *string) *pb.DataSchemaDetails_SearchStrategy_SearchStrategyType {
	if in == nil {
		return nil
	}
	v := direct.Enum_ToProto[pb.DataSchemaDetails_SearchStrategy_SearchStrategyType](mapCtx, in)
	return &v
}
