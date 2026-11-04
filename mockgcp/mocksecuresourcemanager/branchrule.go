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

func (s *secureSourceManagerServer) GetBranchRule(ctx context.Context, req *pb.GetBranchRuleRequest) (*pb.BranchRule, error) {
	name, err := s.parseBranchRuleName(req.Name)
	if err != nil {
		return nil, err
	}

	fqn := name.String()

	obj := &pb.BranchRule{}
	if err := s.storage.Get(ctx, fqn, obj); err != nil {
		if status.Code(err) == codes.NotFound {
			return nil, status.Errorf(codes.NotFound, "Resource %q was not found", fqn)
		}
		return nil, err
	}

	return obj, nil
}

func (s *secureSourceManagerServer) CreateBranchRule(ctx context.Context, req *pb.CreateBranchRuleRequest) (*longrunning.Operation, error) {
	reqName := req.Parent + "/branchRules/" + req.BranchRuleId
	name, err := s.parseBranchRuleName(reqName)
	if err != nil {
		return nil, err
	}

	fqn := name.String()

	now := time.Now()

	obj := proto.CloneOf(req.BranchRule)
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
	return s.operations.DoneLRO(ctx, "", op, func() *pb.BranchRule {
		return obj
	}())
}

func (s *secureSourceManagerServer) UpdateBranchRule(ctx context.Context, req *pb.UpdateBranchRuleRequest) (*longrunning.Operation, error) {
	name, err := s.parseBranchRuleName(req.GetBranchRule().GetName())
	if err != nil {
		return nil, err
	}

	fqn := name.String()

	obj := &pb.BranchRule{}
	if err := s.storage.Get(ctx, fqn, obj); err != nil {
		if status.Code(err) == codes.NotFound {
			return nil, status.Errorf(codes.NotFound, "Resource %q was not found", fqn)
		}
		return nil, err
	}

	now := time.Now()

	if req.GetUpdateMask() != nil && len(req.GetUpdateMask().GetPaths()) > 0 {
		if err := fields.UpdateByFieldMask(obj, req.GetBranchRule(), req.GetUpdateMask().GetPaths()); err != nil {
			return nil, err
		}
	} else {
		obj = proto.CloneOf(req.GetBranchRule())
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

	return s.operations.DoneLRO(ctx, "", op, func() *pb.BranchRule {
		return obj
	}())
}

func (s *secureSourceManagerServer) DeleteBranchRule(ctx context.Context, req *pb.DeleteBranchRuleRequest) (*longrunning.Operation, error) {
	name, err := s.parseBranchRuleName(req.GetName())
	if err != nil {
		return nil, err
	}

	fqn := name.String()
	now := time.Now()

	deleted := &pb.BranchRule{}
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

type BranchRuleName struct {
	Project      *projects.ProjectData
	Location     string
	RepositoryID string
	BranchRuleID string
}

func (n *BranchRuleName) String() string {
	return fmt.Sprintf("projects/%s/locations/%s/repositories/%s/branchRules/%s", n.Project.ID, n.Location, n.RepositoryID, n.BranchRuleID)
}

func (s *MockService) parseBranchRuleName(name string) (*BranchRuleName, error) {
	tokens := strings.Split(name, "/")

	if len(tokens) == 8 && tokens[0] == "projects" && tokens[2] == "locations" && tokens[4] == "repositories" && tokens[6] == "branchRules" {
		projectName, err := projects.ParseProjectName(tokens[0] + "/" + tokens[1])
		if err != nil {
			return nil, err
		}
		project, err := s.Projects.GetProject(projectName)
		if err != nil {
			return nil, err
		}

		name := &BranchRuleName{
			Project:      project,
			Location:     tokens[3],
			RepositoryID: tokens[5],
			BranchRuleID: tokens[7],
		}
		return name, nil
	} else {
		return nil, status.Errorf(codes.InvalidArgument, "name %q is not valid", name)
	}
}
