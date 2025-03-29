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
	krm "github.com/GoogleCloudPlatform/k8s-config-connector/apis/dataproc/v1beta1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/apis/refs/v1beta1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct"
)

func DataprocWorkflowTemplateObservedState_FromProto(mapCtx *direct.MapContext, in *pb.WorkflowTemplate) *krm.DataprocWorkflowTemplateObservedState {
	if in == nil {
		return nil
	}
	out := &krm.DataprocWorkflowTemplateObservedState{}
	out.Name = direct.LazyPtr(in.Name)
	out.CreateTime = direct.StringTimestamp_FromProto(mapCtx, in.GetCreateTime())
	out.UpdateTime = direct.StringTimestamp_FromProto(mapCtx, in.GetUpdateTime())
	out.Placement = WorkflowTemplatePlacementObservedState_FromProto(mapCtx, in.GetPlacement())
	return out
}
func DataprocWorkflowTemplateObservedState_ToProto(mapCtx *direct.MapContext, in *krm.DataprocWorkflowTemplateObservedState) *pb.WorkflowTemplate {
	if in == nil {
		return nil
	}
	out := &pb.WorkflowTemplate{}
	out.Name = direct.ValueOf(in.Name)
	out.CreateTime = direct.StringTimestamp_ToProto(mapCtx, in.CreateTime)
	out.UpdateTime = direct.StringTimestamp_ToProto(mapCtx, in.UpdateTime)
	out.Placement = WorkflowTemplatePlacementObservedState_ToProto(mapCtx, in.Placement)
	return out
}
func DataprocWorkflowTemplateSpec_FromProto(mapCtx *direct.MapContext, in *pb.WorkflowTemplate) *krm.DataprocWorkflowTemplateSpec {
	if in == nil {
		return nil
	}
	out := &krm.DataprocWorkflowTemplateSpec{}
	out.ID = direct.LazyPtr(in.Id)
	out.Version = direct.LazyPtr(in.GetVersion())
	out.Labels = in.Labels
	out.Placement = WorkflowTemplatePlacement_FromProto(mapCtx, in.GetPlacement())
	out.Jobs = direct.Slice_FromProto(mapCtx, in.GetJobs(), OrderedJob_FromProto)
	out.Parameters = direct.Slice_FromProto(mapCtx, in.GetParameters(), TemplateParameter_FromProto)
	out.DagTimeout = direct.Duration_FromProto(mapCtx, in.GetDagTimeout())
	out.EncryptionConfig = WorkflowTemplate_EncryptionConfig_FromProto(mapCtx, in.GetEncryptionConfig())
	return out
}
func DataprocWorkflowTemplateSpec_ToProto(mapCtx *direct.MapContext, in *krm.DataprocWorkflowTemplateSpec) *pb.WorkflowTemplate {
	if in == nil {
		return nil
	}
	out := &pb.WorkflowTemplate{}
	out.Id = direct.ValueOf(in.ID)
	out.Version = direct.ValueOf(in.Version)
	out.Labels = in.Labels
	out.Placement = WorkflowTemplatePlacement_ToProto(mapCtx, in.Placement)
	out.Jobs = direct.Slice_ToProto(mapCtx, in.Jobs, OrderedJob_ToProto)
	out.Parameters = direct.Slice_ToProto(mapCtx, in.Parameters, TemplateParameter_ToProto)
	out.DagTimeout = direct.Duration_ToProto(mapCtx, in.DagTimeout)
	out.EncryptionConfig = WorkflowTemplate_EncryptionConfig_ToProto(mapCtx, in.EncryptionConfig)
	return out
}
func OrderedJob_FromProto(mapCtx *direct.MapContext, in *pb.OrderedJob) *krm.OrderedJob {
	if in == nil {
		return nil
	}
	out := &krm.OrderedJob{}
	out.StepID = direct.LazyPtr(in.GetStepId())
	out.HadoopJob = WorkflowTemplateHadoopJob_FromProto(mapCtx, in.GetHadoopJob())
	out.SparkJob = WorkflowTemplateSparkJob_FromProto(mapCtx, in.GetSparkJob())
	out.PysparkJob = WorkflowTemplatePySparkJob_FromProto(mapCtx, in.GetPysparkJob())
	out.HiveJob = WorkflowTemplateHiveJob_FromProto(mapCtx, in.GetHiveJob())
	out.PigJob = WorkflowTemplatePigJob_FromProto(mapCtx, in.GetPigJob())
	out.SparkRJob = WorkflowTemplateSparkRJob_FromProto(mapCtx, in.GetSparkRJob())
	out.SparkSQLJob = WorkflowTemplateSparkSQLJob_FromProto(mapCtx, in.GetSparkSqlJob())
	out.PrestoJob = WorkflowTemplatePrestoJob_FromProto(mapCtx, in.GetPrestoJob())
	out.TrinoJob = WorkflowTemplateTrinoJob_FromProto(mapCtx, in.GetTrinoJob())
	out.FlinkJob = WorkflowTemplateFlinkJob_FromProto(mapCtx, in.GetFlinkJob())
	out.Labels = in.Labels
	out.Scheduling = WorkflowTemplateJobScheduling_FromProto(mapCtx, in.GetScheduling())
	out.PrerequisiteStepIDs = in.PrerequisiteStepIds
	return out
}
func OrderedJob_ToProto(mapCtx *direct.MapContext, in *krm.OrderedJob) *pb.OrderedJob {
	if in == nil {
		return nil
	}
	out := &pb.OrderedJob{}
	out.StepId = direct.ValueOf(in.StepID)
	if oneof := WorkflowTemplateHadoopJob_ToProto(mapCtx, in.HadoopJob); oneof != nil {
		out.JobType = &pb.OrderedJob_HadoopJob{HadoopJob: oneof}
	}
	if oneof := WorkflowTemplateSparkJob_ToProto(mapCtx, in.SparkJob); oneof != nil {
		out.JobType = &pb.OrderedJob_SparkJob{SparkJob: oneof}
	}
	if oneof := WorkflowTemplatePySparkJob_ToProto(mapCtx, in.PysparkJob); oneof != nil {
		out.JobType = &pb.OrderedJob_PysparkJob{PysparkJob: oneof}
	}
	if oneof := WorkflowTemplateHiveJob_ToProto(mapCtx, in.HiveJob); oneof != nil {
		out.JobType = &pb.OrderedJob_HiveJob{HiveJob: oneof}
	}
	if oneof := WorkflowTemplatePigJob_ToProto(mapCtx, in.PigJob); oneof != nil {
		out.JobType = &pb.OrderedJob_PigJob{PigJob: oneof}
	}
	if oneof := WorkflowTemplateSparkRJob_ToProto(mapCtx, in.SparkRJob); oneof != nil {
		out.JobType = &pb.OrderedJob_SparkRJob{SparkRJob: oneof}
	}
	if oneof := WorkflowTemplateSparkSQLJob_ToProto(mapCtx, in.SparkSQLJob); oneof != nil {
		out.JobType = &pb.OrderedJob_SparkSqlJob{SparkSqlJob: oneof}
	}
	if oneof := WorkflowTemplatePrestoJob_ToProto(mapCtx, in.PrestoJob); oneof != nil {
		out.JobType = &pb.OrderedJob_PrestoJob{PrestoJob: oneof}
	}
	if oneof := WorkflowTemplateTrinoJob_ToProto(mapCtx, in.TrinoJob); oneof != nil {
		out.JobType = &pb.OrderedJob_TrinoJob{TrinoJob: oneof}
	}
	if oneof := WorkflowTemplateFlinkJob_ToProto(mapCtx, in.FlinkJob); oneof != nil {
		out.JobType = &pb.OrderedJob_FlinkJob{FlinkJob: oneof}
	}
	out.Labels = in.Labels
	out.Scheduling = WorkflowTemplateJobScheduling_ToProto(mapCtx, in.Scheduling)
	out.PrerequisiteStepIds = in.PrerequisiteStepIDs
	return out
}

