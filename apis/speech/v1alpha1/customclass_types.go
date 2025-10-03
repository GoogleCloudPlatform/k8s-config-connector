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
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/apis/k8s/v1alpha1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

var SpeechCustomClassGVK = GroupVersion.WithKind("SpeechCustomClass")

// SpeechCustomClassSpec defines the desired state of SpeechCustomClass
// +kcc:spec:proto=google.cloud.speech.v2.CustomClass
type SpeechCustomClassSpec struct {
	// The SpeechCustomClass name. If not given, the metadata.name will be used.
	ResourceID *string `json:"resourceID,omitempty"`

	Parent `json:",inline"`

	// Optional. User-settable, human-readable name for the CustomClass. Must be
	//  63 characters or less.
	// +kcc:proto:field=google.cloud.speech.v2.CustomClass.display_name
	DisplayName *string `json:"displayName,omitempty"`

	// A collection of class items.
	// +kcc:proto:field=google.cloud.speech.v2.CustomClass.items
	Items []CustomClass_ClassItem `json:"items,omitempty"`

	// Optional. Allows users to store small amounts of arbitrary data.
	//  Both the key and the value must be 63 characters or less each.
	//  At most 100 annotations.
	// +kcc:proto:field=google.cloud.speech.v2.CustomClass.annotations
	Annotations map[string]string `json:"annotations,omitempty"`
}

// SpeechCustomClassStatus defines the config connector machine state of SpeechCustomClass
type SpeechCustomClassStatus struct {
	/* Conditions represent the latest available observations of the
	   object's current state. */
	Conditions []v1alpha1.Condition `json:"conditions,omitempty"`

	// ObservedGeneration is the generation of the resource that was most recently observed by the Config Connector controller. If this is equal to metadata.generation, then that means that the current reported status reflects the most recent desired state of the resource.
	ObservedGeneration *int64 `json:"observedGeneration,omitempty"`

	// A unique specifier for the SpeechCustomClass resource in GCP.
	ExternalRef *string `json:"externalRef,omitempty"`

	// ObservedState is the state of the resource as most recently observed in GCP.
	ObservedState *SpeechCustomClassObservedState `json:"observedState,omitempty"`
}

// SpeechCustomClassObservedState is the state of the SpeechCustomClass resource as most recently observed in GCP.
// +kcc:observedstate:proto=google.cloud.speech.v2.CustomClass
type SpeechCustomClassObservedState struct {
	// Output only. Identifier. The resource name of the CustomClass.
	//  Format:
	//  `projects/{project}/locations/{location}/customClasses/{custom_class}`.
	// +kcc:proto:field=google.cloud.speech.v2.CustomClass.name
	// NOTYET: this field serves the same purpose as externalRef
	// Name *string `json:"name,omitempty"`

	// Output only. System-assigned unique identifier for the CustomClass.
	// +kcc:proto:field=google.cloud.speech.v2.CustomClass.uid
	UID *string `json:"uid,omitempty"`

	// Output only. The CustomClass lifecycle state.
	// +kcc:proto:field=google.cloud.speech.v2.CustomClass.state
	State *string `json:"state,omitempty"`

	// Output only. Creation time.
	// +kcc:proto:field=google.cloud.speech.v2.CustomClass.create_time
	CreateTime *string `json:"createTime,omitempty"`

	// Output only. The most recent time this resource was modified.
	// +kcc:proto:field=google.cloud.speech.v2.CustomClass.update_time
	UpdateTime *string `json:"updateTime,omitempty"`

	// Output only. The time at which this resource was requested for deletion.
	// +kcc:proto:field=google.cloud.speech.v2.CustomClass.delete_time
	DeleteTime *string `json:"deleteTime,omitempty"`

	// Output only. The time at which this resource will be purged.
	// +kcc:proto:field=google.cloud.speech.v2.CustomClass.expire_time
	ExpireTime *string `json:"expireTime,omitempty"`

	// Output only. This checksum is computed by the server based on the value of
	//  other fields. This may be sent on update, undelete, and delete requests to
	//  ensure the client has an up-to-date value before proceeding.
	// +kcc:proto:field=google.cloud.speech.v2.CustomClass.etag
	Etag *string `json:"etag,omitempty"`

	// Output only. Whether or not this CustomClass is in the process of being
	//  updated.
	// +kcc:proto:field=google.cloud.speech.v2.CustomClass.reconciling
	Reconciling *bool `json:"reconciling,omitempty"`

	// Output only. The [KMS key
	//  name](https://cloud.google.com/kms/docs/resource-hierarchy#keys) with which
	//  the CustomClass is encrypted. The expected format is
	//  `projects/{project}/locations/{location}/keyRings/{key_ring}/cryptoKeys/{crypto_key}`.
	// +kcc:proto:field=google.cloud.speech.v2.CustomClass.kms_key_name
	KMSKeyName *string `json:"kmsKeyName,omitempty"`

	// Output only. The [KMS key version
	//  name](https://cloud.google.com/kms/docs/resource-hierarchy#key_versions)
	//  with which the CustomClass is encrypted. The expected format is
	//  `projects/{project}/locations/{location}/keyRings/{key_ring}/cryptoKeys/{crypto_key}/cryptoKeyVersions/{crypto_key_version}`.
	// +kcc:proto:field=google.cloud.speech.v2.CustomClass.kms_key_version_name
	KMSKeyVersionName *string `json:"kmsKeyVersionName,omitempty"`
}

// +genclient
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:resource:categories=gcp,shortName=gcpspeechcustomclass;gcpspeechcustomclasses
// +kubebuilder:subresource:status
// +kubebuilder:metadata:labels="cnrm.cloud.google.com/managed-by-kcc=true";"cnrm.cloud.google.com/system=true"
// +kubebuilder:printcolumn:name="Age",JSONPath=".metadata.creationTimestamp",type="date"
// +kubebuilder:printcolumn:name="Ready",JSONPath=".status.conditions[?(@.type=='Ready')].status",type="string",description="When 'True', the most recent reconcile of the resource succeeded"
// +kubebuilder:printcolumn:name="Status",JSONPath=".status.conditions[?(@.type=='Ready')].reason",type="string",description="The reason for the value in 'Ready'"
// +kubebuilder:printcolumn:name="Status Age",JSONPath=".status.conditions[?(@.type=='Ready')].lastTransitionTime",type="date",description="The last transition time for the value in 'Status'"

// SpeechCustomClass is the Schema for the SpeechCustomClass API
// +k8s:openapi-gen=true
type SpeechCustomClass struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	// +required
	Spec   SpeechCustomClassSpec   `json:"spec,omitempty"`
	Status SpeechCustomClassStatus `json:"status,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// SpeechCustomClassList contains a list of SpeechCustomClass
type SpeechCustomClassList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []SpeechCustomClass `json:"items"`
}

func init() {
	SchemeBuilder.Register(&SpeechCustomClass{}, &SpeechCustomClassList{})
}
