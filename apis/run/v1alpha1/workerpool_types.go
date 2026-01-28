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
	secretmanagerv1beta1 "github.com/GoogleCloudPlatform/k8s-config-connector/apis/secretmanager/v1beta1"
	vpcaccessv1beta1 "github.com/GoogleCloudPlatform/k8s-config-connector/apis/vpcaccess/v1beta1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/apis/k8s/v1alpha1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

var RunWorkerPoolGVK = GroupVersion.WithKind("RunWorkerPool")

// RunWorkerPoolSpec defines the desired state of RunWorkerPool
// +kcc:spec:proto=google.cloud.run.v2.WorkerPool
type RunWorkerPoolSpec struct {
	// The project that this resource belongs to.
	ProjectRef *refs.ProjectRef `json:"projectRef"`

	// The location of this resource.
	Location string `json:"location"`

	// The RunWorkerPool name. If not given, the metadata.name will be used.
	ResourceID *string `json:"resourceID,omitempty"`

	// User-provided description of the WorkerPool. This field currently has a
	//  512-character limit.
	// +kcc:proto:field=google.cloud.run.v2.WorkerPool.description
	Description *string `json:"description,omitempty"`

	// Optional. Unstructured key value map that can be used to organize and
	//  categorize objects. User-provided labels are shared with Google's billing
	//  system, so they can be used to filter, or break down billing charges by
	//  team, component, environment, state, etc. For more information, visit
	//  https://cloud.google.com/resource-manager/docs/creating-managing-labels or
	//  https://cloud.google.com/run/docs/configuring/labels.
	//
	//  Cloud Run API v2 does not support labels with  `run.googleapis.com`,
	//  `cloud.googleapis.com`, `serving.knative.dev`, or `autoscaling.knative.dev`
	//  namespaces, and they will be rejected. All system labels in v1 now have a
	//  corresponding field in v2 WorkerPool.
	// +kcc:proto:field=google.cloud.run.v2.WorkerPool.labels
	Labels map[string]string `json:"labels,omitempty"`

	// Optional. Unstructured key value map that may be set by external tools to
	//  store and arbitrary metadata. They are not queryable and should be
	//  preserved when modifying objects.
	//
	//  Cloud Run API v2 does not support annotations with `run.googleapis.com`,
	//  `cloud.googleapis.com`, `serving.knative.dev`, or `autoscaling.knative.dev`
	//  namespaces, and they will be rejected in new resources. All system
	//  annotations in v1 now have a corresponding field in v2 WorkerPool.
	//
	//  <p>This field follows Kubernetes
	//  annotations' namespacing, limits, and rules.
	// +kcc:proto:field=google.cloud.run.v2.WorkerPool.annotations
	Annotations map[string]string `json:"annotations,omitempty"`

	// Arbitrary identifier for the API client.
	// +kcc:proto:field=google.cloud.run.v2.WorkerPool.client
	Client *string `json:"client,omitempty"`

	// Arbitrary version identifier for the API client.
	// +kcc:proto:field=google.cloud.run.v2.WorkerPool.client_version
	ClientVersion *string `json:"clientVersion,omitempty"`

	// Optional. The launch stage as defined by [Google Cloud Platform
	//   Launch Stages](https://cloud.google.com/terms/launch-stages).
	//   Cloud Run supports `ALPHA`, `BETA`, and `GA`. If no value is specified, GA
	//   is assumed.
	//   Set the launch stage to a preview stage on input to allow use of preview
	//   features in that stage. On read (or output), describes whether the
	//   resource uses preview features.
	//
	//   For example, if ALPHA is provided as input, but only BETA and GA-level
	//   features are used, this field will be BETA on output.
	// +kcc:proto:field=google.cloud.run.v2.WorkerPool.launch_stage
	LaunchStage *string `json:"launchStage,omitempty"`

	// Optional. Settings for the Binary Authorization feature.
	// +kcc:proto:field=google.cloud.run.v2.WorkerPool.binary_authorization
	BinaryAuthorization *WorkerPoolBinaryAuthorization `json:"binaryAuthorization,omitempty"`

	// Required. The template used to create revisions for this WorkerPool.
	// +kcc:proto:field=google.cloud.run.v2.WorkerPool.template
	Template *WorkerPoolRevisionTemplate `json:"template,omitempty"`

	// Optional. Specifies how to distribute instances over a collection of
	//  Revisions belonging to the WorkerPool. If instance split is empty or not
	//  provided, defaults to 100% instances assigned to the latest `Ready`
	//  Revision.
	// +kcc:proto:field=google.cloud.run.v2.WorkerPool.instance_splits
	InstanceSplits []WorkerPoolInstanceSplit `json:"instanceSplits,omitempty"`

	// Optional. Specifies worker-pool-level scaling settings
	// +kcc:proto:field=google.cloud.run.v2.WorkerPool.scaling
	Scaling *WorkerPoolScaling `json:"scaling,omitempty"`

	// Not supported, and ignored by Cloud Run.
	// +kcc:proto:field=google.cloud.run.v2.WorkerPool.custom_audiences
	CustomAudiences []string `json:"customAudiences,omitempty"`
}

