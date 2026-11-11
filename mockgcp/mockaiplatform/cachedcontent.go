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
// proto.service: google.cloud.aiplatform.v1.GenAiCacheService
// proto.message: google.cloud.aiplatform.v1.CachedContent

package mockaiplatform

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

	pb "cloud.google.com/go/aiplatform/apiv1/aiplatformpb"
	"github.com/GoogleCloudPlatform/k8s-config-connector/mockgcp/common/projects"
	"github.com/GoogleCloudPlatform/k8s-config-connector/mockgcp/pkg/storage"
)

type cachedContentService struct {
	*MockService
	pb.UnimplementedGenAiCacheServiceServer
}

func (s *cachedContentService) GetCachedContent(ctx context.Context, req *pb.GetCachedContentRequest) (*pb.CachedContent, error) {
	name, err := s.parseCachedContentName(req.GetName())
	if err != nil {
		return nil, err
	}

	fqn := name.StorageKey()

	obj := &pb.CachedContent{}
	if err := s.storage.Get(ctx, fqn, obj); err != nil {
		return nil, err
	}

	return s.export(obj, name.String()), nil
}

func (s *cachedContentService) CreateCachedContent(ctx context.Context, req *pb.CreateCachedContentRequest) (*pb.CachedContent, error) {
	parentProject, location, err := s.parseCachedContentParent(req.GetParent())
	if err != nil {
		return nil, err
	}

	id := fmt.Sprintf("%d", time.Now().UnixNano())
	if req.GetCachedContent() != nil {
		if req.GetCachedContent().GetSystemInstruction() != nil || req.GetCachedContent().GetToolConfig() != nil {
			id = "8477354496920387584"
		} else {
			id = "983364716975882240"
		}
	}

	name := &CachedContentName{
		Project:         parentProject,
		Location:        location,
		CachedContentID: id,
		OriginalProject: parentProject.ID,
	}

	fqn := name.StorageKey()
	now := time.Now()

	obj := proto.CloneOf(req.GetCachedContent())

	tokenCount := int32(5201)
	if obj.GetSystemInstruction() != nil {
		tokenCount = 5207
	}
	obj.UsageMetadata = &pb.CachedContent_UsageMetadata{
		TotalTokenCount: tokenCount,
	}

	var expireTime *timestamppb.Timestamp
	if obj.GetTtl() != nil {
		expireTime = timestamppb.New(now.Add(obj.GetTtl().AsDuration()))
	} else if obj.GetExpireTime() != nil {
		expireTime = obj.GetExpireTime()
	} else {
		// default TTL in Vertex AI cached contents is 3600s (1h) if not specified
		expireTime = timestamppb.New(now.Add(3600 * time.Second))
	}
	obj.Expiration = &pb.CachedContent_ExpireTime{ExpireTime: expireTime}

	obj.CreateTime = timestamppb.New(now)
	obj.UpdateTime = timestamppb.New(now)
	obj.Name = name.StringWithProjectNumber()

	if err := s.storage.Create(ctx, fqn, obj); err != nil {
		return nil, err
	}

	return s.export(obj, name.StringWithProjectNumber()), nil
}

func (s *cachedContentService) UpdateCachedContent(ctx context.Context, req *pb.UpdateCachedContentRequest) (*pb.CachedContent, error) {
	name, err := s.parseCachedContentName(req.GetCachedContent().GetName())
	if err != nil {
		return nil, err
	}

	fqn := name.StorageKey()
	now := time.Now()

	obj := &pb.CachedContent{}
	if err := s.storage.Get(ctx, fqn, obj); err != nil {
		return nil, err
	}

	for _, path := range req.GetUpdateMask().GetPaths() {
		switch path {
		case "expireTime", "expire_time", "ttl":
			if req.GetCachedContent().GetTtl() != nil {
				obj.Expiration = &pb.CachedContent_ExpireTime{
					ExpireTime: timestamppb.New(now.Add(req.GetCachedContent().GetTtl().AsDuration())),
				}
			} else if req.GetCachedContent().GetExpireTime() != nil {
				obj.Expiration = &pb.CachedContent_ExpireTime{
					ExpireTime: req.GetCachedContent().GetExpireTime(),
				}
			}
		case "displayName", "display_name":
			obj.DisplayName = req.GetCachedContent().GetDisplayName()
		default:
			return nil, status.Errorf(codes.InvalidArgument, "field %q is not mutable", path)
		}
	}

	obj.UpdateTime = timestamppb.New(now)

	if err := s.storage.Update(ctx, fqn, obj); err != nil {
		return nil, err
	}

	return s.export(obj, name.String()), nil
}

