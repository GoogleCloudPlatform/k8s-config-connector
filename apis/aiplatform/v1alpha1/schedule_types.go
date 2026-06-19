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

var VertexAIScheduleGVK = GroupVersion.WithKind("VertexAISchedule")

// VertexAIScheduleSpec defines the desired state of VertexAISchedule
// +kcc:spec:proto=google.cloud.aiplatform.v1.Schedule
type VertexAIScheduleSpec struct {
	// The project that this resource belongs to.
	// +required
	ProjectRef *refsv1beta1.ProjectRef `json:"projectRef"`

	// The location of this resource.
	// +required
	Location *string `json:"location"`

	// The VertexAISchedule name. If not given, the metadata.name will be used.
	ResourceID *string `json:"resourceID,omitempty"`

	// Required. User provided name of the Schedule.
	// The name can be up to 128 characters long and can consist of any UTF-8
	// characters.
	// +kcc:proto:field=google.cloud.aiplatform.v1.Schedule.display_name
	// +required
	DisplayName *string `json:"displayName"`

	// Optional. Timestamp after which the first run can be scheduled.
	// Default to Schedule create time if not specified.
	// +kcc:proto:field=google.cloud.aiplatform.v1.Schedule.start_time
	StartTime *string `json:"startTime,omitempty"`

	// Optional. Timestamp after which no new runs can be scheduled.
	// If specified, The schedule will be completed when either
	// end_time is reached or when scheduled_run_count >= max_run_count.
	// If not specified, new runs will keep getting scheduled until this Schedule
	// is paused or deleted. Already scheduled runs will be allowed to complete.
	// Unset if not specified.
	// +kcc:proto:field=google.cloud.aiplatform.v1.Schedule.end_time
	EndTime *string `json:"endTime,omitempty"`

	// Optional. Maximum run count of the schedule.
	// If specified, The schedule will be completed when either
	// started_run_count >= max_run_count or when end_time is reached.
	// If not specified, new runs will keep getting scheduled until this Schedule
	// is paused or deleted. Already scheduled runs will be allowed to complete.
	// Unset if not specified.
	// +kcc:proto:field=google.cloud.aiplatform.v1.Schedule.max_run_count
	MaxRunCount *int64 `json:"maxRunCount,omitempty"`

	// Required. Maximum number of runs that can be started concurrently for this
	// Schedule. This is the limit for starting the scheduled requests and not the
	// execution of the operations/jobs created by the requests (if applicable).
	// +kcc:proto:field=google.cloud.aiplatform.v1.Schedule.max_concurrent_run_count
	// +required
	MaxConcurrentRunCount *int64 `json:"maxConcurrentRunCount"`

	// Optional. Whether new scheduled runs can be queued when max_concurrent_runs
	// limit is reached. If set to true, new runs will be queued instead of
	// skipped. Default to false.
	// +kcc:proto:field=google.cloud.aiplatform.v1.Schedule.allow_queueing
	AllowQueueing *bool `json:"allowQueueing,omitempty"`

	// Cron schedule (https://en.wikipedia.org/wiki/Cron) to launch scheduled
	// runs. To explicitly set a timezone to the cron tab, apply a prefix in the
	// cron tab: "CRON_TZ=${IANA_TIME_ZONE}" or "TZ=${IANA_TIME_ZONE}".
	// The ${IANA_TIME_ZONE} may only be a valid string from IANA time zone
	// database. For example, "CRON_TZ=America/New_York 1 * * * *", or
	// "TZ=America/New_York 1 * * * *".
	// +kcc:proto:field=google.cloud.aiplatform.v1.Schedule.cron
	Cron *string `json:"cron,omitempty"`

	// Request for
	// [PipelineService.CreatePipelineJob][google.cloud.aiplatform.v1.PipelineService.CreatePipelineJob].
	// CreatePipelineJobRequest.parent field is required (format:
	// projects/{project}/locations/{location}).
	// +kcc:proto:field=google.cloud.aiplatform.v1.Schedule.create_pipeline_job_request
	CreatePipelineJobRequest *CreatePipelineJobRequest `json:"createPipelineJobRequest,omitempty"`

	// Request for
	// [NotebookService.CreateNotebookExecutionJob][google.cloud.aiplatform.v1.NotebookService.CreateNotebookExecutionJob].
	// +kcc:proto:field=google.cloud.aiplatform.v1.Schedule.create_notebook_execution_job_request
	CreateNotebookExecutionJobRequest *CreateNotebookExecutionJobRequest `json:"createNotebookExecutionJobRequest,omitempty"`
}

