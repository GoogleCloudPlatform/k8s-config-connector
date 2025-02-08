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


// +kcc:proto=google.cloud.run.v2.CloudSqlInstance
type CloudSqlInstance struct {
	// The Cloud SQL instance connection names, as can be found in
	//  https://console.cloud.google.com/sql/instances. Visit
	//  https://cloud.google.com/sql/docs/mysql/connect-run for more information on
	//  how to connect Cloud SQL and Cloud Run. Format:
	//  {project}:{location}:{instance}
	// +kcc:proto:field=google.cloud.run.v2.CloudSqlInstance.instances
	Instances []string `json:"instances,omitempty"`
}

// +kcc:proto=google.cloud.run.v2.Condition
type Condition struct {
	// type is used to communicate the status of the reconciliation process.
	//  See also:
	//  https://github.com/knative/serving/blob/main/docs/spec/errors.md#error-conditions-and-reporting
	//  Types common to all resources include:
	//  * "Ready": True when the Resource is ready.
	// +kcc:proto:field=google.cloud.run.v2.Condition.type
	Type *string `json:"type,omitempty"`

	// State of the condition.
	// +kcc:proto:field=google.cloud.run.v2.Condition.state
	State *string `json:"state,omitempty"`

	// Human readable message indicating details about the current status.
	// +kcc:proto:field=google.cloud.run.v2.Condition.message
	Message *string `json:"message,omitempty"`

	// Last time the condition transitioned from one status to another.
	// +kcc:proto:field=google.cloud.run.v2.Condition.last_transition_time
	LastTransitionTime *string `json:"lastTransitionTime,omitempty"`

	// How to interpret failures of this condition, one of Error, Warning, Info
	// +kcc:proto:field=google.cloud.run.v2.Condition.severity
	Severity *string `json:"severity,omitempty"`
}

// +kcc:proto=google.cloud.run.v2.Container
type Container struct {
	// Name of the container specified as a DNS_LABEL (RFC 1123).
	// +kcc:proto:field=google.cloud.run.v2.Container.name
	Name *string `json:"name,omitempty"`

	// Required. Name of the container image in Dockerhub, Google Artifact
	//  Registry, or Google Container Registry. If the host is not provided,
	//  Dockerhub is assumed.
	// +kcc:proto:field=google.cloud.run.v2.Container.image
	Image *string `json:"image,omitempty"`

	// Entrypoint array. Not executed within a shell.
	//  The docker image's ENTRYPOINT is used if this is not provided.
	// +kcc:proto:field=google.cloud.run.v2.Container.command
	Command []string `json:"command,omitempty"`

	// Arguments to the entrypoint.
	//  The docker image's CMD is used if this is not provided.
	// +kcc:proto:field=google.cloud.run.v2.Container.args
	Args []string `json:"args,omitempty"`

	// List of environment variables to set in the container.
	// +kcc:proto:field=google.cloud.run.v2.Container.env
	Env []EnvVar `json:"env,omitempty"`

	// Compute Resource requirements by this container.
	// +kcc:proto:field=google.cloud.run.v2.Container.resources
	Resources *ResourceRequirements `json:"resources,omitempty"`

	// List of ports to expose from the container. Only a single port can be
	//  specified. The specified ports must be listening on all interfaces
	//  (0.0.0.0) within the container to be accessible.
	//
	//  If omitted, a port number will be chosen and passed to the container
	//  through the PORT environment variable for the container to listen on.
	// +kcc:proto:field=google.cloud.run.v2.Container.ports
	Ports []ContainerPort `json:"ports,omitempty"`

	// Volume to mount into the container's filesystem.
	// +kcc:proto:field=google.cloud.run.v2.Container.volume_mounts
	VolumeMounts []VolumeMount `json:"volumeMounts,omitempty"`

	// Container's working directory.
	//  If not specified, the container runtime's default will be used, which
	//  might be configured in the container image.
	// +kcc:proto:field=google.cloud.run.v2.Container.working_dir
	WorkingDir *string `json:"workingDir,omitempty"`

	// Periodic probe of container liveness.
	//  Container will be restarted if the probe fails.
	// +kcc:proto:field=google.cloud.run.v2.Container.liveness_probe
	LivenessProbe *Probe `json:"livenessProbe,omitempty"`

	// Startup probe of application within the container.
	//  All other probes are disabled if a startup probe is provided, until it
	//  succeeds. Container will not be added to service endpoints if the probe
	//  fails.
	// +kcc:proto:field=google.cloud.run.v2.Container.startup_probe
	StartupProbe *Probe `json:"startupProbe,omitempty"`

	// Names of the containers that must start before this container.
	// +kcc:proto:field=google.cloud.run.v2.Container.depends_on
	DependsOn []string `json:"dependsOn,omitempty"`
}

