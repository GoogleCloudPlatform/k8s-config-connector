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


// +kcc:proto=google.cloud.dataproc.v1.AutotuningConfig
type AutotuningConfig struct {
	// Optional. Scenarios for which tunings are applied.
	// +kcc:proto:field=google.cloud.dataproc.v1.AutotuningConfig.scenarios
	Scenarios []string `json:"scenarios,omitempty"`
}

// +kcc:proto=google.cloud.dataproc.v1.EnvironmentConfig
type EnvironmentConfig struct {
	// Optional. Execution configuration for a workload.
	// +kcc:proto:field=google.cloud.dataproc.v1.EnvironmentConfig.execution_config
	ExecutionConfig *ExecutionConfig `json:"executionConfig,omitempty"`

	// Optional. Peripherals configuration that workload has access to.
	// +kcc:proto:field=google.cloud.dataproc.v1.EnvironmentConfig.peripherals_config
	PeripheralsConfig *PeripheralsConfig `json:"peripheralsConfig,omitempty"`
}

// +kcc:proto=google.cloud.dataproc.v1.ExecutionConfig
type ExecutionConfig struct {
	// Optional. Service account that used to execute workload.
	// +kcc:proto:field=google.cloud.dataproc.v1.ExecutionConfig.service_account
	ServiceAccount *string `json:"serviceAccount,omitempty"`

	// Optional. Network URI to connect workload to.
	// +kcc:proto:field=google.cloud.dataproc.v1.ExecutionConfig.network_uri
	NetworkURI *string `json:"networkURI,omitempty"`

	// Optional. Subnetwork URI to connect workload to.
	// +kcc:proto:field=google.cloud.dataproc.v1.ExecutionConfig.subnetwork_uri
	SubnetworkURI *string `json:"subnetworkURI,omitempty"`

	// Optional. Tags used for network traffic control.
	// +kcc:proto:field=google.cloud.dataproc.v1.ExecutionConfig.network_tags
	NetworkTags []string `json:"networkTags,omitempty"`

	// Optional. The Cloud KMS key to use for encryption.
	// +kcc:proto:field=google.cloud.dataproc.v1.ExecutionConfig.kms_key
	KMSKey *string `json:"kmsKey,omitempty"`

	// Optional. Applies to sessions only. The duration to keep the session alive
	//  while it's idling. Exceeding this threshold causes the session to
	//  terminate. This field cannot be set on a batch workload. Minimum value is
	//  10 minutes; maximum value is 14 days (see JSON representation of
	//  [Duration](https://developers.google.com/protocol-buffers/docs/proto3#json)).
	//  Defaults to 1 hour if not set.
	//  If both `ttl` and `idle_ttl` are specified for an interactive session,
	//  the conditions are treated as `OR` conditions: the workload will be
	//  terminated when it has been idle for `idle_ttl` or when `ttl` has been
	//  exceeded, whichever occurs first.
	// +kcc:proto:field=google.cloud.dataproc.v1.ExecutionConfig.idle_ttl
	IdleTtl *string `json:"idleTtl,omitempty"`

	// Optional. The duration after which the workload will be terminated,
	//  specified as the JSON representation for
	//  [Duration](https://protobuf.dev/programming-guides/proto3/#json).
	//  When the workload exceeds this duration, it will be unconditionally
	//  terminated without waiting for ongoing work to finish. If `ttl` is not
	//  specified for a batch workload, the workload will be allowed to run until
	//  it exits naturally (or run forever without exiting). If `ttl` is not
	//  specified for an interactive session, it defaults to 24 hours. If `ttl` is
	//  not specified for a batch that uses 2.1+ runtime version, it defaults to 4
	//  hours. Minimum value is 10 minutes; maximum value is 14 days. If both `ttl`
	//  and `idle_ttl` are specified (for an interactive session), the conditions
	//  are treated as `OR` conditions: the workload will be terminated when it has
	//  been idle for `idle_ttl` or when `ttl` has been exceeded, whichever occurs
	//  first.
	// +kcc:proto:field=google.cloud.dataproc.v1.ExecutionConfig.ttl
	Ttl *string `json:"ttl,omitempty"`

	// Optional. A Cloud Storage bucket used to stage workload dependencies,
	//  config files, and store workload output and other ephemeral data, such as
	//  Spark history files. If you do not specify a staging bucket, Cloud Dataproc
	//  will determine a Cloud Storage location according to the region where your
	//  workload is running, and then create and manage project-level, per-location
	//  staging and temporary buckets.
	//  **This field requires a Cloud Storage bucket name, not a `gs://...` URI to
	//  a Cloud Storage bucket.**
	// +kcc:proto:field=google.cloud.dataproc.v1.ExecutionConfig.staging_bucket
	StagingBucket *string `json:"stagingBucket,omitempty"`
}

