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
	bigqueryv1beta1 "github.com/GoogleCloudPlatform/k8s-config-connector/apis/bigquery/v1beta1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/apis/common"
	refsv1beta1 "github.com/GoogleCloudPlatform/k8s-config-connector/apis/refs/v1beta1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/apis/k8s/v1alpha1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

var EnterpriseKnowledgeGraphEntityReconciliationJobGVK = GroupVersion.WithKind("EnterpriseKnowledgeGraphEntityReconciliationJob")

// EnterpriseKnowledgeGraphEntityReconciliationJobSpec defines the desired state of EnterpriseKnowledgeGraphEntityReconciliationJob
// +kcc:spec:proto=google.cloud.enterpriseknowledgegraph.v1.EntityReconciliationJob
type EnterpriseKnowledgeGraphEntityReconciliationJobSpec struct {
	// The project that this resource belongs to.
	ProjectRef *refsv1beta1.ProjectRef `json:"projectRef"`

	// The location of this resource.
	Location *string `json:"location"`

	// The EnterpriseKnowledgeGraphEntityReconciliationJob name. If not given, the metadata.name will be used.
	ResourceID *string `json:"resourceID,omitempty"`

	// Required. Information about the input BigQuery tables.
	// +kcc:proto:field=google.cloud.enterpriseknowledgegraph.v1.EntityReconciliationJob.input_config
	InputConfig *InputConfig `json:"inputConfig,omitempty"`

	// Required. The desired output location.
	// +kcc:proto:field=google.cloud.enterpriseknowledgegraph.v1.EntityReconciliationJob.output_config
	OutputConfig *OutputConfig `json:"outputConfig,omitempty"`

	// Optional. Recon configs to adjust the clustering behavior.
	// +kcc:proto:field=google.cloud.enterpriseknowledgegraph.v1.EntityReconciliationJob.recon_config
	ReconConfig *ReconConfig `json:"reconConfig,omitempty"`
}

// +kcc:proto=google.cloud.enterpriseknowledgegraph.v1.InputConfig
type InputConfig struct {
	// Set of input BigQuery tables.
	// +kcc:proto:field=google.cloud.enterpriseknowledgegraph.v1.InputConfig.bigquery_input_configs
	BigqueryInputConfigs []BigQueryInputConfig `json:"bigqueryInputConfigs,omitempty"`

	// Entity type
	// +kubebuilder:validation:Enum=ENTITY_TYPE_UNSPECIFIED;PEOPLE;ESTABLISHMENT;PROPERTY;PRODUCT;ORGANIZATION;LOCAL_BUSINESS;PERSON
	// +kcc:proto:field=google.cloud.enterpriseknowledgegraph.v1.InputConfig.entity_type
	EntityType *string `json:"entityType,omitempty"`

	// Optional. Provide the bigquery table containing the previous results if
	//  cluster ID stability is desired. Format is
	//  `projects/{project}/datasets/{dataset}/tables/{table}`.
	// +kcc:proto:field=google.cloud.enterpriseknowledgegraph.v1.InputConfig.previous_result_bigquery_table
	PreviousResultBigqueryTableRef *bigqueryv1beta1.BigQueryTableRef `json:"previousResultBigqueryTableRef,omitempty"`
}

// +kcc:proto=google.cloud.enterpriseknowledgegraph.v1.BigQueryInputConfig
type BigQueryInputConfig struct {
	// Required. Reference to BigQuery table. Format is `projects/{project}/datasets/{dataset}/tables/{table}`.
	// +kcc:proto:field=google.cloud.enterpriseknowledgegraph.v1.BigQueryInputConfig.bigquery_table
	BigqueryTableRef *bigqueryv1beta1.BigQueryTableRef `json:"bigqueryTableRef,omitempty"`

	// Required. Schema mapping file
	// +kcc:proto:field=google.cloud.enterpriseknowledgegraph.v1.BigQueryInputConfig.gcs_uri
	GCSURI *string `json:"gcsURI,omitempty"`
}

// +kcc:proto=google.cloud.enterpriseknowledgegraph.v1.OutputConfig
type OutputConfig struct {
	// Format is `projects/{project}/datasets/{dataset}`.
	// +kcc:proto:field=google.cloud.enterpriseknowledgegraph.v1.OutputConfig.bigquery_dataset
	BigqueryDatasetRef *bigqueryv1beta1.DatasetRef `json:"bigqueryDatasetRef,omitempty"`
}

