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

package mockcompute

import (
	"context"
	"fmt"
	"strings"

	"github.com/GoogleCloudPlatform/k8s-config-connector/mockgcp/common/projects"
	pb "github.com/GoogleCloudPlatform/k8s-config-connector/mockgcp/generated/mockgcp/cloud/compute/v1"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/proto"
)

type RegionNetworkEndpointGroupV1 struct {
	*MockService
	pb.UnimplementedRegionNetworkEndpointGroupsServer
}

func (s *RegionNetworkEndpointGroupV1) Get(ctx context.Context, req *pb.GetRegionNetworkEndpointGroupRequest) (*pb.NetworkEndpointGroup, error) {
	reqName := "projects/" + req.GetProject() + "/regions/" + req.GetRegion() + "/networkEndpointGroups/" + req.GetNetworkEndpointGroup()
	name, err := s.parseRegionNetworkEndpointGroupName(reqName)
	if err != nil {
		return nil, err
	}

	fqn := name.String()

	obj := &pb.NetworkEndpointGroup{}
	if err := s.storage.Get(ctx, fqn, obj); err != nil {
		return nil, status.Errorf(codes.NotFound, "The resource '%s' was not found", fqn)
	}

	return obj, nil
}

func (s *RegionNetworkEndpointGroupV1) Insert(ctx context.Context, req *pb.InsertRegionNetworkEndpointGroupRequest) (*pb.Operation, error) {
	reqName := "projects/" + req.GetProject() + "/regions/" + req.GetRegion() + "/networkEndpointGroups/" + req.GetNetworkEndpointGroupResource().GetName()
	name, err := s.parseRegionNetworkEndpointGroupName(reqName)
	if err != nil {
		return nil, err
	}

	fqn := name.String()

	id := s.generateID()

	obj := proto.Clone(req.GetNetworkEndpointGroupResource()).(*pb.NetworkEndpointGroup)
	obj.SelfLink = PtrTo(buildComputeSelfLink(ctx, fqn))
	obj.CreationTimestamp = PtrTo(s.nowString())
	obj.Id = &id
	obj.Kind = PtrTo("compute#networkEndpointGroup")
	obj.Region = PtrTo(buildComputeSelfLink(ctx, fmt.Sprintf("projects/%s/regions/%s", req.GetProject(), req.GetRegion())))
	obj.Size = PtrTo(int32(0))

	if err := s.storage.Create(ctx, fqn, obj); err != nil {
		return nil, err
	}

	op := &pb.Operation{
		TargetId:      obj.Id,
		TargetLink:    obj.SelfLink,
		OperationType: PtrTo("compute.regionNetworkEndpointGroups.insert"),
		User:          PtrTo("user@example.com"),
	}
	return s.startRegionalLRO(ctx, name.Project.ID, name.Region, op, func() (proto.Message, error) {
		return obj, nil
	})
}

func (s *RegionNetworkEndpointGroupV1) Delete(ctx context.Context, req *pb.DeleteRegionNetworkEndpointGroupRequest) (*pb.Operation, error) {
	reqName := "projects/" + req.GetProject() + "/regions/" + req.GetRegion() + "/networkEndpointGroups/" + req.GetNetworkEndpointGroup()
	name, err := s.parseRegionNetworkEndpointGroupName(reqName)
	if err != nil {
		return nil, err
	}

	fqn := name.String()

	deleted := &pb.NetworkEndpointGroup{}
	if err := s.storage.Delete(ctx, fqn, deleted); err != nil {
		return nil, err
	}

	op := &pb.Operation{
		TargetId:      deleted.Id,
		TargetLink:    deleted.SelfLink,
		OperationType: PtrTo("compute.regionNetworkEndpointGroups.delete"),
		User:          PtrTo("user@example.com"),
	}
	return s.startRegionalLRO(ctx, name.Project.ID, name.Region, op, func() (proto.Message, error) {
		return deleted, nil
	})
}

type regionNetworkEndpointGroupName struct {
	Project *projects.ProjectData
	Region  string
	Name    string
}

func (n *regionNetworkEndpointGroupName) String() string {
	return "projects/" + n.Project.ID + "/regions/" + n.Region + "/networkEndpointGroups/" + n.Name
}

// parseRegionNetworkEndpointGroupName parses a string into a regionNetworkEndpointGroupName.
// The expected form is `projects/*/regions/*/networkEndpointGroups/*`.
func (s *MockService) parseRegionNetworkEndpointGroupName(name string) (*regionNetworkEndpointGroupName, error) {
	tokens := strings.Split(name, "/")

	if len(tokens) == 6 && tokens[0] == "projects" && tokens[2] == "regions" && tokens[4] == "networkEndpointGroups" {
		project, err := s.Projects.GetProjectByID(tokens[1])
		if err != nil {
			return nil, err
		}
		name := &regionNetworkEndpointGroupName{
			Project: project,
			Region:  tokens[3],
			Name:    tokens[5],
		}

		return name, nil
	} else {
		return nil, status.Errorf(codes.InvalidArgument, "name %q is not valid", name)
	}
}
