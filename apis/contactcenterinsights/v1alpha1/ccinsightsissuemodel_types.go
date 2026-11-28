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

// API sources for CCInsightsIssueModel, recorded by generate-types:
// +kcc:source:proto=https://github.com/googleapis/googleapis/blob/1765b559c42386788ff0c6412491277b4791107a/google/cloud/contactcenterinsights/v1/resources.proto
// +kcc:source:service-docs=https://cloud.google.com/contact-center/insights/docs
// +kcc:source:resource-docs=https://docs.cloud.google.com/contact-center/insights/docs/reference/rest/v1/projects.locations.issueModels

package v1alpha1

import (
	parent "github.com/GoogleCloudPlatform/k8s-config-connector/apis/common/parent"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/apis/k8s/v1alpha1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

var CCInsightsIssueModelGVK = GroupVersion.WithKind("CCInsightsIssueModel")

// CCInsightsIssueModelSpec defines the desired state of CCInsightsIssueModel
// +kcc:spec:proto=google.cloud.contactcenterinsights.v1.IssueModel
// +kcc:required-from-proto
type CCInsightsIssueModelSpec struct {
	*parent.ProjectAndLocationRef `json:",inline"`

	// Immutable. Optional. The service-generated name of the resource. Used for acquisition only. Leave unset to create a new resource.
	ResourceID *string `json:"resourceID,omitempty"`

	// The representative name for the issue model.
	// +kcc:proto:field=google.cloud.contactcenterinsights.v1.IssueModel.display_name
	DisplayName *string `json:"displayName,omitempty"`

	// Configs for the input data that used to create the issue model.
	// +kcc:proto:field=google.cloud.contactcenterinsights.v1.IssueModel.input_data_config
	InputDataConfig *IssueModel_InputDataConfig `json:"inputDataConfig,omitempty"`

	// Type of the model.
	// +kcc:proto:field=google.cloud.contactcenterinsights.v1.IssueModel.model_type
	ModelType *string `json:"modelType,omitempty"`

	// Language of the model.
	// +kcc:proto:field=google.cloud.contactcenterinsights.v1.IssueModel.language_code
	LanguageCode *string `json:"languageCode,omitempty"`
}

// CCInsightsIssueModelStatus defines the config connector machine state of CCInsightsIssueModel
type CCInsightsIssueModelStatus struct {
	/* Conditions represent the latest available observations of the
	   object's current state. */
	Conditions []v1alpha1.Condition `json:"conditions,omitempty"`

	// ObservedGeneration is the generation of the resource that was most recently observed by the Config Connector controller. If this is equal to metadata.generation, then that means that the current reported status reflects the most recent desired state of the resource.
	ObservedGeneration *int64 `json:"observedGeneration,omitempty"`

	// A unique specifier for the CCInsightsIssueModel resource in GCP.
	ExternalRef *string `json:"externalRef,omitempty"`

	// ObservedState is the state of the resource as most recently observed in GCP.
	ObservedState *CCInsightsIssueModelObservedState `json:"observedState,omitempty"`
}

// CCInsightsIssueModelObservedState is the state of the CCInsightsIssueModel resource as most recently observed in GCP.
// +kcc:observedstate:proto=google.cloud.contactcenterinsights.v1.IssueModel
type CCInsightsIssueModelObservedState struct {
	// Output only. The time at which this issue model was created.
	// +kcc:proto:field=google.cloud.contactcenterinsights.v1.IssueModel.create_time
	CreateTime *string `json:"createTime,omitempty"`

	// Output only. The most recent time at which the issue model was updated.
	// +kcc:proto:field=google.cloud.contactcenterinsights.v1.IssueModel.update_time
	UpdateTime *string `json:"updateTime,omitempty"`

	// Output only. Number of issues in this issue model.
	// +kcc:proto:field=google.cloud.contactcenterinsights.v1.IssueModel.issue_count
	IssueCount *int64 `json:"issueCount,omitempty"`

	// Output only. State of the model.
	// +kcc:proto:field=google.cloud.contactcenterinsights.v1.IssueModel.state
	State *string `json:"state,omitempty"`

	// Configs for the input data that used to create the issue model.
	// +kcc:proto:field=google.cloud.contactcenterinsights.v1.IssueModel.input_data_config
	InputDataConfig *IssueModel_InputDataConfigObservedState `json:"inputDataConfig,omitempty"`

	// Output only. Immutable. The issue model's label statistics on its training
	//  data.
	// +kcc:proto:field=google.cloud.contactcenterinsights.v1.IssueModel.training_stats
	TrainingStats *IssueModelLabelStats `json:"trainingStats,omitempty"`
}

// +genclient
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:resource:categories=gcp,shortName=gcpccinsightsissuemodel;gcpccinsightsissuemodels
// +kubebuilder:subresource:status
// +kubebuilder:metadata:labels="cnrm.cloud.google.com/managed-by-kcc=true"
// +kubebuilder:metadata:labels="cnrm.cloud.google.com/system=true"
// +kubebuilder:printcolumn:name="Age",JSONPath=".metadata.creationTimestamp",type="date"
// +kubebuilder:printcolumn:name="Ready",JSONPath=".status.conditions[?(@.type=='Ready')].status",type="string",description="When 'True', the most recent reconcile of the resource succeeded"
// +kubebuilder:printcolumn:name="Status",JSONPath=".status.conditions[?(@.type=='Ready')].reason",type="string",description="The reason for the value in 'Ready'"
// +kubebuilder:printcolumn:name="Status Age",JSONPath=".status.conditions[?(@.type=='Ready')].lastTransitionTime",type="date",description="The last transition time for the value in 'Status'"

// CCInsightsIssueModel is the Schema for the CCInsightsIssueModel API
// +k8s:openapi-gen=true
type CCInsightsIssueModel struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	// +required
	Spec   CCInsightsIssueModelSpec   `json:"spec,omitempty"`
	Status CCInsightsIssueModelStatus `json:"status,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// CCInsightsIssueModelList contains a list of CCInsightsIssueModel
type CCInsightsIssueModelList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []CCInsightsIssueModel `json:"items"`
}

func init() {
	SchemeBuilder.Register(&CCInsightsIssueModel{}, &CCInsightsIssueModelList{})
}
