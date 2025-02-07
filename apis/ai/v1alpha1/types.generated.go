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


// +kcc:proto=google.ai.generativelanguage.v1beta.Blob
type Blob struct {
	// The IANA standard MIME type of the source data.
	//  Examples:
	//    - image/png
	//    - image/jpeg
	//  If an unsupported MIME type is provided, an error will be returned. For a
	//  complete list of supported types, see [Supported file
	//  formats](https://ai.google.dev/gemini-api/docs/prompting_with_media#supported_file_formats).
	// +kcc:proto:field=google.ai.generativelanguage.v1beta.Blob.mime_type
	MimeType *string `json:"mimeType,omitempty"`

	// Raw bytes for media formats.
	// +kcc:proto:field=google.ai.generativelanguage.v1beta.Blob.data
	Data []byte `json:"data,omitempty"`
}

// +kcc:proto=google.ai.generativelanguage.v1beta.CachedContent
type CachedContent struct {
	// Timestamp in UTC of when this resource is considered expired.
	//  This is *always* provided on output, regardless of what was sent
	//  on input.
	// +kcc:proto:field=google.ai.generativelanguage.v1beta.CachedContent.expire_time
	ExpireTime *string `json:"expireTime,omitempty"`

	// Input only. New TTL for this resource, input only.
	// +kcc:proto:field=google.ai.generativelanguage.v1beta.CachedContent.ttl
	Ttl *string `json:"ttl,omitempty"`

	// Optional. Identifier. The resource name referring to the cached content.
	//  Format: `cachedContents/{id}`
	// +kcc:proto:field=google.ai.generativelanguage.v1beta.CachedContent.name
	Name *string `json:"name,omitempty"`

	// Optional. Immutable. The user-generated meaningful display name of the
	//  cached content. Maximum 128 Unicode characters.
	// +kcc:proto:field=google.ai.generativelanguage.v1beta.CachedContent.display_name
	DisplayName *string `json:"displayName,omitempty"`

	// Required. Immutable. The name of the `Model` to use for cached content
	//  Format: `models/{model}`
	// +kcc:proto:field=google.ai.generativelanguage.v1beta.CachedContent.model
	Model *string `json:"model,omitempty"`

	// Optional. Input only. Immutable. Developer set system instruction.
	//  Currently text only.
	// +kcc:proto:field=google.ai.generativelanguage.v1beta.CachedContent.system_instruction
	SystemInstruction *Content `json:"systemInstruction,omitempty"`

	// Optional. Input only. Immutable. The content to cache.
	// +kcc:proto:field=google.ai.generativelanguage.v1beta.CachedContent.contents
	Contents []Content `json:"contents,omitempty"`

	// Optional. Input only. Immutable. A list of `Tools` the model may use to
	//  generate the next response
	// +kcc:proto:field=google.ai.generativelanguage.v1beta.CachedContent.tools
	Tools []Tool `json:"tools,omitempty"`

	// Optional. Input only. Immutable. Tool config. This config is shared for all
	//  tools.
	// +kcc:proto:field=google.ai.generativelanguage.v1beta.CachedContent.tool_config
	ToolConfig *ToolConfig `json:"toolConfig,omitempty"`
}

// +kcc:proto=google.ai.generativelanguage.v1beta.CachedContent.UsageMetadata
type CachedContent_UsageMetadata struct {
	// Total number of tokens that the cached content consumes.
	// +kcc:proto:field=google.ai.generativelanguage.v1beta.CachedContent.UsageMetadata.total_token_count
	TotalTokenCount *int32 `json:"totalTokenCount,omitempty"`
}

// +kcc:proto=google.ai.generativelanguage.v1beta.CodeExecution
type CodeExecution struct {
}

// +kcc:proto=google.ai.generativelanguage.v1beta.CodeExecutionResult
type CodeExecutionResult struct {
	// Required. Outcome of the code execution.
	// +kcc:proto:field=google.ai.generativelanguage.v1beta.CodeExecutionResult.outcome
	Outcome *string `json:"outcome,omitempty"`

	// Optional. Contains stdout when code execution is successful, stderr or
	//  other description otherwise.
	// +kcc:proto:field=google.ai.generativelanguage.v1beta.CodeExecutionResult.output
	Output *string `json:"output,omitempty"`
}