// +kcc:proto=google.cloud.run.v2.ContainerPort
type ContainerPort struct {
	// If specified, used to specify which protocol to use.
	//  Allowed values are "http1" and "h2c".
	// +kcc:proto:field=google.cloud.run.v2.ContainerPort.name
	Name *string `json:"name,omitempty"`

	// Port number the container listens on.
	//  This must be a valid TCP port number, 0 < container_port < 65536.
	// +kcc:proto:field=google.cloud.run.v2.ContainerPort.container_port
	ContainerPort *int32 `json:"containerPort,omitempty"`
}

// +kcc:proto=google.cloud.run.v2.EmptyDirVolumeSource
type EmptyDirVolumeSource struct {
	// The medium on which the data is stored. Acceptable values today is only
	//  MEMORY or none. When none, the default will currently be backed by memory
	//  but could change over time. +optional
	// +kcc:proto:field=google.cloud.run.v2.EmptyDirVolumeSource.medium
	Medium *string `json:"medium,omitempty"`

	// Limit on the storage usable by this EmptyDir volume.
	//  The size limit is also applicable for memory medium.
	//  The maximum usage on memory medium EmptyDir would be the minimum value
	//  between the SizeLimit specified here and the sum of memory limits of all
	//  containers. The default is nil which means that the limit is undefined.
	//  More info:
	//  https://cloud.google.com/run/docs/configuring/in-memory-volumes#configure-volume.
	//  Info in Kubernetes:
	//  https://kubernetes.io/docs/concepts/storage/volumes/#emptydir
	// +kcc:proto:field=google.cloud.run.v2.EmptyDirVolumeSource.size_limit
	SizeLimit *string `json:"sizeLimit,omitempty"`
}

// +kcc:proto=google.cloud.run.v2.EnvVar
type EnvVar struct {
	// Required. Name of the environment variable. Must not exceed 32768
	//  characters.
	// +kcc:proto:field=google.cloud.run.v2.EnvVar.name
	Name *string `json:"name,omitempty"`

	// Literal value of the environment variable.
	//  Defaults to "", and the maximum length is 32768 bytes.
	//  Variable references are not supported in Cloud Run.
	// +kcc:proto:field=google.cloud.run.v2.EnvVar.value
	Value *string `json:"value,omitempty"`

	// Source for the environment variable's value.
	// +kcc:proto:field=google.cloud.run.v2.EnvVar.value_source
	ValueSource *EnvVarSource `json:"valueSource,omitempty"`
}

// +kcc:proto=google.cloud.run.v2.EnvVarSource
type EnvVarSource struct {
	// Selects a secret and a specific version from Cloud Secret Manager.
	// +kcc:proto:field=google.cloud.run.v2.EnvVarSource.secret_key_ref
	SecretKeyRef *SecretKeySelector `json:"secretKeyRef,omitempty"`
}

// +kcc:proto=google.cloud.run.v2.Execution
type Execution struct {

	// The least stable launch stage needed to create this resource, as defined by
	//  [Google Cloud Platform Launch
	//  Stages](https://cloud.google.com/terms/launch-stages). Cloud Run supports
	//  `ALPHA`, `BETA`, and `GA`.
	//  <p>Note that this value might not be what was used
	//  as input. For example, if ALPHA was provided as input in the parent
	//  resource, but only BETA and GA-level features are were, this field will be
	//  BETA.
	// +kcc:proto:field=google.cloud.run.v2.Execution.launch_stage
	LaunchStage *string `json:"launchStage,omitempty"`
}

