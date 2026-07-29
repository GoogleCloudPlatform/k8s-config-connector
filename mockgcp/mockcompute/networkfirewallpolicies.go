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

type NetworkFirewallPoliciesV1 struct {
	*MockService
	pb.UnimplementedNetworkFirewallPoliciesServer
}

func (s *NetworkFirewallPoliciesV1) resolvePolicy(ctx context.Context, project string, policyNameOrID string) (*pb.FirewallPolicy, string, string, error) {
	projectObj, err := s.Projects.GetProjectByIDOrNumber(project)
	if err != nil {
		return nil, "", "", err
	}

	// Try parsing as ID (number)
	if _, err := strconv.ParseUint(policyNameOrID, 10, 64); err == nil {
		obj, fqn, err := s.findNetworkFirewallPolicyByID(ctx, projectObj, policyNameOrID)
		return obj, fqn, projectObj.ID, err
	}

	// Otherwise treat as name
	name, err := s.newNetworkFirewallPolicyName(projectObj, policyNameOrID)
	if err != nil {
		return nil, "", "", err
	}
	fqn := name.String()
	obj := &pb.FirewallPolicy{}
	if err := s.storage.Get(ctx, fqn, obj); err != nil {
		if status.Code(err) == codes.NotFound {
			return nil, "", "", status.Errorf(codes.NotFound, "The resource '%s' was not found", fqn)
		}
		return nil, "", "", err
	}
	return obj, fqn, projectObj.ID, nil
}

func (s *NetworkFirewallPoliciesV1) findNetworkFirewallPolicyByID(ctx context.Context, projectObj *projects.ProjectData, idStr string) (*pb.FirewallPolicy, string, error) {
	id, err := strconv.ParseUint(idStr, 10, 64)
	if err != nil {
		return nil, "", status.Errorf(codes.InvalidArgument, "invalid policy ID: %s", idStr)
	}

	prefix := "projects/" + projectObj.ID + "/global/firewallPolicies/"

	var foundObj *pb.FirewallPolicy
	var foundFQN string

	err = s.storage.List(ctx, (&pb.FirewallPolicy{}).ProtoReflect().Descriptor(), storage.ListOptions{Prefix: prefix}, func(obj proto.Message) error {
		policy := obj.(*pb.FirewallPolicy)
		if policy.GetId() == id {
			foundObj = policy
			foundFQN = prefix + policy.GetName()
		}
		return nil
	})

	if err != nil {
		return nil, "", err
	}

	if foundObj == nil {
		fqnWithID := prefix + idStr
		return nil, "", status.Errorf(codes.NotFound, "The resource '%s' was not found", fqnWithID)
	}

	return foundObj, foundFQN, nil
}

func (s *NetworkFirewallPoliciesV1) Get(ctx context.Context, req *pb.GetNetworkFirewallPolicyRequest) (*pb.FirewallPolicy, error) {
	obj, _, _, err := s.resolvePolicy(ctx, req.GetProject(), req.GetFirewallPolicy())
	if err != nil {
		return nil, err
	}
	return obj, nil
}

func (s *NetworkFirewallPoliciesV1) Insert(ctx context.Context, req *pb.InsertNetworkFirewallPolicyRequest) (*pb.Operation, error) {
	id := s.generateID()
	policyId := strconv.FormatUint(id, 10)

	policyName := req.GetFirewallPolicyResource().GetName()
	if policyName == "" {
		policyName = policyId
	}

	projectObj, err := s.Projects.GetProjectByIDOrNumber(req.GetProject())
	if err != nil {
		return nil, err
	}

	name, err := s.newNetworkFirewallPolicyName(projectObj, policyName)
	if err != nil {
		return nil, err
	}

	fqn := name.String()

	obj := proto.CloneOf(req.GetFirewallPolicyResource())
	obj.SelfLink = PtrTo(BuildComputeSelfLink(ctx, name.String()))
	obj.SelfLinkWithId = PtrTo(BuildComputeSelfLink(ctx, name.String()))
	obj.RuleTupleCount = PtrTo(int32(8))
	obj.Id = PtrTo(id)
	obj.CreationTimestamp = PtrTo(s.nowString())
	obj.Kind = PtrTo("compute#firewallPolicy")

	if obj.Fingerprint == nil {
		obj.Fingerprint = PtrTo(computeFingerprint(obj))
	}

	if obj.Rules == nil {
		populateDefaultRules(obj)
	}

	if err := s.storage.Create(ctx, fqn, obj); err != nil {
		return nil, err
	}

	op := &pb.Operation{
		TargetId:      obj.Id,
		TargetLink:    obj.SelfLink,
		OperationType: PtrTo("compute.networkFirewallPolicy.insert"),
		User:          PtrTo("user@example.com"),
	}
	return s.startGlobalLRO(ctx, name.Project.ID, op, func() (proto.Message, error) {
		return obj, nil
	})
}

