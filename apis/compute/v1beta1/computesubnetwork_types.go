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
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/apis/k8s/v1alpha1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

var ComputeSubnetworkGVK = GroupVersion.WithKind("ComputeSubnetwork")

// ComputeSubnetworkSpec defines the desired state of ComputeSubnetwork
// +kcc:spec:proto=google.cloud.compute.v1.Subnetwork
type ComputeSubnetworkSpec struct {
	// Immutable. An optional description of this resource. Provide this property when
	// you create the resource. This field can be set only at resource
	// creation time.
	// +kcc:proto:field=google.cloud.compute.v1.Subnetwork.description
	Description *string `json:"description,omitempty"`

	// The range of internal addresses that are owned by this subnetwork.
	// Provide this property when you create the subnetwork. For example,
	// 10.0.0.0/8 or 192.168.0.0/16. Ranges must be unique and
	// non-overlapping within a network. Only IPv4 is supported.
	// +required
	// +kcc:proto:field=google.cloud.compute.v1.Subnetwork.ip_cidr_range
	IPCIDRRange *string `json:"ipCidrRange,omitempty"`

	// The access type of IPv6 address this subnet holds. It's immutable and can only be specified during creation
	// or the first time the subnet is updated into IPV4_IPV6 dual stack. If the ipv6_type is EXTERNAL then this subnet
	// cannot enable direct path. Possible values: ["EXTERNAL", "INTERNAL"].
	// +kcc:proto:field=google.cloud.compute.v1.Subnetwork.ipv6_access_type
	IPV6AccessType *string `json:"ipv6AccessType,omitempty"`

	// This field denotes the VPC flow logging options for this subnetwork. If
	// logging is enabled, logs are exported to Cloud Logging. Flow logging
	// isn't supported if the subnet 'purpose' field is set to subnetwork is
	// 'REGIONAL_MANAGED_PROXY' or 'GLOBAL_MANAGED_PROXY'.
	// +kcc:proto:field=google.cloud.compute.v1.Subnetwork.log_config
	LogConfig *SubnetworkLogConfig `json:"logConfig,omitempty"`

	// The network this subnet belongs to. Only networks that are in the
	// distributed mode can have subnetworks.
	// +required
	// +kcc:proto:field=google.cloud.compute.v1.Subnetwork.network
	NetworkRef *ComputeNetworkRef `json:"networkRef"`

	// When enabled, VMs in this subnetwork without external IP addresses can
	// access Google APIs and services by using Private Google Access.
	// +kcc:proto:field=google.cloud.compute.v1.Subnetwork.private_ip_google_access
	PrivateIPGoogleAccess *bool `json:"privateIpGoogleAccess,omitempty"`

	// The private IPv6 google access type for the VMs in this
	// subnet.
	// +kcc:proto:field=google.cloud.compute.v1.Subnetwork.private_ipv6_google_access
	PrivateIPV6GoogleAccess *string `json:"privateIpv6GoogleAccess,omitempty"`

	// Immutable. The purpose of the resource. This field can be either 'PRIVATE_RFC_1918', 'REGIONAL_MANAGED_PROXY', 'GLOBAL_MANAGED_PROXY', or 'PRIVATE_SERVICE_CONNECT'.
	// A subnet with purpose set to 'REGIONAL_MANAGED_PROXY' is a user-created subnetwork that is reserved for regional Envoy-based load balancers.
	// A subnetwork in a given region with purpose set to 'GLOBAL_MANAGED_PROXY' is a proxy-only subnet and is shared between all the cross-regional Envoy-based load balancers.
	// A subnetwork with purpose set to 'PRIVATE_SERVICE_CONNECT' reserves the subnet for hosting a Private Service Connect published service.
	// Note that 'REGIONAL_MANAGED_PROXY' is the preferred setting for all regional Envoy load balancers.
	// If unspecified, the purpose defaults to 'PRIVATE_RFC_1918'.
	// +kcc:proto:field=google.cloud.compute.v1.Subnetwork.purpose
	Purpose *string `json:"purpose,omitempty"`

	// Immutable. The GCP region for this subnetwork.
	// +required
	Region *string `json:"region"`

	// Immutable. Optional. The name of the resource. Used for
	// creation and acquisition. When unset, the value of `metadata.name`
	// is used as the default.
	ResourceID *string `json:"resourceID,omitempty"`

	// The role of subnetwork.
	// Currently, this field is only used when 'purpose' is 'REGIONAL_MANAGED_PROXY'.
	// The value can be set to 'ACTIVE' or 'BACKUP'.
	// An 'ACTIVE' subnetwork is one that is currently being used for Envoy-based load balancers in a region.
	// A 'BACKUP' subnetwork is one that is ready to be promoted to 'ACTIVE' or is currently draining. Possible values: ["ACTIVE", "BACKUP"].
	// +kcc:proto:field=google.cloud.compute.v1.Subnetwork.role
	Role *string `json:"role,omitempty"`

	// +kcc:proto:field=google.cloud.compute.v1.Subnetwork.secondary_ip_ranges
	SecondaryIPRanges []SubnetworkSecondaryRange `json:"secondaryIpRange,omitempty"`

	// The stack type for this subnet to identify whether the IPv6 feature is enabled or not.
	// If not specified IPV4_ONLY will be used. Possible values: ["IPV4_ONLY", "IPV4_IPV6"].
	// +kcc:proto:field=google.cloud.compute.v1.Subnetwork.stack_type
	StackType *string `json:"stackType,omitempty"`
}

