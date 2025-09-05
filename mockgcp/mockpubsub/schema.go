// Copyright 2024 Google LLC
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package mockpubsub

import (
	"context"
	"fmt"
	"strings"
	"time"

	"github.com/GoogleCloudPlatform/k8s-config-connector/mockgcp/common/projects"
	pb "github.com/GoogleCloudPlatform/k8s-config-connector/mockgcp/generated/mockgcp/pubsub/v1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/mockgcp/pkg/storage"
	"github.com/golang/protobuf/ptypes/empty"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/types/known/timestamppb"
)

type schemaService struct {
	*MockService
	pb.UnimplementedSchemaServiceServer
}

func (s *schemaService) CreateSchema(ctx context.Context, req *pb.CreateSchemaRequest) (*pb.Schema, error) {
	reqName := req.Parent + "/schemas/" + req.GetSchemaId()
	name, err := s.parseSchemaName(reqName)
	if err != nil {
		return nil, err
	}
	fqn := name.String()
	now := time.Now()
	obj := proto.Clone(req.GetSchema()).(*pb.Schema)
	obj.Name = name.String()
	obj.RevisionId = fmt.Sprintf("r%d", now.Unix())
	obj.RevisionCreateTime = timestamppb.New(now)
	s.populateDefaultsForSchema(name, obj)
	if err := s.schemas.Create(ctx, fqn, obj); err != nil {
		return nil, err
	}
	return obj, nil
}

func (s *schemaService) populateDefaultsForSchema(name *schemaName, obj *pb.Schema) {
	// TODO: populate any default values here
}

func (s *schemaService) GetSchema(ctx context.Context, req *pb.GetSchemaRequest) (*pb.Schema, error) {
	name, err := s.parseSchemaName(req.Name)
	if err != nil {
		return nil, err
	}
	fqn := name.String()
	obj := &pb.Schema{}
	if err := s.schemas.Get(ctx, fqn, obj); err != nil {
		if status.Code(err) == codes.NotFound {
			return nil, status.Errorf(codes.NotFound, "Resource not found (resource=%s).", name.String())
		}
		return nil, err
	}
	return obj, nil
}

func (s *schemaService) ListSchemas(ctx context.Context, req *pb.ListSchemasRequest) (*pb.ListSchemasResponse, error) {
	project, err := s.Projects.GetProjectByID(req.Parent)
	if err != nil {
		return nil, err
	}

	findPrefix := fmt.Sprintf("projects/%v/", project.ID)

	var schemas []*pb.Schema

	if err := s.schemas.List(ctx, storage.ListOptions{}, func(obj *pb.Schema) error {
		if strings.HasPrefix(obj.GetName(), findPrefix) {
			schemas = append(schemas, obj)
		}
		return nil
	}); err != nil {
		return nil, err
	}

	return &pb.ListSchemasResponse{
		Schemas: schemas,
	}, nil
}

func (s *schemaService) DeleteSchema(ctx context.Context, req *pb.DeleteSchemaRequest) (*empty.Empty, error) {
	name, err := s.parseSchemaName(req.Name)
	if err != nil {
		return nil, err
	}
	fqn := name.String()
	deletedObj := &pb.Schema{}
	if err := s.schemas.Delete(ctx, fqn, deletedObj); err != nil {
		return nil, err
	}
	return &empty.Empty{}, nil
}

type schemaName struct {
	Project *projects.ProjectData
	ID      string
}

func (n *schemaName) String() string {
	return "projects/" + n.Project.ID + "/schemas/" + n.ID
}

// parseSchemaName parses a string into a schemaName.
// The expected form is `projects/*/schemas/*`.
func (s *MockService) parseSchemaName(name string) (*schemaName, error) {
	tokens := strings.Split(name, "/")
	if len(tokens) == 4 && tokens[0] == "projects" && tokens[2] == "schemas" {
		project, err := s.Projects.GetProjectByID(tokens[1])
		if err != nil {
			return nil, err
		}
		name := &schemaName{
			Project: project,
			ID:      tokens[3],
		}
		return name, nil
	} else {
		return nil, status.Errorf(codes.InvalidArgument, "name %q is not valid", name)
	}
}
