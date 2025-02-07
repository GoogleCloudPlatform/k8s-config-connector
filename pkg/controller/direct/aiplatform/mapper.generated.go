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
func AiplatformModelEvaluationObservedState_FromProto(mapCtx *direct.MapContext, in *pb.ModelEvaluation) *krm.AiplatformModelEvaluationObservedState {
	if in == nil {
		return nil
	}
	out := &krm.AiplatformModelEvaluationObservedState{}
	// MISSING: Name
	// MISSING: DisplayName
	// MISSING: MetricsSchemaURI
	// MISSING: Metrics
	// MISSING: CreateTime
	// MISSING: SliceDimensions
	// MISSING: DataItemSchemaURI
	// MISSING: AnnotationSchemaURI
	// MISSING: ModelExplanation
	// MISSING: ExplanationSpecs
	// MISSING: Metadata
	return out
}
func AiplatformModelEvaluationObservedState_ToProto(mapCtx *direct.MapContext, in *krm.AiplatformModelEvaluationObservedState) *pb.ModelEvaluation {
	if in == nil {
		return nil
	}
	out := &pb.ModelEvaluation{}
	// MISSING: Name
	// MISSING: DisplayName
	// MISSING: MetricsSchemaURI
	// MISSING: Metrics
	// MISSING: CreateTime
	// MISSING: SliceDimensions
	// MISSING: DataItemSchemaURI
	// MISSING: AnnotationSchemaURI
	// MISSING: ModelExplanation
	// MISSING: ExplanationSpecs
	// MISSING: Metadata
	return out
}
func AiplatformModelEvaluationSliceObservedState_FromProto(mapCtx *direct.MapContext, in *pb.ModelEvaluationSlice) *krm.AiplatformModelEvaluationSliceObservedState {
	if in == nil {
		return nil
	}
	out := &krm.AiplatformModelEvaluationSliceObservedState{}
	// MISSING: Name
	// MISSING: Slice
	// MISSING: MetricsSchemaURI
	// MISSING: Metrics
	// MISSING: CreateTime
	// MISSING: ModelExplanation
	return out
}
func AiplatformModelEvaluationSliceObservedState_ToProto(mapCtx *direct.MapContext, in *krm.AiplatformModelEvaluationSliceObservedState) *pb.ModelEvaluationSlice {
	if in == nil {
		return nil
	}
	out := &pb.ModelEvaluationSlice{}
	// MISSING: Name
	// MISSING: Slice
	// MISSING: MetricsSchemaURI
	// MISSING: Metrics
	// MISSING: CreateTime
	// MISSING: ModelExplanation
	return out
}
func AiplatformModelEvaluationSliceSpec_FromProto(mapCtx *direct.MapContext, in *pb.ModelEvaluationSlice) *krm.AiplatformModelEvaluationSliceSpec {
	if in == nil {
		return nil
	}
	out := &krm.AiplatformModelEvaluationSliceSpec{}
	// MISSING: Name
	// MISSING: Slice
	// MISSING: MetricsSchemaURI
	// MISSING: Metrics
	// MISSING: CreateTime
	// MISSING: ModelExplanation
	return out
}
func AiplatformModelEvaluationSliceSpec_ToProto(mapCtx *direct.MapContext, in *krm.AiplatformModelEvaluationSliceSpec) *pb.ModelEvaluationSlice {
	if in == nil {
		return nil
	}
	out := &pb.ModelEvaluationSlice{}
	// MISSING: Name
	// MISSING: Slice
	// MISSING: MetricsSchemaURI
	// MISSING: Metrics
	// MISSING: CreateTime
	// MISSING: ModelExplanation
	return out
}
func AiplatformModelEvaluationSpec_FromProto(mapCtx *direct.MapContext, in *pb.ModelEvaluation) *krm.AiplatformModelEvaluationSpec {
	if in == nil {
		return nil
	}
	out := &krm.AiplatformModelEvaluationSpec{}
	// MISSING: Name
	// MISSING: DisplayName
	// MISSING: MetricsSchemaURI
	// MISSING: Metrics
	// MISSING: CreateTime
	// MISSING: SliceDimensions
	// MISSING: DataItemSchemaURI
	// MISSING: AnnotationSchemaURI
	// MISSING: ModelExplanation
	// MISSING: ExplanationSpecs
	// MISSING: Metadata
	return out
}
func AiplatformModelEvaluationSpec_ToProto(mapCtx *direct.MapContext, in *krm.AiplatformModelEvaluationSpec) *pb.ModelEvaluation {
	if in == nil {
		return nil
	}
	out := &pb.ModelEvaluation{}
	// MISSING: Name
	// MISSING: DisplayName
	// MISSING: MetricsSchemaURI
	// MISSING: Metrics
	// MISSING: CreateTime
	// MISSING: SliceDimensions
	// MISSING: DataItemSchemaURI
	// MISSING: AnnotationSchemaURI
	// MISSING: ModelExplanation
	// MISSING: ExplanationSpecs
	// MISSING: Metadata
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
func AiplatformNasJobObservedState_FromProto(mapCtx *direct.MapContext, in *pb.NasJob) *krm.AiplatformNasJobObservedState {
	if in == nil {
		return nil
	}
	out := &krm.AiplatformNasJobObservedState{}
	// MISSING: Name
	// MISSING: DisplayName
	// MISSING: NasJobSpec
	// MISSING: NasJobOutput
	// MISSING: State
	// MISSING: CreateTime
	// MISSING: StartTime
	// MISSING: EndTime
	// MISSING: UpdateTime
	// MISSING: Error
	// MISSING: Labels
	// MISSING: EncryptionSpec
	// MISSING: EnableRestrictedImageTraining
	// MISSING: SatisfiesPzs
	// MISSING: SatisfiesPzi
	return out
}
func AiplatformNasJobObservedState_ToProto(mapCtx *direct.MapContext, in *krm.AiplatformNasJobObservedState) *pb.NasJob {
	if in == nil {
		return nil
	}
	out := &pb.NasJob{}
	// MISSING: Name
	// MISSING: DisplayName
	// MISSING: NasJobSpec
	// MISSING: NasJobOutput
	// MISSING: State
	// MISSING: CreateTime
	// MISSING: StartTime
	// MISSING: EndTime
	// MISSING: UpdateTime
	// MISSING: Error
	// MISSING: Labels
	// MISSING: EncryptionSpec
	// MISSING: EnableRestrictedImageTraining
	// MISSING: SatisfiesPzs
	// MISSING: SatisfiesPzi
	return out
}
func AiplatformNasJobSpec_FromProto(mapCtx *direct.MapContext, in *pb.NasJob) *krm.AiplatformNasJobSpec {
	if in == nil {
		return nil
	}
	out := &krm.AiplatformNasJobSpec{}
	// MISSING: Name
	// MISSING: DisplayName
	// MISSING: NasJobSpec
	// MISSING: NasJobOutput
	// MISSING: State
	// MISSING: CreateTime
	// MISSING: StartTime
	// MISSING: EndTime
	// MISSING: UpdateTime
	// MISSING: Error
	// MISSING: Labels
	// MISSING: EncryptionSpec
	// MISSING: EnableRestrictedImageTraining
	// MISSING: SatisfiesPzs
	// MISSING: SatisfiesPzi
	return out
}
func AiplatformNasJobSpec_ToProto(mapCtx *direct.MapContext, in *krm.AiplatformNasJobSpec) *pb.NasJob {
	if in == nil {
		return nil
	}
	out := &pb.NasJob{}
	// MISSING: Name
	// MISSING: DisplayName
	// MISSING: NasJobSpec
	// MISSING: NasJobOutput
	// MISSING: State
	// MISSING: CreateTime
	// MISSING: StartTime
	// MISSING: EndTime
	// MISSING: UpdateTime
	// MISSING: Error
	// MISSING: Labels
	// MISSING: EncryptionSpec
	// MISSING: EnableRestrictedImageTraining
	// MISSING: SatisfiesPzs
	// MISSING: SatisfiesPzi
	return out
}
func AiplatformNasTrialDetailObservedState_FromProto(mapCtx *direct.MapContext, in *pb.NasTrialDetail) *krm.AiplatformNasTrialDetailObservedState {
	if in == nil {
		return nil
	}
	out := &krm.AiplatformNasTrialDetailObservedState{}
	// MISSING: Name
	// MISSING: Parameters
	// MISSING: SearchTrial
	// MISSING: TrainTrial
	return out
}
func AiplatformNasTrialDetailObservedState_ToProto(mapCtx *direct.MapContext, in *krm.AiplatformNasTrialDetailObservedState) *pb.NasTrialDetail {
	if in == nil {
		return nil
	}
	out := &pb.NasTrialDetail{}
	// MISSING: Name
	// MISSING: Parameters
	// MISSING: SearchTrial
	// MISSING: TrainTrial
	return out
}
func AiplatformNasTrialDetailSpec_FromProto(mapCtx *direct.MapContext, in *pb.NasTrialDetail) *krm.AiplatformNasTrialDetailSpec {
	if in == nil {
		return nil
	}
	out := &krm.AiplatformNasTrialDetailSpec{}
	// MISSING: Name
	// MISSING: Parameters
	// MISSING: SearchTrial
	// MISSING: TrainTrial
	return out
}
func AiplatformNasTrialDetailSpec_ToProto(mapCtx *direct.MapContext, in *krm.AiplatformNasTrialDetailSpec) *pb.NasTrialDetail {
	if in == nil {
		return nil
	}
	out := &pb.NasTrialDetail{}
	// MISSING: Name
	// MISSING: Parameters
	// MISSING: SearchTrial
	// MISSING: TrainTrial
	return out
}
func AiplatformNotebookExecutionJobObservedState_FromProto(mapCtx *direct.MapContext, in *pb.NotebookExecutionJob) *krm.AiplatformNotebookExecutionJobObservedState {
	if in == nil {
		return nil
	}
	out := &krm.AiplatformNotebookExecutionJobObservedState{}
	// MISSING: DataformRepositorySource
	// MISSING: GcsNotebookSource
	// MISSING: DirectNotebookSource
	// MISSING: NotebookRuntimeTemplateResourceName
	// MISSING: CustomEnvironmentSpec
	// MISSING: GcsOutputURI
	// MISSING: ExecutionUser
	// MISSING: ServiceAccount
	// MISSING: WorkbenchRuntime
	// MISSING: Name
	// MISSING: DisplayName
	// MISSING: ExecutionTimeout
	// MISSING: ScheduleResourceName
	// MISSING: JobState
	// MISSING: Status
	// MISSING: CreateTime
	// MISSING: UpdateTime
	// MISSING: Labels
	// MISSING: KernelName
	// MISSING: EncryptionSpec
	return out
}
func AiplatformNotebookExecutionJobObservedState_ToProto(mapCtx *direct.MapContext, in *krm.AiplatformNotebookExecutionJobObservedState) *pb.NotebookExecutionJob {
	if in == nil {
		return nil
	}
	out := &pb.NotebookExecutionJob{}
	// MISSING: DataformRepositorySource
	// MISSING: GcsNotebookSource
	// MISSING: DirectNotebookSource
	// MISSING: NotebookRuntimeTemplateResourceName
	// MISSING: CustomEnvironmentSpec
	// MISSING: GcsOutputURI
	// MISSING: ExecutionUser
	// MISSING: ServiceAccount
	// MISSING: WorkbenchRuntime
	// MISSING: Name
	// MISSING: DisplayName
	// MISSING: ExecutionTimeout
	// MISSING: ScheduleResourceName
	// MISSING: JobState
	// MISSING: Status
	// MISSING: CreateTime
	// MISSING: UpdateTime
	// MISSING: Labels
	// MISSING: KernelName
	// MISSING: EncryptionSpec
	return out
}
func AiplatformNotebookExecutionJobSpec_FromProto(mapCtx *direct.MapContext, in *pb.NotebookExecutionJob) *krm.AiplatformNotebookExecutionJobSpec {
	if in == nil {
		return nil
	}
	out := &krm.AiplatformNotebookExecutionJobSpec{}
	// MISSING: DataformRepositorySource
	// MISSING: GcsNotebookSource
	// MISSING: DirectNotebookSource
	// MISSING: NotebookRuntimeTemplateResourceName
	// MISSING: CustomEnvironmentSpec
	// MISSING: GcsOutputURI
	// MISSING: ExecutionUser
	// MISSING: ServiceAccount
	// MISSING: WorkbenchRuntime
	// MISSING: Name
	// MISSING: DisplayName
	// MISSING: ExecutionTimeout
	// MISSING: ScheduleResourceName
	// MISSING: JobState
	// MISSING: Status
	// MISSING: CreateTime
	// MISSING: UpdateTime
	// MISSING: Labels
	// MISSING: KernelName
	// MISSING: EncryptionSpec
	return out
}
func AiplatformNotebookExecutionJobSpec_ToProto(mapCtx *direct.MapContext, in *krm.AiplatformNotebookExecutionJobSpec) *pb.NotebookExecutionJob {
	if in == nil {
		return nil
	}
	out := &pb.NotebookExecutionJob{}
	// MISSING: DataformRepositorySource
	// MISSING: GcsNotebookSource
	// MISSING: DirectNotebookSource
	// MISSING: NotebookRuntimeTemplateResourceName
	// MISSING: CustomEnvironmentSpec
	// MISSING: GcsOutputURI
	// MISSING: ExecutionUser
	// MISSING: ServiceAccount
	// MISSING: WorkbenchRuntime
	// MISSING: Name
	// MISSING: DisplayName
	// MISSING: ExecutionTimeout
	// MISSING: ScheduleResourceName
	// MISSING: JobState
	// MISSING: Status
	// MISSING: CreateTime
	// MISSING: UpdateTime
	// MISSING: Labels
	// MISSING: KernelName
	// MISSING: EncryptionSpec
	return out
}
func AiplatformNotebookRuntimeObservedState_FromProto(mapCtx *direct.MapContext, in *pb.NotebookRuntime) *krm.AiplatformNotebookRuntimeObservedState {
	if in == nil {
		return nil
	}
	out := &krm.AiplatformNotebookRuntimeObservedState{}
	// MISSING: Name
	// MISSING: RuntimeUser
	// MISSING: NotebookRuntimeTemplateRef
	// MISSING: ProxyURI
	// MISSING: CreateTime
	// MISSING: UpdateTime
	// MISSING: HealthState
	// MISSING: DisplayName
	// MISSING: Description
	// MISSING: ServiceAccount
	// MISSING: RuntimeState
	// MISSING: IsUpgradable
	// MISSING: Labels
	// MISSING: ExpirationTime
	// MISSING: Version
	// MISSING: NotebookRuntimeType
	// MISSING: IdleShutdownConfig
	// MISSING: NetworkTags
	// MISSING: EncryptionSpec
	// MISSING: SatisfiesPzs
	// MISSING: SatisfiesPzi
	return out
}
func AiplatformNotebookRuntimeObservedState_ToProto(mapCtx *direct.MapContext, in *krm.AiplatformNotebookRuntimeObservedState) *pb.NotebookRuntime {
	if in == nil {
		return nil
	}
	out := &pb.NotebookRuntime{}
	// MISSING: Name
	// MISSING: RuntimeUser
	// MISSING: NotebookRuntimeTemplateRef
	// MISSING: ProxyURI
	// MISSING: CreateTime
	// MISSING: UpdateTime
	// MISSING: HealthState
	// MISSING: DisplayName
	// MISSING: Description
	// MISSING: ServiceAccount
	// MISSING: RuntimeState
	// MISSING: IsUpgradable
	// MISSING: Labels
	// MISSING: ExpirationTime
	// MISSING: Version
	// MISSING: NotebookRuntimeType
	// MISSING: IdleShutdownConfig
	// MISSING: NetworkTags
	// MISSING: EncryptionSpec
	// MISSING: SatisfiesPzs
	// MISSING: SatisfiesPzi
	return out
}
func AiplatformNotebookRuntimeSpec_FromProto(mapCtx *direct.MapContext, in *pb.NotebookRuntime) *krm.AiplatformNotebookRuntimeSpec {
	if in == nil {
		return nil
	}
	out := &krm.AiplatformNotebookRuntimeSpec{}
	// MISSING: Name
	// MISSING: RuntimeUser
	// MISSING: NotebookRuntimeTemplateRef
	// MISSING: ProxyURI
	// MISSING: CreateTime
	// MISSING: UpdateTime
	// MISSING: HealthState
	// MISSING: DisplayName
	// MISSING: Description
	// MISSING: ServiceAccount
	// MISSING: RuntimeState
	// MISSING: IsUpgradable
	// MISSING: Labels
	// MISSING: ExpirationTime
	// MISSING: Version
	// MISSING: NotebookRuntimeType
	// MISSING: IdleShutdownConfig
	// MISSING: NetworkTags
	// MISSING: EncryptionSpec
	// MISSING: SatisfiesPzs
	// MISSING: SatisfiesPzi
	return out
}
func AiplatformNotebookRuntimeSpec_ToProto(mapCtx *direct.MapContext, in *krm.AiplatformNotebookRuntimeSpec) *pb.NotebookRuntime {
	if in == nil {
		return nil
	}
	out := &pb.NotebookRuntime{}
	// MISSING: Name
	// MISSING: RuntimeUser
	// MISSING: NotebookRuntimeTemplateRef
	// MISSING: ProxyURI
	// MISSING: CreateTime
	// MISSING: UpdateTime
	// MISSING: HealthState
	// MISSING: DisplayName
	// MISSING: Description
	// MISSING: ServiceAccount
	// MISSING: RuntimeState
	// MISSING: IsUpgradable
	// MISSING: Labels
	// MISSING: ExpirationTime
	// MISSING: Version
	// MISSING: NotebookRuntimeType
	// MISSING: IdleShutdownConfig
	// MISSING: NetworkTags
	// MISSING: EncryptionSpec
	// MISSING: SatisfiesPzs
	// MISSING: SatisfiesPzi
	return out
}
func AiplatformNotebookRuntimeTemplateObservedState_FromProto(mapCtx *direct.MapContext, in *pb.NotebookRuntimeTemplate) *krm.AiplatformNotebookRuntimeTemplateObservedState {
	if in == nil {
		return nil
	}
	out := &krm.AiplatformNotebookRuntimeTemplateObservedState{}
	// MISSING: Name
	// MISSING: DisplayName
	// MISSING: Description
	// MISSING: IsDefault
	// MISSING: MachineSpec
	// MISSING: DataPersistentDiskSpec
	// MISSING: NetworkSpec
	// MISSING: ServiceAccount
	// MISSING: Etag
	// MISSING: Labels
	// MISSING: IdleShutdownConfig
	// MISSING: EucConfig
	// MISSING: CreateTime
	// MISSING: UpdateTime
	// MISSING: NotebookRuntimeType
	// MISSING: ShieldedVmConfig
	// MISSING: NetworkTags
	// MISSING: EncryptionSpec
	return out
}
func AiplatformNotebookRuntimeTemplateObservedState_ToProto(mapCtx *direct.MapContext, in *krm.AiplatformNotebookRuntimeTemplateObservedState) *pb.NotebookRuntimeTemplate {
	if in == nil {
		return nil
	}
	out := &pb.NotebookRuntimeTemplate{}
	// MISSING: Name
	// MISSING: DisplayName
	// MISSING: Description
	// MISSING: IsDefault
	// MISSING: MachineSpec
	// MISSING: DataPersistentDiskSpec
	// MISSING: NetworkSpec
	// MISSING: ServiceAccount
	// MISSING: Etag
	// MISSING: Labels
	// MISSING: IdleShutdownConfig
	// MISSING: EucConfig
	// MISSING: CreateTime
	// MISSING: UpdateTime
	// MISSING: NotebookRuntimeType
	// MISSING: ShieldedVmConfig
	// MISSING: NetworkTags
	// MISSING: EncryptionSpec
	return out
}
func AiplatformNotebookRuntimeTemplateSpec_FromProto(mapCtx *direct.MapContext, in *pb.NotebookRuntimeTemplate) *krm.AiplatformNotebookRuntimeTemplateSpec {
	if in == nil {
		return nil
	}
	out := &krm.AiplatformNotebookRuntimeTemplateSpec{}
	// MISSING: Name
	// MISSING: DisplayName
	// MISSING: Description
	// MISSING: IsDefault
	// MISSING: MachineSpec
	// MISSING: DataPersistentDiskSpec
	// MISSING: NetworkSpec
	// MISSING: ServiceAccount
	// MISSING: Etag
	// MISSING: Labels
	// MISSING: IdleShutdownConfig
	// MISSING: EucConfig
	// MISSING: CreateTime
	// MISSING: UpdateTime
	// MISSING: NotebookRuntimeType
	// MISSING: ShieldedVmConfig
	// MISSING: NetworkTags
	// MISSING: EncryptionSpec
	return out
}
func AiplatformNotebookRuntimeTemplateSpec_ToProto(mapCtx *direct.MapContext, in *krm.AiplatformNotebookRuntimeTemplateSpec) *pb.NotebookRuntimeTemplate {
	if in == nil {
		return nil
	}
	out := &pb.NotebookRuntimeTemplate{}
	// MISSING: Name
	// MISSING: DisplayName
	// MISSING: Description
	// MISSING: IsDefault
	// MISSING: MachineSpec
	// MISSING: DataPersistentDiskSpec
	// MISSING: NetworkSpec
	// MISSING: ServiceAccount
	// MISSING: Etag
	// MISSING: Labels
	// MISSING: IdleShutdownConfig
	// MISSING: EucConfig
	// MISSING: CreateTime
	// MISSING: UpdateTime
	// MISSING: NotebookRuntimeType
	// MISSING: ShieldedVmConfig
	// MISSING: NetworkTags
	// MISSING: EncryptionSpec
	return out
}
func AiplatformPersistentResourceObservedState_FromProto(mapCtx *direct.MapContext, in *pb.PersistentResource) *krm.AiplatformPersistentResourceObservedState {
	if in == nil {
		return nil
	}
	out := &krm.AiplatformPersistentResourceObservedState{}
	// MISSING: Name
	// MISSING: DisplayName
	// MISSING: ResourcePools
	// MISSING: State
	// MISSING: Error
	// MISSING: CreateTime
	// MISSING: StartTime
	// MISSING: UpdateTime
	// MISSING: Labels
	// MISSING: Network
	// MISSING: EncryptionSpec
	// MISSING: ResourceRuntimeSpec
	// MISSING: ResourceRuntime
	// MISSING: ReservedIPRanges
	return out
}
func AiplatformPersistentResourceObservedState_ToProto(mapCtx *direct.MapContext, in *krm.AiplatformPersistentResourceObservedState) *pb.PersistentResource {
	if in == nil {
		return nil
	}
	out := &pb.PersistentResource{}
	// MISSING: Name
	// MISSING: DisplayName
	// MISSING: ResourcePools
	// MISSING: State
	// MISSING: Error
	// MISSING: CreateTime
	// MISSING: StartTime
	// MISSING: UpdateTime
	// MISSING: Labels
	// MISSING: Network
	// MISSING: EncryptionSpec
	// MISSING: ResourceRuntimeSpec
	// MISSING: ResourceRuntime
	// MISSING: ReservedIPRanges
	return out
}
func AiplatformPersistentResourceSpec_FromProto(mapCtx *direct.MapContext, in *pb.PersistentResource) *krm.AiplatformPersistentResourceSpec {
	if in == nil {
		return nil
	}
	out := &krm.AiplatformPersistentResourceSpec{}
	// MISSING: Name
	// MISSING: DisplayName
	// MISSING: ResourcePools
	// MISSING: State
	// MISSING: Error
	// MISSING: CreateTime
	// MISSING: StartTime
	// MISSING: UpdateTime
	// MISSING: Labels
	// MISSING: Network
	// MISSING: EncryptionSpec
	// MISSING: ResourceRuntimeSpec
	// MISSING: ResourceRuntime
	// MISSING: ReservedIPRanges
	return out
}
func AiplatformPersistentResourceSpec_ToProto(mapCtx *direct.MapContext, in *krm.AiplatformPersistentResourceSpec) *pb.PersistentResource {
	if in == nil {
		return nil
	}
	out := &pb.PersistentResource{}
	// MISSING: Name
	// MISSING: DisplayName
	// MISSING: ResourcePools
	// MISSING: State
	// MISSING: Error
	// MISSING: CreateTime
	// MISSING: StartTime
	// MISSING: UpdateTime
	// MISSING: Labels
	// MISSING: Network
	// MISSING: EncryptionSpec
	// MISSING: ResourceRuntimeSpec
	// MISSING: ResourceRuntime
	// MISSING: ReservedIPRanges
	return out
}
func AiplatformPipelineJobObservedState_FromProto(mapCtx *direct.MapContext, in *pb.PipelineJob) *krm.AiplatformPipelineJobObservedState {
	if in == nil {
		return nil
	}
	out := &krm.AiplatformPipelineJobObservedState{}
	// MISSING: Name
	// MISSING: DisplayName
	// MISSING: CreateTime
	// MISSING: StartTime
	// MISSING: EndTime
	// MISSING: UpdateTime
	// MISSING: PipelineSpec
	// MISSING: State
	// MISSING: JobDetail
	// MISSING: Error
	// MISSING: Labels
	// MISSING: RuntimeConfig
	// MISSING: EncryptionSpec
	// MISSING: ServiceAccount
	// MISSING: Network
	// MISSING: ReservedIPRanges
	// MISSING: TemplateURI
	// MISSING: TemplateMetadata
	// MISSING: ScheduleName
	// MISSING: PreflightValidations
	return out
}
func AiplatformPipelineJobObservedState_ToProto(mapCtx *direct.MapContext, in *krm.AiplatformPipelineJobObservedState) *pb.PipelineJob {
	if in == nil {
		return nil
	}
	out := &pb.PipelineJob{}
	// MISSING: Name
	// MISSING: DisplayName
	// MISSING: CreateTime
	// MISSING: StartTime
	// MISSING: EndTime
	// MISSING: UpdateTime
	// MISSING: PipelineSpec
	// MISSING: State
	// MISSING: JobDetail
	// MISSING: Error
	// MISSING: Labels
	// MISSING: RuntimeConfig
	// MISSING: EncryptionSpec
	// MISSING: ServiceAccount
	// MISSING: Network
	// MISSING: ReservedIPRanges
	// MISSING: TemplateURI
	// MISSING: TemplateMetadata
	// MISSING: ScheduleName
	// MISSING: PreflightValidations
	return out
}
func AiplatformPipelineJobSpec_FromProto(mapCtx *direct.MapContext, in *pb.PipelineJob) *krm.AiplatformPipelineJobSpec {
	if in == nil {
		return nil
	}
	out := &krm.AiplatformPipelineJobSpec{}
	// MISSING: Name
	// MISSING: DisplayName
	// MISSING: CreateTime
	// MISSING: StartTime
	// MISSING: EndTime
	// MISSING: UpdateTime
	// MISSING: PipelineSpec
	// MISSING: State
	// MISSING: JobDetail
	// MISSING: Error
	// MISSING: Labels
	// MISSING: RuntimeConfig
	// MISSING: EncryptionSpec
	// MISSING: ServiceAccount
	// MISSING: Network
	// MISSING: ReservedIPRanges
	// MISSING: TemplateURI
	// MISSING: TemplateMetadata
	// MISSING: ScheduleName
	// MISSING: PreflightValidations
	return out
}
func AiplatformPipelineJobSpec_ToProto(mapCtx *direct.MapContext, in *krm.AiplatformPipelineJobSpec) *pb.PipelineJob {
	if in == nil {
		return nil
	}
	out := &pb.PipelineJob{}
	// MISSING: Name
	// MISSING: DisplayName
	// MISSING: CreateTime
	// MISSING: StartTime
	// MISSING: EndTime
	// MISSING: UpdateTime
	// MISSING: PipelineSpec
	// MISSING: State
	// MISSING: JobDetail
	// MISSING: Error
	// MISSING: Labels
	// MISSING: RuntimeConfig
	// MISSING: EncryptionSpec
	// MISSING: ServiceAccount
	// MISSING: Network
	// MISSING: ReservedIPRanges
	// MISSING: TemplateURI
	// MISSING: TemplateMetadata
	// MISSING: ScheduleName
	// MISSING: PreflightValidations
	return out
}
func AiplatformPublisherModelObservedState_FromProto(mapCtx *direct.MapContext, in *pb.PublisherModel) *krm.AiplatformPublisherModelObservedState {
	if in == nil {
		return nil
	}
	out := &krm.AiplatformPublisherModelObservedState{}
	// MISSING: Name
	// MISSING: VersionID
	// MISSING: OpenSourceCategory
	// MISSING: SupportedActions
	// MISSING: Frameworks
	// MISSING: LaunchStage
	// MISSING: VersionState
	// MISSING: PublisherModelTemplate
	// MISSING: PredictSchemata
	return out
}
func AiplatformPublisherModelObservedState_ToProto(mapCtx *direct.MapContext, in *krm.AiplatformPublisherModelObservedState) *pb.PublisherModel {
	if in == nil {
		return nil
	}
	out := &pb.PublisherModel{}
	// MISSING: Name
	// MISSING: VersionID
	// MISSING: OpenSourceCategory
	// MISSING: SupportedActions
	// MISSING: Frameworks
	// MISSING: LaunchStage
	// MISSING: VersionState
	// MISSING: PublisherModelTemplate
	// MISSING: PredictSchemata
	return out
}
func AiplatformPublisherModelSpec_FromProto(mapCtx *direct.MapContext, in *pb.PublisherModel) *krm.AiplatformPublisherModelSpec {
	if in == nil {
		return nil
	}
	out := &krm.AiplatformPublisherModelSpec{}
	// MISSING: Name
	// MISSING: VersionID
	// MISSING: OpenSourceCategory
	// MISSING: SupportedActions
	// MISSING: Frameworks
	// MISSING: LaunchStage
	// MISSING: VersionState
	// MISSING: PublisherModelTemplate
	// MISSING: PredictSchemata
	return out
}
func AiplatformPublisherModelSpec_ToProto(mapCtx *direct.MapContext, in *krm.AiplatformPublisherModelSpec) *pb.PublisherModel {
	if in == nil {
		return nil
	}
	out := &pb.PublisherModel{}
	// MISSING: Name
	// MISSING: VersionID
	// MISSING: OpenSourceCategory
	// MISSING: SupportedActions
	// MISSING: Frameworks
	// MISSING: LaunchStage
	// MISSING: VersionState
	// MISSING: PublisherModelTemplate
	// MISSING: PredictSchemata
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
func EnvVar_FromProto(mapCtx *direct.MapContext, in *pb.EnvVar) *krm.EnvVar {
	if in == nil {
		return nil
	}
	out := &krm.EnvVar{}
	out.Name = direct.LazyPtr(in.GetName())
	out.Value = direct.LazyPtr(in.GetValue())
	return out
}
func EnvVar_ToProto(mapCtx *direct.MapContext, in *krm.EnvVar) *pb.EnvVar {
	if in == nil {
		return nil
	}
	out := &pb.EnvVar{}
	out.Name = direct.ValueOf(in.Name)
	out.Value = direct.ValueOf(in.Value)
	return out
}
func LargeModelReference_FromProto(mapCtx *direct.MapContext, in *pb.LargeModelReference) *krm.LargeModelReference {
	if in == nil {
		return nil
	}
	out := &krm.LargeModelReference{}
	out.Name = direct.LazyPtr(in.GetName())
	return out
}
func LargeModelReference_ToProto(mapCtx *direct.MapContext, in *krm.LargeModelReference) *pb.LargeModelReference {
	if in == nil {
		return nil
	}
	out := &pb.LargeModelReference{}
	out.Name = direct.ValueOf(in.Name)
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
func ModelContainerSpec_FromProto(mapCtx *direct.MapContext, in *pb.ModelContainerSpec) *krm.ModelContainerSpec {
	if in == nil {
		return nil
	}
	out := &krm.ModelContainerSpec{}
	out.ImageURI = direct.LazyPtr(in.GetImageUri())
	out.Command = in.Command
	out.Args = in.Args
	out.Env = direct.Slice_FromProto(mapCtx, in.Env, EnvVar_FromProto)
	out.Ports = direct.Slice_FromProto(mapCtx, in.Ports, Port_FromProto)
	out.PredictRoute = direct.LazyPtr(in.GetPredictRoute())
	out.HealthRoute = direct.LazyPtr(in.GetHealthRoute())
	out.GrpcPorts = direct.Slice_FromProto(mapCtx, in.GrpcPorts, Port_FromProto)
	out.DeploymentTimeout = direct.StringDuration_FromProto(mapCtx, in.GetDeploymentTimeout())
	out.SharedMemorySizeMb = direct.LazyPtr(in.GetSharedMemorySizeMb())
	out.StartupProbe = Probe_FromProto(mapCtx, in.GetStartupProbe())
	out.HealthProbe = Probe_FromProto(mapCtx, in.GetHealthProbe())
	return out
}
func ModelContainerSpec_ToProto(mapCtx *direct.MapContext, in *krm.ModelContainerSpec) *pb.ModelContainerSpec {
	if in == nil {
		return nil
	}
	out := &pb.ModelContainerSpec{}
	out.ImageUri = direct.ValueOf(in.ImageURI)
	out.Command = in.Command
	out.Args = in.Args
	out.Env = direct.Slice_ToProto(mapCtx, in.Env, EnvVar_ToProto)
	out.Ports = direct.Slice_ToProto(mapCtx, in.Ports, Port_ToProto)
	out.PredictRoute = direct.ValueOf(in.PredictRoute)
	out.HealthRoute = direct.ValueOf(in.HealthRoute)
	out.GrpcPorts = direct.Slice_ToProto(mapCtx, in.GrpcPorts, Port_ToProto)
	out.DeploymentTimeout = direct.StringDuration_ToProto(mapCtx, in.DeploymentTimeout)
	out.SharedMemorySizeMb = direct.ValueOf(in.SharedMemorySizeMb)
	out.StartupProbe = Probe_ToProto(mapCtx, in.StartupProbe)
	out.HealthProbe = Probe_ToProto(mapCtx, in.HealthProbe)
	return out
}
func Port_FromProto(mapCtx *direct.MapContext, in *pb.Port) *krm.Port {
	if in == nil {
		return nil
	}
	out := &krm.Port{}
	out.ContainerPort = direct.LazyPtr(in.GetContainerPort())
	return out
}
func Port_ToProto(mapCtx *direct.MapContext, in *krm.Port) *pb.Port {
	if in == nil {
		return nil
	}
	out := &pb.Port{}
	out.ContainerPort = direct.ValueOf(in.ContainerPort)
	return out
}
func PredictSchemata_FromProto(mapCtx *direct.MapContext, in *pb.PredictSchemata) *krm.PredictSchemata {
	if in == nil {
		return nil
	}
	out := &krm.PredictSchemata{}
	out.InstanceSchemaURI = direct.LazyPtr(in.GetInstanceSchemaUri())
	out.ParametersSchemaURI = direct.LazyPtr(in.GetParametersSchemaUri())
	out.PredictionSchemaURI = direct.LazyPtr(in.GetPredictionSchemaUri())
	return out
}
func PredictSchemata_ToProto(mapCtx *direct.MapContext, in *krm.PredictSchemata) *pb.PredictSchemata {
	if in == nil {
		return nil
	}
	out := &pb.PredictSchemata{}
	out.InstanceSchemaUri = direct.ValueOf(in.InstanceSchemaURI)
	out.ParametersSchemaUri = direct.ValueOf(in.ParametersSchemaURI)
	out.PredictionSchemaUri = direct.ValueOf(in.PredictionSchemaURI)
	return out
}
func Probe_FromProto(mapCtx *direct.MapContext, in *pb.Probe) *krm.Probe {
	if in == nil {
		return nil
	}
	out := &krm.Probe{}
	out.Exec = Probe_ExecAction_FromProto(mapCtx, in.GetExec())
	out.PeriodSeconds = direct.LazyPtr(in.GetPeriodSeconds())
	out.TimeoutSeconds = direct.LazyPtr(in.GetTimeoutSeconds())
	return out
}
func Probe_ToProto(mapCtx *direct.MapContext, in *krm.Probe) *pb.Probe {
	if in == nil {
		return nil
	}
	out := &pb.Probe{}
	if oneof := Probe_ExecAction_ToProto(mapCtx, in.Exec); oneof != nil {
		out.ProbeType = &pb.Probe_Exec{Exec: oneof}
	}
	out.PeriodSeconds = direct.ValueOf(in.PeriodSeconds)
	out.TimeoutSeconds = direct.ValueOf(in.TimeoutSeconds)
	return out
}
func Probe_ExecAction_FromProto(mapCtx *direct.MapContext, in *pb.Probe_ExecAction) *krm.Probe_ExecAction {
	if in == nil {
		return nil
	}
	out := &krm.Probe_ExecAction{}
	out.Command = in.Command
	return out
}
func Probe_ExecAction_ToProto(mapCtx *direct.MapContext, in *krm.Probe_ExecAction) *pb.Probe_ExecAction {
	if in == nil {
		return nil
	}
	out := &pb.Probe_ExecAction{}
	out.Command = in.Command
	return out
}
func PublisherModel_FromProto(mapCtx *direct.MapContext, in *pb.PublisherModel) *krm.PublisherModel {
	if in == nil {
		return nil
	}
	out := &krm.PublisherModel{}
	// MISSING: Name
	// MISSING: VersionID
	out.OpenSourceCategory = direct.Enum_FromProto(mapCtx, in.GetOpenSourceCategory())
	out.SupportedActions = PublisherModel_CallToAction_FromProto(mapCtx, in.GetSupportedActions())
	out.Frameworks = in.Frameworks
	out.LaunchStage = direct.Enum_FromProto(mapCtx, in.GetLaunchStage())
	out.VersionState = direct.Enum_FromProto(mapCtx, in.GetVersionState())
	// MISSING: PublisherModelTemplate
	out.PredictSchemata = PredictSchemata_FromProto(mapCtx, in.GetPredictSchemata())
	return out
}
func PublisherModel_ToProto(mapCtx *direct.MapContext, in *krm.PublisherModel) *pb.PublisherModel {
	if in == nil {
		return nil
	}
	out := &pb.PublisherModel{}
	// MISSING: Name
	// MISSING: VersionID
	out.OpenSourceCategory = direct.Enum_ToProto[pb.PublisherModel_OpenSourceCategory](mapCtx, in.OpenSourceCategory)
	out.SupportedActions = PublisherModel_CallToAction_ToProto(mapCtx, in.SupportedActions)
	out.Frameworks = in.Frameworks
	out.LaunchStage = direct.Enum_ToProto[pb.PublisherModel_LaunchStage](mapCtx, in.LaunchStage)
	out.VersionState = direct.Enum_ToProto[pb.PublisherModel_VersionState](mapCtx, in.VersionState)
	// MISSING: PublisherModelTemplate
	out.PredictSchemata = PredictSchemata_ToProto(mapCtx, in.PredictSchemata)
	return out
}
func PublisherModelObservedState_FromProto(mapCtx *direct.MapContext, in *pb.PublisherModel) *krm.PublisherModelObservedState {
	if in == nil {
		return nil
	}
	out := &krm.PublisherModelObservedState{}
	out.Name = direct.LazyPtr(in.GetName())
	out.VersionID = direct.LazyPtr(in.GetVersionId())
	// MISSING: OpenSourceCategory
	// MISSING: SupportedActions
	// MISSING: Frameworks
	// MISSING: LaunchStage
	// MISSING: VersionState
	out.PublisherModelTemplate = direct.LazyPtr(in.GetPublisherModelTemplate())
	// MISSING: PredictSchemata
	return out
}
func PublisherModelObservedState_ToProto(mapCtx *direct.MapContext, in *krm.PublisherModelObservedState) *pb.PublisherModel {
	if in == nil {
		return nil
	}
	out := &pb.PublisherModel{}
	out.Name = direct.ValueOf(in.Name)
	out.VersionId = direct.ValueOf(in.VersionID)
	// MISSING: OpenSourceCategory
	// MISSING: SupportedActions
	// MISSING: Frameworks
	// MISSING: LaunchStage
	// MISSING: VersionState
	out.PublisherModelTemplate = direct.ValueOf(in.PublisherModelTemplate)
	// MISSING: PredictSchemata
	return out
}
func PublisherModel_CallToAction_FromProto(mapCtx *direct.MapContext, in *pb.PublisherModel_CallToAction) *krm.PublisherModel_CallToAction {
	if in == nil {
		return nil
	}
	out := &krm.PublisherModel_CallToAction{}
	out.ViewRestApi = PublisherModel_CallToAction_ViewRestApi_FromProto(mapCtx, in.GetViewRestApi())
	out.OpenNotebook = PublisherModel_CallToAction_RegionalResourceReferences_FromProto(mapCtx, in.GetOpenNotebook())
	out.OpenNotebooks = PublisherModel_CallToAction_OpenNotebooks_FromProto(mapCtx, in.GetOpenNotebooks())
	out.CreateApplication = PublisherModel_CallToAction_RegionalResourceReferences_FromProto(mapCtx, in.GetCreateApplication())
	out.OpenFineTuningPipeline = PublisherModel_CallToAction_RegionalResourceReferences_FromProto(mapCtx, in.GetOpenFineTuningPipeline())
	out.OpenFineTuningPipelines = PublisherModel_CallToAction_OpenFineTuningPipelines_FromProto(mapCtx, in.GetOpenFineTuningPipelines())
	out.OpenPromptTuningPipeline = PublisherModel_CallToAction_RegionalResourceReferences_FromProto(mapCtx, in.GetOpenPromptTuningPipeline())
	out.OpenGenie = PublisherModel_CallToAction_RegionalResourceReferences_FromProto(mapCtx, in.GetOpenGenie())
	out.Deploy = PublisherModel_CallToAction_Deploy_FromProto(mapCtx, in.GetDeploy())
	out.DeployGke = PublisherModel_CallToAction_DeployGke_FromProto(mapCtx, in.GetDeployGke())
	out.OpenGenerationAiStudio = PublisherModel_CallToAction_RegionalResourceReferences_FromProto(mapCtx, in.GetOpenGenerationAiStudio())
	out.RequestAccess = PublisherModel_CallToAction_RegionalResourceReferences_FromProto(mapCtx, in.GetRequestAccess())
	out.OpenEvaluationPipeline = PublisherModel_CallToAction_RegionalResourceReferences_FromProto(mapCtx, in.GetOpenEvaluationPipeline())
	return out
}
func PublisherModel_CallToAction_ToProto(mapCtx *direct.MapContext, in *krm.PublisherModel_CallToAction) *pb.PublisherModel_CallToAction {
	if in == nil {
		return nil
	}
	out := &pb.PublisherModel_CallToAction{}
	out.ViewRestApi = PublisherModel_CallToAction_ViewRestApi_ToProto(mapCtx, in.ViewRestApi)
	out.OpenNotebook = PublisherModel_CallToAction_RegionalResourceReferences_ToProto(mapCtx, in.OpenNotebook)
	if oneof := PublisherModel_CallToAction_OpenNotebooks_ToProto(mapCtx, in.OpenNotebooks); oneof != nil {
		out.OpenNotebooks = &pb.PublisherModel_CallToAction_OpenNotebooks_{OpenNotebooks: oneof}
	}
	out.CreateApplication = PublisherModel_CallToAction_RegionalResourceReferences_ToProto(mapCtx, in.CreateApplication)
	out.OpenFineTuningPipeline = PublisherModel_CallToAction_RegionalResourceReferences_ToProto(mapCtx, in.OpenFineTuningPipeline)
	if oneof := PublisherModel_CallToAction_OpenFineTuningPipelines_ToProto(mapCtx, in.OpenFineTuningPipelines); oneof != nil {
		out.OpenFineTuningPipelines = &pb.PublisherModel_CallToAction_OpenFineTuningPipelines_{OpenFineTuningPipelines: oneof}
	}
	out.OpenPromptTuningPipeline = PublisherModel_CallToAction_RegionalResourceReferences_ToProto(mapCtx, in.OpenPromptTuningPipeline)
	out.OpenGenie = PublisherModel_CallToAction_RegionalResourceReferences_ToProto(mapCtx, in.OpenGenie)
	out.Deploy = PublisherModel_CallToAction_Deploy_ToProto(mapCtx, in.Deploy)
	out.DeployGke = PublisherModel_CallToAction_DeployGke_ToProto(mapCtx, in.DeployGke)
	out.OpenGenerationAiStudio = PublisherModel_CallToAction_RegionalResourceReferences_ToProto(mapCtx, in.OpenGenerationAiStudio)
	out.RequestAccess = PublisherModel_CallToAction_RegionalResourceReferences_ToProto(mapCtx, in.RequestAccess)
	out.OpenEvaluationPipeline = PublisherModel_CallToAction_RegionalResourceReferences_ToProto(mapCtx, in.OpenEvaluationPipeline)
	return out
}
func PublisherModel_CallToAction_Deploy_FromProto(mapCtx *direct.MapContext, in *pb.PublisherModel_CallToAction_Deploy) *krm.PublisherModel_CallToAction_Deploy {
	if in == nil {
		return nil
	}
	out := &krm.PublisherModel_CallToAction_Deploy{}
	out.DedicatedResources = DedicatedResources_FromProto(mapCtx, in.GetDedicatedResources())
	out.AutomaticResources = AutomaticResources_FromProto(mapCtx, in.GetAutomaticResources())
	out.SharedResources = direct.LazyPtr(in.GetSharedResources())
	out.ModelDisplayName = direct.LazyPtr(in.GetModelDisplayName())
	out.LargeModelReference = LargeModelReference_FromProto(mapCtx, in.GetLargeModelReference())
	out.ContainerSpec = ModelContainerSpec_FromProto(mapCtx, in.GetContainerSpec())
	out.ArtifactURI = direct.LazyPtr(in.GetArtifactUri())
	out.DeployTaskName = in.DeployTaskName
	out.DeployMetadata = PublisherModel_CallToAction_Deploy_DeployMetadata_FromProto(mapCtx, in.GetDeployMetadata())
	out.Title = direct.LazyPtr(in.GetTitle())
	out.PublicArtifactURI = direct.LazyPtr(in.GetPublicArtifactUri())
	return out
}
func PublisherModel_CallToAction_Deploy_ToProto(mapCtx *direct.MapContext, in *krm.PublisherModel_CallToAction_Deploy) *pb.PublisherModel_CallToAction_Deploy {
	if in == nil {
		return nil
	}
	out := &pb.PublisherModel_CallToAction_Deploy{}
	if oneof := DedicatedResources_ToProto(mapCtx, in.DedicatedResources); oneof != nil {
		out.PredictionResources = &pb.PublisherModel_CallToAction_Deploy_DedicatedResources{DedicatedResources: oneof}
	}
	if oneof := AutomaticResources_ToProto(mapCtx, in.AutomaticResources); oneof != nil {
		out.PredictionResources = &pb.PublisherModel_CallToAction_Deploy_AutomaticResources{AutomaticResources: oneof}
	}
	if oneof := PublisherModel_CallToAction_Deploy_SharedResources_ToProto(mapCtx, in.SharedResources); oneof != nil {
		out.PredictionResources = oneof
	}
	out.ModelDisplayName = direct.ValueOf(in.ModelDisplayName)
	out.LargeModelReference = LargeModelReference_ToProto(mapCtx, in.LargeModelReference)
	out.ContainerSpec = ModelContainerSpec_ToProto(mapCtx, in.ContainerSpec)
	out.ArtifactUri = direct.ValueOf(in.ArtifactURI)
	out.DeployTaskName = in.DeployTaskName
	if oneof := PublisherModel_CallToAction_Deploy_DeployMetadata_ToProto(mapCtx, in.DeployMetadata); oneof != nil {
		out.DeployMetadata = &pb.PublisherModel_CallToAction_Deploy_DeployMetadata_{DeployMetadata: oneof}
	}
	out.Title = direct.ValueOf(in.Title)
	out.PublicArtifactUri = direct.ValueOf(in.PublicArtifactURI)
	return out
}
func PublisherModel_CallToAction_DeployGke_FromProto(mapCtx *direct.MapContext, in *pb.PublisherModel_CallToAction_DeployGke) *krm.PublisherModel_CallToAction_DeployGke {
	if in == nil {
		return nil
	}
	out := &krm.PublisherModel_CallToAction_DeployGke{}
	out.GkeYamlConfigs = in.GkeYamlConfigs
	return out
}
func PublisherModel_CallToAction_DeployGke_ToProto(mapCtx *direct.MapContext, in *krm.PublisherModel_CallToAction_DeployGke) *pb.PublisherModel_CallToAction_DeployGke {
	if in == nil {
		return nil
	}
	out := &pb.PublisherModel_CallToAction_DeployGke{}
	out.GkeYamlConfigs = in.GkeYamlConfigs
	return out
}
func PublisherModel_CallToAction_OpenFineTuningPipelines_FromProto(mapCtx *direct.MapContext, in *pb.PublisherModel_CallToAction_OpenFineTuningPipelines) *krm.PublisherModel_CallToAction_OpenFineTuningPipelines {
	if in == nil {
		return nil
	}
	out := &krm.PublisherModel_CallToAction_OpenFineTuningPipelines{}
	out.FineTuningPipelines = direct.Slice_FromProto(mapCtx, in.FineTuningPipelines, PublisherModel_CallToAction_RegionalResourceReferences_FromProto)
	return out
}
func PublisherModel_CallToAction_OpenFineTuningPipelines_ToProto(mapCtx *direct.MapContext, in *krm.PublisherModel_CallToAction_OpenFineTuningPipelines) *pb.PublisherModel_CallToAction_OpenFineTuningPipelines {
	if in == nil {
		return nil
	}
	out := &pb.PublisherModel_CallToAction_OpenFineTuningPipelines{}
	out.FineTuningPipelines = direct.Slice_ToProto(mapCtx, in.FineTuningPipelines, PublisherModel_CallToAction_RegionalResourceReferences_ToProto)
	return out
}
func PublisherModel_CallToAction_OpenNotebooks_FromProto(mapCtx *direct.MapContext, in *pb.PublisherModel_CallToAction_OpenNotebooks) *krm.PublisherModel_CallToAction_OpenNotebooks {
	if in == nil {
		return nil
	}
	out := &krm.PublisherModel_CallToAction_OpenNotebooks{}
	out.Notebooks = direct.Slice_FromProto(mapCtx, in.Notebooks, PublisherModel_CallToAction_RegionalResourceReferences_FromProto)
	return out
}
func PublisherModel_CallToAction_OpenNotebooks_ToProto(mapCtx *direct.MapContext, in *krm.PublisherModel_CallToAction_OpenNotebooks) *pb.PublisherModel_CallToAction_OpenNotebooks {
	if in == nil {
		return nil
	}
	out := &pb.PublisherModel_CallToAction_OpenNotebooks{}
	out.Notebooks = direct.Slice_ToProto(mapCtx, in.Notebooks, PublisherModel_CallToAction_RegionalResourceReferences_ToProto)
	return out
}
func PublisherModel_CallToAction_RegionalResourceReferences_FromProto(mapCtx *direct.MapContext, in *pb.PublisherModel_CallToAction_RegionalResourceReferences) *krm.PublisherModel_CallToAction_RegionalResourceReferences {
	if in == nil {
		return nil
	}
	out := &krm.PublisherModel_CallToAction_RegionalResourceReferences{}
	// MISSING: References
	out.Title = direct.LazyPtr(in.GetTitle())
	out.ResourceTitle = in.ResourceTitle
	out.ResourceUseCase = in.ResourceUseCase
	out.ResourceDescription = in.ResourceDescription
	return out
}
func PublisherModel_CallToAction_RegionalResourceReferences_ToProto(mapCtx *direct.MapContext, in *krm.PublisherModel_CallToAction_RegionalResourceReferences) *pb.PublisherModel_CallToAction_RegionalResourceReferences {
	if in == nil {
		return nil
	}
	out := &pb.PublisherModel_CallToAction_RegionalResourceReferences{}
	// MISSING: References
	out.Title = direct.ValueOf(in.Title)
	out.ResourceTitle = in.ResourceTitle
	out.ResourceUseCase = in.ResourceUseCase
	out.ResourceDescription = in.ResourceDescription
	return out
}
func PublisherModel_CallToAction_ViewRestApi_FromProto(mapCtx *direct.MapContext, in *pb.PublisherModel_CallToAction_ViewRestApi) *krm.PublisherModel_CallToAction_ViewRestApi {
	if in == nil {
		return nil
	}
	out := &krm.PublisherModel_CallToAction_ViewRestApi{}
	out.Documentations = direct.Slice_FromProto(mapCtx, in.Documentations, PublisherModel_Documentation_FromProto)
	out.Title = direct.LazyPtr(in.GetTitle())
	return out
}
func PublisherModel_CallToAction_ViewRestApi_ToProto(mapCtx *direct.MapContext, in *krm.PublisherModel_CallToAction_ViewRestApi) *pb.PublisherModel_CallToAction_ViewRestApi {
	if in == nil {
		return nil
	}
	out := &pb.PublisherModel_CallToAction_ViewRestApi{}
	out.Documentations = direct.Slice_ToProto(mapCtx, in.Documentations, PublisherModel_Documentation_ToProto)
	out.Title = direct.ValueOf(in.Title)
	return out
}
func PublisherModel_Documentation_FromProto(mapCtx *direct.MapContext, in *pb.PublisherModel_Documentation) *krm.PublisherModel_Documentation {
	if in == nil {
		return nil
	}
	out := &krm.PublisherModel_Documentation{}
	out.Title = direct.LazyPtr(in.GetTitle())
	out.Content = direct.LazyPtr(in.GetContent())
	return out
}
func PublisherModel_Documentation_ToProto(mapCtx *direct.MapContext, in *krm.PublisherModel_Documentation) *pb.PublisherModel_Documentation {
	if in == nil {
		return nil
	}
	out := &pb.PublisherModel_Documentation{}
	out.Title = direct.ValueOf(in.Title)
	out.Content = direct.ValueOf(in.Content)
	return out
}
func PublisherModel_ResourceReference_FromProto(mapCtx *direct.MapContext, in *pb.PublisherModel_ResourceReference) *krm.PublisherModel_ResourceReference {
	if in == nil {
		return nil
	}
	out := &krm.PublisherModel_ResourceReference{}
	out.URI = direct.LazyPtr(in.GetUri())
	out.ResourceName = direct.LazyPtr(in.GetResourceName())
	out.UseCase = direct.LazyPtr(in.GetUseCase())
	out.Description = direct.LazyPtr(in.GetDescription())
	return out
}
func PublisherModel_ResourceReference_ToProto(mapCtx *direct.MapContext, in *krm.PublisherModel_ResourceReference) *pb.PublisherModel_ResourceReference {
	if in == nil {
		return nil
	}
	out := &pb.PublisherModel_ResourceReference{}
	if oneof := PublisherModel_ResourceReference_Uri_ToProto(mapCtx, in.URI); oneof != nil {
		out.Reference = oneof
	}
	if oneof := PublisherModel_ResourceReference_ResourceName_ToProto(mapCtx, in.ResourceName); oneof != nil {
		out.Reference = oneof
	}
	if oneof := PublisherModel_ResourceReference_UseCase_ToProto(mapCtx, in.UseCase); oneof != nil {
		out.Reference = oneof
	}
	if oneof := PublisherModel_ResourceReference_Description_ToProto(mapCtx, in.Description); oneof != nil {
		out.Reference = oneof
	}
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
