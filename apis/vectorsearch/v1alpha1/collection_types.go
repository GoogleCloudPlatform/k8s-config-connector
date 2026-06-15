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
	apiextensionsv1 "k8s.io/apiextensions-apiserver/pkg/apis/apiextensions/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

var VectorSearchCollectionGVK = GroupVersion.WithKind("VectorSearchCollection")

// VectorSearchCollectionSpec defines the desired state of VectorSearchCollection
// +kcc:spec:proto=google.cloud.vectorsearch.v1.Collection
type VectorSearchCollectionSpec struct {
	// The project that this resource belongs to.
	ProjectRef *refsv1beta1.ProjectRef `json:"projectRef"`

	// The location of this resource.
	Location string `json:"location"`

	// The VectorSearchCollection name. If not given, the metadata.name will be used.
	ResourceID *string `json:"resourceID,omitempty"`

	// Optional. User-specified display name of the collection
	DisplayName *string `json:"displayName,omitempty"`

	// Optional. User-specified description of the collection
	Description *string `json:"description,omitempty"`

	// Optional. Labels as key value pairs.
	Labels map[string]string `json:"labels,omitempty"`

	// Optional. Schema for vector fields. Only vector fields in this schema will
	//  be searchable. Field names must contain only alphanumeric characters,
	//  underscores, and hyphens.
	VectorSchema map[string]VectorField `json:"vectorSchema,omitempty"`

	// Optional. JSON Schema for data.
	//  Field names must contain only alphanumeric characters,
	//  underscores, and hyphens.
	DataSchema *apiextensionsv1.JSON `json:"dataSchema,omitempty"`
}

// +kcc:proto=google.cloud.vectorsearch.v1.VectorField
type VectorField struct {
	// Dense vector field.
	// +kcc:proto:field=google.cloud.vectorsearch.v1.VectorField.dense_vector
	DenseVector *DenseVectorField `json:"denseVector,omitempty"`

	// Sparse vector field.
	// +kcc:proto:field=google.cloud.vectorsearch.v1.VectorField.sparse_vector
	SparseVector *SparseVectorField `json:"sparseVector,omitempty"`
}

// +kcc:proto=google.cloud.vectorsearch.v1.DenseVectorField
type DenseVectorField struct {
	// Dimensionality of the vector field.
	// +kcc:proto:field=google.cloud.vectorsearch.v1.DenseVectorField.dimensions
	Dimensions *int32 `json:"dimensions,omitempty"`

	// Optional. Configuration for generating embeddings for the vector field. If
	//  not specified, the embedding field must be populated in the DataObject.
	// +kcc:proto:field=google.cloud.vectorsearch.v1.DenseVectorField.vertex_embedding_config
	VertexEmbeddingConfig *VertexEmbeddingConfig `json:"vertexEmbeddingConfig,omitempty"`
}

// +kubebuilder:validation:XPreserveUnknownFields
// +kcc:proto=google.cloud.vectorsearch.v1.SparseVectorField
type SparseVectorField struct {
}

// +kcc:proto=google.cloud.vectorsearch.v1.VertexEmbeddingConfig
type VertexEmbeddingConfig struct {
	// Required. Required: ID of the embedding model to use. See
	//  https://cloud.google.com/vertex-ai/generative-ai/docs/learn/models#embeddings-models
	//  for the list of supported models.
	// +required
	// +kcc:proto:field=google.cloud.vectorsearch.v1.VertexEmbeddingConfig.model_id
	ModelID *string `json:"modelID,omitempty"`

	// Required. Required: Text template for the input to the model. The template
	//  must contain one or more references to fields in the DataObject, e.g.:
	//  "Movie Title: {title} ---- Movie Plot: {plot}".
	// +required
	// +kcc:proto:field=google.cloud.vectorsearch.v1.VertexEmbeddingConfig.text_template
	TextTemplate *string `json:"textTemplate,omitempty"`

	// Required. Required: Task type for the embeddings.
	// +required
	// +kcc:proto:field=google.cloud.vectorsearch.v1.VertexEmbeddingConfig.task_type
	TaskType *string `json:"taskType,omitempty"`
}

// VectorSearchCollectionStatus defines the config connector machine state of VectorSearchCollection
type VectorSearchCollectionStatus struct {
	/* Conditions represent the latest available observations of the
	   object's current state. */
	Conditions []v1alpha1.Condition `json:"conditions,omitempty"`

	// ObservedGeneration is the generation of the resource that was most recently observed by the Config Connector controller. If this is equal to metadata.generation, then that means that the current reported status reflects the most recent desired state of the resource.
	ObservedGeneration *int64 `json:"observedGeneration,omitempty"`

	// A unique specifier for the VectorSearchCollection resource in GCP.
	ExternalRef *string `json:"externalRef,omitempty"`

	// ObservedState is the state of the resource as most recently observed in GCP.
	ObservedState *VectorSearchCollectionObservedState `json:"observedState,omitempty"`
}

// VectorSearchCollectionObservedState is the state of the VectorSearchCollection resource as most recently observed in GCP.
// +kcc:observedstate:proto=google.cloud.vectorsearch.v1.Collection
type VectorSearchCollectionObservedState struct {
	// Output only. [Output only] Create time stamp
	// +kcc:proto:field=google.cloud.vectorsearch.v1.Collection.create_time
	CreateTime *string `json:"createTime,omitempty"`

	// Output only. [Output only] Update time stamp
	// +kcc:proto:field=google.cloud.vectorsearch.v1.Collection.update_time
	UpdateTime *string `json:"updateTime,omitempty"`
}

// +genclient
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:resource:categories=gcp,shortName=gcpvectorsearchcollection;gcpvectorsearchcollections
// +kubebuilder:subresource:status
// +kubebuilder:metadata:labels="cnrm.cloud.google.com/managed-by-kcc=true"
// +kubebuilder:metadata:labels="cnrm.cloud.google.com/system=true"
// +kubebuilder:metadata:labels="cnrm.cloud.google.com/stability-level=alpha"
// +kubebuilder:printcolumn:name="Age",JSONPath=".metadata.creationTimestamp",type="date"
// +kubebuilder:printcolumn:name="Ready",JSONPath=".status.conditions[?(@.type=='Ready')].status",type="string",description="When 'True', the most recent reconcile of the resource succeeded"
// +kubebuilder:printcolumn:name="Status",JSONPath=".status.conditions[?(@.type=='Ready')].reason",type="string",description="The reason for the value in 'Ready'"
// +kubebuilder:printcolumn:name="Status Age",JSONPath=".status.conditions[?(@.type=='Ready')].lastTransitionTime",type="date",description="The last transition time for the value in 'Status'"

// VectorSearchCollection is the Schema for the VectorSearchCollection API
// +k8s:openapi-gen=true
type VectorSearchCollection struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	// +required
	Spec   VectorSearchCollectionSpec   `json:"spec,omitempty"`
	Status VectorSearchCollectionStatus `json:"status,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// VectorSearchCollectionList contains a list of VectorSearchCollection
type VectorSearchCollectionList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []VectorSearchCollection `json:"items"`
}

func init() {
	SchemeBuilder.Register(&VectorSearchCollection{}, &VectorSearchCollectionList{})
}
