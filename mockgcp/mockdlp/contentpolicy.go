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

package mockdlp

import (
	"context"
	"strings"

	pb "cloud.google.com/go/dlp/apiv2/dlppb"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/types/known/emptypb"
	"google.golang.org/protobuf/types/known/timestamppb"

	"github.com/GoogleCloudPlatform/k8s-config-connector/mockgcp/common/fields"
	"github.com/GoogleCloudPlatform/k8s-config-connector/mockgcp/common/projects"
	"github.com/GoogleCloudPlatform/k8s-config-connector/mockgcp/pkg/storage"
)

type dlpServiceServer struct {
	*MockService
	pb.UnimplementedDlpServiceServer
}

func newDLPServiceServer(s *MockService) *dlpServiceServer {
	return &dlpServiceServer{MockService: s}
}

// CreateContentPolicy creates a ContentPolicy.
func (s *dlpServiceServer) CreateContentPolicy(ctx context.Context, req *pb.CreateContentPolicyRequest) (*pb.ContentPolicy, error) {
	contentPolicyID := req.GetContentPolicyId()
	if contentPolicyID == "" {
		return nil, status.Errorf(codes.InvalidArgument, "content_policy_id is required")
	}

	parent, err := s.parseParent(req.GetParent())
	if err != nil {
		return nil, err
	}

	name := &contentPolicyName{
		Project:         parent.Project,
		Location:        parent.Location,
		ContentPolicyID: contentPolicyID,
	}

	if req.GetContentPolicy() == nil {
		return nil, status.Errorf(codes.InvalidArgument, "content_policy is required")
	}
	if req.GetContentPolicy().GetInspectConfig() == nil {
		return nil, status.Errorf(codes.InvalidArgument, "inspect_config is required.")
	}

	fqn := name.String()

	existing := &pb.ContentPolicy{}
	if err := s.storage.Get(ctx, fqn, existing); err == nil {
		return nil, status.Errorf(codes.AlreadyExists, "ContentPolicy with name %s already exists.", fqn)
	}

	now := timestamppb.Now()
	obj := proto.Clone(req.GetContentPolicy()).(*pb.ContentPolicy)
	obj.Name = fqn
	obj.CreateTime = now
	obj.UpdateTime = now

	populateDefaultsForContentPolicy(obj)

	if err := s.storage.Create(ctx, fqn, obj); err != nil {
		return nil, err
	}

	return obj, nil
}

type parentName struct {
	Project  *projects.ProjectData
	Location string
}

func (n *parentName) String() string {
	return "projects/" + n.Project.ID + "/locations/" + n.Location
}

func (s *dlpServiceServer) parseParent(parent string) (*parentName, error) {
	tokens := strings.Split(parent, "/")
	if len(tokens) == 4 && tokens[0] == "projects" && tokens[2] == "locations" {
		projectName, err := projects.ParseProjectName("projects/" + tokens[1])
		if err != nil {
			return nil, err
		}

		project, err := s.Projects.GetProject(projectName)
		if err != nil {
			return nil, err
		}

		return &parentName{
			Project:  project,
			Location: tokens[3],
		}, nil
	}
	return nil, status.Errorf(
		codes.InvalidArgument,
		"parent %q is not valid. expected format is projects/<projectID>/locations/<location>",
		parent,
	)
}

type contentPolicyName struct {
	Project         *projects.ProjectData
	Location        string
	ContentPolicyID string
}

func (n *contentPolicyName) String() string {
	return "projects/" + n.Project.ID + "/locations/" + n.Location + "/contentPolicies/" + n.ContentPolicyID
}

func (s *dlpServiceServer) parseContentPolicyName(name string) (*contentPolicyName, error) {
	tokens := strings.Split(name, "/")
	if len(tokens) == 6 && tokens[0] == "projects" && tokens[2] == "locations" && tokens[4] == "contentPolicies" {
		projectName, err := projects.ParseProjectName("projects/" + tokens[1])
		if err != nil {
			return nil, err
		}

		project, err := s.Projects.GetProject(projectName)
		if err != nil {
			return nil, err
		}

		return &contentPolicyName{
			Project:         project,
			Location:        tokens[3],
			ContentPolicyID: tokens[5],
		}, nil
	}
	return nil, status.Errorf(
		codes.InvalidArgument,
		"name %q is not valid. expected format is projects/<projectID>/locations/<location>/contentPolicies/<contentPolicyID>",
		name,
	)
}

