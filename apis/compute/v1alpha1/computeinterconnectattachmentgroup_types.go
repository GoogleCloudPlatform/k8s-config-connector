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

// API sources for ComputeInterconnectAttachmentGroup, recorded by generate-types:
// +kcc:source:proto=https://github.com/googleapis/googleapis/blob/1765b559c42386788ff0c6412491277b4791107a/google/cloud/compute/v1/compute.proto
// +kcc:source:service-docs=https://developers.google.com/compute/docs/reference/latest/
// +kcc:source:resource-docs=https://cloud.google.com/compute/docs/reference/rest/v1/interconnectAttachmentGroups

package v1alpha1

import (
	computev1beta1 "github.com/GoogleCloudPlatform/k8s-config-connector/apis/compute/v1beta1"
	refsv1beta1 "github.com/GoogleCloudPlatform/k8s-config-connector/apis/refs/v1beta1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/apis/k8s/v1alpha1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

var ComputeInterconnectAttachmentGroupGVK = GroupVersion.WithKind("ComputeInterconnectAttachmentGroup")

// ComputeInterconnectAttachmentGroupSpec defines the desired state of ComputeInterconnectAttachmentGroup
// +kcc:spec:proto=google.cloud.compute.v1.InterconnectAttachmentGroup
// +kcc:required-from-proto
type ComputeInterconnectAttachmentGroupSpec struct {
	// The project that this resource belongs to.
	// +required
	ProjectRef *refsv1beta1.ProjectRef `json:"projectRef"`

	// The ComputeInterconnectAttachmentGroup name. If not given, the metadata.name will be used.
	ResourceID *string `json:"resourceID,omitempty"`

	// Attachments in the AttachmentGroup. Keys are arbitrary user-specified strings. Users are encouraged, but not required, to use their preferred format for resource links as keys. Note that there are add-members and remove-members methods in gcloud. The size of this map is limited by an "Attachments per group" quota.
	// +kcc:proto:field=google.cloud.compute.v1.InterconnectAttachmentGroup.attachments
	Attachments map[string]InterconnectAttachmentGroupAttachment `json:"attachments,omitempty"`

	// An optional description of this resource. Provide this property when you create the resource.
	// +kcc:proto:field=google.cloud.compute.v1.InterconnectAttachmentGroup.description
	Description *string `json:"description,omitempty"`

	// +kcc:proto:field=google.cloud.compute.v1.InterconnectAttachmentGroup.intent
	Intent *InterconnectAttachmentGroupIntent `json:"intent,omitempty"`

	// The URL of an InterconnectGroup that groups these Attachments' Interconnects. Customers do not need to set this unless directed by Google Support.
	// +kcc:proto:field=google.cloud.compute.v1.InterconnectAttachmentGroup.interconnect_group
	InterconnectGroupRef *ComputeInterconnectGroupRef `json:"interconnectGroupRef,omitempty"`
}

// +kcc:proto=google.cloud.compute.v1.InterconnectAttachmentGroupAttachment
type InterconnectAttachmentGroupAttachment struct {
	// The URL of an InterconnectAttachment in this group.
	// +kcc:proto:field=google.cloud.compute.v1.InterconnectAttachmentGroupAttachment.attachment
	AttachmentRef *computev1beta1.ComputeInterconnectAttachmentRef `json:"attachmentRef,omitempty"`
}

// ComputeInterconnectAttachmentGroupStatus defines the config connector machine state of ComputeInterconnectAttachmentGroup
type ComputeInterconnectAttachmentGroupStatus struct {
	/* Conditions represent the latest available observations of the
	   object's current state. */
	Conditions []v1alpha1.Condition `json:"conditions,omitempty"`

	// ObservedGeneration is the generation of the resource that was most recently observed by the Config Connector controller. If this is equal to metadata.generation, then that means that the current reported status reflects the most recent desired state of the resource.
	ObservedGeneration *int64 `json:"observedGeneration,omitempty"`

	// A unique specifier for the ComputeInterconnectAttachmentGroup resource in GCP.
	ExternalRef *string `json:"externalRef,omitempty"`

	// ObservedState is the state of the resource as most recently observed in GCP.
	ObservedState *ComputeInterconnectAttachmentGroupObservedState `json:"observedState,omitempty"`
}

// ComputeInterconnectAttachmentGroupObservedState is the state of the ComputeInterconnectAttachmentGroup resource as most recently observed in GCP.
// +kcc:observedstate:proto=google.cloud.compute.v1.InterconnectAttachmentGroup
type ComputeInterconnectAttachmentGroupObservedState struct {
	// [Output Only] The redundancy this group is configured to support.
	// +kcc:proto:field=google.cloud.compute.v1.InterconnectAttachmentGroup.configured
	Configured *InterconnectAttachmentGroupConfigured `json:"configured,omitempty"`

	// [Output Only] Creation timestamp in RFC3339 text format.
	// +kcc:proto:field=google.cloud.compute.v1.InterconnectAttachmentGroup.creation_timestamp
	CreationTimestamp *string `json:"creationTimestamp,omitempty"`

	// Opaque system-generated token that uniquely identifies the configuration. If provided when patching a configuration in update mode, the provided token must match the current token or the update is rejected. This provides a reliable means of doing read-modify-write (optimistic locking) as described by AIP 154.
	// +kcc:proto:field=google.cloud.compute.v1.InterconnectAttachmentGroup.etag
	Etag *string `json:"etag,omitempty"`

	// [Output Only] The unique identifier for the resource type. The server generates this identifier.
	// +kcc:proto:field=google.cloud.compute.v1.InterconnectAttachmentGroup.id
	ID *uint64 `json:"id,omitempty"`

	// [Output Only] Type of the resource. Always compute#interconnectAttachmentGroup.
	// +kcc:proto:field=google.cloud.compute.v1.InterconnectAttachmentGroup.kind
	Kind *string `json:"kind,omitempty"`

	// [Output Only] An analysis of the logical layout of Attachments in this group.
	// +kcc:proto:field=google.cloud.compute.v1.InterconnectAttachmentGroup.logical_structure
	LogicalStructure *InterconnectAttachmentGroupLogicalStructure `json:"logicalStructure,omitempty"`

	// [Output Only] Server-defined URL for the resource.
	// +kcc:proto:field=google.cloud.compute.v1.InterconnectAttachmentGroup.self_link
	SelfLink *string `json:"selfLink,omitempty"`
}

// +genclient
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:resource:categories=gcp,shortName=gcpcomputeinterconnectattachmentgroup;gcpcomputeinterconnectattachmentgroups
// +kubebuilder:subresource:status
// +kubebuilder:metadata:labels="cnrm.cloud.google.com/managed-by-kcc=true"
// +kubebuilder:metadata:labels="cnrm.cloud.google.com/system=true"
// +kubebuilder:printcolumn:name="Age",JSONPath=".metadata.creationTimestamp",type="date"
// +kubebuilder:printcolumn:name="Ready",JSONPath=".status.conditions[?(@.type=='Ready')].status",type="string",description="When 'True', the most recent reconcile of the resource succeeded"
// +kubebuilder:printcolumn:name="Status",JSONPath=".status.conditions[?(@.type=='Ready')].reason",type="string",description="The reason for the value in 'Ready'"
// +kubebuilder:printcolumn:name="Status Age",JSONPath=".status.conditions[?(@.type=='Ready')].lastTransitionTime",type="date",description="The last transition time for the value in 'Status'"

// ComputeInterconnectAttachmentGroup is the Schema for the ComputeInterconnectAttachmentGroup API
// +k8s:openapi-gen=true
type ComputeInterconnectAttachmentGroup struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	// +required
	Spec   ComputeInterconnectAttachmentGroupSpec   `json:"spec,omitempty"`
	Status ComputeInterconnectAttachmentGroupStatus `json:"status,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// ComputeInterconnectAttachmentGroupList contains a list of ComputeInterconnectAttachmentGroup
type ComputeInterconnectAttachmentGroupList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []ComputeInterconnectAttachmentGroup `json:"items"`
}

func init() {
	SchemeBuilder.Register(&ComputeInterconnectAttachmentGroup{}, &ComputeInterconnectAttachmentGroupList{})
}
