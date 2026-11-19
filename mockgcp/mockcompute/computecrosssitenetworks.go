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

// +tool:mockgcp-support
// proto.service: google.cloud.compute.v1.CrossSiteNetworks
// proto.message: google.cloud.compute.v1.CrossSiteNetwork

package mockcompute

import (
	"context"
	"fmt"
	"strconv"
	"strings"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/types/known/emptypb"

	"github.com/GoogleCloudPlatform/k8s-config-connector/mockgcp/common/projects"
	pb "github.com/GoogleCloudPlatform/k8s-config-connector/mockgcp/generated/mockgcp/cloud/compute/v1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/mockgcp/pkg/storage"
)

type crossSiteNetworks struct {
	*MockService
	pb.UnimplementedCrossSiteNetworksServer
}

func (s *crossSiteNetworks) Get(ctx context.Context, req *pb.GetCrossSiteNetworkRequest) (*pb.CrossSiteNetwork, error) {
	reqName := fmt.Sprintf("projects/%s/global/crossSiteNetworks/%s", req.GetProject(), req.GetCrossSiteNetwork())
	name, err := s.parseCrossSiteNetworkName(reqName)
	if err != nil {
		return nil, err
	}
	fqn := name.String()

	obj := &pb.CrossSiteNetwork{}
	if err := s.storage.Get(ctx, fqn, obj); err != nil {
		if status.Code(err) == codes.NotFound {
			return nil, status.Errorf(codes.NotFound, "The resource '%s' was not found", fqn)
		}
		return nil, err
	}

	return obj, nil
}

func (s *crossSiteNetworks) Insert(ctx context.Context, req *pb.InsertCrossSiteNetworkRequest) (*pb.Operation, error) {
	reqName := fmt.Sprintf("projects/%s/global/crossSiteNetworks/%s", req.GetProject(), req.GetCrossSiteNetworkResource().GetName())
	name, err := s.parseCrossSiteNetworkName(reqName)
	if err != nil {
		return nil, err
	}
	fqn := name.String()

	obj := proto.CloneOf(req.GetCrossSiteNetworkResource())
	obj.Id = proto.Uint64(s.generateID())
	obj.SelfLink = PtrTo(BuildComputeSelfLink(ctx, fqn))
	obj.Kind = PtrTo("compute#crossSiteNetwork")
	obj.CreationTimestamp = PtrTo(s.nowString())

	if obj.Description == nil {
		obj.Description = PtrTo("")
	}

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

func (s *crossSiteNetworks) Delete(ctx context.Context, req *pb.DeleteCrossSiteNetworkRequest) (*pb.Operation, error) {
	reqName := fmt.Sprintf("projects/%s/global/crossSiteNetworks/%s", req.GetProject(), req.GetCrossSiteNetwork())
	name, err := s.parseCrossSiteNetworkName(reqName)
	if err != nil {
		return nil, err
	}
	fqn := name.String()

	deleted := &pb.CrossSiteNetwork{}
	if err := s.storage.Delete(ctx, fqn, deleted); err != nil {
		if status.Code(err) == codes.NotFound {
			return nil, status.Errorf(codes.NotFound, "The resource '%s' was not found", fqn)
		}
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

func (s *crossSiteNetworks) Patch(ctx context.Context, req *pb.PatchCrossSiteNetworkRequest) (*pb.Operation, error) {
	reqName := fmt.Sprintf("projects/%s/global/crossSiteNetworks/%s", req.GetProject(), req.GetCrossSiteNetwork())
	name, err := s.parseCrossSiteNetworkName(reqName)
	if err != nil {
		return nil, err
	}
	fqn := name.String()

	obj := &pb.CrossSiteNetwork{}
	if err := s.storage.Get(ctx, fqn, obj); err != nil {
		if status.Code(err) == codes.NotFound {
			return nil, status.Errorf(codes.NotFound, "The resource '%s' was not found", fqn)
		}
		return nil, err
	}

	if req.GetCrossSiteNetworkResource().Description != nil {
		obj.Description = req.GetCrossSiteNetworkResource().Description
	}

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

func (s *crossSiteNetworks) List(ctx context.Context, req *pb.ListCrossSiteNetworksRequest) (*pb.CrossSiteNetworkList, error) {
	project, err := s.Projects.GetProjectByID(req.GetProject())
	if err != nil {
		return nil, err
	}

	findPrefix := fmt.Sprintf("projects/%s/global/crossSiteNetworks/", project.ID)
	response := &pb.CrossSiteNetworkList{
		Id:       PtrTo(strconv.FormatUint(s.generateID(), 10)),
		Kind:     PtrTo("compute#crossSiteNetworkList"),
		SelfLink: PtrTo(BuildComputeSelfLink(ctx, fmt.Sprintf("projects/%s/global/crossSiteNetworks", project.ID))),
	}

	findKind := (&pb.CrossSiteNetwork{}).ProtoReflect().Descriptor()
	if err := s.storage.List(ctx, findKind, storage.ListOptions{Prefix: findPrefix}, func(obj proto.Message) error {
		network := obj.(*pb.CrossSiteNetwork)
		response.Items = append(response.Items, network)
		return nil
	}); err != nil {
		return nil, err
	}

	return response, nil
}

type crossSiteNetworkName struct {
	Project *projects.ProjectData
	Name    string
}

func (n *crossSiteNetworkName) String() string {
	return "projects/" + n.Project.ID + "/global/crossSiteNetworks/" + n.Name
}

func (s *MockService) parseCrossSiteNetworkName(name string) (*crossSiteNetworkName, error) {
	tokens := strings.Split(name, "/")

	if len(tokens) == 5 && tokens[0] == "projects" && tokens[2] == "global" && tokens[3] == "crossSiteNetworks" {
		project, err := s.Projects.GetProjectByID(tokens[1])
		if err != nil {
			return nil, err
		}

		name := &crossSiteNetworkName{
			Project: project,
			Name:    tokens[4],
		}

		return name, nil
	} else {
		return nil, status.Errorf(codes.InvalidArgument, "name %q is not valid", name)
	}
}
