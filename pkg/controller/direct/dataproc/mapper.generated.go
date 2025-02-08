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
func DataprocSessionObservedState_FromProto(mapCtx *direct.MapContext, in *pb.Session) *krm.DataprocSessionObservedState {
	if in == nil {
		return nil
	}
	out := &krm.DataprocSessionObservedState{}
	// MISSING: Name
	// MISSING: Uuid
	// MISSING: CreateTime
	// MISSING: JupyterSession
	// MISSING: SparkConnectSession
	// MISSING: RuntimeInfo
	// MISSING: State
	// MISSING: StateMessage
	// MISSING: StateTime
	// MISSING: Creator
	// MISSING: Labels
	// MISSING: RuntimeConfig
	// MISSING: EnvironmentConfig
	// MISSING: User
	// MISSING: StateHistory
	// MISSING: SessionTemplate
	return out
}
func DataprocSessionObservedState_ToProto(mapCtx *direct.MapContext, in *krm.DataprocSessionObservedState) *pb.Session {
	if in == nil {
		return nil
	}
	out := &pb.Session{}
	// MISSING: Name
	// MISSING: Uuid
	// MISSING: CreateTime
	// MISSING: JupyterSession
	// MISSING: SparkConnectSession
	// MISSING: RuntimeInfo
	// MISSING: State
	// MISSING: StateMessage
	// MISSING: StateTime
	// MISSING: Creator
	// MISSING: Labels
	// MISSING: RuntimeConfig
	// MISSING: EnvironmentConfig
	// MISSING: User
	// MISSING: StateHistory
	// MISSING: SessionTemplate
	return out
}
func DataprocSessionSpec_FromProto(mapCtx *direct.MapContext, in *pb.Session) *krm.DataprocSessionSpec {
	if in == nil {
		return nil
	}
	out := &krm.DataprocSessionSpec{}
	// MISSING: Name
	// MISSING: Uuid
	// MISSING: CreateTime
	// MISSING: JupyterSession
	// MISSING: SparkConnectSession
	// MISSING: RuntimeInfo
	// MISSING: State
	// MISSING: StateMessage
	// MISSING: StateTime
	// MISSING: Creator
	// MISSING: Labels
	// MISSING: RuntimeConfig
	// MISSING: EnvironmentConfig
	// MISSING: User
	// MISSING: StateHistory
	// MISSING: SessionTemplate
	return out
}
func DataprocSessionSpec_ToProto(mapCtx *direct.MapContext, in *krm.DataprocSessionSpec) *pb.Session {
	if in == nil {
		return nil
	}
	out := &pb.Session{}
	// MISSING: Name
	// MISSING: Uuid
	// MISSING: CreateTime
	// MISSING: JupyterSession
	// MISSING: SparkConnectSession
	// MISSING: RuntimeInfo
	// MISSING: State
	// MISSING: StateMessage
	// MISSING: StateTime
	// MISSING: Creator
	// MISSING: Labels
	// MISSING: RuntimeConfig
	// MISSING: EnvironmentConfig
	// MISSING: User
	// MISSING: StateHistory
	// MISSING: SessionTemplate
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
func RuntimeInfo_FromProto(mapCtx *direct.MapContext, in *pb.RuntimeInfo) *krm.RuntimeInfo {
	if in == nil {
		return nil
	}
	out := &krm.RuntimeInfo{}
	// MISSING: Endpoints
	// MISSING: OutputURI
	// MISSING: DiagnosticOutputURI
	// MISSING: ApproximateUsage
	// MISSING: CurrentUsage
	return out
}
func RuntimeInfo_ToProto(mapCtx *direct.MapContext, in *krm.RuntimeInfo) *pb.RuntimeInfo {
	if in == nil {
		return nil
	}
	out := &pb.RuntimeInfo{}
	// MISSING: Endpoints
	// MISSING: OutputURI
	// MISSING: DiagnosticOutputURI
	// MISSING: ApproximateUsage
	// MISSING: CurrentUsage
	return out
}
func RuntimeInfoObservedState_FromProto(mapCtx *direct.MapContext, in *pb.RuntimeInfo) *krm.RuntimeInfoObservedState {
	if in == nil {
		return nil
	}
	out := &krm.RuntimeInfoObservedState{}
	out.Endpoints = in.Endpoints
	out.OutputURI = direct.LazyPtr(in.GetOutputUri())
	out.DiagnosticOutputURI = direct.LazyPtr(in.GetDiagnosticOutputUri())
	out.ApproximateUsage = UsageMetrics_FromProto(mapCtx, in.GetApproximateUsage())
	out.CurrentUsage = UsageSnapshot_FromProto(mapCtx, in.GetCurrentUsage())
	return out
}
func RuntimeInfoObservedState_ToProto(mapCtx *direct.MapContext, in *krm.RuntimeInfoObservedState) *pb.RuntimeInfo {
	if in == nil {
		return nil
	}
	out := &pb.RuntimeInfo{}
	out.Endpoints = in.Endpoints
	out.OutputUri = direct.ValueOf(in.OutputURI)
	out.DiagnosticOutputUri = direct.ValueOf(in.DiagnosticOutputURI)
	out.ApproximateUsage = UsageMetrics_ToProto(mapCtx, in.ApproximateUsage)
	out.CurrentUsage = UsageSnapshot_ToProto(mapCtx, in.CurrentUsage)
	return out
}
func Session_FromProto(mapCtx *direct.MapContext, in *pb.Session) *krm.Session {
	if in == nil {
		return nil
	}
	out := &krm.Session{}
	out.Name = direct.LazyPtr(in.GetName())
	// MISSING: Uuid
	// MISSING: CreateTime
	out.JupyterSession = JupyterConfig_FromProto(mapCtx, in.GetJupyterSession())
	out.SparkConnectSession = SparkConnectConfig_FromProto(mapCtx, in.GetSparkConnectSession())
	// MISSING: RuntimeInfo
	// MISSING: State
	// MISSING: StateMessage
	// MISSING: StateTime
	// MISSING: Creator
	out.Labels = in.Labels
	out.RuntimeConfig = RuntimeConfig_FromProto(mapCtx, in.GetRuntimeConfig())
	out.EnvironmentConfig = EnvironmentConfig_FromProto(mapCtx, in.GetEnvironmentConfig())
	out.User = direct.LazyPtr(in.GetUser())
	// MISSING: StateHistory
	out.SessionTemplate = direct.LazyPtr(in.GetSessionTemplate())
	return out
}
func Session_ToProto(mapCtx *direct.MapContext, in *krm.Session) *pb.Session {
	if in == nil {
		return nil
	}
	out := &pb.Session{}
	out.Name = direct.ValueOf(in.Name)
	// MISSING: Uuid
	// MISSING: CreateTime
	if oneof := JupyterConfig_ToProto(mapCtx, in.JupyterSession); oneof != nil {
		out.SessionConfig = &pb.Session_JupyterSession{JupyterSession: oneof}
	}
	if oneof := SparkConnectConfig_ToProto(mapCtx, in.SparkConnectSession); oneof != nil {
		out.SessionConfig = &pb.Session_SparkConnectSession{SparkConnectSession: oneof}
	}
	// MISSING: RuntimeInfo
	// MISSING: State
	// MISSING: StateMessage
	// MISSING: StateTime
	// MISSING: Creator
	out.Labels = in.Labels
	out.RuntimeConfig = RuntimeConfig_ToProto(mapCtx, in.RuntimeConfig)
	out.EnvironmentConfig = EnvironmentConfig_ToProto(mapCtx, in.EnvironmentConfig)
	out.User = direct.ValueOf(in.User)
	// MISSING: StateHistory
	out.SessionTemplate = direct.ValueOf(in.SessionTemplate)
	return out
}
func SessionObservedState_FromProto(mapCtx *direct.MapContext, in *pb.Session) *krm.SessionObservedState {
	if in == nil {
		return nil
	}
	out := &krm.SessionObservedState{}
	// MISSING: Name
	out.Uuid = direct.LazyPtr(in.GetUuid())
	out.CreateTime = direct.StringTimestamp_FromProto(mapCtx, in.GetCreateTime())
	// MISSING: JupyterSession
	// MISSING: SparkConnectSession
	out.RuntimeInfo = RuntimeInfo_FromProto(mapCtx, in.GetRuntimeInfo())
	out.State = direct.Enum_FromProto(mapCtx, in.GetState())
	out.StateMessage = direct.LazyPtr(in.GetStateMessage())
	out.StateTime = direct.StringTimestamp_FromProto(mapCtx, in.GetStateTime())
	out.Creator = direct.LazyPtr(in.GetCreator())
	// MISSING: Labels
	// MISSING: RuntimeConfig
	// MISSING: EnvironmentConfig
	// MISSING: User
	out.StateHistory = direct.Slice_FromProto(mapCtx, in.StateHistory, Session_SessionStateHistory_FromProto)
	// MISSING: SessionTemplate
	return out
}
func SessionObservedState_ToProto(mapCtx *direct.MapContext, in *krm.SessionObservedState) *pb.Session {
	if in == nil {
		return nil
	}
	out := &pb.Session{}
	// MISSING: Name
	out.Uuid = direct.ValueOf(in.Uuid)
	out.CreateTime = direct.StringTimestamp_ToProto(mapCtx, in.CreateTime)
	// MISSING: JupyterSession
	// MISSING: SparkConnectSession
	out.RuntimeInfo = RuntimeInfo_ToProto(mapCtx, in.RuntimeInfo)
	out.State = direct.Enum_ToProto[pb.Session_State](mapCtx, in.State)
	out.StateMessage = direct.ValueOf(in.StateMessage)
	out.StateTime = direct.StringTimestamp_ToProto(mapCtx, in.StateTime)
	out.Creator = direct.ValueOf(in.Creator)
	// MISSING: Labels
	// MISSING: RuntimeConfig
	// MISSING: EnvironmentConfig
	// MISSING: User
	out.StateHistory = direct.Slice_ToProto(mapCtx, in.StateHistory, Session_SessionStateHistory_ToProto)
	// MISSING: SessionTemplate
	return out
}
func Session_SessionStateHistory_FromProto(mapCtx *direct.MapContext, in *pb.Session_SessionStateHistory) *krm.Session_SessionStateHistory {
	if in == nil {
		return nil
	}
	out := &krm.Session_SessionStateHistory{}
	// MISSING: State
	// MISSING: StateMessage
	// MISSING: StateStartTime
	return out
}
func Session_SessionStateHistory_ToProto(mapCtx *direct.MapContext, in *krm.Session_SessionStateHistory) *pb.Session_SessionStateHistory {
	if in == nil {
		return nil
	}
	out := &pb.Session_SessionStateHistory{}
	// MISSING: State
	// MISSING: StateMessage
	// MISSING: StateStartTime
	return out
}
func Session_SessionStateHistoryObservedState_FromProto(mapCtx *direct.MapContext, in *pb.Session_SessionStateHistory) *krm.Session_SessionStateHistoryObservedState {
	if in == nil {
		return nil
	}
	out := &krm.Session_SessionStateHistoryObservedState{}
	out.State = direct.Enum_FromProto(mapCtx, in.GetState())
	out.StateMessage = direct.LazyPtr(in.GetStateMessage())
	out.StateStartTime = direct.StringTimestamp_FromProto(mapCtx, in.GetStateStartTime())
	return out
}
func Session_SessionStateHistoryObservedState_ToProto(mapCtx *direct.MapContext, in *krm.Session_SessionStateHistoryObservedState) *pb.Session_SessionStateHistory {
	if in == nil {
		return nil
	}
	out := &pb.Session_SessionStateHistory{}
	out.State = direct.Enum_ToProto[pb.Session_State](mapCtx, in.State)
	out.StateMessage = direct.ValueOf(in.StateMessage)
	out.StateStartTime = direct.StringTimestamp_ToProto(mapCtx, in.StateStartTime)
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
func UsageMetrics_FromProto(mapCtx *direct.MapContext, in *pb.UsageMetrics) *krm.UsageMetrics {
	if in == nil {
		return nil
	}
	out := &krm.UsageMetrics{}
	out.MilliDcuSeconds = direct.LazyPtr(in.GetMilliDcuSeconds())
	out.ShuffleStorageGBSeconds = direct.LazyPtr(in.GetShuffleStorageGbSeconds())
	out.MilliAcceleratorSeconds = direct.LazyPtr(in.GetMilliAcceleratorSeconds())
	out.AcceleratorType = direct.LazyPtr(in.GetAcceleratorType())
	return out
}
func UsageMetrics_ToProto(mapCtx *direct.MapContext, in *krm.UsageMetrics) *pb.UsageMetrics {
	if in == nil {
		return nil
	}
	out := &pb.UsageMetrics{}
	out.MilliDcuSeconds = direct.ValueOf(in.MilliDcuSeconds)
	out.ShuffleStorageGbSeconds = direct.ValueOf(in.ShuffleStorageGBSeconds)
	out.MilliAcceleratorSeconds = direct.ValueOf(in.MilliAcceleratorSeconds)
	out.AcceleratorType = direct.ValueOf(in.AcceleratorType)
	return out
}
func UsageSnapshot_FromProto(mapCtx *direct.MapContext, in *pb.UsageSnapshot) *krm.UsageSnapshot {
	if in == nil {
		return nil
	}
	out := &krm.UsageSnapshot{}
	out.MilliDcu = direct.LazyPtr(in.GetMilliDcu())
	out.ShuffleStorageGB = direct.LazyPtr(in.GetShuffleStorageGb())
	out.MilliDcuPremium = direct.LazyPtr(in.GetMilliDcuPremium())
	out.ShuffleStorageGBPremium = direct.LazyPtr(in.GetShuffleStorageGbPremium())
	out.MilliAccelerator = direct.LazyPtr(in.GetMilliAccelerator())
	out.AcceleratorType = direct.LazyPtr(in.GetAcceleratorType())
	out.SnapshotTime = direct.StringTimestamp_FromProto(mapCtx, in.GetSnapshotTime())
	return out
}
func UsageSnapshot_ToProto(mapCtx *direct.MapContext, in *krm.UsageSnapshot) *pb.UsageSnapshot {
	if in == nil {
		return nil
	}
	out := &pb.UsageSnapshot{}
	out.MilliDcu = direct.ValueOf(in.MilliDcu)
	out.ShuffleStorageGb = direct.ValueOf(in.ShuffleStorageGB)
	out.MilliDcuPremium = direct.ValueOf(in.MilliDcuPremium)
	out.ShuffleStorageGbPremium = direct.ValueOf(in.ShuffleStorageGBPremium)
	out.MilliAccelerator = direct.ValueOf(in.MilliAccelerator)
	out.AcceleratorType = direct.ValueOf(in.AcceleratorType)
	out.SnapshotTime = direct.StringTimestamp_ToProto(mapCtx, in.SnapshotTime)
	return out
}
