// Copyright 2024 Google LLC
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
// proto.service: google.firestore.v1.Firestore
// proto.message: google.firestore.v1.Document

package mockfirestore

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

	pb "cloud.google.com/go/firestore/apiv1/firestorepb"
	"github.com/GoogleCloudPlatform/k8s-config-connector/mockgcp/common/projects"
	"github.com/GoogleCloudPlatform/k8s-config-connector/mockgcp/pkg/storage"
	"github.com/google/uuid"
)

type firestoreServer struct {
	*MockService
	pb.UnimplementedFirestoreServer
}


func (s *firestoreServer) GetDocument(ctx context.Context, req *pb.GetDocumentRequest) (*pb.Document, error) {
	name, err := s.parseDocumentName(req.GetName())
	if err != nil {
		return nil, err
	}

	fqn := name.String()

	obj := &pb.Document{}
	if err := s.storage.Get(ctx, fqn, obj); err != nil {
		if status.Code(err) == codes.NotFound {
			return nil, status.Errorf(codes.NotFound, "Document %q not found.", fqn)
		}
		return nil, err
	}

	return obj, nil
}

func (s *firestoreServer) CreateDocument(ctx context.Context, req *pb.CreateDocumentRequest) (*pb.Document, error) {
	docID := req.GetDocumentId()
	if docID == "" {
		// Firestore auto-generates an ID
		docID = uuid.New().String()[:20]
	}

	reqName := fmt.Sprintf("%s/%s/%s", req.GetParent(), req.GetCollectionId(), docID)
	name, err := s.parseDocumentName(reqName)
	if err != nil {
		return nil, err
	}

	fqn := name.String()

	now := time.Now()

	obj := &pb.Document{
		Name:       fqn,
		Fields:     req.GetDocument().GetFields(),
		CreateTime: timestamppb.New(now),
		UpdateTime: timestamppb.New(now),
	}

	if err := s.storage.Create(ctx, fqn, obj); err != nil {
		return nil, err
	}

	return obj, nil
}

func (s *firestoreServer) UpdateDocument(ctx context.Context, req *pb.UpdateDocumentRequest) (*pb.Document, error) {
	name, err := s.parseDocumentName(req.GetDocument().GetName())
	if err != nil {
		return nil, err
	}
	fqn := name.String()

	obj := &pb.Document{}
	if err := s.storage.Get(ctx, fqn, obj); err != nil {
		return nil, err
	}

	// This is a simple merge, a real implementation would need to handle field masks correctly
	// including deleting fields.
	for k, v := range req.GetDocument().GetFields() {
		obj.Fields[k] = v
	}
	obj.UpdateTime = timestamppb.New(time.Now())

	if err := s.storage.Update(ctx, fqn, obj); err != nil {
		return nil, err
	}

	return obj, nil
}

func (s *firestoreServer) DeleteDocument(ctx context.Context, req *pb.DeleteDocumentRequest) (*emptypb.Empty, error) {
	name, err := s.parseDocumentName(req.Name)
	if err != nil {
		return nil, err
	}

	fqn := name.String()

	deleted := &pb.Document{}
	if err := s.storage.Delete(ctx, fqn, deleted); err != nil {
		// if status.Code(err) == codes.NotFound {
		// 	// Deleting a non-existent document is a no-op
		// 	return &emptypb.Empty{}, nil
		// }
		return nil, err
	}

	return &emptypb.Empty{}, nil
}

func (s *firestoreServer) ListDocuments(ctx context.Context, req *pb.ListDocumentsRequest) (*pb.ListDocumentsResponse, error) {
	// The parent is the path to a document, and collection_id is the collection under that document.
	prefix := fmt.Sprintf("%s/%s/", req.GetParent(), req.GetCollectionId())

	response := &pb.ListDocumentsResponse{}

	documentKind := (&pb.Document{}).ProtoReflect().Descriptor()
	if err := s.storage.List(ctx, documentKind, storage.ListOptions{Prefix: prefix}, func(obj proto.Message) error {
		doc := obj.(*pb.Document)
		// Ensure we only list direct children, not documents in sub-collections.
		if !strings.Contains(strings.TrimPrefix(doc.Name, prefix), "/") {
			response.Documents = append(response.Documents, doc)
		}
		return nil
	}); err != nil {
		return nil, err
	}

	return response, nil
}

type documentName struct {
	Project    *projects.ProjectData
	Database   string
	Collection string
	Document   string
}

func (n *documentName) String() string {
	return fmt.Sprintf("projects/%s/databases/%s/documents/%s/%s", n.Project.ID, n.Database, n.Collection, n.Document)
}

// parseDocumentName parses a string into a documentName.
// The expected form is projects/{project}/databases/{database}/documents/{document_path}
func (s *firestoreServer) parseDocumentName(name string) (*documentName, error) {
	tokens := strings.Split(name, "/")
	if len(tokens) == 7 && tokens[0] == "projects" && tokens[2] == "databases" && tokens[4] == "documents" {
		project, err := s.Projects.GetProjectByID(tokens[1])
		if err != nil {
			return nil, err
		}

		n := &documentName{
			Project:    project,
			Database:   tokens[3],
			Collection: tokens[5],
			Document:   tokens[6],
		}
		return n, nil
	}

	return nil, status.Errorf(codes.InvalidArgument, "invalid document name: %q", name)
}
