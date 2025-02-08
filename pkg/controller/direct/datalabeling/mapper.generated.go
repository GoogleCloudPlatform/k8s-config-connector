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
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct"
	refs "github.com/GoogleCloudPlatform/k8s-config-connector/apis/refs/v1beta1"
	pb "cloud.google.com/go/datalabeling/apiv1beta1/datalabelingpb"
	krm "github.com/GoogleCloudPlatform/k8s-config-connector/apis/datalabeling/v1alpha1"
)
func AnnotatedDataset_FromProto(mapCtx *direct.MapContext, in *pb.AnnotatedDataset) *krm.AnnotatedDataset {
	if in == nil {
		return nil
	}
	out := &krm.AnnotatedDataset{}
	out.Name = direct.LazyPtr(in.GetName())
	out.DisplayName = direct.LazyPtr(in.GetDisplayName())
	out.Description = direct.LazyPtr(in.GetDescription())
	out.AnnotationSource = direct.Enum_FromProto(mapCtx, in.GetAnnotationSource())
	out.AnnotationType = direct.Enum_FromProto(mapCtx, in.GetAnnotationType())
	out.ExampleCount = direct.LazyPtr(in.GetExampleCount())
	out.CompletedExampleCount = direct.LazyPtr(in.GetCompletedExampleCount())
	out.LabelStats = LabelStats_FromProto(mapCtx, in.GetLabelStats())
	out.CreateTime = direct.StringTimestamp_FromProto(mapCtx, in.GetCreateTime())
	out.Metadata = AnnotatedDatasetMetadata_FromProto(mapCtx, in.GetMetadata())
	out.BlockingResources = in.BlockingResources
	return out
}
func AnnotatedDataset_ToProto(mapCtx *direct.MapContext, in *krm.AnnotatedDataset) *pb.AnnotatedDataset {
	if in == nil {
		return nil
	}
	out := &pb.AnnotatedDataset{}
	out.Name = direct.ValueOf(in.Name)
	out.DisplayName = direct.ValueOf(in.DisplayName)
	out.Description = direct.ValueOf(in.Description)
	out.AnnotationSource = direct.Enum_ToProto[pb.AnnotationSource](mapCtx, in.AnnotationSource)
	out.AnnotationType = direct.Enum_ToProto[pb.AnnotationType](mapCtx, in.AnnotationType)
	out.ExampleCount = direct.ValueOf(in.ExampleCount)
	out.CompletedExampleCount = direct.ValueOf(in.CompletedExampleCount)
	out.LabelStats = LabelStats_ToProto(mapCtx, in.LabelStats)
	out.CreateTime = direct.StringTimestamp_ToProto(mapCtx, in.CreateTime)
	out.Metadata = AnnotatedDatasetMetadata_ToProto(mapCtx, in.Metadata)
	out.BlockingResources = in.BlockingResources
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
func DatalabelingAnnotatedDatasetObservedState_FromProto(mapCtx *direct.MapContext, in *pb.AnnotatedDataset) *krm.DatalabelingAnnotatedDatasetObservedState {
	if in == nil {
		return nil
	}
	out := &krm.DatalabelingAnnotatedDatasetObservedState{}
	// MISSING: Name
	// MISSING: DisplayName
	// MISSING: Description
	// MISSING: AnnotationSource
	// MISSING: AnnotationType
	// MISSING: ExampleCount
	// MISSING: CompletedExampleCount
	// MISSING: LabelStats
	// MISSING: CreateTime
	// MISSING: Metadata
	// MISSING: BlockingResources
	return out
}
func DatalabelingAnnotatedDatasetObservedState_ToProto(mapCtx *direct.MapContext, in *krm.DatalabelingAnnotatedDatasetObservedState) *pb.AnnotatedDataset {
	if in == nil {
		return nil
	}
	out := &pb.AnnotatedDataset{}
	// MISSING: Name
	// MISSING: DisplayName
	// MISSING: Description
	// MISSING: AnnotationSource
	// MISSING: AnnotationType
	// MISSING: ExampleCount
	// MISSING: CompletedExampleCount
	// MISSING: LabelStats
	// MISSING: CreateTime
	// MISSING: Metadata
	// MISSING: BlockingResources
	return out
}
func DatalabelingAnnotatedDatasetSpec_FromProto(mapCtx *direct.MapContext, in *pb.AnnotatedDataset) *krm.DatalabelingAnnotatedDatasetSpec {
	if in == nil {
		return nil
	}
	out := &krm.DatalabelingAnnotatedDatasetSpec{}
	// MISSING: Name
	// MISSING: DisplayName
	// MISSING: Description
	// MISSING: AnnotationSource
	// MISSING: AnnotationType
	// MISSING: ExampleCount
	// MISSING: CompletedExampleCount
	// MISSING: LabelStats
	// MISSING: CreateTime
	// MISSING: Metadata
	// MISSING: BlockingResources
	return out
}
func DatalabelingAnnotatedDatasetSpec_ToProto(mapCtx *direct.MapContext, in *krm.DatalabelingAnnotatedDatasetSpec) *pb.AnnotatedDataset {
	if in == nil {
		return nil
	}
	out := &pb.AnnotatedDataset{}
	// MISSING: Name
	// MISSING: DisplayName
	// MISSING: Description
	// MISSING: AnnotationSource
	// MISSING: AnnotationType
	// MISSING: ExampleCount
	// MISSING: CompletedExampleCount
	// MISSING: LabelStats
	// MISSING: CreateTime
	// MISSING: Metadata
	// MISSING: BlockingResources
	return out
}
func EventConfig_FromProto(mapCtx *direct.MapContext, in *pb.EventConfig) *krm.EventConfig {
	if in == nil {
		return nil
	}
	out := &krm.EventConfig{}
	out.AnnotationSpecSets = in.AnnotationSpecSets
	return out
}
func EventConfig_ToProto(mapCtx *direct.MapContext, in *krm.EventConfig) *pb.EventConfig {
	if in == nil {
		return nil
	}
	out := &pb.EventConfig{}
	out.AnnotationSpecSets = in.AnnotationSpecSets
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
func LabelStats_FromProto(mapCtx *direct.MapContext, in *pb.LabelStats) *krm.LabelStats {
	if in == nil {
		return nil
	}
	out := &krm.LabelStats{}
	out.ExampleCount = in.ExampleCount
	return out
}
func LabelStats_ToProto(mapCtx *direct.MapContext, in *krm.LabelStats) *pb.LabelStats {
	if in == nil {
		return nil
	}
	out := &pb.LabelStats{}
	out.ExampleCount = in.ExampleCount
	return out
}
func ObjectDetectionConfig_FromProto(mapCtx *direct.MapContext, in *pb.ObjectDetectionConfig) *krm.ObjectDetectionConfig {
	if in == nil {
		return nil
	}
	out := &krm.ObjectDetectionConfig{}
	out.AnnotationSpecSet = direct.LazyPtr(in.GetAnnotationSpecSet())
	out.ExtractionFrameRate = direct.LazyPtr(in.GetExtractionFrameRate())
	return out
}
func ObjectDetectionConfig_ToProto(mapCtx *direct.MapContext, in *krm.ObjectDetectionConfig) *pb.ObjectDetectionConfig {
	if in == nil {
		return nil
	}
	out := &pb.ObjectDetectionConfig{}
	out.AnnotationSpecSet = direct.ValueOf(in.AnnotationSpecSet)
	out.ExtractionFrameRate = direct.ValueOf(in.ExtractionFrameRate)
	return out
}
func ObjectTrackingConfig_FromProto(mapCtx *direct.MapContext, in *pb.ObjectTrackingConfig) *krm.ObjectTrackingConfig {
	if in == nil {
		return nil
	}
	out := &krm.ObjectTrackingConfig{}
	out.AnnotationSpecSet = direct.LazyPtr(in.GetAnnotationSpecSet())
	return out
}
func ObjectTrackingConfig_ToProto(mapCtx *direct.MapContext, in *krm.ObjectTrackingConfig) *pb.ObjectTrackingConfig {
	if in == nil {
		return nil
	}
	out := &pb.ObjectTrackingConfig{}
	out.AnnotationSpecSet = direct.ValueOf(in.AnnotationSpecSet)
	return out
}
func PolylineConfig_FromProto(mapCtx *direct.MapContext, in *pb.PolylineConfig) *krm.PolylineConfig {
	if in == nil {
		return nil
	}
	out := &krm.PolylineConfig{}
	out.AnnotationSpecSet = direct.LazyPtr(in.GetAnnotationSpecSet())
	out.InstructionMessage = direct.LazyPtr(in.GetInstructionMessage())
	return out
}
func PolylineConfig_ToProto(mapCtx *direct.MapContext, in *krm.PolylineConfig) *pb.PolylineConfig {
	if in == nil {
		return nil
	}
	out := &pb.PolylineConfig{}
	out.AnnotationSpecSet = direct.ValueOf(in.AnnotationSpecSet)
	out.InstructionMessage = direct.ValueOf(in.InstructionMessage)
	return out
}
func SegmentationConfig_FromProto(mapCtx *direct.MapContext, in *pb.SegmentationConfig) *krm.SegmentationConfig {
	if in == nil {
		return nil
	}
	out := &krm.SegmentationConfig{}
	out.AnnotationSpecSet = direct.LazyPtr(in.GetAnnotationSpecSet())
	out.InstructionMessage = direct.LazyPtr(in.GetInstructionMessage())
	return out
}
func SegmentationConfig_ToProto(mapCtx *direct.MapContext, in *krm.SegmentationConfig) *pb.SegmentationConfig {
	if in == nil {
		return nil
	}
	out := &pb.SegmentationConfig{}
	out.AnnotationSpecSet = direct.ValueOf(in.AnnotationSpecSet)
	out.InstructionMessage = direct.ValueOf(in.InstructionMessage)
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
func TextEntityExtractionConfig_FromProto(mapCtx *direct.MapContext, in *pb.TextEntityExtractionConfig) *krm.TextEntityExtractionConfig {
	if in == nil {
		return nil
	}
	out := &krm.TextEntityExtractionConfig{}
	out.AnnotationSpecSet = direct.LazyPtr(in.GetAnnotationSpecSet())
	return out
}
func TextEntityExtractionConfig_ToProto(mapCtx *direct.MapContext, in *krm.TextEntityExtractionConfig) *pb.TextEntityExtractionConfig {
	if in == nil {
		return nil
	}
	out := &pb.TextEntityExtractionConfig{}
	out.AnnotationSpecSet = direct.ValueOf(in.AnnotationSpecSet)
	return out
}
func VideoClassificationConfig_FromProto(mapCtx *direct.MapContext, in *pb.VideoClassificationConfig) *krm.VideoClassificationConfig {
	if in == nil {
		return nil
	}
	out := &krm.VideoClassificationConfig{}
	out.AnnotationSpecSetConfigs = direct.Slice_FromProto(mapCtx, in.AnnotationSpecSetConfigs, VideoClassificationConfig_AnnotationSpecSetConfig_FromProto)
	out.ApplyShotDetection = direct.LazyPtr(in.GetApplyShotDetection())
	return out
}
func VideoClassificationConfig_ToProto(mapCtx *direct.MapContext, in *krm.VideoClassificationConfig) *pb.VideoClassificationConfig {
	if in == nil {
		return nil
	}
	out := &pb.VideoClassificationConfig{}
	out.AnnotationSpecSetConfigs = direct.Slice_ToProto(mapCtx, in.AnnotationSpecSetConfigs, VideoClassificationConfig_AnnotationSpecSetConfig_ToProto)
	out.ApplyShotDetection = direct.ValueOf(in.ApplyShotDetection)
	return out
}
func VideoClassificationConfig_AnnotationSpecSetConfig_FromProto(mapCtx *direct.MapContext, in *pb.VideoClassificationConfig_AnnotationSpecSetConfig) *krm.VideoClassificationConfig_AnnotationSpecSetConfig {
	if in == nil {
		return nil
	}
	out := &krm.VideoClassificationConfig_AnnotationSpecSetConfig{}
	out.AnnotationSpecSet = direct.LazyPtr(in.GetAnnotationSpecSet())
	out.AllowMultiLabel = direct.LazyPtr(in.GetAllowMultiLabel())
	return out
}
func VideoClassificationConfig_AnnotationSpecSetConfig_ToProto(mapCtx *direct.MapContext, in *krm.VideoClassificationConfig_AnnotationSpecSetConfig) *pb.VideoClassificationConfig_AnnotationSpecSetConfig {
	if in == nil {
		return nil
	}
	out := &pb.VideoClassificationConfig_AnnotationSpecSetConfig{}
	out.AnnotationSpecSet = direct.ValueOf(in.AnnotationSpecSet)
	out.AllowMultiLabel = direct.ValueOf(in.AllowMultiLabel)
	return out
}
