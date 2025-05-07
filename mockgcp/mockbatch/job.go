// Copyright 2024 Google LLC
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
// proto.service: google.cloud.batch.v1.BatchService
// proto.message: google.cloud.batch.v1.Job

package mockbatch

import (
	"context"
	"fmt"
	"regexp"
	"time"

	"cloud.google.com/go/longrunning/autogen/longrunningpb"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/types/known/durationpb"
	"google.golang.org/protobuf/types/known/timestamppb"

	"github.com/GoogleCloudPlatform/k8s-config-connector/mockgcp/common/projects"
	pb "github.com/GoogleCloudPlatform/k8s-config-connector/mockgcp/generated/mockgcp/cloud/batch/v1"
)

func (s *BatchV1) GetJob(ctx context.Context, req *pb.GetJobRequest) (*pb.Job, error) {
	name, err := s.parseJobName(req.Name)
	if err != nil {
		return nil, err
	}

	fqn := name.String()

	obj := &pb.Job{}
	if err := s.storage.Get(ctx, fqn, obj); err != nil {
		return nil, err
	}

	obj.Status.State = pb.JobStatus_QUEUED
	return obj, nil
}

func (s *BatchV1) CreateJob(ctx context.Context, req *pb.CreateJobRequest) (*pb.Job, error) {
	reqName := req.Parent + "/jobs/" + req.JobId
	name, err := s.parseJobName(reqName)
	if err != nil {
		return nil, err
	}

	fqn := name.String()

	obj := proto.Clone(req.Job).(*pb.Job)
	obj.Name = fqn
	obj.Uid = "b9a676df-c595-4c81-9963-f44b8e44e50c"
	obj.CreateTime = timestamppb.Now()
	obj.UpdateTime = timestamppb.Now()
	allocationPolicyLabels := make(map[string]string)
	allocationPolicyLabels["batch-job-id"] = req.JobId
	obj.AllocationPolicy = &pb.AllocationPolicy{
		Labels: allocationPolicyLabels,
		Location: &pb.AllocationPolicy_LocationPolicy{
			AllowedLocations: []string{fmt.Sprintf("regions/%s", name.Location), fmt.Sprintf("zones/%s-a", name.Location), fmt.Sprintf("zones/%s-b", name.Location), fmt.Sprintf("zones/%s-c", name.Location)},
		},
		ServiceAccount: &pb.ServiceAccount{
			Email: "${projectNumber}-compute@developer.gserviceaccount.com",
		},
	}
	if obj.TaskGroups != nil {
		if obj.TaskGroups[0] != nil {
			obj.TaskGroups[0].Name = fmt.Sprintf("%s/taskGroups/group0", fqn)
			obj.TaskGroups[0].Parallelism = 1
			obj.TaskGroups[0].TaskCount = 1
			if obj.TaskGroups[0].TaskSpec != nil && obj.TaskGroups[0].TaskSpec.ComputeResource == nil {
				obj.TaskGroups[0].TaskSpec.ComputeResource = &pb.ComputeResource{
					CpuMilli:  2000,
					MemoryMib: 2000,
				}
			}
		} else {
			obj.TaskGroups[0] = &pb.TaskGroup{
				Name:        fmt.Sprintf("%s/taskGroups/group0", fqn),
				Parallelism: 1,
				TaskCount:   1,
				TaskSpec: &pb.TaskSpec{
					ComputeResource: &pb.ComputeResource{
						CpuMilli:  2000,
						MemoryMib: 2000,
					},
				},
			}
		}
	}
	obj.Status = &pb.JobStatus{
		State:       pb.JobStatus_QUEUED,
		RunDuration: durationpb.New(0 * time.Second),
	}
	if err := s.storage.Create(ctx, fqn, obj); err != nil {
		return nil, err
	}

	// create a task within the job
	taskObj := &pb.Task{}
	taskObjName := fqn + "/taskGroups/group0/tasks/0"
	taskObj.Name = taskObjName
	taskObj.Status = &pb.TaskStatus{}
	if err := s.storage.Create(ctx, taskObjName, taskObj); err != nil {
		return nil, err
	}

	return obj, nil
}

func (s *BatchV1) DeleteJob(ctx context.Context, req *pb.DeleteJobRequest) (*longrunningpb.Operation, error) {
	name, err := s.parseJobName(req.Name)
	if err != nil {
		return nil, err
	}

	fqn := name.String()
	now := time.Now()

	deleted := &pb.Job{}
	if err := s.storage.Delete(ctx, fqn, deleted); err != nil {
		return nil, err
	}

	taskObj := &pb.Task{}
	taskObjName := fqn + "/taskGroups/group0/tasks/0"
	taskObj.Name = taskObjName
	taskObj.Status = &pb.TaskStatus{}
	if err := s.storage.Delete(ctx, taskObjName, taskObj); err != nil {
		return nil, err
	}
	operationMetadata := &pb.OperationMetadata{
		CreateTime:            timestamppb.New(now),
		EndTime:               timestamppb.New(now),
		ApiVersion:            "v1",
		RequestedCancellation: false,
		Verb:                  "delete",
		Target:                fqn,
	}
	return s.operations.DoneLRO(ctx, fqn, operationMetadata, nil)
}

func (s *BatchV1) ListJobs(ctx context.Context, req *pb.ListJobsRequest) (*pb.ListJobsResponse, error) {
	_, err := s.parseJobName(req.Parent + "/jobs/optionalJobId")

	if err != nil {
		return nil, err
	}
	// TODO: Support List

	return &pb.ListJobsResponse{}, nil
}

type jobName struct {
	Project  *projects.ProjectData
	Location string
	Job      string
}

func (n *jobName) String() string {
	return fmt.Sprintf("projects/%s/locations/%s/jobs/%s", n.Project.ID, n.Location, n.Job)
}

// parseJobName parses a string into a jobName.
// The expected form is projects/*/locations/*/jobs/*`.
func (s *MockService) parseJobName(name string) (*jobName, error) {
	// Example: `projects/*/locations/*/jobs/*`
	r := regexp.MustCompile("^projects/(?P<project>[^/]+)/locations/(?P<location>[^/]+)/jobs/(?P<job>[^/]+)$")
	match := r.FindStringSubmatch(name)
	if len(match) == 0 {
		return nil, status.Errorf(codes.InvalidArgument, "name %q is not valid", name)
	}
	m := make(map[string]string)
	for i, n := range r.SubexpNames() {
		if len(n) > 0 {
			m[n] = match[i]
		}
	}

	project, err := s.Projects.GetProjectByID(m["project"])
	if err != nil {
		return nil, err
	}

	jobName := &jobName{
		Project:  project,
		Location: m["location"],
		Job:      m["job"],
	}

	return jobName, nil
}
