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

// +tool:mockgcp-support
// proto.service: google.cloud.aiplatform.v1.JobService
// proto.message: google.cloud.aiplatform.v1.ModelDeploymentMonitoringJob

package mockaiplatform

import (
	"context"
	"fmt"
	"strings"
	"time"

	pb "cloud.google.com/go/aiplatform/apiv1/aiplatformpb"
	"github.com/GoogleCloudPlatform/k8s-config-connector/mockgcp/common/fields"
	"github.com/GoogleCloudPlatform/k8s-config-connector/mockgcp/common/projects"
	"github.com/GoogleCloudPlatform/k8s-config-connector/mockgcp/pkg/storage"
	longrunning "google.golang.org/genproto/googleapis/longrunning"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/types/known/emptypb"
	"google.golang.org/protobuf/types/known/timestamppb"
)

func (s *jobService) GetModelDeploymentMonitoringJob(ctx context.Context, req *pb.GetModelDeploymentMonitoringJobRequest) (*pb.ModelDeploymentMonitoringJob, error) {
	name, err := s.parseModelDeploymentMonitoringJobName(req.GetName())
	if err != nil {
		return nil, err
	}

	fqn := name.StorageKey()

	obj := &pb.ModelDeploymentMonitoringJob{}
	if err := s.storage.Get(ctx, fqn, obj); err != nil {
		return nil, err
	}

	result := proto.Clone(obj).(*pb.ModelDeploymentMonitoringJob)
	result.Name = name.StringWithProjectNumber()
	populateModelDeploymentMonitoringJobDefaults(result)
	return result, nil
}

func (s *jobService) CreateModelDeploymentMonitoringJob(ctx context.Context, req *pb.CreateModelDeploymentMonitoringJobRequest) (*pb.ModelDeploymentMonitoringJob, error) {
	parent, err := s.parseModelDeploymentMonitoringJobParent(req.GetParent())
	if err != nil {
		return nil, err
	}

	id := fmt.Sprintf("%d", time.Now().UnixNano())
	name := &ModelDeploymentMonitoringJobName{
		Project:                      parent.Project,
		Location:                     parent.Location,
		ModelDeploymentMonitoringJob: id,
	}

	fqn := name.StorageKey()
	now := time.Now()

	obj := proto.Clone(req.GetModelDeploymentMonitoringJob()).(*pb.ModelDeploymentMonitoringJob)
	obj.Name = name.StringWithProjectNumber()
	obj.State = pb.JobState_JOB_STATE_RUNNING
	obj.ScheduleState = pb.ModelDeploymentMonitoringJob_OFFLINE
	obj.CreateTime = timestamppb.New(now)
	obj.UpdateTime = timestamppb.New(now)
	obj.NextScheduleTime = timestamppb.New(now.Add(time.Hour))

	populateModelDeploymentMonitoringJobDefaults(obj)

	if err := s.storage.Create(ctx, fqn, obj); err != nil {
		return nil, err
	}

	result := proto.Clone(obj).(*pb.ModelDeploymentMonitoringJob)
	return result, nil
}

func (s *jobService) UpdateModelDeploymentMonitoringJob(ctx context.Context, req *pb.UpdateModelDeploymentMonitoringJobRequest) (*longrunning.Operation, error) {
	reqObj := req.GetModelDeploymentMonitoringJob()
	if reqObj == nil {
		return nil, status.Errorf(codes.InvalidArgument, "model_deployment_monitoring_job is required")
	}

	name, err := s.parseModelDeploymentMonitoringJobName(reqObj.GetName())
	if err != nil {
		return nil, err
	}

	fqn := name.StorageKey()
	now := time.Now()

	obj := &pb.ModelDeploymentMonitoringJob{}
	if err := s.storage.Get(ctx, fqn, obj); err != nil {
		return nil, err
	}

	paths := req.GetUpdateMask().GetPaths()
	if len(paths) == 0 {
		return nil, status.Errorf(codes.InvalidArgument, "update_mask must be provided")
	}

	if err := fields.UpdateByFieldMask(obj, reqObj, paths); err != nil {
		return nil, status.Errorf(codes.Internal, "failed to update ModelDeploymentMonitoringJob fields: %v", err)
	}

	obj.UpdateTime = timestamppb.New(now)
	populateModelDeploymentMonitoringJobDefaults(obj)

	if err := s.storage.Update(ctx, fqn, obj); err != nil {
		return nil, err
	}

	op := &pb.UpdateModelDeploymentMonitoringJobOperationMetadata{}
	op.GenericMetadata = &pb.GenericOperationMetadata{
		CreateTime: timestamppb.New(now),
		UpdateTime: timestamppb.New(now),
	}
	opPrefix := fmt.Sprintf("projects/%d/locations/%s", name.Project.Number, name.Location)
	return s.operations.DoneLRO(ctx, opPrefix, op, obj)
}

