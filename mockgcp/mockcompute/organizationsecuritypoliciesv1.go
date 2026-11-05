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
	"strconv"
	"strings"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/proto"

	pb "github.com/GoogleCloudPlatform/k8s-config-connector/mockgcp/generated/mockgcp/cloud/compute/v1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/mockgcp/pkg/storage"
)

type OrganizationSecurityPoliciesV1 struct {
	*MockService
	pb.UnimplementedOrganizationSecurityPoliciesServer
}

type organizationSecurityPolicyName struct {
	Name string
}

func (n *organizationSecurityPolicyName) String() string {
	return "locations/global/securityPolicies/" + n.Name
}

func (s *MockService) parseOrganizationSecurityPolicyName(name string) (*organizationSecurityPolicyName, error) {
	tokens := strings.Split(name, "/")

	if len(tokens) == 4 && tokens[0] == "locations" && tokens[1] == "global" && tokens[2] == "securityPolicies" {
		return &organizationSecurityPolicyName{
			Name: tokens[3],
		}, nil
	} else if len(tokens) == 1 {
		return &organizationSecurityPolicyName{
			Name: tokens[0],
		}, nil
	} else {
		return nil, status.Errorf(codes.InvalidArgument, "name %q is not valid", name)
	}
}

func (s *OrganizationSecurityPoliciesV1) Get(ctx context.Context, req *pb.GetOrganizationSecurityPolicyRequest) (*pb.SecurityPolicy, error) {
	reqName := "locations/global/securityPolicies/" + req.GetSecurityPolicy()
	name, err := s.parseOrganizationSecurityPolicyName(reqName)
	if err != nil {
		return nil, err
	}

	fqn := name.String()

	obj := &pb.SecurityPolicy{}
	if err := s.storage.Get(ctx, fqn, obj); err != nil {
		if status.Code(err) == codes.NotFound {
			return nil, status.Errorf(codes.NotFound, "The resource '%s' was not found", fqn)
		}
		return nil, err
	}

	return obj, nil
}

func (s *OrganizationSecurityPoliciesV1) Insert(ctx context.Context, req *pb.InsertOrganizationSecurityPolicyRequest) (*pb.Operation, error) {
	id := s.generateID()
	policyId := strconv.FormatUint(id, 10)

	reqName := "locations/global/securityPolicies/" + policyId
	policyName, err := s.parseOrganizationSecurityPolicyName(reqName)
	if err != nil {
		return nil, err
	}

	fqn := policyName.String()

	obj := proto.Clone(req.GetSecurityPolicyResource()).(*pb.SecurityPolicy)
	obj.SelfLink = PtrTo(BuildComputeSelfLink(ctx, policyName.String()))
	obj.Parent = req.ParentId
	obj.Id = PtrTo(id)
	obj.Name = PtrTo(policyId)
	obj.CreationTimestamp = PtrTo(s.nowString())
	obj.Kind = PtrTo("compute#securityPolicy")
	if obj.Fingerprint == nil {
		obj.Fingerprint = PtrTo(computeFingerprint(obj))
	}
	if obj.LabelFingerprint == nil {
		obj.LabelFingerprint = PtrTo(computeFingerprint(obj))
	}

	// Default rules for CLOUD_ARMOR
	if len(obj.Rules) == 0 {
		obj.Rules = []*pb.SecurityPolicyRule{
			{
				Action:      PtrTo("goto_next"),
				Description: PtrTo("default rule"),
				Kind:        PtrTo("compute#securityPolicyRule"),
				Match: &pb.SecurityPolicyRuleMatcher{
					Config: &pb.SecurityPolicyRuleMatcherConfig{
						SrcIpRanges: []string{"*"},
					},
					VersionedExpr: PtrTo("SRC_IPS_V1"),
				},
				Preview:  PtrTo(false),
				Priority: PtrTo(int32(2147483647)),
			},
		}
	}

	if err := s.storage.Create(ctx, fqn, obj); err != nil {
		return nil, err
	}

	op := &pb.Operation{
		TargetId:      obj.Id,
		TargetLink:    obj.SelfLink,
		OperationType: PtrTo("createOrganizationSecurityPolicy"),
		User:          PtrTo("user@example.com"),
	}
	return s.startGlobalOrganizationLRO(ctx, op, func() (proto.Message, error) {
		return obj, nil
	})
}

