// Copyright 2025 Google LLC
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


// +kcc:proto=google.cloud.aiplatform.v1.Blob
type Blob struct {
	// Required. The IANA standard MIME type of the source data.
	// +kcc:proto:field=google.cloud.aiplatform.v1.Blob.mime_type
	MimeType *string `json:"mimeType,omitempty"`

	// Required. Raw bytes.
	// +kcc:proto:field=google.cloud.aiplatform.v1.Blob.data
	Data []byte `json:"data,omitempty"`
}

// +kcc:proto=google.cloud.aiplatform.v1.CachedContent
type CachedContent struct {
	// Timestamp of when this resource is considered expired.
	//  This is *always* provided on output, regardless of what was sent
	//  on input.
	// +kcc:proto:field=google.cloud.aiplatform.v1.CachedContent.expire_time
	ExpireTime *string `json:"expireTime,omitempty"`

	// Input only. The TTL for this resource. The expiration time is computed:
	//  now + TTL.
	// +kcc:proto:field=google.cloud.aiplatform.v1.CachedContent.ttl
	Ttl *string `json:"ttl,omitempty"`

	// Immutable. Identifier. The server-generated resource name of the cached
	//  content Format:
	//  projects/{project}/locations/{location}/cachedContents/{cached_content}
	// +kcc:proto:field=google.cloud.aiplatform.v1.CachedContent.name
	Name *string `json:"name,omitempty"`

	// Optional. Immutable. The user-generated meaningful display name of the
	//  cached content.
	// +kcc:proto:field=google.cloud.aiplatform.v1.CachedContent.display_name
	DisplayName *string `json:"displayName,omitempty"`

	// Immutable. The name of the publisher model to use for cached content.
	//  Format:
	//  projects/{project}/locations/{location}/publishers/{publisher}/models/{model}
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
}

// +kcc:proto=google.cloud.aiplatform.v1.CachedContent.UsageMetadata
type CachedContent_UsageMetadata struct {
	// Total number of tokens that the cached content consumes.
	// +kcc:proto:field=google.cloud.aiplatform.v1.CachedContent.UsageMetadata.total_token_count
	TotalTokenCount *int32 `json:"totalTokenCount,omitempty"`

	// Number of text characters.
	// +kcc:proto:field=google.cloud.aiplatform.v1.CachedContent.UsageMetadata.text_count
	TextCount *int32 `json:"textCount,omitempty"`

	// Number of images.
	// +kcc:proto:field=google.cloud.aiplatform.v1.CachedContent.UsageMetadata.image_count
	ImageCount *int32 `json:"imageCount,omitempty"`

	// Duration of video in seconds.
	// +kcc:proto:field=google.cloud.aiplatform.v1.CachedContent.UsageMetadata.video_duration_seconds
	VideoDurationSeconds *int32 `json:"videoDurationSeconds,omitempty"`

	// Duration of audio in seconds.
	// +kcc:proto:field=google.cloud.aiplatform.v1.CachedContent.UsageMetadata.audio_duration_seconds
	AudioDurationSeconds *int32 `json:"audioDurationSeconds,omitempty"`
}

// +kcc:proto=google.cloud.aiplatform.v1.Content
type Content struct {
	// Optional. The producer of the content. Must be either 'user' or 'model'.
	//
	//  Useful to set for multi-turn conversations, otherwise can be left blank
	//  or unset.
	// +kcc:proto:field=google.cloud.aiplatform.v1.Content.role
	Role *string `json:"role,omitempty"`

	// Required. Ordered `Parts` that constitute a single message. Parts may have
	//  different IANA MIME types.
	// +kcc:proto:field=google.cloud.aiplatform.v1.Content.parts
	Parts []Part `json:"parts,omitempty"`
}

