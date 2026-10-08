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

package mockvmmigration

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

	pb "cloud.google.com/go/vmmigration/apiv1/vmmigrationpb"
	"github.com/GoogleCloudPlatform/k8s-config-connector/mockgcp/common/projects"
	longrunningpb "google.golang.org/genproto/googleapis/longrunning"
)

func (s *VmMigrationV1) GetGroup(ctx context.Context, req *pb.GetGroupRequest) (*pb.Group, error) {
	name, err := s.parseGroupName(req.Name)
	if err != nil {
		return nil, err
	}

	fqn := name.String()

	obj := &pb.Group{}
	if err := s.storage.Get(ctx, fqn, obj); err != nil {
		if status.Code(err) == codes.NotFound {
			return nil, status.Errorf(codes.NotFound, "Resource '%v' was not found", name)
		}
		return nil, err
	}
	return obj, nil
}

func (s *VmMigrationV1) CreateGroup(ctx context.Context, req *pb.CreateGroupRequest) (*longrunningpb.Operation, error) {
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

	if obj.MigrationTargetType == pb.Group_MIGRATION_TARGET_TYPE_UNSPECIFIED {
		obj.MigrationTargetType = pb.Group_MIGRATION_TARGET_TYPE_GCE
	}

	if err := s.storage.Create(ctx, fqn, obj); err != nil {
		return nil, err
	}

	lroPrefix := fmt.Sprintf("projects/%s/locations/%s", name.Project.ID, name.Location)
	lroMetadata := &pb.OperationMetadata{
		ApiVersion: "v1",
		CreateTime: timestamppb.Now(),
		Target:     fqn,
		Verb:       "create",
	}
	return s.operations.StartLRO(ctx, lroPrefix, lroMetadata, func() (proto.Message, error) {
		lroMetadata.EndTime = timestamppb.Now()
		return obj, nil
	})
}

func (s *VmMigrationV1) DeleteGroup(ctx context.Context, req *pb.DeleteGroupRequest) (*longrunningpb.Operation, error) {
	name, err := s.parseGroupName(req.Name)
	if err != nil {
		return nil, err
	}
	fqn := name.String()

	deleted := &pb.Group{}
	if err := s.storage.Delete(ctx, fqn, deleted); err != nil {
		return nil, err
	}

	lroPrefix := fmt.Sprintf("projects/%s/locations/%s", name.Project.ID, name.Location)
	lroMetadata := &pb.OperationMetadata{
		ApiVersion: "v1",
		CreateTime: timestamppb.Now(),
		Target:     fqn,
		Verb:       "delete",
	}
	return s.operations.StartLRO(ctx, lroPrefix, lroMetadata, func() (proto.Message, error) {
		lroMetadata.EndTime = timestamppb.Now()
		return &emptypb.Empty{}, nil
	})
}

func (s *VmMigrationV1) UpdateGroup(ctx context.Context, req *pb.UpdateGroupRequest) (*longrunningpb.Operation, error) {
	name, err := s.parseGroupName(req.GetGroup().GetName())
	if err != nil {
		return nil, err
	}
	fqn := name.String()

	existing := &pb.Group{}
	if err := s.storage.Get(ctx, fqn, existing); err != nil {
		return nil, err
	}
	now := time.Now()
	updated := proto.Clone(existing).(*pb.Group)

	updated.UpdateTime = timestamppb.New(now)

	// Required. The update mask applies to the resource.
	paths := req.GetUpdateMask().GetPaths()
	if len(paths) == 0 {
		return nil, status.Errorf(codes.InvalidArgument, "update_mask must be provided")
	}

	for _, path := range paths {
		switch path {
		case "description":
			updated.Description = req.GetGroup().GetDescription()
		case "displayName", "display_name":
			updated.DisplayName = req.GetGroup().GetDisplayName()
		default:
			return nil, status.Errorf(codes.InvalidArgument, "update_mask path %q not valid", path)
		}
	}

	if err := s.storage.Update(ctx, fqn, updated); err != nil {
		return nil, err
	}

	lroPrefix := fmt.Sprintf("projects/%s/locations/%s", name.Project.ID, name.Location)
	lroMetadata := &pb.OperationMetadata{
		ApiVersion: "v1",
		CreateTime: timestamppb.New(now),
		Target:     fqn,
		Verb:       "update",
	}
	return s.operations.StartLRO(ctx, lroPrefix, lroMetadata, func() (proto.Message, error) {
		lroMetadata.EndTime = timestamppb.Now()
		return updated, nil
	})
}

type groupName struct {
	Project  *projects.ProjectData
	Location string
	Group    string
}

func (n *groupName) String() string {
	return fmt.Sprintf("projects/%s/locations/%s/groups/%s", n.Project.ID, n.Location, n.Group)
}

func (s *MockService) parseGroupName(name string) (*groupName, error) {
	tokens := strings.Split(name, "/")
	if len(tokens) == 6 && tokens[0] == "projects" && tokens[2] == "locations" && tokens[4] == "groups" {
		project, err := s.Projects.GetProjectByID(tokens[1])
		if err != nil {
			return nil, err
		}

		n := &groupName{
			Project:  project,
			Location: tokens[3],
			Group:    tokens[5],
		}
		return n, nil
	}
	return nil, status.Errorf(codes.InvalidArgument, "invalid name %q", name)
}
