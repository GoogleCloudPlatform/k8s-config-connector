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
	secretv1beta1 "github.com/GoogleCloudPlatform/k8s-config-connector/apis/refs/v1beta1/secret"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/apis/k8s/v1alpha1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// ContainerClusterSpec defines the desired state of ContainerCluster
// +kcc:spec:proto=google.container.v1.Cluster
type ContainerClusterSpec struct {
	// The location of this resource.
	// +kubebuilder:validation:Required
	Location string `json:"location"`

	// The ContainerCluster name. If not given, the metadata.name will be used.
	ResourceID *string `json:"resourceID,omitempty"`

	// The configuration for addons supported by GKE.
	AddonsConfig *ContainerClusterAddonsConfig `json:"addonsConfig,omitempty"`

	// Enable NET_ADMIN for this cluster.
	AllowNetAdmin *bool `json:"allowNetAdmin,omitempty"`

	// Configuration for the Google Groups for GKE feature.
	AuthenticatorGroupsConfig *ContainerClusterAuthenticatorGroupsConfig `json:"authenticatorGroupsConfig,omitempty"`

	// Configuration options for the Binary Authorization feature.
	BinaryAuthorization *ContainerClusterBinaryAuthorization `json:"binaryAuthorization,omitempty"`

	// Per-cluster configuration of Node Auto-Provisioning with Cluster Autoscaler
	// to automatically adjust the size of the cluster and create/delete node pools
	// based on the current needs of the cluster's workload.
	ClusterAutoscaling *ContainerClusterClusterAutoscaling `json:"clusterAutoscaling,omitempty"`

	// The IP address range of the Kubernetes pods in this cluster in CIDR notation.
	// Immutable.
	ClusterIpv4Cidr *string `json:"clusterIpv4Cidr,omitempty"`

	// Telemetry integration for the cluster.
	ClusterTelemetry *ContainerClusterClusterTelemetry `json:"clusterTelemetry,omitempty"`

	// Configuration for the confidential nodes feature. Immutable.
	ConfidentialNodes *ContainerClusterConfidentialNodes `json:"confidentialNodes,omitempty"`

	// Configuration for all of the cluster's control plane endpoints.
	ControlPlaneEndpointsConfig *ContainerClusterControlPlaneEndpointsConfig `json:"controlPlaneEndpointsConfig,omitempty"`

	// Cost management configuration for the cluster.
	CostManagementConfig *ContainerClusterCostManagementConfig `json:"costManagementConfig,omitempty"`

	// Application-layer Secrets Encryption settings.
	DatabaseEncryption *ContainerClusterDatabaseEncryption `json:"databaseEncryption,omitempty"`

	// The desired datapath provider for this cluster. Immutable.
	DatapathProvider *string `json:"datapathProvider,omitempty"`

	// The default maximum number of pods per node in this cluster. Immutable.
	DefaultMaxPodsPerNode *int32 `json:"defaultMaxPodsPerNode,omitempty"`

	// Whether the cluster disables default in-node sNAT rules.
	DefaultSnatStatus *ContainerClusterDefaultSnatStatus `json:"defaultSnatStatus,omitempty"`

	// Description of the cluster. Immutable.
	Description *string `json:"description,omitempty"`

	// Configuration for Cloud DNS for Kubernetes Engine. Immutable.
	DnsConfig *ContainerClusterDnsConfig `json:"dnsConfig,omitempty"`

	// Enable Autopilot for this cluster. Immutable.
	EnableAutopilot *bool `json:"enableAutopilot,omitempty"`

	// DEPRECATED. Deprecated in favor of binary_authorization.
	EnableBinaryAuthorization *bool `json:"enableBinaryAuthorization,omitempty"`

	// Whether Cilium cluster-wide network policy is enabled on this cluster.
	EnableCiliumClusterwideNetworkPolicy *bool `json:"enableCiliumClusterwideNetworkPolicy,omitempty"`

	// Whether FQDN Network Policy is enabled on this cluster.
	EnableFqdnNetworkPolicy *bool `json:"enableFqdnNetworkPolicy,omitempty"`

	// Whether Intra-node visibility is enabled for this cluster.
	EnableIntranodeVisibility *bool `json:"enableIntranodeVisibility,omitempty"`

	// Configuration for Kubernetes Beta APIs.
	EnableK8sBetaApis *ContainerClusterEnableK8sBetaApis `json:"enableK8sBetaApis,omitempty"`

	// Whether to enable Kubernetes Alpha features for this cluster. Immutable.
	EnableKubernetesAlpha *bool `json:"enableKubernetesAlpha,omitempty"`

	// Whether L4ILB Subsetting is enabled for this cluster.
	EnableL4IlbSubsetting *bool `json:"enableL4IlbSubsetting,omitempty"`

	// Whether the ABAC authorizer is enabled for this cluster.
	EnableLegacyAbac *bool `json:"enableLegacyAbac,omitempty"`

	// Whether multi-networking is enabled for this cluster. Immutable.
	EnableMultiNetworking *bool `json:"enableMultiNetworking,omitempty"`

	// Enable Shielded Nodes features on all nodes in this cluster.
	EnableShieldedNodes *bool `json:"enableShieldedNodes,omitempty"`

	// Whether to enable Cloud TPU resources in this cluster. Immutable.
	EnableTpu *bool `json:"enableTpu,omitempty"`

	// Configuration for GKE Gateway API controller.
	GatewayApiConfig *ContainerClusterGatewayApiConfig `json:"gatewayApiConfig,omitempty"`

	// Configuration for Identity Service.
	IdentityServiceConfig *ContainerClusterIdentityServiceConfig `json:"identityServiceConfig,omitempty"`

	// The number of nodes to create in this cluster's default node pool. Immutable.
	InitialNodeCount *int32 `json:"initialNodeCount,omitempty"`

	// Configuration of cluster IP allocation for VPC-native clusters. Immutable.
	IpAllocationPolicy *ContainerClusterIpAllocationPolicy `json:"ipAllocationPolicy,omitempty"`

	// Logging configuration for the cluster.
	LoggingConfig *ContainerClusterLoggingConfig `json:"loggingConfig,omitempty"`

	// The logging service that the cluster should write logs to.
	LoggingService *string `json:"loggingService,omitempty"`

	// The maintenance policy to use for the cluster.
	MaintenancePolicy *ContainerClusterMaintenancePolicy `json:"maintenancePolicy,omitempty"`

	// DEPRECATED. The authentication information for accessing the Kubernetes master.
	MasterAuth *ContainerClusterMasterAuth `json:"masterAuth,omitempty"`

	// The desired configuration options for master authorized networks.
	MasterAuthorizedNetworksConfig *ContainerClusterMasterAuthorizedNetworksConfig `json:"masterAuthorizedNetworksConfig,omitempty"`

	// If set, and enable_certificates=true, the GKE Workload Identity Certificates
	// controller and node agent will be deployed in the cluster.
	MeshCertificates *ContainerClusterMeshCertificates `json:"meshCertificates,omitempty"`

	// The minimum version of the master.
	MinMasterVersion *string `json:"minMasterVersion,omitempty"`

	// Monitoring configuration for the cluster.
	MonitoringConfig *ContainerClusterMonitoringConfig `json:"monitoringConfig,omitempty"`

	// The monitoring service that the cluster should write metrics to.
	MonitoringService *string `json:"monitoringService,omitempty"`

	// Configuration options for the NetworkPolicy feature.
	NetworkPolicy *ContainerClusterNetworkPolicy `json:"networkPolicy,omitempty"`

	// The network (VPC) to host the cluster in.
	NetworkRef *computev1beta1.ComputeNetworkRef `json:"networkRef,omitempty"`

	// Determines whether alias IPs or routes will be used for pod IPs in the cluster. Immutable.
	NetworkingMode *string `json:"networkingMode,omitempty"`

	// The configuration of the nodepool. Immutable.
	NodeConfig *ContainerClusterNodeConfig `json:"nodeConfig,omitempty"`

	// The list of zones in which the cluster's nodes are located.
	NodeLocations []string `json:"nodeLocations,omitempty"`

	// Node pool configs that apply to all auto-provisioned node pools.
	NodePoolAutoConfig *ContainerClusterNodePoolAutoConfig `json:"nodePoolAutoConfig,omitempty"`

	// The default node pool settings for the entire cluster.
	NodePoolDefaults *ContainerClusterNodePoolDefaults `json:"nodePoolDefaults,omitempty"`

	// The Kubernetes version of the cluster.
	NodeVersion *string `json:"nodeVersion,omitempty"`

	// The notification config for sending cluster upgrade notifications.
	NotificationConfig *ContainerClusterNotificationConfig `json:"notificationConfig,omitempty"`

	// Configuration for the PodSecurityPolicy feature.
	PodSecurityPolicyConfig *ContainerClusterPodSecurityPolicyConfig `json:"podSecurityPolicyConfig,omitempty"`

	// Configuration for private clusters, clusters with private nodes.
	PrivateClusterConfig *ContainerClusterPrivateClusterConfig `json:"privateClusterConfig,omitempty"`

	// The desired state of IPv6 connectivity to Google Services.
	PrivateIpv6GoogleAccess *string `json:"privateIpv6GoogleAccess,omitempty"`

	// Enable/Disable Protect API features for the cluster.
	ProtectConfig *ContainerClusterProtectConfig `json:"protectConfig,omitempty"`

	// Configuration options for the Release channel feature.
	ReleaseChannel *ContainerClusterReleaseChannel `json:"releaseChannel,omitempty"`

	// Configuration for the ResourceUsageExportConfig feature.
	ResourceUsageExportConfig *ContainerClusterResourceUsageExportConfig `json:"resourceUsageExportConfig,omitempty"`

	// Defines the config needed to enable/disable features for the Security Posture API.
	SecurityPostureConfig *ContainerClusterSecurityPostureConfig `json:"securityPostureConfig,omitempty"`

	// If set, and enabled=true, services with external ips field will not be blocked.
	ServiceExternalIpsConfig *ContainerClusterServiceExternalIpsConfig `json:"serviceExternalIpsConfig,omitempty"`

	// The subnetwork to host the cluster in.
	SubnetworkRef *refsv1beta1.ComputeSubnetworkRef `json:"subnetworkRef,omitempty"`

	// Vertical Pod Autoscaling automatically adjusts the resources of pods controlled by it.
	VerticalPodAutoscaling *ContainerClusterVerticalPodAutoscaling `json:"verticalPodAutoscaling,omitempty"`

	// Configuration for the use of Kubernetes Service Accounts in GCP IAM policies.
	WorkloadIdentityConfig *ContainerClusterWorkloadIdentityConfig `json:"workloadIdentityConfig,omitempty"`
}

