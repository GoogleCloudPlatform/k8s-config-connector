// Copyright 2024 Google LLC
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

	pb "github.com/GoogleCloudPlatform/k8s-config-connector/mockgcp/generated/mockgcp/cloud/compute/v1"
)

type FirewallPoliciesV1 struct {
	*MockService
	pb.UnimplementedFirewallPoliciesServer
}

func (s *FirewallPoliciesV1) Get(ctx context.Context, req *pb.GetFirewallPolicyRequest) (*pb.FirewallPolicy, error) {
	reqName := "locations/global/firewallPolicies/" + req.GetFirewallPolicy()
	name, err := s.parseFirewallPolicyName(reqName)
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

func (s *FirewallPoliciesV1) Insert(ctx context.Context, req *pb.InsertFirewallPolicyRequest) (*pb.Operation, error) {
	id := s.generateID()
	policyId := strconv.FormatUint(id, 10)

	reqName := "locations/global/firewallPolicies/" + policyId
	policyName, err := s.parseFirewallPolicyName(reqName)
	if err != nil {
		return nil, err
	}

	fqn := policyName.String()

	obj := proto.Clone(req.GetFirewallPolicyResource()).(*pb.FirewallPolicy)
	obj.SelfLink = PtrTo(buildComputeSelfLink(ctx, policyName.String()))
	obj.SelfLinkWithId = PtrTo(buildComputeSelfLink(ctx, policyName.String()) + "/" + policyId)
	obj.Parent = PtrTo(req.ParentId)
	obj.RuleTupleCount = PtrTo(int32(8))
	obj.Id = PtrTo(id)
	obj.Name = PtrTo(policyId)
	obj.CreationTimestamp = PtrTo(s.nowString())
	obj.Kind = PtrTo("compute#firewallPolicy")
	obj.DisplayName = PtrTo(obj.GetShortName())

	if obj.Fingerprint == nil {
		obj.Fingerprint = PtrTo(computeFingerprint(obj))
	}

	// Use default rules
	if obj.Rules == nil {
		populateDefaultRules(obj)
	}

	if err := s.storage.Create(ctx, fqn, obj); err != nil {
		return nil, err
	}

	op := &pb.Operation{
		TargetId:      obj.Id,
		TargetLink:    obj.SelfLink,
		OperationType: PtrTo("createFirewallPolicy"),
		User:          PtrTo("user@example.com"),
	}
	return s.startGlobalOrganizationLRO(ctx, op, func() (proto.Message, error) {
		return obj, nil
	})
}

func (s *FirewallPoliciesV1) Patch(ctx context.Context, req *pb.PatchFirewallPolicyRequest) (*pb.Operation, error) {
	reqName := "locations/global/firewallPolicies/" + req.GetFirewallPolicy()

	name, err := s.parseFirewallPolicyName(reqName)
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
		TargetId:      obj.Id,
		TargetLink:    obj.SelfLink,
		OperationType: PtrTo("updateFirewallPolicy"),
		User:          PtrTo("user@example.com"),
		// patch operation finished super fast
		Progress: PtrTo(int32(100)),
		Status:   PtrTo(pb.Operation_DONE),
		EndTime:  PtrTo(s.nowString()),
	}
	return s.startGlobalOrganizationLRO(ctx, op, func() (proto.Message, error) {
		return obj, nil
	})
}
func (s *FirewallPoliciesV1) Delete(ctx context.Context, req *pb.DeleteFirewallPolicyRequest) (*pb.Operation, error) {
	reqName := "locations/global/firewallPolicies/" + req.GetFirewallPolicy()
	name, err := s.parseFirewallPolicyName(reqName)
	if err != nil {
		return nil, err
	}

	fqn := name.String()

	deleted := &pb.FirewallPolicy{}
	if err := s.storage.Delete(ctx, fqn, deleted); err != nil {
		return nil, err
	}

	op := &pb.Operation{
		TargetId:      deleted.Id,
		TargetLink:    deleted.SelfLink,
		OperationType: PtrTo("deleteFirewallPolicy"),
		User:          PtrTo("user@example.com"),
	}
	return s.startGlobalOrganizationLRO(ctx, op, func() (proto.Message, error) {
		return deleted, nil
	})
}

