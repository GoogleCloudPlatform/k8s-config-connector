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

// +tool:mockgcp-support
// proto.service: google.appengine.v1.Firewall
// proto.message: google.appengine.v1.FirewallRule

package mockappengine

import (
	"context"
	"fmt"
	"sort"
	"strconv"
	"strings"

	pb "cloud.google.com/go/appengine/apiv1/appenginepb"
	"github.com/GoogleCloudPlatform/k8s-config-connector/mockgcp/pkg/storage"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/types/known/emptypb"
)

type firewallServer struct {
	*MockService
	pb.UnimplementedFirewallServer
}

func parseFirewallRuleName(name string) (appID string, priority int32, err error) {
	// apps/{appsId}/firewall/ingressRules/{ingressRulesId}
	parts := strings.Split(name, "/")
	if len(parts) != 5 || parts[0] != "apps" || parts[2] != "firewall" || parts[3] != "ingressRules" {
		return "", 0, status.Errorf(codes.InvalidArgument, "invalid firewall rule name: %q", name)
	}
	appID = parts[1]
	p, err := strconv.ParseInt(parts[4], 10, 32)
	if err != nil {
		return "", 0, status.Errorf(codes.InvalidArgument, "invalid firewall rule priority: %q", parts[4])
	}
	return appID, int32(p), nil
}

func parseFirewallRuleParent(parent string) (appID string, err error) {
	parts := strings.Split(parent, "/")
	if len(parts) == 2 && parts[0] == "apps" {
		return parts[1], nil
	}
	if len(parts) == 4 && parts[0] == "apps" && parts[2] == "firewall" && parts[3] == "ingressRules" {
		return parts[1], nil
	}
	return "", status.Errorf(codes.InvalidArgument, "invalid firewall rule parent: %q", parent)
}

func (s *firewallServer) ensureDefaultRule(ctx context.Context, appID string) {
	fqn := fmt.Sprintf("apps/%s/firewall/ingressRules/2147483647", appID)
	existing := &pb.FirewallRule{}
	if err := s.storage.Get(ctx, fqn, existing); err != nil && status.Code(err) == codes.NotFound {
		defaultRule := &pb.FirewallRule{
			Priority:    2147483647,
			Action:      pb.FirewallRule_ALLOW,
			SourceRange: "*",
			Description: "The default action.",
		}
		_ = s.storage.Create(ctx, fqn, defaultRule)
	}
}

func (s *firewallServer) GetIngressRule(ctx context.Context, req *pb.GetIngressRuleRequest) (*pb.FirewallRule, error) {
	appID, priority, err := parseFirewallRuleName(req.Name)
	if err != nil {
		return nil, err
	}
	s.ensureDefaultRule(ctx, appID)
	fqn := fmt.Sprintf("apps/%s/firewall/ingressRules/%d", appID, priority)
	obj := &pb.FirewallRule{}
	if err := s.storage.Get(ctx, fqn, obj); err != nil {
		if status.Code(err) == codes.NotFound {
			return nil, status.Errorf(codes.NotFound, "No rule found at position %d", priority)
		}
		return nil, err
	}
	return obj, nil
}

func (s *firewallServer) CreateIngressRule(ctx context.Context, req *pb.CreateIngressRuleRequest) (*pb.FirewallRule, error) {
	appID, err := parseFirewallRuleParent(req.Parent)
	if err != nil {
		return nil, err
	}
	if req.Rule == nil {
		return nil, status.Errorf(codes.InvalidArgument, "rule is required")
	}
	priority := req.Rule.Priority
	if priority <= 0 {
		return nil, status.Errorf(codes.InvalidArgument, "priority must be greater than 0")
	}
	s.ensureDefaultRule(ctx, appID)
	fqn := fmt.Sprintf("apps/%s/firewall/ingressRules/%d", appID, priority)
	obj := proto.Clone(req.Rule).(*pb.FirewallRule)
	if err := s.storage.Create(ctx, fqn, obj); err != nil {
		if status.Code(err) == codes.AlreadyExists {
			return nil, status.Errorf(codes.AlreadyExists, "Rule already exists at position %d", priority)
		}
		return nil, err
	}
	return obj, nil
}

