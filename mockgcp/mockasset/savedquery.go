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
// proto.service: google.cloud.asset.v1.AssetService
// proto.message: google.cloud.asset.v1.SavedQuery

package mockasset

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

	pb "github.com/GoogleCloudPlatform/k8s-config-connector/mockgcp/generated/mockgcp/cloud/asset/v1"
)

func (s *AssetService) GetSavedQuery(ctx context.Context, req *pb.GetSavedQueryRequest) (*pb.SavedQuery, error) {
	name, err := s.parseSavedQueryName(req.Name)
	if err != nil {
		return nil, err
	}

	obj := &pb.SavedQuery{}
	if err := s.storage.Get(ctx, name.String(), obj); err != nil {
		if status.Code(err) == codes.NotFound {
			return nil, status.Errorf(codes.NotFound, "Requested entity was not found.")
		}
		return nil, err
	}

	return obj, nil
}

func (s *AssetService) CreateSavedQuery(ctx context.Context, req *pb.CreateSavedQueryRequest) (*pb.SavedQuery, error) {
	name, err := s.parseSavedQueryName(req.Parent + "/savedQueries/" + req.GetSavedQueryId())
	if err != nil {
		return nil, err
	}
	reqName := name.String()

	now := time.Now()

	obj := proto.Clone(req.GetSavedQuery()).(*pb.SavedQuery)
	obj.Name = reqName
	obj.CreateTime = timestamppb.New(now)
	obj.LastUpdateTime = timestamppb.New(now)
	// Remove creator and lastUpdater as they're not expected in test output
	//obj.Creator = "test-only@example.com"
	//obj.LastUpdater = "test-only@example.com"

	if err := s.storage.Create(ctx, reqName, obj); err != nil {
		return nil, err
	}

	return obj, nil
}

func (s *AssetService) UpdateSavedQuery(ctx context.Context, req *pb.UpdateSavedQueryRequest) (*pb.SavedQuery, error) {
	name, err := s.parseSavedQueryName(req.GetSavedQuery().GetName())
	if err != nil {
		return nil, err
	}

	fqn := name.String()
	obj := &pb.SavedQuery{}
	if err := s.storage.Get(ctx, fqn, obj); err != nil {
		return nil, err
	}

	if req.UpdateMask != nil {
		for _, path := range req.UpdateMask.Paths {
			switch path {
			case "content":
				obj.Content = req.GetSavedQuery().Content
			case "description":
				obj.Description = req.GetSavedQuery().Description
			case "labels":
				obj.Labels = req.GetSavedQuery().Labels
			default:
				return nil, fmt.Errorf("unexpected field mask path: %q", path)
			}
		}
	}
	obj.LastUpdateTime = timestamppb.New(time.Now())
	//obj.LastUpdater = "test-only@example.com"

	if err := s.storage.Update(ctx, fqn, obj); err != nil {
		return nil, err
	}

	return obj, nil
}

func (s *AssetService) DeleteSavedQuery(ctx context.Context, req *pb.DeleteSavedQueryRequest) (*emptypb.Empty, error) {
	name, err := s.parseSavedQueryName(req.Name)
	if err != nil {
		return nil, err
	}

	deleted := &pb.SavedQuery{}
	if err := s.storage.Delete(ctx, name.String(), deleted); err != nil {
		if status.Code(err) == codes.NotFound {
			return nil, status.Errorf(codes.NotFound, "savedQuery %q not found", req.Name)
		}
		return nil, err
	}

	return &emptypb.Empty{}, nil
}

type savedQueryName struct {
	projectNumber  int64
	folderID       string
	organizationID string
	savedQueryID   string
}

func (n *savedQueryName) String() string {
	if n.organizationID != "" {
		return fmt.Sprintf("organizations/%s/savedQueries/%s", n.organizationID, n.savedQueryID)
	}
	if n.folderID != "" {
		return fmt.Sprintf("folders/%s/savedQueries/%s", n.folderID, n.savedQueryID)
	}
	return fmt.Sprintf("projects/%d/savedQueries/%s", n.projectNumber, n.savedQueryID)
}

// parseSavedQueryName parses a string into an savedQueryName.
// The expected form is `projects/*/savedQueries/*`.
func (s *MockService) parseSavedQueryName(name string) (*savedQueryName, error) {
	tokens := strings.Split(name, "/")

	if len(tokens) != 4 || tokens[2] != "savedQueries" {
		return nil, status.Errorf(codes.InvalidArgument, "name %q is not valid", name)
	}

	savedQueryName := &savedQueryName{}
	savedQueryName.savedQueryID = tokens[3]
	switch tokens[0] {
	case "projects":
		project, err := s.Projects.GetProjectByIDOrNumber(tokens[1])
		if err != nil {
			return nil, err
		}
		savedQueryName.projectNumber = project.Number
	case "folders":
		savedQueryName.folderID = tokens[1]
	case "organizations":
		savedQueryName.organizationID = tokens[1]
	default:
		return nil, status.Errorf(codes.InvalidArgument, "name %q is not valid", name)
	}

	return savedQueryName, nil
}
