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

type ClusterManagerV1 struct {
	*MockService
	pb.UnimplementedClusterManagerServer
}

func (s *ClusterManagerV1) GetCluster(ctx context.Context, req *pb.GetClusterRequest) (*pb.Cluster, error) {
	name, err := s.parseClusterName(req.Name)
	if err != nil {
		return nil, err
	}

	fqn := name.String()

	obj := &pb.Cluster{}
	if err := s.storage.Get(ctx, fqn, obj); err != nil {
		return nil, err
	}

	return obj, nil
}

func (s *ClusterManagerV1) CreateCluster(ctx context.Context, req *pb.CreateClusterRequest) (*pb.Operation, error) {
	klog.Infof("cluster create %v", prototext.Format(req))

	reqName := req.GetParent() + "/clusters/" + req.GetCluster().GetName()
	name, err := s.parseClusterName(reqName)
	if err != nil {
		return nil, err
	}

	fqn := name.String()

	region := name.Location

	obj := proto.Clone(req.Cluster).(*pb.Cluster)
	obj.Status = pb.Cluster_RUNNING
	obj.Location = name.Location
	obj.Locations = []string{name.Location}

	if obj.NetworkConfig == nil {
		obj.NetworkConfig = &pb.NetworkConfig{}
	}
	if obj.NetworkConfig.Network == "" {
		obj.NetworkConfig.Network = fmt.Sprintf("projects/%s/global/networks/%s", name.Project.ID, "default")
	}
	if obj.NetworkConfig.Subnetwork == "" {
		obj.NetworkConfig.Network = fmt.Sprintf("projects/%s/regions/%s/subnetworks/%s", name.Project.ID, region, "default")
	}

	if len(obj.NodePools) != 0 {
		// TODO: Do we support this?
		return nil, fmt.Errorf("nodePools must be empty when creating a cluster")
	}
	defaultNodePool := &pb.NodePool{
		Name:      "default-pool",
		Status:    pb.NodePool_RUNNING,
		Locations: []string{name.Location},
	}
	if err := s.populateNodePoolDefaults(defaultNodePool); err != nil {
		return nil, status.Errorf(codes.InvalidArgument, "%v", err)
	}

	obj.NodePools = append(obj.NodePools, defaultNodePool)

	nodePools := obj.NodePools
	obj.NodePools = nil

	klog.Infof("creating cluster %v", prototext.Format(obj))
	if err := s.storage.Create(ctx, fqn, obj); err != nil {
		return nil, err
	}

	for _, nodePool := range nodePools {
		nodePoolFqn := name.String() + "/nodePools/" + nodePool.GetName()
		nodePoolObj := proto.Clone(nodePool).(*pb.NodePool)
		klog.Infof("creating nodePool %q %v", nodePoolFqn, prototext.Format(nodePoolObj))
		if err := s.storage.Create(ctx, nodePoolFqn, nodePoolObj); err != nil {
			return nil, err
		}
	}

	return s.newLRO(ctx, name.Project.ID)
}

func (s *ClusterManagerV1) UpdateCluster(ctx context.Context, req *pb.UpdateClusterRequest) (*pb.Operation, error) {
	reqName := req.GetName()

	name, err := s.parseClusterName(reqName)
	if err != nil {
		return nil, err
	}

	fqn := name.String()
	obj := &pb.Cluster{}
	if err := s.storage.Get(ctx, fqn, obj); err != nil {
		return nil, err
	}

	klog.Infof("UpdateCluster %v", prototext.Format(req))

	update := proto.Clone(req.GetUpdate()).(*pb.ClusterUpdate)

	if update.DesiredMonitoringService != "" {
		obj.MonitoringService = update.DesiredMonitoringService
		update.DesiredMonitoringService = ""
	}
	if update.DesiredLoggingService != "" {
		obj.LoggingService = update.DesiredLoggingService
		update.DesiredLoggingService = ""
	}

	if update.DesiredNodePoolAutoscaling != nil {
		nodePoolID := update.GetDesiredNodePoolId()
		if nodePoolID == "" {
			return nil, status.Errorf(codes.InvalidArgument, "desiredNodePoolId must be specified")
		}

		nodePoolName := name.NodePool(nodePoolID)

		nodePool := &pb.NodePool{}
		if err := s.storage.Get(ctx, nodePoolName.String(), nodePool); err != nil {
			return nil, err
		}

		nodePool.Autoscaling = update.DesiredNodePoolAutoscaling
		update.DesiredNodePoolAutoscaling = nil

		if err := s.storage.Update(ctx, nodePoolName.String(), nodePool); err != nil {
			return nil, err
		}

		update.DesiredNodePoolAutoscaling = nil
		update.DesiredNodePoolId = ""
	}

	// TODO: Support more updates!

	if !proto.Equal(update, &pb.ClusterUpdate{}) {
		return nil, status.Errorf(codes.InvalidArgument, "update was not fully implemented ClusterUpdate=%v", prototext.Format(update))
	}

	if err := s.storage.Update(ctx, fqn, obj); err != nil {
		return nil, err
	}

	return s.newLRO(ctx, name.Project.ID)
}

func (s *ClusterManagerV1) DeleteCluster(ctx context.Context, req *pb.DeleteClusterRequest) (*pb.Operation, error) {
	name, err := s.parseClusterName(req.Name)
	if err != nil {
		return nil, err
	}

	fqn := name.String()

	oldObj := &pb.Cluster{}
	if err := s.storage.Delete(ctx, fqn, oldObj); err != nil {
		return nil, err
	}

	return s.newLRO(ctx, name.Project.ID)
}

type clusterName struct {
	Project  *projects.ProjectData
	Location string
	Cluster  string
}

func (n *clusterName) String() string {
	return "projects/" + n.Project.ID + "/locations/" + n.Location + "/clusters/" + n.Cluster
}

func (n *clusterName) NodePool(nodePool string) *nodePoolName {
	return &nodePoolName{
		Project:  n.Project,
		Location: n.Location,
		Cluster:  n.Cluster,
		NodePool: nodePool,
	}
}

// parseClusterName parses a string into a clusterName.
// The expected form is `projects/*/locations/*/clusters/*`.
func (s *MockService) parseClusterName(name string) (*clusterName, error) {
	tokens := strings.Split(name, "/")

	if len(tokens) == 6 && tokens[0] == "projects" && tokens[2] == "locations" && tokens[4] == "clusters" {
		project, err := s.projects.GetProjectByID(tokens[1])
		if err != nil {
			return nil, err
		}

		name := &clusterName{
			Project:  project,
			Location: tokens[3],
			Cluster:  tokens[5],
		}

		return name, nil
	} else {
		return nil, status.Errorf(codes.InvalidArgument, "name %q is not valid", name)
	}
}
