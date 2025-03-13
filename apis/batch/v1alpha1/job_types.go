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
	*Parent `json:",inline"`

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

// +kcc:proto=google.cloud.batch.v1.AllocationPolicy
type AllocationPolicy struct {
	// Location where compute resources should be allocated for the Job.
	// +kcc:proto:field=google.cloud.batch.v1.AllocationPolicy.location
	Location *AllocationPolicy_LocationPolicy `json:"location,omitempty"`

	// Describe instances that can be created by this AllocationPolicy.
	//  Only instances[0] is supported now.
	// +kcc:proto:field=google.cloud.batch.v1.AllocationPolicy.instances
	Instances []AllocationPolicy_InstancePolicyOrTemplate `json:"instances,omitempty"`

	// Defines the service account for Batch-created VMs. If omitted, the [default
	//  Compute Engine service
	//  account](https://cloud.google.com/compute/docs/access/service-accounts#default_service_account)
	//  is used. Must match the service account specified in any used instance
	//  template configured in the Batch job.
	//
	//  Includes the following fields:
	//   * email: The service account's email address. If not set, the default
	//   Compute Engine service account is used.
	//   * scopes: Additional OAuth scopes to grant the service account, beyond the
	//   default cloud-platform scope. (list of strings)
	// +kcc:proto:field=google.cloud.batch.v1.AllocationPolicy.service_account
	ServiceAccount *v1beta1.IAMServiceAccountRef `json:"serviceAccount,omitempty"`

	// Custom labels to apply to the job and all the Compute Engine resources
	//  that both are created by this allocation policy and support labels.
	//
	//  Use labels to group and describe the resources they are applied to. Batch
	//  automatically applies predefined labels and supports multiple `labels`
	//  fields for each job, which each let you apply custom labels to various
	//  resources. Label names that start with "goog-" or "google-" are
	//  reserved for predefined labels. For more information about labels with
	//  Batch, see
	//  [Organize resources using
	//  labels](https://cloud.google.com/batch/docs/organize-resources-using-labels).
	// +kcc:proto:field=google.cloud.batch.v1.AllocationPolicy.labels
	Labels map[string]string `json:"labels,omitempty"`

	// The network policy.
	//
	//  If you define an instance template in the `InstancePolicyOrTemplate` field,
	//  Batch will use the network settings in the instance template instead of
	//  this field.
	// +kcc:proto:field=google.cloud.batch.v1.AllocationPolicy.network
	Network *AllocationPolicy_NetworkPolicy `json:"network,omitempty"`

	// The placement policy.
	// +kcc:proto:field=google.cloud.batch.v1.AllocationPolicy.placement
	Placement *AllocationPolicy_PlacementPolicy `json:"placement,omitempty"`

	// Optional. Tags applied to the VM instances.
	//
	//  The tags identify valid sources or targets for network firewalls.
	//  Each tag must be 1-63 characters long, and comply with
	//  [RFC1035](https://www.ietf.org/rfc/rfc1035.txt).
	// +kcc:proto:field=google.cloud.batch.v1.AllocationPolicy.tags
	Tags []string `json:"tags,omitempty"`
}

// +kcc:proto=google.cloud.batch.v1.ComputeResource
type ComputeResource struct {
	// The milliCPU count.
	//
	//  `cpuMilli` defines the amount of CPU resources per task in milliCPU units.
	//  For example, `1000` corresponds to 1 vCPU per task. If undefined, the
	//  default value is `2000`.
	//
	//  If you also define the VM's machine type using the `machineType` in
	//  [InstancePolicy](https://cloud.google.com/batch/docs/reference/rest/v1/projects.locations.jobs#instancepolicy)
	//  field or inside the `instanceTemplate` in the
	//  [InstancePolicyOrTemplate](https://cloud.google.com/batch/docs/reference/rest/v1/projects.locations.jobs#instancepolicyortemplate)
	//  field, make sure the CPU resources for both fields are compatible with each
	//  other and with how many tasks you want to allow to run on the same VM at
	//  the same time.
	//
	//  For example, if you specify the `n2-standard-2` machine type, which has 2
	//  vCPUs each, you are recommended to set `cpuMilli` no more than `2000`, or
	//  you are recommended to run two tasks on the same VM if you set `cpuMilli`
	//  to `1000` or less.
	// +kcc:proto:field=google.cloud.batch.v1.ComputeResource.cpu_milli
	CPUMilli *int64 `json:"cpuMilli,omitempty"`

	// Memory in MiB.
	//
	//  `memoryMib` defines the amount of memory per task in MiB units.
	//  If undefined, the default value is `2000`.
	//  If you also define the VM's machine type using the `machineType` in
	//  [InstancePolicy](https://cloud.google.com/batch/docs/reference/rest/v1/projects.locations.jobs#instancepolicy)
	//  field or inside the `instanceTemplate` in the
	//  [InstancePolicyOrTemplate](https://cloud.google.com/batch/docs/reference/rest/v1/projects.locations.jobs#instancepolicyortemplate)
	//  field, make sure the memory resources for both fields are compatible with
	//  each other and with how many tasks you want to allow to run on the same VM
	//  at the same time.
	//
	//  For example, if you specify the `n2-standard-2` machine type, which has 8
	//  GiB each, you are recommended to set `memoryMib` to no more than `8192`,
	//  or you are recommended to run two tasks on the same VM if you set
	//  `memoryMib` to `4096` or less.
	// +kcc:proto:field=google.cloud.batch.v1.ComputeResource.memory_mib
	MemoryMiB *int64 `json:"memoryMib,omitempty"`

	// Extra boot disk size in MiB for each task.
	// +kcc:proto:field=google.cloud.batch.v1.ComputeResource.boot_disk_mib
	BootDiskMiB *int64 `json:"bootDiskMib,omitempty"`
}

// +kcc:proto=google.cloud.batch.v1.Environment.KMSEnvMap
type Environment_KMSEnvMap struct {
	// The name of the KMS key that will be used to decrypt the cipher text.
	// +kcc:proto:field=google.cloud.batch.v1.Environment.KMSEnvMap.key_name
	KMSKeyRef *v1beta1.KMSCryptoKeyRef `json:"kmsKeyRef,omitempty"`

	// The value of the cipherText response from the `encrypt` method.
	// +kcc:proto:field=google.cloud.batch.v1.Environment.KMSEnvMap.cipher_text
	CipherText *string `json:"cipherText,omitempty"`
}

// +kcc:proto=google.cloud.batch.v1.JobNotification
type JobNotification struct {
	// The Pub/Sub topic where notifications for the job, like state
	//  changes, will be published. If undefined, no Pub/Sub notifications
	//  are sent for this job.
	//
	//  Specify the topic using the following format:
	//  `projects/{project}/topics/{topic}`.
	//  Notably, if you want to specify a Pub/Sub topic that is in a
	//  different project than the job, your administrator must grant your
	//  project's Batch service agent permission to publish to that topic.
	//
	//  For more information about configuring Pub/Sub notifications for
	//  a job, see
	//  https://cloud.google.com/batch/docs/enable-notifications.
	// +kcc:proto:field=google.cloud.batch.v1.JobNotification.pubsub_topic
	PubsubTopicRef *v1beta1.PubSubTopicRef `json:"pubsubTopicRef,omitempty"`

	// The attribute requirements of messages to be sent to this Pub/Sub topic.
	//  Without this field, no message will be sent.
	// +kcc:proto:field=google.cloud.batch.v1.JobNotification.message
	Message *JobNotification_Message `json:"message,omitempty"`
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
