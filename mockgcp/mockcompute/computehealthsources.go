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

type RegionHealthSourcesV1 struct {
	*MockService
	pb.UnimplementedRegionHealthSourcesServer
}

type regionalHealthSourceName struct {
	Project *projects.ProjectData
	Region  string
	Name    string
}

func (n *regionalHealthSourceName) String() string {
	return "projects/" + n.Project.ID + "/regions/" + n.Region + "/healthSources/" + n.Name
}

func (s *MockService) buildRegionalHealthSourceName(projectID, region, name string) (*regionalHealthSourceName, error) {
	project, err := s.Projects.GetProjectByID(projectID)
	if err != nil {
		return nil, err
	}

	return &regionalHealthSourceName{
		Project: project,
		Region:  region,
		Name:    name,
	}, nil
}

func (s *RegionHealthSourcesV1) Get(ctx context.Context, req *pb.GetRegionHealthSourceRequest) (*pb.HealthSource, error) {
	name, err := s.buildRegionalHealthSourceName(req.GetProject(), req.GetRegion(), req.GetHealthSource())
	if err != nil {
		return nil, err
	}

	fqn := name.String()

	obj := &pb.HealthSource{}
	if err := s.storage.Get(ctx, fqn, obj); err != nil {
		if status.Code(err) == codes.NotFound {
			return nil, status.Errorf(codes.NotFound, "The resource '%s' was not found", fqn)
		}
		return nil, err
	}

	return obj, nil
}

func (s *RegionHealthSourcesV1) List(ctx context.Context, req *pb.ListRegionHealthSourcesRequest) (*pb.HealthSourceList, error) {
	project, err := s.Projects.GetProjectByID(req.GetProject())
	if err != nil {
		return nil, err
	}

	prefix := fmt.Sprintf("projects/%s/regions/%s/healthSources/", project.ID, req.GetRegion())

	response := &pb.HealthSourceList{
		Id:       PtrTo(fmt.Sprintf("%d", s.generateID())),
		Kind:     PtrTo("compute#healthSourceList"),
		SelfLink: PtrTo(BuildComputeSelfLink(ctx, fmt.Sprintf("projects/%s/regions/%s/healthSources", project.ID, req.GetRegion()))),
	}

	findKind := (&pb.HealthSource{}).ProtoReflect().Descriptor()
	if err := s.storage.List(ctx, findKind, storage.ListOptions{Prefix: prefix}, func(obj proto.Message) error {
		item := obj.(*pb.HealthSource)
		response.Items = append(response.Items, item)
		return nil
	}); err != nil {
		return nil, err
	}

	return response, nil
}

func (s *RegionHealthSourcesV1) Insert(ctx context.Context, req *pb.InsertRegionHealthSourceRequest) (*pb.Operation, error) {
	if req.GetHealthSourceResource() == nil {
		return nil, status.Errorf(codes.InvalidArgument, "HealthSourceResource must be provided")
	}
	name, err := s.buildRegionalHealthSourceName(req.GetProject(), req.GetRegion(), req.GetHealthSourceResource().GetName())
	if err != nil {
		return nil, err
	}

	fqn := name.String()

	id := s.generateID()

	obj := proto.CloneOf(req.GetHealthSourceResource())
	obj.SelfLink = PtrTo(BuildComputeSelfLink(ctx, fqn))
	obj.SelfLinkWithId = PtrTo(BuildComputeSelfLink(ctx, fqn))
	obj.CreationTimestamp = PtrTo(s.nowString())
	obj.Id = &id
	obj.Kind = PtrTo("compute#healthSource")
	obj.Name = PtrTo(name.Name)
	obj.Region = PtrTo(BuildComputeSelfLink(ctx, "projects/"+name.Project.ID+"/regions/"+name.Region))

	if obj.HealthAggregationPolicy != nil {
		obj.HealthAggregationPolicy = PtrTo(ExpandComputeLink(ctx, *obj.HealthAggregationPolicy))
	}
	for i, src := range obj.Sources {
		obj.Sources[i] = ExpandComputeLink(ctx, src)
	}

	s.populateDefaults(obj)

	obj.Fingerprint = PtrTo(computeFingerprint(obj))

	if err := s.storage.Create(ctx, fqn, obj); err != nil {
		return nil, err
	}

	op := &pb.Operation{
		TargetId:      obj.Id,
		TargetLink:    obj.SelfLink,
		OperationType: PtrTo("compute.healthSources.insert"),
		User:          PtrTo("user@example.com"),
	}
	return s.startRegionalLRO(ctx, name.Project.ID, name.Region, op, func() (proto.Message, error) {
		return obj, nil
	})
}

