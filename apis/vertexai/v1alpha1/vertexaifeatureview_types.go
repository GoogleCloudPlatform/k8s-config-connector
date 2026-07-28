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
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

var VertexAIFeatureViewGVK = GroupVersion.WithKind("VertexAIFeatureView")

// VertexAIFeatureViewSpec defines the desired state of VertexAIFeatureView
type VertexAIFeatureViewSpec struct {
	// The project that this resource belongs to.
	// +required
	ProjectRef *refsv1beta1.ProjectRef `json:"projectRef"`

	// The location of this resource.
	// +required
	Location *string `json:"location"`

	// The parent FeatureOnlineStore of this resource.
	// +required
	FeatureOnlineStoreRef *VertexAIFeatureOnlineStoreRef `json:"featureOnlineStoreRef"`

	// The VertexAIFeatureView name. If not given, the metadata.name will be used.
	ResourceID *string `json:"resourceID,omitempty"`

	// Optional. Configures how data is supposed to be extracted from a BigQuery
	//  source to be loaded onto the FeatureOnlineStore.
	BigQuerySource *FeatureView_BigQuerySource `json:"bigQuerySource,omitempty"`

	// Optional. Configures the features from a Feature Registry source that
	//  need to be loaded onto the FeatureOnlineStore.
	FeatureRegistrySource *FeatureView_FeatureRegistrySource `json:"featureRegistrySource,omitempty"`

	// Optional. The Vertex RAG Source that the FeatureView is linked to.
	VertexRagSource *FeatureView_VertexRagSource `json:"vertexRagSource,omitempty"`

	// Optional. The labels with user-defined metadata to organize your
	//  FeatureViews.
	//
	//  Label keys and values can be no longer than 64 characters
	//  (Unicode codepoints), can only contain lowercase letters, numeric
	//  characters, underscores and dashes. International characters are allowed.
	//
	//  See https://goo.gl/xmQnxf for more information on and examples of labels.
	//  No more than 64 user labels can be associated with one
	//  FeatureOnlineStore(System labels are excluded)." System reserved label keys
	//  are prefixed with "aiplatform.googleapis.com/" and are immutable.
	Labels map[string]string `json:"labels,omitempty"`

	// Configures when data is to be synced/updated for this FeatureView. At the
	//  end of the sync the latest featureValues for each entityId of this
	//  FeatureView are made ready for online serving.
	SyncConfig *FeatureView_SyncConfig `json:"syncConfig,omitempty"`

	// Optional. Deprecated: please use
	//  [FeatureView.index_config][google.cloud.aiplatform.v1beta1.FeatureView.index_config]
	//  instead.
	VectorSearchConfig *FeatureView_VectorSearchConfig `json:"vectorSearchConfig,omitempty"`

	// Optional. Configuration for index preparation for vector search. It
	//  contains the required configurations to create an index from source data,
	//  so that approximate nearest neighbor (a.k.a. an ANN) algorithms search can be
	//  performed during online serving.
	IndexConfig *FeatureView_IndexConfig `json:"indexConfig,omitempty"`

	// Optional. Configuration for FeatureView created under Optimized
	//  FeatureOnlineStore.
	OptimizedConfig *FeatureView_OptimizedConfig `json:"optimizedConfig,omitempty"`

	// Optional. Service agent type used during data sync. By default, the Vertex
	//  AI Service Agent is used. When using an IAM Policy to isolate this
	//  FeatureView within a project, a separate service account should be
	//  provisioned by setting this field to `SERVICE_AGENT_TYPE_FEATURE_VIEW`.
	//  This will generate a separate service account to access the BigQuery source
	//  table.
	ServiceAgentType *string `json:"serviceAgentType,omitempty"`
}

// VertexAIFeatureViewStatus defines the config connector machine state of VertexAIFeatureView
type VertexAIFeatureViewStatus struct {
	/* Conditions represent the latest available observations of the
	   object's current state. */
	Conditions []v1alpha1.Condition `json:"conditions,omitempty"`

	// ObservedGeneration is the generation of the resource that was most recently observed by the Config Connector controller. If this is equal to metadata.generation, then that means that the current reported status reflects the most recent desired state of the resource.
	ObservedGeneration *int64 `json:"observedGeneration,omitempty"`

	// A unique specifier for the VertexAIFeatureView resource in GCP.
	ExternalRef *string `json:"externalRef,omitempty"`

	// ObservedState is the state of the resource as most recently observed in GCP.
	ObservedState *VertexAIFeatureViewObservedState `json:"observedState,omitempty"`
}

// VertexAIFeatureViewObservedState is the state of the VertexAIFeatureView resource as most recently observed in GCP.
type VertexAIFeatureViewObservedState struct {
	// Output only. Timestamp when this FeatureView was created.
	CreateTime *string `json:"createTime,omitempty"`

	// Output only. Timestamp when this FeatureView was last updated.
	UpdateTime *string `json:"updateTime,omitempty"`

	// Output only. A Service Account unique to this FeatureView. The role
	//  bigquery.dataViewer should be granted to this service account to allow
	//  Vertex AI Feature Store to sync data to the online store.
	ServiceAccountEmail *string `json:"serviceAccountEmail,omitempty"`

	// Output only. Reserved for future use.
	SatisfiesPzs *bool `json:"satisfiesPzs,omitempty"`

	// Output only. Reserved for future use.
	SatisfiesPzi *bool `json:"satisfiesPzi,omitempty"`

	// Metadata containing information about the Cloud Bigtable.
	BigtableMetadata *FeatureView_BigtableMetadataObservedState `json:"bigtableMetadata,omitempty"`
}

// +genclient
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:resource:categories=gcp,shortName=gcpvertexaifeatureview;gcpvertexaifeatureviews
// +kubebuilder:subresource:status
// +kubebuilder:metadata:labels="cnrm.cloud.google.com/managed-by-kcc=true"
// +kubebuilder:metadata:labels="cnrm.cloud.google.com/system=true"
// +kubebuilder:metadata:labels="cnrm.cloud.google.com/stability-level=alpha"
// +kubebuilder:printcolumn:name="Age",JSONPath=".metadata.creationTimestamp",type="date"
// +kubebuilder:printcolumn:name="Ready",JSONPath=".status.conditions[?(@.type=='Ready')].status",type="string",description="When 'True', the most recent reconcile of the resource succeeded"
// +kubebuilder:printcolumn:name="Status",JSONPath=".status.conditions[?(@.type=='Ready')].reason",type="string",description="The reason for the value in 'Ready'"
// +kubebuilder:printcolumn:name="Status Age",JSONPath=".status.conditions[?(@.type=='Ready')].lastTransitionTime",type="date",description="The last transition time for the value in 'Status'"

// VertexAIFeatureView is the Schema for the VertexAIFeatureView API
// +k8s:openapi-gen=true
type VertexAIFeatureView struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	// +required
	Spec   VertexAIFeatureViewSpec   `json:"spec,omitempty"`
	Status VertexAIFeatureViewStatus `json:"status,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// VertexAIFeatureViewList contains a list of VertexAIFeatureView
type VertexAIFeatureViewList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []VertexAIFeatureView `json:"items"`
}

func init() {
	SchemeBuilder.Register(&VertexAIFeatureView{}, &VertexAIFeatureViewList{})
}
