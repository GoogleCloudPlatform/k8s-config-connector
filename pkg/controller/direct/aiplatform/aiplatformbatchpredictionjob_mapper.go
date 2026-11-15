// Copyright 2026 Google LLC
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
)

func AIPlatformBatchPredictionJobObservedState_FromProto(mapCtx *direct.MapContext, in *pb.BatchPredictionJob) *krm.AIPlatformBatchPredictionJobObservedState {
	if in == nil {
		return nil
	}
	out := &krm.AIPlatformBatchPredictionJobObservedState{}
	// MISSING: Name
	out.ModelVersionID = direct.LazyPtr(in.GetModelVersionId())
	out.OutputInfo = BatchPredictionJob_OutputInfoObservedState_FromProto(mapCtx, in.GetOutputInfo())
	out.State = direct.Enum_FromProto(mapCtx, in.GetState())
	out.Error = direct.Status_FromProto(mapCtx, in.GetError())
	out.PartialFailures = direct.Slice_FromProto(mapCtx, in.PartialFailures, Status_FromProto)
	out.ResourcesConsumed = ResourcesConsumedObservedState_FromProto(mapCtx, in.GetResourcesConsumed())
	out.CompletionStats = CompletionStatsObservedState_FromProto(mapCtx, in.GetCompletionStats())
	out.CreateTime = direct.StringTimestamp_FromProto(mapCtx, in.GetCreateTime())
	out.StartTime = direct.StringTimestamp_FromProto(mapCtx, in.GetStartTime())
	out.EndTime = direct.StringTimestamp_FromProto(mapCtx, in.GetEndTime())
	out.UpdateTime = direct.StringTimestamp_FromProto(mapCtx, in.GetUpdateTime())
	out.SatisfiesPzs = direct.LazyPtr(in.GetSatisfiesPzs())
	out.SatisfiesPzi = direct.LazyPtr(in.GetSatisfiesPzi())
	return out
}

func AIPlatformBatchPredictionJobObservedState_ToProto(mapCtx *direct.MapContext, in *krm.AIPlatformBatchPredictionJobObservedState) *pb.BatchPredictionJob {
	if in == nil {
		return nil
	}
	out := &pb.BatchPredictionJob{}
	// MISSING: Name
	out.ModelVersionId = direct.ValueOf(in.ModelVersionID)
	out.OutputInfo = BatchPredictionJob_OutputInfoObservedState_ToProto(mapCtx, in.OutputInfo)
	out.State = direct.Enum_ToProto[pb.JobState](mapCtx, in.State)
	out.Error = direct.Status_ToProto(mapCtx, in.Error)
	out.PartialFailures = direct.Slice_ToProto(mapCtx, in.PartialFailures, Status_ToProto)
	out.ResourcesConsumed = ResourcesConsumedObservedState_ToProto(mapCtx, in.ResourcesConsumed)
	out.CompletionStats = CompletionStatsObservedState_ToProto(mapCtx, in.CompletionStats)
	out.CreateTime = direct.StringTimestamp_ToProto(mapCtx, in.CreateTime)
	out.StartTime = direct.StringTimestamp_ToProto(mapCtx, in.StartTime)
	out.EndTime = direct.StringTimestamp_ToProto(mapCtx, in.EndTime)
	out.UpdateTime = direct.StringTimestamp_ToProto(mapCtx, in.UpdateTime)
	out.SatisfiesPzs = direct.ValueOf(in.SatisfiesPzs)
	out.SatisfiesPzi = direct.ValueOf(in.SatisfiesPzi)
	return out
}

func AIPlatformBatchPredictionJobSpec_FromProto(mapCtx *direct.MapContext, in *pb.BatchPredictionJob) *krm.AIPlatformBatchPredictionJobSpec {
	if in == nil {
		return nil
	}
	out := &krm.AIPlatformBatchPredictionJobSpec{}
	// MISSING: Name
	out.DisplayName = direct.LazyPtr(in.GetDisplayName())
	out.Model = direct.LazyPtr(in.GetModel())
	out.UnmanagedContainerModel = UnmanagedContainerModel_FromProto(mapCtx, in.GetUnmanagedContainerModel())
	out.InputConfig = BatchPredictionJob_InputConfig_FromProto(mapCtx, in.GetInputConfig())
	out.InstanceConfig = BatchPredictionJob_InstanceConfig_FromProto(mapCtx, in.GetInstanceConfig())
	out.ModelParameters = Value_FromProto(mapCtx, in.GetModelParameters())
	out.OutputConfig = BatchPredictionJob_OutputConfig_FromProto(mapCtx, in.GetOutputConfig())
	out.DedicatedResources = batchDedicatedResources_v1_FromProto(mapCtx, in.GetDedicatedResources())
	out.ServiceAccount = direct.LazyPtr(in.GetServiceAccount())
	out.ManualBatchTuningParameters = ManualBatchTuningParameters_FromProto(mapCtx, in.GetManualBatchTuningParameters())
	out.GenerateExplanation = direct.LazyPtr(in.GetGenerateExplanation())
	out.ExplanationSpec = explanationSpecWithRequired_v1_FromProto(mapCtx, in.GetExplanationSpec())
	out.Labels = in.Labels
	out.EncryptionSpec = EncryptionSpec_FromProto(mapCtx, in.GetEncryptionSpec())
	out.DisableContainerLogging = direct.LazyPtr(in.GetDisableContainerLogging())
	return out
}

