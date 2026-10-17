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
	"github.com/GoogleCloudPlatform/k8s-config-connector/mockgcp/pkg/storage"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/proto"
)

type NetworkEndpointGroupsV1 struct {
	*MockService
	pb.UnimplementedNetworkEndpointGroupsServer
}

func (s *NetworkEndpointGroupsV1) Get(ctx context.Context, req *pb.GetNetworkEndpointGroupRequest) (*pb.NetworkEndpointGroup, error) {
	reqName := "projects/" + req.GetProject() + "/zones/" + req.GetZone() + "/networkEndpointGroups/" + req.GetNetworkEndpointGroup()
	name, err := s.parseZonalNetworkEndpointGroupName(reqName)
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

func (s *NetworkEndpointGroupsV1) Insert(ctx context.Context, req *pb.InsertNetworkEndpointGroupRequest) (*pb.Operation, error) {
	reqName := "projects/" + req.GetProject() + "/zones/" + req.GetZone() + "/networkEndpointGroups/" + req.GetNetworkEndpointGroupResource().GetName()
	name, err := s.parseZonalNetworkEndpointGroupName(reqName)
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

	if obj.Network != nil {
		obj.Network = PtrTo(ExpandComputeLink(ctx, *obj.Network))
	}
	if obj.Subnetwork != nil {
		obj.Subnetwork = PtrTo(ExpandComputeLink(ctx, *obj.Subnetwork))
	}

	if err := s.storage.Create(ctx, fqn, obj); err != nil {
		return nil, err
	}

	op := &pb.Operation{
		TargetId:      obj.Id,
		TargetLink:    obj.SelfLink,
		OperationType: PtrTo("insert"),
		User:          PtrTo("user@example.com"),
	}
	return s.startZonalLRO(ctx, name.Project.ID, name.Zone, op, func() (proto.Message, error) {
		return obj, nil
	})
}

func (s *NetworkEndpointGroupsV1) Delete(ctx context.Context, req *pb.DeleteNetworkEndpointGroupRequest) (*pb.Operation, error) {
	reqName := "projects/" + req.GetProject() + "/zones/" + req.GetZone() + "/networkEndpointGroups/" + req.GetNetworkEndpointGroup()
	name, err := s.parseZonalNetworkEndpointGroupName(reqName)
	if err != nil {
		return nil, err
	}

	fqn := name.String()

	deleted := &pb.NetworkEndpointGroup{}
	if err := s.storage.Delete(ctx, fqn, deleted); err != nil {
		return nil, err
	}

	// Also delete all associated network endpoints
	prefix := fqn + "/networkEndpoints/"
	var toDelete []string
	kind := (&pb.NetworkEndpoint{}).ProtoReflect().Descriptor()
	err = s.storage.List(ctx, kind, storage.ListOptions{Prefix: prefix}, func(obj proto.Message) error {
		endpoint := obj.(*pb.NetworkEndpoint)
		instanceName := ""
		if endpoint.Instance != nil {
			instanceName = lastComponent(*endpoint.Instance)
		}
		ipAddress := endpoint.GetIpAddress()
		port := endpoint.GetPort()
		endpointFQN := fmt.Sprintf("%s%s-%d-%s", prefix, ipAddress, port, instanceName)
		toDelete = append(toDelete, endpointFQN)
		return nil
	})
	if err == nil {
		for _, endpointFQN := range toDelete {
			_ = s.storage.Delete(ctx, endpointFQN, &pb.NetworkEndpoint{})
		}
	}

	op := &pb.Operation{
		TargetId:      deleted.Id,
		TargetLink:    deleted.SelfLink,
		OperationType: PtrTo("delete"),
		User:          PtrTo("user@example.com"),
	}
	return s.startZonalLRO(ctx, name.Project.ID, name.Zone, op, func() (proto.Message, error) {
		return deleted, nil
	})
}

func (s *NetworkEndpointGroupsV1) AttachNetworkEndpoints(ctx context.Context, req *pb.AttachNetworkEndpointsNetworkEndpointGroupRequest) (*pb.Operation, error) {
	reqName := "projects/" + req.GetProject() + "/zones/" + req.GetZone() + "/networkEndpointGroups/" + req.GetNetworkEndpointGroup()
	name, err := s.parseZonalNetworkEndpointGroupName(reqName)
	if err != nil {
		return nil, err
	}

	fqn := name.String()

	neg := &pb.NetworkEndpointGroup{}
	if err := s.storage.Get(ctx, fqn, neg); err != nil {
		return nil, status.Errorf(codes.NotFound, "The resource '%s' was not found", fqn)
	}

	endpoints := req.GetNetworkEndpointGroupsAttachEndpointsRequestResource().GetNetworkEndpoints()
	for _, endpoint := range endpoints {
		instanceName := ""
		if endpoint.Instance != nil {
			instanceName = lastComponent(*endpoint.Instance)
		}
		ipAddress := endpoint.GetIpAddress()
		port := endpoint.GetPort()
		endpointFQN := fmt.Sprintf("%s/networkEndpoints/%s-%d-%s", fqn, ipAddress, port, instanceName)

		existing := &pb.NetworkEndpoint{}
		if err := s.storage.Get(ctx, endpointFQN, existing); err != nil {
			if status.Code(err) == codes.NotFound {
				if err := s.storage.Create(ctx, endpointFQN, endpoint); err != nil {
					return nil, err
				}
			} else {
				return nil, err
			}
		} else {
			if err := s.storage.Update(ctx, endpointFQN, endpoint); err != nil {
				return nil, err
			}
		}
	}

	var count int32
	kind := (&pb.NetworkEndpoint{}).ProtoReflect().Descriptor()
	prefix := fqn + "/networkEndpoints/"
	err = s.storage.List(ctx, kind, storage.ListOptions{Prefix: prefix}, func(obj proto.Message) error {
		count++
		return nil
	})
	if err == nil {
		neg.Size = &count
		_ = s.storage.Update(ctx, fqn, neg)
	}

	op := &pb.Operation{
		TargetId:      neg.Id,
		TargetLink:    neg.SelfLink,
		OperationType: PtrTo("AttachNetworkEndpoints"),
		User:          PtrTo("user@example.com"),
	}
	return s.startZonalLRO(ctx, name.Project.ID, name.Zone, op, func() (proto.Message, error) {
		return neg, nil
	})
}

func (s *NetworkEndpointGroupsV1) DetachNetworkEndpoints(ctx context.Context, req *pb.DetachNetworkEndpointsNetworkEndpointGroupRequest) (*pb.Operation, error) {
	reqName := "projects/" + req.GetProject() + "/zones/" + req.GetZone() + "/networkEndpointGroups/" + req.GetNetworkEndpointGroup()
	name, err := s.parseZonalNetworkEndpointGroupName(reqName)
	if err != nil {
		return nil, err
	}

	fqn := name.String()

	neg := &pb.NetworkEndpointGroup{}
	if err := s.storage.Get(ctx, fqn, neg); err != nil {
		return nil, status.Errorf(codes.NotFound, "The resource '%s' was not found", fqn)
	}

	endpoints := req.GetNetworkEndpointGroupsDetachEndpointsRequestResource().GetNetworkEndpoints()
	for _, endpoint := range endpoints {
		instanceName := ""
		if endpoint.Instance != nil {
			instanceName = lastComponent(*endpoint.Instance)
		}
		ipAddress := endpoint.GetIpAddress()
		port := endpoint.GetPort()
		endpointFQN := fmt.Sprintf("%s/networkEndpoints/%s-%d-%s", fqn, ipAddress, port, instanceName)

		deleted := &pb.NetworkEndpoint{}
		if err := s.storage.Delete(ctx, endpointFQN, deleted); err != nil {
			if status.Code(err) != codes.NotFound {
				return nil, err
			}
		}
	}

	var count int32
	kind := (&pb.NetworkEndpoint{}).ProtoReflect().Descriptor()
	prefix := fqn + "/networkEndpoints/"
	err = s.storage.List(ctx, kind, storage.ListOptions{Prefix: prefix}, func(obj proto.Message) error {
		count++
		return nil
	})
	if err == nil {
		neg.Size = &count
		_ = s.storage.Update(ctx, fqn, neg)
	}

	op := &pb.Operation{
		TargetId:      neg.Id,
		TargetLink:    neg.SelfLink,
		OperationType: PtrTo("DetachNetworkEndpoints"),
		User:          PtrTo("user@example.com"),
	}
	return s.startZonalLRO(ctx, name.Project.ID, name.Zone, op, func() (proto.Message, error) {
		return neg, nil
	})
}

func (s *NetworkEndpointGroupsV1) ListNetworkEndpoints(ctx context.Context, req *pb.ListNetworkEndpointsNetworkEndpointGroupsRequest) (*pb.NetworkEndpointGroupsListNetworkEndpoints, error) {
	reqName := "projects/" + req.GetProject() + "/zones/" + req.GetZone() + "/networkEndpointGroups/" + req.GetNetworkEndpointGroup()
	name, err := s.parseZonalNetworkEndpointGroupName(reqName)
	if err != nil {
		return nil, err
	}

	fqn := name.String()

	neg := &pb.NetworkEndpointGroup{}
	if err := s.storage.Get(ctx, fqn, neg); err != nil {
		return nil, status.Errorf(codes.NotFound, "The resource '%s' was not found", fqn)
	}

	var items []*pb.NetworkEndpointWithHealthStatus
	kind := (&pb.NetworkEndpoint{}).ProtoReflect().Descriptor()
	prefix := fqn + "/networkEndpoints/"
	err = s.storage.List(ctx, kind, storage.ListOptions{Prefix: prefix}, func(obj proto.Message) error {
		endpoint := obj.(*pb.NetworkEndpoint)
		if endpoint.Instance != nil {
			inst := ExpandComputeLink(ctx, *endpoint.Instance)
			endpoint.Instance = &inst
		} else {
			endpoint.Instance = PtrTo("")
		}
		items = append(items, &pb.NetworkEndpointWithHealthStatus{
			NetworkEndpoint: endpoint,
		})
		return nil
	})
	if err != nil {
		return nil, err
	}

	response := &pb.NetworkEndpointGroupsListNetworkEndpoints{
		Id:    PtrTo(fmt.Sprintf("%d", ValueOf(neg.Id))),
		Kind:  PtrTo("compute#networkEndpointGroupsListNetworkEndpoints"),
		Items: items,
	}
	return response, nil
}

func (s *NetworkEndpointGroupsV1) List(ctx context.Context, req *pb.ListNetworkEndpointGroupsRequest) (*pb.NetworkEndpointGroupList, error) {
	prefix := fmt.Sprintf("projects/%s/zones/%s/networkEndpointGroups/", req.GetProject(), req.GetZone())

	var items []*pb.NetworkEndpointGroup
	kind := (&pb.NetworkEndpointGroup{}).ProtoReflect().Descriptor()
	err := s.storage.List(ctx, kind, storage.ListOptions{Prefix: prefix}, func(obj proto.Message) error {
		neg := obj.(*pb.NetworkEndpointGroup)
		items = append(items, neg)
		return nil
	})
	if err != nil {
		return nil, err
	}

	response := &pb.NetworkEndpointGroupList{
		Id:    PtrTo(""),
		Kind:  PtrTo("compute#networkEndpointGroupList"),
		Items: items,
	}
	return response, nil
}

type zonalNetworkEndpointGroupName struct {
	Project *projects.ProjectData
	Zone    string
	Name    string
}

func (n *zonalNetworkEndpointGroupName) String() string {
	return "projects/" + n.Project.ID + "/zones/" + n.Zone + "/networkEndpointGroups/" + n.Name
}

func (s *MockService) parseZonalNetworkEndpointGroupName(name string) (*zonalNetworkEndpointGroupName, error) {
	tokens := strings.Split(name, "/")

	if len(tokens) == 6 && tokens[0] == "projects" && tokens[2] == "zones" && tokens[4] == "networkEndpointGroups" {
		project, err := s.Projects.GetProjectByID(tokens[1])
		if err != nil {
			return nil, err
		}
		name := &zonalNetworkEndpointGroupName{
			Project: project,
			Zone:    tokens[3],
			Name:    tokens[5],
		}

		return name, nil
	} else {
		return nil, status.Errorf(codes.InvalidArgument, "name %q is not valid", name)
	}
}
