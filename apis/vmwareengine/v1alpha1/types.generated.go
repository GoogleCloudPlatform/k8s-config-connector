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


// +kcc:proto=google.cloud.vmwareengine.v1.Hcx
type Hcx struct {
	// Internal IP address of the appliance.
	// +kcc:proto:field=google.cloud.vmwareengine.v1.Hcx.internal_ip
	InternalIP *string `json:"internalIP,omitempty"`

	// Version of the appliance.
	// +kcc:proto:field=google.cloud.vmwareengine.v1.Hcx.version
	Version *string `json:"version,omitempty"`

	// Fully qualified domain name of the appliance.
	// +kcc:proto:field=google.cloud.vmwareengine.v1.Hcx.fqdn
	Fqdn *string `json:"fqdn,omitempty"`
}

// +kcc:proto=google.cloud.vmwareengine.v1.NetworkConfig
type NetworkConfig struct {
	// Required. Management CIDR used by VMware management appliances.
	// +kcc:proto:field=google.cloud.vmwareengine.v1.NetworkConfig.management_cidr
	ManagementCidr *string `json:"managementCidr,omitempty"`

	// Optional. The relative resource name of the VMware Engine network attached
	//  to the private cloud. Specify the name in the following form:
	//  `projects/{project}/locations/{location}/vmwareEngineNetworks/{vmware_engine_network_id}`
	//  where `{project}` can either be a project number or a project ID.
	// +kcc:proto:field=google.cloud.vmwareengine.v1.NetworkConfig.vmware_engine_network
	VmwareEngineNetwork *string `json:"vmwareEngineNetwork,omitempty"`
}

// +kcc:proto=google.cloud.vmwareengine.v1.NodeTypeConfig
type NodeTypeConfig struct {
	// Required. The number of nodes of this type in the cluster
	// +kcc:proto:field=google.cloud.vmwareengine.v1.NodeTypeConfig.node_count
	NodeCount *int32 `json:"nodeCount,omitempty"`

	// Optional. Customized number of cores available to each node of the type.
	//  This number must always be one of `nodeType.availableCustomCoreCounts`.
	//  If zero is provided max value from `nodeType.availableCustomCoreCounts`
	//  will be used.
	// +kcc:proto:field=google.cloud.vmwareengine.v1.NodeTypeConfig.custom_core_count
	CustomCoreCount *int32 `json:"customCoreCount,omitempty"`
}

// +kcc:proto=google.cloud.vmwareengine.v1.Nsx
type Nsx struct {
	// Internal IP address of the appliance.
	// +kcc:proto:field=google.cloud.vmwareengine.v1.Nsx.internal_ip
	InternalIP *string `json:"internalIP,omitempty"`

	// Version of the appliance.
	// +kcc:proto:field=google.cloud.vmwareengine.v1.Nsx.version
	Version *string `json:"version,omitempty"`

	// Fully qualified domain name of the appliance.
	// +kcc:proto:field=google.cloud.vmwareengine.v1.Nsx.fqdn
	Fqdn *string `json:"fqdn,omitempty"`
}

// +kcc:proto=google.cloud.vmwareengine.v1.PrivateCloud
type PrivateCloud struct {

	// Required. Network configuration of the private cloud.
	// +kcc:proto:field=google.cloud.vmwareengine.v1.PrivateCloud.network_config
	NetworkConfig *NetworkConfig `json:"networkConfig,omitempty"`

	// Required. Input only. The management cluster for this private cloud.
	//  This field is required during creation of the private cloud to provide
	//  details for the default cluster.
	//
	//  The following fields can't be changed after private cloud creation:
	//  `ManagementCluster.clusterId`, `ManagementCluster.nodeTypeId`.
	// +kcc:proto:field=google.cloud.vmwareengine.v1.PrivateCloud.management_cluster
	ManagementCluster *PrivateCloud_ManagementCluster `json:"managementCluster,omitempty"`

	// User-provided description for this private cloud.
	// +kcc:proto:field=google.cloud.vmwareengine.v1.PrivateCloud.description
	Description *string `json:"description,omitempty"`

	// Optional. Type of the private cloud. Defaults to STANDARD.
	// +kcc:proto:field=google.cloud.vmwareengine.v1.PrivateCloud.type
	Type *string `json:"type,omitempty"`
}

