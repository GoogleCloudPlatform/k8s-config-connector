// Copyright 2025 Google LLC
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
	"github.com/GoogleCloudPlatform/k8s-config-connector/apis/refs/v1beta1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/apis/k8s/v1alpha1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

var AIPlatformModelGVK = GroupVersion.WithKind("AIPlatformModel")

// AIPlatformModelSpec defines the desired state of AIPlatformModel
// +kcc:proto=google.cloud.aiplatform.v1.Model
type AIPlatformModelSpec struct {
	// User provided version aliases so that a model version can be referenced via
	//  alias (i.e.
	//  `projects/{project}/locations/{location}/models/{model_id}@{version_alias}`
	//  instead of auto-generated version id (i.e.
	//  `projects/{project}/locations/{location}/models/{model_id}@{version_id})`.
	//  The format is [a-z][a-zA-Z0-9-]{0,126}[a-z0-9] to distinguish from
	//  version_id. A default version alias will be created for the first version
	//  of the model, and there must be exactly one default version alias for a
	//  model.
	// +kcc:proto:field=google.cloud.aiplatform.v1.Model.version_aliases
	VersionAliases []string `json:"versionAliases,omitempty"`

	// Required. The display name of the Model.
	//  The name can be up to 128 characters long and can consist of any UTF-8
	//  characters.
	// +kcc:proto:field=google.cloud.aiplatform.v1.Model.display_name
	DisplayName *string `json:"displayName,omitempty"`

	// The description of the Model.
	// +kcc:proto:field=google.cloud.aiplatform.v1.Model.description
	Description *string `json:"description,omitempty"`

	// The description of this version.
	// +kcc:proto:field=google.cloud.aiplatform.v1.Model.version_description
	VersionDescription *string `json:"versionDescription,omitempty"`

	// The schemata that describe formats of the Model's predictions and
	//  explanations as given and returned via
	//  [PredictionService.Predict][google.cloud.aiplatform.v1.PredictionService.Predict]
	//  and
	//  [PredictionService.Explain][google.cloud.aiplatform.v1.PredictionService.Explain].
	// +kcc:proto:field=google.cloud.aiplatform.v1.Model.predict_schemata
	PredictSchemata *PredictSchemata `json:"predictSchemata,omitempty"`

	// Immutable. Points to a YAML file stored on Google Cloud Storage describing
	//  additional information about the Model, that is specific to it. Unset if
	//  the Model does not have any additional information. The schema is defined
	//  as an OpenAPI 3.0.2 [Schema
	//  Object](https://github.com/OAI/OpenAPI-Specification/blob/main/versions/3.0.2.md#schemaObject).
	//  AutoML Models always have this field populated by Vertex AI, if no
	//  additional metadata is needed, this field is set to an empty string.
	//  Note: The URI given on output will be immutable and probably different,
	//  including the URI scheme, than the one given on input. The output URI will
	//  point to a location where the user only has a read access.
	// +kcc:proto:field=google.cloud.aiplatform.v1.Model.metadata_schema_uri
	MetadataSchemaURI *string `json:"metadataSchemaURI,omitempty"`

	// Immutable. An additional information about the Model; the schema of the
	//  metadata can be found in
	//  [metadata_schema][google.cloud.aiplatform.v1.Model.metadata_schema_uri].
	//  Unset if the Model does not have any additional information.
	// +kcc:proto:field=google.cloud.aiplatform.v1.Model.metadata
	Metadata *Value `json:"metadata,omitempty"`

	// Optional. This field is populated if the model is produced by a pipeline
	//  job.
	// +kcc:proto:field=google.cloud.aiplatform.v1.Model.pipeline_job
	PipelineJob *string `json:"pipelineJob,omitempty"`

	// Input only. The specification of the container that is to be used when
	//  deploying this Model. The specification is ingested upon
	//  [ModelService.UploadModel][google.cloud.aiplatform.v1.ModelService.UploadModel],
	//  and all binaries it contains are copied and stored internally by Vertex AI.
	//  Not required for AutoML Models.
	// +kcc:proto:field=google.cloud.aiplatform.v1.Model.container_spec
	ContainerSpec *ModelContainerSpec `json:"containerSpec,omitempty"`

	// Immutable. The path to the directory containing the Model artifact and any
	//  of its supporting files. Not required for AutoML Models.
	// +kcc:proto:field=google.cloud.aiplatform.v1.Model.artifact_uri
	ArtifactURI *string `json:"artifactURI,omitempty"`

	// The default explanation specification for this Model.
	//
	//  The Model can be used for
	//  [requesting
	//  explanation][google.cloud.aiplatform.v1.PredictionService.Explain] after
	//  being [deployed][google.cloud.aiplatform.v1.EndpointService.DeployModel] if
	//  it is populated. The Model can be used for [batch
	//  explanation][google.cloud.aiplatform.v1.BatchPredictionJob.generate_explanation]
	//  if it is populated.
	//
	//  All fields of the explanation_spec can be overridden by
	//  [explanation_spec][google.cloud.aiplatform.v1.DeployedModel.explanation_spec]
	//  of
	//  [DeployModelRequest.deployed_model][google.cloud.aiplatform.v1.DeployModelRequest.deployed_model],
	//  or
	//  [explanation_spec][google.cloud.aiplatform.v1.BatchPredictionJob.explanation_spec]
	//  of [BatchPredictionJob][google.cloud.aiplatform.v1.BatchPredictionJob].
	//
	//  If the default explanation specification is not set for this Model, this
	//  Model can still be used for
	//  [requesting
	//  explanation][google.cloud.aiplatform.v1.PredictionService.Explain] by
	//  setting
	//  [explanation_spec][google.cloud.aiplatform.v1.DeployedModel.explanation_spec]
	//  of
	//  [DeployModelRequest.deployed_model][google.cloud.aiplatform.v1.DeployModelRequest.deployed_model]
	//  and for [batch
	//  explanation][google.cloud.aiplatform.v1.BatchPredictionJob.generate_explanation]
	//  by setting
	//  [explanation_spec][google.cloud.aiplatform.v1.BatchPredictionJob.explanation_spec]
	//  of [BatchPredictionJob][google.cloud.aiplatform.v1.BatchPredictionJob].
	// +kcc:proto:field=google.cloud.aiplatform.v1.Model.explanation_spec
	ExplanationSpec *ExplanationSpec `json:"explanationSpec,omitempty"`

	// Used to perform consistent read-modify-write updates. If not set, a blind
	//  "overwrite" update happens.
	// +kcc:proto:field=google.cloud.aiplatform.v1.Model.etag
	Etag *string `json:"etag,omitempty"`

	// The labels with user-defined metadata to organize your Models.
	//
	//  Label keys and values can be no longer than 64 characters
	//  (Unicode codepoints), can only contain lowercase letters, numeric
	//  characters, underscores and dashes. International characters are allowed.
	//
	//  See https://goo.gl/xmQnxf for more information and examples of labels.
	// +kcc:proto:field=google.cloud.aiplatform.v1.Model.labels
	Labels map[string]string `json:"labels,omitempty"`

	// Stats of data used for training or evaluating the Model.
	//
	//  Only populated when the Model is trained by a TrainingPipeline with
	//  [data_input_config][google.cloud.aiplatform.v1.TrainingPipeline.input_data_config].
	// +kcc:proto:field=google.cloud.aiplatform.v1.Model.data_stats
	DataStats *Model_DataStats `json:"dataStats,omitempty"`

	// Customer-managed encryption key spec for a Model. If set, this
	//  Model and all sub-resources of this Model will be secured by this key.
	// +kcc:proto:field=google.cloud.aiplatform.v1.Model.encryption_spec
	EncryptionSpec *EncryptionSpec `json:"encryptionSpec,omitempty"`

	// Optional. User input field to specify the base model source. Currently it
	//  only supports specifying the Model Garden models and Genie models.
	// +kcc:proto:field=google.cloud.aiplatform.v1.Model.base_model_source
	BaseModelSource *Model_BaseModelSource `json:"baseModelSource,omitempty"`

	// Required. The resource name of the Location into which to upload the Model.
	// Format: projects/{project}/locations/{location}
	Parent *Parent `json:"parent,omitempty"`

	// The AIPlatformModel name. If not given, the metadata.name will be used.
	ResourceID *string `json:"resourceID,omitempty"`
}

