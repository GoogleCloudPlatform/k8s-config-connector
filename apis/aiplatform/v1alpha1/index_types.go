// Copyright 2026 Google LLC
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
	apiextensionsv1 "k8s.io/apiextensions-apiserver/pkg/apis/apiextensions/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

var VertexAIIndexGVK = GroupVersion.WithKind("VertexAIIndex")

// VertexAIIndexSpec defines the desired state of VertexAIIndex
// +kcc:spec:proto=google.cloud.aiplatform.v1.Index
type VertexAIIndexSpec struct {
	// The project that this resource belongs to.
	// +required
	ProjectRef *refsv1beta1.ProjectRef `json:"projectRef"`

	// The location of this resource.
	// +required
	Location string `json:"location"`

	// The VertexAIIndex name. If not given, the metadata.name will be used.
	// +kubebuilder:validation:Optional
	ResourceID *string `json:"resourceID,omitempty"`

	// Required. The display name of the Index.
	//  The name can be up to 128 characters long and can consist of any UTF-8
	//  characters.
	// +kcc:proto:field=google.cloud.aiplatform.v1.Index.display_name
	// +kubebuilder:validation:Required
	DisplayName *string `json:"displayName,omitempty"`

	// The description of the Index.
	// +kcc:proto:field=google.cloud.aiplatform.v1.Index.description
	// +kubebuilder:validation:Optional
	Description *string `json:"description,omitempty"`

	// Immutable. Points to a YAML file stored on Google Cloud Storage describing
	//  additional information about the Index, that is specific to it. Unset if
	//  the Index does not have any additional information. The schema is defined
	//  as an OpenAPI 3.0.2 [Schema
	//  Object](https://github.com/OAI/OpenAPI-Specification/blob/main/versions/3.0.2.md#schemaObject).
	//  Note: The URI given on output will be immutable and probably different,
	//  including the URI scheme, than the one given on input. The output URI will
	//  point to a location where the user only has a read access.
	// +kcc:proto:field=google.cloud.aiplatform.v1.Index.metadata_schema_uri
	// +kubebuilder:validation:Optional
	MetadataSchemaURI *string `json:"metadataSchemaURI,omitempty"`

	// An additional information about the Index; the schema of the metadata can
	//  be found in
	//  [metadata_schema][google.cloud.aiplatform.v1.Index.metadata_schema_uri].
	// +kcc:proto:field=google.cloud.aiplatform.v1.Index.metadata
	// +kubebuilder:validation:Optional
	Metadata *apiextensionsv1.JSON `json:"metadata,omitempty"`

	// Used to perform consistent read-modify-write updates. If not set, a blind
	//  "overwrite" update happens.
	// +kcc:proto:field=google.cloud.aiplatform.v1.Index.etag
	// +kubebuilder:validation:Optional
	Etag *string `json:"etag,omitempty"`

	// The labels with user-defined metadata to organize your Indexes.
	//
	//  Label keys and values can be no longer than 64 characters
	//  (Unicode codepoints), can only contain lowercase letters, numeric
	//  characters, underscores and dashes. International characters are allowed.
	//
	//  See https://goo.gl/xmQnxf for more information and examples of labels.
	// +kcc:proto:field=google.cloud.aiplatform.v1.Index.labels
	// +kubebuilder:validation:Optional
	Labels map[string]string `json:"labels,omitempty"`

	// Immutable. The update method to use with this Index. If not set,
	//  BATCH_UPDATE will be used by default.
	// +kcc:proto:field=google.cloud.aiplatform.v1.Index.index_update_method
	// +kubebuilder:validation:Optional
	IndexUpdateMethod *string `json:"indexUpdateMethod,omitempty"`

	// Immutable. Customer-managed encryption key spec for an Index. If set, this
	//  Index and all sub-resources of this Index will be secured by this key.
	// +kcc:proto:field=google.cloud.aiplatform.v1.Index.encryption_spec
	// +kubebuilder:validation:Optional
	EncryptionSpec *VertexAIEncryptionSpec `json:"encryptionSpec,omitempty"`
}

// VertexAIIndexStatus defines the config connector machine state of VertexAIIndex
type VertexAIIndexStatus struct {
	/* Conditions represent the latest available observations of the
	   object's current state. */
	// +kubebuilder:validation:Optional
	Conditions []v1alpha1.Condition `json:"conditions,omitempty"`

	// ObservedGeneration is the generation of the resource that was most recently observed by the Config Connector controller. If this is equal to metadata.generation, then that means that the current reported status reflects the most recent desired state of the resource.
	// +kubebuilder:validation:Optional
	ObservedGeneration *int64 `json:"observedGeneration,omitempty"`

	// A unique specifier for the VertexAIIndex resource in GCP.
	// +kubebuilder:validation:Optional
	ExternalRef *string `json:"externalRef,omitempty"`

	// ObservedState is the state of the resource as most recently observed in GCP.
	// +kubebuilder:validation:Optional
	ObservedState *VertexAIIndexObservedState `json:"observedState,omitempty"`
}

