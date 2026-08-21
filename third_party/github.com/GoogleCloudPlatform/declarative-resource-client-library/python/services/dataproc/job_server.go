// Copyright 2021 Google LLC. All Rights Reserved.
// 
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
// 
//     http://www.apache.org/licenses/LICENSE-2.0
// 
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.
package server

import (
	"context"

	"github.com/GoogleCloudPlatform/declarative-resource-client-library/dcl"
	dataprocpb "github.com/GoogleCloudPlatform/declarative-resource-client-library/python/proto/dataproc/dataproc_go_proto"
	emptypb "github.com/GoogleCloudPlatform/declarative-resource-client-library/python/proto/empty_go_proto"
	"github.com/GoogleCloudPlatform/declarative-resource-client-library/services/google/dataproc"
)

// Server implements the gRPC interface for Job.
type JobServer struct{}

// ProtoToJobStatusStateEnum converts a JobStatusStateEnum enum from its proto representation.
func ProtoToDataprocJobStatusStateEnum(e dataprocpb.DataprocJobStatusStateEnum) *dataproc.JobStatusStateEnum {
	if e == 0 {
		return nil
	}
	if n, ok := dataprocpb.DataprocJobStatusStateEnum_name[int32(e)]; ok {
		e := dataproc.JobStatusStateEnum(n[len("DataprocJobStatusStateEnum"):])
		return &e
	}
	return nil
}

// ProtoToJobStatusSubstateEnum converts a JobStatusSubstateEnum enum from its proto representation.
func ProtoToDataprocJobStatusSubstateEnum(e dataprocpb.DataprocJobStatusSubstateEnum) *dataproc.JobStatusSubstateEnum {
	if e == 0 {
		return nil
	}
	if n, ok := dataprocpb.DataprocJobStatusSubstateEnum_name[int32(e)]; ok {
		e := dataproc.JobStatusSubstateEnum(n[len("DataprocJobStatusSubstateEnum"):])
		return &e
	}
	return nil
}

// ProtoToJobStatusHistoryStateEnum converts a JobStatusHistoryStateEnum enum from its proto representation.
func ProtoToDataprocJobStatusHistoryStateEnum(e dataprocpb.DataprocJobStatusHistoryStateEnum) *dataproc.JobStatusHistoryStateEnum {
	if e == 0 {
		return nil
	}
	if n, ok := dataprocpb.DataprocJobStatusHistoryStateEnum_name[int32(e)]; ok {
		e := dataproc.JobStatusHistoryStateEnum(n[len("DataprocJobStatusHistoryStateEnum"):])
		return &e
	}
	return nil
}

// ProtoToJobStatusHistorySubstateEnum converts a JobStatusHistorySubstateEnum enum from its proto representation.
func ProtoToDataprocJobStatusHistorySubstateEnum(e dataprocpb.DataprocJobStatusHistorySubstateEnum) *dataproc.JobStatusHistorySubstateEnum {
	if e == 0 {
		return nil
	}
	if n, ok := dataprocpb.DataprocJobStatusHistorySubstateEnum_name[int32(e)]; ok {
		e := dataproc.JobStatusHistorySubstateEnum(n[len("DataprocJobStatusHistorySubstateEnum"):])
		return &e
	}
	return nil
}

// ProtoToJobYarnApplicationsStateEnum converts a JobYarnApplicationsStateEnum enum from its proto representation.
func ProtoToDataprocJobYarnApplicationsStateEnum(e dataprocpb.DataprocJobYarnApplicationsStateEnum) *dataproc.JobYarnApplicationsStateEnum {
	if e == 0 {
		return nil
	}
	if n, ok := dataprocpb.DataprocJobYarnApplicationsStateEnum_name[int32(e)]; ok {
		e := dataproc.JobYarnApplicationsStateEnum(n[len("DataprocJobYarnApplicationsStateEnum"):])
		return &e
	}
	return nil
}

// ProtoToJobReference converts a JobReference resource from its proto representation.
func ProtoToDataprocJobReference(p *dataprocpb.DataprocJobReference) *dataproc.JobReference {
	if p == nil {
		return nil
	}
	obj := &dataproc.JobReference{
		ProjectId: dcl.StringOrNil(p.ProjectId),
		JobId:     dcl.StringOrNil(p.JobId),
	}
	return obj
}

// ProtoToJobPlacement converts a JobPlacement resource from its proto representation.
func ProtoToDataprocJobPlacement(p *dataprocpb.DataprocJobPlacement) *dataproc.JobPlacement {
	if p == nil {
		return nil
	}
	obj := &dataproc.JobPlacement{
		ClusterName: dcl.StringOrNil(p.ClusterName),
		ClusterUuid: dcl.StringOrNil(p.ClusterUuid),
	}
	return obj
}

// ProtoToJobHadoopJob converts a JobHadoopJob resource from its proto representation.
func ProtoToDataprocJobHadoopJob(p *dataprocpb.DataprocJobHadoopJob) *dataproc.JobHadoopJob {
	if p == nil {
		return nil
	}
	obj := &dataproc.JobHadoopJob{
		MainJarFileUri: dcl.StringOrNil(p.MainJarFileUri),
		MainClass:      dcl.StringOrNil(p.MainClass),
		LoggingConfig:  ProtoToDataprocJobHadoopJobLoggingConfig(p.GetLoggingConfig()),
	}
	for _, r := range p.GetArgs() {
		obj.Args = append(obj.Args, r)
	}
	for _, r := range p.GetJarFileUris() {
		obj.JarFileUris = append(obj.JarFileUris, r)
	}
	for _, r := range p.GetFileUris() {
		obj.FileUris = append(obj.FileUris, r)
	}
	for _, r := range p.GetArchiveUris() {
		obj.ArchiveUris = append(obj.ArchiveUris, r)
	}
	return obj
}