// +kcc:proto=google.cloud.aiplatform.v1.DynamicRetrievalConfig
type DynamicRetrievalConfig struct {
	// The mode of the predictor to be used in dynamic retrieval.
	// +kcc:proto:field=google.cloud.aiplatform.v1.DynamicRetrievalConfig.mode
	Mode *string `json:"mode,omitempty"`

	// Optional. The threshold to be used in dynamic retrieval.
	//  If not set, a system default value is used.
	// +kcc:proto:field=google.cloud.aiplatform.v1.DynamicRetrievalConfig.dynamic_threshold
	DynamicThreshold *float32 `json:"dynamicThreshold,omitempty"`
}

// +kcc:proto=google.cloud.aiplatform.v1.FileData
type FileData struct {
	// Required. The IANA standard MIME type of the source data.
	// +kcc:proto:field=google.cloud.aiplatform.v1.FileData.mime_type
	MimeType *string `json:"mimeType,omitempty"`

	// Required. URI.
	// +kcc:proto:field=google.cloud.aiplatform.v1.FileData.file_uri
	FileURI *string `json:"fileURI,omitempty"`
}

// +kcc:proto=google.cloud.aiplatform.v1.FunctionCall
type FunctionCall struct {
	// Required. The name of the function to call.
	//  Matches [FunctionDeclaration.name].
	// +kcc:proto:field=google.cloud.aiplatform.v1.FunctionCall.name
	Name *string `json:"name,omitempty"`

	// Optional. Required. The function parameters and values in JSON object
	//  format. See [FunctionDeclaration.parameters] for parameter details.
	// +kcc:proto:field=google.cloud.aiplatform.v1.FunctionCall.args
	Args map[string]string `json:"args,omitempty"`
}

// +kcc:proto=google.cloud.aiplatform.v1.FunctionCallingConfig
type FunctionCallingConfig struct {
	// Optional. Function calling mode.
	// +kcc:proto:field=google.cloud.aiplatform.v1.FunctionCallingConfig.mode
	Mode *string `json:"mode,omitempty"`

	// Optional. Function names to call. Only set when the Mode is ANY. Function
	//  names should match [FunctionDeclaration.name]. With mode set to ANY, model
	//  will predict a function call from the set of function names provided.
	// +kcc:proto:field=google.cloud.aiplatform.v1.FunctionCallingConfig.allowed_function_names
	AllowedFunctionNames []string `json:"allowedFunctionNames,omitempty"`
}

// +kcc:proto=google.cloud.aiplatform.v1.FunctionDeclaration
type FunctionDeclaration struct {
	// Required. The name of the function to call.
	//  Must start with a letter or an underscore.
	//  Must be a-z, A-Z, 0-9, or contain underscores, dots and dashes, with a
	//  maximum length of 64.
	// +kcc:proto:field=google.cloud.aiplatform.v1.FunctionDeclaration.name
	Name *string `json:"name,omitempty"`

	// Optional. Description and purpose of the function.
	//  Model uses it to decide how and whether to call the function.
	// +kcc:proto:field=google.cloud.aiplatform.v1.FunctionDeclaration.description
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
	// +kcc:proto:field=google.cloud.aiplatform.v1.FunctionDeclaration.parameters
	Parameters *Schema `json:"parameters,omitempty"`

	// Optional. Describes the output from this function in JSON Schema format.
	//  Reflects the Open API 3.03 Response Object. The Schema defines the type
	//  used for the response value of the function.
	// +kcc:proto:field=google.cloud.aiplatform.v1.FunctionDeclaration.response
	Response *Schema `json:"response,omitempty"`
}

// +kcc:proto=google.cloud.aiplatform.v1.FunctionResponse
type FunctionResponse struct {
	// Required. The name of the function to call.
	//  Matches [FunctionDeclaration.name] and [FunctionCall.name].
	// +kcc:proto:field=google.cloud.aiplatform.v1.FunctionResponse.name
	Name *string `json:"name,omitempty"`

	// Required. The function response in JSON object format.
	//  Use "output" key to specify function output and "error" key to specify
	//  error details (if any). If "output" and "error" keys are not specified,
	//  then whole "response" is treated as function output.
	// +kcc:proto:field=google.cloud.aiplatform.v1.FunctionResponse.response
	Response map[string]string `json:"response,omitempty"`
}