// +kcc:proto=google.cloud.run.v2.GCSVolumeSource
type GCSVolumeSource struct {
	// Cloud Storage Bucket name.
	// +kcc:proto:field=google.cloud.run.v2.GCSVolumeSource.bucket
	Bucket *string `json:"bucket,omitempty"`

	// If true, the volume will be mounted as read only for all mounts.
	// +kcc:proto:field=google.cloud.run.v2.GCSVolumeSource.read_only
	ReadOnly *bool `json:"readOnly,omitempty"`

	// A list of additional flags to pass to the gcsfuse CLI.
	//  Options should be specified without the leading "--".
	// +kcc:proto:field=google.cloud.run.v2.GCSVolumeSource.mount_options
	MountOptions []string `json:"mountOptions,omitempty"`
}

// +kcc:proto=google.cloud.run.v2.GRPCAction
type GRPCAction struct {
	// Optional. Port number of the gRPC service. Number must be in the range 1 to
	//  65535. If not specified, defaults to the exposed port of the container,
	//  which is the value of container.ports[0].containerPort.
	// +kcc:proto:field=google.cloud.run.v2.GRPCAction.port
	Port *int32 `json:"port,omitempty"`

	// Optional. Service is the name of the service to place in the gRPC
	//  HealthCheckRequest (see
	//  https://github.com/grpc/grpc/blob/master/doc/health-checking.md ). If this
	//  is not specified, the default behavior is defined by gRPC.
	// +kcc:proto:field=google.cloud.run.v2.GRPCAction.service
	Service *string `json:"service,omitempty"`
}

// +kcc:proto=google.cloud.run.v2.HTTPGetAction
type HTTPGetAction struct {
	// Optional. Path to access on the HTTP server. Defaults to '/'.
	// +kcc:proto:field=google.cloud.run.v2.HTTPGetAction.path
	Path *string `json:"path,omitempty"`

	// Optional. Custom headers to set in the request. HTTP allows repeated
	//  headers.
	// +kcc:proto:field=google.cloud.run.v2.HTTPGetAction.http_headers
	HTTPHeaders []HTTPHeader `json:"httpHeaders,omitempty"`

	// Optional. Port number to access on the container. Must be in the range 1 to
	//  65535. If not specified, defaults to the exposed port of the container,
	//  which is the value of container.ports[0].containerPort.
	// +kcc:proto:field=google.cloud.run.v2.HTTPGetAction.port
	Port *int32 `json:"port,omitempty"`
}

// +kcc:proto=google.cloud.run.v2.HTTPHeader
type HTTPHeader struct {
	// Required. The header field name
	// +kcc:proto:field=google.cloud.run.v2.HTTPHeader.name
	Name *string `json:"name,omitempty"`

	// Optional. The header field value
	// +kcc:proto:field=google.cloud.run.v2.HTTPHeader.value
	Value *string `json:"value,omitempty"`
}

// +kcc:proto=google.cloud.run.v2.NFSVolumeSource
type NFSVolumeSource struct {
	// Hostname or IP address of the NFS server
	// +kcc:proto:field=google.cloud.run.v2.NFSVolumeSource.server
	Server *string `json:"server,omitempty"`

	// Path that is exported by the NFS server.
	// +kcc:proto:field=google.cloud.run.v2.NFSVolumeSource.path
	Path *string `json:"path,omitempty"`

	// If true, the volume will be mounted as read only for all mounts.
	// +kcc:proto:field=google.cloud.run.v2.NFSVolumeSource.read_only
	ReadOnly *bool `json:"readOnly,omitempty"`
}

