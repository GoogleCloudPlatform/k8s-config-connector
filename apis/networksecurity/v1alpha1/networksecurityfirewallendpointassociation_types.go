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

var NetworkSecurityFirewallEndpointAssociationGVK = GroupVersion.WithKind("NetworkSecurityFirewallEndpointAssociation")

// NetworkSecurityFirewallEndpointAssociationSpec defines the desired state of NetworkSecurityFirewallEndpointAssociation
// +kcc:spec:proto=google.cloud.networksecurity.v1.FirewallEndpointAssociation
type NetworkSecurityFirewallEndpointAssociationSpec struct {
	// The project that this resource belongs to.
	// +required
	ProjectRef *refsv1beta1.ProjectRef `json:"projectRef"`

	// The location of this resource.
	// +required
	Location *string `json:"location"`

	// The NetworkSecurityFirewallEndpointAssociation name. If not given, the metadata.name will be used.
	// +optional
	ResourceID *string `json:"resourceID,omitempty"`

	// Optional. Labels as key value pairs
	// +optional
	Labels map[string]string `json:"labels,omitempty"`

	// Required. The URL of the network that is being associated.
	// +required
	NetworkRef *NetworkRef `json:"networkRef"`

	// Required. The URL of the FirewallEndpoint that is being associated.
	// +required
	FirewallEndpointRef *FirewallEndpointRef `json:"firewallEndpointRef"`

	// Optional. The URL of the TlsInspectionPolicy that is being associated.
	// +optional
	TLSInspectionPolicyRef *TLSInspectionPolicyRef `json:"tlsInspectionPolicyRef,omitempty"`

	// Optional. Whether the association is disabled.
	// True indicates that traffic won't be intercepted
	// +optional
	Disabled *bool `json:"disabled,omitempty"`
}

type NetworkRef struct {
	/* The network selflink of form "projects/{{project}}/global/networks/{{name}}", when not managed by Config Connector. */
	External string `json:"external,omitempty"`
	/* The `name` field of a `ComputeNetwork` resource. */
	Name string `json:"name,omitempty"`
	/* The `namespace` field of a `ComputeNetwork` resource. */
	Namespace string `json:"namespace,omitempty"`
}

type FirewallEndpointRef struct {
	/* The firewall endpoint selflink, when not managed by Config Connector. */
	External string `json:"external,omitempty"`
	/* The `name` field of a `NetworkSecurityFirewallEndpoint` resource. */
	Name string `json:"name,omitempty"`
	/* The `namespace` field of a `NetworkSecurityFirewallEndpoint` resource. */
	Namespace string `json:"namespace,omitempty"`
}

type TLSInspectionPolicyRef struct {
	/* The TLS inspection policy selflink, when not managed by Config Connector. */
	External string `json:"external,omitempty"`
	/* The `name` field of a `NetworkSecurityTLSInspectionPolicy` resource. */
	Name string `json:"name,omitempty"`
	/* The `namespace` field of a `NetworkSecurityTLSInspectionPolicy` resource. */
	Namespace string `json:"namespace,omitempty"`
}

// NetworkSecurityFirewallEndpointAssociationStatus defines the config connector machine state of NetworkSecurityFirewallEndpointAssociation
type NetworkSecurityFirewallEndpointAssociationStatus struct {
	/* Conditions represent the latest available observations of the
	   object's current state. */
	Conditions []v1alpha1.Condition `json:"conditions,omitempty"`

	// ObservedGeneration is the generation of the resource that was most recently observed by the Config Connector controller. If this is equal to metadata.generation, then that means that the current reported status reflects the most recent desired state of the resource.
	ObservedGeneration *int64 `json:"observedGeneration,omitempty"`

	// A unique specifier for the NetworkSecurityFirewallEndpointAssociation resource in GCP.
	ExternalRef *string `json:"externalRef,omitempty"`

	// ObservedState is the state of the resource as most recently observed in GCP.
	ObservedState *NetworkSecurityFirewallEndpointAssociationObservedState `json:"observedState,omitempty"`
}

// NetworkSecurityFirewallEndpointAssociationObservedState is the state of the NetworkSecurityFirewallEndpointAssociation resource as most recently observed in GCP.
// +kcc:observedstate:proto=google.cloud.networksecurity.v1.FirewallEndpointAssociation
type NetworkSecurityFirewallEndpointAssociationObservedState struct {
	// Output only. Create time stamp
	CreateTime *string `json:"createTime,omitempty"`

	// Output only. Update time stamp
	UpdateTime *string `json:"updateTime,omitempty"`

	// Output only. Current state of the association.
	State *string `json:"state,omitempty"`

	// Output only. Whether reconciling is in progress, recommended per
	// https://google.aip.dev/128.
	Reconciling *bool `json:"reconciling,omitempty"`
}

// +genclient
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:resource:categories=gcp,shortName=gcpnetworksecurityfirewallendpointassociation;gcpnetworksecurityfirewallendpointassociations
// +kubebuilder:subresource:status
// +kubebuilder:metadata:labels="cnrm.cloud.google.com/managed-by-kcc=true"
// +kubebuilder:metadata:labels="cnrm.cloud.google.com/system=true"
// +kubebuilder:printcolumn:name="Age",JSONPath=".metadata.creationTimestamp",type="date"
// +kubebuilder:printcolumn:name="Ready",JSONPath=".status.conditions[?(@.type=='Ready')].status",type="string",description="When 'True', the most recent reconcile of the resource succeeded"
// +kubebuilder:printcolumn:name="Status",JSONPath=".status.conditions[?(@.type=='Ready')].reason",type="string",description="The reason for the value in 'Ready'"
// +kubebuilder:printcolumn:name="Status Age",JSONPath=".status.conditions[?(@.type=='Ready')].lastTransitionTime",type="date",description="The last transition time for the value in 'Status'"

// NetworkSecurityFirewallEndpointAssociation is the Schema for the NetworkSecurityFirewallEndpointAssociation API
// +k8s:openapi-gen=true
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
