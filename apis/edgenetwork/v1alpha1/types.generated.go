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


// +kcc:proto=google.cloud.edgenetwork.v1.Network
type Network struct {
	// Required. The canonical resource name of the network.
	// +kcc:proto:field=google.cloud.edgenetwork.v1.Network.name
	Name *string `json:"name,omitempty"`

	// Labels associated with this resource.
	// +kcc:proto:field=google.cloud.edgenetwork.v1.Network.labels
	Labels map[string]string `json:"labels,omitempty"`

	// Optional. A free-text description of the resource. Max length 1024
	//  characters.
	// +kcc:proto:field=google.cloud.edgenetwork.v1.Network.description
	Description *string `json:"description,omitempty"`

	// IP (L3) MTU value of the network.
	//  Valid values are: 1500 and 9000.
	//  Default to 1500 if not set.
	// +kcc:proto:field=google.cloud.edgenetwork.v1.Network.mtu
	Mtu *int32 `json:"mtu,omitempty"`
}

// +kcc:proto=google.cloud.edgenetwork.v1.Network
type NetworkObservedState struct {
	// Output only. The time when the network was created.
	// +kcc:proto:field=google.cloud.edgenetwork.v1.Network.create_time
	CreateTime *string `json:"createTime,omitempty"`

	// Output only. The time when the network was last updated.
	// +kcc:proto:field=google.cloud.edgenetwork.v1.Network.update_time
	UpdateTime *string `json:"updateTime,omitempty"`
}
