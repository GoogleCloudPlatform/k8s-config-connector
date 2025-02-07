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
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct"
	refs "github.com/GoogleCloudPlatform/k8s-config-connector/apis/refs/v1beta1"
	pb "cloud.google.com/go/automl/apiv1/automlpb"
	krm "github.com/GoogleCloudPlatform/k8s-config-connector/apis/automl/v1alpha1"
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
func AutomlModelEvaluationObservedState_FromProto(mapCtx *direct.MapContext, in *pb.ModelEvaluation) *krm.AutomlModelEvaluationObservedState {
	if in == nil {
		return nil
	}
	out := &krm.AutomlModelEvaluationObservedState{}
	// MISSING: ClassificationEvaluationMetrics
	// MISSING: TranslationEvaluationMetrics
	// MISSING: ImageObjectDetectionEvaluationMetrics
	// MISSING: TextSentimentEvaluationMetrics
	// MISSING: TextExtractionEvaluationMetrics
	// MISSING: Name
	// MISSING: AnnotationSpecID
	// MISSING: DisplayName
	// MISSING: CreateTime
	// MISSING: EvaluatedExampleCount
	return out
}
func AutomlModelEvaluationObservedState_ToProto(mapCtx *direct.MapContext, in *krm.AutomlModelEvaluationObservedState) *pb.ModelEvaluation {
	if in == nil {
		return nil
	}
	out := &pb.ModelEvaluation{}
	// MISSING: ClassificationEvaluationMetrics
	// MISSING: TranslationEvaluationMetrics
	// MISSING: ImageObjectDetectionEvaluationMetrics
	// MISSING: TextSentimentEvaluationMetrics
	// MISSING: TextExtractionEvaluationMetrics
	// MISSING: Name
	// MISSING: AnnotationSpecID
	// MISSING: DisplayName
	// MISSING: CreateTime
	// MISSING: EvaluatedExampleCount
	return out
}
func AutomlModelEvaluationSpec_FromProto(mapCtx *direct.MapContext, in *pb.ModelEvaluation) *krm.AutomlModelEvaluationSpec {
	if in == nil {
		return nil
	}
	out := &krm.AutomlModelEvaluationSpec{}
	// MISSING: ClassificationEvaluationMetrics
	// MISSING: TranslationEvaluationMetrics
	// MISSING: ImageObjectDetectionEvaluationMetrics
	// MISSING: TextSentimentEvaluationMetrics
	// MISSING: TextExtractionEvaluationMetrics
	// MISSING: Name
	// MISSING: AnnotationSpecID
	// MISSING: DisplayName
	// MISSING: CreateTime
	// MISSING: EvaluatedExampleCount
	return out
}
func AutomlModelEvaluationSpec_ToProto(mapCtx *direct.MapContext, in *krm.AutomlModelEvaluationSpec) *pb.ModelEvaluation {
	if in == nil {
		return nil
	}
	out := &pb.ModelEvaluation{}
	// MISSING: ClassificationEvaluationMetrics
	// MISSING: TranslationEvaluationMetrics
	// MISSING: ImageObjectDetectionEvaluationMetrics
	// MISSING: TextSentimentEvaluationMetrics
	// MISSING: TextExtractionEvaluationMetrics
	// MISSING: Name
	// MISSING: AnnotationSpecID
	// MISSING: DisplayName
	// MISSING: CreateTime
	// MISSING: EvaluatedExampleCount
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
func BoundingBoxMetricsEntry_FromProto(mapCtx *direct.MapContext, in *pb.BoundingBoxMetricsEntry) *krm.BoundingBoxMetricsEntry {
	if in == nil {
		return nil
	}
	out := &krm.BoundingBoxMetricsEntry{}
	out.IouThreshold = direct.LazyPtr(in.GetIouThreshold())
	out.MeanAveragePrecision = direct.LazyPtr(in.GetMeanAveragePrecision())
	out.ConfidenceMetricsEntries = direct.Slice_FromProto(mapCtx, in.ConfidenceMetricsEntries, BoundingBoxMetricsEntry_ConfidenceMetricsEntry_FromProto)
	return out
}
func BoundingBoxMetricsEntry_ToProto(mapCtx *direct.MapContext, in *krm.BoundingBoxMetricsEntry) *pb.BoundingBoxMetricsEntry {
	if in == nil {
		return nil
	}
	out := &pb.BoundingBoxMetricsEntry{}
	out.IouThreshold = direct.ValueOf(in.IouThreshold)
	out.MeanAveragePrecision = direct.ValueOf(in.MeanAveragePrecision)
	out.ConfidenceMetricsEntries = direct.Slice_ToProto(mapCtx, in.ConfidenceMetricsEntries, BoundingBoxMetricsEntry_ConfidenceMetricsEntry_ToProto)
	return out
}
func BoundingBoxMetricsEntry_ConfidenceMetricsEntry_FromProto(mapCtx *direct.MapContext, in *pb.BoundingBoxMetricsEntry_ConfidenceMetricsEntry) *krm.BoundingBoxMetricsEntry_ConfidenceMetricsEntry {
	if in == nil {
		return nil
	}
	out := &krm.BoundingBoxMetricsEntry_ConfidenceMetricsEntry{}
	out.ConfidenceThreshold = direct.LazyPtr(in.GetConfidenceThreshold())
	out.Recall = direct.LazyPtr(in.GetRecall())
	out.Precision = direct.LazyPtr(in.GetPrecision())
	out.F1Score = direct.LazyPtr(in.GetF1Score())
	return out
}
func BoundingBoxMetricsEntry_ConfidenceMetricsEntry_ToProto(mapCtx *direct.MapContext, in *krm.BoundingBoxMetricsEntry_ConfidenceMetricsEntry) *pb.BoundingBoxMetricsEntry_ConfidenceMetricsEntry {
	if in == nil {
		return nil
	}
	out := &pb.BoundingBoxMetricsEntry_ConfidenceMetricsEntry{}
	out.ConfidenceThreshold = direct.ValueOf(in.ConfidenceThreshold)
	out.Recall = direct.ValueOf(in.Recall)
	out.Precision = direct.ValueOf(in.Precision)
	out.F1Score = direct.ValueOf(in.F1Score)
	return out
}
func ClassificationEvaluationMetrics_FromProto(mapCtx *direct.MapContext, in *pb.ClassificationEvaluationMetrics) *krm.ClassificationEvaluationMetrics {
	if in == nil {
		return nil
	}
	out := &krm.ClassificationEvaluationMetrics{}
	out.AuPrc = direct.LazyPtr(in.GetAuPrc())
	out.AuRoc = direct.LazyPtr(in.GetAuRoc())
	out.LogLoss = direct.LazyPtr(in.GetLogLoss())
	out.ConfidenceMetricsEntry = direct.Slice_FromProto(mapCtx, in.ConfidenceMetricsEntry, ClassificationEvaluationMetrics_ConfidenceMetricsEntry_FromProto)
	out.ConfusionMatrix = ClassificationEvaluationMetrics_ConfusionMatrix_FromProto(mapCtx, in.GetConfusionMatrix())
	out.AnnotationSpecID = in.AnnotationSpecId
	return out
}
func ClassificationEvaluationMetrics_ToProto(mapCtx *direct.MapContext, in *krm.ClassificationEvaluationMetrics) *pb.ClassificationEvaluationMetrics {
	if in == nil {
		return nil
	}
	out := &pb.ClassificationEvaluationMetrics{}
	out.AuPrc = direct.ValueOf(in.AuPrc)
	out.AuRoc = direct.ValueOf(in.AuRoc)
	out.LogLoss = direct.ValueOf(in.LogLoss)
	out.ConfidenceMetricsEntry = direct.Slice_ToProto(mapCtx, in.ConfidenceMetricsEntry, ClassificationEvaluationMetrics_ConfidenceMetricsEntry_ToProto)
	out.ConfusionMatrix = ClassificationEvaluationMetrics_ConfusionMatrix_ToProto(mapCtx, in.ConfusionMatrix)
	out.AnnotationSpecId = in.AnnotationSpecID
	return out
}
func ClassificationEvaluationMetrics_ConfidenceMetricsEntry_FromProto(mapCtx *direct.MapContext, in *pb.ClassificationEvaluationMetrics_ConfidenceMetricsEntry) *krm.ClassificationEvaluationMetrics_ConfidenceMetricsEntry {
	if in == nil {
		return nil
	}
	out := &krm.ClassificationEvaluationMetrics_ConfidenceMetricsEntry{}
	out.ConfidenceThreshold = direct.LazyPtr(in.GetConfidenceThreshold())
	out.PositionThreshold = direct.LazyPtr(in.GetPositionThreshold())
	out.Recall = direct.LazyPtr(in.GetRecall())
	out.Precision = direct.LazyPtr(in.GetPrecision())
	out.FalsePositiveRate = direct.LazyPtr(in.GetFalsePositiveRate())
	out.F1Score = direct.LazyPtr(in.GetF1Score())
	out.RecallAt1 = direct.LazyPtr(in.GetRecallAt1())
	out.PrecisionAt1 = direct.LazyPtr(in.GetPrecisionAt1())
	out.FalsePositiveRateAt1 = direct.LazyPtr(in.GetFalsePositiveRateAt1())
	out.F1ScoreAt1 = direct.LazyPtr(in.GetF1ScoreAt1())
	out.TruePositiveCount = direct.LazyPtr(in.GetTruePositiveCount())
	out.FalsePositiveCount = direct.LazyPtr(in.GetFalsePositiveCount())
	out.FalseNegativeCount = direct.LazyPtr(in.GetFalseNegativeCount())
	out.TrueNegativeCount = direct.LazyPtr(in.GetTrueNegativeCount())
	return out
}
func ClassificationEvaluationMetrics_ConfidenceMetricsEntry_ToProto(mapCtx *direct.MapContext, in *krm.ClassificationEvaluationMetrics_ConfidenceMetricsEntry) *pb.ClassificationEvaluationMetrics_ConfidenceMetricsEntry {
	if in == nil {
		return nil
	}
	out := &pb.ClassificationEvaluationMetrics_ConfidenceMetricsEntry{}
	out.ConfidenceThreshold = direct.ValueOf(in.ConfidenceThreshold)
	out.PositionThreshold = direct.ValueOf(in.PositionThreshold)
	out.Recall = direct.ValueOf(in.Recall)
	out.Precision = direct.ValueOf(in.Precision)
	out.FalsePositiveRate = direct.ValueOf(in.FalsePositiveRate)
	out.F1Score = direct.ValueOf(in.F1Score)
	out.RecallAt1 = direct.ValueOf(in.RecallAt1)
	out.PrecisionAt1 = direct.ValueOf(in.PrecisionAt1)
	out.FalsePositiveRateAt1 = direct.ValueOf(in.FalsePositiveRateAt1)
	out.F1ScoreAt1 = direct.ValueOf(in.F1ScoreAt1)
	out.TruePositiveCount = direct.ValueOf(in.TruePositiveCount)
	out.FalsePositiveCount = direct.ValueOf(in.FalsePositiveCount)
	out.FalseNegativeCount = direct.ValueOf(in.FalseNegativeCount)
	out.TrueNegativeCount = direct.ValueOf(in.TrueNegativeCount)
	return out
}
func ClassificationEvaluationMetrics_ConfusionMatrix_FromProto(mapCtx *direct.MapContext, in *pb.ClassificationEvaluationMetrics_ConfusionMatrix) *krm.ClassificationEvaluationMetrics_ConfusionMatrix {
	if in == nil {
		return nil
	}
	out := &krm.ClassificationEvaluationMetrics_ConfusionMatrix{}
	out.AnnotationSpecID = in.AnnotationSpecId
	out.DisplayName = in.DisplayName
	out.Row = direct.Slice_FromProto(mapCtx, in.Row, ClassificationEvaluationMetrics_ConfusionMatrix_Row_FromProto)
	return out
}
func ClassificationEvaluationMetrics_ConfusionMatrix_ToProto(mapCtx *direct.MapContext, in *krm.ClassificationEvaluationMetrics_ConfusionMatrix) *pb.ClassificationEvaluationMetrics_ConfusionMatrix {
	if in == nil {
		return nil
	}
	out := &pb.ClassificationEvaluationMetrics_ConfusionMatrix{}
	out.AnnotationSpecId = in.AnnotationSpecID
	out.DisplayName = in.DisplayName
	out.Row = direct.Slice_ToProto(mapCtx, in.Row, ClassificationEvaluationMetrics_ConfusionMatrix_Row_ToProto)
	return out
}
func ClassificationEvaluationMetrics_ConfusionMatrix_Row_FromProto(mapCtx *direct.MapContext, in *pb.ClassificationEvaluationMetrics_ConfusionMatrix_Row) *krm.ClassificationEvaluationMetrics_ConfusionMatrix_Row {
	if in == nil {
		return nil
	}
	out := &krm.ClassificationEvaluationMetrics_ConfusionMatrix_Row{}
	out.ExampleCount = in.ExampleCount
	return out
}
func ClassificationEvaluationMetrics_ConfusionMatrix_Row_ToProto(mapCtx *direct.MapContext, in *krm.ClassificationEvaluationMetrics_ConfusionMatrix_Row) *pb.ClassificationEvaluationMetrics_ConfusionMatrix_Row {
	if in == nil {
		return nil
	}
	out := &pb.ClassificationEvaluationMetrics_ConfusionMatrix_Row{}
	out.ExampleCount = in.ExampleCount
	return out
}
func ImageObjectDetectionEvaluationMetrics_FromProto(mapCtx *direct.MapContext, in *pb.ImageObjectDetectionEvaluationMetrics) *krm.ImageObjectDetectionEvaluationMetrics {
	if in == nil {
		return nil
	}
	out := &krm.ImageObjectDetectionEvaluationMetrics{}
	out.EvaluatedBoundingBoxCount = direct.LazyPtr(in.GetEvaluatedBoundingBoxCount())
	out.BoundingBoxMetricsEntries = direct.Slice_FromProto(mapCtx, in.BoundingBoxMetricsEntries, BoundingBoxMetricsEntry_FromProto)
	out.BoundingBoxMeanAveragePrecision = direct.LazyPtr(in.GetBoundingBoxMeanAveragePrecision())
	return out
}
func ImageObjectDetectionEvaluationMetrics_ToProto(mapCtx *direct.MapContext, in *krm.ImageObjectDetectionEvaluationMetrics) *pb.ImageObjectDetectionEvaluationMetrics {
	if in == nil {
		return nil
	}
	out := &pb.ImageObjectDetectionEvaluationMetrics{}
	out.EvaluatedBoundingBoxCount = direct.ValueOf(in.EvaluatedBoundingBoxCount)
	out.BoundingBoxMetricsEntries = direct.Slice_ToProto(mapCtx, in.BoundingBoxMetricsEntries, BoundingBoxMetricsEntry_ToProto)
	out.BoundingBoxMeanAveragePrecision = direct.ValueOf(in.BoundingBoxMeanAveragePrecision)
	return out
}
func ModelEvaluation_FromProto(mapCtx *direct.MapContext, in *pb.ModelEvaluation) *krm.ModelEvaluation {
	if in == nil {
		return nil
	}
	out := &krm.ModelEvaluation{}
	out.ClassificationEvaluationMetrics = ClassificationEvaluationMetrics_FromProto(mapCtx, in.GetClassificationEvaluationMetrics())
	out.TranslationEvaluationMetrics = TranslationEvaluationMetrics_FromProto(mapCtx, in.GetTranslationEvaluationMetrics())
	out.ImageObjectDetectionEvaluationMetrics = ImageObjectDetectionEvaluationMetrics_FromProto(mapCtx, in.GetImageObjectDetectionEvaluationMetrics())
	out.TextSentimentEvaluationMetrics = TextSentimentEvaluationMetrics_FromProto(mapCtx, in.GetTextSentimentEvaluationMetrics())
	out.TextExtractionEvaluationMetrics = TextExtractionEvaluationMetrics_FromProto(mapCtx, in.GetTextExtractionEvaluationMetrics())
	out.Name = direct.LazyPtr(in.GetName())
	out.AnnotationSpecID = direct.LazyPtr(in.GetAnnotationSpecId())
	out.DisplayName = direct.LazyPtr(in.GetDisplayName())
	out.CreateTime = direct.StringTimestamp_FromProto(mapCtx, in.GetCreateTime())
	out.EvaluatedExampleCount = direct.LazyPtr(in.GetEvaluatedExampleCount())
	return out
}
func ModelEvaluation_ToProto(mapCtx *direct.MapContext, in *krm.ModelEvaluation) *pb.ModelEvaluation {
	if in == nil {
		return nil
	}
	out := &pb.ModelEvaluation{}
	if oneof := ClassificationEvaluationMetrics_ToProto(mapCtx, in.ClassificationEvaluationMetrics); oneof != nil {
		out.Metrics = &pb.ModelEvaluation_ClassificationEvaluationMetrics{ClassificationEvaluationMetrics: oneof}
	}
	if oneof := TranslationEvaluationMetrics_ToProto(mapCtx, in.TranslationEvaluationMetrics); oneof != nil {
		out.Metrics = &pb.ModelEvaluation_TranslationEvaluationMetrics{TranslationEvaluationMetrics: oneof}
	}
	if oneof := ImageObjectDetectionEvaluationMetrics_ToProto(mapCtx, in.ImageObjectDetectionEvaluationMetrics); oneof != nil {
		out.Metrics = &pb.ModelEvaluation_ImageObjectDetectionEvaluationMetrics{ImageObjectDetectionEvaluationMetrics: oneof}
	}
	if oneof := TextSentimentEvaluationMetrics_ToProto(mapCtx, in.TextSentimentEvaluationMetrics); oneof != nil {
		out.Metrics = &pb.ModelEvaluation_TextSentimentEvaluationMetrics{TextSentimentEvaluationMetrics: oneof}
	}
	if oneof := TextExtractionEvaluationMetrics_ToProto(mapCtx, in.TextExtractionEvaluationMetrics); oneof != nil {
		out.Metrics = &pb.ModelEvaluation_TextExtractionEvaluationMetrics{TextExtractionEvaluationMetrics: oneof}
	}
	out.Name = direct.ValueOf(in.Name)
	out.AnnotationSpecId = direct.ValueOf(in.AnnotationSpecID)
	out.DisplayName = direct.ValueOf(in.DisplayName)
	out.CreateTime = direct.StringTimestamp_ToProto(mapCtx, in.CreateTime)
	out.EvaluatedExampleCount = direct.ValueOf(in.EvaluatedExampleCount)
	return out
}
func TextExtractionEvaluationMetrics_FromProto(mapCtx *direct.MapContext, in *pb.TextExtractionEvaluationMetrics) *krm.TextExtractionEvaluationMetrics {
	if in == nil {
		return nil
	}
	out := &krm.TextExtractionEvaluationMetrics{}
	out.AuPrc = direct.LazyPtr(in.GetAuPrc())
	out.ConfidenceMetricsEntries = direct.Slice_FromProto(mapCtx, in.ConfidenceMetricsEntries, TextExtractionEvaluationMetrics_ConfidenceMetricsEntry_FromProto)
	return out
}
func TextExtractionEvaluationMetrics_ToProto(mapCtx *direct.MapContext, in *krm.TextExtractionEvaluationMetrics) *pb.TextExtractionEvaluationMetrics {
	if in == nil {
		return nil
	}
	out := &pb.TextExtractionEvaluationMetrics{}
	out.AuPrc = direct.ValueOf(in.AuPrc)
	out.ConfidenceMetricsEntries = direct.Slice_ToProto(mapCtx, in.ConfidenceMetricsEntries, TextExtractionEvaluationMetrics_ConfidenceMetricsEntry_ToProto)
	return out
}
func TextExtractionEvaluationMetrics_ConfidenceMetricsEntry_FromProto(mapCtx *direct.MapContext, in *pb.TextExtractionEvaluationMetrics_ConfidenceMetricsEntry) *krm.TextExtractionEvaluationMetrics_ConfidenceMetricsEntry {
	if in == nil {
		return nil
	}
	out := &krm.TextExtractionEvaluationMetrics_ConfidenceMetricsEntry{}
	out.ConfidenceThreshold = direct.LazyPtr(in.GetConfidenceThreshold())
	out.Recall = direct.LazyPtr(in.GetRecall())
	out.Precision = direct.LazyPtr(in.GetPrecision())
	out.F1Score = direct.LazyPtr(in.GetF1Score())
	return out
}
func TextExtractionEvaluationMetrics_ConfidenceMetricsEntry_ToProto(mapCtx *direct.MapContext, in *krm.TextExtractionEvaluationMetrics_ConfidenceMetricsEntry) *pb.TextExtractionEvaluationMetrics_ConfidenceMetricsEntry {
	if in == nil {
		return nil
	}
	out := &pb.TextExtractionEvaluationMetrics_ConfidenceMetricsEntry{}
	out.ConfidenceThreshold = direct.ValueOf(in.ConfidenceThreshold)
	out.Recall = direct.ValueOf(in.Recall)
	out.Precision = direct.ValueOf(in.Precision)
	out.F1Score = direct.ValueOf(in.F1Score)
	return out
}
func TextSentimentEvaluationMetrics_FromProto(mapCtx *direct.MapContext, in *pb.TextSentimentEvaluationMetrics) *krm.TextSentimentEvaluationMetrics {
	if in == nil {
		return nil
	}
	out := &krm.TextSentimentEvaluationMetrics{}
	out.Precision = direct.LazyPtr(in.GetPrecision())
	out.Recall = direct.LazyPtr(in.GetRecall())
	out.F1Score = direct.LazyPtr(in.GetF1Score())
	out.MeanAbsoluteError = direct.LazyPtr(in.GetMeanAbsoluteError())
	out.MeanSquaredError = direct.LazyPtr(in.GetMeanSquaredError())
	out.LinearKappa = direct.LazyPtr(in.GetLinearKappa())
	out.QuadraticKappa = direct.LazyPtr(in.GetQuadraticKappa())
	out.ConfusionMatrix = ClassificationEvaluationMetrics_ConfusionMatrix_FromProto(mapCtx, in.GetConfusionMatrix())
	return out
}
func TextSentimentEvaluationMetrics_ToProto(mapCtx *direct.MapContext, in *krm.TextSentimentEvaluationMetrics) *pb.TextSentimentEvaluationMetrics {
	if in == nil {
		return nil
	}
	out := &pb.TextSentimentEvaluationMetrics{}
	out.Precision = direct.ValueOf(in.Precision)
	out.Recall = direct.ValueOf(in.Recall)
	out.F1Score = direct.ValueOf(in.F1Score)
	out.MeanAbsoluteError = direct.ValueOf(in.MeanAbsoluteError)
	out.MeanSquaredError = direct.ValueOf(in.MeanSquaredError)
	out.LinearKappa = direct.ValueOf(in.LinearKappa)
	out.QuadraticKappa = direct.ValueOf(in.QuadraticKappa)
	out.ConfusionMatrix = ClassificationEvaluationMetrics_ConfusionMatrix_ToProto(mapCtx, in.ConfusionMatrix)
	return out
}
func TranslationEvaluationMetrics_FromProto(mapCtx *direct.MapContext, in *pb.TranslationEvaluationMetrics) *krm.TranslationEvaluationMetrics {
	if in == nil {
		return nil
	}
	out := &krm.TranslationEvaluationMetrics{}
	out.BleuScore = direct.LazyPtr(in.GetBleuScore())
	out.BaseBleuScore = direct.LazyPtr(in.GetBaseBleuScore())
	return out
}
func TranslationEvaluationMetrics_ToProto(mapCtx *direct.MapContext, in *krm.TranslationEvaluationMetrics) *pb.TranslationEvaluationMetrics {
	if in == nil {
		return nil
	}
	out := &pb.TranslationEvaluationMetrics{}
	out.BleuScore = direct.ValueOf(in.BleuScore)
	out.BaseBleuScore = direct.ValueOf(in.BaseBleuScore)
	return out
}
