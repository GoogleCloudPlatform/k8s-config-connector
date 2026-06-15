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

var WorkloadManagerEvaluationGVK = GroupVersion.WithKind("WorkloadManagerEvaluation")

// WorkloadManagerEvaluationSpec defines the desired state of WorkloadManagerEvaluation
// +kcc:spec:proto=google.cloud.workloadmanager.v1.Evaluation
type WorkloadManagerEvaluationSpec struct {
	// The project that this resource belongs to.
	ProjectRef *refsv1beta1.ProjectRef `json:"projectRef"`

	// The location of this resource.
	Location string `json:"location"`

	// The WorkloadManagerEvaluation name. If not given, the metadata.name will be used.
	ResourceID *string `json:"resourceID,omitempty"`

	// Description of the Evaluation.
	// +optional
	Description *string `json:"description,omitempty"`

	// Resource filter for an evaluation defining the scope of resources to be
	//  evaluated.
	// +optional
	ResourceFilter *ResourceFilter `json:"resourceFilter,omitempty"`

	// The names of the rules used for this evaluation.
	// +optional
	RuleNames []string `json:"ruleNames,omitempty"`

	// Labels as key value pairs.
	// +optional
	Labels map[string]string `json:"labels,omitempty"`

	// Crontab format schedule for scheduled evaluation, currently only supports
	//  the following fixed schedules:
	//  * `0 *-/1 * * *` # Hourly
	//  * `0 *-/6 * * *` # Every 6 hours
	//  * `0 *-/12 * * *` # Every 12 hours
	//  * `0 0 *-/1 * *` # Daily
	//  * `0 0 *-/7 * *` # Weekly
	//  * `0 0 *-/14 * *` # Every 14 days
	//  * `0 0 1 *-/1 *` # Monthly
	// +optional
	Schedule *string `json:"schedule,omitempty"`

	// The Cloud Storage bucket name for custom rules.
	// +optional
	CustomRulesBucket *string `json:"customRulesBucket,omitempty"`

	// Evaluation type.
	// +optional
	// +kubebuilder:validation:Enum=EVALUATION_TYPE_UNSPECIFIED;SAP;SQL_SERVER;OTHER
	EvaluationType *string `json:"evaluationType,omitempty"`

	// Optional. The BigQuery destination for detailed evaluation results.
	//  If this field is specified, the results of each evaluation execution are
	//  exported to BigQuery.
	// +optional
	BigQueryDestination *BigQueryDestination `json:"bigQueryDestination,omitempty"`

	// Optional. Immutable. Customer-managed encryption key name.
	// +optional
	KMSKeyRef *refsv1beta1.KMSCryptoKeyRef `json:"kmsKeyRef,omitempty"`
}

// WorkloadManagerEvaluationStatus defines the config connector machine state of WorkloadManagerEvaluation
type WorkloadManagerEvaluationStatus struct {
	/* Conditions represent the latest available observations of the
	   object's current state. */
	Conditions []v1alpha1.Condition `json:"conditions,omitempty"`

	// ObservedGeneration is the generation of the resource that was most recently observed by the Config Connector controller. If this is equal to metadata.generation, then that means that the current reported status reflects the most recent desired state of the resource.
	ObservedGeneration *int64 `json:"observedGeneration,omitempty"`

	// A unique specifier for the WorkloadManagerEvaluation resource in GCP.
	ExternalRef *string `json:"externalRef,omitempty"`

	// ObservedState is the state of the resource as most recently observed in GCP.
	ObservedState *WorkloadManagerEvaluationObservedState `json:"observedState,omitempty"`
}

// WorkloadManagerEvaluationObservedState is the state of the WorkloadManagerEvaluation resource as most recently observed in GCP.
// +kcc:observedstate:proto=google.cloud.workloadmanager.v1.Evaluation
type WorkloadManagerEvaluationObservedState struct {
	// [Output only] The current lifecycle state of the evaluation resource.
	// +optional
	ResourceStatus *ResourceStatus `json:"resourceStatus,omitempty"`

	// [Output only] Create time stamp.
	// +optional
	CreateTime *string `json:"createTime,omitempty"`

	// [Output only] Update time stamp.
	// +optional
	UpdateTime *string `json:"updateTime,omitempty"`
}

// +genclient
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:resource:categories=gcp,shortName=gcpworkloadmanagerevaluation;gcpworkloadmanagerevaluations
// +kubebuilder:subresource:status
// +kubebuilder:metadata:labels="cnrm.cloud.google.com/managed-by-kcc=true"
// +kubebuilder:metadata:labels="cnrm.cloud.google.com/system=true"
// +kubebuilder:metadata:labels="cnrm.cloud.google.com/stability-level=alpha"
// +kubebuilder:printcolumn:name="Age",JSONPath=".metadata.creationTimestamp",type="date"
// +kubebuilder:printcolumn:name="Ready",JSONPath=".status.conditions[?(@.type=='Ready')].status",type="string",description="When 'True', the most recent reconcile of the resource succeeded"
// +kubebuilder:printcolumn:name="Status",JSONPath=".status.conditions[?(@.type=='Ready')].reason",type="string",description="The reason for the value in 'Ready'"
// +kubebuilder:printcolumn:name="Status Age",JSONPath=".status.conditions[?(@.type=='Ready')].lastTransitionTime",type="date",description="The last transition time for the value in 'Status'"

// WorkloadManagerEvaluation is the Schema for the WorkloadManagerEvaluation API
// +k8s:openapi-gen=true
type WorkloadManagerEvaluation struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	// +required
	Spec   WorkloadManagerEvaluationSpec   `json:"spec,omitempty"`
	Status WorkloadManagerEvaluationStatus `json:"status,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// WorkloadManagerEvaluationList contains a list of WorkloadManagerEvaluation
type WorkloadManagerEvaluationList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []WorkloadManagerEvaluation `json:"items"`
}

func init() {
	SchemeBuilder.Register(&WorkloadManagerEvaluation{}, &WorkloadManagerEvaluationList{})
}
