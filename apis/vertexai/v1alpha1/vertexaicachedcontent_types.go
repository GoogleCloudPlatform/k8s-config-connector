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
	aiplatformv1alpha1 "github.com/GoogleCloudPlatform/k8s-config-connector/apis/aiplatform/v1alpha1"
	discoveryenginev1alpha1 "github.com/GoogleCloudPlatform/k8s-config-connector/apis/discoveryengine/v1alpha1"
	refsv1beta1 "github.com/GoogleCloudPlatform/k8s-config-connector/apis/refs/v1beta1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/apis/k8s/v1alpha1"
	apiextensionsv1 "k8s.io/apiextensions-apiserver/pkg/apis/apiextensions/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

var VertexAICachedContentGVK = GroupVersion.WithKind("VertexAICachedContent")

// VertexAICachedContentSpec defines the desired state of VertexAICachedContent
// +kcc:spec:proto=google.cloud.aiplatform.v1beta1.CachedContent
type VertexAICachedContentSpec struct {
	// The project that this resource belongs to.
	// +required
	ProjectRef *refsv1beta1.ProjectRef `json:"projectRef"`

	// The location of this resource.
	// +required
	Location *string `json:"location"`

	// The VertexAICachedContent name. If not given, the metadata.name will be used.
	ResourceID *string `json:"resourceID,omitempty"`

	// Optional. Immutable. The user-generated meaningful display name of the
	//  cached content.
	// +kcc:proto:field=google.cloud.aiplatform.v1beta1.CachedContent.display_name
	DisplayName *string `json:"displayName,omitempty"`

	// Immutable. The name of the `Model` to use for cached content. Currently,
	//  only the published Gemini base models are supported, in form of
	//  projects/{PROJECT}/locations/{LOCATION}/publishers/google/models/{MODEL}
	// +kcc:proto:field=google.cloud.aiplatform.v1beta1.CachedContent.model
	// +required
	ModelRef *aiplatformv1alpha1.AIPlatformModelRef `json:"modelRef"`

	// Optional. Input only. Immutable. Developer set system instruction.
	//  Currently, text only
	// +kcc:proto:field=google.cloud.aiplatform.v1beta1.CachedContent.system_instruction
	SystemInstruction *Content `json:"systemInstruction,omitempty"`

	// Optional. Input only. Immutable. The content to cache
	// +kcc:proto:field=google.cloud.aiplatform.v1beta1.CachedContent.contents
	Contents []Content `json:"contents,omitempty"`

	// Optional. Input only. Immutable. A list of `Tools` the model may use to
	//  generate the next response
	// +kcc:proto:field=google.cloud.aiplatform.v1beta1.CachedContent.tools
	Tools []Tool `json:"tools,omitempty"`

	// Optional. Input only. Immutable. Tool config. This config is shared for all
	//  tools
	// +kcc:proto:field=google.cloud.aiplatform.v1beta1.CachedContent.tool_config
	ToolConfig *ToolConfig `json:"toolConfig,omitempty"`

	// Input only. Immutable. Customer-managed encryption key spec for a
	//  `CachedContent`. If set, this `CachedContent` and all its sub-resources
	//  will be secured by this key.
	// +kcc:proto:field=google.cloud.aiplatform.v1beta1.CachedContent.encryption_spec
	EncryptionSpec *EncryptionSpec `json:"encryptionSpec,omitempty"`

	// Timestamp of when this resource is considered expired.
	//  This is *always* provided on output, regardless of what was sent
	//  on input.
	// +kcc:proto:field=google.cloud.aiplatform.v1beta1.CachedContent.expire_time
	ExpireTime *string `json:"expireTime,omitempty"`

	// Input only. The TTL for this resource. The expiration time is computed:
	//  now + TTL.
	// +kcc:proto:field=google.cloud.aiplatform.v1beta1.CachedContent.ttl
	TTL *string `json:"ttl,omitempty"`
}

// VertexAICachedContentStatus defines the config connector machine state of VertexAICachedContent
type VertexAICachedContentStatus struct {
	/* Conditions represent the latest available observations of the
	   object's current state. */
	Conditions []v1alpha1.Condition `json:"conditions,omitempty"`

	// ObservedGeneration is the generation of the resource that was most recently observed by the Config Connector controller. If this is equal to metadata.generation, then that means that the current reported status reflects the most recent desired state of the resource.
	ObservedGeneration *int64 `json:"observedGeneration,omitempty"`

	// A unique specifier for the VertexAICachedContent resource in GCP.
	ExternalRef *string `json:"externalRef,omitempty"`

	// ObservedState is the state of the resource as most recently observed in GCP.
	ObservedState *VertexAICachedContentObservedState `json:"observedState,omitempty"`
}

