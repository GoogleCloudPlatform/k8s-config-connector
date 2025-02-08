// Copyright 2025 Google LLC
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

package v1alpha1


// +kcc:proto=google.cloud.gkemulticloud.v1.AzureAuthorization
type AzureAuthorization struct {
	// Optional. Users that can perform operations as a cluster admin. A managed
	//  ClusterRoleBinding will be created to grant the `cluster-admin` ClusterRole
	//  to the users. Up to ten admin users can be provided.
	//
	//  For more info on RBAC, see
	//  https://kubernetes.io/docs/reference/access-authn-authz/rbac/#user-facing-roles
	// +kcc:proto:field=google.cloud.gkemulticloud.v1.AzureAuthorization.admin_users
	AdminUsers []AzureClusterUser `json:"adminUsers,omitempty"`

	// Optional. Groups of users that can perform operations as a cluster admin. A
	//  managed ClusterRoleBinding will be created to grant the `cluster-admin`
	//  ClusterRole to the groups. Up to ten admin groups can be provided.
	//
	//  For more info on RBAC, see
	//  https://kubernetes.io/docs/reference/access-authn-authz/rbac/#user-facing-roles
	// +kcc:proto:field=google.cloud.gkemulticloud.v1.AzureAuthorization.admin_groups
	AdminGroups []AzureClusterGroup `json:"adminGroups,omitempty"`
}

