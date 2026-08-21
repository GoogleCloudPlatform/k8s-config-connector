// Copyright 2026 Google LLC
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

package mockdataflow

import (
	"context"
	"fmt"
	"strings"
	"time"

	"github.com/GoogleCloudPlatform/k8s-config-connector/mockgcp/common/projects"
	pb "github.com/GoogleCloudPlatform/k8s-config-connector/mockgcp/generated/mockgcp/dataflow/v1beta3"
	"github.com/GoogleCloudPlatform/k8s-config-connector/mockgcp/pkg/storage"
	"google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/types/known/structpb"
	"google.golang.org/protobuf/types/known/timestamppb"
)

type templatesServer struct {
	*MockService
	pb.UnimplementedTemplatesServiceServer
}

func isStreamingJob(jobName string, gcsPath string) bool {
	gcsPath = strings.ToLower(gcsPath)
	if strings.Contains(strings.ToLower(jobName), "streaming") {
		return true
	}
	if strings.Contains(gcsPath, "pubsub") || strings.Contains(gcsPath, "streaming") {
		return true
	}
	return false
}

func (r *templatesServer) stopExistingJob(ctx context.Context, projectID string, location string, targetJobName string) {
	prefix := fmt.Sprintf("projects/%s/locations/%s/jobs/", projectID, location)
	var activeJobs []*pb.Job
	var activeFqns []string

	findKind := (&pb.Job{}).ProtoReflect().Descriptor()
	_ = r.storage.List(ctx, findKind, storage.ListOptions{
		Prefix: prefix,
	}, func(obj proto.Message) error {
		job := obj.(*pb.Job)
		if job.Name == targetJobName && isJobActive(job.CurrentState) {
			activeJobs = append(activeJobs, job)

			jobNameObj := &jobName{
				Project:  &projects.ProjectData{ID: projectID},
				Location: location,
				Name:     job.Id,
			}
			activeFqns = append(activeFqns, jobNameObj.String())
		}
		return nil
	})

	for i, job := range activeJobs {
		now := time.Now()
		job.CurrentState = pb.JobState_JOB_STATE_CANCELLING
		job.CurrentStateTime = timestamppb.New(now)
		job.RequestedState = pb.JobState_JOB_STATE_CANCELLED
		fqn := activeFqns[i]
		_ = r.storage.Update(ctx, fqn, job)

		go func(fqn string) {
			time.Sleep(10 * time.Second)
			_ = r.StopJob(fqn)
		}(fqn)
	}
}

func (r *templatesServer) CreateJobFromTemplate(ctx context.Context, req *pb.CreateJobFromTemplateRequest) (*pb.Job, error) {
	now := time.Now()
	jobID := now.Format("2006-01-02-15_04_05") + fmt.Sprintf("-%d", now.UnixNano())

	project, err := r.Projects.GetProjectByID(req.GetProjectId())
	if err != nil {
		return nil, err
	}

	location := req.GetLocation()
	if location == "" {
		location = "us-central1"
	}

	r.stopExistingJob(ctx, project.ID, location, req.GetJobName())

	jobName := &jobName{
		Project:  project,
		Location: location,
		Name:     jobID,
	}

	fqn := jobName.String()

	job := &pb.Job{}
	job.Name = req.GetJobName()
	job.CreateTime = timestamppb.New(now)
	job.CurrentStateTime = timestamppb.New(now)
	job.Id = jobID
	job.Location = location
	job.ProjectId = project.ID
	job.StartTime = timestamppb.New(now)
	job.CurrentState = pb.JobState_JOB_STATE_RUNNING
	job.Type = pb.JobType_JOB_TYPE_BATCH

	isStreaming := isStreamingJob(req.GetJobName(), req.GetGcsPath())
	if isStreaming {
		job.Type = pb.JobType_JOB_TYPE_STREAMING
	}

	labels := make(map[string]string)
	for k, v := range req.GetEnvironment().GetAdditionalUserLabels() {
		labels[k] = v
	}
	job.Labels = labels

	job.Environment = &pb.Environment{
		Dataset: "bigquery.googleapis.com/cloud_dataflow",
	}

	optionsMap := map[string]interface{}{
		"templateLocation":      req.GetGcsPath(),
		"tempLocation":          req.GetEnvironment().GetTempLocation(),
		"machineType":           req.GetEnvironment().GetMachineType(),
		"network":               req.GetEnvironment().GetNetwork(),
		"serviceAccountEmail":   req.GetEnvironment().GetServiceAccountEmail(),
		"subnetwork":            req.GetEnvironment().GetSubnetwork(),
		"enableStreamingEngine": isStreaming,
	}

	if len(req.GetEnvironment().GetAdditionalExperiments()) > 0 {
		var exps []interface{}
		for _, e := range req.GetEnvironment().GetAdditionalExperiments() {
			exps = append(exps, e)
		}
		optionsMap["experiments"] = exps
	}

	sdkPipelineOptions := map[string]interface{}{
		"options": optionsMap,
	}

	sdkPipelineOptionsVal, err := structpb.NewStruct(sdkPipelineOptions)
	if err == nil {
		job.Environment.SdkPipelineOptions = sdkPipelineOptionsVal
	}

	if err := r.storage.Create(ctx, fqn, job); err != nil {
		return nil, err
	}

	return job, nil
}

