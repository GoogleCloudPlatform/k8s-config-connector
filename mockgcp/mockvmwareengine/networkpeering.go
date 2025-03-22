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

// +tool:mockgcp-support
// proto.service: google.cloud.vmwareengine.v1.VmwareEngine
// proto.message: google.cloud.vmwareengine.v1.NetworkPeering

package mockvmwareengine

import (
	"context"
	"fmt"
	"strings"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/types/known/emptypb"
	"google.golang.org/protobuf/types/known/timestamppb"

	"github.com/GoogleCloudPlatform/k8s-config-connector/mockgcp/common/projects"
	pb "github.com/GoogleCloudPlatform/k8s-config-connector/mockgcp/generated/mockgcp/cloud/vmwareengine/v1"
	longrunningpb "google.golang.org/genproto/googleapis/longrunning"
)

func (s *VMwareEngineV1) GetNetworkPeering(ctx context.Context, req *pb.GetNetworkPeeringRequest) (*pb.NetworkPeering, error) {
	name, err := s.parseNetworkPeeringName(req.Name)
	if err != nil {
		return nil, err
	}

	fqn := name.String()

	obj := &pb.NetworkPeering{}
	if err := s.storage.Get(ctx, fqn, obj); err != nil {
		if status.Code(err) == codes.NotFound {
			return nil, status.Errorf(codes.NotFound, "Resource '%s' was not found", fqn)
		}
		return nil, err
	}

	return obj, nil
}

func (s *VMwareEngineV1) CreateNetworkPeering(ctx context.Context, req *pb.CreateNetworkPeeringRequest) (*longrunningpb.Operation, error) {
	reqName := req.Parent + "/networkPeerings/" + req.NetworkPeeringId
	name, err := s.parseNetworkPeeringName(reqName)
	if err != nil {
		return nil, err
	}

	fqn := name.String()

	obj := proto.Clone(req.GetNetworkPeering()).(*pb.NetworkPeering)
	obj.Name = fqn
	obj.CreateTime = timestamppb.Now()
	obj.State = pb.NetworkPeering_ACTIVE
	s.setNetworkPeeringFields(obj)

	if err := s.storage.Create(ctx, fqn, obj); err != nil {
		return nil, err
	}

	prefix := fmt.Sprintf("projects/%s/locations/%s", name.Project.ID, name.Location)
	metadata := &pb.OperationMetadata{
		CreateTime: timestamppb.Now(),
		Target:     fqn,
		ApiVersion: "v1",
		Verb:       "create",
	}

	return s.operations.StartLRO(ctx, prefix, metadata, func() (proto.Message, error) {
		metadata.EndTime = timestamppb.Now()
		obj.UpdateTime = timestamppb.Now()
		if err := s.storage.Update(ctx, fqn, obj); err != nil {
			return nil, err
		}
		return obj, nil
	})
}

func (s *VMwareEngineV1) UpdateNetworkPeering(ctx context.Context, req *pb.UpdateNetworkPeeringRequest) (*longrunningpb.Operation, error) {
	name, err := s.parseNetworkPeeringName(req.GetNetworkPeering().GetName())
	if err != nil {
		return nil, err
	}

	fqn := name.String()

	obj := &pb.NetworkPeering{}
	if err := s.storage.Get(ctx, fqn, obj); err != nil {
		return nil, err
	}

	paths := req.GetUpdateMask().GetPaths()
	if len(paths) == 0 {
		return nil, status.Errorf(codes.InvalidArgument, "update_mask must be provided")
	}

	for _, path := range paths {
		switch path {
		case "description":
			obj.Description = req.GetNetworkPeering().Description
		// TODO: add support for other fields
		default:
			return nil, status.Errorf(codes.InvalidArgument, "update_mask path %q not valid", path)
		}
	}

	obj.UpdateTime = timestamppb.Now()

	if err := s.storage.Update(ctx, fqn, obj); err != nil {
		return nil, err
	}

	prefix := fmt.Sprintf("projects/%s/locations/%s", name.Project.ID, name.Location)
	metadata := &pb.OperationMetadata{
		ApiVersion: "v1",
		CreateTime: timestamppb.Now(),
		Target:     name.String(),
		Verb:       "update",
	}

	return s.operations.StartLRO(ctx, prefix, metadata, func() (proto.Message, error) {
		metadata.EndTime = timestamppb.Now()
		return obj, nil
	})
}

func (s *VMwareEngineV1) DeleteNetworkPeering(ctx context.Context, req *pb.DeleteNetworkPeeringRequest) (*longrunningpb.Operation, error) {
	name, err := s.parseNetworkPeeringName(req.Name)
	if err != nil {
		return nil, err
	}

	fqn := name.String()

	deleted := &pb.NetworkPeering{}
	if err := s.storage.Delete(ctx, fqn, deleted); err != nil {
		return nil, err
	}

	prefix := fmt.Sprintf("projects/%s/locations/%s", name.Project.ID, name.Location)
	metadata := &pb.OperationMetadata{
		CreateTime: timestamppb.Now(),
		Target:     fqn,
		ApiVersion: "v1",
		Verb:       "delete",
	}
	return s.operations.StartLRO(ctx, prefix, metadata, func() (proto.Message, error) {
		metadata.EndTime = timestamppb.Now()
		return &emptypb.Empty{}, nil
	})
}

type networkPeeringParent struct {
	Project  *projects.ProjectData
	Location string
}

func (n *networkPeeringParent) String() string {
	return fmt.Sprintf("projects/%s/locations/%s", n.Project.ID, n.Location)
}

type networkPeeringName struct {
	Project          *projects.ProjectData
	Location         string
	NetworkPeeringID string
}

func (n *networkPeeringName) String() string {
	return fmt.Sprintf("projects/%s/locations/%s/networkPeerings/%s", n.Project.ID, n.Location, n.NetworkPeeringID)
}

func (s *MockService) parseNetworkPeeringName(name string) (*networkPeeringName, error) {
	tokens := strings.Split(name, "/")

	if len(tokens) == 6 && tokens[0] == "projects" && tokens[2] == "locations" && tokens[4] == "networkPeerings" {
		project, err := s.Projects.GetProjectByID(tokens[1])
		if err != nil {
			return nil, err
		}

		name := &networkPeeringName{
			Project:          project,
			Location:         tokens[3],
			NetworkPeeringID: tokens[5],
		}

		return name, nil
	}

	return nil, status.Errorf(codes.InvalidArgument, "name %q is not valid", name)
}

func (s *MockService) setNetworkPeeringFields(obj *pb.NetworkPeering) {
	if obj.ExchangeSubnetRoutes == nil {
		obj.ExchangeSubnetRoutes = proto.Bool(true)
	}
	if obj.ExportCustomRoutes == nil {
		obj.ExportCustomRoutes = proto.Bool(true)
	}
	if obj.ExportCustomRoutesWithPublicIp == nil {
		obj.ExportCustomRoutesWithPublicIp = proto.Bool(true)
	}
	if obj.ImportCustomRoutes == nil {
		obj.ImportCustomRoutes = proto.Bool(true)
	}
	if obj.ImportCustomRoutesWithPublicIp == nil {
		obj.ImportCustomRoutesWithPublicIp = proto.Bool(true)
	}
	if obj.PeerMtu == 0 {
		obj.PeerMtu = 1500
	}
	obj.StateDetails = "[2025-01-02T03:04:05.678-09:00]: Connected."
	obj.Uid = "111111111111111111111"
}
