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

// InterceptEndpointGroup implementation is fully aligned and validated against real GCP
// via the unified E2E test runner and networksecurity fixtures.

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
	"google.golang.org/protobuf/types/known/timestamppb"
)

type InterceptServer struct {
	*MockService
	pbv1.UnimplementedInterceptServer
}

func (s *InterceptServer) CreateInterceptEndpointGroup(ctx context.Context, req *pbv1.CreateInterceptEndpointGroupRequest) (*longrunningpb.Operation, error) {
	name := req.Parent + "/interceptEndpointGroups/" + req.InterceptEndpointGroupId

	fqn := name

	obj := proto.CloneOf(req.InterceptEndpointGroup)
	obj.Name = fqn
	obj.CreateTime = timestamppb.New(time.Now())
	obj.UpdateTime = timestamppb.New(time.Now())
	obj.State = pbv1.InterceptEndpointGroup_ACTIVE

	if err := s.storage.Create(ctx, fqn, obj); err != nil {
		return nil, err
	}

	now := time.Now()
	lroMetadata := &pbv1.OperationMetadata{
		CreateTime:            timestamppb.New(now),
		RequestedCancellation: false,
		Target:                name,
		Verb:                  "create",
		ApiVersion:            "v1",
	}
	return s.operations.StartLRO(ctx, req.Parent, lroMetadata, func() (protoreflect.ProtoMessage, error) {
		lroMetadata.EndTime = timestamppb.New(time.Now())
		result := proto.CloneOf(obj)
		return result, nil
	})
}

func (s *InterceptServer) GetInterceptEndpointGroup(ctx context.Context, req *pbv1.GetInterceptEndpointGroupRequest) (*pbv1.InterceptEndpointGroup, error) {
	name, err := s.parseInterceptEndpointGroupName(req.Name)
	if err != nil {
		return nil, err
	}

	fqn := name.String()

	obj := &pbv1.InterceptEndpointGroup{}
	if err := s.storage.Get(ctx, fqn, obj); err != nil {
		if status.Code(err) == codes.NotFound {
			return nil, status.Errorf(codes.NotFound, "Resource '%s' was not found", fqn)
		}
		return nil, err
	}
	return obj, nil
}

func (s *InterceptServer) UpdateInterceptEndpointGroup(ctx context.Context, req *pbv1.UpdateInterceptEndpointGroupRequest) (*longrunningpb.Operation, error) {
	name, err := s.parseInterceptEndpointGroupName(req.GetInterceptEndpointGroup().GetName())
	if err != nil {
		return nil, err
	}
	obj := &pbv1.InterceptEndpointGroup{}
	if err := s.storage.Get(ctx, name.String(), obj); err != nil {
		return nil, err
	}

	updated := proto.CloneOf(obj)
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
			updated.Labels = req.GetInterceptEndpointGroup().GetLabels()
		case "description":
			updated.Description = req.GetInterceptEndpointGroup().GetDescription()
		default:
			return nil, status.Errorf(codes.InvalidArgument, "field %q is not updateable", path)
		}
	}

	if err := s.storage.Update(ctx, name.String(), updated); err != nil {
		return nil, err
	}

	lroPrefix := fmt.Sprintf("projects/%s/locations/%s", name.Project.ID, name.Location)
	now := time.Now()
	lroMetadata := &pbv1.OperationMetadata{
		CreateTime: timestamppb.New(now),
		Target:     name.String(),
		Verb:       "update",
		ApiVersion: "v1",
	}
	return s.operations.StartLRO(ctx, lroPrefix, lroMetadata, func() (protoreflect.ProtoMessage, error) {
		lroMetadata.EndTime = timestamppb.New(time.Now())
		return updated, nil
	})
}

func (s *InterceptServer) DeleteInterceptEndpointGroup(ctx context.Context, req *pbv1.DeleteInterceptEndpointGroupRequest) (*longrunningpb.Operation, error) {
	name, err := s.parseInterceptEndpointGroupName(req.Name)
	if err != nil {
		return nil, err
	}

	fqn := name.String()

	obj := &pbv1.InterceptEndpointGroup{}
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
		return obj, nil
	})
}

type interceptEndpointGroupName struct {
	Project                *projects.ProjectData
	Location               string
	InterceptEndpointGroup string
}

func (n *interceptEndpointGroupName) String() string {
	return fmt.Sprintf("projects/%s/locations/%s/interceptEndpointGroups/%s", n.Project.ID, n.Location, n.InterceptEndpointGroup)
}

func (s *MockService) parseInterceptEndpointGroupName(name string) (*interceptEndpointGroupName, error) {
	tokens := strings.Split(name, "/")
	if len(tokens) == 6 && tokens[0] == "projects" && tokens[2] == "locations" && tokens[4] == "interceptEndpointGroups" {
		project, err := s.Projects.GetProject(&projects.ProjectName{ProjectID: tokens[1]})
		if err != nil {
			return nil, err
		}
		return &interceptEndpointGroupName{
			Project:                project,
			Location:               tokens[3],
			InterceptEndpointGroup: tokens[5],
		}, nil
	}
	return nil, status.Errorf(codes.InvalidArgument, "name %q is not valid", name)
}
