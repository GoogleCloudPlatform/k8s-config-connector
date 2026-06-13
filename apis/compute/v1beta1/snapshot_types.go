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

var ComputeSnapshotGVK = GroupVersion.WithKind("ComputeSnapshot")

// +kcc:proto=google.cloud.compute.v1.CustomerEncryptionKey
type CustomerEncryptionKey struct {
	// The name of the encryption key that is stored in Google Cloud KMS.
	// +optional
	KmsKeyRef *refsv1beta1.KMSCryptoKeyRef `json:"kmsKeyRef,omitempty"`

	// The service account being used for the encryption request for the given KMS key. If absent, the Compute Engine default service account is used.
	// +optional
	// +kcc:proto:field=google.cloud.compute.v1.CustomerEncryptionKey.kms_key_service_account
	KmsKeyServiceAccountRef *refsv1beta1.IAMServiceAccountRef `json:"kmsKeyServiceAccountRef,omitempty"`

	// Specifies a 256-bit customer-supplied encryption key, encoded in RFC 4648 base64 to either encrypt or decrypt this resource.
	// +optional
	// +kcc:proto:field=google.cloud.compute.v1.CustomerEncryptionKey.raw_key
	RawKey *string `json:"rawKey,omitempty"`

	// Specifies an RFC 4648 base64 encoded, RSA-wrapped 2048-bit customer-supplied encryption key to either encrypt or decrypt this resource.
	// +optional
	// +kcc:proto:field=google.cloud.compute.v1.CustomerEncryptionKey.rsa_encrypted_key
	RsaEncryptedKey *string `json:"rsaEncryptedKey,omitempty"`

	// [Output only] The RFC 4648 base64 encoded SHA-256 hash of the customer-supplied encryption key that protects this resource.
	// +optional
	// +kcc:proto:field=google.cloud.compute.v1.CustomerEncryptionKey.sha256
	Sha256 *string `json:"sha256,omitempty"`
}

// ComputeSnapshotSpec defines the desired state of ComputeSnapshot
// +kcc:spec:proto=google.cloud.compute.v1.Snapshot
type ComputeSnapshotSpec struct {
	// The project that this resource belongs to.
	ProjectRef *refsv1beta1.ProjectRef `json:"projectRef"`

	// The location of this resource.
	Location string `json:"location"`

	// Immutable. Creates the new snapshot in the snapshot chain labeled with the
	// specified name. The chain name must be 1-63 characters long and
	// comply with RFC1035. This is an uncommon option only for advanced
	// service owners who needs to create separate snapshot chains, for
	// example, for chargeback tracking. When you describe your snapshot
	// resource, this field is visible only if it has a non-empty value.
	// +optional
	// +kcc:proto:field=google.cloud.compute.v1.Snapshot.chain_name
	ChainName *string `json:"chainName,omitempty"`

	// Immutable. An optional description of this resource.
	// +optional
	// +kcc:proto:field=google.cloud.compute.v1.Snapshot.description
	Description *string `json:"description,omitempty"`

	// The ComputeSnapshot name. If not given, the metadata.name will be used.
	// +optional
	ResourceID *string `json:"resourceID,omitempty"`

	// Immutable. Encrypts the snapshot using a customer-supplied encryption key.
	// +optional
	// +kcc:proto:field=google.cloud.compute.v1.Snapshot.snapshot_encryption_key
	SnapshotEncryptionKey *CustomerEncryptionKey `json:"snapshotEncryptionKey,omitempty"`

	// Immutable. The customer-supplied encryption key of the source disk. Required
	// if the source disk is protected by a customer-supplied encryption key.
	// +optional
	// +kcc:proto:field=google.cloud.compute.v1.Snapshot.source_disk_encryption_key
	SourceDiskEncryptionKey *CustomerEncryptionKey `json:"sourceDiskEncryptionKey,omitempty"`

	// A reference to the disk used to create this snapshot.
	// +required
	SourceDiskRef *refsv1beta1.ComputeDiskRef `json:"sourceDiskRef"`

	// Immutable. Cloud Storage bucket storage location of the snapshot (regional or multi-regional).
	// +optional
	// +kcc:proto:field=google.cloud.compute.v1.Snapshot.storage_locations
	StorageLocations []string `json:"storageLocations,omitempty"`

	// Immutable. A reference to the zone where the disk is hosted.
	// +optional
	Zone *string `json:"zone,omitempty"`
}

// ComputeSnapshotStatus defines the config connector machine state of ComputeSnapshot
type ComputeSnapshotStatus struct {
	/* Conditions represent the latest available observations of the
	   object's current state. */
	Conditions []v1alpha1.Condition `json:"conditions,omitempty"`

	// ObservedGeneration is the generation of the resource that was most recently observed by the Config Connector controller. If this is equal to metadata.generation, then that means that the current reported status reflects the most recent desired state of the resource.
	ObservedGeneration *int64 `json:"observedGeneration,omitempty"`

	// A unique specifier for the ComputeSnapshot resource in GCP.
	ExternalRef *string `json:"externalRef,omitempty"`

	// ObservedState is the state of the resource as most recently observed in GCP.
	ObservedState *ComputeSnapshotObservedState `json:"observedState,omitempty"`
}

