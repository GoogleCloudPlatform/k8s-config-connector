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


// +kcc:proto=google.cloud.aiplatform.v1.EncryptionSpec
type EncryptionSpec struct {
	// Required. The Cloud KMS resource identifier of the customer managed
	//  encryption key used to protect a resource. Has the form:
	//  `projects/my-project/locations/my-region/keyRings/my-kr/cryptoKeys/my-key`.
	//  The key needs to be in the same region as where the compute resource is
	//  created.
	// +kcc:proto:field=google.cloud.aiplatform.v1.EncryptionSpec.kms_key_name
	KMSKeyName *string `json:"kmsKeyName,omitempty"`
}

// +kcc:proto=google.cloud.aiplatform.v1.MachineSpec
type MachineSpec struct {
	// Immutable. The type of the machine.
	//
	//  See the [list of machine types supported for
	//  prediction](https://cloud.google.com/vertex-ai/docs/predictions/configure-compute#machine-types)
	//
	//  See the [list of machine types supported for custom
	//  training](https://cloud.google.com/vertex-ai/docs/training/configure-compute#machine-types).
	//
	//  For [DeployedModel][google.cloud.aiplatform.v1.DeployedModel] this field is
	//  optional, and the default value is `n1-standard-2`. For
	//  [BatchPredictionJob][google.cloud.aiplatform.v1.BatchPredictionJob] or as
	//  part of [WorkerPoolSpec][google.cloud.aiplatform.v1.WorkerPoolSpec] this
	//  field is required.
	// +kcc:proto:field=google.cloud.aiplatform.v1.MachineSpec.machine_type
	MachineType *string `json:"machineType,omitempty"`

	// Immutable. The type of accelerator(s) that may be attached to the machine
	//  as per
	//  [accelerator_count][google.cloud.aiplatform.v1.MachineSpec.accelerator_count].
	// +kcc:proto:field=google.cloud.aiplatform.v1.MachineSpec.accelerator_type
	AcceleratorType *string `json:"acceleratorType,omitempty"`

	// The number of accelerators to attach to the machine.
	// +kcc:proto:field=google.cloud.aiplatform.v1.MachineSpec.accelerator_count
	AcceleratorCount *int32 `json:"acceleratorCount,omitempty"`

	// Immutable. The topology of the TPUs. Corresponds to the TPU topologies
	//  available from GKE. (Example: tpu_topology: "2x2x1").
	// +kcc:proto:field=google.cloud.aiplatform.v1.MachineSpec.tpu_topology
	TpuTopology *string `json:"tpuTopology,omitempty"`

	// Optional. Immutable. Configuration controlling how this resource pool
	//  consumes reservation.
	// +kcc:proto:field=google.cloud.aiplatform.v1.MachineSpec.reservation_affinity
	ReservationAffinity *ReservationAffinity `json:"reservationAffinity,omitempty"`
}

// +kcc:proto=google.cloud.aiplatform.v1.NetworkSpec
type NetworkSpec struct {
	// Whether to enable public internet access. Default false.
	// +kcc:proto:field=google.cloud.aiplatform.v1.NetworkSpec.enable_internet_access
	EnableInternetAccess *bool `json:"enableInternetAccess,omitempty"`

	// The full name of the Google Compute Engine
	//  [network](https://cloud.google.com//compute/docs/networks-and-firewalls#networks)
	// +kcc:proto:field=google.cloud.aiplatform.v1.NetworkSpec.network
	Network *string `json:"network,omitempty"`

	// The name of the subnet that this instance is in.
	//  Format:
	//  `projects/{project_id_or_number}/regions/{region}/subnetworks/{subnetwork_id}`
	// +kcc:proto:field=google.cloud.aiplatform.v1.NetworkSpec.subnetwork
	Subnetwork *string `json:"subnetwork,omitempty"`
}

