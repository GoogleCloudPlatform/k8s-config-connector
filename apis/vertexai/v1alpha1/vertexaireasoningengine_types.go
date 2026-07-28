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

var VertexAIReasoningEngineGVK = GroupVersion.WithKind("VertexAIReasoningEngine")

// VertexAIReasoningEngineSpec defines the desired state of VertexAIReasoningEngine
type VertexAIReasoningEngineSpec struct {
	// The project that this resource belongs to.
	// +required
	ProjectRef *refsv1beta1.ProjectRef `json:"projectRef"`

	// The location of this resource.
	// +required
	Location *string `json:"location"`

	// The VertexAIReasoningEngine name. If not given, the metadata.name will be used.
	ResourceID *string `json:"resourceID,omitempty"`

	// Required. The display name of the ReasoningEngine.
	DisplayName *string `json:"displayName,omitempty"`

	// Optional. The description of the ReasoningEngine.
	Description *string `json:"description,omitempty"`

	// Optional. Configurations of the ReasoningEngine
	Spec *ReasoningEngineSpec `json:"spec,omitempty"`

	// Optional. Configuration for how Agent Engine sub-resources should manage
	//  context.
	ContextSpec *ReasoningEngineContextSpec `json:"contextSpec,omitempty"`

	// Customer-managed encryption key spec for a ReasoningEngine. If set, this
	//  ReasoningEngine and all sub-resources of this ReasoningEngine will be
	//  secured by this key.
	EncryptionSpec *EncryptionSpec `json:"encryptionSpec,omitempty"`

	// Labels for the ReasoningEngine.
	Labels map[string]string `json:"labels,omitempty"`

	// Optional. Traffic distribution configuration for the Reasoning Engine.
	TrafficConfig *ReasoningEngine_TrafficConfig `json:"trafficConfig,omitempty"`
}

// VertexAIReasoningEngineStatus defines the config connector machine state of VertexAIReasoningEngine
type VertexAIReasoningEngineStatus struct {
	/* Conditions represent the latest available observations of the
	   object's current state. */
	Conditions []v1alpha1.Condition `json:"conditions,omitempty"`

	// ObservedGeneration is the generation of the resource that was most recently observed by the Config Connector controller. If this is equal to metadata.generation, then that means that the current reported status reflects the most recent desired state of the resource.
	ObservedGeneration *int64 `json:"observedGeneration,omitempty"`

	// A unique specifier for the VertexAIReasoningEngine resource in GCP.
	ExternalRef *string `json:"externalRef,omitempty"`

	// ObservedState is the state of the resource as most recently observed in GCP.
	ObservedState *VertexAIReasoningEngineObservedState `json:"observedState,omitempty"`
}

// VertexAIReasoningEngineObservedState is the state of the VertexAIReasoningEngine resource as most recently observed in GCP.
type VertexAIReasoningEngineObservedState struct {
	// Output only. Timestamp when this ReasoningEngine was created.
	CreateTime *string `json:"createTime,omitempty"`

	// Output only. Timestamp when this ReasoningEngine was most recently updated.
	UpdateTime *string `json:"updateTime,omitempty"`
}

// +genclient
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:resource:categories=gcp,shortName=gcpvertexaireasoningengine;gcpvertexaireasoningengines
// +kubebuilder:subresource:status
// +kubebuilder:metadata:labels="cnrm.cloud.google.com/managed-by-kcc=true"
// +kubebuilder:metadata:labels="cnrm.cloud.google.com/system=true"
// +kubebuilder:metadata:labels="cnrm.cloud.google.com/stability-level=alpha"
// +kubebuilder:printcolumn:name="Age",JSONPath=".metadata.creationTimestamp",type="date"
// +kubebuilder:printcolumn:name="Ready",JSONPath=".status.conditions[?(@.type=='Ready')].status",type="string",description="When 'True', the most recent reconcile of the resource succeeded"
// +kubebuilder:printcolumn:name="Status",JSONPath=".status.conditions[?(@.type=='Ready')].reason",type="string",description="The reason for the value in 'Ready'"
// +kubebuilder:printcolumn:name="Status Age",JSONPath=".status.conditions[?(@.type=='Ready')].lastTransitionTime",type="date",description="The last transition time for the value in 'Status'"

// VertexAIReasoningEngine is the Schema for the VertexAIReasoningEngine API
// +k8s:openapi-gen=true
type VertexAIReasoningEngine struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	// +required
	Spec   VertexAIReasoningEngineSpec   `json:"spec,omitempty"`
	Status VertexAIReasoningEngineStatus `json:"status,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// VertexAIReasoningEngineList contains a list of VertexAIReasoningEngine
type VertexAIReasoningEngineList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []VertexAIReasoningEngine `json:"items"`
}

func init() {
	SchemeBuilder.Register(&VertexAIReasoningEngine{}, &VertexAIReasoningEngineList{})
}
