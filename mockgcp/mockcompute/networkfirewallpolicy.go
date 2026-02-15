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
// proto.message: google.cloud.compute.v1.FirewallPolicy

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

type NetworkFirewallPoliciesV1 struct {
	*MockService
	pb.UnimplementedNetworkFirewallPoliciesServer
}

func (s *NetworkFirewallPoliciesV1) Get(ctx context.Context, req *pb.GetNetworkFirewallPolicyRequest) (*pb.FirewallPolicy, error) {
	reqName := fmt.Sprintf("projects/%s/global/firewallPolicies/%s", req.GetProject(), req.GetFirewallPolicy())
	name, err := s.parseNetworkFirewallPolicyName(reqName)
	if err != nil {
		return nil, err
	}

	fqn := name.String()

	obj := &pb.FirewallPolicy{}
	if err := s.storage.Get(ctx, fqn, obj); err != nil {
		if status.Code(err) == codes.NotFound {
			return nil, status.Errorf(codes.NotFound, "The resource '%s' was not found", fqn)
		}
		return nil, err
	}

	return obj, nil
}

func (s *NetworkFirewallPoliciesV1) Insert(ctx context.Context, req *pb.InsertNetworkFirewallPolicyRequest) (*pb.Operation, error) {
	reqName := fmt.Sprintf("projects/%s/global/firewallPolicies/%s", req.GetProject(), req.GetFirewallPolicyResource().GetName())
	name, err := s.parseNetworkFirewallPolicyName(reqName)
	if err != nil {
		return nil, err
	}

	fqn := name.String()

	obj := ProtoClone(req.GetFirewallPolicyResource())
	obj.Name = PtrTo(name.Name)
	obj.Kind = PtrTo("compute#firewallPolicy")
	obj.Id = PtrTo(s.generateID())
	obj.CreationTimestamp = PtrTo(s.nowString())
	obj.SelfLink = PtrTo(buildComputeSelfLink(ctx, fqn))
	obj.SelfLinkWithId = PtrTo(fmt.Sprintf("https://www.googleapis.com/compute/v1/projects/%s/global/firewallPolicies/%d", name.Project.ID, obj.GetId()))

	if obj.Description == nil {
		obj.Description = PtrTo("")
	}

	// Use default rules
	if obj.Rules == nil {
		populateDefaultNetworkFirewallPolicyRules(obj)
	}

	ruleTupleCount := int32(0)
	for _, rule := range obj.Rules {
		ruleTupleCount += rule.GetRuleTupleCount()
	}
	obj.RuleTupleCount = PtrTo(ruleTupleCount)

	if obj.Fingerprint == nil {
		obj.Fingerprint = PtrTo(computeFingerprint(obj))
	}

	if err := s.storage.Create(ctx, fqn, obj); err != nil {
		return nil, err
	}

	op := &pb.Operation{
		OperationType: PtrTo("compute.networkFirewallPolicy.insert"),
		TargetId:      obj.Id,
		TargetLink:    obj.SelfLink,
		User:          PtrTo("user@example.com"),
	}
	return s.computeOperations.startGlobalLRO(ctx, name.Project.ID, op, func() (proto.Message, error) {
		return obj, nil
	})
}

func (s *NetworkFirewallPoliciesV1) Patch(ctx context.Context, req *pb.PatchNetworkFirewallPolicyRequest) (*pb.Operation, error) {
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

	proto.Merge(obj, req.GetFirewallPolicyResource())

	if err := s.storage.Update(ctx, fqn, obj); err != nil {
		return nil, err
	}

	op := &pb.Operation{
		OperationType: PtrTo("compute.networkFirewallPolicy.patch"),
		TargetId:      obj.Id,
		TargetLink:    obj.SelfLink,
		User:          PtrTo("user@example.com"),
	}
	return s.computeOperations.startGlobalLRO(ctx, name.Project.ID, op, func() (proto.Message, error) {
		return obj, nil
	})
}

func (s *NetworkFirewallPoliciesV1) Delete(ctx context.Context, req *pb.DeleteNetworkFirewallPolicyRequest) (*pb.Operation, error) {
	reqName := fmt.Sprintf("projects/%s/global/firewallPolicies/%s", req.GetProject(), req.GetFirewallPolicy())
	name, err := s.parseNetworkFirewallPolicyName(reqName)
	if err != nil {
		return nil, err
	}

	fqn := name.String()

	deleted := &pb.FirewallPolicy{}
	if err := s.storage.Delete(ctx, fqn, deleted); err != nil {
		return nil, err
	}

	op := &pb.Operation{
		OperationType: PtrTo("compute.networkFirewallPolicy.delete"),
		TargetId:      deleted.Id,
		TargetLink:    deleted.SelfLink,
		User:          PtrTo("user@example.com"),
	}
	return s.computeOperations.startGlobalLRO(ctx, name.Project.ID, op, func() (proto.Message, error) {
		return &emptypb.Empty{}, nil
	})
}

