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

var VertexAIMetadataStoreGVK = GroupVersion.WithKind("VertexAIMetadataStore")

// VertexAIMetadataStoreSpec defines the desired state of VertexAIMetadataStore
// +kcc:proto=google.cloud.aiplatform.v1beta1.MetadataStore
type VertexAIMetadataStoreSpec struct {
	// Customer-managed encryption key spec for a Metadata Store. If set, this
	// Metadata Store and all sub-resources of this Metadata Store are secured
	// using this key.
	// +kcc:proto:field=google.cloud.aiplatform.v1beta1.MetadataStore.encryption_spec
	EncryptionSpec *EncryptionSpec `json:"encryptionSpec,omitempty"`

	// Optional. Dataplex integration settings.
	// +kcc:proto:field=google.cloud.aiplatform.v1beta1.MetadataStore.dataplex_config
	DataplexConfig *MetadataStore_DataplexConfig `json:"dataplexConfig,omitempty"`

	// Description of the MetadataStore.
	// +kcc:proto:field=google.cloud.aiplatform.v1beta1.MetadataStore.description
	Description string `json:"description,omitempty"`

	// The region of the Metadata Store. eg us-central1.
	Region string `json:"region,omitempty"`

	// The project that this resource belongs to.
	ProjectRef *refsv1beta1.ProjectRef `json:"projectRef,omitempty"`

	// The VertexAIMetadataStore name. If not given, the metadata.name will be used.
	ResourceID *string `json:"resourceID,omitempty"`
}

// VertexAIMetadataStoreStatus defines the config connector machine state of VertexAIMetadataStore
type VertexAIMetadataStoreStatus struct {
	/* Conditions represent the latest available observations of the
	   object's current state. */
	Conditions []v1alpha1.Condition `json:"conditions,omitempty"`

	// ObservedGeneration is the generation of the resource that was most recently observed by the Config Connector controller. If this is equal to metadata.generation, then that means that the current reported status reflects the most recent desired state of the resource.
	ObservedGeneration *int64 `json:"observedGeneration,omitempty"`

	// A unique specifier for the VertexAIMetadataStore resource in GCP.
	ExternalRef *string `json:"externalRef,omitempty"`

	// ObservedState is the state of the resource as most recently observed in GCP.
	ObservedState *VertexAIMetadataStoreObservedState `json:"observedState,omitempty"`
}

// VertexAIMetadataStoreObservedState is the state of the VertexAIMetadataStore resource as most recently observed in GCP.
// +kcc:proto=google.cloud.aiplatform.v1beta1.MetadataStore
type VertexAIMetadataStoreObservedState struct {
	// Output only. The resource name of the MetadataStore instance.
	// +kcc:proto:field=google.cloud.aiplatform.v1beta1.MetadataStore.name
	Name *string `json:"name,omitempty"`

	// Output only. Timestamp when this MetadataStore was created.
	// +kcc:proto:field=google.cloud.aiplatform.v1beta1.MetadataStore.create_time
	CreateTime *string `json:"createTime,omitempty"`

	// Output only. Timestamp when this MetadataStore was last updated.
	// +kcc:proto:field=google.cloud.aiplatform.v1beta1.MetadataStore.update_time
	UpdateTime *string `json:"updateTime,omitempty"`

	// Output only. State information of the MetadataStore.
	// +kcc:proto:field=google.cloud.aiplatform.v1beta1.MetadataStore.state
	State *MetadataStore_MetadataStoreState `json:"state,omitempty"`
}

// +kcc:proto=google.cloud.aiplatform.v1beta1.EncryptionSpec
type EncryptionSpec struct {
	// Required. The Cloud KMS resource identifier of the customer managed
	//  encryption key used to protect a resource. Has the form:
	//  `projects/my-project/locations/my-region/keyRings/my-kr/cryptoKeys/my-key`.
	//  The key needs to be in the same region as where the compute resource is
	//  created.
	// +kcc:proto:field=google.cloud.aiplatform.v1beta1.EncryptionSpec.kms_key_name
	KMSKeyNameRef *refsv1beta1.KMSCryptoKeyRef `json:"kmsKeyNameRef,omitempty"`
}

// +genclient
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// TODO(user): make sure the pluralizaiton below is correct"
// +kubebuilder:resource:categories=gcp,shortName=gcpvertexaimetadatastore;gcpvertexaimetadatastores
// +kubebuilder:subresource:status
// +kubebuilder:metadata:labels="cnrm.cloud.google.com/managed-by-kcc=true";"cnrm.cloud.google.com/system=true";"cnrm.cloud.google.com/tf2crd=true";"cnrm.cloud.google.com/stability-level=alpha"
// +kubebuilder:printcolumn:name="Age",JSONPath=".metadata.creationTimestamp",type="date"
// +kubebuilder:printcolumn:name="Ready",JSONPath=".status.conditions[?(@.type=='Ready')].status",type="string",description="When 'True', the most recent reconcile of the resource succeeded"
// +kubebuilder:printcolumn:name="Status",JSONPath=".status.conditions[?(@.type=='Ready')].reason",type="string",description="The reason for the value in 'Ready'"
// +kubebuilder:printcolumn:name="Status Age",JSONPath=".status.conditions[?(@.type=='Ready')].lastTransitionTime",type="date",description="The last transition time for the value in 'Status'"

// VertexAIMetadataStore is the Schema for the VertexAIMetadataStore API
// +k8s:openapi-gen=true
type VertexAIMetadataStore struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	// +required
	Spec   VertexAIMetadataStoreSpec   `json:"spec,omitempty"`
	Status VertexAIMetadataStoreStatus `json:"status,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// VertexAIMetadataStoreList contains a list of VertexAIMetadataStore
type VertexAIMetadataStoreList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []VertexAIMetadataStore `json:"items"`
}

func init() {
	SchemeBuilder.Register(&VertexAIMetadataStore{}, &VertexAIMetadataStoreList{})
}
