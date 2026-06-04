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
	pubsubv1beta1 "github.com/GoogleCloudPlatform/k8s-config-connector/apis/pubsub/v1beta1"
	refsv1beta1 "github.com/GoogleCloudPlatform/k8s-config-connector/apis/refs/v1beta1"
	secretv1beta1 "github.com/GoogleCloudPlatform/k8s-config-connector/apis/refs/v1beta1/secret"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/apis/k8s/v1alpha1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// +kcc:proto=google.container.v1.RangeInfo
type RangeInfo struct {
	/* Output only. Name of a range. */
	// +kcc:proto:field=google.container.v1.RangeInfo.range_name
	RangeName *string `json:"rangeName,omitempty"`

	/* Output only. The utilization of the range. */
	// +kcc:proto:field=google.container.v1.RangeInfo.utilization
	Utilization *float64 `json:"utilization,omitempty"`
}

// +kcc:proto=google.container.v1.AdditionalPodRangesConfig
type AdditionalPodRangesConfig struct {
	/* Name for pod secondary ipv4 range which has the actual range defined ahead. */
	// +kcc:proto:field=google.container.v1.AdditionalPodRangesConfig.pod_range_names
	PodRangeNames []string `json:"podRangeNames"`
}

// +kcc:proto=google.container.v1.AddonsConfig
type AddonsConfig struct {
	/* The status of the CloudRun addon. It is disabled by default. Set disabled = false to enable. */
	// +kcc:proto:field=google.container.v1.AddonsConfig.cloud_run_config
	CloudRunConfig *CloudRunConfig `json:"cloudrunConfig,omitempty"`

	/* The of the Config Connector addon. */
	// +kcc:proto:field=google.container.v1.AddonsConfig.config_connector_config
	ConfigConnectorConfig *ConfigConnectorConfig `json:"configConnectorConfig,omitempty"`

	/* The status of the NodeLocal DNSCache addon. It is disabled by default. Set enabled = true to enable. */
	// +kcc:proto:field=google.container.v1.AddonsConfig.dns_cache_config
	DNSCacheConfig *DNSCacheConfig `json:"dnsCacheConfig,omitempty"`

	/* Whether this cluster should enable the Google Compute Engine Persistent Disk Container Storage Interface (CSI) Driver. Set enabled = true to enable. The Compute Engine persistent disk CSI Driver is enabled by default on newly created clusters for the following versions: Linux clusters: GKE version 1.18.10-gke.2100 or later, or 1.19.3-gke.2100 or later. */
	// +kcc:proto:field=google.container.v1.AddonsConfig.gce_persistent_disk_csi_driver_config
	GCEPersistentDiskCSIDriverConfig *GCEPersistentDiskCSIDriverConfig `json:"gcePersistentDiskCsiDriverConfig,omitempty"`

	/* The status of the Filestore CSI driver addon, which allows the usage of filestore instance as volumes. Defaults to disabled; set enabled = true to enable. */
	// +kcc:proto:field=google.container.v1.AddonsConfig.gcp_filestore_csi_driver_config
	GCPFilestoreCSIDriverConfig *GCPFilestoreCSIDriverConfig `json:"gcpFilestoreCsiDriverConfig,omitempty"`

	/* The status of the GCS Fuse CSI driver addon, which allows the usage of GCS bucket as volumes. Defaults to disabled; set enabled = true to enable. */
	// +kcc:proto:field=google.container.v1.AddonsConfig.gcs_fuse_csi_driver_config
	GCSFuseCSIDriverConfig *GCSFuseCSIDriverConfig `json:"gcsFuseCsiDriverConfig,omitempty"`

	/* The status of the Backup for GKE Agent addon. It is disabled by default. Set enabled = true to enable. */
	// +kcc:proto:field=google.container.v1.AddonsConfig.gke_backup_agent_config
	GKEBackupAgentConfig *GKEBackupAgentConfig `json:"gkeBackupAgentConfig,omitempty"`

	/* The status of the Horizontal Pod Autoscaling addon, which increases or decreases the number of replica pods a replication controller has based on the resource usage of the existing pods. It ensures that a Heapster pod is running in the cluster, which is also used by the Cloud Monitoring service. It is enabled by default; set disabled = true to disable. */
	// +kcc:proto:field=google.container.v1.AddonsConfig.horizontal_pod_autoscaling
	HorizontalPodAutoscaling *HorizontalPodAutoscaling `json:"horizontalPodAutoscaling,omitempty"`

	/* The status of the HTTP (L7) load balancing controller addon, which makes it easy to set up HTTP load balancers for services in a cluster. It is enabled by default; set disabled = true to disable. */
	// +kcc:proto:field=google.container.v1.AddonsConfig.http_load_balancing
	HTTPLoadBalancing *HTTPLoadBalancing `json:"httpLoadBalancing,omitempty"`

	/* The status of the Istio addon. */
	// +kcc:proto:field=google.container.v1.AddonsConfig.istio_config
	IstioConfig *IstioConfig `json:"istioConfig,omitempty"`

	/* Configuration for the KALM addon, which manages the lifecycle of k8s. It is disabled by default; Set enabled = true to enable. */
	// +kcc:proto:field=google.container.v1.AddonsConfig.kalm_config
	KalmConfig *KalmConfig `json:"kalmConfig,omitempty"`

	/* Whether we should enable the network policy addon for the master. This must be enabled in order to enable network policy for the nodes. To enable this, you must also define a network_policy block, otherwise nothing will happen. It can only be disabled if the nodes already do not have network policies enabled. Defaults to disabled; set disabled = false to enable. */
	// +kcc:proto:field=google.container.v1.AddonsConfig.network_policy_config
	NetworkPolicyConfig *NetworkPolicyConfig `json:"networkPolicyConfig,omitempty"`
}

// +kcc:proto=google.container.v1.AdvancedDatapathObservabilityConfig
type AdvancedDatapathObservabilityConfig struct {
	/* Whether or not the advanced datapath metrics are enabled. */
	// +required
	// +kcc:proto:field=google.container.v1.AdvancedDatapathObservabilityConfig.enable_metrics
	EnableMetrics *bool `json:"enableMetrics,omitempty"`

	/* Mode used to make Relay available. */
	// +kcc:proto:field=google.container.v1.AdvancedDatapathObservabilityConfig.relay_mode
	RelayMode *string `json:"relayMode,omitempty"`
}

// +kcc:proto=google.container.v1.NodeConfig.AdvancedMachineFeatures
type NodeConfig_AdvancedMachineFeatures struct {
	/* Immutable. Whether or not to enable nested virtualization (defaults to false). */
	// +kcc:proto:field=google.container.v1.NodeConfig.AdvancedMachineFeatures.enable_nested_virtualization
	EnableNestedVirtualization *bool `json:"enableNestedVirtualization,omitempty"`

	/* Immutable. The number of threads per physical core. To disable simultaneous multithreading (SMT) set this to 1. If unset, the maximum number of threads supported per core by the underlying processor is assumed. */
	// +kcc:proto:field=google.container.v1.NodeConfig.AdvancedMachineFeatures.threads_per_core
	ThreadsPerCore *int `json:"threadsPerCore,omitempty"`
}

// +kcc:proto=google.container.v1.AuthenticatorGroupsConfig
type AuthenticatorGroupsConfig struct {
	/* The name of the RBAC security group for use with Google security groups in Kubernetes RBAC. Group name must be in format gke-security-groups@yourdomain.com. */
	// +required
	// +kcc:proto:field=google.container.v1.AuthenticatorGroupsConfig.security_group
	SecurityGroup *string `json:"securityGroup,omitempty"`
}

// +kcc:proto=google.container.v1.AutoprovisioningNodePoolDefaults
type ClusterAutoscaling_AutoProvisioningDefaults struct {
	/* Immutable. The Customer Managed Encryption Key used to encrypt the
	boot disk attached to each node in the node pool. */
	// +kcc:proto:field=google.container.v1.AutoprovisioningNodePoolDefaults.boot_disk_kms_key
	BootDiskKMSKeyRef *refsv1beta1.KMSCryptoKeyRef `json:"bootDiskKMSKeyRef,omitempty"`

	/* Size of the disk attached to each node, specified in GB. The smallest allowed disk size is 10GB. */
	// +kcc:proto:field=google.container.v1.AutoprovisioningNodePoolDefaults.disk_size_gb
	DiskSize *int `json:"diskSize,omitempty"`

	/* The default image type used by NAP once a new node pool is being created. */
	// +kcc:proto:field=google.container.v1.AutoprovisioningNodePoolDefaults.image_type
	ImageType *string `json:"imageType,omitempty"`

	/* NodeManagement configuration for this NodePool. */
	// +kcc:proto:field=google.container.v1.AutoprovisioningNodePoolDefaults.management
	NodeManagement *NodeManagement `json:"management,omitempty"`

	/* Minimum CPU platform to be used by this instance. The instance may be scheduled on the specified or newer CPU platform. Applicable values are the friendly names of CPU platforms, such as Intel Haswell. */
	// +kcc:proto:field=google.container.v1.AutoprovisioningNodePoolDefaults.min_cpu_platform
	MinCPUPlatform *string `json:"minCpuPlatform,omitempty"`

	/* Scopes that are used by NAP when creating node pools. */
	// +kcc:proto:field=google.container.v1.AutoprovisioningNodePoolDefaults.oauth_scopes
	OauthScopes []string `json:"oauthScopes,omitempty"`

	// +kcc:proto:field=google.container.v1.AutoprovisioningNodePoolDefaults.service_account
	ServiceAccountRef *refsv1beta1.IAMServiceAccountRef `json:"serviceAccountRef,omitempty"`

	/* Shielded Instance options. */
	// +kcc:proto:field=google.container.v1.AutoprovisioningNodePoolDefaults.shielded_instance_config
	ShieldedInstanceConfig *ShieldedInstanceConfig `json:"shieldedInstanceConfig,omitempty"`

	/* Specifies the upgrade settings for NAP created node pools. */
	// +kcc:proto:field=google.container.v1.AutoprovisioningNodePoolDefaults.upgrade_settings
	NodePool_UpgradeSettings *NodePool_UpgradeSettings `json:"upgradeSettings,omitempty"`
}

// +kcc:proto=google.container.v1.ResourceUsageExportConfig.BigQueryDestination
type ResourceUsageExportConfig_BigQueryDestination struct {
	/* The ID of a BigQuery Dataset. */
	// +required
	// +kcc:proto:field=google.container.v1.ResourceUsageExportConfig.BigQueryDestination.dataset_id
	DatasetID *string `json:"datasetId,omitempty"`
}

// +kcc:proto=google.container.v1.BinaryAuthorization
type BinaryAuthorization struct {
	/* DEPRECATED. Deprecated in favor of evaluation_mode. Enable Binary Authorization for this cluster. */
	// +kcc:proto:field=google.container.v1.BinaryAuthorization.enabled
	Enabled *bool `json:"enabled,omitempty"`

	/* Mode of operation for Binary Authorization policy evaluation. */
	// +kcc:proto:field=google.container.v1.BinaryAuthorization.evaluation_mode
	EvaluationMode *string `json:"evaluationMode,omitempty"`
}

// +kcc:proto=google.container.v1.NodePool.UpdateConfig.BlueGreenSettings
type NodePool_UpdateConfig_BlueGreenSettings struct {
	/* Time needed after draining entire blue pool. After this period, blue pool will be cleaned up.

	A duration in seconds with up to nine fractional digits, ending with 's'. Example: "3.5s". */
	// +kcc:proto:field=google.container.v1.NodePool.UpdateConfig.NodePool_UpdateConfig_BlueGreenSettings.node_pool_soak_duration
	NodePoolSoakDuration *string `json:"nodePoolSoakDuration,omitempty"`

	/* Standard policy for the blue-green upgrade. */
	// +kcc:proto:field=google.container.v1.NodePool.UpdateConfig.NodePool_UpdateConfig_BlueGreenSettings.standard_rollout_policy
	StandardRolloutPolicy *StandardRolloutPolicy `json:"standardRolloutPolicy,omitempty"`
}

