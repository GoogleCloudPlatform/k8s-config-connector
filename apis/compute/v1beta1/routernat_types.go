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
	refs "github.com/GoogleCloudPlatform/k8s-config-connector/apis/refs/v1beta1"
	commonv1alpha1 "github.com/GoogleCloudPlatform/k8s-config-connector/pkg/apis/common/v1alpha1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime/schema"
)

var (
	ComputeRouterNATGVK = schema.GroupVersionKind{
		Group:   GroupVersion.Group,
		Version: GroupVersion.Version,
		Kind:    "ComputeRouterNAT",
	}
)

// +kcc:proto=google.cloud.compute.v1.RouterNatLogConfig
type RouterNatLogConfig struct {
	/* Indicates whether or not to export logs. */
	// +kcc:proto:field=google.cloud.compute.v1.RouterNatLogConfig.enable
	Enable *bool `json:"enable"`

	/* Specifies the desired filtering of logs on this NAT. Possible values: ["ERRORS_ONLY", "TRANSLATIONS_ONLY", "ALL"]. */
	// +kcc:proto:field=google.cloud.compute.v1.RouterNatLogConfig.filter
	Filter *string `json:"filter"`
}

// +kcc:proto=google.cloud.compute.v1.RouterNatRuleAction
type RouterNatRuleAction struct {
	/* A list of URLs of the IP resources used for this NAT rule. These IP addresses must be valid static external IP addresses assigned to the project. This field is used for public NAT. */
	// +kcc:proto:field=google.cloud.compute.v1.RouterNatRuleAction.source_nat_active_ips
	// +optional
	SourceNatActiveIpsRefs []refs.ComputeAddressRef `json:"sourceNatActiveIpsRefs,omitempty"`

	/* A list of URLs of the subnetworks used as source ranges for this NAT Rule. These subnetworks must have purpose set to PRIVATE_NAT. This field is used for private NAT. */
	// +kcc:proto:field=google.cloud.compute.v1.RouterNatRuleAction.source_nat_active_ranges
	// +optional
	SourceNatActiveRangesRefs []refs.ComputeSubnetworkRef `json:"sourceNatActiveRangesRefs,omitempty"`

	/* A list of URLs of the IP resources to be drained. These IPs must be valid static external IPs that have been assigned to the NAT. These IPs should be used for updating/patching a NAT rule only. This field is used for public NAT. */
	// +kcc:proto:field=google.cloud.compute.v1.RouterNatRuleAction.source_nat_drain_ips
	// +optional
	SourceNatDrainIpsRefs []refs.ComputeAddressRef `json:"sourceNatDrainIpsRefs,omitempty"`

	/* A list of URLs of subnetworks representing source ranges to be drained. This is only supported on patch/update, and these subnetworks must have previously been used as active ranges in this NAT Rule. This field is used for private NAT. */
	// +kcc:proto:field=google.cloud.compute.v1.RouterNatRuleAction.source_nat_drain_ranges
	// +optional
	SourceNatDrainRangesRefs []refs.ComputeSubnetworkRef `json:"sourceNatDrainRangesRefs,omitempty"`
}

// +kcc:proto=google.cloud.compute.v1.RouterNatRule
type RouterNatRule struct {
	/* The action to be enforced for traffic that matches this rule. */
	// +optional
	Action *RouterNatRuleAction `json:"action,omitempty"`

	/* An optional description of this rule. */
	// +optional
	Description *string `json:"description,omitempty"`

	/* CEL expression that specifies the match condition that egress traffic from a VM is evaluated against. If it evaluates to true, the corresponding `action` is enforced. */
	// +kcc:proto:field=google.cloud.compute.v1.RouterNatRule.match
	Match *string `json:"match"`

	/* An integer uniquely identifying a rule in the list. The rule number must be a positive value between 0 and 65000, and must be unique among rules within a NAT. */
	// +kcc:proto:field=google.cloud.compute.v1.RouterNatRule.rule_number
	RuleNumber *uint32 `json:"ruleNumber"`
}

// +kcc:proto=google.cloud.compute.v1.RouterNatSubnetworkToNat
type RouterNatSubnetwork struct {
	/* The subnetwork to NAT. */
	SubnetworkRef refs.ComputeSubnetworkRef `json:"subnetworkRef"`

	/* List of the secondary ranges of the subnetwork that are allowed to use NAT. This can be populated only if "LIST_OF_SECONDARY_IP_RANGES" is one of the values in sourceIpRangesToNat. */
	// +optional
	SecondaryIpRangeNames []string `json:"secondaryIpRangeNames,omitempty"`

	/* List of options for which source IPs in the subnetwork should have NAT enabled. Supported values include: 'ALL_IP_RANGES', 'LIST_OF_SECONDARY_IP_RANGES', 'PRIMARY_IP_RANGE'. */
	// +kcc:proto:field=google.cloud.compute.v1.RouterNatSubnetworkToNat.source_ip_ranges_to_nat
	SourceIpRangesToNat []string `json:"sourceIpRangesToNat"`
}