// +kcc:proto=google.cloud.gkemulticloud.v1.AzureCluster
type AzureCluster struct {
	// The name of this resource.
	//
	//  Cluster names are formatted as
	//  `projects/<project-number>/locations/<region>/azureClusters/<cluster-id>`.
	//
	//  See [Resource Names](https://cloud.google.com/apis/design/resource_names)
	//  for more details on Google Cloud Platform resource names.
	// +kcc:proto:field=google.cloud.gkemulticloud.v1.AzureCluster.name
	Name *string `json:"name,omitempty"`

	// Optional. A human readable description of this cluster.
	//  Cannot be longer than 255 UTF-8 encoded bytes.
	// +kcc:proto:field=google.cloud.gkemulticloud.v1.AzureCluster.description
	Description *string `json:"description,omitempty"`

	// Required. The Azure region where the cluster runs.
	//
	//  Each Google Cloud region supports a subset of nearby Azure regions.
	//  You can call
	//  [GetAzureServerConfig][google.cloud.gkemulticloud.v1.AzureClusters.GetAzureServerConfig]
	//  to list all supported Azure regions within a given Google Cloud region.
	// +kcc:proto:field=google.cloud.gkemulticloud.v1.AzureCluster.azure_region
	AzureRegion *string `json:"azureRegion,omitempty"`

	// Required. The ARM ID of the resource group where the cluster resources are
	//  deployed. For example:
	//  `/subscriptions/<subscription-id>/resourceGroups/<resource-group-name>`
	// +kcc:proto:field=google.cloud.gkemulticloud.v1.AzureCluster.resource_group_id
	ResourceGroupID *string `json:"resourceGroupID,omitempty"`

	// Optional. Name of the
	//  [AzureClient][google.cloud.gkemulticloud.v1.AzureClient] that contains
	//  authentication configuration for how the Anthos Multi-Cloud API connects to
	//  Azure APIs.
	//
	//  Either azure_client or azure_services_authentication should be provided.
	//
	//  The `AzureClient` resource must reside on the same Google Cloud Platform
	//  project and region as the `AzureCluster`.
	//
	//  `AzureClient` names are formatted as
	//  `projects/<project-number>/locations/<region>/azureClients/<client-id>`.
	//
	//  See [Resource Names](https://cloud.google.com/apis/design/resource_names)
	//  for more details on Google Cloud resource names.
	// +kcc:proto:field=google.cloud.gkemulticloud.v1.AzureCluster.azure_client
	AzureClient *string `json:"azureClient,omitempty"`

	// Required. Cluster-wide networking configuration.
	// +kcc:proto:field=google.cloud.gkemulticloud.v1.AzureCluster.networking
	Networking *AzureClusterNetworking `json:"networking,omitempty"`

	// Required. Configuration related to the cluster control plane.
	// +kcc:proto:field=google.cloud.gkemulticloud.v1.AzureCluster.control_plane
	ControlPlane *AzureControlPlane `json:"controlPlane,omitempty"`

	// Required. Configuration related to the cluster RBAC settings.
	// +kcc:proto:field=google.cloud.gkemulticloud.v1.AzureCluster.authorization
	Authorization *AzureAuthorization `json:"authorization,omitempty"`

	// Optional. Authentication configuration for management of Azure resources.
	//
	//  Either azure_client or azure_services_authentication should be provided.
	// +kcc:proto:field=google.cloud.gkemulticloud.v1.AzureCluster.azure_services_authentication
	AzureServicesAuthentication *AzureServicesAuthentication `json:"azureServicesAuthentication,omitempty"`

	// Allows clients to perform consistent read-modify-writes
	//  through optimistic concurrency control.
	//
	//  Can be sent on update and delete requests to ensure the
	//  client has an up-to-date value before proceeding.
	// +kcc:proto:field=google.cloud.gkemulticloud.v1.AzureCluster.etag
	Etag *string `json:"etag,omitempty"`

	// Optional. Annotations on the cluster.
	//
	//  This field has the same restrictions as Kubernetes annotations.
	//  The total size of all keys and values combined is limited to 256k.
	//  Keys can have 2 segments: prefix (optional) and name (required),
	//  separated by a slash (/).
	//  Prefix must be a DNS subdomain.
	//  Name must be 63 characters or less, begin and end with alphanumerics,
	//  with dashes (-), underscores (_), dots (.), and alphanumerics between.
	// +kcc:proto:field=google.cloud.gkemulticloud.v1.AzureCluster.annotations
	Annotations map[string]string `json:"annotations,omitempty"`

	// Required. Fleet configuration.
	// +kcc:proto:field=google.cloud.gkemulticloud.v1.AzureCluster.fleet
	Fleet *Fleet `json:"fleet,omitempty"`

	// Optional. Logging configuration for this cluster.
	// +kcc:proto:field=google.cloud.gkemulticloud.v1.AzureCluster.logging_config
	LoggingConfig *LoggingConfig `json:"loggingConfig,omitempty"`

	// Optional. Monitoring configuration for this cluster.
	// +kcc:proto:field=google.cloud.gkemulticloud.v1.AzureCluster.monitoring_config
	MonitoringConfig *MonitoringConfig `json:"monitoringConfig,omitempty"`
}

// +kcc:proto=google.cloud.gkemulticloud.v1.AzureClusterError
type AzureClusterError struct {
	// Human-friendly description of the error.
	// +kcc:proto:field=google.cloud.gkemulticloud.v1.AzureClusterError.message
	Message *string `json:"message,omitempty"`
}

// +kcc:proto=google.cloud.gkemulticloud.v1.AzureClusterGroup
type AzureClusterGroup struct {
	// Required. The name of the group, e.g. `my-group@domain.com`.
	// +kcc:proto:field=google.cloud.gkemulticloud.v1.AzureClusterGroup.group
	Group *string `json:"group,omitempty"`
}

