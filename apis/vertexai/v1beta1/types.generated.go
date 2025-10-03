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
// krm.group: vertexai.cnrm.cloud.google.com
// krm.version: v1beta1
// proto.service: google.cloud.aiplatform.v1beta1
// resource: VertexAIMetadataStore:MetadataStore

package v1beta1

// +kcc:proto=google.cloud.aiplatform.v1beta1.MetadataStore.DataplexConfig
type MetadataStore_DataplexConfig struct {
	// Optional. Whether or not Data Lineage synchronization is enabled for
	//  Vertex Pipelines.
	// +kcc:proto:field=google.cloud.aiplatform.v1beta1.MetadataStore.DataplexConfig.enabled_pipelines_lineage
	EnabledPipelinesLineage *bool `json:"enabledPipelinesLineage,omitempty"`
}

// +kcc:proto=google.cloud.aiplatform.v1beta1.MetadataStore.MetadataStoreState
type MetadataStore_MetadataStoreState struct {
	// The disk utilization of the MetadataStore in bytes.
	// +kcc:proto:field=google.cloud.aiplatform.v1beta1.MetadataStore.MetadataStoreState.disk_utilization_bytes
	DiskUtilizationBytes *int64 `json:"diskUtilizationBytes,omitempty"`
}
