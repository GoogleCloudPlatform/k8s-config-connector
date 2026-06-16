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
	refsv1beta1 "github.com/GoogleCloudPlatform/k8s-config-connector/apis/refs/v1beta1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/apis/k8s/v1alpha1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

var APIKeysKeyGVK = GroupVersion.WithKind("APIKeysKey")

// APIKeysKeySpec defines the desired state of APIKeysKey
// +kcc:spec:proto=google.api.apikeys.v2.Key
type APIKeysKeySpec struct {
	/* Immutable. The Project that this resource belongs to. */
	ProjectRef refsv1beta1.ProjectRef `json:"projectRef"`

	/* Immutable. Optional. The name of the resource. Used for creation and acquisition. When unset, the value of `metadata.name` is used as the default. */
	// +optional
	ResourceID *string `json:"resourceID,omitempty"`

	// Human-readable display name of this key that you can modify.
	//  The maximum length is 63 characters.
	// +kcc:proto:field=google.api.apikeys.v2.Key.display_name
	DisplayName *string `json:"displayName,omitempty"`

	// Key restrictions.
	// +kcc:proto:field=google.api.apikeys.v2.Key.restrictions
	Restrictions *Restrictions `json:"restrictions,omitempty"`
}

// APIKeysKeyStatus defines the config connector machine state of APIKeysKey
type APIKeysKeyStatus struct {
	/* Conditions represent the latest available observations of the
	   object's current state. */
	Conditions []v1alpha1.Condition `json:"conditions,omitempty"`

	// ObservedGeneration is the generation of the resource that was most recently observed by the Config Connector controller. If this is equal to metadata.generation, then that means that the current reported status reflects the most recent desired state of the resource.
	// +optional
	ObservedGeneration *int64 `json:"observedGeneration,omitempty"`

	// A unique specifier for the APIKeysKey resource in Google Cloud.
	// +optional
	ExternalRef *string `json:"externalRef,omitempty"`

	// ObservedState is the state of the resource as most recently observed in Google Cloud.
	// +optional
	ObservedState *APIKeysKeyObservedState `json:"observedState,omitempty"`
}

// APIKeysKeyObservedState is the state of the APIKeysKey resource as most recently observed in Google Cloud.
// +kcc:observedstate:proto=google.api.apikeys.v2.Key
type APIKeysKeyObservedState struct {
	// Output only. Unique id in UUID4 format.
	// +kcc:proto:field=google.api.apikeys.v2.Key.uid
	Uid *string `json:"uid,omitempty"`
}

// +genclient
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:resource:categories=gcp,shortName=gcpapikeyskey;gcpapikeyskeys
// +kubebuilder:subresource:status
// +kubebuilder:metadata:labels="cnrm.cloud.google.com/tf2crd=true"
// +kubebuilder:metadata:labels="cnrm.cloud.google.com/managed-by-kcc=true"
// +kubebuilder:metadata:labels="cnrm.cloud.google.com/stability-level=alpha"
// +kubebuilder:metadata:labels="cnrm.cloud.google.com/system=true"
// +kubebuilder:printcolumn:name="Age",JSONPath=".metadata.creationTimestamp",type="date"
// +kubebuilder:printcolumn:name="Ready",JSONPath=".status.conditions[?(@.type=='Ready')].status",type="string",description="When 'True', the most recent reconcile of the resource succeeded"
// +kubebuilder:printcolumn:name="Status",JSONPath=".status.conditions[?(@.type=='Ready')].reason",type="string",description="The reason for the value in 'Ready'"
// +kubebuilder:printcolumn:name="Status Age",JSONPath=".status.conditions[?(@.type=='Ready')].lastTransitionTime",type="date",description="The last transition time for the value in 'Status'"

// APIKeysKey is the Schema for the APIKeysKey API
// +k8s:openapi-gen=true
type APIKeysKey struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	// +required
	Spec   APIKeysKeySpec   `json:"spec,omitempty"`
	Status APIKeysKeyStatus `json:"status,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// APIKeysKeyList contains a list of APIKeysKey
type APIKeysKeyList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []APIKeysKey `json:"items"`
}

func init() {
	SchemeBuilder.Register(&APIKeysKey{}, &APIKeysKeyList{})
}
