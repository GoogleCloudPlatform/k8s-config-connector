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
	"time"

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
	reqName := req.GetParent() + "/clusters/" + req.GetCluster().GetName()
	name, err := s.parseClusterName(reqName)
	if err != nil {
		return nil, err
	}

	fqn := name.String()

	obj := proto.Clone(req.Cluster).(*pb.Cluster)

	obj.Status = pb.Cluster_RUNNING

	now := time.Now().UTC()
	obj.CreateTime = now.Format(time.RFC3339Nano)

	region := name.Location

	obj.Location = name.Location

	if len(obj.Locations) == 0 {
		// We probably need to expand this to zones, but we can wait for a test
		obj.Locations = []string{name.Location}
	}

	obj.SelfLink = fmt.Sprintf("https://container.googleapis.com/v1beta1/projects/%s/locations/%s/clusters/%s", name.Project.ID, name.Location, name.Cluster)

	if obj.NetworkConfig == nil {
		obj.NetworkConfig = &pb.NetworkConfig{}
	}
	if obj.NetworkConfig.Network == "" {
		obj.NetworkConfig.Network = obj.Network
		if obj.NetworkConfig.Network == "" {
			obj.NetworkConfig.Network = "default"
		}
	}
	if obj.NetworkConfig.Subnetwork == "" {
		obj.NetworkConfig.Subnetwork = obj.Subnetwork
		if obj.NetworkConfig.Subnetwork == "" {
			obj.NetworkConfig.Subnetwork = fmt.Sprintf("projects/%s/regions/%s/subnetworks/%s", name.Project.ID, region, "default")
		}
	}

	if err := s.populateClusterDefaults(obj); err != nil {
		return nil, err
	}

	if len(obj.NodePools) != 0 {
		return nil, fmt.Errorf("nodePools must be empty when creating a cluster")
	}
	defaultNodePool := &pb.NodePool{
		Name:      "default-pool",
		Status:    pb.NodePool_RUNNING,
		Locations: []string{name.Location},
	}

	obj.NodePools = append(obj.NodePools, defaultNodePool)

	for i, nodePool := range obj.NodePools {
		nodePoolObj := proto.Clone(nodePool).(*pb.NodePool)
		if err := s.populateNodePoolDefaults(nodePoolObj); err != nil {
			return nil, err
		}
		obj.NodePools[i] = nodePoolObj
	}

	if err := s.storage.Create(ctx, fqn, obj); err != nil {
		return nil, err
	}

	for _, nodePool := range obj.NodePools {
		nodePoolFqn := name.String() + "/nodePools/" + nodePool.GetName()
		if err := s.storage.Create(ctx, nodePoolFqn, nodePool); err != nil {
			return nil, err
		}
	}

	op := &pb.Operation{
		Zone:          name.Location,
		OperationType: pb.Operation_CREATE_CLUSTER,
		TargetLink:    obj.SelfLink,
	}
	return s.startLRO(ctx, name.Project, op, func() (proto.Message, error) {
		return obj, nil
	})
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

	// We clear each field of the update as we go, so we know if we've missed one!

	if update.DesiredClusterAutoscaling != nil {
		obj.Autoscaling = update.DesiredClusterAutoscaling
		update.DesiredClusterAutoscaling = nil
	}

	if update.DesiredLoggingService != "" {
		obj.LoggingService = update.DesiredLoggingService
		update.DesiredLoggingService = ""
	}

	if update.DesiredMonitoringService != "" {
		obj.MonitoringService = update.DesiredMonitoringService
		update.DesiredMonitoringService = ""
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

	op := &pb.Operation{
		Zone:          name.Location,
		OperationType: pb.Operation_UPDATE_CLUSTER,
		TargetLink:    obj.SelfLink,
	}
	return s.startLRO(ctx, name.Project, op, func() (proto.Message, error) {
		return obj, nil
	})
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

	op := &pb.Operation{
		Zone:          name.Location,
		OperationType: pb.Operation_DELETE_CLUSTER,
		TargetLink:    oldObj.SelfLink,
	}
	return s.startLRO(ctx, name.Project, op, func() (proto.Message, error) {
		return oldObj, nil
	})
}

