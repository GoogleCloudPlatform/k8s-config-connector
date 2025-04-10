// Copyright 2025 Google LLC
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
	refs "github.com/GoogleCloudPlatform/k8s-config-connector/apis/refs/v1beta1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/apis/k8s/v1alpha1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

var KMSImportJobGVK = GroupVersion.WithKind("KMSImportJob")

// KMSImportJobSpec defines the desired state of KMSImportJob
// +kcc:spec:proto=google.cloud.kms.v1.ImportJob
type KMSImportJobSpec struct {
	KMSKeyRingRef *refs.KMSKeyRingRef `json:"kmsKeyRingRef"`

	// The KMSImportJob name. If not given, the metadata.name will be used.
	ResourceID *string `json:"resourceID,omitempty"`

	// Required. Immutable. The wrapping method to be used for incoming key
	// material.
	// +kcc:proto:field=google.cloud.kms.v1.ImportJob.import_method
	ImportMethod *string `json:"importMethod"`

	// Required. Immutable. The protection level of the
	// [ImportJob][google.cloud.kms.v1.ImportJob]. This must match the
	// [protection_level][google.cloud.kms.v1.CryptoKeyVersionTemplate.protection_level]
	// of the [version_template][google.cloud.kms.v1.CryptoKey.version_template]
	// on the [CryptoKey][google.cloud.kms.v1.CryptoKey] you attempt to import
	// into.
	// +kcc:proto:field=google.cloud.kms.v1.ImportJob.protection_level
	ProtectionLevel *string `json:"protectionLevel"`
}

// KMSImportJobStatus defines the config connector machine state of KMSImportJob
type KMSImportJobStatus struct {
	/* Conditions represent the latest available observations of the
	   object's current state. */
	Conditions []v1alpha1.Condition `json:"conditions,omitempty"`

	// ObservedGeneration is the generation of the resource that was most recently observed by the Config Connector controller. If this is equal to metadata.generation, then that means that the current reported status reflects the most recent desired state of the resource.
	ObservedGeneration *int64 `json:"observedGeneration,omitempty"`

	// A unique specifier for the KMSImportJob resource in GCP.
	ExternalRef *string `json:"externalRef,omitempty"`

	// ObservedState is the state of the resource as most recently observed in GCP.
	ObservedState *KMSImportJobObservedState `json:"observedState,omitempty"`
}

// KMSImportJobObservedState is the state of the KMSImportJob resource as most recently observed in GCP.
// +kcc:observedstate:proto=google.cloud.kms.v1.ImportJob
type KMSImportJobObservedState struct {

	// Output only. The time at which this
	//  [ImportJob][google.cloud.kms.v1.ImportJob] was created.
	// +kcc:proto:field=google.cloud.kms.v1.ImportJob.create_time
	CreateTime *string `json:"createTime,omitempty"`

	// Output only. The time this [ImportJob][google.cloud.kms.v1.ImportJob]'s key
	//  material was generated.
	// +kcc:proto:field=google.cloud.kms.v1.ImportJob.generate_time
	GenerateTime *string `json:"generateTime,omitempty"`

	// Output only. The time at which this
	//  [ImportJob][google.cloud.kms.v1.ImportJob] is scheduled for expiration and
	//  can no longer be used to import key material.
	// +kcc:proto:field=google.cloud.kms.v1.ImportJob.expire_time
	ExpireTime *string `json:"expireTime,omitempty"`

	// Output only. The time this [ImportJob][google.cloud.kms.v1.ImportJob]
	//  expired. Only present if [state][google.cloud.kms.v1.ImportJob.state] is
	//  [EXPIRED][google.cloud.kms.v1.ImportJob.ImportJobState.EXPIRED].
	// +kcc:proto:field=google.cloud.kms.v1.ImportJob.expire_event_time
	ExpireEventTime *string `json:"expireEventTime,omitempty"`

	// Output only. The current state of the
	//  [ImportJob][google.cloud.kms.v1.ImportJob], indicating if it can be used.
	// +kcc:proto:field=google.cloud.kms.v1.ImportJob.state
	State *string `json:"state,omitempty"`

	// Output only. The public key with which to wrap key material prior to
	//  import. Only returned if [state][google.cloud.kms.v1.ImportJob.state] is
	//  [ACTIVE][google.cloud.kms.v1.ImportJob.ImportJobState.ACTIVE].
	// +kcc:proto:field=google.cloud.kms.v1.ImportJob.public_key
	PublicKey *ImportJob_WrappingPublicKeyObservedState `json:"publicKey,omitempty"`

	// Output only. Statement that was generated and signed by the key creator
	//  (for example, an HSM) at key creation time. Use this statement to verify
	//  attributes of the key as stored on the HSM, independently of Google.
	//  Only present if the chosen
	//  [ImportMethod][google.cloud.kms.v1.ImportJob.ImportMethod] is one with a
	//  protection level of [HSM][google.cloud.kms.v1.ProtectionLevel.HSM].
	// +kcc:proto:field=google.cloud.kms.v1.ImportJob.attestation
	Attestation *KeyOperationAttestationObservedState `json:"attestation,omitempty"`
}

// +kcc:observedstate:proto=google.cloud.kms.v1.ImportJob.WrappingPublicKey
type ImportJob_WrappingPublicKeyObservedState struct {
	// The public key, encoded in PEM format. For more information, see the [RFC
	//  7468](https://tools.ietf.org/html/rfc7468) sections for [General
	//  Considerations](https://tools.ietf.org/html/rfc7468#section-2) and
	//  [Textual Encoding of Subject Public Key Info]
	//  (https://tools.ietf.org/html/rfc7468#section-13).
	// +kcc:proto:field=google.cloud.kms.v1.ImportJob.WrappingPublicKey.pem
	Pem *string `json:"pem,omitempty"`
}

// +genclient
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:resource:categories=gcp,shortName=gcpkmsimportjob;gcpkmsimportjobs
// +kubebuilder:subresource:status
// +kubebuilder:metadata:labels="cnrm.cloud.google.com/managed-by-kcc=true";"cnrm.cloud.google.com/system=true"
// +kubebuilder:printcolumn:name="Age",JSONPath=".metadata.creationTimestamp",type="date"
// +kubebuilder:printcolumn:name="Ready",JSONPath=".status.conditions[?(@.type=='Ready')].status",type="string",description="When 'True', the most recent reconcile of the resource succeeded"
// +kubebuilder:printcolumn:name="Status",JSONPath=".status.conditions[?(@.type=='Ready')].reason",type="string",description="The reason for the value in 'Ready'"
// +kubebuilder:printcolumn:name="Status Age",JSONPath=".status.conditions[?(@.type=='Ready')].lastTransitionTime",type="date",description="The last transition time for the value in 'Status'"

// KMSImportJob is the Schema for the KMSImportJob API
// +k8s:openapi-gen=true
type KMSImportJob struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	// +required
	Spec   KMSImportJobSpec   `json:"spec,omitempty"`
	Status KMSImportJobStatus `json:"status,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// KMSImportJobList contains a list of KMSImportJob
type KMSImportJobList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []KMSImportJob `json:"items"`
}

func init() {
	SchemeBuilder.Register(&KMSImportJob{}, &KMSImportJobList{})
}
