// Copyright 2024 Google LLC
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

package v1beta1

// +kcc:proto=google.dataflow.v1beta3.AutoscalingEvent
type AutoscalingEvent struct {
	// The current number of workers the job has.
	CurrentNumWorkers *int64 `json:"currentNumWorkers,omitempty"`

	// The target number of workers the worker pool wants to resize to use.
	TargetNumWorkers *int64 `json:"targetNumWorkers,omitempty"`

	// The type of autoscaling event to report.
	EventType *string `json:"eventType,omitempty"`

	// A message describing why the system decided to adjust the current
	//  number of workers, why it failed, or why the system decided to
	//  not make any changes to the number of workers.
	Description *StructuredMessage `json:"description,omitempty"`

	// The time this event was emitted to indicate a new target or current
	//  num_workers value.
	Time *string `json:"time,omitempty"`

	// A short and friendly name for the worker pool this event refers to.
	WorkerPool *string `json:"workerPool,omitempty"`
}

// +kcc:proto=google.dataflow.v1beta3.AutoscalingSettings
type AutoscalingSettings struct {
	// The algorithm to use for autoscaling.
	Algorithm *string `json:"algorithm,omitempty"`

	// The maximum number of workers to cap scaling at.
	MaxNumWorkers *int32 `json:"maxNumWorkers,omitempty"`
}

// +kcc:proto=google.dataflow.v1beta3.BigQueryIODetails
type BigQueryIODetails struct {
	// Table accessed in the connection.
	Table *string `json:"table,omitempty"`

	// Dataset accessed in the connection.
	Dataset *string `json:"dataset,omitempty"`

	// Project accessed in the connection.
	ProjectID *string `json:"projectID,omitempty"`

	// Query used to access data in the connection.
	Query *string `json:"query,omitempty"`
}

// +kcc:proto=google.dataflow.v1beta3.BigTableIODetails
type BigTableIODetails struct {
	// ProjectId accessed in the connection.
	ProjectID *string `json:"projectID,omitempty"`

	// InstanceId accessed in the connection.
	InstanceID *string `json:"instanceID,omitempty"`

	// TableId accessed in the connection.
	TableID *string `json:"tableID,omitempty"`
}

// +kcc:proto=google.dataflow.v1beta3.ComputationTopology
type ComputationTopology struct {
	// The system stage name.
	SystemStageName *string `json:"systemStageName,omitempty"`

	// The ID of the computation.
	ComputationID *string `json:"computationID,omitempty"`

	// The key ranges processed by the computation.
	KeyRanges []KeyRangeLocation `json:"keyRanges,omitempty"`

	// The inputs to the computation.
	Inputs []StreamLocation `json:"inputs,omitempty"`

	// The outputs from the computation.
	Outputs []StreamLocation `json:"outputs,omitempty"`

	// The state family values.
	StateFamilies []StateFamilyConfig `json:"stateFamilies,omitempty"`
}

// +kcc:proto=google.dataflow.v1beta3.ContainerSpec
type ContainerSpec struct {
	// Name of the docker container image. E.g., gcr.io/project/some-image
	Image *string `json:"image,omitempty"`

	// Metadata describing a template including description and validation rules.
	Metadata *TemplateMetadata `json:"metadata,omitempty"`

	// Required. SDK info of the Flex Template.
	SdkInfo *SDKInfo `json:"sdkInfo,omitempty"`

	// Default runtime environment for the job.
	DefaultEnvironment *FlexTemplateRuntimeEnvironment `json:"defaultEnvironment,omitempty"`
}

// +kcc:proto=google.dataflow.v1beta3.CustomSourceLocation
type CustomSourceLocation struct {
	// Whether this source is stateful.
	Stateful *bool `json:"stateful,omitempty"`
}

// +kcc:proto=google.dataflow.v1beta3.DataDiskAssignment
type DataDiskAssignment struct {
	// VM instance name the data disks mounted to, for example
	//  "myproject-1014-104817-4c2-harness-0".
	VmInstance *string `json:"vmInstance,omitempty"`

	// Mounted data disks. The order is important a data disk's 0-based index in
	//  this list defines which persistent directory the disk is mounted to, for
	//  example the list of { "myproject-1014-104817-4c2-harness-0-disk-0" },
	//  { "myproject-1014-104817-4c2-harness-0-disk-1" }.
	DataDisks []string `json:"dataDisks,omitempty"`
}

// +kcc:proto=google.dataflow.v1beta3.DatastoreIODetails
type DatastoreIODetails struct {
	// Namespace used in the connection.
	Namespace *string `json:"namespace,omitempty"`

	// ProjectId accessed in the connection.
	ProjectID *string `json:"projectID,omitempty"`
}

// +kcc:proto=google.dataflow.v1beta3.DebugOptions
type DebugOptions struct {
	// When true, enables the logging of the literal hot key to the user's Cloud
	//  Logging.
	EnableHotKeyLogging *bool `json:"enableHotKeyLogging,omitempty"`
}

// +kcc:proto=google.dataflow.v1beta3.Disk
type Disk struct {
	// Size of disk in GB.  If zero or unspecified, the service will
	//  attempt to choose a reasonable default.
	SizeGb *int32 `json:"sizeGb,omitempty"`

	// Disk storage type, as defined by Google Compute Engine.  This
	//  must be a disk type appropriate to the project and zone in which
	//  the workers will run.  If unknown or unspecified, the service
	//  will attempt to choose a reasonable default.
	//
	//  For example, the standard persistent disk type is a resource name
	//  typically ending in "pd-standard".  If SSD persistent disks are
	//  available, the resource name typically ends with "pd-ssd".  The
	//  actual valid values are defined the Google Compute Engine API,
	//  not by the Cloud Dataflow API; consult the Google Compute Engine
	//  documentation for more information about determining the set of
	//  available disk types for a particular project and zone.
	//
	//  Google Compute Engine Disk types are local to a particular
	//  project in a particular zone, and so the resource name will
	//  typically look something like this:
	//
	//  compute.googleapis.com/projects/project-id/zones/zone/diskTypes/pd-standard
	DiskType *string `json:"diskType,omitempty"`

	// Directory in a VM where disk is mounted.
	MountPoint *string `json:"mountPoint,omitempty"`
}

// +kcc:proto=google.dataflow.v1beta3.DisplayData
type DisplayData struct {
	// The key identifying the display data.
	//  This is intended to be used as a label for the display data
	//  when viewed in a dax monitoring system.
	Key *string `json:"key,omitempty"`

	// The namespace for the key. This is usually a class name or programming
	//  language namespace (i.e. python module) which defines the display data.
	//  This allows a dax monitoring system to specially handle the data
	//  and perform custom rendering.
	Namespace *string `json:"namespace,omitempty"`

	// Contains value if the data is of string type.
	StrValue *string `json:"strValue,omitempty"`

	// Contains value if the data is of int64 type.
	Int64Value *int64 `json:"int64Value,omitempty"`

	// Contains value if the data is of float type.
	FloatValue *float32 `json:"floatValue,omitempty"`

	// Contains value if the data is of java class type.
	JavaClassValue *string `json:"javaClassValue,omitempty"`

	// Contains value if the data is of timestamp type.
	TimestampValue *string `json:"timestampValue,omitempty"`

	// Contains value if the data is of duration type.
	DurationValue *string `json:"durationValue,omitempty"`

	// Contains value if the data is of a boolean type.
	BoolValue *bool `json:"boolValue,omitempty"`

	// A possible additional shorter value to display.
	//  For example a java_class_name_value of com.mypackage.MyDoFn
	//  will be stored with MyDoFn as the short_str_value and
	//  com.mypackage.MyDoFn as the java_class_name value.
	//  short_str_value can be displayed and java_class_name_value
	//  will be displayed as a tooltip.
	ShortStrValue *string `json:"shortStrValue,omitempty"`

	// An optional full URL.
	URL *string `json:"url,omitempty"`

	// An optional label to display in a dax UI for the element.
	Label *string `json:"label,omitempty"`
}

