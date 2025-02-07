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

// +kcc:proto=google.cloud.aiplatform.v1.FeatureOnlineStore
type FeatureOnlineStore struct {
	// Contains settings for the Cloud Bigtable instance that will be created
	//  to serve featureValues for all FeatureViews under this
	//  FeatureOnlineStore.
	// +kcc:proto:field=google.cloud.aiplatform.v1.FeatureOnlineStore.bigtable
	Bigtable *FeatureOnlineStore_Bigtable `json:"bigtable,omitempty"`

	// Contains settings for the Optimized store that will be created
	//  to serve featureValues for all FeatureViews under this
	//  FeatureOnlineStore. When choose Optimized storage type, need to set
	//  [PrivateServiceConnectConfig.enable_private_service_connect][google.cloud.aiplatform.v1.PrivateServiceConnectConfig.enable_private_service_connect]
	//  to use private endpoint. Otherwise will use public endpoint by default.
	// +kcc:proto:field=google.cloud.aiplatform.v1.FeatureOnlineStore.optimized
	Optimized *FeatureOnlineStore_Optimized `json:"optimized,omitempty"`

	// Identifier. Name of the FeatureOnlineStore. Format:
	//  `projects/{project}/locations/{location}/featureOnlineStores/{featureOnlineStore}`
	// +kcc:proto:field=google.cloud.aiplatform.v1.FeatureOnlineStore.name
	Name *string `json:"name,omitempty"`

	// Optional. Used to perform consistent read-modify-write updates. If not set,
	//  a blind "overwrite" update happens.
	// +kcc:proto:field=google.cloud.aiplatform.v1.FeatureOnlineStore.etag
	Etag *string `json:"etag,omitempty"`

	// Optional. The labels with user-defined metadata to organize your
	//  FeatureOnlineStore.
	//
	//  Label keys and values can be no longer than 64 characters
	//  (Unicode codepoints), can only contain lowercase letters, numeric
	//  characters, underscores and dashes. International characters are allowed.
	//
	//  See https://goo.gl/xmQnxf for more information on and examples of labels.
	//  No more than 64 user labels can be associated with one
	//  FeatureOnlineStore(System labels are excluded)." System reserved label keys
	//  are prefixed with "aiplatform.googleapis.com/" and are immutable.
	// +kcc:proto:field=google.cloud.aiplatform.v1.FeatureOnlineStore.labels
	Labels map[string]string `json:"labels,omitempty"`

	// Optional. The dedicated serving endpoint for this FeatureOnlineStore, which
	//  is different from common Vertex service endpoint.
	// +kcc:proto:field=google.cloud.aiplatform.v1.FeatureOnlineStore.dedicated_serving_endpoint
	DedicatedServingEndpoint *FeatureOnlineStore_DedicatedServingEndpoint `json:"dedicatedServingEndpoint,omitempty"`

	// Optional. Customer-managed encryption key spec for data storage. If set,
	//  online store will be secured by this key.
	// +kcc:proto:field=google.cloud.aiplatform.v1.FeatureOnlineStore.encryption_spec
	EncryptionSpec *EncryptionSpec `json:"encryptionSpec,omitempty"`
}

// +kcc:proto=google.cloud.aiplatform.v1.FeatureOnlineStore.Bigtable
type FeatureOnlineStore_Bigtable struct {
	// Required. Autoscaling config applied to Bigtable Instance.
	// +kcc:proto:field=google.cloud.aiplatform.v1.FeatureOnlineStore.Bigtable.auto_scaling
	AutoScaling *FeatureOnlineStore_Bigtable_AutoScaling `json:"autoScaling,omitempty"`
}

// +kcc:proto=google.cloud.aiplatform.v1.FeatureOnlineStore.Bigtable.AutoScaling
type FeatureOnlineStore_Bigtable_AutoScaling struct {
	// Required. The minimum number of nodes to scale down to. Must be greater
	//  than or equal to 1.
	// +kcc:proto:field=google.cloud.aiplatform.v1.FeatureOnlineStore.Bigtable.AutoScaling.min_node_count
	MinNodeCount *int32 `json:"minNodeCount,omitempty"`

	// Required. The maximum number of nodes to scale up to. Must be greater
	//  than or equal to min_node_count, and less than or equal to 10 times of
	//  'min_node_count'.
	// +kcc:proto:field=google.cloud.aiplatform.v1.FeatureOnlineStore.Bigtable.AutoScaling.max_node_count
	MaxNodeCount *int32 `json:"maxNodeCount,omitempty"`

	// Optional. A percentage of the cluster's CPU capacity. Can be from 10%
	//  to 80%. When a cluster's CPU utilization exceeds the target that you
	//  have set, Bigtable immediately adds nodes to the cluster. When CPU
	//  utilization is substantially lower than the target, Bigtable removes
	//  nodes. If not set will default to 50%.
	// +kcc:proto:field=google.cloud.aiplatform.v1.FeatureOnlineStore.Bigtable.AutoScaling.cpu_utilization_target
	CpuUtilizationTarget *int32 `json:"cpuUtilizationTarget,omitempty"`
}

