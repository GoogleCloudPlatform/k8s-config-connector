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
	"strings"

	"github.com/GoogleCloudPlatform/k8s-config-connector/mockgcp/common/projects"
	pb "github.com/GoogleCloudPlatform/k8s-config-connector/mockgcp/generated/mockgcp/cloud/compute/v1"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/proto"
)

type RoutersV1 struct {
	*MockService
	pb.UnimplementedRoutersServer
}

func (s *RoutersV1) Get(ctx context.Context, req *pb.GetRouterRequest) (*pb.Router, error) {
	reqName := "projects/" + req.GetProject() + "/regions/" + req.GetRegion() + "/routers/" + req.GetRouter()
	name, err := s.parseRouterName(reqName)
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

func (s *RoutersV1) Insert(ctx context.Context, req *pb.InsertRouterRequest) (*pb.Operation, error) {
	reqName := "projects/" + req.GetProject() + "/regions/" + req.GetRegion() + "/routers/" + req.GetRouterResource().GetName()
	name, err := s.parseRouterName(reqName)
	if err != nil {
		return nil, err
	}

	fqn := name.String()

	id := s.generateID()

	obj := proto.Clone(req.GetRouterResource()).(*pb.Router)
	obj.SelfLink = PtrTo(BuildComputeSelfLink(ctx, fqn))
	obj.CreationTimestamp = PtrTo(s.nowString())
	obj.Id = &id
	obj.Kind = PtrTo("compute#router")

	if obj.Description == nil {
		obj.Description = PtrTo("")
	}

	if obj.EncryptedInterconnectRouter == nil {
		obj.EncryptedInterconnectRouter = PtrTo(false)
	}

	if obj.Network != nil {
		networkName, err := s.parseNetworkSelfLink(obj.GetNetwork())
		if err != nil {
			return nil, status.Errorf(codes.InvalidArgument, "network %q is not valid", obj.GetNetwork())
		}
		obj.Network = PtrTo(BuildComputeSelfLink(ctx, fmt.Sprintf("projects/%s/global/networks/%s", networkName.Project.ID, networkName.Name)))
	}

	// output only fields
	obj.Region = PtrTo(BuildComputeSelfLink(ctx, fmt.Sprintf("projects/%s/regions/%s", name.Project.ID, name.Region)))

	s.populateRouter(obj)

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

func (s *RoutersV1) Patch(ctx context.Context, req *pb.PatchRouterRequest) (*pb.Operation, error) {
	reqName := "projects/" + req.GetProject() + "/regions/" + req.GetRegion() + "/routers/" + req.GetRouter()
	name, err := s.parseRouterName(reqName)
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

	hasInterfaces := req.GetRouterResource().Interfaces != nil
	hasBgpPeers := req.GetRouterResource().BgpPeers != nil

	proto.Merge(obj, req.GetRouterResource())

	if hasInterfaces {
		obj.Interfaces = req.GetRouterResource().Interfaces
	}
	if hasBgpPeers {
		obj.BgpPeers = req.GetRouterResource().BgpPeers
	}

	s.populateRouter(obj)

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

func (s *RoutersV1) Delete(ctx context.Context, req *pb.DeleteRouterRequest) (*pb.Operation, error) {
	reqName := "projects/" + req.GetProject() + "/regions/" + req.GetRegion() + "/routers/" + req.GetRouter()
	name, err := s.parseRouterName(reqName)
	if err != nil {
		return nil, err
	}

	fqn := name.String()

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

func (s *MockService) parseRouterName(name string) (*routerName, error) {
	tokens := strings.Split(name, "/")

	if len(tokens) == 6 && tokens[0] == "projects" && tokens[2] == "regions" && tokens[4] == "routers" {
		project, err := s.Projects.GetProjectByID(tokens[1])
		if err != nil {
			return nil, err
		}

		name := &routerName{
			Project: project,
			Region:  tokens[3],
			Name:    tokens[5],
		}

		return name, nil
	} else {
		return nil, status.Errorf(codes.InvalidArgument, "name %q is not valid", name)
	}
}

func (s *RoutersV1) populateRouter(obj *pb.Router) {
	for _, iface := range obj.Interfaces {
		if iface.IpVersion == nil {
			iface.IpVersion = PtrTo("IPV4")
		}
	}
}
