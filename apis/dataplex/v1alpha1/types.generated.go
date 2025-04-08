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

// +generated:types
// krm.group: dataplex.cnrm.cloud.google.com
// krm.version: v1alpha1
// proto.service: google.cloud.dataplex.v1
// resource: DataplexTask:Task

package v1alpha1

import (
	refsv1beta1 "github.com/GoogleCloudPlatform/k8s-config-connector/apis/refs/v1beta1"
)

// +kcc:proto=google.cloud.dataplex.v1.AssetStatus
type AssetStatus struct {
	// Last update time of the status.
	// +kcc:proto:field=google.cloud.dataplex.v1.AssetStatus.update_time
	UpdateTime *string `json:"updateTime,omitempty"`

	// Number of active assets.
	// +kcc:proto:field=google.cloud.dataplex.v1.AssetStatus.active_assets
	ActiveAssets *int32 `json:"activeAssets,omitempty"`

	// Number of assets that are in process of updating the security policy on
	//  attached resources.
	// +kcc:proto:field=google.cloud.dataplex.v1.AssetStatus.security_policy_applying_assets
	SecurityPolicyApplyingAssets *int32 `json:"securityPolicyApplyingAssets,omitempty"`
}

// +kcc:proto=google.cloud.dataplex.v1.Lake
type Lake struct {

	// Optional. User friendly display name.
	// +kcc:proto:field=google.cloud.dataplex.v1.Lake.display_name
	DisplayName *string `json:"displayName,omitempty"`

	// Optional. User-defined labels for the lake.
	// +kcc:proto:field=google.cloud.dataplex.v1.Lake.labels
	Labels map[string]string `json:"labels,omitempty"`

	// Optional. Description of the lake.
	// +kcc:proto:field=google.cloud.dataplex.v1.Lake.description
	Description *string `json:"description,omitempty"`

	// Optional. Settings to manage lake and Dataproc Metastore service instance
	//  association.
	// +kcc:proto:field=google.cloud.dataplex.v1.Lake.metastore
	Metastore *Lake_Metastore `json:"metastore,omitempty"`
}

// +kcc:proto=google.cloud.dataplex.v1.Lake.MetastoreStatus
type Lake_MetastoreStatus struct {
	// Current state of association.
	// +kcc:proto:field=google.cloud.dataplex.v1.Lake.MetastoreStatus.state
	State *string `json:"state,omitempty"`

	// Additional information about the current status.
	// +kcc:proto:field=google.cloud.dataplex.v1.Lake.MetastoreStatus.message
	Message *string `json:"message,omitempty"`

	// Last update time of the metastore status of the lake.
	// +kcc:proto:field=google.cloud.dataplex.v1.Lake.MetastoreStatus.update_time
	UpdateTime *string `json:"updateTime,omitempty"`

	// The URI of the endpoint used to access the Metastore service.
	// +kcc:proto:field=google.cloud.dataplex.v1.Lake.MetastoreStatus.endpoint
	Endpoint *string `json:"endpoint,omitempty"`
}

// +kcc:proto=google.cloud.dataplex.v1.Lake
type LakeObservedState struct {
	// Output only. The relative resource name of the lake, of the form:
	//  `projects/{project_number}/locations/{location_id}/lakes/{lake_id}`.
	// +kcc:proto:field=google.cloud.dataplex.v1.Lake.name
	Name *string `json:"name,omitempty"`

	// Output only. System generated globally unique ID for the lake. This ID will
	//  be different if the lake is deleted and re-created with the same name.
	// +kcc:proto:field=google.cloud.dataplex.v1.Lake.uid
	Uid *string `json:"uid,omitempty"`

	// Output only. The time when the lake was created.
	// +kcc:proto:field=google.cloud.dataplex.v1.Lake.create_time
	CreateTime *string `json:"createTime,omitempty"`

	// Output only. The time when the lake was last updated.
	// +kcc:proto:field=google.cloud.dataplex.v1.Lake.update_time
	UpdateTime *string `json:"updateTime,omitempty"`

	// Output only. Current state of the lake.
	// +kcc:proto:field=google.cloud.dataplex.v1.Lake.state
	State *string `json:"state,omitempty"`

	// Output only. Service account associated with this lake. This service
	//  account must be authorized to access or operate on resources managed by the
	//  lake.
	// +kcc:proto:field=google.cloud.dataplex.v1.Lake.service_account
	ServiceAccount *string `json:"serviceAccount,omitempty"`

	// Output only. Aggregated status of the underlying assets of the lake.
	// +kcc:proto:field=google.cloud.dataplex.v1.Lake.asset_status
	AssetStatus *AssetStatus `json:"assetStatus,omitempty"`

	// Output only. Metastore status of the lake.
	// +kcc:proto:field=google.cloud.dataplex.v1.Lake.metastore_status
	MetastoreStatus *Lake_MetastoreStatus `json:"metastoreStatus,omitempty"`
}

