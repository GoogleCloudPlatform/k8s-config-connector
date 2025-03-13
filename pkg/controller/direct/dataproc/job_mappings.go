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
	out.JobUUid = direct.LazyPtr(in.JobUuid)
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
	out.JobUuid = direct.ValueOf(in.JobUUid)
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

func FlinkJob_FromProto(mapCtx *direct.MapContext, in *pb.FlinkJob) *krm.FlinkJob {
	if in == nil {
		return nil
	}
	out := &krm.FlinkJob{}
	out.MainJarFileURI = direct.LazyPtr(in.GetMainJarFileUri())
	out.MainClass = direct.LazyPtr(in.GetMainClass())
	out.Args = in.Args
	out.JarFileURIs = in.GetJarFileUris()
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
	out.JarFileUris = in.JarFileURIs
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
	out.JarFileURIs = in.JarFileUris
	out.FileURIs = in.FileUris
	out.ArchiveURIs = in.ArchiveUris
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
	out.JarFileUris = in.JarFileURIs
	out.FileUris = in.FileURIs
	out.ArchiveUris = in.ArchiveURIs
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
	out.JarFileURIs = in.JarFileUris
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
	out.JarFileUris = in.JarFileURIs
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
	out.JarFileURIs = in.JarFileUris
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
	out.JarFileUris = in.JarFileURIs
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
	out.PythonFileURIs = in.PythonFileUris
	out.JarFileURIs = in.JarFileUris
	out.FileURIs = in.FileUris
	out.ArchiveURIs = in.ArchiveUris
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
	out.PythonFileUris = in.PythonFileURIs
	out.JarFileUris = in.JarFileURIs
	out.FileUris = in.FileURIs
	out.ArchiveUris = in.ArchiveURIs
	out.Properties = in.Properties
	out.LoggingConfig = LoggingConfig_ToProto(mapCtx, in.LoggingConfig)
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
	out.JarFileURIs = in.JarFileUris
	out.FileURIs = in.FileUris
	out.ArchiveURIs = in.ArchiveUris
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
	out.JarFileUris = in.JarFileURIs
	out.FileUris = in.FileURIs
	out.ArchiveUris = in.ArchiveURIs
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
	out.FileURIs = in.FileUris
	out.ArchiveURIs = in.ArchiveUris
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
	out.FileUris = in.FileURIs
	out.ArchiveUris = in.ArchiveURIs
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
	out.JarFileURIs = in.JarFileUris
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
	out.JarFileUris = in.JarFileURIs
	out.LoggingConfig = LoggingConfig_ToProto(mapCtx, in.LoggingConfig)
	return out
}

func FlinkJob_MainJarFileUri_ToProto(mapCtx *direct.MapContext, in *string) *pb.FlinkJob_MainJarFileUri {
	if in == nil {
		return nil
	}
	out := &pb.FlinkJob_MainJarFileUri{
		MainJarFileUri: direct.ValueOf(in),
	}
	return out
}

func FlinkJob_MainClass_ToProto(mapCtx *direct.MapContext, in *string) *pb.FlinkJob_MainClass {
	if in == nil {
		return nil
	}
	out := &pb.FlinkJob_MainClass{
		MainClass: direct.ValueOf(in),
	}
	return out
}

func HadoopJob_MainJarFileUri_ToProto(mapCtx *direct.MapContext, in *string) *pb.HadoopJob_MainJarFileUri {
	if in == nil {
		return nil
	}
	out := &pb.HadoopJob_MainJarFileUri{
		MainJarFileUri: direct.ValueOf(in),
	}
	return out
}

func HadoopJob_MainClass_ToProto(mapCtx *direct.MapContext, in *string) *pb.HadoopJob_MainClass {
	if in == nil {
		return nil
	}
	out := &pb.HadoopJob_MainClass{
		MainClass: direct.ValueOf(in),
	}
	return out
}

func HiveJob_QueryFileUri_ToProto(mapCtx *direct.MapContext, in *string) *pb.HiveJob_QueryFileUri {
	if in == nil {
		return nil
	}
	out := &pb.HiveJob_QueryFileUri{
		QueryFileUri: direct.ValueOf(in),
	}
	return out
}

func PrestoJob_QueryFileUri_ToProto(mapCtx *direct.MapContext, in *string) *pb.PrestoJob_QueryFileUri {
	if in == nil {
		return nil
	}
	out := &pb.PrestoJob_QueryFileUri{
		QueryFileUri: direct.ValueOf(in),
	}
	return out
}

func SparkJob_MainJarFileUri_ToProto(mapCtx *direct.MapContext, in *string) *pb.SparkJob_MainJarFileUri {
	if in == nil {
		return nil
	}
	out := &pb.SparkJob_MainJarFileUri{
		MainJarFileUri: direct.ValueOf(in),
	}
	return out
}
func SparkJob_MainClass_ToProto(mapCtx *direct.MapContext, in *string) *pb.SparkJob_MainClass {
	if in == nil {
		return nil
	}
	out := &pb.SparkJob_MainClass{
		MainClass: direct.ValueOf(in),
	}
	return out
}

func PigJob_QueryFileUri_ToProto(mapCtx *direct.MapContext, in *string) *pb.PigJob_QueryFileUri {
	if in == nil {
		return nil
	}
	out := &pb.PigJob_QueryFileUri{
		QueryFileUri: direct.ValueOf(in),
	}
	return out
}

func SparkSQLJob_QueryFileUri_ToProto(mapCtx *direct.MapContext, in *string) *pb.SparkSqlJob_QueryFileUri {
	if in == nil {
		return nil
	}
	out := &pb.SparkSqlJob_QueryFileUri{
		QueryFileUri: direct.ValueOf(in),
	}
	return out
}
func TrinoJob_QueryFileUri_ToProto(mapCtx *direct.MapContext, in *string) *pb.TrinoJob_QueryFileUri {
	if in == nil {
		return nil
	}
	out := &pb.TrinoJob_QueryFileUri{
		QueryFileUri: direct.ValueOf(in),
	}
	return out
}
