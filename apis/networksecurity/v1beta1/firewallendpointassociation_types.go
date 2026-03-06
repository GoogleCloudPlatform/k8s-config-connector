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
	computev1beta1 "github.com/GoogleCloudPlatform/k8s-config-connector/apis/compute/v1beta1"
	k8sv1alpha1 "github.com/GoogleCloudPlatform/k8s-config-connector/pkg/apis/k8s/v1alpha1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

var NetworkSecurityFirewallEndpointAssociationGVK = GroupVersion.WithKind("NetworkSecurityFirewallEndpointAssociation")

// NetworkSecurityFirewallEndpointAssociationSpec defines the desired state of NetworkSecurityFirewallEndpointAssociation
// +kcc:spec:proto=google.cloud.networksecurity.v1beta1.FirewallEndpointAssociation
type NetworkSecurityFirewallEndpointAssociationSpec struct {
	// The NetworkSecurityFirewallEndpointAssociation name. If not given, the metadata.name will be used.
	// +kubebuilder:validation:XValidation:rule="self == oldSelf",message="Field is immutable"
	ResourceID *string `json:"resourceID,omitempty"`

	Parent `json:",inline"`

	// Optional. Whether the association is disabled. True indicates that traffic won't be intercepted.
	// +kcc:proto:field=google.cloud.networksecurity.v1beta1.FirewallEndpointAssociation.disabled
	Disabled *bool `json:"disabled,omitempty"`

	// Required. The URL of the FirewallEndpoint that is being associated.
	// +kcc:proto:field=google.cloud.networksecurity.v1beta1.FirewallEndpointAssociation.firewall_endpoint
	// +required
	// +kubebuilder:validation:XValidation:rule="self == oldSelf",message="Field is immutable"
	FirewallEndpointRef *NetworkSecurityFirewallEndpointRef `json:"firewallEndpointRef,omitempty"`

	// Required. The URL of the network that is being associated.
	// +kcc:proto:field=google.cloud.networksecurity.v1beta1.FirewallEndpointAssociation.network
	// +required
	// +kubebuilder:validation:XValidation:rule="self == oldSelf",message="Field is immutable"
	NetworkRef *computev1beta1.ComputeNetworkRef `json:"networkRef,omitempty"`

	// Optional. Labels as key value pairs
	// +kcc:proto:field=google.cloud.networksecurity.v1beta1.FirewallEndpointAssociation.labels
	Labels map[string]string `json:"labels,omitempty"`

	// Optional. The URL of the TlsInspectionPolicy that is being associated.
	// +kcc:proto:field=google.cloud.networksecurity.v1beta1.FirewallEndpointAssociation.tls_inspection_policy
	// +kubebuilder:validation:XValidation:rule="self == oldSelf",message="Field is immutable"
	TlsInspectionPolicyRef *NetworkSecurityTlsInspectionPolicyRef `json:"tlsInspectionPolicyRef,omitempty"`
}

// NetworkSecurityFirewallEndpointAssociationStatus defines the config connector machine state of NetworkSecurityFirewallEndpointAssociation
type NetworkSecurityFirewallEndpointAssociationStatus struct {
	/* Conditions represent the latest available observations of the
	   object's current state. */
	Conditions []k8sv1alpha1.Condition `json:"conditions,omitempty"`

	// ObservedGeneration is the generation of the resource that was most recently observed by the Config Connector controller. If this is equal to metadata.generation, then that means that the current reported status reflects the most recent desired state of the resource.
	ObservedGeneration *int64 `json:"observedGeneration,omitempty"`

	// A unique specifier for the NetworkSecurityFirewallEndpointAssociation resource in GCP.
	ExternalRef *string `json:"externalRef,omitempty"`

	// ObservedState is the state of the resource as most recently observed in GCP.
	ObservedState *NetworkSecurityFirewallEndpointAssociationObservedState `json:"observedState,omitempty"`
}

// NetworkSecurityFirewallEndpointAssociationObservedState defines the observed state of NetworkSecurityFirewallEndpointAssociation
// +kcc:observedstate:proto=google.cloud.networksecurity.v1beta1.FirewallEndpointAssociation
type NetworkSecurityFirewallEndpointAssociationObservedState struct {
	// Output only. Create time stamp
	// +kcc:proto:field=google.cloud.networksecurity.v1beta1.FirewallEndpointAssociation.create_time
	CreateTime *string `json:"createTime,omitempty"`

	// Output only. Update time stamp
	// +kcc:proto:field=google.cloud.networksecurity.v1beta1.FirewallEndpointAssociation.update_time
	UpdateTime *string `json:"updateTime,omitempty"`

	// Output only. Current state of the association.
	// +kcc:proto:field=google.cloud.networksecurity.v1beta1.FirewallEndpointAssociation.state
	State *string `json:"state,omitempty"`

	// Output only. Whether reconciling is in progress.
	// +kcc:proto:field=google.cloud.networksecurity.v1beta1.FirewallEndpointAssociation.reconciling
	Reconciling *bool `json:"reconciling,omitempty"`
}

// +genclient
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:resource:categories=gcp,shortName=gcpnetworksecurityfirewallendpointassociation;gcpnetworksecurityfirewallendpointassociations
// +kubebuilder:subresource:status
// +kubebuilder:metadata:labels="cnrm.cloud.google.com/managed-by-kcc=true"
// +kubebuilder:metadata:labels="cnrm.cloud.google.com/system=true"
// +kubebuilder:metadata:labels="cnrm.cloud.google.com/direct=true"
// +kubebuilder:printcolumn:name="Age",JSONPath=".metadata.creationTimestamp",type="date"
// +kubebuilder:printcolumn:name="Ready",JSONPath=".status.conditions[?(@.type=='Ready')].status",type="string",description="When 'True', the most recent reconcile of the resource succeeded"
// +kubebuilder:printcolumn:name="Status",JSONPath=".status.conditions[?(@.type=='Ready')].reason",type="string",description="The reason for the value in 'Ready'"
// +kubebuilder:printcolumn:name="Status Age",JSONPath=".status.conditions[?(@.type=='Ready')].lastTransitionTime",type="date",description="The last transition time for the value in 'Status'"

// NetworkSecurityFirewallEndpointAssociation is the Schema for the NetworkSecurityFirewallEndpointAssociation API
// +k8s:openapi-gen=true
// +kubebuilder:storageversion
type NetworkSecurityFirewallEndpointAssociation struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	// +required
	Spec   NetworkSecurityFirewallEndpointAssociationSpec   `json:"spec,omitempty"`
	Status NetworkSecurityFirewallEndpointAssociationStatus `json:"status,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// NetworkSecurityFirewallEndpointAssociationList contains a list of NetworkSecurityFirewallEndpointAssociation
type NetworkSecurityFirewallEndpointAssociationList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []NetworkSecurityFirewallEndpointAssociation `json:"items"`
}

func init() {
	SchemeBuilder.Register(&NetworkSecurityFirewallEndpointAssociation{}, &NetworkSecurityFirewallEndpointAssociationList{})
}