// +kcc:proto=google.cloud.dataplex.v1.Content
type Content struct {

	// Required. The path for the Content file, represented as directory
	//  structure. Unique within a lake. Limited to alphanumerics, hyphens,
	//  underscores, dots and slashes.
	// +kcc:proto:field=google.cloud.dataplex.v1.Content.path
	Path *string `json:"path,omitempty"`

	// Optional. User defined labels for the content.
	// +kcc:proto:field=google.cloud.dataplex.v1.Content.labels
	Labels map[string]string `json:"labels,omitempty"`

	// Optional. Description of the content.
	// +kcc:proto:field=google.cloud.dataplex.v1.Content.description
	Description *string `json:"description,omitempty"`

	// Required. Content data in string format.
	// +kcc:proto:field=google.cloud.dataplex.v1.Content.data_text
	DataText *string `json:"dataText,omitempty"`

	// Sql Script related configurations.
	// +kcc:proto:field=google.cloud.dataplex.v1.Content.sql_script
	SQLScript *Content_SQLScript `json:"sqlScript,omitempty"`

	// Notebook related configurations.
	// +kcc:proto:field=google.cloud.dataplex.v1.Content.notebook
	Notebook *Content_Notebook `json:"notebook,omitempty"`
}

// +kcc:proto=google.cloud.dataplex.v1.Content.Notebook
type Content_Notebook struct {
	// Required. Kernel Type of the notebook.
	// +kcc:proto:field=google.cloud.dataplex.v1.Content.Notebook.kernel_type
	KernelType *string `json:"kernelType,omitempty"`
}

// +kcc:proto=google.cloud.dataplex.v1.Content.SqlScript
type Content_SQLScript struct {
	// Required. Query Engine to be used for the Sql Query.
	// +kcc:proto:field=google.cloud.dataplex.v1.Content.SqlScript.engine
	Engine *string `json:"engine,omitempty"`
}

// +kcc:proto=google.cloud.dataplex.v1.Content
type ContentObservedState struct {
	// Output only. The relative resource name of the content, of the form:
	//  projects/{project_id}/locations/{location_id}/lakes/{lake_id}/content/{content_id}
	// +kcc:proto:field=google.cloud.dataplex.v1.Content.name
	Name *string `json:"name,omitempty"`

	// Output only. System generated globally unique ID for the content. This ID
	//  will be different if the content is deleted and re-created with the same
	//  name.
	// +kcc:proto:field=google.cloud.dataplex.v1.Content.uid
	Uid *string `json:"uid,omitempty"`

	// Output only. Content creation time.
	// +kcc:proto:field=google.cloud.dataplex.v1.Content.create_time
	CreateTime *string `json:"createTime,omitempty"`

	// Output only. The time when the content was last updated.
	// +kcc:proto:field=google.cloud.dataplex.v1.Content.update_time
	UpdateTime *string `json:"updateTime,omitempty"`
}

