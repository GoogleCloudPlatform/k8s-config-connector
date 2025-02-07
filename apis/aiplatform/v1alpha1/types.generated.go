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


// +kcc:proto=google.cloud.aiplatform.v1.EncryptionSpec
type EncryptionSpec struct {
	// Required. The Cloud KMS resource identifier of the customer managed
	//  encryption key used to protect a resource. Has the form:
	//  `projects/my-project/locations/my-region/keyRings/my-kr/cryptoKeys/my-key`.
	//  The key needs to be in the same region as where the compute resource is
	//  created.
	// +kcc:proto:field=google.cloud.aiplatform.v1.EncryptionSpec.kms_key_name
	KMSKeyName *string `json:"kmsKeyName,omitempty"`
}

// +kcc:proto=google.cloud.aiplatform.v1.MetadataStore
type MetadataStore struct {

	// Customer-managed encryption key spec for a Metadata Store. If set, this
	//  Metadata Store and all sub-resources of this Metadata Store are secured
	//  using this key.
	// +kcc:proto:field=google.cloud.aiplatform.v1.MetadataStore.encryption_spec
	EncryptionSpec *EncryptionSpec `json:"encryptionSpec,omitempty"`

	// Description of the MetadataStore.
	// +kcc:proto:field=google.cloud.aiplatform.v1.MetadataStore.description
	Description *string `json:"description,omitempty"`

	// Optional. Dataplex integration settings.
	// +kcc:proto:field=google.cloud.aiplatform.v1.MetadataStore.dataplex_config
	DataplexConfig *MetadataStore_DataplexConfig `json:"dataplexConfig,omitempty"`
}

// +kcc:proto=google.cloud.aiplatform.v1.MetadataStore.DataplexConfig
type MetadataStore_DataplexConfig struct {
	// Optional. Whether or not Data Lineage synchronization is enabled for
	//  Vertex Pipelines.
	// +kcc:proto:field=google.cloud.aiplatform.v1.MetadataStore.DataplexConfig.enabled_pipelines_lineage
	EnabledPipelinesLineage *bool `json:"enabledPipelinesLineage,omitempty"`
}

// +kcc:proto=google.cloud.aiplatform.v1.MetadataStore.MetadataStoreState
type MetadataStore_MetadataStoreState struct {
	// The disk utilization of the MetadataStore in bytes.
	// +kcc:proto:field=google.cloud.aiplatform.v1.MetadataStore.MetadataStoreState.disk_utilization_bytes
	DiskUtilizationBytes *int64 `json:"diskUtilizationBytes,omitempty"`
}

// +kcc:proto=google.cloud.aiplatform.v1.MetadataStore
type MetadataStoreObservedState struct {
	// Output only. The resource name of the MetadataStore instance.
	// +kcc:proto:field=google.cloud.aiplatform.v1.MetadataStore.name
	Name *string `json:"name,omitempty"`

	// Output only. Timestamp when this MetadataStore was created.
	// +kcc:proto:field=google.cloud.aiplatform.v1.MetadataStore.create_time
	CreateTime *string `json:"createTime,omitempty"`

	// Output only. Timestamp when this MetadataStore was last updated.
	// +kcc:proto:field=google.cloud.aiplatform.v1.MetadataStore.update_time
	UpdateTime *string `json:"updateTime,omitempty"`

	// Output only. State information of the MetadataStore.
	// +kcc:proto:field=google.cloud.aiplatform.v1.MetadataStore.state
	State *MetadataStore_MetadataStoreState `json:"state,omitempty"`
}