// +kcc:proto=google.cloud.compute.v1.RouterNatSubnetworkToNat64
type RouterNat64Subnetwork struct {
	/* The subnetwork to NAT. */
	SubnetworkRef refs.ComputeSubnetworkRef `json:"subnetworkRef"`
}

// +kcc:spec:proto=google.cloud.compute.v1.RouterNat
type ComputeRouterNATSpec struct {
	/* The network tier to use when automatically reserving NAT IP addresses. Must be one of: PREMIUM, STANDARD. If not specified, then the current project-level default tier is used. Possible values: ["PREMIUM", "STANDARD"] */
	// +kcc:proto:field=google.cloud.compute.v1.RouterNat.auto_network_tier
	// +optional
	AutoNetworkTier *string `json:"autoNetworkTier,omitempty"`

	/* A list of URLs of the IP resources to be drained. These IPs must be valid static external IPs that have been assigned to the NAT. These IPs should be used for updating/patching a NAT only. */
	// +kcc:proto:field=google.cloud.compute.v1.RouterNat.drain_nat_ips
	// +optional
	DrainNatIps []refs.ComputeAddressRef `json:"drainNatIps,omitempty"`

	/* Enable Dynamic Port Allocation. If not specified, it is disabled by default. */
	// +kcc:proto:field=google.cloud.compute.v1.RouterNat.enable_dynamic_port_allocation
	// +optional
	EnableDynamicPortAllocation *bool `json:"enableDynamicPortAllocation,omitempty"`

	/* Specifies if endpoint independent mapping is enabled. This is enabled by default. */
	// +kcc:proto:field=google.cloud.compute.v1.RouterNat.enable_endpoint_independent_mapping
	// +optional
	EnableEndpointIndependentMapping *bool `json:"enableEndpointIndependentMapping,omitempty"`

	/* List of NAT-ted endpoint types supported by the Nat Gateway. If the list is empty, then it will be equivalent to include ENDPOINT_TYPE_VM. Possible values: ["ENDPOINT_TYPE_VM", "ENDPOINT_TYPE_SWG"] */
	// +kcc:proto:field=google.cloud.compute.v1.RouterNat.endpoint_types
	// +optional
	EndpointTypes []string `json:"endpointTypes,omitempty"`

	/* Timeout (in seconds) for ICMP connections. Defaults to 30s if not set. */
	// +kcc:proto:field=google.cloud.compute.v1.RouterNat.icmp_idle_timeout_sec
	// +optional
	IcmpIdleTimeoutSec *int32 `json:"icmpIdleTimeoutSec,omitempty"`

	/* Configure logging on this NAT. */
	// +kcc:proto:field=google.cloud.compute.v1.RouterNat.log_config
	// +optional
	LogConfig *RouterNatLogConfig `json:"logConfig,omitempty"`

	/* Maximum number of ports allocated to a VM from this NAT config when Dynamic Port Allocation is enabled. */
	// +kcc:proto:field=google.cloud.compute.v1.RouterNat.max_ports_per_vm
	// +optional
	MaxPortsPerVm *int32 `json:"maxPortsPerVm,omitempty"`

	/* Minimum number of ports allocated to a VM from this NAT config. If not set, a default number of ports is allocated to a VM. */
	// +kcc:proto:field=google.cloud.compute.v1.RouterNat.min_ports_per_vm
	// +optional
	MinPortsPerVm *int32 `json:"minPortsPerVm,omitempty"`

	/* How external IPs should be allocated for this NAT. Valid values are 'AUTO_ONLY' for only allowing NAT IPs allocated by Google Cloud Platform, or 'MANUAL_ONLY' for only user-allocated NAT IP addresses. Possible values: ["MANUAL_ONLY", "AUTO_ONLY"]. */
	// +kcc:proto:field=google.cloud.compute.v1.RouterNat.nat_ip_allocate_option
	// +optional
	NatIpAllocateOption *string `json:"natIpAllocateOption,omitempty"`

	/* A list of URLs of the IP resources used for this Nat service. These IP addresses must be valid static external IP addresses assigned to the project. */
	// +kcc:proto:field=google.cloud.compute.v1.RouterNat.nat_ips
	// +optional
	NatIps []refs.ComputeAddressRef `json:"natIps,omitempty"`

	/* Immutable. Region where the router and NAT reside. */
	Region string `json:"region"`

	/* Immutable. Optional. The name of the resource. Used for creation and acquisition. When unset, the value of `metadata.name` is used as the default. */
	// +optional
	ResourceID *string `json:"resourceID,omitempty"`

	/* The Cloud Router in which this NAT will be configured. */
	RouterRef refs.ComputeRouterRef `json:"routerRef"`

	/* A list of rules associated with this NAT. */
	// +kcc:proto:field=google.cloud.compute.v1.RouterNat.rules
	// +optional
	Rules []RouterNatRule `json:"rules,omitempty"`

	/* How NAT should be configured per Subnetwork. Possible values: ["ALL_SUBNETWORKS_ALL_IP_RANGES", "ALL_SUBNETWORKS_ALL_PRIMARY_IP_RANGES", "LIST_OF_SUBNETWORKS"]. */
	// +kcc:proto:field=google.cloud.compute.v1.RouterNat.source_subnetwork_ip_ranges_to_nat
	SourceSubnetworkIpRangesToNat string `json:"sourceSubnetworkIpRangesToNat"`

	/* How NAT64 should be configured per Subnetwork. Possible values: ["ALL_IPV6_SUBNETWORKS", "LIST_OF_IPV6_SUBNETWORKS"]. */
	// +kcc:proto:field=google.cloud.compute.v1.RouterNat.source_subnetwork_ip_ranges_to_nat64
	// +optional
	SourceSubnetworkIpRangesToNat64 *string `json:"sourceSubnetworkIpRangesToNat64,omitempty"`

	/* A list of Subnetwork resources whose traffic should be translated by NAT Gateway. It is used only when LIST_OF_SUBNETWORKS is selected for the SubnetworkIpRangeToNatOption above. */
	// +kcc:proto:field=google.cloud.compute.v1.RouterNat.subnetworks
	// +optional
	Subnetwork []RouterNatSubnetwork `json:"subnetwork,omitempty"`

	/* A list of Subnetwork resources whose traffic should be translated by NAT64 Gateway. It is used only when LIST_OF_IPV6_SUBNETWORKS is selected for the SubnetworkIpRangeToNat64Option above. */
	// +kcc:proto:field=google.cloud.compute.v1.RouterNat.nat64_subnetworks
	// +optional
	Nat64Subnetworks []RouterNat64Subnetwork `json:"nat64Subnetworks,omitempty"`

	/* Timeout (in seconds) for TCP established connections. Defaults to 1200s if not set. */
	// +kcc:proto:field=google.cloud.compute.v1.RouterNat.tcp_established_idle_timeout_sec
	// +optional
	TcpEstablishedIdleTimeoutSec *int32 `json:"tcpEstablishedIdleTimeoutSec,omitempty"`

	/* Timeout (in seconds) for TCP connections that are in TIME_WAIT state. Defaults to 120s if not set. */
	// +kcc:proto:field=google.cloud.compute.v1.RouterNat.tcp_time_wait_timeout_sec
	// +optional
	TcpTimeWaitTimeoutSec *int32 `json:"tcpTimeWaitTimeoutSec,omitempty"`

	/* Timeout (in seconds) for TCP transitory connections. Defaults to 30s if not set. */
	// +kcc:proto:field=google.cloud.compute.v1.RouterNat.tcp_transitory_idle_timeout_sec
	// +optional
	TcpTransitoryIdleTimeoutSec *int32 `json:"tcpTransitoryIdleTimeoutSec,omitempty"`

	/* Indicates whether this NAT is used for public or private IP translation. If unspecified, it defaults to PUBLIC. Possible values: ["PUBLIC", "PRIVATE"]. */
	// +kcc:proto:field=google.cloud.compute.v1.RouterNat.type
	// +optional
	Type *string `json:"type,omitempty"`

	/* Timeout (in seconds) for UDP connections. Defaults to 30s if not set. */
	// +kcc:proto:field=google.cloud.compute.v1.RouterNat.udp_idle_timeout_sec
	// +optional
	UdpIdleTimeoutSec *int32 `json:"udpIdleTimeoutSec,omitempty"`
}

