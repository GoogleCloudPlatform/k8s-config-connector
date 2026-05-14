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
	refsv1beta1 "github.com/GoogleCloudPlatform/k8s-config-connector/apis/refs/v1beta1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/apis/k8s/v1alpha1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

var VertexAICustomJobGVK = GroupVersion.WithKind("VertexAICustomJob")

// VertexAICustomJobSpec defines the desired state of VertexAICustomJob
// +kcc:spec:proto=google.cloud.aiplatform.v1.CustomJob
type VertexAICustomJobSpec struct {
	// The project that this resource belongs to.
	ProjectRef *refsv1beta1.ProjectRef `json:"projectRef"`

	// The location of this resource.
	Location string `json:"location"`

	// The VertexAICustomJob name. If not given, the metadata.name will be used.
	ResourceID *string `json:"resourceID,omitempty"`

	// Required. The display name of the CustomJob.
	// The name can be up to 128 characters long and can consist of any UTF-8
	// characters.
	// +kcc:proto:field=google.cloud.aiplatform.v1.CustomJob.display_name
	// +kubebuilder:validation:Required
	DisplayName *string `json:"displayName"`

	// Required. Job spec.
	// +kcc:proto:field=google.cloud.aiplatform.v1.CustomJob.job_spec
	// +kubebuilder:validation:Required
	JobSpec *CustomJobSpec `json:"jobSpec"`

	// Customer-managed encryption key options for a CustomJob. If this is set,
	// then all resources created by the CustomJob will be encrypted with the
	// provided encryption key.
	// +kcc:proto:field=google.cloud.aiplatform.v1.CustomJob.encryption_spec
	EncryptionSpec *CustomJobEncryptionSpec `json:"encryptionSpec,omitempty"`

	// The labels with user-defined metadata to organize CustomJobs.
	// +kcc:proto:field=google.cloud.aiplatform.v1.CustomJob.labels
	Labels map[string]string `json:"labels,omitempty"`
}

// +kcc:proto=google.cloud.aiplatform.v1.EncryptionSpec
type CustomJobEncryptionSpec struct {
	// Required. The Cloud KMS resource identifier of the customer managed
	// encryption key used to protect a resource.
	// +kcc:proto:field=google.cloud.aiplatform.v1.EncryptionSpec.kms_key_name
	// +kubebuilder:validation:Required
	KMSKeyRef *refsv1beta1.KMSCryptoKeyRef `json:"kmsKeyRef"`
}

// Represents the spec of a CustomJob.
// +kcc:proto=google.cloud.aiplatform.v1.CustomJobSpec
type CustomJobSpec struct {
	// Optional. The ID of the PersistentResource in the same Project and Location
	// which to run
	// +kcc:proto:field=google.cloud.aiplatform.v1.CustomJobSpec.persistent_resource_id
	// +optional
	PersistentResourceRef *VertexAIPersistentResourceRef `json:"persistentResourceRef,omitempty"`

	// Required. The spec of the worker pools including machine type and Docker
	// image. All worker pools except the first one are optional and can be
	// skipped by providing an empty value.
	// +kcc:proto:field=google.cloud.aiplatform.v1.CustomJobSpec.worker_pool_specs
	// +kubebuilder:validation:Required
	CustomJobWorkerPoolSpecs []CustomJobWorkerPoolSpec `json:"workerPoolSpecs"`

	// CustomJobScheduling options for a CustomJob.
	// +kcc:proto:field=google.cloud.aiplatform.v1.CustomJobSpec.scheduling
	CustomJobScheduling *CustomJobScheduling `json:"scheduling,omitempty"`

	// Specifies the service account for workload run-as account.
	// +kcc:proto:field=google.cloud.aiplatform.v1.CustomJobSpec.service_account
	ServiceAccountRef *refsv1beta1.IAMServiceAccountRef `json:"serviceAccountRef,omitempty"`

	// Optional. The full name of the Compute Engine
	// network to which the Job should be peered.
	// +kcc:proto:field=google.cloud.aiplatform.v1.CustomJobSpec.network
	// +optional
	NetworkRef *computev1beta1.ComputeNetworkRef `json:"networkRef,omitempty"`

	// Optional. A list of names for the reserved ip ranges under the VPC network
	// that can be used for this job.
	// +kcc:proto:field=google.cloud.aiplatform.v1.CustomJobSpec.reserved_ip_ranges
	// +optional
	ReservedIPRanges []string `json:"reservedIPRanges,omitempty"`

	// The Cloud Storage location to store the output of this CustomJob or
	// HyperparameterTuningJob.
	// +kcc:proto:field=google.cloud.aiplatform.v1.CustomJobSpec.base_output_directory
	BaseOutputDirectory *CustomJobGcsDestination `json:"baseOutputDirectory,omitempty"`

	// The ID of the location to store protected artifacts. e.g. us-central1.
	// +kcc:proto:field=google.cloud.aiplatform.v1.CustomJobSpec.protected_artifact_location_id
	ProtectedArtifactLocationID *string `json:"protectedArtifactLocationID,omitempty"`

	// Optional. The name of a Vertex AI Tensorboard resource to which this CustomJob will upload Tensorboard logs.
	// +kcc:proto:field=google.cloud.aiplatform.v1.CustomJobSpec.tensorboard
	// +optional
	TensorboardRef *VertexAITensorboardRef `json:"tensorboardRef,omitempty"`

	// Optional. Whether you want Vertex AI to enable interactive shell
	// access to training containers.
	// +kcc:proto:field=google.cloud.aiplatform.v1.CustomJobSpec.enable_web_access
	// +optional
	EnableWebAccess *bool `json:"enableWebAccess,omitempty"`

	// Optional. Whether you want Vertex AI to enable access to the customized
	// dashboard in training chief container.
	// +kcc:proto:field=google.cloud.aiplatform.v1.CustomJobSpec.enable_dashboard_access
	// +optional
	EnableDashboardAccess *bool `json:"enableDashboardAccess,omitempty"`
}