func (s *NetworkFirewallPoliciesV1) Patch(ctx context.Context, req *pb.PatchNetworkFirewallPolicyRequest) (*pb.Operation, error) {
	obj, fqn, projectID, err := s.resolvePolicy(ctx, req.GetProject(), req.GetFirewallPolicy())
	if err != nil {
		return nil, err
	}

	proto.Merge(obj, req.GetFirewallPolicyResource())

	if err := s.storage.Update(ctx, fqn, obj); err != nil {
		return nil, err
	}

	op := &pb.Operation{
		TargetId:      obj.Id,
		TargetLink:    obj.SelfLink,
		OperationType: PtrTo("compute.networkFirewallPolicy.patch"),
		User:          PtrTo("user@example.com"),
	}
	return s.startGlobalLRO(ctx, projectID, op, func() (proto.Message, error) {
		return obj, nil
	})
}

func (s *NetworkFirewallPoliciesV1) Delete(ctx context.Context, req *pb.DeleteNetworkFirewallPolicyRequest) (*pb.Operation, error) {
	_, fqn, projectID, err := s.resolvePolicy(ctx, req.GetProject(), req.GetFirewallPolicy())
	if err != nil {
		return nil, err
	}

	deleted := &pb.FirewallPolicy{}
	if err := s.storage.Delete(ctx, fqn, deleted); err != nil {
		return nil, err
	}

	op := &pb.Operation{
		TargetId:      deleted.Id,
		TargetLink:    deleted.SelfLink,
		OperationType: PtrTo("compute.networkFirewallPolicy.delete"),
		User:          PtrTo("user@example.com"),
	}
	return s.startGlobalLRO(ctx, projectID, op, func() (proto.Message, error) {
		return deleted, nil
	})
}

func (s *NetworkFirewallPoliciesV1) GetRule(ctx context.Context, req *pb.GetRuleNetworkFirewallPolicyRequest) (*pb.FirewallPolicyRule, error) {
	obj, _, _, err := s.resolvePolicy(ctx, req.GetProject(), req.GetFirewallPolicy())
	if err != nil {
		return nil, err
	}

	var rule *pb.FirewallPolicyRule
	rules := obj.GetRules()

	for _, r := range rules {
		if r.Priority != nil && *r.Priority == *req.Priority {
			rule = r
		}
	}
	if rule == nil {
		return nil, status.Errorf(codes.InvalidArgument, "Invalid value for field 'priority': '%d'. The firewall policy does not contain a rule at priority %d.", int(*req.Priority), int(*req.Priority))
	}

	return rule, nil
}

func (s *NetworkFirewallPoliciesV1) AddRule(ctx context.Context, req *pb.AddRuleNetworkFirewallPolicyRequest) (*pb.Operation, error) {
	obj, fqn, projectID, err := s.resolvePolicy(ctx, req.GetProject(), req.GetFirewallPolicy())
	if err != nil {
		return nil, err
	}

	r := req.GetFirewallPolicyRuleResource()
	mockFieldValuesForRule(r)

	obj.Rules = append(obj.Rules, r)

	if err := s.storage.Update(ctx, fqn, obj); err != nil {
		return nil, err
	}

	op := &pb.Operation{
		TargetId:      obj.Id,
		TargetLink:    obj.SelfLink,
		OperationType: PtrTo("compute.networkFirewallPolicy.addRule"),
		User:          PtrTo("user@example.com"),
	}
	return s.startGlobalLRO(ctx, projectID, op, func() (proto.Message, error) {
		return obj, nil
	})
}

