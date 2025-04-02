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
// proto.service: google.cloud.dataproc.v1
// proto.message: google.cloud.dataproc.v1.Job

package mockdataproc

import (
	"context"
	"fmt"
	"strings"
	"time"

	"github.com/google/uuid"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/types/known/emptypb"
	"google.golang.org/protobuf/types/known/timestamppb"

	pb "cloud.google.com/go/dataproc/v2/apiv1/dataprocpb"
	"github.com/GoogleCloudPlatform/k8s-config-connector/mockgcp/common/projects"
)

type jobControllerServer struct {
	*MockService
	pb.UnimplementedJobControllerServer
}

func (s *jobControllerServer) GetJob(ctx context.Context, req *pb.GetJobRequest) (*pb.Job, error) {
	name, err := s.parseJobName(req.ProjectId, req.Region, req.JobId)
	if err != nil {
		return nil, err
	}

	fqn := name.String()

	obj := &pb.Job{}
	if err := s.storage.Get(ctx, fqn, obj); err != nil {
		return nil, err
	}

	return obj, nil
}

func (s *jobControllerServer) SubmitJob(ctx context.Context, req *pb.SubmitJobRequest) (*pb.Job, error) {
	if req.RequestId == "" {
		req.RequestId = uuid.New().String()
	}

	name, err := s.buildJobName(req.ProjectId, req.Region, "")
	if err != nil {
		return nil, err
	}

	fqn := name.String()

	now := time.Now()

	obj := proto.Clone(req.GetJob()).(*pb.Job)
	obj.Reference = &pb.JobReference{
		JobId:     name.JobID,
		ProjectId: name.Project.ID,
	}
	obj.Status = &pb.JobStatus{
		State:          pb.JobStatus_PENDING,
		StateStartTime: timestamppb.New(now),
	}

	s.populateDefaultsForJob(obj, name)

	if err := s.storage.Create(ctx, fqn, obj); err != nil {
		return nil, err
	}
	updated, err := mutateObject(ctx, s.storage, fqn, func(obj *pb.Job) error {
		obj.Status.State = pb.JobStatus_SETUP_DONE
		obj.StatusHistory = append(obj.StatusHistory, &pb.JobStatus{
			State:          pb.JobStatus_PENDING,
			StateStartTime: timestamppb.New(now),
		})
		return nil
	})
	if err != nil {
		return nil, err
	}
	updated, err = mutateObject(ctx, s.storage, fqn, func(obj *pb.Job) error {
		obj.Status.State = pb.JobStatus_RUNNING
		obj.StatusHistory = append(obj.StatusHistory, &pb.JobStatus{
			State:          pb.JobStatus_SETUP_DONE,
			StateStartTime: timestamppb.New(now),
		})
		return nil
	})
	if err != nil {
		return nil, err
	}
	updated, err = mutateObject(ctx, s.storage, fqn, func(obj *pb.Job) error {
		obj.Status.State = pb.JobStatus_DONE
		obj.StatusHistory = append(obj.StatusHistory, &pb.JobStatus{
			State:          pb.JobStatus_RUNNING,
			StateStartTime: timestamppb.New(now),
			Details:        "Agent reported job success",
		})
		return nil
	})
	if err != nil {
		return nil, err
	}

	return updated, nil
}

func (s *jobControllerServer) populateDefaultsForJob(obj *pb.Job, name *jobName) {
	if obj.Placement == nil {
		obj.Placement = &pb.JobPlacement{}
	}
	if obj.Placement.ClusterName == "" {
		obj.Placement.ClusterName = "default"
	}

	// Output only fields, set by service
	obj.DriverOutputResourceUri = fmt.Sprintf("gs://dataproc-staging-%s-%d-abcdef/google-cloud-dataproc-metainfo/%s/jobs/%s/driveroutput", name.Region, name.Project.Number, obj.Placement.ClusterName, name.JobID)
	obj.DriverControlFilesUri = fmt.Sprintf("gs://dataproc-staging-%s-%d-abcdef/google-cloud-dataproc-metainfo/%s/jobs/%s/", name.Region, name.Project.Number, obj.Placement.ClusterName, name.JobID)

}

func (s *jobControllerServer) CancelJob(ctx context.Context, req *pb.CancelJobRequest) (*pb.Job, error) {
	name, err := s.buildJobName(req.ProjectId, req.Region, req.JobId)
	if err != nil {
		return nil, err
	}

	fqn := name.String()

	now := time.Now()

	updated, err := mutateObject(ctx, s.storage, fqn, func(obj *pb.Job) error {
		obj.Status.State = pb.JobStatus_CANCELLED
		obj.Status.StateStartTime = timestamppb.New(now)
		return nil
	})
	if err != nil {
		return nil, err
	}

	return updated, nil
}

func (s *jobControllerServer) DeleteJob(ctx context.Context, req *pb.DeleteJobRequest) (*emptypb.Empty, error) {
	name, err := s.buildJobName(req.ProjectId, req.Region, req.JobId)
	if err != nil {
		return nil, err
	}

	fqn := name.String()

	deleted := &pb.Job{}
	if err := s.storage.Delete(ctx, fqn, deleted); err != nil {
		return nil, err
	}

	return &emptypb.Empty{}, nil
}

type jobName struct {
	Project *projects.ProjectData
	Region  string
	JobID   string
}

func (n *jobName) String() string {
	return fmt.Sprintf("projects/%s/regions/%s/jobs/%s", n.Project.ID, n.Region, n.JobID)
}

// parseJobName parses a string into a jobName.
// The expected form is `projects/*/regions/*/jobs/*`.
func (s *MockService) parseJobName(projectID, region, jobID string) (*jobName, error) {
	project, err := s.Projects.GetProjectByID(projectID)
	if err != nil {
		return nil, err
	}

	if region == "" {
		return nil, status.Errorf(codes.InvalidArgument, "region is required")
	}

	name := &jobName{
		Project: project,
		Region:  region,
		JobID:   jobID,
	}

	return name, nil
}

// buildJobName builds a jobName from the components.
func (s *MockService) buildJobName(projectName, region, jobID string) (*jobName, error) {

	project, err := s.Projects.GetProjectByID(projectName)
	if err != nil {
		return nil, err
	}

	if region == "" {
		return nil, status.Errorf(codes.InvalidArgument, "region is required")
	}

	return &jobName{
		Project: project,
		Region:  region,
		JobID:   jobID,
	}, nil
}

// parseJobNameFromHTTPPath parses a string into a jobName based on the HTTP binding pattern.
func (s *MockService) parseJobNameFromHTTPPath(path string) (*jobName, error) {
	tokens := strings.Split(path, "/")
	if len(tokens) == 6 && tokens[0] == "projects" && tokens[2] == "regions" && tokens[4] == "jobs" {
		return s.buildJobName(tokens[1], tokens[3], tokens[5])
	}

	return nil, status.Errorf(codes.InvalidArgument, "invalid http path format %q", path)
}
