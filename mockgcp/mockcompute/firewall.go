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

// +tool:mockgcp-support
// proto.service: google.cloud.compute.v1.Firewalls
// proto.message: google.cloud.compute.v1.Firewall

package mockcompute

import (
	"context"
	"fmt"
	"strings"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/types/known/emptypb"

	"github.com/GoogleCloudPlatform/k8s-config-connector/mockgcp/common/projects"
	pb "github.com/GoogleCloudPlatform/k8s-config-connector/mockgcp/generated/mockgcp/cloud/compute/v1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/mockgcp/pkg/storage"
)

func (s *firewalls) Get(ctx context.Context, req *pb.GetFirewallRequest) (*pb.Firewall, error) {
	reqName := fmt.Sprintf("projects/%s/global/firewalls/%s", req.GetProject(), req.GetFirewall())
	name, err := s.parseFirewallName(reqName)
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

func (s *firewalls) Insert(ctx context.Context, req *pb.InsertFirewallRequest) (*pb.Operation, error) {
	reqName := fmt.Sprintf("projects/%s/global/firewalls/%s", req.GetProject(), req.GetFirewallResource().GetName())
	name, err := s.parseFirewallName(reqName)
	if err != nil {
		return nil, err
	}

	fqn := name.String()

	obj := ProtoClone(req.GetFirewallResource())
	obj.Name = PtrTo(name.Name)
	obj.Kind = PtrTo("compute#firewall")
	obj.Id = PtrTo(s.generateID())
	obj.CreationTimestamp = PtrTo(s.nowString())
	obj.SelfLink = PtrTo(buildComputeSelfLink(ctx, fqn))

	// Defaults
	if obj.Priority == nil {
		obj.Priority = PtrTo(int32(1000))
	}

	if err := s.normalizeNetworkURL(ctx, req.GetProject(), &obj.Network); err != nil {
		return nil, err
	}

	if obj.Direction == nil {
		obj.Direction = PtrTo("INGRESS")
	}
	if obj.Disabled == nil {
		obj.Disabled = PtrTo(false)
	}

	if obj.LogConfig == nil {
		obj.LogConfig = &pb.FirewallLogConfig{}
	}
	if obj.LogConfig.Enable == nil {
		obj.LogConfig.Enable = PtrTo(false)
	}

	if err := s.storage.Create(ctx, fqn, obj); err != nil {
		return nil, err
	}

	op := &pb.Operation{
		OperationType: PtrTo("insert"),
		TargetId:      obj.Id,
		TargetLink:    obj.SelfLink,
		User:          PtrTo("user@example.com"),
	}
	return s.computeOperations.startGlobalLRO(ctx, name.Project.ID, op, func() (proto.Message, error) {
		return obj, nil
	})
}

func (s *firewalls) Delete(ctx context.Context, req *pb.DeleteFirewallRequest) (*pb.Operation, error) {
	reqName := fmt.Sprintf("projects/%s/global/firewalls/%s", req.GetProject(), req.GetFirewall())
	name, err := s.parseFirewallName(reqName)
	if err != nil {
		return nil, err
	}

	fqn := name.String()

	deleted := &pb.Firewall{}
	if err := s.storage.Delete(ctx, fqn, deleted); err != nil {
		return nil, err
	}

	op := &pb.Operation{
		OperationType: PtrTo("delete"),
		TargetId:      deleted.Id,
		TargetLink:    deleted.SelfLink,
		User:          PtrTo("user@example.com"),
	}
	return s.computeOperations.startGlobalLRO(ctx, name.Project.ID, op, func() (proto.Message, error) {
		return &emptypb.Empty{}, nil
	})
}

func (s *firewalls) Patch(ctx context.Context, req *pb.PatchFirewallRequest) (*pb.Operation, error) {
	reqName := fmt.Sprintf("projects/%s/global/firewalls/%s", req.GetProject(), req.GetFirewall())
	name, err := s.parseFirewallName(reqName)
	if err != nil {
		return nil, err
	}

	fqn := name.String()
	obj := &pb.Firewall{}
	if err := s.storage.Get(ctx, fqn, obj); err != nil {
		return nil, err
	}

	// TODO: Implement JSON merge patch semantics if needed. For now, simple proto merge.
	// Note: proto.Merge appends repeated fields, which might not be exact JSON merge patch behavior for lists (which replace).
	// Compute Engine API documentation says "This method supports PATCH semantics and uses the JSON merge patch format".
	proto.Merge(obj, req.GetFirewallResource())

	// Allowed list replaces existing list rather than appending.
	if req.GetFirewallResource().GetAllowed() != nil {
		obj.Allowed = req.GetFirewallResource().GetAllowed()
	}

	// SourcRanges replaces existing list rather than appending.
	if req.GetFirewallResource().GetSourceRanges() != nil {
		obj.SourceRanges = req.GetFirewallResource().GetSourceRanges()
	}

	if err := s.storage.Update(ctx, fqn, obj); err != nil {
		return nil, err
	}

	op := &pb.Operation{
		OperationType: PtrTo("patch"),
		TargetId:      obj.Id,
		TargetLink:    obj.SelfLink,
		User:          PtrTo("user@example.com"),
	}
	return s.computeOperations.startGlobalLRO(ctx, name.Project.ID, op, func() (proto.Message, error) {
		return obj, nil
	})
}

func (s *firewalls) Update(ctx context.Context, req *pb.UpdateFirewallRequest) (*pb.Operation, error) {
	reqName := fmt.Sprintf("projects/%s/global/firewalls/%s", req.GetProject(), req.GetFirewall())
	name, err := s.parseFirewallName(reqName)
	if err != nil {
		return nil, err
	}

	fqn := name.String()
	obj := &pb.Firewall{}
	if err := s.storage.Get(ctx, fqn, obj); err != nil {
		return nil, err
	}

	// Update replaces the resource, but preserves immutable fields like ID, CreationTimestamp, SelfLink
	id := obj.Id
	creationTimestamp := obj.CreationTimestamp
	selfLink := obj.SelfLink

	obj = ProtoClone(req.GetFirewallResource())
	obj.Name = PtrTo(name.Name)
	obj.Id = id
	obj.CreationTimestamp = creationTimestamp
	obj.SelfLink = selfLink
	obj.Kind = PtrTo("compute#firewall")

	// Ensure defaults are re-applied if missing in PUT body (PUT replaces, so if fields are missing they are reset)
	if obj.Priority == nil {
		obj.Priority = PtrTo(int32(1000))
	}
	// Note: Network is usually required or defaults, but in PUT it should probably be present.
	if obj.Direction == nil {
		obj.Direction = PtrTo("INGRESS")
	}
	if obj.Disabled == nil {
		obj.Disabled = PtrTo(false)
	}

	if err := s.storage.Update(ctx, fqn, obj); err != nil {
		return nil, err
	}

	op := &pb.Operation{
		OperationType: PtrTo("update"),
		TargetId:      obj.Id,
		TargetLink:    obj.SelfLink,
		User:          PtrTo("user@example.com"),
	}
	return s.computeOperations.startGlobalLRO(ctx, name.Project.ID, op, func() (proto.Message, error) {
		return obj, nil
	})
}

func (s *firewalls) List(ctx context.Context, req *pb.ListFirewallsRequest) (*pb.FirewallList, error) {
	// Name format for listing is usually just the project
	// The request has `project` field.
	projectID := req.GetProject()

	// Filter by project
	// The storage keys are fully qualified: projects/{project}/global/firewalls/{firewall}
	prefix := fmt.Sprintf("projects/%s/global/firewalls/", projectID)

	var items []*pb.Firewall
	kind := (&pb.Firewall{}).ProtoReflect().Descriptor()
	if err := s.storage.List(ctx, kind, storage.ListOptions{Prefix: prefix}, func(msg proto.Message) error {
		item := msg.(*pb.Firewall)
		items = append(items, item)
		return nil
	}); err != nil {
		return nil, err
	}

	return &pb.FirewallList{
		Items:    items,
		Kind:     PtrTo("compute#firewallList"),
		SelfLink: PtrTo(buildComputeSelfLink(ctx, fmt.Sprintf("projects/%s/global/firewalls", projectID))),
	}, nil
}

type firewallName struct {
	Project *projects.ProjectData
	Name    string
}

func (n *firewallName) String() string {
	return fmt.Sprintf("projects/%s/global/firewalls/%s", n.Project.ID, n.Name)
}

// parseFirewallName parses a string into a firewallName.
// The expected form is `projects/{project}/global/firewalls/{firewall}`.
func (s *MockService) parseFirewallName(name string) (*firewallName, error) {
	tokens := strings.Split(name, "/")
	if len(tokens) == 5 && tokens[0] == "projects" && tokens[2] == "global" && tokens[3] == "firewalls" {
		project, err := s.Projects.GetProjectByID(tokens[1])
		if err != nil {
			return nil, err
		}

		return &firewallName{
			Project: project,
			Name:    tokens[4],
		}, nil
	}
	return nil, status.Errorf(codes.InvalidArgument, "name %q is not valid", name)
}
