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

var ComputeFirewallGVK = GroupVersion.WithKind("ComputeFirewall")

// +kcc:proto=google.cloud.compute.v1.Allowed
type FirewallAllow struct {
	// +optional
	// +kcc:proto:field=google.cloud.compute.v1.Allowed.ports
	Ports []string `json:"ports,omitempty"`

	// +required
	// +kcc:proto:field=google.cloud.compute.v1.Allowed.I_p_protocol
	Protocol string `json:"protocol"`
}

// +kcc:proto=google.cloud.compute.v1.Denied
type FirewallDeny struct {
	// +optional
	// +kcc:proto:field=google.cloud.compute.v1.Denied.ports
	Ports []string `json:"ports,omitempty"`

	// +required
	// +kcc:proto:field=google.cloud.compute.v1.Denied.I_p_protocol
	Protocol string `json:"protocol"`
}

// +kcc:proto=google.cloud.compute.v1.FirewallLogConfig
type FirewallLogConfig struct {
	// +required
	// +kcc:proto:field=google.cloud.compute.v1.FirewallLogConfig.metadata
	Metadata string `json:"metadata"`
}

// ComputeFirewallSpec defines the desired state of ComputeFirewall
// +kcc:spec:proto=google.cloud.compute.v1.Firewall
type ComputeFirewallSpec struct {
	// The project that this resource belongs to.
	ProjectRef *refsv1beta1.ProjectRef `json:"projectRef"`

	// The ComputeFirewall name. If not given, the metadata.name will be used.
	// +kcc:proto:field=google.cloud.compute.v1.Firewall.name
	ResourceID *string `json:"resourceID,omitempty"`

	// +optional
	// +kcc:proto:field=google.cloud.compute.v1.Firewall.allowed
	Allow []FirewallAllow `json:"allow,omitempty"`

	// +optional
	// +kcc:proto:field=google.cloud.compute.v1.Firewall.denied
	Deny []FirewallDeny `json:"deny,omitempty"`

	// +optional
	// +kcc:proto:field=google.cloud.compute.v1.Firewall.description
	Description *string `json:"description,omitempty"`

	// +optional
	// +kcc:proto:field=google.cloud.compute.v1.Firewall.destination_ranges
	DestinationRanges []string `json:"destinationRanges,omitempty"`

	// +optional
	// +kcc:proto:field=google.cloud.compute.v1.Firewall.direction
	Direction *string `json:"direction,omitempty"`

	// +optional
	// +kcc:proto:field=google.cloud.compute.v1.Firewall.disabled
	Disabled *bool `json:"disabled,omitempty"`

	// +optional
	// +kcc:proto:field=google.cloud.compute.v1.Firewall.log_config.enable
	EnableLogging *bool `json:"enableLogging,omitempty"`

	// +optional
	// +kcc:proto:field=google.cloud.compute.v1.Firewall.log_config
	LogConfig *FirewallLogConfig `json:"logConfig,omitempty"`

	// +required
	// +kcc:proto:field=google.cloud.compute.v1.Firewall.network
	NetworkRef *ComputeNetworkRef `json:"networkRef"`

	// +optional
	// +kcc:proto:field=google.cloud.compute.v1.Firewall.priority
	Priority *int64 `json:"priority,omitempty"`

	// +optional
	// +kcc:proto:field=google.cloud.compute.v1.Firewall.source_ranges
	SourceRanges []string `json:"sourceRanges,omitempty"`

	// +optional
	// +kcc:proto:field=google.cloud.compute.v1.Firewall.source_service_accounts
	SourceServiceAccounts []*refsv1beta1.IAMServiceAccountRef `json:"sourceServiceAccounts,omitempty"`

	// +optional
	// +kcc:proto:field=google.cloud.compute.v1.Firewall.source_tags
	SourceTags []string `json:"sourceTags,omitempty"`

	// +optional
	// +kcc:proto:field=google.cloud.compute.v1.Firewall.target_service_accounts
	TargetServiceAccounts []*refsv1beta1.IAMServiceAccountRef `json:"targetServiceAccounts,omitempty"`

	// +optional
	// +kcc:proto:field=google.cloud.compute.v1.Firewall.target_tags
	TargetTags []string `json:"targetTags,omitempty"`
}

// ComputeFirewallStatus defines the config connector machine state of ComputeFirewall
type ComputeFirewallStatus struct {
	/* Conditions represent the latest available observations of the
	   object's current state. */
	Conditions []v1alpha1.Condition `json:"conditions,omitempty"`

	// ObservedGeneration is the generation of the resource that was most recently observed by the Config Connector controller. If this is equal to metadata.generation, then that means that the current reported status reflects the most recent desired state of the resource.
	ObservedGeneration *int64 `json:"observedGeneration,omitempty"`

	// A unique specifier for the ComputeFirewall resource in GCP.
	ExternalRef *string `json:"externalRef,omitempty"`

	// ObservedState is the state of the resource as most recently observed in GCP.
	ObservedState *ComputeFirewallObservedState `json:"observedState,omitempty"`
}

// ComputeFirewallObservedState is the state of the ComputeFirewall resource as most recently observed in GCP.
// +kcc:observedstate:proto=google.cloud.compute.v1.Firewall
type ComputeFirewallObservedState struct {
	// +optional
	// +kcc:proto:field=google.cloud.compute.v1.Firewall.creation_timestamp
	CreationTimestamp *string `json:"creationTimestamp,omitempty"`

	// +optional
	// +kcc:proto:field=google.cloud.compute.v1.Firewall.self_link
	SelfLink *string `json:"selfLink,omitempty"`
}

// +genclient
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:resource:categories=gcp,shortName=gcpcomputefirewall;gcpcomputefirewalls
// +kubebuilder:subresource:status
// +kubebuilder:metadata:labels="cnrm.cloud.google.com/managed-by-kcc=true"
// +kubebuilder:metadata:labels="cnrm.cloud.google.com/stability-level=stable"
// +kubebuilder:metadata:labels="cnrm.cloud.google.com/system=true"
// +kubebuilder:printcolumn:name="Age",JSONPath=".metadata.creationTimestamp",type="date"
// +kubebuilder:printcolumn:name="Ready",JSONPath=".status.conditions[?(@.type=='Ready')].status",type="string",description="When 'True', the most recent reconcile of the resource succeeded"
// +kubebuilder:printcolumn:name="Status",JSONPath=".status.conditions[?(@.type=='Ready')].reason",type="string",description="The reason for the value in 'Ready'"
// +kubebuilder:printcolumn:name="Status Age",JSONPath=".status.conditions[?(@.type=='Ready')].lastTransitionTime",type="date",description="The last transition time for the value in 'Status'"

// ComputeFirewall is the Schema for the ComputeFirewall API
// +k8s:openapi-gen=true
type ComputeFirewall struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	// +required
	Spec   ComputeFirewallSpec   `json:"spec,omitempty"`
	Status ComputeFirewallStatus `json:"status,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// ComputeFirewallList contains a list of ComputeFirewall
type ComputeFirewallList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []ComputeFirewall `json:"items"`
}

func init() {
	SchemeBuilder.Register(&ComputeFirewall{}, &ComputeFirewallList{})
}
