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
	aiplatformv1alpha1 "github.com/GoogleCloudPlatform/k8s-config-connector/apis/aiplatform/v1alpha1"
	refsv1beta1 "github.com/GoogleCloudPlatform/k8s-config-connector/apis/refs/v1beta1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/apis/k8s/v1alpha1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"

	"github.com/GoogleCloudPlatform/k8s-config-connector/apis/common"
)

var VertexAIBatchPredictionJobGVK = GroupVersion.WithKind("VertexAIBatchPredictionJob")

// VertexAIBatchPredictionJobSpec defines the desired state of VertexAIBatchPredictionJob
type VertexAIBatchPredictionJobSpec struct {
	// The project that this resource belongs to.
	// +required
	ProjectRef *refsv1beta1.ProjectRef `json:"projectRef"`

	// The location of this resource.
	// +required
	Location *string `json:"location"`

	// The VertexAIBatchPredictionJob name. If not given, the metadata.name will be used.
	ResourceID *string `json:"resourceID,omitempty"`

	// Required. The user-defined name of this BatchPredictionJob.
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
	ModelRef *aiplatformv1alpha1.AIPlatformModelRef `json:"modelRef,omitempty"`

	// Contains model information necessary to perform batch prediction without
	//  requiring uploading to model registry.
	//  Exactly one of model and unmanaged_container_model must be set.
	UnmanagedContainerModel *UnmanagedContainerModel `json:"unmanagedContainerModel,omitempty"`

	// Required. Input configuration of the instances on which predictions are
	//  performed. The schema of any single instance may be specified via the
	//  [Model's][google.cloud.aiplatform.v1beta1.BatchPredictionJob.model]
	//  [PredictSchemata's][google.cloud.aiplatform.v1beta1.Model.predict_schemata]
	//  [instance_schema_uri][google.cloud.aiplatform.v1beta1.PredictSchemata.instance_schema_uri].
	InputConfig *BatchPredictionJob_InputConfig `json:"inputConfig,omitempty"`

	// Configuration for how to convert batch prediction input instances to the
	//  prediction instances that are sent to the Model.
	InstanceConfig *BatchPredictionJob_InstanceConfig `json:"instanceConfig,omitempty"`

	// The parameters that govern the predictions. The schema of the parameters
	//  may be specified via the
	//  [Model's][google.cloud.aiplatform.v1beta1.BatchPredictionJob.model]
	//  [PredictSchemata's][google.cloud.aiplatform.v1beta1.Model.predict_schemata]
	//  [parameters_schema_uri][google.cloud.aiplatform.v1beta1.PredictSchemata.parameters_schema_uri].
	ModelParameters *Value `json:"modelParameters,omitempty"`

	// Required. The Configuration specifying where output predictions should
	//  be written.
	//  The schema of any single prediction may be specified as a concatenation
	//  of [Model's][google.cloud.aiplatform.v1beta1.BatchPredictionJob.model]
	//  [PredictSchemata's][google.cloud.aiplatform.v1beta1.Model.predict_schemata]
	//  [instance_schema_uri][google.cloud.aiplatform.v1beta1.PredictSchemata.instance_schema_uri]
	//  and
	//  [prediction_schema_uri][google.cloud.aiplatform.v1beta1.PredictSchemata.prediction_schema_uri].
	OutputConfig *BatchPredictionJob_OutputConfig `json:"outputConfig,omitempty"`

	// The config of resources used by the Model during the batch prediction. If
	//  the Model
	//  [supports][google.cloud.aiplatform.v1beta1.Model.supported_deployment_resources_types]
	//  DEDICATED_RESOURCES this config may be provided (and the job will use these
	//  resources), if the Model doesn't support AUTOMATIC_RESOURCES, this config
	//  must be provided.
	DedicatedResources *BatchDedicatedResources `json:"dedicatedResources,omitempty"`

	// The service account that the DeployedModel's container runs as. If not
	//  specified, a system generated one will be used, which
	//  has minimal permissions and the custom container, if used, may not have
	//  enough permission to access other Google Cloud resources.
	//
	//  Users deploying the Model must have the `iam.serviceAccounts.actAs`
	//  permission on this service account.
	ServiceAccountRef *refsv1beta1.IAMServiceAccountRef `json:"serviceAccountRef,omitempty"`

	// Immutable. Parameters configuring the batch behavior. Currently only
	//  applicable when
	//  [dedicated_resources][google.cloud.aiplatform.v1beta1.BatchPredictionJob.dedicated_resources]
	//  are used (in other cases Vertex AI does the tuning itself).
	ManualBatchTuningParameters *ManualBatchTuningParameters `json:"manualBatchTuningParameters,omitempty"`

	// Generate explanation with the batch prediction results.
	//
	//  When set to `true`, the batch prediction output changes based on the
	//  `predictions_format` field of the
	//  [BatchPredictionJob.output_config][google.cloud.aiplatform.v1beta1.BatchPredictionJob.output_config]
	//  object:
	//
	//   * `bigquery`: output includes a column named `explanation`. The value
	//     is a struct that conforms to the
	//     [Explanation][google.cloud.aiplatform.v1beta1.Explanation] object.
	//   * `jsonl`: The JSON objects on each line include an additional entry
	//     keyed `explanation`. The value of the entry is a JSON object that
	//     conforms to the
	//     [Explanation][google.cloud.aiplatform.v1beta1.Explanation] object.
	//   * `csv`: Generating explanations for CSV format is not supported.
	//
	//  If this field is set to true, either the
	//  [Model.explanation_spec][google.cloud.aiplatform.v1beta1.Model.explanation_spec]
	//  or
	//  [explanation_spec][google.cloud.aiplatform.v1beta1.BatchPredictionJob.explanation_spec]
	//  must be populated.
	GenerateExplanation *bool `json:"generateExplanation,omitempty"`

	// Explanation configuration for this BatchPredictionJob. Can be
	//  specified only if
	//  [generate_explanation][google.cloud.aiplatform.v1beta1.BatchPredictionJob.generate_explanation]
	//  is set to `true`.
	//
	//  This value overrides the value of
	//  [Model.explanation_spec][google.cloud.aiplatform.v1beta1.Model.explanation_spec].
	//  All fields of
	//  [explanation_spec][google.cloud.aiplatform.v1beta1.BatchPredictionJob.explanation_spec]
	//  are optional in the request. If a field of the
	//  [explanation_spec][google.cloud.aiplatform.v1beta1.BatchPredictionJob.explanation_spec]
	//  object is not populated, the corresponding field of the
	//  [Model.explanation_spec][google.cloud.aiplatform.v1beta1.Model.explanation_spec]
	//  object is inherited.
	ExplanationSpec *ExplanationSpec `json:"explanationSpec,omitempty"`

	// The labels with user-defined metadata to organize BatchPredictionJobs.
	//
	//  Label keys and values can be no longer than 64 characters
	//  (Unicode codepoints), can only contain lowercase letters, numeric
	//  characters, underscores and dashes. International characters are allowed.
	//
	//  See https://goo.gl/xmQnxf for more information and examples of labels.
	Labels map[string]string `json:"labels,omitempty"`

	// Customer-managed encryption key options for a BatchPredictionJob. If this
	//  is set, then all resources created by the BatchPredictionJob will be
	//  encrypted with the provided encryption key.
	EncryptionSpec *EncryptionSpec `json:"encryptionSpec,omitempty"`

	// Model monitoring config will be used for analysis model behaviors, based on
	//  the input and output to the batch prediction job, as well as the provided
	//  training dataset.
	ModelMonitoringConfig *ModelMonitoringConfig `json:"modelMonitoringConfig,omitempty"`

	// Get batch prediction job monitoring statistics.
	ModelMonitoringStatsAnomalies []ModelMonitoringStatsAnomalies `json:"modelMonitoringStatsAnomalies,omitempty"`

	// For custom-trained Models and AutoML Tabular Models, the container of the
	//  DeployedModel instances will send `stderr` and `stdout` streams to
	//  Cloud Logging by default. Please note that the logs incur cost,
	//  which are subject to [Cloud Logging
	//  pricing](https://cloud.google.com/logging/pricing).
	//
	//  User can disable container logging by setting this flag to true.
	DisableContainerLogging *bool `json:"disableContainerLogging,omitempty"`
}