// +kcc:proto=google.dataflow.v1beta3.DynamicTemplateLaunchParams
type DynamicTemplateLaunchParams struct {
	// Path to dynamic template spec file on Cloud Storage.
	//  The file must be a Json serialized DynamicTemplateFieSpec object.
	GcsPath *string `json:"gcsPath,omitempty"`

	// Cloud Storage path for staging dependencies.
	//  Must be a valid Cloud Storage URL, beginning with `gs://`.
	StagingLocation *string `json:"stagingLocation,omitempty"`
}

// +kcc:proto=google.dataflow.v1beta3.Environment
type Environment struct {
	// The prefix of the resources the system should use for temporary
	//  storage.  The system will append the suffix "/temp-{JOBNAME} to
	//  this resource prefix, where {JOBNAME} is the value of the
	//  job_name field.  The resulting bucket and object prefix is used
	//  as the prefix of the resources used to store temporary data
	//  needed during the job execution.  NOTE: This will override the
	//  value in taskrunner_settings.
	//  The supported resource type is:
	//
	//  Google Cloud Storage:
	//
	//    storage.googleapis.com/{bucket}/{object}
	//    bucket.storage.googleapis.com/{object}
	TempStoragePrefix *string `json:"tempStoragePrefix,omitempty"`

	// The type of cluster manager API to use.  If unknown or
	//  unspecified, the service will attempt to choose a reasonable
	//  default.  This should be in the form of the API service name,
	//  e.g. "compute.googleapis.com".
	ClusterManagerApiService *string `json:"clusterManagerApiService,omitempty"`

	// The list of experiments to enable. This field should be used for SDK
	//  related experiments and not for service related experiments. The proper
	//  field for service related experiments is service_options.
	Experiments []string `json:"experiments,omitempty"`

	// The list of service options to enable. This field should be used for
	//  service related experiments only. These experiments, when graduating to GA,
	//  should be replaced by dedicated fields or become default (i.e. always on).
	ServiceOptions []string `json:"serviceOptions,omitempty"`

	// If set, contains the Cloud KMS key identifier used to encrypt data
	//  at rest, AKA a Customer Managed Encryption Key (CMEK).
	//
	//  Format:
	//    projects/PROJECT_ID/locations/LOCATION/keyRings/KEY_RING/cryptoKeys/KEY
	ServiceKmsKeyName *string `json:"serviceKmsKeyName,omitempty"`

	// The worker pools. At least one "harness" worker pool must be
	//  specified in order for the job to have workers.
	WorkerPools []WorkerPool `json:"workerPools,omitempty"`

	// A description of the process that generated the request.
	UserAgent *google_protobuf_Struct `json:"userAgent,omitempty"`

	// A structure describing which components and their versions of the service
	//  are required in order to run the job.
	Version *google_protobuf_Struct `json:"version,omitempty"`

	// The dataset for the current project where various workflow
	//  related tables are stored.
	//
	//  The supported resource type is:
	//
	//  Google BigQuery:
	//    bigquery.googleapis.com/{dataset}
	Dataset *string `json:"dataset,omitempty"`

	// The Cloud Dataflow SDK pipeline options specified by the user. These
	//  options are passed through the service and are used to recreate the
	//  SDK pipeline options on the worker in a language agnostic and platform
	//  independent way.
	SdkPipelineOptions *google_protobuf_Struct `json:"sdkPipelineOptions,omitempty"`

	// Experimental settings.
	InternalExperiments *google_protobuf_Any `json:"internalExperiments,omitempty"`

	// Identity to run virtual machines as. Defaults to the default account.
	ServiceAccountEmail *string `json:"serviceAccountEmail,omitempty"`

	// Which Flexible Resource Scheduling mode to run in.
	FlexResourceSchedulingGoal *string `json:"flexResourceSchedulingGoal,omitempty"`

	// The Compute Engine region
	//  (https://cloud.google.com/compute/docs/regions-zones/regions-zones) in
	//  which worker processing should occur, e.g. "us-west1". Mutually exclusive
	//  with worker_zone. If neither worker_region nor worker_zone is specified,
	//  default to the control plane's region.
	WorkerRegion *string `json:"workerRegion,omitempty"`

	// The Compute Engine zone
	//  (https://cloud.google.com/compute/docs/regions-zones/regions-zones) in
	//  which worker processing should occur, e.g. "us-west1-a". Mutually exclusive
	//  with worker_region. If neither worker_region nor worker_zone is specified,
	//  a zone in the control plane's region is chosen based on available capacity.
	WorkerZone *string `json:"workerZone,omitempty"`

	// Output only. The shuffle mode used for the job.
	ShuffleMode *string `json:"shuffleMode,omitempty"`

	// Any debugging options to be supplied to the job.
	DebugOptions *DebugOptions `json:"debugOptions,omitempty"`
}

// +kcc:proto=google.dataflow.v1beta3.ExecutionStageState
type ExecutionStageState struct {
	// The name of the execution stage.
	ExecutionStageName *string `json:"executionStageName,omitempty"`

	// Executions stage states allow the same set of values as JobState.
	ExecutionStageState *string `json:"executionStageState,omitempty"`

	// The time at which the stage transitioned to this state.
	CurrentStateTime *string `json:"currentStateTime,omitempty"`
}

// +kcc:proto=google.dataflow.v1beta3.ExecutionStageSummary
type ExecutionStageSummary struct {
	// Dataflow service generated name for this stage.
	Name *string `json:"name,omitempty"`

	// Dataflow service generated id for this stage.
	ID *string `json:"id,omitempty"`

	// Type of transform this stage is executing.
	Kind *string `json:"kind,omitempty"`

	// Input sources for this stage.
	InputSource []ExecutionStageSummary_StageSource `json:"inputSource,omitempty"`

	// Output sources for this stage.
	OutputSource []ExecutionStageSummary_StageSource `json:"outputSource,omitempty"`

	// Other stages that must complete before this stage can run.
	PrerequisiteStage []string `json:"prerequisiteStage,omitempty"`

	// Transforms that comprise this execution stage.
	ComponentTransform []ExecutionStageSummary_ComponentTransform `json:"componentTransform,omitempty"`

	// Collections produced and consumed by component transforms of this stage.
	ComponentSource []ExecutionStageSummary_ComponentSource `json:"componentSource,omitempty"`
}

// +kcc:proto=google.dataflow.v1beta3.ExecutionStageSummary.ComponentSource
type ExecutionStageSummary_ComponentSource struct {
	// Human-readable name for this transform; may be user or system generated.
	UserName *string `json:"userName,omitempty"`

	// Dataflow service generated name for this source.
	Name *string `json:"name,omitempty"`

	// User name for the original user transform or collection with which this
	//  source is most closely associated.
	OriginalTransformOrCollection *string `json:"originalTransformOrCollection,omitempty"`
}

