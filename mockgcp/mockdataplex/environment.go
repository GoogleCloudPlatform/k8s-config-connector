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
// proto.service: google.cloud.dataplex.v1.DataplexService
// proto.message: google.cloud.dataplex.v1.Environment

package mockdataplex

import (
	"context"
	"fmt"
	"strings"
	"time"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/types/known/durationpb"
	"google.golang.org/protobuf/types/known/emptypb"
	"google.golang.org/protobuf/types/known/timestamppb"

	"github.com/GoogleCloudPlatform/k8s-config-connector/mockgcp/common/projects"
	pb "github.com/GoogleCloudPlatform/k8s-config-connector/mockgcp/generated/mockgcp/cloud/dataplex/v1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/mockgcp/pkg/storage"
	"github.com/google/uuid"
	longrunningpb "google.golang.org/genproto/googleapis/longrunning"
)

func (s *DataplexV1) GetEnvironment(ctx context.Context, req *pb.GetEnvironmentRequest) (*pb.Environment, error) {
	name, err := s.parseEnvironmentName(req.Name)
	if err != nil {
		return nil, err
	}

	fqn := name.String()

	obj := &pb.Environment{}
	if err := s.storage.Get(ctx, fqn, obj); err != nil {
		if status.Code(err) == codes.NotFound {
			return nil, status.Errorf(codes.NotFound, "Resource '%s' was not found", fqn)
		}
		return nil, err
	}

	return obj, nil
}

func (s *DataplexV1) CreateEnvironment(ctx context.Context, req *pb.CreateEnvironmentRequest) (*longrunningpb.Operation, error) {
	reqName := fmt.Sprintf("%s/environments/%s", req.GetParent(), req.GetEnvironmentId())
	name, err := s.parseEnvironmentName(reqName)
	if err != nil {
		return nil, err
	}

	fqn := name.String()

	now := time.Now()

	obj := proto.Clone(req.GetEnvironment()).(*pb.Environment)
	obj.Name = fqn
	obj.Uid = uuid.NewString()
	obj.CreateTime = timestamppb.New(now)
	obj.UpdateTime = timestamppb.New(now)
	obj.State = pb.State_ACTIVE
	s.populateEnvironmentDefaults(obj)

	if err := s.storage.Create(ctx, fqn, obj); err != nil {
		return nil, err
	}

	prefix := fmt.Sprintf("projects/%s/locations/%s", name.Project.ID, name.Location)
	metadata := &pb.OperationMetadata{
		CreateTime: timestamppb.New(now),
		Target:     fqn,
		Verb:       "create",
		ApiVersion: "v1",
	}
	return s.operations.StartLRO(ctx, prefix, metadata, func() (proto.Message, error) {
		metadata.EndTime = timestamppb.Now()
		return obj, nil
	})
}

func (s *DataplexV1) UpdateEnvironment(ctx context.Context, req *pb.UpdateEnvironmentRequest) (*longrunningpb.Operation, error) {
	name, err := s.parseEnvironmentName(req.GetEnvironment().GetName())
	if err != nil {
		return nil, err
	}
	fqn := name.String()

	obj := &pb.Environment{}
	if err := s.storage.Get(ctx, fqn, obj); err != nil {
		return nil, err
	}

	now := time.Now()

	paths := req.GetUpdateMask().GetPaths()
	if len(paths) == 0 {
		return nil, status.Errorf(codes.InvalidArgument, "update_mask must be provided")
	}

	// Only care about fields we support for update.
	for _, path := range paths {
		switch path {
		case "description":
			obj.Description = req.GetEnvironment().GetDescription()
		case "display_name":
			obj.DisplayName = req.GetEnvironment().GetDisplayName()
		case "labels":
			obj.Labels = req.GetEnvironment().GetLabels()
		case "infrastructure_spec":
			obj.InfrastructureSpec = req.GetEnvironment().GetInfrastructureSpec()
		case "session_spec":
			obj.SessionSpec = req.GetEnvironment().GetSessionSpec()
		default:
			return nil, status.Errorf(codes.InvalidArgument, "update_mask path %q not valid", path)
		}
	}

	obj.UpdateTime = timestamppb.New(now)

	if err := s.storage.Update(ctx, fqn, obj); err != nil {
		return nil, err
	}

	prefix := fmt.Sprintf("projects/%s/locations/%s", name.Project.ID, name.Location)
	metadata := &pb.OperationMetadata{
		CreateTime: timestamppb.New(now),
		Target:     fqn,
		Verb:       "update",
		ApiVersion: "v1",
	}
	return s.operations.StartLRO(ctx, prefix, metadata, func() (proto.Message, error) {
		metadata.EndTime = timestamppb.Now()
		return obj, nil
	})
}

