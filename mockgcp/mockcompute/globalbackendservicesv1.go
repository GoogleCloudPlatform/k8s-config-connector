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

// +tool:mockgcp-support
// proto.service: google.cloud.compute.v1.BackendServices
// proto.message: google.cloud.compute.v1.BackendService

package mockcompute

import (
	"context"
	"strings"

	"github.com/GoogleCloudPlatform/k8s-config-connector/mockgcp/common/projects"
	pb "github.com/GoogleCloudPlatform/k8s-config-connector/mockgcp/generated/mockgcp/cloud/compute/v1"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/types/known/emptypb"
)

type GlobalBackendServicesV1 struct {
	*MockService
	pb.UnimplementedBackendServicesServer
}

func (s *GlobalBackendServicesV1) Get(ctx context.Context, req *pb.GetBackendServiceRequest) (*pb.BackendService, error) {
	reqName := "projects/" + req.GetProject() + "/global" + "/backendServices/" + req.GetBackendService()
	name, err := s.parseGlobalBackendServiceName(reqName)
	if err != nil {
		return nil, err
	}

	fqn := name.String()

	obj := &pb.BackendService{}
	if err := s.storage.Get(ctx, fqn, obj); err != nil {
		return nil, status.Errorf(codes.NotFound, "The resource '%s' was not found", fqn)
	}

	return obj, nil
}

func (s *GlobalBackendServicesV1) Insert(ctx context.Context, req *pb.InsertBackendServiceRequest) (*pb.Operation, error) {
	reqName := "projects/" + req.GetProject() + "/global" + "/backendServices/" + req.GetBackendServiceResource().GetName()
	name, err := s.parseGlobalBackendServiceName(reqName)
	if err != nil {
		return nil, err
	}

	fqn := name.String()

	id := s.generateID()

	obj := proto.Clone(req.GetBackendServiceResource()).(*pb.BackendService)
	obj.SelfLink = PtrTo(buildComputeSelfLink(ctx, fqn))
	obj.CreationTimestamp = PtrTo(s.nowString())
	obj.Id = &id
	obj.Kind = PtrTo("compute#backendService")

	s.populateDefaultsForBackendService(obj)

	if err := s.storage.Create(ctx, fqn, obj); err != nil {
		return nil, err
	}

	op := &pb.Operation{
		OperationType: PtrTo("insert"),
		TargetId:      obj.Id,
		TargetLink:    obj.SelfLink,
		User:          PtrTo("user@example.com"),
	}
	return s.computeOperations.startGlobalLRO(ctx, name.Project.ID, op, func() (proto.Message, error) {
		return obj, nil
	})
}

func (s *GlobalBackendServicesV1) populateDefaultsForBackendService(obj *pb.BackendService) {
	if obj.AffinityCookieTtlSec == nil {
		obj.AffinityCookieTtlSec = PtrTo(int32(0))
	}
	if obj.ConnectionDraining == nil {
		obj.ConnectionDraining = &pb.ConnectionDraining{}
	}
	if obj.ConnectionDraining.DrainingTimeoutSec == nil {
		obj.ConnectionDraining.DrainingTimeoutSec = PtrTo(int32(0))
	}
	if obj.Description == nil {
		obj.Description = PtrTo("")
	}
	if obj.EnableCDN == nil {
		obj.EnableCDN = PtrTo(false)
	}
	if obj.LoadBalancingScheme == nil {
		obj.LoadBalancingScheme = PtrTo("EXTERNAL")
	}
	if obj.Port == nil {
		obj.Port = PtrTo(int32(80))
	}
	if obj.SessionAffinity == nil {
		obj.SessionAffinity = PtrTo("NONE")

	}

	obj.Fingerprint = PtrTo(computeFingerprint(obj))
}

