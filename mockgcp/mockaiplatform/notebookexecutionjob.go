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
// proto.service: google.cloud.aiplatform.v1.NotebookService
// proto.message: google.cloud.aiplatform.v1.NotebookExecutionJob

package mockaiplatform

import (
	"context"
	"fmt"
	"strings"
	"time"

	pb "cloud.google.com/go/aiplatform/apiv1/aiplatformpb"
	"github.com/GoogleCloudPlatform/k8s-config-connector/mockgcp/common/projects"
	"github.com/GoogleCloudPlatform/k8s-config-connector/mockgcp/pkg/storage"
	longrunningpb "google.golang.org/genproto/googleapis/longrunning"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/types/known/emptypb"
	"google.golang.org/protobuf/types/known/timestamppb"
)

type notebookServiceV1 struct {
	*MockService
	pb.UnimplementedNotebookServiceServer
}

type NotebookExecutionJobName struct {
	Project                *projects.ProjectData
	Location               string
	NotebookExecutionJobID string
}

func (n *NotebookExecutionJobName) String() string {
	return fmt.Sprintf("projects/%d/locations/%s/notebookExecutionJobs/%s", n.Project.Number, n.Location, n.NotebookExecutionJobID)
}

// parseNotebookExecutionJobName parses a string into a NotebookExecutionJobName.
// The expected form of input string is projects/<projectID|Number>/locations/<location>/notebookExecutionJobs/<notebookExecutionJobID>
func (s *MockService) parseNotebookExecutionJobName(name string) (*NotebookExecutionJobName, error) {
	tokens := strings.Split(name, "/")

	if len(tokens) == 6 && tokens[0] == "projects" && tokens[2] == "locations" && tokens[4] == "notebookExecutionJobs" {
		projectName, err := projects.ParseProjectName(tokens[0] + "/" + tokens[1])
		if err != nil {
			return nil, err
		}
		project, err := s.Projects.GetProject(projectName)
		if err != nil {
			return nil, err
		}

		return &NotebookExecutionJobName{
			Project:                project,
			Location:               tokens[3],
			NotebookExecutionJobID: tokens[5],
		}, nil
	}
	return nil, status.Errorf(codes.InvalidArgument, "name %q is not valid", name)
}

type NotebookExecutionJobParent struct {
	Project  *projects.ProjectData
	Location string
}

func (s *MockService) parseNotebookExecutionJobParent(parent string) (*NotebookExecutionJobParent, error) {
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

		return &NotebookExecutionJobParent{
			Project:  project,
			Location: tokens[3],
		}, nil
	}
	return nil, status.Errorf(codes.InvalidArgument, "parent %q is not valid", parent)
}

func (s *MockService) populateNotebookExecutionJobDefaults(obj *pb.NotebookExecutionJob, name *NotebookExecutionJobName) {
	if obj.JobState == pb.JobState_JOB_STATE_UNSPECIFIED {
		obj.JobState = pb.JobState_JOB_STATE_SUCCEEDED
	}
	if obj.GetServiceAccount() == "" && obj.GetExecutionUser() == "" {
		obj.ExecutionIdentity = &pb.NotebookExecutionJob_ServiceAccount{
			ServiceAccount: fmt.Sprintf("%d-compute@developer.gserviceaccount.com", name.Project.Number),
		}
	}
	// directNotebookSource.content is write-only and omitted on output by real GCP
	if obj.GetDirectNotebookSource() != nil {
		obj.NotebookSource = &pb.NotebookExecutionJob_DirectNotebookSource_{
			DirectNotebookSource: &pb.NotebookExecutionJob_DirectNotebookSource{},
		}
	}
}