// ContainerClusterAddonsConfig defines the configuration for addons supported by GKE.
type ContainerClusterAddonsConfig struct {
	// The status of the CloudRun addon.
	CloudrunConfig *ContainerClusterCloudrunConfig `json:"cloudrunConfig,omitempty"`

	// The of the Config Connector addon.
	ConfigConnectorConfig *ContainerClusterConfigConnectorConfig `json:"configConnectorConfig,omitempty"`

	// The status of the NodeLocal DNSCache addon.
	DnsCacheConfig *ContainerClusterDnsCacheConfig `json:"dnsCacheConfig,omitempty"`

	// Whether this cluster should enable the Google Compute Engine Persistent Disk CSI Driver.
	GcePersistentDiskCsiDriverConfig *ContainerClusterGcePersistentDiskCsiDriverConfig `json:"gcePersistentDiskCsiDriverConfig,omitempty"`

	// The status of the Filestore CSI driver addon.
	GcpFilestoreCsiDriverConfig *ContainerClusterGcpFilestoreCsiDriverConfig `json:"gcpFilestoreCsiDriverConfig,omitempty"`

	// The status of the GCS Fuse CSI driver addon.
	GcsFuseCsiDriverConfig *ContainerClusterGcsFuseCsiDriverConfig `json:"gcsFuseCsiDriverConfig,omitempty"`

	// The status of the Backup for GKE Agent addon.
	GkeBackupAgentConfig *ContainerClusterGkeBackupAgentConfig `json:"gkeBackupAgentConfig,omitempty"`

	// The status of the Horizontal Pod Autoscaling addon.
	HorizontalPodAutoscaling *ContainerClusterHorizontalPodAutoscaling `json:"horizontalPodAutoscaling,omitempty"`

	// The status of the HTTP (L7) load balancing controller addon.
	HttpLoadBalancing *ContainerClusterHttpLoadBalancing `json:"httpLoadBalancing,omitempty"`

	// The status of the Istio addon.
	IstioConfig *ContainerClusterIstioConfig `json:"istioConfig,omitempty"`

	// Configuration for the KALM addon.
	KalmConfig *ContainerClusterKalmConfig `json:"kalmConfig,omitempty"`

	// Whether we should enable the network policy addon for the master.
	NetworkPolicyConfig *ContainerClusterNetworkPolicyConfig `json:"networkPolicyConfig,omitempty"`
}

type ContainerClusterCloudrunConfig struct {
	// +kubebuilder:validation:Required
	Disabled         bool    `json:"disabled"`
	LoadBalancerType *string `json:"loadBalancerType,omitempty"`
}

type ContainerClusterConfigConnectorConfig struct {
	// +kubebuilder:validation:Required
	Enabled bool `json:"enabled"`
}

type ContainerClusterDnsCacheConfig struct {
	// +kubebuilder:validation:Required
	Enabled bool `json:"enabled"`
}

type ContainerClusterGcePersistentDiskCsiDriverConfig struct {
	// +kubebuilder:validation:Required
	Enabled bool `json:"enabled"`
}

type ContainerClusterGcpFilestoreCsiDriverConfig struct {
	// +kubebuilder:validation:Required
	Enabled bool `json:"enabled"`
}

type ContainerClusterGcsFuseCsiDriverConfig struct {
	// +kubebuilder:validation:Required
	Enabled bool `json:"enabled"`
}

type ContainerClusterGkeBackupAgentConfig struct {
	// +kubebuilder:validation:Required
	Enabled bool `json:"enabled"`
}

