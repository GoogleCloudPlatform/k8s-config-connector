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
func PersistentResource_FromProto(mapCtx *direct.MapContext, in *pb.PersistentResource) *krm.PersistentResource {
	if in == nil {
		return nil
	}
	out := &krm.PersistentResource{}
	out.Name = direct.LazyPtr(in.GetName())
	out.DisplayName = direct.LazyPtr(in.GetDisplayName())
	out.ResourcePools = direct.Slice_FromProto(mapCtx, in.ResourcePools, ResourcePool_FromProto)
	// MISSING: State
	// MISSING: Error
	// MISSING: CreateTime
	// MISSING: StartTime
	// MISSING: UpdateTime
	out.Labels = in.Labels
	out.Network = direct.LazyPtr(in.GetNetwork())
	out.EncryptionSpec = EncryptionSpec_FromProto(mapCtx, in.GetEncryptionSpec())
	out.ResourceRuntimeSpec = ResourceRuntimeSpec_FromProto(mapCtx, in.GetResourceRuntimeSpec())
	// MISSING: ResourceRuntime
	out.ReservedIPRanges = in.ReservedIpRanges
	return out
}
func PersistentResource_ToProto(mapCtx *direct.MapContext, in *krm.PersistentResource) *pb.PersistentResource {
	if in == nil {
		return nil
	}
	out := &pb.PersistentResource{}
	out.Name = direct.ValueOf(in.Name)
	out.DisplayName = direct.ValueOf(in.DisplayName)
	out.ResourcePools = direct.Slice_ToProto(mapCtx, in.ResourcePools, ResourcePool_ToProto)
	// MISSING: State
	// MISSING: Error
	// MISSING: CreateTime
	// MISSING: StartTime
	// MISSING: UpdateTime
	out.Labels = in.Labels
	out.Network = direct.ValueOf(in.Network)
	out.EncryptionSpec = EncryptionSpec_ToProto(mapCtx, in.EncryptionSpec)
	out.ResourceRuntimeSpec = ResourceRuntimeSpec_ToProto(mapCtx, in.ResourceRuntimeSpec)
	// MISSING: ResourceRuntime
	out.ReservedIpRanges = in.ReservedIPRanges
	return out
}
func PersistentResourceObservedState_FromProto(mapCtx *direct.MapContext, in *pb.PersistentResource) *krm.PersistentResourceObservedState {
	if in == nil {
		return nil
	}
	out := &krm.PersistentResourceObservedState{}
	// MISSING: Name
	// MISSING: DisplayName
	out.ResourcePools = direct.Slice_FromProto(mapCtx, in.ResourcePools, ResourcePoolObservedState_FromProto)
	out.State = direct.Enum_FromProto(mapCtx, in.GetState())
	out.Error = Status_FromProto(mapCtx, in.GetError())
	out.CreateTime = direct.StringTimestamp_FromProto(mapCtx, in.GetCreateTime())
	out.StartTime = direct.StringTimestamp_FromProto(mapCtx, in.GetStartTime())
	out.UpdateTime = direct.StringTimestamp_FromProto(mapCtx, in.GetUpdateTime())
	// MISSING: Labels
	// MISSING: Network
	// MISSING: EncryptionSpec
	// MISSING: ResourceRuntimeSpec
	out.ResourceRuntime = ResourceRuntime_FromProto(mapCtx, in.GetResourceRuntime())
	// MISSING: ReservedIPRanges
	return out
}
func PersistentResourceObservedState_ToProto(mapCtx *direct.MapContext, in *krm.PersistentResourceObservedState) *pb.PersistentResource {
	if in == nil {
		return nil
	}
	out := &pb.PersistentResource{}
	// MISSING: Name
	// MISSING: DisplayName
	out.ResourcePools = direct.Slice_ToProto(mapCtx, in.ResourcePools, ResourcePoolObservedState_ToProto)
	out.State = direct.Enum_ToProto[pb.PersistentResource_State](mapCtx, in.State)
	out.Error = Status_ToProto(mapCtx, in.Error)
	out.CreateTime = direct.StringTimestamp_ToProto(mapCtx, in.CreateTime)
	out.StartTime = direct.StringTimestamp_ToProto(mapCtx, in.StartTime)
	out.UpdateTime = direct.StringTimestamp_ToProto(mapCtx, in.UpdateTime)
	// MISSING: Labels
	// MISSING: Network
	// MISSING: EncryptionSpec
	// MISSING: ResourceRuntimeSpec
	out.ResourceRuntime = ResourceRuntime_ToProto(mapCtx, in.ResourceRuntime)
	// MISSING: ReservedIPRanges
	return out
}
func RayLogsSpec_FromProto(mapCtx *direct.MapContext, in *pb.RayLogsSpec) *krm.RayLogsSpec {
	if in == nil {
		return nil
	}
	out := &krm.RayLogsSpec{}
	out.Disabled = direct.LazyPtr(in.GetDisabled())
	return out
}
func RayLogsSpec_ToProto(mapCtx *direct.MapContext, in *krm.RayLogsSpec) *pb.RayLogsSpec {
	if in == nil {
		return nil
	}
	out := &pb.RayLogsSpec{}
	out.Disabled = direct.ValueOf(in.Disabled)
	return out
}
func RayMetricSpec_FromProto(mapCtx *direct.MapContext, in *pb.RayMetricSpec) *krm.RayMetricSpec {
	if in == nil {
		return nil
	}
	out := &krm.RayMetricSpec{}
	out.Disabled = direct.LazyPtr(in.GetDisabled())
	return out
}
func RayMetricSpec_ToProto(mapCtx *direct.MapContext, in *krm.RayMetricSpec) *pb.RayMetricSpec {
	if in == nil {
		return nil
	}
	out := &pb.RayMetricSpec{}
	out.Disabled = direct.ValueOf(in.Disabled)
	return out
}
func RaySpec_FromProto(mapCtx *direct.MapContext, in *pb.RaySpec) *krm.RaySpec {
	if in == nil {
		return nil
	}
	out := &krm.RaySpec{}
	out.ImageURI = direct.LazyPtr(in.GetImageUri())
	out.ResourcePoolImages = in.ResourcePoolImages
	out.HeadNodeResourcePoolID = direct.LazyPtr(in.GetHeadNodeResourcePoolId())
	out.RayMetricSpec = RayMetricSpec_FromProto(mapCtx, in.GetRayMetricSpec())
	out.RayLogsSpec = RayLogsSpec_FromProto(mapCtx, in.GetRayLogsSpec())
	return out
}
func RaySpec_ToProto(mapCtx *direct.MapContext, in *krm.RaySpec) *pb.RaySpec {
	if in == nil {
		return nil
	}
	out := &pb.RaySpec{}
	out.ImageUri = direct.ValueOf(in.ImageURI)
	out.ResourcePoolImages = in.ResourcePoolImages
	out.HeadNodeResourcePoolId = direct.ValueOf(in.HeadNodeResourcePoolID)
	out.RayMetricSpec = RayMetricSpec_ToProto(mapCtx, in.RayMetricSpec)
	out.RayLogsSpec = RayLogsSpec_ToProto(mapCtx, in.RayLogsSpec)
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
func ResourcePool_FromProto(mapCtx *direct.MapContext, in *pb.ResourcePool) *krm.ResourcePool {
	if in == nil {
		return nil
	}
	out := &krm.ResourcePool{}
	out.ID = direct.LazyPtr(in.GetId())
	out.MachineSpec = MachineSpec_FromProto(mapCtx, in.GetMachineSpec())
	out.ReplicaCount = in.ReplicaCount
	out.DiskSpec = DiskSpec_FromProto(mapCtx, in.GetDiskSpec())
	// MISSING: UsedReplicaCount
	out.AutoscalingSpec = ResourcePool_AutoscalingSpec_FromProto(mapCtx, in.GetAutoscalingSpec())
	return out
}
func ResourcePool_ToProto(mapCtx *direct.MapContext, in *krm.ResourcePool) *pb.ResourcePool {
	if in == nil {
		return nil
	}
	out := &pb.ResourcePool{}
	out.Id = direct.ValueOf(in.ID)
	out.MachineSpec = MachineSpec_ToProto(mapCtx, in.MachineSpec)
	out.ReplicaCount = in.ReplicaCount
	out.DiskSpec = DiskSpec_ToProto(mapCtx, in.DiskSpec)
	// MISSING: UsedReplicaCount
	out.AutoscalingSpec = ResourcePool_AutoscalingSpec_ToProto(mapCtx, in.AutoscalingSpec)
	return out
}
func ResourcePoolObservedState_FromProto(mapCtx *direct.MapContext, in *pb.ResourcePool) *krm.ResourcePoolObservedState {
	if in == nil {
		return nil
	}
	out := &krm.ResourcePoolObservedState{}
	// MISSING: ID
	// MISSING: MachineSpec
	// MISSING: ReplicaCount
	// MISSING: DiskSpec
	out.UsedReplicaCount = direct.LazyPtr(in.GetUsedReplicaCount())
	// MISSING: AutoscalingSpec
	return out
}
func ResourcePoolObservedState_ToProto(mapCtx *direct.MapContext, in *krm.ResourcePoolObservedState) *pb.ResourcePool {
	if in == nil {
		return nil
	}
	out := &pb.ResourcePool{}
	// MISSING: ID
	// MISSING: MachineSpec
	// MISSING: ReplicaCount
	// MISSING: DiskSpec
	out.UsedReplicaCount = direct.ValueOf(in.UsedReplicaCount)
	// MISSING: AutoscalingSpec
	return out
}
func ResourcePool_AutoscalingSpec_FromProto(mapCtx *direct.MapContext, in *pb.ResourcePool_AutoscalingSpec) *krm.ResourcePool_AutoscalingSpec {
	if in == nil {
		return nil
	}
	out := &krm.ResourcePool_AutoscalingSpec{}
	out.MinReplicaCount = in.MinReplicaCount
	out.MaxReplicaCount = in.MaxReplicaCount
	return out
}
func ResourcePool_AutoscalingSpec_ToProto(mapCtx *direct.MapContext, in *krm.ResourcePool_AutoscalingSpec) *pb.ResourcePool_AutoscalingSpec {
	if in == nil {
		return nil
	}
	out := &pb.ResourcePool_AutoscalingSpec{}
	out.MinReplicaCount = in.MinReplicaCount
	out.MaxReplicaCount = in.MaxReplicaCount
	return out
}
func ResourceRuntime_FromProto(mapCtx *direct.MapContext, in *pb.ResourceRuntime) *krm.ResourceRuntime {
	if in == nil {
		return nil
	}
	out := &krm.ResourceRuntime{}
	// MISSING: AccessUris
	return out
}
func ResourceRuntime_ToProto(mapCtx *direct.MapContext, in *krm.ResourceRuntime) *pb.ResourceRuntime {
	if in == nil {
		return nil
	}
	out := &pb.ResourceRuntime{}
	// MISSING: AccessUris
	return out
}
func ResourceRuntimeObservedState_FromProto(mapCtx *direct.MapContext, in *pb.ResourceRuntime) *krm.ResourceRuntimeObservedState {
	if in == nil {
		return nil
	}
	out := &krm.ResourceRuntimeObservedState{}
	out.AccessUris = in.AccessUris
	return out
}
func ResourceRuntimeObservedState_ToProto(mapCtx *direct.MapContext, in *krm.ResourceRuntimeObservedState) *pb.ResourceRuntime {
	if in == nil {
		return nil
	}
	out := &pb.ResourceRuntime{}
	out.AccessUris = in.AccessUris
	return out
}
func ResourceRuntimeSpec_FromProto(mapCtx *direct.MapContext, in *pb.ResourceRuntimeSpec) *krm.ResourceRuntimeSpec {
	if in == nil {
		return nil
	}
	out := &krm.ResourceRuntimeSpec{}
	out.ServiceAccountSpec = ServiceAccountSpec_FromProto(mapCtx, in.GetServiceAccountSpec())
	out.RaySpec = RaySpec_FromProto(mapCtx, in.GetRaySpec())
	return out
}
func ResourceRuntimeSpec_ToProto(mapCtx *direct.MapContext, in *krm.ResourceRuntimeSpec) *pb.ResourceRuntimeSpec {
	if in == nil {
		return nil
	}
	out := &pb.ResourceRuntimeSpec{}
	out.ServiceAccountSpec = ServiceAccountSpec_ToProto(mapCtx, in.ServiceAccountSpec)
	out.RaySpec = RaySpec_ToProto(mapCtx, in.RaySpec)
	return out
}
func ServiceAccountSpec_FromProto(mapCtx *direct.MapContext, in *pb.ServiceAccountSpec) *krm.ServiceAccountSpec {
	if in == nil {
		return nil
	}
	out := &krm.ServiceAccountSpec{}
	out.EnableCustomServiceAccount = direct.LazyPtr(in.GetEnableCustomServiceAccount())
	out.ServiceAccount = direct.LazyPtr(in.GetServiceAccount())
	return out
}
func ServiceAccountSpec_ToProto(mapCtx *direct.MapContext, in *krm.ServiceAccountSpec) *pb.ServiceAccountSpec {
	if in == nil {
		return nil
	}
	out := &pb.ServiceAccountSpec{}
	out.EnableCustomServiceAccount = direct.ValueOf(in.EnableCustomServiceAccount)
	out.ServiceAccount = direct.ValueOf(in.ServiceAccount)
	return out
}