// +kcc:proto=google.dataflow.v1beta3.ExecutionStageSummary.ComponentTransform
type ExecutionStageSummary_ComponentTransform struct {
	// Human-readable name for this transform; may be user or system generated.
	UserName *string `json:"userName,omitempty"`

	// Dataflow service generated name for this source.
	Name *string `json:"name,omitempty"`

	// User name for the original user transform with which this transform is
	//  most closely associated.
	OriginalTransform *string `json:"originalTransform,omitempty"`
}

// +kcc:proto=google.dataflow.v1beta3.ExecutionStageSummary.StageSource
type ExecutionStageSummary_StageSource struct {
	// Human-readable name for this source; may be user or system generated.
	UserName *string `json:"userName,omitempty"`

	// Dataflow service generated name for this source.
	Name *string `json:"name,omitempty"`

	// User name for the original user transform or collection with which this
	//  source is most closely associated.
	OriginalTransformOrCollection *string `json:"originalTransformOrCollection,omitempty"`

	// Size of the source, if measurable.
	SizeBytes *int64 `json:"sizeBytes,omitempty"`
}

// +kcc:proto=google.dataflow.v1beta3.FailedLocation
type FailedLocation struct {
	// The name of the [regional endpoint]
	//  (https://cloud.google.com/dataflow/docs/concepts/regional-endpoints) that
	//  failed to respond.
	Name *string `json:"name,omitempty"`
}

// +kcc:proto=google.dataflow.v1beta3.FileIODetails
type FileIODetails struct {
	// File Pattern used to access files by the connector.
	FilePattern *string `json:"filePattern,omitempty"`
}

// +kcc:proto=google.dataflow.v1beta3.InvalidTemplateParameters
type InvalidTemplateParameters struct {
	// Describes all parameter violations in a template request.
	ParameterViolations []InvalidTemplateParameters_ParameterViolation `json:"parameterViolations,omitempty"`
}

// +kcc:proto=google.dataflow.v1beta3.InvalidTemplateParameters.ParameterViolation
type InvalidTemplateParameters_ParameterViolation struct {
	// The parameter that failed to validate.
	Parameter *string `json:"parameter,omitempty"`

	// A description of why the parameter failed to validate.
	Description *string `json:"description,omitempty"`
}

// +kcc:proto=google.dataflow.v1beta3.Job
type Job struct {
	// The unique ID of this job.
	//
	//  This field is set by the Cloud Dataflow service when the Job is
	//  created, and is immutable for the life of the job.
	ID *string `json:"id,omitempty"`

	// The ID of the Cloud Platform project that the job belongs to.
	ProjectID *string `json:"projectID,omitempty"`

	// The user-specified Cloud Dataflow job name.
	//
	//  Only one Job with a given name may exist in a project at any
	//  given time. If a caller attempts to create a Job with the same
	//  name as an already-existing Job, the attempt returns the
	//  existing Job.
	//
	//  The name must match the regular expression
	//  `[a-z]([-a-z0-9]{0,1022}[a-z0-9])?`
	Name *string `json:"name,omitempty"`

	// The type of Cloud Dataflow job.
	Type *string `json:"type,omitempty"`

	// The environment for the job.
	Environment *Environment `json:"environment,omitempty"`

	// Exactly one of step or steps_location should be specified.
	//
	//  The top-level steps that constitute the entire job. Only retrieved with
	//  JOB_VIEW_ALL.
	Steps []Step `json:"steps,omitempty"`

	// The Cloud Storage location where the steps are stored.
	StepsLocation *string `json:"stepsLocation,omitempty"`

	// The current state of the job.
	//
	//  Jobs are created in the `JOB_STATE_STOPPED` state unless otherwise
	//  specified.
	//
	//  A job in the `JOB_STATE_RUNNING` state may asynchronously enter a
	//  terminal state. After a job has reached a terminal state, no
	//  further state updates may be made.
	//
	//  This field may be mutated by the Cloud Dataflow service;
	//  callers cannot mutate it.
	CurrentState *string `json:"currentState,omitempty"`

	// The timestamp associated with the current state.
	CurrentStateTime *string `json:"currentStateTime,omitempty"`

	// The job's requested state.
	//
	//  `UpdateJob` may be used to switch between the `JOB_STATE_STOPPED` and
	//  `JOB_STATE_RUNNING` states, by setting requested_state.  `UpdateJob` may
	//  also be used to directly set a job's requested state to
	//  `JOB_STATE_CANCELLED` or `JOB_STATE_DONE`, irrevocably terminating the
	//  job if it has not already reached a terminal state.
	RequestedState *string `json:"requestedState,omitempty"`

	// Deprecated.
	ExecutionInfo *JobExecutionInfo `json:"executionInfo,omitempty"`

	// The timestamp when the job was initially created. Immutable and set by the
	//  Cloud Dataflow service.
	CreateTime *string `json:"createTime,omitempty"`

	// If this job is an update of an existing job, this field is the job ID
	//  of the job it replaced.
	//
	//  When sending a `CreateJobRequest`, you can update a job by specifying it
	//  here. The job named here is stopped, and its intermediate state is
	//  transferred to this job.
	ReplaceJobID *string `json:"replaceJobID,omitempty"`

	// The map of transform name prefixes of the job to be replaced to the
	//  corresponding name prefixes of the new job.
	TransformNameMapping map[string]string `json:"transformNameMapping,omitempty"`

	// The client's unique identifier of the job, re-used across retried attempts.
	//  If this field is set, the service will ensure its uniqueness.
	//  The request to create a job will fail if the service has knowledge of a
	//  previously submitted job with the same client's ID and job name.
	//  The caller may use this field to ensure idempotence of job
	//  creation across retried attempts to create a job.
	//  By default, the field is empty and, in that case, the service ignores it.
	ClientRequestID *string `json:"clientRequestID,omitempty"`

	// If another job is an update of this job (and thus, this job is in
	//  `JOB_STATE_UPDATED`), this field contains the ID of that job.
	ReplacedByJobID *string `json:"replacedByJobID,omitempty"`

	// A set of files the system should be aware of that are used
	//  for temporary storage. These temporary files will be
	//  removed on job completion.
	//  No duplicates are allowed.
	//  No file patterns are supported.
	//
	//  The supported files are:
	//
	//  Google Cloud Storage:
	//
	//     storage.googleapis.com/{bucket}/{object}
	//     bucket.storage.googleapis.com/{object}
	TempFiles []string `json:"tempFiles,omitempty"`

	// User-defined labels for this job.
	//
	//  The labels map can contain no more than 64 entries.  Entries of the labels
	//  map are UTF8 strings that comply with the following restrictions:
	//
	//  * Keys must conform to regexp:  [\p{Ll}\p{Lo}][\p{Ll}\p{Lo}\p{N}_-]{0,62}
	//  * Values must conform to regexp:  [\p{Ll}\p{Lo}\p{N}_-]{0,63}
	//  * Both keys and values are additionally constrained to be <= 128 bytes in
	//  size.
	Labels map[string]string `json:"labels,omitempty"`

	// The [regional endpoint]
	//  (https://cloud.google.com/dataflow/docs/concepts/regional-endpoints) that
	//  contains this job.
	Location *string `json:"location,omitempty"`

	// Preliminary field: The format of this data may change at any time.
	//  A description of the user pipeline and stages through which it is executed.
	//  Created by Cloud Dataflow service.  Only retrieved with
	//  JOB_VIEW_DESCRIPTION or JOB_VIEW_ALL.
	PipelineDescription *PipelineDescription `json:"pipelineDescription,omitempty"`

	// This field may be mutated by the Cloud Dataflow service;
	//  callers cannot mutate it.
	StageStates []ExecutionStageState `json:"stageStates,omitempty"`

	// This field is populated by the Dataflow service to support filtering jobs
	//  by the metadata values provided here. Populated for ListJobs and all GetJob
	//  views SUMMARY and higher.
	JobMetadata *JobMetadata `json:"jobMetadata,omitempty"`

	// The timestamp when the job was started (transitioned to JOB_STATE_PENDING).
	//  Flexible resource scheduling jobs are started with some delay after job
	//  creation, so start_time is unset before start and is updated when the
	//  job is started by the Cloud Dataflow service. For other jobs, start_time
	//  always equals create_time and is immutable and set by the Cloud Dataflow
	//  service.
	StartTime *string `json:"startTime,omitempty"`

	// If this is specified, the job's initial state is populated from the given
	//  snapshot.
	CreatedFromSnapshotID *string `json:"createdFromSnapshotID,omitempty"`

	// Reserved for future use. This field is set only in responses from the
	//  server; it is ignored if it is set in any requests.
	SatisfiesPzs *bool `json:"satisfiesPzs,omitempty"`
}