func WorkflowTemplateJobScheduling_FromProto(mapCtx *direct.MapContext, in *pb.JobScheduling) *krm.JobScheduling {
	if in == nil {
		return nil
	}
	out := &krm.JobScheduling{}
	out.MaxFailuresPerHour = direct.LazyPtr(in.GetMaxFailuresPerHour())
	out.MaxFailuresTotal = direct.LazyPtr(in.GetMaxFailuresTotal())
	return out
}
func WorkflowTemplateJobScheduling_ToProto(mapCtx *direct.MapContext, in *krm.JobScheduling) *pb.JobScheduling {
	if in == nil {
		return nil
	}
	out := &pb.JobScheduling{}
	out.MaxFailuresPerHour = direct.ValueOf(in.MaxFailuresPerHour)
	out.MaxFailuresTotal = direct.ValueOf(in.MaxFailuresTotal)
	return out
}
func WorkflowTemplateFlinkJob_FromProto(mapCtx *direct.MapContext, in *pb.FlinkJob) *krm.FlinkJob {
	if in == nil {
		return nil
	}
	out := &krm.FlinkJob{}
	out.MainJarFileURI = direct.LazyPtr(in.GetMainJarFileUri())
	out.MainClass = direct.LazyPtr(in.GetMainClass())
	out.Args = append(out.Args, in.Args...)
	out.JarFileURIs = in.GetJarFileUris()
	out.SavepointURI = direct.LazyPtr(in.GetSavepointUri())
	out.Properties = in.Properties
	out.LoggingConfig = WorkflowTemplateLoggingConfig_FromProto(mapCtx, in.GetLoggingConfig())
	return out
}
func WorkflowTemplateFlinkJob_ToProto(mapCtx *direct.MapContext, in *krm.FlinkJob) *pb.FlinkJob {
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
	out.Args = append(out.Args, in.Args...)
	out.JarFileUris = in.JarFileURIs
	out.SavepointUri = direct.ValueOf(in.SavepointURI)
	out.Properties = in.Properties
	out.LoggingConfig = WorkflowTemplateLoggingConfig_ToProto(mapCtx, in.LoggingConfig)
	return out
}
func WorkflowTemplateHadoopJob_FromProto(mapCtx *direct.MapContext, in *pb.HadoopJob) *krm.HadoopJob {
	if in == nil {
		return nil
	}
	out := &krm.HadoopJob{}
	out.MainJarFileURI = direct.LazyPtr(in.GetMainJarFileUri())
	out.MainClass = direct.LazyPtr(in.GetMainClass())
	out.Args = append(out.Args, in.Args...)
	out.JarFileURIs = in.JarFileUris
	out.FileURIs = in.FileUris
	out.ArchiveURIs = in.ArchiveUris
	out.Properties = in.Properties
	out.LoggingConfig = WorkflowTemplateLoggingConfig_FromProto(mapCtx, in.GetLoggingConfig())
	return out
}
func WorkflowTemplateHadoopJob_ToProto(mapCtx *direct.MapContext, in *krm.HadoopJob) *pb.HadoopJob {
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
	out.Args = append(out.Args, in.Args...)
	out.JarFileUris = in.JarFileURIs
	out.FileUris = in.FileURIs
	out.ArchiveUris = in.ArchiveURIs
	out.Properties = in.Properties
	out.LoggingConfig = WorkflowTemplateLoggingConfig_ToProto(mapCtx, in.LoggingConfig)
	return out
}
func WorkflowTemplateHiveJob_FromProto(mapCtx *direct.MapContext, in *pb.HiveJob) *krm.HiveJob {
	if in == nil {
		return nil
	}
	out := &krm.HiveJob{}
	out.QueryFileURI = direct.LazyPtr(in.GetQueryFileUri())
	out.QueryList = WorkflowTemplateQueryList_FromProto(mapCtx, in.GetQueryList())
	out.ContinueOnFailure = direct.LazyPtr(in.GetContinueOnFailure())
	out.ScriptVariables = in.ScriptVariables
	out.Properties = in.Properties
	out.JarFileURIs = in.JarFileUris
	return out
}
func WorkflowTemplateHiveJob_ToProto(mapCtx *direct.MapContext, in *krm.HiveJob) *pb.HiveJob {
	if in == nil {
		return nil
	}
	out := &pb.HiveJob{}
	if oneof := HiveJob_QueryFileUri_ToProto(mapCtx, in.QueryFileURI); oneof != nil {
		out.Queries = oneof
	}
	if oneof := WorkflowTemplateQueryList_ToProto(mapCtx, in.QueryList); oneof != nil {
		out.Queries = &pb.HiveJob_QueryList{QueryList: oneof}
	}
	out.ContinueOnFailure = direct.ValueOf(in.ContinueOnFailure)
	out.ScriptVariables = in.ScriptVariables
	out.Properties = in.Properties
	out.JarFileUris = in.JarFileURIs
	return out
}