type ComputeRouterNATStatus struct {
	commonv1alpha1.CommonStatus `json:",inline"`
}

// +genclient
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:resource:categories=gcp,shortName=gcpcomputerouternat;gcpcomputerouternats
// +kubebuilder:subresource:status
// +kubebuilder:metadata:labels="cnrm.cloud.google.com/managed-by-kcc=true"
// +kubebuilder:metadata:labels="cnrm.cloud.google.com/stability-level=stable"
// +kubebuilder:metadata:labels="cnrm.cloud.google.com/system=true"
// +kubebuilder:metadata:labels="cnrm.cloud.google.com/tf2crd=true"
// +kubebuilder:metadata:labels="cnrm.cloud.google.com/default-controller=direct"
// +kubebuilder:printcolumn:name="Age",JSONPath=".metadata.creationTimestamp",type="date"
// +kubebuilder:printcolumn:name="Ready",JSONPath=".status.conditions[?(@.type=='Ready')].status",type="string",description="When 'True', the most recent reconcile of the resource succeeded"
// +kubebuilder:printcolumn:name="Status",JSONPath=".status.conditions[?(@.type=='Ready')].reason",type="string",description="The reason for the value in 'Ready'"
// +kubebuilder:printcolumn:name="Status Age",JSONPath=".status.conditions[?(@.type=='Ready')].lastTransitionTime",type="date",description="The last transition time for the value in 'Status'"

// ComputeRouterNAT is the Schema for the compute API
// +k8s:openapi-gen=true
type ComputeRouterNAT struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   ComputeRouterNATSpec   `json:"spec,omitempty"`
	Status ComputeRouterNATStatus `json:"status,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// ComputeRouterNATList contains a list of ComputeRouterNAT
type ComputeRouterNATList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []ComputeRouterNAT `json:"items"`
}

func init() {
	SchemeBuilder.Register(&ComputeRouterNAT{}, &ComputeRouterNATList{})
}
