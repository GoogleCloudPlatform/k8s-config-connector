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
	kmsv1beta1 "github.com/GoogleCloudPlatform/k8s-config-connector/apis/kms/v1beta1"
	refsv1beta1 "github.com/GoogleCloudPlatform/k8s-config-connector/apis/refs/v1beta1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/apis/refs/v1beta1/secret"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/apis/k8s/v1alpha1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// +kcc:proto=google.cloud.compute.v1.CustomerEncryptionKey
type SnapshotSnapshotEncryptionKey struct {
	/* The encryption key that is stored in Google Cloud KMS. */
	// +optional
	KmsKeyRef *kmsv1beta1.KMSCryptoKeyRef `json:"kmsKeyRef,omitempty"`

	/* The service account used for the encryption request for the given KMS key.
	If absent, the Compute Engine Service Agent service account is used. */
	// +optional
	KmsKeyServiceAccountRef *refsv1beta1.IAMServiceAccountRef `json:"kmsKeyServiceAccountRef,omitempty"`

	/* Immutable. Specifies a 256-bit customer-supplied encryption key, encoded in
	RFC 4648 base64 to either encrypt or decrypt this resource. */
	// +optional
	RawKey *secret.Legacy `json:"rawKey,omitempty"`

	/* The RFC 4648 base64 encoded SHA-256 hash of the customer-supplied
	encryption key that protects this resource. */
	// +optional
	Sha256 *string `json:"sha256,omitempty"`
}

// +kcc:proto=google.cloud.compute.v1.CustomerEncryptionKey
type SnapshotSourceDiskEncryptionKey struct {
	/* The service account used for the encryption request for the given KMS key.
	If absent, the Compute Engine Service Agent service account is used. */
	// +optional
	KmsKeyServiceAccountRef *refsv1beta1.IAMServiceAccountRef `json:"kmsKeyServiceAccountRef,omitempty"`

	/* Immutable. Specifies a 256-bit customer-supplied encryption key, encoded in
	RFC 4648 base64 to either encrypt or decrypt this resource. */
	// +optional
	RawKey *secret.Legacy `json:"rawKey,omitempty"`
}

// ComputeSnapshotSpec defines the desired state of ComputeSnapshot
// +kcc:spec:proto=google.cloud.compute.v1.Snapshot
type ComputeSnapshotSpec struct {
	/* Immutable. Creates the new snapshot in the snapshot chain labeled with the
	specified name. The chain name must be 1-63 characters long and
	comply with RFC1035. This is an uncommon option only for advanced
	service owners who needs to create separate snapshot chains, for
	example, for chargeback tracking.  When you describe your snapshot
	resource, this field is visible only if it has a non-empty value. */
	// +optional
	// +kcc:proto:field=google.cloud.compute.v1.Snapshot.chain_name
	ChainName *string `json:"chainName,omitempty"`

	/* Immutable. An optional description of this resource. */
	// +optional
	// +kcc:proto:field=google.cloud.compute.v1.Snapshot.description
	Description *string `json:"description,omitempty"`

	/* Immutable. Optional. The name of the resource. Used for creation and acquisition. When unset, the value of `metadata.name` is used as the default. */
	// +optional
	ResourceID *string `json:"resourceID,omitempty"`

	/* Immutable. Encrypts the snapshot using a customer-supplied encryption key.

	After you encrypt a snapshot using a customer-supplied key, you must
	provide the same key if you use the snapshot later. For example, you
	must provide the encryption key when you create a disk from the
	encrypted snapshot in a future request.

	Customer-supplied encryption keys do not protect access to metadata of
	the snapshot.

	If you do not provide an encryption key when creating the snapshot,
	then the snapshot will be encrypted using an automatically generated
	key and you do not need to provide a key to use the snapshot later. */
	// +optional
	// +kcc:proto:field=google.cloud.compute.v1.Snapshot.snapshot_encryption_key
	SnapshotEncryptionKey *SnapshotSnapshotEncryptionKey `json:"snapshotEncryptionKey,omitempty"`

	/* Immutable. The customer-supplied encryption key of the source snapshot. Required
	if the source snapshot is protected by a customer-supplied encryption
	key. */
	// +optional
	// +kcc:proto:field=google.cloud.compute.v1.Snapshot.source_disk_encryption_key
	SourceDiskEncryptionKey *SnapshotSourceDiskEncryptionKey `json:"sourceDiskEncryptionKey,omitempty"`

	/* A reference to the disk used to create this snapshot. */
	// +kcc:proto:field=google.cloud.compute.v1.Snapshot.source_disk
	SourceDiskRef *ComputeDiskRef `json:"sourceDiskRef"`

	/* Immutable. Cloud Storage bucket storage location of the snapshot (regional or multi-regional). */
	// +optional
	// +kcc:proto:field=google.cloud.compute.v1.Snapshot.storage_locations
	StorageLocations []string `json:"storageLocations,omitempty"`

	/* Immutable. A reference to the zone where the disk is hosted. */
	// +optional
	// +kcc:proto:field=google.cloud.compute.v1.Snapshot.zone
	Zone *string `json:"zone,omitempty"`
}