func (s *ClusterManagerV1) populateClusterDefaults(obj *pb.Cluster) error {
	if obj.AddonsConfig == nil {
		obj.AddonsConfig = &pb.AddonsConfig{}
	}
	if obj.AddonsConfig.GcePersistentDiskCsiDriverConfig == nil {
		obj.AddonsConfig.GcePersistentDiskCsiDriverConfig = &pb.GcePersistentDiskCsiDriverConfig{
			Enabled: true,
		}
	}
	if obj.AddonsConfig.KubernetesDashboard == nil {
		obj.AddonsConfig.KubernetesDashboard = &pb.KubernetesDashboard{
			Disabled: true,
		}
	}
	if obj.AddonsConfig.NetworkPolicyConfig == nil {
		obj.AddonsConfig.NetworkPolicyConfig = &pb.NetworkPolicyConfig{
			Disabled: true,
		}
	}

	if obj.Autoscaling == nil {
		obj.Autoscaling = &pb.ClusterAutoscaling{}
	}

	if obj.Autoscaling.AutoscalingProfile == pb.ClusterAutoscaling_PROFILE_UNSPECIFIED {
		obj.Autoscaling.AutoscalingProfile = pb.ClusterAutoscaling_BALANCED
	}

	if obj.ClusterTelemetry == nil {
		obj.ClusterTelemetry = &pb.ClusterTelemetry{}
	}

	if obj.ClusterTelemetry.Type == pb.ClusterTelemetry_UNSPECIFIED {
		obj.ClusterTelemetry.Type = pb.ClusterTelemetry_ENABLED
	}

	if obj.DatabaseEncryption == nil {
		obj.DatabaseEncryption = &pb.DatabaseEncryption{}
	}

	if obj.DatabaseEncryption.State == pb.DatabaseEncryption_UNKNOWN {
		obj.DatabaseEncryption.State = pb.DatabaseEncryption_DECRYPTED
	}

	if obj.EnterpriseConfig == nil {
		obj.EnterpriseConfig = &pb.EnterpriseConfig{}
	}

	if obj.EnterpriseConfig.ClusterTier == pb.EnterpriseConfig_CLUSTER_TIER_UNSPECIFIED {
		obj.EnterpriseConfig.ClusterTier = pb.EnterpriseConfig_STANDARD
	}

	if obj.LoggingConfig == nil {
		obj.LoggingConfig = &pb.LoggingConfig{}
	}

	if obj.LoggingConfig.ComponentConfig == nil {
		obj.LoggingConfig.ComponentConfig = &pb.LoggingComponentConfig{}
	}

	if obj.LoggingConfig.ComponentConfig.EnableComponents == nil {
		obj.LoggingConfig.ComponentConfig.EnableComponents = []pb.LoggingComponentConfig_Component{
			pb.LoggingComponentConfig_SYSTEM_COMPONENTS,
			pb.LoggingComponentConfig_WORKLOADS,
		}
	}

	if obj.MonitoringConfig == nil {
		obj.MonitoringConfig = &pb.MonitoringConfig{}
	}

	if obj.MonitoringConfig.AdvancedDatapathObservabilityConfig == nil {
		obj.MonitoringConfig.AdvancedDatapathObservabilityConfig = &pb.AdvancedDatapathObservabilityConfig{}
	}

	if obj.MonitoringConfig.ComponentConfig == nil {
		obj.MonitoringConfig.ComponentConfig = &pb.MonitoringComponentConfig{}
	}

	if obj.MonitoringConfig.ComponentConfig.EnableComponents == nil {
		obj.MonitoringConfig.ComponentConfig.EnableComponents = []pb.MonitoringComponentConfig_Component{
			pb.MonitoringComponentConfig_SYSTEM_COMPONENTS,
		}
	}

	if obj.MonitoringConfig.ManagedPrometheusConfig == nil {
		obj.MonitoringConfig.ManagedPrometheusConfig = &pb.ManagedPrometheusConfig{
			Enabled: true,
		}
	}

	if obj.ReleaseChannel == nil {
		obj.ReleaseChannel = &pb.ReleaseChannel{
			Channel: pb.ReleaseChannel_REGULAR,
		}
	}

	if obj.SecurityPostureConfig == nil {
		obj.SecurityPostureConfig = &pb.SecurityPostureConfig{}
	}

	if obj.SecurityPostureConfig.Mode == nil {
		obj.SecurityPostureConfig.Mode = PtrTo(pb.SecurityPostureConfig_BASIC)
	}

	if obj.SecurityPostureConfig.VulnerabilityMode == nil {
		obj.SecurityPostureConfig.VulnerabilityMode = PtrTo(pb.SecurityPostureConfig_VULNERABILITY_MODE_UNSPECIFIED)
	}

	return nil
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
		project, err := s.Projects.GetProjectByID(tokens[1])
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
