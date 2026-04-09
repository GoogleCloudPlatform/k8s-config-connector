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
	storagev1beta1 "github.com/GoogleCloudPlatform/k8s-config-connector/apis/storage/v1beta1"
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
	Description *string `json:"description,omitempty"`

	// Optional. Unstructured key value map that may be set by external tools to
	//  store and arbitrary metadata. They are not queryable and should be
	//  preserved when modifying objects.
	Annotations map[string]string `json:"annotations,omitempty"`

	// Arbitrary identifier for the API client.
	Client *string `json:"client,omitempty"`

	// Arbitrary version identifier for the API client.
	ClientVersion *string `json:"clientVersion,omitempty"`

	// Optional. The launch stage as defined by [Google Cloud Platform
	//   Launch Stages](https://cloud.google.com/terms/launch-stages).
	LaunchStage *string `json:"launchStage,omitempty"`

	// Optional. Settings for the Binary Authorization feature.
	BinaryAuthorization *BinaryAuthorization `json:"binaryAuthorization,omitempty"`

	// Required. The template used to create revisions for this WorkerPool.
	Template *WorkerPoolRevisionTemplate `json:"template,omitempty"`

	// Optional. Specifies how to distribute instances over a collection of
	//  Revisions belonging to the WorkerPool. If instance split is empty or not
	//  provided, defaults to 100% instances assigned to the latest `Ready`
	//  Revision.
	InstanceSplits []InstanceSplit `json:"instanceSplits,omitempty"`

	// Optional. Specifies worker-pool-level scaling settings
	Scaling *WorkerPoolScaling `json:"scaling,omitempty"`

	// One or more custom audiences that you want this worker pool to support.
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
	Uid *string `json:"uid,omitempty"`

	// Output only. A number that monotonically increases every time the user
	//  modifies the desired state.
	Generation *int64 `json:"generation,omitempty"`

	// Output only. The creation time.
	CreateTime *string `json:"createTime,omitempty"`

	// Output only. The last-modified time.
	UpdateTime *string `json:"updateTime,omitempty"`

	// Output only. The deletion time. It is only populated as a response to a
	//  Delete request.
	DeleteTime *string `json:"deleteTime,omitempty"`

	// Output only. For a deleted resource, the time after which it will be
	//  permanently deleted.
	ExpireTime *string `json:"expireTime,omitempty"`

	// Output only. Email address of the authenticated creator.
	Creator *string `json:"creator,omitempty"`

	// Output only. Email address of the last authenticated modifier.
	LastModifier *string `json:"lastModifier,omitempty"`

	// Required. The template used to create revisions for this WorkerPool.
	Template *WorkerPoolRevisionTemplateObservedState `json:"template,omitempty"`

	// Output only. The generation of this WorkerPool currently serving traffic.
	ObservedGeneration *int64 `json:"observedGeneration,omitempty"`

	// Output only. The Condition of this WorkerPool, containing its readiness
	//  status, and detailed error information in case it did not reach a serving
	//  state.
	TerminalCondition *Condition `json:"terminalCondition,omitempty"`

	// Output only. The Conditions of all other associated sub-resources. They
	//  contain additional diagnostics information in case the WorkerPool does not
	//  reach its Serving state.
	Conditions []Condition `json:"conditions,omitempty"`

	// Output only. Name of the latest revision that is serving traffic.
	LatestReadyRevision *string `json:"latestReadyRevision,omitempty"`

	// Output only. Name of the last created revision.
	LatestCreatedRevision *string `json:"latestCreatedRevision,omitempty"`

	// Output only. Detailed status information for corresponding instance splits.
	InstanceSplitStatuses []InstanceSplitStatus `json:"instanceSplitStatuses,omitempty"`

	// Output only. Reserved for future use.
	SatisfiesPzs *bool `json:"satisfiesPzs,omitempty"`

	// Output only. Returns true if the WorkerPool is currently being acted upon
	//  by the system to bring it into the desired state.
	Reconciling *bool `json:"reconciling,omitempty"`

	// Output only. A system-generated fingerprint for this version of the
	//  resource. May be used to detect modification conflict during updates.
	Etag *string `json:"etag,omitempty"`
}

