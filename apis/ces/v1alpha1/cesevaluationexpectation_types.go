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

// API sources for CESEvaluationExpectation, recorded by generate-types:
// +kcc:source:proto=https://github.com/googleapis/googleapis/blob/1765b559c42386788ff0c6412491277b4791107a/google/cloud/ces/v1beta/evaluation.proto
// +kcc:source:service-docs=https://docs.cloud.google.com/customer-engagement-ai/conversational-agents/ps
// +kcc:guess=source-link reason=verify-resource-docs-link
// +kcc:source:resource-docs=https://docs.cloud.google.com/customer-engagement-ai/conversational-agents/ps/docs/reference/rest/v1beta/projects.locations.apps.evaluationExpectations

package v1alpha1

import (
	refsv1beta1 "github.com/GoogleCloudPlatform/k8s-config-connector/apis/refs/v1beta1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/apis/k8s/v1alpha1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

var CESEvaluationExpectationGVK = GroupVersion.WithKind("CESEvaluationExpectation")

// CESEvaluationExpectationSpec defines the desired state of CESEvaluationExpectation
// +kcc:spec:proto=google.cloud.ces.v1beta.EvaluationExpectation
// +kcc:required-from-proto
type CESEvaluationExpectationSpec struct {
	// The project that this resource belongs to.
	ProjectRef *refsv1beta1.ProjectRef `json:"projectRef"`

	// The location of this resource.
	Location string `json:"location"`

	// The CESEvaluationExpectation name. If not given, the metadata.name will be used.
	ResourceID *string `json:"resourceID,omitempty"`
	// Optional. Evaluation criteria based on an LLM prompt.
	// +kcc:proto:field=google.cloud.ces.v1beta.EvaluationExpectation.llm_criteria
	LlmCriteria *EvaluationExpectation_LlmCriteria `json:"llmCriteria,omitempty"`

	// Required. User-defined display name. Must be unique within the app.
	// +kcc:proto:field=google.cloud.ces.v1beta.EvaluationExpectation.display_name
	// +required
	DisplayName *string `json:"displayName,omitempty"`

	// Optional. User-defined tags for expectations. Can be used to filter
	//  expectations.
	// +kcc:proto:field=google.cloud.ces.v1beta.EvaluationExpectation.tags
	Tags []string `json:"tags,omitempty"`
}

// CESEvaluationExpectationStatus defines the config connector machine state of CESEvaluationExpectation
type CESEvaluationExpectationStatus struct {
	/* Conditions represent the latest available observations of the
	   object's current state. */
	Conditions []v1alpha1.Condition `json:"conditions,omitempty"`

	// ObservedGeneration is the generation of the resource that was most recently observed by the Config Connector controller. If this is equal to metadata.generation, then that means that the current reported status reflects the most recent desired state of the resource.
	ObservedGeneration *int64 `json:"observedGeneration,omitempty"`

	// A unique specifier for the CESEvaluationExpectation resource in GCP.
	ExternalRef *string `json:"externalRef,omitempty"`

	// ObservedState is the state of the resource as most recently observed in GCP.
	ObservedState *CESEvaluationExpectationObservedState `json:"observedState,omitempty"`
}

// CESEvaluationExpectationObservedState is the state of the CESEvaluationExpectation resource as most recently observed in GCP.
// +kcc:observedstate:proto=google.cloud.ces.v1beta.EvaluationExpectation
type CESEvaluationExpectationObservedState struct {
	// Output only. Timestamp when the evaluation expectation was created.
	// +kcc:proto:field=google.cloud.ces.v1beta.EvaluationExpectation.create_time
	CreateTime *string `json:"createTime,omitempty"`

	// Output only. Timestamp when the evaluation expectation was last updated.
	// +kcc:proto:field=google.cloud.ces.v1beta.EvaluationExpectation.update_time
	UpdateTime *string `json:"updateTime,omitempty"`

	// Output only. Etag used to ensure the object hasn't changed during a
	//  read-modify-write operation. If the etag is empty, the update will
	//  overwrite any concurrent changes.
	// +kcc:proto:field=google.cloud.ces.v1beta.EvaluationExpectation.etag
	Etag *string `json:"etag,omitempty"`
}

// +genclient
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:resource:categories=gcp,shortName=gcpcesevaluationexpectation;gcpcesevaluationexpectations
// +kubebuilder:subresource:status
// +kubebuilder:metadata:labels="cnrm.cloud.google.com/managed-by-kcc=true"
// +kubebuilder:metadata:labels="cnrm.cloud.google.com/system=true"
// +kubebuilder:printcolumn:name="Age",JSONPath=".metadata.creationTimestamp",type="date"
// +kubebuilder:printcolumn:name="Ready",JSONPath=".status.conditions[?(@.type=='Ready')].status",type="string",description="When 'True', the most recent reconcile of the resource succeeded"
// +kubebuilder:printcolumn:name="Status",JSONPath=".status.conditions[?(@.type=='Ready')].reason",type="string",description="The reason for the value in 'Ready'"
// +kubebuilder:printcolumn:name="Status Age",JSONPath=".status.conditions[?(@.type=='Ready')].lastTransitionTime",type="date",description="The last transition time for the value in 'Status'"

// CESEvaluationExpectation is the Schema for the CESEvaluationExpectation API
// +k8s:openapi-gen=true
type CESEvaluationExpectation struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	// +required
	Spec   CESEvaluationExpectationSpec   `json:"spec,omitempty"`
	Status CESEvaluationExpectationStatus `json:"status,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// CESEvaluationExpectationList contains a list of CESEvaluationExpectation
type CESEvaluationExpectationList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []CESEvaluationExpectation `json:"items"`
}

func init() {
	SchemeBuilder.Register(&CESEvaluationExpectation{}, &CESEvaluationExpectationList{})
}