// ProtoToJobHadoopJobLoggingConfig converts a JobHadoopJobLoggingConfig resource from its proto representation.
func ProtoToDataprocJobHadoopJobLoggingConfig(p *dataprocpb.DataprocJobHadoopJobLoggingConfig) *dataproc.JobHadoopJobLoggingConfig {
	if p == nil {
		return nil
	}
	obj := &dataproc.JobHadoopJobLoggingConfig{}
	return obj
}

// ProtoToJobSparkJob converts a JobSparkJob resource from its proto representation.
func ProtoToDataprocJobSparkJob(p *dataprocpb.DataprocJobSparkJob) *dataproc.JobSparkJob {
	if p == nil {
		return nil
	}
	obj := &dataproc.JobSparkJob{
		MainJarFileUri: dcl.StringOrNil(p.MainJarFileUri),
		MainClass:      dcl.StringOrNil(p.MainClass),
		LoggingConfig:  ProtoToDataprocJobSparkJobLoggingConfig(p.GetLoggingConfig()),
	}
	for _, r := range p.GetArgs() {
		obj.Args = append(obj.Args, r)
	}
	for _, r := range p.GetJarFileUris() {
		obj.JarFileUris = append(obj.JarFileUris, r)
	}
	for _, r := range p.GetFileUris() {
		obj.FileUris = append(obj.FileUris, r)
	}
	for _, r := range p.GetArchiveUris() {
		obj.ArchiveUris = append(obj.ArchiveUris, r)
	}
	return obj
}

// ProtoToJobSparkJobLoggingConfig converts a JobSparkJobLoggingConfig resource from its proto representation.
func ProtoToDataprocJobSparkJobLoggingConfig(p *dataprocpb.DataprocJobSparkJobLoggingConfig) *dataproc.JobSparkJobLoggingConfig {
	if p == nil {
		return nil
	}
	obj := &dataproc.JobSparkJobLoggingConfig{}
	return obj
}

// ProtoToJobPysparkJob converts a JobPysparkJob resource from its proto representation.
func ProtoToDataprocJobPysparkJob(p *dataprocpb.DataprocJobPysparkJob) *dataproc.JobPysparkJob {
	if p == nil {
		return nil
	}
	obj := &dataproc.JobPysparkJob{
		MainPythonFileUri: dcl.StringOrNil(p.MainPythonFileUri),
		LoggingConfig:     ProtoToDataprocJobPysparkJobLoggingConfig(p.GetLoggingConfig()),
	}
	for _, r := range p.GetArgs() {
		obj.Args = append(obj.Args, r)
	}
	for _, r := range p.GetPythonFileUris() {
		obj.PythonFileUris = append(obj.PythonFileUris, r)
	}
	for _, r := range p.GetJarFileUris() {
		obj.JarFileUris = append(obj.JarFileUris, r)
	}
	for _, r := range p.GetFileUris() {
		obj.FileUris = append(obj.FileUris, r)
	}
	for _, r := range p.GetArchiveUris() {
		obj.ArchiveUris = append(obj.ArchiveUris, r)
	}
	return obj
}

// ProtoToJobPysparkJobLoggingConfig converts a JobPysparkJobLoggingConfig resource from its proto representation.
func ProtoToDataprocJobPysparkJobLoggingConfig(p *dataprocpb.DataprocJobPysparkJobLoggingConfig) *dataproc.JobPysparkJobLoggingConfig {
	if p == nil {
		return nil
	}
	obj := &dataproc.JobPysparkJobLoggingConfig{}
	return obj
}

// ProtoToJobHiveJob converts a JobHiveJob resource from its proto representation.
func ProtoToDataprocJobHiveJob(p *dataprocpb.DataprocJobHiveJob) *dataproc.JobHiveJob {
	if p == nil {
		return nil
	}
	obj := &dataproc.JobHiveJob{
		QueryFileUri:      dcl.StringOrNil(p.QueryFileUri),
		QueryList:         ProtoToDataprocJobHiveJobQueryList(p.GetQueryList()),
		ContinueOnFailure: dcl.Bool(p.ContinueOnFailure),
	}
	for _, r := range p.GetJarFileUris() {
		obj.JarFileUris = append(obj.JarFileUris, r)
	}
	return obj
}

// ProtoToJobHiveJobQueryList converts a JobHiveJobQueryList resource from its proto representation.
func ProtoToDataprocJobHiveJobQueryList(p *dataprocpb.DataprocJobHiveJobQueryList) *dataproc.JobHiveJobQueryList {
	if p == nil {
		return nil
	}
	obj := &dataproc.JobHiveJobQueryList{}
	for _, r := range p.GetQueries() {
		obj.Queries = append(obj.Queries, r)
	}
	return obj
}

// ProtoToJobPigJob converts a JobPigJob resource from its proto representation.
func ProtoToDataprocJobPigJob(p *dataprocpb.DataprocJobPigJob) *dataproc.JobPigJob {
	if p == nil {
		return nil
	}
	obj := &dataproc.JobPigJob{
		QueryFileUri:      dcl.StringOrNil(p.QueryFileUri),
		QueryList:         ProtoToDataprocJobPigJobQueryList(p.GetQueryList()),
		ContinueOnFailure: dcl.Bool(p.ContinueOnFailure),
		LoggingConfig:     ProtoToDataprocJobPigJobLoggingConfig(p.GetLoggingConfig()),
	}
	for _, r := range p.GetJarFileUris() {
		obj.JarFileUris = append(obj.JarFileUris, r)
	}
	return obj
}

// ProtoToJobPigJobQueryList converts a JobPigJobQueryList resource from its proto representation.
func ProtoToDataprocJobPigJobQueryList(p *dataprocpb.DataprocJobPigJobQueryList) *dataproc.JobPigJobQueryList {
	if p == nil {
		return nil
	}
	obj := &dataproc.JobPigJobQueryList{}
	for _, r := range p.GetQueries() {
		obj.Queries = append(obj.Queries, r)
	}
	return obj
}