type ContainerClusterHorizontalPodAutoscaling struct {
	// +kubebuilder:validation:Required
	Disabled bool `json:"disabled"`
}

type ContainerClusterHttpLoadBalancing struct {
	// +kubebuilder:validation:Required
	Disabled bool `json:"disabled"`
}

type ContainerClusterIstioConfig struct {
	// The authentication type between services in Istio.
	Auth *string `json:"auth,omitempty"`
	// The status of the Istio addon.
	// +kubebuilder:validation:Required
	Disabled bool `json:"disabled"`
}

type ContainerClusterKalmConfig struct {
	// +kubebuilder:validation:Required
	Enabled bool `json:"enabled"`
}

type ContainerClusterNetworkPolicyConfig struct {
	// +kubebuilder:validation:Required
	Disabled bool `json:"disabled"`
}

// ContainerClusterAuthenticatorGroupsConfig is the configuration for Google Groups for GKE.
type ContainerClusterAuthenticatorGroupsConfig struct {
	// The name of the RBAC security group for use with Google security groups in Kubernetes RBAC.
	// +kubebuilder:validation:Required
	SecurityGroup string `json:"securityGroup"`
}

// ContainerClusterBinaryAuthorization defines Binary Authorization configuration.
type ContainerClusterBinaryAuthorization struct {
	// DEPRECATED. Enable Binary Authorization for this cluster.
	Enabled *bool `json:"enabled,omitempty"`
	// Mode of operation for Binary Authorization policy evaluation.
	EvaluationMode *string `json:"evaluationMode,omitempty"`
}

// ContainerClusterClusterAutoscaling defines per-cluster configuration of Node Auto-Provisioning.
type ContainerClusterClusterAutoscaling struct {
	// Contains defaults for a node pool created by NAP.
	AutoProvisioningDefaults *ContainerClusterAutoProvisioningDefaults `json:"autoProvisioningDefaults,omitempty"`
	// Configuration options for the Autoscaling profile feature.
	AutoscalingProfile *string `json:"autoscalingProfile,omitempty"`
	// Default compute class is a configuration for default compute class.
	DefaultComputeClassConfig *ContainerClusterDefaultComputeClassConfig `json:"defaultComputeClassConfig,omitempty"`
	// Whether node auto-provisioning is enabled.
	Enabled *bool `json:"enabled,omitempty"`
	// Global constraints for machine resources in the cluster.
	ResourceLimits []ContainerClusterResourceLimit `json:"resourceLimits,omitempty"`
}

type ContainerClusterAutoProvisioningDefaults struct {
	// The Customer Managed Encryption Key used to encrypt the boot disk attached to each node.
	BootDiskKMSKeyRef *refsv1beta1.KMSCryptoKeyRef `json:"bootDiskKMSKeyRef,omitempty"`
	// Size of the disk attached to each node, specified in GB.
	DiskSize *int32 `json:"diskSize,omitempty"`
	// The default image type used by NAP once a new node pool is being created.
	ImageType *string `json:"imageType,omitempty"`
	// NodeManagement configuration for this NodePool.
	Management *ContainerClusterNodeManagement `json:"management,omitempty"`
	// Minimum CPU platform to be used by this instance.
	MinCpuPlatform *string `json:"minCpuPlatform,omitempty"`
	// Scopes that are used by NAP when creating node pools.
	OauthScopes []string `json:"oauthScopes,omitempty"`
	// The Google Cloud Platform Service Account to be used by the node VMs.
	ServiceAccountRef *refsv1beta1.IAMServiceAccountRef `json:"serviceAccountRef,omitempty"`
	// Shielded Instance options.
	ShieldedInstanceConfig *ContainerClusterShieldedInstanceConfig `json:"shieldedInstanceConfig,omitempty"`
	// Specifies the upgrade settings for NAP created node pools.
	UpgradeSettings *ContainerClusterUpgradeSettings `json:"upgradeSettings,omitempty"`
}

type ContainerClusterNodeManagement struct {
	// Specifies whether the node auto-repair is enabled for the node pool.
	AutoRepair *bool `json:"autoRepair,omitempty"`
	// Specifies whether node auto-upgrade is enabled for the node pool.
	AutoUpgrade *bool `json:"autoUpgrade,omitempty"`
	// Specifies the Auto Upgrade knobs for the node pool.
	UpgradeOptions []ContainerClusterAutoUpgradeOptions `json:"upgradeOptions,omitempty"`
}

type ContainerClusterAutoUpgradeOptions struct {
	// This field is set when upgrades are about to commence.
	AutoUpgradeStartTime *string `json:"autoUpgradeStartTime,omitempty"`
	// This field is set when upgrades are about to commence with the description.
	Description *string `json:"description,omitempty"`
}

type ContainerClusterShieldedInstanceConfig struct {
	// Defines whether the instance has integrity monitoring enabled.
	EnableIntegrityMonitoring *bool `json:"enableIntegrityMonitoring,omitempty"`
	// Defines whether the instance has Secure Boot enabled.
	EnableSecureBoot *bool `json:"enableSecureBoot,omitempty"`
}

type ContainerClusterUpgradeSettings struct {
	// Settings for blue-green upgrade strategy.
	BlueGreenSettings *ContainerClusterBlueGreenSettings `json:"blueGreenSettings,omitempty"`
	// The maximum number of nodes that can be created beyond the current size.
	MaxSurge *int32 `json:"maxSurge,omitempty"`
	// The maximum number of nodes that can be simultaneously unavailable.
	MaxUnavailable *int32 `json:"maxUnavailable,omitempty"`
	// Update strategy of the node pool.
	Strategy *string `json:"strategy,omitempty"`
}

type ContainerClusterBlueGreenSettings struct {
	// Time needed after draining entire blue pool.
	NodePoolSoakDuration *string `json:"nodePoolSoakDuration,omitempty"`
	// Standard policy for the blue-green upgrade.
	StandardRolloutPolicy *ContainerClusterStandardRolloutPolicy `json:"standardRolloutPolicy,omitempty"`
}

type ContainerClusterStandardRolloutPolicy struct {
	// Number of blue nodes to drain in a batch.
	BatchNodeCount *int32 `json:"batchNodeCount,omitempty"`
	// Percentage of the bool pool nodes to drain in a batch.
	BatchPercentage *float64 `json:"batchPercentage,omitempty"`
	// Soak time after each batch gets drained.
	BatchSoakDuration *string `json:"batchSoakDuration,omitempty"`
}

type ContainerClusterDefaultComputeClassConfig struct {
	// Enables default compute class.
	// +kubebuilder:validation:Required
	Enabled bool `json:"enabled"`
}

type ContainerClusterResourceLimit struct {
	// Maximum amount of the resource in the cluster.
	Maximum *int32 `json:"maximum,omitempty"`
	// Minimum amount of the resource in the cluster.
	Minimum *int32 `json:"minimum,omitempty"`
	// The type of the resource.
	// +kubebuilder:validation:Required
	ResourceType string `json:"resourceType"`
}

// ContainerClusterClusterTelemetry defines telemetry integration for the cluster.
type ContainerClusterClusterTelemetry struct {
	// Type of the integration.
	// +kubebuilder:validation:Required
	Type string `json:"type"`
}