func (s *NetworkFirewallPoliciesV1) PatchRule(ctx context.Context, req *pb.PatchRuleNetworkFirewallPolicyRequest) (*pb.Operation, error) {
	obj, fqn, projectID, err := s.resolvePolicy(ctx, req.GetProject(), req.GetFirewallPolicy())
	if err != nil {
		return nil, err
	}

	rules := []*pb.FirewallPolicyRule{}
	for _, rule := range obj.Rules {
		if rule.Priority != nil && *rule.Priority == *req.Priority {
			r := req.GetFirewallPolicyRuleResource()
			r.Priority = PtrTo(*rule.Priority)
			mockFieldValuesForRule(r)
			rules = append(rules, r)
		} else {
			rules = append(rules, rule)
		}
	}

	obj.Rules = rules
	if err := s.storage.Update(ctx, fqn, obj); err != nil {
		return nil, err
	}

	op := &pb.Operation{
		TargetId:      obj.Id,
		TargetLink:    obj.SelfLink,
		OperationType: PtrTo("compute.networkFirewallPolicy.patchRule"),
		User:          PtrTo("user@example.com"),
	}
	return s.startGlobalLRO(ctx, projectID, op, func() (proto.Message, error) {
		return obj, nil
	})
}

func (s *NetworkFirewallPoliciesV1) RemoveRule(ctx context.Context, req *pb.RemoveRuleNetworkFirewallPolicyRequest) (*pb.Operation, error) {
	obj, fqn, projectID, err := s.resolvePolicy(ctx, req.GetProject(), req.GetFirewallPolicy())
	if err != nil {
		return nil, err
	}

	rules := []*pb.FirewallPolicyRule{}
	for _, rule := range obj.Rules {
		if rule.Priority != nil && *rule.Priority == *req.Priority {
			continue
		} else {
			rules = append(rules, rule)
		}
	}

	if len(rules) == 0 {
		populateDefaultRules(obj)
	} else {
		obj.Rules = rules
	}

	if err := s.storage.Update(ctx, fqn, obj); err != nil {
		return nil, err
	}

	op := &pb.Operation{
		TargetId:      obj.Id,
		TargetLink:    obj.SelfLink,
		OperationType: PtrTo("compute.networkFirewallPolicy.removeRule"),
		User:          PtrTo("user@example.com"),
	}
	return s.startGlobalLRO(ctx, projectID, op, func() (proto.Message, error) {
		return obj, nil
	})
}

type networkFirewallPolicyName struct {
	Project *projects.ProjectData
	Name    string
}

func (n *networkFirewallPolicyName) String() string {
	return "projects/" + n.Project.ID + "/global/firewallPolicies/" + n.Name
}

func (s *MockService) parseNetworkFirewallPolicyName(name string) (*networkFirewallPolicyName, error) {
	tokens := strings.Split(name, "/")

	if len(tokens) == 5 && tokens[0] == "projects" && tokens[2] == "global" && tokens[3] == "firewallPolicies" {
		projectObj, err := s.Projects.GetProjectByIDOrNumber(tokens[1])
		if err != nil {
			return nil, err
		}
		return s.newNetworkFirewallPolicyName(projectObj, tokens[4])
	}
	return nil, status.Errorf(codes.InvalidArgument, "name %q is not valid", name)
}

func (s *MockService) newNetworkFirewallPolicyName(projectObj *projects.ProjectData, name string) (*networkFirewallPolicyName, error) {
	return &networkFirewallPolicyName{
		Project: projectObj,
		Name:    name,
	}, nil
}
