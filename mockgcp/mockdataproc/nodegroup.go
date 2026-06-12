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
// proto.service: google.cloud.dataproc.v1.NodeGroupController
// proto.message: google.cloud.dataproc.v1.NodeGroup

package mockdataproc

import (
	"context"
	"fmt"
	"strings"

	pb "cloud.google.com/go/dataproc/v2/apiv1/dataprocpb"
	longrunningpb "google.golang.org/genproto/googleapis/longrunning"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/proto"
)

type nodeGroupControllerServer struct {
	*MockService
	pb.UnimplementedNodeGroupControllerServer
}

func (s *nodeGroupControllerServer) GetNodeGroup(ctx context.Context, req *pb.GetNodeGroupRequest) (*pb.NodeGroup, error) {
	name, err := s.parseNodeGroupName(req.Name)
	if err != nil {
		return nil, err
	}

	fqn := name.String()

	obj := &pb.NodeGroup{}
	if err := s.storage.Get(ctx, fqn, obj); err != nil {
		return nil, err
	}

	return obj, nil
}

func (s *nodeGroupControllerServer) CreateNodeGroup(ctx context.Context, req *pb.CreateNodeGroupRequest) (*longrunningpb.Operation, error) {
	name, err := s.parseNodeGroupName(req.Parent + "/nodeGroups/" + req.NodeGroupId)
	if err != nil {
		return nil, err
	}

	fqn := name.String()

	obj := proto.Clone(req.NodeGroup).(*pb.NodeGroup)
	obj.Name = fqn

	if err := s.storage.Create(ctx, fqn, obj); err != nil {
		return nil, err
	}

	lroPrefix := fmt.Sprintf("projects/%s/regions/%s", name.Project, name.Region)
	lroMetadata := &pb.NodeGroupOperationMetadata{
		NodeGroupId:   name.ID,
		ClusterUuid:   "cluster-uuid",
		OperationType: pb.NodeGroupOperationMetadata_CREATE,
		Description:   "Create Node Group",
	}
	return s.operations.StartLRO(ctx, lroPrefix, lroMetadata, func() (proto.Message, error) {
		return obj, nil
	})
}

func (s *nodeGroupControllerServer) ResizeNodeGroup(ctx context.Context, req *pb.ResizeNodeGroupRequest) (*longrunningpb.Operation, error) {
	name, err := s.parseNodeGroupName(req.Name)
	if err != nil {
		return nil, err
	}

	fqn := name.String()

	obj := &pb.NodeGroup{}
	if err := s.storage.Get(ctx, fqn, obj); err != nil {
		return nil, err
	}

	if obj.NodeGroupConfig == nil {
		obj.NodeGroupConfig = &pb.InstanceGroupConfig{}
	}
	obj.NodeGroupConfig.NumInstances = req.Size

	if err := s.storage.Update(ctx, fqn, obj); err != nil {
		return nil, err
	}

	lroPrefix := fmt.Sprintf("projects/%s/regions/%s", name.Project, name.Region)
	lroMetadata := &pb.NodeGroupOperationMetadata{
		NodeGroupId:   name.ID,
		ClusterUuid:   "cluster-uuid",
		OperationType: pb.NodeGroupOperationMetadata_RESIZE,
		Description:   "Resize Node Group",
	}
	return s.operations.StartLRO(ctx, lroPrefix, lroMetadata, func() (proto.Message, error) {
		return obj, nil
	})
}

type nodeGroupName struct {
	Project string
	Region  string
	Cluster string
	ID      string
}

func (n *nodeGroupName) String() string {
	return fmt.Sprintf("projects/%s/regions/%s/clusters/%s/nodeGroups/%s", n.Project, n.Region, n.Cluster, n.ID)
}

func (s *MockService) parseNodeGroupName(name string) (*nodeGroupName, error) {
	tokens := strings.Split(name, "/")
	if len(tokens) != 8 || tokens[0] != "projects" || tokens[2] != "regions" || tokens[4] != "clusters" || tokens[6] != "nodeGroups" {
		return nil, status.Errorf(codes.InvalidArgument, "invalid node group name %q", name)
	}
	return &nodeGroupName{
		Project: tokens[1],
		Region:  tokens[3],
		Cluster: tokens[5],
		ID:      tokens[7],
	}, nil
}
