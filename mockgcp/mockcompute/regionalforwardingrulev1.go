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
	"fmt"
	"strings"

	"github.com/GoogleCloudPlatform/k8s-config-connector/mockgcp/common/projects"
	pb "github.com/GoogleCloudPlatform/k8s-config-connector/mockgcp/generated/mockgcp/cloud/compute/v1"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/proto"
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
		if status.Code(err) == codes.NotFound {
			return nil, status.Errorf(codes.NotFound, "The resource '%s' was not found", fqn)
		}
		return nil, err
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
	obj.SelfLink = PtrTo("https://www.googleapis.com/compute/v1/" + name.String())
	obj.CreationTimestamp = PtrTo(s.nowString())
	obj.Id = &id
	obj.Kind = PtrTo("compute#forwardingRule")
	// labels will be added separately with setLabels
	obj.Labels = nil
	// If below values are not provided by user, it appears to default by GCP
	if obj.LabelFingerprint == nil {
		obj.LabelFingerprint = PtrTo(computeFingerprint(obj))
	}
	if obj.Fingerprint == nil {
		obj.Fingerprint = PtrTo(computeFingerprint(obj))
	}
	if obj.IPProtocol == nil {
		obj.IPProtocol = PtrTo("TCP")
	}
	if obj.NetworkTier == nil {
		obj.NetworkTier = PtrTo("PREMIUM")
	}
	if obj.LoadBalancingScheme != nil && *obj.LoadBalancingScheme == "" {
		obj.LoadBalancingScheme = nil
	}

	// pattern: \d+(?:-\d+)?
	if obj.PortRange != nil {
		r := obj.GetPortRange()
		token := strings.Split(r, "-")
		if len(token) == 1 {
			obj.PortRange = PtrTo(fmt.Sprintf("%s-%s", token[0], token[0]))
		} else if len(token) == 2 {
			obj.PortRange = PtrTo(fmt.Sprintf("%s-%s", token[0], token[1]))
		} else {
			return nil, status.Errorf(codes.InvalidArgument, "portRange %s is not valid", obj.GetPortRange())
		}
	}

	if obj.Network != nil {
		networkName, err := s.parseNetworkName(obj.GetNetwork())
		if err != nil {
			return nil, status.Errorf(codes.InvalidArgument, "network %q is not valid", obj.GetNetwork())
		}
		obj.Network = PtrTo(fmt.Sprintf("https://www.googleapis.com/compute/v1/projects/%s/global/networks/%s", networkName.Project.ID, networkName.Name))
	}

	if obj.Subnetwork != nil {
		subnetworkName, err := s.parseSubnetName(obj.GetSubnetwork())
		if err != nil {
			return nil, status.Errorf(codes.InvalidArgument, "subnetwork %q is not valid", obj.GetSubnetwork())
		}
		obj.Subnetwork = PtrTo(fmt.Sprintf("https://www.googleapis.com/compute/v1/projects/%s/regions/%s/subnetworks/%s", subnetworkName.Project.ID, subnetworkName.Region, subnetworkName.Name))
	}

	// output only fields
	obj.Region = PtrTo(fmt.Sprintf("https://www.googleapis.com/compute/v1/projects/%s/regions/%s", name.Project.ID, name.Region))
	// output only field, this field is only used for internal load balancing.
	if obj.LoadBalancingScheme != nil && *obj.LoadBalancingScheme == "INTERNAL" {
		if obj.ServiceLabel != nil {
			obj.ServiceName = PtrTo(fmt.Sprintf("%s.%s.il4.%s.lb.%s.internal", obj.GetServiceLabel(), name.Name, name.Region, name.Project.ID))
		}
	}

	if err := s.storage.Create(ctx, fqn, obj); err != nil {
		return nil, err
	}

	op := &pb.Operation{
		TargetId:      obj.Id,
		TargetLink:    obj.SelfLink,
		OperationType: PtrTo("insert"),
		User:          PtrTo("user@example.com"),
	}
	return s.startRegionalLRO(ctx, name.Project.ID, name.Region, op, func() (proto.Message, error) {
		return obj, nil
	})
}

func (s *RegionalForwardingRulesV1) Patch(ctx context.Context, req *pb.PatchForwardingRuleRequest) (*pb.Operation, error) {
	reqName := "projects/" + req.GetProject() + "/regions/" + req.GetRegion() + "/forwardingRules/" + req.GetForwardingRule()
	name, err := s.parseRegionalForwardingRuleName(reqName)
	if err != nil {
		return nil, err
	}

	fqn := name.String()

	obj := &pb.ForwardingRule{}

	if err := s.storage.Get(ctx, fqn, obj); err != nil {
		if status.Code(err) == codes.NotFound {
			return nil, status.Errorf(codes.NotFound, "The resource '%s' was not found", fqn)
		}
		return nil, err
	}

	proto.Merge(obj, req.GetForwardingRuleResource())
	// checked GCP log, when AllowGlobalAccess is false, the field will be ignored
	if obj.AllowGlobalAccess != nil && *obj.AllowGlobalAccess == false {
		obj.AllowGlobalAccess = nil
	}

	if err := s.storage.Update(ctx, fqn, obj); err != nil {
		return nil, err
	}

	op := &pb.Operation{
		TargetId:      obj.Id,
		TargetLink:    obj.SelfLink,
		OperationType: PtrTo("patch"),
		User:          PtrTo("user@example.com"),
	}
	return s.startRegionalLRO(ctx, name.Project.ID, name.Region, op, func() (proto.Message, error) {
		return obj, nil
	})
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
		return nil, err
	}

	op := &pb.Operation{
		TargetId:      deleted.Id,
		TargetLink:    deleted.SelfLink,
		OperationType: PtrTo("delete"),
		User:          PtrTo("user@example.com"),
	}
	return s.startRegionalLRO(ctx, name.Project.ID, name.Region, op, func() (proto.Message, error) {
		return deleted, nil
	})
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
		return nil, err
	}

	obj.Labels = req.GetRegionSetLabelsRequestResource().GetLabels()
	if err := s.storage.Update(ctx, fqn, obj); err != nil {
		return nil, err
	}

	op := &pb.Operation{
		TargetId:      obj.Id,
		TargetLink:    obj.SelfLink,
		OperationType: PtrTo("setLabels"),
		User:          PtrTo("user@example.com"),
		// SetLabels operation has EndTime in response
		EndTime: PtrTo("2024-04-01T12:34:56.123456Z"),
		// SetLabels operation finished super fast
		Progress: PtrTo(int32(100)),
		Status:   PtrTo(pb.Operation_DONE),
	}
	return s.startRegionalLRO(ctx, name.Project.ID, name.Region, op, func() (proto.Message, error) {
		return obj, nil
	})
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
		return nil, err
	}

	obj.Target = req.GetTargetReferenceResource().Target
	if err := s.storage.Update(ctx, fqn, obj); err != nil {
		return nil, err
	}

	op := &pb.Operation{
		TargetId:      obj.Id,
		TargetLink:    obj.SelfLink,
		OperationType: PtrTo("SetTarget"),
		User:          PtrTo("user@example.com"),
	}
	return s.startRegionalLRO(ctx, name.Project.ID, name.Region, op, func() (proto.Message, error) {
		return obj, nil
	})
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
		project, err := s.Projects.GetProjectByID(tokens[1])
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