// +kcc:proto=google.cloud.run.v2.WorkerPoolRevisionTemplate
type WorkerPoolRevisionTemplate struct {
	// Optional. The unique name for the revision. If this field is omitted, it
	//  will be automatically generated based on the WorkerPool name.
	Revision *string `json:"revision,omitempty"`

	// Optional. Unstructured key value map that may be set by external tools to
	//  store and arbitrary metadata.
	Annotations map[string]string `json:"annotations,omitempty"`

	// Optional. VPC Access configuration to use for this Revision.
	VPCAccess *VPCAccess `json:"vpcAccess,omitempty"`

	// Optional. Email address of the IAM service account associated with the
	//  revision of the service.
	ServiceAccountRef *refs.IAMServiceAccountRef `json:"serviceAccountRef,omitempty"`

	// Holds list of the containers that defines the unit of execution for this
	//  Revision.
	Containers []Container `json:"containers,omitempty"`

	// Optional. A list of Volumes to make available to containers.
	Volumes []Volume `json:"volumes,omitempty"`

	// A reference to a customer managed encryption key (CMEK) to use to encrypt
	//  this container image.
	EncryptionKeyRef *refs.KMSCryptoKeyRef `json:"encryptionKeyRef,omitempty"`

	// Optional. Enables service mesh connectivity.
	ServiceMesh *ServiceMesh `json:"serviceMesh,omitempty"`

	// Optional. The action to take if the encryption key is revoked.
	EncryptionKeyRevocationAction *string `json:"encryptionKeyRevocationAction,omitempty"`

	// Optional. If encryption_key_revocation_action is SHUTDOWN, the duration
	//  before shutting down all instances.
	EncryptionKeyShutdownDuration *string `json:"encryptionKeyShutdownDuration,omitempty"`

	// Optional. The node selector for the revision template.
	NodeSelector *NodeSelector `json:"nodeSelector,omitempty"`
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

// +kcc:proto=google.cloud.run.v2.VpcAccess
type VPCAccess struct {
	// VPC Access connector name.
	ConnectorRef *vpcaccessv1beta1.VPCAccessConnectorRef `json:"connectorRef,omitempty"`

	// Optional. Traffic VPC egress settings.
	Egress *string `json:"egress,omitempty"`

	// Optional. Direct VPC egress settings. Currently only single network
	//  interface is supported.
	NetworkInterfaces []VPCAccess_NetworkInterface `json:"networkInterfaces,omitempty"`
}

// +kcc:proto=google.cloud.run.v2.VpcAccess.NetworkInterface
type VPCAccess_NetworkInterface struct {
	// Optional. The VPC network that the Cloud Run resource will be able to
	//  send traffic to.
	NetworkRef *computev1beta1.ComputeNetworkRef `json:"networkRef,omitempty"`

	// Optional. The VPC subnetwork that the Cloud Run resource will get IPs
	//  from.
	SubnetworkRef *refs.ComputeSubnetworkRef `json:"subnetworkRef,omitempty"`

	// Optional. Network tags applied to this Cloud Run resource.
	Tags []string `json:"tags,omitempty"`
}

// +kcc:proto=google.cloud.run.v2.Container
type Container struct {
	// Name of the container specified as a DNS_LABEL (RFC 1123).
	Name *string `json:"name,omitempty"`

	// Required. Name of the container image in Dockerhub, Google Artifact
	//  Registry, or Google Container Registry.
	Image *string `json:"image,omitempty"`

	// Entrypoint array. Not executed within a shell.
	Command []string `json:"command,omitempty"`

	// Arguments to the entrypoint.
	Args []string `json:"args,omitempty"`

	// List of environment variables to set in the container.
	Env []EnvVar `json:"env,omitempty"`

	// Compute Resource requirements by this container.
	Resources *ResourceRequirements `json:"resources,omitempty"`

	// List of ports to expose from the container.
	Ports []ContainerPort `json:"ports,omitempty"`

	// Volume to mount into the container's filesystem.
	VolumeMounts []VolumeMount `json:"volumeMounts,omitempty"`

	// Container's working directory.
	WorkingDir *string `json:"workingDir,omitempty"`

	// Periodic probe of container liveness.
	LivenessProbe *Probe `json:"livenessProbe,omitempty"`

	// Startup probe of application within the container.
	StartupProbe *Probe `json:"startupProbe,omitempty"`

	// Names of the containers that must start before this container.
	DependsOn []string `json:"dependsOn,omitempty"`

	// Base image for this container. Only supported for services.
	BaseImageURI *string `json:"baseImageURI,omitempty"`
}

// +kcc:proto=google.cloud.run.v2.EnvVar
type EnvVar struct {
	// Required. Name of the environment variable. Must not exceed 32768
	//  characters.
	Name *string `json:"name,omitempty"`

	// Literal value of the environment variable.
	Value *string `json:"value,omitempty"`

	// Source for the environment variable's value.
	ValueSource *EnvVarSource `json:"valueSource,omitempty"`
}

// +kcc:proto=google.cloud.run.v2.EnvVarSource
type EnvVarSource struct {
	// Selects a secret and a specific version from Cloud Secret Manager.
	SecretKeyRef *SecretKeySelector `json:"secretKeyRef,omitempty"`
}

// +kcc:proto=google.cloud.run.v2.SecretKeySelector
type SecretKeySelector struct {
	// Required. The name of the secret in Cloud Secret Manager.
	SecretRef *secretmanagerv1beta1.SecretRef `json:"secretRef,omitempty"`

	// The Cloud Secret Manager secret version.
	Version *string `json:"version,omitempty"`
}

// +kcc:proto=google.cloud.run.v2.Volume
type Volume struct {
	// Required. Volume's name.
	Name *string `json:"name,omitempty"`

	// Secret represents a secret that should populate this volume.
	Secret *SecretVolumeSource `json:"secret,omitempty"`

	// For Cloud SQL volumes, contains the specific instances that should be
	//  mounted.
	CloudSQLInstance *CloudSQLInstance `json:"cloudSQLInstance,omitempty"`

	// Ephemeral storage used as a shared volume.
	EmptyDir *EmptyDirVolumeSource `json:"emptyDir,omitempty"`

	// For NFS Voumes, contains the path to the nfs Volume
	Nfs *NfsVolumeSource `json:"nfs,omitempty"`

	// Persistent storage backed by a Google Cloud Storage bucket.
	GCS *GCSVolumeSource `json:"gcs,omitempty"`
}

// +kcc:proto=google.cloud.run.v2.SecretVolumeSource
type SecretVolumeSource struct {
	// Required. The name of the secret in Cloud Secret Manager.
	SecretRef *secretmanagerv1beta1.SecretRef `json:"secretRef,omitempty"`

	// If unspecified, the volume will expose a file whose name is the
	//  secret, relative to VolumeMount.mount_path.
	Items []VersionToPath `json:"items,omitempty"`

	// Integer representation of mode bits to use on created files by default.
	DefaultMode *int32 `json:"defaultMode,omitempty"`
}

// +kcc:proto=google.cloud.run.v2.GCSVolumeSource
type GCSVolumeSource struct {
	// Cloud Storage Bucket name.
	BucketRef *storagev1beta1.StorageBucketRef `json:"bucketRef,omitempty"`

	// If true, the volume will be mounted as read only for all mounts.
	ReadOnly *bool `json:"readOnly,omitempty"`

	// A list of additional flags to pass to the gcsfuse CLI.
	MountOptions []string `json:"mountOptions,omitempty"`
}

// +kcc:proto=google.cloud.run.v2.CloudSqlInstance
type CloudSQLInstance struct {
	// The Cloud SQL instance connection names.
	InstanceRefs []*refs.SQLInstanceRef `json:"instanceRefs,omitempty"`
}

// +kcc:observedstate:proto=google.cloud.run.v2.WorkerPoolRevisionTemplate
type WorkerPoolRevisionTemplateObservedState struct {
	// Holds list of the containers that defines the unit of execution for this
	//  Revision.
	Containers []ContainerObservedState `json:"containers,omitempty"`
}

// +kcc:observedstate:proto=google.cloud.run.v2.Container
type ContainerObservedState struct {
	// Output only. The build info of the container image.
	BuildInfo *BuildInfoObservedState `json:"buildInfo,omitempty"`
}

// +kcc:observedstate:proto=google.cloud.run.v2.BuildInfo
type BuildInfoObservedState struct {
	// Output only. Entry point of the function when the image is a Cloud Run
	//  function.
	FunctionTarget *string `json:"functionTarget,omitempty"`

	// Output only. Source code location of the image.
	SourceLocation *string `json:"sourceLocation,omitempty"`
}

// +kcc:proto=google.cloud.run.v2.BinaryAuthorization
type BinaryAuthorization struct {
	// Optional. If True, indicates to use the default project's binary
	//  authorization policy. If False, binary authorization will be disabled.
	UseDefault *bool `json:"useDefault,omitempty"`

	// Optional. The path to a binary authorization policy.
	//  Format: `projects/{project}/platforms/cloudRun/{policy-name}`
	//  Commented out because this field refers to Binary Authorization Platform Policies,
	//  which are not yet supported as a KRM resource in KCC.
	// Policy *string `json:"policy,omitempty"`

	// Optional. If present, indicates to use Breakglass using this justification.
	//  If use_default is False, then it must be empty.
	//  For more information on breakglass, see
	//  https://cloud.google.com/binary-authorization/docs/using-breakglass
	BreakglassJustification *string `json:"breakglassJustification,omitempty"`
}