func (r *templatesServer) LaunchTemplate(ctx context.Context, req *pb.LaunchTemplateRequest) (*pb.LaunchTemplateResponse, error) {
	now := time.Now()
	jobID := now.Format("2006-01-02-15_04_05") + fmt.Sprintf("-%d", now.UnixNano())

	project, err := r.Projects.GetProjectByID(req.GetProjectId())
	if err != nil {
		return nil, err
	}

	location := req.GetLocation()
	if location == "" {
		location = "us-central1"
	}

	r.stopExistingJob(ctx, project.ID, location, req.GetLaunchParameters().GetJobName())

	jobName := &jobName{
		Project:  project,
		Location: location,
		Name:     jobID,
	}

	fqn := jobName.String()

	job := &pb.Job{}
	job.Name = req.GetLaunchParameters().GetJobName()
	job.CreateTime = timestamppb.New(now)
	job.CurrentStateTime = timestamppb.New(now)
	job.Id = jobID
	job.Location = location
	job.ProjectId = project.ID
	job.StartTime = timestamppb.New(now)
	job.CurrentState = pb.JobState_JOB_STATE_RUNNING
	job.Type = pb.JobType_JOB_TYPE_BATCH

	isStreaming := isStreamingJob(req.GetLaunchParameters().GetJobName(), req.GetGcsPath())
	if isStreaming {
		job.Type = pb.JobType_JOB_TYPE_STREAMING
	}

	labels := make(map[string]string)
	for k, v := range req.GetLaunchParameters().GetEnvironment().GetAdditionalUserLabels() {
		labels[k] = v
	}
	job.Labels = labels

	job.Environment = &pb.Environment{
		Dataset: "bigquery.googleapis.com/cloud_dataflow",
	}

	optionsMap := map[string]interface{}{
		"templateLocation":      req.GetGcsPath(),
		"tempLocation":          req.GetLaunchParameters().GetEnvironment().GetTempLocation(),
		"machineType":           req.GetLaunchParameters().GetEnvironment().GetMachineType(),
		"network":               req.GetLaunchParameters().GetEnvironment().GetNetwork(),
		"serviceAccountEmail":   req.GetLaunchParameters().GetEnvironment().GetServiceAccountEmail(),
		"subnetwork":            req.GetLaunchParameters().GetEnvironment().GetSubnetwork(),
		"enableStreamingEngine": isStreaming,
	}

	if len(req.GetLaunchParameters().GetEnvironment().GetAdditionalExperiments()) > 0 {
		var exps []interface{}
		for _, e := range req.GetLaunchParameters().GetEnvironment().GetAdditionalExperiments() {
			exps = append(exps, e)
		}
		optionsMap["experiments"] = exps
	}

	sdkPipelineOptions := map[string]interface{}{
		"options": optionsMap,
	}

	sdkPipelineOptionsVal, err := structpb.NewStruct(sdkPipelineOptions)
	if err == nil {
		job.Environment.SdkPipelineOptions = sdkPipelineOptionsVal
	}

	if err := r.storage.Create(ctx, fqn, job); err != nil {
		return nil, err
	}

	return &pb.LaunchTemplateResponse{
		Job: job,
	}, nil
}
