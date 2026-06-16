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

var CCInsightsIssueModelGVK = GroupVersion.WithKind("CCInsightsIssueModel")

// CCInsightsIssueModelSpec defines the desired state of CCInsightsIssueModel
// +kcc:spec:proto=google.cloud.contactcenterinsights.v1.IssueModel
type CCInsightsIssueModelSpec struct {
	// The project that this resource belongs to.
	ProjectRef *refsv1beta1.ProjectRef `json:"projectRef"`

	// The location of this resource.
	Location string `json:"location"`

	// The CCInsightsIssueModel name. If not given, the metadata.name will be used.
	ResourceID *string `json:"resourceID,omitempty"`

	// The representative name for the issue model.
	DisplayName *string `json:"displayName,omitempty"`

	// Configs for the input data that used to create the issue model.
	InputDataConfig *IssueModel_InputDataConfig `json:"inputDataConfig,omitempty"`

	// Type of the model.
	ModelType *string `json:"modelType,omitempty"`

	// Language of the model.
	LanguageCode *string `json:"languageCode,omitempty"`
}

// CCInsightsIssueModelStatus defines the config connector machine state of CCInsightsIssueModel
type CCInsightsIssueModelStatus struct {
	/* Conditions represent the latest available observations of the
	   object's current state. */
	Conditions []v1alpha1.Condition `json:"conditions,omitempty"`

	// ObservedGeneration is the generation of the resource that was most recently observed by the Config Connector controller. If this is equal to metadata.generation, then that means that the current reported status reflects the most recent desired state of the resource.
	ObservedGeneration *int64 `json:"observedGeneration,omitempty"`

	// A unique specifier for the CCInsightsIssueModel resource in Google Cloud.
	ExternalRef *string `json:"externalRef,omitempty"`

	// ObservedState is the state of the resource as most recently observed in Google Cloud.
	ObservedState *CCInsightsIssueModelObservedState `json:"observedState,omitempty"`
}

// CCInsightsIssueModelObservedState is the state of the CCInsightsIssueModel resource as most recently observed in Google Cloud.
// +kcc:observedstate:proto=google.cloud.contactcenterinsights.v1.IssueModel
type CCInsightsIssueModelObservedState struct {
	// Output only. The time at which this issue model was created.
	CreateTime *string `json:"createTime,omitempty"`

	// Output only. The most recent time at which the issue model was updated.
	UpdateTime *string `json:"updateTime,omitempty"`

	// Output only. Number of issues in this issue model.
	IssueCount *int64 `json:"issueCount,omitempty"`

	// Output only. State of the model.
	State *string `json:"state,omitempty"`

	// Configs for the input data that used to create the issue model.
	InputDataConfig *IssueModel_InputDataConfigObservedState `json:"inputDataConfig,omitempty"`

	// Output only. Immutable. The issue model's label statistics on its training data.
	TrainingStats *IssueModelLabelStats `json:"trainingStats,omitempty"`
}

// +genclient
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:resource:categories=gcp,shortName=gcpccinsightsissuemodel;gcpccinsightsissuemodels
// +kubebuilder:subresource:status
// +kubebuilder:metadata:labels="cnrm.cloud.google.com/managed-by-kcc=true"
// +kubebuilder:metadata:labels="cnrm.cloud.google.com/system=true"
// +kubebuilder:metadata:labels="cnrm.cloud.google.com/stability-level=alpha"
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