func AIPlatformBatchPredictionJobSpec_ToProto(mapCtx *direct.MapContext, in *krm.AIPlatformBatchPredictionJobSpec) *pb.BatchPredictionJob {
	if in == nil {
		return nil
	}
	out := &pb.BatchPredictionJob{}
	// MISSING: Name
	out.DisplayName = direct.ValueOf(in.DisplayName)
	out.Model = direct.ValueOf(in.Model)
	out.UnmanagedContainerModel = UnmanagedContainerModel_ToProto(mapCtx, in.UnmanagedContainerModel)
	out.InputConfig = BatchPredictionJob_InputConfig_ToProto(mapCtx, in.InputConfig)
	out.InstanceConfig = BatchPredictionJob_InstanceConfig_ToProto(mapCtx, in.InstanceConfig)
	out.ModelParameters = Value_ToProto(mapCtx, in.ModelParameters)
	out.OutputConfig = BatchPredictionJob_OutputConfig_ToProto(mapCtx, in.OutputConfig)
	out.DedicatedResources = batchDedicatedResources_v1_ToProto(mapCtx, in.DedicatedResources)
	out.ServiceAccount = direct.ValueOf(in.ServiceAccount)
	out.ManualBatchTuningParameters = ManualBatchTuningParameters_ToProto(mapCtx, in.ManualBatchTuningParameters)
	out.GenerateExplanation = direct.ValueOf(in.GenerateExplanation)
	out.ExplanationSpec = explanationSpecWithRequired_v1_ToProto(mapCtx, in.ExplanationSpec)
	out.Labels = in.Labels
	out.EncryptionSpec = EncryptionSpec_ToProto(mapCtx, in.EncryptionSpec)
	out.DisableContainerLogging = direct.ValueOf(in.DisableContainerLogging)
	return out
}

func batchDedicatedResources_v1_FromProto(mapCtx *direct.MapContext, in *pb.BatchDedicatedResources) *krm.BatchDedicatedResources {
	if in == nil {
		return nil
	}
	out := &krm.BatchDedicatedResources{}
	out.MachineSpec = MachineSpec_FromProto(mapCtx, in.GetMachineSpec())
	out.StartingReplicaCount = direct.LazyPtr(in.GetStartingReplicaCount())
	out.MaxReplicaCount = direct.LazyPtr(in.GetMaxReplicaCount())
	return out
}

func batchDedicatedResources_v1_ToProto(mapCtx *direct.MapContext, in *krm.BatchDedicatedResources) *pb.BatchDedicatedResources {
	if in == nil {
		return nil
	}
	out := &pb.BatchDedicatedResources{}
	out.MachineSpec = MachineSpec_ToProto(mapCtx, in.MachineSpec)
	out.StartingReplicaCount = direct.ValueOf(in.StartingReplicaCount)
	out.MaxReplicaCount = direct.ValueOf(in.MaxReplicaCount)
	return out
}

func explanationSpecWithRequired_v1_FromProto(mapCtx *direct.MapContext, in *pb.ExplanationSpec) *krm.ExplanationSpecWithRequired {
	if in == nil {
		return nil
	}
	out := &krm.ExplanationSpecWithRequired{}
	out.Parameters = ExplanationParameters_FromProto(mapCtx, in.GetParameters())
	out.Metadata = ExplanationMetadata_FromProto(mapCtx, in.GetMetadata())
	return out
}

func explanationSpecWithRequired_v1_ToProto(mapCtx *direct.MapContext, in *krm.ExplanationSpecWithRequired) *pb.ExplanationSpec {
	if in == nil {
		return nil
	}
	out := &pb.ExplanationSpec{}
	out.Parameters = ExplanationParameters_ToProto(mapCtx, in.Parameters)
	out.Metadata = ExplanationMetadata_ToProto(mapCtx, in.Metadata)
	return out
}
