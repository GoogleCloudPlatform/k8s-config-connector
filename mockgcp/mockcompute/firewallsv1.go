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
	"strconv"
	"strings"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/proto"

	"github.com/GoogleCloudPlatform/k8s-config-connector/mockgcp/common/projects"
	pb "github.com/GoogleCloudPlatform/k8s-config-connector/mockgcp/generated/mockgcp/cloud/compute/v1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/mockgcp/pkg/storage"
)

type FirewallsV1 struct {
	*MockService
	pb.UnimplementedFirewallsServer
}

func (s *FirewallsV1) List(ctx context.Context, req *pb.ListFirewallsRequest) (*pb.FirewallList, error) {
	name, err := s.newFirewallName(req.GetProject(), "placeholder")
	if err != nil {
		return nil, err
	}

	findPrefix := strings.TrimSuffix(name.String(), "placeholder")

	response := &pb.FirewallList{}
	response.Id = PtrTo(strconv.FormatUint(s.generateID(), 10))
	response.Kind = PtrTo("compute#firewallList")
	response.SelfLink = PtrTo(BuildComputeSelfLink(ctx, strings.TrimSuffix(findPrefix, "/")))

	findKind := (&pb.Firewall{}).ProtoReflect().Descriptor()
	if err := s.storage.List(ctx, findKind, storage.ListOptions{Prefix: findPrefix}, func(obj proto.Message) error {
		firewall := obj.(*pb.Firewall)
		isMatch, err := matchFilter(req.GetFilter(), firewall)
		if err != nil {
			return err
		}
		if isMatch {
			response.Items = append(response.Items, firewall)
		}
		return nil
	}); err != nil {
		return nil, err
	}

	return response, nil
}

func (s *FirewallsV1) Get(ctx context.Context, req *pb.GetFirewallRequest) (*pb.Firewall, error) {
	name, err := s.newFirewallName(req.GetProject(), req.GetFirewall())
	if err != nil {
		return nil, err
	}

	fqn := name.String()

	obj := &pb.Firewall{}
	if err := s.storage.Get(ctx, fqn, obj); err != nil {
		if status.Code(err) == codes.NotFound {
			return nil, status.Errorf(codes.NotFound, "The resource '%s' was not found", fqn)
		}
		return nil, err
	}

	return obj, nil
}

func (s *FirewallsV1) Insert(ctx context.Context, req *pb.InsertFirewallRequest) (*pb.Operation, error) {
	name, err := s.newFirewallName(req.GetProject(), req.GetFirewallResource().GetName())
	if err != nil {
		return nil, err
	}

	fqn := name.String()

	id := s.generateID()

	obj := proto.Clone(req.GetFirewallResource()).(*pb.Firewall)
	obj.CreationTimestamp = PtrTo(s.nowString())
	obj.Id = &id
	obj.SelfLink = PtrTo(BuildComputeSelfLink(ctx, name.String()))
	obj.Kind = PtrTo("compute#firewall")

	// Set default values if omitted
	if obj.Priority == nil {
		obj.Priority = PtrTo(int32(1000))
	}
	if obj.Direction == nil {
		obj.Direction = PtrTo("INGRESS")
	}
	if obj.Disabled == nil {
		obj.Disabled = PtrTo(false)
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
	return s.startGlobalLRO(ctx, name.Project.ID, op, func() (proto.Message, error) {
		return obj, nil
	})
}

func (s *FirewallsV1) Patch(ctx context.Context, req *pb.PatchFirewallRequest) (*pb.Operation, error) {
	name, err := s.newFirewallName(req.GetProject(), req.GetFirewall())
	if err != nil {
		return nil, err
	}

	fqn := name.String()
	obj := &pb.Firewall{}
	if err := s.storage.Get(ctx, fqn, obj); err != nil {
		return nil, err
	}

	// Update mutable fields from requested resource.
	reqObj := req.GetFirewallResource()
	if reqObj.Allowed != nil {
		obj.Allowed = reqObj.Allowed
	}
	if reqObj.Denied != nil {
		obj.Denied = reqObj.Denied
	}
	if reqObj.Description != nil {
		obj.Description = reqObj.Description
	}
	if reqObj.DestinationRanges != nil {
		obj.DestinationRanges = reqObj.DestinationRanges
	}
	if reqObj.Direction != nil {
		obj.Direction = reqObj.Direction
	}
	if reqObj.Disabled != nil {
		obj.Disabled = reqObj.Disabled
	}
	if reqObj.LogConfig != nil {
		obj.LogConfig = reqObj.LogConfig
	}
	if reqObj.Network != nil {
		obj.Network = reqObj.Network
	}
	if reqObj.Priority != nil {
		obj.Priority = reqObj.Priority
	}
	if reqObj.SourceRanges != nil {
		obj.SourceRanges = reqObj.SourceRanges
	}
	if reqObj.SourceServiceAccounts != nil {
		obj.SourceServiceAccounts = reqObj.SourceServiceAccounts
	}
	if reqObj.SourceTags != nil {
		obj.SourceTags = reqObj.SourceTags
	}
	if reqObj.TargetServiceAccounts != nil {
		obj.TargetServiceAccounts = reqObj.TargetServiceAccounts
	}
	if reqObj.TargetTags != nil {
		obj.TargetTags = reqObj.TargetTags
	}

	if err := s.storage.Update(ctx, fqn, obj); err != nil {
		return nil, err
	}

	op := &pb.Operation{
		TargetId:      obj.Id,
		TargetLink:    obj.SelfLink,
		OperationType: PtrTo("compute.firewalls.patch"),
		User:          PtrTo("user@example.com"),
	}
	return s.startGlobalLRO(ctx, name.Project.ID, op, func() (proto.Message, error) {
		return obj, nil
	})
}

func (s *FirewallsV1) Delete(ctx context.Context, req *pb.DeleteFirewallRequest) (*pb.Operation, error) {
	name, err := s.newFirewallName(req.GetProject(), req.GetFirewall())
	if err != nil {
		return nil, err
	}

	fqn := name.String()

	deleted := &pb.Firewall{}
	if err := s.storage.Delete(ctx, fqn, deleted); err != nil {
		return nil, err
	}

	op := &pb.Operation{
		TargetId:      deleted.Id,
		TargetLink:    deleted.SelfLink,
		OperationType: PtrTo("delete"),
		User:          PtrTo("user@example.com"),
	}
	return s.startGlobalLRO(ctx, name.Project.ID, op, func() (proto.Message, error) {
		return deleted, nil
	})
}

type firewallName struct {
	Project *projects.ProjectData
	Name    string
}

func (n *firewallName) String() string {
	return "projects/" + n.Project.ID + "/global" + "/firewalls/" + n.Name
}

func (s *MockService) parseFirewallName(name string) (*firewallName, error) {
	tokens := strings.Split(name, "/")

	if len(tokens) == 5 && tokens[0] == "projects" && tokens[2] == "global" && tokens[3] == "firewalls" {
		return s.newFirewallName(tokens[1], tokens[4])
	}
	return nil, status.Errorf(codes.InvalidArgument, "name %q is not valid", name)
}

func (s *MockService) newFirewallName(project string, name string) (*firewallName, error) {
	projectObj, err := s.Projects.GetProjectByIDOrNumber(project)
	if err != nil {
		return nil, err
	}

	return &firewallName{
		Project: projectObj,
		Name:    name,
	}, nil
}