// +kcc:proto=google.cloud.aiplatform.v1.GoogleSearchRetrieval
type GoogleSearchRetrieval struct {
	// Specifies the dynamic retrieval configuration for the given source.
	// +kcc:proto:field=google.cloud.aiplatform.v1.GoogleSearchRetrieval.dynamic_retrieval_config
	DynamicRetrievalConfig *DynamicRetrievalConfig `json:"dynamicRetrievalConfig,omitempty"`
}

// +kcc:proto=google.cloud.aiplatform.v1.Part
type Part struct {
	// Optional. Text part (can be code).
	// +kcc:proto:field=google.cloud.aiplatform.v1.Part.text
	Text *string `json:"text,omitempty"`

	// Optional. Inlined bytes data.
	// +kcc:proto:field=google.cloud.aiplatform.v1.Part.inline_data
	InlineData *Blob `json:"inlineData,omitempty"`

	// Optional. URI based data.
	// +kcc:proto:field=google.cloud.aiplatform.v1.Part.file_data
	FileData *FileData `json:"fileData,omitempty"`

	// Optional. A predicted [FunctionCall] returned from the model that
	//  contains a string representing the [FunctionDeclaration.name] with the
	//  parameters and their values.
	// +kcc:proto:field=google.cloud.aiplatform.v1.Part.function_call
	FunctionCall *FunctionCall `json:"functionCall,omitempty"`

	// Optional. The result output of a [FunctionCall] that contains a string
	//  representing the [FunctionDeclaration.name] and a structured JSON object
	//  containing any output from the function call. It is used as context to
	//  the model.
	// +kcc:proto:field=google.cloud.aiplatform.v1.Part.function_response
	FunctionResponse *FunctionResponse `json:"functionResponse,omitempty"`

	// Optional. Video metadata. The metadata should only be specified while the
	//  video data is presented in inline_data or file_data.
	// +kcc:proto:field=google.cloud.aiplatform.v1.Part.video_metadata
	VideoMetadata *VideoMetadata `json:"videoMetadata,omitempty"`
}

// +kcc:proto=google.cloud.aiplatform.v1.RagRetrievalConfig
type RagRetrievalConfig struct {
	// Optional. The number of contexts to retrieve.
	// +kcc:proto:field=google.cloud.aiplatform.v1.RagRetrievalConfig.top_k
	TopK *int32 `json:"topK,omitempty"`

	// Optional. Config for filters.
	// +kcc:proto:field=google.cloud.aiplatform.v1.RagRetrievalConfig.filter
	Filter *RagRetrievalConfig_Filter `json:"filter,omitempty"`
}

// +kcc:proto=google.cloud.aiplatform.v1.RagRetrievalConfig.Filter
type RagRetrievalConfig_Filter struct {
	// Optional. Only returns contexts with vector distance smaller than the
	//  threshold.
	// +kcc:proto:field=google.cloud.aiplatform.v1.RagRetrievalConfig.Filter.vector_distance_threshold
	VectorDistanceThreshold *float64 `json:"vectorDistanceThreshold,omitempty"`

	// Optional. Only returns contexts with vector similarity larger than the
	//  threshold.
	// +kcc:proto:field=google.cloud.aiplatform.v1.RagRetrievalConfig.Filter.vector_similarity_threshold
	VectorSimilarityThreshold *float64 `json:"vectorSimilarityThreshold,omitempty"`

	// Optional. String for metadata filtering.
	// +kcc:proto:field=google.cloud.aiplatform.v1.RagRetrievalConfig.Filter.metadata_filter
	MetadataFilter *string `json:"metadataFilter,omitempty"`
}

