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
func AnnotationSpec_FromProto(mapCtx *direct.MapContext, in *pb.AnnotationSpec) *krm.AnnotationSpec {
	if in == nil {
		return nil
	}
	out := &krm.AnnotationSpec{}
	out.DisplayName = direct.LazyPtr(in.GetDisplayName())
	out.Description = direct.LazyPtr(in.GetDescription())
	return out
}
func AnnotationSpec_ToProto(mapCtx *direct.MapContext, in *krm.AnnotationSpec) *pb.AnnotationSpec {
	if in == nil {
		return nil
	}
	out := &pb.AnnotationSpec{}
	out.DisplayName = direct.ValueOf(in.DisplayName)
	out.Description = direct.ValueOf(in.Description)
	return out
}
func BoundingBoxEvaluationOptions_FromProto(mapCtx *direct.MapContext, in *pb.BoundingBoxEvaluationOptions) *krm.BoundingBoxEvaluationOptions {
	if in == nil {
		return nil
	}
	out := &krm.BoundingBoxEvaluationOptions{}
	out.IouThreshold = direct.LazyPtr(in.GetIouThreshold())
	return out
}
func BoundingBoxEvaluationOptions_ToProto(mapCtx *direct.MapContext, in *krm.BoundingBoxEvaluationOptions) *pb.BoundingBoxEvaluationOptions {
	if in == nil {
		return nil
	}
	out := &pb.BoundingBoxEvaluationOptions{}
	out.IouThreshold = direct.ValueOf(in.IouThreshold)
	return out
}
func ClassificationMetrics_FromProto(mapCtx *direct.MapContext, in *pb.ClassificationMetrics) *krm.ClassificationMetrics {
	if in == nil {
		return nil
	}
	out := &krm.ClassificationMetrics{}
	out.PrCurve = PrCurve_FromProto(mapCtx, in.GetPrCurve())
	out.ConfusionMatrix = ConfusionMatrix_FromProto(mapCtx, in.GetConfusionMatrix())
	return out
}
func ClassificationMetrics_ToProto(mapCtx *direct.MapContext, in *krm.ClassificationMetrics) *pb.ClassificationMetrics {
	if in == nil {
		return nil
	}
	out := &pb.ClassificationMetrics{}
	out.PrCurve = PrCurve_ToProto(mapCtx, in.PrCurve)
	out.ConfusionMatrix = ConfusionMatrix_ToProto(mapCtx, in.ConfusionMatrix)
	return out
}
func ConfusionMatrix_FromProto(mapCtx *direct.MapContext, in *pb.ConfusionMatrix) *krm.ConfusionMatrix {
	if in == nil {
		return nil
	}
	out := &krm.ConfusionMatrix{}
	out.Row = direct.Slice_FromProto(mapCtx, in.Row, ConfusionMatrix_Row_FromProto)
	return out
}
func ConfusionMatrix_ToProto(mapCtx *direct.MapContext, in *krm.ConfusionMatrix) *pb.ConfusionMatrix {
	if in == nil {
		return nil
	}
	out := &pb.ConfusionMatrix{}
	out.Row = direct.Slice_ToProto(mapCtx, in.Row, ConfusionMatrix_Row_ToProto)
	return out
}
func ConfusionMatrix_ConfusionMatrixEntry_FromProto(mapCtx *direct.MapContext, in *pb.ConfusionMatrix_ConfusionMatrixEntry) *krm.ConfusionMatrix_ConfusionMatrixEntry {
	if in == nil {
		return nil
	}
	out := &krm.ConfusionMatrix_ConfusionMatrixEntry{}
	out.AnnotationSpec = AnnotationSpec_FromProto(mapCtx, in.GetAnnotationSpec())
	out.ItemCount = direct.LazyPtr(in.GetItemCount())
	return out
}
func ConfusionMatrix_ConfusionMatrixEntry_ToProto(mapCtx *direct.MapContext, in *krm.ConfusionMatrix_ConfusionMatrixEntry) *pb.ConfusionMatrix_ConfusionMatrixEntry {
	if in == nil {
		return nil
	}
	out := &pb.ConfusionMatrix_ConfusionMatrixEntry{}
	out.AnnotationSpec = AnnotationSpec_ToProto(mapCtx, in.AnnotationSpec)
	out.ItemCount = direct.ValueOf(in.ItemCount)
	return out
}
func ConfusionMatrix_Row_FromProto(mapCtx *direct.MapContext, in *pb.ConfusionMatrix_Row) *krm.ConfusionMatrix_Row {
	if in == nil {
		return nil
	}
	out := &krm.ConfusionMatrix_Row{}
	out.AnnotationSpec = AnnotationSpec_FromProto(mapCtx, in.GetAnnotationSpec())
	out.Entries = direct.Slice_FromProto(mapCtx, in.Entries, ConfusionMatrix_ConfusionMatrixEntry_FromProto)
	return out
}
func ConfusionMatrix_Row_ToProto(mapCtx *direct.MapContext, in *krm.ConfusionMatrix_Row) *pb.ConfusionMatrix_Row {
	if in == nil {
		return nil
	}
	out := &pb.ConfusionMatrix_Row{}
	out.AnnotationSpec = AnnotationSpec_ToProto(mapCtx, in.AnnotationSpec)
	out.Entries = direct.Slice_ToProto(mapCtx, in.Entries, ConfusionMatrix_ConfusionMatrixEntry_ToProto)
	return out
}
func DatalabelingEvaluationObservedState_FromProto(mapCtx *direct.MapContext, in *pb.Evaluation) *krm.DatalabelingEvaluationObservedState {
	if in == nil {
		return nil
	}
	out := &krm.DatalabelingEvaluationObservedState{}
	// MISSING: Name
	// MISSING: Config
	// MISSING: EvaluationJobRunTime
	// MISSING: CreateTime
	// MISSING: EvaluationMetrics
	// MISSING: AnnotationType
	// MISSING: EvaluatedItemCount
	return out
}
func DatalabelingEvaluationObservedState_ToProto(mapCtx *direct.MapContext, in *krm.DatalabelingEvaluationObservedState) *pb.Evaluation {
	if in == nil {
		return nil
	}
	out := &pb.Evaluation{}
	// MISSING: Name
	// MISSING: Config
	// MISSING: EvaluationJobRunTime
	// MISSING: CreateTime
	// MISSING: EvaluationMetrics
	// MISSING: AnnotationType
	// MISSING: EvaluatedItemCount
	return out
}
func DatalabelingEvaluationSpec_FromProto(mapCtx *direct.MapContext, in *pb.Evaluation) *krm.DatalabelingEvaluationSpec {
	if in == nil {
		return nil
	}
	out := &krm.DatalabelingEvaluationSpec{}
	// MISSING: Name
	// MISSING: Config
	// MISSING: EvaluationJobRunTime
	// MISSING: CreateTime
	// MISSING: EvaluationMetrics
	// MISSING: AnnotationType
	// MISSING: EvaluatedItemCount
	return out
}
func DatalabelingEvaluationSpec_ToProto(mapCtx *direct.MapContext, in *krm.DatalabelingEvaluationSpec) *pb.Evaluation {
	if in == nil {
		return nil
	}
	out := &pb.Evaluation{}
	// MISSING: Name
	// MISSING: Config
	// MISSING: EvaluationJobRunTime
	// MISSING: CreateTime
	// MISSING: EvaluationMetrics
	// MISSING: AnnotationType
	// MISSING: EvaluatedItemCount
	return out
}
func Evaluation_FromProto(mapCtx *direct.MapContext, in *pb.Evaluation) *krm.Evaluation {
	if in == nil {
		return nil
	}
	out := &krm.Evaluation{}
	out.Name = direct.LazyPtr(in.GetName())
	out.Config = EvaluationConfig_FromProto(mapCtx, in.GetConfig())
	out.EvaluationJobRunTime = direct.StringTimestamp_FromProto(mapCtx, in.GetEvaluationJobRunTime())
	out.CreateTime = direct.StringTimestamp_FromProto(mapCtx, in.GetCreateTime())
	out.EvaluationMetrics = EvaluationMetrics_FromProto(mapCtx, in.GetEvaluationMetrics())
	out.AnnotationType = direct.Enum_FromProto(mapCtx, in.GetAnnotationType())
	out.EvaluatedItemCount = direct.LazyPtr(in.GetEvaluatedItemCount())
	return out
}
func Evaluation_ToProto(mapCtx *direct.MapContext, in *krm.Evaluation) *pb.Evaluation {
	if in == nil {
		return nil
	}
	out := &pb.Evaluation{}
	out.Name = direct.ValueOf(in.Name)
	out.Config = EvaluationConfig_ToProto(mapCtx, in.Config)
	out.EvaluationJobRunTime = direct.StringTimestamp_ToProto(mapCtx, in.EvaluationJobRunTime)
	out.CreateTime = direct.StringTimestamp_ToProto(mapCtx, in.CreateTime)
	out.EvaluationMetrics = EvaluationMetrics_ToProto(mapCtx, in.EvaluationMetrics)
	out.AnnotationType = direct.Enum_ToProto[pb.AnnotationType](mapCtx, in.AnnotationType)
	out.EvaluatedItemCount = direct.ValueOf(in.EvaluatedItemCount)
	return out
}
func EvaluationConfig_FromProto(mapCtx *direct.MapContext, in *pb.EvaluationConfig) *krm.EvaluationConfig {
	if in == nil {
		return nil
	}
	out := &krm.EvaluationConfig{}
	out.BoundingBoxEvaluationOptions = BoundingBoxEvaluationOptions_FromProto(mapCtx, in.GetBoundingBoxEvaluationOptions())
	return out
}
func EvaluationConfig_ToProto(mapCtx *direct.MapContext, in *krm.EvaluationConfig) *pb.EvaluationConfig {
	if in == nil {
		return nil
	}
	out := &pb.EvaluationConfig{}
	if oneof := BoundingBoxEvaluationOptions_ToProto(mapCtx, in.BoundingBoxEvaluationOptions); oneof != nil {
		out.VerticalOption = &pb.EvaluationConfig_BoundingBoxEvaluationOptions{BoundingBoxEvaluationOptions: oneof}
	}
	return out
}
func EvaluationMetrics_FromProto(mapCtx *direct.MapContext, in *pb.EvaluationMetrics) *krm.EvaluationMetrics {
	if in == nil {
		return nil
	}
	out := &krm.EvaluationMetrics{}
	out.ClassificationMetrics = ClassificationMetrics_FromProto(mapCtx, in.GetClassificationMetrics())
	out.ObjectDetectionMetrics = ObjectDetectionMetrics_FromProto(mapCtx, in.GetObjectDetectionMetrics())
	return out
}
func EvaluationMetrics_ToProto(mapCtx *direct.MapContext, in *krm.EvaluationMetrics) *pb.EvaluationMetrics {
	if in == nil {
		return nil
	}
	out := &pb.EvaluationMetrics{}
	if oneof := ClassificationMetrics_ToProto(mapCtx, in.ClassificationMetrics); oneof != nil {
		out.Metrics = &pb.EvaluationMetrics_ClassificationMetrics{ClassificationMetrics: oneof}
	}
	if oneof := ObjectDetectionMetrics_ToProto(mapCtx, in.ObjectDetectionMetrics); oneof != nil {
		out.Metrics = &pb.EvaluationMetrics_ObjectDetectionMetrics{ObjectDetectionMetrics: oneof}
	}
	return out
}
func ObjectDetectionMetrics_FromProto(mapCtx *direct.MapContext, in *pb.ObjectDetectionMetrics) *krm.ObjectDetectionMetrics {
	if in == nil {
		return nil
	}
	out := &krm.ObjectDetectionMetrics{}
	out.PrCurve = PrCurve_FromProto(mapCtx, in.GetPrCurve())
	return out
}
func ObjectDetectionMetrics_ToProto(mapCtx *direct.MapContext, in *krm.ObjectDetectionMetrics) *pb.ObjectDetectionMetrics {
	if in == nil {
		return nil
	}
	out := &pb.ObjectDetectionMetrics{}
	out.PrCurve = PrCurve_ToProto(mapCtx, in.PrCurve)
	return out
}
func PrCurve_FromProto(mapCtx *direct.MapContext, in *pb.PrCurve) *krm.PrCurve {
	if in == nil {
		return nil
	}
	out := &krm.PrCurve{}
	out.AnnotationSpec = AnnotationSpec_FromProto(mapCtx, in.GetAnnotationSpec())
	out.AreaUnderCurve = direct.LazyPtr(in.GetAreaUnderCurve())
	out.ConfidenceMetricsEntries = direct.Slice_FromProto(mapCtx, in.ConfidenceMetricsEntries, PrCurve_ConfidenceMetricsEntry_FromProto)
	out.MeanAveragePrecision = direct.LazyPtr(in.GetMeanAveragePrecision())
	return out
}
func PrCurve_ToProto(mapCtx *direct.MapContext, in *krm.PrCurve) *pb.PrCurve {
	if in == nil {
		return nil
	}
	out := &pb.PrCurve{}
	out.AnnotationSpec = AnnotationSpec_ToProto(mapCtx, in.AnnotationSpec)
	out.AreaUnderCurve = direct.ValueOf(in.AreaUnderCurve)
	out.ConfidenceMetricsEntries = direct.Slice_ToProto(mapCtx, in.ConfidenceMetricsEntries, PrCurve_ConfidenceMetricsEntry_ToProto)
	out.MeanAveragePrecision = direct.ValueOf(in.MeanAveragePrecision)
	return out
}
func PrCurve_ConfidenceMetricsEntry_FromProto(mapCtx *direct.MapContext, in *pb.PrCurve_ConfidenceMetricsEntry) *krm.PrCurve_ConfidenceMetricsEntry {
	if in == nil {
		return nil
	}
	out := &krm.PrCurve_ConfidenceMetricsEntry{}
	out.ConfidenceThreshold = direct.LazyPtr(in.GetConfidenceThreshold())
	out.Recall = direct.LazyPtr(in.GetRecall())
	out.Precision = direct.LazyPtr(in.GetPrecision())
	out.F1Score = direct.LazyPtr(in.GetF1Score())
	out.RecallAt1 = direct.LazyPtr(in.GetRecallAt1())
	out.PrecisionAt1 = direct.LazyPtr(in.GetPrecisionAt1())
	out.F1ScoreAt1 = direct.LazyPtr(in.GetF1ScoreAt1())
	out.RecallAt5 = direct.LazyPtr(in.GetRecallAt5())
	out.PrecisionAt5 = direct.LazyPtr(in.GetPrecisionAt5())
	out.F1ScoreAt5 = direct.LazyPtr(in.GetF1ScoreAt5())
	return out
}
func PrCurve_ConfidenceMetricsEntry_ToProto(mapCtx *direct.MapContext, in *krm.PrCurve_ConfidenceMetricsEntry) *pb.PrCurve_ConfidenceMetricsEntry {
	if in == nil {
		return nil
	}
	out := &pb.PrCurve_ConfidenceMetricsEntry{}
	out.ConfidenceThreshold = direct.ValueOf(in.ConfidenceThreshold)
	out.Recall = direct.ValueOf(in.Recall)
	out.Precision = direct.ValueOf(in.Precision)
	out.F1Score = direct.ValueOf(in.F1Score)
	out.RecallAt1 = direct.ValueOf(in.RecallAt1)
	out.PrecisionAt1 = direct.ValueOf(in.PrecisionAt1)
	out.F1ScoreAt1 = direct.ValueOf(in.F1ScoreAt1)
	out.RecallAt5 = direct.ValueOf(in.RecallAt5)
	out.PrecisionAt5 = direct.ValueOf(in.PrecisionAt5)
	out.F1ScoreAt5 = direct.ValueOf(in.F1ScoreAt5)
	return out
}