// ProtoToJobPigJobLoggingConfig converts a JobPigJobLoggingConfig resource from its proto representation.
func ProtoToDataprocJobPigJobLoggingConfig(p *dataprocpb.DataprocJobPigJobLoggingConfig) *dataproc.JobPigJobLoggingConfig {
	if p == nil {
		return nil
	}
	obj := &dataproc.JobPigJobLoggingConfig{}
	return obj
}

// ProtoToJobSparkRJob converts a JobSparkRJob resource from its proto representation.
func ProtoToDataprocJobSparkRJob(p *dataprocpb.DataprocJobSparkRJob) *dataproc.JobSparkRJob {
	if p == nil {
		return nil
	}
	obj := &dataproc.JobSparkRJob{
		MainRFileUri:  dcl.StringOrNil(p.MainRFileUri),
		LoggingConfig: ProtoToDataprocJobSparkRJobLoggingConfig(p.GetLoggingConfig()),
	}
	for _, r := range p.GetArgs() {
		obj.Args = append(obj.Args, r)
	}
	for _, r := range p.GetFileUris() {
		obj.FileUris = append(obj.FileUris, r)
	}
	for _, r := range p.GetArchiveUris() {
		obj.ArchiveUris = append(obj.ArchiveUris, r)
	}
	return obj
}

// ProtoToJobSparkRJobLoggingConfig converts a JobSparkRJobLoggingConfig resource from its proto representation.
func ProtoToDataprocJobSparkRJobLoggingConfig(p *dataprocpb.DataprocJobSparkRJobLoggingConfig) *dataproc.JobSparkRJobLoggingConfig {
	if p == nil {
		return nil
	}
	obj := &dataproc.JobSparkRJobLoggingConfig{}
	return obj
}

// ProtoToJobSparkSqlJob converts a JobSparkSqlJob resource from its proto representation.
func ProtoToDataprocJobSparkSqlJob(p *dataprocpb.DataprocJobSparkSqlJob) *dataproc.JobSparkSqlJob {
	if p == nil {
		return nil
	}
	obj := &dataproc.JobSparkSqlJob{
		QueryFileUri:  dcl.StringOrNil(p.QueryFileUri),
		QueryList:     ProtoToDataprocJobSparkSqlJobQueryList(p.GetQueryList()),
		LoggingConfig: ProtoToDataprocJobSparkSqlJobLoggingConfig(p.GetLoggingConfig()),
	}
	for _, r := range p.GetJarFileUris() {
		obj.JarFileUris = append(obj.JarFileUris, r)
	}
	return obj
}

// ProtoToJobSparkSqlJobQueryList converts a JobSparkSqlJobQueryList resource from its proto representation.
func ProtoToDataprocJobSparkSqlJobQueryList(p *dataprocpb.DataprocJobSparkSqlJobQueryList) *dataproc.JobSparkSqlJobQueryList {
	if p == nil {
		return nil
	}
	obj := &dataproc.JobSparkSqlJobQueryList{}
	for _, r := range p.GetQueries() {
		obj.Queries = append(obj.Queries, r)
	}
	return obj
}

// ProtoToJobSparkSqlJobLoggingConfig converts a JobSparkSqlJobLoggingConfig resource from its proto representation.
func ProtoToDataprocJobSparkSqlJobLoggingConfig(p *dataprocpb.DataprocJobSparkSqlJobLoggingConfig) *dataproc.JobSparkSqlJobLoggingConfig {
	if p == nil {
		return nil
	}
	obj := &dataproc.JobSparkSqlJobLoggingConfig{}
	return obj
}

// ProtoToJobPrestoJob converts a JobPrestoJob resource from its proto representation.
func ProtoToDataprocJobPrestoJob(p *dataprocpb.DataprocJobPrestoJob) *dataproc.JobPrestoJob {
	if p == nil {
		return nil
	}
	obj := &dataproc.JobPrestoJob{
		QueryFileUri:      dcl.StringOrNil(p.QueryFileUri),
		QueryList:         ProtoToDataprocJobPrestoJobQueryList(p.GetQueryList()),
		ContinueOnFailure: dcl.Bool(p.ContinueOnFailure),
		OutputFormat:      dcl.StringOrNil(p.OutputFormat),
		LoggingConfig:     ProtoToDataprocJobPrestoJobLoggingConfig(p.GetLoggingConfig()),
	}
	for _, r := range p.GetClientTags() {
		obj.ClientTags = append(obj.ClientTags, r)
	}
	return obj
}

// ProtoToJobPrestoJobQueryList converts a JobPrestoJobQueryList resource from its proto representation.
func ProtoToDataprocJobPrestoJobQueryList(p *dataprocpb.DataprocJobPrestoJobQueryList) *dataproc.JobPrestoJobQueryList {
	if p == nil {
		return nil
	}
	obj := &dataproc.JobPrestoJobQueryList{}
	for _, r := range p.GetQueries() {
		obj.Queries = append(obj.Queries, r)
	}
	return obj
}

// ProtoToJobPrestoJobLoggingConfig converts a JobPrestoJobLoggingConfig resource from its proto representation.
func ProtoToDataprocJobPrestoJobLoggingConfig(p *dataprocpb.DataprocJobPrestoJobLoggingConfig) *dataproc.JobPrestoJobLoggingConfig {
	if p == nil {
		return nil
	}
	obj := &dataproc.JobPrestoJobLoggingConfig{}
	return obj
}

// ProtoToJobStatus converts a JobStatus resource from its proto representation.
func ProtoToDataprocJobStatus(p *dataprocpb.DataprocJobStatus) *dataproc.JobStatus {
	if p == nil {
		return nil
	}
	obj := &dataproc.JobStatus{
		State:          ProtoToDataprocJobStatusStateEnum(p.GetState()),
		Details:        dcl.StringOrNil(p.Details),
		StateStartTime: dcl.StringOrNil(p.GetStateStartTime()),
		Substate:       ProtoToDataprocJobStatusSubstateEnum(p.GetSubstate()),
	}
	return obj
}

