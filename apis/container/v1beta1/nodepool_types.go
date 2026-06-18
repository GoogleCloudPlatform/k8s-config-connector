// Copyright 2026 Google LLC
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//    http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package v1beta1

import (
	computev1beta1 "github.com/GoogleCloudPlatform/k8s-config-connector/apis/compute/v1beta1"
	refsv1beta1 "github.com/GoogleCloudPlatform/k8s-config-connector/apis/refs/v1beta1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/apis/k8s/v1alpha1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

var ContainerNodePoolGVK = GroupVersion.WithKind("ContainerNodePool")

// +kcc:proto=google.container.v1.NodePoolAutoscaling
type NodePoolAutoscaling struct {
	/* Location policy used when scaling up a nodepool. */
	// +kcc:proto:field=google.container.v1.NodePoolAutoscaling.location_policy
	LocationPolicy *string `json:"locationPolicy,omitempty"`

	/* Maximum number of nodes for one location in the node pool. Must be >= min_node_count. There has to be enough quota to scale up the cluster. */
	// +kcc:proto:field=google.container.v1.NodePoolAutoscaling.max_node_count
	MaxNodeCount *int32 `json:"maxNodeCount,omitempty"`

	/* Minimum number of nodes for one location in the node pool. Must be greater than or equal to 0 and less than or equal to max_node_count. */
	// +kcc:proto:field=google.container.v1.NodePoolAutoscaling.min_node_count
	MinNodeCount *int32 `json:"minNodeCount,omitempty"`

	/* Maximum number of nodes in the node pool. Must be greater than or equal to total_min_node_count. There has to be enough quota to scale up the cluster. The total_*_node_count fields are mutually exclusive with the *_node_count fields. */
	// +kcc:proto:field=google.container.v1.NodePoolAutoscaling.total_max_node_count
	TotalMaxNodeCount *int32 `json:"totalMaxNodeCount,omitempty"`

	/* Minimum number of nodes in the node pool. Must be greater than or equal to 0 and less than or equal to total_max_node_count. The total_*_node_count fields are mutually exclusive with the *_node_count fields. */
	// +kcc:proto:field=google.container.v1.NodePoolAutoscaling.total_min_node_count
	TotalMinNodeCount *int32 `json:"totalMinNodeCount,omitempty"`
}

// +kcc:proto=google.container.v1.NodeManagement
type NodePoolManagement struct {
	/* A flag that specifies whether the node auto-repair is enabled for the node pool. If enabled, the nodes in this node pool will be monitored and, if they fail health checks too many times, an automatic repair action will be triggered. */
	// +kcc:proto:field=google.container.v1.NodeManagement.auto_repair
	AutoRepair *bool `json:"autoRepair,omitempty"`

	/* A flag that specifies whether node auto-upgrade is enabled for the node pool. If enabled, node auto-upgrade helps keep the nodes in your node pool up to date with the latest release version of Kubernetes. */
	// +kcc:proto:field=google.container.v1.NodeManagement.auto_upgrade
	AutoUpgrade *bool `json:"autoUpgrade,omitempty"`
}

// +kcc:proto=google.container.v1.AdditionalNodeNetworkConfig
type AdditionalNodeNetworkConfig struct {
	/* ComputeNetworkRef is a reference to a GCP ComputeNetwork. */
	NetworkRef *computev1beta1.ComputeNetworkRef `json:"networkRef,omitempty"`
	/* ComputeSubnetworkRef is a reference to a GCP ComputeSubnetwork. */
	SubnetworkRef *computev1beta1.ComputeSubnetworkRef `json:"subnetworkRef,omitempty"`
}

// +kcc:proto=google.container.v1.AdditionalPodNetworkConfig
type AdditionalPodNetworkConfig struct {
	/* ComputeSubnetworkRef is a reference to a GCP ComputeSubnetwork. */
	SubnetworkRef *computev1beta1.ComputeSubnetworkRef `json:"subnetworkRef,omitempty"`
	/* The name of the secondary range on the subnet which provides IP address for this pod range. */
	SecondaryPodRange *string `json:"secondaryPodRange,omitempty"`
	/* The maximum number of pods per node which use this pod network. */
	MaxPodsPerNode *int `json:"maxPodsPerNode,omitempty"`
}

// +kcc:proto=google.container.v1.NodeNetworkConfig
type NodeNetworkConfig struct {
	/* We specify the additional node networks for this node pool using this list. Each node network corresponds to an additional interface. */
	// +kcc:proto:field=google.container.v1.NodeNetworkConfig.additional_node_network_configs
	AdditionalNodeNetworkConfigs []AdditionalNodeNetworkConfig `json:"additionalNodeNetworkConfigs,omitempty"`

	/* We specify the additional pod networks for this node pool using this list. Each pod network corresponds to an additional alias IP range for the node. */
	// +kcc:proto:field=google.container.v1.NodeNetworkConfig.additional_pod_network_configs
	AdditionalPodNetworkConfigs []AdditionalPodNetworkConfig `json:"additionalPodNetworkConfigs,omitempty"`

	/* Input only. Whether to create a new range for pod IPs in this node pool. Defaults are provided for `pod_range` and `pod_ipv4_cidr_block` if they are not specified. If neither `create_pod_range` or `pod_range` are specified, the cluster-level default (`ip_allocation_policy.cluster_ipv4_cidr_block`) is used. Only applicable if `ip_allocation_policy.use_ip_aliases` is true. This field cannot be changed after the node pool has been created. */
	// +kcc:proto:field=google.container.v1.NodeNetworkConfig.create_pod_range
	CreatePodRange *bool `json:"createPodRange,omitempty"`

	/* Whether nodes have internal IP addresses only. If enable_private_nodes is not specified, then the value is derived from [Cluster.NetworkConfig.default_enable_private_nodes]. */
	// +kcc:proto:field=google.container.v1.NodeNetworkConfig.enable_private_nodes
	EnablePrivateNodes *bool `json:"enablePrivateNodes,omitempty"`

	/* [PRIVATE FIELD] Pod CIDR size overprovisioning config for the nodepool. Pod CIDR size per node depends on max_pods_per_node. By default, the value of max_pods_per_node is rounded off to next power of 2 and we then double that to get the size of pod CIDR block per node. Example: max_pods_per_node of 30 would result in 64 IPs (/26). This config can disable the doubling of IPs (we still round off to next power of 2) Example: max_pods_per_node of 30 will result in 32 IPs (/27) when overprovisioning is disabled. */
	// +kcc:proto:field=google.container.v1.NodeNetworkConfig.pod_cidr_overprovision_config
	PodCidrOverprovisionConfig *PodCIDROverprovisionConfig `json:"podCidrOverprovisionConfig,omitempty"`

	/* The IP address range for pod IPs in this node pool. Only applicable if `create_pod_range` is true. Set to blank to have a range chosen with the default size. Set to /netmask (e.g. `/14`) to have a range chosen with a specific netmask. Set to a CIDR notation (e.g. `10.96.0.0/14`) to pick a specific range to use. Only applicable if `ip_allocation_policy.use_ip_aliases` is true. This field cannot be changed after the node pool has been created. */
	// +kcc:proto:field=google.container.v1.NodeNetworkConfig.pod_ipv4_cidr_block
	PodIpv4CidrBlock *string `json:"podIpv4CidrBlock,omitempty"`

	/* The ID of the secondary range for pod IPs. If `create_pod_range` is true, this ID is used for the new range. If `create_pod_range` is false, uses an existing secondary range with this ID. Only applicable if `ip_allocation_policy.use_ip_aliases` is true. This field cannot be changed after the node pool has been created. */
	// +kcc:proto:field=google.container.v1.NodeNetworkConfig.pod_range
	PodRange *string `json:"podRange,omitempty"`

	/* ComputeSubnetworkRef is a reference to a GCP ComputeSubnetwork. */
	SubnetworkRef *computev1beta1.ComputeSubnetworkRef `json:"subnetworkRef,omitempty"`
}

// +kcc:proto=google.container.v1.NodePool.PlacementPolicy
type NodePool_PlacementPolicy struct {
	/* The type of placement. */
	// +kcc:proto:field=google.container.v1.NodePool.PlacementPolicy.type
	Type string `json:"type"`

	/* Optional. TPU placement topology for pod slice node pool. */
	// +kcc:proto:field=google.container.v1.NodePool.PlacementPolicy.tpu_topology
	TpuTopology *string `json:"tpuTopology,omitempty"`

	/* If set, refers to the name of a custom resource policy supplied by the user. The resource policy must be in the same project and region as the node pool. If not found, InvalidArgument error is returned. */
	PolicyNameRef *computev1beta1.ComputeResourcePolicyRef `json:"policyNameRef,omitempty"`
}

// +kcc:proto=google.container.v1.NodePool.UpdateConfig.BlueGreenSettings
type NodePoolBlueGreenSettings struct {
	/* Time needed after draining entire blue pool. After this period, blue pool will be cleaned up. */
	// +kcc:proto:field=google.container.v1.NodePool.UpdateConfig.NodePool_UpdateConfig_BlueGreenSettings.node_pool_soak_duration
	NodePoolSoakDuration *string `json:"nodePoolSoakDuration,omitempty"`

	// +required
	/* Standard policy for the blue-green upgrade. */
	// +kcc:proto:field=google.container.v1.NodePool.UpdateConfig.NodePool_UpdateConfig_BlueGreenSettings.standard_rollout_policy
	StandardRolloutPolicy *StandardRolloutPolicy `json:"standardRolloutPolicy"`
}

// +kcc:proto=google.container.v1.NodePool.UpgradeSettings
type NodePoolUpgradeSettings struct {
	/* Settings for blue-green upgrade strategy. */
	// +kcc:proto:field=google.container.v1.NodePool.NodePool_UpgradeSettings.blue_green_settings
	BlueGreenSettings *NodePoolBlueGreenSettings `json:"blueGreenSettings,omitempty"`

	/* The maximum number of nodes that can be created beyond the current size of the node pool during the upgrade process. */
	// +kcc:proto:field=google.container.v1.NodePool.NodePool_UpgradeSettings.max_surge
	MaxSurge *int32 `json:"maxSurge,omitempty"`

	/* The maximum number of nodes that can be simultaneously unavailable during the upgrade process. A node is considered available if its status is Ready. */
	// +kcc:proto:field=google.container.v1.NodePool.NodePool_UpgradeSettings.max_unavailable
	MaxUnavailable *int32 `json:"maxUnavailable,omitempty"`

	/* Update strategy of the node pool. */
	// +kcc:proto:field=google.container.v1.NodePool.NodePool_UpgradeSettings.strategy
	Strategy *string `json:"strategy,omitempty"`
}

// +kcc:proto=google.container.v1.WindowsNodeConfig
type WindowsNodeConfig struct {
	/* Operating system version of the Windows nodes. */
	// +kcc:proto:field=google.container.v1.WindowsNodeConfig.os_version
	OSVersion *string `json:"osVersion,omitempty"`
}

// +kcc:proto=google.container.v1.NodeConfig
type NodePoolNodeConfig struct {
	/* Immutable. Specifies options for controlling advanced machine features. */
	// +kcc:proto:field=google.container.v1.NodeConfig.advanced_machine_features
	AdvancedMachineFeatures *NodeConfig_AdvancedMachineFeatures `json:"advancedMachineFeatures,omitempty"`

	/* Immutable. Cryptographic key used to encrypt the boot disk. */
	// +kcc:proto:field=google.container.v1.NodeConfig.boot_disk_kms_key
	BootDiskKMSCryptoKeyRef *refsv1beta1.KMSCryptoKeyRef `json:"bootDiskKMSCryptoKeyRef,omitempty"`

	/* Immutable. Configuration for the confidential nodes feature, which makes nodes run on confidential VMs. Warning: This configuration can't be changed (or added/removed) after pool creation without deleting and recreating the entire pool. */
	// +kcc:proto:field=google.container.v1.NodeConfig.confidential_nodes
	ConfidentialNodes *ConfidentialNodes `json:"confidentialNodes,omitempty"`

	/* Immutable. Size of the disk attached to each node, specified in GB. The smallest allowed disk size is 10GB. */
	// +kcc:proto:field=google.container.v1.NodeConfig.disk_size_gb
	DiskSizeGb *int32 `json:"diskSizeGb,omitempty"`

	/* Immutable. Type of the disk attached to each node. Such as pd-standard, pd-balanced or pd-ssd. */
	// +kcc:proto:field=google.container.v1.NodeConfig.disk_type
	DiskType *string `json:"diskType,omitempty"`

	/* Immutable. Parameters for the ephemeral storage filesystem. If unspecified, ephemeral storage is backed by the boot disk. */
	// +kcc:proto:field=google.container.v1.NodeConfig.ephemeral_storage_config
	EphemeralStorageConfig *EphemeralStorageConfig `json:"ephemeralStorageConfig,omitempty"`

	/* Immutable. Parameters for the ephemeral storage filesystem. If unspecified, ephemeral storage is backed by the boot disk. */
	// +kcc:proto:field=google.container.v1.NodeConfig.ephemeral_storage_local_ssd_config
	EphemeralStorageLocalSsdConfig *EphemeralStorageLocalSsdConfig `json:"ephemeralStorageLocalSsdConfig,omitempty"`

	/* Enable or disable NCCL Fast Socket in the node pool. */
	// +kcc:proto:field=google.container.v1.NodeConfig.fast_socket
	FastSocket *FastSocket `json:"fastSocket,omitempty"`

	/* Immutable. GCFS configuration for this node. */
	// +kcc:proto:field=google.container.v1.NodeConfig.gcfs_config
	GcfsConfig *GcfsConfig `json:"gcfsConfig,omitempty"`

	/* Immutable. List of the type and count of accelerator cards attached to the instance. */
	// +kcc:proto:field=google.container.v1.NodeConfig.accelerators
	AcceleratorConfig []AcceleratorConfig `json:"guestAccelerator,omitempty"`

	/* Immutable. Enable or disable gvnic in the node pool. */
	// +kcc:proto:field=google.container.v1.NodeConfig.gvnic
	VirtualNic *VirtualNic `json:"gvnic,omitempty"`

	/* Immutable. The maintenance policy for the hosts on which the GKE VMs run on. */
	// +kcc:proto:field=google.container.v1.NodeConfig.host_maintenance_policy
	HostMaintenancePolicy *HostMaintenancePolicy `json:"hostMaintenancePolicy,omitempty"`

	/* The image type to use for this node. Note that for a given image type, the latest version of it will be used. */
	// +kcc:proto:field=google.container.v1.NodeConfig.image_type
	ImageType *string `json:"imageType,omitempty"`

	/* Node kubelet configs. */
	// +kcc:proto:field=google.container.v1.NodeConfig.kubelet_config
	KubeletConfig *KubeletConfig `json:"kubeletConfig,omitempty"`

	/* Immutable. The map of Kubernetes labels (key/value pairs) to be applied to each node. These will added in addition to any default label(s) that Kubernetes may apply to the node. */
	// +kcc:proto:field=google.container.v1.NodeConfig.labels
	Labels map[string]string `json:"labels,omitempty"`

	/* Parameters that can be configured on Linux nodes. */
	// +kcc:proto:field=google.container.v1.NodeConfig.linux_node_config
	LinuxNodeConfig *LinuxNodeConfig `json:"linuxNodeConfig,omitempty"`

	/* Immutable. Parameters for raw-block local NVMe SSDs. */
	// +kcc:proto:field=google.container.v1.NodeConfig.local_nvme_ssd_block_config
	NodeConfig_LocalNvmeSsdBlockConfig *NodeConfig_LocalNvmeSsdBlockConfig `json:"localNvmeSsdBlockConfig,omitempty"`

	/* Immutable. The number of local SSD disks to be attached to the node. */
	// +kcc:proto:field=google.container.v1.NodeConfig.local_ssd_count
	LocalSsdCount *int32 `json:"localSsdCount,omitempty"`

	/* Type of logging agent that is used as the default value for node pools in the cluster. Valid values include DEFAULT and MAX_THROUGHPUT. */
	// +kcc:proto:field=google.container.v1.NodeConfig.logging_config
	LoggingVariant *string `json:"loggingVariant,omitempty"`

	/* Immutable. The name of a Google Compute Engine machine type. */
	// +kcc:proto:field=google.container.v1.NodeConfig.machine_type
	MachineType *string `json:"machineType,omitempty"`

	/* Immutable. The metadata key/value pairs assigned to instances in the cluster. */
	// +kcc:proto:field=google.container.v1.NodeConfig.metadata
	Metadata map[string]string `json:"metadata,omitempty"`

	/* Immutable. Minimum CPU platform to be used by this instance. The instance may be scheduled on the specified or newer CPU platform. */
	// +kcc:proto:field=google.container.v1.NodeConfig.min_cpu_platform
	MinCPUPlatform *string `json:"minCpuPlatform,omitempty"`

	/* Immutable. Setting this field will assign instances
	of this pool to run on the specified node group. This is useful
	for running workloads on sole tenant nodes. */
	// +kcc:proto:field=google.container.v1.NodeConfig.node_group
	NodeGroupRef *computev1beta1.ComputeNodeGroupRef `json:"nodeGroupRef,omitempty"`

	/* Immutable. The set of Google API scopes to be made available on all of the node VMs. */
	// +kcc:proto:field=google.container.v1.NodeConfig.oauth_scopes
	OauthScopes []string `json:"oauthScopes,omitempty"`

	/* Immutable. Whether the nodes are created as preemptible VM instances. */
	// +kcc:proto:field=google.container.v1.NodeConfig.preemptible
	Preemptible *bool `json:"preemptible,omitempty"`

	/* Immutable. The reservation affinity configuration for the node pool. */
	// +kcc:proto:field=google.container.v1.NodeConfig.reservation_affinity
	ReservationAffinity *ReservationAffinity `json:"reservationAffinity,omitempty"`

	/* The GCE resource labels (a map of key/value pairs) to be applied to the node pool. */
	// +kcc:proto:field=google.container.v1.NodeConfig.resource_labels
	ResourceLabels map[string]string `json:"resourceLabels,omitempty"`

	/* Immutable. Sandbox configuration for this node. */
	// +kcc:proto:field=google.container.v1.NodeConfig.sandbox_config
	SandboxConfig *SandboxConfig `json:"sandboxConfig,omitempty"`

	// +kcc:proto:field=google.container.v1.NodeConfig.service_account
	ServiceAccountRef *refsv1beta1.IAMServiceAccountRef `json:"serviceAccountRef,omitempty"`

	/* Immutable. Shielded Instance options. */
	// +kcc:proto:field=google.container.v1.NodeConfig.shielded_instance_config
	ShieldedInstanceConfig *ShieldedInstanceConfig `json:"shieldedInstanceConfig,omitempty"`

	/* Immutable. Node affinity options for sole tenant node pools. */
	// +kcc:proto:field=google.container.v1.NodeConfig.sole_tenant_config
	NodeConfig_SoleTenantConfig *NodeConfig_SoleTenantConfig `json:"soleTenantConfig,omitempty"`

	/* Immutable. Whether the nodes are created as spot VM instances. */
	// +kcc:proto:field=google.container.v1.NodeConfig.spot
	Spot *bool `json:"spot,omitempty"`

	/* The list of instance tags applied to all nodes. */
	// +kcc:proto:field=google.container.v1.NodeConfig.tags
	Tags []string `json:"tags,omitempty"`

	/* List of Kubernetes taints to be applied to each node. */
	// +kcc:proto:field=google.container.v1.NodeConfig.taints
	Taint []NodeTaint `json:"taint,omitempty"`

	/* Parameters that can be configured on Windows nodes. */
	// +kcc:proto:field=google.container.v1.NodeConfig.windows_node_config
	WindowsNodeConfig *WindowsNodeConfig `json:"windowsNodeConfig,omitempty"`

	/* The workload metadata configuration for this node. */
	// +kcc:proto:field=google.container.v1.NodeConfig.workload_metadata_config
	NodeConfig_WorkloadMetadataConfig *NodeConfig_WorkloadMetadataConfig `json:"workloadMetadataConfig,omitempty"`
}

// ContainerNodePoolSpec defines the desired state of ContainerNodePool
// +kcc:spec:proto=google.container.v1.NodePool
type ContainerNodePoolSpec struct {
	/* Autoscaler configuration for this NodePool. Autoscaler is enabled only if a valid configuration is present. */
	// +kcc:proto:field=google.container.v1.NodePool.autoscaling
	Autoscaling *NodePoolAutoscaling `json:"autoscaling,omitempty"`

	/* The GKE cluster this node pool belongs to. */
	ClusterRef ContainerClusterRef `json:"clusterRef"`

	/* The initial node count for the pool. You must ensure that your Compute Engine resource quota is sufficient for this number of instances. You must also have available firewall and routes quota. */
	// +kcc:proto:field=google.container.v1.NodePool.initial_node_count
	InitialNodeCount *int32 `json:"initialNodeCount,omitempty"`

	/* Immutable. The location (region or zone) of the cluster. */
	Location string `json:"location"`

	/* NodeManagement configuration for this NodePool. */
	// +kcc:proto:field=google.container.v1.NodePool.management
	Management *NodePoolManagement `json:"management,omitempty"`

	/* The constraint on the maximum number of pods that can be run simultaneously on a node in the node pool. */
	// +kcc:proto:field=google.container.v1.NodePool.max_pods_constraint.max_pods_per_node
	MaxPodsPerNode *int `json:"maxPodsPerNode,omitempty"`

	/* Immutable. Creates a unique name for the node pool beginning with the specified prefix. Conflicts with name. */
	NamePrefix *string `json:"namePrefix,omitempty"`

	/* Networking configuration for this NodePool. If specified, it overrides the cluster-level defaults. */
	// +kcc:proto:field=google.container.v1.NodePool.network_config
	NetworkConfig *NodeNetworkConfig `json:"networkConfig,omitempty"`

	/* The node configuration of the pool. */
	// +kcc:proto:field=google.container.v1.NodePool.config
	NodeConfig *NodePoolNodeConfig `json:"nodeConfig,omitempty"`

	/* The node count of the pool. */
	NodeCount *int32 `json:"nodeCount,omitempty"`

	/* The list of Google Compute Engine zones in which the NodePool's nodes should be located. */
	// +kcc:proto:field=google.container.v1.NodePool.locations
	NodeLocations []string `json:"nodeLocations,omitempty"`

	/* Specifies the node placement policy. */
	// +kcc:proto:field=google.container.v1.NodePool.placement_policy
	PlacementPolicy *NodePool_PlacementPolicy `json:"placementPolicy,omitempty"`

	/* Immutable. Optional. The name of the resource. Used for creation and acquisition. When unset, the value of `metadata.name` is used as the default. */
	ResourceID *string `json:"resourceID,omitempty"`

	/* Upgrade settings control disruption and speed of the upgrade. */
	// +kcc:proto:field=google.container.v1.NodePool.upgrade_settings
	UpgradeSettings *NodePoolUpgradeSettings `json:"upgradeSettings,omitempty"`

	/* The version of Kubernetes running on this NodePool's nodes. */
	// +kcc:proto:field=google.container.v1.NodePool.version
	Version *string `json:"version,omitempty"`
}

// ContainerNodePoolStatus defines the config connector machine state of ContainerNodePool
type ContainerNodePoolStatus struct {
	/* Conditions represent the latest available observations of the
	   ContainerNodePool's current state. */
	Conditions []v1alpha1.Condition `json:"conditions,omitempty"`

	/* The resource URLs of the managed instance groups associated with this node pool. */
	InstanceGroupUrls []string `json:"instanceGroupUrls,omitempty"`

	/* List of instance group URLs which have been assigned to this node pool. */
	ManagedInstanceGroupUrls []string `json:"managedInstanceGroupUrls,omitempty"`

	/* ObservedGeneration is the generation of the resource that was most recently observed by the Config Connector controller. If this is equal to metadata.generation, then that means that the current reported status reflects the most recent desired state of the resource. */
	ObservedGeneration *int64 `json:"observedGeneration,omitempty"`

	/* The observed state of the underlying GCP resource. */
	ObservedState *NodepoolObservedStateStatus `json:"observedState,omitempty"`

	Operation *string `json:"operation,omitempty"`
}

type NodepoolObservedStateStatus struct {
	Version *string `json:"version,omitempty"`
}

// +genclient
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:resource:categories=gcp,shortName=gcpcontainernodepool;gcpcontainernodepools
// +kubebuilder:subresource:status
// +kubebuilder:metadata:labels="cnrm.cloud.google.com/managed-by-kcc=true"
// +kubebuilder:metadata:labels="cnrm.cloud.google.com/stability-level=stable"
// +kubebuilder:metadata:labels="cnrm.cloud.google.com/system=true"
// +kubebuilder:metadata:labels="cnrm.cloud.google.com/tf2crd=true"
// +kubebuilder:printcolumn:name="Age",JSONPath=".metadata.creationTimestamp",type="date"
// +kubebuilder:printcolumn:name="Ready",JSONPath=".status.conditions[?(@.type=='Ready')].status",type="string",description="When 'True', the most recent reconcile of the resource succeeded"
// +kubebuilder:printcolumn:name="Status",JSONPath=".status.conditions[?(@.type=='Ready')].reason",type="string",description="The reason for the value in 'Ready'"
// +kubebuilder:printcolumn:name="Status Age",JSONPath=".status.conditions[?(@.type=='Ready')].lastTransitionTime",type="date",description="The last transition time for the value in 'Status'"

// ContainerNodePool is the Schema for the ContainerNodePool API
// +k8s:openapi-gen=true
type ContainerNodePool struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	// +required
	Spec   ContainerNodePoolSpec   `json:"spec,omitempty"`
	Status ContainerNodePoolStatus `json:"status,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// ContainerNodePoolList contains a list of ContainerNodePool
type ContainerNodePoolList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []ContainerNodePool `json:"items"`
}

func init() {
	SchemeBuilder.Register(&ContainerNodePool{}, &ContainerNodePoolList{})
}
