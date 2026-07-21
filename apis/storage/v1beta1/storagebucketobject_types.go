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

var StorageBucketObjectGVK = GroupVersion.WithKind("StorageBucketObject")

// StorageBucketObjectSpec defines the desired state of StorageBucketObject
// +kcc:spec:proto=google.storage.v1.Object
type StorageBucketObjectSpec struct {
	// The project that this resource belongs to.
	ProjectRef *refsv1beta1.ProjectRef `json:"projectRef"`

	// The StorageBucket containing this object.
	// +required
	BucketRef *StorageBucketRef `json:"bucketRef"`

	// The StorageBucketObject name. If not given, the metadata.name will be used.
	ResourceID *string `json:"resourceID,omitempty"`

	// Cache-Control directive for the object data, matching
	//  [https://tools.ietf.org/html/rfc7234#section-5.2"][RFC 7234 §5.2].
	// +kcc:proto:field=google.storage.v1.Object.cache_control
	CacheControl *string `json:"cacheControl,omitempty"`

	// Content-Disposition of the object data, matching
	//  [https://tools.ietf.org/html/rfc6266][RFC 6266].
	// +kcc:proto:field=google.storage.v1.Object.content_disposition
	ContentDisposition *string `json:"contentDisposition,omitempty"`

	// Content-Encoding of the object data, matching
	//  [https://tools.ietf.org/html/rfc7231#section-3.1.2.2][RFC 7231 §3.1.2.2]
	// +kcc:proto:field=google.storage.v1.Object.content_encoding
	ContentEncoding *string `json:"contentEncoding,omitempty"`

	// Content-Language of the object data, matching
	//  [https://tools.ietf.org/html/rfc7231#section-3.1.3.2][RFC 7231 §3.1.3.2].
	// +kcc:proto:field=google.storage.v1.Object.content_language
	ContentLanguage *string `json:"contentLanguage,omitempty"`

	// Content-Type of the object data, matching
	//  [https://tools.ietf.org/html/rfc7231#section-3.1.1.5][RFC 7231 §3.1.1.5].
	// +kcc:proto:field=google.storage.v1.Object.content_type
	ContentType *string `json:"contentType,omitempty"`

	// Storage class of the object.
	// +kcc:proto:field=google.storage.v1.Object.storage_class
	StorageClass *string `json:"storageClass,omitempty"`

	// A reference to the KMS Crypto Key that will be used to encrypt the object.
	// +kcc:proto:field=google.storage.v1.Object.kms_key_name
	KmsKeyRef *refsv1beta1.KMSCryptoKeyRef `json:"kmsKeyRef,omitempty"`

	// Whether an object is under temporary hold. While this flag is set to true,
	//  the object is protected against deletion and overwrites.
	// +kcc:proto:field=google.storage.v1.Object.temporary_hold
	TemporaryHold *bool `json:"temporaryHold,omitempty"`

	// Whether an object is under event-based hold.
	// +kcc:proto:field=google.storage.v1.Object.event_based_hold
	EventBasedHold *bool `json:"eventBasedHold,omitempty"`

	// User-provided metadata, in key/value pairs.
	// +kcc:proto:field=google.storage.v1.Object.metadata
	Metadata map[string]string `json:"metadata,omitempty"`

	// A user-specified timestamp set on an object.
	// +kcc:proto:field=google.storage.v1.Object.custom_time
	CustomTime *string `json:"customTime,omitempty"`

	// Access controls on the object.
	// +kcc:proto:field=google.storage.v1.Object.acl
	Acl []ObjectAccessControl `json:"acl,omitempty"`
}

// StorageBucketObjectStatus defines the config connector machine state of StorageBucketObject
type StorageBucketObjectStatus struct {
	/* Conditions represent the latest available observations of the
	   object's current state. */
	Conditions []v1alpha1.Condition `json:"conditions,omitempty"`

	// ObservedGeneration is the generation of the resource that was most recently observed by the Config Connector controller. If this is equal to metadata.generation, then that means that the current reported status reflects the most recent desired state of the resource.
	ObservedGeneration *int64 `json:"observedGeneration,omitempty"`

	// A unique specifier for the StorageBucketObject resource in GCP.
	ExternalRef *string `json:"externalRef,omitempty"`

	// ObservedState is the state of the resource as most recently observed in GCP.
	ObservedState *StorageBucketObjectObservedState `json:"observedState,omitempty"`
}

