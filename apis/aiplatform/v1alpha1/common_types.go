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
	apiextensionsv1 "k8s.io/apiextensions-apiserver/pkg/apis/apiextensions/v1"
)

// +kcc:proto=google.cloud.aiplatform.v1.BleuSpec
// +kcc:proto=google.cloud.aiplatform.v1beta1.BleuSpec
type BleuSpec struct {
	// Optional. Whether to use_effective_order to compute bleu score.
	// +kcc:proto:field=google.cloud.aiplatform.v1.BleuSpec.use_effective_order
	// +kcc:proto:field=google.cloud.aiplatform.v1beta1.BleuSpec.use_effective_order
	UseEffectiveOrder *bool `json:"useEffectiveOrder,omitempty"`
}

// +kcc:proto=google.cloud.aiplatform.v1.ComputationBasedMetricSpec
// +kcc:proto=google.cloud.aiplatform.v1beta1.ComputationBasedMetricSpec
type ComputationBasedMetricSpec struct {
	// Required. The type of the computation based metric.
	// +kcc:proto:field=google.cloud.aiplatform.v1.ComputationBasedMetricSpec.type
	// +kcc:proto:field=google.cloud.aiplatform.v1beta1.ComputationBasedMetricSpec.type
	Type *string `json:"type,omitempty"`

	// Optional. A map of parameters for the metric, e.g. {"rouge_type":
	//  "rougeL"}.
	// +kcc:proto:field=google.cloud.aiplatform.v1.ComputationBasedMetricSpec.parameters
	// +kcc:proto:field=google.cloud.aiplatform.v1beta1.ComputationBasedMetricSpec.parameters
	Parameters apiextensionsv1.JSON `json:"parameters,omitempty"`
}

// +kcc:proto=google.cloud.aiplatform.v1.CustomOutputFormatConfig
// +kcc:proto=google.cloud.aiplatform.v1beta1.CustomOutputFormatConfig
type CustomOutputFormatConfig struct {
	// Optional. Whether to return raw output.
	// +kcc:proto:field=google.cloud.aiplatform.v1.CustomOutputFormatConfig.return_raw_output
	// +kcc:proto:field=google.cloud.aiplatform.v1beta1.CustomOutputFormatConfig.return_raw_output
	ReturnRawOutput *bool `json:"returnRawOutput,omitempty"`
}

// +kcc:proto=google.cloud.aiplatform.v1.ExactMatchSpec
// +kcc:proto=google.cloud.aiplatform.v1beta1.ExactMatchSpec
type ExactMatchSpec struct {
}

