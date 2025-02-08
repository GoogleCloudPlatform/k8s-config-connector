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

// +kcc:proto=google.cloud.gkemulticloud.v1.AzureDiskTemplate
type AzureDiskTemplate struct {
	// Optional. The size of the disk, in GiBs.
	//
	//  When unspecified, a default value is provided. See the specific reference
	//  in the parent resource.
	// +kcc:proto:field=google.cloud.gkemulticloud.v1.AzureDiskTemplate.size_gib
	SizeGib *int32 `json:"sizeGib,omitempty"`
}

// +kcc:proto=google.cloud.gkemulticloud.v1.AzureNodeConfig
type AzureNodeConfig struct {
	// Optional. The Azure VM size name. Example: `Standard_DS2_v2`.
	//
	//  See [Supported VM
	//  sizes](/anthos/clusters/docs/azure/reference/supported-vms) for options.
	//
	//  When unspecified, it defaults to `Standard_DS2_v2`.
	// +kcc:proto:field=google.cloud.gkemulticloud.v1.AzureNodeConfig.vm_size
	VmSize *string `json:"vmSize,omitempty"`

	// Optional. Configuration related to the root volume provisioned for each
	//  node pool machine.
	//
	//  When unspecified, it defaults to a 32-GiB Azure Disk.
	// +kcc:proto:field=google.cloud.gkemulticloud.v1.AzureNodeConfig.root_volume
	RootVolume *AzureDiskTemplate `json:"rootVolume,omitempty"`

	// Optional. A set of tags to apply to all underlying Azure resources for this
	//  node pool. This currently only includes Virtual Machine Scale Sets.
	//
	//  Specify at most 50 pairs containing alphanumerics, spaces, and symbols
	//  (.+-=_:@/). Keys can be up to 127 Unicode characters. Values can be up to
	//  255 Unicode characters.
	// +kcc:proto:field=google.cloud.gkemulticloud.v1.AzureNodeConfig.tags
	Tags map[string]string `json:"tags,omitempty"`

	// Optional. The OS image type to use on node pool instances.
	//  Can be unspecified, or have a value of `ubuntu`.
	//
	//  When unspecified, it defaults to `ubuntu`.
	// +kcc:proto:field=google.cloud.gkemulticloud.v1.AzureNodeConfig.image_type
	ImageType *string `json:"imageType,omitempty"`

	// Required. SSH configuration for how to access the node pool machines.
	// +kcc:proto:field=google.cloud.gkemulticloud.v1.AzureNodeConfig.ssh_config
	SSHConfig *AzureSshConfig `json:"sshConfig,omitempty"`

	// Optional. Proxy configuration for outbound HTTP(S) traffic.
	// +kcc:proto:field=google.cloud.gkemulticloud.v1.AzureNodeConfig.proxy_config
	ProxyConfig *AzureProxyConfig `json:"proxyConfig,omitempty"`

	// Optional. Configuration related to vm config encryption.
	// +kcc:proto:field=google.cloud.gkemulticloud.v1.AzureNodeConfig.config_encryption
	ConfigEncryption *AzureConfigEncryption `json:"configEncryption,omitempty"`

	// Optional. The initial taints assigned to nodes of this node pool.
	// +kcc:proto:field=google.cloud.gkemulticloud.v1.AzureNodeConfig.taints
	Taints []NodeTaint `json:"taints,omitempty"`

	// Optional. The initial labels assigned to nodes of this node pool. An object
	//  containing a list of "key": value pairs. Example: { "name": "wrench",
	//  "mass": "1.3kg", "count": "3" }.
	// +kcc:proto:field=google.cloud.gkemulticloud.v1.AzureNodeConfig.labels
	Labels map[string]string `json:"labels,omitempty"`
}

