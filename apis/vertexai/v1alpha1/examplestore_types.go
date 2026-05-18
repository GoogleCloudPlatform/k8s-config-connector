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

var VertexAIExampleStoreGVK = GroupVersion.WithKind("VertexAIExampleStore")

// VertexAIExampleStoreSpec defines the desired state of VertexAIExampleStore
// +kcc:spec:proto=google.cloud.aiplatform.v1beta1.ExampleStore
type VertexAIExampleStoreSpec struct {
	// The project that this resource belongs to.
	// +required
	ProjectRef *refsv1beta1.ProjectRef `json:"projectRef"`

	// The location of this resource.
	// +required
	Location string `json:"location"`

	// The VertexAIExampleStore name. If not given, the metadata.name will be used.
	ResourceID *string `json:"resourceID,omitempty"`

	// Required. Display name of the ExampleStore.
	// +kcc:proto:field=google.cloud.aiplatform.v1beta1.ExampleStore.display_name
	// +required
	DisplayName string `json:"displayName"`

	// Optional. Description of the ExampleStore.
	// +kcc:proto:field=google.cloud.aiplatform.v1beta1.ExampleStore.description
	Description *string `json:"description,omitempty"`

	// Required. Example Store config.
	// +kcc:proto:field=google.cloud.aiplatform.v1beta1.ExampleStore.example_store_config
	// +required
	ExampleStoreConfig *ExampleStoreConfig `json:"exampleStoreConfig"`
}

// VertexAIExampleStoreStatus defines the config connector machine state of VertexAIExampleStore
type VertexAIExampleStoreStatus struct {
	/* Conditions represent the latest available observations of the
	   object's current state. */
	Conditions []v1alpha1.Condition `json:"conditions,omitempty"`

	// ObservedGeneration is the generation of the resource that was most recently observed by the Config Connector controller. If this is equal to metadata.generation, then that means that the current reported status reflects the most recent desired state of the resource.
	ObservedGeneration *int64 `json:"observedGeneration,omitempty"`

	// A unique specifier for the VertexAIExampleStore resource in GCP.
	ExternalRef *string `json:"externalRef,omitempty"`

	// ObservedState is the state of the resource as most recently observed in GCP.
	ObservedState *VertexAIExampleStoreObservedState `json:"observedState,omitempty"`
}

// VertexAIExampleStoreObservedState is the state of the VertexAIExampleStore resource as most recently observed in GCP.
// +kcc:observedstate:proto=google.cloud.aiplatform.v1beta1.ExampleStore
type VertexAIExampleStoreObservedState struct {
	// Output only. Timestamp when this ExampleStore was created.
	// +kcc:proto:field=google.cloud.aiplatform.v1beta1.ExampleStore.create_time
	CreateTime *string `json:"createTime,omitempty"`

	// Output only. Timestamp when this ExampleStore was last updated.
	// +kcc:proto:field=google.cloud.aiplatform.v1beta1.ExampleStore.update_time
	UpdateTime *string `json:"updateTime,omitempty"`
}

// +genclient
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:resource:categories=gcp,shortName=gcpvertexaiexamplestore;gcpvertexaiexamplestores
// +kubebuilder:subresource:status
// +kubebuilder:metadata:labels="cnrm.cloud.google.com/managed-by-kcc=true"
// +kubebuilder:metadata:labels="cnrm.cloud.google.com/system=true"
// +kubebuilder:metadata:labels="cnrm.cloud.google.com/stability-level=alpha"
// +kubebuilder:printcolumn:name="Age",JSONPath=".metadata.creationTimestamp",type="date"
// +kubebuilder:printcolumn:name="Ready",JSONPath=".status.conditions[?(@.type=='Ready')].status",type="string",description="When 'True', the most recent reconcile of the resource succeeded"
// +kubebuilder:printcolumn:name="Status",JSONPath=".status.conditions[?(@.type=='Ready')].reason",type="string",description="The reason for the value in 'Ready'"
// +kubebuilder:printcolumn:name="Status Age",JSONPath=".status.conditions[?(@.type=='Ready')].lastTransitionTime",type="date",description="The last transition time for the value in 'Status'"

// VertexAIExampleStore is the Schema for the VertexAIExampleStore API
// +k8s:openapi-gen=true
type VertexAIExampleStore struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	// +required
	Spec   VertexAIExampleStoreSpec   `json:"spec,omitempty"`
	Status VertexAIExampleStoreStatus `json:"status,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// VertexAIExampleStoreList contains a list of VertexAIExampleStore
type VertexAIExampleStoreList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []VertexAIExampleStore `json:"items"`
}

func init() {
	SchemeBuilder.Register(&VertexAIExampleStore{}, &VertexAIExampleStoreList{})
}
