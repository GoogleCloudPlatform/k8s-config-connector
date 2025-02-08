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

// +kcc:proto=google.cloud.tpu.v2.AcceleratorType
type AcceleratorType struct {
	// The resource name.
	// +kcc:proto:field=google.cloud.tpu.v2.AcceleratorType.name
	Name *string `json:"name,omitempty"`

	// The accelerator type.
	// +kcc:proto:field=google.cloud.tpu.v2.AcceleratorType.type
	Type *string `json:"type,omitempty"`

	// The accelerator config.
	// +kcc:proto:field=google.cloud.tpu.v2.AcceleratorType.accelerator_configs
	AcceleratorConfigs []AcceleratorConfig `json:"acceleratorConfigs,omitempty"`
}
