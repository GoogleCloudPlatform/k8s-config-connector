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
// krm.group: dataproc.cnrm.cloud.google.com
// krm.version: v1alpha1
// proto.service: google.cloud.dataproc.v1

package dataproc

import (
	pb "cloud.google.com/go/dataproc/v2/apiv1/dataprocpb"
	krm "github.com/GoogleCloudPlatform/k8s-config-connector/apis/dataproc/v1alpha1"
	krmdataprocv1beta1 "github.com/GoogleCloudPlatform/k8s-config-connector/apis/dataproc/v1beta1"
	refsv1beta1 "github.com/GoogleCloudPlatform/k8s-config-connector/apis/refs/v1beta1"
	krmstoragev1beta1 "github.com/GoogleCloudPlatform/k8s-config-connector/apis/storage/v1beta1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct"
)

func AcceleratorConfig_FromProto(mapCtx *direct.MapContext, in *pb.AcceleratorConfig) *krm.AcceleratorConfig {
	if in == nil {
		return nil
	}
	out := &krm.AcceleratorConfig{}
	out.AcceleratorTypeURI = direct.LazyPtr(in.GetAcceleratorTypeUri())
	out.AcceleratorCount = direct.LazyPtr(in.GetAcceleratorCount())
	return out
}
func AcceleratorConfig_ToProto(mapCtx *direct.MapContext, in *krm.AcceleratorConfig) *pb.AcceleratorConfig {
	if in == nil {
		return nil
	}
	out := &pb.AcceleratorConfig{}
	out.AcceleratorTypeUri = direct.ValueOf(in.AcceleratorTypeURI)
	out.AcceleratorCount = direct.ValueOf(in.AcceleratorCount)
	return out
}
func AuthenticationConfig_FromProto(mapCtx *direct.MapContext, in *pb.AuthenticationConfig) *krm.AuthenticationConfig {
	if in == nil {
		return nil
	}
	out := &krm.AuthenticationConfig{}
	out.UserWorkloadAuthenticationType = direct.Enum_FromProto(mapCtx, in.GetUserWorkloadAuthenticationType())
	return out
}
func AuthenticationConfig_ToProto(mapCtx *direct.MapContext, in *krm.AuthenticationConfig) *pb.AuthenticationConfig {
	if in == nil {
		return nil
	}
	out := &pb.AuthenticationConfig{}
	out.UserWorkloadAuthenticationType = direct.Enum_ToProto[pb.AuthenticationConfig_AuthenticationType](mapCtx, in.UserWorkloadAuthenticationType)
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
func DataprocJobObservedState_FromProto(mapCtx *direct.MapContext, in *pb.Job) *krm.DataprocJobObservedState {
	if in == nil {
		return nil
	}
	out := &krm.DataprocJobObservedState{}
	out.Placement = JobPlacementObservedState_FromProto(mapCtx, in.GetPlacement())
	out.Status = JobStatusObservedState_FromProto(mapCtx, in.GetStatus())
	out.StatusHistory = direct.Slice_FromProto(mapCtx, in.StatusHistory, JobStatusObservedState_FromProto)
	out.YarnApplications = direct.Slice_FromProto(mapCtx, in.YarnApplications, YarnApplication_FromProto)
	out.DriverOutputResourceURI = direct.LazyPtr(in.GetDriverOutputResourceUri())
	out.DriverControlFilesURI = direct.LazyPtr(in.GetDriverControlFilesUri())
	// MISSING: JobUuid
	// (near miss): "JobUuid" vs "JobUUid"
	out.Done = direct.LazyPtr(in.GetDone())
	return out
}
func DataprocJobObservedState_ToProto(mapCtx *direct.MapContext, in *krm.DataprocJobObservedState) *pb.Job {
	if in == nil {
		return nil
	}
	out := &pb.Job{}
	out.Placement = JobPlacementObservedState_ToProto(mapCtx, in.Placement)
	out.Status = JobStatusObservedState_ToProto(mapCtx, in.Status)
	out.StatusHistory = direct.Slice_ToProto(mapCtx, in.StatusHistory, JobStatusObservedState_ToProto)
	out.YarnApplications = direct.Slice_ToProto(mapCtx, in.YarnApplications, YarnApplication_ToProto)
	out.DriverOutputResourceUri = direct.ValueOf(in.DriverOutputResourceURI)
	out.DriverControlFilesUri = direct.ValueOf(in.DriverControlFilesURI)
	// MISSING: JobUuid
	// (near miss): "JobUuid" vs "JobUUid"
	out.Done = direct.ValueOf(in.Done)
	return out
}
func DataprocJobSpec_FromProto(mapCtx *direct.MapContext, in *pb.Job) *krm.DataprocJobSpec {
	if in == nil {
		return nil
	}
	out := &krm.DataprocJobSpec{}
	out.Reference = JobReference_FromProto(mapCtx, in.GetReference())
	out.Placement = JobPlacement_FromProto(mapCtx, in.GetPlacement())
	out.HadoopJob = HadoopJob_FromProto(mapCtx, in.GetHadoopJob())
	out.SparkJob = SparkJob_FromProto(mapCtx, in.GetSparkJob())
	out.PysparkJob = PySparkJob_FromProto(mapCtx, in.GetPysparkJob())
	out.HiveJob = HiveJob_FromProto(mapCtx, in.GetHiveJob())
	out.PigJob = PigJob_FromProto(mapCtx, in.GetPigJob())
	out.SparkRJob = SparkRJob_FromProto(mapCtx, in.GetSparkRJob())
	out.SparkSQLJob = SparkSQLJob_FromProto(mapCtx, in.GetSparkSqlJob())
	out.PrestoJob = PrestoJob_FromProto(mapCtx, in.GetPrestoJob())
	out.TrinoJob = TrinoJob_FromProto(mapCtx, in.GetTrinoJob())
	out.FlinkJob = FlinkJob_FromProto(mapCtx, in.GetFlinkJob())
	out.Labels = in.Labels
	out.Scheduling = JobScheduling_FromProto(mapCtx, in.GetScheduling())
	// MISSING: JobUuid
	out.DriverSchedulingConfig = DriverSchedulingConfig_FromProto(mapCtx, in.GetDriverSchedulingConfig())
	return out
}
func DataprocJobSpec_ToProto(mapCtx *direct.MapContext, in *krm.DataprocJobSpec) *pb.Job {
	if in == nil {
		return nil
	}
	out := &pb.Job{}
	out.Reference = JobReference_ToProto(mapCtx, in.Reference)
	out.Placement = JobPlacement_ToProto(mapCtx, in.Placement)
	if oneof := HadoopJob_ToProto(mapCtx, in.HadoopJob); oneof != nil {
		out.TypeJob = &pb.Job_HadoopJob{HadoopJob: oneof}
	}
	if oneof := SparkJob_ToProto(mapCtx, in.SparkJob); oneof != nil {
		out.TypeJob = &pb.Job_SparkJob{SparkJob: oneof}
	}
	if oneof := PySparkJob_ToProto(mapCtx, in.PysparkJob); oneof != nil {
		out.TypeJob = &pb.Job_PysparkJob{PysparkJob: oneof}
	}
	if oneof := HiveJob_ToProto(mapCtx, in.HiveJob); oneof != nil {
		out.TypeJob = &pb.Job_HiveJob{HiveJob: oneof}
	}
	if oneof := PigJob_ToProto(mapCtx, in.PigJob); oneof != nil {
		out.TypeJob = &pb.Job_PigJob{PigJob: oneof}
	}
	if oneof := SparkRJob_ToProto(mapCtx, in.SparkRJob); oneof != nil {
		out.TypeJob = &pb.Job_SparkRJob{SparkRJob: oneof}
	}
	if oneof := SparkSQLJob_ToProto(mapCtx, in.SparkSQLJob); oneof != nil {
		out.TypeJob = &pb.Job_SparkSqlJob{SparkSqlJob: oneof}
	}
	if oneof := PrestoJob_ToProto(mapCtx, in.PrestoJob); oneof != nil {
		out.TypeJob = &pb.Job_PrestoJob{PrestoJob: oneof}
	}
	if oneof := TrinoJob_ToProto(mapCtx, in.TrinoJob); oneof != nil {
		out.TypeJob = &pb.Job_TrinoJob{TrinoJob: oneof}
	}
	if oneof := FlinkJob_ToProto(mapCtx, in.FlinkJob); oneof != nil {
		out.TypeJob = &pb.Job_FlinkJob{FlinkJob: oneof}
	}
	out.Labels = in.Labels
	out.Scheduling = JobScheduling_ToProto(mapCtx, in.Scheduling)
	// MISSING: JobUuid
	out.DriverSchedulingConfig = DriverSchedulingConfig_ToProto(mapCtx, in.DriverSchedulingConfig)
	return out
}
func DataprocNodeGroupObservedState_FromProto(mapCtx *direct.MapContext, in *pb.NodeGroup) *krm.DataprocNodeGroupObservedState {
	if in == nil {
		return nil
	}
	out := &krm.DataprocNodeGroupObservedState{}
	// MISSING: Name
	out.NodeGroupConfig = InstanceGroupConfigObservedState_FromProto(mapCtx, in.GetNodeGroupConfig())
	return out
}
func DataprocNodeGroupObservedState_ToProto(mapCtx *direct.MapContext, in *krm.DataprocNodeGroupObservedState) *pb.NodeGroup {
	if in == nil {
		return nil
	}
	out := &pb.NodeGroup{}
	// MISSING: Name
	out.NodeGroupConfig = InstanceGroupConfigObservedState_ToProto(mapCtx, in.NodeGroupConfig)
	return out
}
func DataprocNodeGroupSpec_FromProto(mapCtx *direct.MapContext, in *pb.NodeGroup) *krm.DataprocNodeGroupSpec {
	if in == nil {
		return nil
	}
	out := &krm.DataprocNodeGroupSpec{}
	// MISSING: Name
	out.Roles = direct.EnumSlice_FromProto(mapCtx, in.Roles)
	out.NodeGroupConfig = InstanceGroupConfig_FromProto(mapCtx, in.GetNodeGroupConfig())
	out.Labels = in.Labels
	return out
}
func DataprocNodeGroupSpec_ToProto(mapCtx *direct.MapContext, in *krm.DataprocNodeGroupSpec) *pb.NodeGroup {
	if in == nil {
		return nil
	}
	out := &pb.NodeGroup{}
	// MISSING: Name
	out.Roles = direct.EnumSlice_ToProto[pb.NodeGroup_Role](mapCtx, in.Roles)
	out.NodeGroupConfig = InstanceGroupConfig_ToProto(mapCtx, in.NodeGroupConfig)
	out.Labels = in.Labels
	return out
}
func DiskConfig_FromProto(mapCtx *direct.MapContext, in *pb.DiskConfig) *krm.DiskConfig {
	if in == nil {
		return nil
	}
	out := &krm.DiskConfig{}
	out.BootDiskType = direct.LazyPtr(in.GetBootDiskType())
	out.BootDiskSizeGB = direct.LazyPtr(in.GetBootDiskSizeGb())
	// MISSING: NumLocalSsds
	// (near miss): "NumLocalSsds" vs "NumLocalSSDs"
	// MISSING: LocalSsdInterface
	// (near miss): "LocalSsdInterface" vs "LocalSSDInterface"
	// MISSING: BootDiskProvisionedIops
	// (near miss): "BootDiskProvisionedIops" vs "BootDiskProvisionedIOPs"
	out.BootDiskProvisionedThroughput = in.BootDiskProvisionedThroughput
	return out
}
func DiskConfig_ToProto(mapCtx *direct.MapContext, in *krm.DiskConfig) *pb.DiskConfig {
	if in == nil {
		return nil
	}
	out := &pb.DiskConfig{}
	out.BootDiskType = direct.ValueOf(in.BootDiskType)
	out.BootDiskSizeGb = direct.ValueOf(in.BootDiskSizeGB)
	// MISSING: NumLocalSsds
	// (near miss): "NumLocalSsds" vs "NumLocalSSDs"
	// MISSING: LocalSsdInterface
	// (near miss): "LocalSsdInterface" vs "LocalSSDInterface"
	// MISSING: BootDiskProvisionedIops
	// (near miss): "BootDiskProvisionedIops" vs "BootDiskProvisionedIOPs"
	out.BootDiskProvisionedThroughput = in.BootDiskProvisionedThroughput
	return out
}
func DriverSchedulingConfig_FromProto(mapCtx *direct.MapContext, in *pb.DriverSchedulingConfig) *krm.DriverSchedulingConfig {
	if in == nil {
		return nil
	}
	out := &krm.DriverSchedulingConfig{}
	out.MemoryMb = direct.LazyPtr(in.GetMemoryMb())
	out.Vcores = direct.LazyPtr(in.GetVcores())
	return out
}
func DriverSchedulingConfig_ToProto(mapCtx *direct.MapContext, in *krm.DriverSchedulingConfig) *pb.DriverSchedulingConfig {
	if in == nil {
		return nil
	}
	out := &pb.DriverSchedulingConfig{}
	out.MemoryMb = direct.ValueOf(in.MemoryMb)
	out.Vcores = direct.ValueOf(in.Vcores)
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
	if in.GetServiceAccount() != "" {
		out.ServiceAccountRef = &refsv1beta1.IAMServiceAccountRef{External: in.GetServiceAccount()}
	}
	out.NetworkURI = direct.LazyPtr(in.GetNetworkUri())
	out.SubnetworkURI = direct.LazyPtr(in.GetSubnetworkUri())
	out.NetworkTags = in.NetworkTags
	if in.GetKmsKey() != "" {
		out.KMSKeyRef = &refsv1beta1.KMSCryptoKeyRef{External: in.GetKmsKey()}
	}
	out.IdleTTL = direct.StringDuration_FromProto(mapCtx, in.GetIdleTtl())
	out.TTL = direct.StringDuration_FromProto(mapCtx, in.GetTtl())
	if in.GetStagingBucket() != "" {
		out.StagingBucketRef = &krmstoragev1beta1.StorageBucketRef{External: in.GetStagingBucket()}
	}
	// MISSING: AuthenticationConfig
	return out
}
func ExecutionConfig_ToProto(mapCtx *direct.MapContext, in *krm.ExecutionConfig) *pb.ExecutionConfig {
	if in == nil {
		return nil
	}
	out := &pb.ExecutionConfig{}
	if in.ServiceAccountRef != nil {
		out.ServiceAccount = in.ServiceAccountRef.External
	}
	if oneof := ExecutionConfig_NetworkUri_ToProto(mapCtx, in.NetworkURI); oneof != nil {
		out.Network = oneof
	}
	if oneof := ExecutionConfig_SubnetworkUri_ToProto(mapCtx, in.SubnetworkURI); oneof != nil {
		out.Network = oneof
	}
	out.NetworkTags = in.NetworkTags
	if in.KMSKeyRef != nil {
		out.KmsKey = in.KMSKeyRef.External
	}
	out.IdleTtl = direct.StringDuration_ToProto(mapCtx, in.IdleTTL)
	out.Ttl = direct.StringDuration_ToProto(mapCtx, in.TTL)
	if in.StagingBucketRef != nil {
		out.StagingBucket = in.StagingBucketRef.External
	}
	// MISSING: AuthenticationConfig
	return out
}
func ExecutionConfig_NetworkUri_ToProto(mapCtx *direct.MapContext, in *string) *pb.ExecutionConfig_NetworkUri {
	if in == nil {
		return nil
	}
	return &pb.ExecutionConfig_NetworkUri{NetworkUri: *in}
}
func ExecutionConfig_SubnetworkUri_ToProto(mapCtx *direct.MapContext, in *string) *pb.ExecutionConfig_SubnetworkUri {
	if in == nil {
		return nil
	}
	return &pb.ExecutionConfig_SubnetworkUri{SubnetworkUri: *in}
}
func FlinkJob_FromProto(mapCtx *direct.MapContext, in *pb.FlinkJob) *krm.FlinkJob {
	if in == nil {
		return nil
	}
	out := &krm.FlinkJob{}
	out.MainJarFileURI = direct.LazyPtr(in.GetMainJarFileUri())
	out.MainClass = direct.LazyPtr(in.GetMainClass())
	out.Args = in.Args
	// MISSING: JarFileUris
	// (near miss): "JarFileUris" vs "JarFileURIs"
	out.SavepointURI = direct.LazyPtr(in.GetSavepointUri())
	out.Properties = in.Properties
	out.LoggingConfig = LoggingConfig_FromProto(mapCtx, in.GetLoggingConfig())
	return out
}
func FlinkJob_ToProto(mapCtx *direct.MapContext, in *krm.FlinkJob) *pb.FlinkJob {
	if in == nil {
		return nil
	}
	out := &pb.FlinkJob{}
	if oneof := FlinkJob_MainJarFileUri_ToProto(mapCtx, in.MainJarFileURI); oneof != nil {
		out.Driver = oneof
	}
	if oneof := FlinkJob_MainClass_ToProto(mapCtx, in.MainClass); oneof != nil {
		out.Driver = oneof
	}
	out.Args = in.Args
	// MISSING: JarFileUris
	// (near miss): "JarFileUris" vs "JarFileURIs"
	out.SavepointUri = direct.ValueOf(in.SavepointURI)
	out.Properties = in.Properties
	out.LoggingConfig = LoggingConfig_ToProto(mapCtx, in.LoggingConfig)
	return out
}
func FlinkJob_MainJarFileUri_ToProto(mapCtx *direct.MapContext, in *string) *pb.FlinkJob_MainJarFileUri {
	if in == nil {
		return nil
	}
	return &pb.FlinkJob_MainJarFileUri{MainJarFileUri: *in}
}
func FlinkJob_MainClass_ToProto(mapCtx *direct.MapContext, in *string) *pb.FlinkJob_MainClass {
	if in == nil {
		return nil
	}
	return &pb.FlinkJob_MainClass{MainClass: *in}
}
func HadoopJob_FromProto(mapCtx *direct.MapContext, in *pb.HadoopJob) *krm.HadoopJob {
	if in == nil {
		return nil
	}
	out := &krm.HadoopJob{}
	out.MainJarFileURI = direct.LazyPtr(in.GetMainJarFileUri())
	out.MainClass = direct.LazyPtr(in.GetMainClass())
	out.Args = in.Args
	// MISSING: JarFileUris
	// (near miss): "JarFileUris" vs "JarFileURIs"
	// MISSING: FileUris
	// (near miss): "FileUris" vs "FileURIs"
	// MISSING: ArchiveUris
	// (near miss): "ArchiveUris" vs "ArchiveURIs"
	out.Properties = in.Properties
	out.LoggingConfig = LoggingConfig_FromProto(mapCtx, in.GetLoggingConfig())
	return out
}
func HadoopJob_ToProto(mapCtx *direct.MapContext, in *krm.HadoopJob) *pb.HadoopJob {
	if in == nil {
		return nil
	}
	out := &pb.HadoopJob{}
	if oneof := HadoopJob_MainJarFileUri_ToProto(mapCtx, in.MainJarFileURI); oneof != nil {
		out.Driver = oneof
	}
	if oneof := HadoopJob_MainClass_ToProto(mapCtx, in.MainClass); oneof != nil {
		out.Driver = oneof
	}
	out.Args = in.Args
	// MISSING: JarFileUris
	// (near miss): "JarFileUris" vs "JarFileURIs"
	// MISSING: FileUris
	// (near miss): "FileUris" vs "FileURIs"
	// MISSING: ArchiveUris
	// (near miss): "ArchiveUris" vs "ArchiveURIs"
	out.Properties = in.Properties
	out.LoggingConfig = LoggingConfig_ToProto(mapCtx, in.LoggingConfig)
	return out
}
func HadoopJob_MainJarFileUri_ToProto(mapCtx *direct.MapContext, in *string) *pb.HadoopJob_MainJarFileUri {
	if in == nil {
		return nil
	}
	return &pb.HadoopJob_MainJarFileUri{MainJarFileUri: *in}
}
func HadoopJob_MainClass_ToProto(mapCtx *direct.MapContext, in *string) *pb.HadoopJob_MainClass {
	if in == nil {
		return nil
	}
	return &pb.HadoopJob_MainClass{MainClass: *in}
}
func HiveJob_FromProto(mapCtx *direct.MapContext, in *pb.HiveJob) *krm.HiveJob {
	if in == nil {
		return nil
	}
	out := &krm.HiveJob{}
	out.QueryFileURI = direct.LazyPtr(in.GetQueryFileUri())
	out.QueryList = QueryList_FromProto(mapCtx, in.GetQueryList())
	out.ContinueOnFailure = direct.LazyPtr(in.GetContinueOnFailure())
	out.ScriptVariables = in.ScriptVariables
	out.Properties = in.Properties
	// MISSING: JarFileUris
	// (near miss): "JarFileUris" vs "JarFileURIs"
	return out
}
func HiveJob_ToProto(mapCtx *direct.MapContext, in *krm.HiveJob) *pb.HiveJob {
	if in == nil {
		return nil
	}
	out := &pb.HiveJob{}
	if oneof := HiveJob_QueryFileUri_ToProto(mapCtx, in.QueryFileURI); oneof != nil {
		out.Queries = oneof
	}
	if oneof := QueryList_ToProto(mapCtx, in.QueryList); oneof != nil {
		out.Queries = &pb.HiveJob_QueryList{QueryList: oneof}
	}
	out.ContinueOnFailure = direct.ValueOf(in.ContinueOnFailure)
	out.ScriptVariables = in.ScriptVariables
	out.Properties = in.Properties
	// MISSING: JarFileUris
	// (near miss): "JarFileUris" vs "JarFileURIs"
	return out
}
func HiveJob_QueryFileUri_ToProto(mapCtx *direct.MapContext, in *string) *pb.HiveJob_QueryFileUri {
	if in == nil {
		return nil
	}
	return &pb.HiveJob_QueryFileUri{QueryFileUri: *in}
}
func InstanceFlexibilityPolicy_FromProto(mapCtx *direct.MapContext, in *pb.InstanceFlexibilityPolicy) *krm.InstanceFlexibilityPolicy {
	if in == nil {
		return nil
	}
	out := &krm.InstanceFlexibilityPolicy{}
	out.ProvisioningModelMix = InstanceFlexibilityPolicy_ProvisioningModelMix_FromProto(mapCtx, in.GetProvisioningModelMix())
	out.InstanceSelectionList = direct.Slice_FromProto(mapCtx, in.InstanceSelectionList, InstanceFlexibilityPolicy_InstanceSelection_FromProto)
	// MISSING: InstanceSelectionResults
	return out
}
func InstanceFlexibilityPolicy_ToProto(mapCtx *direct.MapContext, in *krm.InstanceFlexibilityPolicy) *pb.InstanceFlexibilityPolicy {
	if in == nil {
		return nil
	}
	out := &pb.InstanceFlexibilityPolicy{}
	out.ProvisioningModelMix = InstanceFlexibilityPolicy_ProvisioningModelMix_ToProto(mapCtx, in.ProvisioningModelMix)
	out.InstanceSelectionList = direct.Slice_ToProto(mapCtx, in.InstanceSelectionList, InstanceFlexibilityPolicy_InstanceSelection_ToProto)
	// MISSING: InstanceSelectionResults
	return out
}
func InstanceFlexibilityPolicyObservedState_FromProto(mapCtx *direct.MapContext, in *pb.InstanceFlexibilityPolicy) *krm.InstanceFlexibilityPolicyObservedState {
	if in == nil {
		return nil
	}
	out := &krm.InstanceFlexibilityPolicyObservedState{}
	// MISSING: ProvisioningModelMix
	// MISSING: InstanceSelectionList
	out.InstanceSelectionResults = direct.Slice_FromProto(mapCtx, in.InstanceSelectionResults, InstanceFlexibilityPolicy_InstanceSelectionResult_FromProto)
	return out
}
func InstanceFlexibilityPolicyObservedState_ToProto(mapCtx *direct.MapContext, in *krm.InstanceFlexibilityPolicyObservedState) *pb.InstanceFlexibilityPolicy {
	if in == nil {
		return nil
	}
	out := &pb.InstanceFlexibilityPolicy{}
	// MISSING: ProvisioningModelMix
	// MISSING: InstanceSelectionList
	out.InstanceSelectionResults = direct.Slice_ToProto(mapCtx, in.InstanceSelectionResults, InstanceFlexibilityPolicy_InstanceSelectionResult_ToProto)
	return out
}
func InstanceFlexibilityPolicy_InstanceSelection_FromProto(mapCtx *direct.MapContext, in *pb.InstanceFlexibilityPolicy_InstanceSelection) *krm.InstanceFlexibilityPolicy_InstanceSelection {
	if in == nil {
		return nil
	}
	out := &krm.InstanceFlexibilityPolicy_InstanceSelection{}
	out.MachineTypes = in.MachineTypes
	out.Rank = direct.LazyPtr(in.GetRank())
	return out
}
func InstanceFlexibilityPolicy_InstanceSelection_ToProto(mapCtx *direct.MapContext, in *krm.InstanceFlexibilityPolicy_InstanceSelection) *pb.InstanceFlexibilityPolicy_InstanceSelection {
	if in == nil {
		return nil
	}
	out := &pb.InstanceFlexibilityPolicy_InstanceSelection{}
	out.MachineTypes = in.MachineTypes
	out.Rank = direct.ValueOf(in.Rank)
	return out
}
func InstanceFlexibilityPolicy_InstanceSelectionResult_FromProto(mapCtx *direct.MapContext, in *pb.InstanceFlexibilityPolicy_InstanceSelectionResult) *krm.InstanceFlexibilityPolicy_InstanceSelectionResult {
	if in == nil {
		return nil
	}
	out := &krm.InstanceFlexibilityPolicy_InstanceSelectionResult{}
	// MISSING: MachineType
	// MISSING: VMCount
	return out
}
func InstanceFlexibilityPolicy_InstanceSelectionResult_ToProto(mapCtx *direct.MapContext, in *krm.InstanceFlexibilityPolicy_InstanceSelectionResult) *pb.InstanceFlexibilityPolicy_InstanceSelectionResult {
	if in == nil {
		return nil
	}
	out := &pb.InstanceFlexibilityPolicy_InstanceSelectionResult{}
	// MISSING: MachineType
	// MISSING: VMCount
	return out
}
func InstanceFlexibilityPolicy_InstanceSelectionResultObservedState_FromProto(mapCtx *direct.MapContext, in *pb.InstanceFlexibilityPolicy_InstanceSelectionResult) *krm.InstanceFlexibilityPolicy_InstanceSelectionResultObservedState {
	if in == nil {
		return nil
	}
	out := &krm.InstanceFlexibilityPolicy_InstanceSelectionResultObservedState{}
	out.MachineType = in.MachineType
	out.VMCount = in.VmCount
	return out
}
func InstanceFlexibilityPolicy_InstanceSelectionResultObservedState_ToProto(mapCtx *direct.MapContext, in *krm.InstanceFlexibilityPolicy_InstanceSelectionResultObservedState) *pb.InstanceFlexibilityPolicy_InstanceSelectionResult {
	if in == nil {
		return nil
	}
	out := &pb.InstanceFlexibilityPolicy_InstanceSelectionResult{}
	out.MachineType = in.MachineType
	out.VmCount = in.VMCount
	return out
}
func InstanceFlexibilityPolicy_ProvisioningModelMix_FromProto(mapCtx *direct.MapContext, in *pb.InstanceFlexibilityPolicy_ProvisioningModelMix) *krm.InstanceFlexibilityPolicy_ProvisioningModelMix {
	if in == nil {
		return nil
	}
	out := &krm.InstanceFlexibilityPolicy_ProvisioningModelMix{}
	out.StandardCapacityBase = in.StandardCapacityBase
	out.StandardCapacityPercentAboveBase = in.StandardCapacityPercentAboveBase
	return out
}
func InstanceFlexibilityPolicy_ProvisioningModelMix_ToProto(mapCtx *direct.MapContext, in *krm.InstanceFlexibilityPolicy_ProvisioningModelMix) *pb.InstanceFlexibilityPolicy_ProvisioningModelMix {
	if in == nil {
		return nil
	}
	out := &pb.InstanceFlexibilityPolicy_ProvisioningModelMix{}
	out.StandardCapacityBase = in.StandardCapacityBase
	out.StandardCapacityPercentAboveBase = in.StandardCapacityPercentAboveBase
	return out
}
func InstanceGroupConfig_FromProto(mapCtx *direct.MapContext, in *pb.InstanceGroupConfig) *krm.InstanceGroupConfig {
	if in == nil {
		return nil
	}
	out := &krm.InstanceGroupConfig{}
	out.NumInstances = direct.LazyPtr(in.GetNumInstances())
	// MISSING: InstanceNames
	// MISSING: InstanceReferences
	out.ImageURI = direct.LazyPtr(in.GetImageUri())
	out.MachineTypeURI = direct.LazyPtr(in.GetMachineTypeUri())
	out.DiskConfig = DiskConfig_FromProto(mapCtx, in.GetDiskConfig())
	// MISSING: IsPreemptible
	out.Preemptibility = direct.Enum_FromProto(mapCtx, in.GetPreemptibility())
	// MISSING: ManagedGroupConfig
	out.Accelerators = direct.Slice_FromProto(mapCtx, in.Accelerators, AcceleratorConfig_FromProto)
	out.MinCPUPlatform = direct.LazyPtr(in.GetMinCpuPlatform())
	out.MinNumInstances = direct.LazyPtr(in.GetMinNumInstances())
	out.InstanceFlexibilityPolicy = InstanceFlexibilityPolicy_FromProto(mapCtx, in.GetInstanceFlexibilityPolicy())
	out.StartupConfig = StartupConfig_FromProto(mapCtx, in.GetStartupConfig())
	return out
}
func InstanceGroupConfig_ToProto(mapCtx *direct.MapContext, in *krm.InstanceGroupConfig) *pb.InstanceGroupConfig {
	if in == nil {
		return nil
	}
	out := &pb.InstanceGroupConfig{}
	out.NumInstances = direct.ValueOf(in.NumInstances)
	// MISSING: InstanceNames
	// MISSING: InstanceReferences
	out.ImageUri = direct.ValueOf(in.ImageURI)
	out.MachineTypeUri = direct.ValueOf(in.MachineTypeURI)
	out.DiskConfig = DiskConfig_ToProto(mapCtx, in.DiskConfig)
	// MISSING: IsPreemptible
	out.Preemptibility = direct.Enum_ToProto[pb.InstanceGroupConfig_Preemptibility](mapCtx, in.Preemptibility)
	// MISSING: ManagedGroupConfig
	out.Accelerators = direct.Slice_ToProto(mapCtx, in.Accelerators, AcceleratorConfig_ToProto)
	out.MinCpuPlatform = direct.ValueOf(in.MinCPUPlatform)
	out.MinNumInstances = direct.ValueOf(in.MinNumInstances)
	out.InstanceFlexibilityPolicy = InstanceFlexibilityPolicy_ToProto(mapCtx, in.InstanceFlexibilityPolicy)
	out.StartupConfig = StartupConfig_ToProto(mapCtx, in.StartupConfig)
	return out
}
func InstanceGroupConfigObservedState_FromProto(mapCtx *direct.MapContext, in *pb.InstanceGroupConfig) *krm.InstanceGroupConfigObservedState {
	if in == nil {
		return nil
	}
	out := &krm.InstanceGroupConfigObservedState{}
	// MISSING: NumInstances
	out.InstanceNames = in.InstanceNames
	out.InstanceReferences = direct.Slice_FromProto(mapCtx, in.InstanceReferences, InstanceReference_FromProto)
	// MISSING: ImageURI
	// MISSING: MachineTypeURI
	// MISSING: DiskConfig
	out.IsPreemptible = direct.LazyPtr(in.GetIsPreemptible())
	// MISSING: Preemptibility
	out.ManagedGroupConfig = ManagedGroupConfig_FromProto(mapCtx, in.GetManagedGroupConfig())
	// MISSING: Accelerators
	// MISSING: MinCPUPlatform
	// MISSING: MinNumInstances
	out.InstanceFlexibilityPolicy = InstanceFlexibilityPolicyObservedState_FromProto(mapCtx, in.GetInstanceFlexibilityPolicy())
	// MISSING: StartupConfig
	return out
}
func InstanceGroupConfigObservedState_ToProto(mapCtx *direct.MapContext, in *krm.InstanceGroupConfigObservedState) *pb.InstanceGroupConfig {
	if in == nil {
		return nil
	}
	out := &pb.InstanceGroupConfig{}
	// MISSING: NumInstances
	out.InstanceNames = in.InstanceNames
	out.InstanceReferences = direct.Slice_ToProto(mapCtx, in.InstanceReferences, InstanceReference_ToProto)
	// MISSING: ImageURI
	// MISSING: MachineTypeURI
	// MISSING: DiskConfig
	out.IsPreemptible = direct.ValueOf(in.IsPreemptible)
	// MISSING: Preemptibility
	out.ManagedGroupConfig = ManagedGroupConfig_ToProto(mapCtx, in.ManagedGroupConfig)
	// MISSING: Accelerators
	// MISSING: MinCPUPlatform
	// MISSING: MinNumInstances
	out.InstanceFlexibilityPolicy = InstanceFlexibilityPolicyObservedState_ToProto(mapCtx, in.InstanceFlexibilityPolicy)
	// MISSING: StartupConfig
	return out
}
func InstanceReference_FromProto(mapCtx *direct.MapContext, in *pb.InstanceReference) *krm.InstanceReference {
	if in == nil {
		return nil
	}
	out := &krm.InstanceReference{}
	out.InstanceName = direct.LazyPtr(in.GetInstanceName())
	out.InstanceID = direct.LazyPtr(in.GetInstanceId())
	out.PublicKey = direct.LazyPtr(in.GetPublicKey())
	out.PublicEciesKey = direct.LazyPtr(in.GetPublicEciesKey())
	return out
}
func InstanceReference_ToProto(mapCtx *direct.MapContext, in *krm.InstanceReference) *pb.InstanceReference {
	if in == nil {
		return nil
	}
	out := &pb.InstanceReference{}
	out.InstanceName = direct.ValueOf(in.InstanceName)
	out.InstanceId = direct.ValueOf(in.InstanceID)
	out.PublicKey = direct.ValueOf(in.PublicKey)
	out.PublicEciesKey = direct.ValueOf(in.PublicEciesKey)
	return out
}
func JobPlacement_FromProto(mapCtx *direct.MapContext, in *pb.JobPlacement) *krm.JobPlacement {
	if in == nil {
		return nil
	}
	out := &krm.JobPlacement{}
	out.ClusterName = direct.LazyPtr(in.GetClusterName())
	// MISSING: ClusterUuid
	out.ClusterLabels = in.ClusterLabels
	return out
}
func JobPlacement_ToProto(mapCtx *direct.MapContext, in *krm.JobPlacement) *pb.JobPlacement {
	if in == nil {
		return nil
	}
	out := &pb.JobPlacement{}
	out.ClusterName = direct.ValueOf(in.ClusterName)
	// MISSING: ClusterUuid
	out.ClusterLabels = in.ClusterLabels
	return out
}
func JobPlacementObservedState_FromProto(mapCtx *direct.MapContext, in *pb.JobPlacement) *krm.JobPlacementObservedState {
	if in == nil {
		return nil
	}
	out := &krm.JobPlacementObservedState{}
	// MISSING: ClusterName
	out.ClusterUuid = direct.LazyPtr(in.GetClusterUuid())
	// MISSING: ClusterLabels
	return out
}
func JobPlacementObservedState_ToProto(mapCtx *direct.MapContext, in *krm.JobPlacementObservedState) *pb.JobPlacement {
	if in == nil {
		return nil
	}
	out := &pb.JobPlacement{}
	// MISSING: ClusterName
	out.ClusterUuid = direct.ValueOf(in.ClusterUuid)
	// MISSING: ClusterLabels
	return out
}
func JobReference_FromProto(mapCtx *direct.MapContext, in *pb.JobReference) *krm.JobReference {
	if in == nil {
		return nil
	}
	out := &krm.JobReference{}
	out.ProjectID = direct.LazyPtr(in.GetProjectId())
	out.JobID = direct.LazyPtr(in.GetJobId())
	return out
}
func JobReference_ToProto(mapCtx *direct.MapContext, in *krm.JobReference) *pb.JobReference {
	if in == nil {
		return nil
	}
	out := &pb.JobReference{}
	out.ProjectId = direct.ValueOf(in.ProjectID)
	out.JobId = direct.ValueOf(in.JobID)
	return out
}
func JobScheduling_FromProto(mapCtx *direct.MapContext, in *pb.JobScheduling) *krm.JobScheduling {
	if in == nil {
		return nil
	}
	out := &krm.JobScheduling{}
	out.MaxFailuresPerHour = direct.LazyPtr(in.GetMaxFailuresPerHour())
	out.MaxFailuresTotal = direct.LazyPtr(in.GetMaxFailuresTotal())
	return out
}
func JobScheduling_ToProto(mapCtx *direct.MapContext, in *krm.JobScheduling) *pb.JobScheduling {
	if in == nil {
		return nil
	}
	out := &pb.JobScheduling{}
	out.MaxFailuresPerHour = direct.ValueOf(in.MaxFailuresPerHour)
	out.MaxFailuresTotal = direct.ValueOf(in.MaxFailuresTotal)
	return out
}
func JobStatus_FromProto(mapCtx *direct.MapContext, in *pb.JobStatus) *krm.JobStatus {
	if in == nil {
		return nil
	}
	out := &krm.JobStatus{}
	// MISSING: State
	// MISSING: Details
	// MISSING: StateStartTime
	// MISSING: Substate
	return out
}
func JobStatus_ToProto(mapCtx *direct.MapContext, in *krm.JobStatus) *pb.JobStatus {
	if in == nil {
		return nil
	}
	out := &pb.JobStatus{}
	// MISSING: State
	// MISSING: Details
	// MISSING: StateStartTime
	// MISSING: Substate
	return out
}
func JobStatusObservedState_FromProto(mapCtx *direct.MapContext, in *pb.JobStatus) *krm.JobStatusObservedState {
	if in == nil {
		return nil
	}
	out := &krm.JobStatusObservedState{}
	out.State = direct.Enum_FromProto(mapCtx, in.GetState())
	out.Details = direct.LazyPtr(in.GetDetails())
	out.StateStartTime = direct.StringTimestamp_FromProto(mapCtx, in.GetStateStartTime())
	out.Substate = direct.Enum_FromProto(mapCtx, in.GetSubstate())
	return out
}
func JobStatusObservedState_ToProto(mapCtx *direct.MapContext, in *krm.JobStatusObservedState) *pb.JobStatus {
	if in == nil {
		return nil
	}
	out := &pb.JobStatus{}
	out.State = direct.Enum_ToProto[pb.JobStatus_State](mapCtx, in.State)
	out.Details = direct.ValueOf(in.Details)
	out.StateStartTime = direct.StringTimestamp_ToProto(mapCtx, in.StateStartTime)
	out.Substate = direct.Enum_ToProto[pb.JobStatus_Substate](mapCtx, in.Substate)
	return out
}
func LoggingConfig_FromProto(mapCtx *direct.MapContext, in *pb.LoggingConfig) *krm.LoggingConfig {
	if in == nil {
		return nil
	}
	out := &krm.LoggingConfig{}
	// MISSING: DriverLogLevels
	return out
}
func LoggingConfig_ToProto(mapCtx *direct.MapContext, in *krm.LoggingConfig) *pb.LoggingConfig {
	if in == nil {
		return nil
	}
	out := &pb.LoggingConfig{}
	// MISSING: DriverLogLevels
	return out
}
func ManagedGroupConfig_FromProto(mapCtx *direct.MapContext, in *pb.ManagedGroupConfig) *krm.ManagedGroupConfig {
	if in == nil {
		return nil
	}
	out := &krm.ManagedGroupConfig{}
	// MISSING: InstanceTemplateName
	// MISSING: InstanceGroupManagerName
	// MISSING: InstanceGroupManagerURI
	return out
}
func ManagedGroupConfig_ToProto(mapCtx *direct.MapContext, in *krm.ManagedGroupConfig) *pb.ManagedGroupConfig {
	if in == nil {
		return nil
	}
	out := &pb.ManagedGroupConfig{}
	// MISSING: InstanceTemplateName
	// MISSING: InstanceGroupManagerName
	// MISSING: InstanceGroupManagerURI
	return out
}
func ManagedGroupConfigObservedState_FromProto(mapCtx *direct.MapContext, in *pb.ManagedGroupConfig) *krm.ManagedGroupConfigObservedState {
	if in == nil {
		return nil
	}
	out := &krm.ManagedGroupConfigObservedState{}
	out.InstanceTemplateName = direct.LazyPtr(in.GetInstanceTemplateName())
	out.InstanceGroupManagerName = direct.LazyPtr(in.GetInstanceGroupManagerName())
	out.InstanceGroupManagerURI = direct.LazyPtr(in.GetInstanceGroupManagerUri())
	return out
}
func ManagedGroupConfigObservedState_ToProto(mapCtx *direct.MapContext, in *krm.ManagedGroupConfigObservedState) *pb.ManagedGroupConfig {
	if in == nil {
		return nil
	}
	out := &pb.ManagedGroupConfig{}
	out.InstanceTemplateName = direct.ValueOf(in.InstanceTemplateName)
	out.InstanceGroupManagerName = direct.ValueOf(in.InstanceGroupManagerName)
	out.InstanceGroupManagerUri = direct.ValueOf(in.InstanceGroupManagerURI)
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
func PigJob_FromProto(mapCtx *direct.MapContext, in *pb.PigJob) *krm.PigJob {
	if in == nil {
		return nil
	}
	out := &krm.PigJob{}
	out.QueryFileURI = direct.LazyPtr(in.GetQueryFileUri())
	out.QueryList = QueryList_FromProto(mapCtx, in.GetQueryList())
	out.ContinueOnFailure = direct.LazyPtr(in.GetContinueOnFailure())
	out.ScriptVariables = in.ScriptVariables
	out.Properties = in.Properties
	// MISSING: JarFileUris
	// (near miss): "JarFileUris" vs "JarFileURIs"
	out.LoggingConfig = LoggingConfig_FromProto(mapCtx, in.GetLoggingConfig())
	return out
}
func PigJob_ToProto(mapCtx *direct.MapContext, in *krm.PigJob) *pb.PigJob {
	if in == nil {
		return nil
	}
	out := &pb.PigJob{}
	if oneof := PigJob_QueryFileUri_ToProto(mapCtx, in.QueryFileURI); oneof != nil {
		out.Queries = oneof
	}
	if oneof := QueryList_ToProto(mapCtx, in.QueryList); oneof != nil {
		out.Queries = &pb.PigJob_QueryList{QueryList: oneof}
	}
	out.ContinueOnFailure = direct.ValueOf(in.ContinueOnFailure)
	out.ScriptVariables = in.ScriptVariables
	out.Properties = in.Properties
	// MISSING: JarFileUris
	// (near miss): "JarFileUris" vs "JarFileURIs"
	out.LoggingConfig = LoggingConfig_ToProto(mapCtx, in.LoggingConfig)
	return out
}
func PigJob_QueryFileUri_ToProto(mapCtx *direct.MapContext, in *string) *pb.PigJob_QueryFileUri {
	if in == nil {
		return nil
	}
	return &pb.PigJob_QueryFileUri{QueryFileUri: *in}
}
func PrestoJob_FromProto(mapCtx *direct.MapContext, in *pb.PrestoJob) *krm.PrestoJob {
	if in == nil {
		return nil
	}
	out := &krm.PrestoJob{}
	out.QueryFileURI = direct.LazyPtr(in.GetQueryFileUri())
	out.QueryList = QueryList_FromProto(mapCtx, in.GetQueryList())
	out.ContinueOnFailure = direct.LazyPtr(in.GetContinueOnFailure())
	out.OutputFormat = direct.LazyPtr(in.GetOutputFormat())
	out.ClientTags = in.ClientTags
	out.Properties = in.Properties
	out.LoggingConfig = LoggingConfig_FromProto(mapCtx, in.GetLoggingConfig())
	return out
}
func PrestoJob_ToProto(mapCtx *direct.MapContext, in *krm.PrestoJob) *pb.PrestoJob {
	if in == nil {
		return nil
	}
	out := &pb.PrestoJob{}
	if oneof := PrestoJob_QueryFileUri_ToProto(mapCtx, in.QueryFileURI); oneof != nil {
		out.Queries = oneof
	}
	if oneof := QueryList_ToProto(mapCtx, in.QueryList); oneof != nil {
		out.Queries = &pb.PrestoJob_QueryList{QueryList: oneof}
	}
	out.ContinueOnFailure = direct.ValueOf(in.ContinueOnFailure)
	out.OutputFormat = direct.ValueOf(in.OutputFormat)
	out.ClientTags = in.ClientTags
	out.Properties = in.Properties
	out.LoggingConfig = LoggingConfig_ToProto(mapCtx, in.LoggingConfig)
	return out
}
func PrestoJob_QueryFileUri_ToProto(mapCtx *direct.MapContext, in *string) *pb.PrestoJob_QueryFileUri {
	if in == nil {
		return nil
	}
	return &pb.PrestoJob_QueryFileUri{QueryFileUri: *in}
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
func PySparkBatch_FromProto(mapCtx *direct.MapContext, in *pb.PySparkBatch) *krm.PySparkBatch {
	if in == nil {
		return nil
	}
	out := &krm.PySparkBatch{}
	out.MainPythonFileURI = direct.LazyPtr(in.GetMainPythonFileUri())
	out.Args = in.Args
	// MISSING: PythonFileUris
	// (near miss): "PythonFileUris" vs "PythonFileURIs"
	// MISSING: JarFileUris
	// (near miss): "JarFileUris" vs "JarFileURIs"
	// MISSING: FileUris
	// (near miss): "FileUris" vs "FileURIs"
	// MISSING: ArchiveUris
	// (near miss): "ArchiveUris" vs "ArchiveURIs"
	return out
}
func PySparkBatch_ToProto(mapCtx *direct.MapContext, in *krm.PySparkBatch) *pb.PySparkBatch {
	if in == nil {
		return nil
	}
	out := &pb.PySparkBatch{}
	out.MainPythonFileUri = direct.ValueOf(in.MainPythonFileURI)
	out.Args = in.Args
	// MISSING: PythonFileUris
	// (near miss): "PythonFileUris" vs "PythonFileURIs"
	// MISSING: JarFileUris
	// (near miss): "JarFileUris" vs "JarFileURIs"
	// MISSING: FileUris
	// (near miss): "FileUris" vs "FileURIs"
	// MISSING: ArchiveUris
	// (near miss): "ArchiveUris" vs "ArchiveURIs"
	return out
}
func PySparkJob_FromProto(mapCtx *direct.MapContext, in *pb.PySparkJob) *krm.PySparkJob {
	if in == nil {
		return nil
	}
	out := &krm.PySparkJob{}
	out.MainPythonFileURI = direct.LazyPtr(in.GetMainPythonFileUri())
	out.Args = in.Args
	// MISSING: PythonFileUris
	// (near miss): "PythonFileUris" vs "PythonFileURIs"
	// MISSING: JarFileUris
	// (near miss): "JarFileUris" vs "JarFileURIs"
	// MISSING: FileUris
	// (near miss): "FileUris" vs "FileURIs"
	// MISSING: ArchiveUris
	// (near miss): "ArchiveUris" vs "ArchiveURIs"
	out.Properties = in.Properties
	out.LoggingConfig = LoggingConfig_FromProto(mapCtx, in.GetLoggingConfig())
	return out
}
func PySparkJob_ToProto(mapCtx *direct.MapContext, in *krm.PySparkJob) *pb.PySparkJob {
	if in == nil {
		return nil
	}
	out := &pb.PySparkJob{}
	out.MainPythonFileUri = direct.ValueOf(in.MainPythonFileURI)
	out.Args = in.Args
	// MISSING: PythonFileUris
	// (near miss): "PythonFileUris" vs "PythonFileURIs"
	// MISSING: JarFileUris
	// (near miss): "JarFileUris" vs "JarFileURIs"
	// MISSING: FileUris
	// (near miss): "FileUris" vs "FileURIs"
	// MISSING: ArchiveUris
	// (near miss): "ArchiveUris" vs "ArchiveURIs"
	out.Properties = in.Properties
	out.LoggingConfig = LoggingConfig_ToProto(mapCtx, in.LoggingConfig)
	return out
}
func QueryList_FromProto(mapCtx *direct.MapContext, in *pb.QueryList) *krm.QueryList {
	if in == nil {
		return nil
	}
	out := &krm.QueryList{}
	out.Queries = in.Queries
	return out
}
func QueryList_ToProto(mapCtx *direct.MapContext, in *krm.QueryList) *pb.QueryList {
	if in == nil {
		return nil
	}
	out := &pb.QueryList{}
	out.Queries = in.Queries
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
func SparkBatch_FromProto(mapCtx *direct.MapContext, in *pb.SparkBatch) *krm.SparkBatch {
	if in == nil {
		return nil
	}
	out := &krm.SparkBatch{}
	out.MainJarFileURI = direct.LazyPtr(in.GetMainJarFileUri())
	out.MainClass = direct.LazyPtr(in.GetMainClass())
	out.Args = in.Args
	// MISSING: JarFileUris
	// (near miss): "JarFileUris" vs "JarFileURIs"
	// MISSING: FileUris
	// (near miss): "FileUris" vs "FileURIs"
	// MISSING: ArchiveUris
	// (near miss): "ArchiveUris" vs "ArchiveURIs"
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
	// MISSING: JarFileUris
	// (near miss): "JarFileUris" vs "JarFileURIs"
	// MISSING: FileUris
	// (near miss): "FileUris" vs "FileURIs"
	// MISSING: ArchiveUris
	// (near miss): "ArchiveUris" vs "ArchiveURIs"
	return out
}
func SparkBatch_MainJarFileUri_ToProto(mapCtx *direct.MapContext, in *string) *pb.SparkBatch_MainJarFileUri {
	if in == nil {
		return nil
	}
	return &pb.SparkBatch_MainJarFileUri{MainJarFileUri: *in}
}
func SparkBatch_MainClass_ToProto(mapCtx *direct.MapContext, in *string) *pb.SparkBatch_MainClass {
	if in == nil {
		return nil
	}
	return &pb.SparkBatch_MainClass{MainClass: *in}
}
func SparkHistoryServerConfig_FromProto(mapCtx *direct.MapContext, in *pb.SparkHistoryServerConfig) *krm.SparkHistoryServerConfig {
	if in == nil {
		return nil
	}
	out := &krm.SparkHistoryServerConfig{}
	if in.GetDataprocCluster() != "" {
		out.DataprocClusterRef = &krmdataprocv1beta1.DataprocClusterRef{External: in.GetDataprocCluster()}
	}
	return out
}
func SparkHistoryServerConfig_ToProto(mapCtx *direct.MapContext, in *krm.SparkHistoryServerConfig) *pb.SparkHistoryServerConfig {
	if in == nil {
		return nil
	}
	out := &pb.SparkHistoryServerConfig{}
	if in.DataprocClusterRef != nil {
		out.DataprocCluster = in.DataprocClusterRef.External
	}
	return out
}
func SparkJob_FromProto(mapCtx *direct.MapContext, in *pb.SparkJob) *krm.SparkJob {
	if in == nil {
		return nil
	}
	out := &krm.SparkJob{}
	out.MainJarFileURI = direct.LazyPtr(in.GetMainJarFileUri())
	out.MainClass = direct.LazyPtr(in.GetMainClass())
	out.Args = in.Args
	// MISSING: JarFileUris
	// (near miss): "JarFileUris" vs "JarFileURIs"
	// MISSING: FileUris
	// (near miss): "FileUris" vs "FileURIs"
	// MISSING: ArchiveUris
	// (near miss): "ArchiveUris" vs "ArchiveURIs"
	out.Properties = in.Properties
	out.LoggingConfig = LoggingConfig_FromProto(mapCtx, in.GetLoggingConfig())
	return out
}
func SparkJob_ToProto(mapCtx *direct.MapContext, in *krm.SparkJob) *pb.SparkJob {
	if in == nil {
		return nil
	}
	out := &pb.SparkJob{}
	if oneof := SparkJob_MainJarFileUri_ToProto(mapCtx, in.MainJarFileURI); oneof != nil {
		out.Driver = oneof
	}
	if oneof := SparkJob_MainClass_ToProto(mapCtx, in.MainClass); oneof != nil {
		out.Driver = oneof
	}
	out.Args = in.Args
	// MISSING: JarFileUris
	// (near miss): "JarFileUris" vs "JarFileURIs"
	// MISSING: FileUris
	// (near miss): "FileUris" vs "FileURIs"
	// MISSING: ArchiveUris
	// (near miss): "ArchiveUris" vs "ArchiveURIs"
	out.Properties = in.Properties
	out.LoggingConfig = LoggingConfig_ToProto(mapCtx, in.LoggingConfig)
	return out
}
func SparkJob_MainJarFileUri_ToProto(mapCtx *direct.MapContext, in *string) *pb.SparkJob_MainJarFileUri {
	if in == nil {
		return nil
	}
	return &pb.SparkJob_MainJarFileUri{MainJarFileUri: *in}
}
func SparkJob_MainClass_ToProto(mapCtx *direct.MapContext, in *string) *pb.SparkJob_MainClass {
	if in == nil {
		return nil
	}
	return &pb.SparkJob_MainClass{MainClass: *in}
}
func SparkRBatch_FromProto(mapCtx *direct.MapContext, in *pb.SparkRBatch) *krm.SparkRBatch {
	if in == nil {
		return nil
	}
	out := &krm.SparkRBatch{}
	out.MainRFileURI = direct.LazyPtr(in.GetMainRFileUri())
	out.Args = in.Args
	// MISSING: FileUris
	// (near miss): "FileUris" vs "FileURIs"
	// MISSING: ArchiveUris
	// (near miss): "ArchiveUris" vs "ArchiveURIs"
	return out
}
func SparkRBatch_ToProto(mapCtx *direct.MapContext, in *krm.SparkRBatch) *pb.SparkRBatch {
	if in == nil {
		return nil
	}
	out := &pb.SparkRBatch{}
	out.MainRFileUri = direct.ValueOf(in.MainRFileURI)
	out.Args = in.Args
	// MISSING: FileUris
	// (near miss): "FileUris" vs "FileURIs"
	// MISSING: ArchiveUris
	// (near miss): "ArchiveUris" vs "ArchiveURIs"
	return out
}
func SparkRJob_FromProto(mapCtx *direct.MapContext, in *pb.SparkRJob) *krm.SparkRJob {
	if in == nil {
		return nil
	}
	out := &krm.SparkRJob{}
	out.MainRFileURI = direct.LazyPtr(in.GetMainRFileUri())
	out.Args = in.Args
	// MISSING: FileUris
	// (near miss): "FileUris" vs "FileURIs"
	// MISSING: ArchiveUris
	// (near miss): "ArchiveUris" vs "ArchiveURIs"
	out.Properties = in.Properties
	out.LoggingConfig = LoggingConfig_FromProto(mapCtx, in.GetLoggingConfig())
	return out
}
func SparkRJob_ToProto(mapCtx *direct.MapContext, in *krm.SparkRJob) *pb.SparkRJob {
	if in == nil {
		return nil
	}
	out := &pb.SparkRJob{}
	out.MainRFileUri = direct.ValueOf(in.MainRFileURI)
	out.Args = in.Args
	// MISSING: FileUris
	// (near miss): "FileUris" vs "FileURIs"
	// MISSING: ArchiveUris
	// (near miss): "ArchiveUris" vs "ArchiveURIs"
	out.Properties = in.Properties
	out.LoggingConfig = LoggingConfig_ToProto(mapCtx, in.LoggingConfig)
	return out
}
func SparkSQLBatch_FromProto(mapCtx *direct.MapContext, in *pb.SparkSqlBatch) *krm.SparkSQLBatch {
	if in == nil {
		return nil
	}
	out := &krm.SparkSQLBatch{}
	out.QueryFileURI = direct.LazyPtr(in.GetQueryFileUri())
	out.QueryVariables = in.QueryVariables
	// MISSING: JarFileUris
	// (near miss): "JarFileUris" vs "JarFileURIs"
	return out
}
func SparkSQLBatch_ToProto(mapCtx *direct.MapContext, in *krm.SparkSQLBatch) *pb.SparkSqlBatch {
	if in == nil {
		return nil
	}
	out := &pb.SparkSqlBatch{}
	out.QueryFileUri = direct.ValueOf(in.QueryFileURI)
	out.QueryVariables = in.QueryVariables
	// MISSING: JarFileUris
	// (near miss): "JarFileUris" vs "JarFileURIs"
	return out
}
func SparkSQLJob_FromProto(mapCtx *direct.MapContext, in *pb.SparkSqlJob) *krm.SparkSQLJob {
	if in == nil {
		return nil
	}
	out := &krm.SparkSQLJob{}
	out.QueryFileURI = direct.LazyPtr(in.GetQueryFileUri())
	out.QueryList = QueryList_FromProto(mapCtx, in.GetQueryList())
	out.ScriptVariables = in.ScriptVariables
	out.Properties = in.Properties
	// MISSING: JarFileUris
	// (near miss): "JarFileUris" vs "JarFileURIs"
	out.LoggingConfig = LoggingConfig_FromProto(mapCtx, in.GetLoggingConfig())
	return out
}
func SparkSQLJob_ToProto(mapCtx *direct.MapContext, in *krm.SparkSQLJob) *pb.SparkSqlJob {
	if in == nil {
		return nil
	}
	out := &pb.SparkSqlJob{}
	if oneof := SparkSQLJob_QueryFileUri_ToProto(mapCtx, in.QueryFileURI); oneof != nil {
		out.Queries = oneof
	}
	if oneof := QueryList_ToProto(mapCtx, in.QueryList); oneof != nil {
		out.Queries = &pb.SparkSqlJob_QueryList{QueryList: oneof}
	}
	out.ScriptVariables = in.ScriptVariables
	out.Properties = in.Properties
	// MISSING: JarFileUris
	// (near miss): "JarFileUris" vs "JarFileURIs"
	out.LoggingConfig = LoggingConfig_ToProto(mapCtx, in.LoggingConfig)
	return out
}
func SparkSQLJob_QueryFileUri_ToProto(mapCtx *direct.MapContext, in *string) *pb.SparkSqlJob_QueryFileUri {
	if in == nil {
		return nil
	}
	return &pb.SparkSqlJob_QueryFileUri{QueryFileUri: *in}
}
func StartupConfig_FromProto(mapCtx *direct.MapContext, in *pb.StartupConfig) *krm.StartupConfig {
	if in == nil {
		return nil
	}
	out := &krm.StartupConfig{}
	out.RequiredRegistrationFraction = in.RequiredRegistrationFraction
	return out
}
func StartupConfig_ToProto(mapCtx *direct.MapContext, in *krm.StartupConfig) *pb.StartupConfig {
	if in == nil {
		return nil
	}
	out := &pb.StartupConfig{}
	out.RequiredRegistrationFraction = in.RequiredRegistrationFraction
	return out
}
func TrinoJob_FromProto(mapCtx *direct.MapContext, in *pb.TrinoJob) *krm.TrinoJob {
	if in == nil {
		return nil
	}
	out := &krm.TrinoJob{}
	out.QueryFileURI = direct.LazyPtr(in.GetQueryFileUri())
	out.QueryList = QueryList_FromProto(mapCtx, in.GetQueryList())
	out.ContinueOnFailure = direct.LazyPtr(in.GetContinueOnFailure())
	out.OutputFormat = direct.LazyPtr(in.GetOutputFormat())
	out.ClientTags = in.ClientTags
	out.Properties = in.Properties
	out.LoggingConfig = LoggingConfig_FromProto(mapCtx, in.GetLoggingConfig())
	return out
}
func TrinoJob_ToProto(mapCtx *direct.MapContext, in *krm.TrinoJob) *pb.TrinoJob {
	if in == nil {
		return nil
	}
	out := &pb.TrinoJob{}
	if oneof := TrinoJob_QueryFileUri_ToProto(mapCtx, in.QueryFileURI); oneof != nil {
		out.Queries = oneof
	}
	if oneof := QueryList_ToProto(mapCtx, in.QueryList); oneof != nil {
		out.Queries = &pb.TrinoJob_QueryList{QueryList: oneof}
	}
	out.ContinueOnFailure = direct.ValueOf(in.ContinueOnFailure)
	out.OutputFormat = direct.ValueOf(in.OutputFormat)
	out.ClientTags = in.ClientTags
	out.Properties = in.Properties
	out.LoggingConfig = LoggingConfig_ToProto(mapCtx, in.LoggingConfig)
	return out
}
func TrinoJob_QueryFileUri_ToProto(mapCtx *direct.MapContext, in *string) *pb.TrinoJob_QueryFileUri {
	if in == nil {
		return nil
	}
	return &pb.TrinoJob_QueryFileUri{QueryFileUri: *in}
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
func YarnApplication_FromProto(mapCtx *direct.MapContext, in *pb.YarnApplication) *krm.YarnApplication {
	if in == nil {
		return nil
	}
	out := &krm.YarnApplication{}
	out.Name = direct.LazyPtr(in.GetName())
	out.State = direct.Enum_FromProto(mapCtx, in.GetState())
	out.Progress = direct.LazyPtr(in.GetProgress())
	out.TrackingURL = direct.LazyPtr(in.GetTrackingUrl())
	return out
}
func YarnApplication_ToProto(mapCtx *direct.MapContext, in *krm.YarnApplication) *pb.YarnApplication {
	if in == nil {
		return nil
	}
	out := &pb.YarnApplication{}
	out.Name = direct.ValueOf(in.Name)
	out.State = direct.Enum_ToProto[pb.YarnApplication_State](mapCtx, in.State)
	out.Progress = direct.ValueOf(in.Progress)
	out.TrackingUrl = direct.ValueOf(in.TrackingURL)
	return out
}
