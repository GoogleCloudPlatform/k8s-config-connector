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
func BatchDedicatedResources_FromProto(mapCtx *direct.MapContext, in *pb.BatchDedicatedResources) *krm.BatchDedicatedResources {
	if in == nil {
		return nil
	}
	out := &krm.BatchDedicatedResources{}
	out.MachineSpec = MachineSpec_FromProto(mapCtx, in.GetMachineSpec())
	out.StartingReplicaCount = direct.LazyPtr(in.GetStartingReplicaCount())
	out.MaxReplicaCount = direct.LazyPtr(in.GetMaxReplicaCount())
	return out
}
func BatchDedicatedResources_ToProto(mapCtx *direct.MapContext, in *krm.BatchDedicatedResources) *pb.BatchDedicatedResources {
	if in == nil {
		return nil
	}
	out := &pb.BatchDedicatedResources{}
	out.MachineSpec = MachineSpec_ToProto(mapCtx, in.MachineSpec)
	out.StartingReplicaCount = direct.ValueOf(in.StartingReplicaCount)
	out.MaxReplicaCount = direct.ValueOf(in.MaxReplicaCount)
	return out
}
func BatchPredictionJob_FromProto(mapCtx *direct.MapContext, in *pb.BatchPredictionJob) *krm.BatchPredictionJob {
	if in == nil {
		return nil
	}
	out := &krm.BatchPredictionJob{}
	// MISSING: Name
	out.DisplayName = direct.LazyPtr(in.GetDisplayName())
	out.Model = direct.LazyPtr(in.GetModel())
	// MISSING: ModelVersionID
	out.UnmanagedContainerModel = UnmanagedContainerModel_FromProto(mapCtx, in.GetUnmanagedContainerModel())
	out.InputConfig = BatchPredictionJob_InputConfig_FromProto(mapCtx, in.GetInputConfig())
	out.InstanceConfig = BatchPredictionJob_InstanceConfig_FromProto(mapCtx, in.GetInstanceConfig())
	out.ModelParameters = Value_FromProto(mapCtx, in.GetModelParameters())
	out.OutputConfig = BatchPredictionJob_OutputConfig_FromProto(mapCtx, in.GetOutputConfig())
	out.DedicatedResources = BatchDedicatedResources_FromProto(mapCtx, in.GetDedicatedResources())
	out.ServiceAccount = direct.LazyPtr(in.GetServiceAccount())
	out.ManualBatchTuningParameters = ManualBatchTuningParameters_FromProto(mapCtx, in.GetManualBatchTuningParameters())
	out.GenerateExplanation = direct.LazyPtr(in.GetGenerateExplanation())
	out.ExplanationSpec = ExplanationSpec_FromProto(mapCtx, in.GetExplanationSpec())
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
	out.Labels = in.Labels
	out.EncryptionSpec = EncryptionSpec_FromProto(mapCtx, in.GetEncryptionSpec())
	out.DisableContainerLogging = direct.LazyPtr(in.GetDisableContainerLogging())
	// MISSING: SatisfiesPzs
	// MISSING: SatisfiesPzi
	return out
}
func BatchPredictionJob_ToProto(mapCtx *direct.MapContext, in *krm.BatchPredictionJob) *pb.BatchPredictionJob {
	if in == nil {
		return nil
	}
	out := &pb.BatchPredictionJob{}
	// MISSING: Name
	out.DisplayName = direct.ValueOf(in.DisplayName)
	out.Model = direct.ValueOf(in.Model)
	// MISSING: ModelVersionID
	out.UnmanagedContainerModel = UnmanagedContainerModel_ToProto(mapCtx, in.UnmanagedContainerModel)
	out.InputConfig = BatchPredictionJob_InputConfig_ToProto(mapCtx, in.InputConfig)
	out.InstanceConfig = BatchPredictionJob_InstanceConfig_ToProto(mapCtx, in.InstanceConfig)
	out.ModelParameters = Value_ToProto(mapCtx, in.ModelParameters)
	out.OutputConfig = BatchPredictionJob_OutputConfig_ToProto(mapCtx, in.OutputConfig)
	out.DedicatedResources = BatchDedicatedResources_ToProto(mapCtx, in.DedicatedResources)
	out.ServiceAccount = direct.ValueOf(in.ServiceAccount)
	out.ManualBatchTuningParameters = ManualBatchTuningParameters_ToProto(mapCtx, in.ManualBatchTuningParameters)
	out.GenerateExplanation = direct.ValueOf(in.GenerateExplanation)
	out.ExplanationSpec = ExplanationSpec_ToProto(mapCtx, in.ExplanationSpec)
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
	out.Labels = in.Labels
	out.EncryptionSpec = EncryptionSpec_ToProto(mapCtx, in.EncryptionSpec)
	out.DisableContainerLogging = direct.ValueOf(in.DisableContainerLogging)
	// MISSING: SatisfiesPzs
	// MISSING: SatisfiesPzi
	return out
}
func BatchPredictionJobObservedState_FromProto(mapCtx *direct.MapContext, in *pb.BatchPredictionJob) *krm.BatchPredictionJobObservedState {
	if in == nil {
		return nil
	}
	out := &krm.BatchPredictionJobObservedState{}
	out.Name = direct.LazyPtr(in.GetName())
	// MISSING: DisplayName
	// MISSING: Model
	out.ModelVersionID = direct.LazyPtr(in.GetModelVersionId())
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
	out.OutputInfo = BatchPredictionJob_OutputInfo_FromProto(mapCtx, in.GetOutputInfo())
	out.State = direct.Enum_FromProto(mapCtx, in.GetState())
	out.Error = Status_FromProto(mapCtx, in.GetError())
	out.PartialFailures = direct.Slice_FromProto(mapCtx, in.PartialFailures, Status_FromProto)
	out.ResourcesConsumed = ResourcesConsumed_FromProto(mapCtx, in.GetResourcesConsumed())
	out.CompletionStats = CompletionStats_FromProto(mapCtx, in.GetCompletionStats())
	out.CreateTime = direct.StringTimestamp_FromProto(mapCtx, in.GetCreateTime())
	out.StartTime = direct.StringTimestamp_FromProto(mapCtx, in.GetStartTime())
	out.EndTime = direct.StringTimestamp_FromProto(mapCtx, in.GetEndTime())
	out.UpdateTime = direct.StringTimestamp_FromProto(mapCtx, in.GetUpdateTime())
	// MISSING: Labels
	// MISSING: EncryptionSpec
	// MISSING: DisableContainerLogging
	out.SatisfiesPzs = direct.LazyPtr(in.GetSatisfiesPzs())
	out.SatisfiesPzi = direct.LazyPtr(in.GetSatisfiesPzi())
	return out
}
func BatchPredictionJobObservedState_ToProto(mapCtx *direct.MapContext, in *krm.BatchPredictionJobObservedState) *pb.BatchPredictionJob {
	if in == nil {
		return nil
	}
	out := &pb.BatchPredictionJob{}
	out.Name = direct.ValueOf(in.Name)
	// MISSING: DisplayName
	// MISSING: Model
	out.ModelVersionId = direct.ValueOf(in.ModelVersionID)
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
	out.OutputInfo = BatchPredictionJob_OutputInfo_ToProto(mapCtx, in.OutputInfo)
	out.State = direct.Enum_ToProto[pb.JobState](mapCtx, in.State)
	out.Error = Status_ToProto(mapCtx, in.Error)
	out.PartialFailures = direct.Slice_ToProto(mapCtx, in.PartialFailures, Status_ToProto)
	out.ResourcesConsumed = ResourcesConsumed_ToProto(mapCtx, in.ResourcesConsumed)
	out.CompletionStats = CompletionStats_ToProto(mapCtx, in.CompletionStats)
	out.CreateTime = direct.StringTimestamp_ToProto(mapCtx, in.CreateTime)
	out.StartTime = direct.StringTimestamp_ToProto(mapCtx, in.StartTime)
	out.EndTime = direct.StringTimestamp_ToProto(mapCtx, in.EndTime)
	out.UpdateTime = direct.StringTimestamp_ToProto(mapCtx, in.UpdateTime)
	// MISSING: Labels
	// MISSING: EncryptionSpec
	// MISSING: DisableContainerLogging
	out.SatisfiesPzs = direct.ValueOf(in.SatisfiesPzs)
	out.SatisfiesPzi = direct.ValueOf(in.SatisfiesPzi)
	return out
}
func BatchPredictionJob_InputConfig_FromProto(mapCtx *direct.MapContext, in *pb.BatchPredictionJob_InputConfig) *krm.BatchPredictionJob_InputConfig {
	if in == nil {
		return nil
	}
	out := &krm.BatchPredictionJob_InputConfig{}
	out.GcsSource = GcsSource_FromProto(mapCtx, in.GetGcsSource())
	out.BigquerySource = BigQuerySource_FromProto(mapCtx, in.GetBigquerySource())
	out.InstancesFormat = direct.LazyPtr(in.GetInstancesFormat())
	return out
}
func BatchPredictionJob_InputConfig_ToProto(mapCtx *direct.MapContext, in *krm.BatchPredictionJob_InputConfig) *pb.BatchPredictionJob_InputConfig {
	if in == nil {
		return nil
	}
	out := &pb.BatchPredictionJob_InputConfig{}
	if oneof := GcsSource_ToProto(mapCtx, in.GcsSource); oneof != nil {
		out.Source = &pb.BatchPredictionJob_InputConfig_GcsSource{GcsSource: oneof}
	}
	if oneof := BigQuerySource_ToProto(mapCtx, in.BigquerySource); oneof != nil {
		out.Source = &pb.BatchPredictionJob_InputConfig_BigquerySource{BigquerySource: oneof}
	}
	out.InstancesFormat = direct.ValueOf(in.InstancesFormat)
	return out
}
func BatchPredictionJob_InstanceConfig_FromProto(mapCtx *direct.MapContext, in *pb.BatchPredictionJob_InstanceConfig) *krm.BatchPredictionJob_InstanceConfig {
	if in == nil {
		return nil
	}
	out := &krm.BatchPredictionJob_InstanceConfig{}
	out.InstanceType = direct.LazyPtr(in.GetInstanceType())
	out.KeyField = direct.LazyPtr(in.GetKeyField())
	out.IncludedFields = in.IncludedFields
	out.ExcludedFields = in.ExcludedFields
	return out
}
func BatchPredictionJob_InstanceConfig_ToProto(mapCtx *direct.MapContext, in *krm.BatchPredictionJob_InstanceConfig) *pb.BatchPredictionJob_InstanceConfig {
	if in == nil {
		return nil
	}
	out := &pb.BatchPredictionJob_InstanceConfig{}
	out.InstanceType = direct.ValueOf(in.InstanceType)
	out.KeyField = direct.ValueOf(in.KeyField)
	out.IncludedFields = in.IncludedFields
	out.ExcludedFields = in.ExcludedFields
	return out
}
func BatchPredictionJob_OutputConfig_FromProto(mapCtx *direct.MapContext, in *pb.BatchPredictionJob_OutputConfig) *krm.BatchPredictionJob_OutputConfig {
	if in == nil {
		return nil
	}
	out := &krm.BatchPredictionJob_OutputConfig{}
	out.GcsDestination = GcsDestination_FromProto(mapCtx, in.GetGcsDestination())
	out.BigqueryDestination = BigQueryDestination_FromProto(mapCtx, in.GetBigqueryDestination())
	out.PredictionsFormat = direct.LazyPtr(in.GetPredictionsFormat())
	return out
}
func BatchPredictionJob_OutputConfig_ToProto(mapCtx *direct.MapContext, in *krm.BatchPredictionJob_OutputConfig) *pb.BatchPredictionJob_OutputConfig {
	if in == nil {
		return nil
	}
	out := &pb.BatchPredictionJob_OutputConfig{}
	if oneof := GcsDestination_ToProto(mapCtx, in.GcsDestination); oneof != nil {
		out.Destination = &pb.BatchPredictionJob_OutputConfig_GcsDestination{GcsDestination: oneof}
	}
	if oneof := BigQueryDestination_ToProto(mapCtx, in.BigqueryDestination); oneof != nil {
		out.Destination = &pb.BatchPredictionJob_OutputConfig_BigqueryDestination{BigqueryDestination: oneof}
	}
	out.PredictionsFormat = direct.ValueOf(in.PredictionsFormat)
	return out
}
func BatchPredictionJob_OutputInfo_FromProto(mapCtx *direct.MapContext, in *pb.BatchPredictionJob_OutputInfo) *krm.BatchPredictionJob_OutputInfo {
	if in == nil {
		return nil
	}
	out := &krm.BatchPredictionJob_OutputInfo{}
	// MISSING: GcsOutputDirectory
	// MISSING: BigqueryOutputDataset
	// MISSING: BigqueryOutputTable
	return out
}
func BatchPredictionJob_OutputInfo_ToProto(mapCtx *direct.MapContext, in *krm.BatchPredictionJob_OutputInfo) *pb.BatchPredictionJob_OutputInfo {
	if in == nil {
		return nil
	}
	out := &pb.BatchPredictionJob_OutputInfo{}
	// MISSING: GcsOutputDirectory
	// MISSING: BigqueryOutputDataset
	// MISSING: BigqueryOutputTable
	return out
}
func BatchPredictionJob_OutputInfoObservedState_FromProto(mapCtx *direct.MapContext, in *pb.BatchPredictionJob_OutputInfo) *krm.BatchPredictionJob_OutputInfoObservedState {
	if in == nil {
		return nil
	}
	out := &krm.BatchPredictionJob_OutputInfoObservedState{}
	out.GcsOutputDirectory = direct.LazyPtr(in.GetGcsOutputDirectory())
	out.BigqueryOutputDataset = direct.LazyPtr(in.GetBigqueryOutputDataset())
	out.BigqueryOutputTable = direct.LazyPtr(in.GetBigqueryOutputTable())
	return out
}
func BatchPredictionJob_OutputInfoObservedState_ToProto(mapCtx *direct.MapContext, in *krm.BatchPredictionJob_OutputInfoObservedState) *pb.BatchPredictionJob_OutputInfo {
	if in == nil {
		return nil
	}
	out := &pb.BatchPredictionJob_OutputInfo{}
	if oneof := BatchPredictionJob_OutputInfoObservedState_GcsOutputDirectory_ToProto(mapCtx, in.GcsOutputDirectory); oneof != nil {
		out.OutputLocation = oneof
	}
	if oneof := BatchPredictionJob_OutputInfoObservedState_BigqueryOutputDataset_ToProto(mapCtx, in.BigqueryOutputDataset); oneof != nil {
		out.OutputLocation = oneof
	}
	out.BigqueryOutputTable = direct.ValueOf(in.BigqueryOutputTable)
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
func CompletionStats_FromProto(mapCtx *direct.MapContext, in *pb.CompletionStats) *krm.CompletionStats {
	if in == nil {
		return nil
	}
	out := &krm.CompletionStats{}
	// MISSING: SuccessfulCount
	// MISSING: FailedCount
	// MISSING: IncompleteCount
	// MISSING: SuccessfulForecastPointCount
	return out
}
func CompletionStats_ToProto(mapCtx *direct.MapContext, in *krm.CompletionStats) *pb.CompletionStats {
	if in == nil {
		return nil
	}
	out := &pb.CompletionStats{}
	// MISSING: SuccessfulCount
	// MISSING: FailedCount
	// MISSING: IncompleteCount
	// MISSING: SuccessfulForecastPointCount
	return out
}
func CompletionStatsObservedState_FromProto(mapCtx *direct.MapContext, in *pb.CompletionStats) *krm.CompletionStatsObservedState {
	if in == nil {
		return nil
	}
	out := &krm.CompletionStatsObservedState{}
	out.SuccessfulCount = direct.LazyPtr(in.GetSuccessfulCount())
	out.FailedCount = direct.LazyPtr(in.GetFailedCount())
	out.IncompleteCount = direct.LazyPtr(in.GetIncompleteCount())
	out.SuccessfulForecastPointCount = direct.LazyPtr(in.GetSuccessfulForecastPointCount())
	return out
}
func CompletionStatsObservedState_ToProto(mapCtx *direct.MapContext, in *krm.CompletionStatsObservedState) *pb.CompletionStats {
	if in == nil {
		return nil
	}
	out := &pb.CompletionStats{}
	out.SuccessfulCount = direct.ValueOf(in.SuccessfulCount)
	out.FailedCount = direct.ValueOf(in.FailedCount)
	out.IncompleteCount = direct.ValueOf(in.IncompleteCount)
	out.SuccessfulForecastPointCount = direct.ValueOf(in.SuccessfulForecastPointCount)
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
func ManualBatchTuningParameters_FromProto(mapCtx *direct.MapContext, in *pb.ManualBatchTuningParameters) *krm.ManualBatchTuningParameters {
	if in == nil {
		return nil
	}
	out := &krm.ManualBatchTuningParameters{}
	out.BatchSize = direct.LazyPtr(in.GetBatchSize())
	return out
}
func ManualBatchTuningParameters_ToProto(mapCtx *direct.MapContext, in *krm.ManualBatchTuningParameters) *pb.ManualBatchTuningParameters {
	if in == nil {
		return nil
	}
	out := &pb.ManualBatchTuningParameters{}
	out.BatchSize = direct.ValueOf(in.BatchSize)
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
func ResourcesConsumed_FromProto(mapCtx *direct.MapContext, in *pb.ResourcesConsumed) *krm.ResourcesConsumed {
	if in == nil {
		return nil
	}
	out := &krm.ResourcesConsumed{}
	// MISSING: ReplicaHours
	return out
}
func ResourcesConsumed_ToProto(mapCtx *direct.MapContext, in *krm.ResourcesConsumed) *pb.ResourcesConsumed {
	if in == nil {
		return nil
	}
	out := &pb.ResourcesConsumed{}
	// MISSING: ReplicaHours
	return out
}
func ResourcesConsumedObservedState_FromProto(mapCtx *direct.MapContext, in *pb.ResourcesConsumed) *krm.ResourcesConsumedObservedState {
	if in == nil {
		return nil
	}
	out := &krm.ResourcesConsumedObservedState{}
	out.ReplicaHours = direct.LazyPtr(in.GetReplicaHours())
	return out
}
func ResourcesConsumedObservedState_ToProto(mapCtx *direct.MapContext, in *krm.ResourcesConsumedObservedState) *pb.ResourcesConsumed {
	if in == nil {
		return nil
	}
	out := &pb.ResourcesConsumed{}
	out.ReplicaHours = direct.ValueOf(in.ReplicaHours)
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
func UnmanagedContainerModel_FromProto(mapCtx *direct.MapContext, in *pb.UnmanagedContainerModel) *krm.UnmanagedContainerModel {
	if in == nil {
		return nil
	}
	out := &krm.UnmanagedContainerModel{}
	out.ArtifactURI = direct.LazyPtr(in.GetArtifactUri())
	out.PredictSchemata = PredictSchemata_FromProto(mapCtx, in.GetPredictSchemata())
	out.ContainerSpec = ModelContainerSpec_FromProto(mapCtx, in.GetContainerSpec())
	return out
}
func UnmanagedContainerModel_ToProto(mapCtx *direct.MapContext, in *krm.UnmanagedContainerModel) *pb.UnmanagedContainerModel {
	if in == nil {
		return nil
	}
	out := &pb.UnmanagedContainerModel{}
	out.ArtifactUri = direct.ValueOf(in.ArtifactURI)
	out.PredictSchemata = PredictSchemata_ToProto(mapCtx, in.PredictSchemata)
	out.ContainerSpec = ModelContainerSpec_ToProto(mapCtx, in.ContainerSpec)
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