func (s *DataplexV1) DeleteEnvironment(ctx context.Context, req *pb.DeleteEnvironmentRequest) (*longrunningpb.Operation, error) {
	name, err := s.parseEnvironmentName(req.Name)
	if err != nil {
		return nil, err
	}

	fqn := name.String()

	deleted := &pb.Environment{}
	if err := s.storage.Delete(ctx, fqn, deleted); err != nil {
		return nil, err
	}

	prefix := fmt.Sprintf("projects/%s/locations/%s", name.Project.ID, name.Location)
	now := time.Now()
	metadata := &pb.OperationMetadata{
		CreateTime: timestamppb.New(now),
		Target:     fqn,
		Verb:       "delete",
		ApiVersion: "v1",
	}
	return s.operations.StartLRO(ctx, prefix, metadata, func() (proto.Message, error) {
		metadata.EndTime = timestamppb.New(now)
		return &emptypb.Empty{}, nil
	})
}

func (s *DataplexV1) ListEnvironments(ctx context.Context, req *pb.ListEnvironmentsRequest) (*pb.ListEnvironmentsResponse, error) {
	parent, err := s.parseLakeName(req.Parent)
	if err != nil {
		return nil, err
	}

	response := &pb.ListEnvironmentsResponse{}

	environmentKind := (&pb.Environment{}).ProtoReflect().Descriptor()
	if err := s.storage.List(ctx, environmentKind, storage.ListOptions{}, func(obj proto.Message) error {
		env := obj.(*pb.Environment)
		if strings.HasPrefix(env.GetName(), parent.String()+"/environments/") {
			response.Environments = append(response.Environments, env)
		}
		return nil
	}); err != nil {
		return nil, err
	}

	return response, nil
}

func (s *DataplexV1) populateEnvironmentDefaults(obj *pb.Environment) {
	if obj.SessionSpec == nil {
		obj.SessionSpec = &pb.Environment_SessionSpec{}
	}
	if obj.SessionSpec.MaxIdleDuration == nil {
		obj.SessionSpec.MaxIdleDuration = durationpb.New(3 * time.Hour) // 10800s
	}
	if obj.SessionStatus == nil {
		obj.SessionStatus = &pb.Environment_SessionStatus{Active: true}
	}
	if obj.InfrastructureSpec == nil {
		obj.InfrastructureSpec = &pb.Environment_InfrastructureSpec{}
	}
	if obj.InfrastructureSpec.GetOsImage() == nil {
		obj.InfrastructureSpec.Runtime = &pb.Environment_InfrastructureSpec_OsImage{
			OsImage: &pb.Environment_InfrastructureSpec_OsImageRuntime{
				ImageVersion: "1.0.0", // Default if not provided?
			},
		}
	}
	if obj.InfrastructureSpec.GetCompute() == nil {
		obj.InfrastructureSpec.Resources = &pb.Environment_InfrastructureSpec_Compute{
			Compute: &pb.Environment_InfrastructureSpec_ComputeResources{
				DiskSizeGb: 100, // Default disk size
			},
		}
	}
	if obj.Endpoints == nil {
		obj.Endpoints = &pb.Environment_Endpoints{
			Notebooks: fmt.Sprintf("https://notebooks.dataplex.goog/%s", obj.Name), // Example format
			Sql:       fmt.Sprintf("https://sql.dataplex.goog/%s", obj.Name),       // Example format
		}
	}
}

type environmentName struct {
	Project       *projects.ProjectData
	Location      string
	LakeID        string
	EnvironmentID string
}

func (n *environmentName) String() string {
	return fmt.Sprintf("projects/%s/locations/%s/lakes/%s/environments/%s", n.Project.ID, n.Location, n.LakeID, n.EnvironmentID)
}

// parseEnvironmentName parses a string into an environmentName.
// The expected form is `projects/*/locations/*/lakes/*/environments/*`.
func (s *MockService) parseEnvironmentName(name string) (*environmentName, error) {
	tokens := strings.Split(name, "/")

	if len(tokens) == 8 && tokens[0] == "projects" && tokens[2] == "locations" && tokens[4] == "lakes" && tokens[6] == "environments" {
		project, err := s.Projects.GetProjectByID(tokens[1])
		if err != nil {
			return nil, err
		}

		name := &environmentName{
			Project:       project,
			Location:      tokens[3],
			LakeID:        tokens[5],
			EnvironmentID: tokens[7],
		}

		return name, nil
	}

	return nil, status.Errorf(codes.InvalidArgument, "name %q is not valid", name)
}
