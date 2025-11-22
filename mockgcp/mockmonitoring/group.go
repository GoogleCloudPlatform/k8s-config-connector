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

package mockmonitoring

import (
	"context"
	"fmt"
	"strings"
	"time"

	pb "cloud.google.com/go/monitoring/apiv3/v2/monitoringpb"
	"github.com/GoogleCloudPlatform/k8s-config-connector/mockgcp/common/projects"
	"github.com/golang/protobuf/ptypes/empty"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/proto"
)

type GroupService struct {
	*MockService
	pb.UnimplementedGroupServiceServer
}

func (s *GroupService) GetGroup(ctx context.Context, req *pb.GetGroupRequest) (*pb.Group, error) {
	name, err := s.parseGroupName(req.GetName())
	if err != nil {
		return nil, err
	}

	fqn := name.String()

	obj := &pb.Group{}
	if err := s.storage.Get(ctx, fqn, obj); err != nil {
		if status.Code(err) == codes.NotFound {
			return nil, status.Errorf(codes.NotFound, "Could not find group definition for '%s'.", name.Name)
		}
		return nil, err
	}

	return obj, nil
}

func populateDefaultsForGroup(obj *pb.Group) {
}

func (s *GroupService) CreateGroup(ctx context.Context, req *pb.CreateGroupRequest) (*pb.Group, error) {
	now := time.Now()

	groupID := fmt.Sprintf("%d", now.UnixNano())

	reqName := req.GetName() + "/groups/" + groupID
	name, err := s.parseGroupName(reqName)
	if err != nil {
		return nil, err
	}

	fqn := name.String()

	obj := proto.Clone(req.Group).(*pb.Group)
	obj.Name = fqn

	populateDefaultsForGroup(obj)

	if err := s.storage.Create(ctx, fqn, obj); err != nil {
		return nil, status.Errorf(codes.Internal, "error creating group: %v", err)
	}

	return obj, nil
}

func (s *GroupService) UpdateGroup(ctx context.Context, req *pb.UpdateGroupRequest) (*pb.Group, error) {
	name, err := s.parseGroupName(req.GetGroup().GetName())
	if err != nil {
		return nil, err
	}

	fqn := name.String()

	existing := &pb.Group{}
	if err := s.storage.Get(ctx, fqn, existing); err != nil {
		return nil, err
	}

	// All fields are replaced except name
	updated := proto.Clone(req.GetGroup()).(*pb.Group)
	updated.Name = existing.Name

	populateDefaultsForGroup(updated)

	if err := s.storage.Update(ctx, fqn, updated); err != nil {
		return nil, status.Errorf(codes.Internal, "error updating group: %v", err)
	}

	return updated, nil
}

func (s *GroupService) DeleteGroup(ctx context.Context, req *pb.DeleteGroupRequest) (*empty.Empty, error) {
	name, err := s.parseGroupName(req.Name)
	if err != nil {
		return nil, err
	}

	fqn := name.String()

	deleted := &pb.Group{}
	if err := s.storage.Delete(ctx, fqn, deleted); err != nil {
		return nil, err
	}

	return &empty.Empty{}, nil
}

type groupName struct {
	Project *projects.ProjectData
	Name    string
}

func (n *groupName) String() string {
	return "projects/" + n.Project.ID + "/groups/" + n.Name
}

// parseGroupName parses a string into a groupName.
// The expected form is projects/[PROJECT_ID_OR_NUMBER]/groups/[UPTIME_CHECK_ID]
func (s *MockService) parseGroupName(name string) (*groupName, error) {
	tokens := strings.Split(name, "/")

	if len(tokens) == 4 && tokens[0] == "projects" && tokens[2] == "groups" {
		project, err := s.Projects.GetProjectByID(tokens[1])
		if err != nil {
			return nil, err
		}

		name := &groupName{
			Project: project,
			Name:    tokens[3],
		}

		return name, nil
	} else {
		return nil, status.Errorf(codes.InvalidArgument, "name %q is not valid", name)
	}
}
