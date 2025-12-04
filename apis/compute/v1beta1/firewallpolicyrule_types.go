// Copyright 2024 Google LLC
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
)

var (
	ComputeFirewallPolicyRuleGVK = GroupVersion.WithKind("ComputeFirewallPolicyRule")
)

// +kcc:proto=google.cloud.compute.v1.FirewallPolicyRuleMatcherLayer4Config
type FirewallPolicyRuleMatcherLayer4Config struct {
	/* The IP protocol to which this rule applies. The protocol type is required when creating a firewall rule. This value can either be one of the following well known protocol strings (`tcp`, `udp`, `icmp`, `esp`, `ah`, `ipip`, `sctp`), or the IP protocol number. */
	// +kcc:proto:field=google.cloud.compute.v1.FirewallPolicyRuleMatcherLayer4Config.ip_protocol
	IPProtocol string `json:"ipProtocol"`

	/* An optional list of ports to which this rule applies. This field is only applicable for UDP or TCP protocol. Each entry must be either an integer or a range. If not specified, this rule applies to connections through any port. Example inputs include: ``. */
	// +kcc:proto:field=google.cloud.compute.v1.FirewallPolicyRuleMatcherLayer4Config.ports
	Ports []string `json:"ports,omitempty"`
}

// +kcc:proto=google.cloud.compute.v1.FirewallPolicyRuleMatcher
type FirewallPolicyRuleMatcher struct {
	/* Address groups which should be matched against the traffic destination. Maximum number of destination address groups is 10. Destination address groups is only supported in Egress rules. */
	// +kcc:proto:field=google.cloud.compute.v1.FirewallPolicyRuleMatcher.dest_address_groups
	DestAddressGroups []string `json:"destAddressGroups,omitempty"`

	/* Domain names that will be used to match against the resolved domain name of destination of traffic. Can only be specified if DIRECTION is egress. */
	// +kcc:proto:field=google.cloud.compute.v1.FirewallPolicyRuleMatcher.dest_fqdns
	DestFqdns []string `json:"destFqdns,omitempty"`

	/* CIDR IP address range. Maximum number of destination CIDR IP ranges allowed is 256. */
	// +kcc:proto:field=google.cloud.compute.v1.FirewallPolicyRuleMatcher.dest_ip_ranges
	DestIPRanges []string `json:"destIPRanges,omitempty"`

	/* The Unicode country codes whose IP addresses will be used to match against the source of traffic. Can only be specified if DIRECTION is egress. */
	// +kcc:proto:field=google.cloud.compute.v1.FirewallPolicyRuleMatcher.dest_region_codes
	DestRegionCodes []string `json:"destRegionCodes,omitempty"`

	/* Name of the Google Cloud Threat Intelligence list. */
	// +kcc:proto:field=google.cloud.compute.v1.FirewallPolicyRuleMatcher.dest_threat_intelligences
	DestThreatIntelligences []string `json:"destThreatIntelligences,omitempty"`

	/* Pairs of IP protocols and ports that the rule should match. */
	// +kcc:proto:field=google.cloud.compute.v1.FirewallPolicyRuleMatcher.layer4_configs
	Layer4Configs []FirewallPolicyRuleMatcherLayer4Config `json:"layer4Configs"`

	/* Address groups which should be matched against the traffic source. Maximum number of source address groups is 10. Source address groups is only supported in Ingress rules. */
	// +kcc:proto:field=google.cloud.compute.v1.FirewallPolicyRuleMatcher.src_address_groups
	SrcAddressGroups []string `json:"srcAddressGroups,omitempty"`

	/* Domain names that will be used to match against the resolved domain name of source of traffic. Can only be specified if DIRECTION is ingress. */
	// +kcc:proto:field=google.cloud.compute.v1.FirewallPolicyRuleMatcher.src_fqdns
	SrcFqdns []string `json:"srcFqdns,omitempty"`

	/* CIDR IP address range. Maximum number of source CIDR IP ranges allowed is 256. */
	// +kcc:proto:field=google.cloud.compute.v1.FirewallPolicyRuleMatcher.src_ip_ranges
	SrcIPRanges []string `json:"srcIPRanges,omitempty"`

	/* The Unicode country codes whose IP addresses will be used to match against the source of traffic. Can only be specified if DIRECTION is ingress. */
	// +kcc:proto:field=google.cloud.compute.v1.FirewallPolicyRuleMatcher.src_region_codes
	SrcRegionCodes []string `json:"srcRegionCodes,omitempty"`

	/* Name of the Google Cloud Threat Intelligence list. */
	// +kcc:proto:field=google.cloud.compute.v1.FirewallPolicyRuleMatcher.src_threat_intelligences
	SrcThreatIntelligences []string `json:"srcThreatIntelligences,omitempty"`
}

