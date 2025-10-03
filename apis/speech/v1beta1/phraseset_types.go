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

package v1beta1

import (
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/apis/k8s/v1alpha1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

var SpeechPhraseSetGVK = GroupVersion.WithKind("SpeechPhraseSet")

// SpeechPhraseSetSpec defines the desired state of SpeechPhraseSet
// +kcc:spec:proto=google.cloud.speech.v2.PhraseSet
type SpeechPhraseSetSpec struct {
	// The SpeechPhraseSet name. If not given, the metadata.name will be used.
	ResourceID *string `json:"resourceID,omitempty"`

	Parent `json:",inline"`

	// A list of word and phrases.
	// +kcc:proto:field=google.cloud.speech.v2.PhraseSet.phrases
	Phrases []PhraseSet_Phrase `json:"phrases,omitempty"`

	// Hint Boost. Positive value will increase the probability that a specific
	//  phrase will be recognized over other similar sounding phrases. The higher
	//  the boost, the higher the chance of false positive recognition as well.
	//  Valid `boost` values are between 0 (exclusive) and 20. We recommend using a
	//  binary search approach to finding the optimal value for your use case as
	//  well as adding phrases both with and without boost to your requests.
	// +kcc:proto:field=google.cloud.speech.v2.PhraseSet.boost
	/* NOTYET Boost *string `json:"boost,omitempty"` */

	// User-settable, human-readable name for the PhraseSet. Must be 63
	//  characters or less.
	// +kcc:proto:field=google.cloud.speech.v2.PhraseSet.display_name
	DisplayName *string `json:"displayName,omitempty"`

	// Allows users to store small amounts of arbitrary data.
	//  Both the key and the value must be 63 characters or less each.
	//  At most 100 annotations.
	// +kcc:proto:field=google.cloud.speech.v2.PhraseSet.annotations
	Annotations map[string]string `json:"annotations,omitempty"`
}

// SpeechPhraseSetStatus defines the config connector machine state of SpeechPhraseSet
type SpeechPhraseSetStatus struct {
	/* Conditions represent the latest available observations of the
	   object's current state. */
	Conditions []v1alpha1.Condition `json:"conditions,omitempty"`

	// ObservedGeneration is the generation of the resource that was most recently observed by the Config Connector controller. If this is equal to metadata.generation, then that means that the current reported status reflects the most recent desired state of the resource.
	ObservedGeneration *int64 `json:"observedGeneration,omitempty"`

	// A unique specifier for the SpeechPhraseSet resource in GCP.
	ExternalRef *string `json:"externalRef,omitempty"`

	// ObservedState is the state of the resource as most recently observed in GCP.
	ObservedState *SpeechPhraseSetObservedState `json:"observedState,omitempty"`
}

// SpeechPhraseSetObservedState is the state of the SpeechPhraseSet resource as most recently observed in GCP.
// +kcc:observedstate:proto=google.cloud.speech.v2.PhraseSet
type SpeechPhraseSetObservedState struct {
	// Output only. Identifier. The resource name of the PhraseSet.
	//  Format: `projects/{project}/locations/{location}/phraseSets/{phrase_set}`.
	// +kcc:proto:field=google.cloud.speech.v2.PhraseSet.name
	// NOTYET: this field serves the same purpose as externalRef
	// Name *string `json:"name,omitempty"`

	// Output only. System-assigned unique identifier for the PhraseSet.
	// +kcc:proto:field=google.cloud.speech.v2.PhraseSet.uid
	UID *string `json:"uid,omitempty"`

	// Output only. The PhraseSet lifecycle state.
	// +kcc:proto:field=google.cloud.speech.v2.PhraseSet.state
	State *string `json:"state,omitempty"`

	// Output only. Creation time.
	// +kcc:proto:field=google.cloud.speech.v2.PhraseSet.create_time
	CreateTime *string `json:"createTime,omitempty"`

	// Output only. The most recent time this resource was modified.
	// +kcc:proto:field=google.cloud.speech.v2.PhraseSet.update_time
	UpdateTime *string `json:"updateTime,omitempty"`

	// Output only. The time at which this resource was requested for deletion.
	// +kcc:proto:field=google.cloud.speech.v2.PhraseSet.delete_time
	DeleteTime *string `json:"deleteTime,omitempty"`

	// Output only. The time at which this resource will be purged.
	// +kcc:proto:field=google.cloud.speech.v2.PhraseSet.expire_time
	ExpireTime *string `json:"expireTime,omitempty"`

	// Output only. This checksum is computed by the server based on the value of
	//  other fields. This may be sent on update, undelete, and delete requests to
	//  ensure the client has an up-to-date value before proceeding.
	// +kcc:proto:field=google.cloud.speech.v2.PhraseSet.etag
	Etag *string `json:"etag,omitempty"`

	// Output only. Whether or not this PhraseSet is in the process of being
	//  updated.
	// +kcc:proto:field=google.cloud.speech.v2.PhraseSet.reconciling
	Reconciling *bool `json:"reconciling,omitempty"`

	// Output only. The [KMS key
	//  name](https://cloud.google.com/kms/docs/resource-hierarchy#keys) with which
	//  the PhraseSet is encrypted. The expected format is
	//  `projects/{project}/locations/{location}/keyRings/{key_ring}/cryptoKeys/{crypto_key}`.
	// +kcc:proto:field=google.cloud.speech.v2.PhraseSet.kms_key_name
	KMSKeyName *string `json:"kmsKeyName,omitempty"`

	// Output only. The [KMS key version
	//  name](https://cloud.google.com/kms/docs/resource-hierarchy#key_versions)
	//  with which the PhraseSet is encrypted. The expected format is
	//  `projects/{project}/locations/{location}/keyRings/{key_ring}/cryptoKeys/{crypto_key}/cryptoKeyVersions/{crypto_key_version}`.
	// +kcc:proto:field=google.cloud.speech.v2.PhraseSet.kms_key_version_name
	KMSKeyVersionName *string `json:"kmsKeyVersionName,omitempty"`
}

// +genclient
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:resource:categories=gcp,shortName=gcpspeechphraseset;gcpspeechphrasesets
// +kubebuilder:subresource:status
// +kubebuilder:metadata:labels="cnrm.cloud.google.com/managed-by-kcc=true";"cnrm.cloud.google.com/system=true"
// +kubebuilder:printcolumn:name="Age",JSONPath=".metadata.creationTimestamp",type="date"
// +kubebuilder:printcolumn:name="Ready",JSONPath=".status.conditions[?(@.type=='Ready')].status",type="string",description="When 'True', the most recent reconcile of the resource succeeded"
// +kubebuilder:printcolumn:name="Status",JSONPath=".status.conditions[?(@.type=='Ready')].reason",type="string",description="The reason for the value in 'Ready'"
// +kubebuilder:printcolumn:name="Status Age",JSONPath=".status.conditions[?(@.type=='Ready')].lastTransitionTime",type="date",description="The last transition time for the value in 'Status'"

// SpeechPhraseSet is the Schema for the SpeechPhraseSet API
// +k8s:openapi-gen=true
// +kubebuilder:storageversion
type SpeechPhraseSet struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	// +required
	Spec   SpeechPhraseSetSpec   `json:"spec,omitempty"`
	Status SpeechPhraseSetStatus `json:"status,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// SpeechPhraseSetList contains a list of SpeechPhraseSet
type SpeechPhraseSetList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []SpeechPhraseSet `json:"items"`
}

func init() {
	SchemeBuilder.Register(&SpeechPhraseSet{}, &SpeechPhraseSetList{})
}

// +kcc:proto=google.cloud.speech.v2.PhraseSet.Phrase
type PhraseSet_Phrase struct {
	// The phrase itself.
	// +kcc:proto:field=google.cloud.speech.v2.PhraseSet.Phrase.value
	Value *string `json:"value,omitempty"`

	// Hint Boost. Overrides the boost set at the phrase set level.
	//  Positive value will increase the probability that a specific phrase will
	//  be recognized over other similar sounding phrases. The higher the boost,
	//  the higher the chance of false positive recognition as well. Negative
	//  boost values would correspond to anti-biasing. Anti-biasing is not
	//  enabled, so negative boost values will return an error. Boost values must
	//  be between 0 and 20. Any values outside that range will return an error.
	//  We recommend using a binary search approach to finding the optimal value
	//  for your use case as well as adding phrases both with and without boost
	//  to your requests.
	// +kcc:proto:field=google.cloud.speech.v2.PhraseSet.Phrase.boost
	Boost *string `json:"boost,omitempty"`
}
