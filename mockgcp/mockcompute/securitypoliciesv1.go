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
	"sort"
	"strings"

	"github.com/GoogleCloudPlatform/k8s-config-connector/mockgcp/common/projects"
	pb "github.com/GoogleCloudPlatform/k8s-config-connector/mockgcp/generated/mockgcp/cloud/compute/v1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/mockgcp/pkg/storage"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/proto"
)

type SecurityPoliciesV1 struct {
	*MockService
	pb.UnimplementedSecurityPoliciesServer
}

func (s *SecurityPoliciesV1) Get(ctx context.Context, req *pb.GetSecurityPolicyRequest) (*pb.SecurityPolicy, error) {
	reqName := "projects/" + req.GetProject() + "/global/securityPolicies/" + req.GetSecurityPolicy()
	name, err := s.parseSecurityPolicyName(reqName)
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

func (s *SecurityPoliciesV1) Insert(ctx context.Context, req *pb.InsertSecurityPolicyRequest) (*pb.Operation, error) {
	reqName := "projects/" + req.GetProject() + "/global/securityPolicies/" + req.GetSecurityPolicyResource().GetName()
	name, err := s.parseSecurityPolicyName(reqName)
	if err != nil {
		return nil, err
	}

	fqn := name.String()
	id := s.generateID()

	obj := proto.CloneOf(req.GetSecurityPolicyResource())
	obj.SelfLink = PtrTo(BuildComputeSelfLink(ctx, fqn))
	obj.CreationTimestamp = PtrTo(s.nowString())
	obj.Id = &id
	obj.Kind = PtrTo("compute#securityPolicy")

	if obj.Fingerprint == nil {
		obj.Fingerprint = PtrTo("abcdef0123A=")
	}
	if obj.LabelFingerprint == nil {
		obj.LabelFingerprint = PtrTo("abcdef0123A=")
	}
	if obj.Type == nil {
		obj.Type = PtrTo("CLOUD_ARMOR")
	}

	for _, rule := range obj.Rules {
		if rule.Kind == nil {
			rule.Kind = PtrTo("compute#securityPolicyRule")
		}
		if rule.Description == nil {
			rule.Description = PtrTo("")
		}
		if rule.Preview == nil {
			rule.Preview = PtrTo(false)
		}
	}

	sort.Slice(obj.Rules, func(i, j int) bool {
		return obj.Rules[i].GetPriority() < obj.Rules[j].GetPriority()
	})

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

func (s *SecurityPoliciesV1) Patch(ctx context.Context, req *pb.PatchSecurityPolicyRequest) (*pb.Operation, error) {
	reqName := "projects/" + req.GetProject() + "/global/securityPolicies/" + req.GetSecurityPolicy()
	name, err := s.parseSecurityPolicyName(reqName)
	if err != nil {
		return nil, err
	}

	fqn := name.String()
	obj := &pb.SecurityPolicy{}
	if err := s.storage.Get(ctx, fqn, obj); err != nil {
		return nil, err
	}

	patchObj := req.GetSecurityPolicyResource()
	if patchObj.AdaptiveProtectionConfig != nil {
		obj.AdaptiveProtectionConfig = patchObj.AdaptiveProtectionConfig
	}
	if patchObj.AdvancedOptionsConfig != nil {
		obj.AdvancedOptionsConfig = patchObj.AdvancedOptionsConfig
	}
	if patchObj.Description != nil {
		obj.Description = patchObj.Description
	}
	if patchObj.RecaptchaOptionsConfig != nil {
		obj.RecaptchaOptionsConfig = patchObj.RecaptchaOptionsConfig
	}
	if patchObj.Rules != nil {
		obj.Rules = patchObj.Rules
	}
	for _, rule := range obj.Rules {
		if rule.Kind == nil {
			rule.Kind = PtrTo("compute#securityPolicyRule")
		}
	}

	obj.Fingerprint = PtrTo("abcdef0123A=")

	if err := s.storage.Update(ctx, fqn, obj); err != nil {
		return nil, err
	}

	op := &pb.Operation{
		TargetId:      obj.Id,
		TargetLink:    obj.SelfLink,
		OperationType: PtrTo("patch"),
		User:          PtrTo("user@example.com"),
	}
	return s.startGlobalLRO(ctx, name.Project.ID, op, func() (proto.Message, error) {
		return obj, nil
	})
}

func (s *SecurityPoliciesV1) Delete(ctx context.Context, req *pb.DeleteSecurityPolicyRequest) (*pb.Operation, error) {
	reqName := "projects/" + req.GetProject() + "/global/securityPolicies/" + req.GetSecurityPolicy()
	name, err := s.parseSecurityPolicyName(reqName)
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
		OperationType: PtrTo("delete"),
		User:          PtrTo("user@example.com"),
	}
	return s.startGlobalLRO(ctx, name.Project.ID, op, func() (proto.Message, error) {
		return deleted, nil
	})
}

func (s *SecurityPoliciesV1) AddRule(ctx context.Context, req *pb.AddRuleSecurityPolicyRequest) (*pb.Operation, error) {
	reqName := "projects/" + req.GetProject() + "/global/securityPolicies/" + req.GetSecurityPolicy()
	name, err := s.parseSecurityPolicyName(reqName)
	if err != nil {
		return nil, err
	}
	fqn := name.String()
	obj := &pb.SecurityPolicy{}
	if err := s.storage.Get(ctx, fqn, obj); err != nil {
		return nil, err
	}

	rule := proto.CloneOf(req.GetSecurityPolicyRuleResource())
	if rule.Kind == nil {
		rule.Kind = PtrTo("compute#securityPolicyRule")
	}
	obj.Rules = append(obj.Rules, rule)

	sort.Slice(obj.Rules, func(i, j int) bool {
		return obj.Rules[i].GetPriority() < obj.Rules[j].GetPriority()
	})

	if err := s.storage.Update(ctx, fqn, obj); err != nil {
		return nil, err
	}

	op := &pb.Operation{
		TargetId:      obj.Id,
		TargetLink:    obj.SelfLink,
		OperationType: PtrTo("AddRule"),
		User:          PtrTo("user@example.com"),
	}
	return s.startGlobalLRO(ctx, name.Project.ID, op, func() (proto.Message, error) {
		return obj, nil
	})
}

func (s *SecurityPoliciesV1) PatchRule(ctx context.Context, req *pb.PatchRuleSecurityPolicyRequest) (*pb.Operation, error) {
	reqName := "projects/" + req.GetProject() + "/global/securityPolicies/" + req.GetSecurityPolicy()
	name, err := s.parseSecurityPolicyName(reqName)
	if err != nil {
		return nil, err
	}
	fqn := name.String()
	obj := &pb.SecurityPolicy{}
	if err := s.storage.Get(ctx, fqn, obj); err != nil {
		return nil, err
	}

	priority := req.GetPriority()
	found := false
	for idx, rule := range obj.Rules {
		if rule.GetPriority() == priority {
			newRule := req.GetSecurityPolicyRuleResource()
			proto.Merge(rule, newRule)
			if newRule.Kind == nil {
				rule.Kind = PtrTo("compute#securityPolicyRule")
			}
			obj.Rules[idx] = rule
			found = true
			break
		}
	}
	if !found {
		return nil, status.Errorf(codes.NotFound, "Rule with priority %d not found", priority)
	}

	sort.Slice(obj.Rules, func(i, j int) bool {
		return obj.Rules[i].GetPriority() < obj.Rules[j].GetPriority()
	})

	if err := s.storage.Update(ctx, fqn, obj); err != nil {
		return nil, err
	}

	op := &pb.Operation{
		TargetId:      obj.Id,
		TargetLink:    obj.SelfLink,
		OperationType: PtrTo("PatchRule"),
		User:          PtrTo("user@example.com"),
	}
	return s.startGlobalLRO(ctx, name.Project.ID, op, func() (proto.Message, error) {
		return obj, nil
	})
}

func (s *SecurityPoliciesV1) RemoveRule(ctx context.Context, req *pb.RemoveRuleSecurityPolicyRequest) (*pb.Operation, error) {
	reqName := "projects/" + req.GetProject() + "/global/securityPolicies/" + req.GetSecurityPolicy()
	name, err := s.parseSecurityPolicyName(reqName)
	if err != nil {
		return nil, err
	}
	fqn := name.String()
	obj := &pb.SecurityPolicy{}
	if err := s.storage.Get(ctx, fqn, obj); err != nil {
		return nil, err
	}

	priority := req.GetPriority()
	found := false
	var newRules []*pb.SecurityPolicyRule
	for _, rule := range obj.Rules {
		if rule.GetPriority() == priority {
			found = true
		} else {
			newRules = append(newRules, rule)
		}
	}
	if !found {
		return nil, status.Errorf(codes.NotFound, "Rule with priority %d not found", priority)
	}
	obj.Rules = newRules

	sort.Slice(obj.Rules, func(i, j int) bool {
		return obj.Rules[i].GetPriority() < obj.Rules[j].GetPriority()
	})

	if err := s.storage.Update(ctx, fqn, obj); err != nil {
		return nil, err
	}

	op := &pb.Operation{
		TargetId:      obj.Id,
		TargetLink:    obj.SelfLink,
		OperationType: PtrTo("RemoveRule"),
		User:          PtrTo("user@example.com"),
	}
	return s.startGlobalLRO(ctx, name.Project.ID, op, func() (proto.Message, error) {
		return obj, nil
	})
}

func (s *SecurityPoliciesV1) List(ctx context.Context, req *pb.ListSecurityPoliciesRequest) (*pb.SecurityPolicyList, error) {
	reqName := "projects/" + req.GetProject() + "/global/securityPolicies/placeholder"
	name, err := s.parseSecurityPolicyName(reqName)
	if err != nil {
		return nil, err
	}

	findPrefix := strings.TrimSuffix(name.String(), "placeholder")

	response := &pb.SecurityPolicyList{}
	response.Id = PtrTo("0123456789")
	response.Kind = PtrTo("compute#securityPolicyList")

	findKind := (&pb.SecurityPolicy{}).ProtoReflect().Descriptor()
	if err := s.storage.List(ctx, findKind, storage.ListOptions{Prefix: findPrefix}, func(obj proto.Message) error {
		policy := obj.(*pb.SecurityPolicy)
		response.Items = append(response.Items, policy)
		return nil
	}); err != nil {
		return nil, err
	}

	return response, nil
}

type securityPolicyName struct {
	Project *projects.ProjectData
	Region  string
	Name    string
}

func (n *securityPolicyName) String() string {
	if n.Region != "" {
		return "projects/" + n.Project.ID + "/regions/" + n.Region + "/securityPolicies/" + n.Name
	}
	return "projects/" + n.Project.ID + "/global/securityPolicies/" + n.Name
}

func (s *MockService) parseSecurityPolicyName(name string) (*securityPolicyName, error) {
	tokens := strings.Split(name, "/")

	if len(tokens) == 5 && tokens[0] == "projects" && tokens[2] == "global" && tokens[3] == "securityPolicies" {
		project, err := s.Projects.GetProjectByID(tokens[1])
		if err != nil {
			return nil, err
		}

		name := &securityPolicyName{
			Project: project,
			Name:    tokens[4],
		}

		return name, nil
	} else if len(tokens) == 6 && tokens[0] == "projects" && tokens[2] == "regions" && tokens[4] == "securityPolicies" {
		project, err := s.Projects.GetProjectByID(tokens[1])
		if err != nil {
			return nil, err
		}

		name := &securityPolicyName{
			Project: project,
			Region:  tokens[3],
			Name:    tokens[5],
		}

		return name, nil
	} else {
		return nil, status.Errorf(codes.InvalidArgument, "name %q is not valid", name)
	}
}

type RegionSecurityPoliciesV1 struct {
	*MockService
	pb.UnimplementedRegionSecurityPoliciesServer
}

func (s *RegionSecurityPoliciesV1) Get(ctx context.Context, req *pb.GetRegionSecurityPolicyRequest) (*pb.SecurityPolicy, error) {
	reqName := "projects/" + req.GetProject() + "/regions/" + req.GetRegion() + "/securityPolicies/" + req.GetSecurityPolicy()
	name, err := s.parseSecurityPolicyName(reqName)
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

func (s *RegionSecurityPoliciesV1) Insert(ctx context.Context, req *pb.InsertRegionSecurityPolicyRequest) (*pb.Operation, error) {
	reqName := "projects/" + req.GetProject() + "/regions/" + req.GetRegion() + "/securityPolicies/" + req.GetSecurityPolicyResource().GetName()
	name, err := s.parseSecurityPolicyName(reqName)
	if err != nil {
		return nil, err
	}

	fqn := name.String()
	id := s.generateID()

	obj := proto.CloneOf(req.GetSecurityPolicyResource())
	obj.SelfLink = PtrTo(BuildComputeSelfLink(ctx, fqn))
	obj.CreationTimestamp = PtrTo(s.nowString())
	obj.Id = &id
	obj.Kind = PtrTo("compute#securityPolicy")
	obj.Region = PtrTo("https://www.googleapis.com/compute/v1/projects/" + name.Project.ID + "/regions/" + name.Region)

	if obj.Fingerprint == nil {
		obj.Fingerprint = PtrTo("abcdef0123A=")
	}
	if obj.LabelFingerprint == nil {
		obj.LabelFingerprint = PtrTo("abcdef0123A=")
	}
	if obj.Type == nil {
		obj.Type = PtrTo("CLOUD_ARMOR")
	}

	for _, rule := range obj.Rules {
		if rule.Kind == nil {
			rule.Kind = PtrTo("compute#securityPolicyRule")
		}
	}

	sort.Slice(obj.Rules, func(i, j int) bool {
		return obj.Rules[i].GetPriority() < obj.Rules[j].GetPriority()
	})

	if err := s.storage.Create(ctx, fqn, obj); err != nil {
		return nil, err
	}

	op := &pb.Operation{
		TargetId:      obj.Id,
		TargetLink:    obj.SelfLink,
		OperationType: PtrTo("insert"),
		User:          PtrTo("user@example.com"),
	}
	return s.startRegionalLRO(ctx, name.Project.ID, name.Region, op, func() (proto.Message, error) {
		return obj, nil
	})
}

func (s *RegionSecurityPoliciesV1) Patch(ctx context.Context, req *pb.PatchRegionSecurityPolicyRequest) (*pb.Operation, error) {
	reqName := "projects/" + req.GetProject() + "/regions/" + req.GetRegion() + "/securityPolicies/" + req.GetSecurityPolicy()
	name, err := s.parseSecurityPolicyName(reqName)
	if err != nil {
		return nil, err
	}

	fqn := name.String()
	obj := &pb.SecurityPolicy{}
	if err := s.storage.Get(ctx, fqn, obj); err != nil {
		return nil, err
	}

	patchObj := req.GetSecurityPolicyResource()
	if patchObj.AdaptiveProtectionConfig != nil {
		obj.AdaptiveProtectionConfig = patchObj.AdaptiveProtectionConfig
	}
	if patchObj.AdvancedOptionsConfig != nil {
		obj.AdvancedOptionsConfig = patchObj.AdvancedOptionsConfig
	}
	if patchObj.Description != nil {
		obj.Description = patchObj.Description
	}
	if patchObj.RecaptchaOptionsConfig != nil {
		obj.RecaptchaOptionsConfig = patchObj.RecaptchaOptionsConfig
	}
	if patchObj.Rules != nil {
		obj.Rules = patchObj.Rules
	}
	for _, rule := range obj.Rules {
		if rule.Kind == nil {
			rule.Kind = PtrTo("compute#securityPolicyRule")
		}
	}

	obj.Fingerprint = PtrTo("abcdef0123A=")

	if err := s.storage.Update(ctx, fqn, obj); err != nil {
		return nil, err
	}

	op := &pb.Operation{
		TargetId:      obj.Id,
		TargetLink:    obj.SelfLink,
		OperationType: PtrTo("patch"),
		User:          PtrTo("user@example.com"),
	}
	return s.startRegionalLRO(ctx, name.Project.ID, name.Region, op, func() (proto.Message, error) {
		return obj, nil
	})
}

func (s *RegionSecurityPoliciesV1) Delete(ctx context.Context, req *pb.DeleteRegionSecurityPolicyRequest) (*pb.Operation, error) {
	reqName := "projects/" + req.GetProject() + "/regions/" + req.GetRegion() + "/securityPolicies/" + req.GetSecurityPolicy()
	name, err := s.parseSecurityPolicyName(reqName)
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
		OperationType: PtrTo("delete"),
		User:          PtrTo("user@example.com"),
	}
	return s.startRegionalLRO(ctx, name.Project.ID, name.Region, op, func() (proto.Message, error) {
		return deleted, nil
	})
}

func (s *RegionSecurityPoliciesV1) List(ctx context.Context, req *pb.ListRegionSecurityPoliciesRequest) (*pb.SecurityPolicyList, error) {
	reqName := "projects/" + req.GetProject() + "/regions/" + req.GetRegion() + "/securityPolicies/placeholder"
	name, err := s.parseSecurityPolicyName(reqName)
	if err != nil {
		return nil, err
	}

	findPrefix := strings.TrimSuffix(name.String(), "placeholder")

	response := &pb.SecurityPolicyList{}
	response.Id = PtrTo("0123456789")
	response.Kind = PtrTo("compute#securityPolicyList")

	findKind := (&pb.SecurityPolicy{}).ProtoReflect().Descriptor()
	if err := s.storage.List(ctx, findKind, storage.ListOptions{Prefix: findPrefix}, func(obj proto.Message) error {
		policy := obj.(*pb.SecurityPolicy)
		response.Items = append(response.Items, policy)
		return nil
	}); err != nil {
		return nil, err
	}

	return response, nil
}
