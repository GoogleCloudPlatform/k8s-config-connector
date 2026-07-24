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
	"google.golang.org/genproto/googleapis/longrunning"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/reflect/protoreflect"
	"google.golang.org/protobuf/types/known/emptypb"
	"google.golang.org/protobuf/types/known/timestamppb"
)

type SecurityProfileGroupServer struct {
	*MockService
	pbv1.UnimplementedSecurityProfileGroupServiceServer
}

func (s *SecurityProfileGroupServer) CreateSecurityProfileGroup(ctx context.Context, req *pbv1.CreateSecurityProfileGroupRequest) (*longrunning.Operation, error) {
	project, location, err := s.parseParent(req.Parent)
	if err != nil {
		return nil, err
	}

	name := req.Parent + "/securityProfileGroups/" + req.SecurityProfileGroupId
	fqn := name

	obj := proto.CloneOf(req.SecurityProfileGroup)
	obj.Name = fqn
	obj.CreateTime = timestamppb.New(time.Now())
	obj.UpdateTime = timestamppb.New(time.Now())
	obj.Etag = "dummy-etag"
	obj.DataPathId = 12345

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
	lroPrefix := fmt.Sprintf("projects/%s/locations/%s", project.ID, location)
	return s.operations.StartLRO(ctx, lroPrefix, lroMetadata, func() (protoreflect.ProtoMessage, error) {
		lroMetadata.EndTime = timestamppb.New(time.Now())
		result := proto.CloneOf(obj)
		return result, nil
	})
}

func (s *SecurityProfileGroupServer) GetSecurityProfileGroup(ctx context.Context, req *pbv1.GetSecurityProfileGroupRequest) (*pbv1.SecurityProfileGroup, error) {
	name, err := s.parseSecurityProfileGroupName(req.Name)
	if err != nil {
		return nil, err
	}

	fqn := name.String()

	obj := &pbv1.SecurityProfileGroup{}
	if err := s.storage.Get(ctx, fqn, obj); err != nil {
		if status.Code(err) == codes.NotFound {
			return nil, status.Errorf(codes.NotFound, "Resource '%s' was not found", fqn)
		}
		return nil, err
	}
	return obj, nil
}

func (s *SecurityProfileGroupServer) UpdateSecurityProfileGroup(ctx context.Context, req *pbv1.UpdateSecurityProfileGroupRequest) (*longrunning.Operation, error) {
	name, err := s.parseSecurityProfileGroupName(req.SecurityProfileGroup.Name)
	if err != nil {
		return nil, err
	}

	fqn := name.String()

	existing := &pbv1.SecurityProfileGroup{}
	if err := s.storage.Get(ctx, fqn, existing); err != nil {
		return nil, err
	}

	obj := proto.CloneOf(req.SecurityProfileGroup)
	obj.CreateTime = existing.CreateTime
	obj.UpdateTime = timestamppb.New(time.Now())
	obj.Etag = "updated-etag"
	obj.DataPathId = existing.DataPathId

	if err := s.storage.Update(ctx, fqn, obj); err != nil {
		return nil, err
	}

	now := time.Now()
	lroMetadata := &pbv1.OperationMetadata{
		CreateTime:            timestamppb.New(now),
		RequestedCancellation: false,
		Target:                fqn,
		Verb:                  "update",
		ApiVersion:            "v1",
	}
	lroPrefix := fmt.Sprintf("projects/%s/locations/%s", name.Project.ID, name.Location)
	return s.operations.StartLRO(ctx, lroPrefix, lroMetadata, func() (protoreflect.ProtoMessage, error) {
		lroMetadata.EndTime = timestamppb.New(time.Now())
		return obj, nil
	})
}

func (s *SecurityProfileGroupServer) DeleteSecurityProfileGroup(ctx context.Context, req *pbv1.DeleteSecurityProfileGroupRequest) (*longrunning.Operation, error) {
	name, err := s.parseSecurityProfileGroupName(req.Name)
	if err != nil {
		return nil, err
	}

	if err := s.storage.Delete(ctx, name.String(), &pbv1.SecurityProfileGroup{}); err != nil {
		return nil, err
	}
	now := time.Now()
	lroMetadata := &pbv1.OperationMetadata{
		CreateTime:            timestamppb.New(now),
		RequestedCancellation: false,
		Target:                name.String(),
		Verb:                  "delete",
		ApiVersion:            "v1",
	}
	lroPrefix := fmt.Sprintf("projects/%s/locations/%s", name.Project.ID, name.Location)
	return s.operations.StartLRO(ctx, lroPrefix, lroMetadata, func() (protoreflect.ProtoMessage, error) {
		lroMetadata.EndTime = timestamppb.New(time.Now())
		return &emptypb.Empty{}, nil
	})
}

type securityProfileGroupName struct {
	Project                *projects.ProjectData
	Location               string
	SecurityProfileGroupID string
}

func (n *securityProfileGroupName) String() string {
	return "projects/" + n.Project.ID + "/locations/" + n.Location + "/securityProfileGroups/" + n.SecurityProfileGroupID
}

func (s *SecurityProfileGroupServer) parseSecurityProfileGroupName(name string) (*securityProfileGroupName, error) {
	tokens := strings.Split(name, "/")
	if len(tokens) == 6 && tokens[0] == "projects" && tokens[2] == "locations" && tokens[4] == "securityProfileGroups" {
		project, err := s.Projects.GetProject(&projects.ProjectName{ProjectID: tokens[1]})
		if err != nil {
			return nil, err
		}
		name := &securityProfileGroupName{
			Project:                project,
			Location:               tokens[3],
			SecurityProfileGroupID: tokens[5],
		}
		return name, nil
	} else {
		return nil, status.Errorf(codes.InvalidArgument, "name %q is not valid", name)
	}
}

func (s *SecurityProfileGroupServer) parseParent(parent string) (*projects.ProjectData, string, error) {
	tokens := strings.Split(parent, "/")
	if len(tokens) == 4 && tokens[0] == "projects" && tokens[2] == "locations" {
		project, err := s.Projects.GetProject(&projects.ProjectName{ProjectID: tokens[1]})
		if err != nil {
			return nil, "", err
		}
		return project, tokens[3], nil
	}
	return nil, "", status.Errorf(codes.InvalidArgument, "parent %q is not valid", parent)
}
