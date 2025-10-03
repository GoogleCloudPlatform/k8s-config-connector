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
// proto.message: google.cloud.vmwareengine.v1.VmwareEngineNetwork

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

	"github.com/GoogleCloudPlatform/k8s-config-connector/mockgcp/common/fields"
	"github.com/GoogleCloudPlatform/k8s-config-connector/mockgcp/common/projects"
	pb "github.com/GoogleCloudPlatform/k8s-config-connector/mockgcp/generated/mockgcp/cloud/vmwareengine/v1"
	"github.com/google/uuid"
	longrunningpb "google.golang.org/genproto/googleapis/longrunning"
)

func (s *VMwareEngineV1) GetVmwareEngineNetwork(ctx context.Context, req *pb.GetVmwareEngineNetworkRequest) (*pb.VmwareEngineNetwork, error) {
	name, err := s.parseVmwareEngineNetworkName(req.Name)
	if err != nil {
		return nil, err
	}

	fqn := name.String()

	obj := &pb.VmwareEngineNetwork{}
	if err := s.storage.Get(ctx, fqn, obj); err != nil {
		if status.Code(err) == codes.NotFound {
			return nil, status.Errorf(codes.NotFound, "Resource '%s' was not found", fqn)
		}
		return nil, err
	}

	return obj, nil
}

func (s *VMwareEngineV1) CreateVmwareEngineNetwork(ctx context.Context, req *pb.CreateVmwareEngineNetworkRequest) (*longrunningpb.Operation, error) {
	reqName := req.Parent + "/vmwareEngineNetworks/" + req.VmwareEngineNetworkId
	name, err := s.parseVmwareEngineNetworkName(reqName)
	if err != nil {
		return nil, err
	}

	fqn := name.String()

	now := time.Now()

	obj := proto.Clone(req.GetVmwareEngineNetwork()).(*pb.VmwareEngineNetwork)
	obj.Name = fqn
	obj.CreateTime = timestamppb.New(now)
	obj.State = pb.VmwareEngineNetwork_ACTIVE
	obj.Uid = "111111111111111111111"
	obj.Etag = fields.ComputeWeakEtag(obj)
	s.generateVPCNetworks(obj)

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
		obj.UpdateTime = timestamppb.New(now)
		if err := s.storage.Update(ctx, fqn, obj); err != nil {
			return nil, err
		}
		return obj, nil
	})
}

func (s *VMwareEngineV1) UpdateVmwareEngineNetwork(ctx context.Context, req *pb.UpdateVmwareEngineNetworkRequest) (*longrunningpb.Operation, error) {
	name, err := s.parseVmwareEngineNetworkName(req.GetVmwareEngineNetwork().GetName())
	if err != nil {
		return nil, err
	}
	fqn := name.String()

	obj := &pb.VmwareEngineNetwork{}
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
			obj.Description = req.GetVmwareEngineNetwork().Description
		case "state":
			obj.State = req.GetVmwareEngineNetwork().State
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

func (s *VMwareEngineV1) DeleteVmwareEngineNetwork(ctx context.Context, req *pb.DeleteVmwareEngineNetworkRequest) (*longrunningpb.Operation, error) {
	name, err := s.parseVmwareEngineNetworkName(req.Name)
	if err != nil {
		return nil, err
	}

	fqn := name.String()

	deleted := &pb.VmwareEngineNetwork{}
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

type vmwareEngineNetworkName struct {
	Project                 *projects.ProjectData
	Location                string
	VmwareEngineNetworkName string
}

func (n *vmwareEngineNetworkName) String() string {
	return fmt.Sprintf("projects/%s/locations/%s/vmwareEngineNetworks/%s", n.Project.ID, n.Location, n.VmwareEngineNetworkName)
}

// parseVmwareEngineNetworkName parses a string into a vmwareEngineNetworkName.
// The expected form is `projects/*/locations/*/VmwareEngineNetwork/*`.
func (s *MockService) parseVmwareEngineNetworkName(name string) (*vmwareEngineNetworkName, error) {
	tokens := strings.Split(name, "/")

	if len(tokens) == 6 && tokens[0] == "projects" && tokens[2] == "locations" && tokens[4] == "vmwareEngineNetworks" {
		project, err := s.Projects.GetProjectByID(tokens[1])
		if err != nil {
			return nil, err
		}

		name := &vmwareEngineNetworkName{
			Project:                 project,
			Location:                tokens[3],
			VmwareEngineNetworkName: tokens[5],
		}

		return name, nil
	}

	return nil, status.Errorf(codes.InvalidArgument, "name %q is not valid", name)
}

func (s *VMwareEngineV1) generateVPCNetworks(obj *pb.VmwareEngineNetwork) {
	if obj.VpcNetworks == nil || len(obj.VpcNetworks) == 0 {
		uuid := uuid.New().String()
		projectID := "b3e854f0b4bedfea6-tp"
		obj.VpcNetworks = []*pb.VmwareEngineNetwork_VpcNetwork{
			{
				Network: fmt.Sprintf("projects/%s/global/networks/internet-%s", projectID, uuid),
				Type:    pb.VmwareEngineNetwork_VpcNetwork_INTERNET,
			},
			{
				Network: fmt.Sprintf("projects/%s/global/networks/intranet-%s", projectID, uuid),
				Type:    pb.VmwareEngineNetwork_VpcNetwork_INTRANET,
			},
			{
				Network: fmt.Sprintf("projects/%s/global/networks/gcp-%s", projectID, uuid),
				Type:    pb.VmwareEngineNetwork_VpcNetwork_GOOGLE_CLOUD,
			},
		}
	}
}