func (s *FirewallPoliciesV1) GetRule(ctx context.Context, req *pb.GetRuleFirewallPolicyRequest) (*pb.FirewallPolicyRule, error) {
	reqName := "locations/global/firewallPolicies/" + req.GetFirewallPolicy()
	name, err := s.parseFirewallPolicyName(reqName)
	if err != nil {
		return nil, err
	}

	fqn := name.String()

	obj := &pb.FirewallPolicy{}
	if err := s.storage.Get(ctx, fqn, obj); err != nil {
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

func (s *FirewallPoliciesV1) AddRule(ctx context.Context, req *pb.AddRuleFirewallPolicyRequest) (*pb.Operation, error) {
	reqName := "locations/global/firewallPolicies/" + req.GetFirewallPolicy()
	name, err := s.parseFirewallPolicyName(reqName)
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

	obj.Rules = []*pb.FirewallPolicyRule{r}
	if err := s.storage.Update(ctx, fqn, obj); err != nil {
		return nil, err
	}

	op := &pb.Operation{
		TargetId:      obj.Id,
		TargetLink:    obj.SelfLink,
		OperationType: PtrTo("addFirewallRuleToFirewallPolicy"),
		User:          PtrTo("user@example.com"),
	}
	return s.startGlobalOrganizationLRO(ctx, op, func() (proto.Message, error) {
		return obj, nil
	})
}

func (s *FirewallPoliciesV1) PatchRule(ctx context.Context, req *pb.PatchRuleFirewallPolicyRequest) (*pb.Operation, error) {
	reqName := "locations/global/firewallPolicies/" + req.GetFirewallPolicy()

	name, err := s.parseFirewallPolicyName(reqName)
	if err != nil {
		return nil, err
	}

	fqn := name.String()
	obj := &pb.FirewallPolicy{}
	if err := s.storage.Get(ctx, fqn, obj); err != nil {
		return nil, err
	}

	rules := []*pb.FirewallPolicyRule{}
	for _, rule := range obj.Rules {
		if rule.Priority != nil && *rule.Priority == *req.Priority {
			// update the rule
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
		OperationType: PtrTo("patchFirewallRuleInFirewallPolicy"),
		User:          PtrTo("user@example.com"),
	}
	return s.startGlobalOrganizationLRO(ctx, op, func() (proto.Message, error) {
		return obj, nil
	})
}

func (s *FirewallPoliciesV1) RemoveRule(ctx context.Context, req *pb.RemoveRuleFirewallPolicyRequest) (*pb.Operation, error) {
	reqName := "locations/global/firewallPolicies/" + req.GetFirewallPolicy()
	name, err := s.parseFirewallPolicyName(reqName)
	if err != nil {
		return nil, err
	}

	fqn := name.String()

	obj := &pb.FirewallPolicy{}
	if err := s.storage.Get(ctx, fqn, obj); err != nil {
		return nil, err
	}

	rules := []*pb.FirewallPolicyRule{}
	for _, rule := range obj.Rules {
		if rule.Priority != nil && *rule.Priority == *req.Priority {
			// remove the rule
			continue
		} else {
			rules = append(rules, rule)
		}
	}

	if len(rules) == 0 {
		// When the target policy has no rules, i.e. all the custom rules are deleted,
		// we update the policy to add default rules to it.
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
		OperationType: PtrTo("removeFirewallRuleFromFirewallPolicy"),
		User:          PtrTo("user@example.com"),
	}
	return s.startGlobalOrganizationLRO(ctx, op, func() (proto.Message, error) {
		return obj, nil
	})
}

type firewallPolicyName struct {
	Name string
}

func (n *firewallPolicyName) String() string {
	return "locations/global/firewallPolicies/" + n.Name
}

// parseFirewallPolicyName parses a string into a firewallPolicyName.
// The expected form is `locations/global/firewallPolicies/*`.
func (s *MockService) parseFirewallPolicyName(name string) (*firewallPolicyName, error) {
	tokens := strings.Split(name, "/")

	if len(tokens) == 4 && tokens[0] == "locations" && tokens[1] == "global" && tokens[2] == "firewallPolicies" {
		name := &firewallPolicyName{
			Name: tokens[3],
		}
		return name, nil
	} else {
		return nil, status.Errorf(codes.InvalidArgument, "name %q is not valid", name)
	}
}

func populateDefaultRules(obj *pb.FirewallPolicy) {
	obj.Rules = []*pb.FirewallPolicyRule{
		{
			Action:        PtrTo("goto_next"),
			Description:   PtrTo("default egress rule ipv6"),
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
			Priority:       PtrTo(int32(2147483644)),
			RuleTupleCount: PtrTo(int32(2)),
		},
		{
			Action:        PtrTo("goto_next"),
			Description:   PtrTo("default ingress rule ipv6"),
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
			Priority:       PtrTo(int32(2147483645)),
			RuleTupleCount: PtrTo(int32(2)),
		},
		{
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
			Priority:       PtrTo(int32(2147483646)),
			RuleTupleCount: PtrTo(int32(2)),
		},
		{
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
			Priority:       PtrTo(int32(2147483647)),
			RuleTupleCount: PtrTo(int32(2)),
		},
	}
}

func mockFieldValuesForRule(r *pb.FirewallPolicyRule) {
	// RuleTupleCount is output only, calculation of the complexity of a single firewall policy rule.
	// Manually set different ruleTupleCount to match the realGCP log
	if r.TargetResources != nil {
		r.RuleTupleCount = PtrTo(int32(4))
	} else {
		r.RuleTupleCount = PtrTo(int32(2))
	}
	r.Kind = PtrTo("compute#firewallPolicyRule")
	if r.Description == nil {
		r.Description = PtrTo("")
	}
}
