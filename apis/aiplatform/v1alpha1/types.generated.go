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

// +kcc:proto=google.cloud.aiplatform.v1.ModelEvaluationSlice
type ModelEvaluationSlice struct {
}

// +kcc:proto=google.cloud.aiplatform.v1.ModelEvaluationSlice.Slice
type ModelEvaluationSlice_Slice struct {
}

// +kcc:proto=google.cloud.aiplatform.v1.ModelEvaluationSlice.Slice.SliceSpec
type ModelEvaluationSlice_Slice_SliceSpec struct {

	// TODO: unsupported map type with key string and value message

}

// +kcc:proto=google.cloud.aiplatform.v1.ModelEvaluationSlice.Slice.SliceSpec.Range
type ModelEvaluationSlice_Slice_SliceSpec_Range struct {
	// Inclusive low value for the range.
	// +kcc:proto:field=google.cloud.aiplatform.v1.ModelEvaluationSlice.Slice.SliceSpec.Range.low
	Low *float32 `json:"low,omitempty"`

	// Exclusive high value for the range.
	// +kcc:proto:field=google.cloud.aiplatform.v1.ModelEvaluationSlice.Slice.SliceSpec.Range.high
	High *float32 `json:"high,omitempty"`
}

// +kcc:proto=google.cloud.aiplatform.v1.ModelEvaluationSlice.Slice.SliceSpec.SliceConfig
type ModelEvaluationSlice_Slice_SliceSpec_SliceConfig struct {
	// A unique specific value for a given feature.
	//  Example: `{ "value": { "string_value": "12345" } }`
	// +kcc:proto:field=google.cloud.aiplatform.v1.ModelEvaluationSlice.Slice.SliceSpec.SliceConfig.value
	Value *ModelEvaluationSlice_Slice_SliceSpec_Value `json:"value,omitempty"`

	// A range of values for a numerical feature.
	//  Example: `{"range":{"low":10000.0,"high":50000.0}}`
	//  will capture 12345 and 23334 in the slice.
	// +kcc:proto:field=google.cloud.aiplatform.v1.ModelEvaluationSlice.Slice.SliceSpec.SliceConfig.range
	Range *ModelEvaluationSlice_Slice_SliceSpec_Range `json:"range,omitempty"`

	// If all_values is set to true, then all possible labels of the keyed
	//  feature will have another slice computed.
	//  Example: `{"all_values":{"value":true}}`
	// +kcc:proto:field=google.cloud.aiplatform.v1.ModelEvaluationSlice.Slice.SliceSpec.SliceConfig.all_values
	AllValues *bool `json:"allValues,omitempty"`
}

// +kcc:proto=google.cloud.aiplatform.v1.ModelEvaluationSlice.Slice.SliceSpec.Value
type ModelEvaluationSlice_Slice_SliceSpec_Value struct {
	// String type.
	// +kcc:proto:field=google.cloud.aiplatform.v1.ModelEvaluationSlice.Slice.SliceSpec.Value.string_value
	StringValue *string `json:"stringValue,omitempty"`

	// Float type.
	// +kcc:proto:field=google.cloud.aiplatform.v1.ModelEvaluationSlice.Slice.SliceSpec.Value.float_value
	FloatValue *float32 `json:"floatValue,omitempty"`
}

// +kcc:proto=google.cloud.aiplatform.v1.ModelExplanation
type ModelExplanation struct {
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

// +kcc:proto=google.cloud.aiplatform.v1.ModelEvaluationSlice
type ModelEvaluationSliceObservedState struct {
	// Output only. The resource name of the ModelEvaluationSlice.
	// +kcc:proto:field=google.cloud.aiplatform.v1.ModelEvaluationSlice.name
	Name *string `json:"name,omitempty"`

	// Output only. The slice of the test data that is used to evaluate the Model.
	// +kcc:proto:field=google.cloud.aiplatform.v1.ModelEvaluationSlice.slice
	Slice *ModelEvaluationSlice_Slice `json:"slice,omitempty"`

	// Output only. Points to a YAML file stored on Google Cloud Storage
	//  describing the
	//  [metrics][google.cloud.aiplatform.v1.ModelEvaluationSlice.metrics] of this
	//  ModelEvaluationSlice. The schema is defined as an OpenAPI 3.0.2 [Schema
	//  Object](https://github.com/OAI/OpenAPI-Specification/blob/main/versions/3.0.2.md#schemaObject).
	// +kcc:proto:field=google.cloud.aiplatform.v1.ModelEvaluationSlice.metrics_schema_uri
	MetricsSchemaURI *string `json:"metricsSchemaURI,omitempty"`

	// Output only. Sliced evaluation metrics of the Model. The schema of the
	//  metrics is stored in
	//  [metrics_schema_uri][google.cloud.aiplatform.v1.ModelEvaluationSlice.metrics_schema_uri]
	// +kcc:proto:field=google.cloud.aiplatform.v1.ModelEvaluationSlice.metrics
	Metrics *Value `json:"metrics,omitempty"`

	// Output only. Timestamp when this ModelEvaluationSlice was created.
	// +kcc:proto:field=google.cloud.aiplatform.v1.ModelEvaluationSlice.create_time
	CreateTime *string `json:"createTime,omitempty"`

	// Output only. Aggregated explanation metrics for the Model's prediction
	//  output over the data this ModelEvaluation uses. This field is populated
	//  only if the Model is evaluated with explanations, and only for tabular
	//  Models.
	// +kcc:proto:field=google.cloud.aiplatform.v1.ModelEvaluationSlice.model_explanation
	ModelExplanation *ModelExplanation `json:"modelExplanation,omitempty"`
}

// +kcc:proto=google.cloud.aiplatform.v1.ModelEvaluationSlice.Slice
type ModelEvaluationSlice_SliceObservedState struct {
	// Output only. The dimension of the slice.
	//  Well-known dimensions are:
	//    * `annotationSpec`: This slice is on the test data that has either
	//      ground truth or prediction with
	//      [AnnotationSpec.display_name][google.cloud.aiplatform.v1.AnnotationSpec.display_name]
	//      equals to
	//      [value][google.cloud.aiplatform.v1.ModelEvaluationSlice.Slice.value].
	//    * `slice`: This slice is a user customized slice defined by its
	//      SliceSpec.
	// +kcc:proto:field=google.cloud.aiplatform.v1.ModelEvaluationSlice.Slice.dimension
	Dimension *string `json:"dimension,omitempty"`

	// Output only. The value of the dimension in this slice.
	// +kcc:proto:field=google.cloud.aiplatform.v1.ModelEvaluationSlice.Slice.value
	Value *string `json:"value,omitempty"`

	// Output only. Specification for how the data was sliced.
	// +kcc:proto:field=google.cloud.aiplatform.v1.ModelEvaluationSlice.Slice.slice_spec
	SliceSpec *ModelEvaluationSlice_Slice_SliceSpec `json:"sliceSpec,omitempty"`
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