// ProtoToJobStatusHistory converts a JobStatusHistory resource from its proto representation.
func ProtoToDataprocJobStatusHistory(p *dataprocpb.DataprocJobStatusHistory) *dataproc.JobStatusHistory {
	if p == nil {
		return nil
	}
	obj := &dataproc.JobStatusHistory{
		State:          ProtoToDataprocJobStatusHistoryStateEnum(p.GetState()),
		Details:        dcl.StringOrNil(p.Details),
		StateStartTime: dcl.StringOrNil(p.GetStateStartTime()),
		Substate:       ProtoToDataprocJobStatusHistorySubstateEnum(p.GetSubstate()),
	}
	return obj
}

// ProtoToJobYarnApplications converts a JobYarnApplications resource from its proto representation.
func ProtoToDataprocJobYarnApplications(p *dataprocpb.DataprocJobYarnApplications) *dataproc.JobYarnApplications {
	if p == nil {
		return nil
	}
	obj := &dataproc.JobYarnApplications{
		Name:        dcl.StringOrNil(p.Name),
		State:       ProtoToDataprocJobYarnApplicationsStateEnum(p.GetState()),
		Progress:    dcl.Float64OrNil(p.Progress),
		TrackingUrl: dcl.StringOrNil(p.TrackingUrl),
	}
	return obj
}

// ProtoToJobScheduling converts a JobScheduling resource from its proto representation.
func ProtoToDataprocJobScheduling(p *dataprocpb.DataprocJobScheduling) *dataproc.JobScheduling {
	if p == nil {
		return nil
	}
	obj := &dataproc.JobScheduling{
		MaxFailuresPerHour: dcl.Int64OrNil(p.MaxFailuresPerHour),
		MaxFailuresTotal:   dcl.Int64OrNil(p.MaxFailuresTotal),
	}
	return obj
}

// ProtoToJob converts a Job resource from its proto representation.
func ProtoToJob(p *dataprocpb.DataprocJob) *dataproc.Job {
	obj := &dataproc.Job{
		Reference:               ProtoToDataprocJobReference(p.GetReference()),
		Placement:               ProtoToDataprocJobPlacement(p.GetPlacement()),
		HadoopJob:               ProtoToDataprocJobHadoopJob(p.GetHadoopJob()),
		SparkJob:                ProtoToDataprocJobSparkJob(p.GetSparkJob()),
		PysparkJob:              ProtoToDataprocJobPysparkJob(p.GetPysparkJob()),
		HiveJob:                 ProtoToDataprocJobHiveJob(p.GetHiveJob()),
		PigJob:                  ProtoToDataprocJobPigJob(p.GetPigJob()),
		SparkRJob:               ProtoToDataprocJobSparkRJob(p.GetSparkRJob()),
		SparkSqlJob:             ProtoToDataprocJobSparkSqlJob(p.GetSparkSqlJob()),
		PrestoJob:               ProtoToDataprocJobPrestoJob(p.GetPrestoJob()),
		Status:                  ProtoToDataprocJobStatus(p.GetStatus()),
		DriverOutputResourceUri: dcl.StringOrNil(p.DriverOutputResourceUri),
		DriverControlFilesUri:   dcl.StringOrNil(p.DriverControlFilesUri),
		Scheduling:              ProtoToDataprocJobScheduling(p.GetScheduling()),
		Name:                    dcl.StringOrNil(p.Name),
		Done:                    dcl.Bool(p.Done),
		Region:                  dcl.StringOrNil(p.Region),
		Project:                 dcl.StringOrNil(p.Project),
	}
	for _, r := range p.GetStatusHistory() {
		obj.StatusHistory = append(obj.StatusHistory, *ProtoToDataprocJobStatusHistory(r))
	}
	for _, r := range p.GetYarnApplications() {
		obj.YarnApplications = append(obj.YarnApplications, *ProtoToDataprocJobYarnApplications(r))
	}
	return obj
}

// JobStatusStateEnumToProto converts a JobStatusStateEnum enum to its proto representation.
func DataprocJobStatusStateEnumToProto(e *dataproc.JobStatusStateEnum) dataprocpb.DataprocJobStatusStateEnum {
	if e == nil {
		return dataprocpb.DataprocJobStatusStateEnum(0)
	}
	if v, ok := dataprocpb.DataprocJobStatusStateEnum_value["JobStatusStateEnum"+string(*e)]; ok {
		return dataprocpb.DataprocJobStatusStateEnum(v)
	}
	return dataprocpb.DataprocJobStatusStateEnum(0)
}

// JobStatusSubstateEnumToProto converts a JobStatusSubstateEnum enum to its proto representation.
func DataprocJobStatusSubstateEnumToProto(e *dataproc.JobStatusSubstateEnum) dataprocpb.DataprocJobStatusSubstateEnum {
	if e == nil {
		return dataprocpb.DataprocJobStatusSubstateEnum(0)
	}
	if v, ok := dataprocpb.DataprocJobStatusSubstateEnum_value["JobStatusSubstateEnum"+string(*e)]; ok {
		return dataprocpb.DataprocJobStatusSubstateEnum(v)
	}
	return dataprocpb.DataprocJobStatusSubstateEnum(0)
}

// JobStatusHistoryStateEnumToProto converts a JobStatusHistoryStateEnum enum to its proto representation.
func DataprocJobStatusHistoryStateEnumToProto(e *dataproc.JobStatusHistoryStateEnum) dataprocpb.DataprocJobStatusHistoryStateEnum {
	if e == nil {
		return dataprocpb.DataprocJobStatusHistoryStateEnum(0)
	}
	if v, ok := dataprocpb.DataprocJobStatusHistoryStateEnum_value["JobStatusHistoryStateEnum"+string(*e)]; ok {
		return dataprocpb.DataprocJobStatusHistoryStateEnum(v)
	}
	return dataprocpb.DataprocJobStatusHistoryStateEnum(0)
}