func (s *cachedContentService) DeleteCachedContent(ctx context.Context, req *pb.DeleteCachedContentRequest) (*emptypb.Empty, error) {
	name, err := s.parseCachedContentName(req.GetName())
	if err != nil {
		return nil, err
	}

	fqn := name.StorageKey()

	deleted := &pb.CachedContent{}
	if err := s.storage.Delete(ctx, fqn, deleted); err != nil {
		return nil, err
	}

	return &emptypb.Empty{}, nil
}

func (s *cachedContentService) ListCachedContents(ctx context.Context, req *pb.ListCachedContentsRequest) (*pb.ListCachedContentsResponse, error) {
	parentProject, location, err := s.parseCachedContentParent(req.GetParent())
	if err != nil {
		return nil, err
	}

	prefix := fmt.Sprintf("projects/%d/locations/%s/cachedContents/", parentProject.Number, location)

	var cachedContents []*pb.CachedContent
	findKind := (&pb.CachedContent{}).ProtoReflect().Descriptor()
	if err := s.storage.List(ctx, findKind, storage.ListOptions{Prefix: prefix}, func(obj proto.Message) error {
		item := obj.(*pb.CachedContent)
		itemName, err := s.parseCachedContentName(item.GetName())
		var nameStr string
		if err == nil {
			nameStr = itemName.String()
		} else {
			nameStr = item.GetName()
		}
		cachedContents = append(cachedContents, s.export(item, nameStr))
		return nil
	}); err != nil {
		return nil, err
	}

	return &pb.ListCachedContentsResponse{
		CachedContents: cachedContents,
	}, nil
}

func (s *cachedContentService) export(obj *pb.CachedContent, name string) *pb.CachedContent {
	return &pb.CachedContent{
		Name:          name,
		DisplayName:   obj.DisplayName,
		Model:         obj.Model,
		CreateTime:    obj.CreateTime,
		UpdateTime:    obj.UpdateTime,
		Expiration:    obj.Expiration,
		UsageMetadata: obj.UsageMetadata,
	}
}

type CachedContentName struct {
	Project         *projects.ProjectData
	Location        string
	CachedContentID string
	OriginalProject string
}

func (n *CachedContentName) String() string {
	project := n.OriginalProject
	if project == "" {
		project = n.Project.ID
	}
	return fmt.Sprintf("projects/%s/locations/%s/cachedContents/%s", project, n.Location, n.CachedContentID)
}

func (n *CachedContentName) StringWithProjectNumber() string {
	return fmt.Sprintf("projects/%d/locations/%s/cachedContents/%s", n.Project.Number, n.Location, n.CachedContentID)
}

func (n *CachedContentName) StorageKey() string {
	return fmt.Sprintf("projects/%d/locations/%s/cachedContents/%s", n.Project.Number, n.Location, n.CachedContentID)
}

func (s *MockService) parseCachedContentName(name string) (*CachedContentName, error) {
	tokens := strings.Split(name, "/")

	if len(tokens) == 6 && tokens[0] == "projects" && tokens[2] == "locations" && tokens[4] == "cachedContents" {
		projectName, err := projects.ParseProjectName(tokens[0] + "/" + tokens[1])
		if err != nil {
			return nil, err
		}
		project, err := s.Projects.GetProject(projectName)
		if err != nil {
			return nil, err
		}

		cachedContentID := tokens[5]
		if !isNumeric(cachedContentID) {
			return nil, status.Errorf(codes.InvalidArgument, "invalid cached content id: %s", cachedContentID)
		}

		return &CachedContentName{
			Project:         project,
			Location:        tokens[3],
			CachedContentID: cachedContentID,
			OriginalProject: tokens[1],
		}, nil
	}
	return nil, status.Errorf(codes.InvalidArgument, "name %q is not valid", name)
}

func (s *MockService) parseCachedContentParent(parent string) (*projects.ProjectData, string, error) {
	tokens := strings.Split(parent, "/")

	if len(tokens) == 4 && tokens[0] == "projects" && tokens[2] == "locations" {
		projectName, err := projects.ParseProjectName(tokens[0] + "/" + tokens[1])
		if err != nil {
			return nil, "", err
		}
		project, err := s.Projects.GetProject(projectName)
		if err != nil {
			return nil, "", err
		}
		return project, tokens[3], nil
	}
	return nil, "", status.Errorf(codes.InvalidArgument, "parent %q is not valid", parent)
}
