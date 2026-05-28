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

	pb "github.com/GoogleCloudPlatform/k8s-config-connector/mockgcp/generated/mockgcp/cloud/tpu/v1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/mockgcp/pkg/storage"
	longrunningpb "google.golang.org/genproto/googleapis/longrunning"
)

func (s *TpuV1Server) GetNode(ctx context.Context, req *pb.GetNodeRequest) (*pb.Node, error) {
	name, err := s.parseNodeName(req.Name)
	if err != nil {
		return nil, err
	}

	fqn := name.String()

	obj := &pb.Node{}
	if err := s.storage.Get(ctx, fqn, obj); err != nil {
		if status.Code(err) == codes.NotFound {
			return nil, status.Errorf(codes.NotFound, "TPU node %q not found", fqn)
		}
		return nil, err
	}

	return obj, nil
}

func (s *TpuV1Server) ListNodes(ctx context.Context, req *pb.ListNodesRequest) (*pb.ListNodesResponse, error) {
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

func (s *TpuV1Server) CreateNode(ctx context.Context, req *pb.CreateNodeRequest) (*longrunningpb.Operation, error) {
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

	s.populateNodeV1(obj, name)

	if err := s.storage.Create(ctx, fqn, obj); err != nil {
		return nil, err
	}

	lroPrefix := fmt.Sprintf("projects/%s/locations/%s", name.Project.ID, name.Location)
	lroMetadata := &pb.OperationMetadata{
		CreateTime: timestamppb.New(now),
		Target:     name.String(),
		Verb:       "create",
		ApiVersion: "v1",
	}

	return s.operations.StartLRO(ctx, lroPrefix, lroMetadata, func() (proto.Message, error) {
		lroMetadata.EndTime = timestamppb.Now()

		update := func(obj *pb.Node) (*pb.Node, error) {
			obj.State = pb.Node_READY
			return obj, nil
		}

		obj, err := s.updateTPUNodeV1(ctx, fqn, update)
		if err != nil {
			return nil, err
		}

		return obj, nil
	})
}

func (s *TpuV1Server) populateNodeV1(obj *pb.Node, name *nodeName) {
	if obj.Network == "" {
		obj.Network = "default"
	}
	if !strings.Contains(obj.Network, "/") {
		obj.Network = fmt.Sprintf("projects/%s/global/networks/%s", name.Project.ID, obj.Network)
	}

	if obj.CidrBlock == "" {
		obj.CidrBlock = "10.2.0.0/29"
	}

	if len(obj.NetworkEndpoints) == 0 {
		obj.NetworkEndpoints = []*pb.NetworkEndpoint{
			{
				IpAddress: "10.2.0.2",
				Port:      8470,
			},
		}
	}

	if obj.ServiceAccount == "" {
		obj.ServiceAccount = fmt.Sprintf("%d-compute@developer.gserviceaccount.com", name.Project.Number)
	}
}

func (s *TpuV1Server) DeleteNode(ctx context.Context, req *pb.DeleteNodeRequest) (*longrunningpb.Operation, error) {
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
	lroMetadata := &pb.OperationMetadata{
		CreateTime: timestamppb.New(now),
		Target:     name.String(),
		Verb:       "delete",
		ApiVersion: "v1",
	}

	return s.operations.StartLRO(ctx, lroPrefix, lroMetadata, func() (proto.Message, error) {
		lroMetadata.EndTime = timestamppb.Now()
		return &emptypb.Empty{}, nil
	})
}

func (s *TpuV1Server) updateTPUNodeV1(ctx context.Context, fqn string, update func(obj *pb.Node) (*pb.Node, error)) (*pb.Node, error) {
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
