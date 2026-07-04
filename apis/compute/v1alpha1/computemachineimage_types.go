// Copyright 2026 Google LLC
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package v1alpha1

import (
	computev1beta1 "github.com/GoogleCloudPlatform/k8s-config-connector/apis/compute/v1beta1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/apis/refs"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/apis/k8s/v1alpha1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

var ComputeMachineImageGVK = GroupVersion.WithKind("ComputeMachineImage")

// +kcc:proto=google.cloud.compute.v1.CustomerEncryptionKey
type MachineImageEncryptionKey struct {
	// Immutable. The name of the encryption key that is stored in Google Cloud KMS.
	// +kcc:proto:field=google.cloud.compute.v1.CustomerEncryptionKey.kms_key_name
	KMSKeyName *string `json:"kmsKeyName,omitempty"`

	// Immutable. The service account used for the encryption request for the given KMS key.
	// If absent, the Compute Engine Service Agent service account is used.
	// +kcc:proto:field=google.cloud.compute.v1.CustomerEncryptionKey.kms_key_service_account
	KMSKeyServiceAccount *string `json:"kmsKeyServiceAccount,omitempty"`

	// Immutable. Specifies a 256-bit customer-supplied encryption key, encoded in
	// RFC 4648 base64 to either encrypt or decrypt this resource.
	// +kcc:proto:field=google.cloud.compute.v1.CustomerEncryptionKey.raw_key
	RawKey *string `json:"rawKey,omitempty"`

	// The RFC 4648 base64 encoded SHA-256 hash of the
	// customer-supplied encryption key that protects this resource.
	// +kcc:proto:field=google.cloud.compute.v1.CustomerEncryptionKey.sha256
	Sha256 *string `json:"sha256,omitempty"`
}

// ComputeMachineImageSpec defines the desired state of ComputeMachineImage
// +kcc:spec:proto=google.cloud.compute.v1.MachineImage
type ComputeMachineImageSpec struct {
	// The project that this resource belongs to.
	// +required
	ProjectRef *refs.ProjectRef `json:"projectRef"`

	// Immutable. A text description of the resource.
	// +kcc:proto:field=google.cloud.compute.v1.MachineImage.description
	Description *string `json:"description,omitempty"`

	// Immutable. Specify this to create an application consistent machine image by informing the OS to prepare for the snapshot process.
	// Currently only supported on Windows instances using the Volume Shadow Copy Service (VSS).
	// +kcc:proto:field=google.cloud.compute.v1.MachineImage.guest_flush
	GuestFlush *bool `json:"guestFlush,omitempty"`

	// Immutable. Encrypts the machine image using a customer-supplied encryption key.
	// +kcc:proto:field=google.cloud.compute.v1.MachineImage.machine_image_encryption_key
	MachineImageEncryptionKey *MachineImageEncryptionKey `json:"machineImageEncryptionKey,omitempty"`

	// Immutable. Optional. The name of the resource. Used for creation and acquisition. When unset, the value of `metadata.name` is used as the default.
	ResourceID *string `json:"resourceID,omitempty"`

	// The source instance used to create the machine image.
	// +required
	// +kcc:proto:field=google.cloud.compute.v1.MachineImage.source_instance
	SourceInstanceRef *computev1beta1.InstanceRef `json:"sourceInstanceRef"`
}

// ComputeMachineImageStatus defines the config connector machine state of ComputeMachineImage
// +kcc:proto=google.cloud.compute.v1.MachineImage
type ComputeMachineImageStatus struct {
	/* Conditions represent the latest available observations of the
	   object's current state. */
	Conditions []v1alpha1.Condition `json:"conditions,omitempty"`

	// ObservedGeneration is the generation of the resource that was most recently observed by the Config Connector controller. If this is equal to metadata.generation, then that means that the current reported status reflects the most recent desired state of the resource.
	ObservedGeneration *int64 `json:"observedGeneration,omitempty"`

	// Server-defined URL for the resource.
	// +kcc:proto:field=google.cloud.compute.v1.MachineImage.self_link
	SelfLink *string `json:"selfLink,omitempty"`

	// The regional or multi-regional Cloud Storage bucket location where the machine image is stored.
	// +kcc:proto:field=google.cloud.compute.v1.MachineImage.storage_locations
	StorageLocations []string `json:"storageLocations,omitempty"`
}

// +genclient
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:resource:categories=gcp,shortName=gcpcomputemachineimage;gcpcomputemachineimages
// +kubebuilder:subresource:status
// +kubebuilder:metadata:labels="cnrm.cloud.google.com/managed-by-kcc=true"
// +kubebuilder:metadata:labels="cnrm.cloud.google.com/stability-level=alpha"
// +kubebuilder:metadata:labels="cnrm.cloud.google.com/system=true"
// +kubebuilder:metadata:labels="cnrm.cloud.google.com/tf2crd=true"
// +kubebuilder:printcolumn:name="Age",JSONPath=".metadata.creationTimestamp",type="date"
// +kubebuilder:printcolumn:name="Ready",JSONPath=".status.conditions[?(@.type=='Ready')].status",type="string",description="When 'True', the most recent reconcile of the resource succeeded"
// +kubebuilder:printcolumn:name="Status",JSONPath=".status.conditions[?(@.type=='Ready')].reason",type="string",description="The reason for the value in 'Ready'"
// +kubebuilder:printcolumn:name="Status Age",JSONPath=".status.conditions[?(@.type=='Ready')].lastTransitionTime",type="date",description="The last transition time for the value in 'Status'"

// ComputeMachineImage is the Schema for the computemachineimages API
type ComputeMachineImage struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   ComputeMachineImageSpec   `json:"spec"`
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
