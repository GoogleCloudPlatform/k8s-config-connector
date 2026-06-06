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
	"encoding/json"
	"fmt"
	"strings"
	"time"

	pb "github.com/GoogleCloudPlatform/k8s-config-connector/mockgcp/generated/mockgcp/dataflow/v1beta3"
	"google.golang.org/protobuf/types/known/structpb"
	"google.golang.org/protobuf/types/known/timestamppb"
	"k8s.io/klog/v2"
)

type templatesServer struct {
	*MockService
	pb.UnimplementedTemplatesServiceServer
}

func buildStruct(m map[string]any) (*structpb.Struct, error) {
	bytes, err := json.Marshal(m)
	if err != nil {
		return nil, err
	}
	var cleanMap map[string]any
	if err := json.Unmarshal(bytes, &cleanMap); err != nil {
		return nil, err
	}
	return structpb.NewStruct(cleanMap)
}

func (r *templatesServer) CreateJobFromTemplate(ctx context.Context, req *pb.CreateJobFromTemplateRequest) (*pb.Job, error) {
	now := time.Now()
	jobID := now.Format("2006-01-02-15_04_05") + fmt.Sprintf("-%d", now.UnixNano())

	project, err := r.Projects.GetProjectByID(req.GetProjectId())
	if err != nil {
		return nil, err
	}

	jobName := &jobName{
		Project:  project,
		Location: req.GetLocation(),
		Name:     jobID,
	}

	fqn := jobName.String()

	job := &pb.Job{}
	job.Name = req.GetJobName()
	job.CreateTime = timestamppb.New(now)
	job.CurrentStateTime = timestamppb.New(time.Unix(0, 0))
	job.Id = jobID
	job.Location = req.GetLocation()
	job.ProjectId = req.GetProjectId()
	job.StartTime = timestamppb.New(now)

	// Determine job type
	job.Type = pb.JobType_JOB_TYPE_BATCH
	gcsPath := req.GetGcsPath()
	if strings.Contains(gcsPath, "PubSub_to_") || strings.Contains(gcsPath, "streaming") {
		job.Type = pb.JobType_JOB_TYPE_STREAMING
	}

	job.Labels = req.GetEnvironment().GetAdditionalUserLabels()

	env := req.GetEnvironment()
	if env == nil {
		env = &pb.RuntimeEnvironment{}
	}
	job.Environment = &pb.Environment{
		ServiceKmsKeyName:   env.KmsKeyName,
		ServiceAccountEmail: env.ServiceAccountEmail,
		WorkerRegion:        env.WorkerRegion,
		WorkerZone:          env.Zone,
		TempStoragePrefix:   env.TempLocation,
		WorkerPools: []*pb.WorkerPool{
			{
				Kind:            "harness",
				MachineType:     env.MachineType,
				Network:         env.Network,
				Subnetwork:      env.Subnetwork,
				IpConfiguration: env.IpConfiguration,
			},
		},
	}

	pipelineOptions := map[string]any{
		"options": map[string]any{
			"templateLocation":      req.GetGcsPath(),
			"tempLocation":          env.TempLocation,
			"machineType":           env.MachineType,
			"maxNumWorkers":         env.MaxWorkers,
			"network":               env.Network,
			"subnetwork":            env.Subnetwork,
			"serviceAccountEmail":   env.ServiceAccountEmail,
			"zone":                  env.Zone,
			"ipConfiguration":       env.IpConfiguration.String(),
			"experiments":           env.AdditionalExperiments,
			"enableStreamingEngine": env.EnableStreamingEngine,
		},
	}
	if pipelineOptionsVal, err := buildStruct(pipelineOptions); err == nil {
		job.Environment.SdkPipelineOptions = pipelineOptionsVal
	} else {
		klog.Errorf("failed to build pipeline options struct: %v", err)
	}

	if err := r.storage.Create(ctx, fqn, job); err != nil {
		return nil, err
	}

	go func() {
		if err := r.StartClassicJob(fqn, project, req); err != nil {
			klog.Fatalf("error starting job: %v", err)
		}
	}()

	return job, nil
}

