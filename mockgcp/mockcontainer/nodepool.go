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

package mockcontainer

import (
	"context"
	"strings"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/encoding/prototext"
	"google.golang.org/protobuf/proto"
	"k8s.io/klog/v2"

	"github.com/GoogleCloudPlatform/k8s-config-connector/mockgcp/common/projects"
	pb "github.com/GoogleCloudPlatform/k8s-config-connector/mockgcp/generated/mockgcp/container/v1beta1"
)

func (s *ClusterManagerV1) GetNodePool(ctx context.Context, req *pb.GetNodePoolRequest) (*pb.NodePool, error) {
	name, err := s.parseNodePoolName(req.Name)
	if err != nil {
		return nil, err
	}

	fqn := name.String()

	obj := &pb.NodePool{}
	if err := s.storage.Get(ctx, fqn, obj); err != nil {
		return nil, err
	}

	return obj, nil
}

func (s *ClusterManagerV1) CreateNodePool(ctx context.Context, req *pb.CreateNodePoolRequest) (*pb.Operation, error) {
	reqName := req.GetParent() + "/nodePools/" + req.GetNodePool().GetName()
	name, err := s.parseNodePoolName(reqName)
	if err != nil {
		return nil, err
	}

	fqn := name.String()

	obj := proto.Clone(req.NodePool).(*pb.NodePool)
	if err := s.populateNodePoolDefaults(obj); err != nil {
		return nil, status.Errorf(codes.InvalidArgument, "%v", err)
	}

	if err := s.storage.Create(ctx, fqn, obj); err != nil {
		return nil, err
	}

	return s.newLRO(ctx, name.Project.ID)
}

func (s *ClusterManagerV1) populateNodePoolDefaults(obj *pb.NodePool) error {
	obj.Status = pb.NodePool_RUNNING

	if obj.Management == nil {
		obj.Management = &pb.NodeManagement{
			AutoUpgrade: true,
			AutoRepair:  true,
		}
	}

	if obj.MaxPodsConstraint == nil {
		obj.MaxPodsConstraint = &pb.MaxPodsConstraint{
			MaxPodsPerNode: 110,
		}
	}

	return nil
}

func (s *ClusterManagerV1) UpdateNodePool(ctx context.Context, req *pb.UpdateNodePoolRequest) (*pb.Operation, error) {
	reqName := req.GetName()

	name, err := s.parseNodePoolName(reqName)
	if err != nil {
		return nil, err
	}

	fqn := name.String()
	obj := &pb.NodePool{}
	if err := s.storage.Get(ctx, fqn, obj); err != nil {
		return nil, err
	}

	klog.Infof("UpdateNodePool %v", prototext.Format(req))

	update := proto.Clone(req).(*pb.UpdateNodePoolRequest)
	update.Name = ""

	// TODO: Support more updates!

	if !proto.Equal(update, &pb.UpdateNodePoolRequest{}) {
		return nil, status.Errorf(codes.InvalidArgument, "update was not fully implemented UpdateNodePoolRequest=%v", prototext.Format(update))
	}

	// // Required. The update mask applies to the resource.
	// paths := req.GetUpdateMask().GetPaths()
	// if len(paths) == 0 {
	// 	klog.Warningf("update_mask was not provided in request, should be required")
	// }

	// // TODO: Some sort of helper for fieldmask?
	// for _, path := range paths {
	// 	switch path {
	// 	case "description":
	// 		obj.Description = req.GetNodePool().GetDescription()
	// 	case "labels":
	// 		obj.Labels = req.GetNodePool().GetLabels()
	// 	case "timeout":
	// 		obj.Timeout = req.GetNodePool().GetTimeout()
	// 	default:
	// 		return nil, status.Errorf(codes.InvalidArgument, "update_mask path %q not valid", path)
	// 	}
	// }

	if err := s.storage.Update(ctx, fqn, obj); err != nil {
		return nil, err
	}

	return s.newLRO(ctx, name.Project.ID)
}

func (s *ClusterManagerV1) DeleteNodePool(ctx context.Context, req *pb.DeleteNodePoolRequest) (*pb.Operation, error) {
	name, err := s.parseNodePoolName(req.Name)
	if err != nil {
		return nil, err
	}

	fqn := name.String()

	oldObj := &pb.NodePool{}
	if err := s.storage.Delete(ctx, fqn, oldObj); err != nil {
		return nil, err
	}

	return s.newLRO(ctx, name.Project.ID)
}

type nodePoolName struct {
	Project  *projects.ProjectData
	Location string
	Cluster  string
	NodePool string
}

func (n *nodePoolName) String() string {
	return "projects/" + n.Project.ID + "/locations/" + n.Location + "/clusters/" + n.Cluster + "/nodePools/" + n.NodePool
}

func (n *nodePoolName) ClusterName() *clusterName {
	return &clusterName{
		Project:  n.Project,
		Location: n.Location,
		Cluster:  n.Cluster,
	}
}

// parseNodePoolName parses a string into a nodePoolName.
// The expected form is `projects/*/locations/*/clusters/*`.
func (s *MockService) parseNodePoolName(name string) (*nodePoolName, error) {
	tokens := strings.Split(name, "/")

	if len(tokens) == 8 && tokens[0] == "projects" && tokens[2] == "locations" && tokens[4] == "clusters" && tokens[6] == "nodePools" {
		project, err := s.projects.GetProjectByID(tokens[1])
		if err != nil {
			return nil, err
		}

		name := &nodePoolName{
			Project:  project,
			Location: tokens[3],
			Cluster:  tokens[5],
			NodePool: tokens[7],
		}

		return name, nil
	} else {
		return nil, status.Errorf(codes.InvalidArgument, "name %q is not valid", name)
	}
}
