// Copyright 2025 Google LLC
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//    http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package v1alpha1

import (
	refsv1beta1 "github.com/GoogleCloudPlatform/k8s-config-connector/apis/refs/v1beta1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/apis/k8s/v1alpha1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

var VertexAIFeatureOnlineStoreGVK = GroupVersion.WithKind("VertexAIFeatureOnlineStore")

// VertexAIFeatureOnlineStoreSpec defines the desired state of VertexAIFeatureOnlineStore
// +kcc:spec:proto=google.cloud.aiplatform.v1.FeatureOnlineStore
type VertexAIFeatureOnlineStoreSpec struct {
	// The project that this resource belongs to.
	ProjectRef *refsv1beta1.ProjectRef `json:"projectRef"`

	// The location of this resource.
	Location string `json:"location"`

	// The VertexAIFeatureOnlineStore name. If not given, the metadata.name will be used.
	ResourceID *string `json:"resourceID,omitempty"`

	// Contains settings for the Cloud Bigtable instance that will be created
	// to serve featureValues for all FeatureViews under this
	// FeatureOnlineStore.
	// +kubebuilder:validation:Optional
	Bigtable *FeatureOnlineStore_Bigtable `json:"bigtable,omitempty"`

	// Contains settings for the Optimized store that will be created
	// to serve featureValues for all FeatureViews under this
	// FeatureOnlineStore. When choose Optimized storage type, need to set
	// [PrivateServiceConnectConfig.enable_private_service_connect][google.cloud.aiplatform.v1.PrivateServiceConnectConfig.enable_private_service_connect]
	// to use private endpoint. Otherwise will use public endpoint by default.
	// +kubebuilder:validation:Optional
	Optimized *FeatureOnlineStore_Optimized `json:"optimized,omitempty"`

	// Optional. The labels with user-defined metadata to organize your
	// FeatureOnlineStore.
	// +kubebuilder:validation:Optional
	Labels map[string]string `json:"labels,omitempty"`

	// Optional. The dedicated serving endpoint for this FeatureOnlineStore, which
	// is different from common Vertex service endpoint.
	// +kubebuilder:validation:Optional
	DedicatedServingEndpoint *FeatureOnlineStore_DedicatedServingEndpoint `json:"dedicatedServingEndpoint,omitempty"`

	// Optional. Customer-managed encryption key spec for data storage. If set,
	// online store will be secured by this key.
	// +kubebuilder:validation:Optional
	EncryptionSpec *EncryptionSpec `json:"encryptionSpec,omitempty"`
}

// VertexAIFeatureOnlineStoreStatus defines the config connector machine state of VertexAIFeatureOnlineStore
type VertexAIFeatureOnlineStoreStatus struct {
	/* Conditions represent the latest available observations of the
	   object's current state. */
	Conditions []v1alpha1.Condition `json:"conditions,omitempty"`

	// ObservedGeneration is the generation of the resource that was most recently observed by the Config Connector controller. If this is equal to metadata.generation, then that means that the current reported status reflects the most recent desired state of the resource.
	ObservedGeneration *int64 `json:"observedGeneration,omitempty"`

	// A unique specifier for the VertexAIFeatureOnlineStore resource in GCP.
	ExternalRef *string `json:"externalRef,omitempty"`

	// ObservedState is the state of the resource as most recently observed in GCP.
	ObservedState *VertexAIFeatureOnlineStoreObservedState `json:"observedState,omitempty"`
}

// VertexAIFeatureOnlineStoreObservedState is the state of the VertexAIFeatureOnlineStore resource as most recently observed in GCP.
// +kcc:observedstate:proto=google.cloud.aiplatform.v1.FeatureOnlineStore
type VertexAIFeatureOnlineStoreObservedState struct {

	// Output only. Timestamp when this FeatureOnlineStore was created.
	CreateTime *string `json:"createTime,omitempty"`

	// Output only. Timestamp when this FeatureOnlineStore was last updated.
	UpdateTime *string `json:"updateTime,omitempty"`

	// Optional. Used to perform consistent read-modify-write updates. If not set,
	// a blind "overwrite" update happens.
	Etag *string `json:"etag,omitempty"`

	// Output only. State of the featureOnlineStore.
	State *string `json:"state,omitempty"`

	// Output only. Reserved for future use.
	SatisfiesPzs *bool `json:"satisfiesPzs,omitempty"`

	// Output only. Reserved for future use.
	SatisfiesPzi *bool `json:"satisfiesPzi,omitempty"`
}

// +genclient
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:resource:categories=gcp,shortName=gcpvertexaifeatureonlinestore;gcpvertexaifeatureonlinestores
// +kubebuilder:subresource:status
// +kubebuilder:metadata:labels="cnrm.cloud.google.com/managed-by-kcc=true"
// +kubebuilder:metadata:labels="cnrm.cloud.google.com/system=true"
// +kubebuilder:printcolumn:name="Age",JSONPath=".metadata.creationTimestamp",type="date"
// +kubebuilder:printcolumn:name="Ready",JSONPath=".status.conditions[?(@.type=='Ready')].status",type="string",description="When 'True', the most recent reconcile of the resource succeeded"
// +kubebuilder:printcolumn:name="Status",JSONPath=".status.conditions[?(@.type=='Ready')].reason",type="string",description="The reason for the value in 'Ready'"
// +kubebuilder:printcolumn:name="Status Age",JSONPath=".status.conditions[?(@.type=='Ready')].lastTransitionTime",type="date",description="The last transition time for the value in 'Status'"

// VertexAIFeatureOnlineStore is the Schema for the VertexAIFeatureOnlineStore API
// +k8s:openapi-gen=true
type VertexAIFeatureOnlineStore struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	// +required
	Spec   VertexAIFeatureOnlineStoreSpec   `json:"spec,omitempty"`
	Status VertexAIFeatureOnlineStoreStatus `json:"status,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// VertexAIFeatureOnlineStoreList contains a list of VertexAIFeatureOnlineStore
type VertexAIFeatureOnlineStoreList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []VertexAIFeatureOnlineStore `json:"items"`
}

func init() {
	SchemeBuilder.Register(&VertexAIFeatureOnlineStore{}, &VertexAIFeatureOnlineStoreList{})
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
	CPUUtilizationTarget *int32 `json:"cpuUtilizationTarget,omitempty"`
}

// +kcc:proto=google.cloud.aiplatform.v1.FeatureOnlineStore.Optimized
type FeatureOnlineStore_Optimized struct {
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

// +kcc:proto=google.cloud.aiplatform.v1.PrivateServiceConnectConfig
type PrivateServiceConnectConfig struct {
	// Required. If true, expose the IndexEndpoint via private service connect.
	// +kcc:proto:field=google.cloud.aiplatform.v1.PrivateServiceConnectConfig.enable_private_service_connect
	EnablePrivateServiceConnect *bool `json:"enablePrivateServiceConnect,omitempty"`

	// A list of Projects from which the forwarding rule will target the service
	//  attachment.
	// +kcc:proto:field=google.cloud.aiplatform.v1.PrivateServiceConnectConfig.project_allowlist
	ProjectAllowlist []string `json:"projectAllowlist,omitempty"`

	// Optional. List of projects and networks where the PSC endpoints will be
	//  created. This field is used by Online Inference(Prediction) only.
	// +kcc:proto:field=google.cloud.aiplatform.v1.PrivateServiceConnectConfig.psc_automation_configs
	PSCAutomationConfigs []PSCAutomationConfig `json:"pscAutomationConfigs,omitempty"`
}

// +kcc:proto=google.cloud.aiplatform.v1.PSCAutomationConfig
type PSCAutomationConfig struct {
	// Required. Project id used to create forwarding rule.
	// +kcc:proto:field=google.cloud.aiplatform.v1.PSCAutomationConfig.project_id
	ProjectID *string `json:"projectID,omitempty"`

	// Required. The full name of the Google Compute Engine
	//  [network](https://cloud.google.com/compute/docs/networks-and-firewalls#networks).
	//  [Format](https://cloud.google.com/compute/docs/reference/rest/v1/networks/get):
	//  `projects/{project}/global/networks/{network}`.
	// +kcc:proto:field=google.cloud.aiplatform.v1.PSCAutomationConfig.network
	Network *string `json:"network,omitempty"`
}
