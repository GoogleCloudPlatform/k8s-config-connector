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
	computev1beta1 "github.com/GoogleCloudPlatform/k8s-config-connector/apis/compute/v1beta1"
	refsv1beta1 "github.com/GoogleCloudPlatform/k8s-config-connector/apis/refs/v1beta1"
	vertexaiv1alpha1 "github.com/GoogleCloudPlatform/k8s-config-connector/apis/vertexai/v1alpha1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/apis/k8s/v1alpha1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

var VertexAIHyperparameterTuningJobGVK = GroupVersion.WithKind("VertexAIHyperparameterTuningJob")

// +kcc:proto=google.cloud.aiplatform.v1.CustomJobSpec
type CustomJobSpec struct {
	// Optional. The ID of the PersistentResource in the same Project and Location
	//  which to run
	//
	//  If this is specified, the job will be run on existing machines held by the
	//  PersistentResource instead of on-demand short-live machines.
	//  The network and CMEK configs on the job should be consistent with those on
	//  the PersistentResource, otherwise, the job will be rejected.
	// +kcc:proto:field=google.cloud.aiplatform.v1.CustomJobSpec.persistent_resource_id
	PersistentResourceID *string `json:"persistentResourceID,omitempty"`

	// Required. The spec of the worker pools including machine type and Docker
	//  image. All worker pools except the first one are optional and can be
	//  skipped by providing an empty value.
	// +kcc:proto:field=google.cloud.aiplatform.v1.CustomJobSpec.worker_pool_specs
	WorkerPoolSpecs []WorkerPoolSpec `json:"workerPoolSpecs,omitempty"`

	// Scheduling options for a CustomJob.
	// +kcc:proto:field=google.cloud.aiplatform.v1.CustomJobSpec.scheduling
	Scheduling *Scheduling `json:"scheduling,omitempty"`

	// Specifies the service account for workload run-as account.
	//  Users submitting jobs must have act-as permission on this run-as account.
	//  If unspecified, the [Vertex AI Custom Code Service
	//  Agent](https://cloud.google.com/vertex-ai/docs/general/access-control#service-agents)
	//  for the CustomJob's project is used.
	// +kcc:proto:field=google.cloud.aiplatform.v1.CustomJobSpec.service_account
	ServiceAccountRef *refsv1beta1.IAMServiceAccountRef `json:"serviceAccountRef,omitempty"`

	// Optional. The full name of the Compute Engine
	//  [network](/compute/docs/networks-and-firewalls#networks) to which the Job
	//  should be peered. For example, `projects/12345/global/networks/myVPC`.
	//  [Format](/compute/docs/reference/rest/v1/networks/insert)
	//  is of the form `projects/{project}/global/networks/{network}`.
	//  Where {project} is a project number, as in `12345`, and {network} is a
	//  network name.
	//
	//  To specify this field, you must have already [configured VPC Network
	//  Peering for Vertex
	//  AI](https://cloud.google.com/vertex-ai/docs/general/vpc-peering).
	//
	//  If this field is left unspecified, the job is not peered with any network.
	// +kcc:proto:field=google.cloud.aiplatform.v1.CustomJobSpec.network
	NetworkRef *computev1beta1.ComputeNetworkRef `json:"networkRef,omitempty"`

	// Optional. A list of names for the reserved ip ranges under the VPC network
	//  that can be used for this job.
	//
	//  If set, we will deploy the job within the provided ip ranges. Otherwise,
	//  the job will be deployed to any ip ranges under the provided VPC
	//  network.
	//
	//  Example: ['vertex-ai-ip-range'].
	// +kcc:proto:field=google.cloud.aiplatform.v1.CustomJobSpec.reserved_ip_ranges
	ReservedIPRanges []string `json:"reservedIPRanges,omitempty"`

	// Optional. Configuration for PSC-I for CustomJob.
	// +kcc:proto:field=google.cloud.aiplatform.v1.CustomJobSpec.psc_interface_config
	PSCInterfaceConfig *PSCInterfaceConfig `json:"pscInterfaceConfig,omitempty"`

	// The Cloud Storage location to store the output of this CustomJob or
	//  HyperparameterTuningJob. For HyperparameterTuningJob,
	//  the baseOutputDirectory of
	//  each child CustomJob backing a Trial is set to a subdirectory of name
	//  [id][google.cloud.aiplatform.v1.Trial.id] under its parent
	//  HyperparameterTuningJob's baseOutputDirectory.
	//
	//  The following Vertex AI environment variables will be passed to
	//  containers or python modules when this field is set:
	//
	//    For CustomJob:
	//
	//    * AIP_MODEL_DIR = `<base_output_directory>/model/`
	//    * AIP_CHECKPOINT_DIR = `<base_output_directory>/checkpoints/`
	//    * AIP_TENSORBOARD_LOG_DIR = `<base_output_directory>/logs/`
	//
	//    For CustomJob backing a Trial of HyperparameterTuningJob:
	//
	//    * AIP_MODEL_DIR = `<base_output_directory>/<trial_id>/model/`
	//    * AIP_CHECKPOINT_DIR = `<base_output_directory>/<trial_id>/checkpoints/`
	//    * AIP_TENSORBOARD_LOG_DIR = `<base_output_directory>/<trial_id>/logs/`
	// +kcc:proto:field=google.cloud.aiplatform.v1.CustomJobSpec.base_output_directory
	BaseOutputDirectory *GCSDestination `json:"baseOutputDirectory,omitempty"`

	// The ID of the location to store protected artifacts. e.g. us-central1.
	//  Populate only when the location is different than CustomJob location.
	//  List of supported locations:
	//  https://cloud.google.com/vertex-ai/docs/general/locations
	// +kcc:proto:field=google.cloud.aiplatform.v1.CustomJobSpec.protected_artifact_location_id
	ProtectedArtifactLocationID *string `json:"protectedArtifactLocationID,omitempty"`

	// Optional. The name of a Vertex AI
	//  [Tensorboard][google.cloud.aiplatform.v1.Tensorboard] resource to which
	//  this CustomJob will upload Tensorboard logs. Format:
	//  `projects/{project}/locations/{location}/tensorboards/{tensorboard}`
	// +kcc:proto:field=google.cloud.aiplatform.v1.CustomJobSpec.tensorboard
	TensorboardRef *vertexaiv1alpha1.VertexAITensorboardRef `json:"tensorboardRef,omitempty"`

	// Optional. Whether you want Vertex AI to enable [interactive shell
	//  access](https://cloud.google.com/vertex-ai/docs/training/monitor-debug-interactive-shell)
	//  to training containers.
	//
	//  If set to `true`, you can access interactive shells at the URIs given
	//  by
	//  [CustomJob.web_access_uris][google.cloud.aiplatform.v1.CustomJob.web_access_uris]
	//  or
	//  [Trial.web_access_uris][google.cloud.aiplatform.v1.Trial.web_access_uris]
	//  (within
	//  [HyperparameterTuningJob.trials][google.cloud.aiplatform.v1.HyperparameterTuningJob.trials]).
	// +kcc:proto:field=google.cloud.aiplatform.v1.CustomJobSpec.enable_web_access
	EnableWebAccess *bool `json:"enableWebAccess,omitempty"`

	// Optional. Whether you want Vertex AI to enable access to the customized
	//  dashboard in training chief container.
	//
	//  If set to `true`, you can access the dashboard at the URIs given
	//  by
	//  [CustomJob.web_access_uris][google.cloud.aiplatform.v1.CustomJob.web_access_uris]
	//  or
	//  [Trial.web_access_uris][google.cloud.aiplatform.v1.Trial.web_access_uris]
	//  (within
	//  [HyperparameterTuningJob.trials][google.cloud.aiplatform.v1.HyperparameterTuningJob.trials]).
	// +kcc:proto:field=google.cloud.aiplatform.v1.CustomJobSpec.enable_dashboard_access
	EnableDashboardAccess *bool `json:"enableDashboardAccess,omitempty"`

	// Optional. The Experiment associated with this job.
	//  Format:
	//  `projects/{project}/locations/{location}/metadataStores/{metadataStores}/contexts/{experiment-name}`
	// +kcc:proto:field=google.cloud.aiplatform.v1.CustomJobSpec.experiment
	Experiment *string `json:"experiment,omitempty"`

	// Optional. The Experiment Run associated with this job.
	//  Format:
	//  `projects/{project}/locations/{location}/metadataStores/{metadataStores}/contexts/{experiment-name}-{experiment-run-name}`
	// +kcc:proto:field=google.cloud.aiplatform.v1.CustomJobSpec.experiment_run
	ExperimentRun *string `json:"experimentRun,omitempty"`

	// Optional. The name of the Model resources for which to generate a mapping
	//  to artifact URIs. Applicable only to some of the Google-provided custom
	//  jobs. Format: `projects/{project}/locations/{location}/models/{model}`
	//
	//  In order to retrieve a specific version of the model, also provide
	//  the version ID or version alias.
	//    Example: `projects/{project}/locations/{location}/models/{model}@2`
	//               or
	//             `projects/{project}/locations/{location}/models/{model}@golden`
	//  If no version ID or alias is specified, the "default" version will be
	//  returned. The "default" version alias is created for the first version of
	//  the model, and can be moved to other versions later on. There will be
	//  exactly one default version.
	// +kcc:proto:field=google.cloud.aiplatform.v1.CustomJobSpec.models
	ModelRefs []AIPlatformModelRef `json:"modelRefs,omitempty"`
}

