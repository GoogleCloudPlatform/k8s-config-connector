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
// proto.service: google.cloud.compute.v1.NetworkFirewallPolicies
// proto.message: google.cloud.compute.v1.FirewallPolicyAssociation

package mockcompute

import (
	"context"
	"fmt"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/proto"

	pb "github.com/GoogleCloudPlatform/k8s-config-connector/mockgcp/generated/mockgcp/cloud/compute/v1"
)

func (s *NetworkFirewallPoliciesV1) GetAssociation(ctx context.Context, req *pb.GetAssociationNetworkFirewallPolicyRequest) (*pb.FirewallPolicyAssociation, error) {
	reqName := fmt.Sprintf("projects/%s/global/firewallPolicies/%s", req.GetProject(), req.GetFirewallPolicy())
	name, err := s.parseNetworkFirewallPolicyName(reqName)
	if err != nil {
		return nil, err
	}

	fqn := name.String()

	obj := &pb.FirewallPolicy{}
	if err := s.storage.Get(ctx, fqn, obj); err != nil {
		return nil, err
	}

	for _, association := range obj.GetAssociations() {
		if association.GetName() == req.GetName() {
			return association, nil
		}
	}

	return nil, status.Errorf(codes.NotFound, "Association %q not found", req.GetName())
}

func (s *NetworkFirewallPoliciesV1) AddAssociation(ctx context.Context, req *pb.AddAssociationNetworkFirewallPolicyRequest) (*pb.Operation, error) {
	reqName := fmt.Sprintf("projects/%s/global/firewallPolicies/%s", req.GetProject(), req.GetFirewallPolicy())
	name, err := s.parseNetworkFirewallPolicyName(reqName)
	if err != nil {
		return nil, err
	}

	fqn := name.String()
	obj := &pb.FirewallPolicy{}
	if err := s.storage.Get(ctx, fqn, obj); err != nil {
		return nil, err
	}

	association := req.GetFirewallPolicyAssociationResource()

	if err := s.normalizeNetworkURL(ctx, name.Project.ID, &association.AttachmentTarget); err != nil {
		return nil, err
	}

	// Check for duplicates
	for _, existing := range obj.Associations {
		if existing.GetName() == association.GetName() {
			return nil, status.Errorf(codes.AlreadyExists, "Association %q already exists", association.GetName())
		}
	}

	obj.Associations = append(obj.Associations, association)

	if err := s.storage.Update(ctx, fqn, obj); err != nil {
		return nil, err
	}

	op := &pb.Operation{
		OperationType: PtrTo("compute.networkFirewallPolicy.attach"),
		TargetId:      obj.Id,
		TargetLink:    obj.SelfLink,
		User:          PtrTo("user@example.com"),
	}
	return s.computeOperations.startGlobalLRO(ctx, name.Project.ID, op, func() (proto.Message, error) {
		return nil, nil
	})
}

func (s *NetworkFirewallPoliciesV1) RemoveAssociation(ctx context.Context, req *pb.RemoveAssociationNetworkFirewallPolicyRequest) (*pb.Operation, error) {
	reqName := fmt.Sprintf("projects/%s/global/firewallPolicies/%s", req.GetProject(), req.GetFirewallPolicy())
	name, err := s.parseNetworkFirewallPolicyName(reqName)
	if err != nil {
		return nil, err
	}

	fqn := name.String()
	obj := &pb.FirewallPolicy{}
	if err := s.storage.Get(ctx, fqn, obj); err != nil {
		return nil, err
	}

	var associations []*pb.FirewallPolicyAssociation
	found := false
	for _, association := range obj.Associations {
		if association.GetName() == req.GetName() {
			found = true
			continue
		}
		associations = append(associations, association)
	}

	if !found {
		return nil, status.Errorf(codes.NotFound, "Association %q not found", req.GetName())
	}

	obj.Associations = associations

	if err := s.storage.Update(ctx, fqn, obj); err != nil {
		return nil, err
	}

	op := &pb.Operation{
		OperationType: PtrTo("compute.networkFirewallPolicy.detach"),
		TargetId:      obj.Id,
		TargetLink:    obj.SelfLink,
		User:          PtrTo("user@example.com"),
	}
	return s.computeOperations.startGlobalLRO(ctx, name.Project.ID, op, func() (proto.Message, error) {
		return nil, nil
	})
}
