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
// proto.service: google.cloud.apphub.v1.AppHub
// proto.message: google.cloud.apphub.v1.Application

package mockapphub

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

	"github.com/GoogleCloudPlatform/k8s-config-connector/mockgcp/common/projects"
	pb "github.com/GoogleCloudPlatform/k8s-config-connector/mockgcp/generated/mockgcp/cloud/apphub/v1"
	longrunningpb "google.golang.org/genproto/googleapis/longrunning"
)

type AppHubV1Service struct {
	*MockService
	pb.UnimplementedAppHubServer
}

func (s *AppHubV1Service) GetApplication(ctx context.Context, req *pb.GetApplicationRequest) (*pb.Application, error) {
	name, err := s.parseApplicationName(req.Name)
	if err != nil {
		return nil, err
	}

	fqn := name.String()

	obj := &pb.Application{}
	if err := s.storage.Get(ctx, fqn, obj); err != nil {
		return nil, err
	}

	return obj, nil
}

func (s *AppHubV1Service) UpdateApplication(ctx context.Context, req *pb.UpdateApplicationRequest) (*longrunning.Operation, error) {
	name, err := s.parseApplicationName(req.GetApplication().GetName())
	if err != nil {
		return nil, err
	}
	fqn := name.String()

	obj := &pb.Application{}
	if err := s.storage.Get(ctx, fqn, obj); err != nil {
		return nil, err
	}
	obj.UpdateTime = timestamppb.New(time.Now())
	paths := req.GetUpdateMask().GetPaths()
	for _, path := range paths {
		switch path {
		case "description":
			obj.Description = req.GetApplication().GetDescription()
		case "display_name":
			obj.DisplayName = req.GetApplication().GetDisplayName()
		case "attributes":
			obj.Attributes = req.GetApplication().GetAttributes()
		default:
			return nil, status.Errorf(codes.InvalidArgument, "update_mask path %q not valid", path)
		}
	}
	if err := s.storage.Update(ctx, fqn, obj); err != nil {
		return nil, err
	}
	lroPrefix := fmt.Sprintf("projects/%s/locations/%s", name.Project.ID, name.Location)
	lroMetadata := &pb.OperationMetadata{
		ApiVersion:    "v1",
		CreateTime:    timestamppb.Now(),
		Target:        name.String(),
		Verb:          "update",
		EndTime:       timestamppb.Now(),
		StatusMessage: "Updating application",
	}
	return s.operations.StartLRO(ctx, lroPrefix, lroMetadata, func() (proto.Message, error) {
		lroMetadata.EndTime = timestamppb.Now()
		return obj, nil
	})
}

func (s *AppHubV1Service) CreateApplication(ctx context.Context, req *pb.CreateApplicationRequest) (*longrunningpb.Operation, error) {
	reqName := req.Parent + "/applications/" + req.ApplicationId
	name, err := s.parseApplicationName(reqName)
	if err != nil {
		return nil, err
	}

	fqn := name.String()

	obj := proto.Clone(req.Application).(*pb.Application)
	obj.Name = fqn
	obj.CreateTime = timestamppb.New(time.Now())
	obj.UpdateTime = timestamppb.New(time.Now())
	obj.Uid = "7683c772-1f55-4270-a0ea-8b49c9f42d00" // TODO: generate a unique UUID
	obj.State = pb.Application_ACTIVE

	if err := s.storage.Create(ctx, fqn, obj); err != nil {
		return nil, err
	}

	prefix := fmt.Sprintf("projects/%s/locations/%s", name.Project.ID, name.Location)
	lroMetadata := &pb.OperationMetadata{
		Target:     name.String(),
		Verb:       "create",
		ApiVersion: "v1",
		CreateTime: timestamppb.New(time.Now()),
	}
	return s.operations.StartLRO(ctx, prefix, lroMetadata, func() (proto.Message, error) {
		lroMetadata.EndTime = timestamppb.Now()
		return obj, nil
	})
}

func (s *AppHubV1Service) DeleteApplication(ctx context.Context, req *pb.DeleteApplicationRequest) (*longrunningpb.Operation, error) {
	name, err := s.parseApplicationName(req.Name)
	if err != nil {
		return nil, err
	}

	fqn := name.String()

	deleted := &pb.Application{}
	if err := s.storage.Delete(ctx, fqn, deleted); err != nil {
		return nil, err
	}

	prefix := fmt.Sprintf("projects/%s/locations/%s", name.Project.ID, name.Location)
	lroMetadata := &pb.OperationMetadata{
		Target:     name.String(),
		Verb:       "delete",
		ApiVersion: "v1",
		CreateTime: timestamppb.New(time.Now()),
	}
	return s.operations.StartLRO(ctx, prefix, lroMetadata, func() (proto.Message, error) {
		lroMetadata.EndTime = timestamppb.New(time.Now())
		return &emptypb.Empty{}, nil
	})
}

type applicationName struct {
	Project         *projects.ProjectData
	Location        string
	ApplicationName string
}

func (n *applicationName) String() string {
	return fmt.Sprintf("projects/%s/locations/%s/applications/%s", n.Project.ID, n.Location, n.ApplicationName)
}

// parseApplicationName parses a string into an applicationName.
// The expected form is `projects/*/locations/*/applications/*`.
func (s *AppHubV1Service) parseApplicationName(name string) (*applicationName, error) {
	tokens := strings.Split(name, "/")

	if len(tokens) == 6 && tokens[0] == "projects" && tokens[2] == "locations" && tokens[4] == "applications" {
		project, err := s.Projects.GetProjectByID(tokens[1])
		if err != nil {
			return nil, err
		}

		name := &applicationName{
			Project:         project,
			Location:        tokens[3],
			ApplicationName: tokens[5],
		}

		return name, nil
	}

	return nil, status.Errorf(codes.InvalidArgument, "name %q is not valid", name)
}
