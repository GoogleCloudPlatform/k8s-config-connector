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
func ContainerSpec_FromProto(mapCtx *direct.MapContext, in *pb.ContainerSpec) *krm.ContainerSpec {
	if in == nil {
		return nil
	}
	out := &krm.ContainerSpec{}
	out.ImageURI = direct.LazyPtr(in.GetImageUri())
	out.Command = in.Command
	out.Args = in.Args
	out.Env = direct.Slice_FromProto(mapCtx, in.Env, EnvVar_FromProto)
	return out
}
func ContainerSpec_ToProto(mapCtx *direct.MapContext, in *krm.ContainerSpec) *pb.ContainerSpec {
	if in == nil {
		return nil
	}
	out := &pb.ContainerSpec{}
	out.ImageUri = direct.ValueOf(in.ImageURI)
	out.Command = in.Command
	out.Args = in.Args
	out.Env = direct.Slice_ToProto(mapCtx, in.Env, EnvVar_ToProto)
	return out
}
func CustomJobSpec_FromProto(mapCtx *direct.MapContext, in *pb.CustomJobSpec) *krm.CustomJobSpec {
	if in == nil {
		return nil
	}
	out := &krm.CustomJobSpec{}
	out.PersistentResourceID = direct.LazyPtr(in.GetPersistentResourceId())
	out.WorkerPoolSpecs = direct.Slice_FromProto(mapCtx, in.WorkerPoolSpecs, WorkerPoolSpec_FromProto)
	out.Scheduling = Scheduling_FromProto(mapCtx, in.GetScheduling())
	out.ServiceAccount = direct.LazyPtr(in.GetServiceAccount())
	out.Network = direct.LazyPtr(in.GetNetwork())
	out.ReservedIPRanges = in.ReservedIpRanges
	out.BaseOutputDirectory = GcsDestination_FromProto(mapCtx, in.GetBaseOutputDirectory())
	out.ProtectedArtifactLocationID = direct.LazyPtr(in.GetProtectedArtifactLocationId())
	out.Tensorboard = direct.LazyPtr(in.GetTensorboard())
	out.EnableWebAccess = direct.LazyPtr(in.GetEnableWebAccess())
	out.EnableDashboardAccess = direct.LazyPtr(in.GetEnableDashboardAccess())
	out.Experiment = direct.LazyPtr(in.GetExperiment())
	out.ExperimentRun = direct.LazyPtr(in.GetExperimentRun())
	out.Models = in.Models
	return out
}
func CustomJobSpec_ToProto(mapCtx *direct.MapContext, in *krm.CustomJobSpec) *pb.CustomJobSpec {
	if in == nil {
		return nil
	}
	out := &pb.CustomJobSpec{}
	out.PersistentResourceId = direct.ValueOf(in.PersistentResourceID)
	out.WorkerPoolSpecs = direct.Slice_ToProto(mapCtx, in.WorkerPoolSpecs, WorkerPoolSpec_ToProto)
	out.Scheduling = Scheduling_ToProto(mapCtx, in.Scheduling)
	out.ServiceAccount = direct.ValueOf(in.ServiceAccount)
	out.Network = direct.ValueOf(in.Network)
	out.ReservedIpRanges = in.ReservedIPRanges
	out.BaseOutputDirectory = GcsDestination_ToProto(mapCtx, in.BaseOutputDirectory)
	out.ProtectedArtifactLocationId = direct.ValueOf(in.ProtectedArtifactLocationID)
	out.Tensorboard = direct.ValueOf(in.Tensorboard)
	out.EnableWebAccess = direct.ValueOf(in.EnableWebAccess)
	out.EnableDashboardAccess = direct.ValueOf(in.EnableDashboardAccess)
	out.Experiment = direct.ValueOf(in.Experiment)
	out.ExperimentRun = direct.ValueOf(in.ExperimentRun)
	out.Models = in.Models
	return out
}
func DiskSpec_FromProto(mapCtx *direct.MapContext, in *pb.DiskSpec) *krm.DiskSpec {
	if in == nil {
		return nil
	}
	out := &krm.DiskSpec{}
	out.BootDiskType = direct.LazyPtr(in.GetBootDiskType())
	out.BootDiskSizeGB = direct.LazyPtr(in.GetBootDiskSizeGb())
	return out
}
func DiskSpec_ToProto(mapCtx *direct.MapContext, in *krm.DiskSpec) *pb.DiskSpec {
	if in == nil {
		return nil
	}
	out := &pb.DiskSpec{}
	out.BootDiskType = direct.ValueOf(in.BootDiskType)
	out.BootDiskSizeGb = direct.ValueOf(in.BootDiskSizeGB)
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
func HyperparameterTuningJob_FromProto(mapCtx *direct.MapContext, in *pb.HyperparameterTuningJob) *krm.HyperparameterTuningJob {
	if in == nil {
		return nil
	}
	out := &krm.HyperparameterTuningJob{}
	// MISSING: Name
	out.DisplayName = direct.LazyPtr(in.GetDisplayName())
	out.StudySpec = StudySpec_FromProto(mapCtx, in.GetStudySpec())
	out.MaxTrialCount = direct.LazyPtr(in.GetMaxTrialCount())
	out.ParallelTrialCount = direct.LazyPtr(in.GetParallelTrialCount())
	out.MaxFailedTrialCount = direct.LazyPtr(in.GetMaxFailedTrialCount())
	out.TrialJobSpec = CustomJobSpec_FromProto(mapCtx, in.GetTrialJobSpec())
	// MISSING: Trials
	// MISSING: State
	// MISSING: CreateTime
	// MISSING: StartTime
	// MISSING: EndTime
	// MISSING: UpdateTime
	// MISSING: Error
	out.Labels = in.Labels
	out.EncryptionSpec = EncryptionSpec_FromProto(mapCtx, in.GetEncryptionSpec())
	// MISSING: SatisfiesPzs
	// MISSING: SatisfiesPzi
	return out
}
func HyperparameterTuningJob_ToProto(mapCtx *direct.MapContext, in *krm.HyperparameterTuningJob) *pb.HyperparameterTuningJob {
	if in == nil {
		return nil
	}
	out := &pb.HyperparameterTuningJob{}
	// MISSING: Name
	out.DisplayName = direct.ValueOf(in.DisplayName)
	out.StudySpec = StudySpec_ToProto(mapCtx, in.StudySpec)
	out.MaxTrialCount = direct.ValueOf(in.MaxTrialCount)
	out.ParallelTrialCount = direct.ValueOf(in.ParallelTrialCount)
	out.MaxFailedTrialCount = direct.ValueOf(in.MaxFailedTrialCount)
	out.TrialJobSpec = CustomJobSpec_ToProto(mapCtx, in.TrialJobSpec)
	// MISSING: Trials
	// MISSING: State
	// MISSING: CreateTime
	// MISSING: StartTime
	// MISSING: EndTime
	// MISSING: UpdateTime
	// MISSING: Error
	out.Labels = in.Labels
	out.EncryptionSpec = EncryptionSpec_ToProto(mapCtx, in.EncryptionSpec)
	// MISSING: SatisfiesPzs
	// MISSING: SatisfiesPzi
	return out
}
func HyperparameterTuningJobObservedState_FromProto(mapCtx *direct.MapContext, in *pb.HyperparameterTuningJob) *krm.HyperparameterTuningJobObservedState {
	if in == nil {
		return nil
	}
	out := &krm.HyperparameterTuningJobObservedState{}
	out.Name = direct.LazyPtr(in.GetName())
	// MISSING: DisplayName
	// MISSING: StudySpec
	// MISSING: MaxTrialCount
	// MISSING: ParallelTrialCount
	// MISSING: MaxFailedTrialCount
	// MISSING: TrialJobSpec
	out.Trials = direct.Slice_FromProto(mapCtx, in.Trials, Trial_FromProto)
	out.State = direct.Enum_FromProto(mapCtx, in.GetState())
	out.CreateTime = direct.StringTimestamp_FromProto(mapCtx, in.GetCreateTime())
	out.StartTime = direct.StringTimestamp_FromProto(mapCtx, in.GetStartTime())
	out.EndTime = direct.StringTimestamp_FromProto(mapCtx, in.GetEndTime())
	out.UpdateTime = direct.StringTimestamp_FromProto(mapCtx, in.GetUpdateTime())
	out.Error = Status_FromProto(mapCtx, in.GetError())
	// MISSING: Labels
	// MISSING: EncryptionSpec
	out.SatisfiesPzs = direct.LazyPtr(in.GetSatisfiesPzs())
	out.SatisfiesPzi = direct.LazyPtr(in.GetSatisfiesPzi())
	return out
}
func HyperparameterTuningJobObservedState_ToProto(mapCtx *direct.MapContext, in *krm.HyperparameterTuningJobObservedState) *pb.HyperparameterTuningJob {
	if in == nil {
		return nil
	}
	out := &pb.HyperparameterTuningJob{}
	out.Name = direct.ValueOf(in.Name)
	// MISSING: DisplayName
	// MISSING: StudySpec
	// MISSING: MaxTrialCount
	// MISSING: ParallelTrialCount
	// MISSING: MaxFailedTrialCount
	// MISSING: TrialJobSpec
	out.Trials = direct.Slice_ToProto(mapCtx, in.Trials, Trial_ToProto)
	out.State = direct.Enum_ToProto[pb.JobState](mapCtx, in.State)
	out.CreateTime = direct.StringTimestamp_ToProto(mapCtx, in.CreateTime)
	out.StartTime = direct.StringTimestamp_ToProto(mapCtx, in.StartTime)
	out.EndTime = direct.StringTimestamp_ToProto(mapCtx, in.EndTime)
	out.UpdateTime = direct.StringTimestamp_ToProto(mapCtx, in.UpdateTime)
	out.Error = Status_ToProto(mapCtx, in.Error)
	// MISSING: Labels
	// MISSING: EncryptionSpec
	out.SatisfiesPzs = direct.ValueOf(in.SatisfiesPzs)
	out.SatisfiesPzi = direct.ValueOf(in.SatisfiesPzi)
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
func Measurement_FromProto(mapCtx *direct.MapContext, in *pb.Measurement) *krm.Measurement {
	if in == nil {
		return nil
	}
	out := &krm.Measurement{}
	// MISSING: ElapsedDuration
	// MISSING: StepCount
	// MISSING: Metrics
	return out
}
func Measurement_ToProto(mapCtx *direct.MapContext, in *krm.Measurement) *pb.Measurement {
	if in == nil {
		return nil
	}
	out := &pb.Measurement{}
	// MISSING: ElapsedDuration
	// MISSING: StepCount
	// MISSING: Metrics
	return out
}
func MeasurementObservedState_FromProto(mapCtx *direct.MapContext, in *pb.Measurement) *krm.MeasurementObservedState {
	if in == nil {
		return nil
	}
	out := &krm.MeasurementObservedState{}
	out.ElapsedDuration = direct.StringDuration_FromProto(mapCtx, in.GetElapsedDuration())
	out.StepCount = direct.LazyPtr(in.GetStepCount())
	out.Metrics = direct.Slice_FromProto(mapCtx, in.Metrics, Measurement_Metric_FromProto)
	return out
}
func MeasurementObservedState_ToProto(mapCtx *direct.MapContext, in *krm.MeasurementObservedState) *pb.Measurement {
	if in == nil {
		return nil
	}
	out := &pb.Measurement{}
	out.ElapsedDuration = direct.StringDuration_ToProto(mapCtx, in.ElapsedDuration)
	out.StepCount = direct.ValueOf(in.StepCount)
	out.Metrics = direct.Slice_ToProto(mapCtx, in.Metrics, Measurement_Metric_ToProto)
	return out
}
func Measurement_Metric_FromProto(mapCtx *direct.MapContext, in *pb.Measurement_Metric) *krm.Measurement_Metric {
	if in == nil {
		return nil
	}
	out := &krm.Measurement_Metric{}
	// MISSING: MetricID
	// MISSING: Value
	return out
}
func Measurement_Metric_ToProto(mapCtx *direct.MapContext, in *krm.Measurement_Metric) *pb.Measurement_Metric {
	if in == nil {
		return nil
	}
	out := &pb.Measurement_Metric{}
	// MISSING: MetricID
	// MISSING: Value
	return out
}
func Measurement_MetricObservedState_FromProto(mapCtx *direct.MapContext, in *pb.Measurement_Metric) *krm.Measurement_MetricObservedState {
	if in == nil {
		return nil
	}
	out := &krm.Measurement_MetricObservedState{}
	out.MetricID = direct.LazyPtr(in.GetMetricId())
	out.Value = direct.LazyPtr(in.GetValue())
	return out
}
func Measurement_MetricObservedState_ToProto(mapCtx *direct.MapContext, in *krm.Measurement_MetricObservedState) *pb.Measurement_Metric {
	if in == nil {
		return nil
	}
	out := &pb.Measurement_Metric{}
	out.MetricId = direct.ValueOf(in.MetricID)
	out.Value = direct.ValueOf(in.Value)
	return out
}
func NfsMount_FromProto(mapCtx *direct.MapContext, in *pb.NfsMount) *krm.NfsMount {
	if in == nil {
		return nil
	}
	out := &krm.NfsMount{}
	out.Server = direct.LazyPtr(in.GetServer())
	out.Path = direct.LazyPtr(in.GetPath())
	out.MountPoint = direct.LazyPtr(in.GetMountPoint())
	return out
}
func NfsMount_ToProto(mapCtx *direct.MapContext, in *krm.NfsMount) *pb.NfsMount {
	if in == nil {
		return nil
	}
	out := &pb.NfsMount{}
	out.Server = direct.ValueOf(in.Server)
	out.Path = direct.ValueOf(in.Path)
	out.MountPoint = direct.ValueOf(in.MountPoint)
	return out
}
func PythonPackageSpec_FromProto(mapCtx *direct.MapContext, in *pb.PythonPackageSpec) *krm.PythonPackageSpec {
	if in == nil {
		return nil
	}
	out := &krm.PythonPackageSpec{}
	out.ExecutorImageURI = direct.LazyPtr(in.GetExecutorImageUri())
	out.PackageUris = in.PackageUris
	out.PythonModule = direct.LazyPtr(in.GetPythonModule())
	out.Args = in.Args
	out.Env = direct.Slice_FromProto(mapCtx, in.Env, EnvVar_FromProto)
	return out
}
func PythonPackageSpec_ToProto(mapCtx *direct.MapContext, in *krm.PythonPackageSpec) *pb.PythonPackageSpec {
	if in == nil {
		return nil
	}
	out := &pb.PythonPackageSpec{}
	out.ExecutorImageUri = direct.ValueOf(in.ExecutorImageURI)
	out.PackageUris = in.PackageUris
	out.PythonModule = direct.ValueOf(in.PythonModule)
	out.Args = in.Args
	out.Env = direct.Slice_ToProto(mapCtx, in.Env, EnvVar_ToProto)
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
func Scheduling_FromProto(mapCtx *direct.MapContext, in *pb.Scheduling) *krm.Scheduling {
	if in == nil {
		return nil
	}
	out := &krm.Scheduling{}
	out.Timeout = direct.StringDuration_FromProto(mapCtx, in.GetTimeout())
	out.RestartJobOnWorkerRestart = direct.LazyPtr(in.GetRestartJobOnWorkerRestart())
	out.Strategy = direct.Enum_FromProto(mapCtx, in.GetStrategy())
	out.DisableRetries = direct.LazyPtr(in.GetDisableRetries())
	out.MaxWaitDuration = direct.StringDuration_FromProto(mapCtx, in.GetMaxWaitDuration())
	return out
}
func Scheduling_ToProto(mapCtx *direct.MapContext, in *krm.Scheduling) *pb.Scheduling {
	if in == nil {
		return nil
	}
	out := &pb.Scheduling{}
	out.Timeout = direct.StringDuration_ToProto(mapCtx, in.Timeout)
	out.RestartJobOnWorkerRestart = direct.ValueOf(in.RestartJobOnWorkerRestart)
	out.Strategy = direct.Enum_ToProto[pb.Scheduling_Strategy](mapCtx, in.Strategy)
	out.DisableRetries = direct.ValueOf(in.DisableRetries)
	out.MaxWaitDuration = direct.StringDuration_ToProto(mapCtx, in.MaxWaitDuration)
	return out
}
func StudySpec_FromProto(mapCtx *direct.MapContext, in *pb.StudySpec) *krm.StudySpec {
	if in == nil {
		return nil
	}
	out := &krm.StudySpec{}
	out.DecayCurveStoppingSpec = StudySpec_DecayCurveAutomatedStoppingSpec_FromProto(mapCtx, in.GetDecayCurveStoppingSpec())
	out.MedianAutomatedStoppingSpec = StudySpec_MedianAutomatedStoppingSpec_FromProto(mapCtx, in.GetMedianAutomatedStoppingSpec())
	out.ConvexAutomatedStoppingSpec = StudySpec_ConvexAutomatedStoppingSpec_FromProto(mapCtx, in.GetConvexAutomatedStoppingSpec())
	out.Metrics = direct.Slice_FromProto(mapCtx, in.Metrics, StudySpec_MetricSpec_FromProto)
	out.Parameters = direct.Slice_FromProto(mapCtx, in.Parameters, StudySpec_ParameterSpec_FromProto)
	out.Algorithm = direct.Enum_FromProto(mapCtx, in.GetAlgorithm())
	out.ObservationNoise = direct.Enum_FromProto(mapCtx, in.GetObservationNoise())
	out.MeasurementSelectionType = direct.Enum_FromProto(mapCtx, in.GetMeasurementSelectionType())
	out.StudyStoppingConfig = StudySpec_StudyStoppingConfig_FromProto(mapCtx, in.GetStudyStoppingConfig())
	return out
}
func StudySpec_ToProto(mapCtx *direct.MapContext, in *krm.StudySpec) *pb.StudySpec {
	if in == nil {
		return nil
	}
	out := &pb.StudySpec{}
	if oneof := StudySpec_DecayCurveAutomatedStoppingSpec_ToProto(mapCtx, in.DecayCurveStoppingSpec); oneof != nil {
		out.AutomatedStoppingSpec = &pb.StudySpec_DecayCurveStoppingSpec{DecayCurveStoppingSpec: oneof}
	}
	if oneof := StudySpec_MedianAutomatedStoppingSpec_ToProto(mapCtx, in.MedianAutomatedStoppingSpec); oneof != nil {
		out.AutomatedStoppingSpec = &pb.StudySpec_MedianAutomatedStoppingSpec_{MedianAutomatedStoppingSpec: oneof}
	}
	if oneof := StudySpec_ConvexAutomatedStoppingSpec_ToProto(mapCtx, in.ConvexAutomatedStoppingSpec); oneof != nil {
		out.AutomatedStoppingSpec = &pb.StudySpec_ConvexAutomatedStoppingSpec_{ConvexAutomatedStoppingSpec: oneof}
	}
	out.Metrics = direct.Slice_ToProto(mapCtx, in.Metrics, StudySpec_MetricSpec_ToProto)
	out.Parameters = direct.Slice_ToProto(mapCtx, in.Parameters, StudySpec_ParameterSpec_ToProto)
	out.Algorithm = direct.Enum_ToProto[pb.StudySpec_Algorithm](mapCtx, in.Algorithm)
	out.ObservationNoise = direct.Enum_ToProto[pb.StudySpec_ObservationNoise](mapCtx, in.ObservationNoise)
	out.MeasurementSelectionType = direct.Enum_ToProto[pb.StudySpec_MeasurementSelectionType](mapCtx, in.MeasurementSelectionType)
	if oneof := StudySpec_StudyStoppingConfig_ToProto(mapCtx, in.StudyStoppingConfig); oneof != nil {
		out.StudyStoppingConfig = &pb.StudySpec_StudyStoppingConfig_{StudyStoppingConfig: oneof}
	}
	return out
}
func StudySpec_ConvexAutomatedStoppingSpec_FromProto(mapCtx *direct.MapContext, in *pb.StudySpec_ConvexAutomatedStoppingSpec) *krm.StudySpec_ConvexAutomatedStoppingSpec {
	if in == nil {
		return nil
	}
	out := &krm.StudySpec_ConvexAutomatedStoppingSpec{}
	out.MaxStepCount = direct.LazyPtr(in.GetMaxStepCount())
	out.MinStepCount = direct.LazyPtr(in.GetMinStepCount())
	out.MinMeasurementCount = direct.LazyPtr(in.GetMinMeasurementCount())
	out.LearningRateParameterName = direct.LazyPtr(in.GetLearningRateParameterName())
	out.UseElapsedDuration = direct.LazyPtr(in.GetUseElapsedDuration())
	out.UpdateAllStoppedTrials = in.UpdateAllStoppedTrials
	return out
}
func StudySpec_ConvexAutomatedStoppingSpec_ToProto(mapCtx *direct.MapContext, in *krm.StudySpec_ConvexAutomatedStoppingSpec) *pb.StudySpec_ConvexAutomatedStoppingSpec {
	if in == nil {
		return nil
	}
	out := &pb.StudySpec_ConvexAutomatedStoppingSpec{}
	out.MaxStepCount = direct.ValueOf(in.MaxStepCount)
	out.MinStepCount = direct.ValueOf(in.MinStepCount)
	out.MinMeasurementCount = direct.ValueOf(in.MinMeasurementCount)
	out.LearningRateParameterName = direct.ValueOf(in.LearningRateParameterName)
	out.UseElapsedDuration = direct.ValueOf(in.UseElapsedDuration)
	out.UpdateAllStoppedTrials = in.UpdateAllStoppedTrials
	return out
}
func StudySpec_DecayCurveAutomatedStoppingSpec_FromProto(mapCtx *direct.MapContext, in *pb.StudySpec_DecayCurveAutomatedStoppingSpec) *krm.StudySpec_DecayCurveAutomatedStoppingSpec {
	if in == nil {
		return nil
	}
	out := &krm.StudySpec_DecayCurveAutomatedStoppingSpec{}
	out.UseElapsedDuration = direct.LazyPtr(in.GetUseElapsedDuration())
	return out
}
func StudySpec_DecayCurveAutomatedStoppingSpec_ToProto(mapCtx *direct.MapContext, in *krm.StudySpec_DecayCurveAutomatedStoppingSpec) *pb.StudySpec_DecayCurveAutomatedStoppingSpec {
	if in == nil {
		return nil
	}
	out := &pb.StudySpec_DecayCurveAutomatedStoppingSpec{}
	out.UseElapsedDuration = direct.ValueOf(in.UseElapsedDuration)
	return out
}
func StudySpec_MedianAutomatedStoppingSpec_FromProto(mapCtx *direct.MapContext, in *pb.StudySpec_MedianAutomatedStoppingSpec) *krm.StudySpec_MedianAutomatedStoppingSpec {
	if in == nil {
		return nil
	}
	out := &krm.StudySpec_MedianAutomatedStoppingSpec{}
	out.UseElapsedDuration = direct.LazyPtr(in.GetUseElapsedDuration())
	return out
}
func StudySpec_MedianAutomatedStoppingSpec_ToProto(mapCtx *direct.MapContext, in *krm.StudySpec_MedianAutomatedStoppingSpec) *pb.StudySpec_MedianAutomatedStoppingSpec {
	if in == nil {
		return nil
	}
	out := &pb.StudySpec_MedianAutomatedStoppingSpec{}
	out.UseElapsedDuration = direct.ValueOf(in.UseElapsedDuration)
	return out
}
func StudySpec_MetricSpec_FromProto(mapCtx *direct.MapContext, in *pb.StudySpec_MetricSpec) *krm.StudySpec_MetricSpec {
	if in == nil {
		return nil
	}
	out := &krm.StudySpec_MetricSpec{}
	out.MetricID = direct.LazyPtr(in.GetMetricId())
	out.Goal = direct.Enum_FromProto(mapCtx, in.GetGoal())
	out.SafetyConfig = StudySpec_MetricSpec_SafetyMetricConfig_FromProto(mapCtx, in.GetSafetyConfig())
	return out
}
func StudySpec_MetricSpec_ToProto(mapCtx *direct.MapContext, in *krm.StudySpec_MetricSpec) *pb.StudySpec_MetricSpec {
	if in == nil {
		return nil
	}
	out := &pb.StudySpec_MetricSpec{}
	out.MetricId = direct.ValueOf(in.MetricID)
	out.Goal = direct.Enum_ToProto[pb.StudySpec_MetricSpec_GoalType](mapCtx, in.Goal)
	if oneof := StudySpec_MetricSpec_SafetyMetricConfig_ToProto(mapCtx, in.SafetyConfig); oneof != nil {
		out.SafetyConfig = &pb.StudySpec_MetricSpec_SafetyConfig{SafetyConfig: oneof}
	}
	return out
}
func StudySpec_MetricSpec_SafetyMetricConfig_FromProto(mapCtx *direct.MapContext, in *pb.StudySpec_MetricSpec_SafetyMetricConfig) *krm.StudySpec_MetricSpec_SafetyMetricConfig {
	if in == nil {
		return nil
	}
	out := &krm.StudySpec_MetricSpec_SafetyMetricConfig{}
	out.SafetyThreshold = direct.LazyPtr(in.GetSafetyThreshold())
	out.DesiredMinSafeTrialsFraction = in.DesiredMinSafeTrialsFraction
	return out
}
func StudySpec_MetricSpec_SafetyMetricConfig_ToProto(mapCtx *direct.MapContext, in *krm.StudySpec_MetricSpec_SafetyMetricConfig) *pb.StudySpec_MetricSpec_SafetyMetricConfig {
	if in == nil {
		return nil
	}
	out := &pb.StudySpec_MetricSpec_SafetyMetricConfig{}
	out.SafetyThreshold = direct.ValueOf(in.SafetyThreshold)
	out.DesiredMinSafeTrialsFraction = in.DesiredMinSafeTrialsFraction
	return out
}
func StudySpec_ParameterSpec_FromProto(mapCtx *direct.MapContext, in *pb.StudySpec_ParameterSpec) *krm.StudySpec_ParameterSpec {
	if in == nil {
		return nil
	}
	out := &krm.StudySpec_ParameterSpec{}
	out.DoubleValueSpec = StudySpec_ParameterSpec_DoubleValueSpec_FromProto(mapCtx, in.GetDoubleValueSpec())
	out.IntegerValueSpec = StudySpec_ParameterSpec_IntegerValueSpec_FromProto(mapCtx, in.GetIntegerValueSpec())
	out.CategoricalValueSpec = StudySpec_ParameterSpec_CategoricalValueSpec_FromProto(mapCtx, in.GetCategoricalValueSpec())
	out.DiscreteValueSpec = StudySpec_ParameterSpec_DiscreteValueSpec_FromProto(mapCtx, in.GetDiscreteValueSpec())
	out.ParameterID = direct.LazyPtr(in.GetParameterId())
	out.ScaleType = direct.Enum_FromProto(mapCtx, in.GetScaleType())
	out.ConditionalParameterSpecs = direct.Slice_FromProto(mapCtx, in.ConditionalParameterSpecs, StudySpec_ParameterSpec_ConditionalParameterSpec_FromProto)
	return out
}
func StudySpec_ParameterSpec_ToProto(mapCtx *direct.MapContext, in *krm.StudySpec_ParameterSpec) *pb.StudySpec_ParameterSpec {
	if in == nil {
		return nil
	}
	out := &pb.StudySpec_ParameterSpec{}
	if oneof := StudySpec_ParameterSpec_DoubleValueSpec_ToProto(mapCtx, in.DoubleValueSpec); oneof != nil {
		out.ParameterValueSpec = &pb.StudySpec_ParameterSpec_DoubleValueSpec_{DoubleValueSpec: oneof}
	}
	if oneof := StudySpec_ParameterSpec_IntegerValueSpec_ToProto(mapCtx, in.IntegerValueSpec); oneof != nil {
		out.ParameterValueSpec = &pb.StudySpec_ParameterSpec_IntegerValueSpec_{IntegerValueSpec: oneof}
	}
	if oneof := StudySpec_ParameterSpec_CategoricalValueSpec_ToProto(mapCtx, in.CategoricalValueSpec); oneof != nil {
		out.ParameterValueSpec = &pb.StudySpec_ParameterSpec_CategoricalValueSpec_{CategoricalValueSpec: oneof}
	}
	if oneof := StudySpec_ParameterSpec_DiscreteValueSpec_ToProto(mapCtx, in.DiscreteValueSpec); oneof != nil {
		out.ParameterValueSpec = &pb.StudySpec_ParameterSpec_DiscreteValueSpec_{DiscreteValueSpec: oneof}
	}
	out.ParameterId = direct.ValueOf(in.ParameterID)
	out.ScaleType = direct.Enum_ToProto[pb.StudySpec_ParameterSpec_ScaleType](mapCtx, in.ScaleType)
	out.ConditionalParameterSpecs = direct.Slice_ToProto(mapCtx, in.ConditionalParameterSpecs, StudySpec_ParameterSpec_ConditionalParameterSpec_ToProto)
	return out
}
func StudySpec_ParameterSpec_CategoricalValueSpec_FromProto(mapCtx *direct.MapContext, in *pb.StudySpec_ParameterSpec_CategoricalValueSpec) *krm.StudySpec_ParameterSpec_CategoricalValueSpec {
	if in == nil {
		return nil
	}
	out := &krm.StudySpec_ParameterSpec_CategoricalValueSpec{}
	out.Values = in.Values
	out.DefaultValue = in.DefaultValue
	return out
}
func StudySpec_ParameterSpec_CategoricalValueSpec_ToProto(mapCtx *direct.MapContext, in *krm.StudySpec_ParameterSpec_CategoricalValueSpec) *pb.StudySpec_ParameterSpec_CategoricalValueSpec {
	if in == nil {
		return nil
	}
	out := &pb.StudySpec_ParameterSpec_CategoricalValueSpec{}
	out.Values = in.Values
	out.DefaultValue = in.DefaultValue
	return out
}
func StudySpec_ParameterSpec_ConditionalParameterSpec_FromProto(mapCtx *direct.MapContext, in *pb.StudySpec_ParameterSpec_ConditionalParameterSpec) *krm.StudySpec_ParameterSpec_ConditionalParameterSpec {
	if in == nil {
		return nil
	}
	out := &krm.StudySpec_ParameterSpec_ConditionalParameterSpec{}
	out.ParentDiscreteValues = StudySpec_ParameterSpec_ConditionalParameterSpec_DiscreteValueCondition_FromProto(mapCtx, in.GetParentDiscreteValues())
	out.ParentIntValues = StudySpec_ParameterSpec_ConditionalParameterSpec_IntValueCondition_FromProto(mapCtx, in.GetParentIntValues())
	out.ParentCategoricalValues = StudySpec_ParameterSpec_ConditionalParameterSpec_CategoricalValueCondition_FromProto(mapCtx, in.GetParentCategoricalValues())
	out.ParameterSpec = StudySpec_ParameterSpec_FromProto(mapCtx, in.GetParameterSpec())
	return out
}
func StudySpec_ParameterSpec_ConditionalParameterSpec_ToProto(mapCtx *direct.MapContext, in *krm.StudySpec_ParameterSpec_ConditionalParameterSpec) *pb.StudySpec_ParameterSpec_ConditionalParameterSpec {
	if in == nil {
		return nil
	}
	out := &pb.StudySpec_ParameterSpec_ConditionalParameterSpec{}
	if oneof := StudySpec_ParameterSpec_ConditionalParameterSpec_DiscreteValueCondition_ToProto(mapCtx, in.ParentDiscreteValues); oneof != nil {
		out.ParentValueCondition = &pb.StudySpec_ParameterSpec_ConditionalParameterSpec_ParentDiscreteValues{ParentDiscreteValues: oneof}
	}
	if oneof := StudySpec_ParameterSpec_ConditionalParameterSpec_IntValueCondition_ToProto(mapCtx, in.ParentIntValues); oneof != nil {
		out.ParentValueCondition = &pb.StudySpec_ParameterSpec_ConditionalParameterSpec_ParentIntValues{ParentIntValues: oneof}
	}
	if oneof := StudySpec_ParameterSpec_ConditionalParameterSpec_CategoricalValueCondition_ToProto(mapCtx, in.ParentCategoricalValues); oneof != nil {
		out.ParentValueCondition = &pb.StudySpec_ParameterSpec_ConditionalParameterSpec_ParentCategoricalValues{ParentCategoricalValues: oneof}
	}
	out.ParameterSpec = StudySpec_ParameterSpec_ToProto(mapCtx, in.ParameterSpec)
	return out
}
func StudySpec_ParameterSpec_ConditionalParameterSpec_CategoricalValueCondition_FromProto(mapCtx *direct.MapContext, in *pb.StudySpec_ParameterSpec_ConditionalParameterSpec_CategoricalValueCondition) *krm.StudySpec_ParameterSpec_ConditionalParameterSpec_CategoricalValueCondition {
	if in == nil {
		return nil
	}
	out := &krm.StudySpec_ParameterSpec_ConditionalParameterSpec_CategoricalValueCondition{}
	out.Values = in.Values
	return out
}
func StudySpec_ParameterSpec_ConditionalParameterSpec_CategoricalValueCondition_ToProto(mapCtx *direct.MapContext, in *krm.StudySpec_ParameterSpec_ConditionalParameterSpec_CategoricalValueCondition) *pb.StudySpec_ParameterSpec_ConditionalParameterSpec_CategoricalValueCondition {
	if in == nil {
		return nil
	}
	out := &pb.StudySpec_ParameterSpec_ConditionalParameterSpec_CategoricalValueCondition{}
	out.Values = in.Values
	return out
}
func StudySpec_ParameterSpec_ConditionalParameterSpec_DiscreteValueCondition_FromProto(mapCtx *direct.MapContext, in *pb.StudySpec_ParameterSpec_ConditionalParameterSpec_DiscreteValueCondition) *krm.StudySpec_ParameterSpec_ConditionalParameterSpec_DiscreteValueCondition {
	if in == nil {
		return nil
	}
	out := &krm.StudySpec_ParameterSpec_ConditionalParameterSpec_DiscreteValueCondition{}
	out.Values = in.Values
	return out
}
func StudySpec_ParameterSpec_ConditionalParameterSpec_DiscreteValueCondition_ToProto(mapCtx *direct.MapContext, in *krm.StudySpec_ParameterSpec_ConditionalParameterSpec_DiscreteValueCondition) *pb.StudySpec_ParameterSpec_ConditionalParameterSpec_DiscreteValueCondition {
	if in == nil {
		return nil
	}
	out := &pb.StudySpec_ParameterSpec_ConditionalParameterSpec_DiscreteValueCondition{}
	out.Values = in.Values
	return out
}
func StudySpec_ParameterSpec_ConditionalParameterSpec_IntValueCondition_FromProto(mapCtx *direct.MapContext, in *pb.StudySpec_ParameterSpec_ConditionalParameterSpec_IntValueCondition) *krm.StudySpec_ParameterSpec_ConditionalParameterSpec_IntValueCondition {
	if in == nil {
		return nil
	}
	out := &krm.StudySpec_ParameterSpec_ConditionalParameterSpec_IntValueCondition{}
	out.Values = in.Values
	return out
}
func StudySpec_ParameterSpec_ConditionalParameterSpec_IntValueCondition_ToProto(mapCtx *direct.MapContext, in *krm.StudySpec_ParameterSpec_ConditionalParameterSpec_IntValueCondition) *pb.StudySpec_ParameterSpec_ConditionalParameterSpec_IntValueCondition {
	if in == nil {
		return nil
	}
	out := &pb.StudySpec_ParameterSpec_ConditionalParameterSpec_IntValueCondition{}
	out.Values = in.Values
	return out
}
func StudySpec_ParameterSpec_DiscreteValueSpec_FromProto(mapCtx *direct.MapContext, in *pb.StudySpec_ParameterSpec_DiscreteValueSpec) *krm.StudySpec_ParameterSpec_DiscreteValueSpec {
	if in == nil {
		return nil
	}
	out := &krm.StudySpec_ParameterSpec_DiscreteValueSpec{}
	out.Values = in.Values
	out.DefaultValue = in.DefaultValue
	return out
}
func StudySpec_ParameterSpec_DiscreteValueSpec_ToProto(mapCtx *direct.MapContext, in *krm.StudySpec_ParameterSpec_DiscreteValueSpec) *pb.StudySpec_ParameterSpec_DiscreteValueSpec {
	if in == nil {
		return nil
	}
	out := &pb.StudySpec_ParameterSpec_DiscreteValueSpec{}
	out.Values = in.Values
	out.DefaultValue = in.DefaultValue
	return out
}
func StudySpec_ParameterSpec_DoubleValueSpec_FromProto(mapCtx *direct.MapContext, in *pb.StudySpec_ParameterSpec_DoubleValueSpec) *krm.StudySpec_ParameterSpec_DoubleValueSpec {
	if in == nil {
		return nil
	}
	out := &krm.StudySpec_ParameterSpec_DoubleValueSpec{}
	out.MinValue = direct.LazyPtr(in.GetMinValue())
	out.MaxValue = direct.LazyPtr(in.GetMaxValue())
	out.DefaultValue = in.DefaultValue
	return out
}
func StudySpec_ParameterSpec_DoubleValueSpec_ToProto(mapCtx *direct.MapContext, in *krm.StudySpec_ParameterSpec_DoubleValueSpec) *pb.StudySpec_ParameterSpec_DoubleValueSpec {
	if in == nil {
		return nil
	}
	out := &pb.StudySpec_ParameterSpec_DoubleValueSpec{}
	out.MinValue = direct.ValueOf(in.MinValue)
	out.MaxValue = direct.ValueOf(in.MaxValue)
	out.DefaultValue = in.DefaultValue
	return out
}
func StudySpec_ParameterSpec_IntegerValueSpec_FromProto(mapCtx *direct.MapContext, in *pb.StudySpec_ParameterSpec_IntegerValueSpec) *krm.StudySpec_ParameterSpec_IntegerValueSpec {
	if in == nil {
		return nil
	}
	out := &krm.StudySpec_ParameterSpec_IntegerValueSpec{}
	out.MinValue = direct.LazyPtr(in.GetMinValue())
	out.MaxValue = direct.LazyPtr(in.GetMaxValue())
	out.DefaultValue = in.DefaultValue
	return out
}
func StudySpec_ParameterSpec_IntegerValueSpec_ToProto(mapCtx *direct.MapContext, in *krm.StudySpec_ParameterSpec_IntegerValueSpec) *pb.StudySpec_ParameterSpec_IntegerValueSpec {
	if in == nil {
		return nil
	}
	out := &pb.StudySpec_ParameterSpec_IntegerValueSpec{}
	out.MinValue = direct.ValueOf(in.MinValue)
	out.MaxValue = direct.ValueOf(in.MaxValue)
	out.DefaultValue = in.DefaultValue
	return out
}
func StudySpec_StudyStoppingConfig_FromProto(mapCtx *direct.MapContext, in *pb.StudySpec_StudyStoppingConfig) *krm.StudySpec_StudyStoppingConfig {
	if in == nil {
		return nil
	}
	out := &krm.StudySpec_StudyStoppingConfig{}
	out.ShouldStopAsap = direct.BoolValue_FromProto(mapCtx, in.GetShouldStopAsap())
	out.MinimumRuntimeConstraint = StudyTimeConstraint_FromProto(mapCtx, in.GetMinimumRuntimeConstraint())
	out.MaximumRuntimeConstraint = StudyTimeConstraint_FromProto(mapCtx, in.GetMaximumRuntimeConstraint())
	out.MinNumTrials = Int32Value_FromProto(mapCtx, in.GetMinNumTrials())
	out.MaxNumTrials = Int32Value_FromProto(mapCtx, in.GetMaxNumTrials())
	out.MaxNumTrialsNoProgress = Int32Value_FromProto(mapCtx, in.GetMaxNumTrialsNoProgress())
	out.MaxDurationNoProgress = direct.StringDuration_FromProto(mapCtx, in.GetMaxDurationNoProgress())
	return out
}
func StudySpec_StudyStoppingConfig_ToProto(mapCtx *direct.MapContext, in *krm.StudySpec_StudyStoppingConfig) *pb.StudySpec_StudyStoppingConfig {
	if in == nil {
		return nil
	}
	out := &pb.StudySpec_StudyStoppingConfig{}
	out.ShouldStopAsap = direct.BoolValue_ToProto(mapCtx, in.ShouldStopAsap)
	out.MinimumRuntimeConstraint = StudyTimeConstraint_ToProto(mapCtx, in.MinimumRuntimeConstraint)
	out.MaximumRuntimeConstraint = StudyTimeConstraint_ToProto(mapCtx, in.MaximumRuntimeConstraint)
	out.MinNumTrials = Int32Value_ToProto(mapCtx, in.MinNumTrials)
	out.MaxNumTrials = Int32Value_ToProto(mapCtx, in.MaxNumTrials)
	out.MaxNumTrialsNoProgress = Int32Value_ToProto(mapCtx, in.MaxNumTrialsNoProgress)
	out.MaxDurationNoProgress = direct.StringDuration_ToProto(mapCtx, in.MaxDurationNoProgress)
	return out
}
func StudyTimeConstraint_FromProto(mapCtx *direct.MapContext, in *pb.StudyTimeConstraint) *krm.StudyTimeConstraint {
	if in == nil {
		return nil
	}
	out := &krm.StudyTimeConstraint{}
	out.MaxDuration = direct.StringDuration_FromProto(mapCtx, in.GetMaxDuration())
	out.EndTime = direct.StringTimestamp_FromProto(mapCtx, in.GetEndTime())
	return out
}
func StudyTimeConstraint_ToProto(mapCtx *direct.MapContext, in *krm.StudyTimeConstraint) *pb.StudyTimeConstraint {
	if in == nil {
		return nil
	}
	out := &pb.StudyTimeConstraint{}
	if oneof := direct.StringDuration_ToProto(mapCtx, in.MaxDuration); oneof != nil {
		out.Constraint = &pb.StudyTimeConstraint_MaxDuration{MaxDuration: oneof}
	}
	if oneof := direct.StringTimestamp_ToProto(mapCtx, in.EndTime); oneof != nil {
		out.Constraint = &pb.StudyTimeConstraint_EndTime{EndTime: oneof}
	}
	return out
}
func Trial_FromProto(mapCtx *direct.MapContext, in *pb.Trial) *krm.Trial {
	if in == nil {
		return nil
	}
	out := &krm.Trial{}
	// MISSING: Name
	// MISSING: ID
	// MISSING: State
	// MISSING: Parameters
	// MISSING: FinalMeasurement
	// MISSING: Measurements
	// MISSING: StartTime
	// MISSING: EndTime
	// MISSING: ClientID
	// MISSING: InfeasibleReason
	// MISSING: CustomJob
	// MISSING: WebAccessUris
	return out
}
func Trial_ToProto(mapCtx *direct.MapContext, in *krm.Trial) *pb.Trial {
	if in == nil {
		return nil
	}
	out := &pb.Trial{}
	// MISSING: Name
	// MISSING: ID
	// MISSING: State
	// MISSING: Parameters
	// MISSING: FinalMeasurement
	// MISSING: Measurements
	// MISSING: StartTime
	// MISSING: EndTime
	// MISSING: ClientID
	// MISSING: InfeasibleReason
	// MISSING: CustomJob
	// MISSING: WebAccessUris
	return out
}
func TrialObservedState_FromProto(mapCtx *direct.MapContext, in *pb.Trial) *krm.TrialObservedState {
	if in == nil {
		return nil
	}
	out := &krm.TrialObservedState{}
	out.Name = direct.LazyPtr(in.GetName())
	out.ID = direct.LazyPtr(in.GetId())
	out.State = direct.Enum_FromProto(mapCtx, in.GetState())
	out.Parameters = direct.Slice_FromProto(mapCtx, in.Parameters, Trial_Parameter_FromProto)
	out.FinalMeasurement = Measurement_FromProto(mapCtx, in.GetFinalMeasurement())
	out.Measurements = direct.Slice_FromProto(mapCtx, in.Measurements, Measurement_FromProto)
	out.StartTime = direct.StringTimestamp_FromProto(mapCtx, in.GetStartTime())
	out.EndTime = direct.StringTimestamp_FromProto(mapCtx, in.GetEndTime())
	out.ClientID = direct.LazyPtr(in.GetClientId())
	out.InfeasibleReason = direct.LazyPtr(in.GetInfeasibleReason())
	out.CustomJob = direct.LazyPtr(in.GetCustomJob())
	out.WebAccessUris = in.WebAccessUris
	return out
}
func TrialObservedState_ToProto(mapCtx *direct.MapContext, in *krm.TrialObservedState) *pb.Trial {
	if in == nil {
		return nil
	}
	out := &pb.Trial{}
	out.Name = direct.ValueOf(in.Name)
	out.Id = direct.ValueOf(in.ID)
	out.State = direct.Enum_ToProto[pb.Trial_State](mapCtx, in.State)
	out.Parameters = direct.Slice_ToProto(mapCtx, in.Parameters, Trial_Parameter_ToProto)
	out.FinalMeasurement = Measurement_ToProto(mapCtx, in.FinalMeasurement)
	out.Measurements = direct.Slice_ToProto(mapCtx, in.Measurements, Measurement_ToProto)
	out.StartTime = direct.StringTimestamp_ToProto(mapCtx, in.StartTime)
	out.EndTime = direct.StringTimestamp_ToProto(mapCtx, in.EndTime)
	out.ClientId = direct.ValueOf(in.ClientID)
	out.InfeasibleReason = direct.ValueOf(in.InfeasibleReason)
	out.CustomJob = direct.ValueOf(in.CustomJob)
	out.WebAccessUris = in.WebAccessUris
	return out
}
func Trial_Parameter_FromProto(mapCtx *direct.MapContext, in *pb.Trial_Parameter) *krm.Trial_Parameter {
	if in == nil {
		return nil
	}
	out := &krm.Trial_Parameter{}
	// MISSING: ParameterID
	// MISSING: Value
	return out
}
func Trial_Parameter_ToProto(mapCtx *direct.MapContext, in *krm.Trial_Parameter) *pb.Trial_Parameter {
	if in == nil {
		return nil
	}
	out := &pb.Trial_Parameter{}
	// MISSING: ParameterID
	// MISSING: Value
	return out
}
func Trial_ParameterObservedState_FromProto(mapCtx *direct.MapContext, in *pb.Trial_Parameter) *krm.Trial_ParameterObservedState {
	if in == nil {
		return nil
	}
	out := &krm.Trial_ParameterObservedState{}
	out.ParameterID = direct.LazyPtr(in.GetParameterId())
	out.Value = Value_FromProto(mapCtx, in.GetValue())
	return out
}
func Trial_ParameterObservedState_ToProto(mapCtx *direct.MapContext, in *krm.Trial_ParameterObservedState) *pb.Trial_Parameter {
	if in == nil {
		return nil
	}
	out := &pb.Trial_Parameter{}
	out.ParameterId = direct.ValueOf(in.ParameterID)
	out.Value = Value_ToProto(mapCtx, in.Value)
	return out
}
func WorkerPoolSpec_FromProto(mapCtx *direct.MapContext, in *pb.WorkerPoolSpec) *krm.WorkerPoolSpec {
	if in == nil {
		return nil
	}
	out := &krm.WorkerPoolSpec{}
	out.ContainerSpec = ContainerSpec_FromProto(mapCtx, in.GetContainerSpec())
	out.PythonPackageSpec = PythonPackageSpec_FromProto(mapCtx, in.GetPythonPackageSpec())
	out.MachineSpec = MachineSpec_FromProto(mapCtx, in.GetMachineSpec())
	out.ReplicaCount = direct.LazyPtr(in.GetReplicaCount())
	out.NfsMounts = direct.Slice_FromProto(mapCtx, in.NfsMounts, NfsMount_FromProto)
	out.DiskSpec = DiskSpec_FromProto(mapCtx, in.GetDiskSpec())
	return out
}
func WorkerPoolSpec_ToProto(mapCtx *direct.MapContext, in *krm.WorkerPoolSpec) *pb.WorkerPoolSpec {
	if in == nil {
		return nil
	}
	out := &pb.WorkerPoolSpec{}
	if oneof := ContainerSpec_ToProto(mapCtx, in.ContainerSpec); oneof != nil {
		out.Task = &pb.WorkerPoolSpec_ContainerSpec{ContainerSpec: oneof}
	}
	if oneof := PythonPackageSpec_ToProto(mapCtx, in.PythonPackageSpec); oneof != nil {
		out.Task = &pb.WorkerPoolSpec_PythonPackageSpec{PythonPackageSpec: oneof}
	}
	out.MachineSpec = MachineSpec_ToProto(mapCtx, in.MachineSpec)
	out.ReplicaCount = direct.ValueOf(in.ReplicaCount)
	out.NfsMounts = direct.Slice_ToProto(mapCtx, in.NfsMounts, NfsMount_ToProto)
	out.DiskSpec = DiskSpec_ToProto(mapCtx, in.DiskSpec)
	return out
}