// RunWorkerPoolStatus defines the config connector machine state of RunWorkerPool
type RunWorkerPoolStatus struct {
	/* Conditions represent the latest available observations of the
	   object's current state. */
	Conditions []v1alpha1.Condition `json:"conditions,omitempty"`

	// ObservedGeneration is the generation of the resource that was most recently observed by the Config Connector controller. If this is equal to metadata.generation, then that means that the current reported status reflects the most recent desired state of the resource.
	ObservedGeneration *int64 `json:"observedGeneration,omitempty"`

	// A unique specifier for the RunWorkerPool resource in GCP.
	ExternalRef *string `json:"externalRef,omitempty"`

	// ObservedState is the state of the resource as most recently observed in GCP.
	ObservedState *RunWorkerPoolObservedState `json:"observedState,omitempty"`
}

// RunWorkerPoolObservedState is the state of the RunWorkerPool resource as most recently observed in GCP.
// +kcc:observedstate:proto=google.cloud.run.v2.WorkerPool
type RunWorkerPoolObservedState struct {
	// Output only. Server assigned unique identifier for the trigger. The value
	//  is a UUID4 string and guaranteed to remain unchanged until the resource is
	//  deleted.
	// +kcc:proto:field=google.cloud.run.v2.WorkerPool.uid
	Uid *string `json:"uid,omitempty"`

	// Output only. A number that monotonically increases every time the user
	//  modifies the desired state.
	//  Please note that unlike v1, this is an int64 value. As with most Google
	//  APIs, its JSON representation will be a `string` instead of an `integer`.
	// +kcc:proto:field=google.cloud.run.v2.WorkerPool.generation
	Generation *int64 `json:"generation,omitempty"`

	// Output only. The creation time.
	// +kcc:proto:field=google.cloud.run.v2.WorkerPool.create_time
	CreateTime *string `json:"createTime,omitempty"`

	// Output only. The last-modified time.
	// +kcc:proto:field=google.cloud.run.v2.WorkerPool.update_time
	UpdateTime *string `json:"updateTime,omitempty"`

	// Output only. The deletion time. It is only populated as a response to a
	//  Delete request.
	// +kcc:proto:field=google.cloud.run.v2.WorkerPool.delete_time
	DeleteTime *string `json:"deleteTime,omitempty"`

	// Output only. For a deleted resource, the time after which it will be
	//  permamently deleted.
	// +kcc:proto:field=google.cloud.run.v2.WorkerPool.expire_time
	ExpireTime *string `json:"expireTime,omitempty"`

	// Output only. Email address of the authenticated creator.
	// +kcc:proto:field=google.cloud.run.v2.WorkerPool.creator
	Creator *string `json:"creator,omitempty"`

	// Output only. Email address of the last authenticated modifier.
	// +kcc:proto:field=google.cloud.run.v2.WorkerPool.last_modifier
	LastModifier *string `json:"lastModifier,omitempty"`

	// Required. The template used to create revisions for this WorkerPool.
	// +kcc:proto:field=google.cloud.run.v2.WorkerPool.template
	Template *WorkerPoolRevisionTemplateObservedState `json:"template,omitempty"`

	// Output only. The generation of this WorkerPool currently serving workloads.
	//  See comments in `reconciling` for additional information on reconciliation
	//  process in Cloud Run. Please note that unlike v1, this is an int64 value.
	//  As with most Google APIs, its JSON representation will be a `string`
	//  instead of an `integer`.
	// +kcc:proto:field=google.cloud.run.v2.WorkerPool.observed_generation
	ObservedGeneration *int64 `json:"observedGeneration,omitempty"`

	// Output only. The Condition of this WorkerPool, containing its readiness
	//  status, and detailed error information in case it did not reach a serving
	//  state. See comments in `reconciling` for additional information on
	//  reconciliation process in Cloud Run.
	// +kcc:proto:field=google.cloud.run.v2.WorkerPool.terminal_condition
	TerminalCondition *WorkerPoolCondition `json:"terminalCondition,omitempty"`

	// Output only. The Conditions of all other associated sub-resources. They
	//  contain additional diagnostics information in case the WorkerPool does not
	//  reach its Serving state. See comments in `reconciling` for additional
	//  information on reconciliation process in Cloud Run.
	// +kcc:proto:field=google.cloud.run.v2.WorkerPool.conditions
	Conditions []WorkerPoolCondition `json:"conditions,omitempty"`

	// Output only. Name of the latest revision that is serving workloads. See
	//  comments in `reconciling` for additional information on reconciliation
	//  process in Cloud Run.
	// +kcc:proto:field=google.cloud.run.v2.WorkerPool.latest_ready_revision
	LatestReadyRevision *string `json:"latestReadyRevision,omitempty"`

	// Output only. Name of the last created revision. See comments in
	//  `reconciling` for additional information on reconciliation process in Cloud
	//  Run.
	// +kcc:proto:field=google.cloud.run.v2.WorkerPool.latest_created_revision
	LatestCreatedRevision *string `json:"latestCreatedRevision,omitempty"`

	// Output only. Detailed status information for corresponding instance splits.
	//  See comments in `reconciling` for additional information on reconciliation
	//  process in Cloud Run.
	// +kcc:proto:field=google.cloud.run.v2.WorkerPool.instance_split_statuses
	InstanceSplitStatuses []WorkerPoolInstanceSplitStatus `json:"instanceSplitStatuses,omitempty"`

	// Output only. Indicates whether Cloud Run Threat Detection monitoring is
	//  enabled for the parent project of this worker pool.
	// +kcc:proto:field=google.cloud.run.v2.WorkerPool.threat_detection_enabled
	ThreatDetectionEnabled *bool `json:"threatDetectionEnabled,omitempty"`

	// Output only. Reserved for future use.
	// +kcc:proto:field=google.cloud.run.v2.WorkerPool.satisfies_pzs
	SatisfiesPzs *bool `json:"satisfiesPzs,omitempty"`

	// Output only. Returns true if the WorkerPool is currently being acted upon
	//  by the system to bring it into the desired state.
	// +kcc:proto:field=google.cloud.run.v2.WorkerPool.reconciling
	Reconciling *bool `json:"reconciling,omitempty"`
}

