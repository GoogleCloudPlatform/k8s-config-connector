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
// proto.service: google.cloud.datacatalog.v1.DataCatalog
// proto.message: google.cloud.datacatalog.v1.EntryGroup

package mockdatacatalog

import (
	"context"
	"strings"
	"time"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/types/known/emptypb"
	"google.golang.org/protobuf/types/known/timestamppb"

	pb "github.com/GoogleCloudPlatform/k8s-config-connector/mockgcp/generated/mockgcp/cloud/datacatalog/v1"
)

func (s *DataCatalogV1) GetEntryGroup(ctx context.Context, req *pb.GetEntryGroupRequest) (*pb.EntryGroup, error) {
	name, err := s.parseEntryGroupName(req.Name)
	if err != nil {
		return nil, err
	}

	fqn := name.String()

	obj := &pb.EntryGroup{}
	if err := s.storage.Get(ctx, fqn, obj); err != nil {
		if status.Code(err) == codes.NotFound {
			return nil, status.Errorf(codes.NotFound, "EntryGroup %%q not found: %s", fqn)
		}
		return nil, err
	}

	return obj, nil
}

func (s *DataCatalogV1) UpdateEntryGroup(ctx context.Context, req *pb.UpdateEntryGroupRequest) (*pb.EntryGroup, error) {
	if req.EntryGroup == nil {
		return nil, status.Errorf(codes.InvalidArgument, "entry_group is required")
	}
	name, err := s.parseEntryGroupName(req.EntryGroup.Name)
	if err != nil {
		return nil, err
	}

	fqn := name.String()

	obj := &pb.EntryGroup{}
	if err := s.storage.Get(ctx, fqn, obj); err != nil {
		if status.Code(err) == codes.NotFound {
			return nil, status.Errorf(codes.NotFound, "EntryGroup %q not found", fqn)
		}
		return nil, err
	}

	// Apply field mask updates.
	if req.UpdateMask != nil && len(req.UpdateMask.Paths) > 0 {
		for _, path := range req.UpdateMask.Paths {
			switch path {
			case "description":
				obj.Description = req.EntryGroup.Description
				// Add other updatable fields here.
			}
		}
	} else {
		proto.Merge(obj, req.EntryGroup)
	}

	// Update the update timestamp
	obj.DataCatalogTimestamps.UpdateTime = timestamppb.New(time.Now())

	if err := s.storage.Update(ctx, fqn, obj); err != nil {
		return nil, err
	}

	return obj, nil
}

func (s *DataCatalogV1) DeleteEntryGroup(ctx context.Context, req *pb.DeleteEntryGroupRequest) (*emptypb.Empty, error) {
	name, err := s.parseEntryGroupName(req.Name)
	if err != nil {
		return nil, err
	}

	fqn := name.String()

	if err := s.storage.Delete(ctx, fqn, &pb.EntryGroup{}); err != nil {
		if status.Code(err) == codes.NotFound {
			return nil, status.Errorf(codes.NotFound, "EntryGroup %q not found", fqn)
		}
		return nil, err
	}

	return &emptypb.Empty{}, nil
}

func (s *DataCatalogV1) CreateEntryGroup(ctx context.Context, req *pb.CreateEntryGroupRequest) (*pb.EntryGroup, error) {
	reqName := req.Parent + "/entryGroups/" + req.EntryGroupId
	name, err := s.parseEntryGroupName(reqName)
	if err != nil {
		return nil, err
	}

	fqn := name.String()

	obj := proto.Clone(req.EntryGroup).(*pb.EntryGroup)
	obj.Name = fqn

	// Add timestamps for creation
	now := timestamppb.New(time.Now())
	obj.DataCatalogTimestamps = &pb.SystemTimestamps{
		CreateTime: now,
		UpdateTime: now,
	}

	if err := s.storage.Create(ctx, fqn, obj); err != nil {
		return nil, err
	}

	return obj, nil
}

type entryGroupName struct {
	Project        string
	Location       string
	EntryGroupName string
}

func (n *entryGroupName) String() string {
	return "projects/" + n.Project + "/locations/" + n.Location + "/entryGroups/" + n.EntryGroupName
}

// parseEntryGroupName parses a string into a entryGroupName.
// The expected form is projects/<projectID>/locations/<location>/entryGroups/<entryGroupID>`.
func (s *DataCatalogV1) parseEntryGroupName(name string) (*entryGroupName, error) {
	tokens := strings.Split(name, "/")
	if len(tokens) == 6 && tokens[0] == "projects" && tokens[2] == "locations" && tokens[4] == "entryGroups" {
		name := &entryGroupName{
			Project:        tokens[1],
			Location:       tokens[3],
			EntryGroupName: tokens[5],
		}

		return name, nil
	}

	return nil, status.Errorf(codes.InvalidArgument, "name %q is not valid", name)
}