type Parent struct {

	// Immutable. The location where the model should reside.
	// +required
	Location *string `json:"location,omitempty"`

	// The project that this resource belongs to.
	// +required
	ProjectRef *v1beta1.ProjectRef `json:"projectRef,omitempty"`
}

// AIPlatformModelStatus defines the config connector machine state of AIPlatformModel
type AIPlatformModelStatus struct {
	/* Conditions represent the latest available observations of the
	   object's current state. */
	Conditions []v1alpha1.Condition `json:"conditions,omitempty"`

	// ObservedGeneration is the generation of the resource that was most recently observed by the Config Connector controller. If this is equal to metadata.generation, then that means that the current reported status reflects the most recent desired state of the resource.
	ObservedGeneration *int64 `json:"observedGeneration,omitempty"`

	// A unique specifier for the AIPlatformModel resource in GCP.
	ExternalRef *string `json:"externalRef,omitempty"`

	// ObservedState is the state of the resource as most recently observed in GCP.
	ObservedState *AIPlatformModelObservedState `json:"observedState,omitempty"`
}

// AIPlatformModelObservedState is the state of the AIPlatformModel resource as most recently observed in GCP.
// +kcc:proto=google.cloud.aiplatform.v1.Model
type AIPlatformModelObservedState struct {
	// Output only. Immutable. The version ID of the model.
	//  A new version is committed when a new model version is uploaded or
	//  trained under an existing model id. It is an auto-incrementing decimal
	//  number in string representation.
	// +kcc:proto:field=google.cloud.aiplatform.v1.Model.version_id
	VersionID *string `json:"versionID,omitempty"`

	// Output only. Timestamp when this version was created.
	// +kcc:proto:field=google.cloud.aiplatform.v1.Model.version_create_time
	VersionCreateTime *string `json:"versionCreateTime,omitempty"`

	// Output only. Timestamp when this version was most recently updated.
	// +kcc:proto:field=google.cloud.aiplatform.v1.Model.version_update_time
	VersionUpdateTime *string `json:"versionUpdateTime,omitempty"`

	// Output only. The formats in which this Model may be exported. If empty,
	//  this Model is not available for export.
	// +kcc:proto:field=google.cloud.aiplatform.v1.Model.supported_export_formats
	SupportedExportFormats []Model_ExportFormat `json:"supportedExportFormats,omitempty"`

	// Output only. The resource name of the TrainingPipeline that uploaded this
	//  Model, if any.
	// +kcc:proto:field=google.cloud.aiplatform.v1.Model.training_pipeline
	TrainingPipeline *string `json:"trainingPipeline,omitempty"`

	// Output only. When this Model is deployed, its prediction resources are
	//  described by the `prediction_resources` field of the
	//  [Endpoint.deployed_models][google.cloud.aiplatform.v1.Endpoint.deployed_models]
	//  object. Because not all Models support all resource configuration types,
	//  the configuration types this Model supports are listed here. If no
	//  configuration types are listed, the Model cannot be deployed to an
	//  [Endpoint][google.cloud.aiplatform.v1.Endpoint] and does not support
	//  online predictions
	//  ([PredictionService.Predict][google.cloud.aiplatform.v1.PredictionService.Predict]
	//  or
	//  [PredictionService.Explain][google.cloud.aiplatform.v1.PredictionService.Explain]).
	//  Such a Model can serve predictions by using a
	//  [BatchPredictionJob][google.cloud.aiplatform.v1.BatchPredictionJob], if it
	//  has at least one entry each in
	//  [supported_input_storage_formats][google.cloud.aiplatform.v1.Model.supported_input_storage_formats]
	//  and
	//  [supported_output_storage_formats][google.cloud.aiplatform.v1.Model.supported_output_storage_formats].
	// +kcc:proto:field=google.cloud.aiplatform.v1.Model.supported_deployment_resources_types
	SupportedDeploymentResourcesTypes []string `json:"supportedDeploymentResourcesTypes,omitempty"`

	// Output only. The formats this Model supports in
	//  [BatchPredictionJob.input_config][google.cloud.aiplatform.v1.BatchPredictionJob.input_config].
	//  If
	//  [PredictSchemata.instance_schema_uri][google.cloud.aiplatform.v1.PredictSchemata.instance_schema_uri]
	//  exists, the instances should be given as per that schema.
	//
	//  The possible formats are:
	//
	//  * `jsonl`
	//  The JSON Lines format, where each instance is a single line. Uses
	//  [GcsSource][google.cloud.aiplatform.v1.BatchPredictionJob.InputConfig.gcs_source].
	//
	//  * `csv`
	//  The CSV format, where each instance is a single comma-separated line.
	//  The first line in the file is the header, containing comma-separated field
	//  names. Uses
	//  [GcsSource][google.cloud.aiplatform.v1.BatchPredictionJob.InputConfig.gcs_source].
	//
	//  * `tf-record`
	//  The TFRecord format, where each instance is a single record in tfrecord
	//  syntax. Uses
	//  [GcsSource][google.cloud.aiplatform.v1.BatchPredictionJob.InputConfig.gcs_source].
	//
	//  * `tf-record-gzip`
	//  Similar to `tf-record`, but the file is gzipped. Uses
	//  [GcsSource][google.cloud.aiplatform.v1.BatchPredictionJob.InputConfig.gcs_source].
	//
	//  * `bigquery`
	//  Each instance is a single row in BigQuery. Uses
	//  [BigQuerySource][google.cloud.aiplatform.v1.BatchPredictionJob.InputConfig.bigquery_source].
	//
	//  * `file-list`
	//  Each line of the file is the location of an instance to process, uses
	//  `gcs_source` field of the
	//  [InputConfig][google.cloud.aiplatform.v1.BatchPredictionJob.InputConfig]
	//  object.
	//
	//
	//  If this Model doesn't support any of these formats it means it cannot be
	//  used with a
	//  [BatchPredictionJob][google.cloud.aiplatform.v1.BatchPredictionJob].
	//  However, if it has
	//  [supported_deployment_resources_types][google.cloud.aiplatform.v1.Model.supported_deployment_resources_types],
	//  it could serve online predictions by using
	//  [PredictionService.Predict][google.cloud.aiplatform.v1.PredictionService.Predict]
	//  or
	//  [PredictionService.Explain][google.cloud.aiplatform.v1.PredictionService.Explain].
	// +kcc:proto:field=google.cloud.aiplatform.v1.Model.supported_input_storage_formats
	SupportedInputStorageFormats []string `json:"supportedInputStorageFormats,omitempty"`

	// Output only. The formats this Model supports in
	//  [BatchPredictionJob.output_config][google.cloud.aiplatform.v1.BatchPredictionJob.output_config].
	//  If both
	//  [PredictSchemata.instance_schema_uri][google.cloud.aiplatform.v1.PredictSchemata.instance_schema_uri]
	//  and
	//  [PredictSchemata.prediction_schema_uri][google.cloud.aiplatform.v1.PredictSchemata.prediction_schema_uri]
	//  exist, the predictions are returned together with their instances. In other
	//  words, the prediction has the original instance data first, followed by the
	//  actual prediction content (as per the schema).
	//
	//  The possible formats are:
	//
	//  * `jsonl`
	//  The JSON Lines format, where each prediction is a single line. Uses
	//  [GcsDestination][google.cloud.aiplatform.v1.BatchPredictionJob.OutputConfig.gcs_destination].
	//
	//  * `csv`
	//  The CSV format, where each prediction is a single comma-separated line.
	//  The first line in the file is the header, containing comma-separated field
	//  names. Uses
	//  [GcsDestination][google.cloud.aiplatform.v1.BatchPredictionJob.OutputConfig.gcs_destination].
	//
	//  * `bigquery`
	//  Each prediction is a single row in a BigQuery table, uses
	//  [BigQueryDestination][google.cloud.aiplatform.v1.BatchPredictionJob.OutputConfig.bigquery_destination]
	//  .
	//
	//
	//  If this Model doesn't support any of these formats it means it cannot be
	//  used with a
	//  [BatchPredictionJob][google.cloud.aiplatform.v1.BatchPredictionJob].
	//  However, if it has
	//  [supported_deployment_resources_types][google.cloud.aiplatform.v1.Model.supported_deployment_resources_types],
	//  it could serve online predictions by using
	//  [PredictionService.Predict][google.cloud.aiplatform.v1.PredictionService.Predict]
	//  or
	//  [PredictionService.Explain][google.cloud.aiplatform.v1.PredictionService.Explain].
	// +kcc:proto:field=google.cloud.aiplatform.v1.Model.supported_output_storage_formats
	SupportedOutputStorageFormats []string `json:"supportedOutputStorageFormats,omitempty"`

	// Output only. Timestamp when this Model was uploaded into Vertex AI.
	// +kcc:proto:field=google.cloud.aiplatform.v1.Model.create_time
	CreateTime *string `json:"createTime,omitempty"`

	// Output only. Timestamp when this Model was most recently updated.
	// +kcc:proto:field=google.cloud.aiplatform.v1.Model.update_time
	UpdateTime *string `json:"updateTime,omitempty"`

	// Output only. The pointers to DeployedModels created from this Model. Note
	//  that Model could have been deployed to Endpoints in different Locations.
	// +kcc:proto:field=google.cloud.aiplatform.v1.Model.deployed_models
	DeployedModels []DeployedModelRef `json:"deployedModels,omitempty"`

	// Output only. Source of a model. It can either be automl training pipeline,
	//  custom training pipeline, BigQuery ML, or saved and tuned from Genie or
	//  Model Garden.
	// +kcc:proto:field=google.cloud.aiplatform.v1.Model.model_source_info
	ModelSourceInfo *ModelSourceInfo `json:"modelSourceInfo,omitempty"`

	// Output only. If this Model is a copy of another Model, this contains info
	//  about the original.
	// +kcc:proto:field=google.cloud.aiplatform.v1.Model.original_model_info
	OriginalModelInfo *Model_OriginalModelInfo `json:"originalModelInfo,omitempty"`

	// Output only. The resource name of the Artifact that was created in
	//  MetadataStore when creating the Model. The Artifact resource name pattern
	//  is
	//  `projects/{project}/locations/{location}/metadataStores/{metadata_store}/artifacts/{artifact}`.
	// +kcc:proto:field=google.cloud.aiplatform.v1.Model.metadata_artifact
	MetadataArtifact *string `json:"metadataArtifact,omitempty"`

	// Output only. Reserved for future use.
	// +kcc:proto:field=google.cloud.aiplatform.v1.Model.satisfies_pzs
	SatisfiesPzs *bool `json:"satisfiesPzs,omitempty"`

	// Output only. Reserved for future use.
	// +kcc:proto:field=google.cloud.aiplatform.v1.Model.satisfies_pzi
	SatisfiesPzi *bool `json:"satisfiesPzi,omitempty"`
}

