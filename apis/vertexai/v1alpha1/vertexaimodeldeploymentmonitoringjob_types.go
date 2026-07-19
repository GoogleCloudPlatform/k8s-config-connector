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
	common "github.com/GoogleCloudPlatform/k8s-config-connector/apis/common"
	refsv1beta1 "github.com/GoogleCloudPlatform/k8s-config-connector/apis/refs/v1beta1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/apis/k8s/v1alpha1"
	apiextensionsv1 "k8s.io/apiextensions-apiserver/pkg/apis/apiextensions/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

var VertexAIModelDeploymentMonitoringJobGVK = GroupVersion.WithKind("VertexAIModelDeploymentMonitoringJob")

// VertexAIModelDeploymentMonitoringJobSpec defines the desired state of VertexAIModelDeploymentMonitoringJob
// +kcc:spec:proto=google.cloud.aiplatform.v1beta1.ModelDeploymentMonitoringJob
type VertexAIModelDeploymentMonitoringJobSpec struct {
	// The project that this resource belongs to.
	// +required
	ProjectRef *refsv1beta1.ProjectRef `json:"projectRef"`

	// The location of this resource.
	// +required
	Location *string `json:"location"`

	// The VertexAIModelDeploymentMonitoringJob name. If not given, the metadata.name will be used.
	ResourceID *string `json:"resourceID,omitempty"`

	// Required. The user-defined name of the ModelDeploymentMonitoringJob.
	// The name can be up to 128 characters long and can consist of any UTF-8
	// characters.
	// Display name of a ModelDeploymentMonitoringJob.
	// +kcc:proto:field=google.cloud.aiplatform.v1beta1.ModelDeploymentMonitoringJob.display_name
	// +required
	DisplayName string `json:"displayName"`

	// Required. Endpoint resource name.
	// Format: `projects/{project}/locations/{location}/endpoints/{endpoint}`
	// +kcc:proto:field=google.cloud.aiplatform.v1beta1.ModelDeploymentMonitoringJob.endpoint
	// +required
	EndpointRef *VertexAIEndpointRef `json:"endpointRef"`

	// Required. The config for monitoring objectives. This is a per DeployedModel
	// config. Each DeployedModel needs to be configured separately.
	// +kcc:proto:field=google.cloud.aiplatform.v1beta1.ModelDeploymentMonitoringJob.model_deployment_monitoring_objective_configs
	// +required
	ModelDeploymentMonitoringObjectiveConfigs []ModelDeploymentMonitoringObjectiveConfig `json:"modelDeploymentMonitoringObjectiveConfigs"`

	// Required. Schedule config for running the monitoring job.
	// +kcc:proto:field=google.cloud.aiplatform.v1beta1.ModelDeploymentMonitoringJob.model_deployment_monitoring_schedule_config
	// +required
	ModelDeploymentMonitoringScheduleConfig *ModelDeploymentMonitoringScheduleConfig `json:"modelDeploymentMonitoringScheduleConfig"`

	// Required. Sample Strategy for logging.
	// +kcc:proto:field=google.cloud.aiplatform.v1beta1.ModelDeploymentMonitoringJob.logging_sampling_strategy
	// +required
	LoggingSamplingStrategy *SamplingStrategy `json:"loggingSamplingStrategy"`

	// Alert config for model monitoring.
	// +kcc:proto:field=google.cloud.aiplatform.v1beta1.ModelDeploymentMonitoringJob.model_monitoring_alert_config
	ModelMonitoringAlertConfig *ModelMonitoringAlertConfig `json:"modelMonitoringAlertConfig,omitempty"`

	// YAML schema file uri describing the format of a single instance,
	// which are given to format this Endpoint's prediction (and explanation).
	// If not set, we will generate predict schema from collected predict
	// requests.
	// +kcc:proto:field=google.cloud.aiplatform.v1beta1.ModelDeploymentMonitoringJob.predict_instance_schema_uri
	PredictInstanceSchemaURI *string `json:"predictInstanceSchemaURI,omitempty"`

	// Sample Predict instance, same format as
	// [PredictRequest.instances][google.cloud.aiplatform.v1beta1.PredictRequest.instances],
	// this can be set as a replacement of
	// [ModelDeploymentMonitoringJob.predict_instance_schema_uri][google.cloud.aiplatform.v1beta1.ModelDeploymentMonitoringJob.predict_instance_schema_uri].
	// If not set, we will generate predict schema from collected predict
	// requests.
	// +kcc:proto:field=google.cloud.aiplatform.v1beta1.ModelDeploymentMonitoringJob.sample_predict_instance
	SamplePredictInstance *apiextensionsv1.JSON `json:"samplePredictInstance,omitempty"`

	// YAML schema file uri describing the format of a single instance that you
	// want Tensorflow Data Validation (TFDV) to analyze.
	//
	// If this field is empty, all the feature data types are inferred from
	// [predict_instance_schema_uri][google.cloud.aiplatform.v1beta1.ModelDeploymentMonitoringJob.predict_instance_schema_uri],
	// meaning that TFDV will use the data in the exact format(data type) as
	// prediction request/response.
	// If there are any data type differences between predict instance and TFDV
	// instance, this field can be used to override the schema.
	// For models trained with Vertex AI, this field must be set as all the
	// fields in predict instance formatted as string.
	// +kcc:proto:field=google.cloud.aiplatform.v1beta1.ModelDeploymentMonitoringJob.analysis_instance_schema_uri
	AnalysisInstanceSchemaURI *string `json:"analysisInstanceSchemaURI,omitempty"`

	// The TTL of BigQuery tables in user projects which stores logs.
	// A day is the basic unit of the TTL and we take the ceil of TTL/86400(a
	// day). e.g. { second: 3600} indicates ttl = 1 day.
	// +kcc:proto:field=google.cloud.aiplatform.v1beta1.ModelDeploymentMonitoringJob.log_ttl
	LogTTL *string `json:"logTTL,omitempty"`

	// The labels with user-defined metadata to organize your
	// ModelDeploymentMonitoringJob.
	//
	// Label keys and values can be no longer than 64 characters
	// (Unicode codepoints), can only contain lowercase letters, numeric
	// characters, underscores and dashes. International characters are allowed.
	//
	// See https://goo.gl/xmQnxf for more information and examples of labels.
	// +kcc:proto:field=google.cloud.aiplatform.v1beta1.ModelDeploymentMonitoringJob.labels
	Labels map[string]string `json:"labels,omitempty"`

	// Stats anomalies base folder path.
	// +kcc:proto:field=google.cloud.aiplatform.v1beta1.ModelDeploymentMonitoringJob.stats_anomalies_base_directory
	StatsAnomaliesBaseDirectory *GCSDestination `json:"statsAnomaliesBaseDirectory,omitempty"`

	// Customer-managed encryption key spec for a ModelDeploymentMonitoringJob. If
	// set, this ModelDeploymentMonitoringJob and all sub-resources of this
	// ModelDeploymentMonitoringJob will be secured by this key.
	// +kcc:proto:field=google.cloud.aiplatform.v1beta1.ModelDeploymentMonitoringJob.encryption_spec
	EncryptionSpec *EncryptionSpec `json:"encryptionSpec,omitempty"`

	// If true, the scheduled monitoring pipeline logs are sent to
	// Google Cloud Logging, including pipeline status and anomalies detected.
	// Please note the logs incur cost, which are subject to [Cloud Logging
	// pricing](https://cloud.google.com/logging#pricing).
	// +kcc:proto:field=google.cloud.aiplatform.v1beta1.ModelDeploymentMonitoringJob.enable_monitoring_pipeline_logs
	EnableMonitoringPipelineLogs *bool `json:"enableMonitoringPipelineLogs,omitempty"`
}