// VertexAIHyperparameterTuningJobSpec defines the desired state of VertexAIHyperparameterTuningJob
// +kcc:spec:proto=google.cloud.aiplatform.v1.HyperparameterTuningJob
type VertexAIHyperparameterTuningJobSpec struct {
	// The project that this resource belongs to.
	ProjectRef *refsv1beta1.ProjectRef `json:"projectRef"`

	// The location of this resource.
	Location *string `json:"location"`

	// The VertexAIHyperparameterTuningJob name. If not given, the metadata.name will be used.
	ResourceID *string `json:"resourceID,omitempty"`

	// Required. The display name of the HyperparameterTuningJob.
	// The name can be up to 128 characters long and can consist of any UTF-8
	// characters.
	// +kcc:proto:field=google.cloud.aiplatform.v1.HyperparameterTuningJob.display_name
	// +kubebuilder:validation:Required
	DisplayName *string `json:"displayName"`

	// Required. Study configuration of the HyperparameterTuningJob.
	// +kcc:proto:field=google.cloud.aiplatform.v1.HyperparameterTuningJob.study_spec
	// +kubebuilder:validation:Required
	StudySpec *StudySpec `json:"studySpec"`

	// Required. The desired total number of Trials.
	// +kcc:proto:field=google.cloud.aiplatform.v1.HyperparameterTuningJob.max_trial_count
	// +kubebuilder:validation:Required
	MaxTrialCount *int32 `json:"maxTrialCount"`

	// Required. The desired number of Trials to run in parallel.
	// +kcc:proto:field=google.cloud.aiplatform.v1.HyperparameterTuningJob.parallel_trial_count
	// +kubebuilder:validation:Required
	ParallelTrialCount *int32 `json:"parallelTrialCount"`

	// The number of failed Trials that need to be seen before failing
	// the HyperparameterTuningJob.
	//
	// If set to 0, Vertex AI decides how many Trials must fail
	// before the whole job fails.
	// +kcc:proto:field=google.cloud.aiplatform.v1.HyperparameterTuningJob.max_failed_trial_count
	MaxFailedTrialCount *int32 `json:"maxFailedTrialCount,omitempty"`

	// Required. The spec of a trial job. The same spec applies to the CustomJobs
	// created in all the trials.
	// +kcc:proto:field=google.cloud.aiplatform.v1.HyperparameterTuningJob.trial_job_spec
	// +kubebuilder:validation:Required
	TrialJobSpec *CustomJobSpec `json:"trialJobSpec"`

	// The labels with user-defined metadata to organize HyperparameterTuningJobs.
	//
	// Label keys and values can be no longer than 64 characters
	// (Unicode codepoints), can only contain lowercase letters, numeric
	// characters, underscores and dashes. International characters are allowed.
	//
	// See https://goo.gl/xmQnxf for more information and examples of labels.
	// +kcc:proto:field=google.cloud.aiplatform.v1.HyperparameterTuningJob.labels
	Labels map[string]string `json:"labels,omitempty"`

	// Customer-managed encryption key options for a HyperparameterTuningJob.
	// If this is set, then all resources created by the HyperparameterTuningJob
	// will be encrypted with the provided encryption key.
	// +kcc:proto:field=google.cloud.aiplatform.v1.HyperparameterTuningJob.encryption_spec
	EncryptionSpec *EncryptionSpec `json:"encryptionSpec,omitempty"`
}