// ComputeSnapshotStatus defines the config connector machine state of ComputeSnapshot
// +kcc:status:proto=google.cloud.compute.v1.Snapshot
type ComputeSnapshotStatus struct {
	/* Conditions represent the latest available observations of the
	   ComputeSnapshot's current state. */
	// +optional
	Conditions []v1alpha1.Condition `json:"conditions,omitempty"`

	/* Creation timestamp in RFC3339 text format. */
	// +optional
	// +kcc:proto:field=google.cloud.compute.v1.Snapshot.creation_timestamp
	CreationTimestamp *string `json:"creationTimestamp,omitempty"`

	/* Size of the snapshot, specified in GB. */
	// +optional
	// +kcc:proto:field=google.cloud.compute.v1.Snapshot.disk_size_gb
	DiskSizeGb *int64 `json:"diskSizeGb,omitempty"`

	/* The fingerprint used for optimistic locking of this resource. Used
	internally during updates. */
	// +optional
	// +kcc:proto:field=google.cloud.compute.v1.Snapshot.label_fingerprint
	LabelFingerprint *string `json:"labelFingerprint,omitempty"`

	/* A list of public visible licenses that apply to this snapshot. This
	can be because the original image had licenses attached (such as a
	Windows image).  snapshotEncryptionKey nested object Encrypts the
	snapshot using a customer-supplied encryption key. */
	// +optional
	// +kcc:proto:field=google.cloud.compute.v1.Snapshot.licenses
	Licenses []string `json:"licenses,omitempty"`

	/* ObservedGeneration is the generation of the resource that was most recently observed by the Config Connector controller. If this is equal to metadata.generation, then that means that the current reported status reflects the most recent desired state of the resource. */
	// +optional
	ObservedGeneration *int64 `json:"observedGeneration,omitempty"`

	// +optional
	// +kcc:proto:field=google.cloud.compute.v1.Snapshot.self_link
	SelfLink *string `json:"selfLink,omitempty"`

	/* The unique identifier for the resource. */
	// +optional
	// +kcc:proto:field=google.cloud.compute.v1.Snapshot.id
	SnapshotId *int64 `json:"snapshotId,omitempty"`

	/* A size of the storage used by the snapshot. As snapshots share
	storage, this number is expected to change with snapshot
	creation/deletion. */
	// +optional
	// +kcc:proto:field=google.cloud.compute.v1.Snapshot.storage_bytes
	StorageBytes *int64 `json:"storageBytes,omitempty"`
}

// +genclient
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:resource:categories=gcp,shortName=gcpcomputesnapshot;gcpcomputesnapshots
// +kubebuilder:subresource:status
// +kubebuilder:metadata:labels="cnrm.cloud.google.com/managed-by-kcc=true"
// +kubebuilder:metadata:labels="cnrm.cloud.google.com/stability-level=stable"
// +kubebuilder:metadata:labels="cnrm.cloud.google.com/system=true"
// +kubebuilder:metadata:labels="cnrm.cloud.google.com/tf2crd=true"
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
