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

func (s *FirewallsV1) List(ctx context.Context, req *pb.ListFirewallsRequest) (*pb.FirewallList, error) {
	name, err := s.newFirewallName(req.GetProject(), "placeholder")
	if err != nil {
		return nil, err
	}

	findPrefix := strings.TrimSuffix(name.String(), "placeholder")

	response := &pb.FirewallList{}
	response.Id = PtrTo("0123456789")
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

func (s *FirewallsV1) Insert(ctx context.Context, req *pb.InsertFirewallRequest) (*pb.Operation, error) {
	name, err := s.newFirewallName(req.GetProject(), req.GetFirewallResource().GetName())
	if err != nil {
		return nil, err
	}

	fqn := name.String()

	id := s.generateID()

	obj := proto.CloneOf(req.GetFirewallResource())
	obj.CreationTimestamp = PtrTo(s.nowString())
	obj.Id = &id
	obj.SelfLink = PtrTo(BuildComputeSelfLink(ctx, name.String()))
	obj.Kind = PtrTo("compute#firewall")

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

	proto.Merge(obj, req.GetFirewallResource())

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

func (s *FirewallsV1) Update(ctx context.Context, req *pb.UpdateFirewallRequest) (*pb.Operation, error) {
	name, err := s.newFirewallName(req.GetProject(), req.GetFirewall())
	if err != nil {
		return nil, err
	}

	fqn := name.String()
	obj := &pb.Firewall{}
	if err := s.storage.Get(ctx, fqn, obj); err != nil {
		return nil, err
	}

	updatedObj := proto.CloneOf(req.GetFirewallResource())
	updatedObj.CreationTimestamp = obj.CreationTimestamp
	updatedObj.Id = obj.Id
	updatedObj.SelfLink = obj.SelfLink
	updatedObj.Kind = obj.Kind

	if err := s.storage.Update(ctx, fqn, updatedObj); err != nil {
		return nil, err
	}

	op := &pb.Operation{
		TargetId:      updatedObj.Id,
		TargetLink:    updatedObj.SelfLink,
		OperationType: PtrTo("compute.firewalls.update"),
		User:          PtrTo("user@example.com"),
	}
	return s.startGlobalLRO(ctx, name.Project.ID, op, func() (proto.Message, error) {
		return updatedObj, nil
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
