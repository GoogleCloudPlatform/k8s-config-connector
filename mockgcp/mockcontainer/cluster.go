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
		if status.Code(err) == codes.NotFound {
			return nil, status.Errorf(codes.NotFound, "Not found: %s.", AsZonalLink(fqn))
		}
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

	region, err := locationToRegion(name.Location)
	if err != nil {
		return nil, err
	}

	obj.Location = name.Location

	if len(obj.Locations) == 0 {
		// We probably need to expand this to zones, but we can wait for a test
		obj.Locations = []string{name.Location}
	}

	obj.SelfLink = buildSelfLink(ctx, fmt.Sprintf("projects/%s/locations/%s/clusters/%s", name.Project.ID, name.Location, name.Cluster))
	obj.SelfLink = AsZonalLink(obj.SelfLink)

	if obj.NetworkConfig == nil {
		obj.NetworkConfig = &pb.NetworkConfig{}
	}
	if obj.Network == "" && obj.NetworkConfig.Network == "" {
		obj.Network = "default"
		obj.NetworkConfig.Network = "default"
	} else if obj.Network != "" && obj.NetworkConfig.Network == "" {
		obj.NetworkConfig.Network = obj.Network
	}
	if obj.Subnetwork == "" && obj.NetworkConfig.Subnetwork == "" {
		obj.Subnetwork = "default"
		obj.NetworkConfig.Subnetwork = fmt.Sprintf("projects/%s/regions/%s/subnetworks/%s", name.Project.ID, region, obj.Subnetwork)
	} else if obj.Subnetwork != "" && obj.NetworkConfig.Subnetwork == "" {
		if strings.Contains(obj.Subnetwork, "projects/") {
			obj.NetworkConfig.Subnetwork = obj.Subnetwork
		} else {
			obj.NetworkConfig.Subnetwork = fmt.Sprintf("projects/%s/regions/%s/subnetworks/%s", name.Project.ID, region, obj.Subnetwork)
		}
	}

	if obj.NetworkConfig.DefaultSnatStatus == nil {
		obj.NetworkConfig.DefaultSnatStatus = &pb.DefaultSnatStatus{}
	}

	if obj.NetworkConfig.ServiceExternalIpsConfig == nil {
		obj.NetworkConfig.ServiceExternalIpsConfig = &pb.ServiceExternalIPsConfig{}
	}

	// On output, Network and Subnetwork show the ID instead of the full name
	obj.Network = lastComponent(obj.Network)
	obj.Subnetwork = lastComponent(obj.Subnetwork)

	if isZone(obj.Location) {
		obj.Zone = obj.Location
	}

	// PrivateCluster is now the default??
	obj.PrivateCluster = true

	obj.ServicesIpv4Cidr = "34.118.224.0/20"

	if err := s.populateClusterDefaults(obj); err != nil {
		return nil, err
	}

	if len(obj.NodePools) == 0 {
		defaultNodePool := &pb.NodePool{
			Name:      "default-pool",
			Status:    pb.NodePool_RUNNING,
			Locations: []string{name.Location},
		}
		obj.NodePools = append(obj.NodePools, defaultNodePool)
	}

	for i, nodePool := range obj.NodePools {
		nodePoolObj := proto.Clone(nodePool).(*pb.NodePool)
		if err := s.populateNodePoolDefaults(obj, nodePoolObj); err != nil {
			return nil, err
		}
		obj.NodePools[i] = nodePoolObj
	}

	if obj.Autoscaling != nil && obj.Autoscaling.AutoprovisioningNodePoolDefaults != nil &&
		obj.Autoscaling.AutoprovisioningNodePoolDefaults.UpgradeSettings != nil {
		upgradeSettings := obj.Autoscaling.AutoprovisioningNodePoolDefaults.UpgradeSettings
		if *upgradeSettings.Strategy == pb.NodePoolUpdateStrategy_SURGE && upgradeSettings.MaxSurge > 0 {
			obj.Autoscaling.AutoprovisioningNodePoolDefaults.UpgradeSettings.BlueGreenSettings = nil
		}
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
		TargetLink:    buildTargetLink(ctx, name),
	}
	return s.startLRO(ctx, name.Project, op, func() (proto.Message, error) {
		op.Progress = &pb.OperationProgress{
			Metrics: []*pb.OperationProgress_Metric{
				{Name: "CLUSTER_CONFIGURING", Value: &pb.OperationProgress_Metric_IntValue{IntValue: 10}},
				{Name: "CLUSTER_CONFIGURING_TOTAL", Value: &pb.OperationProgress_Metric_IntValue{IntValue: 10}},
				{Name: "CLUSTER_DEPLOYING", Value: &pb.OperationProgress_Metric_IntValue{IntValue: 12}},
				{Name: "CLUSTER_DEPLOYING_TOTAL", Value: &pb.OperationProgress_Metric_IntValue{IntValue: 12}},
				{Name: "CLUSTER_HEALTHCHECKING", Value: &pb.OperationProgress_Metric_IntValue{IntValue: 1}},
				{Name: "CLUSTER_HEALTHCHECKING_TOTAL", Value: &pb.OperationProgress_Metric_IntValue{IntValue: 2}},
			},
		}
		return obj, nil
	})
}

