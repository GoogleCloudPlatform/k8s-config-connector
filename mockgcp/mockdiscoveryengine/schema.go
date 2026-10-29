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
// proto.service: google.cloud.discoveryengine.v1.SchemaService
// proto.message: google.cloud.discoveryengine.v1.Schema

package mockdiscoveryengine

import (
	"bytes"
	"context"
	"encoding/json"
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
	"github.com/GoogleCloudPlatform/k8s-config-connector/mockgcp/pkg/storage"
)

type schemaService struct {
	*MockService
	pb.UnimplementedSchemaServiceServer
}

func (s *schemaService) GetSchema(ctx context.Context, req *pb.GetSchemaRequest) (*pb.Schema, error) {
	name, err := s.parseSchemaName(req.GetName())
	if err != nil {
		return nil, err
	}

	fqn := name.String()
	obj := &pb.Schema{}
	if err := s.storage.Get(ctx, fqn, obj); err != nil {
		if status.Code(err) == codes.NotFound {
			return nil, status.Errorf(codes.NotFound, "Schema with name %s does not exist.", fqn)
		}
		return nil, err
	}
	return obj, nil
}

func (s *schemaService) CreateSchema(ctx context.Context, req *pb.CreateSchemaRequest) (*longrunningpb.Operation, error) {
	parent, err := s.parseDataStoreParent(req.GetParent())
	if err != nil {
		return nil, err
	}

	var schemaExists bool
	prefix := fmt.Sprintf("%s/schemas/", parent.String())
	findKind := (&pb.Schema{}).ProtoReflect().Descriptor()
	if err := s.storage.List(ctx, findKind, storage.ListOptions{Prefix: prefix}, func(obj proto.Message) error {
		schemaExists = true
		return nil
	}); err != nil {
		return nil, err
	}
	if schemaExists {
		return nil, status.Errorf(codes.AlreadyExists, "Active schema already exists for data store %s", parent.String())
	}

	reqName := fmt.Sprintf("%s/schemas/%s", req.GetParent(), req.GetSchemaId())
	name, err := s.parseSchemaName(reqName)
	if err != nil {
		return nil, err
	}

	fqn := name.String()
	obj := proto.Clone(req.GetSchema()).(*pb.Schema)
	obj.Name = fqn

	if js := obj.GetJsonSchema(); js != "" {
		var buf bytes.Buffer
		if err := json.Compact(&buf, []byte(js)); err == nil {
			obj.Schema = &pb.Schema_JsonSchema{JsonSchema: buf.String()}
		}
	}

	if err := s.storage.Create(ctx, fqn, obj); err != nil {
		return nil, err
	}

	prefixLRO := fqn
	return s.operations.DoneLRO(ctx, prefixLRO, &pb.CreateSchemaMetadata{}, obj)
}

func (s *schemaService) UpdateSchema(ctx context.Context, req *pb.UpdateSchemaRequest) (*longrunningpb.Operation, error) {
	name, err := s.parseSchemaName(req.GetSchema().GetName())
	if err != nil {
		return nil, err
	}

	fqn := name.String()
	obj := &pb.Schema{}
	err = s.storage.Get(ctx, fqn, obj)
	if err != nil {
		if status.Code(err) == codes.NotFound {
			if !req.GetAllowMissing() {
				return nil, status.Errorf(codes.NotFound, "Schema with name %q does not exist.", fqn)
			}
		} else {
			return nil, err
		}
	}

	prefix := fqn
	metadata := &pb.UpdateSchemaMetadata{}
	return s.operations.StartLRO(ctx, prefix, metadata, func() (proto.Message, error) {
		now := time.Now()
		metadata.CreateTime = timestamppb.New(now)
		metadata.UpdateTime = timestamppb.New(now)

		updated := proto.Clone(req.GetSchema()).(*pb.Schema)
		updated.Name = fqn
		if js := updated.GetJsonSchema(); js != "" {
			var buf bytes.Buffer
			if err := json.Compact(&buf, []byte(js)); err == nil {
				updated.Schema = &pb.Schema_JsonSchema{JsonSchema: buf.String()}
			}
		}

		if err != nil && status.Code(err) == codes.NotFound {
			if err := s.storage.Create(ctx, fqn, updated); err != nil {
				return nil, err
			}
		} else {
			if err := s.storage.Update(ctx, fqn, updated); err != nil {
				return nil, err
			}
		}
		return updated, nil
	})
}

