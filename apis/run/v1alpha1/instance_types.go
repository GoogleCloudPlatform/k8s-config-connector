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
	computev1beta1 "github.com/GoogleCloudPlatform/k8s-config-connector/apis/compute/v1beta1"
	refs "github.com/GoogleCloudPlatform/k8s-config-connector/apis/refs/v1beta1"
	refsv1beta1 "github.com/GoogleCloudPlatform/k8s-config-connector/apis/refs/v1beta1"
	secretmanagerv1beta1 "github.com/GoogleCloudPlatform/k8s-config-connector/apis/secretmanager/v1beta1"
	storagev1beta1 "github.com/GoogleCloudPlatform/k8s-config-connector/apis/storage/v1beta1"
	vpcaccessv1beta1 "github.com/GoogleCloudPlatform/k8s-config-connector/apis/vpcaccess/v1beta1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/apis/k8s/v1alpha1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

var CloudRunInstanceGVK = GroupVersion.WithKind("CloudRunInstance")

// CloudRunInstanceSpec defines the desired state of CloudRunInstance
// +kcc:spec:proto=google.cloud.run.v2.Instance
type CloudRunInstanceSpec struct {
	ProjectRef                    *refsv1beta1.ProjectRef           `json:"projectRef"`
	Location                      *string                           `json:"location,omitempty"`
	ResourceID                    *string                           `json:"resourceID,omitempty"`
	Description                   *string                           `json:"description,omitempty"`
	Labels                        map[string]string                 `json:"labels,omitempty"`
	Annotations                   map[string]string                 `json:"annotations,omitempty"`
	Client                        *string                           `json:"client,omitempty"`
	ClientVersion                 *string                           `json:"clientVersion,omitempty"`
	LaunchStage                   *string                           `json:"launchStage,omitempty"`
	BinaryAuthorization           *InstanceBinaryAuthorization      `json:"binaryAuthorization,omitempty"`
	VpcAccess                     *InstanceVPCAccess                `json:"vpcAccess,omitempty"`
	ServiceAccountRef             *refsv1beta1.IAMServiceAccountRef `json:"serviceAccountRef,omitempty"`
	Containers                    []InstanceContainer               `json:"containers"`
	Volumes                       []InstanceVolume                  `json:"volumes,omitempty"`
	EncryptionKeyRef              *refsv1beta1.KMSCryptoKeyRef      `json:"encryptionKeyRef,omitempty"`
	EncryptionKeyRevocationAction *string                           `json:"encryptionKeyRevocationAction,omitempty"`
	EncryptionKeyShutdownDuration *string                           `json:"encryptionKeyShutdownDuration,omitempty"`
	NodeSelector                  *InstanceNodeSelector             `json:"nodeSelector,omitempty"`
	GpuZonalRedundancyDisabled    *bool                             `json:"gpuZonalRedundancyDisabled,omitempty"`
	Ingress                       *string                           `json:"ingress,omitempty"`
	InvokerIAMDisabled            *bool                             `json:"invokerIAMDisabled,omitempty"`
	IapEnabled                    *bool                             `json:"iapEnabled,omitempty"`
}

// +kcc:observedstate:proto=google.cloud.run.v2.Instance
type CloudRunInstanceObservedState struct {
	Uid                *string                   `json:"uid,omitempty"`
	Generation         *int64                    `json:"generation,omitempty"`
	CreateTime         *string                   `json:"createTime,omitempty"`
	UpdateTime         *string                   `json:"updateTime,omitempty"`
	DeleteTime         *string                   `json:"deleteTime,omitempty"`
	ExpireTime         *string                   `json:"expireTime,omitempty"`
	Creator            *string                   `json:"creator,omitempty"`
	LastModifier       *string                   `json:"lastModifier,omitempty"`
	ObservedGeneration *int64                    `json:"observedGeneration,omitempty"`
	LogURI             *string                   `json:"logURI,omitempty"`
	TerminalCondition  *InstanceCondition        `json:"terminalCondition,omitempty"`
	ContainerStatuses  []InstanceContainerStatus `json:"containerStatuses,omitempty"`
	SatisfiesPzs       *bool                     `json:"satisfiesPzs,omitempty"`
	Urls               []string                  `json:"urls,omitempty"`
	Reconciling        *bool                     `json:"reconciling,omitempty"`
	Etag               *string                   `json:"etag,omitempty"`
}

