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

var TasksQueueGVK = GroupVersion.WithKind("TasksQueue")

// TasksQueueSpec defines the desired state of TasksQueue
// +kcc:proto=google.cloud.tasks.v2.Queue
type TasksQueueSpec struct {

	// Overrides for
	//  [task-level
	//  app_engine_routing][google.cloud.tasks.v2.AppEngineHttpRequest.app_engine_routing].
	//  These settings apply only to
	//  [App Engine tasks][google.cloud.tasks.v2.AppEngineHttpRequest] in this
	//  queue. [Http tasks][google.cloud.tasks.v2.HttpRequest] are not affected.
	//
	//  If set, `app_engine_routing_override` is used for all
	//  [App Engine tasks][google.cloud.tasks.v2.AppEngineHttpRequest] in the
	//  queue, no matter what the setting is for the [task-level
	//  app_engine_routing][google.cloud.tasks.v2.AppEngineHttpRequest.app_engine_routing].
	// +kcc:proto:field=google.cloud.tasks.v2.Queue.app_engine_routing_override
	AppEngineRoutingOverride *AppEngineRouting `json:"appEngineRoutingOverride,omitempty"`

	// Rate limits for task dispatches.
	//
	//  [rate_limits][google.cloud.tasks.v2.Queue.rate_limits] and
	//  [retry_config][google.cloud.tasks.v2.Queue.retry_config] are related
	//  because they both control task attempts. However they control task attempts
	//  in different ways:
	//
	//  * [rate_limits][google.cloud.tasks.v2.Queue.rate_limits] controls the total
	//  rate of
	//    dispatches from a queue (i.e. all traffic dispatched from the
	//    queue, regardless of whether the dispatch is from a first
	//    attempt or a retry).
	//  * [retry_config][google.cloud.tasks.v2.Queue.retry_config] controls what
	//  happens to
	//    particular a task after its first attempt fails. That is,
	//    [retry_config][google.cloud.tasks.v2.Queue.retry_config] controls task
	//    retries (the second attempt, third attempt, etc).
	//
	//  The queue's actual dispatch rate is the result of:
	//
	//  * Number of tasks in the queue
	//  * User-specified throttling:
	//  [rate_limits][google.cloud.tasks.v2.Queue.rate_limits],
	//    [retry_config][google.cloud.tasks.v2.Queue.retry_config], and the
	//    [queue's state][google.cloud.tasks.v2.Queue.state].
	//  * System throttling due to `429` (Too Many Requests) or `503` (Service
	//    Unavailable) responses from the worker, high error rates, or to smooth
	//    sudden large traffic spikes.
	// +kcc:proto:field=google.cloud.tasks.v2.Queue.rate_limits
	RateLimits *RateLimits `json:"rateLimits,omitempty"`

	// Settings that determine the retry behavior.
	//
	//  * For tasks created using Cloud Tasks: the queue-level retry settings
	//    apply to all tasks in the queue that were created using Cloud Tasks.
	//    Retry settings cannot be set on individual tasks.
	//  * For tasks created using the App Engine SDK: the queue-level retry
	//    settings apply to all tasks in the queue which do not have retry settings
	//    explicitly set on the task and were created by the App Engine SDK. See
	//    [App Engine
	//    documentation](https://cloud.google.com/appengine/docs/standard/python/taskqueue/push/retrying-tasks).
	// +kcc:proto:field=google.cloud.tasks.v2.Queue.retry_config
	RetryConfig *RetryConfig `json:"retryConfig,omitempty"`

	// Configuration options for writing logs to
	//  [Stackdriver Logging](https://cloud.google.com/logging/docs/). If this
	//  field is unset, then no logs are written.
	// +kcc:proto:field=google.cloud.tasks.v2.Queue.stackdriver_logging_config
	StackdriverLoggingConfig *StackdriverLoggingConfig `json:"stackdriverLoggingConfig,omitempty"`

	// Required. The location of the queue.
	Location string `json:"location,omitempty"`

	// Required. The host project of the queue.
	ProjectRef v1beta1.ProjectRef `json:"projectRef,omitempty"`

	// The TasksQueue name. If not given, the metadata.name will be used.
	ResourceID *string `json:"resourceID,omitempty"`
}