// +kcc:proto=google.container.v1.MasterAuthorizedNetworksConfig.CIDRBlock
type MasterAuthorizedNetworksConfig_CIDRBlock struct {
	/* External network that can access Kubernetes master through HTTPS. Must be specified in CIDR notation. */
	// +required
	// +kcc:proto:field=google.container.v1.MasterAuthorizedNetworksConfig.CIDRBlock.cidr_block
	CIDRBlock *string `json:"cidrBlock,omitempty"`

	/* Field for users to identify CIDR blocks. */
	// +kcc:proto:field=google.container.v1.MasterAuthorizedNetworksConfig.CIDRBlock.display_name
	DisplayName *string `json:"displayName,omitempty"`
}

// +kcc:proto=google.container.v1.MasterAuth.ClientCertificateConfig
type MasterAuth_ClientCertificateConfig struct {
	/* Immutable. Whether client certificate authorization is enabled for this cluster. */
	// +required
	// +kcc:proto:field=google.container.v1.MasterAuth.MasterAuth_ClientCertificateConfig.issue_client_certificate
	IssueClientCertificate *bool `json:"issueClientCertificate,omitempty"`
}

// +kcc:proto=google.container.v1.CloudRunConfig
type CloudRunConfig struct {
	// +required
	// +kcc:proto:field=google.container.v1.CloudRunConfig.disabled
	Disabled *bool `json:"disabled,omitempty"`

	// +kcc:proto:field=google.container.v1.CloudRunConfig.load_balancer_type
	LoadBalancerType *string `json:"loadBalancerType,omitempty"`
}

// +kcc:proto=google.container.v1.ClusterAutoscaling
type ClusterAutoscaling struct {
	/* Contains defaults for a node pool created by NAP. */
	// +kcc:proto:field=google.container.v1.ClusterAutoscaling.auto_provisioning_defaults
	ClusterAutoscaling_AutoProvisioningDefaults *ClusterAutoscaling_AutoProvisioningDefaults `json:"autoProvisioningDefaults,omitempty"`

	/* Configuration options for the Autoscaling profile feature, which lets you choose whether the cluster autoscaler should optimize for resource utilization or resource availability when deciding to remove nodes from a cluster. Can be BALANCED or OPTIMIZE_UTILIZATION. Defaults to BALANCED. */
	// +kcc:proto:field=google.container.v1.ClusterAutoscaling.autoscaling_profile
	AutoscalingProfile *string `json:"autoscalingProfile,omitempty"`

	/* Default compute class is a configuration for default compute class. */
	// +kcc:proto:field=google.container.v1.ClusterAutoscaling.default_compute_class_config
	ClusterAutoscaling_DefaultComputeClassConfig *ClusterAutoscaling_DefaultComputeClassConfig `json:"defaultComputeClassConfig,omitempty"`

	/* Whether node auto-provisioning is enabled. Resource limits for cpu and memory must be defined to enable node auto-provisioning. */
	// +kcc:proto:field=google.container.v1.ClusterAutoscaling.enable_node_autoprovisioning
	Enabled *bool `json:"enabled,omitempty"`

	/* Global constraints for machine resources in the cluster. Configuring the cpu and memory types is required if node auto-provisioning is enabled. These limits will apply to node pool autoscaling in addition to node auto-provisioning. */
	// +kcc:proto:field=google.container.v1.ClusterAutoscaling.resource_limits
	ResourceLimits []ResourceLimits `json:"resourceLimits,omitempty"`
}

// +kcc:proto=google.container.v1.ClusterTelemetry
type ClusterTelemetry struct {
	/* Type of the integration. */
	// +required
	// +kcc:proto:field=google.container.v1.ClusterTelemetry.type
	Type *string `json:"type,omitempty"`
}

// +kcc:proto=google.container.v1.ConfidentialNodes
type ConfidentialNodes struct {
	/* Immutable. Confidential instance type for the nodes in the pool. Valid values are SEV, SEV_SNP, and TDX. */
	// +kcc:proto:field=google.container.v1.ConfidentialNodes.confidential_instance_type
	// +optional
	ConfidentialInstanceType *string `json:"confidentialInstanceType,omitempty"`

	/* Immutable. Whether Confidential Nodes feature is enabled for all nodes in this pool. */
	// +required
	// +kcc:proto:field=google.container.v1.ConfidentialNodes.enabled
	Enabled *bool `json:"enabled,omitempty"`
}

// +kcc:proto=google.container.v1.ConfigConnectorConfig
type ConfigConnectorConfig struct {
	// +required
	// +kcc:proto:field=google.container.v1.ConfigConnectorConfig.enabled
	Enabled *bool `json:"enabled,omitempty"`
}

// +kcc:proto=google.container.v1.ControlPlaneEndpointsConfig
type ControlPlaneEndpointsConfig struct {
	/* DNS endpoint configuration. */
	// +kcc:proto:field=google.container.v1.ControlPlaneEndpointsConfig.dns_endpoint_config
	ControlPlaneEndpointsConfig_DNSEndpointConfig *ControlPlaneEndpointsConfig_DNSEndpointConfig `json:"dnsEndpointConfig,omitempty"`

	/* IP endpoint configuration. */
	// +kcc:proto:field=google.container.v1.ControlPlaneEndpointsConfig.ip_endpoints_config
	ControlPlaneEndpointsConfig_IPEndpointsConfig *ControlPlaneEndpointsConfig_IPEndpointsConfig `json:"ipEndpointsConfig,omitempty"`
}

// +kcc:proto=google.container.v1.CostManagementConfig
type CostManagementConfig struct {
	/* Whether to enable GKE cost allocation. When you enable GKE cost allocation, the cluster name and namespace of your GKE workloads appear in the labels field of the billing export to BigQuery. Defaults to false. */
	// +required
	// +kcc:proto:field=google.container.v1.CostManagementConfig.enabled
	Enabled *bool `json:"enabled,omitempty"`
}

// +kcc:proto=google.container.v1.DailyMaintenanceWindow
type DailyMaintenanceWindow struct {
	// +kcc:proto:field=google.container.v1.DailyMaintenanceWindow.duration
	Duration *string `json:"duration,omitempty"`

	// +required
	// +kcc:proto:field=google.container.v1.DailyMaintenanceWindow.start_time
	StartTime *string `json:"startTime,omitempty"`
}

// +kcc:proto=google.container.v1.DatabaseEncryption
type DatabaseEncryption struct {
	/* The key to use to encrypt/decrypt secrets. */
	// +kcc:proto:field=google.container.v1.DatabaseEncryption.key_name
	KeyName *string `json:"keyName,omitempty"`

	/* ENCRYPTED, ALL_OBJECTS_ENCRYPTION_ENABLED or DECRYPTED. */
	// +required
	// +kcc:proto:field=google.container.v1.DatabaseEncryption.state
	State *string `json:"state,omitempty"`
}

// +kcc:proto=google.container.v1.ClusterAutoscaling.DefaultComputeClassConfig
type ClusterAutoscaling_DefaultComputeClassConfig struct {
	/* Enables default compute class. */
	// +required
	// +kcc:proto:field=google.container.v1.ClusterAutoscaling.ClusterAutoscaling_DefaultComputeClassConfig.enabled
	Enabled *bool `json:"enabled,omitempty"`
}

// +kcc:proto=google.container.v1.DefaultSnatStatus
type DefaultSnatStatus struct {
	/* When disabled is set to false, default IP masquerade rules will be applied to the nodes to prevent sNAT on cluster internal traffic. */
	// +required
	// +kcc:proto:field=google.container.v1.DefaultSnatStatus.disabled
	Disabled *bool `json:"disabled,omitempty"`
}

// +kcc:proto=google.container.v1.DnsCacheConfig
type DNSCacheConfig struct {
	// +required
	// +kcc:proto:field=google.container.v1.DNSCacheConfig.enabled
	Enabled *bool `json:"enabled,omitempty"`
}

// +kcc:proto=google.container.v1.DNSConfig
type DNSConfig struct {
	/* Which in-cluster DNS provider should be used. */
	// +kcc:proto:field=google.container.v1.DNSConfig.cluster_dns
	ClusterDNS *string `json:"clusterDns,omitempty"`

	/* The suffix used for all cluster service records. */
	// +kcc:proto:field=google.container.v1.DNSConfig.cluster_dns_domain
	ClusterDNSDomain *string `json:"clusterDnsDomain,omitempty"`

	/* The scope of access to cluster DNS records. */
	// +kcc:proto:field=google.container.v1.DNSConfig.cluster_dns_scope
	ClusterDNSScope *string `json:"clusterDnsScope,omitempty"`
}

// +kcc:proto=google.container.v1.ControlPlaneEndpointsConfig.DNSEndpointConfig
type ControlPlaneEndpointsConfig_DNSEndpointConfig struct {
	/* Controls whether user traffic is allowed over this endpoint. Note that GCP-managed services may still use the endpoint even if this is false. */
	// +kcc:proto:field=google.container.v1.ControlPlaneEndpointsConfig.DNSEndpointConfig.allow_external_traffic
	AllowExternalTraffic *bool `json:"allowExternalTraffic,omitempty"`

	/* Controls whether the k8s token auth is allowed via DNS. */
	// +kcc:proto:field=google.container.v1.ControlPlaneEndpointsConfig.DNSEndpointConfig.enable_k8s_tokens_via_dns
	EnableK8STokensViaDNS *bool `json:"enableK8sTokensViaDns,omitempty"`
}

// +kcc:proto=google.container.v1.K8sBetaAPIConfig
type K8SBetaAPIConfig struct {
	/* Enabled Kubernetes Beta APIs. */
	// +kcc:proto:field=google.container.v1.K8sBetaAPIConfig.enabled_apis
	EnabledApis []string `json:"enabledApis"`
}

// +kcc:proto=google.container.v1.NodeConfig.EphemeralStorageConfig
type EphemeralStorageConfig struct {
	/* Immutable. Number of local SSDs to use to back ephemeral storage. Uses NVMe interfaces. Each local SSD must be 375 or 3000 GB in size, and all local SSDs must share the same size. */
	// +required
	// +kcc:proto:field=google.container.v1.NodeConfig.EphemeralStorageConfig.local_ssd_count
	LocalSsdCount *int `json:"localSsdCount,omitempty"`
}

// +kcc:proto=google.container.v1.EphemeralStorageLocalSsdConfig
type EphemeralStorageLocalSsdConfig struct {
	/* Immutable. Number of local SSDs to be utilized for GKE Data Cache. Uses NVMe interfaces. */
	// +kcc:proto:field=google.container.v1.EphemeralStorageLocalSsdConfig.data_cache_count
	DataCacheCount *int `json:"dataCacheCount,omitempty"`

	/* Immutable. Number of local SSDs to use to back ephemeral storage. Uses NVMe interfaces. Each local SSD must be 375 or 3000 GB in size, and all local SSDs must share the same size. */
	// +required
	// +kcc:proto:field=google.container.v1.EphemeralStorageLocalSsdConfig.local_ssd_count
	LocalSsdCount *int `json:"localSsdCount,omitempty"`
}

// +kcc:proto=google.container.v1.MaintenanceExclusionOptions
type MaintenanceExclusionOptions struct {
	/* The scope of automatic upgrades to restrict in the exclusion window. */
	// +required
	// +kcc:proto:field=google.container.v1.MaintenancePolicy.MaintenanceExclusion.MaintenanceExclusionOptions.scope
	Scope *string `json:"scope,omitempty"`
}

// +kcc:proto=google.container.v1.FastSocket
type FastSocket struct {
	/* Whether or not NCCL Fast Socket is enabled. */
	// +required
	// +kcc:proto:field=google.container.v1.FastSocket.enabled
	Enabled *bool `json:"enabled,omitempty"`
}

