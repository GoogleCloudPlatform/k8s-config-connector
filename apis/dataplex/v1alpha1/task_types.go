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
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/apis/k8s/v1alpha1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

var DataplexTaskGVK = GroupVersion.WithKind("DataplexTask")

// DataplexTaskSpec defines the desired state of DataplexTask
// +kcc:proto=google.cloud.dataplex.v1.Task
type DataplexTaskSpec struct {
	// The DataplexTask name. If not given, the metadata.name will be used.
	ResourceID *string `json:"resourceID,omitempty"` // Existing field preserved

	// Optional. Description of the task.
	// +optional
	Description *string `json:"description,omitempty"`

	// Optional. User friendly display name.
	// +optional
	DisplayName *string `json:"displayName,omitempty"`

	// Optional. User-defined labels for the task.
	// +optional
	Labels map[string]string `json:"labels,omitempty"`

	// Required. Spec related to how often and when a task should be triggered.
	// +required
	TriggerSpec *Task_TriggerSpec `json:"triggerSpec,omitempty"`

	// Required. Spec related to how a task is executed.
	// +required
	ExecutionSpec *Task_ExecutionSpec `json:"executionSpec,omitempty"`

	// Config related to running custom Spark tasks.
	// Exactly one of spark or notebook must be set.
	// +optional
	Spark *Task_SparkTaskConfig `json:"spark,omitempty"`

	// Config related to running scheduled Notebooks.
	// Exactly one of spark or notebook must be set.
	// +optional
	Notebook *Task_NotebookTaskConfig `json:"notebook,omitempty"`
}

// DataplexTaskStatus defines the config connector machine state of DataplexTask
type DataplexTaskStatus struct {
	/* Conditions represent the latest available observations of the
	   object's current state. */
	Conditions []v1alpha1.Condition `json:"conditions,omitempty"`

	// ObservedGeneration is the generation of the resource that was most recently observed by the Config Connector controller. If this is equal to metadata.generation, then that means that the current reported status reflects the most recent desired state of the resource.
	ObservedGeneration *int64 `json:"observedGeneration,omitempty"`

	// A unique specifier for the DataplexTask resource in GCP.
	ExternalRef *string `json:"externalRef,omitempty"`

	// ObservedState is the state of the resource as most recently observed in GCP.
	ObservedState *DataplexTaskObservedState `json:"observedState,omitempty"`
}

// DataplexTaskObservedState is the state of the DataplexTask resource as most recently observed in GCP.
// Based on fields from Task_ExecutionStatusObservedState in types.generated.go
type DataplexTaskObservedState struct {
	// Output only. Last update time of the status.
	// +kcc:proto:field=google.cloud.dataplex.v1.Task.ExecutionStatus.update_time // Copied from Task_ExecutionStatusObservedState
	UpdateTime *string `json:"updateTime,omitempty"`

	// Output only. latest job execution
	// +kcc:proto:field=google.cloud.dataplex.v1.Task.ExecutionStatus.latest_job // Copied from Task_ExecutionStatusObservedState
	LatestJob *Job `json:"latestJob,omitempty"` // Requires 'Job' type defined in types.generated.go
}

// +genclient
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// TODO(user): make sure the pluralizaiton below is correct
// +kubebuilder:resource:categories=gcp,shortName=gcpdataplextask;gcpdataplextasks
// +kubebuilder:subresource:status
// +kubebuilder:metadata:labels="cnrm.cloud.google.com/managed-by-kcc=true";"cnrm.cloud.google.com/system=true"
// +kubebuilder:printcolumn:name="Age",JSONPath=".metadata.creationTimestamp",type="date"
// +kubebuilder:printcolumn:name="Ready",JSONPath=".status.conditions[?(@.type=='Ready')].status",type="string",description="When 'True', the most recent reconcile of the resource succeeded"
// +kubebuilder:printcolumn:name="Status",JSONPath=".status.conditions[?(@.type=='Ready')].reason",type="string",description="The reason for the value in 'Ready'"
// +kubebuilder:printcolumn:name="Status Age",JSONPath=".status.conditions[?(@.type=='Ready')].lastTransitionTime",type="date",description="The last transition time for the value in 'Status'"

// DataplexTask is the Schema for the DataplexTask API
// +k8s:openapi-gen=true
type DataplexTask struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	// +required
	Spec   DataplexTaskSpec   `json:"spec,omitempty"`
	Status DataplexTaskStatus `json:"status,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// DataplexTaskList contains a list of DataplexTask
type DataplexTaskList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []DataplexTask `json:"items"`
}

func init() {
	SchemeBuilder.Register(&DataplexTask{}, &DataplexTaskList{})
}