// VertexAIModelDeploymentMonitoringJobStatus defines the config connector machine state of VertexAIModelDeploymentMonitoringJob
type VertexAIModelDeploymentMonitoringJobStatus struct {
	/* Conditions represent the latest available observations of the
	   object's current state. */
	Conditions []v1alpha1.Condition `json:"conditions,omitempty"`

	// ObservedGeneration is the generation of the resource that was most recently observed by the Config Connector controller. If this is equal to metadata.generation, then that means that the current reported status reflects the most recent desired state of the resource.
	ObservedGeneration *int64 `json:"observedGeneration,omitempty"`

	// A unique specifier for the VertexAIModelDeploymentMonitoringJob resource in GCP.
	ExternalRef *string `json:"externalRef,omitempty"`

	// ObservedState is the state of the resource as most recently observed in GCP.
	ObservedState *VertexAIModelDeploymentMonitoringJobObservedState `json:"observedState,omitempty"`
}

// VertexAIModelDeploymentMonitoringJobObservedState is the state of the VertexAIModelDeploymentMonitoringJob resource as most recently observed in GCP.
// +kcc:observedstate:proto=google.cloud.aiplatform.v1beta1.ModelDeploymentMonitoringJob
type VertexAIModelDeploymentMonitoringJobObservedState struct {
	// Output only. The detailed state of the monitoring job.
	// When the job is still creating, the state will be 'PENDING'.
	// Once the job is successfully created, the state will be 'RUNNING'.
	// Pause the job, the state will be 'PAUSED'.
	// Resume the job, the state will return to 'RUNNING'.
	// +kcc:proto:field=google.cloud.aiplatform.v1beta1.ModelDeploymentMonitoringJob.state
	State *string `json:"state,omitempty"`

	// Output only. Schedule state when the monitoring job is in Running state.
	// +kcc:proto:field=google.cloud.aiplatform.v1beta1.ModelDeploymentMonitoringJob.schedule_state
	ScheduleState *string `json:"scheduleState,omitempty"`

	// Output only. Latest triggered monitoring pipeline metadata.
	// +kcc:proto:field=google.cloud.aiplatform.v1beta1.ModelDeploymentMonitoringJob.latest_monitoring_pipeline_metadata
	LatestMonitoringPipelineMetadata *ModelDeploymentMonitoringJob_LatestMonitoringPipelineMetadata `json:"latestMonitoringPipelineMetadata,omitempty"`

	// Output only. The created bigquery tables for the job under customer
	// project. Customer could do their own query & analysis. There could be 4 log
	// tables in maximum:
	// 1. Training data logging predict request/response
	// 2. Serving data logging predict request/response
	// +kcc:proto:field=google.cloud.aiplatform.v1beta1.ModelDeploymentMonitoringJob.bigquery_tables
	BigqueryTables []ModelDeploymentMonitoringBigQueryTable `json:"bigqueryTables,omitempty"`

	// Output only. Timestamp when this ModelDeploymentMonitoringJob was created.
	// +kcc:proto:field=google.cloud.aiplatform.v1beta1.ModelDeploymentMonitoringJob.create_time
	CreateTime *string `json:"createTime,omitempty"`

	// Output only. Timestamp when this ModelDeploymentMonitoringJob was updated
	// most recently.
	// +kcc:proto:field=google.cloud.aiplatform.v1beta1.ModelDeploymentMonitoringJob.update_time
	UpdateTime *string `json:"updateTime,omitempty"`

	// Output only. Timestamp when this monitoring pipeline will be scheduled to
	// run for the next round.
	// +kcc:proto:field=google.cloud.aiplatform.v1beta1.ModelDeploymentMonitoringJob.next_schedule_time
	NextScheduleTime *string `json:"nextScheduleTime,omitempty"`

	// Output only. Only populated when the job's state is `JOB_STATE_FAILED` or
	// `JOB_STATE_CANCELLED`.
	// +kcc:proto:field=google.cloud.aiplatform.v1beta1.ModelDeploymentMonitoringJob.error
	Error *common.Status `json:"error,omitempty"`

	// Output only. Reserved for future use.
	// +kcc:proto:field=google.cloud.aiplatform.v1beta1.ModelDeploymentMonitoringJob.satisfies_pzs
	SatisfiesPzs *bool `json:"satisfiesPzs,omitempty"`

	// Output only. Reserved for future use.
	// +kcc:proto:field=google.cloud.aiplatform.v1beta1.ModelDeploymentMonitoringJob.satisfies_pzi
	SatisfiesPzi *bool `json:"satisfiesPzi,omitempty"`
}

