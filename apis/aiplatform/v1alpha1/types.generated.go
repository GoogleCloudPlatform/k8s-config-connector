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


// +kcc:proto=google.cloud.aiplatform.v1.BatchDedicatedResources
type BatchDedicatedResources struct {
	// Required. Immutable. The specification of a single machine.
	// +kcc:proto:field=google.cloud.aiplatform.v1.BatchDedicatedResources.machine_spec
	MachineSpec *MachineSpec `json:"machineSpec,omitempty"`

	// Immutable. The number of machine replicas used at the start of the batch
	//  operation. If not set, Vertex AI decides starting number, not greater than
	//  [max_replica_count][google.cloud.aiplatform.v1.BatchDedicatedResources.max_replica_count]
	// +kcc:proto:field=google.cloud.aiplatform.v1.BatchDedicatedResources.starting_replica_count
	StartingReplicaCount *int32 `json:"startingReplicaCount,omitempty"`

	// Immutable. The maximum number of machine replicas the batch operation may
	//  be scaled to. The default value is 10.
	// +kcc:proto:field=google.cloud.aiplatform.v1.BatchDedicatedResources.max_replica_count
	MaxReplicaCount *int32 `json:"maxReplicaCount,omitempty"`
}

// +kcc:proto=google.cloud.aiplatform.v1.BatchPredictionJob
type BatchPredictionJob struct {

	// Required. The user-defined name of this BatchPredictionJob.
	// +kcc:proto:field=google.cloud.aiplatform.v1.BatchPredictionJob.display_name
	DisplayName *string `json:"displayName,omitempty"`

	// The name of the Model resource that produces the predictions via this job,
	//  must share the same ancestor Location.
	//  Starting this job has no impact on any existing deployments of the Model
	//  and their resources.
	//  Exactly one of model and unmanaged_container_model must be set.
	//
	//  The model resource name may contain version id or version alias to specify
	//  the version.
	//   Example: `projects/{project}/locations/{location}/models/{model}@2`
	//               or
	//             `projects/{project}/locations/{location}/models/{model}@golden`
	//  if no version is specified, the default version will be deployed.
	//
	//  The model resource could also be a publisher model.
	//   Example: `publishers/{publisher}/models/{model}`
	//               or
	//            `projects/{project}/locations/{location}/publishers/{publisher}/models/{model}`
	// +kcc:proto:field=google.cloud.aiplatform.v1.BatchPredictionJob.model
	Model *string `json:"model,omitempty"`

	// Contains model information necessary to perform batch prediction without
	//  requiring uploading to model registry.
	//  Exactly one of model and unmanaged_container_model must be set.
	// +kcc:proto:field=google.cloud.aiplatform.v1.BatchPredictionJob.unmanaged_container_model
	UnmanagedContainerModel *UnmanagedContainerModel `json:"unmanagedContainerModel,omitempty"`

	// Required. Input configuration of the instances on which predictions are
	//  performed. The schema of any single instance may be specified via the
	//  [Model's][google.cloud.aiplatform.v1.BatchPredictionJob.model]
	//  [PredictSchemata's][google.cloud.aiplatform.v1.Model.predict_schemata]
	//  [instance_schema_uri][google.cloud.aiplatform.v1.PredictSchemata.instance_schema_uri].
	// +kcc:proto:field=google.cloud.aiplatform.v1.BatchPredictionJob.input_config
	InputConfig *BatchPredictionJob_InputConfig `json:"inputConfig,omitempty"`

	// Configuration for how to convert batch prediction input instances to the
	//  prediction instances that are sent to the Model.
	// +kcc:proto:field=google.cloud.aiplatform.v1.BatchPredictionJob.instance_config
	InstanceConfig *BatchPredictionJob_InstanceConfig `json:"instanceConfig,omitempty"`

	// The parameters that govern the predictions. The schema of the parameters
	//  may be specified via the
	//  [Model's][google.cloud.aiplatform.v1.BatchPredictionJob.model]
	//  [PredictSchemata's][google.cloud.aiplatform.v1.Model.predict_schemata]
	//  [parameters_schema_uri][google.cloud.aiplatform.v1.PredictSchemata.parameters_schema_uri].
	// +kcc:proto:field=google.cloud.aiplatform.v1.BatchPredictionJob.model_parameters
	ModelParameters *Value `json:"modelParameters,omitempty"`

	// Required. The Configuration specifying where output predictions should
	//  be written.
	//  The schema of any single prediction may be specified as a concatenation
	//  of [Model's][google.cloud.aiplatform.v1.BatchPredictionJob.model]
	//  [PredictSchemata's][google.cloud.aiplatform.v1.Model.predict_schemata]
	//  [instance_schema_uri][google.cloud.aiplatform.v1.PredictSchemata.instance_schema_uri]
	//  and
	//  [prediction_schema_uri][google.cloud.aiplatform.v1.PredictSchemata.prediction_schema_uri].
	// +kcc:proto:field=google.cloud.aiplatform.v1.BatchPredictionJob.output_config
	OutputConfig *BatchPredictionJob_OutputConfig `json:"outputConfig,omitempty"`

	// The config of resources used by the Model during the batch prediction. If
	//  the Model
	//  [supports][google.cloud.aiplatform.v1.Model.supported_deployment_resources_types]
	//  DEDICATED_RESOURCES this config may be provided (and the job will use these
	//  resources), if the Model doesn't support AUTOMATIC_RESOURCES, this config
	//  must be provided.
	// +kcc:proto:field=google.cloud.aiplatform.v1.BatchPredictionJob.dedicated_resources
	DedicatedResources *BatchDedicatedResources `json:"dedicatedResources,omitempty"`

	// The service account that the DeployedModel's container runs as. If not
	//  specified, a system generated one will be used, which
	//  has minimal permissions and the custom container, if used, may not have
	//  enough permission to access other Google Cloud resources.
	//
	//  Users deploying the Model must have the `iam.serviceAccounts.actAs`
	//  permission on this service account.
	// +kcc:proto:field=google.cloud.aiplatform.v1.BatchPredictionJob.service_account
	ServiceAccount *string `json:"serviceAccount,omitempty"`

	// Immutable. Parameters configuring the batch behavior. Currently only
	//  applicable when
	//  [dedicated_resources][google.cloud.aiplatform.v1.BatchPredictionJob.dedicated_resources]
	//  are used (in other cases Vertex AI does the tuning itself).
	// +kcc:proto:field=google.cloud.aiplatform.v1.BatchPredictionJob.manual_batch_tuning_parameters
	ManualBatchTuningParameters *ManualBatchTuningParameters `json:"manualBatchTuningParameters,omitempty"`

	// Generate explanation with the batch prediction results.
	//
	//  When set to `true`, the batch prediction output changes based on the
	//  `predictions_format` field of the
	//  [BatchPredictionJob.output_config][google.cloud.aiplatform.v1.BatchPredictionJob.output_config]
	//  object:
	//
	//   * `bigquery`: output includes a column named `explanation`. The value
	//     is a struct that conforms to the
	//     [Explanation][google.cloud.aiplatform.v1.Explanation] object.
	//   * `jsonl`: The JSON objects on each line include an additional entry
	//     keyed `explanation`. The value of the entry is a JSON object that
	//     conforms to the [Explanation][google.cloud.aiplatform.v1.Explanation]
	//     object.
	//   * `csv`: Generating explanations for CSV format is not supported.
	//
	//  If this field is set to true, either the
	//  [Model.explanation_spec][google.cloud.aiplatform.v1.Model.explanation_spec]
	//  or
	//  [explanation_spec][google.cloud.aiplatform.v1.BatchPredictionJob.explanation_spec]
	//  must be populated.
	// +kcc:proto:field=google.cloud.aiplatform.v1.BatchPredictionJob.generate_explanation
	GenerateExplanation *bool `json:"generateExplanation,omitempty"`

	// Explanation configuration for this BatchPredictionJob. Can be
	//  specified only if
	//  [generate_explanation][google.cloud.aiplatform.v1.BatchPredictionJob.generate_explanation]
	//  is set to `true`.
	//
	//  This value overrides the value of
	//  [Model.explanation_spec][google.cloud.aiplatform.v1.Model.explanation_spec].
	//  All fields of
	//  [explanation_spec][google.cloud.aiplatform.v1.BatchPredictionJob.explanation_spec]
	//  are optional in the request. If a field of the
	//  [explanation_spec][google.cloud.aiplatform.v1.BatchPredictionJob.explanation_spec]
	//  object is not populated, the corresponding field of the
	//  [Model.explanation_spec][google.cloud.aiplatform.v1.Model.explanation_spec]
	//  object is inherited.
	// +kcc:proto:field=google.cloud.aiplatform.v1.BatchPredictionJob.explanation_spec
	ExplanationSpec *ExplanationSpec `json:"explanationSpec,omitempty"`

	// The labels with user-defined metadata to organize BatchPredictionJobs.
	//
	//  Label keys and values can be no longer than 64 characters
	//  (Unicode codepoints), can only contain lowercase letters, numeric
	//  characters, underscores and dashes. International characters are allowed.
	//
	//  See https://goo.gl/xmQnxf for more information and examples of labels.
	// +kcc:proto:field=google.cloud.aiplatform.v1.BatchPredictionJob.labels
	Labels map[string]string `json:"labels,omitempty"`

	// Customer-managed encryption key options for a BatchPredictionJob. If this
	//  is set, then all resources created by the BatchPredictionJob will be
	//  encrypted with the provided encryption key.
	// +kcc:proto:field=google.cloud.aiplatform.v1.BatchPredictionJob.encryption_spec
	EncryptionSpec *EncryptionSpec `json:"encryptionSpec,omitempty"`

	// For custom-trained Models and AutoML Tabular Models, the container of the
	//  DeployedModel instances will send `stderr` and `stdout` streams to
	//  Cloud Logging by default. Please note that the logs incur cost,
	//  which are subject to [Cloud Logging
	//  pricing](https://cloud.google.com/logging/pricing).
	//
	//  User can disable container logging by setting this flag to true.
	// +kcc:proto:field=google.cloud.aiplatform.v1.BatchPredictionJob.disable_container_logging
	DisableContainerLogging *bool `json:"disableContainerLogging,omitempty"`
}

