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


// +kcc:proto=google.cloud.aiplatform.v1.Attribution
type Attribution struct {
}

// +kcc:proto=google.cloud.aiplatform.v1.BlurBaselineConfig
type BlurBaselineConfig struct {
	// The standard deviation of the blur kernel for the blurred baseline. The
	//  same blurring parameter is used for both the height and the width
	//  dimension. If not set, the method defaults to the zero (i.e. black for
	//  images) baseline.
	// +kcc:proto:field=google.cloud.aiplatform.v1.BlurBaselineConfig.max_blur_sigma
	MaxBlurSigma *float32 `json:"maxBlurSigma,omitempty"`
}

// +kcc:proto=google.cloud.aiplatform.v1.Examples
type Examples struct {
	// The Cloud Storage input instances.
	// +kcc:proto:field=google.cloud.aiplatform.v1.Examples.example_gcs_source
	ExampleGcsSource *Examples_ExampleGcsSource `json:"exampleGcsSource,omitempty"`

	// The full configuration for the generated index, the semantics are the
	//  same as [metadata][google.cloud.aiplatform.v1.Index.metadata] and should
	//  match
	//  [NearestNeighborSearchConfig](https://cloud.google.com/vertex-ai/docs/explainable-ai/configuring-explanations-example-based#nearest-neighbor-search-config).
	// +kcc:proto:field=google.cloud.aiplatform.v1.Examples.nearest_neighbor_search_config
	NearestNeighborSearchConfig *Value `json:"nearestNeighborSearchConfig,omitempty"`

	// Simplified preset configuration, which automatically sets configuration
	//  values based on the desired query speed-precision trade-off and modality.
	// +kcc:proto:field=google.cloud.aiplatform.v1.Examples.presets
	Presets *Presets `json:"presets,omitempty"`

	// The number of neighbors to return when querying for examples.
	// +kcc:proto:field=google.cloud.aiplatform.v1.Examples.neighbor_count
	NeighborCount *int32 `json:"neighborCount,omitempty"`
}

// +kcc:proto=google.cloud.aiplatform.v1.Examples.ExampleGcsSource
type Examples_ExampleGcsSource struct {
	// The format in which instances are given, if not specified, assume it's
	//  JSONL format. Currently only JSONL format is supported.
	// +kcc:proto:field=google.cloud.aiplatform.v1.Examples.ExampleGcsSource.data_format
	DataFormat *string `json:"dataFormat,omitempty"`

	// The Cloud Storage location for the input instances.
	// +kcc:proto:field=google.cloud.aiplatform.v1.Examples.ExampleGcsSource.gcs_source
	GcsSource *GcsSource `json:"gcsSource,omitempty"`
}

// +kcc:proto=google.cloud.aiplatform.v1.ExplanationMetadata
type ExplanationMetadata struct {

	// TODO: unsupported map type with key string and value message


	// TODO: unsupported map type with key string and value message


	// Points to a YAML file stored on Google Cloud Storage describing the format
	//  of the [feature
	//  attributions][google.cloud.aiplatform.v1.Attribution.feature_attributions].
	//  The schema is defined as an OpenAPI 3.0.2 [Schema
	//  Object](https://github.com/OAI/OpenAPI-Specification/blob/main/versions/3.0.2.md#schemaObject).
	//  AutoML tabular Models always have this field populated by Vertex AI.
	//  Note: The URI given on output may be different, including the URI scheme,
	//  than the one given on input. The output URI will point to a location where
	//  the user only has a read access.
	// +kcc:proto:field=google.cloud.aiplatform.v1.ExplanationMetadata.feature_attributions_schema_uri
	FeatureAttributionsSchemaURI *string `json:"featureAttributionsSchemaURI,omitempty"`

	// Name of the source to generate embeddings for example based explanations.
	// +kcc:proto:field=google.cloud.aiplatform.v1.ExplanationMetadata.latent_space_source
	LatentSpaceSource *string `json:"latentSpaceSource,omitempty"`
}