// +kcc:proto=google.container.v1.NotificationConfig.Filter
type NotificationConfig_Filter struct {
	/* Can be used to filter what notifications are sent. Valid values include include UPGRADE_AVAILABLE_EVENT, UPGRADE_EVENT and SECURITY_BULLETIN_EVENT. */
	// +required
	// +kcc:proto:field=google.container.v1.NotificationConfig.NotificationConfig_Filter.event_type
	EventType []string `json:"eventType,omitempty"`
}

// +kcc:proto=google.container.v1.GatewayAPIConfig
type GatewayAPIConfig struct {
	/* The Gateway API release channel to use for Gateway API. */
	// +required
	// +kcc:proto:field=google.container.v1.GatewayAPIConfig.channel
	// +required
	Channel *string `json:"channel,omitempty"`
}

// +kcc:proto=google.container.v1.GcePersistentDiskCSIDriverConfig
type GCEPersistentDiskCSIDriverConfig struct {
	// +required
	// +kcc:proto:field=google.container.v1.GcePersistentDiskCSIDriverConfig.enabled
	Enabled *bool `json:"enabled,omitempty"`
}

// +kcc:proto=google.container.v1.GcfsConfig
type GcfsConfig struct {
	/* Whether or not GCFS is enabled. */
	// +required
	// +kcc:proto:field=google.container.v1.GcfsConfig.enabled
	Enabled *bool `json:"enabled,omitempty"`
}

// +kcc:proto=google.container.v1.GCPFilestoreCSIDriverConfig
type GCPFilestoreCSIDriverConfig struct {
	// +required
	// +kcc:proto:field=google.container.v1.GCPFilestoreCSIDriverConfig.enabled
	Enabled *bool `json:"enabled,omitempty"`
}

// +kcc:proto=google.container.v1.GcsFuseCSIDriverConfig
type GCSFuseCSIDriverConfig struct {
	// +required
	// +kcc:proto:field=google.container.v1.GcsFuseCSIDriverConfig.enabled
	Enabled *bool `json:"enabled,omitempty"`
}

// +kcc:proto=google.container.v1.GkeBackupAgentConfig
type GKEBackupAgentConfig struct {
	// +required
	// +kcc:proto:field=google.container.v1.GkeBackupAgentConfig.enabled
	Enabled *bool `json:"enabled,omitempty"`
}

// +kcc:proto=google.container.v1.GPUDriverInstallationConfig
type GPUDriverInstallationConfig struct {
	/* Immutable. Mode for how the GPU driver is installed. */
	// +required
	// +kcc:proto:field=google.container.v1.GPUDriverInstallationConfig.gpu_driver_version
	GPUDriverVersion *string `json:"gpuDriverVersion,omitempty"`
}

// +kcc:proto=google.container.v1.GPUSharingConfig
type GPUSharingConfig struct {
	/* Immutable. The type of GPU sharing strategy to enable on the GPU node. Possible values are described in the API package (https://pkg.go.dev/google.golang.org/api/container/v1#GPUSharingConfig). */
	// +required
	// +kcc:proto:field=google.container.v1.GPUSharingConfig.gpu_sharing_strategy
	GPUSharingStrategy *string `json:"gpuSharingStrategy,omitempty"`

	/* Immutable. The maximum number of containers that can share a GPU. */
	// +required
	// +kcc:proto:field=google.container.v1.GPUSharingConfig.max_shared_clients_per_gpu
	MaxSharedClientsPerGPU *int `json:"maxSharedClientsPerGpu,omitempty"`
}

// +kcc:proto=google.container.v1.AcceleratorConfig
type AcceleratorConfig struct {
	/* Immutable. The number of the accelerator cards exposed to an instance. */
	// +required
	// +kcc:proto:field=google.container.v1.AcceleratorConfig.accelerator_count
	Count *int `json:"count,omitempty"`

	/* Immutable. Configuration for auto installation of GPU driver. */
	// +kcc:proto:field=google.container.v1.AcceleratorConfig.gpu_driver_installation_config
	GPUDriverInstallationConfig *GPUDriverInstallationConfig `json:"gpuDriverInstallationConfig,omitempty"`

	/* Immutable. Size of partitions to create on the GPU. Valid values are described in the NVIDIA mig user guide (https://docs.nvidia.com/datacenter/tesla/mig-user-guide/#partitioning). */
	// +kcc:proto:field=google.container.v1.AcceleratorConfig.gpu_partition_size
	GPUPartitionSize *string `json:"gpuPartitionSize,omitempty"`

	/* Immutable. Configuration for GPU sharing. */
	// +kcc:proto:field=google.container.v1.AcceleratorConfig.gpu_sharing_config
	GPUSharingConfig *GPUSharingConfig `json:"gpuSharingConfig,omitempty"`

	/* Immutable. The accelerator type resource name. */
	// +required
	// +kcc:proto:field=google.container.v1.AcceleratorConfig.accelerator_type
	Type *string `json:"type,omitempty"`
}

// +kcc:proto=google.container.v1.VirtualNic
type VirtualNic struct {
	/* Immutable. Whether or not gvnic is enabled. */
	// +kcc:proto:field=google.container.v1.VirtualNic.enabled
	// +required
	Enabled *bool `json:"enabled,omitempty"`
}

// +kcc:proto=google.container.v1.HorizontalPodAutoscaling
type HorizontalPodAutoscaling struct {
	// +required
	// +kcc:proto:field=google.container.v1.HorizontalPodAutoscaling.disabled
	Disabled *bool `json:"disabled,omitempty"`
}

// +kcc:proto=google.container.v1.NodeConfig.HostMaintenancePolicy
type HostMaintenancePolicy struct {
	/* Immutable. . */
	// +required
	// +kcc:proto:field=google.container.v1.NodeConfig.HostMaintenancePolicy.maintenance_interval
	MaintenanceInterval *string `json:"maintenanceInterval,omitempty"`
}

// +kcc:proto=google.container.v1.HttpLoadBalancing
type HTTPLoadBalancing struct {
	// +required
	// +kcc:proto:field=google.container.v1.HttpLoadBalancing.disabled
	Disabled *bool `json:"disabled,omitempty"`
}

// +kcc:proto=google.container.v1.IdentityServiceConfig
type IdentityServiceConfig struct {
	/* Whether to enable the Identity Service component. */
	// +kcc:proto:field=google.container.v1.IdentityServiceConfig.enabled
	Enabled *bool `json:"enabled,omitempty"`
}

// +kcc:proto=google.container.v1.AdditionalIPRangesConfig
type AdditionalIPRangesConfig struct {
	/* The subnetwork path for the additional IP range. Format: projects/{project}/regions/{region}/subnetworks/{subnetwork}. */
	// +required
	// +kcc:proto:field=google.container.v1.AdditionalIPRangesConfig.subnetwork
	SubnetworkRef *refsv1beta1.ComputeSubnetworkRef `json:"subnetworkRef,omitempty"`

	/* List of secondary ranges names within this subnetwork that can be used for pod IPs. */
	// +kcc:proto:field=google.container.v1.AdditionalIPRangesConfig.pod_ipv4_range_names
	PodIPV4RangeNames []string `json:"podIpv4RangeNames,omitempty"`

	/* Status of the subnetwork, If in draining status, subnet will not be selected for new node pools. */
	Status *string `json:"status,omitempty"`
}

// +kcc:proto=google.container.v1.IPAllocationPolicy
type IPAllocationPolicy struct {
	/* AdditionalIPRangesConfig is the configuration for additional pod secondary ranges supporting the ClusterUpdate message. Each AdditionalIPRangesConfig corresponds to a single subnetwork. */
	// +kcc:proto:field=google.container.v1.IPAllocationPolicy.additional_ip_ranges_configs
	AdditionalIPRangesConfigs []AdditionalIPRangesConfig `json:"additionalIpRangesConfigs,omitempty"`

	/* AdditionalPodRangesConfig is the configuration for additional pod secondary ranges supporting the ClusterUpdate message. */
	// +kcc:proto:field=google.container.v1.IPAllocationPolicy.additional_pod_ranges_config
	AdditionalPodRangesConfig *AdditionalPodRangesConfig `json:"additionalPodRangesConfig,omitempty"`

	/* Immutable. The IP address range for the cluster pod IPs. Set to blank to have a range chosen with the default size. Set to /netmask (e.g. /14) to have a range chosen with a specific netmask. Set to a CIDR notation (e.g. 10.96.0.0/14) from the RFC-1918 private networks (e.g. 10.0.0.0/8, 172.16.0.0/12, 192.168.0.0/16) to pick a specific range to use. */
	// +kcc:proto:field=google.container.v1.IPAllocationPolicy.cluster_ipv4_cidr_block
	ClusterIPV4CIDRBlock *string `json:"clusterIpv4CidrBlock,omitempty"`

	/* Immutable. The name of the existing secondary range in the cluster's subnetwork to use for pod IP addresses. Alternatively, cluster_ipv4_cidr_block can be used to automatically create a GKE-managed one. */
	// +kcc:proto:field=google.container.v1.IPAllocationPolicy.cluster_secondary_range_name
	ClusterSecondaryRangeName *string `json:"clusterSecondaryRangeName,omitempty"`

	/* Immutable. Configuration for cluster level pod cidr overprovision. Default is disabled=false. */
	// +kcc:proto:field=google.container.v1.IPAllocationPolicy.pod_cidr_overprovision_config
	PodCIDROverprovisionConfig *PodCIDROverprovisionConfig `json:"podCidrOverprovisionConfig,omitempty"`

	/* Immutable. The IP address range of the services IPs in this cluster. Set to blank to have a range chosen with the default size. Set to /netmask (e.g. /14) to have a range chosen with a specific netmask. Set to a CIDR notation (e.g. 10.96.0.0/14) from the RFC-1918 private networks (e.g. 10.0.0.0/8, 172.16.0.0/12, 192.168.0.0/16) to pick a specific range to use. */
	// +kcc:proto:field=google.container.v1.IPAllocationPolicy.services_ipv4_cidr_block
	ServicesIPV4CIDRBlock *string `json:"servicesIpv4CidrBlock,omitempty"`

	/* Immutable. The name of the existing secondary range in the cluster's subnetwork to use for service ClusterIPs. Alternatively, services_ipv4_cidr_block can be used to automatically create a GKE-managed one. */
	// +kcc:proto:field=google.container.v1.IPAllocationPolicy.services_secondary_range_name
	ServicesSecondaryRangeName *string `json:"servicesSecondaryRangeName,omitempty"`

	/* Immutable. The IP Stack type of the cluster. Choose between IPV4 and IPV4_IPV6. Default type is IPV4 Only if not set. */
	// +kcc:proto:field=google.container.v1.IPAllocationPolicy.stack_type
	StackType *string `json:"stackType,omitempty"`
}

// +kcc:proto=google.container.v1.ControlPlaneEndpointsConfig.IPEndpointsConfig
type ControlPlaneEndpointsConfig_IPEndpointsConfig struct {
	/* Controls whether to allow direct IP access. When false, configuration of masterAuthorizedNetworksConfig, privateClusterConfig.enablePrivateEndpoint, privateClusterConfig.privateEndpointSubnetwork and privateClusterConfig.masterGlobalAccessConfig fields won't be used, and privateClusterConfig.privateEndpoint and privateClusterConfig.publicEndpoint fields won't be populated. */
	// +kcc:proto:field=google.container.v1.ControlPlaneEndpointsConfig.IPEndpointsConfig.enabled
	Enabled *bool `json:"enabled,omitempty"`
}