// +kcc:proto=google.cloud.run.v2.Probe
type Probe struct {
	// Optional. Number of seconds after the container has started before the
	//  probe is initiated. Defaults to 0 seconds. Minimum value is 0. Maximum
	//  value for liveness probe is 3600. Maximum value for startup probe is 240.
	// +kcc:proto:field=google.cloud.run.v2.Probe.initial_delay_seconds
	InitialDelaySeconds *int32 `json:"initialDelaySeconds,omitempty"`

	// Optional. Number of seconds after which the probe times out.
	//  Defaults to 1 second. Minimum value is 1. Maximum value is 3600.
	//  Must be smaller than period_seconds.
	// +kcc:proto:field=google.cloud.run.v2.Probe.timeout_seconds
	TimeoutSeconds *int32 `json:"timeoutSeconds,omitempty"`

	// Optional. How often (in seconds) to perform the probe.
	//  Default to 10 seconds. Minimum value is 1. Maximum value for liveness probe
	//  is 3600. Maximum value for startup probe is 240.
	//  Must be greater or equal than timeout_seconds.
	// +kcc:proto:field=google.cloud.run.v2.Probe.period_seconds
	PeriodSeconds *int32 `json:"periodSeconds,omitempty"`

	// Optional. Minimum consecutive failures for the probe to be considered
	//  failed after having succeeded. Defaults to 3. Minimum value is 1.
	// +kcc:proto:field=google.cloud.run.v2.Probe.failure_threshold
	FailureThreshold *int32 `json:"failureThreshold,omitempty"`

	// Optional. HTTPGet specifies the http request to perform.
	//  Exactly one of httpGet, tcpSocket, or grpc must be specified.
	// +kcc:proto:field=google.cloud.run.v2.Probe.http_get
	HTTPGet *HTTPGetAction `json:"httpGet,omitempty"`

	// Optional. TCPSocket specifies an action involving a TCP port.
	//  Exactly one of httpGet, tcpSocket, or grpc must be specified.
	// +kcc:proto:field=google.cloud.run.v2.Probe.tcp_socket
	TcpSocket *TCPSocketAction `json:"tcpSocket,omitempty"`

	// Optional. GRPC specifies an action involving a gRPC port.
	//  Exactly one of httpGet, tcpSocket, or grpc must be specified.
	// +kcc:proto:field=google.cloud.run.v2.Probe.grpc
	Grpc *GRPCAction `json:"grpc,omitempty"`
}

// +kcc:proto=google.cloud.run.v2.ResourceRequirements
type ResourceRequirements struct {
	// Only `memory` and `cpu` keys in the map are supported.
	//
	//  <p>Notes:
	//   * The only supported values for CPU are '1', '2', '4', and '8'. Setting 4
	//  CPU requires at least 2Gi of memory. For more information, go to
	//  https://cloud.google.com/run/docs/configuring/cpu.
	//    * For supported 'memory' values and syntax, go to
	//   https://cloud.google.com/run/docs/configuring/memory-limits
	// +kcc:proto:field=google.cloud.run.v2.ResourceRequirements.limits
	Limits map[string]string `json:"limits,omitempty"`

	// Determines whether CPU is only allocated during requests (true by default).
	//  However, if ResourceRequirements is set, the caller must explicitly
	//  set this field to true to preserve the default behavior.
	// +kcc:proto:field=google.cloud.run.v2.ResourceRequirements.cpu_idle
	CpuIdle *bool `json:"cpuIdle,omitempty"`

	// Determines whether CPU should be boosted on startup of a new container
	//  instance above the requested CPU threshold, this can help reduce cold-start
	//  latency.
	// +kcc:proto:field=google.cloud.run.v2.ResourceRequirements.startup_cpu_boost
	StartupCpuBoost *bool `json:"startupCpuBoost,omitempty"`
}

// +kcc:proto=google.cloud.run.v2.SecretKeySelector
type SecretKeySelector struct {
	// Required. The name of the secret in Cloud Secret Manager.
	//  Format: {secret_name} if the secret is in the same project.
	//  projects/{project}/secrets/{secret_name} if the secret is
	//  in a different project.
	// +kcc:proto:field=google.cloud.run.v2.SecretKeySelector.secret
	Secret *string `json:"secret,omitempty"`

	// The Cloud Secret Manager secret version.
	//  Can be 'latest' for the latest version, an integer for a specific version,
	//  or a version alias.
	// +kcc:proto:field=google.cloud.run.v2.SecretKeySelector.version
	Version *string `json:"version,omitempty"`
}

