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

// +kcc:proto=google.cloud.aiplatform.v1.ModelDeploymentMonitoringBigQueryTable
type ModelDeploymentMonitoringBigQueryTable struct {
	// The source of log.
	// +kcc:proto:field=google.cloud.aiplatform.v1.ModelDeploymentMonitoringBigQueryTable.log_source
	LogSource *string `json:"logSource,omitempty"`

	// The type of log.
	// +kcc:proto:field=google.cloud.aiplatform.v1.ModelDeploymentMonitoringBigQueryTable.log_type
	LogType *string `json:"logType,omitempty"`

	// The created BigQuery table to store logs. Customer could do their own query
	//  & analysis. Format:
	//  `bq://<project_id>.model_deployment_monitoring_<endpoint_id>.<tolower(log_source)>_<tolower(log_type)>`
	// +kcc:proto:field=google.cloud.aiplatform.v1.ModelDeploymentMonitoringBigQueryTable.bigquery_table_path
	BigqueryTablePath *string `json:"bigqueryTablePath,omitempty"`
}

// +kcc:proto=google.cloud.aiplatform.v1.ModelDeploymentMonitoringJob
type ModelDeploymentMonitoringJob struct {

	// Required. The user-defined name of the ModelDeploymentMonitoringJob.
	//  The name can be up to 128 characters long and can consist of any UTF-8
	//  characters.
	//  Display name of a ModelDeploymentMonitoringJob.
	// +kcc:proto:field=google.cloud.aiplatform.v1.ModelDeploymentMonitoringJob.display_name
	DisplayName *string `json:"displayName,omitempty"`

	// Required. Endpoint resource name.
	//  Format: `projects/{project}/locations/{location}/endpoints/{endpoint}`
	// +kcc:proto:field=google.cloud.aiplatform.v1.ModelDeploymentMonitoringJob.endpoint
	Endpoint *string `json:"endpoint,omitempty"`

	// Required. The config for monitoring objectives. This is a per DeployedModel
	//  config. Each DeployedModel needs to be configured separately.
	// +kcc:proto:field=google.cloud.aiplatform.v1.ModelDeploymentMonitoringJob.model_deployment_monitoring_objective_configs
	ModelDeploymentMonitoringObjectiveConfigs []ModelDeploymentMonitoringObjectiveConfig `json:"modelDeploymentMonitoringObjectiveConfigs,omitempty"`

	// Required. Schedule config for running the monitoring job.
	// +kcc:proto:field=google.cloud.aiplatform.v1.ModelDeploymentMonitoringJob.model_deployment_monitoring_schedule_config
	ModelDeploymentMonitoringScheduleConfig *ModelDeploymentMonitoringScheduleConfig `json:"modelDeploymentMonitoringScheduleConfig,omitempty"`

	// Required. Sample Strategy for logging.
	// +kcc:proto:field=google.cloud.aiplatform.v1.ModelDeploymentMonitoringJob.logging_sampling_strategy
	LoggingSamplingStrategy *SamplingStrategy `json:"loggingSamplingStrategy,omitempty"`

	// Alert config for model monitoring.
	// +kcc:proto:field=google.cloud.aiplatform.v1.ModelDeploymentMonitoringJob.model_monitoring_alert_config
	ModelMonitoringAlertConfig *ModelMonitoringAlertConfig `json:"modelMonitoringAlertConfig,omitempty"`

	// YAML schema file uri describing the format of a single instance,
	//  which are given to format this Endpoint's prediction (and explanation).
	//  If not set, we will generate predict schema from collected predict
	//  requests.
	// +kcc:proto:field=google.cloud.aiplatform.v1.ModelDeploymentMonitoringJob.predict_instance_schema_uri
	PredictInstanceSchemaURI *string `json:"predictInstanceSchemaURI,omitempty"`

	// Sample Predict instance, same format as
	//  [PredictRequest.instances][google.cloud.aiplatform.v1.PredictRequest.instances],
	//  this can be set as a replacement of
	//  [ModelDeploymentMonitoringJob.predict_instance_schema_uri][google.cloud.aiplatform.v1.ModelDeploymentMonitoringJob.predict_instance_schema_uri].
	//  If not set, we will generate predict schema from collected predict
	//  requests.
	// +kcc:proto:field=google.cloud.aiplatform.v1.ModelDeploymentMonitoringJob.sample_predict_instance
	SamplePredictInstance *Value `json:"samplePredictInstance,omitempty"`

	// YAML schema file uri describing the format of a single instance that you
	//  want Tensorflow Data Validation (TFDV) to analyze.
	//
	//  If this field is empty, all the feature data types are inferred from
	//  [predict_instance_schema_uri][google.cloud.aiplatform.v1.ModelDeploymentMonitoringJob.predict_instance_schema_uri],
	//  meaning that TFDV will use the data in the exact format(data type) as
	//  prediction request/response.
	//  If there are any data type differences between predict instance and TFDV
	//  instance, this field can be used to override the schema.
	//  For models trained with Vertex AI, this field must be set as all the
	//  fields in predict instance formatted as string.
	// +kcc:proto:field=google.cloud.aiplatform.v1.ModelDeploymentMonitoringJob.analysis_instance_schema_uri
	AnalysisInstanceSchemaURI *string `json:"analysisInstanceSchemaURI,omitempty"`

	// The TTL of BigQuery tables in user projects which stores logs.
	//  A day is the basic unit of the TTL and we take the ceil of TTL/86400(a
	//  day). e.g. { second: 3600} indicates ttl = 1 day.
	// +kcc:proto:field=google.cloud.aiplatform.v1.ModelDeploymentMonitoringJob.log_ttl
	LogTtl *string `json:"logTtl,omitempty"`

	// The labels with user-defined metadata to organize your
	//  ModelDeploymentMonitoringJob.
	//
	//  Label keys and values can be no longer than 64 characters
	//  (Unicode codepoints), can only contain lowercase letters, numeric
	//  characters, underscores and dashes. International characters are allowed.
	//
	//  See https://goo.gl/xmQnxf for more information and examples of labels.
	// +kcc:proto:field=google.cloud.aiplatform.v1.ModelDeploymentMonitoringJob.labels
	Labels map[string]string `json:"labels,omitempty"`

	// Stats anomalies base folder path.
	// +kcc:proto:field=google.cloud.aiplatform.v1.ModelDeploymentMonitoringJob.stats_anomalies_base_directory
	StatsAnomaliesBaseDirectory *GcsDestination `json:"statsAnomaliesBaseDirectory,omitempty"`

	// Customer-managed encryption key spec for a ModelDeploymentMonitoringJob. If
	//  set, this ModelDeploymentMonitoringJob and all sub-resources of this
	//  ModelDeploymentMonitoringJob will be secured by this key.
	// +kcc:proto:field=google.cloud.aiplatform.v1.ModelDeploymentMonitoringJob.encryption_spec
	EncryptionSpec *EncryptionSpec `json:"encryptionSpec,omitempty"`

	// If true, the scheduled monitoring pipeline logs are sent to
	//  Google Cloud Logging, including pipeline status and anomalies detected.
	//  Please note the logs incur cost, which are subject to [Cloud Logging
	//  pricing](https://cloud.google.com/logging#pricing).
	// +kcc:proto:field=google.cloud.aiplatform.v1.ModelDeploymentMonitoringJob.enable_monitoring_pipeline_logs
	EnableMonitoringPipelineLogs *bool `json:"enableMonitoringPipelineLogs,omitempty"`
}