// +kcc:proto=google.container.v1.IstioConfig
type IstioConfig struct {
	/* The authentication type between services in Istio. Available options include AUTH_MUTUAL_TLS. */
	// +kcc:proto:field=google.container.v1.IstioConfig.auth
	Auth *string `json:"auth,omitempty"`

	/* The status of the Istio addon, which makes it easy to set up Istio for services in a cluster. It is disabled by default. Set disabled = false to enable. */
	// +required
	// +kcc:proto:field=google.container.v1.IstioConfig.disabled
	Disabled *bool `json:"disabled,omitempty"`
}

// +kcc:proto=google.container.v1.KalmConfig
type KalmConfig struct {
	// +required
	// +kcc:proto:field=google.container.v1.KalmConfig.enabled
	Enabled *bool `json:"enabled,omitempty"`
}

// +kcc:proto=google.container.v1.NodeKubeletConfig
type KubeletConfig struct {
	/* Enable CPU CFS quota enforcement for containers that specify CPU limits. */
	// +kcc:proto:field=google.container.v1.NodeKubeletConfig.cpu_cfs_quota
	CPUCfsQuota *bool `json:"cpuCfsQuota,omitempty"`

	/* Set the CPU CFS quota period value 'cpu.cfs_period_us'. */
	// +kcc:proto:field=google.container.v1.NodeKubeletConfig.cpu_cfs_quota_period
	CPUCfsQuotaPeriod *string `json:"cpuCfsQuotaPeriod,omitempty"`

	/* Control the CPU management policy on the node. */
	// +required
	// +kcc:proto:field=google.container.v1.NodeKubeletConfig.cpu_manager_policy
	CPUManagerPolicy *string `json:"cpuManagerPolicy,omitempty"`

	/* Controls the maximum number of processes allowed to run in a pod. */
	// +kcc:proto:field=google.container.v1.NodeKubeletConfig.pod_pids_limit
	PodPidsLimit *int `json:"podPidsLimit,omitempty"`
}

// +kcc:proto=google.container.v1.LinuxNodeConfig
type LinuxNodeConfig struct {
	/* cgroupMode specifies the cgroup mode to be used on the node. */
	// +kcc:proto:field=google.container.v1.LinuxNodeConfig.cgroup_mode
	CgroupMode *string `json:"cgroupMode,omitempty"`

	/* The Linux kernel parameters to be applied to the nodes and all pods running on the nodes. */
	// +kcc:proto:field=google.container.v1.LinuxNodeConfig.sysctls
	Sysctls map[string]string `json:"sysctls,omitempty"`
}

// +kcc:proto=google.container.v1.NodeConfig.LocalNvmeSsdBlockConfig
type NodeConfig_LocalNvmeSsdBlockConfig struct {
	/* Immutable. Number of raw-block local NVMe SSD disks to be attached to the node. Each local SSD is 375 GB in size. */
	// +required
	// +kcc:proto:field=google.container.v1.NodeConfig.LocalNvmeSsdBlockConfig.local_ssd_count
	LocalSsdCount *int `json:"localSsdCount,omitempty"`
}

// +kcc:proto=google.container.v1.LoggingConfig
type LoggingConfig struct {
	/* GKE components exposing logs. Valid values include SYSTEM_COMPONENTS, APISERVER, CONTROLLER_MANAGER, SCHEDULER, and WORKLOADS. */
	// +kcc:proto:field=google.container.v1.LoggingConfig.enable_components
	EnableComponents []string `json:"enableComponents"`
}

// +kcc:proto=google.container.v1.MaintenancePolicy.MaintenanceExclusion
type MaintenanceExclusion struct {
	// +kcc:proto:field=google.container.v1.MaintenancePolicy.MaintenanceExclusion.end_time
	// +required
	EndTime *string `json:"endTime,omitempty"`

	// +kcc:proto:field=google.container.v1.MaintenancePolicy.MaintenanceExclusion.exclusion_name
	// +required
	ExclusionName *string `json:"exclusionName,omitempty"`

	/* Maintenance exclusion related options. */
	// +kcc:proto:field=google.container.v1.MaintenancePolicy.MaintenanceExclusion.exclusion_options
	MaintenanceExclusionOptions *MaintenanceExclusionOptions `json:"exclusionOptions,omitempty"`

	// +kcc:proto:field=google.container.v1.MaintenancePolicy.MaintenanceExclusion.start_time
	// +required
	StartTime *string `json:"startTime,omitempty"`
}

// +kcc:proto=google.container.v1.MaintenancePolicy
type MaintenancePolicy struct {
	/* Time window specified for daily maintenance operations. Specify start_time in RFC3339 format "HH:MM", where HH : [00-23] and MM : [00-59] GMT. */
	// +kcc:proto:field=google.container.v1.MaintenancePolicy.window.daily_maintenance_window
	DailyMaintenanceWindow *DailyMaintenanceWindow `json:"dailyMaintenanceWindow,omitempty"`

	/* Exceptions to maintenance window. Non-emergency maintenance should not occur in these windows. */
	// +kcc:proto:field=google.container.v1.MaintenancePolicy.window.maintenance_exclusions
	MaintenanceExclusion []MaintenanceExclusion `json:"maintenanceExclusion,omitempty"`

	/* Time window for recurring maintenance operations. */
	// +kcc:proto:field=google.container.v1.MaintenancePolicy.window.recurring_window
	RecurringWindow *RecurringWindow `json:"recurringWindow,omitempty"`
}

// +kcc:proto=google.container.v1.ManagedPrometheusConfig
type ManagedPrometheus struct {
	/* Whether or not the managed collection is enabled. */
	// +kcc:proto:field=google.container.v1.ManagedPrometheusConfig.enabled
	// +required
	Enabled *bool `json:"enabled,omitempty"`
}

// +kcc:proto=google.container.v1.NodeManagement
type NodeManagement struct {
	/* Specifies whether the node auto-repair is enabled for the node pool. If enabled, the nodes in this node pool will be monitored and, if they fail health checks too many times, an automatic repair action will be triggered. */
	// +kcc:proto:field=google.container.v1.NodeManagement.auto_repair
	AutoRepair *bool `json:"autoRepair,omitempty"`

	/* Specifies whether node auto-upgrade is enabled for the node pool. If enabled, node auto-upgrade helps keep the nodes in your node pool up to date with the latest release version of Kubernetes. */
	// +kcc:proto:field=google.container.v1.NodeManagement.auto_upgrade
	AutoUpgrade *bool `json:"autoUpgrade,omitempty"`

	/* Specifies the Auto Upgrade knobs for the node pool. */
	// +kcc:proto:field=google.container.v1.NodeManagement.upgrade_options
	NodeManagement_UpgradeOptions []NodeManagement_UpgradeOptions `json:"upgradeOptions,omitempty"`
}

// +kcc:proto=google.container.v1.MasterAuth
type MasterAuth struct {
	/* Base64 encoded public certificate used by clients to authenticate to the cluster endpoint. */
	// +kcc:proto:field=google.container.v1.MasterAuth.client_certificate
	ClientCertificate *string `json:"clientCertificate,omitempty"`

	/* Immutable. Whether client certificate authorization is enabled for this cluster. */
	// +kcc:proto:field=google.container.v1.MasterAuth.client_certificate_config
	MasterAuth_ClientCertificateConfig *MasterAuth_ClientCertificateConfig `json:"clientCertificateConfig,omitempty"`

	/* Base64 encoded private key used by clients to authenticate to the cluster endpoint. */
	// +kcc:proto:field=google.container.v1.MasterAuth.client_key
	ClientKey *string `json:"clientKey,omitempty"`

	/* Base64 encoded public certificate that is the root of trust for the cluster. */
	// +kcc:proto:field=google.container.v1.MasterAuth.cluster_ca_certificate
	ClusterCACertificate *string `json:"clusterCaCertificate,omitempty"`

	/* The password to use for HTTP basic authentication when accessing the Kubernetes master endpoint. */
	// +kcc:proto:field=google.container.v1.MasterAuth.password
	Password *secretv1beta1.Legacy `json:"password,omitempty"`

	/* The username to use for HTTP basic authentication when accessing the Kubernetes master endpoint. If not present basic auth will be disabled. */
	// +kcc:proto:field=google.container.v1.MasterAuth.username
	Username *string `json:"username,omitempty"`
}

// +kcc:proto=google.container.v1.MasterAuthorizedNetworksConfig
type MasterAuthorizedNetworksConfig struct {
	/* External networks that can access the Kubernetes cluster master through HTTPS. */
	// +kcc:proto:field=google.container.v1.MasterAuthorizedNetworksConfig.cidr_blocks
	MasterAuthorizedNetworksConfig_CIDRBlock []MasterAuthorizedNetworksConfig_CIDRBlock `json:"cidrBlocks,omitempty"`

	/* Whether master is accessible via Google Compute Engine Public IP addresses. */
	// +kcc:proto:field=google.container.v1.MasterAuthorizedNetworksConfig.gcp_public_cidrs_access_enabled
	GCPPublicCidrsAccessEnabled *bool `json:"gcpPublicCidrsAccessEnabled,omitempty"`
}

// +kcc:proto=google.container.v1.PrivateClusterConfig.PrivateClusterMasterGlobalAccessConfig
type PrivateClusterMasterGlobalAccessConfig struct {
	/* Whether the cluster master is accessible globally or not. */
	// +kcc:proto:field=google.container.v1.PrivateClusterConfig.PrivateClusterMasterGlobalAccessConfig.enabled
	// +required
	Enabled *bool `json:"enabled,omitempty"`
}

// +kcc:proto=google.container.v1.MeshCertificates
type MeshCertificates struct {
	/* When enabled the GKE Workload Identity Certificates controller and node agent will be deployed in the cluster. */
	// +required
	// +kcc:proto:field=google.container.v1.MeshCertificates.enable_certificates
	EnableCertificates *bool `json:"enableCertificates,omitempty"`
}

// +kcc:proto=google.container.v1.MonitoringConfig
type MonitoringConfig struct {
	/* Configuration of Advanced Datapath Observability features. */
	// +kcc:proto:field=google.container.v1.MonitoringConfig.advanced_datapath_observability_config
	AdvancedDatapathObservabilityConfig []*AdvancedDatapathObservabilityConfig `json:"advancedDatapathObservabilityConfig,omitempty"`

	/* GKE components exposing metrics. Valid values include SYSTEM_COMPONENTS, APISERVER, SCHEDULER, CONTROLLER_MANAGER, STORAGE, HPA, POD, DAEMONSET, DEPLOYMENT, STATEFULSET and WORKLOADS. */
	// +kcc:proto:field=google.container.v1.MonitoringConfig.enable_components
	EnableComponents []string `json:"enableComponents,omitempty"`

	/* Configuration for Google Cloud Managed Services for Prometheus. */
	// +kcc:proto:field=google.container.v1.MonitoringConfig.managed_prometheus_config
	ManagedPrometheus *ManagedPrometheus `json:"managedPrometheus,omitempty"`
}

// +kcc:proto=google.container.v1.NetworkPolicy
type NetworkPolicy struct {
	/* Whether network policy is enabled on the cluster. */
	// +required
	// +kcc:proto:field=google.container.v1.NetworkPolicy.enabled
	Enabled *bool `json:"enabled,omitempty"`

	/* The selected network policy provider. Defaults to PROVIDER_UNSPECIFIED. */
	// +kcc:proto:field=google.container.v1.NetworkPolicy.provider
	Provider *string `json:"provider,omitempty"`
}

// +kcc:proto=google.container.v1.NetworkPolicyConfig
type NetworkPolicyConfig struct {
	// +required
	// +kcc:proto:field=google.container.v1.NetworkPolicyConfig.disabled
	Disabled *bool `json:"disabled,omitempty"`
}

// +kcc:proto=google.container.v1.NetworkTags
type NetworkTags struct {
	/* List of network tags applied to auto-provisioned node pools. */
	// +kcc:proto:field=google.container.v1.NetworkTags.tags
	Tags []string `json:"tags,omitempty"`
}

