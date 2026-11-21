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

package mockdialogflow

import (
	"context"
	"fmt"
	"strings"

	"github.com/google/uuid"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/types/known/emptypb"

	pb "cloud.google.com/go/dialogflow/apiv2beta1/dialogflowpb"
	"github.com/GoogleCloudPlatform/k8s-config-connector/mockgcp/pkg/storage"
)

type entityTypesServer struct {
	*MockService
	pb.UnimplementedEntityTypesServer
}

type entityTypeName struct {
	Project    string
	Location   string
	EntityType string
}

func (n *entityTypeName) String() string {
	if n.Location != "" {
		return fmt.Sprintf("projects/%s/locations/%s/agent/entityTypes/%s", n.Project, n.Location, n.EntityType)
	}
	return fmt.Sprintf("projects/%s/agent/entityTypes/%s", n.Project, n.EntityType)
}

func (n *entityTypeName) Parent() string {
	if n.Location != "" {
		return fmt.Sprintf("projects/%s/locations/%s/agent", n.Project, n.Location)
	}
	return fmt.Sprintf("projects/%s/agent", n.Project)
}

func (s *MockService) parseEntityTypeName(name string) (*entityTypeName, error) {
	if name == "" {
		return nil, status.Errorf(codes.InvalidArgument, "name must be provided")
	}

	tokens := strings.Split(name, "/")
	if len(tokens) == 7 && tokens[0] == "projects" && tokens[2] == "locations" && tokens[4] == "agent" && tokens[5] == "entityTypes" {
		project, err := s.Projects.GetProjectByIDOrNumber(tokens[1])
		if err != nil {
			return nil, err
		}

		return &entityTypeName{
			Project:    project.ID,
			Location:   tokens[3],
			EntityType: tokens[6],
		}, nil
	}
	if len(tokens) == 5 && tokens[0] == "projects" && tokens[2] == "agent" && tokens[3] == "entityTypes" {
		project, err := s.Projects.GetProjectByIDOrNumber(tokens[1])
		if err != nil {
			return nil, err
		}

		return &entityTypeName{
			Project:    project.ID,
			EntityType: tokens[4],
		}, nil
	}

	return nil, status.Errorf(codes.InvalidArgument, "name %q is not valid", name)
}

func (s *MockService) parseEntityTypeParent(parent string) (*entityTypeName, error) {
	if parent == "" {
		return nil, status.Errorf(codes.InvalidArgument, "parent must be provided")
	}

	tokens := strings.Split(parent, "/")
	if len(tokens) == 5 && tokens[0] == "projects" && tokens[2] == "locations" && tokens[4] == "agent" {
		project, err := s.Projects.GetProjectByIDOrNumber(tokens[1])
		if err != nil {
			return nil, err
		}
		return &entityTypeName{
			Project:  project.ID,
			Location: tokens[3],
		}, nil
	}
	if len(tokens) == 3 && tokens[0] == "projects" && tokens[2] == "agent" {
		project, err := s.Projects.GetProjectByIDOrNumber(tokens[1])
		if err != nil {
			return nil, err
		}
		return &entityTypeName{
			Project: project.ID,
		}, nil
	}
	return nil, status.Errorf(codes.InvalidArgument, "parent %q is not valid", parent)
}

func (s *entityTypesServer) GetEntityType(ctx context.Context, req *pb.GetEntityTypeRequest) (*pb.EntityType, error) {
	name, err := s.parseEntityTypeName(req.GetName())
	if err != nil {
		return nil, err
	}

	fqn := name.String()
	obj := &pb.EntityType{}
	if err := s.storage.Get(ctx, fqn, obj); err != nil {
		if status.Code(err) == codes.NotFound {
			return nil, status.Errorf(codes.NotFound, "EntityType %q not found", fqn)
		}
		return nil, err
	}

	return obj, nil
}

func (s *entityTypesServer) CreateEntityType(ctx context.Context, req *pb.CreateEntityTypeRequest) (*pb.EntityType, error) {
	parent, err := s.parseEntityTypeParent(req.GetParent())
	if err != nil {
		return nil, err
	}

	entityType := req.GetEntityType()
	if entityType == nil {
		return nil, status.Errorf(codes.InvalidArgument, "entity_type is required")
	}

	parentPrefix := parent.Parent() + "/entityTypes/"
	kind := (&pb.EntityType{}).ProtoReflect().Descriptor()
	if err := s.storage.List(ctx, kind, storage.ListOptions{Prefix: parentPrefix}, func(obj proto.Message) error {
		existing := obj.(*pb.EntityType)
		if existing.GetDisplayName() == entityType.GetDisplayName() {
			return status.Errorf(codes.FailedPrecondition, "com.google.apps.framework.request.FailedPreconditionException: This entity's displayName '%s' already exists with current agent. Code: FAILED_PRECONDITION", entityType.GetDisplayName())
		}
		return nil
	}); err != nil {
		return nil, err
	}

	id := uuid.New().String()
	name := parentPrefix + id

	obj := proto.Clone(entityType).(*pb.EntityType)
	obj.Name = name

	if err := s.storage.Create(ctx, name, obj); err != nil {
		return nil, err
	}

	return obj, nil
}

func (s *entityTypesServer) UpdateEntityType(ctx context.Context, req *pb.UpdateEntityTypeRequest) (*pb.EntityType, error) {
	entityType := req.GetEntityType()
	if entityType == nil {
		return nil, status.Errorf(codes.InvalidArgument, "entity_type is required")
	}

	name, err := s.parseEntityTypeName(entityType.GetName())
	if err != nil {
		return nil, err
	}

	fqn := name.String()
	existing := &pb.EntityType{}
	if err := s.storage.Get(ctx, fqn, existing); err != nil {
		if status.Code(err) == codes.NotFound {
			return nil, status.Errorf(codes.NotFound, "EntityType %q not found", fqn)
		}
		return nil, err
	}

	obj := proto.Clone(existing).(*pb.EntityType)
	if entityType.GetDisplayName() != "" {
		obj.DisplayName = entityType.GetDisplayName()
	}
	if entityType.GetKind() != pb.EntityType_KIND_UNSPECIFIED {
		obj.Kind = entityType.GetKind()
	}
	if len(entityType.GetEntities()) > 0 {
		obj.Entities = entityType.GetEntities()
	}
	obj.EnableFuzzyExtraction = entityType.GetEnableFuzzyExtraction()

	if err := s.storage.Update(ctx, fqn, obj); err != nil {
		return nil, err
	}

	return obj, nil
}

func (s *entityTypesServer) DeleteEntityType(ctx context.Context, req *pb.DeleteEntityTypeRequest) (*emptypb.Empty, error) {
	name, err := s.parseEntityTypeName(req.GetName())
	if err != nil {
		return nil, err
	}

	fqn := name.String()
	deletedObj := &pb.EntityType{}
	if err := s.storage.Delete(ctx, fqn, deletedObj); err != nil {
		if status.Code(err) == codes.NotFound {
			return nil, status.Errorf(codes.NotFound, "EntityType %q not found", fqn)
		}
		return nil, err
	}

	return &emptypb.Empty{}, nil
}