// +kcc:proto=google.cloud.vmwareengine.v1.PrivateCloud.ManagementCluster
type PrivateCloud_ManagementCluster struct {
	// Required. The user-provided identifier of the new `Cluster`.
	//  The identifier must meet the following requirements:
	//
	//  * Only contains 1-63 alphanumeric characters and hyphens
	//  * Begins with an alphabetical character
	//  * Ends with a non-hyphen character
	//  * Not formatted as a UUID
	//  * Complies with [RFC
	//  1034](https://datatracker.ietf.org/doc/html/rfc1034) (section 3.5)
	// +kcc:proto:field=google.cloud.vmwareengine.v1.PrivateCloud.ManagementCluster.cluster_id
	ClusterID *string `json:"clusterID,omitempty"`

	// TODO: unsupported map type with key string and value message


	// Optional. Configuration of a stretched cluster. Required for STRETCHED
	//  private clouds.
	// +kcc:proto:field=google.cloud.vmwareengine.v1.PrivateCloud.ManagementCluster.stretched_cluster_config
	StretchedClusterConfig *StretchedClusterConfig `json:"stretchedClusterConfig,omitempty"`
}

// +kcc:proto=google.cloud.vmwareengine.v1.StretchedClusterConfig
type StretchedClusterConfig struct {
	// Required. Zone that will remain operational when connection between the two
	//  zones is lost. Specify the resource name of a zone that belongs to the
	//  region of the private cloud. For example:
	//  `projects/{project}/locations/europe-west3-a` where `{project}` can either
	//  be a project number or a project ID.
	// +kcc:proto:field=google.cloud.vmwareengine.v1.StretchedClusterConfig.preferred_location
	PreferredLocation *string `json:"preferredLocation,omitempty"`

	// Required. Additional zone for a higher level of availability and load
	//  balancing. Specify the resource name of a zone that belongs to the region
	//  of the private cloud. For example:
	//  `projects/{project}/locations/europe-west3-b` where `{project}` can either
	//  be a project number or a project ID.
	// +kcc:proto:field=google.cloud.vmwareengine.v1.StretchedClusterConfig.secondary_location
	SecondaryLocation *string `json:"secondaryLocation,omitempty"`
}

// +kcc:proto=google.cloud.vmwareengine.v1.Vcenter
type Vcenter struct {
	// Internal IP address of the appliance.
	// +kcc:proto:field=google.cloud.vmwareengine.v1.Vcenter.internal_ip
	InternalIP *string `json:"internalIP,omitempty"`

	// Version of the appliance.
	// +kcc:proto:field=google.cloud.vmwareengine.v1.Vcenter.version
	Version *string `json:"version,omitempty"`

	// Fully qualified domain name of the appliance.
	// +kcc:proto:field=google.cloud.vmwareengine.v1.Vcenter.fqdn
	Fqdn *string `json:"fqdn,omitempty"`
}

// +kcc:proto=google.cloud.vmwareengine.v1.Hcx
type HcxObservedState struct {
	// Output only. The state of the appliance.
	// +kcc:proto:field=google.cloud.vmwareengine.v1.Hcx.state
	State *string `json:"state,omitempty"`
}

