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
	"time"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/types/known/emptypb"
	apierrors "k8s.io/apimachinery/pkg/api/errors"

	"github.com/GoogleCloudPlatform/k8s-config-connector/mockgcp/common/projects"
	pb "github.com/GoogleCloudPlatform/k8s-config-connector/mockgcp/generated/google/cloud/dns/v1"
)

type responsePoliciesService struct {
	*MockService
	pb.UnimplementedResponsePoliciesServerServer
}

type responsePolicyName struct {
	Project *projects.ProjectData
	Name    string
}

func (n *responsePolicyName) String() string {
	return "projects/" + n.Project.ID + "/responsePolicies/" + n.Name
}

func (s *MockService) parseResponsePolicyName(name string) (*responsePolicyName, error) {
	tokens := strings.Split(name, "/")
	if len(tokens) == 4 && tokens[0] == "projects" && tokens[2] == "responsePolicies" {
		project, err := s.Projects.GetProjectByIDOrNumber(tokens[1])
		if err != nil {
			return nil, err
		}
		name := &responsePolicyName{
			Project: project,
			Name:    tokens[3],
		}
		return name, nil
	} else {
		return nil, status.Errorf(codes.InvalidArgument, "name %q is not valid", name)
	}
}

func (s *responsePoliciesService) GetResponsePolicy(ctx context.Context, req *pb.GetResponsePolicyRequest) (*pb.ResponsePolicy, error) {
	name, err := s.parseResponsePolicyName("projects/" + req.GetProject() + "/responsePolicies/" + req.GetName())
	if err != nil {
		return nil, err
	}

	fqn := name.String()

	obj := &pb.ResponsePolicy{}
	if err := s.storage.Get(ctx, fqn, obj); err != nil {
		if apierrors.IsNotFound(err) {
			return nil, status.Errorf(codes.NotFound, "The 'parameters.responsePolicy' resource named '%s' does not exist.", name.Name)
		}
		return nil, err
	}

	return obj, nil
}

func (s *responsePoliciesService) CreateResponsePolicy(ctx context.Context, req *pb.CreateResponsePolicyRequest) (*pb.ResponsePolicy, error) {
	name, err := s.parseResponsePolicyName("projects/" + req.GetProject() + "/responsePolicies/" + req.GetResponsePolicy().GetResponsePolicyName())
	if err != nil {
		return nil, err
	}

	fqn := name.String()

	obj := proto.Clone(req.ResponsePolicy).(*pb.ResponsePolicy)

	obj.Id = PtrTo[int64](int64(time.Now().UnixNano()))
	obj.Kind = PtrTo("dns#responsePolicy")

	for i := range obj.GkeClusters {
		obj.GkeClusters[i].Kind = PtrTo("dns#responsePolicyGKECluster")
	}
	for i := range obj.Networks {
		obj.Networks[i].Kind = PtrTo("dns#responsePolicyNetwork")
	}

	if err := s.storage.Create(ctx, fqn, obj); err != nil {
		return nil, err
	}

	return obj, nil
}

func (s *responsePoliciesService) UpdateResponsePolicy(ctx context.Context, req *pb.UpdateResponsePolicyRequest) (*pb.ResponsePoliciesUpdateResponse, error) {
	name, err := s.parseResponsePolicyName("projects/" + req.GetProject() + "/responsePolicies/" + req.GetName())
	if err != nil {
		return nil, err
	}

	fqn := name.String()

	var existing pb.ResponsePolicy
	if err := s.storage.Get(ctx, fqn, &existing); err != nil {
		if apierrors.IsNotFound(err) {
			return nil, status.Errorf(codes.NotFound, "The 'parameters.responsePolicy' resource named '%s' does not exist.", name.Name)
		}
		return nil, err
	}

	updated := proto.Clone(req.ResponsePolicy).(*pb.ResponsePolicy)

	updated.Id = existing.Id
	updated.Kind = existing.Kind
	updated.ResponsePolicyName = existing.ResponsePolicyName

	for i := range updated.GkeClusters {
		updated.GkeClusters[i].Kind = PtrTo("dns#responsePolicyGKECluster")
	}
	for i := range updated.Networks {
		updated.Networks[i].Kind = PtrTo("dns#responsePolicyNetwork")
	}

	if err := s.storage.Update(ctx, fqn, updated); err != nil {
		return nil, err
	}

	return &pb.ResponsePoliciesUpdateResponse{
		ResponsePolicy: updated,
	}, nil
}

func (s *responsePoliciesService) PatchResponsePolicy(ctx context.Context, req *pb.PatchResponsePolicyRequest) (*pb.ResponsePoliciesPatchResponse, error) {
	name, err := s.parseResponsePolicyName("projects/" + req.GetProject() + "/responsePolicies/" + req.GetName())
	if err != nil {
		return nil, err
	}

	fqn := name.String()

	var existing pb.ResponsePolicy
	if err := s.storage.Get(ctx, fqn, &existing); err != nil {
		if apierrors.IsNotFound(err) {
			return nil, status.Errorf(codes.NotFound, "The 'parameters.responsePolicy' resource named '%s' does not exist.", name.Name)
		}
		return nil, err
	}

	updated := proto.Clone(req.ResponsePolicy).(*pb.ResponsePolicy)

	updated.Id = existing.Id
	updated.Kind = existing.Kind
	updated.ResponsePolicyName = existing.ResponsePolicyName

	for i := range updated.GkeClusters {
		updated.GkeClusters[i].Kind = PtrTo("dns#responsePolicyGKECluster")
	}
	for i := range updated.Networks {
		updated.Networks[i].Kind = PtrTo("dns#responsePolicyNetwork")
	}

	if err := s.storage.Update(ctx, fqn, updated); err != nil {
		return nil, err
	}

	return &pb.ResponsePoliciesPatchResponse{
		ResponsePolicy: updated,
	}, nil
}

func (s *responsePoliciesService) DeleteResponsePolicy(ctx context.Context, req *pb.DeleteResponsePolicyRequest) (*emptypb.Empty, error) {
	name, err := s.parseResponsePolicyName("projects/" + req.GetProject() + "/responsePolicies/" + req.GetName())
	if err != nil {
		return nil, err
	}

	fqn := name.String()

	var existing pb.ResponsePolicy
	if err := s.storage.Get(ctx, fqn, &existing); err != nil {
		if apierrors.IsNotFound(err) {
			return nil, status.Errorf(codes.NotFound, "The 'parameters.responsePolicy' resource named '%s' does not exist.", name.Name)
		}
		return nil, err
	}

	if err := s.storage.Delete(ctx, fqn, &existing); err != nil {
		return nil, err
	}

	return &emptypb.Empty{}, nil
}

func (s *responsePoliciesService) ListResponsePolicies(ctx context.Context, req *pb.ListResponsePoliciesRequest) (*pb.ResponsePoliciesListResponse, error) {
	return &pb.ResponsePoliciesListResponse{}, nil
}
