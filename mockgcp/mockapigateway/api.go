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
// proto.service: google.cloud.apigateway.v1.ApiGatewayService
// proto.message: google.cloud.apigateway.v1.Api

package mockapigateway

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
	pb "github.com/GoogleCloudPlatform/k8s-config-connector/mockgcp/generated/mockgcp/cloud/apigateway/v1"
	longrunningpb "google.golang.org/genproto/googleapis/longrunning"
)

func (s *ApiGatewayV1) GetApi(ctx context.Context, req *pb.GetApiRequest) (*pb.Api, error) {
	name, err := s.parseAPIName(req.Name)
	if err != nil {
		return nil, err
	}

	fqn := name.String()

	obj := &pb.Api{}
	if err := s.storage.Get(ctx, fqn, obj); err != nil {
		if status.Code(err) == codes.NotFound {
			return nil, status.Errorf(codes.NotFound, "API %q not found", fqn)
		}
		return nil, err
	}

	return obj, nil
}

func (s *ApiGatewayV1) CreateApi(ctx context.Context, req *pb.CreateApiRequest) (*longrunningpb.Operation, error) {
	reqName := fmt.Sprintf("%s/apis/%s", req.GetParent(), req.GetApiId())
	name, err := s.parseAPIName(reqName)
	if err != nil {
		return nil, err
	}

	fqn := name.String()

	now := time.Now()

	obj := proto.Clone(req.GetApi()).(*pb.Api)
	obj.Name = fqn
	obj.CreateTime = timestamppb.New(now)
	obj.UpdateTime = timestamppb.New(now)
	obj.State = pb.Api_ACTIVE
	obj.DisplayName = name.Api

	if err := s.storage.Create(ctx, fqn, obj); err != nil {
		return nil, err
	}

	lroPrefix := name.Parent()
	lroMetadata := &pb.OperationMetadata{
		CreateTime: timestamppb.New(now),
		Target:     name.String(),
		Verb:       "create",
		ApiVersion: "v1",
	}
	return s.operations.StartLRO(ctx, lroPrefix, lroMetadata, func() (proto.Message, error) {
		lroMetadata.EndTime = timestamppb.New(time.Now())
		return obj, nil
	})
}

func (s *ApiGatewayV1) UpdateApi(ctx context.Context, req *pb.UpdateApiRequest) (*longrunningpb.Operation, error) {
	name, err := s.parseAPIName(req.GetApi().GetName())
	if err != nil {
		return nil, err
	}
	fqn := name.String()

	obj := &pb.Api{}
	if err := s.storage.Get(ctx, fqn, obj); err != nil {
		return nil, err
	}

	reqObj := req.GetApi()
	paths := req.GetUpdateMask().GetPaths()
	if len(paths) == 0 {
		return nil, status.Errorf(codes.InvalidArgument, "update_mask must be provided")
	}

	// TODO: Some sort of helper for fieldmask?
	for _, path := range paths {
		switch path {
		case "display_name", "displayName":
			obj.DisplayName = reqObj.GetDisplayName()
		case "labels":
			obj.Labels = reqObj.GetLabels()
		default:
			return nil, status.Errorf(codes.InvalidArgument, "update_mask path %q not valid", path)
		}
	}

	obj.UpdateTime = timestamppb.New(time.Now())

	if err := s.storage.Update(ctx, fqn, obj); err != nil {
		return nil, err
	}

	lroPrefix := name.Parent()
	lroMetadata := &pb.OperationMetadata{
		CreateTime: obj.CreateTime,
		Target:     name.String(),
		Verb:       "update",
		ApiVersion: "v1",
	}
	return s.operations.StartLRO(ctx, lroPrefix, lroMetadata, func() (proto.Message, error) {
		lroMetadata.EndTime = timestamppb.Now()
		return obj, nil
	})
}

func (s *ApiGatewayV1) DeleteApi(ctx context.Context, req *pb.DeleteApiRequest) (*longrunningpb.Operation, error) {
	name, err := s.parseAPIName(req.Name)
	if err != nil {
		return nil, err
	}
	fqn := name.String()

	deleted := &pb.Api{}
	if err := s.storage.Delete(ctx, fqn, deleted); err != nil {
		return nil, err
	}

	lroPrefix := name.Parent()
	lroMetadata := &pb.OperationMetadata{
		CreateTime: timestamppb.Now(),
		Target:     name.String(),
		Verb:       "delete",
		ApiVersion: "v1",
	}

	return s.operations.StartLRO(ctx, lroPrefix, lroMetadata, func() (proto.Message, error) {
		lroMetadata.EndTime = timestamppb.Now()
		return &emptypb.Empty{}, nil
	})
}

type apiName struct {
	Project  *projects.ProjectData
	Location string
	Api      string
}

func (n *apiName) Parent() string {
	return fmt.Sprintf("projects/%s/locations/%s", n.Project.ID, n.Location)
}

func (n *apiName) String() string {
	return fmt.Sprintf("%s/apis/%s", n.Parent(), n.Api)
}

// parseAPIName parses a string into an apiName.
// The expected form is `projects/*/locations/global/apis/*`.
func (s *MockService) parseAPIName(name string) (*apiName, error) {
	tokens := strings.Split(name, "/")

	if len(tokens) == 6 && tokens[0] == "projects" && tokens[2] == "locations" && tokens[4] == "apis" {
		if tokens[3] != "global" {
			return nil, status.Errorf(codes.InvalidArgument, "only 'global' location is supported, but got %q", tokens[3])
		}
		projectName, err := projects.ParseProjectName(tokens[0] + "/" + tokens[1])
		if err != nil {
			return nil, err
		}
		project, err := s.Projects.GetProject(projectName)
		if err != nil {
			return nil, err
		}

		name := &apiName{
			Project:  project,
			Location: tokens[3],
			Api:      tokens[5],
		}

		return name, nil
	}
	return nil, status.Errorf(codes.InvalidArgument, "name %q is not valid", name)
}
