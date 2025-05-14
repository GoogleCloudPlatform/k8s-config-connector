// Copyright 2025 Google LLC
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
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/apis/k8s/v1alpha1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

var VMwareEngineExternalAccessRuleGVK = GroupVersion.WithKind("VMwareEngineExternalAccessRule")

// +kcc:proto=google.cloud.vmwareengine.v1.ExternalAccessRule.IpRange
type ExternalAccessRule_IPRange struct {
	// A single IP address. For example: `10.0.0.5`.
	// +kcc:proto:field=google.cloud.vmwareengine.v1.ExternalAccessRule.IpRange.ip_address
	IPAddress *string `json:"ipAddress,omitempty"`

	// An IP address range in the CIDR format. For example: `10.0.0.0/24`.
	// +kcc:proto:field=google.cloud.vmwareengine.v1.ExternalAccessRule.IpRange.ip_address_range
	IPAddressRange *string `json:"ipAddressRange,omitempty"`

	// The name of an `ExternalAddress` resource. The external address must
	//  have been reserved in the scope of this external access rule's parent
	//  network policy.
	// +kcc:proto:field=google.cloud.vmwareengine.v1.ExternalAccessRule.IpRange.external_address
	ExternalAddressRef *ExternalAddressRef `json:"externalAddressRef,omitempty"`
}

// VMwareEngineExternalAccessRuleSpec defines the desired state of VMwareEngineExternalAccessRule
// +kcc:spec:proto=google.cloud.vmwareengine.v1.ExternalAccessRule
type VMwareEngineExternalAccessRuleSpec struct {
	// The VMwareEngineExternalAccessRule name. If not given, the metadata.name will be used.
	ResourceID *string `json:"resourceID,omitempty"`

	// Required. The resource name of the network policy to create a new external access firewall rule in.
	// +required
	NetworkPolicyRef *NetworkPolicyRef `json:"networkPolicyRef"`

	// User-provided description for this external access rule.
	// +kcc:proto:field=google.cloud.vmwareengine.v1.ExternalAccessRule.description
	Description *string `json:"description,omitempty"`

	// External access rule priority, which determines the external access rule to
	//  use when multiple rules apply. If multiple rules have the same priority,
	//  their ordering is non-deterministic. If specific ordering is required,
	//  assign unique priorities to enforce such ordering. The external access rule
	//  priority is an integer from 100 to 4096, both inclusive. Lower integers
	//  indicate higher precedence. For example, a rule with priority `100` has
	//  higher precedence than a rule with priority `101`.
	// +kcc:proto:field=google.cloud.vmwareengine.v1.ExternalAccessRule.priority
	Priority *int32 `json:"priority,omitempty"`

	// The action that the external access rule performs.
	// +kcc:proto:field=google.cloud.vmwareengine.v1.ExternalAccessRule.action
	Action *string `json:"action,omitempty"`

	// The IP protocol to which the external access rule applies. This value can
	//  be one of the following three protocol strings (not case-sensitive):
	//  `tcp`, `udp`, or `icmp`.
	// +kcc:proto:field=google.cloud.vmwareengine.v1.ExternalAccessRule.ip_protocol
	IPProtocol *string `json:"ipProtocol,omitempty"`

	// If source ranges are specified, the external access rule applies only to
	//  traffic that has a source IP address in these ranges. These ranges can
	//  either be expressed in the CIDR format or as an IP address. As only inbound
	//  rules are supported, `ExternalAddress` resources cannot be the source IP
	//  addresses of an external access rule. To match all source addresses,
	//  specify `0.0.0.0/0`.
	// +kcc:proto:field=google.cloud.vmwareengine.v1.ExternalAccessRule.source_ip_ranges
	SourceIPRanges []ExternalAccessRule_IPRange `json:"sourceIPRanges,omitempty"`

	// A list of source ports to which the external access rule applies. This
	//  field is only applicable for the UDP or TCP protocol.
	//  Each entry must be either an integer or a range. For example: `["22"]`,
	//  `["80","443"]`, or `["12345-12349"]`. To match all source ports, specify
	//  `["0-65535"]`.
	// +kcc:proto:field=google.cloud.vmwareengine.v1.ExternalAccessRule.source_ports
	SourcePorts []string `json:"sourcePorts,omitempty"`

	// If destination ranges are specified, the external access rule applies only
	//  to the traffic that has a destination IP address in these ranges. The
	//  specified IP addresses must have reserved external IP addresses in the
	//  scope of the parent network policy. To match all external IP addresses in
	//  the scope of the parent network policy, specify `0.0.0.0/0`. To match a
	//  specific external IP address, specify it using the
	//  `IpRange.external_address` property.
	// +kcc:proto:field=google.cloud.vmwareengine.v1.ExternalAccessRule.destination_ip_ranges
	DestinationIPRanges []ExternalAccessRule_IPRange `json:"destinationIPRanges,omitempty"`

	// A list of destination ports to which the external access rule applies. This
	//  field is only applicable for the UDP or TCP protocol.
	//  Each entry must be either an integer or a range. For example: `["22"]`,
	//  `["80","443"]`, or `["12345-12349"]`. To match all destination ports,
	//  specify `["0-65535"]`.
	// +kcc:proto:field=google.cloud.vmwareengine.v1.ExternalAccessRule.destination_ports
	DestinationPorts []string `json:"destinationPorts,omitempty"`
}

