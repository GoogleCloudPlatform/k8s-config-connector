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

package mockdiscoveryengine

import (
	"context"
	"fmt"
	"strings"

	"google.golang.org/genproto/googleapis/longrunning"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/types/known/timestamppb"

	pb "cloud.google.com/go/discoveryengine/apiv1/discoveryenginepb"
	"github.com/GoogleCloudPlatform/k8s-config-connector/mockgcp/common/projects"
)

func (s *engineService) GetEngine(ctx context.Context, req *pb.GetEngineRequest) (*pb.Engine, error) {
	name, err := s.parseEngineName(req.Name)
	if err != nil {
		return nil, err
	}

	fqn := name.String()

	obj := &pb.Engine{}
	if err := s.storage.Get(ctx, fqn, obj); err != nil {
		if status.Code(err) == codes.NotFound {
			return nil, status.Errorf(codes.NotFound, "Engine %q not found.", fqn)
		}
		return nil, err
	}

	return obj, nil
}

func (s *engineService) CreateEngine(ctx context.Context, req *pb.CreateEngineRequest) (*longrunning.Operation, error) {
	parent, err := s.parseCollectionName(req.Parent)
	if err != nil {
		return nil, err
	}

	id := req.EngineId
	fqn := parent.String() + "/engines/" + id

	obj := proto.Clone(req.Engine).(*pb.Engine)
	obj.Name = fqn
	obj.CreateTime = timestamppb.Now()
	obj.UpdateTime = timestamppb.Now()

	if err := s.storage.Create(ctx, fqn, obj); err != nil {
		return nil, err
	}

	prefix := fmt.Sprintf("projects/%s/locations/%s/collections/%s", parent.Project.ID, parent.Location, parent.Collection)
	return s.operations.DoneLRO(ctx, prefix, nil, obj)
}

func (s *engineService) UpdateEngine(ctx context.Context, req *pb.UpdateEngineRequest) (*pb.Engine, error) {
	name, err := s.parseEngineName(req.Engine.Name)
	if err != nil {
		return nil, err
	}
	fqn := name.String()
	obj := &pb.Engine{}

	if err := s.storage.Get(ctx, fqn, obj); err != nil {
		if status.Code(err) == codes.NotFound {
			return nil, status.Errorf(codes.NotFound, "engine %q not found", fqn)
		}
		return nil, err
	}

	// TODO: support update mask

	proto.Merge(obj, req.GetEngine())

	if err := s.storage.Update(ctx, fqn, obj); err != nil {
		return nil, err
	}
	return obj, nil
}

func (s *engineService) DeleteEngine(ctx context.Context, req *pb.DeleteEngineRequest) (*longrunning.Operation, error) {
	name, err := s.parseEngineName(req.Name)
	if err != nil {
		return nil, err
	}

	fqn := name.String()

	obj := &pb.Engine{}
	if err := s.storage.Delete(ctx, fqn, obj); err != nil {
		if status.Code(err) == codes.NotFound {
			return nil, status.Errorf(codes.NotFound, "Engine %q not found.", fqn)
		}
		return nil, err
	}

	prefix := fmt.Sprintf("projects/%s/locations/%s/collections/%s", name.Project.ID, name.Location, name.Collection)
	return s.operations.DoneLRO(ctx, prefix, nil, nil)
}

type engineName struct {
	Project    *projects.ProjectData
	Location   string
	Collection string
	Engine     string
}

func (n *engineName) String() string {
	return fmt.Sprintf("projects/%s/locations/%s/collections/%s/engines/%s", n.Project.ID, n.Location, n.Collection, n.Engine)
}

func (s *MockService) parseEngineName(name string) (*engineName, error) {
	tokens := strings.Split(name, "/")
	if len(tokens) == 8 && tokens[0] == "projects" && tokens[2] == "locations" && tokens[4] == "collections" && tokens[6] == "engines" {
		project, err := s.Projects.GetProjectByID(tokens[1])
		if err != nil {
			return nil, err
		}

		return &engineName{
			Project:    project,
			Location:   tokens[3],
			Collection: tokens[5],
			Engine:     tokens[7],
		}, nil
	}
	return nil, status.Errorf(codes.InvalidArgument, "invalid engine name %q", name)
}

type collectionName struct {
	Project    *projects.ProjectData
	Location   string
	Collection string
}

func (n *collectionName) String() string {
	return fmt.Sprintf("projects/%s/locations/%s/collections/%s", n.Project.ID, n.Location, n.Collection)
}

func (s *MockService) parseCollectionName(name string) (*collectionName, error) {
	tokens := strings.Split(name, "/")
	if len(tokens) == 6 && tokens[0] == "projects" && tokens[2] == "locations" && tokens[4] == "collections" {
		project, err := s.Projects.GetProjectByID(tokens[1])
		if err != nil {
			return nil, err
		}

		return &collectionName{
			Project:    project,
			Location:   tokens[3],
			Collection: tokens[5],
		}, nil
	}
	return nil, status.Errorf(codes.InvalidArgument, "invalid collection name %q", name)
}