// +kcc:proto=google.cloud.run.v2.SecretVolumeSource
type SecretVolumeSource struct {
	// Required. The name of the secret in Cloud Secret Manager.
	//  Format: {secret} if the secret is in the same project.
	//  projects/{project}/secrets/{secret} if the secret is
	//  in a different project.
	// +kcc:proto:field=google.cloud.run.v2.SecretVolumeSource.secret
	Secret *string `json:"secret,omitempty"`

	// If unspecified, the volume will expose a file whose name is the
	//  secret, relative to VolumeMount.mount_path.
	//  If specified, the key will be used as the version to fetch from Cloud
	//  Secret Manager and the path will be the name of the file exposed in the
	//  volume. When items are defined, they must specify a path and a version.
	// +kcc:proto:field=google.cloud.run.v2.SecretVolumeSource.items
	Items []VersionToPath `json:"items,omitempty"`

	// Integer representation of mode bits to use on created files by default.
	//  Must be a value between 0000 and 0777 (octal), defaulting to 0444.
	//  Directories within the path are not affected by  this setting.
	//
	//  Notes
	//
	//  * Internally, a umask of 0222 will be applied to any non-zero value.
	//  * This is an integer representation of the mode bits. So, the octal
	//  integer value should look exactly as the chmod numeric notation with a
	//  leading zero. Some examples: for chmod 640 (u=rw,g=r), set to 0640 (octal)
	//  or 416 (base-10). For chmod 755 (u=rwx,g=rx,o=rx), set to 0755 (octal) or
	//  493 (base-10).
	//  * This might be in conflict with other options that affect the
	//  file mode, like fsGroup, and the result can be other mode bits set.
	//
	//  This might be in conflict with other options that affect the
	//  file mode, like fsGroup, and as a result, other mode bits could be set.
	// +kcc:proto:field=google.cloud.run.v2.SecretVolumeSource.default_mode
	DefaultMode *int32 `json:"defaultMode,omitempty"`
}

// +kcc:proto=google.cloud.run.v2.TCPSocketAction
type TCPSocketAction struct {
	// Optional. Port number to access on the container. Must be in the range 1 to
	//  65535. If not specified, defaults to the exposed port of the container,
	//  which is the value of container.ports[0].containerPort.
	// +kcc:proto:field=google.cloud.run.v2.TCPSocketAction.port
	Port *int32 `json:"port,omitempty"`
}

// +kcc:proto=google.cloud.run.v2.TaskTemplate
type TaskTemplate struct {
	// Holds the single container that defines the unit of execution for this
	//  task.
	// +kcc:proto:field=google.cloud.run.v2.TaskTemplate.containers
	Containers []Container `json:"containers,omitempty"`

	// Optional. A list of Volumes to make available to containers.
	// +kcc:proto:field=google.cloud.run.v2.TaskTemplate.volumes
	Volumes []Volume `json:"volumes,omitempty"`

	// Number of retries allowed per Task, before marking this Task failed.
	//  Defaults to 3.
	// +kcc:proto:field=google.cloud.run.v2.TaskTemplate.max_retries
	MaxRetries *int32 `json:"maxRetries,omitempty"`

	// Optional. Max allowed time duration the Task may be active before the
	//  system will actively try to mark it failed and kill associated containers.
	//  This applies per attempt of a task, meaning each retry can run for the full
	//  timeout. Defaults to 600 seconds.
	// +kcc:proto:field=google.cloud.run.v2.TaskTemplate.timeout
	Timeout *string `json:"timeout,omitempty"`

	// Optional. Email address of the IAM service account associated with the Task
	//  of a Job. The service account represents the identity of the running task,
	//  and determines what permissions the task has. If not provided, the task
	//  will use the project's default service account.
	// +kcc:proto:field=google.cloud.run.v2.TaskTemplate.service_account
	ServiceAccount *string `json:"serviceAccount,omitempty"`

	// Optional. The execution environment being used to host this Task.
	// +kcc:proto:field=google.cloud.run.v2.TaskTemplate.execution_environment
	ExecutionEnvironment *string `json:"executionEnvironment,omitempty"`

	// A reference to a customer managed encryption key (CMEK) to use to encrypt
	//  this container image. For more information, go to
	//  https://cloud.google.com/run/docs/securing/using-cmek
	// +kcc:proto:field=google.cloud.run.v2.TaskTemplate.encryption_key
	EncryptionKey *string `json:"encryptionKey,omitempty"`

	// Optional. VPC Access configuration to use for this Task. For more
	//  information, visit
	//  https://cloud.google.com/run/docs/configuring/connecting-vpc.
	// +kcc:proto:field=google.cloud.run.v2.TaskTemplate.vpc_access
	VpcAccess *VpcAccess `json:"vpcAccess,omitempty"`
}

