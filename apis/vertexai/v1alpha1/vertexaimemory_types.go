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
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/apis/k8s/v1alpha1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

var VertexAIMemoryGVK = GroupVersion.WithKind("VertexAIMemory")

// VertexAIMemorySpec defines the desired state of VertexAIMemory
type VertexAIMemorySpec struct {
	// The project that this resource belongs to.
	// +required
	ProjectRef *refsv1beta1.ProjectRef `json:"projectRef"`

	// The location of this resource.
	// +required
	Location *string `json:"location"`

	// The parent ReasoningEngine of this resource.
	// +required
	ReasoningEngineRef *VertexAIReasoningEngineRef `json:"reasoningEngineRef"`

	// The VertexAIMemory name. If not given, the metadata.name will be used.
	ResourceID *string `json:"resourceID,omitempty"`

	// Optional. Timestamp of when this resource is considered expired.
	//  This is *always* provided on output, regardless of what `expiration` was
	//  sent on input.
	ExpireTime *string `json:"expireTime,omitempty"`

	// Optional. Input only. The TTL for this resource. The expiration time is
	//  computed: now + TTL.
	TTL *string `json:"ttl,omitempty"`

	// Optional. Display name of the Memory.
	DisplayName *string `json:"displayName,omitempty"`

	// Optional. Description of the Memory.
	Description *string `json:"description,omitempty"`

	// Required. Semantic knowledge extracted from the source content.
	Fact *string `json:"fact,omitempty"`

	// Required. Immutable. The scope of the Memory. Memories are isolated
	//  within their scope. The scope is defined when creating or generating
	//  memories. Scope values cannot contain the wildcard character '*'.
	Scope map[string]string `json:"scope,omitempty"`
}

// VertexAIMemoryStatus defines the config connector machine state of VertexAIMemory
type VertexAIMemoryStatus struct {
	/* Conditions represent the latest available observations of the
	   object's current state. */
	Conditions []v1alpha1.Condition `json:"conditions,omitempty"`

	// ObservedGeneration is the generation of the resource that was most recently observed by the Config Connector controller. If this is equal to metadata.generation, then that means that the current reported status reflects the most recent desired state of the resource.
	ObservedGeneration *int64 `json:"observedGeneration,omitempty"`

	// A unique specifier for the VertexAIMemory resource in GCP.
	ExternalRef *string `json:"externalRef,omitempty"`

	// ObservedState is the state of the resource as most recently observed in GCP.
	ObservedState *VertexAIMemoryObservedState `json:"observedState,omitempty"`
}

// VertexAIMemoryObservedState is the state of the VertexAIMemory resource as most recently observed in GCP.
type VertexAIMemoryObservedState struct {
	// Output only. Timestamp when this Memory was created.
	CreateTime *string `json:"createTime,omitempty"`

	// Output only. Timestamp when this Memory was most recently updated.
	UpdateTime *string `json:"updateTime,omitempty"`
}

// +genclient
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:resource:categories=gcp,shortName=gcpvertexaimemory;gcpvertexaimemories
// +kubebuilder:subresource:status
// +kubebuilder:metadata:labels="cnrm.cloud.google.com/managed-by-kcc=true"
// +kubebuilder:metadata:labels="cnrm.cloud.google.com/system=true"
// +kubebuilder:metadata:labels="cnrm.cloud.google.com/stability-level=alpha"
// +kubebuilder:printcolumn:name="Age",JSONPath=".metadata.creationTimestamp",type="date"
// +kubebuilder:printcolumn:name="Ready",JSONPath=".status.conditions[?(@.type=='Ready')].status",type="string",description="When 'True', the most recent reconcile of the resource succeeded"
// +kubebuilder:printcolumn:name="Status",JSONPath=".status.conditions[?(@.type=='Ready')].reason",type="string",description="The reason for the value in 'Ready'"
// +kubebuilder:printcolumn:name="Status Age",JSONPath=".status.conditions[?(@.type=='Ready')].lastTransitionTime",type="date",description="The last transition time for the value in 'Status'"

// VertexAIMemory is the Schema for the VertexAIMemory API
// +k8s:openapi-gen=true
type VertexAIMemory struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	// +required
	Spec   VertexAIMemorySpec   `json:"spec,omitempty"`
	Status VertexAIMemoryStatus `json:"status,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// VertexAIMemoryList contains a list of VertexAIMemory
type VertexAIMemoryList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []VertexAIMemory `json:"items"`
}

func init() {
	SchemeBuilder.Register(&VertexAIMemory{}, &VertexAIMemoryList{})
}
