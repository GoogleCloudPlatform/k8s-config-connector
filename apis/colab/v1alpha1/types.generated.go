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

// +kcc:proto=google.cloud.aiplatform.v1.NotebookEucConfig
type NotebookEucConfig struct {
	// Input only. Whether EUC is disabled in this NotebookRuntimeTemplate.
	//  In proto3, the default value of a boolean is false. In this way, by default
	//  EUC will be enabled for NotebookRuntimeTemplate.
	// +kcc:proto:field=google.cloud.aiplatform.v1.NotebookEucConfig.euc_disabled
	EucDisabled *bool `json:"eucDisabled,omitempty"`
}

// +kcc:proto=google.cloud.aiplatform.v1.NotebookIdleShutdownConfig
type NotebookIdleShutdownConfig struct {
	// Required. Duration is accurate to the second. In Notebook, Idle Timeout is
	//  accurate to minute so the range of idle_timeout (second) is: 10 * 60 ~ 1440
	//  * 60.
	// +kcc:proto:field=google.cloud.aiplatform.v1.NotebookIdleShutdownConfig.idle_timeout
	IdleTimeout *string `json:"idleTimeout,omitempty"`

	// Whether Idle Shutdown is disabled in this NotebookRuntimeTemplate.
	// +kcc:proto:field=google.cloud.aiplatform.v1.NotebookIdleShutdownConfig.idle_shutdown_disabled
	IdleShutdownDisabled *bool `json:"idleShutdownDisabled,omitempty"`
}

