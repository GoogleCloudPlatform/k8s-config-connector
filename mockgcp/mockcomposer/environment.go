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
// proto.service: google.cloud.orchestration.airflow.service.v1.Environments
// proto.message: google.cloud.orchestration.airflow.service.v1.Environment

package mockcomposer

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
	pb "github.com/GoogleCloudPlatform/k8s-config-connector/mockgcp/generated/mockgcp/cloud/orchestration/airflow/service/v1"
	longrunningpb "google.golang.org/genproto/googleapis/longrunning"
)

func (s *ComposerV1) GetEnvironment(ctx context.Context, req *pb.GetEnvironmentRequest) (*pb.Environment, error) {
	name, err := s.parseEnvironmentName(req.Name)
	if err != nil {
		return nil, err
	}

	fqn := name.String()

	obj := &pb.Environment{}
	if err := s.storage.Get(ctx, fqn, obj); err != nil {
		return nil, err
	}
	return obj, nil
}

func (s *ComposerV1) CreateEnvironment(ctx context.Context, req *pb.CreateEnvironmentRequest) (*longrunningpb.Operation, error) {
	// The parent field is projects/{projectId}/locations/{locationId}.
	name, err := s.parseParentEnvironment(req.Parent)
	if err != nil {
		return nil, err
	}
	obj := proto.Clone(req.GetEnvironment()).(*pb.Environment)
	// The resource name of the environment.
	obj.Name = fmt.Sprintf("projects/%s/locations/%s/environments/%s", name.Project.ID, name.Location, "TEMP")

	obj.CreateTime = timestamppb.New(time.Now())
	obj.UpdateTime = timestamppb.New(time.Now())

	s.populateDefaultsForEnvironment(obj)

	lroPrefix := fmt.Sprintf("projects/%s/locations/%s/environments", name.Project.ID, name.Location)
	lroRet := proto.Clone(obj).(*pb.Environment)
	lroRet.CreateTime = nil
	lroRet.UpdateTime = nil
	lroRet.Uuid = ""
	// State can only be set to running with UpdateEnvironment
	lroRet.State = pb.Environment_CREATING
	return s.operations.StartLRO(ctx, lroPrefix, nil, func() (proto.Message, error) {
		// Environment ID is parsed from environment.Name, projects/*/locations/*/environments/*
		tokens := strings.Split(lroRet.Name, "/")
		environmentId := tokens[5]
		// The resource name of the environment.
		obj.Name = fmt.Sprintf("projects/%s/locations/%s/environments/%s", name.Project.ID, name.Location, environmentId)
		fqn := obj.Name
		obj.Uuid = "123e4567-e89b-12d3-a456-426614174000" // this doesn't seem to be configurable

		// Create needs to set state
		obj.State = pb.Environment_CREATING

		if err := s.storage.Create(ctx, fqn, obj); err != nil {
			return nil, err
		}
		return obj, nil
	})
}

func (s *ComposerV1) UpdateEnvironment(ctx context.Context, req *pb.UpdateEnvironmentRequest) (*longrunningpb.Operation, error) {
	name, err := s.parseEnvironmentName(req.Name)
	if err != nil {
		return nil, err
	}

	fqn := name.String()

	obj := &pb.Environment{}
	if err := s.storage.Get(ctx, fqn, obj); err != nil {
		return nil, err
	}

	// Required. The update mask applies to the resource.
	paths := req.GetUpdateMask().GetPaths()
	if len(paths) == 0 {
		return nil, status.Errorf(codes.InvalidArgument, "update_mask must be provided")
	}

	// TODO: support update mask

	// TODO:  only labels are updateable with a mask
	proto.Merge(obj, req.GetEnvironment())

	s.populateDefaultsForEnvironment(obj)

	obj.State = pb.Environment_UPDATING
	obj.UpdateTime = timestamppb.New(time.Now())

	if err := s.storage.Update(ctx, fqn, obj); err != nil {
		return nil, err
	}
	updatedObj := proto.Clone(obj).(*pb.Environment)
	updatedObj.CreateTime = nil
	updatedObj.UpdateTime = nil
	updatedObj.Uuid = ""
	prefix := fmt.Sprintf("projects/%s/locations/%s/environments", name.Project.ID, name.Location)
	return s.operations.StartLRO(ctx, prefix, nil, func() (proto.Message, error) {
		// State can only be set to running with UpdateEnvironment
		updatedObj.State = pb.Environment_RUNNING
		obj.State = pb.Environment_RUNNING
		obj.UpdateTime = timestamppb.New(time.Now())
		if err := s.storage.Update(ctx, fqn, obj); err != nil {
			return nil, err
		}
		return updatedObj, nil
	})
}

