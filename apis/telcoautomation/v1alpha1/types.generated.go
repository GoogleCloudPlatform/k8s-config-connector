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


// +kcc:proto=google.cloud.telcoautomation.v1.FullManagementConfig
type FullManagementConfig struct {
	// Optional. Name of the VPC Network to put the GKE cluster and nodes in. The
	//  VPC will be created if it doesn't exist.
	// +kcc:proto:field=google.cloud.telcoautomation.v1.FullManagementConfig.network
	Network *string `json:"network,omitempty"`

	// Optional. Specifies the subnet that the interface will be part of. Network
	//  key must be specified and the subnet must be a subnetwork of the specified
	//  network.
	// +kcc:proto:field=google.cloud.telcoautomation.v1.FullManagementConfig.subnet
	Subnet *string `json:"subnet,omitempty"`

	// Optional. The /28 network that the masters will use.
	// +kcc:proto:field=google.cloud.telcoautomation.v1.FullManagementConfig.master_ipv4_cidr_block
	MasterIpv4CidrBlock *string `json:"masterIpv4CidrBlock,omitempty"`

	// Optional. The IP address range for the cluster pod IPs. Set to blank to
	//  have a range chosen with the default size. Set to /netmask (e.g. /14) to
	//  have a range chosen with a specific netmask. Set to a CIDR notation
	//  (e.g. 10.96.0.0/14) from the RFC-1918 private networks (e.g. 10.0.0.0/8,
	//  172.16.0.0/12, 192.168.0.0/16) to pick a specific range to use.
	// +kcc:proto:field=google.cloud.telcoautomation.v1.FullManagementConfig.cluster_cidr_block
	ClusterCidrBlock *string `json:"clusterCidrBlock,omitempty"`

	// Optional. The IP address range for the cluster service IPs. Set to blank to
	//  have a range chosen with the default size. Set to /netmask (e.g. /14) to
	//  have a range chosen with a specific netmask. Set to a CIDR notation (e.g.
	//  10.96.0.0/14) from the RFC-1918 private networks (e.g. 10.0.0.0/8,
	//  172.16.0.0/12, 192.168.0.0/16) to pick a specific range to use.
	// +kcc:proto:field=google.cloud.telcoautomation.v1.FullManagementConfig.services_cidr_block
	ServicesCidrBlock *string `json:"servicesCidrBlock,omitempty"`

	// Optional. The name of the existing secondary range in the cluster's
	//  subnetwork to use for pod IP addresses. Alternatively, cluster_cidr_block
	//  can be used to automatically create a GKE-managed one.
	// +kcc:proto:field=google.cloud.telcoautomation.v1.FullManagementConfig.cluster_named_range
	ClusterNamedRange *string `json:"clusterNamedRange,omitempty"`

	// Optional. The name of the existing secondary range in the cluster's
	//  subnetwork to use for service ClusterIPs. Alternatively,
	//  services_cidr_block can be used to automatically create a GKE-managed one.
	// +kcc:proto:field=google.cloud.telcoautomation.v1.FullManagementConfig.services_named_range
	ServicesNamedRange *string `json:"servicesNamedRange,omitempty"`

	// Optional. Master Authorized Network that supports multiple CIDR blocks.
	//  Allows access to the k8s master from multiple blocks. It cannot be set at
	//  the same time with the field man_block.
	// +kcc:proto:field=google.cloud.telcoautomation.v1.FullManagementConfig.master_authorized_networks_config
	MasterAuthorizedNetworksConfig *MasterAuthorizedNetworksConfig `json:"masterAuthorizedNetworksConfig,omitempty"`
}

// +kcc:proto=google.cloud.telcoautomation.v1.ManagementConfig
type ManagementConfig struct {
	// Configuration of the standard (GKE) cluster management
	// +kcc:proto:field=google.cloud.telcoautomation.v1.ManagementConfig.standard_management_config
	StandardManagementConfig *StandardManagementConfig `json:"standardManagementConfig,omitempty"`

	// Configuration of the full (Autopilot) cluster management. Full cluster
	//  management is a preview feature.
	// +kcc:proto:field=google.cloud.telcoautomation.v1.ManagementConfig.full_management_config
	FullManagementConfig *FullManagementConfig `json:"fullManagementConfig,omitempty"`
}

// +kcc:proto=google.cloud.telcoautomation.v1.MasterAuthorizedNetworksConfig
type MasterAuthorizedNetworksConfig struct {
	// Optional. cidr_blocks define up to 50 external networks that could access
	//  Kubernetes master through HTTPS.
	// +kcc:proto:field=google.cloud.telcoautomation.v1.MasterAuthorizedNetworksConfig.cidr_blocks
	CidrBlocks []MasterAuthorizedNetworksConfig_CidrBlock `json:"cidrBlocks,omitempty"`
}

// +kcc:proto=google.cloud.telcoautomation.v1.MasterAuthorizedNetworksConfig.CidrBlock
type MasterAuthorizedNetworksConfig_CidrBlock struct {
	// Optional. display_name is an optional field for users to identify CIDR
	//  blocks.
	// +kcc:proto:field=google.cloud.telcoautomation.v1.MasterAuthorizedNetworksConfig.CidrBlock.display_name
	DisplayName *string `json:"displayName,omitempty"`

	// Optional. cidr_block must be specified in CIDR notation when using
	//  master_authorized_networks_config. Currently, the user could still use
	//  the deprecated man_block field, so this field is currently optional, but
	//  will be required in the future.
	// +kcc:proto:field=google.cloud.telcoautomation.v1.MasterAuthorizedNetworksConfig.CidrBlock.cidr_block
	CidrBlock *string `json:"cidrBlock,omitempty"`
}