// ComputeSnapshotObservedState is the state of the ComputeSnapshot resource as most recently observed in GCP.
// +kcc:observedstate:proto=google.cloud.compute.v1.Snapshot
type ComputeSnapshotObservedState struct {
	// [Output Only] Size in bytes of the snapshot at creation time.
	// +optional
	// +kcc:proto:field=google.cloud.compute.v1.Snapshot.creation_size_bytes
	CreationSizeBytes *int64 `json:"creationSizeBytes,omitempty"`

	// [Output Only] Creation timestamp in RFC3339 text format.
	// +optional
	// +kcc:proto:field=google.cloud.compute.v1.Snapshot.creation_timestamp
	CreationTimestamp *string `json:"creationTimestamp,omitempty"`

	// [Output Only] Size of the source disk, specified in GB.
	// +optional
	// +kcc:proto:field=google.cloud.compute.v1.Snapshot.disk_size_gb
	DiskSizeGB *int64 `json:"diskSizeGb,omitempty"`

	// [Output Only] Number of bytes downloaded to restore a snapshot to a disk.
	// +optional
	// +kcc:proto:field=google.cloud.compute.v1.Snapshot.download_bytes
	DownloadBytes *int64 `json:"downloadBytes,omitempty"`

	// Whether this snapshot is created from a confidential compute mode disk. [Output Only]: This field is not set by user, but from source disk.
	// +optional
	// +kcc:proto:field=google.cloud.compute.v1.Snapshot.enable_confidential_compute
	EnableConfidentialCompute *bool `json:"enableConfidentialCompute,omitempty"`

	// A fingerprint for the labels being applied to this snapshot, which is essentially a hash of the labels set used for optimistic locking.
	// +optional
	// +kcc:proto:field=google.cloud.compute.v1.Snapshot.label_fingerprint
	LabelFingerprint *string `json:"labelFingerprint,omitempty"`

	// [Output Only] A list of public visible licenses that apply to this snapshot.
	// +optional
	// +kcc:proto:field=google.cloud.compute.v1.Snapshot.licenses
	Licenses []string `json:"licenses,omitempty"`

	// [Output Only] The unique identifier for the resource. This identifier is defined by the server.
	// +optional
	// +kcc:proto:field=google.cloud.compute.v1.Snapshot.id
	SnapshotID *int64 `json:"snapshotId,omitempty"`

	// [Output Only] A size of the storage used by the snapshot. As snapshots share storage, this number is expected to change with snapshot creation/deletion.
	// +optional
	// +kcc:proto:field=google.cloud.compute.v1.Snapshot.storage_bytes
	StorageBytes *int64 `json:"storageBytes,omitempty"`

	// [Output Only] An indicator whether storageBytes is in a stable state or it is being adjusted as a result of shared storage reallocation.
	// +optional
	// +kcc:proto:field=google.cloud.compute.v1.Snapshot.storage_bytes_status
	StorageBytesStatus *string `json:"storageBytesStatus,omitempty"`
}

// +genclient
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:resource:categories=gcp,shortName=gcpcomputesnapshot;gcpcomputesnapshots
// +kubebuilder:subresource:status
// +kubebuilder:metadata:labels="cnrm.cloud.google.com/managed-by-kcc=true"
// +kubebuilder:metadata:labels="cnrm.cloud.google.com/stability-level=stable"
// +kubebuilder:metadata:labels="cnrm.cloud.google.com/system=true"
// +kubebuilder:printcolumn:name="Age",JSONPath=".metadata.creationTimestamp",type="date"
// +kubebuilder:printcolumn:name="Ready",JSONPath=".status.conditions[?(@.type=='Ready')].status",type="string",description="When 'True', the most recent reconcile of the resource succeeded"
// +kubebuilder:printcolumn:name="Status",JSONPath=".status.conditions[?(@.type=='Ready')].reason",type="string",description="The reason for the value in 'Ready'"
// +kubebuilder:printcolumn:name="Status Age",JSONPath=".status.conditions[?(@.type=='Ready')].lastTransitionTime",type="date",description="The last transition time for the value in 'Status'"

// ComputeSnapshot is the Schema for the ComputeSnapshot API
// +k8s:openapi-gen=true
type ComputeSnapshot struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	// +required
	Spec   ComputeSnapshotSpec   `json:"spec,omitempty"`
	Status ComputeSnapshotStatus `json:"status,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// ComputeSnapshotList contains a list of ComputeSnapshot
type ComputeSnapshotList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []ComputeSnapshot `json:"items"`
}

func init() {
	SchemeBuilder.Register(&ComputeSnapshot{}, &ComputeSnapshotList{})
}
