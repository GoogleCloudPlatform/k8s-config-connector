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
	k8sv1alpha1 "github.com/GoogleCloudPlatform/k8s-config-connector/pkg/apis/k8s/v1alpha1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

var NetworkSecurityFirewallEndpointGVK = GroupVersion.WithKind("NetworkSecurityFirewallEndpoint")

// NetworkSecurityFirewallEndpointSpec defines the desired state of NetworkSecurityFirewallEndpoint
// +kcc:spec:proto=google.cloud.networksecurity.v1beta1.FirewallEndpoint
type NetworkSecurityFirewallEndpointSpec struct {
	// The Organization that this resource belongs to.
	// +required
	OrganizationRef *refsv1beta1.OrganizationRef `json:"organizationRef"`

	// Immutable. The location for the resource
	// +required
	Location string `json:"location"`

	// The NetworkSecurityFirewallEndpoint name. If not given, the metadata.name will be used.
	// +kubebuilder:validation:XValidation:rule="self == oldSelf",message="Field is immutable"
	ResourceID *string `json:"resourceID,omitempty"`

	// Required. Project to bill on endpoint uptime usage.
	// +kcc:proto:field=google.cloud.networksecurity.v1beta1.FirewallEndpoint.billing_project_id
	// +required
	// +kubebuilder:validation:XValidation:rule="self == oldSelf",message="Field is immutable"
	BillingProjectID *string `json:"billingProjectID,omitempty"`

	// Optional. Description of the firewall endpoint. Max length 2048 characters.
	// +kcc:proto:field=google.cloud.networksecurity.v1beta1.FirewallEndpoint.description
	// +kubebuilder:validation:MaxLength=2048
	Description *string `json:"description,omitempty"`

	// Optional. Labels as key value pairs
	// +kcc:proto:field=google.cloud.networksecurity.v1beta1.FirewallEndpoint.labels
	Labels map[string]string `json:"labels,omitempty"`
}

// NetworkSecurityFirewallEndpointStatus defines the config connector machine state of NetworkSecurityFirewallEndpoint
type NetworkSecurityFirewallEndpointStatus struct {
	/* Conditions represent the latest available observations of the
	   object's current state. */
	Conditions []k8sv1alpha1.Condition `json:"conditions,omitempty"`

	// ObservedGeneration is the generation of the resource that was most recently observed by the Config Connector controller. If this is equal to metadata.generation, then that means that the current reported status reflects the most recent desired state of the resource.
	ObservedGeneration *int64 `json:"observedGeneration,omitempty"`

	// A unique specifier for the NetworkSecurityFirewallEndpoint resource in GCP.
	ExternalRef *string `json:"externalRef,omitempty"`

	// ObservedState is the state of the resource as most recently observed in GCP.
	ObservedState *NetworkSecurityFirewallEndpointObservedState `json:"observedState,omitempty"`
}

// NetworkSecurityFirewallEndpointObservedState defines the observed state of NetworkSecurityFirewallEndpoint
// +kcc:observedstate:proto=google.cloud.networksecurity.v1beta1.FirewallEndpoint
type NetworkSecurityFirewallEndpointObservedState struct {
	// Output only. List of networks that are associated with this endpoint in the local zone. This is a projection of the FirewallEndpointAssociations pointing at this endpoint. A network will only appear in this list after traffic routing is fully configured. Format: projects/{project}/global/networks/{name}.
	// +kcc:proto:field=google.cloud.networksecurity.v1beta1.FirewallEndpoint.associated_networks
	AssociatedNetworks []string `json:"associatedNetworks,omitempty"`

	// Output only. Update time stamp
	// +kcc:proto:field=google.cloud.networksecurity.v1beta1.FirewallEndpoint.update_time
	UpdateTime *string `json:"updateTime,omitempty"`
}

// +genclient
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:resource:categories=gcp,shortName=gcpnetworksecurityfirewallendpoint;gcpnetworksecurityfirewallendpoints
// +kubebuilder:subresource:status
// +kubebuilder:metadata:labels="cnrm.cloud.google.com/managed-by-kcc=true"
// +kubebuilder:metadata:labels="cnrm.cloud.google.com/system=true"
// +kubebuilder:metadata:labels="cnrm.cloud.google.com/direct=true"
// +kubebuilder:printcolumn:name="Age",JSONPath=".metadata.creationTimestamp",type="date"
// +kubebuilder:printcolumn:name="Ready",JSONPath=".status.conditions[?(@.type=='Ready')].status",type="string",description="When 'True', the most recent reconcile of the resource succeeded"
// +kubebuilder:printcolumn:name="Status",JSONPath=".status.conditions[?(@.type=='Ready')].reason",type="string",description="The reason for the value in 'Ready'"
// +kubebuilder:printcolumn:name="Status Age",JSONPath=".status.conditions[?(@.type=='Ready')].lastTransitionTime",type="date",description="The last transition time for the value in 'Status'"

// NetworkSecurityFirewallEndpoint is the Schema for the NetworkSecurityFirewallEndpoint API
// +k8s:openapi-gen=true
// +kubebuilder:storageversion
type NetworkSecurityFirewallEndpoint struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	// +required
	Spec   NetworkSecurityFirewallEndpointSpec   `json:"spec,omitempty"`
	Status NetworkSecurityFirewallEndpointStatus `json:"status,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// NetworkSecurityFirewallEndpointList contains a list of NetworkSecurityFirewallEndpoint
type NetworkSecurityFirewallEndpointList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []NetworkSecurityFirewallEndpoint `json:"items"`
}

func init() {
	SchemeBuilder.Register(&NetworkSecurityFirewallEndpoint{}, &NetworkSecurityFirewallEndpointList{})
}