func locationToRegion(location string) (string, error) {
	tokens := strings.Split(location, "-")
	if len(tokens) == 2 {
		return location, nil
	}
	if len(tokens) == 3 {
		return fmt.Sprintf("%s-%s", tokens[0], tokens[1]), nil
	}
	return "", fmt.Errorf("incorrect location: %v", location)
}

func locationToZone(location string) (string, error) {
	tokens := strings.Split(location, "-")
	if len(tokens) == 3 {
		return location, nil
	}
	if len(tokens) == 2 {
		return fmt.Sprintf("%s-a", location), nil
	}
	return "", fmt.Errorf("incorrect location: %v", location)
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
		if obj.Autoscaling.AutoprovisioningNodePoolDefaults != nil &&
			obj.Autoscaling.AutoprovisioningNodePoolDefaults.UpgradeSettings != nil &&
			obj.Autoscaling.AutoprovisioningNodePoolDefaults.UpgradeSettings.BlueGreenSettings != nil &&
			obj.Autoscaling.AutoprovisioningNodePoolDefaults.UpgradeSettings.BlueGreenSettings.String() == "" {
			obj.Autoscaling.AutoprovisioningNodePoolDefaults.UpgradeSettings.BlueGreenSettings = nil
		}
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

	if update.DesiredAddonsConfig != nil {
		obj.AddonsConfig = update.DesiredAddonsConfig
		update.DesiredAddonsConfig = nil
	}

	if update.DesiredNodePoolAutoConfigNetworkTags != nil {
		if obj.NodePoolAutoConfig == nil {
			obj.NodePoolAutoConfig = &pb.NodePoolAutoConfig{}
		}
		obj.NodePoolAutoConfig.NetworkTags = update.DesiredNodePoolAutoConfigNetworkTags
		update.DesiredNodePoolAutoConfigNetworkTags = nil
	}

	if update.DesiredMasterAuthorizedNetworksConfig != nil {
		obj.MasterAuthorizedNetworksConfig = update.DesiredMasterAuthorizedNetworksConfig
		update.DesiredMasterAuthorizedNetworksConfig = nil
	}

	if update.DesiredPrivateClusterConfig != nil {
		obj.PrivateClusterConfig = update.DesiredPrivateClusterConfig
		update.DesiredPrivateClusterConfig = nil
	}

	if update.DesiredEnablePrivateEndpoint != nil {
		if obj.PrivateClusterConfig == nil {
			obj.PrivateClusterConfig = &pb.PrivateClusterConfig{}
		}
		obj.PrivateClusterConfig.EnablePrivateEndpoint = *update.DesiredEnablePrivateEndpoint
		update.DesiredEnablePrivateEndpoint = nil
	}

	if update.DesiredControlPlaneEndpointsConfig != nil {
		obj.ControlPlaneEndpointsConfig = update.DesiredControlPlaneEndpointsConfig

		if update.DesiredControlPlaneEndpointsConfig.IpEndpointsConfig != nil {
			if update.DesiredControlPlaneEndpointsConfig.IpEndpointsConfig.EnablePublicEndpoint != nil {
				obj.PrivateClusterConfig.EnablePrivateEndpoint = !*update.DesiredControlPlaneEndpointsConfig.IpEndpointsConfig.EnablePublicEndpoint
			}
			if update.DesiredControlPlaneEndpointsConfig.IpEndpointsConfig.AuthorizedNetworksConfig != nil {
				obj.MasterAuthorizedNetworksConfig = update.DesiredControlPlaneEndpointsConfig.IpEndpointsConfig.AuthorizedNetworksConfig
			}
		}

		update.DesiredControlPlaneEndpointsConfig = nil

	}

	if update.DesiredDefaultEnablePrivateNodes != nil {
		obj.NetworkConfig.DefaultEnablePrivateNodes = update.DesiredDefaultEnablePrivateNodes
		update.DesiredDefaultEnablePrivateNodes = nil
	}

	if update.DesiredEnableCiliumClusterwideNetworkPolicy != nil {
		if obj.NetworkConfig == nil {
			obj.NetworkConfig = &pb.NetworkConfig{}
		}
		obj.NetworkConfig.EnableCiliumClusterwideNetworkPolicy = update.DesiredEnableCiliumClusterwideNetworkPolicy
		update.DesiredEnableCiliumClusterwideNetworkPolicy = nil
	}

	// TODO: Support more updates!

	if !proto.Equal(update, &pb.ClusterUpdate{}) {
		return nil, status.Errorf(codes.InvalidArgument, "update was not fully implemented ClusterUpdate=%v", prototext.Format(update))
	}

	if err := s.populateClusterDefaults(obj); err != nil {
		return nil, err
	}

	if err := s.storage.Update(ctx, fqn, obj); err != nil {
		return nil, err
	}

	op := &pb.Operation{
		Zone:          name.Location,
		OperationType: pb.Operation_UPDATE_CLUSTER,
		TargetLink:    buildTargetLink(ctx, name),
	}
	return s.startLRO(ctx, name.Project, op, func() (proto.Message, error) {
		return obj, nil
	})
}