// +kcc:proto=google.cloud.aiplatform.v1.GenerationConfig
// +kcc:proto=google.cloud.aiplatform.v1beta1.GenerationConfig
type GenerationConfig struct {
	// Optional. Controls the randomness of predictions.
	// +kcc:proto:field=google.cloud.aiplatform.v1.GenerationConfig.temperature
	// +kcc:proto:field=google.cloud.aiplatform.v1beta1.GenerationConfig.temperature
	Temperature *float32 `json:"temperature,omitempty"`

	// Optional. If specified, nucleus sampling will be used.
	// +kcc:proto:field=google.cloud.aiplatform.v1.GenerationConfig.top_p
	// +kcc:proto:field=google.cloud.aiplatform.v1beta1.GenerationConfig.top_p
	TopP *float32 `json:"topP,omitempty"`

	// Optional. If specified, top-k sampling will be used.
	// +kcc:proto:field=google.cloud.aiplatform.v1.GenerationConfig.top_k
	// +kcc:proto:field=google.cloud.aiplatform.v1beta1.GenerationConfig.top_k
	TopK *float32 `json:"topK,omitempty"`

	// Optional. Number of candidates to generate.
	// +kcc:proto:field=google.cloud.aiplatform.v1.GenerationConfig.candidate_count
	// +kcc:proto:field=google.cloud.aiplatform.v1beta1.GenerationConfig.candidate_count
	CandidateCount *int32 `json:"candidateCount,omitempty"`

	// Optional. The maximum number of output tokens to generate per message.
	// +kcc:proto:field=google.cloud.aiplatform.v1.GenerationConfig.max_output_tokens
	// +kcc:proto:field=google.cloud.aiplatform.v1beta1.GenerationConfig.max_output_tokens
	MaxOutputTokens *int32 `json:"maxOutputTokens,omitempty"`

	// Optional. Stop sequences.
	// +kcc:proto:field=google.cloud.aiplatform.v1.GenerationConfig.stop_sequences
	// +kcc:proto:field=google.cloud.aiplatform.v1beta1.GenerationConfig.stop_sequences
	StopSequences []string `json:"stopSequences,omitempty"`

	// Optional. If true, export the logprobs results in response.
	// +kcc:proto:field=google.cloud.aiplatform.v1.GenerationConfig.response_logprobs
	// +kcc:proto:field=google.cloud.aiplatform.v1beta1.GenerationConfig.response_logprobs
	ResponseLogprobs *bool `json:"responseLogprobs,omitempty"`

	// Optional. Logit probabilities.
	// +kcc:proto:field=google.cloud.aiplatform.v1.GenerationConfig.logprobs
	// +kcc:proto:field=google.cloud.aiplatform.v1beta1.GenerationConfig.logprobs
	Logprobs *int32 `json:"logprobs,omitempty"`

	// Optional. Positive penalties.
	// +kcc:proto:field=google.cloud.aiplatform.v1.GenerationConfig.presence_penalty
	// +kcc:proto:field=google.cloud.aiplatform.v1beta1.GenerationConfig.presence_penalty
	PresencePenalty *float32 `json:"presencePenalty,omitempty"`

	// Optional. Frequency penalties.
	// +kcc:proto:field=google.cloud.aiplatform.v1.GenerationConfig.frequency_penalty
	// +kcc:proto:field=google.cloud.aiplatform.v1beta1.GenerationConfig.frequency_penalty
	FrequencyPenalty *float32 `json:"frequencyPenalty,omitempty"`

	// Optional. Seed.
	// +kcc:proto:field=google.cloud.aiplatform.v1.GenerationConfig.seed
	// +kcc:proto:field=google.cloud.aiplatform.v1beta1.GenerationConfig.seed
	Seed *int32 `json:"seed,omitempty"`

	// Optional. Output response mimetype of the generated candidate text.
	//  Supported mimetype:
	//  - `text/plain`: (default) Text output.
	//  - `application/json`: JSON response in the candidates.
	//  The model needs to be prompted to output the appropriate response type,
	//  otherwise the behavior is undefined.
	//  This is a preview feature.
	// +kcc:proto:field=google.cloud.aiplatform.v1.GenerationConfig.response_mime_type
	// +kcc:proto:field=google.cloud.aiplatform.v1beta1.GenerationConfig.response_mime_type
	ResponseMimeType *string `json:"responseMimeType,omitempty"`

	// Optional. The `Schema` object allows the definition of input and output
	//  data types. These types can be objects, but also primitives and arrays.
	//  Represents a select subset of an [OpenAPI 3.0 schema
	//  object](https://spec.openapis.org/oas/v3.0.3#schema).
	//  If set, a compatible response_mime_type must also be set.
	//  Compatible mimetypes:
	//  `application/json`: Schema for JSON response.
	// +kcc:proto:field=google.cloud.aiplatform.v1.GenerationConfig.response_schema
	// +kcc:proto:field=google.cloud.aiplatform.v1beta1.GenerationConfig.response_schema
	ResponseSchema *Schema `json:"responseSchema,omitempty"`

	// Optional. Output schema of the generated response. This is an alternative
	//  to `response_schema` that accepts [JSON Schema](https://json-schema.org/).
	//
	//  If set, `response_schema` must be omitted, but `response_mime_type` is
	//  required.
	//
	//  While the full JSON Schema may be sent, not all features are supported.
	//  Specifically, only the following properties are supported:
	//
	//  - `$id`
	//  - `$defs`
	//  - `$ref`
	//  - `$anchor`
	//  - `type`
	//  - `format`
	//  - `title`
	//  - `description`
	//  - `enum` (for strings and numbers)
	//  - `items`
	//  - `prefixItems`
	//  - `minItems`
	//  - `maxItems`
	//  - `minimum`
	//  - `maximum`
	//  - `anyOf`
	//  - `oneOf` (interpreted the same as `anyOf`)
	//  - `properties`
	//  - `additionalProperties`
	//  - `required`
	//
	//  The non-standard `propertyOrdering` property may also be set.
	//
	//  Cyclic references are unrolled to a limited degree and, as such, may only
	//  be used within non-required properties. (Nullable properties are not
	//  sufficient.) If `$ref` is set on a sub-schema, no other properties, except
	//  for than those starting as a `$`, may be set.
	// +kcc:proto:field=google.cloud.aiplatform.v1.GenerationConfig.response_json_schema
	// +kcc:proto:field=google.cloud.aiplatform.v1beta1.GenerationConfig.response_json_schema
	ResponseJsonSchema *Value `json:"responseJsonSchema,omitempty"`

	// Optional. Routing configuration.
	// +kcc:proto:field=google.cloud.aiplatform.v1.GenerationConfig.routing_config
	// +kcc:proto:field=google.cloud.aiplatform.v1beta1.GenerationConfig.routing_config
	RoutingConfig *GenerationConfig_RoutingConfig `json:"routingConfig,omitempty"`

	// Optional. If enabled, audio timestamps will be included in the request to
	//  the model. This can be useful for synchronizing audio with other modalities
	//  in the response.
	// +kcc:proto:field=google.cloud.aiplatform.v1.GenerationConfig.audio_timestamp
	// +kcc:proto:field=google.cloud.aiplatform.v1beta1.GenerationConfig.audio_timestamp
	AudioTimestamp *bool `json:"audioTimestamp,omitempty"`

	// Optional. The modalities of the response. The model will generate a
	//  response that includes all the specified modalities. For example, if this
	//  is set to `[TEXT, IMAGE]`, the response will include both text and an
	//  image.
	// +kcc:proto:field=google.cloud.aiplatform.v1.GenerationConfig.response_modalities
	// +kcc:proto:field=google.cloud.aiplatform.v1beta1.GenerationConfig.response_modalities
	ResponseModalities []string `json:"responseModalities,omitempty"`

	// Optional. The token resolution at which input media content is sampled.
	//  This is used to control the trade-off between the quality of the response
	//  and the number of tokens used to represent the media. A higher resolution
	//  allows the model to perceive more detail, which can lead to a more nuanced
	//  response, but it will also use more tokens. This does not affect the
	//  image dimensions sent to the model.
	// +kcc:proto:field=google.cloud.aiplatform.v1.GenerationConfig.media_resolution
	// +kcc:proto:field=google.cloud.aiplatform.v1beta1.GenerationConfig.media_resolution
	MediaResolution *string `json:"mediaResolution,omitempty"`

	// Optional. The speech generation config.
	// +kcc:proto:field=google.cloud.aiplatform.v1.GenerationConfig.speech_config
	// +kcc:proto:field=google.cloud.aiplatform.v1beta1.GenerationConfig.speech_config
	SpeechConfig *SpeechConfig `json:"speechConfig,omitempty"`

	// Optional. Config for thinking features.
	//  An error will be returned if this field is set for models that don't
	//  support thinking.
	// +kcc:proto:field=google.cloud.aiplatform.v1.GenerationConfig.thinking_config
	// +kcc:proto:field=google.cloud.aiplatform.v1beta1.GenerationConfig.thinking_config
	ThinkingConfig *GenerationConfig_ThinkingConfig `json:"thinkingConfig,omitempty"`

	// Optional. Config for image generation features.
	// +kcc:proto:field=google.cloud.aiplatform.v1.GenerationConfig.image_config
	// +kcc:proto:field=google.cloud.aiplatform.v1beta1.GenerationConfig.image_config
	ImageConfig *ImageConfig `json:"imageConfig,omitempty"`
}

