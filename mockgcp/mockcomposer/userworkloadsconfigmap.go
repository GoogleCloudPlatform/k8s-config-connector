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
// proto.service: google.cloud.orchestration.airflow.service.v1.Environments
// proto.message: google.cloud.orchestration.airflow.service.v1.UserWorkloadsConfigMap

package mockcomposer

import (
	"context"
	"fmt"
	"strings"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/types/known/emptypb"

	"github.com/GoogleCloudPlatform/k8s-config-connector/mockgcp/common/projects"
	pb "github.com/GoogleCloudPlatform/k8s-config-connector/mockgcp/generated/mockgcp/cloud/orchestration/airflow/service/v1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/mockgcp/pkg/storage"
)

type userWorkloadsConfigMapName struct {
	Project                *projects.ProjectData
	Location               string
	Environment            string
	UserWorkloadsConfigMap string
}

func (n *userWorkloadsConfigMapName) String() string {
	return fmt.Sprintf("projects/%s/locations/%s/environments/%s/userWorkloadsConfigMaps/%s", n.Project.ID, n.Location, n.Environment, n.UserWorkloadsConfigMap)
}

func (s *MockService) parseUserWorkloadsConfigMapName(name string) (*userWorkloadsConfigMapName, error) {
	tokens := strings.Split(name, "/")
	if len(tokens) == 8 && tokens[0] == "projects" && tokens[2] == "locations" && tokens[4] == "environments" && tokens[6] == "userWorkloadsConfigMaps" {
		project, err := s.Projects.GetProjectByID(tokens[1])
		if err != nil {
			return nil, err
		}
		return &userWorkloadsConfigMapName{
			Project:                project,
			Location:               tokens[3],
			Environment:            tokens[5],
			UserWorkloadsConfigMap: tokens[7],
		}, nil
	}
	return nil, status.Errorf(codes.InvalidArgument, "name %q is not valid", name)
}

func (s *ComposerV1) GetUserWorkloadsConfigMap(ctx context.Context, req *pb.GetUserWorkloadsConfigMapRequest) (*pb.UserWorkloadsConfigMap, error) {
	name, err := s.parseUserWorkloadsConfigMapName(req.Name)
	if err != nil {
		return nil, err
	}

	fqn := name.String()

	obj := &pb.UserWorkloadsConfigMap{}
	if err := s.storage.Get(ctx, fqn, obj); err != nil {
		if status.Code(err) == codes.NotFound {
			return nil, status.Errorf(codes.NotFound, "configmaps %q not found", name.UserWorkloadsConfigMap)
		}
		return nil, err
	}

	return proto.Clone(obj).(*pb.UserWorkloadsConfigMap), nil
}

func (s *ComposerV1) CreateUserWorkloadsConfigMap(ctx context.Context, req *pb.CreateUserWorkloadsConfigMapRequest) (*pb.UserWorkloadsConfigMap, error) {
	configMapName := req.GetUserWorkloadsConfigMap().GetName()
	if configMapName == "" {
		return nil, status.Errorf(codes.InvalidArgument, "config map name is required")
	}
	if !strings.HasPrefix(configMapName, "projects/") {
		configMapName = fmt.Sprintf("%s/userWorkloadsConfigMaps/%s", req.GetParent(), configMapName)
	}

	name, err := s.parseUserWorkloadsConfigMapName(configMapName)
	if err != nil {
		return nil, err
	}

	fqn := name.String()

	obj := proto.Clone(req.GetUserWorkloadsConfigMap()).(*pb.UserWorkloadsConfigMap)
	obj.Name = fqn

	if err := s.storage.Create(ctx, fqn, obj); err != nil {
		return nil, err
	}

	return proto.Clone(obj).(*pb.UserWorkloadsConfigMap), nil
}

func (s *ComposerV1) UpdateUserWorkloadsConfigMap(ctx context.Context, req *pb.UpdateUserWorkloadsConfigMapRequest) (*pb.UserWorkloadsConfigMap, error) {
	configMapName := req.GetUserWorkloadsConfigMap().GetName()
	name, err := s.parseUserWorkloadsConfigMapName(configMapName)
	if err != nil {
		return nil, err
	}

	fqn := name.String()

	obj := &pb.UserWorkloadsConfigMap{}
	if err := s.storage.Get(ctx, fqn, obj); err != nil {
		if status.Code(err) == codes.NotFound {
			return nil, status.Errorf(codes.NotFound, "configmaps %q not found", name.UserWorkloadsConfigMap)
		}
		return nil, err
	}

	obj.Data = req.GetUserWorkloadsConfigMap().GetData()

	if err := s.storage.Update(ctx, fqn, obj); err != nil {
		return nil, err
	}

	return proto.Clone(obj).(*pb.UserWorkloadsConfigMap), nil
}

func (s *ComposerV1) DeleteUserWorkloadsConfigMap(ctx context.Context, req *pb.DeleteUserWorkloadsConfigMapRequest) (*emptypb.Empty, error) {
	name, err := s.parseUserWorkloadsConfigMapName(req.Name)
	if err != nil {
		return nil, err
	}

	fqn := name.String()

	obj := &pb.UserWorkloadsConfigMap{}
	if err := s.storage.Delete(ctx, fqn, obj); err != nil {
		if status.Code(err) == codes.NotFound {
			return nil, status.Errorf(codes.NotFound, "configmaps %q not found", name.UserWorkloadsConfigMap)
		}
		return nil, err
	}

	return &emptypb.Empty{}, nil
}

func (s *ComposerV1) ListUserWorkloadsConfigMaps(ctx context.Context, req *pb.ListUserWorkloadsConfigMapsRequest) (*pb.ListUserWorkloadsConfigMapsResponse, error) {
	findPrefix := req.GetParent() + "/userWorkloadsConfigMaps/"

	response := &pb.ListUserWorkloadsConfigMapsResponse{}
	findKind := (&pb.UserWorkloadsConfigMap{}).ProtoReflect().Descriptor()

	if err := s.storage.List(ctx, findKind, storage.ListOptions{Prefix: findPrefix}, func(obj proto.Message) error {
		configMap := proto.Clone(obj).(*pb.UserWorkloadsConfigMap)
		response.UserWorkloadsConfigMaps = append(response.UserWorkloadsConfigMaps, configMap)
		return nil
	}); err != nil {
		return nil, err
	}

	return response, nil
}
