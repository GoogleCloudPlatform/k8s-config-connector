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
// proto.service: google.cloud.dataplex.v1.ContentService
// proto.message: google.cloud.dataplex.v1.Content

package mockdataplex

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

	"github.com/google/uuid"

	// Note: we use the "real" proto (not mockgcp), because the client uses GRPC.
	pb "cloud.google.com/go/dataplex/apiv1/dataplexpb"
)

type ContentService struct {
	*MockService
	pb.UnimplementedContentServiceServer
}

func (s *ContentService) GetContent(ctx context.Context, req *pb.GetContentRequest) (*pb.Content, error) {
	name, err := s.parseContentName(req.Name)
	if err != nil {
		return nil, err
	}

	fqn := name.String()

	obj := &pb.Content{}
	if err := s.storage.Get(ctx, fqn, obj); err != nil {
		if status.Code(err) == codes.NotFound {
			return nil, status.Errorf(codes.NotFound, "content %q not found", name)
		}
		return nil, err
	}

	// Clear data_text for non-GET requests as per API behavior
	if req.View != pb.GetContentRequest_FULL {
		obj.Data = nil // Clear the data oneof
	}

	return obj, nil
}

func (s *ContentService) CreateContent(ctx context.Context, req *pb.CreateContentRequest) (*pb.Content, error) {
	name, err := s.parseContentName(req.GetContent().GetName())
	if err != nil {
		// This parsing should ideally succeed if parent parsing and contentID construction are correct.
		return nil, status.Errorf(codes.Internal, "failed to parse constructed content name %q: %v", name, err)
	}

	fqn := name.String()
	now := time.Now()

	obj := proto.Clone(req.GetContent()).(*pb.Content)
	obj.Name = fqn
	obj.Uid = uuid.NewString()
	obj.CreateTime = timestamppb.New(now)
	obj.UpdateTime = timestamppb.New(now)

	s.populateDefaultsForContent(obj)

	if err := s.storage.Create(ctx, fqn, obj); err != nil {
		return nil, err
	}
	// Clear data_text for non-GET requests as per API behavior
	obj.Data = nil // Clear the data oneof

	return obj, nil
}

func (s *ContentService) populateDefaultsForContent(obj *pb.Content) {
	// No specific defaults mentioned, but ensure oneof fields are valid if needed
	switch obj.Content.(type) {
	case *pb.Content_SqlScript_:
		if obj.GetSqlScript().GetEngine() == pb.Content_SqlScript_QUERY_ENGINE_UNSPECIFIED {
			// Default engine? Let's assume Spark if not specified.
			// obj.GetSqlScript().Engine = pb.Content_SqlScript_SPARK
		}
	case *pb.Content_Notebook_:
		if obj.GetNotebook().GetKernelType() == pb.Content_Notebook_KERNEL_TYPE_UNSPECIFIED {
			// Default kernel? Let's assume PYTHON3 if not specified.
			// obj.GetNotebook().KernelType = pb.Content_Notebook_PYTHON3
		}
	}
}

func (s *ContentService) UpdateContent(ctx context.Context, req *pb.UpdateContentRequest) (*pb.Content, error) {
	name, err := s.parseContentName(req.GetContent().GetName())
	if err != nil {
		return nil, err
	}
	fqn := name.String()

	existing := &pb.Content{}
	if err := s.storage.Get(ctx, fqn, existing); err != nil {
		return nil, err
	}

	now := time.Now()
	updated := proto.Clone(existing).(*pb.Content)
	updated.UpdateTime = timestamppb.New(now)

	paths := req.GetUpdateMask().GetPaths()
	for _, path := range paths {
		switch path {
		case "labels":
			updated.Labels = req.GetContent().GetLabels()
		case "description":
			updated.Description = req.GetContent().GetDescription()
		case "path":
			updated.Path = req.GetContent().GetPath()
		case "data_text":
			updated.Data = &pb.Content_DataText{DataText: req.GetContent().GetDataText()}
		case "sql_script":
			updated.Content = &pb.Content_SqlScript_{SqlScript: req.GetContent().GetSqlScript()}
		case "notebook":
			updated.Content = &pb.Content_Notebook_{Notebook: req.GetContent().GetNotebook()}
		default:
			return nil, status.Errorf(codes.InvalidArgument, "update_mask path %q not valid for Content", path)
		}
	}

	s.populateDefaultsForContent(updated) // Re-apply defaults if necessary after update

	if err := s.storage.Update(ctx, fqn, updated); err != nil {
		return nil, err
	}

	// Clear data_text for non-GET requests as per API behavior
	updated.Data = nil // Clear the data oneof

	return updated, nil
}

func (s *ContentService) DeleteContent(ctx context.Context, req *pb.DeleteContentRequest) (*emptypb.Empty, error) {
	name, err := s.parseContentName(req.Name)
	if err != nil {
		return nil, err
	}

	fqn := name.String()

	deleted := &pb.Content{}
	if err := s.storage.Delete(ctx, fqn, deleted); err != nil {
		return nil, err
	}

	return &emptypb.Empty{}, nil
}

type contentName struct {
	Lake      string
	ContentID string
}

func (n *contentName) String() string {
	return fmt.Sprintf("%s/content/%s", n.Lake, n.ContentID)
}

// parseContentName parses a string into a contentName.
// The expected form is `projects/*/locations/*/lakes/*/content/{content_path}`
func (s *MockService) parseContentName(name string) (*contentName, error) {
	tokens := strings.Split(name, "/")
	if len(tokens) == 8 && tokens[0] == "projects" && tokens[2] == "locations" && tokens[4] == "lakes" && tokens[6] == "content" {
		name := &contentName{
			Lake:      strings.Join(tokens[:len(tokens)-2], "/"),
			ContentID: tokens[7],
		}

		return name, nil
	}

	return nil, status.Errorf(codes.InvalidArgument, "name %q is not valid", name)
}