// +kcc:proto=google.cloud.aiplatform.v1.GenerationConfig.RoutingConfig
// +kcc:proto=google.cloud.aiplatform.v1beta1.GenerationConfig.RoutingConfig
type GenerationConfig_RoutingConfig struct {
	// Automated routing.
	// +kcc:proto:field=google.cloud.aiplatform.v1.GenerationConfig.RoutingConfig.auto_mode
	// +kcc:proto:field=google.cloud.aiplatform.v1beta1.GenerationConfig.RoutingConfig.auto_mode
	AutoMode *GenerationConfig_RoutingConfig_AutoRoutingMode `json:"autoMode,omitempty"`

	// Manual routing.
	// +kcc:proto:field=google.cloud.aiplatform.v1.GenerationConfig.RoutingConfig.manual_mode
	// +kcc:proto:field=google.cloud.aiplatform.v1beta1.GenerationConfig.RoutingConfig.manual_mode
	ManualMode *GenerationConfig_RoutingConfig_ManualRoutingMode `json:"manualMode,omitempty"`
}

// +kcc:proto=google.cloud.aiplatform.v1.GenerationConfig.RoutingConfig.AutoRoutingMode
// +kcc:proto=google.cloud.aiplatform.v1beta1.GenerationConfig.RoutingConfig.AutoRoutingMode
type GenerationConfig_RoutingConfig_AutoRoutingMode struct {
	// The model routing preference.
	// +kcc:proto:field=google.cloud.aiplatform.v1.GenerationConfig.RoutingConfig.AutoRoutingMode.model_routing_preference
	// +kcc:proto:field=google.cloud.aiplatform.v1beta1.GenerationConfig.RoutingConfig.AutoRoutingMode.model_routing_preference
	ModelRoutingPreference *string `json:"modelRoutingPreference,omitempty"`
}

