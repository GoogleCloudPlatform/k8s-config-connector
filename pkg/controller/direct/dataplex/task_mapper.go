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

// +generated:mapper
// krm.group: dataplex.cnrm.cloud.google.com
// krm.version: v1alpha1
// proto.service: google.cloud.dataplex.v1

package dataplex

import (
	pb "cloud.google.com/go/dataplex/apiv1/dataplexpb"
	krmv1alpha1 "github.com/GoogleCloudPlatform/k8s-config-connector/apis/dataplex/v1alpha1"
	kmsv1beta1 "github.com/GoogleCloudPlatform/k8s-config-connector/apis/kms/v1beta1"
	refsv1beta1 "github.com/GoogleCloudPlatform/k8s-config-connector/apis/refs/v1beta1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct"
)

func DataplexTaskSpec_FromProto(mapCtx *direct.MapContext, in *pb.Task) *krmv1alpha1.DataplexTaskSpec {
	if in == nil {
		return nil
	}
	out := &krmv1alpha1.DataplexTaskSpec{}
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
func DataplexTaskSpec_ToProto(mapCtx *direct.MapContext, in *krmv1alpha1.DataplexTaskSpec) *pb.Task {
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

func DataplexTaskObservedState_FromProto(mapCtx *direct.MapContext, in *pb.Task) *krmv1alpha1.DataplexTaskObservedState {
	if in == nil {
		return nil
	}
	out := &krmv1alpha1.DataplexTaskObservedState{}
	out.UID = direct.LazyPtr(in.GetUid())
	out.CreateTime = direct.StringTimestamp_FromProto(mapCtx, in.GetCreateTime())
	out.UpdateTime = direct.StringTimestamp_FromProto(mapCtx, in.GetUpdateTime())
	out.State = direct.Enum_FromProto(mapCtx, in.GetState())
	out.ExecutionStatus = Task_ExecutionStatusObservedState_FromProto(mapCtx, in.GetExecutionStatus())
	return out
}
func DataplexTaskObservedState_ToProto(mapCtx *direct.MapContext, in *krmv1alpha1.DataplexTaskObservedState) *pb.Task {
	if in == nil {
		return nil
	}
	out := &pb.Task{}
	out.Uid = direct.ValueOf(in.UID)
	out.CreateTime = direct.StringTimestamp_ToProto(mapCtx, in.CreateTime)
	out.UpdateTime = direct.StringTimestamp_ToProto(mapCtx, in.UpdateTime)
	out.State = direct.Enum_ToProto[pb.State](mapCtx, in.State)
	out.ExecutionStatus = Task_ExecutionStatusObservedState_ToProto(mapCtx, in.ExecutionStatus)
	return out
}

func JobObservedState_FromProto(mapCtx *direct.MapContext, in *pb.Job) *krmv1alpha1.JobObservedState {
	if in == nil {
		return nil
	}
	out := &krmv1alpha1.JobObservedState{}
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
	out.ExecutionSpec = Task_ExecutionSpecObservedStatus_FromProto(mapCtx, in.GetExecutionSpec())
	return out
}
func JobObservedState_ToProto(mapCtx *direct.MapContext, in *krmv1alpha1.JobObservedState) *pb.Job {
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
	out.ExecutionSpec = Task_ExecutionSpecObservedStatus_ToProto(mapCtx, in.ExecutionSpec)
	return out
}

func Task_ExecutionSpec_FromProto(mapCtx *direct.MapContext, in *pb.Task_ExecutionSpec) *krmv1alpha1.Task_ExecutionSpec {
	if in == nil {
		return nil
	}
	out := &krmv1alpha1.Task_ExecutionSpec{}
	out.Args = in.Args
	if in.ServiceAccount != "" {
		out.ServiceAccountRef = &refsv1beta1.IAMServiceAccountRef{External: in.GetServiceAccount()}
	}
	out.Project = direct.LazyPtr(in.GetProject())
	out.MaxJobExecutionLifetime = direct.StringDuration_FromProto(mapCtx, in.GetMaxJobExecutionLifetime())
	if in.GetKmsKey() != "" {
		out.KMSKeyRef = &kmsv1beta1.KMSKeyRef_OneOf{External: in.GetKmsKey()}
	}
	return out
}
func Task_ExecutionSpec_ToProto(mapCtx *direct.MapContext, in *krmv1alpha1.Task_ExecutionSpec) *pb.Task_ExecutionSpec {
	if in == nil {
		return nil
	}
	out := &pb.Task_ExecutionSpec{}
	out.Args = in.Args
	if in.ServiceAccountRef != nil {
		out.ServiceAccount = in.ServiceAccountRef.External
	}
	out.Project = direct.ValueOf(in.Project)
	out.MaxJobExecutionLifetime = direct.StringDuration_ToProto(mapCtx, in.MaxJobExecutionLifetime)
	if in.KMSKeyRef != nil {
		out.KmsKey = in.KMSKeyRef.External
	}
	return out
}

// Duplicate of Task_ExecutionSpec_FromProto, removing reference fields in observed status
func Task_ExecutionSpecObservedStatus_FromProto(mapCtx *direct.MapContext, in *pb.Task_ExecutionSpec) *krmv1alpha1.Task_ExecutionSpecObservedState {
	if in == nil {
		return nil
	}
	out := &krmv1alpha1.Task_ExecutionSpecObservedState{}
	out.Args = in.Args
	out.ServiceAccount = direct.LazyPtr(in.ServiceAccount)
	out.Project = direct.LazyPtr(in.GetProject())
	out.MaxJobExecutionLifetime = direct.StringDuration_FromProto(mapCtx, in.GetMaxJobExecutionLifetime())
	out.KMSKey = direct.LazyPtr(in.GetKmsKey())
	return out
}

// Duplicate of Task_ExecutionSpec_ToProto, removing reference fields in observed status
func Task_ExecutionSpecObservedStatus_ToProto(mapCtx *direct.MapContext, in *krmv1alpha1.Task_ExecutionSpecObservedState) *pb.Task_ExecutionSpec {
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
func Task_ExecutionStatusObservedState_FromProto(mapCtx *direct.MapContext, in *pb.Task_ExecutionStatus) *krmv1alpha1.Task_ExecutionStatusObservedState {
	if in == nil {
		return nil
	}
	out := &krmv1alpha1.Task_ExecutionStatusObservedState{}
	out.UpdateTime = direct.StringTimestamp_FromProto(mapCtx, in.GetUpdateTime())
	out.LatestJob = JobObservedState_FromProto(mapCtx, in.GetLatestJob())
	return out
}
func Task_ExecutionStatusObservedState_ToProto(mapCtx *direct.MapContext, in *krmv1alpha1.Task_ExecutionStatusObservedState) *pb.Task_ExecutionStatus {
	if in == nil {
		return nil
	}
	out := &pb.Task_ExecutionStatus{}
	out.UpdateTime = direct.StringTimestamp_ToProto(mapCtx, in.UpdateTime)
	out.LatestJob = JobObservedState_ToProto(mapCtx, in.LatestJob)
	return out
}
func Task_InfrastructureSpec_FromProto(mapCtx *direct.MapContext, in *pb.Task_InfrastructureSpec) *krmv1alpha1.Task_InfrastructureSpec {
	if in == nil {
		return nil
	}
	out := &krmv1alpha1.Task_InfrastructureSpec{}
	out.Batch = Task_InfrastructureSpec_BatchComputeResources_FromProto(mapCtx, in.GetBatch())
	out.ContainerImage = Task_InfrastructureSpec_ContainerImageRuntime_FromProto(mapCtx, in.GetContainerImage())
	out.VpcNetwork = Task_InfrastructureSpec_VpcNetwork_FromProto(mapCtx, in.GetVpcNetwork())
	return out
}
func Task_InfrastructureSpec_ToProto(mapCtx *direct.MapContext, in *krmv1alpha1.Task_InfrastructureSpec) *pb.Task_InfrastructureSpec {
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
func Task_InfrastructureSpec_BatchComputeResources_FromProto(mapCtx *direct.MapContext, in *pb.Task_InfrastructureSpec_BatchComputeResources) *krmv1alpha1.Task_InfrastructureSpec_BatchComputeResources {
	if in == nil {
		return nil
	}
	out := &krmv1alpha1.Task_InfrastructureSpec_BatchComputeResources{}
	out.ExecutorsCount = direct.LazyPtr(in.GetExecutorsCount())
	out.MaxExecutorsCount = direct.LazyPtr(in.GetMaxExecutorsCount())
	return out
}
func Task_InfrastructureSpec_BatchComputeResources_ToProto(mapCtx *direct.MapContext, in *krmv1alpha1.Task_InfrastructureSpec_BatchComputeResources) *pb.Task_InfrastructureSpec_BatchComputeResources {
	if in == nil {
		return nil
	}
	out := &pb.Task_InfrastructureSpec_BatchComputeResources{}
	out.ExecutorsCount = direct.ValueOf(in.ExecutorsCount)
	out.MaxExecutorsCount = direct.ValueOf(in.MaxExecutorsCount)
	return out
}
func Task_InfrastructureSpec_ContainerImageRuntime_FromProto(mapCtx *direct.MapContext, in *pb.Task_InfrastructureSpec_ContainerImageRuntime) *krmv1alpha1.Task_InfrastructureSpec_ContainerImageRuntime {
	if in == nil {
		return nil
	}
	out := &krmv1alpha1.Task_InfrastructureSpec_ContainerImageRuntime{}
	out.Image = direct.LazyPtr(in.GetImage())
	out.JavaJars = in.JavaJars
	out.PythonPackages = in.PythonPackages
	out.Properties = in.Properties
	return out
}
func Task_InfrastructureSpec_ContainerImageRuntime_ToProto(mapCtx *direct.MapContext, in *krmv1alpha1.Task_InfrastructureSpec_ContainerImageRuntime) *pb.Task_InfrastructureSpec_ContainerImageRuntime {
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
func Task_InfrastructureSpec_VpcNetwork_FromProto(mapCtx *direct.MapContext, in *pb.Task_InfrastructureSpec_VpcNetwork) *krmv1alpha1.Task_InfrastructureSpec_VpcNetwork {
	if in == nil {
		return nil
	}
	out := &krmv1alpha1.Task_InfrastructureSpec_VpcNetwork{}
	out.Network = direct.LazyPtr(in.GetNetwork())
	out.SubNetwork = direct.LazyPtr(in.GetSubNetwork())
	out.NetworkTags = in.NetworkTags
	return out
}
func Task_InfrastructureSpec_VpcNetwork_ToProto(mapCtx *direct.MapContext, in *krmv1alpha1.Task_InfrastructureSpec_VpcNetwork) *pb.Task_InfrastructureSpec_VpcNetwork {
	if in == nil {
		return nil
	}
	out := &pb.Task_InfrastructureSpec_VpcNetwork{}
	if oneof := in.Network; oneof != nil {
		out.NetworkName = &pb.Task_InfrastructureSpec_VpcNetwork_Network{Network: direct.ValueOf(oneof)}
	}
	if oneof := in.SubNetwork; oneof != nil {
		out.NetworkName = &pb.Task_InfrastructureSpec_VpcNetwork_SubNetwork{SubNetwork: direct.ValueOf(oneof)}
	}
	out.NetworkTags = in.NetworkTags
	return out
}
func Task_NotebookTaskConfig_FromProto(mapCtx *direct.MapContext, in *pb.Task_NotebookTaskConfig) *krmv1alpha1.Task_NotebookTaskConfig {
	if in == nil {
		return nil
	}
	out := &krmv1alpha1.Task_NotebookTaskConfig{}
	out.Notebook = direct.LazyPtr(in.GetNotebook())
	out.InfrastructureSpec = Task_InfrastructureSpec_FromProto(mapCtx, in.GetInfrastructureSpec())
	out.FileURIs = in.FileUris
	out.ArchiveURIs = in.ArchiveUris
	return out
}
func Task_NotebookTaskConfig_ToProto(mapCtx *direct.MapContext, in *krmv1alpha1.Task_NotebookTaskConfig) *pb.Task_NotebookTaskConfig {
	if in == nil {
		return nil
	}
	out := &pb.Task_NotebookTaskConfig{}
	out.Notebook = direct.ValueOf(in.Notebook)
	out.InfrastructureSpec = Task_InfrastructureSpec_ToProto(mapCtx, in.InfrastructureSpec)
	out.FileUris = in.FileURIs
	out.ArchiveUris = in.ArchiveURIs
	return out
}
func Task_SparkTaskConfig_FromProto(mapCtx *direct.MapContext, in *pb.Task_SparkTaskConfig) *krmv1alpha1.Task_SparkTaskConfig {
	if in == nil {
		return nil
	}
	out := &krmv1alpha1.Task_SparkTaskConfig{}
	out.MainJarFileURI = direct.LazyPtr(in.GetMainJarFileUri())
	out.MainClass = direct.LazyPtr(in.GetMainClass())
	out.PythonScriptFile = direct.LazyPtr(in.GetPythonScriptFile())
	out.SQLScriptFile = direct.LazyPtr(in.GetSqlScriptFile())
	out.SQLScript = direct.LazyPtr(in.GetSqlScript())
	out.FileURIs = in.FileUris
	out.ArchiveURIs = in.ArchiveUris
	out.InfrastructureSpec = Task_InfrastructureSpec_FromProto(mapCtx, in.GetInfrastructureSpec())
	return out
}
func Task_SparkTaskConfig_ToProto(mapCtx *direct.MapContext, in *krmv1alpha1.Task_SparkTaskConfig) *pb.Task_SparkTaskConfig {
	if in == nil {
		return nil
	}
	out := &pb.Task_SparkTaskConfig{}
	if oneof := in.MainJarFileURI; oneof != nil {
		out.Driver = &pb.Task_SparkTaskConfig_MainJarFileUri{MainJarFileUri: direct.ValueOf(oneof)}
	}
	if oneof := in.MainClass; oneof != nil {
		out.Driver = &pb.Task_SparkTaskConfig_MainClass{MainClass: direct.ValueOf(oneof)}
	}
	if oneof := in.PythonScriptFile; oneof != nil {
		out.Driver = &pb.Task_SparkTaskConfig_PythonScriptFile{PythonScriptFile: direct.ValueOf(oneof)}
	}
	if oneof := in.SQLScriptFile; oneof != nil {
		out.Driver = &pb.Task_SparkTaskConfig_SqlScriptFile{SqlScriptFile: direct.ValueOf(oneof)}
	}
	if oneof := in.SQLScript; oneof != nil {
		out.Driver = &pb.Task_SparkTaskConfig_SqlScript{SqlScript: direct.ValueOf(oneof)}
	}
	out.FileUris = in.FileURIs
	out.ArchiveUris = in.ArchiveURIs
	out.InfrastructureSpec = Task_InfrastructureSpec_ToProto(mapCtx, in.InfrastructureSpec)
	return out
}
func Task_TriggerSpec_FromProto(mapCtx *direct.MapContext, in *pb.Task_TriggerSpec) *krmv1alpha1.Task_TriggerSpec {
	if in == nil {
		return nil
	}
	out := &krmv1alpha1.Task_TriggerSpec{}
	out.Type = direct.Enum_FromProto(mapCtx, in.GetType())
	out.StartTime = direct.StringTimestamp_FromProto(mapCtx, in.GetStartTime())
	out.Disabled = direct.LazyPtr(in.GetDisabled())
	out.MaxRetries = direct.LazyPtr(in.GetMaxRetries())
	out.Schedule = direct.LazyPtr(in.GetSchedule())
	return out
}
func Task_TriggerSpec_ToProto(mapCtx *direct.MapContext, in *krmv1alpha1.Task_TriggerSpec) *pb.Task_TriggerSpec {
	if in == nil {
		return nil
	}
	out := &pb.Task_TriggerSpec{}
	out.Type = direct.Enum_ToProto[pb.Task_TriggerSpec_Type](mapCtx, in.Type)
	out.StartTime = direct.StringTimestamp_ToProto(mapCtx, in.StartTime)
	out.Disabled = direct.ValueOf(in.Disabled)
	out.MaxRetries = direct.ValueOf(in.MaxRetries)
	if oneof := in.Schedule; oneof != nil {
		out.Trigger = &pb.Task_TriggerSpec_Schedule{Schedule: direct.ValueOf(oneof)}
	}
	return out
}
