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

package mocknetworksecurity

import (
	"context"
	"fmt"
	"strings"

	"google.golang.org/genproto/googleapis/longrunning"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/types/known/emptypb"
	"google.golang.org/protobuf/types/known/timestamppb"

	pb "github.com/GoogleCloudPlatform/k8s-config-connector/mockgcp/generated/mockgcp/cloud/networksecurity/v1"
)

type projectsLocationsAddressGroupsServer struct {
	*MockService
	pb.UnimplementedProjectsLocationsAddressGroupsServerServer
}

func (s *projectsLocationsAddressGroupsServer) GetProjectsLocationsAddressGroup(ctx context.Context, req *pb.GetProjectsLocationsAddressGroupRequest) (*pb.AddressGroup, error) {
	name, err := s.parseGroupName(req.GetName())
	if err != nil {
		return nil, err
	}

	fqn := name.String()

	obj := &pb.AddressGroup{}
	if err := s.storage.Get(ctx, fqn, obj); err != nil {
		if status.Code(err) == codes.NotFound {
			return nil, err
		}
		return nil, err
	}

	return obj, nil
}

func (s *projectsLocationsAddressGroupsServer) CreateProjectsLocationsAddressGroup(ctx context.Context, req *pb.CreateProjectsLocationsAddressGroupRequest) (*longrunning.Operation, error) {
	reqName := req.GetParent() + "/addressGroups/" + req.GetAddressGroupId()
	name, err := s.parseGroupName(reqName)
	if err != nil {
		return nil, err
	}

	fqn := name.String()

	obj := proto.Clone(req.ProjectsLocationsAddressGroup).(*pb.AddressGroup)

	if obj.Purpose == nil {
		obj.Purpose = []string{"DEFAULT"}
	}

	if err := s.storage.Create(ctx, fqn, obj); err != nil {
		return nil, err
	}

	metadata := &pb.OperationMetadata{
		ApiVersion:            PtrTo("v1"),
		CreateTime:            timestamppb.Now(),
		RequestedCancellation: PtrTo(false),
		Target:                PtrTo(fqn),
		Verb:                  PtrTo("create"),
	}

	prefix := fmt.Sprintf("projects/%s/locations/%s", name.Project, name.Location)
	return s.operations.StartLRO(ctx, prefix, metadata, func() (proto.Message, error) {
		metadata.EndTime = timestamppb.Now()

		result := proto.Clone(obj).(*pb.AddressGroup)
		result.CreateTime = timestamppb.Now()
		result.UpdateTime = timestamppb.Now()
		result.Name = PtrTo(fqn)

		if err := s.storage.Update(ctx, fqn, result); err != nil {
			return nil, err
		}

		return result, nil
	})
}

func (s *projectsLocationsAddressGroupsServer) PatchProjectsLocationsAddressGroup(ctx context.Context, req *pb.PatchProjectsLocationsAddressGroupRequest) (*longrunning.Operation, error) {
	reqName := req.GetName()

	name, err := s.parseGroupName(reqName)
	if err != nil {
		return nil, err
	}

	fqn := name.String()
	obj := &pb.AddressGroup{}
	if err := s.storage.Get(ctx, fqn, obj); err != nil {
		return nil, err
	}

	// TODO: Some sort of helper for fieldmask?
	for _, path := range strings.Split(req.GetUpdateMask(), ",") {
		switch path {
		case "description":
			obj.Description = req.GetProjectsLocationsAddressGroup().Description
		default:
			return nil, status.Errorf(codes.InvalidArgument, "update_mask path %q not valid", path)
		}
	}

	obj.UpdateTime = timestamppb.Now()

	if err := s.storage.Update(ctx, fqn, obj); err != nil {
		return nil, err
	}

	metadata := &pb.OperationMetadata{
		ApiVersion:            PtrTo("v1"),
		CreateTime:            timestamppb.Now(),
		Target:                PtrTo(fqn),
		RequestedCancellation: PtrTo(false),
		Verb:                  PtrTo("update"),
	}

	prefix := fmt.Sprintf("projects/%s/locations/%s", name.Project, name.Location)
	return s.operations.StartLRO(ctx, prefix, metadata, func() (proto.Message, error) {
		metadata.EndTime = timestamppb.Now()

		result := proto.Clone(obj).(*pb.AddressGroup)
		result.CreateTime = timestamppb.Now()
		result.UpdateTime = timestamppb.Now()
		if err := s.storage.Update(ctx, fqn, result); err != nil {
			return nil, err
		}

		return result, nil
	})
}

func (s *projectsLocationsAddressGroupsServer) DeleteProjectsLocationsAddressGroup(ctx context.Context, req *pb.DeleteProjectsLocationsAddressGroupRequest) (*longrunning.Operation, error) {
	name, err := s.parseGroupName(req.GetName())
	if err != nil {
		return nil, err
	}

	fqn := name.String()

	deleted := &pb.AddressGroup{}
	if err := s.storage.Delete(ctx, fqn, deleted); err != nil {
		return nil, err
	}

	metadata := &pb.OperationMetadata{
		ApiVersion:            PtrTo("v1"),
		CreateTime:            timestamppb.Now(),
		Target:                PtrTo(fqn),
		RequestedCancellation: PtrTo(false),
		Verb:                  PtrTo("delete"),
	}

	prefix := fmt.Sprintf("projects/%s/locations/%s", name.Project, name.Location)
	return s.operations.StartLRO(ctx, prefix, metadata, func() (proto.Message, error) {
		metadata.EndTime = timestamppb.Now()
		return &emptypb.Empty{}, nil
	})
}

type groupName struct {
	Project  string
	Location string
	Name     string
}

func (n *groupName) String() string {
	return "projects/" + n.Project + "/locations/" + n.Location + "/addressGroups/" + n.Name
}

func (s *MockService) parseGroupName(name string) (*groupName, error) {
	tokens := strings.Split(name, "/")
	if len(tokens) == 6 && tokens[0] == "projects" && tokens[2] == "locations" && tokens[4] == "addressGroups" {
		name := &groupName{
			Project:  tokens[1],
			Location: tokens[3],
			Name:     tokens[5],
		}
		return name, nil
	} else {
		return nil, status.Errorf(codes.InvalidArgument, "name %q is not valid", name)
	}
}
