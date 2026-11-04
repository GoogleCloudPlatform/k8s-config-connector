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
// proto.service: google.cloud.dataplex.v1.CatalogService
// proto.message: google.cloud.dataplex.v1.Entry

package mockdataplex

import (
	"context"
	"fmt"
	"strings"
	"time"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/types/known/timestamppb"

	"github.com/GoogleCloudPlatform/k8s-config-connector/mockgcp/common/projects"

	// Note: we use the "real" proto (not mockgcp), because the client uses GRPC.
	pb "cloud.google.com/go/dataplex/apiv1/dataplexpb"
)

func (s *CatalogService) GetEntry(ctx context.Context, req *pb.GetEntryRequest) (*pb.Entry, error) {
	name, err := s.parseEntryName(req.Name)
	if err != nil {
		return nil, err
	}

	fqn := name.String()

	obj := &pb.Entry{}
	if err := s.storage.Get(ctx, fqn, obj); err != nil {
		if status.Code(err) == codes.NotFound {
			return nil, status.Errorf(codes.NotFound, "Entry `%s` does not exist.", name)
		}
		return nil, err
	}

	if obj.EntryType != "" {
		obj.EntryType = s.formatProjectNumberInLink(obj.EntryType)
	}
	if obj.ParentEntry != "" {
		obj.ParentEntry = s.formatProjectNumberInLink(obj.ParentEntry)
	}

	return obj, nil
}

func (s *CatalogService) CreateEntry(ctx context.Context, req *pb.CreateEntryRequest) (*pb.Entry, error) {
	reqName := fmt.Sprintf("%s/entries/%s", req.GetParent(), req.GetEntryId())
	name, err := s.parseEntryName(reqName)
	if err != nil {
		return nil, err
	}

	fqn := name.String()
	now := time.Now()

	obj := proto.CloneOf(req.GetEntry())
	obj.Name = fqn
	obj.CreateTime = timestamppb.New(now)
	obj.UpdateTime = timestamppb.New(now)

	if obj.EntrySource == nil {
		obj.EntrySource = &pb.EntrySource{}
	}
	if obj.EntrySource.Location == "" {
		obj.EntrySource.Location = name.Location
	}

	if err := s.storage.Create(ctx, fqn, obj); err != nil {
		return nil, err
	}

	return obj, nil
}

func (s *CatalogService) UpdateEntry(ctx context.Context, req *pb.UpdateEntryRequest) (*pb.Entry, error) {
	name, err := s.parseEntryName(req.GetEntry().GetName())
	if err != nil {
		return nil, err
	}
	fqn := name.String()
	now := time.Now()

	obj := &pb.Entry{}
	if err := s.storage.Get(ctx, fqn, obj); err != nil {
		if status.Code(err) == codes.NotFound {
			return nil, status.Errorf(codes.NotFound, "Entry `%s` does not exist.", name)
		}
		return nil, err
	}

	paths := req.GetUpdateMask().GetPaths()
	if len(paths) == 0 {
		proto.Merge(obj, req.GetEntry())
	} else {
		for _, path := range paths {
			switch path {
			case "entry_source", "entrySource":
				obj.EntrySource = req.GetEntry().GetEntrySource()
				if obj.EntrySource != nil && obj.EntrySource.Location == "" {
					obj.EntrySource.Location = name.Location
				}
			case "entry_type", "entryType":
				obj.EntryType = req.GetEntry().GetEntryType()
			case "parent_entry", "parentEntry":
				obj.ParentEntry = req.GetEntry().GetParentEntry()
			case "fully_qualified_name", "fullyQualifiedName":
				obj.FullyQualifiedName = req.GetEntry().GetFullyQualifiedName()
			case "aspects":
				obj.Aspects = req.GetEntry().GetAspects()
			default:
				return nil, status.Errorf(codes.InvalidArgument, "unhandled update mask path %q", path)
			}
		}
	}

	obj.UpdateTime = timestamppb.New(now)

	if err := s.storage.Update(ctx, fqn, obj); err != nil {
		return nil, err
	}

	return obj, nil
}

func (s *CatalogService) DeleteEntry(ctx context.Context, req *pb.DeleteEntryRequest) (*pb.Entry, error) {
	name, err := s.parseEntryName(req.Name)
	if err != nil {
		return nil, err
	}

	fqn := name.String()

	deleted := &pb.Entry{}
	if err := s.storage.Delete(ctx, fqn, deleted); err != nil {
		if status.Code(err) == codes.NotFound {
			return nil, status.Errorf(codes.NotFound, "Entry `%s` does not exist.", name)
		}
		return nil, err
	}

	if deleted.EntryType != "" {
		deleted.EntryType = s.formatProjectNumberInLink(deleted.EntryType)
	}
	if deleted.ParentEntry != "" {
		deleted.ParentEntry = s.formatProjectNumberInLink(deleted.ParentEntry)
	}

	return deleted, nil
}

type entryName struct {
	Project      *projects.ProjectData
	Location     string
	EntryGroupID string
	EntryID      string
}

func (n *entryName) String() string {
	return fmt.Sprintf("projects/%s/locations/%s/entryGroups/%s/entries/%s", n.Project.ID, n.Location, n.EntryGroupID, n.EntryID)
}

func (s *MockService) formatProjectNumberInLink(link string) string {
	if !strings.HasPrefix(link, "projects/") {
		return link
	}
	tokens := strings.Split(link, "/")
	if len(tokens) >= 2 {
		project, err := s.Projects.GetProjectByIDOrNumber(tokens[1])
		if err == nil {
			tokens[1] = fmt.Sprintf("%d", project.Number)
			return strings.Join(tokens, "/")
		}
	}
	return link
}

// parseEntryName parses a string into an entryName.
// The expected form is `projects/*/locations/*/entryGroups/*/entries/*`.
func (s *MockService) parseEntryName(name string) (*entryName, error) {
	tokens := strings.Split(name, "/")

	if len(tokens) == 8 && tokens[0] == "projects" && tokens[2] == "locations" && tokens[4] == "entryGroups" && tokens[6] == "entries" {
		project, err := s.Projects.GetProjectByIDOrNumber(tokens[1])
		if err != nil {
			return nil, err
		}

		return &entryName{
			Project:      project,
			Location:     tokens[3],
			EntryGroupID: tokens[5],
			EntryID:      tokens[7],
		}, nil
	}

	return nil, status.Errorf(codes.InvalidArgument, "name %q is not valid", name)
}
