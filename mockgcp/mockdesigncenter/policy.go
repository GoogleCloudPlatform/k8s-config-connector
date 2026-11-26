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

package mockdesigncenter

import (
	"context"
	"fmt"
	"time"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/types/known/emptypb"
	"google.golang.org/protobuf/types/known/timestamppb"

	pb "cloud.google.com/go/designcenter/apiv1/designcenterpb"
	"github.com/GoogleCloudPlatform/k8s-config-connector/mockgcp/common/fields"
	"github.com/GoogleCloudPlatform/k8s-config-connector/mockgcp/pkg/storage"
)

func (s *DesignCenterServer) GetPolicy(ctx context.Context, req *pb.GetPolicyRequest) (*pb.Policy, error) {
	name, err := s.parsePolicyName(req.Name)
	if err != nil {
		return nil, err
	}

	fqn := name.String()

	obj := &pb.Policy{}
	if err := s.storage.Get(ctx, fqn, obj); err != nil {
		if status.Code(err) == codes.NotFound {
			return nil, status.Errorf(codes.NotFound, "Resource '%s' was not found", fqn)
		}
		return nil, err
	}

	return s.populatePolicy(obj), nil
}

func (s *DesignCenterServer) CreatePolicy(ctx context.Context, req *pb.CreatePolicyRequest) (*pb.Policy, error) {
	parent, err := s.parseApplicationTemplateName(req.Parent)
	if err != nil {
		return nil, err
	}

	// Reject if parent application template does not exist
	parentAppTemplate := &pb.ApplicationTemplate{}
	if err := s.storage.Get(ctx, parent.String(), parentAppTemplate); err != nil {
		if status.Code(err) == codes.NotFound {
			return nil, status.Errorf(codes.NotFound, "Resource '%s' was not found", parent.String())
		}
		return nil, err
	}

	if req.PolicyId == "" {
		return nil, status.Errorf(codes.InvalidArgument, "policyId is required")
	}

	if req.Policy == nil {
		return nil, status.Errorf(codes.InvalidArgument, "policy is required")
	}

	policyID := req.PolicyId
	fqn := fmt.Sprintf("%s/policies/%s", parent.String(), policyID)

	name, err := s.parsePolicyName(fqn)
	if err != nil {
		return nil, err
	}

	now := time.Now()
	obj := proto.Clone(req.Policy).(*pb.Policy)
	obj.Name = name.String()
	obj.CreateTime = timestamppb.New(now)
	obj.UpdateTime = timestamppb.New(now)

	obj = s.populatePolicy(obj)

	if err := s.storage.Create(ctx, fqn, obj); err != nil {
		return nil, err
	}

	return obj, nil
}

func (s *DesignCenterServer) UpdatePolicy(ctx context.Context, req *pb.UpdatePolicyRequest) (*pb.Policy, error) {
	if req.Policy == nil {
		return nil, status.Errorf(codes.InvalidArgument, "policy is required")
	}

	name, err := s.parsePolicyName(req.Policy.Name)
	if err != nil {
		return nil, err
	}

	fqn := name.String()

	existing := &pb.Policy{}
	if err := s.storage.Get(ctx, fqn, existing); err != nil {
		if status.Code(err) == codes.NotFound {
			return nil, status.Errorf(codes.NotFound, "Resource '%s' was not found", fqn)
		}
		return nil, err
	}

	updated := proto.Clone(existing).(*pb.Policy)
	now := time.Now()
	updated.UpdateTime = timestamppb.New(now)

	paths := req.GetUpdateMask().GetPaths()
	if len(paths) == 0 {
		return nil, status.Errorf(codes.InvalidArgument, "update_mask is required")
	}

	if err := fields.UpdateByFieldMask(updated, req.Policy, paths); err != nil {
		return nil, err
	}

	updated = s.populatePolicy(updated)

	if err := s.storage.Update(ctx, fqn, updated); err != nil {
		return nil, err
	}

	return updated, nil
}

func (s *DesignCenterServer) DeletePolicy(ctx context.Context, req *pb.DeletePolicyRequest) (*emptypb.Empty, error) {
	name, err := s.parsePolicyName(req.Name)
	if err != nil {
		return nil, err
	}

	fqn := name.String()

	deleted := &pb.Policy{}
	if err := s.storage.Delete(ctx, fqn, deleted); err != nil {
		if status.Code(err) == codes.NotFound {
			return nil, status.Errorf(codes.NotFound, "Resource '%s' was not found", fqn)
		}
		return nil, err
	}

	return &emptypb.Empty{}, nil
}

func (s *DesignCenterServer) ListPolicies(ctx context.Context, req *pb.ListPoliciesRequest) (*pb.ListPoliciesResponse, error) {
	parent, err := s.parseApplicationTemplateName(req.Parent)
	if err != nil {
		return nil, err
	}

	prefix := parent.String() + "/policies/"
	response := &pb.ListPoliciesResponse{}

	findOpts := storage.ListOptions{
		Prefix: prefix,
	}
	policyKind := (&pb.Policy{}).ProtoReflect().Descriptor()
	if err := s.storage.List(ctx, policyKind, findOpts, func(item proto.Message) error {
		obj := item.(*pb.Policy)
		response.Policies = append(response.Policies, s.populatePolicy(obj))
		return nil
	}); err != nil {
		return nil, err
	}

	return response, nil
}

func (s *DesignCenterServer) populatePolicy(obj *pb.Policy) *pb.Policy {
	return obj
}