// +kcc:proto=google.cloud.aiplatform.v1.GenerationConfig.RoutingConfig.ManualRoutingMode
// +kcc:proto=google.cloud.aiplatform.v1beta1.GenerationConfig.RoutingConfig.ManualRoutingMode
type GenerationConfig_RoutingConfig_ManualRoutingMode struct {
	// The model name to use. Only the public LLM models are accepted. e.g.
	//  'gemini-1.5-pro-001'.
	// +kcc:proto:field=google.cloud.aiplatform.v1.GenerationConfig.RoutingConfig.ManualRoutingMode.model_name
	// +kcc:proto:field=google.cloud.aiplatform.v1beta1.GenerationConfig.RoutingConfig.ManualRoutingMode.model_name
	ModelName *string `json:"modelName,omitempty"`
}

// +kcc:proto=google.cloud.aiplatform.v1.GenerationConfig.ThinkingConfig
// +kcc:proto=google.cloud.aiplatform.v1beta1.GenerationConfig.ThinkingConfig
type GenerationConfig_ThinkingConfig struct {
	// Indicates whether to include thoughts in the response.
	//  If true, thoughts are returned only when available.
	// +kcc:proto:field=google.cloud.aiplatform.v1.GenerationConfig.ThinkingConfig.include_thoughts
	// +kcc:proto:field=google.cloud.aiplatform.v1beta1.GenerationConfig.ThinkingConfig.include_thoughts
	IncludeThoughts *bool `json:"includeThoughts,omitempty"`

	// Optional. Indicates the thinking budget in tokens.
	//  This is only applied when enable_thinking is true.
	// +kcc:proto:field=google.cloud.aiplatform.v1.GenerationConfig.ThinkingConfig.thinking_budget
	// +kcc:proto:field=google.cloud.aiplatform.v1beta1.GenerationConfig.ThinkingConfig.thinking_budget
	ThinkingBudget *int32 `json:"thinkingBudget,omitempty"`

	// Optional. The number of thoughts tokens that the model should generate.
	// +kcc:proto:field=google.cloud.aiplatform.v1.GenerationConfig.ThinkingConfig.thinking_level
	// +kcc:proto:field=google.cloud.aiplatform.v1beta1.GenerationConfig.ThinkingConfig.thinking_level
	ThinkingLevel *string `json:"thinkingLevel,omitempty"`
}

// +kcc:proto=google.cloud.aiplatform.v1.ImageConfig
// +kcc:proto=google.cloud.aiplatform.v1beta1.ImageConfig
type ImageConfig struct {
	// Optional. The image output format for generated images.
	// +kcc:proto:field=google.cloud.aiplatform.v1.ImageConfig.image_output_options
	// +kcc:proto:field=google.cloud.aiplatform.v1beta1.ImageConfig.image_output_options
	ImageOutputOptions *ImageConfig_ImageOutputOptions `json:"imageOutputOptions,omitempty"`

	// Optional. The desired aspect ratio for the generated images. The following
	//  aspect ratios are supported:
	//
	//  "1:1"
	//  "2:3", "3:2"
	//  "3:4", "4:3"
	//  "4:5", "5:4"
	//  "9:16", "16:9"
	//  "21:9"
	// +kcc:proto:field=google.cloud.aiplatform.v1.ImageConfig.aspect_ratio
	// +kcc:proto:field=google.cloud.aiplatform.v1beta1.ImageConfig.aspect_ratio
	AspectRatio *string `json:"aspectRatio,omitempty"`

	// Optional. Controls whether the model can generate people.
	// +kcc:proto:field=google.cloud.aiplatform.v1.ImageConfig.person_generation
	// +kcc:proto:field=google.cloud.aiplatform.v1beta1.ImageConfig.person_generation
	PersonGeneration *string `json:"personGeneration,omitempty"`

	// Optional. Specifies the size of generated images. Supported values are
	//  `1K`, `2K`, `4K`. If not specified, the model will use default value `1K`.
	// +kcc:proto:field=google.cloud.aiplatform.v1.ImageConfig.image_size
	// +kcc:proto:field=google.cloud.aiplatform.v1beta1.ImageConfig.image_size
	ImageSize *string `json:"imageSize,omitempty"`
}

