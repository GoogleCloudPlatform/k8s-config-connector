// Copyright 2024 Google LLC
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
	refs "github.com/GoogleCloudPlatform/k8s-config-connector/apis/refs/v1beta1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/apis/k8s/v1alpha1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

var KMSKeyHandleGVK = SchemeGroupVersion.WithKind("KMSKeyHandle")

// NOTE: json tags are required.  Any new fields you add must have json tags for the fields to be serialized.

// KMSKeyHandleSpec defines the desired state of KMSKeyHandle
// +kcc:spec:proto=google.cloud.kms.v1.KeyHandle
type KMSKeyHandleSpec struct {
	// The KMS Key Handle ID used for resource creation or acquisition.
	// For creation: If specified, this value is used as the key handle ID. If not provided, a UUID will be generated and assigned as the key handle ID.
	// For acquisition: This field must be provided to identify the key handle resource to acquire.
	ResourceID *string `json:"resourceID,omitempty"`

	// Project hosting KMSKeyHandle
	ProjectRef *refs.ProjectRef `json:"projectRef,omitempty"`

	// Location name to create KeyHandle
	Location *string `json:"location,omitempty"`

	// Indicates the resource type that the resulting [CryptoKey][] is meant to
	// protect, e.g. `{SERVICE}.googleapis.com/{TYPE}`. See documentation for
	// supported resource types https://cloud.google.com/kms/docs/autokey-overview#compatible-services.
	ResourceTypeSelector *string `json:"resourceTypeSelector,omitempty"`
}

// KMSKeyHandleStatus defines the config connector machine state of KMSKeyHandle
type KMSKeyHandleStatus struct {
	/* Conditions represent the latest available observations of the
	   object's current state. */
	Conditions []v1alpha1.Condition `json:"conditions,omitempty"`

	// ObservedGeneration is the generation of the resource that was most recently observed by the Config Connector controller. If this is equal to metadata.generation, then that means that the current reported status reflects the most recent desired state of the resource.
	ObservedGeneration *int64 `json:"observedGeneration,omitempty"`

	// A unique specifier for the KMSKeyHandle resource in GCP.
	ExternalRef *string `json:"externalRef,omitempty"`

	// ObservedState is the state of the resource as most recently observed in GCP.
	ObservedState *KMSKeyHandleObservedState `json:"observedState,omitempty"`
}

// KMSKeyHandleObservedState is the state of the KMSKeyHandle resource as most recently observed in GCP.
type KMSKeyHandleObservedState struct {
	KMSKey *string `json:"kmsKey,omitempty"`
}

// +genclient
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:resource:categories=gcp,shortName=gcpkmskeyhandle;gcpkmskeyhandles
// +kubebuilder:subresource:status
// +kubebuilder:metadata:labels="cnrm.cloud.google.com/managed-by-kcc=true";"cnrm.cloud.google.com/system=true";"cnrm.cloud.google.com/stability-level=beta"
// +kubebuilder:printcolumn:name="Age",JSONPath=".metadata.creationTimestamp",type="date"
// +kubebuilder:printcolumn:name="Ready",JSONPath=".status.conditions[?(@.type=='Ready')].status",type="string",description="When 'True', the most recent reconcile of the resource succeeded"
// +kubebuilder:printcolumn:name="Status",JSONPath=".status.conditions[?(@.type=='Ready')].reason",type="string",description="The reason for the value in 'Ready'"
// +kubebuilder:printcolumn:name="Status Age",JSONPath=".status.conditions[?(@.type=='Ready')].lastTransitionTime",type="date",description="The last transition time for the value in 'Status'"

// KMSKeyHandle is the Schema for the KMSKeyHandle API
// +k8s:openapi-gen=true
// +kubebuilder:storageversion
type KMSKeyHandle struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	// +required
	Spec   KMSKeyHandleSpec   `json:"spec,omitempty"`
	Status KMSKeyHandleStatus `json:"status,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// KMSKeyHandleList contains a list of KMSKeyHandle
type KMSKeyHandleList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []KMSKeyHandle `json:"items"`
}

func init() {
	SchemeBuilder.Register(&KMSKeyHandle{}, &KMSKeyHandleList{})
}
