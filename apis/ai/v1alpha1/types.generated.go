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


// +kcc:proto=google.ai.generativelanguage.v1beta2.Model
type Model struct {
	// Required. The resource name of the `Model`.
	//
	//  Format: `models/{model}` with a `{model}` naming convention of:
	//
	//  * "{base_model_id}-{version}"
	//
	//  Examples:
	//
	//  * `models/chat-bison-001`
	// +kcc:proto:field=google.ai.generativelanguage.v1beta2.Model.name
	Name *string `json:"name,omitempty"`

	// Required. The name of the base model, pass this to the generation request.
	//
	//  Examples:
	//
	//  * `chat-bison`
	// +kcc:proto:field=google.ai.generativelanguage.v1beta2.Model.base_model_id
	BaseModelID *string `json:"baseModelID,omitempty"`

	// Required. The version number of the model.
	//
	//  This represents the major version
	// +kcc:proto:field=google.ai.generativelanguage.v1beta2.Model.version
	Version *string `json:"version,omitempty"`

	// The human-readable name of the model. E.g. "Chat Bison".
	//
	//  The name can be up to 128 characters long and can consist of any UTF-8
	//  characters.
	// +kcc:proto:field=google.ai.generativelanguage.v1beta2.Model.display_name
	DisplayName *string `json:"displayName,omitempty"`

	// A short description of the model.
	// +kcc:proto:field=google.ai.generativelanguage.v1beta2.Model.description
	Description *string `json:"description,omitempty"`

	// Maximum number of input tokens allowed for this model.
	// +kcc:proto:field=google.ai.generativelanguage.v1beta2.Model.input_token_limit
	InputTokenLimit *int32 `json:"inputTokenLimit,omitempty"`

	// Maximum number of output tokens available for this model.
	// +kcc:proto:field=google.ai.generativelanguage.v1beta2.Model.output_token_limit
	OutputTokenLimit *int32 `json:"outputTokenLimit,omitempty"`

	// The model's supported generation methods.
	//
	//  The method names are defined as Pascal case
	//  strings, such as `generateMessage` which correspond to API methods.
	// +kcc:proto:field=google.ai.generativelanguage.v1beta2.Model.supported_generation_methods
	SupportedGenerationMethods []string `json:"supportedGenerationMethods,omitempty"`

	// Controls the randomness of the output.
	//
	//  Values can range over `[0.0,1.0]`, inclusive. A value closer to `1.0` will
	//  produce responses that are more varied, while a value closer to `0.0` will
	//  typically result in less surprising responses from the model.
	//  This value specifies default to be used by the backend while making the
	//  call to the model.
	// +kcc:proto:field=google.ai.generativelanguage.v1beta2.Model.temperature
	Temperature *float32 `json:"temperature,omitempty"`

	// For Nucleus sampling.
	//
	//  Nucleus sampling considers the smallest set of tokens whose probability
	//  sum is at least `top_p`.
	//  This value specifies default to be used by the backend while making the
	//  call to the model.
	// +kcc:proto:field=google.ai.generativelanguage.v1beta2.Model.top_p
	TopP *float32 `json:"topP,omitempty"`

	// For Top-k sampling.
	//
	//  Top-k sampling considers the set of `top_k` most probable tokens.
	//  This value specifies default to be used by the backend while making the
	//  call to the model.
	// +kcc:proto:field=google.ai.generativelanguage.v1beta2.Model.top_k
	TopK *int32 `json:"topK,omitempty"`
}
