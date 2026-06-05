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

// +kcc:proto=google.cloud.enterpriseknowledgegraph.v1.BigQueryInputConfig
type BigQueryInputConfig struct {
	// Required. Format is `projects/*/datasets/*/tables/*`.
	// +required
	BigqueryTableRef *bigqueryv1beta1.BigQueryTableRef `json:"bigqueryTableRef"`

	// Required. Schema mapping file
	// +required
	GCSURI *string `json:"gcsURI"`
}

// +kcc:proto=google.cloud.enterpriseknowledgegraph.v1.InputConfig
type InputConfig struct {
	// Set of input BigQuery tables.
	// +required
	BigqueryInputConfigs []BigQueryInputConfig `json:"bigqueryInputConfigs,omitempty"`

	// Entity type
	// +optional
	EntityType *string `json:"entityType,omitempty"`

	// Optional. Provide the bigquery table containing the previous results if
	//  cluster ID stability is desired. Format is
	//  `projects/-*-/datasets/-*-/tables/-*`.
	// +optional
	PreviousResultBigqueryTableRef *bigqueryv1beta1.BigQueryTableRef `json:"previousResultBigqueryTableRef,omitempty"`
}

// +kcc:proto=google.cloud.enterpriseknowledgegraph.v1.OutputConfig
type OutputConfig struct {
	// Format is “projects/-*-/datasets/-*”.
	// +required
	BigqueryDatasetRef *refsv1beta1.BigQueryDatasetRef `json:"bigqueryDatasetRef"`
}

// EnterpriseKnowledgeGraphEntityReconciliationJobSpec defines the desired state of EnterpriseKnowledgeGraphEntityReconciliationJob
// +kcc:spec:proto=google.cloud.enterpriseknowledgegraph.v1.EntityReconciliationJob
type EnterpriseKnowledgeGraphEntityReconciliationJobSpec struct {
	// The project that this resource belongs to.
	ProjectRef *refsv1beta1.ProjectRef `json:"projectRef"`

	// The location of this resource.
	Location string `json:"location"`

	// The EnterpriseKnowledgeGraphEntityReconciliationJob name. If not given, the metadata.name will be used.
	ResourceID *string `json:"resourceID,omitempty"`

	// Required. Information about the input BigQuery tables.
	// +required
	InputConfig *InputConfig `json:"inputConfig"`

	// Required. The desired output location.
	// +required
	OutputConfig *OutputConfig `json:"outputConfig"`

	// Optional. Recon configs to adjust the clustering behavior.
	// +optional
	ReconConfig *ReconConfig `json:"reconConfig,omitempty"`
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
	// Output only. Resource name of the EntityReconciliationJob.
	// +optional
	Name *string `json:"name,omitempty"`

	// Output only. The detailed state of the job.
	// +optional
	State *string `json:"state,omitempty"`

	// Output only. Only populated when the job's state is JOB_STATE_FAILED or
	//  JOB_STATE_CANCELLED.
	// +optional
	Error *common.Status `json:"error,omitempty"`

	// Output only. Time when the EntityReconciliationJob was created.
	// +optional
	CreateTime *string `json:"createTime,omitempty"`

	// Output only. Time when the EntityReconciliationJob entered any of the
	//  following states: `JOB_STATE_SUCCEEDED`, `JOB_STATE_FAILED`,
	//  `JOB_STATE_CANCELLED`.
	// +optional
	EndTime *string `json:"endTime,omitempty"`

	// Output only. Time when the EntityReconciliationJob was most recently
	//  updated.
	// +optional
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