// VertexAICachedContentObservedState is the state of the VertexAICachedContent resource as most recently observed in GCP.
// +kcc:observedstate:proto=google.cloud.aiplatform.v1beta1.CachedContent
type VertexAICachedContentObservedState struct {
	// Output only. Creation time of the cache entry.
	// +kcc:proto:field=google.cloud.aiplatform.v1beta1.CachedContent.create_time
	CreateTime *string `json:"createTime,omitempty"`

	// Output only. When the cache entry was last updated in UTC time.
	// +kcc:proto:field=google.cloud.aiplatform.v1beta1.CachedContent.update_time
	UpdateTime *string `json:"updateTime,omitempty"`

	// Output only. Metadata on the usage of the cached content.
	// +kcc:proto:field=google.cloud.aiplatform.v1beta1.CachedContent.usage_metadata
	UsageMetadata *CachedContent_UsageMetadata `json:"usageMetadata,omitempty"`
}

// +genclient
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:resource:categories=gcp,shortName=gcpvertexaicachedcontent;gcpvertexaicachedcontents
// +kubebuilder:subresource:status
// +kubebuilder:metadata:labels="cnrm.cloud.google.com/managed-by-kcc=true"
// +kubebuilder:metadata:labels="cnrm.cloud.google.com/system=true"
// +kubebuilder:metadata:labels="cnrm.cloud.google.com/stability-level=alpha"
// +kubebuilder:printcolumn:name="Age",JSONPath=".metadata.creationTimestamp",type="date"
// +kubebuilder:printcolumn:name="Ready",JSONPath=".status.conditions[?(@.type=='Ready')].status",type="string",description="When 'True', the most recent reconcile of the resource succeeded"
// +kubebuilder:printcolumn:name="Status",JSONPath=".status.conditions[?(@.type=='Ready')].reason",type="string",description="The reason for the value in 'Ready'"
// +kubebuilder:printcolumn:name="Status Age",JSONPath=".status.conditions[?(@.type=='Ready')].lastTransitionTime",type="date",description="The last transition time for the value in 'Status'"