func (s *ClusterManagerV1) SetLabels(ctx context.Context, req *pb.SetLabelsRequest) (*pb.Operation, error) {
	reqName := req.GetName()

	name, err := s.parseClusterName(reqName)
	if err != nil {
		return nil, err
	}

	fqn := name.String()
	existing := &pb.Cluster{}
	if err := s.storage.Get(ctx, fqn, existing); err != nil {
		return nil, err
	}

	klog.Infof("SetLabels %v", prototext.Format(req))

	if existing.GetLabelFingerprint() != req.GetLabelFingerprint() {
		return nil, status.Errorf(codes.FailedPrecondition, "label fingerprint does not match")
	}

	update := proto.Clone(existing).(*pb.Cluster)
	update.ResourceLabels = req.ResourceLabels

	if err := s.storage.Update(ctx, fqn, update); err != nil {
		return nil, err
	}

	op := &pb.Operation{
		Zone:          name.Location,
		OperationType: pb.Operation_SET_LABELS,
		TargetLink:    existing.SelfLink,
	}
	return s.startLRO(ctx, name.Project, op, func() (proto.Message, error) {
		return existing, nil
	})
}

func (s *ClusterManagerV1) SetMaintenancePolicy(ctx context.Context, req *pb.SetMaintenancePolicyRequest) (*pb.Operation, error) {
	name, err := s.parseClusterName(req.Name)
	if err != nil {
		return nil, err
	}

	fqn := name.String()
	obj := &pb.Cluster{}
	if err := s.storage.Get(ctx, fqn, obj); err != nil {
		return nil, err
	}

	obj.MaintenancePolicy = req.MaintenancePolicy
	if obj.MaintenancePolicy != nil && obj.MaintenancePolicy.GetWindow() != nil &&
		obj.MaintenancePolicy.GetWindow().GetDailyMaintenanceWindow() != nil {

		obj.MaintenancePolicy.Window.GetDailyMaintenanceWindow().Duration = "PT4H0M0S"
	}

	if err := s.storage.Update(ctx, fqn, obj); err != nil {
		return nil, err
	}

	op := &pb.Operation{
		Zone:          name.Location,
		OperationType: pb.Operation_UPDATE_CLUSTER,
		TargetLink:    buildTargetLink(ctx, name),
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
		TargetLink:    buildTargetLink(ctx, name),
	}
	return s.startLRO(ctx, name.Project, op, func() (proto.Message, error) {
		return oldObj, nil
	})
}