// +kcc:proto=google.dataflow.v1beta3.JobExecutionDetails
type JobExecutionDetails struct {
	// The stages of the job execution.
	Stages []StageSummary `json:"stages,omitempty"`

	// If present, this response does not contain all requested tasks.  To obtain
	//  the next page of results, repeat the request with page_token set to this
	//  value.
	NextPageToken *string `json:"nextPageToken,omitempty"`
}

// +kcc:proto=google.dataflow.v1beta3.JobExecutionInfo
type JobExecutionInfo struct {

	// TODO: map type string message for stages

}

// +kcc:proto=google.dataflow.v1beta3.JobExecutionStageInfo
type JobExecutionStageInfo struct {
	// The steps associated with the execution stage.
	//  Note that stages may have several steps, and that a given step
	//  might be run by more than one stage.
	StepName []string `json:"stepName,omitempty"`
}

// +kcc:proto=google.dataflow.v1beta3.JobMessage
type JobMessage struct {
	// Deprecated.
	ID *string `json:"id,omitempty"`

	// The timestamp of the message.
	Time *string `json:"time,omitempty"`

	// The text of the message.
	MessageText *string `json:"messageText,omitempty"`

	// Importance level of the message.
	MessageImportance *string `json:"messageImportance,omitempty"`
}

// +kcc:proto=google.dataflow.v1beta3.JobMetrics
type JobMetrics struct {
	// Timestamp as of which metric values are current.
	MetricTime *string `json:"metricTime,omitempty"`

	// All metrics for this job.
	Metrics []MetricUpdate `json:"metrics,omitempty"`
}

// +kcc:proto=google.dataflow.v1beta3.KeyRangeDataDiskAssignment
type KeyRangeDataDiskAssignment struct {
	// The start (inclusive) of the key range.
	Start *string `json:"start,omitempty"`

	// The end (exclusive) of the key range.
	End *string `json:"end,omitempty"`

	// The name of the data disk where data for this range is stored.
	//  This name is local to the Google Cloud Platform project and uniquely
	//  identifies the disk within that project, for example
	//  "myproject-1014-104817-4c2-harness-0-disk-1".
	DataDisk *string `json:"dataDisk,omitempty"`
}

// +kcc:proto=google.dataflow.v1beta3.KeyRangeLocation
type KeyRangeLocation struct {
	// The start (inclusive) of the key range.
	Start *string `json:"start,omitempty"`

	// The end (exclusive) of the key range.
	End *string `json:"end,omitempty"`

	// The physical location of this range assignment to be used for
	//  streaming computation cross-worker message delivery.
	DeliveryEndpoint *string `json:"deliveryEndpoint,omitempty"`

	// The name of the data disk where data for this range is stored.
	//  This name is local to the Google Cloud Platform project and uniquely
	//  identifies the disk within that project, for example
	//  "myproject-1014-104817-4c2-harness-0-disk-1".
	DataDisk *string `json:"dataDisk,omitempty"`

	// DEPRECATED. The location of the persistent state for this range, as a
	//  persistent directory in the worker local filesystem.
	DeprecatedPersistentDirectory *string `json:"deprecatedPersistentDirectory,omitempty"`
}

// +kcc:proto=google.dataflow.v1beta3.LaunchFlexTemplateParameter
type LaunchFlexTemplateParameter struct {
	// Required. The job name to use for the created job. For update job request,
	//  job name should be same as the existing running job.
	JobName *string `json:"jobName,omitempty"`

	// Spec about the container image to launch.
	ContainerSpec *ContainerSpec `json:"containerSpec,omitempty"`

	// Cloud Storage path to a file with json serialized ContainerSpec as
	//  content.
	ContainerSpecGcsPath *string `json:"containerSpecGcsPath,omitempty"`

	// The parameters for FlexTemplate.
	//  Ex. {"num_workers":"5"}
	Parameters map[string]string `json:"parameters,omitempty"`

	// Launch options for this flex template job. This is a common set of options
	//  across languages and templates. This should not be used to pass job
	//  parameters.
	LaunchOptions map[string]string `json:"launchOptions,omitempty"`

	// The runtime environment for the FlexTemplate job
	Environment *FlexTemplateRuntimeEnvironment `json:"environment,omitempty"`

	// Set this to true if you are sending a request to update a running
	//  streaming job. When set, the job name should be the same as the
	//  running job.
	Update *bool `json:"update,omitempty"`

	// Use this to pass transform_name_mappings for streaming update jobs.
	//  Ex:{"oldTransformName":"newTransformName",...}'
	TransformNameMappings map[string]string `json:"transformNameMappings,omitempty"`
}

// +kcc:proto=google.dataflow.v1beta3.LaunchTemplateParameters
type LaunchTemplateParameters struct {
	// Required. The job name to use for the created job.
	JobName *string `json:"jobName,omitempty"`

	// The runtime parameters to pass to the job.
	Parameters map[string]string `json:"parameters,omitempty"`

	// The runtime environment for the job.
	Environment *RuntimeEnvironment `json:"environment,omitempty"`

	// If set, replace the existing pipeline with the name specified by jobName
	//  with this pipeline, preserving state.
	Update *bool `json:"update,omitempty"`

	// Only applicable when updating a pipeline. Map of transform name prefixes of
	//  the job to be replaced to the corresponding name prefixes of the new job.
	TransformNameMapping map[string]string `json:"transformNameMapping,omitempty"`
}

// +kcc:proto=google.dataflow.v1beta3.MetricStructuredName
type MetricStructuredName struct {
	// Origin (namespace) of metric name. May be blank for user-define metrics;
	//  will be "dataflow" for metrics defined by the Dataflow service or SDK.
	Origin *string `json:"origin,omitempty"`

	// Worker-defined metric name.
	Name *string `json:"name,omitempty"`

	// Zero or more labeled fields which identify the part of the job this
	//  metric is associated with, such as the name of a step or collection.
	//
	//  For example, built-in counters associated with steps will have
	//  context['step'] = <step-name>. Counters associated with PCollections
	//  in the SDK will have context['pcollection'] = <pcollection-name>.
	Context map[string]string `json:"context,omitempty"`
}

