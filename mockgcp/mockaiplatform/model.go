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
// proto.service: google.cloud.aiplatform.v1.ModelService
// proto.message: google.cloud.aiplatform.v1.Model

package mockaiplatform

import (
	"context"
	"strings"
	"time"

	longrunning "google.golang.org/genproto/googleapis/longrunning"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/types/known/emptypb"
	"google.golang.org/protobuf/types/known/timestamppb"

	"github.com/GoogleCloudPlatform/k8s-config-connector/mockgcp/common/projects"
	pb "github.com/GoogleCloudPlatform/k8s-config-connector/mockgcp/generated/mockgcp/cloud/aiplatform/v1beta1"
)

type modelService struct {
	*MockService
	pb.UnimplementedModelServiceServer
}

func (s *modelService) UploadModel(ctx context.Context, req *pb.UploadModelRequest) (*longrunning.Operation, error) {
	reqName := req.Parent + "/models/" + "mockgenerated"
	name, err := s.parseModelName(reqName)
	if err != nil {
		return nil, err
	}

	fqn := name.String()

	now := time.Now()

	obj := proto.Clone(req.Model).(*pb.Model)
	obj.Name = fqn

	obj.CreateTime = timestamppb.New(now)
	obj.UpdateTime = timestamppb.New(now)

	if err := s.storage.Create(ctx, fqn, obj); err != nil {
		return nil, err
	}

	op := &pb.UploadModelOperationMetadata{}
	op.GenericMetadata = &pb.GenericOperationMetadata{
		CreateTime: timestamppb.New(now),
		UpdateTime: timestamppb.New(now),
	}
	opPrefix := name.String()
	return s.operations.StartLRO(ctx, opPrefix, op, func() (proto.Message, error) {
		// Many fields are not populated in the LRO result
		result := &pb.UploadModelResponse{
			Model: obj.Name,
		}
		return result, nil
	})
}

type ModelName struct {
	Project  *projects.ProjectData
	Location string
	ModelID  string
}

func (n *ModelName) String() string {
	return "projects/" + n.Project.ID + "/locations/" + n.Location + "/models/" + n.ModelID
}

// parseModelName parses a string into a modelName.
// The expected form is `projects/*/locations/*/models/*`.
func (s *MockService) parseModelName(name string) (*ModelName, error) {
	tokens := strings.Split(name, "/")
	if len(tokens) == 6 && tokens[0] == "projects" && tokens[2] == "locations" && tokens[4] == "models" {
		project, err := s.Projects.GetProjectByID(tokens[1])
		if err != nil {
			return nil, err
		}

		name := &ModelName{
			Project:  project,
			Location: tokens[3],
			ModelID:  tokens[5],
		}

		return name, nil
	}

	return nil, status.Errorf(codes.InvalidArgument, "name %q is not valid", name)
}

func (s *modelService) ListModels(ctx context.Context, req *pb.ListModelsRequest) (*pb.ListModelsResponse, error) {
	_, err := s.parseModelName(req.Parent + "/models/" + "id-can-be-anything")
	if err != nil {
		return nil, err
	}

	var models []*pb.Model
	return &pb.ListModelsResponse{Models: models}, nil
}

func (s *modelService) GetModel(ctx context.Context, req *pb.GetModelRequest) (*pb.Model, error) {
	name, err := s.parseModelName(req.Name)
	if err != nil {
		return nil, err
	}

	fqn := name.String()

	obj := &pb.Model{}
	if err := s.storage.Get(ctx, fqn, obj); err != nil {
		return nil, err
	}

	return obj, nil
}

func (s *modelService) DeleteModel(ctx context.Context, req *pb.DeleteModelRequest) (*longrunning.Operation, error) {
	name, err := s.parseModelName(req.Name)
	if err != nil {
		return nil, err
	}

	fqn := name.String()
	now := time.Now()

	deleted := &pb.Model{}
	if err := s.storage.Delete(ctx, fqn, deleted); err != nil {
		return nil, err
	}

	op := &pb.DeleteOperationMetadata{}
	op.GenericMetadata = &pb.GenericOperationMetadata{
		CreateTime: timestamppb.New(now),
		UpdateTime: timestamppb.New(now),
	}
	opPrefix := name.String()

	return s.operations.DoneLRO(ctx, opPrefix, op, &emptypb.Empty{})
}

func (s *modelService) UpdateModel(ctx context.Context, req *pb.UpdateModelRequest) (*pb.Model, error) {
	name, err := s.parseModelName(req.GetModel().Name)
	if err != nil {
		return nil, err
	}

	fqn := name.String()

	obj := &pb.Model{}
	if err := s.storage.Get(ctx, fqn, obj); err != nil {
		return nil, err
	}

	// Apply the update mask to the object.
	paths := req.GetUpdateMask().GetPaths()
	if len(paths) == 0 {
		return nil, status.Errorf(codes.InvalidArgument, "update_mask must be provided")
	}
	for _, path := range paths {
		switch path {
		case "displayName":
			obj.DisplayName = req.GetModel().GetDisplayName()
		case "description":
			obj.Description = req.GetModel().GetDescription()
		case "labels":
			obj.Labels = req.GetModel().GetLabels()
		default:
			return nil, status.Errorf(codes.InvalidArgument, "update_mask path %q not valid", path)
		}
	}
	// Version ID and alias can't be updated
	obj.VersionId = req.GetModel().GetVersionId()
	obj.VersionAliases = req.GetModel().GetVersionAliases()
	obj.VersionCreateTime = req.GetModel().GetVersionCreateTime()
	obj.VersionUpdateTime = timestamppb.New(time.Now())
	obj.UpdateTime = timestamppb.New(time.Now())

	if err := s.storage.Update(ctx, fqn, obj); err != nil {
		return nil, err
	}

	return obj, nil
}
