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

// Adapted from networkfirewallpolicies.go (global network firewall policies) for regional
// ones: projects/{project}/regions/{region}/firewallPolicies/{name}, regional
// operations, the region's URL in the policy, policyType VPC_POLICY when
// unset (the API's documented default), and a new fingerprint on each write.

package mockcompute

import (
	"context"
	"strconv"
	"strings"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/proto"

	"github.com/GoogleCloudPlatform/k8s-config-connector/mockgcp/common/projects"
	pb "github.com/GoogleCloudPlatform/k8s-config-connector/mockgcp/generated/mockgcp/cloud/compute/v1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/mockgcp/pkg/storage"
)

type RegionNetworkFirewallPoliciesV1 struct {
	*MockService
	pb.UnimplementedRegionNetworkFirewallPoliciesServer
}

type regionNetworkFirewallPolicyName struct {
	Project *projects.ProjectData
	Region  string
	Name    string
}

func (n *regionNetworkFirewallPolicyName) String() string {
	return "projects/" + n.Project.ID + "/regions/" + n.Region + "/firewallPolicies/" + n.Name
}

// resolvePolicy finds a policy by name or by numeric ID, as the API does.
func (s *RegionNetworkFirewallPoliciesV1) resolvePolicy(ctx context.Context, project, region, policyNameOrID string) (*pb.FirewallPolicy, *regionNetworkFirewallPolicyName, error) {
	projectObj, err := s.Projects.GetProjectByIDOrNumber(project)
	if err != nil {
		return nil, nil, err
	}

	if id, err := strconv.ParseUint(policyNameOrID, 10, 64); err == nil {
		prefix := (&regionNetworkFirewallPolicyName{Project: projectObj, Region: region}).String()
		var found *pb.FirewallPolicy
		if err := s.storage.List(ctx, (&pb.FirewallPolicy{}).ProtoReflect().Descriptor(), storage.ListOptions{Prefix: prefix}, func(obj proto.Message) error {
			if p := obj.(*pb.FirewallPolicy); p.GetId() == id {
				found = p
			}
			return nil
		}); err != nil {
			return nil, nil, err
		}
		if found == nil {
			return nil, nil, status.Errorf(codes.NotFound, "The resource '%s%s' was not found", prefix, policyNameOrID)
		}
		return found, &regionNetworkFirewallPolicyName{Project: projectObj, Region: region, Name: found.GetName()}, nil
	}

	name := &regionNetworkFirewallPolicyName{Project: projectObj, Region: region, Name: policyNameOrID}
	obj := &pb.FirewallPolicy{}
	if err := s.storage.Get(ctx, name.String(), obj); err != nil {
		if status.Code(err) == codes.NotFound {
			return nil, nil, status.Errorf(codes.NotFound, "The resource '%s' was not found", name.String())
		}
		return nil, nil, err
	}
	return obj, name, nil
}

func (s *RegionNetworkFirewallPoliciesV1) Get(ctx context.Context, req *pb.GetRegionNetworkFirewallPolicyRequest) (*pb.FirewallPolicy, error) {
	obj, _, err := s.resolvePolicy(ctx, req.GetProject(), req.GetRegion(), req.GetFirewallPolicy())
	if err != nil {
		return nil, err
	}
	return obj, nil
}

func (s *RegionNetworkFirewallPoliciesV1) List(ctx context.Context, req *pb.ListRegionNetworkFirewallPoliciesRequest) (*pb.FirewallPolicyList, error) {
	projectObj, err := s.Projects.GetProjectByIDOrNumber(req.GetProject())
	if err != nil {
		return nil, err
	}
	prefix := (&regionNetworkFirewallPolicyName{Project: projectObj, Region: req.GetRegion()}).String()
	out := &pb.FirewallPolicyList{Kind: PtrTo("compute#firewallPolicyList")}
	if err := s.storage.List(ctx, (&pb.FirewallPolicy{}).ProtoReflect().Descriptor(), storage.ListOptions{Prefix: prefix}, func(obj proto.Message) error {
		out.Items = append(out.Items, obj.(*pb.FirewallPolicy))
		return nil
	}); err != nil {
		return nil, err
	}
	return out, nil
}

