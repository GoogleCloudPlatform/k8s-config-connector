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
// proto.service: google.cloud.dataproc.v1
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
	labels := make(map[string]string)
	labels["cnrm-test"] = "true"
	labels["managed-by-cnrm"] = "true"

	obj := &pb.NodeGroup{}
	obj.Name = fqn
	obj.Labels = labels
	if err := s.storage.Create(ctx, fqn, obj); err != nil {
		return nil, err
	}

	lroPrefix := fmt.Sprintf("projects/%s/regions/%s", name.Project, name.Location)
	lroMetadata := &pb.NodeGroupOperationMetadata{
		NodeGroupId:   fqn,
		OperationType: pb.NodeGroupOperationMetadata_CREATE,
		Description:   "NodeGroup",
		ClusterUuid:   name.Cluster,
		Labels:        labels,
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

	lroPrefix := fmt.Sprintf("projects/%s/regions/%s", name.Project, name.Location)
	lroMetadata := &pb.NodeGroupOperationMetadata{
		NodeGroupId:   fqn,
		OperationType: pb.NodeGroupOperationMetadata_RESIZE,
		Description:   "NodeGroup",
		ClusterUuid:   name.Cluster,
	}
	return s.operations.StartLRO(ctx, lroPrefix, lroMetadata, func() (proto.Message, error) {

		return obj, nil
	})
}

type nodeGroupName struct {
	Project   string
	Cluster   string
	Location  string
	NodeGroup string
}

func (n *nodeGroupName) String() string {
	return fmt.Sprintf("projects/%s/regions/%s/clusters/%s/nodeGroups/%s", n.Project, n.Location, n.Cluster, n.NodeGroup)
}

// parseNodeGroupName parses a string into a batchName.
// The expected form is `projects/*/regions/*/clusters/*/nodeGroups/*`.
func (s *MockService) parseNodeGroupName(name string) (*nodeGroupName, error) {
	tokens := strings.Split(name, "/")

	if len(tokens) == 8 && tokens[0] == "projects" && tokens[2] == "regions" && tokens[4] == "clusters" && tokens[6] == "nodeGroups" {
		name := &nodeGroupName{
			Project:   tokens[1],
			Location:  tokens[3],
			Cluster:   tokens[5],
			NodeGroup: tokens[7],
		}

		return name, nil
	}

	return nil, status.Errorf(codes.InvalidArgument, "name %q is not valid", name)
}