// +kcc:proto=google.container.v1.NodeAffinity
type NodeAffinity struct {
	/* Immutable. . */
	// +kcc:proto:field=google.container.v1.NodeAffinity.key
	// +required
	Key *string `json:"key,omitempty"`

	/* Immutable. . */
	// +kcc:proto:field=google.container.v1.NodeAffinity.operator
	// +required
	Operator *string `json:"operator,omitempty"`

	/* Immutable. . */
	// +kcc:proto:field=google.container.v1.NodeAffinity.values
	// +required
	Values []string `json:"values,omitempty"`
}

// +kcc:proto=google.container.v1.NodeConfig
type NodeConfig struct {
	/* Immutable. Specifies options for controlling advanced machine features. */
	// +kcc:proto:field=google.container.v1.NodeConfig.advanced_machine_features
	NodeConfig_AdvancedMachineFeatures *NodeConfig_AdvancedMachineFeatures `json:"advancedMachineFeatures,omitempty"`

	// +kcc:proto:field=google.container.v1.NodeConfig.boot_disk_kms_key
	BootDiskKMSCryptoKeyRef *refsv1beta1.KMSCryptoKeyRef `json:"bootDiskKMSCryptoKeyRef,omitempty"`

	/* Immutable. Configuration for the confidential nodes feature, which makes nodes run on confidential VMs. Warning: This configuration can't be changed (or added/removed) after pool creation without deleting and recreating the entire pool. */
	// +kcc:proto:field=google.container.v1.NodeConfig.confidential_nodes
	ConfidentialNodes *ConfidentialNodes `json:"confidentialNodes,omitempty"`

	/* Immutable. Size of the disk attached to each node, specified in GB. The smallest allowed disk size is 10GB. */
	// +kcc:proto:field=google.container.v1.NodeConfig.disk_size_gb
	DiskSizeGb *int `json:"diskSizeGb,omitempty"`

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
	LocalSsdCount *int `json:"localSsdCount,omitempty"`

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

	/* Immutable. The workload metadata configuration for this node. */
	// +kcc:proto:field=google.container.v1.NodeConfig.workload_metadata_config
	NodeConfig_WorkloadMetadataConfig *NodeConfig_WorkloadMetadataConfig `json:"workloadMetadataConfig,omitempty"`
}

// +kcc:proto=google.container.v1.NodeConfigDefaults
type NodeConfigDefaults struct {
	/* GCFS configuration for this node. */
	// +kcc:proto:field=google.container.v1.NodeConfigDefaults.gcfs_config
	GcfsConfig *GcfsConfig `json:"gcfsConfig,omitempty"`

	/* Type of logging agent that is used as the default value for node pools in the cluster. Valid values include DEFAULT and MAX_THROUGHPUT. */
	// +kcc:proto:field=google.container.v1.NodeConfigDefaults.logging_config
	LoggingVariant *string `json:"loggingVariant,omitempty"`
}

// +kcc:proto=google.container.v1.NodePoolAutoConfig
type NodePoolAutoConfig struct {
	/* Collection of Compute Engine network tags that can be applied to a node's underlying VM instance. */
	// +kcc:proto:field=google.container.v1.NetworkTags.tags
	NetworkTags *NetworkTags `json:"networkTags,omitempty"`
}

// +kcc:proto=google.container.v1.NodePoolDefaults
type NodePoolDefaults struct {
	/* Subset of NodeConfig message that has defaults. */
	// +kcc:proto:field=google.container.v1.NodePoolDefaults.node_config_defaults
	NodeConfigDefaults *NodeConfigDefaults `json:"nodeConfigDefaults,omitempty"`
}

// +kcc:proto=google.container.v1.NotificationConfig
type NotificationConfig struct {
	/* Notification config for Cloud Pub/Sub. */
	// +kcc:proto:field=google.container.v1.NotificationConfig.pubsub
	// +required
	Pubsub *NotificationConfig_PubSub `json:"pubsub,omitempty"`
}

// +kcc:proto=google.container.v1.PodCIDROverprovisionConfig
type PodCIDROverprovisionConfig struct {
	// +kcc:proto:field=google.container.v1.PodCIDROverprovisionConfig.disable
	// +required
	Disabled *bool `json:"disabled,omitempty"`
}

// +kcc:proto=google.container.v1.PodSecurityPolicyConfig
type PodSecurityPolicyConfig struct {
	/* Enable the PodSecurityPolicy controller for this cluster. If enabled, pods must be valid under a PodSecurityPolicy to be created. */
	// +kcc:proto:field=google.container.v1.PodSecurityPolicyConfig.enabled
	// +required
	Enabled *bool `json:"enabled,omitempty"`
}

// +kcc:proto=google.container.v1.PrivateClusterConfig
type PrivateClusterConfig struct {
	/* When true, the cluster's private endpoint is used as the cluster endpoint and access through the public endpoint is disabled. When false, either endpoint can be used. */
	// +kcc:proto:field=google.container.v1.PrivateClusterConfig.enable_private_endpoint
	EnablePrivateEndpoint *bool `json:"enablePrivateEndpoint,omitempty"`

	/* Enables the private cluster feature, creating a private endpoint on the cluster. In a private cluster, nodes only have RFC 1918 private addresses and communicate with the master's private endpoint via private networking. */
	// +kcc:proto:field=google.container.v1.PrivateClusterConfig.enable_private_nodes
	EnablePrivateNodes *bool `json:"enablePrivateNodes,omitempty"`

	/* Controls cluster master global access settings. */
	// +kcc:proto:field=google.container.v1.PrivateClusterConfig.master_global_access_config
	PrivateClusterMasterGlobalAccessConfig *PrivateClusterMasterGlobalAccessConfig `json:"masterGlobalAccessConfig,omitempty"`

	/* Immutable. The IP range in CIDR notation to use for the hosted master network. This range will be used for assigning private IP addresses to the cluster master(s) and the ILB VIP. This range must not overlap with any other ranges in use within the cluster's network, and it must be a /28 subnet. See Private Cluster Limitations for more details. This field only applies to private clusters, when enable_private_nodes is true. */
	// +kcc:proto:field=google.container.v1.PrivateClusterConfig.master_ipv4_cidr_block
	MasterIPV4CIDRBlock *string `json:"masterIpv4CidrBlock,omitempty"`

	/* The name of the peering between this cluster and the Google owned VPC. */
	// +kcc:proto:field=google.container.v1.PrivateClusterConfig.peering_name
	PeeringName *string `json:"peeringName,omitempty"`

	/* The internal IP address of this cluster's master endpoint. */
	// +kcc:proto:field=google.container.v1.PrivateClusterConfig.private_endpoint
	PrivateEndpoint *string `json:"privateEndpoint,omitempty"`

	/* Immutable. Subnetwork in cluster's network where master's endpoint
	will be provisioned. */
	// +kcc:proto:field=google.container.v1.PrivateClusterConfig.private_endpoint_subnetwork
	PrivateEndpointSubnetworkRef *refsv1beta1.ComputeSubnetworkRef `json:"privateEndpointSubnetworkRef,omitempty"`

	/* The external IP address of this cluster's master endpoint. */
	// +kcc:proto:field=google.container.v1.PrivateClusterConfig.public_endpoint
	PublicEndpoint *string `json:"publicEndpoint,omitempty"`
}

type ProtectConfig struct {
	/* WorkloadConfig defines which actions are enabled for a cluster's workload configurations. */
	WorkloadConfig *WorkloadConfig `json:"workloadConfig,omitempty"`

	/* Sets which mode to use for Protect workload vulnerability scanning feature. Accepted values are DISABLED, BASIC. */
	WorkloadVulnerabilityMode *string `json:"workloadVulnerabilityMode,omitempty"`
}

// +kcc:proto=google.container.v1.NotificationConfig.PubSub
type NotificationConfig_PubSub struct {
	/* Whether or not the notification config is enabled. */
	// +kcc:proto:field=google.container.v1.NotificationConfig.NotificationConfig_PubSub.enabled
	// +required
	Enabled *bool `json:"enabled,omitempty"`

	/* Allows filtering to one or more specific event types. If event types are present, those and only those event types will be transmitted to the cluster. Other types will be skipped. If no filter is specified, or no event types are present, all event types will be sent. */
	// +kcc:proto:field=google.container.v1.NotificationConfig.NotificationConfig_PubSub.filter
	NotificationConfig_Filter *NotificationConfig_Filter `json:"filter,omitempty"`

	/* The PubSubTopic to send the notification to. */
	// +kcc:proto:field=google.container.v1.NotificationConfig.NotificationConfig_PubSub.topic
	TopicRef *pubsubv1beta1.PubSubTopicRef `json:"topicRef,omitempty"`
}

// +kcc:proto=google.container.v1.RecurringTimeWindow.window
type RecurringWindow struct {
	// +kcc:proto:field=google.container.v1.RecurringTimeWindow.window.end_time
	// +required
	EndTime *string `json:"endTime,omitempty"`

	// +kcc:proto:field=google.container.v1.RecurringTimeWindow.recurrence
	// +required
	Recurrence *string `json:"recurrence,omitempty"`

	// +kcc:proto:field=google.container.v1.RecurringTimeWindow.window.start_time
	// +required
	StartTime *string `json:"startTime,omitempty"`
}

// +kcc:proto=google.container.v1.ReleaseChannel
type ReleaseChannel struct {
	/* The selected release channel. Accepted values are:
	* UNSPECIFIED: Not set.
	* RAPID: Weekly upgrade cadence; Early testers and developers who requires new features.
	* REGULAR: Multiple per month upgrade cadence; Production users who need features not yet offered in the Stable channel.
	* STABLE: Every few months upgrade cadence; Production users who need stability above all else, and for whom frequent upgrades are too risky. */
	// +required
	// +kcc:proto:field=google.container.v1.ReleaseChannel.channel
	// +required
	Channel *string `json:"channel,omitempty"`
}

// +kcc:proto=google.container.v1.ReservationAffinity
type ReservationAffinity struct {
	/* Immutable. Corresponds to the type of reservation consumption. */
	// +required
	// +kcc:proto:field=google.container.v1.ReservationAffinity.consume_reservation_type
	ConsumeReservationType *string `json:"consumeReservationType,omitempty"`

	/* Immutable. The label key of a reservation resource. */
	// +kcc:proto:field=google.container.v1.ReservationAffinity.key
	Key *string `json:"key,omitempty"`

	/* Immutable. The label values of the reservation resource. */
	// +kcc:proto:field=google.container.v1.ReservationAffinity.values
	Values []string `json:"values,omitempty"`
}

// +kcc:proto=google.container.v1.ResourceLimit
type ResourceLimits struct {
	/* Maximum amount of the resource in the cluster. */
	// +kcc:proto:field=google.container.v1.ResourceLimit.maximum
	Maximum *int `json:"maximum,omitempty"`

	/* Minimum amount of the resource in the cluster. */
	// +kcc:proto:field=google.container.v1.ResourceLimit.minimum
	Minimum *int `json:"minimum,omitempty"`

	/* The type of the resource. For example, cpu and memory. See the guide to using Node Auto-Provisioning for a list of types. */
	// +required
	// +kcc:proto:field=google.container.v1.ResourceLimit.resource_type
	ResourceType *string `json:"resourceType,omitempty"`
}

// +kcc:proto=google.container.v1.ResourceUsageExportConfig
type ResourceUsageExportConfig struct {
	/* Parameters for using BigQuery as the destination of resource usage export. */
	// +kcc:proto:field=google.container.v1.ResourceUsageExportConfig.bigquery_destination
	ResourceUsageExportConfig_BigQueryDestination ResourceUsageExportConfig_BigQueryDestination `json:"bigqueryDestination"`

	/* Whether to enable network egress metering for this cluster. If enabled, a daemonset will be created in the cluster to meter network egress traffic. */
	// +kcc:proto:field=google.container.v1.ResourceUsageExportConfig.enable_network_egress_metering
	EnableNetworkEgressMetering *bool `json:"enableNetworkEgressMetering,omitempty"`

	/* Whether to enable resource consumption metering on this cluster. When enabled, a table will be created in the resource export BigQuery dataset to store resource consumption data. The resulting table can be joined with the resource usage table or with BigQuery billing export. Defaults to true. */
	// +kcc:proto:field=google.container.v1.ResourceUsageExportConfig.consumption_metering_config.enabled
	EnableResourceConsumptionMetering *bool `json:"enableResourceConsumptionMetering,omitempty"`
}

