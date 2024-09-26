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
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/proto"
	"strconv"
	"strings"

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
	policyId := req.GetFirewallPolicyResource().GetName()
	id := s.generateID()
	if policyId == "" {
		policyId = strconv.FormatUint(id, 10)
	}
	reqName := "locations/global/firewallPolicies/" + policyId
	policyName, err := s.parseFirewallPolicyName(reqName)
	if err != nil {
		return nil, err
	}

	fqn := policyName.String()

	obj := proto.Clone(req.GetFirewallPolicyResource()).(*pb.FirewallPolicy)
	obj.SelfLink = PtrTo("https://www.googleapis.com/compute/v1/" + policyName.String())
	obj.SelfLinkWithId = PtrTo("https://www.googleapis.com/compute/v1/" + policyName.String() + "/" + policyId)
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
		EndTime:  PtrTo("2024-04-01T12:34:56.123456Z"),
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

	if len(tokens) == 4 && tokens[2] == "firewallPolicies" {
		name := &firewallPolicyName{
			Name: tokens[3],
		}
		return name, nil
	} else {
		return nil, status.Errorf(codes.InvalidArgument, "name %q is not valid", name)
	}
}
