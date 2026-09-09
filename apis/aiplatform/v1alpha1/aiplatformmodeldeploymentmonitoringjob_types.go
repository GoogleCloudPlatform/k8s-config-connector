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
	"github.com/GoogleCloudPlatform/k8s-config-connector/apis/common"
	refsv1beta1 "github.com/GoogleCloudPlatform/k8s-config-connector/apis/refs/v1beta1"
	vertexaiv1beta1 "github.com/GoogleCloudPlatform/k8s-config-connector/apis/vertexai/v1beta1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/apis/k8s/v1alpha1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

var AIPlatformModelDeploymentMonitoringJobGVK = GroupVersion.WithKind("AIPlatformModelDeploymentMonitoringJob")

// AIPlatformModelDeploymentMonitoringJobSpec defines the desired state of AIPlatformModelDeploymentMonitoringJob
// +kcc:spec:proto=google.cloud.aiplatform.v1.ModelDeploymentMonitoringJob
type AIPlatformModelDeploymentMonitoringJobSpec struct {
	// The project that this resource belongs to.
	ProjectRef *refsv1beta1.ProjectRef `json:"projectRef"`

	// The location of this resource.
	Location *string `json:"location"`

	// The AIPlatformModelDeploymentMonitoringJob name. If not given, the metadata.name will be used.
	ResourceID *string `json:"resourceID,omitempty"`

	// Required. The user-defined name of the ModelDeploymentMonitoringJob.
	// The name can be up to 128 characters long and can consist of any UTF-8
	// characters.
	// Display name of a ModelDeploymentMonitoringJob.
	// +kcc:proto:field=google.cloud.aiplatform.v1.ModelDeploymentMonitoringJob.display_name
	DisplayName *string `json:"displayName,omitempty"`

	// Required. Endpoint resource name.
	// Format: `projects/{project}/locations/{location}/endpoints/{endpoint}`
	// +kcc:proto:field=google.cloud.aiplatform.v1.ModelDeploymentMonitoringJob.endpoint
	EndpointRef *vertexaiv1beta1.VertexAIEndpointRef `json:"endpointRef"`

	// Required. The config for monitoring objectives. This is a per DeployedModel
	// config. Each DeployedModel needs to be configured separately.
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
	// which are given to format this Endpoint's prediction (and explanation).
	// If not set, we will generate predict schema from collected predict
	// requests.
	// +kcc:proto:field=google.cloud.aiplatform.v1.ModelDeploymentMonitoringJob.predict_instance_schema_uri
	PredictInstanceSchemaURI *string `json:"predictInstanceSchemaURI,omitempty"`

	// Sample Predict instance, same format as
	// [PredictRequest.instances][google.cloud.aiplatform.v1.PredictRequest.instances],
	// this can be set as a replacement of
	// [ModelDeploymentMonitoringJob.predict_instance_schema_uri][google.cloud.aiplatform.v1.ModelDeploymentMonitoringJob.predict_instance_schema_uri].
	// If not set, we will generate predict schema from collected predict
	// requests.
	// +kcc:proto:field=google.cloud.aiplatform.v1.ModelDeploymentMonitoringJob.sample_predict_instance
	SamplePredictInstance *Value `json:"samplePredictInstance,omitempty"`

	// YAML schema file uri describing the format of a single instance that you
	// want Tensorflow Data Validation (TFDV) to analyze.
	//
	// If this field is empty, all the feature data types are inferred from
	// [predict_instance_schema_uri][google.cloud.aiplatform.v1.ModelDeploymentMonitoringJob.predict_instance_schema_uri],
	// meaning that TFDV will use the data in the exact format(data type) as
	// prediction request/response.
	// If there are any data type differences between predict instance and TFDV
	// instance, this field can be used to override the schema.
	// For models trained with Vertex AI, this field must be set as all the
	// fields in predict instance formatted as string.
	// +kcc:proto:field=google.cloud.aiplatform.v1.ModelDeploymentMonitoringJob.analysis_instance_schema_uri
	AnalysisInstanceSchemaURI *string `json:"analysisInstanceSchemaURI,omitempty"`

	// The TTL of BigQuery tables in user projects which stores logs.
	// A day is the basic unit of the TTL and we take the ceil of TTL/86400(a
	// day). e.g. { second: 3600} indicates ttl = 1 day.
	// +kcc:proto:field=google.cloud.aiplatform.v1.ModelDeploymentMonitoringJob.log_ttl
	LogTTL *string `json:"logTTL,omitempty"`

	// The labels with user-defined metadata to organize your
	// ModelDeploymentMonitoringJob.
	//
	// Label keys and values can be no longer than 64 characters
	// (Unicode codepoints), can only contain lowercase letters, numeric
	// characters, underscores and dashes. International characters are allowed.
	//
	// See https://goo.gl/xmQnxf for more information and examples of labels.
	// +kcc:proto:field=google.cloud.aiplatform.v1.ModelDeploymentMonitoringJob.labels
	Labels map[string]string `json:"labels,omitempty"`

	// Stats anomalies base folder path.
	// +kcc:proto:field=google.cloud.aiplatform.v1.ModelDeploymentMonitoringJob.stats_anomalies_base_directory
	StatsAnomaliesBaseDirectory *GCSDestination `json:"statsAnomaliesBaseDirectory,omitempty"`

	// Customer-managed encryption key spec for a ModelDeploymentMonitoringJob. If
	// set, this ModelDeploymentMonitoringJob and all sub-resources of this
	// ModelDeploymentMonitoringJob will be secured by this key.
	// +kcc:proto:field=google.cloud.aiplatform.v1.ModelDeploymentMonitoringJob.encryption_spec
	EncryptionSpec *EncryptionSpec `json:"encryptionSpec,omitempty"`

	// If true, the scheduled monitoring pipeline logs are sent to
	// Google Cloud Logging, including pipeline status and anomalies detected.
	// Please note the logs incur cost, which are subject to [Cloud Logging
	// pricing](https://cloud.google.com/logging#pricing).
	// +kcc:proto:field=google.cloud.aiplatform.v1.ModelDeploymentMonitoringJob.enable_monitoring_pipeline_logs
	EnableMonitoringPipelineLogs *bool `json:"enableMonitoringPipelineLogs,omitempty"`
}

