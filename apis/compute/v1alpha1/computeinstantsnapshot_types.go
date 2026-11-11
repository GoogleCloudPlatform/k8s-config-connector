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

// API sources for ComputeInstantSnapshot, recorded by generate-types:
// +kcc:source:proto=https://github.com/googleapis/googleapis/blob/1765b559c42386788ff0c6412491277b4791107a/google/cloud/compute/v1/compute.proto
// +kcc:source:service-docs=https://developers.google.com/compute/docs/reference/latest/
// +kcc:source:resource-docs=https://cloud.google.com/compute/docs/reference/rest/v1/instantSnapshots

package v1alpha1

import (
	computev1beta1 "github.com/GoogleCloudPlatform/k8s-config-connector/apis/compute/v1beta1"
	refsv1beta1 "github.com/GoogleCloudPlatform/k8s-config-connector/apis/refs/v1beta1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/apis/k8s/v1alpha1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

var ComputeInstantSnapshotGVK = GroupVersion.WithKind("ComputeInstantSnapshot")

// ComputeInstantSnapshotSpec defines the desired state of ComputeInstantSnapshot
// +kcc:spec:proto=google.cloud.compute.v1.InstantSnapshot
// +kcc:required-from-proto
type ComputeInstantSnapshotSpec struct {
	// The project that this resource belongs to.
	// +required
	ProjectRef *refsv1beta1.ProjectRef `json:"projectRef"`

	// Location represents the geographical location of the ComputeInstantSnapshot. Specify a region name or a zone name. Reference: GCP definition of regions/zones (https://cloud.google.com/compute/docs/regions-zones/)
	// +required
	Location string `json:"location"`

	// The ComputeInstantSnapshot name. If not given, the metadata.name will be used.
	ResourceID *string `json:"resourceID,omitempty"`

	// An optional description of this resource. Provide this property when you create the resource.
	// +kcc:proto:field=google.cloud.compute.v1.InstantSnapshot.description
	Description *string `json:"description,omitempty"`

	// A reference to the disk used to create this instant snapshot. Note that the source disk must be in the same zone/region as the instant snapshot to be created.
	// +kcc:proto:field=google.cloud.compute.v1.InstantSnapshot.source_disk
	SourceDiskRef *computev1beta1.ComputeDiskRef `json:"sourceDiskRef,omitempty"`
}

// ComputeInstantSnapshotStatus defines the config connector machine state of ComputeInstantSnapshot
type ComputeInstantSnapshotStatus struct {
	/* Conditions represent the latest available observations of the
	   object's current state. */
	Conditions []v1alpha1.Condition `json:"conditions,omitempty"`

	// ObservedGeneration is the generation of the resource that was most recently observed by the Config Connector controller. If this is equal to metadata.generation, then that means that the current reported status reflects the most recent desired state of the resource.
	ObservedGeneration *int64 `json:"observedGeneration,omitempty"`

	// A unique specifier for the ComputeInstantSnapshot resource in GCP.
	ExternalRef *string `json:"externalRef,omitempty"`

	// ObservedState is the state of the resource as most recently observed in GCP.
	ObservedState *ComputeInstantSnapshotObservedState `json:"observedState,omitempty"`
}

// ComputeInstantSnapshotObservedState is the state of the ComputeInstantSnapshot resource as most recently observed in GCP.
// +kcc:observedstate:proto=google.cloud.compute.v1.InstantSnapshot
type ComputeInstantSnapshotObservedState struct {
	// [Output Only] The architecture of the instant snapshot. Valid values are ARM64 or X86_64.
	// Check the Architecture enum for the list of possible values.
	// +kcc:proto:field=google.cloud.compute.v1.InstantSnapshot.architecture
	Architecture *string `json:"architecture,omitempty"`

	// [Output Only] Creation timestamp in RFC3339 text format.
	// +kcc:proto:field=google.cloud.compute.v1.InstantSnapshot.creation_timestamp
	CreationTimestamp *string `json:"creationTimestamp,omitempty"`

	// [Output Only] Size of the source disk, specified in GB.
	// +kcc:proto:field=google.cloud.compute.v1.InstantSnapshot.disk_size_gb
	DiskSizeGB *int64 `json:"diskSizeGB,omitempty"`

	// [Output Only] The unique identifier for the resource. This identifier is defined by the server.
	// +kcc:proto:field=google.cloud.compute.v1.InstantSnapshot.id
	ID *uint64 `json:"id,omitempty"`

	// [Output Only] Type of the resource. Always compute#instantSnapshot for InstantSnapshot resources.
	// +kcc:proto:field=google.cloud.compute.v1.InstantSnapshot.kind
	Kind *string `json:"kind,omitempty"`

	// [Output Only] A fingerprint for the labels being applied to this InstantSnapshot, which is essentially a hash of the labels set used for optimistic locking.
	// +kcc:proto:field=google.cloud.compute.v1.InstantSnapshot.label_fingerprint
	LabelFingerprint *string `json:"labelFingerprint,omitempty"`

	// [Output Only] URL of the region where the instant snapshot resides. You must specify this field as part of the HTTP request URL. It is not settable as a field in the request body.
	// +kcc:proto:field=google.cloud.compute.v1.InstantSnapshot.region
	Region *string `json:"region,omitempty"`

	// [Output Only] Status information for the instant snapshot resource.
	// +kcc:proto:field=google.cloud.compute.v1.InstantSnapshot.resource_status
	ResourceStatus *InstantSnapshotResourceStatus `json:"resourceStatus,omitempty"`

	// Output only. Reserved for future use.
	// +kcc:proto:field=google.cloud.compute.v1.InstantSnapshot.satisfies_pzi
	SatisfiesPzi *bool `json:"satisfiesPzi,omitempty"`

	// [Output Only] Reserved for future use.
	// +kcc:proto:field=google.cloud.compute.v1.InstantSnapshot.satisfies_pzs
	SatisfiesPzs *bool `json:"satisfiesPzs,omitempty"`

	// [Output Only] Server-defined URL for the resource.
	// +kcc:proto:field=google.cloud.compute.v1.InstantSnapshot.self_link
	SelfLink *string `json:"selfLink,omitempty"`

	// [Output Only] Server-defined URL for this resource's resource id.
	// +kcc:proto:field=google.cloud.compute.v1.InstantSnapshot.self_link_with_id
	SelfLinkWithID *string `json:"selfLinkWithID,omitempty"`

	// [Output Only] The ID value of the disk used to create this InstantSnapshot. This value may be used to determine whether the InstantSnapshot was taken from the current or a previous instance of a given disk name.
	// +kcc:proto:field=google.cloud.compute.v1.InstantSnapshot.source_disk_id
	SourceDiskID *string `json:"sourceDiskID,omitempty"`

	// [Output Only] The status of the instantSnapshot. This can be CREATING, DELETING, FAILED, or READY.
	// Check the Status enum for the list of possible values.
	// +kcc:proto:field=google.cloud.compute.v1.InstantSnapshot.status
	Status *string `json:"status,omitempty"`

	// [Output Only] URL of the zone where the instant snapshot resides. You must specify this field as part of the HTTP request URL. It is not settable as a field in the request body.
	// +kcc:proto:field=google.cloud.compute.v1.InstantSnapshot.zone
	Zone *string `json:"zone,omitempty"`
}

// +genclient
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:resource:categories=gcp,shortName=gcpcomputeinstantsnapshot;gcpcomputeinstantsnapshots
// +kubebuilder:subresource:status
// +kubebuilder:metadata:labels="cnrm.cloud.google.com/managed-by-kcc=true"
// +kubebuilder:metadata:labels="cnrm.cloud.google.com/system=true"
// +kubebuilder:printcolumn:name="Age",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:printcolumn:name="Ready",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status",description="When 'True', the most recent reconcile of the resource succeeded"
// +kubebuilder:printcolumn:name="Status",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].reason",description="The reason for the value in 'Ready'"
// +kubebuilder:printcolumn:name="Status Age",type="date",JSONPath=".status.conditions[?(@.type=='Ready')].lastTransitionTime",description="The last transition time for the value in 'Status'"

// ComputeInstantSnapshot is the Schema for the ComputeInstantSnapshot API
// +k8s:openapi-gen=true
type ComputeInstantSnapshot struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   ComputeInstantSnapshotSpec   `json:"spec,omitempty"`
	Status ComputeInstantSnapshotStatus `json:"status,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// ComputeInstantSnapshotList contains a list of ComputeInstantSnapshot
type ComputeInstantSnapshotList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []ComputeInstantSnapshot `json:"items"`
}

func init() {
	SchemeBuilder.Register(&ComputeInstantSnapshot{}, &ComputeInstantSnapshotList{})
}
