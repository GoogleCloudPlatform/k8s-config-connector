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

// +tool:mockgcp-support
// proto.service: google.cloud.vmwareengine.v1.VmwareEngine
// proto.message: google.cloud.vmwareengine.v1.NetworkPolicy

package mockvmwareengine

import (
	"context"
	"fmt"
	"strings"
	"time"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/types/known/emptypb"
	"google.golang.org/protobuf/types/known/timestamppb"

	"github.com/GoogleCloudPlatform/k8s-config-connector/mockgcp/common/projects"
	pb "github.com/GoogleCloudPlatform/k8s-config-connector/mockgcp/generated/mockgcp/cloud/vmwareengine/v1"
	longrunningpb "google.golang.org/genproto/googleapis/longrunning"
)

func (s *VMwareEngineV1) GetNetworkPolicy(ctx context.Context, req *pb.GetNetworkPolicyRequest) (*pb.NetworkPolicy, error) {
	name, err := s.parseNetworkPolicyName(req.Name)
	if err != nil {
		return nil, err
	}

	fqn := name.String()

	obj := &pb.NetworkPolicy{}
	if err := s.storage.Get(ctx, fqn, obj); err != nil {
		if status.Code(err) == codes.NotFound {
			return nil, status.Errorf(codes.NotFound, "Resource '%s' was not found", fqn)
		}
		return nil, err
	}

	return obj, nil
}

func (s *VMwareEngineV1) ListNetworkPolicies(ctx context.Context, req *pb.ListNetworkPoliciesRequest) (*pb.ListNetworkPoliciesResponse, error) {
	parent, err := s.parseNetworkPolicyParent(req.Parent)
	if err != nil {
		return nil, err
	}

	prefix := parent.String() + "/"

	var objs []*pb.NetworkPolicy
	if err := s.storage.List(ctx, prefix, func(fqn string, obj proto.Message) error {
		objs = append(objs, obj.(*pb.NetworkPolicy))
		return nil
	}); err != nil {
		return nil, err
	}

	return &pb.ListNetworkPoliciesResponse{NetworkPolicies: objs}, nil
}

func (s *VMwareEngineV1) CreateNetworkPolicy(ctx context.Context, req *pb.CreateNetworkPolicyRequest) (*longrunningpb.Operation, error) {
	reqName := req.Parent + "/networkPolicies/" + req.NetworkPolicyId
	name, err := s.parseNetworkPolicyName(reqName)
	if err != nil {
		return nil, err
	}

	fqn := name.String()

	// Check that the referenced VmwareEngineNetwork exists.
	if _, err := s.GetVmwareEngineNetwork(ctx, &pb.GetVmwareEngineNetworkRequest{Name: req.NetworkPolicy.VmwareEngineNetwork}); err != nil {
		return nil, status.Errorf(codes.FailedPrecondition, "vmwareEngineNetwork '%s' not found", req.NetworkPolicy.VmwareEngineNetwork)
	}

	now := time.Now()

	obj := proto.Clone(req.GetNetworkPolicy()).(*pb.NetworkPolicy)
	obj.Name = fqn
	obj.CreateTime = timestamppb.New(now)
	obj.UpdateTime = timestamppb.New(now)
	obj.Uid = "111111111111111111111"
	obj.VmwareEngineNetworkCanonical = req.NetworkPolicy.VmwareEngineNetwork

	if err := s.storage.Create(ctx, fqn, obj); err != nil {
		return nil, err
	}

	prefix := fmt.Sprintf("projects/%s/locations/%s", name.Project.ID, name.Location)
	metadata := &pb.OperationMetadata{
		ApiVersion: "v1",
		CreateTime: timestamppb.New(now),
		Target:     name.String(),
		Verb:       "create",
	}
	return s.operations.StartLRO(ctx, prefix, metadata, func() (proto.Message, error) {
		metadata.EndTime = timestamppb.New(now)
		obj.UpdateTime = timestamppb.New(now)
		if err := s.storage.Update(ctx, fqn, obj); err != nil {
			return nil, err
		}
		return obj, nil
	})
}

