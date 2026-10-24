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
// proto.service: google.cloud.netapp.v1.NetApp
// proto.message: google.cloud.netapp.v1.HostGroup

package mocknetapp

import (
	"context"
	"fmt"
	"strings"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/types/known/emptypb"
	"google.golang.org/protobuf/types/known/timestamppb"

	"cloud.google.com/go/longrunning/autogen/longrunningpb"
	pb "cloud.google.com/go/netapp/apiv1/netapppb"
	"github.com/GoogleCloudPlatform/k8s-config-connector/mockgcp/common/fields"
	"github.com/GoogleCloudPlatform/k8s-config-connector/mockgcp/common/projects"
	"github.com/GoogleCloudPlatform/k8s-config-connector/mockgcp/pkg/storage"
)

func (s *backupVaultsService) GetHostGroup(ctx context.Context, req *pb.GetHostGroupRequest) (*pb.HostGroup, error) {
	name, err := s.parseHostGroupName(req.Name)
	if err != nil {
		return nil, err
	}

	fqn := name.String()

	obj := &pb.HostGroup{}
	if err := s.storage.Get(ctx, fqn, obj); err != nil {
		if status.Code(err) == codes.NotFound {
			return nil, status.Errorf(codes.NotFound, "Resource '%v' was not found", fqn)
		}
		return nil, err
	}

	return obj, nil
}

func (s *backupVaultsService) ListHostGroups(ctx context.Context, req *pb.ListHostGroupsRequest) (*pb.ListHostGroupsResponse, error) {
	parent := req.GetParent()
	response := &pb.ListHostGroupsResponse{}

	findPrefix := fmt.Sprintf("%s/hostGroups/", parent)

	hostGroupKind := (&pb.HostGroup{}).ProtoReflect().Descriptor()
	if err := s.storage.List(ctx, hostGroupKind, storage.ListOptions{Prefix: findPrefix}, func(obj proto.Message) error {
		cloned := proto.Clone(obj).(*pb.HostGroup)
		response.HostGroups = append(response.HostGroups, cloned)
		return nil
	}); err != nil {
		return nil, err
	}

	return response, nil
}

func (s *backupVaultsService) CreateHostGroup(ctx context.Context, req *pb.CreateHostGroupRequest) (*longrunningpb.Operation, error) {
	reqName := fmt.Sprintf("%s/hostGroups/%s", req.GetParent(), req.GetHostGroupId())
	name, err := s.parseHostGroupName(reqName)
	if err != nil {
		return nil, err
	}

	fqn := name.String()

	obj := proto.Clone(req.GetHostGroup()).(*pb.HostGroup)
	obj.Name = fqn
	obj.State = pb.HostGroup_READY

	if err := s.storage.Create(ctx, fqn, obj); err != nil {
		return nil, err
	}

	lroPrefix := fmt.Sprintf("projects/%s/locations/%s", name.Project.ID, name.Location)
	lroMetadata := &pb.OperationMetadata{
		CreateTime:            timestamppb.Now(),
		Target:                name.String(),
		Verb:                  "create",
		ApiVersion:            "v1",
		RequestedCancellation: false,
	}
	return s.operations.StartLRO(ctx, lroPrefix, lroMetadata, func() (proto.Message, error) {
		lroMetadata.EndTime = timestamppb.Now()
		return obj, nil
	})
}

func (s *backupVaultsService) UpdateHostGroup(ctx context.Context, req *pb.UpdateHostGroupRequest) (*longrunningpb.Operation, error) {
	name, err := s.parseHostGroupName(req.GetHostGroup().GetName())
	if err != nil {
		return nil, err
	}
	fqn := name.String()

	obj := &pb.HostGroup{}
	if err := s.storage.Get(ctx, fqn, obj); err != nil {
		if status.Code(err) == codes.NotFound {
			return nil, status.Errorf(codes.NotFound, "Resource '%v' was not found", fqn)
		}
		return nil, err
	}

	paths := req.GetUpdateMask().GetPaths()
	if len(paths) == 0 {
		return nil, status.Errorf(codes.InvalidArgument, "update_mask must be provided")
	}
	if err := fields.UpdateByFieldMask(obj, req.GetHostGroup(), req.UpdateMask.Paths); err != nil {
		return nil, fmt.Errorf("update field_mask.paths: %w", err)
	}

	if err := s.storage.Update(ctx, fqn, obj); err != nil {
		return nil, err
	}

	prefix := fmt.Sprintf("projects/%s/locations/%s", name.Project.ID, name.Location)
	lroMetadata := &pb.OperationMetadata{
		CreateTime:            timestamppb.Now(),
		Target:                name.String(),
		Verb:                  "update",
		ApiVersion:            "v1",
		RequestedCancellation: false,
	}
	return s.operations.StartLRO(ctx, prefix, lroMetadata, func() (proto.Message, error) {
		lroMetadata.EndTime = timestamppb.Now()
		return obj, nil
	})
}

func (s *backupVaultsService) DeleteHostGroup(ctx context.Context, req *pb.DeleteHostGroupRequest) (*longrunningpb.Operation, error) {
	name, err := s.parseHostGroupName(req.Name)
	if err != nil {
		return nil, err
	}

	fqn := name.String()

	deleted := &pb.HostGroup{}
	if err := s.storage.Delete(ctx, fqn, deleted); err != nil {
		return nil, err
	}

	prefix := fmt.Sprintf("projects/%s/locations/%s", name.Project.ID, name.Location)
	lroMetadata := &pb.OperationMetadata{
		CreateTime:            timestamppb.Now(),
		Target:                fqn,
		Verb:                  "delete",
		ApiVersion:            "v1",
		RequestedCancellation: false,
	}
	return s.operations.StartLRO(ctx, prefix, lroMetadata, func() (proto.Message, error) {
		lroMetadata.EndTime = timestamppb.Now()
		return &emptypb.Empty{}, nil
	})
}

type hostGroupName struct {
	Project     *projects.ProjectData
	Location    string
	HostGroupID string
}

func (n *hostGroupName) String() string {
	return fmt.Sprintf("projects/%s/locations/%s/hostGroups/%s", n.Project.ID, n.Location, n.HostGroupID)
}

// parseHostGroupName parses a string into a hostGroupName.
// The expected form is `projects/*/locations/*/hostGroups/*`.
func (s *MockService) parseHostGroupName(name string) (*hostGroupName, error) {
	tokens := strings.Split(name, "/")

	if len(tokens) == 6 && tokens[0] == "projects" && tokens[2] == "locations" && tokens[4] == "hostGroups" {
		project, err := s.Projects.GetProjectByID(tokens[1])
		if err != nil {
			return nil, err
		}

		name := &hostGroupName{
			Project:     project,
			Location:    tokens[3],
			HostGroupID: tokens[5],
		}

		return name, nil
	}

	return nil, status.Errorf(codes.InvalidArgument, "name %q is not valid", name)
}