// +kcc:proto=google.dataflow.v1beta3.MetricUpdate
type MetricUpdate struct {
	// Name of the metric.
	Name *MetricStructuredName `json:"name,omitempty"`

	// Metric aggregation kind.  The possible metric aggregation kinds are
	//  "Sum", "Max", "Min", "Mean", "Set", "And", "Or", and "Distribution".
	//  The specified aggregation kind is case-insensitive.
	//
	//  If omitted, this is not an aggregated value but instead
	//  a single metric sample value.
	Kind *string `json:"kind,omitempty"`

	// True if this metric is reported as the total cumulative aggregate
	//  value accumulated since the worker started working on this WorkItem.
	//  By default this is false, indicating that this metric is reported
	//  as a delta that is not associated with any WorkItem.
	Cumulative *bool `json:"cumulative,omitempty"`

	// Worker-computed aggregate value for aggregation kinds "Sum", "Max", "Min",
	//  "And", and "Or".  The possible value types are Long, Double, and Boolean.
	Scalar *google_protobuf_Value `json:"scalar,omitempty"`

	// Worker-computed aggregate value for the "Mean" aggregation kind.
	//  This holds the sum of the aggregated values and is used in combination
	//  with mean_count below to obtain the actual mean aggregate value.
	//  The only possible value types are Long and Double.
	MeanSum *google_protobuf_Value `json:"meanSum,omitempty"`

	// Worker-computed aggregate value for the "Mean" aggregation kind.
	//  This holds the count of the aggregated values and is used in combination
	//  with mean_sum above to obtain the actual mean aggregate value.
	//  The only possible value type is Long.
	MeanCount *google_protobuf_Value `json:"meanCount,omitempty"`

	// Worker-computed aggregate value for the "Set" aggregation kind.  The only
	//  possible value type is a list of Values whose type can be Long, Double,
	//  or String, according to the metric's type.  All Values in the list must
	//  be of the same type.
	Set *google_protobuf_Value `json:"set,omitempty"`

	// A struct value describing properties of a distribution of numeric values.
	Distribution *google_protobuf_Value `json:"distribution,omitempty"`

	// A struct value describing properties of a Gauge.
	//  Metrics of gauge type show the value of a metric across time, and is
	//  aggregated based on the newest value.
	Gauge *google_protobuf_Value `json:"gauge,omitempty"`

	// Worker-computed aggregate value for internal use by the Dataflow
	//  service.
	Internal *google_protobuf_Value `json:"internal,omitempty"`

	// Timestamp associated with the metric value. Optional when workers are
	//  reporting work progress; it will be filled in responses from the
	//  metrics API.
	UpdateTime *string `json:"updateTime,omitempty"`
}

// +kcc:proto=google.dataflow.v1beta3.MountedDataDisk
type MountedDataDisk struct {
	// The name of the data disk.
	//  This name is local to the Google Cloud Platform project and uniquely
	//  identifies the disk within that project, for example
	//  "myproject-1014-104817-4c2-harness-0-disk-1".
	DataDisk *string `json:"dataDisk,omitempty"`
}

// +kcc:proto=google.dataflow.v1beta3.Package
type Package struct {
	// The name of the package.
	Name *string `json:"name,omitempty"`

	// The resource to read the package from. The supported resource type is:
	//
	//  Google Cloud Storage:
	//
	//    storage.googleapis.com/{bucket}
	//    bucket.storage.googleapis.com/
	Location *string `json:"location,omitempty"`
}

// +kcc:proto=google.dataflow.v1beta3.PipelineDescription
type PipelineDescription struct {
	// Description of each transform in the pipeline and collections between them.
	OriginalPipelineTransform []TransformSummary `json:"originalPipelineTransform,omitempty"`

	// Description of each stage of execution of the pipeline.
	ExecutionPipelineStage []ExecutionStageSummary `json:"executionPipelineStage,omitempty"`

	// Pipeline level display data.
	DisplayData []DisplayData `json:"displayData,omitempty"`
}

// +kcc:proto=google.dataflow.v1beta3.ProgressTimeseries
type ProgressTimeseries struct {
	// The current progress of the component, in the range [0,1].
	CurrentProgress *float64 `json:"currentProgress,omitempty"`

	// History of progress for the component.
	//
	//  Points are sorted by time.
	DataPoints []ProgressTimeseries_Point `json:"dataPoints,omitempty"`
}

// +kcc:proto=google.dataflow.v1beta3.ProgressTimeseries.Point
type ProgressTimeseries_Point struct {
	// The timestamp of the point.
	Time *string `json:"time,omitempty"`

	// The value of the point.
	Value *float64 `json:"value,omitempty"`
}

// +kcc:proto=google.dataflow.v1beta3.PubSubIODetails
type PubSubIODetails struct {
	// Topic accessed in the connection.
	Topic *string `json:"topic,omitempty"`

	// Subscription used in the connection.
	Subscription *string `json:"subscription,omitempty"`
}

// +kcc:proto=google.dataflow.v1beta3.PubsubLocation
type PubsubLocation struct {
	// A pubsub topic, in the form of
	//  "pubsub.googleapis.com/topics/<project-id>/<topic-name>"
	Topic *string `json:"topic,omitempty"`

	// A pubsub subscription, in the form of
	//  "pubsub.googleapis.com/subscriptions/<project-id>/<subscription-name>"
	Subscription *string `json:"subscription,omitempty"`

	// If set, contains a pubsub label from which to extract record timestamps.
	//  If left empty, record timestamps will be generated upon arrival.
	TimestampLabel *string `json:"timestampLabel,omitempty"`

	// If set, contains a pubsub label from which to extract record ids.
	//  If left empty, record deduplication will be strictly best effort.
	IDLabel *string `json:"idLabel,omitempty"`

	// Indicates whether the pipeline allows late-arriving data.
	DropLateData *bool `json:"dropLateData,omitempty"`

	// If set, specifies the pubsub subscription that will be used for tracking
	//  custom time timestamps for watermark estimation.
	TrackingSubscription *string `json:"trackingSubscription,omitempty"`

	// If true, then the client has requested to get pubsub attributes.
	WithAttributes *bool `json:"withAttributes,omitempty"`
}

