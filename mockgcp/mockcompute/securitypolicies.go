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
	"strings"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/proto"

	"github.com/GoogleCloudPlatform/k8s-config-connector/mockgcp/common/projects"
	pb "github.com/GoogleCloudPlatform/k8s-config-connector/mockgcp/generated/mockgcp/cloud/compute/v1"
)

// SecurityPoliciesV1 implements SecurityPoliciesServer
type SecurityPoliciesV1 struct {
	*MockService
	pb.UnimplementedSecurityPoliciesServer
}

func (s *SecurityPoliciesV1) Get(ctx context.Context, req *pb.GetSecurityPolicyRequest) (*pb.SecurityPolicy, error) {
	reqName := fmt.Sprintf("projects/%s/global/securityPolicies/%s", req.GetProject(), req.GetSecurityPolicy())
	name, err := s.parseSecurityPolicyName(reqName)
	if err != nil {
		return nil, err
	}
	fqn := name.String()

	obj := &pb.SecurityPolicy{}
	if err := s.storage.Get(ctx, fqn, obj); err != nil {
		if status.Code(err) == codes.NotFound {
			return nil, status.Errorf(codes.NotFound, "The resource '%s' was not found", name)
		}
		return nil, err
	}

	return obj, nil
}

func (s *SecurityPoliciesV1) Insert(ctx context.Context, req *pb.InsertSecurityPolicyRequest) (*pb.Operation, error) {
	reqName := fmt.Sprintf("projects/%s/global/securityPolicies/%s", req.GetProject(), req.GetSecurityPolicyResource().GetName())
	name, err := s.parseSecurityPolicyName(reqName)
	if err != nil {
		return nil, err
	}
	fqn := name.String()
	id := s.generateID()

	obj := proto.Clone(req.GetSecurityPolicyResource()).(*pb.SecurityPolicy)
	obj.Id = proto.Uint64(id)
	obj.SelfLink = PtrTo(BuildComputeSelfLink(ctx, fqn))
	obj.Kind = PtrTo("compute#securityPolicy")
	obj.CreationTimestamp = PtrTo(s.nowString())
	obj.Fingerprint = PtrTo("abcdef0123A=")

	if err := s.storage.Create(ctx, fqn, obj); err != nil {
		return nil, err
	}

	op := &pb.Operation{
		OperationType: PtrTo("insertSecurityPolicy"),
		TargetId:      obj.Id,
		TargetLink:    obj.SelfLink,
		User:          PtrTo("user@example.com"),
	}
	return s.computeOperations.startGlobalLRO(ctx, name.Project.ID, op, func() (proto.Message, error) {
		return obj, nil
	})
}

func (s *SecurityPoliciesV1) Patch(ctx context.Context, req *pb.PatchSecurityPolicyRequest) (*pb.Operation, error) {
	reqName := fmt.Sprintf("projects/%s/global/securityPolicies/%s", req.GetProject(), req.GetSecurityPolicy())
	name, err := s.parseSecurityPolicyName(reqName)
	if err != nil {
		return nil, err
	}
	fqn := name.String()

	obj := &pb.SecurityPolicy{}
	if err := s.storage.Get(ctx, fqn, obj); err != nil {
		return nil, err
	}

	// Update fields
	newObj := req.GetSecurityPolicyResource()
	if newObj.AdaptiveProtectionConfig != nil {
		obj.AdaptiveProtectionConfig = newObj.AdaptiveProtectionConfig
	}
	if newObj.AdvancedOptionsConfig != nil {
		obj.AdvancedOptionsConfig = newObj.AdvancedOptionsConfig
	}
	if newObj.Description != nil {
		obj.Description = newObj.Description
	}
	if newObj.RecaptchaOptionsConfig != nil {
		obj.RecaptchaOptionsConfig = newObj.RecaptchaOptionsConfig
	}
	if newObj.Rules != nil {
		obj.Rules = newObj.Rules
	}
	if newObj.Type != nil {
		obj.Type = newObj.Type
	}
	obj.Fingerprint = PtrTo("updated_fingerprint_A=")

	if err := s.storage.Update(ctx, fqn, obj); err != nil {
		return nil, err
	}

	op := &pb.Operation{
		OperationType: PtrTo("patchSecurityPolicy"),
		TargetId:      obj.Id,
		TargetLink:    obj.SelfLink,
		User:          PtrTo("user@example.com"),
	}
	return s.computeOperations.startGlobalLRO(ctx, name.Project.ID, op, func() (proto.Message, error) {
		return obj, nil
	})
}

