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
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/clients/generated/apis/k8s/v1alpha1"

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
	ServiceAccountRef *v1beta1.IAMServiceAccountRef `json:"serviceAccount,omitempty"`

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

// +kcc:proto=google.cloud.batch.v1.AllocationPolicy.Disk
type AllocationPolicy_Disk struct {
	// URL for a VM image to use as the data source for this disk.
	//  For example, the following are all valid URLs:
	//
	//  * Specify the image by its family name:
	//  projects/{project}/global/images/family/{image_family}
	//  * Specify the image version:
	//  projects/{project}/global/images/{image_version}
	//
	//  You can also use Batch customized image in short names.
	//  The following image values are supported for a boot disk:
	//
	//  * `batch-debian`: use Batch Debian images.
	//  * `batch-cos`: use Batch Container-Optimized images.
	//  * `batch-hpc-rocky`: use Batch HPC Rocky Linux images.
	// +kcc:proto:field=google.cloud.batch.v1.AllocationPolicy.Disk.image
	ImageRef *v1alpha1.ResourceRef `json:"imageRef,omitempty"`

	// Name of a snapshot used as the data source.
	//  Snapshot is not supported as boot disk now.
	// +kcc:proto:field=google.cloud.batch.v1.AllocationPolicy.Disk.snapshot
	Snapshot *string `json:"snapshot,omitempty"`

	// Disk type as shown in `gcloud compute disk-types list`.
	//  For example, local SSD uses type "local-ssd".
	//  Persistent disks and boot disks use "pd-balanced", "pd-extreme", "pd-ssd"
	//  or "pd-standard". If not specified, "pd-standard" will be used as the
	//  default type for non-boot disks, "pd-balanced" will be used as the
	//  default type for boot disks.
	// +kcc:proto:field=google.cloud.batch.v1.AllocationPolicy.Disk.type
	Type *string `json:"type,omitempty"`

	// Disk size in GB.
	//
	//  **Non-Boot Disk**:
	//  If the `type` specifies a persistent disk, this field
	//  is ignored if `data_source` is set as `image` or `snapshot`.
	//  If the `type` specifies a local SSD, this field should be a multiple of
	//  375 GB, otherwise, the final size will be the next greater multiple of
	//  375 GB.
	//
	//  **Boot Disk**:
	//  Batch will calculate the boot disk size based on source
	//  image and task requirements if you do not speicify the size.
	//  If both this field and the `boot_disk_mib` field in task spec's
	//  `compute_resource` are defined, Batch will only honor this field.
	//  Also, this field should be no smaller than the source disk's
	//  size when the `data_source` is set as `snapshot` or `image`.
	//  For example, if you set an image as the `data_source` field and the
	//  image's default disk size 30 GB, you can only use this field to make the
	//  disk larger or equal to 30 GB.
	// +kcc:proto:field=google.cloud.batch.v1.AllocationPolicy.Disk.size_gb
	SizeGB *int64 `json:"sizeGB,omitempty"`

	// Local SSDs are available through both "SCSI" and "NVMe" interfaces.
	//  If not indicated, "NVMe" will be the default one for local ssds.
	//  This field is ignored for persistent disks as the interface is chosen
	//  automatically. See
	//  https://cloud.google.com/compute/docs/disks/persistent-disks#choose_an_interface.
	// +kcc:proto:field=google.cloud.batch.v1.AllocationPolicy.Disk.disk_interface
	DiskInterface *string `json:"diskInterface,omitempty"`
}

// +kcc:proto=google.cloud.batch.v1.AllocationPolicy.NetworkInterface
type AllocationPolicy_NetworkInterface struct {
	// The URL of an existing network resource.
	//  You can specify the network as a full or partial URL.
	//
	//  For example, the following are all valid URLs:
	//
	//  * https://www.googleapis.com/compute/v1/projects/{project}/global/networks/{network}
	//  * projects/{project}/global/networks/{network}
	//  * global/networks/{network}
	// +kcc:proto:field=google.cloud.batch.v1.AllocationPolicy.NetworkInterface.network
	NetworkRef *v1beta1.ComputeNetworkRef `json:"networkRef,omitempty"`

	// The URL of an existing subnetwork resource in the network.
	//  You can specify the subnetwork as a full or partial URL.
	//
	//  For example, the following are all valid URLs:
	//
	//  * https://www.googleapis.com/compute/v1/projects/{project}/regions/{region}/subnetworks/{subnetwork}
	//  * projects/{project}/regions/{region}/subnetworks/{subnetwork}
	//  * regions/{region}/subnetworks/{subnetwork}
	// +kcc:proto:field=google.cloud.batch.v1.AllocationPolicy.NetworkInterface.subnetwork
	SubnetworkRef *v1beta1.ComputeSubnetworkRef `json:"subnetworkRef,omitempty"`

	// Default is false (with an external IP address). Required if
	//  no external public IP address is attached to the VM. If no external
	//  public IP address, additional configuration is required to allow the VM
	//  to access Google Services. See
	//  https://cloud.google.com/vpc/docs/configure-private-google-access and
	//  https://cloud.google.com/nat/docs/gce-example#create-nat for more
	//  information.
	// +kcc:proto:field=google.cloud.batch.v1.AllocationPolicy.NetworkInterface.no_external_ip_address
	NoExternalIPAddress *bool `json:"noExternalIPAddress,omitempty"`
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
	MemoryMiB *int64 `json:"memoryMiB,omitempty"`

	// Extra boot disk size in MiB for each task.
	// +kcc:proto:field=google.cloud.batch.v1.ComputeResource.boot_disk_mib
	BootDiskMiB *int64 `json:"bootDiskMiB,omitempty"`
}