// +kcc:proto=google.cloud.run.v2.VersionToPath
type VersionToPath struct {
	// Required. The relative path of the secret in the container.
	// +kcc:proto:field=google.cloud.run.v2.VersionToPath.path
	Path *string `json:"path,omitempty"`

	// The Cloud Secret Manager secret version.
	//  Can be 'latest' for the latest value, or an integer or a secret alias for a
	//  specific version.
	// +kcc:proto:field=google.cloud.run.v2.VersionToPath.version
	Version *string `json:"version,omitempty"`

	// Integer octal mode bits to use on this file, must be a value between
	//  01 and 0777 (octal). If 0 or not set, the Volume's default mode will be
	//  used.
	//
	//  Notes
	//
	//  * Internally, a umask of 0222 will be applied to any non-zero value.
	//  * This is an integer representation of the mode bits. So, the octal
	//  integer value should look exactly as the chmod numeric notation with a
	//  leading zero. Some examples: for chmod 640 (u=rw,g=r), set to 0640 (octal)
	//  or 416 (base-10). For chmod 755 (u=rwx,g=rx,o=rx), set to 0755 (octal) or
	//  493 (base-10).
	//  * This might be in conflict with other options that affect the
	//  file mode, like fsGroup, and the result can be other mode bits set.
	// +kcc:proto:field=google.cloud.run.v2.VersionToPath.mode
	Mode *int32 `json:"mode,omitempty"`
}

// +kcc:proto=google.cloud.run.v2.Volume
type Volume struct {
	// Required. Volume's name.
	// +kcc:proto:field=google.cloud.run.v2.Volume.name
	Name *string `json:"name,omitempty"`

	// Secret represents a secret that should populate this volume.
	// +kcc:proto:field=google.cloud.run.v2.Volume.secret
	Secret *SecretVolumeSource `json:"secret,omitempty"`

	// For Cloud SQL volumes, contains the specific instances that should be
	//  mounted. Visit https://cloud.google.com/sql/docs/mysql/connect-run for
	//  more information on how to connect Cloud SQL and Cloud Run.
	// +kcc:proto:field=google.cloud.run.v2.Volume.cloud_sql_instance
	CloudSqlInstance *CloudSqlInstance `json:"cloudSqlInstance,omitempty"`

	// Ephemeral storage used as a shared volume.
	// +kcc:proto:field=google.cloud.run.v2.Volume.empty_dir
	EmptyDir *EmptyDirVolumeSource `json:"emptyDir,omitempty"`

	// For NFS Voumes, contains the path to the nfs Volume
	// +kcc:proto:field=google.cloud.run.v2.Volume.nfs
	Nfs *NFSVolumeSource `json:"nfs,omitempty"`

	// Persistent storage backed by a Google Cloud Storage bucket.
	// +kcc:proto:field=google.cloud.run.v2.Volume.gcs
	Gcs *GCSVolumeSource `json:"gcs,omitempty"`
}

// +kcc:proto=google.cloud.run.v2.VolumeMount
type VolumeMount struct {
	// Required. This must match the Name of a Volume.
	// +kcc:proto:field=google.cloud.run.v2.VolumeMount.name
	Name *string `json:"name,omitempty"`

	// Required. Path within the container at which the volume should be mounted.
	//  Must not contain ':'. For Cloud SQL volumes, it can be left empty, or must
	//  otherwise be `/cloudsql`. All instances defined in the Volume will be
	//  available as `/cloudsql/[instance]`. For more information on Cloud SQL
	//  volumes, visit https://cloud.google.com/sql/docs/mysql/connect-run
	// +kcc:proto:field=google.cloud.run.v2.VolumeMount.mount_path
	MountPath *string `json:"mountPath,omitempty"`
}

// +kcc:proto=google.cloud.run.v2.VpcAccess
type VpcAccess struct {
	// VPC Access connector name.
	//  Format: `projects/{project}/locations/{location}/connectors/{connector}`,
	//  where `{project}` can be project id or number.
	//  For more information on sending traffic to a VPC network via a connector,
	//  visit https://cloud.google.com/run/docs/configuring/vpc-connectors.
	// +kcc:proto:field=google.cloud.run.v2.VpcAccess.connector
	Connector *string `json:"connector,omitempty"`

	// Optional. Traffic VPC egress settings. If not provided, it defaults to
	//  PRIVATE_RANGES_ONLY.
	// +kcc:proto:field=google.cloud.run.v2.VpcAccess.egress
	Egress *string `json:"egress,omitempty"`

	// Optional. Direct VPC egress settings. Currently only single network
	//  interface is supported.
	// +kcc:proto:field=google.cloud.run.v2.VpcAccess.network_interfaces
	NetworkInterfaces []VpcAccess_NetworkInterface `json:"networkInterfaces,omitempty"`
}