// AIPlatformModelDeploymentMonitoringJobStatus defines the config connector machine state of AIPlatformModelDeploymentMonitoringJob
type AIPlatformModelDeploymentMonitoringJobStatus struct {
	/* Conditions represent the latest available observations of the
	   object's current state. */
	Conditions []v1alpha1.Condition `json:"conditions,omitempty"`

	// ObservedGeneration is the generation of the resource that was most recently observed by the Config Connector controller. If this is equal to metadata.generation, then that means that the current reported status reflects the most recent desired state of the resource.
	ObservedGeneration *int64 `json:"observedGeneration,omitempty"`

	// A unique specifier for the AIPlatformModelDeploymentMonitoringJob resource in GCP.
	ExternalRef *string `json:"externalRef,omitempty"`

	// ObservedState is the state of the resource as most recently observed in GCP.
	ObservedState *AIPlatformModelDeploymentMonitoringJobObservedState `json:"observedState,omitempty"`
}

// AIPlatformModelDeploymentMonitoringJobObservedState is the state of the AIPlatformModelDeploymentMonitoringJob resource as most recently observed in GCP.
// +kcc:observedstate:proto=google.cloud.aiplatform.v1.ModelDeploymentMonitoringJob
type AIPlatformModelDeploymentMonitoringJobObservedState struct {
	// Output only. Resource name of a ModelDeploymentMonitoringJob.
	// +kcc:proto:field=google.cloud.aiplatform.v1.ModelDeploymentMonitoringJob.name
	Name *string `json:"name,omitempty"`

	// +kcc:proto:field=google.cloud.aiplatform.v1.ModelDeploymentMonitoringJob.state
	State *string `json:"state,omitempty"`

	// +kcc:proto:field=google.cloud.aiplatform.v1.ModelDeploymentMonitoringJob.schedule_state
	ScheduleState *string `json:"scheduleState,omitempty"`

	// +kcc:proto:field=google.cloud.aiplatform.v1.ModelDeploymentMonitoringJob.latest_monitoring_pipeline_metadata
	LatestMonitoringPipelineMetadata *ModelDeploymentMonitoringJob_LatestMonitoringPipelineMetadata `json:"latestMonitoringPipelineMetadata,omitempty"`

	// +kcc:proto:field=google.cloud.aiplatform.v1.ModelDeploymentMonitoringJob.bigquery_tables
	BigqueryTables []ModelDeploymentMonitoringBigQueryTable `json:"bigqueryTables,omitempty"`

	// Output only. Timestamp when this ModelDeploymentMonitoringJob was created.
	// +kcc:proto:field=google.cloud.aiplatform.v1.ModelDeploymentMonitoringJob.create_time
	CreateTime *string `json:"createTime,omitempty"`

	// Output only. Timestamp when this ModelDeploymentMonitoringJob was updated
	// +kcc:proto:field=google.cloud.aiplatform.v1.ModelDeploymentMonitoringJob.update_time
	UpdateTime *string `json:"updateTime,omitempty"`

	// +kcc:proto:field=google.cloud.aiplatform.v1.ModelDeploymentMonitoringJob.next_schedule_time
	NextScheduleTime *string `json:"nextScheduleTime,omitempty"`

	// +kcc:proto:field=google.cloud.aiplatform.v1.ModelDeploymentMonitoringJob.error
	Error *common.Status `json:"error,omitempty"`

	// +kcc:proto:field=google.cloud.aiplatform.v1.ModelDeploymentMonitoringJob.satisfies_pzs
	SatisfiesPzs *bool `json:"satisfiesPzs,omitempty"`

	// +kcc:proto:field=google.cloud.aiplatform.v1.ModelDeploymentMonitoringJob.satisfies_pzi
	SatisfiesPzi *bool `json:"satisfiesPzi,omitempty"`
}

