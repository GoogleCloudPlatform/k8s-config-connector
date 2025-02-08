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


// +kcc:proto=google.cloud.networkservices.v1.Mesh
type Mesh struct {
	// Required. Name of the Mesh resource. It matches pattern
	//  `projects/*/locations/global/meshes/<mesh_name>`.
	// +kcc:proto:field=google.cloud.networkservices.v1.Mesh.name
	Name *string `json:"name,omitempty"`

	// Optional. Set of label tags associated with the Mesh resource.
	// +kcc:proto:field=google.cloud.networkservices.v1.Mesh.labels
	Labels map[string]string `json:"labels,omitempty"`

	// Optional. A free-text description of the resource. Max length 1024
	//  characters.
	// +kcc:proto:field=google.cloud.networkservices.v1.Mesh.description
	Description *string `json:"description,omitempty"`

	// Optional. If set to a valid TCP port (1-65535), instructs the SIDECAR proxy
	//  to listen on the specified port of localhost (127.0.0.1) address. The
	//  SIDECAR proxy will expect all traffic to be redirected to this port
	//  regardless of its actual ip:port destination. If unset, a port '15001' is
	//  used as the interception port. This is applicable only for sidecar proxy
	//  deployments.
	// +kcc:proto:field=google.cloud.networkservices.v1.Mesh.interception_port
	InterceptionPort *int32 `json:"interceptionPort,omitempty"`
}

// +kcc:proto=google.cloud.networkservices.v1.Mesh
type MeshObservedState struct {
	// Output only. Server-defined URL of this resource
	// +kcc:proto:field=google.cloud.networkservices.v1.Mesh.self_link
	SelfLink *string `json:"selfLink,omitempty"`

	// Output only. The timestamp when the resource was created.
	// +kcc:proto:field=google.cloud.networkservices.v1.Mesh.create_time
	CreateTime *string `json:"createTime,omitempty"`

	// Output only. The timestamp when the resource was updated.
	// +kcc:proto:field=google.cloud.networkservices.v1.Mesh.update_time
	UpdateTime *string `json:"updateTime,omitempty"`
}
