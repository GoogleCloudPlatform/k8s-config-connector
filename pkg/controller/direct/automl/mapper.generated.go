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

package automl

import (
	refs "github.com/GoogleCloudPlatform/k8s-config-connector/apis/refs/v1beta1"
	pb "cloud.google.com/go/automl/apiv1/automlpb"
	krm "github.com/GoogleCloudPlatform/k8s-config-connector/apis/automl/v1alpha1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct"
)
func AutomlAnnotationSpecObservedState_FromProto(mapCtx *direct.MapContext, in *pb.AnnotationSpec) *krm.AutomlAnnotationSpecObservedState {
	if in == nil {
		return nil
	}
	out := &krm.AutomlAnnotationSpecObservedState{}
	// MISSING: Name
	// MISSING: DisplayName
	// MISSING: ExampleCount
	return out
}
func AutomlAnnotationSpecObservedState_ToProto(mapCtx *direct.MapContext, in *krm.AutomlAnnotationSpecObservedState) *pb.AnnotationSpec {
	if in == nil {
		return nil
	}
	out := &pb.AnnotationSpec{}
	// MISSING: Name
	// MISSING: DisplayName
	// MISSING: ExampleCount
	return out
}
func AutomlAnnotationSpecSpec_FromProto(mapCtx *direct.MapContext, in *pb.AnnotationSpec) *krm.AutomlAnnotationSpecSpec {
	if in == nil {
		return nil
	}
	out := &krm.AutomlAnnotationSpecSpec{}
	// MISSING: Name
	// MISSING: DisplayName
	// MISSING: ExampleCount
	return out
}
func AutomlAnnotationSpecSpec_ToProto(mapCtx *direct.MapContext, in *krm.AutomlAnnotationSpecSpec) *pb.AnnotationSpec {
	if in == nil {
		return nil
	}
	out := &pb.AnnotationSpec{}
	// MISSING: Name
	// MISSING: DisplayName
	// MISSING: ExampleCount
	return out
}
func AutomlDatasetObservedState_FromProto(mapCtx *direct.MapContext, in *pb.Dataset) *krm.AutomlDatasetObservedState {
	if in == nil {
		return nil
	}
	out := &krm.AutomlDatasetObservedState{}
	// MISSING: TranslationDatasetMetadata
	// MISSING: ImageClassificationDatasetMetadata
	// MISSING: TextClassificationDatasetMetadata
	// MISSING: ImageObjectDetectionDatasetMetadata
	// MISSING: TextExtractionDatasetMetadata
	// MISSING: TextSentimentDatasetMetadata
	// MISSING: Name
	// MISSING: DisplayName
	// MISSING: Description
	// MISSING: ExampleCount
	// MISSING: CreateTime
	// MISSING: Etag
	// MISSING: Labels
	return out
}
func AutomlDatasetObservedState_ToProto(mapCtx *direct.MapContext, in *krm.AutomlDatasetObservedState) *pb.Dataset {
	if in == nil {
		return nil
	}
	out := &pb.Dataset{}
	// MISSING: TranslationDatasetMetadata
	// MISSING: ImageClassificationDatasetMetadata
	// MISSING: TextClassificationDatasetMetadata
	// MISSING: ImageObjectDetectionDatasetMetadata
	// MISSING: TextExtractionDatasetMetadata
	// MISSING: TextSentimentDatasetMetadata
	// MISSING: Name
	// MISSING: DisplayName
	// MISSING: Description
	// MISSING: ExampleCount
	// MISSING: CreateTime
	// MISSING: Etag
	// MISSING: Labels
	return out
}
func AutomlDatasetSpec_FromProto(mapCtx *direct.MapContext, in *pb.Dataset) *krm.AutomlDatasetSpec {
	if in == nil {
		return nil
	}
	out := &krm.AutomlDatasetSpec{}
	// MISSING: TranslationDatasetMetadata
	// MISSING: ImageClassificationDatasetMetadata
	// MISSING: TextClassificationDatasetMetadata
	// MISSING: ImageObjectDetectionDatasetMetadata
	// MISSING: TextExtractionDatasetMetadata
	// MISSING: TextSentimentDatasetMetadata
	// MISSING: Name
	// MISSING: DisplayName
	// MISSING: Description
	// MISSING: ExampleCount
	// MISSING: CreateTime
	// MISSING: Etag
	// MISSING: Labels
	return out
}
func AutomlDatasetSpec_ToProto(mapCtx *direct.MapContext, in *krm.AutomlDatasetSpec) *pb.Dataset {
	if in == nil {
		return nil
	}
	out := &pb.Dataset{}
	// MISSING: TranslationDatasetMetadata
	// MISSING: ImageClassificationDatasetMetadata
	// MISSING: TextClassificationDatasetMetadata
	// MISSING: ImageObjectDetectionDatasetMetadata
	// MISSING: TextExtractionDatasetMetadata
	// MISSING: TextSentimentDatasetMetadata
	// MISSING: Name
	// MISSING: DisplayName
	// MISSING: Description
	// MISSING: ExampleCount
	// MISSING: CreateTime
	// MISSING: Etag
	// MISSING: Labels
	return out
}
func AutomlModelObservedState_FromProto(mapCtx *direct.MapContext, in *pb.Model) *krm.AutomlModelObservedState {
	if in == nil {
		return nil
	}
	out := &krm.AutomlModelObservedState{}
	// MISSING: TranslationModelMetadata
	// MISSING: ImageClassificationModelMetadata
	// MISSING: TextClassificationModelMetadata
	// MISSING: ImageObjectDetectionModelMetadata
	// MISSING: TextExtractionModelMetadata
	// MISSING: TextSentimentModelMetadata
	// MISSING: Name
	// MISSING: DisplayName
	// MISSING: DatasetID
	// MISSING: CreateTime
	// MISSING: UpdateTime
	// MISSING: DeploymentState
	// MISSING: Etag
	// MISSING: Labels
	return out
}
func AutomlModelObservedState_ToProto(mapCtx *direct.MapContext, in *krm.AutomlModelObservedState) *pb.Model {
	if in == nil {
		return nil
	}
	out := &pb.Model{}
	// MISSING: TranslationModelMetadata
	// MISSING: ImageClassificationModelMetadata
	// MISSING: TextClassificationModelMetadata
	// MISSING: ImageObjectDetectionModelMetadata
	// MISSING: TextExtractionModelMetadata
	// MISSING: TextSentimentModelMetadata
	// MISSING: Name
	// MISSING: DisplayName
	// MISSING: DatasetID
	// MISSING: CreateTime
	// MISSING: UpdateTime
	// MISSING: DeploymentState
	// MISSING: Etag
	// MISSING: Labels
	return out
}
func AutomlModelSpec_FromProto(mapCtx *direct.MapContext, in *pb.Model) *krm.AutomlModelSpec {
	if in == nil {
		return nil
	}
	out := &krm.AutomlModelSpec{}
	// MISSING: TranslationModelMetadata
	// MISSING: ImageClassificationModelMetadata
	// MISSING: TextClassificationModelMetadata
	// MISSING: ImageObjectDetectionModelMetadata
	// MISSING: TextExtractionModelMetadata
	// MISSING: TextSentimentModelMetadata
	// MISSING: Name
	// MISSING: DisplayName
	// MISSING: DatasetID
	// MISSING: CreateTime
	// MISSING: UpdateTime
	// MISSING: DeploymentState
	// MISSING: Etag
	// MISSING: Labels
	return out
}
func AutomlModelSpec_ToProto(mapCtx *direct.MapContext, in *krm.AutomlModelSpec) *pb.Model {
	if in == nil {
		return nil
	}
	out := &pb.Model{}
	// MISSING: TranslationModelMetadata
	// MISSING: ImageClassificationModelMetadata
	// MISSING: TextClassificationModelMetadata
	// MISSING: ImageObjectDetectionModelMetadata
	// MISSING: TextExtractionModelMetadata
	// MISSING: TextSentimentModelMetadata
	// MISSING: Name
	// MISSING: DisplayName
	// MISSING: DatasetID
	// MISSING: CreateTime
	// MISSING: UpdateTime
	// MISSING: DeploymentState
	// MISSING: Etag
	// MISSING: Labels
	return out
}
func Model_FromProto(mapCtx *direct.MapContext, in *pb.Model) *krm.Model {
	if in == nil {
		return nil
	}
	out := &krm.Model{}
	out.TranslationModelMetadata = TranslationModelMetadata_FromProto(mapCtx, in.GetTranslationModelMetadata())
	out.ImageClassificationModelMetadata = ImageClassificationModelMetadata_FromProto(mapCtx, in.GetImageClassificationModelMetadata())
	out.TextClassificationModelMetadata = TextClassificationModelMetadata_FromProto(mapCtx, in.GetTextClassificationModelMetadata())
	out.ImageObjectDetectionModelMetadata = ImageObjectDetectionModelMetadata_FromProto(mapCtx, in.GetImageObjectDetectionModelMetadata())
	out.TextExtractionModelMetadata = TextExtractionModelMetadata_FromProto(mapCtx, in.GetTextExtractionModelMetadata())
	out.TextSentimentModelMetadata = TextSentimentModelMetadata_FromProto(mapCtx, in.GetTextSentimentModelMetadata())
	out.Name = direct.LazyPtr(in.GetName())
	out.DisplayName = direct.LazyPtr(in.GetDisplayName())
	out.DatasetID = direct.LazyPtr(in.GetDatasetId())
	out.CreateTime = direct.StringTimestamp_FromProto(mapCtx, in.GetCreateTime())
	out.UpdateTime = direct.StringTimestamp_FromProto(mapCtx, in.GetUpdateTime())
	out.DeploymentState = direct.Enum_FromProto(mapCtx, in.GetDeploymentState())
	out.Etag = direct.LazyPtr(in.GetEtag())
	out.Labels = in.Labels
	return out
}
func Model_ToProto(mapCtx *direct.MapContext, in *krm.Model) *pb.Model {
	if in == nil {
		return nil
	}
	out := &pb.Model{}
	if oneof := TranslationModelMetadata_ToProto(mapCtx, in.TranslationModelMetadata); oneof != nil {
		out.ModelMetadata = &pb.Model_TranslationModelMetadata{TranslationModelMetadata: oneof}
	}
	if oneof := ImageClassificationModelMetadata_ToProto(mapCtx, in.ImageClassificationModelMetadata); oneof != nil {
		out.ModelMetadata = &pb.Model_ImageClassificationModelMetadata{ImageClassificationModelMetadata: oneof}
	}
	if oneof := TextClassificationModelMetadata_ToProto(mapCtx, in.TextClassificationModelMetadata); oneof != nil {
		out.ModelMetadata = &pb.Model_TextClassificationModelMetadata{TextClassificationModelMetadata: oneof}
	}
	if oneof := ImageObjectDetectionModelMetadata_ToProto(mapCtx, in.ImageObjectDetectionModelMetadata); oneof != nil {
		out.ModelMetadata = &pb.Model_ImageObjectDetectionModelMetadata{ImageObjectDetectionModelMetadata: oneof}
	}
	if oneof := TextExtractionModelMetadata_ToProto(mapCtx, in.TextExtractionModelMetadata); oneof != nil {
		out.ModelMetadata = &pb.Model_TextExtractionModelMetadata{TextExtractionModelMetadata: oneof}
	}
	if oneof := TextSentimentModelMetadata_ToProto(mapCtx, in.TextSentimentModelMetadata); oneof != nil {
		out.ModelMetadata = &pb.Model_TextSentimentModelMetadata{TextSentimentModelMetadata: oneof}
	}
	out.Name = direct.ValueOf(in.Name)
	out.DisplayName = direct.ValueOf(in.DisplayName)
	out.DatasetId = direct.ValueOf(in.DatasetID)
	out.CreateTime = direct.StringTimestamp_ToProto(mapCtx, in.CreateTime)
	out.UpdateTime = direct.StringTimestamp_ToProto(mapCtx, in.UpdateTime)
	out.DeploymentState = direct.Enum_ToProto[pb.Model_DeploymentState](mapCtx, in.DeploymentState)
	out.Etag = direct.ValueOf(in.Etag)
	out.Labels = in.Labels
	return out
}
func ModelObservedState_FromProto(mapCtx *direct.MapContext, in *pb.Model) *krm.ModelObservedState {
	if in == nil {
		return nil
	}
	out := &krm.ModelObservedState{}
	// MISSING: TranslationModelMetadata
	out.ImageClassificationModelMetadata = ImageClassificationModelMetadataObservedState_FromProto(mapCtx, in.GetImageClassificationModelMetadata())
	// MISSING: TextClassificationModelMetadata
	out.ImageObjectDetectionModelMetadata = ImageObjectDetectionModelMetadataObservedState_FromProto(mapCtx, in.GetImageObjectDetectionModelMetadata())
	// MISSING: TextExtractionModelMetadata
	// MISSING: TextSentimentModelMetadata
	// MISSING: Name
	// MISSING: DisplayName
	// MISSING: DatasetID
	// MISSING: CreateTime
	// MISSING: UpdateTime
	// MISSING: DeploymentState
	// MISSING: Etag
	// MISSING: Labels
	return out
}
func ModelObservedState_ToProto(mapCtx *direct.MapContext, in *krm.ModelObservedState) *pb.Model {
	if in == nil {
		return nil
	}
	out := &pb.Model{}
	// MISSING: TranslationModelMetadata
	if oneof := ImageClassificationModelMetadataObservedState_ToProto(mapCtx, in.ImageClassificationModelMetadata); oneof != nil {
		out.ModelMetadata = &pb.Model_ImageClassificationModelMetadata{ImageClassificationModelMetadata: oneof}
	}
	// MISSING: TextClassificationModelMetadata
	if oneof := ImageObjectDetectionModelMetadataObservedState_ToProto(mapCtx, in.ImageObjectDetectionModelMetadata); oneof != nil {
		out.ModelMetadata = &pb.Model_ImageObjectDetectionModelMetadata{ImageObjectDetectionModelMetadata: oneof}
	}
	// MISSING: TextExtractionModelMetadata
	// MISSING: TextSentimentModelMetadata
	// MISSING: Name
	// MISSING: DisplayName
	// MISSING: DatasetID
	// MISSING: CreateTime
	// MISSING: UpdateTime
	// MISSING: DeploymentState
	// MISSING: Etag
	// MISSING: Labels
	return out
}