func (s *GlobalBackendServicesV1) Update(ctx context.Context, req *pb.UpdateBackendServiceRequest) (*pb.Operation, error) {
	reqName := "projects/" + req.GetProject() + "/global" + "/backendServices/" + req.GetBackendServiceResource().GetName()
	name, err := s.parseGlobalBackendServiceName(reqName)
	if err != nil {
		return nil, err
	}

	fqn := name.String()
	obj := &pb.BackendService{}
	if err := s.storage.Get(ctx, fqn, obj); err != nil {
		return nil, err
	}

	// TODO: Implement helper to implement the full rules here
	proto.Merge(obj, req.GetBackendServiceResource())

	if err := s.storage.Update(ctx, fqn, obj); err != nil {
		return nil, err
	}

	op := &pb.Operation{
		OperationType: PtrTo("insert"),
		TargetId:      obj.Id,
		TargetLink:    obj.SelfLink,
		User:          PtrTo("user@example.com"),
	}
	return s.computeOperations.startGlobalLRO(ctx, name.Project.ID, op, func() (proto.Message, error) {
		return obj, nil
	})
}

func (s *GlobalBackendServicesV1) Patch(ctx context.Context, req *pb.PatchBackendServiceRequest) (*pb.Operation, error) {
	reqName := "projects/" + req.GetProject() + "/global" + "/backendServices/" + req.GetBackendServiceResource().GetName()
	name, err := s.parseGlobalBackendServiceName(reqName)
	if err != nil {
		return nil, err
	}

	fqn := name.String()
	obj := &pb.BackendService{}
	if err := s.storage.Get(ctx, fqn, obj); err != nil {
		return nil, err
	}

	// TODO: Implement helper to implement the full rules here
	proto.Merge(obj, req.GetBackendServiceResource())

	if err := s.storage.Update(ctx, fqn, obj); err != nil {
		return nil, err
	}

	op := &pb.Operation{
		OperationType: PtrTo("patch"),
		TargetId:      obj.Id,
		TargetLink:    obj.SelfLink,
		User:          PtrTo("user@example.com"),
	}
	return s.computeOperations.startGlobalLRO(ctx, name.Project.ID, op, func() (proto.Message, error) {
		return obj, nil
	})
}

func (s *GlobalBackendServicesV1) Delete(ctx context.Context, req *pb.DeleteBackendServiceRequest) (*pb.Operation, error) {
	reqName := "projects/" + req.GetProject() + "/global" + "/backendServices/" + req.GetBackendService()
	name, err := s.parseGlobalBackendServiceName(reqName)
	if err != nil {
		return nil, err
	}

	fqn := name.String()

	deleted := &pb.BackendService{}
	if err := s.storage.Delete(ctx, fqn, deleted); err != nil {
		return nil, err
	}

	op := &pb.Operation{
		OperationType: PtrTo("delete"),
		TargetId:      deleted.Id,
		TargetLink:    deleted.SelfLink,
		User:          PtrTo("user@example.com"),
	}
	return s.computeOperations.startGlobalLRO(ctx, name.Project.ID, op, func() (proto.Message, error) {
		return &emptypb.Empty{}, nil
	})
}

type globalBackendServiceName struct {
	Project *projects.ProjectData
	Name    string
}

func (n *globalBackendServiceName) String() string {
	return "projects/" + n.Project.ID + "/global" + "/backendServices/" + n.Name
}

// parseBackendServiceName parses a string into a backendserviceName.
// The expected form is `projects/*/regions/*/backendservice/*`.
func (s *MockService) parseGlobalBackendServiceName(name string) (*globalBackendServiceName, error) {
	tokens := strings.Split(name, "/")

	if len(tokens) == 5 && tokens[0] == "projects" && tokens[2] == "global" && tokens[3] == "backendServices" {
		project, err := s.Projects.GetProjectByID(tokens[1])
		if err != nil {
			return nil, err
		}

		name := &globalBackendServiceName{
			Project: project,
			Name:    tokens[4],
		}

		return name, nil
	} else {
		return nil, status.Errorf(codes.InvalidArgument, "name %q is not valid", name)
	}
}