// +kcc:proto=google.cloud.aiplatform.v1.ModelDeploymentMonitoringJob.LatestMonitoringPipelineMetadata
type ModelDeploymentMonitoringJob_LatestMonitoringPipelineMetadata struct {
	// The time that most recent monitoring pipelines that is related to this
	//  run.
	// +kcc:proto:field=google.cloud.aiplatform.v1.ModelDeploymentMonitoringJob.LatestMonitoringPipelineMetadata.run_time
	RunTime *string `json:"runTime,omitempty"`

	// The status of the most recent monitoring pipeline.
	// +kcc:proto:field=google.cloud.aiplatform.v1.ModelDeploymentMonitoringJob.LatestMonitoringPipelineMetadata.status
	Status *Status `json:"status,omitempty"`
}

// +kcc:proto=google.cloud.aiplatform.v1.ModelDeploymentMonitoringObjectiveConfig
type ModelDeploymentMonitoringObjectiveConfig struct {
	// The DeployedModel ID of the objective config.
	// +kcc:proto:field=google.cloud.aiplatform.v1.ModelDeploymentMonitoringObjectiveConfig.deployed_model_id
	DeployedModelID *string `json:"deployedModelID,omitempty"`

	// The objective config of for the modelmonitoring job of this deployed model.
	// +kcc:proto:field=google.cloud.aiplatform.v1.ModelDeploymentMonitoringObjectiveConfig.objective_config
	ObjectiveConfig *ModelMonitoringObjectiveConfig `json:"objectiveConfig,omitempty"`
}