// +kcc:proto=google.cloud.aiplatform.v1.BatchPredictionJob.InputConfig
type BatchPredictionJob_InputConfig struct {
	// The Cloud Storage location for the input instances.
	// +kcc:proto:field=google.cloud.aiplatform.v1.BatchPredictionJob.InputConfig.gcs_source
	GcsSource *GcsSource `json:"gcsSource,omitempty"`

	// The BigQuery location of the input table.
	//  The schema of the table should be in the format described by the given
	//  context OpenAPI Schema, if one is provided. The table may contain
	//  additional columns that are not described by the schema, and they will
	//  be ignored.
	// +kcc:proto:field=google.cloud.aiplatform.v1.BatchPredictionJob.InputConfig.bigquery_source
	BigquerySource *BigQuerySource `json:"bigquerySource,omitempty"`

	// Required. The format in which instances are given, must be one of the
	//  [Model's][google.cloud.aiplatform.v1.BatchPredictionJob.model]
	//  [supported_input_storage_formats][google.cloud.aiplatform.v1.Model.supported_input_storage_formats].
	// +kcc:proto:field=google.cloud.aiplatform.v1.BatchPredictionJob.InputConfig.instances_format
	InstancesFormat *string `json:"instancesFormat,omitempty"`
}

// +kcc:proto=google.cloud.aiplatform.v1.BatchPredictionJob.InstanceConfig
type BatchPredictionJob_InstanceConfig struct {
	// The format of the instance that the Model accepts. Vertex AI will
	//  convert compatible
	//  [batch prediction input instance
	//  formats][google.cloud.aiplatform.v1.BatchPredictionJob.InputConfig.instances_format]
	//  to the specified format.
	//
	//  Supported values are:
	//
	//  * `object`: Each input is converted to JSON object format.
	//      * For `bigquery`, each row is converted to an object.
	//      * For `jsonl`, each line of the JSONL input must be an object.
	//      * Does not apply to `csv`, `file-list`, `tf-record`, or
	//        `tf-record-gzip`.
	//
	//  * `array`: Each input is converted to JSON array format.
	//      * For `bigquery`, each row is converted to an array. The order
	//        of columns is determined by the BigQuery column order, unless
	//        [included_fields][google.cloud.aiplatform.v1.BatchPredictionJob.InstanceConfig.included_fields]
	//        is populated.
	//        [included_fields][google.cloud.aiplatform.v1.BatchPredictionJob.InstanceConfig.included_fields]
	//        must be populated for specifying field orders.
	//      * For `jsonl`, if each line of the JSONL input is an object,
	//        [included_fields][google.cloud.aiplatform.v1.BatchPredictionJob.InstanceConfig.included_fields]
	//        must be populated for specifying field orders.
	//      * Does not apply to `csv`, `file-list`, `tf-record`, or
	//        `tf-record-gzip`.
	//
	//  If not specified, Vertex AI converts the batch prediction input as
	//  follows:
	//
	//   * For `bigquery` and `csv`, the behavior is the same as `array`. The
	//     order of columns is the same as defined in the file or table, unless
	//     [included_fields][google.cloud.aiplatform.v1.BatchPredictionJob.InstanceConfig.included_fields]
	//     is populated.
	//   * For `jsonl`, the prediction instance format is determined by
	//     each line of the input.
	//   * For `tf-record`/`tf-record-gzip`, each record will be converted to
	//     an object in the format of `{"b64": <value>}`, where `<value>` is
	//     the Base64-encoded string of the content of the record.
	//   * For `file-list`, each file in the list will be converted to an
	//     object in the format of `{"b64": <value>}`, where `<value>` is
	//     the Base64-encoded string of the content of the file.
	// +kcc:proto:field=google.cloud.aiplatform.v1.BatchPredictionJob.InstanceConfig.instance_type
	InstanceType *string `json:"instanceType,omitempty"`

	// The name of the field that is considered as a key.
	//
	//  The values identified by the key field is not included in the transformed
	//  instances that is sent to the Model. This is similar to
	//  specifying this name of the field in
	//  [excluded_fields][google.cloud.aiplatform.v1.BatchPredictionJob.InstanceConfig.excluded_fields].
	//  In addition, the batch prediction output will not include the instances.
	//  Instead the output will only include the value of the key field, in a
	//  field named `key` in the output:
	//
	//   * For `jsonl` output format, the output will have a `key` field
	//     instead of the `instance` field.
	//   * For `csv`/`bigquery` output format, the output will have have a `key`
	//     column instead of the instance feature columns.
	//
	//  The input must be JSONL with objects at each line, CSV, BigQuery
	//  or TfRecord.
	// +kcc:proto:field=google.cloud.aiplatform.v1.BatchPredictionJob.InstanceConfig.key_field
	KeyField *string `json:"keyField,omitempty"`

	// Fields that will be included in the prediction instance that is
	//  sent to the Model.
	//
	//  If
	//  [instance_type][google.cloud.aiplatform.v1.BatchPredictionJob.InstanceConfig.instance_type]
	//  is `array`, the order of field names in included_fields also determines
	//  the order of the values in the array.
	//
	//  When included_fields is populated,
	//  [excluded_fields][google.cloud.aiplatform.v1.BatchPredictionJob.InstanceConfig.excluded_fields]
	//  must be empty.
	//
	//  The input must be JSONL with objects at each line, BigQuery
	//  or TfRecord.
	// +kcc:proto:field=google.cloud.aiplatform.v1.BatchPredictionJob.InstanceConfig.included_fields
	IncludedFields []string `json:"includedFields,omitempty"`

	// Fields that will be excluded in the prediction instance that is
	//  sent to the Model.
	//
	//  Excluded will be attached to the batch prediction output if
	//  [key_field][google.cloud.aiplatform.v1.BatchPredictionJob.InstanceConfig.key_field]
	//  is not specified.
	//
	//  When excluded_fields is populated,
	//  [included_fields][google.cloud.aiplatform.v1.BatchPredictionJob.InstanceConfig.included_fields]
	//  must be empty.
	//
	//  The input must be JSONL with objects at each line, BigQuery
	//  or TfRecord.
	// +kcc:proto:field=google.cloud.aiplatform.v1.BatchPredictionJob.InstanceConfig.excluded_fields
	ExcludedFields []string `json:"excludedFields,omitempty"`
}

