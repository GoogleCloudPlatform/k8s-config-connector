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

package dataproc

import (
	refs "github.com/GoogleCloudPlatform/k8s-config-connector/apis/refs/v1beta1"
	pb "cloud.google.com/go/dataproc/v2/apiv1/dataprocpb"
	krm "github.com/GoogleCloudPlatform/k8s-config-connector/apis/dataproc/v1alpha1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct"
)
func AutotuningConfig_FromProto(mapCtx *direct.MapContext, in *pb.AutotuningConfig) *krm.AutotuningConfig {
	if in == nil {
		return nil
	}
	out := &krm.AutotuningConfig{}
	out.Scenarios = direct.EnumSlice_FromProto(mapCtx, in.Scenarios)
	return out
}
func AutotuningConfig_ToProto(mapCtx *direct.MapContext, in *krm.AutotuningConfig) *pb.AutotuningConfig {
	if in == nil {
		return nil
	}
	out := &pb.AutotuningConfig{}
	out.Scenarios = direct.EnumSlice_ToProto[pb.AutotuningConfig_Scenario](mapCtx, in.Scenarios)
	return out
}
func DataprocSessionTemplateObservedState_FromProto(mapCtx *direct.MapContext, in *pb.SessionTemplate) *krm.DataprocSessionTemplateObservedState {
	if in == nil {
		return nil
	}
	out := &krm.DataprocSessionTemplateObservedState{}
	// MISSING: Name
	// MISSING: Description
	// MISSING: CreateTime
	// MISSING: JupyterSession
	// MISSING: SparkConnectSession
	// MISSING: Creator
	// MISSING: Labels
	// MISSING: RuntimeConfig
	// MISSING: EnvironmentConfig
	// MISSING: UpdateTime
	// MISSING: Uuid
	return out
}
func DataprocSessionTemplateObservedState_ToProto(mapCtx *direct.MapContext, in *krm.DataprocSessionTemplateObservedState) *pb.SessionTemplate {
	if in == nil {
		return nil
	}
	out := &pb.SessionTemplate{}
	// MISSING: Name
	// MISSING: Description
	// MISSING: CreateTime
	// MISSING: JupyterSession
	// MISSING: SparkConnectSession
	// MISSING: Creator
	// MISSING: Labels
	// MISSING: RuntimeConfig
	// MISSING: EnvironmentConfig
	// MISSING: UpdateTime
	// MISSING: Uuid
	return out
}
func DataprocSessionTemplateSpec_FromProto(mapCtx *direct.MapContext, in *pb.SessionTemplate) *krm.DataprocSessionTemplateSpec {
	if in == nil {
		return nil
	}
	out := &krm.DataprocSessionTemplateSpec{}
	// MISSING: Name
	// MISSING: Description
	// MISSING: CreateTime
	// MISSING: JupyterSession
	// MISSING: SparkConnectSession
	// MISSING: Creator
	// MISSING: Labels
	// MISSING: RuntimeConfig
	// MISSING: EnvironmentConfig
	// MISSING: UpdateTime
	// MISSING: Uuid
	return out
}
func DataprocSessionTemplateSpec_ToProto(mapCtx *direct.MapContext, in *krm.DataprocSessionTemplateSpec) *pb.SessionTemplate {
	if in == nil {
		return nil
	}
	out := &pb.SessionTemplate{}
	// MISSING: Name
	// MISSING: Description
	// MISSING: CreateTime
	// MISSING: JupyterSession
	// MISSING: SparkConnectSession
	// MISSING: Creator
	// MISSING: Labels
	// MISSING: RuntimeConfig
	// MISSING: EnvironmentConfig
	// MISSING: UpdateTime
	// MISSING: Uuid
	return out
}
func EnvironmentConfig_FromProto(mapCtx *direct.MapContext, in *pb.EnvironmentConfig) *krm.EnvironmentConfig {
	if in == nil {
		return nil
	}
	out := &krm.EnvironmentConfig{}
	out.ExecutionConfig = ExecutionConfig_FromProto(mapCtx, in.GetExecutionConfig())
	out.PeripheralsConfig = PeripheralsConfig_FromProto(mapCtx, in.GetPeripheralsConfig())
	return out
}
func EnvironmentConfig_ToProto(mapCtx *direct.MapContext, in *krm.EnvironmentConfig) *pb.EnvironmentConfig {
	if in == nil {
		return nil
	}
	out := &pb.EnvironmentConfig{}
	out.ExecutionConfig = ExecutionConfig_ToProto(mapCtx, in.ExecutionConfig)
	out.PeripheralsConfig = PeripheralsConfig_ToProto(mapCtx, in.PeripheralsConfig)
	return out
}
func ExecutionConfig_FromProto(mapCtx *direct.MapContext, in *pb.ExecutionConfig) *krm.ExecutionConfig {
	if in == nil {
		return nil
	}
	out := &krm.ExecutionConfig{}
	out.ServiceAccount = direct.LazyPtr(in.GetServiceAccount())
	out.NetworkURI = direct.LazyPtr(in.GetNetworkUri())
	out.SubnetworkURI = direct.LazyPtr(in.GetSubnetworkUri())
	out.NetworkTags = in.NetworkTags
	out.KMSKey = direct.LazyPtr(in.GetKmsKey())
	out.IdleTtl = direct.StringDuration_FromProto(mapCtx, in.GetIdleTtl())
	out.Ttl = direct.StringDuration_FromProto(mapCtx, in.GetTtl())
	out.StagingBucket = direct.LazyPtr(in.GetStagingBucket())
	return out
}
func ExecutionConfig_ToProto(mapCtx *direct.MapContext, in *krm.ExecutionConfig) *pb.ExecutionConfig {
	if in == nil {
		return nil
	}
	out := &pb.ExecutionConfig{}
	out.ServiceAccount = direct.ValueOf(in.ServiceAccount)
	if oneof := ExecutionConfig_NetworkUri_ToProto(mapCtx, in.NetworkURI); oneof != nil {
		out.Network = oneof
	}
	if oneof := ExecutionConfig_SubnetworkUri_ToProto(mapCtx, in.SubnetworkURI); oneof != nil {
		out.Network = oneof
	}
	out.NetworkTags = in.NetworkTags
	out.KmsKey = direct.ValueOf(in.KMSKey)
	out.IdleTtl = direct.StringDuration_ToProto(mapCtx, in.IdleTtl)
	out.Ttl = direct.StringDuration_ToProto(mapCtx, in.Ttl)
	out.StagingBucket = direct.ValueOf(in.StagingBucket)
	return out
}
func JupyterConfig_FromProto(mapCtx *direct.MapContext, in *pb.JupyterConfig) *krm.JupyterConfig {
	if in == nil {
		return nil
	}
	out := &krm.JupyterConfig{}
	out.Kernel = direct.Enum_FromProto(mapCtx, in.GetKernel())
	out.DisplayName = direct.LazyPtr(in.GetDisplayName())
	return out
}
func JupyterConfig_ToProto(mapCtx *direct.MapContext, in *krm.JupyterConfig) *pb.JupyterConfig {
	if in == nil {
		return nil
	}
	out := &pb.JupyterConfig{}
	out.Kernel = direct.Enum_ToProto[pb.JupyterConfig_Kernel](mapCtx, in.Kernel)
	out.DisplayName = direct.ValueOf(in.DisplayName)
	return out
}
func PeripheralsConfig_FromProto(mapCtx *direct.MapContext, in *pb.PeripheralsConfig) *krm.PeripheralsConfig {
	if in == nil {
		return nil
	}
	out := &krm.PeripheralsConfig{}
	out.MetastoreService = direct.LazyPtr(in.GetMetastoreService())
	out.SparkHistoryServerConfig = SparkHistoryServerConfig_FromProto(mapCtx, in.GetSparkHistoryServerConfig())
	return out
}
func PeripheralsConfig_ToProto(mapCtx *direct.MapContext, in *krm.PeripheralsConfig) *pb.PeripheralsConfig {
	if in == nil {
		return nil
	}
	out := &pb.PeripheralsConfig{}
	out.MetastoreService = direct.ValueOf(in.MetastoreService)
	out.SparkHistoryServerConfig = SparkHistoryServerConfig_ToProto(mapCtx, in.SparkHistoryServerConfig)
	return out
}
func PyPiRepositoryConfig_FromProto(mapCtx *direct.MapContext, in *pb.PyPiRepositoryConfig) *krm.PyPiRepositoryConfig {
	if in == nil {
		return nil
	}
	out := &krm.PyPiRepositoryConfig{}
	out.PypiRepository = direct.LazyPtr(in.GetPypiRepository())
	return out
}
func PyPiRepositoryConfig_ToProto(mapCtx *direct.MapContext, in *krm.PyPiRepositoryConfig) *pb.PyPiRepositoryConfig {
	if in == nil {
		return nil
	}
	out := &pb.PyPiRepositoryConfig{}
	out.PypiRepository = direct.ValueOf(in.PypiRepository)
	return out
}
func RepositoryConfig_FromProto(mapCtx *direct.MapContext, in *pb.RepositoryConfig) *krm.RepositoryConfig {
	if in == nil {
		return nil
	}
	out := &krm.RepositoryConfig{}
	out.PypiRepositoryConfig = PyPiRepositoryConfig_FromProto(mapCtx, in.GetPypiRepositoryConfig())
	return out
}
func RepositoryConfig_ToProto(mapCtx *direct.MapContext, in *krm.RepositoryConfig) *pb.RepositoryConfig {
	if in == nil {
		return nil
	}
	out := &pb.RepositoryConfig{}
	out.PypiRepositoryConfig = PyPiRepositoryConfig_ToProto(mapCtx, in.PypiRepositoryConfig)
	return out
}
func RuntimeConfig_FromProto(mapCtx *direct.MapContext, in *pb.RuntimeConfig) *krm.RuntimeConfig {
	if in == nil {
		return nil
	}
	out := &krm.RuntimeConfig{}
	out.Version = direct.LazyPtr(in.GetVersion())
	out.ContainerImage = direct.LazyPtr(in.GetContainerImage())
	out.Properties = in.Properties
	out.RepositoryConfig = RepositoryConfig_FromProto(mapCtx, in.GetRepositoryConfig())
	out.AutotuningConfig = AutotuningConfig_FromProto(mapCtx, in.GetAutotuningConfig())
	out.Cohort = direct.LazyPtr(in.GetCohort())
	return out
}
func RuntimeConfig_ToProto(mapCtx *direct.MapContext, in *krm.RuntimeConfig) *pb.RuntimeConfig {
	if in == nil {
		return nil
	}
	out := &pb.RuntimeConfig{}
	out.Version = direct.ValueOf(in.Version)
	out.ContainerImage = direct.ValueOf(in.ContainerImage)
	out.Properties = in.Properties
	out.RepositoryConfig = RepositoryConfig_ToProto(mapCtx, in.RepositoryConfig)
	out.AutotuningConfig = AutotuningConfig_ToProto(mapCtx, in.AutotuningConfig)
	out.Cohort = direct.ValueOf(in.Cohort)
	return out
}
func SessionTemplate_FromProto(mapCtx *direct.MapContext, in *pb.SessionTemplate) *krm.SessionTemplate {
	if in == nil {
		return nil
	}
	out := &krm.SessionTemplate{}
	out.Name = direct.LazyPtr(in.GetName())
	out.Description = direct.LazyPtr(in.GetDescription())
	// MISSING: CreateTime
	out.JupyterSession = JupyterConfig_FromProto(mapCtx, in.GetJupyterSession())
	out.SparkConnectSession = SparkConnectConfig_FromProto(mapCtx, in.GetSparkConnectSession())
	// MISSING: Creator
	out.Labels = in.Labels
	out.RuntimeConfig = RuntimeConfig_FromProto(mapCtx, in.GetRuntimeConfig())
	out.EnvironmentConfig = EnvironmentConfig_FromProto(mapCtx, in.GetEnvironmentConfig())
	// MISSING: UpdateTime
	// MISSING: Uuid
	return out
}
func SessionTemplate_ToProto(mapCtx *direct.MapContext, in *krm.SessionTemplate) *pb.SessionTemplate {
	if in == nil {
		return nil
	}
	out := &pb.SessionTemplate{}
	out.Name = direct.ValueOf(in.Name)
	out.Description = direct.ValueOf(in.Description)
	// MISSING: CreateTime
	if oneof := JupyterConfig_ToProto(mapCtx, in.JupyterSession); oneof != nil {
		out.SessionConfig = &pb.SessionTemplate_JupyterSession{JupyterSession: oneof}
	}
	if oneof := SparkConnectConfig_ToProto(mapCtx, in.SparkConnectSession); oneof != nil {
		out.SessionConfig = &pb.SessionTemplate_SparkConnectSession{SparkConnectSession: oneof}
	}
	// MISSING: Creator
	out.Labels = in.Labels
	out.RuntimeConfig = RuntimeConfig_ToProto(mapCtx, in.RuntimeConfig)
	out.EnvironmentConfig = EnvironmentConfig_ToProto(mapCtx, in.EnvironmentConfig)
	// MISSING: UpdateTime
	// MISSING: Uuid
	return out
}
func SessionTemplateObservedState_FromProto(mapCtx *direct.MapContext, in *pb.SessionTemplate) *krm.SessionTemplateObservedState {
	if in == nil {
		return nil
	}
	out := &krm.SessionTemplateObservedState{}
	// MISSING: Name
	// MISSING: Description
	out.CreateTime = direct.StringTimestamp_FromProto(mapCtx, in.GetCreateTime())
	// MISSING: JupyterSession
	// MISSING: SparkConnectSession
	out.Creator = direct.LazyPtr(in.GetCreator())
	// MISSING: Labels
	// MISSING: RuntimeConfig
	// MISSING: EnvironmentConfig
	out.UpdateTime = direct.StringTimestamp_FromProto(mapCtx, in.GetUpdateTime())
	out.Uuid = direct.LazyPtr(in.GetUuid())
	return out
}
func SessionTemplateObservedState_ToProto(mapCtx *direct.MapContext, in *krm.SessionTemplateObservedState) *pb.SessionTemplate {
	if in == nil {
		return nil
	}
	out := &pb.SessionTemplate{}
	// MISSING: Name
	// MISSING: Description
	out.CreateTime = direct.StringTimestamp_ToProto(mapCtx, in.CreateTime)
	// MISSING: JupyterSession
	// MISSING: SparkConnectSession
	out.Creator = direct.ValueOf(in.Creator)
	// MISSING: Labels
	// MISSING: RuntimeConfig
	// MISSING: EnvironmentConfig
	out.UpdateTime = direct.StringTimestamp_ToProto(mapCtx, in.UpdateTime)
	out.Uuid = direct.ValueOf(in.Uuid)
	return out
}
func SparkConnectConfig_FromProto(mapCtx *direct.MapContext, in *pb.SparkConnectConfig) *krm.SparkConnectConfig {
	if in == nil {
		return nil
	}
	out := &krm.SparkConnectConfig{}
	return out
}
func SparkConnectConfig_ToProto(mapCtx *direct.MapContext, in *krm.SparkConnectConfig) *pb.SparkConnectConfig {
	if in == nil {
		return nil
	}
	out := &pb.SparkConnectConfig{}
	return out
}
func SparkHistoryServerConfig_FromProto(mapCtx *direct.MapContext, in *pb.SparkHistoryServerConfig) *krm.SparkHistoryServerConfig {
	if in == nil {
		return nil
	}
	out := &krm.SparkHistoryServerConfig{}
	out.DataprocCluster = direct.LazyPtr(in.GetDataprocCluster())
	return out
}
func SparkHistoryServerConfig_ToProto(mapCtx *direct.MapContext, in *krm.SparkHistoryServerConfig) *pb.SparkHistoryServerConfig {
	if in == nil {
		return nil
	}
	out := &pb.SparkHistoryServerConfig{}
	out.DataprocCluster = direct.ValueOf(in.DataprocCluster)
	return out
}