// +kcc:proto=google.cloud.enterpriseknowledgegraph.v1.ReconConfig
type ReconConfig struct {
	// Configs for connected components.
	// +kcc:proto:field=google.cloud.enterpriseknowledgegraph.v1.ReconConfig.connected_components_config
	ConnectedComponentsConfig *ConnectedComponentsConfig `json:"connectedComponentsConfig,omitempty"`

	// Configs for affinity clustering.
	// +kcc:proto:field=google.cloud.enterpriseknowledgegraph.v1.ReconConfig.affinity_clustering_config
	AffinityClusteringConfig *AffinityClusteringConfig `json:"affinityClusteringConfig,omitempty"`

	// Extra options that affect entity clustering behavior.
	// +kcc:proto:field=google.cloud.enterpriseknowledgegraph.v1.ReconConfig.options
	Options *ReconConfig_Options `json:"options,omitempty"`

	// Model Configs
	// +kcc:proto:field=google.cloud.enterpriseknowledgegraph.v1.ReconConfig.model_config
	ModelConfig *ReconConfig_ModelConfig `json:"modelConfig,omitempty"`
}

// +kcc:proto=google.cloud.enterpriseknowledgegraph.v1.ConnectedComponentsConfig
type ConnectedComponentsConfig struct {
	// Threshold used for connected components. Default value is 0.85.
	// +kcc:proto:field=google.cloud.enterpriseknowledgegraph.v1.ConnectedComponentsConfig.weight_threshold
	WeightThreshold *float32 `json:"weightThreshold,omitempty"`
}

// +kcc:proto=google.cloud.enterpriseknowledgegraph.v1.AffinityClusteringConfig
type AffinityClusteringConfig struct {
	// Number of iterations to perform. Default value is 1.
	// +kcc:proto:field=google.cloud.enterpriseknowledgegraph.v1.AffinityClusteringConfig.compression_round_count
	CompressionRoundCount *int64 `json:"compressionRoundCount,omitempty"`
}

// +kcc:proto=google.cloud.enterpriseknowledgegraph.v1.ReconConfig.Options
type ReconConfig_Options struct {
	// If true, separate clusters by their geographic region (from geocoding).
	//  Uses the following entity features:
	//
	//  - schema.org/addressLocality
	//  - schema.org/addressRegion
	//  - schema.org/addressCountry
	//  Warning: processing will no longer be regionalized!
	// +kcc:proto:field=google.cloud.enterpriseknowledgegraph.v1.ReconConfig.Options.enable_geocoding_separation
	EnableGeocodingSeparation *bool `json:"enableGeocodingSeparation,omitempty"`
}

// +kcc:proto=google.cloud.enterpriseknowledgegraph.v1.ReconConfig.ModelConfig
type ReconConfig_ModelConfig struct {
	// Model name. Refer to external documentation for valid names.
	//  If unspecified, it defaults to the one mentioned in the documentation.
	// +kcc:proto:field=google.cloud.enterpriseknowledgegraph.v1.ReconConfig.ModelConfig.model_name
	ModelName *string `json:"modelName,omitempty"`

	// Model version tag. Refer to external documentation for valid tags.
	//  If unspecified, it defaults to the one mentioned in the documentation.
	// +kcc:proto:field=google.cloud.enterpriseknowledgegraph.v1.ReconConfig.ModelConfig.version_tag
	VersionTag *string `json:"versionTag,omitempty"`
}

// EnterpriseKnowledgeGraphEntityReconciliationJobStatus defines the config connector machine state of EnterpriseKnowledgeGraphEntityReconciliationJob
type EnterpriseKnowledgeGraphEntityReconciliationJobStatus struct {
	/* Conditions represent the latest available observations of the
	   object's current state. */
	Conditions []v1alpha1.Condition `json:"conditions,omitempty"`

	// ObservedGeneration is the generation of the resource that was most recently observed by the Config Connector controller. If this is equal to metadata.generation, then that means that the current reported status reflects the most recent desired state of the resource.
	ObservedGeneration *int64 `json:"observedGeneration,omitempty"`

	// A unique specifier for the EnterpriseKnowledgeGraphEntityReconciliationJob resource in GCP.
	ExternalRef *string `json:"externalRef,omitempty"`

	// ObservedState is the state of the resource as most recently observed in GCP.
	ObservedState *EnterpriseKnowledgeGraphEntityReconciliationJobObservedState `json:"observedState,omitempty"`
}