// +kcc:proto=google.cloud.gkemulticloud.v1.AzureClusterNetworking
type AzureClusterNetworking struct {
	// Required. The Azure Resource Manager (ARM) ID of the VNet associated with
	//  your cluster.
	//
	//  All components in the cluster (i.e. control plane and node pools) run on a
	//  single VNet.
	//
	//  Example:
	//  `/subscriptions/<subscription-id>/resourceGroups/<resource-group-id>/providers/Microsoft.Network/virtualNetworks/<vnet-id>`
	//
	//  This field cannot be changed after creation.
	// +kcc:proto:field=google.cloud.gkemulticloud.v1.AzureClusterNetworking.virtual_network_id
	VirtualNetworkID *string `json:"virtualNetworkID,omitempty"`

	// Required. The IP address range of the pods in this cluster, in CIDR
	//  notation (e.g. `10.96.0.0/14`).
	//
	//  All pods in the cluster get assigned a unique IPv4 address from these
	//  ranges. Only a single range is supported.
	//
	//  This field cannot be changed after creation.
	// +kcc:proto:field=google.cloud.gkemulticloud.v1.AzureClusterNetworking.pod_address_cidr_blocks
	PodAddressCidrBlocks []string `json:"podAddressCidrBlocks,omitempty"`

	// Required. The IP address range for services in this cluster, in CIDR
	//  notation (e.g. `10.96.0.0/14`).
	//
	//  All services in the cluster get assigned a unique IPv4 address from these
	//  ranges. Only a single range is supported.
	//
	//  This field cannot be changed after creating a cluster.
	// +kcc:proto:field=google.cloud.gkemulticloud.v1.AzureClusterNetworking.service_address_cidr_blocks
	ServiceAddressCidrBlocks []string `json:"serviceAddressCidrBlocks,omitempty"`

	// Optional. The ARM ID of the subnet where Kubernetes private service type
	//  load balancers are deployed. When unspecified, it defaults to
	//  AzureControlPlane.subnet_id.
	//
	//  Example:
	//  "/subscriptions/d00494d6-6f3c-4280-bbb2-899e163d1d30/resourceGroups/anthos_cluster_gkeust4/providers/Microsoft.Network/virtualNetworks/gke-vnet-gkeust4/subnets/subnetid456"
	// +kcc:proto:field=google.cloud.gkemulticloud.v1.AzureClusterNetworking.service_load_balancer_subnet_id
	ServiceLoadBalancerSubnetID *string `json:"serviceLoadBalancerSubnetID,omitempty"`
}

// +kcc:proto=google.cloud.gkemulticloud.v1.AzureClusterResources
type AzureClusterResources struct {
}

// +kcc:proto=google.cloud.gkemulticloud.v1.AzureClusterUser
type AzureClusterUser struct {
	// Required. The name of the user, e.g. `my-gcp-id@gmail.com`.
	// +kcc:proto:field=google.cloud.gkemulticloud.v1.AzureClusterUser.username
	Username *string `json:"username,omitempty"`
}

// +kcc:proto=google.cloud.gkemulticloud.v1.AzureConfigEncryption
type AzureConfigEncryption struct {
	// Required. The ARM ID of the Azure Key Vault key to encrypt / decrypt config
	//  data.
	//
	//  For example:
	//  `/subscriptions/<subscription-id>/resourceGroups/<resource-group-id>/providers/Microsoft.KeyVault/vaults/<key-vault-id>/keys/<key-name>`
	// +kcc:proto:field=google.cloud.gkemulticloud.v1.AzureConfigEncryption.key_id
	KeyID *string `json:"keyID,omitempty"`

	// Optional. RSA key of the Azure Key Vault public key to use for encrypting
	//  the data.
	//
	//  This key must be formatted as a PEM-encoded SubjectPublicKeyInfo (RFC 5280)
	//  in ASN.1 DER form. The string must be comprised of a single PEM block of
	//  type "PUBLIC KEY".
	// +kcc:proto:field=google.cloud.gkemulticloud.v1.AzureConfigEncryption.public_key
	PublicKey *string `json:"publicKey,omitempty"`
}