// +kcc:proto=google.cloud.dataproc.v1.JupyterConfig
type JupyterConfig struct {
	// Optional. Kernel
	// +kcc:proto:field=google.cloud.dataproc.v1.JupyterConfig.kernel
	Kernel *string `json:"kernel,omitempty"`

	// Optional. Display name, shown in the Jupyter kernelspec card.
	// +kcc:proto:field=google.cloud.dataproc.v1.JupyterConfig.display_name
	DisplayName *string `json:"displayName,omitempty"`
}

// +kcc:proto=google.cloud.dataproc.v1.PeripheralsConfig
type PeripheralsConfig struct {
	// Optional. Resource name of an existing Dataproc Metastore service.
	//
	//  Example:
	//
	//  * `projects/[project_id]/locations/[region]/services/[service_id]`
	// +kcc:proto:field=google.cloud.dataproc.v1.PeripheralsConfig.metastore_service
	MetastoreService *string `json:"metastoreService,omitempty"`

	// Optional. The Spark History Server configuration for the workload.
	// +kcc:proto:field=google.cloud.dataproc.v1.PeripheralsConfig.spark_history_server_config
	SparkHistoryServerConfig *SparkHistoryServerConfig `json:"sparkHistoryServerConfig,omitempty"`
}

// +kcc:proto=google.cloud.dataproc.v1.PyPiRepositoryConfig
type PyPiRepositoryConfig struct {
	// Optional. PyPi repository address
	// +kcc:proto:field=google.cloud.dataproc.v1.PyPiRepositoryConfig.pypi_repository
	PypiRepository *string `json:"pypiRepository,omitempty"`
}

// +kcc:proto=google.cloud.dataproc.v1.RepositoryConfig
type RepositoryConfig struct {
	// Optional. Configuration for PyPi repository.
	// +kcc:proto:field=google.cloud.dataproc.v1.RepositoryConfig.pypi_repository_config
	PypiRepositoryConfig *PyPiRepositoryConfig `json:"pypiRepositoryConfig,omitempty"`
}

// +kcc:proto=google.cloud.dataproc.v1.RuntimeConfig
type RuntimeConfig struct {
	// Optional. Version of the batch runtime.
	// +kcc:proto:field=google.cloud.dataproc.v1.RuntimeConfig.version
	Version *string `json:"version,omitempty"`

	// Optional. Optional custom container image for the job runtime environment.
	//  If not specified, a default container image will be used.
	// +kcc:proto:field=google.cloud.dataproc.v1.RuntimeConfig.container_image
	ContainerImage *string `json:"containerImage,omitempty"`

	// Optional. A mapping of property names to values, which are used to
	//  configure workload execution.
	// +kcc:proto:field=google.cloud.dataproc.v1.RuntimeConfig.properties
	Properties map[string]string `json:"properties,omitempty"`

	// Optional. Dependency repository configuration.
	// +kcc:proto:field=google.cloud.dataproc.v1.RuntimeConfig.repository_config
	RepositoryConfig *RepositoryConfig `json:"repositoryConfig,omitempty"`

	// Optional. Autotuning configuration of the workload.
	// +kcc:proto:field=google.cloud.dataproc.v1.RuntimeConfig.autotuning_config
	AutotuningConfig *AutotuningConfig `json:"autotuningConfig,omitempty"`

	// Optional. Cohort identifier. Identifies families of the workloads having
	//  the same shape, e.g. daily ETL jobs.
	// +kcc:proto:field=google.cloud.dataproc.v1.RuntimeConfig.cohort
	Cohort *string `json:"cohort,omitempty"`
}