// +genclient
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// TODO(user): make sure the pluralizaiton below is correct
// +kubebuilder:resource:categories=gcp,shortName=gcpaiplatformmodel;gcpaiplatformmodels
// +kubebuilder:subresource:status
// +kubebuilder:metadata:labels="cnrm.cloud.google.com/managed-by-kcc=true";"cnrm.cloud.google.com/system=true"
// +kubebuilder:printcolumn:name="Age",JSONPath=".metadata.creationTimestamp",type="date"
// +kubebuilder:printcolumn:name="Ready",JSONPath=".status.conditions[?(@.type=='Ready')].status",type="string",description="When 'True', the most recent reconcile of the resource succeeded"
// +kubebuilder:printcolumn:name="Status",JSONPath=".status.conditions[?(@.type=='Ready')].reason",type="string",description="The reason for the value in 'Ready'"
// +kubebuilder:printcolumn:name="Status Age",JSONPath=".status.conditions[?(@.type=='Ready')].lastTransitionTime",type="date",description="The last transition time for the value in 'Status'"

// AIPlatformModel is the Schema for the AIPlatformModel API
// +k8s:openapi-gen=true
type AIPlatformModel struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	// +required
	Spec   AIPlatformModelSpec   `json:"spec,omitempty"`
	Status AIPlatformModelStatus `json:"status,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// AIPlatformModelList contains a list of AIPlatformModel
type AIPlatformModelList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []AIPlatformModel `json:"items"`
}

func init() {
	SchemeBuilder.Register(&AIPlatformModel{}, &AIPlatformModelList{})
}