func WorkflowTemplatePigJob_FromProto(mapCtx *direct.MapContext, in *pb.PigJob) *krm.PigJob {
	if in == nil {
		return nil
	}
	out := &krm.PigJob{}
	out.QueryFileURI = direct.LazyPtr(in.GetQueryFileUri())
	out.QueryList = WorkflowTemplateQueryList_FromProto(mapCtx, in.GetQueryList())
	out.ContinueOnFailure = direct.LazyPtr(in.GetContinueOnFailure())
	out.ScriptVariables = in.ScriptVariables
	out.Properties = in.Properties
	out.JarFileURIs = in.JarFileUris
	out.LoggingConfig = WorkflowTemplateLoggingConfig_FromProto(mapCtx, in.GetLoggingConfig())
	return out
}
func WorkflowTemplatePigJob_ToProto(mapCtx *direct.MapContext, in *krm.PigJob) *pb.PigJob {
	if in == nil {
		return nil
	}
	out := &pb.PigJob{}
	if oneof := PigJob_QueryFileUri_ToProto(mapCtx, in.QueryFileURI); oneof != nil {
		out.Queries = oneof
	}
	if oneof := WorkflowTemplateQueryList_ToProto(mapCtx, in.QueryList); oneof != nil {
		out.Queries = &pb.PigJob_QueryList{QueryList: oneof}
	}
	out.ContinueOnFailure = direct.ValueOf(in.ContinueOnFailure)
	out.ScriptVariables = in.ScriptVariables
	out.Properties = in.Properties
	out.JarFileUris = in.JarFileURIs
	out.LoggingConfig = WorkflowTemplateLoggingConfig_ToProto(mapCtx, in.LoggingConfig)
	return out
}