func (s *SecurityPoliciesV1) Delete(ctx context.Context, req *pb.DeleteSecurityPolicyRequest) (*pb.Operation, error) {
	reqName := fmt.Sprintf("projects/%s/global/securityPolicies/%s", req.GetProject(), req.GetSecurityPolicy())
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
		OperationType: PtrTo("deleteSecurityPolicy"),
		TargetId:      deleted.Id,
		TargetLink:    deleted.SelfLink,
		User:          PtrTo("user@example.com"),
	}
	return s.computeOperations.startGlobalLRO(ctx, name.Project.ID, op, func() (proto.Message, error) {
		return deleted, nil
	})
}

func (s *SecurityPoliciesV1) AddRule(ctx context.Context, req *pb.AddRuleSecurityPolicyRequest) (*pb.Operation, error) {
	reqName := fmt.Sprintf("projects/%s/global/securityPolicies/%s", req.GetProject(), req.GetSecurityPolicy())
	name, err := s.parseSecurityPolicyName(reqName)
	if err != nil {
		return nil, err
	}
	fqn := name.String()

	obj := &pb.SecurityPolicy{}
	if err := s.storage.Get(ctx, fqn, obj); err != nil {
		return nil, err
	}

	rule := req.GetSecurityPolicyRuleResource()
	obj.Rules = append(obj.Rules, rule)
	obj.Fingerprint = PtrTo("updated_fingerprint_A=")

	if err := s.storage.Update(ctx, fqn, obj); err != nil {
		return nil, err
	}

	op := &pb.Operation{
		OperationType: PtrTo("addRuleSecurityPolicy"),
		TargetId:      obj.Id,
		TargetLink:    obj.SelfLink,
		User:          PtrTo("user@example.com"),
	}
	return s.computeOperations.startGlobalLRO(ctx, name.Project.ID, op, func() (proto.Message, error) {
		return obj, nil
	})
}

func (s *SecurityPoliciesV1) PatchRule(ctx context.Context, req *pb.PatchRuleSecurityPolicyRequest) (*pb.Operation, error) {
	reqName := fmt.Sprintf("projects/%s/global/securityPolicies/%s", req.GetProject(), req.GetSecurityPolicy())
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
	newRule := req.GetSecurityPolicyRuleResource()
	found := false
	for i, rule := range obj.Rules {
		if rule.GetPriority() == priority {
			obj.Rules[i] = newRule
			found = true
			break
		}
	}
	if !found {
		obj.Rules = append(obj.Rules, newRule)
	}
	obj.Fingerprint = PtrTo("updated_fingerprint_A=")

	if err := s.storage.Update(ctx, fqn, obj); err != nil {
		return nil, err
	}

	op := &pb.Operation{
		OperationType: PtrTo("patchRuleSecurityPolicy"),
		TargetId:      obj.Id,
		TargetLink:    obj.SelfLink,
		User:          PtrTo("user@example.com"),
	}
	return s.computeOperations.startGlobalLRO(ctx, name.Project.ID, op, func() (proto.Message, error) {
		return obj, nil
	})
}

func (s *SecurityPoliciesV1) RemoveRule(ctx context.Context, req *pb.RemoveRuleSecurityPolicyRequest) (*pb.Operation, error) {
	reqName := fmt.Sprintf("projects/%s/global/securityPolicies/%s", req.GetProject(), req.GetSecurityPolicy())
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
	var rules []*pb.SecurityPolicyRule
	for _, rule := range obj.Rules {
		if rule.GetPriority() != priority {
			rules = append(rules, rule)
		}
	}
	obj.Rules = rules
	obj.Fingerprint = PtrTo("updated_fingerprint_A=")

	if err := s.storage.Update(ctx, fqn, obj); err != nil {
		return nil, err
	}

	op := &pb.Operation{
		OperationType: PtrTo("removeRuleSecurityPolicy"),
		TargetId:      obj.Id,
		TargetLink:    obj.SelfLink,
		User:          PtrTo("user@example.com"),
	}
	return s.computeOperations.startGlobalLRO(ctx, name.Project.ID, op, func() (proto.Message, error) {
		return obj, nil
	})
}

// RegionSecurityPoliciesV1 implements RegionSecurityPoliciesServer
type RegionSecurityPoliciesV1 struct {
	*MockService
	pb.UnimplementedRegionSecurityPoliciesServer
}

func (s *RegionSecurityPoliciesV1) Get(ctx context.Context, req *pb.GetRegionSecurityPolicyRequest) (*pb.SecurityPolicy, error) {
	reqName := fmt.Sprintf("projects/%s/regions/%s/securityPolicies/%s", req.GetProject(), req.GetRegion(), req.GetSecurityPolicy())
	name, err := s.parseSecurityPolicyName(reqName)
	if err != nil {
		return nil, err
	}
	fqn := name.String()

	obj := &pb.SecurityPolicy{}
	if err := s.storage.Get(ctx, fqn, obj); err != nil {
		if status.Code(err) == codes.NotFound {
			return nil, status.Errorf(codes.NotFound, "The resource '%s' was not found", name)
		}
		return nil, err
	}

	return obj, nil
}