func (s *RegionHealthSourcesV1) Patch(ctx context.Context, req *pb.PatchRegionHealthSourceRequest) (*pb.Operation, error) {
	name, err := s.buildRegionalHealthSourceName(req.GetProject(), req.GetRegion(), req.GetHealthSource())
	if err != nil {
		return nil, err
	}

	fqn := name.String()
	obj := &pb.HealthSource{}
	if err := s.storage.Get(ctx, fqn, obj); err != nil {
		if status.Code(err) == codes.NotFound {
			return nil, status.Errorf(codes.NotFound, "The resource '%s' was not found", fqn)
		}
		return nil, err
	}

	patchObj := req.GetHealthSourceResource()
	if patchObj == nil {
		patchObj = &pb.HealthSource{}
	}

	if patchObj.GetFingerprint() != "" && patchObj.GetFingerprint() != obj.GetFingerprint() {
		return nil, status.Errorf(codes.FailedPrecondition, "Condition Not Met: fingerprint mismatch")
	}

	if patchObj.Description != nil {
		obj.Description = patchObj.Description
	}
	if patchObj.HealthAggregationPolicy != nil {
		obj.HealthAggregationPolicy = PtrTo(ExpandComputeLink(ctx, patchObj.GetHealthAggregationPolicy()))
	}
	if len(patchObj.Sources) > 0 {
		var sources []string
		for _, src := range patchObj.Sources {
			sources = append(sources, ExpandComputeLink(ctx, src))
		}
		obj.Sources = sources
	}

	obj.Fingerprint = nil
	obj.Fingerprint = PtrTo(computeFingerprint(obj))

	if err := s.storage.Update(ctx, fqn, obj); err != nil {
		return nil, err
	}

	op := &pb.Operation{
		TargetId:      obj.Id,
		TargetLink:    obj.SelfLink,
		OperationType: PtrTo("compute.healthSources.patch"),
		User:          PtrTo("user@example.com"),
	}
	return s.startRegionalLRO(ctx, name.Project.ID, name.Region, op, func() (proto.Message, error) {
		return obj, nil
	})
}

func (s *RegionHealthSourcesV1) Delete(ctx context.Context, req *pb.DeleteRegionHealthSourceRequest) (*pb.Operation, error) {
	name, err := s.buildRegionalHealthSourceName(req.GetProject(), req.GetRegion(), req.GetHealthSource())
	if err != nil {
		return nil, err
	}

	fqn := name.String()

	deleted := &pb.HealthSource{}
	if err := s.storage.Delete(ctx, fqn, deleted); err != nil {
		if status.Code(err) == codes.NotFound {
			return nil, status.Errorf(codes.NotFound, "The resource '%s' was not found", fqn)
		}
		return nil, err
	}

	op := &pb.Operation{
		TargetId:      deleted.Id,
		TargetLink:    deleted.SelfLink,
		OperationType: PtrTo("compute.healthSources.delete"),
		User:          PtrTo("user@example.com"),
	}
	return s.startRegionalLRO(ctx, name.Project.ID, name.Region, op, func() (proto.Message, error) {
		return deleted, nil
	})
}

func (s *RegionHealthSourcesV1) populateDefaults(obj *pb.HealthSource) {
	if obj.SourceType == nil {
		obj.SourceType = PtrTo("BACKEND_SERVICE")
	}
}