// JobStatusHistorySubstateEnumToProto converts a JobStatusHistorySubstateEnum enum to its proto representation.
func DataprocJobStatusHistorySubstateEnumToProto(e *dataproc.JobStatusHistorySubstateEnum) dataprocpb.DataprocJobStatusHistorySubstateEnum {
	if e == nil {
		return dataprocpb.DataprocJobStatusHistorySubstateEnum(0)
	}
	if v, ok := dataprocpb.DataprocJobStatusHistorySubstateEnum_value["JobStatusHistorySubstateEnum"+string(*e)]; ok {
		return dataprocpb.DataprocJobStatusHistorySubstateEnum(v)
	}
	return dataprocpb.DataprocJobStatusHistorySubstateEnum(0)
}

// JobYarnApplicationsStateEnumToProto converts a JobYarnApplicationsStateEnum enum to its proto representation.
func DataprocJobYarnApplicationsStateEnumToProto(e *dataproc.JobYarnApplicationsStateEnum) dataprocpb.DataprocJobYarnApplicationsStateEnum {
	if e == nil {
		return dataprocpb.DataprocJobYarnApplicationsStateEnum(0)
	}
	if v, ok := dataprocpb.DataprocJobYarnApplicationsStateEnum_value["JobYarnApplicationsStateEnum"+string(*e)]; ok {
		return dataprocpb.DataprocJobYarnApplicationsStateEnum(v)
	}
	return dataprocpb.DataprocJobYarnApplicationsStateEnum(0)
}

// JobReferenceToProto converts a JobReference resource to its proto representation.
func DataprocJobReferenceToProto(o *dataproc.JobReference) *dataprocpb.DataprocJobReference {
	if o == nil {
		return nil
	}
	p := &dataprocpb.DataprocJobReference{
		ProjectId: dcl.ValueOrEmptyString(o.ProjectId),
		JobId:     dcl.ValueOrEmptyString(o.JobId),
	}
	return p
}

// JobPlacementToProto converts a JobPlacement resource to its proto representation.
func DataprocJobPlacementToProto(o *dataproc.JobPlacement) *dataprocpb.DataprocJobPlacement {
	if o == nil {
		return nil
	}
	p := &dataprocpb.DataprocJobPlacement{
		ClusterName: dcl.ValueOrEmptyString(o.ClusterName),
		ClusterUuid: dcl.ValueOrEmptyString(o.ClusterUuid),
	}
	p.ClusterLabels = make(map[string]string)
	for k, r := range o.ClusterLabels {
		p.ClusterLabels[k] = r
	}
	return p
}

// JobHadoopJobToProto converts a JobHadoopJob resource to its proto representation.
func DataprocJobHadoopJobToProto(o *dataproc.JobHadoopJob) *dataprocpb.DataprocJobHadoopJob {
	if o == nil {
		return nil
	}
	p := &dataprocpb.DataprocJobHadoopJob{
		MainJarFileUri: dcl.ValueOrEmptyString(o.MainJarFileUri),
		MainClass:      dcl.ValueOrEmptyString(o.MainClass),
		LoggingConfig:  DataprocJobHadoopJobLoggingConfigToProto(o.LoggingConfig),
	}
	for _, r := range o.Args {
		p.Args = append(p.Args, r)
	}
	for _, r := range o.JarFileUris {
		p.JarFileUris = append(p.JarFileUris, r)
	}
	for _, r := range o.FileUris {
		p.FileUris = append(p.FileUris, r)
	}
	for _, r := range o.ArchiveUris {
		p.ArchiveUris = append(p.ArchiveUris, r)
	}
	p.Properties = make(map[string]string)
	for k, r := range o.Properties {
		p.Properties[k] = r
	}
	return p
}

// JobHadoopJobLoggingConfigToProto converts a JobHadoopJobLoggingConfig resource to its proto representation.
func DataprocJobHadoopJobLoggingConfigToProto(o *dataproc.JobHadoopJobLoggingConfig) *dataprocpb.DataprocJobHadoopJobLoggingConfig {
	if o == nil {
		return nil
	}
	p := &dataprocpb.DataprocJobHadoopJobLoggingConfig{}
	p.DriverLogLevels = make(map[string]string)
	for k, r := range o.DriverLogLevels {
		p.DriverLogLevels[k] = r
	}
	return p
}

// JobSparkJobToProto converts a JobSparkJob resource to its proto representation.
func DataprocJobSparkJobToProto(o *dataproc.JobSparkJob) *dataprocpb.DataprocJobSparkJob {
	if o == nil {
		return nil
	}
	p := &dataprocpb.DataprocJobSparkJob{
		MainJarFileUri: dcl.ValueOrEmptyString(o.MainJarFileUri),
		MainClass:      dcl.ValueOrEmptyString(o.MainClass),
		LoggingConfig:  DataprocJobSparkJobLoggingConfigToProto(o.LoggingConfig),
	}
	for _, r := range o.Args {
		p.Args = append(p.Args, r)
	}
	for _, r := range o.JarFileUris {
		p.JarFileUris = append(p.JarFileUris, r)
	}
	for _, r := range o.FileUris {
		p.FileUris = append(p.FileUris, r)
	}
	for _, r := range o.ArchiveUris {
		p.ArchiveUris = append(p.ArchiveUris, r)
	}
	p.Properties = make(map[string]string)
	for k, r := range o.Properties {
		p.Properties[k] = r
	}
	return p
}

// JobSparkJobLoggingConfigToProto converts a JobSparkJobLoggingConfig resource to its proto representation.
func DataprocJobSparkJobLoggingConfigToProto(o *dataproc.JobSparkJobLoggingConfig) *dataprocpb.DataprocJobSparkJobLoggingConfig {
	if o == nil {
		return nil
	}
	p := &dataprocpb.DataprocJobSparkJobLoggingConfig{}
	p.DriverLogLevels = make(map[string]string)
	for k, r := range o.DriverLogLevels {
		p.DriverLogLevels[k] = r
	}
	return p
}

