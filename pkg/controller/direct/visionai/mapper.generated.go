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
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct"
	refs "github.com/GoogleCloudPlatform/k8s-config-connector/apis/refs/v1beta1"
	pb "cloud.google.com/go/visionai/apiv1/visionaipb"
	krm "github.com/GoogleCloudPlatform/k8s-config-connector/apis/visionai/v1alpha1"
)
func Annotation_FromProto(mapCtx *direct.MapContext, in *pb.Annotation) *krm.Annotation {
	if in == nil {
		return nil
	}
	out := &krm.Annotation{}
	out.Name = direct.LazyPtr(in.GetName())
	out.UserSpecifiedAnnotation = UserSpecifiedAnnotation_FromProto(mapCtx, in.GetUserSpecifiedAnnotation())
	return out
}
func Annotation_ToProto(mapCtx *direct.MapContext, in *krm.Annotation) *pb.Annotation {
	if in == nil {
		return nil
	}
	out := &pb.Annotation{}
	out.Name = direct.ValueOf(in.Name)
	out.UserSpecifiedAnnotation = UserSpecifiedAnnotation_ToProto(mapCtx, in.UserSpecifiedAnnotation)
	return out
}
func AnnotationCustomizedStruct_FromProto(mapCtx *direct.MapContext, in *pb.AnnotationCustomizedStruct) *krm.AnnotationCustomizedStruct {
	if in == nil {
		return nil
	}
	out := &krm.AnnotationCustomizedStruct{}
	// MISSING: Elements
	return out
}
func AnnotationCustomizedStruct_ToProto(mapCtx *direct.MapContext, in *krm.AnnotationCustomizedStruct) *pb.AnnotationCustomizedStruct {
	if in == nil {
		return nil
	}
	out := &pb.AnnotationCustomizedStruct{}
	// MISSING: Elements
	return out
}
func AnnotationList_FromProto(mapCtx *direct.MapContext, in *pb.AnnotationList) *krm.AnnotationList {
	if in == nil {
		return nil
	}
	out := &krm.AnnotationList{}
	out.Values = direct.Slice_FromProto(mapCtx, in.Values, AnnotationValue_FromProto)
	return out
}
func AnnotationList_ToProto(mapCtx *direct.MapContext, in *krm.AnnotationList) *pb.AnnotationList {
	if in == nil {
		return nil
	}
	out := &pb.AnnotationList{}
	out.Values = direct.Slice_ToProto(mapCtx, in.Values, AnnotationValue_ToProto)
	return out
}
func AnnotationValue_FromProto(mapCtx *direct.MapContext, in *pb.AnnotationValue) *krm.AnnotationValue {
	if in == nil {
		return nil
	}
	out := &krm.AnnotationValue{}
	out.IntValue = direct.LazyPtr(in.GetIntValue())
	out.FloatValue = direct.LazyPtr(in.GetFloatValue())
	out.StrValue = direct.LazyPtr(in.GetStrValue())
	out.DatetimeValue = direct.LazyPtr(in.GetDatetimeValue())
	out.GeoCoordinate = GeoCoordinate_FromProto(mapCtx, in.GetGeoCoordinate())
	out.ProtoAnyValue = Any_FromProto(mapCtx, in.GetProtoAnyValue())
	out.BoolValue = direct.LazyPtr(in.GetBoolValue())
	out.CustomizedStructDataValue = CustomizedStructDataValue_FromProto(mapCtx, in.GetCustomizedStructDataValue())
	out.ListValue = AnnotationList_FromProto(mapCtx, in.GetListValue())
	out.CustomizedStructValue = AnnotationCustomizedStruct_FromProto(mapCtx, in.GetCustomizedStructValue())
	return out
}
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
	if oneof := CustomizedStructDataValue_ToProto(mapCtx, in.CustomizedStructDataValue); oneof != nil {
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
func GeoCoordinate_FromProto(mapCtx *direct.MapContext, in *pb.GeoCoordinate) *krm.GeoCoordinate {
	if in == nil {
		return nil
	}
	out := &krm.GeoCoordinate{}
	out.Latitude = direct.LazyPtr(in.GetLatitude())
	out.Longitude = direct.LazyPtr(in.GetLongitude())
	return out
}
func GeoCoordinate_ToProto(mapCtx *direct.MapContext, in *krm.GeoCoordinate) *pb.GeoCoordinate {
	if in == nil {
		return nil
	}
	out := &pb.GeoCoordinate{}
	out.Latitude = direct.ValueOf(in.Latitude)
	out.Longitude = direct.ValueOf(in.Longitude)
	return out
}
func Partition_FromProto(mapCtx *direct.MapContext, in *pb.Partition) *krm.Partition {
	if in == nil {
		return nil
	}
	out := &krm.Partition{}
	out.TemporalPartition = Partition_TemporalPartition_FromProto(mapCtx, in.GetTemporalPartition())
	out.SpatialPartition = Partition_SpatialPartition_FromProto(mapCtx, in.GetSpatialPartition())
	out.RelativeTemporalPartition = Partition_RelativeTemporalPartition_FromProto(mapCtx, in.GetRelativeTemporalPartition())
	return out
}
func Partition_ToProto(mapCtx *direct.MapContext, in *krm.Partition) *pb.Partition {
	if in == nil {
		return nil
	}
	out := &pb.Partition{}
	out.TemporalPartition = Partition_TemporalPartition_ToProto(mapCtx, in.TemporalPartition)
	out.SpatialPartition = Partition_SpatialPartition_ToProto(mapCtx, in.SpatialPartition)
	out.RelativeTemporalPartition = Partition_RelativeTemporalPartition_ToProto(mapCtx, in.RelativeTemporalPartition)
	return out
}
func Partition_RelativeTemporalPartition_FromProto(mapCtx *direct.MapContext, in *pb.Partition_RelativeTemporalPartition) *krm.Partition_RelativeTemporalPartition {
	if in == nil {
		return nil
	}
	out := &krm.Partition_RelativeTemporalPartition{}
	out.StartOffset = direct.StringDuration_FromProto(mapCtx, in.GetStartOffset())
	out.EndOffset = direct.StringDuration_FromProto(mapCtx, in.GetEndOffset())
	return out
}
func Partition_RelativeTemporalPartition_ToProto(mapCtx *direct.MapContext, in *krm.Partition_RelativeTemporalPartition) *pb.Partition_RelativeTemporalPartition {
	if in == nil {
		return nil
	}
	out := &pb.Partition_RelativeTemporalPartition{}
	out.StartOffset = direct.StringDuration_ToProto(mapCtx, in.StartOffset)
	out.EndOffset = direct.StringDuration_ToProto(mapCtx, in.EndOffset)
	return out
}
func Partition_SpatialPartition_FromProto(mapCtx *direct.MapContext, in *pb.Partition_SpatialPartition) *krm.Partition_SpatialPartition {
	if in == nil {
		return nil
	}
	out := &krm.Partition_SpatialPartition{}
	out.XMin = in.XMin
	out.YMin = in.YMin
	out.XMax = in.XMax
	out.YMax = in.YMax
	return out
}
func Partition_SpatialPartition_ToProto(mapCtx *direct.MapContext, in *krm.Partition_SpatialPartition) *pb.Partition_SpatialPartition {
	if in == nil {
		return nil
	}
	out := &pb.Partition_SpatialPartition{}
	out.XMin = in.XMin
	out.YMin = in.YMin
	out.XMax = in.XMax
	out.YMax = in.YMax
	return out
}
func Partition_TemporalPartition_FromProto(mapCtx *direct.MapContext, in *pb.Partition_TemporalPartition) *krm.Partition_TemporalPartition {
	if in == nil {
		return nil
	}
	out := &krm.Partition_TemporalPartition{}
	out.StartTime = direct.StringTimestamp_FromProto(mapCtx, in.GetStartTime())
	out.EndTime = direct.StringTimestamp_FromProto(mapCtx, in.GetEndTime())
	return out
}
func Partition_TemporalPartition_ToProto(mapCtx *direct.MapContext, in *krm.Partition_TemporalPartition) *pb.Partition_TemporalPartition {
	if in == nil {
		return nil
	}
	out := &pb.Partition_TemporalPartition{}
	out.StartTime = direct.StringTimestamp_ToProto(mapCtx, in.StartTime)
	out.EndTime = direct.StringTimestamp_ToProto(mapCtx, in.EndTime)
	return out
}
func UserSpecifiedAnnotation_FromProto(mapCtx *direct.MapContext, in *pb.UserSpecifiedAnnotation) *krm.UserSpecifiedAnnotation {
	if in == nil {
		return nil
	}
	out := &krm.UserSpecifiedAnnotation{}
	out.Key = direct.LazyPtr(in.GetKey())
	out.Value = AnnotationValue_FromProto(mapCtx, in.GetValue())
	out.Partition = Partition_FromProto(mapCtx, in.GetPartition())
	return out
}
func UserSpecifiedAnnotation_ToProto(mapCtx *direct.MapContext, in *krm.UserSpecifiedAnnotation) *pb.UserSpecifiedAnnotation {
	if in == nil {
		return nil
	}
	out := &pb.UserSpecifiedAnnotation{}
	out.Key = direct.ValueOf(in.Key)
	out.Value = AnnotationValue_ToProto(mapCtx, in.Value)
	out.Partition = Partition_ToProto(mapCtx, in.Partition)
	return out
}
func VisionaiAnnotationObservedState_FromProto(mapCtx *direct.MapContext, in *pb.Annotation) *krm.VisionaiAnnotationObservedState {
	if in == nil {
		return nil
	}
	out := &krm.VisionaiAnnotationObservedState{}
	// MISSING: Name
	// MISSING: UserSpecifiedAnnotation
	return out
}
func VisionaiAnnotationObservedState_ToProto(mapCtx *direct.MapContext, in *krm.VisionaiAnnotationObservedState) *pb.Annotation {
	if in == nil {
		return nil
	}
	out := &pb.Annotation{}
	// MISSING: Name
	// MISSING: UserSpecifiedAnnotation
	return out
}
func VisionaiAnnotationSpec_FromProto(mapCtx *direct.MapContext, in *pb.Annotation) *krm.VisionaiAnnotationSpec {
	if in == nil {
		return nil
	}
	out := &krm.VisionaiAnnotationSpec{}
	// MISSING: Name
	// MISSING: UserSpecifiedAnnotation
	return out
}
func VisionaiAnnotationSpec_ToProto(mapCtx *direct.MapContext, in *krm.VisionaiAnnotationSpec) *pb.Annotation {
	if in == nil {
		return nil
	}
	out := &pb.Annotation{}
	// MISSING: Name
	// MISSING: UserSpecifiedAnnotation
	return out
}