// ContainerClusterConfidentialNodes defines configuration for the confidential nodes feature.
type ContainerClusterConfidentialNodes struct {
	// Whether Confidential Nodes feature is enabled for all nodes in this cluster. Immutable.
	// +kubebuilder:validation:Required
	Enabled bool `json:"enabled"`
}

// ContainerClusterControlPlaneEndpointsConfig defines configuration for all of the cluster's control plane endpoints.
type ContainerClusterControlPlaneEndpointsConfig struct {
	// DNS endpoint configuration.
	DnsEndpointConfig *ContainerClusterDnsEndpointConfig `json:"dnsEndpointConfig,omitempty"`
	// IP endpoint configuration.
	IpEndpointsConfig *ContainerClusterIpEndpointsConfig `json:"ipEndpointsConfig,omitempty"`
}

type ContainerClusterDnsEndpointConfig struct {
	// Controls whether user traffic is allowed over this endpoint.
	AllowExternalTraffic *bool `json:"allowExternalTraffic,omitempty"`
	// Controls whether the k8s token auth is allowed via DNS.
	EnableK8sTokensViaDns *bool `json:"enableK8sTokensViaDns,omitempty"`
}

type ContainerClusterIpEndpointsConfig struct {
	// Controls whether to allow direct IP access.
	Enabled *bool `json:"enabled,omitempty"`
}

// ContainerClusterCostManagementConfig defines cost management configuration.
type ContainerClusterCostManagementConfig struct {
	// Whether to enable GKE cost allocation.
	// +kubebuilder:validation:Required
	Enabled bool `json:"enabled"`
}

// ContainerClusterDatabaseEncryption defines Application-layer Secrets Encryption settings.
type ContainerClusterDatabaseEncryption struct {
	// The key to use to encrypt/decrypt secrets.
	KeyName *string `json:"keyName,omitempty"`
	// ENCRYPTED or DECRYPTED.
	// +kubebuilder:validation:Required
	State string `json:"state"`
}

// ContainerClusterDefaultSnatStatus defines whether the cluster disables default in-node sNAT rules.
type ContainerClusterDefaultSnatStatus struct {
	// When disabled is set to false, default IP masquerade rules will be applied.
	// +kubebuilder:validation:Required
	Disabled bool `json:"disabled"`
}

// ContainerClusterDnsConfig defines configuration for Cloud DNS for Kubernetes Engine.
type ContainerClusterDnsConfig struct {
	// Which in-cluster DNS provider should be used.
	ClusterDns *string `json:"clusterDns,omitempty"`
	// The suffix used for all cluster service records.
	ClusterDnsDomain *string `json:"clusterDnsDomain,omitempty"`
	// The scope of access to cluster DNS records.
	ClusterDnsScope *string `json:"clusterDnsScope,omitempty"`
}

// ContainerClusterEnableK8sBetaApis defines configuration for Kubernetes Beta APIs.
type ContainerClusterEnableK8sBetaApis struct {
	// Enabled Kubernetes Beta APIs.
	// +kubebuilder:validation:Required
	EnabledApis []string `json:"enabledApis"`
}

// ContainerClusterGatewayApiConfig defines configuration for GKE Gateway API controller.
type ContainerClusterGatewayApiConfig struct {
	// The Gateway API release channel to use for Gateway API.
	// +kubebuilder:validation:Required
	Channel string `json:"channel"`
}

// ContainerClusterIdentityServiceConfig defines configuration for Identity Service.
type ContainerClusterIdentityServiceConfig struct {
	// Whether to enable the Identity Service component.
	Enabled *bool `json:"enabled,omitempty"`
}

// ContainerClusterIpAllocationPolicy defines configuration of cluster IP allocation.
type ContainerClusterIpAllocationPolicy struct {
	// AdditionalPodRangesConfig is the configuration for additional pod secondary ranges.
	AdditionalPodRangesConfig *ContainerClusterAdditionalPodRangesConfig `json:"additionalPodRangesConfig,omitempty"`
	// The IP address range for the cluster pod IPs. Immutable.
	ClusterIpv4CidrBlock *string `json:"clusterIpv4CidrBlock,omitempty"`
	// The name of the existing secondary range in the cluster's subnetwork to use for pod IP addresses. Immutable.
	ClusterSecondaryRangeName *string `json:"clusterSecondaryRangeName,omitempty"`
	// Configuration for cluster level pod cidr overprovision. Immutable.
	PodCidrOverprovisionConfig *ContainerClusterPodCidrOverprovisionConfig `json:"podCidrOverprovisionConfig,omitempty"`
	// The IP address range of the services IPs in this cluster. Immutable.
	ServicesIpv4CidrBlock *string `json:"servicesIpv4CidrBlock,omitempty"`
	// The name of the existing secondary range in the cluster's subnetwork to use for service ClusterIPs. Immutable.
	ServicesSecondaryRangeName *string `json:"servicesSecondaryRangeName,omitempty"`
	// The IP Stack type of the cluster. Immutable.
	StackType *string `json:"stackType,omitempty"`
}

type ContainerClusterAdditionalPodRangesConfig struct {
	// Name for pod secondary ipv4 range which has the actual range defined ahead.
	// +kubebuilder:validation:Required
	PodRangeNames []string `json:"podRangeNames"`
}

type ContainerClusterPodCidrOverprovisionConfig struct {
	// +kubebuilder:validation:Required
	Disabled bool `json:"disabled"`
}

// ContainerClusterLoggingConfig defines logging configuration for the cluster.
type ContainerClusterLoggingConfig struct {
	// GKE components exposing logs.
	// +kubebuilder:validation:Required
	EnableComponents []string `json:"enableComponents"`
}

// ContainerClusterMaintenancePolicy defines the maintenance policy to use for the cluster.
type ContainerClusterMaintenancePolicy struct {
	// Time window specified for daily maintenance operations.
	DailyMaintenanceWindow *ContainerClusterDailyMaintenanceWindow `json:"dailyMaintenanceWindow,omitempty"`
	// Exceptions to maintenance window.
	MaintenanceExclusion []ContainerClusterMaintenanceExclusion `json:"maintenanceExclusion,omitempty"`
	// Time window for recurring maintenance operations.
	RecurringWindow *ContainerClusterRecurringWindow `json:"recurringWindow,omitempty"`
}

type ContainerClusterDailyMaintenanceWindow struct {
	Duration *string `json:"duration,omitempty"`
	// +kubebuilder:validation:Required
	StartTime string `json:"startTime"`
}

type ContainerClusterMaintenanceExclusion struct {
	// +kubebuilder:validation:Required
	EndTime string `json:"endTime"`
	// +kubebuilder:validation:Required
	ExclusionName string `json:"exclusionName"`
	// Maintenance exclusion related options.
	ExclusionOptions *ContainerClusterExclusionOptions `json:"exclusionOptions,omitempty"`
	// +kubebuilder:validation:Required
	StartTime string `json:"startTime"`
}

type ContainerClusterExclusionOptions struct {
	// The scope of automatic upgrades to restrict in the exclusion window.
	// +kubebuilder:validation:Required
	Scope string `json:"scope"`
}

