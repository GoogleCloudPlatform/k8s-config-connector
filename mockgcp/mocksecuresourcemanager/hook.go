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

package mocksecuresourcemanager

import (
	"context"
	"fmt"
	"strings"
	"time"

	longrunning "google.golang.org/genproto/googleapis/longrunning"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/types/known/emptypb"
	"google.golang.org/protobuf/types/known/timestamppb"

	pb "cloud.google.com/go/securesourcemanager/apiv1/securesourcemanagerpb"
	"github.com/GoogleCloudPlatform/k8s-config-connector/mockgcp/common/fields"
	"github.com/GoogleCloudPlatform/k8s-config-connector/mockgcp/common/projects"
)

func (s *secureSourceManagerServer) GetHook(ctx context.Context, req *pb.GetHookRequest) (*pb.Hook, error) {
	name, err := s.parseHookName(req.Name)
	if err != nil {
		return nil, err
	}

	fqn := name.String()

	obj := &pb.Hook{}
	if err := s.storage.Get(ctx, fqn, obj); err != nil {
		if status.Code(err) == codes.NotFound {
			return nil, status.Errorf(codes.NotFound, "Resource %q was not found", fqn)
		}
		return nil, err
	}

	return obj, nil
}

func (s *secureSourceManagerServer) CreateHook(ctx context.Context, req *pb.CreateHookRequest) (*longrunning.Operation, error) {
	reqName := req.Parent + "/hooks/" + req.HookId
	name, err := s.parseHookName(reqName)
	if err != nil {
		return nil, err
	}

	fqn := name.String()

	now := time.Now()

	obj := proto.CloneOf(req.Hook)
	obj.Name = fqn
	obj.CreateTime = timestamppb.New(now)
	obj.UpdateTime = timestamppb.New(now)

	op := &pb.OperationMetadata{
		CreateTime: timestamppb.New(now),
		EndTime:    timestamppb.New(now),
		Target:     name.String(),
		Verb:       "create",
		ApiVersion: "v1",
	}
	if err := s.storage.Create(ctx, fqn, obj); err != nil {
		return nil, err
	}
	return s.operations.DoneLRO(ctx, "", op, func() *pb.Hook {
		return obj
	}())
}

func (s *secureSourceManagerServer) UpdateHook(ctx context.Context, req *pb.UpdateHookRequest) (*longrunning.Operation, error) {
	name, err := s.parseHookName(req.GetHook().GetName())
	if err != nil {
		return nil, err
	}

	fqn := name.String()

	obj := &pb.Hook{}
	if err := s.storage.Get(ctx, fqn, obj); err != nil {
		if status.Code(err) == codes.NotFound {
			return nil, status.Errorf(codes.NotFound, "Resource %q was not found", fqn)
		}
		return nil, err
	}

	now := time.Now()

	if req.GetUpdateMask() != nil && len(req.GetUpdateMask().GetPaths()) > 0 {
		if err := fields.UpdateByFieldMask(obj, req.GetHook(), req.GetUpdateMask().GetPaths()); err != nil {
			return nil, err
		}
	} else {
		obj = proto.CloneOf(req.GetHook())
	}

	obj.UpdateTime = timestamppb.New(now)

	op := &pb.OperationMetadata{
		CreateTime: timestamppb.New(now),
		EndTime:    timestamppb.New(now),
		Target:     name.String(),
		Verb:       "update",
		ApiVersion: "v1",
	}

	if err := s.storage.Update(ctx, fqn, obj); err != nil {
		return nil, err
	}

	return s.operations.DoneLRO(ctx, "", op, func() *pb.Hook {
		return obj
	}())
}

func (s *secureSourceManagerServer) DeleteHook(ctx context.Context, req *pb.DeleteHookRequest) (*longrunning.Operation, error) {
	name, err := s.parseHookName(req.GetName())
	if err != nil {
		return nil, err
	}

	fqn := name.String()
	now := time.Now()

	deleted := &pb.Hook{}
	if err := s.storage.Delete(ctx, fqn, deleted); err != nil {
		return nil, err
	}

	op := &pb.OperationMetadata{
		CreateTime: timestamppb.New(now),
		EndTime:    timestamppb.New(now),
		Target:     name.String(),
		Verb:       "delete",
		ApiVersion: "v1",
	}

	return s.operations.DoneLRO(ctx, "", op, func() *emptypb.Empty {
		return &emptypb.Empty{}
	}())
}

type HookName struct {
	Project      *projects.ProjectData
	Location     string
	RepositoryID string
	HookID       string
}

func (n *HookName) String() string {
	return fmt.Sprintf("projects/%s/locations/%s/repositories/%s/hooks/%s", n.Project.ID, n.Location, n.RepositoryID, n.HookID)
}

func (s *MockService) parseHookName(name string) (*HookName, error) {
	tokens := strings.Split(name, "/")

	if len(tokens) == 8 && tokens[0] == "projects" && tokens[2] == "locations" && tokens[4] == "repositories" && tokens[6] == "hooks" {
		projectName, err := projects.ParseProjectName(tokens[0] + "/" + tokens[1])
		if err != nil {
			return nil, err
		}
		project, err := s.Projects.GetProject(projectName)
		if err != nil {
			return nil, err
		}

		name := &HookName{
			Project:      project,
			Location:     tokens[3],
			RepositoryID: tokens[5],
			HookID:       tokens[7],
		}
		return name, nil
	} else {
		return nil, status.Errorf(codes.InvalidArgument, "name %q is not valid", name)
	}
}
