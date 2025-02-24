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
	betapb "github.com/GoogleCloudPlatform/declarative-resource-client-library/python/proto/dataproc/beta/dataproc_beta_go_proto"
	emptypb "github.com/GoogleCloudPlatform/declarative-resource-client-library/python/proto/empty_go_proto"
	"github.com/GoogleCloudPlatform/declarative-resource-client-library/services/google/dataproc/beta"
)

// Server implements the gRPC interface for Job.
type JobServer struct{}

// ProtoToJobStatusStateEnum converts a JobStatusStateEnum enum from its proto representation.
func ProtoToDataprocBetaJobStatusStateEnum(e betapb.DataprocBetaJobStatusStateEnum) *beta.JobStatusStateEnum {
	if e == 0 {
		return nil
	}
	if n, ok := betapb.DataprocBetaJobStatusStateEnum_name[int32(e)]; ok {
		e := beta.JobStatusStateEnum(n[len("DataprocBetaJobStatusStateEnum"):])
		return &e
	}
	return nil
}

// ProtoToJobStatusSubstateEnum converts a JobStatusSubstateEnum enum from its proto representation.
func ProtoToDataprocBetaJobStatusSubstateEnum(e betapb.DataprocBetaJobStatusSubstateEnum) *beta.JobStatusSubstateEnum {
	if e == 0 {
		return nil
	}
	if n, ok := betapb.DataprocBetaJobStatusSubstateEnum_name[int32(e)]; ok {
		e := beta.JobStatusSubstateEnum(n[len("DataprocBetaJobStatusSubstateEnum"):])
		return &e
	}
	return nil
}

// ProtoToJobStatusHistoryStateEnum converts a JobStatusHistoryStateEnum enum from its proto representation.
func ProtoToDataprocBetaJobStatusHistoryStateEnum(e betapb.DataprocBetaJobStatusHistoryStateEnum) *beta.JobStatusHistoryStateEnum {
	if e == 0 {
		return nil
	}
	if n, ok := betapb.DataprocBetaJobStatusHistoryStateEnum_name[int32(e)]; ok {
		e := beta.JobStatusHistoryStateEnum(n[len("DataprocBetaJobStatusHistoryStateEnum"):])
		return &e
	}
	return nil
}

// ProtoToJobStatusHistorySubstateEnum converts a JobStatusHistorySubstateEnum enum from its proto representation.
func ProtoToDataprocBetaJobStatusHistorySubstateEnum(e betapb.DataprocBetaJobStatusHistorySubstateEnum) *beta.JobStatusHistorySubstateEnum {
	if e == 0 {
		return nil
	}
	if n, ok := betapb.DataprocBetaJobStatusHistorySubstateEnum_name[int32(e)]; ok {
		e := beta.JobStatusHistorySubstateEnum(n[len("DataprocBetaJobStatusHistorySubstateEnum"):])
		return &e
	}
	return nil
}

// ProtoToJobYarnApplicationsStateEnum converts a JobYarnApplicationsStateEnum enum from its proto representation.
func ProtoToDataprocBetaJobYarnApplicationsStateEnum(e betapb.DataprocBetaJobYarnApplicationsStateEnum) *beta.JobYarnApplicationsStateEnum {
	if e == 0 {
		return nil
	}
	if n, ok := betapb.DataprocBetaJobYarnApplicationsStateEnum_name[int32(e)]; ok {
		e := beta.JobYarnApplicationsStateEnum(n[len("DataprocBetaJobYarnApplicationsStateEnum"):])
		return &e
	}
	return nil
}

// ProtoToJobReference converts a JobReference resource from its proto representation.
func ProtoToDataprocBetaJobReference(p *betapb.DataprocBetaJobReference) *beta.JobReference {
	if p == nil {
		return nil
	}
	obj := &beta.JobReference{
		ProjectId: dcl.StringOrNil(p.ProjectId),
		JobId:     dcl.StringOrNil(p.JobId),
	}
	return obj
}

// ProtoToJobPlacement converts a JobPlacement resource from its proto representation.
func ProtoToDataprocBetaJobPlacement(p *betapb.DataprocBetaJobPlacement) *beta.JobPlacement {
	if p == nil {
		return nil
	}
	obj := &beta.JobPlacement{
		ClusterName: dcl.StringOrNil(p.ClusterName),
		ClusterUuid: dcl.StringOrNil(p.ClusterUuid),
	}
	return obj
}