// +kcc:proto=google.cloud.aiplatform.v1.Retrieval
type Retrieval struct {
	// Set to use data source powered by Vertex AI Search.
	// +kcc:proto:field=google.cloud.aiplatform.v1.Retrieval.vertex_ai_search
	VertexAiSearch *VertexAISearch `json:"vertexAiSearch,omitempty"`

	// Set to use data source powered by Vertex RAG store.
	//  User data is uploaded via the VertexRagDataService.
	// +kcc:proto:field=google.cloud.aiplatform.v1.Retrieval.vertex_rag_store
	VertexRagStore *VertexRagStore `json:"vertexRagStore,omitempty"`

	// Optional. Deprecated. This option is no longer supported.
	// +kcc:proto:field=google.cloud.aiplatform.v1.Retrieval.disable_attribution
	DisableAttribution *bool `json:"disableAttribution,omitempty"`
}

// +kcc:proto=google.cloud.aiplatform.v1.RetrievalConfig
type RetrievalConfig struct {
	// The location of the user.
	// +kcc:proto:field=google.cloud.aiplatform.v1.RetrievalConfig.lat_lng
	LatLng *LatLng `json:"latLng,omitempty"`

	// The language code of the user.
	// +kcc:proto:field=google.cloud.aiplatform.v1.RetrievalConfig.language_code
	LanguageCode *string `json:"languageCode,omitempty"`
}

// +kcc:proto=google.cloud.aiplatform.v1.Schema
type Schema struct {
	// Optional. The type of the data.
	// +kcc:proto:field=google.cloud.aiplatform.v1.Schema.type
	Type *string `json:"type,omitempty"`

	// Optional. The format of the data.
	//  Supported formats:
	//   for NUMBER type: "float", "double"
	//   for INTEGER type: "int32", "int64"
	//   for STRING type: "email", "byte", etc
	// +kcc:proto:field=google.cloud.aiplatform.v1.Schema.format
	Format *string `json:"format,omitempty"`

	// Optional. The title of the Schema.
	// +kcc:proto:field=google.cloud.aiplatform.v1.Schema.title
	Title *string `json:"title,omitempty"`

	// Optional. The description of the data.
	// +kcc:proto:field=google.cloud.aiplatform.v1.Schema.description
	Description *string `json:"description,omitempty"`

	// Optional. Indicates if the value may be null.
	// +kcc:proto:field=google.cloud.aiplatform.v1.Schema.nullable
	Nullable *bool `json:"nullable,omitempty"`

	// Optional. Default value of the data.
	// +kcc:proto:field=google.cloud.aiplatform.v1.Schema.default
	Default *Value `json:"default,omitempty"`

	// Optional. SCHEMA FIELDS FOR TYPE ARRAY
	//  Schema of the elements of Type.ARRAY.
	// +kcc:proto:field=google.cloud.aiplatform.v1.Schema.items
	Items *Schema `json:"items,omitempty"`

	// Optional. Minimum number of the elements for Type.ARRAY.
	// +kcc:proto:field=google.cloud.aiplatform.v1.Schema.min_items
	MinItems *int64 `json:"minItems,omitempty"`

	// Optional. Maximum number of the elements for Type.ARRAY.
	// +kcc:proto:field=google.cloud.aiplatform.v1.Schema.max_items
	MaxItems *int64 `json:"maxItems,omitempty"`

	// Optional. Possible values of the element of primitive type with enum
	//  format. Examples:
	//  1. We can define direction as :
	//  {type:STRING, format:enum, enum:["EAST", NORTH", "SOUTH", "WEST"]}
	//  2. We can define apartment number as :
	//  {type:INTEGER, format:enum, enum:["101", "201", "301"]}
	// +kcc:proto:field=google.cloud.aiplatform.v1.Schema.enum
	Enum []string `json:"enum,omitempty"`

	// TODO: unsupported map type with key string and value message


	// Optional. The order of the properties.
	//  Not a standard field in open api spec. Only used to support the order of
	//  the properties.
	// +kcc:proto:field=google.cloud.aiplatform.v1.Schema.property_ordering
	PropertyOrdering []string `json:"propertyOrdering,omitempty"`

	// Optional. Required properties of Type.OBJECT.
	// +kcc:proto:field=google.cloud.aiplatform.v1.Schema.required
	Required []string `json:"required,omitempty"`

	// Optional. Minimum number of the properties for Type.OBJECT.
	// +kcc:proto:field=google.cloud.aiplatform.v1.Schema.min_properties
	MinProperties *int64 `json:"minProperties,omitempty"`

	// Optional. Maximum number of the properties for Type.OBJECT.
	// +kcc:proto:field=google.cloud.aiplatform.v1.Schema.max_properties
	MaxProperties *int64 `json:"maxProperties,omitempty"`

	// Optional. SCHEMA FIELDS FOR TYPE INTEGER and NUMBER
	//  Minimum value of the Type.INTEGER and Type.NUMBER
	// +kcc:proto:field=google.cloud.aiplatform.v1.Schema.minimum
	Minimum *float64 `json:"minimum,omitempty"`

	// Optional. Maximum value of the Type.INTEGER and Type.NUMBER
	// +kcc:proto:field=google.cloud.aiplatform.v1.Schema.maximum
	Maximum *float64 `json:"maximum,omitempty"`

	// Optional. SCHEMA FIELDS FOR TYPE STRING
	//  Minimum length of the Type.STRING
	// +kcc:proto:field=google.cloud.aiplatform.v1.Schema.min_length
	MinLength *int64 `json:"minLength,omitempty"`

	// Optional. Maximum length of the Type.STRING
	// +kcc:proto:field=google.cloud.aiplatform.v1.Schema.max_length
	MaxLength *int64 `json:"maxLength,omitempty"`

	// Optional. Pattern of the Type.STRING to restrict a string to a regular
	//  expression.
	// +kcc:proto:field=google.cloud.aiplatform.v1.Schema.pattern
	Pattern *string `json:"pattern,omitempty"`

	// Optional. Example of the object. Will only populated when the object is the
	//  root.
	// +kcc:proto:field=google.cloud.aiplatform.v1.Schema.example
	Example *Value `json:"example,omitempty"`

	// Optional. The value should be validated against any (one or more) of the
	//  subschemas in the list.
	// +kcc:proto:field=google.cloud.aiplatform.v1.Schema.any_of
	AnyOf []Schema `json:"anyOf,omitempty"`
}

