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
// krm.group: colab.cnrm.cloud.google.com
// krm.version: v1alpha1
// proto.service: google.cloud.aiplatform.v1beta1
// resource: ColabRuntimeTemplate:NotebookRuntimeTemplate
// resource: ColabRuntime:NotebookRuntime

package v1alpha1

// +kcc:proto=google.cloud.aiplatform.v1beta1.MachineSpec
type MachineSpec struct {
	// Immutable. The type of the machine.
	//
	//  See the [list of machine types supported for
	//  prediction](https://cloud.google.com/vertex-ai/docs/predictions/configure-compute#machine-types)
	//
	//  See the [list of machine types supported for custom
	//  training](https://cloud.google.com/vertex-ai/docs/training/configure-compute#machine-types).
	//
	//  For [DeployedModel][google.cloud.aiplatform.v1beta1.DeployedModel] this
	//  field is optional, and the default value is `n1-standard-2`. For
	//  [BatchPredictionJob][google.cloud.aiplatform.v1beta1.BatchPredictionJob] or
	//  as part of [WorkerPoolSpec][google.cloud.aiplatform.v1beta1.WorkerPoolSpec]
	//  this field is required.
	// +kcc:proto:field=google.cloud.aiplatform.v1beta1.MachineSpec.machine_type
	MachineType *string `json:"machineType,omitempty"`

	// Immutable. The type of accelerator(s) that may be attached to the machine
	//  as per
	//  [accelerator_count][google.cloud.aiplatform.v1beta1.MachineSpec.accelerator_count].
	// +kcc:proto:field=google.cloud.aiplatform.v1beta1.MachineSpec.accelerator_type
	AcceleratorType *string `json:"acceleratorType,omitempty"`

	// The number of accelerators to attach to the machine.
	// +kcc:proto:field=google.cloud.aiplatform.v1beta1.MachineSpec.accelerator_count
	AcceleratorCount *int32 `json:"acceleratorCount,omitempty"`

	// Immutable. The topology of the TPUs. Corresponds to the TPU topologies
	//  available from GKE. (Example: tpu_topology: "2x2x1").
	// +kcc:proto:field=google.cloud.aiplatform.v1beta1.MachineSpec.tpu_topology
	TpuTopology *string `json:"tpuTopology,omitempty"`

	// Optional. Immutable. Configuration controlling how this resource pool
	//  consumes reservation.
	// +kcc:proto:field=google.cloud.aiplatform.v1beta1.MachineSpec.reservation_affinity
	ReservationAffinity *ReservationAffinity `json:"reservationAffinity,omitempty"`
}

// +kcc:proto=google.cloud.aiplatform.v1beta1.NotebookEucConfig
type NotebookEUCConfig struct {
	// Input only. Whether EUC is disabled in this NotebookRuntimeTemplate.
	//  In proto3, the default value of a boolean is false. In this way, by default
	//  EUC will be enabled for NotebookRuntimeTemplate.
	// +kcc:proto:field=google.cloud.aiplatform.v1beta1.NotebookEucConfig.euc_disabled
	EUCDisabled *bool `json:"eucDisabled,omitempty"`
}

// +kcc:proto=google.cloud.aiplatform.v1beta1.NotebookIdleShutdownConfig
type NotebookIdleShutdownConfig struct {
	// Required. Duration is accurate to the second. In Notebook, Idle Timeout is
	//  accurate to minute so the range of idle_timeout (second) is: 10 * 60 ~ 1440
	//  * 60.
	// +kcc:proto:field=google.cloud.aiplatform.v1beta1.NotebookIdleShutdownConfig.idle_timeout
	IdleTimeout *string `json:"idleTimeout,omitempty"`

	// Whether Idle Shutdown is disabled in this NotebookRuntimeTemplate.
	// +kcc:proto:field=google.cloud.aiplatform.v1beta1.NotebookIdleShutdownConfig.idle_shutdown_disabled
	IdleShutdownDisabled *bool `json:"idleShutdownDisabled,omitempty"`
}

// +kcc:proto=google.cloud.aiplatform.v1beta1.PersistentDiskSpec
type PersistentDiskSpec struct {
	// Type of the disk (default is "pd-standard").
	//  Valid values: "pd-ssd" (Persistent Disk Solid State Drive)
	//  "pd-standard" (Persistent Disk Hard Disk Drive)
	//  "pd-balanced" (Balanced Persistent Disk)
	//  "pd-extreme" (Extreme Persistent Disk)
	// +kcc:proto:field=google.cloud.aiplatform.v1beta1.PersistentDiskSpec.disk_type
	DiskType *string `json:"diskType,omitempty"`

	// Size in GB of the disk (default is 100GB).
	// +kcc:proto:field=google.cloud.aiplatform.v1beta1.PersistentDiskSpec.disk_size_gb
	DiskSizeGB *int64 `json:"diskSizeGB,omitempty"`
}

// +kcc:proto=google.cloud.aiplatform.v1beta1.ReservationAffinity
type ReservationAffinity struct {
	// Required. Specifies the reservation affinity type.
	// +kcc:proto:field=google.cloud.aiplatform.v1beta1.ReservationAffinity.reservation_affinity_type
	ReservationAffinityType *string `json:"reservationAffinityType,omitempty"`

	// Optional. Corresponds to the label key of a reservation resource. To target
	//  a SPECIFIC_RESERVATION by name, use
	//  `compute.googleapis.com/reservation-name` as the key and specify the name
	//  of your reservation as its value.
	// +kcc:proto:field=google.cloud.aiplatform.v1beta1.ReservationAffinity.key
	Key *string `json:"key,omitempty"`

	// Optional. Corresponds to the label values of a reservation resource. This
	//  must be the full resource name of the reservation.
	// +kcc:proto:field=google.cloud.aiplatform.v1beta1.ReservationAffinity.values
	Values []string `json:"values,omitempty"`
}

// +kcc:proto=google.cloud.aiplatform.v1beta1.ShieldedVmConfig
type ShieldedVMConfig struct {
	// Defines whether the instance has [Secure
	//  Boot](https://cloud.google.com/compute/shielded-vm/docs/shielded-vm#secure-boot)
	//  enabled.
	//
	//  Secure Boot helps ensure that the system only runs authentic software by
	//  verifying the digital signature of all boot components, and halting the
	//  boot process if signature verification fails.
	// +kcc:proto:field=google.cloud.aiplatform.v1beta1.ShieldedVmConfig.enable_secure_boot
	EnableSecureBoot *bool `json:"enableSecureBoot,omitempty"`
}

// +kcc:proto=google.cloud.aiplatform.v1beta1.NotebookEucConfig
type NotebookEUCConfigObservedState struct {
	// Output only. Whether ActAs check is bypassed for service account attached
	//  to the VM. If false, we need ActAs check for the default Compute Engine
	//  Service account. When a Runtime is created, a VM is allocated using Default
	//  Compute Engine Service Account. Any user requesting to use this Runtime
	//  requires Service Account User (ActAs) permission over this SA. If true,
	//  Runtime owner is using EUC and does not require the above permission as VM
	//  no longer use default Compute Engine SA, but a P4SA.
	// +kcc:proto:field=google.cloud.aiplatform.v1beta1.NotebookEucConfig.bypass_actas_check
	BypassActasCheck *bool `json:"bypassActasCheck,omitempty"`
}