// +kcc:proto=google.cloud.aiplatform.v1.ExplanationMetadata.InputMetadata
type ExplanationMetadata_InputMetadata struct {
	// Baseline inputs for this feature.
	//
	//  If no baseline is specified, Vertex AI chooses the baseline for this
	//  feature. If multiple baselines are specified, Vertex AI returns the
	//  average attributions across them in
	//  [Attribution.feature_attributions][google.cloud.aiplatform.v1.Attribution.feature_attributions].
	//
	//  For Vertex AI-provided Tensorflow images (both 1.x and 2.x), the shape
	//  of each baseline must match the shape of the input tensor. If a scalar is
	//  provided, we broadcast to the same shape as the input tensor.
	//
	//  For custom images, the element of the baselines must be in the same
	//  format as the feature's input in the
	//  [instance][google.cloud.aiplatform.v1.ExplainRequest.instances][]. The
	//  schema of any single instance may be specified via Endpoint's
	//  DeployedModels' [Model's][google.cloud.aiplatform.v1.DeployedModel.model]
	//  [PredictSchemata's][google.cloud.aiplatform.v1.Model.predict_schemata]
	//  [instance_schema_uri][google.cloud.aiplatform.v1.PredictSchemata.instance_schema_uri].
	// +kcc:proto:field=google.cloud.aiplatform.v1.ExplanationMetadata.InputMetadata.input_baselines
	InputBaselines []Value `json:"inputBaselines,omitempty"`

	// Name of the input tensor for this feature. Required and is only
	//  applicable to Vertex AI-provided images for Tensorflow.
	// +kcc:proto:field=google.cloud.aiplatform.v1.ExplanationMetadata.InputMetadata.input_tensor_name
	InputTensorName *string `json:"inputTensorName,omitempty"`

	// Defines how the feature is encoded into the input tensor. Defaults to
	//  IDENTITY.
	// +kcc:proto:field=google.cloud.aiplatform.v1.ExplanationMetadata.InputMetadata.encoding
	Encoding *string `json:"encoding,omitempty"`

	// Modality of the feature. Valid values are: numeric, image. Defaults to
	//  numeric.
	// +kcc:proto:field=google.cloud.aiplatform.v1.ExplanationMetadata.InputMetadata.modality
	Modality *string `json:"modality,omitempty"`

	// The domain details of the input feature value. Like min/max, original
	//  mean or standard deviation if normalized.
	// +kcc:proto:field=google.cloud.aiplatform.v1.ExplanationMetadata.InputMetadata.feature_value_domain
	FeatureValueDomain *ExplanationMetadata_InputMetadata_FeatureValueDomain `json:"featureValueDomain,omitempty"`

	// Specifies the index of the values of the input tensor.
	//  Required when the input tensor is a sparse representation. Refer to
	//  Tensorflow documentation for more details:
	//  https://www.tensorflow.org/api_docs/python/tf/sparse/SparseTensor.
	// +kcc:proto:field=google.cloud.aiplatform.v1.ExplanationMetadata.InputMetadata.indices_tensor_name
	IndicesTensorName *string `json:"indicesTensorName,omitempty"`

	// Specifies the shape of the values of the input if the input is a sparse
	//  representation. Refer to Tensorflow documentation for more details:
	//  https://www.tensorflow.org/api_docs/python/tf/sparse/SparseTensor.
	// +kcc:proto:field=google.cloud.aiplatform.v1.ExplanationMetadata.InputMetadata.dense_shape_tensor_name
	DenseShapeTensorName *string `json:"denseShapeTensorName,omitempty"`

	// A list of feature names for each index in the input tensor.
	//  Required when the input
	//  [InputMetadata.encoding][google.cloud.aiplatform.v1.ExplanationMetadata.InputMetadata.encoding]
	//  is BAG_OF_FEATURES, BAG_OF_FEATURES_SPARSE, INDICATOR.
	// +kcc:proto:field=google.cloud.aiplatform.v1.ExplanationMetadata.InputMetadata.index_feature_mapping
	IndexFeatureMapping []string `json:"indexFeatureMapping,omitempty"`

	// Encoded tensor is a transformation of the input tensor. Must be provided
	//  if choosing
	//  [Integrated Gradients
	//  attribution][google.cloud.aiplatform.v1.ExplanationParameters.integrated_gradients_attribution]
	//  or [XRAI
	//  attribution][google.cloud.aiplatform.v1.ExplanationParameters.xrai_attribution]
	//  and the input tensor is not differentiable.
	//
	//  An encoded tensor is generated if the input tensor is encoded by a lookup
	//  table.
	// +kcc:proto:field=google.cloud.aiplatform.v1.ExplanationMetadata.InputMetadata.encoded_tensor_name
	EncodedTensorName *string `json:"encodedTensorName,omitempty"`

	// A list of baselines for the encoded tensor.
	//
	//  The shape of each baseline should match the shape of the encoded tensor.
	//  If a scalar is provided, Vertex AI broadcasts to the same shape as the
	//  encoded tensor.
	// +kcc:proto:field=google.cloud.aiplatform.v1.ExplanationMetadata.InputMetadata.encoded_baselines
	EncodedBaselines []Value `json:"encodedBaselines,omitempty"`

	// Visualization configurations for image explanation.
	// +kcc:proto:field=google.cloud.aiplatform.v1.ExplanationMetadata.InputMetadata.visualization
	Visualization *ExplanationMetadata_InputMetadata_Visualization `json:"visualization,omitempty"`

	// Name of the group that the input belongs to. Features with the same group
	//  name will be treated as one feature when computing attributions. Features
	//  grouped together can have different shapes in value. If provided, there
	//  will be one single attribution generated in
	//  [Attribution.feature_attributions][google.cloud.aiplatform.v1.Attribution.feature_attributions],
	//  keyed by the group name.
	// +kcc:proto:field=google.cloud.aiplatform.v1.ExplanationMetadata.InputMetadata.group_name
	GroupName *string `json:"groupName,omitempty"`
}

// +kcc:proto=google.cloud.aiplatform.v1.ExplanationMetadata.InputMetadata.FeatureValueDomain
type ExplanationMetadata_InputMetadata_FeatureValueDomain struct {
	// The minimum permissible value for this feature.
	// +kcc:proto:field=google.cloud.aiplatform.v1.ExplanationMetadata.InputMetadata.FeatureValueDomain.min_value
	MinValue *float32 `json:"minValue,omitempty"`

	// The maximum permissible value for this feature.
	// +kcc:proto:field=google.cloud.aiplatform.v1.ExplanationMetadata.InputMetadata.FeatureValueDomain.max_value
	MaxValue *float32 `json:"maxValue,omitempty"`

	// If this input feature has been normalized to a mean value of 0,
	//  the original_mean specifies the mean value of the domain prior to
	//  normalization.
	// +kcc:proto:field=google.cloud.aiplatform.v1.ExplanationMetadata.InputMetadata.FeatureValueDomain.original_mean
	OriginalMean *float32 `json:"originalMean,omitempty"`

	// If this input feature has been normalized to a standard deviation of
	//  1.0, the original_stddev specifies the standard deviation of the domain
	//  prior to normalization.
	// +kcc:proto:field=google.cloud.aiplatform.v1.ExplanationMetadata.InputMetadata.FeatureValueDomain.original_stddev
	OriginalStddev *float32 `json:"originalStddev,omitempty"`
}