type ContainerClusterRecurringWindow struct {
	// +kubebuilder:validation:Required
	EndTime string `json:"endTime"`
	// +kubebuilder:validation:Required
	Recurrence string `json:"recurrence"`
	// +kubebuilder:validation:Required
	StartTime string `json:"startTime"`
}

// ContainerClusterMasterAuth defines the authentication information for accessing the Kubernetes master.
type ContainerClusterMasterAuth struct {
	// Base64 encoded public certificate used by clients to authenticate to the cluster endpoint.
	ClientCertificate *string `json:"clientCertificate,omitempty"`
	// Whether client certificate authorization is enabled for this cluster. Immutable.
	ClientCertificateConfig *ContainerClusterClientCertificateConfig `json:"clientCertificateConfig,omitempty"`
	// Base64 encoded private key used by clients to authenticate to the cluster endpoint.
	ClientKey *string `json:"clientKey,omitempty"`
	// Base64 encoded public certificate that is the root of trust for the cluster.
	ClusterCaCertificate *string `json:"clusterCaCertificate,omitempty"`
	// The password to use for HTTP basic authentication when accessing the Kubernetes master endpoint.
	Password *secretv1beta1.Legacy `json:"password,omitempty"`
	// The username to use for HTTP basic authentication when accessing the Kubernetes master endpoint.
	Username *string `json:"username,omitempty"`
}

type ContainerClusterClientCertificateConfig struct {
	// Whether client certificate authorization is enabled for this cluster. Immutable.
	// +kubebuilder:validation:Required
	IssueClientCertificate bool `json:"issueClientCertificate"`
}

// ContainerClusterMasterAuthorizedNetworksConfig defines desired configuration options for master authorized networks.
type ContainerClusterMasterAuthorizedNetworksConfig struct {
	// External networks that can access the Kubernetes cluster master through HTTPS.
	CidrBlocks []ContainerClusterCidrBlock `json:"cidrBlocks,omitempty"`
	// Whether master is accessbile via Google Compute Engine Public IP addresses.
	GcpPublicCidrsAccessEnabled *bool `json:"gcpPublicCidrsAccessEnabled,omitempty"`
}

type ContainerClusterCidrBlock struct {
	// External network that can access Kubernetes master through HTTPS.
	// +kubebuilder:validation:Required
	CidrBlock string `json:"cidrBlock"`
	// Field for users to identify CIDR blocks.
	DisplayName *string `json:"displayName,omitempty"`
}

// ContainerClusterMeshCertificates defines configuration for GKE Workload Identity Certificates.
type ContainerClusterMeshCertificates struct {
	// When enabled the GKE Workload Identity Certificates controller and node agent will be deployed.
	// +kubebuilder:validation:Required
	EnableCertificates bool `json:"enableCertificates"`
}

// ContainerClusterMonitoringConfig defines monitoring configuration for the cluster.
type ContainerClusterMonitoringConfig struct {
	// Configuration of Advanced Datapath Observability features.
	AdvancedDatapathObservabilityConfig []ContainerClusterAdvancedDatapathObservabilityConfig `json:"advancedDatapathObservabilityConfig,omitempty"`
	// GKE components exposing metrics.
	EnableComponents []string `json:"enableComponents,omitempty"`
	// Configuration for Google Cloud Managed Services for Prometheus.
	ManagedPrometheus *ContainerClusterManagedPrometheus `json:"managedPrometheus,omitempty"`
}

type ContainerClusterAdvancedDatapathObservabilityConfig struct {
	// Whether or not the advanced datapath metrics are enabled.
	// +kubebuilder:validation:Required
	EnableMetrics bool `json:"enableMetrics"`
	// Mode used to make Relay available.
	RelayMode *string `json:"relayMode,omitempty"`
}

type ContainerClusterManagedPrometheus struct {
	// Whether or not the managed collection is enabled.
	// +kubebuilder:validation:Required
	Enabled bool `json:"enabled"`
}

// ContainerClusterNetworkPolicy defines configuration options for the NetworkPolicy feature.
type ContainerClusterNetworkPolicy struct {
	// Whether network policy is enabled on the cluster.
	// +kubebuilder:validation:Required
	Enabled bool `json:"enabled"`
	// The selected network policy provider.
	Provider *string `json:"provider,omitempty"`
}

// ContainerClusterNodeConfig defines the configuration of the nodepool.
type ContainerClusterNodeConfig struct {
	// Specifies options for controlling advanced machine features. Immutable.
	AdvancedMachineFeatures *ContainerClusterAdvancedMachineFeatures `json:"advancedMachineFeatures,omitempty"`
	// The Customer Managed Encryption Key used to encrypt the boot disk. Immutable.
	BootDiskKMSCryptoKeyRef *refsv1beta1.KMSCryptoKeyRef `json:"bootDiskKMSCryptoKeyRef,omitempty"`
	// Configuration for the confidential nodes feature. Immutable.
	ConfidentialNodes *ContainerClusterNodeConfidentialNodes `json:"confidentialNodes,omitempty"`
	// Size of the disk attached to each node. Immutable.
	DiskSizeGb *int32 `json:"diskSizeGb,omitempty"`
	// Type of the disk attached to each node. Immutable.
	DiskType *string `json:"diskType,omitempty"`
	// Parameters for the ephemeral storage filesystem. Immutable.
	EphemeralStorageConfig *ContainerClusterEphemeralStorageConfig `json:"ephemeralStorageConfig,omitempty"`
	// Parameters for the ephemeral storage filesystem. Immutable.
	EphemeralStorageLocalSsdConfig *ContainerClusterEphemeralStorageLocalSsdConfig `json:"ephemeralStorageLocalSsdConfig,omitempty"`
	// Enable or disable NCCL Fast Socket in the node pool.
	FastSocket *ContainerClusterFastSocket `json:"fastSocket,omitempty"`
	// GCFS configuration for this node. Immutable.
	GcfsConfig *ContainerClusterGcfsConfig `json:"gcfsConfig,omitempty"`
	// List of the type and count of accelerator cards attached to the instance. Immutable.
	GuestAccelerator []ContainerClusterGuestAccelerator `json:"guestAccelerator,omitempty"`
	// Enable or disable gvnic in the node pool. Immutable.
	Gvnic *ContainerClusterGvnic `json:"gvnic,omitempty"`
	// The maintenance policy for the hosts on which the GKE VMs run on. Immutable.
	HostMaintenancePolicy *ContainerClusterHostMaintenancePolicy `json:"hostMaintenancePolicy,omitempty"`
	// The image type to use for this node.
	ImageType *string `json:"imageType,omitempty"`
	// Node kubelet configs.
	KubeletConfig *ContainerClusterKubeletConfig `json:"kubeletConfig,omitempty"`
	// The map of Kubernetes labels to be applied to each node. Immutable.
	Labels map[string]string `json:"labels,omitempty"`
	// Parameters that can be configured on Linux nodes.
	LinuxNodeConfig *ContainerClusterLinuxNodeConfig `json:"linuxNodeConfig,omitempty"`
	// Parameters for raw-block local NVMe SSDs. Immutable.
	LocalNvmeSsdBlockConfig *ContainerClusterLocalNvmeSsdBlockConfig `json:"localNvmeSsdBlockConfig,omitempty"`
	// The number of local SSD disks to be attached to the node. Immutable.
	LocalSsdCount *int32 `json:"localSsdCount,omitempty"`
	// Type of logging agent that is used as the default value for node pools in the cluster.
	LoggingVariant *string `json:"loggingVariant,omitempty"`
	// The name of a Google Compute Engine machine type. Immutable.
	MachineType *string `json:"machineType,omitempty"`
	// The metadata key/value pairs assigned to instances in the cluster. Immutable.
	Metadata map[string]string `json:"metadata,omitempty"`
	// Minimum CPU platform to be used by this instance. Immutable.
	MinCpuPlatform *string `json:"minCpuPlatform,omitempty"`
	// Setting this field will assign instances of this pool to run on the specified node group. Immutable.
	NodeGroupRef *ContainerClusterNodeGroupRef `json:"nodeGroupRef,omitempty"`
	// The set of Google API scopes to be made available on all of the node VMs. Immutable.
	OauthScopes []string `json:"oauthScopes,omitempty"`
	// Whether the nodes are created as preemptible VM instances. Immutable.
	Preemptible *bool `json:"preemptible,omitempty"`
	// The reservation affinity configuration for the node pool. Immutable.
	ReservationAffinity *ContainerClusterReservationAffinity `json:"reservationAffinity,omitempty"`
	// The GCE resource labels to be applied to the node pool.
	ResourceLabels map[string]string `json:"resourceLabels,omitempty"`
	// Sandbox configuration for this node. Immutable.
	SandboxConfig *ContainerClusterSandboxConfig `json:"sandboxConfig,omitempty"`
	// The Google Cloud Platform Service Account to be used by the node VMs.
	ServiceAccountRef *refsv1beta1.IAMServiceAccountRef `json:"serviceAccountRef,omitempty"`
	// Shielded Instance options. Immutable.
	ShieldedInstanceConfig *ContainerClusterShieldedInstanceConfig `json:"shieldedInstanceConfig,omitempty"`
	// Node affinity options for sole tenant node pools. Immutable.
	SoleTenantConfig *ContainerClusterSoleTenantConfig `json:"soleTenantConfig,omitempty"`
	// Whether the nodes are created as spot VM instances. Immutable.
	Spot *bool `json:"spot,omitempty"`
	// The list of instance tags applied to all nodes.
	Tags []string `json:"tags,omitempty"`
	// List of Kubernetes taints to be applied to each node.
	Taint []ContainerClusterNodeTaint `json:"taint,omitempty"`
	// The workload metadata configuration for this node. Immutable.
	WorkloadMetadataConfig *ContainerClusterWorkloadMetadataConfig `json:"workloadMetadataConfig,omitempty"`
}