// +kcc:proto=google.cloud.aiplatform.v1.BatchPredictionJob.OutputConfig
type BatchPredictionJob_OutputConfig struct {
	// The Cloud Storage location of the directory where the output is
	//  to be written to. In the given directory a new directory is created.
	//  Its name is `prediction-<model-display-name>-<job-create-time>`,
	//  where timestamp is in YYYY-MM-DDThh:mm:ss.sssZ ISO-8601 format.
	//  Inside of it files `predictions_0001.<extension>`,
	//  `predictions_0002.<extension>`, ..., `predictions_N.<extension>`
	//  are created where `<extension>` depends on chosen
	//  [predictions_format][google.cloud.aiplatform.v1.BatchPredictionJob.OutputConfig.predictions_format],
	//  and N may equal 0001 and depends on the total number of successfully
	//  predicted instances. If the Model has both
	//  [instance][google.cloud.aiplatform.v1.PredictSchemata.instance_schema_uri]
	//  and
	//  [prediction][google.cloud.aiplatform.v1.PredictSchemata.parameters_schema_uri]
	//  schemata defined then each such file contains predictions as per the
	//  [predictions_format][google.cloud.aiplatform.v1.BatchPredictionJob.OutputConfig.predictions_format].
	//  If prediction for any instance failed (partially or completely), then
	//  an additional `errors_0001.<extension>`, `errors_0002.<extension>`,...,
	//  `errors_N.<extension>` files are created (N depends on total number
	//  of failed predictions). These files contain the failed instances,
	//  as per their schema, followed by an additional `error` field which as
	//  value has [google.rpc.Status][google.rpc.Status]
	//  containing only `code` and `message` fields.
	// +kcc:proto:field=google.cloud.aiplatform.v1.BatchPredictionJob.OutputConfig.gcs_destination
	GcsDestination *GcsDestination `json:"gcsDestination,omitempty"`

	// The BigQuery project or dataset location where the output is to be
	//  written to. If project is provided, a new dataset is created with name
	//  `prediction_<model-display-name>_<job-create-time>`
	//  where <model-display-name> is made
	//  BigQuery-dataset-name compatible (for example, most special characters
	//  become underscores), and timestamp is in
	//  YYYY_MM_DDThh_mm_ss_sssZ "based on ISO-8601" format. In the dataset
	//  two tables will be created, `predictions`, and `errors`.
	//  If the Model has both
	//  [instance][google.cloud.aiplatform.v1.PredictSchemata.instance_schema_uri]
	//  and
	//  [prediction][google.cloud.aiplatform.v1.PredictSchemata.parameters_schema_uri]
	//  schemata defined then the tables have columns as follows: The
	//  `predictions` table contains instances for which the prediction
	//  succeeded, it has columns as per a concatenation of the Model's
	//  instance and prediction schemata. The `errors` table contains rows for
	//  which the prediction has failed, it has instance columns, as per the
	//  instance schema, followed by a single "errors" column, which as values
	//  has [google.rpc.Status][google.rpc.Status]
	//  represented as a STRUCT, and containing only `code` and `message`.
	// +kcc:proto:field=google.cloud.aiplatform.v1.BatchPredictionJob.OutputConfig.bigquery_destination
	BigqueryDestination *BigQueryDestination `json:"bigqueryDestination,omitempty"`

	// Required. The format in which Vertex AI gives the predictions, must be
	//  one of the [Model's][google.cloud.aiplatform.v1.BatchPredictionJob.model]
	//  [supported_output_storage_formats][google.cloud.aiplatform.v1.Model.supported_output_storage_formats].
	// +kcc:proto:field=google.cloud.aiplatform.v1.BatchPredictionJob.OutputConfig.predictions_format
	PredictionsFormat *string `json:"predictionsFormat,omitempty"`
}

