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
		return nil, err
	}

	if err := s.storage.Create(ctx, fqn, obj); err != nil {
		return nil, err
	}

	op := &pb.Operation{
		Zone:          name.Location,
		OperationType: pb.Operation_CREATE_NODE_POOL,
		TargetLink:    obj.SelfLink,
	}
	return s.startLRO(ctx, name.Project, op, func() (proto.Message, error) {
		return obj, nil
	})
}

func (s *ClusterManagerV1) populateNodePoolDefaults(obj *pb.NodePool) error {
	obj.Status = pb.NodePool_RUNNING

	if obj.Config == nil {
		obj.Config = &pb.NodeConfig{}
	}

	if obj.Config.DiskSizeGb == 0 {
		obj.Config.DiskSizeGb = 100
	}
	if obj.Config.DiskType == "" {
		obj.Config.DiskType = "pd-balanced"
	}
	if obj.Config.ImageType == "" {
		obj.Config.ImageType = "COS_CONTAINERD"
	}

	if obj.Config.MachineType == "" {
		obj.Config.MachineType = "e2-standard-4"
	}

	if obj.Config.Metadata == nil {
		obj.Config.Metadata = make(map[string]string)
	}

	if obj.Config.Metadata["disable-legacy-endpoints"] == "" {
		obj.Config.Metadata["disable-legacy-endpoints"] = "true"
	}

	if obj.Config.OauthScopes == nil {
		obj.Config.OauthScopes = []string{
			"https://www.googleapis.com/auth/devstorage.read_only",
			"https://www.googleapis.com/auth/logging.write",
			"https://www.googleapis.com/auth/monitoring",
			"https://www.googleapis.com/auth/service.management.readonly",
			"https://www.googleapis.com/auth/servicecontrol",
			"https://www.googleapis.com/auth/trace.append",
		}
	}

	if obj.Config.ServiceAccount == "" {
		obj.Config.ServiceAccount = "default"
	}

	if obj.Config.ShieldedInstanceConfig == nil {
		obj.Config.ShieldedInstanceConfig = &pb.ShieldedInstanceConfig{
			EnableIntegrityMonitoring: true,
		}
	}

	if obj.Config.WindowsNodeConfig == nil {
		obj.Config.WindowsNodeConfig = &pb.WindowsNodeConfig{}
	}

	if obj.InitialNodeCount == 0 {
		obj.InitialNodeCount = 1
	}

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

	if obj.NetworkConfig == nil {
		obj.NetworkConfig = &pb.NodeNetworkConfig{}
	}

	if obj.NetworkConfig.EnablePrivateNodes == nil {
		obj.NetworkConfig.EnablePrivateNodes = PtrTo(false)
	}

	if obj.UpgradeSettings == nil {
		obj.UpgradeSettings = &pb.NodePool_UpgradeSettings{
			MaxSurge:       1,
			MaxUnavailable: 0,
			Strategy:       PtrTo(pb.NodePoolUpdateStrategy_SURGE),
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

	if update.Taints != nil {
		obj.Config.Taints = update.GetTaints().Taints
		update.Taints = nil
	}

	// TODO: Support more updates!

	if !proto.Equal(update, &pb.UpdateNodePoolRequest{}) {
		return nil, status.Errorf(codes.InvalidArgument, "update was not fully implemented UpdateNodePoolRequest=%v", prototext.Format(update))
	}

	if err := s.storage.Update(ctx, fqn, obj); err != nil {
		return nil, err
	}

	op := &pb.Operation{
		Zone:       name.Location,
		TargetLink: obj.SelfLink,
	}
	return s.startLRO(ctx, name.Project, op, func() (proto.Message, error) {
		return obj, nil
	})
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

	op := &pb.Operation{
		Zone:          name.Location,
		OperationType: pb.Operation_DELETE_NODE_POOL,
		TargetLink:    oldObj.SelfLink,
	}
	return s.startLRO(ctx, name.Project, op, func() (proto.Message, error) {
		return oldObj, nil
	})
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
		project, err := s.Projects.GetProjectByID(tokens[1])
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