// +kcc:proto=google.container.v1.SandboxConfig
type SandboxConfig struct {
	/* Type of the sandbox to use for the node (e.g. 'gvisor'). */
	// +required
	// +kcc:proto:field=google.container.v1.SandboxConfig.type
	SandboxType *string `json:"sandboxType,omitempty"`
}

// +kcc:proto=google.container.v1.SecurityPostureConfig
type SecurityPostureConfig struct {
	/* Sets the mode of the Kubernetes security posture API's off-cluster features. Available options include DISABLED and BASIC. */
	// +kcc:proto:field=google.container.v1.SecurityPostureConfig.mode
	Mode *string `json:"mode,omitempty"`

	/* Sets the mode of the Kubernetes security posture API's workload vulnerability scanning. Available options include VULNERABILITY_DISABLED and VULNERABILITY_BASIC. */
	// +kcc:proto:field=google.container.v1.SecurityPostureConfig.vulnerability_mode
	VulnerabilityMode *string `json:"vulnerabilityMode,omitempty"`
}

// +kcc:proto=google.container.v1.ServiceExternalIPsConfig
type ServiceExternalIPsConfig struct {
	/* When enabled, services with exterenal ips specified will be allowed. */
	// +required
	// +kcc:proto:field=google.container.v1.ServiceExternalIPsConfig.enabled
	Enabled *bool `json:"enabled,omitempty"`
}

// +kcc:proto=google.container.v1.ShieldedInstanceConfig
type ShieldedInstanceConfig struct {
	/* Immutable. Defines whether the instance has integrity monitoring enabled. */
	// +kcc:proto:field=google.container.v1.ShieldedInstanceConfig.enable_integrity_monitoring
	EnableIntegrityMonitoring *bool `json:"enableIntegrityMonitoring,omitempty"`

	/* Immutable. Defines whether the instance has Secure Boot enabled. */
	// +kcc:proto:field=google.container.v1.ShieldedInstanceConfig.enable_secure_boot
	EnableSecureBoot *bool `json:"enableSecureBoot,omitempty"`
}

// +kcc:proto=google.container.v1.NodeConfig.SoleTenantConfig
type NodeConfig_SoleTenantConfig struct {
	/* Immutable. . */
	// +kcc:proto:field=google.container.v1.NodeConfig.SoleTenantConfig.node_affinities
	NodeAffinity []NodeAffinity `json:"nodeAffinity"`
}

// +kcc:proto=google.container.v1.NodePool.UpdateConfig.BlueGreenSettings.StandardRolloutPolicy
type StandardRolloutPolicy struct {
	/* Number of blue nodes to drain in a batch. */
	// +kcc:proto:field=google.container.v1.NodePool.UpdateConfig.NodePool_UpdateConfig_BlueGreenSettings.StandardRolloutPolicy.batch_node_count
	BatchNodeCount *int `json:"batchNodeCount,omitempty"`

	/* Percentage of the bool pool nodes to drain in a batch. The range of this field should be (0.0, 1.0]. */
	// +kcc:proto:field=google.container.v1.NodePool.UpdateConfig.NodePool_UpdateConfig_BlueGreenSettings.StandardRolloutPolicy.batch_percentage
	BatchPercentage *float64 `json:"batchPercentage,omitempty"`

	/* Soak time after each batch gets drained.
	A duration in seconds with up to nine fractional digits, ending with 's'. Example: "3.5s". */
	// +kcc:proto:field=google.container.v1.NodePool.UpdateConfig.NodePool_UpdateConfig_BlueGreenSettings.StandardRolloutPolicy.batch_soak_duration
	BatchSoakDuration *string `json:"batchSoakDuration,omitempty"`
}

// +kcc:proto=google.container.v1.NodeTaint
type NodeTaint struct {
	// +required
	/* Effect for taint. */
	// +kcc:proto:field=google.container.v1.NodeTaint.effect
	Effect *string `json:"effect,omitempty"`

	// +required
	/* Key for taint. */
	// +kcc:proto:field=google.container.v1.NodeTaint.key
	Key *string `json:"key,omitempty"`

	// +required
	/* Value for taint. */
	// +kcc:proto:field=google.container.v1.NodeTaint.value
	Value *string `json:"value,omitempty"`
}

// +kcc:proto=google.container.v1.NodeManagement.UpgradeOptions
type NodeManagement_UpgradeOptions struct {
	/* This field is set when upgrades are about to commence with the approximate start time for the upgrades, in RFC3339 text format. */
	// +kcc:proto:field=google.container.v1.NodeManagement.NodeManagement_UpgradeOptions.auto_upgrade_start_time
	AutoUpgradeStartTime *string `json:"autoUpgradeStartTime,omitempty"`

	/* This field is set when upgrades are about to commence with the description of the upgrade. */
	// +kcc:proto:field=google.container.v1.NodeManagement.NodeManagement_UpgradeOptions.description
	Description *string `json:"description,omitempty"`
}

// +kcc:proto=google.container.v1.NodePool.UpgradeSettings
type NodePool_UpgradeSettings struct {
	/* Settings for blue-green upgrade strategy. */
	// +kcc:proto:field=google.container.v1.NodePool.NodePool_UpgradeSettings.blue_green_settings
	NodePool_UpdateConfig_BlueGreenSettings *NodePool_UpdateConfig_BlueGreenSettings `json:"blueGreenSettings,omitempty"`

	/* The maximum number of nodes that can be created beyond the current size of the node pool during the upgrade process. */
	// +kcc:proto:field=google.container.v1.NodePool.NodePool_UpgradeSettings.max_surge
	MaxSurge *int `json:"maxSurge,omitempty"`

	/* The maximum number of nodes that can be simultaneously unavailable during the upgrade process. */
	// +kcc:proto:field=google.container.v1.NodePool.NodePool_UpgradeSettings.max_unavailable
	MaxUnavailable *int `json:"maxUnavailable,omitempty"`

	/* Update strategy of the node pool. */
	// +kcc:proto:field=google.container.v1.NodePool.NodePool_UpgradeSettings.strategy
	Strategy *string `json:"strategy,omitempty"`
}

// +kcc:proto=google.container.v1.VerticalPodAutoscaling
type VerticalPodAutoscaling struct {
	/* Enables vertical pod autoscaling. */
	// +kcc:proto:field=google.container.v1.VerticalPodAutoscaling.enabled
	// +required
	Enabled *bool `json:"enabled,omitempty"`
}

// +kcc:proto=google.container.v1.SecurityPostureConfig.WorkloadConfig
type WorkloadConfig struct {
	/* Sets which mode of auditing should be used for the cluster's workloads. Accepted values are DISABLED, BASIC. */
	// +required
	// +kcc:proto:field=google.container.v1.SecurityPostureConfig.WorkloadConfig.audit_mode
	AuditMode *string `json:"auditMode,omitempty"`
}

// +kcc:proto=google.container.v1.WorkloadIdentityConfig
type WorkloadIdentityConfig struct {
	/* DEPRECATED. This field will be removed in a future major release as it has been deprecated in the API. Use `workloadPool` instead; `workloadPool` field will supersede this field.
	Enables workload identity. */
	// +kcc:proto:field=google.container.v1.WorkloadIdentityConfig.identity_namespace
	IdentityNamespace *string `json:"identityNamespace,omitempty"`

	/* The workload pool to attach all Kubernetes service accounts to. */
	// +kcc:proto:field=google.container.v1.WorkloadIdentityConfig.workload_pool
	WorkloadPool *string `json:"workloadPool,omitempty"`
}

// +kcc:proto=google.container.v1.NodeConfig.WorkloadMetadataConfig
type NodeConfig_WorkloadMetadataConfig struct {
	/* Mode is the configuration for how to expose metadata to workloads running on the node. */
	// +kcc:proto:field=google.container.v1.NodeConfig.WorkloadMetadataConfig.mode
	Mode *string `json:"mode,omitempty"`

	/* DEPRECATED. Deprecated in favor of mode. NodeMetadata is the configuration for how to expose metadata to the workloads running on the node. */
	// +kcc:proto:field=google.container.v1.NodeConfig.WorkloadMetadataConfig.node_metadata
	NodeMetadata *string `json:"nodeMetadata,omitempty"`
}