// +kcc:proto=google.cloud.dataplex.v1.Environment
type Environment struct {

	// Optional. User friendly display name.
	// +kcc:proto:field=google.cloud.dataplex.v1.Environment.display_name
	DisplayName *string `json:"displayName,omitempty"`

	// Optional. User defined labels for the environment.
	// +kcc:proto:field=google.cloud.dataplex.v1.Environment.labels
	Labels map[string]string `json:"labels,omitempty"`

	// Optional. Description of the environment.
	// +kcc:proto:field=google.cloud.dataplex.v1.Environment.description
	Description *string `json:"description,omitempty"`

	// Required. Infrastructure specification for the Environment.
	// +kcc:proto:field=google.cloud.dataplex.v1.Environment.infrastructure_spec
	InfrastructureSpec *Environment_InfrastructureSpec `json:"infrastructureSpec,omitempty"`

	// Optional. Configuration for sessions created for this environment.
	// +kcc:proto:field=google.cloud.dataplex.v1.Environment.session_spec
	SessionSpec *Environment_SessionSpec `json:"sessionSpec,omitempty"`
}

// +kcc:proto=google.cloud.dataplex.v1.Environment.Endpoints
type Environment_Endpoints struct {
}

// +kcc:proto=google.cloud.dataplex.v1.Environment.InfrastructureSpec
type Environment_InfrastructureSpec struct {
	// Optional. Compute resources needed for analyze interactive workloads.
	// +kcc:proto:field=google.cloud.dataplex.v1.Environment.InfrastructureSpec.compute
	Compute *Environment_InfrastructureSpec_ComputeResources `json:"compute,omitempty"`

	// Required. Software Runtime Configuration for analyze interactive
	//  workloads.
	// +kcc:proto:field=google.cloud.dataplex.v1.Environment.InfrastructureSpec.os_image
	OSImage *Environment_InfrastructureSpec_OSImageRuntime `json:"osImage,omitempty"`
}

// +kcc:proto=google.cloud.dataplex.v1.Environment.InfrastructureSpec.ComputeResources
type Environment_InfrastructureSpec_ComputeResources struct {
	// Optional. Size in GB of the disk. Default is 100 GB.
	// +kcc:proto:field=google.cloud.dataplex.v1.Environment.InfrastructureSpec.ComputeResources.disk_size_gb
	DiskSizeGB *int32 `json:"diskSizeGB,omitempty"`

	// Optional. Total number of nodes in the sessions created for this
	//  environment.
	// +kcc:proto:field=google.cloud.dataplex.v1.Environment.InfrastructureSpec.ComputeResources.node_count
	NodeCount *int32 `json:"nodeCount,omitempty"`

	// Optional. Max configurable nodes.
	//  If max_node_count > node_count, then auto-scaling is enabled.
	// +kcc:proto:field=google.cloud.dataplex.v1.Environment.InfrastructureSpec.ComputeResources.max_node_count
	MaxNodeCount *int32 `json:"maxNodeCount,omitempty"`
}

// +kcc:proto=google.cloud.dataplex.v1.Environment.InfrastructureSpec.OsImageRuntime
type Environment_InfrastructureSpec_OSImageRuntime struct {
	// Required. Dataplex Image version.
	// +kcc:proto:field=google.cloud.dataplex.v1.Environment.InfrastructureSpec.OsImageRuntime.image_version
	ImageVersion *string `json:"imageVersion,omitempty"`

	// Optional. List of Java jars to be included in the runtime environment.
	//  Valid input includes Cloud Storage URIs to Jar binaries.
	//  For example, gs://bucket-name/my/path/to/file.jar
	// +kcc:proto:field=google.cloud.dataplex.v1.Environment.InfrastructureSpec.OsImageRuntime.java_libraries
	JavaLibraries []string `json:"javaLibraries,omitempty"`

	// Optional. A list of python packages to be installed.
	//  Valid formats include Cloud Storage URI to a PIP installable library.
	//  For example, gs://bucket-name/my/path/to/lib.tar.gz
	// +kcc:proto:field=google.cloud.dataplex.v1.Environment.InfrastructureSpec.OsImageRuntime.python_packages
	PythonPackages []string `json:"pythonPackages,omitempty"`

	// Optional. Spark properties to provide configuration for use in sessions
	//  created for this environment. The properties to set on daemon config
	//  files. Property keys are specified in `prefix:property` format. The
	//  prefix must be "spark".
	// +kcc:proto:field=google.cloud.dataplex.v1.Environment.InfrastructureSpec.OsImageRuntime.properties
	Properties map[string]string `json:"properties,omitempty"`
}

