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
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/apis/k8s/v1alpha1"

	refs "github.com/GoogleCloudPlatform/k8s-config-connector/apis/refs/v1beta1"

	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

var KMSAutokeyConfigGVK = SchemeGroupVersion.WithKind("KMSAutokeyConfig")

// NOTE: json tags are required.  Any new fields you add must have json tags for the fields to be serialized.

// KMSAutokeyConfigSpec defines the desired state of KMSAutokeyConfig
// +kcc:spec:proto=google.cloud.kms.v1.AutokeyConfig
type KMSAutokeyConfigSpec struct {

	// NOTE: ResourceID field is not required for AutokeyConfig as its ID has the format folders/<folderID>/autokeyConfig i.e., it doesnt have any unique ID of its own and relies on folderID for uniqueness.

	// Immutable. The folder that this resource belongs to.
	// +required
	FolderRef *refs.FolderRef `json:"folderRef"`

	// +optional
	KeyProjectRef *refs.ProjectRef `json:"keyProject,omitempty"`
}

// KMSAutokeyConfigStatus defines the config connector machine state of KMSAutokeyConfig
type KMSAutokeyConfigStatus struct {
	/* Conditions represent the latest available observations of the
	   object's current state. */
	Conditions []v1alpha1.Condition `json:"conditions,omitempty"`

	// ObservedGeneration is the generation of the resource that was most recently observed by the Config Connector controller. If this is equal to metadata.generation, then that means that the current reported status reflects the most recent desired state of the resource.
	ObservedGeneration *int64 `json:"observedGeneration,omitempty"`

	// A unique specifier for the KMSAutokeyConfig resource in GCP.
	ExternalRef *string `json:"externalRef,omitempty"`

	// ObservedState is the state of the resource as most recently observed in GCP.
	ObservedState *KMSAutokeyConfigObservedState `json:"observedState,omitempty"`
}

// KMSAutokeyConfigSpec defines the desired state of KMSAutokeyConfig
// +kcc:observedstate:proto=google.cloud.kms.v1.AutokeyConfig
type KMSAutokeyConfigObservedState struct {
	// Output only. Current state of this AutokeyConfig.
	// +optional
	State *string `json:"state,omitempty"`
}

// +genclient
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:resource:categories=gcp,shortName=gcpkmsautokeyconfig;gcpkmsautokeyconfigs
// +kubebuilder:subresource:status
// +kubebuilder:metadata:labels="cnrm.cloud.google.com/managed-by-kcc=true";"cnrm.cloud.google.com/system=true";"cnrm.cloud.google.com/stability-level=beta"
// +kubebuilder:printcolumn:name="Age",JSONPath=".metadata.creationTimestamp",type="date"
// +kubebuilder:printcolumn:name="Ready",JSONPath=".status.conditions[?(@.type=='Ready')].status",type="string",description="When 'True', the most recent reconcile of the resource succeeded"
// +kubebuilder:printcolumn:name="Status",JSONPath=".status.conditions[?(@.type=='Ready')].reason",type="string",description="The reason for the value in 'Ready'"
// +kubebuilder:printcolumn:name="Status Age",JSONPath=".status.conditions[?(@.type=='Ready')].lastTransitionTime",type="date",description="The last transition time for the value in 'Status'"

// KMSAutokeyConfig is the Schema for the KMSAutokeyConfig API
// +k8s:openapi-gen=true
// +kubebuilder:storageversion
type KMSAutokeyConfig struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   KMSAutokeyConfigSpec   `json:"spec,omitempty"`
	Status KMSAutokeyConfigStatus `json:"status,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// KMSAutokeyConfigList contains a list of KMSAutokeyConfig
type KMSAutokeyConfigList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []KMSAutokeyConfig `json:"items"`
}

func init() {
	SchemeBuilder.Register(&KMSAutokeyConfig{}, &KMSAutokeyConfigList{})
}
