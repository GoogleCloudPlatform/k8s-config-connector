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
	pb "cloud.google.com/go/dataproc/v2/apiv1/dataprocpb"
	krm "github.com/GoogleCloudPlatform/k8s-config-connector/apis/dataproc/v1alpha1"
	dataprocv1beta1 "github.com/GoogleCloudPlatform/k8s-config-connector/apis/dataproc/v1beta1"
	refs "github.com/GoogleCloudPlatform/k8s-config-connector/apis/refs/v1beta1"
	storagev1beta1 "github.com/GoogleCloudPlatform/k8s-config-connector/apis/storage/v1beta1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct"
)

func ExecutionConfig_NetworkUri_ToProto(mapCtx *direct.MapContext, in *string) *pb.ExecutionConfig_NetworkUri {
	if in == nil {
		return nil
	}
	out := &pb.ExecutionConfig_NetworkUri{
		NetworkUri: direct.ValueOf(in),
	}
	return out
}

func ExecutionConfig_SubnetworkUri_ToProto(mapCtx *direct.MapContext, in *string) *pb.ExecutionConfig_SubnetworkUri {
	if in == nil {
		return nil
	}
	out := &pb.ExecutionConfig_SubnetworkUri{
		SubnetworkUri: direct.ValueOf(in),
	}
	return out
}

func SparkBatch_MainJarFileUri_ToProto(mapCtx *direct.MapContext, in *string) *pb.SparkBatch_MainJarFileUri {
	if in == nil {
		return nil
	}
	out := &pb.SparkBatch_MainJarFileUri{
		MainJarFileUri: direct.ValueOf(in),
	}
	return out
}

func SparkBatch_MainClass_ToProto(mapCtx *direct.MapContext, in *string) *pb.SparkBatch_MainClass {
	if in == nil {
		return nil
	}
	out := &pb.SparkBatch_MainClass{
		MainClass: direct.ValueOf(in),
	}
	return out
}