// CloudRunInstanceStatus defines the config connector machine state of CloudRunInstance
type CloudRunInstanceStatus struct {
	Conditions         []v1alpha1.Condition           `json:"conditions,omitempty"`
	ObservedGeneration *int64                         `json:"observedGeneration,omitempty"`
	ExternalRef        *string                        `json:"externalRef,omitempty"`
	ObservedState      *CloudRunInstanceObservedState `json:"observedState,omitempty"`
}

// +genclient
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:resource:categories=gcp,shortName=gcpcloudruninstance;gcpcloudruninstances
// +kubebuilder:subresource:status
// +kubebuilder:metadata:labels="cnrm.cloud.google.com/managed-by-kcc=true"
// +kubebuilder:metadata:labels="cnrm.cloud.google.com/system=true"
// +kubebuilder:printcolumn:name="Age",JSONPath=".metadata.creationTimestamp",type="date"
// +kubebuilder:printcolumn:name="Ready",JSONPath=".status.conditions[?(@.type=='Ready')].status",type="string",description="When 'True', the most recent reconcile of the resource succeeded"
// +kubebuilder:printcolumn:name="Status",JSONPath=".status.conditions[?(@.type=='Ready')].reason",type="string",description="The reason for the value in 'Ready'"
// +kubebuilder:printcolumn:name="Status Age",JSONPath=".status.conditions[?(@.type=='Ready')].lastTransitionTime",type="date",description="The last transition time for the value in 'Status'"

