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
func PipelineJob_FromProto(mapCtx *direct.MapContext, in *pb.PipelineJob) *krm.PipelineJob {
	if in == nil {
		return nil
	}
	out := &krm.PipelineJob{}
	// MISSING: Name
	out.DisplayName = direct.LazyPtr(in.GetDisplayName())
	// MISSING: CreateTime
	// MISSING: StartTime
	// MISSING: EndTime
	// MISSING: UpdateTime
	out.PipelineSpec = PipelineSpec_FromProto(mapCtx, in.GetPipelineSpec())
	// MISSING: State
	// MISSING: JobDetail
	// MISSING: Error
	out.Labels = in.Labels
	out.RuntimeConfig = PipelineJob_RuntimeConfig_FromProto(mapCtx, in.GetRuntimeConfig())
	out.EncryptionSpec = EncryptionSpec_FromProto(mapCtx, in.GetEncryptionSpec())
	out.ServiceAccount = direct.LazyPtr(in.GetServiceAccount())
	out.Network = direct.LazyPtr(in.GetNetwork())
	out.ReservedIPRanges = in.ReservedIpRanges
	out.TemplateURI = direct.LazyPtr(in.GetTemplateUri())
	// MISSING: TemplateMetadata
	// MISSING: ScheduleName
	out.PreflightValidations = direct.LazyPtr(in.GetPreflightValidations())
	return out
}
func PipelineJob_ToProto(mapCtx *direct.MapContext, in *krm.PipelineJob) *pb.PipelineJob {
	if in == nil {
		return nil
	}
	out := &pb.PipelineJob{}
	// MISSING: Name
	out.DisplayName = direct.ValueOf(in.DisplayName)
	// MISSING: CreateTime
	// MISSING: StartTime
	// MISSING: EndTime
	// MISSING: UpdateTime
	out.PipelineSpec = PipelineSpec_ToProto(mapCtx, in.PipelineSpec)
	// MISSING: State
	// MISSING: JobDetail
	// MISSING: Error
	out.Labels = in.Labels
	out.RuntimeConfig = PipelineJob_RuntimeConfig_ToProto(mapCtx, in.RuntimeConfig)
	out.EncryptionSpec = EncryptionSpec_ToProto(mapCtx, in.EncryptionSpec)
	out.ServiceAccount = direct.ValueOf(in.ServiceAccount)
	out.Network = direct.ValueOf(in.Network)
	out.ReservedIpRanges = in.ReservedIPRanges
	out.TemplateUri = direct.ValueOf(in.TemplateURI)
	// MISSING: TemplateMetadata
	// MISSING: ScheduleName
	out.PreflightValidations = direct.ValueOf(in.PreflightValidations)
	return out
}
func PipelineJobDetail_FromProto(mapCtx *direct.MapContext, in *pb.PipelineJobDetail) *krm.PipelineJobDetail {
	if in == nil {
		return nil
	}
	out := &krm.PipelineJobDetail{}
	// MISSING: PipelineContext
	// MISSING: PipelineRunContext
	// MISSING: TaskDetails
	return out
}
func PipelineJobDetail_ToProto(mapCtx *direct.MapContext, in *krm.PipelineJobDetail) *pb.PipelineJobDetail {
	if in == nil {
		return nil
	}
	out := &pb.PipelineJobDetail{}
	// MISSING: PipelineContext
	// MISSING: PipelineRunContext
	// MISSING: TaskDetails
	return out
}
func PipelineJobDetailObservedState_FromProto(mapCtx *direct.MapContext, in *pb.PipelineJobDetail) *krm.PipelineJobDetailObservedState {
	if in == nil {
		return nil
	}
	out := &krm.PipelineJobDetailObservedState{}
	out.PipelineContext = Context_FromProto(mapCtx, in.GetPipelineContext())
	out.PipelineRunContext = Context_FromProto(mapCtx, in.GetPipelineRunContext())
	out.TaskDetails = direct.Slice_FromProto(mapCtx, in.TaskDetails, PipelineTaskDetail_FromProto)
	return out
}
func PipelineJobDetailObservedState_ToProto(mapCtx *direct.MapContext, in *krm.PipelineJobDetailObservedState) *pb.PipelineJobDetail {
	if in == nil {
		return nil
	}
	out := &pb.PipelineJobDetail{}
	out.PipelineContext = Context_ToProto(mapCtx, in.PipelineContext)
	out.PipelineRunContext = Context_ToProto(mapCtx, in.PipelineRunContext)
	out.TaskDetails = direct.Slice_ToProto(mapCtx, in.TaskDetails, PipelineTaskDetail_ToProto)
	return out
}
func PipelineJobObservedState_FromProto(mapCtx *direct.MapContext, in *pb.PipelineJob) *krm.PipelineJobObservedState {
	if in == nil {
		return nil
	}
	out := &krm.PipelineJobObservedState{}
	out.Name = direct.LazyPtr(in.GetName())
	// MISSING: DisplayName
	out.CreateTime = direct.StringTimestamp_FromProto(mapCtx, in.GetCreateTime())
	out.StartTime = direct.StringTimestamp_FromProto(mapCtx, in.GetStartTime())
	out.EndTime = direct.StringTimestamp_FromProto(mapCtx, in.GetEndTime())
	out.UpdateTime = direct.StringTimestamp_FromProto(mapCtx, in.GetUpdateTime())
	// MISSING: PipelineSpec
	out.State = direct.Enum_FromProto(mapCtx, in.GetState())
	out.JobDetail = PipelineJobDetail_FromProto(mapCtx, in.GetJobDetail())
	out.Error = Status_FromProto(mapCtx, in.GetError())
	// MISSING: Labels
	// MISSING: RuntimeConfig
	// MISSING: EncryptionSpec
	// MISSING: ServiceAccount
	// MISSING: Network
	// MISSING: ReservedIPRanges
	// MISSING: TemplateURI
	out.TemplateMetadata = PipelineTemplateMetadata_FromProto(mapCtx, in.GetTemplateMetadata())
	out.ScheduleName = direct.LazyPtr(in.GetScheduleName())
	// MISSING: PreflightValidations
	return out
}
func PipelineJobObservedState_ToProto(mapCtx *direct.MapContext, in *krm.PipelineJobObservedState) *pb.PipelineJob {
	if in == nil {
		return nil
	}
	out := &pb.PipelineJob{}
	out.Name = direct.ValueOf(in.Name)
	// MISSING: DisplayName
	out.CreateTime = direct.StringTimestamp_ToProto(mapCtx, in.CreateTime)
	out.StartTime = direct.StringTimestamp_ToProto(mapCtx, in.StartTime)
	out.EndTime = direct.StringTimestamp_ToProto(mapCtx, in.EndTime)
	out.UpdateTime = direct.StringTimestamp_ToProto(mapCtx, in.UpdateTime)
	// MISSING: PipelineSpec
	out.State = direct.Enum_ToProto[pb.PipelineState](mapCtx, in.State)
	out.JobDetail = PipelineJobDetail_ToProto(mapCtx, in.JobDetail)
	out.Error = Status_ToProto(mapCtx, in.Error)
	// MISSING: Labels
	// MISSING: RuntimeConfig
	// MISSING: EncryptionSpec
	// MISSING: ServiceAccount
	// MISSING: Network
	// MISSING: ReservedIPRanges
	// MISSING: TemplateURI
	out.TemplateMetadata = PipelineTemplateMetadata_ToProto(mapCtx, in.TemplateMetadata)
	out.ScheduleName = direct.ValueOf(in.ScheduleName)
	// MISSING: PreflightValidations
	return out
}
func PipelineJob_RuntimeConfig_FromProto(mapCtx *direct.MapContext, in *pb.PipelineJob_RuntimeConfig) *krm.PipelineJob_RuntimeConfig {
	if in == nil {
		return nil
	}
	out := &krm.PipelineJob_RuntimeConfig{}
	// MISSING: Parameters
	out.GcsOutputDirectory = direct.LazyPtr(in.GetGcsOutputDirectory())
	// MISSING: ParameterValues
	out.FailurePolicy = direct.Enum_FromProto(mapCtx, in.GetFailurePolicy())
	// MISSING: InputArtifacts
	return out
}
func PipelineJob_RuntimeConfig_ToProto(mapCtx *direct.MapContext, in *krm.PipelineJob_RuntimeConfig) *pb.PipelineJob_RuntimeConfig {
	if in == nil {
		return nil
	}
	out := &pb.PipelineJob_RuntimeConfig{}
	// MISSING: Parameters
	out.GcsOutputDirectory = direct.ValueOf(in.GcsOutputDirectory)
	// MISSING: ParameterValues
	out.FailurePolicy = direct.Enum_ToProto[pb.PipelineFailurePolicy](mapCtx, in.FailurePolicy)
	// MISSING: InputArtifacts
	return out
}
func PipelineJob_RuntimeConfig_InputArtifact_FromProto(mapCtx *direct.MapContext, in *pb.PipelineJob_RuntimeConfig_InputArtifact) *krm.PipelineJob_RuntimeConfig_InputArtifact {
	if in == nil {
		return nil
	}
	out := &krm.PipelineJob_RuntimeConfig_InputArtifact{}
	out.ArtifactID = direct.LazyPtr(in.GetArtifactId())
	return out
}
func PipelineJob_RuntimeConfig_InputArtifact_ToProto(mapCtx *direct.MapContext, in *krm.PipelineJob_RuntimeConfig_InputArtifact) *pb.PipelineJob_RuntimeConfig_InputArtifact {
	if in == nil {
		return nil
	}
	out := &pb.PipelineJob_RuntimeConfig_InputArtifact{}
	if oneof := PipelineJob_RuntimeConfig_InputArtifact_ArtifactId_ToProto(mapCtx, in.ArtifactID); oneof != nil {
		out.Kind = oneof
	}
	return out
}
func PipelineTaskDetail_FromProto(mapCtx *direct.MapContext, in *pb.PipelineTaskDetail) *krm.PipelineTaskDetail {
	if in == nil {
		return nil
	}
	out := &krm.PipelineTaskDetail{}
	// MISSING: TaskID
	// MISSING: ParentTaskID
	// MISSING: TaskName
	// MISSING: CreateTime
	// MISSING: StartTime
	// MISSING: EndTime
	// MISSING: ExecutorDetail
	// MISSING: State
	// MISSING: Execution
	// MISSING: Error
	// MISSING: PipelineTaskStatus
	// MISSING: Inputs
	// MISSING: Outputs
	return out
}
func PipelineTaskDetail_ToProto(mapCtx *direct.MapContext, in *krm.PipelineTaskDetail) *pb.PipelineTaskDetail {
	if in == nil {
		return nil
	}
	out := &pb.PipelineTaskDetail{}
	// MISSING: TaskID
	// MISSING: ParentTaskID
	// MISSING: TaskName
	// MISSING: CreateTime
	// MISSING: StartTime
	// MISSING: EndTime
	// MISSING: ExecutorDetail
	// MISSING: State
	// MISSING: Execution
	// MISSING: Error
	// MISSING: PipelineTaskStatus
	// MISSING: Inputs
	// MISSING: Outputs
	return out
}
func PipelineTaskDetailObservedState_FromProto(mapCtx *direct.MapContext, in *pb.PipelineTaskDetail) *krm.PipelineTaskDetailObservedState {
	if in == nil {
		return nil
	}
	out := &krm.PipelineTaskDetailObservedState{}
	out.TaskID = direct.LazyPtr(in.GetTaskId())
	out.ParentTaskID = direct.LazyPtr(in.GetParentTaskId())
	out.TaskName = direct.LazyPtr(in.GetTaskName())
	out.CreateTime = direct.StringTimestamp_FromProto(mapCtx, in.GetCreateTime())
	out.StartTime = direct.StringTimestamp_FromProto(mapCtx, in.GetStartTime())
	out.EndTime = direct.StringTimestamp_FromProto(mapCtx, in.GetEndTime())
	out.ExecutorDetail = PipelineTaskExecutorDetail_FromProto(mapCtx, in.GetExecutorDetail())
	out.State = direct.Enum_FromProto(mapCtx, in.GetState())
	out.Execution = Execution_FromProto(mapCtx, in.GetExecution())
	out.Error = Status_FromProto(mapCtx, in.GetError())
	out.PipelineTaskStatus = direct.Slice_FromProto(mapCtx, in.PipelineTaskStatus, PipelineTaskDetail_PipelineTaskStatus_FromProto)
	// MISSING: Inputs
	// MISSING: Outputs
	return out
}
func PipelineTaskDetailObservedState_ToProto(mapCtx *direct.MapContext, in *krm.PipelineTaskDetailObservedState) *pb.PipelineTaskDetail {
	if in == nil {
		return nil
	}
	out := &pb.PipelineTaskDetail{}
	out.TaskId = direct.ValueOf(in.TaskID)
	out.ParentTaskId = direct.ValueOf(in.ParentTaskID)
	out.TaskName = direct.ValueOf(in.TaskName)
	out.CreateTime = direct.StringTimestamp_ToProto(mapCtx, in.CreateTime)
	out.StartTime = direct.StringTimestamp_ToProto(mapCtx, in.StartTime)
	out.EndTime = direct.StringTimestamp_ToProto(mapCtx, in.EndTime)
	out.ExecutorDetail = PipelineTaskExecutorDetail_ToProto(mapCtx, in.ExecutorDetail)
	out.State = direct.Enum_ToProto[pb.PipelineTaskDetail_State](mapCtx, in.State)
	out.Execution = Execution_ToProto(mapCtx, in.Execution)
	out.Error = Status_ToProto(mapCtx, in.Error)
	out.PipelineTaskStatus = direct.Slice_ToProto(mapCtx, in.PipelineTaskStatus, PipelineTaskDetail_PipelineTaskStatus_ToProto)
	// MISSING: Inputs
	// MISSING: Outputs
	return out
}
func PipelineTaskDetail_ArtifactList_FromProto(mapCtx *direct.MapContext, in *pb.PipelineTaskDetail_ArtifactList) *krm.PipelineTaskDetail_ArtifactList {
	if in == nil {
		return nil
	}
	out := &krm.PipelineTaskDetail_ArtifactList{}
	// MISSING: Artifacts
	return out
}
func PipelineTaskDetail_ArtifactList_ToProto(mapCtx *direct.MapContext, in *krm.PipelineTaskDetail_ArtifactList) *pb.PipelineTaskDetail_ArtifactList {
	if in == nil {
		return nil
	}
	out := &pb.PipelineTaskDetail_ArtifactList{}
	// MISSING: Artifacts
	return out
}
func PipelineTaskDetail_PipelineTaskStatus_FromProto(mapCtx *direct.MapContext, in *pb.PipelineTaskDetail_PipelineTaskStatus) *krm.PipelineTaskDetail_PipelineTaskStatus {
	if in == nil {
		return nil
	}
	out := &krm.PipelineTaskDetail_PipelineTaskStatus{}
	// MISSING: UpdateTime
	// MISSING: State
	// MISSING: Error
	return out
}
func PipelineTaskDetail_PipelineTaskStatus_ToProto(mapCtx *direct.MapContext, in *krm.PipelineTaskDetail_PipelineTaskStatus) *pb.PipelineTaskDetail_PipelineTaskStatus {
	if in == nil {
		return nil
	}
	out := &pb.PipelineTaskDetail_PipelineTaskStatus{}
	// MISSING: UpdateTime
	// MISSING: State
	// MISSING: Error
	return out
}
func PipelineTaskDetail_PipelineTaskStatusObservedState_FromProto(mapCtx *direct.MapContext, in *pb.PipelineTaskDetail_PipelineTaskStatus) *krm.PipelineTaskDetail_PipelineTaskStatusObservedState {
	if in == nil {
		return nil
	}
	out := &krm.PipelineTaskDetail_PipelineTaskStatusObservedState{}
	out.UpdateTime = direct.StringTimestamp_FromProto(mapCtx, in.GetUpdateTime())
	out.State = direct.Enum_FromProto(mapCtx, in.GetState())
	out.Error = Status_FromProto(mapCtx, in.GetError())
	return out
}
func PipelineTaskDetail_PipelineTaskStatusObservedState_ToProto(mapCtx *direct.MapContext, in *krm.PipelineTaskDetail_PipelineTaskStatusObservedState) *pb.PipelineTaskDetail_PipelineTaskStatus {
	if in == nil {
		return nil
	}
	out := &pb.PipelineTaskDetail_PipelineTaskStatus{}
	out.UpdateTime = direct.StringTimestamp_ToProto(mapCtx, in.UpdateTime)
	out.State = direct.Enum_ToProto[pb.PipelineTaskDetail_State](mapCtx, in.State)
	out.Error = Status_ToProto(mapCtx, in.Error)
	return out
}
func PipelineTaskExecutorDetail_FromProto(mapCtx *direct.MapContext, in *pb.PipelineTaskExecutorDetail) *krm.PipelineTaskExecutorDetail {
	if in == nil {
		return nil
	}
	out := &krm.PipelineTaskExecutorDetail{}
	// MISSING: ContainerDetail
	// MISSING: CustomJobDetail
	return out
}
func PipelineTaskExecutorDetail_ToProto(mapCtx *direct.MapContext, in *krm.PipelineTaskExecutorDetail) *pb.PipelineTaskExecutorDetail {
	if in == nil {
		return nil
	}
	out := &pb.PipelineTaskExecutorDetail{}
	// MISSING: ContainerDetail
	// MISSING: CustomJobDetail
	return out
}
func PipelineTaskExecutorDetailObservedState_FromProto(mapCtx *direct.MapContext, in *pb.PipelineTaskExecutorDetail) *krm.PipelineTaskExecutorDetailObservedState {
	if in == nil {
		return nil
	}
	out := &krm.PipelineTaskExecutorDetailObservedState{}
	out.ContainerDetail = PipelineTaskExecutorDetail_ContainerDetail_FromProto(mapCtx, in.GetContainerDetail())
	out.CustomJobDetail = PipelineTaskExecutorDetail_CustomJobDetail_FromProto(mapCtx, in.GetCustomJobDetail())
	return out
}
func PipelineTaskExecutorDetailObservedState_ToProto(mapCtx *direct.MapContext, in *krm.PipelineTaskExecutorDetailObservedState) *pb.PipelineTaskExecutorDetail {
	if in == nil {
		return nil
	}
	out := &pb.PipelineTaskExecutorDetail{}
	if oneof := PipelineTaskExecutorDetail_ContainerDetail_ToProto(mapCtx, in.ContainerDetail); oneof != nil {
		out.Details = &pb.PipelineTaskExecutorDetail_ContainerDetail_{ContainerDetail: oneof}
	}
	if oneof := PipelineTaskExecutorDetail_CustomJobDetail_ToProto(mapCtx, in.CustomJobDetail); oneof != nil {
		out.Details = &pb.PipelineTaskExecutorDetail_CustomJobDetail_{CustomJobDetail: oneof}
	}
	return out
}
func PipelineTaskExecutorDetail_ContainerDetail_FromProto(mapCtx *direct.MapContext, in *pb.PipelineTaskExecutorDetail_ContainerDetail) *krm.PipelineTaskExecutorDetail_ContainerDetail {
	if in == nil {
		return nil
	}
	out := &krm.PipelineTaskExecutorDetail_ContainerDetail{}
	// MISSING: MainJob
	// MISSING: PreCachingCheckJob
	// MISSING: FailedMainJobs
	// MISSING: FailedPreCachingCheckJobs
	return out
}
func PipelineTaskExecutorDetail_ContainerDetail_ToProto(mapCtx *direct.MapContext, in *krm.PipelineTaskExecutorDetail_ContainerDetail) *pb.PipelineTaskExecutorDetail_ContainerDetail {
	if in == nil {
		return nil
	}
	out := &pb.PipelineTaskExecutorDetail_ContainerDetail{}
	// MISSING: MainJob
	// MISSING: PreCachingCheckJob
	// MISSING: FailedMainJobs
	// MISSING: FailedPreCachingCheckJobs
	return out
}
func PipelineTaskExecutorDetail_ContainerDetailObservedState_FromProto(mapCtx *direct.MapContext, in *pb.PipelineTaskExecutorDetail_ContainerDetail) *krm.PipelineTaskExecutorDetail_ContainerDetailObservedState {
	if in == nil {
		return nil
	}
	out := &krm.PipelineTaskExecutorDetail_ContainerDetailObservedState{}
	out.MainJob = direct.LazyPtr(in.GetMainJob())
	out.PreCachingCheckJob = direct.LazyPtr(in.GetPreCachingCheckJob())
	out.FailedMainJobs = in.FailedMainJobs
	out.FailedPreCachingCheckJobs = in.FailedPreCachingCheckJobs
	return out
}
func PipelineTaskExecutorDetail_ContainerDetailObservedState_ToProto(mapCtx *direct.MapContext, in *krm.PipelineTaskExecutorDetail_ContainerDetailObservedState) *pb.PipelineTaskExecutorDetail_ContainerDetail {
	if in == nil {
		return nil
	}
	out := &pb.PipelineTaskExecutorDetail_ContainerDetail{}
	out.MainJob = direct.ValueOf(in.MainJob)
	out.PreCachingCheckJob = direct.ValueOf(in.PreCachingCheckJob)
	out.FailedMainJobs = in.FailedMainJobs
	out.FailedPreCachingCheckJobs = in.FailedPreCachingCheckJobs
	return out
}
func PipelineTaskExecutorDetail_CustomJobDetail_FromProto(mapCtx *direct.MapContext, in *pb.PipelineTaskExecutorDetail_CustomJobDetail) *krm.PipelineTaskExecutorDetail_CustomJobDetail {
	if in == nil {
		return nil
	}
	out := &krm.PipelineTaskExecutorDetail_CustomJobDetail{}
	// MISSING: Job
	// MISSING: FailedJobs
	return out
}
func PipelineTaskExecutorDetail_CustomJobDetail_ToProto(mapCtx *direct.MapContext, in *krm.PipelineTaskExecutorDetail_CustomJobDetail) *pb.PipelineTaskExecutorDetail_CustomJobDetail {
	if in == nil {
		return nil
	}
	out := &pb.PipelineTaskExecutorDetail_CustomJobDetail{}
	// MISSING: Job
	// MISSING: FailedJobs
	return out
}
func PipelineTaskExecutorDetail_CustomJobDetailObservedState_FromProto(mapCtx *direct.MapContext, in *pb.PipelineTaskExecutorDetail_CustomJobDetail) *krm.PipelineTaskExecutorDetail_CustomJobDetailObservedState {
	if in == nil {
		return nil
	}
	out := &krm.PipelineTaskExecutorDetail_CustomJobDetailObservedState{}
	out.Job = direct.LazyPtr(in.GetJob())
	out.FailedJobs = in.FailedJobs
	return out
}
func PipelineTaskExecutorDetail_CustomJobDetailObservedState_ToProto(mapCtx *direct.MapContext, in *krm.PipelineTaskExecutorDetail_CustomJobDetailObservedState) *pb.PipelineTaskExecutorDetail_CustomJobDetail {
	if in == nil {
		return nil
	}
	out := &pb.PipelineTaskExecutorDetail_CustomJobDetail{}
	out.Job = direct.ValueOf(in.Job)
	out.FailedJobs = in.FailedJobs
	return out
}
func Value_FromProto(mapCtx *direct.MapContext, in *pb.Value) *krm.Value {
	if in == nil {
		return nil
	}
	out := &krm.Value{}
	out.IntValue = direct.LazyPtr(in.GetIntValue())
	out.DoubleValue = direct.LazyPtr(in.GetDoubleValue())
	out.StringValue = direct.LazyPtr(in.GetStringValue())
	return out
}
func Value_ToProto(mapCtx *direct.MapContext, in *krm.Value) *pb.Value {
	if in == nil {
		return nil
	}
	out := &pb.Value{}
	if oneof := Value_IntValue_ToProto(mapCtx, in.IntValue); oneof != nil {
		out.Value = oneof
	}
	if oneof := Value_DoubleValue_ToProto(mapCtx, in.DoubleValue); oneof != nil {
		out.Value = oneof
	}
	if oneof := Value_StringValue_ToProto(mapCtx, in.StringValue); oneof != nil {
		out.Value = oneof
	}
	return out
}
