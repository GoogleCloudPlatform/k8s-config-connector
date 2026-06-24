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
// proto.service: google.cloud.deploy.v1.CloudDeploy
// proto.message: google.cloud.deploy.v1.CustomTargetType

package mockclouddeploy

import (
	"context"
	"fmt"
	"strings"
	"time"

	longrunningpb "google.golang.org/genproto/googleapis/longrunning"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/types/known/emptypb"
	"google.golang.org/protobuf/types/known/timestamppb"

	pb "cloud.google.com/go/deploy/apiv1/deploypb"
	"github.com/GoogleCloudPlatform/k8s-config-connector/mockgcp/common/fields"
	"github.com/GoogleCloudPlatform/k8s-config-connector/mockgcp/common/projects"
	"github.com/GoogleCloudPlatform/k8s-config-connector/mockgcp/pkg/storage"
	"github.com/google/uuid"
)

func (s *cloudDeploy) GetCustomTargetType(ctx context.Context, req *pb.GetCustomTargetTypeRequest) (*pb.CustomTargetType, error) {
	name, err := s.parseCustomTargetTypeName(req.Name)
	if err != nil {
		return nil, err
	}

	fqn := name.String()

	obj := &pb.CustomTargetType{}
	if err := s.storage.Get(ctx, fqn, obj); err != nil {
		if status.Code(err) == codes.NotFound {
			return nil, status.Errorf(codes.NotFound, "Resource '%s' was not found", fqn)
		}
		return nil, err
	}

	return obj, nil
}