// JobPysparkJobToProto converts a JobPysparkJob resource to its proto representation.
func DataprocJobPysparkJobToProto(o *dataproc.JobPysparkJob) *dataprocpb.DataprocJobPysparkJob {
	if o == nil {
		return nil
	}
	p := &dataprocpb.DataprocJobPysparkJob{
		MainPythonFileUri: dcl.ValueOrEmptyString(o.MainPythonFileUri),
		LoggingConfig:     DataprocJobPysparkJobLoggingConfigToProto(o.LoggingConfig),
	}
	for _, r := range o.Args {
		p.Args = append(p.Args, r)
	}
	for _, r := range o.PythonFileUris {
		p.PythonFileUris = append(p.PythonFileUris, r)
	}
	for _, r := range o.JarFileUris {
		p.JarFileUris = append(p.JarFileUris, r)
	}
	for _, r := range o.FileUris {
		p.FileUris = append(p.FileUris, r)
	}
	for _, r := range o.ArchiveUris {
		p.ArchiveUris = append(p.ArchiveUris, r)
	}
	p.Properties = make(map[string]string)
	for k, r := range o.Properties {
		p.Properties[k] = r
	}
	return p
}

// JobPysparkJobLoggingConfigToProto converts a JobPysparkJobLoggingConfig resource to its proto representation.
func DataprocJobPysparkJobLoggingConfigToProto(o *dataproc.JobPysparkJobLoggingConfig) *dataprocpb.DataprocJobPysparkJobLoggingConfig {
	if o == nil {
		return nil
	}
	p := &dataprocpb.DataprocJobPysparkJobLoggingConfig{}
	p.DriverLogLevels = make(map[string]string)
	for k, r := range o.DriverLogLevels {
		p.DriverLogLevels[k] = r
	}
	return p
}

// JobHiveJobToProto converts a JobHiveJob resource to its proto representation.
func DataprocJobHiveJobToProto(o *dataproc.JobHiveJob) *dataprocpb.DataprocJobHiveJob {
	if o == nil {
		return nil
	}
	p := &dataprocpb.DataprocJobHiveJob{
		QueryFileUri:      dcl.ValueOrEmptyString(o.QueryFileUri),
		QueryList:         DataprocJobHiveJobQueryListToProto(o.QueryList),
		ContinueOnFailure: dcl.ValueOrEmptyBool(o.ContinueOnFailure),
	}
	p.ScriptVariables = make(map[string]string)
	for k, r := range o.ScriptVariables {
		p.ScriptVariables[k] = r
	}
	p.Properties = make(map[string]string)
	for k, r := range o.Properties {
		p.Properties[k] = r
	}
	for _, r := range o.JarFileUris {
		p.JarFileUris = append(p.JarFileUris, r)
	}
	return p
}

// JobHiveJobQueryListToProto converts a JobHiveJobQueryList resource to its proto representation.
func DataprocJobHiveJobQueryListToProto(o *dataproc.JobHiveJobQueryList) *dataprocpb.DataprocJobHiveJobQueryList {
	if o == nil {
		return nil
	}
	p := &dataprocpb.DataprocJobHiveJobQueryList{}
	for _, r := range o.Queries {
		p.Queries = append(p.Queries, r)
	}
	return p
}

// JobPigJobToProto converts a JobPigJob resource to its proto representation.
func DataprocJobPigJobToProto(o *dataproc.JobPigJob) *dataprocpb.DataprocJobPigJob {
	if o == nil {
		return nil
	}
	p := &dataprocpb.DataprocJobPigJob{
		QueryFileUri:      dcl.ValueOrEmptyString(o.QueryFileUri),
		QueryList:         DataprocJobPigJobQueryListToProto(o.QueryList),
		ContinueOnFailure: dcl.ValueOrEmptyBool(o.ContinueOnFailure),
		LoggingConfig:     DataprocJobPigJobLoggingConfigToProto(o.LoggingConfig),
	}
	p.ScriptVariables = make(map[string]string)
	for k, r := range o.ScriptVariables {
		p.ScriptVariables[k] = r
	}
	p.Properties = make(map[string]string)
	for k, r := range o.Properties {
		p.Properties[k] = r
	}
	for _, r := range o.JarFileUris {
		p.JarFileUris = append(p.JarFileUris, r)
	}
	return p
}

// JobPigJobQueryListToProto converts a JobPigJobQueryList resource to its proto representation.
func DataprocJobPigJobQueryListToProto(o *dataproc.JobPigJobQueryList) *dataprocpb.DataprocJobPigJobQueryList {
	if o == nil {
		return nil
	}
	p := &dataprocpb.DataprocJobPigJobQueryList{}
	for _, r := range o.Queries {
		p.Queries = append(p.Queries, r)
	}
	return p
}

// JobPigJobLoggingConfigToProto converts a JobPigJobLoggingConfig resource to its proto representation.
func DataprocJobPigJobLoggingConfigToProto(o *dataproc.JobPigJobLoggingConfig) *dataprocpb.DataprocJobPigJobLoggingConfig {
	if o == nil {
		return nil
	}
	p := &dataprocpb.DataprocJobPigJobLoggingConfig{}
	p.DriverLogLevels = make(map[string]string)
	for k, r := range o.DriverLogLevels {
		p.DriverLogLevels[k] = r
	}
	return p
}

// JobSparkRJobToProto converts a JobSparkRJob resource to its proto representation.
func DataprocJobSparkRJobToProto(o *dataproc.JobSparkRJob) *dataprocpb.DataprocJobSparkRJob {
	if o == nil {
		return nil
	}
	p := &dataprocpb.DataprocJobSparkRJob{
		MainRFileUri:  dcl.ValueOrEmptyString(o.MainRFileUri),
		LoggingConfig: DataprocJobSparkRJobLoggingConfigToProto(o.LoggingConfig),
	}
	for _, r := range o.Args {
		p.Args = append(p.Args, r)
	}
	for _, r := range o.FileUris {
		p.FileUris = append(p.FileUris, r)
	}
	for _, r := range o.ArchiveUris {
		p.ArchiveUris = append(p.ArchiveUris, r)
	}
	p.Properties = make(map[string]string)
	for k, r := range o.Properties {
		p.Properties[k] = r
	}
	return p
}