// ComputeSubnetworkStatus defines the config connector machine state of ComputeSubnetwork
// +kcc:status:proto=google.cloud.compute.v1.Subnetwork
type ComputeSubnetworkStatus struct {
	/* Conditions represent the latest available observations of the
	   object's current state. */
	Conditions []v1alpha1.Condition `json:"conditions,omitempty"`

	// ObservedGeneration is the generation of the resource that was most recently observed by the Config Connector controller. If this is equal to metadata.generation, then that means that the current reported status reflects the most recent desired state of the resource.
	ObservedGeneration *int64 `json:"observedGeneration,omitempty"`

	// Creation timestamp in RFC3339 text format.
	// +kcc:proto:field=google.cloud.compute.v1.Subnetwork.creation_timestamp
	CreationTimestamp *string `json:"creationTimestamp,omitempty"`

	// The range of external IPv6 addresses that are owned by
	// this subnetwork.
	// +kcc:proto:field=google.cloud.compute.v1.Subnetwork.external_ipv6_prefix
	ExternalIPV6Prefix *string `json:"externalIpv6Prefix,omitempty"`

	// DEPRECATED. This field is not useful for users, and has
	// been removed as an output. Fingerprint of this resource. This field
	// is used internally during updates of this resource.
	// +kcc:proto:field=google.cloud.compute.v1.Subnetwork.fingerprint
	Fingerprint *string `json:"fingerprint,omitempty"`

	// The gateway address for default routes to reach destination addresses
	// outside this subnetwork.
	// +kcc:proto:field=google.cloud.compute.v1.Subnetwork.gateway_address
	GatewayAddress *string `json:"gatewayAddress,omitempty"`

	// The internal IPv6 address range that is assigned to this
	// subnetwork.
	// +kcc:proto:field=google.cloud.compute.v1.Subnetwork.internal_ipv6_prefix
	InternalIPV6Prefix *string `json:"internalIpv6Prefix,omitempty"`

	// The range of internal IPv6 addresses that are owned by
	// this subnetwork.
	// +kcc:proto:field=google.cloud.compute.v1.Subnetwork.ipv6_cidr_range
	IPV6CIDRRange *string `json:"ipv6CidrRange,omitempty"`

	// Server-defined URL for the resource.
	// +kcc:proto:field=google.cloud.compute.v1.Subnetwork.self_link
	SelfLink *string `json:"selfLink,omitempty"`
}

// +genclient
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:resource:categories=gcp,shortName=gcpcomputesubnetwork;gcpcomputesubnetworks
// +kubebuilder:subresource:status
// +kubebuilder:metadata:labels="cnrm.cloud.google.com/managed-by-kcc=true"
// +kubebuilder:metadata:labels="cnrm.cloud.google.com/stability-level=stable"
// +kubebuilder:metadata:labels="cnrm.cloud.google.com/system=true"
// +kubebuilder:metadata:labels="cnrm.cloud.google.com/tf2crd=true"
// +kubebuilder:printcolumn:name="Age",JSONPath=".metadata.creationTimestamp",type="date"
// +kubebuilder:printcolumn:name="Ready",JSONPath=".status.conditions[?(@.type=='Ready')].status",type="string",description="When 'True', the most recent reconcile of the resource succeeded"
// +kubebuilder:printcolumn:name="Status",JSONPath=".status.conditions[?(@.type=='Ready')].reason",type="string",description="The reason for the value in 'Ready'"
// +kubebuilder:printcolumn:name="Status Age",JSONPath=".status.conditions[?(@.type=='Ready')].lastTransitionTime",type="date",description="The last transition time for the value in 'Status'"