func WorkflowTemplatePySparkJob_FromProto(mapCtx *direct.MapContext, in *pb.PySparkJob) *krm.PySparkJob {
	if in == nil {
		return nil
	}
	out := &krm.PySparkJob{}
	out.MainPythonFileURI = direct.LazyPtr(in.GetMainPythonFileUri())
	out.Args = append(out.Args, in.Args...)
	out.PythonFileURIs = in.PythonFileUris
	out.JarFileURIs = in.JarFileUris
	out.FileURIs = in.FileUris
	out.ArchiveURIs = in.ArchiveUris
	out.Properties = in.Properties
	out.LoggingConfig = WorkflowTemplateLoggingConfig_FromProto(mapCtx, in.GetLoggingConfig())
	return out
}
func WorkflowTemplatePySparkJob_ToProto(mapCtx *direct.MapContext, in *krm.PySparkJob) *pb.PySparkJob {
	if in == nil {
		return nil
	}
	out := &pb.PySparkJob{}
	out.MainPythonFileUri = direct.ValueOf(in.MainPythonFileURI)
	out.Args = append(out.Args, in.Args...)
	out.PythonFileUris = in.PythonFileURIs
	out.JarFileUris = in.JarFileURIs
	out.FileUris = in.FileURIs
	out.ArchiveUris = in.ArchiveURIs
	out.Properties = in.Properties
	out.LoggingConfig = WorkflowTemplateLoggingConfig_ToProto(mapCtx, in.LoggingConfig)
	return out
}
func WorkflowTemplateSparkJob_FromProto(mapCtx *direct.MapContext, in *pb.SparkJob) *krm.SparkJob {
	if in == nil {
		return nil
	}
	out := &krm.SparkJob{}
	out.MainJarFileURI = direct.LazyPtr(in.GetMainJarFileUri())
	out.MainClass = direct.LazyPtr(in.GetMainClass())
	out.Args = append(out.Args, in.Args...)
	out.JarFileURIs = in.JarFileUris
	out.FileURIs = in.FileUris
	out.ArchiveURIs = in.ArchiveUris
	out.Properties = in.Properties
	out.LoggingConfig = WorkflowTemplateLoggingConfig_FromProto(mapCtx, in.GetLoggingConfig())
	return out
}
func WorkflowTemplateSparkJob_ToProto(mapCtx *direct.MapContext, in *krm.SparkJob) *pb.SparkJob {
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
	out.Args = append(out.Args, in.Args...)
	out.JarFileUris = in.JarFileURIs
	out.FileUris = in.FileURIs
	out.ArchiveUris = in.ArchiveURIs
	out.Properties = in.Properties
	out.LoggingConfig = WorkflowTemplateLoggingConfig_ToProto(mapCtx, in.LoggingConfig)
	return out
}
func WorkflowTemplateSparkRJob_FromProto(mapCtx *direct.MapContext, in *pb.SparkRJob) *krm.SparkRJob {
	if in == nil {
		return nil
	}
	out := &krm.SparkRJob{}
	out.MainRFileURI = direct.LazyPtr(in.GetMainRFileUri())
	out.Args = append(out.Args, in.Args...)
	out.FileURIs = in.FileUris
	out.ArchiveURIs = in.ArchiveUris
	out.Properties = in.Properties
	out.LoggingConfig = WorkflowTemplateLoggingConfig_FromProto(mapCtx, in.GetLoggingConfig())
	return out
}
func WorkflowTemplateSparkRJob_ToProto(mapCtx *direct.MapContext, in *krm.SparkRJob) *pb.SparkRJob {
	if in == nil {
		return nil
	}
	out := &pb.SparkRJob{}
	out.MainRFileUri = direct.ValueOf(in.MainRFileURI)
	out.Args = append(out.Args, in.Args...)
	out.FileUris = in.FileURIs
	out.ArchiveUris = in.ArchiveURIs
	out.Properties = in.Properties
	out.LoggingConfig = WorkflowTemplateLoggingConfig_ToProto(mapCtx, in.LoggingConfig)
	return out
}
func WorkflowTemplateSparkSQLJob_FromProto(mapCtx *direct.MapContext, in *pb.SparkSqlJob) *krm.SparkSQLJob {
	if in == nil {
		return nil
	}
	out := &krm.SparkSQLJob{}
	out.QueryFileURI = direct.LazyPtr(in.GetQueryFileUri())
	out.QueryList = WorkflowTemplateQueryList_FromProto(mapCtx, in.GetQueryList())
	out.ScriptVariables = in.ScriptVariables
	out.Properties = in.Properties
	out.JarFileURIs = in.JarFileUris
	out.LoggingConfig = WorkflowTemplateLoggingConfig_FromProto(mapCtx, in.GetLoggingConfig())
	return out
}
func WorkflowTemplateSparkSQLJob_ToProto(mapCtx *direct.MapContext, in *krm.SparkSQLJob) *pb.SparkSqlJob {
	if in == nil {
		return nil
	}
	out := &pb.SparkSqlJob{}
	if oneof := SparkSQLJob_QueryFileUri_ToProto(mapCtx, in.QueryFileURI); oneof != nil {
		out.Queries = oneof
	}
	if oneof := WorkflowTemplateQueryList_ToProto(mapCtx, in.QueryList); oneof != nil {
		out.Queries = &pb.SparkSqlJob_QueryList{QueryList: oneof}
	}
	out.ScriptVariables = in.ScriptVariables
	out.Properties = in.Properties
	out.JarFileUris = in.JarFileURIs
	out.LoggingConfig = WorkflowTemplateLoggingConfig_ToProto(mapCtx, in.LoggingConfig)
	return out
}
func WorkflowTemplateLoggingConfig_FromProto(mapCtx *direct.MapContext, in *pb.LoggingConfig) *krm.LoggingConfig {
	if in == nil {
		return nil
	}
	out := &krm.LoggingConfig{}
	// MISSING: DriverLogLevels
	return out
}
func WorkflowTemplateLoggingConfig_ToProto(mapCtx *direct.MapContext, in *krm.LoggingConfig) *pb.LoggingConfig {
	if in == nil {
		return nil
	}
	out := &pb.LoggingConfig{}
	// MISSING: DriverLogLevels
	return out
}
func WorkflowTemplatePrestoJob_FromProto(mapCtx *direct.MapContext, in *pb.PrestoJob) *krm.PrestoJob {
	if in == nil {
		return nil
	}
	out := &krm.PrestoJob{}
	out.QueryFileURI = direct.LazyPtr(in.GetQueryFileUri())
	out.QueryList = WorkflowTemplateQueryList_FromProto(mapCtx, in.GetQueryList())
	out.ContinueOnFailure = direct.LazyPtr(in.GetContinueOnFailure())
	out.OutputFormat = direct.LazyPtr(in.GetOutputFormat())
	out.ClientTags = in.ClientTags
	out.Properties = in.Properties
	out.LoggingConfig = WorkflowTemplateLoggingConfig_FromProto(mapCtx, in.GetLoggingConfig())
	return out
}
func WorkflowTemplatePrestoJob_ToProto(mapCtx *direct.MapContext, in *krm.PrestoJob) *pb.PrestoJob {
	if in == nil {
		return nil
	}
	out := &pb.PrestoJob{}
	if oneof := PrestoJob_QueryFileUri_ToProto(mapCtx, in.QueryFileURI); oneof != nil {
		out.Queries = oneof
	}
	if oneof := WorkflowTemplateQueryList_ToProto(mapCtx, in.QueryList); oneof != nil {
		out.Queries = &pb.PrestoJob_QueryList{QueryList: oneof}
	}
	out.ContinueOnFailure = direct.ValueOf(in.ContinueOnFailure)
	out.OutputFormat = direct.ValueOf(in.OutputFormat)
	out.ClientTags = in.ClientTags
	out.Properties = in.Properties
	out.LoggingConfig = WorkflowTemplateLoggingConfig_ToProto(mapCtx, in.LoggingConfig)
	return out
}
func WorkflowTemplateQueryList_FromProto(mapCtx *direct.MapContext, in *pb.QueryList) *krm.QueryList {
	if in == nil {
		return nil
	}
	out := &krm.QueryList{}
	out.Queries = in.Queries
	return out
}
func WorkflowTemplateQueryList_ToProto(mapCtx *direct.MapContext, in *krm.QueryList) *pb.QueryList {
	if in == nil {
		return nil
	}
	out := &pb.QueryList{}
	out.Queries = in.Queries
	return out
}
func WorkflowTemplateTrinoJob_FromProto(mapCtx *direct.MapContext, in *pb.TrinoJob) *krm.TrinoJob {
	if in == nil {
		return nil
	}
	out := &krm.TrinoJob{}
	out.QueryFileURI = direct.LazyPtr(in.GetQueryFileUri())
	out.QueryList = WorkflowTemplateQueryList_FromProto(mapCtx, in.GetQueryList())
	out.ContinueOnFailure = direct.LazyPtr(in.GetContinueOnFailure())
	out.OutputFormat = direct.LazyPtr(in.GetOutputFormat())
	out.ClientTags = in.ClientTags
	out.Properties = in.Properties
	out.LoggingConfig = WorkflowTemplateLoggingConfig_FromProto(mapCtx, in.GetLoggingConfig())
	return out
}
func WorkflowTemplateTrinoJob_ToProto(mapCtx *direct.MapContext, in *krm.TrinoJob) *pb.TrinoJob {
	if in == nil {
		return nil
	}
	out := &pb.TrinoJob{}
	if oneof := TrinoJob_QueryFileUri_ToProto(mapCtx, in.QueryFileURI); oneof != nil {
		out.Queries = oneof
	}
	if oneof := WorkflowTemplateQueryList_ToProto(mapCtx, in.QueryList); oneof != nil {
		out.Queries = &pb.TrinoJob_QueryList{QueryList: oneof}
	}
	out.ContinueOnFailure = direct.ValueOf(in.ContinueOnFailure)
	out.OutputFormat = direct.ValueOf(in.OutputFormat)
	out.ClientTags = in.ClientTags
	out.Properties = in.Properties
	out.LoggingConfig = WorkflowTemplateLoggingConfig_ToProto(mapCtx, in.LoggingConfig)
	return out
}
func WorkflowTemplate_EncryptionConfig_FromProto(mapCtx *direct.MapContext, in *pb.WorkflowTemplate_EncryptionConfig) *krm.WorkflowTemplate_EncryptionConfig {
	if in == nil {
		return nil
	}
	out := &krm.WorkflowTemplate_EncryptionConfig{}
	out.KMSKeyRef = &v1beta1.KMSCryptoKeyRef{
		External: in.GetKmsKey(),
	}
	return out
}
func WorkflowTemplate_EncryptionConfig_ToProto(mapCtx *direct.MapContext, in *krm.WorkflowTemplate_EncryptionConfig) *pb.WorkflowTemplate_EncryptionConfig {
	if in == nil {
		return nil
	}
	out := &pb.WorkflowTemplate_EncryptionConfig{}
	if in.KMSKeyRef != nil {
		out.KmsKey = in.KMSKeyRef.External
	}
	return out
}
