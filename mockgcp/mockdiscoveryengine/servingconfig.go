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
// proto.service: google.cloud.discoveryengine.v1beta.ServingConfigService
// proto.message: google.cloud.discoveryengine.v1beta.ServingConfig

package mockdiscoveryengine

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

	pb_v1beta "cloud.google.com/go/discoveryengine/apiv1beta/discoveryenginepb"

	"github.com/GoogleCloudPlatform/k8s-config-connector/mockgcp/common/projects"
)

type servingConfigService struct {
	*MockService
	pb_v1beta.UnimplementedServingConfigServiceServer
}

func (s *servingConfigService) CreateServingConfig(ctx context.Context, req *pb_v1beta.CreateServingConfigRequest) (*pb_v1beta.ServingConfig, error) {
	reqName := fmt.Sprintf("%s/servingConfigs/%s", req.GetParent(), req.GetServingConfigId())
	name, err := s.parseServingConfigName(reqName)
	if err != nil {
		return nil, err
	}
	now := time.Now()

	fqn := name.String()
	obj := proto.Clone(req.GetServingConfig()).(*pb_v1beta.ServingConfig)
	obj.Name = fqn
	obj.CreateTime = timestamppb.New(now)
	obj.UpdateTime = timestamppb.New(now)

	if err := s.storage.Create(ctx, fqn, obj); err != nil {
		return nil, err
	}

	return obj, nil
}

func (s *servingConfigService) GetServingConfig(ctx context.Context, req *pb_v1beta.GetServingConfigRequest) (*pb_v1beta.ServingConfig, error) {
	name, err := s.parseServingConfigName(req.GetName())
	if err != nil {
		return nil, err
	}

	fqn := name.String()
	obj := &pb_v1beta.ServingConfig{}
	if err := s.storage.Get(ctx, fqn, obj); err != nil {
		if status.Code(err) == codes.NotFound {
			return nil, status.Errorf(codes.NotFound, "ServingConfig with name %q does not exist.", fqn)
		}
		return nil, err
	}
	return obj, nil
}

func (s *servingConfigService) UpdateServingConfig(ctx context.Context, req *pb_v1beta.UpdateServingConfigRequest) (*pb_v1beta.ServingConfig, error) {
	name, err := s.parseServingConfigName(req.GetServingConfig().GetName())
	if err != nil {
		return nil, err
	}
	fqn := name.String()
	obj := &pb_v1beta.ServingConfig{}

	if err := s.storage.Get(ctx, fqn, obj); err != nil {
		if status.Code(err) == codes.NotFound {
			return nil, status.Errorf(codes.NotFound, "ServingConfig %q not found", name)
		}
		return nil, err
	}

	// simple merge for now
	proto.Merge(obj, req.GetServingConfig())
	obj.Name = fqn
	obj.UpdateTime = timestamppb.New(time.Now())

	if err := s.storage.Update(ctx, fqn, obj); err != nil {
		return nil, err
	}
	return obj, nil
}

func (s *servingConfigService) DeleteServingConfig(ctx context.Context, req *pb_v1beta.DeleteServingConfigRequest) (*emptypb.Empty, error) {
	name, err := s.parseServingConfigName(req.GetName())
	if err != nil {
		return nil, err
	}

	fqn := name.String()

	deleted := &pb_v1beta.ServingConfig{}
	if err := s.storage.Delete(ctx, fqn, deleted); err != nil {
		return nil, err
	}

	return &emptypb.Empty{}, nil
}

type servingConfigName struct {
	Project       *projects.ProjectData
	Location      string
	Collection    string
	Engine        string
	ServingConfig string
}

func (n *servingConfigName) String() string {
	return fmt.Sprintf("projects/%d/locations/%s/collections/%s/engines/%s/servingConfigs/%s", n.Project.Number, n.Location, n.Collection, n.Engine, n.ServingConfig)
}

func (s *MockService) parseServingConfigName(name string) (*servingConfigName, error) {
	tokens := strings.Split(name, "/")
	if len(tokens) == 10 && tokens[0] == "projects" && tokens[2] == "locations" && tokens[4] == "collections" && tokens[6] == "engines" && tokens[8] == "servingConfigs" {
		project, err := s.Projects.GetProjectByIDOrNumber(tokens[1])
		if err != nil {
			return nil, err
		}
		return &servingConfigName{
			Project:       project,
			Location:      tokens[3],
			Collection:    tokens[5],
			Engine:        tokens[7],
			ServingConfig: tokens[9],
		}, nil
	}
	return nil, status.Errorf(codes.InvalidArgument, "invalid serving config name %q", name)
}