// VertexAIIndexObservedState is the state of the VertexAIIndex resource as most recently observed in GCP.
// +kcc:observedstate:proto=google.cloud.aiplatform.v1.Index
type VertexAIIndexObservedState struct {

	// Output only. The resource name of the Index.
	// +kcc:proto:field=google.cloud.aiplatform.v1.Index.name
	// +kubebuilder:validation:Optional
	Name *string `json:"name,omitempty"`

	// Output only. The pointers to DeployedIndexes created from this Index.
	//  An Index can be only deleted if all its DeployedIndexes had been undeployed
	//  first.
	// +kcc:proto:field=google.cloud.aiplatform.v1.Index.deployed_indexes
	// +kubebuilder:validation:Optional
	DeployedIndexes []VertexAIDeployedIndexRefObservedState `json:"deployedIndexes,omitempty"`

	// Output only. Timestamp when this Index was created.
	// +kcc:proto:field=google.cloud.aiplatform.v1.Index.create_time
	// +kubebuilder:validation:Optional
	CreateTime *string `json:"createTime,omitempty"`

	// Output only. Timestamp when this Index was most recently updated.
	//  This also includes any update to the contents of the Index.
	//  Note that Operations working on this Index may have their
	//  [Operations.metadata.generic_metadata.update_time]
	//  [google.cloud.aiplatform.v1.GenericOperationMetadata.update_time] a little
	//  after the value of this timestamp, yet that does not mean their results are
	//  not already reflected in the Index. Result of any successfully completed
	//  Operation on the Index is reflected in it.
	// +kcc:proto:field=google.cloud.aiplatform.v1.Index.update_time
	// +kubebuilder:validation:Optional
	UpdateTime *string `json:"updateTime,omitempty"`

	// Output only. Stats of the index resource.
	// +kcc:proto:field=google.cloud.aiplatform.v1.Index.index_stats
	// +kubebuilder:validation:Optional
	IndexStats *VertexAIIndexStatsObservedState `json:"indexStats,omitempty"`

	// Output only. Reserved for future use.
	// +kcc:proto:field=google.cloud.aiplatform.v1.Index.satisfies_pzs
	// +kubebuilder:validation:Optional
	SatisfiesPzs *bool `json:"satisfiesPzs,omitempty"`

	// Output only. Reserved for future use.
	// +kcc:proto:field=google.cloud.aiplatform.v1.Index.satisfies_pzi
	// +kubebuilder:validation:Optional
	SatisfiesPzi *bool `json:"satisfiesPzi,omitempty"`
}

type VertexAIDeployedIndexRef struct {

	// Immutable. A resource name of the IndexEndpoint.
	// +kcc:proto:field=google.cloud.aiplatform.v1.DeployedIndexRef.index_endpoint
	// +kubebuilder:validation:Optional
	IndexEndpoint *string `json:"indexEndpoint,omitempty"`

	// Immutable. The ID of the DeployedIndex in the above IndexEndpoint.
	// +kcc:proto:field=google.cloud.aiplatform.v1.DeployedIndexRef.deployed_index_id
	// +kubebuilder:validation:Optional
	DeployedIndexID *string `json:"deployedIndexID,omitempty"`
}

type VertexAIEncryptionSpec struct {

	// Required. The Cloud KMS resource identifier of the customer managed
	//  encryption key used to protect a resource. Has the form:
	//  `projects/my-project/locations/my-region/keyRings/my-kr/cryptoKeys/my-key`.
	//  The key needs to be in the same region as where the compute resource is
	//  created.
	// +kcc:proto:field=google.cloud.aiplatform.v1.EncryptionSpec.kms_key_name
	// +kubebuilder:validation:Required
	KMSKeyName *string `json:"kmsKeyName,omitempty"`
}

type VertexAIDeployedIndexRefObservedState struct {

	// Output only. The display name of the DeployedIndex.
	// +kcc:proto:field=google.cloud.aiplatform.v1.DeployedIndexRef.display_name
	// +kubebuilder:validation:Optional
	DisplayName *string `json:"displayName,omitempty"`
}

type VertexAIIndexStatsObservedState struct {

	// Output only. The number of dense vectors in the Index.
	// +kcc:proto:field=google.cloud.aiplatform.v1.IndexStats.vectors_count
	// +kubebuilder:validation:Optional
	VectorsCount *int64 `json:"vectorsCount,omitempty"`

	// Output only. The number of sparse vectors in the Index.
	// +kcc:proto:field=google.cloud.aiplatform.v1.IndexStats.sparse_vectors_count
	// +kubebuilder:validation:Optional
	SparseVectorsCount *int64 `json:"sparseVectorsCount,omitempty"`

	// Output only. The number of shards in the Index.
	// +kcc:proto:field=google.cloud.aiplatform.v1.IndexStats.shards_count
	// +kubebuilder:validation:Optional
	ShardsCount *int32 `json:"shardsCount,omitempty"`
}

// +genclient
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:resource:categories=gcp,shortName=gcpvertexaiindex;gcpvertexaiindexs
// +kubebuilder:subresource:status
// +kubebuilder:metadata:labels="cnrm.cloud.google.com/managed-by-kcc=true"
// +kubebuilder:metadata:labels="cnrm.cloud.google.com/system=true"
// +kubebuilder:printcolumn:name="Age",JSONPath=".metadata.creationTimestamp",type="date"
// +kubebuilder:printcolumn:name="Ready",JSONPath=".status.conditions[?(@.type=='Ready')].status",type="string",description="When 'True', the most recent reconcile of the resource succeeded"
// +kubebuilder:printcolumn:name="Status",JSONPath=".status.conditions[?(@.type=='Ready')].reason",type="string",description="The reason for the value in 'Ready'"
// +kubebuilder:printcolumn:name="Status Age",JSONPath=".status.conditions[?(@.type=='Ready')].lastTransitionTime",type="date",description="The last transition time for the value in 'Status'"

// VertexAIIndex is the Schema for the VertexAIIndex API
// +k8s:openapi-gen=true
type VertexAIIndex struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	// +required
	Spec VertexAIIndexSpec `json:"spec,omitempty"`
	// +kubebuilder:validation:Optional
	Status VertexAIIndexStatus `json:"status,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// VertexAIIndexList contains a list of VertexAIIndex
type VertexAIIndexList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// +kubebuilder:validation:Optional
	Items []VertexAIIndex `json:"items"`
}

func init() {
	SchemeBuilder.Register(&VertexAIIndex{}, &VertexAIIndexList{})
}