func (s *ComposerV1) DeleteEnvironment(ctx context.Context, req *pb.DeleteEnvironmentRequest) (*longrunningpb.Operation, error) {
	name, err := s.parseEnvironmentName(req.Name)
	if err != nil {
		return nil, err
	}

	fqn := name.String()

	deletedObj := &pb.Environment{}
	if err := s.storage.Delete(ctx, fqn, deletedObj); err != nil {
		return nil, err
	}
	prefix := fmt.Sprintf("projects/%s/locations/%s/environments", name.Project.ID, name.Location)
	deletedObj.CreateTime = nil
	deletedObj.UpdateTime = nil
	deletedObj.Uuid = ""
	// State can only be set to running with UpdateEnvironment
	deletedObj.State = pb.Environment_DELETING
	return s.operations.StartLRO(ctx, prefix, nil, func() (proto.Message, error) {
		return &emptypb.Empty{}, nil
	})
}

func (s *ComposerV1) populateDefaultsForEnvironment(obj *pb.Environment) {
	if obj.Config == nil {
		obj.Config = &pb.EnvironmentConfig{}
	}

	s.populateDefaultsForEnvironmentConfig(obj.Config)
}

func (s *ComposerV1) populateDefaultsForEnvironmentConfig(config *pb.EnvironmentConfig) {
	if config.NodeCount == 0 {
		config.NodeCount = 3
	}

	if config.SoftwareConfig == nil {
		config.SoftwareConfig = &pb.SoftwareConfig{}
	}

	s.populateDefaultsForSoftwareConfig(config.SoftwareConfig)

	if config.PrivateEnvironmentConfig == nil {
		config.PrivateEnvironmentConfig = &pb.PrivateEnvironmentConfig{}
	}
	// TODO: more populating

}

func (s *ComposerV1) populateDefaultsForSoftwareConfig(config *pb.SoftwareConfig) {
	// TODO:
}

type environmentName struct {
	Project         *projects.ProjectData
	Location        string
	EnvironmentName string
}

func (n *environmentName) String() string {
	return "projects/" + n.Project.ID + "/locations/" + n.Location + "/environments/" + n.EnvironmentName
}

// parseEnvironmentName parses a string into a environmentName.
// The expected form is `projects/*/locations/*/environments/*`.
func (s *MockService) parseEnvironmentName(name string) (*environmentName, error) {
	tokens := strings.Split(name, "/")

	if len(tokens) == 6 && tokens[0] == "projects" && tokens[2] == "locations" && tokens[4] == "environments" {
		project, err := s.Projects.GetProjectByID(tokens[1])
		if err != nil {
			return nil, err
		}

		name := &environmentName{
			Project:         project,
			Location:        tokens[3],
			EnvironmentName: tokens[5],
		}

		return name, nil
	}

	return nil, status.Errorf(codes.InvalidArgument, "name %q is not valid", name)
}

// parseParentEnvironment parses a string into project and location.
// The expected form is `projects/*/locations/*`.
func (s *MockService) parseParentEnvironment(name string) (*environmentName, error) {
	tokens := strings.Split(name, "/")

	if len(tokens) == 4 && tokens[0] == "projects" && tokens[2] == "locations" {
		project, err := s.Projects.GetProjectByID(tokens[1])
		if err != nil {
			return nil, err
		}

		name := &environmentName{
			Project:  project,
			Location: tokens[3],
		}

		return name, nil
	}

	return nil, status.Errorf(codes.InvalidArgument, "name %q is not valid", name)
}