// +kcc:proto=google.cloud.aiplatform.v1.NotebookExecutionJob
type NotebookExecutionJob struct {
	// The Dataform Repository pointing to a single file notebook repository.
	// +kcc:proto:field=google.cloud.aiplatform.v1.NotebookExecutionJob.dataform_repository_source
	DataformRepositorySource *NotebookExecutionJob_DataformRepositorySource `json:"dataformRepositorySource,omitempty"`

	// The Cloud Storage url pointing to the ipynb file. Format:
	//  `gs://bucket/notebook_file.ipynb`
	// +kcc:proto:field=google.cloud.aiplatform.v1.NotebookExecutionJob.gcs_notebook_source
	GcsNotebookSource *NotebookExecutionJob_GcsNotebookSource `json:"gcsNotebookSource,omitempty"`

	// The contents of an input notebook file.
	// +kcc:proto:field=google.cloud.aiplatform.v1.NotebookExecutionJob.direct_notebook_source
	DirectNotebookSource *NotebookExecutionJob_DirectNotebookSource `json:"directNotebookSource,omitempty"`

	// The NotebookRuntimeTemplate to source compute configuration from.
	// +kcc:proto:field=google.cloud.aiplatform.v1.NotebookExecutionJob.notebook_runtime_template_resource_name
	NotebookRuntimeTemplateResourceName *string `json:"notebookRuntimeTemplateResourceName,omitempty"`

	// The custom compute configuration for an execution job.
	// +kcc:proto:field=google.cloud.aiplatform.v1.NotebookExecutionJob.custom_environment_spec
	CustomEnvironmentSpec *NotebookExecutionJob_CustomEnvironmentSpec `json:"customEnvironmentSpec,omitempty"`

	// The Cloud Storage location to upload the result to. Format:
	//  `gs://bucket-name`
	// +kcc:proto:field=google.cloud.aiplatform.v1.NotebookExecutionJob.gcs_output_uri
	GcsOutputURI *string `json:"gcsOutputURI,omitempty"`

	// The user email to run the execution as. Only supported by Colab runtimes.
	// +kcc:proto:field=google.cloud.aiplatform.v1.NotebookExecutionJob.execution_user
	ExecutionUser *string `json:"executionUser,omitempty"`

	// The service account to run the execution as.
	// +kcc:proto:field=google.cloud.aiplatform.v1.NotebookExecutionJob.service_account
	ServiceAccount *string `json:"serviceAccount,omitempty"`

	// The Workbench runtime configuration to use for the notebook execution.
	// +kcc:proto:field=google.cloud.aiplatform.v1.NotebookExecutionJob.workbench_runtime
	WorkbenchRuntime *NotebookExecutionJob_WorkbenchRuntime `json:"workbenchRuntime,omitempty"`

	// The display name of the NotebookExecutionJob. The name can be up to 128
	//  characters long and can consist of any UTF-8 characters.
	// +kcc:proto:field=google.cloud.aiplatform.v1.NotebookExecutionJob.display_name
	DisplayName *string `json:"displayName,omitempty"`

	// Max running time of the execution job in seconds (default 86400s / 24 hrs).
	// +kcc:proto:field=google.cloud.aiplatform.v1.NotebookExecutionJob.execution_timeout
	ExecutionTimeout *string `json:"executionTimeout,omitempty"`

	// The labels with user-defined metadata to organize NotebookExecutionJobs.
	//
	//  Label keys and values can be no longer than 64 characters
	//  (Unicode codepoints), can only contain lowercase letters, numeric
	//  characters, underscores and dashes. International characters are allowed.
	//
	//  See https://goo.gl/xmQnxf for more information and examples of labels.
	//  System reserved label keys are prefixed with "aiplatform.googleapis.com/"
	//  and are immutable.
	// +kcc:proto:field=google.cloud.aiplatform.v1.NotebookExecutionJob.labels
	Labels map[string]string `json:"labels,omitempty"`

	// The name of the kernel to use during notebook execution. If unset, the
	//  default kernel is used.
	// +kcc:proto:field=google.cloud.aiplatform.v1.NotebookExecutionJob.kernel_name
	KernelName *string `json:"kernelName,omitempty"`

	// Customer-managed encryption key spec for the notebook execution job.
	//  This field is auto-populated if the
	//  [NotebookRuntimeTemplate][google.cloud.aiplatform.v1.NotebookRuntimeTemplate]
	//  has an encryption spec.
	// +kcc:proto:field=google.cloud.aiplatform.v1.NotebookExecutionJob.encryption_spec
	EncryptionSpec *EncryptionSpec `json:"encryptionSpec,omitempty"`
}

