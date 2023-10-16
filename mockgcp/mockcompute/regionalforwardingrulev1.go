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

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/proto"
	apierrors "k8s.io/apimachinery/pkg/api/errors"

	"github.com/GoogleCloudPlatform/k8s-config-connector/mockgcp/common/projects"
	pb "github.com/GoogleCloudPlatform/k8s-config-connector/mockgcp/generated/mockgcp/cloud/compute/v1"
)

type RegionalForwardingRulesV1 struct {
	*MockService
	pb.UnimplementedForwardingRulesServer
}

func (s *RegionalForwardingRulesV1) Get(ctx context.Context, req *pb.GetForwardingRuleRequest) (*pb.ForwardingRule, error) {
	reqName := "projects/" + req.GetProject() + "/regions/" + req.GetRegion() + "/forwardingRules/" + req.GetForwardingRule()
	name, err := s.parseRegionalForwardingRuleName(reqName)
	if err != nil {
		return nil, err
	}

	fqn := name.String()

	obj := &pb.ForwardingRule{}
	if err := s.storage.Get(ctx, fqn, obj); err != nil {
		if apierrors.IsNotFound(err) {
			return nil, status.Errorf(codes.NotFound, "forwardingRule %q not found", name)
		} else {
			return nil, status.Errorf(codes.Internal, "error reading forwardingRule: %v", err)
		}
	}

	return obj, nil
}

func (s *RegionalForwardingRulesV1) Insert(ctx context.Context, req *pb.InsertForwardingRuleRequest) (*pb.Operation, error) {
	reqName := "projects/" + req.GetProject() + "/regions/" + req.GetRegion() + "/forwardingRules/" + req.GetForwardingRuleResource().GetName()
	name, err := s.parseRegionalForwardingRuleName(reqName)
	if err != nil {
		return nil, err
	}

	fqn := name.String()

	id := s.generateID()

	obj := proto.Clone(req.GetForwardingRuleResource()).(*pb.ForwardingRule)
	obj.SelfLink = PtrTo("https://compute.googleapis.com/compute/v1/" + name.String())
	obj.CreationTimestamp = PtrTo(s.nowString())
	obj.Id = &id
	obj.Kind = PtrTo("compute#forwardingRule")

	if err := s.storage.Create(ctx, fqn, obj); err != nil {
		return nil, status.Errorf(codes.Internal, "error creating forwardingRule: %v", err)
	}

	return s.newLRO(ctx, name.Project.ID)
}

func (s *RegionalForwardingRulesV1) Delete(ctx context.Context, req *pb.DeleteForwardingRuleRequest) (*pb.Operation, error) {
	reqName := "projects/" + req.GetProject() + "/regions/" + req.GetRegion() + "/forwardingRules/" + req.GetForwardingRule()
	name, err := s.parseRegionalForwardingRuleName(reqName)
	if err != nil {
		return nil, err
	}

	fqn := name.String()

	deleted := &pb.ForwardingRule{}
	if err := s.storage.Delete(ctx, fqn, deleted); err != nil {
		if apierrors.IsNotFound(err) {
			return nil, status.Errorf(codes.NotFound, "forwardingRule %q not found", name)
		} else {
			return nil, status.Errorf(codes.Internal, "error deleting forwardingRule: %v", err)
		}
	}

	return s.newLRO(ctx, name.Project.ID)
}

func (s *RegionalForwardingRulesV1) SetLabels(ctx context.Context, req *pb.SetLabelsForwardingRuleRequest) (*pb.Operation, error) {
	reqName := "projects/" + req.GetProject() + "/regions/" + req.GetRegion() + "/forwardingRules/" + req.GetResource()
	name, err := s.parseRegionalForwardingRuleName(reqName)
	if err != nil {
		return nil, err
	}

	fqn := name.String()

	obj := &pb.ForwardingRule{}
	if err := s.storage.Get(ctx, fqn, obj); err != nil {
		if apierrors.IsNotFound(err) {
			return nil, status.Errorf(codes.NotFound, "forwardingRule %q not found", name)
		} else {
			return nil, status.Errorf(codes.Internal, "error reading forwardingRule: %v", err)
		}
	}

	obj.Labels = req.GetRegionSetLabelsRequestResource().GetLabels()
	if err := s.storage.Update(ctx, fqn, obj); err != nil {
		return nil, status.Errorf(codes.Internal, "error updating forwardingRule: %v", err)
	}

	return s.newLRO(ctx, name.Project.ID)
}

func (s *RegionalForwardingRulesV1) SetTarget(ctx context.Context, req *pb.SetTargetForwardingRuleRequest) (*pb.Operation, error) {
	reqName := "projects/" + req.GetProject() + "/regions/" + req.GetRegion() + "/forwardingRules/" + req.GetForwardingRule()
	name, err := s.parseRegionalForwardingRuleName(reqName)
	if err != nil {
		return nil, err
	}

	fqn := name.String()

	obj := &pb.ForwardingRule{}
	if err := s.storage.Get(ctx, fqn, obj); err != nil {
		if apierrors.IsNotFound(err) {
			return nil, status.Errorf(codes.NotFound, "forwardingRule %q not found", name)
		} else {
			return nil, status.Errorf(codes.Internal, "error reading forwardingRule: %v", err)
		}
	}

	obj.Target = req.GetTargetReferenceResource().Target
	if err := s.storage.Update(ctx, fqn, obj); err != nil {
		return nil, status.Errorf(codes.Internal, "error updating forwardingRule: %v", err)
	}

	return s.newLRO(ctx, name.Project.ID)
}

type regionalForwardingRuleName struct {
	Project *projects.ProjectData
	Region  string
	Name    string
}

func (n *regionalForwardingRuleName) String() string {
	return "projects/" + n.Project.ID + "/regions/" + n.Region + "/forwardingRules/" + n.Name
}

// parseForwardingRuleName parses a string into a forwardingruleName.
// The expected form is `projects/*/regions/*/forwardingrule/*`.
func (s *MockService) parseRegionalForwardingRuleName(name string) (*regionalForwardingRuleName, error) {
	tokens := strings.Split(name, "/")

	if len(tokens) == 6 && tokens[0] == "projects" && tokens[2] == "regions" && tokens[4] == "forwardingRules" {
		project, err := s.projects.GetProjectByID(tokens[1])
		if err != nil {
			return nil, err
		}

		name := &regionalForwardingRuleName{
			Project: project,
			Region:  tokens[3],
			Name:    tokens[5],
		}

		return name, nil
	} else {
		return nil, status.Errorf(codes.InvalidArgument, "name %q is not valid", name)
	}
}
