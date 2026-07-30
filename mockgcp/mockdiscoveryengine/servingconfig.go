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
	"google.golang.org/protobuf/types/known/timestamppb"

	pbbeta "cloud.google.com/go/discoveryengine/apiv1beta/discoveryenginepb"
	"github.com/GoogleCloudPlatform/k8s-config-connector/mockgcp/common/projects"
)

type servingConfigService struct {
	*MockService
	pbbeta.UnimplementedServingConfigServiceServer
}

func (s *servingConfigService) GetServingConfig(ctx context.Context, req *pbbeta.GetServingConfigRequest) (*pbbeta.ServingConfig, error) {
	name, err := s.parseServingConfigName(req.GetName())
	if err != nil {
		return nil, err
	}

	fqn := name.String()
	obj := &pbbeta.ServingConfig{}
	if err := s.storage.Get(ctx, fqn, obj); err != nil {
		if status.Code(err) == codes.NotFound {
			// If not found in storage, we dynamically return a default pre-created config
			// (as it is pre-created by GCP when the datastore/engine is created)
			now := timestamppb.New(time.Now())
			obj = &pbbeta.ServingConfig{
				Name:         fqn,
				DisplayName:  "default_search",
				SolutionType: pbbeta.SolutionType_SOLUTION_TYPE_SEARCH,
				CreateTime:   now,
				UpdateTime:   now,
			}
			if err := s.storage.Create(ctx, fqn, obj); err != nil {
				return nil, err
			}
			return obj, nil
		}
		return nil, err
	}
	return obj, nil
}

func (s *servingConfigService) UpdateServingConfig(ctx context.Context, req *pbbeta.UpdateServingConfigRequest) (*pbbeta.ServingConfig, error) {
	name, err := s.parseServingConfigName(req.GetServingConfig().GetName())
	if err != nil {
		return nil, err
	}
	fqn := name.String()
	obj := &pbbeta.ServingConfig{}

	if err := s.storage.Get(ctx, fqn, obj); err != nil {
		if status.Code(err) == codes.NotFound {
			// Pre-create if it doesn't exist yet
			now := timestamppb.New(time.Now())
			obj = &pbbeta.ServingConfig{
				Name:         fqn,
				DisplayName:  "default_search",
				SolutionType: pbbeta.SolutionType_SOLUTION_TYPE_SEARCH,
				CreateTime:   now,
				UpdateTime:   now,
			}
			if err := s.storage.Create(ctx, fqn, obj); err != nil {
				return nil, err
			}
		} else {
			return nil, err
		}
	}

	// Simple merge of updated fields
	proto.Merge(obj, req.GetServingConfig())
	obj.UpdateTime = timestamppb.New(time.Now())

	if err := s.storage.Update(ctx, fqn, obj); err != nil {
		return nil, err
	}
	return obj, nil
}

type servingConfigName struct {
	Project       *projects.ProjectData
	Location      string
	Collection    string
	DataStore     string
	Engine        string
	ServingConfig string
}

func (n *servingConfigName) String() string {
	if n.Engine != "" {
		return fmt.Sprintf("projects/%d/locations/%s/collections/%s/engines/%s/servingConfigs/%s", n.Project.Number, n.Location, n.Collection, n.Engine, n.ServingConfig)
	}
	return fmt.Sprintf("projects/%d/locations/%s/collections/%s/dataStores/%s/servingConfigs/%s", n.Project.Number, n.Location, n.Collection, n.DataStore, n.ServingConfig)
}

func (s *MockService) parseServingConfigName(name string) (*servingConfigName, error) {
	name = strings.TrimPrefix(name, "//discoveryengine.googleapis.com/")
	name = strings.TrimPrefix(name, "/")
	tokens := strings.Split(name, "/")
	if len(tokens) == 10 && tokens[0] == "projects" && tokens[2] == "locations" && tokens[4] == "collections" && tokens[8] == "servingConfigs" {
		project, err := s.Projects.GetProjectByID(tokens[1])
		if err != nil {
			return nil, err
		}
		if tokens[6] == "engines" {
			return &servingConfigName{
				Project:       project,
				Location:      tokens[3],
				Collection:    tokens[5],
				Engine:        tokens[7],
				ServingConfig: tokens[9],
			}, nil
		}
		if tokens[6] == "dataStores" {
			return &servingConfigName{
				Project:       project,
				Location:      tokens[3],
				Collection:    tokens[5],
				DataStore:     tokens[7],
				ServingConfig: tokens[9],
			}, nil
		}
	}
	return nil, status.Errorf(codes.InvalidArgument, "invalid serving config name: %q", name)
}