// +kcc:proto=google.cloud.dataplex.v1.Environment.SessionSpec
type Environment_SessionSpec struct {
	// Optional. The idle time configuration of the session. The session will be
	//  auto-terminated at the end of this period.
	// +kcc:proto:field=google.cloud.dataplex.v1.Environment.SessionSpec.max_idle_duration
	MaxIdleDuration *string `json:"maxIdleDuration,omitempty"`

	// Optional. If True, this causes sessions to be pre-created and available
	//  for faster startup to enable interactive exploration use-cases. This
	//  defaults to False to avoid additional billed charges. These can only be
	//  set to True for the environment with name set to "default", and with
	//  default configuration.
	// +kcc:proto:field=google.cloud.dataplex.v1.Environment.SessionSpec.enable_fast_startup
	EnableFastStartup *bool `json:"enableFastStartup,omitempty"`
}

// +kcc:proto=google.cloud.dataplex.v1.Environment.SessionStatus
type Environment_SessionStatus struct {
}

// +kcc:proto=google.cloud.dataplex.v1.Environment
type EnvironmentObservedState struct {
	// Output only. The relative resource name of the environment, of the form:
	//  projects/{project_id}/locations/{location_id}/lakes/{lake_id}/environment/{environment_id}
	// +kcc:proto:field=google.cloud.dataplex.v1.Environment.name
	Name *string `json:"name,omitempty"`

	// Output only. System generated globally unique ID for the environment. This
	//  ID will be different if the environment is deleted and re-created with the
	//  same name.
	// +kcc:proto:field=google.cloud.dataplex.v1.Environment.uid
	Uid *string `json:"uid,omitempty"`

	// Output only. Environment creation time.
	// +kcc:proto:field=google.cloud.dataplex.v1.Environment.create_time
	CreateTime *string `json:"createTime,omitempty"`

	// Output only. The time when the environment was last updated.
	// +kcc:proto:field=google.cloud.dataplex.v1.Environment.update_time
	UpdateTime *string `json:"updateTime,omitempty"`

	// Output only. Current state of the environment.
	// +kcc:proto:field=google.cloud.dataplex.v1.Environment.state
	State *string `json:"state,omitempty"`

	// Output only. Status of sessions created for this environment.
	// +kcc:proto:field=google.cloud.dataplex.v1.Environment.session_status
	SessionStatus *Environment_SessionStatusObservedState `json:"sessionStatus,omitempty"`

	// Output only. URI Endpoints to access sessions associated with the
	//  Environment.
	// +kcc:proto:field=google.cloud.dataplex.v1.Environment.endpoints
	Endpoints *Environment_EndpointsObservedState `json:"endpoints,omitempty"`
}

// +kcc:proto=google.cloud.dataplex.v1.Environment.Endpoints
type Environment_EndpointsObservedState struct {
	// Output only. URI to serve notebook APIs
	// +kcc:proto:field=google.cloud.dataplex.v1.Environment.Endpoints.notebooks
	Notebooks *string `json:"notebooks,omitempty"`

	// Output only. URI to serve SQL APIs
	// +kcc:proto:field=google.cloud.dataplex.v1.Environment.Endpoints.sql
	SQL *string `json:"sql,omitempty"`
}

// +kcc:proto=google.cloud.dataplex.v1.Environment.SessionStatus
type Environment_SessionStatusObservedState struct {
	// Output only. Queries over sessions to mark whether the environment is
	//  currently active or not
	// +kcc:proto:field=google.cloud.dataplex.v1.Environment.SessionStatus.active
	Active *bool `json:"active,omitempty"`
}

