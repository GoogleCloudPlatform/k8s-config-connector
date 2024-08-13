// Copyright 2022 Google LLC
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

type GlobalForwardingRulesV1 struct {
	*MockService
	pb.UnimplementedGlobalForwardingRulesServer
}

func (s *GlobalForwardingRulesV1) Get(ctx context.Context, req *pb.GetGlobalForwardingRuleRequest) (*pb.ForwardingRule, error) {
	reqName := "projects/" + req.GetProject() + "/global" + "/forwardingRules/" + req.GetForwardingRule()
	name, err := s.parseGlobalForwardingRuleName(reqName)
	if err != nil {
		return nil, err
	}

	fqn := name.String()

	obj := &pb.ForwardingRule{}
	if err := s.storage.Get(ctx, fqn, obj); err != nil {
		return nil, err
	}

	return obj, nil
}

func (s *GlobalForwardingRulesV1) Insert(ctx context.Context, req *pb.InsertGlobalForwardingRuleRequest) (*pb.Operation, error) {
	reqName := "projects/" + req.GetProject() + "/global" + "/forwardingRules/" + req.GetForwardingRuleResource().GetName()
	name, err := s.parseGlobalForwardingRuleName(reqName)
	if err != nil {
		return nil, err
	}

	fqn := name.String()

	id := s.generateID()

	obj := proto.Clone(req.GetForwardingRuleResource()).(*pb.ForwardingRule)
	obj.SelfLink = PtrTo("https://www.googleapis.com/compute/v1/" + name.String())
	obj.CreationTimestamp = PtrTo(s.nowString())
	obj.Id = &id
	obj.Kind = PtrTo("compute#forwardingRule")
	obj.LabelFingerprint = PtrTo("abcdef0123A=")

	if err := s.storage.Create(ctx, fqn, obj); err != nil {
		return nil, err
	}

	return s.newLRO(ctx, name.Project.ID)
}

func (s *GlobalForwardingRulesV1) Delete(ctx context.Context, req *pb.DeleteGlobalForwardingRuleRequest) (*pb.Operation, error) {
	reqName := "projects/" + req.GetProject() + "/global" + "/forwardingRules/" + req.GetForwardingRule()
	name, err := s.parseGlobalForwardingRuleName(reqName)
	if err != nil {
		return nil, err
	}

	fqn := name.String()

	deleted := &pb.ForwardingRule{}
	if err := s.storage.Delete(ctx, fqn, deleted); err != nil {
		return nil, err
	}

	return s.newLRO(ctx, name.Project.ID)
}

func (s *GlobalForwardingRulesV1) SetLabels(ctx context.Context, req *pb.SetLabelsGlobalForwardingRuleRequest) (*pb.Operation, error) {
	reqName := "projects/" + req.GetProject() + "/global" + "/forwardingRules/" + req.GetResource()
	name, err := s.parseGlobalForwardingRuleName(reqName)
	if err != nil {
		return nil, err
	}

	fqn := name.String()

	obj := &pb.ForwardingRule{}
	if err := s.storage.Get(ctx, fqn, obj); err != nil {
		return nil, err
	}

	obj.Labels = req.GetGlobalSetLabelsRequestResource().GetLabels()
	if err := s.storage.Update(ctx, fqn, obj); err != nil {
		return nil, err
	}

	return s.newLRO(ctx, name.Project.ID)
}

func (s *GlobalForwardingRulesV1) SetTarget(ctx context.Context, req *pb.SetTargetGlobalForwardingRuleRequest) (*pb.Operation, error) {
	reqName := "projects/" + req.GetProject() + "/global" + "/forwardingRules/" + req.GetForwardingRule()
	name, err := s.parseGlobalForwardingRuleName(reqName)
	if err != nil {
		return nil, err
	}

	fqn := name.String()

	obj := &pb.ForwardingRule{}
	if err := s.storage.Get(ctx, fqn, obj); err != nil {
		return nil, err
	}

	obj.Target = req.GetTargetReferenceResource().Target
	if err := s.storage.Update(ctx, fqn, obj); err != nil {
		return nil, err
	}

	return s.newLRO(ctx, name.Project.ID)
}

type globalForwardingRuleName struct {
	Project *projects.ProjectData
	Name    string
}

func (n *globalForwardingRuleName) String() string {
	return "projects/" + n.Project.ID + "/global" + "/forwardingRules/" + n.Name
}

// parseForwardingRuleName parses a string into a forwardingruleName.
// The expected form is `projects/*/regions/*/forwardingrule/*`.
func (s *MockService) parseGlobalForwardingRuleName(name string) (*globalForwardingRuleName, error) {
	tokens := strings.Split(name, "/")

	if len(tokens) == 5 && tokens[0] == "projects" && tokens[2] == "global" && tokens[3] == "forwardingRules" {
		project, err := s.Projects.GetProjectByID(tokens[1])
		if err != nil {
			return nil, err
		}

		name := &globalForwardingRuleName{
			Project: project,
			Name:    tokens[4],
		}

		return name, nil
	} else {
		return nil, status.Errorf(codes.InvalidArgument, "name %q is not valid", name)
	}
}
