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

package mockcontentwarehouse

import (
	"context"
	"fmt"
	"strings"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/types/known/emptypb"
	"google.golang.org/protobuf/types/known/timestamppb"

	"github.com/GoogleCloudPlatform/k8s-config-connector/mockgcp/common/projects"
	pb "github.com/GoogleCloudPlatform/k8s-config-connector/mockgcp/generated/mockgcp/cloud/contentwarehouse/v1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/mockgcp/pkg/storage"
)

type DocumentSchemaName struct {
	Project        *projects.ProjectName
	Location       string
	DocumentSchema string
}

func (n *DocumentSchemaName) String() string {
	return fmt.Sprintf("projects/%s/locations/%s/documentSchemas/%s", n.Project.OriginalValue, n.Location, n.DocumentSchema)
}

func (s *DocumentSchemasV1) parseDocumentSchemaName(name string) (*DocumentSchemaName, error) {
	tokens := strings.Split(name, "/")

	if len(tokens) == 6 && tokens[0] == "projects" && tokens[2] == "locations" && tokens[4] == "documentSchemas" {
		projectName, err := projects.ParseProjectIDOrNumber(tokens[1])
		if err != nil {
			return nil, err
		}
		return &DocumentSchemaName{
			Project:        projectName,
			Location:       tokens[3],
			DocumentSchema: tokens[5],
		}, nil
	} else {
		return nil, status.Errorf(codes.InvalidArgument, "name %q is not valid", name)
	}
}

func (s *DocumentSchemasV1) canonicalizeName(name string) (string, error) {
	parsed, err := s.parseDocumentSchemaName(name)
	if err != nil {
		return "", err
	}
	proj, err := s.Projects.GetProjectByIDOrNumber(parsed.Project.OriginalValue)
	if err == nil {
		// Canonical format uses project number
		return fmt.Sprintf("projects/%d/locations/%s/documentSchemas/%s", proj.Number, parsed.Location, parsed.DocumentSchema), nil
	}
	return parsed.String(), nil
}

func (s *DocumentSchemasV1) CreateDocumentSchema(ctx context.Context, req *pb.CreateDocumentSchemaRequest) (*pb.DocumentSchema, error) {
	parent := req.GetParent() // projects/{project_number}/locations/{location}
	docSchema := req.GetDocumentSchema()

	if docSchema == nil {
		return nil, status.Errorf(codes.InvalidArgument, "document_schema is required")
	}

	schemaID := ""
	if docSchema.GetName() != "" {
		parsed, err := s.parseDocumentSchemaName(docSchema.GetName())
		if err == nil {
			schemaID = parsed.DocumentSchema
		}
	}
	if schemaID == "" {
		schemaID = "schema-10001"
	}

	fqn := fmt.Sprintf("%s/documentSchemas/%s", parent, schemaID)
	canonicalFQN, err := s.canonicalizeName(fqn)
	if err != nil {
		return nil, err
	}

	docSchema.Name = canonicalFQN
	docSchema.CreateTime = timestamppb.Now()
	docSchema.UpdateTime = timestamppb.Now()

	if err := s.storage.Create(ctx, canonicalFQN, docSchema); err != nil {
		return nil, status.Errorf(codes.Internal, "error creating document schema: %v", err)
	}

	return docSchema, nil
}

func (s *DocumentSchemasV1) GetDocumentSchema(ctx context.Context, req *pb.GetDocumentSchemaRequest) (*pb.DocumentSchema, error) {
	canonicalFQN, err := s.canonicalizeName(req.GetName())
	if err != nil {
		return nil, err
	}

	obj := &pb.DocumentSchema{}
	if err := s.storage.Get(ctx, canonicalFQN, obj); err != nil {
		return nil, s.mapNotFoundError(err, canonicalFQN)
	}

	return obj, nil
}

func (s *DocumentSchemasV1) UpdateDocumentSchema(ctx context.Context, req *pb.UpdateDocumentSchemaRequest) (*pb.DocumentSchema, error) {
	name := req.GetName()
	if name == "" {
		name = req.GetDocumentSchema().GetName()
	}
	canonicalFQN, err := s.canonicalizeName(name)
	if err != nil {
		return nil, err
	}

	docSchema := req.GetDocumentSchema()
	if docSchema == nil {
		return nil, status.Errorf(codes.InvalidArgument, "document_schema is required")
	}

	existing := &pb.DocumentSchema{}
	if err := s.storage.Get(ctx, canonicalFQN, existing); err != nil {
		return nil, s.mapNotFoundError(err, canonicalFQN)
	}

	docSchema.Name = canonicalFQN
	docSchema.CreateTime = existing.CreateTime
	docSchema.UpdateTime = timestamppb.Now()

	if err := s.storage.Update(ctx, canonicalFQN, docSchema); err != nil {
		return nil, status.Errorf(codes.Internal, "error updating document schema: %v", err)
	}

	return docSchema, nil
}

func (s *DocumentSchemasV1) DeleteDocumentSchema(ctx context.Context, req *pb.DeleteDocumentSchemaRequest) (*emptypb.Empty, error) {
	canonicalFQN, err := s.canonicalizeName(req.GetName())
	if err != nil {
		return nil, err
	}

	deletedObject := &pb.DocumentSchema{}
	if err := s.storage.Delete(ctx, canonicalFQN, deletedObject); err != nil {
		return nil, s.mapNotFoundError(err, canonicalFQN)
	}

	return &emptypb.Empty{}, nil
}

func (s *DocumentSchemasV1) mapNotFoundError(err error, name string) error {
	if status.Code(err) == codes.NotFound {
		formattedName := name
		if strings.HasPrefix(name, "projects/") {
			formattedName = "Projects/" + strings.TrimPrefix(name, "projects/")
		}
		return status.Errorf(codes.NotFound, "DocumentSchema %s not found.", formattedName)
	}
	return err
}

func (s *DocumentSchemasV1) ListDocumentSchemas(ctx context.Context, req *pb.ListDocumentSchemasRequest) (*pb.ListDocumentSchemasResponse, error) {
	parent := req.GetParent()
	if parent == "" {
		return nil, status.Errorf(codes.InvalidArgument, "parent must be specified")
	}

	var documentSchemas []*pb.DocumentSchema
	if err := s.storage.List(ctx, (&pb.DocumentSchema{}).ProtoReflect().Descriptor(), storage.ListOptions{
		Prefix: parent + "/documentSchemas/",
	}, func(obj proto.Message) error {
		docSchema := obj.(*pb.DocumentSchema)
		documentSchemas = append(documentSchemas, docSchema)
		return nil
	}); err != nil {
		return nil, err
	}

	return &pb.ListDocumentSchemasResponse{
		DocumentSchemas: documentSchemas,
	}, nil
}
