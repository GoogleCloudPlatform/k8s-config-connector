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


// +kcc:proto=google.cloud.visionai.v1.Cluster
type Cluster struct {

	// Labels as key value pairs
	// +kcc:proto:field=google.cloud.visionai.v1.Cluster.labels
	Labels map[string]string `json:"labels,omitempty"`

	// Annotations to allow clients to store small amounts of arbitrary data.
	// +kcc:proto:field=google.cloud.visionai.v1.Cluster.annotations
	Annotations map[string]string `json:"annotations,omitempty"`
}

// +kcc:proto=google.cloud.visionai.v1.Cluster
type ClusterObservedState struct {
	// Output only. Name of the resource.
	// +kcc:proto:field=google.cloud.visionai.v1.Cluster.name
	Name *string `json:"name,omitempty"`

	// Output only. The create timestamp.
	// +kcc:proto:field=google.cloud.visionai.v1.Cluster.create_time
	CreateTime *string `json:"createTime,omitempty"`

	// Output only. The update timestamp.
	// +kcc:proto:field=google.cloud.visionai.v1.Cluster.update_time
	UpdateTime *string `json:"updateTime,omitempty"`

	// Output only. The DNS name of the data plane service
	// +kcc:proto:field=google.cloud.visionai.v1.Cluster.dataplane_service_endpoint
	DataplaneServiceEndpoint *string `json:"dataplaneServiceEndpoint,omitempty"`

	// Output only. The current state of the cluster.
	// +kcc:proto:field=google.cloud.visionai.v1.Cluster.state
	State *string `json:"state,omitempty"`

	// Output only. The private service connection service target name.
	// +kcc:proto:field=google.cloud.visionai.v1.Cluster.psc_target
	PscTarget *string `json:"pscTarget,omitempty"`
}
