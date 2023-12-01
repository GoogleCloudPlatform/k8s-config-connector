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

type GlobalHealthCheckV1 struct {
	*MockService
	pb.UnimplementedHealthChecksServer
}

func (s *GlobalHealthCheckV1) Get(ctx context.Context, req *pb.GetHealthCheckRequest) (*pb.HealthCheck, error) {
	reqName := "projects/" + req.GetProject() + "/global" + "/healthChecks/" + req.GetHealthCheck()
	name, err := s.parseGlobalHealthCheckName(reqName)
	if err != nil {
		return nil, err
	}

	fqn := name.String()

	obj := &pb.HealthCheck{}
	if err := s.storage.Get(ctx, fqn, obj); err != nil {
		if apierrors.IsNotFound(err) {
			return nil, status.Errorf(codes.NotFound, "healthCheck %q not found", name)
		} else {
			return nil, status.Errorf(codes.Internal, "error reading healthCheck: %v", err)
		}
	}

	return obj, nil
}

func (s *GlobalHealthCheckV1) Insert(ctx context.Context, req *pb.InsertHealthCheckRequest) (*pb.Operation, error) {
	reqName := "projects/" + req.GetProject() + "/global" + "/healthChecks/" + req.GetHealthCheckResource().GetName()
	name, err := s.parseGlobalHealthCheckName(reqName)
	if err != nil {
		return nil, err
	}

	fqn := name.String()

	id := s.generateID()

	obj := proto.Clone(req.GetHealthCheckResource()).(*pb.HealthCheck)
	obj.SelfLink = PtrTo("https://compute.googleapis.com/compute/v1/" + name.String())
	obj.CreationTimestamp = PtrTo(s.nowString())
	obj.Id = &id
	obj.Kind = PtrTo("compute#healthCheck")

	if err := s.storage.Create(ctx, fqn, obj); err != nil {
		return nil, status.Errorf(codes.Internal, "error creating healthCheck: %v", err)
	}

	return s.newLRO(ctx, name.Project.ID)
}

// Updates a HealthCheck resource in the specified project using the data included in the request.
// This method supports PATCH semantics and uses the JSON merge patch format and processing rules.
func (s *GlobalHealthCheckV1) Patch(ctx context.Context, req *pb.PatchHealthCheckRequest) (*pb.Operation, error) {
	reqName := "projects/" + req.GetProject() + "/global" + "/healthChecks/" + req.GetHealthCheckResource().GetName()
	name, err := s.parseGlobalHealthCheckName(reqName)
	if err != nil {
		return nil, err
	}

	fqn := name.String()
	obj := &pb.HealthCheck{}
	if err := s.storage.Get(ctx, fqn, obj); err != nil {
		if apierrors.IsNotFound(err) {
			return nil, status.Errorf(codes.NotFound, "healthCheck %q not found", fqn)
		}
		return nil, status.Errorf(codes.Internal, "error reading healthCheck: %v", err)
	}

	// TODO: Implement helper to implement the full rules here
	proto.Merge(obj, req.GetHealthCheckResource())

	if err := s.storage.Update(ctx, fqn, obj); err != nil {
		return nil, status.Errorf(codes.Internal, "error updating healthCheck: %v", err)
	}

	return s.newLRO(ctx, name.Project.ID)
}

// Updates a HealthCheck resource in the specified project using the data included in the request.
func (s *GlobalHealthCheckV1) Update(ctx context.Context, req *pb.UpdateHealthCheckRequest) (*pb.Operation, error) {
	reqName := "projects/" + req.GetProject() + "/global" + "/healthChecks/" + req.GetHealthCheckResource().GetName()
	name, err := s.parseGlobalHealthCheckName(reqName)
	if err != nil {
		return nil, err
	}

	fqn := name.String()
	obj := &pb.HealthCheck{}
	if err := s.storage.Get(ctx, fqn, obj); err != nil {
		if apierrors.IsNotFound(err) {
			return nil, status.Errorf(codes.NotFound, "healthCheck %q not found", fqn)
		}
		return nil, status.Errorf(codes.Internal, "error reading healthCheck: %v", err)
	}

	// TODO: Implement helper to implement the full rules here
	proto.Merge(obj, req.GetHealthCheckResource())

	if err := s.storage.Update(ctx, fqn, obj); err != nil {
		return nil, status.Errorf(codes.Internal, "error updating healthCheck: %v", err)
	}

	return s.newLRO(ctx, name.Project.ID)
}

func (s *GlobalHealthCheckV1) Delete(ctx context.Context, req *pb.DeleteHealthCheckRequest) (*pb.Operation, error) {
	reqName := "projects/" + req.GetProject() + "/global" + "/healthChecks/" + req.GetHealthCheck()
	name, err := s.parseGlobalHealthCheckName(reqName)
	if err != nil {
		return nil, err
	}

	fqn := name.String()

	deleted := &pb.HealthCheck{}
	if err := s.storage.Delete(ctx, fqn, deleted); err != nil {
		if apierrors.IsNotFound(err) {
			return nil, status.Errorf(codes.NotFound, "healthCheck %q not found", name)
		} else {
			return nil, status.Errorf(codes.Internal, "error deleting healthCheck: %v", err)
		}
	}

	return s.newLRO(ctx, name.Project.ID)
}

type globalHealthCheckName struct {
	Project *projects.ProjectData
	Name    string
}

func (n *globalHealthCheckName) String() string {
	return "projects/" + n.Project.ID + "/global" + "/healthChecks/" + n.Name
}

// parseGlobalHealthCheckName parses a string into a globalHealthCheckName.
// The expected form is `projects/*/regions/*/healthcheck/*`.
func (s *MockService) parseGlobalHealthCheckName(name string) (*globalHealthCheckName, error) {
	tokens := strings.Split(name, "/")

	if len(tokens) == 5 && tokens[0] == "projects" && tokens[2] == "global" && tokens[3] == "healthChecks" {
		project, err := s.projects.GetProjectByID(tokens[1])
		if err != nil {
			return nil, err
		}

		name := &globalHealthCheckName{
			Project: project,
			Name:    tokens[4],
		}

		return name, nil
	} else {
		return nil, status.Errorf(codes.InvalidArgument, "name %q is not valid", name)
	}
}