// +kcc:proto=google.ai.generativelanguage.v1beta.Content
type Content struct {
	// Ordered `Parts` that constitute a single message. Parts may have different
	//  MIME types.
	// +kcc:proto:field=google.ai.generativelanguage.v1beta.Content.parts
	Parts []Part `json:"parts,omitempty"`

	// Optional. The producer of the content. Must be either 'user' or 'model'.
	//
	//  Useful to set for multi-turn conversations, otherwise can be left blank
	//  or unset.
	// +kcc:proto:field=google.ai.generativelanguage.v1beta.Content.role
	Role *string `json:"role,omitempty"`
}

// +kcc:proto=google.ai.generativelanguage.v1beta.DynamicRetrievalConfig
type DynamicRetrievalConfig struct {
	// The mode of the predictor to be used in dynamic retrieval.
	// +kcc:proto:field=google.ai.generativelanguage.v1beta.DynamicRetrievalConfig.mode
	Mode *string `json:"mode,omitempty"`

	// The threshold to be used in dynamic retrieval.
	//  If not set, a system default value is used.
	// +kcc:proto:field=google.ai.generativelanguage.v1beta.DynamicRetrievalConfig.dynamic_threshold
	DynamicThreshold *float32 `json:"dynamicThreshold,omitempty"`
}

// +kcc:proto=google.ai.generativelanguage.v1beta.ExecutableCode
type ExecutableCode struct {
	// Required. Programming language of the `code`.
	// +kcc:proto:field=google.ai.generativelanguage.v1beta.ExecutableCode.language
	Language *string `json:"language,omitempty"`

	// Required. The code to be executed.
	// +kcc:proto:field=google.ai.generativelanguage.v1beta.ExecutableCode.code
	Code *string `json:"code,omitempty"`
}

// +kcc:proto=google.ai.generativelanguage.v1beta.FileData
type FileData struct {
	// Optional. The IANA standard MIME type of the source data.
	// +kcc:proto:field=google.ai.generativelanguage.v1beta.FileData.mime_type
	MimeType *string `json:"mimeType,omitempty"`

	// Required. URI.
	// +kcc:proto:field=google.ai.generativelanguage.v1beta.FileData.file_uri
	FileURI *string `json:"fileURI,omitempty"`
}

// +kcc:proto=google.ai.generativelanguage.v1beta.FunctionCall
type FunctionCall struct {
	// Optional. The unique id of the function call. If populated, the client to
	//  execute the `function_call` and return the response with the matching `id`.
	// +kcc:proto:field=google.ai.generativelanguage.v1beta.FunctionCall.id
	ID *string `json:"id,omitempty"`

	// Required. The name of the function to call.
	//  Must be a-z, A-Z, 0-9, or contain underscores and dashes, with a maximum
	//  length of 63.
	// +kcc:proto:field=google.ai.generativelanguage.v1beta.FunctionCall.name
	Name *string `json:"name,omitempty"`

	// Optional. The function parameters and values in JSON object format.
	// +kcc:proto:field=google.ai.generativelanguage.v1beta.FunctionCall.args
	Args map[string]string `json:"args,omitempty"`
}

// +kcc:proto=google.ai.generativelanguage.v1beta.FunctionCallingConfig
type FunctionCallingConfig struct {
	// Optional. Specifies the mode in which function calling should execute. If
	//  unspecified, the default value will be set to AUTO.
	// +kcc:proto:field=google.ai.generativelanguage.v1beta.FunctionCallingConfig.mode
	Mode *string `json:"mode,omitempty"`

	// Optional. A set of function names that, when provided, limits the functions
	//  the model will call.
	//
	//  This should only be set when the Mode is ANY. Function names
	//  should match [FunctionDeclaration.name]. With mode set to ANY, model will
	//  predict a function call from the set of function names provided.
	// +kcc:proto:field=google.ai.generativelanguage.v1beta.FunctionCallingConfig.allowed_function_names
	AllowedFunctionNames []string `json:"allowedFunctionNames,omitempty"`
}

// +kcc:proto=google.ai.generativelanguage.v1beta.FunctionDeclaration
type FunctionDeclaration struct {
	// Required. The name of the function.
	//  Must be a-z, A-Z, 0-9, or contain underscores and dashes, with a maximum
	//  length of 63.
	// +kcc:proto:field=google.ai.generativelanguage.v1beta.FunctionDeclaration.name
	Name *string `json:"name,omitempty"`

	// Required. A brief description of the function.
	// +kcc:proto:field=google.ai.generativelanguage.v1beta.FunctionDeclaration.description
	Description *string `json:"description,omitempty"`

	// Optional. Describes the parameters to this function. Reflects the Open
	//  API 3.03 Parameter Object string Key: the name of the parameter. Parameter
	//  names are case sensitive. Schema Value: the Schema defining the type used
	//  for the parameter.
	// +kcc:proto:field=google.ai.generativelanguage.v1beta.FunctionDeclaration.parameters
	Parameters *Schema `json:"parameters,omitempty"`

	// Optional. Describes the output from this function in JSON Schema format.
	//  Reflects the Open API 3.03 Response Object. The Schema defines the type
	//  used for the response value of the function.
	// +kcc:proto:field=google.ai.generativelanguage.v1beta.FunctionDeclaration.response
	Response *Schema `json:"response,omitempty"`
}