type ContainerClusterAdvancedMachineFeatures struct {
	// Whether or not to enable nested virtualization. Immutable.
	EnableNestedVirtualization *bool `json:"enableNestedVirtualization,omitempty"`
	// The number of threads per physical core. Immutable.
	ThreadsPerCore *int32 `json:"threadsPerCore,omitempty"`
}

type ContainerClusterNodeConfidentialNodes struct {
	// Whether Confidential Nodes feature is enabled for all nodes in this pool. Immutable.
	// +kubebuilder:validation:Required
	Enabled bool `json:"enabled"`
}

type ContainerClusterEphemeralStorageConfig struct {
	// Number of local SSDs to use to back ephemeral storage. Immutable.
	// +kubebuilder:validation:Required
	LocalSsdCount int32 `json:"localSsdCount"`
}

type ContainerClusterEphemeralStorageLocalSsdConfig struct {
	// Number of local SSDs to use to back ephemeral storage. Immutable.
	// +kubebuilder:validation:Required
	LocalSsdCount int32 `json:"localSsdCount"`
}

type ContainerClusterFastSocket struct {
	// Whether or not NCCL Fast Socket is enabled.
	// +kubebuilder:validation:Required
	Enabled bool `json:"enabled"`
}

type ContainerClusterGcfsConfig struct {
	// Whether or not GCFS is enabled. Immutable.
	// +kubebuilder:validation:Required
	Enabled bool `json:"enabled"`
}

type ContainerClusterGuestAccelerator struct {
	// The number of the accelerator cards exposed to an instance. Immutable.
	// +kubebuilder:validation:Required
	Count int32 `json:"count"`
	// Configuration for auto installation of GPU driver. Immutable.
	GpuDriverInstallationConfig *ContainerClusterGpuDriverInstallationConfig `json:"gpuDriverInstallationConfig,omitempty"`
	// Size of partitions to create on the GPU. Immutable.
	GpuPartitionSize *string `json:"gpuPartitionSize,omitempty"`
	// Configuration for GPU sharing. Immutable.
	GpuSharingConfig *ContainerClusterGpuSharingConfig `json:"gpuSharingConfig,omitempty"`
	// The accelerator type resource name. Immutable.
	// +kubebuilder:validation:Required
	Type string `json:"type"`
}

type ContainerClusterGpuDriverInstallationConfig struct {
	// Mode for how the GPU driver is installed. Immutable.
	// +kubebuilder:validation:Required
	GpuDriverVersion string `json:"gpuDriverVersion"`
}

type ContainerClusterGpuSharingConfig struct {
	// The type of GPU sharing strategy to enable on the GPU node. Immutable.
	// +kubebuilder:validation:Required
	GpuSharingStrategy string `json:"gpuSharingStrategy"`
	// The maximum number of containers that can share a GPU. Immutable.
	// +kubebuilder:validation:Required
	MaxSharedClientsPerGpu int32 `json:"maxSharedClientsPerGpu"`
}

type ContainerClusterGvnic struct {
	// Whether or not gvnic is enabled. Immutable.
	// +kubebuilder:validation:Required
	Enabled bool `json:"enabled"`
}

type ContainerClusterHostMaintenancePolicy struct {
	// Immutable.
	// +kubebuilder:validation:Required
	MaintenanceInterval string `json:"maintenanceInterval"`
}

type ContainerClusterKubeletConfig struct {
	// Enable CPU CFS quota enforcement for containers that specify CPU limits.
	CpuCfsQuota *bool `json:"cpuCfsQuota,omitempty"`
	// Set the CPU CFS quota period value 'cpu.cfs_period_us'.
	CpuCfsQuotaPeriod *string `json:"cpuCfsQuotaPeriod,omitempty"`
	// Control the CPU management policy on the node.
	// +kubebuilder:validation:Required
	CpuManagerPolicy string `json:"cpuManagerPolicy"`
	// Controls the maximum number of processes allowed to run in a pod.
	PodPidsLimit *int32 `json:"podPidsLimit,omitempty"`
}

type ContainerClusterLinuxNodeConfig struct {
	// cgroupMode specifies the cgroup mode to be used on the node.
	CgroupMode *string `json:"cgroupMode,omitempty"`
	// The Linux kernel parameters to be applied to the nodes and all pods running on the nodes.
	Sysctls map[string]string `json:"sysctls,omitempty"`
}