// +kcc:proto=google.cloud.aiplatform.v1.NotebookExecutionJob.CustomEnvironmentSpec
type NotebookExecutionJob_CustomEnvironmentSpec struct {
	// The specification of a single machine for the execution job.
	// +kcc:proto:field=google.cloud.aiplatform.v1.NotebookExecutionJob.CustomEnvironmentSpec.machine_spec
	MachineSpec *MachineSpec `json:"machineSpec,omitempty"`

	// The specification of a persistent disk to attach for the execution job.
	// +kcc:proto:field=google.cloud.aiplatform.v1.NotebookExecutionJob.CustomEnvironmentSpec.persistent_disk_spec
	PersistentDiskSpec *PersistentDiskSpec `json:"persistentDiskSpec,omitempty"`

	// The network configuration to use for the execution job.
	// +kcc:proto:field=google.cloud.aiplatform.v1.NotebookExecutionJob.CustomEnvironmentSpec.network_spec
	NetworkSpec *NetworkSpec `json:"networkSpec,omitempty"`
}

// +kcc:proto=google.cloud.aiplatform.v1.NotebookExecutionJob.DataformRepositorySource
type NotebookExecutionJob_DataformRepositorySource struct {
	// The resource name of the Dataform Repository. Format:
	//  `projects/{project_id}/locations/{location}/repositories/{repository_id}`
	// +kcc:proto:field=google.cloud.aiplatform.v1.NotebookExecutionJob.DataformRepositorySource.dataform_repository_resource_name
	DataformRepositoryResourceName *string `json:"dataformRepositoryResourceName,omitempty"`

	// The commit SHA to read repository with. If unset, the file will be read
	//  at HEAD.
	// +kcc:proto:field=google.cloud.aiplatform.v1.NotebookExecutionJob.DataformRepositorySource.commit_sha
	CommitSha *string `json:"commitSha,omitempty"`
}

// +kcc:proto=google.cloud.aiplatform.v1.NotebookExecutionJob.DirectNotebookSource
type NotebookExecutionJob_DirectNotebookSource struct {
	// The base64-encoded contents of the input notebook file.
	// +kcc:proto:field=google.cloud.aiplatform.v1.NotebookExecutionJob.DirectNotebookSource.content
	Content []byte `json:"content,omitempty"`
}

// +kcc:proto=google.cloud.aiplatform.v1.NotebookExecutionJob.GcsNotebookSource
type NotebookExecutionJob_GcsNotebookSource struct {
	// The Cloud Storage uri pointing to the ipynb file. Format:
	//  `gs://bucket/notebook_file.ipynb`
	// +kcc:proto:field=google.cloud.aiplatform.v1.NotebookExecutionJob.GcsNotebookSource.uri
	URI *string `json:"uri,omitempty"`

	// The version of the Cloud Storage object to read. If unset, the current
	//  version of the object is read. See
	//  https://cloud.google.com/storage/docs/metadata#generation-number.
	// +kcc:proto:field=google.cloud.aiplatform.v1.NotebookExecutionJob.GcsNotebookSource.generation
	Generation *string `json:"generation,omitempty"`
}

// +kcc:proto=google.cloud.aiplatform.v1.NotebookExecutionJob.WorkbenchRuntime
type NotebookExecutionJob_WorkbenchRuntime struct {
}

// +kcc:proto=google.cloud.aiplatform.v1.PersistentDiskSpec
type PersistentDiskSpec struct {
	// Type of the disk (default is "pd-standard").
	//  Valid values: "pd-ssd" (Persistent Disk Solid State Drive)
	//  "pd-standard" (Persistent Disk Hard Disk Drive)
	//  "pd-balanced" (Balanced Persistent Disk)
	//  "pd-extreme" (Extreme Persistent Disk)
	// +kcc:proto:field=google.cloud.aiplatform.v1.PersistentDiskSpec.disk_type
	DiskType *string `json:"diskType,omitempty"`

	// Size in GB of the disk (default is 100GB).
	// +kcc:proto:field=google.cloud.aiplatform.v1.PersistentDiskSpec.disk_size_gb
	DiskSizeGB *int64 `json:"diskSizeGB,omitempty"`
}

// +kcc:proto=google.cloud.aiplatform.v1.ReservationAffinity
type ReservationAffinity struct {
	// Required. Specifies the reservation affinity type.
	// +kcc:proto:field=google.cloud.aiplatform.v1.ReservationAffinity.reservation_affinity_type
	ReservationAffinityType *string `json:"reservationAffinityType,omitempty"`

	// Optional. Corresponds to the label key of a reservation resource. To target
	//  a SPECIFIC_RESERVATION by name, use
	//  `compute.googleapis.com/reservation-name` as the key and specify the name
	//  of your reservation as its value.
	// +kcc:proto:field=google.cloud.aiplatform.v1.ReservationAffinity.key
	Key *string `json:"key,omitempty"`

	// Optional. Corresponds to the label values of a reservation resource. This
	//  must be the full resource name of the reservation.
	// +kcc:proto:field=google.cloud.aiplatform.v1.ReservationAffinity.values
	Values []string `json:"values,omitempty"`
}