// +kcc:spec:proto=google.cloud.compute.v1.FirewallPolicyRule
type ComputeFirewallPolicyRuleSpec struct {
	/* The Action to perform when the client connection triggers the rule. Valid actions are "allow", "deny" and "goto_next". */
	// +kcc:proto:field=google.cloud.compute.v1.FirewallPolicyRule.action
	Action string `json:"action"`

	/* An optional description for this resource. */
	// +kcc:proto:field=google.cloud.compute.v1.FirewallPolicyRule.description
	Description *string `json:"description,omitempty"`

	/* The direction in which this rule applies. Possible values: INGRESS, EGRESS */
	// +kcc:proto:field=google.cloud.compute.v1.FirewallPolicyRule.direction
	Direction string `json:"direction"`

	/* Denotes whether the firewall policy rule is disabled. When set to true, the firewall policy rule is not enforced and traffic behaves as if it did not exist. If this is unspecified, the firewall policy rule will be enabled. */
	// +kcc:proto:field=google.cloud.compute.v1.FirewallPolicyRule.disabled
	Disabled *bool `json:"disabled,omitempty"`

	/* Denotes whether to enable logging for a particular rule. If logging is enabled, logs will be exported to the configured export destination in Stackdriver. Logs may be exported to BigQuery or Pub/Sub. Note: you cannot enable logging on "goto_next" rules. */
	// +kcc:proto:field=google.cloud.compute.v1.FirewallPolicyRule.enable_logging
	EnableLogging *bool `json:"enableLogging,omitempty"`

	/* Immutable. */
	FirewallPolicyRef *refs.ComputeFirewallPolicyRef `json:"firewallPolicyRef"`

	/* A match condition that incoming traffic is evaluated against. If it evaluates to true, the corresponding 'action' is enforced. */
	// +kcc:proto:field=google.cloud.compute.v1.FirewallPolicyRule.match
	Match *FirewallPolicyRuleMatcher `json:"match"`

	/* Immutable. An integer indicating the priority of a rule in the list. The priority must be a positive value between 0 and 2147483647. Rules are evaluated from highest to lowest priority where 0 is the highest priority and 2147483647 is the lowest prority. */
	// +kcc:proto:field=google.cloud.compute.v1.FirewallPolicyRule.priority
	Priority int64 `json:"priority"`

	// +kcc:proto:field=google.cloud.compute.v1.FirewallPolicyRule.target_resources
	TargetResources []*ComputeNetworkRef `json:"targetResources,omitempty"`

	// +kcc:proto:field=google.cloud.compute.v1.FirewallPolicyRule.target_service_accounts
	TargetServiceAccounts []*refs.IAMServiceAccountRef `json:"targetServiceAccounts,omitempty"`
}

// +kcc:status:proto=google.cloud.compute.v1.FirewallPolicyRule
type ComputeFirewallPolicyRuleStatus struct {
	commonv1alpha1.CommonStatus `json:",inline"`
	/* Type of the resource. Always `compute#firewallPolicyRule` for firewall policy rules */
	// +kcc:proto:field=google.cloud.compute.v1.FirewallPolicyRule.kind
	Kind *string `json:"kind,omitempty"`

	/* Calculation of the complexity of a single firewall policy rule. */
	// +kcc:proto:field=google.cloud.compute.v1.FirewallPolicyRule.rule_tuple_count
	RuleTupleCount *int64 `json:"ruleTupleCount,omitempty"`
}

// +genclient
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:resource:categories=gcp,shortName=gcpcomputefirewallpolicyrule;gcpcomputefirewallpolicyrules
// +kubebuilder:subresource:status
// +kubebuilder:metadata:labels="cnrm.cloud.google.com/managed-by-kcc=true";"cnrm.cloud.google.com/stability-level=stable";"cnrm.cloud.google.com/system=true"
// +kubebuilder:printcolumn:name="Age",JSONPath=".metadata.creationTimestamp",type="date"
// +kubebuilder:printcolumn:name="Ready",JSONPath=".status.conditions[?(@.type=='Ready')].status",type="string",description="When 'True', the most recent reconcile of the resource succeeded"
// +kubebuilder:printcolumn:name="Status",JSONPath=".status.conditions[?(@.type=='Ready')].reason",type="string",description="The reason for the value in 'Ready'"
// +kubebuilder:printcolumn:name="Status Age",JSONPath=".status.conditions[?(@.type=='Ready')].lastTransitionTime",type="date",description="The last transition time for the value in 'Status'"

// ComputeFirewallPolicyRule is the Schema for the compute API
// +k8s:openapi-gen=true
type ComputeFirewallPolicyRule struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   ComputeFirewallPolicyRuleSpec   `json:"spec,omitempty"`
	Status ComputeFirewallPolicyRuleStatus `json:"status,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// ComputeFirewallPolicyRuleList contains a list of ComputeFirewallPolicyRule
type ComputeFirewallPolicyRuleList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []ComputeFirewallPolicyRule `json:"items"`
}

func init() {
	SchemeBuilder.Register(&ComputeFirewallPolicyRule{}, &ComputeFirewallPolicyRuleList{})
}