// +kcc:proto=google.cloud.aiplatform.v1.ModelDeploymentMonitoringScheduleConfig
type ModelDeploymentMonitoringScheduleConfig struct {
	// Required. The model monitoring job scheduling interval. It will be rounded
	//  up to next full hour. This defines how often the monitoring jobs are
	//  triggered.
	// +kcc:proto:field=google.cloud.aiplatform.v1.ModelDeploymentMonitoringScheduleConfig.monitor_interval
	MonitorInterval *string `json:"monitorInterval,omitempty"`

	// The time window of the prediction data being included in each prediction
	//  dataset. This window specifies how long the data should be collected from
	//  historical model results for each run. If not set,
	//  [ModelDeploymentMonitoringScheduleConfig.monitor_interval][google.cloud.aiplatform.v1.ModelDeploymentMonitoringScheduleConfig.monitor_interval]
	//  will be used. e.g. If currently the cutoff time is 2022-01-08 14:30:00 and
	//  the monitor_window is set to be 3600, then data from 2022-01-08 13:30:00 to
	//  2022-01-08 14:30:00 will be retrieved and aggregated to calculate the
	//  monitoring statistics.
	// +kcc:proto:field=google.cloud.aiplatform.v1.ModelDeploymentMonitoringScheduleConfig.monitor_window
	MonitorWindow *string `json:"monitorWindow,omitempty"`
}

// +kcc:proto=google.cloud.aiplatform.v1.ModelMonitoringAlertConfig
type ModelMonitoringAlertConfig struct {
	// Email alert config.
	// +kcc:proto:field=google.cloud.aiplatform.v1.ModelMonitoringAlertConfig.email_alert_config
	EmailAlertConfig *ModelMonitoringAlertConfig_EmailAlertConfig `json:"emailAlertConfig,omitempty"`

	// Dump the anomalies to Cloud Logging. The anomalies will be put to json
	//  payload encoded from proto
	//  [ModelMonitoringStatsAnomalies][google.cloud.aiplatform.v1.ModelMonitoringStatsAnomalies].
	//  This can be further synced to Pub/Sub or any other services supported by
	//  Cloud Logging.
	// +kcc:proto:field=google.cloud.aiplatform.v1.ModelMonitoringAlertConfig.enable_logging
	EnableLogging *bool `json:"enableLogging,omitempty"`

	// Resource names of the NotificationChannels to send alert.
	//  Must be of the format
	//  `projects/<project_id_or_number>/notificationChannels/<channel_id>`
	// +kcc:proto:field=google.cloud.aiplatform.v1.ModelMonitoringAlertConfig.notification_channels
	NotificationChannels []string `json:"notificationChannels,omitempty"`
}

// +kcc:proto=google.cloud.aiplatform.v1.ModelMonitoringAlertConfig.EmailAlertConfig
type ModelMonitoringAlertConfig_EmailAlertConfig struct {
	// The email addresses to send the alert.
	// +kcc:proto:field=google.cloud.aiplatform.v1.ModelMonitoringAlertConfig.EmailAlertConfig.user_emails
	UserEmails []string `json:"userEmails,omitempty"`
}

