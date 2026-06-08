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
	"github.com/GoogleCloudPlatform/k8s-config-connector/mockgcp/pkg/storage"
)

type policiesService struct {
	*MockService
	pb.UnimplementedPoliciesServerServer
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
		for _, target := range obj.AlternativeNameServerConfig.TargetNameServers {
			target.Kind = PtrTo("dns#policyAlternativeNameServerConfigTargetNameServer")
		}
	}
	for _, net := range obj.Networks {
		net.Kind = PtrTo("dns#policyNetwork")
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

	// These fields are output only and cannot be changed.
	updated.Id = existing.Id
	updated.Kind = existing.Kind

	if updated.AlternativeNameServerConfig != nil {
		updated.AlternativeNameServerConfig.Kind = PtrTo("dns#policyAlternativeNameServerConfig")
		for _, target := range updated.AlternativeNameServerConfig.TargetNameServers {
			target.Kind = PtrTo("dns#policyAlternativeNameServerConfigTargetNameServer")
		}
	}
	for _, net := range updated.Networks {
		net.Kind = PtrTo("dns#policyNetwork")
	}

	if err := s.storage.Update(ctx, fqn, updated); err != nil {
		return nil, err
	}

	return &pb.PoliciesUpdateResponse{
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

	updated := proto.Clone(&existing).(*pb.Policy)

	// Merge fields from request
	reqPolicy := req.GetPolicy()
	if reqPolicy.Description != nil {
		updated.Description = reqPolicy.Description
	}
	if reqPolicy.AlternativeNameServerConfig != nil {
		updated.AlternativeNameServerConfig = reqPolicy.AlternativeNameServerConfig
		updated.AlternativeNameServerConfig.Kind = PtrTo("dns#policyAlternativeNameServerConfig")
		for _, target := range updated.AlternativeNameServerConfig.TargetNameServers {
			target.Kind = PtrTo("dns#policyAlternativeNameServerConfigTargetNameServer")
		}
	}
	if reqPolicy.EnableInboundForwarding != nil {
		updated.EnableInboundForwarding = reqPolicy.EnableInboundForwarding
	}
	if reqPolicy.EnableLogging != nil {
		updated.EnableLogging = reqPolicy.EnableLogging
	}
	if reqPolicy.Networks != nil {
		updated.Networks = reqPolicy.Networks
		for _, net := range updated.Networks {
			net.Kind = PtrTo("dns#policyNetwork")
		}
	}

	if err := s.storage.Update(ctx, fqn, updated); err != nil {
		return nil, err
	}

	return &pb.PoliciesPatchResponse{
		Policy: updated,
	}, nil
}

func (s *policiesService) ListPolicies(ctx context.Context, req *pb.ListPoliciesRequest) (*pb.PoliciesListResponse, error) {
	name, err := s.parsePolicyName("projects/" + req.GetProject() + "/policies/dummy")
	if err != nil {
		return nil, err
	}

	prefix := strings.TrimSuffix(name.String(), "dummy")

	policyKind := (&pb.Policy{}).ProtoReflect().Descriptor()

	items := make([]*pb.Policy, 0)
	if err := s.storage.List(ctx, policyKind, storage.ListOptions{Prefix: prefix}, func(item proto.Message) error {
		p := item.(*pb.Policy)
		items = append(items, p)
		return nil
	}); err != nil {
		return nil, err
	}

	response := &pb.PoliciesListResponse{
		Policies: items,
		Kind:     PtrTo("dns#policiesListResponse"),
	}

	if req.MaxResults != nil {
		maxResults := req.GetMaxResults()
		if len(items) > int(maxResults) {
			response.Policies = items[:maxResults]
			response.NextPageToken = PtrTo("mock-pagination-token")
		}
	}

	return response, nil
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
	}
	return nil, status.Errorf(codes.InvalidArgument, "invalid policy name %q", name)
}
