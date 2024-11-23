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

package mockcompute

import (
	"context"

	"google.golang.org/protobuf/proto"

	"github.com/GoogleCloudPlatform/k8s-config-connector/mockgcp/common/projects"
	pb "github.com/GoogleCloudPlatform/k8s-config-connector/mockgcp/generated/mockgcp/cloud/compute/v1"
)

type RegionalHealthCheckV1 struct {
	*MockService
	pb.UnimplementedRegionHealthChecksServer
}

func (s *RegionalHealthCheckV1) Get(ctx context.Context, req *pb.GetRegionHealthCheckRequest) (*pb.HealthCheck, error) {
	name, err := s.buildRegionalHealthCheckName(req.GetProject(), req.GetRegion(), req.GetHealthCheck())
	if err != nil {
		return nil, err
	}

	fqn := name.String()

	obj := &pb.HealthCheck{}
	if err := s.storage.Get(ctx, fqn, obj); err != nil {
		return nil, err
	}

	return obj, nil
}

func (s *RegionalHealthCheckV1) Insert(ctx context.Context, req *pb.InsertRegionHealthCheckRequest) (*pb.Operation, error) {
	name, err := s.buildRegionalHealthCheckName(req.GetProject(), req.GetRegion(), req.GetHealthCheckResource().GetName())
	if err != nil {
		return nil, err
	}

	fqn := name.String()

	id := s.generateID()

	obj := proto.Clone(req.GetHealthCheckResource()).(*pb.HealthCheck)
	obj.SelfLink = PtrTo(buildComputeSelfLink(ctx, fqn))
	obj.CreationTimestamp = PtrTo(s.nowString())
	obj.Id = &id
	obj.Kind = PtrTo("compute#healthCheck")

	if err := s.storage.Create(ctx, fqn, obj); err != nil {
		return nil, err
	}

	return s.newLRO(ctx, name.Project.ID)
}

// Updates a HealthCheck resource in the specified project using the data included in the request.
// This method supports PATCH semantics and uses the JSON merge patch format and processing rules.
func (s *RegionalHealthCheckV1) Patch(ctx context.Context, req *pb.PatchRegionHealthCheckRequest) (*pb.Operation, error) {
	name, err := s.buildRegionalHealthCheckName(req.GetProject(), req.GetRegion(), req.GetHealthCheck())
	if err != nil {
		return nil, err
	}

	fqn := name.String()
	obj := &pb.HealthCheck{}
	if err := s.storage.Get(ctx, fqn, obj); err != nil {
		return nil, err
	}

	// TODO: Implement helper to implement the full rules here
	proto.Merge(obj, req.GetHealthCheckResource())

	if err := s.storage.Update(ctx, fqn, obj); err != nil {
		return nil, err
	}

	return s.newLRO(ctx, name.Project.ID)
}

// Updates a HealthCheck resource in the specified project using the data included in the request.
func (s *RegionalHealthCheckV1) Update(ctx context.Context, req *pb.UpdateRegionHealthCheckRequest) (*pb.Operation, error) {
	name, err := s.buildRegionalHealthCheckName(req.GetProject(), req.GetRegion(), req.GetHealthCheck())
	if err != nil {
		return nil, err
	}

	fqn := name.String()
	obj := &pb.HealthCheck{}
	if err := s.storage.Get(ctx, fqn, obj); err != nil {
		return nil, err
	}

	// TODO: Implement helper to implement the full rules here
	proto.Merge(obj, req.GetHealthCheckResource())

	if err := s.storage.Update(ctx, fqn, obj); err != nil {
		return nil, err
	}

	return s.newLRO(ctx, name.Project.ID)
}

func (s *RegionalHealthCheckV1) Delete(ctx context.Context, req *pb.DeleteRegionHealthCheckRequest) (*pb.Operation, error) {
	name, err := s.buildRegionalHealthCheckName(req.GetProject(), req.GetRegion(), req.GetHealthCheck())
	if err != nil {
		return nil, err
	}

	fqn := name.String()

	deleted := &pb.HealthCheck{}
	if err := s.storage.Delete(ctx, fqn, deleted); err != nil {
		return nil, err
	}

	return s.newLRO(ctx, name.Project.ID)
}

type regionalHealthCheckName struct {
	Project *projects.ProjectData
	Region  string
	Name    string
}

func (n *regionalHealthCheckName) String() string {
	return "projects/" + n.Project.ID + "/regions/" + n.Region + "/healthChecks/" + n.Name
}

func (s *MockService) buildRegionalHealthCheckName(projectID, region, name string) (*regionalHealthCheckName, error) {
	project, err := s.Projects.GetProjectByID(projectID)
	if err != nil {
		return nil, err
	}

	return &regionalHealthCheckName{
		Project: project,
		Region:  region,
		Name:    name,
	}, nil
}
