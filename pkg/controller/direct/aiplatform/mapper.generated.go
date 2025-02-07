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
func Attribution_FromProto(mapCtx *direct.MapContext, in *pb.Attribution) *krm.Attribution {
	if in == nil {
		return nil
	}
	out := &krm.Attribution{}
	// MISSING: BaselineOutputValue
	// MISSING: InstanceOutputValue
	// MISSING: FeatureAttributions
	// MISSING: OutputIndex
	// MISSING: OutputDisplayName
	// MISSING: ApproximationError
	// MISSING: OutputName
	return out
}
func Attribution_ToProto(mapCtx *direct.MapContext, in *krm.Attribution) *pb.Attribution {
	if in == nil {
		return nil
	}
	out := &pb.Attribution{}
	// MISSING: BaselineOutputValue
	// MISSING: InstanceOutputValue
	// MISSING: FeatureAttributions
	// MISSING: OutputIndex
	// MISSING: OutputDisplayName
	// MISSING: ApproximationError
	// MISSING: OutputName
	return out
}
func AttributionObservedState_FromProto(mapCtx *direct.MapContext, in *pb.Attribution) *krm.AttributionObservedState {
	if in == nil {
		return nil
	}
	out := &krm.AttributionObservedState{}
	out.BaselineOutputValue = direct.LazyPtr(in.GetBaselineOutputValue())
	out.InstanceOutputValue = direct.LazyPtr(in.GetInstanceOutputValue())
	out.FeatureAttributions = Value_FromProto(mapCtx, in.GetFeatureAttributions())
	out.OutputIndex = in.OutputIndex
	out.OutputDisplayName = direct.LazyPtr(in.GetOutputDisplayName())
	out.ApproximationError = direct.LazyPtr(in.GetApproximationError())
	out.OutputName = direct.LazyPtr(in.GetOutputName())
	return out
}
func AttributionObservedState_ToProto(mapCtx *direct.MapContext, in *krm.AttributionObservedState) *pb.Attribution {
	if in == nil {
		return nil
	}
	out := &pb.Attribution{}
	out.BaselineOutputValue = direct.ValueOf(in.BaselineOutputValue)
	out.InstanceOutputValue = direct.ValueOf(in.InstanceOutputValue)
	out.FeatureAttributions = Value_ToProto(mapCtx, in.FeatureAttributions)
	out.OutputIndex = in.OutputIndex
	out.OutputDisplayName = direct.ValueOf(in.OutputDisplayName)
	out.ApproximationError = direct.ValueOf(in.ApproximationError)
	out.OutputName = direct.ValueOf(in.OutputName)
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
func ModelEvaluation_FromProto(mapCtx *direct.MapContext, in *pb.ModelEvaluation) *krm.ModelEvaluation {
	if in == nil {
		return nil
	}
	out := &krm.ModelEvaluation{}
	// MISSING: Name
	out.DisplayName = direct.LazyPtr(in.GetDisplayName())
	out.MetricsSchemaURI = direct.LazyPtr(in.GetMetricsSchemaUri())
	out.Metrics = Value_FromProto(mapCtx, in.GetMetrics())
	// MISSING: CreateTime
	out.SliceDimensions = in.SliceDimensions
	out.DataItemSchemaURI = direct.LazyPtr(in.GetDataItemSchemaUri())
	out.AnnotationSchemaURI = direct.LazyPtr(in.GetAnnotationSchemaUri())
	out.ModelExplanation = ModelExplanation_FromProto(mapCtx, in.GetModelExplanation())
	out.ExplanationSpecs = direct.Slice_FromProto(mapCtx, in.ExplanationSpecs, ModelEvaluation_ModelEvaluationExplanationSpec_FromProto)
	out.Metadata = Value_FromProto(mapCtx, in.GetMetadata())
	return out
}
func ModelEvaluation_ToProto(mapCtx *direct.MapContext, in *krm.ModelEvaluation) *pb.ModelEvaluation {
	if in == nil {
		return nil
	}
	out := &pb.ModelEvaluation{}
	// MISSING: Name
	out.DisplayName = direct.ValueOf(in.DisplayName)
	out.MetricsSchemaUri = direct.ValueOf(in.MetricsSchemaURI)
	out.Metrics = Value_ToProto(mapCtx, in.Metrics)
	// MISSING: CreateTime
	out.SliceDimensions = in.SliceDimensions
	out.DataItemSchemaUri = direct.ValueOf(in.DataItemSchemaURI)
	out.AnnotationSchemaUri = direct.ValueOf(in.AnnotationSchemaURI)
	out.ModelExplanation = ModelExplanation_ToProto(mapCtx, in.ModelExplanation)
	out.ExplanationSpecs = direct.Slice_ToProto(mapCtx, in.ExplanationSpecs, ModelEvaluation_ModelEvaluationExplanationSpec_ToProto)
	out.Metadata = Value_ToProto(mapCtx, in.Metadata)
	return out
}
func ModelEvaluationObservedState_FromProto(mapCtx *direct.MapContext, in *pb.ModelEvaluation) *krm.ModelEvaluationObservedState {
	if in == nil {
		return nil
	}
	out := &krm.ModelEvaluationObservedState{}
	out.Name = direct.LazyPtr(in.GetName())
	// MISSING: DisplayName
	// MISSING: MetricsSchemaURI
	// MISSING: Metrics
	out.CreateTime = direct.StringTimestamp_FromProto(mapCtx, in.GetCreateTime())
	// MISSING: SliceDimensions
	// MISSING: DataItemSchemaURI
	// MISSING: AnnotationSchemaURI
	out.ModelExplanation = ModelExplanationObservedState_FromProto(mapCtx, in.GetModelExplanation())
	// MISSING: ExplanationSpecs
	// MISSING: Metadata
	return out
}
func ModelEvaluationObservedState_ToProto(mapCtx *direct.MapContext, in *krm.ModelEvaluationObservedState) *pb.ModelEvaluation {
	if in == nil {
		return nil
	}
	out := &pb.ModelEvaluation{}
	out.Name = direct.ValueOf(in.Name)
	// MISSING: DisplayName
	// MISSING: MetricsSchemaURI
	// MISSING: Metrics
	out.CreateTime = direct.StringTimestamp_ToProto(mapCtx, in.CreateTime)
	// MISSING: SliceDimensions
	// MISSING: DataItemSchemaURI
	// MISSING: AnnotationSchemaURI
	out.ModelExplanation = ModelExplanationObservedState_ToProto(mapCtx, in.ModelExplanation)
	// MISSING: ExplanationSpecs
	// MISSING: Metadata
	return out
}
func ModelEvaluation_ModelEvaluationExplanationSpec_FromProto(mapCtx *direct.MapContext, in *pb.ModelEvaluation_ModelEvaluationExplanationSpec) *krm.ModelEvaluation_ModelEvaluationExplanationSpec {
	if in == nil {
		return nil
	}
	out := &krm.ModelEvaluation_ModelEvaluationExplanationSpec{}
	out.ExplanationType = direct.LazyPtr(in.GetExplanationType())
	out.ExplanationSpec = ExplanationSpec_FromProto(mapCtx, in.GetExplanationSpec())
	return out
}
func ModelEvaluation_ModelEvaluationExplanationSpec_ToProto(mapCtx *direct.MapContext, in *krm.ModelEvaluation_ModelEvaluationExplanationSpec) *pb.ModelEvaluation_ModelEvaluationExplanationSpec {
	if in == nil {
		return nil
	}
	out := &pb.ModelEvaluation_ModelEvaluationExplanationSpec{}
	out.ExplanationType = direct.ValueOf(in.ExplanationType)
	out.ExplanationSpec = ExplanationSpec_ToProto(mapCtx, in.ExplanationSpec)
	return out
}
func ModelExplanation_FromProto(mapCtx *direct.MapContext, in *pb.ModelExplanation) *krm.ModelExplanation {
	if in == nil {
		return nil
	}
	out := &krm.ModelExplanation{}
	// MISSING: MeanAttributions
	return out
}
func ModelExplanation_ToProto(mapCtx *direct.MapContext, in *krm.ModelExplanation) *pb.ModelExplanation {
	if in == nil {
		return nil
	}
	out := &pb.ModelExplanation{}
	// MISSING: MeanAttributions
	return out
}
func ModelExplanationObservedState_FromProto(mapCtx *direct.MapContext, in *pb.ModelExplanation) *krm.ModelExplanationObservedState {
	if in == nil {
		return nil
	}
	out := &krm.ModelExplanationObservedState{}
	out.MeanAttributions = direct.Slice_FromProto(mapCtx, in.MeanAttributions, Attribution_FromProto)
	return out
}
func ModelExplanationObservedState_ToProto(mapCtx *direct.MapContext, in *krm.ModelExplanationObservedState) *pb.ModelExplanation {
	if in == nil {
		return nil
	}
	out := &pb.ModelExplanation{}
	out.MeanAttributions = direct.Slice_ToProto(mapCtx, in.MeanAttributions, Attribution_ToProto)
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
