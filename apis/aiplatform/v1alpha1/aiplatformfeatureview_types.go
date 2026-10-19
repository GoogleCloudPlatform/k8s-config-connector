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

// API sources for AIPlatformFeatureView, recorded by generate-types:
// +kcc:source:proto=https://github.com/googleapis/googleapis/blob/e7513c828b8db536fa12fc5493d541f323aedf90/google/cloud/aiplatform/v1/feature_view.proto
// +kcc:source:service-docs=https://cloud.google.com/vertex-ai/
// +kcc:guess=source-link reason=verify-resource-docs-link
// +kcc:source:resource-docs=https://docs.cloud.google.com/vertex-ai/docs/reference/rest/v1/projects.locations.featureOnlineStores.featureViews

package v1alpha1

import (
	refsv1beta1 "github.com/GoogleCloudPlatform/k8s-config-connector/apis/refs/v1beta1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/apis/k8s/v1alpha1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

var AIPlatformFeatureViewGVK = GroupVersion.WithKind("AIPlatformFeatureView")

// AIPlatformFeatureViewSpec defines the desired state of AIPlatformFeatureView
// +kcc:spec:proto=google.cloud.aiplatform.v1.FeatureView
// +kcc:required-from-proto
type AIPlatformFeatureViewSpec struct {
	// The project that this resource belongs to.
	ProjectRef *refsv1beta1.ProjectRef `json:"projectRef"`

	// The location of this resource.
	Location *string `json:"location"`

	// The AIPlatformFeatureView name. If not given, the metadata.name will be used.
	ResourceID *string `json:"resourceID,omitempty"`
	// Optional. Configures how data is supposed to be extracted from a BigQuery
	//  source to be loaded onto the FeatureOnlineStore.
	// +kcc:proto:field=google.cloud.aiplatform.v1.FeatureView.big_query_source
	BigQuerySource *FeatureView_BigQuerySource `json:"bigQuerySource,omitempty"`

	// Optional. Configures the features from a Feature Registry source that
	//  need to be loaded onto the FeatureOnlineStore.
	// +kcc:proto:field=google.cloud.aiplatform.v1.FeatureView.feature_registry_source
	FeatureRegistrySource *FeatureView_FeatureRegistrySource `json:"featureRegistrySource,omitempty"`

	// Optional. The Vertex RAG Source that the FeatureView is linked to.
	// +kcc:proto:field=google.cloud.aiplatform.v1.FeatureView.vertex_rag_source
	VertexRagSource *FeatureView_VertexRagSource `json:"vertexRagSource,omitempty"`

	// Optional. Used to perform consistent read-modify-write updates. If not set,
	//  a blind "overwrite" update happens.
	// +kcc:proto:field=google.cloud.aiplatform.v1.FeatureView.etag
	Etag *string `json:"etag,omitempty"`

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
	// +kcc:proto:field=google.cloud.aiplatform.v1.FeatureView.labels
	Labels map[string]string `json:"labels,omitempty"`

	// Configures when data is to be synced/updated for this FeatureView. At the
	//  end of the sync the latest featureValues for each entityId of this
	//  FeatureView are made ready for online serving.
	// +kcc:proto:field=google.cloud.aiplatform.v1.FeatureView.sync_config
	SyncConfig *FeatureView_SyncConfig `json:"syncConfig,omitempty"`

	// Optional. Configuration for index preparation for vector search. It
	//  contains the required configurations to create an index from source data,
	//  so that approximate nearest neighbor (a.k.a ANN) algorithms search can be
	//  performed during online serving.
	// +kcc:proto:field=google.cloud.aiplatform.v1.FeatureView.index_config
	IndexConfig *FeatureView_IndexConfig `json:"indexConfig,omitempty"`

	// Optional. Configuration for FeatureView created under Optimized
	//  FeatureOnlineStore.
	// +kcc:proto:field=google.cloud.aiplatform.v1.FeatureView.optimized_config
	OptimizedConfig *FeatureView_OptimizedConfig `json:"optimizedConfig,omitempty"`

	// Optional. Service agent type used during data sync. By default, the Vertex
	//  AI Service Agent is used. When using an IAM Policy to isolate this
	//  FeatureView within a project, a separate service account should be
	//  provisioned by setting this field to `SERVICE_AGENT_TYPE_FEATURE_VIEW`.
	//  This will generate a separate service account to access the BigQuery source
	//  table.
	// +kcc:proto:field=google.cloud.aiplatform.v1.FeatureView.service_agent_type
	ServiceAgentType *string `json:"serviceAgentType,omitempty"`
}

// AIPlatformFeatureViewStatus defines the config connector machine state of AIPlatformFeatureView
type AIPlatformFeatureViewStatus struct {
	/* Conditions represent the latest available observations of the
	   object's current state. */
	Conditions []v1alpha1.Condition `json:"conditions,omitempty"`

	// ObservedGeneration is the generation of the resource that was most recently observed by the Config Connector controller. If this is equal to metadata.generation, then that means that the current reported status reflects the most recent desired state of the resource.
	ObservedGeneration *int64 `json:"observedGeneration,omitempty"`

	// A unique specifier for the AIPlatformFeatureView resource in GCP.
	ExternalRef *string `json:"externalRef,omitempty"`

	// ObservedState is the state of the resource as most recently observed in GCP.
	ObservedState *AIPlatformFeatureViewObservedState `json:"observedState,omitempty"`
}

// AIPlatformFeatureViewObservedState is the state of the AIPlatformFeatureView resource as most recently observed in GCP.
// +kcc:observedstate:proto=google.cloud.aiplatform.v1.FeatureView
type AIPlatformFeatureViewObservedState struct {
	// Output only. Timestamp when this FeatureView was created.
	// +kcc:proto:field=google.cloud.aiplatform.v1.FeatureView.create_time
	CreateTime *string `json:"createTime,omitempty"`

	// Output only. Timestamp when this FeatureView was last updated.
	// +kcc:proto:field=google.cloud.aiplatform.v1.FeatureView.update_time
	UpdateTime *string `json:"updateTime,omitempty"`

	// Output only. A Service Account unique to this FeatureView. The role
	//  bigquery.dataViewer should be granted to this service account to allow
	//  Vertex AI Feature Store to sync data to the online store.
	// +kcc:proto:field=google.cloud.aiplatform.v1.FeatureView.service_account_email
	ServiceAccountEmail *string `json:"serviceAccountEmail,omitempty"`

	// Output only. Reserved for future use.
	// +kcc:proto:field=google.cloud.aiplatform.v1.FeatureView.satisfies_pzs
	SatisfiesPzs *bool `json:"satisfiesPzs,omitempty"`

	// Output only. Reserved for future use.
	// +kcc:proto:field=google.cloud.aiplatform.v1.FeatureView.satisfies_pzi
	SatisfiesPzi *bool `json:"satisfiesPzi,omitempty"`

	// Metadata containing information about the Cloud Bigtable.
	// +kcc:proto:field=google.cloud.aiplatform.v1.FeatureView.bigtable_metadata
	BigtableMetadata *FeatureView_BigtableMetadataObservedState `json:"bigtableMetadata,omitempty"`
}

// +genclient
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:resource:categories=gcp,shortName=gcpaiplatformfeatureview;gcpaiplatformfeatureviews
// +kubebuilder:subresource:status
// +kubebuilder:metadata:labels="cnrm.cloud.google.com/managed-by-kcc=true"
// +kubebuilder:metadata:labels="cnrm.cloud.google.com/system=true"
// +kubebuilder:printcolumn:name="Age",JSONPath=".metadata.creationTimestamp",type="date"
// +kubebuilder:printcolumn:name="Ready",JSONPath=".status.conditions[?(@.type=='Ready')].status",type="string",description="When 'True', the most recent reconcile of the resource succeeded"
// +kubebuilder:printcolumn:name="Status",JSONPath=".status.conditions[?(@.type=='Ready')].reason",type="string",description="The reason for the value in 'Ready'"
// +kubebuilder:printcolumn:name="Status Age",JSONPath=".status.conditions[?(@.type=='Ready')].lastTransitionTime",type="date",description="The last transition time for the value in 'Status'"

// AIPlatformFeatureView is the Schema for the AIPlatformFeatureView API
// +k8s:openapi-gen=true
type AIPlatformFeatureView struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	// +required
	Spec   AIPlatformFeatureViewSpec   `json:"spec,omitempty"`
	Status AIPlatformFeatureViewStatus `json:"status,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// AIPlatformFeatureViewList contains a list of AIPlatformFeatureView
type AIPlatformFeatureViewList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []AIPlatformFeatureView `json:"items"`
}

func init() {
	SchemeBuilder.Register(&AIPlatformFeatureView{}, &AIPlatformFeatureViewList{})
}