// VertexAIHyperparameterTuningJobStatus defines the config connector machine state of VertexAIHyperparameterTuningJob
type VertexAIHyperparameterTuningJobStatus struct {
	/* Conditions represent the latest available observations of the
	   object's current state. */
	Conditions []v1alpha1.Condition `json:"conditions,omitempty"`

	// ObservedGeneration is the generation of the resource that was most recently observed by the Config Connector controller. If this is equal to metadata.generation, then that means that the current reported status reflects the most recent desired state of the resource.
	ObservedGeneration *int64 `json:"observedGeneration,omitempty"`

	// A unique specifier for the VertexAIHyperparameterTuningJob resource in GCP.
	ExternalRef *string `json:"externalRef,omitempty"`

	// ObservedState is the state of the resource as most recently observed in GCP.
	ObservedState *VertexAIHyperparameterTuningJobObservedState `json:"observedState,omitempty"`
}

// VertexAIHyperparameterTuningJobObservedState is the state of the VertexAIHyperparameterTuningJob resource as most recently observed in GCP.
// +kcc:observedstate:proto=google.cloud.aiplatform.v1.HyperparameterTuningJob
type VertexAIHyperparameterTuningJobObservedState struct {
	// Output only. Trials of the HyperparameterTuningJob.
	// +kcc:proto:field=google.cloud.aiplatform.v1.HyperparameterTuningJob.trials
	Trials []TrialObservedState `json:"trials,omitempty"`

	// Output only. The detailed state of the job.
	// +kcc:proto:field=google.cloud.aiplatform.v1.HyperparameterTuningJob.state
	State *string `json:"state,omitempty"`

	// Output only. Time when the HyperparameterTuningJob was created.
	// +kcc:proto:field=google.cloud.aiplatform.v1.HyperparameterTuningJob.create_time
	CreateTime *string `json:"createTime,omitempty"`

	// Output only. Time when the HyperparameterTuningJob for the first time
	//  entered the `JOB_STATE_RUNNING` state.
	// +kcc:proto:field=google.cloud.aiplatform.v1.HyperparameterTuningJob.start_time
	StartTime *string `json:"startTime,omitempty"`

	// Output only. Time when the HyperparameterTuningJob entered any of the
	//  following states: `JOB_STATE_SUCCEEDED`, `JOB_STATE_FAILED`,
	//  `JOB_STATE_CANCELLED`.
	// +kcc:proto:field=google.cloud.aiplatform.v1.HyperparameterTuningJob.end_time
	EndTime *string `json:"endTime,omitempty"`

	// Output only. Time when the HyperparameterTuningJob was most recently
	//  updated.
	// +kcc:proto:field=google.cloud.aiplatform.v1.HyperparameterTuningJob.update_time
	UpdateTime *string `json:"updateTime,omitempty"`

	// Output only. Only populated when job's state is JOB_STATE_FAILED or
	//  JOB_STATE_CANCELLED.
	// +kcc:proto:field=google.cloud.aiplatform.v1.HyperparameterTuningJob.error
	Error *common.Status `json:"error,omitempty"`

	// Output only. Reserved for future use.
	// +kcc:proto:field=google.cloud.aiplatform.v1.HyperparameterTuningJob.satisfies_pzs
	SatisfiesPzs *bool `json:"satisfiesPzs,omitempty"`

	// Output only. Reserved for future use.
	// +kcc:proto:field=google.cloud.aiplatform.v1.HyperparameterTuningJob.satisfies_pzi
	SatisfiesPzi *bool `json:"satisfiesPzi,omitempty"`
}

