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
// proto.service: google.cloud.aiplatform.v1.MetadataService
// proto.message: google.cloud.aiplatform.v1.Execution

package mockaiplatform

import (
	"context"
	"fmt"
	"strings"
	"time"

	longrunning "google.golang.org/genproto/googleapis/longrunning"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/types/known/emptypb"
	"google.golang.org/protobuf/types/known/structpb"
	"google.golang.org/protobuf/types/known/timestamppb"

	pbv1 "cloud.google.com/go/aiplatform/apiv1/aiplatformpb"
	pb "cloud.google.com/go/aiplatform/apiv1beta1/aiplatformpb"
	"github.com/GoogleCloudPlatform/k8s-config-connector/mockgcp/common/projects"
	"github.com/GoogleCloudPlatform/k8s-config-connector/mockgcp/pkg/storage"
	"github.com/google/uuid"
)

type executionName struct {
	Project         *projects.ProjectData
	Location        string
	MetadataStoreId string
	ExecutionId     string
}

func (n *executionName) String() string {
	return fmt.Sprintf("projects/%d/locations/%s/metadataStores/%s/executions/%s", n.Project.Number, n.Location, n.MetadataStoreId, n.ExecutionId)
}

// parseExecutionName parses a string into an executionName.
// Expected form: projects/<projectID>/locations/<location>/metadataStores/<metadataStoreID>/executions/<executionID>
func (s *MockService) parseExecutionName(name string) (*executionName, error) {
	tokens := strings.Split(name, "/")

	if len(tokens) == 8 && tokens[0] == "projects" && tokens[2] == "locations" && tokens[4] == "metadataStores" && tokens[6] == "executions" {
		projectName, err := projects.ParseProjectName(tokens[0] + "/" + tokens[1])
		if err != nil {
			return nil, err
		}
		project, err := s.Projects.GetProject(projectName)
		if err != nil {
			return nil, err
		}

		return &executionName{
			Project:         project,
			Location:        tokens[3],
			MetadataStoreId: tokens[5],
			ExecutionId:     tokens[7],
		}, nil
	}
	return nil, status.Errorf(codes.InvalidArgument, "name %q is not valid", name)
}

func (s *metadataStoreService) GetExecution(ctx context.Context, req *pb.GetExecutionRequest) (*pb.Execution, error) {
	name, err := s.parseExecutionName(req.GetName())
	if err != nil {
		return nil, err
	}

	fqn := name.String()

	obj := &pb.Execution{}
	if err := s.storage.Get(ctx, fqn, obj); err != nil {
		if status.Code(err) == codes.NotFound {
			return nil, status.Errorf(codes.NotFound, "Resource not found.; GetExecution is unable to find execution resource with name: %s", fqn)
		}
		return nil, err
	}

	return obj, nil
}

func (s *metadataStoreService) CreateExecution(ctx context.Context, req *pb.CreateExecutionRequest) (*pb.Execution, error) {
	executionId := req.GetExecutionId()
	if executionId == "" {
		executionId = uuid.NewString()
	}
	reqName := req.GetParent() + "/executions/" + executionId
	name, err := s.parseExecutionName(reqName)
	if err != nil {
		return nil, err
	}

	fqn := name.String()

	now := time.Now()

	obj := proto.Clone(req.GetExecution()).(*pb.Execution)
	obj.Name = fqn
	obj.CreateTime = timestamppb.New(now)
	obj.UpdateTime = timestamppb.New(now)
	if obj.SchemaVersion == "" {
		obj.SchemaVersion = "0.0.1"
	}
	if obj.Metadata == nil {
		obj.Metadata = &structpb.Struct{
			Fields: make(map[string]*structpb.Value),
		}
	}
	obj.Etag = computeEtag(obj)

	if err := s.storage.Create(ctx, fqn, obj); err != nil {
		return nil, err
	}

	return obj, nil
}

