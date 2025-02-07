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
func BlurBaselineConfig_FromProto(mapCtx *direct.MapContext, in *pb.BlurBaselineConfig) *krm.BlurBaselineConfig {
	if in == nil {
		return nil
	}
	out := &krm.BlurBaselineConfig{}
	out.MaxBlurSigma = direct.LazyPtr(in.GetMaxBlurSigma())
	return out
}
func BlurBaselineConfig_ToProto(mapCtx *direct.MapContext, in *krm.BlurBaselineConfig) *pb.BlurBaselineConfig {
	if in == nil {
		return nil
	}
	out := &pb.BlurBaselineConfig{}
	out.MaxBlurSigma = direct.ValueOf(in.MaxBlurSigma)
	return out
}
func ClientConnectionConfig_FromProto(mapCtx *direct.MapContext, in *pb.ClientConnectionConfig) *krm.ClientConnectionConfig {
	if in == nil {
		return nil
	}
	out := &krm.ClientConnectionConfig{}
	out.InferenceTimeout = direct.StringDuration_FromProto(mapCtx, in.GetInferenceTimeout())
	return out
}
func ClientConnectionConfig_ToProto(mapCtx *direct.MapContext, in *krm.ClientConnectionConfig) *pb.ClientConnectionConfig {
	if in == nil {
		return nil
	}
	out := &pb.ClientConnectionConfig{}
	out.InferenceTimeout = direct.StringDuration_ToProto(mapCtx, in.InferenceTimeout)
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
func DeployedModel_FromProto(mapCtx *direct.MapContext, in *pb.DeployedModel) *krm.DeployedModel {
	if in == nil {
		return nil
	}
	out := &krm.DeployedModel{}
	out.DedicatedResources = DedicatedResources_FromProto(mapCtx, in.GetDedicatedResources())
	out.AutomaticResources = AutomaticResources_FromProto(mapCtx, in.GetAutomaticResources())
	out.SharedResources = direct.LazyPtr(in.GetSharedResources())
	out.ID = direct.LazyPtr(in.GetId())
	out.Model = direct.LazyPtr(in.GetModel())
	// MISSING: ModelVersionID
	out.DisplayName = direct.LazyPtr(in.GetDisplayName())
	// MISSING: CreateTime
	out.ExplanationSpec = ExplanationSpec_FromProto(mapCtx, in.GetExplanationSpec())
	out.DisableExplanations = direct.LazyPtr(in.GetDisableExplanations())
	out.ServiceAccount = direct.LazyPtr(in.GetServiceAccount())
	out.DisableContainerLogging = direct.LazyPtr(in.GetDisableContainerLogging())
	out.EnableAccessLogging = direct.LazyPtr(in.GetEnableAccessLogging())
	// MISSING: PrivateEndpoints
	out.FasterDeploymentConfig = FasterDeploymentConfig_FromProto(mapCtx, in.GetFasterDeploymentConfig())
	// MISSING: Status
	out.SystemLabels = in.SystemLabels
	return out
}
func DeployedModel_ToProto(mapCtx *direct.MapContext, in *krm.DeployedModel) *pb.DeployedModel {
	if in == nil {
		return nil
	}
	out := &pb.DeployedModel{}
	if oneof := DedicatedResources_ToProto(mapCtx, in.DedicatedResources); oneof != nil {
		out.PredictionResources = &pb.DeployedModel_DedicatedResources{DedicatedResources: oneof}
	}
	if oneof := AutomaticResources_ToProto(mapCtx, in.AutomaticResources); oneof != nil {
		out.PredictionResources = &pb.DeployedModel_AutomaticResources{AutomaticResources: oneof}
	}
	if oneof := DeployedModel_SharedResources_ToProto(mapCtx, in.SharedResources); oneof != nil {
		out.PredictionResources = oneof
	}
	out.Id = direct.ValueOf(in.ID)
	out.Model = direct.ValueOf(in.Model)
	// MISSING: ModelVersionID
	out.DisplayName = direct.ValueOf(in.DisplayName)
	// MISSING: CreateTime
	out.ExplanationSpec = ExplanationSpec_ToProto(mapCtx, in.ExplanationSpec)
	out.DisableExplanations = direct.ValueOf(in.DisableExplanations)
	out.ServiceAccount = direct.ValueOf(in.ServiceAccount)
	out.DisableContainerLogging = direct.ValueOf(in.DisableContainerLogging)
	out.EnableAccessLogging = direct.ValueOf(in.EnableAccessLogging)
	// MISSING: PrivateEndpoints
	out.FasterDeploymentConfig = FasterDeploymentConfig_ToProto(mapCtx, in.FasterDeploymentConfig)
	// MISSING: Status
	out.SystemLabels = in.SystemLabels
	return out
}
func DeployedModelObservedState_FromProto(mapCtx *direct.MapContext, in *pb.DeployedModel) *krm.DeployedModelObservedState {
	if in == nil {
		return nil
	}
	out := &krm.DeployedModelObservedState{}
	// MISSING: DedicatedResources
	// MISSING: AutomaticResources
	// MISSING: SharedResources
	// MISSING: ID
	// MISSING: Model
	out.ModelVersionID = direct.LazyPtr(in.GetModelVersionId())
	// MISSING: DisplayName
	out.CreateTime = direct.StringTimestamp_FromProto(mapCtx, in.GetCreateTime())
	// MISSING: ExplanationSpec
	// MISSING: DisableExplanations
	// MISSING: ServiceAccount
	// MISSING: DisableContainerLogging
	// MISSING: EnableAccessLogging
	out.PrivateEndpoints = PrivateEndpoints_FromProto(mapCtx, in.GetPrivateEndpoints())
	// MISSING: FasterDeploymentConfig
	out.Status = DeployedModel_Status_FromProto(mapCtx, in.GetStatus())
	// MISSING: SystemLabels
	return out
}
func DeployedModelObservedState_ToProto(mapCtx *direct.MapContext, in *krm.DeployedModelObservedState) *pb.DeployedModel {
	if in == nil {
		return nil
	}
	out := &pb.DeployedModel{}
	// MISSING: DedicatedResources
	// MISSING: AutomaticResources
	// MISSING: SharedResources
	// MISSING: ID
	// MISSING: Model
	out.ModelVersionId = direct.ValueOf(in.ModelVersionID)
	// MISSING: DisplayName
	out.CreateTime = direct.StringTimestamp_ToProto(mapCtx, in.CreateTime)
	// MISSING: ExplanationSpec
	// MISSING: DisableExplanations
	// MISSING: ServiceAccount
	// MISSING: DisableContainerLogging
	// MISSING: EnableAccessLogging
	out.PrivateEndpoints = PrivateEndpoints_ToProto(mapCtx, in.PrivateEndpoints)
	// MISSING: FasterDeploymentConfig
	out.Status = DeployedModel_Status_ToProto(mapCtx, in.Status)
	// MISSING: SystemLabels
	return out
}
func DeployedModel_Status_FromProto(mapCtx *direct.MapContext, in *pb.DeployedModel_Status) *krm.DeployedModel_Status {
	if in == nil {
		return nil
	}
	out := &krm.DeployedModel_Status{}
	// MISSING: Message
	// MISSING: LastUpdateTime
	// MISSING: AvailableReplicaCount
	return out
}
func DeployedModel_Status_ToProto(mapCtx *direct.MapContext, in *krm.DeployedModel_Status) *pb.DeployedModel_Status {
	if in == nil {
		return nil
	}
	out := &pb.DeployedModel_Status{}
	// MISSING: Message
	// MISSING: LastUpdateTime
	// MISSING: AvailableReplicaCount
	return out
}
func DeployedModel_StatusObservedState_FromProto(mapCtx *direct.MapContext, in *pb.DeployedModel_Status) *krm.DeployedModel_StatusObservedState {
	if in == nil {
		return nil
	}
	out := &krm.DeployedModel_StatusObservedState{}
	out.Message = direct.LazyPtr(in.GetMessage())
	out.LastUpdateTime = direct.StringTimestamp_FromProto(mapCtx, in.GetLastUpdateTime())
	out.AvailableReplicaCount = direct.LazyPtr(in.GetAvailableReplicaCount())
	return out
}
func DeployedModel_StatusObservedState_ToProto(mapCtx *direct.MapContext, in *krm.DeployedModel_StatusObservedState) *pb.DeployedModel_Status {
	if in == nil {
		return nil
	}
	out := &pb.DeployedModel_Status{}
	out.Message = direct.ValueOf(in.Message)
	out.LastUpdateTime = direct.StringTimestamp_ToProto(mapCtx, in.LastUpdateTime)
	out.AvailableReplicaCount = direct.ValueOf(in.AvailableReplicaCount)
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
func Endpoint_FromProto(mapCtx *direct.MapContext, in *pb.Endpoint) *krm.Endpoint {
	if in == nil {
		return nil
	}
	out := &krm.Endpoint{}
	// MISSING: Name
	out.DisplayName = direct.LazyPtr(in.GetDisplayName())
	out.Description = direct.LazyPtr(in.GetDescription())
	// MISSING: DeployedModels
	// MISSING: TrafficSplit
	out.Etag = direct.LazyPtr(in.GetEtag())
	out.Labels = in.Labels
	// MISSING: CreateTime
	// MISSING: UpdateTime
	out.EncryptionSpec = EncryptionSpec_FromProto(mapCtx, in.GetEncryptionSpec())
	out.Network = direct.LazyPtr(in.GetNetwork())
	out.EnablePrivateServiceConnect = direct.LazyPtr(in.GetEnablePrivateServiceConnect())
	out.PrivateServiceConnectConfig = PrivateServiceConnectConfig_FromProto(mapCtx, in.GetPrivateServiceConnectConfig())
	// MISSING: ModelDeploymentMonitoringJob
	out.PredictRequestResponseLoggingConfig = PredictRequestResponseLoggingConfig_FromProto(mapCtx, in.GetPredictRequestResponseLoggingConfig())
	out.DedicatedEndpointEnabled = direct.LazyPtr(in.GetDedicatedEndpointEnabled())
	// MISSING: DedicatedEndpointDns
	out.ClientConnectionConfig = ClientConnectionConfig_FromProto(mapCtx, in.GetClientConnectionConfig())
	// MISSING: SatisfiesPzs
	// MISSING: SatisfiesPzi
	return out
}
func Endpoint_ToProto(mapCtx *direct.MapContext, in *krm.Endpoint) *pb.Endpoint {
	if in == nil {
		return nil
	}
	out := &pb.Endpoint{}
	// MISSING: Name
	out.DisplayName = direct.ValueOf(in.DisplayName)
	out.Description = direct.ValueOf(in.Description)
	// MISSING: DeployedModels
	// MISSING: TrafficSplit
	out.Etag = direct.ValueOf(in.Etag)
	out.Labels = in.Labels
	// MISSING: CreateTime
	// MISSING: UpdateTime
	out.EncryptionSpec = EncryptionSpec_ToProto(mapCtx, in.EncryptionSpec)
	out.Network = direct.ValueOf(in.Network)
	out.EnablePrivateServiceConnect = direct.ValueOf(in.EnablePrivateServiceConnect)
	out.PrivateServiceConnectConfig = PrivateServiceConnectConfig_ToProto(mapCtx, in.PrivateServiceConnectConfig)
	// MISSING: ModelDeploymentMonitoringJob
	out.PredictRequestResponseLoggingConfig = PredictRequestResponseLoggingConfig_ToProto(mapCtx, in.PredictRequestResponseLoggingConfig)
	out.DedicatedEndpointEnabled = direct.ValueOf(in.DedicatedEndpointEnabled)
	// MISSING: DedicatedEndpointDns
	out.ClientConnectionConfig = ClientConnectionConfig_ToProto(mapCtx, in.ClientConnectionConfig)
	// MISSING: SatisfiesPzs
	// MISSING: SatisfiesPzi
	return out
}
func EndpointObservedState_FromProto(mapCtx *direct.MapContext, in *pb.Endpoint) *krm.EndpointObservedState {
	if in == nil {
		return nil
	}
	out := &krm.EndpointObservedState{}
	out.Name = direct.LazyPtr(in.GetName())
	// MISSING: DisplayName
	// MISSING: Description
	out.DeployedModels = direct.Slice_FromProto(mapCtx, in.DeployedModels, DeployedModel_FromProto)
	// MISSING: TrafficSplit
	// MISSING: Etag
	// MISSING: Labels
	out.CreateTime = direct.StringTimestamp_FromProto(mapCtx, in.GetCreateTime())
	out.UpdateTime = direct.StringTimestamp_FromProto(mapCtx, in.GetUpdateTime())
	// MISSING: EncryptionSpec
	// MISSING: Network
	// MISSING: EnablePrivateServiceConnect
	out.PrivateServiceConnectConfig = PrivateServiceConnectConfigObservedState_FromProto(mapCtx, in.GetPrivateServiceConnectConfig())
	out.ModelDeploymentMonitoringJob = direct.LazyPtr(in.GetModelDeploymentMonitoringJob())
	// MISSING: PredictRequestResponseLoggingConfig
	// MISSING: DedicatedEndpointEnabled
	out.DedicatedEndpointDns = direct.LazyPtr(in.GetDedicatedEndpointDns())
	// MISSING: ClientConnectionConfig
	out.SatisfiesPzs = direct.LazyPtr(in.GetSatisfiesPzs())
	out.SatisfiesPzi = direct.LazyPtr(in.GetSatisfiesPzi())
	return out
}
func EndpointObservedState_ToProto(mapCtx *direct.MapContext, in *krm.EndpointObservedState) *pb.Endpoint {
	if in == nil {
		return nil
	}
	out := &pb.Endpoint{}
	out.Name = direct.ValueOf(in.Name)
	// MISSING: DisplayName
	// MISSING: Description
	out.DeployedModels = direct.Slice_ToProto(mapCtx, in.DeployedModels, DeployedModel_ToProto)
	// MISSING: TrafficSplit
	// MISSING: Etag
	// MISSING: Labels
	out.CreateTime = direct.StringTimestamp_ToProto(mapCtx, in.CreateTime)
	out.UpdateTime = direct.StringTimestamp_ToProto(mapCtx, in.UpdateTime)
	// MISSING: EncryptionSpec
	// MISSING: Network
	// MISSING: EnablePrivateServiceConnect
	out.PrivateServiceConnectConfig = PrivateServiceConnectConfigObservedState_ToProto(mapCtx, in.PrivateServiceConnectConfig)
	out.ModelDeploymentMonitoringJob = direct.ValueOf(in.ModelDeploymentMonitoringJob)
	// MISSING: PredictRequestResponseLoggingConfig
	// MISSING: DedicatedEndpointEnabled
	out.DedicatedEndpointDns = direct.ValueOf(in.DedicatedEndpointDns)
	// MISSING: ClientConnectionConfig
	out.SatisfiesPzs = direct.ValueOf(in.SatisfiesPzs)
	out.SatisfiesPzi = direct.ValueOf(in.SatisfiesPzi)
	return out
}
func Examples_FromProto(mapCtx *direct.MapContext, in *pb.Examples) *krm.Examples {
	if in == nil {
		return nil
	}
	out := &krm.Examples{}
	out.ExampleGcsSource = Examples_ExampleGcsSource_FromProto(mapCtx, in.GetExampleGcsSource())
	out.NearestNeighborSearchConfig = Value_FromProto(mapCtx, in.GetNearestNeighborSearchConfig())
	out.Presets = Presets_FromProto(mapCtx, in.GetPresets())
	out.NeighborCount = direct.LazyPtr(in.GetNeighborCount())
	return out
}
func Examples_ToProto(mapCtx *direct.MapContext, in *krm.Examples) *pb.Examples {
	if in == nil {
		return nil
	}
	out := &pb.Examples{}
	if oneof := Examples_ExampleGcsSource_ToProto(mapCtx, in.ExampleGcsSource); oneof != nil {
		out.Source = &pb.Examples_ExampleGcsSource_{ExampleGcsSource: oneof}
	}
	if oneof := Value_ToProto(mapCtx, in.NearestNeighborSearchConfig); oneof != nil {
		out.Config = &pb.Examples_NearestNeighborSearchConfig{NearestNeighborSearchConfig: oneof}
	}
	if oneof := Presets_ToProto(mapCtx, in.Presets); oneof != nil {
		out.Config = &pb.Examples_Presets{Presets: oneof}
	}
	out.NeighborCount = direct.ValueOf(in.NeighborCount)
	return out
}
func Examples_ExampleGcsSource_FromProto(mapCtx *direct.MapContext, in *pb.Examples_ExampleGcsSource) *krm.Examples_ExampleGcsSource {
	if in == nil {
		return nil
	}
	out := &krm.Examples_ExampleGcsSource{}
	out.DataFormat = direct.Enum_FromProto(mapCtx, in.GetDataFormat())
	out.GcsSource = GcsSource_FromProto(mapCtx, in.GetGcsSource())
	return out
}
func Examples_ExampleGcsSource_ToProto(mapCtx *direct.MapContext, in *krm.Examples_ExampleGcsSource) *pb.Examples_ExampleGcsSource {
	if in == nil {
		return nil
	}
	out := &pb.Examples_ExampleGcsSource{}
	out.DataFormat = direct.Enum_ToProto[pb.Examples_ExampleGcsSource_DataFormat](mapCtx, in.DataFormat)
	out.GcsSource = GcsSource_ToProto(mapCtx, in.GcsSource)
	return out
}
func ExplanationParameters_FromProto(mapCtx *direct.MapContext, in *pb.ExplanationParameters) *krm.ExplanationParameters {
	if in == nil {
		return nil
	}
	out := &krm.ExplanationParameters{}
	out.SampledShapleyAttribution = SampledShapleyAttribution_FromProto(mapCtx, in.GetSampledShapleyAttribution())
	out.IntegratedGradientsAttribution = IntegratedGradientsAttribution_FromProto(mapCtx, in.GetIntegratedGradientsAttribution())
	out.XraiAttribution = XraiAttribution_FromProto(mapCtx, in.GetXraiAttribution())
	out.Examples = Examples_FromProto(mapCtx, in.GetExamples())
	out.TopK = direct.LazyPtr(in.GetTopK())
	out.OutputIndices = ListValue_FromProto(mapCtx, in.GetOutputIndices())
	return out
}
func ExplanationParameters_ToProto(mapCtx *direct.MapContext, in *krm.ExplanationParameters) *pb.ExplanationParameters {
	if in == nil {
		return nil
	}
	out := &pb.ExplanationParameters{}
	if oneof := SampledShapleyAttribution_ToProto(mapCtx, in.SampledShapleyAttribution); oneof != nil {
		out.Method = &pb.ExplanationParameters_SampledShapleyAttribution{SampledShapleyAttribution: oneof}
	}
	if oneof := IntegratedGradientsAttribution_ToProto(mapCtx, in.IntegratedGradientsAttribution); oneof != nil {
		out.Method = &pb.ExplanationParameters_IntegratedGradientsAttribution{IntegratedGradientsAttribution: oneof}
	}
	if oneof := XraiAttribution_ToProto(mapCtx, in.XraiAttribution); oneof != nil {
		out.Method = &pb.ExplanationParameters_XraiAttribution{XraiAttribution: oneof}
	}
	if oneof := Examples_ToProto(mapCtx, in.Examples); oneof != nil {
		out.Method = &pb.ExplanationParameters_Examples{Examples: oneof}
	}
	out.TopK = direct.ValueOf(in.TopK)
	out.OutputIndices = ListValue_ToProto(mapCtx, in.OutputIndices)
	return out
}
func ExplanationSpec_FromProto(mapCtx *direct.MapContext, in *pb.ExplanationSpec) *krm.ExplanationSpec {
	if in == nil {
		return nil
	}
	out := &krm.ExplanationSpec{}
	out.Parameters = ExplanationParameters_FromProto(mapCtx, in.GetParameters())
	out.Metadata = ExplanationMetadata_FromProto(mapCtx, in.GetMetadata())
	return out
}
func ExplanationSpec_ToProto(mapCtx *direct.MapContext, in *krm.ExplanationSpec) *pb.ExplanationSpec {
	if in == nil {
		return nil
	}
	out := &pb.ExplanationSpec{}
	out.Parameters = ExplanationParameters_ToProto(mapCtx, in.Parameters)
	out.Metadata = ExplanationMetadata_ToProto(mapCtx, in.Metadata)
	return out
}
func FasterDeploymentConfig_FromProto(mapCtx *direct.MapContext, in *pb.FasterDeploymentConfig) *krm.FasterDeploymentConfig {
	if in == nil {
		return nil
	}
	out := &krm.FasterDeploymentConfig{}
	out.FastTryoutEnabled = direct.LazyPtr(in.GetFastTryoutEnabled())
	return out
}
func FasterDeploymentConfig_ToProto(mapCtx *direct.MapContext, in *krm.FasterDeploymentConfig) *pb.FasterDeploymentConfig {
	if in == nil {
		return nil
	}
	out := &pb.FasterDeploymentConfig{}
	out.FastTryoutEnabled = direct.ValueOf(in.FastTryoutEnabled)
	return out
}
func FeatureNoiseSigma_FromProto(mapCtx *direct.MapContext, in *pb.FeatureNoiseSigma) *krm.FeatureNoiseSigma {
	if in == nil {
		return nil
	}
	out := &krm.FeatureNoiseSigma{}
	out.NoiseSigma = direct.Slice_FromProto(mapCtx, in.NoiseSigma, FeatureNoiseSigma_NoiseSigmaForFeature_FromProto)
	return out
}
func FeatureNoiseSigma_ToProto(mapCtx *direct.MapContext, in *krm.FeatureNoiseSigma) *pb.FeatureNoiseSigma {
	if in == nil {
		return nil
	}
	out := &pb.FeatureNoiseSigma{}
	out.NoiseSigma = direct.Slice_ToProto(mapCtx, in.NoiseSigma, FeatureNoiseSigma_NoiseSigmaForFeature_ToProto)
	return out
}
func FeatureNoiseSigma_NoiseSigmaForFeature_FromProto(mapCtx *direct.MapContext, in *pb.FeatureNoiseSigma_NoiseSigmaForFeature) *krm.FeatureNoiseSigma_NoiseSigmaForFeature {
	if in == nil {
		return nil
	}
	out := &krm.FeatureNoiseSigma_NoiseSigmaForFeature{}
	out.Name = direct.LazyPtr(in.GetName())
	out.Sigma = direct.LazyPtr(in.GetSigma())
	return out
}
func FeatureNoiseSigma_NoiseSigmaForFeature_ToProto(mapCtx *direct.MapContext, in *krm.FeatureNoiseSigma_NoiseSigmaForFeature) *pb.FeatureNoiseSigma_NoiseSigmaForFeature {
	if in == nil {
		return nil
	}
	out := &pb.FeatureNoiseSigma_NoiseSigmaForFeature{}
	out.Name = direct.ValueOf(in.Name)
	out.Sigma = direct.ValueOf(in.Sigma)
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
func IntegratedGradientsAttribution_FromProto(mapCtx *direct.MapContext, in *pb.IntegratedGradientsAttribution) *krm.IntegratedGradientsAttribution {
	if in == nil {
		return nil
	}
	out := &krm.IntegratedGradientsAttribution{}
	out.StepCount = direct.LazyPtr(in.GetStepCount())
	out.SmoothGradConfig = SmoothGradConfig_FromProto(mapCtx, in.GetSmoothGradConfig())
	out.BlurBaselineConfig = BlurBaselineConfig_FromProto(mapCtx, in.GetBlurBaselineConfig())
	return out
}
func IntegratedGradientsAttribution_ToProto(mapCtx *direct.MapContext, in *krm.IntegratedGradientsAttribution) *pb.IntegratedGradientsAttribution {
	if in == nil {
		return nil
	}
	out := &pb.IntegratedGradientsAttribution{}
	out.StepCount = direct.ValueOf(in.StepCount)
	out.SmoothGradConfig = SmoothGradConfig_ToProto(mapCtx, in.SmoothGradConfig)
	out.BlurBaselineConfig = BlurBaselineConfig_ToProto(mapCtx, in.BlurBaselineConfig)
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
func PredictRequestResponseLoggingConfig_FromProto(mapCtx *direct.MapContext, in *pb.PredictRequestResponseLoggingConfig) *krm.PredictRequestResponseLoggingConfig {
	if in == nil {
		return nil
	}
	out := &krm.PredictRequestResponseLoggingConfig{}
	out.Enabled = direct.LazyPtr(in.GetEnabled())
	out.SamplingRate = direct.LazyPtr(in.GetSamplingRate())
	out.BigqueryDestination = BigQueryDestination_FromProto(mapCtx, in.GetBigqueryDestination())
	return out
}
func PredictRequestResponseLoggingConfig_ToProto(mapCtx *direct.MapContext, in *krm.PredictRequestResponseLoggingConfig) *pb.PredictRequestResponseLoggingConfig {
	if in == nil {
		return nil
	}
	out := &pb.PredictRequestResponseLoggingConfig{}
	out.Enabled = direct.ValueOf(in.Enabled)
	out.SamplingRate = direct.ValueOf(in.SamplingRate)
	out.BigqueryDestination = BigQueryDestination_ToProto(mapCtx, in.BigqueryDestination)
	return out
}
func Presets_FromProto(mapCtx *direct.MapContext, in *pb.Presets) *krm.Presets {
	if in == nil {
		return nil
	}
	out := &krm.Presets{}
	out.Query = direct.Enum_FromProto(mapCtx, in.GetQuery())
	out.Modality = direct.Enum_FromProto(mapCtx, in.GetModality())
	return out
}
func Presets_ToProto(mapCtx *direct.MapContext, in *krm.Presets) *pb.Presets {
	if in == nil {
		return nil
	}
	out := &pb.Presets{}
	if oneof := Presets_Query_ToProto(mapCtx, in.Query); oneof != nil {
		out.Query = oneof
	}
	out.Modality = direct.Enum_ToProto[pb.Presets_Modality](mapCtx, in.Modality)
	return out
}
func PrivateEndpoints_FromProto(mapCtx *direct.MapContext, in *pb.PrivateEndpoints) *krm.PrivateEndpoints {
	if in == nil {
		return nil
	}
	out := &krm.PrivateEndpoints{}
	// MISSING: PredictHTTPURI
	// MISSING: ExplainHTTPURI
	// MISSING: HealthHTTPURI
	// MISSING: ServiceAttachment
	return out
}
func PrivateEndpoints_ToProto(mapCtx *direct.MapContext, in *krm.PrivateEndpoints) *pb.PrivateEndpoints {
	if in == nil {
		return nil
	}
	out := &pb.PrivateEndpoints{}
	// MISSING: PredictHTTPURI
	// MISSING: ExplainHTTPURI
	// MISSING: HealthHTTPURI
	// MISSING: ServiceAttachment
	return out
}
func PrivateEndpointsObservedState_FromProto(mapCtx *direct.MapContext, in *pb.PrivateEndpoints) *krm.PrivateEndpointsObservedState {
	if in == nil {
		return nil
	}
	out := &krm.PrivateEndpointsObservedState{}
	out.PredictHTTPURI = direct.LazyPtr(in.GetPredictHttpUri())
	out.ExplainHTTPURI = direct.LazyPtr(in.GetExplainHttpUri())
	out.HealthHTTPURI = direct.LazyPtr(in.GetHealthHttpUri())
	out.ServiceAttachment = direct.LazyPtr(in.GetServiceAttachment())
	return out
}
func PrivateEndpointsObservedState_ToProto(mapCtx *direct.MapContext, in *krm.PrivateEndpointsObservedState) *pb.PrivateEndpoints {
	if in == nil {
		return nil
	}
	out := &pb.PrivateEndpoints{}
	out.PredictHttpUri = direct.ValueOf(in.PredictHTTPURI)
	out.ExplainHttpUri = direct.ValueOf(in.ExplainHTTPURI)
	out.HealthHttpUri = direct.ValueOf(in.HealthHTTPURI)
	out.ServiceAttachment = direct.ValueOf(in.ServiceAttachment)
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
func SampledShapleyAttribution_FromProto(mapCtx *direct.MapContext, in *pb.SampledShapleyAttribution) *krm.SampledShapleyAttribution {
	if in == nil {
		return nil
	}
	out := &krm.SampledShapleyAttribution{}
	out.PathCount = direct.LazyPtr(in.GetPathCount())
	return out
}
func SampledShapleyAttribution_ToProto(mapCtx *direct.MapContext, in *krm.SampledShapleyAttribution) *pb.SampledShapleyAttribution {
	if in == nil {
		return nil
	}
	out := &pb.SampledShapleyAttribution{}
	out.PathCount = direct.ValueOf(in.PathCount)
	return out
}
func SmoothGradConfig_FromProto(mapCtx *direct.MapContext, in *pb.SmoothGradConfig) *krm.SmoothGradConfig {
	if in == nil {
		return nil
	}
	out := &krm.SmoothGradConfig{}
	out.NoiseSigma = direct.LazyPtr(in.GetNoiseSigma())
	out.FeatureNoiseSigma = FeatureNoiseSigma_FromProto(mapCtx, in.GetFeatureNoiseSigma())
	out.NoisySampleCount = direct.LazyPtr(in.GetNoisySampleCount())
	return out
}
func SmoothGradConfig_ToProto(mapCtx *direct.MapContext, in *krm.SmoothGradConfig) *pb.SmoothGradConfig {
	if in == nil {
		return nil
	}
	out := &pb.SmoothGradConfig{}
	if oneof := SmoothGradConfig_NoiseSigma_ToProto(mapCtx, in.NoiseSigma); oneof != nil {
		out.GradientNoiseSigma = oneof
	}
	if oneof := FeatureNoiseSigma_ToProto(mapCtx, in.FeatureNoiseSigma); oneof != nil {
		out.GradientNoiseSigma = &pb.SmoothGradConfig_FeatureNoiseSigma{FeatureNoiseSigma: oneof}
	}
	out.NoisySampleCount = direct.ValueOf(in.NoisySampleCount)
	return out
}
func XraiAttribution_FromProto(mapCtx *direct.MapContext, in *pb.XraiAttribution) *krm.XraiAttribution {
	if in == nil {
		return nil
	}
	out := &krm.XraiAttribution{}
	out.StepCount = direct.LazyPtr(in.GetStepCount())
	out.SmoothGradConfig = SmoothGradConfig_FromProto(mapCtx, in.GetSmoothGradConfig())
	out.BlurBaselineConfig = BlurBaselineConfig_FromProto(mapCtx, in.GetBlurBaselineConfig())
	return out
}
func XraiAttribution_ToProto(mapCtx *direct.MapContext, in *krm.XraiAttribution) *pb.XraiAttribution {
	if in == nil {
		return nil
	}
	out := &pb.XraiAttribution{}
	out.StepCount = direct.ValueOf(in.StepCount)
	out.SmoothGradConfig = SmoothGradConfig_ToProto(mapCtx, in.SmoothGradConfig)
	out.BlurBaselineConfig = BlurBaselineConfig_ToProto(mapCtx, in.BlurBaselineConfig)
	return out
}
