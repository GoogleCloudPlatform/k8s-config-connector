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

// +kcc:proto=google.cloud.aiplatform.v1beta1.Featurestore.OnlineServingConfig
type Featurestore_OnlineServingConfig struct {
	// The number of nodes for the online store. The number of nodes doesn't
	//  scale automatically, but you can manually update the number of
	//  nodes. If set to 0, the featurestore will not have an
	//  online store and cannot be used for online serving.
	// +kcc:proto:field=google.cloud.aiplatform.v1beta1.Featurestore.OnlineServingConfig.fixed_node_count
	FixedNodeCount *int32 `json:"fixedNodeCount,omitempty"`

	// Online serving scaling configuration.
	//  Only one of `fixed_node_count` and `scaling` can be set. Setting one will
	//  reset the other.
	// +kcc:proto:field=google.cloud.aiplatform.v1beta1.Featurestore.OnlineServingConfig.scaling
	Scaling *Featurestore_OnlineServingConfig_Scaling `json:"scaling,omitempty"`
}

// +kcc:proto=google.cloud.aiplatform.v1beta1.Featurestore.OnlineServingConfig.Scaling
type Featurestore_OnlineServingConfig_Scaling struct {
	// Required. The minimum number of nodes to scale down to. Must be greater
	//  than or equal to 1.
	// +kcc:proto:field=google.cloud.aiplatform.v1beta1.Featurestore.OnlineServingConfig.Scaling.min_node_count
	MinNodeCount *int32 `json:"minNodeCount,omitempty"`

	// The maximum number of nodes to scale up to. Must be greater than
	//  min_node_count, and less than or equal to 10 times of 'min_node_count'.
	// +kcc:proto:field=google.cloud.aiplatform.v1beta1.Featurestore.OnlineServingConfig.Scaling.max_node_count
	MaxNodeCount *int32 `json:"maxNodeCount,omitempty"`

	// Optional. The cpu utilization that the Autoscaler should be trying to
	//  achieve. This number is on a scale from 0 (no utilization) to 100
	//  (total utilization), and is limited between 10 and 80. When a cluster's
	//  CPU utilization exceeds the target that you have set, Bigtable
	//  immediately adds nodes to the cluster. When CPU utilization is
	//  substantially lower than the target, Bigtable removes nodes. If not set
	//  or set to 0, default to 50.
	// +kcc:proto:field=google.cloud.aiplatform.v1beta1.Featurestore.OnlineServingConfig.Scaling.cpu_utilization_target
	CPUUtilizationTarget *int32 `json:"cpuUtilizationTarget,omitempty"`
}
import (
	refsv1beta1 "github.com/GoogleCloudPlatform/k8s-config-connector/apis/refs/v1beta1"
)

// +kcc:proto=google.cloud.aiplatform.v1beta1.EncryptionSpec
type EncryptionSpec struct {
	// Required. The Cloud KMS resource identifier of the customer managed
	//  encryption key used to protect a resource. Has the form:
	//  `projects/my-project/locations/my-region/keyRings/my-kr/cryptoKeys/my-key`.
	//  The key needs to be in the same region as where the compute resource is
	//  created.
	// +kcc:proto:field=google.cloud.aiplatform.v1beta1.EncryptionSpec.kms_key_ref
	KMSKeyRef *refsv1beta1.KMSCryptoKeyRef `json:"kmsKeyRef,omitempty"`
}

// +kcc:proto=google.cloud.aiplatform.v1beta1.MetadataStore
type MetadataStore struct {

	// Customer-managed encryption key spec for a Metadata Store. If set, this
	//  Metadata Store and all sub-resources of this Metadata Store are secured
	//  using this key.
	// +kcc:proto:field=google.cloud.aiplatform.v1beta1.MetadataStore.encryption_spec
	EncryptionSpec *EncryptionSpec `json:"encryptionSpec,omitempty"`

	// Description of the MetadataStore.
	// +kcc:proto:field=google.cloud.aiplatform.v1beta1.MetadataStore.description
	Description *string `json:"description,omitempty"`

	// Optional. Dataplex integration settings.
	// +kcc:proto:field=google.cloud.aiplatform.v1beta1.MetadataStore.dataplex_config
	DataplexConfig *MetadataStore_DataplexConfig `json:"dataplexConfig,omitempty"`
}

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