func (s *RegionSecurityPoliciesV1) Insert(ctx context.Context, req *pb.InsertRegionSecurityPolicyRequest) (*pb.Operation, error) {
	reqName := fmt.Sprintf("projects/%s/regions/%s/securityPolicies/%s", req.GetProject(), req.GetRegion(), req.GetSecurityPolicyResource().GetName())
	name, err := s.parseSecurityPolicyName(reqName)
	if err != nil {
		return nil, err
	}
	fqn := name.String()
	id := s.generateID()

	obj := proto.Clone(req.GetSecurityPolicyResource()).(*pb.SecurityPolicy)
	obj.Id = proto.Uint64(id)
	obj.SelfLink = PtrTo(BuildComputeSelfLink(ctx, fqn))
	obj.Kind = PtrTo("compute#securityPolicy")
	obj.CreationTimestamp = PtrTo(s.nowString())
	obj.Fingerprint = PtrTo("abcdef0123RegionA=")
	obj.Region = PtrTo(fmt.Sprintf("https://www.googleapis.com/compute/v1/projects/%s/regions/%s", name.Project.ID, name.Region))

	if err := s.storage.Create(ctx, fqn, obj); err != nil {
		return nil, err
	}

	op := &pb.Operation{
		OperationType: PtrTo("insertRegionSecurityPolicy"),
		TargetId:      obj.Id,
		TargetLink:    obj.SelfLink,
		User:          PtrTo("user@example.com"),
	}
	return s.computeOperations.startRegionalLRO(ctx, name.Project.ID, name.Region, op, func() (proto.Message, error) {
		return obj, nil
	})
}

func (s *RegionSecurityPoliciesV1) Patch(ctx context.Context, req *pb.PatchRegionSecurityPolicyRequest) (*pb.Operation, error) {
	reqName := fmt.Sprintf("projects/%s/regions/%s/securityPolicies/%s", req.GetProject(), req.GetRegion(), req.GetSecurityPolicy())
	name, err := s.parseSecurityPolicyName(reqName)
	if err != nil {
		return nil, err
	}
	fqn := name.String()

	obj := &pb.SecurityPolicy{}
	if err := s.storage.Get(ctx, fqn, obj); err != nil {
		return nil, err
	}

	// Update fields
	newObj := req.GetSecurityPolicyResource()
	if newObj.AdaptiveProtectionConfig != nil {
		obj.AdaptiveProtectionConfig = newObj.AdaptiveProtectionConfig
	}
	if newObj.AdvancedOptionsConfig != nil {
		obj.AdvancedOptionsConfig = newObj.AdvancedOptionsConfig
	}
	if newObj.Description != nil {
		obj.Description = newObj.Description
	}
	if newObj.RecaptchaOptionsConfig != nil {
		obj.RecaptchaOptionsConfig = newObj.RecaptchaOptionsConfig
	}
	if newObj.Rules != nil {
		obj.Rules = newObj.Rules
	}
	if newObj.Type != nil {
		obj.Type = newObj.Type
	}
	obj.Fingerprint = PtrTo("updated_fingerprint_RegionA=")

	if err := s.storage.Update(ctx, fqn, obj); err != nil {
		return nil, err
	}

	op := &pb.Operation{
		OperationType: PtrTo("patchRegionSecurityPolicy"),
		TargetId:      obj.Id,
		TargetLink:    obj.SelfLink,
		User:          PtrTo("user@example.com"),
	}
	return s.computeOperations.startRegionalLRO(ctx, name.Project.ID, name.Region, op, func() (proto.Message, error) {
		return obj, nil
	})
}

func (s *RegionSecurityPoliciesV1) Delete(ctx context.Context, req *pb.DeleteRegionSecurityPolicyRequest) (*pb.Operation, error) {
	reqName := fmt.Sprintf("projects/%s/regions/%s/securityPolicies/%s", req.GetProject(), req.GetRegion(), req.GetSecurityPolicy())
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
		OperationType: PtrTo("deleteRegionSecurityPolicy"),
		TargetId:      deleted.Id,
		TargetLink:    deleted.SelfLink,
		User:          PtrTo("user@example.com"),
	}
	return s.computeOperations.startRegionalLRO(ctx, name.Project.ID, name.Region, op, func() (proto.Message, error) {
		return deleted, nil
	})
}