// +kcc:proto=google.dataflow.v1beta3.RuntimeEnvironment
type RuntimeEnvironment struct {
	// The initial number of Google Compute Engine instance for the job.
	NumWorkers *int32 `json:"numWorkers,omitempty"`

	// The maximum number of Google Compute Engine instances to be made
	//  available to your pipeline during execution, from 1 to 1000.
	MaxWorkers *int32 `json:"maxWorkers,omitempty"`

	// The Compute Engine [availability
	//  zone](https://cloud.google.com/compute/docs/regions-zones/regions-zones)
	//  for launching worker instances to run your pipeline.
	//  In the future, worker_zone will take precedence.
	Zone *string `json:"zone,omitempty"`

	// The email address of the service account to run the job as.
	ServiceAccountEmail *string `json:"serviceAccountEmail,omitempty"`

	// The Cloud Storage path to use for temporary files.
	//  Must be a valid Cloud Storage URL, beginning with `gs://`.
	TempLocation *string `json:"tempLocation,omitempty"`

	// Whether to bypass the safety checks for the job's temporary directory.
	//  Use with caution.
	BypassTempDirValidation *bool `json:"bypassTempDirValidation,omitempty"`

	// The machine type to use for the job. Defaults to the value from the
	//  template if not specified.
	MachineType *string `json:"machineType,omitempty"`

	// Additional experiment flags for the job, specified with the
	//  `--experiments` option.
	AdditionalExperiments []string `json:"additionalExperiments,omitempty"`

	// Network to which VMs will be assigned.  If empty or unspecified,
	//  the service will use the network "default".
	Network *string `json:"network,omitempty"`

	// Subnetwork to which VMs will be assigned, if desired. You can specify a
	//  subnetwork using either a complete URL or an abbreviated path. Expected to
	//  be of the form
	//  "https://www.googleapis.com/compute/v1/projects/HOST_PROJECT_ID/regions/REGION/subnetworks/SUBNETWORK"
	//  or "regions/REGION/subnetworks/SUBNETWORK". If the subnetwork is located in
	//  a Shared VPC network, you must use the complete URL.
	Subnetwork *string `json:"subnetwork,omitempty"`

	// Additional user labels to be specified for the job.
	//  Keys and values should follow the restrictions specified in the [labeling
	//  restrictions](https://cloud.google.com/compute/docs/labeling-resources#restrictions)
	//  page.
	//  An object containing a list of "key": value pairs.
	//  Example: { "name": "wrench", "mass": "1kg", "count": "3" }.
	AdditionalUserLabels map[string]string `json:"additionalUserLabels,omitempty"`

	// Name for the Cloud KMS key for the job.
	//  Key format is:
	//  projects/<project>/locations/<location>/keyRings/<keyring>/cryptoKeys/<key>
	KmsKeyName *string `json:"kmsKeyName,omitempty"`

	// Configuration for VM IPs.
	IpConfiguration *string `json:"ipConfiguration,omitempty"`

	// The Compute Engine region
	//  (https://cloud.google.com/compute/docs/regions-zones/regions-zones) in
	//  which worker processing should occur, e.g. "us-west1". Mutually exclusive
	//  with worker_zone. If neither worker_region nor worker_zone is specified,
	//  default to the control plane's region.
	WorkerRegion *string `json:"workerRegion,omitempty"`

	// The Compute Engine zone
	//  (https://cloud.google.com/compute/docs/regions-zones/regions-zones) in
	//  which worker processing should occur, e.g. "us-west1-a". Mutually exclusive
	//  with worker_region. If neither worker_region nor worker_zone is specified,
	//  a zone in the control plane's region is chosen based on available capacity.
	//  If both `worker_zone` and `zone` are set, `worker_zone` takes precedence.
	WorkerZone *string `json:"workerZone,omitempty"`

	// Whether to enable Streaming Engine for the job.
	EnableStreamingEngine *bool `json:"enableStreamingEngine,omitempty"`
}

// +kcc:proto=google.dataflow.v1beta3.SDKInfo
type SDKInfo struct {
	// Required. The SDK Language.
	Language *string `json:"language,omitempty"`

	// Optional. The SDK version.
	Version *string `json:"version,omitempty"`
}

// +kcc:proto=google.dataflow.v1beta3.SdkHarnessContainerImage
type SdkHarnessContainerImage struct {
	// A docker container image that resides in Google Container Registry.
	ContainerImage *string `json:"containerImage,omitempty"`

	// If true, recommends the Dataflow service to use only one core per SDK
	//  container instance with this image. If false (or unset) recommends using
	//  more than one core per SDK container instance with this image for
	//  efficiency. Note that Dataflow service may choose to override this property
	//  if needed.
	UseSingleCorePerContainer *bool `json:"useSingleCorePerContainer,omitempty"`

	// Environment ID for the Beam runner API proto Environment that corresponds
	//  to the current SDK Harness.
	EnvironmentID *string `json:"environmentID,omitempty"`

	// The set of capabilities enumerated in the above Environment proto. See also
	//  https://github.com/apache/beam/blob/master/model/pipeline/src/main/proto/beam_runner_api.proto
	Capabilities []string `json:"capabilities,omitempty"`
}

// +kcc:proto=google.dataflow.v1beta3.SdkVersion
type SdkVersion struct {
	// The version of the SDK used to run the job.
	Version *string `json:"version,omitempty"`

	// A readable string describing the version of the SDK.
	VersionDisplayName *string `json:"versionDisplayName,omitempty"`

	// The support status for this SDK version.
	SdkSupportStatus *string `json:"sdkSupportStatus,omitempty"`
}

// +kcc:proto=google.dataflow.v1beta3.Snapshot
type Snapshot struct {
	// The unique ID of this snapshot.
	ID *string `json:"id,omitempty"`

	// The project this snapshot belongs to.
	ProjectID *string `json:"projectID,omitempty"`

	// The job this snapshot was created from.
	SourceJobID *string `json:"sourceJobID,omitempty"`

	// The time this snapshot was created.
	CreationTime *string `json:"creationTime,omitempty"`

	// The time after which this snapshot will be automatically deleted.
	Ttl *string `json:"ttl,omitempty"`

	// State of the snapshot.
	State *string `json:"state,omitempty"`

	// Pub/Sub snapshot metadata.
	PubsubMetadata []PubsubSnapshotMetadata `json:"pubsubMetadata,omitempty"`

	// User specified description of the snapshot. Maybe empty.
	Description *string `json:"description,omitempty"`

	// The disk byte size of the snapshot. Only available for snapshots in READY
	//  state.
	DiskSizeBytes *int64 `json:"diskSizeBytes,omitempty"`

	// Cloud region where this snapshot lives in, e.g., "us-central1".
	Region *string `json:"region,omitempty"`
}

// +kcc:proto=google.dataflow.v1beta3.SpannerIODetails
type SpannerIODetails struct {
	// ProjectId accessed in the connection.
	ProjectID *string `json:"projectID,omitempty"`

	// InstanceId accessed in the connection.
	InstanceID *string `json:"instanceID,omitempty"`

	// DatabaseId accessed in the connection.
	DatabaseID *string `json:"databaseID,omitempty"`
}

// +kcc:proto=google.dataflow.v1beta3.StageExecutionDetails
type StageExecutionDetails struct {
	// Workers that have done work on the stage.
	Workers []WorkerDetails `json:"workers,omitempty"`

	// If present, this response does not contain all requested tasks.  To obtain
	//  the next page of results, repeat the request with page_token set to this
	//  value.
	NextPageToken *string `json:"nextPageToken,omitempty"`
}

// +kcc:proto=google.dataflow.v1beta3.StageSummary
type StageSummary struct {
	// ID of this stage
	StageID *string `json:"stageID,omitempty"`

	// State of this stage.
	State *string `json:"state,omitempty"`

	// Start time of this stage.
	StartTime *string `json:"startTime,omitempty"`

	// End time of this stage.
	//
	//  If the work item is completed, this is the actual end time of the stage.
	//  Otherwise, it is the predicted end time.
	EndTime *string `json:"endTime,omitempty"`

	// Progress for this stage.
	//  Only applicable to Batch jobs.
	Progress *ProgressTimeseries `json:"progress,omitempty"`

	// Metrics for this stage.
	Metrics []MetricUpdate `json:"metrics,omitempty"`
}

// +kcc:proto=google.dataflow.v1beta3.StateFamilyConfig
type StateFamilyConfig struct {
	// The state family value.
	StateFamily *string `json:"stateFamily,omitempty"`

	// If true, this family corresponds to a read operation.
	IsRead *bool `json:"isRead,omitempty"`
}

// +kcc:proto=google.dataflow.v1beta3.Step
type Step struct {
	// The kind of step in the Cloud Dataflow job.
	Kind *string `json:"kind,omitempty"`

	// The name that identifies the step. This must be unique for each
	//  step with respect to all other steps in the Cloud Dataflow job.
	Name *string `json:"name,omitempty"`

	// Named properties associated with the step. Each kind of
	//  predefined step has its own required set of properties.
	//  Must be provided on Create.  Only retrieved with JOB_VIEW_ALL.
	Properties *google_protobuf_Struct `json:"properties,omitempty"`
}