// +genclient
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:resource:categories=gcp,shortName=gcpvertexaimodeldeploymentmonitoringjob;gcpvertexaimodeldeploymentmonitoringjobs
// +kubebuilder:subresource:status
// +kubebuilder:metadata:labels="cnrm.cloud.google.com/managed-by-kcc=true"
// +kubebuilder:metadata:labels="cnrm.cloud.google.com/system=true"
// +kubebuilder:metadata:labels="cnrm.cloud.google.com/stability-level=alpha"
// +kubebuilder:printcolumn:name="Age",JSONPath=".metadata.creationTimestamp",type="date"
// +kubebuilder:printcolumn:name="Ready",JSONPath=".status.conditions[?(@.type=='Ready')].status",type="string",description="When 'True', the most recent reconcile of the resource succeeded"
// +kubebuilder:printcolumn:name="Status",JSONPath=".status.conditions[?(@.type=='Ready')].reason",type="string",description="The reason for the value in 'Ready'"
// +kubebuilder:printcolumn:name="Status Age",JSONPath=".status.conditions[?(@.type=='Ready')].lastTransitionTime",type="date",description="The last transition time for the value in 'Status'"

// VertexAIModelDeploymentMonitoringJob is the Schema for the VertexAIModelDeploymentMonitoringJob API
// +k8s:openapi-gen=true
type VertexAIModelDeploymentMonitoringJob struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	// +required
	Spec   VertexAIModelDeploymentMonitoringJobSpec   `json:"spec,omitempty"`
	Status VertexAIModelDeploymentMonitoringJobStatus `json:"status,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// VertexAIModelDeploymentMonitoringJobList contains a list of VertexAIModelDeploymentMonitoringJob
type VertexAIModelDeploymentMonitoringJobList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []VertexAIModelDeploymentMonitoringJob `json:"items"`
}

func init() {
	SchemeBuilder.Register(&VertexAIModelDeploymentMonitoringJob{}, &VertexAIModelDeploymentMonitoringJobList{})
}