// EnterpriseKnowledgeGraphEntityReconciliationJobObservedState is the state of the EnterpriseKnowledgeGraphEntityReconciliationJob resource as most recently observed in GCP.
// +kcc:observedstate:proto=google.cloud.enterpriseknowledgegraph.v1.EntityReconciliationJob
type EnterpriseKnowledgeGraphEntityReconciliationJobObservedState struct {
	// Output only. The detailed state of the job.
	// +kubebuilder:validation:Enum=JOB_STATE_UNSPECIFIED;JOB_STATE_PENDING;JOB_STATE_RUNNING;JOB_STATE_SUCCEEDED;JOB_STATE_FAILED;JOB_STATE_CANCELLED;JOB_STATE_KNOWLEDGE_EXTRACTION;JOB_STATE_RECON_PREPROCESSING;JOB_STATE_CLUSTERING;JOB_STATE_EXPORTING_CLUSTERS
	// +kcc:proto:field=google.cloud.enterpriseknowledgegraph.v1.EntityReconciliationJob.state
	State *string `json:"state,omitempty"`

	// Output only. Only populated when the job's state is JOB_STATE_FAILED or
	//  JOB_STATE_CANCELLED.
	// +kcc:proto:field=google.cloud.enterpriseknowledgegraph.v1.EntityReconciliationJob.error
	Error *common.Status `json:"error,omitempty"`

	// Output only. Time when the EntityReconciliationJob was created.
	// +kcc:proto:field=google.cloud.enterpriseknowledgegraph.v1.EntityReconciliationJob.create_time
	CreateTime *string `json:"createTime,omitempty"`

	// Output only. Time when the EntityReconciliationJob entered any of the
	//  following states: `JOB_STATE_SUCCEEDED`, `JOB_STATE_FAILED`,
	//  `JOB_STATE_CANCELLED`.
	// +kcc:proto:field=google.cloud.enterpriseknowledgegraph.v1.EntityReconciliationJob.end_time
	EndTime *string `json:"endTime,omitempty"`

	// Output only. Time when the EntityReconciliationJob was most recently
	//  updated.
	// +kcc:proto:field=google.cloud.enterpriseknowledgegraph.v1.EntityReconciliationJob.update_time
	UpdateTime *string `json:"updateTime,omitempty"`
}

// +genclient
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:resource:categories=gcp,shortName=gcpenterpriseknowledgegraphentityreconciliationjob;gcpenterpriseknowledgegraphentityreconciliationjobs
// +kubebuilder:subresource:status
// +kubebuilder:metadata:labels="cnrm.cloud.google.com/managed-by-kcc=true"
// +kubebuilder:metadata:labels="cnrm.cloud.google.com/system=true"
// +kubebuilder:metadata:labels="cnrm.cloud.google.com/stability-level=alpha"
// +kubebuilder:printcolumn:name="Age",JSONPath=".metadata.creationTimestamp",type="date"
// +kubebuilder:printcolumn:name="Ready",JSONPath=".status.conditions[?(@.type=='Ready')].status",type="string",description="When 'True', the most recent reconcile of the resource succeeded"
// +kubebuilder:printcolumn:name="Status",JSONPath=".status.conditions[?(@.type=='Ready')].reason",type="string",description="The reason for the value in 'Ready'"
// +kubebuilder:printcolumn:name="Status Age",JSONPath=".status.conditions[?(@.type=='Ready')].lastTransitionTime",type="date",description="The last transition time for the value in 'Status'"

// EnterpriseKnowledgeGraphEntityReconciliationJob is the Schema for the EnterpriseKnowledgeGraphEntityReconciliationJob API
// +k8s:openapi-gen=true
type EnterpriseKnowledgeGraphEntityReconciliationJob struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	// +required
	Spec   EnterpriseKnowledgeGraphEntityReconciliationJobSpec   `json:"spec,omitempty"`
	Status EnterpriseKnowledgeGraphEntityReconciliationJobStatus `json:"status,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// EnterpriseKnowledgeGraphEntityReconciliationJobList contains a list of EnterpriseKnowledgeGraphEntityReconciliationJob
type EnterpriseKnowledgeGraphEntityReconciliationJobList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []EnterpriseKnowledgeGraphEntityReconciliationJob `json:"items"`
}

func init() {
	SchemeBuilder.Register(&EnterpriseKnowledgeGraphEntityReconciliationJob{}, &EnterpriseKnowledgeGraphEntityReconciliationJobList{})
}
