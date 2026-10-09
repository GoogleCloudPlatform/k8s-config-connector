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
// +kcc:guess=source-link reason=verify-resource-docs-link
// +kcc:source:resource-docs=https://docs.cloud.google.com/vertex-ai/docs/reference/rest/v1/projects.locations.cachedContents

package v1alpha1

import (
	refsv1beta1 "github.com/GoogleCloudPlatform/k8s-config-connector/apis/refs/v1beta1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/apis/k8s/v1alpha1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

var AIPlatformCachedContentGVK = GroupVersion.WithKind("AIPlatformCachedContent")

// AIPlatformCachedContentSpec defines the desired state of AIPlatformCachedContent
// +kcc:spec:proto=google.cloud.aiplatform.v1.CachedContent
type AIPlatformCachedContentSpec struct {
	// The project that this resource belongs to.
	ProjectRef *refsv1beta1.ProjectRef `json:"projectRef"`

	// The location of this resource.
	Location string `json:"location"`

	// The AIPlatformCachedContent name. If not given, the metadata.name will be used.
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
	Model *string `json:"model,omitempty"`

	// Optional. Input only. Immutable. Developer set system instruction.
	//  Currently, text only
	// +kcc:proto:field=google.cloud.aiplatform.v1.CachedContent.system_instruction
	SystemInstruction *Content `json:"systemInstruction,omitempty"`

	// Optional. Input only. Immutable. The content to cache
	// +kcc:proto:field=google.cloud.aiplatform.v1.CachedContent.contents
	Contents []Content `json:"contents,omitempty"`

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
	EncryptionSpec *EncryptionSpec `json:"encryptionSpec,omitempty"`
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

// Tool, Retrieval and VertexRagStore are written by hand to leave out
// deprecated fields. generate-types skips types that a hand-written file in
// this package defines.
// Tool.google_search_retrieval is deprecated in the API docs.
// Retrieval.disable_attribution, VertexRagStore.similarity_top_k and
// VertexRagStore.vector_distance_threshold are [deprecated = true] in the
// proto.

// +kcc:proto=google.cloud.aiplatform.v1.Tool
type Tool struct {
	// Optional. Function tool type.
	//  One or more function declarations to be passed to the model along with the
	//  current user query. Model may decide to call a subset of these functions
	//  by populating [FunctionCall][google.cloud.aiplatform.v1.Part.function_call]
	//  in the response. User should provide a
	//  [FunctionResponse][google.cloud.aiplatform.v1.Part.function_response] for
	//  each function call in the next turn. Based on the function responses, Model
	//  will generate the final response back to the user. Maximum 128 function
	//  declarations can be provided.
	// +kcc:proto:field=google.cloud.aiplatform.v1.Tool.function_declarations
	FunctionDeclarations []FunctionDeclaration `json:"functionDeclarations,omitempty"`

	// Optional. Retrieval tool type.
	//  System will always execute the provided retrieval tool(s) to get external
	//  knowledge to answer the prompt. Retrieval results are presented to the
	//  model for generation.
	// +kcc:proto:field=google.cloud.aiplatform.v1.Tool.retrieval
	Retrieval *Retrieval `json:"retrieval,omitempty"`

	// Optional. GoogleSearch tool type.
	//  Tool to support Google Search in Model. Powered by Google.
	// +kcc:proto:field=google.cloud.aiplatform.v1.Tool.google_search
	GoogleSearch *Tool_GoogleSearch `json:"googleSearch,omitempty"`

	// Optional. GoogleMaps tool type.
	//  Tool to support Google Maps in Model.
	// +kcc:proto:field=google.cloud.aiplatform.v1.Tool.google_maps
	GoogleMaps *GoogleMaps `json:"googleMaps,omitempty"`

	// Optional. Tool to support searching public web data, powered by Vertex AI
	//  Search and Sec4 compliance.
	// +kcc:proto:field=google.cloud.aiplatform.v1.Tool.enterprise_web_search
	EnterpriseWebSearch *EnterpriseWebSearch `json:"enterpriseWebSearch,omitempty"`

	// Optional. CodeExecution tool type.
	//  Enables the model to execute code as part of generation.
	// +kcc:proto:field=google.cloud.aiplatform.v1.Tool.code_execution
	CodeExecution *Tool_CodeExecution `json:"codeExecution,omitempty"`

	// Optional. Tool to support URL context retrieval.
	// +kcc:proto:field=google.cloud.aiplatform.v1.Tool.url_context
	URLContext *URLContext `json:"urlContext,omitempty"`

	// Optional. Tool to support the model interacting directly with the computer.
	//  If enabled, it automatically populates computer-use specific Function
	//  Declarations.
	// +kcc:proto:field=google.cloud.aiplatform.v1.Tool.computer_use
	ComputerUse *Tool_ComputerUse `json:"computerUse,omitempty"`
}

// +kcc:proto=google.cloud.aiplatform.v1.Retrieval
type Retrieval struct {
	// Set to use data source powered by Vertex AI Search.
	// +kcc:proto:field=google.cloud.aiplatform.v1.Retrieval.vertex_ai_search
	VertexAiSearch *VertexAiSearch `json:"vertexAiSearch,omitempty"`

	// Set to use data source powered by Vertex RAG store.
	//  User data is uploaded via the VertexRagDataService.
	// +kcc:proto:field=google.cloud.aiplatform.v1.Retrieval.vertex_rag_store
	VertexRagStore *VertexRagStore `json:"vertexRagStore,omitempty"`
}

// +kcc:proto=google.cloud.aiplatform.v1.VertexRagStore
type VertexRagStore struct {
	// Optional. The representation of the rag source. It can be used to specify
	//  corpus only or ragfiles. Currently only support one corpus or multiple
	//  files from one corpus. In the future we may open up multiple corpora
	//  support.
	// +kcc:proto:field=google.cloud.aiplatform.v1.VertexRagStore.rag_resources
	RagResources []VertexRagStore_RagResource `json:"ragResources,omitempty"`

	// Optional. The retrieval config for the Rag query.
	// +kcc:proto:field=google.cloud.aiplatform.v1.VertexRagStore.rag_retrieval_config
	RagRetrievalConfig *RagRetrievalConfig `json:"ragRetrievalConfig,omitempty"`
}
