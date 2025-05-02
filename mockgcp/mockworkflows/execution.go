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

// +tool:mockgcp-support
// proto.service: google.cloud.workflows.executions.v1.Executions
// proto.message: google.cloud.workflows.executions.v1.Execution

package mockworkflows

import (
	"context"
	// "fmt"
	// "strings"
	"time"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/types/known/durationpb"
	"google.golang.org/protobuf/types/known/timestamppb"

	// "github.com/GoogleCloudPlatform/k8s-config-connector/mockgcp/common/projects"
	pb "cloud.google.com/go/workflows/executions/apiv1/executionspb"
	"github.com/GoogleCloudPlatform/k8s-config-connector/mockgcp/pkg/storage"
)

type WorkflowExecutionsV1 struct {
	*MockService
	pb.UnimplementedExecutionsServer
}

func (s *WorkflowExecutionsV1) CreateExecution(ctx context.Context, req *pb.CreateExecutionRequest) (*pb.Execution, error) {
	fqn := req.GetParent() + "/executions/123456789"
	now := time.Now()
	obj := proto.Clone(req.GetExecution()).(*pb.Execution)
	obj.Name = fqn
	obj.StartTime = timestamppb.New(now)
	// EndTime not to be returned on Create
	obj.State = pb.Execution_ACTIVE
	obj.WorkflowRevisionId = "000001-609"
	// Result not to be returned on Create
	obj.Status = &pb.Execution_Status{}
	if err := s.storage.Create(ctx, fqn, obj); err != nil {
		return nil, err
	}

	return obj, nil
}

func (s *WorkflowExecutionsV1) GetExecution(ctx context.Context, req *pb.GetExecutionRequest) (*pb.Execution, error) {
	name, err := s.parseExecutionName(req.GetName())
	if err != nil {
		return nil, err
	}

	fqn := name.String()

	obj := &pb.Execution{}
	if err := s.storage.Get(ctx, fqn, obj); err != nil {
		if status.Code(err) == codes.NotFound {
			return nil, status.Errorf(codes.NotFound, "Execution %q not found.", fqn)
		}
		return nil, err
	}

	// Pretend execution finished
	now := time.Now()
	obj.Duration = durationpb.New(100 * time.Millisecond)
	obj.EndTime = timestamppb.New(now)
	obj.Result = "\"Hello initial value\""
	obj.State = pb.Execution_SUCCEEDED
	obj.Status = &pb.Execution_Status{
		CurrentSteps: []*pb.Execution_Status_Step{
			&pb.Execution_Status_Step{
				Routine: "main",
				Step: "returnOutput",
			},
		},
	}

	return obj, nil
}

func (s *WorkflowExecutionsV1) ListExecutions(ctx context.Context, req *pb.ListExecutionsRequest) (*pb.ListExecutionsResponse, error) {
	findPrefix := req.GetParent()

	response := &pb.ListExecutionsResponse{}
	findKind := (&pb.Execution{}).ProtoReflect().Descriptor()
	if err := s.storage.List(ctx, findKind, storage.ListOptions{Prefix: findPrefix}, func(obj proto.Message) error {
		execution := obj.(*pb.Execution)
		response.Executions = append(response.Executions, execution)
		return nil
	}); err != nil {
		return nil, err
	}

	return response, nil
}

func (s *WorkflowExecutionsV1) CancelExecution(ctx context.Context, req *pb.CancelExecutionRequest) (*pb.Execution, error) {
	name, err := s.parseExecutionName(req.Name)
	if err != nil {
		return nil, err
	}

	fqn := name.String()
	obj := &pb.Execution{}
	if err := s.storage.Get(ctx, fqn, obj); err != nil {
		return nil, err
	}
	obj.State = pb.Execution_CANCELLED
	if err := s.storage.Update(ctx, fqn, obj); err != nil {
		return nil, err
	}
	return obj, nil
}