// +kcc:proto=google.dataflow.v1beta3.StreamLocation
type StreamLocation struct {
	// The stream is part of another computation within the current
	//  streaming Dataflow job.
	StreamingStageLocation *StreamingStageLocation `json:"streamingStageLocation,omitempty"`

	// The stream is a pubsub stream.
	PubsubLocation *PubsubLocation `json:"pubsubLocation,omitempty"`

	// The stream is a streaming side input.
	SideInputLocation *StreamingSideInputLocation `json:"sideInputLocation,omitempty"`

	// The stream is a custom source.
	CustomSourceLocation *CustomSourceLocation `json:"customSourceLocation,omitempty"`
}

// +kcc:proto=google.dataflow.v1beta3.StreamingApplianceSnapshotConfig
type StreamingApplianceSnapshotConfig struct {
	// If set, indicates the snapshot id for the snapshot being performed.
	SnapshotID *string `json:"snapshotID,omitempty"`

	// Indicates which endpoint is used to import appliance state.
	ImportStateEndpoint *string `json:"importStateEndpoint,omitempty"`
}

// +kcc:proto=google.dataflow.v1beta3.StreamingComputationRanges
type StreamingComputationRanges struct {
	// The ID of the computation.
	ComputationID *string `json:"computationID,omitempty"`

	// Data disk assignments for ranges from this computation.
	RangeAssignments []KeyRangeDataDiskAssignment `json:"rangeAssignments,omitempty"`
}

// +kcc:proto=google.dataflow.v1beta3.StreamingSideInputLocation
type StreamingSideInputLocation struct {
	// Identifies the particular side input within the streaming Dataflow job.
	Tag *string `json:"tag,omitempty"`

	// Identifies the state family where this side input is stored.
	StateFamily *string `json:"stateFamily,omitempty"`
}

// +kcc:proto=google.dataflow.v1beta3.StreamingStageLocation
type StreamingStageLocation struct {
	// Identifies the particular stream within the streaming Dataflow
	//  job.
	StreamID *string `json:"streamID,omitempty"`
}

// +kcc:proto=google.dataflow.v1beta3.StructuredMessage
type StructuredMessage struct {
	// Human-readable version of message.
	MessageText *string `json:"messageText,omitempty"`

	// Identifier for this message type.  Used by external systems to
	//  internationalize or personalize message.
	MessageKey *string `json:"messageKey,omitempty"`

	// The structured data associated with this message.
	Parameters []StructuredMessage_Parameter `json:"parameters,omitempty"`
}

// +kcc:proto=google.dataflow.v1beta3.StructuredMessage.Parameter
type StructuredMessage_Parameter struct {
	// Key or name for this parameter.
	Key *string `json:"key,omitempty"`

	// Value for this parameter.
	Value *google_protobuf_Value `json:"value,omitempty"`
}

// +kcc:proto=google.dataflow.v1beta3.TaskRunnerSettings
type TaskRunnerSettings struct {
	// The UNIX user ID on the worker VM to use for tasks launched by
	//  taskrunner; e.g. "root".
	TaskUser *string `json:"taskUser,omitempty"`

	// The UNIX group ID on the worker VM to use for tasks launched by
	//  taskrunner; e.g. "wheel".
	TaskGroup *string `json:"taskGroup,omitempty"`

	// The OAuth2 scopes to be requested by the taskrunner in order to
	//  access the Cloud Dataflow API.
	OauthScopes []string `json:"oauthScopes,omitempty"`

	// The base URL for the taskrunner to use when accessing Google Cloud APIs.
	//
	//  When workers access Google Cloud APIs, they logically do so via
	//  relative URLs.  If this field is specified, it supplies the base
	//  URL to use for resolving these relative URLs.  The normative
	//  algorithm used is defined by RFC 1808, "Relative Uniform Resource
	//  Locators".
	//
	//  If not specified, the default value is "http://www.googleapis.com/"
	BaseURL *string `json:"baseURL,omitempty"`

	// The API version of endpoint, e.g. "v1b3"
	DataflowApiVersion *string `json:"dataflowApiVersion,omitempty"`

	// The settings to pass to the parallel worker harness.
	ParallelWorkerSettings *WorkerSettings `json:"parallelWorkerSettings,omitempty"`

	// The location on the worker for task-specific subdirectories.
	BaseTaskDir *string `json:"baseTaskDir,omitempty"`

	// Whether to continue taskrunner if an exception is hit.
	ContinueOnException *bool `json:"continueOnException,omitempty"`

	// Whether to send taskrunner log info to Google Compute Engine VM serial
	//  console.
	LogToSerialconsole *bool `json:"logToSerialconsole,omitempty"`

	// Whether to also send taskrunner log info to stderr.
	Alsologtostderr *bool `json:"alsologtostderr,omitempty"`

	// Indicates where to put logs.  If this is not specified, the logs
	//  will not be uploaded.
	//
	//  The supported resource type is:
	//
	//  Google Cloud Storage:
	//    storage.googleapis.com/{bucket}/{object}
	//    bucket.storage.googleapis.com/{object}
	LogUploadLocation *string `json:"logUploadLocation,omitempty"`

	// The directory on the VM to store logs.
	LogDir *string `json:"logDir,omitempty"`

	// The prefix of the resources the taskrunner should use for
	//  temporary storage.
	//
	//  The supported resource type is:
	//
	//  Google Cloud Storage:
	//    storage.googleapis.com/{bucket}/{object}
	//    bucket.storage.googleapis.com/{object}
	TempStoragePrefix *string `json:"tempStoragePrefix,omitempty"`

	// The command to launch the worker harness.
	HarnessCommand *string `json:"harnessCommand,omitempty"`

	// The file to store the workflow in.
	WorkflowFileName *string `json:"workflowFileName,omitempty"`

	// The file to store preprocessing commands in.
	CommandlinesFileName *string `json:"commandlinesFileName,omitempty"`

	// The ID string of the VM.
	VmID *string `json:"vmID,omitempty"`

	// The suggested backend language.
	LanguageHint *string `json:"languageHint,omitempty"`

	// The streaming worker main class name.
	StreamingWorkerMainClass *string `json:"streamingWorkerMainClass,omitempty"`
}

// +kcc:proto=google.dataflow.v1beta3.TopologyConfig
type TopologyConfig struct {
	// The computations associated with a streaming Dataflow job.
	Computations []ComputationTopology `json:"computations,omitempty"`

	// The disks assigned to a streaming Dataflow job.
	DataDiskAssignments []DataDiskAssignment `json:"dataDiskAssignments,omitempty"`

	// Maps user stage names to stable computation names.
	UserStageToComputationNameMap map[string]string `json:"userStageToComputationNameMap,omitempty"`

	// The size (in bits) of keys that will be assigned to source messages.
	ForwardingKeyBits *int32 `json:"forwardingKeyBits,omitempty"`

	// Version number for persistent state.
	PersistentStateVersion *int32 `json:"persistentStateVersion,omitempty"`
}

// +kcc:proto=google.dataflow.v1beta3.TransformSummary
type TransformSummary struct {
	// Type of transform.
	Kind *string `json:"kind,omitempty"`

	// SDK generated id of this transform instance.
	ID *string `json:"id,omitempty"`

	// User provided name for this transform instance.
	Name *string `json:"name,omitempty"`

	// Transform-specific display data.
	DisplayData []DisplayData `json:"displayData,omitempty"`

	// User  names for all collection outputs to this transform.
	OutputCollectionName []string `json:"outputCollectionName,omitempty"`

	// User names for all collection inputs to this transform.
	InputCollectionName []string `json:"inputCollectionName,omitempty"`
}

