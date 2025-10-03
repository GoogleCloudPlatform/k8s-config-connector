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
// proto.service: google.cloud.edgecontainer.v1.Service
// proto.message: google.cloud.edgecontainer.v1.VpnConnection

package mockedgecontainer

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

	"github.com/GoogleCloudPlatform/k8s-config-connector/mockgcp/common/projects"
	pb "github.com/GoogleCloudPlatform/k8s-config-connector/mockgcp/generated/mockgcp/cloud/edgecontainer/v1"
	longrunningpb "google.golang.org/genproto/googleapis/longrunning"
)

func (s *EdgeContainerV1) GetVpnConnection(ctx context.Context, req *pb.GetVpnConnectionRequest) (*pb.VpnConnection, error) {
	name, err := s.parseVpnConnectionName(req.Name)
	if err != nil {
		return nil, err
	}

	fqn := name.String()

	obj := &pb.VpnConnection{}
	if err := s.storage.Get(ctx, fqn, obj); err != nil {
		return nil, err
	}

	return obj, nil
}

func (s *EdgeContainerV1) CreateVpnConnection(ctx context.Context, req *pb.CreateVpnConnectionRequest) (*longrunningpb.Operation, error) {
	reqName := fmt.Sprintf("%s/vpnConnections/%s", req.GetParent(), req.GetVpnConnectionId())
	name, err := s.parseVpnConnectionName(reqName)
	if err != nil {
		return nil, err
	}

	fqn := name.String()

	now := time.Now()

	obj := proto.Clone(req.GetVpnConnection()).(*pb.VpnConnection)
	obj.Name = fqn
	obj.CreateTime = timestamppb.New(now)
	obj.UpdateTime = timestamppb.New(now)

	if err := s.storage.Create(ctx, fqn, obj); err != nil {
		return nil, err
	}

	lroPrefix := fmt.Sprintf("projects/%s/locations/%s", name.Project.ID, name.Location)
	lroMetadata := &pb.OperationMetadata{
		CreateTime: timestamppb.New(now),
		Target:     obj.Name,
		Verb:       "create",
	}
	return s.operations.StartLRO(ctx, lroPrefix, lroMetadata, func() (proto.Message, error) {
		lroMetadata.EndTime = timestamppb.Now()
		if err := s.storage.Get(ctx, fqn, obj); err != nil {
			return nil, err
		}
		if err := s.storage.Update(ctx, fqn, obj); err != nil {
			return nil, err
		}
		return obj, nil
	})
}

func (s *EdgeContainerV1) DeleteVpnConnection(ctx context.Context, req *pb.DeleteVpnConnectionRequest) (*longrunningpb.Operation, error) {
	name, err := s.parseVpnConnectionName(req.Name)
	if err != nil {
		return nil, err
	}

	fqn := name.String()

	deleted := &pb.VpnConnection{}
	if err := s.storage.Delete(ctx, fqn, deleted); err != nil {
		return nil, err
	}

	prefix := fmt.Sprintf("projects/%s/locations/%s", name.Project.ID, name.Location)
	metadata := &pb.OperationMetadata{
		CreateTime: timestamppb.Now(),
		Target:     fqn,
		Verb:       "delete",
	}
	return s.operations.StartLRO(ctx, prefix, metadata, func() (proto.Message, error) {
		metadata.EndTime = timestamppb.Now()
		return &emptypb.Empty{}, nil
	})
}

type vpnConnectionName struct {
	Project         *projects.ProjectData
	Location        string
	VpnConnectionID string
}

func (n *vpnConnectionName) String() string {
	return fmt.Sprintf("projects/%s/locations/%s/vpnConnections/%s", n.Project.ID, n.Location, n.VpnConnectionID)
}

// parseVpnConnectionName parses a string into a VpnConnectionName.
// The expected form is `projects/*/locations/*/VpnConnections/*`.
func (s *MockService) parseVpnConnectionName(name string) (*vpnConnectionName, error) {
	tokens := strings.Split(name, "/")

	if len(tokens) == 6 && tokens[0] == "projects" && tokens[2] == "locations" && tokens[4] == "vpnConnections" {
		project, err := s.Projects.GetProjectByID(tokens[1])
		if err != nil {
			return nil, err
		}

		name := &vpnConnectionName{
			Project:         project,
			Location:        tokens[3],
			VpnConnectionID: tokens[5],
		}

		return name, nil
	}

	return nil, status.Errorf(codes.InvalidArgument, "name %q is not valid", name)
}