// TasksQueueStatus defines the config connector machine state of TasksQueue
type TasksQueueStatus struct {
	/* Conditions represent the latest available observations of the
	   object's current state. */
	Conditions []v1alpha1.Condition `json:"conditions,omitempty"`

	// ObservedGeneration is the generation of the resource that was most recently observed by the Config Connector controller. If this is equal to metadata.generation, then that means that the current reported status reflects the most recent desired state of the resource.
	ObservedGeneration *int64 `json:"observedGeneration,omitempty"`

	// A unique specifier for the TasksQueue resource in GCP.
	ExternalRef *string `json:"externalRef,omitempty"`

	// ObservedState is the state of the resource as most recently observed in GCP.
	ObservedState *TasksQueueObservedState `json:"observedState,omitempty"`
}

// TasksQueueObservedState is the state of the TasksQueue resource as most recently observed in GCP.
// +kcc:proto=google.cloud.tasks.v2.Queue
type TasksQueueObservedState struct {

	// Output only. The last time this queue was purged.
	//
	//  All tasks that were [created][google.cloud.tasks.v2.Task.create_time]
	//  before this time were purged.
	//
	//  A queue can be purged using
	//  [PurgeQueue][google.cloud.tasks.v2.CloudTasks.PurgeQueue], the [App Engine
	//  Task Queue SDK, or the Cloud
	//  Console](https://cloud.google.com/appengine/docs/standard/python/taskqueue/push/deleting-tasks-and-queues#purging_all_tasks_from_a_queue).
	//
	//  Purge time will be truncated to the nearest microsecond. Purge
	//  time will be unset if the queue has never been purged.
	// +kcc:proto:field=google.cloud.tasks.v2.Queue.purge_time
	PurgeTime *string `json:"purgeTime,omitempty"`

	// Output only. The state of the queue.
	//
	//  `state` can only be changed by calling
	//  [PauseQueue][google.cloud.tasks.v2.CloudTasks.PauseQueue],
	//  [ResumeQueue][google.cloud.tasks.v2.CloudTasks.ResumeQueue], or uploading
	//  [queue.yaml/xml](https://cloud.google.com/appengine/docs/python/config/queueref).
	//  [UpdateQueue][google.cloud.tasks.v2.CloudTasks.UpdateQueue] cannot be used
	//  to change `state`.
	// +kcc:proto:field=google.cloud.tasks.v2.Queue.state
	State *string `json:"state,omitempty"`
}

// +genclient
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// TODO(user): make sure the pluralizaiton below is correct
// +kubebuilder:resource:categories=gcp,shortName=gcptasksqueue;gcptasksqueues
// +kubebuilder:subresource:status
// +kubebuilder:metadata:labels="cnrm.cloud.google.com/managed-by-kcc=true";"cnrm.cloud.google.com/system=true"
// +kubebuilder:printcolumn:name="Age",JSONPath=".metadata.creationTimestamp",type="date"
// +kubebuilder:printcolumn:name="Ready",JSONPath=".status.conditions[?(@.type=='Ready')].status",type="string",description="When 'True', the most recent reconcile of the resource succeeded"
// +kubebuilder:printcolumn:name="Status",JSONPath=".status.conditions[?(@.type=='Ready')].reason",type="string",description="The reason for the value in 'Ready'"
// +kubebuilder:printcolumn:name="Status Age",JSONPath=".status.conditions[?(@.type=='Ready')].lastTransitionTime",type="date",description="The last transition time for the value in 'Status'"

// TasksQueue is the Schema for the TasksQueue API
// +k8s:openapi-gen=true
type TasksQueue struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	// +required
	Spec   TasksQueueSpec   `json:"spec,omitempty"`
	Status TasksQueueStatus `json:"status,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// TasksQueueList contains a list of TasksQueue
type TasksQueueList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []TasksQueue `json:"items"`
}

func init() {
	SchemeBuilder.Register(&TasksQueue{}, &TasksQueueList{})
}
