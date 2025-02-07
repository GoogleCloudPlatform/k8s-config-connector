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

// +kcc:proto=google.cloud.aiplatform.v1.NasJob
type NasJob struct {

	// Required. The display name of the NasJob.
	//  The name can be up to 128 characters long and can consist of any UTF-8
	//  characters.
	// +kcc:proto:field=google.cloud.aiplatform.v1.NasJob.display_name
	DisplayName *string `json:"displayName,omitempty"`

	// Required. The specification of a NasJob.
	// +kcc:proto:field=google.cloud.aiplatform.v1.NasJob.nas_job_spec
	NasJobSpec *NasJobSpec `json:"nasJobSpec,omitempty"`

	// The labels with user-defined metadata to organize NasJobs.
	//
	//  Label keys and values can be no longer than 64 characters
	//  (Unicode codepoints), can only contain lowercase letters, numeric
	//  characters, underscores and dashes. International characters are allowed.
	//
	//  See https://goo.gl/xmQnxf for more information and examples of labels.
	// +kcc:proto:field=google.cloud.aiplatform.v1.NasJob.labels
	Labels map[string]string `json:"labels,omitempty"`

	// Customer-managed encryption key options for a NasJob.
	//  If this is set, then all resources created by the NasJob
	//  will be encrypted with the provided encryption key.
	// +kcc:proto:field=google.cloud.aiplatform.v1.NasJob.encryption_spec
	EncryptionSpec *EncryptionSpec `json:"encryptionSpec,omitempty"`

	// Optional. Enable a separation of Custom model training
	//  and restricted image training for tenant project.
	// +kcc:proto:field=google.cloud.aiplatform.v1.NasJob.enable_restricted_image_training
	EnableRestrictedImageTraining *bool `json:"enableRestrictedImageTraining,omitempty"`
}

// +kcc:proto=google.cloud.aiplatform.v1.NasJobOutput
type NasJobOutput struct {
}

// +kcc:proto=google.cloud.aiplatform.v1.NasJobOutput.MultiTrialJobOutput
type NasJobOutput_MultiTrialJobOutput struct {
}

// +kcc:proto=google.cloud.aiplatform.v1.NasJobSpec
type NasJobSpec struct {
	// The spec of multi-trial algorithms.
	// +kcc:proto:field=google.cloud.aiplatform.v1.NasJobSpec.multi_trial_algorithm_spec
	MultiTrialAlgorithmSpec *NasJobSpec_MultiTrialAlgorithmSpec `json:"multiTrialAlgorithmSpec,omitempty"`

	// The ID of the existing NasJob in the same Project and Location
	//  which will be used to resume search. search_space_spec and
	//  nas_algorithm_spec are obtained from previous NasJob hence should not
	//  provide them again for this NasJob.
	// +kcc:proto:field=google.cloud.aiplatform.v1.NasJobSpec.resume_nas_job_id
	ResumeNasJobID *string `json:"resumeNasJobID,omitempty"`

	// It defines the search space for Neural Architecture Search (NAS).
	// +kcc:proto:field=google.cloud.aiplatform.v1.NasJobSpec.search_space_spec
	SearchSpaceSpec *string `json:"searchSpaceSpec,omitempty"`
}

// +kcc:proto=google.cloud.aiplatform.v1.NasJobSpec.MultiTrialAlgorithmSpec
type NasJobSpec_MultiTrialAlgorithmSpec struct {
	// The multi-trial Neural Architecture Search (NAS) algorithm
	//  type. Defaults to `REINFORCEMENT_LEARNING`.
	// +kcc:proto:field=google.cloud.aiplatform.v1.NasJobSpec.MultiTrialAlgorithmSpec.multi_trial_algorithm
	MultiTrialAlgorithm *string `json:"multiTrialAlgorithm,omitempty"`

	// Metric specs for the NAS job.
	//  Validation for this field is done at `multi_trial_algorithm_spec` field.
	// +kcc:proto:field=google.cloud.aiplatform.v1.NasJobSpec.MultiTrialAlgorithmSpec.metric
	Metric *NasJobSpec_MultiTrialAlgorithmSpec_MetricSpec `json:"metric,omitempty"`

	// Required. Spec for search trials.
	// +kcc:proto:field=google.cloud.aiplatform.v1.NasJobSpec.MultiTrialAlgorithmSpec.search_trial_spec
	SearchTrialSpec *NasJobSpec_MultiTrialAlgorithmSpec_SearchTrialSpec `json:"searchTrialSpec,omitempty"`

	// Spec for train trials. Top N [TrainTrialSpec.max_parallel_trial_count]
	//  search trials will be trained for every M
	//  [TrainTrialSpec.frequency] trials searched.
	// +kcc:proto:field=google.cloud.aiplatform.v1.NasJobSpec.MultiTrialAlgorithmSpec.train_trial_spec
	TrainTrialSpec *NasJobSpec_MultiTrialAlgorithmSpec_TrainTrialSpec `json:"trainTrialSpec,omitempty"`
}