// +kcc:proto=google.cloud.telcoautomation.v1.OrchestrationCluster
type OrchestrationCluster struct {
	// Name of the orchestration cluster. The name of orchestration cluster cannot
	//  be more than 24 characters.
	// +kcc:proto:field=google.cloud.telcoautomation.v1.OrchestrationCluster.name
	Name *string `json:"name,omitempty"`

	// Management configuration of the underlying GKE cluster.
	// +kcc:proto:field=google.cloud.telcoautomation.v1.OrchestrationCluster.management_config
	ManagementConfig *ManagementConfig `json:"managementConfig,omitempty"`

	// Labels as key value pairs.
	// +kcc:proto:field=google.cloud.telcoautomation.v1.OrchestrationCluster.labels
	Labels map[string]string `json:"labels,omitempty"`
}

// +kcc:proto=google.cloud.telcoautomation.v1.StandardManagementConfig
type StandardManagementConfig struct {
	// Optional. Name of the VPC Network to put the GKE cluster and nodes in. The
	//  VPC will be created if it doesn't exist.
	// +kcc:proto:field=google.cloud.telcoautomation.v1.StandardManagementConfig.network
	Network *string `json:"network,omitempty"`

	// Optional. Specifies the subnet that the interface will be part of. Network
	//  key must be specified and the subnet must be a subnetwork of the specified
	//  network.
	// +kcc:proto:field=google.cloud.telcoautomation.v1.StandardManagementConfig.subnet
	Subnet *string `json:"subnet,omitempty"`

	// Optional. The /28 network that the masters will use. It should be free
	//  within the network.
	// +kcc:proto:field=google.cloud.telcoautomation.v1.StandardManagementConfig.master_ipv4_cidr_block
	MasterIpv4CidrBlock *string `json:"masterIpv4CidrBlock,omitempty"`

	// Optional. The IP address range for the cluster pod IPs. Set to blank to
	//  have a range chosen with the default size. Set to /netmask (e.g. /14) to
	//  have a range chosen with a specific netmask. Set to a CIDR notation
	//  (e.g. 10.96.0.0/14) from the RFC-1918 private networks (e.g. 10.0.0.0/8,
	//  172.16.0.0/12, 192.168.0.0/16) to pick a specific range to use.
	// +kcc:proto:field=google.cloud.telcoautomation.v1.StandardManagementConfig.cluster_cidr_block
	ClusterCidrBlock *string `json:"clusterCidrBlock,omitempty"`

	// Optional. The IP address range for the cluster service IPs. Set to blank to
	//  have a range chosen with the default size. Set to /netmask (e.g. /14) to
	//  have a range chosen with a specific netmask. Set to a CIDR notation (e.g.
	//  10.96.0.0/14) from the RFC-1918 private networks (e.g. 10.0.0.0/8,
	//  172.16.0.0/12, 192.168.0.0/16) to pick a specific range to use.
	// +kcc:proto:field=google.cloud.telcoautomation.v1.StandardManagementConfig.services_cidr_block
	ServicesCidrBlock *string `json:"servicesCidrBlock,omitempty"`

	// Optional. The name of the existing secondary range in the cluster's
	//  subnetwork to use for pod IP addresses. Alternatively, cluster_cidr_block
	//  can be used to automatically create a GKE-managed one.
	// +kcc:proto:field=google.cloud.telcoautomation.v1.StandardManagementConfig.cluster_named_range
	ClusterNamedRange *string `json:"clusterNamedRange,omitempty"`

	// Optional. The name of the existing secondary range in the cluster's
	//  subnetwork to use for service ClusterIPs. Alternatively,
	//  services_cidr_block can be used to automatically create a GKE-managed one.
	// +kcc:proto:field=google.cloud.telcoautomation.v1.StandardManagementConfig.services_named_range
	ServicesNamedRange *string `json:"servicesNamedRange,omitempty"`

	// Optional. Master Authorized Network that supports multiple CIDR blocks.
	//  Allows access to the k8s master from multiple blocks. It cannot be set at
	//  the same time with the field man_block.
	// +kcc:proto:field=google.cloud.telcoautomation.v1.StandardManagementConfig.master_authorized_networks_config
	MasterAuthorizedNetworksConfig *MasterAuthorizedNetworksConfig `json:"masterAuthorizedNetworksConfig,omitempty"`
}

// +kcc:proto=google.cloud.telcoautomation.v1.OrchestrationCluster
type OrchestrationClusterObservedState struct {
	// Output only. [Output only] Create time stamp.
	// +kcc:proto:field=google.cloud.telcoautomation.v1.OrchestrationCluster.create_time
	CreateTime *string `json:"createTime,omitempty"`

	// Output only. [Output only] Update time stamp.
	// +kcc:proto:field=google.cloud.telcoautomation.v1.OrchestrationCluster.update_time
	UpdateTime *string `json:"updateTime,omitempty"`

	// Output only. Provides the TNA version installed on the cluster.
	// +kcc:proto:field=google.cloud.telcoautomation.v1.OrchestrationCluster.tna_version
	TnaVersion *string `json:"tnaVersion,omitempty"`

	// Output only. State of the Orchestration Cluster.
	// +kcc:proto:field=google.cloud.telcoautomation.v1.OrchestrationCluster.state
	State *string `json:"state,omitempty"`
}
