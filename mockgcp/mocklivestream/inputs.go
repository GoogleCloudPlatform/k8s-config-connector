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

package mocklivestream

import (
	"context"
	"fmt"
	"strings"
	"time"

	"google.golang.org/genproto/googleapis/longrunning"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/types/known/emptypb"
	"google.golang.org/protobuf/types/known/timestamppb"

	pb "cloud.google.com/go/video/livestream/apiv1/livestreampb"
	"github.com/GoogleCloudPlatform/k8s-config-connector/mockgcp/common/projects"
	"github.com/GoogleCloudPlatform/k8s-config-connector/mockgcp/pkg/storage"
)

type LivestreamServer struct {
	*MockService
	pb.UnimplementedLivestreamServiceServer
}

type inputName struct {
	Project  *projects.ProjectData
	Location string
	Input    string
}

func (n *inputName) String() string {
	return "projects/" + n.Project.ID + "/locations/" + n.Location + "/inputs/" + n.Input
}

func (s *MockService) parseInputName(name string) (*inputName, error) {
	tokens := strings.Split(name, "/")

	if len(tokens) == 6 && tokens[0] == "projects" && tokens[2] == "locations" && tokens[4] == "inputs" {
		project, err := s.Projects.GetProjectByID(tokens[1])
		if err != nil {
			return nil, err
		}

		return &inputName{
			Project:  project,
			Location: tokens[3],
			Input:    tokens[5],
		}, nil
	}
	return nil, status.Errorf(codes.InvalidArgument, "name %q is not valid", name)
}

func (s *LivestreamServer) GetInput(ctx context.Context, req *pb.GetInputRequest) (*pb.Input, error) {
	name, err := s.parseInputName(req.Name)
	if err != nil {
		return nil, err
	}

	fqn := name.String()

	obj := &pb.Input{}
	if err := s.storage.Get(ctx, fqn, obj); err != nil {
		if status.Code(err) == codes.NotFound {
			return nil, status.Errorf(codes.NotFound, "Resource '%s' was not found", fqn)
		}
		return nil, err
	}

	return obj, nil
}

func (s *LivestreamServer) CreateInput(ctx context.Context, req *pb.CreateInputRequest) (*longrunning.Operation, error) {
	reqName := req.Parent + "/inputs/" + req.InputId
	name, err := s.parseInputName(reqName)
	if err != nil {
		return nil, err
	}

	now := time.Now()
	fqn := name.String()

	obj := proto.Clone(req.Input).(*pb.Input)
	obj.Name = fqn
	obj.CreateTime = timestamppb.New(now)
	obj.UpdateTime = timestamppb.New(now)

	if obj.Tier == pb.Input_TIER_UNSPECIFIED {
		obj.Tier = pb.Input_HD
	}

	if obj.Uri == "" {
		if obj.Type == pb.Input_SRT_PUSH {
			obj.Uri = "srt://35.239.41.154/live/9fa80b77-5819-4ca3-aeba-ca41dffe4636"
		} else {
			obj.Uri = "rtmp://35.239.41.154/live/9fa80b77-5819-4ca3-aeba-ca41dffe4636"
		}
	}

	if err := s.storage.Create(ctx, fqn, obj); err != nil {
		return nil, err
	}

	opMetadata := &pb.OperationMetadata{
		ApiVersion: "v1",
		CreateTime: timestamppb.New(now),
		Verb:       "create",
		Target:     fqn,
	}
	opPrefix := fmt.Sprintf("projects/%s/locations/%s", name.Project.ID, name.Location)
	return s.operations.StartLRO(ctx, opPrefix, opMetadata, func() (proto.Message, error) {
		opMetadata.EndTime = timestamppb.Now()
		return obj, nil
	})
}

func (s *LivestreamServer) UpdateInput(ctx context.Context, req *pb.UpdateInputRequest) (*longrunning.Operation, error) {
	reqName := req.GetInput().GetName()
	name, err := s.parseInputName(reqName)
	if err != nil {
		return nil, err
	}

	fqn := name.String()

	existing := &pb.Input{}
	if err := s.storage.Get(ctx, fqn, existing); err != nil {
		if status.Code(err) == codes.NotFound {
			return nil, status.Errorf(codes.NotFound, "Resource '%s' was not found", fqn)
		}
		return nil, err
	}

	now := time.Now()
	updated := proto.Clone(existing).(*pb.Input)

	paths := req.GetUpdateMask().GetPaths()
	if len(paths) == 0 {
		updated = proto.Clone(req.GetInput()).(*pb.Input)
		updated.Name = fqn
		updated.CreateTime = existing.CreateTime
		updated.Uri = existing.Uri
	} else {
		for _, path := range paths {
			switch path {
			case "preprocessing_config", "preprocessingConfig":
				updated.PreprocessingConfig = req.GetInput().GetPreprocessingConfig()
			case "security_rules", "securityRules":
				updated.SecurityRules = req.GetInput().GetSecurityRules()
			case "labels":
				updated.Labels = req.GetInput().GetLabels()
			case "tier":
				updated.Tier = req.GetInput().GetTier()
			default:
				return nil, status.Errorf(codes.InvalidArgument, "field %q not supported for update", path)
			}
		}
	}
	updated.UpdateTime = timestamppb.New(now)

	if err := s.storage.Update(ctx, fqn, updated); err != nil {
		return nil, err
	}

	opMetadata := &pb.OperationMetadata{
		ApiVersion: "v1",
		CreateTime: timestamppb.New(now),
		Verb:       "update",
		Target:     fqn,
	}
	opPrefix := fmt.Sprintf("projects/%s/locations/%s", name.Project.ID, name.Location)
	return s.operations.StartLRO(ctx, opPrefix, opMetadata, func() (proto.Message, error) {
		opMetadata.EndTime = timestamppb.Now()
		return updated, nil
	})
}

func (s *LivestreamServer) DeleteInput(ctx context.Context, req *pb.DeleteInputRequest) (*longrunning.Operation, error) {
	name, err := s.parseInputName(req.Name)
	if err != nil {
		return nil, err
	}

	fqn := name.String()

	deleted := &pb.Input{}
	if err := s.storage.Delete(ctx, fqn, deleted); err != nil {
		if status.Code(err) == codes.NotFound {
			return nil, status.Errorf(codes.NotFound, "Resource '%s' was not found", fqn)
		}
		return nil, err
	}

	now := time.Now()
	opMetadata := &pb.OperationMetadata{
		ApiVersion: "v1",
		CreateTime: timestamppb.New(now),
		Verb:       "delete",
		Target:     fqn,
	}
	opPrefix := fmt.Sprintf("projects/%s/locations/%s", name.Project.ID, name.Location)
	return s.operations.StartLRO(ctx, opPrefix, opMetadata, func() (proto.Message, error) {
		opMetadata.EndTime = timestamppb.Now()
		return &emptypb.Empty{}, nil
	})
}

func (s *LivestreamServer) ListInputs(ctx context.Context, req *pb.ListInputsRequest) (*pb.ListInputsResponse, error) {
	response := &pb.ListInputsResponse{}
	findPrefix := req.Parent + "/inputs/"

	inputKind := (&pb.Input{}).ProtoReflect().Descriptor()
	if err := s.storage.List(ctx, inputKind, storage.ListOptions{
		Prefix: findPrefix,
	}, func(obj proto.Message) error {
		response.Inputs = append(response.Inputs, obj.(*pb.Input))
		return nil
	}); err != nil {
		return nil, err
	}

	return response, nil
}
