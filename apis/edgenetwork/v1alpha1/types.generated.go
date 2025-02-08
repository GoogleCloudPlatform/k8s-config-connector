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


// +kcc:proto=google.cloud.edgenetwork.v1.Interconnect
type Interconnect struct {
	// Required. The canonical resource name of the interconnect.
	// +kcc:proto:field=google.cloud.edgenetwork.v1.Interconnect.name
	Name *string `json:"name,omitempty"`

	// Labels associated with this resource.
	// +kcc:proto:field=google.cloud.edgenetwork.v1.Interconnect.labels
	Labels map[string]string `json:"labels,omitempty"`

	// Optional. A free-text description of the resource. Max length 1024
	//  characters.
	// +kcc:proto:field=google.cloud.edgenetwork.v1.Interconnect.description
	Description *string `json:"description,omitempty"`

	// Optional. Type of interconnect, which takes only the value 'DEDICATED' for
	//  now.
	// +kcc:proto:field=google.cloud.edgenetwork.v1.Interconnect.interconnect_type
	InterconnectType *string `json:"interconnectType,omitempty"`
}

// +kcc:proto=google.cloud.edgenetwork.v1.Interconnect
type InterconnectObservedState struct {
	// Output only. The time when the subnet was created.
	// +kcc:proto:field=google.cloud.edgenetwork.v1.Interconnect.create_time
	CreateTime *string `json:"createTime,omitempty"`

	// Output only. The time when the subnet was last updated.
	// +kcc:proto:field=google.cloud.edgenetwork.v1.Interconnect.update_time
	UpdateTime *string `json:"updateTime,omitempty"`

	// Output only. Unique identifier for the link.
	// +kcc:proto:field=google.cloud.edgenetwork.v1.Interconnect.uuid
	Uuid *string `json:"uuid,omitempty"`

	// Output only. Cloud resource name of the switch device.
	// +kcc:proto:field=google.cloud.edgenetwork.v1.Interconnect.device_cloud_resource_name
	DeviceCloudResourceName *string `json:"deviceCloudResourceName,omitempty"`

	// Output only. Physical ports (e.g., TenGigE0/0/0/1) that form the
	//  interconnect.
	// +kcc:proto:field=google.cloud.edgenetwork.v1.Interconnect.physical_ports
	PhysicalPorts []string `json:"physicalPorts,omitempty"`
}