// +kcc:proto=google.cloud.aiplatform.v1.NasJobSpec.MultiTrialAlgorithmSpec.MetricSpec
type NasJobSpec_MultiTrialAlgorithmSpec_MetricSpec struct {
	// Required. The ID of the metric. Must not contain whitespaces.
	// +kcc:proto:field=google.cloud.aiplatform.v1.NasJobSpec.MultiTrialAlgorithmSpec.MetricSpec.metric_id
	MetricID *string `json:"metricID,omitempty"`

	// Required. The optimization goal of the metric.
	// +kcc:proto:field=google.cloud.aiplatform.v1.NasJobSpec.MultiTrialAlgorithmSpec.MetricSpec.goal
	Goal *string `json:"goal,omitempty"`
}

// +kcc:proto=google.cloud.aiplatform.v1.NasJobSpec.MultiTrialAlgorithmSpec.SearchTrialSpec
type NasJobSpec_MultiTrialAlgorithmSpec_SearchTrialSpec struct {
	// Required. The spec of a search trial job. The same spec applies to
	//  all search trials.
	// +kcc:proto:field=google.cloud.aiplatform.v1.NasJobSpec.MultiTrialAlgorithmSpec.SearchTrialSpec.search_trial_job_spec
	SearchTrialJobSpec *CustomJobSpec `json:"searchTrialJobSpec,omitempty"`

	// Required. The maximum number of Neural Architecture Search (NAS) trials
	//  to run.
	// +kcc:proto:field=google.cloud.aiplatform.v1.NasJobSpec.MultiTrialAlgorithmSpec.SearchTrialSpec.max_trial_count
	MaxTrialCount *int32 `json:"maxTrialCount,omitempty"`

	// Required. The maximum number of trials to run in parallel.
	// +kcc:proto:field=google.cloud.aiplatform.v1.NasJobSpec.MultiTrialAlgorithmSpec.SearchTrialSpec.max_parallel_trial_count
	MaxParallelTrialCount *int32 `json:"maxParallelTrialCount,omitempty"`

	// The number of failed trials that need to be seen before failing
	//  the NasJob.
	//
	//  If set to 0, Vertex AI decides how many trials must fail
	//  before the whole job fails.
	// +kcc:proto:field=google.cloud.aiplatform.v1.NasJobSpec.MultiTrialAlgorithmSpec.SearchTrialSpec.max_failed_trial_count
	MaxFailedTrialCount *int32 `json:"maxFailedTrialCount,omitempty"`
}

// +kcc:proto=google.cloud.aiplatform.v1.NasJobSpec.MultiTrialAlgorithmSpec.TrainTrialSpec
type NasJobSpec_MultiTrialAlgorithmSpec_TrainTrialSpec struct {
	// Required. The spec of a train trial job. The same spec applies to
	//  all train trials.
	// +kcc:proto:field=google.cloud.aiplatform.v1.NasJobSpec.MultiTrialAlgorithmSpec.TrainTrialSpec.train_trial_job_spec
	TrainTrialJobSpec *CustomJobSpec `json:"trainTrialJobSpec,omitempty"`

	// Required. The maximum number of trials to run in parallel.
	// +kcc:proto:field=google.cloud.aiplatform.v1.NasJobSpec.MultiTrialAlgorithmSpec.TrainTrialSpec.max_parallel_trial_count
	MaxParallelTrialCount *int32 `json:"maxParallelTrialCount,omitempty"`

	// Required. Frequency of search trials to start train stage. Top N
	//  [TrainTrialSpec.max_parallel_trial_count]
	//  search trials will be trained for every M
	//  [TrainTrialSpec.frequency] trials searched.
	// +kcc:proto:field=google.cloud.aiplatform.v1.NasJobSpec.MultiTrialAlgorithmSpec.TrainTrialSpec.frequency
	Frequency *int32 `json:"frequency,omitempty"`
}

