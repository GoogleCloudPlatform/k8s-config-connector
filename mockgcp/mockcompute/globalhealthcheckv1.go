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

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/proto"

	"github.com/GoogleCloudPlatform/k8s-config-connector/mockgcp/common/projects"
	pb "github.com/GoogleCloudPlatform/k8s-config-connector/mockgcp/generated/mockgcp/cloud/compute/v1"
)

type GlobalHealthCheckV1 struct {
	*MockService
	pb.UnimplementedHealthChecksServer
}

func (s *GlobalHealthCheckV1) Get(ctx context.Context, req *pb.GetHealthCheckRequest) (*pb.HealthCheck, error) {
	name, err := s.buildGlobalHealthCheckName(req.GetProject(), req.GetHealthCheck())
	if err != nil {
		return nil, err
	}

	fqn := name.String()

	obj := &pb.HealthCheck{}
	if err := s.storage.Get(ctx, fqn, obj); err != nil {
		if status.Code(err) == codes.NotFound {
			return nil, status.Errorf(codes.NotFound, "The resource '%s' was not found", fqn)
		}
		return nil, err
	}

	return obj, nil
}

func (s *GlobalHealthCheckV1) Insert(ctx context.Context, req *pb.InsertHealthCheckRequest) (*pb.Operation, error) {
	name, err := s.buildGlobalHealthCheckName(req.GetProject(), req.GetHealthCheckResource().GetName())
	if err != nil {
		return nil, err
	}

	fqn := name.String()

	id := s.generateID()

	obj := proto.CloneOf(req.GetHealthCheckResource())
	obj.SelfLink = PtrTo(BuildComputeSelfLink(ctx, fqn))
	obj.CreationTimestamp = PtrTo(s.nowString())
	obj.Id = &id
	obj.Kind = PtrTo("compute#healthCheck")

	s.populateHealthCheckDefaults(obj)

	if err := s.storage.Create(ctx, fqn, obj); err != nil {
		return nil, err
	}

	op := &pb.Operation{
		TargetId:      obj.Id,
		TargetLink:    obj.SelfLink,
		OperationType: PtrTo("insert"),
		User:          PtrTo("user@example.com"),
	}
	return s.startGlobalLRO(ctx, name.Project.ID, op, func() (proto.Message, error) {
		return obj, nil
	})
}

// Updates a HealthCheck resource in the specified project using the data included in the request.
// This method supports PATCH semantics and uses the JSON merge patch format and processing rules.
func (s *GlobalHealthCheckV1) Patch(ctx context.Context, req *pb.PatchHealthCheckRequest) (*pb.Operation, error) {
	name, err := s.buildGlobalHealthCheckName(req.GetProject(), req.GetHealthCheck())
	if err != nil {
		return nil, err
	}

	fqn := name.String()
	obj := &pb.HealthCheck{}
	if err := s.storage.Get(ctx, fqn, obj); err != nil {
		return nil, err
	}

	// TODO: Implement helper to implement the full rules here
	if req.GetHealthCheckResource() != nil && req.GetHealthCheckResource().SourceRegions != nil {
		obj.SourceRegions = nil
	}
	proto.Merge(obj, req.GetHealthCheckResource())

	s.populateHealthCheckDefaults(obj)

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

// Updates a HealthCheck resource in the specified project using the data included in the request.
func (s *GlobalHealthCheckV1) Update(ctx context.Context, req *pb.UpdateHealthCheckRequest) (*pb.Operation, error) {
	name, err := s.buildGlobalHealthCheckName(req.GetProject(), req.GetHealthCheck())
	if err != nil {
		return nil, err
	}

	fqn := name.String()
	obj := &pb.HealthCheck{}
	if err := s.storage.Get(ctx, fqn, obj); err != nil {
		return nil, err
	}

	// For PUT (Update), we replace the resource but preserve system fields
	id := obj.Id
	selfLink := obj.SelfLink
	creationTimestamp := obj.CreationTimestamp
	kind := obj.Kind

	updatedObj := proto.Clone(req.GetHealthCheckResource()).(*pb.HealthCheck)
	updatedObj.Id = id
	updatedObj.SelfLink = selfLink
	updatedObj.CreationTimestamp = creationTimestamp
	updatedObj.Kind = kind

	s.populateHealthCheckDefaults(updatedObj)

	if err := s.storage.Update(ctx, fqn, updatedObj); err != nil {
		return nil, err
	}

	op := &pb.Operation{
		TargetId:      updatedObj.Id,
		TargetLink:    updatedObj.SelfLink,
		OperationType: PtrTo("update"),
		User:          PtrTo("user@example.com"),
	}
	return s.startGlobalLRO(ctx, name.Project.ID, op, func() (proto.Message, error) {
		return updatedObj, nil
	})
}

func (s *GlobalHealthCheckV1) Delete(ctx context.Context, req *pb.DeleteHealthCheckRequest) (*pb.Operation, error) {
	name, err := s.buildGlobalHealthCheckName(req.GetProject(), req.GetHealthCheck())
	if err != nil {
		return nil, err
	}

	fqn := name.String()

	deleted := &pb.HealthCheck{}
	if err := s.storage.Delete(ctx, fqn, deleted); err != nil {
		return nil, err
	}

	op := &pb.Operation{
		TargetId:      deleted.Id,
		TargetLink:    deleted.SelfLink,
		OperationType: PtrTo("delete"),
		User:          PtrTo("user@example.com"),
	}
	return s.startGlobalLRO(ctx, name.Project.ID, op, func() (proto.Message, error) {
		return deleted, nil
	})
}

type globalHealthCheckName struct {
	Project *projects.ProjectData
	Name    string
}

func (n *globalHealthCheckName) String() string {
	return "projects/" + n.Project.ID + "/global" + "/healthChecks/" + n.Name
}

func (s *MockService) buildGlobalHealthCheckName(projectName, name string) (*globalHealthCheckName, error) {
	project, err := s.Projects.GetProjectByID(projectName)
	if err != nil {
		return nil, err
	}

	return &globalHealthCheckName{
		Project: project,
		Name:    name,
	}, nil
}
