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
	"net/url"
	"strings"
	"time"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/types/known/emptypb"
	"google.golang.org/protobuf/types/known/timestamppb"

	"github.com/GoogleCloudPlatform/k8s-config-connector/mockgcp/common/projects"
	"github.com/GoogleCloudPlatform/k8s-config-connector/mockgcp/pkg/storage"
	"github.com/google/uuid"

	// Note: we use the "real" proto (not mockgcp), because the client uses GRPC.
	pb "cloud.google.com/go/dataplex/apiv1/dataplexpb"
	iampb "google.golang.org/genproto/googleapis/iam/v1"
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
	parentName, err := s.parseLakeName(req.GetParent())
	if err != nil {
		return nil, err
	}

	// The ID is derived from the path. Path needs URL encoding.
	// Path might contain slashes, so the ID needs careful construction.
	// Let's assume the API automatically handles the full path as the {content} segment identifier.
	// We need to URL-escape the path to form the final name segment safely.
	contentID := url.PathEscape(req.GetContent().GetPath())
	// Double-encode slashes in the path as per API behavior
	contentID = strings.ReplaceAll(contentID, "%2F", "%252F")

	reqName := fmt.Sprintf("%s/content/%s", req.GetParent(), contentID)

	name, err := s.parseContentName(reqName)
	if err != nil {
		// This parsing should ideally succeed if parent parsing and contentID construction are correct.
		return nil, status.Errorf(codes.Internal, "failed to parse constructed content name %q: %v", reqName, err)
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
	case *pb.Content_SqlScript:
		if obj.GetSqlScript().GetEngine() == pb.Content_SqlScript_QUERY_ENGINE_UNSPECIFIED {
			// Default engine? Let's assume Spark if not specified.
			// obj.GetSqlScript().Engine = pb.Content_SqlScript_SPARK
		}
	case *pb.Content_Notebook:
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

	// Per API documentation: Only supports full resource update.
	// However, typically `update_mask` is used. Let's follow the mask if provided.
	paths := req.GetUpdateMask().GetPaths()
	if len(paths) == 0 {
		// If no mask, assume full update of mutable fields based on documentation
		updated.Labels = req.GetContent().GetLabels()
		updated.Description = req.GetContent().GetDescription()
		updated.Data = req.GetContent().GetData()       // Update the data oneof
		updated.Content = req.GetContent().GetContent() // Update the content oneof
	} else {
		// Apply partial updates based on the mask
		// TODO: Use a fieldmask helper if available
		for _, path := range paths {
			switch path {
			case "labels":
				updated.Labels = req.GetContent().GetLabels()
			case "description":
				updated.Description = req.GetContent().GetDescription()
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

func (s *ContentService) ListContent(ctx context.Context, req *pb.ListContentRequest) (*pb.ListContentResponse, error) {
	parentName, err := s.parseLakeName(req.GetParent())
	if err != nil {
		return nil, err
	}

	parentFQNPrefix := parentName.String() + "/content/"

	response := &pb.ListContentResponse{}

	contentKind := (&pb.Content{}).ProtoReflect().Descriptor()
	if err := s.storage.List(ctx, contentKind, storage.ListOptions{}, func(obj proto.Message) error {
		content := obj.(*pb.Content)
		if strings.HasPrefix(content.GetName(), parentFQNPrefix) {
			// Clear data_text for list requests as per API behavior
			content.Data = nil
			response.Content = append(response.Content, content)
		}
		return nil
	}); err != nil {
		return nil, err
	}

	// TODO: Implement filtering if req.Filter is set
	// TODO: Implement pagination

	return response, nil
}

func (s *ContentService) GetIamPolicy(ctx context.Context, req *iampb.GetIamPolicyRequest) (*iampb.Policy, error) {
	// Mock implementation - return empty policy
	return &iampb.Policy{}, nil
}

func (s *ContentService) SetIamPolicy(ctx context.Context, req *iampb.SetIamPolicyRequest) (*iampb.Policy, error) {
	// Mock implementation - return the policy passed in
	return req.Policy, nil
}

func (s *ContentService) TestIamPermissions(ctx context.Context, req *iampb.TestIamPermissionsRequest) (*iampb.TestIamPermissionsResponse, error) {
	// Mock implementation - return all permissions passed in
	return &iampb.TestIamPermissionsResponse{Permissions: req.Permissions}, nil
}

type contentName struct {
	Project  *projects.ProjectData
	Location string
	Lake     string
	// Content can contain slashes, representing a path.
	ContentPath string
}

func (n *contentName) String() string {
	// Use the URL-escaped and slash-encoded version for the final segment
	escapedContentPath := url.PathEscape(n.ContentPath)
	encodedContentPath := strings.ReplaceAll(escapedContentPath, "%2F", "%252F")
	return fmt.Sprintf("projects/%s/locations/%s/lakes/%s/content/%s", n.Project.ID, n.Location, n.Lake, encodedContentPath)
}

// parseContentName parses a string into a contentName.
// The expected form is `projects/*/locations/*/lakes/*/content/{content_path}`
// or `projects/*/locations/*/lakes/*/contentitems/{content_path}`
// where {content_path} can contain multiple segments separated by '/'.
func (s *MockService) parseContentName(name string) (*contentName, error) {
	pattern1Prefix := "projects/"
	pattern1Suffix := "/content/"
	pattern2Suffix := "/contentitems/"

	var projectID, location, lake, contentPath string
	var tokens []string

	if strings.Contains(name, pattern1Suffix) {
		parts := strings.SplitN(name, pattern1Suffix, 2)
		prefixPart := parts[0]
		contentPath = parts[1]
		tokens = strings.Split(prefixPart, "/")
		if len(tokens) != 6 || tokens[0] != "projects" || tokens[2] != "locations" || tokens[4] != "lakes" {
			return nil, status.Errorf(codes.InvalidArgument, "name %q does not match pattern projects/*/locations/*/lakes/*/content/**", name)
		}
	} else if strings.Contains(name, pattern2Suffix) {
		parts := strings.SplitN(name, pattern2Suffix, 2)
		prefixPart := parts[0]
		contentPath = parts[1]
		tokens = strings.Split(prefixPart, "/")
		if len(tokens) != 6 || tokens[0] != "projects" || tokens[2] != "locations" || tokens[4] != "lakes" {
			return nil, status.Errorf(codes.InvalidArgument, "name %q does not match pattern projects/*/locations/*/lakes/*/contentitems/**", name)
		}
	} else {
		return nil, status.Errorf(codes.InvalidArgument, "name %q does not contain /content/ or /contentitems/", name)
	}

	projectID = tokens[1]
	location = tokens[3]
	lake = tokens[5]

	// URL-decode the content path, including the double-encoded slashes
	decodedContentPath, err := url.PathUnescape(strings.ReplaceAll(contentPath, "%252F", "%2F"))
	if err != nil {
		return nil, status.Errorf(codes.InvalidArgument, "invalid content path encoding in name %q: %v", name, err)
	}

	project, err := s.Projects.GetProjectByID(projectID)
	if err != nil {
		return nil, err // Or wrap with context
	}

	parsedName := &contentName{
		Project:     project,
		Location:    location,
		Lake:        lake,
		ContentPath: decodedContentPath,
	}

	return parsedName, nil
}
