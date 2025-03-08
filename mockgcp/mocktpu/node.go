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
// proto.service: google.cloud.tpu.v2.Tpu
// proto.message: google.cloud.tpu.v2.Node

package mocktpu

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

	pb "cloud.google.com/go/tpu/apiv2/tpupb"
	"github.com/GoogleCloudPlatform/k8s-config-connector/mockgcp/common/projects"
	"github.com/GoogleCloudPlatform/k8s-config-connector/mockgcp/pkg/storage"
	commonpb "google.golang.org/genproto/googleapis/cloud/common"
	longrunningpb "google.golang.org/genproto/googleapis/longrunning"
)

func (s *TpuServer) GetNode(ctx context.Context, req *pb.GetNodeRequest) (*pb.Node, error) {
	name, err := s.parseNodeName(req.Name)
	if err != nil {
		return nil, err
	}

	fqn := name.String()

	obj := &pb.Node{}
	if err := s.storage.Get(ctx, fqn, obj); err != nil {
		if status.Code(err) == codes.NotFound {
			return nil, status.Errorf(codes.NotFound, "Resource '%v' was not found", fqn)
		}
		return nil, err
	}

	return obj, nil
}

func (s *TpuServer) ListNodes(ctx context.Context, req *pb.ListNodesRequest) (*pb.ListNodesResponse, error) {
	response := &pb.ListNodesResponse{}

	_, err := s.parseNodeName(req.Parent)
	if err != nil {
		return nil, err
	}

	findPrefix := req.Parent + "/nodes/"

	nodeKind := (&pb.Node{}).ProtoReflect().Descriptor()
	if err := s.storage.List(ctx, nodeKind, storage.ListOptions{}, func(obj proto.Message) error {
		node := obj.(*pb.Node)
		if strings.HasPrefix(node.Name, findPrefix) {
			response.Nodes = append(response.Nodes, node)
		}
		return nil
	}); err != nil {
		return nil, fmt.Errorf("error listing %s: %w", nodeKind.Name(), err)
	}

	return response, nil
}

func (s *TpuServer) CreateNode(ctx context.Context, req *pb.CreateNodeRequest) (*longrunningpb.Operation, error) {
	reqName := req.Parent + "/nodes/" + req.NodeId
	name, err := s.parseNodeName(reqName)
	if err != nil {
		return nil, err
	}

	fqn := name.String()

	now := time.Now()

	obj := proto.Clone(req.GetNode()).(*pb.Node)
	obj.Name = fqn
	obj.CreateTime = timestamppb.New(now)
	obj.State = pb.Node_CREATING

	obj.Id = time.Now().UnixNano()

	s.populateNode(obj, name)

	if err := s.storage.Create(ctx, fqn, obj); err != nil {
		return nil, err
	}

	lroPrefix := fmt.Sprintf("projects/%s/locations/%s", name.Project.ID, name.Location)
	lroMetadata := &commonpb.OperationMetadata{
		CreateTime: timestamppb.New(now),
		Target:     name.String(),
		Verb:       "create",
		ApiVersion: "v2",
	}

	return s.operations.StartLRO(ctx, lroPrefix, lroMetadata, func() (proto.Message, error) {
		lroMetadata.EndTime = timestamppb.Now()

		update := func(obj *pb.Node) (*pb.Node, error) {
			obj.State = pb.Node_READY
			obj.Health = pb.Node_HEALTHY

			return obj, nil
		}

		obj, err := s.updateTPUNode(ctx, fqn, update)
		if err != nil {
			return nil, err
		}

		// Return without Health
		ret := ProtoClone(obj)
		ret.Health = pb.Node_HEALTH_UNSPECIFIED
		return ret, nil
	})
}

func (s *TpuServer) populateNode(obj *pb.Node, name *nodeName) {

	if obj.ApiVersion == pb.Node_API_VERSION_UNSPECIFIED {
		obj.ApiVersion = pb.Node_V2
	}
	if obj.AcceleratorConfig == nil {
		obj.AcceleratorConfig = &pb.AcceleratorConfig{}
	}
	if obj.AcceleratorConfig.Type == pb.AcceleratorConfig_TYPE_UNSPECIFIED {
		obj.AcceleratorConfig.Type = pb.AcceleratorConfig_V5P
	}
	if obj.AcceleratorConfig.Topology == "" {
		obj.AcceleratorConfig.Topology = "2x2x1"
	}

	if obj.NetworkConfig == nil {
		obj.NetworkConfig = &pb.NetworkConfig{}
	}
	if obj.NetworkConfig.Network == "" {
		obj.NetworkConfig.Network = "default"
	}
	if tokens := strings.Split(obj.NetworkConfig.Network, "/"); len(tokens) == 1 {
		obj.NetworkConfig.Network = fmt.Sprintf("projects/%s/global/networks/%s", name.Project.ID, tokens[0])
	}
	if obj.NetworkConfig.Subnetwork == "" {
		obj.NetworkConfig.Subnetwork = fmt.Sprintf("projects/%s/regions/%s/subnetworks/default", name.Project.ID, regionForLocation(name.Location))
	}
	if obj.SchedulingConfig == nil {
		obj.SchedulingConfig = &pb.SchedulingConfig{}
	}
	if obj.ServiceAccount == nil {
		obj.ServiceAccount = &pb.ServiceAccount{}
	}
	if obj.ServiceAccount.Email == "" {
		obj.ServiceAccount.Email = fmt.Sprintf("%d-compute@developer.gserviceaccount.com", name.Project.Number)
	}
	if obj.ServiceAccount.Scope == nil {
		obj.ServiceAccount.Scope = []string{
			"https://www.googleapis.com/auth/devstorage.read_write",
			"https://www.googleapis.com/auth/logging.write",
			"https://www.googleapis.com/auth/service.management",
			"https://www.googleapis.com/auth/servicecontrol",
			"https://www.googleapis.com/auth/cloud-platform",
			"https://www.googleapis.com/auth/pubsub",
		}
	}
	if obj.ShieldedInstanceConfig == nil {
		obj.ShieldedInstanceConfig = &pb.ShieldedInstanceConfig{}
	}
	if obj.CidrBlock == "" {
		obj.CidrBlock = "10.32.0.0/20"
	}

	if len(obj.NetworkEndpoints) == 0 {
		obj.NetworkEndpoints = []*pb.NetworkEndpoint{
			{
				IpAddress:    "10.32.0.27",
				Port:         8470,
				AccessConfig: &pb.AccessConfig{},
			},
		}
		if obj.GetNetworkConfig().GetEnableExternalIps() {
			obj.NetworkEndpoints[0].AccessConfig.ExternalIp = "8.8.8.8"
		}
	}
}