// +kcc:proto=google.cloud.aiplatform.v1.ExplanationMetadata.InputMetadata.Visualization
type ExplanationMetadata_InputMetadata_Visualization struct {
	// Type of the image visualization. Only applicable to
	//  [Integrated Gradients
	//  attribution][google.cloud.aiplatform.v1.ExplanationParameters.integrated_gradients_attribution].
	//  OUTLINES shows regions of attribution, while PIXELS shows per-pixel
	//  attribution. Defaults to OUTLINES.
	// +kcc:proto:field=google.cloud.aiplatform.v1.ExplanationMetadata.InputMetadata.Visualization.type
	Type *string `json:"type,omitempty"`

	// Whether to only highlight pixels with positive contributions, negative
	//  or both. Defaults to POSITIVE.
	// +kcc:proto:field=google.cloud.aiplatform.v1.ExplanationMetadata.InputMetadata.Visualization.polarity
	Polarity *string `json:"polarity,omitempty"`

	// The color scheme used for the highlighted areas.
	//
	//  Defaults to PINK_GREEN for
	//  [Integrated Gradients
	//  attribution][google.cloud.aiplatform.v1.ExplanationParameters.integrated_gradients_attribution],
	//  which shows positive attributions in green and negative in pink.
	//
	//  Defaults to VIRIDIS for
	//  [XRAI
	//  attribution][google.cloud.aiplatform.v1.ExplanationParameters.xrai_attribution],
	//  which highlights the most influential regions in yellow and the least
	//  influential in blue.
	// +kcc:proto:field=google.cloud.aiplatform.v1.ExplanationMetadata.InputMetadata.Visualization.color_map
	ColorMap *string `json:"colorMap,omitempty"`

	// Excludes attributions above the specified percentile from the
	//  highlighted areas. Using the clip_percent_upperbound and
	//  clip_percent_lowerbound together can be useful for filtering out noise
	//  and making it easier to see areas of strong attribution. Defaults to
	//  99.9.
	// +kcc:proto:field=google.cloud.aiplatform.v1.ExplanationMetadata.InputMetadata.Visualization.clip_percent_upperbound
	ClipPercentUpperbound *float32 `json:"clipPercentUpperbound,omitempty"`

	// Excludes attributions below the specified percentile, from the
	//  highlighted areas. Defaults to 62.
	// +kcc:proto:field=google.cloud.aiplatform.v1.ExplanationMetadata.InputMetadata.Visualization.clip_percent_lowerbound
	ClipPercentLowerbound *float32 `json:"clipPercentLowerbound,omitempty"`

	// How the original image is displayed in the visualization.
	//  Adjusting the overlay can help increase visual clarity if the original
	//  image makes it difficult to view the visualization. Defaults to NONE.
	// +kcc:proto:field=google.cloud.aiplatform.v1.ExplanationMetadata.InputMetadata.Visualization.overlay_type
	OverlayType *string `json:"overlayType,omitempty"`
}

// +kcc:proto=google.cloud.aiplatform.v1.ExplanationMetadata.OutputMetadata
type ExplanationMetadata_OutputMetadata struct {
	// Static mapping between the index and display name.
	//
	//  Use this if the outputs are a deterministic n-dimensional array, e.g. a
	//  list of scores of all the classes in a pre-defined order for a
	//  multi-classification Model. It's not feasible if the outputs are
	//  non-deterministic, e.g. the Model produces top-k classes or sort the
	//  outputs by their values.
	//
	//  The shape of the value must be an n-dimensional array of strings. The
	//  number of dimensions must match that of the outputs to be explained.
	//  The
	//  [Attribution.output_display_name][google.cloud.aiplatform.v1.Attribution.output_display_name]
	//  is populated by locating in the mapping with
	//  [Attribution.output_index][google.cloud.aiplatform.v1.Attribution.output_index].
	// +kcc:proto:field=google.cloud.aiplatform.v1.ExplanationMetadata.OutputMetadata.index_display_name_mapping
	IndexDisplayNameMapping *Value `json:"indexDisplayNameMapping,omitempty"`

	// Specify a field name in the prediction to look for the display name.
	//
	//  Use this if the prediction contains the display names for the outputs.
	//
	//  The display names in the prediction must have the same shape of the
	//  outputs, so that it can be located by
	//  [Attribution.output_index][google.cloud.aiplatform.v1.Attribution.output_index]
	//  for a specific output.
	// +kcc:proto:field=google.cloud.aiplatform.v1.ExplanationMetadata.OutputMetadata.display_name_mapping_key
	DisplayNameMappingKey *string `json:"displayNameMappingKey,omitempty"`

	// Name of the output tensor. Required and is only applicable to Vertex
	//  AI provided images for Tensorflow.
	// +kcc:proto:field=google.cloud.aiplatform.v1.ExplanationMetadata.OutputMetadata.output_tensor_name
	OutputTensorName *string `json:"outputTensorName,omitempty"`
}

