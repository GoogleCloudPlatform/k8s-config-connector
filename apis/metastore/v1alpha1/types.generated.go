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

// +kcc:proto=google.cloud.metastore.v1.BackendMetastore
type BackendMetastore struct {
	// The relative resource name of the metastore that is being federated.
	//  The formats of the relative resource names for the currently supported
	//  metastores are listed below:
	//
	//  * BigQuery
	//      * `projects/{project_id}`
	//  * Dataproc Metastore
	//      * `projects/{project_id}/locations/{location}/services/{service_id}`
	// +kcc:proto:field=google.cloud.metastore.v1.BackendMetastore.name
	Name *string `json:"name,omitempty"`

	// The type of the backend metastore.
	// +kcc:proto:field=google.cloud.metastore.v1.BackendMetastore.metastore_type
	MetastoreType *string `json:"metastoreType,omitempty"`
}

// +kcc:proto=google.cloud.metastore.v1.Federation
type Federation struct {
	// Immutable. The relative resource name of the federation, of the
	//  form:
	//  projects/{project_number}/locations/{location_id}/federations/{federation_id}`.
	// +kcc:proto:field=google.cloud.metastore.v1.Federation.name
	Name *string `json:"name,omitempty"`

	// User-defined labels for the metastore federation.
	// +kcc:proto:field=google.cloud.metastore.v1.Federation.labels
	Labels map[string]string `json:"labels,omitempty"`

	// Immutable. The Apache Hive metastore version of the federation. All backend
	//  metastore versions must be compatible with the federation version.
	// +kcc:proto:field=google.cloud.metastore.v1.Federation.version
	Version *string `json:"version,omitempty"`

	// TODO: unsupported map type with key int32 and value message

}

// +kcc:proto=google.cloud.metastore.v1.Federation
type FederationObservedState struct {
	// Output only. The time when the metastore federation was created.
	// +kcc:proto:field=google.cloud.metastore.v1.Federation.create_time
	CreateTime *string `json:"createTime,omitempty"`

	// Output only. The time when the metastore federation was last updated.
	// +kcc:proto:field=google.cloud.metastore.v1.Federation.update_time
	UpdateTime *string `json:"updateTime,omitempty"`

	// Output only. The federation endpoint.
	// +kcc:proto:field=google.cloud.metastore.v1.Federation.endpoint_uri
	EndpointURI *string `json:"endpointURI,omitempty"`

	// Output only. The current state of the federation.
	// +kcc:proto:field=google.cloud.metastore.v1.Federation.state
	State *string `json:"state,omitempty"`

	// Output only. Additional information about the current state of the
	//  metastore federation, if available.
	// +kcc:proto:field=google.cloud.metastore.v1.Federation.state_message
	StateMessage *string `json:"stateMessage,omitempty"`

	// Output only. The globally unique resource identifier of the metastore
	//  federation.
	// +kcc:proto:field=google.cloud.metastore.v1.Federation.uid
	Uid *string `json:"uid,omitempty"`
}
