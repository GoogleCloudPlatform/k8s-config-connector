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
package mocknetworksecurity

import (
	"context"
	"strings"
	"time"

	"github.com/GoogleCloudPlatform/k8s-config-connector/mockgcp/common"
	"github.com/GoogleCloudPlatform/k8s-config-connector/mockgcp/common"
	"github.com/GoogleCloudPlatform/k8s-config-connector/mockgcp/common"
	"github.com/GoogleCloudPlatform/k8s-config-connector/mockgcp/common/projects"
	pb "google.golang.org/genproto/googleapis/cloud/networksecurity/v1beta1"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/types/known/timestamppb"
)

func (s *MockService) CreateAuthorizationPolicy(ctx context.Context, req *pb.CreateAuthorizationPolicyRequest) (*common.Operation, error) {
	name, err := s.buildAuthorizationPolicyName(req.Parent, req.AuthorizationPolicyId)
	if err != nil {
		return nil, err
	}

	fqn := name.String()

	obj := proto.Clone(req.AuthorizationPolicy).(*pb.AuthorizationPolicy)
	obj.Name = fqn
	obj.CreateTime = timestamppb.New(time.Now())
	obj.UpdateTime = timestamppb.New(time.Now())

	if err := s.storage.Create(ctx, fqn, obj); err != nil {
		return nil, err
	}

	return s.operations.NewLRO(ctx)
}

func (s *MockService) GetAuthorizationPolicy(ctx context.Context, req *pb.GetAuthorizationPolicyRequest) (*pb.AuthorizationPolicy, error) {
	fqn := req.Name
	obj := &pb.AuthorizationPolicy{}
	if err := s.storage.Get(ctx, fqn, obj); err != nil {
		if status.Code(err) == codes.NotFound {
			return nil, status.Errorf(codes.NotFound, "Resource '%s' was not found", fqn)
		}
		return nil, err
	}
	return obj, nil
}

func (s *MockService) UpdateAuthorizationPolicy(ctx context.Context, req *pb.UpdateAuthorizationPolicyRequest) (*common.Operation, error) {
	fqn := req.AuthorizationPolicy.Name
	obj := &pb.AuthorizationPolicy{}
	if err := s.storage.Get(ctx, fqn, obj); err != nil {
		return nil, err
	}

	// TODO: implement field mask
	updated := proto.Clone(req.AuthorizationPolicy).(*pb.AuthorizationPolicy)
	updated.CreateTime = obj.CreateTime
	updated.UpdateTime = timestamppb.New(time.Now())

	if err := s.storage.Update(ctx, fqn, updated); err != nil {
		return nil, err
	}

	return s.operations.NewLRO(ctx)
}

func (s *MockService) DeleteAuthorizationPolicy(ctx context.Context, req *pb.DeleteAuthorizationPolicyRequest) (*common.Operation, error) {
	fqn := req.Name
	if err := s.storage.Delete(ctx, fqn, &pb.AuthorizationPolicy{}); err != nil {
		return nil, err
	}
	return s.operations.NewLRO(ctx)
}

type authorizationPolicyName struct {
	Project               *projects.ProjectData
	Location              string
	AuthorizationPolicyID string
}

func (n *authorizationPolicyName) String() string {
	return "projects/" + n.Project.ID + "/locations/" + n.Location + "/authorizationPolicies/" + n.AuthorizationPolicyID
}

func (s *MockService) parseAuthorizationPolicyName(name string) (*authorizationPolicyName, error) {
	tokens := strings.Split(name, "/")
	if len(tokens) == 6 && tokens[0] == "projects" && tokens[2] == "locations" && tokens[4] == "authorizationPolicies" {
		project, err := s.Projects.GetProject(&projects.ProjectName{ProjectID: tokens[1]})
		if err != nil {
			return nil, err
		}
		name := &authorizationPolicyName{
			Project:               project,
			Location:              tokens[3],
			AuthorizationPolicyID: tokens[5],
		}
		return name, nil
	} else {
		return nil, status.Errorf(codes.InvalidArgument, "name %q is not valid", name)
	}
}

func (s *MockService) buildAuthorizationPolicyName(parent string, authorizationPolicyID string) (*authorizationPolicyName, error) {
	tokens := strings.Split(parent, "/")
	if len(tokens) == 4 && tokens[0] == "projects" && tokens[2] == "locations" {
		project, err := s.Projects.GetProject(&projects.ProjectName{ProjectID: tokens[1]})
		if err != nil {
			return nil, err
		}
		name := &authorizationPolicyName{
			Project:               project,
			Location:              tokens[3],
			AuthorizationPolicyID: authorizationPolicyID,
		}
		return name, nil
	} else {
		return nil, status.Errorf(codes.InvalidArgument, "parent %q is not valid", parent)
	}
}