func (s *RegionSecurityPoliciesV1) AddRule(ctx context.Context, req *pb.AddRuleRegionSecurityPolicyRequest) (*pb.Operation, error) {
	reqName := fmt.Sprintf("projects/%s/regions/%s/securityPolicies/%s", req.GetProject(), req.GetRegion(), req.GetSecurityPolicy())
	name, err := s.parseSecurityPolicyName(reqName)
	if err != nil {
		return nil, err
	}
	fqn := name.String()

	obj := &pb.SecurityPolicy{}
	if err := s.storage.Get(ctx, fqn, obj); err != nil {
		return nil, err
	}

	rule := req.GetSecurityPolicyRuleResource()
	obj.Rules = append(obj.Rules, rule)
	obj.Fingerprint = PtrTo("updated_fingerprint_RegionA=")

	if err := s.storage.Update(ctx, fqn, obj); err != nil {
		return nil, err
	}

	op := &pb.Operation{
		OperationType: PtrTo("addRuleRegionSecurityPolicy"),
		TargetId:      obj.Id,
		TargetLink:    obj.SelfLink,
		User:          PtrTo("user@example.com"),
	}
	return s.computeOperations.startRegionalLRO(ctx, name.Project.ID, name.Region, op, func() (proto.Message, error) {
		return obj, nil
	})
}

func (s *RegionSecurityPoliciesV1) PatchRule(ctx context.Context, req *pb.PatchRuleRegionSecurityPolicyRequest) (*pb.Operation, error) {
	reqName := fmt.Sprintf("projects/%s/regions/%s/securityPolicies/%s", req.GetProject(), req.GetRegion(), req.GetSecurityPolicy())
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
	newRule := req.GetSecurityPolicyRuleResource()
	found := false
	for i, rule := range obj.Rules {
		if rule.GetPriority() == priority {
			obj.Rules[i] = newRule
			found = true
			break
		}
	}
	if !found {
		obj.Rules = append(obj.Rules, newRule)
	}
	obj.Fingerprint = PtrTo("updated_fingerprint_RegionA=")

	if err := s.storage.Update(ctx, fqn, obj); err != nil {
		return nil, err
	}

	op := &pb.Operation{
		OperationType: PtrTo("patchRuleRegionSecurityPolicy"),
		TargetId:      obj.Id,
		TargetLink:    obj.SelfLink,
		User:          PtrTo("user@example.com"),
	}
	return s.computeOperations.startRegionalLRO(ctx, name.Project.ID, name.Region, op, func() (proto.Message, error) {
		return obj, nil
	})
}

func (s *RegionSecurityPoliciesV1) RemoveRule(ctx context.Context, req *pb.RemoveRuleRegionSecurityPolicyRequest) (*pb.Operation, error) {
	reqName := fmt.Sprintf("projects/%s/regions/%s/securityPolicies/%s", req.GetProject(), req.GetRegion(), req.GetSecurityPolicy())
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
	var rules []*pb.SecurityPolicyRule
	for _, rule := range obj.Rules {
		if rule.GetPriority() != priority {
			rules = append(rules, rule)
		}
	}
	obj.Rules = rules
	obj.Fingerprint = PtrTo("updated_fingerprint_RegionA=")

	if err := s.storage.Update(ctx, fqn, obj); err != nil {
		return nil, err
	}

	op := &pb.Operation{
		OperationType: PtrTo("removeRuleRegionSecurityPolicy"),
		TargetId:      obj.Id,
		TargetLink:    obj.SelfLink,
		User:          PtrTo("user@example.com"),
	}
	return s.computeOperations.startRegionalLRO(ctx, name.Project.ID, name.Region, op, func() (proto.Message, error) {
		return obj, nil
	})
}

type securityPolicyName struct {
	Project *projects.ProjectData
	Region  string
	Name    string
}

func (n *securityPolicyName) String() string {
	if n.Region == "" {
		return fmt.Sprintf("projects/%s/global/securityPolicies/%s", n.Project.ID, n.Name)
	}
	return fmt.Sprintf("projects/%s/regions/%s/securityPolicies/%s", n.Project.ID, n.Region, n.Name)
}

func (s *MockService) parseSecurityPolicyName(name string) (*securityPolicyName, error) {
	tokens := strings.Split(name, "/")
	if len(tokens) == 5 && tokens[0] == "projects" && tokens[2] == "global" && tokens[3] == "securityPolicies" {
		project, err := s.Projects.GetProjectByID(tokens[1])
		if err != nil {
			return nil, err
		}
		return &securityPolicyName{
			Project: project,
			Name:    tokens[4],
		}, nil
	}
	if len(tokens) == 6 && tokens[0] == "projects" && tokens[2] == "regions" && tokens[4] == "securityPolicies" {
		project, err := s.Projects.GetProjectByID(tokens[1])
		if err != nil {
			return nil, err
		}
		return &securityPolicyName{
			Project: project,
			Region:  tokens[3],
			Name:    tokens[5],
		}, nil
	}
	return nil, status.Errorf(codes.InvalidArgument, "invalid security policy name %q", name)
}