// +kcc:proto=google.cloud.run.v2.VpcAccess.NetworkInterface
type VpcAccess_NetworkInterface struct {
	// Optional. The VPC network that the Cloud Run resource will be able to
	//  send traffic to. At least one of network or subnetwork must be specified.
	//  If both network and subnetwork are specified, the given VPC subnetwork
	//  must belong to the given VPC network. If network is not specified, it
	//  will be looked up from the subnetwork.
	// +kcc:proto:field=google.cloud.run.v2.VpcAccess.NetworkInterface.network
	Network *string `json:"network,omitempty"`

	// Optional. The VPC subnetwork that the Cloud Run resource will get IPs
	//  from. At least one of network or subnetwork must be specified. If both
	//  network and subnetwork are specified, the given VPC subnetwork must
	//  belong to the given VPC network. If subnetwork is not specified, the
	//  subnetwork with the same name with the network will be used.
	// +kcc:proto:field=google.cloud.run.v2.VpcAccess.NetworkInterface.subnetwork
	Subnetwork *string `json:"subnetwork,omitempty"`

	// Optional. Network tags applied to this Cloud Run resource.
	// +kcc:proto:field=google.cloud.run.v2.VpcAccess.NetworkInterface.tags
	Tags []string `json:"tags,omitempty"`
}

// +kcc:proto=google.cloud.run.v2.Condition
type ConditionObservedState struct {
	// Output only. A common (service-level) reason for this condition.
	// +kcc:proto:field=google.cloud.run.v2.Condition.reason
	Reason *string `json:"reason,omitempty"`

	// Output only. A reason for the revision condition.
	// +kcc:proto:field=google.cloud.run.v2.Condition.revision_reason
	RevisionReason *string `json:"revisionReason,omitempty"`

	// Output only. A reason for the execution condition.
	// +kcc:proto:field=google.cloud.run.v2.Condition.execution_reason
	ExecutionReason *string `json:"executionReason,omitempty"`
}