// +kcc:proto=google.ai.generativelanguage.v1beta.FunctionResponse
type FunctionResponse struct {
	// Optional. The id of the function call this response is for. Populated by
	//  the client to match the corresponding function call `id`.
	// +kcc:proto:field=google.ai.generativelanguage.v1beta.FunctionResponse.id
	ID *string `json:"id,omitempty"`

	// Required. The name of the function to call.
	//  Must be a-z, A-Z, 0-9, or contain underscores and dashes, with a maximum
	//  length of 63.
	// +kcc:proto:field=google.ai.generativelanguage.v1beta.FunctionResponse.name
	Name *string `json:"name,omitempty"`

	// Required. The function response in JSON object format.
	// +kcc:proto:field=google.ai.generativelanguage.v1beta.FunctionResponse.response
	Response map[string]string `json:"response,omitempty"`
}

// +kcc:proto=google.ai.generativelanguage.v1beta.GoogleSearchRetrieval
type GoogleSearchRetrieval struct {
	// Specifies the dynamic retrieval configuration for the given source.
	// +kcc:proto:field=google.ai.generativelanguage.v1beta.GoogleSearchRetrieval.dynamic_retrieval_config
	DynamicRetrievalConfig *DynamicRetrievalConfig `json:"dynamicRetrievalConfig,omitempty"`
}

// +kcc:proto=google.ai.generativelanguage.v1beta.Part
type Part struct {
	// Inline text.
	// +kcc:proto:field=google.ai.generativelanguage.v1beta.Part.text
	Text *string `json:"text,omitempty"`

	// Inline media bytes.
	// +kcc:proto:field=google.ai.generativelanguage.v1beta.Part.inline_data
	InlineData *Blob `json:"inlineData,omitempty"`

	// A predicted `FunctionCall` returned from the model that contains
	//  a string representing the `FunctionDeclaration.name` with the
	//  arguments and their values.
	// +kcc:proto:field=google.ai.generativelanguage.v1beta.Part.function_call
	FunctionCall *FunctionCall `json:"functionCall,omitempty"`

	// The result output of a `FunctionCall` that contains a string
	//  representing the `FunctionDeclaration.name` and a structured JSON
	//  object containing any output from the function is used as context to
	//  the model.
	// +kcc:proto:field=google.ai.generativelanguage.v1beta.Part.function_response
	FunctionResponse *FunctionResponse `json:"functionResponse,omitempty"`

	// URI based data.
	// +kcc:proto:field=google.ai.generativelanguage.v1beta.Part.file_data
	FileData *FileData `json:"fileData,omitempty"`

	// Code generated by the model that is meant to be executed.
	// +kcc:proto:field=google.ai.generativelanguage.v1beta.Part.executable_code
	ExecutableCode *ExecutableCode `json:"executableCode,omitempty"`

	// Result of executing the `ExecutableCode`.
	// +kcc:proto:field=google.ai.generativelanguage.v1beta.Part.code_execution_result
	CodeExecutionResult *CodeExecutionResult `json:"codeExecutionResult,omitempty"`
}

