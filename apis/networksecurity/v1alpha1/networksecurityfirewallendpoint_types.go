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

var NetworkSecurityFirewallEndpointGVK = GroupVersion.WithKind("NetworkSecurityFirewallEndpoint")

type FirewallEndpointEndpointSettings struct {
	// Optional. Immutable. Indicates whether Jumbo Frames are enabled.
	// +kcc:proto:field=google.cloud.networksecurity.v1.FirewallEndpoint.EndpointSettings.jumbo_frames_enabled
	// +kubebuilder:validation:Optional
	JumboFramesEnabled *bool `json:"jumboFramesEnabled,omitempty"`
}

type FirewallEndpointAssociationReference struct {
	// Output only. The resource name of the FirewallEndpointAssociation.
	// +kcc:proto:field=google.cloud.networksecurity.v1.FirewallEndpoint.AssociationReference.name
	Name *string `json:"name,omitempty"`

	// Output only. The VPC network associated.
	// +kcc:proto:field=google.cloud.networksecurity.v1.FirewallEndpoint.AssociationReference.network
	Network *string `json:"network,omitempty"`
}

// NetworkSecurityFirewallEndpointSpec defines the desired state of NetworkSecurityFirewallEndpoint
// +kcc:spec:proto=google.cloud.networksecurity.v1.FirewallEndpoint
type NetworkSecurityFirewallEndpointSpec struct {
	// The project that this resource belongs to.
	// +required
	// +kubebuilder:validation:Required
	ProjectRef *refsv1beta1.ProjectRef `json:"projectRef"`

	// The location of this resource.
	// +required
	// +kubebuilder:validation:Required
	Location string `json:"location"`

	// The NetworkSecurityFirewallEndpoint name. If not given, the metadata.name will be used.
	// +optional
	// +kubebuilder:validation:Optional
	ResourceID *string `json:"resourceID,omitempty"`

	// Optional. Description of the firewall endpoint. Max length 2048 characters.
	// +kcc:proto:field=google.cloud.networksecurity.v1.FirewallEndpoint.description
	// +optional
	// +kubebuilder:validation:Optional
	Description *string `json:"description,omitempty"`

	// Optional. Labels as key value pairs.
	// +kcc:proto:field=google.cloud.networksecurity.v1.FirewallEndpoint.labels
	// +optional
	// +kubebuilder:validation:Optional
	Labels map[string]string `json:"labels,omitempty"`

	// Optional. Billing project ID.
	// +kcc:proto:field=google.cloud.networksecurity.v1.FirewallEndpoint.billing_project_id
	// +optional
	// +kubebuilder:validation:Optional
	BillingProjectID *string `json:"billingProjectID,omitempty"`

	// Optional. Settings for the endpoint.
	// +kcc:proto:field=google.cloud.networksecurity.v1.FirewallEndpoint.endpoint_settings
	// +optional
	// +kubebuilder:validation:Optional
	EndpointSettings *FirewallEndpointEndpointSettings `json:"endpointSettings,omitempty"`
}

// NetworkSecurityFirewallEndpointStatus defines the config connector machine state of NetworkSecurityFirewallEndpoint
type NetworkSecurityFirewallEndpointStatus struct {
	/* Conditions represent the latest available observations of the
	   object's current state. */
	Conditions []v1alpha1.Condition `json:"conditions,omitempty"`

	// ObservedGeneration is the generation of the resource that was most recently observed by the Config Connector controller. If this is equal to metadata.generation, then that means that the current reported status reflects the most recent desired state of the resource.
	ObservedGeneration *int64 `json:"observedGeneration,omitempty"`

	// A unique specifier for the NetworkSecurityFirewallEndpoint resource in GCP.
	ExternalRef *string `json:"externalRef,omitempty"`

	// ObservedState is the state of the resource as most recently observed in GCP.
	ObservedState *NetworkSecurityFirewallEndpointObservedState `json:"observedState,omitempty"`
}

// NetworkSecurityFirewallEndpointObservedState is the state of the NetworkSecurityFirewallEndpoint resource as most recently observed in GCP.
// +kcc:observedstate:proto=google.cloud.networksecurity.v1.FirewallEndpoint
type NetworkSecurityFirewallEndpointObservedState struct {
	// Output only. Create time stamp.
	// +kcc:proto:field=google.cloud.networksecurity.v1.FirewallEndpoint.create_time
	CreateTime *string `json:"createTime,omitempty"`

	// Output only. Update time stamp.
	// +kcc:proto:field=google.cloud.networksecurity.v1.FirewallEndpoint.update_time
	UpdateTime *string `json:"updateTime,omitempty"`

	// Output only. Current state of the endpoint.
	// +kcc:proto:field=google.cloud.networksecurity.v1.FirewallEndpoint.state
	State *string `json:"state,omitempty"`

	// Output only. Whether reconciling is in progress, recommended per https://google.aip.dev/128.
	// +kcc:proto:field=google.cloud.networksecurity.v1.FirewallEndpoint.reconciling
	Reconciling *bool `json:"reconciling,omitempty"`

	// Output only. Deprecated: List of networks that are associated with this endpoint in the local zone.
	// +kcc:proto:field=google.cloud.networksecurity.v1.FirewallEndpoint.associated_networks
	AssociatedNetworks []string `json:"associatedNetworks,omitempty"`

	// Output only. List of FirewallEndpointAssociations that are associated to this endpoint.
	// +kcc:proto:field=google.cloud.networksecurity.v1.FirewallEndpoint.associations
	Associations []FirewallEndpointAssociationReference `json:"associations,omitempty"`
}

// +genclient
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:resource:categories=gcp,shortName=gcpnetworksecurityfirewallendpoint;gcpnetworksecurityfirewallendpoints
// +kubebuilder:subresource:status
// +kubebuilder:metadata:labels="cnrm.cloud.google.com/managed-by-kcc=true"
// +kubebuilder:metadata:labels="cnrm.cloud.google.com/system=true"
// +kubebuilder:metadata:labels="cnrm.cloud.google.com/stability-level=alpha"
// +kubebuilder:printcolumn:name="Age",JSONPath=".metadata.creationTimestamp",type="date"
// +kubebuilder:printcolumn:name="Ready",JSONPath=".status.conditions[?(@.type=='Ready')].status",type="string",description="When 'True', the most recent reconcile of the resource succeeded"
// +kubebuilder:printcolumn:name="Status",JSONPath=".status.conditions[?(@.type=='Ready')].reason",type="string",description="The reason for the value in 'Ready'"
// +kubebuilder:printcolumn:name="Status Age",JSONPath=".status.conditions[?(@.type=='Ready')].lastTransitionTime",type="date",description="The last transition time for the value in 'Status'"

// NetworkSecurityFirewallEndpoint is the Schema for the NetworkSecurityFirewallEndpoint API
// +k8s:openapi-gen=true
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