// +kcc:proto=google.cloud.gkemulticloud.v1.AzureNodeManagement
type AzureNodeManagement struct {
	// Optional. Whether or not the nodes will be automatically repaired. When set
	//  to true, the nodes in this node pool will be monitored and if they fail
	//  health checks consistently over a period of time, an automatic repair
	//  action will be triggered to replace them with new nodes.
	// +kcc:proto:field=google.cloud.gkemulticloud.v1.AzureNodeManagement.auto_repair
	AutoRepair *bool `json:"autoRepair,omitempty"`
}

// +kcc:proto=google.cloud.gkemulticloud.v1.AzureNodePool
type AzureNodePool struct {
	// The name of this resource.
	//
	//  Node pool names are formatted as
	//  `projects/<project-number>/locations/<region>/azureClusters/<cluster-id>/azureNodePools/<node-pool-id>`.
	//
	//  For more details on Google Cloud resource names,
	//  see [Resource Names](https://cloud.google.com/apis/design/resource_names)
	// +kcc:proto:field=google.cloud.gkemulticloud.v1.AzureNodePool.name
	Name *string `json:"name,omitempty"`

	// Required. The Kubernetes version (e.g. `1.19.10-gke.1000`) running on this
	//  node pool.
	// +kcc:proto:field=google.cloud.gkemulticloud.v1.AzureNodePool.version
	Version *string `json:"version,omitempty"`

	// Required. The node configuration of the node pool.
	// +kcc:proto:field=google.cloud.gkemulticloud.v1.AzureNodePool.config
	Config *AzureNodeConfig `json:"config,omitempty"`

	// Required. The ARM ID of the subnet where the node pool VMs run. Make sure
	//  it's a subnet under the virtual network in the cluster configuration.
	// +kcc:proto:field=google.cloud.gkemulticloud.v1.AzureNodePool.subnet_id
	SubnetID *string `json:"subnetID,omitempty"`

	// Required. Autoscaler configuration for this node pool.
	// +kcc:proto:field=google.cloud.gkemulticloud.v1.AzureNodePool.autoscaling
	Autoscaling *AzureNodePoolAutoscaling `json:"autoscaling,omitempty"`

	// Allows clients to perform consistent read-modify-writes
	//  through optimistic concurrency control.
	//
	//  Can be sent on update and delete requests to ensure the
	//  client has an up-to-date value before proceeding.
	// +kcc:proto:field=google.cloud.gkemulticloud.v1.AzureNodePool.etag
	Etag *string `json:"etag,omitempty"`

	// Optional. Annotations on the node pool.
	//
	//  This field has the same restrictions as Kubernetes annotations.
	//  The total size of all keys and values combined is limited to 256k.
	//  Keys can have 2 segments: prefix (optional) and name (required),
	//  separated by a slash (/).
	//  Prefix must be a DNS subdomain.
	//  Name must be 63 characters or less, begin and end with alphanumerics,
	//  with dashes (-), underscores (_), dots (.), and alphanumerics between.
	// +kcc:proto:field=google.cloud.gkemulticloud.v1.AzureNodePool.annotations
	Annotations map[string]string `json:"annotations,omitempty"`

	// Required. The constraint on the maximum number of pods that can be run
	//  simultaneously on a node in the node pool.
	// +kcc:proto:field=google.cloud.gkemulticloud.v1.AzureNodePool.max_pods_constraint
	MaxPodsConstraint *MaxPodsConstraint `json:"maxPodsConstraint,omitempty"`

	// Optional. The Azure availability zone of the nodes in this nodepool.
	//
	//  When unspecified, it defaults to `1`.
	// +kcc:proto:field=google.cloud.gkemulticloud.v1.AzureNodePool.azure_availability_zone
	AzureAvailabilityZone *string `json:"azureAvailabilityZone,omitempty"`

	// Optional. The Management configuration for this node pool.
	// +kcc:proto:field=google.cloud.gkemulticloud.v1.AzureNodePool.management
	Management *AzureNodeManagement `json:"management,omitempty"`
}

