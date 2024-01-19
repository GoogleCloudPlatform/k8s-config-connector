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

package mockdns

import (
	"context"
	"strings"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/types/known/emptypb"

	"github.com/GoogleCloudPlatform/k8s-config-connector/mockgcp/common/projects"
	pb "github.com/GoogleCloudPlatform/k8s-config-connector/mockgcp/generated/mockgcp/cloud/dns/v1beta2"
)

type policiesServer struct {
	*MockService
	pb.UnimplementedPoliciesServerServer
}

func (s *policiesServer) GetPolicy(ctx context.Context, req *pb.GetPolicyRequest) (*pb.Policy, error) {
	name, err := s.newPolicyName(req.GetProject(), req.GetName())
	if err != nil {
		return nil, err
	}

	fqn := name.String()

	obj := &pb.Policy{}
	if err := s.storage.Get(ctx, fqn, obj); err != nil {
		return nil, err
	}

	return obj, nil
}

func (s *policiesServer) CreatePolicy(ctx context.Context, req *pb.CreatePolicyRequest) (*pb.Policy, error) {
	name, err := s.newPolicyName(req.GetProject(), req.GetPolicy().GetName())
	if err != nil {
		return nil, err
	}

	fqn := name.String()

	obj := proto.Clone(req.Policy).(*pb.Policy)

	if err := s.storage.Create(ctx, fqn, obj); err != nil {
		return nil, err
	}

	return obj, nil
}

func (s *policiesServer) PatchPolicy(ctx context.Context, req *pb.PatchPolicyRequest) (*pb.PoliciesPatchResponse, error) {
	name, err := s.newPolicyName(req.GetProject(), req.GetName())
	if err != nil {
		return nil, err
	}

	fqn := name.String()
	obj := &pb.Policy{}
	if err := s.storage.Get(ctx, fqn, obj); err != nil {
		return nil, err
	}

	update := req.GetPolicy()
	if update.AlternativeNameServerConfig != nil {
		obj.AlternativeNameServerConfig = update.AlternativeNameServerConfig
	}

	if update.Description != nil {
		obj.Description = update.Description
	}
	if update.EnableInboundForwarding != nil {
		obj.EnableInboundForwarding = update.EnableInboundForwarding
	}
	if update.EnableLogging != nil {
		obj.EnableLogging = update.EnableLogging
	}

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

func (s *policiesServer) DeletePolicy(ctx context.Context, req *pb.DeletePolicyRequest) (*emptypb.Empty, error) {
	name, err := s.newPolicyName(req.GetProject(), req.GetName())
	if err != nil {
		return nil, err
	}

	fqn := name.String()

	deletedObj := &pb.Policy{}
	if err := s.storage.Delete(ctx, fqn, deletedObj); err != nil {
		return nil, err
	}

	response := &emptypb.Empty{}
	return response, nil
}

type policyName struct {
	Project *projects.ProjectData
	Name    string
}

func (n *policyName) String() string {
	return "projects/" + n.Project.ID + "/policies/" + n.Name
}

// parsePolicyName parses a string into a policyName.
// The expected form is `projects/*/policies/*`.
func (s *MockService) parsePolicyName(name string) (*policyName, error) {
	tokens := strings.Split(name, "/")

	if len(tokens) == 4 && tokens[0] == "projects" && tokens[2] == "policies" {
		return s.newPolicyName(tokens[1], tokens[3])
	} else {
		return nil, status.Errorf(codes.InvalidArgument, "name %q is not valid", name)
	}
}

// newPolicyName builds a policyName from its components.
func (s *MockService) newPolicyName(projectID string, name string) (*policyName, error) {
	project, err := s.Projects.GetProjectByID(projectID)
	if err != nil {
		return nil, err
	}

	return &policyName{
		Project: project,
		Name:    name,
	}, nil
}