func (s *cloudDeploy) CreateCustomTargetType(ctx context.Context, req *pb.CreateCustomTargetTypeRequest) (*longrunningpb.Operation, error) {
	if req.CustomTargetTypeId == "" {
		return nil, status.Errorf(codes.InvalidArgument, "custom_target_type_id must be provided")
	}

	reqName := fmt.Sprintf("%s/customTargetTypes/%s", req.Parent, req.CustomTargetTypeId)
	name, err := s.parseCustomTargetTypeName(reqName)
	if err != nil {
		return nil, err
	}

	fqn := name.String()
	obj := proto.Clone(req.CustomTargetType).(*pb.CustomTargetType)
	obj.Name = fqn
	obj.CustomTargetTypeId = name.CustomTargetType

	obj.Uid = uuid.NewString()
	obj.CreateTime = timestamppb.New(time.Now())
	obj.UpdateTime = timestamppb.New(time.Now())
	obj.Etag = uuid.NewString()

	if err := s.storage.Create(ctx, fqn, obj); err != nil {
		return nil, err
	}

	// By default, immediately finish the LRO with success.
	lroPrefix := name.LocationPrefix()
	lroMetadata := &pb.OperationMetadata{
		CreateTime: timestamppb.New(time.Now()),
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
	if req.CustomTargetType == nil {
		return nil, status.Errorf(codes.InvalidArgument, "custom_target_type must be provided")
	}

	name, err := s.parseCustomTargetTypeName(req.CustomTargetType.Name)
	if err != nil {
		return nil, err
	}

	fqn := name.String()

	obj := &pb.CustomTargetType{}
	exists := true
	if err := s.storage.Get(ctx, fqn, obj); err != nil {
		if status.Code(err) == codes.NotFound && req.GetAllowMissing() {
			exists = false
			obj = proto.Clone(req.CustomTargetType).(*pb.CustomTargetType)
			obj.Name = fqn
			obj.CustomTargetTypeId = name.CustomTargetType
			obj.Uid = uuid.NewString()
			obj.CreateTime = timestamppb.New(time.Now())
			obj.UpdateTime = timestamppb.New(time.Now())
			obj.Etag = uuid.NewString()
		} else {
			if status.Code(err) == codes.NotFound {
				return nil, status.Errorf(codes.NotFound, "Resource '%s' was not found", fqn)
			}
			return nil, err
		}
	}

	if exists {
		// Apply the update mask to the object.
		paths := req.GetUpdateMask().GetPaths()
		if len(paths) == 0 {
			return nil, status.Errorf(codes.InvalidArgument, "update_mask must be provided")
		}

		if err := fields.UpdateByFieldMask(obj, req.CustomTargetType, req.UpdateMask.Paths); err != nil {
			return nil, fmt.Errorf("update field_mask.paths: %w", err)
		}

		obj.UpdateTime = timestamppb.New(time.Now())
		obj.Etag = uuid.NewString()

		if err := s.storage.Update(ctx, fqn, obj); err != nil {
			return nil, err
		}
	} else {
		if err := s.storage.Create(ctx, fqn, obj); err != nil {
			return nil, err
		}
	}

	lroPrefix := name.LocationPrefix()
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

func (s *cloudDeploy) ListCustomTargetTypes(ctx context.Context, req *pb.ListCustomTargetTypesRequest) (*pb.ListCustomTargetTypesResponse, error) {
	parent, err := s.parseLocationName(req.Parent)
	if err != nil {
		return nil, err
	}

	var objs []*pb.CustomTargetType
	kind := (&pb.CustomTargetType{}).ProtoReflect().Descriptor()
	err = s.storage.List(ctx, kind, storage.ListOptions{}, func(obj proto.Message) error {
		item := obj.(*pb.CustomTargetType)
		name, err := s.parseCustomTargetTypeName(item.Name)
		if err != nil {
			return nil
		}
		if name.Project.ID == parent.Project.ID && name.Location == parent.Location {
			objs = append(objs, item)
		}
		return nil
	})
	if err != nil {
		return nil, err
	}

	return &pb.ListCustomTargetTypesResponse{
		CustomTargetTypes: objs,
	}, nil
}

func (s *cloudDeploy) DeleteCustomTargetType(ctx context.Context, req *pb.DeleteCustomTargetTypeRequest) (*longrunningpb.Operation, error) {
	name, err := s.parseCustomTargetTypeName(req.Name)
	if err != nil {
		return nil, err
	}

	fqn := name.String()

	if err := s.storage.Delete(ctx, fqn, &pb.CustomTargetType{}); err != nil {
		if status.Code(err) == codes.NotFound && req.GetAllowMissing() {
			// Return success (LRO) if not found and AllowMissing is true
		} else {
			if status.Code(err) == codes.NotFound {
				return nil, status.Errorf(codes.NotFound, "Resource '%s' was not found", fqn)
			}
			return nil, err
		}
	}

	// By default, immediately finish the LRO with success.
	lroPrefix := name.LocationPrefix()
	lroMetadata := &pb.OperationMetadata{
		CreateTime: timestamppb.New(time.Now()),
		Target:     name.String(),
		Verb:       "delete",
		ApiVersion: "v1",
	}
	return s.operations.StartLRO(ctx, lroPrefix, lroMetadata, func() (proto.Message, error) {
		lroMetadata.EndTime = timestamppb.Now()
		return &emptypb.Empty{}, nil
	})
}

type customTargetTypeName struct {
	Project          *projects.ProjectData
	Location         string
	CustomTargetType string
}

func (n *customTargetTypeName) String() string {
	return fmt.Sprintf("projects/%s/locations/%s/customTargetTypes/%s", n.Project.ID, n.Location, n.CustomTargetType)
}

func (n *customTargetTypeName) LocationPrefix() string {
	return fmt.Sprintf("projects/%s/locations/%s", n.Project.ID, n.Location)
}

func (s *MockService) parseCustomTargetTypeName(name string) (*customTargetTypeName, error) {
	tokens := strings.Split(name, "/")

	if len(tokens) == 6 && tokens[0] == "projects" && tokens[2] == "locations" && tokens[4] == "customTargetTypes" {
		for i := 1; i < len(tokens); i += 2 {
			if tokens[i] == "" {
				return nil, status.Errorf(codes.InvalidArgument, "name %q is not valid", name)
			}
		}

		project, err := s.Projects.GetProjectByID(tokens[1])
		if err != nil {
			return nil, err
		}

		name := &customTargetTypeName{
			Project:          project,
			Location:         tokens[3],
			CustomTargetType: tokens[5],
		}

		return name, nil
	}

	return nil, status.Errorf(codes.InvalidArgument, "name %q is not valid", name)
}
