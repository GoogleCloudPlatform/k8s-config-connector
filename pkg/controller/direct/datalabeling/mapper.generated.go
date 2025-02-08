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
func Attempt_FromProto(mapCtx *direct.MapContext, in *pb.Attempt) *krm.Attempt {
	if in == nil {
		return nil
	}
	out := &krm.Attempt{}
	out.AttemptTime = direct.StringTimestamp_FromProto(mapCtx, in.GetAttemptTime())
	out.PartialFailures = direct.Slice_FromProto(mapCtx, in.PartialFailures, Status_FromProto)
	return out
}
func Attempt_ToProto(mapCtx *direct.MapContext, in *krm.Attempt) *pb.Attempt {
	if in == nil {
		return nil
	}
	out := &pb.Attempt{}
	out.AttemptTime = direct.StringTimestamp_ToProto(mapCtx, in.AttemptTime)
	out.PartialFailures = direct.Slice_ToProto(mapCtx, in.PartialFailures, Status_ToProto)
	return out
}
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
func BoundingPolyConfig_FromProto(mapCtx *direct.MapContext, in *pb.BoundingPolyConfig) *krm.BoundingPolyConfig {
	if in == nil {
		return nil
	}
	out := &krm.BoundingPolyConfig{}
	out.AnnotationSpecSet = direct.LazyPtr(in.GetAnnotationSpecSet())
	out.InstructionMessage = direct.LazyPtr(in.GetInstructionMessage())
	return out
}
func BoundingPolyConfig_ToProto(mapCtx *direct.MapContext, in *krm.BoundingPolyConfig) *pb.BoundingPolyConfig {
	if in == nil {
		return nil
	}
	out := &pb.BoundingPolyConfig{}
	out.AnnotationSpecSet = direct.ValueOf(in.AnnotationSpecSet)
	out.InstructionMessage = direct.ValueOf(in.InstructionMessage)
	return out
}
func DatalabelingEvaluationJobObservedState_FromProto(mapCtx *direct.MapContext, in *pb.EvaluationJob) *krm.DatalabelingEvaluationJobObservedState {
	if in == nil {
		return nil
	}
	out := &krm.DatalabelingEvaluationJobObservedState{}
	// MISSING: Name
	// MISSING: Description
	// MISSING: State
	// MISSING: Schedule
	// MISSING: ModelVersion
	// MISSING: EvaluationJobConfig
	// MISSING: AnnotationSpecSet
	// MISSING: LabelMissingGroundTruth
	// MISSING: Attempts
	// MISSING: CreateTime
	return out
}
func DatalabelingEvaluationJobObservedState_ToProto(mapCtx *direct.MapContext, in *krm.DatalabelingEvaluationJobObservedState) *pb.EvaluationJob {
	if in == nil {
		return nil
	}
	out := &pb.EvaluationJob{}
	// MISSING: Name
	// MISSING: Description
	// MISSING: State
	// MISSING: Schedule
	// MISSING: ModelVersion
	// MISSING: EvaluationJobConfig
	// MISSING: AnnotationSpecSet
	// MISSING: LabelMissingGroundTruth
	// MISSING: Attempts
	// MISSING: CreateTime
	return out
}
func DatalabelingEvaluationJobSpec_FromProto(mapCtx *direct.MapContext, in *pb.EvaluationJob) *krm.DatalabelingEvaluationJobSpec {
	if in == nil {
		return nil
	}
	out := &krm.DatalabelingEvaluationJobSpec{}
	// MISSING: Name
	// MISSING: Description
	// MISSING: State
	// MISSING: Schedule
	// MISSING: ModelVersion
	// MISSING: EvaluationJobConfig
	// MISSING: AnnotationSpecSet
	// MISSING: LabelMissingGroundTruth
	// MISSING: Attempts
	// MISSING: CreateTime
	return out
}
func DatalabelingEvaluationJobSpec_ToProto(mapCtx *direct.MapContext, in *krm.DatalabelingEvaluationJobSpec) *pb.EvaluationJob {
	if in == nil {
		return nil
	}
	out := &pb.EvaluationJob{}
	// MISSING: Name
	// MISSING: Description
	// MISSING: State
	// MISSING: Schedule
	// MISSING: ModelVersion
	// MISSING: EvaluationJobConfig
	// MISSING: AnnotationSpecSet
	// MISSING: LabelMissingGroundTruth
	// MISSING: Attempts
	// MISSING: CreateTime
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
func EvaluationJob_FromProto(mapCtx *direct.MapContext, in *pb.EvaluationJob) *krm.EvaluationJob {
	if in == nil {
		return nil
	}
	out := &krm.EvaluationJob{}
	out.Name = direct.LazyPtr(in.GetName())
	out.Description = direct.LazyPtr(in.GetDescription())
	out.State = direct.Enum_FromProto(mapCtx, in.GetState())
	out.Schedule = direct.LazyPtr(in.GetSchedule())
	out.ModelVersion = direct.LazyPtr(in.GetModelVersion())
	out.EvaluationJobConfig = EvaluationJobConfig_FromProto(mapCtx, in.GetEvaluationJobConfig())
	out.AnnotationSpecSet = direct.LazyPtr(in.GetAnnotationSpecSet())
	out.LabelMissingGroundTruth = direct.LazyPtr(in.GetLabelMissingGroundTruth())
	out.Attempts = direct.Slice_FromProto(mapCtx, in.Attempts, Attempt_FromProto)
	out.CreateTime = direct.StringTimestamp_FromProto(mapCtx, in.GetCreateTime())
	return out
}
func EvaluationJob_ToProto(mapCtx *direct.MapContext, in *krm.EvaluationJob) *pb.EvaluationJob {
	if in == nil {
		return nil
	}
	out := &pb.EvaluationJob{}
	out.Name = direct.ValueOf(in.Name)
	out.Description = direct.ValueOf(in.Description)
	out.State = direct.Enum_ToProto[pb.EvaluationJob_State](mapCtx, in.State)
	out.Schedule = direct.ValueOf(in.Schedule)
	out.ModelVersion = direct.ValueOf(in.ModelVersion)
	out.EvaluationJobConfig = EvaluationJobConfig_ToProto(mapCtx, in.EvaluationJobConfig)
	out.AnnotationSpecSet = direct.ValueOf(in.AnnotationSpecSet)
	out.LabelMissingGroundTruth = direct.ValueOf(in.LabelMissingGroundTruth)
	out.Attempts = direct.Slice_ToProto(mapCtx, in.Attempts, Attempt_ToProto)
	out.CreateTime = direct.StringTimestamp_ToProto(mapCtx, in.CreateTime)
	return out
}
func EvaluationJobAlertConfig_FromProto(mapCtx *direct.MapContext, in *pb.EvaluationJobAlertConfig) *krm.EvaluationJobAlertConfig {
	if in == nil {
		return nil
	}
	out := &krm.EvaluationJobAlertConfig{}
	out.Email = direct.LazyPtr(in.GetEmail())
	out.MinAcceptableMeanAveragePrecision = direct.LazyPtr(in.GetMinAcceptableMeanAveragePrecision())
	return out
}
func EvaluationJobAlertConfig_ToProto(mapCtx *direct.MapContext, in *krm.EvaluationJobAlertConfig) *pb.EvaluationJobAlertConfig {
	if in == nil {
		return nil
	}
	out := &pb.EvaluationJobAlertConfig{}
	out.Email = direct.ValueOf(in.Email)
	out.MinAcceptableMeanAveragePrecision = direct.ValueOf(in.MinAcceptableMeanAveragePrecision)
	return out
}
func EvaluationJobConfig_FromProto(mapCtx *direct.MapContext, in *pb.EvaluationJobConfig) *krm.EvaluationJobConfig {
	if in == nil {
		return nil
	}
	out := &krm.EvaluationJobConfig{}
	out.ImageClassificationConfig = ImageClassificationConfig_FromProto(mapCtx, in.GetImageClassificationConfig())
	out.BoundingPolyConfig = BoundingPolyConfig_FromProto(mapCtx, in.GetBoundingPolyConfig())
	out.TextClassificationConfig = TextClassificationConfig_FromProto(mapCtx, in.GetTextClassificationConfig())
	out.InputConfig = InputConfig_FromProto(mapCtx, in.GetInputConfig())
	out.EvaluationConfig = EvaluationConfig_FromProto(mapCtx, in.GetEvaluationConfig())
	out.HumanAnnotationConfig = HumanAnnotationConfig_FromProto(mapCtx, in.GetHumanAnnotationConfig())
	out.BigqueryImportKeys = in.BigqueryImportKeys
	out.ExampleCount = direct.LazyPtr(in.GetExampleCount())
	out.ExampleSamplePercentage = direct.LazyPtr(in.GetExampleSamplePercentage())
	out.EvaluationJobAlertConfig = EvaluationJobAlertConfig_FromProto(mapCtx, in.GetEvaluationJobAlertConfig())
	return out
}
func EvaluationJobConfig_ToProto(mapCtx *direct.MapContext, in *krm.EvaluationJobConfig) *pb.EvaluationJobConfig {
	if in == nil {
		return nil
	}
	out := &pb.EvaluationJobConfig{}
	if oneof := ImageClassificationConfig_ToProto(mapCtx, in.ImageClassificationConfig); oneof != nil {
		out.HumanAnnotationRequestConfig = &pb.EvaluationJobConfig_ImageClassificationConfig{ImageClassificationConfig: oneof}
	}
	if oneof := BoundingPolyConfig_ToProto(mapCtx, in.BoundingPolyConfig); oneof != nil {
		out.HumanAnnotationRequestConfig = &pb.EvaluationJobConfig_BoundingPolyConfig{BoundingPolyConfig: oneof}
	}
	if oneof := TextClassificationConfig_ToProto(mapCtx, in.TextClassificationConfig); oneof != nil {
		out.HumanAnnotationRequestConfig = &pb.EvaluationJobConfig_TextClassificationConfig{TextClassificationConfig: oneof}
	}
	out.InputConfig = InputConfig_ToProto(mapCtx, in.InputConfig)
	out.EvaluationConfig = EvaluationConfig_ToProto(mapCtx, in.EvaluationConfig)
	out.HumanAnnotationConfig = HumanAnnotationConfig_ToProto(mapCtx, in.HumanAnnotationConfig)
	out.BigqueryImportKeys = in.BigqueryImportKeys
	out.ExampleCount = direct.ValueOf(in.ExampleCount)
	out.ExampleSamplePercentage = direct.ValueOf(in.ExampleSamplePercentage)
	out.EvaluationJobAlertConfig = EvaluationJobAlertConfig_ToProto(mapCtx, in.EvaluationJobAlertConfig)
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
func HumanAnnotationConfig_FromProto(mapCtx *direct.MapContext, in *pb.HumanAnnotationConfig) *krm.HumanAnnotationConfig {
	if in == nil {
		return nil
	}
	out := &krm.HumanAnnotationConfig{}
	out.Instruction = direct.LazyPtr(in.GetInstruction())
	out.AnnotatedDatasetDisplayName = direct.LazyPtr(in.GetAnnotatedDatasetDisplayName())
	out.AnnotatedDatasetDescription = direct.LazyPtr(in.GetAnnotatedDatasetDescription())
	out.LabelGroup = direct.LazyPtr(in.GetLabelGroup())
	out.LanguageCode = direct.LazyPtr(in.GetLanguageCode())
	out.ReplicaCount = direct.LazyPtr(in.GetReplicaCount())
	out.QuestionDuration = direct.StringDuration_FromProto(mapCtx, in.GetQuestionDuration())
	out.ContributorEmails = in.ContributorEmails
	out.UserEmailAddress = direct.LazyPtr(in.GetUserEmailAddress())
	return out
}
func HumanAnnotationConfig_ToProto(mapCtx *direct.MapContext, in *krm.HumanAnnotationConfig) *pb.HumanAnnotationConfig {
	if in == nil {
		return nil
	}
	out := &pb.HumanAnnotationConfig{}
	out.Instruction = direct.ValueOf(in.Instruction)
	out.AnnotatedDatasetDisplayName = direct.ValueOf(in.AnnotatedDatasetDisplayName)
	out.AnnotatedDatasetDescription = direct.ValueOf(in.AnnotatedDatasetDescription)
	out.LabelGroup = direct.ValueOf(in.LabelGroup)
	out.LanguageCode = direct.ValueOf(in.LanguageCode)
	out.ReplicaCount = direct.ValueOf(in.ReplicaCount)
	out.QuestionDuration = direct.StringDuration_ToProto(mapCtx, in.QuestionDuration)
	out.ContributorEmails = in.ContributorEmails
	out.UserEmailAddress = direct.ValueOf(in.UserEmailAddress)
	return out
}
func ImageClassificationConfig_FromProto(mapCtx *direct.MapContext, in *pb.ImageClassificationConfig) *krm.ImageClassificationConfig {
	if in == nil {
		return nil
	}
	out := &krm.ImageClassificationConfig{}
	out.AnnotationSpecSet = direct.LazyPtr(in.GetAnnotationSpecSet())
	out.AllowMultiLabel = direct.LazyPtr(in.GetAllowMultiLabel())
	out.AnswerAggregationType = direct.Enum_FromProto(mapCtx, in.GetAnswerAggregationType())
	return out
}
func ImageClassificationConfig_ToProto(mapCtx *direct.MapContext, in *krm.ImageClassificationConfig) *pb.ImageClassificationConfig {
	if in == nil {
		return nil
	}
	out := &pb.ImageClassificationConfig{}
	out.AnnotationSpecSet = direct.ValueOf(in.AnnotationSpecSet)
	out.AllowMultiLabel = direct.ValueOf(in.AllowMultiLabel)
	out.AnswerAggregationType = direct.Enum_ToProto[pb.StringAggregationType](mapCtx, in.AnswerAggregationType)
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
func SentimentConfig_FromProto(mapCtx *direct.MapContext, in *pb.SentimentConfig) *krm.SentimentConfig {
	if in == nil {
		return nil
	}
	out := &krm.SentimentConfig{}
	out.EnableLabelSentimentSelection = direct.LazyPtr(in.GetEnableLabelSentimentSelection())
	return out
}
func SentimentConfig_ToProto(mapCtx *direct.MapContext, in *krm.SentimentConfig) *pb.SentimentConfig {
	if in == nil {
		return nil
	}
	out := &pb.SentimentConfig{}
	out.EnableLabelSentimentSelection = direct.ValueOf(in.EnableLabelSentimentSelection)
	return out
}
func TextClassificationConfig_FromProto(mapCtx *direct.MapContext, in *pb.TextClassificationConfig) *krm.TextClassificationConfig {
	if in == nil {
		return nil
	}
	out := &krm.TextClassificationConfig{}
	out.AllowMultiLabel = direct.LazyPtr(in.GetAllowMultiLabel())
	out.AnnotationSpecSet = direct.LazyPtr(in.GetAnnotationSpecSet())
	out.SentimentConfig = SentimentConfig_FromProto(mapCtx, in.GetSentimentConfig())
	return out
}
func TextClassificationConfig_ToProto(mapCtx *direct.MapContext, in *krm.TextClassificationConfig) *pb.TextClassificationConfig {
	if in == nil {
		return nil
	}
	out := &pb.TextClassificationConfig{}
	out.AllowMultiLabel = direct.ValueOf(in.AllowMultiLabel)
	out.AnnotationSpecSet = direct.ValueOf(in.AnnotationSpecSet)
	out.SentimentConfig = SentimentConfig_ToProto(mapCtx, in.SentimentConfig)
	return out
}