// +kcc:proto=google.cloud.gkemulticloud.v1.AzureNodePoolAutoscaling
type AzureNodePoolAutoscaling struct {
	// Required. Minimum number of nodes in the node pool. Must be greater than or
	//  equal to 1 and less than or equal to max_node_count.
	// +kcc:proto:field=google.cloud.gkemulticloud.v1.AzureNodePoolAutoscaling.min_node_count
	MinNodeCount *int32 `json:"minNodeCount,omitempty"`

	// Required. Maximum number of nodes in the node pool. Must be greater than or
	//  equal to min_node_count and less than or equal to 50.
	// +kcc:proto:field=google.cloud.gkemulticloud.v1.AzureNodePoolAutoscaling.max_node_count
	MaxNodeCount *int32 `json:"maxNodeCount,omitempty"`
}

// +kcc:proto=google.cloud.gkemulticloud.v1.AzureNodePoolError
type AzureNodePoolError struct {
	// Human-friendly description of the error.
	// +kcc:proto:field=google.cloud.gkemulticloud.v1.AzureNodePoolError.message
	Message *string `json:"message,omitempty"`
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

// +kcc:proto=google.cloud.gkemulticloud.v1.AzureSshConfig
type AzureSshConfig struct {
	// Required. The SSH public key data for VMs managed by Anthos. This accepts
	//  the authorized_keys file format used in OpenSSH according to the sshd(8)
	//  manual page.
	// +kcc:proto:field=google.cloud.gkemulticloud.v1.AzureSshConfig.authorized_key
	AuthorizedKey *string `json:"authorizedKey,omitempty"`
}

// +kcc:proto=google.cloud.gkemulticloud.v1.MaxPodsConstraint
type MaxPodsConstraint struct {
	// Required. The maximum number of pods to schedule on a single node.
	// +kcc:proto:field=google.cloud.gkemulticloud.v1.MaxPodsConstraint.max_pods_per_node
	MaxPodsPerNode *int64 `json:"maxPodsPerNode,omitempty"`
}

// +kcc:proto=google.cloud.gkemulticloud.v1.NodeTaint
type NodeTaint struct {
	// Required. Key for the taint.
	// +kcc:proto:field=google.cloud.gkemulticloud.v1.NodeTaint.key
	Key *string `json:"key,omitempty"`

	// Required. Value for the taint.
	// +kcc:proto:field=google.cloud.gkemulticloud.v1.NodeTaint.value
	Value *string `json:"value,omitempty"`

	// Required. The taint effect.
	// +kcc:proto:field=google.cloud.gkemulticloud.v1.NodeTaint.effect
	Effect *string `json:"effect,omitempty"`
}

// +kcc:proto=google.cloud.gkemulticloud.v1.AzureNodePool
type AzureNodePoolObservedState struct {
	// Output only. The current state of the node pool.
	// +kcc:proto:field=google.cloud.gkemulticloud.v1.AzureNodePool.state
	State *string `json:"state,omitempty"`

	// Output only. A globally unique identifier for the node pool.
	// +kcc:proto:field=google.cloud.gkemulticloud.v1.AzureNodePool.uid
	Uid *string `json:"uid,omitempty"`

	// Output only. If set, there are currently pending changes to the node
	//  pool.
	// +kcc:proto:field=google.cloud.gkemulticloud.v1.AzureNodePool.reconciling
	Reconciling *bool `json:"reconciling,omitempty"`

	// Output only. The time at which this node pool was created.
	// +kcc:proto:field=google.cloud.gkemulticloud.v1.AzureNodePool.create_time
	CreateTime *string `json:"createTime,omitempty"`

	// Output only. The time at which this node pool was last updated.
	// +kcc:proto:field=google.cloud.gkemulticloud.v1.AzureNodePool.update_time
	UpdateTime *string `json:"updateTime,omitempty"`

	// Output only. A set of errors found in the node pool.
	// +kcc:proto:field=google.cloud.gkemulticloud.v1.AzureNodePool.errors
	Errors []AzureNodePoolError `json:"errors,omitempty"`
}
