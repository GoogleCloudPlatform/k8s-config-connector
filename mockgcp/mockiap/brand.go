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
// proto.service: google.cloud.iap.v1.IdentityAwareProxyOAuthService
// proto.message: google.cloud.iap.v1.Brand

package mockiap

import (
	"context"
	"fmt"
	"strconv"
	"strings"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/proto"

	pb "cloud.google.com/go/iap/apiv1/iappb"
	"github.com/GoogleCloudPlatform/k8s-config-connector/mockgcp/common/projects"
	"github.com/GoogleCloudPlatform/k8s-config-connector/mockgcp/pkg/storage"
)

// CreateBrand creates a new OAuth brand for the project.
func (s *IdentityAwareProxyOAuthService) CreateBrand(ctx context.Context, req *pb.CreateBrandRequest) (*pb.Brand, error) {
	parentName, err := s.parseBrandName(req.GetParent() + "/brands/dummy")
	if err != nil {
		return nil, err
	}

	name, err := s.parseBrandName(req.GetParent() + "/brands/" + strconv.FormatInt(parentName.Project.Number, 10))
	if err != nil {
		return nil, err
	}

	// Check if a brand already exists for this project. Only one is allowed.
	fqn := name.String()

	existing := &pb.Brand{}
	if err := s.storage.Get(ctx, fqn, existing); err == nil {
		return nil, status.Errorf(codes.AlreadyExists, "A brand already exists for the project: %s", parentName.Project.ID)
	} else if status.Code(err) != codes.NotFound {
		return nil, status.Errorf(codes.Internal, "failed to check for existing brand: %v", err)
	}

	obj := ProtoClone(req.GetBrand())
	obj.Name = name.String()
	obj.OrgInternalOnly = true

	if err := s.storage.Create(ctx, fqn, obj); err != nil {
		return nil, status.Errorf(codes.Internal, "failed to create brand: %v", err)
	}

	return obj, nil
}

// GetBrand retrieves the OAuth brand for the project.
func (s *IdentityAwareProxyOAuthService) GetBrand(ctx context.Context, req *pb.GetBrandRequest) (*pb.Brand, error) {
	name, err := s.parseBrandName(req.GetName())
	if err != nil {
		return nil, err
	}

	fqn := name.String()

	obj := &pb.Brand{}
	if err := s.storage.Get(ctx, fqn, obj); err != nil {
		if status.Code(err) == codes.NotFound {
			return nil, status.Errorf(codes.NotFound, "Brand for project %q not found.", name.Project.ID)
		}
		return nil, status.Errorf(codes.Internal, "failed to get brand: %v", err)
	}

	return obj, nil
}

// ListBrands lists the brand for a project (at most one).
func (s *IdentityAwareProxyOAuthService) ListBrands(ctx context.Context, req *pb.ListBrandsRequest) (*pb.ListBrandsResponse, error) {
	name, err := s.parseBrandName(req.GetParent() + "/brands/dummy")
	if err != nil {
		return nil, err
	}

	prefix := strings.TrimSuffix(name.String(), "dummy")

	response := &pb.ListBrandsResponse{}

	brandKind := (&pb.Brand{}).ProtoReflect().Descriptor()
	if err := s.storage.List(ctx, brandKind, storage.ListOptions{Prefix: prefix}, func(obj proto.Message) error {
		brand := obj.(*pb.Brand)
		response.Brands = append(response.Brands, brand)
		return nil
	}); err != nil {
		return nil, err
	}

	return response, nil

}

// brandName is a parsed IAP brand resource name.
type brandName struct {
	Project *projects.ProjectData
	Brand   string
}

// String returns the string representation of the Brand's name, which includes the project number.
// Format: projects/{project_number}/brands/{brand}
func (n *brandName) String() string {
	return fmt.Sprintf("projects/%d/brands/%s", n.Project.Number, n.Brand)
}

// parseBrandResourceName parses a string in the format "projects/{project}/brands/{brand_id}".
func (s *MockService) parseBrandName(name string) (*brandName, error) {
	tokens := strings.Split(name, "/")
	if len(tokens) == 4 && tokens[0] == "projects" && tokens[2] == "brands" {
		project, err := s.Projects.GetProjectByIDOrNumber(tokens[1])
		if err != nil {
			return nil, status.Errorf(codes.NotFound, "project %q not found", tokens[1])
		}

		brandID := tokens[3]

		// The brandID is typically the project number, but this is not something we enforce here.

		return &brandName{Project: project, Brand: brandID}, nil
	}
	return nil, status.Errorf(codes.InvalidArgument, "invalid brand name format: %q", name)
}
