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

package documentai

import (
	krm "github.com/GoogleCloudPlatform/k8s-config-connector/apis/documentai/v1alpha1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct"
	refs "github.com/GoogleCloudPlatform/k8s-config-connector/apis/refs/v1beta1"
	pb "cloud.google.com/go/documentai/apiv1beta3/documentaipb"
)
func Evaluation_FromProto(mapCtx *direct.MapContext, in *pb.Evaluation) *krm.Evaluation {
	if in == nil {
		return nil
	}
	out := &krm.Evaluation{}
	out.Name = direct.LazyPtr(in.GetName())
	out.CreateTime = direct.StringTimestamp_FromProto(mapCtx, in.GetCreateTime())
	out.DocumentCounters = Evaluation_Counters_FromProto(mapCtx, in.GetDocumentCounters())
	out.AllEntitiesMetrics = Evaluation_MultiConfidenceMetrics_FromProto(mapCtx, in.GetAllEntitiesMetrics())
	// MISSING: EntityMetrics
	out.KMSKeyName = direct.LazyPtr(in.GetKmsKeyName())
	out.KMSKeyVersionName = direct.LazyPtr(in.GetKmsKeyVersionName())
	return out
}
func Evaluation_ToProto(mapCtx *direct.MapContext, in *krm.Evaluation) *pb.Evaluation {
	if in == nil {
		return nil
	}
	out := &pb.Evaluation{}
	out.Name = direct.ValueOf(in.Name)
	out.CreateTime = direct.StringTimestamp_ToProto(mapCtx, in.CreateTime)
	out.DocumentCounters = Evaluation_Counters_ToProto(mapCtx, in.DocumentCounters)
	out.AllEntitiesMetrics = Evaluation_MultiConfidenceMetrics_ToProto(mapCtx, in.AllEntitiesMetrics)
	// MISSING: EntityMetrics
	out.KmsKeyName = direct.ValueOf(in.KMSKeyName)
	out.KmsKeyVersionName = direct.ValueOf(in.KMSKeyVersionName)
	return out
}
func Evaluation_ConfidenceLevelMetrics_FromProto(mapCtx *direct.MapContext, in *pb.Evaluation_ConfidenceLevelMetrics) *krm.Evaluation_ConfidenceLevelMetrics {
	if in == nil {
		return nil
	}
	out := &krm.Evaluation_ConfidenceLevelMetrics{}
	out.ConfidenceLevel = direct.LazyPtr(in.GetConfidenceLevel())
	out.Metrics = Evaluation_Metrics_FromProto(mapCtx, in.GetMetrics())
	return out
}
func Evaluation_ConfidenceLevelMetrics_ToProto(mapCtx *direct.MapContext, in *krm.Evaluation_ConfidenceLevelMetrics) *pb.Evaluation_ConfidenceLevelMetrics {
	if in == nil {
		return nil
	}
	out := &pb.Evaluation_ConfidenceLevelMetrics{}
	out.ConfidenceLevel = direct.ValueOf(in.ConfidenceLevel)
	out.Metrics = Evaluation_Metrics_ToProto(mapCtx, in.Metrics)
	return out
}
func Evaluation_Counters_FromProto(mapCtx *direct.MapContext, in *pb.Evaluation_Counters) *krm.Evaluation_Counters {
	if in == nil {
		return nil
	}
	out := &krm.Evaluation_Counters{}
	out.InputDocumentsCount = direct.LazyPtr(in.GetInputDocumentsCount())
	out.InvalidDocumentsCount = direct.LazyPtr(in.GetInvalidDocumentsCount())
	out.FailedDocumentsCount = direct.LazyPtr(in.GetFailedDocumentsCount())
	out.EvaluatedDocumentsCount = direct.LazyPtr(in.GetEvaluatedDocumentsCount())
	return out
}
func Evaluation_Counters_ToProto(mapCtx *direct.MapContext, in *krm.Evaluation_Counters) *pb.Evaluation_Counters {
	if in == nil {
		return nil
	}
	out := &pb.Evaluation_Counters{}
	out.InputDocumentsCount = direct.ValueOf(in.InputDocumentsCount)
	out.InvalidDocumentsCount = direct.ValueOf(in.InvalidDocumentsCount)
	out.FailedDocumentsCount = direct.ValueOf(in.FailedDocumentsCount)
	out.EvaluatedDocumentsCount = direct.ValueOf(in.EvaluatedDocumentsCount)
	return out
}
func Evaluation_Metrics_FromProto(mapCtx *direct.MapContext, in *pb.Evaluation_Metrics) *krm.Evaluation_Metrics {
	if in == nil {
		return nil
	}
	out := &krm.Evaluation_Metrics{}
	out.Precision = direct.LazyPtr(in.GetPrecision())
	out.Recall = direct.LazyPtr(in.GetRecall())
	out.F1Score = direct.LazyPtr(in.GetF1Score())
	out.PredictedOccurrencesCount = direct.LazyPtr(in.GetPredictedOccurrencesCount())
	out.GroundTruthOccurrencesCount = direct.LazyPtr(in.GetGroundTruthOccurrencesCount())
	out.PredictedDocumentCount = direct.LazyPtr(in.GetPredictedDocumentCount())
	out.GroundTruthDocumentCount = direct.LazyPtr(in.GetGroundTruthDocumentCount())
	out.TruePositivesCount = direct.LazyPtr(in.GetTruePositivesCount())
	out.FalsePositivesCount = direct.LazyPtr(in.GetFalsePositivesCount())
	out.FalseNegativesCount = direct.LazyPtr(in.GetFalseNegativesCount())
	out.TotalDocumentsCount = direct.LazyPtr(in.GetTotalDocumentsCount())
	return out
}
func Evaluation_Metrics_ToProto(mapCtx *direct.MapContext, in *krm.Evaluation_Metrics) *pb.Evaluation_Metrics {
	if in == nil {
		return nil
	}
	out := &pb.Evaluation_Metrics{}
	out.Precision = direct.ValueOf(in.Precision)
	out.Recall = direct.ValueOf(in.Recall)
	out.F1Score = direct.ValueOf(in.F1Score)
	out.PredictedOccurrencesCount = direct.ValueOf(in.PredictedOccurrencesCount)
	out.GroundTruthOccurrencesCount = direct.ValueOf(in.GroundTruthOccurrencesCount)
	out.PredictedDocumentCount = direct.ValueOf(in.PredictedDocumentCount)
	out.GroundTruthDocumentCount = direct.ValueOf(in.GroundTruthDocumentCount)
	out.TruePositivesCount = direct.ValueOf(in.TruePositivesCount)
	out.FalsePositivesCount = direct.ValueOf(in.FalsePositivesCount)
	out.FalseNegativesCount = direct.ValueOf(in.FalseNegativesCount)
	out.TotalDocumentsCount = direct.ValueOf(in.TotalDocumentsCount)
	return out
}
func Evaluation_MultiConfidenceMetrics_FromProto(mapCtx *direct.MapContext, in *pb.Evaluation_MultiConfidenceMetrics) *krm.Evaluation_MultiConfidenceMetrics {
	if in == nil {
		return nil
	}
	out := &krm.Evaluation_MultiConfidenceMetrics{}
	out.ConfidenceLevelMetrics = direct.Slice_FromProto(mapCtx, in.ConfidenceLevelMetrics, Evaluation_ConfidenceLevelMetrics_FromProto)
	out.ConfidenceLevelMetricsExact = direct.Slice_FromProto(mapCtx, in.ConfidenceLevelMetricsExact, Evaluation_ConfidenceLevelMetrics_FromProto)
	out.Auprc = direct.LazyPtr(in.GetAuprc())
	out.EstimatedCalibrationError = direct.LazyPtr(in.GetEstimatedCalibrationError())
	out.AuprcExact = direct.LazyPtr(in.GetAuprcExact())
	out.EstimatedCalibrationErrorExact = direct.LazyPtr(in.GetEstimatedCalibrationErrorExact())
	out.MetricsType = direct.Enum_FromProto(mapCtx, in.GetMetricsType())
	return out
}
func Evaluation_MultiConfidenceMetrics_ToProto(mapCtx *direct.MapContext, in *krm.Evaluation_MultiConfidenceMetrics) *pb.Evaluation_MultiConfidenceMetrics {
	if in == nil {
		return nil
	}
	out := &pb.Evaluation_MultiConfidenceMetrics{}
	out.ConfidenceLevelMetrics = direct.Slice_ToProto(mapCtx, in.ConfidenceLevelMetrics, Evaluation_ConfidenceLevelMetrics_ToProto)
	out.ConfidenceLevelMetricsExact = direct.Slice_ToProto(mapCtx, in.ConfidenceLevelMetricsExact, Evaluation_ConfidenceLevelMetrics_ToProto)
	out.Auprc = direct.ValueOf(in.Auprc)
	out.EstimatedCalibrationError = direct.ValueOf(in.EstimatedCalibrationError)
	out.AuprcExact = direct.ValueOf(in.AuprcExact)
	out.EstimatedCalibrationErrorExact = direct.ValueOf(in.EstimatedCalibrationErrorExact)
	out.MetricsType = direct.Enum_ToProto[pb.Evaluation_MultiConfidenceMetrics_MetricsType](mapCtx, in.MetricsType)
	return out
}