// +kcc:proto=google.cloud.aiplatform.v1.WorkerPoolSpec
type CustomJobWorkerPoolSpec struct {
	// The custom container task.
	// +kcc:proto:field=google.cloud.aiplatform.v1.WorkerPoolSpec.container_spec
	CustomJobContainerSpec *CustomJobContainerSpec `json:"containerSpec,omitempty"`

	// The Python packaged task.
	// +kcc:proto:field=google.cloud.aiplatform.v1.WorkerPoolSpec.python_package_spec
	CustomJobPythonPackageSpec *CustomJobPythonPackageSpec `json:"pythonPackageSpec,omitempty"`

	// Optional. Immutable. The specification of a single machine.
	// +kcc:proto:field=google.cloud.aiplatform.v1.WorkerPoolSpec.machine_spec
	// +optional
	CustomJobMachineSpec *CustomJobMachineSpec `json:"machineSpec,omitempty"`

	// Optional. The number of worker replicas to use for this worker pool.
	// +kcc:proto:field=google.cloud.aiplatform.v1.WorkerPoolSpec.replica_count
	// +optional
	ReplicaCount *int64 `json:"replicaCount,omitempty"`

	// Optional. List of NFS mount spec.
	// +kcc:proto:field=google.cloud.aiplatform.v1.WorkerPoolSpec.nfs_mounts
	// +optional
	CustomJobNfsMounts []CustomJobNfsMount `json:"nfsMounts,omitempty"`

	// Disk spec.
	// +kcc:proto:field=google.cloud.aiplatform.v1.WorkerPoolSpec.disk_spec
	CustomJobDiskSpec *CustomJobDiskSpec `json:"diskSpec,omitempty"`
}

// +kcc:proto=google.cloud.aiplatform.v1.ContainerSpec
type CustomJobContainerSpec struct {
	// Required. The URI of a container image in the Container Registry that is to
	// be run on each worker replica.
	// +kcc:proto:field=google.cloud.aiplatform.v1.ContainerSpec.image_uri
	// +kubebuilder:validation:Required
	ImageURI *string `json:"imageURI"`

	// The command to be invoked when the container is started.
	// It overrides the entrypoint instruction in Dockerfile when provided.
	// +kcc:proto:field=google.cloud.aiplatform.v1.ContainerSpec.command
	Command []string `json:"command,omitempty"`

	// The arguments to be passed when starting the container.
	// +kcc:proto:field=google.cloud.aiplatform.v1.ContainerSpec.args
	Args []string `json:"args,omitempty"`

	// Environment variables to be passed to the container.
	// Maximum limit is 100.
	// +kcc:proto:field=google.cloud.aiplatform.v1.ContainerSpec.env
	Env []CustomJobEnvVar `json:"env,omitempty"`
}