// +kcc:spec:proto=google.container.v1.Cluster
type ContainerClusterSpec struct {
	/* The configuration for addons supported by GKE. */
	// +kcc:proto:field=google.container.v1.Cluster.addons_config
	AddonsConfig *AddonsConfig `json:"addonsConfig,omitempty"`

	/* Enable NET_ADMIN for this cluster. */
	// +kcc:proto:field=google.container.v1.Cluster.allow_net_admin
	AllowNetAdmin *bool `json:"allowNetAdmin,omitempty"`

	/* Configuration for the Google Groups for GKE feature. */
	// +kcc:proto:field=google.container.v1.Cluster.authenticator_groups_config
	AuthenticatorGroupsConfig *AuthenticatorGroupsConfig `json:"authenticatorGroupsConfig,omitempty"`

	/* Configuration options for the Binary Authorization feature. */
	// +kcc:proto:field=google.container.v1.Cluster.binary_authorization
	BinaryAuthorization *BinaryAuthorization `json:"binaryAuthorization,omitempty"`

	/* Per-cluster configuration of Node Auto-Provisioning with Cluster Autoscaler to automatically adjust the size of the cluster and create/delete node pools based on the current needs of the cluster's workload. See the guide to using Node Auto-Provisioning for more details. */
	// +kcc:proto:field=google.container.v1.Cluster.autoscaling
	ClusterAutoscaling *ClusterAutoscaling `json:"clusterAutoscaling,omitempty"`

	/* Immutable. The IP address range of the Kubernetes pods in this cluster in CIDR notation (e.g. 10.96.0.0/14). Leave blank to have one automatically chosen or specify a /14 block in 10.0.0.0/8. This field will only work for routes-based clusters, where ip_allocation_policy is not defined. */
	// +kcc:proto:field=google.container.v1.Cluster.cluster_ipv4_cidr
	ClusterIPV4CIDR *string `json:"clusterIpv4Cidr,omitempty"`

	/* Telemetry integration for the cluster. */
	// +kcc:proto:field=google.container.v1.Cluster.cluster_telemetry
	ClusterTelemetry *ClusterTelemetry `json:"clusterTelemetry,omitempty"`

	/* Immutable. Configuration for the confidential nodes feature, which makes nodes run on confidential VMs. Warning: This configuration can't be changed (or added/removed) after cluster creation without deleting and recreating the entire cluster. */
	// +kcc:proto:field=google.container.v1.Cluster.confidential_nodes
	ConfidentialNodes *ConfidentialNodes `json:"confidentialNodes,omitempty"`

	/* Configuration for all of the cluster's control plane endpoints. Currently supports only DNS endpoint configuration and disable IP endpoint. Other IP endpoint configurations are available in private_cluster_config. */
	// +kcc:proto:field=google.container.v1.Cluster.control_plane_endpoints_config
	ControlPlaneEndpointsConfig *ControlPlaneEndpointsConfig `json:"controlPlaneEndpointsConfig,omitempty"`

	/* Cost management configuration for the cluster. */
	// +kcc:proto:field=google.container.v1.Cluster.cost_management_config
	CostManagementConfig *CostManagementConfig `json:"costManagementConfig,omitempty"`

	/* Application-layer Secrets Encryption settings. The object format is {state = string, key_name = string}. Valid values of state are: "ENCRYPTED"; "DECRYPTED". key_name is the name of a CloudKMS key. */
	// +kcc:proto:field=google.container.v1.Cluster.database_encryption
	DatabaseEncryption *DatabaseEncryption `json:"databaseEncryption,omitempty"`

	/* Immutable. The desired datapath provider for this cluster. By default, uses the IPTables-based kube-proxy implementation. */
	// +kcc:proto:field=google.container.v1.Cluster.datapath_provider
	DatapathProvider *string `json:"datapathProvider,omitempty"`

	/* Immutable. The default maximum number of pods per node in this cluster. This doesn't work on "routes-based" clusters, clusters that don't have IP Aliasing enabled. */
	// +kcc:proto:field=google.container.v1.Cluster.default_max_pods_constraint.max_pods_per_node
	DefaultMaxPodsPerNode *int `json:"defaultMaxPodsPerNode,omitempty"`

	/* Whether the cluster disables default in-node sNAT rules. In-node sNAT rules will be disabled when defaultSnatStatus is disabled. */
	// +kcc:proto:field=google.container.v1.Cluster.default_snat_status
	DefaultSnatStatus *DefaultSnatStatus `json:"defaultSnatStatus,omitempty"`

	/* Immutable.  Description of the cluster. */
	// +kcc:proto:field=google.container.v1.Cluster.description
	Description *string `json:"description,omitempty"`

	/* Immutable. Configuration for Cloud DNS for Kubernetes Engine. */
	// +kcc:proto:field=google.container.v1.Cluster.dns_config
	DnsConfig *DNSConfig `json:"dnsConfig,omitempty"`

	/* Immutable. Enable Autopilot for this cluster. */
	// +kcc:proto:field=google.container.v1.Cluster.autopilot.enabled
	EnableAutopilot *bool `json:"enableAutopilot,omitempty"`

	/* DEPRECATED. Deprecated in favor of binary_authorization. Enable Binary Authorization for this cluster. If enabled, all container images will be validated by Google Binary Authorization. */
	// +kcc:proto:field=google.container.v1.Cluster.enable_binary_authorization
	EnableBinaryAuthorization *bool `json:"enableBinaryAuthorization,omitempty"`

	/* Whether Cilium cluster-wide network policy is enabled on this cluster. */
	// +kcc:proto:field=google.container.v1.Cluster.network_config.enable_cilium_clusterwide_network_policy
	EnableCiliumClusterwideNetworkPolicy *bool `json:"enableCiliumClusterwideNetworkPolicy,omitempty"`

	/* Whether FQDN Network Policy is enabled on this cluster. */
	// +kcc:proto:field=google.container.v1.Cluster.network_config.enable_fqdn_network_policy
	EnableFQDNNetworkPolicy *bool `json:"enableFqdnNetworkPolicy,omitempty"`

	/* Whether Intra-node visibility is enabled for this cluster. This makes same node pod to pod traffic visible for VPC network. */
	// +kcc:proto:field=google.container.v1.Cluster.network_config.enable_intranode_visibility
	EnableIntranodeVisibility *bool `json:"enableIntranodeVisibility,omitempty"`

	/* Configuration for Kubernetes Beta APIs. */
	// +kcc:proto:field=google.container.v1.Cluster.enable_k8s_beta_apis
	EnableK8SBetaApis *K8SBetaAPIConfig `json:"enableK8sBetaApis,omitempty"`

	/* Immutable. Whether to enable Kubernetes Alpha features for this cluster. Note that when this option is enabled, the cluster cannot be upgraded and will be automatically deleted after 30 days. */
	// +kcc:proto:field=google.container.v1.Cluster.enable_kubernetes_alpha
	EnableKubernetesAlpha *bool `json:"enableKubernetesAlpha,omitempty"`

	/* Whether L4ILB Subsetting is enabled for this cluster. */
	// +kcc:proto:field=google.container.v1.Cluster.network_config.enable_l4_ilb_subsetting
	EnableL4ILBSubsetting *bool `json:"enableL4IlbSubsetting,omitempty"`

	/* Whether the ABAC authorizer is enabled for this cluster. When enabled, identities in the system, including service accounts, nodes, and controllers, will have statically granted permissions beyond those provided by the RBAC configuration or IAM. Defaults to false. */
	// +kcc:proto:field=google.container.v1.Cluster.legacy_abac.enabled
	EnableLegacyAbac *bool `json:"enableLegacyAbac,omitempty"`

	/* Immutable. Whether multi-networking is enabled for this cluster. */
	// +kcc:proto:field=google.container.v1.Cluster.network_config.enable_multi_networking
	EnableMultiNetworking *bool `json:"enableMultiNetworking,omitempty"`

	/* Enable Shielded Nodes features on all nodes in this cluster. Defaults to true. */
	// +kcc:proto:field=google.container.v1.Cluster.shielded_nodes.enabled
	EnableShieldedNodes *bool `json:"enableShieldedNodes,omitempty"`

	/* Immutable. Whether to enable Cloud TPU resources in this cluster. */
	// +kcc:proto:field=google.container.v1.Cluster.enable_tpu
	EnableTpu *bool `json:"enableTpu,omitempty"`

	/* Configuration for GKE Gateway API controller. */
	// +kcc:proto:field=google.container.v1.Cluster.gateway_api_config
	GatewayApiConfig *GatewayAPIConfig `json:"gatewayApiConfig,omitempty"`

	/* Configuration for Identity Service which allows customers to use external identity providers with the K8S API. */
	// +kcc:proto:field=google.container.v1.Cluster.identity_service_config
	IdentityServiceConfig *IdentityServiceConfig `json:"identityServiceConfig,omitempty"`

	/* Immutable. The number of nodes to create in this cluster's default node pool. In regional or multi-zonal clusters, this is the number of nodes per zone. Must be set if node_pool is not set. If you're using google_container_node_pool objects with no default node pool, you'll need to set this to a value of at least 1, alongside setting remove_default_node_pool to true. */
	// +kcc:proto:field=google.container.v1.Cluster.initial_node_count
	InitialNodeCount *int `json:"initialNodeCount,omitempty"`

	/* Immutable. Configuration of cluster IP allocation for VPC-native clusters. Adding this block enables IP aliasing, making the cluster VPC-native instead of routes-based. */
	// +kcc:proto:field=google.container.v1.Cluster.ip_allocation_policy
	IPAllocationPolicy *IPAllocationPolicy `json:"ipAllocationPolicy,omitempty"`

	/* Immutable. The location (region or zone) in which the cluster master will be created, as well as the default node location. If you specify a zone (such as us-central1-a), the cluster will be a zonal cluster with a single cluster master. If you specify a region (such as us-west1), the cluster will be a regional cluster with multiple masters spread across zones in the region, and with default node locations in those zones as well. */
	// +required
	Location *string `json:"location,omitempty"`

	/* Logging configuration for the cluster. */
	// +kcc:proto:field=google.container.v1.Cluster.logging_config
	LoggingConfig *LoggingConfig `json:"loggingConfig,omitempty"`

	/* The logging service that the cluster should write logs to. Available options include logging.googleapis.com(Legacy Stackdriver), logging.googleapis.com/kubernetes(Stackdriver Kubernetes Engine Logging), and none. Defaults to logging.googleapis.com/kubernetes. */
	// +kcc:proto:field=google.container.v1.Cluster.logging_service
	LoggingService *string `json:"loggingService,omitempty"`

	/* The maintenance policy to use for the cluster. */
	// +kcc:proto:field=google.container.v1.Cluster.maintenance_policy
	MaintenancePolicy *MaintenancePolicy `json:"maintenancePolicy,omitempty"`

	/* DEPRECATED. Basic authentication was removed for GKE cluster versions >= 1.19. The authentication information for accessing the Kubernetes master. Some values in this block are only returned by the API if your service account has permission to get credentials for your GKE cluster. If you see an unexpected diff unsetting your client cert, ensure you have the container.clusters.getCredentials permission. */
	// +kcc:proto:field=google.container.v1.Cluster.master_auth
	MasterAuth *MasterAuth `json:"masterAuth,omitempty"`

	/* The desired configuration options for master authorized networks. Omit the nested cidr_blocks attribute to disallow external access (except the cluster node IPs, which GKE automatically whitelists). */
	// +kcc:proto:field=google.container.v1.Cluster.master_authorized_networks_config
	MasterAuthorizedNetworksConfig *MasterAuthorizedNetworksConfig `json:"masterAuthorizedNetworksConfig,omitempty"`

	/* If set, and enable_certificates=true, the GKE Workload Identity Certificates controller and node agent will be deployed in the cluster. */
	// +kcc:proto:field=google.container.v1.Cluster.mesh_certificates
	MeshCertificates *MeshCertificates `json:"meshCertificates,omitempty"`

	/* The minimum version of the master. GKE will auto-update the master to new versions, so this does not guarantee the current master version--use the read-only master_version field to obtain that. If unset, the cluster's version will be set by GKE to the version of the most recent official release (which is not necessarily the latest version). */
	// +kcc:proto:field=google.container.v1.Cluster.min_master_version
	MinMasterVersion *string `json:"minMasterVersion,omitempty"`

	/* Monitoring configuration for the cluster. */
	// +kcc:proto:field=google.container.v1.Cluster.monitoring_config
	MonitoringConfig *MonitoringConfig `json:"monitoringConfig,omitempty"`

	/* The monitoring service that the cluster should write metrics to. Automatically send metrics from pods in the cluster to the Google Cloud Monitoring API. VM metrics will be collected by Google Compute Engine regardless of this setting Available options include monitoring.googleapis.com(Legacy Stackdriver), monitoring.googleapis.com/kubernetes(Stackdriver Kubernetes Engine Monitoring), and none. Defaults to monitoring.googleapis.com/kubernetes. */
	// +kcc:proto:field=google.container.v1.Cluster.monitoring_service
	MonitoringService *string `json:"monitoringService,omitempty"`

	/* Configuration options for the NetworkPolicy feature. */
	// +kcc:proto:field=google.container.v1.Cluster.network_policy
	NetworkPolicy *NetworkPolicy `json:"networkPolicy,omitempty"`

	// +kcc:proto:field=google.container.v1.Cluster.network
	NetworkRef *computev1beta1.ComputeNetworkRef `json:"networkRef,omitempty"`

	/* Immutable. Determines whether alias IPs or routes will be used for pod IPs in the cluster. */
	// +kcc:proto:field=google.container.v1.Cluster.networking_mode
	NetworkingMode *string `json:"networkingMode,omitempty"`

	/* Immutable. The configuration of the nodepool. */
	// +kcc:proto:field=google.container.v1.Cluster.node_config
	NodeConfig *NodeConfig `json:"nodeConfig,omitempty"`

	/* The list of zones in which the cluster's nodes are located. Nodes must be in the region of their regional cluster or in the same region as their cluster's zone for zonal clusters. If this is specified for a zonal cluster, omit the cluster's zone. */
	// +kcc:proto:field=google.container.v1.Cluster.locations
	NodeLocations []string `json:"nodeLocations,omitempty"`

	/* Node pool configs that apply to all auto-provisioned node pools in autopilot clusters and node auto-provisioning enabled clusters. */
	// +kcc:proto:field=google.container.v1.Cluster.node_pool_auto_config
	NodePoolAutoConfig *NodePoolAutoConfig `json:"nodePoolAutoConfig,omitempty"`

	/* The default nodel pool settings for the entire cluster. */
	// +kcc:proto:field=google.container.v1.Cluster.node_pool_defaults
	NodePoolDefaults *NodePoolDefaults `json:"nodePoolDefaults,omitempty"`

	// +kcc:proto:field=google.container.v1.Cluster.node_version
	NodeVersion *string `json:"nodeVersion,omitempty"`

	/* The notification config for sending cluster upgrade notifications. */
	// +kcc:proto:field=google.container.v1.Cluster.notification_config
	NotificationConfig *NotificationConfig `json:"notificationConfig,omitempty"`

	/* Configuration for the PodSecurityPolicy feature. */
	// +kcc:proto:field=google.container.v1.Cluster.pod_security_policy_config
	PodSecurityPolicyConfig *PodSecurityPolicyConfig `json:"podSecurityPolicyConfig,omitempty"`

	/* Configuration for private clusters, clusters with private nodes. */
	// +kcc:proto:field=google.container.v1.Cluster.private_cluster_config
	PrivateClusterConfig *PrivateClusterConfig `json:"privateClusterConfig,omitempty"`

	/* The desired state of IPv6 connectivity to Google Services. By default, no private IPv6 access to or from Google Services (all access will be via IPv4). */
	// +kcc:proto:field=google.container.v1.Cluster.private_ipv6_google_access
	PrivateIpv6GoogleAccess *string `json:"privateIpv6GoogleAccess,omitempty"`

	/* Enable/Disable Protect API features for the cluster. */
	ProtectConfig *ProtectConfig `json:"protectConfig,omitempty"`

	/* Configuration options for the Release channel feature, which provide more control over automatic upgrades of your GKE clusters. Note that removing this field from your config will not unenroll it. Instead, use the "UNSPECIFIED" channel. */
	// +kcc:proto:field=google.container.v1.Cluster.release_channel
	ReleaseChannel *ReleaseChannel `json:"releaseChannel,omitempty"`

	/* Immutable. Optional. The name of the resource. Used for creation and acquisition. When unset, the value of `metadata.name` is used as the default. */
	ResourceID *string `json:"resourceID,omitempty"`

	/* Configuration for the ResourceUsageExportConfig feature. */
	// +kcc:proto:field=google.container.v1.Cluster.resource_usage_export_config
	ResourceUsageExportConfig *ResourceUsageExportConfig `json:"resourceUsageExportConfig,omitempty"`

	/* Defines the config needed to enable/disable features for the Security Posture API. */
	// +kcc:proto:field=google.container.v1.Cluster.security_posture_config
	SecurityPostureConfig *SecurityPostureConfig `json:"securityPostureConfig,omitempty"`

	/* If set, and enabled=true, services with external ips field will not be blocked. */
	// +kcc:proto:field=google.container.v1.Cluster.service_external_ips_config
	ServiceExternalIPsConfig *ServiceExternalIPsConfig `json:"serviceExternalIpsConfig,omitempty"`

	// +kcc:proto:field=google.container.v1.Cluster.subnetwork
	SubnetworkRef *refsv1beta1.ComputeSubnetworkRef `json:"subnetworkRef,omitempty"`

	/* Vertical Pod Autoscaling automatically adjusts the resources of pods controlled by it. */
	// +kcc:proto:field=google.container.v1.Cluster.vertical_pod_autoscaling
	VerticalPodAutoscaling *VerticalPodAutoscaling `json:"verticalPodAutoscaling,omitempty"`

	/* Configuration for the use of Kubernetes Service Accounts in GCP IAM policies. */
	// +kcc:proto:field=google.container.v1.Cluster.workload_identity_config
	WorkloadIdentityConfig *WorkloadIdentityConfig `json:"workloadIdentityConfig,omitempty"`
}