// +kcc:proto=google.cloud.gkemulticloud.v1.AzureControlPlane
type AzureControlPlane struct {
	// Required. The Kubernetes version to run on control plane replicas
	//  (e.g. `1.19.10-gke.1000`).
	//
	//  You can list all supported versions on a given Google Cloud region by
	//  calling
	//  [GetAzureServerConfig][google.cloud.gkemulticloud.v1.AzureClusters.GetAzureServerConfig].
	// +kcc:proto:field=google.cloud.gkemulticloud.v1.AzureControlPlane.version
	Version *string `json:"version,omitempty"`

	// Optional. The ARM ID of the default subnet for the control plane. The
	//  control plane VMs are deployed in this subnet, unless
	//  `AzureControlPlane.replica_placements` is specified. This subnet will also
	//  be used as default for `AzureControlPlane.endpoint_subnet_id` if
	//  `AzureControlPlane.endpoint_subnet_id` is not specified. Similarly it will
	//  be used as default for
	//  `AzureClusterNetworking.service_load_balancer_subnet_id`.
	//
	//  Example:
	//  `/subscriptions/<subscription-id>/resourceGroups/<resource-group-id>/providers/Microsoft.Network/virtualNetworks/<vnet-id>/subnets/default`.
	// +kcc:proto:field=google.cloud.gkemulticloud.v1.AzureControlPlane.subnet_id
	SubnetID *string `json:"subnetID,omitempty"`

	// Optional. The Azure VM size name. Example: `Standard_DS2_v2`.
	//
	//  For available VM sizes, see
	//  https://docs.microsoft.com/en-us/azure/virtual-machines/vm-naming-conventions.
	//
	//  When unspecified, it defaults to `Standard_DS2_v2`.
	// +kcc:proto:field=google.cloud.gkemulticloud.v1.AzureControlPlane.vm_size
	VmSize *string `json:"vmSize,omitempty"`

	// Required. SSH configuration for how to access the underlying control plane
	//  machines.
	// +kcc:proto:field=google.cloud.gkemulticloud.v1.AzureControlPlane.ssh_config
	SSHConfig *AzureSshConfig `json:"sshConfig,omitempty"`

	// Optional. Configuration related to the root volume provisioned for each
	//  control plane replica.
	//
	//  When unspecified, it defaults to 32-GiB Azure Disk.
	// +kcc:proto:field=google.cloud.gkemulticloud.v1.AzureControlPlane.root_volume
	RootVolume *AzureDiskTemplate `json:"rootVolume,omitempty"`

	// Optional. Configuration related to the main volume provisioned for each
	//  control plane replica.
	//  The main volume is in charge of storing all of the cluster's etcd state.
	//
	//  When unspecified, it defaults to a 8-GiB Azure Disk.
	// +kcc:proto:field=google.cloud.gkemulticloud.v1.AzureControlPlane.main_volume
	MainVolume *AzureDiskTemplate `json:"mainVolume,omitempty"`

	// Optional. Configuration related to application-layer secrets encryption.
	// +kcc:proto:field=google.cloud.gkemulticloud.v1.AzureControlPlane.database_encryption
	DatabaseEncryption *AzureDatabaseEncryption `json:"databaseEncryption,omitempty"`

	// Optional. Proxy configuration for outbound HTTP(S) traffic.
	// +kcc:proto:field=google.cloud.gkemulticloud.v1.AzureControlPlane.proxy_config
	ProxyConfig *AzureProxyConfig `json:"proxyConfig,omitempty"`

	// Optional. Configuration related to vm config encryption.
	// +kcc:proto:field=google.cloud.gkemulticloud.v1.AzureControlPlane.config_encryption
	ConfigEncryption *AzureConfigEncryption `json:"configEncryption,omitempty"`

	// Optional. A set of tags to apply to all underlying control plane Azure
	//  resources.
	// +kcc:proto:field=google.cloud.gkemulticloud.v1.AzureControlPlane.tags
	Tags map[string]string `json:"tags,omitempty"`

	// Optional. Configuration for where to place the control plane replicas.
	//
	//  Up to three replica placement instances can be specified. If
	//  replica_placements is set, the replica placement instances will be applied
	//  to the three control plane replicas as evenly as possible.
	// +kcc:proto:field=google.cloud.gkemulticloud.v1.AzureControlPlane.replica_placements
	ReplicaPlacements []ReplicaPlacement `json:"replicaPlacements,omitempty"`

	// Optional. The ARM ID of the subnet where the control plane load balancer is
	//  deployed. When unspecified, it defaults to AzureControlPlane.subnet_id.
	//
	//  Example:
	//  "/subscriptions/d00494d6-6f3c-4280-bbb2-899e163d1d30/resourceGroups/anthos_cluster_gkeust4/providers/Microsoft.Network/virtualNetworks/gke-vnet-gkeust4/subnets/subnetid123"
	// +kcc:proto:field=google.cloud.gkemulticloud.v1.AzureControlPlane.endpoint_subnet_id
	EndpointSubnetID *string `json:"endpointSubnetID,omitempty"`
}