// +kcc:proto=google.cloud.dataproc.v1.RuntimeInfo
type RuntimeInfo struct {
}

// +kcc:proto=google.cloud.dataproc.v1.Session
type Session struct {
	// Required. The resource name of the session.
	// +kcc:proto:field=google.cloud.dataproc.v1.Session.name
	Name *string `json:"name,omitempty"`

	// Optional. Jupyter session config.
	// +kcc:proto:field=google.cloud.dataproc.v1.Session.jupyter_session
	JupyterSession *JupyterConfig `json:"jupyterSession,omitempty"`

	// Optional. Spark Connect session config.
	// +kcc:proto:field=google.cloud.dataproc.v1.Session.spark_connect_session
	SparkConnectSession *SparkConnectConfig `json:"sparkConnectSession,omitempty"`

	// Optional. The labels to associate with the session.
	//  Label **keys** must contain 1 to 63 characters, and must conform to
	//  [RFC 1035](https://www.ietf.org/rfc/rfc1035.txt).
	//  Label **values** may be empty, but, if present, must contain 1 to 63
	//  characters, and must conform to [RFC
	//  1035](https://www.ietf.org/rfc/rfc1035.txt). No more than 32 labels can be
	//  associated with a session.
	// +kcc:proto:field=google.cloud.dataproc.v1.Session.labels
	Labels map[string]string `json:"labels,omitempty"`

	// Optional. Runtime configuration for the session execution.
	// +kcc:proto:field=google.cloud.dataproc.v1.Session.runtime_config
	RuntimeConfig *RuntimeConfig `json:"runtimeConfig,omitempty"`

	// Optional. Environment configuration for the session execution.
	// +kcc:proto:field=google.cloud.dataproc.v1.Session.environment_config
	EnvironmentConfig *EnvironmentConfig `json:"environmentConfig,omitempty"`

	// Optional. The email address of the user who owns the session.
	// +kcc:proto:field=google.cloud.dataproc.v1.Session.user
	User *string `json:"user,omitempty"`

	// Optional. The session template used by the session.
	//
	//  Only resource names, including project ID and location, are valid.
	//
	//  Example:
	//  * `https://www.googleapis.com/compute/v1/projects/[project_id]/locations/[dataproc_region]/sessionTemplates/[template_id]`
	//  * `projects/[project_id]/locations/[dataproc_region]/sessionTemplates/[template_id]`
	//
	//  The template must be in the same project and Dataproc region as the
	//  session.
	// +kcc:proto:field=google.cloud.dataproc.v1.Session.session_template
	SessionTemplate *string `json:"sessionTemplate,omitempty"`
}

// +kcc:proto=google.cloud.dataproc.v1.Session.SessionStateHistory
type Session_SessionStateHistory struct {
}

// +kcc:proto=google.cloud.dataproc.v1.SparkConnectConfig
type SparkConnectConfig struct {
}

// +kcc:proto=google.cloud.dataproc.v1.SparkHistoryServerConfig
type SparkHistoryServerConfig struct {
	// Optional. Resource name of an existing Dataproc Cluster to act as a Spark
	//  History Server for the workload.
	//
	//  Example:
	//
	//  * `projects/[project_id]/regions/[region]/clusters/[cluster_name]`
	// +kcc:proto:field=google.cloud.dataproc.v1.SparkHistoryServerConfig.dataproc_cluster
	DataprocCluster *string `json:"dataprocCluster,omitempty"`
}