// +kcc:proto=google.cloud.batch.v1.Runnable.Container
type Runnable_Container struct {
	// Required. The URI to pull the container image from.
	// +kcc:proto:field=google.cloud.batch.v1.Runnable.Container.image_uri
	ImageURI *string `json:"imageURI,omitempty"`

	// Required for some container images. Overrides the `CMD` specified in the
	//  container. If there is an `ENTRYPOINT` (either in the container image or
	//  with the `entrypoint` field below) then these commands are appended as
	//  arguments to the `ENTRYPOINT`.
	// +kcc:proto:field=google.cloud.batch.v1.Runnable.Container.commands
	Commands []string `json:"commands,omitempty"`

	// Required for some container images. Overrides the `ENTRYPOINT` specified
	//  in the container.
	// +kcc:proto:field=google.cloud.batch.v1.Runnable.Container.entrypoint
	Entrypoint *string `json:"entrypoint,omitempty"`

	// Volumes to mount (bind mount) from the host machine files or directories
	//  into the container, formatted to match `--volume` option for the
	//  `docker run` command&mdash;for example, `/foo:/bar` or `/foo:/bar:ro`.
	//
	//  If the `TaskSpec.Volumes` field is specified but this field is not, Batch
	//  will mount each volume from the host machine to the container with the
	//  same mount path by default. In this case, the default mount option for
	//  containers will be read-only (`ro`) for existing persistent disks and
	//  read-write (`rw`) for other volume types, regardless of the original
	//  mount options specified in `TaskSpec.Volumes`. If you need different
	//  mount settings, you can explicitly configure them in this field.
	// +kcc:proto:field=google.cloud.batch.v1.Runnable.Container.volumes
	Volumes []string `json:"volumes,omitempty"`

	// Required for some container images. Arbitrary additional options to
	//  include in the `docker run` command when running this container&mdash;for
	//  example, `--network host`. For the `--volume` option, use the `volumes`
	//  field for the container.
	// +kcc:proto:field=google.cloud.batch.v1.Runnable.Container.options
	Options *string `json:"options,omitempty"`

	// If set to true, external network access to and from container will be
	//  blocked, containers that are with block_external_network as true can
	//  still communicate with each other, network cannot be specified in the
	//  `container.options` field.
	// +kcc:proto:field=google.cloud.batch.v1.Runnable.Container.block_external_network
	BlockExternalNetwork *bool `json:"blockExternalNetwork,omitempty"`

	// Required if the container image is from a private Docker registry. The
	//  username to login to the Docker registry that contains the image.
	//
	//  You can either specify the username directly by using plain text or
	//  specify an encrypted username by using a Secret Manager secret:
	//  `projects/*/secrets/*/versions/*`. However, using a secret is
	//  recommended for enhanced security.
	//
	//  Caution: If you specify the username using plain text, you risk the
	//  username being exposed to any users who can view the job or its logs.
	//  To avoid this risk, specify a secret that contains the username instead.
	//
	//  Learn more about [Secret
	//  Manager](https://cloud.google.com/secret-manager/docs/) and [using
	//  Secret Manager with
	//  Batch](https://cloud.google.com/batch/docs/create-run-job-secret-manager).
	// +kcc:proto:field=google.cloud.batch.v1.Runnable.Container.username
	Username *string `json:"username,omitempty"`

	// Required if the container image is from a private Docker registry. The
	//  password to login to the Docker registry that contains the image.
	//
	//  For security, it is strongly recommended to specify an
	//  encrypted password by using a Secret Manager secret:
	//  `projects/*/secrets/*/versions/*`.
	//
	//  Warning: If you specify the password using plain text, you risk the
	//  password being exposed to any users who can view the job or its logs.
	//  To avoid this risk, specify a secret that contains the password instead.
	//
	//  Learn more about [Secret
	//  Manager](https://cloud.google.com/secret-manager/docs/) and [using
	//  Secret Manager with
	//  Batch](https://cloud.google.com/batch/docs/create-run-job-secret-manager).
	// +kcc:proto:field=google.cloud.batch.v1.Runnable.Container.password

	//	TODO: comment out the password for now and add it back when secreteManager
	//  and k8s secret is ready before promoting to beta.
	// SecretRef *refsv1beta1secret.BasicAuthSecretRef `json:"secretRef,omitempty"`

	// Optional. If set to true, this container runnable uses Image streaming.
	//
	//  Use Image streaming to allow the runnable to initialize without
	//  waiting for the entire container image to download, which can
	//  significantly reduce startup time for large container images.
	//
	//  When `enableImageStreaming` is set to true, the container
	//  runtime is [containerd](https://containerd.io/) instead of Docker.
	//  Additionally, this container runnable only supports the following
	//  `container` subfields: `imageUri`,
	//  `commands[]`, `entrypoint`, and
	//  `volumes[]`; any other `container` subfields are ignored.
	//
	//  For more information about the requirements and limitations for using
	//  Image streaming with Batch, see the [`image-streaming`
	//  sample on
	//  GitHub](https://github.com/GoogleCloudPlatform/batch-samples/tree/main/api-samples/image-streaming).
	// +kcc:proto:field=google.cloud.batch.v1.Runnable.Container.enable_image_streaming
	EnableImageStreaming *bool `json:"enableImageStreaming,omitempty"`
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