func (s *RegionNetworkFirewallPoliciesV1) Insert(ctx context.Context, req *pb.InsertRegionNetworkFirewallPolicyRequest) (*pb.Operation, error) {
	id := s.generateID()

	policyName := req.GetFirewallPolicyResource().GetName()
	if policyName == "" {
		return nil, status.Errorf(codes.InvalidArgument, "Invalid value for field 'resource.name': ''. Must be a match of regex '[a-z](?:[-a-z0-9]{0,61}[a-z0-9])?'")
	}
	if strings.Contains(policyName, "/") {
		return nil, status.Errorf(codes.InvalidArgument, "Invalid value for field 'resource.name': '%s'.", policyName)
	}

	projectObj, err := s.Projects.GetProjectByIDOrNumber(req.GetProject())
	if err != nil {
		return nil, err
	}
	name := &regionNetworkFirewallPolicyName{Project: projectObj, Region: req.GetRegion(), Name: policyName}
	fqn := name.String()

	obj := proto.CloneOf(req.GetFirewallPolicyResource())
	obj.SelfLink = PtrTo(BuildComputeSelfLink(ctx, fqn))
	obj.SelfLinkWithId = PtrTo(BuildComputeSelfLink(ctx, "projects/"+projectObj.ID+"/regions/"+req.GetRegion()+"/firewallPolicies/"+strconv.FormatUint(id, 10)))
	obj.Region = PtrTo(BuildComputeSelfLink(ctx, "projects/"+projectObj.ID+"/regions/"+req.GetRegion()))
	obj.RuleTupleCount = PtrTo(int32(8))
	obj.Id = PtrTo(id)
	obj.CreationTimestamp = PtrTo(s.nowString())
	obj.Kind = PtrTo("compute#firewallPolicy")
	if obj.PolicyType == nil {
		obj.PolicyType = PtrTo("VPC_POLICY")
	}
	if obj.Rules == nil {
		populateDefaultRules(obj, false)
	}
	obj.Fingerprint = nil
	obj.Fingerprint = PtrTo(computeFingerprint(obj))

	if err := s.storage.Create(ctx, fqn, obj); err != nil {
		return nil, err
	}

	op := &pb.Operation{
		TargetId:      obj.Id,
		TargetLink:    obj.SelfLink,
		OperationType: PtrTo("compute.regionNetworkFirewallPolicies.insert"),
		User:          PtrTo("user@example.com"),
	}
	return s.startRegionalLRO(ctx, projectObj.ID, req.GetRegion(), op, func() (proto.Message, error) {
		return obj, nil
	})
}

func (s *RegionNetworkFirewallPoliciesV1) Patch(ctx context.Context, req *pb.PatchRegionNetworkFirewallPolicyRequest) (*pb.Operation, error) {
	obj, name, err := s.resolvePolicy(ctx, req.GetProject(), req.GetRegion(), req.GetFirewallPolicy())
	if err != nil {
		return nil, err
	}

	patch := proto.CloneOf(req.GetFirewallPolicyResource())
	patch.Fingerprint = nil
	proto.Merge(obj, patch)
	obj.Fingerprint = nil
	obj.Fingerprint = PtrTo(computeFingerprint(obj))

	if err := s.storage.Update(ctx, name.String(), obj); err != nil {
		return nil, err
	}

	op := &pb.Operation{
		TargetId:      obj.Id,
		TargetLink:    obj.SelfLink,
		OperationType: PtrTo("compute.regionNetworkFirewallPolicies.patch"),
		User:          PtrTo("user@example.com"),
	}
	return s.startRegionalLRO(ctx, name.Project.ID, name.Region, op, func() (proto.Message, error) {
		return obj, nil
	})
}

func (s *RegionNetworkFirewallPoliciesV1) Delete(ctx context.Context, req *pb.DeleteRegionNetworkFirewallPolicyRequest) (*pb.Operation, error) {
	_, name, err := s.resolvePolicy(ctx, req.GetProject(), req.GetRegion(), req.GetFirewallPolicy())
	if err != nil {
		return nil, err
	}

	deleted := &pb.FirewallPolicy{}
	if err := s.storage.Delete(ctx, name.String(), deleted); err != nil {
		return nil, err
	}

	op := &pb.Operation{
		TargetId:      deleted.Id,
		TargetLink:    deleted.SelfLink,
		OperationType: PtrTo("compute.regionNetworkFirewallPolicies.delete"),
		User:          PtrTo("user@example.com"),
	}
	return s.startRegionalLRO(ctx, name.Project.ID, name.Region, op, func() (proto.Message, error) {
		return deleted, nil
	})
}