// +kcc:proto=google.cloud.aiplatform.v1.PythonPackageSpec
type CustomJobPythonPackageSpec struct {
	// Required. The URI of a container image in Artifact Registry that will run
	// the provided Python package.
	// +kcc:proto:field=google.cloud.aiplatform.v1.PythonPackageSpec.executor_image_uri
	// +kubebuilder:validation:Required
	ExecutorImageURI *string `json:"executorImageURI"`

	// Required. The Google Cloud Storage location of the Python package files
	// which are the training program and its dependent packages.
	// +kcc:proto:field=google.cloud.aiplatform.v1.PythonPackageSpec.package_uris
	// +kubebuilder:validation:Required
	PackageURIs []string `json:"packageURIs"`

	// Required. The Python module name to run after installing the packages.
	// +kcc:proto:field=google.cloud.aiplatform.v1.PythonPackageSpec.python_module
	// +kubebuilder:validation:Required
	PythonModule *string `json:"pythonModule"`

	// Command line arguments to be passed to the Python task.
	// +kcc:proto:field=google.cloud.aiplatform.v1.PythonPackageSpec.args
	Args []string `json:"args,omitempty"`

	// Environment variables to be passed to the python module.
	// Maximum limit is 100.
	// +kcc:proto:field=google.cloud.aiplatform.v1.PythonPackageSpec.env
	Env []CustomJobEnvVar `json:"env,omitempty"`
}

// +kcc:proto=google.cloud.aiplatform.v1.MachineSpec
type CustomJobMachineSpec struct {
	// Immutable. The type of the machine.
	// +kcc:proto:field=google.cloud.aiplatform.v1.MachineSpec.machine_type
	MachineType *string `json:"machineType,omitempty"`

	// Immutable. The type of accelerator(s) that may be attached to the machine as per
	// accelerator_count.
	// +kcc:proto:field=google.cloud.aiplatform.v1.MachineSpec.accelerator_type
	// +kubebuilder:validation:Enum=ACCELERATOR_TYPE_UNSPECIFIED;NVIDIA_TESLA_K80;NVIDIA_TESLA_P100;NVIDIA_TESLA_V100;NVIDIA_TESLA_P4;NVIDIA_TESLA_T4;NVIDIA_TESLA_A100;NVIDIA_A100_80GB;NVIDIA_L4;NVIDIA_H100_80GB;NVIDIA_H100_MEGA_80GB;TPU_V2;TPU_V3;TPU_V4_POD;TPU_V5_LITEPOD
	AcceleratorType *string `json:"acceleratorType,omitempty"`

	// The number of accelerators to attach to the machine.
	// +kcc:proto:field=google.cloud.aiplatform.v1.MachineSpec.accelerator_count
	AcceleratorCount *int32 `json:"acceleratorCount,omitempty"`
}

// +kcc:proto=google.cloud.aiplatform.v1.NfsMount
type CustomJobNfsMount struct {
	// Required. IP address of the NFS server.
	// +kcc:proto:field=google.cloud.aiplatform.v1.NfsMount.server
	// +kubebuilder:validation:Required
	Server *string `json:"server"`

	// Required. Source path exported from NFS server.
	// +kcc:proto:field=google.cloud.aiplatform.v1.NfsMount.path
	// +kubebuilder:validation:Required
	Path *string `json:"path"`

	// Required. Destination mount path.
	// +kcc:proto:field=google.cloud.aiplatform.v1.NfsMount.mount_point
	// +kubebuilder:validation:Required
	MountPoint *string `json:"mountPoint"`
}

// +kcc:proto=google.cloud.aiplatform.v1.DiskSpec
type CustomJobDiskSpec struct {
	// Type of the boot disk (default is "pd-standard").
	// +kcc:proto:field=google.cloud.aiplatform.v1.DiskSpec.boot_disk_type
	BootDiskType *string `json:"bootDiskType,omitempty"`

	// Size in GB of the boot disk (default is 100GB).
	// +kcc:proto:field=google.cloud.aiplatform.v1.DiskSpec.boot_disk_size_gb
	BootDiskSizeGB *int32 `json:"bootDiskSizeGB,omitempty"`
}

// +kcc:proto=google.cloud.aiplatform.v1.EnvVar
type CustomJobEnvVar struct {
	// Required. Name of the environment variable.
	// +kcc:proto:field=google.cloud.aiplatform.v1.EnvVar.name
	// +kubebuilder:validation:Required
	Name *string `json:"name"`

	// Required. Variables that reference a $(VAR_NAME) are expanded
	// using the previous defined environment variables in the container and
	// any service environment variables.
	// +kcc:proto:field=google.cloud.aiplatform.v1.EnvVar.value
	// +kubebuilder:validation:Required
	Value *string `json:"value"`
}

