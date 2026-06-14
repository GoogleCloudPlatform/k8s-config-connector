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

var VertexAIRAGCorpusGVK = GroupVersion.WithKind("VertexAIRAGCorpus")

// VertexAIRAGCorpusSpec defines the desired state of VertexAIRAGCorpus
// +kcc:spec:proto=google.cloud.aiplatform.v1.RagCorpus
type VertexAIRAGCorpusSpec struct {
	// The project that this resource belongs to.
	ProjectRef *refsv1beta1.ProjectRef `json:"projectRef"`

	// Immutable. The location of this resource.
	Location *string `json:"location"`

	// The VertexAIRAGCorpus name. If not given, the metadata.name will be used.
	ResourceID *string `json:"resourceID,omitempty"`

	// Required. The display name of the RagCorpus.
	// +kcc:proto:field=google.cloud.aiplatform.v1.RagCorpus.display_name
	DisplayName *string `json:"displayName,omitempty"`

	// Optional. The description of the RagCorpus.
	// +kcc:proto:field=google.cloud.aiplatform.v1.RagCorpus.description
	Description *string `json:"description,omitempty"`

	// Optional. Immutable. The config for the Vector DBs.
	// +kcc:proto:field=google.cloud.aiplatform.v1.RagCorpus.vector_db_config
	VectorDbConfig *RagVectorDbConfig `json:"vectorDbConfig,omitempty"`

	// Optional. Immutable. The config for the Vertex AI Search.
	// +kcc:proto:field=google.cloud.aiplatform.v1.RagCorpus.vertex_ai_search_config
	VertexAiSearchConfig *VertexAiSearchConfig `json:"vertexAiSearchConfig,omitempty"`

	// Optional. Immutable. The CMEK key name used to encrypt at-rest data related to this Corpus.
	// +kcc:proto:field=google.cloud.aiplatform.v1.RagCorpus.encryption_spec
	EncryptionSpec *EncryptionSpec `json:"encryptionSpec,omitempty"`
}

// VertexAIRAGCorpusStatus defines the config connector machine state of VertexAIRAGCorpus
type VertexAIRAGCorpusStatus struct {
	/* Conditions represent the latest available observations of the
	   object's current state. */
	Conditions []v1alpha1.Condition `json:"conditions,omitempty"`

	// ObservedGeneration is the generation of the resource that was most recently observed by the Config Connector controller. If this is equal to metadata.generation, then that means that the current reported status reflects the most recent desired state of the resource.
	ObservedGeneration *int64 `json:"observedGeneration,omitempty"`

	// A unique specifier for the VertexAIRAGCorpus resource in GCP.
	ExternalRef *string `json:"externalRef,omitempty"`

	// ObservedState is the state of the resource as most recently observed in GCP.
	ObservedState *VertexAIRAGCorpusObservedState `json:"observedState,omitempty"`
}

// VertexAIRAGCorpusObservedState is the state of the VertexAIRAGCorpus resource as most recently observed in GCP.
// +kcc:observedstate:proto=google.cloud.aiplatform.v1.RagCorpus
type VertexAIRAGCorpusObservedState struct {
	// Output only. Timestamp when this RagCorpus was created.
	// +kcc:proto:field=google.cloud.aiplatform.v1.RagCorpus.create_time
	CreateTime *string `json:"createTime,omitempty"`

	// Output only. Timestamp when this RagCorpus was last updated.
	// +kcc:proto:field=google.cloud.aiplatform.v1.RagCorpus.update_time
	UpdateTime *string `json:"updateTime,omitempty"`

	// Output only. RagCorpus state.
	// +kcc:proto:field=google.cloud.aiplatform.v1.RagCorpus.corpus_status
	CorpusStatus *CorpusStatusObservedState `json:"corpusStatus,omitempty"`
}

// +kcc:proto=google.cloud.aiplatform.v1.ApiAuth
type APIAuth struct {
	// The API secret.
	// +kcc:proto:field=google.cloud.aiplatform.v1.ApiAuth.api_key_config
	APIKeyConfig *APIAuth_APIKeyConfig `json:"apiKeyConfig,omitempty"`
}