// +kcc:proto=google.cloud.aiplatform.v1.NasTrial
type NasTrial struct {
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

// +kcc:proto=google.cloud.aiplatform.v1.NasJob
type NasJobObservedState struct {
	// Output only. Resource name of the NasJob.
	// +kcc:proto:field=google.cloud.aiplatform.v1.NasJob.name
	Name *string `json:"name,omitempty"`

	// Output only. Output of the NasJob.
	// +kcc:proto:field=google.cloud.aiplatform.v1.NasJob.nas_job_output
	NasJobOutput *NasJobOutput `json:"nasJobOutput,omitempty"`

	// Output only. The detailed state of the job.
	// +kcc:proto:field=google.cloud.aiplatform.v1.NasJob.state
	State *string `json:"state,omitempty"`

	// Output only. Time when the NasJob was created.
	// +kcc:proto:field=google.cloud.aiplatform.v1.NasJob.create_time
	CreateTime *string `json:"createTime,omitempty"`

	// Output only. Time when the NasJob for the first time entered the
	//  `JOB_STATE_RUNNING` state.
	// +kcc:proto:field=google.cloud.aiplatform.v1.NasJob.start_time
	StartTime *string `json:"startTime,omitempty"`

	// Output only. Time when the NasJob entered any of the following states:
	//  `JOB_STATE_SUCCEEDED`, `JOB_STATE_FAILED`, `JOB_STATE_CANCELLED`.
	// +kcc:proto:field=google.cloud.aiplatform.v1.NasJob.end_time
	EndTime *string `json:"endTime,omitempty"`

	// Output only. Time when the NasJob was most recently updated.
	// +kcc:proto:field=google.cloud.aiplatform.v1.NasJob.update_time
	UpdateTime *string `json:"updateTime,omitempty"`

	// Output only. Only populated when job's state is JOB_STATE_FAILED or
	//  JOB_STATE_CANCELLED.
	// +kcc:proto:field=google.cloud.aiplatform.v1.NasJob.error
	Error *Status `json:"error,omitempty"`

	// Output only. Reserved for future use.
	// +kcc:proto:field=google.cloud.aiplatform.v1.NasJob.satisfies_pzs
	SatisfiesPzs *bool `json:"satisfiesPzs,omitempty"`

	// Output only. Reserved for future use.
	// +kcc:proto:field=google.cloud.aiplatform.v1.NasJob.satisfies_pzi
	SatisfiesPzi *bool `json:"satisfiesPzi,omitempty"`
}

// +kcc:proto=google.cloud.aiplatform.v1.NasJobOutput
type NasJobOutputObservedState struct {
	// Output only. The output of this multi-trial Neural Architecture Search
	//  (NAS) job.
	// +kcc:proto:field=google.cloud.aiplatform.v1.NasJobOutput.multi_trial_job_output
	MultiTrialJobOutput *NasJobOutput_MultiTrialJobOutput `json:"multiTrialJobOutput,omitempty"`
}

// +kcc:proto=google.cloud.aiplatform.v1.NasJobOutput.MultiTrialJobOutput
type NasJobOutput_MultiTrialJobOutputObservedState struct {
	// Output only. List of NasTrials that were started as part of search stage.
	// +kcc:proto:field=google.cloud.aiplatform.v1.NasJobOutput.MultiTrialJobOutput.search_trials
	SearchTrials []NasTrial `json:"searchTrials,omitempty"`

	// Output only. List of NasTrials that were started as part of train stage.
	// +kcc:proto:field=google.cloud.aiplatform.v1.NasJobOutput.MultiTrialJobOutput.train_trials
	TrainTrials []NasTrial `json:"trainTrials,omitempty"`
}

// +kcc:proto=google.cloud.aiplatform.v1.NasTrial
type NasTrialObservedState struct {
	// Output only. The identifier of the NasTrial assigned by the service.
	// +kcc:proto:field=google.cloud.aiplatform.v1.NasTrial.id
	ID *string `json:"id,omitempty"`

	// Output only. The detailed state of the NasTrial.
	// +kcc:proto:field=google.cloud.aiplatform.v1.NasTrial.state
	State *string `json:"state,omitempty"`

	// Output only. The final measurement containing the objective value.
	// +kcc:proto:field=google.cloud.aiplatform.v1.NasTrial.final_measurement
	FinalMeasurement *Measurement `json:"finalMeasurement,omitempty"`

	// Output only. Time when the NasTrial was started.
	// +kcc:proto:field=google.cloud.aiplatform.v1.NasTrial.start_time
	StartTime *string `json:"startTime,omitempty"`

	// Output only. Time when the NasTrial's status changed to `SUCCEEDED` or
	//  `INFEASIBLE`.
	// +kcc:proto:field=google.cloud.aiplatform.v1.NasTrial.end_time
	EndTime *string `json:"endTime,omitempty"`
}