func (s *OrganizationSecurityPoliciesV1) Patch(ctx context.Context, req *pb.PatchOrganizationSecurityPolicyRequest) (*pb.Operation, error) {
	reqName := "locations/global/securityPolicies/" + req.GetSecurityPolicy()

	name, err := s.parseOrganizationSecurityPolicyName(reqName)
	if err != nil {
		return nil, err
	}

	fqn := name.String()
	obj := &pb.SecurityPolicy{}
	if err := s.storage.Get(ctx, fqn, obj); err != nil {
		return nil, err
	}

	if req.GetSecurityPolicyResource().Description != nil {
		obj.Description = req.GetSecurityPolicyResource().Description
	}
	obj.Fingerprint = PtrTo(computeFingerprint(obj))

	if err := s.storage.Update(ctx, fqn, obj); err != nil {
		return nil, err
	}

	op := &pb.Operation{
		TargetId:      obj.Id,
		TargetLink:    obj.SelfLink,
		OperationType: PtrTo("patchOrganizationSecurityPolicy"),
		User:          PtrTo("user@example.com"),
	}
	return s.startGlobalOrganizationLRO(ctx, op, func() (proto.Message, error) {
		return obj, nil
	})
}

func (s *OrganizationSecurityPoliciesV1) Delete(ctx context.Context, req *pb.DeleteOrganizationSecurityPolicyRequest) (*pb.Operation, error) {
	reqName := "locations/global/securityPolicies/" + req.GetSecurityPolicy()
	name, err := s.parseOrganizationSecurityPolicyName(reqName)
	if err != nil {
		return nil, err
	}

	fqn := name.String()

	deleted := &pb.SecurityPolicy{}
	if err := s.storage.Delete(ctx, fqn, deleted); err != nil {
		return nil, err
	}

	op := &pb.Operation{
		TargetId:      deleted.Id,
		TargetLink:    deleted.SelfLink,
		OperationType: PtrTo("deleteOrganizationSecurityPolicy"),
		User:          PtrTo("user@example.com"),
	}
	return s.startGlobalOrganizationLRO(ctx, op, func() (proto.Message, error) {
		return deleted, nil
	})
}

func (s *OrganizationSecurityPoliciesV1) List(ctx context.Context, req *pb.ListOrganizationSecurityPoliciesRequest) (*pb.SecurityPolicyList, error) {
	response := &pb.SecurityPolicyList{
		Kind: PtrTo("compute#securityPolicyList"),
		Id:   PtrTo(""),
	}

	policyKind := (&pb.SecurityPolicy{}).ProtoReflect().Descriptor()
	if err := s.storage.List(ctx, policyKind, storage.ListOptions{}, func(obj proto.Message) error {
		policy := obj.(*pb.SecurityPolicy)
		if strings.HasPrefix(policy.GetSelfLink(), "https://www.googleapis.com/compute/v1/locations/global/securityPolicies/") {
			if req.ParentId == nil || policy.GetParent() == req.GetParentId() {
				response.Items = append(response.Items, policy)
			}
		}
		return nil
	}); err != nil {
		return nil, err
	}

	return response, nil
}

func (s *OrganizationSecurityPoliciesV1) AddAssociation(ctx context.Context, req *pb.AddAssociationOrganizationSecurityPolicyRequest) (*pb.Operation, error) {
	reqName := "locations/global/securityPolicies/" + req.GetSecurityPolicy()
	name, err := s.parseOrganizationSecurityPolicyName(reqName)
	if err != nil {
		return nil, err
	}

	fqn := name.String()

	policy := &pb.SecurityPolicy{}
	if err := s.storage.Get(ctx, fqn, policy); err != nil {
		return nil, err
	}

	association := proto.Clone(req.GetSecurityPolicyAssociationResource()).(*pb.SecurityPolicyAssociation)
	assocName := association.GetName()
	if assocName == "" {
		return nil, status.Errorf(codes.InvalidArgument, "association name is required")
	}

	found := false
	for i, existing := range policy.Associations {
		if existing.GetName() == assocName {
			if !req.GetReplaceExistingAssociation() {
				return nil, status.Errorf(codes.AlreadyExists, "association %q already exists", assocName)
			}
			policy.Associations[i] = association
			found = true
			break
		}
	}
	if !found {
		policy.Associations = append(policy.Associations, association)
	}

	if err := s.storage.Update(ctx, fqn, policy); err != nil {
		return nil, err
	}

	assocFQN := fmt.Sprintf("%s/associations/%s", fqn, assocName)
	if found {
		if err := s.storage.Update(ctx, assocFQN, association); err != nil {
			return nil, err
		}
	} else {
		if err := s.storage.Create(ctx, assocFQN, association); err != nil {
			return nil, err
		}
	}

	op := &pb.Operation{
		TargetId:      policy.Id,
		TargetLink:    policy.SelfLink,
		OperationType: PtrTo("attachOrganizationSecurityPolicy"),
		User:          PtrTo("user@example.com"),
	}
	return s.startGlobalOrganizationLRO(ctx, op, func() (proto.Message, error) {
		return association, nil
	})
}