// +kcc:proto=google.cloud.aiplatform.v1.ApiAuth.ApiKeyConfig
type APIAuth_APIKeyConfig struct {
	// Required. The SecretManager secret version resource name storing API key.
	// +kcc:proto:field=google.cloud.aiplatform.v1.ApiAuth.ApiKeyConfig.api_key_secret_version
	APIKeySecretVersionRef *refsv1beta1.SecretManagerSecretVersionRef `json:"apiKeySecretVersionRef,omitempty"`
}

// +kcc:proto=google.cloud.aiplatform.v1.RagEmbeddingModelConfig
type RagEmbeddingModelConfig struct {
	// The Vertex AI Prediction Endpoint that either refers to a publisher model
	//  or an endpoint that is hosting a 1P fine-tuned text embedding model.
	// +kcc:proto:field=google.cloud.aiplatform.v1.RagEmbeddingModelConfig.vertex_prediction_endpoint
	VertexPredictionEndpoint *RagEmbeddingModelConfig_VertexPredictionEndpoint `json:"vertexPredictionEndpoint,omitempty"`
}

// +kcc:proto=google.cloud.aiplatform.v1.RagEmbeddingModelConfig.VertexPredictionEndpoint
type RagEmbeddingModelConfig_VertexPredictionEndpoint struct {
	// Required. The endpoint resource name.
	// +kcc:proto:field=google.cloud.aiplatform.v1.RagEmbeddingModelConfig.VertexPredictionEndpoint.endpoint
	EndpointRef *refsv1beta1.VertexAIEndpointRef `json:"endpointRef,omitempty"`
}

// +kcc:proto=google.cloud.aiplatform.v1.RagVectorDbConfig.VertexVectorSearch
type RagVectorDbConfig_VertexVectorSearch struct {
	// The resource name of the Index Endpoint.
	// +kcc:proto:field=google.cloud.aiplatform.v1.RagVectorDbConfig.VertexVectorSearch.index_endpoint
	IndexEndpointRef *refsv1beta1.VertexAIIndexEndpointRef `json:"indexEndpointRef,omitempty"`

	// The resource name of the Index.
	// +kcc:proto:field=google.cloud.aiplatform.v1.RagVectorDbConfig.VertexVectorSearch.index
	IndexRef *refsv1beta1.VertexAIIndexRef `json:"indexRef,omitempty"`
}

// +kcc:proto=google.cloud.aiplatform.v1.VertexAiSearchConfig
type VertexAiSearchConfig struct {
	// Vertex AI Search Serving Config resource full name.
	// +kcc:proto:field=google.cloud.aiplatform.v1.VertexAiSearchConfig.serving_config
	ServingConfigRef *refsv1beta1.VertexAISearchServingConfigRef `json:"servingConfigRef,omitempty"`
}

// +genclient
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:resource:path=vertexairagcorpora
// +kubebuilder:resource:categories=gcp,shortName=gcpvertexairagcorpus;gcpvertexairagcorpora
// +kubebuilder:subresource:status
// +kubebuilder:metadata:labels="cnrm.cloud.google.com/managed-by-kcc=true"
// +kubebuilder:metadata:labels="cnrm.cloud.google.com/system=true"
// +kubebuilder:metadata:labels="cnrm.cloud.google.com/stability-level=alpha"
// +kubebuilder:printcolumn:name="Age",JSONPath=".metadata.creationTimestamp",type="date"
// +kubebuilder:printcolumn:name="Ready",JSONPath=".status.conditions[?(@.type=='Ready')].status",type="string",description="When 'True', the most recent reconcile of the resource succeeded"
// +kubebuilder:printcolumn:name="Status",JSONPath=".status.conditions[?(@.type=='Ready')].reason",type="string",description="The reason for the value in 'Ready'"
// +kubebuilder:printcolumn:name="Status Age",JSONPath=".status.conditions[?(@.type=='Ready')].lastTransitionTime",type="date",description="The last transition time for the value in 'Status'"

// VertexAIRAGCorpus is the Schema for the VertexAIRAGCorpus API
// +k8s:openapi-gen=true
type VertexAIRAGCorpus struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	// +required
	Spec   VertexAIRAGCorpusSpec   `json:"spec,omitempty"`
	Status VertexAIRAGCorpusStatus `json:"status,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// VertexAIRAGCorpusList contains a list of VertexAIRAGCorpus
type VertexAIRAGCorpusList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []VertexAIRAGCorpus `json:"items"`
}

func init() {
	SchemeBuilder.Register(&VertexAIRAGCorpus{}, &VertexAIRAGCorpusList{})
}