// VertexAIBatchPredictionJobStatus defines the config connector machine state of VertexAIBatchPredictionJob
type VertexAIBatchPredictionJobStatus struct {
	/* Conditions represent the latest available observations of the
	   object's current state. */
	Conditions []v1alpha1.Condition `json:"conditions,omitempty"`

	// ObservedGeneration is the generation of the resource that was most recently observed by the Config Connector controller. If this is equal to metadata.generation, then that means that the current reported status reflects the most recent desired state of the resource.
	ObservedGeneration *int64 `json:"observedGeneration,omitempty"`

	// A unique specifier for the VertexAIBatchPredictionJob resource in GCP.
	ExternalRef *string `json:"externalRef,omitempty"`

	// ObservedState is the state of the resource as most recently observed in GCP.
	ObservedState *VertexAIBatchPredictionJobObservedState `json:"observedState,omitempty"`
}

// VertexAIBatchPredictionJobObservedState is the state of the VertexAIBatchPredictionJob resource as most recently observed in GCP.
type VertexAIBatchPredictionJobObservedState struct {
	// Output only. Resource name of the BatchPredictionJob.
	Name *string `json:"name,omitempty"`

	// Output only. The version ID of the Model that produces the predictions via
	//  this job.
	ModelVersionID *string `json:"modelVersionID,omitempty"`

	// Output only. Information further describing the output of this job.
	OutputInfo *BatchPredictionJob_OutputInfoObservedState `json:"outputInfo,omitempty"`

	// Output only. The detailed state of the job.
	State *string `json:"state,omitempty"`

	// Output only. Only populated when the job's state is JOB_STATE_FAILED or
	//  JOB_STATE_CANCELLED.
	Error *common.Status `json:"error,omitempty"`

	// Output only. Partial failures encountered.
	//  For example, single files that can't be read.
	//  This field never exceeds 20 entries.
	//  Status details fields contain standard Google Cloud error details.
	PartialFailures []common.Status `json:"partialFailures,omitempty"`

	// Output only. Information about resources that had been consumed by this
	//  job. Provided in real time at best effort basis, as well as a final value
	//  once the job completes.
	//
	//  Note: This field currently may be not populated for batch predictions that
	//  use AutoML Models.
	ResourcesConsumed *ResourcesConsumedObservedState `json:"resourcesConsumed,omitempty"`

	// Output only. Statistics on completed and failed prediction instances.
	CompletionStats *CompletionStatsObservedState `json:"completionStats,omitempty"`

	// Output only. Time when the BatchPredictionJob was created.
	CreateTime *string `json:"createTime,omitempty"`

	// Output only. Time when the BatchPredictionJob for the first time entered
	//  the `JOB_STATE_RUNNING` state.
	StartTime *string `json:"startTime,omitempty"`

	// Output only. Time when the BatchPredictionJob entered any of the following
	//  states: `JOB_STATE_SUCCEEDED`, `JOB_STATE_FAILED`, `JOB_STATE_CANCELLED`.
	EndTime *string `json:"endTime,omitempty"`

	// Output only. Time when the BatchPredictionJob was most recently updated.
	UpdateTime *string `json:"updateTime,omitempty"`

	// Output only. The running status of the model monitoring pipeline.
	ModelMonitoringStatus *common.Status `json:"modelMonitoringStatus,omitempty"`

	// Output only. Reserved for future use.
	SatisfiesPzs *bool `json:"satisfiesPzs,omitempty"`

	// Output only. Reserved for future use.
	SatisfiesPzi *bool `json:"satisfiesPzi,omitempty"`
}