// VertexAIScheduleStatus defines the config connector machine state of VertexAISchedule
type VertexAIScheduleStatus struct {
	/* Conditions represent the latest available observations of the
	   object's current state. */
	Conditions []v1alpha1.Condition `json:"conditions,omitempty"`

	// ObservedGeneration is the generation of the resource that was most recently observed by the Config Connector controller. If this is equal to metadata.generation, then that means that the current reported status reflects the most recent desired state of the resource.
	ObservedGeneration *int64 `json:"observedGeneration,omitempty"`

	// A unique specifier for the VertexAISchedule resource in GCP.
	ExternalRef *string `json:"externalRef,omitempty"`

	// ObservedState is the state of the resource as most recently observed in GCP.
	ObservedState *VertexAIScheduleObservedState `json:"observedState,omitempty"`
}

// VertexAIScheduleObservedState is the state of the VertexAISchedule resource as most recently observed in GCP.
// +kcc:observedstate:proto=google.cloud.aiplatform.v1.Schedule
type VertexAIScheduleObservedState struct {
	// Output only. The number of runs started by this schedule.
	// +kcc:proto:field=google.cloud.aiplatform.v1.Schedule.started_run_count
	StartedRunCount *int64 `json:"startedRunCount,omitempty"`

	// Output only. The state of this Schedule.
	// +kcc:proto:field=google.cloud.aiplatform.v1.Schedule.state
	State *string `json:"state,omitempty"`

	// Output only. Timestamp when this Schedule was created.
	// +kcc:proto:field=google.cloud.aiplatform.v1.Schedule.create_time
	CreateTime *string `json:"createTime,omitempty"`

	// Output only. Timestamp when this Schedule was updated.
	// +kcc:proto:field=google.cloud.aiplatform.v1.Schedule.update_time
	UpdateTime *string `json:"updateTime,omitempty"`

	// Output only. Timestamp when this Schedule should schedule the next run.
	// Having a next_run_time in the past means the runs are being started
	// behind schedule.
	// +kcc:proto:field=google.cloud.aiplatform.v1.Schedule.next_run_time
	NextRunTime *string `json:"nextRunTime,omitempty"`

	// Output only. Timestamp when this Schedule was last paused.
	// Unset if never paused.
	// +kcc:proto:field=google.cloud.aiplatform.v1.Schedule.last_pause_time
	LastPauseTime *string `json:"lastPauseTime,omitempty"`

	// Output only. Timestamp when this Schedule was last resumed.
	// Unset if never resumed from pause.
	// +kcc:proto:field=google.cloud.aiplatform.v1.Schedule.last_resume_time
	LastResumeTime *string `json:"lastResumeTime,omitempty"`

	// Output only. Whether to backfill missed runs when the schedule is resumed
	// from PAUSED state. If set to true, all missed runs will be scheduled. New
	// runs will be scheduled after the backfill is complete. Default to false.
	// +kcc:proto:field=google.cloud.aiplatform.v1.Schedule.catch_up
	CatchUp *bool `json:"catchUp,omitempty"`

	// Output only. Response of the last scheduled run.
	// This is the response for starting the scheduled requests and not the
	// execution of the operations/jobs created by the requests (if applicable).
	// Unset if no run has been scheduled yet.
	// +kcc:proto:field=google.cloud.aiplatform.v1.Schedule.last_scheduled_run_response
	LastScheduledRunResponse *Schedule_RunResponse `json:"lastScheduledRunResponse,omitempty"`
}

// +genclient
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:resource:categories=gcp,shortName=gcpvertexaischedule;gcpvertexaischedules
// +kubebuilder:subresource:status
// +kubebuilder:metadata:labels="cnrm.cloud.google.com/managed-by-kcc=true"
// +kubebuilder:metadata:labels="cnrm.cloud.google.com/system=true"
// +kubebuilder:metadata:labels="cnrm.cloud.google.com/stability-level=alpha"
// +kubebuilder:printcolumn:name="Age",JSONPath=".metadata.creationTimestamp",type="date"
// +kubebuilder:printcolumn:name="Ready",JSONPath=".status.conditions[?(@.type=='Ready')].status",type="string",description="When 'True', the most recent reconcile of the resource succeeded"
// +kubebuilder:printcolumn:name="Status",JSONPath=".status.conditions[?(@.type=='Ready')].reason",type="string",description="The reason for the value in 'Ready'"
// +kubebuilder:printcolumn:name="Status Age",JSONPath=".status.conditions[?(@.type=='Ready')].lastTransitionTime",type="date",description="The last transition time for the value in 'Status'"

// VertexAISchedule is the Schema for the VertexAISchedule API
// +k8s:openapi-gen=true
type VertexAISchedule struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	// +required
	Spec   VertexAIScheduleSpec   `json:"spec,omitempty"`
	Status VertexAIScheduleStatus `json:"status,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// VertexAIScheduleList contains a list of VertexAISchedule
type VertexAIScheduleList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []VertexAISchedule `json:"items"`
}

func init() {
	SchemeBuilder.Register(&VertexAISchedule{}, &VertexAIScheduleList{})
}