// +kcc:proto=google.cloud.aiplatform.v1.LLMBasedMetricSpec
// +kcc:proto=google.cloud.aiplatform.v1beta1.LLMBasedMetricSpec
type LlmBasedMetricSpec struct {
	// Use a pre-defined group of rubrics associated with the input.
	//  Refers to a key in the rubric_groups map of EvaluationInstance.
	// +kcc:proto:field=google.cloud.aiplatform.v1.LLMBasedMetricSpec.rubric_group_key
	// +kcc:proto:field=google.cloud.aiplatform.v1beta1.LLMBasedMetricSpec.rubric_group_key
	RubricGroupKey *string `json:"rubricGroupKey,omitempty"`

	// Dynamically generate rubrics using a predefined spec.
	// +kcc:proto:field=google.cloud.aiplatform.v1.LLMBasedMetricSpec.predefined_rubric_generation_spec
	// +kcc:proto:field=google.cloud.aiplatform.v1beta1.LLMBasedMetricSpec.predefined_rubric_generation_spec
	PredefinedRubricGenerationSpec *PredefinedMetricSpec `json:"predefinedRubricGenerationSpec,omitempty"`

	// Required. Template for the prompt sent to the judge model.
	// +kcc:proto:field=google.cloud.aiplatform.v1.LLMBasedMetricSpec.metric_prompt_template
	// +kcc:proto:field=google.cloud.aiplatform.v1beta1.LLMBasedMetricSpec.metric_prompt_template
	MetricPromptTemplate *string `json:"metricPromptTemplate,omitempty"`

	// Optional. System instructions for the judge model.
	// +kcc:proto:field=google.cloud.aiplatform.v1.LLMBasedMetricSpec.system_instruction
	// +kcc:proto:field=google.cloud.aiplatform.v1beta1.LLMBasedMetricSpec.system_instruction
	SystemInstruction *string `json:"systemInstruction,omitempty"`

	// Optional. Optional configuration for the judge LLM (Autorater).
	// +kcc:proto:field=google.cloud.aiplatform.v1.LLMBasedMetricSpec.judge_autorater_config
	// +kcc:proto:field=google.cloud.aiplatform.v1beta1.LLMBasedMetricSpec.judge_autorater_config
	JudgeAutoraterConfig *AutoraterConfig `json:"judgeAutoraterConfig,omitempty"`

	// Optional. Optional additional configuration for the metric.
	// +kcc:proto:field=google.cloud.aiplatform.v1.LLMBasedMetricSpec.additional_config
	// +kcc:proto:field=google.cloud.aiplatform.v1beta1.LLMBasedMetricSpec.additional_config
	AdditionalConfig apiextensionsv1.JSON `json:"additionalConfig,omitempty"`
}

// +kcc:proto=google.cloud.aiplatform.v1.Metric
// +kcc:proto=google.cloud.aiplatform.v1beta1.Metric
type Metric struct {
	// The spec for a pre-defined metric.
	// +kcc:proto:field=google.cloud.aiplatform.v1.Metric.predefined_metric_spec
	// +kcc:proto:field=google.cloud.aiplatform.v1beta1.Metric.predefined_metric_spec
	PredefinedMetricSpec *PredefinedMetricSpec `json:"predefinedMetricSpec,omitempty"`

	// Spec for a computation based metric.
	// +kcc:proto:field=google.cloud.aiplatform.v1.Metric.computation_based_metric_spec
	// +kcc:proto:field=google.cloud.aiplatform.v1beta1.Metric.computation_based_metric_spec
	ComputationBasedMetricSpec *ComputationBasedMetricSpec `json:"computationBasedMetricSpec,omitempty"`

	// Spec for an LLM based metric.
	// +kcc:proto:field=google.cloud.aiplatform.v1.Metric.llm_based_metric_spec
	// +kcc:proto:field=google.cloud.aiplatform.v1beta1.Metric.llm_based_metric_spec
	LlmBasedMetricSpec *LlmBasedMetricSpec `json:"llmBasedMetricSpec,omitempty"`

	// Spec for pointwise metric.
	// +kcc:proto:field=google.cloud.aiplatform.v1.Metric.pointwise_metric_spec
	// +kcc:proto:field=google.cloud.aiplatform.v1beta1.Metric.pointwise_metric_spec
	PointwiseMetricSpec *PointwiseMetricSpec `json:"pointwiseMetricSpec,omitempty"`

	// Spec for pairwise metric.
	// +kcc:proto:field=google.cloud.aiplatform.v1.Metric.pairwise_metric_spec
	// +kcc:proto:field=google.cloud.aiplatform.v1beta1.Metric.pairwise_metric_spec
	PairwiseMetricSpec *PairwiseMetricSpec `json:"pairwiseMetricSpec,omitempty"`

	// Spec for exact match metric.
	// +kcc:proto:field=google.cloud.aiplatform.v1.Metric.exact_match_spec
	// +kcc:proto:field=google.cloud.aiplatform.v1beta1.Metric.exact_match_spec
	ExactMatchSpec *ExactMatchSpec `json:"exactMatchSpec,omitempty"`

	// Spec for bleu metric.
	// +kcc:proto:field=google.cloud.aiplatform.v1.Metric.bleu_spec
	// +kcc:proto:field=google.cloud.aiplatform.v1beta1.Metric.bleu_spec
	BleuSpec *BleuSpec `json:"bleuSpec,omitempty"`

	// Spec for rouge metric.
	// +kcc:proto:field=google.cloud.aiplatform.v1.Metric.rouge_spec
	// +kcc:proto:field=google.cloud.aiplatform.v1beta1.Metric.rouge_spec
	RougeSpec *RougeSpec `json:"rougeSpec,omitempty"`

	// Optional. The aggregation metrics to use.
	// +kcc:proto:field=google.cloud.aiplatform.v1.Metric.aggregation_metrics
	// +kcc:proto:field=google.cloud.aiplatform.v1beta1.Metric.aggregation_metrics
	AggregationMetrics []string `json:"aggregationMetrics,omitempty"`
}