// VertexAICachedContent is the Schema for the VertexAICachedContent API
// +k8s:openapi-gen=true
type VertexAICachedContent struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	// +required
	Spec   VertexAICachedContentSpec   `json:"spec,omitempty"`
	Status VertexAICachedContentStatus `json:"status,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// VertexAICachedContentList contains a list of VertexAICachedContent
type VertexAICachedContentList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []VertexAICachedContent `json:"items"`
}

func init() {
	SchemeBuilder.Register(&VertexAICachedContent{}, &VertexAICachedContentList{})
}

// VertexAIRagCorpusRef is a reference to a VertexAIRagCorpus.
type VertexAIRagCorpusRef struct {
	// A reference to an externally managed VertexAIRagCorpus resource. Should be in the format "projects/{{projectID}}/locations/{{location}}/ragCorpora/{{ragCorpus}}".
	External string `json:"external,omitempty"`

	// The name of a VertexAIRagCorpus resource.
	Name string `json:"name,omitempty"`

	// The namespace of a VertexAIRagCorpus resource.
	Namespace string `json:"namespace,omitempty"`
}

// +kcc:proto=google.cloud.aiplatform.v1beta1.VertexAISearch
type VertexAiSearch struct {
	// Optional. Fully-qualified Vertex AI Search data store resource ID.
	//  Format:
	//  `projects/{project}/locations/{location}/collections/{collection}/dataStores/{dataStore}`
	// +kcc:proto:field=google.cloud.aiplatform.v1beta1.VertexAISearch.datastore
	DatastoreRef *discoveryenginev1alpha1.DiscoveryEngineDataStoreRef `json:"datastoreRef,omitempty"`

	// Optional. Fully-qualified Vertex AI Search engine resource ID.
	//  Format:
	//  `projects/{project}/locations/{location}/collections/{collection}/engines/{engine}`
	// +kcc:proto:field=google.cloud.aiplatform.v1beta1.VertexAISearch.engine
	EngineRef *discoveryenginev1alpha1.DiscoveryEngineEngineRef `json:"engineRef,omitempty"`

	// Optional. Number of search results to return per query.
	//  The default value is 10.
	//  The maximumm allowed value is 10.
	// +kcc:proto:field=google.cloud.aiplatform.v1beta1.VertexAISearch.max_results
	MaxResults *int32 `json:"maxResults,omitempty"`

	// Optional. Filter strings to be passed to the search API.
	// +kcc:proto:field=google.cloud.aiplatform.v1beta1.VertexAISearch.filter
	Filter *string `json:"filter,omitempty"`

	// Specifications that define the specific DataStores to be searched, along
	//  with configurations for those data stores. This is only considered for
	//  Engines with multiple data stores.
	//  It should only be set if engine is used.
	// +kcc:proto:field=google.cloud.aiplatform.v1beta1.VertexAISearch.data_store_specs
	DataStoreSpecs []VertexAiSearch_DataStoreSpec `json:"dataStoreSpecs,omitempty"`
}

// +kcc:proto=google.cloud.aiplatform.v1beta1.VertexAISearch.DataStoreSpec
type VertexAiSearch_DataStoreSpec struct {
	// Full resource name of DataStore, such as
	//  Format:
	//  `projects/{project}/locations/{location}/collections/{collection}/dataStores/{dataStore}`
	// +kcc:proto:field=google.cloud.aiplatform.v1beta1.VertexAISearch.DataStoreSpec.data_store
	DataStoreRef *discoveryenginev1alpha1.DiscoveryEngineDataStoreRef `json:"dataStoreRef,omitempty"`

	// Optional. Filter specification to filter documents in the data store
	//  specified by data_store field. For more information on filtering, see
	//  [Filtering](https://cloud.google.com/generative-ai-app-builder/docs/filter-search-metadata)
	// +kcc:proto:field=google.cloud.aiplatform.v1beta1.VertexAISearch.DataStoreSpec.filter
	Filter *string `json:"filter,omitempty"`
}

// +kcc:proto=google.cloud.aiplatform.v1beta1.VertexRagStore.RagResource
type VertexRagStore_RagResource struct {
	// Optional. RagCorpora resource name.
	//  Format:
	//  `projects/{project}/locations/{location}/ragCorpora/{rag_corpus}`
	// +kcc:proto:field=google.cloud.aiplatform.v1beta1.VertexRagStore.RagResource.rag_corpus
	RagCorpusRef *VertexAIRagCorpusRef `json:"ragCorpusRef,omitempty"`

	// Optional. rag_file_id. The files should be in the same rag_corpus set in
	//  rag_corpus field.
	// +kcc:proto:field=google.cloud.aiplatform.v1beta1.VertexRagStore.RagResource.rag_file_ids
	RagFileIDs []string `json:"ragFileIDs,omitempty"`
}

// +kcc:proto=google.cloud.aiplatform.v1beta1.Content
type Content struct {
	// Optional. The producer of the content. Must be either 'user' or 'model'.
	//
	//  Useful to set for multi-turn conversations, otherwise can be left blank
	//  or unset.
	// +kcc:proto:field=google.cloud.aiplatform.v1beta1.Content.role
	Role *string `json:"role,omitempty"`

	// Required. Ordered `Parts` that constitute a single message. Parts may have
	//  different IANA MIME types.
	// +kcc:proto:field=google.cloud.aiplatform.v1beta1.Content.parts
	Parts []Part `json:"parts,omitempty"`
}

// +kcc:proto=google.cloud.aiplatform.v1beta1.Part
type Part struct {
	// Optional. Text part (can be code).
	// +kcc:proto:field=google.cloud.aiplatform.v1beta1.Part.text
	Text *string `json:"text,omitempty"`

	// Optional. Inlined bytes data.
	// +kcc:proto:field=google.cloud.aiplatform.v1beta1.Part.inline_data
	InlineData *Blob `json:"inlineData,omitempty"`

	// Optional. URI based data.
	// +kcc:proto:field=google.cloud.aiplatform.v1beta1.Part.file_data
	FileData *FileData `json:"fileData,omitempty"`

	// Optional. A predicted [FunctionCall] returned from the model that
	//  contains a string representing the [FunctionDeclaration.name] with the
	//  parameters and their values.
	// +kcc:proto:field=google.cloud.aiplatform.v1beta1.Part.function_call
	FunctionCall *FunctionCall `json:"functionCall,omitempty"`

	// Optional. The result output of a [FunctionCall] that contains a string
	//  representing the [FunctionDeclaration.name] and a structured JSON object
	//  containing any output from the function call. It is used as context to
	//  the model.
	// +kcc:proto:field=google.cloud.aiplatform.v1beta1.Part.function_response
	FunctionResponse *FunctionResponse `json:"functionResponse,omitempty"`

	// Optional. Code generated by the model that is meant to be executed.
	// +kcc:proto:field=google.cloud.aiplatform.v1beta1.Part.executable_code
	ExecutableCode *ExecutableCode `json:"executableCode,omitempty"`

	// Optional. Result of executing the [ExecutableCode].
	// +kcc:proto:field=google.cloud.aiplatform.v1beta1.Part.code_execution_result
	CodeExecutionResult *CodeExecutionResult `json:"codeExecutionResult,omitempty"`

	// Optional. Video metadata. The metadata should only be specified while the
	//  video data is presented in inline_data or file_data.
	// +kcc:proto:field=google.cloud.aiplatform.v1beta1.Part.video_metadata
	VideoMetadata *VideoMetadata `json:"videoMetadata,omitempty"`

	// Indicates if the part is thought from the model.
	// +kcc:proto:field=google.cloud.aiplatform.v1beta1.Part.thought
	Thought *bool `json:"thought,omitempty"`

	// An opaque signature for the thought so it can be reused in subsequent
	//  requests.
	// +kcc:proto:field=google.cloud.aiplatform.v1beta1.Part.thought_signature
	ThoughtSignature []byte `json:"thoughtSignature,omitempty"`
}

// +kcc:proto=google.cloud.aiplatform.v1beta1.Blob
type Blob struct {
	// Required. The IANA standard MIME type of the source data.
	// +kcc:proto:field=google.cloud.aiplatform.v1beta1.Blob.mime_type
	MimeType *string `json:"mimeType,omitempty"`

	// Required. Raw bytes.
	// +kcc:proto:field=google.cloud.aiplatform.v1beta1.Blob.data
	Data []byte `json:"data,omitempty"`
}

// +kcc:proto=google.cloud.aiplatform.v1beta1.FileData
type FileData struct {
	// Required. The IANA standard MIME type of the source data.
	// +kcc:proto:field=google.cloud.aiplatform.v1beta1.FileData.mime_type
	MimeType *string `json:"mimeType,omitempty"`

	// Required. URI.
	// +kcc:proto:field=google.cloud.aiplatform.v1beta1.FileData.file_uri
	FileURI *string `json:"fileURI,omitempty"`
}

// +kcc:proto=google.cloud.aiplatform.v1beta1.FunctionCall
type FunctionCall struct {
	// Optional. The unique id of the function call. If populated, the client to
	//  execute the `function_call` and return the response with the matching `id`.
	// +kcc:proto:field=google.cloud.aiplatform.v1beta1.FunctionCall.id
	ID *string `json:"id,omitempty"`

	// Required. The name of the function to call.
	//  Matches [FunctionDeclaration.name].
	// +kcc:proto:field=google.cloud.aiplatform.v1beta1.FunctionCall.name
	Name *string `json:"name,omitempty"`

	// Optional. Required. The function parameters and values in JSON object
	//  format. See [FunctionDeclaration.parameters] for parameter details.
	// +kcc:proto:field=google.cloud.aiplatform.v1beta1.FunctionCall.args
	Args apiextensionsv1.JSON `json:"args,omitempty"`
}

// +kcc:proto=google.cloud.aiplatform.v1beta1.FunctionResponse
type FunctionResponse struct {
	// Optional. The id of the function call this response is for. Populated by
	//  the client to match the corresponding function call `id`.
	// +kcc:proto:field=google.cloud.aiplatform.v1beta1.FunctionResponse.id
	ID *string `json:"id,omitempty"`

	// Required. The name of the function to call.
	//  Matches [FunctionDeclaration.name] and [FunctionCall.name].
	// +kcc:proto:field=google.cloud.aiplatform.v1beta1.FunctionResponse.name
	Name *string `json:"name,omitempty"`

	// Required. The function response in JSON object format.
	//  Use "output" key to specify function output and "error" key to specify
	//  error details (if any). If "output" and "error" keys are not specified,
	//  then whole "response" is treated as function output.
	// +kcc:proto:field=google.cloud.aiplatform.v1beta1.FunctionResponse.response
	Response apiextensionsv1.JSON `json:"response,omitempty"`
}

// +kcc:proto=google.cloud.aiplatform.v1beta1.ExecutableCode
type ExecutableCode struct {
	// Required. Programming language of the `code`.
	// +kcc:proto:field=google.cloud.aiplatform.v1beta1.ExecutableCode.language
	Language *string `json:"language,omitempty"`

	// Required. The code to be executed.
	// +kcc:proto:field=google.cloud.aiplatform.v1beta1.ExecutableCode.code
	Code *string `json:"code,omitempty"`
}

// +kcc:proto=google.cloud.aiplatform.v1beta1.CodeExecutionResult
type CodeExecutionResult struct {
	// Required. Outcome of the code execution.
	// +kcc:proto:field=google.cloud.aiplatform.v1beta1.CodeExecutionResult.outcome
	Outcome *string `json:"outcome,omitempty"`

	// Optional. Contains stdout when code execution is successful, stderr or
	//  other description otherwise.
	// +kcc:proto:field=google.cloud.aiplatform.v1beta1.CodeExecutionResult.output
	Output *string `json:"output,omitempty"`
}

// +kcc:proto=google.cloud.aiplatform.v1beta1.VideoMetadata
type VideoMetadata struct {
	// Optional. The start offset of the video.
	// +kcc:proto:field=google.cloud.aiplatform.v1beta1.VideoMetadata.start_offset
	StartOffset *string `json:"startOffset,omitempty"`

	// Optional. The end offset of the video.
	// +kcc:proto:field=google.cloud.aiplatform.v1beta1.VideoMetadata.end_offset
	EndOffset *string `json:"endOffset,omitempty"`
}

// +kcc:proto=google.cloud.aiplatform.v1beta1.Tool
type Tool struct {
	// Optional. Function tool type.
	//  One or more function declarations to be passed to the model along with the
	//  current user query. Model may decide to call a subset of these functions
	//  by populating
	//  [FunctionCall][google.cloud.aiplatform.v1beta1.Part.function_call] in the
	//  response. User should provide a
	//  [FunctionResponse][google.cloud.aiplatform.v1beta1.Part.function_response]
	//  for each function call in the next turn. Based on the function responses,
	//  Model will generate the final response back to the user.
	//  Maximum 128 function declarations can be provided.
	// +kcc:proto:field=google.cloud.aiplatform.v1beta1.Tool.function_declarations
	FunctionDeclarations []FunctionDeclaration `json:"functionDeclarations,omitempty"`

	// Optional. Retrieval tool type.
	//  System will always execute the provided retrieval tool(s) to get external
	//  knowledge to answer the prompt. Retrieval results are presented to the
	//  model for generation.
	// +kcc:proto:field=google.cloud.aiplatform.v1beta1.Tool.retrieval
	Retrieval *Retrieval `json:"retrieval,omitempty"`

	// Optional. GoogleSearch tool type.
	//  Tool to support Google Search in Model. Powered by Google.
	// +kcc:proto:field=google.cloud.aiplatform.v1beta1.Tool.google_search
	GoogleSearch *Tool_GoogleSearch `json:"googleSearch,omitempty"`

	// Optional. GoogleSearchRetrieval tool type.
	//  Specialized retrieval tool that is powered by Google search.
	// +kcc:proto:field=google.cloud.aiplatform.v1beta1.Tool.google_search_retrieval
	GoogleSearchRetrieval *GoogleSearchRetrieval `json:"googleSearchRetrieval,omitempty"`

	// Optional. GoogleMaps tool type.
	//  Tool to support Google Maps in Model.
	// +kcc:proto:field=google.cloud.aiplatform.v1beta1.Tool.google_maps
	GoogleMaps *GoogleMaps `json:"googleMaps,omitempty"`

	// Optional. Tool to support searching public web data, powered by Vertex AI
	//  Search and Sec4 compliance.
	// +kcc:proto:field=google.cloud.aiplatform.v1beta1.Tool.enterprise_web_search
	EnterpriseWebSearch *EnterpriseWebSearch `json:"enterpriseWebSearch,omitempty"`

	// Optional. CodeExecution tool type.
	//  Enables the model to execute code as part of generation.
	// +kcc:proto:field=google.cloud.aiplatform.v1beta1.Tool.code_execution
	CodeExecution *Tool_CodeExecution `json:"codeExecution,omitempty"`

	// Optional. Tool to support URL context retrieval.
	// +kcc:proto:field=google.cloud.aiplatform.v1beta1.Tool.url_context
	URLContext *URLContext `json:"urlContext,omitempty"`

	// Optional. Tool to support computer use.
	// +kcc:proto:field=google.cloud.aiplatform.v1beta1.Tool.computer_use
	ComputerUse *Tool_ComputerUse `json:"computerUse,omitempty"`
}

// +kcc:proto=google.cloud.aiplatform.v1beta1.Tool.CodeExecution
type Tool_CodeExecution struct {
}

// +kcc:proto=google.cloud.aiplatform.v1beta1.Tool.GoogleSearch
type Tool_GoogleSearch struct {
	// Optional. List of domains to be excluded from the search results.
	//  The default limit is 2000 domains.
	//  Example: ["amazon.com", "facebook.com"].
	// +kcc:proto:field=google.cloud.aiplatform.v1beta1.Tool.google_search.exclude_domains
	ExcludeDomains []string `json:"excludeDomains,omitempty"`
}

// +kcc:proto=google.cloud.aiplatform.v1beta1.FunctionDeclaration
type FunctionDeclaration struct {
	// Required. The name of the function to call.
	//  Must start with a letter or an underscore.
	//  Must be a-z, A-Z, 0-9, or contain underscores, dots and dashes, with a
	//  maximum length of 64.
	// +kcc:proto:field=google.cloud.aiplatform.v1beta1.FunctionDeclaration.name
	Name *string `json:"name,omitempty"`

	// Optional. Description and purpose of the function.
	//  Model uses it to decide how and whether to call the function.
	// +kcc:proto:field=google.cloud.aiplatform.v1beta1.FunctionDeclaration.description
	Description *string `json:"description,omitempty"`

	// Optional. Describes the parameters to this function in JSON Schema Object
	//  format. Reflects the Open API 3.03 Parameter Object. string Key: the name
	//  of the parameter. Parameter names are case sensitive. Schema Value: the
	//  Schema defining the type used for the parameter. For function with no
	//  parameters, this can be left unset. Parameter names must start with a
	//  letter or an underscore and must only contain chars a-z, A-Z, 0-9, or
	//  underscores with a maximum length of 64. Example with 1 required and 1
	//  optional parameter: type: OBJECT properties:
	//   param1:
	//     type: STRING
	//   param2:
	//     type: INTEGER
	//  required:
	//   - param1
	// +kcc:proto:field=google.cloud.aiplatform.v1beta1.FunctionDeclaration.parameters
	Parameters apiextensionsv1.JSON `json:"parameters,omitempty"`

	// Optional. Describes the parameters to the function in JSON Schema format.
	//  The schema must describe an object where the properties are the parameters
	//  to the function.
	// +kcc:proto:field=google.cloud.aiplatform.v1beta1.FunctionDeclaration.parameters_json_schema
	ParametersJsonSchema apiextensionsv1.JSON `json:"parametersJsonSchema,omitempty"`

	// Optional. Describes the output from this function in JSON Schema format.
	//  Reflects the Open API 3.03 Response Object. The Schema defines the type
	//  used for the response value of the function.
	// +kcc:proto:field=google.cloud.aiplatform.v1beta1.FunctionDeclaration.response
	Response apiextensionsv1.JSON `json:"response,omitempty"`

	// Optional. Describes the output from this function in JSON Schema format.
	//  The value specified by the schema is the response value of the function.
	//
	//  This field is mutually exclusive with `response`.
	// +kcc:proto:field=google.cloud.aiplatform.v1beta1.FunctionDeclaration.response_json_schema
	ResponseJsonSchema apiextensionsv1.JSON `json:"responseJsonSchema,omitempty"`
}

// +kcc:proto=google.cloud.aiplatform.v1beta1.Retrieval
type Retrieval struct {
	// Set to use data source powered by Vertex AI Search.
	// +kcc:proto:field=google.cloud.aiplatform.v1beta1.Retrieval.vertex_ai_search
	VertexAiSearch *VertexAiSearch `json:"vertexAiSearch,omitempty"`

	// Set to use data source powered by Vertex RAG store.
	//  User data is uploaded via the VertexRagDataService.
	// +kcc:proto:field=google.cloud.aiplatform.v1beta1.Retrieval.vertex_rag_store
	VertexRagStore *VertexRagStore `json:"vertexRagStore,omitempty"`

	// Optional. Deprecated. This option is no longer supported.
	// +kcc:proto:field=google.cloud.aiplatform.v1beta1.Retrieval.disable_attribution
	DisableAttribution *bool `json:"disableAttribution,omitempty"`
}

// +kcc:proto=google.cloud.aiplatform.v1beta1.GoogleSearchRetrieval
type GoogleSearchRetrieval struct {
	// Specifies the dynamic retrieval configuration for the given source.
	// +kcc:proto:field=google.cloud.aiplatform.v1beta1.GoogleSearchRetrieval.dynamic_retrieval_config
	DynamicRetrievalConfig *DynamicRetrievalConfig `json:"dynamicRetrievalConfig,omitempty"`
}

// +kcc:proto=google.cloud.aiplatform.v1beta1.DynamicRetrievalConfig
type DynamicRetrievalConfig struct {
	// The mode of the predictor to be used in dynamic retrieval.
	// +kcc:proto:field=google.cloud.aiplatform.v1beta1.DynamicRetrievalConfig.mode
	Mode *string `json:"mode,omitempty"`

	// Optional. The threshold to be used in dynamic retrieval.
	//  If not set, a system default value is used.
	// +kcc:proto:field=google.cloud.aiplatform.v1beta1.DynamicRetrievalConfig.dynamic_threshold
	DynamicThreshold *float32 `json:"dynamicThreshold,omitempty"`
}

// +kcc:proto=google.cloud.aiplatform.v1beta1.ToolConfig
type ToolConfig struct {
	// Optional. Function calling config.
	// +kcc:proto:field=google.cloud.aiplatform.v1beta1.ToolConfig.function_calling_config
	FunctionCallingConfig *FunctionCallingConfig `json:"functionCallingConfig,omitempty"`

	// Optional. Retrieval config.
	// +kcc:proto:field=google.cloud.aiplatform.v1beta1.ToolConfig.retrieval_config
	RetrievalConfig *RetrievalConfig `json:"retrievalConfig,omitempty"`
}

// +kcc:proto=google.cloud.aiplatform.v1beta1.FunctionCallingConfig
type FunctionCallingConfig struct {
	// Optional. Function calling mode.
	// +kcc:proto:field=google.cloud.aiplatform.v1beta1.FunctionCallingConfig.mode
	Mode *string `json:"mode,omitempty"`

	// Optional. Function names to call. Only set when the Mode is ANY. Function
	//  names should match [FunctionDeclaration.name]. With mode set to ANY, model
	//  will predict a function call from the set of function names provided.
	// +kcc:proto:field=google.cloud.aiplatform.v1beta1.FunctionCallingConfig.allowed_function_names
	AllowedFunctionNames []string `json:"allowedFunctionNames,omitempty"`
}

// +kcc:proto=google.cloud.aiplatform.v1beta1.RetrievalConfig
type RetrievalConfig struct {
	// The location of the user.
	// +kcc:proto:field=google.cloud.aiplatform.v1beta1.RetrievalConfig.lat_lng
	LatLng *LatLng `json:"latLng,omitempty"`

	// The language code of the user.
	// +kcc:proto:field=google.cloud.aiplatform.v1beta1.RetrievalConfig.language_code
	LanguageCode *string `json:"languageCode,omitempty"`
}

// +kcc:proto=google.cloud.aiplatform.v1beta1.CachedContent.UsageMetadata
type CachedContent_UsageMetadata struct {
	// Total number of tokens that the cached content consumes.
	// +kcc:proto:field=google.cloud.aiplatform.v1beta1.CachedContent.UsageMetadata.total_token_count
	TotalTokenCount *int32 `json:"totalTokenCount,omitempty"`

	// Number of text characters.
	// +kcc:proto:field=google.cloud.aiplatform.v1beta1.CachedContent.UsageMetadata.text_count
	TextCount *int32 `json:"textCount,omitempty"`

	// Number of images.
	// +kcc:proto:field=google.cloud.aiplatform.v1beta1.CachedContent.UsageMetadata.image_count
	ImageCount *int32 `json:"imageCount,omitempty"`

	// Duration of video in seconds.
	// +kcc:proto:field=google.cloud.aiplatform.v1beta1.CachedContent.UsageMetadata.video_duration_seconds
	VideoDurationSeconds *int32 `json:"videoDurationSeconds,omitempty"`

	// Duration of audio in seconds.
	// +kcc:proto:field=google.cloud.aiplatform.v1beta1.CachedContent.UsageMetadata.audio_duration_seconds
	AudioDurationSeconds *int32 `json:"audioDurationSeconds,omitempty"`
}

// +kcc:proto=google.cloud.aiplatform.v1beta1.UrlContext
type URLContext struct {
}

// +kcc:proto=google.cloud.aiplatform.v1beta1.VertexRagStore
type VertexRagStore struct {
	// Optional. Deprecated. Please use rag_resources instead.
	// +kcc:proto:field=google.cloud.aiplatform.v1beta1.VertexRagStore.rag_corpora
	RagCorpora []string `json:"ragCorpora,omitempty"`

	// Optional. The representation of the rag source. It can be used to specify
	//  corpus only or ragfiles. Currently only support one corpus or multiple
	//  files from one corpus. In the future we may open up multiple corpora
	//  support.
	// +kcc:proto:field=google.cloud.aiplatform.v1beta1.VertexRagStore.rag_resources
	RagResources []VertexRagStore_RagResource `json:"ragResources,omitempty"`

	// Optional. Number of top k results to return from the selected corpora.
	// +kcc:proto:field=google.cloud.aiplatform.v1beta1.VertexRagStore.similarity_top_k
	SimilarityTopK *int32 `json:"similarityTopK,omitempty"`

	// Optional. Only return results with vector distance smaller than the
	//  threshold.
	// +kcc:proto:field=google.cloud.aiplatform.v1beta1.VertexRagStore.vector_distance_threshold
	VectorDistanceThreshold *float64 `json:"vectorDistanceThreshold,omitempty"`

	// Optional. The retrieval config for the Rag query.
	// +kcc:proto:field=google.cloud.aiplatform.v1beta1.VertexRagStore.rag_retrieval_config
	RagRetrievalConfig *RagRetrievalConfig `json:"ragRetrievalConfig,omitempty"`

	// Optional. Currently only supported for Gemini Multimodal Live API.
	// +kcc:proto:field=google.cloud.aiplatform.v1beta1.VertexRagStore.store_context
	StoreContext *bool `json:"storeContext,omitempty"`
}

// +kcc:proto=google.cloud.aiplatform.v1beta1.GoogleMaps
type GoogleMaps struct {
}

// +kcc:proto=google.cloud.aiplatform.v1beta1.EnterpriseWebSearch
type EnterpriseWebSearch struct {
	// Optional. List of domains to be excluded from the search results.
	//  The default limit is 2000 domains.
	// +kcc:proto:field=google.cloud.aiplatform.v1beta1.EnterpriseWebSearch.exclude_domains
	ExcludeDomains []string `json:"excludeDomains,omitempty"`
}

// +kcc:proto=google.type.LatLng
type LatLng struct {
	// The latitude in degrees. It must be in the range [-90.0, +90.0].
	// +kcc:proto:field=google.type.LatLng.latitude
	Latitude *float64 `json:"latitude,omitempty"`

	// The longitude in degrees. It must be in the range [-180.0, +180.0].
	// +kcc:proto:field=google.type.LatLng.longitude
	Longitude *float64 `json:"longitude,omitempty"`
}

// +kcc:proto=google.cloud.aiplatform.v1beta1.RagRetrievalConfig
type RagRetrievalConfig struct {
	// Optional. The number of contexts to retrieve.
	// +kcc:proto:field=google.cloud.aiplatform.v1beta1.RagRetrievalConfig.top_k
	TopK *int32 `json:"topK,omitempty"`

	// Optional. Config for hybrid search.
	// +kcc:proto:field=google.cloud.aiplatform.v1beta1.RagRetrievalConfig.hybrid_search
	HybridSearch *RagRetrievalConfig_HybridSearch `json:"hybridSearch,omitempty"`

	// Optional. Config for filters.
	// +kcc:proto:field=google.cloud.aiplatform.v1beta1.RagRetrievalConfig.filter
	Filter *RagRetrievalConfig_Filter `json:"filter,omitempty"`

	// Optional. Config for ranking.
	// +kcc:proto:field=google.cloud.aiplatform.v1beta1.RagRetrievalConfig.ranking
	Ranking *RagRetrievalConfig_Ranking `json:"ranking,omitempty"`
}

// +kcc:proto=google.cloud.aiplatform.v1beta1.RagRetrievalConfig.HybridSearch
type RagRetrievalConfig_HybridSearch struct {
	// Optional. The alpha value of the hybrid search.
	// +kcc:proto:field=google.cloud.aiplatform.v1beta1.RagRetrievalConfig.HybridSearch.alpha
	Alpha *float32 `json:"alpha,omitempty"`
}

// +kcc:proto=google.cloud.aiplatform.v1beta1.RagRetrievalConfig.Filter
type RagRetrievalConfig_Filter struct {
	// Optional. The vector distance threshold.
	// +kcc:proto:field=google.cloud.aiplatform.v1beta1.RagRetrievalConfig.Filter.vector_distance_threshold
	VectorDistanceThreshold *float64 `json:"vectorDistanceThreshold,omitempty"`

	// Optional. The vector similarity threshold.
	// +kcc:proto:field=google.cloud.aiplatform.v1beta1.RagRetrievalConfig.Filter.vector_similarity_threshold
	VectorSimilarityThreshold *float64 `json:"vectorSimilarityThreshold,omitempty"`

	// Optional. The metadata filter.
	// +kcc:proto:field=google.cloud.aiplatform.v1beta1.RagRetrievalConfig.Filter.metadata_filter
	MetadataFilter *string `json:"metadataFilter,omitempty"`
}

// +kcc:proto=google.cloud.aiplatform.v1beta1.RagRetrievalConfig.Ranking
type RagRetrievalConfig_Ranking struct {
	// Optional. The rank service.
	// +kcc:proto:field=google.cloud.aiplatform.v1beta1.RagRetrievalConfig.Ranking.rank_service
	RankService *RagRetrievalConfig_Ranking_RankService `json:"rankService,omitempty"`

	// Optional. The LLM ranker.
	// +kcc:proto:field=google.cloud.aiplatform.v1beta1.RagRetrievalConfig.Ranking.llm_ranker
	LlmRanker *RagRetrievalConfig_Ranking_LlmRanker `json:"llmRanker,omitempty"`
}

// +kcc:proto=google.cloud.aiplatform.v1beta1.RagRetrievalConfig.Ranking.RankService
type RagRetrievalConfig_Ranking_RankService struct {
	// Optional. The model name of the rank service.
	// +kcc:proto:field=google.cloud.aiplatform.v1beta1.RagRetrievalConfig.Ranking.RankService.model_name
	ModelName *string `json:"modelName,omitempty"`
}

// +kcc:proto=google.cloud.aiplatform.v1beta1.RagRetrievalConfig.Ranking.LlmRanker
type RagRetrievalConfig_Ranking_LlmRanker struct {
	// Optional. The model name of the LLM ranker.
	// +kcc:proto:field=google.cloud.aiplatform.v1beta1.RagRetrievalConfig.Ranking.LlmRanker.model_name
	ModelName *string `json:"modelName,omitempty"`
}

// +kcc:proto=google.cloud.aiplatform.v1beta1.Tool.ComputerUse
type Tool_ComputerUse struct {
	// Optional. The environment to use for computer use.
	// +kcc:proto:field=google.cloud.aiplatform.v1beta1.Tool.ComputerUse.environment
	Environment *string `json:"environment,omitempty"`
}