func (s *schemaService) DeleteSchema(ctx context.Context, req *pb.DeleteSchemaRequest) (*longrunningpb.Operation, error) {
	name, err := s.parseSchemaName(req.GetName())
	if err != nil {
		return nil, err
	}

	fqn := name.String()
	oldObj := &pb.Schema{}
	if err := s.storage.Delete(ctx, fqn, oldObj); err != nil {
		return nil, err
	}

	dataStoreFQN := fmt.Sprintf("projects/%d/locations/%s/collections/%s/dataStores/%s", name.Project.Number, name.Location, name.Collection, name.DataStore)
	dataStore := &pb.DataStore{}
	if err := s.storage.Get(ctx, dataStoreFQN, dataStore); err == nil {
		if dataStore.DefaultSchemaId == name.Schema {
			dataStore.DefaultSchemaId = ""
			if err := s.storage.Update(ctx, dataStoreFQN, dataStore); err != nil {
				return nil, err
			}
		}
	}

	prefix := fqn
	return s.operations.DoneLRO(ctx, prefix, &pb.DeleteSchemaMetadata{}, nil)
}

func (s *schemaService) ListSchemas(ctx context.Context, req *pb.ListSchemasRequest) (*pb.ListSchemasResponse, error) {
	parent, err := s.parseDataStoreParent(req.GetParent())
	if err != nil {
		return nil, err
	}

	prefix := fmt.Sprintf("%s/schemas/", parent.String())
	response := &pb.ListSchemasResponse{}
	findKind := (&pb.Schema{}).ProtoReflect().Descriptor()
	if err := s.storage.List(ctx, findKind, storage.ListOptions{Prefix: prefix}, func(obj proto.Message) error {
		schema := obj.(*pb.Schema)
		response.Schemas = append(response.Schemas, schema)
		return nil
	}); err != nil {
		return nil, err
	}

	return response, nil
}

type schemaName struct {
	Project    *projects.ProjectData
	Location   string
	Collection string
	DataStore  string
	Schema     string
}

func (n *schemaName) String() string {
	return fmt.Sprintf("projects/%d/locations/%s/collections/%s/dataStores/%s/schemas/%s", n.Project.Number, n.Location, n.Collection, n.DataStore, n.Schema)
}

func (s *MockService) parseSchemaName(name string) (*schemaName, error) {
	tokens := strings.Split(name, "/")
	if len(tokens) == 10 && tokens[0] == "projects" && tokens[2] == "locations" && tokens[4] == "collections" && tokens[6] == "dataStores" && tokens[8] == "schemas" {
		project, err := s.Projects.GetProjectByIDOrNumber(tokens[1])
		if err != nil {
			return nil, err
		}
		return &schemaName{
			Project:    project,
			Location:   tokens[3],
			Collection: tokens[5],
			DataStore:  tokens[7],
			Schema:     tokens[9],
		}, nil
	}
	if len(tokens) == 8 && tokens[0] == "projects" && tokens[2] == "locations" && tokens[4] == "dataStores" && tokens[6] == "schemas" {
		project, err := s.Projects.GetProjectByIDOrNumber(tokens[1])
		if err != nil {
			return nil, err
		}
		return &schemaName{
			Project:    project,
			Location:   tokens[3],
			Collection: "default_collection",
			DataStore:  tokens[5],
			Schema:     tokens[7],
		}, nil
	}
	return nil, status.Errorf(codes.InvalidArgument, "invalid schema name: %q", name)
}

type dataStoreParentName struct {
	Project    *projects.ProjectData
	Location   string
	Collection string
	DataStore  string
}

func (n *dataStoreParentName) String() string {
	return fmt.Sprintf("projects/%d/locations/%s/collections/%s/dataStores/%s", n.Project.Number, n.Location, n.Collection, n.DataStore)
}

func (s *MockService) parseDataStoreParent(parent string) (*dataStoreParentName, error) {
	tokens := strings.Split(parent, "/")
	if len(tokens) == 8 && tokens[0] == "projects" && tokens[2] == "locations" && tokens[4] == "collections" && tokens[6] == "dataStores" {
		project, err := s.Projects.GetProjectByIDOrNumber(tokens[1])
		if err != nil {
			return nil, err
		}
		return &dataStoreParentName{
			Project:    project,
			Location:   tokens[3],
			Collection: tokens[5],
			DataStore:  tokens[7],
		}, nil
	}
	if len(tokens) == 6 && tokens[0] == "projects" && tokens[2] == "locations" && tokens[4] == "dataStores" {
		project, err := s.Projects.GetProjectByIDOrNumber(tokens[1])
		if err != nil {
			return nil, err
		}
		return &dataStoreParentName{
			Project:    project,
			Location:   tokens[3],
			Collection: "default_collection",
			DataStore:  tokens[5],
		}, nil
	}
	return nil, status.Errorf(codes.InvalidArgument, "invalid parent name: %q", parent)
}