// JobSparkRJobLoggingConfigToProto converts a JobSparkRJobLoggingConfig resource to its proto representation.
func DataprocJobSparkRJobLoggingConfigToProto(o *dataproc.JobSparkRJobLoggingConfig) *dataprocpb.DataprocJobSparkRJobLoggingConfig {
	if o == nil {
		return nil
	}
	p := &dataprocpb.DataprocJobSparkRJobLoggingConfig{}
	p.DriverLogLevels = make(map[string]string)
	for k, r := range o.DriverLogLevels {
		p.DriverLogLevels[k] = r
	}
	return p
}

// JobSparkSqlJobToProto converts a JobSparkSqlJob resource to its proto representation.
func DataprocJobSparkSqlJobToProto(o *dataproc.JobSparkSqlJob) *dataprocpb.DataprocJobSparkSqlJob {
	if o == nil {
		return nil
	}
	p := &dataprocpb.DataprocJobSparkSqlJob{
		QueryFileUri:  dcl.ValueOrEmptyString(o.QueryFileUri),
		QueryList:     DataprocJobSparkSqlJobQueryListToProto(o.QueryList),
		LoggingConfig: DataprocJobSparkSqlJobLoggingConfigToProto(o.LoggingConfig),
	}
	p.ScriptVariables = make(map[string]string)
	for k, r := range o.ScriptVariables {
		p.ScriptVariables[k] = r
	}
	p.Properties = make(map[string]string)
	for k, r := range o.Properties {
		p.Properties[k] = r
	}
	for _, r := range o.JarFileUris {
		p.JarFileUris = append(p.JarFileUris, r)
	}
	return p
}

// JobSparkSqlJobQueryListToProto converts a JobSparkSqlJobQueryList resource to its proto representation.
func DataprocJobSparkSqlJobQueryListToProto(o *dataproc.JobSparkSqlJobQueryList) *dataprocpb.DataprocJobSparkSqlJobQueryList {
	if o == nil {
		return nil
	}
	p := &dataprocpb.DataprocJobSparkSqlJobQueryList{}
	for _, r := range o.Queries {
		p.Queries = append(p.Queries, r)
	}
	return p
}

// JobSparkSqlJobLoggingConfigToProto converts a JobSparkSqlJobLoggingConfig resource to its proto representation.
func DataprocJobSparkSqlJobLoggingConfigToProto(o *dataproc.JobSparkSqlJobLoggingConfig) *dataprocpb.DataprocJobSparkSqlJobLoggingConfig {
	if o == nil {
		return nil
	}
	p := &dataprocpb.DataprocJobSparkSqlJobLoggingConfig{}
	p.DriverLogLevels = make(map[string]string)
	for k, r := range o.DriverLogLevels {
		p.DriverLogLevels[k] = r
	}
	return p
}

// JobPrestoJobToProto converts a JobPrestoJob resource to its proto representation.
func DataprocJobPrestoJobToProto(o *dataproc.JobPrestoJob) *dataprocpb.DataprocJobPrestoJob {
	if o == nil {
		return nil
	}
	p := &dataprocpb.DataprocJobPrestoJob{
		QueryFileUri:      dcl.ValueOrEmptyString(o.QueryFileUri),
		QueryList:         DataprocJobPrestoJobQueryListToProto(o.QueryList),
		ContinueOnFailure: dcl.ValueOrEmptyBool(o.ContinueOnFailure),
		OutputFormat:      dcl.ValueOrEmptyString(o.OutputFormat),
		LoggingConfig:     DataprocJobPrestoJobLoggingConfigToProto(o.LoggingConfig),
	}
	for _, r := range o.ClientTags {
		p.ClientTags = append(p.ClientTags, r)
	}
	p.Properties = make(map[string]string)
	for k, r := range o.Properties {
		p.Properties[k] = r
	}
	return p
}

// JobPrestoJobQueryListToProto converts a JobPrestoJobQueryList resource to its proto representation.
func DataprocJobPrestoJobQueryListToProto(o *dataproc.JobPrestoJobQueryList) *dataprocpb.DataprocJobPrestoJobQueryList {
	if o == nil {
		return nil
	}
	p := &dataprocpb.DataprocJobPrestoJobQueryList{}
	for _, r := range o.Queries {
		p.Queries = append(p.Queries, r)
	}
	return p
}

// JobPrestoJobLoggingConfigToProto converts a JobPrestoJobLoggingConfig resource to its proto representation.
func DataprocJobPrestoJobLoggingConfigToProto(o *dataproc.JobPrestoJobLoggingConfig) *dataprocpb.DataprocJobPrestoJobLoggingConfig {
	if o == nil {
		return nil
	}
	p := &dataprocpb.DataprocJobPrestoJobLoggingConfig{}
	p.DriverLogLevels = make(map[string]string)
	for k, r := range o.DriverLogLevels {
		p.DriverLogLevels[k] = r
	}
	return p
}

// JobStatusToProto converts a JobStatus resource to its proto representation.
func DataprocJobStatusToProto(o *dataproc.JobStatus) *dataprocpb.DataprocJobStatus {
	if o == nil {
		return nil
	}
	p := &dataprocpb.DataprocJobStatus{
		State:          DataprocJobStatusStateEnumToProto(o.State),
		Details:        dcl.ValueOrEmptyString(o.Details),
		StateStartTime: dcl.ValueOrEmptyString(o.StateStartTime),
		Substate:       DataprocJobStatusSubstateEnumToProto(o.Substate),
	}
	return p
}

