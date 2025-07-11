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
// krm.group: networksecurity.cnrm.cloud.google.com
// krm.version: v1alpha1
// proto.service: google.cloud.networksecurity.v1
// resource: NetworkSecurityAddressGroup:AddressGroup

package v1alpha1

// +kcc:proto=google.cloud.networksecurity.v1.AddressGroup
type AddressGroup struct {
	// Required. Capacity of the Address Group
	// +kcc:proto:field=google.cloud.networksecurity.v1.AddressGroup.capacity
	Capacity *int32 `json:"capacity,omitempty"`

	// Output only. The timestamp when the resource was created.
	// +kcc:proto:field=google.cloud.networksecurity.v1.AddressGroup.create_time
	CreateTime *string `json:"createTime,omitempty"`

	// Optional. Free-text description of the resource.
	// +kcc:proto:field=google.cloud.networksecurity.v1.AddressGroup.description
	Description *string `json:"description,omitempty"`

	// Optional. List of items.
	// +kcc:proto:field=google.cloud.networksecurity.v1.AddressGroup.items
	Items []string `json:"items,omitempty"`

	// Optional. Set of label tags associated with the AddressGroup resource.
	// +kcc:proto:field=google.cloud.networksecurity.v1.AddressGroup.labels
	Labels map[string]string `json:"labels,omitempty"`

	// Required. Name of the AddressGroup resource. It matches pattern `projects/*/locations/{location}/addressGroups/`.
	// +kcc:proto:field=google.cloud.networksecurity.v1.AddressGroup.name
	Name *string `json:"name,omitempty"`

	// Optional. List of supported purposes of the Address Group.
	// +kcc:proto:field=google.cloud.networksecurity.v1.AddressGroup.purpose
	Purpose []string `json:"purpose,omitempty"`

	// Output only. Server-defined fully-qualified URL for this resource.
	// +kcc:proto:field=google.cloud.networksecurity.v1.AddressGroup.self_link
	SelfLink *string `json:"selfLink,omitempty"`

	// Required. The type of the Address Group. Possible values are "IPv4" or "IPV6".
	// +kcc:proto:field=google.cloud.networksecurity.v1.AddressGroup.type
	Type *string `json:"type,omitempty"`

	// Output only. The timestamp when the resource was updated.
	// +kcc:proto:field=google.cloud.networksecurity.v1.AddressGroup.update_time
	UpdateTime *string `json:"updateTime,omitempty"`
}
