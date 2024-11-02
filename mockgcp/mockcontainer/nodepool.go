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
	"fmt"
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

	cluster, err := s.GetCluster(ctx, &pb.GetClusterRequest{Name: req.GetParent()})
	if err != nil {
		return nil, err
	}

	fqn := name.String()

	obj := proto.Clone(req.NodePool).(*pb.NodePool)

	obj.SelfLink = fmt.Sprintf("https://container.googlerapis.com/v1beta1/%s", fqn)

	if err := s.populateNodePoolDefaults(cluster, obj); err != nil {
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

func (s *ClusterManagerV1) populateNodePoolDefaults(cluster *pb.Cluster, obj *pb.NodePool) error {
	obj.Status = pb.NodePool_RUNNING
	if obj.Version == "" {
		obj.Version = cluster.CurrentNodeVersion
	}

	if obj.Config == nil {
		obj.Config = &pb.NodeConfig{}
	}
	if err := s.populateNodeConfig(obj.Config); err != nil {
		return err
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
	if obj.NetworkConfig.PodIpv4CidrBlock == "" {
		obj.NetworkConfig.PodIpv4CidrBlock = "10.92.0.0/14"
	}
	if obj.NetworkConfig.PodIpv4RangeUtilization == 0 {
		obj.NetworkConfig.PodIpv4RangeUtilization = 0.001
	}
	if obj.NetworkConfig.PodRange == "" {
		obj.NetworkConfig.PodRange = obj.Name + "-pods-12345678"
	}

	if obj.PodIpv4CidrSize == 0 {
		obj.PodIpv4CidrSize = 24
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

func (s *ClusterManagerV1) populateNodeConfig(obj *pb.NodeConfig) error {
	if obj.DiskSizeGb == 0 {
		obj.DiskSizeGb = 100
	}
	if obj.DiskType == "" {
		obj.DiskType = "pd-balanced"
	}

	if obj.ImageType == "" {
		obj.ImageType = "COS_CONTAINERD"
	}

	if obj.MachineType == "" {
		obj.MachineType = "e2-medium"
	}

	if obj.Metadata == nil {
		obj.Metadata = make(map[string]string)
	}

	if obj.Metadata["disable-legacy-endpoints"] == "" {
		obj.Metadata["disable-legacy-endpoints"] = "true"
	}

	if obj.OauthScopes == nil {
		obj.OauthScopes = []string{
			"https://www.googleapis.com/auth/devstorage.read_only",
			"https://www.googleapis.com/auth/logging.write",
			"https://www.googleapis.com/auth/monitoring",
			"https://www.googleapis.com/auth/service.management.readonly",
			"https://www.googleapis.com/auth/servicecontrol",
			"https://www.googleapis.com/auth/trace.append",
		}
	}

	if obj.ServiceAccount == "" {
		obj.ServiceAccount = "default"
	}

	if obj.ShieldedInstanceConfig == nil {
		obj.ShieldedInstanceConfig = &pb.ShieldedInstanceConfig{
			EnableIntegrityMonitoring: true,
		}
	}

	if obj.WindowsNodeConfig == nil {
		obj.WindowsNodeConfig = &pb.WindowsNodeConfig{}
	}

	return nil
}

func (s *ClusterManagerV1) populateAutoprovisioningNodePoolDefaults(obj *pb.AutoprovisioningNodePoolDefaults) error {

	if obj.OauthScopes == nil {
		obj.OauthScopes = []string{
			"https://www.googleapis.com/auth/userinfo.email",
			"https://www.googleapis.com/auth/cloud-platform",
		}
	}

	if obj.UpgradeSettings == nil {
		obj.UpgradeSettings = &pb.NodePool_UpgradeSettings{}
	}
	if obj.UpgradeSettings.Strategy == nil {
		obj.UpgradeSettings.Strategy = PtrTo(pb.NodePoolUpdateStrategy_SURGE)
	}

	if obj.Management == nil {
		obj.Management = &pb.NodeManagement{
			AutoRepair:  true,
			AutoUpgrade: true,
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