// StorageBucketObjectObservedState is the state of the StorageBucketObject resource as most recently observed in GCP.
// +kcc:observedstate:proto=google.storage.v1.Object
type StorageBucketObjectObservedState struct {
	// The version of the metadata for this object at this generation. Used for
	//  preconditions and for detecting changes in metadata.
	// +kcc:proto:field=google.storage.v1.Object.metageneration
	Metageneration *int64 `json:"metageneration,omitempty"`

	// The deletion time of the object. Will be returned if and only if this
	//  version of the object has been deleted.
	// +kcc:proto:field=google.storage.v1.Object.time_deleted
	TimeDeleted *string `json:"timeDeleted,omitempty"`

	// Content-Length of the object data in bytes, matching
	//  [https://tools.ietf.org/html/rfc7230#section-3.3.2][RFC 7230 §3.3.2].
	// +kcc:proto:field=google.storage.v1.Object.size
	Size *int64 `json:"size,omitempty"`

	// The creation time of the object.
	// +kcc:proto:field=google.storage.v1.Object.time_created
	TimeCreated *string `json:"timeCreated,omitempty"`

	// CRC32C checksum.
	// +kcc:proto:field=google.storage.v1.Object.crc32c
	Crc32C *uint32 `json:"crc32c,omitempty"`

	// Number of underlying components that make up this object. Components are
	//  accumulated by compose operations.
	// +kcc:proto:field=google.storage.v1.Object.component_count
	ComponentCount *int32 `json:"componentCount,omitempty"`

	// MD5 hash of the data; encoded using base64 as per
	//  [https://tools.ietf.org/html/rfc4648#section-4][RFC 4648 §4].
	// +kcc:proto:field=google.storage.v1.Object.md5_hash
	Md5Hash *string `json:"md5Hash,omitempty"`

	// HTTP 1.1 Entity tag for the object. See
	//  [https://tools.ietf.org/html/rfc7232#section-2.3][RFC 7232 §2.3].
	// +kcc:proto:field=google.storage.v1.Object.etag
	Etag *string `json:"etag,omitempty"`

	// The modification time of the object metadata.
	// +kcc:proto:field=google.storage.v1.Object.updated
	Updated *string `json:"updated,omitempty"`

	// The time at which the object's storage class was last changed. When the
	//  object is initially created, it will be set to time_created.
	// +kcc:proto:field=google.storage.v1.Object.time_storage_class_updated
	TimeStorageClassUpdated *string `json:"timeStorageClassUpdated,omitempty"`

	// A server-determined value that specifies the earliest time that the
	//  object's retention period expires.
	// +kcc:proto:field=google.storage.v1.Object.retention_expiration_time
	RetentionExpirationTime *string `json:"retentionExpirationTime,omitempty"`

	// The content generation of this object. Used for object versioning.
	// +kcc:proto:field=google.storage.v1.Object.generation
	Generation *int64 `json:"generation,omitempty"`

	// The owner of the object. This will always be the uploader of the object.
	// +kcc:proto:field=google.storage.v1.Object.owner
	Owner *Owner `json:"owner,omitempty"`

	// Metadata of customer-supplied encryption key, if the object is encrypted by
	//  such a key.
	// +kcc:proto:field=google.storage.v1.Object.customer_encryption
	CustomerEncryption *Object_CustomerEncryption `json:"customerEncryption,omitempty"`
}

// +genclient
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:resource:categories=gcp,shortName=gcpstoragebucketobject;gcpstoragebucketobjects
// +kubebuilder:subresource:status
// +kubebuilder:metadata:labels="cnrm.cloud.google.com/managed-by-kcc=true"
// +kubebuilder:metadata:labels="cnrm.cloud.google.com/system=true"
// +kubebuilder:metadata:labels="cnrm.cloud.google.com/stability-level=alpha"
// +kubebuilder:printcolumn:name="Age",JSONPath=".metadata.creationTimestamp",type="date"
// +kubebuilder:printcolumn:name="Ready",JSONPath=".status.conditions[?(@.type=='Ready')].status",type="string",description="When 'True', the most recent reconcile of the resource succeeded"
// +kubebuilder:printcolumn:name="Status",JSONPath=".status.conditions[?(@.type=='Ready')].reason",type="string",description="The reason for the value in 'Ready'"
// +kubebuilder:printcolumn:name="Status Age",JSONPath=".status.conditions[?(@.type=='Ready')].lastTransitionTime",type="date",description="The last transition time for the value in 'Status'"

// StorageBucketObject is the Schema for the StorageBucketObject API
// +k8s:openapi-gen=true
type StorageBucketObject struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	// +required
	Spec   StorageBucketObjectSpec   `json:"spec,omitempty"`
	Status StorageBucketObjectStatus `json:"status,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// StorageBucketObjectList contains a list of StorageBucketObject
type StorageBucketObjectList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []StorageBucketObject `json:"items"`
}

func init() {
	SchemeBuilder.Register(&StorageBucketObject{}, &StorageBucketObjectList{})
}
