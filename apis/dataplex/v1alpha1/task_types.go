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

type DataplexTaskParent struct {
	LakeRef *LakeRef `json:"lakeRef,omitempty"`
}

// DataplexTaskSpec defines the desired state of DataplexTask
// +kcc:spec:proto=google.cloud.dataplex.v1.Task
type DataplexTaskSpec struct {
	// The DataplexTask name. If not given, the metadata.name will be used.
	ResourceID *string `json:"resourceID,omitempty"`

	DataplexTaskParent `json:",inline"`

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

type DataplexTaskObservedState struct {
	// Output only. System generated globally unique ID for the task. This ID will
	// be different if the task is deleted and re-created with the same name.
	// +kcc:proto:field=google.cloud.dataplex.v1.Task.uid
	UID *string `json:"uid,omitempty"`
	// Output only. The time when the task was created.
	// +kcc:proto:field=google.cloud.dataplex.v1.Task.create_time
	CreateTime *string `json:"createTime,omitempty"`

	// Output only. The time when the task was last updated.
	// +kcc:proto:field=google.cloud.dataplex.v1.Task.update_time
	UpdateTime *string `json:"updateTime,omitempty"`

	// Output only. Current state of the task.
	State *string `json:"state,omitempty"`

	//  Status of the task execution (e.g. Jobs).
	// +kcc:proto:field=google.cloud.dataplex.v1.Task.ExecutionStatus
	ExecutionStatus *Task_ExecutionStatusObservedState `json:"executionStatus,omitempty"`
}

// Duplicate of Task_ExecutionSpec struct. As ServiceAccount and KMSKey cannot be a reference field when it's in status.
// +kcc:proto=google.cloud.dataplex.v1.Task.ExecutionSpec
type Task_ExecutionSpecObservedState struct {
	// The arguments to pass to the task.
	//  The args can use placeholders of the format ${placeholder} as
	//  part of key/value string. These will be interpolated before passing the
	//  args to the driver. Currently supported placeholders:
	//  - ${task_id}
	//  - ${job_time}
	//  To pass positional args, set the key as TASK_ARGS. The value should be a
	//  comma-separated string of all the positional arguments. To use a
	//  delimiter other than comma, refer to
	//  https://cloud.google.com/sdk/gcloud/reference/topic/escaping. In case of
	//  other keys being present in the args, then TASK_ARGS will be passed as
	//  the last argument.
	// +kcc:proto:field=google.cloud.dataplex.v1.Task.ExecutionSpec.args
	Args map[string]string `json:"args,omitempty"`

	// Service account to use to execute a task.
	//  If not provided, the default Compute service account for the project is
	//  used.
	// +required
	// +kcc:proto:field=google.cloud.dataplex.v1.Task.ExecutionSpec.service_account
	ServiceAccount *string `json:"serviceAccount,omitempty"`

	// The project in which jobs are run. By default, the project
	//  containing the Lake is used. If a project is provided, the
	//  [ExecutionSpec.service_account][google.cloud.dataplex.v1.Task.ExecutionSpec.service_account]
	//  must belong to this project.
	// +kcc:proto:field=google.cloud.dataplex.v1.Task.ExecutionSpec.project
	Project *string `json:"project,omitempty"`

	// The maximum duration after which the job execution is expired.
	// +kcc:proto:field=google.cloud.dataplex.v1.Task.ExecutionSpec.max_job_execution_lifetime
	MaxJobExecutionLifetime *string `json:"maxJobExecutionLifetime,omitempty"`

	// The Cloud KMS key to use for encryption, of the form:
	//  `projects/{project_number}/locations/{location_id}/keyRings/{key-ring-name}/cryptoKeys/{key-name}`.
	// +kcc:proto:field=google.cloud.dataplex.v1.Task.ExecutionSpec.kms_key
	KMSKey *string `json:"kmsKey,omitempty"`
}

// +genclient
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
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
