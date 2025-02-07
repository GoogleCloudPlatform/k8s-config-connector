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
	krm "github.com/GoogleCloudPlatform/k8s-config-connector/apis/automl/v1alpha1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct"
	refs "github.com/GoogleCloudPlatform/k8s-config-connector/apis/refs/v1beta1"
	pb "cloud.google.com/go/automl/apiv1/automlpb"
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
func Dataset_FromProto(mapCtx *direct.MapContext, in *pb.Dataset) *krm.Dataset {
	if in == nil {
		return nil
	}
	out := &krm.Dataset{}
	out.TranslationDatasetMetadata = TranslationDatasetMetadata_FromProto(mapCtx, in.GetTranslationDatasetMetadata())
	out.ImageClassificationDatasetMetadata = ImageClassificationDatasetMetadata_FromProto(mapCtx, in.GetImageClassificationDatasetMetadata())
	out.TextClassificationDatasetMetadata = TextClassificationDatasetMetadata_FromProto(mapCtx, in.GetTextClassificationDatasetMetadata())
	out.ImageObjectDetectionDatasetMetadata = ImageObjectDetectionDatasetMetadata_FromProto(mapCtx, in.GetImageObjectDetectionDatasetMetadata())
	out.TextExtractionDatasetMetadata = TextExtractionDatasetMetadata_FromProto(mapCtx, in.GetTextExtractionDatasetMetadata())
	out.TextSentimentDatasetMetadata = TextSentimentDatasetMetadata_FromProto(mapCtx, in.GetTextSentimentDatasetMetadata())
	out.Name = direct.LazyPtr(in.GetName())
	out.DisplayName = direct.LazyPtr(in.GetDisplayName())
	out.Description = direct.LazyPtr(in.GetDescription())
	out.ExampleCount = direct.LazyPtr(in.GetExampleCount())
	out.CreateTime = direct.StringTimestamp_FromProto(mapCtx, in.GetCreateTime())
	out.Etag = direct.LazyPtr(in.GetEtag())
	out.Labels = in.Labels
	return out
}
func Dataset_ToProto(mapCtx *direct.MapContext, in *krm.Dataset) *pb.Dataset {
	if in == nil {
		return nil
	}
	out := &pb.Dataset{}
	if oneof := TranslationDatasetMetadata_ToProto(mapCtx, in.TranslationDatasetMetadata); oneof != nil {
		out.DatasetMetadata = &pb.Dataset_TranslationDatasetMetadata{TranslationDatasetMetadata: oneof}
	}
	if oneof := ImageClassificationDatasetMetadata_ToProto(mapCtx, in.ImageClassificationDatasetMetadata); oneof != nil {
		out.DatasetMetadata = &pb.Dataset_ImageClassificationDatasetMetadata{ImageClassificationDatasetMetadata: oneof}
	}
	if oneof := TextClassificationDatasetMetadata_ToProto(mapCtx, in.TextClassificationDatasetMetadata); oneof != nil {
		out.DatasetMetadata = &pb.Dataset_TextClassificationDatasetMetadata{TextClassificationDatasetMetadata: oneof}
	}
	if oneof := ImageObjectDetectionDatasetMetadata_ToProto(mapCtx, in.ImageObjectDetectionDatasetMetadata); oneof != nil {
		out.DatasetMetadata = &pb.Dataset_ImageObjectDetectionDatasetMetadata{ImageObjectDetectionDatasetMetadata: oneof}
	}
	if oneof := TextExtractionDatasetMetadata_ToProto(mapCtx, in.TextExtractionDatasetMetadata); oneof != nil {
		out.DatasetMetadata = &pb.Dataset_TextExtractionDatasetMetadata{TextExtractionDatasetMetadata: oneof}
	}
	if oneof := TextSentimentDatasetMetadata_ToProto(mapCtx, in.TextSentimentDatasetMetadata); oneof != nil {
		out.DatasetMetadata = &pb.Dataset_TextSentimentDatasetMetadata{TextSentimentDatasetMetadata: oneof}
	}
	out.Name = direct.ValueOf(in.Name)
	out.DisplayName = direct.ValueOf(in.DisplayName)
	out.Description = direct.ValueOf(in.Description)
	out.ExampleCount = direct.ValueOf(in.ExampleCount)
	out.CreateTime = direct.StringTimestamp_ToProto(mapCtx, in.CreateTime)
	out.Etag = direct.ValueOf(in.Etag)
	out.Labels = in.Labels
	return out
}