// +kcc:proto=google.cloud.aiplatform.v1.MultiSpeakerVoiceConfig
// +kcc:proto=google.cloud.aiplatform.v1beta1.MultiSpeakerVoiceConfig
type MultiSpeakerVoiceConfig struct {
	// Required. A list of configurations for the voices of the speakers. Exactly
	//  two speaker voice configurations must be provided.
	// +kcc:proto:field=google.cloud.aiplatform.v1.MultiSpeakerVoiceConfig.speaker_voice_configs
	// +kcc:proto:field=google.cloud.aiplatform.v1beta1.MultiSpeakerVoiceConfig.speaker_voice_configs
	SpeakerVoiceConfigs []SpeakerVoiceConfig `json:"speakerVoiceConfigs,omitempty"`
}

// +kcc:proto=google.cloud.aiplatform.v1.PairwiseMetricSpec
// +kcc:proto=google.cloud.aiplatform.v1beta1.PairwiseMetricSpec
type PairwiseMetricSpec struct {
	// Required. Metric prompt template for pairwise metric.
	// +kcc:proto:field=google.cloud.aiplatform.v1.PairwiseMetricSpec.metric_prompt_template
	// +kcc:proto:field=google.cloud.aiplatform.v1beta1.PairwiseMetricSpec.metric_prompt_template
	MetricPromptTemplate *string `json:"metricPromptTemplate,omitempty"`

	// Optional. The field name of the candidate response.
	// +kcc:proto:field=google.cloud.aiplatform.v1.PairwiseMetricSpec.candidate_response_field_name
	// +kcc:proto:field=google.cloud.aiplatform.v1beta1.PairwiseMetricSpec.candidate_response_field_name
	CandidateResponseFieldName *string `json:"candidateResponseFieldName,omitempty"`

	// Optional. The field name of the baseline response.
	// +kcc:proto:field=google.cloud.aiplatform.v1.PairwiseMetricSpec.baseline_response_field_name
	// +kcc:proto:field=google.cloud.aiplatform.v1beta1.PairwiseMetricSpec.baseline_response_field_name
	BaselineResponseFieldName *string `json:"baselineResponseFieldName,omitempty"`

	// Optional. System instructions for pairwise metric.
	// +kcc:proto:field=google.cloud.aiplatform.v1.PairwiseMetricSpec.system_instruction
	// +kcc:proto:field=google.cloud.aiplatform.v1beta1.PairwiseMetricSpec.system_instruction
	SystemInstruction *string `json:"systemInstruction,omitempty"`

	// Optional. CustomOutputFormatConfig allows customization of metric output.
	//  When this config is set, the default output is replaced with
	//  the raw output string.
	//  If a custom format is chosen, the `pairwise_choice` and `explanation`
	//  fields in the corresponding metric result will be empty.
	// +kcc:proto:field=google.cloud.aiplatform.v1.PairwiseMetricSpec.custom_output_format_config
	// +kcc:proto:field=google.cloud.aiplatform.v1beta1.PairwiseMetricSpec.custom_output_format_config
	CustomOutputFormatConfig *CustomOutputFormatConfig `json:"customOutputFormatConfig,omitempty"`
}