func (s *VMwareEngineV1) UpdateNetworkPolicy(ctx context.Context, req *pb.UpdateNetworkPolicyRequest) (*longrunningpb.Operation, error) {
	name, err := s.parseNetworkPolicyName(req.GetNetworkPolicy().GetName())
	if err != nil {
		return nil, err
	}

	fqn := name.String()

	// Check that the referenced VmwareEngineNetwork exists.
	if req.NetworkPolicy.VmwareEngineNetwork != "" {
		if _, err := s.GetVmwareEngineNetwork(ctx, &pb.GetVmwareEngineNetworkRequest{Name: req.NetworkPolicy.VmwareEngineNetwork}); err != nil {
			return nil, status.Errorf(codes.FailedPrecondition, "vmwareEngineNetwork '%s' not found", req.NetworkPolicy.VmwareEngineNetwork)
		}
	}

	original := &pb.NetworkPolicy{}
	if err := s.storage.Get(ctx, fqn, original); err != nil {
		return nil, err
	}

	obj := proto.Clone(original).(*pb.NetworkPolicy)
	if req.NetworkPolicy.InternetAccess != nil {
		obj.InternetAccess = req.NetworkPolicy.InternetAccess
	}
	if req.NetworkPolicy.ExternalIp != nil {
		obj.ExternalIp = req.NetworkPolicy.ExternalIp
	}
	if req.NetworkPolicy.EdgeServicesCidr != "" {
		obj.EdgeServicesCidr = req.NetworkPolicy.EdgeServicesCidr
	}
	if req.NetworkPolicy.Description != "" {
		obj.Description = req.NetworkPolicy.Description
	}

	obj.UpdateTime = timestamppb.Now()

	if err := s.storage.Update(ctx, fqn, obj); err != nil {
		return nil, err
	}

	prefix := fmt.Sprintf("projects/%s/locations/%s", name.Project.ID, name.Location)
	metadata := &pb.OperationMetadata{
		ApiVersion: "v1",
		CreateTime: timestamppb.Now(),
		Target:     name.String(),
		Verb:       "update",
	}
	return s.operations.StartLRO(ctx, prefix, metadata, func() (proto.Message, error) {
		metadata.EndTime = timestamppb.Now()
		if err := s.storage.Update(ctx, fqn, obj); err != nil {
			return nil, err
		}
		return obj, nil
	})
}

func (s *VMwareEngineV1) DeleteNetworkPolicy(ctx context.Context, req *pb.DeleteNetworkPolicyRequest) (*longrunningpb.Operation, error) {
	name, err := s.parseNetworkPolicyName(req.Name)
	if err != nil {
		return nil, err
	}

	fqn := name.String()

	deleted := &pb.NetworkPolicy{}
	if err := s.storage.Delete(ctx, fqn, deleted); err != nil {
		return nil, err
	}

	prefix := fmt.Sprintf("projects/%s/locations/%s", name.Project.ID, name.Location)
	metadata := &pb.OperationMetadata{
		CreateTime: timestamppb.Now(),
		Target:     fqn,
		ApiVersion: "v1",
		Verb:       "delete",
	}
	return s.operations.StartLRO(ctx, prefix, metadata, func() (proto.Message, error) {
		metadata.EndTime = timestamppb.Now()
		return &emptypb.Empty{}, nil
	})
}

type networkPolicyName struct {
	Project         *projects.ProjectData
	Location        string
	NetworkPolicyID string
}

func (n *networkPolicyName) String() string {
	return fmt.Sprintf("projects/%s/locations/%s/networkPolicies/%s", n.Project.ID, n.Location, n.NetworkPolicyID)
}

// parseNetworkPolicyName parses a string into a networkPolicyName.
// The expected form is `projects/*/locations/*/networkPolicies/*`.
func (s *MockService) parseNetworkPolicyName(name string) (*networkPolicyName, error) {
	tokens := strings.Split(name, "/")

	if len(tokens) == 6 && tokens[0] == "projects" && tokens[2] == "locations" && tokens[4] == "networkPolicies" {
		project, err := s.Projects.GetProjectByID(tokens[1])
		if err != nil {
			return nil, err
		}

		name := &networkPolicyName{
			Project:         project,
			Location:        tokens[3],
			NetworkPolicyID: tokens[5],
		}

		return name, nil
	}

	return nil, status.Errorf(codes.InvalidArgument, "name %q is not valid", name)
}

type networkPolicyParent struct {
	Project  *projects.ProjectData
	Location string
}

func (n *networkPolicyParent) String() string {
	return fmt.Sprintf("projects/%s/locations/%s", n.Project.ID, n.Location)
}

// parseNetworkPolicyParent parses a string into a networkPolicyParent.
// The expected form is `projects/*/locations/*`.
func (s *MockService) parseNetworkPolicyParent(name string) (*networkPolicyParent, error) {
	tokens := strings.Split(name, "/")

	if len(tokens) == 4 && tokens[0] == "projects" && tokens[2] == "locations" {
		project, err := s.Projects.GetProjectByID(tokens[1])
		if err != nil {
			return nil, err
		}

		name := &networkPolicyParent{
			Project:  project,
			Location: tokens[3],
		}

		return name, nil
	}

	return nil, status.Errorf(codes.InvalidArgument, "parent name %q is not valid", name)
}