// populateDefaultsForContentPolicy populates default values on both write and read paths.
// Note: If default_action, failed_to_scan_supported_file_type, input_too_large, or
// unsupported_file_type are not configured, they remain unspecified (zero-value)
// and are omitted by proto3 JSON serialization, matching real GCP API behavior.
func populateDefaultsForContentPolicy(obj *pb.ContentPolicy) {
	if obj.InspectConfig != nil {
		if obj.InspectConfig.Limits == nil {
			obj.InspectConfig.Limits = &pb.InspectConfig_FindingLimits{}
		}
	}
}

// GetContentPolicy gets a ContentPolicy.
func (s *dlpServiceServer) GetContentPolicy(ctx context.Context, req *pb.GetContentPolicyRequest) (*pb.ContentPolicy, error) {
	name, err := s.parseContentPolicyName(req.GetName())
	if err != nil {
		return nil, err
	}

	fqn := name.String()
	obj := &pb.ContentPolicy{}
	if err := s.storage.Get(ctx, fqn, obj); err != nil {
		if status.Code(err) == codes.NotFound {
			return nil, status.Errorf(codes.NotFound, "ContentPolicy with name %s not found.", fqn)
		}
		return nil, err
	}

	populateDefaultsForContentPolicy(obj)
	return obj, nil
}

// UpdateContentPolicy updates a ContentPolicy.
func (s *dlpServiceServer) UpdateContentPolicy(ctx context.Context, req *pb.UpdateContentPolicyRequest) (*pb.ContentPolicy, error) {
	name, err := s.parseContentPolicyName(req.GetName())
	if err != nil {
		return nil, err
	}

	fqn := name.String()
	existing := &pb.ContentPolicy{}
	if err := s.storage.Get(ctx, fqn, existing); err != nil {
		if status.Code(err) == codes.NotFound {
			return nil, status.Errorf(codes.NotFound, "ContentPolicy with name %s not found.", fqn)
		}
		return nil, err
	}

	updated := proto.Clone(existing).(*pb.ContentPolicy)

	if req.GetContentPolicy() != nil {
		paths := req.GetUpdateMask().GetPaths()
		if len(paths) == 0 {
			updated = proto.Clone(req.GetContentPolicy()).(*pb.ContentPolicy)
			updated.CreateTime = existing.CreateTime
		} else {
			if err := fields.UpdateByFieldMask(updated, req.GetContentPolicy(), paths); err != nil {
				return nil, status.Errorf(codes.InvalidArgument, "applying update_mask: %v", err)
			}
		}
	}

	updated.Name = fqn
	updated.UpdateTime = timestamppb.Now()
	populateDefaultsForContentPolicy(updated)

	if err := s.storage.Update(ctx, fqn, updated); err != nil {
		return nil, err
	}

	return updated, nil
}

// DeleteContentPolicy deletes a ContentPolicy.
func (s *dlpServiceServer) DeleteContentPolicy(ctx context.Context, req *pb.DeleteContentPolicyRequest) (*emptypb.Empty, error) {
	name, err := s.parseContentPolicyName(req.GetName())
	if err != nil {
		return nil, err
	}

	fqn := name.String()
	oldObj := &pb.ContentPolicy{}
	if err := s.storage.Delete(ctx, fqn, oldObj); err != nil {
		if status.Code(err) == codes.NotFound {
			return nil, status.Errorf(codes.NotFound, "ContentPolicy with name %s not found.", fqn)
		}
		return nil, err
	}

	return &emptypb.Empty{}, nil
}

// ListContentPolicies lists ContentPolicies in a parent.
func (s *dlpServiceServer) ListContentPolicies(ctx context.Context, req *pb.ListContentPoliciesRequest) (*pb.ListContentPoliciesResponse, error) {
	parent, err := s.parseParent(req.GetParent())
	if err != nil {
		return nil, err
	}

	prefix := parent.String() + "/contentPolicies/"
	response := &pb.ListContentPoliciesResponse{}
	findKind := (&pb.ContentPolicy{}).ProtoReflect().Descriptor()
	if err := s.storage.List(ctx, findKind, storage.ListOptions{
		Prefix: prefix,
	}, func(obj proto.Message) error {
		item := proto.Clone(obj).(*pb.ContentPolicy)
		populateDefaultsForContentPolicy(item)
		response.ContentPolicies = append(response.ContentPolicies, item)
		return nil
	}); err != nil {
		return nil, err
	}

	return response, nil
}