type ContainerClusterLocalNvmeSsdBlockConfig struct {
	// Number of raw-block local NVMe SSD disks to be attached to the node. Immutable.
	// +kubebuilder:validation:Required
	LocalSsdCount int32 `json:"localSsdCount"`
}

// ContainerClusterNodeGroupRef defines a reference to a ComputeNodeGroup resource.
type ContainerClusterNodeGroupRef struct {
	// Allowed value: The `name` field of a `ComputeNodeGroup` resource.
	External string `json:"external,omitempty"`
	// Name of the referent.
	Name string `json:"name,omitempty"`
	// Namespace of the referent.
	Namespace string `json:"namespace,omitempty"`
}

type ContainerClusterReservationAffinity struct {
	// Corresponds to the type of reservation consumption. Immutable.
	// +kubebuilder:validation:Required
	ConsumeReservationType string `json:"consumeReservationType"`
	// The label key of a reservation resource. Immutable.
	Key *string `json:"key,omitempty"`
	// The label values of the reservation resource. Immutable.
	Values []string `json:"values,omitempty"`
}

type ContainerClusterSandboxConfig struct {
	// Type of the sandbox to use for the node. Immutable.
	// +kubebuilder:validation:Required
	SandboxType string `json:"sandboxType"`
}

type ContainerClusterSoleTenantConfig struct {
	// Immutable.
	// +kubebuilder:validation:Required
	NodeAffinity []ContainerClusterNodeAffinity `json:"nodeAffinity"`
}

type ContainerClusterNodeAffinity struct {
	// Immutable.
	// +kubebuilder:validation:Required
	Key string `json:"key"`
	// Immutable.
	// +kubebuilder:validation:Required
	Operator string `json:"operator"`
	// Immutable.
	// +kubebuilder:validation:Required
	Values []string `json:"values"`
}

type ContainerClusterNodeTaint struct {
	// Effect for taint.
	// +kubebuilder:validation:Required
	Effect string `json:"effect"`
	// Key for taint.
	// +kubebuilder:validation:Required
	Key string `json:"key"`
	// Value for taint.
	// +kubebuilder:validation:Required
	Value string `json:"value"`
}

type ContainerClusterWorkloadMetadataConfig struct {
	// Mode is the configuration for how to expose metadata to workloads running on the node.
	Mode *string `json:"mode,omitempty"`
	// DEPRECATED. Deprecated in favor of mode.
	NodeMetadata *string `json:"nodeMetadata,omitempty"`
}

// ContainerClusterNodePoolAutoConfig defines node pool configs that apply to all auto-provisioned node pools.
type ContainerClusterNodePoolAutoConfig struct {
	// Collection of Compute Engine network tags that can be applied to a node's underlying VM instance.
	NetworkTags *ContainerClusterNetworkTags `json:"networkTags,omitempty"`
}

type ContainerClusterNetworkTags struct {
	// List of network tags applied to auto-provisioned node pools.
	Tags []string `json:"tags,omitempty"`
}

// ContainerClusterNodePoolDefaults defines the default node pool settings for the entire cluster.
type ContainerClusterNodePoolDefaults struct {
	// Subset of NodeConfig message that has defaults.
	NodeConfigDefaults *ContainerClusterNodeConfigDefaults `json:"nodeConfigDefaults,omitempty"`
}

type ContainerClusterNodeConfigDefaults struct {
	// GCFS configuration for this node.
	GcfsConfig *ContainerClusterGcfsConfigDefault `json:"gcfsConfig,omitempty"`
	// Type of logging agent that is used as the default value for node pools in the cluster.
	LoggingVariant *string `json:"loggingVariant,omitempty"`
}

type ContainerClusterGcfsConfigDefault struct {
	// Whether or not GCFS is enabled.
	// +kubebuilder:validation:Required
	Enabled bool `json:"enabled"`
}

// ContainerClusterNotificationConfig defines the notification config for sending cluster upgrade notifications.
type ContainerClusterNotificationConfig struct {
	// Notification config for Cloud Pub/Sub.
	// +kubebuilder:validation:Required
	Pubsub ContainerClusterPubSub `json:"pubsub"`
}

type ContainerClusterPubSub struct {
	// Whether or not the notification config is enabled.
	// +kubebuilder:validation:Required
	Enabled bool `json:"enabled"`
	// Allows filtering to one or more specific event types.
	Filter *ContainerClusterPubSubFilter `json:"filter,omitempty"`
	// The PubSubTopic to send the notification to.
	TopicRef *ContainerClusterTopicRef `json:"topicRef,omitempty"`
}

type ContainerClusterPubSubFilter struct {
	// Can be used to filter what notifications are sent.
	// +kubebuilder:validation:Required
	EventType []string `json:"eventType"`
}

// ContainerClusterTopicRef defines a reference to a PubSubTopic resource.
type ContainerClusterTopicRef struct {
	// Allowed value: string of the format `projects/{{project}}/topics/{{value}}`.
	External string `json:"external,omitempty"`
	// Name of the referent.
	Name string `json:"name,omitempty"`
	// Namespace of the referent.
	Namespace string `json:"namespace,omitempty"`
}

// ContainerClusterPodSecurityPolicyConfig defines configuration for the PodSecurityPolicy feature.
type ContainerClusterPodSecurityPolicyConfig struct {
	// Enable the PodSecurityPolicy controller for this cluster.
	// +kubebuilder:validation:Required
	Enabled bool `json:"enabled"`
}

// ContainerClusterPrivateClusterConfig defines configuration for private clusters.
type ContainerClusterPrivateClusterConfig struct {
	// When true, the cluster's private endpoint is used as the cluster endpoint.
	EnablePrivateEndpoint *bool `json:"enablePrivateEndpoint,omitempty"`
	// Enables the private cluster feature, creating a private endpoint on the cluster.
	EnablePrivateNodes *bool `json:"enablePrivateNodes,omitempty"`
	// Controls cluster master global access settings.
	MasterGlobalAccessConfig *ContainerClusterMasterGlobalAccessConfig `json:"masterGlobalAccessConfig,omitempty"`
	// The IP range in CIDR notation to use for the hosted master network. Immutable.
	MasterIpv4CidrBlock *string `json:"masterIpv4CidrBlock,omitempty"`
	// The name of the peering between this cluster and the Google owned VPC.
	PeeringName *string `json:"peeringName,omitempty"`
	// The internal IP address of this cluster's master endpoint.
	PrivateEndpoint *string `json:"privateEndpoint,omitempty"`
	// Subnetwork in cluster's network where master's endpoint will be provisioned. Immutable.
	PrivateEndpointSubnetworkRef *refsv1beta1.ComputeSubnetworkRef `json:"privateEndpointSubnetworkRef,omitempty"`
	// The external IP address of this cluster's master endpoint.
	PublicEndpoint *string `json:"publicEndpoint,omitempty"`
}

type ContainerClusterMasterGlobalAccessConfig struct {
	// Whether the cluster master is accessible globally or not.
	// +kubebuilder:validation:Required
	Enabled bool `json:"enabled"`
}