// +kcc:proto=google.cloud.dataplex.v1.Task.ExecutionSpec
type Task_ExecutionSpec struct {
	// Optional. The arguments to pass to the task.
	//  The args can use placeholders of the format ${placeholder} as
	//  part of key/value string. These will be interpolated before passing the
	//  args to the driver. Currently supported placeholders:
	//  - ${task_id}
	//  - ${job_time}
	//  To pass positional args, set the key as TASK_ARGS. The value should be a
	//  comma-separated string of all the positional arguments. To use a
	//  delimiter other than comma, refer to
	//  https://cloud.google.com/sdk/gcloud/reference/topic/escaping. In case of
	//  other keys being present in the args, then TASK_ARGS will be passed as
	//  the last argument.
	// +kcc:proto:field=google.cloud.dataplex.v1.Task.ExecutionSpec.args
	Args map[string]string `json:"args,omitempty"`

	// Required. Service account to use to execute a task.
	//  If not provided, the default Compute service account for the project is
	//  used.
	// +required
	// +kcc:proto:field=google.cloud.dataplex.v1.Task.ExecutionSpec.service_account
	ServiceAccount *string `json:"serviceAccount,omitempty"`

	// Optional. The project in which jobs are run. By default, the project
	//  containing the Lake is used. If a project is provided, the
	//  [ExecutionSpec.service_account][google.cloud.dataplex.v1.Task.ExecutionSpec.service_account]
	//  must belong to this project.
	// +kcc:proto:field=google.cloud.dataplex.v1.Task.ExecutionSpec.project
	Project *string `json:"project,omitempty"`

	// Optional. The maximum duration after which the job execution is expired.
	// +kcc:proto:field=google.cloud.dataplex.v1.Task.ExecutionSpec.max_job_execution_lifetime
	MaxJobExecutionLifetime *string `json:"maxJobExecutionLifetime,omitempty"`

	// Optional. The Cloud KMS key to use for encryption, of the form:
	//  `projects/{project_number}/locations/{location_id}/keyRings/{key-ring-name}/cryptoKeys/{key-name}`.
	// +kcc:proto:field=google.cloud.dataplex.v1.Task.ExecutionSpec.kms_key
	KMSKeyRef *refsv1beta1.KMSCryptoKeyRef `json:"kmsKeyRef,omitempty"`
}

// +kcc:proto=google.cloud.dataplex.v1.Task.InfrastructureSpec
type Task_InfrastructureSpec struct {
	// Compute resources needed for a Task when using Dataproc Serverless.
	// +kcc:proto:field=google.cloud.dataplex.v1.Task.InfrastructureSpec.batch
	Batch *Task_InfrastructureSpec_BatchComputeResources `json:"batch,omitempty"`

	// Container Image Runtime Configuration.
	// +kcc:proto:field=google.cloud.dataplex.v1.Task.InfrastructureSpec.container_image
	ContainerImage *Task_InfrastructureSpec_ContainerImageRuntime `json:"containerImage,omitempty"`

	// Vpc network.
	// +kcc:proto:field=google.cloud.dataplex.v1.Task.InfrastructureSpec.vpc_network
	VpcNetwork *Task_InfrastructureSpec_VpcNetwork `json:"vpcNetwork,omitempty"`
}

// +kcc:proto=google.cloud.dataplex.v1.Task.InfrastructureSpec.BatchComputeResources
type Task_InfrastructureSpec_BatchComputeResources struct {
	// Optional. Total number of job executors.
	//  Executor Count should be between 2 and 100. [Default=2]
	// +kcc:proto:field=google.cloud.dataplex.v1.Task.InfrastructureSpec.BatchComputeResources.executors_count
	ExecutorsCount *int32 `json:"executorsCount,omitempty"`

	// Optional. Max configurable executors.
	//  If max_executors_count > executors_count, then auto-scaling is enabled.
	//  Max Executor Count should be between 2 and 1000. [Default=1000]
	// +kcc:proto:field=google.cloud.dataplex.v1.Task.InfrastructureSpec.BatchComputeResources.max_executors_count
	MaxExecutorsCount *int32 `json:"maxExecutorsCount,omitempty"`
}