func regionForLocation(location string) string {
	tokens := strings.Split(location, "-")
	if len(tokens) == 3 {
		return tokens[0] + "-" + tokens[1]
	}
	return location
}

func (s *TpuServer) UpdateNode(ctx context.Context, req *pb.UpdateNodeRequest) (*longrunningpb.Operation, error) {
	name, err := s.parseNodeName(req.GetNode().GetName())
	if err != nil {
		return nil, err
	}
	fqn := name.String()
	existing := &pb.Node{}
	if err := s.storage.Get(ctx, fqn, existing); err != nil {
		return nil, err
	}

	now := time.Now()

	updated := ProtoClone(existing)
	updated.Name = name.String()

	// Required. The update mask applies to the resource.
	paths := req.GetUpdateMask().GetPaths()

	// TODO: Some sort of helper for fieldmask?
	for _, path := range paths {
		switch path {
		case "description":
			updated.Description = req.GetNode().GetDescription()
		default:
			return nil, status.Errorf(codes.InvalidArgument, "update_mask path %q not supported by mock", path)
		}
	}

	s.populateNode(updated, name)

	if err := s.storage.Update(ctx, fqn, updated); err != nil {
		return nil, err
	}

	lroPrefix := fmt.Sprintf("projects/%s/locations/%s", name.Project.ID, name.Location)
	lroMetadata := &commonpb.OperationMetadata{
		CreateTime: timestamppb.New(now),
		Target:     name.String(),
		Verb:       "update",
		ApiVersion: "v2",
	}

	return s.operations.StartLRO(ctx, lroPrefix, lroMetadata, func() (proto.Message, error) {
		lroMetadata.EndTime = timestamppb.Now()
		return updated, nil
	})
}

func (s *TpuServer) DeleteNode(ctx context.Context, req *pb.DeleteNodeRequest) (*longrunningpb.Operation, error) {
	name, err := s.parseNodeName(req.Name)
	if err != nil {
		return nil, err
	}

	fqn := name.String()

	deleted := &pb.Node{}
	if err := s.storage.Delete(ctx, fqn, deleted); err != nil {
		return nil, err
	}

	now := time.Now()

	lroPrefix := fmt.Sprintf("projects/%s/locations/%s", name.Project.ID, name.Location)
	lroMetadata := &commonpb.OperationMetadata{
		CreateTime: timestamppb.New(now),
		Target:     name.String(),
		Verb:       "delete",
		ApiVersion: "v2",
	}

	return s.operations.StartLRO(ctx, lroPrefix, lroMetadata, func() (proto.Message, error) {
		lroMetadata.EndTime = timestamppb.Now()
		return &emptypb.Empty{}, nil
	})
}

func (s *TpuServer) updateTPUNode(ctx context.Context, fqn string, update func(obj *pb.Node) (*pb.Node, error)) (*pb.Node, error) {
	obj := &pb.Node{}
	if err := s.storage.Get(ctx, fqn, obj); err != nil {
		return nil, err
	}

	updateObj, err := update(obj)
	if err != nil {
		return nil, err
	}

	if err := s.storage.Update(ctx, fqn, updateObj); err != nil {
		return nil, err
	}
	return updateObj, nil
}

type nodeName struct {
	Project  *projects.ProjectData
	Location string
	Node     string
}

func (n *nodeName) String() string {
	return fmt.Sprintf("projects/%s/locations/%s/nodes/%s", n.Project.ID, n.Location, n.Node)
}

// parseNodeName parses a string into an nodeName.
// The expected form is `projects/*/locations/*/nodes/*`.
func (s *MockService) parseNodeName(name string) (*nodeName, error) {
	tokens := strings.Split(name, "/")

	if len(tokens) == 6 && tokens[0] == "projects" && tokens[2] == "locations" && tokens[4] == "nodes" {
		project, err := s.Projects.GetProjectByID(tokens[1])
		if err != nil {
			return nil, err
		}

		name := &nodeName{
			Project:  project,
			Location: tokens[3],
			Node:     tokens[5],
		}

		return name, nil
	}

	return nil, status.Errorf(codes.InvalidArgument, "name %q is not valid", name)
}
