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

var ComputeOrganizationSecurityPolicyRuleGVK = GroupVersion.WithKind("ComputeOrganizationSecurityPolicyRule")

// +kcc:proto=google.cloud.compute.v1.SecurityPolicyRuleMatcherConfigLayer4Config
type SecurityPolicyRuleMatcherConfigLayer4Config struct {
	/* The IP protocol to which this rule applies. The protocol type is required when creating a firewall rule. This value can either be one of the following well known protocol strings (`tcp`, `udp`, `icmp`, `esp`, `ah`, `ipip`, `sctp`), or the IP protocol number. */
	// +kcc:proto:field=google.cloud.compute.v1.SecurityPolicyRuleMatcherConfigLayer4Config.ip_protocol
	IPProtocol string `json:"ipProtocol"`

	/* An optional list of ports to which this rule applies. This field is only applicable for UDP or TCP protocol. Each entry must be either an integer or a range. If not specified, this rule applies to connections through any port. */
	// +kcc:proto:field=google.cloud.compute.v1.SecurityPolicyRuleMatcherConfigLayer4Config.ports
	Ports []string `json:"ports,omitempty"`
}

// +kcc:proto=google.cloud.compute.v1.SecurityPolicyRuleMatcherConfig
type SecurityPolicyRuleMatcherConfig struct {
	/* Destination IP address range in CIDR format. Required for EGRESS rules. */
	// +kcc:proto:field=google.cloud.compute.v1.SecurityPolicyRuleMatcherConfig.dest_ip_ranges
	// +optional
	DestIPRanges []string `json:"destIPRanges,omitempty"`

	/* Pairs of IP protocols and ports that the rule should match. */
	// +kcc:proto:field=google.cloud.compute.v1.SecurityPolicyRuleMatcherConfig.layer4_configs
	Layer4Configs []SecurityPolicyRuleMatcherConfigLayer4Config `json:"layer4Configs,omitempty"`

	/* Source IP address range in CIDR format. Required for INGRESS rules. */
	// +kcc:proto:field=google.cloud.compute.v1.SecurityPolicyRuleMatcherConfig.src_ip_ranges
	// +optional
	SrcIPRanges []string `json:"srcIPRanges,omitempty"`
}

// +kcc:proto=google.cloud.compute.v1.SecurityPolicyRuleMatcher
type SecurityPolicyRuleMatcher struct {
	/* The configuration options for matching the rule. */
	// +kcc:proto:field=google.cloud.compute.v1.SecurityPolicyRuleMatcher.config
	Config *SecurityPolicyRuleMatcherConfig `json:"config,omitempty"`

	/* A description of the rule. */
	// +kcc:proto:field=google.cloud.compute.v1.SecurityPolicyRuleMatcher.description
	// +optional
	Description *string `json:"description,omitempty"`

	/* Preconfigured versioned expression. For organization security policy rules, the only supported type is "FIREWALL". Default value: "FIREWALL" Possible values: ["FIREWALL"]. */
	// +kcc:proto:field=google.cloud.compute.v1.SecurityPolicyRuleMatcher.versioned_expr
	// +optional
	VersionedExpr *string `json:"versionedExpr,omitempty"`
}

// ComputeOrganizationSecurityPolicyRuleSpec defines the desired state of ComputeOrganizationSecurityPolicyRule
// +kcc:spec:proto=google.cloud.compute.v1.SecurityPolicyRule
type ComputeOrganizationSecurityPolicyRuleSpec struct {
	// The project that this resource belongs to.
	// +required
	ProjectRef *refsv1beta1.ProjectRef `json:"projectRef"`

	// The location of this resource.
	// +required
	Location string `json:"location"`

	// The ComputeOrganizationSecurityPolicyRule name. If not given, the metadata.name will be used.
	// +optional
	ResourceID *string `json:"resourceID,omitempty"`

	/* The Action to perform when the client connection triggers the rule. Can currently be either "allow", "deny" or "goto_next". */
	// +kcc:proto:field=google.cloud.compute.v1.SecurityPolicyRule.action
	// +required
	Action string `json:"action"`

	/* A description of the rule. */
	// +kcc:proto:field=google.cloud.compute.v1.SecurityPolicyRule.description
	// +optional
	Description *string `json:"description,omitempty"`

	/* The direction in which this rule applies. If unspecified an INGRESS rule is created. Possible values: ["INGRESS", "EGRESS"]. */
	// +kcc:proto:field=google.cloud.compute.v1.SecurityPolicyRule.direction
	// +optional
	Direction *string `json:"direction,omitempty"`

	/* Denotes whether to enable logging for a particular rule. If logging is enabled, logs will be exported to the configured export destination in Stackdriver. */
	// +kcc:proto:field=google.cloud.compute.v1.SecurityPolicyRule.enable_logging
	// +optional
	EnableLogging *bool `json:"enableLogging,omitempty"`

	/* A match condition that incoming traffic is evaluated against. If it evaluates to true, the corresponding 'action' is enforced. */
	// +kcc:proto:field=google.cloud.compute.v1.SecurityPolicyRule.match
	// +required
	Match *SecurityPolicyRuleMatcher `json:"match"`

	/* Immutable. The ID of the OrganizationSecurityPolicy this rule applies to. */
	// +required
	PolicyRef *ComputeOrganizationSecurityPolicyRef `json:"policyRef"`

	/* If set to true, the specified action is not enforced. */
	// +kcc:proto:field=google.cloud.compute.v1.SecurityPolicyRule.preview
	// +optional
	Preview *bool `json:"preview,omitempty"`

	/* A list of network resource URLs to which this rule applies. This field allows you to control which network's VMs get this rule. If this field is left blank, all VMs within the organization will receive the rule. */
	// +kcc:proto:field=google.cloud.compute.v1.SecurityPolicyRule.target_resources
	// +optional
	TargetResources []string `json:"targetResources,omitempty"`

	/* A list of service accounts indicating the sets of instances that are applied with this rule. */
	// +kcc:proto:field=google.cloud.compute.v1.SecurityPolicyRule.target_service_accounts
	// +optional
	TargetServiceAccounts []string `json:"targetServiceAccounts,omitempty"`
}