// ComputeSubnetwork is the Schema for the ComputeSubnetwork API
// +k8s:openapi-gen=true
type ComputeSubnetwork struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	// +required
	Spec   ComputeSubnetworkSpec   `json:"spec,omitempty"`
	Status ComputeSubnetworkStatus `json:"status,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// ComputeSubnetworkList contains a list of ComputeSubnetwork
type ComputeSubnetworkList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []ComputeSubnetwork `json:"items"`
}

func init() {
	SchemeBuilder.Register(&ComputeSubnetwork{}, &ComputeSubnetworkList{})
}

// +kcc:proto=google.cloud.compute.v1.SubnetworkLogConfig
type SubnetworkLogConfig struct {
	// Can only be specified if VPC flow logging for this subnetwork is enabled. Toggles the aggregation interval for collecting flow logs. Increasing the interval time will reduce the amount of generated flow logs for long lasting connections. Default is an interval of 5 seconds per connection.
	//  Check the AggregationInterval enum for the list of possible values.
	// +kcc:proto:field=google.cloud.compute.v1.SubnetworkLogConfig.aggregation_interval
	AggregationInterval *string `json:"aggregationInterval,omitempty"`

	// Can only be specified if VPC flow logs for this subnetwork is enabled. The filter expression is used to define which VPC flow logs should be exported to Cloud Logging.
	// +kcc:proto:field=google.cloud.compute.v1.SubnetworkLogConfig.filter_expr
	FilterExpr *string `json:"filterExpr,omitempty"`

	// Can only be specified if VPC flow logging for this subnetwork is enabled. The value of the field must be in [0, 1]. Set the sampling rate of VPC flow logs within the subnetwork where 1.0 means all collected logs are reported and 0.0 means no logs are reported. Default is 0.5 unless otherwise specified by the org policy, which means half of all collected logs are reported.
	// +kcc:proto:field=google.cloud.compute.v1.SubnetworkLogConfig.flow_sampling
	FlowSampling *float32 `json:"flowSampling,omitempty"`

	// Can only be specified if VPC flow logs for this subnetwork is enabled. Configures whether all, none or a subset of metadata fields should be added to the reported VPC flow logs. Default is EXCLUDE_ALL_METADATA.
	//  Check the Metadata enum for the list of possible values.
	// +kcc:proto:field=google.cloud.compute.v1.SubnetworkLogConfig.metadata
	Metadata *string `json:"metadata,omitempty"`

	// Can only be specified if VPC flow logs for this subnetwork is enabled and "metadata" was set to CUSTOM_METADATA.
	// +kcc:proto:field=google.cloud.compute.v1.SubnetworkLogConfig.metadata_fields
	MetadataFields []string `json:"metadataFields,omitempty"`
}

// +kcc:proto=google.cloud.compute.v1.SubnetworkSecondaryRange
type SubnetworkSecondaryRange struct {
	// The range of IP addresses belonging to this subnetwork secondary range. Provide this property when you create the subnetwork. Ranges must be unique and non-overlapping with all primary and secondary IP ranges within a network. Only IPv4 is supported. The range can be any range listed in the Valid ranges list.
	// +required
	// +kcc:proto:field=google.cloud.compute.v1.SubnetworkSecondaryRange.ip_cidr_range
	IPCIDRRange *string `json:"ipCidrRange,omitempty"`

	// The name associated with this subnetwork secondary range, used when adding an alias IP range to a VM instance. The name must be 1-63 characters long, and comply with RFC1035. The name must be unique within the subnetwork.
	// +required
	// +kcc:proto:field=google.cloud.compute.v1.SubnetworkSecondaryRange.range_name
	RangeName *string `json:"rangeName,omitempty"`
}
