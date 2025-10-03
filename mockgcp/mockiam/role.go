// Copyright 2025 Google LLC
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
// proto.service: google.iam.admin.v1.IAM
// proto.message: google.iam.admin.v1.Role

package mockiam

import (
	"context"
	"fmt"
	"strings"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	pb "github.com/GoogleCloudPlatform/k8s-config-connector/mockgcp/generated/mockgcp/iam/admin/v1"
	"google.golang.org/protobuf/proto"
)

func (s *IAMServer) GetRole(ctx context.Context, req *pb.GetRoleRequest) (*pb.Role, error) {
	name, err := s.parseRoleName(req.Name)
	if err != nil {
		return nil, err
	}

	fqn := name.String()

	obj := &pb.Role{}
	if err := s.storage.Get(ctx, fqn, obj); err != nil {
		if status.Code(err) == codes.NotFound {
			return nil, status.Errorf(codes.NotFound, "role %q not found", fqn)
		}
		return nil, err
	}

	return obj, nil
}

func (s *IAMServer) CreateRole(ctx context.Context, req *pb.CreateRoleRequest) (*pb.Role, error) {
	reqName := fmt.Sprintf("%s/roles/%s", req.GetParent(), req.GetRoleId())
	name, err := s.parseRoleName(reqName)
	if err != nil {
		return nil, err
	}

	fqn := name.String()

	obj := proto.Clone(req.GetRole()).(*pb.Role)
	obj.Name = fqn
	obj.Etag = computeEtag(obj)

	if err := s.storage.Create(ctx, fqn, obj); err != nil {
		return nil, err
	}

	return obj, nil
}

func (s *IAMServer) UpdateRole(ctx context.Context, req *pb.UpdateRoleRequest) (*pb.Role, error) {
	name, err := s.parseRoleName(req.GetName())
	if err != nil {
		return nil, err
	}

	fqn := name.String()
	obj := &pb.Role{}
	if err := s.storage.Get(ctx, fqn, obj); err != nil {
		return nil, err
	}

	for _, field := range req.GetUpdateMask().GetPaths() {
		switch field {
		case "title":
			obj.Title = req.GetRole().GetTitle()
		case "includedPermissions":
			obj.IncludedPermissions = req.GetRole().GetIncludedPermissions()
		default:
			return nil, fmt.Errorf("mockgcp does not implement update_mask %q", field)
		}
	}
	obj.Etag = computeEtag(obj)

	if err := s.storage.Update(ctx, fqn, obj); err != nil {
		return nil, err
	}

	return obj, nil
}

func (s *IAMServer) DeleteRole(ctx context.Context, req *pb.DeleteRoleRequest) (*pb.Role, error) {
	name, err := s.parseRoleName(req.Name)
	if err != nil {
		return nil, err
	}

	fqn := name.String()

	obj := &pb.Role{}
	if err := s.storage.Get(ctx, fqn, obj); err != nil {
		if status.Code(err) == codes.NotFound {
			return nil, status.Errorf(codes.NotFound, "role %q not found", fqn)
		}
		return nil, err
	}

	obj.Deleted = true

	obj.Etag = computeEtag(obj)

	if err := s.storage.Update(ctx, fqn, obj); err != nil {
		return nil, err
	}

	return obj, nil
}

func (s *IAMServer) UndeleteRole(ctx context.Context, req *pb.UndeleteRoleRequest) (*pb.Role, error) {
	name, err := s.parseRoleName(req.Name)
	if err != nil {
		return nil, err
	}

	fqn := name.String()

	obj := &pb.Role{}
	if err := s.storage.Get(ctx, fqn, obj); err != nil {
		if status.Code(err) == codes.NotFound {
			return nil, status.Errorf(codes.NotFound, "role %q not found", fqn)
		}
		return nil, err
	}

	obj.Deleted = false

	obj.Etag = computeEtag(obj)

	if err := s.storage.Update(ctx, fqn, obj); err != nil {
		return nil, err
	}

	return obj, nil
}

// QueryTestablePermissions returns all/many permissions that can be set
// gcloud calls this endpoint, so we stub-implement it
func (s *IAMServer) QueryTestablePermissions(ctx context.Context, req *pb.QueryTestablePermissionsRequest) (*pb.QueryTestablePermissionsResponse, error) {
	response := &pb.QueryTestablePermissionsResponse{}
	response.Permissions = []*pb.Permission{}
	return response, nil
}

type roleName struct {
	Parent   string
	Resource string
}

func (r *roleName) String() string {
	return r.Parent + "/roles/" + r.Resource
}

// parseRoleName parses a string into a roleName.
func (s *IAMServer) parseRoleName(name string) (*roleName, error) {
	tokens := strings.Split(name, "/")
	if len(tokens) == 4 && tokens[2] == "roles" {
		// Predefined roles have the format "roles/{resource}"
		if tokens[0] == "roles" {
			return &roleName{Parent: tokens[0], Resource: tokens[3]}, nil
		}

		if tokens[0] == "projects" || tokens[0] == "organizations" {
			name := &roleName{
				Parent:   strings.Join(tokens[0:2], "/"),
				Resource: tokens[3],
			}

			return name, nil
		}
	}

	return nil, status.Errorf(codes.InvalidArgument, "name %q is not valid", name)
}