// ContainerClusterProtectConfig defines Enable/Disable Protect API features for the cluster.
type ContainerClusterProtectConfig struct {
	// WorkloadConfig defines which actions are enabled for a cluster's workload configurations.
	WorkloadConfig *ContainerClusterWorkloadConfig `json:"workloadConfig,omitempty"`
	// Sets which mode to use for Protect workload vulnerability scanning feature.
	WorkloadVulnerabilityMode *string `json:"workloadVulnerabilityMode,omitempty"`
}

type ContainerClusterWorkloadConfig struct {
	// Sets which mode of auditing should be used for the cluster's workloads.
	// +kubebuilder:validation:Required
	AuditMode string `json:"auditMode"`
}

// ContainerClusterReleaseChannel defines configuration options for the Release channel feature.
type ContainerClusterReleaseChannel struct {
	// The selected release channel.
	// +kubebuilder:validation:Required
	Channel string `json:"channel"`
}

// ContainerClusterResourceUsageExportConfig defines configuration for the ResourceUsageExportConfig feature.
type ContainerClusterResourceUsageExportConfig struct {
	// Parameters for using BigQuery as the destination of resource usage export.
	// +kubebuilder:validation:Required
	BigqueryDestination ContainerClusterBigqueryDestination `json:"bigqueryDestination"`
	// Whether to enable network egress metering for this cluster.
	EnableNetworkEgressMetering *bool `json:"enableNetworkEgressMetering,omitempty"`
	// Whether to enable resource consumption metering on this cluster.
	EnableResourceConsumptionMetering *bool `json:"enableResourceConsumptionMetering,omitempty"`
}

type ContainerClusterBigqueryDestination struct {
	// The ID of a BigQuery Dataset.
	// +kubebuilder:validation:Required
	DatasetId string `json:"datasetId"`
}

// ContainerClusterSecurityPostureConfig defines the config needed to enable/disable features for the Security Posture API.
type ContainerClusterSecurityPostureConfig struct {
	// Sets the mode of the Kubernetes security posture API's off-cluster features.
	Mode *string `json:"mode,omitempty"`
	// Sets the mode of the Kubernetes security posture API's workload vulnerability scanning.
	VulnerabilityMode *string `json:"vulnerabilityMode,omitempty"`
}

// ContainerClusterServiceExternalIpsConfig defines service external IPs configuration.
type ContainerClusterServiceExternalIpsConfig struct {
	// When enabled, services with exterenal ips specified will be allowed.
	// +kubebuilder:validation:Required
	Enabled bool `json:"enabled"`
}

// ContainerClusterVerticalPodAutoscaling defines Vertical Pod Autoscaling configuration.
type ContainerClusterVerticalPodAutoscaling struct {
	// Enables vertical pod autoscaling.
	// +kubebuilder:validation:Required
	Enabled bool `json:"enabled"`
}

// ContainerClusterWorkloadIdentityConfig defines configuration for the use of Kubernetes Service Accounts in GCP IAM policies.
type ContainerClusterWorkloadIdentityConfig struct {
	// DEPRECATED. Use workloadPool instead.
	IdentityNamespace *string `json:"identityNamespace,omitempty"`
	// The workload pool to attach all Kubernetes service accounts to.
	WorkloadPool *string `json:"workloadPool,omitempty"`
}

// ContainerClusterStatus defines the config connector machine state of ContainerCluster
type ContainerClusterStatus struct {
	/* Conditions represent the latest available observations of the
	   object's current state. */
	Conditions []v1alpha1.Condition `json:"conditions,omitempty"`

	// ObservedGeneration is the generation of the resource that was most recently observed by the Config Connector controller. If this is equal to metadata.generation, then that means that the current reported status reflects the most recent desired state of the resource.
	ObservedGeneration *int64 `json:"observedGeneration,omitempty"`

	// A unique specifier for the ContainerCluster resource in GCP.
	ExternalRef *string `json:"externalRef,omitempty"`

	// ObservedState is the state of the resource as most recently observed in GCP.
	ObservedState *ContainerClusterObservedState `json:"observedState,omitempty"`

	// The IP address of this cluster's master endpoint.
	Endpoint *string `json:"endpoint,omitempty"`

	// The fingerprint of the set of labels for this cluster.
	LabelFingerprint *string `json:"labelFingerprint,omitempty"`

	// The current software version of the master endpoint.
	MasterVersion *string `json:"masterVersion,omitempty"`

	// The name of the Google Compute Engine operation associated with this cluster operation (if any).
	Operation *string `json:"operation,omitempty"`

	// Server-defined URL for the resource.
	SelfLink *string `json:"selfLink,omitempty"`

	// The IP address range of the Kubernetes services in this cluster, in CIDR notation.
	ServicesIpv4Cidr *string `json:"servicesIpv4Cidr,omitempty"`

	// The IP address range of the Cloud TPUs in this cluster, in CIDR notation.
	TpuIpv4CidrBlock *string `json:"tpuIpv4CidrBlock,omitempty"`
}

// ContainerClusterObservedState is the state of the ContainerCluster resource as most recently observed in GCP.
// +kcc:observedstate:proto=google.container.v1.Cluster
type ContainerClusterObservedState struct {
	// Configuration for all of the cluster's control plane endpoints.
	ControlPlaneEndpointsConfig *ContainerClusterControlPlaneEndpointsConfigObservedState `json:"controlPlaneEndpointsConfig,omitempty"`

	// DEPRECATED. The authentication information for accessing the Kubernetes master.
	MasterAuth *ContainerClusterMasterAuthObservedState `json:"masterAuth,omitempty"`

	// Configuration for private clusters, clusters with private nodes.
	PrivateClusterConfig *ContainerClusterPrivateClusterConfigObservedState `json:"privateClusterConfig,omitempty"`
}

type ContainerClusterControlPlaneEndpointsConfigObservedState struct {
	// DNS endpoint configuration.
	DnsEndpointConfig *ContainerClusterDnsEndpointConfigObservedState `json:"dnsEndpointConfig,omitempty"`
}

type ContainerClusterDnsEndpointConfigObservedState struct {
	// The cluster's DNS endpoint.
	Endpoint *string `json:"endpoint,omitempty"`
}

type ContainerClusterMasterAuthObservedState struct {
	// Base64 encoded public certificate used by clients to authenticate to the cluster endpoint.
	ClientCertificate *string `json:"clientCertificate,omitempty"`
	// Base64 encoded public certificate that is the root of trust for the cluster.
	ClusterCaCertificate *string `json:"clusterCaCertificate,omitempty"`
}

type ContainerClusterPrivateClusterConfigObservedState struct {
	// The internal IP address of this cluster's master endpoint.
	PrivateEndpoint *string `json:"privateEndpoint,omitempty"`
	// The external IP address of this cluster's master endpoint.
	PublicEndpoint *string `json:"publicEndpoint,omitempty"`
}

// +genclient
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:resource:categories=gcp,shortName=gcpcontainercluster;gcpcontainerclusters
// +kubebuilder:subresource:status
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

	// +required
	Spec   ContainerClusterSpec   `json:"spec,omitempty"`
	Status ContainerClusterStatus `json:"status,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// ContainerClusterList contains a list of ContainerCluster
type ContainerClusterList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []ContainerCluster `json:"items"`
}

func init() {
	SchemeBuilder.Register(&ContainerCluster{}, &ContainerClusterList{})
}
