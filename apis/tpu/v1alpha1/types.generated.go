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

// +kcc:proto=google.cloud.tpu.v1.NetworkEndpoint
type NetworkEndpoint struct {
	// The IP address of this network endpoint.
	// +kcc:proto:field=google.cloud.tpu.v1.NetworkEndpoint.ip_address
	IPAddress *string `json:"ipAddress,omitempty"`

	// The port of this network endpoint.
	// +kcc:proto:field=google.cloud.tpu.v1.NetworkEndpoint.port
	Port *int32 `json:"port,omitempty"`
}

// +kcc:proto=google.cloud.tpu.v1.SchedulingConfig
type SchedulingConfig struct {
	// Defines whether the node is preemptible.
	// +kcc:proto:field=google.cloud.tpu.v1.SchedulingConfig.preemptible
	Preemptible *bool `json:"preemptible,omitempty"`

	// Whether the node is created under a reservation.
	// +kcc:proto:field=google.cloud.tpu.v1.SchedulingConfig.reserved
	Reserved *bool `json:"reserved,omitempty"`
}

// +kcc:proto=google.cloud.tpu.v1.Symptom
type Symptom struct {
	// Timestamp when the Symptom is created.
	// +kcc:proto:field=google.cloud.tpu.v1.Symptom.create_time
	CreateTime *string `json:"createTime,omitempty"`

	// Type of the Symptom.
	// +kcc:proto:field=google.cloud.tpu.v1.Symptom.symptom_type
	SymptomType *string `json:"symptomType,omitempty"`

	// Detailed information of the current Symptom.
	// +kcc:proto:field=google.cloud.tpu.v1.Symptom.details
	Details *string `json:"details,omitempty"`

	// A string used to uniquely distinguish a worker within a TPU node.
	// +kcc:proto:field=google.cloud.tpu.v1.Symptom.worker_id
	WorkerID *string `json:"workerID,omitempty"`
}
