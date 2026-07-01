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
// proto.service: google.cloud.notebooks.v1.NotebookService
// proto.message: google.cloud.notebooks.v1.Execution

package mocknotebooks

import (
	"context"
	"fmt"
	"strings"
	"time"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/types/known/emptypb"
	"google.golang.org/protobuf/types/known/timestamppb"

	pb "cloud.google.com/go/notebooks/apiv1/notebookspb"
	"github.com/GoogleCloudPlatform/k8s-config-connector/mockgcp/common/projects"
	longrunningpb "google.golang.org/genproto/googleapis/longrunning"
)

func (s *NotebookServiceV1) GetExecution(ctx context.Context, req *pb.GetExecutionRequest) (*pb.Execution, error) {
	name, err := s.parseExecutionName(req.GetName())
	if err != nil {
		return nil, err
	}
	fqn := name.String()

	obj := &pb.Execution{}
	if err := s.storage.Get(ctx, fqn, obj); err != nil {
		if status.Code(err) == codes.NotFound {
			return nil, status.Errorf(codes.NotFound, "execution %q not found", fqn)
		}
		return nil, err
	}

	return obj, nil
}

func (s *NotebookServiceV1) CreateExecution(ctx context.Context, req *pb.CreateExecutionRequest) (*longrunningpb.Operation, error) {
	reqName := fmt.Sprintf("%s/executions/%s", req.GetParent(), req.GetExecutionId())
	name, err := s.parseExecutionName(reqName)
	if err != nil {
		return nil, err
	}

	fqn := name.String()
	obj := proto.CloneOf(req.GetExecution())
	obj.Name = fqn
	obj.CreateTime = timestamppb.New(time.Now())
	obj.UpdateTime = timestamppb.New(time.Now())
	obj.State = pb.Execution_SUCCEEDED

	if err := s.storage.Create(ctx, fqn, obj); err != nil {
		return nil, err
	}

	prefix := fmt.Sprintf("projects/%s/locations/%s", name.Project.ID, name.Location)
	metadata := &pb.OperationMetadata{
		ApiVersion:            "v1",
		CreateTime:            timestamppb.New(time.Now()),
		RequestedCancellation: false,
		Target:                name.String(),
		Verb:                  "create",
		Endpoint:              "CreateExecution",
	}
	return s.operations.StartLRO(ctx, prefix, metadata, func() (proto.Message, error) {
		metadata.EndTime = timestamppb.New(time.Now())
		return obj, nil
	})
}

func (s *NotebookServiceV1) DeleteExecution(ctx context.Context, req *pb.DeleteExecutionRequest) (*longrunningpb.Operation, error) {
	name, err := s.parseExecutionName(req.GetName())
	if err != nil {
		return nil, err
	}
	fqn := name.String()

	deletedObj := &pb.Execution{}
	if err := s.storage.Delete(ctx, fqn, deletedObj); err != nil {
		return nil, err
	}

	prefix := fmt.Sprintf("projects/%s/locations/%s", name.Project.ID, name.Location)
	metadata := &pb.OperationMetadata{
		ApiVersion:            "v1",
		CreateTime:            timestamppb.Now(),
		RequestedCancellation: false,
		Target:                name.String(),
		Verb:                  "delete",
		Endpoint:              "DeleteExecution",
	}
	return s.operations.StartLRO(ctx, prefix, metadata, func() (proto.Message, error) {
		metadata.EndTime = timestamppb.New(time.Now())
		return &emptypb.Empty{}, nil
	})
}

type executionName struct {
	Project   *projects.ProjectData
	Location  string
	Execution string
}

func (n *executionName) String() string {
	return "projects/" + n.Project.ID + "/locations/" + n.Location + "/executions/" + n.Execution
}

// parseExecutionName parses a string into an executionName.
// The expected form is projects/<projectID>/locations/<location>/executions/<execution>.
func (s *MockService) parseExecutionName(name string) (*executionName, error) {
	tokens := strings.Split(name, "/")
	if len(tokens) == 6 && tokens[0] == "projects" && tokens[2] == "locations" && tokens[4] == "executions" {
		project, err := s.Projects.GetProjectByID(tokens[1])
		if err != nil {
			return nil, err
		}

		name := &executionName{
			Project:   project,
			Location:  tokens[3],
			Execution: tokens[5],
		}

		return name, nil
	} else {
		return nil, status.Errorf(codes.InvalidArgument, "name %q is not valid", name)
	}
}