// +kcc:proto=google.cloud.gkemulticloud.v1.AzureDatabaseEncryption
type AzureDatabaseEncryption struct {
	// Required. The ARM ID of the Azure Key Vault key to encrypt / decrypt data.
	//
	//  For example:
	//  `/subscriptions/<subscription-id>/resourceGroups/<resource-group-id>/providers/Microsoft.KeyVault/vaults/<key-vault-id>/keys/<key-name>`
	//  Encryption will always take the latest version of the key and hence
	//  specific version is not supported.
	// +kcc:proto:field=google.cloud.gkemulticloud.v1.AzureDatabaseEncryption.key_id
	KeyID *string `json:"keyID,omitempty"`
}

// +kcc:proto=google.cloud.gkemulticloud.v1.AzureDiskTemplate
type AzureDiskTemplate struct {
	// Optional. The size of the disk, in GiBs.
	//
	//  When unspecified, a default value is provided. See the specific reference
	//  in the parent resource.
	// +kcc:proto:field=google.cloud.gkemulticloud.v1.AzureDiskTemplate.size_gib
	SizeGib *int32 `json:"sizeGib,omitempty"`
}

// +kcc:proto=google.cloud.gkemulticloud.v1.AzureProxyConfig
type AzureProxyConfig struct {
	// The ARM ID the of the resource group containing proxy keyvault.
	//
	//  Resource group ids are formatted as
	//  `/subscriptions/<subscription-id>/resourceGroups/<resource-group-name>`.
	// +kcc:proto:field=google.cloud.gkemulticloud.v1.AzureProxyConfig.resource_group_id
	ResourceGroupID *string `json:"resourceGroupID,omitempty"`

	// The URL the of the proxy setting secret with its version.
	//
	//  The secret must be a JSON encoded proxy configuration
	//  as described in
	//  https://cloud.google.com/kubernetes-engine/multi-cloud/docs/azure/how-to/use-a-proxy#create_a_proxy_configuration_file
	//
	//  Secret ids are formatted as
	//  `https://<key-vault-name>.vault.azure.net/secrets/<secret-name>/<secret-version>`.
	// +kcc:proto:field=google.cloud.gkemulticloud.v1.AzureProxyConfig.secret_id
	SecretID *string `json:"secretID,omitempty"`
}

// +kcc:proto=google.cloud.gkemulticloud.v1.AzureServicesAuthentication
type AzureServicesAuthentication struct {
	// Required. The Azure Active Directory Tenant ID.
	// +kcc:proto:field=google.cloud.gkemulticloud.v1.AzureServicesAuthentication.tenant_id
	TenantID *string `json:"tenantID,omitempty"`

	// Required. The Azure Active Directory Application ID.
	// +kcc:proto:field=google.cloud.gkemulticloud.v1.AzureServicesAuthentication.application_id
	ApplicationID *string `json:"applicationID,omitempty"`
}

// +kcc:proto=google.cloud.gkemulticloud.v1.AzureSshConfig
type AzureSshConfig struct {
	// Required. The SSH public key data for VMs managed by Anthos. This accepts
	//  the authorized_keys file format used in OpenSSH according to the sshd(8)
	//  manual page.
	// +kcc:proto:field=google.cloud.gkemulticloud.v1.AzureSshConfig.authorized_key
	AuthorizedKey *string `json:"authorizedKey,omitempty"`
}

