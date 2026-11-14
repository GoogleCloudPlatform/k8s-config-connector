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

// API sources for AIPlatformModelMonitor, recorded by generate-types:
// +kcc:source:proto=https://github.com/googleapis/googleapis/blob/e7513c828b8db536fa12fc5493d541f323aedf90/google/cloud/aiplatform/v1beta1/model_monitor.proto
// +kcc:source:service-docs=https://cloud.google.com/vertex-ai/
// +kcc:source:resource-docs=https://cloud.google.com/vertex-ai/docs/reference/rest/v1beta1/projects.locations.modelMonitors

package v1alpha1

import (
	parent "github.com/GoogleCloudPlatform/k8s-config-connector/apis/common/parent"
	monitoringv1beta1 "github.com/GoogleCloudPlatform/k8s-config-connector/apis/monitoring/v1beta1"
	vertexaiv1beta1 "github.com/GoogleCloudPlatform/k8s-config-connector/apis/vertexai/v1beta1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/apis/k8s/v1alpha1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

var AIPlatformModelMonitorGVK = GroupVersion.WithKind("AIPlatformModelMonitor")

// AIPlatformModelMonitorSpec defines the desired state of AIPlatformModelMonitor
// +kcc:spec:proto=google.cloud.aiplatform.v1beta1.ModelMonitor
// +kcc:required-from-proto
type AIPlatformModelMonitorSpec struct {
	*parent.ProjectAndLocationRef `json:",inline"`

	// The AIPlatformModelMonitor name. If not given, the metadata.name will be used.
	ResourceID *string `json:"resourceID,omitempty"`

	// Optional default tabular model monitoring objective.
	// +kcc:proto:field=google.cloud.aiplatform.v1beta1.ModelMonitor.tabular_objective
	TabularObjective *ModelMonitoringObjectiveSpec_TabularObjectiveWithRequired `json:"tabularObjective,omitempty"`

	// The display name of the ModelMonitor.
	//  The name can be up to 128 characters long and can consist of any UTF-8.
	// +kcc:proto:field=google.cloud.aiplatform.v1beta1.ModelMonitor.display_name
	DisplayName *string `json:"displayName,omitempty"`

	// The entity that is subject to analysis.
	//  Currently only models in Vertex AI Model Registry are supported. If you
	//  want to analyze the model which is outside the Vertex AI, you could
	//  register a model in Vertex AI Model Registry using just a display name.
	// +kcc:proto:field=google.cloud.aiplatform.v1beta1.ModelMonitor.model_monitoring_target
	ModelMonitoringTarget *ModelMonitor_ModelMonitoringTarget `json:"modelMonitoringTarget,omitempty"`

	// Optional training dataset used to train the model.
	//  It can serve as a reference dataset to identify changes in production.
	// +kcc:proto:field=google.cloud.aiplatform.v1beta1.ModelMonitor.training_dataset
	TrainingDataset *ModelMonitoringInput `json:"trainingDataset,omitempty"`

	// Optional default notification spec, it can be overridden in the
	//  ModelMonitoringJob notification spec.
	// +kcc:proto:field=google.cloud.aiplatform.v1beta1.ModelMonitor.notification_spec
	NotificationSpec *ModelMonitoringNotificationSpec `json:"notificationSpec,omitempty"`

	// Optional default monitoring metrics/logs export spec, it can be overridden
	//  in the ModelMonitoringJob output spec.
	//  If not specified, a default Google Cloud Storage bucket will be created
	//  under your project.
	// +kcc:proto:field=google.cloud.aiplatform.v1beta1.ModelMonitor.output_spec
	OutputSpec *ModelMonitoringOutputSpec `json:"outputSpec,omitempty"`

	// Optional model explanation spec. It is used for feature attribution
	//  monitoring.
	// +kcc:proto:field=google.cloud.aiplatform.v1beta1.ModelMonitor.explanation_spec
	ExplanationSpec *ExplanationSpecWithRequired `json:"explanationSpec,omitempty"`

	// Monitoring Schema is to specify the model's features, prediction outputs
	//  and ground truth properties. It is used to extract pertinent data from the
	//  dataset and to process features based on their properties.
	//  Make sure that the schema aligns with your dataset, if it does not, we will
	//  be unable to extract data from the dataset.
	//  It is required for most models, but optional for Vertex AI AutoML Tables
	//  unless the schem information is not available.
	// +kcc:proto:field=google.cloud.aiplatform.v1beta1.ModelMonitor.model_monitoring_schema
	ModelMonitoringSchema *ModelMonitoringSchema `json:"modelMonitoringSchema,omitempty"`

	// Customer-managed encryption key spec for a ModelMonitor. If
	//  set, this ModelMonitor and all sub-resources of this
	//  ModelMonitor will be secured by this key.
	// +kcc:proto:field=google.cloud.aiplatform.v1beta1.ModelMonitor.encryption_spec
	EncryptionSpec *EncryptionSpecWithRequired `json:"encryptionSpec,omitempty"`
}

// +kcc:proto=google.cloud.aiplatform.v1beta1.ExplanationSpec
type ExplanationSpecWithRequired struct {
	// Required. Parameters that configure explaining of the Model's predictions.
	// +kcc:proto:field=google.cloud.aiplatform.v1beta1.ExplanationSpec.parameters
	// +required
	Parameters *ExplanationParametersWithRequired `json:"parameters,omitempty"`

	// Optional. Metadata describing the Model's input and output for explanation.
	// +kcc:proto:field=google.cloud.aiplatform.v1beta1.ExplanationSpec.metadata
	Metadata *ExplanationMetadataWithRequired `json:"metadata,omitempty"`
}

// +kcc:proto=google.cloud.aiplatform.v1.ExplanationParameters
type ExplanationParametersWithRequired struct {
	// An attribution method that approximates Shapley values for features that
	//  contribute to the label being predicted. A sampling strategy is used to
	//  approximate the value rather than considering all subsets of features.
	//  Refer to this paper for model details: https://arxiv.org/abs/1306.4265.
	// +kcc:proto:field=google.cloud.aiplatform.v1.ExplanationParameters.sampled_shapley_attribution
	SampledShapleyAttribution *SampledShapleyAttributionWithRequired `json:"sampledShapleyAttribution,omitempty"`

	// An attribution method that computes Aumann-Shapley values taking
	//  advantage of the model's fully differentiable structure. Refer to this
	//  paper for more details: https://arxiv.org/abs/1703.01365
	// +kcc:proto:field=google.cloud.aiplatform.v1.ExplanationParameters.integrated_gradients_attribution
	IntegratedGradientsAttribution *IntegratedGradientsAttributionWithRequired `json:"integratedGradientsAttribution,omitempty"`

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
	XraiAttribution *XraiAttributionWithRequired `json:"xraiAttribution,omitempty"`

	// Example-based explanations that returns the nearest neighbors from the
	//  provided dataset.
	// +kcc:proto:field=google.cloud.aiplatform.v1.ExplanationParameters.examples
	Examples *ExamplesWithRequired `json:"examples,omitempty"`

	// If populated, returns attributions for top K indices of outputs
	//  (defaults to 1). Only applies to Models that predicts more than one outputs
	//  (e,g, multi-class Models). When set to -1, returns explanations for all
	//  outputs.
	// +kcc:proto:field=google.cloud.aiplatform.v1.ExplanationParameters.top_k
	TopK *int32 `json:"topK,omitempty"`
}

// +kcc:proto=google.cloud.aiplatform.v1.ExplanationMetadata
type ExplanationMetadataWithRequired struct {
	// Required. Map from feature names to feature input metadata. Keys are the
	// name of the features. Values are the specification of the feature.
	//
	// An empty InputMetadata is valid. It describes a text feature which has the
	// name specified as the key in
	// [ExplanationMetadata.inputs][google.cloud.aiplatform.v1.ExplanationMetadata.inputs].
	// The baseline of the empty feature is chosen by Vertex AI.
	//
	// For Vertex AI-provided Tensorflow images, the key can be any friendly
	// name of the feature. Once specified,
	// [featureAttributions][google.cloud.aiplatform.v1.Attribution.feature_attributions]
	// are keyed by this key (if not grouped with another feature).
	//
	// For custom images, the key must match with the key in
	// [instance][google.cloud.aiplatform.v1.ExplainRequest.instances].
	// +required
	// +kcc:proto:field=google.cloud.aiplatform.v1.ExplanationMetadata.inputs
	Inputs map[string]*ExplanationMetadata_InputMetadata `json:"inputs,omitempty"`

	// Required. Map from output names to output metadata.
	//
	// For Vertex AI-provided Tensorflow images, keys can be any user defined
	// string that consists of any UTF-8 characters.
	//
	// For custom images, keys are the name of the output field in the prediction
	// to be explained.
	//
	// Currently only one key is allowed.
	// +required
	// +kcc:proto:field=google.cloud.aiplatform.v1.ExplanationMetadata.outputs
	Outputs map[string]*ExplanationMetadata_OutputMetadata `json:"outputs,omitempty"`

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

	// Points to a YAML file stored on Google Cloud Storage describing the
	//  format of the [ExplanationMetadata.latent_space_source].
	// +kcc:proto:field=google.cloud.aiplatform.v1.ExplanationMetadata.latent_space_source
	LatentSpaceSource *string `json:"latentSpaceSource,omitempty"`
}

// +kcc:proto=google.cloud.aiplatform.v1beta1.ModelMonitoringObjectiveSpec.TabularObjective
type ModelMonitoringObjectiveSpec_TabularObjectiveWithRequired struct {
	// Input feature distribution drift monitoring spec.
	// +kcc:proto:field=google.cloud.aiplatform.v1beta1.ModelMonitoringObjectiveSpec.TabularObjective.feature_drift_spec
	FeatureDriftSpec *ModelMonitoringObjectiveSpec_DataDriftSpec `json:"featureDriftSpec,omitempty"`

	// Prediction output distribution drift monitoring spec.
	// +kcc:proto:field=google.cloud.aiplatform.v1beta1.ModelMonitoringObjectiveSpec.TabularObjective.prediction_output_drift_spec
	PredictionOutputDriftSpec *ModelMonitoringObjectiveSpec_DataDriftSpec `json:"predictionOutputDriftSpec,omitempty"`

	// Feature attribution monitoring spec.
	// +kcc:proto:field=google.cloud.aiplatform.v1beta1.ModelMonitoringObjectiveSpec.TabularObjective.feature_attribution_spec
	FeatureAttributionSpec *ModelMonitoringObjectiveSpec_FeatureAttributionSpecWithRequired `json:"featureAttributionSpec,omitempty"`
}

// +kcc:proto=google.cloud.aiplatform.v1beta1.ModelMonitoringObjectiveSpec.FeatureAttributionSpec
type ModelMonitoringObjectiveSpec_FeatureAttributionSpecWithRequired struct {
	// Feature names interested in monitoring.
	//  These should be a subset of the input feature names specified in the
	//  monitoring schema. If the field is not specified all features outlied in
	//  the monitoring schema will be used.
	// +kcc:proto:field=google.cloud.aiplatform.v1beta1.ModelMonitoringObjectiveSpec.FeatureAttributionSpec.features
	Features []string `json:"features,omitempty"`

	// Optional default alert condition applied to all the features.
	// +kcc:proto:field=google.cloud.aiplatform.v1beta1.ModelMonitoringObjectiveSpec.FeatureAttributionSpec.default_alert_condition
	DefaultAlertCondition *ModelMonitoringAlertCondition `json:"defaultAlertCondition,omitempty"`

	// Per feature alert condition will override default alert condition.
	// +kcc:proto:field=google.cloud.aiplatform.v1beta1.ModelMonitoringObjectiveSpec.FeatureAttributionSpec.feature_alert_conditions
	FeatureAlertConditions map[string]ModelMonitoringAlertCondition `json:"featureAlertConditions,omitempty"`

	// The config of resources used by the Model Monitoring during the batch
	//  explanation for non-AutoML models. If not set, `n1-standard-2` machine
	//  type will be used by default.
	// +kcc:proto:field=google.cloud.aiplatform.v1beta1.ModelMonitoringObjectiveSpec.FeatureAttributionSpec.batch_explanation_dedicated_resources
	BatchExplanationDedicatedResources *BatchDedicatedResourcesWithRequired `json:"batchExplanationDedicatedResources,omitempty"`
}

// +kcc:proto=google.cloud.aiplatform.v1beta1.BatchDedicatedResources
type BatchDedicatedResourcesWithRequired struct {
	// Required. Immutable. The specification of a single machine.
	// +kcc:proto:field=google.cloud.aiplatform.v1beta1.BatchDedicatedResources.machine_spec
	// +required
	MachineSpec *MachineSpecWithRequired `json:"machineSpec,omitempty"`

	// Immutable. The number of machine replicas used at the start of the batch
	//  operation. If not set, Vertex AI decides starting number, not greater than
	//  [max_replica_count][google.cloud.aiplatform.v1beta1.BatchDedicatedResources.max_replica_count]
	// +kcc:proto:field=google.cloud.aiplatform.v1beta1.BatchDedicatedResources.starting_replica_count
	StartingReplicaCount *int32 `json:"startingReplicaCount,omitempty"`

	// Immutable. The maximum number of machine replicas the batch operation may
	//  be scaled to. The default value is 10.
	// +kcc:proto:field=google.cloud.aiplatform.v1beta1.BatchDedicatedResources.max_replica_count
	MaxReplicaCount *int32 `json:"maxReplicaCount,omitempty"`

	// Optional. Immutable. If set, use DWS resource to schedule the deployment
	//  workload. reference:
	//  (https://cloud.google.com/blog/products/compute/introducing-dynamic-workload-scheduler)
	// +kcc:proto:field=google.cloud.aiplatform.v1beta1.BatchDedicatedResources.flex_start
	FlexStart *FlexStart `json:"flexStart,omitempty"`

	// Optional. Immutable. If true, the batch workload will be run on spot VM.
	// +kcc:proto:field=google.cloud.aiplatform.v1beta1.BatchDedicatedResources.spot
	Spot *bool `json:"spot,omitempty"`
}

// +kcc:proto=google.cloud.aiplatform.v1beta1.ModelMonitor.ModelMonitoringTarget
type ModelMonitor_ModelMonitoringTarget struct {
	// Model in Vertex AI Model Registry.
	// +kcc:proto:field=google.cloud.aiplatform.v1beta1.ModelMonitor.ModelMonitoringTarget.vertex_model
	VertexModel *ModelMonitor_ModelMonitoringTarget_VertexModelSource `json:"vertexModel,omitempty"`
}

// +kcc:proto=google.cloud.aiplatform.v1beta1.ModelMonitor.ModelMonitoringTarget.VertexModelSource
type ModelMonitor_ModelMonitoringTarget_VertexModelSource struct {
	// Model resource name. Format:
	//  projects/{project}/locations/{location}/models/{model}.
	// +kcc:proto:field=google.cloud.aiplatform.v1beta1.ModelMonitor.ModelMonitoringTarget.VertexModelSource.model
	ModelRef *AIPlatformModelRef `json:"modelRef,omitempty"`

	// Model version id.
	// +kcc:proto:field=google.cloud.aiplatform.v1beta1.ModelMonitor.ModelMonitoringTarget.VertexModelSource.model_version_id
	ModelVersionID *string `json:"modelVersionID,omitempty"`
}

// +kcc:proto=google.cloud.aiplatform.v1beta1.ModelMonitoringNotificationSpec
type ModelMonitoringNotificationSpec struct {
	// Email alert config.
	// +kcc:proto:field=google.cloud.aiplatform.v1beta1.ModelMonitoringNotificationSpec.email_config
	EmailConfig *ModelMonitoringNotificationSpec_EmailConfig `json:"emailConfig,omitempty"`

	// Dump the anomalies to Cloud Logging.
	// +kcc:proto:field=google.cloud.aiplatform.v1beta1.ModelMonitoringNotificationSpec.enable_cloud_logging
	EnableCloudLogging *bool `json:"enableCloudLogging,omitempty"`

	// Notification channel config.
	// +kcc:proto:field=google.cloud.aiplatform.v1beta1.ModelMonitoringNotificationSpec.notification_channel_configs
	NotificationChannelConfigs []ModelMonitoringNotificationSpec_NotificationChannelConfig `json:"notificationChannelConfigs,omitempty"`
}

// +kcc:proto=google.cloud.aiplatform.v1beta1.ModelMonitoringNotificationSpec.EmailConfig
type ModelMonitoringNotificationSpec_EmailConfig struct {
	// The email addresses to send the alerts.
	// +kcc:proto:field=google.cloud.aiplatform.v1beta1.ModelMonitoringNotificationSpec.EmailConfig.user_emails
	UserEmails []string `json:"userEmails,omitempty"`
}

// +kcc:proto=google.cloud.aiplatform.v1beta1.ModelMonitoringNotificationSpec.NotificationChannelConfig
type ModelMonitoringNotificationSpec_NotificationChannelConfig struct {
	// Resource names of the NotificationChannels.
	// +kcc:proto:field=google.cloud.aiplatform.v1beta1.ModelMonitoringNotificationSpec.NotificationChannelConfig.notification_channel
	NotificationChannelRef *monitoringv1beta1.MonitoringNotificationChannelRef `json:"notificationChannelRef,omitempty"`
}

// +kcc:proto=google.cloud.aiplatform.v1beta1.ModelMonitoringInput.ModelMonitoringDataset
type ModelMonitoringInput_ModelMonitoringDataset struct {
	// Resource name of the Vertex AI managed dataset.
	// +kcc:proto:field=google.cloud.aiplatform.v1beta1.ModelMonitoringInput.ModelMonitoringDataset.vertex_dataset
	VertexDatasetRef *vertexaiv1beta1.VertexAIDatasetRef `json:"vertexDatasetRef,omitempty"`

	// Google Cloud Storage data source.
	// +kcc:proto:field=google.cloud.aiplatform.v1beta1.ModelMonitoringInput.ModelMonitoringDataset.gcs_source
	GCSSource *ModelMonitoringInput_ModelMonitoringDataset_ModelMonitoringGCSSource `json:"gcsSource,omitempty"`

	// BigQuery data source.
	// +kcc:proto:field=google.cloud.aiplatform.v1beta1.ModelMonitoringInput.ModelMonitoringDataset.bigquery_source
	BigquerySource *ModelMonitoringInput_ModelMonitoringDataset_ModelMonitoringBigQuerySource `json:"bigquerySource,omitempty"`

	// The timestamp field. Usually for serving data.
	// +kcc:proto:field=google.cloud.aiplatform.v1beta1.ModelMonitoringInput.ModelMonitoringDataset.timestamp_field
	TimestampField *string `json:"timestampField,omitempty"`
}

// +kcc:proto=google.cloud.aiplatform.v1beta1.ModelMonitoringInput.VertexEndpointLogs
type ModelMonitoringInput_VertexEndpointLogs struct {
	// List of endpoint resource names.
	// +kcc:proto:field=google.cloud.aiplatform.v1beta1.ModelMonitoringInput.VertexEndpointLogs.endpoints
	EndpointRefs []vertexaiv1beta1.VertexAIEndpointRef `json:"endpointRefs,omitempty"`
}

// +kcc:proto=google.cloud.aiplatform.v1beta1.ModelMonitoringInput.BatchPredictionOutput
type ModelMonitoringInput_BatchPredictionOutput struct {
	// Vertex AI Batch prediction Job resource name.
	// +kcc:proto:field=google.cloud.aiplatform.v1beta1.ModelMonitoringInput.BatchPredictionOutput.batch_prediction_job
	BatchPredictionJobRef *AIPlatformBatchPredictionJobRef `json:"batchPredictionJobRef,omitempty"`
}

// AIPlatformModelMonitorStatus defines the config connector machine state of AIPlatformModelMonitor
type AIPlatformModelMonitorStatus struct {
	/* Conditions represent the latest available observations of the
	   object's current state. */
	Conditions []v1alpha1.Condition `json:"conditions,omitempty"`

	// ObservedGeneration is the generation of the resource that was most recently observed by the Config Connector controller. If this is equal to metadata.generation, then that means that the current reported status reflects the most recent desired state of the resource.
	ObservedGeneration *int64 `json:"observedGeneration,omitempty"`

	// A unique specifier for the AIPlatformModelMonitor resource in GCP.
	ExternalRef *string `json:"externalRef,omitempty"`

	// ObservedState is the state of the resource as most recently observed in GCP.
	ObservedState *AIPlatformModelMonitorObservedState `json:"observedState,omitempty"`
}

// AIPlatformModelMonitorObservedState is the state of the AIPlatformModelMonitor resource as most recently observed in GCP.
// +kcc:observedstate:proto=google.cloud.aiplatform.v1beta1.ModelMonitor
type AIPlatformModelMonitorObservedState struct {
	// Output only. Timestamp when this ModelMonitor was created.
	// +kcc:proto:field=google.cloud.aiplatform.v1beta1.ModelMonitor.create_time
	CreateTime *string `json:"createTime,omitempty"`

	// Output only. Timestamp when this ModelMonitor was updated most recently.
	// +kcc:proto:field=google.cloud.aiplatform.v1beta1.ModelMonitor.update_time
	UpdateTime *string `json:"updateTime,omitempty"`

	// Output only. Reserved for future use.
	// +kcc:proto:field=google.cloud.aiplatform.v1beta1.ModelMonitor.satisfies_pzs
	SatisfiesPzs *bool `json:"satisfiesPzs,omitempty"`

	// Output only. Reserved for future use.
	// +kcc:proto:field=google.cloud.aiplatform.v1beta1.ModelMonitor.satisfies_pzi
	SatisfiesPzi *bool `json:"satisfiesPzi,omitempty"`
}

// +genclient
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:resource:categories=gcp,shortName=gcpaiplatformmodelmonitor;gcpaiplatformmodelmonitors
// +kubebuilder:subresource:status
// +kubebuilder:metadata:labels="cnrm.cloud.google.com/managed-by-kcc=true"
// +kubebuilder:metadata:labels="cnrm.cloud.google.com/system=true"
// +kubebuilder:printcolumn:name="Age",JSONPath=".metadata.creationTimestamp",type="date"
// +kubebuilder:printcolumn:name="Ready",JSONPath=".status.conditions[?(@.type=='Ready')].status",type="string",description="When 'True', the most recent reconcile of the resource succeeded"
// +kubebuilder:printcolumn:name="Status",JSONPath=".status.conditions[?(@.type=='Ready')].reason",type="string",description="The reason for the value in 'Ready'"
// +kubebuilder:printcolumn:name="Status Age",JSONPath=".status.conditions[?(@.type=='Ready')].lastTransitionTime",type="date",description="The last transition time for the value in 'Status'"

// AIPlatformModelMonitor is the Schema for the AIPlatformModelMonitor API
// +k8s:openapi-gen=true
type AIPlatformModelMonitor struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	// +required
	Spec   AIPlatformModelMonitorSpec   `json:"spec,omitempty"`
	Status AIPlatformModelMonitorStatus `json:"status,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// AIPlatformModelMonitorList contains a list of AIPlatformModelMonitor
type AIPlatformModelMonitorList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []AIPlatformModelMonitor `json:"items"`
}

func init() {
	SchemeBuilder.Register(&AIPlatformModelMonitor{}, &AIPlatformModelMonitorList{})
}