func (s *NetworkFirewallPoliciesV1) List(ctx context.Context, req *pb.ListNetworkFirewallPoliciesRequest) (*pb.FirewallPolicyList, error) {
	projectID := req.GetProject()
	prefix := fmt.Sprintf("projects/%s/global/firewallPolicies/", projectID)

	var items []*pb.FirewallPolicy
	kind := (&pb.FirewallPolicy{}).ProtoReflect().Descriptor()
	if err := s.storage.List(ctx, kind, storage.ListOptions{Prefix: prefix}, func(msg proto.Message) error {
		item := msg.(*pb.FirewallPolicy)
		items = append(items, item)
		return nil
	}); err != nil {
		return nil, err
	}

	return &pb.FirewallPolicyList{
		Id:    PtrTo("000000000000000000000"),
		Items: items,
		Kind:  PtrTo("compute#firewallPolicyList"),
	}, nil
}

func (s *NetworkFirewallPoliciesV1) GetRule(ctx context.Context, req *pb.GetRuleNetworkFirewallPolicyRequest) (*pb.FirewallPolicyRule, error) {
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

	for _, r := range obj.GetRules() {
		if r.Priority != nil && *r.Priority == *req.Priority {
			return r, nil
		}
	}

	return nil, status.Errorf(codes.NotFound, "Rule with priority %d not found", *req.Priority)
}

func (s *NetworkFirewallPoliciesV1) AddRule(ctx context.Context, req *pb.AddRuleNetworkFirewallPolicyRequest) (*pb.Operation, error) {
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

	r := req.GetFirewallPolicyRuleResource()
	mockFieldValuesForRule(r)

	// Check for duplicates? Simple append for now.
	obj.Rules = append(obj.Rules, r)

	if err := s.storage.Update(ctx, fqn, obj); err != nil {
		return nil, err
	}

	op := &pb.Operation{
		OperationType: PtrTo("compute.networkFirewallPolicy.addRule"),
		TargetId:      obj.Id,
		TargetLink:    obj.SelfLink,
		User:          PtrTo("user@example.com"),
	}
	return s.computeOperations.startGlobalLRO(ctx, name.Project.ID, op, func() (proto.Message, error) {
		return nil, nil
	})
}

func (s *NetworkFirewallPoliciesV1) PatchRule(ctx context.Context, req *pb.PatchRuleNetworkFirewallPolicyRequest) (*pb.Operation, error) {
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

	var rules []*pb.FirewallPolicyRule
	found := false
	for _, rule := range obj.Rules {
		if rule.Priority != nil && *rule.Priority == *req.Priority {
			r := req.GetFirewallPolicyRuleResource()
			r.Priority = PtrTo(*rule.Priority)
			mockFieldValuesForRule(r)
			rules = append(rules, r)
			found = true
		} else {
			rules = append(rules, rule)
		}
	}
	if !found {
		return nil, status.Errorf(codes.NotFound, "Rule with priority %d not found", *req.Priority)
	}

	obj.Rules = rules
	if err := s.storage.Update(ctx, fqn, obj); err != nil {
		return nil, err
	}

	op := &pb.Operation{
		OperationType: PtrTo("compute.networkFirewallPolicy.patchRule"),
		TargetId:      obj.Id,
		TargetLink:    obj.SelfLink,
		User:          PtrTo("user@example.com"),
	}
	return s.computeOperations.startGlobalLRO(ctx, name.Project.ID, op, func() (proto.Message, error) {
		return nil, nil
	})
}

func (s *NetworkFirewallPoliciesV1) RemoveRule(ctx context.Context, req *pb.RemoveRuleNetworkFirewallPolicyRequest) (*pb.Operation, error) {
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

	var rules []*pb.FirewallPolicyRule
	for _, rule := range obj.Rules {
		if rule.Priority != nil && *rule.Priority == *req.Priority {
			continue
		}
		rules = append(rules, rule)
	}

	// Check if rule list is empty and populate defaults if needed?
	// For RemoveRule on a specific priority, if it leaves empty it implies user deleted all.
	// But usually default rules (priority 2147483647) should be there unless user deleted that too.
	// However, the helper populateDefaultNetworkFirewallPolicyRules adds 65535 implied rule?

	obj.Rules = rules

	if err := s.storage.Update(ctx, fqn, obj); err != nil {
		return nil, err
	}

	op := &pb.Operation{
		OperationType: PtrTo("compute.networkFirewallPolicy.removeRule"),
		TargetId:      obj.Id,
		TargetLink:    obj.SelfLink,
		User:          PtrTo("user@example.com"),
	}
	return s.computeOperations.startGlobalLRO(ctx, name.Project.ID, op, func() (proto.Message, error) {
		return nil, nil
	})
}