// +kcc:proto=google.cloud.gkemulticloud.v1.CloudMonitoringConfig
type CloudMonitoringConfig struct {
	// Enable GKE-native logging and metrics.
	//  Only for Attached Clusters.
	// +kcc:proto:field=google.cloud.gkemulticloud.v1.CloudMonitoringConfig.enabled
	Enabled *bool `json:"enabled,omitempty"`
}

// +kcc:proto=google.cloud.gkemulticloud.v1.Fleet
type Fleet struct {
	// Required. The name of the Fleet host project where this cluster will be
	//  registered.
	//
	//  Project names are formatted as
	//  `projects/<project-number>`.
	// +kcc:proto:field=google.cloud.gkemulticloud.v1.Fleet.project
	Project *string `json:"project,omitempty"`
}

// +kcc:proto=google.cloud.gkemulticloud.v1.LoggingComponentConfig
type LoggingComponentConfig struct {
	// The components to be enabled.
	// +kcc:proto:field=google.cloud.gkemulticloud.v1.LoggingComponentConfig.enable_components
	EnableComponents []string `json:"enableComponents,omitempty"`
}

// +kcc:proto=google.cloud.gkemulticloud.v1.LoggingConfig
type LoggingConfig struct {
	// The configuration of the logging components;
	// +kcc:proto:field=google.cloud.gkemulticloud.v1.LoggingConfig.component_config
	ComponentConfig *LoggingComponentConfig `json:"componentConfig,omitempty"`
}

// +kcc:proto=google.cloud.gkemulticloud.v1.ManagedPrometheusConfig
type ManagedPrometheusConfig struct {
	// Enable Managed Collection.
	// +kcc:proto:field=google.cloud.gkemulticloud.v1.ManagedPrometheusConfig.enabled
	Enabled *bool `json:"enabled,omitempty"`
}

// +kcc:proto=google.cloud.gkemulticloud.v1.MonitoringConfig
type MonitoringConfig struct {
	// Enable Google Cloud Managed Service for Prometheus in the cluster.
	// +kcc:proto:field=google.cloud.gkemulticloud.v1.MonitoringConfig.managed_prometheus_config
	ManagedPrometheusConfig *ManagedPrometheusConfig `json:"managedPrometheusConfig,omitempty"`

	// Optionally enable GKE metrics.
	//  Only for Attached Clusters.
	// +kcc:proto:field=google.cloud.gkemulticloud.v1.MonitoringConfig.cloud_monitoring_config
	CloudMonitoringConfig *CloudMonitoringConfig `json:"cloudMonitoringConfig,omitempty"`
}

// +kcc:proto=google.cloud.gkemulticloud.v1.ReplicaPlacement
type ReplicaPlacement struct {
	// Required. For a given replica, the ARM ID of the subnet where the control
	//  plane VM is deployed. Make sure it's a subnet under the virtual network in
	//  the cluster configuration.
	// +kcc:proto:field=google.cloud.gkemulticloud.v1.ReplicaPlacement.subnet_id
	SubnetID *string `json:"subnetID,omitempty"`

	// Required. For a given replica, the Azure availability zone where to
	//  provision the control plane VM and the ETCD disk.
	// +kcc:proto:field=google.cloud.gkemulticloud.v1.ReplicaPlacement.azure_availability_zone
	AzureAvailabilityZone *string `json:"azureAvailabilityZone,omitempty"`
}

// +kcc:proto=google.cloud.gkemulticloud.v1.WorkloadIdentityConfig
type WorkloadIdentityConfig struct {
	// The OIDC issuer URL for this cluster.
	// +kcc:proto:field=google.cloud.gkemulticloud.v1.WorkloadIdentityConfig.issuer_uri
	IssuerURI *string `json:"issuerURI,omitempty"`

	// The Workload Identity Pool associated to the cluster.
	// +kcc:proto:field=google.cloud.gkemulticloud.v1.WorkloadIdentityConfig.workload_pool
	WorkloadPool *string `json:"workloadPool,omitempty"`

	// The ID of the OIDC Identity Provider (IdP) associated to the Workload
	//  Identity Pool.
	// +kcc:proto:field=google.cloud.gkemulticloud.v1.WorkloadIdentityConfig.identity_provider
	IdentityProvider *string `json:"identityProvider,omitempty"`
}

