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

package v1alpha1

import (
	refsv1beta1 "github.com/GoogleCloudPlatform/k8s-config-connector/apis/refs/v1beta1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/apis/k8s/v1alpha1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

var ComputeNetworkFirewallPolicyRuleGVK = GroupVersion.WithKind("ComputeNetworkFirewallPolicyRule")

// +kcc:proto=google.cloud.compute.v1.FirewallPolicyRuleSecureTag
type FirewallPolicyRuleSecureTag struct {
	// Name of the secure tag, created with TagManager's TagValue API.
	// +kcc:proto:field=google.cloud.compute.v1.FirewallPolicyRuleSecureTag.name
	Name *string `json:"name,omitempty"`

	// [Output Only] State of the secure tag, either `EFFECTIVE` or `INEFFECTIVE`. A secure tag is `INEFFECTIVE` when it is deleted or its network is deleted.
	// +kcc:proto:field=google.cloud.compute.v1.FirewallPolicyRuleSecureTag.state
	State *string `json:"state,omitempty"`
}

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

	/* List of secure tag values, which should be matched at the source of the traffic. For INGRESS rule, if all the <code>srcSecureTag</code> are INEFFECTIVE, and there is no <code>srcIpRange</code>, this rule will be ignored. Maximum number of source tag values allowed is 256. */
	// +kcc:proto:field=google.cloud.compute.v1.FirewallPolicyRuleMatcher.src_secure_tags
	SrcSecureTags []FirewallPolicyRuleSecureTag `json:"srcSecureTags,omitempty"`

	/* Name of the Google Cloud Threat Intelligence list. */
	// +kcc:proto:field=google.cloud.compute.v1.FirewallPolicyRuleMatcher.src_threat_intelligences
	SrcThreatIntelligences []string `json:"srcThreatIntelligences,omitempty"`
}

// ComputeNetworkFirewallPolicyRuleSpec defines the desired state of ComputeNetworkFirewallPolicyRule
// +kcc:spec:proto=google.cloud.compute.v1.FirewallPolicyRule
type ComputeNetworkFirewallPolicyRuleSpec struct {
	// The project that this resource belongs to.
	// +required
	ProjectRef *refsv1beta1.ProjectRef `json:"projectRef"`

	// The ComputeNetworkFirewallPolicyRule name / priority. If not given, the metadata.name will be used.
	// +optional
	ResourceID *string `json:"resourceID,omitempty"`

	/* The Action to perform when the client connection triggers the rule. Valid actions are "allow", "deny" and "goto_next". */
	// +required
	// +kcc:proto:field=google.cloud.compute.v1.FirewallPolicyRule.action
	Action string `json:"action"`

	/* An optional description for this resource. */
	// +optional
	// +kcc:proto:field=google.cloud.compute.v1.FirewallPolicyRule.description
	Description *string `json:"description,omitempty"`

	/* The direction in which this rule applies. Possible values: INGRESS, EGRESS. */
	// +required
	// +kcc:proto:field=google.cloud.compute.v1.FirewallPolicyRule.direction
	Direction string `json:"direction"`

	/* Denotes whether the firewall policy rule is disabled. When set to true, the firewall policy rule is not enforced and traffic behaves as if it did not exist. If this is unspecified, the firewall policy rule will be enabled. */
	// +optional
	// +kcc:proto:field=google.cloud.compute.v1.FirewallPolicyRule.disabled
	Disabled *bool `json:"disabled,omitempty"`

	/* Denotes whether to enable logging for a particular rule. If logging is enabled, logs will be exported to the configured export destination in Stackdriver. Logs may be exported to BigQuery or Pub/Sub. Note: you cannot enable logging on "goto_next" rules. */
	// +optional
	// +kcc:proto:field=google.cloud.compute.v1.FirewallPolicyRule.enable_logging
	EnableLogging *bool `json:"enableLogging,omitempty"`

	/* The firewall policy that this rule belongs to. */
	// +required
	FirewallPolicyRef *refsv1beta1.ComputeNetworkFirewallPolicyRef `json:"firewallPolicyRef"`

	/* A match condition that incoming traffic is evaluated against. If it evaluates to true, the corresponding 'action' is enforced. */
	// +required
	// +kcc:proto:field=google.cloud.compute.v1.FirewallPolicyRule.match
	Match *FirewallPolicyRuleMatcher `json:"match"`

	/* An optional name for the rule. This field is not a unique identifier and can be updated. */
	// +optional
	// +kcc:proto:field=google.cloud.compute.v1.FirewallPolicyRule.rule_name
	RuleName *string `json:"ruleName,omitempty"`

	/* A list of secure tags that controls which instances the firewall rule applies to. If <code>targetSecureTag</code> are specified, then the firewall rule applies only to instances in the VPC network that have one of those EFFECTIVE secure tags, if all the target_secure_tag are in INEFFECTIVE state, then this rule will be ignored. <code>targetSecureTag</code> may not be set at the same time as <code>targetServiceAccounts</code>. If neither <code>targetServiceAccounts</code> nor <code>targetSecureTag</code> are specified, the firewall rule applies to all instances on the specified network. Maximum number of target label tags allowed is 256. */
	// +optional
	// +kcc:proto:field=google.cloud.compute.v1.FirewallPolicyRule.target_secure_tags
	TargetSecureTags []FirewallPolicyRuleSecureTag `json:"targetSecureTags,omitempty"`

	// +optional
	// +kcc:proto:field=google.cloud.compute.v1.FirewallPolicyRule.target_service_accounts
	TargetServiceAccountRefs []*refsv1beta1.IAMServiceAccountRef `json:"targetServiceAccountRefs,omitempty"`
}

