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
		if status.Code(err) == codes.NotFound {
			return nil, status.Errorf(codes.NotFound, "The resource '%s' was not found", fqn)
		}
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
	obj.SelfLink = PtrTo(buildComputeSelfLink(ctx, fqn))
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
	if *obj.LoadBalancingScheme == "" {
		obj.LoadBalancingScheme = nil
	}
	if isPSCForwardingRule(obj) {
		var num uint64 = 111111111111
		obj.PscConnectionId = &num
		obj.ServiceDirectoryRegistrations = []*pb.ForwardingRuleServiceDirectoryRegistration{
			{
				Namespace:              PtrTo("goog-psc-${networkID}-${networkID}"),
				ServiceDirectoryRegion: PtrTo("us-central1"),
			},
		}

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
		obj.Network = PtrTo(buildComputeSelfLink(ctx, fmt.Sprintf("projects/%s/global/networks/%s", networkName.Project.ID, networkName.Name)))
	}

	if obj.Subnetwork != nil {
		subnetworkName, err := s.parseSubnetName(obj.GetSubnetwork())
		if err != nil {
			return nil, status.Errorf(codes.InvalidArgument, "subnetwork %q is not valid", obj.GetSubnetwork())
		}
		obj.Subnetwork = PtrTo(buildComputeSelfLink(ctx, fmt.Sprintf("projects/%s/regions/%s/subnetworks/%s", subnetworkName.Project.ID, subnetworkName.Region, subnetworkName.Name)))
	}

	// output only field. This field is only used for internal load balancing.
	if obj.LoadBalancingScheme != nil && *obj.LoadBalancingScheme == "INTERNAL" {
		if obj.ServiceLabel != nil {
			obj.ServiceName = PtrTo(fmt.Sprintf("%s.%s.il4.global.lb.%s.internal", obj.GetServiceLabel(), name.Name, name.Project.ID))
		}
	}

	// network field is only used for global internal load balancing.
	// If neither subnetwork nor network field is specified, the default network will be used.
	if obj.Network == nil && obj.Subnetwork == nil && obj.LoadBalancingScheme != nil && *obj.LoadBalancingScheme != "EXTERNAL" {
		obj.Network = PtrTo(buildComputeSelfLink(ctx, fmt.Sprintf("projects/%s/global/networks/default", name.Project.ID)))
	}

	if err := s.storage.Create(ctx, fqn, obj); err != nil {
		return nil, err
	}

	var opType *string
	if isPSCForwardingRule(obj) {
		opType = PtrTo("createPSCServiceEndpoint")
	} else {
		opType = PtrTo("insert")
	}

	op := &pb.Operation{
		TargetId:      obj.Id,
		TargetLink:    obj.SelfLink,
		OperationType: opType,
		User:          PtrTo("user@example.com"),
	}
	return s.startGlobalLRO(ctx, name.Project.ID, op, func() (proto.Message, error) {
		return obj, nil
	})
}

func (s *GlobalForwardingRulesV1) Patch(ctx context.Context, req *pb.PatchGlobalForwardingRuleRequest) (*pb.Operation, error) {
	reqName := "projects/" + req.GetProject() + "/global" + "/forwardingRules/" + req.GetForwardingRule()
	name, err := s.parseGlobalForwardingRuleName(reqName)
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
	return s.startGlobalLRO(ctx, name.Project.ID, op, func() (proto.Message, error) {
		return obj, nil
	})
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

	var opType *string
	if isPSCForwardingRule(deleted) {
		opType = PtrTo("deletePscForwardingRule")
	} else {
		opType = PtrTo("delete")
	}

	op := &pb.Operation{
		TargetId:      deleted.Id,
		TargetLink:    deleted.SelfLink,
		OperationType: opType,
		User:          PtrTo("user@example.com"),
	}
	return s.startGlobalLRO(ctx, name.Project.ID, op, func() (proto.Message, error) {
		return deleted, nil
	})
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
	return s.startGlobalLRO(ctx, name.Project.ID, op, func() (proto.Message, error) {
		return obj, nil
	})
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

	op := &pb.Operation{
		TargetId:      obj.Id,
		TargetLink:    obj.SelfLink,
		OperationType: PtrTo("SetTarget"),
		User:          PtrTo("user@example.com"),
	}
	return s.startGlobalLRO(ctx, name.Project.ID, op, func() (proto.Message, error) {
		return obj, nil
	})
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

func isPSCForwardingRule(obj *pb.ForwardingRule) bool {
	target := *obj.Target
	if target == "all-apis" || target == "vpc-sc" || strings.Contains(target, "/serviceAttachments/") {
		return true
	}
	return false
}