// +kcc:proto=google.cloud.aiplatform.v1.ModelMonitoringObjectiveConfig
type ModelMonitoringObjectiveConfig struct {
	// Training dataset for models. This field has to be set only if
	//  TrainingPredictionSkewDetectionConfig is specified.
	// +kcc:proto:field=google.cloud.aiplatform.v1.ModelMonitoringObjectiveConfig.training_dataset
	TrainingDataset *ModelMonitoringObjectiveConfig_TrainingDataset `json:"trainingDataset,omitempty"`

	// The config for skew between training data and prediction data.
	// +kcc:proto:field=google.cloud.aiplatform.v1.ModelMonitoringObjectiveConfig.training_prediction_skew_detection_config
	TrainingPredictionSkewDetectionConfig *ModelMonitoringObjectiveConfig_TrainingPredictionSkewDetectionConfig `json:"trainingPredictionSkewDetectionConfig,omitempty"`

	// The config for drift of prediction data.
	// +kcc:proto:field=google.cloud.aiplatform.v1.ModelMonitoringObjectiveConfig.prediction_drift_detection_config
	PredictionDriftDetectionConfig *ModelMonitoringObjectiveConfig_PredictionDriftDetectionConfig `json:"predictionDriftDetectionConfig,omitempty"`

	// The config for integrating with Vertex Explainable AI.
	// +kcc:proto:field=google.cloud.aiplatform.v1.ModelMonitoringObjectiveConfig.explanation_config
	ExplanationConfig *ModelMonitoringObjectiveConfig_ExplanationConfig `json:"explanationConfig,omitempty"`
}

// +kcc:proto=google.cloud.aiplatform.v1.ModelMonitoringObjectiveConfig.ExplanationConfig
type ModelMonitoringObjectiveConfig_ExplanationConfig struct {
	// If want to analyze the Vertex Explainable AI feature attribute scores or
	//  not. If set to true, Vertex AI will log the feature attributions from
	//  explain response and do the skew/drift detection for them.
	// +kcc:proto:field=google.cloud.aiplatform.v1.ModelMonitoringObjectiveConfig.ExplanationConfig.enable_feature_attributes
	EnableFeatureAttributes *bool `json:"enableFeatureAttributes,omitempty"`

	// Predictions generated by the BatchPredictionJob using baseline dataset.
	// +kcc:proto:field=google.cloud.aiplatform.v1.ModelMonitoringObjectiveConfig.ExplanationConfig.explanation_baseline
	ExplanationBaseline *ModelMonitoringObjectiveConfig_ExplanationConfig_ExplanationBaseline `json:"explanationBaseline,omitempty"`
}

// +kcc:proto=google.cloud.aiplatform.v1.ModelMonitoringObjectiveConfig.ExplanationConfig.ExplanationBaseline
type ModelMonitoringObjectiveConfig_ExplanationConfig_ExplanationBaseline struct {
	// Cloud Storage location for BatchExplain output.
	// +kcc:proto:field=google.cloud.aiplatform.v1.ModelMonitoringObjectiveConfig.ExplanationConfig.ExplanationBaseline.gcs
	Gcs *GcsDestination `json:"gcs,omitempty"`

	// BigQuery location for BatchExplain output.
	// +kcc:proto:field=google.cloud.aiplatform.v1.ModelMonitoringObjectiveConfig.ExplanationConfig.ExplanationBaseline.bigquery
	Bigquery *BigQueryDestination `json:"bigquery,omitempty"`

	// The storage format of the predictions generated BatchPrediction job.
	// +kcc:proto:field=google.cloud.aiplatform.v1.ModelMonitoringObjectiveConfig.ExplanationConfig.ExplanationBaseline.prediction_format
	PredictionFormat *string `json:"predictionFormat,omitempty"`
}

// +kcc:proto=google.cloud.aiplatform.v1.ModelMonitoringObjectiveConfig.PredictionDriftDetectionConfig
type ModelMonitoringObjectiveConfig_PredictionDriftDetectionConfig struct {

	// TODO: unsupported map type with key string and value message


	// TODO: unsupported map type with key string and value message


	// Drift anomaly detection threshold used by all features.
	//  When the per-feature thresholds are not set, this field can be used to
	//  specify a threshold for all features.
	// +kcc:proto:field=google.cloud.aiplatform.v1.ModelMonitoringObjectiveConfig.PredictionDriftDetectionConfig.default_drift_threshold
	DefaultDriftThreshold *ThresholdConfig `json:"defaultDriftThreshold,omitempty"`
}