// +kcc:proto=google.cloud.aiplatform.v1.ExplanationParameters
type ExplanationParameters struct {
	// An attribution method that approximates Shapley values for features that
	//  contribute to the label being predicted. A sampling strategy is used to
	//  approximate the value rather than considering all subsets of features.
	//  Refer to this paper for model details: https://arxiv.org/abs/1306.4265.
	// +kcc:proto:field=google.cloud.aiplatform.v1.ExplanationParameters.sampled_shapley_attribution
	SampledShapleyAttribution *SampledShapleyAttribution `json:"sampledShapleyAttribution,omitempty"`

	// An attribution method that computes Aumann-Shapley values taking
	//  advantage of the model's fully differentiable structure. Refer to this
	//  paper for more details: https://arxiv.org/abs/1703.01365
	// +kcc:proto:field=google.cloud.aiplatform.v1.ExplanationParameters.integrated_gradients_attribution
	IntegratedGradientsAttribution *IntegratedGradientsAttribution `json:"integratedGradientsAttribution,omitempty"`

	// An attribution method that redistributes Integrated Gradients
	//  attribution to segmented regions, taking advantage of the model's fully
	//  differentiable structure. Refer to this paper for
	//  more details: https://arxiv.org/abs/1906.02825
	//
	//  XRAI currently performs better on natural images, like a picture of a
	//  house or an animal. If the images are taken in artificial environments,
	//  like a lab or manufacturing line, or from diagnostic equipment, like
	//  x-rays or quality-control cameras, use Integrated Gradients instead.
	// +kcc:proto:field=google.cloud.aiplatform.v1.ExplanationParameters.xrai_attribution
	XraiAttribution *XraiAttribution `json:"xraiAttribution,omitempty"`

	// Example-based explanations that returns the nearest neighbors from the
	//  provided dataset.
	// +kcc:proto:field=google.cloud.aiplatform.v1.ExplanationParameters.examples
	Examples *Examples `json:"examples,omitempty"`

	// If populated, returns attributions for top K indices of outputs
	//  (defaults to 1). Only applies to Models that predicts more than one outputs
	//  (e,g, multi-class Models). When set to -1, returns explanations for all
	//  outputs.
	// +kcc:proto:field=google.cloud.aiplatform.v1.ExplanationParameters.top_k
	TopK *int32 `json:"topK,omitempty"`

	// If populated, only returns attributions that have
	//  [output_index][google.cloud.aiplatform.v1.Attribution.output_index]
	//  contained in output_indices. It must be an ndarray of integers, with the
	//  same shape of the output it's explaining.
	//
	//  If not populated, returns attributions for
	//  [top_k][google.cloud.aiplatform.v1.ExplanationParameters.top_k] indices of
	//  outputs. If neither top_k nor output_indices is populated, returns the
	//  argmax index of the outputs.
	//
	//  Only applicable to Models that predict multiple outputs (e,g, multi-class
	//  Models that predict multiple classes).
	// +kcc:proto:field=google.cloud.aiplatform.v1.ExplanationParameters.output_indices
	OutputIndices *ListValue `json:"outputIndices,omitempty"`
}

// +kcc:proto=google.cloud.aiplatform.v1.ExplanationSpec
type ExplanationSpec struct {
	// Required. Parameters that configure explaining of the Model's predictions.
	// +kcc:proto:field=google.cloud.aiplatform.v1.ExplanationSpec.parameters
	Parameters *ExplanationParameters `json:"parameters,omitempty"`

	// Optional. Metadata describing the Model's input and output for explanation.
	// +kcc:proto:field=google.cloud.aiplatform.v1.ExplanationSpec.metadata
	Metadata *ExplanationMetadata `json:"metadata,omitempty"`
}

// +kcc:proto=google.cloud.aiplatform.v1.FeatureNoiseSigma
type FeatureNoiseSigma struct {
	// Noise sigma per feature. No noise is added to features that are not set.
	// +kcc:proto:field=google.cloud.aiplatform.v1.FeatureNoiseSigma.noise_sigma
	NoiseSigma []FeatureNoiseSigma_NoiseSigmaForFeature `json:"noiseSigma,omitempty"`
}

// +kcc:proto=google.cloud.aiplatform.v1.FeatureNoiseSigma.NoiseSigmaForFeature
type FeatureNoiseSigma_NoiseSigmaForFeature struct {
	// The name of the input feature for which noise sigma is provided. The
	//  features are defined in
	//  [explanation metadata
	//  inputs][google.cloud.aiplatform.v1.ExplanationMetadata.inputs].
	// +kcc:proto:field=google.cloud.aiplatform.v1.FeatureNoiseSigma.NoiseSigmaForFeature.name
	Name *string `json:"name,omitempty"`

	// This represents the standard deviation of the Gaussian kernel that will
	//  be used to add noise to the feature prior to computing gradients. Similar
	//  to [noise_sigma][google.cloud.aiplatform.v1.SmoothGradConfig.noise_sigma]
	//  but represents the noise added to the current feature. Defaults to 0.1.
	// +kcc:proto:field=google.cloud.aiplatform.v1.FeatureNoiseSigma.NoiseSigmaForFeature.sigma
	Sigma *float32 `json:"sigma,omitempty"`
}