// +kcc:proto=google.ai.generativelanguage.v1beta.Schema
type Schema struct {
	// Required. Data type.
	// +kcc:proto:field=google.ai.generativelanguage.v1beta.Schema.type
	Type *string `json:"type,omitempty"`

	// Optional. The format of the data. This is used only for primitive
	//  datatypes. Supported formats:
	//   for NUMBER type: float, double
	//   for INTEGER type: int32, int64
	//   for STRING type: enum
	// +kcc:proto:field=google.ai.generativelanguage.v1beta.Schema.format
	Format *string `json:"format,omitempty"`

	// Optional. A brief description of the parameter. This could contain examples
	//  of use. Parameter description may be formatted as Markdown.
	// +kcc:proto:field=google.ai.generativelanguage.v1beta.Schema.description
	Description *string `json:"description,omitempty"`

	// Optional. Indicates if the value may be null.
	// +kcc:proto:field=google.ai.generativelanguage.v1beta.Schema.nullable
	Nullable *bool `json:"nullable,omitempty"`

	// Optional. Possible values of the element of Type.STRING with enum format.
	//  For example we can define an Enum Direction as :
	//  {type:STRING, format:enum, enum:["EAST", NORTH", "SOUTH", "WEST"]}
	// +kcc:proto:field=google.ai.generativelanguage.v1beta.Schema.enum
	Enum []string `json:"enum,omitempty"`

	// Optional. Schema of the elements of Type.ARRAY.
	// +kcc:proto:field=google.ai.generativelanguage.v1beta.Schema.items
	Items *Schema `json:"items,omitempty"`

	// Optional. Maximum number of the elements for Type.ARRAY.
	// +kcc:proto:field=google.ai.generativelanguage.v1beta.Schema.max_items
	MaxItems *int64 `json:"maxItems,omitempty"`

	// Optional. Minimum number of the elements for Type.ARRAY.
	// +kcc:proto:field=google.ai.generativelanguage.v1beta.Schema.min_items
	MinItems *int64 `json:"minItems,omitempty"`

	// TODO: unsupported map type with key string and value message


	// Optional. Required properties of Type.OBJECT.
	// +kcc:proto:field=google.ai.generativelanguage.v1beta.Schema.required
	Required []string `json:"required,omitempty"`
}

// +kcc:proto=google.ai.generativelanguage.v1beta.Tool
type Tool struct {
	// Optional. A list of `FunctionDeclarations` available to the model that can
	//  be used for function calling.
	//
	//  The model or system does not execute the function. Instead the defined
	//  function may be returned as a
	//  [FunctionCall][google.ai.generativelanguage.v1beta.Part.function_call] with
	//  arguments to the client side for execution. The model may decide to call a
	//  subset of these functions by populating
	//  [FunctionCall][google.ai.generativelanguage.v1beta.Part.function_call] in
	//  the response. The next conversation turn may contain a
	//  [FunctionResponse][google.ai.generativelanguage.v1beta.Part.function_response]
	//  with the [Content.role][google.ai.generativelanguage.v1beta.Content.role]
	//  "function" generation context for the next model turn.
	// +kcc:proto:field=google.ai.generativelanguage.v1beta.Tool.function_declarations
	FunctionDeclarations []FunctionDeclaration `json:"functionDeclarations,omitempty"`

	// Optional. Retrieval tool that is powered by Google search.
	// +kcc:proto:field=google.ai.generativelanguage.v1beta.Tool.google_search_retrieval
	GoogleSearchRetrieval *GoogleSearchRetrieval `json:"googleSearchRetrieval,omitempty"`

	// Optional. Enables the model to execute code as part of generation.
	// +kcc:proto:field=google.ai.generativelanguage.v1beta.Tool.code_execution
	CodeExecution *CodeExecution `json:"codeExecution,omitempty"`

	// Optional. GoogleSearch tool type.
	//  Tool to support Google Search in Model. Powered by Google.
	// +kcc:proto:field=google.ai.generativelanguage.v1beta.Tool.google_search
	GoogleSearch *Tool_GoogleSearch `json:"googleSearch,omitempty"`
}

// +kcc:proto=google.ai.generativelanguage.v1beta.Tool.GoogleSearch
type Tool_GoogleSearch struct {
}

// +kcc:proto=google.ai.generativelanguage.v1beta.ToolConfig
type ToolConfig struct {
	// Optional. Function calling config.
	// +kcc:proto:field=google.ai.generativelanguage.v1beta.ToolConfig.function_calling_config
	FunctionCallingConfig *FunctionCallingConfig `json:"functionCallingConfig,omitempty"`
}

// +kcc:proto=google.ai.generativelanguage.v1beta.CachedContent
type CachedContentObservedState struct {
	// Output only. Creation time of the cache entry.
	// +kcc:proto:field=google.ai.generativelanguage.v1beta.CachedContent.create_time
	CreateTime *string `json:"createTime,omitempty"`

	// Output only. When the cache entry was last updated in UTC time.
	// +kcc:proto:field=google.ai.generativelanguage.v1beta.CachedContent.update_time
	UpdateTime *string `json:"updateTime,omitempty"`

	// Output only. Metadata on the usage of the cached content.
	// +kcc:proto:field=google.ai.generativelanguage.v1beta.CachedContent.usage_metadata
	UsageMetadata *CachedContent_UsageMetadata `json:"usageMetadata,omitempty"`
}
