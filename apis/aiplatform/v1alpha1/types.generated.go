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


// +kcc:proto=google.cloud.aiplatform.v1.ContainerSpec
type ContainerSpec struct {
	// Required. The URI of a container image in the Container Registry that is to
	//  be run on each worker replica.
	// +kcc:proto:field=google.cloud.aiplatform.v1.ContainerSpec.image_uri
	ImageURI *string `json:"imageURI,omitempty"`

	// The command to be invoked when the container is started.
	//  It overrides the entrypoint instruction in Dockerfile when provided.
	// +kcc:proto:field=google.cloud.aiplatform.v1.ContainerSpec.command
	Command []string `json:"command,omitempty"`

	// The arguments to be passed when starting the container.
	// +kcc:proto:field=google.cloud.aiplatform.v1.ContainerSpec.args
	Args []string `json:"args,omitempty"`

	// Environment variables to be passed to the container.
	//  Maximum limit is 100.
	// +kcc:proto:field=google.cloud.aiplatform.v1.ContainerSpec.env
	Env []EnvVar `json:"env,omitempty"`
}

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
	ServiceAccount *string `json:"serviceAccount,omitempty"`

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
	Network *string `json:"network,omitempty"`

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
	BaseOutputDirectory *GcsDestination `json:"baseOutputDirectory,omitempty"`

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
	Tensorboard *string `json:"tensorboard,omitempty"`

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
	Models []string `json:"models,omitempty"`
}

