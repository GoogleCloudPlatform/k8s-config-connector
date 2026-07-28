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

var VertexAIModelMonitoringJobGVK = GroupVersion.WithKind("VertexAIModelMonitoringJob")

// VertexAIModelMonitoringJobSpec defines the desired state of VertexAIModelMonitoringJob
type VertexAIModelMonitoringJobSpec struct {
	// The project that this resource belongs to.
	// +required
	ProjectRef *refsv1beta1.ProjectRef `json:"projectRef"`

	// The location of this resource.
	// +required
	Location *string `json:"location"`

	// The parent ModelMonitor of this resource.
	// +required
	ModelMonitorRef *VertexAIModelMonitorRef `json:"modelMonitorRef"`

	// The VertexAIModelMonitoringJob name. If not given, the metadata.name will be used.
	ResourceID *string `json:"resourceID,omitempty"`

	// The display name of the ModelMonitoringJob.
	//  The name can be up to 128 characters long and can consist of any UTF-8.
	DisplayName *string `json:"displayName,omitempty"`

	// Monitoring monitoring job spec. It outlines the specifications for
	//  monitoring objectives, notifications, and result exports. If left blank,
	//  the default monitoring specifications from the top-level resource
	//  'ModelMonitor' will be applied. If provided, we will use the specification
	//  defined here rather than the default one.
	ModelMonitoringSpec *ModelMonitoringSpec `json:"modelMonitoringSpec,omitempty"`
}

// VertexAIModelMonitoringJobStatus defines the config connector machine state of VertexAIModelMonitoringJob
type VertexAIModelMonitoringJobStatus struct {
	/* Conditions represent the latest available observations of the
	   object's current state. */
	Conditions []v1alpha1.Condition `json:"conditions,omitempty"`

	// ObservedGeneration is the generation of the resource that was most recently observed by the Config Connector controller. If this is equal to metadata.generation, then that means that the current reported status reflects the most recent desired state of the resource.
	ObservedGeneration *int64 `json:"observedGeneration,omitempty"`

	// A unique specifier for the VertexAIModelMonitoringJob resource in GCP.
	ExternalRef *string `json:"externalRef,omitempty"`

	// ObservedState is the state of the resource as most recently observed in GCP.
	ObservedState *VertexAIModelMonitoringJobObservedState `json:"observedState,omitempty"`
}

// VertexAIModelMonitoringJobObservedState is the state of the VertexAIModelMonitoringJob resource as most recently observed in GCP.
type VertexAIModelMonitoringJobObservedState struct {
	// Output only. Resource name of a ModelMonitoringJob. Format:
	//  `projects/{project_id}/locations/{location_id}/modelMonitors/{model_monitor_id}/modelMonitoringJobs/{model_monitoring_job_id}`
	Name *string `json:"name,omitempty"`

	// Output only. Timestamp when this ModelMonitoringJob was created.
	CreateTime *string `json:"createTime,omitempty"`

	// Output only. Timestamp when this ModelMonitoringJob was updated most
	//  recently.
	UpdateTime *string `json:"updateTime,omitempty"`

	// Output only. The state of the monitoring job.
	//   * When the job is still creating, the state will be 'JOB_STATE_PENDING'.
	//   * Once the job is successfully created, the state will be
	//     'JOB_STATE_RUNNING'.
	//   * Once the job is finished, the state will be one of
	//     'JOB_STATE_FAILED', 'JOB_STATE_SUCCEEDED',
	//     'JOB_STATE_PARTIALLY_SUCCEEDED'.
	State *string `json:"state,omitempty"`

	// Output only. Schedule resource name. It will only appear when this job is
	//  triggered by a schedule.
	Schedule *string `json:"schedule,omitempty"`

	// Output only. Execution results for all the monitoring objectives.
	JobExecutionDetail *ModelMonitoringJobExecutionDetail `json:"jobExecutionDetail,omitempty"`

	// Output only. Timestamp when this ModelMonitoringJob was scheduled. It will
	//  only appear when this job is triggered by a schedule.
	ScheduleTime *string `json:"scheduleTime,omitempty"`
}

// +genclient
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:resource:categories=gcp,shortName=gcpvertexaimodelmonitoringjob;gcpvertexaimodelmonitoringjobs
// +kubebuilder:subresource:status
// +kubebuilder:metadata:labels="cnrm.cloud.google.com/managed-by-kcc=true"
// +kubebuilder:metadata:labels="cnrm.cloud.google.com/system=true"
// +kubebuilder:metadata:labels="cnrm.cloud.google.com/stability-level=alpha"
// +kubebuilder:printcolumn:name="Age",JSONPath=".metadata.creationTimestamp",type="date"
// +kubebuilder:printcolumn:name="Ready",JSONPath=".status.conditions[?(@.type=='Ready')].status",type="string",description="When 'True', the most recent reconcile of the resource succeeded"
// +kubebuilder:printcolumn:name="Status",JSONPath=".status.conditions[?(@.type=='Ready')].reason",type="string",description="The reason for the value in 'Ready'"
// +kubebuilder:printcolumn:name="Status Age",JSONPath=".status.conditions[?(@.type=='Ready')].lastTransitionTime",type="date",description="The last transition time for the value in 'Status'"

// VertexAIModelMonitoringJob is the Schema for the VertexAIModelMonitoringJob API
// +k8s:openapi-gen=true
type VertexAIModelMonitoringJob struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	// +required
	Spec   VertexAIModelMonitoringJobSpec   `json:"spec,omitempty"`
	Status VertexAIModelMonitoringJobStatus `json:"status,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// VertexAIModelMonitoringJobList contains a list of VertexAIModelMonitoringJob
type VertexAIModelMonitoringJobList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []VertexAIModelMonitoringJob `json:"items"`
}

func init() {
	SchemeBuilder.Register(&VertexAIModelMonitoringJob{}, &VertexAIModelMonitoringJobList{})
}
