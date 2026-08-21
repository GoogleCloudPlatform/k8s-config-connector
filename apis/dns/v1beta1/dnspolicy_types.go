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
	computerefs "github.com/GoogleCloudPlatform/k8s-config-connector/apis/compute/refs"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/apis/k8s/v1alpha1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

var DNSPolicyGVK = GroupVersion.WithKind("DNSPolicy")

// +openapi:PolicyAlternativeNameServerConfig
type PolicyAlternativeNameServerConfig struct {
	/* Sets an alternative name server for the associated networks. When specified,
	all DNS queries are forwarded to a name server that you choose. Names such as .internal
	are not available when an alternative name server is specified. */
	TargetNameServers []PolicyTargetNameServers `json:"targetNameServers"`
}

// +openapi:PolicyAlternativeNameServerConfigTargetNameServer
type PolicyTargetNameServers struct {
	/* Forwarding path for this TargetNameServer. If unset or 'default' Cloud DNS will make forwarding
	decision based on address ranges, i.e. RFC1918 addresses go to the VPC, Non-RFC1918 addresses go
	to the Internet. When set to 'private', Cloud DNS will always send queries through VPC for this target Possible values: ["default", "private"]. */
	// +optional
	ForwardingPath *string `json:"forwardingPath,omitempty"`

	/* IPv4 address to forward to. */
	Ipv4Address string `json:"ipv4Address"`
}

// +openapi:PolicyNetwork
type PolicyNetworks struct {
	/* VPC network to bind to. */
	NetworkRef computerefs.ComputeNetworkRef `json:"networkRef"`
}

// +openapi:Policy
type DNSPolicySpec struct {
	/* Sets an alternative name server for the associated networks.
	When specified, all DNS queries are forwarded to a name server that you choose.
	Names such as .internal are not available when an alternative name server is specified. */
	// +optional
	AlternativeNameServerConfig *PolicyAlternativeNameServerConfig `json:"alternativeNameServerConfig,omitempty"`

	/* A textual description field. Defaults to 'Managed by Config Connector'. */
	// +optional
	Description *string `json:"description,omitempty"`

	/* Allows networks bound to this policy to receive DNS queries sent
	by VMs or applications over VPN connections. When enabled, a
	virtual IP address will be allocated from each of the sub-networks
	that are bound to this policy. */
	// +optional
	EnableInboundForwarding *bool `json:"enableInboundForwarding,omitempty"`

	/* Controls whether logging is enabled for the networks bound to this policy.
	Defaults to no logging if not set. */
	// +optional
	EnableLogging *bool `json:"enableLogging,omitempty"`

	/* List of network names specifying networks to which this policy is applied. */
	// +optional
	Networks []PolicyNetworks `json:"networks,omitempty"`

	/* Immutable. Optional. The name of the resource. Used for creation and acquisition. When unset, the value of `metadata.name` is used as the default. */
	// +optional
	ResourceID *string `json:"resourceID,omitempty"`
}

type DNSPolicyStatus struct {
	/* Conditions represent the latest available observations of the
	   DNSPolicy's current state. */
	Conditions []v1alpha1.Condition `json:"conditions,omitempty"`
	/* ObservedGeneration is the generation of the resource that was most recently observed by the Config Connector controller. If this is equal to metadata.generation, then that means that the current reported status reflects the most recent desired state of the resource. */
	// +optional
	ObservedGeneration *int64 `json:"observedGeneration,omitempty"`
}

// +genclient
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:resource:categories=gcp,shortName=gcpdnspolicy;gcpdnspolicies
// +kubebuilder:subresource:status
// +kubebuilder:metadata:labels="cnrm.cloud.google.com/managed-by-kcc=true"
// +kubebuilder:metadata:labels="cnrm.cloud.google.com/stability-level=stable"
// +kubebuilder:metadata:labels="cnrm.cloud.google.com/system=true"
// +kubebuilder:metadata:labels="cnrm.cloud.google.com/tf2crd=true"
// +kubebuilder:printcolumn:name="Age",JSONPath=".metadata.creationTimestamp",type="date"
// +kubebuilder:printcolumn:name="Ready",JSONPath=".status.conditions[?(@.type=='Ready')].status",type="string",description="When 'True', the most recent reconcile of the resource succeeded"
// +kubebuilder:printcolumn:name="Status",JSONPath=".status.conditions[?(@.type=='Ready')].reason",type="string",description="The reason for the value in 'Ready'"
// +kubebuilder:printcolumn:name="Status Age",JSONPath=".status.conditions[?(@.type=='Ready')].lastTransitionTime",type="date",description="The last transition time for the value in 'Status'"

// DNSPolicy is the Schema for the dns API
// +k8s:openapi-gen=true
type DNSPolicy struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   DNSPolicySpec   `json:"spec,omitempty"`
	Status DNSPolicyStatus `json:"status,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// DNSPolicyList contains a list of DNSPolicy
type DNSPolicyList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []DNSPolicy `json:"items"`
}

func init() {
	SchemeBuilder.Register(&DNSPolicy{}, &DNSPolicyList{})
}
