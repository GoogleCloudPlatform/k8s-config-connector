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
// proto.service: google.cloud.workflows.v1.Workflows
// proto.message: google.cloud.workflows.v1.Workflow

package mockworkflows

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

	"github.com/GoogleCloudPlatform/k8s-config-connector/mockgcp/common/projects"
	pb "github.com/GoogleCloudPlatform/k8s-config-connector/mockgcp/generated/mockgcp/cloud/workflows/v1"
	longrunningpb "google.golang.org/genproto/googleapis/longrunning"
)

func (s *WorkflowsV1) GetWorkflow(ctx context.Context, req *pb.GetWorkflowRequest) (*pb.Workflow, error) {
	name, err := s.parseWorkflowName(req.Name)
	if err != nil {
		return nil, err
	}

	fqn := name.String()

	obj := &pb.Workflow{}
	if err := s.storage.Get(ctx, fqn, obj); err != nil {
		if status.Code(err) == codes.NotFound {
			return nil, status.Errorf(codes.NotFound, "Resource '%v' was not found", name)
		}
		return nil, err
	}
	return obj, nil
}

func (s *WorkflowsV1) CreateWorkflow(ctx context.Context, req *pb.CreateWorkflowRequest) (*longrunningpb.Operation, error) {
	reqName := fmt.Sprintf("%s/workflows/%s", req.GetParent(), req.GetWorkflowId())
	name, err := s.parseWorkflowName(reqName)
	if err != nil {
		return nil, err
	}
	fqn := name.String()

	now := time.Now()
	obj := proto.Clone(req.GetWorkflow()).(*pb.Workflow)
	obj.Name = fqn
	obj.CreateTime = timestamppb.New(now)
	obj.UpdateTime = timestamppb.New(now)
	obj.RevisionCreateTime = timestamppb.New(now)
	obj.ServiceAccount = fmt.Sprintf("projects/%s/serviceAccounts/%d-compute@developer.gserviceaccount.com", name.Project.ID, name.Project.Number)
	obj.RevisionId = "000001-a4d" // TODO: increment
	obj.State = pb.Workflow_ACTIVE
	s.populateDefaultsForWorkflow(obj)

	if err := s.storage.Create(ctx, fqn, obj); err != nil {
		return nil, err
	}

	lroPrefix := fmt.Sprintf("projects/%s/locations/%s", name.Project.ID, name.Location)
	// // Returns with no createTime
	// lroRet := proto.Clone(obj).(*pb.Workflow)
	// lroRet.CreateTime = nil
	// lroRet.UpdateTime = nil
	// lroRet.RevisionCreateTime = nil

	lroMetadata := &pb.OperationMetadata{
		ApiVersion: "v1",
		CreateTime: timestamppb.Now(),
		Target:     fqn,
		Verb:       "create",
	}
	return s.operations.StartLRO(ctx, lroPrefix, lroMetadata, func() (proto.Message, error) {
		lroMetadata.EndTime = timestamppb.Now()
		return obj, nil
	})
}

func (s *WorkflowsV1) populateDefaultsForWorkflow(obj *pb.Workflow) {
}

func (s *WorkflowsV1) DeleteWorkflow(ctx context.Context, req *pb.DeleteWorkflowRequest) (*longrunningpb.Operation, error) {
	name, err := s.parseWorkflowName(req.Name)
	if err != nil {
		return nil, err
	}
	fqn := name.String()

	deleted := &pb.Workflow{}
	if err := s.storage.Delete(ctx, fqn, deleted); err != nil {
		return nil, err
	}

	lroPrefix := fmt.Sprintf("projects/%s/locations/%s", name.Project.ID, name.Location)
	lroMetadata := &pb.OperationMetadata{
		ApiVersion: "v1",
		CreateTime: timestamppb.Now(),
		Target:     fqn,
		Verb:       "delete",
	}
	return s.operations.StartLRO(ctx, lroPrefix, lroMetadata, func() (proto.Message, error) {
		lroMetadata.EndTime = timestamppb.Now()
		return &emptypb.Empty{}, nil
	})
}

func (s *WorkflowsV1) UpdateWorkflow(ctx context.Context, req *pb.UpdateWorkflowRequest) (*longrunningpb.Operation, error) {
	name, err := s.parseWorkflowName(req.GetWorkflow().GetName())
	if err != nil {
		return nil, err
	}
	fqn := name.String()

	existing := &pb.Workflow{}
	if err := s.storage.Get(ctx, fqn, existing); err != nil {
		return nil, err
	}
	now := time.Now()
	updated := proto.Clone(existing).(*pb.Workflow)

	updated.UpdateTime = timestamppb.New(now)
	updated.RevisionCreateTime = timestamppb.New(now)
	updated.RevisionId = "000002-a4d"

	// Required. The update mask applies to the resource.
	paths := req.GetUpdateMask().GetPaths()
	if len(paths) == 0 {
		return nil, status.Errorf(codes.InvalidArgument, "update_mask must be provided")
	}

	// TODO: Some sort of helper for fieldmask?
	for _, path := range paths {
		switch path {
		case "description":
			updated.Description = req.GetWorkflow().GetDescription()
		case "labels":
			updated.Labels = req.GetWorkflow().GetLabels()
		case "source_contents", "sourceContents":
			updated.SourceCode = &pb.Workflow_SourceContents{
				SourceContents: req.GetWorkflow().GetSourceContents(),
			}
		// case "service_account":
		// 	updated.ServiceAccount = req.GetWorkflow().GetServiceAccount()
		// case "crypto_key_name":
		// 	updated.CryptoKeyName = req.GetWorkflow().GetCryptoKeyName()
		// case "call_log_level":
		// 	updated.CallLogLevel = req.GetWorkflow().GetCallLogLevel()
		// case "user_env_vars":
		// 	updated.UserEnvVars = req.GetWorkflow().GetUserEnvVars()
		default:
			return nil, status.Errorf(codes.InvalidArgument, "update_mask path %q not valid", path)
		}
	}

	if err := s.storage.Update(ctx, fqn, updated); err != nil {
		return nil, err
	}

	lroPrefix := fmt.Sprintf("projects/%s/locations/%s", name.Project.ID, name.Location)
	// // Returns with no createTime
	// lroRet := proto.Clone(obj).(*pb.Workflow)
	// lroRet.CreateTime = nil
	// lroRet.UpdateTime = nil
	// lroRet.RevisionCreateTime = nil
	lroMetadata := &pb.OperationMetadata{
		ApiVersion: "v1",
		CreateTime: timestamppb.New(now),
		Target:     fqn,
		Verb:       "update",
	}
	return s.operations.StartLRO(ctx, lroPrefix, lroMetadata, func() (proto.Message, error) {
		lroMetadata.EndTime = timestamppb.Now()
		return updated, nil
	})
}

type workflowName struct {
	Project  *projects.ProjectData
	Location string
	Workflow string
}

func (n *workflowName) String() string {
	return fmt.Sprintf("projects/%s/locations/%s/workflows/%s", n.Project.ID, n.Location, n.Workflow)
}

func (s *MockService) parseWorkflowName(name string) (*workflowName, error) {
	tokens := strings.Split(name, "/")
	if len(tokens) == 6 && tokens[0] == "projects" && tokens[2] == "locations" && tokens[4] == "workflows" {
		project, err := s.Projects.GetProjectByID(tokens[1])
		if err != nil {
			return nil, err
		}

		n := &workflowName{
			Project:  project,
			Location: tokens[3],
			Workflow: tokens[5],
		}
		return n, nil
	}
	return nil, status.Errorf(codes.InvalidArgument, "invalid name %q", name)
}
