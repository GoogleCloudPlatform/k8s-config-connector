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

var NotebooksRuntimeGVK = GroupVersion.WithKind("NotebooksRuntime")

// NotebooksRuntimeSpec defines the desired state of NotebooksRuntime
// +kcc:proto=google.cloud.notebooks.v1.Runtime
type NotebooksRuntimeSpec struct {
	// The NotebooksRuntime name. If not given, the metadata.name will be used.
	ResourceID *string `json:"resourceID,omitempty"`
}

// NotebooksRuntimeStatus defines the config connector machine state of NotebooksRuntime
type NotebooksRuntimeStatus struct {
	/* Conditions represent the latest available observations of the
	   object's current state. */
	Conditions []v1alpha1.Condition `json:"conditions,omitempty"`

	// ObservedGeneration is the generation of the resource that was most recently observed by the Config Connector controller. If this is equal to metadata.generation, then that means that the current reported status reflects the most recent desired state of the resource.
	ObservedGeneration *int64 `json:"observedGeneration,omitempty"`

	// A unique specifier for the NotebooksRuntime resource in GCP.
	ExternalRef *string `json:"externalRef,omitempty"`

	// ObservedState is the state of the resource as most recently observed in GCP.
	ObservedState *NotebooksRuntimeObservedState `json:"observedState,omitempty"`
}

// NotebooksRuntimeObservedState is the state of the NotebooksRuntime resource as most recently observed in GCP.
// +kcc:proto=google.cloud.notebooks.v1.Runtime
type NotebooksRuntimeObservedState struct {
}

// +genclient
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// TODO(user): make sure the pluralizaiton below is correct
// +kubebuilder:resource:categories=gcp,shortName=gcpnotebooksruntime;gcpnotebooksruntimes
// +kubebuilder:subresource:status
// +kubebuilder:metadata:labels="cnrm.cloud.google.com/managed-by-kcc=true";"cnrm.cloud.google.com/system=true"
// +kubebuilder:printcolumn:name="Age",JSONPath=".metadata.creationTimestamp",type="date"
// +kubebuilder:printcolumn:name="Ready",JSONPath=".status.conditions[?(@.type=='Ready')].status",type="string",description="When 'True', the most recent reconcile of the resource succeeded"
// +kubebuilder:printcolumn:name="Status",JSONPath=".status.conditions[?(@.type=='Ready')].reason",type="string",description="The reason for the value in 'Ready'"
// +kubebuilder:printcolumn:name="Status Age",JSONPath=".status.conditions[?(@.type=='Ready')].lastTransitionTime",type="date",description="The last transition time for the value in 'Status'"

// NotebooksRuntime is the Schema for the NotebooksRuntime API
// +k8s:openapi-gen=true
type NotebooksRuntime struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	// +required
	Spec   NotebooksRuntimeSpec   `json:"spec,omitempty"`
	Status NotebooksRuntimeStatus `json:"status,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// NotebooksRuntimeList contains a list of NotebooksRuntime
type NotebooksRuntimeList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []NotebooksRuntime `json:"items"`
}

func init() {
	SchemeBuilder.Register(&NotebooksRuntime{}, &NotebooksRuntimeList{})
}
