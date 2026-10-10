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

// API sources for CESEvaluationDataset, recorded by generate-types:
// +kcc:source:proto=https://github.com/googleapis/googleapis/blob/1765b559c42386788ff0c6412491277b4791107a/google/cloud/ces/v1beta/evaluation.proto
// +kcc:source:service-docs=https://docs.cloud.google.com/customer-engagement-ai/conversational-agents/ps
// +kcc:guess=source-link reason=verify-resource-docs-link
// +kcc:source:resource-docs=https://docs.cloud.google.com/customer-engagement-ai/conversational-agents/ps/docs/reference/rest/v1beta/projects.locations.apps.evaluationDatasets

package v1alpha1

import (
	refsv1beta1 "github.com/GoogleCloudPlatform/k8s-config-connector/apis/refs/v1beta1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/apis/k8s/v1alpha1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

var CESEvaluationDatasetGVK = GroupVersion.WithKind("CESEvaluationDataset")

// CESEvaluationDatasetSpec defines the desired state of CESEvaluationDataset
// +kcc:spec:proto=google.cloud.ces.v1beta.EvaluationDataset
// +kcc:required-from-proto
type CESEvaluationDatasetSpec struct {
	// The project that this resource belongs to.
	ProjectRef *refsv1beta1.ProjectRef `json:"projectRef"`

	// The location of this resource.
	Location string `json:"location"`

	// The CESEvaluationDataset name. If not given, the metadata.name will be used.
	ResourceID *string `json:"resourceID,omitempty"`
	// Required. User-defined display name of the evaluation dataset. Unique
	//  within an App.
	// +kcc:proto:field=google.cloud.ces.v1beta.EvaluationDataset.display_name
	// +required
	DisplayName *string `json:"displayName,omitempty"`

	// Optional. Evaluations that are included in this dataset.
	// +kcc:guess=possible-reference target=CESEvaluation
	// +kcc:proto:field=google.cloud.ces.v1beta.EvaluationDataset.evaluations
	Evaluations []string `json:"evaluations,omitempty"`
}

// CESEvaluationDatasetStatus defines the config connector machine state of CESEvaluationDataset
type CESEvaluationDatasetStatus struct {
	/* Conditions represent the latest available observations of the
	   object's current state. */
	Conditions []v1alpha1.Condition `json:"conditions,omitempty"`

	// ObservedGeneration is the generation of the resource that was most recently observed by the Config Connector controller. If this is equal to metadata.generation, then that means that the current reported status reflects the most recent desired state of the resource.
	ObservedGeneration *int64 `json:"observedGeneration,omitempty"`

	// A unique specifier for the CESEvaluationDataset resource in GCP.
	ExternalRef *string `json:"externalRef,omitempty"`

	// ObservedState is the state of the resource as most recently observed in GCP.
	ObservedState *CESEvaluationDatasetObservedState `json:"observedState,omitempty"`
}

// CESEvaluationDatasetObservedState is the state of the CESEvaluationDataset resource as most recently observed in GCP.
// +kcc:observedstate:proto=google.cloud.ces.v1beta.EvaluationDataset
type CESEvaluationDatasetObservedState struct {
	// Output only. Timestamp when the evaluation dataset was created.
	// +kcc:proto:field=google.cloud.ces.v1beta.EvaluationDataset.create_time
	CreateTime *string `json:"createTime,omitempty"`

	// Output only. Timestamp when the evaluation dataset was last updated.
	// +kcc:proto:field=google.cloud.ces.v1beta.EvaluationDataset.update_time
	UpdateTime *string `json:"updateTime,omitempty"`

	// Output only. Etag used to ensure the object hasn't changed during a
	//  read-modify-write operation. If the etag is empty, the update will
	//  overwrite any concurrent changes.
	// +kcc:proto:field=google.cloud.ces.v1beta.EvaluationDataset.etag
	Etag *string `json:"etag,omitempty"`

	// Output only. The user who created the evaluation dataset.
	// +kcc:proto:field=google.cloud.ces.v1beta.EvaluationDataset.created_by
	CreatedBy *string `json:"createdBy,omitempty"`

	// Output only. The user who last updated the evaluation dataset.
	// +kcc:proto:field=google.cloud.ces.v1beta.EvaluationDataset.last_updated_by
	LastUpdatedBy *string `json:"lastUpdatedBy,omitempty"`

	// Output only. The aggregated metrics for this evaluation dataset across all
	//  runs.
	// +kcc:proto:field=google.cloud.ces.v1beta.EvaluationDataset.aggregated_metrics
	AggregatedMetrics *AggregatedMetricsObservedState `json:"aggregatedMetrics,omitempty"`
}

// +genclient
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:resource:categories=gcp,shortName=gcpcesevaluationdataset;gcpcesevaluationdatasets
// +kubebuilder:subresource:status
// +kubebuilder:metadata:labels="cnrm.cloud.google.com/managed-by-kcc=true"
// +kubebuilder:metadata:labels="cnrm.cloud.google.com/system=true"
// +kubebuilder:printcolumn:name="Age",JSONPath=".metadata.creationTimestamp",type="date"
// +kubebuilder:printcolumn:name="Ready",JSONPath=".status.conditions[?(@.type=='Ready')].status",type="string",description="When 'True', the most recent reconcile of the resource succeeded"
// +kubebuilder:printcolumn:name="Status",JSONPath=".status.conditions[?(@.type=='Ready')].reason",type="string",description="The reason for the value in 'Ready'"
// +kubebuilder:printcolumn:name="Status Age",JSONPath=".status.conditions[?(@.type=='Ready')].lastTransitionTime",type="date",description="The last transition time for the value in 'Status'"

// CESEvaluationDataset is the Schema for the CESEvaluationDataset API
// +k8s:openapi-gen=true
type CESEvaluationDataset struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	// +required
	Spec   CESEvaluationDatasetSpec   `json:"spec,omitempty"`
	Status CESEvaluationDatasetStatus `json:"status,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// CESEvaluationDatasetList contains a list of CESEvaluationDataset
type CESEvaluationDatasetList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []CESEvaluationDataset `json:"items"`
}

func init() {
	SchemeBuilder.Register(&CESEvaluationDataset{}, &CESEvaluationDatasetList{})
}