// +genclient
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:resource:categories=gcp,shortName=gcpaiplatformmodeldeploymentmonitoringjob;gcpaiplatformmodeldeploymentmonitoringjobs
// +kubebuilder:subresource:status
// +kubebuilder:metadata:labels="cnrm.cloud.google.com/managed-by-kcc=true"
// +kubebuilder:metadata:labels="cnrm.cloud.google.com/system=true"
// +kubebuilder:metadata:labels="cnrm.cloud.google.com/stability-level=alpha"
// +kubebuilder:printcolumn:name="Age",JSONPath=".metadata.creationTimestamp",type="date"
// +kubebuilder:printcolumn:name="Ready",JSONPath=".status.conditions[?(@.type=='Ready')].status",type="string",description="When 'True', the most recent reconcile of the resource succeeded"
// +kubebuilder:printcolumn:name="Status",JSONPath=".status.conditions[?(@.type=='Ready')].reason",type="string",description="The reason for the value in 'Ready'"
// +kubebuilder:printcolumn:name="Status Age",JSONPath=".status.conditions[?(@.type=='Ready')].lastTransitionTime",type="date",description="The last transition time for the value in 'Status'"

// AIPlatformModelDeploymentMonitoringJob is the Schema for the AIPlatformModelDeploymentMonitoringJob API
// +k8s:openapi-gen=true
type AIPlatformModelDeploymentMonitoringJob struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	// +required
	Spec   AIPlatformModelDeploymentMonitoringJobSpec   `json:"spec,omitempty"`
	Status AIPlatformModelDeploymentMonitoringJobStatus `json:"status,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// AIPlatformModelDeploymentMonitoringJobList contains a list of AIPlatformModelDeploymentMonitoringJob
type AIPlatformModelDeploymentMonitoringJobList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []AIPlatformModelDeploymentMonitoringJob `json:"items"`
}

func init() {
	SchemeBuilder.Register(&AIPlatformModelDeploymentMonitoringJob{}, &AIPlatformModelDeploymentMonitoringJobList{})
}