// +kcc:proto=google.cloud.aiplatform.v1.Tool
type Tool struct {
	// Optional. Function tool type.
	//  One or more function declarations to be passed to the model along with the
	//  current user query. Model may decide to call a subset of these functions
	//  by populating [FunctionCall][content.part.function_call] in the response.
	//  User should provide a [FunctionResponse][content.part.function_response]
	//  for each function call in the next turn. Based on the function responses,
	//  Model will generate the final response back to the user.
	//  Maximum 128 function declarations can be provided.
	// +kcc:proto:field=google.cloud.aiplatform.v1.Tool.function_declarations
	FunctionDeclarations []FunctionDeclaration `json:"functionDeclarations,omitempty"`

	// Optional. Retrieval tool type.
	//  System will always execute the provided retrieval tool(s) to get external
	//  knowledge to answer the prompt. Retrieval results are presented to the
	//  model for generation.
	// +kcc:proto:field=google.cloud.aiplatform.v1.Tool.retrieval
	Retrieval *Retrieval `json:"retrieval,omitempty"`

	// Optional. GoogleSearchRetrieval tool type.
	//  Specialized retrieval tool that is powered by Google search.
	// +kcc:proto:field=google.cloud.aiplatform.v1.Tool.google_search_retrieval
	GoogleSearchRetrieval *GoogleSearchRetrieval `json:"googleSearchRetrieval,omitempty"`
}

