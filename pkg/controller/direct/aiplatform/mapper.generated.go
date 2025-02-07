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
func AiplatformMetadataSchemaObservedState_FromProto(mapCtx *direct.MapContext, in *pb.MetadataSchema) *krm.AiplatformMetadataSchemaObservedState {
	if in == nil {
		return nil
	}
	out := &krm.AiplatformMetadataSchemaObservedState{}
	// MISSING: Name
	// MISSING: SchemaVersion
	// MISSING: Schema
	// MISSING: SchemaType
	// MISSING: CreateTime
	// MISSING: Description
	return out
}
func AiplatformMetadataSchemaObservedState_ToProto(mapCtx *direct.MapContext, in *krm.AiplatformMetadataSchemaObservedState) *pb.MetadataSchema {
	if in == nil {
		return nil
	}
	out := &pb.MetadataSchema{}
	// MISSING: Name
	// MISSING: SchemaVersion
	// MISSING: Schema
	// MISSING: SchemaType
	// MISSING: CreateTime
	// MISSING: Description
	return out
}
func AiplatformMetadataSchemaSpec_FromProto(mapCtx *direct.MapContext, in *pb.MetadataSchema) *krm.AiplatformMetadataSchemaSpec {
	if in == nil {
		return nil
	}
	out := &krm.AiplatformMetadataSchemaSpec{}
	// MISSING: Name
	// MISSING: SchemaVersion
	// MISSING: Schema
	// MISSING: SchemaType
	// MISSING: CreateTime
	// MISSING: Description
	return out
}
func AiplatformMetadataSchemaSpec_ToProto(mapCtx *direct.MapContext, in *krm.AiplatformMetadataSchemaSpec) *pb.MetadataSchema {
	if in == nil {
		return nil
	}
	out := &pb.MetadataSchema{}
	// MISSING: Name
	// MISSING: SchemaVersion
	// MISSING: Schema
	// MISSING: SchemaType
	// MISSING: CreateTime
	// MISSING: Description
	return out
}
func AiplatformMetadataStoreObservedState_FromProto(mapCtx *direct.MapContext, in *pb.MetadataStore) *krm.AiplatformMetadataStoreObservedState {
	if in == nil {
		return nil
	}
	out := &krm.AiplatformMetadataStoreObservedState{}
	// MISSING: Name
	// MISSING: CreateTime
	// MISSING: UpdateTime
	// MISSING: EncryptionSpec
	// MISSING: Description
	// MISSING: State
	// MISSING: DataplexConfig
	return out
}
func AiplatformMetadataStoreObservedState_ToProto(mapCtx *direct.MapContext, in *krm.AiplatformMetadataStoreObservedState) *pb.MetadataStore {
	if in == nil {
		return nil
	}
	out := &pb.MetadataStore{}
	// MISSING: Name
	// MISSING: CreateTime
	// MISSING: UpdateTime
	// MISSING: EncryptionSpec
	// MISSING: Description
	// MISSING: State
	// MISSING: DataplexConfig
	return out
}
func AiplatformMetadataStoreSpec_FromProto(mapCtx *direct.MapContext, in *pb.MetadataStore) *krm.AiplatformMetadataStoreSpec {
	if in == nil {
		return nil
	}
	out := &krm.AiplatformMetadataStoreSpec{}
	// MISSING: Name
	// MISSING: CreateTime
	// MISSING: UpdateTime
	// MISSING: EncryptionSpec
	// MISSING: Description
	// MISSING: State
	// MISSING: DataplexConfig
	return out
}
func AiplatformMetadataStoreSpec_ToProto(mapCtx *direct.MapContext, in *krm.AiplatformMetadataStoreSpec) *pb.MetadataStore {
	if in == nil {
		return nil
	}
	out := &pb.MetadataStore{}
	// MISSING: Name
	// MISSING: CreateTime
	// MISSING: UpdateTime
	// MISSING: EncryptionSpec
	// MISSING: Description
	// MISSING: State
	// MISSING: DataplexConfig
	return out
}
func AiplatformModelDeploymentMonitoringJobObservedState_FromProto(mapCtx *direct.MapContext, in *pb.ModelDeploymentMonitoringJob) *krm.AiplatformModelDeploymentMonitoringJobObservedState {
	if in == nil {
		return nil
	}
	out := &krm.AiplatformModelDeploymentMonitoringJobObservedState{}
	// MISSING: Name
	// MISSING: DisplayName
	// MISSING: Endpoint
	// MISSING: State
	// MISSING: ScheduleState
	// MISSING: LatestMonitoringPipelineMetadata
	// MISSING: ModelDeploymentMonitoringObjectiveConfigs
	// MISSING: ModelDeploymentMonitoringScheduleConfig
	// MISSING: LoggingSamplingStrategy
	// MISSING: ModelMonitoringAlertConfig
	// MISSING: PredictInstanceSchemaURI
	// MISSING: SamplePredictInstance
	// MISSING: AnalysisInstanceSchemaURI
	// MISSING: BigqueryTables
	// MISSING: LogTtl
	// MISSING: Labels
	// MISSING: CreateTime
	// MISSING: UpdateTime
	// MISSING: NextScheduleTime
	// MISSING: StatsAnomaliesBaseDirectory
	// MISSING: EncryptionSpec
	// MISSING: EnableMonitoringPipelineLogs
	// MISSING: Error
	// MISSING: SatisfiesPzs
	// MISSING: SatisfiesPzi
	return out
}
func AiplatformModelDeploymentMonitoringJobObservedState_ToProto(mapCtx *direct.MapContext, in *krm.AiplatformModelDeploymentMonitoringJobObservedState) *pb.ModelDeploymentMonitoringJob {
	if in == nil {
		return nil
	}
	out := &pb.ModelDeploymentMonitoringJob{}
	// MISSING: Name
	// MISSING: DisplayName
	// MISSING: Endpoint
	// MISSING: State
	// MISSING: ScheduleState
	// MISSING: LatestMonitoringPipelineMetadata
	// MISSING: ModelDeploymentMonitoringObjectiveConfigs
	// MISSING: ModelDeploymentMonitoringScheduleConfig
	// MISSING: LoggingSamplingStrategy
	// MISSING: ModelMonitoringAlertConfig
	// MISSING: PredictInstanceSchemaURI
	// MISSING: SamplePredictInstance
	// MISSING: AnalysisInstanceSchemaURI
	// MISSING: BigqueryTables
	// MISSING: LogTtl
	// MISSING: Labels
	// MISSING: CreateTime
	// MISSING: UpdateTime
	// MISSING: NextScheduleTime
	// MISSING: StatsAnomaliesBaseDirectory
	// MISSING: EncryptionSpec
	// MISSING: EnableMonitoringPipelineLogs
	// MISSING: Error
	// MISSING: SatisfiesPzs
	// MISSING: SatisfiesPzi
	return out
}
func AiplatformModelDeploymentMonitoringJobSpec_FromProto(mapCtx *direct.MapContext, in *pb.ModelDeploymentMonitoringJob) *krm.AiplatformModelDeploymentMonitoringJobSpec {
	if in == nil {
		return nil
	}
	out := &krm.AiplatformModelDeploymentMonitoringJobSpec{}
	// MISSING: Name
	// MISSING: DisplayName
	// MISSING: Endpoint
	// MISSING: State
	// MISSING: ScheduleState
	// MISSING: LatestMonitoringPipelineMetadata
	// MISSING: ModelDeploymentMonitoringObjectiveConfigs
	// MISSING: ModelDeploymentMonitoringScheduleConfig
	// MISSING: LoggingSamplingStrategy
	// MISSING: ModelMonitoringAlertConfig
	// MISSING: PredictInstanceSchemaURI
	// MISSING: SamplePredictInstance
	// MISSING: AnalysisInstanceSchemaURI
	// MISSING: BigqueryTables
	// MISSING: LogTtl
	// MISSING: Labels
	// MISSING: CreateTime
	// MISSING: UpdateTime
	// MISSING: NextScheduleTime
	// MISSING: StatsAnomaliesBaseDirectory
	// MISSING: EncryptionSpec
	// MISSING: EnableMonitoringPipelineLogs
	// MISSING: Error
	// MISSING: SatisfiesPzs
	// MISSING: SatisfiesPzi
	return out
}
func AiplatformModelDeploymentMonitoringJobSpec_ToProto(mapCtx *direct.MapContext, in *krm.AiplatformModelDeploymentMonitoringJobSpec) *pb.ModelDeploymentMonitoringJob {
	if in == nil {
		return nil
	}
	out := &pb.ModelDeploymentMonitoringJob{}
	// MISSING: Name
	// MISSING: DisplayName
	// MISSING: Endpoint
	// MISSING: State
	// MISSING: ScheduleState
	// MISSING: LatestMonitoringPipelineMetadata
	// MISSING: ModelDeploymentMonitoringObjectiveConfigs
	// MISSING: ModelDeploymentMonitoringScheduleConfig
	// MISSING: LoggingSamplingStrategy
	// MISSING: ModelMonitoringAlertConfig
	// MISSING: PredictInstanceSchemaURI
	// MISSING: SamplePredictInstance
	// MISSING: AnalysisInstanceSchemaURI
	// MISSING: BigqueryTables
	// MISSING: LogTtl
	// MISSING: Labels
	// MISSING: CreateTime
	// MISSING: UpdateTime
	// MISSING: NextScheduleTime
	// MISSING: StatsAnomaliesBaseDirectory
	// MISSING: EncryptionSpec
	// MISSING: EnableMonitoringPipelineLogs
	// MISSING: Error
	// MISSING: SatisfiesPzs
	// MISSING: SatisfiesPzi
	return out
}
func AiplatformModelObservedState_FromProto(mapCtx *direct.MapContext, in *pb.Model) *krm.AiplatformModelObservedState {
	if in == nil {
		return nil
	}
	out := &krm.AiplatformModelObservedState{}
	// MISSING: Name
	// MISSING: VersionID
	// MISSING: VersionAliases
	// MISSING: VersionCreateTime
	// MISSING: VersionUpdateTime
	// MISSING: DisplayName
	// MISSING: Description
	// MISSING: VersionDescription
	// MISSING: PredictSchemata
	// MISSING: MetadataSchemaURI
	// MISSING: Metadata
	// MISSING: SupportedExportFormats
	// MISSING: TrainingPipeline
	// MISSING: PipelineJob
	// MISSING: ContainerSpec
	// MISSING: ArtifactURI
	// MISSING: SupportedDeploymentResourcesTypes
	// MISSING: SupportedInputStorageFormats
	// MISSING: SupportedOutputStorageFormats
	// MISSING: CreateTime
	// MISSING: UpdateTime
	// MISSING: DeployedModels
	// MISSING: ExplanationSpec
	// MISSING: Etag
	// MISSING: Labels
	// MISSING: DataStats
	// MISSING: EncryptionSpec
	// MISSING: ModelSourceInfo
	// MISSING: OriginalModelInfo
	// MISSING: MetadataArtifact
	// MISSING: BaseModelSource
	// MISSING: SatisfiesPzs
	// MISSING: SatisfiesPzi
	return out
}
func AiplatformModelObservedState_ToProto(mapCtx *direct.MapContext, in *krm.AiplatformModelObservedState) *pb.Model {
	if in == nil {
		return nil
	}
	out := &pb.Model{}
	// MISSING: Name
	// MISSING: VersionID
	// MISSING: VersionAliases
	// MISSING: VersionCreateTime
	// MISSING: VersionUpdateTime
	// MISSING: DisplayName
	// MISSING: Description
	// MISSING: VersionDescription
	// MISSING: PredictSchemata
	// MISSING: MetadataSchemaURI
	// MISSING: Metadata
	// MISSING: SupportedExportFormats
	// MISSING: TrainingPipeline
	// MISSING: PipelineJob
	// MISSING: ContainerSpec
	// MISSING: ArtifactURI
	// MISSING: SupportedDeploymentResourcesTypes
	// MISSING: SupportedInputStorageFormats
	// MISSING: SupportedOutputStorageFormats
	// MISSING: CreateTime
	// MISSING: UpdateTime
	// MISSING: DeployedModels
	// MISSING: ExplanationSpec
	// MISSING: Etag
	// MISSING: Labels
	// MISSING: DataStats
	// MISSING: EncryptionSpec
	// MISSING: ModelSourceInfo
	// MISSING: OriginalModelInfo
	// MISSING: MetadataArtifact
	// MISSING: BaseModelSource
	// MISSING: SatisfiesPzs
	// MISSING: SatisfiesPzi
	return out
}
func AiplatformModelSpec_FromProto(mapCtx *direct.MapContext, in *pb.Model) *krm.AiplatformModelSpec {
	if in == nil {
		return nil
	}
	out := &krm.AiplatformModelSpec{}
	// MISSING: Name
	// MISSING: VersionID
	// MISSING: VersionAliases
	// MISSING: VersionCreateTime
	// MISSING: VersionUpdateTime
	// MISSING: DisplayName
	// MISSING: Description
	// MISSING: VersionDescription
	// MISSING: PredictSchemata
	// MISSING: MetadataSchemaURI
	// MISSING: Metadata
	// MISSING: SupportedExportFormats
	// MISSING: TrainingPipeline
	// MISSING: PipelineJob
	// MISSING: ContainerSpec
	// MISSING: ArtifactURI
	// MISSING: SupportedDeploymentResourcesTypes
	// MISSING: SupportedInputStorageFormats
	// MISSING: SupportedOutputStorageFormats
	// MISSING: CreateTime
	// MISSING: UpdateTime
	// MISSING: DeployedModels
	// MISSING: ExplanationSpec
	// MISSING: Etag
	// MISSING: Labels
	// MISSING: DataStats
	// MISSING: EncryptionSpec
	// MISSING: ModelSourceInfo
	// MISSING: OriginalModelInfo
	// MISSING: MetadataArtifact
	// MISSING: BaseModelSource
	// MISSING: SatisfiesPzs
	// MISSING: SatisfiesPzi
	return out
}
func AiplatformModelSpec_ToProto(mapCtx *direct.MapContext, in *krm.AiplatformModelSpec) *pb.Model {
	if in == nil {
		return nil
	}
	out := &pb.Model{}
	// MISSING: Name
	// MISSING: VersionID
	// MISSING: VersionAliases
	// MISSING: VersionCreateTime
	// MISSING: VersionUpdateTime
	// MISSING: DisplayName
	// MISSING: Description
	// MISSING: VersionDescription
	// MISSING: PredictSchemata
	// MISSING: MetadataSchemaURI
	// MISSING: Metadata
	// MISSING: SupportedExportFormats
	// MISSING: TrainingPipeline
	// MISSING: PipelineJob
	// MISSING: ContainerSpec
	// MISSING: ArtifactURI
	// MISSING: SupportedDeploymentResourcesTypes
	// MISSING: SupportedInputStorageFormats
	// MISSING: SupportedOutputStorageFormats
	// MISSING: CreateTime
	// MISSING: UpdateTime
	// MISSING: DeployedModels
	// MISSING: ExplanationSpec
	// MISSING: Etag
	// MISSING: Labels
	// MISSING: DataStats
	// MISSING: EncryptionSpec
	// MISSING: ModelSourceInfo
	// MISSING: OriginalModelInfo
	// MISSING: MetadataArtifact
	// MISSING: BaseModelSource
	// MISSING: SatisfiesPzs
	// MISSING: SatisfiesPzi
	return out
}
func BigQueryDestination_FromProto(mapCtx *direct.MapContext, in *pb.BigQueryDestination) *krm.BigQueryDestination {
	if in == nil {
		return nil
	}
	out := &krm.BigQueryDestination{}
	out.OutputURI = direct.LazyPtr(in.GetOutputUri())
	return out
}
func BigQueryDestination_ToProto(mapCtx *direct.MapContext, in *krm.BigQueryDestination) *pb.BigQueryDestination {
	if in == nil {
		return nil
	}
	out := &pb.BigQueryDestination{}
	out.OutputUri = direct.ValueOf(in.OutputURI)
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
func GcsDestination_FromProto(mapCtx *direct.MapContext, in *pb.GcsDestination) *krm.GcsDestination {
	if in == nil {
		return nil
	}
	out := &krm.GcsDestination{}
	out.OutputURIPrefix = direct.LazyPtr(in.GetOutputUriPrefix())
	return out
}
func GcsDestination_ToProto(mapCtx *direct.MapContext, in *krm.GcsDestination) *pb.GcsDestination {
	if in == nil {
		return nil
	}
	out := &pb.GcsDestination{}
	out.OutputUriPrefix = direct.ValueOf(in.OutputURIPrefix)
	return out
}
func GcsSource_FromProto(mapCtx *direct.MapContext, in *pb.GcsSource) *krm.GcsSource {
	if in == nil {
		return nil
	}
	out := &krm.GcsSource{}
	out.Uris = in.Uris
	return out
}
func GcsSource_ToProto(mapCtx *direct.MapContext, in *krm.GcsSource) *pb.GcsSource {
	if in == nil {
		return nil
	}
	out := &pb.GcsSource{}
	out.Uris = in.Uris
	return out
}
func ModelDeploymentMonitoringBigQueryTable_FromProto(mapCtx *direct.MapContext, in *pb.ModelDeploymentMonitoringBigQueryTable) *krm.ModelDeploymentMonitoringBigQueryTable {
	if in == nil {
		return nil
	}
	out := &krm.ModelDeploymentMonitoringBigQueryTable{}
	out.LogSource = direct.Enum_FromProto(mapCtx, in.GetLogSource())
	out.LogType = direct.Enum_FromProto(mapCtx, in.GetLogType())
	out.BigqueryTablePath = direct.LazyPtr(in.GetBigqueryTablePath())
	// MISSING: RequestResponseLoggingSchemaVersion
	return out
}
func ModelDeploymentMonitoringBigQueryTable_ToProto(mapCtx *direct.MapContext, in *krm.ModelDeploymentMonitoringBigQueryTable) *pb.ModelDeploymentMonitoringBigQueryTable {
	if in == nil {
		return nil
	}
	out := &pb.ModelDeploymentMonitoringBigQueryTable{}
	out.LogSource = direct.Enum_ToProto[pb.ModelDeploymentMonitoringBigQueryTable_LogSource](mapCtx, in.LogSource)
	out.LogType = direct.Enum_ToProto[pb.ModelDeploymentMonitoringBigQueryTable_LogType](mapCtx, in.LogType)
	out.BigqueryTablePath = direct.ValueOf(in.BigqueryTablePath)
	// MISSING: RequestResponseLoggingSchemaVersion
	return out
}
func ModelDeploymentMonitoringBigQueryTableObservedState_FromProto(mapCtx *direct.MapContext, in *pb.ModelDeploymentMonitoringBigQueryTable) *krm.ModelDeploymentMonitoringBigQueryTableObservedState {
	if in == nil {
		return nil
	}
	out := &krm.ModelDeploymentMonitoringBigQueryTableObservedState{}
	// MISSING: LogSource
	// MISSING: LogType
	// MISSING: BigqueryTablePath
	out.RequestResponseLoggingSchemaVersion = direct.LazyPtr(in.GetRequestResponseLoggingSchemaVersion())
	return out
}
func ModelDeploymentMonitoringBigQueryTableObservedState_ToProto(mapCtx *direct.MapContext, in *krm.ModelDeploymentMonitoringBigQueryTableObservedState) *pb.ModelDeploymentMonitoringBigQueryTable {
	if in == nil {
		return nil
	}
	out := &pb.ModelDeploymentMonitoringBigQueryTable{}
	// MISSING: LogSource
	// MISSING: LogType
	// MISSING: BigqueryTablePath
	out.RequestResponseLoggingSchemaVersion = direct.ValueOf(in.RequestResponseLoggingSchemaVersion)
	return out
}
func ModelDeploymentMonitoringJob_FromProto(mapCtx *direct.MapContext, in *pb.ModelDeploymentMonitoringJob) *krm.ModelDeploymentMonitoringJob {
	if in == nil {
		return nil
	}
	out := &krm.ModelDeploymentMonitoringJob{}
	// MISSING: Name
	out.DisplayName = direct.LazyPtr(in.GetDisplayName())
	out.Endpoint = direct.LazyPtr(in.GetEndpoint())
	// MISSING: State
	// MISSING: ScheduleState
	// MISSING: LatestMonitoringPipelineMetadata
	out.ModelDeploymentMonitoringObjectiveConfigs = direct.Slice_FromProto(mapCtx, in.ModelDeploymentMonitoringObjectiveConfigs, ModelDeploymentMonitoringObjectiveConfig_FromProto)
	out.ModelDeploymentMonitoringScheduleConfig = ModelDeploymentMonitoringScheduleConfig_FromProto(mapCtx, in.GetModelDeploymentMonitoringScheduleConfig())
	out.LoggingSamplingStrategy = SamplingStrategy_FromProto(mapCtx, in.GetLoggingSamplingStrategy())
	out.ModelMonitoringAlertConfig = ModelMonitoringAlertConfig_FromProto(mapCtx, in.GetModelMonitoringAlertConfig())
	out.PredictInstanceSchemaURI = direct.LazyPtr(in.GetPredictInstanceSchemaUri())
	out.SamplePredictInstance = Value_FromProto(mapCtx, in.GetSamplePredictInstance())
	out.AnalysisInstanceSchemaURI = direct.LazyPtr(in.GetAnalysisInstanceSchemaUri())
	// MISSING: BigqueryTables
	out.LogTtl = direct.StringDuration_FromProto(mapCtx, in.GetLogTtl())
	out.Labels = in.Labels
	// MISSING: CreateTime
	// MISSING: UpdateTime
	// MISSING: NextScheduleTime
	out.StatsAnomaliesBaseDirectory = GcsDestination_FromProto(mapCtx, in.GetStatsAnomaliesBaseDirectory())
	out.EncryptionSpec = EncryptionSpec_FromProto(mapCtx, in.GetEncryptionSpec())
	out.EnableMonitoringPipelineLogs = direct.LazyPtr(in.GetEnableMonitoringPipelineLogs())
	// MISSING: Error
	// MISSING: SatisfiesPzs
	// MISSING: SatisfiesPzi
	return out
}
func ModelDeploymentMonitoringJob_ToProto(mapCtx *direct.MapContext, in *krm.ModelDeploymentMonitoringJob) *pb.ModelDeploymentMonitoringJob {
	if in == nil {
		return nil
	}
	out := &pb.ModelDeploymentMonitoringJob{}
	// MISSING: Name
	out.DisplayName = direct.ValueOf(in.DisplayName)
	out.Endpoint = direct.ValueOf(in.Endpoint)
	// MISSING: State
	// MISSING: ScheduleState
	// MISSING: LatestMonitoringPipelineMetadata
	out.ModelDeploymentMonitoringObjectiveConfigs = direct.Slice_ToProto(mapCtx, in.ModelDeploymentMonitoringObjectiveConfigs, ModelDeploymentMonitoringObjectiveConfig_ToProto)
	out.ModelDeploymentMonitoringScheduleConfig = ModelDeploymentMonitoringScheduleConfig_ToProto(mapCtx, in.ModelDeploymentMonitoringScheduleConfig)
	out.LoggingSamplingStrategy = SamplingStrategy_ToProto(mapCtx, in.LoggingSamplingStrategy)
	out.ModelMonitoringAlertConfig = ModelMonitoringAlertConfig_ToProto(mapCtx, in.ModelMonitoringAlertConfig)
	out.PredictInstanceSchemaUri = direct.ValueOf(in.PredictInstanceSchemaURI)
	out.SamplePredictInstance = Value_ToProto(mapCtx, in.SamplePredictInstance)
	out.AnalysisInstanceSchemaUri = direct.ValueOf(in.AnalysisInstanceSchemaURI)
	// MISSING: BigqueryTables
	out.LogTtl = direct.StringDuration_ToProto(mapCtx, in.LogTtl)
	out.Labels = in.Labels
	// MISSING: CreateTime
	// MISSING: UpdateTime
	// MISSING: NextScheduleTime
	out.StatsAnomaliesBaseDirectory = GcsDestination_ToProto(mapCtx, in.StatsAnomaliesBaseDirectory)
	out.EncryptionSpec = EncryptionSpec_ToProto(mapCtx, in.EncryptionSpec)
	out.EnableMonitoringPipelineLogs = direct.ValueOf(in.EnableMonitoringPipelineLogs)
	// MISSING: Error
	// MISSING: SatisfiesPzs
	// MISSING: SatisfiesPzi
	return out
}
func ModelDeploymentMonitoringJobObservedState_FromProto(mapCtx *direct.MapContext, in *pb.ModelDeploymentMonitoringJob) *krm.ModelDeploymentMonitoringJobObservedState {
	if in == nil {
		return nil
	}
	out := &krm.ModelDeploymentMonitoringJobObservedState{}
	out.Name = direct.LazyPtr(in.GetName())
	// MISSING: DisplayName
	// MISSING: Endpoint
	out.State = direct.Enum_FromProto(mapCtx, in.GetState())
	out.ScheduleState = direct.Enum_FromProto(mapCtx, in.GetScheduleState())
	out.LatestMonitoringPipelineMetadata = ModelDeploymentMonitoringJob_LatestMonitoringPipelineMetadata_FromProto(mapCtx, in.GetLatestMonitoringPipelineMetadata())
	// MISSING: ModelDeploymentMonitoringObjectiveConfigs
	// MISSING: ModelDeploymentMonitoringScheduleConfig
	// MISSING: LoggingSamplingStrategy
	// MISSING: ModelMonitoringAlertConfig
	// MISSING: PredictInstanceSchemaURI
	// MISSING: SamplePredictInstance
	// MISSING: AnalysisInstanceSchemaURI
	out.BigqueryTables = direct.Slice_FromProto(mapCtx, in.BigqueryTables, ModelDeploymentMonitoringBigQueryTable_FromProto)
	// MISSING: LogTtl
	// MISSING: Labels
	out.CreateTime = direct.StringTimestamp_FromProto(mapCtx, in.GetCreateTime())
	out.UpdateTime = direct.StringTimestamp_FromProto(mapCtx, in.GetUpdateTime())
	out.NextScheduleTime = direct.StringTimestamp_FromProto(mapCtx, in.GetNextScheduleTime())
	// MISSING: StatsAnomaliesBaseDirectory
	// MISSING: EncryptionSpec
	// MISSING: EnableMonitoringPipelineLogs
	out.Error = Status_FromProto(mapCtx, in.GetError())
	out.SatisfiesPzs = direct.LazyPtr(in.GetSatisfiesPzs())
	out.SatisfiesPzi = direct.LazyPtr(in.GetSatisfiesPzi())
	return out
}
func ModelDeploymentMonitoringJobObservedState_ToProto(mapCtx *direct.MapContext, in *krm.ModelDeploymentMonitoringJobObservedState) *pb.ModelDeploymentMonitoringJob {
	if in == nil {
		return nil
	}
	out := &pb.ModelDeploymentMonitoringJob{}
	out.Name = direct.ValueOf(in.Name)
	// MISSING: DisplayName
	// MISSING: Endpoint
	out.State = direct.Enum_ToProto[pb.JobState](mapCtx, in.State)
	out.ScheduleState = direct.Enum_ToProto[pb.ModelDeploymentMonitoringJob_MonitoringScheduleState](mapCtx, in.ScheduleState)
	out.LatestMonitoringPipelineMetadata = ModelDeploymentMonitoringJob_LatestMonitoringPipelineMetadata_ToProto(mapCtx, in.LatestMonitoringPipelineMetadata)
	// MISSING: ModelDeploymentMonitoringObjectiveConfigs
	// MISSING: ModelDeploymentMonitoringScheduleConfig
	// MISSING: LoggingSamplingStrategy
	// MISSING: ModelMonitoringAlertConfig
	// MISSING: PredictInstanceSchemaURI
	// MISSING: SamplePredictInstance
	// MISSING: AnalysisInstanceSchemaURI
	out.BigqueryTables = direct.Slice_ToProto(mapCtx, in.BigqueryTables, ModelDeploymentMonitoringBigQueryTable_ToProto)
	// MISSING: LogTtl
	// MISSING: Labels
	out.CreateTime = direct.StringTimestamp_ToProto(mapCtx, in.CreateTime)
	out.UpdateTime = direct.StringTimestamp_ToProto(mapCtx, in.UpdateTime)
	out.NextScheduleTime = direct.StringTimestamp_ToProto(mapCtx, in.NextScheduleTime)
	// MISSING: StatsAnomaliesBaseDirectory
	// MISSING: EncryptionSpec
	// MISSING: EnableMonitoringPipelineLogs
	out.Error = Status_ToProto(mapCtx, in.Error)
	out.SatisfiesPzs = direct.ValueOf(in.SatisfiesPzs)
	out.SatisfiesPzi = direct.ValueOf(in.SatisfiesPzi)
	return out
}
func ModelDeploymentMonitoringObjectiveConfig_FromProto(mapCtx *direct.MapContext, in *pb.ModelDeploymentMonitoringObjectiveConfig) *krm.ModelDeploymentMonitoringObjectiveConfig {
	if in == nil {
		return nil
	}
	out := &krm.ModelDeploymentMonitoringObjectiveConfig{}
	out.DeployedModelID = direct.LazyPtr(in.GetDeployedModelId())
	out.ObjectiveConfig = ModelMonitoringObjectiveConfig_FromProto(mapCtx, in.GetObjectiveConfig())
	return out
}
func ModelDeploymentMonitoringObjectiveConfig_ToProto(mapCtx *direct.MapContext, in *krm.ModelDeploymentMonitoringObjectiveConfig) *pb.ModelDeploymentMonitoringObjectiveConfig {
	if in == nil {
		return nil
	}
	out := &pb.ModelDeploymentMonitoringObjectiveConfig{}
	out.DeployedModelId = direct.ValueOf(in.DeployedModelID)
	out.ObjectiveConfig = ModelMonitoringObjectiveConfig_ToProto(mapCtx, in.ObjectiveConfig)
	return out
}
func ModelDeploymentMonitoringScheduleConfig_FromProto(mapCtx *direct.MapContext, in *pb.ModelDeploymentMonitoringScheduleConfig) *krm.ModelDeploymentMonitoringScheduleConfig {
	if in == nil {
		return nil
	}
	out := &krm.ModelDeploymentMonitoringScheduleConfig{}
	out.MonitorInterval = direct.StringDuration_FromProto(mapCtx, in.GetMonitorInterval())
	out.MonitorWindow = direct.StringDuration_FromProto(mapCtx, in.GetMonitorWindow())
	return out
}
func ModelDeploymentMonitoringScheduleConfig_ToProto(mapCtx *direct.MapContext, in *krm.ModelDeploymentMonitoringScheduleConfig) *pb.ModelDeploymentMonitoringScheduleConfig {
	if in == nil {
		return nil
	}
	out := &pb.ModelDeploymentMonitoringScheduleConfig{}
	out.MonitorInterval = direct.StringDuration_ToProto(mapCtx, in.MonitorInterval)
	out.MonitorWindow = direct.StringDuration_ToProto(mapCtx, in.MonitorWindow)
	return out
}
func ModelMonitoringAlertConfig_FromProto(mapCtx *direct.MapContext, in *pb.ModelMonitoringAlertConfig) *krm.ModelMonitoringAlertConfig {
	if in == nil {
		return nil
	}
	out := &krm.ModelMonitoringAlertConfig{}
	out.EmailAlertConfig = ModelMonitoringAlertConfig_EmailAlertConfig_FromProto(mapCtx, in.GetEmailAlertConfig())
	out.EnableLogging = direct.LazyPtr(in.GetEnableLogging())
	out.NotificationChannels = in.NotificationChannels
	return out
}
func ModelMonitoringAlertConfig_ToProto(mapCtx *direct.MapContext, in *krm.ModelMonitoringAlertConfig) *pb.ModelMonitoringAlertConfig {
	if in == nil {
		return nil
	}
	out := &pb.ModelMonitoringAlertConfig{}
	if oneof := ModelMonitoringAlertConfig_EmailAlertConfig_ToProto(mapCtx, in.EmailAlertConfig); oneof != nil {
		out.Alert = &pb.ModelMonitoringAlertConfig_EmailAlertConfig_{EmailAlertConfig: oneof}
	}
	out.EnableLogging = direct.ValueOf(in.EnableLogging)
	out.NotificationChannels = in.NotificationChannels
	return out
}
func ModelMonitoringAlertConfig_EmailAlertConfig_FromProto(mapCtx *direct.MapContext, in *pb.ModelMonitoringAlertConfig_EmailAlertConfig) *krm.ModelMonitoringAlertConfig_EmailAlertConfig {
	if in == nil {
		return nil
	}
	out := &krm.ModelMonitoringAlertConfig_EmailAlertConfig{}
	out.UserEmails = in.UserEmails
	return out
}
func ModelMonitoringAlertConfig_EmailAlertConfig_ToProto(mapCtx *direct.MapContext, in *krm.ModelMonitoringAlertConfig_EmailAlertConfig) *pb.ModelMonitoringAlertConfig_EmailAlertConfig {
	if in == nil {
		return nil
	}
	out := &pb.ModelMonitoringAlertConfig_EmailAlertConfig{}
	out.UserEmails = in.UserEmails
	return out
}
func ModelMonitoringObjectiveConfig_FromProto(mapCtx *direct.MapContext, in *pb.ModelMonitoringObjectiveConfig) *krm.ModelMonitoringObjectiveConfig {
	if in == nil {
		return nil
	}
	out := &krm.ModelMonitoringObjectiveConfig{}
	out.TrainingDataset = ModelMonitoringObjectiveConfig_TrainingDataset_FromProto(mapCtx, in.GetTrainingDataset())
	out.TrainingPredictionSkewDetectionConfig = ModelMonitoringObjectiveConfig_TrainingPredictionSkewDetectionConfig_FromProto(mapCtx, in.GetTrainingPredictionSkewDetectionConfig())
	out.PredictionDriftDetectionConfig = ModelMonitoringObjectiveConfig_PredictionDriftDetectionConfig_FromProto(mapCtx, in.GetPredictionDriftDetectionConfig())
	out.ExplanationConfig = ModelMonitoringObjectiveConfig_ExplanationConfig_FromProto(mapCtx, in.GetExplanationConfig())
	return out
}
func ModelMonitoringObjectiveConfig_ToProto(mapCtx *direct.MapContext, in *krm.ModelMonitoringObjectiveConfig) *pb.ModelMonitoringObjectiveConfig {
	if in == nil {
		return nil
	}
	out := &pb.ModelMonitoringObjectiveConfig{}
	out.TrainingDataset = ModelMonitoringObjectiveConfig_TrainingDataset_ToProto(mapCtx, in.TrainingDataset)
	out.TrainingPredictionSkewDetectionConfig = ModelMonitoringObjectiveConfig_TrainingPredictionSkewDetectionConfig_ToProto(mapCtx, in.TrainingPredictionSkewDetectionConfig)
	out.PredictionDriftDetectionConfig = ModelMonitoringObjectiveConfig_PredictionDriftDetectionConfig_ToProto(mapCtx, in.PredictionDriftDetectionConfig)
	out.ExplanationConfig = ModelMonitoringObjectiveConfig_ExplanationConfig_ToProto(mapCtx, in.ExplanationConfig)
	return out
}
func ModelMonitoringObjectiveConfig_ExplanationConfig_FromProto(mapCtx *direct.MapContext, in *pb.ModelMonitoringObjectiveConfig_ExplanationConfig) *krm.ModelMonitoringObjectiveConfig_ExplanationConfig {
	if in == nil {
		return nil
	}
	out := &krm.ModelMonitoringObjectiveConfig_ExplanationConfig{}
	out.EnableFeatureAttributes = direct.LazyPtr(in.GetEnableFeatureAttributes())
	out.ExplanationBaseline = ModelMonitoringObjectiveConfig_ExplanationConfig_ExplanationBaseline_FromProto(mapCtx, in.GetExplanationBaseline())
	return out
}
func ModelMonitoringObjectiveConfig_ExplanationConfig_ToProto(mapCtx *direct.MapContext, in *krm.ModelMonitoringObjectiveConfig_ExplanationConfig) *pb.ModelMonitoringObjectiveConfig_ExplanationConfig {
	if in == nil {
		return nil
	}
	out := &pb.ModelMonitoringObjectiveConfig_ExplanationConfig{}
	out.EnableFeatureAttributes = direct.ValueOf(in.EnableFeatureAttributes)
	out.ExplanationBaseline = ModelMonitoringObjectiveConfig_ExplanationConfig_ExplanationBaseline_ToProto(mapCtx, in.ExplanationBaseline)
	return out
}
func ModelMonitoringObjectiveConfig_ExplanationConfig_ExplanationBaseline_FromProto(mapCtx *direct.MapContext, in *pb.ModelMonitoringObjectiveConfig_ExplanationConfig_ExplanationBaseline) *krm.ModelMonitoringObjectiveConfig_ExplanationConfig_ExplanationBaseline {
	if in == nil {
		return nil
	}
	out := &krm.ModelMonitoringObjectiveConfig_ExplanationConfig_ExplanationBaseline{}
	out.Gcs = GcsDestination_FromProto(mapCtx, in.GetGcs())
	out.Bigquery = BigQueryDestination_FromProto(mapCtx, in.GetBigquery())
	out.PredictionFormat = direct.Enum_FromProto(mapCtx, in.GetPredictionFormat())
	return out
}
func ModelMonitoringObjectiveConfig_ExplanationConfig_ExplanationBaseline_ToProto(mapCtx *direct.MapContext, in *krm.ModelMonitoringObjectiveConfig_ExplanationConfig_ExplanationBaseline) *pb.ModelMonitoringObjectiveConfig_ExplanationConfig_ExplanationBaseline {
	if in == nil {
		return nil
	}
	out := &pb.ModelMonitoringObjectiveConfig_ExplanationConfig_ExplanationBaseline{}
	if oneof := GcsDestination_ToProto(mapCtx, in.Gcs); oneof != nil {
		out.Destination = &pb.ModelMonitoringObjectiveConfig_ExplanationConfig_ExplanationBaseline_Gcs{Gcs: oneof}
	}
	if oneof := BigQueryDestination_ToProto(mapCtx, in.Bigquery); oneof != nil {
		out.Destination = &pb.ModelMonitoringObjectiveConfig_ExplanationConfig_ExplanationBaseline_Bigquery{Bigquery: oneof}
	}
	out.PredictionFormat = direct.Enum_ToProto[pb.ModelMonitoringObjectiveConfig_ExplanationConfig_ExplanationBaseline_PredictionFormat](mapCtx, in.PredictionFormat)
	return out
}
func ModelMonitoringObjectiveConfig_PredictionDriftDetectionConfig_FromProto(mapCtx *direct.MapContext, in *pb.ModelMonitoringObjectiveConfig_PredictionDriftDetectionConfig) *krm.ModelMonitoringObjectiveConfig_PredictionDriftDetectionConfig {
	if in == nil {
		return nil
	}
	out := &krm.ModelMonitoringObjectiveConfig_PredictionDriftDetectionConfig{}
	// MISSING: DriftThresholds
	// MISSING: AttributionScoreDriftThresholds
	out.DefaultDriftThreshold = ThresholdConfig_FromProto(mapCtx, in.GetDefaultDriftThreshold())
	return out
}
func ModelMonitoringObjectiveConfig_PredictionDriftDetectionConfig_ToProto(mapCtx *direct.MapContext, in *krm.ModelMonitoringObjectiveConfig_PredictionDriftDetectionConfig) *pb.ModelMonitoringObjectiveConfig_PredictionDriftDetectionConfig {
	if in == nil {
		return nil
	}
	out := &pb.ModelMonitoringObjectiveConfig_PredictionDriftDetectionConfig{}
	// MISSING: DriftThresholds
	// MISSING: AttributionScoreDriftThresholds
	out.DefaultDriftThreshold = ThresholdConfig_ToProto(mapCtx, in.DefaultDriftThreshold)
	return out
}
func ModelMonitoringObjectiveConfig_TrainingDataset_FromProto(mapCtx *direct.MapContext, in *pb.ModelMonitoringObjectiveConfig_TrainingDataset) *krm.ModelMonitoringObjectiveConfig_TrainingDataset {
	if in == nil {
		return nil
	}
	out := &krm.ModelMonitoringObjectiveConfig_TrainingDataset{}
	out.Dataset = direct.LazyPtr(in.GetDataset())
	out.GcsSource = GcsSource_FromProto(mapCtx, in.GetGcsSource())
	out.BigquerySource = BigQuerySource_FromProto(mapCtx, in.GetBigquerySource())
	out.DataFormat = direct.LazyPtr(in.GetDataFormat())
	out.TargetField = direct.LazyPtr(in.GetTargetField())
	out.LoggingSamplingStrategy = SamplingStrategy_FromProto(mapCtx, in.GetLoggingSamplingStrategy())
	return out
}
func ModelMonitoringObjectiveConfig_TrainingDataset_ToProto(mapCtx *direct.MapContext, in *krm.ModelMonitoringObjectiveConfig_TrainingDataset) *pb.ModelMonitoringObjectiveConfig_TrainingDataset {
	if in == nil {
		return nil
	}
	out := &pb.ModelMonitoringObjectiveConfig_TrainingDataset{}
	if oneof := ModelMonitoringObjectiveConfig_TrainingDataset_Dataset_ToProto(mapCtx, in.Dataset); oneof != nil {
		out.DataSource = oneof
	}
	if oneof := GcsSource_ToProto(mapCtx, in.GcsSource); oneof != nil {
		out.DataSource = &pb.ModelMonitoringObjectiveConfig_TrainingDataset_GcsSource{GcsSource: oneof}
	}
	if oneof := BigQuerySource_ToProto(mapCtx, in.BigquerySource); oneof != nil {
		out.DataSource = &pb.ModelMonitoringObjectiveConfig_TrainingDataset_BigquerySource{BigquerySource: oneof}
	}
	out.DataFormat = direct.ValueOf(in.DataFormat)
	out.TargetField = direct.ValueOf(in.TargetField)
	out.LoggingSamplingStrategy = SamplingStrategy_ToProto(mapCtx, in.LoggingSamplingStrategy)
	return out
}
func ModelMonitoringObjectiveConfig_TrainingPredictionSkewDetectionConfig_FromProto(mapCtx *direct.MapContext, in *pb.ModelMonitoringObjectiveConfig_TrainingPredictionSkewDetectionConfig) *krm.ModelMonitoringObjectiveConfig_TrainingPredictionSkewDetectionConfig {
	if in == nil {
		return nil
	}
	out := &krm.ModelMonitoringObjectiveConfig_TrainingPredictionSkewDetectionConfig{}
	// MISSING: SkewThresholds
	// MISSING: AttributionScoreSkewThresholds
	out.DefaultSkewThreshold = ThresholdConfig_FromProto(mapCtx, in.GetDefaultSkewThreshold())
	return out
}
func ModelMonitoringObjectiveConfig_TrainingPredictionSkewDetectionConfig_ToProto(mapCtx *direct.MapContext, in *krm.ModelMonitoringObjectiveConfig_TrainingPredictionSkewDetectionConfig) *pb.ModelMonitoringObjectiveConfig_TrainingPredictionSkewDetectionConfig {
	if in == nil {
		return nil
	}
	out := &pb.ModelMonitoringObjectiveConfig_TrainingPredictionSkewDetectionConfig{}
	// MISSING: SkewThresholds
	// MISSING: AttributionScoreSkewThresholds
	out.DefaultSkewThreshold = ThresholdConfig_ToProto(mapCtx, in.DefaultSkewThreshold)
	return out
}
func SamplingStrategy_FromProto(mapCtx *direct.MapContext, in *pb.SamplingStrategy) *krm.SamplingStrategy {
	if in == nil {
		return nil
	}
	out := &krm.SamplingStrategy{}
	out.RandomSampleConfig = SamplingStrategy_RandomSampleConfig_FromProto(mapCtx, in.GetRandomSampleConfig())
	return out
}
func SamplingStrategy_ToProto(mapCtx *direct.MapContext, in *krm.SamplingStrategy) *pb.SamplingStrategy {
	if in == nil {
		return nil
	}
	out := &pb.SamplingStrategy{}
	out.RandomSampleConfig = SamplingStrategy_RandomSampleConfig_ToProto(mapCtx, in.RandomSampleConfig)
	return out
}
func SamplingStrategy_RandomSampleConfig_FromProto(mapCtx *direct.MapContext, in *pb.SamplingStrategy_RandomSampleConfig) *krm.SamplingStrategy_RandomSampleConfig {
	if in == nil {
		return nil
	}
	out := &krm.SamplingStrategy_RandomSampleConfig{}
	out.SampleRate = direct.LazyPtr(in.GetSampleRate())
	return out
}
func SamplingStrategy_RandomSampleConfig_ToProto(mapCtx *direct.MapContext, in *krm.SamplingStrategy_RandomSampleConfig) *pb.SamplingStrategy_RandomSampleConfig {
	if in == nil {
		return nil
	}
	out := &pb.SamplingStrategy_RandomSampleConfig{}
	out.SampleRate = direct.ValueOf(in.SampleRate)
	return out
}
func ThresholdConfig_FromProto(mapCtx *direct.MapContext, in *pb.ThresholdConfig) *krm.ThresholdConfig {
	if in == nil {
		return nil
	}
	out := &krm.ThresholdConfig{}
	out.Value = direct.LazyPtr(in.GetValue())
	return out
}
func ThresholdConfig_ToProto(mapCtx *direct.MapContext, in *krm.ThresholdConfig) *pb.ThresholdConfig {
	if in == nil {
		return nil
	}
	out := &pb.ThresholdConfig{}
	if oneof := ThresholdConfig_Value_ToProto(mapCtx, in.Value); oneof != nil {
		out.Threshold = oneof
	}
	return out
}
