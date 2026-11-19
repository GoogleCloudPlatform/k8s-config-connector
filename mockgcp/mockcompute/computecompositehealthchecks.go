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
	"fmt"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/proto"

	"github.com/GoogleCloudPlatform/k8s-config-connector/mockgcp/common/projects"
	pb "github.com/GoogleCloudPlatform/k8s-config-connector/mockgcp/generated/mockgcp/cloud/compute/v1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/mockgcp/pkg/storage"
)

type RegionCompositeHealthChecksV1 struct {
	*MockService
	pb.UnimplementedRegionCompositeHealthChecksServer
}

type regionalCompositeHealthCheckName struct {
	Project *projects.ProjectData
	Region  string
	Name    string
}

func (n *regionalCompositeHealthCheckName) String() string {
	return "projects/" + n.Project.ID + "/regions/" + n.Region + "/compositeHealthChecks/" + n.Name
}

func (s *MockService) buildRegionalCompositeHealthCheckName(projectID, region, name string) (*regionalCompositeHealthCheckName, error) {
	project, err := s.Projects.GetProjectByID(projectID)
	if err != nil {
		return nil, err
	}

	return &regionalCompositeHealthCheckName{
		Project: project,
		Region:  region,
		Name:    name,
	}, nil
}

func (s *RegionCompositeHealthChecksV1) Get(ctx context.Context, req *pb.GetRegionCompositeHealthCheckRequest) (*pb.CompositeHealthCheck, error) {
	name, err := s.buildRegionalCompositeHealthCheckName(req.GetProject(), req.GetRegion(), req.GetCompositeHealthCheck())
	if err != nil {
		return nil, err
	}

	fqn := name.String()

	obj := &pb.CompositeHealthCheck{}
	if err := s.storage.Get(ctx, fqn, obj); err != nil {
		if status.Code(err) == codes.NotFound {
			return nil, status.Errorf(codes.NotFound, "The resource '%s' was not found", fqn)
		}
		return nil, err
	}

	return obj, nil
}

func (s *RegionCompositeHealthChecksV1) List(ctx context.Context, req *pb.ListRegionCompositeHealthChecksRequest) (*pb.CompositeHealthCheckList, error) {
	project, err := s.Projects.GetProjectByID(req.GetProject())
	if err != nil {
		return nil, err
	}

	prefix := fmt.Sprintf("projects/%s/regions/%s/compositeHealthChecks/", project.ID, req.GetRegion())

	response := &pb.CompositeHealthCheckList{
		Id:       PtrTo(fmt.Sprintf("%d", s.generateID())),
		Kind:     PtrTo("compute#compositeHealthCheckList"),
		SelfLink: PtrTo(BuildComputeSelfLink(ctx, fmt.Sprintf("projects/%s/regions/%s/compositeHealthChecks", project.ID, req.GetRegion()))),
	}

	findKind := (&pb.CompositeHealthCheck{}).ProtoReflect().Descriptor()
	if err := s.storage.List(ctx, findKind, storage.ListOptions{Prefix: prefix}, func(obj proto.Message) error {
		item := obj.(*pb.CompositeHealthCheck)
		response.Items = append(response.Items, item)
		return nil
	}); err != nil {
		return nil, err
	}

	return response, nil
}

func (s *RegionCompositeHealthChecksV1) Insert(ctx context.Context, req *pb.InsertRegionCompositeHealthCheckRequest) (*pb.Operation, error) {
	if req.GetCompositeHealthCheckResource() == nil {
		return nil, status.Errorf(codes.InvalidArgument, "CompositeHealthCheckResource must be provided")
	}
	name, err := s.buildRegionalCompositeHealthCheckName(req.GetProject(), req.GetRegion(), req.GetCompositeHealthCheckResource().GetName())
	if err != nil {
		return nil, err
	}

	fqn := name.String()

	id := s.generateID()

	obj := proto.CloneOf(req.GetCompositeHealthCheckResource())
	obj.SelfLink = PtrTo(BuildComputeSelfLink(ctx, fqn))
	obj.SelfLinkWithId = PtrTo(BuildComputeSelfLink(ctx, fqn))
	obj.CreationTimestamp = PtrTo(s.nowString())
	obj.Id = &id
	obj.Kind = PtrTo("compute#compositeHealthCheck")
	obj.Name = PtrTo(name.Name)
	obj.Region = PtrTo(BuildComputeSelfLink(ctx, "projects/"+name.Project.ID+"/regions/"+name.Region))

	if obj.HealthDestination != nil {
		obj.HealthDestination = PtrTo(ExpandComputeLink(ctx, *obj.HealthDestination))
	}
	for i, src := range obj.HealthSources {
		obj.HealthSources[i] = ExpandComputeLink(ctx, src)
	}

	s.populateDefaults(obj)

	obj.Fingerprint = PtrTo(computeFingerprint(obj))

	if err := s.storage.Create(ctx, fqn, obj); err != nil {
		return nil, err
	}

	op := &pb.Operation{
		TargetId:      obj.Id,
		TargetLink:    obj.SelfLink,
		OperationType: PtrTo("compute.compositeHealthChecks.insert"),
		User:          PtrTo("user@example.com"),
	}
	return s.startRegionalLRO(ctx, name.Project.ID, name.Region, op, func() (proto.Message, error) {
		return obj, nil
	})
}