// +kcc:proto=google.cloud.aiplatform.v1.ModelMonitoringObjectiveConfig.TrainingDataset
type ModelMonitoringObjectiveConfig_TrainingDataset struct {
	// The resource name of the Dataset used to train this Model.
	// +kcc:proto:field=google.cloud.aiplatform.v1.ModelMonitoringObjectiveConfig.TrainingDataset.dataset
	Dataset *string `json:"dataset,omitempty"`

	// The Google Cloud Storage uri of the unmanaged Dataset used to train
	//  this Model.
	// +kcc:proto:field=google.cloud.aiplatform.v1.ModelMonitoringObjectiveConfig.TrainingDataset.gcs_source
	GcsSource *GcsSource `json:"gcsSource,omitempty"`

	// The BigQuery table of the unmanaged Dataset used to train this
	//  Model.
	// +kcc:proto:field=google.cloud.aiplatform.v1.ModelMonitoringObjectiveConfig.TrainingDataset.bigquery_source
	BigquerySource *BigQuerySource `json:"bigquerySource,omitempty"`

	// Data format of the dataset, only applicable if the input is from
	//  Google Cloud Storage.
	//  The possible formats are:
	//
	//  "tf-record"
	//  The source file is a TFRecord file.
	//
	//  "csv"
	//  The source file is a CSV file.
	//  "jsonl"
	//  The source file is a JSONL file.
	// +kcc:proto:field=google.cloud.aiplatform.v1.ModelMonitoringObjectiveConfig.TrainingDataset.data_format
	DataFormat *string `json:"dataFormat,omitempty"`

	// The target field name the model is to predict.
	//  This field will be excluded when doing Predict and (or) Explain for the
	//  training data.
	// +kcc:proto:field=google.cloud.aiplatform.v1.ModelMonitoringObjectiveConfig.TrainingDataset.target_field
	TargetField *string `json:"targetField,omitempty"`

	// Strategy to sample data from Training Dataset.
	//  If not set, we process the whole dataset.
	// +kcc:proto:field=google.cloud.aiplatform.v1.ModelMonitoringObjectiveConfig.TrainingDataset.logging_sampling_strategy
	LoggingSamplingStrategy *SamplingStrategy `json:"loggingSamplingStrategy,omitempty"`
}

// +kcc:proto=google.cloud.aiplatform.v1.ModelMonitoringObjectiveConfig.TrainingPredictionSkewDetectionConfig
type ModelMonitoringObjectiveConfig_TrainingPredictionSkewDetectionConfig struct {

	// TODO: unsupported map type with key string and value message


	// TODO: unsupported map type with key string and value message


	// Skew anomaly detection threshold used by all features.
	//  When the per-feature thresholds are not set, this field can be used to
	//  specify a threshold for all features.
	// +kcc:proto:field=google.cloud.aiplatform.v1.ModelMonitoringObjectiveConfig.TrainingPredictionSkewDetectionConfig.default_skew_threshold
	DefaultSkewThreshold *ThresholdConfig `json:"defaultSkewThreshold,omitempty"`
}

// +kcc:proto=google.cloud.aiplatform.v1.SamplingStrategy
type SamplingStrategy struct {
	// Random sample config. Will support more sampling strategies later.
	// +kcc:proto:field=google.cloud.aiplatform.v1.SamplingStrategy.random_sample_config
	RandomSampleConfig *SamplingStrategy_RandomSampleConfig `json:"randomSampleConfig,omitempty"`
}

// +kcc:proto=google.cloud.aiplatform.v1.SamplingStrategy.RandomSampleConfig
type SamplingStrategy_RandomSampleConfig struct {
	// Sample rate (0, 1]
	// +kcc:proto:field=google.cloud.aiplatform.v1.SamplingStrategy.RandomSampleConfig.sample_rate
	SampleRate *float64 `json:"sampleRate,omitempty"`
}

