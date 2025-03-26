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
	refv1beta1 "github.com/GoogleCloudPlatform/k8s-config-connector/apis/refs/v1beta1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/apis/k8s/v1alpha1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

var ComputeNetworkAttachmentGVK = GroupVersion.WithKind("ComputeNetworkAttachment")

// ComputeNetworkAttachmentParent holds the fields describing the parent of the NetworkAttachment resource.
type ComputeNetworkAttachmentParent struct {
	// +required
	ProjectRef *refv1beta1.ProjectRef `json:"projectRef"`
	// +required
	Location string `json:"location"`
}

// ComputeNetworkAttachmentSpec defines the desired state of ComputeNetworkAttachment
// +kcc:proto=google.cloud.compute.v1.NetworkAttachment
type ComputeNetworkAttachmentSpec struct {
	// Parent reference.
	ComputeNetworkAttachmentParent `json:",inline"`

	// The ComputeNetworkAttachment name. If not given, the metadata.name will be used.
	ResourceID *string `json:"resourceID,omitempty"`

	// Check the ConnectionPreference enum for the list of possible values.
	// +kcc:proto:field=google.cloud.compute.v1.NetworkAttachment.connection_preference
	ConnectionPreference *string `json:"connectionPreference,omitempty"`

	// An optional description of this resource. Provide this property when you create the resource.
	// +kcc:proto:field=google.cloud.compute.v1.NetworkAttachment.description
	Description *string `json:"description,omitempty"`

	// Fingerprint of this resource. A hash of the contents stored in this object. This field is used in optimistic locking. An up-to-date fingerprint must be provided in order to patch.
	// +kcc:proto:field=google.cloud.compute.v1.NetworkAttachment.fingerprint
	Fingerprint *string `json:"fingerprint,omitempty"`

	// Projects that are allowed to connect to this network attachment. The project can be specified using its id or number.
	// +kcc:proto:field=google.cloud.compute.v1.NetworkAttachment.producer_accept_lists
	ProducerAcceptLists []*refv1beta1.ProjectRef `json:"producerAcceptLists,omitempty"`

	// Projects that are not allowed to connect to this network attachment. The project can be specified using its id or number.
	// +kcc:proto:field=google.cloud.compute.v1.NetworkAttachment.producer_reject_lists
	ProducerRejectLists []*refv1beta1.ProjectRef `json:"producerRejectLists,omitempty"`

	// An array of URLs where each entry is the URL of a subnet provided by the service consumer to use for endpoints in the producers that connect to this network attachment.
	// +kcc:proto:field=google.cloud.compute.v1.NetworkAttachment.subnetworks
	Subnetworks []*refv1beta1.ComputeSubnetworkRef `json:"subnetworks,omitempty"`
}

// ComputeNetworkAttachmentStatus defines the config connector machine state of ComputeNetworkAttachment
type ComputeNetworkAttachmentStatus struct {
	/* Conditions represent the latest available observations of the
	   object's current state. */
	Conditions []v1alpha1.Condition `json:"conditions,omitempty"`

	// ObservedGeneration is the generation of the resource that was most recently observed by the Config Connector controller. If this is equal to metadata.generation, then that means that the current reported status reflects the most recent desired state of the resource.
	ObservedGeneration *int64 `json:"observedGeneration,omitempty"`

	// A unique specifier for the ComputeNetworkAttachment resource in GCP.
	ExternalRef *string `json:"externalRef,omitempty"`

	// ObservedState is the state of the resource as most recently observed in GCP.
	ObservedState *ComputeNetworkAttachmentObservedState `json:"observedState,omitempty"`
}

// ComputeNetworkAttachmentObservedState is the state of the ComputeNetworkAttachment resource as most recently observed in GCP.
// +kcc:proto=google.cloud.compute.v1.NetworkAttachment
type ComputeNetworkAttachmentObservedState struct {
	// [Output Only] An array of connections for all the producers connected to this network attachment.
	// +kcc:proto:field=google.cloud.compute.v1.NetworkAttachment.connection_endpoints
	ConnectionEndpoints []NetworkAttachmentConnectedEndpoint `json:"connectionEndpoints,omitempty"`

	// [Output Only] Creation timestamp in RFC3339 text format.
	// +kcc:proto:field=google.cloud.compute.v1.NetworkAttachment.creation_timestamp
	CreationTimestamp *string `json:"creationTimestamp,omitempty"`

	// [Output Only] The unique identifier for the resource type. The server generates this identifier.
	// +kcc:proto:field=google.cloud.compute.v1.NetworkAttachment.id
	ID *uint64 `json:"id,omitempty"`

	// [Output Only] Type of the resource.
	// +kcc:proto:field=google.cloud.compute.v1.NetworkAttachment.kind
	Kind *string `json:"kind,omitempty"`

	// [Output Only] URL of the region where the network attachment resides. This field applies only to the region resource. You must specify this field as part of the HTTP request URL. It is not settable as a field in the request body.
	// +kcc:proto:field=google.cloud.compute.v1.NetworkAttachment.region
	Region *string `json:"region,omitempty"`

	// [Output Only] Server-defined URL for the resource.
	// +kcc:proto:field=google.cloud.compute.v1.NetworkAttachment.self_link
	SelfLink *string `json:"selfLink,omitempty"`

	// [Output Only] Server-defined URL for this resource's resource id.
	// +kcc:proto:field=google.cloud.compute.v1.NetworkAttachment.self_link_with_id
	SelfLinkWithID *string `json:"selfLinkWithID,omitempty"`

	// [Output Only] The URL of the network which the Network Attachment belongs to. Practically it is inferred by fetching the network of the first subnetwork associated. Because it is required that all the subnetworks must be from the same network, it is assured that the Network Attachment belongs to the same network as all the subnetworks.
	// +kcc:proto:field=google.cloud.compute.v1.NetworkAttachment.network
	Network *string `json:"network,omitempty"`
}

// +genclient
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:resource:categories=gcp,shortName=gcpcomputenetworkattachment;gcpcomputenetworkattachments
// +kubebuilder:subresource:status
// +kubebuilder:metadata:labels="cnrm.cloud.google.com/managed-by-kcc=true";"cnrm.cloud.google.com/system=true"
// +kubebuilder:printcolumn:name="Age",JSONPath=".metadata.creationTimestamp",type="date"
// +kubebuilder:printcolumn:name="Ready",JSONPath=".status.conditions[?(@.type=='Ready')].status",type="string",description="When 'True', the most recent reconcile of the resource succeeded"
// +kubebuilder:printcolumn:name="Status",JSONPath=".status.conditions[?(@.type=='Ready')].reason",type="string",description="The reason for the value in 'Ready'"
// +kubebuilder:printcolumn:name="Status Age",JSONPath=".status.conditions[?(@.type=='Ready')].lastTransitionTime",type="date",description="The last transition time for the value in 'Status'"

// ComputeNetworkAttachment is the Schema for the ComputeNetworkAttachment API
// +k8s:openapi-gen=true
type ComputeNetworkAttachment struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	// +required
	Spec   ComputeNetworkAttachmentSpec   `json:"spec,omitempty"`
	Status ComputeNetworkAttachmentStatus `json:"status,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// ComputeNetworkAttachmentList contains a list of ComputeNetworkAttachment
type ComputeNetworkAttachmentList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []ComputeNetworkAttachment `json:"items"`
}

func init() {
	SchemeBuilder.Register(&ComputeNetworkAttachment{}, &ComputeNetworkAttachmentList{})
}