// +kcc:proto=google.cloud.aiplatform.v1.GcsSource
type GcsSource struct {
	// Required. Google Cloud Storage URI(-s) to the input file(s). May contain
	//  wildcards. For more information on wildcards, see
	//  https://cloud.google.com/storage/docs/gsutil/addlhelp/WildcardNames.
	// +kcc:proto:field=google.cloud.aiplatform.v1.GcsSource.uris
	Uris []string `json:"uris,omitempty"`
}

// +kcc:proto=google.cloud.aiplatform.v1.IntegratedGradientsAttribution
type IntegratedGradientsAttribution struct {
	// Required. The number of steps for approximating the path integral.
	//  A good value to start is 50 and gradually increase until the
	//  sum to diff property is within the desired error range.
	//
	//  Valid range of its value is [1, 100], inclusively.
	// +kcc:proto:field=google.cloud.aiplatform.v1.IntegratedGradientsAttribution.step_count
	StepCount *int32 `json:"stepCount,omitempty"`

	// Config for SmoothGrad approximation of gradients.
	//
	//  When enabled, the gradients are approximated by averaging the gradients
	//  from noisy samples in the vicinity of the inputs. Adding
	//  noise can help improve the computed gradients. Refer to this paper for more
	//  details: https://arxiv.org/pdf/1706.03825.pdf
	// +kcc:proto:field=google.cloud.aiplatform.v1.IntegratedGradientsAttribution.smooth_grad_config
	SmoothGradConfig *SmoothGradConfig `json:"smoothGradConfig,omitempty"`

	// Config for IG with blur baseline.
	//
	//  When enabled, a linear path from the maximally blurred image to the input
	//  image is created. Using a blurred baseline instead of zero (black image) is
	//  motivated by the BlurIG approach explained here:
	//  https://arxiv.org/abs/2004.03383
	// +kcc:proto:field=google.cloud.aiplatform.v1.IntegratedGradientsAttribution.blur_baseline_config
	BlurBaselineConfig *BlurBaselineConfig `json:"blurBaselineConfig,omitempty"`
}

// +kcc:proto=google.cloud.aiplatform.v1.ModelEvaluation
type ModelEvaluation struct {

	// The display name of the ModelEvaluation.
	// +kcc:proto:field=google.cloud.aiplatform.v1.ModelEvaluation.display_name
	DisplayName *string `json:"displayName,omitempty"`

	// Points to a YAML file stored on Google Cloud Storage describing the
	//  [metrics][google.cloud.aiplatform.v1.ModelEvaluation.metrics] of this
	//  ModelEvaluation. The schema is defined as an OpenAPI 3.0.2 [Schema
	//  Object](https://github.com/OAI/OpenAPI-Specification/blob/main/versions/3.0.2.md#schemaObject).
	// +kcc:proto:field=google.cloud.aiplatform.v1.ModelEvaluation.metrics_schema_uri
	MetricsSchemaURI *string `json:"metricsSchemaURI,omitempty"`

	// Evaluation metrics of the Model. The schema of the metrics is stored in
	//  [metrics_schema_uri][google.cloud.aiplatform.v1.ModelEvaluation.metrics_schema_uri]
	// +kcc:proto:field=google.cloud.aiplatform.v1.ModelEvaluation.metrics
	Metrics *Value `json:"metrics,omitempty"`

	// All possible
	//  [dimensions][google.cloud.aiplatform.v1.ModelEvaluationSlice.Slice.dimension]
	//  of ModelEvaluationSlices. The dimensions can be used as the filter of the
	//  [ModelService.ListModelEvaluationSlices][google.cloud.aiplatform.v1.ModelService.ListModelEvaluationSlices]
	//  request, in the form of `slice.dimension = <dimension>`.
	// +kcc:proto:field=google.cloud.aiplatform.v1.ModelEvaluation.slice_dimensions
	SliceDimensions []string `json:"sliceDimensions,omitempty"`

	// Points to a YAML file stored on Google Cloud Storage describing
	//  [EvaluatedDataItemView.data_item_payload][] and
	//  [EvaluatedAnnotation.data_item_payload][google.cloud.aiplatform.v1.EvaluatedAnnotation.data_item_payload].
	//  The schema is defined as an OpenAPI 3.0.2 [Schema
	//  Object](https://github.com/OAI/OpenAPI-Specification/blob/main/versions/3.0.2.md#schemaObject).
	//
	//  This field is not populated if there are neither EvaluatedDataItemViews nor
	//  EvaluatedAnnotations under this ModelEvaluation.
	// +kcc:proto:field=google.cloud.aiplatform.v1.ModelEvaluation.data_item_schema_uri
	DataItemSchemaURI *string `json:"dataItemSchemaURI,omitempty"`

	// Points to a YAML file stored on Google Cloud Storage describing
	//  [EvaluatedDataItemView.predictions][],
	//  [EvaluatedDataItemView.ground_truths][],
	//  [EvaluatedAnnotation.predictions][google.cloud.aiplatform.v1.EvaluatedAnnotation.predictions],
	//  and
	//  [EvaluatedAnnotation.ground_truths][google.cloud.aiplatform.v1.EvaluatedAnnotation.ground_truths].
	//  The schema is defined as an OpenAPI 3.0.2 [Schema
	//  Object](https://github.com/OAI/OpenAPI-Specification/blob/main/versions/3.0.2.md#schemaObject).
	//
	//  This field is not populated if there are neither EvaluatedDataItemViews nor
	//  EvaluatedAnnotations under this ModelEvaluation.
	// +kcc:proto:field=google.cloud.aiplatform.v1.ModelEvaluation.annotation_schema_uri
	AnnotationSchemaURI *string `json:"annotationSchemaURI,omitempty"`

	// Aggregated explanation metrics for the Model's prediction output over the
	//  data this ModelEvaluation uses. This field is populated only if the Model
	//  is evaluated with explanations, and only for AutoML tabular Models.
	// +kcc:proto:field=google.cloud.aiplatform.v1.ModelEvaluation.model_explanation
	ModelExplanation *ModelExplanation `json:"modelExplanation,omitempty"`

	// Describes the values of
	//  [ExplanationSpec][google.cloud.aiplatform.v1.ExplanationSpec] that are used
	//  for explaining the predicted values on the evaluated data.
	// +kcc:proto:field=google.cloud.aiplatform.v1.ModelEvaluation.explanation_specs
	ExplanationSpecs []ModelEvaluation_ModelEvaluationExplanationSpec `json:"explanationSpecs,omitempty"`

	// The metadata of the ModelEvaluation.
	//  For the ModelEvaluation uploaded from Managed Pipeline, metadata contains a
	//  structured value with keys of "pipeline_job_id", "evaluation_dataset_type",
	//  "evaluation_dataset_path", "row_based_metrics_path".
	// +kcc:proto:field=google.cloud.aiplatform.v1.ModelEvaluation.metadata
	Metadata *Value `json:"metadata,omitempty"`
}

