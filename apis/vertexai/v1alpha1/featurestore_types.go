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

var VertexAIFeaturestoreGVK = GroupVersion.WithKind("VertexAIFeaturestore")

// +kcc:spec:proto=google.cloud.aiplatform.v1beta1.EncryptionSpec
type EncryptionSpec struct {
	// Required. The Cloud KMS resource identifier of the customer managed
	//  encryption key used to protect a resource.
	//  The key needs to be in the same region as where the compute resource is
	//  created.
	// +kcc:proto:field=google.cloud.aiplatform.v1beta1.EncryptionSpec.kms_key_name
	// +required
	KMSKeyRef *refsv1beta1.KMSCryptoKeyRef `json:"kmsKeyRef,omitempty"`
}

// VertexAIFeaturestoreSpec defines the desired state of VertexAIFeaturestore
// +kcc:spec:proto=google.cloud.aiplatform.v1beta1.Featurestore
type VertexAIFeaturestoreSpec struct {
	// The Project that this resource belongs to.
	// +required
	ProjectRef *refsv1beta1.ProjectRef `json:"projectRef"`

	// The location for the resource.
	// +required
	Location string `json:"location"`

	// The VertexAIFeaturestore name. If not given, the metadata.name will be used.
	ResourceID *string `json:"resourceID,omitempty"`

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
	OnlineStorageTTLDays *int32 `json:"onlineStorageTTLDays,omitempty"`

	// Optional. Customer-managed encryption key spec for data storage. If set,
	//  both of the online and offline data storage will be secured by this key.
	// +kcc:proto:field=google.cloud.aiplatform.v1beta1.Featurestore.encryption_spec
	EncryptionSpec *EncryptionSpec `json:"encryptionSpec,omitempty"`
}

// VertexAIFeaturestoreStatus defines the config connector machine state of VertexAIFeaturestore
type VertexAIFeaturestoreStatus struct {
	/* Conditions represent the latest available observations of the
	   object's current state. */
	Conditions []v1alpha1.Condition `json:"conditions,omitempty"`

	// ObservedGeneration is the generation of the resource that was most recently observed by the Config Connector controller. If this is equal to metadata.generation, then that means that the current reported status reflects the most recent desired state of the resource.
	ObservedGeneration *int64 `json:"observedGeneration,omitempty"`

	// A unique specifier for the VertexAIFeaturestore resource in GCP.
	ExternalRef *string `json:"externalRef,omitempty"`

	// ObservedState is the state of the resource as most recently observed in GCP.
	ObservedState *VertexAIFeaturestoreObservedState `json:"observedState,omitempty"`
}

// VertexAIFeaturestoreObservedState is the state of the VertexAIFeaturestore resource as most recently observed in GCP.
// +kcc:observedstate:proto=google.cloud.aiplatform.v1beta1.Featurestore
type VertexAIFeaturestoreObservedState struct {
	// Output only. Name of the Featurestore. Format:
	//  `projects/{project}/locations/{location}/featurestores/{featurestore}`
	// +kcc:proto:field=google.cloud.aiplatform.v1beta1.Featurestore.name
	// NOTYET
	// Name *string `json:"name,omitempty"`

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
	// NOTYET
	// SatisfiesPzs *bool `json:"satisfiesPzs,omitempty"`

	// Output only. Reserved for future use.
	// +kcc:proto:field=google.cloud.aiplatform.v1beta1.Featurestore.satisfies_pzi
	// NOTYET
	// SatisfiesPzi *bool `json:"satisfiesPzi,omitempty"`
}

// +genclient
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:resource:categories=gcp,shortName=gcpvertexaifeaturestore;gcpvertexaifeaturestores
// +kubebuilder:subresource:status
// +kubebuilder:metadata:labels="cnrm.cloud.google.com/managed-by-kcc=true";"cnrm.cloud.google.com/system=true"
// +kubebuilder:printcolumn:name="Age",JSONPath=".metadata.creationTimestamp",type="date"
// +kubebuilder:printcolumn:name="Ready",JSONPath=".status.conditions[?(@.type=='Ready')].status",type="string",description="When 'True', the most recent reconcile of the resource succeeded"
// +kubebuilder:printcolumn:name="Status",JSONPath=".status.conditions[?(@.type=='Ready')].reason",type="string",description="The reason for the value in 'Ready'"
// +kubebuilder:printcolumn:name="Status Age",JSONPath=".status.conditions[?(@.type=='Ready')].lastTransitionTime",type="date",description="The last transition time for the value in 'Status'"

// VertexAIFeaturestore is the Schema for the VertexAIFeaturestore API
// +k8s:openapi-gen=true
type VertexAIFeaturestore struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	// +required
	Spec   VertexAIFeaturestoreSpec   `json:"spec,omitempty"`
	Status VertexAIFeaturestoreStatus `json:"status,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// VertexAIFeaturestoreList contains a list of VertexAIFeaturestore
type VertexAIFeaturestoreList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []VertexAIFeaturestore `json:"items"`
}

func init() {
	SchemeBuilder.Register(&VertexAIFeaturestore{}, &VertexAIFeaturestoreList{})
}

type Featurestore struct {
	// Output only. Timestamp when this Featurestore was created.
	CreateTime *string `json:"createTime,omitempty"`

	// Optional. Used to perform consistent read-modify-write updates.
	Etag *string `json:"etag,omitempty"`

	// Optional. The labels with user-defined metadata to organize your Featurestore.
	Labels map[string]string `json:"labels,omitempty"`

	// Output only. Name of the Featurestore. Format:
	// `projects/{project}/locations/{location}/featurestores/{featurestore}`
	Name *string `json:"name,omitempty"`

	// Optional. Config for online serving resources.
	OnlineServingConfig *Featurestore_OnlineServingConfig `json:"onlineServingConfig,omitempty"`

	// Output only. State of the featurestore.
	State *string `json:"state,omitempty"`

	// Output only. Timestamp when this Featurestore was last updated.
	UpdateTime *string `json:"updateTime,omitempty"`
}
