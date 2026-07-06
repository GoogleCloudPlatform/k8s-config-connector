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

var NotebooksScheduleGVK = GroupVersion.WithKind("NotebooksSchedule")

// NotebooksScheduleSpec defines the desired state of NotebooksSchedule
// +kcc:spec:proto=google.cloud.notebooks.v1.Schedule
type NotebooksScheduleSpec struct {
	// The Project that this resource belongs to.
	// +required
	ProjectRef *refsv1beta1.ProjectRef `json:"projectRef"`

	// The location of this resource.
	// +required
	Location *string `json:"location"`

	// The NotebooksSchedule name. If not given, the metadata.name will be used.
	ResourceID *string `json:"resourceID,omitempty"`

	// A brief description of this environment.
	// +kcc:proto:field=google.cloud.notebooks.v1.Schedule.description
	Description *string `json:"description,omitempty"`

	// +kcc:proto:field=google.cloud.notebooks.v1.Schedule.state
	State *string `json:"state,omitempty"`

	// Cron-tab formatted schedule by which the job will execute.
	//  Format: minute, hour, day of month, month, day of week,
	//  e.g. `0 0 * * WED` = every Wednesday
	//  More examples: https://crontab.guru/examples.html
	// +kcc:proto:field=google.cloud.notebooks.v1.Schedule.cron_schedule
	CronSchedule *string `json:"cronSchedule,omitempty"`

	// Timezone on which the cron_schedule.
	//  The value of this field must be a time zone name from the tz database.
	//  TZ Database: https://en.wikipedia.org/wiki/List_of_tz_database_time_zones
	//
	//  Note that some time zones include a provision for daylight savings time.
	//  The rules for daylight saving time are determined by the chosen tz.
	//  For UTC use the string "utc". If a time zone is not specified,
	//  the default will be in UTC (also known as GMT).
	// +kcc:proto:field=google.cloud.notebooks.v1.Schedule.time_zone
	TimeZone *string `json:"timeZone,omitempty"`

	// Notebook Execution Template corresponding to this schedule.
	// +kcc:proto:field=google.cloud.notebooks.v1.Schedule.execution_template
	ExecutionTemplate *ExecutionTemplate `json:"executionTemplate,omitempty"`
}

// NotebooksScheduleStatus defines the config connector machine state of NotebooksSchedule
type NotebooksScheduleStatus struct {
	/* Conditions represent the latest available observations of the
	   object's current state. */
	Conditions []v1alpha1.Condition `json:"conditions,omitempty"`

	// ObservedGeneration is the generation of the resource that was most recently observed by the Config Connector controller. If this is equal to metadata.generation, then that means that the current reported status reflects the most recent desired state of the resource.
	ObservedGeneration *int64 `json:"observedGeneration,omitempty"`

	// A unique specifier for the NotebooksSchedule resource in GCP.
	ExternalRef *string `json:"externalRef,omitempty"`

	// ObservedState is the state of the resource as most recently observed in GCP.
	ObservedState *NotebooksScheduleObservedState `json:"observedState,omitempty"`
}

// NotebooksScheduleObservedState is the state of the NotebooksSchedule resource as most recently observed in GCP.
// +kcc:observedstate:proto=google.cloud.notebooks.v1.Schedule
type NotebooksScheduleObservedState struct {
	// Output only. Display name used for UI purposes.
	//  Name can only contain alphanumeric characters, hyphens `-`,
	//  and underscores `_`.
	// +kcc:proto:field=google.cloud.notebooks.v1.Schedule.display_name
	DisplayName *string `json:"displayName,omitempty"`

	// Output only. Time the schedule was created.
	// +kcc:proto:field=google.cloud.notebooks.v1.Schedule.create_time
	CreateTime *string `json:"createTime,omitempty"`

	// Output only. Time the schedule was last updated.
	// +kcc:proto:field=google.cloud.notebooks.v1.Schedule.update_time
	UpdateTime *string `json:"updateTime,omitempty"`

	// Output only. The most recent execution names triggered from this schedule and their
	//  corresponding states.
	// +kcc:proto:field=google.cloud.notebooks.v1.Schedule.recent_executions
	RecentExecutions []NotebooksExecutionObservedState `json:"recentExecutions,omitempty"`
}

// +genclient
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:resource:categories=gcp,shortName=gcpnotebooksschedule;gcpnotebooksschedules
// +kubebuilder:subresource:status
// +kubebuilder:metadata:labels="cnrm.cloud.google.com/managed-by-kcc=true"
// +kubebuilder:metadata:labels="cnrm.cloud.google.com/system=true"
// +kubebuilder:metadata:labels="cnrm.cloud.google.com/stability-level=alpha"
// +kubebuilder:printcolumn:name="Age",JSONPath=".metadata.creationTimestamp",type="date"
// +kubebuilder:printcolumn:name="Ready",JSONPath=".status.conditions[?(@.type=='Ready')].status",type="string",description="When 'True', the most recent reconcile of the resource succeeded"
// +kubebuilder:printcolumn:name="Status",JSONPath=".status.conditions[?(@.type=='Ready')].reason",type="string",description="The reason for the value in 'Ready'"
// +kubebuilder:printcolumn:name="Status Age",JSONPath=".status.conditions[?(@.type=='Ready')].lastTransitionTime",type="date",description="The last transition time for the value in 'Status'"

// NotebooksSchedule is the Schema for the NotebooksSchedule API
// +k8s:openapi-gen=true
type NotebooksSchedule struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	// +required
	Spec   NotebooksScheduleSpec   `json:"spec,omitempty"`
	Status NotebooksScheduleStatus `json:"status,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// NotebooksScheduleList contains a list of NotebooksSchedule
type NotebooksScheduleList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []NotebooksSchedule `json:"items"`
}

func init() {
	SchemeBuilder.Register(&NotebooksSchedule{}, &NotebooksScheduleList{})
}
