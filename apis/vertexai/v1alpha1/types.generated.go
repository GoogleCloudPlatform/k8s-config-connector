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

// +kcc:proto=google.cloud.aiplatform.v1beta1.EncryptionSpec
type EncryptionSpec struct {
	// Required. The Cloud KMS resource identifier of the customer managed
	//  encryption key used to protect a resource. Has the form:
	//  `projects/my-project/locations/my-region/keyRings/my-kr/cryptoKeys/my-key`.
	//  The key needs to be in the same region as where the compute resource is
	//  created.
	// +kcc:proto:field=google.cloud.aiplatform.v1beta1.EncryptionSpec.kms_key_name
	KMSKeyName *string `json:"kmsKeyName,omitempty"`
}

// +kcc:proto=google.cloud.aiplatform.v1beta1.Featurestore
type Featurestore struct {

	// Optional. Used to perform consistent read-modify-write updates. If not set,
	//  a blind "overwrite" update happens.
	// +kcc:proto:field=google.cloud.aiplatform.v1beta1.Featurestore.etag
	Etag *string `json:"etag,omitempty"`

	// Optional. The labels with user-defined metadata to organize your
	//  Featurestore.
	//
	//  Label keys and values can be no longer than 64 characters
	//  (Unicode codepoints), can only contain lowercase letters, numeric
	//  characters, underscores and dashes. International characters are allowed.
	//
	//  See https://goo.gl/xmQnxf for more information on and examples of labels.
	//  No more than 64 user labels can be associated with one Featurestore(System
	//  labels are excluded)."
	//  System reserved label keys are prefixed with "aiplatform.googleapis.com/"
	//  and are immutable.
	// +kcc:proto:field=google.cloud.aiplatform.v1beta1.Featurestore.labels
	Labels map[string]string `json:"labels,omitempty"`

	// Optional. Config for online storage resources. The field should not
	//  co-exist with the field of `OnlineStoreReplicationConfig`. If both of it
	//  and OnlineStoreReplicationConfig are unset, the feature store will not have
	//  an online store and cannot be used for online serving.
	// +kcc:proto:field=google.cloud.aiplatform.v1beta1.Featurestore.online_serving_config
	OnlineServingConfig *Featurestore_OnlineServingConfig `json:"onlineServingConfig,omitempty"`

	// Optional. TTL in days for feature values that will be stored in online
	//  serving storage. The Feature Store online storage periodically removes
	//  obsolete feature values older than `online_storage_ttl_days` since the
	//  feature generation time. Note that `online_storage_ttl_days` should be less
	//  than or equal to `offline_storage_ttl_days` for each EntityType under a
	//  featurestore. If not set, default to 4000 days
	// +kcc:proto:field=google.cloud.aiplatform.v1beta1.Featurestore.online_storage_ttl_days
	OnlineStorageTtlDays *int32 `json:"onlineStorageTtlDays,omitempty"`

	// Optional. Customer-managed encryption key spec for data storage. If set,
	//  both of the online and offline data storage will be secured by this key.
	// +kcc:proto:field=google.cloud.aiplatform.v1beta1.Featurestore.encryption_spec
	EncryptionSpec *EncryptionSpec `json:"encryptionSpec,omitempty"`
}

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

// +kcc:proto=google.cloud.aiplatform.v1beta1.Featurestore
type FeaturestoreObservedState struct {
	// Output only. Name of the Featurestore. Format:
	//  `projects/{project}/locations/{location}/featurestores/{featurestore}`
	// +kcc:proto:field=google.cloud.aiplatform.v1beta1.Featurestore.name
	Name *string `json:"name,omitempty"`

	// Output only. Timestamp when this Featurestore was created.
	// +kcc:proto:field=google.cloud.aiplatform.v1beta1.Featurestore.create_time
	CreateTime *string `json:"createTime,omitempty"`

	// Output only. Timestamp when this Featurestore was last updated.
	// +kcc:proto:field=google.cloud.aiplatform.v1beta1.Featurestore.update_time
	UpdateTime *string `json:"updateTime,omitempty"`

	// Output only. State of the featurestore.
	// +kcc:proto:field=google.cloud.aiplatform.v1beta1.Featurestore.state
	State *string `json:"state,omitempty"`

	// Output only. Reserved for future use.
	// +kcc:proto:field=google.cloud.aiplatform.v1beta1.Featurestore.satisfies_pzs
	SatisfiesPzs *bool `json:"satisfiesPzs,omitempty"`

	// Output only. Reserved for future use.
	// +kcc:proto:field=google.cloud.aiplatform.v1beta1.Featurestore.satisfies_pzi
	SatisfiesPzi *bool `json:"satisfiesPzi,omitempty"`
}