// +kcc:proto=google.dataflow.v1beta3.WorkItemDetails
type WorkItemDetails struct {
	// Name of this work item.
	TaskID *string `json:"taskID,omitempty"`

	// Attempt ID of this work item
	AttemptID *string `json:"attemptID,omitempty"`

	// Start time of this work item attempt.
	StartTime *string `json:"startTime,omitempty"`

	// End time of this work item attempt.
	//
	//  If the work item is completed, this is the actual end time of the work
	//  item.  Otherwise, it is the predicted end time.
	EndTime *string `json:"endTime,omitempty"`

	// State of this work item.
	State *string `json:"state,omitempty"`

	// Progress of this work item.
	Progress *ProgressTimeseries `json:"progress,omitempty"`

	// Metrics for this work item.
	Metrics []MetricUpdate `json:"metrics,omitempty"`
}

// +kcc:proto=google.dataflow.v1beta3.WorkerDetails
type WorkerDetails struct {
	// Name of this worker
	WorkerName *string `json:"workerName,omitempty"`

	// Work items processed by this worker, sorted by time.
	WorkItems []WorkItemDetails `json:"workItems,omitempty"`
}

// +kcc:proto=google.dataflow.v1beta3.WorkerPool
type WorkerPool struct {
	// The kind of the worker pool; currently only `harness` and `shuffle`
	//  are supported.
	Kind *string `json:"kind,omitempty"`

	// Number of Google Compute Engine workers in this pool needed to
	//  execute the job.  If zero or unspecified, the service will
	//  attempt to choose a reasonable default.
	NumWorkers *int32 `json:"numWorkers,omitempty"`

	// Packages to be installed on workers.
	Packages []Package `json:"packages,omitempty"`

	// The default package set to install.  This allows the service to
	//  select a default set of packages which are useful to worker
	//  harnesses written in a particular language.
	DefaultPackageSet *string `json:"defaultPackageSet,omitempty"`

	// Machine type (e.g. "n1-standard-1").  If empty or unspecified, the
	//  service will attempt to choose a reasonable default.
	MachineType *string `json:"machineType,omitempty"`

	// Sets the policy for determining when to turndown worker pool.
	//  Allowed values are: `TEARDOWN_ALWAYS`, `TEARDOWN_ON_SUCCESS`, and
	//  `TEARDOWN_NEVER`.
	//  `TEARDOWN_ALWAYS` means workers are always torn down regardless of whether
	//  the job succeeds. `TEARDOWN_ON_SUCCESS` means workers are torn down
	//  if the job succeeds. `TEARDOWN_NEVER` means the workers are never torn
	//  down.
	//
	//  If the workers are not torn down by the service, they will
	//  continue to run and use Google Compute Engine VM resources in the
	//  user's project until they are explicitly terminated by the user.
	//  Because of this, Google recommends using the `TEARDOWN_ALWAYS`
	//  policy except for small, manually supervised test jobs.
	//
	//  If unknown or unspecified, the service will attempt to choose a reasonable
	//  default.
	TeardownPolicy *string `json:"teardownPolicy,omitempty"`

	// Size of root disk for VMs, in GB.  If zero or unspecified, the service will
	//  attempt to choose a reasonable default.
	DiskSizeGb *int32 `json:"diskSizeGb,omitempty"`

	// Type of root disk for VMs.  If empty or unspecified, the service will
	//  attempt to choose a reasonable default.
	DiskType *string `json:"diskType,omitempty"`

	// Fully qualified source image for disks.
	DiskSourceImage *string `json:"diskSourceImage,omitempty"`

	// Zone to run the worker pools in.  If empty or unspecified, the service
	//  will attempt to choose a reasonable default.
	Zone *string `json:"zone,omitempty"`

	// Settings passed through to Google Compute Engine workers when
	//  using the standard Dataflow task runner.  Users should ignore
	//  this field.
	TaskrunnerSettings *TaskRunnerSettings `json:"taskrunnerSettings,omitempty"`

	// The action to take on host maintenance, as defined by the Google
	//  Compute Engine API.
	OnHostMaintenance *string `json:"onHostMaintenance,omitempty"`

	// Data disks that are used by a VM in this workflow.
	DataDisks []Disk `json:"dataDisks,omitempty"`

	// Metadata to set on the Google Compute Engine VMs.
	Metadata map[string]string `json:"metadata,omitempty"`

	// Settings for autoscaling of this WorkerPool.
	AutoscalingSettings *AutoscalingSettings `json:"autoscalingSettings,omitempty"`

	// Extra arguments for this worker pool.
	PoolArgs *google_protobuf_Any `json:"poolArgs,omitempty"`

	// Network to which VMs will be assigned.  If empty or unspecified,
	//  the service will use the network "default".
	Network *string `json:"network,omitempty"`

	// Subnetwork to which VMs will be assigned, if desired.  Expected to be of
	//  the form "regions/REGION/subnetworks/SUBNETWORK".
	Subnetwork *string `json:"subnetwork,omitempty"`

	// Required. Docker container image that executes the Cloud Dataflow worker
	//  harness, residing in Google Container Registry.
	//
	//  Deprecated for the Fn API path. Use sdk_harness_container_images instead.
	WorkerHarnessContainerImage *string `json:"workerHarnessContainerImage,omitempty"`

	// The number of threads per worker harness. If empty or unspecified, the
	//  service will choose a number of threads (according to the number of cores
	//  on the selected machine type for batch, or 1 by convention for streaming).
	NumThreadsPerWorker *int32 `json:"numThreadsPerWorker,omitempty"`

	// Configuration for VM IPs.
	IpConfiguration *string `json:"ipConfiguration,omitempty"`

	// Set of SDK harness containers needed to execute this pipeline. This will
	//  only be set in the Fn API path. For non-cross-language pipelines this
	//  should have only one entry. Cross-language pipelines will have two or more
	//  entries.
	SdkHarnessContainerImages []SdkHarnessContainerImage `json:"sdkHarnessContainerImages,omitempty"`
}

// +kcc:proto=google.dataflow.v1beta3.WorkerSettings
type WorkerSettings struct {
	// The base URL for accessing Google Cloud APIs.
	//
	//  When workers access Google Cloud APIs, they logically do so via
	//  relative URLs.  If this field is specified, it supplies the base
	//  URL to use for resolving these relative URLs.  The normative
	//  algorithm used is defined by RFC 1808, "Relative Uniform Resource
	//  Locators".
	//
	//  If not specified, the default value is "http://www.googleapis.com/"
	BaseURL *string `json:"baseURL,omitempty"`

	// Whether to send work progress updates to the service.
	ReportingEnabled *bool `json:"reportingEnabled,omitempty"`

	// The Cloud Dataflow service path relative to the root URL, for example,
	//  "dataflow/v1b3/projects".
	ServicePath *string `json:"servicePath,omitempty"`

	// The Shuffle service path relative to the root URL, for example,
	//  "shuffle/v1beta1".
	ShuffleServicePath *string `json:"shuffleServicePath,omitempty"`

	// The ID of the worker running this pipeline.
	WorkerID *string `json:"workerID,omitempty"`

	// The prefix of the resources the system should use for temporary
	//  storage.
	//
	//  The supported resource type is:
	//
	//  Google Cloud Storage:
	//
	//    storage.googleapis.com/{bucket}/{object}
	//    bucket.storage.googleapis.com/{object}
	TempStoragePrefix *string `json:"tempStoragePrefix,omitempty"`
}
