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

package mockmigrationcenter

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

	"github.com/GoogleCloudPlatform/k8s-config-connector/mockgcp/common/projects"
	pb "github.com/GoogleCloudPlatform/k8s-config-connector/mockgcp/generated/mockgcp/cloud/migrationcenter/v1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/mockgcp/pkg/storage"
	longrunningpb "google.golang.org/genproto/googleapis/longrunning"
)

type groupName struct {
	Project   *projects.ProjectData
	Location  string
	GroupName string
}

func (n *groupName) String() string {
	return fmt.Sprintf("projects/%s/locations/%s/groups/%s", n.Project.ID, n.Location, n.GroupName)
}

func (s *MockService) parseGroupName(name string) (*groupName, error) {
	tokens := strings.Split(name, "/")
	if len(tokens) == 6 && tokens[0] == "projects" && tokens[2] == "locations" && tokens[4] == "groups" {
		project, err := s.Projects.GetProjectByID(tokens[1])
		if err != nil {
			return nil, err
		}
		return &groupName{
			Project:   project,
			Location:  tokens[3],
			GroupName: tokens[5],
		}, nil
	}
	return nil, status.Errorf(codes.InvalidArgument, "name %q is not in expected format", name)
}

func (s *MigrationCenterV1) GetGroup(ctx context.Context, req *pb.GetGroupRequest) (*pb.Group, error) {
	name, err := s.parseGroupName(req.Name)
	if err != nil {
		return nil, err
	}

	fqn := name.String()

	obj := &pb.Group{}
	if err := s.storage.Get(ctx, fqn, obj); err != nil {
		if status.Code(err) == codes.NotFound {
			return nil, status.Errorf(codes.NotFound, "Group %q not found", fqn)
		}
		return nil, err
	}

	return obj, nil
}

func (s *MigrationCenterV1) CreateGroup(ctx context.Context, req *pb.CreateGroupRequest) (*longrunningpb.Operation, error) {
	reqName := fmt.Sprintf("%s/groups/%s", req.GetParent(), req.GetGroupId())
	name, err := s.parseGroupName(reqName)
	if err != nil {
		return nil, err
	}

	fqn := name.String()
	now := time.Now()

	obj := proto.Clone(req.GetGroup()).(*pb.Group)
	obj.Name = fqn
	obj.CreateTime = timestamppb.New(now)
	obj.UpdateTime = timestamppb.New(now)

	if err := s.storage.Create(ctx, fqn, obj); err != nil {
		return nil, err
	}

	lroPrefix := fmt.Sprintf("projects/%s/locations/%s", name.Project.ID, name.Location)
	lroMetadata := &pb.OperationMetadata{
		CreateTime: timestamppb.New(now),
		Target:     name.String(),
		Verb:       "create",
		ApiVersion: "v1",
	}
	return s.operations.StartLRO(ctx, lroPrefix, lroMetadata, func() (proto.Message, error) {
		lroMetadata.EndTime = timestamppb.Now()
		return obj, nil
	})
}

func (s *MigrationCenterV1) UpdateGroup(ctx context.Context, req *pb.UpdateGroupRequest) (*longrunningpb.Operation, error) {
	name, err := s.parseGroupName(req.GetGroup().GetName())
	if err != nil {
		return nil, err
	}
	fqn := name.String()

	obj := &pb.Group{}
	if err := s.storage.Get(ctx, fqn, obj); err != nil {
		return nil, err
	}
	now := time.Now()

	paths := req.GetUpdateMask().GetPaths()
	for _, path := range paths {
		switch path {
		case "labels":
			obj.Labels = req.GetGroup().GetLabels()
		case "description":
			obj.Description = req.GetGroup().GetDescription()
		case "display_name", "displayName":
			obj.DisplayName = req.GetGroup().GetDisplayName()
		default:
			return nil, status.Errorf(codes.InvalidArgument, "update_mask path %q not valid/supported in mock", path)
		}
	}
	obj.UpdateTime = timestamppb.New(now)

	if err := s.storage.Update(ctx, fqn, obj); err != nil {
		return nil, err
	}

	lroPrefix := fmt.Sprintf("projects/%s/locations/%s", name.Project.ID, name.Location)
	lroMetadata := &pb.OperationMetadata{
		CreateTime: timestamppb.New(now),
		Target:     name.String(),
		Verb:       "update",
		ApiVersion: "v1",
	}
	return s.operations.StartLRO(ctx, lroPrefix, lroMetadata, func() (proto.Message, error) {
		lroMetadata.EndTime = timestamppb.Now()
		return obj, nil
	})
}

func (s *MigrationCenterV1) DeleteGroup(ctx context.Context, req *pb.DeleteGroupRequest) (*longrunningpb.Operation, error) {
	name, err := s.parseGroupName(req.Name)
	if err != nil {
		return nil, err
	}
	fqn := name.String()

	obj := &pb.Group{}
	if err := s.storage.Get(ctx, fqn, obj); err != nil {
		if status.Code(err) == codes.NotFound {
			return nil, status.Errorf(codes.NotFound, "Group %q not found", fqn)
		}
		return nil, err
	}

	if err := s.storage.Delete(ctx, fqn, obj); err != nil {
		return nil, err
	}

	now := time.Now()
	lroPrefix := fmt.Sprintf("projects/%s/locations/%s", name.Project.ID, name.Location)
	lroMetadata := &pb.OperationMetadata{
		CreateTime: timestamppb.New(now),
		Target:     name.String(),
		Verb:       "delete",
		ApiVersion: "v1",
	}
	return s.operations.StartLRO(ctx, lroPrefix, lroMetadata, func() (proto.Message, error) {
		lroMetadata.EndTime = timestamppb.Now()
		return &emptypb.Empty{}, nil
	})
}

func (s *MigrationCenterV1) ListGroups(ctx context.Context, req *pb.ListGroupsRequest) (*pb.ListGroupsResponse, error) {
	tokens := strings.Split(req.GetParent(), "/")
	if len(tokens) != 4 || tokens[0] != "projects" || tokens[2] != "locations" {
		return nil, status.Errorf(codes.InvalidArgument, "parent %q is invalid", req.GetParent())
	}

	var groups []*pb.Group
	listKind := (&pb.Group{}).ProtoReflect().Descriptor()
	err := s.storage.List(ctx, listKind, storage.ListOptions{}, func(obj proto.Message) error {
		group := obj.(*pb.Group)
		name, err := s.parseGroupName(group.GetName())
		if err == nil && name.Project.ID == tokens[1] && name.Location == tokens[3] {
			groups = append(groups, group)
		}
		return nil
	})
	if err != nil {
		return nil, err
	}

	return &pb.ListGroupsResponse{
		Groups: groups,
	}, nil
}
