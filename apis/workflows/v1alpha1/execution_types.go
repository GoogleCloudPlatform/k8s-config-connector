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

package v1alpha1

import (
	"github.com/GoogleCloudPlatform/k8s-config-connector/apis/refs/v1beta1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/apis/k8s/v1alpha1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

var WorkflowsExecutionGVK = GroupVersion.WithKind("WorkflowsExecution")

// WorkflowsExecutionSpec defines the desired state of WorkflowsExecution
// +kcc:proto=google.cloud.workflows.executions.v1.Execution
type WorkflowsExecutionSpec struct {

	// Input parameters of the execution represented as a JSON string.
	//  The size limit is 32KB.
	//
	//  *Note*: If you are using the REST API directly to run your workflow, you
	//  must escape any JSON string value of `argument`. Example:
	//  `'{"argument":"{\"firstName\":\"FIRST\",\"lastName\":\"LAST\"}"}'`
	// +kcc:proto:field=google.cloud.workflows.executions.v1.Execution.argument
	Argument *string `json:"argument,omitempty"`

	// The call logging level associated to this execution.
	// +kcc:proto:field=google.cloud.workflows.executions.v1.Execution.call_log_level
	CallLogLevel *string `json:"callLogLevel,omitempty"`

	// Labels associated with this execution.
	//  Labels can contain at most 64 entries. Keys and values can be no longer
	//  than 63 characters and can only contain lowercase letters, numeric
	//  characters, underscores, and dashes. Label keys must start with a letter.
	//  International characters are allowed.
	//  By default, labels are inherited from the workflow but are overridden by
	//  any labels associated with the execution.
	// +kcc:proto:field=google.cloud.workflows.executions.v1.Execution.labels
	Labels map[string]string `json:"labels,omitempty"`

	// Required. Name of the workflow for which an execution should be created.
	// Format: projects/{project}/locations/{location}/workflows/{workflow}.
	// The latest revision of the workflow will be used.
	Parent *Parent `json:"parent,omitempty"`

	// The WorkflowsExecution name. If not given, the metadata.name will be used.
	ResourceID *string `json:"resourceID,omitempty"`
}

type Parent struct {
	// Required. The location of the application.
	Location string `json:"location,omitempty"`

	// Required. The host project of the application.
	ProjectRef *v1beta1.ProjectRef `json:"projectRef,omitempty"`

	// Required.
	WorkflowRef *v1beta1.WorkflowRef `json:"workflowRef,omitempty"`
}

// WorkflowsExecutionStatus defines the config connector machine state of WorkflowsExecution
type WorkflowsExecutionStatus struct {
	/* Conditions represent the latest available observations of the
	   object's current state. */
	Conditions []v1alpha1.Condition `json:"conditions,omitempty"`

	// ObservedGeneration is the generation of the resource that was most recently observed by the Config Connector controller. If this is equal to metadata.generation, then that means that the current reported status reflects the most recent desired state of the resource.
	ObservedGeneration *int64 `json:"observedGeneration,omitempty"`

	// A unique specifier for the WorkflowsExecution resource in GCP.
	ExternalRef *string `json:"externalRef,omitempty"`

	// ObservedState is the state of the resource as most recently observed in GCP.
	ObservedState *WorkflowsExecutionObservedState `json:"observedState,omitempty"`
}

// WorkflowsExecutionObservedState is the state of the WorkflowsExecution resource as most recently observed in GCP.
// +kcc:proto=google.cloud.workflows.executions.v1.Execution
type WorkflowsExecutionObservedState struct {

	// Output only. Marks the beginning of execution.
	// +kcc:proto:field=google.cloud.workflows.executions.v1.Execution.start_time
	StartTime *string `json:"startTime,omitempty"`

	// Output only. Marks the end of execution, successful or not.
	// +kcc:proto:field=google.cloud.workflows.executions.v1.Execution.end_time
	EndTime *string `json:"endTime,omitempty"`

	// Output only. Measures the duration of the execution.
	// +kcc:proto:field=google.cloud.workflows.executions.v1.Execution.duration
	Duration *string `json:"duration,omitempty"`

	// Output only. Current state of the execution.
	// +kcc:proto:field=google.cloud.workflows.executions.v1.Execution.state
	State *string `json:"state,omitempty"`

	// Output only. Output of the execution represented as a JSON string. The
	//  value can only be present if the execution's state is `SUCCEEDED`.
	// +kcc:proto:field=google.cloud.workflows.executions.v1.Execution.result
	Result *string `json:"result,omitempty"`

	// Output only. The error which caused the execution to finish prematurely.
	//  The value is only present if the execution's state is `FAILED`
	//  or `CANCELLED`.
	// +kcc:proto:field=google.cloud.workflows.executions.v1.Execution.error
	Error *Execution_Error `json:"error,omitempty"`

	// Output only. Revision of the workflow this execution is using.
	// +kcc:proto:field=google.cloud.workflows.executions.v1.Execution.workflow_revision_id
	WorkflowRevisionID *string `json:"workflowRevisionID,omitempty"`

	// Output only. Status tracks the current steps and progress data of this
	//  execution.
	// +kcc:proto:field=google.cloud.workflows.executions.v1.Execution.status
	Status *Execution_Status `json:"status,omitempty"`

	// Output only. Error regarding the state of the Execution resource. For
	//  example, this field will have error details if the execution data is
	//  unavailable due to revoked KMS key permissions.
	// +kcc:proto:field=google.cloud.workflows.executions.v1.Execution.state_error
	StateError *Execution_StateError `json:"stateError,omitempty"`
}

// +genclient
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// TODO(user): make sure the pluralizaiton below is correct
// +kubebuilder:resource:categories=gcp,shortName=gcpworkflowsexecution;gcpworkflowsexecutions
// +kubebuilder:subresource:status
// +kubebuilder:metadata:labels="cnrm.cloud.google.com/managed-by-kcc=true";"cnrm.cloud.google.com/system=true"
// +kubebuilder:printcolumn:name="Age",JSONPath=".metadata.creationTimestamp",type="date"
// +kubebuilder:printcolumn:name="Ready",JSONPath=".status.conditions[?(@.type=='Ready')].status",type="string",description="When 'True', the most recent reconcile of the resource succeeded"
// +kubebuilder:printcolumn:name="Status",JSONPath=".status.conditions[?(@.type=='Ready')].reason",type="string",description="The reason for the value in 'Ready'"
// +kubebuilder:printcolumn:name="Status Age",JSONPath=".status.conditions[?(@.type=='Ready')].lastTransitionTime",type="date",description="The last transition time for the value in 'Status'"

// WorkflowsExecution is the Schema for the WorkflowsExecution API
// +k8s:openapi-gen=true
type WorkflowsExecution struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	// +required
	Spec   WorkflowsExecutionSpec   `json:"spec,omitempty"`
	Status WorkflowsExecutionStatus `json:"status,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// WorkflowsExecutionList contains a list of WorkflowsExecution
type WorkflowsExecutionList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []WorkflowsExecution `json:"items"`
}

func init() {
	SchemeBuilder.Register(&WorkflowsExecution{}, &WorkflowsExecutionList{})
}
