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

// API sources for CESScheduledEvaluationRun, recorded by generate-types:
// +kcc:source:proto=https://github.com/googleapis/googleapis/blob/1765b559c42386788ff0c6412491277b4791107a/google/cloud/ces/v1beta/evaluation.proto
// +kcc:source:service-docs=https://docs.cloud.google.com/customer-engagement-ai/conversational-agents/ps
// +kcc:guess=source-link reason=verify-resource-docs-link
// +kcc:source:resource-docs=https://docs.cloud.google.com/customer-engagement-ai/conversational-agents/ps/docs/reference/rest/v1beta/projects.locations.apps.scheduledEvaluationRuns

package v1alpha1

import (
	refsv1beta1 "github.com/GoogleCloudPlatform/k8s-config-connector/apis/refs/v1beta1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/apis/k8s/v1alpha1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

var CESScheduledEvaluationRunGVK = GroupVersion.WithKind("CESScheduledEvaluationRun")

// CESScheduledEvaluationRunSpec defines the desired state of CESScheduledEvaluationRun
// +kcc:spec:proto=google.cloud.ces.v1beta.ScheduledEvaluationRun
// +kcc:required-from-proto
type CESScheduledEvaluationRunSpec struct {
	// The project that this resource belongs to.
	ProjectRef *refsv1beta1.ProjectRef `json:"projectRef"`

	// The location of this resource.
	Location string `json:"location"`

	// The CESScheduledEvaluationRun name. If not given, the metadata.name will be used.
	ResourceID *string `json:"resourceID,omitempty"`
	// Required. User-defined display name of the scheduled evaluation run config.
	// +kcc:proto:field=google.cloud.ces.v1beta.ScheduledEvaluationRun.display_name
	// +required
	DisplayName *string `json:"displayName,omitempty"`

	// Required. The RunEvaluationRequest to schedule
	// +kcc:proto:field=google.cloud.ces.v1beta.ScheduledEvaluationRun.request
	// +required
	Request *RunEvaluationRequest `json:"request,omitempty"`

	// Optional. User-defined description of the scheduled evaluation run.
	// +kcc:proto:field=google.cloud.ces.v1beta.ScheduledEvaluationRun.description
	Description *string `json:"description,omitempty"`

	// Required. Configuration for the timing and frequency with which to execute
	//  the evaluations.
	// +kcc:proto:field=google.cloud.ces.v1beta.ScheduledEvaluationRun.scheduling_config
	// +required
	SchedulingConfig *ScheduledEvaluationRun_SchedulingConfig `json:"schedulingConfig,omitempty"`

	// Optional. Whether this config is active
	// +kcc:proto:field=google.cloud.ces.v1beta.ScheduledEvaluationRun.active
	Active *bool `json:"active,omitempty"`
}

// CESScheduledEvaluationRunStatus defines the config connector machine state of CESScheduledEvaluationRun
type CESScheduledEvaluationRunStatus struct {
	/* Conditions represent the latest available observations of the
	   object's current state. */
	Conditions []v1alpha1.Condition `json:"conditions,omitempty"`

	// ObservedGeneration is the generation of the resource that was most recently observed by the Config Connector controller. If this is equal to metadata.generation, then that means that the current reported status reflects the most recent desired state of the resource.
	ObservedGeneration *int64 `json:"observedGeneration,omitempty"`

	// A unique specifier for the CESScheduledEvaluationRun resource in GCP.
	ExternalRef *string `json:"externalRef,omitempty"`

	// ObservedState is the state of the resource as most recently observed in GCP.
	ObservedState *CESScheduledEvaluationRunObservedState `json:"observedState,omitempty"`
}

// CESScheduledEvaluationRunObservedState is the state of the CESScheduledEvaluationRun resource as most recently observed in GCP.
// +kcc:observedstate:proto=google.cloud.ces.v1beta.ScheduledEvaluationRun
type CESScheduledEvaluationRunObservedState struct {
	// Required. The RunEvaluationRequest to schedule
	// +kcc:proto:field=google.cloud.ces.v1beta.ScheduledEvaluationRun.request
	Request *RunEvaluationRequestObservedState `json:"request,omitempty"`

	// Output only. The last successful EvaluationRun of this scheduled execution.
	//  Format:
	//  `projects/{project}/locations/{location}/apps/{app}/evaluationRuns/{evaluationRun}`
	// +kcc:proto:field=google.cloud.ces.v1beta.ScheduledEvaluationRun.last_completed_run
	LastCompletedRun *string `json:"lastCompletedRun,omitempty"`

	// Output only. The total number of times this run has been executed
	// +kcc:proto:field=google.cloud.ces.v1beta.ScheduledEvaluationRun.total_executions
	TotalExecutions *int32 `json:"totalExecutions,omitempty"`

	// Output only. The next time this is scheduled to execute
	// +kcc:proto:field=google.cloud.ces.v1beta.ScheduledEvaluationRun.next_scheduled_execution_time
	NextScheduledExecutionTime *string `json:"nextScheduledExecutionTime,omitempty"`

	// Output only. Timestamp when the scheduled evaluation run was created.
	// +kcc:proto:field=google.cloud.ces.v1beta.ScheduledEvaluationRun.create_time
	CreateTime *string `json:"createTime,omitempty"`

	// Output only. The user who created the scheduled evaluation run.
	// +kcc:proto:field=google.cloud.ces.v1beta.ScheduledEvaluationRun.created_by
	CreatedBy *string `json:"createdBy,omitempty"`

	// Output only. Timestamp when the evaluation was last updated.
	// +kcc:proto:field=google.cloud.ces.v1beta.ScheduledEvaluationRun.update_time
	UpdateTime *string `json:"updateTime,omitempty"`

	// Output only. The user who last updated the evaluation.
	// +kcc:proto:field=google.cloud.ces.v1beta.ScheduledEvaluationRun.last_updated_by
	LastUpdatedBy *string `json:"lastUpdatedBy,omitempty"`

	// Output only. Etag used to ensure the object hasn't changed during a
	//  read-modify-write operation. If the etag is empty, the update will
	//  overwrite any concurrent changes.
	// +kcc:proto:field=google.cloud.ces.v1beta.ScheduledEvaluationRun.etag
	Etag *string `json:"etag,omitempty"`
}

// +genclient
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:resource:categories=gcp,shortName=gcpcesscheduledevaluationrun;gcpcesscheduledevaluationruns
// +kubebuilder:subresource:status
// +kubebuilder:metadata:labels="cnrm.cloud.google.com/managed-by-kcc=true"
// +kubebuilder:metadata:labels="cnrm.cloud.google.com/system=true"
// +kubebuilder:printcolumn:name="Age",JSONPath=".metadata.creationTimestamp",type="date"
// +kubebuilder:printcolumn:name="Ready",JSONPath=".status.conditions[?(@.type=='Ready')].status",type="string",description="When 'True', the most recent reconcile of the resource succeeded"
// +kubebuilder:printcolumn:name="Status",JSONPath=".status.conditions[?(@.type=='Ready')].reason",type="string",description="The reason for the value in 'Ready'"
// +kubebuilder:printcolumn:name="Status Age",JSONPath=".status.conditions[?(@.type=='Ready')].lastTransitionTime",type="date",description="The last transition time for the value in 'Status'"

// CESScheduledEvaluationRun is the Schema for the CESScheduledEvaluationRun API
// +k8s:openapi-gen=true
type CESScheduledEvaluationRun struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	// +required
	Spec   CESScheduledEvaluationRunSpec   `json:"spec,omitempty"`
	Status CESScheduledEvaluationRunStatus `json:"status,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// CESScheduledEvaluationRunList contains a list of CESScheduledEvaluationRun
type CESScheduledEvaluationRunList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []CESScheduledEvaluationRun `json:"items"`
}

func init() {
	SchemeBuilder.Register(&CESScheduledEvaluationRun{}, &CESScheduledEvaluationRunList{})
}