func (s *ClusterManagerV1) populateClusterDefaults(obj *pb.Cluster) error {
	if obj.NodeConfig == nil {
		obj.NodeConfig = &pb.NodeConfig{}
	}
	if err := s.populateNodeConfig(obj.NodeConfig); err != nil {
		return err
	}

	if obj.InitialClusterVersion == "" {
		obj.InitialClusterVersion = "1.30.5-gke.1014001"
	}

	if obj.AddonsConfig == nil {
		obj.AddonsConfig = &pb.AddonsConfig{}
	}
	if obj.AddonsConfig.CloudRunConfig != nil &&
		obj.AddonsConfig.CloudRunConfig.LoadBalancerType == pb.CloudRunConfig_LOAD_BALANCER_TYPE_UNSPECIFIED {
		obj.AddonsConfig.CloudRunConfig.LoadBalancerType = pb.CloudRunConfig_LOAD_BALANCER_TYPE_EXTERNAL
	}
	if obj.AddonsConfig.GcePersistentDiskCsiDriverConfig == nil {
		obj.AddonsConfig.GcePersistentDiskCsiDriverConfig = &pb.GcePersistentDiskCsiDriverConfig{}
	}
	// Weird behavior:
	// Even if `.AddonsConfig.GcePersistentDiskCsiDriverConfig.Enabled` is set to
	// `false` in request, the returned value is still `true`.
	obj.AddonsConfig.GcePersistentDiskCsiDriverConfig.Enabled = true

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
	if obj.Autoscaling.EnableNodeAutoprovisioning == true {
		if obj.Autoscaling.AutoprovisioningNodePoolDefaults == nil {
			obj.Autoscaling.AutoprovisioningNodePoolDefaults = &pb.AutoprovisioningNodePoolDefaults{}
		}

		if err := s.populateAutoprovisioningNodePoolDefaults(obj.Autoscaling.AutoprovisioningNodePoolDefaults); err != nil {
			return err
		}
	}

	if obj.MaintenancePolicy != nil {
		if obj.MaintenancePolicy.ResourceVersion == "" {
			obj.MaintenancePolicy.ResourceVersion = "1234abcd"
		}
		if obj.MaintenancePolicy.GetWindow() != nil {
			window := obj.MaintenancePolicy.GetWindow()
			if len(window.MaintenanceExclusions) == 0 && window.Policy == nil {
				obj.MaintenancePolicy.Window = nil
			} else {
				if obj.MaintenancePolicy.Window.GetDailyMaintenanceWindow() != nil {
					obj.MaintenancePolicy.Window.GetDailyMaintenanceWindow().Duration = "PT4H0M0S"
				}
				for _, exclusion := range obj.MaintenancePolicy.Window.GetMaintenanceExclusions() {
					if exclusion.GetOptions() == nil {
						exclusion.Options = &pb.TimeWindow_MaintenanceExclusionOptions{}
					}
				}
			}
		}
	}

	if obj.BinaryAuthorization == nil {
		obj.BinaryAuthorization = &pb.BinaryAuthorization{}
	}

	if obj.ClusterIpv4Cidr == "" {
		obj.ClusterIpv4Cidr = "10.92.0.0/14"
	}

	if obj.ClusterTelemetry == nil {
		obj.ClusterTelemetry = &pb.ClusterTelemetry{}
	}

	if obj.ClusterTelemetry.Type == pb.ClusterTelemetry_UNSPECIFIED {
		obj.ClusterTelemetry.Type = pb.ClusterTelemetry_ENABLED
	}

	if obj.CurrentMasterVersion == "" {
		obj.CurrentMasterVersion = obj.InitialClusterVersion
	}
	if obj.CurrentNodeVersion == "" {
		obj.CurrentNodeVersion = obj.InitialClusterVersion
	}

	if obj.CurrentNodeCount == 0 {
		obj.CurrentNodeCount = 1
	}

	if obj.DatabaseEncryption == nil {
		obj.DatabaseEncryption = &pb.DatabaseEncryption{}
	}

	if obj.DatabaseEncryption.State == pb.DatabaseEncryption_UNKNOWN {
		obj.DatabaseEncryption.State = pb.DatabaseEncryption_DECRYPTED
	}
	if obj.DatabaseEncryption.CurrentState == nil {
		obj.DatabaseEncryption.CurrentState = PtrTo(pb.DatabaseEncryption_CURRENT_STATE_DECRYPTED)
	}

	if obj.DefaultMaxPodsConstraint == nil {
		obj.DefaultMaxPodsConstraint = &pb.MaxPodsConstraint{}
	}
	if obj.DefaultMaxPodsConstraint.MaxPodsPerNode == 0 {
		obj.DefaultMaxPodsConstraint.MaxPodsPerNode = 110
	}

	if obj.Endpoint == "" {
		obj.Endpoint = "1.23.456.78"
	}

	if obj.EnterpriseConfig == nil {
		obj.EnterpriseConfig = &pb.EnterpriseConfig{}
	}

	if obj.EnterpriseConfig.ClusterTier == pb.EnterpriseConfig_CLUSTER_TIER_UNSPECIFIED {
		obj.EnterpriseConfig.ClusterTier = pb.EnterpriseConfig_STANDARD
	}

	if obj.Etag == "" {
		obj.Etag = "abcdef0123A="
	}

	if obj.Id == "" {
		obj.Id = "000000000000000000000"
	}

	zone, err := locationToZone(obj.Location)
	if err != nil {
		return err
	}

	if obj.InstanceGroupUrls == nil {
		obj.InstanceGroupUrls = []string{
			fmt.Sprintf("https://www.googleapis.com/compute/v1/projects/${projectId}/zones/%s/instanceGroupManagers/gke-containercluster-abcdef", zone),
		}
	}

	if obj.LabelFingerprint == "" {
		obj.LabelFingerprint = "abcdef0123A="
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

	if obj.LoggingService == "" {
		obj.LoggingService = "logging.googleapis.com/kubernetes"
	}

	if obj.Master == nil {
		obj.Master = &pb.Master{}
	}

	if obj.MasterAuth == nil {
		obj.MasterAuth = &pb.MasterAuth{
			ClientCertificateConfig: &pb.ClientCertificateConfig{},
			ClusterCaCertificate:    "1234567890abcdefghijklmn",
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

	if obj.MonitoringService == "" {
		obj.MonitoringService = "monitoring.googleapis.com/kubernetes"
	}

	if obj.NetworkPolicy != nil && obj.NetworkPolicy.String() == "" {
		obj.NetworkPolicy = nil
	}

	if obj.PrivateClusterConfig == nil {
		obj.PrivateClusterConfig = &pb.PrivateClusterConfig{}
	}
	if obj.PrivateClusterConfig.PublicEndpoint == "" {
		obj.PrivateClusterConfig.PublicEndpoint = "8.8.8.8"
	}
	if obj.NetworkConfig != nil && obj.NetworkConfig.DefaultEnablePrivateNodes != nil {
		obj.PrivateClusterConfig.EnablePrivateNodes = *obj.NetworkConfig.DefaultEnablePrivateNodes
	}

	if obj.ControlPlaneEndpointsConfig == nil {
		obj.ControlPlaneEndpointsConfig = &pb.ControlPlaneEndpointsConfig{}
	}
	if obj.ControlPlaneEndpointsConfig.IpEndpointsConfig != nil {

		if obj.ControlPlaneEndpointsConfig.IpEndpointsConfig.Enabled != nil &&
			*obj.ControlPlaneEndpointsConfig.IpEndpointsConfig.Enabled {

			// The deprecated PrivateClusterConfig.PrivateEndpoint field seems to
			// be populated only when IP endpoints config is enabled.
			if obj.PrivateClusterConfig.PrivateEndpoint == "" {
				obj.PrivateClusterConfig.PrivateEndpoint = "10.128.0.2"
			}

			if obj.ControlPlaneEndpointsConfig.IpEndpointsConfig.PrivateEndpoint == "" {
				obj.ControlPlaneEndpointsConfig.IpEndpointsConfig.PrivateEndpoint = "10.128.0.2"
			}
			if obj.ControlPlaneEndpointsConfig.IpEndpointsConfig.EnablePublicEndpoint != nil &&
				*obj.ControlPlaneEndpointsConfig.IpEndpointsConfig.EnablePublicEndpoint {

				if obj.ControlPlaneEndpointsConfig.IpEndpointsConfig.PublicEndpoint == "" {
					obj.ControlPlaneEndpointsConfig.IpEndpointsConfig.PublicEndpoint = "8.8.8.8"
				}
				if obj.ControlPlaneEndpointsConfig.IpEndpointsConfig.AuthorizedNetworksConfig == nil {
					obj.ControlPlaneEndpointsConfig.IpEndpointsConfig.AuthorizedNetworksConfig = &pb.MasterAuthorizedNetworksConfig{}
				}
				obj.ControlPlaneEndpointsConfig.IpEndpointsConfig.AuthorizedNetworksConfig.GcpPublicCidrsAccessEnabled = PtrTo(true)

				if obj.MasterAuthorizedNetworksConfig == nil {
					obj.MasterAuthorizedNetworksConfig = &pb.MasterAuthorizedNetworksConfig{}
				}
				obj.MasterAuthorizedNetworksConfig.GcpPublicCidrsAccessEnabled = PtrTo(true)
			} else {
				if !obj.MasterAuthorizedNetworksConfig.Enabled {
					return fmt.Errorf("'masterAuthorizedNetworksConfig' must be enabled when private endpoint is enabled")
				} else if obj.MasterAuthorizedNetworksConfig.GcpPublicCidrsAccessEnabled != nil &&
					*obj.MasterAuthorizedNetworksConfig.GcpPublicCidrsAccessEnabled {
					return fmt.Errorf("'masterAuthorizedNetworksConfig.gcpPublicCidrsAccessEnabled' cannot be true if private endpoint is enabled")
				}
			}
		}
	}
	if obj.ControlPlaneEndpointsConfig.DnsEndpointConfig == nil {
		obj.ControlPlaneEndpointsConfig.DnsEndpointConfig = &pb.ControlPlaneEndpointsConfig_DNSEndpointConfig{}
	}
	// Always set the computed endpoint field
	obj.ControlPlaneEndpointsConfig.DnsEndpointConfig.Endpoint = fmt.Sprintf("gke-12345trewq-${projectNumber}.%s.gke.goog", obj.Location)

	if obj.ProtectConfig == nil {
		obj.ProtectConfig = &pb.ProtectConfig{}
	}
	if obj.ProtectConfig.WorkloadConfig == nil {
		obj.ProtectConfig.WorkloadConfig = &pb.WorkloadConfig{}
	}
	if obj.ProtectConfig.WorkloadConfig.AuditMode == nil {
		obj.ProtectConfig.WorkloadConfig.AuditMode = PtrTo(pb.WorkloadConfig_BASIC)
	}
	if obj.ProtectConfig.WorkloadVulnerabilityMode == nil {
		obj.ProtectConfig.WorkloadVulnerabilityMode = PtrTo(pb.ProtectConfig_WORKLOAD_VULNERABILITY_MODE_UNSPECIFIED)
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

func (n *clusterName) LinkWithNumber() string {
	return fmt.Sprintf("projects/%d/locations/%s/clusters/%s", n.Project.Number, n.Location, n.Cluster)
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

func buildTargetLink(ctx context.Context, name *clusterName) string {
	return buildSelfLink(ctx, AsZonalLink(name.LinkWithNumber()))
}

func lastComponent(s string) string {
	return s[strings.LastIndex(s, "/")+1:]
}
