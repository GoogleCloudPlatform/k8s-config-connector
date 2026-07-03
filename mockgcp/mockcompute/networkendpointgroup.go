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

type NetworkEndpointGroupV1 struct {
	*MockService
	pb.UnimplementedNetworkEndpointGroupsServer
}

func (s *NetworkEndpointGroupV1) Get(ctx context.Context, req *pb.GetNetworkEndpointGroupRequest) (*pb.NetworkEndpointGroup, error) {
	reqName := "projects/" + req.GetProject() + "/zones/" + req.GetZone() + "/networkEndpointGroups/" + req.GetNetworkEndpointGroup()
	name, err := s.parseNetworkEndpointGroupName(reqName)
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

func (s *NetworkEndpointGroupV1) Insert(ctx context.Context, req *pb.InsertNetworkEndpointGroupRequest) (*pb.Operation, error) {
	reqName := "projects/" + req.GetProject() + "/zones/" + req.GetZone() + "/networkEndpointGroups/" + req.GetNetworkEndpointGroupResource().GetName()
	name, err := s.parseNetworkEndpointGroupName(reqName)
	if err != nil {
		return nil, err
	}

	fqn := name.String()

	id := s.generateID()

	obj := proto.CloneOf(req.GetNetworkEndpointGroupResource())
	obj.SelfLink = PtrTo(BuildComputeSelfLink(ctx, fqn))
	obj.CreationTimestamp = PtrTo(s.nowString())
	obj.Id = &id
	obj.Kind = PtrTo("compute#networkEndpointGroup")
	obj.Zone = PtrTo(BuildComputeSelfLink(ctx, fmt.Sprintf("projects/%s/zones/%s", req.GetProject(), req.GetZone())))
	obj.Size = PtrTo(int32(0))

	if err := s.storage.Create(ctx, fqn, obj); err != nil {
		return nil, err
	}

	op := &pb.Operation{
		TargetId:      obj.Id,
		TargetLink:    obj.SelfLink,
		OperationType: PtrTo("compute.networkEndpointGroups.insert"),
		User:          PtrTo("user@example.com"),
	}
	return s.startZonalLRO(ctx, name.Project.ID, name.Zone, op, func() (proto.Message, error) {
		return obj, nil
	})
}

func (s *NetworkEndpointGroupV1) Delete(ctx context.Context, req *pb.DeleteNetworkEndpointGroupRequest) (*pb.Operation, error) {
	reqName := "projects/" + req.GetProject() + "/zones/" + req.GetZone() + "/networkEndpointGroups/" + req.GetNetworkEndpointGroup()
	name, err := s.parseNetworkEndpointGroupName(reqName)
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
		OperationType: PtrTo("compute.networkEndpointGroups.delete"),
		User:          PtrTo("user@example.com"),
	}
	return s.startZonalLRO(ctx, name.Project.ID, name.Zone, op, func() (proto.Message, error) {
		return deleted, nil
	})
}

type networkEndpointGroupName struct {
	Project *projects.ProjectData
	Zone    string
	Name    string
}

func (n *networkEndpointGroupName) String() string {
	return "projects/" + n.Project.ID + "/zones/" + n.Zone + "/networkEndpointGroups/" + n.Name
}

// parseNetworkEndpointGroupName parses a string into a networkEndpointGroupName.
// The expected form is `projects/*/zones/*/networkEndpointGroups/*`.
func (s *MockService) parseNetworkEndpointGroupName(name string) (*networkEndpointGroupName, error) {
	tokens := strings.Split(name, "/")

	if len(tokens) == 6 && tokens[0] == "projects" && tokens[2] == "zones" && tokens[4] == "networkEndpointGroups" {
		project, err := s.Projects.GetProjectByID(tokens[1])
		if err != nil {
			return nil, err
		}
		name := &networkEndpointGroupName{
			Project: project,
			Zone:    tokens[3],
			Name:    tokens[5],
		}

		return name, nil
	} else {
		return nil, status.Errorf(codes.InvalidArgument, "name %q is not valid", name)
	}
}
