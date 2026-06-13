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
	refsv1beta1 "github.com/GoogleCloudPlatform/k8s-config-connector/apis/refs/v1beta1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/apis/k8s/v1alpha1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

var ComputeRouterNATGVK = GroupVersion.WithKind("ComputeRouterNAT")

// ComputeRouterNATSpec defines the desired state of ComputeRouterNAT
// +kcc:spec:proto=google.cloud.compute.v1.RouterNat
type ComputeRouterNATSpec struct {
	// The project that this resource belongs to.
	// +required
	ProjectRef *refsv1beta1.ProjectRef `json:"projectRef"`

	// Immutable. Region where the router and NAT reside.
	// +required
	Region *string `json:"region"`

	// The Cloud Router in which this NAT will be configured.
	// +required
	RouterRef *ComputeRouterRef `json:"routerRef"`

	// Immutable. Optional. The name of the resource. Used for
	// creation and acquisition. When unset, the value of `metadata.name`
	// is used as the default.
	ResourceID *string `json:"resourceID,omitempty"`

	// +kcc:proto:field=google.cloud.compute.v1.RouterNat.drain_nat_ips
	DrainNatIps []refsv1beta1.ComputeAddressRef `json:"drainNatIps,omitempty"`

	// +kcc:proto:field=google.cloud.compute.v1.RouterNat.enable_dynamic_port_allocation
	EnableDynamicPortAllocation *bool `json:"enableDynamicPortAllocation,omitempty"`

	// +kcc:proto:field=google.cloud.compute.v1.RouterNat.enable_endpoint_independent_mapping
	EnableEndpointIndependentMapping *bool `json:"enableEndpointIndependentMapping,omitempty"`

	// +kcc:proto:field=google.cloud.compute.v1.RouterNat.icmp_idle_timeout_sec
	IcmpIdleTimeoutSec *int32 `json:"icmpIdleTimeoutSec,omitempty"`

	// +kcc:proto:field=google.cloud.compute.v1.RouterNat.log_config
	LogConfig *RouterNATLogConfig `json:"logConfig,omitempty"`

	// +kcc:proto:field=google.cloud.compute.v1.RouterNat.max_ports_per_vm
	MaxPortsPerVM *int32 `json:"maxPortsPerVm,omitempty"`

	// +kcc:proto:field=google.cloud.compute.v1.RouterNat.min_ports_per_vm
	MinPortsPerVM *int32 `json:"minPortsPerVm,omitempty"`

	// +required
	// +kcc:proto:field=google.cloud.compute.v1.RouterNat.nat_ip_allocate_option
	NatIpAllocateOption *string `json:"natIpAllocateOption"`

	// +kcc:proto:field=google.cloud.compute.v1.RouterNat.nat_ips
	NatIps []refsv1beta1.ComputeAddressRef `json:"natIps,omitempty"`

	// +kcc:proto:field=google.cloud.compute.v1.RouterNat.rules
	Rules []RouterNATRule `json:"rules,omitempty"`

	// +kcc:proto:field=google.cloud.compute.v1.RouterNat.source_subnetwork_ip_ranges_to_nat
	SourceSubnetworkIpRangesToNat *string `json:"sourceSubnetworkIpRangesToNat,omitempty"`

	// +kcc:proto:field=google.cloud.compute.v1.RouterNat.subnetworks
	Subnetwork []RouterNATSubnetwork `json:"subnetwork,omitempty"`

	// +kcc:proto:field=google.cloud.compute.v1.RouterNat.tcp_established_idle_timeout_sec
	TcpEstablishedIdleTimeoutSec *int32 `json:"tcpEstablishedIdleTimeoutSec,omitempty"`

	// +kcc:proto:field=google.cloud.compute.v1.RouterNat.tcp_time_wait_timeout_sec
	TcpTimeWaitTimeoutSec *int32 `json:"tcpTimeWaitTimeoutSec,omitempty"`

	// +kcc:proto:field=google.cloud.compute.v1.RouterNat.tcp_transitory_idle_timeout_sec
	TcpTransitoryIdleTimeoutSec *int32 `json:"tcpTransitoryIdleTimeoutSec,omitempty"`

	// +kcc:proto:field=google.cloud.compute.v1.RouterNat.udp_idle_timeout_sec
	UdpIdleTimeoutSec *int32 `json:"udpIdleTimeoutSec,omitempty"`
}

// +kcc:proto=google.cloud.compute.v1.RouterNatLogConfig
type RouterNATLogConfig struct {
	// +required
	// +kcc:proto:field=google.cloud.compute.v1.RouterNatLogConfig.enable
	Enable *bool `json:"enable"`

	// +required
	// +kcc:proto:field=google.cloud.compute.v1.RouterNatLogConfig.filter
	Filter *string `json:"filter"`
}

