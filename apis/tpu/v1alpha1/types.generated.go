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

// +kcc:proto=google.cloud.tpu.v2.AcceleratorConfig
type AcceleratorConfig struct {
	// Required. Type of TPU.
	// +kcc:proto:field=google.cloud.tpu.v2.AcceleratorConfig.type
	Type *string `json:"type,omitempty"`

	// Required. Topology of TPU in chips.
	// +kcc:proto:field=google.cloud.tpu.v2.AcceleratorConfig.topology
	Topology *string `json:"topology,omitempty"`
}

// +kcc:proto=google.cloud.tpu.v2.AttachedDisk
type AttachedDisk struct {
	// Specifies the full path to an existing disk.
	//  For example: "projects/my-project/zones/us-central1-c/disks/my-disk".
	// +kcc:proto:field=google.cloud.tpu.v2.AttachedDisk.source_disk
	SourceDisk *string `json:"sourceDisk,omitempty"`

	// The mode in which to attach this disk.
	//  If not specified, the default is READ_WRITE mode.
	//  Only applicable to data_disks.
	// +kcc:proto:field=google.cloud.tpu.v2.AttachedDisk.mode
	Mode *string `json:"mode,omitempty"`
}

// +kcc:proto=google.cloud.tpu.v2.SchedulingConfig
type SchedulingConfig struct {
	// Defines whether the node is preemptible.
	// +kcc:proto:field=google.cloud.tpu.v2.SchedulingConfig.preemptible
	Preemptible *bool `json:"preemptible,omitempty"`

	// Whether the node is created under a reservation.
	// +kcc:proto:field=google.cloud.tpu.v2.SchedulingConfig.reserved
	Reserved *bool `json:"reserved,omitempty"`

	// Optional. Defines whether the node is Spot VM.
	// +kcc:proto:field=google.cloud.tpu.v2.SchedulingConfig.spot
	Spot *bool `json:"spot,omitempty"`
}

// +kcc:proto=google.cloud.tpu.v2.ShieldedInstanceConfig
type ShieldedInstanceConfig struct {
	// Defines whether the instance has Secure Boot enabled.
	// +kcc:proto:field=google.cloud.tpu.v2.ShieldedInstanceConfig.enable_secure_boot
	EnableSecureBoot *bool `json:"enableSecureBoot,omitempty"`
}

// +kcc:proto=google.cloud.tpu.v2.Symptom
type Symptom struct {
	// Timestamp when the Symptom is created.
	// +kcc:proto:field=google.cloud.tpu.v2.Symptom.create_time
	CreateTime *string `json:"createTime,omitempty"`

	// Type of the Symptom.
	// +kcc:proto:field=google.cloud.tpu.v2.Symptom.symptom_type
	SymptomType *string `json:"symptomType,omitempty"`

	// Detailed information of the current Symptom.
	// +kcc:proto:field=google.cloud.tpu.v2.Symptom.details
	Details *string `json:"details,omitempty"`

	// A string used to uniquely distinguish a worker within a TPU node.
	// +kcc:proto:field=google.cloud.tpu.v2.Symptom.worker_id
	WorkerID *string `json:"workerID,omitempty"`
}
