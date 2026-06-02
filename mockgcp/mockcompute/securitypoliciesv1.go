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
	"strings"

	"github.com/GoogleCloudPlatform/k8s-config-connector/mockgcp/common/projects"
	pb "github.com/GoogleCloudPlatform/k8s-config-connector/mockgcp/generated/mockgcp/cloud/compute/v1"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/proto"
)

type GlobalSecurityPoliciesV1 struct {
	*MockService
	pb.UnimplementedSecurityPoliciesServer
}

type RegionalSecurityPoliciesV1 struct {
	*MockService
	pb.UnimplementedRegionSecurityPoliciesServer
}

type securityPolicyName struct {
	Project *projects.ProjectData
	Region  string // empty for global
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
		return &securityPolicyName{
			Project: project,
			Name:    tokens[4],
		}, nil
	} else if len(tokens) == 6 && tokens[0] == "projects" && tokens[2] == "regions" && tokens[4] == "securityPolicies" {
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
	return nil, status.Errorf(codes.InvalidArgument, "name %q is not valid", name)
}

// Global SecurityPolicies
func (s *GlobalSecurityPoliciesV1) Get(ctx context.Context, req *pb.GetSecurityPolicyRequest) (*pb.SecurityPolicy, error) {
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

func (s *GlobalSecurityPoliciesV1) Insert(ctx context.Context, req *pb.InsertSecurityPolicyRequest) (*pb.Operation, error) {
	reqName := "projects/" + req.GetProject() + "/global/securityPolicies/" + req.GetSecurityPolicyResource().GetName()
	name, err := s.parseSecurityPolicyName(reqName)
	if err != nil {
		return nil, err
	}

	fqn := name.String()
	id := s.generateID()

	obj := proto.Clone(req.GetSecurityPolicyResource()).(*pb.SecurityPolicy)
	obj.SelfLink = PtrTo(BuildComputeSelfLink(ctx, fqn))
	obj.CreationTimestamp = PtrTo(s.nowString())
	obj.Id = &id
	obj.Kind = PtrTo("compute#securityPolicy")
	if obj.Fingerprint == nil {
		obj.Fingerprint = PtrTo(computeFingerprint(obj))
	}

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

func (s *GlobalSecurityPoliciesV1) Patch(ctx context.Context, req *pb.PatchSecurityPolicyRequest) (*pb.Operation, error) {
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

	proto.Merge(obj, req.GetSecurityPolicyResource())
	obj.Fingerprint = PtrTo(computeFingerprint(obj))

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

func (s *GlobalSecurityPoliciesV1) Delete(ctx context.Context, req *pb.DeleteSecurityPolicyRequest) (*pb.Operation, error) {
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

// Regional SecurityPolicies
func (s *RegionalSecurityPoliciesV1) Get(ctx context.Context, req *pb.GetRegionSecurityPolicyRequest) (*pb.SecurityPolicy, error) {
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

func (s *RegionalSecurityPoliciesV1) Insert(ctx context.Context, req *pb.InsertRegionSecurityPolicyRequest) (*pb.Operation, error) {
	reqName := "projects/" + req.GetProject() + "/regions/" + req.GetRegion() + "/securityPolicies/" + req.GetSecurityPolicyResource().GetName()
	name, err := s.parseSecurityPolicyName(reqName)
	if err != nil {
		return nil, err
	}

	fqn := name.String()
	id := s.generateID()

	obj := proto.Clone(req.GetSecurityPolicyResource()).(*pb.SecurityPolicy)
	obj.SelfLink = PtrTo(BuildComputeSelfLink(ctx, fqn))
	obj.CreationTimestamp = PtrTo(s.nowString())
	obj.Id = &id
	obj.Kind = PtrTo("compute#securityPolicy")
	if obj.Fingerprint == nil {
		obj.Fingerprint = PtrTo(computeFingerprint(obj))
	}

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

func (s *RegionalSecurityPoliciesV1) Patch(ctx context.Context, req *pb.PatchRegionSecurityPolicyRequest) (*pb.Operation, error) {
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

	proto.Merge(obj, req.GetSecurityPolicyResource())
	obj.Fingerprint = PtrTo(computeFingerprint(obj))

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

func (s *RegionalSecurityPoliciesV1) Delete(ctx context.Context, req *pb.DeleteRegionSecurityPolicyRequest) (*pb.Operation, error) {
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
