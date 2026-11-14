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

// API sources for AIPlatformCachedContent, recorded by generate-types:
// +kcc:source:proto=https://github.com/googleapis/googleapis/blob/1765b559c42386788ff0c6412491277b4791107a/google/cloud/aiplatform/v1/cached_content.proto
// +kcc:source:service-docs=https://cloud.google.com/vertex-ai/
// +kcc:source:resource-docs=https://cloud.google.com/vertex-ai/docs/reference/rest/v1/projects.locations.cachedContents

package v1alpha1

import (
	"github.com/GoogleCloudPlatform/k8s-config-connector/apis/common/parent"
	discoveryenginev1alpha1 "github.com/GoogleCloudPlatform/k8s-config-connector/apis/discoveryengine/v1alpha1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/apis/k8s/v1alpha1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

var AIPlatformCachedContentGVK = GroupVersion.WithKind("AIPlatformCachedContent")

// AIPlatformCachedContentSpec defines the desired state of AIPlatformCachedContent
// +kcc:spec:proto=google.cloud.aiplatform.v1.CachedContent
// +kcc:required-from-proto
type AIPlatformCachedContentSpec struct {
	*parent.ProjectAndLocationRef `json:",inline"`

	// Immutable. Optional. The service-generated name of the resource. Used for acquisition only. Leave unset to create a new resource.
	ResourceID *string `json:"resourceID,omitempty"`
	// Timestamp of when this resource is considered expired.
	//  This is *always* provided on output, regardless of what was sent
	//  on input.
	// +kcc:proto:field=google.cloud.aiplatform.v1.CachedContent.expire_time
	ExpireTime *string `json:"expireTime,omitempty"`

	// Input only. The TTL for this resource. The expiration time is computed:
	//  now + TTL.
	// +kcc:proto:field=google.cloud.aiplatform.v1.CachedContent.ttl
	TTL *string `json:"ttl,omitempty"`

	// Optional. Immutable. The user-generated meaningful display name of the
	//  cached content.
	// +kcc:proto:field=google.cloud.aiplatform.v1.CachedContent.display_name
	DisplayName *string `json:"displayName,omitempty"`

	// Immutable. The name of the `Model` to use for cached content. Currently,
	//  only the published Gemini base models are supported, in form of
	//  projects/{PROJECT}/locations/{LOCATION}/publishers/google/models/{MODEL}
	// +kcc:proto:field=google.cloud.aiplatform.v1.CachedContent.model
	ModelRef *AIPlatformModelRef `json:"modelRef,omitempty"`

	// Optional. Input only. Immutable. Developer set system instruction.
	//  Currently, text only
	// +kcc:proto:field=google.cloud.aiplatform.v1.CachedContent.system_instruction
	SystemInstruction *ContentWithRequired `json:"systemInstruction,omitempty"`

	// Optional. Input only. Immutable. The content to cache
	// +kcc:proto:field=google.cloud.aiplatform.v1.CachedContent.contents
	Contents []ContentWithRequired `json:"contents,omitempty"`

	// Optional. Input only. Immutable. A list of `Tools` the model may use to
	//  generate the next response
	// +kcc:proto:field=google.cloud.aiplatform.v1.CachedContent.tools
	Tools []Tool `json:"tools,omitempty"`

	// Optional. Input only. Immutable. Tool config. This config is shared for all
	//  tools
	// +kcc:proto:field=google.cloud.aiplatform.v1.CachedContent.tool_config
	ToolConfig *ToolConfig `json:"toolConfig,omitempty"`

	// Input only. Immutable. Customer-managed encryption key spec for a
	//  `CachedContent`. If set, this `CachedContent` and all its sub-resources
	//  will be secured by this key.
	// +kcc:proto:field=google.cloud.aiplatform.v1.CachedContent.encryption_spec
	EncryptionSpec *EncryptionSpecWithRequired `json:"encryptionSpec,omitempty"`
}

// AIPlatformCachedContentStatus defines the config connector machine state of AIPlatformCachedContent
type AIPlatformCachedContentStatus struct {
	/* Conditions represent the latest available observations of the
	   object's current state. */
	Conditions []v1alpha1.Condition `json:"conditions,omitempty"`

	// ObservedGeneration is the generation of the resource that was most recently observed by the Config Connector controller. If this is equal to metadata.generation, then that means that the current reported status reflects the most recent desired state of the resource.
	ObservedGeneration *int64 `json:"observedGeneration,omitempty"`

	// A unique specifier for the AIPlatformCachedContent resource in GCP.
	ExternalRef *string `json:"externalRef,omitempty"`

	// ObservedState is the state of the resource as most recently observed in GCP.
	ObservedState *AIPlatformCachedContentObservedState `json:"observedState,omitempty"`
}

// AIPlatformCachedContentObservedState is the state of the AIPlatformCachedContent resource as most recently observed in GCP.
// +kcc:observedstate:proto=google.cloud.aiplatform.v1.CachedContent
type AIPlatformCachedContentObservedState struct {
	// Output only. Creation time of the cache entry.
	// +kcc:proto:field=google.cloud.aiplatform.v1.CachedContent.create_time
	CreateTime *string `json:"createTime,omitempty"`

	// Output only. When the cache entry was last updated in UTC time.
	// +kcc:proto:field=google.cloud.aiplatform.v1.CachedContent.update_time
	UpdateTime *string `json:"updateTime,omitempty"`

	// Output only. Metadata on the usage of the cached content.
	// +kcc:proto:field=google.cloud.aiplatform.v1.CachedContent.usage_metadata
	UsageMetadata *CachedContent_UsageMetadata `json:"usageMetadata,omitempty"`
}

// +genclient
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:resource:categories=gcp,shortName=gcpaiplatformcachedcontent;gcpaiplatformcachedcontents
// +kubebuilder:subresource:status
// +kubebuilder:metadata:labels="cnrm.cloud.google.com/managed-by-kcc=true"
// +kubebuilder:metadata:labels="cnrm.cloud.google.com/system=true"
// +kubebuilder:printcolumn:name="Age",JSONPath=".metadata.creationTimestamp",type="date"
// +kubebuilder:printcolumn:name="Ready",JSONPath=".status.conditions[?(@.type=='Ready')].status",type="string",description="When 'True', the most recent reconcile of the resource succeeded"
// +kubebuilder:printcolumn:name="Status",JSONPath=".status.conditions[?(@.type=='Ready')].reason",type="string",description="The reason for the value in 'Ready'"
// +kubebuilder:printcolumn:name="Status Age",JSONPath=".status.conditions[?(@.type=='Ready')].lastTransitionTime",type="date",description="The last transition time for the value in 'Status'"

// AIPlatformCachedContent is the Schema for the AIPlatformCachedContent API
// +k8s:openapi-gen=true
type AIPlatformCachedContent struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	// +required
	Spec   AIPlatformCachedContentSpec   `json:"spec,omitempty"`
	Status AIPlatformCachedContentStatus `json:"status,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// AIPlatformCachedContentList contains a list of AIPlatformCachedContent
type AIPlatformCachedContentList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []AIPlatformCachedContent `json:"items"`
}

func init() {
	SchemeBuilder.Register(&AIPlatformCachedContent{}, &AIPlatformCachedContentList{})
}

// +kcc:proto=google.cloud.aiplatform.v1.VertexAISearch
type VertexAiSearch struct {
	// Optional. Fully-qualified Vertex AI Search data store resource ID.
	// Format:
	// `projects/{project}/locations/{location}/collections/{collection}/dataStores/{dataStore}`
	// +kcc:proto:field=google.cloud.aiplatform.v1.VertexAISearch.datastore
	DatastoreRef *discoveryenginev1alpha1.DiscoveryEngineDataStoreRef `json:"datastoreRef,omitempty"`

	// Optional. Fully-qualified Vertex AI Search engine resource ID.
	// Format:
	// `projects/{project}/locations/{location}/collections/{collection}/engines/{engine}`
	// +kcc:proto:field=google.cloud.aiplatform.v1.VertexAISearch.engine
	EngineRef *discoveryenginev1alpha1.DiscoveryEngineEngineRef `json:"engineRef,omitempty"`

	// Optional. Number of search results to return per query.
	// The default value is 10.
	// The maximumm allowed value is 10.
	// +kcc:proto:field=google.cloud.aiplatform.v1.VertexAISearch.max_results
	MaxResults *int32 `json:"maxResults,omitempty"`

	// Optional. Filter strings to be passed to the search API.
	// +kcc:proto:field=google.cloud.aiplatform.v1.VertexAISearch.filter
	Filter *string `json:"filter,omitempty"`

	// Specifications that define the specific DataStores to be searched, along
	// with configurations for those data stores. This is only considered for
	// Engines with multiple data stores.
	// It should only be set if engine is used.
	// +kcc:proto:field=google.cloud.aiplatform.v1.VertexAISearch.data_store_specs
	DataStoreSpecs []VertexAiSearch_DataStoreSpec `json:"dataStoreSpecs,omitempty"`
}

// +kcc:proto=google.cloud.aiplatform.v1.VertexAISearch.DataStoreSpec
type VertexAiSearch_DataStoreSpec struct {
	// Full resource name of DataStore, such as
	// Format:
	// `projects/{project}/locations/{location}/collections/{collection}/dataStores/{dataStore}`
	// +kcc:proto:field=google.cloud.aiplatform.v1.VertexAISearch.DataStoreSpec.data_store
	DataStoreRef *discoveryenginev1alpha1.DiscoveryEngineDataStoreRef `json:"dataStoreRef,omitempty"`

	// Optional. Filter specification to filter documents in the data store
	// specified by data_store field. For more information on filtering, see
	// [Filtering](https://cloud.google.com/generative-ai-app-builder/docs/filter-search-metadata)
	// +kcc:proto:field=google.cloud.aiplatform.v1.VertexAISearch.DataStoreSpec.filter
	Filter *string `json:"filter,omitempty"`
}

// +kcc:proto=google.cloud.aiplatform.v1.VertexRagStore.RagResource
type VertexRagStore_RagResource struct {
	// Optional. RagCorpora resource name.
	// Format:
	// `projects/{project}/locations/{location}/ragCorpora/{rag_corpus}`
	// +kcc:proto:field=google.cloud.aiplatform.v1.VertexRagStore.RagResource.rag_corpus
	RagCorpusRef *AIPlatformRAGCorpusRef `json:"ragCorpusRef,omitempty"`

	// Optional. rag_file_id. The files should be in the same rag_corpus set in
	// rag_corpus field.
	// +kcc:proto:field=google.cloud.aiplatform.v1.VertexRagStore.RagResource.rag_file_ids
	RagFileIDs []string `json:"ragFileIDs,omitempty"`
}