// +kcc:proto=google.cloud.dataplex.v1.Task.InfrastructureSpec.ContainerImageRuntime
type Task_InfrastructureSpec_ContainerImageRuntime struct {
	// Optional. Container image to use.
	// +kcc:proto:field=google.cloud.dataplex.v1.Task.InfrastructureSpec.ContainerImageRuntime.image
	Image *string `json:"image,omitempty"`

	// Optional. A list of Java JARS to add to the classpath.
	//  Valid input includes Cloud Storage URIs to Jar binaries.
	//  For example, gs://bucket-name/my/path/to/file.jar
	// +kcc:proto:field=google.cloud.dataplex.v1.Task.InfrastructureSpec.ContainerImageRuntime.java_jars
	JavaJars []string `json:"javaJars,omitempty"`

	// Optional. A list of python packages to be installed.
	//  Valid formats include Cloud Storage URI to a PIP installable library.
	//  For example, gs://bucket-name/my/path/to/lib.tar.gz
	// +kcc:proto:field=google.cloud.dataplex.v1.Task.InfrastructureSpec.ContainerImageRuntime.python_packages
	PythonPackages []string `json:"pythonPackages,omitempty"`

	// Optional. Override to common configuration of open source components
	//  installed on the Dataproc cluster. The properties to set on daemon
	//  config files. Property keys are specified in `prefix:property` format,
	//  for example `core:hadoop.tmp.dir`. For more information, see [Cluster
	//  properties](https://cloud.google.com/dataproc/docs/concepts/cluster-properties).
	// +kcc:proto:field=google.cloud.dataplex.v1.Task.InfrastructureSpec.ContainerImageRuntime.properties
	Properties map[string]string `json:"properties,omitempty"`
}

// +kcc:proto=google.cloud.dataplex.v1.Task.InfrastructureSpec.VpcNetwork
type Task_InfrastructureSpec_VpcNetwork struct {
	// Optional. The Cloud VPC network in which the job is run. By default,
	//  the Cloud VPC network named Default within the project is used.
	// +kcc:proto:field=google.cloud.dataplex.v1.Task.InfrastructureSpec.VpcNetwork.network
	Network *string `json:"network,omitempty"`

	// Optional. The Cloud VPC sub-network in which the job is run.
	// +kcc:proto:field=google.cloud.dataplex.v1.Task.InfrastructureSpec.VpcNetwork.sub_network
	SubNetwork *string `json:"subNetwork,omitempty"`

	// Optional. List of network tags to apply to the job.
	// +kcc:proto:field=google.cloud.dataplex.v1.Task.InfrastructureSpec.VpcNetwork.network_tags
	NetworkTags []string `json:"networkTags,omitempty"`
}

// +kcc:proto=google.cloud.dataplex.v1.Task.NotebookTaskConfig
type Task_NotebookTaskConfig struct {
	// Required. Path to input notebook. This can be the Cloud Storage URI of
	//  the notebook file or the path to a Notebook Content. The execution args
	//  are accessible as environment variables
	//  (`TASK_key=value`).
	// +required
	// +kcc:proto:field=google.cloud.dataplex.v1.Task.NotebookTaskConfig.notebook
	Notebook *string `json:"notebook,omitempty"`

	// Optional. Infrastructure specification for the execution.
	// +kcc:proto:field=google.cloud.dataplex.v1.Task.NotebookTaskConfig.infrastructure_spec
	InfrastructureSpec *Task_InfrastructureSpec `json:"infrastructureSpec,omitempty"`

	// Optional. Cloud Storage URIs of files to be placed in the working
	//  directory of each executor.
	// +kcc:proto:field=google.cloud.dataplex.v1.Task.NotebookTaskConfig.file_uris
	FileUris []string `json:"fileUris,omitempty"`

	// Optional. Cloud Storage URIs of archives to be extracted into the working
	//  directory of each executor. Supported file types: .jar, .tar, .tar.gz,
	//  .tgz, and .zip.
	// +kcc:proto:field=google.cloud.dataplex.v1.Task.NotebookTaskConfig.archive_uris
	ArchiveUris []string `json:"archiveUris,omitempty"`
}

