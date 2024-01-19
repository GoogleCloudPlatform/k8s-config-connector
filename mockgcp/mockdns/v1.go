// Copyright 2023 Google LLC
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

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/proto"
	apierrors "k8s.io/apimachinery/pkg/api/errors"

	pb "github.com/GoogleCloudPlatform/k8s-config-connector/mockgcp/generated/mockgcp/cloud/dns/v1beta2"
)

type DNSV1 struct {
	*MockService
	pb.UnimplementedPoliciesServer
}

func (s *DNSV1) GetPolicy(ctx context.Context, req *pb.GetPolicyRequest) (*pb.Policy, error) {
	reqName := "projects/" + req.Project + "/policies/" + req.Name
	name, err := s.parsePolicyName(reqName)
	if err != nil {
		return nil, err
	}

	fqn := name.String()

	obj := &pb.Policy{}
	if err := s.storage.Get(ctx, fqn, obj); err != nil {
		if apierrors.IsNotFound(err) {
			return nil, status.Errorf(codes.NotFound, "policy %q not found", name)
		} else {
			return nil, status.Errorf(codes.Internal, "error reading policy: %v", err)
		}
	}

	return obj, nil
}

func (s *DNSV1) CreatePolicy(ctx context.Context, req *pb.CreatePolicyRequest) (*pb.Policy, error) {
	reqName := "projects/" + req.Project + "/policies/" + req.GetPolicy().Name
	name, err := s.parsePolicyName(reqName)
	if err != nil {
		return nil, err
	}

	fqn := name.String()

	obj := proto.Clone(req.Policy).(*pb.Policy)

	if err := s.storage.Create(ctx, fqn, obj); err != nil {
		return nil, status.Errorf(codes.Internal, "error creating Policy: %v", err)
	}

	return obj, nil
}

func (s *DNSV1) PatchPolicy(ctx context.Context, req *pb.PatchPolicyRequest) (*pb.PoliciesPatchResponse, error) {
	reqName := "projects/" + req.Project + "/policies/" + req.Name
	name, err := s.parsePolicyName(reqName)
	if err != nil {
		return nil, err
	}

	fqn := name.String()
	obj := &pb.Policy{}
	if err := s.storage.Get(ctx, fqn, obj); err != nil {
		if apierrors.IsNotFound(err) {
			return nil, status.Errorf(codes.NotFound, "Policy %q not found", reqName)
		}
		return nil, status.Errorf(codes.Internal, "error reading Policy: %v", err)
	}

	update := req.GetPolicy()
	if update.AlternativeNameServerConfig != nil {
		obj.AlternativeNameServerConfig = update.AlternativeNameServerConfig
	}

	// TODO: Do we need to use proto2?
	// if update.Description != nil {
	// 	obj.Description = update.Description
	// }
	// if update.EnableInboundForwarding != nil {
	// 	obj.EnableInboundForwarding = update.EnableInboundForwarding
	// }
	// if update.EnableLoging != nil {
	// 	obj.EnableLogging = update.EnableLoging
	// }
	if update.Networks != nil {
		obj.Networks = update.Networks
	}

	if err := s.storage.Update(ctx, fqn, obj); err != nil {
		return nil, status.Errorf(codes.Internal, "error updating Policy: %v", err)
	}

	response := &pb.PoliciesPatchResponse{
		Policy: obj,
	}
	return response, nil
}

func (s *DNSV1) DeletePolicy(ctx context.Context, req *pb.DeletePolicyRequest) (*pb.DeletePolicyResponse, error) {
	reqName := "projects/" + req.Project + "/policies/" + req.Name
	name, err := s.parsePolicyName(reqName)
	if err != nil {
		return nil, err
	}

	fqn := name.String()

	deletedObj := &pb.Policy{}
	if err := s.storage.Delete(ctx, fqn, deletedObj); err != nil {
		if apierrors.IsNotFound(err) {
			return nil, status.Errorf(codes.NotFound, "Policy %q not found", name)
		} else {
			return nil, status.Errorf(codes.Internal, "error deleting Policy: %v", err)
		}
	}

	response := &pb.DeletePolicyResponse{}
	return response, nil
}