// VMwareEngineExternalAccessRuleStatus defines the config connector machine state of VMwareEngineExternalAccessRule
type VMwareEngineExternalAccessRuleStatus struct {
	/* Conditions represent the latest available observations of the
	   object's current state. */
	Conditions []v1alpha1.Condition `json:"conditions,omitempty"`

	// ObservedGeneration is the generation of the resource that was most recently observed by the Config Connector controller. If this is equal to metadata.generation, then that means that the current reported status reflects the most recent desired state of the resource.
	ObservedGeneration *int64 `json:"observedGeneration,omitempty"`

	// A unique specifier for the VMwareEngineExternalAccessRule resource in GCP.
	ExternalRef *string `json:"externalRef,omitempty"`

	// ObservedState is the state of the resource as most recently observed in GCP.
	ObservedState *VMwareEngineExternalAccessRuleObservedState `json:"observedState,omitempty"`
}

// VMwareEngineExternalAccessRuleObservedState is the state of the VMwareEngineExternalAccessRule resource as most recently observed in GCP.
// +kcc:observedstate:proto=google.cloud.vmwareengine.v1.ExternalAccessRule
type VMwareEngineExternalAccessRuleObservedState struct {
	// Output only. The resource name of this external access rule.
	//  Resource names are schemeless URIs that follow the conventions in
	//  https://cloud.google.com/apis/design/resource_names.
	//  For example:
	//  `projects/my-project/locations/us-central1/networkPolicies/my-policy/externalAccessRules/my-rule`
	// +kcc:proto:field=google.cloud.vmwareengine.v1.ExternalAccessRule.name
	// NOTYET: this field serves the same purpose as externalRef
	// Name *string `json:"name,omitempty"`

	// Output only. Creation time of this resource.
	// +kcc:proto:field=google.cloud.vmwareengine.v1.ExternalAccessRule.create_time
	CreateTime *string `json:"createTime,omitempty"`

	// Output only. Last update time of this resource.
	// +kcc:proto:field=google.cloud.vmwareengine.v1.ExternalAccessRule.update_time
	UpdateTime *string `json:"updateTime,omitempty"`

	// Output only. The state of the resource.
	// +kcc:proto:field=google.cloud.vmwareengine.v1.ExternalAccessRule.state
	State *string `json:"state,omitempty"`

	// Output only. System-generated unique identifier for the resource.
	// +kcc:proto:field=google.cloud.vmwareengine.v1.ExternalAccessRule.uid
	UID *string `json:"uid,omitempty"`
}

// +genclient
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:resource:categories=gcp,shortName=gcpvmwareengineexternalaccessrule;gcpvmwareengineexternalaccessrules
// +kubebuilder:subresource:status
// +kubebuilder:metadata:labels="cnrm.cloud.google.com/managed-by-kcc=true";"cnrm.cloud.google.com/system=true"
// +kubebuilder:printcolumn:name="Age",JSONPath=".metadata.creationTimestamp",type="date"
// +kubebuilder:printcolumn:name="Ready",JSONPath=".status.conditions[?(@.type=='Ready')].status",type="string",description="When 'True', the most recent reconcile of the resource succeeded"
// +kubebuilder:printcolumn:name="Status",JSONPath=".status.conditions[?(@.type=='Ready')].reason",type="string",description="The reason for the value in 'Ready'"
// +kubebuilder:printcolumn:name="Status Age",JSONPath=".status.conditions[?(@.type=='Ready')].lastTransitionTime",type="date",description="The last transition time for the value in 'Status'"

// VMwareEngineExternalAccessRule is the Schema for the VMwareEngineExternalAccessRule API
// +k8s:openapi-gen=true
type VMwareEngineExternalAccessRule struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	// +required
	Spec   VMwareEngineExternalAccessRuleSpec   `json:"spec,omitempty"`
	Status VMwareEngineExternalAccessRuleStatus `json:"status,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// VMwareEngineExternalAccessRuleList contains a list of VMwareEngineExternalAccessRule
type VMwareEngineExternalAccessRuleList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []VMwareEngineExternalAccessRule `json:"items"`
}

func init() {
	SchemeBuilder.Register(&VMwareEngineExternalAccessRule{}, &VMwareEngineExternalAccessRuleList{})
}
