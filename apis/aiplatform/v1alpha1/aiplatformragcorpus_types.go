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

// API sources for AIPlatformRAGCorpus, recorded by generate-types:
// +kcc:source:proto=https://github.com/googleapis/googleapis/blob/1765b559c42386788ff0c6412491277b4791107a/google/cloud/aiplatform/v1/vertex_rag_data.proto
// +kcc:source:service-docs=https://cloud.google.com/vertex-ai/
// +kcc:source:resource-docs=https://cloud.google.com/vertex-ai/docs/reference/rest/v1/projects.locations.ragCorpora

package v1alpha1

import (
	discoveryenginev1alpha1 "github.com/GoogleCloudPlatform/k8s-config-connector/apis/discoveryengine/v1alpha1"
	refsv1beta1 "github.com/GoogleCloudPlatform/k8s-config-connector/apis/refs/v1beta1"
	vertexaiv1beta1 "github.com/GoogleCloudPlatform/k8s-config-connector/apis/vertexai/v1beta1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/apis/k8s/v1alpha1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

var AIPlatformRAGCorpusGVK = GroupVersion.WithKind("AIPlatformRAGCorpus")

// AIPlatformRAGCorpusSpec defines the desired state of AIPlatformRAGCorpus
// +kcc:spec:proto=google.cloud.aiplatform.v1.RagCorpus
// +kcc:required-from-proto
type AIPlatformRAGCorpusSpec struct {
	// The project that this resource belongs to.
	ProjectRef *refsv1beta1.ProjectRef `json:"projectRef"`

	// The location of this resource.
	Location *string `json:"location"`

	// Immutable. Optional. The service-generated name of the resource. Used for acquisition only. Leave unset to create a new resource.
	ResourceID *string `json:"resourceID,omitempty"`
	// Optional. Immutable. The config for the Vector DBs.
	// +kcc:proto:field=google.cloud.aiplatform.v1.RagCorpus.vector_db_config
	VectorDbConfig *RagVectorDbConfig `json:"vectorDbConfig,omitempty"`

	// Optional. Immutable. The config for the Vertex AI Search.
	// +kcc:proto:field=google.cloud.aiplatform.v1.RagCorpus.vertex_ai_search_config
	VertexAiSearchConfig *VertexAiSearchConfig `json:"vertexAiSearchConfig,omitempty"`

	// Required. The display name of the RagCorpus.
	//  The name can be up to 128 characters long and can consist of any UTF-8
	//  characters.
	// +kcc:proto:field=google.cloud.aiplatform.v1.RagCorpus.display_name
	// +required
	DisplayName *string `json:"displayName,omitempty"`

	// Optional. The description of the RagCorpus.
	// +kcc:proto:field=google.cloud.aiplatform.v1.RagCorpus.description
	Description *string `json:"description,omitempty"`

	// Optional. Immutable. The CMEK key name used to encrypt at-rest data related
	//  to this Corpus. Only applicable to RagManagedDb option for Vector DB. This
	//  field can only be set at corpus creation time, and cannot be updated or
	//  deleted.
	// +kcc:proto:field=google.cloud.aiplatform.v1.RagCorpus.encryption_spec
	EncryptionSpec *EncryptionSpecWithRequired `json:"encryptionSpec,omitempty"`
}

// AIPlatformRAGCorpusStatus defines the config connector machine state of AIPlatformRAGCorpus
type AIPlatformRAGCorpusStatus struct {
	/* Conditions represent the latest available observations of the
	   object's current state. */
	Conditions []v1alpha1.Condition `json:"conditions,omitempty"`

	// ObservedGeneration is the generation of the resource that was most recently observed by the Config Connector controller. If this is equal to metadata.generation, then that means that the current reported status reflects the most recent desired state of the resource.
	ObservedGeneration *int64 `json:"observedGeneration,omitempty"`

	// A unique specifier for the AIPlatformRAGCorpus resource in GCP.
	ExternalRef *string `json:"externalRef,omitempty"`

	// ObservedState is the state of the resource as most recently observed in GCP.
	ObservedState *AIPlatformRAGCorpusObservedState `json:"observedState,omitempty"`
}

// AIPlatformRAGCorpusObservedState is the state of the AIPlatformRAGCorpus resource as most recently observed in GCP.
// +kcc:observedstate:proto=google.cloud.aiplatform.v1.RagCorpus
type AIPlatformRAGCorpusObservedState struct {
	// Optional. Immutable. The config for the Vector DBs.
	// +kcc:proto:field=google.cloud.aiplatform.v1.RagCorpus.vector_db_config
	VectorDbConfig *RagVectorDbConfigObservedState `json:"vectorDbConfig,omitempty"`

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

// +genclient
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:resource:categories=gcp,shortName=gcpaiplatformragcorpus;gcpaiplatformragcorpora
// +kubebuilder:subresource:status
// +kubebuilder:metadata:labels="cnrm.cloud.google.com/managed-by-kcc=true"
// +kubebuilder:metadata:labels="cnrm.cloud.google.com/system=true"
// +kubebuilder:printcolumn:name="Age",JSONPath=".metadata.creationTimestamp",type="date"
// +kubebuilder:printcolumn:name="Ready",JSONPath=".status.conditions[?(@.type=='Ready')].status",type="string",description="When 'True', the most recent reconcile of the resource succeeded"
// +kubebuilder:printcolumn:name="Status",JSONPath=".status.conditions[?(@.type=='Ready')].reason",type="string",description="The reason for the value in 'Ready'"
// +kubebuilder:printcolumn:name="Status Age",JSONPath=".status.conditions[?(@.type=='Ready')].lastTransitionTime",type="date",description="The last transition time for the value in 'Status'"

// AIPlatformRAGCorpus is the Schema for the AIPlatformRAGCorpus API
// +k8s:openapi-gen=true
type AIPlatformRAGCorpus struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	// +required
	Spec   AIPlatformRAGCorpusSpec   `json:"spec,omitempty"`
	Status AIPlatformRAGCorpusStatus `json:"status,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// AIPlatformRAGCorpusList contains a list of AIPlatformRAGCorpus
type AIPlatformRAGCorpusList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []AIPlatformRAGCorpus `json:"items"`
}

func init() {
	SchemeBuilder.Register(&AIPlatformRAGCorpus{}, &AIPlatformRAGCorpusList{})
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
	//  e.g. projects/{project}/secrets/{secret}/versions/{version}
	// +kcc:proto:field=google.cloud.aiplatform.v1.ApiAuth.ApiKeyConfig.api_key_secret_version
	// +required
	APIKeySecretVersionRef *refsv1beta1.SecretManagerSecretVersionRef `json:"apiKeySecretVersionRef,omitempty"`
}

// +kcc:proto=google.cloud.aiplatform.v1.RagEmbeddingModelConfig.VertexPredictionEndpoint
type RagEmbeddingModelConfig_VertexPredictionEndpoint struct {
	// Required. The endpoint resource name.
	//  Format:
	//  `projects/{project}/locations/{location}/publishers/{publisher}/models/{model}`
	//  or
	//  `projects/{project}/locations/{location}/endpoints/{endpoint}`
	// +kcc:proto:field=google.cloud.aiplatform.v1.RagEmbeddingModelConfig.VertexPredictionEndpoint.endpoint
	// +required
	EndpointRef *vertexaiv1beta1.VertexAIEndpointRef `json:"endpointRef,omitempty"`
}

// +kcc:observedstate:proto=google.cloud.aiplatform.v1.RagEmbeddingModelConfig.VertexPredictionEndpoint
type RagEmbeddingModelConfig_VertexPredictionEndpointObservedState struct {
	// Output only. The resource name of the model that is deployed on the
	//  endpoint. Present only when the endpoint is not a publisher model.
	//  Pattern:
	//  `projects/{project}/locations/{location}/models/{model}`
	// +kcc:proto:field=google.cloud.aiplatform.v1.RagEmbeddingModelConfig.VertexPredictionEndpoint.model
	Model *string `json:"model,omitempty"`

	// Output only. Version ID of the model that is deployed on the endpoint.
	//  Present only when the endpoint is not a publisher model.
	// +kcc:proto:field=google.cloud.aiplatform.v1.RagEmbeddingModelConfig.VertexPredictionEndpoint.model_version_id
	ModelVersionID *string `json:"modelVersionID,omitempty"`
}

// +kcc:proto=google.cloud.aiplatform.v1.RagVectorDbConfig.VertexVectorSearch
type RagVectorDbConfig_VertexVectorSearch struct {
	// The resource name of the Index Endpoint.
	//  Format:
	//  `projects/{project}/locations/{location}/indexEndpoints/{index_endpoint}`
	// +kcc:proto:field=google.cloud.aiplatform.v1.RagVectorDbConfig.VertexVectorSearch.index_endpoint
	IndexEndpointRef *AIPlatformIndexEndpointRef `json:"indexEndpointRef,omitempty"`

	// The resource name of the Index.
	//  Format:
	//  `projects/{project}/locations/{location}/indexes/{index}`
	// +kcc:proto:field=google.cloud.aiplatform.v1.RagVectorDbConfig.VertexVectorSearch.index
	IndexRef *vertexaiv1beta1.VertexAIIndexRef `json:"indexRef,omitempty"`
}

// +kcc:proto=google.cloud.aiplatform.v1.VertexAiSearchConfig
type VertexAiSearchConfig struct {
	// Vertex AI Search Serving Config resource full name. For example,
	//  `projects/{project}/locations/{location}/collections/{collection}/engines/{engine}/servingConfigs/{serving_config}`
	//  or
	//  `projects/{project}/locations/{location}/collections/{collection}/dataStores/{data_store}/servingConfigs/{serving_config}`.
	// +kcc:proto:field=google.cloud.aiplatform.v1.VertexAiSearchConfig.serving_config
	ServingConfigRef *discoveryenginev1alpha1.DiscoveryEngineServingConfigRef `json:"servingConfigRef,omitempty"`
}
