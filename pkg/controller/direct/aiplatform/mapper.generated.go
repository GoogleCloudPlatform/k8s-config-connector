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
func CustomJob_FromProto(mapCtx *direct.MapContext, in *pb.CustomJob) *krm.CustomJob {
	if in == nil {
		return nil
	}
	out := &krm.CustomJob{}
	// MISSING: Name
	out.DisplayName = direct.LazyPtr(in.GetDisplayName())
	out.JobSpec = CustomJobSpec_FromProto(mapCtx, in.GetJobSpec())
	// MISSING: State
	// MISSING: CreateTime
	// MISSING: StartTime
	// MISSING: EndTime
	// MISSING: UpdateTime
	// MISSING: Error
	out.Labels = in.Labels
	out.EncryptionSpec = EncryptionSpec_FromProto(mapCtx, in.GetEncryptionSpec())
	// MISSING: WebAccessUris
	// MISSING: SatisfiesPzs
	// MISSING: SatisfiesPzi
	return out
}
func CustomJob_ToProto(mapCtx *direct.MapContext, in *krm.CustomJob) *pb.CustomJob {
	if in == nil {
		return nil
	}
	out := &pb.CustomJob{}
	// MISSING: Name
	out.DisplayName = direct.ValueOf(in.DisplayName)
	out.JobSpec = CustomJobSpec_ToProto(mapCtx, in.JobSpec)
	// MISSING: State
	// MISSING: CreateTime
	// MISSING: StartTime
	// MISSING: EndTime
	// MISSING: UpdateTime
	// MISSING: Error
	out.Labels = in.Labels
	out.EncryptionSpec = EncryptionSpec_ToProto(mapCtx, in.EncryptionSpec)
	// MISSING: WebAccessUris
	// MISSING: SatisfiesPzs
	// MISSING: SatisfiesPzi
	return out
}
func CustomJobObservedState_FromProto(mapCtx *direct.MapContext, in *pb.CustomJob) *krm.CustomJobObservedState {
	if in == nil {
		return nil
	}
	out := &krm.CustomJobObservedState{}
	out.Name = direct.LazyPtr(in.GetName())
	// MISSING: DisplayName
	// MISSING: JobSpec
	out.State = direct.Enum_FromProto(mapCtx, in.GetState())
	out.CreateTime = direct.StringTimestamp_FromProto(mapCtx, in.GetCreateTime())
	out.StartTime = direct.StringTimestamp_FromProto(mapCtx, in.GetStartTime())
	out.EndTime = direct.StringTimestamp_FromProto(mapCtx, in.GetEndTime())
	out.UpdateTime = direct.StringTimestamp_FromProto(mapCtx, in.GetUpdateTime())
	out.Error = Status_FromProto(mapCtx, in.GetError())
	// MISSING: Labels
	// MISSING: EncryptionSpec
	out.WebAccessUris = in.WebAccessUris
	out.SatisfiesPzs = direct.LazyPtr(in.GetSatisfiesPzs())
	out.SatisfiesPzi = direct.LazyPtr(in.GetSatisfiesPzi())
	return out
}
func CustomJobObservedState_ToProto(mapCtx *direct.MapContext, in *krm.CustomJobObservedState) *pb.CustomJob {
	if in == nil {
		return nil
	}
	out := &pb.CustomJob{}
	out.Name = direct.ValueOf(in.Name)
	// MISSING: DisplayName
	// MISSING: JobSpec
	out.State = direct.Enum_ToProto[pb.JobState](mapCtx, in.State)
	out.CreateTime = direct.StringTimestamp_ToProto(mapCtx, in.CreateTime)
	out.StartTime = direct.StringTimestamp_ToProto(mapCtx, in.StartTime)
	out.EndTime = direct.StringTimestamp_ToProto(mapCtx, in.EndTime)
	out.UpdateTime = direct.StringTimestamp_ToProto(mapCtx, in.UpdateTime)
	out.Error = Status_ToProto(mapCtx, in.Error)
	// MISSING: Labels
	// MISSING: EncryptionSpec
	out.WebAccessUris = in.WebAccessUris
	out.SatisfiesPzs = direct.ValueOf(in.SatisfiesPzs)
	out.SatisfiesPzi = direct.ValueOf(in.SatisfiesPzi)
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