func (r *templatesServer) LaunchTemplate(ctx context.Context, req *pb.LaunchTemplateRequest) (*pb.LaunchTemplateResponse, error) {
	now := time.Now()
	jobID := now.Format("2006-01-02-15_04_05") + fmt.Sprintf("-%d", now.UnixNano())

	project, err := r.Projects.GetProjectByID(req.GetProjectId())
	if err != nil {
		return nil, err
	}

	launchParams := req.GetLaunchParameters()
	jobNameVal := launchParams.GetJobName()

	jobNameObj := &jobName{
		Project:  project,
		Location: req.GetLocation(),
		Name:     jobID,
	}

	fqn := jobNameObj.String()

	job := &pb.Job{}
	job.Name = jobNameVal
	job.CreateTime = timestamppb.New(now)
	job.CurrentStateTime = timestamppb.New(time.Unix(0, 0))
	job.Id = jobID
	job.Location = req.GetLocation()
	job.ProjectId = req.GetProjectId()
	job.StartTime = timestamppb.New(now)

	// Determine job type
	job.Type = pb.JobType_JOB_TYPE_BATCH
	gcsPath := req.GetGcsPath()
	if strings.Contains(gcsPath, "PubSub_to_") || strings.Contains(gcsPath, "streaming") {
		job.Type = pb.JobType_JOB_TYPE_STREAMING
	}
	job.Labels = launchParams.GetEnvironment().GetAdditionalUserLabels()

	env := launchParams.GetEnvironment()
	if env == nil {
		env = &pb.RuntimeEnvironment{}
	}
	job.Environment = &pb.Environment{
		ServiceKmsKeyName:   env.KmsKeyName,
		ServiceAccountEmail: env.ServiceAccountEmail,
		WorkerRegion:        env.WorkerRegion,
		WorkerZone:          env.Zone,
		TempStoragePrefix:   env.TempLocation,
		WorkerPools: []*pb.WorkerPool{
			{
				Kind:            "harness",
				MachineType:     env.MachineType,
				Network:         env.Network,
				Subnetwork:      env.Subnetwork,
				IpConfiguration: env.IpConfiguration,
			},
		},
	}

	pipelineOptions := map[string]any{
		"options": map[string]any{
			"templateLocation":      req.GetGcsPath(),
			"tempLocation":          env.TempLocation,
			"machineType":           env.MachineType,
			"maxNumWorkers":         env.MaxWorkers,
			"network":               env.Network,
			"subnetwork":            env.Subnetwork,
			"serviceAccountEmail":   env.ServiceAccountEmail,
			"zone":                  env.Zone,
			"ipConfiguration":       env.IpConfiguration.String(),
			"experiments":           env.AdditionalExperiments,
			"enableStreamingEngine": env.EnableStreamingEngine,
		},
	}
	if pipelineOptionsVal, err := buildStruct(pipelineOptions); err == nil {
		job.Environment.SdkPipelineOptions = pipelineOptionsVal
	} else {
		klog.Errorf("failed to build pipeline options struct: %v", err)
	}

	if launchParams.GetUpdate() {
		existingJob, err := findJobByJobName(ctx, r.storage, project.ID, req.GetLocation(), job.Name)
		if err != nil {
			return nil, err
		}
		if existingJob == nil {
			return nil, fmt.Errorf("existing job not found")
		}
		job.ReplaceJobId = existingJob.GetId()
	}

	if err := r.storage.Create(ctx, fqn, job); err != nil {
		return nil, err
	}

	go func() {
		createReq := &pb.CreateJobFromTemplateRequest{
			ProjectId: req.GetProjectId(),
			Location:  req.GetLocation(),
			JobName:   jobNameVal,
			Template: &pb.CreateJobFromTemplateRequest_GcsPath{
				GcsPath: req.GetGcsPath(),
			},
			Environment: env,
		}
		if err := r.StartClassicJob(fqn, project, createReq); err != nil {
			klog.Fatalf("error starting classic job: %v", err)
		}
	}()

	return &pb.LaunchTemplateResponse{
		Job: job,
	}, nil
}
