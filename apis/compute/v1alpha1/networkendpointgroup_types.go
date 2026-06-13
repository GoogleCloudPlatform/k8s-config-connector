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

var ComputeGlobalNetworkEndpointGroupGVK = GroupVersion.WithKind("ComputeGlobalNetworkEndpointGroup")

type ComputeGlobalNetworkEndpointGroupParent struct {
	// +required
	ProjectRef *refsv1beta1.ProjectRef `json:"projectRef"`
}

// ComputeGlobalNetworkEndpointGroupSpec defines the desired state of ComputeGlobalNetworkEndpointGroup
// +kcc:spec:proto=google.cloud.compute.v1.NetworkEndpointGroup
type ComputeGlobalNetworkEndpointGroupSpec struct {
	ComputeGlobalNetworkEndpointGroupParent `json:",inline"`

	// The ComputeGlobalNetworkEndpointGroup name. If not given, the metadata.name will be used.
	ResourceID *string `json:"resourceID,omitempty"`

	// The default port used if the port number is not specified in the network endpoint. Optional. If the network endpoint type is either GCE_VM_IP, SERVERLESS or PRIVATE_SERVICE_CONNECT, this field must not be specified.
	// +kcc:proto:field=google.cloud.compute.v1.NetworkEndpointGroup.default_port
	// +optional
	DefaultPort *int32 `json:"defaultPort,omitempty"`

	// An optional description of this resource. Provide this property when you create the resource.
	// +kcc:proto:field=google.cloud.compute.v1.NetworkEndpointGroup.description
	// +optional
	Description *string `json:"description,omitempty"`

	// Type of network endpoints in this network endpoint group. Can be one of GCE_VM_IP, GCE_VM_IP_PORT, NON_GCP_PRIVATE_IP_PORT, INTERNET_FQDN_PORT, INTERNET_IP_PORT, SERVERLESS, PRIVATE_SERVICE_CONNECT, GCE_VM_IP_PORTMAP.
	// Check the NetworkEndpointType enum for the list of possible values.
	// +kcc:proto:field=google.cloud.compute.v1.NetworkEndpointGroup.network_endpoint_type
	// +required
	// +kubebuilder:validation:Enum=INTERNET_IP_PORT;INTERNET_FQDN_PORT
	NetworkEndpointType string `json:"networkEndpointType"`
}

// ComputeGlobalNetworkEndpointGroupStatus defines the config connector machine state of ComputeGlobalNetworkEndpointGroup
type ComputeGlobalNetworkEndpointGroupStatus struct {
	/* Conditions represent the latest available observations of the
	   object's current state. */
	Conditions []v1alpha1.Condition `json:"conditions,omitempty"`

	// ObservedGeneration is the generation of the resource that was most recently observed by the Config Connector controller. If this is equal to metadata.generation, then that means that the current reported status reflects the most recent desired state of the resource.
	ObservedGeneration *int64 `json:"observedGeneration,omitempty"`

	// A unique specifier for the ComputeGlobalNetworkEndpointGroup resource in GCP.
	ExternalRef *string `json:"externalRef,omitempty"`

	// ObservedState is the state of the resource as most recently observed in GCP.
	ObservedState *ComputeGlobalNetworkEndpointGroupObservedState `json:"observedState,omitempty"`
}

// ComputeGlobalNetworkEndpointGroupObservedState is the state of the ComputeGlobalNetworkEndpointGroup resource as most recently observed in GCP.
// +kcc:observedstate:proto=google.cloud.compute.v1.NetworkEndpointGroup
type ComputeGlobalNetworkEndpointGroupObservedState struct {
	// [Output Only] Creation timestamp in RFC3339 text format.
	// +kcc:proto:field=google.cloud.compute.v1.NetworkEndpointGroup.creation_timestamp
	CreationTimestamp *string `json:"creationTimestamp,omitempty"`

	// [Output Only] The unique identifier for the resource. This identifier is defined by the server.
	// +kcc:proto:field=google.cloud.compute.v1.NetworkEndpointGroup.id
	ID *uint64 `json:"id,omitempty"`

	// [Output Only] Server-defined URL for the resource.
	// +kcc:proto:field=google.cloud.compute.v1.NetworkEndpointGroup.self_link
	SelfLink *string `json:"selfLink,omitempty"`

	// [Output only] Number of network endpoints in the network endpoint group.
	// +kcc:proto:field=google.cloud.compute.v1.NetworkEndpointGroup.size
	Size *int32 `json:"size,omitempty"`
}

// +genclient
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:resource:categories=gcp,shortName=gcpcomputeglobalnetworkendpointgroup;gcpcomputeglobalnetworkendpointgroups
// +kubebuilder:subresource:status
// +kubebuilder:metadata:labels="cnrm.cloud.google.com/managed-by-kcc=true"
// +kubebuilder:metadata:labels="cnrm.cloud.google.com/system=true"
// +kubebuilder:metadata:labels="cnrm.cloud.google.com/stability-level=alpha"
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