// +kcc:proto=google.cloud.run.v2.Execution
type ExecutionObservedState struct {
	// Output only. The unique name of this Execution.
	// +kcc:proto:field=google.cloud.run.v2.Execution.name
	Name *string `json:"name,omitempty"`

	// Output only. Server assigned unique identifier for the Execution. The value
	//  is a UUID4 string and guaranteed to remain unchanged until the resource is
	//  deleted.
	// +kcc:proto:field=google.cloud.run.v2.Execution.uid
	Uid *string `json:"uid,omitempty"`

	// Output only. A number that monotonically increases every time the user
	//  modifies the desired state.
	// +kcc:proto:field=google.cloud.run.v2.Execution.generation
	Generation *int64 `json:"generation,omitempty"`

	// Output only. Unstructured key value map that can be used to organize and
	//  categorize objects. User-provided labels are shared with Google's billing
	//  system, so they can be used to filter, or break down billing charges by
	//  team, component, environment, state, etc. For more information, visit
	//  https://cloud.google.com/resource-manager/docs/creating-managing-labels or
	//  https://cloud.google.com/run/docs/configuring/labels
	// +kcc:proto:field=google.cloud.run.v2.Execution.labels
	Labels map[string]string `json:"labels,omitempty"`

	// Output only. Unstructured key value map that may
	//  be set by external tools to store and arbitrary metadata.
	//  They are not queryable and should be preserved
	//  when modifying objects.
	// +kcc:proto:field=google.cloud.run.v2.Execution.annotations
	Annotations map[string]string `json:"annotations,omitempty"`

	// Output only. Represents time when the execution was acknowledged by the
	//  execution controller. It is not guaranteed to be set in happens-before
	//  order across separate operations.
	// +kcc:proto:field=google.cloud.run.v2.Execution.create_time
	CreateTime *string `json:"createTime,omitempty"`

	// Output only. Represents time when the execution started to run.
	//  It is not guaranteed to be set in happens-before order across separate
	//  operations.
	// +kcc:proto:field=google.cloud.run.v2.Execution.start_time
	StartTime *string `json:"startTime,omitempty"`

	// Output only. Represents time when the execution was completed. It is not
	//  guaranteed to be set in happens-before order across separate operations.
	// +kcc:proto:field=google.cloud.run.v2.Execution.completion_time
	CompletionTime *string `json:"completionTime,omitempty"`

	// Output only. The last-modified time.
	// +kcc:proto:field=google.cloud.run.v2.Execution.update_time
	UpdateTime *string `json:"updateTime,omitempty"`

	// Output only. For a deleted resource, the deletion time. It is only
	//  populated as a response to a Delete request.
	// +kcc:proto:field=google.cloud.run.v2.Execution.delete_time
	DeleteTime *string `json:"deleteTime,omitempty"`

	// Output only. For a deleted resource, the time after which it will be
	//  permamently deleted. It is only populated as a response to a Delete
	//  request.
	// +kcc:proto:field=google.cloud.run.v2.Execution.expire_time
	ExpireTime *string `json:"expireTime,omitempty"`

	// Output only. The name of the parent Job.
	// +kcc:proto:field=google.cloud.run.v2.Execution.job
	Job *string `json:"job,omitempty"`

	// Output only. Specifies the maximum desired number of tasks the execution
	//  should run at any given time. Must be <= task_count. The actual number of
	//  tasks running in steady state will be less than this number when
	//  ((.spec.task_count - .status.successful) < .spec.parallelism), i.e. when
	//  the work left to do is less than max parallelism.
	// +kcc:proto:field=google.cloud.run.v2.Execution.parallelism
	Parallelism *int32 `json:"parallelism,omitempty"`

	// Output only. Specifies the desired number of tasks the execution should
	//  run. Setting to 1 means that parallelism is limited to 1 and the success of
	//  that task signals the success of the execution.
	// +kcc:proto:field=google.cloud.run.v2.Execution.task_count
	TaskCount *int32 `json:"taskCount,omitempty"`

	// Output only. The template used to create tasks for this execution.
	// +kcc:proto:field=google.cloud.run.v2.Execution.template
	Template *TaskTemplate `json:"template,omitempty"`

	// Output only. Indicates whether the resource's reconciliation is still in
	//  progress. See comments in `Job.reconciling` for additional information on
	//  reconciliation process in Cloud Run.
	// +kcc:proto:field=google.cloud.run.v2.Execution.reconciling
	Reconciling *bool `json:"reconciling,omitempty"`

	// Output only. The Condition of this Execution, containing its readiness
	//  status, and detailed error information in case it did not reach the desired
	//  state.
	// +kcc:proto:field=google.cloud.run.v2.Execution.conditions
	Conditions []Condition `json:"conditions,omitempty"`

	// Output only. The generation of this Execution. See comments in
	//  `reconciling` for additional information on reconciliation process in Cloud
	//  Run.
	// +kcc:proto:field=google.cloud.run.v2.Execution.observed_generation
	ObservedGeneration *int64 `json:"observedGeneration,omitempty"`

	// Output only. The number of actively running tasks.
	// +kcc:proto:field=google.cloud.run.v2.Execution.running_count
	RunningCount *int32 `json:"runningCount,omitempty"`

	// Output only. The number of tasks which reached phase Succeeded.
	// +kcc:proto:field=google.cloud.run.v2.Execution.succeeded_count
	SucceededCount *int32 `json:"succeededCount,omitempty"`

	// Output only. The number of tasks which reached phase Failed.
	// +kcc:proto:field=google.cloud.run.v2.Execution.failed_count
	FailedCount *int32 `json:"failedCount,omitempty"`

	// Output only. The number of tasks which reached phase Cancelled.
	// +kcc:proto:field=google.cloud.run.v2.Execution.cancelled_count
	CancelledCount *int32 `json:"cancelledCount,omitempty"`

	// Output only. The number of tasks which have retried at least once.
	// +kcc:proto:field=google.cloud.run.v2.Execution.retried_count
	RetriedCount *int32 `json:"retriedCount,omitempty"`

	// Output only. URI where logs for this execution can be found in Cloud
	//  Console.
	// +kcc:proto:field=google.cloud.run.v2.Execution.log_uri
	LogURI *string `json:"logURI,omitempty"`

	// Output only. Reserved for future use.
	// +kcc:proto:field=google.cloud.run.v2.Execution.satisfies_pzs
	SatisfiesPzs *bool `json:"satisfiesPzs,omitempty"`

	// Output only. A system-generated fingerprint for this version of the
	//  resource. May be used to detect modification conflict during updates.
	// +kcc:proto:field=google.cloud.run.v2.Execution.etag
	Etag *string `json:"etag,omitempty"`
}