func (s *jobService) DeleteModelDeploymentMonitoringJob(ctx context.Context, req *pb.DeleteModelDeploymentMonitoringJobRequest) (*longrunning.Operation, error) {
	name, err := s.parseModelDeploymentMonitoringJobName(req.GetName())
	if err != nil {
		return nil, err
	}

	fqn := name.StorageKey()
	now := time.Now()

	deleted := &pb.ModelDeploymentMonitoringJob{}
	if err := s.storage.Delete(ctx, fqn, deleted); err != nil {
		return nil, err
	}

	op := &pb.DeleteOperationMetadata{}
	op.GenericMetadata = &pb.GenericOperationMetadata{
		CreateTime: timestamppb.New(now),
		UpdateTime: timestamppb.New(now),
	}
	opPrefix := fmt.Sprintf("projects/%d/locations/%s", name.Project.Number, name.Location)
	return s.operations.DoneLRO(ctx, opPrefix, op, &emptypb.Empty{})
}

func (s *jobService) PauseModelDeploymentMonitoringJob(ctx context.Context, req *pb.PauseModelDeploymentMonitoringJobRequest) (*emptypb.Empty, error) {
	name, err := s.parseModelDeploymentMonitoringJobName(req.GetName())
	if err != nil {
		return nil, err
	}

	fqn := name.StorageKey()
	obj := &pb.ModelDeploymentMonitoringJob{}
	if err := s.storage.Get(ctx, fqn, obj); err != nil {
		return nil, err
	}

	obj.State = pb.JobState_JOB_STATE_PAUSED
	obj.UpdateTime = timestamppb.New(time.Now())

	if err := s.storage.Update(ctx, fqn, obj); err != nil {
		return nil, err
	}

	return &emptypb.Empty{}, nil
}

func (s *jobService) ResumeModelDeploymentMonitoringJob(ctx context.Context, req *pb.ResumeModelDeploymentMonitoringJobRequest) (*emptypb.Empty, error) {
	name, err := s.parseModelDeploymentMonitoringJobName(req.GetName())
	if err != nil {
		return nil, err
	}

	fqn := name.StorageKey()
	obj := &pb.ModelDeploymentMonitoringJob{}
	if err := s.storage.Get(ctx, fqn, obj); err != nil {
		return nil, err
	}

	obj.State = pb.JobState_JOB_STATE_RUNNING
	obj.UpdateTime = timestamppb.New(time.Now())

	if err := s.storage.Update(ctx, fqn, obj); err != nil {
		return nil, err
	}

	return &emptypb.Empty{}, nil
}

func (s *jobService) ListModelDeploymentMonitoringJobs(ctx context.Context, req *pb.ListModelDeploymentMonitoringJobsRequest) (*pb.ListModelDeploymentMonitoringJobsResponse, error) {
	parent, err := s.parseModelDeploymentMonitoringJobParent(req.GetParent())
	if err != nil {
		return nil, err
	}

	findPrefix := fmt.Sprintf("projects/%d/locations/%s/modelDeploymentMonitoringJobs/", parent.Project.Number, parent.Location)
	findPrefixID := fmt.Sprintf("projects/%s/locations/%s/modelDeploymentMonitoringJobs/", parent.Project.ID, parent.Location)

	var jobs []*pb.ModelDeploymentMonitoringJob
	kind := (&pb.ModelDeploymentMonitoringJob{}).ProtoReflect().Descriptor()
	if err := s.storage.List(ctx, kind, storage.ListOptions{}, func(obj proto.Message) error {
		job := obj.(*pb.ModelDeploymentMonitoringJob)
		if strings.HasPrefix(job.GetName(), findPrefix) || strings.HasPrefix(job.GetName(), findPrefixID) {
			result := proto.Clone(job).(*pb.ModelDeploymentMonitoringJob)
			populateModelDeploymentMonitoringJobDefaults(result)
			jobs = append(jobs, result)
		}
		return nil
	}); err != nil {
		return nil, err
	}

	return &pb.ListModelDeploymentMonitoringJobsResponse{
		ModelDeploymentMonitoringJobs: jobs,
	}, nil
}