func (s *RegionCompositeHealthChecksV1) Patch(ctx context.Context, req *pb.PatchRegionCompositeHealthCheckRequest) (*pb.Operation, error) {
	name, err := s.buildRegionalCompositeHealthCheckName(req.GetProject(), req.GetRegion(), req.GetCompositeHealthCheck())
	if err != nil {
		return nil, err
	}

	fqn := name.String()
	obj := &pb.CompositeHealthCheck{}
	if err := s.storage.Get(ctx, fqn, obj); err != nil {
		if status.Code(err) == codes.NotFound {
			return nil, status.Errorf(codes.NotFound, "The resource '%s' was not found", fqn)
		}
		return nil, err
	}

	patchObj := req.GetCompositeHealthCheckResource()
	if patchObj == nil {
		patchObj = &pb.CompositeHealthCheck{}
	}

	if patchObj.GetFingerprint() != "" && patchObj.GetFingerprint() != obj.GetFingerprint() {
		return nil, status.Errorf(codes.FailedPrecondition, "Condition Not Met: fingerprint mismatch")
	}

	if patchObj.Description != nil {
		obj.Description = patchObj.Description
	}
	if patchObj.HealthDestination != nil {
		obj.HealthDestination = PtrTo(ExpandComputeLink(ctx, patchObj.GetHealthDestination()))
	}
	if len(patchObj.HealthSources) > 0 {
		var sources []string
		for _, src := range patchObj.HealthSources {
			sources = append(sources, ExpandComputeLink(ctx, src))
		}
		obj.HealthSources = sources
	}

	s.populateDefaults(obj)

	obj.Fingerprint = nil
	obj.Fingerprint = PtrTo(computeFingerprint(obj))

	if err := s.storage.Update(ctx, fqn, obj); err != nil {
		return nil, err
	}

	op := &pb.Operation{
		TargetId:      obj.Id,
		TargetLink:    obj.SelfLink,
		OperationType: PtrTo("compute.compositeHealthChecks.patch"),
		User:          PtrTo("user@example.com"),
	}
	return s.startRegionalLRO(ctx, name.Project.ID, name.Region, op, func() (proto.Message, error) {
		return obj, nil
	})
}

func (s *RegionCompositeHealthChecksV1) Delete(ctx context.Context, req *pb.DeleteRegionCompositeHealthCheckRequest) (*pb.Operation, error) {
	name, err := s.buildRegionalCompositeHealthCheckName(req.GetProject(), req.GetRegion(), req.GetCompositeHealthCheck())
	if err != nil {
		return nil, err
	}

	fqn := name.String()

	deleted := &pb.CompositeHealthCheck{}
	if err := s.storage.Delete(ctx, fqn, deleted); err != nil {
		if status.Code(err) == codes.NotFound {
			return nil, status.Errorf(codes.NotFound, "The resource '%s' was not found", fqn)
		}
		return nil, err
	}

	op := &pb.Operation{
		TargetId:      deleted.Id,
		TargetLink:    deleted.SelfLink,
		OperationType: PtrTo("compute.compositeHealthChecks.delete"),
		User:          PtrTo("user@example.com"),
	}
	return s.startRegionalLRO(ctx, name.Project.ID, name.Region, op, func() (proto.Message, error) {
		return deleted, nil
	})
}

func (s *RegionCompositeHealthChecksV1) populateDefaults(obj *pb.CompositeHealthCheck) {
}
