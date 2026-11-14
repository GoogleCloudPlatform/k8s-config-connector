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

package mockcompute

import (
	"context"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/proto"

	"github.com/GoogleCloudPlatform/k8s-config-connector/mockgcp/common/projects"
	pb "github.com/GoogleCloudPlatform/k8s-config-connector/mockgcp/generated/mockgcp/cloud/compute/v1"
)

type RegionHealthAggregationPoliciesV1 struct {
	*MockService
	pb.UnimplementedRegionHealthAggregationPoliciesServer
}

type regionalHealthAggregationPolicyName struct {
	Project *projects.ProjectData
	Region  string
	Name    string
}

func (n *regionalHealthAggregationPolicyName) String() string {
	return "projects/" + n.Project.ID + "/regions/" + n.Region + "/healthAggregationPolicies/" + n.Name
}

func (s *MockService) buildRegionalHealthAggregationPolicyName(projectID, region, name string) (*regionalHealthAggregationPolicyName, error) {
	project, err := s.Projects.GetProjectByID(projectID)
	if err != nil {
		return nil, err
	}

	return &regionalHealthAggregationPolicyName{
		Project: project,
		Region:  region,
		Name:    name,
	}, nil
}

func (s *RegionHealthAggregationPoliciesV1) Get(ctx context.Context, req *pb.GetRegionHealthAggregationPolicyRequest) (*pb.HealthAggregationPolicy, error) {
	name, err := s.buildRegionalHealthAggregationPolicyName(req.GetProject(), req.GetRegion(), req.GetHealthAggregationPolicy())
	if err != nil {
		return nil, err
	}

	fqn := name.String()

	obj := &pb.HealthAggregationPolicy{}
	if err := s.storage.Get(ctx, fqn, obj); err != nil {
		if status.Code(err) == codes.NotFound {
			return nil, status.Errorf(codes.NotFound, "The resource '%s' was not found", fqn)
		}
		return nil, err
	}

	return obj, nil
}

func (s *RegionHealthAggregationPoliciesV1) Insert(ctx context.Context, req *pb.InsertRegionHealthAggregationPolicyRequest) (*pb.Operation, error) {
	name, err := s.buildRegionalHealthAggregationPolicyName(req.GetProject(), req.GetRegion(), req.GetHealthAggregationPolicyResource().GetName())
	if err != nil {
		return nil, err
	}

	fqn := name.String()

	id := s.generateID()

	obj := proto.CloneOf(req.GetHealthAggregationPolicyResource())
	obj.SelfLink = PtrTo(BuildComputeSelfLink(ctx, fqn))
	obj.SelfLinkWithId = PtrTo(BuildComputeSelfLink(ctx, fqn))
	obj.CreationTimestamp = PtrTo(s.nowString())
	obj.Id = &id
	obj.Kind = PtrTo("compute#healthAggregationPolicy")
	obj.Name = PtrTo(name.Name)
	obj.Region = PtrTo(BuildComputeSelfLink(ctx, "projects/"+name.Project.ID+"/regions/"+name.Region))

	s.populateDefaults(obj)

	obj.Fingerprint = PtrTo(computeFingerprint(obj))

	if err := s.storage.Create(ctx, fqn, obj); err != nil {
		return nil, err
	}

	op := &pb.Operation{
		TargetId:      obj.Id,
		TargetLink:    obj.SelfLink,
		OperationType: PtrTo("compute.regionHealthAggregationPolicies.insert"),
		User:          PtrTo("user@example.com"),
	}
	return s.startRegionalLRO(ctx, name.Project.ID, name.Region, op, func() (proto.Message, error) {
		return obj, nil
	})
}

func (s *RegionHealthAggregationPoliciesV1) Patch(ctx context.Context, req *pb.PatchRegionHealthAggregationPolicyRequest) (*pb.Operation, error) {
	name, err := s.buildRegionalHealthAggregationPolicyName(req.GetProject(), req.GetRegion(), req.GetHealthAggregationPolicy())
	if err != nil {
		return nil, err
	}

	fqn := name.String()
	obj := &pb.HealthAggregationPolicy{}
	if err := s.storage.Get(ctx, fqn, obj); err != nil {
		if status.Code(err) == codes.NotFound {
			return nil, status.Errorf(codes.NotFound, "The resource '%s' was not found", fqn)
		}
		return nil, err
	}

	patchObj := req.GetHealthAggregationPolicyResource()
	if patchObj == nil {
		patchObj = &pb.HealthAggregationPolicy{}
	}

	if patchObj.GetFingerprint() != "" && patchObj.GetFingerprint() != obj.GetFingerprint() {
		return nil, status.Errorf(codes.FailedPrecondition, "Condition Not Met: fingerprint mismatch")
	}

	if patchObj.Description != nil {
		obj.Description = patchObj.Description
	}
	if patchObj.HealthyPercentThreshold != nil {
		obj.HealthyPercentThreshold = patchObj.HealthyPercentThreshold
	}
	if patchObj.MinHealthyThreshold != nil {
		obj.MinHealthyThreshold = patchObj.MinHealthyThreshold
	}

	obj.Fingerprint = nil
	obj.Fingerprint = PtrTo(computeFingerprint(obj))

	if err := s.storage.Update(ctx, fqn, obj); err != nil {
		return nil, err
	}

	op := &pb.Operation{
		TargetId:      obj.Id,
		TargetLink:    obj.SelfLink,
		OperationType: PtrTo("compute.healthAggregationPolicies.patch"),
		User:          PtrTo("user@example.com"),
	}
	return s.startRegionalLRO(ctx, name.Project.ID, name.Region, op, func() (proto.Message, error) {
		return obj, nil
	})
}

func (s *RegionHealthAggregationPoliciesV1) Delete(ctx context.Context, req *pb.DeleteRegionHealthAggregationPolicyRequest) (*pb.Operation, error) {
	name, err := s.buildRegionalHealthAggregationPolicyName(req.GetProject(), req.GetRegion(), req.GetHealthAggregationPolicy())
	if err != nil {
		return nil, err
	}

	fqn := name.String()

	deleted := &pb.HealthAggregationPolicy{}
	if err := s.storage.Delete(ctx, fqn, deleted); err != nil {
		if status.Code(err) == codes.NotFound {
			return nil, status.Errorf(codes.NotFound, "The resource '%s' was not found", fqn)
		}
		return nil, err
	}

	op := &pb.Operation{
		TargetId:      deleted.Id,
		TargetLink:    deleted.SelfLink,
		OperationType: PtrTo("compute.regionHealthAggregationPolicies.delete"),
		User:          PtrTo("user@example.com"),
	}
	return s.startRegionalLRO(ctx, name.Project.ID, name.Region, op, func() (proto.Message, error) {
		return deleted, nil
	})
}

func (s *RegionHealthAggregationPoliciesV1) populateDefaults(obj *pb.HealthAggregationPolicy) {
	if obj.HealthyPercentThreshold == nil {
		obj.HealthyPercentThreshold = PtrTo(uint32(60))
	}
	if obj.MinHealthyThreshold == nil {
		obj.MinHealthyThreshold = PtrTo(uint32(1))
	}
}