// CloudRunInstance is the Schema for the CloudRunInstance API
// +k8s:openapi-gen=true
type CloudRunInstance struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	// +required
	Spec   CloudRunInstanceSpec   `json:"spec,omitempty"`
	Status CloudRunInstanceStatus `json:"status,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// CloudRunInstanceList contains a list of CloudRunInstance
type CloudRunInstanceList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []CloudRunInstance `json:"items"`
}

func init() {
	SchemeBuilder.Register(&CloudRunInstance{}, &CloudRunInstanceList{})
}

// +kcc:proto=google.cloud.run.v2.BinaryAuthorization
type InstanceBinaryAuthorization struct {
	// Optional. If True, indicates to use the default project's binary
	//  authorization policy. If False, binary authorization will be disabled.
	// +kcc:proto:field=google.cloud.run.v2.BinaryAuthorization.use_default
	UseDefault *bool `json:"useDefault,omitempty"`

	// Optional. The path to a binary authorization policy.
	//  Format: `projects/{project}/platforms/cloudRun/{policy-name}`
	// +kcc:proto:field=google.cloud.run.v2.BinaryAuthorization.policy
	// Policy *string `json:"policy,omitempty"`

	// Optional. If present, indicates to use Breakglass using this justification.
	//  If use_default is False, then it must be empty.
	//  For more information on breakglass, see
	//  https://cloud.google.com/binary-authorization/docs/using-breakglass
	// +kcc:proto:field=google.cloud.run.v2.BinaryAuthorization.breakglass_justification
	BreakglassJustification *string `json:"breakglassJustification,omitempty"`
}

// +kcc:proto=google.cloud.run.v2.VpcAccess
type InstanceVPCAccess struct {
	// VPC Access connector name.
	//  Format: `projects/{project}/locations/{location}/connectors/{connector}`,
	//  where `{project}` can be project id or number.
	//  For more information on sending traffic to a VPC network via a connector,
	//  visit https://cloud.google.com/run/docs/configuring/vpc-connectors.
	// +kcc:proto:field=google.cloud.run.v2.VpcAccess.connector
	// +kcc:ref=VPCAccessConnector
	ConnectorRef *vpcaccessv1beta1.VPCAccessConnectorRef `json:"connectorRef,omitempty"`

	// Optional. Traffic VPC egress settings. If not provided, it defaults to
	//  PRIVATE_RANGES_ONLY.
	// +kcc:proto:field=google.cloud.run.v2.VpcAccess.egress
	Egress *string `json:"egress,omitempty"`

	// Optional. Direct VPC egress settings. Currently only single network
	//  interface is supported.
	// +kcc:proto:field=google.cloud.run.v2.VpcAccess.network_interfaces
	NetworkInterfaces []InstanceVPCAccess_NetworkInterface `json:"networkInterfaces,omitempty"`
}

// +kcc:proto=google.cloud.run.v2.Container
type InstanceContainer struct {
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
	Env []InstanceEnvVar `json:"env,omitempty"`

	// Compute Resource requirements by this container.
	// +kcc:proto:field=google.cloud.run.v2.Container.resources
	Resources *InstanceResourceRequirements `json:"resources,omitempty"`

	// List of ports to expose from the container. Only a single port can be
	//  specified. The specified ports must be listening on all interfaces
	//  (0.0.0.0) within the container to be accessible.
	//
	//  If omitted, a port number will be chosen and passed to the container
	//  through the PORT environment variable for the container to listen on.
	// +kcc:proto:field=google.cloud.run.v2.Container.ports
	Ports []InstanceContainerPort `json:"ports,omitempty"`

	// InstanceVolume to mount into the container's filesystem.
	// +kcc:proto:field=google.cloud.run.v2.Container.volume_mounts
	VolumeMounts []InstanceVolumeMount `json:"volumeMounts,omitempty"`

	// Container's working directory.
	//  If not specified, the container runtime's default will be used, which
	//  might be configured in the container image.
	// +kcc:proto:field=google.cloud.run.v2.Container.working_dir
	WorkingDir *string `json:"workingDir,omitempty"`

	// Periodic probe of container liveness.
	//  Container will be restarted if the probe fails.
	// +kcc:proto:field=google.cloud.run.v2.Container.liveness_probe
	LivenessProbe *InstanceProbe `json:"livenessProbe,omitempty"`

	// Startup probe of application within the container.
	//  All other probes are disabled if a startup probe is provided, until it
	//  succeeds. Container will not be added to service endpoints if the probe
	//  fails.
	// +kcc:proto:field=google.cloud.run.v2.Container.startup_probe
	StartupProbe *InstanceProbe `json:"startupProbe,omitempty"`

	// Names of the containers that must start before this container.
	// +kcc:proto:field=google.cloud.run.v2.Container.depends_on
	DependsOn []string `json:"dependsOn,omitempty"`

	// Names of the containers that must start before this container.
	// +kcc:proto:field=google.cloud.run.v2.Container.depends_on
	// DependsOn []string `json:"dependsOn,omitempty"`

	// Base image for this container. Only supported for services. If set, it
	//  indicates that the service is enrolled into automatic base image update.
	// +kcc:proto:field=google.cloud.run.v2.Container.base_image_uri
	// BaseImageURI *string `json:"baseImageURI,omitempty"`
}

// +kcc:proto=google.cloud.run.v2.Volume
type InstanceVolume struct {
	// Required. Volume's name.
	// +kcc:proto:field=google.cloud.run.v2.Volume.name
	Name *string `json:"name,omitempty"`

	// Secret represents a secret that should populate this volume.
	// +kcc:proto:field=google.cloud.run.v2.Volume.secret
	Secret *InstanceSecretVolumeSource `json:"secret,omitempty"`

	// For Cloud SQL volumes, contains the specific instances that should be
	//  mounted. Visit https://cloud.google.com/sql/docs/mysql/connect-run for
	//  more information on how to connect Cloud SQL and Cloud Run.
	// +kcc:proto:field=google.cloud.run.v2.Volume.cloud_sql_instance
	CloudSQLInstance *InstanceCloudSQLInstance `json:"cloudSQLInstance,omitempty"`

	// Ephemeral storage used as a shared volume.
	// +kcc:proto:field=google.cloud.run.v2.Volume.empty_dir
	EmptyDir *InstanceEmptyDirVolumeSource `json:"emptyDir,omitempty"`

	// For Nfs Volumes, contains the path to the nfs Volume
	// +kcc:proto:field=google.cloud.run.v2.Volume.nfs
	Nfs *InstanceNfsVolumeSource `json:"nfs,omitempty"`

	// Persistent storage backed by a Google Cloud Storage bucket.
	// +kcc:proto:field=google.cloud.run.v2.Volume.gcs
	GCS *InstanceGCSVolumeSource `json:"gcs,omitempty"`
}

// +kcc:proto=google.cloud.run.v2.Condition
type InstanceCondition struct {
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

	// Note: `Reason`, `RevisionReason` and `ExecutionReason` are OneOf fields. The GCP proto uses a single interface field
	// but the KCC (DCL-based) has split them into individual fields.

	// A common (service-level) reason for this condition.
	Reason *string `json:"reason,omitempty"`

	// A reason for the revision condition.
	RevisionReason *string `json:"revisionReason,omitempty"`

	// A reason for the execution condition.
	ExecutionReason *string `json:"executionReason,omitempty"`
}

// +kcc:proto=google.cloud.run.v2.EnvVarSource
type InstanceEnvVarSource struct {
	// Selects a secret and a specific version from Cloud Secret Manager.
	// +kcc:proto:field=google.cloud.run.v2.InstanceEnvVarSource.secret_key_ref
	SecretKeyRef *InstanceSecretKeySelector `json:"secretKeyRef,omitempty"`
}

// +kcc:proto=google.cloud.run.v2.SecretKeySelector
type InstanceSecretKeySelector struct {
	// Required.
	// The name of the secret in Cloud Secret  Manager. Format: {secret} if the secret is in
	// the same project. projects/{project}/secrets/{secret}
	// +kcc:proto:field=google.cloud.run.v2.SecretKeySelector.secret
	SecretRef *secretmanagerv1beta1.SecretRef `json:"secretRef,omitempty"`

	// The Cloud Secret Manager secret version.
	//  Can be 'latest' for the latest version, an integer for a specific version,
	//  or a version alias.
	// +kcc:proto:field=google.cloud.run.v2.SecretKeySelector.version
	VersionRef *secretmanagerv1beta1.SecretVersionRef `json:"versionRef,omitempty"`
}

// +kcc:proto=google.cloud.run.v2.ResourceRequirements
type InstanceResourceRequirements struct {
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
	// CPUIdle *bool `json:"cpuIdle,omitempty"`

	// Determines whether CPU should be boosted on startup of a new container
	//  instance above the requested CPU threshold, this can help reduce cold-start
	//  latency.
	// +kcc:proto:field=google.cloud.run.v2.ResourceRequirements.startup_cpu_boost
	// StartupCPUBoost *bool `json:"startupCPUBoost,omitempty"`
}

// +kcc:proto=google.cloud.run.v2.Probe
type InstanceProbe struct {
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
	HTTPGet *InstanceHTTPGetAction `json:"httpGet,omitempty"`

	// Optional. TCPSocket specifies an action involving a TCP port.
	//  Exactly one of httpGet, tcpSocket, or grpc must be specified.
	// +kcc:proto:field=google.cloud.run.v2.Probe.tcp_socket
	TCPSocket *InstanceTCPSocketAction `json:"tcpSocket,omitempty"`

	// Optional. GRPC specifies an action involving a gRPC port.
	//  Exactly one of httpGet, tcpSocket, or grpc must be specified.
	// +kcc:proto:field=google.cloud.run.v2.Probe.grpc
	// Grpc *InstanceGrpcAction `json:"grpc,omitempty"`
}

// +kcc:proto=google.cloud.run.v2.HTTPGetAction
type InstanceHTTPGetAction struct {
	// Optional. Path to access on the HTTP server. Defaults to '/'.
	// +kcc:proto:field=google.cloud.run.v2.HTTPGetAction.path
	Path *string `json:"path,omitempty"`

	// Optional. Custom headers to set in the request. HTTP allows repeated
	//  headers.
	// +kcc:proto:field=google.cloud.run.v2.HTTPGetAction.http_headers
	HttpHeaders []InstanceHTTPHeader `json:"httpHeaders,omitempty"`

	// Optional. Port number to access on the container. Must be in the range 1 to
	//  65535. If not specified, defaults to the exposed port of the container,
	//  which is the value of container.ports[0].containerPort.
	// +kcc:proto:field=google.cloud.run.v2.HTTPGetAction.port
	Port *int32 `json:"port,omitempty"`
}

// +kcc:proto=google.cloud.run.v2.SecretVolumeSource
type InstanceSecretVolumeSource struct {
	// Required. The name of the secret in Cloud Secret Manager.
	//  Format: {secret} if the secret is in the same project.
	//  projects/{project}/secrets/{secret} if the secret is
	//  in a different project.
	// +kcc:proto:field=google.cloud.run.v2.SecretVolumeSource.secret
	SecretRef *secretmanagerv1beta1.SecretRef `json:"secretRef,omitempty"`

	// If unspecified, the volume will expose a file whose name is the
	//  secret, relative to InstanceVolumeMount.mount_path.
	//  If specified, the key will be used as the version to fetch from Cloud
	//  Secret Manager and the path will be the name of the file exposed in the
	//  volume. When items are defined, they must specify a path and a version.
	// +kcc:proto:field=google.cloud.run.v2.SecretVolumeSource.items
	Items []InstanceVersionToPath `json:"items,omitempty"`

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

// +kcc:proto=google.cloud.run.v2.CloudSqlInstance
type InstanceCloudSQLInstance struct {
	// The Cloud SQL instance connection names, as can be found in
	//  https://console.cloud.google.com/sql/instances. Visit
	//  https://cloud.google.com/sql/docs/mysql/connect-run for more information on
	//  how to connect Cloud SQL and Cloud Run. Format:
	//  {project}:{location}:{instance}
	// +kcc:proto:field=google.cloud.run.v2.CloudSqlInstance.instances
	InstanceRefs []*refs.SQLInstanceRef `json:"instanceRefs,omitempty"`
}

// +kcc:proto=google.cloud.run.v2.GCSVolumeSource
type InstanceGCSVolumeSource struct {
	// Cloud Storage Bucket name.
	// +kcc:proto:field=google.cloud.run.v2.GCSVolumeSource.bucket
	BucketRef *storagev1beta1.StorageBucketRef `json:"bucketRef,omitempty"`

	// If true, the volume will be mounted as read only for all mounts.
	// +kcc:proto:field=google.cloud.run.v2.GCSVolumeSource.read_only
	ReadOnly *bool `json:"readOnly,omitempty"`

	// A list of additional flags to pass to the gcsfuse CLI.
	//  Options should be specified without the leading "--".
	// +kcc:proto:field=google.cloud.run.v2.GCSVolumeSource.mount_options
	MountOptions []string `json:"mountOptions,omitempty"`
}

// +kcc:proto=google.cloud.run.v2.VersionToPath
type InstanceVersionToPath struct {
	// Required. The relative path of the secret in the container.
	// +kcc:proto:field=google.cloud.run.v2.VersionToPath.path
	Path *string `json:"path,omitempty"`

	// The Cloud Secret Manager secret version.
	//  Can be 'latest' for the latest value, or an integer or a secret alias for a
	//  specific version.
	// +kcc:proto:field=google.cloud.run.v2.VersionToPath.version
	VersionRef *secretmanagerv1beta1.SecretVersionRef `json:"versionRef,omitempty"`

	// Integer octal mode bits to use on this file, must be a value between
	//  01 and 0777 (octal). If 0 or not set, the InstanceVolume's default mode will be
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

// +kcc:proto=google.cloud.run.v2.NodeSelector
type InstanceNodeSelector struct {
	Accelerator *string `json:"accelerator,omitempty"`
}

// +kcc:proto=google.cloud.run.v2.ContainerStatus
type InstanceContainerStatus struct {
	Name        *string `json:"name,omitempty"`
	ImageDigest *string `json:"imageDigest,omitempty"`
}

// +kcc:proto=google.cloud.run.v2.EnvVar
type InstanceEnvVar struct {
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
	ValueSource *InstanceEnvVarSource `json:"valueSource,omitempty"`
}

// +kcc:proto=google.cloud.run.v2.ContainerPort
type InstanceContainerPort struct {
	// If specified, used to specify which protocol to use.
	//  Allowed values are "http1" and "h2c".
	// +kcc:proto:field=google.cloud.run.v2.ContainerPort.name
	Name *string `json:"name,omitempty"`

	// Port number the container listens on.
	//  This must be a valid TCP port number, 0 < container_port < 65536.
	// +kcc:proto:field=google.cloud.run.v2.ContainerPort.container_port
	ContainerPort *int32 `json:"containerPort,omitempty"`
}

// +kcc:proto=google.cloud.run.v2.VolumeMount
type InstanceVolumeMount struct {
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

// +kcc:proto=google.cloud.run.v2.TCPSocketAction
type InstanceTCPSocketAction struct {
	// Optional. Port number to access on the container. Must be in the range 1 to
	//  65535. If not specified, defaults to the exposed port of the container,
	//  which is the value of container.ports[0].containerPort.
	// +kcc:proto:field=google.cloud.run.v2.TCPSocketAction.port
	Port *int32 `json:"port,omitempty"`
}

// +kcc:proto=google.cloud.run.v2.HTTPHeader
type InstanceHTTPHeader struct {
	// Required. The header field name
	// +kcc:proto:field=google.cloud.run.v2.HTTPHeader.name
	Name *string `json:"name,omitempty"`

	// Optional. The header field value
	// +kcc:proto:field=google.cloud.run.v2.HTTPHeader.value
	Value *string `json:"value,omitempty"`
}

// +kcc:proto=google.cloud.run.v2.VpcAccess.NetworkInterface
type InstanceVPCAccess_NetworkInterface struct {
	// Optional. The VPC network that the Cloud Run resource will be able to
	//  send traffic to. At least one of network or subnetwork must be specified.
	//  If both network and subnetwork are specified, the given VPC subnetwork
	//  must belong to the given VPC network. If network is not specified, it
	//  will be looked up from the subnetwork.
	// +kcc:proto:field=google.cloud.run.v2.VpcAccess.NetworkInterface.network
	NetworkRef *computev1beta1.ComputeNetworkRef `json:"networkRef,omitempty"`

	// Optional. The VPC subnetwork that the Cloud Run resource will get IPs
	//  from. At least one of network or subnetwork must be specified. If both
	//  network and subnetwork are specified, the given VPC subnetwork must
	//  belong to the given VPC network. If subnetwork is not specified, the
	//  subnetwork with the same name with the network will be used.
	// +kcc:proto:field=google.cloud.run.v2.VpcAccess.NetworkInterface.subnetwork
	SubnetworkRef *computev1beta1.ComputeSubnetworkRef `json:"subnetworkRef,omitempty"`

	// Optional. Network tags applied to this Cloud Run resource.
	// +kcc:proto:field=google.cloud.run.v2.VpcAccess.NetworkInterface.tags
	Tags []string `json:"tags,omitempty"`
}

// +kcc:proto=google.cloud.run.v2.GRPCAction
type InstanceGrpcAction struct {
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

// +kcc:proto=google.cloud.run.v2.EmptyDirVolumeSource
type InstanceEmptyDirVolumeSource struct {
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

// +kcc:proto=google.cloud.run.v2.NFSVolumeSource
type InstanceNfsVolumeSource struct {
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