func (s *metadataStoreService) UpdateExecution(ctx context.Context, req *pb.UpdateExecutionRequest) (*pb.Execution, error) {
	name, err := s.parseExecutionName(req.GetExecution().GetName())
	if err != nil {
		return nil, err
	}

	fqn := name.String()

	existing := &pb.Execution{}
	if err := s.storage.Get(ctx, fqn, existing); err != nil {
		return nil, err
	}

	updateMask := req.GetUpdateMask()
	if updateMask == nil || len(updateMask.GetPaths()) == 0 {
		return nil, status.Errorf(codes.InvalidArgument, "update_mask is required")
	}

	for _, path := range updateMask.GetPaths() {
		switch path {
		case "description":
			existing.Description = req.GetExecution().GetDescription()
		case "displayName", "display_name":
			existing.DisplayName = req.GetExecution().GetDisplayName()
		case "metadata":
			existing.Metadata = req.GetExecution().GetMetadata()
		case "state":
			existing.State = req.GetExecution().GetState()
		case "labels":
			existing.Labels = req.GetExecution().GetLabels()
		case "schemaTitle", "schema_title":
			existing.SchemaTitle = req.GetExecution().GetSchemaTitle()
		case "schemaVersion", "schema_version":
			existing.SchemaVersion = req.GetExecution().GetSchemaVersion()
		case "etag":
			// Etag check/update
		default:
			return nil, status.Errorf(codes.InvalidArgument, "field %q is not updatable", path)
		}
	}

	now := time.Now()
	existing.UpdateTime = timestamppb.New(now)
	existing.Etag = computeEtag(existing)

	if err := s.storage.Update(ctx, fqn, existing); err != nil {
		return nil, err
	}

	return existing, nil
}

func (s *metadataStoreService) DeleteExecution(ctx context.Context, req *pb.DeleteExecutionRequest) (*longrunning.Operation, error) {
	name, err := s.parseExecutionName(req.GetName())
	if err != nil {
		return nil, err
	}

	fqn := name.String()

	existing := &pb.Execution{}
	if err := s.storage.Get(ctx, fqn, existing); err != nil {
		return nil, err
	}

	if err := s.storage.Delete(ctx, fqn, existing); err != nil {
		return nil, err
	}

	now := time.Now()
	op := &pbv1.DeleteOperationMetadata{
		GenericMetadata: &pbv1.GenericOperationMetadata{
			CreateTime: timestamppb.New(now),
			UpdateTime: timestamppb.New(now),
		},
	}
	opPrefix := fqn
	return s.operations.DoneLRO(ctx, opPrefix, op, &emptypb.Empty{})
}

func (s *metadataStoreService) ListExecutions(ctx context.Context, req *pb.ListExecutionsRequest) (*pb.ListExecutionsResponse, error) {
	tokens := strings.Split(req.GetParent(), "/")
	if len(tokens) != 6 || tokens[0] != "projects" || tokens[2] != "locations" || tokens[4] != "metadataStores" {
		return nil, status.Errorf(codes.InvalidArgument, "parent %q is not valid", req.GetParent())
	}
	projectName, err := projects.ParseProjectName(tokens[0] + "/" + tokens[1])
	if err != nil {
		return nil, err
	}
	project, err := s.Projects.GetProject(projectName)
	if err != nil {
		return nil, err
	}

	findPrefix := fmt.Sprintf("projects/%d/locations/%s/metadataStores/%s/executions/", project.Number, tokens[3], tokens[5])

	var executions []*pb.Execution
	executionKind := (&pb.Execution{}).ProtoReflect().Descriptor()
	if err := s.storage.List(ctx, executionKind, storage.ListOptions{}, func(obj proto.Message) error {
		execution := obj.(*pb.Execution)
		if strings.HasPrefix(execution.GetName(), findPrefix) {
			executions = append(executions, execution)
		}
		return nil
	}); err != nil {
		return nil, err
	}

	return &pb.ListExecutionsResponse{
		Executions: executions,
	}, nil
}
