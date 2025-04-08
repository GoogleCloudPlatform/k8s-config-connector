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
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct"
	refs "github.com/GoogleCloudPlatform/k8s-config-connector/apis/refs/v1beta1"
	pb "cloud.google.com/go/dataplex/apiv1/dataplexpb"
	krmv1alpha1 "github.com/GoogleCloudPlatform/k8s-config-connector/apis/dataplex/v1alpha1"
)
func Content_FromProto(mapCtx *direct.MapContext, in *pb.Content) *krmv1alpha1.Content {
	if in == nil {
		return nil
	}
	out := &krmv1alpha1.Content{}
	// MISSING: Name
	// MISSING: Uid
	out.Path = direct.LazyPtr(in.GetPath())
	// MISSING: CreateTime
	// MISSING: UpdateTime
	out.Labels = in.Labels
	out.Description = direct.LazyPtr(in.GetDescription())
	out.DataText = direct.LazyPtr(in.GetDataText())
	out.SQLScript = Content_SQLScript_FromProto(mapCtx, in.GetSqlScript())
	out.Notebook = Content_Notebook_FromProto(mapCtx, in.GetNotebook())
	return out
}
func Content_ToProto(mapCtx *direct.MapContext, in *krmv1alpha1.Content) *pb.Content {
	if in == nil {
		return nil
	}
	out := &pb.Content{}
	// MISSING: Name
	// MISSING: Uid
	out.Path = direct.ValueOf(in.Path)
	// MISSING: CreateTime
	// MISSING: UpdateTime
	out.Labels = in.Labels
	out.Description = direct.ValueOf(in.Description)
	if oneof := Content_DataText_ToProto(mapCtx, in.DataText); oneof != nil {
		out.Data = oneof
	}
	if oneof := Content_SQLScript_ToProto(mapCtx, in.SQLScript); oneof != nil {
		out.Content = &pb.Content_SqlScript_{SqlScript: oneof}
	}
	if oneof := Content_Notebook_ToProto(mapCtx, in.Notebook); oneof != nil {
		out.Content = &pb.Content_Notebook_{Notebook: oneof}
	}
	return out
}
func ContentObservedState_FromProto(mapCtx *direct.MapContext, in *pb.Content) *krmv1alpha1.ContentObservedState {
	if in == nil {
		return nil
	}
	out := &krmv1alpha1.ContentObservedState{}
	out.Name = direct.LazyPtr(in.GetName())
	out.Uid = direct.LazyPtr(in.GetUid())
	// MISSING: Path
	out.CreateTime = direct.StringTimestamp_FromProto(mapCtx, in.GetCreateTime())
	out.UpdateTime = direct.StringTimestamp_FromProto(mapCtx, in.GetUpdateTime())
	// MISSING: Labels
	// MISSING: Description
	// MISSING: DataText
	// MISSING: SQLScript
	// MISSING: Notebook
	return out
}
func ContentObservedState_ToProto(mapCtx *direct.MapContext, in *krmv1alpha1.ContentObservedState) *pb.Content {
	if in == nil {
		return nil
	}
	out := &pb.Content{}
	out.Name = direct.ValueOf(in.Name)
	out.Uid = direct.ValueOf(in.Uid)
	// MISSING: Path
	out.CreateTime = direct.StringTimestamp_ToProto(mapCtx, in.CreateTime)
	out.UpdateTime = direct.StringTimestamp_ToProto(mapCtx, in.UpdateTime)
	// MISSING: Labels
	// MISSING: Description
	// MISSING: DataText
	// MISSING: SQLScript
	// MISSING: Notebook
	return out
}
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
func Environment_FromProto(mapCtx *direct.MapContext, in *pb.Environment) *krmv1alpha1.Environment {
	if in == nil {
		return nil
	}
	out := &krmv1alpha1.Environment{}
	// MISSING: Name
	out.DisplayName = direct.LazyPtr(in.GetDisplayName())
	// MISSING: Uid
	// MISSING: CreateTime
	// MISSING: UpdateTime
	out.Labels = in.Labels
	out.Description = direct.LazyPtr(in.GetDescription())
	// MISSING: State
	out.InfrastructureSpec = Environment_InfrastructureSpec_FromProto(mapCtx, in.GetInfrastructureSpec())
	out.SessionSpec = Environment_SessionSpec_FromProto(mapCtx, in.GetSessionSpec())
	// MISSING: SessionStatus
	// MISSING: Endpoints
	return out
}
func Environment_ToProto(mapCtx *direct.MapContext, in *krmv1alpha1.Environment) *pb.Environment {
	if in == nil {
		return nil
	}
	out := &pb.Environment{}
	// MISSING: Name
	out.DisplayName = direct.ValueOf(in.DisplayName)
	// MISSING: Uid
	// MISSING: CreateTime
	// MISSING: UpdateTime
	out.Labels = in.Labels
	out.Description = direct.ValueOf(in.Description)
	// MISSING: State
	out.InfrastructureSpec = Environment_InfrastructureSpec_ToProto(mapCtx, in.InfrastructureSpec)
	out.SessionSpec = Environment_SessionSpec_ToProto(mapCtx, in.SessionSpec)
	// MISSING: SessionStatus
	// MISSING: Endpoints
	return out
}
func EnvironmentObservedState_FromProto(mapCtx *direct.MapContext, in *pb.Environment) *krmv1alpha1.EnvironmentObservedState {
	if in == nil {
		return nil
	}
	out := &krmv1alpha1.EnvironmentObservedState{}
	out.Name = direct.LazyPtr(in.GetName())
	// MISSING: DisplayName
	out.Uid = direct.LazyPtr(in.GetUid())
	out.CreateTime = direct.StringTimestamp_FromProto(mapCtx, in.GetCreateTime())
	out.UpdateTime = direct.StringTimestamp_FromProto(mapCtx, in.GetUpdateTime())
	// MISSING: Labels
	// MISSING: Description
	out.State = direct.Enum_FromProto(mapCtx, in.GetState())
	// MISSING: InfrastructureSpec
	// MISSING: SessionSpec
	out.SessionStatus = Environment_SessionStatusObservedState_FromProto(mapCtx, in.GetSessionStatus())
	out.Endpoints = Environment_EndpointsObservedState_FromProto(mapCtx, in.GetEndpoints())
	return out
}
func EnvironmentObservedState_ToProto(mapCtx *direct.MapContext, in *krmv1alpha1.EnvironmentObservedState) *pb.Environment {
	if in == nil {
		return nil
	}
	out := &pb.Environment{}
	out.Name = direct.ValueOf(in.Name)
	// MISSING: DisplayName
	out.Uid = direct.ValueOf(in.Uid)
	out.CreateTime = direct.StringTimestamp_ToProto(mapCtx, in.CreateTime)
	out.UpdateTime = direct.StringTimestamp_ToProto(mapCtx, in.UpdateTime)
	// MISSING: Labels
	// MISSING: Description
	out.State = direct.Enum_ToProto[pb.State](mapCtx, in.State)
	// MISSING: InfrastructureSpec
	// MISSING: SessionSpec
	out.SessionStatus = Environment_SessionStatusObservedState_ToProto(mapCtx, in.SessionStatus)
	out.Endpoints = Environment_EndpointsObservedState_ToProto(mapCtx, in.Endpoints)
	return out
}
func Environment_Endpoints_FromProto(mapCtx *direct.MapContext, in *pb.Environment_Endpoints) *krmv1alpha1.Environment_Endpoints {
	if in == nil {
		return nil
	}
	out := &krmv1alpha1.Environment_Endpoints{}
	// MISSING: Notebooks
	// MISSING: SQL
	return out
}
func Environment_Endpoints_ToProto(mapCtx *direct.MapContext, in *krmv1alpha1.Environment_Endpoints) *pb.Environment_Endpoints {
	if in == nil {
		return nil
	}
	out := &pb.Environment_Endpoints{}
	// MISSING: Notebooks
	// MISSING: SQL
	return out
}
func Environment_SessionStatus_FromProto(mapCtx *direct.MapContext, in *pb.Environment_SessionStatus) *krmv1alpha1.Environment_SessionStatus {
	if in == nil {
		return nil
	}
	out := &krmv1alpha1.Environment_SessionStatus{}
	// MISSING: Active
	return out
}
func Environment_SessionStatus_ToProto(mapCtx *direct.MapContext, in *krmv1alpha1.Environment_SessionStatus) *pb.Environment_SessionStatus {
	if in == nil {
		return nil
	}
	out := &pb.Environment_SessionStatus{}
	// MISSING: Active
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
	out.ExecutionSpec = Task_ExecutionSpec_FromProto(mapCtx, in.GetExecutionSpec())
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
	out.ExecutionSpec = Task_ExecutionSpec_ToProto(mapCtx, in.ExecutionSpec)
	return out
}
func Lake_FromProto(mapCtx *direct.MapContext, in *pb.Lake) *krmv1alpha1.Lake {
	if in == nil {
		return nil
	}
	out := &krmv1alpha1.Lake{}
	// MISSING: Name
	out.DisplayName = direct.LazyPtr(in.GetDisplayName())
	// MISSING: Uid
	// MISSING: CreateTime
	// MISSING: UpdateTime
	out.Labels = in.Labels
	out.Description = direct.LazyPtr(in.GetDescription())
	// MISSING: State
	// MISSING: ServiceAccount
	out.Metastore = Lake_Metastore_FromProto(mapCtx, in.GetMetastore())
	// MISSING: AssetStatus
	// MISSING: MetastoreStatus
	return out
}
func Lake_ToProto(mapCtx *direct.MapContext, in *krmv1alpha1.Lake) *pb.Lake {
	if in == nil {
		return nil
	}
	out := &pb.Lake{}
	// MISSING: Name
	out.DisplayName = direct.ValueOf(in.DisplayName)
	// MISSING: Uid
	// MISSING: CreateTime
	// MISSING: UpdateTime
	out.Labels = in.Labels
	out.Description = direct.ValueOf(in.Description)
	// MISSING: State
	// MISSING: ServiceAccount
	out.Metastore = Lake_Metastore_ToProto(mapCtx, in.Metastore)
	// MISSING: AssetStatus
	// MISSING: MetastoreStatus
	return out
}
func LakeObservedState_FromProto(mapCtx *direct.MapContext, in *pb.Lake) *krmv1alpha1.LakeObservedState {
	if in == nil {
		return nil
	}
	out := &krmv1alpha1.LakeObservedState{}
	out.Name = direct.LazyPtr(in.GetName())
	// MISSING: DisplayName
	out.Uid = direct.LazyPtr(in.GetUid())
	out.CreateTime = direct.StringTimestamp_FromProto(mapCtx, in.GetCreateTime())
	out.UpdateTime = direct.StringTimestamp_FromProto(mapCtx, in.GetUpdateTime())
	// MISSING: Labels
	// MISSING: Description
	out.State = direct.Enum_FromProto(mapCtx, in.GetState())
	out.ServiceAccount = direct.LazyPtr(in.GetServiceAccount())
	// MISSING: Metastore
	out.AssetStatus = AssetStatus_FromProto(mapCtx, in.GetAssetStatus())
	out.MetastoreStatus = Lake_MetastoreStatus_FromProto(mapCtx, in.GetMetastoreStatus())
	return out
}
func LakeObservedState_ToProto(mapCtx *direct.MapContext, in *krmv1alpha1.LakeObservedState) *pb.Lake {
	if in == nil {
		return nil
	}
	out := &pb.Lake{}
	out.Name = direct.ValueOf(in.Name)
	// MISSING: DisplayName
	out.Uid = direct.ValueOf(in.Uid)
	out.CreateTime = direct.StringTimestamp_ToProto(mapCtx, in.CreateTime)
	out.UpdateTime = direct.StringTimestamp_ToProto(mapCtx, in.UpdateTime)
	// MISSING: Labels
	// MISSING: Description
	out.State = direct.Enum_ToProto[pb.State](mapCtx, in.State)
	out.ServiceAccount = direct.ValueOf(in.ServiceAccount)
	// MISSING: Metastore
	out.AssetStatus = AssetStatus_ToProto(mapCtx, in.AssetStatus)
	out.MetastoreStatus = Lake_MetastoreStatus_ToProto(mapCtx, in.MetastoreStatus)
	return out
}
func Task_ExecutionSpec_FromProto(mapCtx *direct.MapContext, in *pb.Task_ExecutionSpec) *krmv1alpha1.Task_ExecutionSpec {
	if in == nil {
		return nil
	}
	out := &krmv1alpha1.Task_ExecutionSpec{}
	out.Args = in.Args
	out.ServiceAccount = direct.LazyPtr(in.GetServiceAccount())
	out.Project = direct.LazyPtr(in.GetProject())
	out.MaxJobExecutionLifetime = direct.StringDuration_FromProto(mapCtx, in.GetMaxJobExecutionLifetime())
	if in.GetKmsKey() != "" {
		out.KMSKeyRef = &refs.*refsv1beta1.KMSCryptoKeyRef{External: in.GetKmsKey()}
	}
	return out
}
func Task_ExecutionSpec_ToProto(mapCtx *direct.MapContext, in *krmv1alpha1.Task_ExecutionSpec) *pb.Task_ExecutionSpec {
	if in == nil {
		return nil
	}
	out := &pb.Task_ExecutionSpec{}
	out.Args = in.Args
	out.ServiceAccount = direct.ValueOf(in.ServiceAccount)
	out.Project = direct.ValueOf(in.Project)
	out.MaxJobExecutionLifetime = direct.StringDuration_ToProto(mapCtx, in.MaxJobExecutionLifetime)
	if in.KMSKeyRef != nil {
		out.KmsKey = in.KMSKeyRef.External
	}
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
	if oneof := Task_InfrastructureSpec_VpcNetwork_Network_ToProto(mapCtx, in.Network); oneof != nil {
		out.NetworkName = oneof
	}
	if oneof := Task_InfrastructureSpec_VpcNetwork_SubNetwork_ToProto(mapCtx, in.SubNetwork); oneof != nil {
		out.NetworkName = oneof
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
	out.FileUris = in.FileUris
	out.ArchiveUris = in.ArchiveUris
	return out
}
func Task_NotebookTaskConfig_ToProto(mapCtx *direct.MapContext, in *krmv1alpha1.Task_NotebookTaskConfig) *pb.Task_NotebookTaskConfig {
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
	out.FileUris = in.FileUris
	out.ArchiveUris = in.ArchiveUris
	out.InfrastructureSpec = Task_InfrastructureSpec_FromProto(mapCtx, in.GetInfrastructureSpec())
	return out
}
func Task_SparkTaskConfig_ToProto(mapCtx *direct.MapContext, in *krmv1alpha1.Task_SparkTaskConfig) *pb.Task_SparkTaskConfig {
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
	if oneof := Task_TriggerSpec_Schedule_ToProto(mapCtx, in.Schedule); oneof != nil {
		out.Trigger = oneof
	}
	return out
}
