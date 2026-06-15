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
	refsv1beta1 "github.com/GoogleCloudPlatform/k8s-config-connector/apis/refs/v1beta1"
	storagev1beta1 "github.com/GoogleCloudPlatform/k8s-config-connector/apis/storage/v1beta1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/apis/k8s/v1alpha1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

var StorageBatchOperationsJobGVK = GroupVersion.WithKind("StorageBatchOperationsJob")

// StorageBatchOperationsJobSpec defines the desired state of StorageBatchOperationsJob
// +kcc:spec:proto=google.cloud.storagebatchoperations.v1.Job
type StorageBatchOperationsJobSpec struct {
	// The project that this resource belongs to.
	// +required
	ProjectRef *refsv1beta1.ProjectRef `json:"projectRef"`

	// The location of this resource.
	// +required
	Location string `json:"location"`

	// The StorageBatchOperationsJob name. If not given, the metadata.name will be used.
	// +optional
	ResourceID *string `json:"resourceID,omitempty"`

	// Optional. A description provided by the user for the job. Its max length is
	//  1024 bytes when Unicode-encoded.
	// +optional
	Description *string `json:"description,omitempty"`

	// Specifies a list of buckets and their objects to be transformed.
	// +optional
	BucketList *BucketList `json:"bucketList,omitempty"`

	// Changes object hold status.
	// +optional
	PutObjectHold *PutObjectHold `json:"putObjectHold,omitempty"`

	// Delete objects.
	// +optional
	DeleteObject *DeleteObject `json:"deleteObject,omitempty"`

	// Updates object metadata. Allows updating fixed-key and custom metadata
	//  and fixed-key metadata i.e. Cache-Control, Content-Disposition,
	//  Content-Encoding, Content-Language, Content-Type, Custom-Time.
	// +optional
	PutMetadata *PutMetadata `json:"putMetadata,omitempty"`

	// Rewrite the object and updates metadata like KMS key.
	// +optional
	RewriteObject *RewriteObject `json:"rewriteObject,omitempty"`

	// Optional. Logging configuration.
	// +optional
	LoggingConfig *LoggingConfig `json:"loggingConfig,omitempty"`
}

// +kcc:proto=google.cloud.storagebatchoperations.v1.BucketList
type BucketList struct {
	// Required. List of buckets and their objects to be transformed. Currently,
	//  only one bucket configuration is supported. If multiple buckets are
	//  specified, an error will be returned.
	// +required
	Buckets []BucketList_Bucket `json:"buckets,omitempty"`
}

// +kcc:proto=google.cloud.storagebatchoperations.v1.BucketList.Bucket
type BucketList_Bucket struct {
	// Required. Bucket name for the objects to be transformed.
	// +required
	BucketRef *storagev1beta1.StorageBucketRef `json:"bucketRef,omitempty"`

	// Specifies objects matching a prefix set.
	// +optional
	PrefixList *PrefixList `json:"prefixList,omitempty"`

	// Specifies objects in a manifest file.
	// +optional
	Manifest *Manifest `json:"manifest,omitempty"`
}

// +kcc:proto=google.cloud.storagebatchoperations.v1.RewriteObject
type RewriteObject struct {
	// Required. Resource name of the Cloud KMS key that will be used to encrypt
	//  the object. The Cloud KMS key must be located in same location as the
	//  object. Refer to
	//  https://cloud.google.com/storage/docs/encryption/using-customer-managed-keys#add-object-key
	//  for additional documentation. Format:
	//  projects/{project}/locations/{location}/keyRings/{keyring}/cryptoKeys/{key}
	//  For example:
	//  "projects/123456/locations/us-central1/keyRings/my-keyring/cryptoKeys/my-key".
	//  The object will be rewritten and set with the specified KMS key.
	// +required
	KMSCryptoKeyRef *refsv1beta1.KMSCryptoKeyRef `json:"kmsCryptoKeyRef,omitempty"`
}

// StorageBatchOperationsJobStatus defines the config connector machine state of StorageBatchOperationsJob
type StorageBatchOperationsJobStatus struct {
	/* Conditions represent the latest available observations of the
	   object's current state. */
	Conditions []v1alpha1.Condition `json:"conditions,omitempty"`

	// ObservedGeneration is the generation of the resource that was most recently observed by the Config Connector controller. If this is equal to metadata.generation, then that means that the current reported status reflects the most recent desired state of the resource.
	ObservedGeneration *int64 `json:"observedGeneration,omitempty"`

	// A unique specifier for the StorageBatchOperationsJob resource in GCP.
	ExternalRef *string `json:"externalRef,omitempty"`

	// ObservedState is the state of the resource as most recently observed in GCP.
	ObservedState *StorageBatchOperationsJobObservedState `json:"observedState,omitempty"`
}

// StorageBatchOperationsJobObservedState is the state of the StorageBatchOperationsJob resource as most recently observed in GCP.
// +kcc:observedstate:proto=google.cloud.storagebatchoperations.v1.Job
type StorageBatchOperationsJobObservedState struct {
	// Output only. The time that the job was created.
	CreateTime *string `json:"createTime,omitempty"`

	// Output only. The time that the job was scheduled.
	ScheduleTime *string `json:"scheduleTime,omitempty"`

	// Output only. The time that the job was completed.
	CompleteTime *string `json:"completeTime,omitempty"`

	// Output only. Information about the progress of the job.
	Counters *CountersObservedState `json:"counters,omitempty"`

	// Output only. Summarizes errors encountered with sample error log entries.
	ErrorSummaries []ErrorSummaryObservedState `json:"errorSummaries,omitempty"`

	// Output only. State of the job.
	State *string `json:"state,omitempty"`
}

// +genclient
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:resource:categories=gcp,shortName=gcpstoragebatchoperationsjob;gcpstoragebatchoperationsjobs
// +kubebuilder:subresource:status
// +kubebuilder:metadata:labels="cnrm.cloud.google.com/managed-by-kcc=true"
// +kubebuilder:metadata:labels="cnrm.cloud.google.com/system=true"
// +kubebuilder:metadata:labels="cnrm.cloud.google.com/stability-level=alpha"
// +kubebuilder:printcolumn:name="Age",JSONPath=".metadata.creationTimestamp",type="date"
// +kubebuilder:printcolumn:name="Ready",JSONPath=".status.conditions[?(@.type=='Ready')].status",type="string",description="When 'True', the most recent reconcile of the resource succeeded"
// +kubebuilder:printcolumn:name="Status",JSONPath=".status.conditions[?(@.type=='Ready')].reason",type="string",description="The reason for the value in 'Ready'"
// +kubebuilder:printcolumn:name="Status Age",JSONPath=".status.conditions[?(@.type=='Ready')].lastTransitionTime",type="date",description="The last transition time for the value in 'Status'"

// StorageBatchOperationsJob is the Schema for the StorageBatchOperationsJob API
// +k8s:openapi-gen=true
type StorageBatchOperationsJob struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	// +required
	Spec   StorageBatchOperationsJobSpec   `json:"spec,omitempty"`
	Status StorageBatchOperationsJobStatus `json:"status,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// StorageBatchOperationsJobList contains a list of StorageBatchOperationsJob
type StorageBatchOperationsJobList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []StorageBatchOperationsJob `json:"items"`
}

func init() {
	SchemeBuilder.Register(&StorageBatchOperationsJob{}, &StorageBatchOperationsJobList{})
}