// +kcc:proto=google.cloud.aiplatform.v1.ModelEvaluation.ModelEvaluationExplanationSpec
type ModelEvaluation_ModelEvaluationExplanationSpec struct {
	// Explanation type.
	//
	//  For AutoML Image Classification models, possible values are:
	//
	//    * `image-integrated-gradients`
	//    * `image-xrai`
	// +kcc:proto:field=google.cloud.aiplatform.v1.ModelEvaluation.ModelEvaluationExplanationSpec.explanation_type
	ExplanationType *string `json:"explanationType,omitempty"`

	// Explanation spec details.
	// +kcc:proto:field=google.cloud.aiplatform.v1.ModelEvaluation.ModelEvaluationExplanationSpec.explanation_spec
	ExplanationSpec *ExplanationSpec `json:"explanationSpec,omitempty"`
}

// +kcc:proto=google.cloud.aiplatform.v1.ModelExplanation
type ModelExplanation struct {
}

// +kcc:proto=google.cloud.aiplatform.v1.Presets
type Presets struct {
	// Preset option controlling parameters for speed-precision trade-off when
	//  querying for examples. If omitted, defaults to `PRECISE`.
	// +kcc:proto:field=google.cloud.aiplatform.v1.Presets.query
	Query *string `json:"query,omitempty"`

	// The modality of the uploaded model, which automatically configures the
	//  distance measurement and feature normalization for the underlying example
	//  index and queries. If your model does not precisely fit one of these types,
	//  it is okay to choose the closest type.
	// +kcc:proto:field=google.cloud.aiplatform.v1.Presets.modality
	Modality *string `json:"modality,omitempty"`
}

// +kcc:proto=google.cloud.aiplatform.v1.SampledShapleyAttribution
type SampledShapleyAttribution struct {
	// Required. The number of feature permutations to consider when approximating
	//  the Shapley values.
	//
	//  Valid range of its value is [1, 50], inclusively.
	// +kcc:proto:field=google.cloud.aiplatform.v1.SampledShapleyAttribution.path_count
	PathCount *int32 `json:"pathCount,omitempty"`
}

// +kcc:proto=google.cloud.aiplatform.v1.SmoothGradConfig
type SmoothGradConfig struct {
	// This is a single float value and will be used to add noise to all the
	//  features. Use this field when all features are normalized to have the
	//  same distribution: scale to range [0, 1], [-1, 1] or z-scoring, where
	//  features are normalized to have 0-mean and 1-variance. Learn more about
	//  [normalization](https://developers.google.com/machine-learning/data-prep/transform/normalization).
	//
	//  For best results the recommended value is about 10% - 20% of the standard
	//  deviation of the input feature. Refer to section 3.2 of the SmoothGrad
	//  paper: https://arxiv.org/pdf/1706.03825.pdf. Defaults to 0.1.
	//
	//  If the distribution is different per feature, set
	//  [feature_noise_sigma][google.cloud.aiplatform.v1.SmoothGradConfig.feature_noise_sigma]
	//  instead for each feature.
	// +kcc:proto:field=google.cloud.aiplatform.v1.SmoothGradConfig.noise_sigma
	NoiseSigma *float32 `json:"noiseSigma,omitempty"`

	// This is similar to
	//  [noise_sigma][google.cloud.aiplatform.v1.SmoothGradConfig.noise_sigma],
	//  but provides additional flexibility. A separate noise sigma can be
	//  provided for each feature, which is useful if their distributions are
	//  different. No noise is added to features that are not set. If this field
	//  is unset,
	//  [noise_sigma][google.cloud.aiplatform.v1.SmoothGradConfig.noise_sigma]
	//  will be used for all features.
	// +kcc:proto:field=google.cloud.aiplatform.v1.SmoothGradConfig.feature_noise_sigma
	FeatureNoiseSigma *FeatureNoiseSigma `json:"featureNoiseSigma,omitempty"`

	// The number of gradient samples to use for
	//  approximation. The higher this number, the more accurate the gradient
	//  is, but the runtime complexity increases by this factor as well.
	//  Valid range of its value is [1, 50]. Defaults to 3.
	// +kcc:proto:field=google.cloud.aiplatform.v1.SmoothGradConfig.noisy_sample_count
	NoisySampleCount *int32 `json:"noisySampleCount,omitempty"`
}

