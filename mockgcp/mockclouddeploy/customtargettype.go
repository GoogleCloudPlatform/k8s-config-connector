// Copyright 2025 Google LLC
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
// proto.service: google.cloud.deploy.v1.CloudDeploy
// proto.message: google.cloud.deploy.v1.CustomTargetType

package mockclouddeploy

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

	"github.com/GoogleCloudPlatform/k8s-config-connector/mockgcp/common/fields"
	"github.com/GoogleCloudPlatform/k8s-config-connector/mockgcp/common/projects"
	pb "github.com/GoogleCloudPlatform/k8s-config-connector/mockgcp/generated/mockgcp/cloud/deploy/v1"
	longrunningpb "google.golang.org/genproto/googleapis/longrunning"
)

func (s *cloudDeploy) GetCustomTargetType(ctx context.Context, req *pb.GetCustomTargetTypeRequest) (*pb.CustomTargetType, error) {
	name, err := s.parseCustomTargetTypeName(req.Name)
	if err != nil {
		return nil, err
	}

	fqn := name.String()

	obj := &pb.CustomTargetType{}
	if err := s.storage.Get(ctx, fqn, obj); err != nil {
		return nil, err
	}

	return obj, nil
}

func (s *cloudDeploy) CreateCustomTargetType(ctx context.Context, req *pb.CreateCustomTargetTypeRequest) (*longrunningpb.Operation, error) {
	reqName := fmt.Sprintf("%s/customTargetTypes/%s", req.GetParent(), req.GetCustomTargetTypeId())

	name, err := s.parseCustomTargetTypeName(reqName)
	if err != nil {
		return nil, err
	}

	fqn := name.String()

	obj := proto.Clone(req.CustomTargetType).(*pb.CustomTargetType)

	now := time.Now()
	obj.Name = fqn
	obj.CreateTime = timestamppb.New(now)
	obj.UpdateTime = timestamppb.New(now)
	obj.Uid = "a-uuid" // TODO: Fix this
	obj.Etag = "etag-1"

	if err := s.storage.Create(ctx, fqn, obj); err != nil {
		return nil, err
	}

	// By default, immediately finish the LRO with success.
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

func (s *cloudDeploy) UpdateCustomTargetType(ctx context.Context, req *pb.UpdateCustomTargetTypeRequest) (*longrunningpb.Operation, error) {
	name, err := s.parseCustomTargetTypeName(req.GetCustomTargetType().GetName())
	if err != nil {
		return nil, err
	}
	fqn := name.String()
	obj := &pb.CustomTargetType{}

	if err := s.storage.Get(ctx, fqn, obj); err != nil {
		return nil, err
	}

	// Apply the update mask to the object.
	paths := req.GetUpdateMask().GetPaths()
	if len(paths) == 0 {
		return nil, status.Errorf(codes.InvalidArgument, "update_mask must be provided")
	}
	if err := fields.UpdateByFieldMask(obj, req.GetCustomTargetType(), req.UpdateMask.Paths); err != nil {
		return nil, fmt.Errorf("update field_mask.paths: %w", err)
	}

	obj.UpdateTime = timestamppb.New(time.Now())
	if err := s.storage.Update(ctx, fqn, obj); err != nil {
		return nil, err
	}

	// By default, immediately finish the LRO with success.
	lroPrefix := fmt.Sprintf("projects/%s/locations/%s", name.Project.ID, name.Location)
	lroMetadata := &pb.OperationMetadata{
		CreateTime: timestamppb.New(time.Now()),
		Target:     name.String(),
		Verb:       "update",
		ApiVersion: "v1",
	}
	return s.operations.StartLRO(ctx, lroPrefix, lroMetadata, func() (proto.Message, error) {
		lroMetadata.EndTime = timestamppb.Now()
		return obj, nil
	})
}

func (s *cloudDeploy) DeleteCustomTargetType(ctx context.Context, req *pb.DeleteCustomTargetTypeRequest) (*longrunningpb.Operation, error) {
	name, err := s.parseCustomTargetTypeName(req.Name)
	if err != nil {
		return nil, err
	}

	fqn := name.String()

	deleted := &pb.CustomTargetType{}
	if err := s.storage.Delete(ctx, fqn, deleted); err != nil {
		return nil, err
	}

	// By default, immediately finish the LRO with success.
	lroPrefix := fmt.Sprintf("projects/%s/locations/%s", name.Project.ID, name.Location)
	lroMetadata := &pb.OperationMetadata{
		CreateTime: timestamppb.New(time.Now()),
		Target:     name.String(),
		Verb:       "delete",
		ApiVersion: "v1",
	}
	return s.operations.StartLRO(ctx, lroPrefix, lroMetadata, func() (proto.Message, error) {
		lroMetadata.EndTime = timestamppb.New(time.Now())
		return &emptypb.Empty{}, nil
	})
}

type customTargetTypeName struct {
	Project            *projects.ProjectData
	Location           string
	CustomTargetTypeID string
}

func (n *customTargetTypeName) String() string {
	return "projects/" + n.Project.ID + "/locations/" + n.Location + "/customTargetTypes/" + n.CustomTargetTypeID
}

// parseCustomTargetTypeName parses a string into a customTargetTypeName.
// The expected form is `projects/*/locations/*/customTargetTypes/*`.
func (s *MockService) parseCustomTargetTypeName(name string) (*customTargetTypeName, error) {
	tokens := strings.Split(name, "/")
	if len(tokens) == 6 && tokens[0] == "projects" && tokens[2] == "locations" && tokens[4] == "customTargetTypes" {
		project, err := s.Projects.GetProjectByID(tokens[1])
		if err != nil {
			return nil, err
		}

		name := &customTargetTypeName{
			Project:            project,
			Location:           tokens[3],
			CustomTargetTypeID: tokens[5],
		}

		return name, nil
	}

	return nil, status.Errorf(codes.InvalidArgument, "name %q is not valid", name)
}
