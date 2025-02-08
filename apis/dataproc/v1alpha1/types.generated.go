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

// +kcc:proto=google.cloud.dataproc.v1.SessionTemplate
type SessionTemplate struct {
	// Required. The resource name of the session template.
	// +kcc:proto:field=google.cloud.dataproc.v1.SessionTemplate.name
	Name *string `json:"name,omitempty"`

	// Optional. Brief description of the template.
	// +kcc:proto:field=google.cloud.dataproc.v1.SessionTemplate.description
	Description *string `json:"description,omitempty"`

	// Optional. Jupyter session config.
	// +kcc:proto:field=google.cloud.dataproc.v1.SessionTemplate.jupyter_session
	JupyterSession *JupyterConfig `json:"jupyterSession,omitempty"`

	// Optional. Spark Connect session config.
	// +kcc:proto:field=google.cloud.dataproc.v1.SessionTemplate.spark_connect_session
	SparkConnectSession *SparkConnectConfig `json:"sparkConnectSession,omitempty"`

	// Optional. Labels to associate with sessions created using this template.
	//  Label **keys** must contain 1 to 63 characters, and must conform to
	//  [RFC 1035](https://www.ietf.org/rfc/rfc1035.txt).
	//  Label **values** can be empty, but, if present, must contain 1 to 63
	//  characters and conform to [RFC
	//  1035](https://www.ietf.org/rfc/rfc1035.txt). No more than 32 labels can be
	//  associated with a session.
	// +kcc:proto:field=google.cloud.dataproc.v1.SessionTemplate.labels
	Labels map[string]string `json:"labels,omitempty"`

	// Optional. Runtime configuration for session execution.
	// +kcc:proto:field=google.cloud.dataproc.v1.SessionTemplate.runtime_config
	RuntimeConfig *RuntimeConfig `json:"runtimeConfig,omitempty"`

	// Optional. Environment configuration for session execution.
	// +kcc:proto:field=google.cloud.dataproc.v1.SessionTemplate.environment_config
	EnvironmentConfig *EnvironmentConfig `json:"environmentConfig,omitempty"`
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

// +kcc:proto=google.cloud.dataproc.v1.SessionTemplate
type SessionTemplateObservedState struct {
	// Output only. The time when the template was created.
	// +kcc:proto:field=google.cloud.dataproc.v1.SessionTemplate.create_time
	CreateTime *string `json:"createTime,omitempty"`

	// Output only. The email address of the user who created the template.
	// +kcc:proto:field=google.cloud.dataproc.v1.SessionTemplate.creator
	Creator *string `json:"creator,omitempty"`

	// Output only. The time the template was last updated.
	// +kcc:proto:field=google.cloud.dataproc.v1.SessionTemplate.update_time
	UpdateTime *string `json:"updateTime,omitempty"`

	// Output only. A session template UUID (Unique Universal Identifier). The
	//  service generates this value when it creates the session template.
	// +kcc:proto:field=google.cloud.dataproc.v1.SessionTemplate.uuid
	Uuid *string `json:"uuid,omitempty"`
}