func PySparkBatch_FromProto(mapCtx *direct.MapContext, in *pb.PySparkBatch) *krm.PySparkBatch {
	if in == nil {
		return nil
	}
	out := &krm.PySparkBatch{}
	out.MainPythonFileURI = direct.LazyPtr(in.GetMainPythonFileUri())
	out.Args = in.Args
	out.PythonFileURIs = in.PythonFileUris
	out.JarFileURIs = in.JarFileUris
	out.FileURIs = in.FileUris
	out.ArchiveURIs = in.ArchiveUris
	return out
}
func PySparkBatch_ToProto(mapCtx *direct.MapContext, in *krm.PySparkBatch) *pb.PySparkBatch {
	if in == nil {
		return nil
	}
	out := &pb.PySparkBatch{}
	out.MainPythonFileUri = direct.ValueOf(in.MainPythonFileURI)
	out.Args = in.Args
	out.PythonFileUris = in.PythonFileURIs
	out.JarFileUris = in.JarFileURIs
	out.FileUris = in.FileURIs
	out.ArchiveUris = in.ArchiveURIs
	return out
}
func SparkBatch_FromProto(mapCtx *direct.MapContext, in *pb.SparkBatch) *krm.SparkBatch {
	if in == nil {
		return nil
	}
	out := &krm.SparkBatch{}
	out.MainJarFileURI = direct.LazyPtr(in.GetMainJarFileUri())
	out.MainClass = direct.LazyPtr(in.GetMainClass())
	out.Args = in.Args
	out.JarFileURIs = in.JarFileUris
	out.FileURIs = in.FileUris
	out.ArchiveURIs = in.ArchiveUris
	return out
}
func SparkBatch_ToProto(mapCtx *direct.MapContext, in *krm.SparkBatch) *pb.SparkBatch {
	if in == nil {
		return nil
	}
	out := &pb.SparkBatch{}
	if oneof := SparkBatch_MainJarFileUri_ToProto(mapCtx, in.MainJarFileURI); oneof != nil {
		out.Driver = oneof
	}
	if oneof := SparkBatch_MainClass_ToProto(mapCtx, in.MainClass); oneof != nil {
		out.Driver = oneof
	}
	out.Args = in.Args
	out.JarFileUris = in.JarFileURIs
	out.FileUris = in.FileURIs
	out.ArchiveUris = in.ArchiveURIs
	return out
}
func SparkRBatch_FromProto(mapCtx *direct.MapContext, in *pb.SparkRBatch) *krm.SparkRBatch {
	if in == nil {
		return nil
	}
	out := &krm.SparkRBatch{}
	out.MainRFileURI = direct.LazyPtr(in.GetMainRFileUri())
	out.Args = in.Args
	out.FileURIs = in.FileUris
	out.ArchiveURIs = in.ArchiveUris
	return out
}
func SparkRBatch_ToProto(mapCtx *direct.MapContext, in *krm.SparkRBatch) *pb.SparkRBatch {
	if in == nil {
		return nil
	}
	out := &pb.SparkRBatch{}
	out.MainRFileUri = direct.ValueOf(in.MainRFileURI)
	out.Args = in.Args
	out.FileUris = in.FileURIs
	out.ArchiveUris = in.ArchiveURIs
	return out
}
func SparkSQLBatch_FromProto(mapCtx *direct.MapContext, in *pb.SparkSqlBatch) *krm.SparkSQLBatch {
	if in == nil {
		return nil
	}
	out := &krm.SparkSQLBatch{}
	out.QueryFileURI = direct.LazyPtr(in.GetQueryFileUri())
	out.QueryVariables = in.QueryVariables
	out.JarFileURIs = in.JarFileUris
	return out
}
func SparkSQLBatch_ToProto(mapCtx *direct.MapContext, in *krm.SparkSQLBatch) *pb.SparkSqlBatch {
	if in == nil {
		return nil
	}
	out := &pb.SparkSqlBatch{}
	out.QueryFileUri = direct.ValueOf(in.QueryFileURI)
	out.QueryVariables = in.QueryVariables
	out.JarFileUris = in.JarFileURIs
	return out
}

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
func Batch_StateHistory_FromProto(mapCtx *direct.MapContext, in *pb.Batch_StateHistory) *krm.Batch_StateHistory {
	if in == nil {
		return nil
	}
	out := &krm.Batch_StateHistory{}
	// MISSING: State
	// MISSING: StateMessage
	// MISSING: StateStartTime
	return out
}
func Batch_StateHistory_ToProto(mapCtx *direct.MapContext, in *krm.Batch_StateHistory) *pb.Batch_StateHistory {
	if in == nil {
		return nil
	}
	out := &pb.Batch_StateHistory{}
	// MISSING: State
	// MISSING: StateMessage
	// MISSING: StateStartTime
	return out
}
func Batch_StateHistoryObservedState_FromProto(mapCtx *direct.MapContext, in *pb.Batch_StateHistory) *krm.Batch_StateHistoryObservedState {
	if in == nil {
		return nil
	}
	out := &krm.Batch_StateHistoryObservedState{}
	out.State = direct.Enum_FromProto(mapCtx, in.GetState())
	out.StateMessage = direct.LazyPtr(in.GetStateMessage())
	out.StateStartTime = direct.StringTimestamp_FromProto(mapCtx, in.GetStateStartTime())
	return out
}
func Batch_StateHistoryObservedState_ToProto(mapCtx *direct.MapContext, in *krm.Batch_StateHistoryObservedState) *pb.Batch_StateHistory {
	if in == nil {
		return nil
	}
	out := &pb.Batch_StateHistory{}
	out.State = direct.Enum_ToProto[pb.Batch_State](mapCtx, in.State)
	out.StateMessage = direct.ValueOf(in.StateMessage)
	out.StateStartTime = direct.StringTimestamp_ToProto(mapCtx, in.StateStartTime)
	return out
}
func DataprocBatchObservedState_FromProto(mapCtx *direct.MapContext, in *pb.Batch) *krm.DataprocBatchObservedState {
	if in == nil {
		return nil
	}
	out := &krm.DataprocBatchObservedState{}
	// MISSING: Name
	out.Uuid = direct.LazyPtr(in.GetUuid())
	out.CreateTime = direct.StringTimestamp_FromProto(mapCtx, in.GetCreateTime())
	out.RuntimeInfo = RuntimeInfoObservedState_FromProto(mapCtx, in.GetRuntimeInfo())
	out.State = direct.Enum_FromProto(mapCtx, in.GetState())
	out.StateMessage = direct.LazyPtr(in.GetStateMessage())
	out.StateTime = direct.StringTimestamp_FromProto(mapCtx, in.GetStateTime())
	out.Creator = direct.LazyPtr(in.GetCreator())
	out.Operation = direct.LazyPtr(in.GetOperation())
	out.StateHistory = direct.Slice_FromProto(mapCtx, in.StateHistory, Batch_StateHistoryObservedState_FromProto)
	return out
}
func DataprocBatchObservedState_ToProto(mapCtx *direct.MapContext, in *krm.DataprocBatchObservedState) *pb.Batch {
	if in == nil {
		return nil
	}
	out := &pb.Batch{}
	// MISSING: Name
	out.Uuid = direct.ValueOf(in.Uuid)
	out.CreateTime = direct.StringTimestamp_ToProto(mapCtx, in.CreateTime)
	out.RuntimeInfo = RuntimeInfoObservedState_ToProto(mapCtx, in.RuntimeInfo)
	out.State = direct.Enum_ToProto[pb.Batch_State](mapCtx, in.State)
	out.StateMessage = direct.ValueOf(in.StateMessage)
	out.StateTime = direct.StringTimestamp_ToProto(mapCtx, in.StateTime)
	out.Creator = direct.ValueOf(in.Creator)
	out.Operation = direct.ValueOf(in.Operation)
	out.StateHistory = direct.Slice_ToProto(mapCtx, in.StateHistory, Batch_StateHistoryObservedState_ToProto)
	return out
}
func DataprocBatchSpec_FromProto(mapCtx *direct.MapContext, in *pb.Batch) *krm.DataprocBatchSpec {
	if in == nil {
		return nil
	}
	out := &krm.DataprocBatchSpec{}
	// MISSING: Name
	out.PysparkBatch = PySparkBatch_FromProto(mapCtx, in.GetPysparkBatch())
	out.SparkBatch = SparkBatch_FromProto(mapCtx, in.GetSparkBatch())
	out.SparkRBatch = SparkRBatch_FromProto(mapCtx, in.GetSparkRBatch())
	out.SparkSQLBatch = SparkSQLBatch_FromProto(mapCtx, in.GetSparkSqlBatch())
	out.Labels = in.Labels
	out.RuntimeConfig = RuntimeConfig_FromProto(mapCtx, in.GetRuntimeConfig())
	out.EnvironmentConfig = EnvironmentConfig_FromProto(mapCtx, in.GetEnvironmentConfig())
	return out
}
func DataprocBatchSpec_ToProto(mapCtx *direct.MapContext, in *krm.DataprocBatchSpec) *pb.Batch {
	if in == nil {
		return nil
	}
	out := &pb.Batch{}
	// MISSING: Name
	if oneof := PySparkBatch_ToProto(mapCtx, in.PysparkBatch); oneof != nil {
		out.BatchConfig = &pb.Batch_PysparkBatch{PysparkBatch: oneof}
	}
	if oneof := SparkBatch_ToProto(mapCtx, in.SparkBatch); oneof != nil {
		out.BatchConfig = &pb.Batch_SparkBatch{SparkBatch: oneof}
	}
	if oneof := SparkRBatch_ToProto(mapCtx, in.SparkRBatch); oneof != nil {
		out.BatchConfig = &pb.Batch_SparkRBatch{SparkRBatch: oneof}
	}
	if oneof := SparkSQLBatch_ToProto(mapCtx, in.SparkSQLBatch); oneof != nil {
		out.BatchConfig = &pb.Batch_SparkSqlBatch{SparkSqlBatch: oneof}
	}
	out.Labels = in.Labels
	out.RuntimeConfig = RuntimeConfig_ToProto(mapCtx, in.RuntimeConfig)
	out.EnvironmentConfig = EnvironmentConfig_ToProto(mapCtx, in.EnvironmentConfig)
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
	out.ServiceAccountRef = &refs.IAMServiceAccountRef{
		External: in.GetServiceAccount(),
	}
	out.NetworkURI = direct.LazyPtr(in.GetNetworkUri())
	out.SubnetworkURI = direct.LazyPtr(in.GetSubnetworkUri())
	out.NetworkTags = in.NetworkTags
	out.KMSKeyRef = &refs.KMSCryptoKeyRef{
		External: in.GetKmsKey(),
	}
	out.IdleTTL = direct.StringDuration_FromProto(mapCtx, in.GetIdleTtl())
	out.TTL = direct.StringDuration_FromProto(mapCtx, in.GetTtl())
	out.StagingBucketRef = &storagev1beta1.StorageBucketRef{
		External: in.GetStagingBucket(),
	}
	return out
}
func ExecutionConfig_ToProto(mapCtx *direct.MapContext, in *krm.ExecutionConfig) *pb.ExecutionConfig {
	if in == nil {
		return nil
	}
	out := &pb.ExecutionConfig{}
	if in.ServiceAccountRef != nil {
		out.ServiceAccount = direct.ValueOf(&in.ServiceAccountRef.External)
	}
	if oneof := ExecutionConfig_NetworkUri_ToProto(mapCtx, in.NetworkURI); oneof != nil {
		out.Network = oneof
	}
	if oneof := ExecutionConfig_SubnetworkUri_ToProto(mapCtx, in.SubnetworkURI); oneof != nil {
		out.Network = oneof
	}
	out.NetworkTags = in.NetworkTags
	if in.KMSKeyRef != nil {
		out.KmsKey = direct.ValueOf(&in.KMSKeyRef.External)
	}
	out.IdleTtl = direct.StringDuration_ToProto(mapCtx, in.IdleTTL)
	out.Ttl = direct.StringDuration_ToProto(mapCtx, in.TTL)
	if in.StagingBucketRef != nil {
		out.StagingBucket = direct.ValueOf(&in.StagingBucketRef.External)
	}
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
func SparkHistoryServerConfig_FromProto(mapCtx *direct.MapContext, in *pb.SparkHistoryServerConfig) *krm.SparkHistoryServerConfig {
	if in == nil {
		return nil
	}
	out := &krm.SparkHistoryServerConfig{}
	out.DataprocClusterRef = &dataprocv1beta1.DataprocClusterRef{
		External: in.GetDataprocCluster(),
	}
	return out
}
func SparkHistoryServerConfig_ToProto(mapCtx *direct.MapContext, in *krm.SparkHistoryServerConfig) *pb.SparkHistoryServerConfig {
	if in == nil {
		return nil
	}
	out := &pb.SparkHistoryServerConfig{}
	if in.DataprocClusterRef != nil {
		out.DataprocCluster = direct.ValueOf(&in.DataprocClusterRef.External)
	}
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
