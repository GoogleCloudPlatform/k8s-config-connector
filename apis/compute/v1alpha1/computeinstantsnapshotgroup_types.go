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

// API sources for ComputeInstantSnapshotGroup, recorded by generate-types:
// +kcc:source:proto=https://github.com/googleapis/googleapis/blob/1765b559c42386788ff0c6412491277b4791107a/google/cloud/compute/v1/compute.proto
// +kcc:source:service-docs=https://developers.google.com/compute/docs/reference/latest/
// +kcc:source:resource-docs=https://cloud.google.com/compute/docs/reference/rest/v1/instantSnapshotGroups

package v1alpha1

import (
	computev1beta1 "github.com/GoogleCloudPlatform/k8s-config-connector/apis/compute/v1beta1"
	refsv1beta1 "github.com/GoogleCloudPlatform/k8s-config-connector/apis/refs/v1beta1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/apis/k8s/v1alpha1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

var ComputeInstantSnapshotGroupGVK = GroupVersion.WithKind("ComputeInstantSnapshotGroup")

// ComputeInstantSnapshotGroupSpec defines the desired state of ComputeInstantSnapshotGroup
// +kcc:spec:proto=google.cloud.compute.v1.InstantSnapshotGroup
// +kcc:required-from-proto
type ComputeInstantSnapshotGroupSpec struct {
	// The project that this resource belongs to.
	// +required
	ProjectRef *refsv1beta1.ProjectRef `json:"projectRef"`

	// Location represents the geographical location of the ComputeInstantSnapshotGroup. Specify a region name or a zone name. Reference: GCP definition of regions/zones (https://cloud.google.com/compute/docs/regions-zones/)
	// +required
	Location string `json:"location"`

	// The ComputeInstantSnapshotGroup name. If not given, the metadata.name will be used.
	ResourceID *string `json:"resourceID,omitempty"`

	// Optional. An optional description of this resource. Provide this property when you create the resource.
	// +kcc:proto:field=google.cloud.compute.v1.InstantSnapshotGroup.description
	Description *string `json:"description,omitempty"`

	// A reference to the ComputeResourcePolicy that defines the consistency group of disks.
	// +kcc:proto:field=google.cloud.compute.v1.InstantSnapshotGroup.source_consistency_group
	SourceConsistencyGroupRef *computev1beta1.ComputeResourcePolicyRef `json:"sourceConsistencyGroupRef,omitempty"`
}

// ComputeInstantSnapshotGroupStatus defines the config connector machine state of ComputeInstantSnapshotGroup
type ComputeInstantSnapshotGroupStatus struct {
	/* Conditions represent the latest available observations of the
	   object's current state. */
	Conditions []v1alpha1.Condition `json:"conditions,omitempty"`

	// ObservedGeneration is the generation of the resource that was most recently observed by the Config Connector controller. If this is equal to metadata.generation, then that means that the current reported status reflects the most recent desired state of the resource.
	ObservedGeneration *int64 `json:"observedGeneration,omitempty"`

	// A unique specifier for the ComputeInstantSnapshotGroup resource in GCP.
	ExternalRef *string `json:"externalRef,omitempty"`

	// ObservedState is the state of the resource as most recently observed in GCP.
	ObservedState *ComputeInstantSnapshotGroupObservedState `json:"observedState,omitempty"`
}

// ComputeInstantSnapshotGroupObservedState is the state of the ComputeInstantSnapshotGroup resource as most recently observed in GCP.
// +kcc:observedstate:proto=google.cloud.compute.v1.InstantSnapshotGroup
type ComputeInstantSnapshotGroupObservedState struct {
	// Output only. [Output Only] Creation timestamp in RFC3339 text format.
	// +kcc:proto:field=google.cloud.compute.v1.InstantSnapshotGroup.creation_timestamp
	CreationTimestamp *string `json:"creationTimestamp,omitempty"`

	// Output only. [Output Only] The unique identifier for the resource. This identifier is defined by the server.
	// +kcc:proto:field=google.cloud.compute.v1.InstantSnapshotGroup.id
	ID *uint64 `json:"id,omitempty"`

	// Output only. [Output Only] Type of the resource. Always compute#instantSnapshotGroup for InstantSnapshotGroup resources.
	// +kcc:proto:field=google.cloud.compute.v1.InstantSnapshotGroup.kind
	Kind *string `json:"kind,omitempty"`

	// Output only. [Output Only] URL of the region where the instant snapshot group resides.
	// +kcc:proto:field=google.cloud.compute.v1.InstantSnapshotGroup.region
	Region *string `json:"region,omitempty"`

	// Output only. [Output Only] Status information for the instant snapshot group resource.
	// +kcc:proto:field=google.cloud.compute.v1.InstantSnapshotGroup.resource_status
	ResourceStatus *InstantSnapshotGroupResourceStatus `json:"resourceStatus,omitempty"`

	// Output only. [Output Only] Server-defined URL for the resource.
	// +kcc:proto:field=google.cloud.compute.v1.InstantSnapshotGroup.self_link
	SelfLink *string `json:"selfLink,omitempty"`

	// Output only. [Output Only] Server-defined URL for this resource's resource id.
	// +kcc:proto:field=google.cloud.compute.v1.InstantSnapshotGroup.self_link_with_id
	SelfLinkWithID *string `json:"selfLinkWithID,omitempty"`

	// Output only. [Output Only] Check the Status enum for the list of possible values.
	// +kcc:proto:field=google.cloud.compute.v1.InstantSnapshotGroup.status
	Status *string `json:"status,omitempty"`

	// Output only. [Output Only] URL of the zone where the instant snapshot group resides.
	// +kcc:proto:field=google.cloud.compute.v1.InstantSnapshotGroup.zone
	Zone *string `json:"zone,omitempty"`
}

// +genclient
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:resource:categories=gcp,shortName=gcpcomputeinstantsnapshotgroup;gcpcomputeinstantsnapshotgroups
// +kubebuilder:subresource:status
// +kubebuilder:metadata:labels="cnrm.cloud.google.com/managed-by-kcc=true"
// +kubebuilder:metadata:labels="cnrm.cloud.google.com/system=true"
// +kubebuilder:printcolumn:name="Age",JSONPath=".metadata.creationTimestamp",type="date"
// +kubebuilder:printcolumn:name="Ready",JSONPath=".status.conditions[?(@.type=='Ready')].status",type="string",description="When 'True', the most recent reconcile of the resource succeeded"
// +kubebuilder:printcolumn:name="Status",JSONPath=".status.conditions[?(@.type=='Ready')].reason",type="string",description="The reason for the value in 'Ready'"
// +kubebuilder:printcolumn:name="Status Age",JSONPath=".status.conditions[?(@.type=='Ready')].lastTransitionTime",type="date",description="The last transition time for the value in 'Status'"

// ComputeInstantSnapshotGroup is the Schema for the ComputeInstantSnapshotGroup API
// +k8s:openapi-gen=true
type ComputeInstantSnapshotGroup struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	// +required
	Spec   ComputeInstantSnapshotGroupSpec   `json:"spec,omitempty"`
	Status ComputeInstantSnapshotGroupStatus `json:"status,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// ComputeInstantSnapshotGroupList contains a list of ComputeInstantSnapshotGroup
type ComputeInstantSnapshotGroupList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []ComputeInstantSnapshotGroup `json:"items"`
}

func init() {
	SchemeBuilder.Register(&ComputeInstantSnapshotGroup{}, &ComputeInstantSnapshotGroupList{})
}