// +kcc:proto=google.cloud.vmwareengine.v1.NetworkConfig
type NetworkConfigObservedState struct {
	// Output only. The canonical name of the VMware Engine network in the form:
	//  `projects/{project_number}/locations/{location}/vmwareEngineNetworks/{vmware_engine_network_id}`
	// +kcc:proto:field=google.cloud.vmwareengine.v1.NetworkConfig.vmware_engine_network_canonical
	VmwareEngineNetworkCanonical *string `json:"vmwareEngineNetworkCanonical,omitempty"`

	// Output only. The IP address layout version of the management IP address
	//  range. Possible versions include:
	//  * `managementIpAddressLayoutVersion=1`: Indicates the legacy IP address
	//  layout used by some existing private clouds. This is no longer supported
	//  for new private clouds as it does not support all features.
	//  * `managementIpAddressLayoutVersion=2`: Indicates the latest IP address
	//  layout used by all newly created private clouds. This version supports all
	//  current features.
	// +kcc:proto:field=google.cloud.vmwareengine.v1.NetworkConfig.management_ip_address_layout_version
	ManagementIPAddressLayoutVersion *int32 `json:"managementIPAddressLayoutVersion,omitempty"`

	// Output only. DNS Server IP of the Private Cloud.
	//  All DNS queries can be forwarded to this address for name resolution of
	//  Private Cloud's management entities like vCenter, NSX-T Manager and
	//  ESXi hosts.
	// +kcc:proto:field=google.cloud.vmwareengine.v1.NetworkConfig.dns_server_ip
	DnsServerIP *string `json:"dnsServerIP,omitempty"`
}

// +kcc:proto=google.cloud.vmwareengine.v1.Nsx
type NsxObservedState struct {
	// Output only. The state of the appliance.
	// +kcc:proto:field=google.cloud.vmwareengine.v1.Nsx.state
	State *string `json:"state,omitempty"`
}

// +kcc:proto=google.cloud.vmwareengine.v1.PrivateCloud
type PrivateCloudObservedState struct {
	// Output only. The resource name of this private cloud.
	//  Resource names are schemeless URIs that follow the conventions in
	//  https://cloud.google.com/apis/design/resource_names.
	//  For example:
	//  `projects/my-project/locations/us-central1-a/privateClouds/my-cloud`
	// +kcc:proto:field=google.cloud.vmwareengine.v1.PrivateCloud.name
	Name *string `json:"name,omitempty"`

	// Output only. Creation time of this resource.
	// +kcc:proto:field=google.cloud.vmwareengine.v1.PrivateCloud.create_time
	CreateTime *string `json:"createTime,omitempty"`

	// Output only. Last update time of this resource.
	// +kcc:proto:field=google.cloud.vmwareengine.v1.PrivateCloud.update_time
	UpdateTime *string `json:"updateTime,omitempty"`

	// Output only. Time when the resource was scheduled for deletion.
	// +kcc:proto:field=google.cloud.vmwareengine.v1.PrivateCloud.delete_time
	DeleteTime *string `json:"deleteTime,omitempty"`

	// Output only. Time when the resource will be irreversibly deleted.
	// +kcc:proto:field=google.cloud.vmwareengine.v1.PrivateCloud.expire_time
	ExpireTime *string `json:"expireTime,omitempty"`

	// Output only. State of the resource. New values may be added to this enum
	//  when appropriate.
	// +kcc:proto:field=google.cloud.vmwareengine.v1.PrivateCloud.state
	State *string `json:"state,omitempty"`

	// Required. Network configuration of the private cloud.
	// +kcc:proto:field=google.cloud.vmwareengine.v1.PrivateCloud.network_config
	NetworkConfig *NetworkConfigObservedState `json:"networkConfig,omitempty"`

	// Output only. HCX appliance.
	// +kcc:proto:field=google.cloud.vmwareengine.v1.PrivateCloud.hcx
	Hcx *Hcx `json:"hcx,omitempty"`

	// Output only. NSX appliance.
	// +kcc:proto:field=google.cloud.vmwareengine.v1.PrivateCloud.nsx
	Nsx *Nsx `json:"nsx,omitempty"`

	// Output only. Vcenter appliance.
	// +kcc:proto:field=google.cloud.vmwareengine.v1.PrivateCloud.vcenter
	Vcenter *Vcenter `json:"vcenter,omitempty"`

	// Output only. System-generated unique identifier for the resource.
	// +kcc:proto:field=google.cloud.vmwareengine.v1.PrivateCloud.uid
	Uid *string `json:"uid,omitempty"`
}

// +kcc:proto=google.cloud.vmwareengine.v1.Vcenter
type VcenterObservedState struct {
	// Output only. The state of the appliance.
	// +kcc:proto:field=google.cloud.vmwareengine.v1.Vcenter.state
	State *string `json:"state,omitempty"`
}