// +kcc:proto=google.container.v1.ControlPlaneEndpointsConfig
type ControlPlaneEndpointsConfigStatus struct {
	/* DNS endpoint configuration. */
	// +kcc:proto:field=google.container.v1.ControlPlaneEndpointsConfig.dns_endpoint_config
	DNSEndpointConfig *DNSEndpointConfigStatus `json:"dnsEndpointConfig,omitempty"`
}

// +kcc:proto=google.container.v1.ControlPlaneEndpointsConfig.DNSEndpointConfig
type DNSEndpointConfigStatus struct {
	/* The cluster's DNS endpoint. */
	// +kcc:proto:field=google.container.v1.ControlPlaneEndpointsConfig.DNSEndpointConfig.endpoint
	Endpoint *string `json:"endpoint,omitempty"`
}

// +kcc:proto=google.container.v1.MasterAuth
type MasterAuthStatus struct {
	/* Base64 encoded public certificate used by clients to authenticate to the cluster endpoint. */
	// +kcc:proto:field=google.container.v1.MasterAuth.client_certificate
	ClientCertificate *string `json:"clientCertificate,omitempty"`

	/* Base64 encoded public certificate that is the root of trust for the cluster. */
	// +kcc:proto:field=google.container.v1.MasterAuth.cluster_ca_certificate
	ClusterCACertificate *string `json:"clusterCaCertificate,omitempty"`
}

// +kcc:observedstate:proto=google.container.v1.Cluster
type ClusterObservedState struct {
	/* Configuration for all of the cluster's control plane endpoints. Currently supports only DNS endpoint configuration and disable IP endpoint. Other IP endpoint configurations are available in private_cluster_config. */
	// +kcc:proto:field=google.container.v1.Cluster.control_plane_endpoints_config
	ControlPlaneEndpointsConfig *ControlPlaneEndpointsConfigStatus `json:"controlPlaneEndpointsConfig,omitempty"`

	/* DEPRECATED. Basic authentication was removed for GKE cluster versions >= 1.19. The authentication information for accessing the Kubernetes master. Some values in this block are only returned by the API if your service account has permission to get credentials for your GKE cluster. If you see an unexpected diff unsetting your client cert, ensure you have the container.clusters.getCredentials permission. */
	// +kcc:proto:field=google.container.v1.Cluster.master_auth
	MasterAuth *MasterAuthStatus `json:"masterAuth,omitempty"`

	/* Configuration for private clusters, clusters with private nodes. */
	// +kcc:proto:field=google.container.v1.Cluster.private_cluster_config
	PrivateClusterConfig *PrivateClusterConfigStatus `json:"privateClusterConfig,omitempty"`
}

// +kcc:proto=google.container.v1.PrivateClusterConfig
type PrivateClusterConfigStatus struct {
	/* The internal IP address of this cluster's master endpoint. */
	// +kcc:proto:field=google.container.v1.PrivateClusterConfig.private_endpoint
	PrivateEndpoint *string `json:"privateEndpoint,omitempty"`

	/* The external IP address of this cluster's master endpoint. */
	// +kcc:proto:field=google.container.v1.PrivateClusterConfig.public_endpoint
	PublicEndpoint *string `json:"publicEndpoint,omitempty"`
}

type ContainerClusterStatus struct {
	/* Conditions represent the latest available observations of the
	   ContainerCluster's current state. */
	Conditions []v1alpha1.Condition `json:"conditions,omitempty"`
	/* The IP address of this cluster's Kubernetes master. */
	Endpoint *string `json:"endpoint,omitempty"`

	/* The fingerprint of the set of labels for this cluster. */
	LabelFingerprint *string `json:"labelFingerprint,omitempty"`

	/* The current version of the master in the cluster. This may be different than the min_master_version set in the config if the master has been updated by GKE. */
	MasterVersion *string `json:"masterVersion,omitempty"`

	/* ObservedGeneration is the generation of the resource that was most recently observed by the Config Connector controller. If this is equal to metadata.generation, then that means that the current reported status reflects the most recent desired state of the resource. */
	ObservedGeneration *int64 `json:"observedGeneration,omitempty"`

	/* The observed state of the underlying GCP resource. */
	ObservedState *ClusterObservedState `json:"observedState,omitempty"`

	Operation *string `json:"operation,omitempty"`

	/* Server-defined URL for the resource. */
	SelfLink *string `json:"selfLink,omitempty"`

	/* The IP address range of the Kubernetes services in this cluster, in CIDR notation (e.g. 1.2.3.4/29). Service addresses are typically put in the last /16 from the container CIDR. */
	ServicesIPV4CIDR *string `json:"servicesIpv4Cidr,omitempty"`

	/* The IP address range of the Cloud TPUs in this cluster, in CIDR notation (e.g. 1.2.3.4/29). */
	TpuIPV4CIDRBlock *string `json:"tpuIpv4CidrBlock,omitempty"`
}

// +genclient
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:object:root=true
// +kubebuilder:resource:categories=gcp,shortName=gcpcontainercluster;gcpcontainerclusters
// +kubebuilder:subresource:status
// +kubebuilder:storageversion
// +kubebuilder:metadata:labels="cnrm.cloud.google.com/managed-by-kcc=true"
// +kubebuilder:metadata:labels="cnrm.cloud.google.com/system=true"
// +kubebuilder:metadata:labels="cnrm.cloud.google.com/stability-level=stable"
// +kubebuilder:metadata:labels="cnrm.cloud.google.com/tf2crd=true"
// +kubebuilder:printcolumn:name="Age",JSONPath=".metadata.creationTimestamp",type="date"
// +kubebuilder:printcolumn:name="Ready",JSONPath=".status.conditions[?(@.type=='Ready')].status",type="string",description="When 'True', the most recent reconcile of the resource succeeded"
// +kubebuilder:printcolumn:name="Status",JSONPath=".status.conditions[?(@.type=='Ready')].reason",type="string",description="The reason for the value in 'Ready'"
// +kubebuilder:printcolumn:name="Status Age",JSONPath=".status.conditions[?(@.type=='Ready')].lastTransitionTime",type="date",description="The last transition time for the value in 'Status'"

// ContainerCluster is the Schema for the ContainerCluster API
// +k8s:openapi-gen=true
type ContainerCluster struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   ContainerClusterSpec   `json:"spec"`
	Status ContainerClusterStatus `json:"status,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:object:root=true
// ContainerClusterList contains a list of ContainerCluster
type ContainerClusterList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []ContainerCluster `json:"items"`
}

func init() {
	SchemeBuilder.Register(&ContainerCluster{}, &ContainerClusterList{})
}

// +kcc:observedstate:proto=google.container.v1.DailyMaintenanceWindow
type DailyMaintenanceWindowObservedState struct {
	// +kcc:proto:field=google.container.v1.DailyMaintenanceWindow.duration
	Duration *string `json:"duration,omitempty"`
}

// +kcc:observedstate:proto=google.container.v1.NodeManagement
type NodeManagementObservedState struct {
	// +kcc:proto:field=google.container.v1.NodeManagement.upgrade_options
	UpgradeOptions *AutoUpgradeOptionsObservedState `json:"upgradeOptions,omitempty"`
}

// +kcc:observedstate:proto=google.container.v1.AutoUpgradeOptions
type AutoUpgradeOptionsObservedState struct {
	// +kcc:proto:field=google.container.v1.AutoUpgradeOptions.auto_upgrade_start_time
	AutoUpgradeStartTime *string `json:"autoUpgradeStartTime,omitempty"`
	// +kcc:proto:field=google.container.v1.AutoUpgradeOptions.description
	Description *string `json:"description,omitempty"`
}

// +kcc:observedstate:proto=google.container.v1.AdditionalPodRangesConfig
type AdditionalPodRangesConfigObservedState struct {
	// +kcc:proto:field=google.container.v1.AdditionalPodRangesConfig.pod_range_info
	PodRangeInfo []RangeInfo `json:"podRangeInfo,omitempty"`
}

// +kcc:observedstate:proto=google.container.v1.MaintenanceWindow
type MaintenanceWindowObservedState struct {
	// DailyMaintenanceWindow specifies a daily maintenance operation window.
	// +kcc:proto:field=google.container.v1.MaintenanceWindow.daily_maintenance_window
	DailyMaintenanceWindow *DailyMaintenanceWindowObservedState `json:"dailyMaintenanceWindow,omitempty"`
}

// +kcc:observedstate:proto=google.container.v1.MaintenancePolicy
type MaintenancePolicyObservedState struct {
	// Specifies the maintenance window in which maintenance may be performed.
	// +kcc:proto:field=google.container.v1.MaintenancePolicy.window
	Window *MaintenanceWindowObservedState `json:"window,omitempty"`
}
