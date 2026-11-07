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
// proto.service: google.cloud.dataplex.v1.BusinessGlossaryService
// proto.message: google.cloud.dataplex.v1.GlossaryCategory

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

	pb "cloud.google.com/go/dataplex/apiv1/dataplexpb"
	"github.com/GoogleCloudPlatform/k8s-config-connector/mockgcp/common/projects"
	"github.com/GoogleCloudPlatform/k8s-config-connector/mockgcp/pkg/storage"
)

func (s *BusinessGlossaryService) GetGlossaryCategory(ctx context.Context, req *pb.GetGlossaryCategoryRequest) (*pb.GlossaryCategory, error) {
	name, err := s.parseGlossaryCategoryName(req.GetName())
	if err != nil {
		return nil, err
	}

	fqn := name.String()

	obj := &pb.GlossaryCategory{}
	if err := s.storage.Get(ctx, fqn, obj); err != nil {
		if status.Code(err) == codes.NotFound {
			return nil, status.Errorf(codes.NotFound, "Category: projects/%d/locations/%s/glossaries/%s/categories/%s not found.", name.Project.Number, name.Location, name.GlossaryID, name.CategoryID)
		}
		return nil, err
	}

	return obj, nil
}

func (s *BusinessGlossaryService) CreateGlossaryCategory(ctx context.Context, req *pb.CreateGlossaryCategoryRequest) (*pb.GlossaryCategory, error) {
	reqName := req.GetParent() + "/categories/" + req.GetCategoryId()
	name, err := s.parseGlossaryCategoryName(reqName)
	if err != nil {
		return nil, err
	}

	if req.GetCategory() == nil {
		return nil, status.Errorf(codes.InvalidArgument, "category is required")
	}

	fqn := name.String()

	now := time.Now()
	obj := proto.Clone(req.GetCategory()).(*pb.GlossaryCategory)
	obj.Name = fqn
	obj.Parent = fmt.Sprintf("projects/%s/locations/%s/glossaries/%s", name.Project.ID, name.Location, name.GlossaryID)
	obj.CreateTime = timestamppb.New(now)
	obj.UpdateTime = timestamppb.New(now)
	obj.Uid = fmt.Sprintf("%d", now.UnixNano())

	if err := s.storage.Create(ctx, fqn, obj); err != nil {
		return nil, err
	}

	return obj, nil
}

func (s *BusinessGlossaryService) UpdateGlossaryCategory(ctx context.Context, req *pb.UpdateGlossaryCategoryRequest) (*pb.GlossaryCategory, error) {
	if req.GetCategory() == nil {
		return nil, status.Errorf(codes.InvalidArgument, "category is required")
	}

	name, err := s.parseGlossaryCategoryName(req.GetCategory().GetName())
	if err != nil {
		return nil, err
	}
	fqn := name.String()

	obj := &pb.GlossaryCategory{}
	if err := s.storage.Get(ctx, fqn, obj); err != nil {
		if status.Code(err) == codes.NotFound {
			return nil, status.Errorf(codes.NotFound, "Category: projects/%d/locations/%s/glossaries/%s/categories/%s not found.", name.Project.Number, name.Location, name.GlossaryID, name.CategoryID)
		}
		return nil, err
	}

	now := time.Now()
	obj.UpdateTime = timestamppb.New(now)

	updateMask := req.GetUpdateMask()
	if updateMask == nil || len(updateMask.GetPaths()) == 0 {
		obj.Description = req.GetCategory().GetDescription()
		obj.DisplayName = req.GetCategory().GetDisplayName()
		obj.Labels = req.GetCategory().GetLabels()
	} else {
		for _, path := range updateMask.GetPaths() {
			switch path {
			case "description":
				obj.Description = req.GetCategory().GetDescription()
			case "display_name", "displayName":
				obj.DisplayName = req.GetCategory().GetDisplayName()
			case "labels":
				obj.Labels = req.GetCategory().GetLabels()
			default:
				return nil, status.Errorf(codes.InvalidArgument, "mock does not implement update of %q", path)
			}
		}
	}

	if err := s.storage.Update(ctx, fqn, obj); err != nil {
		return nil, err
	}

	return obj, nil
}

func (s *BusinessGlossaryService) DeleteGlossaryCategory(ctx context.Context, req *pb.DeleteGlossaryCategoryRequest) (*emptypb.Empty, error) {
	name, err := s.parseGlossaryCategoryName(req.GetName())
	if err != nil {
		return nil, err
	}

	fqn := name.String()

	deleted := &pb.GlossaryCategory{}
	if err := s.storage.Delete(ctx, fqn, deleted); err != nil {
		if status.Code(err) == codes.NotFound {
			return nil, status.Errorf(codes.NotFound, "Category: projects/%d/locations/%s/glossaries/%s/categories/%s not found.", name.Project.Number, name.Location, name.GlossaryID, name.CategoryID)
		}
		return nil, err
	}

	return &emptypb.Empty{}, nil
}

func (s *BusinessGlossaryService) ListGlossaryCategories(ctx context.Context, req *pb.ListGlossaryCategoriesRequest) (*pb.ListGlossaryCategoriesResponse, error) {
	response := &pb.ListGlossaryCategoriesResponse{}

	parentPrefix := strings.TrimSuffix(req.GetParent(), "/") + "/categories/"
	kind := (&pb.GlossaryCategory{}).ProtoReflect().Descriptor()
	if err := s.storage.List(ctx, kind, storage.ListOptions{}, func(obj proto.Message) error {
		category := obj.(*pb.GlossaryCategory)
		if strings.HasPrefix(category.GetName(), parentPrefix) {
			response.Categories = append(response.Categories, category)
		}
		return nil
	}); err != nil {
		return nil, err
	}
	return response, nil
}

type glossaryCategoryName struct {
	Project    *projects.ProjectData
	Location   string
	GlossaryID string
	CategoryID string
}

func (n *glossaryCategoryName) String() string {
	return fmt.Sprintf("projects/%s/locations/%s/glossaries/%s/categories/%s", n.Project.ID, n.Location, n.GlossaryID, n.CategoryID)
}

func (s *MockService) parseGlossaryCategoryName(name string) (*glossaryCategoryName, error) {
	tokens := strings.Split(name, "/")

	if len(tokens) == 8 && tokens[0] == "projects" && tokens[2] == "locations" && tokens[4] == "glossaries" && tokens[6] == "categories" {
		if tokens[1] == "" || tokens[3] == "" || tokens[5] == "" || tokens[7] == "" {
			return nil, status.Errorf(codes.InvalidArgument, "name %q is not valid", name)
		}
		project, err := s.Projects.GetProjectByIDOrNumber(tokens[1])
		if err != nil {
			return nil, err
		}

		return &glossaryCategoryName{
			Project:    project,
			Location:   tokens[3],
			GlossaryID: tokens[5],
			CategoryID: tokens[7],
		}, nil
	}

	return nil, status.Errorf(codes.InvalidArgument, "name %q is not valid", name)
}