type networkFirewallPolicyName struct {
	Project *projects.ProjectData
	Name    string
}

func (n *networkFirewallPolicyName) String() string {
	return fmt.Sprintf("projects/%s/global/firewallPolicies/%s", n.Project.ID, n.Name)
}

// parseNetworkFirewallPolicyName parses a string into a networkFirewallPolicyName.
// The expected form is `projects/{project}/global/firewallPolicies/{firewall_policy}`.
func (s *MockService) parseNetworkFirewallPolicyName(name string) (*networkFirewallPolicyName, error) {
	tokens := strings.Split(name, "/")

	if len(tokens) == 5 && tokens[0] == "projects" && tokens[2] == "global" && tokens[3] == "firewallPolicies" {
		project, err := s.Projects.GetProjectByID(tokens[1])
		if err != nil {
			return nil, err
		}

		return &networkFirewallPolicyName{
			Project: project,
			Name:    tokens[4],
		}, nil
	}
	return nil, status.Errorf(codes.InvalidArgument, "name %q is not valid", name)
}

func populateDefaultNetworkFirewallPolicyRules(obj *pb.FirewallPolicy) {
	// Network Firewall Policies default rules
	obj.Rules = []*pb.FirewallPolicyRule{
		{
			Priority:      PtrTo(int32(2147483644)),
			Action:        PtrTo("goto_next"),
			Description:   PtrTo("default egress rule"),
			Direction:     PtrTo("EGRESS"),
			EnableLogging: PtrTo(false),
			Kind:          PtrTo("compute#firewallPolicyRule"),
			Match: &pb.FirewallPolicyRuleMatcher{
				DestIpRanges: []string{"::/0"},
				Layer4Configs: []*pb.FirewallPolicyRuleMatcherLayer4Config{
					{
						IpProtocol: PtrTo("all"),
					},
				},
			},
			RuleTupleCount: PtrTo(int32(2)),
		},

		{
			Priority:      PtrTo(int32(2147483645)),
			Action:        PtrTo("goto_next"),
			Description:   PtrTo("default ingress rule"),
			Direction:     PtrTo("INGRESS"),
			EnableLogging: PtrTo(false),
			Kind:          PtrTo("compute#firewallPolicyRule"),
			Match: &pb.FirewallPolicyRuleMatcher{
				SrcIpRanges: []string{"::/0"},
				Layer4Configs: []*pb.FirewallPolicyRuleMatcherLayer4Config{
					{
						IpProtocol: PtrTo("all"),
					},
				},
			},
			RuleTupleCount: PtrTo(int32(2)),
		},

		{
			Priority:      PtrTo(int32(2147483646)),
			Action:        PtrTo("goto_next"),
			Description:   PtrTo("default egress rule"),
			Direction:     PtrTo("EGRESS"),
			EnableLogging: PtrTo(false),
			Kind:          PtrTo("compute#firewallPolicyRule"),
			Match: &pb.FirewallPolicyRuleMatcher{
				DestIpRanges: []string{"0.0.0.0/0"},
				Layer4Configs: []*pb.FirewallPolicyRuleMatcherLayer4Config{
					{
						IpProtocol: PtrTo("all"),
					},
				},
			},
			RuleTupleCount: PtrTo(int32(2)),
		},

		{
			Priority:      PtrTo(int32(2147483647)),
			Action:        PtrTo("goto_next"),
			Description:   PtrTo("default ingress rule"),
			Direction:     PtrTo("INGRESS"),
			EnableLogging: PtrTo(false),
			Kind:          PtrTo("compute#firewallPolicyRule"),
			Match: &pb.FirewallPolicyRuleMatcher{
				SrcIpRanges: []string{"0.0.0.0/0"},
				Layer4Configs: []*pb.FirewallPolicyRuleMatcherLayer4Config{
					{
						IpProtocol: PtrTo("all"),
					},
				},
			},
			RuleTupleCount: PtrTo(int32(2)),
		},
	}
}
