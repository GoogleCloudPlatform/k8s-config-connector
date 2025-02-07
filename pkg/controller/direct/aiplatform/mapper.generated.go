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

package aiplatform

import (
	refs "github.com/GoogleCloudPlatform/k8s-config-connector/apis/refs/v1beta1"
	pb "cloud.google.com/go/aiplatform/apiv1/aiplatformpb"
	krm "github.com/GoogleCloudPlatform/k8s-config-connector/apis/aiplatform/v1alpha1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct"
)
func ActiveLearningConfig_FromProto(mapCtx *direct.MapContext, in *pb.ActiveLearningConfig) *krm.ActiveLearningConfig {
	if in == nil {
		return nil
	}
	out := &krm.ActiveLearningConfig{}
	out.MaxDataItemCount = direct.LazyPtr(in.GetMaxDataItemCount())
	out.MaxDataItemPercentage = direct.LazyPtr(in.GetMaxDataItemPercentage())
	out.SampleConfig = SampleConfig_FromProto(mapCtx, in.GetSampleConfig())
	out.TrainingConfig = TrainingConfig_FromProto(mapCtx, in.GetTrainingConfig())
	return out
}
func ActiveLearningConfig_ToProto(mapCtx *direct.MapContext, in *krm.ActiveLearningConfig) *pb.ActiveLearningConfig {
	if in == nil {
		return nil
	}
	out := &pb.ActiveLearningConfig{}
	if oneof := ActiveLearningConfig_MaxDataItemCount_ToProto(mapCtx, in.MaxDataItemCount); oneof != nil {
		out.HumanLabelingBudget = oneof
	}
	if oneof := ActiveLearningConfig_MaxDataItemPercentage_ToProto(mapCtx, in.MaxDataItemPercentage); oneof != nil {
		out.HumanLabelingBudget = oneof
	}
	out.SampleConfig = SampleConfig_ToProto(mapCtx, in.SampleConfig)
	out.TrainingConfig = TrainingConfig_ToProto(mapCtx, in.TrainingConfig)
	return out
}
func AiplatformAnnotationObservedState_FromProto(mapCtx *direct.MapContext, in *pb.Annotation) *krm.AiplatformAnnotationObservedState {
	if in == nil {
		return nil
	}
	out := &krm.AiplatformAnnotationObservedState{}
	// MISSING: Name
	// MISSING: PayloadSchemaURI
	// MISSING: Payload
	// MISSING: CreateTime
	// MISSING: UpdateTime
	// MISSING: Etag
	// MISSING: AnnotationSource
	// MISSING: Labels
	return out
}
func AiplatformAnnotationObservedState_ToProto(mapCtx *direct.MapContext, in *krm.AiplatformAnnotationObservedState) *pb.Annotation {
	if in == nil {
		return nil
	}
	out := &pb.Annotation{}
	// MISSING: Name
	// MISSING: PayloadSchemaURI
	// MISSING: Payload
	// MISSING: CreateTime
	// MISSING: UpdateTime
	// MISSING: Etag
	// MISSING: AnnotationSource
	// MISSING: Labels
	return out
}
func AiplatformAnnotationSpec_FromProto(mapCtx *direct.MapContext, in *pb.Annotation) *krm.AiplatformAnnotationSpec {
	if in == nil {
		return nil
	}
	out := &krm.AiplatformAnnotationSpec{}
	// MISSING: Name
	// MISSING: PayloadSchemaURI
	// MISSING: Payload
	// MISSING: CreateTime
	// MISSING: UpdateTime
	// MISSING: Etag
	// MISSING: AnnotationSource
	// MISSING: Labels
	return out
}
func AiplatformAnnotationSpec_ToProto(mapCtx *direct.MapContext, in *krm.AiplatformAnnotationSpec) *pb.Annotation {
	if in == nil {
		return nil
	}
	out := &pb.Annotation{}
	// MISSING: Name
	// MISSING: PayloadSchemaURI
	// MISSING: Payload
	// MISSING: CreateTime
	// MISSING: UpdateTime
	// MISSING: Etag
	// MISSING: AnnotationSource
	// MISSING: Labels
	return out
}
func AiplatformAnnotationSpecObservedState_FromProto(mapCtx *direct.MapContext, in *pb.AnnotationSpec) *krm.AiplatformAnnotationSpecObservedState {
	if in == nil {
		return nil
	}
	out := &krm.AiplatformAnnotationSpecObservedState{}
	// MISSING: Name
	// MISSING: DisplayName
	// MISSING: CreateTime
	// MISSING: UpdateTime
	// MISSING: Etag
	return out
}
func AiplatformAnnotationSpecObservedState_ToProto(mapCtx *direct.MapContext, in *krm.AiplatformAnnotationSpecObservedState) *pb.AnnotationSpec {
	if in == nil {
		return nil
	}
	out := &pb.AnnotationSpec{}
	// MISSING: Name
	// MISSING: DisplayName
	// MISSING: CreateTime
	// MISSING: UpdateTime
	// MISSING: Etag
	return out
}
func AiplatformAnnotationSpecSpec_FromProto(mapCtx *direct.MapContext, in *pb.AnnotationSpec) *krm.AiplatformAnnotationSpecSpec {
	if in == nil {
		return nil
	}
	out := &krm.AiplatformAnnotationSpecSpec{}
	// MISSING: Name
	// MISSING: DisplayName
	// MISSING: CreateTime
	// MISSING: UpdateTime
	// MISSING: Etag
	return out
}
func AiplatformAnnotationSpecSpec_ToProto(mapCtx *direct.MapContext, in *krm.AiplatformAnnotationSpecSpec) *pb.AnnotationSpec {
	if in == nil {
		return nil
	}
	out := &pb.AnnotationSpec{}
	// MISSING: Name
	// MISSING: DisplayName
	// MISSING: CreateTime
	// MISSING: UpdateTime
	// MISSING: Etag
	return out
}
func AiplatformArtifactObservedState_FromProto(mapCtx *direct.MapContext, in *pb.Artifact) *krm.AiplatformArtifactObservedState {
	if in == nil {
		return nil
	}
	out := &krm.AiplatformArtifactObservedState{}
	// MISSING: Name
	// MISSING: DisplayName
	// MISSING: URI
	// MISSING: Etag
	// MISSING: Labels
	// MISSING: CreateTime
	// MISSING: UpdateTime
	// MISSING: State
	// MISSING: SchemaTitle
	// MISSING: SchemaVersion
	// MISSING: Metadata
	// MISSING: Description
	return out
}
func AiplatformArtifactObservedState_ToProto(mapCtx *direct.MapContext, in *krm.AiplatformArtifactObservedState) *pb.Artifact {
	if in == nil {
		return nil
	}
	out := &pb.Artifact{}
	// MISSING: Name
	// MISSING: DisplayName
	// MISSING: URI
	// MISSING: Etag
	// MISSING: Labels
	// MISSING: CreateTime
	// MISSING: UpdateTime
	// MISSING: State
	// MISSING: SchemaTitle
	// MISSING: SchemaVersion
	// MISSING: Metadata
	// MISSING: Description
	return out
}
func AiplatformArtifactSpec_FromProto(mapCtx *direct.MapContext, in *pb.Artifact) *krm.AiplatformArtifactSpec {
	if in == nil {
		return nil
	}
	out := &krm.AiplatformArtifactSpec{}
	// MISSING: Name
	// MISSING: DisplayName
	// MISSING: URI
	// MISSING: Etag
	// MISSING: Labels
	// MISSING: CreateTime
	// MISSING: UpdateTime
	// MISSING: State
	// MISSING: SchemaTitle
	// MISSING: SchemaVersion
	// MISSING: Metadata
	// MISSING: Description
	return out
}
func AiplatformArtifactSpec_ToProto(mapCtx *direct.MapContext, in *krm.AiplatformArtifactSpec) *pb.Artifact {
	if in == nil {
		return nil
	}
	out := &pb.Artifact{}
	// MISSING: Name
	// MISSING: DisplayName
	// MISSING: URI
	// MISSING: Etag
	// MISSING: Labels
	// MISSING: CreateTime
	// MISSING: UpdateTime
	// MISSING: State
	// MISSING: SchemaTitle
	// MISSING: SchemaVersion
	// MISSING: Metadata
	// MISSING: Description
	return out
}
func AiplatformBatchPredictionJobObservedState_FromProto(mapCtx *direct.MapContext, in *pb.BatchPredictionJob) *krm.AiplatformBatchPredictionJobObservedState {
	if in == nil {
		return nil
	}
	out := &krm.AiplatformBatchPredictionJobObservedState{}
	// MISSING: Name
	// MISSING: DisplayName
	// MISSING: Model
	// MISSING: ModelVersionID
	// MISSING: UnmanagedContainerModel
	// MISSING: InputConfig
	// MISSING: InstanceConfig
	// MISSING: ModelParameters
	// MISSING: OutputConfig
	// MISSING: DedicatedResources
	// MISSING: ServiceAccount
	// MISSING: ManualBatchTuningParameters
	// MISSING: GenerateExplanation
	// MISSING: ExplanationSpec
	// MISSING: OutputInfo
	// MISSING: State
	// MISSING: Error
	// MISSING: PartialFailures
	// MISSING: ResourcesConsumed
	// MISSING: CompletionStats
	// MISSING: CreateTime
	// MISSING: StartTime
	// MISSING: EndTime
	// MISSING: UpdateTime
	// MISSING: Labels
	// MISSING: EncryptionSpec
	// MISSING: DisableContainerLogging
	// MISSING: SatisfiesPzs
	// MISSING: SatisfiesPzi
	return out
}
func AiplatformBatchPredictionJobObservedState_ToProto(mapCtx *direct.MapContext, in *krm.AiplatformBatchPredictionJobObservedState) *pb.BatchPredictionJob {
	if in == nil {
		return nil
	}
	out := &pb.BatchPredictionJob{}
	// MISSING: Name
	// MISSING: DisplayName
	// MISSING: Model
	// MISSING: ModelVersionID
	// MISSING: UnmanagedContainerModel
	// MISSING: InputConfig
	// MISSING: InstanceConfig
	// MISSING: ModelParameters
	// MISSING: OutputConfig
	// MISSING: DedicatedResources
	// MISSING: ServiceAccount
	// MISSING: ManualBatchTuningParameters
	// MISSING: GenerateExplanation
	// MISSING: ExplanationSpec
	// MISSING: OutputInfo
	// MISSING: State
	// MISSING: Error
	// MISSING: PartialFailures
	// MISSING: ResourcesConsumed
	// MISSING: CompletionStats
	// MISSING: CreateTime
	// MISSING: StartTime
	// MISSING: EndTime
	// MISSING: UpdateTime
	// MISSING: Labels
	// MISSING: EncryptionSpec
	// MISSING: DisableContainerLogging
	// MISSING: SatisfiesPzs
	// MISSING: SatisfiesPzi
	return out
}
func AiplatformBatchPredictionJobSpec_FromProto(mapCtx *direct.MapContext, in *pb.BatchPredictionJob) *krm.AiplatformBatchPredictionJobSpec {
	if in == nil {
		return nil
	}
	out := &krm.AiplatformBatchPredictionJobSpec{}
	// MISSING: Name
	// MISSING: DisplayName
	// MISSING: Model
	// MISSING: ModelVersionID
	// MISSING: UnmanagedContainerModel
	// MISSING: InputConfig
	// MISSING: InstanceConfig
	// MISSING: ModelParameters
	// MISSING: OutputConfig
	// MISSING: DedicatedResources
	// MISSING: ServiceAccount
	// MISSING: ManualBatchTuningParameters
	// MISSING: GenerateExplanation
	// MISSING: ExplanationSpec
	// MISSING: OutputInfo
	// MISSING: State
	// MISSING: Error
	// MISSING: PartialFailures
	// MISSING: ResourcesConsumed
	// MISSING: CompletionStats
	// MISSING: CreateTime
	// MISSING: StartTime
	// MISSING: EndTime
	// MISSING: UpdateTime
	// MISSING: Labels
	// MISSING: EncryptionSpec
	// MISSING: DisableContainerLogging
	// MISSING: SatisfiesPzs
	// MISSING: SatisfiesPzi
	return out
}
func AiplatformBatchPredictionJobSpec_ToProto(mapCtx *direct.MapContext, in *krm.AiplatformBatchPredictionJobSpec) *pb.BatchPredictionJob {
	if in == nil {
		return nil
	}
	out := &pb.BatchPredictionJob{}
	// MISSING: Name
	// MISSING: DisplayName
	// MISSING: Model
	// MISSING: ModelVersionID
	// MISSING: UnmanagedContainerModel
	// MISSING: InputConfig
	// MISSING: InstanceConfig
	// MISSING: ModelParameters
	// MISSING: OutputConfig
	// MISSING: DedicatedResources
	// MISSING: ServiceAccount
	// MISSING: ManualBatchTuningParameters
	// MISSING: GenerateExplanation
	// MISSING: ExplanationSpec
	// MISSING: OutputInfo
	// MISSING: State
	// MISSING: Error
	// MISSING: PartialFailures
	// MISSING: ResourcesConsumed
	// MISSING: CompletionStats
	// MISSING: CreateTime
	// MISSING: StartTime
	// MISSING: EndTime
	// MISSING: UpdateTime
	// MISSING: Labels
	// MISSING: EncryptionSpec
	// MISSING: DisableContainerLogging
	// MISSING: SatisfiesPzs
	// MISSING: SatisfiesPzi
	return out
}
func AiplatformCachedContentObservedState_FromProto(mapCtx *direct.MapContext, in *pb.CachedContent) *krm.AiplatformCachedContentObservedState {
	if in == nil {
		return nil
	}
	out := &krm.AiplatformCachedContentObservedState{}
	// MISSING: ExpireTime
	// MISSING: Ttl
	// MISSING: Name
	// MISSING: DisplayName
	// MISSING: Model
	// MISSING: SystemInstruction
	// MISSING: Contents
	// MISSING: Tools
	// MISSING: ToolConfig
	// MISSING: CreateTime
	// MISSING: UpdateTime
	// MISSING: UsageMetadata
	return out
}
func AiplatformCachedContentObservedState_ToProto(mapCtx *direct.MapContext, in *krm.AiplatformCachedContentObservedState) *pb.CachedContent {
	if in == nil {
		return nil
	}
	out := &pb.CachedContent{}
	// MISSING: ExpireTime
	// MISSING: Ttl
	// MISSING: Name
	// MISSING: DisplayName
	// MISSING: Model
	// MISSING: SystemInstruction
	// MISSING: Contents
	// MISSING: Tools
	// MISSING: ToolConfig
	// MISSING: CreateTime
	// MISSING: UpdateTime
	// MISSING: UsageMetadata
	return out
}
func AiplatformCachedContentSpec_FromProto(mapCtx *direct.MapContext, in *pb.CachedContent) *krm.AiplatformCachedContentSpec {
	if in == nil {
		return nil
	}
	out := &krm.AiplatformCachedContentSpec{}
	// MISSING: ExpireTime
	// MISSING: Ttl
	// MISSING: Name
	// MISSING: DisplayName
	// MISSING: Model
	// MISSING: SystemInstruction
	// MISSING: Contents
	// MISSING: Tools
	// MISSING: ToolConfig
	// MISSING: CreateTime
	// MISSING: UpdateTime
	// MISSING: UsageMetadata
	return out
}
func AiplatformCachedContentSpec_ToProto(mapCtx *direct.MapContext, in *krm.AiplatformCachedContentSpec) *pb.CachedContent {
	if in == nil {
		return nil
	}
	out := &pb.CachedContent{}
	// MISSING: ExpireTime
	// MISSING: Ttl
	// MISSING: Name
	// MISSING: DisplayName
	// MISSING: Model
	// MISSING: SystemInstruction
	// MISSING: Contents
	// MISSING: Tools
	// MISSING: ToolConfig
	// MISSING: CreateTime
	// MISSING: UpdateTime
	// MISSING: UsageMetadata
	return out
}
func AiplatformContextObservedState_FromProto(mapCtx *direct.MapContext, in *pb.Context) *krm.AiplatformContextObservedState {
	if in == nil {
		return nil
	}
	out := &krm.AiplatformContextObservedState{}
	// MISSING: Name
	// MISSING: DisplayName
	// MISSING: Etag
	// MISSING: Labels
	// MISSING: CreateTime
	// MISSING: UpdateTime
	// MISSING: ParentContexts
	// MISSING: SchemaTitle
	// MISSING: SchemaVersion
	// MISSING: Metadata
	// MISSING: Description
	return out
}
func AiplatformContextObservedState_ToProto(mapCtx *direct.MapContext, in *krm.AiplatformContextObservedState) *pb.Context {
	if in == nil {
		return nil
	}
	out := &pb.Context{}
	// MISSING: Name
	// MISSING: DisplayName
	// MISSING: Etag
	// MISSING: Labels
	// MISSING: CreateTime
	// MISSING: UpdateTime
	// MISSING: ParentContexts
	// MISSING: SchemaTitle
	// MISSING: SchemaVersion
	// MISSING: Metadata
	// MISSING: Description
	return out
}
func AiplatformContextSpec_FromProto(mapCtx *direct.MapContext, in *pb.Context) *krm.AiplatformContextSpec {
	if in == nil {
		return nil
	}
	out := &krm.AiplatformContextSpec{}
	// MISSING: Name
	// MISSING: DisplayName
	// MISSING: Etag
	// MISSING: Labels
	// MISSING: CreateTime
	// MISSING: UpdateTime
	// MISSING: ParentContexts
	// MISSING: SchemaTitle
	// MISSING: SchemaVersion
	// MISSING: Metadata
	// MISSING: Description
	return out
}
func AiplatformContextSpec_ToProto(mapCtx *direct.MapContext, in *krm.AiplatformContextSpec) *pb.Context {
	if in == nil {
		return nil
	}
	out := &pb.Context{}
	// MISSING: Name
	// MISSING: DisplayName
	// MISSING: Etag
	// MISSING: Labels
	// MISSING: CreateTime
	// MISSING: UpdateTime
	// MISSING: ParentContexts
	// MISSING: SchemaTitle
	// MISSING: SchemaVersion
	// MISSING: Metadata
	// MISSING: Description
	return out
}
func AiplatformCustomJobObservedState_FromProto(mapCtx *direct.MapContext, in *pb.CustomJob) *krm.AiplatformCustomJobObservedState {
	if in == nil {
		return nil
	}
	out := &krm.AiplatformCustomJobObservedState{}
	// MISSING: Name
	// MISSING: DisplayName
	// MISSING: JobSpec
	// MISSING: State
	// MISSING: CreateTime
	// MISSING: StartTime
	// MISSING: EndTime
	// MISSING: UpdateTime
	// MISSING: Error
	// MISSING: Labels
	// MISSING: EncryptionSpec
	// MISSING: WebAccessUris
	// MISSING: SatisfiesPzs
	// MISSING: SatisfiesPzi
	return out
}
func AiplatformCustomJobObservedState_ToProto(mapCtx *direct.MapContext, in *krm.AiplatformCustomJobObservedState) *pb.CustomJob {
	if in == nil {
		return nil
	}
	out := &pb.CustomJob{}
	// MISSING: Name
	// MISSING: DisplayName
	// MISSING: JobSpec
	// MISSING: State
	// MISSING: CreateTime
	// MISSING: StartTime
	// MISSING: EndTime
	// MISSING: UpdateTime
	// MISSING: Error
	// MISSING: Labels
	// MISSING: EncryptionSpec
	// MISSING: WebAccessUris
	// MISSING: SatisfiesPzs
	// MISSING: SatisfiesPzi
	return out
}
func AiplatformCustomJobSpec_FromProto(mapCtx *direct.MapContext, in *pb.CustomJob) *krm.AiplatformCustomJobSpec {
	if in == nil {
		return nil
	}
	out := &krm.AiplatformCustomJobSpec{}
	// MISSING: Name
	// MISSING: DisplayName
	// MISSING: JobSpec
	// MISSING: State
	// MISSING: CreateTime
	// MISSING: StartTime
	// MISSING: EndTime
	// MISSING: UpdateTime
	// MISSING: Error
	// MISSING: Labels
	// MISSING: EncryptionSpec
	// MISSING: WebAccessUris
	// MISSING: SatisfiesPzs
	// MISSING: SatisfiesPzi
	return out
}
func AiplatformCustomJobSpec_ToProto(mapCtx *direct.MapContext, in *krm.AiplatformCustomJobSpec) *pb.CustomJob {
	if in == nil {
		return nil
	}
	out := &pb.CustomJob{}
	// MISSING: Name
	// MISSING: DisplayName
	// MISSING: JobSpec
	// MISSING: State
	// MISSING: CreateTime
	// MISSING: StartTime
	// MISSING: EndTime
	// MISSING: UpdateTime
	// MISSING: Error
	// MISSING: Labels
	// MISSING: EncryptionSpec
	// MISSING: WebAccessUris
	// MISSING: SatisfiesPzs
	// MISSING: SatisfiesPzi
	return out
}
func AiplatformDataItemObservedState_FromProto(mapCtx *direct.MapContext, in *pb.DataItem) *krm.AiplatformDataItemObservedState {
	if in == nil {
		return nil
	}
	out := &krm.AiplatformDataItemObservedState{}
	// MISSING: Name
	// MISSING: CreateTime
	// MISSING: UpdateTime
	// MISSING: Labels
	// MISSING: Payload
	// MISSING: Etag
	// MISSING: SatisfiesPzs
	// MISSING: SatisfiesPzi
	return out
}
func AiplatformDataItemObservedState_ToProto(mapCtx *direct.MapContext, in *krm.AiplatformDataItemObservedState) *pb.DataItem {
	if in == nil {
		return nil
	}
	out := &pb.DataItem{}
	// MISSING: Name
	// MISSING: CreateTime
	// MISSING: UpdateTime
	// MISSING: Labels
	// MISSING: Payload
	// MISSING: Etag
	// MISSING: SatisfiesPzs
	// MISSING: SatisfiesPzi
	return out
}
func AiplatformDataItemSpec_FromProto(mapCtx *direct.MapContext, in *pb.DataItem) *krm.AiplatformDataItemSpec {
	if in == nil {
		return nil
	}
	out := &krm.AiplatformDataItemSpec{}
	// MISSING: Name
	// MISSING: CreateTime
	// MISSING: UpdateTime
	// MISSING: Labels
	// MISSING: Payload
	// MISSING: Etag
	// MISSING: SatisfiesPzs
	// MISSING: SatisfiesPzi
	return out
}
func AiplatformDataItemSpec_ToProto(mapCtx *direct.MapContext, in *krm.AiplatformDataItemSpec) *pb.DataItem {
	if in == nil {
		return nil
	}
	out := &pb.DataItem{}
	// MISSING: Name
	// MISSING: CreateTime
	// MISSING: UpdateTime
	// MISSING: Labels
	// MISSING: Payload
	// MISSING: Etag
	// MISSING: SatisfiesPzs
	// MISSING: SatisfiesPzi
	return out
}
func AiplatformDataLabelingJobObservedState_FromProto(mapCtx *direct.MapContext, in *pb.DataLabelingJob) *krm.AiplatformDataLabelingJobObservedState {
	if in == nil {
		return nil
	}
	out := &krm.AiplatformDataLabelingJobObservedState{}
	// MISSING: Name
	// MISSING: DisplayName
	// MISSING: Datasets
	// MISSING: AnnotationLabels
	// MISSING: LabelerCount
	// MISSING: InstructionURI
	// MISSING: InputsSchemaURI
	// MISSING: Inputs
	// MISSING: State
	// MISSING: LabelingProgress
	// MISSING: CurrentSpend
	// MISSING: CreateTime
	// MISSING: UpdateTime
	// MISSING: Error
	// MISSING: Labels
	// MISSING: SpecialistPools
	// MISSING: EncryptionSpec
	// MISSING: ActiveLearningConfig
	return out
}
func AiplatformDataLabelingJobObservedState_ToProto(mapCtx *direct.MapContext, in *krm.AiplatformDataLabelingJobObservedState) *pb.DataLabelingJob {
	if in == nil {
		return nil
	}
	out := &pb.DataLabelingJob{}
	// MISSING: Name
	// MISSING: DisplayName
	// MISSING: Datasets
	// MISSING: AnnotationLabels
	// MISSING: LabelerCount
	// MISSING: InstructionURI
	// MISSING: InputsSchemaURI
	// MISSING: Inputs
	// MISSING: State
	// MISSING: LabelingProgress
	// MISSING: CurrentSpend
	// MISSING: CreateTime
	// MISSING: UpdateTime
	// MISSING: Error
	// MISSING: Labels
	// MISSING: SpecialistPools
	// MISSING: EncryptionSpec
	// MISSING: ActiveLearningConfig
	return out
}
func AiplatformDataLabelingJobSpec_FromProto(mapCtx *direct.MapContext, in *pb.DataLabelingJob) *krm.AiplatformDataLabelingJobSpec {
	if in == nil {
		return nil
	}
	out := &krm.AiplatformDataLabelingJobSpec{}
	// MISSING: Name
	// MISSING: DisplayName
	// MISSING: Datasets
	// MISSING: AnnotationLabels
	// MISSING: LabelerCount
	// MISSING: InstructionURI
	// MISSING: InputsSchemaURI
	// MISSING: Inputs
	// MISSING: State
	// MISSING: LabelingProgress
	// MISSING: CurrentSpend
	// MISSING: CreateTime
	// MISSING: UpdateTime
	// MISSING: Error
	// MISSING: Labels
	// MISSING: SpecialistPools
	// MISSING: EncryptionSpec
	// MISSING: ActiveLearningConfig
	return out
}
func AiplatformDataLabelingJobSpec_ToProto(mapCtx *direct.MapContext, in *krm.AiplatformDataLabelingJobSpec) *pb.DataLabelingJob {
	if in == nil {
		return nil
	}
	out := &pb.DataLabelingJob{}
	// MISSING: Name
	// MISSING: DisplayName
	// MISSING: Datasets
	// MISSING: AnnotationLabels
	// MISSING: LabelerCount
	// MISSING: InstructionURI
	// MISSING: InputsSchemaURI
	// MISSING: Inputs
	// MISSING: State
	// MISSING: LabelingProgress
	// MISSING: CurrentSpend
	// MISSING: CreateTime
	// MISSING: UpdateTime
	// MISSING: Error
	// MISSING: Labels
	// MISSING: SpecialistPools
	// MISSING: EncryptionSpec
	// MISSING: ActiveLearningConfig
	return out
}
func DataLabelingJob_FromProto(mapCtx *direct.MapContext, in *pb.DataLabelingJob) *krm.DataLabelingJob {
	if in == nil {
		return nil
	}
	out := &krm.DataLabelingJob{}
	// MISSING: Name
	out.DisplayName = direct.LazyPtr(in.GetDisplayName())
	out.Datasets = in.Datasets
	out.AnnotationLabels = in.AnnotationLabels
	out.LabelerCount = direct.LazyPtr(in.GetLabelerCount())
	out.InstructionURI = direct.LazyPtr(in.GetInstructionUri())
	out.InputsSchemaURI = direct.LazyPtr(in.GetInputsSchemaUri())
	out.Inputs = Value_FromProto(mapCtx, in.GetInputs())
	// MISSING: State
	// MISSING: LabelingProgress
	// MISSING: CurrentSpend
	// MISSING: CreateTime
	// MISSING: UpdateTime
	// MISSING: Error
	out.Labels = in.Labels
	out.SpecialistPools = in.SpecialistPools
	out.EncryptionSpec = EncryptionSpec_FromProto(mapCtx, in.GetEncryptionSpec())
	out.ActiveLearningConfig = ActiveLearningConfig_FromProto(mapCtx, in.GetActiveLearningConfig())
	return out
}
func DataLabelingJob_ToProto(mapCtx *direct.MapContext, in *krm.DataLabelingJob) *pb.DataLabelingJob {
	if in == nil {
		return nil
	}
	out := &pb.DataLabelingJob{}
	// MISSING: Name
	out.DisplayName = direct.ValueOf(in.DisplayName)
	out.Datasets = in.Datasets
	out.AnnotationLabels = in.AnnotationLabels
	out.LabelerCount = direct.ValueOf(in.LabelerCount)
	out.InstructionUri = direct.ValueOf(in.InstructionURI)
	out.InputsSchemaUri = direct.ValueOf(in.InputsSchemaURI)
	out.Inputs = Value_ToProto(mapCtx, in.Inputs)
	// MISSING: State
	// MISSING: LabelingProgress
	// MISSING: CurrentSpend
	// MISSING: CreateTime
	// MISSING: UpdateTime
	// MISSING: Error
	out.Labels = in.Labels
	out.SpecialistPools = in.SpecialistPools
	out.EncryptionSpec = EncryptionSpec_ToProto(mapCtx, in.EncryptionSpec)
	out.ActiveLearningConfig = ActiveLearningConfig_ToProto(mapCtx, in.ActiveLearningConfig)
	return out
}
func DataLabelingJobObservedState_FromProto(mapCtx *direct.MapContext, in *pb.DataLabelingJob) *krm.DataLabelingJobObservedState {
	if in == nil {
		return nil
	}
	out := &krm.DataLabelingJobObservedState{}
	out.Name = direct.LazyPtr(in.GetName())
	// MISSING: DisplayName
	// MISSING: Datasets
	// MISSING: AnnotationLabels
	// MISSING: LabelerCount
	// MISSING: InstructionURI
	// MISSING: InputsSchemaURI
	// MISSING: Inputs
	out.State = direct.Enum_FromProto(mapCtx, in.GetState())
	out.LabelingProgress = direct.LazyPtr(in.GetLabelingProgress())
	out.CurrentSpend = Money_FromProto(mapCtx, in.GetCurrentSpend())
	out.CreateTime = direct.StringTimestamp_FromProto(mapCtx, in.GetCreateTime())
	out.UpdateTime = direct.StringTimestamp_FromProto(mapCtx, in.GetUpdateTime())
	out.Error = Status_FromProto(mapCtx, in.GetError())
	// MISSING: Labels
	// MISSING: SpecialistPools
	// MISSING: EncryptionSpec
	// MISSING: ActiveLearningConfig
	return out
}
func DataLabelingJobObservedState_ToProto(mapCtx *direct.MapContext, in *krm.DataLabelingJobObservedState) *pb.DataLabelingJob {
	if in == nil {
		return nil
	}
	out := &pb.DataLabelingJob{}
	out.Name = direct.ValueOf(in.Name)
	// MISSING: DisplayName
	// MISSING: Datasets
	// MISSING: AnnotationLabels
	// MISSING: LabelerCount
	// MISSING: InstructionURI
	// MISSING: InputsSchemaURI
	// MISSING: Inputs
	out.State = direct.Enum_ToProto[pb.JobState](mapCtx, in.State)
	out.LabelingProgress = direct.ValueOf(in.LabelingProgress)
	out.CurrentSpend = Money_ToProto(mapCtx, in.CurrentSpend)
	out.CreateTime = direct.StringTimestamp_ToProto(mapCtx, in.CreateTime)
	out.UpdateTime = direct.StringTimestamp_ToProto(mapCtx, in.UpdateTime)
	out.Error = Status_ToProto(mapCtx, in.Error)
	// MISSING: Labels
	// MISSING: SpecialistPools
	// MISSING: EncryptionSpec
	// MISSING: ActiveLearningConfig
	return out
}
func EncryptionSpec_FromProto(mapCtx *direct.MapContext, in *pb.EncryptionSpec) *krm.EncryptionSpec {
	if in == nil {
		return nil
	}
	out := &krm.EncryptionSpec{}
	out.KMSKeyName = direct.LazyPtr(in.GetKmsKeyName())
	return out
}
func EncryptionSpec_ToProto(mapCtx *direct.MapContext, in *krm.EncryptionSpec) *pb.EncryptionSpec {
	if in == nil {
		return nil
	}
	out := &pb.EncryptionSpec{}
	out.KmsKeyName = direct.ValueOf(in.KMSKeyName)
	return out
}
func SampleConfig_FromProto(mapCtx *direct.MapContext, in *pb.SampleConfig) *krm.SampleConfig {
	if in == nil {
		return nil
	}
	out := &krm.SampleConfig{}
	out.InitialBatchSamplePercentage = direct.LazyPtr(in.GetInitialBatchSamplePercentage())
	out.FollowingBatchSamplePercentage = direct.LazyPtr(in.GetFollowingBatchSamplePercentage())
	out.SampleStrategy = direct.Enum_FromProto(mapCtx, in.GetSampleStrategy())
	return out
}
func SampleConfig_ToProto(mapCtx *direct.MapContext, in *krm.SampleConfig) *pb.SampleConfig {
	if in == nil {
		return nil
	}
	out := &pb.SampleConfig{}
	if oneof := SampleConfig_InitialBatchSamplePercentage_ToProto(mapCtx, in.InitialBatchSamplePercentage); oneof != nil {
		out.InitialBatchSampleSize = oneof
	}
	if oneof := SampleConfig_FollowingBatchSamplePercentage_ToProto(mapCtx, in.FollowingBatchSamplePercentage); oneof != nil {
		out.FollowingBatchSampleSize = oneof
	}
	out.SampleStrategy = direct.Enum_ToProto[pb.SampleConfig_SampleStrategy](mapCtx, in.SampleStrategy)
	return out
}
func TrainingConfig_FromProto(mapCtx *direct.MapContext, in *pb.TrainingConfig) *krm.TrainingConfig {
	if in == nil {
		return nil
	}
	out := &krm.TrainingConfig{}
	out.TimeoutTrainingMilliHours = direct.LazyPtr(in.GetTimeoutTrainingMilliHours())
	return out
}
func TrainingConfig_ToProto(mapCtx *direct.MapContext, in *krm.TrainingConfig) *pb.TrainingConfig {
	if in == nil {
		return nil
	}
	out := &pb.TrainingConfig{}
	out.TimeoutTrainingMilliHours = direct.ValueOf(in.TimeoutTrainingMilliHours)
	return out
}