// +genclient
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:resource:categories=gcp,shortName=gcpvertexaihyperparametertuningjob;gcpvertexaihyperparametertuningjobs
// +kubebuilder:subresource:status
// +kubebuilder:metadata:labels="cnrm.cloud.google.com/managed-by-kcc=true"
// +kubebuilder:metadata:labels="cnrm.cloud.google.com/system=true"
// +kubebuilder:printcolumn:name="Age",JSONPath=".metadata.creationTimestamp",type="date"
// +kubebuilder:printcolumn:name="Ready",JSONPath=".status.conditions[?(@.type=='Ready')].status",type="string",description="When 'True', the most recent reconcile of the resource succeeded"
// +kubebuilder:printcolumn:name="Status",JSONPath=".status.conditions[?(@.type=='Ready')].reason",type="string",description="The reason for the value in 'Ready'"
// +kubebuilder:printcolumn:name="Status Age",JSONPath=".status.conditions[?(@.type=='Ready')].lastTransitionTime",type="date",description="The last transition time for the value in 'Status'"

// VertexAIHyperparameterTuningJob is the Schema for the VertexAIHyperparameterTuningJob API
// +k8s:openapi-gen=true
type VertexAIHyperparameterTuningJob struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	// +required
	Spec   VertexAIHyperparameterTuningJobSpec   `json:"spec,omitempty"`
	Status VertexAIHyperparameterTuningJobStatus `json:"status,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// VertexAIHyperparameterTuningJobList contains a list of VertexAIHyperparameterTuningJob
type VertexAIHyperparameterTuningJobList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []VertexAIHyperparameterTuningJob `json:"items"`
}

func init() {
	SchemeBuilder.Register(&VertexAIHyperparameterTuningJob{}, &VertexAIHyperparameterTuningJobList{})
}
