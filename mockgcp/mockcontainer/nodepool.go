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
	"github.com/GoogleCloudPlatform/k8s-config-connector/mockgcp/pkg/storage"
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

func (s *ClusterManagerV1) ListNodePools(ctx context.Context, req *pb.ListNodePoolsRequest) (*pb.ListNodePoolsResponse, error) {
	response := &pb.ListNodePoolsResponse{}

	prefix := req.GetParent() + "/nodePools/"
	modelKind := (&pb.NodePool{}).ProtoReflect().Descriptor()
	if err := s.storage.List(ctx, modelKind, storage.ListOptions{
		Prefix: prefix,
	}, func(obj proto.Message) error {
		model := obj.(*pb.NodePool)
		response.NodePools = append(response.NodePools, model)

		return nil
	}); err != nil {
		return nil, err
	}
	return response, nil
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

	if err := s.populateNodePoolDefaults(cluster, obj); err != nil {
		return nil, err
	}

	obj.SelfLink = name.SelfLink(ctx)
	obj.Etag = computeEtag(obj)

	if err := s.storage.Create(ctx, fqn, obj); err != nil {
		return nil, err
	}

	op := &pb.Operation{
		Zone:          name.Location,
		OperationType: pb.Operation_CREATE_NODE_POOL,
		TargetLink:    name.TargetLink(ctx),
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

	if obj.Autoscaling == nil {
		obj.Autoscaling = &pb.NodePoolAutoscaling{}
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

	if len(obj.Locations) == 0 {
		obj.Locations = cluster.Locations
	}

	if obj.InstanceGroupUrls == nil {
		zone := obj.Locations[0]
		nodePoolName := "gke-" + cluster.Name + "-hash1-" + obj.Name + "-hash2-grp"
		obj.InstanceGroupUrls = append(obj.InstanceGroupUrls,
			fmt.Sprintf("https://www.googleapis.com/compute/v1/projects/${projectId}/zones/%s/instanceGroupManagers/%s", zone, nodePoolName),
		)
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

	s.populateNetworkConfigDefaults(cluster, obj.NetworkConfig)

	if obj.PodIpv4CidrSize == 0 {
		obj.PodIpv4CidrSize = 24
	}

	if obj.UpgradeSettings == nil {
		obj.UpgradeSettings = &pb.NodePool_UpgradeSettings{
			MaxUnavailable: 0,
		}
	}
	if obj.UpgradeSettings.Strategy == nil {
		obj.UpgradeSettings.Strategy = PtrTo(pb.NodePoolUpdateStrategy_SURGE)
	}
	if obj.UpgradeSettings.MaxSurge == 0 {
		obj.UpgradeSettings.MaxSurge = 1
	}

	return nil
}

func (s *ClusterManagerV1) populateNetworkConfigDefaults(cluster *pb.Cluster, obj *pb.NodeNetworkConfig) {
	if obj.PodIpv4CidrBlock == "" {
		obj.PodIpv4CidrBlock = "10.92.0.0/14"
	}
	if obj.PodIpv4RangeUtilization == 0 {
		obj.PodIpv4RangeUtilization = 0.14
	}

	if obj.PodRange == "" {
		obj.PodRange = "gke-" + cluster.Name + "-pods-12345678"
	}

	if obj.Subnetwork == "" {
		obj.Subnetwork = cluster.NetworkConfig.Subnetwork
	}
}

func (s *ClusterManagerV1) populateNodeConfig(obj *pb.NodeConfig) error {
	if obj.BootDisk == nil {
		obj.BootDisk = &pb.BootDisk{}
	}
	if obj.BootDisk.DiskType == "" {
		obj.BootDisk.DiskType = "pd-balanced"
	}
	if obj.BootDisk.SizeGb == 0 {
		obj.BootDisk.SizeGb = 100
	}
	if obj.DiskSizeGb == 0 {
		obj.DiskSizeGb = 100
	}
	if obj.DiskType == "" {
		obj.DiskType = "pd-balanced"
	}

	if obj.EffectiveCgroupMode == pb.NodeConfig_EFFECTIVE_CGROUP_MODE_UNSPECIFIED {
		obj.EffectiveCgroupMode = pb.NodeConfig_EFFECTIVE_CGROUP_MODE_V2
	}

	if obj.ImageType == "" {
		obj.ImageType = "COS_CONTAINERD"
	}

	if obj.KubeletConfig == nil {
		obj.KubeletConfig = &pb.NodeKubeletConfig{}
	}

	if obj.KubeletConfig.InsecureKubeletReadonlyPortEnabled == nil {
		obj.KubeletConfig.InsecureKubeletReadonlyPortEnabled = PtrTo(false)
	}

	if obj.KubeletConfig.MaxParallelImagePulls == 0 {
		obj.KubeletConfig.MaxParallelImagePulls = 2
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

	if obj.ResourceLabels == nil {
		obj.ResourceLabels = make(map[string]string)
	}
	if obj.ResourceLabels["goog-gke-node-pool-provisioning-model"] == "" {
		obj.ResourceLabels["goog-gke-node-pool-provisioning-model"] = "on-demand"
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
	} else {
		// TODO: Reach out to API team to get clarification of the following behavior:
		// When the input is
		// "oauthScopes": [
		//   "https://www.googleapis.com/auth/devstorage.read_only",
		//   "https://www.googleapis.com/auth/logging.write"
		// ],
		// the output becomes
		// "oauthScopes": [
		//   "https://www.googleapis.com/auth/devstorage.read_only",
		//   "https://www.googleapis.com/auth/logging.write",
		//   "https://www.googleapis.com/auth/monitoring"
		// ],
		hasMonitoring := false
		hasLoggingWrite := false
		for _, scope := range obj.OauthScopes {
			if scope == "https://www.googleapis.com/auth/monitoring" {
				hasMonitoring = true
			}
			if scope == "https://www.googleapis.com/auth/logging.write" {
				hasLoggingWrite = true
			}
		}
		if hasLoggingWrite && !hasMonitoring {
			obj.OauthScopes = append(obj.OauthScopes, "https://www.googleapis.com/auth/monitoring")
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

	// According to the proto:
	//   This field is deprecated, min_cpu_platform should be specified using
	//   `cloud.google.com/requested-min-cpu-platform` label selector on the pod.
	//   To unset the min cpu platform field pass "automatic"
	//   as field value.
	if obj.MinCpuPlatform == "automatic" {
		obj.MinCpuPlatform = ""
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
		TargetLink: name.TargetLink(ctx),
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
		TargetLink:    name.TargetLink(ctx),
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

func (n *nodePoolName) SelfLink(ctx context.Context) string {
	version := getAPIVersion(ctx)
	prefix := "https://container.googleapis.com/" + version + "/"

	if isZone(n.Location) {
		return prefix + fmt.Sprintf("projects/%s/zones/%s/clusters/%s/nodePools/%s", n.Project.ID, n.Location, n.Cluster, n.NodePool)
	}
	return prefix + fmt.Sprintf("projects/%s/locations/%s/clusters/%s/nodePools/%s", n.Project.ID, n.Location, n.Cluster, n.NodePool)
}

func (n *nodePoolName) TargetLink(ctx context.Context) string {
	version := getAPIVersion(ctx)
	prefix := "https://container.googleapis.com/" + version + "/"

	if isZone(n.Location) {
		return prefix + fmt.Sprintf("projects/%d/zones/%s/clusters/%s/nodePools/%s", n.Project.Number, n.Location, n.Cluster, n.NodePool)
	}
	return prefix + fmt.Sprintf("projects/%d/locations/%s/clusters/%s/nodePools/%s", n.Project.Number, n.Location, n.Cluster, n.NodePool)
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
