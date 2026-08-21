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

package automl

import (
	pb "cloud.google.com/go/automl/apiv1/automlpb"
	krm "github.com/GoogleCloudPlatform/k8s-config-connector/apis/automl/v1alpha1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct"
)

func TranslationDatasetMetadata_FromProto(mapCtx *direct.MapContext, in *pb.TranslationDatasetMetadata) *krm.TranslationDatasetMetadata {
	if in == nil {
		return nil
	}
	out := &krm.TranslationDatasetMetadata{}
	out.SourceLanguageCode = direct.LazyPtr(in.GetSourceLanguageCode())
	out.TargetLanguageCode = direct.LazyPtr(in.GetTargetLanguageCode())
	return out
}
func TranslationDatasetMetadata_ToProto(mapCtx *direct.MapContext, in *krm.TranslationDatasetMetadata) *pb.TranslationDatasetMetadata {
	if in == nil {
		return nil
	}
	out := &pb.TranslationDatasetMetadata{}
	out.SourceLanguageCode = direct.ValueOf(in.SourceLanguageCode)
	out.TargetLanguageCode = direct.ValueOf(in.TargetLanguageCode)
	return out
}

func ImageClassificationDatasetMetadata_FromProto(mapCtx *direct.MapContext, in *pb.ImageClassificationDatasetMetadata) *krm.ImageClassificationDatasetMetadata {
	if in == nil {
		return nil
	}
	out := &krm.ImageClassificationDatasetMetadata{}
	// ClassificationType is an enum in proto
	out.ClassificationType = direct.Enum_FromProto(mapCtx, in.GetClassificationType())
	return out
}
func ImageClassificationDatasetMetadata_ToProto(mapCtx *direct.MapContext, in *krm.ImageClassificationDatasetMetadata) *pb.ImageClassificationDatasetMetadata {
	if in == nil {
		return nil
	}
	out := &pb.ImageClassificationDatasetMetadata{}
	out.ClassificationType = direct.Enum_ToProto[pb.ClassificationType](mapCtx, in.ClassificationType)
	return out
}

func TextClassificationDatasetMetadata_FromProto(mapCtx *direct.MapContext, in *pb.TextClassificationDatasetMetadata) *krm.TextClassificationDatasetMetadata {
	if in == nil {
		return nil
	}
	out := &krm.TextClassificationDatasetMetadata{}
	out.ClassificationType = direct.Enum_FromProto(mapCtx, in.GetClassificationType())
	return out
}
func TextClassificationDatasetMetadata_ToProto(mapCtx *direct.MapContext, in *krm.TextClassificationDatasetMetadata) *pb.TextClassificationDatasetMetadata {
	if in == nil {
		return nil
	}
	out := &pb.TextClassificationDatasetMetadata{}
	out.ClassificationType = direct.Enum_ToProto[pb.ClassificationType](mapCtx, in.ClassificationType)
	return out
}

func ImageObjectDetectionDatasetMetadata_FromProto(mapCtx *direct.MapContext, in *pb.ImageObjectDetectionDatasetMetadata) *krm.ImageObjectDetectionDatasetMetadata {
	if in == nil {
		return nil
	}
	out := &krm.ImageObjectDetectionDatasetMetadata{}
	return out
}
func ImageObjectDetectionDatasetMetadata_ToProto(mapCtx *direct.MapContext, in *krm.ImageObjectDetectionDatasetMetadata) *pb.ImageObjectDetectionDatasetMetadata {
	if in == nil {
		return nil
	}
	out := &pb.ImageObjectDetectionDatasetMetadata{}
	return out
}

func TextExtractionDatasetMetadata_FromProto(mapCtx *direct.MapContext, in *pb.TextExtractionDatasetMetadata) *krm.TextExtractionDatasetMetadata {
	if in == nil {
		return nil
	}
	out := &krm.TextExtractionDatasetMetadata{}
	return out
}
func TextExtractionDatasetMetadata_ToProto(mapCtx *direct.MapContext, in *krm.TextExtractionDatasetMetadata) *pb.TextExtractionDatasetMetadata {
	if in == nil {
		return nil
	}
	out := &pb.TextExtractionDatasetMetadata{}
	return out
}

func TextSentimentDatasetMetadata_FromProto(mapCtx *direct.MapContext, in *pb.TextSentimentDatasetMetadata) *krm.TextSentimentDatasetMetadata {
	if in == nil {
		return nil
	}
	out := &krm.TextSentimentDatasetMetadata{}
	out.SentimentMax = direct.LazyPtr(in.GetSentimentMax())
	return out
}
func TextSentimentDatasetMetadata_ToProto(mapCtx *direct.MapContext, in *krm.TextSentimentDatasetMetadata) *pb.TextSentimentDatasetMetadata {
	if in == nil {
		return nil
	}
	out := &pb.TextSentimentDatasetMetadata{}
	out.SentimentMax = direct.ValueOf(in.SentimentMax)
	return out
}
