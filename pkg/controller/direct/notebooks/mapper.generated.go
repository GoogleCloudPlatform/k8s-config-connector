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

package notebooks

import (
	refs "github.com/GoogleCloudPlatform/k8s-config-connector/apis/refs/v1beta1"
	pb "cloud.google.com/go/notebooks/apiv1/notebookspb"
	krm "github.com/GoogleCloudPlatform/k8s-config-connector/apis/notebooks/v1alpha1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct"
)
func Execution_FromProto(mapCtx *direct.MapContext, in *pb.Execution) *krm.Execution {
	if in == nil {
		return nil
	}
	out := &krm.Execution{}
	out.ExecutionTemplate = ExecutionTemplate_FromProto(mapCtx, in.GetExecutionTemplate())
	// MISSING: Name
	// MISSING: DisplayName
	out.Description = direct.LazyPtr(in.GetDescription())
	// MISSING: CreateTime
	// MISSING: UpdateTime
	// MISSING: State
	out.OutputNotebookFile = direct.LazyPtr(in.GetOutputNotebookFile())
	// MISSING: JobURI
	return out
}
func Execution_ToProto(mapCtx *direct.MapContext, in *krm.Execution) *pb.Execution {
	if in == nil {
		return nil
	}
	out := &pb.Execution{}
	out.ExecutionTemplate = ExecutionTemplate_ToProto(mapCtx, in.ExecutionTemplate)
	// MISSING: Name
	// MISSING: DisplayName
	out.Description = direct.ValueOf(in.Description)
	// MISSING: CreateTime
	// MISSING: UpdateTime
	// MISSING: State
	out.OutputNotebookFile = direct.ValueOf(in.OutputNotebookFile)
	// MISSING: JobURI
	return out
}
func ExecutionObservedState_FromProto(mapCtx *direct.MapContext, in *pb.Execution) *krm.ExecutionObservedState {
	if in == nil {
		return nil
	}
	out := &krm.ExecutionObservedState{}
	// MISSING: ExecutionTemplate
	out.Name = direct.LazyPtr(in.GetName())
	out.DisplayName = direct.LazyPtr(in.GetDisplayName())
	// MISSING: Description
	out.CreateTime = direct.StringTimestamp_FromProto(mapCtx, in.GetCreateTime())
	out.UpdateTime = direct.StringTimestamp_FromProto(mapCtx, in.GetUpdateTime())
	out.State = direct.Enum_FromProto(mapCtx, in.GetState())
	// MISSING: OutputNotebookFile
	out.JobURI = direct.LazyPtr(in.GetJobUri())
	return out
}
func ExecutionObservedState_ToProto(mapCtx *direct.MapContext, in *krm.ExecutionObservedState) *pb.Execution {
	if in == nil {
		return nil
	}
	out := &pb.Execution{}
	// MISSING: ExecutionTemplate
	out.Name = direct.ValueOf(in.Name)
	out.DisplayName = direct.ValueOf(in.DisplayName)
	// MISSING: Description
	out.CreateTime = direct.StringTimestamp_ToProto(mapCtx, in.CreateTime)
	out.UpdateTime = direct.StringTimestamp_ToProto(mapCtx, in.UpdateTime)
	out.State = direct.Enum_ToProto[pb.Execution_State](mapCtx, in.State)
	// MISSING: OutputNotebookFile
	out.JobUri = direct.ValueOf(in.JobURI)
	return out
}
func ExecutionTemplate_FromProto(mapCtx *direct.MapContext, in *pb.ExecutionTemplate) *krm.ExecutionTemplate {
	if in == nil {
		return nil
	}
	out := &krm.ExecutionTemplate{}
	out.ScaleTier = direct.Enum_FromProto(mapCtx, in.GetScaleTier())
	out.MasterType = direct.LazyPtr(in.GetMasterType())
	out.AcceleratorConfig = ExecutionTemplate_SchedulerAcceleratorConfig_FromProto(mapCtx, in.GetAcceleratorConfig())
	out.Labels = in.Labels
	out.InputNotebookFile = direct.LazyPtr(in.GetInputNotebookFile())
	out.ContainerImageURI = direct.LazyPtr(in.GetContainerImageUri())
	out.OutputNotebookFolder = direct.LazyPtr(in.GetOutputNotebookFolder())
	out.ParamsYamlFile = direct.LazyPtr(in.GetParamsYamlFile())
	out.Parameters = direct.LazyPtr(in.GetParameters())
	out.ServiceAccount = direct.LazyPtr(in.GetServiceAccount())
	out.JobType = direct.Enum_FromProto(mapCtx, in.GetJobType())
	out.DataprocParameters = ExecutionTemplate_DataprocParameters_FromProto(mapCtx, in.GetDataprocParameters())
	out.VertexAiParameters = ExecutionTemplate_VertexAIParameters_FromProto(mapCtx, in.GetVertexAiParameters())
	out.KernelSpec = direct.LazyPtr(in.GetKernelSpec())
	out.Tensorboard = direct.LazyPtr(in.GetTensorboard())
	return out
}
func ExecutionTemplate_ToProto(mapCtx *direct.MapContext, in *krm.ExecutionTemplate) *pb.ExecutionTemplate {
	if in == nil {
		return nil
	}
	out := &pb.ExecutionTemplate{}
	out.ScaleTier = direct.Enum_ToProto[pb.ExecutionTemplate_ScaleTier](mapCtx, in.ScaleTier)
	out.MasterType = direct.ValueOf(in.MasterType)
	out.AcceleratorConfig = ExecutionTemplate_SchedulerAcceleratorConfig_ToProto(mapCtx, in.AcceleratorConfig)
	out.Labels = in.Labels
	out.InputNotebookFile = direct.ValueOf(in.InputNotebookFile)
	out.ContainerImageUri = direct.ValueOf(in.ContainerImageURI)
	out.OutputNotebookFolder = direct.ValueOf(in.OutputNotebookFolder)
	out.ParamsYamlFile = direct.ValueOf(in.ParamsYamlFile)
	out.Parameters = direct.ValueOf(in.Parameters)
	out.ServiceAccount = direct.ValueOf(in.ServiceAccount)
	out.JobType = direct.Enum_ToProto[pb.ExecutionTemplate_JobType](mapCtx, in.JobType)
	if oneof := ExecutionTemplate_DataprocParameters_ToProto(mapCtx, in.DataprocParameters); oneof != nil {
		out.JobParameters = &pb.ExecutionTemplate_DataprocParameters_{DataprocParameters: oneof}
	}
	if oneof := ExecutionTemplate_VertexAIParameters_ToProto(mapCtx, in.VertexAiParameters); oneof != nil {
		out.JobParameters = &pb.ExecutionTemplate_VertexAiParameters{VertexAiParameters: oneof}
	}
	out.KernelSpec = direct.ValueOf(in.KernelSpec)
	out.Tensorboard = direct.ValueOf(in.Tensorboard)
	return out
}
func ExecutionTemplate_DataprocParameters_FromProto(mapCtx *direct.MapContext, in *pb.ExecutionTemplate_DataprocParameters) *krm.ExecutionTemplate_DataprocParameters {
	if in == nil {
		return nil
	}
	out := &krm.ExecutionTemplate_DataprocParameters{}
	out.Cluster = direct.LazyPtr(in.GetCluster())
	return out
}
func ExecutionTemplate_DataprocParameters_ToProto(mapCtx *direct.MapContext, in *krm.ExecutionTemplate_DataprocParameters) *pb.ExecutionTemplate_DataprocParameters {
	if in == nil {
		return nil
	}
	out := &pb.ExecutionTemplate_DataprocParameters{}
	out.Cluster = direct.ValueOf(in.Cluster)
	return out
}
func ExecutionTemplate_SchedulerAcceleratorConfig_FromProto(mapCtx *direct.MapContext, in *pb.ExecutionTemplate_SchedulerAcceleratorConfig) *krm.ExecutionTemplate_SchedulerAcceleratorConfig {
	if in == nil {
		return nil
	}
	out := &krm.ExecutionTemplate_SchedulerAcceleratorConfig{}
	out.Type = direct.Enum_FromProto(mapCtx, in.GetType())
	out.CoreCount = direct.LazyPtr(in.GetCoreCount())
	return out
}
func ExecutionTemplate_SchedulerAcceleratorConfig_ToProto(mapCtx *direct.MapContext, in *krm.ExecutionTemplate_SchedulerAcceleratorConfig) *pb.ExecutionTemplate_SchedulerAcceleratorConfig {
	if in == nil {
		return nil
	}
	out := &pb.ExecutionTemplate_SchedulerAcceleratorConfig{}
	out.Type = direct.Enum_ToProto[pb.ExecutionTemplate_SchedulerAcceleratorType](mapCtx, in.Type)
	out.CoreCount = direct.ValueOf(in.CoreCount)
	return out
}
func ExecutionTemplate_VertexAIParameters_FromProto(mapCtx *direct.MapContext, in *pb.ExecutionTemplate_VertexAIParameters) *krm.ExecutionTemplate_VertexAIParameters {
	if in == nil {
		return nil
	}
	out := &krm.ExecutionTemplate_VertexAIParameters{}
	out.Network = direct.LazyPtr(in.GetNetwork())
	out.Env = in.Env
	return out
}
func ExecutionTemplate_VertexAIParameters_ToProto(mapCtx *direct.MapContext, in *krm.ExecutionTemplate_VertexAIParameters) *pb.ExecutionTemplate_VertexAIParameters {
	if in == nil {
		return nil
	}
	out := &pb.ExecutionTemplate_VertexAIParameters{}
	out.Network = direct.ValueOf(in.Network)
	out.Env = in.Env
	return out
}
func NotebooksExecutionObservedState_FromProto(mapCtx *direct.MapContext, in *pb.Execution) *krm.NotebooksExecutionObservedState {
	if in == nil {
		return nil
	}
	out := &krm.NotebooksExecutionObservedState{}
	// MISSING: ExecutionTemplate
	// MISSING: Name
	// MISSING: DisplayName
	// MISSING: Description
	// MISSING: CreateTime
	// MISSING: UpdateTime
	// MISSING: State
	// MISSING: OutputNotebookFile
	// MISSING: JobURI
	return out
}
func NotebooksExecutionObservedState_ToProto(mapCtx *direct.MapContext, in *krm.NotebooksExecutionObservedState) *pb.Execution {
	if in == nil {
		return nil
	}
	out := &pb.Execution{}
	// MISSING: ExecutionTemplate
	// MISSING: Name
	// MISSING: DisplayName
	// MISSING: Description
	// MISSING: CreateTime
	// MISSING: UpdateTime
	// MISSING: State
	// MISSING: OutputNotebookFile
	// MISSING: JobURI
	return out
}
func NotebooksExecutionSpec_FromProto(mapCtx *direct.MapContext, in *pb.Execution) *krm.NotebooksExecutionSpec {
	if in == nil {
		return nil
	}
	out := &krm.NotebooksExecutionSpec{}
	// MISSING: ExecutionTemplate
	// MISSING: Name
	// MISSING: DisplayName
	// MISSING: Description
	// MISSING: CreateTime
	// MISSING: UpdateTime
	// MISSING: State
	// MISSING: OutputNotebookFile
	// MISSING: JobURI
	return out
}
func NotebooksExecutionSpec_ToProto(mapCtx *direct.MapContext, in *krm.NotebooksExecutionSpec) *pb.Execution {
	if in == nil {
		return nil
	}
	out := &pb.Execution{}
	// MISSING: ExecutionTemplate
	// MISSING: Name
	// MISSING: DisplayName
	// MISSING: Description
	// MISSING: CreateTime
	// MISSING: UpdateTime
	// MISSING: State
	// MISSING: OutputNotebookFile
	// MISSING: JobURI
	return out
}