// +kcc:proto=google.cloud.aiplatform.v1.PointwiseMetricSpec
// +kcc:proto=google.cloud.aiplatform.v1beta1.PointwiseMetricSpec
type PointwiseMetricSpec struct {
	// Required. Metric prompt template for pointwise metric.
	// +kcc:proto:field=google.cloud.aiplatform.v1.PointwiseMetricSpec.metric_prompt_template
	// +kcc:proto:field=google.cloud.aiplatform.v1beta1.PointwiseMetricSpec.metric_prompt_template
	MetricPromptTemplate *string `json:"metricPromptTemplate,omitempty"`

	// Optional. System instructions for pointwise metric.
	// +kcc:proto:field=google.cloud.aiplatform.v1.PointwiseMetricSpec.system_instruction
	// +kcc:proto:field=google.cloud.aiplatform.v1beta1.PointwiseMetricSpec.system_instruction
	SystemInstruction *string `json:"systemInstruction,omitempty"`

	// Optional. CustomOutputFormatConfig allows customization of metric output.
	//  By default, metrics return a score and explanation.
	//  When this config is set, the default output is replaced with either:
	//   - The raw output string.
	//   - A parsed output based on a user-defined schema.
	//  If a custom format is chosen, the `score` and `explanation` fields in the
	//  corresponding metric result will be empty.
	// +kcc:proto:field=google.cloud.aiplatform.v1.PointwiseMetricSpec.custom_output_format_config
	// +kcc:proto:field=google.cloud.aiplatform.v1beta1.PointwiseMetricSpec.custom_output_format_config
	CustomOutputFormatConfig *CustomOutputFormatConfig `json:"customOutputFormatConfig,omitempty"`
}

// +kcc:proto=google.cloud.aiplatform.v1.PrebuiltVoiceConfig
// +kcc:proto=google.cloud.aiplatform.v1beta1.PrebuiltVoiceConfig
type PrebuiltVoiceConfig struct {
	// The name of the prebuilt voice to use.
	// +kcc:proto:field=google.cloud.aiplatform.v1.PrebuiltVoiceConfig.voice_name
	// +kcc:proto:field=google.cloud.aiplatform.v1beta1.PrebuiltVoiceConfig.voice_name
	VoiceName *string `json:"voiceName,omitempty"`
}

// +kcc:proto=google.cloud.aiplatform.v1.PredefinedMetricSpec
// +kcc:proto=google.cloud.aiplatform.v1beta1.PredefinedMetricSpec
type PredefinedMetricSpec struct {
	// Required. The name of a pre-defined metric, such as
	//  "instruction_following_v1" or "text_quality_v1".
	// +kcc:proto:field=google.cloud.aiplatform.v1.PredefinedMetricSpec.metric_spec_name
	// +kcc:proto:field=google.cloud.aiplatform.v1beta1.PredefinedMetricSpec.metric_spec_name
	MetricSpecName *string `json:"metricSpecName,omitempty"`

	// Optional. The parameters needed to run the pre-defined metric.
	// +kcc:proto:field=google.cloud.aiplatform.v1.PredefinedMetricSpec.metric_spec_parameters
	// +kcc:proto:field=google.cloud.aiplatform.v1beta1.PredefinedMetricSpec.metric_spec_parameters
	MetricSpecParameters apiextensionsv1.JSON `json:"metricSpecParameters,omitempty"`
}

// +kcc:proto=google.cloud.aiplatform.v1.ReplicatedVoiceConfig
// +kcc:proto=google.cloud.aiplatform.v1beta1.ReplicatedVoiceConfig
type ReplicatedVoiceConfig struct {
	// Optional. The mimetype of the voice sample. The only currently supported
	//  value is `audio/wav`. This represents 16-bit signed little-endian wav data,
	//  with a 24kHz sampling rate. `mime_type` will default to `audio/wav` if not
	//  set.
	// +kcc:proto:field=google.cloud.aiplatform.v1.ReplicatedVoiceConfig.mime_type
	// +kcc:proto:field=google.cloud.aiplatform.v1beta1.ReplicatedVoiceConfig.mime_type
	MimeType *string `json:"mimeType,omitempty"`

	// Optional. The sample of the custom voice.
	// +kcc:proto:field=google.cloud.aiplatform.v1.ReplicatedVoiceConfig.voice_sample_audio
	// +kcc:proto:field=google.cloud.aiplatform.v1beta1.ReplicatedVoiceConfig.voice_sample_audio
	VoiceSampleAudio []byte `json:"voiceSampleAudio,omitempty"`
}