// ComputeOrganizationSecurityPolicyRuleStatus defines the config connector machine state of ComputeOrganizationSecurityPolicyRule
type ComputeOrganizationSecurityPolicyRuleStatus struct {
	/* Conditions represent the latest available observations of the
	   object's current state. */
	Conditions []v1alpha1.Condition `json:"conditions,omitempty"`

	// ObservedGeneration is the generation of the resource that was most recently observed by the Config Connector controller. If this is equal to metadata.generation, then that means that the current reported status reflects the most recent desired state of the resource.
	ObservedGeneration *int64 `json:"observedGeneration,omitempty"`

	// A unique specifier for the ComputeOrganizationSecurityPolicyRule resource in GCP.
	ExternalRef *string `json:"externalRef,omitempty"`

	// ObservedState is the state of the resource as most recently observed in GCP.
	ObservedState *ComputeOrganizationSecurityPolicyRuleObservedState `json:"observedState,omitempty"`
}

// ComputeOrganizationSecurityPolicyRuleObservedState is the state of the ComputeOrganizationSecurityPolicyRule resource as most recently observed in GCP.
// +kcc:observedstate:proto=google.cloud.compute.v1.SecurityPolicyRule
type ComputeOrganizationSecurityPolicyRuleObservedState struct {
}

// +genclient
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:resource:categories=gcp,shortName=gcpcomputeorganizationsecuritypolicyrule;gcpcomputeorganizationsecuritypolicyrules
// +kubebuilder:subresource:status
// +kubebuilder:metadata:labels="cnrm.cloud.google.com/managed-by-kcc=true"
// +kubebuilder:metadata:labels="cnrm.cloud.google.com/stability-level=alpha"
// +kubebuilder:metadata:labels="cnrm.cloud.google.com/system=true"
// +kubebuilder:printcolumn:name="Age",JSONPath=".metadata.creationTimestamp",type="date"
// +kubebuilder:printcolumn:name="Ready",JSONPath=".status.conditions[?(@.type=='Ready')].status",type="string",description="When 'True', the most recent reconcile of the resource succeeded"
// +kubebuilder:printcolumn:name="Status",JSONPath=".status.conditions[?(@.type=='Ready')].reason",type="string",description="The reason for the value in 'Ready'"
// +kubebuilder:printcolumn:name="Status Age",JSONPath=".status.conditions[?(@.type=='Ready')].lastTransitionTime",type="date",description="The last transition time for the value in 'Status'"

// ComputeOrganizationSecurityPolicyRule is the Schema for the ComputeOrganizationSecurityPolicyRule API
// +k8s:openapi-gen=true
type ComputeOrganizationSecurityPolicyRule struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	// +required
	Spec   ComputeOrganizationSecurityPolicyRuleSpec   `json:"spec,omitempty"`
	Status ComputeOrganizationSecurityPolicyRuleStatus `json:"status,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// ComputeOrganizationSecurityPolicyRuleList contains a list of ComputeOrganizationSecurityPolicyRule
type ComputeOrganizationSecurityPolicyRuleList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []ComputeOrganizationSecurityPolicyRule `json:"items"`
}

func init() {
	SchemeBuilder.Register(&ComputeOrganizationSecurityPolicyRule{}, &ComputeOrganizationSecurityPolicyRuleList{})
}