// +kcc:proto=google.cloud.gkemulticloud.v1.AzureCluster
type AzureClusterObservedState struct {
	// Output only. The current state of the cluster.
	// +kcc:proto:field=google.cloud.gkemulticloud.v1.AzureCluster.state
	State *string `json:"state,omitempty"`

	// Output only. The endpoint of the cluster's API server.
	// +kcc:proto:field=google.cloud.gkemulticloud.v1.AzureCluster.endpoint
	Endpoint *string `json:"endpoint,omitempty"`

	// Output only. A globally unique identifier for the cluster.
	// +kcc:proto:field=google.cloud.gkemulticloud.v1.AzureCluster.uid
	Uid *string `json:"uid,omitempty"`

	// Output only. If set, there are currently changes in flight to the cluster.
	// +kcc:proto:field=google.cloud.gkemulticloud.v1.AzureCluster.reconciling
	Reconciling *bool `json:"reconciling,omitempty"`

	// Output only. The time at which this cluster was created.
	// +kcc:proto:field=google.cloud.gkemulticloud.v1.AzureCluster.create_time
	CreateTime *string `json:"createTime,omitempty"`

	// Output only. The time at which this cluster was last updated.
	// +kcc:proto:field=google.cloud.gkemulticloud.v1.AzureCluster.update_time
	UpdateTime *string `json:"updateTime,omitempty"`

	// Output only. Workload Identity settings.
	// +kcc:proto:field=google.cloud.gkemulticloud.v1.AzureCluster.workload_identity_config
	WorkloadIdentityConfig *WorkloadIdentityConfig `json:"workloadIdentityConfig,omitempty"`

	// Output only. PEM encoded x509 certificate of the cluster root of trust.
	// +kcc:proto:field=google.cloud.gkemulticloud.v1.AzureCluster.cluster_ca_certificate
	ClusterCaCertificate *string `json:"clusterCaCertificate,omitempty"`

	// Required. Fleet configuration.
	// +kcc:proto:field=google.cloud.gkemulticloud.v1.AzureCluster.fleet
	Fleet *FleetObservedState `json:"fleet,omitempty"`

	// Output only. Managed Azure resources for this cluster.
	// +kcc:proto:field=google.cloud.gkemulticloud.v1.AzureCluster.managed_resources
	ManagedResources *AzureClusterResources `json:"managedResources,omitempty"`

	// Output only. A set of errors found in the cluster.
	// +kcc:proto:field=google.cloud.gkemulticloud.v1.AzureCluster.errors
	Errors []AzureClusterError `json:"errors,omitempty"`
}

// +kcc:proto=google.cloud.gkemulticloud.v1.AzureClusterResources
type AzureClusterResourcesObservedState struct {
	// Output only. The ARM ID of the cluster network security group.
	// +kcc:proto:field=google.cloud.gkemulticloud.v1.AzureClusterResources.network_security_group_id
	NetworkSecurityGroupID *string `json:"networkSecurityGroupID,omitempty"`

	// Output only. The ARM ID of the control plane application security group.
	// +kcc:proto:field=google.cloud.gkemulticloud.v1.AzureClusterResources.control_plane_application_security_group_id
	ControlPlaneApplicationSecurityGroupID *string `json:"controlPlaneApplicationSecurityGroupID,omitempty"`
}

// +kcc:proto=google.cloud.gkemulticloud.v1.Fleet
type FleetObservedState struct {
	// Output only. The name of the managed Hub Membership resource associated to
	//  this cluster.
	//
	//  Membership names are formatted as
	//  `projects/<project-number>/locations/global/membership/<cluster-id>`.
	// +kcc:proto:field=google.cloud.gkemulticloud.v1.Fleet.membership
	Membership *string `json:"membership,omitempty"`
}