// +kcc:proto=google.cloud.aiplatform.v1.Scheduling
type CustomJobScheduling struct {
	// The maximum job running time.
	// +kcc:proto:field=google.cloud.aiplatform.v1.Scheduling.timeout
	Timeout *string `json:"timeout,omitempty"`

	// Restarts the entire CustomJob if a worker gets restarted.
	// +kcc:proto:field=google.cloud.aiplatform.v1.Scheduling.restart_job_on_worker_restart
	RestartJobOnWorkerRestart *bool `json:"restartJobOnWorkerRestart,omitempty"`

	// Optional. This determines which type of scheduling strategy to use.
	// +kcc:proto:field=google.cloud.aiplatform.v1.Scheduling.strategy
	// +optional
	Strategy *string `json:"strategy,omitempty"`
}

// +kcc:proto=google.cloud.aiplatform.v1.GcsDestination
type CustomJobGcsDestination struct {
	// Required. URI of the Cloud Storage directory.
	// +kcc:proto:field=google.cloud.aiplatform.v1.GcsDestination.output_uri_prefix
	// +kubebuilder:validation:Required
	OutputURIPrefix *string `json:"outputURIPrefix"`
}

// VertexAICustomJobStatus defines the config connector machine state of VertexAICustomJob
type VertexAICustomJobStatus struct {
	/* Conditions represent the latest available observations of the
	   object's current state. */
	Conditions []v1alpha1.Condition `json:"conditions,omitempty"`

	// ObservedGeneration is the generation of the resource that was most recently observed by the Config Connector controller. If this is equal to metadata.generation, then that means that the current reported status reflects the most recent desired state of the resource.
	ObservedGeneration *int64 `json:"observedGeneration,omitempty"`

	// A unique specifier for the VertexAICustomJob resource in GCP.
	ExternalRef *string `json:"externalRef,omitempty"`

	// ObservedState is the state of the resource as most recently observed in GCP.
	ObservedState *VertexAICustomJobObservedState `json:"observedState,omitempty"`
}

// VertexAICustomJobObservedState is the state of the VertexAICustomJob resource as most recently observed in GCP.
// +kcc:observedstate:proto=google.cloud.aiplatform.v1.CustomJob
type VertexAICustomJobObservedState struct {
	// Output only. The detailed state of the job.
	// +kcc:proto:field=google.cloud.aiplatform.v1.CustomJob.state
	State *string `json:"state,omitempty"`
}

// +genclient
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:resource:categories=gcp,shortName=gcpvertexaicustomjob;gcpvertexaicustomjobs
// +kubebuilder:subresource:status
// +kubebuilder:metadata:labels="cnrm.cloud.google.com/managed-by-kcc=true"
// +kubebuilder:metadata:labels="cnrm.cloud.google.com/system=true"
// +kubebuilder:metadata:labels="cnrm.cloud.google.com/stability-level=alpha"
// +kubebuilder:printcolumn:name="Age",JSONPath=".metadata.creationTimestamp",type="date"
// +kubebuilder:printcolumn:name="Ready",JSONPath=".status.conditions[?(@.type=='Ready')].status",type="string",description="When 'True', the most recent reconcile of the resource succeeded"
// +kubebuilder:printcolumn:name="Status",JSONPath=".status.conditions[?(@.type=='Ready')].reason",type="string",description="The reason for the value in 'Ready'"
// +kubebuilder:printcolumn:name="Status Age",JSONPath=".status.conditions[?(@.type=='Ready')].lastTransitionTime",type="date",description="The last transition time for the value in 'Status'"

// VertexAICustomJob is the Schema for the VertexAICustomJob API
// +k8s:openapi-gen=true
type VertexAICustomJob struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	// +required
	Spec   VertexAICustomJobSpec   `json:"spec,omitempty"`
	Status VertexAICustomJobStatus `json:"status,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// VertexAICustomJobList contains a list of VertexAICustomJob
type VertexAICustomJobList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []VertexAICustomJob `json:"items"`
}

func init() {
	SchemeBuilder.Register(&VertexAICustomJob{}, &VertexAICustomJobList{})
}