// +kcc:proto=google.protobuf.Any
type Any struct {
	// A URL/resource name that uniquely identifies the type of the serialized
	//  protocol buffer message. This string must contain at least
	//  one "/" character. The last segment of the URL's path must represent
	//  the fully qualified name of the type (as in
	//  `path/google.protobuf.Duration`). The name should be in a canonical form
	//  (e.g., leading "." is not accepted).
	//
	//  In practice, teams usually precompile into the binary all types that they
	//  expect it to use in the context of Any. However, for URLs which use the
	//  scheme `http`, `https`, or no scheme, one can optionally set up a type
	//  server that maps type URLs to message definitions as follows:
	//
	//  * If no scheme is provided, `https` is assumed.
	//  * An HTTP GET on the URL must yield a [google.protobuf.Type][]
	//    value in binary format, or produce an error.
	//  * Applications are allowed to cache lookup results based on the
	//    URL, or have them precompiled into a binary to avoid any
	//    lookup. Therefore, binary compatibility needs to be preserved
	//    on changes to types. (Use versioned type names to manage
	//    breaking changes.)
	//
	//  Note: this functionality is not currently available in the official
	//  protobuf release, and it is not used for type URLs beginning with
	//  type.googleapis.com.
	//
	//  Schemes other than `http`, `https` (or the empty scheme) might be
	//  used with implementation specific semantics.
	// +kcc:proto:field=google.protobuf.Any.type_url
	TypeURL *string `json:"typeURL,omitempty"`

	// Must be a valid serialized protocol buffer of the above specified type.
	// +kcc:proto:field=google.protobuf.Any.value
	Value []byte `json:"value,omitempty"`
}

// +kcc:proto=google.rpc.Status
type Status struct {
	// The status code, which should be an enum value of
	//  [google.rpc.Code][google.rpc.Code].
	// +kcc:proto:field=google.rpc.Status.code
	Code *int32 `json:"code,omitempty"`

	// A developer-facing error message, which should be in English. Any
	//  user-facing error message should be localized and sent in the
	//  [google.rpc.Status.details][google.rpc.Status.details] field, or localized
	//  by the client.
	// +kcc:proto:field=google.rpc.Status.message
	Message *string `json:"message,omitempty"`

	// A list of messages that carry the error details.  There is a common set of
	//  message types for APIs to use.
	// +kcc:proto:field=google.rpc.Status.details
	Details []Any `json:"details,omitempty"`
}

// +kcc:proto=google.cloud.aiplatform.v1.NotebookExecutionJob
type NotebookExecutionJobObservedState struct {
	// Output only. The resource name of this NotebookExecutionJob. Format:
	//  `projects/{project_id}/locations/{location}/notebookExecutionJobs/{job_id}`
	// +kcc:proto:field=google.cloud.aiplatform.v1.NotebookExecutionJob.name
	Name *string `json:"name,omitempty"`

	// Output only. The Schedule resource name if this job is triggered by one.
	//  Format:
	//  `projects/{project_id}/locations/{location}/schedules/{schedule_id}`
	// +kcc:proto:field=google.cloud.aiplatform.v1.NotebookExecutionJob.schedule_resource_name
	ScheduleResourceName *string `json:"scheduleResourceName,omitempty"`

	// Output only. The state of the NotebookExecutionJob.
	// +kcc:proto:field=google.cloud.aiplatform.v1.NotebookExecutionJob.job_state
	JobState *string `json:"jobState,omitempty"`

	// Output only. Populated when the NotebookExecutionJob is completed. When
	//  there is an error during notebook execution, the error details are
	//  populated.
	// +kcc:proto:field=google.cloud.aiplatform.v1.NotebookExecutionJob.status
	Status *Status `json:"status,omitempty"`

	// Output only. Timestamp when this NotebookExecutionJob was created.
	// +kcc:proto:field=google.cloud.aiplatform.v1.NotebookExecutionJob.create_time
	CreateTime *string `json:"createTime,omitempty"`

	// Output only. Timestamp when this NotebookExecutionJob was most recently
	//  updated.
	// +kcc:proto:field=google.cloud.aiplatform.v1.NotebookExecutionJob.update_time
	UpdateTime *string `json:"updateTime,omitempty"`
}
