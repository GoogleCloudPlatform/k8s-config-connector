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

	// Optional. Immutable. The Nvidia GPU partition size.
	//
	//  When specified, the requested accelerators will be partitioned into
	//  smaller GPU partitions. For example, if the request is for 8 units of
	//  NVIDIA A100 GPUs, and gpu_partition_size="1g.10gb", the service will
	//  create 8 * 7 = 56 partitioned MIG instances.
	//
	//  The partition size must be a value supported by the requested accelerator.
	//  Refer to
	//  [Nvidia GPU
	//  Partitioning](https://cloud.google.com/kubernetes-engine/docs/how-to/gpus-multi#multi-instance_gpu_partitions)
	//  for the available partition sizes.
	//
	//  If set, the accelerator_count should be set to 1.
	// +kcc:proto:field=google.cloud.aiplatform.v1beta1.MachineSpec.gpu_partition_size
	GpuPartitionSize *string `json:"gpuPartitionSize,omitempty"`

	// Immutable. The topology of the TPUs. Corresponds to the TPU topologies
	//  available from GKE. (Example: tpu_topology: "2x2x1").
	// +kcc:proto:field=google.cloud.aiplatform.v1.MachineSpec.tpu_topology
	TpuTopology *string `json:"tpuTopology,omitempty"`

	// Optional. Immutable. The number of nodes per replica for multihost GPU
	//  deployments.
	// +kcc:proto:field=google.cloud.aiplatform.v1beta1.MachineSpec.multihost_gpu_node_count
	MultihostGpuNodeCount *int32 `json:"multihostGpuNodeCount,omitempty"`

	// Optional. Immutable. Configuration controlling how this resource pool
	//  consumes reservation.
	// +kcc:proto:field=google.cloud.aiplatform.v1.MachineSpec.reservation_affinity
	ReservationAffinity *ReservationAffinity `json:"reservationAffinity,omitempty"`
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