// +kcc:proto=google.cloud.aiplatform.v1.ThresholdConfig
type ThresholdConfig struct {
	// Specify a threshold value that can trigger the alert.
	//  If this threshold config is for feature distribution distance:
	//    1. For categorical feature, the distribution distance is calculated by
	//       L-inifinity norm.
	//    2. For numerical feature, the distribution distance is calculated by
	//       Jensen–Shannon divergence.
	//  Each feature must have a non-zero threshold if they need to be monitored.
	//  Otherwise no alert will be triggered for that feature.
	// +kcc:proto:field=google.cloud.aiplatform.v1.ThresholdConfig.value
	Value *float64 `json:"value,omitempty"`
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

// +kcc:proto=google.cloud.aiplatform.v1.ModelDeploymentMonitoringBigQueryTable
type ModelDeploymentMonitoringBigQueryTableObservedState struct {
	// Output only. The schema version of the request/response logging BigQuery
	//  table. Default to v1 if unset.
	// +kcc:proto:field=google.cloud.aiplatform.v1.ModelDeploymentMonitoringBigQueryTable.request_response_logging_schema_version
	RequestResponseLoggingSchemaVersion *string `json:"requestResponseLoggingSchemaVersion,omitempty"`
}

// +kcc:proto=google.cloud.aiplatform.v1.ModelDeploymentMonitoringJob
type ModelDeploymentMonitoringJobObservedState struct {
	// Output only. Resource name of a ModelDeploymentMonitoringJob.
	// +kcc:proto:field=google.cloud.aiplatform.v1.ModelDeploymentMonitoringJob.name
	Name *string `json:"name,omitempty"`

	// Output only. The detailed state of the monitoring job.
	//  When the job is still creating, the state will be 'PENDING'.
	//  Once the job is successfully created, the state will be 'RUNNING'.
	//  Pause the job, the state will be 'PAUSED'.
	//  Resume the job, the state will return to 'RUNNING'.
	// +kcc:proto:field=google.cloud.aiplatform.v1.ModelDeploymentMonitoringJob.state
	State *string `json:"state,omitempty"`

	// Output only. Schedule state when the monitoring job is in Running state.
	// +kcc:proto:field=google.cloud.aiplatform.v1.ModelDeploymentMonitoringJob.schedule_state
	ScheduleState *string `json:"scheduleState,omitempty"`

	// Output only. Latest triggered monitoring pipeline metadata.
	// +kcc:proto:field=google.cloud.aiplatform.v1.ModelDeploymentMonitoringJob.latest_monitoring_pipeline_metadata
	LatestMonitoringPipelineMetadata *ModelDeploymentMonitoringJob_LatestMonitoringPipelineMetadata `json:"latestMonitoringPipelineMetadata,omitempty"`

	// Output only. The created bigquery tables for the job under customer
	//  project. Customer could do their own query & analysis. There could be 4 log
	//  tables in maximum:
	//  1. Training data logging predict request/response
	//  2. Serving data logging predict request/response
	// +kcc:proto:field=google.cloud.aiplatform.v1.ModelDeploymentMonitoringJob.bigquery_tables
	BigqueryTables []ModelDeploymentMonitoringBigQueryTable `json:"bigqueryTables,omitempty"`

	// Output only. Timestamp when this ModelDeploymentMonitoringJob was created.
	// +kcc:proto:field=google.cloud.aiplatform.v1.ModelDeploymentMonitoringJob.create_time
	CreateTime *string `json:"createTime,omitempty"`

	// Output only. Timestamp when this ModelDeploymentMonitoringJob was updated
	//  most recently.
	// +kcc:proto:field=google.cloud.aiplatform.v1.ModelDeploymentMonitoringJob.update_time
	UpdateTime *string `json:"updateTime,omitempty"`

	// Output only. Timestamp when this monitoring pipeline will be scheduled to
	//  run for the next round.
	// +kcc:proto:field=google.cloud.aiplatform.v1.ModelDeploymentMonitoringJob.next_schedule_time
	NextScheduleTime *string `json:"nextScheduleTime,omitempty"`

	// Output only. Only populated when the job's state is `JOB_STATE_FAILED` or
	//  `JOB_STATE_CANCELLED`.
	// +kcc:proto:field=google.cloud.aiplatform.v1.ModelDeploymentMonitoringJob.error
	Error *Status `json:"error,omitempty"`

	// Output only. Reserved for future use.
	// +kcc:proto:field=google.cloud.aiplatform.v1.ModelDeploymentMonitoringJob.satisfies_pzs
	SatisfiesPzs *bool `json:"satisfiesPzs,omitempty"`

	// Output only. Reserved for future use.
	// +kcc:proto:field=google.cloud.aiplatform.v1.ModelDeploymentMonitoringJob.satisfies_pzi
	SatisfiesPzi *bool `json:"satisfiesPzi,omitempty"`
}
