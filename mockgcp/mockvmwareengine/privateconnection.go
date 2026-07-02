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
// proto.service: google.cloud.vmwareengine.v1.VmwareEngine
// proto.message: google.cloud.vmwareengine.v1.PrivateConnection

package mockvmwareengine

import (
	"context"
	"fmt"
	"strings"
	"time"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/types/known/emptypb"
	"google.golang.org/protobuf/types/known/timestamppb"

	pb "cloud.google.com/go/vmwareengine/apiv1/vmwareenginepb"
	"github.com/GoogleCloudPlatform/k8s-config-connector/mockgcp/common/projects"
	"github.com/google/uuid"
	longrunningpb "google.golang.org/genproto/googleapis/longrunning"
)

func (s *VMwareEngineV1) GetPrivateConnection(ctx context.Context, req *pb.GetPrivateConnectionRequest) (*pb.PrivateConnection, error) {
	name, err := s.parsePrivateConnectionName(req.Name)
	if err != nil {
		return nil, err
	}

	fqn := name.String()

	obj := &pb.PrivateConnection{}
	if err := s.storage.Get(ctx, fqn, obj); err != nil {
		if status.Code(err) == codes.NotFound {
			return nil, status.Errorf(codes.NotFound, "Resource '%s' was not found", fqn)
		}
		return nil, err
	}

	return obj, nil
}

func (s *VMwareEngineV1) CreatePrivateConnection(ctx context.Context, req *pb.CreatePrivateConnectionRequest) (*longrunningpb.Operation, error) {
	reqName := req.Parent + "/privateConnections/" + req.PrivateConnectionId
	name, err := s.parsePrivateConnectionName(reqName)
	if err != nil {
		return nil, err
	}

	fqn := name.String()

	now := time.Now()

	obj := proto.Clone(req.GetPrivateConnection()).(*pb.PrivateConnection)
	obj.Name = fqn
	obj.CreateTime = timestamppb.New(now)
	obj.UpdateTime = timestamppb.New(now)
	obj.State = pb.PrivateConnection_ACTIVE
	obj.Uid = uuid.New().String()
	obj.PeeringId = "peering-55c50401-d392-4c50-94af-d4a1242f70ac"
	obj.PeeringState = pb.PrivateConnection_PEERING_ACTIVE

	// VmwareEngineNetworkCanonical
	parts := strings.Split(obj.VmwareEngineNetwork, "/")
	if len(parts) >= 6 && parts[0] == "projects" && parts[2] == "locations" && parts[4] == "vmwareEngineNetworks" {
		obj.VmwareEngineNetworkCanonical = fmt.Sprintf("projects/%d/locations/%s/vmwareEngineNetworks/%s",
			name.Project.Number, parts[3], parts[5])
	}

	if err := s.storage.Create(ctx, fqn, obj); err != nil {
		return nil, err
	}

	prefix := fmt.Sprintf("projects/%s/locations/%s", name.Project.ID, name.Location)
	metadata := &pb.OperationMetadata{
		ApiVersion: "v1",
		CreateTime: timestamppb.New(now),
		Target:     name.String(),
		Verb:       "create",
	}
	return s.operations.StartLRO(ctx, prefix, metadata, func() (proto.Message, error) {
		metadata.EndTime = timestamppb.Now()
		if err := s.storage.Update(ctx, fqn, obj); err != nil {
			return nil, err
		}
		return obj, nil
	})
}

func (s *VMwareEngineV1) UpdatePrivateConnection(ctx context.Context, req *pb.UpdatePrivateConnectionRequest) (*longrunningpb.Operation, error) {
	name, err := s.parsePrivateConnectionName(req.GetPrivateConnection().GetName())
	if err != nil {
		return nil, err
	}
	fqn := name.String()

	obj := &pb.PrivateConnection{}
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
			obj.Description = req.GetPrivateConnection().Description
		case "routing_mode":
			obj.RoutingMode = req.GetPrivateConnection().RoutingMode
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

func (s *VMwareEngineV1) DeletePrivateConnection(ctx context.Context, req *pb.DeletePrivateConnectionRequest) (*longrunningpb.Operation, error) {
	name, err := s.parsePrivateConnectionName(req.Name)
	if err != nil {
		return nil, err
	}

	fqn := name.String()

	deleted := &pb.PrivateConnection{}
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

type privateConnectionName struct {
	Project             *projects.ProjectData
	Location            string
	PrivateConnectionID string
}

func (n *privateConnectionName) String() string {
	return fmt.Sprintf("projects/%s/locations/%s/privateConnections/%s", n.Project.ID, n.Location, n.PrivateConnectionID)
}

func (s *MockService) parsePrivateConnectionName(name string) (*privateConnectionName, error) {
	tokens := strings.Split(name, "/")

	if len(tokens) == 6 && tokens[0] == "projects" && tokens[2] == "locations" && tokens[4] == "privateConnections" {
		project, err := s.Projects.GetProjectByID(tokens[1])
		if err != nil {
			return nil, err
		}

		name := &privateConnectionName{
			Project:             project,
			Location:            tokens[3],
			PrivateConnectionID: tokens[5],
		}

		return name, nil
	}

	return nil, status.Errorf(codes.InvalidArgument, "name %q is not valid", name)
}