func populateModelDeploymentMonitoringJobDefaults(job *pb.ModelDeploymentMonitoringJob) {
	// Add BigqueryTables if Endpoint is set and BigqueryTables is empty
	if len(job.BigqueryTables) == 0 && job.Endpoint != "" {
		endpointParts := strings.Split(job.Endpoint, "/")
		endpointID := endpointParts[len(endpointParts)-1]
		projectID := ""
		if len(endpointParts) >= 2 && endpointParts[0] == "projects" {
			projectID = endpointParts[1]
		}
		if projectID != "" && endpointID != "" {
			job.BigqueryTables = []*pb.ModelDeploymentMonitoringBigQueryTable{
				{
					LogSource:                           pb.ModelDeploymentMonitoringBigQueryTable_SERVING,
					LogType:                             pb.ModelDeploymentMonitoringBigQueryTable_PREDICT,
					BigqueryTablePath:                   fmt.Sprintf("bq://%s.model_deployment_monitoring_%s.serving_predict", projectID, endpointID),
					RequestResponseLoggingSchemaVersion: "v1",
				},
			}
		}
	}
}

type ModelDeploymentMonitoringJobName struct {
	Project                      *projects.ProjectData
	Location                     string
	ModelDeploymentMonitoringJob string
}

func (n *ModelDeploymentMonitoringJobName) StringWithProjectNumber() string {
	return fmt.Sprintf("projects/%d/locations/%s/modelDeploymentMonitoringJobs/%s", n.Project.Number, n.Location, n.ModelDeploymentMonitoringJob)
}

func (n *ModelDeploymentMonitoringJobName) StorageKey() string {
	return fmt.Sprintf("projects/%d/locations/%s/modelDeploymentMonitoringJobs/%s", n.Project.Number, n.Location, n.ModelDeploymentMonitoringJob)
}

type ModelDeploymentMonitoringJobParent struct {
	Project  *projects.ProjectData
	Location string
}

func (s *MockService) parseModelDeploymentMonitoringJobParent(parent string) (*ModelDeploymentMonitoringJobParent, error) {
	tokens := strings.Split(parent, "/")

	if len(tokens) == 4 && tokens[0] == "projects" && tokens[2] == "locations" {
		projectName, err := projects.ParseProjectName(tokens[0] + "/" + tokens[1])
		if err != nil {
			return nil, err
		}
		project, err := s.Projects.GetProject(projectName)
		if err != nil {
			return nil, err
		}

		return &ModelDeploymentMonitoringJobParent{
			Project:  project,
			Location: tokens[3],
		}, nil
	}
	return nil, status.Errorf(codes.InvalidArgument, "parent %q is not valid", parent)
}

// parseModelDeploymentMonitoringJobName parses a string into a ModelDeploymentMonitoringJobName.
// The expected form is projects/<projectID_or_number>/locations/<location>/modelDeploymentMonitoringJobs/<id>
func (s *MockService) parseModelDeploymentMonitoringJobName(name string) (*ModelDeploymentMonitoringJobName, error) {
	tokens := strings.Split(name, "/")

	if len(tokens) == 6 && tokens[0] == "projects" && tokens[2] == "locations" && tokens[4] == "modelDeploymentMonitoringJobs" {
		projectName, err := projects.ParseProjectName(tokens[0] + "/" + tokens[1])
		if err != nil {
			return nil, err
		}
		project, err := s.Projects.GetProject(projectName)
		if err != nil {
			return nil, err
		}

		jobID := tokens[5]
		if !isNumeric(jobID) {
			return nil, status.Errorf(codes.InvalidArgument, "invalid model deployment monitoring job id: %s", jobID)
		}

		return &ModelDeploymentMonitoringJobName{
			Project:                      project,
			Location:                     tokens[3],
			ModelDeploymentMonitoringJob: jobID,
		}, nil
	}
	return nil, status.Errorf(codes.InvalidArgument, "name %q is not valid", name)
}