// +kcc:proto=google.cloud.aiplatform.v1.ToolConfig
type ToolConfig struct {
	// Optional. Function calling config.
	// +kcc:proto:field=google.cloud.aiplatform.v1.ToolConfig.function_calling_config
	FunctionCallingConfig *FunctionCallingConfig `json:"functionCallingConfig,omitempty"`

	// Optional. Retrieval config.
	// +kcc:proto:field=google.cloud.aiplatform.v1.ToolConfig.retrieval_config
	RetrievalConfig *RetrievalConfig `json:"retrievalConfig,omitempty"`
}

// +kcc:proto=google.cloud.aiplatform.v1.VertexAISearch
type VertexAISearch struct {
	// Required. Fully-qualified Vertex AI Search data store resource ID.
	//  Format:
	//  `projects/{project}/locations/{location}/collections/{collection}/dataStores/{dataStore}`
	// +kcc:proto:field=google.cloud.aiplatform.v1.VertexAISearch.datastore
	Datastore *string `json:"datastore,omitempty"`
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
	RagCorpus *string `json:"ragCorpus,omitempty"`

	// Optional. rag_file_id. The files should be in the same rag_corpus set in
	//  rag_corpus field.
	// +kcc:proto:field=google.cloud.aiplatform.v1.VertexRagStore.RagResource.rag_file_ids
	RagFileIds []string `json:"ragFileIds,omitempty"`
}

// +kcc:proto=google.cloud.aiplatform.v1.VideoMetadata
type VideoMetadata struct {
	// Optional. The start offset of the video.
	// +kcc:proto:field=google.cloud.aiplatform.v1.VideoMetadata.start_offset
	StartOffset *string `json:"startOffset,omitempty"`

	// Optional. The end offset of the video.
	// +kcc:proto:field=google.cloud.aiplatform.v1.VideoMetadata.end_offset
	EndOffset *string `json:"endOffset,omitempty"`
}

// +kcc:proto=google.protobuf.ListValue
type ListValue struct {
	// Repeated field of dynamically typed values.
	// +kcc:proto:field=google.protobuf.ListValue.values
	Values []Value `json:"values,omitempty"`
}

// +kcc:proto=google.protobuf.Value
type Value struct {
	// Represents a null value.
	// +kcc:proto:field=google.protobuf.Value.null_value
	NullValue *string `json:"nullValue,omitempty"`

	// Represents a double value.
	// +kcc:proto:field=google.protobuf.Value.number_value
	NumberValue *float64 `json:"numberValue,omitempty"`

	// Represents a string value.
	// +kcc:proto:field=google.protobuf.Value.string_value
	StringValue *string `json:"stringValue,omitempty"`

	// Represents a boolean value.
	// +kcc:proto:field=google.protobuf.Value.bool_value
	BoolValue *bool `json:"boolValue,omitempty"`

	// Represents a structured value.
	// +kcc:proto:field=google.protobuf.Value.struct_value
	StructValue map[string]string `json:"structValue,omitempty"`

	// Represents a repeated `Value`.
	// +kcc:proto:field=google.protobuf.Value.list_value
	ListValue *ListValue `json:"listValue,omitempty"`
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

// +kcc:proto=google.cloud.aiplatform.v1.CachedContent
type CachedContentObservedState struct {
	// Output only. Creatation time of the cache entry.
	// +kcc:proto:field=google.cloud.aiplatform.v1.CachedContent.create_time
	CreateTime *string `json:"createTime,omitempty"`

	// Output only. When the cache entry was last updated in UTC time.
	// +kcc:proto:field=google.cloud.aiplatform.v1.CachedContent.update_time
	UpdateTime *string `json:"updateTime,omitempty"`

	// Output only. Metadata on the usage of the cached content.
	// +kcc:proto:field=google.cloud.aiplatform.v1.CachedContent.usage_metadata
	UsageMetadata *CachedContent_UsageMetadata `json:"usageMetadata,omitempty"`
}
