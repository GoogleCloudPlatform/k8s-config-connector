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

// API sources for ComputeInterconnectGroup, recorded by generate-types:
// +kcc:source:proto=https://github.com/googleapis/googleapis/blob/1765b559c42386788ff0c6412491277b4791107a/google/cloud/compute/v1/compute.proto
// +kcc:source:service-docs=https://developers.google.com/compute/docs/reference/latest/
// +kcc:guess=source-link reason=verify-resource-docs-link
// +kcc:source:resource-docs=https://docs.cloud.google.com/compute/docs/reference/rest

package v1alpha1

import (
	refsv1beta1 "github.com/GoogleCloudPlatform/k8s-config-connector/apis/refs/v1beta1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/apis/k8s/v1alpha1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

var ComputeInterconnectGroupGVK = GroupVersion.WithKind("ComputeInterconnectGroup")

// ComputeInterconnectGroupSpec defines the desired state of ComputeInterconnectGroup
// +kcc:spec:proto=google.cloud.compute.v1.InterconnectGroup
// +kcc:required-from-proto
type ComputeInterconnectGroupSpec struct {
	// The project that this resource belongs to.
	ProjectRef *refsv1beta1.ProjectRef `json:"projectRef"`

	// The location of this resource.
	Location *string `json:"location"`

	// The ComputeInterconnectGroup name. If not given, the metadata.name will be used.
	ResourceID *string `json:"resourceID,omitempty"`
	// +kcc:proto:field=google.cloud.compute.v1.InterconnectGroup.configured
	Configured *InterconnectGroupConfigured `json:"configured,omitempty"`

	// An optional description of this resource. Provide this property when you create the resource.
	// +kcc:proto:field=google.cloud.compute.v1.InterconnectGroup.description
	Description *string `json:"description,omitempty"`

	// +kcc:proto:field=google.cloud.compute.v1.InterconnectGroup.intent
	Intent *InterconnectGroupIntent `json:"intent,omitempty"`

	// Interconnects in the InterconnectGroup. Keys are arbitrary user-specified strings. Users are encouraged, but not required, to use their preferred format for resource links as keys. Note that there are add-members and remove-members methods in gcloud. The size of this map is limited by an "Interconnects per group" quota.
	// +kcc:proto:field=google.cloud.compute.v1.InterconnectGroup.interconnects
	Interconnects map[string]InterconnectGroupInterconnect `json:"interconnects,omitempty"`

	// +kcc:proto:field=google.cloud.compute.v1.InterconnectGroup.physical_structure
	PhysicalStructure *InterconnectGroupPhysicalStructure `json:"physicalStructure,omitempty"`
}

// ComputeInterconnectGroupStatus defines the config connector machine state of ComputeInterconnectGroup
type ComputeInterconnectGroupStatus struct {
	/* Conditions represent the latest available observations of the
	   object's current state. */
	Conditions []v1alpha1.Condition `json:"conditions,omitempty"`

	// ObservedGeneration is the generation of the resource that was most recently observed by the Config Connector controller. If this is equal to metadata.generation, then that means that the current reported status reflects the most recent desired state of the resource.
	ObservedGeneration *int64 `json:"observedGeneration,omitempty"`

	// A unique specifier for the ComputeInterconnectGroup resource in GCP.
	ExternalRef *string `json:"externalRef,omitempty"`

	// ObservedState is the state of the resource as most recently observed in GCP.
	ObservedState *ComputeInterconnectGroupObservedState `json:"observedState,omitempty"`
}

// ComputeInterconnectGroupObservedState is the state of the ComputeInterconnectGroup resource as most recently observed in GCP.
// +kcc:observedstate:proto=google.cloud.compute.v1.InterconnectGroup
type ComputeInterconnectGroupObservedState struct {
	// [Output Only] Creation timestamp in RFC3339 text format.
	// +kcc:guess=placement reason=no-field-behavior-on-message
	// +kcc:proto:field=google.cloud.compute.v1.InterconnectGroup.creation_timestamp
	CreationTimestamp *string `json:"creationTimestamp,omitempty"`

	// Opaque system-generated token that uniquely identifies the configuration. If provided when patching a configuration in update mode, the provided token must match the current token or the update is rejected. This provides a reliable means of doing read-modify-write (optimistic locking) as described by API 154.
	// +kcc:guess=placement reason=no-field-behavior-on-message
	// +kcc:proto:field=google.cloud.compute.v1.InterconnectGroup.etag
	Etag *string `json:"etag,omitempty"`

	// [Output Only] The unique identifier for the resource type. The server generates this identifier.
	// +kcc:guess=placement reason=no-field-behavior-on-message
	// +kcc:proto:field=google.cloud.compute.v1.InterconnectGroup.id
	ID *uint64 `json:"id,omitempty"`

	// [Output Only] Type of the resource. Always compute#InterconnectGroup
	// +kcc:guess=placement reason=no-field-behavior-on-message
	// +kcc:proto:field=google.cloud.compute.v1.InterconnectGroup.kind
	Kind *string `json:"kind,omitempty"`

	// [Output Only] Server-defined URL for the resource.
	// +kcc:guess=placement reason=no-field-behavior-on-message
	// +kcc:proto:field=google.cloud.compute.v1.InterconnectGroup.self_link
	SelfLink *string `json:"selfLink,omitempty"`
}

// +genclient
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:resource:categories=gcp,shortName=gcpcomputeinterconnectgroup;gcpcomputeinterconnectgroups
// +kubebuilder:subresource:status
// +kubebuilder:metadata:labels="cnrm.cloud.google.com/managed-by-kcc=true"
// +kubebuilder:metadata:labels="cnrm.cloud.google.com/system=true"
// +kubebuilder:printcolumn:name="Age",JSONPath=".metadata.creationTimestamp",type="date"
// +kubebuilder:printcolumn:name="Ready",JSONPath=".status.conditions[?(@.type=='Ready')].status",type="string",description="When 'True', the most recent reconcile of the resource succeeded"
// +kubebuilder:printcolumn:name="Status",JSONPath=".status.conditions[?(@.type=='Ready')].reason",type="string",description="The reason for the value in 'Ready'"
// +kubebuilder:printcolumn:name="Status Age",JSONPath=".status.conditions[?(@.type=='Ready')].lastTransitionTime",type="date",description="The last transition time for the value in 'Status'"

// ComputeInterconnectGroup is the Schema for the ComputeInterconnectGroup API
// +k8s:openapi-gen=true
type ComputeInterconnectGroup struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	// +required
	Spec   ComputeInterconnectGroupSpec   `json:"spec,omitempty"`
	Status ComputeInterconnectGroupStatus `json:"status,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// ComputeInterconnectGroupList contains a list of ComputeInterconnectGroup
type ComputeInterconnectGroupList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []ComputeInterconnectGroup `json:"items"`
}

func init() {
	SchemeBuilder.Register(&ComputeInterconnectGroup{}, &ComputeInterconnectGroupList{})
}
