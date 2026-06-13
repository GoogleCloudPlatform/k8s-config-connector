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

package v1beta1

import (
	refsv1beta1 "github.com/GoogleCloudPlatform/k8s-config-connector/apis/refs/v1beta1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/apis/k8s/v1alpha1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// ComputeImageSpec defines the desired state of ComputeImage
// +kcc:spec:proto=google.cloud.compute.v1.Image
type ComputeImageSpec struct {
	// The project that this resource belongs to.
	ProjectRef *refsv1beta1.ProjectRef `json:"projectRef"`

	// The location of this resource.
	Location string `json:"location"`

	// The ComputeImage name. If not given, the metadata.name will be used.
	ResourceID *string `json:"resourceID,omitempty"`

	// An optional description of this resource. Provide this property when you create the resource.
	// +optional
	// +kcc:proto:field=google.cloud.compute.v1.Image.description
	Description *string `json:"description,omitempty"`

	// The source disk to create this image based on.
	// +optional
	// +kcc:proto:field=google.cloud.compute.v1.Image.source_disk
	DiskRef *refsv1beta1.ComputeDiskRef `json:"diskRef,omitempty"`

	// Size of the image when restored onto a persistent disk (in GB).
	// +optional
	// +kcc:proto:field=google.cloud.compute.v1.Image.disk_size_gb
	DiskSizeGb *int64 `json:"diskSizeGb,omitempty"`

	// The name of the image family to which this image belongs.
	// +optional
	// +kcc:proto:field=google.cloud.compute.v1.Image.family
	Family *string `json:"family,omitempty"`

	// A list of features to enable on the guest operating system.
	// +optional
	// +kcc:proto:field=google.cloud.compute.v1.Image.guest_os_features
	GuestOsFeatures []ImageGuestOSFeatures `json:"guestOsFeatures,omitempty"`

	// Encrypts the image using a customer-supplied encryption key.
	// +optional
	// +kcc:proto:field=google.cloud.compute.v1.Image.image_encryption_key
	ImageEncryptionKey *ImageImageEncryptionKey `json:"imageEncryptionKey,omitempty"`

	// Any applicable license URI.
	// +optional
	// +kcc:proto:field=google.cloud.compute.v1.Image.licenses
	Licenses []string `json:"licenses,omitempty"`

	// The parameters of the raw disk image.
	// +optional
	// +kcc:proto:field=google.cloud.compute.v1.Image.raw_disk
	RawDisk *ImageRawDisk `json:"rawDisk,omitempty"`

	// The source image used to create this image.
	// +optional
	// +kcc:proto:field=google.cloud.compute.v1.Image.source_image
	SourceImageRef *ComputeImageRef `json:"sourceImageRef,omitempty"`

	// The source snapshot used to create this image.
	// +optional
	// +kcc:proto:field=google.cloud.compute.v1.Image.source_snapshot
	SourceSnapshotRef *refsv1beta1.ComputeSnapshotRef `json:"sourceSnapshotRef,omitempty"`

	// Cloud Storage bucket storage location of the image (regional or multi-regional).
	// +optional
	// +kcc:proto:field=google.cloud.compute.v1.Image.storage_locations
	StorageLocations []string `json:"storageLocations,omitempty"`
}

type ImageGuestOSFeatures struct {
	// The type of supported feature.
	// +optional
	// +kcc:proto:field=google.cloud.compute.v1.GuestOsFeature.type
	Type *string `json:"type,omitempty"`
}

type ImageImageEncryptionKey struct {
	// The self link of the encryption key that is stored in Google Cloud KMS.
	// +optional
	KmsKeySelfLinkRef *refsv1beta1.KMSCryptoKeyRef `json:"kmsKeySelfLinkRef,omitempty"`

	// The service account being used for the encryption request for the given KMS key. If absent, the Compute Engine default service account is used.
	// +optional
	KmsKeyServiceAccountRef *refsv1beta1.IAMServiceAccountRef `json:"kmsKeyServiceAccountRef,omitempty"`
}

type ImageRawDisk struct {
	// The format used to encode and transmit the block device, which should be TAR.
	// +optional
	// +kcc:proto:field=google.cloud.compute.v1.RawDisk.container_type
	ContainerType *string `json:"containerType,omitempty"`

	// An optional SHA1 checksum of the disk image before unpackaging.
	// +optional
	// +kcc:proto:field=google.cloud.compute.v1.RawDisk.sha1_checksum
	Sha1 *string `json:"sha1,omitempty"`

	// The full Google Cloud Storage URL where disk storage is stored.
	// +optional
	// +kcc:proto:field=google.cloud.compute.v1.RawDisk.source
	Source *string `json:"source,omitempty"`
}

// ComputeImageStatus defines the config connector machine state of ComputeImage
type ComputeImageStatus struct {
	/* Conditions represent the latest available observations of the
	   object's current state. */
	Conditions []v1alpha1.Condition `json:"conditions,omitempty"`

	// ObservedGeneration is the generation of the resource that was most recently observed by the Config Connector controller. If this is equal to metadata.generation, then that means that the current reported status reflects the most recent desired state of the resource.
	ObservedGeneration *int64 `json:"observedGeneration,omitempty"`

	// A unique specifier for the ComputeImage resource in GCP.
	ExternalRef *string `json:"externalRef,omitempty"`

	// ObservedState is the state of the resource as most recently observed in GCP.
	ObservedState *ComputeImageObservedState `json:"observedState,omitempty"`

	// Size of the image tar.gz archive stored in Google Cloud Storage (in bytes).
	// +optional
	// +kcc:proto:field=google.cloud.compute.v1.Image.archive_size_bytes
	ArchiveSizeBytes *int64 `json:"archiveSizeBytes,omitempty"`

	// Creation timestamp in RFC3339 text format.
	// +optional
	// +kcc:proto:field=google.cloud.compute.v1.Image.creation_timestamp
	CreationTimestamp *string `json:"creationTimestamp,omitempty"`

	// The fingerprint used for optimistic locking of this resource.
	// +optional
	// +kcc:proto:field=google.cloud.compute.v1.Image.label_fingerprint
	LabelFingerprint *string `json:"labelFingerprint,omitempty"`

	// Server-defined URL for the resource.
	// +optional
	// +kcc:proto:field=google.cloud.compute.v1.Image.self_link
	SelfLink *string `json:"selfLink,omitempty"`
}

// ComputeImageObservedState is the state of the ComputeImage resource as most recently observed in GCP.
// +kcc:observedstate:proto=google.cloud.compute.v1.Image
type ComputeImageObservedState struct {
}

// +genclient
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:resource:categories=gcp,shortName=gcpcomputeimage;gcpcomputeimages
// +kubebuilder:subresource:status
// +kubebuilder:metadata:labels="cnrm.cloud.google.com/managed-by-kcc=true"
// +kubebuilder:metadata:labels="cnrm.cloud.google.com/system=true"
// +kubebuilder:metadata:labels="cnrm.cloud.google.com/stability-level=stable"
// +kubebuilder:printcolumn:name="Age",JSONPath=".metadata.creationTimestamp",type="date"
// +kubebuilder:printcolumn:name="Ready",JSONPath=".status.conditions[?(@.type=='Ready')].status",type="string",description="When 'True', the most recent reconcile of the resource succeeded"
// +kubebuilder:printcolumn:name="Status",JSONPath=".status.conditions[?(@.type=='Ready')].reason",type="string",description="The reason for the value in 'Ready'"
// +kubebuilder:printcolumn:name="Status Age",JSONPath=".status.conditions[?(@.type=='Ready')].lastTransitionTime",type="date",description="The last transition time for the value in 'Status'"

// ComputeImage is the Schema for the ComputeImage API
// +k8s:openapi-gen=true
type ComputeImage struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	// +required
	Spec   ComputeImageSpec   `json:"spec,omitempty"`
	Status ComputeImageStatus `json:"status,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// ComputeImageList contains a list of ComputeImage
type ComputeImageList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []ComputeImage `json:"items"`
}

func init() {
	SchemeBuilder.Register(&ComputeImage{}, &ComputeImageList{})
}