// +kcc:proto=google.cloud.compute.v1.RouterNatRule
type RouterNATRule struct {
	// +kcc:proto:field=google.cloud.compute.v1.RouterNatRule.action
	Action *RouterNATRuleAction `json:"action,omitempty"`

	// +kcc:proto:field=google.cloud.compute.v1.RouterNatRule.description
	Description *string `json:"description,omitempty"`

	// +required
	// +kcc:proto:field=google.cloud.compute.v1.RouterNatRule.match
	Match *string `json:"match"`

	// +required
	// +kcc:proto:field=google.cloud.compute.v1.RouterNatRule.rule_number
	RuleNumber *uint32 `json:"ruleNumber"`
}

// +kcc:proto=google.cloud.compute.v1.RouterNatRuleAction
type RouterNATRuleAction struct {
	// +kcc:proto:field=google.cloud.compute.v1.RouterNatRuleAction.source_nat_active_ips
	SourceNatActiveIpsRefs []refsv1beta1.ComputeAddressRef `json:"sourceNatActiveIpsRefs,omitempty"`

	// +kcc:proto:field=google.cloud.compute.v1.RouterNatRuleAction.source_nat_drain_ips
	SourceNatDrainIpsRefs []refsv1beta1.ComputeAddressRef `json:"sourceNatDrainIpsRefs,omitempty"`
}

// +kcc:proto=google.cloud.compute.v1.RouterNatSubnetworkToNat
type RouterNATSubnetwork struct {
	// +required
	// +kcc:proto:field=google.cloud.compute.v1.RouterNatSubnetworkToNat.name
	SubnetworkRef *ComputeSubnetworkRef `json:"subnetworkRef"`

	// +kcc:proto:field=google.cloud.compute.v1.RouterNatSubnetworkToNat.secondary_ip_range_names
	SecondaryIpRangeNames []string `json:"secondaryIpRangeNames,omitempty"`

	// +required
	// +kcc:proto:field=google.cloud.compute.v1.RouterNatSubnetworkToNat.source_ip_ranges_to_nat
	SourceIpRangesToNat []string `json:"sourceIpRangesToNat"`
}

// ComputeRouterNATStatus defines the config connector machine state of ComputeRouterNAT
type ComputeRouterNATStatus struct {
	/* Conditions represent the latest available observations of the
	   object's current state. */
	Conditions []v1alpha1.Condition `json:"conditions,omitempty"`

	// ObservedGeneration is the generation of the resource that was most recently observed by the Config Connector controller. If this is equal to metadata.generation, then that means that the current reported status reflects the most recent desired state of the resource.
	ObservedGeneration *int64 `json:"observedGeneration,omitempty"`

	// A unique specifier for the ComputeRouterNAT resource in GCP.
	ExternalRef *string `json:"externalRef,omitempty"`

	// ObservedState is the state of the resource as most recently observed in GCP.
	ObservedState *ComputeRouterNATObservedState `json:"observedState,omitempty"`
}

// ComputeRouterNATObservedState is the state of the ComputeRouterNAT resource as most recently observed in GCP.
// +kcc:observedstate:proto=google.cloud.compute.v1.RouterNat
type ComputeRouterNATObservedState struct {
}

// +genclient
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:resource:categories=gcp,shortName=gcpcomputerouternat;gcpcomputerouternats
// +kubebuilder:subresource:status
// +kubebuilder:metadata:labels="cnrm.cloud.google.com/managed-by-kcc=true"
// +kubebuilder:metadata:labels="cnrm.cloud.google.com/stability-level=alpha"
// +kubebuilder:metadata:labels="cnrm.cloud.google.com/system=true"
// +kubebuilder:printcolumn:name="Age",JSONPath=".metadata.creationTimestamp",type="date"
// +kubebuilder:printcolumn:name="Ready",JSONPath=".status.conditions[?(@.type=='Ready')].status",type="string",description="When 'True', the most recent reconcile of the resource succeeded"
// +kubebuilder:printcolumn:name="Status",JSONPath=".status.conditions[?(@.type=='Ready')].reason",type="string",description="The reason for the value in 'Ready'"
// +kubebuilder:printcolumn:name="Status Age",JSONPath=".status.conditions[?(@.type=='Ready')].lastTransitionTime",type="date",description="The last transition time for the value in 'Status'"

// ComputeRouterNAT is the Schema for the ComputeRouterNAT API
// +k8s:openapi-gen=true
type ComputeRouterNAT struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	// +required
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
