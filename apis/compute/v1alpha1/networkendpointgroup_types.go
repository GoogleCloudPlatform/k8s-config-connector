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
	"github.com/GoogleCloudPlatform/k8s-config-connector/apis/refs"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/apis/k8s/v1alpha1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

var ComputeGlobalNetworkEndpointGroupGVK = GroupVersion.WithKind("ComputeGlobalNetworkEndpointGroup")

// ComputeGlobalNetworkEndpointGroupSpec defines the desired state of ComputeGlobalNetworkEndpointGroup
// +kcc:spec:proto=google.cloud.compute.v1.NetworkEndpointGroup
type ComputeGlobalNetworkEndpointGroupSpec struct {
	/* Immutable. The default port used if the port number is not specified in the
	network endpoint. */
	// +optional
	// +kcc:proto:field=google.cloud.compute.v1.NetworkEndpointGroup.default_port
	DefaultPort *int64 `json:"defaultPort,omitempty"`

	/* Immutable. An optional description of this resource. Provide this property when
	you create the resource. */
	// +optional
	// +kcc:proto:field=google.cloud.compute.v1.NetworkEndpointGroup.description
	Description *string `json:"description,omitempty"`

	/* Immutable. Type of network endpoints in this network endpoint group. Possible values: ["INTERNET_IP_PORT", "INTERNET_FQDN_PORT"]. */
	// +required
	// +kcc:proto:field=google.cloud.compute.v1.NetworkEndpointGroup.network_endpoint_type
	NetworkEndpointType string `json:"networkEndpointType"`

	/* The project that this resource belongs to. */
	// +required
	ProjectRef *refs.ProjectRef `json:"projectRef"`

	/* Immutable. Optional. The name of the resource. Used for creation and acquisition. When unset, the value of `metadata.name` is used as the default. */
	// +optional
	ResourceID *string `json:"resourceID,omitempty"`
}

// ComputeGlobalNetworkEndpointGroupStatus defines the config connector machine state of ComputeGlobalNetworkEndpointGroup
type ComputeGlobalNetworkEndpointGroupStatus struct {
	/* Conditions represent the latest available observations of the
	   ComputeGlobalNetworkEndpointGroup's current state. */
	Conditions []v1alpha1.Condition `json:"conditions,omitempty"`

	/* ObservedGeneration is the generation of the resource that was most recently observed by the Config Connector controller. If this is equal to metadata.generation, then that means that the current reported status reflects the most recent desired state of the resource. */
	// +optional
	ObservedGeneration *int64 `json:"observedGeneration,omitempty"`

	// +optional
	// +kcc:proto:field=google.cloud.compute.v1.NetworkEndpointGroup.self_link
	SelfLink *string `json:"selfLink,omitempty"`
}

// +genclient
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:resource:categories=gcp,shortName=gcpcomputeglobalnetworkendpointgroup;gcpcomputeglobalnetworkendpointgroups
// +kubebuilder:subresource:status
// +kubebuilder:metadata:labels="cnrm.cloud.google.com/managed-by-kcc=true"
// +kubebuilder:metadata:labels="cnrm.cloud.google.com/stability-level=alpha"
// +kubebuilder:metadata:labels="cnrm.cloud.google.com/system=true"
// +kubebuilder:metadata:labels="cnrm.cloud.google.com/tf2crd=true"
// +kubebuilder:printcolumn:name="Age",JSONPath=".metadata.creationTimestamp",type="date"
// +kubebuilder:printcolumn:name="Ready",JSONPath=".status.conditions[?(@.type=='Ready')].status",type="string",description="When 'True', the most recent reconcile of the resource succeeded"
// +kubebuilder:printcolumn:name="Status",JSONPath=".status.conditions[?(@.type=='Ready')].reason",type="string",description="The reason for the value in 'Ready'"
// +kubebuilder:printcolumn:name="Status Age",JSONPath=".status.conditions[?(@.type=='Ready')].lastTransitionTime",type="date",description="The last transition time for the value in 'Status'"

// ComputeGlobalNetworkEndpointGroup is the Schema for the ComputeGlobalNetworkEndpointGroup API
// +k8s:openapi-gen=true
type ComputeGlobalNetworkEndpointGroup struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	// +required
	Spec   ComputeGlobalNetworkEndpointGroupSpec   `json:"spec,omitempty"`
	Status ComputeGlobalNetworkEndpointGroupStatus `json:"status,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// ComputeGlobalNetworkEndpointGroupList contains a list of ComputeGlobalNetworkEndpointGroup
type ComputeGlobalNetworkEndpointGroupList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []ComputeGlobalNetworkEndpointGroup `json:"items"`
}

func init() {
	SchemeBuilder.Register(&ComputeGlobalNetworkEndpointGroup{}, &ComputeGlobalNetworkEndpointGroupList{})
}