func (s *firewallServer) UpdateIngressRule(ctx context.Context, req *pb.UpdateIngressRuleRequest) (*pb.FirewallRule, error) {
	appID, priority, err := parseFirewallRuleName(req.Name)
	if err != nil {
		return nil, err
	}
	if req.Rule == nil {
		return nil, status.Errorf(codes.InvalidArgument, "rule is required")
	}
	s.ensureDefaultRule(ctx, appID)
	fqn := fmt.Sprintf("apps/%s/firewall/ingressRules/%d", appID, priority)
	existing := &pb.FirewallRule{}
	if err := s.storage.Get(ctx, fqn, existing); err != nil {
		if status.Code(err) == codes.NotFound {
			return nil, status.Errorf(codes.NotFound, "No rule found at position %d", priority)
		}
		return nil, err
	}

	paths := []string{"action", "description", "source_range"}
	if req.UpdateMask != nil && len(req.UpdateMask.Paths) > 0 {
		paths = req.UpdateMask.Paths
	}

	for _, path := range paths {
		switch path {
		case "action":
			existing.Action = req.Rule.Action
		case "description":
			existing.Description = req.Rule.Description
		case "source_range", "sourceRange":
			existing.SourceRange = req.Rule.SourceRange
		default:
			return nil, status.Errorf(codes.InvalidArgument, "field %q cannot be updated", path)
		}
	}
	existing.Priority = priority

	if err := s.storage.Update(ctx, fqn, existing); err != nil {
		return nil, err
	}
	return existing, nil
}

func (s *firewallServer) DeleteIngressRule(ctx context.Context, req *pb.DeleteIngressRuleRequest) (*emptypb.Empty, error) {
	appID, priority, err := parseFirewallRuleName(req.Name)
	if err != nil {
		return nil, err
	}
	if priority == 2147483647 {
		return nil, status.Errorf(codes.InvalidArgument, "Cannot delete the default rule")
	}
	s.ensureDefaultRule(ctx, appID)
	fqn := fmt.Sprintf("apps/%s/firewall/ingressRules/%d", appID, priority)
	if err := s.storage.Delete(ctx, fqn, &pb.FirewallRule{}); err != nil {
		if status.Code(err) == codes.NotFound {
			return nil, status.Errorf(codes.NotFound, "No rule found at position %d", priority)
		}
		return nil, err
	}
	return &emptypb.Empty{}, nil
}

func (s *firewallServer) ListIngressRules(ctx context.Context, req *pb.ListIngressRulesRequest) (*pb.ListIngressRulesResponse, error) {
	appID, err := parseFirewallRuleParent(req.Parent)
	if err != nil {
		return nil, err
	}
	s.ensureDefaultRule(ctx, appID)

	var rules []*pb.FirewallRule
	prefix := fmt.Sprintf("apps/%s/firewall/ingressRules/", appID)
	findKind := (&pb.FirewallRule{}).ProtoReflect().Descriptor()
	if err := s.storage.List(ctx, findKind, storage.ListOptions{Prefix: prefix}, func(msg proto.Message) error {
		rule := msg.(*pb.FirewallRule)
		rules = append(rules, proto.Clone(rule).(*pb.FirewallRule))
		return nil
	}); err != nil {
		return nil, err
	}

	sort.Slice(rules, func(i, j int) bool {
		return rules[i].Priority < rules[j].Priority
	})

	return &pb.ListIngressRulesResponse{
		IngressRules: rules,
	}, nil
}

func (s *firewallServer) BatchUpdateIngressRules(ctx context.Context, req *pb.BatchUpdateIngressRulesRequest) (*pb.BatchUpdateIngressRulesResponse, error) {
	parts := strings.Split(req.Name, "/")
	if len(parts) != 4 || parts[0] != "apps" || parts[2] != "firewall" || parts[3] != "ingressRules" {
		return nil, status.Errorf(codes.InvalidArgument, "invalid firewall rules name: %q", req.Name)
	}
	appID := parts[1]

	prefix := fmt.Sprintf("apps/%s/firewall/ingressRules/", appID)
	findKind := (&pb.FirewallRule{}).ProtoReflect().Descriptor()
	var toDelete []string
	if err := s.storage.List(ctx, findKind, storage.ListOptions{Prefix: prefix}, func(msg proto.Message) error {
		rule := msg.(*pb.FirewallRule)
		toDelete = append(toDelete, fmt.Sprintf("%s%d", prefix, rule.Priority))
		return nil
	}); err != nil {
		return nil, err
	}
	for _, fqn := range toDelete {
		_ = s.storage.Delete(ctx, fqn, &pb.FirewallRule{})
	}

	var result []*pb.FirewallRule
	for _, rule := range req.IngressRules {
		obj := proto.Clone(rule).(*pb.FirewallRule)
		fqn := fmt.Sprintf("%s%d", prefix, obj.Priority)
		if err := s.storage.Create(ctx, fqn, obj); err != nil {
			return nil, err
		}
		result = append(result, obj)
	}
	s.ensureDefaultRule(ctx, appID)

	return &pb.BatchUpdateIngressRulesResponse{
		IngressRules: result,
	}, nil
}