// +kcc:proto=google.cloud.dataproc.v1.UsageMetrics
type UsageMetrics struct {
	// Optional. DCU (Dataproc Compute Units) usage in (`milliDCU` x `seconds`)
	//  (see [Dataproc Serverless pricing]
	//  (https://cloud.google.com/dataproc-serverless/pricing)).
	// +kcc:proto:field=google.cloud.dataproc.v1.UsageMetrics.milli_dcu_seconds
	MilliDcuSeconds *int64 `json:"milliDcuSeconds,omitempty"`

	// Optional. Shuffle storage usage in (`GB` x `seconds`) (see
	//  [Dataproc Serverless pricing]
	//  (https://cloud.google.com/dataproc-serverless/pricing)).
	// +kcc:proto:field=google.cloud.dataproc.v1.UsageMetrics.shuffle_storage_gb_seconds
	ShuffleStorageGBSeconds *int64 `json:"shuffleStorageGBSeconds,omitempty"`

	// Optional. Accelerator usage in (`milliAccelerator` x `seconds`) (see
	//  [Dataproc Serverless pricing]
	//  (https://cloud.google.com/dataproc-serverless/pricing)).
	// +kcc:proto:field=google.cloud.dataproc.v1.UsageMetrics.milli_accelerator_seconds
	MilliAcceleratorSeconds *int64 `json:"milliAcceleratorSeconds,omitempty"`

	// Optional. Accelerator type being used, if any
	// +kcc:proto:field=google.cloud.dataproc.v1.UsageMetrics.accelerator_type
	AcceleratorType *string `json:"acceleratorType,omitempty"`
}

// +kcc:proto=google.cloud.dataproc.v1.UsageSnapshot
type UsageSnapshot struct {
	// Optional. Milli (one-thousandth) Dataproc Compute Units (DCUs) (see
	//  [Dataproc Serverless pricing]
	//  (https://cloud.google.com/dataproc-serverless/pricing)).
	// +kcc:proto:field=google.cloud.dataproc.v1.UsageSnapshot.milli_dcu
	MilliDcu *int64 `json:"milliDcu,omitempty"`

	// Optional. Shuffle Storage in gigabytes (GB). (see [Dataproc Serverless
	//  pricing] (https://cloud.google.com/dataproc-serverless/pricing))
	// +kcc:proto:field=google.cloud.dataproc.v1.UsageSnapshot.shuffle_storage_gb
	ShuffleStorageGB *int64 `json:"shuffleStorageGB,omitempty"`

	// Optional. Milli (one-thousandth) Dataproc Compute Units (DCUs) charged at
	//  premium tier (see [Dataproc Serverless pricing]
	//  (https://cloud.google.com/dataproc-serverless/pricing)).
	// +kcc:proto:field=google.cloud.dataproc.v1.UsageSnapshot.milli_dcu_premium
	MilliDcuPremium *int64 `json:"milliDcuPremium,omitempty"`

	// Optional. Shuffle Storage in gigabytes (GB) charged at premium tier. (see
	//  [Dataproc Serverless pricing]
	//  (https://cloud.google.com/dataproc-serverless/pricing))
	// +kcc:proto:field=google.cloud.dataproc.v1.UsageSnapshot.shuffle_storage_gb_premium
	ShuffleStorageGBPremium *int64 `json:"shuffleStorageGBPremium,omitempty"`

	// Optional. Milli (one-thousandth) accelerator. (see [Dataproc
	//  Serverless pricing] (https://cloud.google.com/dataproc-serverless/pricing))
	// +kcc:proto:field=google.cloud.dataproc.v1.UsageSnapshot.milli_accelerator
	MilliAccelerator *int64 `json:"milliAccelerator,omitempty"`

	// Optional. Accelerator type being used, if any
	// +kcc:proto:field=google.cloud.dataproc.v1.UsageSnapshot.accelerator_type
	AcceleratorType *string `json:"acceleratorType,omitempty"`

	// Optional. The timestamp of the usage snapshot.
	// +kcc:proto:field=google.cloud.dataproc.v1.UsageSnapshot.snapshot_time
	SnapshotTime *string `json:"snapshotTime,omitempty"`
}