// +kcc:proto=google.cloud.aiplatform.v1.DiskSpec
type DiskSpec struct {
	// Type of the boot disk (default is "pd-ssd").
	//  Valid values: "pd-ssd" (Persistent Disk Solid State Drive) or
	//  "pd-standard" (Persistent Disk Hard Disk Drive).
	// +kcc:proto:field=google.cloud.aiplatform.v1.DiskSpec.boot_disk_type
	BootDiskType *string `json:"bootDiskType,omitempty"`

	// Size in GB of the boot disk (default is 100GB).
	// +kcc:proto:field=google.cloud.aiplatform.v1.DiskSpec.boot_disk_size_gb
	BootDiskSizeGB *int32 `json:"bootDiskSizeGB,omitempty"`
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

// +kcc:proto=google.cloud.aiplatform.v1.GcsDestination
type GcsDestination struct {
	// Required. Google Cloud Storage URI to output directory. If the uri doesn't
	//  end with
	//  '/', a '/' will be automatically appended. The directory is created if it
	//  doesn't exist.
	// +kcc:proto:field=google.cloud.aiplatform.v1.GcsDestination.output_uri_prefix
	OutputURIPrefix *string `json:"outputURIPrefix,omitempty"`
}

// +kcc:proto=google.cloud.aiplatform.v1.HyperparameterTuningJob
type HyperparameterTuningJob struct {

	// Required. The display name of the HyperparameterTuningJob.
	//  The name can be up to 128 characters long and can consist of any UTF-8
	//  characters.
	// +kcc:proto:field=google.cloud.aiplatform.v1.HyperparameterTuningJob.display_name
	DisplayName *string `json:"displayName,omitempty"`

	// Required. Study configuration of the HyperparameterTuningJob.
	// +kcc:proto:field=google.cloud.aiplatform.v1.HyperparameterTuningJob.study_spec
	StudySpec *StudySpec `json:"studySpec,omitempty"`

	// Required. The desired total number of Trials.
	// +kcc:proto:field=google.cloud.aiplatform.v1.HyperparameterTuningJob.max_trial_count
	MaxTrialCount *int32 `json:"maxTrialCount,omitempty"`

	// Required. The desired number of Trials to run in parallel.
	// +kcc:proto:field=google.cloud.aiplatform.v1.HyperparameterTuningJob.parallel_trial_count
	ParallelTrialCount *int32 `json:"parallelTrialCount,omitempty"`

	// The number of failed Trials that need to be seen before failing
	//  the HyperparameterTuningJob.
	//
	//  If set to 0, Vertex AI decides how many Trials must fail
	//  before the whole job fails.
	// +kcc:proto:field=google.cloud.aiplatform.v1.HyperparameterTuningJob.max_failed_trial_count
	MaxFailedTrialCount *int32 `json:"maxFailedTrialCount,omitempty"`

	// Required. The spec of a trial job. The same spec applies to the CustomJobs
	//  created in all the trials.
	// +kcc:proto:field=google.cloud.aiplatform.v1.HyperparameterTuningJob.trial_job_spec
	TrialJobSpec *CustomJobSpec `json:"trialJobSpec,omitempty"`

	// The labels with user-defined metadata to organize HyperparameterTuningJobs.
	//
	//  Label keys and values can be no longer than 64 characters
	//  (Unicode codepoints), can only contain lowercase letters, numeric
	//  characters, underscores and dashes. International characters are allowed.
	//
	//  See https://goo.gl/xmQnxf for more information and examples of labels.
	// +kcc:proto:field=google.cloud.aiplatform.v1.HyperparameterTuningJob.labels
	Labels map[string]string `json:"labels,omitempty"`

	// Customer-managed encryption key options for a HyperparameterTuningJob.
	//  If this is set, then all resources created by the HyperparameterTuningJob
	//  will be encrypted with the provided encryption key.
	// +kcc:proto:field=google.cloud.aiplatform.v1.HyperparameterTuningJob.encryption_spec
	EncryptionSpec *EncryptionSpec `json:"encryptionSpec,omitempty"`
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

// +kcc:proto=google.cloud.aiplatform.v1.Measurement
type Measurement struct {
}

// +kcc:proto=google.cloud.aiplatform.v1.Measurement.Metric
type Measurement_Metric struct {
}

// +kcc:proto=google.cloud.aiplatform.v1.NfsMount
type NfsMount struct {
	// Required. IP address of the NFS server.
	// +kcc:proto:field=google.cloud.aiplatform.v1.NfsMount.server
	Server *string `json:"server,omitempty"`

	// Required. Source path exported from NFS server.
	//  Has to start with '/', and combined with the ip address, it indicates
	//  the source mount path in the form of `server:path`
	// +kcc:proto:field=google.cloud.aiplatform.v1.NfsMount.path
	Path *string `json:"path,omitempty"`

	// Required. Destination mount path. The NFS will be mounted for the user
	//  under /mnt/nfs/<mount_point>
	// +kcc:proto:field=google.cloud.aiplatform.v1.NfsMount.mount_point
	MountPoint *string `json:"mountPoint,omitempty"`
}

// +kcc:proto=google.cloud.aiplatform.v1.PythonPackageSpec
type PythonPackageSpec struct {
	// Required. The URI of a container image in Artifact Registry that will run
	//  the provided Python package. Vertex AI provides a wide range of executor
	//  images with pre-installed packages to meet users' various use cases. See
	//  the list of [pre-built containers for
	//  training](https://cloud.google.com/vertex-ai/docs/training/pre-built-containers).
	//  You must use an image from this list.
	// +kcc:proto:field=google.cloud.aiplatform.v1.PythonPackageSpec.executor_image_uri
	ExecutorImageURI *string `json:"executorImageURI,omitempty"`

	// Required. The Google Cloud Storage location of the Python package files
	//  which are the training program and its dependent packages. The maximum
	//  number of package URIs is 100.
	// +kcc:proto:field=google.cloud.aiplatform.v1.PythonPackageSpec.package_uris
	PackageUris []string `json:"packageUris,omitempty"`

	// Required. The Python module name to run after installing the packages.
	// +kcc:proto:field=google.cloud.aiplatform.v1.PythonPackageSpec.python_module
	PythonModule *string `json:"pythonModule,omitempty"`

	// Command line arguments to be passed to the Python task.
	// +kcc:proto:field=google.cloud.aiplatform.v1.PythonPackageSpec.args
	Args []string `json:"args,omitempty"`

	// Environment variables to be passed to the python module.
	//  Maximum limit is 100.
	// +kcc:proto:field=google.cloud.aiplatform.v1.PythonPackageSpec.env
	Env []EnvVar `json:"env,omitempty"`
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

// +kcc:proto=google.cloud.aiplatform.v1.Scheduling
type Scheduling struct {
	// Optional. The maximum job running time. The default is 7 days.
	// +kcc:proto:field=google.cloud.aiplatform.v1.Scheduling.timeout
	Timeout *string `json:"timeout,omitempty"`

	// Optional. Restarts the entire CustomJob if a worker gets restarted.
	//  This feature can be used by distributed training jobs that are not
	//  resilient to workers leaving and joining a job.
	// +kcc:proto:field=google.cloud.aiplatform.v1.Scheduling.restart_job_on_worker_restart
	RestartJobOnWorkerRestart *bool `json:"restartJobOnWorkerRestart,omitempty"`

	// Optional. This determines which type of scheduling strategy to use.
	// +kcc:proto:field=google.cloud.aiplatform.v1.Scheduling.strategy
	Strategy *string `json:"strategy,omitempty"`

	// Optional. Indicates if the job should retry for internal errors after the
	//  job starts running. If true, overrides
	//  `Scheduling.restart_job_on_worker_restart` to false.
	// +kcc:proto:field=google.cloud.aiplatform.v1.Scheduling.disable_retries
	DisableRetries *bool `json:"disableRetries,omitempty"`

	// Optional. This is the maximum duration that a job will wait for the
	//  requested resources to be provisioned if the scheduling strategy is set to
	//  [Strategy.DWS_FLEX_START].
	//  If set to 0, the job will wait indefinitely. The default is 24 hours.
	// +kcc:proto:field=google.cloud.aiplatform.v1.Scheduling.max_wait_duration
	MaxWaitDuration *string `json:"maxWaitDuration,omitempty"`
}

// +kcc:proto=google.cloud.aiplatform.v1.StudySpec
type StudySpec struct {
	// The automated early stopping spec using decay curve rule.
	// +kcc:proto:field=google.cloud.aiplatform.v1.StudySpec.decay_curve_stopping_spec
	DecayCurveStoppingSpec *StudySpec_DecayCurveAutomatedStoppingSpec `json:"decayCurveStoppingSpec,omitempty"`

	// The automated early stopping spec using median rule.
	// +kcc:proto:field=google.cloud.aiplatform.v1.StudySpec.median_automated_stopping_spec
	MedianAutomatedStoppingSpec *StudySpec_MedianAutomatedStoppingSpec `json:"medianAutomatedStoppingSpec,omitempty"`

	// The automated early stopping spec using convex stopping rule.
	// +kcc:proto:field=google.cloud.aiplatform.v1.StudySpec.convex_automated_stopping_spec
	ConvexAutomatedStoppingSpec *StudySpec_ConvexAutomatedStoppingSpec `json:"convexAutomatedStoppingSpec,omitempty"`

	// Required. Metric specs for the Study.
	// +kcc:proto:field=google.cloud.aiplatform.v1.StudySpec.metrics
	Metrics []StudySpec_MetricSpec `json:"metrics,omitempty"`

	// Required. The set of parameters to tune.
	// +kcc:proto:field=google.cloud.aiplatform.v1.StudySpec.parameters
	Parameters []StudySpec_ParameterSpec `json:"parameters,omitempty"`

	// The search algorithm specified for the Study.
	// +kcc:proto:field=google.cloud.aiplatform.v1.StudySpec.algorithm
	Algorithm *string `json:"algorithm,omitempty"`

	// The observation noise level of the study.
	//  Currently only supported by the Vertex AI Vizier service. Not supported by
	//  HyperparameterTuningJob or TrainingPipeline.
	// +kcc:proto:field=google.cloud.aiplatform.v1.StudySpec.observation_noise
	ObservationNoise *string `json:"observationNoise,omitempty"`

	// Describe which measurement selection type will be used
	// +kcc:proto:field=google.cloud.aiplatform.v1.StudySpec.measurement_selection_type
	MeasurementSelectionType *string `json:"measurementSelectionType,omitempty"`

	// Conditions for automated stopping of a Study. Enable automated stopping by
	//  configuring at least one condition.
	// +kcc:proto:field=google.cloud.aiplatform.v1.StudySpec.study_stopping_config
	StudyStoppingConfig *StudySpec_StudyStoppingConfig `json:"studyStoppingConfig,omitempty"`
}

// +kcc:proto=google.cloud.aiplatform.v1.StudySpec.ConvexAutomatedStoppingSpec
type StudySpec_ConvexAutomatedStoppingSpec struct {
	// Steps used in predicting the final objective for early stopped trials. In
	//  general, it's set to be the same as the defined steps in training /
	//  tuning. If not defined, it will learn it from the completed trials. When
	//  use_steps is false, this field is set to the maximum elapsed seconds.
	// +kcc:proto:field=google.cloud.aiplatform.v1.StudySpec.ConvexAutomatedStoppingSpec.max_step_count
	MaxStepCount *int64 `json:"maxStepCount,omitempty"`

	// Minimum number of steps for a trial to complete. Trials which do not have
	//  a measurement with step_count > min_step_count won't be considered for
	//  early stopping. It's ok to set it to 0, and a trial can be early stopped
	//  at any stage. By default, min_step_count is set to be one-tenth of the
	//  max_step_count.
	//  When use_elapsed_duration is true, this field is set to the minimum
	//  elapsed seconds.
	// +kcc:proto:field=google.cloud.aiplatform.v1.StudySpec.ConvexAutomatedStoppingSpec.min_step_count
	MinStepCount *int64 `json:"minStepCount,omitempty"`

	// The minimal number of measurements in a Trial.  Early-stopping checks
	//  will not trigger if less than min_measurement_count+1 completed trials or
	//  pending trials with less than min_measurement_count measurements. If not
	//  defined, the default value is 5.
	// +kcc:proto:field=google.cloud.aiplatform.v1.StudySpec.ConvexAutomatedStoppingSpec.min_measurement_count
	MinMeasurementCount *int64 `json:"minMeasurementCount,omitempty"`

	// The hyper-parameter name used in the tuning job that stands for learning
	//  rate. Leave it blank if learning rate is not in a parameter in tuning.
	//  The learning_rate is used to estimate the objective value of the ongoing
	//  trial.
	// +kcc:proto:field=google.cloud.aiplatform.v1.StudySpec.ConvexAutomatedStoppingSpec.learning_rate_parameter_name
	LearningRateParameterName *string `json:"learningRateParameterName,omitempty"`

	// This bool determines whether or not the rule is applied based on
	//  elapsed_secs or steps. If use_elapsed_duration==false, the early stopping
	//  decision is made according to the predicted objective values according to
	//  the target steps. If use_elapsed_duration==true, elapsed_secs is used
	//  instead of steps. Also, in this case, the parameters max_num_steps and
	//  min_num_steps are overloaded to contain max_elapsed_seconds and
	//  min_elapsed_seconds.
	// +kcc:proto:field=google.cloud.aiplatform.v1.StudySpec.ConvexAutomatedStoppingSpec.use_elapsed_duration
	UseElapsedDuration *bool `json:"useElapsedDuration,omitempty"`

	// ConvexAutomatedStoppingSpec by default only updates the trials that needs
	//  to be early stopped using a newly trained auto-regressive model. When
	//  this flag is set to True, all stopped trials from the beginning are
	//  potentially updated in terms of their `final_measurement`. Also, note
	//  that the training logic of autoregressive models is different in this
	//  case. Enabling this option has shown better results and this may be the
	//  default option in the future.
	// +kcc:proto:field=google.cloud.aiplatform.v1.StudySpec.ConvexAutomatedStoppingSpec.update_all_stopped_trials
	UpdateAllStoppedTrials *bool `json:"updateAllStoppedTrials,omitempty"`
}

// +kcc:proto=google.cloud.aiplatform.v1.StudySpec.DecayCurveAutomatedStoppingSpec
type StudySpec_DecayCurveAutomatedStoppingSpec struct {
	// True if
	//  [Measurement.elapsed_duration][google.cloud.aiplatform.v1.Measurement.elapsed_duration]
	//  is used as the x-axis of each Trials Decay Curve. Otherwise,
	//  [Measurement.step_count][google.cloud.aiplatform.v1.Measurement.step_count]
	//  will be used as the x-axis.
	// +kcc:proto:field=google.cloud.aiplatform.v1.StudySpec.DecayCurveAutomatedStoppingSpec.use_elapsed_duration
	UseElapsedDuration *bool `json:"useElapsedDuration,omitempty"`
}

// +kcc:proto=google.cloud.aiplatform.v1.StudySpec.MedianAutomatedStoppingSpec
type StudySpec_MedianAutomatedStoppingSpec struct {
	// True if median automated stopping rule applies on
	//  [Measurement.elapsed_duration][google.cloud.aiplatform.v1.Measurement.elapsed_duration].
	//  It means that elapsed_duration field of latest measurement of current
	//  Trial is used to compute median objective value for each completed
	//  Trials.
	// +kcc:proto:field=google.cloud.aiplatform.v1.StudySpec.MedianAutomatedStoppingSpec.use_elapsed_duration
	UseElapsedDuration *bool `json:"useElapsedDuration,omitempty"`
}

// +kcc:proto=google.cloud.aiplatform.v1.StudySpec.MetricSpec
type StudySpec_MetricSpec struct {
	// Required. The ID of the metric. Must not contain whitespaces and must be
	//  unique amongst all MetricSpecs.
	// +kcc:proto:field=google.cloud.aiplatform.v1.StudySpec.MetricSpec.metric_id
	MetricID *string `json:"metricID,omitempty"`

	// Required. The optimization goal of the metric.
	// +kcc:proto:field=google.cloud.aiplatform.v1.StudySpec.MetricSpec.goal
	Goal *string `json:"goal,omitempty"`

	// Used for safe search. In the case, the metric will be a safety
	//  metric. You must provide a separate metric for objective metric.
	// +kcc:proto:field=google.cloud.aiplatform.v1.StudySpec.MetricSpec.safety_config
	SafetyConfig *StudySpec_MetricSpec_SafetyMetricConfig `json:"safetyConfig,omitempty"`
}

// +kcc:proto=google.cloud.aiplatform.v1.StudySpec.MetricSpec.SafetyMetricConfig
type StudySpec_MetricSpec_SafetyMetricConfig struct {
	// Safety threshold (boundary value between safe and unsafe). NOTE that if
	//  you leave SafetyMetricConfig unset, a default value of 0 will be used.
	// +kcc:proto:field=google.cloud.aiplatform.v1.StudySpec.MetricSpec.SafetyMetricConfig.safety_threshold
	SafetyThreshold *float64 `json:"safetyThreshold,omitempty"`

	// Desired minimum fraction of safe trials (over total number of trials)
	//  that should be targeted by the algorithm at any time during the
	//  study (best effort). This should be between 0.0 and 1.0 and a value of
	//  0.0 means that there is no minimum and an algorithm proceeds without
	//  targeting any specific fraction. A value of 1.0 means that the
	//  algorithm attempts to only Suggest safe Trials.
	// +kcc:proto:field=google.cloud.aiplatform.v1.StudySpec.MetricSpec.SafetyMetricConfig.desired_min_safe_trials_fraction
	DesiredMinSafeTrialsFraction *float64 `json:"desiredMinSafeTrialsFraction,omitempty"`
}

// +kcc:proto=google.cloud.aiplatform.v1.StudySpec.ParameterSpec
type StudySpec_ParameterSpec struct {
	// The value spec for a 'DOUBLE' parameter.
	// +kcc:proto:field=google.cloud.aiplatform.v1.StudySpec.ParameterSpec.double_value_spec
	DoubleValueSpec *StudySpec_ParameterSpec_DoubleValueSpec `json:"doubleValueSpec,omitempty"`

	// The value spec for an 'INTEGER' parameter.
	// +kcc:proto:field=google.cloud.aiplatform.v1.StudySpec.ParameterSpec.integer_value_spec
	IntegerValueSpec *StudySpec_ParameterSpec_IntegerValueSpec `json:"integerValueSpec,omitempty"`

	// The value spec for a 'CATEGORICAL' parameter.
	// +kcc:proto:field=google.cloud.aiplatform.v1.StudySpec.ParameterSpec.categorical_value_spec
	CategoricalValueSpec *StudySpec_ParameterSpec_CategoricalValueSpec `json:"categoricalValueSpec,omitempty"`

	// The value spec for a 'DISCRETE' parameter.
	// +kcc:proto:field=google.cloud.aiplatform.v1.StudySpec.ParameterSpec.discrete_value_spec
	DiscreteValueSpec *StudySpec_ParameterSpec_DiscreteValueSpec `json:"discreteValueSpec,omitempty"`

	// Required. The ID of the parameter. Must not contain whitespaces and must
	//  be unique amongst all ParameterSpecs.
	// +kcc:proto:field=google.cloud.aiplatform.v1.StudySpec.ParameterSpec.parameter_id
	ParameterID *string `json:"parameterID,omitempty"`

	// How the parameter should be scaled.
	//  Leave unset for `CATEGORICAL` parameters.
	// +kcc:proto:field=google.cloud.aiplatform.v1.StudySpec.ParameterSpec.scale_type
	ScaleType *string `json:"scaleType,omitempty"`

	// A conditional parameter node is active if the parameter's value matches
	//  the conditional node's parent_value_condition.
	//
	//  If two items in conditional_parameter_specs have the same name, they
	//  must have disjoint parent_value_condition.
	// +kcc:proto:field=google.cloud.aiplatform.v1.StudySpec.ParameterSpec.conditional_parameter_specs
	ConditionalParameterSpecs []StudySpec_ParameterSpec_ConditionalParameterSpec `json:"conditionalParameterSpecs,omitempty"`
}

// +kcc:proto=google.cloud.aiplatform.v1.StudySpec.ParameterSpec.CategoricalValueSpec
type StudySpec_ParameterSpec_CategoricalValueSpec struct {
	// Required. The list of possible categories.
	// +kcc:proto:field=google.cloud.aiplatform.v1.StudySpec.ParameterSpec.CategoricalValueSpec.values
	Values []string `json:"values,omitempty"`

	// A default value for a `CATEGORICAL` parameter that is assumed to be a
	//  relatively good starting point.  Unset value signals that there is no
	//  offered starting point.
	//
	//  Currently only supported by the Vertex AI Vizier service. Not supported
	//  by HyperparameterTuningJob or TrainingPipeline.
	// +kcc:proto:field=google.cloud.aiplatform.v1.StudySpec.ParameterSpec.CategoricalValueSpec.default_value
	DefaultValue *string `json:"defaultValue,omitempty"`
}

// +kcc:proto=google.cloud.aiplatform.v1.StudySpec.ParameterSpec.ConditionalParameterSpec
type StudySpec_ParameterSpec_ConditionalParameterSpec struct {
	// The spec for matching values from a parent parameter of
	//  `DISCRETE` type.
	// +kcc:proto:field=google.cloud.aiplatform.v1.StudySpec.ParameterSpec.ConditionalParameterSpec.parent_discrete_values
	ParentDiscreteValues *StudySpec_ParameterSpec_ConditionalParameterSpec_DiscreteValueCondition `json:"parentDiscreteValues,omitempty"`

	// The spec for matching values from a parent parameter of `INTEGER`
	//  type.
	// +kcc:proto:field=google.cloud.aiplatform.v1.StudySpec.ParameterSpec.ConditionalParameterSpec.parent_int_values
	ParentIntValues *StudySpec_ParameterSpec_ConditionalParameterSpec_IntValueCondition `json:"parentIntValues,omitempty"`

	// The spec for matching values from a parent parameter of
	//  `CATEGORICAL` type.
	// +kcc:proto:field=google.cloud.aiplatform.v1.StudySpec.ParameterSpec.ConditionalParameterSpec.parent_categorical_values
	ParentCategoricalValues *StudySpec_ParameterSpec_ConditionalParameterSpec_CategoricalValueCondition `json:"parentCategoricalValues,omitempty"`

	// Required. The spec for a conditional parameter.
	// +kcc:proto:field=google.cloud.aiplatform.v1.StudySpec.ParameterSpec.ConditionalParameterSpec.parameter_spec
	ParameterSpec *StudySpec_ParameterSpec `json:"parameterSpec,omitempty"`
}

// +kcc:proto=google.cloud.aiplatform.v1.StudySpec.ParameterSpec.ConditionalParameterSpec.CategoricalValueCondition
type StudySpec_ParameterSpec_ConditionalParameterSpec_CategoricalValueCondition struct {
	// Required. Matches values of the parent parameter of 'CATEGORICAL'
	//  type. All values must exist in `categorical_value_spec` of parent
	//  parameter.
	// +kcc:proto:field=google.cloud.aiplatform.v1.StudySpec.ParameterSpec.ConditionalParameterSpec.CategoricalValueCondition.values
	Values []string `json:"values,omitempty"`
}

// +kcc:proto=google.cloud.aiplatform.v1.StudySpec.ParameterSpec.ConditionalParameterSpec.DiscreteValueCondition
type StudySpec_ParameterSpec_ConditionalParameterSpec_DiscreteValueCondition struct {
	// Required. Matches values of the parent parameter of 'DISCRETE' type.
	//  All values must exist in `discrete_value_spec` of parent parameter.
	//
	//  The Epsilon of the value matching is 1e-10.
	// +kcc:proto:field=google.cloud.aiplatform.v1.StudySpec.ParameterSpec.ConditionalParameterSpec.DiscreteValueCondition.values
	Values []float64 `json:"values,omitempty"`
}

// +kcc:proto=google.cloud.aiplatform.v1.StudySpec.ParameterSpec.ConditionalParameterSpec.IntValueCondition
type StudySpec_ParameterSpec_ConditionalParameterSpec_IntValueCondition struct {
	// Required. Matches values of the parent parameter of 'INTEGER' type.
	//  All values must lie in `integer_value_spec` of parent parameter.
	// +kcc:proto:field=google.cloud.aiplatform.v1.StudySpec.ParameterSpec.ConditionalParameterSpec.IntValueCondition.values
	Values []int64 `json:"values,omitempty"`
}

// +kcc:proto=google.cloud.aiplatform.v1.StudySpec.ParameterSpec.DiscreteValueSpec
type StudySpec_ParameterSpec_DiscreteValueSpec struct {
	// Required. A list of possible values.
	//  The list should be in increasing order and at least 1e-10 apart.
	//  For instance, this parameter might have possible settings of 1.5, 2.5,
	//  and 4.0. This list should not contain more than 1,000 values.
	// +kcc:proto:field=google.cloud.aiplatform.v1.StudySpec.ParameterSpec.DiscreteValueSpec.values
	Values []float64 `json:"values,omitempty"`

	// A default value for a `DISCRETE` parameter that is assumed to be a
	//  relatively good starting point.  Unset value signals that there is no
	//  offered starting point.  It automatically rounds to the
	//  nearest feasible discrete point.
	//
	//  Currently only supported by the Vertex AI Vizier service. Not supported
	//  by HyperparameterTuningJob or TrainingPipeline.
	// +kcc:proto:field=google.cloud.aiplatform.v1.StudySpec.ParameterSpec.DiscreteValueSpec.default_value
	DefaultValue *float64 `json:"defaultValue,omitempty"`
}

// +kcc:proto=google.cloud.aiplatform.v1.StudySpec.ParameterSpec.DoubleValueSpec
type StudySpec_ParameterSpec_DoubleValueSpec struct {
	// Required. Inclusive minimum value of the parameter.
	// +kcc:proto:field=google.cloud.aiplatform.v1.StudySpec.ParameterSpec.DoubleValueSpec.min_value
	MinValue *float64 `json:"minValue,omitempty"`

	// Required. Inclusive maximum value of the parameter.
	// +kcc:proto:field=google.cloud.aiplatform.v1.StudySpec.ParameterSpec.DoubleValueSpec.max_value
	MaxValue *float64 `json:"maxValue,omitempty"`

	// A default value for a `DOUBLE` parameter that is assumed to be a
	//  relatively good starting point.  Unset value signals that there is no
	//  offered starting point.
	//
	//  Currently only supported by the Vertex AI Vizier service. Not supported
	//  by HyperparameterTuningJob or TrainingPipeline.
	// +kcc:proto:field=google.cloud.aiplatform.v1.StudySpec.ParameterSpec.DoubleValueSpec.default_value
	DefaultValue *float64 `json:"defaultValue,omitempty"`
}

// +kcc:proto=google.cloud.aiplatform.v1.StudySpec.ParameterSpec.IntegerValueSpec
type StudySpec_ParameterSpec_IntegerValueSpec struct {
	// Required. Inclusive minimum value of the parameter.
	// +kcc:proto:field=google.cloud.aiplatform.v1.StudySpec.ParameterSpec.IntegerValueSpec.min_value
	MinValue *int64 `json:"minValue,omitempty"`

	// Required. Inclusive maximum value of the parameter.
	// +kcc:proto:field=google.cloud.aiplatform.v1.StudySpec.ParameterSpec.IntegerValueSpec.max_value
	MaxValue *int64 `json:"maxValue,omitempty"`

	// A default value for an `INTEGER` parameter that is assumed to be a
	//  relatively good starting point.  Unset value signals that there is no
	//  offered starting point.
	//
	//  Currently only supported by the Vertex AI Vizier service. Not supported
	//  by HyperparameterTuningJob or TrainingPipeline.
	// +kcc:proto:field=google.cloud.aiplatform.v1.StudySpec.ParameterSpec.IntegerValueSpec.default_value
	DefaultValue *int64 `json:"defaultValue,omitempty"`
}

// +kcc:proto=google.cloud.aiplatform.v1.StudySpec.StudyStoppingConfig
type StudySpec_StudyStoppingConfig struct {
	// If true, a Study enters STOPPING_ASAP whenever it would normally enters
	//  STOPPING state.
	//
	//  The bottom line is: set to true if you want to interrupt on-going
	//  evaluations of Trials as soon as the study stopping condition is met.
	//  (Please see Study.State documentation for the source of truth).
	// +kcc:proto:field=google.cloud.aiplatform.v1.StudySpec.StudyStoppingConfig.should_stop_asap
	ShouldStopAsap *bool `json:"shouldStopAsap,omitempty"`

	// Each "stopping rule" in this proto specifies an "if" condition. Before
	//  Vizier would generate a new suggestion, it first checks each specified
	//  stopping rule, from top to bottom in this list.
	//  Note that the first few rules (e.g. minimum_runtime_constraint,
	//  min_num_trials) will prevent other stopping rules from being evaluated
	//  until they are met. For example, setting `min_num_trials=5` and
	//  `always_stop_after= 1 hour` means that the Study will ONLY stop after it
	//  has 5 COMPLETED trials, even if more than an hour has passed since its
	//  creation. It follows the first applicable rule (whose "if" condition is
	//  satisfied) to make a stopping decision. If none of the specified rules
	//  are applicable, then Vizier decides that the study should not stop.
	//  If Vizier decides that the study should stop, the study enters
	//  STOPPING state (or STOPPING_ASAP if should_stop_asap = true).
	//  IMPORTANT: The automatic study state transition happens precisely as
	//  described above; that is, deleting trials or updating StudyConfig NEVER
	//  automatically moves the study state back to ACTIVE. If you want to
	//  _resume_ a Study that was stopped, 1) change the stopping conditions if
	//  necessary, 2) activate the study, and then 3) ask for suggestions.
	//  If the specified time or duration has not passed, do not stop the
	//  study.
	// +kcc:proto:field=google.cloud.aiplatform.v1.StudySpec.StudyStoppingConfig.minimum_runtime_constraint
	MinimumRuntimeConstraint *StudyTimeConstraint `json:"minimumRuntimeConstraint,omitempty"`

	// If the specified time or duration has passed, stop the study.
	// +kcc:proto:field=google.cloud.aiplatform.v1.StudySpec.StudyStoppingConfig.maximum_runtime_constraint
	MaximumRuntimeConstraint *StudyTimeConstraint `json:"maximumRuntimeConstraint,omitempty"`

	// If there are fewer than this many COMPLETED trials, do not stop the
	//  study.
	// +kcc:proto:field=google.cloud.aiplatform.v1.StudySpec.StudyStoppingConfig.min_num_trials
	MinNumTrials *Int32Value `json:"minNumTrials,omitempty"`

	// If there are more than this many trials, stop the study.
	// +kcc:proto:field=google.cloud.aiplatform.v1.StudySpec.StudyStoppingConfig.max_num_trials
	MaxNumTrials *Int32Value `json:"maxNumTrials,omitempty"`

	// If the objective value has not improved for this many consecutive
	//  trials, stop the study.
	//
	//  WARNING: Effective only for single-objective studies.
	// +kcc:proto:field=google.cloud.aiplatform.v1.StudySpec.StudyStoppingConfig.max_num_trials_no_progress
	MaxNumTrialsNoProgress *Int32Value `json:"maxNumTrialsNoProgress,omitempty"`

	// If the objective value has not improved for this much time, stop the
	//  study.
	//
	//  WARNING: Effective only for single-objective studies.
	// +kcc:proto:field=google.cloud.aiplatform.v1.StudySpec.StudyStoppingConfig.max_duration_no_progress
	MaxDurationNoProgress *string `json:"maxDurationNoProgress,omitempty"`
}

// +kcc:proto=google.cloud.aiplatform.v1.StudyTimeConstraint
type StudyTimeConstraint struct {
	// Counts the wallclock time passed since the creation of this Study.
	// +kcc:proto:field=google.cloud.aiplatform.v1.StudyTimeConstraint.max_duration
	MaxDuration *string `json:"maxDuration,omitempty"`

	// Compares the wallclock time to this time. Must use UTC timezone.
	// +kcc:proto:field=google.cloud.aiplatform.v1.StudyTimeConstraint.end_time
	EndTime *string `json:"endTime,omitempty"`
}

// +kcc:proto=google.cloud.aiplatform.v1.Trial
type Trial struct {
}

// +kcc:proto=google.cloud.aiplatform.v1.Trial.Parameter
type Trial_Parameter struct {
}

// +kcc:proto=google.cloud.aiplatform.v1.WorkerPoolSpec
type WorkerPoolSpec struct {
	// The custom container task.
	// +kcc:proto:field=google.cloud.aiplatform.v1.WorkerPoolSpec.container_spec
	ContainerSpec *ContainerSpec `json:"containerSpec,omitempty"`

	// The Python packaged task.
	// +kcc:proto:field=google.cloud.aiplatform.v1.WorkerPoolSpec.python_package_spec
	PythonPackageSpec *PythonPackageSpec `json:"pythonPackageSpec,omitempty"`

	// Optional. Immutable. The specification of a single machine.
	// +kcc:proto:field=google.cloud.aiplatform.v1.WorkerPoolSpec.machine_spec
	MachineSpec *MachineSpec `json:"machineSpec,omitempty"`

	// Optional. The number of worker replicas to use for this worker pool.
	// +kcc:proto:field=google.cloud.aiplatform.v1.WorkerPoolSpec.replica_count
	ReplicaCount *int64 `json:"replicaCount,omitempty"`

	// Optional. List of NFS mount spec.
	// +kcc:proto:field=google.cloud.aiplatform.v1.WorkerPoolSpec.nfs_mounts
	NfsMounts []NfsMount `json:"nfsMounts,omitempty"`

	// Disk spec.
	// +kcc:proto:field=google.cloud.aiplatform.v1.WorkerPoolSpec.disk_spec
	DiskSpec *DiskSpec `json:"diskSpec,omitempty"`
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

// +kcc:proto=google.protobuf.Int32Value
type Int32Value struct {
	// The int32 value.
	// +kcc:proto:field=google.protobuf.Int32Value.value
	Value *int32 `json:"value,omitempty"`
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

// +kcc:proto=google.cloud.aiplatform.v1.HyperparameterTuningJob
type HyperparameterTuningJobObservedState struct {
	// Output only. Resource name of the HyperparameterTuningJob.
	// +kcc:proto:field=google.cloud.aiplatform.v1.HyperparameterTuningJob.name
	Name *string `json:"name,omitempty"`

	// Output only. Trials of the HyperparameterTuningJob.
	// +kcc:proto:field=google.cloud.aiplatform.v1.HyperparameterTuningJob.trials
	Trials []Trial `json:"trials,omitempty"`

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
	Error *Status `json:"error,omitempty"`

	// Output only. Reserved for future use.
	// +kcc:proto:field=google.cloud.aiplatform.v1.HyperparameterTuningJob.satisfies_pzs
	SatisfiesPzs *bool `json:"satisfiesPzs,omitempty"`

	// Output only. Reserved for future use.
	// +kcc:proto:field=google.cloud.aiplatform.v1.HyperparameterTuningJob.satisfies_pzi
	SatisfiesPzi *bool `json:"satisfiesPzi,omitempty"`
}

// +kcc:proto=google.cloud.aiplatform.v1.Measurement
type MeasurementObservedState struct {
	// Output only. Time that the Trial has been running at the point of this
	//  Measurement.
	// +kcc:proto:field=google.cloud.aiplatform.v1.Measurement.elapsed_duration
	ElapsedDuration *string `json:"elapsedDuration,omitempty"`

	// Output only. The number of steps the machine learning model has been
	//  trained for. Must be non-negative.
	// +kcc:proto:field=google.cloud.aiplatform.v1.Measurement.step_count
	StepCount *int64 `json:"stepCount,omitempty"`

	// Output only. A list of metrics got by evaluating the objective functions
	//  using suggested Parameter values.
	// +kcc:proto:field=google.cloud.aiplatform.v1.Measurement.metrics
	Metrics []Measurement_Metric `json:"metrics,omitempty"`
}

// +kcc:proto=google.cloud.aiplatform.v1.Measurement.Metric
type Measurement_MetricObservedState struct {
	// Output only. The ID of the Metric. The Metric should be defined in
	//  [StudySpec's Metrics][google.cloud.aiplatform.v1.StudySpec.metrics].
	// +kcc:proto:field=google.cloud.aiplatform.v1.Measurement.Metric.metric_id
	MetricID *string `json:"metricID,omitempty"`

	// Output only. The value for this metric.
	// +kcc:proto:field=google.cloud.aiplatform.v1.Measurement.Metric.value
	Value *float64 `json:"value,omitempty"`
}

// +kcc:proto=google.cloud.aiplatform.v1.Trial
type TrialObservedState struct {
	// Output only. Resource name of the Trial assigned by the service.
	// +kcc:proto:field=google.cloud.aiplatform.v1.Trial.name
	Name *string `json:"name,omitempty"`

	// Output only. The identifier of the Trial assigned by the service.
	// +kcc:proto:field=google.cloud.aiplatform.v1.Trial.id
	ID *string `json:"id,omitempty"`

	// Output only. The detailed state of the Trial.
	// +kcc:proto:field=google.cloud.aiplatform.v1.Trial.state
	State *string `json:"state,omitempty"`

	// Output only. The parameters of the Trial.
	// +kcc:proto:field=google.cloud.aiplatform.v1.Trial.parameters
	Parameters []Trial_Parameter `json:"parameters,omitempty"`

	// Output only. The final measurement containing the objective value.
	// +kcc:proto:field=google.cloud.aiplatform.v1.Trial.final_measurement
	FinalMeasurement *Measurement `json:"finalMeasurement,omitempty"`

	// Output only. A list of measurements that are strictly lexicographically
	//  ordered by their induced tuples (steps, elapsed_duration).
	//  These are used for early stopping computations.
	// +kcc:proto:field=google.cloud.aiplatform.v1.Trial.measurements
	Measurements []Measurement `json:"measurements,omitempty"`

	// Output only. Time when the Trial was started.
	// +kcc:proto:field=google.cloud.aiplatform.v1.Trial.start_time
	StartTime *string `json:"startTime,omitempty"`

	// Output only. Time when the Trial's status changed to `SUCCEEDED` or
	//  `INFEASIBLE`.
	// +kcc:proto:field=google.cloud.aiplatform.v1.Trial.end_time
	EndTime *string `json:"endTime,omitempty"`

	// Output only. The identifier of the client that originally requested this
	//  Trial. Each client is identified by a unique client_id. When a client asks
	//  for a suggestion, Vertex AI Vizier will assign it a Trial. The client
	//  should evaluate the Trial, complete it, and report back to Vertex AI
	//  Vizier. If suggestion is asked again by same client_id before the Trial is
	//  completed, the same Trial will be returned. Multiple clients with
	//  different client_ids can ask for suggestions simultaneously, each of them
	//  will get their own Trial.
	// +kcc:proto:field=google.cloud.aiplatform.v1.Trial.client_id
	ClientID *string `json:"clientID,omitempty"`

	// Output only. A human readable string describing why the Trial is
	//  infeasible. This is set only if Trial state is `INFEASIBLE`.
	// +kcc:proto:field=google.cloud.aiplatform.v1.Trial.infeasible_reason
	InfeasibleReason *string `json:"infeasibleReason,omitempty"`

	// Output only. The CustomJob name linked to the Trial.
	//  It's set for a HyperparameterTuningJob's Trial.
	// +kcc:proto:field=google.cloud.aiplatform.v1.Trial.custom_job
	CustomJob *string `json:"customJob,omitempty"`

	// Output only. URIs for accessing [interactive
	//  shells](https://cloud.google.com/vertex-ai/docs/training/monitor-debug-interactive-shell)
	//  (one URI for each training node). Only available if this trial is part of
	//  a
	//  [HyperparameterTuningJob][google.cloud.aiplatform.v1.HyperparameterTuningJob]
	//  and the job's
	//  [trial_job_spec.enable_web_access][google.cloud.aiplatform.v1.CustomJobSpec.enable_web_access]
	//  field is `true`.
	//
	//  The keys are names of each node used for the trial; for example,
	//  `workerpool0-0` for the primary node, `workerpool1-0` for the first node in
	//  the second worker pool, and `workerpool1-1` for the second node in the
	//  second worker pool.
	//
	//  The values are the URIs for each node's interactive shell.
	// +kcc:proto:field=google.cloud.aiplatform.v1.Trial.web_access_uris
	WebAccessUris map[string]string `json:"webAccessUris,omitempty"`
}

// +kcc:proto=google.cloud.aiplatform.v1.Trial.Parameter
type Trial_ParameterObservedState struct {
	// Output only. The ID of the parameter. The parameter should be defined in
	//  [StudySpec's
	//  Parameters][google.cloud.aiplatform.v1.StudySpec.parameters].
	// +kcc:proto:field=google.cloud.aiplatform.v1.Trial.Parameter.parameter_id
	ParameterID *string `json:"parameterID,omitempty"`

	// Output only. The value of the parameter.
	//  `number_value` will be set if a parameter defined in StudySpec is
	//  in type 'INTEGER', 'DOUBLE' or 'DISCRETE'.
	//  `string_value` will be set if a parameter defined in StudySpec is
	//  in type 'CATEGORICAL'.
	// +kcc:proto:field=google.cloud.aiplatform.v1.Trial.Parameter.value
	Value *Value `json:"value,omitempty"`
}