// +kcc:proto=google.cloud.aiplatform.v1.RougeSpec
// +kcc:proto=google.cloud.aiplatform.v1beta1.RougeSpec
type RougeSpec struct {
	// Optional. Supported rouge types are rougen[1-9], rougeL, and rougeLsum.
	// +kcc:proto:field=google.cloud.aiplatform.v1.RougeSpec.rouge_type
	// +kcc:proto:field=google.cloud.aiplatform.v1beta1.RougeSpec.rouge_type
	RougeType *string `json:"rougeType,omitempty"`

	// Optional. Whether to use stemmer to compute rouge score.
	// +kcc:proto:field=google.cloud.aiplatform.v1.RougeSpec.use_stemmer
	// +kcc:proto:field=google.cloud.aiplatform.v1beta1.RougeSpec.use_stemmer
	UseStemmer *bool `json:"useStemmer,omitempty"`

	// Optional. Whether to split summaries while using rougeLsum.
	// +kcc:proto:field=google.cloud.aiplatform.v1.RougeSpec.split_summaries
	// +kcc:proto:field=google.cloud.aiplatform.v1beta1.RougeSpec.split_summaries
	SplitSummaries *bool `json:"splitSummaries,omitempty"`
}

// +kcc:proto=google.cloud.aiplatform.v1.SpeakerVoiceConfig
// +kcc:proto=google.cloud.aiplatform.v1beta1.SpeakerVoiceConfig
type SpeakerVoiceConfig struct {
	// Required. The name of the speaker. This should be the same as the speaker
	//  name used in the prompt.
	// +kcc:proto:field=google.cloud.aiplatform.v1.SpeakerVoiceConfig.speaker
	// +kcc:proto:field=google.cloud.aiplatform.v1beta1.SpeakerVoiceConfig.speaker
	Speaker *string `json:"speaker,omitempty"`

	// Required. The configuration for the voice of this speaker.
	// +kcc:proto:field=google.cloud.aiplatform.v1.SpeakerVoiceConfig.voice_config
	// +kcc:proto:field=google.cloud.aiplatform.v1beta1.SpeakerVoiceConfig.voice_config
	VoiceConfig *VoiceConfig `json:"voiceConfig,omitempty"`
}

// +kcc:proto=google.cloud.aiplatform.v1.SpeechConfig
// +kcc:proto=google.cloud.aiplatform.v1beta1.SpeechConfig
type SpeechConfig struct {
	// The configuration for the voice to use.
	// +kcc:proto:field=google.cloud.aiplatform.v1.SpeechConfig.voice_config
	// +kcc:proto:field=google.cloud.aiplatform.v1beta1.SpeechConfig.voice_config
	VoiceConfig *VoiceConfig `json:"voiceConfig,omitempty"`

	// Optional. The language code (ISO 639-1) for the speech synthesis.
	// +kcc:proto:field=google.cloud.aiplatform.v1.SpeechConfig.language_code
	// +kcc:proto:field=google.cloud.aiplatform.v1beta1.SpeechConfig.language_code
	LanguageCode *string `json:"languageCode,omitempty"`

	// The configuration for a multi-speaker text-to-speech request.
	//  This field is mutually exclusive with `voice_config`.
	// +kcc:proto:field=google.cloud.aiplatform.v1.SpeechConfig.multi_speaker_voice_config
	// +kcc:proto:field=google.cloud.aiplatform.v1beta1.SpeechConfig.multi_speaker_voice_config
	MultiSpeakerVoiceConfig *MultiSpeakerVoiceConfig `json:"multiSpeakerVoiceConfig,omitempty"`
}

// +kcc:proto=google.cloud.aiplatform.v1.VoiceConfig
// +kcc:proto=google.cloud.aiplatform.v1beta1.VoiceConfig
type VoiceConfig struct {
	// The configuration for a prebuilt voice.
	// +kcc:proto:field=google.cloud.aiplatform.v1.VoiceConfig.prebuilt_voice_config
	// +kcc:proto:field=google.cloud.aiplatform.v1beta1.VoiceConfig.prebuilt_voice_config
	PrebuiltVoiceConfig *PrebuiltVoiceConfig `json:"prebuiltVoiceConfig,omitempty"`

	// Optional. The configuration for a replicated voice. This enables users to
	//  replicate a voice from an audio sample.
	// +kcc:proto:field=google.cloud.aiplatform.v1.VoiceConfig.replicated_voice_config
	// +kcc:proto:field=google.cloud.aiplatform.v1beta1.VoiceConfig.replicated_voice_config
	ReplicatedVoiceConfig *ReplicatedVoiceConfig `json:"replicatedVoiceConfig,omitempty"`
}
