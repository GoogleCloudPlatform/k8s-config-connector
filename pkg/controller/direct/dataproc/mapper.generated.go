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
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct"
)

func DataprocJobObservedState_FromProto(mapCtx *direct.MapContext, in *pb.Job) *krm.DataprocJobObservedState {
	if in == nil {
		return nil
	}
	out := &krm.DataprocJobObservedState{}
	out.Placement = JobPlacementObservedState_FromProto(mapCtx, in.GetPlacement())
	out.Status = JobStatus_FromProto(mapCtx, in.GetStatus())
	out.StatusHistory = direct.Slice_FromProto(mapCtx, in.StatusHistory, JobStatus_FromProto)
	out.YarnApplications = direct.Slice_FromProto(mapCtx, in.YarnApplications, YarnApplication_FromProto)
	out.DriverOutputResourceURI = direct.LazyPtr(in.GetDriverOutputResourceUri())
	out.DriverControlFilesURI = direct.LazyPtr(in.GetDriverControlFilesUri())
	out.JobUuid = direct.LazyPtr(in.GetJobUuid())
	out.Done = direct.LazyPtr(in.GetDone())
	return out
}
func DataprocJobObservedState_ToProto(mapCtx *direct.MapContext, in *krm.DataprocJobObservedState) *pb.Job {
	if in == nil {
		return nil
	}
	out := &pb.Job{}
	out.Placement = JobPlacementObservedState_ToProto(mapCtx, in.Placement)
	out.Status = JobStatus_ToProto(mapCtx, in.Status)
	out.StatusHistory = direct.Slice_ToProto(mapCtx, in.StatusHistory, JobStatus_ToProto)
	out.YarnApplications = direct.Slice_ToProto(mapCtx, in.YarnApplications, YarnApplication_ToProto)
	out.DriverOutputResourceUri = direct.ValueOf(in.DriverOutputResourceURI)
	out.DriverControlFilesUri = direct.ValueOf(in.DriverControlFilesURI)
	out.JobUuid = direct.ValueOf(in.JobUuid)
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
	out.DriverSchedulingConfig = DriverSchedulingConfig_ToProto(mapCtx, in.DriverSchedulingConfig)
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
func FlinkJob_FromProto(mapCtx *direct.MapContext, in *pb.FlinkJob) *krm.FlinkJob {
	if in == nil {
		return nil
	}
	out := &krm.FlinkJob{}
	out.MainJarFileURI = direct.LazyPtr(in.GetMainJarFileUri())
	out.MainClass = direct.LazyPtr(in.GetMainClass())
	out.Args = in.Args
	out.JarFileUris = in.JarFileUris
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
	out.JarFileUris = in.JarFileUris
	out.SavepointUri = direct.ValueOf(in.SavepointURI)
	out.Properties = in.Properties
	out.LoggingConfig = LoggingConfig_ToProto(mapCtx, in.LoggingConfig)
	return out
}
func HadoopJob_FromProto(mapCtx *direct.MapContext, in *pb.HadoopJob) *krm.HadoopJob {
	if in == nil {
		return nil
	}
	out := &krm.HadoopJob{}
	out.MainJarFileURI = direct.LazyPtr(in.GetMainJarFileUri())
	out.MainClass = direct.LazyPtr(in.GetMainClass())
	out.Args = in.Args
	out.JarFileUris = in.JarFileUris
	out.FileUris = in.FileUris
	out.ArchiveUris = in.ArchiveUris
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
	out.JarFileUris = in.JarFileUris
	out.FileUris = in.FileUris
	out.ArchiveUris = in.ArchiveUris
	out.Properties = in.Properties
	out.LoggingConfig = LoggingConfig_ToProto(mapCtx, in.LoggingConfig)
	return out
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
	out.JarFileUris = in.JarFileUris
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
	out.JarFileUris = in.JarFileUris
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
	out.JarFileUris = in.JarFileUris
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
	out.JarFileUris = in.JarFileUris
	out.LoggingConfig = LoggingConfig_ToProto(mapCtx, in.LoggingConfig)
	return out
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
func PySparkJob_FromProto(mapCtx *direct.MapContext, in *pb.PySparkJob) *krm.PySparkJob {
	if in == nil {
		return nil
	}
	out := &krm.PySparkJob{}
	out.MainPythonFileURI = direct.LazyPtr(in.GetMainPythonFileUri())
	out.Args = in.Args
	out.PythonFileUris = in.PythonFileUris
	out.JarFileUris = in.JarFileUris
	out.FileUris = in.FileUris
	out.ArchiveUris = in.ArchiveUris
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
	out.PythonFileUris = in.PythonFileUris
	out.JarFileUris = in.JarFileUris
	out.FileUris = in.FileUris
	out.ArchiveUris = in.ArchiveUris
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
func SparkJob_FromProto(mapCtx *direct.MapContext, in *pb.SparkJob) *krm.SparkJob {
	if in == nil {
		return nil
	}
	out := &krm.SparkJob{}
	out.MainJarFileURI = direct.LazyPtr(in.GetMainJarFileUri())
	out.MainClass = direct.LazyPtr(in.GetMainClass())
	out.Args = in.Args
	out.JarFileUris = in.JarFileUris
	out.FileUris = in.FileUris
	out.ArchiveUris = in.ArchiveUris
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
	out.JarFileUris = in.JarFileUris
	out.FileUris = in.FileUris
	out.ArchiveUris = in.ArchiveUris
	out.Properties = in.Properties
	out.LoggingConfig = LoggingConfig_ToProto(mapCtx, in.LoggingConfig)
	return out
}
func SparkRJob_FromProto(mapCtx *direct.MapContext, in *pb.SparkRJob) *krm.SparkRJob {
	if in == nil {
		return nil
	}
	out := &krm.SparkRJob{}
	out.MainRFileURI = direct.LazyPtr(in.GetMainRFileUri())
	out.Args = in.Args
	out.FileUris = in.FileUris
	out.ArchiveUris = in.ArchiveUris
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
	out.FileUris = in.FileUris
	out.ArchiveUris = in.ArchiveUris
	out.Properties = in.Properties
	out.LoggingConfig = LoggingConfig_ToProto(mapCtx, in.LoggingConfig)
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
	out.JarFileUris = in.JarFileUris
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
	out.JarFileUris = in.JarFileUris
	out.LoggingConfig = LoggingConfig_ToProto(mapCtx, in.LoggingConfig)
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