// ProtoToJobHadoopJob converts a JobHadoopJob resource from its proto representation.
func ProtoToDataprocBetaJobHadoopJob(p *betapb.DataprocBetaJobHadoopJob) *beta.JobHadoopJob {
	if p == nil {
		return nil
	}
	obj := &beta.JobHadoopJob{
		MainJarFileUri: dcl.StringOrNil(p.MainJarFileUri),
		MainClass:      dcl.StringOrNil(p.MainClass),
		LoggingConfig:  ProtoToDataprocBetaJobHadoopJobLoggingConfig(p.GetLoggingConfig()),
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
func ProtoToDataprocBetaJobHadoopJobLoggingConfig(p *betapb.DataprocBetaJobHadoopJobLoggingConfig) *beta.JobHadoopJobLoggingConfig {
	if p == nil {
		return nil
	}
	obj := &beta.JobHadoopJobLoggingConfig{}
	return obj
}

// ProtoToJobSparkJob converts a JobSparkJob resource from its proto representation.
func ProtoToDataprocBetaJobSparkJob(p *betapb.DataprocBetaJobSparkJob) *beta.JobSparkJob {
	if p == nil {
		return nil
	}
	obj := &beta.JobSparkJob{
		MainJarFileUri: dcl.StringOrNil(p.MainJarFileUri),
		MainClass:      dcl.StringOrNil(p.MainClass),
		LoggingConfig:  ProtoToDataprocBetaJobSparkJobLoggingConfig(p.GetLoggingConfig()),
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
func ProtoToDataprocBetaJobSparkJobLoggingConfig(p *betapb.DataprocBetaJobSparkJobLoggingConfig) *beta.JobSparkJobLoggingConfig {
	if p == nil {
		return nil
	}
	obj := &beta.JobSparkJobLoggingConfig{}
	return obj
}

// ProtoToJobPysparkJob converts a JobPysparkJob resource from its proto representation.
func ProtoToDataprocBetaJobPysparkJob(p *betapb.DataprocBetaJobPysparkJob) *beta.JobPysparkJob {
	if p == nil {
		return nil
	}
	obj := &beta.JobPysparkJob{
		MainPythonFileUri: dcl.StringOrNil(p.MainPythonFileUri),
		LoggingConfig:     ProtoToDataprocBetaJobPysparkJobLoggingConfig(p.GetLoggingConfig()),
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
func ProtoToDataprocBetaJobPysparkJobLoggingConfig(p *betapb.DataprocBetaJobPysparkJobLoggingConfig) *beta.JobPysparkJobLoggingConfig {
	if p == nil {
		return nil
	}
	obj := &beta.JobPysparkJobLoggingConfig{}
	return obj
}

// ProtoToJobHiveJob converts a JobHiveJob resource from its proto representation.
func ProtoToDataprocBetaJobHiveJob(p *betapb.DataprocBetaJobHiveJob) *beta.JobHiveJob {
	if p == nil {
		return nil
	}
	obj := &beta.JobHiveJob{
		QueryFileUri:      dcl.StringOrNil(p.QueryFileUri),
		QueryList:         ProtoToDataprocBetaJobHiveJobQueryList(p.GetQueryList()),
		ContinueOnFailure: dcl.Bool(p.ContinueOnFailure),
	}
	for _, r := range p.GetJarFileUris() {
		obj.JarFileUris = append(obj.JarFileUris, r)
	}
	return obj
}

// ProtoToJobHiveJobQueryList converts a JobHiveJobQueryList resource from its proto representation.
func ProtoToDataprocBetaJobHiveJobQueryList(p *betapb.DataprocBetaJobHiveJobQueryList) *beta.JobHiveJobQueryList {
	if p == nil {
		return nil
	}
	obj := &beta.JobHiveJobQueryList{}
	for _, r := range p.GetQueries() {
		obj.Queries = append(obj.Queries, r)
	}
	return obj
}

// ProtoToJobPigJob converts a JobPigJob resource from its proto representation.
func ProtoToDataprocBetaJobPigJob(p *betapb.DataprocBetaJobPigJob) *beta.JobPigJob {
	if p == nil {
		return nil
	}
	obj := &beta.JobPigJob{
		QueryFileUri:      dcl.StringOrNil(p.QueryFileUri),
		QueryList:         ProtoToDataprocBetaJobPigJobQueryList(p.GetQueryList()),
		ContinueOnFailure: dcl.Bool(p.ContinueOnFailure),
		LoggingConfig:     ProtoToDataprocBetaJobPigJobLoggingConfig(p.GetLoggingConfig()),
	}
	for _, r := range p.GetJarFileUris() {
		obj.JarFileUris = append(obj.JarFileUris, r)
	}
	return obj
}

// ProtoToJobPigJobQueryList converts a JobPigJobQueryList resource from its proto representation.
func ProtoToDataprocBetaJobPigJobQueryList(p *betapb.DataprocBetaJobPigJobQueryList) *beta.JobPigJobQueryList {
	if p == nil {
		return nil
	}
	obj := &beta.JobPigJobQueryList{}
	for _, r := range p.GetQueries() {
		obj.Queries = append(obj.Queries, r)
	}
	return obj
}

// ProtoToJobPigJobLoggingConfig converts a JobPigJobLoggingConfig resource from its proto representation.
func ProtoToDataprocBetaJobPigJobLoggingConfig(p *betapb.DataprocBetaJobPigJobLoggingConfig) *beta.JobPigJobLoggingConfig {
	if p == nil {
		return nil
	}
	obj := &beta.JobPigJobLoggingConfig{}
	return obj
}

// ProtoToJobSparkRJob converts a JobSparkRJob resource from its proto representation.
func ProtoToDataprocBetaJobSparkRJob(p *betapb.DataprocBetaJobSparkRJob) *beta.JobSparkRJob {
	if p == nil {
		return nil
	}
	obj := &beta.JobSparkRJob{
		MainRFileUri:  dcl.StringOrNil(p.MainRFileUri),
		LoggingConfig: ProtoToDataprocBetaJobSparkRJobLoggingConfig(p.GetLoggingConfig()),
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
func ProtoToDataprocBetaJobSparkRJobLoggingConfig(p *betapb.DataprocBetaJobSparkRJobLoggingConfig) *beta.JobSparkRJobLoggingConfig {
	if p == nil {
		return nil
	}
	obj := &beta.JobSparkRJobLoggingConfig{}
	return obj
}

// ProtoToJobSparkSqlJob converts a JobSparkSqlJob resource from its proto representation.
func ProtoToDataprocBetaJobSparkSqlJob(p *betapb.DataprocBetaJobSparkSqlJob) *beta.JobSparkSqlJob {
	if p == nil {
		return nil
	}
	obj := &beta.JobSparkSqlJob{
		QueryFileUri:  dcl.StringOrNil(p.QueryFileUri),
		QueryList:     ProtoToDataprocBetaJobSparkSqlJobQueryList(p.GetQueryList()),
		LoggingConfig: ProtoToDataprocBetaJobSparkSqlJobLoggingConfig(p.GetLoggingConfig()),
	}
	for _, r := range p.GetJarFileUris() {
		obj.JarFileUris = append(obj.JarFileUris, r)
	}
	return obj
}

// ProtoToJobSparkSqlJobQueryList converts a JobSparkSqlJobQueryList resource from its proto representation.
func ProtoToDataprocBetaJobSparkSqlJobQueryList(p *betapb.DataprocBetaJobSparkSqlJobQueryList) *beta.JobSparkSqlJobQueryList {
	if p == nil {
		return nil
	}
	obj := &beta.JobSparkSqlJobQueryList{}
	for _, r := range p.GetQueries() {
		obj.Queries = append(obj.Queries, r)
	}
	return obj
}

// ProtoToJobSparkSqlJobLoggingConfig converts a JobSparkSqlJobLoggingConfig resource from its proto representation.
func ProtoToDataprocBetaJobSparkSqlJobLoggingConfig(p *betapb.DataprocBetaJobSparkSqlJobLoggingConfig) *beta.JobSparkSqlJobLoggingConfig {
	if p == nil {
		return nil
	}
	obj := &beta.JobSparkSqlJobLoggingConfig{}
	return obj
}

// ProtoToJobPrestoJob converts a JobPrestoJob resource from its proto representation.
func ProtoToDataprocBetaJobPrestoJob(p *betapb.DataprocBetaJobPrestoJob) *beta.JobPrestoJob {
	if p == nil {
		return nil
	}
	obj := &beta.JobPrestoJob{
		QueryFileUri:      dcl.StringOrNil(p.QueryFileUri),
		QueryList:         ProtoToDataprocBetaJobPrestoJobQueryList(p.GetQueryList()),
		ContinueOnFailure: dcl.Bool(p.ContinueOnFailure),
		OutputFormat:      dcl.StringOrNil(p.OutputFormat),
		LoggingConfig:     ProtoToDataprocBetaJobPrestoJobLoggingConfig(p.GetLoggingConfig()),
	}
	for _, r := range p.GetClientTags() {
		obj.ClientTags = append(obj.ClientTags, r)
	}
	return obj
}

// ProtoToJobPrestoJobQueryList converts a JobPrestoJobQueryList resource from its proto representation.
func ProtoToDataprocBetaJobPrestoJobQueryList(p *betapb.DataprocBetaJobPrestoJobQueryList) *beta.JobPrestoJobQueryList {
	if p == nil {
		return nil
	}
	obj := &beta.JobPrestoJobQueryList{}
	for _, r := range p.GetQueries() {
		obj.Queries = append(obj.Queries, r)
	}
	return obj
}

// ProtoToJobPrestoJobLoggingConfig converts a JobPrestoJobLoggingConfig resource from its proto representation.
func ProtoToDataprocBetaJobPrestoJobLoggingConfig(p *betapb.DataprocBetaJobPrestoJobLoggingConfig) *beta.JobPrestoJobLoggingConfig {
	if p == nil {
		return nil
	}
	obj := &beta.JobPrestoJobLoggingConfig{}
	return obj
}

// ProtoToJobStatus converts a JobStatus resource from its proto representation.
func ProtoToDataprocBetaJobStatus(p *betapb.DataprocBetaJobStatus) *beta.JobStatus {
	if p == nil {
		return nil
	}
	obj := &beta.JobStatus{
		State:          ProtoToDataprocBetaJobStatusStateEnum(p.GetState()),
		Details:        dcl.StringOrNil(p.Details),
		StateStartTime: dcl.StringOrNil(p.GetStateStartTime()),
		Substate:       ProtoToDataprocBetaJobStatusSubstateEnum(p.GetSubstate()),
	}
	return obj
}

// ProtoToJobStatusHistory converts a JobStatusHistory resource from its proto representation.
func ProtoToDataprocBetaJobStatusHistory(p *betapb.DataprocBetaJobStatusHistory) *beta.JobStatusHistory {
	if p == nil {
		return nil
	}
	obj := &beta.JobStatusHistory{
		State:          ProtoToDataprocBetaJobStatusHistoryStateEnum(p.GetState()),
		Details:        dcl.StringOrNil(p.Details),
		StateStartTime: dcl.StringOrNil(p.GetStateStartTime()),
		Substate:       ProtoToDataprocBetaJobStatusHistorySubstateEnum(p.GetSubstate()),
	}
	return obj
}

// ProtoToJobYarnApplications converts a JobYarnApplications resource from its proto representation.
func ProtoToDataprocBetaJobYarnApplications(p *betapb.DataprocBetaJobYarnApplications) *beta.JobYarnApplications {
	if p == nil {
		return nil
	}
	obj := &beta.JobYarnApplications{
		Name:        dcl.StringOrNil(p.Name),
		State:       ProtoToDataprocBetaJobYarnApplicationsStateEnum(p.GetState()),
		Progress:    dcl.Float64OrNil(p.Progress),
		TrackingUrl: dcl.StringOrNil(p.TrackingUrl),
	}
	return obj
}

// ProtoToJobScheduling converts a JobScheduling resource from its proto representation.
func ProtoToDataprocBetaJobScheduling(p *betapb.DataprocBetaJobScheduling) *beta.JobScheduling {
	if p == nil {
		return nil
	}
	obj := &beta.JobScheduling{
		MaxFailuresPerHour: dcl.Int64OrNil(p.MaxFailuresPerHour),
		MaxFailuresTotal:   dcl.Int64OrNil(p.MaxFailuresTotal),
	}
	return obj
}

// ProtoToJob converts a Job resource from its proto representation.
func ProtoToJob(p *betapb.DataprocBetaJob) *beta.Job {
	obj := &beta.Job{
		Reference:               ProtoToDataprocBetaJobReference(p.GetReference()),
		Placement:               ProtoToDataprocBetaJobPlacement(p.GetPlacement()),
		HadoopJob:               ProtoToDataprocBetaJobHadoopJob(p.GetHadoopJob()),
		SparkJob:                ProtoToDataprocBetaJobSparkJob(p.GetSparkJob()),
		PysparkJob:              ProtoToDataprocBetaJobPysparkJob(p.GetPysparkJob()),
		HiveJob:                 ProtoToDataprocBetaJobHiveJob(p.GetHiveJob()),
		PigJob:                  ProtoToDataprocBetaJobPigJob(p.GetPigJob()),
		SparkRJob:               ProtoToDataprocBetaJobSparkRJob(p.GetSparkRJob()),
		SparkSqlJob:             ProtoToDataprocBetaJobSparkSqlJob(p.GetSparkSqlJob()),
		PrestoJob:               ProtoToDataprocBetaJobPrestoJob(p.GetPrestoJob()),
		Status:                  ProtoToDataprocBetaJobStatus(p.GetStatus()),
		SubmittedBy:             dcl.StringOrNil(p.SubmittedBy),
		DriverOutputResourceUri: dcl.StringOrNil(p.DriverOutputResourceUri),
		DriverControlFilesUri:   dcl.StringOrNil(p.DriverControlFilesUri),
		Scheduling:              ProtoToDataprocBetaJobScheduling(p.GetScheduling()),
		Name:                    dcl.StringOrNil(p.Name),
		Done:                    dcl.Bool(p.Done),
		Region:                  dcl.StringOrNil(p.Region),
		Project:                 dcl.StringOrNil(p.Project),
	}
	for _, r := range p.GetStatusHistory() {
		obj.StatusHistory = append(obj.StatusHistory, *ProtoToDataprocBetaJobStatusHistory(r))
	}
	for _, r := range p.GetYarnApplications() {
		obj.YarnApplications = append(obj.YarnApplications, *ProtoToDataprocBetaJobYarnApplications(r))
	}
	return obj
}

// JobStatusStateEnumToProto converts a JobStatusStateEnum enum to its proto representation.
func DataprocBetaJobStatusStateEnumToProto(e *beta.JobStatusStateEnum) betapb.DataprocBetaJobStatusStateEnum {
	if e == nil {
		return betapb.DataprocBetaJobStatusStateEnum(0)
	}
	if v, ok := betapb.DataprocBetaJobStatusStateEnum_value["JobStatusStateEnum"+string(*e)]; ok {
		return betapb.DataprocBetaJobStatusStateEnum(v)
	}
	return betapb.DataprocBetaJobStatusStateEnum(0)
}

// JobStatusSubstateEnumToProto converts a JobStatusSubstateEnum enum to its proto representation.
func DataprocBetaJobStatusSubstateEnumToProto(e *beta.JobStatusSubstateEnum) betapb.DataprocBetaJobStatusSubstateEnum {
	if e == nil {
		return betapb.DataprocBetaJobStatusSubstateEnum(0)
	}
	if v, ok := betapb.DataprocBetaJobStatusSubstateEnum_value["JobStatusSubstateEnum"+string(*e)]; ok {
		return betapb.DataprocBetaJobStatusSubstateEnum(v)
	}
	return betapb.DataprocBetaJobStatusSubstateEnum(0)
}

// JobStatusHistoryStateEnumToProto converts a JobStatusHistoryStateEnum enum to its proto representation.
func DataprocBetaJobStatusHistoryStateEnumToProto(e *beta.JobStatusHistoryStateEnum) betapb.DataprocBetaJobStatusHistoryStateEnum {
	if e == nil {
		return betapb.DataprocBetaJobStatusHistoryStateEnum(0)
	}
	if v, ok := betapb.DataprocBetaJobStatusHistoryStateEnum_value["JobStatusHistoryStateEnum"+string(*e)]; ok {
		return betapb.DataprocBetaJobStatusHistoryStateEnum(v)
	}
	return betapb.DataprocBetaJobStatusHistoryStateEnum(0)
}

// JobStatusHistorySubstateEnumToProto converts a JobStatusHistorySubstateEnum enum to its proto representation.
func DataprocBetaJobStatusHistorySubstateEnumToProto(e *beta.JobStatusHistorySubstateEnum) betapb.DataprocBetaJobStatusHistorySubstateEnum {
	if e == nil {
		return betapb.DataprocBetaJobStatusHistorySubstateEnum(0)
	}
	if v, ok := betapb.DataprocBetaJobStatusHistorySubstateEnum_value["JobStatusHistorySubstateEnum"+string(*e)]; ok {
		return betapb.DataprocBetaJobStatusHistorySubstateEnum(v)
	}
	return betapb.DataprocBetaJobStatusHistorySubstateEnum(0)
}

// JobYarnApplicationsStateEnumToProto converts a JobYarnApplicationsStateEnum enum to its proto representation.
func DataprocBetaJobYarnApplicationsStateEnumToProto(e *beta.JobYarnApplicationsStateEnum) betapb.DataprocBetaJobYarnApplicationsStateEnum {
	if e == nil {
		return betapb.DataprocBetaJobYarnApplicationsStateEnum(0)
	}
	if v, ok := betapb.DataprocBetaJobYarnApplicationsStateEnum_value["JobYarnApplicationsStateEnum"+string(*e)]; ok {
		return betapb.DataprocBetaJobYarnApplicationsStateEnum(v)
	}
	return betapb.DataprocBetaJobYarnApplicationsStateEnum(0)
}

// JobReferenceToProto converts a JobReference resource to its proto representation.
func DataprocBetaJobReferenceToProto(o *beta.JobReference) *betapb.DataprocBetaJobReference {
	if o == nil {
		return nil
	}
	p := &betapb.DataprocBetaJobReference{
		ProjectId: dcl.ValueOrEmptyString(o.ProjectId),
		JobId:     dcl.ValueOrEmptyString(o.JobId),
	}
	return p
}

// JobPlacementToProto converts a JobPlacement resource to its proto representation.
func DataprocBetaJobPlacementToProto(o *beta.JobPlacement) *betapb.DataprocBetaJobPlacement {
	if o == nil {
		return nil
	}
	p := &betapb.DataprocBetaJobPlacement{
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
func DataprocBetaJobHadoopJobToProto(o *beta.JobHadoopJob) *betapb.DataprocBetaJobHadoopJob {
	if o == nil {
		return nil
	}
	p := &betapb.DataprocBetaJobHadoopJob{
		MainJarFileUri: dcl.ValueOrEmptyString(o.MainJarFileUri),
		MainClass:      dcl.ValueOrEmptyString(o.MainClass),
		LoggingConfig:  DataprocBetaJobHadoopJobLoggingConfigToProto(o.LoggingConfig),
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
func DataprocBetaJobHadoopJobLoggingConfigToProto(o *beta.JobHadoopJobLoggingConfig) *betapb.DataprocBetaJobHadoopJobLoggingConfig {
	if o == nil {
		return nil
	}
	p := &betapb.DataprocBetaJobHadoopJobLoggingConfig{}
	p.DriverLogLevels = make(map[string]string)
	for k, r := range o.DriverLogLevels {
		p.DriverLogLevels[k] = r
	}
	return p
}

// JobSparkJobToProto converts a JobSparkJob resource to its proto representation.
func DataprocBetaJobSparkJobToProto(o *beta.JobSparkJob) *betapb.DataprocBetaJobSparkJob {
	if o == nil {
		return nil
	}
	p := &betapb.DataprocBetaJobSparkJob{
		MainJarFileUri: dcl.ValueOrEmptyString(o.MainJarFileUri),
		MainClass:      dcl.ValueOrEmptyString(o.MainClass),
		LoggingConfig:  DataprocBetaJobSparkJobLoggingConfigToProto(o.LoggingConfig),
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
func DataprocBetaJobSparkJobLoggingConfigToProto(o *beta.JobSparkJobLoggingConfig) *betapb.DataprocBetaJobSparkJobLoggingConfig {
	if o == nil {
		return nil
	}
	p := &betapb.DataprocBetaJobSparkJobLoggingConfig{}
	p.DriverLogLevels = make(map[string]string)
	for k, r := range o.DriverLogLevels {
		p.DriverLogLevels[k] = r
	}
	return p
}

// JobPysparkJobToProto converts a JobPysparkJob resource to its proto representation.
func DataprocBetaJobPysparkJobToProto(o *beta.JobPysparkJob) *betapb.DataprocBetaJobPysparkJob {
	if o == nil {
		return nil
	}
	p := &betapb.DataprocBetaJobPysparkJob{
		MainPythonFileUri: dcl.ValueOrEmptyString(o.MainPythonFileUri),
		LoggingConfig:     DataprocBetaJobPysparkJobLoggingConfigToProto(o.LoggingConfig),
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
func DataprocBetaJobPysparkJobLoggingConfigToProto(o *beta.JobPysparkJobLoggingConfig) *betapb.DataprocBetaJobPysparkJobLoggingConfig {
	if o == nil {
		return nil
	}
	p := &betapb.DataprocBetaJobPysparkJobLoggingConfig{}
	p.DriverLogLevels = make(map[string]string)
	for k, r := range o.DriverLogLevels {
		p.DriverLogLevels[k] = r
	}
	return p
}

// JobHiveJobToProto converts a JobHiveJob resource to its proto representation.
func DataprocBetaJobHiveJobToProto(o *beta.JobHiveJob) *betapb.DataprocBetaJobHiveJob {
	if o == nil {
		return nil
	}
	p := &betapb.DataprocBetaJobHiveJob{
		QueryFileUri:      dcl.ValueOrEmptyString(o.QueryFileUri),
		QueryList:         DataprocBetaJobHiveJobQueryListToProto(o.QueryList),
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
func DataprocBetaJobHiveJobQueryListToProto(o *beta.JobHiveJobQueryList) *betapb.DataprocBetaJobHiveJobQueryList {
	if o == nil {
		return nil
	}
	p := &betapb.DataprocBetaJobHiveJobQueryList{}
	for _, r := range o.Queries {
		p.Queries = append(p.Queries, r)
	}
	return p
}

// JobPigJobToProto converts a JobPigJob resource to its proto representation.
func DataprocBetaJobPigJobToProto(o *beta.JobPigJob) *betapb.DataprocBetaJobPigJob {
	if o == nil {
		return nil
	}
	p := &betapb.DataprocBetaJobPigJob{
		QueryFileUri:      dcl.ValueOrEmptyString(o.QueryFileUri),
		QueryList:         DataprocBetaJobPigJobQueryListToProto(o.QueryList),
		ContinueOnFailure: dcl.ValueOrEmptyBool(o.ContinueOnFailure),
		LoggingConfig:     DataprocBetaJobPigJobLoggingConfigToProto(o.LoggingConfig),
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
func DataprocBetaJobPigJobQueryListToProto(o *beta.JobPigJobQueryList) *betapb.DataprocBetaJobPigJobQueryList {
	if o == nil {
		return nil
	}
	p := &betapb.DataprocBetaJobPigJobQueryList{}
	for _, r := range o.Queries {
		p.Queries = append(p.Queries, r)
	}
	return p
}

// JobPigJobLoggingConfigToProto converts a JobPigJobLoggingConfig resource to its proto representation.
func DataprocBetaJobPigJobLoggingConfigToProto(o *beta.JobPigJobLoggingConfig) *betapb.DataprocBetaJobPigJobLoggingConfig {
	if o == nil {
		return nil
	}
	p := &betapb.DataprocBetaJobPigJobLoggingConfig{}
	p.DriverLogLevels = make(map[string]string)
	for k, r := range o.DriverLogLevels {
		p.DriverLogLevels[k] = r
	}
	return p
}

// JobSparkRJobToProto converts a JobSparkRJob resource to its proto representation.
func DataprocBetaJobSparkRJobToProto(o *beta.JobSparkRJob) *betapb.DataprocBetaJobSparkRJob {
	if o == nil {
		return nil
	}
	p := &betapb.DataprocBetaJobSparkRJob{
		MainRFileUri:  dcl.ValueOrEmptyString(o.MainRFileUri),
		LoggingConfig: DataprocBetaJobSparkRJobLoggingConfigToProto(o.LoggingConfig),
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
func DataprocBetaJobSparkRJobLoggingConfigToProto(o *beta.JobSparkRJobLoggingConfig) *betapb.DataprocBetaJobSparkRJobLoggingConfig {
	if o == nil {
		return nil
	}
	p := &betapb.DataprocBetaJobSparkRJobLoggingConfig{}
	p.DriverLogLevels = make(map[string]string)
	for k, r := range o.DriverLogLevels {
		p.DriverLogLevels[k] = r
	}
	return p
}

// JobSparkSqlJobToProto converts a JobSparkSqlJob resource to its proto representation.
func DataprocBetaJobSparkSqlJobToProto(o *beta.JobSparkSqlJob) *betapb.DataprocBetaJobSparkSqlJob {
	if o == nil {
		return nil
	}
	p := &betapb.DataprocBetaJobSparkSqlJob{
		QueryFileUri:  dcl.ValueOrEmptyString(o.QueryFileUri),
		QueryList:     DataprocBetaJobSparkSqlJobQueryListToProto(o.QueryList),
		LoggingConfig: DataprocBetaJobSparkSqlJobLoggingConfigToProto(o.LoggingConfig),
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
func DataprocBetaJobSparkSqlJobQueryListToProto(o *beta.JobSparkSqlJobQueryList) *betapb.DataprocBetaJobSparkSqlJobQueryList {
	if o == nil {
		return nil
	}
	p := &betapb.DataprocBetaJobSparkSqlJobQueryList{}
	for _, r := range o.Queries {
		p.Queries = append(p.Queries, r)
	}
	return p
}

// JobSparkSqlJobLoggingConfigToProto converts a JobSparkSqlJobLoggingConfig resource to its proto representation.
func DataprocBetaJobSparkSqlJobLoggingConfigToProto(o *beta.JobSparkSqlJobLoggingConfig) *betapb.DataprocBetaJobSparkSqlJobLoggingConfig {
	if o == nil {
		return nil
	}
	p := &betapb.DataprocBetaJobSparkSqlJobLoggingConfig{}
	p.DriverLogLevels = make(map[string]string)
	for k, r := range o.DriverLogLevels {
		p.DriverLogLevels[k] = r
	}
	return p
}

// JobPrestoJobToProto converts a JobPrestoJob resource to its proto representation.
func DataprocBetaJobPrestoJobToProto(o *beta.JobPrestoJob) *betapb.DataprocBetaJobPrestoJob {
	if o == nil {
		return nil
	}
	p := &betapb.DataprocBetaJobPrestoJob{
		QueryFileUri:      dcl.ValueOrEmptyString(o.QueryFileUri),
		QueryList:         DataprocBetaJobPrestoJobQueryListToProto(o.QueryList),
		ContinueOnFailure: dcl.ValueOrEmptyBool(o.ContinueOnFailure),
		OutputFormat:      dcl.ValueOrEmptyString(o.OutputFormat),
		LoggingConfig:     DataprocBetaJobPrestoJobLoggingConfigToProto(o.LoggingConfig),
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
func DataprocBetaJobPrestoJobQueryListToProto(o *beta.JobPrestoJobQueryList) *betapb.DataprocBetaJobPrestoJobQueryList {
	if o == nil {
		return nil
	}
	p := &betapb.DataprocBetaJobPrestoJobQueryList{}
	for _, r := range o.Queries {
		p.Queries = append(p.Queries, r)
	}
	return p
}

// JobPrestoJobLoggingConfigToProto converts a JobPrestoJobLoggingConfig resource to its proto representation.
func DataprocBetaJobPrestoJobLoggingConfigToProto(o *beta.JobPrestoJobLoggingConfig) *betapb.DataprocBetaJobPrestoJobLoggingConfig {
	if o == nil {
		return nil
	}
	p := &betapb.DataprocBetaJobPrestoJobLoggingConfig{}
	p.DriverLogLevels = make(map[string]string)
	for k, r := range o.DriverLogLevels {
		p.DriverLogLevels[k] = r
	}
	return p
}

// JobStatusToProto converts a JobStatus resource to its proto representation.
func DataprocBetaJobStatusToProto(o *beta.JobStatus) *betapb.DataprocBetaJobStatus {
	if o == nil {
		return nil
	}
	p := &betapb.DataprocBetaJobStatus{
		State:          DataprocBetaJobStatusStateEnumToProto(o.State),
		Details:        dcl.ValueOrEmptyString(o.Details),
		StateStartTime: dcl.ValueOrEmptyString(o.StateStartTime),
		Substate:       DataprocBetaJobStatusSubstateEnumToProto(o.Substate),
	}
	return p
}

// JobStatusHistoryToProto converts a JobStatusHistory resource to its proto representation.
func DataprocBetaJobStatusHistoryToProto(o *beta.JobStatusHistory) *betapb.DataprocBetaJobStatusHistory {
	if o == nil {
		return nil
	}
	p := &betapb.DataprocBetaJobStatusHistory{
		State:          DataprocBetaJobStatusHistoryStateEnumToProto(o.State),
		Details:        dcl.ValueOrEmptyString(o.Details),
		StateStartTime: dcl.ValueOrEmptyString(o.StateStartTime),
		Substate:       DataprocBetaJobStatusHistorySubstateEnumToProto(o.Substate),
	}
	return p
}

// JobYarnApplicationsToProto converts a JobYarnApplications resource to its proto representation.
func DataprocBetaJobYarnApplicationsToProto(o *beta.JobYarnApplications) *betapb.DataprocBetaJobYarnApplications {
	if o == nil {
		return nil
	}
	p := &betapb.DataprocBetaJobYarnApplications{
		Name:        dcl.ValueOrEmptyString(o.Name),
		State:       DataprocBetaJobYarnApplicationsStateEnumToProto(o.State),
		Progress:    dcl.ValueOrEmptyDouble(o.Progress),
		TrackingUrl: dcl.ValueOrEmptyString(o.TrackingUrl),
	}
	return p
}

// JobSchedulingToProto converts a JobScheduling resource to its proto representation.
func DataprocBetaJobSchedulingToProto(o *beta.JobScheduling) *betapb.DataprocBetaJobScheduling {
	if o == nil {
		return nil
	}
	p := &betapb.DataprocBetaJobScheduling{
		MaxFailuresPerHour: dcl.ValueOrEmptyInt64(o.MaxFailuresPerHour),
		MaxFailuresTotal:   dcl.ValueOrEmptyInt64(o.MaxFailuresTotal),
	}
	return p
}

// JobToProto converts a Job resource to its proto representation.
func JobToProto(resource *beta.Job) *betapb.DataprocBetaJob {
	p := &betapb.DataprocBetaJob{
		Reference:               DataprocBetaJobReferenceToProto(resource.Reference),
		Placement:               DataprocBetaJobPlacementToProto(resource.Placement),
		HadoopJob:               DataprocBetaJobHadoopJobToProto(resource.HadoopJob),
		SparkJob:                DataprocBetaJobSparkJobToProto(resource.SparkJob),
		PysparkJob:              DataprocBetaJobPysparkJobToProto(resource.PysparkJob),
		HiveJob:                 DataprocBetaJobHiveJobToProto(resource.HiveJob),
		PigJob:                  DataprocBetaJobPigJobToProto(resource.PigJob),
		SparkRJob:               DataprocBetaJobSparkRJobToProto(resource.SparkRJob),
		SparkSqlJob:             DataprocBetaJobSparkSqlJobToProto(resource.SparkSqlJob),
		PrestoJob:               DataprocBetaJobPrestoJobToProto(resource.PrestoJob),
		Status:                  DataprocBetaJobStatusToProto(resource.Status),
		SubmittedBy:             dcl.ValueOrEmptyString(resource.SubmittedBy),
		DriverOutputResourceUri: dcl.ValueOrEmptyString(resource.DriverOutputResourceUri),
		DriverControlFilesUri:   dcl.ValueOrEmptyString(resource.DriverControlFilesUri),
		Scheduling:              DataprocBetaJobSchedulingToProto(resource.Scheduling),
		Name:                    dcl.ValueOrEmptyString(resource.Name),
		Done:                    dcl.ValueOrEmptyBool(resource.Done),
		Region:                  dcl.ValueOrEmptyString(resource.Region),
		Project:                 dcl.ValueOrEmptyString(resource.Project),
	}
	for _, r := range resource.StatusHistory {
		p.StatusHistory = append(p.StatusHistory, DataprocBetaJobStatusHistoryToProto(&r))
	}
	for _, r := range resource.YarnApplications {
		p.YarnApplications = append(p.YarnApplications, DataprocBetaJobYarnApplicationsToProto(&r))
	}

	return p
}

// ApplyJob handles the gRPC request by passing it to the underlying Job Apply() method.
func (s *JobServer) applyJob(ctx context.Context, c *beta.Client, request *betapb.ApplyDataprocBetaJobRequest) (*betapb.DataprocBetaJob, error) {
	p := ProtoToJob(request.GetResource())
	res, err := c.ApplyJob(ctx, p)
	if err != nil {
		return nil, err
	}
	r := JobToProto(res)
	return r, nil
}

// ApplyJob handles the gRPC request by passing it to the underlying Job Apply() method.
func (s *JobServer) ApplyDataprocBetaJob(ctx context.Context, request *betapb.ApplyDataprocBetaJobRequest) (*betapb.DataprocBetaJob, error) {
	cl, err := createConfigJob(ctx, request.ServiceAccountFile)
	if err != nil {
		return nil, err
	}
	return s.applyJob(ctx, cl, request)
}

// DeleteJob handles the gRPC request by passing it to the underlying Job Delete() method.
func (s *JobServer) DeleteDataprocBetaJob(ctx context.Context, request *betapb.DeleteDataprocBetaJobRequest) (*emptypb.Empty, error) {

	cl, err := createConfigJob(ctx, request.ServiceAccountFile)
	if err != nil {
		return nil, err
	}
	return &emptypb.Empty{}, cl.DeleteJob(ctx, ProtoToJob(request.GetResource()))

}

// ListDataprocBetaJob handles the gRPC request by passing it to the underlying JobList() method.
func (s *JobServer) ListDataprocBetaJob(ctx context.Context, request *betapb.ListDataprocBetaJobRequest) (*betapb.ListDataprocBetaJobResponse, error) {
	cl, err := createConfigJob(ctx, request.ServiceAccountFile)
	if err != nil {
		return nil, err
	}

	resources, err := cl.ListJob(ctx, request.Project, request.Region)
	if err != nil {
		return nil, err
	}
	var protos []*betapb.DataprocBetaJob
	for _, r := range resources.Items {
		rp := JobToProto(r)
		protos = append(protos, rp)
	}
	return &betapb.ListDataprocBetaJobResponse{Items: protos}, nil
}

func createConfigJob(ctx context.Context, service_account_file string) (*beta.Client, error) {

	conf := dcl.NewConfig(dcl.WithUserAgent("dcl-test"), dcl.WithCredentialsFile(service_account_file))
	return beta.NewClient(conf), nil
}
