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

func ImageClassificationModelMetadata_FromProto(mapCtx *direct.MapContext, in *pb.ImageClassificationModelMetadata) *krm.ImageClassificationModelMetadata {
	if in == nil {
		return nil
	}
	out := &krm.ImageClassificationModelMetadata{}
	out.BaseModelID = direct.LazyPtr(in.GetBaseModelId())
	out.TrainBudgetMilliNodeHours = direct.LazyPtr(in.GetTrainBudgetMilliNodeHours())
	out.ModelType = direct.LazyPtr(in.GetModelType())
	return out
}

func ImageClassificationModelMetadata_ToProto(mapCtx *direct.MapContext, in *krm.ImageClassificationModelMetadata) *pb.ImageClassificationModelMetadata {
	if in == nil {
		return nil
	}
	out := &pb.ImageClassificationModelMetadata{}
	out.BaseModelId = direct.ValueOf(in.BaseModelID)
	out.TrainBudgetMilliNodeHours = direct.ValueOf(in.TrainBudgetMilliNodeHours)
	out.ModelType = direct.ValueOf(in.ModelType)
	return out
}

func ImageClassificationModelMetadataObservedState_FromProto(mapCtx *direct.MapContext, in *pb.ImageClassificationModelMetadata) *krm.ImageClassificationModelMetadataObservedState {
	if in == nil {
		return nil
	}
	out := &krm.ImageClassificationModelMetadataObservedState{}
	out.TrainCostMilliNodeHours = direct.LazyPtr(in.GetTrainCostMilliNodeHours())
	out.StopReason = direct.LazyPtr(in.GetStopReason())
	out.NodeQps = direct.LazyPtr(in.GetNodeQps())
	out.NodeCount = direct.LazyPtr(in.GetNodeCount())
	return out
}

func ImageClassificationModelMetadataObservedState_ToProto(mapCtx *direct.MapContext, in *krm.ImageClassificationModelMetadataObservedState) *pb.ImageClassificationModelMetadata {
	if in == nil {
		return nil
	}
	out := &pb.ImageClassificationModelMetadata{}
	out.TrainCostMilliNodeHours = direct.ValueOf(in.TrainCostMilliNodeHours)
	out.StopReason = direct.ValueOf(in.StopReason)
	out.NodeQps = direct.ValueOf(in.NodeQps)
	out.NodeCount = direct.ValueOf(in.NodeCount)
	return out
}

func ImageObjectDetectionModelMetadata_FromProto(mapCtx *direct.MapContext, in *pb.ImageObjectDetectionModelMetadata) *krm.ImageObjectDetectionModelMetadata {
	if in == nil {
		return nil
	}
	out := &krm.ImageObjectDetectionModelMetadata{}
	out.ModelType = direct.LazyPtr(in.GetModelType())
	out.TrainBudgetMilliNodeHours = direct.LazyPtr(in.GetTrainBudgetMilliNodeHours())
	return out
}

func ImageObjectDetectionModelMetadata_ToProto(mapCtx *direct.MapContext, in *krm.ImageObjectDetectionModelMetadata) *pb.ImageObjectDetectionModelMetadata {
	if in == nil {
		return nil
	}
	out := &pb.ImageObjectDetectionModelMetadata{}
	out.ModelType = direct.ValueOf(in.ModelType)
	out.TrainBudgetMilliNodeHours = direct.ValueOf(in.TrainBudgetMilliNodeHours)
	return out
}

func ImageObjectDetectionModelMetadataObservedState_FromProto(mapCtx *direct.MapContext, in *pb.ImageObjectDetectionModelMetadata) *krm.ImageObjectDetectionModelMetadataObservedState {
	if in == nil {
		return nil
	}
	out := &krm.ImageObjectDetectionModelMetadataObservedState{}
	out.NodeCount = direct.LazyPtr(in.GetNodeCount())
	out.NodeQps = direct.LazyPtr(in.GetNodeQps())
	out.StopReason = direct.LazyPtr(in.GetStopReason())
	out.TrainCostMilliNodeHours = direct.LazyPtr(in.GetTrainCostMilliNodeHours())
	return out
}

func ImageObjectDetectionModelMetadataObservedState_ToProto(mapCtx *direct.MapContext, in *krm.ImageObjectDetectionModelMetadataObservedState) *pb.ImageObjectDetectionModelMetadata {
	if in == nil {
		return nil
	}
	out := &pb.ImageObjectDetectionModelMetadata{}
	out.NodeCount = direct.ValueOf(in.NodeCount)
	out.NodeQps = direct.ValueOf(in.NodeQps)
	out.StopReason = direct.ValueOf(in.StopReason)
	out.TrainCostMilliNodeHours = direct.ValueOf(in.TrainCostMilliNodeHours)
	return out
}

func TextClassificationModelMetadata_FromProto(mapCtx *direct.MapContext, in *pb.TextClassificationModelMetadata) *krm.TextClassificationModelMetadata {
	if in == nil {
		return nil
	}
	out := &krm.TextClassificationModelMetadata{}
	out.ClassificationType = direct.Enum_FromProto(mapCtx, in.GetClassificationType())
	return out
}

func TextClassificationModelMetadata_ToProto(mapCtx *direct.MapContext, in *krm.TextClassificationModelMetadata) *pb.TextClassificationModelMetadata {
	if in == nil {
		return nil
	}
	out := &pb.TextClassificationModelMetadata{}
	out.ClassificationType = direct.Enum_ToProto[pb.ClassificationType](mapCtx, in.ClassificationType)
	return out
}

func TextExtractionModelMetadata_FromProto(mapCtx *direct.MapContext, in *pb.TextExtractionModelMetadata) *krm.TextExtractionModelMetadata {
	if in == nil {
		return nil
	}
	return &krm.TextExtractionModelMetadata{}
}

func TextExtractionModelMetadata_ToProto(mapCtx *direct.MapContext, in *krm.TextExtractionModelMetadata) *pb.TextExtractionModelMetadata {
	if in == nil {
		return nil
	}
	return &pb.TextExtractionModelMetadata{}
}

func TextSentimentModelMetadata_FromProto(mapCtx *direct.MapContext, in *pb.TextSentimentModelMetadata) *krm.TextSentimentModelMetadata {
	if in == nil {
		return nil
	}
	return &krm.TextSentimentModelMetadata{}
}

func TextSentimentModelMetadata_ToProto(mapCtx *direct.MapContext, in *krm.TextSentimentModelMetadata) *pb.TextSentimentModelMetadata {
	if in == nil {
		return nil
	}
	return &pb.TextSentimentModelMetadata{}
}

func TranslationModelMetadata_FromProto(mapCtx *direct.MapContext, in *pb.TranslationModelMetadata) *krm.TranslationModelMetadata {
	if in == nil {
		return nil
	}
	out := &krm.TranslationModelMetadata{}
	out.BaseModel = direct.LazyPtr(in.GetBaseModel())
	out.SourceLanguageCode = direct.LazyPtr(in.GetSourceLanguageCode())
	out.TargetLanguageCode = direct.LazyPtr(in.GetTargetLanguageCode())
	return out
}

func TranslationModelMetadata_ToProto(mapCtx *direct.MapContext, in *krm.TranslationModelMetadata) *pb.TranslationModelMetadata {
	if in == nil {
		return nil
	}
	out := &pb.TranslationModelMetadata{}
	out.BaseModel = direct.ValueOf(in.BaseModel)
	out.SourceLanguageCode = direct.ValueOf(in.SourceLanguageCode)
	out.TargetLanguageCode = direct.ValueOf(in.TargetLanguageCode)
	return out
}
