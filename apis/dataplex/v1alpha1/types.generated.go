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
	OsImage *Environment_InfrastructureSpec_OsImageRuntime `json:"osImage,omitempty"`
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
type Environment_InfrastructureSpec_OsImageRuntime struct {
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
	SessionStatus *Environment_SessionStatus `json:"sessionStatus,omitempty"`

	// Output only. URI Endpoints to access sessions associated with the
	//  Environment.
	// +kcc:proto:field=google.cloud.dataplex.v1.Environment.endpoints
	Endpoints *Environment_Endpoints `json:"endpoints,omitempty"`
}

// +kcc:proto=google.cloud.dataplex.v1.Environment.Endpoints
type Environment_EndpointsObservedState struct {
	// Output only. URI to serve notebook APIs
	// +kcc:proto:field=google.cloud.dataplex.v1.Environment.Endpoints.notebooks
	Notebooks *string `json:"notebooks,omitempty"`

	// Output only. URI to serve SQL APIs
	// +kcc:proto:field=google.cloud.dataplex.v1.Environment.Endpoints.sql
	Sql *string `json:"sql,omitempty"`
}

// +kcc:proto=google.cloud.dataplex.v1.Environment.SessionStatus
type Environment_SessionStatusObservedState struct {
	// Output only. Queries over sessions to mark whether the environment is
	//  currently active or not
	// +kcc:proto:field=google.cloud.dataplex.v1.Environment.SessionStatus.active
	Active *bool `json:"active,omitempty"`
}
