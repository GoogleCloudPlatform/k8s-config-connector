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
	pb "cloud.google.com/go/documentai/apiv1/documentaipb"
	krmv1beta1 "github.com/GoogleCloudPlatform/k8s-config-connector/apis/documentai/v1beta1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct"
)

func DocumentSchema_Metadata_v1beta1_FromProto(mapCtx *direct.MapContext, in *pb.DocumentSchema_Metadata) *krmv1beta1.DocumentSchema_Metadata {
	if in == nil {
		return nil
	}
	out := &krmv1beta1.DocumentSchema_Metadata{}
	out.DocumentAllowMultipleLabels = direct.LazyPtr(in.DocumentAllowMultipleLabels)
	out.DocumentSplitter = direct.LazyPtr(in.DocumentSplitter)
	out.PrefixedNamingOnProperties = direct.LazyPtr(in.PrefixedNamingOnProperties)
	out.SkipNamingValidation = direct.LazyPtr(in.SkipNamingValidation)
	return out
}

func DocumentSchema_Metadata_v1beta1_ToProto(mapCtx *direct.MapContext, in *krmv1beta1.DocumentSchema_Metadata) *pb.DocumentSchema_Metadata {
	if in == nil {
		return nil
	}
	out := &pb.DocumentSchema_Metadata{}
	out.DocumentAllowMultipleLabels = direct.ValueOf(in.DocumentAllowMultipleLabels)
	out.DocumentSplitter = direct.ValueOf(in.DocumentSplitter)
	out.PrefixedNamingOnProperties = direct.ValueOf(in.PrefixedNamingOnProperties)
	out.SkipNamingValidation = direct.ValueOf(in.SkipNamingValidation)
	return out
}

func Evaluation_Metrics_v1beta1_FromProto(mapCtx *direct.MapContext, in *pb.Evaluation_Metrics) *krmv1beta1.Evaluation_Metrics {
	if in == nil {
		return nil
	}
	out := &krmv1beta1.Evaluation_Metrics{}
	out.F1Score = direct.LazyPtr(in.F1Score)
	out.GroundTruthDocumentCount = direct.LazyPtr(in.GroundTruthDocumentCount)
	out.FalsePositivesCount = direct.LazyPtr(in.FalsePositivesCount)
	out.GroundTruthOccurrencesCount = direct.LazyPtr(in.GroundTruthOccurrencesCount)
	out.FalseNegativesCount = direct.LazyPtr(in.FalseNegativesCount)
	out.Precision = direct.LazyPtr(in.Precision)
	out.PredictedDocumentCount = direct.LazyPtr(in.PredictedDocumentCount)
	out.PredictedOccurrencesCount = direct.LazyPtr(in.PredictedOccurrencesCount)
	out.Recall = direct.LazyPtr(in.Recall)
	out.TotalDocumentsCount = direct.LazyPtr(in.TotalDocumentsCount)
	out.TruePositivesCount = direct.LazyPtr(in.TruePositivesCount)
	return out
}
func Evaluation_Metrics_v1beta1_ToProto(mapCtx *direct.MapContext, in *krmv1beta1.Evaluation_Metrics) *pb.Evaluation_Metrics {
	if in == nil {
		return nil
	}
	out := &pb.Evaluation_Metrics{}
	out.F1Score = direct.ValueOf(in.F1Score)
	out.GroundTruthDocumentCount = direct.ValueOf(in.GroundTruthDocumentCount)
	out.FalsePositivesCount = direct.ValueOf(in.FalsePositivesCount)
	out.GroundTruthOccurrencesCount = direct.ValueOf(in.GroundTruthOccurrencesCount)
	out.FalseNegativesCount = direct.ValueOf(in.FalseNegativesCount)
	out.Precision = direct.ValueOf(in.Precision)
	out.PredictedDocumentCount = direct.ValueOf(in.PredictedDocumentCount)
	out.PredictedOccurrencesCount = direct.ValueOf(in.PredictedOccurrencesCount)
	out.Recall = direct.ValueOf(in.Recall)
	out.TotalDocumentsCount = direct.ValueOf(in.TotalDocumentsCount)
	out.TruePositivesCount = direct.ValueOf(in.TruePositivesCount)
	return out
}
