// Copyright 2025 Google LLC
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

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/proto"

	"github.com/GoogleCloudPlatform/k8s-config-connector/mockgcp/common/projects"
	pb "github.com/GoogleCloudPlatform/k8s-config-connector/mockgcp/generated/mockgcp/cloud/compute/v1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/mockgcp/pkg/storage"
)

type RoutesV1 struct {
	*MockService
	pb.UnimplementedRoutesServer
}

func (s *RoutesV1) Get(ctx context.Context, req *pb.GetRouteRequest) (*pb.Route, error) {
	name, err := s.newRouteName(req.GetProject(), req.GetRoute())
	if err != nil {
		return nil, err
	}

	fqn := name.String()

	obj := &pb.Route{}
	if err := s.storage.Get(ctx, fqn, obj); err != nil {
		if status.Code(err) == codes.NotFound {
			return nil, status.Errorf(codes.NotFound, "The resource '%s' was not found", fqn)
		}
		return nil, err
	}

	return obj, nil
}

func (s *RoutesV1) List(ctx context.Context, req *pb.ListRoutesRequest) (*pb.RouteList, error) {
	name, err := s.newRouteName(req.GetProject(), "placeholder")
	if err != nil {
		return nil, err
	}

	findPrefix := strings.TrimSuffix(name.String(), "placeholder")

	response := &pb.RouteList{}
	response.Id = PtrTo("0123456789")
	response.Kind = PtrTo("compute#routeList")
	response.SelfLink = PtrTo(buildComputeSelfLink(ctx, strings.TrimSuffix(findPrefix, "/")))

	findKind := (&pb.Route{}).ProtoReflect().Descriptor()
	if err := s.storage.List(ctx, findKind, storage.ListOptions{Prefix: findPrefix}, func(obj proto.Message) error {
		route := obj.(*pb.Route)
		isMatch, err := matchFilter(req.GetFilter(), route)
		if err != nil {
			return err
		}
		if isMatch {
			response.Items = append(response.Items, route)
		}
		return nil
	}); err != nil {
		return nil, err
	}

	return response, nil
}

func (s *RoutesV1) Insert(ctx context.Context, req *pb.InsertRouteRequest) (*pb.Operation, error) {
	name, err := s.newRouteName(req.GetProject(), req.GetRouteResource().GetName())
	if err != nil {
		return nil, err
	}

	fqn := name.String()

	id := s.generateID()

	obj := proto.Clone(req.GetRouteResource()).(*pb.Route)
	obj.SelfLink = PtrTo(buildComputeSelfLink(ctx, fqn))
	obj.CreationTimestamp = PtrTo(s.nowString())
	obj.Id = &id
	obj.Kind = PtrTo("compute#route")

	networkName, err := s.parseNetworkSelfLink(obj.GetNetwork())
	if err != nil {
		return nil, status.Errorf(codes.InvalidArgument, "network %q is not valid", obj.GetNetwork())
	}
	obj.Network = PtrTo(buildComputeSelfLink(ctx, fmt.Sprintf("projects/%s/global/networks/%s", networkName.Project.ID, networkName.Name)))

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

func (s *RoutesV1) Delete(ctx context.Context, req *pb.DeleteRouteRequest) (*pb.Operation, error) {
	name, err := s.newRouteName(req.GetProject(), req.GetRoute())
	if err != nil {
		return nil, err
	}

	fqn := name.String()

	existing := &pb.Route{}
	if err := s.storage.Get(ctx, fqn, existing); err != nil {
		if status.Code(err) == codes.NotFound {
			return nil, status.Errorf(codes.NotFound, "The resource '%s' was not found", fqn)
		}
		return nil, err
	}

	deleted := &pb.Route{}
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

type routeName struct {
	Project *projects.ProjectData
	Name    string
}

func (n *routeName) String() string {
	return "projects/" + n.Project.ID + "/global/routes/" + n.Name
}

// parseRouteName parses a string into a routeName.
// The expected form is `projects/*/global/routes/*`.
func (s *MockService) parseRouteName(name string) (*routeName, error) {
	tokens := strings.Split(name, "/")

	if len(tokens) == 5 && tokens[0] == "projects" && tokens[2] == "global" && tokens[3] == "routes" {
		return s.newRouteName(tokens[1], tokens[4])
	} else {
		return nil, status.Errorf(codes.InvalidArgument, "name %q is not valid", name)
	}
}

// newRouteName builds a normalized routeName from the constituent parts.
func (s *MockService) newRouteName(project string, name string) (*routeName, error) {
	projectObj, err := s.Projects.GetProjectByID(project)
	if err != nil {
		return nil, err
	}

	return &routeName{
		Project: projectObj,
		Name:    name,
	}, nil
}