// +genclient
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:resource:categories=gcp,shortName=gcprunworkerpool;gcprunworkerpools
// +kubebuilder:subresource:status
// +kubebuilder:metadata:labels="cnrm.cloud.google.com/managed-by-kcc=true"
// +kubebuilder:metadata:labels="cnrm.cloud.google.com/system=true"
// +kubebuilder:printcolumn:name="Age",JSONPath=".metadata.creationTimestamp",type="date"
// +kubebuilder:printcolumn:name="Ready",JSONPath=".status.conditions[?(@.type=='Ready')].status",type="string",description="When 'True', the most recent reconcile of the resource succeeded"
// +kubebuilder:printcolumn:name="Status",JSONPath=".status.conditions[?(@.type=='Ready')].reason",type="string",description="The reason for the value in 'Ready'"
// +kubebuilder:printcolumn:name="Status Age",JSONPath=".status.conditions[?(@.type=='Ready')].lastTransitionTime",type="date",description="The last transition time for the value in 'Status'"

// RunWorkerPool is the Schema for the RunWorkerPool API
// +k8s:openapi-gen=true
type RunWorkerPool struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	// +required
	Spec   RunWorkerPoolSpec   `json:"spec,omitempty"`
	Status RunWorkerPoolStatus `json:"status,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// RunWorkerPoolList contains a list of RunWorkerPool
type RunWorkerPoolList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []RunWorkerPool `json:"items"`
}

func init() {
	SchemeBuilder.Register(&RunWorkerPool{}, &RunWorkerPoolList{})
}

// +kcc:proto=google.cloud.run.v2.WorkerPoolRevisionTemplate
type WorkerPoolRevisionTemplate struct {
	// Optional. The unique name of the revision. If this field is omitted, it
	//  will be automatically generated based on the WorkerPool name.
	// +kcc:proto:field=google.cloud.run.v2.WorkerPoolRevisionTemplate.revision
	Revision *string `json:"revision,omitempty"`

	// Optional. Unstructured key value map that can be used to organize and
	//  categorize objects. User-provided labels are shared with Google's billing
	//  system, so they can be used to filter, or break down billing charges by
	//  team, component, environment, state, etc. For more information, visit
	//  https://cloud.google.com/resource-manager/docs/creating-managing-labels or
	//  https://cloud.google.com/run/docs/configuring/labels.
	// +kcc:proto:field=google.cloud.run.v2.WorkerPoolRevisionTemplate.labels
	Labels map[string]string `json:"labels,omitempty"`

	// Optional. Unstructured key value map that may be set by external tools to
	//  store and arbitrary metadata. They are not queryable and should be
	//  preserved when modifying objects.
	// +kcc:proto:field=google.cloud.run.v2.WorkerPoolRevisionTemplate.annotations
	Annotations map[string]string `json:"annotations,omitempty"`

	// Optional. VPC Access configuration to use for this Revision. For more
	//  information, visit
	//  https://cloud.google.com/run/docs/configuring/connecting-vpc.
	// +kcc:proto:field=google.cloud.run.v2.WorkerPoolRevisionTemplate.vpc_access
	VPCAccess *WorkerPoolVPCAccess `json:"vpcAccess,omitempty"`

	// Optional. Email address of the IAM service account associated with the
	//  revision of the service. The service account represents the identity of the
	//  running revision, and determines what permissions the revision has. If not
	//  provided, the revision will use the project's default service account.
	// +kcc:proto:field=google.cloud.run.v2.WorkerPoolRevisionTemplate.service_account
	ServiceAccountRef *refs.IAMServiceAccountRef `json:"serviceAccountRef,omitempty"`

	// Holds list of the containers that defines the unit of execution for this
	//  Revision.
	// +kcc:proto:field=google.cloud.run.v2.WorkerPoolRevisionTemplate.containers
	Containers []WorkerPoolContainer `json:"containers,omitempty"`

	// Optional. A list of Volumes to make available to containers.
	// +kcc:proto:field=google.cloud.run.v2.WorkerPoolRevisionTemplate.volumes
	Volumes []WorkerPoolVolume `json:"volumes,omitempty"`

	// A reference to a customer managed encryption key (CMEK) to use to encrypt
	//  this container image. For more information, go to
	//  https://cloud.google.com/run/docs/securing/using-cmek
	// +kcc:proto:field=google.cloud.run.v2.WorkerPoolRevisionTemplate.encryption_key
	EncryptionKeyRef *refs.KMSCryptoKeyRef `json:"encryptionKeyRef,omitempty"`

	// Optional. Enables service mesh connectivity.
	// +kcc:proto:field=google.cloud.run.v2.WorkerPoolRevisionTemplate.service_mesh
	ServiceMesh *WorkerPoolServiceMesh `json:"serviceMesh,omitempty"`

	// Optional. The action to take if the encryption key is revoked.
	// +kcc:proto:field=google.cloud.run.v2.WorkerPoolRevisionTemplate.encryption_key_revocation_action
	EncryptionKeyRevocationAction *string `json:"encryptionKeyRevocationAction,omitempty"`

	// Optional. If encryption_key_revocation_action is SHUTDOWN, the duration
	//  before shutting down all instances. The minimum increment is 1 hour.
	// +kcc:proto:field=google.cloud.run.v2.WorkerPoolRevisionTemplate.encryption_key_shutdown_duration
	EncryptionKeyShutdownDuration *string `json:"encryptionKeyShutdownDuration,omitempty"`

	// Optional. The node selector for the revision template.
	// +kcc:proto:field=google.cloud.run.v2.WorkerPoolRevisionTemplate.node_selector
	NodeSelector *WorkerPoolNodeSelector `json:"nodeSelector,omitempty"`

	// Optional. True if GPU zonal redundancy is disabled on this worker pool.
	// +kcc:proto:field=google.cloud.run.v2.WorkerPoolRevisionTemplate.gpu_zonal_redundancy_disabled
	GpuZonalRedundancyDisabled *bool `json:"gpuZonalRedundancyDisabled,omitempty"`
}

// +kcc:proto=google.cloud.run.v2.BinaryAuthorization
type WorkerPoolBinaryAuthorization struct {
	// Optional. If True, indicates to use the default project's binary
	//  authorization policy. If False, binary authorization will be disabled.
	// +kcc:proto:field=google.cloud.run.v2.BinaryAuthorization.use_default
	UseDefault *bool `json:"useDefault,omitempty"`

	// Optional. The path to a binary authorization policy.
	//  Format: `projects/{project}/platforms/cloudRun/{policy-name}`
	// +kcc:proto:field=google.cloud.run.v2.BinaryAuthorization.policy
	Policy *string `json:"policy,omitempty"`

	// Optional. If present, indicates to use Breakglass using this justification.
	//  If use_default is False, then it must be empty.
	//  For more information on breakglass, see
	//  https://cloud.google.com/binary-authorization/docs/using-breakglass
	// +kcc:proto:field=google.cloud.run.v2.BinaryAuthorization.breakglass_justification
	BreakglassJustification *string `json:"breakglassJustification,omitempty"`
}

// +kcc:proto=google.cloud.run.v2.Container
type WorkerPoolContainer struct {
	// Name of the container specified as a DNS_LABEL (RFC 1123).
	// +kcc:proto:field=google.cloud.run.v2.Container.name
	Name *string `json:"name,omitempty"`

	// Required. Name of the container image in Dockerhub, Google Artifact
	//  Registry, or Google Container Registry. If the host is not provided,
	//  Dockerhub is assumed.
	// +kcc:proto:field=google.cloud.run.v2.Container.image
	Image *string `json:"image,omitempty"`

	// Optional. Location of the source.
	// +kcc:proto:field=google.cloud.run.v2.Container.source_code
	SourceCode *WorkerPoolSourceCode `json:"sourceCode,omitempty"`

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
	Env []WorkerPoolEnvVar `json:"env,omitempty"`

	// Compute Resource requirements by this container.
	// +kcc:proto:field=google.cloud.run.v2.Container.resources
	Resources *WorkerPoolResourceRequirements `json:"resources,omitempty"`

	// List of ports to expose from the container. Only a single port can be
	//  specified. The specified ports must be listening on all interfaces
	//  (0.0.0.0) within the container to be accessible.
	//
	//  If omitted, a port number will be chosen and passed to the container
	//  through the PORT environment variable for the container to listen on.
	// +kcc:proto:field=google.cloud.run.v2.Container.ports
	Ports []WorkerPoolContainerPort `json:"ports,omitempty"`

	// Volume to mount into the container's filesystem.
	// +kcc:proto:field=google.cloud.run.v2.Container.volume_mounts
	VolumeMounts []WorkerPoolVolumeMount `json:"volumeMounts,omitempty"`

	// Container's working directory.
	//  If not specified, the container runtime's default will be used, which
	//  might be configured in the container image.
	// +kcc:proto:field=google.cloud.run.v2.Container.working_dir
	WorkingDir *string `json:"workingDir,omitempty"`

	// Periodic probe of container liveness.
	//  Container will be restarted if the probe fails.
	// +kcc:proto:field=google.cloud.run.v2.Container.liveness_probe
	LivenessProbe *WorkerPoolProbe `json:"livenessProbe,omitempty"`

	// Startup probe of application within the container.
	//  All other probes are disabled if a startup probe is provided, until it
	//  succeeds. Container will not be added to service endpoints if the probe
	//  fails.
	// +kcc:proto:field=google.cloud.run.v2.Container.startup_probe
	StartupProbe *WorkerPoolProbe `json:"startupProbe,omitempty"`

	// Readiness probe to be used for health checks.
	// +kcc:proto:field=google.cloud.run.v2.Container.readiness_probe
	// ReadinessProbe *WorkerPoolProbe `json:"readinessProbe,omitempty"`

	// Names of the containers that must start before this container.
	// +kcc:proto:field=google.cloud.run.v2.Container.depends_on
	DependsOn []string `json:"dependsOn,omitempty"`

	// Base image for this container. Only supported for services. If set, it
	//  indicates that the service is enrolled into automatic base image update.
	// +kcc:proto:field=google.cloud.run.v2.Container.base_image_uri
	BaseImageURI *string `json:"baseImageURI,omitempty"`
}

// +kcc:proto=google.cloud.run.v2.EnvVar
type WorkerPoolEnvVar struct {
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
	ValueSource *WorkerPoolEnvVarSource `json:"valueSource,omitempty"`
}

// +kcc:proto=google.cloud.run.v2.EnvVarSource
type WorkerPoolEnvVarSource struct {
	// Selects a secret and a specific version from Cloud Secret Manager.
	// +kcc:proto:field=google.cloud.run.v2.EnvVarSource.secret_key_ref
	SecretKeyRef *WorkerPoolSecretKeySelector `json:"secretKeyRef,omitempty"`
}

// +kcc:proto=google.cloud.run.v2.SecretKeySelector
type WorkerPoolSecretKeySelector struct {
	// Required. The name of the secret in Cloud Secret Manager.
	// +kcc:proto:field=google.cloud.run.v2.SecretKeySelector.secret
	SecretRef *secretmanagerv1beta1.SecretRef `json:"secretRef,omitempty"`

	// The Cloud Secret Manager secret version.
	//  Can be 'latest' for the latest version, an integer for a specific version,
	//  or a version alias.
	// +kcc:proto:field=google.cloud.run.v2.SecretKeySelector.version
	VersionRef *secretmanagerv1beta1.SecretVersionRef `json:"versionRef,omitempty"`
}

// +kcc:proto=google.cloud.run.v2.Volume
type WorkerPoolVolume struct {
	// Required. Volume's name.
	// +kcc:proto:field=google.cloud.run.v2.Volume.name
	Name *string `json:"name,omitempty"`

	// Secret represents a secret that should populate this volume.
	// +kcc:proto:field=google.cloud.run.v2.Volume.secret
	Secret *WorkerPoolSecretVolumeSource `json:"secret,omitempty"`

	// For Cloud SQL volumes, contains the specific instances that should be
	//  mounted. Visit https://cloud.google.com/sql/docs/mysql/connect-run for
	//  more information on how to connect Cloud SQL and Cloud Run.
	// +kcc:proto:field=google.cloud.run.v2.Volume.cloud_sql_instance
	CloudSQLInstance *WorkerPoolCloudSQLInstance `json:"cloudSQLInstance,omitempty"`

	// Ephemeral storage used as a shared volume.
	// +kcc:proto:field=google.cloud.run.v2.Volume.empty_dir
	EmptyDir *WorkerPoolEmptyDirVolumeSource `json:"emptyDir,omitempty"`

	// For NFS Voumes, contains the path to the nfs Volume
	// +kcc:proto:field=google.cloud.run.v2.Volume.nfs
	Nfs *WorkerPoolNFSVolumeSource `json:"nfs,omitempty"`

	// Persistent storage backed by a Google Cloud Storage bucket.
	// +kcc:proto:field=google.cloud.run.v2.Volume.gcs
	GCS *WorkerPoolGCSVolumeSource `json:"gcs,omitempty"`
}

// +kcc:proto=google.cloud.run.v2.SecretVolumeSource
type WorkerPoolSecretVolumeSource struct {
	// Required. The name of the secret in Cloud Secret Manager.
	// +kcc:proto:field=google.cloud.run.v2.SecretVolumeSource.secret
	SecretRef *secretmanagerv1beta1.SecretRef `json:"secretRef,omitempty"`

	// If unspecified, the volume will expose a file whose name is the
	//  secret, relative to VolumeMount.mount_path + VolumeMount.sub_path.
	//  If specified, the key will be used as the version to fetch from Cloud
	//  Secret Manager and the path will be the name of the file exposed in the
	//  volume. When items are defined, they must specify a path and a version.
	// +kcc:proto:field=google.cloud.run.v2.SecretVolumeSource.items
	Items []WorkerPoolVersionToPath `json:"items,omitempty"`

	// Integer representation of mode bits to use on created files by default.
	//  Must be a value between 0000 and 0777 (octal), defaulting to 0444.
	// +kcc:proto:field=google.cloud.run.v2.SecretVolumeSource.default_mode
	DefaultMode *int32 `json:"defaultMode,omitempty"`
}

// +kcc:proto=google.cloud.run.v2.CloudSqlInstance
type WorkerPoolCloudSQLInstance struct {
	// The Cloud SQL instance connection names, as can be found in
	//  https://console.cloud.google.com/sql/instances. Visit
	//  https://cloud.google.com/sql/docs/mysql/connect-run for more information on
	//  how to connect Cloud SQL and Cloud Run. Format:
	//  {project}:{location}:{instance}
	// +kcc:proto:field=google.cloud.run.v2.CloudSqlInstance.instances
	InstanceRefs []*refs.SQLInstanceRef `json:"instanceRefs,omitempty"`
}

// +kcc:proto=google.cloud.run.v2.VersionToPath
type WorkerPoolVersionToPath struct {
	// Required. The relative path of the secret in the container.
	// +kcc:proto:field=google.cloud.run.v2.VersionToPath.path
	Path *string `json:"path,omitempty"`

	// The Cloud Secret Manager secret version.
	//  Can be 'latest' for the latest value, or an integer or a secret alias for a
	//  specific version.
	// +kcc:proto:field=google.cloud.run.v2.VersionToPath.version
	VersionRef *secretmanagerv1beta1.SecretVersionRef `json:"versionRef,omitempty"`

	// Integer octal mode bits to use on this file, must be a value between
	//  01 and 0777 (octal). If 0 or not set, the Volume's default mode will be
	//  used.
	// +kcc:proto:field=google.cloud.run.v2.VersionToPath.mode
	Mode *int32 `json:"mode,omitempty"`
}

// +kcc:proto=google.cloud.run.v2.VpcAccess
type WorkerPoolVPCAccess struct {
	// VPC Access connector name.
	//  Format: `projects/{project}/locations/{location}/connectors/{connector}`,
	//  where `{project}` can be project id or number.
	// +kcc:proto:field=google.cloud.run.v2.VpcAccess.connector
	ConnectorRef *vpcaccessv1beta1.VPCAccessConnectorRef `json:"connectorRef,omitempty"`

	// Optional. Traffic VPC egress settings. If not provided, it defaults to
	//  PRIVATE_RANGES_ONLY.
	// +kcc:proto:field=google.cloud.run.v2.VpcAccess.egress
	Egress *string `json:"egress,omitempty"`

	// Optional. Direct VPC egress settings. Currently only single network
	//  interface is supported.
	// +kcc:proto:field=google.cloud.run.v2.VpcAccess.network_interfaces
	NetworkInterfaces []WorkerPoolVPCAccessNetworkInterface `json:"networkInterfaces,omitempty"`
}

// +kcc:proto=google.cloud.run.v2.VpcAccess.NetworkInterface
type WorkerPoolVPCAccessNetworkInterface struct {
	// Optional. The VPC network that the Cloud Run resource will be able to
	//  send traffic to. At least one of network or subnetwork must be specified.
	// +kcc:proto:field=google.cloud.run.v2.VpcAccess.NetworkInterface.network
	NetworkRef *computev1beta1.ComputeNetworkRef `json:"networkRef,omitempty"`

	// Optional. The VPC subnetwork that the Cloud Run resource will get IPs
	//  from. At least one of network or subnetwork must be specified.
	// +kcc:proto:field=google.cloud.run.v2.VpcAccess.NetworkInterface.subnetwork
	SubnetworkRef *refs.ComputeSubnetworkRef `json:"subnetworkRef,omitempty"`

	// Optional. Network tags applied to this Cloud Run resource.
	// +kcc:proto:field=google.cloud.run.v2.VpcAccess.NetworkInterface.tags
	Tags []string `json:"tags,omitempty"`
}

// +kcc:proto=google.cloud.run.v2.Condition
type WorkerPoolCondition struct {
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

// +kcc:observedstate:proto=google.cloud.run.v2.WorkerPoolRevisionTemplate
type WorkerPoolRevisionTemplateObservedState struct {
	// Holds list of the containers that defines the unit of execution for this
	//  Revision.
	// +kcc:proto:field=google.cloud.run.v2.WorkerPoolRevisionTemplate.containers
	Containers []WorkerPoolContainerObservedState `json:"containers,omitempty"`
}

// +kcc:observedstate:proto=google.cloud.run.v2.Container
type WorkerPoolContainerObservedState struct {
	// Output only. The build info of the container image.
	// +kcc:proto:field=google.cloud.run.v2.Container.build_info
	BuildInfo *WorkerPoolBuildInfoObservedState `json:"buildInfo,omitempty"`
}

// +kcc:proto=google.cloud.run.v2.ContainerPort
type WorkerPoolContainerPort struct {
	// If specified, used to specify which protocol to use.
	//  Allowed values are "http1" and "h2c".
	// +kcc:proto:field=google.cloud.run.v2.ContainerPort.name
	Name *string `json:"name,omitempty"`

	// Port number the container listens on.
	//  This must be a valid TCP port number, 0 < container_port < 65536.
	// +kcc:proto:field=google.cloud.run.v2.ContainerPort.container_port
	ContainerPort *int32 `json:"containerPort,omitempty"`
}

// +kcc:proto=google.cloud.run.v2.SourceCode
type WorkerPoolSourceCode struct {
	// The source is a Cloud Storage bucket.
	// +kcc:proto:field=google.cloud.run.v2.SourceCode.cloud_storage_source
	CloudStorageSource *WorkerPoolSourceCodeCloudStorageSource `json:"cloudStorageSource,omitempty"`
}

// +kcc:proto=google.cloud.run.v2.SourceCode.CloudStorageSource
type WorkerPoolSourceCodeCloudStorageSource struct {
	// Required. The Cloud Storage bucket name.
	// +kcc:proto:field=google.cloud.run.v2.SourceCode.CloudStorageSource.bucket
	Bucket *string `json:"bucket,omitempty"`

	// Required. The Cloud Storage object name.
	// +kcc:proto:field=google.cloud.run.v2.SourceCode.CloudStorageSource.object
	Object *string `json:"object,omitempty"`

	// Optional. The Cloud Storage object generation.
	// +kcc:proto:field=google.cloud.run.v2.SourceCode.CloudStorageSource.generation
	Generation *int64 `json:"generation,omitempty"`
}

// +kcc:proto=google.cloud.run.v2.EmptyDirVolumeSource
type WorkerPoolEmptyDirVolumeSource struct {
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
type WorkerPoolNFSVolumeSource struct {
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

// +kcc:proto=google.cloud.run.v2.GCSVolumeSource
type WorkerPoolGCSVolumeSource struct {
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

// +kcc:proto=google.cloud.run.v2.ServiceMesh
type WorkerPoolServiceMesh struct {
	// The Mesh resource name. Format:
	//  `projects/{project}/locations/global/meshes/{mesh}`, where `{project}` can
	//  be project id or number.
	// +kcc:proto:field=google.cloud.run.v2.ServiceMesh.mesh
	Mesh *string `json:"mesh,omitempty"`
}

// +kcc:proto=google.cloud.run.v2.NodeSelector
type WorkerPoolNodeSelector struct {
	// Required. GPU accelerator type to attach to an instance.
	// +kcc:proto:field=google.cloud.run.v2.NodeSelector.accelerator
	Accelerator *string `json:"accelerator,omitempty"`
}

// +kcc:proto=google.cloud.run.v2.ResourceRequirements
type WorkerPoolResourceRequirements struct {
	// Only `memory`, `cpu` and `nvidia.com/gpu` keys in the map are supported.
	//
	//  <p>Notes:
	//   * The only supported values for CPU are '1', '2', '4', and '8'. Setting 4
	//  CPU requires at least 2Gi of memory. For more information, go to
	//  https://cloud.google.com/run/docs/configuring/cpu.
	//    * For supported 'memory' values and syntax, go to
	//   https://cloud.google.com/run/docs/configuring/memory-limits
	//   * The only supported 'nvidia.com/gpu' value is '1'.
	// +kcc:proto:field=google.cloud.run.v2.ResourceRequirements.limits
	Limits map[string]string `json:"limits,omitempty"`

	// Determines whether CPU is only allocated during requests (true by default).
	//  However, if ResourceRequirements is set, the caller must explicitly
	//  set this field to true to preserve the default behavior.
	// +kcc:proto:field=google.cloud.run.v2.ResourceRequirements.cpu_idle
	CPUIdle *bool `json:"cpuIdle,omitempty"`

	// Determines whether CPU should be boosted on startup of a new container
	//  instance above the requested CPU threshold, this can help reduce cold-start
	//  latency.
	// +kcc:proto:field=google.cloud.run.v2.ResourceRequirements.startup_cpu_boost
	StartupCPUBoost *bool `json:"startupCPUBoost,omitempty"`
}

// +kcc:proto=google.cloud.run.v2.VolumeMount
type WorkerPoolVolumeMount struct {
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

	// Optional. Path within the volume from which the container's volume should
	//  be mounted. Defaults to "" (volume's root).
	// +kcc:proto:field=google.cloud.run.v2.VolumeMount.sub_path
	SubPath *string `json:"subPath,omitempty"`
}

// +kcc:proto=google.cloud.run.v2.Probe
type WorkerPoolProbe struct {
	// Optional. Number of seconds after the container has started before the
	//  probe is initiated. Defaults to 0 seconds. Minimum value is 0. Maximum
	//  value for liveness probe is 3600. Maximum value for startup probe is 240.
	// +kcc:proto:field=google.cloud.run.v2.Probe.initial_delay_seconds
	InitialDelaySeconds *int32 `json:"initialDelaySeconds,omitempty"`

	// Optional. Number of seconds after which the probe times out.
	//  Defaults to 1 second. Minimum value is 1. Maximum value is 3600.
	//  Must be smaller than period_seconds.
	// +kcc:proto:field=google.cloud.run.v2.Probe.period_seconds
	PeriodSeconds *int32 `json:"periodSeconds,omitempty"`

	// Optional. Minimum consecutive failures for the probe to be considered
	//  failed after having succeeded. Defaults to 3. Minimum value is 1.
	// +kcc:proto:field=google.cloud.run.v2.Probe.failure_threshold
	FailureThreshold *int32 `json:"failureThreshold,omitempty"`

	// Optional. HTTPGet specifies the http request to perform.
	//  Exactly one of httpGet, tcpSocket, or grpc must be specified.
	// +kcc:proto:field=google.cloud.run.v2.Probe.http_get
	HTTPGet *WorkerPoolHTTPGetAction `json:"httpGet,omitempty"`

	// Optional. TCPSocket specifies an action involving a TCP port.
	//  Exactly one of httpGet, tcpSocket, or grpc must be specified.
	// +kcc:proto:field=google.cloud.run.v2.Probe.tcp_socket
	TCPSocket *WorkerPoolTCPSocketAction `json:"tcpSocket,omitempty"`

	// Optional. GRPC specifies an action involving a gRPC port.
	//  Exactly one of httpGet, tcpSocket, or grpc must be specified.
	// +kcc:proto:field=google.cloud.run.v2.Probe.grpc
	Grpc *WorkerPoolGRPCAction `json:"grpc,omitempty"`
}

// +kcc:proto=google.cloud.run.v2.HTTPGetAction
type WorkerPoolHTTPGetAction struct {
	// Optional. Path to access on the HTTP server. Defaults to '/'.
	// +kcc:proto:field=google.cloud.run.v2.HTTPGetAction.path
	Path *string `json:"path,omitempty"`

	// Optional. Custom headers to set in the request. HTTP allows repeated
	//  headers.
	// +kcc:proto:field=google.cloud.run.v2.HTTPGetAction.http_headers
	HTTPHeaders []WorkerPoolHTTPHeader `json:"httpHeaders,omitempty"`

	// Optional. Port number to access on the container. Must be in the range 1 to
	//  65535. If not specified, defaults to the exposed port of the container,
	//  which is the value of container.ports[0].containerPort.
	// +kcc:proto:field=google.cloud.run.v2.HTTPGetAction.port
	Port *int32 `json:"port,omitempty"`
}

// +kcc:proto=google.cloud.run.v2.TCPSocketAction
type WorkerPoolTCPSocketAction struct {
	// Optional. Port number to access on the container. Must be in the range 1 to
	//  65535. If not specified, defaults to the exposed port of the container,
	//  which is the value of container.ports[0].containerPort.
	// +kcc:proto:field=google.cloud.run.v2.TCPSocketAction.port
	Port *int32 `json:"port,omitempty"`
}

// +kcc:proto=google.cloud.run.v2.GRPCAction
type WorkerPoolGRPCAction struct {
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

// +kcc:observedstate:proto=google.cloud.run.v2.BuildInfo
type WorkerPoolBuildInfoObservedState struct {
	// Output only. Entry point of the function when the image is a Cloud Run
	//  function.
	// +kcc:proto:field=google.cloud.run.v2.BuildInfo.function_target
	FunctionTarget *string `json:"functionTarget,omitempty"`

	// Output only. Source code location of the image.
	// +kcc:proto:field=google.cloud.run.v2.BuildInfo.source_location
	SourceLocation *string `json:"sourceLocation,omitempty"`
}

// +kcc:proto=google.cloud.run.v2.InstanceSplit
type WorkerPoolInstanceSplit struct {
	// The allocation type for this instance split.
	// +kcc:proto:field=google.cloud.run.v2.InstanceSplit.type
	Type *string `json:"type,omitempty"`

	// Revision to which to assign this portion of instances, if split allocation
	//  is by revision.
	// +kcc:proto:field=google.cloud.run.v2.InstanceSplit.revision
	Revision *string `json:"revision,omitempty"`

	// Specifies percent of the instance split to this Revision.
	//  This defaults to zero if unspecified.
	// +kcc:proto:field=google.cloud.run.v2.InstanceSplit.percent
	Percent *int32 `json:"percent,omitempty"`
}

// +kcc:proto=google.cloud.run.v2.InstanceSplitStatus
type WorkerPoolInstanceSplitStatus struct {
	// The allocation type for this instance split.
	// +kcc:proto:field=google.cloud.run.v2.InstanceSplitStatus.type
	Type *string `json:"type,omitempty"`

	// Revision to which this instance split is assigned.
	// +kcc:proto:field=google.cloud.run.v2.InstanceSplitStatus.revision
	Revision *string `json:"revision,omitempty"`

	// Specifies percent of the instance split to this Revision.
	// +kcc:proto:field=google.cloud.run.v2.InstanceSplitStatus.percent
	Percent *int32 `json:"percent,omitempty"`
}

// +kcc:proto=google.cloud.run.v2.WorkerPoolScaling
type WorkerPoolScaling struct {
	// Optional. The total number of instances in manual scaling mode.
	// +kcc:proto:field=google.cloud.run.v2.WorkerPoolScaling.manual_instance_count
	ManualInstanceCount *int32 `json:"manualInstanceCount,omitempty"`
}

// +kcc:proto=google.cloud.run.v2.HTTPHeader
type WorkerPoolHTTPHeader struct {
	// Required. The header field name
	// +kcc:proto:field=google.cloud.run.v2.HTTPHeader.name
	Name *string `json:"name,omitempty"`

	// Optional. The header field value
	// +kcc:proto:field=google.cloud.run.v2.HTTPHeader.value
	Value *string `json:"value,omitempty"`
}
