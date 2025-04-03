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

// +kcc:proto=google.cloud.aiplatform.v1.BlurBaselineConfig
type BlurBaselineConfig struct {
	// The standard deviation of the blur kernel for the blurred baseline. The
	//  same blurring parameter is used for both the height and the width
	//  dimension. If not set, the method defaults to the zero (i.e. black for
	//  images) baseline.
	// +kcc:proto:field=google.cloud.aiplatform.v1.BlurBaselineConfig.max_blur_sigma
	MaxBlurSigma *float32 `json:"maxBlurSigma,omitempty"`
}

// +kcc:proto=google.cloud.aiplatform.v1.DeployedModelRef
type DeployedModelRef struct {
	// Immutable. A resource name of an Endpoint.
	// +kcc:proto:field=google.cloud.aiplatform.v1.DeployedModelRef.endpoint
	Endpoint *string `json:"endpoint,omitempty"`

	// Immutable. An ID of a DeployedModel in the above Endpoint.
	// +kcc:proto:field=google.cloud.aiplatform.v1.DeployedModelRef.deployed_model_id
	DeployedModelID *string `json:"deployedModelID,omitempty"`
}

// +kcc:proto=google.cloud.aiplatform.v1.EncryptionSpec
type EncryptionSpec struct {
	// Required. The Cloud KMS resource identifier of the customer managed
	//  encryption key used to protect a resource. Has the form:
	//  `projects/my-project/locations/my-region/keyRings/my-kr/cryptoKeys/my-key`.
	//  The key needs to be in the same region as where the compute resource is
	//  created.
	// +kcc:proto:field=google.cloud.aiplatform.v1.EncryptionSpec.kms_key_name
	KMSKeyName *string `json:"kmsKeyName,omitempty"`
}

// +kcc:proto=google.cloud.aiplatform.v1.EnvVar
type EnvVar struct {
	// Required. Name of the environment variable. Must be a valid C identifier.
	// +kcc:proto:field=google.cloud.aiplatform.v1.EnvVar.name
	Name *string `json:"name,omitempty"`

	// Required. Variables that reference a $(VAR_NAME) are expanded
	//  using the previous defined environment variables in the container and
	//  any service environment variables. If a variable cannot be resolved,
	//  the reference in the input string will be unchanged. The $(VAR_NAME)
	//  syntax can be escaped with a double $$, ie: $$(VAR_NAME). Escaped
	//  references will never be expanded, regardless of whether the variable
	//  exists or not.
	// +kcc:proto:field=google.cloud.aiplatform.v1.EnvVar.value
	Value *string `json:"value,omitempty"`
}