// +kcc:proto=google.cloud.dataplex.v1.Task.SparkTaskConfig
type Task_SparkTaskConfig struct {
	// The Cloud Storage URI of the jar file that contains the main class.
	//  The execution args are passed in as a sequence of named process
	//  arguments (`--key=value`).
	// +kcc:proto:field=google.cloud.dataplex.v1.Task.SparkTaskConfig.main_jar_file_uri
	MainJarFileURI *string `json:"mainJarFileURI,omitempty"`

	// The name of the driver's main class. The jar file that contains the
	//  class must be in the default CLASSPATH or specified in
	//  `jar_file_uris`.
	//  The execution args are passed in as a sequence of named process
	//  arguments (`--key=value`).
	// +kcc:proto:field=google.cloud.dataplex.v1.Task.SparkTaskConfig.main_class
	MainClass *string `json:"mainClass,omitempty"`

	// The Gcloud Storage URI of the main Python file to use as the driver.
	//  Must be a .py file. The execution args are passed in as a sequence of
	//  named process arguments (`--key=value`).
	// +kcc:proto:field=google.cloud.dataplex.v1.Task.SparkTaskConfig.python_script_file
	PythonScriptFile *string `json:"pythonScriptFile,omitempty"`

	// A reference to a query file. This should be the Cloud Storage URI of
	//  the query file. The execution args are used to declare a set of script
	//  variables (`set key="value";`).
	// +kcc:proto:field=google.cloud.dataplex.v1.Task.SparkTaskConfig.sql_script_file
	SQLScriptFile *string `json:"sqlScriptFile,omitempty"`

	// The query text.
	//  The execution args are used to declare a set of script variables
	//  (`set key="value";`).
	// +kcc:proto:field=google.cloud.dataplex.v1.Task.SparkTaskConfig.sql_script
	SQLScript *string `json:"sqlScript,omitempty"`

	// Optional. Cloud Storage URIs of files to be placed in the working
	//  directory of each executor.
	// +kcc:proto:field=google.cloud.dataplex.v1.Task.SparkTaskConfig.file_uris
	FileUris []string `json:"fileUris,omitempty"`

	// Optional. Cloud Storage URIs of archives to be extracted into the working
	//  directory of each executor. Supported file types: .jar, .tar, .tar.gz,
	//  .tgz, and .zip.
	// +kcc:proto:field=google.cloud.dataplex.v1.Task.SparkTaskConfig.archive_uris
	ArchiveUris []string `json:"archiveUris,omitempty"`

	// Optional. Infrastructure specification for the execution.
	// +kcc:proto:field=google.cloud.dataplex.v1.Task.SparkTaskConfig.infrastructure_spec
	InfrastructureSpec *Task_InfrastructureSpec `json:"infrastructureSpec,omitempty"`
}

// +kcc:proto=google.cloud.dataplex.v1.Task.TriggerSpec
type Task_TriggerSpec struct {
	// Required. Immutable. Trigger type of the user-specified Task.
	// +required
	// +kcc:proto:field=google.cloud.dataplex.v1.Task.TriggerSpec.type
	Type *string `json:"type,omitempty"`

	// Optional. The first run of the task will be after this time.
	//  If not specified, the task will run shortly after being submitted if
	//  ON_DEMAND and based on the schedule if RECURRING.
	// +kcc:proto:field=google.cloud.dataplex.v1.Task.TriggerSpec.start_time
	StartTime *string `json:"startTime,omitempty"`

	// Optional. Prevent the task from executing.
	//  This does not cancel already running tasks. It is intended to temporarily
	//  disable RECURRING tasks.
	// +kcc:proto:field=google.cloud.dataplex.v1.Task.TriggerSpec.disabled
	Disabled *bool `json:"disabled,omitempty"`

	// Optional. Number of retry attempts before aborting.
	//  Set to zero to never attempt to retry a failed task.
	// +kcc:proto:field=google.cloud.dataplex.v1.Task.TriggerSpec.max_retries
	MaxRetries *int32 `json:"maxRetries,omitempty"`

	// Optional. Cron schedule (https://en.wikipedia.org/wiki/Cron) for
	//  running tasks periodically. To explicitly set a timezone to the cron
	//  tab, apply a prefix in the cron tab: "CRON_TZ=${IANA_TIME_ZONE}" or
	//  "TZ=${IANA_TIME_ZONE}". The ${IANA_TIME_ZONE} may only be a valid
	//  string from IANA time zone database. For example,
	//  `CRON_TZ=America/New_York 1 * * * *`, or `TZ=America/New_York 1 * * *
	//  *`. This field is required for RECURRING tasks.
	// +kcc:proto:field=google.cloud.dataplex.v1.Task.TriggerSpec.schedule
	Schedule *string `json:"schedule,omitempty"`
}