func (s *OrganizationSecurityPoliciesV1) GetAssociation(ctx context.Context, req *pb.GetAssociationOrganizationSecurityPolicyRequest) (*pb.SecurityPolicyAssociation, error) {
	reqName := "locations/global/securityPolicies/" + req.GetSecurityPolicy()
	name, err := s.parseOrganizationSecurityPolicyName(reqName)
	if err != nil {
		return nil, err
	}

	fqn := name.String()
	assocFQN := fmt.Sprintf("%s/associations/%s", fqn, req.GetName())

	assoc := &pb.SecurityPolicyAssociation{}
	if err := s.storage.Get(ctx, assocFQN, assoc); err != nil {
		if status.Code(err) == codes.NotFound {
			return nil, status.Errorf(codes.InvalidArgument, "Invalid value for field 'name': '%s'. An association with that name does not exist.", req.GetName())
		}
		return nil, err
	}

	return assoc, nil
}

func (s *OrganizationSecurityPoliciesV1) RemoveAssociation(ctx context.Context, req *pb.RemoveAssociationOrganizationSecurityPolicyRequest) (*pb.Operation, error) {
	reqName := "locations/global/securityPolicies/" + req.GetSecurityPolicy()
	name, err := s.parseOrganizationSecurityPolicyName(reqName)
	if err != nil {
		return nil, err
	}

	fqn := name.String()

	policy := &pb.SecurityPolicy{}
	if err := s.storage.Get(ctx, fqn, policy); err != nil {
		return nil, err
	}

	assocFQN := fmt.Sprintf("%s/associations/%s", fqn, req.GetName())
	deleted := &pb.SecurityPolicyAssociation{}
	if err := s.storage.Delete(ctx, assocFQN, deleted); err != nil {
		if status.Code(err) == codes.NotFound {
			return nil, status.Errorf(codes.InvalidArgument, "Invalid value for field 'name': '%s'. An association with that name does not exist.", req.GetName())
		}
		return nil, err
	}

	var associations []*pb.SecurityPolicyAssociation
	for _, assoc := range policy.Associations {
		if assoc.GetName() != req.GetName() {
			associations = append(associations, assoc)
		}
	}
	policy.Associations = associations

	if err := s.storage.Update(ctx, fqn, policy); err != nil {
		return nil, err
	}

	op := &pb.Operation{
		TargetId:      policy.Id,
		TargetLink:    policy.SelfLink,
		OperationType: PtrTo("detachOrganizationSecurityPolicy"),
		User:          PtrTo("user@example.com"),
	}
	return s.startGlobalOrganizationLRO(ctx, op, func() (proto.Message, error) {
		return deleted, nil
	})
}

func (s *OrganizationSecurityPoliciesV1) ListAssociations(ctx context.Context, req *pb.ListAssociationsOrganizationSecurityPolicyRequest) (*pb.OrganizationSecurityPoliciesListAssociationsResponse, error) {
	if req.GetTargetResource() == "" {
		return nil, status.Errorf(codes.InvalidArgument, "targetResource is required")
	}

	response := &pb.OrganizationSecurityPoliciesListAssociationsResponse{
		Kind: PtrTo("compute#organizationSecurityPoliciesListAssociations"),
	}

	policyKind := (&pb.SecurityPolicy{}).ProtoReflect().Descriptor()
	if err := s.storage.List(ctx, policyKind, storage.ListOptions{}, func(obj proto.Message) error {
		policy := obj.(*pb.SecurityPolicy)
		for _, assoc := range policy.Associations {
			if assoc.GetAttachmentId() == req.GetTargetResource() {
				response.Associations = append(response.Associations, assoc)
			}
		}
		return nil
	}); err != nil {
		return nil, err
	}

	return response, nil
}
