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

package dataplex

import (
	pb "cloud.google.com/go/dataplex/apiv1/dataplexpb"
	krm "github.com/GoogleCloudPlatform/k8s-config-connector/apis/dataplex/v1alpha1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct"
)

func DataplexTaskObservedState_FromProto(mapCtx *direct.MapContext, in *pb.Task) *krm.DataplexTaskObservedState {
	if in == nil {
		return nil
	}
	out := &krm.DataplexTaskObservedState{}
	// MISSING: Name
	// MISSING: Uid
	// MISSING: CreateTime
	// MISSING: UpdateTime
	// MISSING: Description
	// MISSING: DisplayName
	// MISSING: State
	// MISSING: Labels
	// MISSING: TriggerSpec
	// MISSING: ExecutionSpec
	// MISSING: ExecutionStatus
	// MISSING: Spark
	// MISSING: Notebook
	return out
}
func DataplexTaskObservedState_ToProto(mapCtx *direct.MapContext, in *krm.DataplexTaskObservedState) *pb.Task {
	if in == nil {
		return nil
	}
	out := &pb.Task{}
	// MISSING: Name
	// MISSING: Uid
	// MISSING: CreateTime
	// MISSING: UpdateTime
	// MISSING: Description
	// MISSING: DisplayName
	// MISSING: State
	// MISSING: Labels
	// MISSING: TriggerSpec
	// MISSING: ExecutionSpec
	// MISSING: ExecutionStatus
	// MISSING: Spark
	// MISSING: Notebook
	return out
}
func DataplexTaskSpec_FromProto(mapCtx *direct.MapContext, in *pb.Task) *krm.DataplexTaskSpec {
	if in == nil {
		return nil
	}
	out := &krm.DataplexTaskSpec{}
	// MISSING: Name
	// MISSING: Uid
	// MISSING: CreateTime
	// MISSING: UpdateTime
	// MISSING: Description
	// MISSING: DisplayName
	// MISSING: State
	// MISSING: Labels
	// MISSING: TriggerSpec
	// MISSING: ExecutionSpec
	// MISSING: ExecutionStatus
	// MISSING: Spark
	// MISSING: Notebook
	return out
}
func DataplexTaskSpec_ToProto(mapCtx *direct.MapContext, in *krm.DataplexTaskSpec) *pb.Task {
	if in == nil {
		return nil
	}
	out := &pb.Task{}
	// MISSING: Name
	// MISSING: Uid
	// MISSING: CreateTime
	// MISSING: UpdateTime
	// MISSING: Description
	// MISSING: DisplayName
	// MISSING: State
	// MISSING: Labels
	// MISSING: TriggerSpec
	// MISSING: ExecutionSpec
	// MISSING: ExecutionStatus
	// MISSING: Spark
	// MISSING: Notebook
	return out
}
func Job_FromProto(mapCtx *direct.MapContext, in *pb.Job) *krm.Job {
	if in == nil {
		return nil
	}
	out := &krm.Job{}
	// MISSING: Name
	// MISSING: Uid
	// MISSING: StartTime
	// MISSING: EndTime
	// MISSING: State
	// MISSING: RetryCount
	// MISSING: Service
	// MISSING: ServiceJob
	// MISSING: Message
	// MISSING: Labels
	// MISSING: Trigger
	// MISSING: ExecutionSpec
	return out
}
func Job_ToProto(mapCtx *direct.MapContext, in *krm.Job) *pb.Job {
	if in == nil {
		return nil
	}
	out := &pb.Job{}
	// MISSING: Name
	// MISSING: Uid
	// MISSING: StartTime
	// MISSING: EndTime
	// MISSING: State
	// MISSING: RetryCount
	// MISSING: Service
	// MISSING: ServiceJob
	// MISSING: Message
	// MISSING: Labels
	// MISSING: Trigger
	// MISSING: ExecutionSpec
	return out
}
func JobObservedState_FromProto(mapCtx *direct.MapContext, in *pb.Job) *krm.JobObservedState {
	if in == nil {
		return nil
	}
	out := &krm.JobObservedState{}
	out.Name = direct.LazyPtr(in.GetName())
	out.Uid = direct.LazyPtr(in.GetUid())
	out.StartTime = direct.StringTimestamp_FromProto(mapCtx, in.GetStartTime())
	out.EndTime = direct.StringTimestamp_FromProto(mapCtx, in.GetEndTime())
	out.State = direct.Enum_FromProto(mapCtx, in.GetState())
	out.RetryCount = direct.LazyPtr(in.GetRetryCount())
	out.Service = direct.Enum_FromProto(mapCtx, in.GetService())
	out.ServiceJob = direct.LazyPtr(in.GetServiceJob())
	out.Message = direct.LazyPtr(in.GetMessage())
	out.Labels = in.Labels
	out.Trigger = direct.Enum_FromProto(mapCtx, in.GetTrigger())
	out.ExecutionSpec = Task_ExecutionSpec_FromProto(mapCtx, in.GetExecutionSpec())
	return out
}
func JobObservedState_ToProto(mapCtx *direct.MapContext, in *krm.JobObservedState) *pb.Job {
	if in == nil {
		return nil
	}
	out := &pb.Job{}
	out.Name = direct.ValueOf(in.Name)
	out.Uid = direct.ValueOf(in.Uid)
	out.StartTime = direct.StringTimestamp_ToProto(mapCtx, in.StartTime)
	out.EndTime = direct.StringTimestamp_ToProto(mapCtx, in.EndTime)
	out.State = direct.Enum_ToProto[pb.Job_State](mapCtx, in.State)
	out.RetryCount = direct.ValueOf(in.RetryCount)
	out.Service = direct.Enum_ToProto[pb.Job_Service](mapCtx, in.Service)
	out.ServiceJob = direct.ValueOf(in.ServiceJob)
	out.Message = direct.ValueOf(in.Message)
	out.Labels = in.Labels
	out.Trigger = direct.Enum_ToProto[pb.Job_Trigger](mapCtx, in.Trigger)
	out.ExecutionSpec = Task_ExecutionSpec_ToProto(mapCtx, in.ExecutionSpec)
	return out
}
func Task_FromProto(mapCtx *direct.MapContext, in *pb.Task) *krm.Task {
	if in == nil {
		return nil
	}
	out := &krm.Task{}
	// MISSING: Name
	// MISSING: Uid
	// MISSING: CreateTime
	// MISSING: UpdateTime
	out.Description = direct.LazyPtr(in.GetDescription())
	out.DisplayName = direct.LazyPtr(in.GetDisplayName())
	// MISSING: State
	out.Labels = in.Labels
	out.TriggerSpec = Task_TriggerSpec_FromProto(mapCtx, in.GetTriggerSpec())
	out.ExecutionSpec = Task_ExecutionSpec_FromProto(mapCtx, in.GetExecutionSpec())
	// MISSING: ExecutionStatus
	out.Spark = Task_SparkTaskConfig_FromProto(mapCtx, in.GetSpark())
	out.Notebook = Task_NotebookTaskConfig_FromProto(mapCtx, in.GetNotebook())
	return out
}
func Task_ToProto(mapCtx *direct.MapContext, in *krm.Task) *pb.Task {
	if in == nil {
		return nil
	}
	out := &pb.Task{}
	// MISSING: Name
	// MISSING: Uid
	// MISSING: CreateTime
	// MISSING: UpdateTime
	out.Description = direct.ValueOf(in.Description)
	out.DisplayName = direct.ValueOf(in.DisplayName)
	// MISSING: State
	out.Labels = in.Labels
	out.TriggerSpec = Task_TriggerSpec_ToProto(mapCtx, in.TriggerSpec)
	out.ExecutionSpec = Task_ExecutionSpec_ToProto(mapCtx, in.ExecutionSpec)
	// MISSING: ExecutionStatus
	if oneof := Task_SparkTaskConfig_ToProto(mapCtx, in.Spark); oneof != nil {
		out.Config = &pb.Task_Spark{Spark: oneof}
	}
	if oneof := Task_NotebookTaskConfig_ToProto(mapCtx, in.Notebook); oneof != nil {
		out.Config = &pb.Task_Notebook{Notebook: oneof}
	}
	return out
}
func TaskObservedState_FromProto(mapCtx *direct.MapContext, in *pb.Task) *krm.TaskObservedState {
	if in == nil {
		return nil
	}
	out := &krm.TaskObservedState{}
	out.Name = direct.LazyPtr(in.GetName())
	out.Uid = direct.LazyPtr(in.GetUid())
	out.CreateTime = direct.StringTimestamp_FromProto(mapCtx, in.GetCreateTime())
	out.UpdateTime = direct.StringTimestamp_FromProto(mapCtx, in.GetUpdateTime())
	// MISSING: Description
	// MISSING: DisplayName
	out.State = direct.Enum_FromProto(mapCtx, in.GetState())
	// MISSING: Labels
	// MISSING: TriggerSpec
	// MISSING: ExecutionSpec
	out.ExecutionStatus = Task_ExecutionStatus_FromProto(mapCtx, in.GetExecutionStatus())
	// MISSING: Spark
	// MISSING: Notebook
	return out
}
func TaskObservedState_ToProto(mapCtx *direct.MapContext, in *krm.TaskObservedState) *pb.Task {
	if in == nil {
		return nil
	}
	out := &pb.Task{}
	out.Name = direct.ValueOf(in.Name)
	out.Uid = direct.ValueOf(in.Uid)
	out.CreateTime = direct.StringTimestamp_ToProto(mapCtx, in.CreateTime)
	out.UpdateTime = direct.StringTimestamp_ToProto(mapCtx, in.UpdateTime)
	// MISSING: Description
	// MISSING: DisplayName
	out.State = direct.Enum_ToProto[pb.State](mapCtx, in.State)
	// MISSING: Labels
	// MISSING: TriggerSpec
	// MISSING: ExecutionSpec
	out.ExecutionStatus = Task_ExecutionStatus_ToProto(mapCtx, in.ExecutionStatus)
	// MISSING: Spark
	// MISSING: Notebook
	return out
}
func Task_ExecutionSpec_FromProto(mapCtx *direct.MapContext, in *pb.Task_ExecutionSpec) *krm.Task_ExecutionSpec {
	if in == nil {
		return nil
	}
	out := &krm.Task_ExecutionSpec{}
	out.Args = in.Args
	out.ServiceAccount = direct.LazyPtr(in.GetServiceAccount())
	out.Project = direct.LazyPtr(in.GetProject())
	out.MaxJobExecutionLifetime = direct.StringDuration_FromProto(mapCtx, in.GetMaxJobExecutionLifetime())
	out.KMSKey = direct.LazyPtr(in.GetKmsKey())
	return out
}
func Task_ExecutionSpec_ToProto(mapCtx *direct.MapContext, in *krm.Task_ExecutionSpec) *pb.Task_ExecutionSpec {
	if in == nil {
		return nil
	}
	out := &pb.Task_ExecutionSpec{}
	out.Args = in.Args
	out.ServiceAccount = direct.ValueOf(in.ServiceAccount)
	out.Project = direct.ValueOf(in.Project)
	out.MaxJobExecutionLifetime = direct.StringDuration_ToProto(mapCtx, in.MaxJobExecutionLifetime)
	out.KmsKey = direct.ValueOf(in.KMSKey)
	return out
}
func Task_ExecutionStatus_FromProto(mapCtx *direct.MapContext, in *pb.Task_ExecutionStatus) *krm.Task_ExecutionStatus {
	if in == nil {
		return nil
	}
	out := &krm.Task_ExecutionStatus{}
	// MISSING: UpdateTime
	// MISSING: LatestJob
	return out
}
func Task_ExecutionStatus_ToProto(mapCtx *direct.MapContext, in *krm.Task_ExecutionStatus) *pb.Task_ExecutionStatus {
	if in == nil {
		return nil
	}
	out := &pb.Task_ExecutionStatus{}
	// MISSING: UpdateTime
	// MISSING: LatestJob
	return out
}
func Task_ExecutionStatusObservedState_FromProto(mapCtx *direct.MapContext, in *pb.Task_ExecutionStatus) *krm.Task_ExecutionStatusObservedState {
	if in == nil {
		return nil
	}
	out := &krm.Task_ExecutionStatusObservedState{}
	out.UpdateTime = direct.StringTimestamp_FromProto(mapCtx, in.GetUpdateTime())
	out.LatestJob = Job_FromProto(mapCtx, in.GetLatestJob())
	return out
}
func Task_ExecutionStatusObservedState_ToProto(mapCtx *direct.MapContext, in *krm.Task_ExecutionStatusObservedState) *pb.Task_ExecutionStatus {
	if in == nil {
		return nil
	}
	out := &pb.Task_ExecutionStatus{}
	out.UpdateTime = direct.StringTimestamp_ToProto(mapCtx, in.UpdateTime)
	out.LatestJob = Job_ToProto(mapCtx, in.LatestJob)
	return out
}
func Task_InfrastructureSpec_FromProto(mapCtx *direct.MapContext, in *pb.Task_InfrastructureSpec) *krm.Task_InfrastructureSpec {
	if in == nil {
		return nil
	}
	out := &krm.Task_InfrastructureSpec{}
	out.Batch = Task_InfrastructureSpec_BatchComputeResources_FromProto(mapCtx, in.GetBatch())
	out.ContainerImage = Task_InfrastructureSpec_ContainerImageRuntime_FromProto(mapCtx, in.GetContainerImage())
	out.VpcNetwork = Task_InfrastructureSpec_VpcNetwork_FromProto(mapCtx, in.GetVpcNetwork())
	return out
}
func Task_InfrastructureSpec_ToProto(mapCtx *direct.MapContext, in *krm.Task_InfrastructureSpec) *pb.Task_InfrastructureSpec {
	if in == nil {
		return nil
	}
	out := &pb.Task_InfrastructureSpec{}
	if oneof := Task_InfrastructureSpec_BatchComputeResources_ToProto(mapCtx, in.Batch); oneof != nil {
		out.Resources = &pb.Task_InfrastructureSpec_Batch{Batch: oneof}
	}
	if oneof := Task_InfrastructureSpec_ContainerImageRuntime_ToProto(mapCtx, in.ContainerImage); oneof != nil {
		out.Runtime = &pb.Task_InfrastructureSpec_ContainerImage{ContainerImage: oneof}
	}
	if oneof := Task_InfrastructureSpec_VpcNetwork_ToProto(mapCtx, in.VpcNetwork); oneof != nil {
		out.Network = &pb.Task_InfrastructureSpec_VpcNetwork_{VpcNetwork: oneof}
	}
	return out
}
func Task_InfrastructureSpec_BatchComputeResources_FromProto(mapCtx *direct.MapContext, in *pb.Task_InfrastructureSpec_BatchComputeResources) *krm.Task_InfrastructureSpec_BatchComputeResources {
	if in == nil {
		return nil
	}
	out := &krm.Task_InfrastructureSpec_BatchComputeResources{}
	out.ExecutorsCount = direct.LazyPtr(in.GetExecutorsCount())
	out.MaxExecutorsCount = direct.LazyPtr(in.GetMaxExecutorsCount())
	return out
}
func Task_InfrastructureSpec_BatchComputeResources_ToProto(mapCtx *direct.MapContext, in *krm.Task_InfrastructureSpec_BatchComputeResources) *pb.Task_InfrastructureSpec_BatchComputeResources {
	if in == nil {
		return nil
	}
	out := &pb.Task_InfrastructureSpec_BatchComputeResources{}
	out.ExecutorsCount = direct.ValueOf(in.ExecutorsCount)
	out.MaxExecutorsCount = direct.ValueOf(in.MaxExecutorsCount)
	return out
}
func Task_InfrastructureSpec_ContainerImageRuntime_FromProto(mapCtx *direct.MapContext, in *pb.Task_InfrastructureSpec_ContainerImageRuntime) *krm.Task_InfrastructureSpec_ContainerImageRuntime {
	if in == nil {
		return nil
	}
	out := &krm.Task_InfrastructureSpec_ContainerImageRuntime{}
	out.Image = direct.LazyPtr(in.GetImage())
	out.JavaJars = in.JavaJars
	out.PythonPackages = in.PythonPackages
	out.Properties = in.Properties
	return out
}
func Task_InfrastructureSpec_ContainerImageRuntime_ToProto(mapCtx *direct.MapContext, in *krm.Task_InfrastructureSpec_ContainerImageRuntime) *pb.Task_InfrastructureSpec_ContainerImageRuntime {
	if in == nil {
		return nil
	}
	out := &pb.Task_InfrastructureSpec_ContainerImageRuntime{}
	out.Image = direct.ValueOf(in.Image)
	out.JavaJars = in.JavaJars
	out.PythonPackages = in.PythonPackages
	out.Properties = in.Properties
	return out
}
func Task_InfrastructureSpec_VpcNetwork_FromProto(mapCtx *direct.MapContext, in *pb.Task_InfrastructureSpec_VpcNetwork) *krm.Task_InfrastructureSpec_VpcNetwork {
	if in == nil {
		return nil
	}
	out := &krm.Task_InfrastructureSpec_VpcNetwork{}
	out.Network = direct.LazyPtr(in.GetNetwork())
	out.SubNetwork = direct.LazyPtr(in.GetSubNetwork())
	out.NetworkTags = in.NetworkTags
	return out
}
func Task_InfrastructureSpec_VpcNetwork_ToProto(mapCtx *direct.MapContext, in *krm.Task_InfrastructureSpec_VpcNetwork) *pb.Task_InfrastructureSpec_VpcNetwork {
	if in == nil {
		return nil
	}
	out := &pb.Task_InfrastructureSpec_VpcNetwork{}
	if oneof := Task_InfrastructureSpec_VpcNetwork_Network_ToProto(mapCtx, in.Network); oneof != nil {
		out.NetworkName = oneof
	}
	if oneof := Task_InfrastructureSpec_VpcNetwork_SubNetwork_ToProto(mapCtx, in.SubNetwork); oneof != nil {
		out.NetworkName = oneof
	}
	out.NetworkTags = in.NetworkTags
	return out
}
func Task_NotebookTaskConfig_FromProto(mapCtx *direct.MapContext, in *pb.Task_NotebookTaskConfig) *krm.Task_NotebookTaskConfig {
	if in == nil {
		return nil
	}
	out := &krm.Task_NotebookTaskConfig{}
	out.Notebook = direct.LazyPtr(in.GetNotebook())
	out.InfrastructureSpec = Task_InfrastructureSpec_FromProto(mapCtx, in.GetInfrastructureSpec())
	out.FileUris = in.FileUris
	out.ArchiveUris = in.ArchiveUris
	return out
}
func Task_NotebookTaskConfig_ToProto(mapCtx *direct.MapContext, in *krm.Task_NotebookTaskConfig) *pb.Task_NotebookTaskConfig {
	if in == nil {
		return nil
	}
	out := &pb.Task_NotebookTaskConfig{}
	out.Notebook = direct.ValueOf(in.Notebook)
	out.InfrastructureSpec = Task_InfrastructureSpec_ToProto(mapCtx, in.InfrastructureSpec)
	out.FileUris = in.FileUris
	out.ArchiveUris = in.ArchiveUris
	return out
}
func Task_SparkTaskConfig_FromProto(mapCtx *direct.MapContext, in *pb.Task_SparkTaskConfig) *krm.Task_SparkTaskConfig {
	if in == nil {
		return nil
	}
	out := &krm.Task_SparkTaskConfig{}
	out.MainJarFileURI = direct.LazyPtr(in.GetMainJarFileUri())
	out.MainClass = direct.LazyPtr(in.GetMainClass())
	out.PythonScriptFile = direct.LazyPtr(in.GetPythonScriptFile())
	out.SQLScriptFile = direct.LazyPtr(in.GetSqlScriptFile())
	out.SQLScript = direct.LazyPtr(in.GetSqlScript())
	out.FileUris = in.FileUris
	out.ArchiveUris = in.ArchiveUris
	out.InfrastructureSpec = Task_InfrastructureSpec_FromProto(mapCtx, in.GetInfrastructureSpec())
	return out
}
func Task_SparkTaskConfig_ToProto(mapCtx *direct.MapContext, in *krm.Task_SparkTaskConfig) *pb.Task_SparkTaskConfig {
	if in == nil {
		return nil
	}
	out := &pb.Task_SparkTaskConfig{}
	if oneof := Task_SparkTaskConfig_MainJarFileUri_ToProto(mapCtx, in.MainJarFileURI); oneof != nil {
		out.Driver = oneof
	}
	if oneof := Task_SparkTaskConfig_MainClass_ToProto(mapCtx, in.MainClass); oneof != nil {
		out.Driver = oneof
	}
	if oneof := Task_SparkTaskConfig_PythonScriptFile_ToProto(mapCtx, in.PythonScriptFile); oneof != nil {
		out.Driver = oneof
	}
	if oneof := Task_SparkTaskConfig_SqlScriptFile_ToProto(mapCtx, in.SQLScriptFile); oneof != nil {
		out.Driver = oneof
	}
	if oneof := Task_SparkTaskConfig_SqlScript_ToProto(mapCtx, in.SQLScript); oneof != nil {
		out.Driver = oneof
	}
	out.FileUris = in.FileUris
	out.ArchiveUris = in.ArchiveUris
	out.InfrastructureSpec = Task_InfrastructureSpec_ToProto(mapCtx, in.InfrastructureSpec)
	return out
}
func Task_TriggerSpec_FromProto(mapCtx *direct.MapContext, in *pb.Task_TriggerSpec) *krm.Task_TriggerSpec {
	if in == nil {
		return nil
	}
	out := &krm.Task_TriggerSpec{}
	out.Type = direct.Enum_FromProto(mapCtx, in.GetType())
	out.StartTime = direct.StringTimestamp_FromProto(mapCtx, in.GetStartTime())
	out.Disabled = direct.LazyPtr(in.GetDisabled())
	out.MaxRetries = direct.LazyPtr(in.GetMaxRetries())
	out.Schedule = direct.LazyPtr(in.GetSchedule())
	return out
}
func Task_TriggerSpec_ToProto(mapCtx *direct.MapContext, in *krm.Task_TriggerSpec) *pb.Task_TriggerSpec {
	if in == nil {
		return nil
	}
	out := &pb.Task_TriggerSpec{}
	out.Type = direct.Enum_ToProto[pb.Task_TriggerSpec_Type](mapCtx, in.Type)
	out.StartTime = direct.StringTimestamp_ToProto(mapCtx, in.StartTime)
	out.Disabled = direct.ValueOf(in.Disabled)
	out.MaxRetries = direct.ValueOf(in.MaxRetries)
	if oneof := Task_TriggerSpec_Schedule_ToProto(mapCtx, in.Schedule); oneof != nil {
		out.Trigger = oneof
	}
	return out
}