// +kcc:proto=google.cloud.aiplatform.v1.Examples
type Examples struct {
	// The Cloud Storage input instances.
	// +kcc:proto:field=google.cloud.aiplatform.v1.Examples.example_gcs_source
	ExampleGCSSource *Examples_ExampleGcsSource `json:"exampleGCSSource,omitempty"`

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

	// ListValue struct refers to Value struct and vice versa, causing it to be recursive.
	// The recursive structure is causing CRD generation failure
	// OutputIndices *ListValue `json:"outputIndices,omitempty"`
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

// +kcc:proto=google.cloud.aiplatform.v1.GenieSource
type GenieSource struct {
	// Required. The public base model URI.
	// +kcc:proto:field=google.cloud.aiplatform.v1.GenieSource.base_model_uri
	BaseModelURI *string `json:"baseModelURI,omitempty"`
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

// +kcc:proto=google.cloud.aiplatform.v1.Model.BaseModelSource
type Model_BaseModelSource struct {
	// Source information of Model Garden models.
	// +kcc:proto:field=google.cloud.aiplatform.v1.Model.BaseModelSource.model_garden_source
	ModelGardenSource *ModelGardenSource `json:"modelGardenSource,omitempty"`

	// Information about the base model of Genie models.
	// +kcc:proto:field=google.cloud.aiplatform.v1.Model.BaseModelSource.genie_source
	GenieSource *GenieSource `json:"genieSource,omitempty"`
}

// +kcc:proto=google.cloud.aiplatform.v1.Model.DataStats
type Model_DataStats struct {
	// Number of DataItems that were used for training this Model.
	// +kcc:proto:field=google.cloud.aiplatform.v1.Model.DataStats.training_data_items_count
	TrainingDataItemsCount *int64 `json:"trainingDataItemsCount,omitempty"`

	// Number of DataItems that were used for validating this Model during
	//  training.
	// +kcc:proto:field=google.cloud.aiplatform.v1.Model.DataStats.validation_data_items_count
	ValidationDataItemsCount *int64 `json:"validationDataItemsCount,omitempty"`

	// Number of DataItems that were used for evaluating this Model. If the
	//  Model is evaluated multiple times, this will be the number of test
	//  DataItems used by the first evaluation. If the Model is not evaluated,
	//  the number is 0.
	// +kcc:proto:field=google.cloud.aiplatform.v1.Model.DataStats.test_data_items_count
	TestDataItemsCount *int64 `json:"testDataItemsCount,omitempty"`

	// Number of Annotations that are used for training this Model.
	// +kcc:proto:field=google.cloud.aiplatform.v1.Model.DataStats.training_annotations_count
	TrainingAnnotationsCount *int64 `json:"trainingAnnotationsCount,omitempty"`

	// Number of Annotations that are used for validating this Model during
	//  training.
	// +kcc:proto:field=google.cloud.aiplatform.v1.Model.DataStats.validation_annotations_count
	ValidationAnnotationsCount *int64 `json:"validationAnnotationsCount,omitempty"`

	// Number of Annotations that are used for evaluating this Model. If the
	//  Model is evaluated multiple times, this will be the number of test
	//  Annotations used by the first evaluation. If the Model is not evaluated,
	//  the number is 0.
	// +kcc:proto:field=google.cloud.aiplatform.v1.Model.DataStats.test_annotations_count
	TestAnnotationsCount *int64 `json:"testAnnotationsCount,omitempty"`
}

// +kcc:proto=google.cloud.aiplatform.v1.Model.ExportFormat
type Model_ExportFormat struct {
}

// +kcc:proto=google.cloud.aiplatform.v1.Model.OriginalModelInfo
type Model_OriginalModelInfo struct {
}

// +kcc:proto=google.cloud.aiplatform.v1.ModelContainerSpec
type ModelContainerSpec struct {
	// Required. Immutable. URI of the Docker image to be used as the custom
	//  container for serving predictions. This URI must identify an image in
	//  Artifact Registry or Container Registry. Learn more about the [container
	//  publishing
	//  requirements](https://cloud.google.com/vertex-ai/docs/predictions/custom-container-requirements#publishing),
	//  including permissions requirements for the Vertex AI Service Agent.
	//
	//  The container image is ingested upon
	//  [ModelService.UploadModel][google.cloud.aiplatform.v1.ModelService.UploadModel],
	//  stored internally, and this original path is afterwards not used.
	//
	//  To learn about the requirements for the Docker image itself, see
	//  [Custom container
	//  requirements](https://cloud.google.com/vertex-ai/docs/predictions/custom-container-requirements#).
	//
	//  You can use the URI to one of Vertex AI's [pre-built container images for
	//  prediction](https://cloud.google.com/vertex-ai/docs/predictions/pre-built-containers)
	//  in this field.
	// +kcc:proto:field=google.cloud.aiplatform.v1.ModelContainerSpec.image_uri
	ImageURI *string `json:"imageURI,omitempty"`

	// Immutable. Specifies the command that runs when the container starts. This
	//  overrides the container's
	//  [ENTRYPOINT](https://docs.docker.com/engine/reference/builder/#entrypoint).
	//  Specify this field as an array of executable and arguments, similar to a
	//  Docker `ENTRYPOINT`'s "exec" form, not its "shell" form.
	//
	//  If you do not specify this field, then the container's `ENTRYPOINT` runs,
	//  in conjunction with the
	//  [args][google.cloud.aiplatform.v1.ModelContainerSpec.args] field or the
	//  container's [`CMD`](https://docs.docker.com/engine/reference/builder/#cmd),
	//  if either exists. If this field is not specified and the container does not
	//  have an `ENTRYPOINT`, then refer to the Docker documentation about [how
	//  `CMD` and `ENTRYPOINT`
	//  interact](https://docs.docker.com/engine/reference/builder/#understand-how-cmd-and-entrypoint-interact).
	//
	//  If you specify this field, then you can also specify the `args` field to
	//  provide additional arguments for this command. However, if you specify this
	//  field, then the container's `CMD` is ignored. See the
	//  [Kubernetes documentation about how the
	//  `command` and `args` fields interact with a container's `ENTRYPOINT` and
	//  `CMD`](https://kubernetes.io/docs/tasks/inject-data-application/define-command-argument-container/#notes).
	//
	//  In this field, you can reference [environment variables set by Vertex
	//  AI](https://cloud.google.com/vertex-ai/docs/predictions/custom-container-requirements#aip-variables)
	//  and environment variables set in the
	//  [env][google.cloud.aiplatform.v1.ModelContainerSpec.env] field. You cannot
	//  reference environment variables set in the Docker image. In order for
	//  environment variables to be expanded, reference them by using the following
	//  syntax: <code>$(<var>VARIABLE_NAME</var>)</code> Note that this differs
	//  from Bash variable expansion, which does not use parentheses. If a variable
	//  cannot be resolved, the reference in the input string is used unchanged. To
	//  avoid variable expansion, you can escape this syntax with `$$`; for
	//  example: <code>$$(<var>VARIABLE_NAME</var>)</code> This field corresponds
	//  to the `command` field of the Kubernetes Containers [v1 core
	//  API](https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.23/#container-v1-core).
	// +kcc:proto:field=google.cloud.aiplatform.v1.ModelContainerSpec.command
	Command []string `json:"command,omitempty"`

	// Immutable. Specifies arguments for the command that runs when the container
	//  starts. This overrides the container's
	//  [`CMD`](https://docs.docker.com/engine/reference/builder/#cmd). Specify
	//  this field as an array of executable and arguments, similar to a Docker
	//  `CMD`'s "default parameters" form.
	//
	//  If you don't specify this field but do specify the
	//  [command][google.cloud.aiplatform.v1.ModelContainerSpec.command] field,
	//  then the command from the `command` field runs without any additional
	//  arguments. See the [Kubernetes documentation about how the `command` and
	//  `args` fields interact with a container's `ENTRYPOINT` and
	//  `CMD`](https://kubernetes.io/docs/tasks/inject-data-application/define-command-argument-container/#notes).
	//
	//  If you don't specify this field and don't specify the `command` field,
	//  then the container's
	//  [`ENTRYPOINT`](https://docs.docker.com/engine/reference/builder/#cmd) and
	//  `CMD` determine what runs based on their default behavior. See the Docker
	//  documentation about [how `CMD` and `ENTRYPOINT`
	//  interact](https://docs.docker.com/engine/reference/builder/#understand-how-cmd-and-entrypoint-interact).
	//
	//  In this field, you can reference [environment variables
	//  set by Vertex
	//  AI](https://cloud.google.com/vertex-ai/docs/predictions/custom-container-requirements#aip-variables)
	//  and environment variables set in the
	//  [env][google.cloud.aiplatform.v1.ModelContainerSpec.env] field. You cannot
	//  reference environment variables set in the Docker image. In order for
	//  environment variables to be expanded, reference them by using the following
	//  syntax: <code>$(<var>VARIABLE_NAME</var>)</code> Note that this differs
	//  from Bash variable expansion, which does not use parentheses. If a variable
	//  cannot be resolved, the reference in the input string is used unchanged. To
	//  avoid variable expansion, you can escape this syntax with `$$`; for
	//  example: <code>$$(<var>VARIABLE_NAME</var>)</code> This field corresponds
	//  to the `args` field of the Kubernetes Containers [v1 core
	//  API](https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.23/#container-v1-core).
	// +kcc:proto:field=google.cloud.aiplatform.v1.ModelContainerSpec.args
	Args []string `json:"args,omitempty"`

	// Immutable. List of environment variables to set in the container. After the
	//  container starts running, code running in the container can read these
	//  environment variables.
	//
	//  Additionally, the
	//  [command][google.cloud.aiplatform.v1.ModelContainerSpec.command] and
	//  [args][google.cloud.aiplatform.v1.ModelContainerSpec.args] fields can
	//  reference these variables. Later entries in this list can also reference
	//  earlier entries. For example, the following example sets the variable
	//  `VAR_2` to have the value `foo bar`:
	//
	//  ```json
	//  [
	//    {
	//      "name": "VAR_1",
	//      "value": "foo"
	//    },
	//    {
	//      "name": "VAR_2",
	//      "value": "$(VAR_1) bar"
	//    }
	//  ]
	//  ```
	//
	//  If you switch the order of the variables in the example, then the expansion
	//  does not occur.
	//
	//  This field corresponds to the `env` field of the Kubernetes Containers
	//  [v1 core
	//  API](https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.23/#container-v1-core).
	// +kcc:proto:field=google.cloud.aiplatform.v1.ModelContainerSpec.env
	Env []EnvVar `json:"env,omitempty"`

	// Immutable. List of ports to expose from the container. Vertex AI sends any
	//  prediction requests that it receives to the first port on this list. Vertex
	//  AI also sends
	//  [liveness and health
	//  checks](https://cloud.google.com/vertex-ai/docs/predictions/custom-container-requirements#liveness)
	//  to this port.
	//
	//  If you do not specify this field, it defaults to following value:
	//
	//  ```json
	//  [
	//    {
	//      "containerPort": 8080
	//    }
	//  ]
	//  ```
	//
	//  Vertex AI does not use ports other than the first one listed. This field
	//  corresponds to the `ports` field of the Kubernetes Containers
	//  [v1 core
	//  API](https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.23/#container-v1-core).
	// +kcc:proto:field=google.cloud.aiplatform.v1.ModelContainerSpec.ports
	Ports []Port `json:"ports,omitempty"`

	// Immutable. HTTP path on the container to send prediction requests to.
	//  Vertex AI forwards requests sent using
	//  [projects.locations.endpoints.predict][google.cloud.aiplatform.v1.PredictionService.Predict]
	//  to this path on the container's IP address and port. Vertex AI then returns
	//  the container's response in the API response.
	//
	//  For example, if you set this field to `/foo`, then when Vertex AI
	//  receives a prediction request, it forwards the request body in a POST
	//  request to the `/foo` path on the port of your container specified by the
	//  first value of this `ModelContainerSpec`'s
	//  [ports][google.cloud.aiplatform.v1.ModelContainerSpec.ports] field.
	//
	//  If you don't specify this field, it defaults to the following value when
	//  you [deploy this Model to an
	//  Endpoint][google.cloud.aiplatform.v1.EndpointService.DeployModel]:
	//  <code>/v1/endpoints/<var>ENDPOINT</var>/deployedModels/<var>DEPLOYED_MODEL</var>:predict</code>
	//  The placeholders in this value are replaced as follows:
	//
	//  * <var>ENDPOINT</var>: The last segment (following `endpoints/`)of the
	//    Endpoint.name][] field of the Endpoint where this Model has been
	//    deployed. (Vertex AI makes this value available to your container code
	//    as the [`AIP_ENDPOINT_ID` environment
	//   variable](https://cloud.google.com/vertex-ai/docs/predictions/custom-container-requirements#aip-variables).)
	//
	//  * <var>DEPLOYED_MODEL</var>:
	//  [DeployedModel.id][google.cloud.aiplatform.v1.DeployedModel.id] of the
	//  `DeployedModel`.
	//    (Vertex AI makes this value available to your container code
	//    as the [`AIP_DEPLOYED_MODEL_ID` environment
	//    variable](https://cloud.google.com/vertex-ai/docs/predictions/custom-container-requirements#aip-variables).)
	// +kcc:proto:field=google.cloud.aiplatform.v1.ModelContainerSpec.predict_route
	PredictRoute *string `json:"predictRoute,omitempty"`

	// Immutable. HTTP path on the container to send health checks to. Vertex AI
	//  intermittently sends GET requests to this path on the container's IP
	//  address and port to check that the container is healthy. Read more about
	//  [health
	//  checks](https://cloud.google.com/vertex-ai/docs/predictions/custom-container-requirements#health).
	//
	//  For example, if you set this field to `/bar`, then Vertex AI
	//  intermittently sends a GET request to the `/bar` path on the port of your
	//  container specified by the first value of this `ModelContainerSpec`'s
	//  [ports][google.cloud.aiplatform.v1.ModelContainerSpec.ports] field.
	//
	//  If you don't specify this field, it defaults to the following value when
	//  you [deploy this Model to an
	//  Endpoint][google.cloud.aiplatform.v1.EndpointService.DeployModel]:
	//  <code>/v1/endpoints/<var>ENDPOINT</var>/deployedModels/<var>DEPLOYED_MODEL</var>:predict</code>
	//  The placeholders in this value are replaced as follows:
	//
	//  * <var>ENDPOINT</var>: The last segment (following `endpoints/`)of the
	//    Endpoint.name][] field of the Endpoint where this Model has been
	//    deployed. (Vertex AI makes this value available to your container code
	//    as the [`AIP_ENDPOINT_ID` environment
	//    variable](https://cloud.google.com/vertex-ai/docs/predictions/custom-container-requirements#aip-variables).)
	//
	//  * <var>DEPLOYED_MODEL</var>:
	//  [DeployedModel.id][google.cloud.aiplatform.v1.DeployedModel.id] of the
	//  `DeployedModel`.
	//    (Vertex AI makes this value available to your container code as the
	//    [`AIP_DEPLOYED_MODEL_ID` environment
	//    variable](https://cloud.google.com/vertex-ai/docs/predictions/custom-container-requirements#aip-variables).)
	// +kcc:proto:field=google.cloud.aiplatform.v1.ModelContainerSpec.health_route
	HealthRoute *string `json:"healthRoute,omitempty"`

	// Immutable. List of ports to expose from the container. Vertex AI sends gRPC
	//  prediction requests that it receives to the first port on this list. Vertex
	//  AI also sends liveness and health checks to this port.
	//
	//  If you do not specify this field, gRPC requests to the container will be
	//  disabled.
	//
	//  Vertex AI does not use ports other than the first one listed. This field
	//  corresponds to the `ports` field of the Kubernetes Containers v1 core API.
	// +kcc:proto:field=google.cloud.aiplatform.v1.ModelContainerSpec.grpc_ports
	GrpcPorts []Port `json:"grpcPorts,omitempty"`

	// Immutable. Deployment timeout.
	//  Limit for deployment timeout is 2 hours.
	// +kcc:proto:field=google.cloud.aiplatform.v1.ModelContainerSpec.deployment_timeout
	DeploymentTimeout *string `json:"deploymentTimeout,omitempty"`

	// Immutable. The amount of the VM memory to reserve as the shared memory for
	//  the model in megabytes.
	// +kcc:proto:field=google.cloud.aiplatform.v1.ModelContainerSpec.shared_memory_size_mb
	SharedMemorySizeMb *int64 `json:"sharedMemorySizeMb,omitempty"`

	// Immutable. Specification for Kubernetes startup probe.
	// +kcc:proto:field=google.cloud.aiplatform.v1.ModelContainerSpec.startup_probe
	StartupProbe *Probe `json:"startupProbe,omitempty"`

	// Immutable. Specification for Kubernetes readiness probe.
	// +kcc:proto:field=google.cloud.aiplatform.v1.ModelContainerSpec.health_probe
	HealthProbe *Probe `json:"healthProbe,omitempty"`
}

// +kcc:proto=google.cloud.aiplatform.v1.ModelGardenSource
type ModelGardenSource struct {
	// Required. The model garden source model resource name.
	// +kcc:proto:field=google.cloud.aiplatform.v1.ModelGardenSource.public_model_name
	PublicModelName *string `json:"publicModelName,omitempty"`
}

// +kcc:proto=google.cloud.aiplatform.v1.ModelSourceInfo
type ModelSourceInfo struct {
	// Type of the model source.
	// +kcc:proto:field=google.cloud.aiplatform.v1.ModelSourceInfo.source_type
	SourceType *string `json:"sourceType,omitempty"`

	// If this Model is copy of another Model. If true then
	//  [source_type][google.cloud.aiplatform.v1.ModelSourceInfo.source_type]
	//  pertains to the original.
	// +kcc:proto:field=google.cloud.aiplatform.v1.ModelSourceInfo.copy
	Copy *bool `json:"copy,omitempty"`
}

// +kcc:proto=google.cloud.aiplatform.v1.Port
type Port struct {
	// The number of the port to expose on the pod's IP address.
	//  Must be a valid port number, between 1 and 65535 inclusive.
	// +kcc:proto:field=google.cloud.aiplatform.v1.Port.container_port
	ContainerPort *int32 `json:"containerPort,omitempty"`
}

// +kcc:proto=google.cloud.aiplatform.v1.PredictSchemata
type PredictSchemata struct {
	// Immutable. Points to a YAML file stored on Google Cloud Storage describing
	//  the format of a single instance, which are used in
	//  [PredictRequest.instances][google.cloud.aiplatform.v1.PredictRequest.instances],
	//  [ExplainRequest.instances][google.cloud.aiplatform.v1.ExplainRequest.instances]
	//  and
	//  [BatchPredictionJob.input_config][google.cloud.aiplatform.v1.BatchPredictionJob.input_config].
	//  The schema is defined as an OpenAPI 3.0.2 [Schema
	//  Object](https://github.com/OAI/OpenAPI-Specification/blob/main/versions/3.0.2.md#schemaObject).
	//  AutoML Models always have this field populated by Vertex AI.
	//  Note: The URI given on output will be immutable and probably different,
	//  including the URI scheme, than the one given on input. The output URI will
	//  point to a location where the user only has a read access.
	// +kcc:proto:field=google.cloud.aiplatform.v1.PredictSchemata.instance_schema_uri
	InstanceSchemaURI *string `json:"instanceSchemaURI,omitempty"`

	// Immutable. Points to a YAML file stored on Google Cloud Storage describing
	//  the parameters of prediction and explanation via
	//  [PredictRequest.parameters][google.cloud.aiplatform.v1.PredictRequest.parameters],
	//  [ExplainRequest.parameters][google.cloud.aiplatform.v1.ExplainRequest.parameters]
	//  and
	//  [BatchPredictionJob.model_parameters][google.cloud.aiplatform.v1.BatchPredictionJob.model_parameters].
	//  The schema is defined as an OpenAPI 3.0.2 [Schema
	//  Object](https://github.com/OAI/OpenAPI-Specification/blob/main/versions/3.0.2.md#schemaObject).
	//  AutoML Models always have this field populated by Vertex AI, if no
	//  parameters are supported, then it is set to an empty string.
	//  Note: The URI given on output will be immutable and probably different,
	//  including the URI scheme, than the one given on input. The output URI will
	//  point to a location where the user only has a read access.
	// +kcc:proto:field=google.cloud.aiplatform.v1.PredictSchemata.parameters_schema_uri
	ParametersSchemaURI *string `json:"parametersSchemaURI,omitempty"`

	// Immutable. Points to a YAML file stored on Google Cloud Storage describing
	//  the format of a single prediction produced by this Model, which are
	//  returned via
	//  [PredictResponse.predictions][google.cloud.aiplatform.v1.PredictResponse.predictions],
	//  [ExplainResponse.explanations][google.cloud.aiplatform.v1.ExplainResponse.explanations],
	//  and
	//  [BatchPredictionJob.output_config][google.cloud.aiplatform.v1.BatchPredictionJob.output_config].
	//  The schema is defined as an OpenAPI 3.0.2 [Schema
	//  Object](https://github.com/OAI/OpenAPI-Specification/blob/main/versions/3.0.2.md#schemaObject).
	//  AutoML Models always have this field populated by Vertex AI.
	//  Note: The URI given on output will be immutable and probably different,
	//  including the URI scheme, than the one given on input. The output URI will
	//  point to a location where the user only has a read access.
	// +kcc:proto:field=google.cloud.aiplatform.v1.PredictSchemata.prediction_schema_uri
	PredictionSchemaURI *string `json:"predictionSchemaURI,omitempty"`
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

// +kcc:proto=google.cloud.aiplatform.v1.Probe
type Probe struct {
	// ExecAction probes the health of a container by executing a command.
	// +kcc:proto:field=google.cloud.aiplatform.v1.Probe.exec
	Exec *Probe_ExecAction `json:"exec,omitempty"`

	// How often (in seconds) to perform the probe. Default to 10 seconds.
	//  Minimum value is 1. Must be less than timeout_seconds.
	//
	//  Maps to Kubernetes probe argument 'periodSeconds'.
	// +kcc:proto:field=google.cloud.aiplatform.v1.Probe.period_seconds
	PeriodSeconds *int32 `json:"periodSeconds,omitempty"`

	// Number of seconds after which the probe times out. Defaults to 1 second.
	//  Minimum value is 1. Must be greater or equal to period_seconds.
	//
	//  Maps to Kubernetes probe argument 'timeoutSeconds'.
	// +kcc:proto:field=google.cloud.aiplatform.v1.Probe.timeout_seconds
	TimeoutSeconds *int32 `json:"timeoutSeconds,omitempty"`
}

// +kcc:proto=google.cloud.aiplatform.v1.Probe.ExecAction
type Probe_ExecAction struct {
	// Command is the command line to execute inside the container, the working
	//  directory for the command is root ('/') in the container's filesystem.
	//  The command is simply exec'd, it is not run inside a shell, so
	//  traditional shell instructions ('|', etc) won't work. To use a shell, you
	//  need to explicitly call out to that shell. Exit status of 0 is treated as
	//  live/healthy and non-zero is unhealthy.
	// +kcc:proto:field=google.cloud.aiplatform.v1.Probe.ExecAction.command
	Command []string `json:"command,omitempty"`
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

	// ListValue refers to Value struct and vice versa, causing it to be recursive.
	// The recursive structure is causing CRD generation failure
	// ListValue []Value `json:"listValue,omitempty"`
}

// +kcc:proto=google.cloud.aiplatform.v1.Model.ExportFormat
type Model_ExportFormatObservedState struct {
	// Output only. The ID of the export format.
	//  The possible format IDs are:
	//
	//  * `tflite`
	//  Used for Android mobile devices.
	//
	//  * `edgetpu-tflite`
	//  Used for [Edge TPU](https://cloud.google.com/edge-tpu/) devices.
	//
	//  * `tf-saved-model`
	//  A tensorflow model in SavedModel format.
	//
	//  * `tf-js`
	//  A [TensorFlow.js](https://www.tensorflow.org/js) model that can be used
	//  in the browser and in Node.js using JavaScript.
	//
	//  * `core-ml`
	//  Used for iOS mobile devices.
	//
	//  * `custom-trained`
	//  A Model that was uploaded or trained by custom code.
	// +kcc:proto:field=google.cloud.aiplatform.v1.Model.ExportFormat.id
	ID *string `json:"id,omitempty"`

	// Output only. The content of this Model that may be exported.
	// +kcc:proto:field=google.cloud.aiplatform.v1.Model.ExportFormat.exportable_contents
	ExportableContents []string `json:"exportableContents,omitempty"`
}

// +kcc:proto=google.cloud.aiplatform.v1.Model.OriginalModelInfo
type Model_OriginalModelInfoObservedState struct {
	// Output only. The resource name of the Model this Model is a copy of,
	//  including the revision. Format:
	//  `projects/{project}/locations/{location}/models/{model_id}@{version_id}`
	// +kcc:proto:field=google.cloud.aiplatform.v1.Model.OriginalModelInfo.model
	Model *string `json:"model,omitempty"`
}