// +kcc:proto=google.cloud.dataproc.v1.RuntimeInfo
type RuntimeInfoObservedState struct {
	// Output only. Map of remote access endpoints (such as web interfaces and
	//  APIs) to their URIs.
	// +kcc:proto:field=google.cloud.dataproc.v1.RuntimeInfo.endpoints
	Endpoints map[string]string `json:"endpoints,omitempty"`

	// Output only. A URI pointing to the location of the stdout and stderr of the
	//  workload.
	// +kcc:proto:field=google.cloud.dataproc.v1.RuntimeInfo.output_uri
	OutputURI *string `json:"outputURI,omitempty"`

	// Output only. A URI pointing to the location of the diagnostics tarball.
	// +kcc:proto:field=google.cloud.dataproc.v1.RuntimeInfo.diagnostic_output_uri
	DiagnosticOutputURI *string `json:"diagnosticOutputURI,omitempty"`

	// Output only. Approximate workload resource usage, calculated when
	//  the workload completes (see [Dataproc Serverless pricing]
	//  (https://cloud.google.com/dataproc-serverless/pricing)).
	//
	//  **Note:** This metric calculation may change in the future, for
	//  example, to capture cumulative workload resource
	//  consumption during workload execution (see the
	//  [Dataproc Serverless release notes]
	//  (https://cloud.google.com/dataproc-serverless/docs/release-notes)
	//  for announcements, changes, fixes
	//  and other Dataproc developments).
	// +kcc:proto:field=google.cloud.dataproc.v1.RuntimeInfo.approximate_usage
	ApproximateUsage *UsageMetrics `json:"approximateUsage,omitempty"`

	// Output only. Snapshot of current workload resource usage.
	// +kcc:proto:field=google.cloud.dataproc.v1.RuntimeInfo.current_usage
	CurrentUsage *UsageSnapshot `json:"currentUsage,omitempty"`
}

// +kcc:proto=google.cloud.dataproc.v1.Session
type SessionObservedState struct {
	// Output only. A session UUID (Unique Universal Identifier). The service
	//  generates this value when it creates the session.
	// +kcc:proto:field=google.cloud.dataproc.v1.Session.uuid
	Uuid *string `json:"uuid,omitempty"`

	// Output only. The time when the session was created.
	// +kcc:proto:field=google.cloud.dataproc.v1.Session.create_time
	CreateTime *string `json:"createTime,omitempty"`

	// Output only. Runtime information about session execution.
	// +kcc:proto:field=google.cloud.dataproc.v1.Session.runtime_info
	RuntimeInfo *RuntimeInfo `json:"runtimeInfo,omitempty"`

	// Output only. A state of the session.
	// +kcc:proto:field=google.cloud.dataproc.v1.Session.state
	State *string `json:"state,omitempty"`

	// Output only. Session state details, such as the failure
	//  description if the state is `FAILED`.
	// +kcc:proto:field=google.cloud.dataproc.v1.Session.state_message
	StateMessage *string `json:"stateMessage,omitempty"`

	// Output only. The time when the session entered the current state.
	// +kcc:proto:field=google.cloud.dataproc.v1.Session.state_time
	StateTime *string `json:"stateTime,omitempty"`

	// Output only. The email address of the user who created the session.
	// +kcc:proto:field=google.cloud.dataproc.v1.Session.creator
	Creator *string `json:"creator,omitempty"`

	// Output only. Historical state information for the session.
	// +kcc:proto:field=google.cloud.dataproc.v1.Session.state_history
	StateHistory []Session_SessionStateHistory `json:"stateHistory,omitempty"`
}

// +kcc:proto=google.cloud.dataproc.v1.Session.SessionStateHistory
type Session_SessionStateHistoryObservedState struct {
	// Output only. The state of the session at this point in the session
	//  history.
	// +kcc:proto:field=google.cloud.dataproc.v1.Session.SessionStateHistory.state
	State *string `json:"state,omitempty"`

	// Output only. Details about the state at this point in the session
	//  history.
	// +kcc:proto:field=google.cloud.dataproc.v1.Session.SessionStateHistory.state_message
	StateMessage *string `json:"stateMessage,omitempty"`

	// Output only. The time when the session entered the historical state.
	// +kcc:proto:field=google.cloud.dataproc.v1.Session.SessionStateHistory.state_start_time
	StateStartTime *string `json:"stateStartTime,omitempty"`
}