// ComputeNetworkFirewallPolicyRuleStatus defines the config connector machine state of ComputeNetworkFirewallPolicyRule
type ComputeNetworkFirewallPolicyRuleStatus struct {
	/* Conditions represent the latest available observations of the
	   object's current state. */
	Conditions []v1alpha1.Condition `json:"conditions,omitempty"`

	// ObservedGeneration is the generation of the resource that was most recently observed by the Config Connector controller. If this is equal to metadata.generation, then that means that the current reported status reflects the most recent desired state of the resource.
	ObservedGeneration *int64 `json:"observedGeneration,omitempty"`

	// A unique specifier for the ComputeNetworkFirewallPolicyRule resource in GCP.
	ExternalRef *string `json:"externalRef,omitempty"`

	// ObservedState is the state of the resource as most recently observed in GCP.
	ObservedState *ComputeNetworkFirewallPolicyRuleObservedState `json:"observedState,omitempty"`
}

// ComputeNetworkFirewallPolicyRuleObservedState is the state of the ComputeNetworkFirewallPolicyRule resource as most recently observed in GCP.
// +kcc:observedstate:proto=google.cloud.compute.v1.FirewallPolicyRule
type ComputeNetworkFirewallPolicyRuleObservedState struct {
	/* Type of the resource. Always `compute#firewallPolicyRule` for firewall policy rules */
	// +kcc:proto:field=google.cloud.compute.v1.FirewallPolicyRule.kind
	Kind *string `json:"kind,omitempty"`

	/* Calculation of the complexity of a single firewall policy rule. */
	// +kcc:proto:field=google.cloud.compute.v1.FirewallPolicyRule.rule_tuple_count
	RuleTupleCount *int64 `json:"ruleTupleCount,omitempty"`
}

// +genclient
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:resource:categories=gcp,shortName=gcpcomputenetworkfirewallpolicyrule;gcpcomputenetworkfirewallpolicyrules
// +kubebuilder:subresource:status
// +kubebuilder:metadata:labels="cnrm.cloud.google.com/managed-by-kcc=true"
// +kubebuilder:metadata:labels="cnrm.cloud.google.com/system=true"
// +kubebuilder:metadata:labels="cnrm.cloud.google.com/stability-level=alpha"
// +kubebuilder:printcolumn:name="Age",JSONPath=".metadata.creationTimestamp",type="date"
// +kubebuilder:printcolumn:name="Ready",JSONPath=".status.conditions[?(@.type=='Ready')].status",type="string",description="When 'True', the most recent reconcile of the resource succeeded"
// +kubebuilder:printcolumn:name="Status",JSONPath=".status.conditions[?(@.type=='Ready')].reason",type="string",description="The reason for the value in 'Ready'"
// +kubebuilder:printcolumn:name="Status Age",JSONPath=".status.conditions[?(@.type=='Ready')].lastTransitionTime",type="date",description="The last transition time for the value in 'Status'"

// ComputeNetworkFirewallPolicyRule is the Schema for the ComputeNetworkFirewallPolicyRule API
// +k8s:openapi-gen=true
type ComputeNetworkFirewallPolicyRule struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	// +required
	Spec   ComputeNetworkFirewallPolicyRuleSpec   `json:"spec,omitempty"`
	Status ComputeNetworkFirewallPolicyRuleStatus `json:"status,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// ComputeNetworkFirewallPolicyRuleList contains a list of ComputeNetworkFirewallPolicyRule
type ComputeNetworkFirewallPolicyRuleList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []ComputeNetworkFirewallPolicyRule `json:"items"`
}

func init() {
	SchemeBuilder.Register(&ComputeNetworkFirewallPolicyRule{}, &ComputeNetworkFirewallPolicyRuleList{})
}