// +kcc:proto=google.cloud.dataplex.v1.Job
type JobObservedState struct {
	// Output only. The relative resource name of the job, of the form:
	//  `projects/{project_number}/locations/{location_id}/lakes/{lake_id}/tasks/{task_id}/jobs/{job_id}`.
	// +kcc:proto:field=google.cloud.dataplex.v1.Job.name
	Name *string `json:"name,omitempty"`

	// Output only. System generated globally unique ID for the job.
	// +kcc:proto:field=google.cloud.dataplex.v1.Job.uid
	Uid *string `json:"uid,omitempty"`

	// Output only. The time when the job was started.
	// +kcc:proto:field=google.cloud.dataplex.v1.Job.start_time
	StartTime *string `json:"startTime,omitempty"`

	// Output only. The time when the job ended.
	// +kcc:proto:field=google.cloud.dataplex.v1.Job.end_time
	EndTime *string `json:"endTime,omitempty"`

	// Output only. Execution state for the job.
	// +kcc:proto:field=google.cloud.dataplex.v1.Job.state
	State *string `json:"state,omitempty"`

	// Output only. The number of times the job has been retried (excluding the
	//  initial attempt).
	// +kcc:proto:field=google.cloud.dataplex.v1.Job.retry_count
	RetryCount *uint32 `json:"retryCount,omitempty"`

	// Output only. The underlying service running a job.
	// +kcc:proto:field=google.cloud.dataplex.v1.Job.service
	Service *string `json:"service,omitempty"`

	// Output only. The full resource name for the job run under a particular
	//  service.
	// +kcc:proto:field=google.cloud.dataplex.v1.Job.service_job
	ServiceJob *string `json:"serviceJob,omitempty"`

	// Output only. Additional information about the current state.
	// +kcc:proto:field=google.cloud.dataplex.v1.Job.message
	Message *string `json:"message,omitempty"`

	// Output only. User-defined labels for the task.
	// +kcc:proto:field=google.cloud.dataplex.v1.Job.labels
	Labels map[string]string `json:"labels,omitempty"`

	// Output only. Job execution trigger.
	// +kcc:proto:field=google.cloud.dataplex.v1.Job.trigger
	Trigger *string `json:"trigger,omitempty"`

	// Output only. Spec related to how a task is executed.
	// +kcc:proto:field=google.cloud.dataplex.v1.Job.execution_spec
	ExecutionSpec *Task_ExecutionSpec `json:"executionSpec,omitempty"`
}

// +kcc:proto=google.cloud.dataplex.v1.Task.ExecutionStatus
type Task_ExecutionStatusObservedState struct {
	// Output only. Last update time of the status.
	// +kcc:proto:field=google.cloud.dataplex.v1.Task.ExecutionStatus.update_time
	UpdateTime *string `json:"updateTime,omitempty"`

	// Output only. latest job execution
	// +kcc:proto:field=google.cloud.dataplex.v1.Task.ExecutionStatus.latest_job
	LatestJob *JobObservedState `json:"latestJob,omitempty"`
}
