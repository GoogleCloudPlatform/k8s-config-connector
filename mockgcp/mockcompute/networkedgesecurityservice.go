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

// +tool:mockgcp-support
// proto.service: google.cloud.compute.v1.NetworkEdgeSecurityServices
// proto.message: google.cloud.compute.v1.NetworkEdgeSecurityService

package mockcompute

import (
	"context"
	"fmt"
	"strings"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/types/known/emptypb"

	"github.com/GoogleCloudPlatform/k8s-config-connector/mockgcp/common/projects"
	pb "github.com/GoogleCloudPlatform/k8s-config-connector/mockgcp/generated/mockgcp/cloud/compute/v1"
)

type networkEdgeSecurityServices struct {
	*MockService
	pb.UnimplementedNetworkEdgeSecurityServicesServer
}

func (s *networkEdgeSecurityServices) Get(ctx context.Context, req *pb.GetNetworkEdgeSecurityServiceRequest) (*pb.NetworkEdgeSecurityService, error) {
	reqName := fmt.Sprintf("projects/%s/regions/%s/networkEdgeSecurityServices/%s", req.GetProject(), req.GetRegion(), req.GetNetworkEdgeSecurityService())
	name, err := s.parseNetworkEdgeSecurityServiceName(reqName)
	if err != nil {
		return nil, err
	}
	fqn := name.String()

	obj := &pb.NetworkEdgeSecurityService{}
	if err := s.storage.Get(ctx, fqn, obj); err != nil {
		if status.Code(err) == codes.NotFound {
			return nil, status.Errorf(codes.NotFound, "NetworkEdgeSecurityService %q not found", name)
		}
		return nil, err
	}

	return obj, nil
}

func (s *networkEdgeSecurityServices) Insert(ctx context.Context, req *pb.InsertNetworkEdgeSecurityServiceRequest) (*pb.Operation, error) {
	reqName := fmt.Sprintf("projects/%s/regions/%s/networkEdgeSecurityServices/%s", req.GetProject(), req.GetRegion(), req.GetNetworkEdgeSecurityServiceResource().GetName())
	name, err := s.parseNetworkEdgeSecurityServiceName(reqName)
	if err != nil {
		return nil, err
	}
	fqn := name.String()

	obj := proto.Clone(req.GetNetworkEdgeSecurityServiceResource()).(*pb.NetworkEdgeSecurityService)
	obj.Id = proto.Uint64(s.generateID())
	obj.SelfLink = PtrTo(buildComputeSelfLink(ctx, fqn))
	obj.Kind = PtrTo("compute#networkEdgeSecurityService")
	obj.CreationTimestamp = PtrTo(s.nowString())
	obj.Region = PtrTo(fmt.Sprintf("https://www.googleapis.com/compute/v1/projects/%s/regions/%s", name.Project.ID, name.Region))
	obj.Fingerprint = PtrTo(computeCRC32C(obj))

	if err := s.storage.Create(ctx, fqn, obj); err != nil {
		return nil, err
	}

	op := &pb.Operation{
		OperationType: PtrTo("insert"),
		TargetId:      obj.Id,
		TargetLink:    obj.SelfLink,
		User:          PtrTo("user@example.com"),
	}
	return s.computeOperations.startRegionalLRO(ctx, name.Project.ID, name.Region, op, func() (proto.Message, error) {
		return obj, nil
	})
}

func (s *networkEdgeSecurityServices) Patch(ctx context.Context, req *pb.PatchNetworkEdgeSecurityServiceRequest) (*pb.Operation, error) {
	reqName := fmt.Sprintf("projects/%s/regions/%s/networkEdgeSecurityServices/%s", req.GetProject(), req.GetRegion(), req.GetNetworkEdgeSecurityService())
	name, err := s.parseNetworkEdgeSecurityServiceName(reqName)
	if err != nil {
		return nil, err
	}
	fqn := name.String()

	obj := &pb.NetworkEdgeSecurityService{}
	if err := s.storage.Get(ctx, fqn, obj); err != nil {
		return nil, err
	}

	// TODO: Apply field mask.
	proto.Merge(obj, req.GetNetworkEdgeSecurityServiceResource())
	obj.Fingerprint = PtrTo(computeCRC32C(obj))

	if err := s.storage.Update(ctx, fqn, obj); err != nil {
		return nil, err
	}

	op := &pb.Operation{
		OperationType: PtrTo("patch"),
		TargetId:      obj.Id,
		TargetLink:    obj.SelfLink,
		User:          PtrTo("user@example.com"),
	}
	return s.computeOperations.startRegionalLRO(ctx, name.Project.ID, name.Region, op, func() (proto.Message, error) {
		return obj, nil
	})
}

func (s *networkEdgeSecurityServices) Delete(ctx context.Context, req *pb.DeleteNetworkEdgeSecurityServiceRequest) (*pb.Operation, error) {
	reqName := fmt.Sprintf("projects/%s/regions/%s/networkEdgeSecurityServices/%s", req.GetProject(), req.GetRegion(), req.GetNetworkEdgeSecurityService())
	name, err := s.parseNetworkEdgeSecurityServiceName(reqName)
	if err != nil {
		return nil, err
	}
	fqn := name.String()

	deleted := &pb.NetworkEdgeSecurityService{}
	if err := s.storage.Delete(ctx, fqn, deleted); err != nil {
		return nil, err
	}

	op := &pb.Operation{
		OperationType: PtrTo("delete"),
		TargetId:      deleted.Id,
		TargetLink:    deleted.SelfLink,
		User:          PtrTo("user@example.com"),
	}
	return s.computeOperations.startRegionalLRO(ctx, name.Project.ID, name.Region, op, func() (proto.Message, error) {
		return &emptypb.Empty{}, nil
	})
}

type networkEdgeSecurityServiceName struct {
	Project *projects.ProjectData
	Region  string
	Name    string
}

func (n *networkEdgeSecurityServiceName) String() string {
	return fmt.Sprintf("projects/%s/regions/%s/networkEdgeSecurityServices/%s", n.Project.ID, n.Region, n.Name)
}

// parseNetworkEdgeSecurityServiceName parses a string into a networkEdgeSecurityServiceName.
// The expected form is `projects/*/regions/*/networkEdgeSecurityServices/*`.
func (s *MockService) parseNetworkEdgeSecurityServiceName(name string) (*networkEdgeSecurityServiceName, error) {
	tokens := strings.Split(name, "/")

	if len(tokens) == 6 && tokens[0] == "projects" && tokens[2] == "regions" && tokens[4] == "networkEdgeSecurityServices" {
		project, err := s.Projects.GetProjectByID(tokens[1])
		if err != nil {
			return nil, err
		}

		name := &networkEdgeSecurityServiceName{
			Project: project,
			Region:  tokens[3],
			Name:    tokens[5],
		}

		return name, nil
	} else {
		return nil, status.Errorf(codes.InvalidArgument, "name %q is not valid", name)
	}
}
