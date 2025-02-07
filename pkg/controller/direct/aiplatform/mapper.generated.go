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
	pb "cloud.google.com/go/aiplatform/apiv1/aiplatformpb"
	krm "github.com/GoogleCloudPlatform/k8s-config-connector/apis/aiplatform/v1alpha1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct"
	refs "github.com/GoogleCloudPlatform/k8s-config-connector/apis/refs/v1beta1"
)
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
func AiplatformDatasetObservedState_FromProto(mapCtx *direct.MapContext, in *pb.Dataset) *krm.AiplatformDatasetObservedState {
	if in == nil {
		return nil
	}
	out := &krm.AiplatformDatasetObservedState{}
	// MISSING: Name
	// MISSING: DisplayName
	// MISSING: Description
	// MISSING: MetadataSchemaURI
	// MISSING: Metadata
	// MISSING: DataItemCount
	// MISSING: CreateTime
	// MISSING: UpdateTime
	// MISSING: Etag
	// MISSING: Labels
	// MISSING: SavedQueries
	// MISSING: EncryptionSpec
	// MISSING: MetadataArtifact
	// MISSING: ModelReference
	// MISSING: SatisfiesPzs
	// MISSING: SatisfiesPzi
	return out
}
func AiplatformDatasetObservedState_ToProto(mapCtx *direct.MapContext, in *krm.AiplatformDatasetObservedState) *pb.Dataset {
	if in == nil {
		return nil
	}
	out := &pb.Dataset{}
	// MISSING: Name
	// MISSING: DisplayName
	// MISSING: Description
	// MISSING: MetadataSchemaURI
	// MISSING: Metadata
	// MISSING: DataItemCount
	// MISSING: CreateTime
	// MISSING: UpdateTime
	// MISSING: Etag
	// MISSING: Labels
	// MISSING: SavedQueries
	// MISSING: EncryptionSpec
	// MISSING: MetadataArtifact
	// MISSING: ModelReference
	// MISSING: SatisfiesPzs
	// MISSING: SatisfiesPzi
	return out
}
func AiplatformDatasetSpec_FromProto(mapCtx *direct.MapContext, in *pb.Dataset) *krm.AiplatformDatasetSpec {
	if in == nil {
		return nil
	}
	out := &krm.AiplatformDatasetSpec{}
	// MISSING: Name
	// MISSING: DisplayName
	// MISSING: Description
	// MISSING: MetadataSchemaURI
	// MISSING: Metadata
	// MISSING: DataItemCount
	// MISSING: CreateTime
	// MISSING: UpdateTime
	// MISSING: Etag
	// MISSING: Labels
	// MISSING: SavedQueries
	// MISSING: EncryptionSpec
	// MISSING: MetadataArtifact
	// MISSING: ModelReference
	// MISSING: SatisfiesPzs
	// MISSING: SatisfiesPzi
	return out
}
func AiplatformDatasetSpec_ToProto(mapCtx *direct.MapContext, in *krm.AiplatformDatasetSpec) *pb.Dataset {
	if in == nil {
		return nil
	}
	out := &pb.Dataset{}
	// MISSING: Name
	// MISSING: DisplayName
	// MISSING: Description
	// MISSING: MetadataSchemaURI
	// MISSING: Metadata
	// MISSING: DataItemCount
	// MISSING: CreateTime
	// MISSING: UpdateTime
	// MISSING: Etag
	// MISSING: Labels
	// MISSING: SavedQueries
	// MISSING: EncryptionSpec
	// MISSING: MetadataArtifact
	// MISSING: ModelReference
	// MISSING: SatisfiesPzs
	// MISSING: SatisfiesPzi
	return out
}
func AiplatformDatasetVersionObservedState_FromProto(mapCtx *direct.MapContext, in *pb.DatasetVersion) *krm.AiplatformDatasetVersionObservedState {
	if in == nil {
		return nil
	}
	out := &krm.AiplatformDatasetVersionObservedState{}
	// MISSING: Name
	// MISSING: CreateTime
	// MISSING: UpdateTime
	// MISSING: Etag
	// MISSING: BigQueryDatasetName
	// MISSING: DisplayName
	// MISSING: Metadata
	// MISSING: ModelReference
	// MISSING: SatisfiesPzs
	// MISSING: SatisfiesPzi
	return out
}
func AiplatformDatasetVersionObservedState_ToProto(mapCtx *direct.MapContext, in *krm.AiplatformDatasetVersionObservedState) *pb.DatasetVersion {
	if in == nil {
		return nil
	}
	out := &pb.DatasetVersion{}
	// MISSING: Name
	// MISSING: CreateTime
	// MISSING: UpdateTime
	// MISSING: Etag
	// MISSING: BigQueryDatasetName
	// MISSING: DisplayName
	// MISSING: Metadata
	// MISSING: ModelReference
	// MISSING: SatisfiesPzs
	// MISSING: SatisfiesPzi
	return out
}
func AiplatformDatasetVersionSpec_FromProto(mapCtx *direct.MapContext, in *pb.DatasetVersion) *krm.AiplatformDatasetVersionSpec {
	if in == nil {
		return nil
	}
	out := &krm.AiplatformDatasetVersionSpec{}
	// MISSING: Name
	// MISSING: CreateTime
	// MISSING: UpdateTime
	// MISSING: Etag
	// MISSING: BigQueryDatasetName
	// MISSING: DisplayName
	// MISSING: Metadata
	// MISSING: ModelReference
	// MISSING: SatisfiesPzs
	// MISSING: SatisfiesPzi
	return out
}
func AiplatformDatasetVersionSpec_ToProto(mapCtx *direct.MapContext, in *krm.AiplatformDatasetVersionSpec) *pb.DatasetVersion {
	if in == nil {
		return nil
	}
	out := &pb.DatasetVersion{}
	// MISSING: Name
	// MISSING: CreateTime
	// MISSING: UpdateTime
	// MISSING: Etag
	// MISSING: BigQueryDatasetName
	// MISSING: DisplayName
	// MISSING: Metadata
	// MISSING: ModelReference
	// MISSING: SatisfiesPzs
	// MISSING: SatisfiesPzi
	return out
}
func AiplatformDeploymentResourcePoolObservedState_FromProto(mapCtx *direct.MapContext, in *pb.DeploymentResourcePool) *krm.AiplatformDeploymentResourcePoolObservedState {
	if in == nil {
		return nil
	}
	out := &krm.AiplatformDeploymentResourcePoolObservedState{}
	// MISSING: Name
	// MISSING: DedicatedResources
	// MISSING: EncryptionSpec
	// MISSING: ServiceAccount
	// MISSING: DisableContainerLogging
	// MISSING: CreateTime
	// MISSING: SatisfiesPzs
	// MISSING: SatisfiesPzi
	return out
}
func AiplatformDeploymentResourcePoolObservedState_ToProto(mapCtx *direct.MapContext, in *krm.AiplatformDeploymentResourcePoolObservedState) *pb.DeploymentResourcePool {
	if in == nil {
		return nil
	}
	out := &pb.DeploymentResourcePool{}
	// MISSING: Name
	// MISSING: DedicatedResources
	// MISSING: EncryptionSpec
	// MISSING: ServiceAccount
	// MISSING: DisableContainerLogging
	// MISSING: CreateTime
	// MISSING: SatisfiesPzs
	// MISSING: SatisfiesPzi
	return out
}
func AiplatformDeploymentResourcePoolSpec_FromProto(mapCtx *direct.MapContext, in *pb.DeploymentResourcePool) *krm.AiplatformDeploymentResourcePoolSpec {
	if in == nil {
		return nil
	}
	out := &krm.AiplatformDeploymentResourcePoolSpec{}
	// MISSING: Name
	// MISSING: DedicatedResources
	// MISSING: EncryptionSpec
	// MISSING: ServiceAccount
	// MISSING: DisableContainerLogging
	// MISSING: CreateTime
	// MISSING: SatisfiesPzs
	// MISSING: SatisfiesPzi
	return out
}
func AiplatformDeploymentResourcePoolSpec_ToProto(mapCtx *direct.MapContext, in *krm.AiplatformDeploymentResourcePoolSpec) *pb.DeploymentResourcePool {
	if in == nil {
		return nil
	}
	out := &pb.DeploymentResourcePool{}
	// MISSING: Name
	// MISSING: DedicatedResources
	// MISSING: EncryptionSpec
	// MISSING: ServiceAccount
	// MISSING: DisableContainerLogging
	// MISSING: CreateTime
	// MISSING: SatisfiesPzs
	// MISSING: SatisfiesPzi
	return out
}
func AiplatformEndpointObservedState_FromProto(mapCtx *direct.MapContext, in *pb.Endpoint) *krm.AiplatformEndpointObservedState {
	if in == nil {
		return nil
	}
	out := &krm.AiplatformEndpointObservedState{}
	// MISSING: Name
	// MISSING: DisplayName
	// MISSING: Description
	// MISSING: DeployedModels
	// MISSING: TrafficSplit
	// MISSING: Etag
	// MISSING: Labels
	// MISSING: CreateTime
	// MISSING: UpdateTime
	// MISSING: EncryptionSpec
	// MISSING: Network
	// MISSING: EnablePrivateServiceConnect
	// MISSING: PrivateServiceConnectConfig
	// MISSING: ModelDeploymentMonitoringJob
	// MISSING: PredictRequestResponseLoggingConfig
	// MISSING: DedicatedEndpointEnabled
	// MISSING: DedicatedEndpointDns
	// MISSING: ClientConnectionConfig
	// MISSING: SatisfiesPzs
	// MISSING: SatisfiesPzi
	return out
}
func AiplatformEndpointObservedState_ToProto(mapCtx *direct.MapContext, in *krm.AiplatformEndpointObservedState) *pb.Endpoint {
	if in == nil {
		return nil
	}
	out := &pb.Endpoint{}
	// MISSING: Name
	// MISSING: DisplayName
	// MISSING: Description
	// MISSING: DeployedModels
	// MISSING: TrafficSplit
	// MISSING: Etag
	// MISSING: Labels
	// MISSING: CreateTime
	// MISSING: UpdateTime
	// MISSING: EncryptionSpec
	// MISSING: Network
	// MISSING: EnablePrivateServiceConnect
	// MISSING: PrivateServiceConnectConfig
	// MISSING: ModelDeploymentMonitoringJob
	// MISSING: PredictRequestResponseLoggingConfig
	// MISSING: DedicatedEndpointEnabled
	// MISSING: DedicatedEndpointDns
	// MISSING: ClientConnectionConfig
	// MISSING: SatisfiesPzs
	// MISSING: SatisfiesPzi
	return out
}
func AiplatformEndpointSpec_FromProto(mapCtx *direct.MapContext, in *pb.Endpoint) *krm.AiplatformEndpointSpec {
	if in == nil {
		return nil
	}
	out := &krm.AiplatformEndpointSpec{}
	// MISSING: Name
	// MISSING: DisplayName
	// MISSING: Description
	// MISSING: DeployedModels
	// MISSING: TrafficSplit
	// MISSING: Etag
	// MISSING: Labels
	// MISSING: CreateTime
	// MISSING: UpdateTime
	// MISSING: EncryptionSpec
	// MISSING: Network
	// MISSING: EnablePrivateServiceConnect
	// MISSING: PrivateServiceConnectConfig
	// MISSING: ModelDeploymentMonitoringJob
	// MISSING: PredictRequestResponseLoggingConfig
	// MISSING: DedicatedEndpointEnabled
	// MISSING: DedicatedEndpointDns
	// MISSING: ClientConnectionConfig
	// MISSING: SatisfiesPzs
	// MISSING: SatisfiesPzi
	return out
}
func AiplatformEndpointSpec_ToProto(mapCtx *direct.MapContext, in *krm.AiplatformEndpointSpec) *pb.Endpoint {
	if in == nil {
		return nil
	}
	out := &pb.Endpoint{}
	// MISSING: Name
	// MISSING: DisplayName
	// MISSING: Description
	// MISSING: DeployedModels
	// MISSING: TrafficSplit
	// MISSING: Etag
	// MISSING: Labels
	// MISSING: CreateTime
	// MISSING: UpdateTime
	// MISSING: EncryptionSpec
	// MISSING: Network
	// MISSING: EnablePrivateServiceConnect
	// MISSING: PrivateServiceConnectConfig
	// MISSING: ModelDeploymentMonitoringJob
	// MISSING: PredictRequestResponseLoggingConfig
	// MISSING: DedicatedEndpointEnabled
	// MISSING: DedicatedEndpointDns
	// MISSING: ClientConnectionConfig
	// MISSING: SatisfiesPzs
	// MISSING: SatisfiesPzi
	return out
}
func AiplatformEntityTypeObservedState_FromProto(mapCtx *direct.MapContext, in *pb.EntityType) *krm.AiplatformEntityTypeObservedState {
	if in == nil {
		return nil
	}
	out := &krm.AiplatformEntityTypeObservedState{}
	// MISSING: Name
	// MISSING: Description
	// MISSING: CreateTime
	// MISSING: UpdateTime
	// MISSING: Labels
	// MISSING: Etag
	// MISSING: MonitoringConfig
	// MISSING: OfflineStorageTtlDays
	// MISSING: SatisfiesPzs
	// MISSING: SatisfiesPzi
	return out
}
func AiplatformEntityTypeObservedState_ToProto(mapCtx *direct.MapContext, in *krm.AiplatformEntityTypeObservedState) *pb.EntityType {
	if in == nil {
		return nil
	}
	out := &pb.EntityType{}
	// MISSING: Name
	// MISSING: Description
	// MISSING: CreateTime
	// MISSING: UpdateTime
	// MISSING: Labels
	// MISSING: Etag
	// MISSING: MonitoringConfig
	// MISSING: OfflineStorageTtlDays
	// MISSING: SatisfiesPzs
	// MISSING: SatisfiesPzi
	return out
}
func AiplatformEntityTypeSpec_FromProto(mapCtx *direct.MapContext, in *pb.EntityType) *krm.AiplatformEntityTypeSpec {
	if in == nil {
		return nil
	}
	out := &krm.AiplatformEntityTypeSpec{}
	// MISSING: Name
	// MISSING: Description
	// MISSING: CreateTime
	// MISSING: UpdateTime
	// MISSING: Labels
	// MISSING: Etag
	// MISSING: MonitoringConfig
	// MISSING: OfflineStorageTtlDays
	// MISSING: SatisfiesPzs
	// MISSING: SatisfiesPzi
	return out
}
func AiplatformEntityTypeSpec_ToProto(mapCtx *direct.MapContext, in *krm.AiplatformEntityTypeSpec) *pb.EntityType {
	if in == nil {
		return nil
	}
	out := &pb.EntityType{}
	// MISSING: Name
	// MISSING: Description
	// MISSING: CreateTime
	// MISSING: UpdateTime
	// MISSING: Labels
	// MISSING: Etag
	// MISSING: MonitoringConfig
	// MISSING: OfflineStorageTtlDays
	// MISSING: SatisfiesPzs
	// MISSING: SatisfiesPzi
	return out
}
func AiplatformExecutionObservedState_FromProto(mapCtx *direct.MapContext, in *pb.Execution) *krm.AiplatformExecutionObservedState {
	if in == nil {
		return nil
	}
	out := &krm.AiplatformExecutionObservedState{}
	// MISSING: Name
	// MISSING: DisplayName
	// MISSING: State
	// MISSING: Etag
	// MISSING: Labels
	// MISSING: CreateTime
	// MISSING: UpdateTime
	// MISSING: SchemaTitle
	// MISSING: SchemaVersion
	// MISSING: Metadata
	// MISSING: Description
	return out
}
func AiplatformExecutionObservedState_ToProto(mapCtx *direct.MapContext, in *krm.AiplatformExecutionObservedState) *pb.Execution {
	if in == nil {
		return nil
	}
	out := &pb.Execution{}
	// MISSING: Name
	// MISSING: DisplayName
	// MISSING: State
	// MISSING: Etag
	// MISSING: Labels
	// MISSING: CreateTime
	// MISSING: UpdateTime
	// MISSING: SchemaTitle
	// MISSING: SchemaVersion
	// MISSING: Metadata
	// MISSING: Description
	return out
}
func AiplatformExecutionSpec_FromProto(mapCtx *direct.MapContext, in *pb.Execution) *krm.AiplatformExecutionSpec {
	if in == nil {
		return nil
	}
	out := &krm.AiplatformExecutionSpec{}
	// MISSING: Name
	// MISSING: DisplayName
	// MISSING: State
	// MISSING: Etag
	// MISSING: Labels
	// MISSING: CreateTime
	// MISSING: UpdateTime
	// MISSING: SchemaTitle
	// MISSING: SchemaVersion
	// MISSING: Metadata
	// MISSING: Description
	return out
}
func AiplatformExecutionSpec_ToProto(mapCtx *direct.MapContext, in *krm.AiplatformExecutionSpec) *pb.Execution {
	if in == nil {
		return nil
	}
	out := &pb.Execution{}
	// MISSING: Name
	// MISSING: DisplayName
	// MISSING: State
	// MISSING: Etag
	// MISSING: Labels
	// MISSING: CreateTime
	// MISSING: UpdateTime
	// MISSING: SchemaTitle
	// MISSING: SchemaVersion
	// MISSING: Metadata
	// MISSING: Description
	return out
}
func AiplatformFeatureGroupObservedState_FromProto(mapCtx *direct.MapContext, in *pb.FeatureGroup) *krm.AiplatformFeatureGroupObservedState {
	if in == nil {
		return nil
	}
	out := &krm.AiplatformFeatureGroupObservedState{}
	// MISSING: BigQuery
	// MISSING: Name
	// MISSING: CreateTime
	// MISSING: UpdateTime
	// MISSING: Etag
	// MISSING: Labels
	// MISSING: Description
	return out
}
func AiplatformFeatureGroupObservedState_ToProto(mapCtx *direct.MapContext, in *krm.AiplatformFeatureGroupObservedState) *pb.FeatureGroup {
	if in == nil {
		return nil
	}
	out := &pb.FeatureGroup{}
	// MISSING: BigQuery
	// MISSING: Name
	// MISSING: CreateTime
	// MISSING: UpdateTime
	// MISSING: Etag
	// MISSING: Labels
	// MISSING: Description
	return out
}
func AiplatformFeatureGroupSpec_FromProto(mapCtx *direct.MapContext, in *pb.FeatureGroup) *krm.AiplatformFeatureGroupSpec {
	if in == nil {
		return nil
	}
	out := &krm.AiplatformFeatureGroupSpec{}
	// MISSING: BigQuery
	// MISSING: Name
	// MISSING: CreateTime
	// MISSING: UpdateTime
	// MISSING: Etag
	// MISSING: Labels
	// MISSING: Description
	return out
}
func AiplatformFeatureGroupSpec_ToProto(mapCtx *direct.MapContext, in *krm.AiplatformFeatureGroupSpec) *pb.FeatureGroup {
	if in == nil {
		return nil
	}
	out := &pb.FeatureGroup{}
	// MISSING: BigQuery
	// MISSING: Name
	// MISSING: CreateTime
	// MISSING: UpdateTime
	// MISSING: Etag
	// MISSING: Labels
	// MISSING: Description
	return out
}
func AiplatformFeatureObservedState_FromProto(mapCtx *direct.MapContext, in *pb.Feature) *krm.AiplatformFeatureObservedState {
	if in == nil {
		return nil
	}
	out := &krm.AiplatformFeatureObservedState{}
	// MISSING: Name
	// MISSING: Description
	// MISSING: ValueType
	// MISSING: CreateTime
	// MISSING: UpdateTime
	// MISSING: Labels
	// MISSING: Etag
	// MISSING: DisableMonitoring
	// MISSING: MonitoringStatsAnomalies
	// MISSING: VersionColumnName
	// MISSING: PointOfContact
	return out
}
func AiplatformFeatureObservedState_ToProto(mapCtx *direct.MapContext, in *krm.AiplatformFeatureObservedState) *pb.Feature {
	if in == nil {
		return nil
	}
	out := &pb.Feature{}
	// MISSING: Name
	// MISSING: Description
	// MISSING: ValueType
	// MISSING: CreateTime
	// MISSING: UpdateTime
	// MISSING: Labels
	// MISSING: Etag
	// MISSING: DisableMonitoring
	// MISSING: MonitoringStatsAnomalies
	// MISSING: VersionColumnName
	// MISSING: PointOfContact
	return out
}
func AiplatformFeatureOnlineStoreObservedState_FromProto(mapCtx *direct.MapContext, in *pb.FeatureOnlineStore) *krm.AiplatformFeatureOnlineStoreObservedState {
	if in == nil {
		return nil
	}
	out := &krm.AiplatformFeatureOnlineStoreObservedState{}
	// MISSING: Bigtable
	// MISSING: Optimized
	// MISSING: Name
	// MISSING: CreateTime
	// MISSING: UpdateTime
	// MISSING: Etag
	// MISSING: Labels
	// MISSING: State
	// MISSING: DedicatedServingEndpoint
	// MISSING: EncryptionSpec
	// MISSING: SatisfiesPzs
	// MISSING: SatisfiesPzi
	return out
}
func AiplatformFeatureOnlineStoreObservedState_ToProto(mapCtx *direct.MapContext, in *krm.AiplatformFeatureOnlineStoreObservedState) *pb.FeatureOnlineStore {
	if in == nil {
		return nil
	}
	out := &pb.FeatureOnlineStore{}
	// MISSING: Bigtable
	// MISSING: Optimized
	// MISSING: Name
	// MISSING: CreateTime
	// MISSING: UpdateTime
	// MISSING: Etag
	// MISSING: Labels
	// MISSING: State
	// MISSING: DedicatedServingEndpoint
	// MISSING: EncryptionSpec
	// MISSING: SatisfiesPzs
	// MISSING: SatisfiesPzi
	return out
}
func AiplatformFeatureOnlineStoreSpec_FromProto(mapCtx *direct.MapContext, in *pb.FeatureOnlineStore) *krm.AiplatformFeatureOnlineStoreSpec {
	if in == nil {
		return nil
	}
	out := &krm.AiplatformFeatureOnlineStoreSpec{}
	// MISSING: Bigtable
	// MISSING: Optimized
	// MISSING: Name
	// MISSING: CreateTime
	// MISSING: UpdateTime
	// MISSING: Etag
	// MISSING: Labels
	// MISSING: State
	// MISSING: DedicatedServingEndpoint
	// MISSING: EncryptionSpec
	// MISSING: SatisfiesPzs
	// MISSING: SatisfiesPzi
	return out
}
func AiplatformFeatureOnlineStoreSpec_ToProto(mapCtx *direct.MapContext, in *krm.AiplatformFeatureOnlineStoreSpec) *pb.FeatureOnlineStore {
	if in == nil {
		return nil
	}
	out := &pb.FeatureOnlineStore{}
	// MISSING: Bigtable
	// MISSING: Optimized
	// MISSING: Name
	// MISSING: CreateTime
	// MISSING: UpdateTime
	// MISSING: Etag
	// MISSING: Labels
	// MISSING: State
	// MISSING: DedicatedServingEndpoint
	// MISSING: EncryptionSpec
	// MISSING: SatisfiesPzs
	// MISSING: SatisfiesPzi
	return out
}
func AiplatformFeatureSpec_FromProto(mapCtx *direct.MapContext, in *pb.Feature) *krm.AiplatformFeatureSpec {
	if in == nil {
		return nil
	}
	out := &krm.AiplatformFeatureSpec{}
	// MISSING: Name
	// MISSING: Description
	// MISSING: ValueType
	// MISSING: CreateTime
	// MISSING: UpdateTime
	// MISSING: Labels
	// MISSING: Etag
	// MISSING: DisableMonitoring
	// MISSING: MonitoringStatsAnomalies
	// MISSING: VersionColumnName
	// MISSING: PointOfContact
	return out
}
func AiplatformFeatureSpec_ToProto(mapCtx *direct.MapContext, in *krm.AiplatformFeatureSpec) *pb.Feature {
	if in == nil {
		return nil
	}
	out := &pb.Feature{}
	// MISSING: Name
	// MISSING: Description
	// MISSING: ValueType
	// MISSING: CreateTime
	// MISSING: UpdateTime
	// MISSING: Labels
	// MISSING: Etag
	// MISSING: DisableMonitoring
	// MISSING: MonitoringStatsAnomalies
	// MISSING: VersionColumnName
	// MISSING: PointOfContact
	return out
}
func AiplatformFeatureViewObservedState_FromProto(mapCtx *direct.MapContext, in *pb.FeatureView) *krm.AiplatformFeatureViewObservedState {
	if in == nil {
		return nil
	}
	out := &krm.AiplatformFeatureViewObservedState{}
	// MISSING: BigQuerySource
	// MISSING: FeatureRegistrySource
	// MISSING: VertexRagSource
	// MISSING: Name
	// MISSING: CreateTime
	// MISSING: UpdateTime
	// MISSING: Etag
	// MISSING: Labels
	// MISSING: SyncConfig
	// MISSING: IndexConfig
	// MISSING: SatisfiesPzs
	// MISSING: SatisfiesPzi
	return out
}
func AiplatformFeatureViewObservedState_ToProto(mapCtx *direct.MapContext, in *krm.AiplatformFeatureViewObservedState) *pb.FeatureView {
	if in == nil {
		return nil
	}
	out := &pb.FeatureView{}
	// MISSING: BigQuerySource
	// MISSING: FeatureRegistrySource
	// MISSING: VertexRagSource
	// MISSING: Name
	// MISSING: CreateTime
	// MISSING: UpdateTime
	// MISSING: Etag
	// MISSING: Labels
	// MISSING: SyncConfig
	// MISSING: IndexConfig
	// MISSING: SatisfiesPzs
	// MISSING: SatisfiesPzi
	return out
}
func AiplatformFeatureViewSpec_FromProto(mapCtx *direct.MapContext, in *pb.FeatureView) *krm.AiplatformFeatureViewSpec {
	if in == nil {
		return nil
	}
	out := &krm.AiplatformFeatureViewSpec{}
	// MISSING: BigQuerySource
	// MISSING: FeatureRegistrySource
	// MISSING: VertexRagSource
	// MISSING: Name
	// MISSING: CreateTime
	// MISSING: UpdateTime
	// MISSING: Etag
	// MISSING: Labels
	// MISSING: SyncConfig
	// MISSING: IndexConfig
	// MISSING: SatisfiesPzs
	// MISSING: SatisfiesPzi
	return out
}
func AiplatformFeatureViewSpec_ToProto(mapCtx *direct.MapContext, in *krm.AiplatformFeatureViewSpec) *pb.FeatureView {
	if in == nil {
		return nil
	}
	out := &pb.FeatureView{}
	// MISSING: BigQuerySource
	// MISSING: FeatureRegistrySource
	// MISSING: VertexRagSource
	// MISSING: Name
	// MISSING: CreateTime
	// MISSING: UpdateTime
	// MISSING: Etag
	// MISSING: Labels
	// MISSING: SyncConfig
	// MISSING: IndexConfig
	// MISSING: SatisfiesPzs
	// MISSING: SatisfiesPzi
	return out
}
func AiplatformFeatureViewSyncObservedState_FromProto(mapCtx *direct.MapContext, in *pb.FeatureViewSync) *krm.AiplatformFeatureViewSyncObservedState {
	if in == nil {
		return nil
	}
	out := &krm.AiplatformFeatureViewSyncObservedState{}
	// MISSING: Name
	// MISSING: CreateTime
	// MISSING: RunTime
	// MISSING: FinalStatus
	// MISSING: SyncSummary
	// MISSING: SatisfiesPzs
	// MISSING: SatisfiesPzi
	return out
}
func AiplatformFeatureViewSyncObservedState_ToProto(mapCtx *direct.MapContext, in *krm.AiplatformFeatureViewSyncObservedState) *pb.FeatureViewSync {
	if in == nil {
		return nil
	}
	out := &pb.FeatureViewSync{}
	// MISSING: Name
	// MISSING: CreateTime
	// MISSING: RunTime
	// MISSING: FinalStatus
	// MISSING: SyncSummary
	// MISSING: SatisfiesPzs
	// MISSING: SatisfiesPzi
	return out
}
func AiplatformFeatureViewSyncSpec_FromProto(mapCtx *direct.MapContext, in *pb.FeatureViewSync) *krm.AiplatformFeatureViewSyncSpec {
	if in == nil {
		return nil
	}
	out := &krm.AiplatformFeatureViewSyncSpec{}
	// MISSING: Name
	// MISSING: CreateTime
	// MISSING: RunTime
	// MISSING: FinalStatus
	// MISSING: SyncSummary
	// MISSING: SatisfiesPzs
	// MISSING: SatisfiesPzi
	return out
}
func AiplatformFeatureViewSyncSpec_ToProto(mapCtx *direct.MapContext, in *krm.AiplatformFeatureViewSyncSpec) *pb.FeatureViewSync {
	if in == nil {
		return nil
	}
	out := &pb.FeatureViewSync{}
	// MISSING: Name
	// MISSING: CreateTime
	// MISSING: RunTime
	// MISSING: FinalStatus
	// MISSING: SyncSummary
	// MISSING: SatisfiesPzs
	// MISSING: SatisfiesPzi
	return out
}
func AiplatformFeaturestoreObservedState_FromProto(mapCtx *direct.MapContext, in *pb.Featurestore) *krm.AiplatformFeaturestoreObservedState {
	if in == nil {
		return nil
	}
	out := &krm.AiplatformFeaturestoreObservedState{}
	// MISSING: Name
	// MISSING: CreateTime
	// MISSING: UpdateTime
	// MISSING: Etag
	// MISSING: Labels
	// MISSING: OnlineServingConfig
	// MISSING: State
	// MISSING: OnlineStorageTtlDays
	// MISSING: EncryptionSpec
	// MISSING: SatisfiesPzs
	// MISSING: SatisfiesPzi
	return out
}
func AiplatformFeaturestoreObservedState_ToProto(mapCtx *direct.MapContext, in *krm.AiplatformFeaturestoreObservedState) *pb.Featurestore {
	if in == nil {
		return nil
	}
	out := &pb.Featurestore{}
	// MISSING: Name
	// MISSING: CreateTime
	// MISSING: UpdateTime
	// MISSING: Etag
	// MISSING: Labels
	// MISSING: OnlineServingConfig
	// MISSING: State
	// MISSING: OnlineStorageTtlDays
	// MISSING: EncryptionSpec
	// MISSING: SatisfiesPzs
	// MISSING: SatisfiesPzi
	return out
}
func AiplatformFeaturestoreSpec_FromProto(mapCtx *direct.MapContext, in *pb.Featurestore) *krm.AiplatformFeaturestoreSpec {
	if in == nil {
		return nil
	}
	out := &krm.AiplatformFeaturestoreSpec{}
	// MISSING: Name
	// MISSING: CreateTime
	// MISSING: UpdateTime
	// MISSING: Etag
	// MISSING: Labels
	// MISSING: OnlineServingConfig
	// MISSING: State
	// MISSING: OnlineStorageTtlDays
	// MISSING: EncryptionSpec
	// MISSING: SatisfiesPzs
	// MISSING: SatisfiesPzi
	return out
}
func AiplatformFeaturestoreSpec_ToProto(mapCtx *direct.MapContext, in *krm.AiplatformFeaturestoreSpec) *pb.Featurestore {
	if in == nil {
		return nil
	}
	out := &pb.Featurestore{}
	// MISSING: Name
	// MISSING: CreateTime
	// MISSING: UpdateTime
	// MISSING: Etag
	// MISSING: Labels
	// MISSING: OnlineServingConfig
	// MISSING: State
	// MISSING: OnlineStorageTtlDays
	// MISSING: EncryptionSpec
	// MISSING: SatisfiesPzs
	// MISSING: SatisfiesPzi
	return out
}
func AiplatformHyperparameterTuningJobObservedState_FromProto(mapCtx *direct.MapContext, in *pb.HyperparameterTuningJob) *krm.AiplatformHyperparameterTuningJobObservedState {
	if in == nil {
		return nil
	}
	out := &krm.AiplatformHyperparameterTuningJobObservedState{}
	// MISSING: Name
	// MISSING: DisplayName
	// MISSING: StudySpec
	// MISSING: MaxTrialCount
	// MISSING: ParallelTrialCount
	// MISSING: MaxFailedTrialCount
	// MISSING: TrialJobSpec
	// MISSING: Trials
	// MISSING: State
	// MISSING: CreateTime
	// MISSING: StartTime
	// MISSING: EndTime
	// MISSING: UpdateTime
	// MISSING: Error
	// MISSING: Labels
	// MISSING: EncryptionSpec
	// MISSING: SatisfiesPzs
	// MISSING: SatisfiesPzi
	return out
}
func AiplatformHyperparameterTuningJobObservedState_ToProto(mapCtx *direct.MapContext, in *krm.AiplatformHyperparameterTuningJobObservedState) *pb.HyperparameterTuningJob {
	if in == nil {
		return nil
	}
	out := &pb.HyperparameterTuningJob{}
	// MISSING: Name
	// MISSING: DisplayName
	// MISSING: StudySpec
	// MISSING: MaxTrialCount
	// MISSING: ParallelTrialCount
	// MISSING: MaxFailedTrialCount
	// MISSING: TrialJobSpec
	// MISSING: Trials
	// MISSING: State
	// MISSING: CreateTime
	// MISSING: StartTime
	// MISSING: EndTime
	// MISSING: UpdateTime
	// MISSING: Error
	// MISSING: Labels
	// MISSING: EncryptionSpec
	// MISSING: SatisfiesPzs
	// MISSING: SatisfiesPzi
	return out
}
func AiplatformHyperparameterTuningJobSpec_FromProto(mapCtx *direct.MapContext, in *pb.HyperparameterTuningJob) *krm.AiplatformHyperparameterTuningJobSpec {
	if in == nil {
		return nil
	}
	out := &krm.AiplatformHyperparameterTuningJobSpec{}
	// MISSING: Name
	// MISSING: DisplayName
	// MISSING: StudySpec
	// MISSING: MaxTrialCount
	// MISSING: ParallelTrialCount
	// MISSING: MaxFailedTrialCount
	// MISSING: TrialJobSpec
	// MISSING: Trials
	// MISSING: State
	// MISSING: CreateTime
	// MISSING: StartTime
	// MISSING: EndTime
	// MISSING: UpdateTime
	// MISSING: Error
	// MISSING: Labels
	// MISSING: EncryptionSpec
	// MISSING: SatisfiesPzs
	// MISSING: SatisfiesPzi
	return out
}
func AiplatformHyperparameterTuningJobSpec_ToProto(mapCtx *direct.MapContext, in *krm.AiplatformHyperparameterTuningJobSpec) *pb.HyperparameterTuningJob {
	if in == nil {
		return nil
	}
	out := &pb.HyperparameterTuningJob{}
	// MISSING: Name
	// MISSING: DisplayName
	// MISSING: StudySpec
	// MISSING: MaxTrialCount
	// MISSING: ParallelTrialCount
	// MISSING: MaxFailedTrialCount
	// MISSING: TrialJobSpec
	// MISSING: Trials
	// MISSING: State
	// MISSING: CreateTime
	// MISSING: StartTime
	// MISSING: EndTime
	// MISSING: UpdateTime
	// MISSING: Error
	// MISSING: Labels
	// MISSING: EncryptionSpec
	// MISSING: SatisfiesPzs
	// MISSING: SatisfiesPzi
	return out
}
func AiplatformIndexEndpointObservedState_FromProto(mapCtx *direct.MapContext, in *pb.IndexEndpoint) *krm.AiplatformIndexEndpointObservedState {
	if in == nil {
		return nil
	}
	out := &krm.AiplatformIndexEndpointObservedState{}
	// MISSING: Name
	// MISSING: DisplayName
	// MISSING: Description
	// MISSING: DeployedIndexes
	// MISSING: Etag
	// MISSING: Labels
	// MISSING: CreateTime
	// MISSING: UpdateTime
	// MISSING: Network
	// MISSING: EnablePrivateServiceConnect
	// MISSING: PrivateServiceConnectConfig
	// MISSING: PublicEndpointEnabled
	// MISSING: PublicEndpointDomainName
	// MISSING: EncryptionSpec
	// MISSING: SatisfiesPzs
	// MISSING: SatisfiesPzi
	return out
}
func AiplatformIndexEndpointObservedState_ToProto(mapCtx *direct.MapContext, in *krm.AiplatformIndexEndpointObservedState) *pb.IndexEndpoint {
	if in == nil {
		return nil
	}
	out := &pb.IndexEndpoint{}
	// MISSING: Name
	// MISSING: DisplayName
	// MISSING: Description
	// MISSING: DeployedIndexes
	// MISSING: Etag
	// MISSING: Labels
	// MISSING: CreateTime
	// MISSING: UpdateTime
	// MISSING: Network
	// MISSING: EnablePrivateServiceConnect
	// MISSING: PrivateServiceConnectConfig
	// MISSING: PublicEndpointEnabled
	// MISSING: PublicEndpointDomainName
	// MISSING: EncryptionSpec
	// MISSING: SatisfiesPzs
	// MISSING: SatisfiesPzi
	return out
}
func AiplatformIndexEndpointSpec_FromProto(mapCtx *direct.MapContext, in *pb.IndexEndpoint) *krm.AiplatformIndexEndpointSpec {
	if in == nil {
		return nil
	}
	out := &krm.AiplatformIndexEndpointSpec{}
	// MISSING: Name
	// MISSING: DisplayName
	// MISSING: Description
	// MISSING: DeployedIndexes
	// MISSING: Etag
	// MISSING: Labels
	// MISSING: CreateTime
	// MISSING: UpdateTime
	// MISSING: Network
	// MISSING: EnablePrivateServiceConnect
	// MISSING: PrivateServiceConnectConfig
	// MISSING: PublicEndpointEnabled
	// MISSING: PublicEndpointDomainName
	// MISSING: EncryptionSpec
	// MISSING: SatisfiesPzs
	// MISSING: SatisfiesPzi
	return out
}
func AiplatformIndexEndpointSpec_ToProto(mapCtx *direct.MapContext, in *krm.AiplatformIndexEndpointSpec) *pb.IndexEndpoint {
	if in == nil {
		return nil
	}
	out := &pb.IndexEndpoint{}
	// MISSING: Name
	// MISSING: DisplayName
	// MISSING: Description
	// MISSING: DeployedIndexes
	// MISSING: Etag
	// MISSING: Labels
	// MISSING: CreateTime
	// MISSING: UpdateTime
	// MISSING: Network
	// MISSING: EnablePrivateServiceConnect
	// MISSING: PrivateServiceConnectConfig
	// MISSING: PublicEndpointEnabled
	// MISSING: PublicEndpointDomainName
	// MISSING: EncryptionSpec
	// MISSING: SatisfiesPzs
	// MISSING: SatisfiesPzi
	return out
}
func AiplatformIndexObservedState_FromProto(mapCtx *direct.MapContext, in *pb.Index) *krm.AiplatformIndexObservedState {
	if in == nil {
		return nil
	}
	out := &krm.AiplatformIndexObservedState{}
	// MISSING: Name
	// MISSING: DisplayName
	// MISSING: Description
	// MISSING: MetadataSchemaURI
	// MISSING: Metadata
	// MISSING: DeployedIndexes
	// MISSING: Etag
	// MISSING: Labels
	// MISSING: CreateTime
	// MISSING: UpdateTime
	// MISSING: IndexStats
	// MISSING: IndexUpdateMethod
	// MISSING: EncryptionSpec
	// MISSING: SatisfiesPzs
	// MISSING: SatisfiesPzi
	return out
}
func AiplatformIndexObservedState_ToProto(mapCtx *direct.MapContext, in *krm.AiplatformIndexObservedState) *pb.Index {
	if in == nil {
		return nil
	}
	out := &pb.Index{}
	// MISSING: Name
	// MISSING: DisplayName
	// MISSING: Description
	// MISSING: MetadataSchemaURI
	// MISSING: Metadata
	// MISSING: DeployedIndexes
	// MISSING: Etag
	// MISSING: Labels
	// MISSING: CreateTime
	// MISSING: UpdateTime
	// MISSING: IndexStats
	// MISSING: IndexUpdateMethod
	// MISSING: EncryptionSpec
	// MISSING: SatisfiesPzs
	// MISSING: SatisfiesPzi
	return out
}
func AiplatformIndexSpec_FromProto(mapCtx *direct.MapContext, in *pb.Index) *krm.AiplatformIndexSpec {
	if in == nil {
		return nil
	}
	out := &krm.AiplatformIndexSpec{}
	// MISSING: Name
	// MISSING: DisplayName
	// MISSING: Description
	// MISSING: MetadataSchemaURI
	// MISSING: Metadata
	// MISSING: DeployedIndexes
	// MISSING: Etag
	// MISSING: Labels
	// MISSING: CreateTime
	// MISSING: UpdateTime
	// MISSING: IndexStats
	// MISSING: IndexUpdateMethod
	// MISSING: EncryptionSpec
	// MISSING: SatisfiesPzs
	// MISSING: SatisfiesPzi
	return out
}
func AiplatformIndexSpec_ToProto(mapCtx *direct.MapContext, in *krm.AiplatformIndexSpec) *pb.Index {
	if in == nil {
		return nil
	}
	out := &pb.Index{}
	// MISSING: Name
	// MISSING: DisplayName
	// MISSING: Description
	// MISSING: MetadataSchemaURI
	// MISSING: Metadata
	// MISSING: DeployedIndexes
	// MISSING: Etag
	// MISSING: Labels
	// MISSING: CreateTime
	// MISSING: UpdateTime
	// MISSING: IndexStats
	// MISSING: IndexUpdateMethod
	// MISSING: EncryptionSpec
	// MISSING: SatisfiesPzs
	// MISSING: SatisfiesPzi
	return out
}
func AutomaticResources_FromProto(mapCtx *direct.MapContext, in *pb.AutomaticResources) *krm.AutomaticResources {
	if in == nil {
		return nil
	}
	out := &krm.AutomaticResources{}
	out.MinReplicaCount = direct.LazyPtr(in.GetMinReplicaCount())
	out.MaxReplicaCount = direct.LazyPtr(in.GetMaxReplicaCount())
	return out
}
func AutomaticResources_ToProto(mapCtx *direct.MapContext, in *krm.AutomaticResources) *pb.AutomaticResources {
	if in == nil {
		return nil
	}
	out := &pb.AutomaticResources{}
	out.MinReplicaCount = direct.ValueOf(in.MinReplicaCount)
	out.MaxReplicaCount = direct.ValueOf(in.MaxReplicaCount)
	return out
}
func AutoscalingMetricSpec_FromProto(mapCtx *direct.MapContext, in *pb.AutoscalingMetricSpec) *krm.AutoscalingMetricSpec {
	if in == nil {
		return nil
	}
	out := &krm.AutoscalingMetricSpec{}
	out.MetricName = direct.LazyPtr(in.GetMetricName())
	out.Target = direct.LazyPtr(in.GetTarget())
	return out
}
func AutoscalingMetricSpec_ToProto(mapCtx *direct.MapContext, in *krm.AutoscalingMetricSpec) *pb.AutoscalingMetricSpec {
	if in == nil {
		return nil
	}
	out := &pb.AutoscalingMetricSpec{}
	out.MetricName = direct.ValueOf(in.MetricName)
	out.Target = direct.ValueOf(in.Target)
	return out
}
func DedicatedResources_FromProto(mapCtx *direct.MapContext, in *pb.DedicatedResources) *krm.DedicatedResources {
	if in == nil {
		return nil
	}
	out := &krm.DedicatedResources{}
	out.MachineSpec = MachineSpec_FromProto(mapCtx, in.GetMachineSpec())
	out.MinReplicaCount = direct.LazyPtr(in.GetMinReplicaCount())
	out.MaxReplicaCount = direct.LazyPtr(in.GetMaxReplicaCount())
	out.RequiredReplicaCount = direct.LazyPtr(in.GetRequiredReplicaCount())
	out.AutoscalingMetricSpecs = direct.Slice_FromProto(mapCtx, in.AutoscalingMetricSpecs, AutoscalingMetricSpec_FromProto)
	out.Spot = direct.LazyPtr(in.GetSpot())
	return out
}
func DedicatedResources_ToProto(mapCtx *direct.MapContext, in *krm.DedicatedResources) *pb.DedicatedResources {
	if in == nil {
		return nil
	}
	out := &pb.DedicatedResources{}
	out.MachineSpec = MachineSpec_ToProto(mapCtx, in.MachineSpec)
	out.MinReplicaCount = direct.ValueOf(in.MinReplicaCount)
	out.MaxReplicaCount = direct.ValueOf(in.MaxReplicaCount)
	out.RequiredReplicaCount = direct.ValueOf(in.RequiredReplicaCount)
	out.AutoscalingMetricSpecs = direct.Slice_ToProto(mapCtx, in.AutoscalingMetricSpecs, AutoscalingMetricSpec_ToProto)
	out.Spot = direct.ValueOf(in.Spot)
	return out
}
func DeployedIndex_FromProto(mapCtx *direct.MapContext, in *pb.DeployedIndex) *krm.DeployedIndex {
	if in == nil {
		return nil
	}
	out := &krm.DeployedIndex{}
	out.ID = direct.LazyPtr(in.GetId())
	out.Index = direct.LazyPtr(in.GetIndex())
	out.DisplayName = direct.LazyPtr(in.GetDisplayName())
	// MISSING: CreateTime
	// MISSING: PrivateEndpoints
	// MISSING: IndexSyncTime
	out.AutomaticResources = AutomaticResources_FromProto(mapCtx, in.GetAutomaticResources())
	out.DedicatedResources = DedicatedResources_FromProto(mapCtx, in.GetDedicatedResources())
	out.EnableAccessLogging = direct.LazyPtr(in.GetEnableAccessLogging())
	out.DeployedIndexAuthConfig = DeployedIndexAuthConfig_FromProto(mapCtx, in.GetDeployedIndexAuthConfig())
	out.ReservedIPRanges = in.ReservedIpRanges
	out.DeploymentGroup = direct.LazyPtr(in.GetDeploymentGroup())
	out.PscAutomationConfigs = direct.Slice_FromProto(mapCtx, in.PscAutomationConfigs, PSCAutomationConfig_FromProto)
	return out
}
func DeployedIndex_ToProto(mapCtx *direct.MapContext, in *krm.DeployedIndex) *pb.DeployedIndex {
	if in == nil {
		return nil
	}
	out := &pb.DeployedIndex{}
	out.Id = direct.ValueOf(in.ID)
	out.Index = direct.ValueOf(in.Index)
	out.DisplayName = direct.ValueOf(in.DisplayName)
	// MISSING: CreateTime
	// MISSING: PrivateEndpoints
	// MISSING: IndexSyncTime
	out.AutomaticResources = AutomaticResources_ToProto(mapCtx, in.AutomaticResources)
	out.DedicatedResources = DedicatedResources_ToProto(mapCtx, in.DedicatedResources)
	out.EnableAccessLogging = direct.ValueOf(in.EnableAccessLogging)
	out.DeployedIndexAuthConfig = DeployedIndexAuthConfig_ToProto(mapCtx, in.DeployedIndexAuthConfig)
	out.ReservedIpRanges = in.ReservedIPRanges
	out.DeploymentGroup = direct.ValueOf(in.DeploymentGroup)
	out.PscAutomationConfigs = direct.Slice_ToProto(mapCtx, in.PscAutomationConfigs, PSCAutomationConfig_ToProto)
	return out
}
func DeployedIndexAuthConfig_FromProto(mapCtx *direct.MapContext, in *pb.DeployedIndexAuthConfig) *krm.DeployedIndexAuthConfig {
	if in == nil {
		return nil
	}
	out := &krm.DeployedIndexAuthConfig{}
	out.AuthProvider = DeployedIndexAuthConfig_AuthProvider_FromProto(mapCtx, in.GetAuthProvider())
	return out
}
func DeployedIndexAuthConfig_ToProto(mapCtx *direct.MapContext, in *krm.DeployedIndexAuthConfig) *pb.DeployedIndexAuthConfig {
	if in == nil {
		return nil
	}
	out := &pb.DeployedIndexAuthConfig{}
	out.AuthProvider = DeployedIndexAuthConfig_AuthProvider_ToProto(mapCtx, in.AuthProvider)
	return out
}
func DeployedIndexAuthConfig_AuthProvider_FromProto(mapCtx *direct.MapContext, in *pb.DeployedIndexAuthConfig_AuthProvider) *krm.DeployedIndexAuthConfig_AuthProvider {
	if in == nil {
		return nil
	}
	out := &krm.DeployedIndexAuthConfig_AuthProvider{}
	out.Audiences = in.Audiences
	out.AllowedIssuers = in.AllowedIssuers
	return out
}
func DeployedIndexAuthConfig_AuthProvider_ToProto(mapCtx *direct.MapContext, in *krm.DeployedIndexAuthConfig_AuthProvider) *pb.DeployedIndexAuthConfig_AuthProvider {
	if in == nil {
		return nil
	}
	out := &pb.DeployedIndexAuthConfig_AuthProvider{}
	out.Audiences = in.Audiences
	out.AllowedIssuers = in.AllowedIssuers
	return out
}
func DeployedIndexObservedState_FromProto(mapCtx *direct.MapContext, in *pb.DeployedIndex) *krm.DeployedIndexObservedState {
	if in == nil {
		return nil
	}
	out := &krm.DeployedIndexObservedState{}
	// MISSING: ID
	// MISSING: Index
	// MISSING: DisplayName
	out.CreateTime = direct.StringTimestamp_FromProto(mapCtx, in.GetCreateTime())
	out.PrivateEndpoints = IndexPrivateEndpoints_FromProto(mapCtx, in.GetPrivateEndpoints())
	out.IndexSyncTime = direct.StringTimestamp_FromProto(mapCtx, in.GetIndexSyncTime())
	// MISSING: AutomaticResources
	// MISSING: DedicatedResources
	// MISSING: EnableAccessLogging
	// MISSING: DeployedIndexAuthConfig
	// MISSING: ReservedIPRanges
	// MISSING: DeploymentGroup
	// MISSING: PscAutomationConfigs
	return out
}
func DeployedIndexObservedState_ToProto(mapCtx *direct.MapContext, in *krm.DeployedIndexObservedState) *pb.DeployedIndex {
	if in == nil {
		return nil
	}
	out := &pb.DeployedIndex{}
	// MISSING: ID
	// MISSING: Index
	// MISSING: DisplayName
	out.CreateTime = direct.StringTimestamp_ToProto(mapCtx, in.CreateTime)
	out.PrivateEndpoints = IndexPrivateEndpoints_ToProto(mapCtx, in.PrivateEndpoints)
	out.IndexSyncTime = direct.StringTimestamp_ToProto(mapCtx, in.IndexSyncTime)
	// MISSING: AutomaticResources
	// MISSING: DedicatedResources
	// MISSING: EnableAccessLogging
	// MISSING: DeployedIndexAuthConfig
	// MISSING: ReservedIPRanges
	// MISSING: DeploymentGroup
	// MISSING: PscAutomationConfigs
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
func IndexEndpoint_FromProto(mapCtx *direct.MapContext, in *pb.IndexEndpoint) *krm.IndexEndpoint {
	if in == nil {
		return nil
	}
	out := &krm.IndexEndpoint{}
	// MISSING: Name
	out.DisplayName = direct.LazyPtr(in.GetDisplayName())
	out.Description = direct.LazyPtr(in.GetDescription())
	// MISSING: DeployedIndexes
	out.Etag = direct.LazyPtr(in.GetEtag())
	out.Labels = in.Labels
	// MISSING: CreateTime
	// MISSING: UpdateTime
	out.Network = direct.LazyPtr(in.GetNetwork())
	out.EnablePrivateServiceConnect = direct.LazyPtr(in.GetEnablePrivateServiceConnect())
	out.PrivateServiceConnectConfig = PrivateServiceConnectConfig_FromProto(mapCtx, in.GetPrivateServiceConnectConfig())
	out.PublicEndpointEnabled = direct.LazyPtr(in.GetPublicEndpointEnabled())
	// MISSING: PublicEndpointDomainName
	out.EncryptionSpec = EncryptionSpec_FromProto(mapCtx, in.GetEncryptionSpec())
	// MISSING: SatisfiesPzs
	// MISSING: SatisfiesPzi
	return out
}
func IndexEndpoint_ToProto(mapCtx *direct.MapContext, in *krm.IndexEndpoint) *pb.IndexEndpoint {
	if in == nil {
		return nil
	}
	out := &pb.IndexEndpoint{}
	// MISSING: Name
	out.DisplayName = direct.ValueOf(in.DisplayName)
	out.Description = direct.ValueOf(in.Description)
	// MISSING: DeployedIndexes
	out.Etag = direct.ValueOf(in.Etag)
	out.Labels = in.Labels
	// MISSING: CreateTime
	// MISSING: UpdateTime
	out.Network = direct.ValueOf(in.Network)
	out.EnablePrivateServiceConnect = direct.ValueOf(in.EnablePrivateServiceConnect)
	out.PrivateServiceConnectConfig = PrivateServiceConnectConfig_ToProto(mapCtx, in.PrivateServiceConnectConfig)
	out.PublicEndpointEnabled = direct.ValueOf(in.PublicEndpointEnabled)
	// MISSING: PublicEndpointDomainName
	out.EncryptionSpec = EncryptionSpec_ToProto(mapCtx, in.EncryptionSpec)
	// MISSING: SatisfiesPzs
	// MISSING: SatisfiesPzi
	return out
}
func IndexEndpointObservedState_FromProto(mapCtx *direct.MapContext, in *pb.IndexEndpoint) *krm.IndexEndpointObservedState {
	if in == nil {
		return nil
	}
	out := &krm.IndexEndpointObservedState{}
	out.Name = direct.LazyPtr(in.GetName())
	// MISSING: DisplayName
	// MISSING: Description
	out.DeployedIndexes = direct.Slice_FromProto(mapCtx, in.DeployedIndexes, DeployedIndex_FromProto)
	// MISSING: Etag
	// MISSING: Labels
	out.CreateTime = direct.StringTimestamp_FromProto(mapCtx, in.GetCreateTime())
	out.UpdateTime = direct.StringTimestamp_FromProto(mapCtx, in.GetUpdateTime())
	// MISSING: Network
	// MISSING: EnablePrivateServiceConnect
	out.PrivateServiceConnectConfig = PrivateServiceConnectConfigObservedState_FromProto(mapCtx, in.GetPrivateServiceConnectConfig())
	// MISSING: PublicEndpointEnabled
	out.PublicEndpointDomainName = direct.LazyPtr(in.GetPublicEndpointDomainName())
	// MISSING: EncryptionSpec
	out.SatisfiesPzs = direct.LazyPtr(in.GetSatisfiesPzs())
	out.SatisfiesPzi = direct.LazyPtr(in.GetSatisfiesPzi())
	return out
}
func IndexEndpointObservedState_ToProto(mapCtx *direct.MapContext, in *krm.IndexEndpointObservedState) *pb.IndexEndpoint {
	if in == nil {
		return nil
	}
	out := &pb.IndexEndpoint{}
	out.Name = direct.ValueOf(in.Name)
	// MISSING: DisplayName
	// MISSING: Description
	out.DeployedIndexes = direct.Slice_ToProto(mapCtx, in.DeployedIndexes, DeployedIndex_ToProto)
	// MISSING: Etag
	// MISSING: Labels
	out.CreateTime = direct.StringTimestamp_ToProto(mapCtx, in.CreateTime)
	out.UpdateTime = direct.StringTimestamp_ToProto(mapCtx, in.UpdateTime)
	// MISSING: Network
	// MISSING: EnablePrivateServiceConnect
	out.PrivateServiceConnectConfig = PrivateServiceConnectConfigObservedState_ToProto(mapCtx, in.PrivateServiceConnectConfig)
	// MISSING: PublicEndpointEnabled
	out.PublicEndpointDomainName = direct.ValueOf(in.PublicEndpointDomainName)
	// MISSING: EncryptionSpec
	out.SatisfiesPzs = direct.ValueOf(in.SatisfiesPzs)
	out.SatisfiesPzi = direct.ValueOf(in.SatisfiesPzi)
	return out
}
func IndexPrivateEndpoints_FromProto(mapCtx *direct.MapContext, in *pb.IndexPrivateEndpoints) *krm.IndexPrivateEndpoints {
	if in == nil {
		return nil
	}
	out := &krm.IndexPrivateEndpoints{}
	// MISSING: MatchGrpcAddress
	// MISSING: ServiceAttachment
	// MISSING: PscAutomatedEndpoints
	return out
}
func IndexPrivateEndpoints_ToProto(mapCtx *direct.MapContext, in *krm.IndexPrivateEndpoints) *pb.IndexPrivateEndpoints {
	if in == nil {
		return nil
	}
	out := &pb.IndexPrivateEndpoints{}
	// MISSING: MatchGrpcAddress
	// MISSING: ServiceAttachment
	// MISSING: PscAutomatedEndpoints
	return out
}
func IndexPrivateEndpointsObservedState_FromProto(mapCtx *direct.MapContext, in *pb.IndexPrivateEndpoints) *krm.IndexPrivateEndpointsObservedState {
	if in == nil {
		return nil
	}
	out := &krm.IndexPrivateEndpointsObservedState{}
	out.MatchGrpcAddress = direct.LazyPtr(in.GetMatchGrpcAddress())
	out.ServiceAttachment = direct.LazyPtr(in.GetServiceAttachment())
	out.PscAutomatedEndpoints = direct.Slice_FromProto(mapCtx, in.PscAutomatedEndpoints, PscAutomatedEndpoints_FromProto)
	return out
}
func IndexPrivateEndpointsObservedState_ToProto(mapCtx *direct.MapContext, in *krm.IndexPrivateEndpointsObservedState) *pb.IndexPrivateEndpoints {
	if in == nil {
		return nil
	}
	out := &pb.IndexPrivateEndpoints{}
	out.MatchGrpcAddress = direct.ValueOf(in.MatchGrpcAddress)
	out.ServiceAttachment = direct.ValueOf(in.ServiceAttachment)
	out.PscAutomatedEndpoints = direct.Slice_ToProto(mapCtx, in.PscAutomatedEndpoints, PscAutomatedEndpoints_ToProto)
	return out
}
func MachineSpec_FromProto(mapCtx *direct.MapContext, in *pb.MachineSpec) *krm.MachineSpec {
	if in == nil {
		return nil
	}
	out := &krm.MachineSpec{}
	out.MachineType = direct.LazyPtr(in.GetMachineType())
	out.AcceleratorType = direct.Enum_FromProto(mapCtx, in.GetAcceleratorType())
	out.AcceleratorCount = direct.LazyPtr(in.GetAcceleratorCount())
	out.TpuTopology = direct.LazyPtr(in.GetTpuTopology())
	out.ReservationAffinity = ReservationAffinity_FromProto(mapCtx, in.GetReservationAffinity())
	return out
}
func MachineSpec_ToProto(mapCtx *direct.MapContext, in *krm.MachineSpec) *pb.MachineSpec {
	if in == nil {
		return nil
	}
	out := &pb.MachineSpec{}
	out.MachineType = direct.ValueOf(in.MachineType)
	out.AcceleratorType = direct.Enum_ToProto[pb.AcceleratorType](mapCtx, in.AcceleratorType)
	out.AcceleratorCount = direct.ValueOf(in.AcceleratorCount)
	out.TpuTopology = direct.ValueOf(in.TpuTopology)
	out.ReservationAffinity = ReservationAffinity_ToProto(mapCtx, in.ReservationAffinity)
	return out
}
func PSCAutomationConfig_FromProto(mapCtx *direct.MapContext, in *pb.PSCAutomationConfig) *krm.PSCAutomationConfig {
	if in == nil {
		return nil
	}
	out := &krm.PSCAutomationConfig{}
	out.ProjectID = direct.LazyPtr(in.GetProjectId())
	out.Network = direct.LazyPtr(in.GetNetwork())
	return out
}
func PSCAutomationConfig_ToProto(mapCtx *direct.MapContext, in *krm.PSCAutomationConfig) *pb.PSCAutomationConfig {
	if in == nil {
		return nil
	}
	out := &pb.PSCAutomationConfig{}
	out.ProjectId = direct.ValueOf(in.ProjectID)
	out.Network = direct.ValueOf(in.Network)
	return out
}
func PrivateServiceConnectConfig_FromProto(mapCtx *direct.MapContext, in *pb.PrivateServiceConnectConfig) *krm.PrivateServiceConnectConfig {
	if in == nil {
		return nil
	}
	out := &krm.PrivateServiceConnectConfig{}
	out.EnablePrivateServiceConnect = direct.LazyPtr(in.GetEnablePrivateServiceConnect())
	out.ProjectAllowlist = in.ProjectAllowlist
	// MISSING: ServiceAttachment
	return out
}
func PrivateServiceConnectConfig_ToProto(mapCtx *direct.MapContext, in *krm.PrivateServiceConnectConfig) *pb.PrivateServiceConnectConfig {
	if in == nil {
		return nil
	}
	out := &pb.PrivateServiceConnectConfig{}
	out.EnablePrivateServiceConnect = direct.ValueOf(in.EnablePrivateServiceConnect)
	out.ProjectAllowlist = in.ProjectAllowlist
	// MISSING: ServiceAttachment
	return out
}
func PrivateServiceConnectConfigObservedState_FromProto(mapCtx *direct.MapContext, in *pb.PrivateServiceConnectConfig) *krm.PrivateServiceConnectConfigObservedState {
	if in == nil {
		return nil
	}
	out := &krm.PrivateServiceConnectConfigObservedState{}
	// MISSING: EnablePrivateServiceConnect
	// MISSING: ProjectAllowlist
	out.ServiceAttachment = direct.LazyPtr(in.GetServiceAttachment())
	return out
}
func PrivateServiceConnectConfigObservedState_ToProto(mapCtx *direct.MapContext, in *krm.PrivateServiceConnectConfigObservedState) *pb.PrivateServiceConnectConfig {
	if in == nil {
		return nil
	}
	out := &pb.PrivateServiceConnectConfig{}
	// MISSING: EnablePrivateServiceConnect
	// MISSING: ProjectAllowlist
	out.ServiceAttachment = direct.ValueOf(in.ServiceAttachment)
	return out
}
func PscAutomatedEndpoints_FromProto(mapCtx *direct.MapContext, in *pb.PscAutomatedEndpoints) *krm.PscAutomatedEndpoints {
	if in == nil {
		return nil
	}
	out := &krm.PscAutomatedEndpoints{}
	out.ProjectID = direct.LazyPtr(in.GetProjectId())
	out.Network = direct.LazyPtr(in.GetNetwork())
	out.MatchAddress = direct.LazyPtr(in.GetMatchAddress())
	return out
}
func PscAutomatedEndpoints_ToProto(mapCtx *direct.MapContext, in *krm.PscAutomatedEndpoints) *pb.PscAutomatedEndpoints {
	if in == nil {
		return nil
	}
	out := &pb.PscAutomatedEndpoints{}
	out.ProjectId = direct.ValueOf(in.ProjectID)
	out.Network = direct.ValueOf(in.Network)
	out.MatchAddress = direct.ValueOf(in.MatchAddress)
	return out
}
func ReservationAffinity_FromProto(mapCtx *direct.MapContext, in *pb.ReservationAffinity) *krm.ReservationAffinity {
	if in == nil {
		return nil
	}
	out := &krm.ReservationAffinity{}
	out.ReservationAffinityType = direct.Enum_FromProto(mapCtx, in.GetReservationAffinityType())
	out.Key = direct.LazyPtr(in.GetKey())
	out.Values = in.Values
	return out
}
func ReservationAffinity_ToProto(mapCtx *direct.MapContext, in *krm.ReservationAffinity) *pb.ReservationAffinity {
	if in == nil {
		return nil
	}
	out := &pb.ReservationAffinity{}
	out.ReservationAffinityType = direct.Enum_ToProto[pb.ReservationAffinity_Type](mapCtx, in.ReservationAffinityType)
	out.Key = direct.ValueOf(in.Key)
	out.Values = in.Values
	return out
}
