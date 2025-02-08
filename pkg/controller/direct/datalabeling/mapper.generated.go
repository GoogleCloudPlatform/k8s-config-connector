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

package datalabeling

import (
	refs "github.com/GoogleCloudPlatform/k8s-config-connector/apis/refs/v1beta1"
	pb "cloud.google.com/go/datalabeling/apiv1beta1/datalabelingpb"
	krm "github.com/GoogleCloudPlatform/k8s-config-connector/apis/datalabeling/v1alpha1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct"
)
func BigQuerySource_FromProto(mapCtx *direct.MapContext, in *pb.BigQuerySource) *krm.BigQuerySource {
	if in == nil {
		return nil
	}
	out := &krm.BigQuerySource{}
	out.InputURI = direct.LazyPtr(in.GetInputUri())
	return out
}
func BigQuerySource_ToProto(mapCtx *direct.MapContext, in *krm.BigQuerySource) *pb.BigQuerySource {
	if in == nil {
		return nil
	}
	out := &pb.BigQuerySource{}
	out.InputUri = direct.ValueOf(in.InputURI)
	return out
}
func DatalabelingDatasetObservedState_FromProto(mapCtx *direct.MapContext, in *pb.Dataset) *krm.DatalabelingDatasetObservedState {
	if in == nil {
		return nil
	}
	out := &krm.DatalabelingDatasetObservedState{}
	// MISSING: Name
	// MISSING: DisplayName
	// MISSING: Description
	// MISSING: CreateTime
	// MISSING: InputConfigs
	// MISSING: BlockingResources
	// MISSING: DataItemCount
	return out
}
func DatalabelingDatasetObservedState_ToProto(mapCtx *direct.MapContext, in *krm.DatalabelingDatasetObservedState) *pb.Dataset {
	if in == nil {
		return nil
	}
	out := &pb.Dataset{}
	// MISSING: Name
	// MISSING: DisplayName
	// MISSING: Description
	// MISSING: CreateTime
	// MISSING: InputConfigs
	// MISSING: BlockingResources
	// MISSING: DataItemCount
	return out
}
func DatalabelingDatasetSpec_FromProto(mapCtx *direct.MapContext, in *pb.Dataset) *krm.DatalabelingDatasetSpec {
	if in == nil {
		return nil
	}
	out := &krm.DatalabelingDatasetSpec{}
	// MISSING: Name
	// MISSING: DisplayName
	// MISSING: Description
	// MISSING: CreateTime
	// MISSING: InputConfigs
	// MISSING: BlockingResources
	// MISSING: DataItemCount
	return out
}
func DatalabelingDatasetSpec_ToProto(mapCtx *direct.MapContext, in *krm.DatalabelingDatasetSpec) *pb.Dataset {
	if in == nil {
		return nil
	}
	out := &pb.Dataset{}
	// MISSING: Name
	// MISSING: DisplayName
	// MISSING: Description
	// MISSING: CreateTime
	// MISSING: InputConfigs
	// MISSING: BlockingResources
	// MISSING: DataItemCount
	return out
}
func Dataset_FromProto(mapCtx *direct.MapContext, in *pb.Dataset) *krm.Dataset {
	if in == nil {
		return nil
	}
	out := &krm.Dataset{}
	out.Name = direct.LazyPtr(in.GetName())
	out.DisplayName = direct.LazyPtr(in.GetDisplayName())
	out.Description = direct.LazyPtr(in.GetDescription())
	out.CreateTime = direct.StringTimestamp_FromProto(mapCtx, in.GetCreateTime())
	out.InputConfigs = direct.Slice_FromProto(mapCtx, in.InputConfigs, InputConfig_FromProto)
	out.BlockingResources = in.BlockingResources
	out.DataItemCount = direct.LazyPtr(in.GetDataItemCount())
	return out
}
func Dataset_ToProto(mapCtx *direct.MapContext, in *krm.Dataset) *pb.Dataset {
	if in == nil {
		return nil
	}
	out := &pb.Dataset{}
	out.Name = direct.ValueOf(in.Name)
	out.DisplayName = direct.ValueOf(in.DisplayName)
	out.Description = direct.ValueOf(in.Description)
	out.CreateTime = direct.StringTimestamp_ToProto(mapCtx, in.CreateTime)
	out.InputConfigs = direct.Slice_ToProto(mapCtx, in.InputConfigs, InputConfig_ToProto)
	out.BlockingResources = in.BlockingResources
	out.DataItemCount = direct.ValueOf(in.DataItemCount)
	return out
}
func GcsSource_FromProto(mapCtx *direct.MapContext, in *pb.GcsSource) *krm.GcsSource {
	if in == nil {
		return nil
	}
	out := &krm.GcsSource{}
	out.InputURI = direct.LazyPtr(in.GetInputUri())
	out.MimeType = direct.LazyPtr(in.GetMimeType())
	return out
}
func GcsSource_ToProto(mapCtx *direct.MapContext, in *krm.GcsSource) *pb.GcsSource {
	if in == nil {
		return nil
	}
	out := &pb.GcsSource{}
	out.InputUri = direct.ValueOf(in.InputURI)
	out.MimeType = direct.ValueOf(in.MimeType)
	return out
}
func InputConfig_FromProto(mapCtx *direct.MapContext, in *pb.InputConfig) *krm.InputConfig {
	if in == nil {
		return nil
	}
	out := &krm.InputConfig{}
	out.TextMetadata = TextMetadata_FromProto(mapCtx, in.GetTextMetadata())
	out.GcsSource = GcsSource_FromProto(mapCtx, in.GetGcsSource())
	out.BigquerySource = BigQuerySource_FromProto(mapCtx, in.GetBigquerySource())
	out.DataType = direct.Enum_FromProto(mapCtx, in.GetDataType())
	out.AnnotationType = direct.Enum_FromProto(mapCtx, in.GetAnnotationType())
	out.ClassificationMetadata = ClassificationMetadata_FromProto(mapCtx, in.GetClassificationMetadata())
	return out
}
func InputConfig_ToProto(mapCtx *direct.MapContext, in *krm.InputConfig) *pb.InputConfig {
	if in == nil {
		return nil
	}
	out := &pb.InputConfig{}
	if oneof := TextMetadata_ToProto(mapCtx, in.TextMetadata); oneof != nil {
		out.DataTypeMetadata = &pb.InputConfig_TextMetadata{TextMetadata: oneof}
	}
	if oneof := GcsSource_ToProto(mapCtx, in.GcsSource); oneof != nil {
		out.Source = &pb.InputConfig_GcsSource{GcsSource: oneof}
	}
	if oneof := BigQuerySource_ToProto(mapCtx, in.BigquerySource); oneof != nil {
		out.Source = &pb.InputConfig_BigquerySource{BigquerySource: oneof}
	}
	out.DataType = direct.Enum_ToProto[pb.DataType](mapCtx, in.DataType)
	out.AnnotationType = direct.Enum_ToProto[pb.AnnotationType](mapCtx, in.AnnotationType)
	out.ClassificationMetadata = ClassificationMetadata_ToProto(mapCtx, in.ClassificationMetadata)
	return out
}