// +kcc:proto=google.cloud.aiplatform.v1.NotebookRuntimeTemplate
type NotebookRuntimeTemplate struct {
	// The resource name of the NotebookRuntimeTemplate.
	// +kcc:proto:field=google.cloud.aiplatform.v1.NotebookRuntimeTemplate.name
	Name *string `json:"name,omitempty"`

	// Required. The display name of the NotebookRuntimeTemplate.
	//  The name can be up to 128 characters long and can consist of any UTF-8
	//  characters.
	// +kcc:proto:field=google.cloud.aiplatform.v1.NotebookRuntimeTemplate.display_name
	DisplayName *string `json:"displayName,omitempty"`

	// The description of the NotebookRuntimeTemplate.
	// +kcc:proto:field=google.cloud.aiplatform.v1.NotebookRuntimeTemplate.description
	Description *string `json:"description,omitempty"`

	// Optional. Immutable. The specification of a single machine for the
	//  template.
	// +kcc:proto:field=google.cloud.aiplatform.v1.NotebookRuntimeTemplate.machine_spec
	MachineSpec *MachineSpec `json:"machineSpec,omitempty"`

	// Optional. The specification of [persistent
	//  disk][https://cloud.google.com/compute/docs/disks/persistent-disks]
	//  attached to the runtime as data disk storage.
	// +kcc:proto:field=google.cloud.aiplatform.v1.NotebookRuntimeTemplate.data_persistent_disk_spec
	DataPersistentDiskSpec *PersistentDiskSpec `json:"dataPersistentDiskSpec,omitempty"`

	// Optional. Network spec.
	// +kcc:proto:field=google.cloud.aiplatform.v1.NotebookRuntimeTemplate.network_spec
	NetworkSpec *NetworkSpec `json:"networkSpec,omitempty"`

	// The service account that the runtime workload runs as.
	//  You can use any service account within the same project, but you
	//  must have the service account user permission to use the instance.
	//
	//  If not specified, the [Compute Engine default service
	//  account](https://cloud.google.com/compute/docs/access/service-accounts#default_service_account)
	//  is used.
	// +kcc:proto:field=google.cloud.aiplatform.v1.NotebookRuntimeTemplate.service_account
	ServiceAccount *string `json:"serviceAccount,omitempty"`

	// Used to perform consistent read-modify-write updates. If not set, a blind
	//  "overwrite" update happens.
	// +kcc:proto:field=google.cloud.aiplatform.v1.NotebookRuntimeTemplate.etag
	Etag *string `json:"etag,omitempty"`

	// The labels with user-defined metadata to organize the
	//  NotebookRuntimeTemplates.
	//
	//  Label keys and values can be no longer than 64 characters
	//  (Unicode codepoints), can only contain lowercase letters, numeric
	//  characters, underscores and dashes. International characters are allowed.
	//
	//  See https://goo.gl/xmQnxf for more information and examples of labels.
	// +kcc:proto:field=google.cloud.aiplatform.v1.NotebookRuntimeTemplate.labels
	Labels map[string]string `json:"labels,omitempty"`

	// The idle shutdown configuration of NotebookRuntimeTemplate. This config
	//  will only be set when idle shutdown is enabled.
	// +kcc:proto:field=google.cloud.aiplatform.v1.NotebookRuntimeTemplate.idle_shutdown_config
	IdleShutdownConfig *NotebookIdleShutdownConfig `json:"idleShutdownConfig,omitempty"`

	// EUC configuration of the NotebookRuntimeTemplate.
	// +kcc:proto:field=google.cloud.aiplatform.v1.NotebookRuntimeTemplate.euc_config
	EucConfig *NotebookEucConfig `json:"eucConfig,omitempty"`

	// Optional. Immutable. The type of the notebook runtime template.
	// +kcc:proto:field=google.cloud.aiplatform.v1.NotebookRuntimeTemplate.notebook_runtime_type
	NotebookRuntimeType *string `json:"notebookRuntimeType,omitempty"`

	// Optional. Immutable. Runtime Shielded VM spec.
	// +kcc:proto:field=google.cloud.aiplatform.v1.NotebookRuntimeTemplate.shielded_vm_config
	ShieldedVmConfig *ShieldedVmConfig `json:"shieldedVmConfig,omitempty"`

	// Optional. The Compute Engine tags to add to runtime (see [Tagging
	//  instances](https://cloud.google.com/vpc/docs/add-remove-network-tags)).
	// +kcc:proto:field=google.cloud.aiplatform.v1.NotebookRuntimeTemplate.network_tags
	NetworkTags []string `json:"networkTags,omitempty"`

	// Customer-managed encryption key spec for the notebook runtime.
	// +kcc:proto:field=google.cloud.aiplatform.v1.NotebookRuntimeTemplate.encryption_spec
	EncryptionSpec *EncryptionSpec `json:"encryptionSpec,omitempty"`
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

// +kcc:proto=google.cloud.aiplatform.v1.ShieldedVmConfig
type ShieldedVmConfig struct {
	// Defines whether the instance has [Secure
	//  Boot](https://cloud.google.com/compute/shielded-vm/docs/shielded-vm#secure-boot)
	//  enabled.
	//
	//  Secure Boot helps ensure that the system only runs authentic software by
	//  verifying the digital signature of all boot components, and halting the
	//  boot process if signature verification fails.
	// +kcc:proto:field=google.cloud.aiplatform.v1.ShieldedVmConfig.enable_secure_boot
	EnableSecureBoot *bool `json:"enableSecureBoot,omitempty"`
}

// +kcc:proto=google.cloud.aiplatform.v1.NotebookEucConfig
type NotebookEucConfigObservedState struct {
	// Output only. Whether ActAs check is bypassed for service account attached
	//  to the VM. If false, we need ActAs check for the default Compute Engine
	//  Service account. When a Runtime is created, a VM is allocated using Default
	//  Compute Engine Service Account. Any user requesting to use this Runtime
	//  requires Service Account User (ActAs) permission over this SA. If true,
	//  Runtime owner is using EUC and does not require the above permission as VM
	//  no longer use default Compute Engine SA, but a P4SA.
	// +kcc:proto:field=google.cloud.aiplatform.v1.NotebookEucConfig.bypass_actas_check
	BypassActasCheck *bool `json:"bypassActasCheck,omitempty"`
}

// +kcc:proto=google.cloud.aiplatform.v1.NotebookRuntimeTemplate
type NotebookRuntimeTemplateObservedState struct {
	// Output only. The default template to use if not specified.
	// +kcc:proto:field=google.cloud.aiplatform.v1.NotebookRuntimeTemplate.is_default
	IsDefault *bool `json:"isDefault,omitempty"`

	// EUC configuration of the NotebookRuntimeTemplate.
	// +kcc:proto:field=google.cloud.aiplatform.v1.NotebookRuntimeTemplate.euc_config
	EucConfig *NotebookEucConfigObservedState `json:"eucConfig,omitempty"`

	// Output only. Timestamp when this NotebookRuntimeTemplate was created.
	// +kcc:proto:field=google.cloud.aiplatform.v1.NotebookRuntimeTemplate.create_time
	CreateTime *string `json:"createTime,omitempty"`

	// Output only. Timestamp when this NotebookRuntimeTemplate was most recently
	//  updated.
	// +kcc:proto:field=google.cloud.aiplatform.v1.NotebookRuntimeTemplate.update_time
	UpdateTime *string `json:"updateTime,omitempty"`
}