// +kcc:proto=google.cloud.aiplatform.v1.FeatureOnlineStore.DedicatedServingEndpoint
type FeatureOnlineStore_DedicatedServingEndpoint struct {

	// Optional. Private service connect config. The private service connection
	//  is available only for Optimized storage type, not for embedding
	//  management now. If
	//  [PrivateServiceConnectConfig.enable_private_service_connect][google.cloud.aiplatform.v1.PrivateServiceConnectConfig.enable_private_service_connect]
	//  set to true, customers will use private service connection to send
	//  request. Otherwise, the connection will set to public endpoint.
	// +kcc:proto:field=google.cloud.aiplatform.v1.FeatureOnlineStore.DedicatedServingEndpoint.private_service_connect_config
	PrivateServiceConnectConfig *PrivateServiceConnectConfig `json:"privateServiceConnectConfig,omitempty"`
}

// +kcc:proto=google.cloud.aiplatform.v1.FeatureOnlineStore.Optimized
type FeatureOnlineStore_Optimized struct {
}

// +kcc:proto=google.cloud.aiplatform.v1.PrivateServiceConnectConfig
type PrivateServiceConnectConfig struct {
	// Required. If true, expose the IndexEndpoint via private service connect.
	// +kcc:proto:field=google.cloud.aiplatform.v1.PrivateServiceConnectConfig.enable_private_service_connect
	EnablePrivateServiceConnect *bool `json:"enablePrivateServiceConnect,omitempty"`

	// A list of Projects from which the forwarding rule will target the service
	//  attachment.
	// +kcc:proto:field=google.cloud.aiplatform.v1.PrivateServiceConnectConfig.project_allowlist
	ProjectAllowlist []string `json:"projectAllowlist,omitempty"`
}

// +kcc:proto=google.cloud.aiplatform.v1.FeatureOnlineStore
type FeatureOnlineStoreObservedState struct {
	// Output only. Timestamp when this FeatureOnlineStore was created.
	// +kcc:proto:field=google.cloud.aiplatform.v1.FeatureOnlineStore.create_time
	CreateTime *string `json:"createTime,omitempty"`

	// Output only. Timestamp when this FeatureOnlineStore was last updated.
	// +kcc:proto:field=google.cloud.aiplatform.v1.FeatureOnlineStore.update_time
	UpdateTime *string `json:"updateTime,omitempty"`

	// Output only. State of the featureOnlineStore.
	// +kcc:proto:field=google.cloud.aiplatform.v1.FeatureOnlineStore.state
	State *string `json:"state,omitempty"`

	// Optional. The dedicated serving endpoint for this FeatureOnlineStore, which
	//  is different from common Vertex service endpoint.
	// +kcc:proto:field=google.cloud.aiplatform.v1.FeatureOnlineStore.dedicated_serving_endpoint
	DedicatedServingEndpoint *FeatureOnlineStore_DedicatedServingEndpointObservedState `json:"dedicatedServingEndpoint,omitempty"`

	// Output only. Reserved for future use.
	// +kcc:proto:field=google.cloud.aiplatform.v1.FeatureOnlineStore.satisfies_pzs
	SatisfiesPzs *bool `json:"satisfiesPzs,omitempty"`

	// Output only. Reserved for future use.
	// +kcc:proto:field=google.cloud.aiplatform.v1.FeatureOnlineStore.satisfies_pzi
	SatisfiesPzi *bool `json:"satisfiesPzi,omitempty"`
}

// +kcc:proto=google.cloud.aiplatform.v1.FeatureOnlineStore.DedicatedServingEndpoint
type FeatureOnlineStore_DedicatedServingEndpointObservedState struct {
	// Output only. This field will be populated with the domain name to use for
	//  this FeatureOnlineStore
	// +kcc:proto:field=google.cloud.aiplatform.v1.FeatureOnlineStore.DedicatedServingEndpoint.public_endpoint_domain_name
	PublicEndpointDomainName *string `json:"publicEndpointDomainName,omitempty"`

	// Optional. Private service connect config. The private service connection
	//  is available only for Optimized storage type, not for embedding
	//  management now. If
	//  [PrivateServiceConnectConfig.enable_private_service_connect][google.cloud.aiplatform.v1.PrivateServiceConnectConfig.enable_private_service_connect]
	//  set to true, customers will use private service connection to send
	//  request. Otherwise, the connection will set to public endpoint.
	// +kcc:proto:field=google.cloud.aiplatform.v1.FeatureOnlineStore.DedicatedServingEndpoint.private_service_connect_config
	PrivateServiceConnectConfig *PrivateServiceConnectConfigObservedState `json:"privateServiceConnectConfig,omitempty"`

	// Output only. The name of the service attachment resource. Populated if
	//  private service connect is enabled and after FeatureViewSync is created.
	// +kcc:proto:field=google.cloud.aiplatform.v1.FeatureOnlineStore.DedicatedServingEndpoint.service_attachment
	ServiceAttachment *string `json:"serviceAttachment,omitempty"`
}

// +kcc:proto=google.cloud.aiplatform.v1.PrivateServiceConnectConfig
type PrivateServiceConnectConfigObservedState struct {
	// Output only. The name of the generated service attachment resource.
	//  This is only populated if the endpoint is deployed with
	//  PrivateServiceConnect.
	// +kcc:proto:field=google.cloud.aiplatform.v1.PrivateServiceConnectConfig.service_attachment
	ServiceAttachment *string `json:"serviceAttachment,omitempty"`
}