// JobStatusHistoryToProto converts a JobStatusHistory resource to its proto representation.
func DataprocJobStatusHistoryToProto(o *dataproc.JobStatusHistory) *dataprocpb.DataprocJobStatusHistory {
	if o == nil {
		return nil
	}
	p := &dataprocpb.DataprocJobStatusHistory{
		State:          DataprocJobStatusHistoryStateEnumToProto(o.State),
		Details:        dcl.ValueOrEmptyString(o.Details),
		StateStartTime: dcl.ValueOrEmptyString(o.StateStartTime),
		Substate:       DataprocJobStatusHistorySubstateEnumToProto(o.Substate),
	}
	return p
}

// JobYarnApplicationsToProto converts a JobYarnApplications resource to its proto representation.
func DataprocJobYarnApplicationsToProto(o *dataproc.JobYarnApplications) *dataprocpb.DataprocJobYarnApplications {
	if o == nil {
		return nil
	}
	p := &dataprocpb.DataprocJobYarnApplications{
		Name:        dcl.ValueOrEmptyString(o.Name),
		State:       DataprocJobYarnApplicationsStateEnumToProto(o.State),
		Progress:    dcl.ValueOrEmptyDouble(o.Progress),
		TrackingUrl: dcl.ValueOrEmptyString(o.TrackingUrl),
	}
	return p
}

// JobSchedulingToProto converts a JobScheduling resource to its proto representation.
func DataprocJobSchedulingToProto(o *dataproc.JobScheduling) *dataprocpb.DataprocJobScheduling {
	if o == nil {
		return nil
	}
	p := &dataprocpb.DataprocJobScheduling{
		MaxFailuresPerHour: dcl.ValueOrEmptyInt64(o.MaxFailuresPerHour),
		MaxFailuresTotal:   dcl.ValueOrEmptyInt64(o.MaxFailuresTotal),
	}
	return p
}

// JobToProto converts a Job resource to its proto representation.
func JobToProto(resource *dataproc.Job) *dataprocpb.DataprocJob {
	p := &dataprocpb.DataprocJob{
		Reference:               DataprocJobReferenceToProto(resource.Reference),
		Placement:               DataprocJobPlacementToProto(resource.Placement),
		HadoopJob:               DataprocJobHadoopJobToProto(resource.HadoopJob),
		SparkJob:                DataprocJobSparkJobToProto(resource.SparkJob),
		PysparkJob:              DataprocJobPysparkJobToProto(resource.PysparkJob),
		HiveJob:                 DataprocJobHiveJobToProto(resource.HiveJob),
		PigJob:                  DataprocJobPigJobToProto(resource.PigJob),
		SparkRJob:               DataprocJobSparkRJobToProto(resource.SparkRJob),
		SparkSqlJob:             DataprocJobSparkSqlJobToProto(resource.SparkSqlJob),
		PrestoJob:               DataprocJobPrestoJobToProto(resource.PrestoJob),
		Status:                  DataprocJobStatusToProto(resource.Status),
		DriverOutputResourceUri: dcl.ValueOrEmptyString(resource.DriverOutputResourceUri),
		DriverControlFilesUri:   dcl.ValueOrEmptyString(resource.DriverControlFilesUri),
		Scheduling:              DataprocJobSchedulingToProto(resource.Scheduling),
		Name:                    dcl.ValueOrEmptyString(resource.Name),
		Done:                    dcl.ValueOrEmptyBool(resource.Done),
		Region:                  dcl.ValueOrEmptyString(resource.Region),
		Project:                 dcl.ValueOrEmptyString(resource.Project),
	}
	for _, r := range resource.StatusHistory {
		p.StatusHistory = append(p.StatusHistory, DataprocJobStatusHistoryToProto(&r))
	}
	for _, r := range resource.YarnApplications {
		p.YarnApplications = append(p.YarnApplications, DataprocJobYarnApplicationsToProto(&r))
	}

	return p
}

// ApplyJob handles the gRPC request by passing it to the underlying Job Apply() method.
func (s *JobServer) applyJob(ctx context.Context, c *dataproc.Client, request *dataprocpb.ApplyDataprocJobRequest) (*dataprocpb.DataprocJob, error) {
	p := ProtoToJob(request.GetResource())
	res, err := c.ApplyJob(ctx, p)
	if err != nil {
		return nil, err
	}
	r := JobToProto(res)
	return r, nil
}

// ApplyJob handles the gRPC request by passing it to the underlying Job Apply() method.
func (s *JobServer) ApplyDataprocJob(ctx context.Context, request *dataprocpb.ApplyDataprocJobRequest) (*dataprocpb.DataprocJob, error) {
	cl, err := createConfigJob(ctx, request.ServiceAccountFile)
	if err != nil {
		return nil, err
	}
	return s.applyJob(ctx, cl, request)
}

// DeleteJob handles the gRPC request by passing it to the underlying Job Delete() method.
func (s *JobServer) DeleteDataprocJob(ctx context.Context, request *dataprocpb.DeleteDataprocJobRequest) (*emptypb.Empty, error) {

	cl, err := createConfigJob(ctx, request.ServiceAccountFile)
	if err != nil {
		return nil, err
	}
	return &emptypb.Empty{}, cl.DeleteJob(ctx, ProtoToJob(request.GetResource()))

}

// ListDataprocJob handles the gRPC request by passing it to the underlying JobList() method.
func (s *JobServer) ListDataprocJob(ctx context.Context, request *dataprocpb.ListDataprocJobRequest) (*dataprocpb.ListDataprocJobResponse, error) {
	cl, err := createConfigJob(ctx, request.ServiceAccountFile)
	if err != nil {
		return nil, err
	}

	resources, err := cl.ListJob(ctx, request.Project, request.Region)
	if err != nil {
		return nil, err
	}
	var protos []*dataprocpb.DataprocJob
	for _, r := range resources.Items {
		rp := JobToProto(r)
		protos = append(protos, rp)
	}
	return &dataprocpb.ListDataprocJobResponse{Items: protos}, nil
}

func createConfigJob(ctx context.Context, service_account_file string) (*dataproc.Client, error) {

	conf := dcl.NewConfig(dcl.WithUserAgent("dcl-test"), dcl.WithCredentialsFile(service_account_file))
	return dataproc.NewClient(conf), nil
}
