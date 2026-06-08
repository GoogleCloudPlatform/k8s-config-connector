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

package mockdns

import (
	"context"
	"strings"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/types/known/emptypb"
	apierrors "k8s.io/apimachinery/pkg/api/errors"

	"github.com/GoogleCloudPlatform/k8s-config-connector/mockgcp/common/projects"
	pb "github.com/GoogleCloudPlatform/k8s-config-connector/mockgcp/generated/google/cloud/dns/v1"
)

type policiesService struct {
	*MockService
	pb.UnimplementedPoliciesServerServer
}

type policyName struct {
	Project *projects.ProjectData
	Name    string
}

func (n *policyName) String() string {
	return "projects/" + n.Project.ID + "/policies/" + n.Name
}

func (s *MockService) parsePolicyName(name string) (*policyName, error) {
	tokens := strings.Split(name, "/")
	if len(tokens) == 4 && tokens[0] == "projects" && tokens[2] == "policies" {
		project, err := s.Projects.GetProjectByIDOrNumber(tokens[1])
		if err != nil {
			return nil, err
		}
		name := &policyName{
			Project: project,
			Name:    tokens[3],
		}
		return name, nil
	} else {
		return nil, status.Errorf(codes.InvalidArgument, "name %q is not valid", name)
	}
}

func (s *policiesService) GetPolicy(ctx context.Context, req *pb.GetPolicyRequest) (*pb.Policy, error) {
	name, err := s.parsePolicyName("projects/" + req.GetProject() + "/policies/" + req.GetName())
	if err != nil {
		return nil, err
	}

	fqn := name.String()

	obj := &pb.Policy{}
	if err := s.storage.Get(ctx, fqn, obj); err != nil {
		if apierrors.IsNotFound(err) {
			return nil, status.Errorf(codes.NotFound, "policy %q not found", name)
		}
		return nil, err
	}

	return obj, nil
}

func (s *policiesService) CreatePolicy(ctx context.Context, req *pb.CreatePolicyRequest) (*pb.Policy, error) {
	name, err := s.parsePolicyName("projects/" + req.GetProject() + "/policies/" + req.GetPolicy().GetName())
	if err != nil {
		return nil, err
	}

	fqn := name.String()

	obj := proto.Clone(req.Policy).(*pb.Policy)

	obj.Id = PtrTo[uint64](1234567890)
	obj.Kind = PtrTo("dns#policy")

	if obj.AlternativeNameServerConfig != nil {
		obj.AlternativeNameServerConfig.Kind = PtrTo("dns#policyAlternativeNameServerConfig")
		for i := range obj.AlternativeNameServerConfig.TargetNameServers {
			obj.AlternativeNameServerConfig.TargetNameServers[i].Kind = PtrTo("dns#policyTargetNameServer")
		}
	}
	for i := range obj.Networks {
		obj.Networks[i].Kind = PtrTo("dns#policyNetwork")
	}

	if err := s.storage.Create(ctx, fqn, obj); err != nil {
		return nil, err
	}

	return obj, nil
}

func (s *policiesService) UpdatePolicy(ctx context.Context, req *pb.UpdatePolicyRequest) (*pb.PoliciesUpdateResponse, error) {
	name, err := s.parsePolicyName("projects/" + req.GetProject() + "/policies/" + req.GetName())
	if err != nil {
		return nil, err
	}

	fqn := name.String()

	var existing pb.Policy
	if err := s.storage.Get(ctx, fqn, &existing); err != nil {
		if apierrors.IsNotFound(err) {
			return nil, status.Errorf(codes.NotFound, "policy %q not found", name)
		}
		return nil, err
	}

	updated := proto.Clone(req.Policy).(*pb.Policy)

	updated.Id = existing.Id
	updated.Kind = existing.Kind
	updated.Name = existing.Name

	if updated.AlternativeNameServerConfig != nil {
		updated.AlternativeNameServerConfig.Kind = PtrTo("dns#policyAlternativeNameServerConfig")
		for i := range updated.AlternativeNameServerConfig.TargetNameServers {
			updated.AlternativeNameServerConfig.TargetNameServers[i].Kind = PtrTo("dns#policyTargetNameServer")
		}
	}
	for i := range updated.Networks {
		updated.Networks[i].Kind = PtrTo("dns#policyNetwork")
	}

	if err := s.storage.Update(ctx, fqn, updated); err != nil {
		return nil, err
	}

	return &pb.PoliciesUpdateResponse{
		Policy: updated,
	}, nil
}

func (s *policiesService) PatchPolicy(ctx context.Context, req *pb.PatchPolicyRequest) (*pb.PoliciesPatchResponse, error) {
	name, err := s.parsePolicyName("projects/" + req.GetProject() + "/policies/" + req.GetName())
	if err != nil {
		return nil, err
	}

	fqn := name.String()

	var existing pb.Policy
	if err := s.storage.Get(ctx, fqn, &existing); err != nil {
		if apierrors.IsNotFound(err) {
			return nil, status.Errorf(codes.NotFound, "policy %q not found", name)
		}
		return nil, err
	}

	updated := proto.Clone(req.Policy).(*pb.Policy)

	updated.Id = existing.Id
	updated.Kind = existing.Kind
	updated.Name = existing.Name

	if updated.AlternativeNameServerConfig != nil {
		updated.AlternativeNameServerConfig.Kind = PtrTo("dns#policyAlternativeNameServerConfig")
		for i := range updated.AlternativeNameServerConfig.TargetNameServers {
			updated.AlternativeNameServerConfig.TargetNameServers[i].Kind = PtrTo("dns#policyTargetNameServer")
		}
	}
	for i := range updated.Networks {
		updated.Networks[i].Kind = PtrTo("dns#policyNetwork")
	}

	if err := s.storage.Update(ctx, fqn, updated); err != nil {
		return nil, err
	}

	return &pb.PoliciesPatchResponse{
		Policy: updated,
	}, nil
}

func (s *policiesService) DeletePolicy(ctx context.Context, req *pb.DeletePolicyRequest) (*emptypb.Empty, error) {
	name, err := s.parsePolicyName("projects/" + req.GetProject() + "/policies/" + req.GetName())
	if err != nil {
		return nil, err
	}

	fqn := name.String()

	var existing pb.Policy
	if err := s.storage.Get(ctx, fqn, &existing); err != nil {
		if apierrors.IsNotFound(err) {
			return nil, status.Errorf(codes.NotFound, "policy %q not found", name)
		}
		return nil, err
	}

	if err := s.storage.Delete(ctx, fqn, &existing); err != nil {
		return nil, err
	}

	return &emptypb.Empty{}, nil
}

func (s *policiesService) ListPolicies(ctx context.Context, req *pb.ListPoliciesRequest) (*pb.PoliciesListResponse, error) {
	return &pb.PoliciesListResponse{}, nil
}
