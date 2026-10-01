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
// proto.service: google.cloud.discoveryengine.v1.ControlService
// proto.message: google.cloud.discoveryengine.v1.Control

package mockdiscoveryengine

import (
	"context"
	"fmt"
	"strings"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/types/known/emptypb"

	pb "cloud.google.com/go/discoveryengine/apiv1/discoveryenginepb"

	"github.com/GoogleCloudPlatform/k8s-config-connector/mockgcp/common/fields"
	"github.com/GoogleCloudPlatform/k8s-config-connector/mockgcp/common/projects"
	"github.com/GoogleCloudPlatform/k8s-config-connector/mockgcp/pkg/storage"
)

type controlService struct {
	*MockService
	pb.UnimplementedControlServiceServer
}

func (s *controlService) CreateControl(ctx context.Context, req *pb.CreateControlRequest) (*pb.Control, error) {
	reqName := fmt.Sprintf("%s/controls/%s", req.GetParent(), req.GetControlId())
	name, err := s.parseControlName(reqName)
	if err != nil {
		return nil, err
	}

	fqn := name.String()
	obj := proto.Clone(req.GetControl()).(*pb.Control)
	obj.Name = s.buildControlResponseName(ctx, name)

	if err := s.storage.Create(ctx, fqn, obj); err != nil {
		return nil, err
	}

	return obj, nil
}

func (s *controlService) GetControl(ctx context.Context, req *pb.GetControlRequest) (*pb.Control, error) {
	name, err := s.parseControlName(req.GetName())
	if err != nil {
		return nil, err
	}

	fqn := name.String()
	obj := &pb.Control{}
	if err := s.storage.Get(ctx, fqn, obj); err != nil {
		if status.Code(err) == codes.NotFound {
			return nil, status.Errorf(codes.NotFound, "Control with name %q does not exist.", fqn)
		}
		return nil, err
	}
	obj.Name = s.buildControlResponseName(ctx, name)
	return obj, nil
}

func (s *controlService) UpdateControl(ctx context.Context, req *pb.UpdateControlRequest) (*pb.Control, error) {
	name, err := s.parseControlName(req.GetControl().GetName())
	if err != nil {
		return nil, err
	}

	fqn := name.String()
	obj := &pb.Control{}

	if err := s.storage.Get(ctx, fqn, obj); err != nil {
		if status.Code(err) == codes.NotFound {
			return nil, status.Errorf(codes.NotFound, "Control with name %q does not exist.", fqn)
		}
		return nil, err
	}

	paths := req.GetUpdateMask().GetPaths()
	if len(paths) > 0 {
		if err := fields.UpdateByFieldMask(obj, req.GetControl(), paths); err != nil {
			return nil, err
		}
	} else {
		proto.Merge(obj, req.GetControl())
	}

	obj.Name = s.buildControlResponseName(ctx, name)

	if err := s.storage.Update(ctx, fqn, obj); err != nil {
		return nil, err
	}
	return obj, nil
}

func (s *controlService) DeleteControl(ctx context.Context, req *pb.DeleteControlRequest) (*emptypb.Empty, error) {
	name, err := s.parseControlName(req.GetName())
	if err != nil {
		return nil, err
	}

	fqn := name.String()

	deleted := &pb.Control{}
	if err := s.storage.Delete(ctx, fqn, deleted); err != nil {
		return nil, err
	}

	return &emptypb.Empty{}, nil
}

type controlName struct {
	Project    *projects.ProjectData
	Location   string
	Collection string
	DataStore  string
	Engine     string
	Control    string
}

func (n *controlName) String() string {
	if n.DataStore != "" {
		return fmt.Sprintf("projects/%d/locations/%s/collections/%s/dataStores/%s/controls/%s", n.Project.Number, n.Location, n.Collection, n.DataStore, n.Control)
	}
	return fmt.Sprintf("projects/%d/locations/%s/collections/%s/engines/%s/controls/%s", n.Project.Number, n.Location, n.Collection, n.Engine, n.Control)
}

func (s *MockService) buildControlResponseName(ctx context.Context, name *controlName) string {
	if name.Engine != "" {
		return fmt.Sprintf("projects/%d/locations/%s/collections/%s/engines/%s/controls/%s", name.Project.Number, name.Location, name.Collection, name.Engine, name.Control)
	}

	// If DataStore is specified, check if an Engine is attached to this DataStore
	if name.DataStore != "" {
		enginePrefix := fmt.Sprintf("projects/%d/locations/%s/collections/%s/engines/", name.Project.Number, name.Location, name.Collection)
		var foundEngineName string
		_ = s.storage.List(ctx, (&pb.Engine{}).ProtoReflect().Descriptor(), storage.ListOptions{Prefix: enginePrefix}, func(obj proto.Message) error {
			if engine, ok := obj.(*pb.Engine); ok {
				for _, dsID := range engine.GetDataStoreIds() {
					if dsID == name.DataStore {
						foundEngineName = engine.GetName()
						return nil
					}
				}
			}
			return nil
		})
		if foundEngineName != "" {
			return fmt.Sprintf("%s/controls/%s", foundEngineName, name.Control)
		}
	}

	return name.String()
}

func (s *MockService) parseControlName(name string) (*controlName, error) {
	tokens := strings.Split(name, "/")
	if len(tokens) == 10 && tokens[0] == "projects" && tokens[2] == "locations" && tokens[4] == "collections" && tokens[8] == "controls" {
		project, err := s.Projects.GetProjectByIDOrNumber(tokens[1])
		if err != nil {
			return nil, err
		}
		if tokens[6] == "dataStores" {
			return &controlName{
				Project:    project,
				Location:   tokens[3],
				Collection: tokens[5],
				DataStore:  tokens[7],
				Control:    tokens[9],
			}, nil
		}
		if tokens[6] == "engines" {
			return &controlName{
				Project:    project,
				Location:   tokens[3],
				Collection: tokens[5],
				Engine:     tokens[7],
				Control:    tokens[9],
			}, nil
		}
	}
	if len(tokens) == 8 && tokens[0] == "projects" && tokens[2] == "locations" && tokens[6] == "controls" {
		project, err := s.Projects.GetProjectByIDOrNumber(tokens[1])
		if err != nil {
			return nil, err
		}
		if tokens[4] == "dataStores" {
			return &controlName{
				Project:    project,
				Location:   tokens[3],
				Collection: "default_collection",
				DataStore:  tokens[5],
				Control:    tokens[7],
			}, nil
		}
		if tokens[4] == "engines" {
			return &controlName{
				Project:    project,
				Location:   tokens[3],
				Collection: "default_collection",
				Engine:     tokens[5],
				Control:    tokens[7],
			}, nil
		}
	}
	return nil, status.Errorf(codes.InvalidArgument, "invalid control name: %q", name)
}