func (s *notebookServiceV1) GetNotebookExecutionJob(ctx context.Context, req *pb.GetNotebookExecutionJobRequest) (*pb.NotebookExecutionJob, error) {
	name, err := s.parseNotebookExecutionJobName(req.GetName())
	if err != nil {
		return nil, err
	}

	fqn := name.String()

	obj := &pb.NotebookExecutionJob{}
	if err := s.storage.Get(ctx, fqn, obj); err != nil {
		if status.Code(err) == codes.NotFound {
			return nil, status.Errorf(codes.NotFound, "NotebookExecutionJob %s is not found.", req.GetName())
		}
		return nil, err
	}

	s.populateNotebookExecutionJobDefaults(obj, name)

	return obj, nil
}

func (s *notebookServiceV1) CreateNotebookExecutionJob(ctx context.Context, req *pb.CreateNotebookExecutionJobRequest) (*longrunningpb.Operation, error) {
	reqName := fmt.Sprintf("%s/notebookExecutionJobs/%s", req.GetParent(), req.GetNotebookExecutionJobId())
	name, err := s.parseNotebookExecutionJobName(reqName)
	if err != nil {
		return nil, err
	}

	fqn := name.String()

	existing := &pb.NotebookExecutionJob{}
	if err := s.storage.Get(ctx, fqn, existing); err == nil {
		return nil, status.Errorf(codes.AlreadyExists, "NotebookExecutionJob %s already exists", fqn)
	}

	now := time.Now()

	obj := proto.Clone(req.GetNotebookExecutionJob()).(*pb.NotebookExecutionJob)
	obj.Name = fqn
	obj.CreateTime = timestamppb.New(now)
	obj.UpdateTime = timestamppb.New(now)
	s.populateNotebookExecutionJobDefaults(obj, name)

	if err := s.storage.Create(ctx, fqn, obj); err != nil {
		return nil, err
	}

	op := &pb.CreateNotebookExecutionJobOperationMetadata{}
	op.GenericMetadata = &pb.GenericOperationMetadata{
		CreateTime: timestamppb.New(now),
		UpdateTime: timestamppb.New(now),
	}

	opPrefix := fqn
	return s.operations.StartLRO(ctx, opPrefix, op, func() (proto.Message, error) {
		op.ProgressMessage = "NotebookExecutionJob is initializing, waiting for execution to start."
		return &pb.NotebookExecutionJob{
			Name: fqn,
		}, nil
	})
}

func (s *notebookServiceV1) DeleteNotebookExecutionJob(ctx context.Context, req *pb.DeleteNotebookExecutionJobRequest) (*longrunningpb.Operation, error) {
	name, err := s.parseNotebookExecutionJobName(req.GetName())
	if err != nil {
		return nil, err
	}

	fqn := name.String()
	now := time.Now()

	deleted := &pb.NotebookExecutionJob{}
	if err := s.storage.Delete(ctx, fqn, deleted); err != nil {
		if status.Code(err) == codes.NotFound {
			return nil, status.Errorf(codes.NotFound, "NotebookExecutionJob %s is not found.", req.GetName())
		}
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

func (s *notebookServiceV1) ListNotebookExecutionJobs(ctx context.Context, req *pb.ListNotebookExecutionJobsRequest) (*pb.ListNotebookExecutionJobsResponse, error) {
	parent, err := s.parseNotebookExecutionJobParent(req.GetParent())
	if err != nil {
		return nil, err
	}

	prefix := fmt.Sprintf("projects/%d/locations/%s/notebookExecutionJobs/", parent.Project.Number, parent.Location)
	response := &pb.ListNotebookExecutionJobsResponse{}
	findKind := (&pb.NotebookExecutionJob{}).ProtoReflect().Descriptor()
	if err := s.storage.List(ctx, findKind, storage.ListOptions{Prefix: prefix}, func(obj proto.Message) error {
		item := obj.(*pb.NotebookExecutionJob)
		name, err := s.parseNotebookExecutionJobName(item.GetName())
		if err == nil {
			s.populateNotebookExecutionJobDefaults(item, name)
		}
		response.NotebookExecutionJobs = append(response.NotebookExecutionJobs, item)
		return nil
	}); err != nil {
		return nil, err
	}

	return response, nil
}
