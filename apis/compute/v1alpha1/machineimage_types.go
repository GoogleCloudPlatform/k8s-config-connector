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
	computev1beta1 "github.com/GoogleCloudPlatform/k8s-config-connector/apis/compute/v1beta1"
	refsv1beta1 "github.com/GoogleCloudPlatform/k8s-config-connector/apis/refs/v1beta1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/apis/k8s/v1alpha1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

var ComputeMachineImageGVK = GroupVersion.WithKind("ComputeMachineImage")

// ComputeMachineImageSpec defines the desired state of ComputeMachineImage
// +kcc:spec:proto=google.cloud.compute.v1.MachineImage
type ComputeMachineImageSpec struct {
	// The project that this resource belongs to.
	// +required
	ProjectRef *refsv1beta1.ProjectRef `json:"projectRef"`

	// The location of this resource.
	// +required
	Location string `json:"location"`

	// The ComputeMachineImage name. If not given, the metadata.name will be used.
	ResourceID *string `json:"resourceID,omitempty"`

	// An optional description of this resource. Provide this property when you create the resource.
	// +kcc:proto:field=google.cloud.compute.v1.MachineImage.description
	Description *string `json:"description,omitempty"`

	// +kcc:proto:field=google.cloud.compute.v1.MachineImage.guest_flush
	GuestFlush *bool `json:"guestFlush,omitempty"`

	// +kcc:proto:field=google.cloud.compute.v1.MachineImage.labels
	Labels map[string]string `json:"labels,omitempty"`

	// +kcc:proto:field=google.cloud.compute.v1.MachineImage.machine_image_encryption_key
	MachineImageEncryptionKey *CustomerEncryptionKey `json:"machineImageEncryptionKey,omitempty"`

	// +required
	// +kcc:proto:field=google.cloud.compute.v1.MachineImage.source_instance
	SourceInstanceRef *computev1beta1.InstanceRef `json:"sourceInstanceRef"`

	// The regional or multi-regional Cloud Storage bucket location where the machine image is stored.
	// +kcc:proto:field=google.cloud.compute.v1.MachineImage.storage_locations
	StorageLocations []string `json:"storageLocations,omitempty"`
}

// ComputeMachineImageStatus defines the config connector machine state of ComputeMachineImage
type ComputeMachineImageStatus struct {
	/* Conditions represent the latest available observations of the
	   object's current state. */
	Conditions []v1alpha1.Condition `json:"conditions,omitempty"`

	// ObservedGeneration is the generation of the resource that was most recently observed by the Config Connector controller. If this is equal to metadata.generation, then that means that the current reported status reflects the most recent desired state of the resource.
	ObservedGeneration *int64 `json:"observedGeneration,omitempty"`

	// A unique specifier for the ComputeMachineImage resource in GCP.
	ExternalRef *string `json:"externalRef,omitempty"`

	// ObservedState is the state of the resource as most recently observed in GCP.
	ObservedState *ComputeMachineImageObservedState `json:"observedState,omitempty"`
}

// ComputeMachineImageObservedState is the state of the ComputeMachineImage resource as most recently observed in GCP.
// +kcc:observedstate:proto=google.cloud.compute.v1.MachineImage
type ComputeMachineImageObservedState struct {
	// [Output Only] Creation timestamp in RFC3339 text format.
	// +kcc:proto:field=google.cloud.compute.v1.MachineImage.creation_timestamp
	CreationTimestamp *string `json:"creationTimestamp,omitempty"`

	// [Output Only] The unique identifier for the resource type. The server generates this identifier.
	// +kcc:proto:field=google.cloud.compute.v1.MachineImage.id
	ID *uint64 `json:"id,omitempty"`

	// +kcc:proto:field=google.cloud.compute.v1.MachineImage.instance_properties
	InstanceProperties *InstanceProperties `json:"instanceProperties,omitempty"`

	// [Output Only] The resource type, which is always compute#machineImage for machine image.
	// +kcc:proto:field=google.cloud.compute.v1.MachineImage.kind
	Kind *string `json:"kind,omitempty"`

	// +kcc:proto:field=google.cloud.compute.v1.MachineImage.label_fingerprint
	LabelFingerprint *string `json:"labelFingerprint,omitempty"`

	// [Output Only] Reserved for future use.
	// +kcc:proto:field=google.cloud.compute.v1.MachineImage.satisfies_pzi
	SatisfiesPzi *bool `json:"satisfiesPzi,omitempty"`

	// [Output Only] Reserved for future use.
	// +kcc:proto:field=google.cloud.compute.v1.MachineImage.satisfies_pzs
	SatisfiesPzs *bool `json:"satisfiesPzs,omitempty"`

	// +kcc:proto:field=google.cloud.compute.v1.MachineImage.saved_disks
	SavedDisks []SavedDisk `json:"savedDisks,omitempty"`

	// [Output Only] Server-defined URL for the resource.
	// +kcc:proto:field=google.cloud.compute.v1.MachineImage.self_link
	SelfLink *string `json:"selfLink,omitempty"`

	// +kcc:proto:field=google.cloud.compute.v1.MachineImage.source_instance_properties
	SourceInstanceProperties *SourceInstanceProperties `json:"sourceInstanceProperties,omitempty"`

	// [Output Only] The status of the machine image. One of: CREATING, DELETING, INVALID, READY, or UPLOADING.
	// +kcc:proto:field=google.cloud.compute.v1.MachineImage.status
	Status *string `json:"status,omitempty"`

	// [Output Only] Total size of the machine image in bytes.
	// +kcc:proto:field=google.cloud.compute.v1.MachineImage.total_storage_bytes
	TotalStorageBytes *int64 `json:"totalStorageBytes,omitempty"`
}

// +genclient
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:resource:categories=gcp,shortName=gcpcomputemachineimage;gcpcomputemachineimages
// +kubebuilder:subresource:status
// +kubebuilder:metadata:labels="cnrm.cloud.google.com/managed-by-kcc=true"
// +kubebuilder:metadata:labels="cnrm.cloud.google.com/system=true"
// +kubebuilder:metadata:labels="cnrm.cloud.google.com/stability-level=alpha"
// +kubebuilder:printcolumn:name="Age",JSONPath=".metadata.creationTimestamp",type="date"
// +kubebuilder:printcolumn:name="Ready",JSONPath=".status.conditions[?(@.type=='Ready')].status",type="string",description="When 'True', the most recent reconcile of the resource succeeded"
// +kubebuilder:printcolumn:name="Status",JSONPath=".status.conditions[?(@.type=='Ready')].reason",type="string",description="The reason for the value in 'Ready'"
// +kubebuilder:printcolumn:name="Status Age",JSONPath=".status.conditions[?(@.type=='Ready')].lastTransitionTime",type="date",description="The last transition time for the value in 'Status'"

// ComputeMachineImage is the Schema for the ComputeMachineImage API
// +k8s:openapi-gen=true
type ComputeMachineImage struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	// +required
	Spec   ComputeMachineImageSpec   `json:"spec,omitempty"`
	Status ComputeMachineImageStatus `json:"status,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// ComputeMachineImageList contains a list of ComputeMachineImage
type ComputeMachineImageList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []ComputeMachineImage `json:"items"`
}

func init() {
	SchemeBuilder.Register(&ComputeMachineImage{}, &ComputeMachineImageList{})
}
