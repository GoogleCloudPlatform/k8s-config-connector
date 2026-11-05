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

package mocknetworksecurity

import (
	"context"
	"fmt"
	"strings"
	"time"

	pbv1 "cloud.google.com/go/networksecurity/apiv1/networksecuritypb"
	"github.com/GoogleCloudPlatform/k8s-config-connector/mockgcp/common/projects"
	longrunningpb "google.golang.org/genproto/googleapis/longrunning"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/reflect/protoreflect"
	"google.golang.org/protobuf/types/known/emptypb"
	"google.golang.org/protobuf/types/known/timestamppb"
)

func (s *InterceptServer) CreateInterceptDeploymentGroup(ctx context.Context, req *pbv1.CreateInterceptDeploymentGroupRequest) (*longrunningpb.Operation, error) {
	reqName := req.Parent + "/interceptDeploymentGroups/" + req.InterceptDeploymentGroupId
	name, err := s.parseInterceptDeploymentGroupName(reqName)
	if err != nil {
		return nil, err
	}

	fqn := name.String()

	obj := proto.Clone(req.InterceptDeploymentGroup).(*pbv1.InterceptDeploymentGroup)
	obj.Name = fqn
	obj.CreateTime = timestamppb.New(time.Now())
	obj.UpdateTime = timestamppb.New(time.Now())
	obj.State = pbv1.InterceptDeploymentGroup_ACTIVE
	// Note: Real GCP returns "reconciling": false in the resource JSON;
	// proto3 JSON marshaling omits false boolean values by default.
	obj.Reconciling = false

	if err := s.storage.Create(ctx, fqn, obj); err != nil {
		if status.Code(err) == codes.AlreadyExists {
			return nil, status.Errorf(codes.AlreadyExists, "Resource '%s' already exists", fqn)
		}
		return nil, err
	}

	now := time.Now()
	lroMetadata := &pbv1.OperationMetadata{
		CreateTime:            timestamppb.New(now),
		RequestedCancellation: false,
		Target:                fqn,
		Verb:                  "create",
		ApiVersion:            "v1",
	}
	return s.operations.StartLRO(ctx, req.Parent, lroMetadata, func() (protoreflect.ProtoMessage, error) {
		lroMetadata.EndTime = timestamppb.New(time.Now())
		result := proto.Clone(obj).(*pbv1.InterceptDeploymentGroup)
		// Real GCP LRO response does not return labels on create
		result.Labels = nil
		return result, nil
	})
}

func (s *InterceptServer) GetInterceptDeploymentGroup(ctx context.Context, req *pbv1.GetInterceptDeploymentGroupRequest) (*pbv1.InterceptDeploymentGroup, error) {
	name, err := s.parseInterceptDeploymentGroupName(req.Name)
	if err != nil {
		return nil, err
	}

	fqn := name.String()

	obj := &pbv1.InterceptDeploymentGroup{}
	if err := s.storage.Get(ctx, fqn, obj); err != nil {
		if status.Code(err) == codes.NotFound {
			return nil, status.Errorf(codes.NotFound, "Resource '%s' was not found", fqn)
		}
		return nil, err
	}
	return obj, nil
}

func (s *InterceptServer) UpdateInterceptDeploymentGroup(ctx context.Context, req *pbv1.UpdateInterceptDeploymentGroupRequest) (*longrunningpb.Operation, error) {
	name, err := s.parseInterceptDeploymentGroupName(req.GetInterceptDeploymentGroup().GetName())
	if err != nil {
		return nil, err
	}

	fqn := name.String()

	obj := &pbv1.InterceptDeploymentGroup{}
	if err := s.storage.Get(ctx, fqn, obj); err != nil {
		if status.Code(err) == codes.NotFound {
			return nil, status.Errorf(codes.NotFound, "Resource '%s' was not found", fqn)
		}
		return nil, err
	}

	updated := proto.Clone(obj).(*pbv1.InterceptDeploymentGroup)
	updated.UpdateTime = timestamppb.New(time.Now())

	// Apply field mask updates
	paths := req.GetUpdateMask().GetPaths()
	if len(paths) == 0 {
		// Default to all updateable spec fields if empty
		paths = []string{"labels", "description"}
	}

	for _, path := range paths {
		switch path {
		case "labels":
			updated.Labels = req.GetInterceptDeploymentGroup().GetLabels()
		case "description":
			updated.Description = req.GetInterceptDeploymentGroup().GetDescription()
		default:
			return nil, status.Errorf(codes.InvalidArgument, "field %q is not updateable", path)
		}
	}

	if err := s.storage.Update(ctx, fqn, updated); err != nil {
		return nil, err
	}

	lroPrefix := fmt.Sprintf("projects/%s/locations/%s", name.Project.ID, name.Location)
	now := time.Now()
	lroMetadata := &pbv1.OperationMetadata{
		CreateTime: timestamppb.New(now),
		Target:     fqn,
		Verb:       "update",
		ApiVersion: "v1",
	}
	return s.operations.StartLRO(ctx, lroPrefix, lroMetadata, func() (protoreflect.ProtoMessage, error) {
		lroMetadata.EndTime = timestamppb.New(time.Now())
		result := proto.Clone(updated).(*pbv1.InterceptDeploymentGroup)
		// Real GCP LRO response does not return labels on update
		result.Labels = nil
		return result, nil
	})
}

func (s *InterceptServer) DeleteInterceptDeploymentGroup(ctx context.Context, req *pbv1.DeleteInterceptDeploymentGroupRequest) (*longrunningpb.Operation, error) {
	name, err := s.parseInterceptDeploymentGroupName(req.Name)
	if err != nil {
		return nil, err
	}

	fqn := name.String()

	obj := &pbv1.InterceptDeploymentGroup{}
	if err := s.storage.Delete(ctx, fqn, obj); err != nil {
		if status.Code(err) == codes.NotFound {
			return nil, status.Errorf(codes.NotFound, "Resource '%s' was not found", fqn)
		}
		return nil, err
	}

	lroPrefix := fmt.Sprintf("projects/%s/locations/%s", name.Project.ID, name.Location)
	now := time.Now()
	lroMetadata := &pbv1.OperationMetadata{
		CreateTime: timestamppb.New(now),
		Target:     fqn,
		Verb:       "delete",
		ApiVersion: "v1",
	}
	return s.operations.StartLRO(ctx, lroPrefix, lroMetadata, func() (protoreflect.ProtoMessage, error) {
		lroMetadata.EndTime = timestamppb.New(time.Now())
		return &emptypb.Empty{}, nil
	})
}

type interceptDeploymentGroupName struct {
	Project                  *projects.ProjectData
	Location                 string
	InterceptDeploymentGroup string
}

func (n *interceptDeploymentGroupName) String() string {
	return fmt.Sprintf("projects/%s/locations/%s/interceptDeploymentGroups/%s", n.Project.ID, n.Location, n.InterceptDeploymentGroup)
}

func (s *MockService) parseInterceptDeploymentGroupName(name string) (*interceptDeploymentGroupName, error) {
	tokens := strings.Split(name, "/")
	if len(tokens) == 6 && tokens[0] == "projects" && tokens[2] == "locations" && tokens[4] == "interceptDeploymentGroups" {
		project, err := s.Projects.GetProject(&projects.ProjectName{ProjectID: tokens[1]})
		if err != nil {
			return nil, err
		}
		return &interceptDeploymentGroupName{
			Project:                  project,
			Location:                 tokens[3],
			InterceptDeploymentGroup: tokens[5],
		}, nil
	}
	return nil, status.Errorf(codes.InvalidArgument, "name %q is not valid", name)
}
