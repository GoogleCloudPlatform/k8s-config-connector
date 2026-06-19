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
// proto.service: google.cloud.discoveryengine.v1.EngineService
// proto.message: google.cloud.discoveryengine.v1.Engine

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

	pb "cloud.google.com/go/discoveryengine/apiv1/discoveryenginepb"
	longrunningpb "google.golang.org/genproto/googleapis/longrunning"

	"github.com/GoogleCloudPlatform/k8s-config-connector/mockgcp/common/projects"
)

type engineService struct {
	*MockService
	pb.UnimplementedEngineServiceServer
}

func (s *engineService) CreateEngine(ctx context.Context, req *pb.CreateEngineRequest) (*longrunningpb.Operation, error) {
	reqName := fmt.Sprintf("%s/engines/%s", req.GetParent(), req.GetEngineId())
	name, err := s.parseEngineName(reqName)
	if err != nil {
		return nil, err
	}
	now := time.Now()

	fqn := name.String()
	obj := proto.CloneOf(req.GetEngine())
	obj.Name = fqn
	obj.CreateTime = timestamppb.New(now)

	if err := s.storage.Create(ctx, fqn, obj); err != nil {
		return nil, err
	}

	prefix := fmt.Sprintf("projects/%d/locations/%s/collections/%s", name.Project.Number, name.Location, name.Collection)
	lroRet := proto.CloneOf(obj)
	lroRet.CreateTime = nil
	return s.operations.DoneLRO(ctx, prefix, nil, lroRet)
}

func (s *engineService) DeleteEngine(ctx context.Context, req *pb.DeleteEngineRequest) (*longrunningpb.Operation, error) {
	name, err := s.parseEngineName(req.Name)
	if err != nil {
		return nil, err
	}

	fqn := name.String()

	deleted := &pb.Engine{}
	if err := s.storage.Delete(ctx, fqn, deleted); err != nil {
		return nil, err
	}

	prefix := fmt.Sprintf("projects/%d/locations/%s/collections/%s", name.Project.Number, name.Location, name.Collection)
	return s.operations.DoneLRO(ctx, prefix, nil, nil)
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
			return nil, status.Errorf(codes.NotFound, "engine %q not found", name)
		}
		return nil, err
	}

	paths := req.GetUpdateMask().GetPaths()
	if len(paths) == 0 {
		return nil, status.Errorf(codes.InvalidArgument, "update_mask must be provided")
	}

	proto.Merge(obj, req.GetEngine())

	if err := s.storage.Update(ctx, fqn, obj); err != nil {
		return nil, err
	}
	return obj, nil
}

func (s *engineService) GetEngine(ctx context.Context, req *pb.GetEngineRequest) (*pb.Engine, error) {
	name, err := s.parseEngineName(req.Name)
	if err != nil {
		return nil, err
	}

	fqn := name.String()
	obj := &pb.Engine{}
	if err := s.storage.Get(ctx, fqn, obj); err != nil {
		if status.Code(err) == codes.NotFound {
			return nil, status.Errorf(codes.NotFound, "Engine %v not found.", name)
		}
		return nil, err
	}
	return obj, nil
}

type engineName struct {
	Project    *projects.ProjectData
	Location   string
	Collection string
	Engine     string
}

func (n *engineName) String() string {
	return fmt.Sprintf("projects/%d/locations/%s/collections/%s/engines/%s", n.Project.Number, n.Location, n.Collection, n.Engine)
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

	return nil, status.Errorf(codes.InvalidArgument, "invalid name %q", name)
}