// +kcc:proto=google.cloud.aiplatform.v1.XraiAttribution
type XraiAttribution struct {
	// Required. The number of steps for approximating the path integral.
	//  A good value to start is 50 and gradually increase until the
	//  sum to diff property is met within the desired error range.
	//
	//  Valid range of its value is [1, 100], inclusively.
	// +kcc:proto:field=google.cloud.aiplatform.v1.XraiAttribution.step_count
	StepCount *int32 `json:"stepCount,omitempty"`

	// Config for SmoothGrad approximation of gradients.
	//
	//  When enabled, the gradients are approximated by averaging the gradients
	//  from noisy samples in the vicinity of the inputs. Adding
	//  noise can help improve the computed gradients. Refer to this paper for more
	//  details: https://arxiv.org/pdf/1706.03825.pdf
	// +kcc:proto:field=google.cloud.aiplatform.v1.XraiAttribution.smooth_grad_config
	SmoothGradConfig *SmoothGradConfig `json:"smoothGradConfig,omitempty"`

	// Config for XRAI with blur baseline.
	//
	//  When enabled, a linear path from the maximally blurred image to the input
	//  image is created. Using a blurred baseline instead of zero (black image) is
	//  motivated by the BlurIG approach explained here:
	//  https://arxiv.org/abs/2004.03383
	// +kcc:proto:field=google.cloud.aiplatform.v1.XraiAttribution.blur_baseline_config
	BlurBaselineConfig *BlurBaselineConfig `json:"blurBaselineConfig,omitempty"`
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

// +kcc:proto=google.cloud.aiplatform.v1.Attribution
type AttributionObservedState struct {
	// Output only. Model predicted output if the input instance is constructed
	//  from the baselines of all the features defined in
	//  [ExplanationMetadata.inputs][google.cloud.aiplatform.v1.ExplanationMetadata.inputs].
	//  The field name of the output is determined by the key in
	//  [ExplanationMetadata.outputs][google.cloud.aiplatform.v1.ExplanationMetadata.outputs].
	//
	//  If the Model's predicted output has multiple dimensions (rank > 1), this is
	//  the value in the output located by
	//  [output_index][google.cloud.aiplatform.v1.Attribution.output_index].
	//
	//  If there are multiple baselines, their output values are averaged.
	// +kcc:proto:field=google.cloud.aiplatform.v1.Attribution.baseline_output_value
	BaselineOutputValue *float64 `json:"baselineOutputValue,omitempty"`

	// Output only. Model predicted output on the corresponding [explanation
	//  instance][ExplainRequest.instances]. The field name of the output is
	//  determined by the key in
	//  [ExplanationMetadata.outputs][google.cloud.aiplatform.v1.ExplanationMetadata.outputs].
	//
	//  If the Model predicted output has multiple dimensions, this is the value in
	//  the output located by
	//  [output_index][google.cloud.aiplatform.v1.Attribution.output_index].
	// +kcc:proto:field=google.cloud.aiplatform.v1.Attribution.instance_output_value
	InstanceOutputValue *float64 `json:"instanceOutputValue,omitempty"`

	// Output only. Attributions of each explained feature. Features are extracted
	//  from the [prediction
	//  instances][google.cloud.aiplatform.v1.ExplainRequest.instances] according
	//  to [explanation metadata for
	//  inputs][google.cloud.aiplatform.v1.ExplanationMetadata.inputs].
	//
	//  The value is a struct, whose keys are the name of the feature. The values
	//  are how much the feature in the
	//  [instance][google.cloud.aiplatform.v1.ExplainRequest.instances] contributed
	//  to the predicted result.
	//
	//  The format of the value is determined by the feature's input format:
	//
	//    * If the feature is a scalar value, the attribution value is a
	//      [floating number][google.protobuf.Value.number_value].
	//
	//    * If the feature is an array of scalar values, the attribution value is
	//      an [array][google.protobuf.Value.list_value].
	//
	//    * If the feature is a struct, the attribution value is a
	//      [struct][google.protobuf.Value.struct_value]. The keys in the
	//      attribution value struct are the same as the keys in the feature
	//      struct. The formats of the values in the attribution struct are
	//      determined by the formats of the values in the feature struct.
	//
	//  The
	//  [ExplanationMetadata.feature_attributions_schema_uri][google.cloud.aiplatform.v1.ExplanationMetadata.feature_attributions_schema_uri]
	//  field, pointed to by the
	//  [ExplanationSpec][google.cloud.aiplatform.v1.ExplanationSpec] field of the
	//  [Endpoint.deployed_models][google.cloud.aiplatform.v1.Endpoint.deployed_models]
	//  object, points to the schema file that describes the features and their
	//  attribution values (if it is populated).
	// +kcc:proto:field=google.cloud.aiplatform.v1.Attribution.feature_attributions
	FeatureAttributions *Value `json:"featureAttributions,omitempty"`

	// Output only. The index that locates the explained prediction output.
	//
	//  If the prediction output is a scalar value, output_index is not populated.
	//  If the prediction output has multiple dimensions, the length of the
	//  output_index list is the same as the number of dimensions of the output.
	//  The i-th element in output_index is the element index of the i-th dimension
	//  of the output vector. Indices start from 0.
	// +kcc:proto:field=google.cloud.aiplatform.v1.Attribution.output_index
	OutputIndex []int32 `json:"outputIndex,omitempty"`

	// Output only. The display name of the output identified by
	//  [output_index][google.cloud.aiplatform.v1.Attribution.output_index]. For
	//  example, the predicted class name by a multi-classification Model.
	//
	//  This field is only populated iff the Model predicts display names as a
	//  separate field along with the explained output. The predicted display name
	//  must has the same shape of the explained output, and can be located using
	//  output_index.
	// +kcc:proto:field=google.cloud.aiplatform.v1.Attribution.output_display_name
	OutputDisplayName *string `json:"outputDisplayName,omitempty"`

	// Output only. Error of
	//  [feature_attributions][google.cloud.aiplatform.v1.Attribution.feature_attributions]
	//  caused by approximation used in the explanation method. Lower value means
	//  more precise attributions.
	//
	//  * For Sampled Shapley
	//  [attribution][google.cloud.aiplatform.v1.ExplanationParameters.sampled_shapley_attribution],
	//  increasing
	//  [path_count][google.cloud.aiplatform.v1.SampledShapleyAttribution.path_count]
	//  might reduce the error.
	//  * For Integrated Gradients
	//  [attribution][google.cloud.aiplatform.v1.ExplanationParameters.integrated_gradients_attribution],
	//  increasing
	//  [step_count][google.cloud.aiplatform.v1.IntegratedGradientsAttribution.step_count]
	//  might reduce the error.
	//  * For [XRAI
	//  attribution][google.cloud.aiplatform.v1.ExplanationParameters.xrai_attribution],
	//  increasing
	//  [step_count][google.cloud.aiplatform.v1.XraiAttribution.step_count] might
	//  reduce the error.
	//
	//  See [this introduction](/vertex-ai/docs/explainable-ai/overview)
	//  for more information.
	// +kcc:proto:field=google.cloud.aiplatform.v1.Attribution.approximation_error
	ApproximationError *float64 `json:"approximationError,omitempty"`

	// Output only. Name of the explain output. Specified as the key in
	//  [ExplanationMetadata.outputs][google.cloud.aiplatform.v1.ExplanationMetadata.outputs].
	// +kcc:proto:field=google.cloud.aiplatform.v1.Attribution.output_name
	OutputName *string `json:"outputName,omitempty"`
}

// +kcc:proto=google.cloud.aiplatform.v1.ModelEvaluation
type ModelEvaluationObservedState struct {
	// Output only. The resource name of the ModelEvaluation.
	// +kcc:proto:field=google.cloud.aiplatform.v1.ModelEvaluation.name
	Name *string `json:"name,omitempty"`

	// Output only. Timestamp when this ModelEvaluation was created.
	// +kcc:proto:field=google.cloud.aiplatform.v1.ModelEvaluation.create_time
	CreateTime *string `json:"createTime,omitempty"`

	// Aggregated explanation metrics for the Model's prediction output over the
	//  data this ModelEvaluation uses. This field is populated only if the Model
	//  is evaluated with explanations, and only for AutoML tabular Models.
	// +kcc:proto:field=google.cloud.aiplatform.v1.ModelEvaluation.model_explanation
	ModelExplanation *ModelExplanationObservedState `json:"modelExplanation,omitempty"`
}

// +kcc:proto=google.cloud.aiplatform.v1.ModelExplanation
type ModelExplanationObservedState struct {
	// Output only. Aggregated attributions explaining the Model's prediction
	//  outputs over the set of instances. The attributions are grouped by outputs.
	//
	//  For Models that predict only one output, such as regression Models that
	//  predict only one score, there is only one attibution that explains the
	//  predicted output. For Models that predict multiple outputs, such as
	//  multiclass Models that predict multiple classes, each element explains one
	//  specific item.
	//  [Attribution.output_index][google.cloud.aiplatform.v1.Attribution.output_index]
	//  can be used to identify which output this attribution is explaining.
	//
	//  The
	//  [baselineOutputValue][google.cloud.aiplatform.v1.Attribution.baseline_output_value],
	//  [instanceOutputValue][google.cloud.aiplatform.v1.Attribution.instance_output_value]
	//  and
	//  [featureAttributions][google.cloud.aiplatform.v1.Attribution.feature_attributions]
	//  fields are averaged over the test data.
	//
	//  NOTE: Currently AutoML tabular classification Models produce only one
	//  attribution, which averages attributions over all the classes it predicts.
	//  [Attribution.approximation_error][google.cloud.aiplatform.v1.Attribution.approximation_error]
	//  is not populated.
	// +kcc:proto:field=google.cloud.aiplatform.v1.ModelExplanation.mean_attributions
	MeanAttributions []Attribution `json:"meanAttributions,omitempty"`
}
