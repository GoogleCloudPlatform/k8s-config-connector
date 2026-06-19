// Copyright 2026 Google LLC
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package v1alpha1

import (
	discoveryenginev1alpha1 "github.com/GoogleCloudPlatform/k8s-config-connector/apis/discoveryengine/v1alpha1"
)

// VertexAIRagCorpusRef defines the resource reference to Vertex AI RagCorpus.
type VertexAIRagCorpusRef struct {
	// A reference to an externally managed Vertex AI RagCorpus resource.
	// Should be in the format "projects/{{projectID}}/locations/{{location}}/ragCorpora/{{ragCorpusID}}".
	External string `json:"external,omitempty"`

	// The name of a Vertex AI RagCorpus resource.
	Name string `json:"name,omitempty"`

	// The namespace of a Vertex AI RagCorpus resource.
	Namespace string `json:"namespace,omitempty"`
}

// +kcc:proto=google.cloud.aiplatform.v1.VertexAISearch
type VertexAiSearch struct {
	// Optional. Fully-qualified Vertex AI Search data store resource ID.
	//  Format:
	//  `projects/{project}/locations/{location}/collections/{collection}/dataStores/{dataStore}`
	// +kcc:proto:field=google.cloud.aiplatform.v1.VertexAISearch.datastore
	DatastoreRef *discoveryenginev1alpha1.DiscoveryEngineDataStoreRef `json:"datastoreRef,omitempty"`

	// Optional. Fully-qualified Vertex AI Search engine resource ID.
	//  Format:
	//  `projects/{project}/locations/{location}/collections/{collection}/engines/{engine}`
	// +kcc:proto:field=google.cloud.aiplatform.v1.VertexAISearch.engine
	EngineRef *discoveryenginev1alpha1.DiscoveryEngineEngineRef `json:"engineRef,omitempty"`

	// Optional. Number of search results to return per query.
	//  The default value is 10.
	//  The maximumm allowed value is 10.
	// +kcc:proto:field=google.cloud.aiplatform.v1.VertexAISearch.max_results
	MaxResults *int32 `json:"maxResults,omitempty"`

	// Optional. Filter strings to be passed to the search API.
	// +kcc:proto:field=google.cloud.aiplatform.v1.VertexAISearch.filter
	Filter *string `json:"filter,omitempty"`

	// Specifications that define the specific DataStores to be searched, along
	//  with configurations for those data stores. This is only considered for
	//  Engines with multiple data stores.
	//  It should only be set if engine is used.
	// +kcc:proto:field=google.cloud.aiplatform.v1.VertexAISearch.data_store_specs
	DataStoreSpecs []VertexAiSearch_DataStoreSpec `json:"dataStoreSpecs,omitempty"`
}

// +kcc:proto=google.cloud.aiplatform.v1.VertexAISearch.DataStoreSpec
type VertexAiSearch_DataStoreSpec struct {
	// Full resource name of DataStore, such as
	//  Format:
	//  `projects/{project}/locations/{location}/collections/{collection}/dataStores/{dataStore}`
	// +kcc:proto:field=google.cloud.aiplatform.v1.VertexAISearch.DataStoreSpec.data_store
	DataStoreRef *discoveryenginev1alpha1.DiscoveryEngineDataStoreRef `json:"dataStoreRef,omitempty"`

	// Optional. Filter specification to filter documents in the data store
	//  specified by data_store field. For more information on filtering, see
	//  [Filtering](https://cloud.google.com/generative-ai-app-builder/docs/filter-search-metadata)
	// +kcc:proto:field=google.cloud.aiplatform.v1.VertexAISearch.DataStoreSpec.filter
	Filter *string `json:"filter,omitempty"`
}

// +kcc:proto=google.cloud.aiplatform.v1.VertexRagStore
type VertexRagStore struct {
	// Optional. The representation of the rag source. It can be used to specify
	//  corpus only or ragfiles. Currently only support one corpus or multiple
	//  files from one corpus. In the future we may open up multiple corpora
	//  support.
	// +kcc:proto:field=google.cloud.aiplatform.v1.VertexRagStore.rag_resources
	RagResources []VertexRagStore_RagResource `json:"ragResources,omitempty"`

	// Optional. Number of top k results to return from the selected corpora.
	// +kcc:proto:field=google.cloud.aiplatform.v1.VertexRagStore.similarity_top_k
	SimilarityTopK *int32 `json:"similarityTopK,omitempty"`

	// Optional. Only return results with vector distance smaller than the
	//  threshold.
	// +kcc:proto:field=google.cloud.aiplatform.v1.VertexRagStore.vector_distance_threshold
	VectorDistanceThreshold *float64 `json:"vectorDistanceThreshold,omitempty"`

	// Optional. The retrieval config for the Rag query.
	// +kcc:proto:field=google.cloud.aiplatform.v1.VertexRagStore.rag_retrieval_config
	RagRetrievalConfig *RagRetrievalConfig `json:"ragRetrievalConfig,omitempty"`
}

// +kcc:proto=google.cloud.aiplatform.v1.VertexRagStore.RagResource
type VertexRagStore_RagResource struct {
	// Optional. RagCorpora resource name.
	//  Format:
	//  `projects/{project}/locations/{location}/ragCorpora/{rag_corpus}`
	// +kcc:proto:field=google.cloud.aiplatform.v1.VertexRagStore.RagResource.rag_corpus
	RagCorpusRef *VertexAIRagCorpusRef `json:"ragCorpusRef,omitempty"`

	// Optional. rag_file_id. The files should be in the same rag_corpus set in
	//  rag_corpus field.
	// +kcc:proto:field=google.cloud.aiplatform.v1.VertexRagStore.RagResource.rag_file_ids
	RagFileIDs []string `json:"ragFileIDs,omitempty"`
}
