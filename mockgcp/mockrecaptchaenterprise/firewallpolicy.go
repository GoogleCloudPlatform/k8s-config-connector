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
// proto.service: google.cloud.recaptchaenterprise.v1.RecaptchaEnterpriseService
// proto.message: google.cloud.recaptchaenterprise.v1.FirewallPolicy

package mockrecaptchaenterprise

import (
	"context"
	"fmt"
	"strings"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/types/known/emptypb"

	"github.com/GoogleCloudPlatform/k8s-config-connector/mockgcp/common/projects"
	pb "github.com/GoogleCloudPlatform/k8s-config-connector/mockgcp/generated/mockgcp/cloud/recaptchaenterprise/v1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/mockgcp/pkg/storage"
)

type recaptchaEnterpriseService struct {
	*MockService
	pb.UnimplementedRecaptchaEnterpriseServiceServer
}

func (s *recaptchaEnterpriseService) GetFirewallPolicy(ctx context.Context, req *pb.GetFirewallPolicyRequest) (*pb.FirewallPolicy, error) {
	name, err := s.parseFirewallPolicyName(req.Name)
	if err != nil {
		return nil, err
	}

	fqn := name.String()

	obj := &pb.FirewallPolicy{}
	if err := s.storage.Get(ctx, fqn, obj); err != nil {
		return nil, err
	}

	return obj, nil
}

func (s *recaptchaEnterpriseService) CreateFirewallPolicy(ctx context.Context, req *pb.CreateFirewallPolicyRequest) (*pb.FirewallPolicy, error) {
	reqName := fmt.Sprintf("%s/firewallpolicies/%d", req.GetParent(), 100)
	name, err := s.parseFirewallPolicyName(reqName)
	if err != nil {
		return nil, err
	}
	fqn := name.String()

	obj := proto.Clone(req.GetFirewallPolicy()).(*pb.FirewallPolicy)
	obj.Name = fqn

	if err := s.storage.Create(ctx, fqn, obj); err != nil {
		return nil, err
	}

	obj.Name = strings.ReplaceAll(obj.Name, name.Project.ID, fmt.Sprintf("%v", name.Project.Number))
	return obj, nil
}

func (s *recaptchaEnterpriseService) UpdateFirewallPolicy(ctx context.Context, req *pb.UpdateFirewallPolicyRequest) (*pb.FirewallPolicy, error) {
	reqName := req.GetFirewallPolicy().GetName()

	name, err := s.parseFirewallPolicyName(reqName)
	if err != nil {
		return nil, err
	}

	fqn := name.String()

	obj := &pb.FirewallPolicy{}
	if err := s.storage.Get(ctx, fqn, obj); err != nil {
		return nil, err
	}

	// Required. A list of fields to be updated in this request.
	paths := req.GetUpdateMask().GetPaths()
	if len(paths) == 0 {
		return nil, status.Errorf(codes.InvalidArgument, "update_mask must be provided")
	}

	for _, path := range paths {
		switch path {
		case "description":
			obj.Description = req.GetFirewallPolicy().GetDescription()
		case "path":
			obj.Path = req.GetFirewallPolicy().GetPath()
		case "condition":
			obj.Condition = req.GetFirewallPolicy().GetCondition()
		case "actions":
			obj.Actions = req.GetFirewallPolicy().GetActions()
		default:
			return nil, status.Errorf(codes.InvalidArgument, "update_mask path %q not valid", path)
		}
	}

	if err := s.storage.Update(ctx, fqn, obj); err != nil {
		return nil, err
	}

	return obj, nil
}

func (s *recaptchaEnterpriseService) DeleteFirewallPolicy(ctx context.Context, req *pb.DeleteFirewallPolicyRequest) (*emptypb.Empty, error) {
	name, err := s.parseFirewallPolicyName(req.Name)
	if err != nil {
		return nil, err
	}
	fqn := name.String()

	deletedObj := &pb.FirewallPolicy{}
	if err := s.storage.Delete(ctx, fqn, deletedObj); err != nil {
		return nil, err
	}

	return &emptypb.Empty{}, nil
}

type firewallPolicyName struct {
	Project *projects.ProjectData
	Name    string
}

func (n *firewallPolicyName) String() string {
	return fmt.Sprintf("projects/%s/firewallpolicies/%s", n.Project.ID, n.Name)
}

// parseFirewallPolicyName parses a string into an FirewallPolicy name.
// The expected form is `projects/*/firewallpolicies/*`.
func (s *MockService) parseFirewallPolicyName(name string) (*firewallPolicyName, error) {
	tokens := strings.Split(name, "/")

	if len(tokens) == 4 && tokens[0] == "projects" && tokens[2] == "firewallpolicies" {
		project, err := s.Projects.GetProjectByID(tokens[1])
		if err != nil {
			return nil, err
		}

		name := &firewallPolicyName{
			Project: project,
			Name:    tokens[3],
		}

		return name, nil
	}
	return nil, status.Errorf(codes.InvalidArgument, "name %q is not valid", name)
}

func (s *recaptchaEnterpriseService) ListFirewallPolicies(ctx context.Context, req *pb.ListFirewallPoliciesRequest) (*pb.ListFirewallPoliciesResponse, error) {
	response := &pb.ListFirewallPoliciesResponse{}
	kind := (&pb.FirewallPolicy{}).ProtoReflect().Descriptor()
	if err := s.storage.List(ctx, kind, storage.ListOptions{Prefix: req.Parent}, func(msg proto.Message) error {
		fPolicy, ok := msg.(*pb.FirewallPolicy)
		if ok {
			response.FirewallPolicies = append(response.FirewallPolicies, fPolicy)
		}
		return nil
	}); err != nil {
		return nil, status.Errorf(codes.Internal, "error listing FirewallPolicy in %q: %v", req.Parent, err)
	}
	return response, nil
}