// +genclient
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:resource:categories=gcp,shortName=gcpvertexaibatchpredictionjob;gcpvertexaibatchpredictionjobs
// +kubebuilder:subresource:status
// +kubebuilder:metadata:labels="cnrm.cloud.google.com/managed-by-kcc=true"
// +kubebuilder:metadata:labels="cnrm.cloud.google.com/system=true"
// +kubebuilder:metadata:labels="cnrm.cloud.google.com/stability-level=alpha"
// +kubebuilder:printcolumn:name="Age",JSONPath=".metadata.creationTimestamp",type="date"
// +kubebuilder:printcolumn:name="Ready",JSONPath=".status.conditions[?(@.type=='Ready')].status",type="string",description="When 'True', the most recent reconcile of the resource succeeded"
// +kubebuilder:printcolumn:name="Status",JSONPath=".status.conditions[?(@.type=='Ready')].reason",type="string",description="The reason for the value in 'Ready'"
// +kubebuilder:printcolumn:name="Status Age",JSONPath=".status.conditions[?(@.type=='Ready')].lastTransitionTime",type="date",description="The last transition time for the value in 'Status'"

// VertexAIBatchPredictionJob is the Schema for the VertexAIBatchPredictionJob API
// +k8s:openapi-gen=true
type VertexAIBatchPredictionJob struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	// +required
	Spec   VertexAIBatchPredictionJobSpec   `json:"spec,omitempty"`
	Status VertexAIBatchPredictionJobStatus `json:"status,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// VertexAIBatchPredictionJobList contains a list of VertexAIBatchPredictionJob
type VertexAIBatchPredictionJobList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []VertexAIBatchPredictionJob `json:"items"`
}

func init() {
	SchemeBuilder.Register(&VertexAIBatchPredictionJob{}, &VertexAIBatchPredictionJobList{})
}
