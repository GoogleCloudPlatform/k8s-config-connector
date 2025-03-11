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

var BatchJobGVK = GroupVersion.WithKind("BatchJob")

// BatchJobSpec defines the desired state of BatchJob
// +kcc:proto=google.cloud.batch.v1.Job
type BatchJobSpec struct {

	// Priority of the Job.
	//  The valid value range is [0, 100). Default value is 0.
	//  Higher value indicates higher priority.
	//  A job with higher priority value is more likely to run earlier if all other
	//  requirements are satisfied.
	// +kcc:proto:field=google.cloud.batch.v1.Job.priority
	Priority *int64 `json:"priority,omitempty"`

	// Required. TaskGroups in the Job. Only one TaskGroup is supported now.
	// +kcc:proto:field=google.cloud.batch.v1.Job.task_groups
	TaskGroups []TaskGroup `json:"taskGroups,omitempty"`

	// Compute resource allocation for all TaskGroups in the Job.
	// +kcc:proto:field=google.cloud.batch.v1.Job.allocation_policy
	AllocationPolicy *AllocationPolicy `json:"allocationPolicy,omitempty"`

	// Custom labels to apply to the job and any Cloud Logging
	//  [LogEntry](https://cloud.google.com/logging/docs/reference/v2/rest/v2/LogEntry)
	//  that it generates.
	//
	//  Use labels to group and describe the resources they are applied to. Batch
	//  automatically applies predefined labels and supports multiple `labels`
	//  fields for each job, which each let you apply custom labels to various
	//  resources. Label names that start with "goog-" or "google-" are
	//  reserved for predefined labels. For more information about labels with
	//  Batch, see
	//  [Organize resources using
	//  labels](https://cloud.google.com/batch/docs/organize-resources-using-labels).
	// +kcc:proto:field=google.cloud.batch.v1.Job.labels
	Labels map[string]string `json:"labels,omitempty"`

	// Log preservation policy for the Job.
	// +kcc:proto:field=google.cloud.batch.v1.Job.logs_policy
	LogsPolicy *LogsPolicy `json:"logsPolicy,omitempty"`

	// Notification configurations.
	// +kcc:proto:field=google.cloud.batch.v1.Job.notifications
	Notifications []JobNotification `json:"notifications,omitempty"`

	// Required. The parent resource name where the Job will be created. Pattern: "projects/{project}/locations/{location}"
	Parent *Parent `json:"parent,omitempty"`

	// The BatchJob name. If not given, the metadata.name will be used.
	ResourceID *string `json:"resourceID,omitempty"`
}

type Parent struct {
	// Immutable. The location where the alloydb cluster should reside.
	// +required
	Location *string `json:"location,omitempty"`

	// The project that this resource belongs to.
	// +required
	ProjectRef *v1beta1.ProjectRef `json:"projectRef,omitempty"`
}

// BatchJobStatus defines the config connector machine state of BatchJob
type BatchJobStatus struct {
	/* Conditions represent the latest available observations of the
	   object's current state. */
	Conditions []v1alpha1.Condition `json:"conditions,omitempty"`

	// ObservedGeneration is the generation of the resource that was most recently observed by the Config Connector controller. If this is equal to metadata.generation, then that means that the current reported status reflects the most recent desired state of the resource.
	ObservedGeneration *int64 `json:"observedGeneration,omitempty"`

	// A unique specifier for the BatchJob resource in GCP.
	ExternalRef *string `json:"externalRef,omitempty"`

	// ObservedState is the state of the resource as most recently observed in GCP.
	ObservedState *BatchJobObservedState `json:"observedState,omitempty"`
}

// BatchJobObservedState is the state of the BatchJob resource as most recently observed in GCP.
// +kcc:proto=google.cloud.batch.v1.Job
type BatchJobObservedState struct {
	// Output only. Job name.
	//  For example: "projects/123456/locations/us-central1/jobs/job01".
	// +kcc:proto:field=google.cloud.batch.v1.Job.name
	Name *string `json:"name,omitempty"`

	// Output only. A system generated unique ID for the Job.
	// +kcc:proto:field=google.cloud.batch.v1.Job.uid
	Uid *string `json:"uid,omitempty"`

	// Required. TaskGroups in the Job. Only one TaskGroup is supported now.
	// +kcc:proto:field=google.cloud.batch.v1.Job.task_groups
	TaskGroups []TaskGroupObservedState `json:"taskGroups,omitempty"`

	// Output only. Job status. It is read only for users.
	// +kcc:proto:field=google.cloud.batch.v1.Job.status
	Status *JobStatus `json:"status,omitempty"`

	// Output only. When the Job was created.
	// +kcc:proto:field=google.cloud.batch.v1.Job.create_time
	CreateTime *string `json:"createTime,omitempty"`

	// Output only. The last time the Job was updated.
	// +kcc:proto:field=google.cloud.batch.v1.Job.update_time
	UpdateTime *string `json:"updateTime,omitempty"`
}

// +genclient
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// TODO(user): make sure the pluralizaiton below is correct
// +kubebuilder:resource:categories=gcp,shortName=gcpbatchjob;gcpbatchjobs
// +kubebuilder:subresource:status
// +kubebuilder:metadata:labels="cnrm.cloud.google.com/managed-by-kcc=true";"cnrm.cloud.google.com/system=true"
// +kubebuilder:printcolumn:name="Age",JSONPath=".metadata.creationTimestamp",type="date"
// +kubebuilder:printcolumn:name="Ready",JSONPath=".status.conditions[?(@.type=='Ready')].status",type="string",description="When 'True', the most recent reconcile of the resource succeeded"
// +kubebuilder:printcolumn:name="Status",JSONPath=".status.conditions[?(@.type=='Ready')].reason",type="string",description="The reason for the value in 'Ready'"
// +kubebuilder:printcolumn:name="Status Age",JSONPath=".status.conditions[?(@.type=='Ready')].lastTransitionTime",type="date",description="The last transition time for the value in 'Status'"

// BatchJob is the Schema for the BatchJob API
// +k8s:openapi-gen=true
type BatchJob struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	// +required
	Spec   BatchJobSpec   `json:"spec,omitempty"`
	Status BatchJobStatus `json:"status,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// BatchJobList contains a list of BatchJob
type BatchJobList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []BatchJob `json:"items"`
}

func init() {
	SchemeBuilder.Register(&BatchJob{}, &BatchJobList{})
}
