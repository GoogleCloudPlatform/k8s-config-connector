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
)

type routersV1 struct {
	*MockService
	pb.UnimplementedRoutersServer
}

func (s *routersV1) Get(ctx context.Context, req *pb.GetRouterRequest) (*pb.Router, error) {
	name, err := s.newRouterName(req.GetProject(), req.GetRegion(), req.GetRouter())
	if err != nil {
		return nil, err
	}

	fqn := name.String()

	obj := &pb.Router{}
	if err := s.storage.Get(ctx, fqn, obj); err != nil {
		if status.Code(err) == codes.NotFound {
			return nil, status.Errorf(codes.NotFound, "The resource '%s' was not found", fqn)
		}
		return nil, err
	}

	return obj, nil
}

func (s *routersV1) Insert(ctx context.Context, req *pb.InsertRouterRequest) (*pb.Operation, error) {
	name, err := s.newRouterName(req.GetProject(), req.GetRegion(), req.GetRouterResource().GetName())
	if err != nil {
		return nil, err
	}

	fqn := name.String()

	id := s.generateID()

	obj := proto.Clone(req.GetRouterResource()).(*pb.Router)
	obj.SelfLink = PtrTo(buildComputeSelfLink(ctx, fqn))
	obj.CreationTimestamp = PtrTo(s.nowString())
	obj.Id = &id
	obj.Kind = PtrTo("compute#router")
	obj.Region = PtrTo(buildComputeSelfLink(ctx, fmt.Sprintf("projects/%s/regions/%s", name.Project.ID, name.Region)))

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

func (s *routersV1) Patch(ctx context.Context, req *pb.PatchRouterRequest) (*pb.Operation, error) {
	name, err := s.newRouterName(req.GetProject(), req.GetRegion(), req.GetRouter())
	if err != nil {
		return nil, err
	}

	fqn := name.String()
	obj := &pb.Router{}
	if err := s.storage.Get(ctx, fqn, obj); err != nil {
		return nil, err
	}

	// Basic patch logic: only update nats for now
	if req.GetRouterResource().Nats != nil {
		obj.Nats = req.GetRouterResource().Nats
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

func (s *routersV1) Delete(ctx context.Context, req *pb.DeleteRouterRequest) (*pb.Operation, error) {
	name, err := s.newRouterName(req.GetProject(), req.GetRegion(), req.GetRouter())
	if err != nil {
		return nil, err
	}

	fqn := name.String()

	existing := &pb.Router{}
	if err := s.storage.Get(ctx, fqn, existing); err != nil {
		if status.Code(err) == codes.NotFound {
			return nil, status.Errorf(codes.NotFound, "The resource '%s' was not found", fqn)
		}
		return nil, err
	}

	deleted := &pb.Router{}
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

type routerName struct {
	Project *projects.ProjectData
	Region  string
	Name    string
}

func (n *routerName) String() string {
	return "projects/" + n.Project.ID + "/regions/" + n.Region + "/routers/" + n.Name
}

func (s *MockService) newRouterName(project string, region string, name string) (*routerName, error) {
	projectObj, err := s.Projects.GetProjectByID(project)
	if err != nil {
		return nil, err
	}

	return &routerName{
		Project: projectObj,
		Region:  region,
		Name:    name,
	}, nil
}