// +kcc:proto=google.cloud.aiplatform.v1.BatchPredictionJob.OutputInfo
type BatchPredictionJob_OutputInfo struct {
}

// +kcc:proto=google.cloud.aiplatform.v1.BigQueryDestination
type BigQueryDestination struct {
	// Required. BigQuery URI to a project or table, up to 2000 characters long.
	//
	//  When only the project is specified, the Dataset and Table is created.
	//  When the full table reference is specified, the Dataset must exist and
	//  table must not exist.
	//
	//  Accepted forms:
	//
	//  *  BigQuery path. For example:
	//  `bq://projectId` or `bq://projectId.bqDatasetId` or
	//  `bq://projectId.bqDatasetId.bqTableId`.
	// +kcc:proto:field=google.cloud.aiplatform.v1.BigQueryDestination.output_uri
	OutputURI *string `json:"outputURI,omitempty"`
}

// +kcc:proto=google.cloud.aiplatform.v1.BigQuerySource
type BigQuerySource struct {
	// Required. BigQuery URI to a table, up to 2000 characters long.
	//  Accepted forms:
	//
	//  *  BigQuery path. For example: `bq://projectId.bqDatasetId.bqTableId`.
	// +kcc:proto:field=google.cloud.aiplatform.v1.BigQuerySource.input_uri
	InputURI *string `json:"inputURI,omitempty"`
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

// +kcc:proto=google.cloud.aiplatform.v1.CompletionStats
type CompletionStats struct {
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

// +kcc:proto=google.cloud.aiplatform.v1.GcsDestination
type GcsDestination struct {
	// Required. Google Cloud Storage URI to output directory. If the uri doesn't
	//  end with
	//  '/', a '/' will be automatically appended. The directory is created if it
	//  doesn't exist.
	// +kcc:proto:field=google.cloud.aiplatform.v1.GcsDestination.output_uri_prefix
	OutputURIPrefix *string `json:"outputURIPrefix,omitempty"`
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

// +kcc:proto=google.cloud.aiplatform.v1.MachineSpec
type MachineSpec struct {
	// Immutable. The type of the machine.
	//
	//  See the [list of machine types supported for
	//  prediction](https://cloud.google.com/vertex-ai/docs/predictions/configure-compute#machine-types)
	//
	//  See the [list of machine types supported for custom
	//  training](https://cloud.google.com/vertex-ai/docs/training/configure-compute#machine-types).
	//
	//  For [DeployedModel][google.cloud.aiplatform.v1.DeployedModel] this field is
	//  optional, and the default value is `n1-standard-2`. For
	//  [BatchPredictionJob][google.cloud.aiplatform.v1.BatchPredictionJob] or as
	//  part of [WorkerPoolSpec][google.cloud.aiplatform.v1.WorkerPoolSpec] this
	//  field is required.
	// +kcc:proto:field=google.cloud.aiplatform.v1.MachineSpec.machine_type
	MachineType *string `json:"machineType,omitempty"`

	// Immutable. The type of accelerator(s) that may be attached to the machine
	//  as per
	//  [accelerator_count][google.cloud.aiplatform.v1.MachineSpec.accelerator_count].
	// +kcc:proto:field=google.cloud.aiplatform.v1.MachineSpec.accelerator_type
	AcceleratorType *string `json:"acceleratorType,omitempty"`

	// The number of accelerators to attach to the machine.
	// +kcc:proto:field=google.cloud.aiplatform.v1.MachineSpec.accelerator_count
	AcceleratorCount *int32 `json:"acceleratorCount,omitempty"`

	// Immutable. The topology of the TPUs. Corresponds to the TPU topologies
	//  available from GKE. (Example: tpu_topology: "2x2x1").
	// +kcc:proto:field=google.cloud.aiplatform.v1.MachineSpec.tpu_topology
	TpuTopology *string `json:"tpuTopology,omitempty"`

	// Optional. Immutable. Configuration controlling how this resource pool
	//  consumes reservation.
	// +kcc:proto:field=google.cloud.aiplatform.v1.MachineSpec.reservation_affinity
	ReservationAffinity *ReservationAffinity `json:"reservationAffinity,omitempty"`
}

// +kcc:proto=google.cloud.aiplatform.v1.ManualBatchTuningParameters
type ManualBatchTuningParameters struct {
	// Immutable. The number of the records (e.g. instances) of the operation
	//  given in each batch to a machine replica. Machine type, and size of a
	//  single record should be considered when setting this parameter, higher
	//  value speeds up the batch operation's execution, but too high value will
	//  result in a whole batch not fitting in a machine's memory, and the whole
	//  operation will fail.
	//  The default value is 64.
	// +kcc:proto:field=google.cloud.aiplatform.v1.ManualBatchTuningParameters.batch_size
	BatchSize *int32 `json:"batchSize,omitempty"`
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

// +kcc:proto=google.cloud.aiplatform.v1.ReservationAffinity
type ReservationAffinity struct {
	// Required. Specifies the reservation affinity type.
	// +kcc:proto:field=google.cloud.aiplatform.v1.ReservationAffinity.reservation_affinity_type
	ReservationAffinityType *string `json:"reservationAffinityType,omitempty"`

	// Optional. Corresponds to the label key of a reservation resource. To target
	//  a SPECIFIC_RESERVATION by name, use
	//  `compute.googleapis.com/reservation-name` as the key and specify the name
	//  of your reservation as its value.
	// +kcc:proto:field=google.cloud.aiplatform.v1.ReservationAffinity.key
	Key *string `json:"key,omitempty"`

	// Optional. Corresponds to the label values of a reservation resource. This
	//  must be the full resource name of the reservation.
	// +kcc:proto:field=google.cloud.aiplatform.v1.ReservationAffinity.values
	Values []string `json:"values,omitempty"`
}

// +kcc:proto=google.cloud.aiplatform.v1.ResourcesConsumed
type ResourcesConsumed struct {
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

// +kcc:proto=google.cloud.aiplatform.v1.UnmanagedContainerModel
type UnmanagedContainerModel struct {
	// The path to the directory containing the Model artifact and any of its
	//  supporting files.
	// +kcc:proto:field=google.cloud.aiplatform.v1.UnmanagedContainerModel.artifact_uri
	ArtifactURI *string `json:"artifactURI,omitempty"`

	// Contains the schemata used in Model's predictions and explanations
	// +kcc:proto:field=google.cloud.aiplatform.v1.UnmanagedContainerModel.predict_schemata
	PredictSchemata *PredictSchemata `json:"predictSchemata,omitempty"`

	// Input only. The specification of the container that is to be used when
	//  deploying this Model.
	// +kcc:proto:field=google.cloud.aiplatform.v1.UnmanagedContainerModel.container_spec
	ContainerSpec *ModelContainerSpec `json:"containerSpec,omitempty"`
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

// +kcc:proto=google.protobuf.Any
type Any struct {
	// A URL/resource name that uniquely identifies the type of the serialized
	//  protocol buffer message. This string must contain at least
	//  one "/" character. The last segment of the URL's path must represent
	//  the fully qualified name of the type (as in
	//  `path/google.protobuf.Duration`). The name should be in a canonical form
	//  (e.g., leading "." is not accepted).
	//
	//  In practice, teams usually precompile into the binary all types that they
	//  expect it to use in the context of Any. However, for URLs which use the
	//  scheme `http`, `https`, or no scheme, one can optionally set up a type
	//  server that maps type URLs to message definitions as follows:
	//
	//  * If no scheme is provided, `https` is assumed.
	//  * An HTTP GET on the URL must yield a [google.protobuf.Type][]
	//    value in binary format, or produce an error.
	//  * Applications are allowed to cache lookup results based on the
	//    URL, or have them precompiled into a binary to avoid any
	//    lookup. Therefore, binary compatibility needs to be preserved
	//    on changes to types. (Use versioned type names to manage
	//    breaking changes.)
	//
	//  Note: this functionality is not currently available in the official
	//  protobuf release, and it is not used for type URLs beginning with
	//  type.googleapis.com.
	//
	//  Schemes other than `http`, `https` (or the empty scheme) might be
	//  used with implementation specific semantics.
	// +kcc:proto:field=google.protobuf.Any.type_url
	TypeURL *string `json:"typeURL,omitempty"`

	// Must be a valid serialized protocol buffer of the above specified type.
	// +kcc:proto:field=google.protobuf.Any.value
	Value []byte `json:"value,omitempty"`
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

// +kcc:proto=google.rpc.Status
type Status struct {
	// The status code, which should be an enum value of
	//  [google.rpc.Code][google.rpc.Code].
	// +kcc:proto:field=google.rpc.Status.code
	Code *int32 `json:"code,omitempty"`

	// A developer-facing error message, which should be in English. Any
	//  user-facing error message should be localized and sent in the
	//  [google.rpc.Status.details][google.rpc.Status.details] field, or localized
	//  by the client.
	// +kcc:proto:field=google.rpc.Status.message
	Message *string `json:"message,omitempty"`

	// A list of messages that carry the error details.  There is a common set of
	//  message types for APIs to use.
	// +kcc:proto:field=google.rpc.Status.details
	Details []Any `json:"details,omitempty"`
}

// +kcc:proto=google.cloud.aiplatform.v1.BatchPredictionJob
type BatchPredictionJobObservedState struct {
	// Output only. Resource name of the BatchPredictionJob.
	// +kcc:proto:field=google.cloud.aiplatform.v1.BatchPredictionJob.name
	Name *string `json:"name,omitempty"`

	// Output only. The version ID of the Model that produces the predictions via
	//  this job.
	// +kcc:proto:field=google.cloud.aiplatform.v1.BatchPredictionJob.model_version_id
	ModelVersionID *string `json:"modelVersionID,omitempty"`

	// Output only. Information further describing the output of this job.
	// +kcc:proto:field=google.cloud.aiplatform.v1.BatchPredictionJob.output_info
	OutputInfo *BatchPredictionJob_OutputInfo `json:"outputInfo,omitempty"`

	// Output only. The detailed state of the job.
	// +kcc:proto:field=google.cloud.aiplatform.v1.BatchPredictionJob.state
	State *string `json:"state,omitempty"`

	// Output only. Only populated when the job's state is JOB_STATE_FAILED or
	//  JOB_STATE_CANCELLED.
	// +kcc:proto:field=google.cloud.aiplatform.v1.BatchPredictionJob.error
	Error *Status `json:"error,omitempty"`

	// Output only. Partial failures encountered.
	//  For example, single files that can't be read.
	//  This field never exceeds 20 entries.
	//  Status details fields contain standard Google Cloud error details.
	// +kcc:proto:field=google.cloud.aiplatform.v1.BatchPredictionJob.partial_failures
	PartialFailures []Status `json:"partialFailures,omitempty"`

	// Output only. Information about resources that had been consumed by this
	//  job. Provided in real time at best effort basis, as well as a final value
	//  once the job completes.
	//
	//  Note: This field currently may be not populated for batch predictions that
	//  use AutoML Models.
	// +kcc:proto:field=google.cloud.aiplatform.v1.BatchPredictionJob.resources_consumed
	ResourcesConsumed *ResourcesConsumed `json:"resourcesConsumed,omitempty"`

	// Output only. Statistics on completed and failed prediction instances.
	// +kcc:proto:field=google.cloud.aiplatform.v1.BatchPredictionJob.completion_stats
	CompletionStats *CompletionStats `json:"completionStats,omitempty"`

	// Output only. Time when the BatchPredictionJob was created.
	// +kcc:proto:field=google.cloud.aiplatform.v1.BatchPredictionJob.create_time
	CreateTime *string `json:"createTime,omitempty"`

	// Output only. Time when the BatchPredictionJob for the first time entered
	//  the `JOB_STATE_RUNNING` state.
	// +kcc:proto:field=google.cloud.aiplatform.v1.BatchPredictionJob.start_time
	StartTime *string `json:"startTime,omitempty"`

	// Output only. Time when the BatchPredictionJob entered any of the following
	//  states: `JOB_STATE_SUCCEEDED`, `JOB_STATE_FAILED`, `JOB_STATE_CANCELLED`.
	// +kcc:proto:field=google.cloud.aiplatform.v1.BatchPredictionJob.end_time
	EndTime *string `json:"endTime,omitempty"`

	// Output only. Time when the BatchPredictionJob was most recently updated.
	// +kcc:proto:field=google.cloud.aiplatform.v1.BatchPredictionJob.update_time
	UpdateTime *string `json:"updateTime,omitempty"`

	// Output only. Reserved for future use.
	// +kcc:proto:field=google.cloud.aiplatform.v1.BatchPredictionJob.satisfies_pzs
	SatisfiesPzs *bool `json:"satisfiesPzs,omitempty"`

	// Output only. Reserved for future use.
	// +kcc:proto:field=google.cloud.aiplatform.v1.BatchPredictionJob.satisfies_pzi
	SatisfiesPzi *bool `json:"satisfiesPzi,omitempty"`
}

// +kcc:proto=google.cloud.aiplatform.v1.BatchPredictionJob.OutputInfo
type BatchPredictionJob_OutputInfoObservedState struct {
	// Output only. The full path of the Cloud Storage directory created, into
	//  which the prediction output is written.
	// +kcc:proto:field=google.cloud.aiplatform.v1.BatchPredictionJob.OutputInfo.gcs_output_directory
	GcsOutputDirectory *string `json:"gcsOutputDirectory,omitempty"`

	// Output only. The path of the BigQuery dataset created, in
	//  `bq://projectId.bqDatasetId`
	//  format, into which the prediction output is written.
	// +kcc:proto:field=google.cloud.aiplatform.v1.BatchPredictionJob.OutputInfo.bigquery_output_dataset
	BigqueryOutputDataset *string `json:"bigqueryOutputDataset,omitempty"`

	// Output only. The name of the BigQuery table created, in
	//  `predictions_<timestamp>`
	//  format, into which the prediction output is written.
	//  Can be used by UI to generate the BigQuery output path, for example.
	// +kcc:proto:field=google.cloud.aiplatform.v1.BatchPredictionJob.OutputInfo.bigquery_output_table
	BigqueryOutputTable *string `json:"bigqueryOutputTable,omitempty"`
}

// +kcc:proto=google.cloud.aiplatform.v1.CompletionStats
type CompletionStatsObservedState struct {
	// Output only. The number of entities that had been processed successfully.
	// +kcc:proto:field=google.cloud.aiplatform.v1.CompletionStats.successful_count
	SuccessfulCount *int64 `json:"successfulCount,omitempty"`

	// Output only. The number of entities for which any error was encountered.
	// +kcc:proto:field=google.cloud.aiplatform.v1.CompletionStats.failed_count
	FailedCount *int64 `json:"failedCount,omitempty"`

	// Output only. In cases when enough errors are encountered a job, pipeline,
	//  or operation may be failed as a whole. Below is the number of entities for
	//  which the processing had not been finished (either in successful or failed
	//  state). Set to -1 if the number is unknown (for example, the operation
	//  failed before the total entity number could be collected).
	// +kcc:proto:field=google.cloud.aiplatform.v1.CompletionStats.incomplete_count
	IncompleteCount *int64 `json:"incompleteCount,omitempty"`

	// Output only. The number of the successful forecast points that are
	//  generated by the forecasting model. This is ONLY used by the forecasting
	//  batch prediction.
	// +kcc:proto:field=google.cloud.aiplatform.v1.CompletionStats.successful_forecast_point_count
	SuccessfulForecastPointCount *int64 `json:"successfulForecastPointCount,omitempty"`
}

// +kcc:proto=google.cloud.aiplatform.v1.ResourcesConsumed
type ResourcesConsumedObservedState struct {
	// Output only. The number of replica hours used. Note that many replicas may
	//  run in parallel, and additionally any given work may be queued for some
	//  time. Therefore this value is not strictly related to wall time.
	// +kcc:proto:field=google.cloud.aiplatform.v1.ResourcesConsumed.replica_hours
	ReplicaHours *float64 `json:"replicaHours,omitempty"`
}
