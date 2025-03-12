// Copyright 2025 Google LLC
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package v1alpha1

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
	ServiceAccount *ServiceAccount `json:"serviceAccount,omitempty"`

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

// +kcc:proto=google.cloud.batch.v1.AllocationPolicy.Accelerator
type AllocationPolicy_Accelerator struct {
	// The accelerator type. For example, "nvidia-tesla-t4".
	//  See `gcloud compute accelerator-types list`.
	// +kcc:proto:field=google.cloud.batch.v1.AllocationPolicy.Accelerator.type
	Type *string `json:"type,omitempty"`

	// The number of accelerators of this type.
	// +kcc:proto:field=google.cloud.batch.v1.AllocationPolicy.Accelerator.count
	Count *int64 `json:"count,omitempty"`

	// Deprecated: please use instances[0].install_gpu_drivers instead.
	// +kcc:proto:field=google.cloud.batch.v1.AllocationPolicy.Accelerator.install_gpu_drivers
	InstallGpuDrivers *bool `json:"installGpuDrivers,omitempty"`

	// Optional. The NVIDIA GPU driver version that should be installed for this
	//  type.
	//
	//  You can define the specific driver version such as "470.103.01",
	//  following the driver version requirements in
	//  https://cloud.google.com/compute/docs/gpus/install-drivers-gpu#minimum-driver.
	//  Batch will install the specific accelerator driver if qualified.
	// +kcc:proto:field=google.cloud.batch.v1.AllocationPolicy.Accelerator.driver_version
	DriverVersion *string `json:"driverVersion,omitempty"`
}

// +kcc:proto=google.cloud.batch.v1.AllocationPolicy.AttachedDisk
type AllocationPolicy_AttachedDisk struct {
	// +kcc:proto:field=google.cloud.batch.v1.AllocationPolicy.AttachedDisk.new_disk
	NewDisk *AllocationPolicy_Disk `json:"newDisk,omitempty"`

	// Name of an existing PD.
	// +kcc:proto:field=google.cloud.batch.v1.AllocationPolicy.AttachedDisk.existing_disk
	ExistingDisk *string `json:"existingDisk,omitempty"`

	// Device name that the guest operating system will see.
	//  It is used by Runnable.volumes field to mount disks. So please specify
	//  the device_name if you want Batch to help mount the disk, and it should
	//  match the device_name field in volumes.
	// +kcc:proto:field=google.cloud.batch.v1.AllocationPolicy.AttachedDisk.device_name
	DeviceName *string `json:"deviceName,omitempty"`
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
	Image *string `json:"image,omitempty"`

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

// +kcc:proto=google.cloud.batch.v1.AllocationPolicy.InstancePolicy
type AllocationPolicy_InstancePolicy struct {
	// The Compute Engine machine type.
	// +kcc:proto:field=google.cloud.batch.v1.AllocationPolicy.InstancePolicy.machine_type
	MachineType *string `json:"machineType,omitempty"`

	// The minimum CPU platform.
	//  See
	//  https://cloud.google.com/compute/docs/instances/specify-min-cpu-platform.
	// +kcc:proto:field=google.cloud.batch.v1.AllocationPolicy.InstancePolicy.min_cpu_platform
	MinCPUPlatform *string `json:"minCPUPlatform,omitempty"`

	// The provisioning model.
	// +kcc:proto:field=google.cloud.batch.v1.AllocationPolicy.InstancePolicy.provisioning_model
	ProvisioningModel *string `json:"provisioningModel,omitempty"`

	// The accelerators attached to each VM instance.
	// +kcc:proto:field=google.cloud.batch.v1.AllocationPolicy.InstancePolicy.accelerators
	Accelerators []AllocationPolicy_Accelerator `json:"accelerators,omitempty"`

	// Boot disk to be created and attached to each VM by this InstancePolicy.
	//  Boot disk will be deleted when the VM is deleted.
	//  Batch API now only supports booting from image.
	// +kcc:proto:field=google.cloud.batch.v1.AllocationPolicy.InstancePolicy.boot_disk
	BootDisk *AllocationPolicy_Disk `json:"bootDisk,omitempty"`

	// Non-boot disks to be attached for each VM created by this InstancePolicy.
	//  New disks will be deleted when the VM is deleted.
	//  A non-boot disk is a disk that can be of a device with a
	//  file system or a raw storage drive that is not ready for data
	//  storage and accessing.
	// +kcc:proto:field=google.cloud.batch.v1.AllocationPolicy.InstancePolicy.disks
	Disks []AllocationPolicy_AttachedDisk `json:"disks,omitempty"`

	// Optional. If not specified (default), VMs will consume any applicable
	//  reservation. If "NO_RESERVATION" is specified, VMs will not consume any
	//  reservation. Otherwise, if specified, VMs will consume only the specified
	//  reservation.
	// +kcc:proto:field=google.cloud.batch.v1.AllocationPolicy.InstancePolicy.reservation
	Reservation *string `json:"reservation,omitempty"`
}

// +kcc:proto=google.cloud.batch.v1.AllocationPolicy.InstancePolicyOrTemplate
type AllocationPolicy_InstancePolicyOrTemplate struct {
	// InstancePolicy.
	// +kcc:proto:field=google.cloud.batch.v1.AllocationPolicy.InstancePolicyOrTemplate.policy
	Policy *AllocationPolicy_InstancePolicy `json:"policy,omitempty"`

	// Name of an instance template used to create VMs.
	//  Named the field as 'instance_template' instead of 'template' to avoid
	//  C++ keyword conflict.
	//
	//  Batch only supports global instance templates from the same project as
	//  the job.
	//  You can specify the global instance template as a full or partial URL.
	// +kcc:proto:field=google.cloud.batch.v1.AllocationPolicy.InstancePolicyOrTemplate.instance_template
	InstanceTemplate *string `json:"instanceTemplate,omitempty"`

	// Set this field true if you want Batch to help fetch drivers from a third
	//  party location and install them for GPUs specified in
	//  `policy.accelerators` or `instance_template` on your behalf. Default is
	//  false.
	//
	//  For Container-Optimized Image cases, Batch will install the
	//  accelerator driver following milestones of
	//  https://cloud.google.com/container-optimized-os/docs/release-notes. For
	//  non Container-Optimized Image cases, following
	//  https://github.com/GoogleCloudPlatform/compute-gpu-installation/blob/main/linux/install_gpu_driver.py.
	// +kcc:proto:field=google.cloud.batch.v1.AllocationPolicy.InstancePolicyOrTemplate.install_gpu_drivers
	InstallGpuDrivers *bool `json:"installGpuDrivers,omitempty"`

	// Optional. Set this field true if you want Batch to install Ops Agent on
	//  your behalf. Default is false.
	// +kcc:proto:field=google.cloud.batch.v1.AllocationPolicy.InstancePolicyOrTemplate.install_ops_agent
	InstallOpsAgent *bool `json:"installOpsAgent,omitempty"`

	// Optional. Set this field to `true` if you want Batch to block
	//  project-level SSH keys from accessing this job's VMs.  Alternatively, you
	//  can configure the job to specify a VM instance template that blocks
	//  project-level SSH keys. In either case, Batch blocks project-level SSH
	//  keys while creating the VMs for this job.
	//
	//  Batch allows project-level SSH keys for a job's VMs only if all
	//  the following are true:
	//
	//  + This field is undefined or set to `false`.
	//  + The job's VM instance template (if any) doesn't block project-level
	//    SSH keys.
	//
	//  Notably, you can override this behavior by manually updating a VM to
	//  block or allow project-level SSH keys. For more information about
	//  blocking project-level SSH keys, see the Compute Engine documentation:
	//  https://cloud.google.com/compute/docs/connect/restrict-ssh-keys#block-keys
	// +kcc:proto:field=google.cloud.batch.v1.AllocationPolicy.InstancePolicyOrTemplate.block_project_ssh_keys
	BlockProjectSSHKeys *bool `json:"blockProjectSSHKeys,omitempty"`
}

// +kcc:proto=google.cloud.batch.v1.AllocationPolicy.LocationPolicy
type AllocationPolicy_LocationPolicy struct {
	// A list of allowed location names represented by internal URLs.
	//
	//  Each location can be a region or a zone.
	//  Only one region or multiple zones in one region is supported now.
	//  For example,
	//  ["regions/us-central1"] allow VMs in any zones in region us-central1.
	//  ["zones/us-central1-a", "zones/us-central1-c"] only allow VMs
	//  in zones us-central1-a and us-central1-c.
	//
	//  Mixing locations from different regions would cause errors.
	//  For example,
	//  ["regions/us-central1", "zones/us-central1-a", "zones/us-central1-b",
	//  "zones/us-west1-a"] contains locations from two distinct regions:
	//  us-central1 and us-west1. This combination will trigger an error.
	// +kcc:proto:field=google.cloud.batch.v1.AllocationPolicy.LocationPolicy.allowed_locations
	AllowedLocations []string `json:"allowedLocations,omitempty"`
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
	Network *string `json:"network,omitempty"`

	// The URL of an existing subnetwork resource in the network.
	//  You can specify the subnetwork as a full or partial URL.
	//
	//  For example, the following are all valid URLs:
	//
	//  * https://www.googleapis.com/compute/v1/projects/{project}/regions/{region}/subnetworks/{subnetwork}
	//  * projects/{project}/regions/{region}/subnetworks/{subnetwork}
	//  * regions/{region}/subnetworks/{subnetwork}
	// +kcc:proto:field=google.cloud.batch.v1.AllocationPolicy.NetworkInterface.subnetwork
	Subnetwork *string `json:"subnetwork,omitempty"`

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

// +kcc:proto=google.cloud.batch.v1.AllocationPolicy.NetworkPolicy
type AllocationPolicy_NetworkPolicy struct {
	// Network configurations.
	// +kcc:proto:field=google.cloud.batch.v1.AllocationPolicy.NetworkPolicy.network_interfaces
	NetworkInterfaces []AllocationPolicy_NetworkInterface `json:"networkInterfaces,omitempty"`
}

// +kcc:proto=google.cloud.batch.v1.AllocationPolicy.PlacementPolicy
type AllocationPolicy_PlacementPolicy struct {
	// UNSPECIFIED vs. COLLOCATED (default UNSPECIFIED). Use COLLOCATED when you
	//  want VMs to be located close to each other for low network latency
	//  between the VMs. No placement policy will be generated when collocation
	//  is UNSPECIFIED.
	// +kcc:proto:field=google.cloud.batch.v1.AllocationPolicy.PlacementPolicy.collocation
	Collocation *string `json:"collocation,omitempty"`

	// When specified, causes the job to fail if more than max_distance logical
	//  switches are required between VMs. Batch uses the most compact possible
	//  placement of VMs even when max_distance is not specified. An explicit
	//  max_distance makes that level of compactness a strict requirement.
	//  Not yet implemented
	// +kcc:proto:field=google.cloud.batch.v1.AllocationPolicy.PlacementPolicy.max_distance
	MaxDistance *int64 `json:"maxDistance,omitempty"`
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
	MemoryMib *int64 `json:"memoryMib,omitempty"`

	// Extra boot disk size in MiB for each task.
	// +kcc:proto:field=google.cloud.batch.v1.ComputeResource.boot_disk_mib
	BootDiskMib *int64 `json:"bootDiskMib,omitempty"`
}

// +kcc:proto=google.cloud.batch.v1.Environment
type Environment struct {
	// A map of environment variable names to values.
	// +kcc:proto:field=google.cloud.batch.v1.Environment.variables
	Variables map[string]string `json:"variables,omitempty"`

	// A map of environment variable names to Secret Manager secret names.
	//  The VM will access the named secrets to set the value of each environment
	//  variable.
	// +kcc:proto:field=google.cloud.batch.v1.Environment.secret_variables
	SecretVariables map[string]string `json:"secretVariables,omitempty"`

	// An encrypted JSON dictionary where the key/value pairs correspond to
	//  environment variable names and their values.
	// +kcc:proto:field=google.cloud.batch.v1.Environment.encrypted_variables
	EncryptedVariables *Environment_KMSEnvMap `json:"encryptedVariables,omitempty"`
}

// +kcc:proto=google.cloud.batch.v1.Environment.KMSEnvMap
type Environment_KMSEnvMap struct {
	// The name of the KMS key that will be used to decrypt the cipher text.
	// +kcc:proto:field=google.cloud.batch.v1.Environment.KMSEnvMap.key_name
	KeyName *string `json:"keyName,omitempty"`

	// The value of the cipherText response from the `encrypt` method.
	// +kcc:proto:field=google.cloud.batch.v1.Environment.KMSEnvMap.cipher_text
	CipherText *string `json:"cipherText,omitempty"`
}

// +kcc:proto=google.cloud.batch.v1.GCS
type Gcs struct {
	// Remote path, either a bucket name or a subdirectory of a bucket, e.g.:
	//  bucket_name, bucket_name/subdirectory/
	// +kcc:proto:field=google.cloud.batch.v1.GCS.remote_path
	RemotePath *string `json:"remotePath,omitempty"`
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
	PubsubTopic *string `json:"pubsubTopic,omitempty"`

	// The attribute requirements of messages to be sent to this Pub/Sub topic.
	//  Without this field, no message will be sent.
	// +kcc:proto:field=google.cloud.batch.v1.JobNotification.message
	Message *JobNotification_Message `json:"message,omitempty"`
}

// +kcc:proto=google.cloud.batch.v1.JobNotification.Message
type JobNotification_Message struct {
	// The message type.
	// +kcc:proto:field=google.cloud.batch.v1.JobNotification.Message.type
	Type *string `json:"type,omitempty"`

	// The new job state.
	// +kcc:proto:field=google.cloud.batch.v1.JobNotification.Message.new_job_state
	NewJobState *string `json:"newJobState,omitempty"`

	// The new task state.
	// +kcc:proto:field=google.cloud.batch.v1.JobNotification.Message.new_task_state
	NewTaskState *string `json:"newTaskState,omitempty"`
}

// +kcc:proto=google.cloud.batch.v1.JobStatus
type JobStatus struct {
	// Job state
	// +kcc:proto:field=google.cloud.batch.v1.JobStatus.state
	State *string `json:"state,omitempty"`

	// Job status events
	// +kcc:proto:field=google.cloud.batch.v1.JobStatus.status_events
	StatusEvents []StatusEvent `json:"statusEvents,omitempty"`

	// TODO: unsupported map type with key string and value message

	// The duration of time that the Job spent in status RUNNING.
	// +kcc:proto:field=google.cloud.batch.v1.JobStatus.run_duration
	RunDuration *string `json:"runDuration,omitempty"`
}

// +kcc:proto=google.cloud.batch.v1.JobStatus.InstanceStatus
type JobStatus_InstanceStatus struct {
	// The Compute Engine machine type.
	// +kcc:proto:field=google.cloud.batch.v1.JobStatus.InstanceStatus.machine_type
	MachineType *string `json:"machineType,omitempty"`

	// The VM instance provisioning model.
	// +kcc:proto:field=google.cloud.batch.v1.JobStatus.InstanceStatus.provisioning_model
	ProvisioningModel *string `json:"provisioningModel,omitempty"`

	// The max number of tasks can be assigned to this instance type.
	// +kcc:proto:field=google.cloud.batch.v1.JobStatus.InstanceStatus.task_pack
	TaskPack *int64 `json:"taskPack,omitempty"`

	// The VM boot disk.
	// +kcc:proto:field=google.cloud.batch.v1.JobStatus.InstanceStatus.boot_disk
	BootDisk *AllocationPolicy_Disk `json:"bootDisk,omitempty"`
}

// +kcc:proto=google.cloud.batch.v1.JobStatus.TaskGroupStatus
type JobStatus_TaskGroupStatus struct {
	// Count of task in each state in the TaskGroup.
	//  The map key is task state name.
	// +kcc:proto:field=google.cloud.batch.v1.JobStatus.TaskGroupStatus.counts
	Counts map[string]int64 `json:"counts,omitempty"`

	// Status of instances allocated for the TaskGroup.
	// +kcc:proto:field=google.cloud.batch.v1.JobStatus.TaskGroupStatus.instances
	Instances []JobStatus_InstanceStatus `json:"instances,omitempty"`
}

// +kcc:proto=google.cloud.batch.v1.LifecyclePolicy
type LifecyclePolicy struct {
	// Action to execute when ActionCondition is true.
	//  When RETRY_TASK is specified, we will retry failed tasks
	//  if we notice any exit code match and fail tasks if no match is found.
	//  Likewise, when FAIL_TASK is specified, we will fail tasks
	//  if we notice any exit code match and retry tasks if no match is found.
	// +kcc:proto:field=google.cloud.batch.v1.LifecyclePolicy.action
	Action *string `json:"action,omitempty"`

	// Conditions that decide why a task failure is dealt with a specific action.
	// +kcc:proto:field=google.cloud.batch.v1.LifecyclePolicy.action_condition
	ActionCondition *LifecyclePolicy_ActionCondition `json:"actionCondition,omitempty"`
}

// +kcc:proto=google.cloud.batch.v1.LifecyclePolicy.ActionCondition
type LifecyclePolicy_ActionCondition struct {
	// Exit codes of a task execution.
	//  If there are more than 1 exit codes,
	//  when task executes with any of the exit code in the list,
	//  the condition is met and the action will be executed.
	// +kcc:proto:field=google.cloud.batch.v1.LifecyclePolicy.ActionCondition.exit_codes
	ExitCodes []int32 `json:"exitCodes,omitempty"`
}

// +kcc:proto=google.cloud.batch.v1.LogsPolicy
type LogsPolicy struct {
	// If and where logs should be saved.
	// +kcc:proto:field=google.cloud.batch.v1.LogsPolicy.destination
	Destination *string `json:"destination,omitempty"`

	// When `destination` is set to `PATH`, you must set this field to the path
	//  where you want logs to be saved. This path can point to a local directory
	//  on the VM or (if congifured) a directory under the mount path of any
	//  Cloud Storage bucket, network file system (NFS), or writable persistent
	//  disk that is mounted to the job. For example, if the job has a bucket with
	//  `mountPath` set to `/mnt/disks/my-bucket`, you can write logs to the
	//  root directory of the `remotePath` of that bucket by setting this field to
	//  `/mnt/disks/my-bucket/`.
	// +kcc:proto:field=google.cloud.batch.v1.LogsPolicy.logs_path
	LogsPath *string `json:"logsPath,omitempty"`

	// Optional. When `destination` is set to `CLOUD_LOGGING`, you can optionally
	//  set this field to configure additional settings for Cloud Logging.
	// +kcc:proto:field=google.cloud.batch.v1.LogsPolicy.cloud_logging_option
	CloudLoggingOption *LogsPolicy_CloudLoggingOption `json:"cloudLoggingOption,omitempty"`
}

// +kcc:proto=google.cloud.batch.v1.LogsPolicy.CloudLoggingOption
type LogsPolicy_CloudLoggingOption struct {
	// Optional. Set this field to `true` to change the [monitored resource
	//  type](https://cloud.google.com/monitoring/api/resources) for
	//  Cloud Logging logs generated by this Batch job from
	//  the
	//  [`batch.googleapis.com/Job`](https://cloud.google.com/monitoring/api/resources#tag_batch.googleapis.com/Job)
	//  type to the formerly used
	//  [`generic_task`](https://cloud.google.com/monitoring/api/resources#tag_generic_task)
	//  type.
	// +kcc:proto:field=google.cloud.batch.v1.LogsPolicy.CloudLoggingOption.use_generic_task_monitored_resource
	UseGenericTaskMonitoredResource *bool `json:"useGenericTaskMonitoredResource,omitempty"`
}

// +kcc:proto=google.cloud.batch.v1.NFS
type Nfs struct {
	// The IP address of the NFS.
	// +kcc:proto:field=google.cloud.batch.v1.NFS.server
	Server *string `json:"server,omitempty"`

	// Remote source path exported from the NFS, e.g., "/share".
	// +kcc:proto:field=google.cloud.batch.v1.NFS.remote_path
	RemotePath *string `json:"remotePath,omitempty"`
}

// +kcc:proto=google.cloud.batch.v1.Runnable
type Runnable struct {
	// Container runnable.
	// +kcc:proto:field=google.cloud.batch.v1.Runnable.container
	Container *Runnable_Container `json:"container,omitempty"`

	// Script runnable.
	// +kcc:proto:field=google.cloud.batch.v1.Runnable.script
	Script *Runnable_Script `json:"script,omitempty"`

	// Barrier runnable.
	// +kcc:proto:field=google.cloud.batch.v1.Runnable.barrier
	Barrier *Runnable_Barrier `json:"barrier,omitempty"`

	// Optional. DisplayName is an optional field that can be provided by the
	//  caller. If provided, it will be used in logs and other outputs to identify
	//  the script, making it easier for users to understand the logs. If not
	//  provided the index of the runnable will be used for outputs.
	// +kcc:proto:field=google.cloud.batch.v1.Runnable.display_name
	DisplayName *string `json:"displayName,omitempty"`

	// Normally, a runnable that returns a non-zero exit status fails and causes
	//  the task to fail. However, you can set this field to `true` to allow the
	//  task to continue executing its other runnables even if this runnable
	//  fails.
	// +kcc:proto:field=google.cloud.batch.v1.Runnable.ignore_exit_status
	IgnoreExitStatus *bool `json:"ignoreExitStatus,omitempty"`

	// Normally, a runnable that doesn't exit causes its task to fail. However,
	//  you can set this field to `true` to configure a background runnable.
	//  Background runnables are allowed continue running in the background while
	//  the task executes subsequent runnables. For example, background runnables
	//  are useful for providing services to other runnables or providing
	//  debugging-support tools like SSH servers.
	//
	//  Specifically, background runnables are killed automatically (if they have
	//  not already exited) a short time after all foreground runnables have
	//  completed. Even though this is likely to result in a non-zero exit status
	//  for the background runnable, these automatic kills are not treated as task
	//  failures.
	// +kcc:proto:field=google.cloud.batch.v1.Runnable.background
	Background *bool `json:"background,omitempty"`

	// By default, after a Runnable fails, no further Runnable are executed. This
	//  flag indicates that this Runnable must be run even if the Task has already
	//  failed. This is useful for Runnables that copy output files off of the VM
	//  or for debugging.
	//
	//  The always_run flag does not override the Task's overall max_run_duration.
	//  If the max_run_duration has expired then no further Runnables will execute,
	//  not even always_run Runnables.
	// +kcc:proto:field=google.cloud.batch.v1.Runnable.always_run
	AlwaysRun *bool `json:"alwaysRun,omitempty"`

	// Environment variables for this Runnable (overrides variables set for the
	//  whole Task or TaskGroup).
	// +kcc:proto:field=google.cloud.batch.v1.Runnable.environment
	Environment *Environment `json:"environment,omitempty"`

	// Timeout for this Runnable.
	// +kcc:proto:field=google.cloud.batch.v1.Runnable.timeout
	Timeout *string `json:"timeout,omitempty"`

	// Labels for this Runnable.
	// +kcc:proto:field=google.cloud.batch.v1.Runnable.labels
	Labels map[string]string `json:"labels,omitempty"`
}

// +kcc:proto=google.cloud.batch.v1.Runnable.Barrier
type Runnable_Barrier struct {
	// Barriers are identified by their index in runnable list.
	//  Names are not required, but if present should be an identifier.
	// +kcc:proto:field=google.cloud.batch.v1.Runnable.Barrier.name
	Name *string `json:"name,omitempty"`
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
	Password *string `json:"password,omitempty"`

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

// +kcc:proto=google.cloud.batch.v1.Runnable.Script
type Runnable_Script struct {
	// The path to a script file that is accessible from the host VM(s).
	//
	//  Unless the script file supports the default `#!/bin/sh` shell
	//  interpreter, you must specify an interpreter by including a
	//  [shebang line](https://en.wikipedia.org/wiki/Shebang_(Unix) as the
	//  first line of the file. For example, to execute the script using bash,
	//  include `#!/bin/bash` as the first line of the file. Alternatively,
	//  to execute the script using Python3, include `#!/usr/bin/env python3`
	//  as the first line of the file.
	// +kcc:proto:field=google.cloud.batch.v1.Runnable.Script.path
	Path *string `json:"path,omitempty"`

	// The text for a script.
	//
	//  Unless the script text supports the default `#!/bin/sh` shell
	//  interpreter, you must specify an interpreter by including a
	//  [shebang line](https://en.wikipedia.org/wiki/Shebang_(Unix) at the
	//  beginning of the text. For example, to execute the script using bash,
	//  include `#!/bin/bash\n` at the beginning of the text. Alternatively,
	//  to execute the script using Python3, include `#!/usr/bin/env python3\n`
	//  at the beginning of the text.
	// +kcc:proto:field=google.cloud.batch.v1.Runnable.Script.text
	Text *string `json:"text,omitempty"`
}

// +kcc:proto=google.cloud.batch.v1.ServiceAccount
type ServiceAccount struct {
	// Email address of the service account.
	// +kcc:proto:field=google.cloud.batch.v1.ServiceAccount.email
	Email *string `json:"email,omitempty"`

	// List of scopes to be enabled for this service account.
	// +kcc:proto:field=google.cloud.batch.v1.ServiceAccount.scopes
	Scopes []string `json:"scopes,omitempty"`
}

// +kcc:proto=google.cloud.batch.v1.StatusEvent
type StatusEvent struct {
	// Type of the event.
	// +kcc:proto:field=google.cloud.batch.v1.StatusEvent.type
	Type *string `json:"type,omitempty"`

	// Description of the event.
	// +kcc:proto:field=google.cloud.batch.v1.StatusEvent.description
	Description *string `json:"description,omitempty"`

	// The time this event occurred.
	// +kcc:proto:field=google.cloud.batch.v1.StatusEvent.event_time
	EventTime *string `json:"eventTime,omitempty"`

	// Task Execution.
	//  This field is only defined for task-level status events where the task
	//  fails.
	// +kcc:proto:field=google.cloud.batch.v1.StatusEvent.task_execution
	TaskExecution *TaskExecution `json:"taskExecution,omitempty"`

	// Task State.
	//  This field is only defined for task-level status events.
	// +kcc:proto:field=google.cloud.batch.v1.StatusEvent.task_state
	TaskState *string `json:"taskState,omitempty"`
}

// +kcc:proto=google.cloud.batch.v1.TaskExecution
type TaskExecution struct {
	// The exit code of a finished task.
	//
	//  If the task succeeded, the exit code will be 0. If the task failed but not
	//  due to the following reasons, the exit code will be 50000.
	//
	//  Otherwise, it can be from different sources:
	//  * Batch known failures:
	//  https://cloud.google.com/batch/docs/troubleshooting#reserved-exit-codes.
	//  * Batch runnable execution failures; you can rely on Batch logs to further
	//  diagnose: https://cloud.google.com/batch/docs/analyze-job-using-logs. If
	//  there are multiple runnables failures, Batch only exposes the first error.
	// +kcc:proto:field=google.cloud.batch.v1.TaskExecution.exit_code
	ExitCode *int32 `json:"exitCode,omitempty"`
}

// +kcc:proto=google.cloud.batch.v1.TaskGroup
type TaskGroup struct {

	// Required. Tasks in the group share the same task spec.
	// +kcc:proto:field=google.cloud.batch.v1.TaskGroup.task_spec
	TaskSpec *TaskSpec `json:"taskSpec,omitempty"`

	// Number of Tasks in the TaskGroup.
	//  Default is 1.
	// +kcc:proto:field=google.cloud.batch.v1.TaskGroup.task_count
	TaskCount *int64 `json:"taskCount,omitempty"`

	// Max number of tasks that can run in parallel.
	//  Default to min(task_count, parallel tasks per job limit).
	//  See: [Job Limits](https://cloud.google.com/batch/quotas#job_limits).
	//  Field parallelism must be 1 if the scheduling_policy is IN_ORDER.
	// +kcc:proto:field=google.cloud.batch.v1.TaskGroup.parallelism
	Parallelism *int64 `json:"parallelism,omitempty"`

	// Scheduling policy for Tasks in the TaskGroup.
	//  The default value is AS_SOON_AS_POSSIBLE.
	// +kcc:proto:field=google.cloud.batch.v1.TaskGroup.scheduling_policy
	SchedulingPolicy *string `json:"schedulingPolicy,omitempty"`

	// An array of environment variable mappings, which are passed to Tasks with
	//  matching indices. If task_environments is used then task_count should
	//  not be specified in the request (and will be ignored). Task count will be
	//  the length of task_environments.
	//
	//  Tasks get a BATCH_TASK_INDEX and BATCH_TASK_COUNT environment variable, in
	//  addition to any environment variables set in task_environments, specifying
	//  the number of Tasks in the Task's parent TaskGroup, and the specific Task's
	//  index in the TaskGroup (0 through BATCH_TASK_COUNT - 1).
	// +kcc:proto:field=google.cloud.batch.v1.TaskGroup.task_environments
	TaskEnvironments []Environment `json:"taskEnvironments,omitempty"`

	// Max number of tasks that can be run on a VM at the same time.
	//  If not specified, the system will decide a value based on available
	//  compute resources on a VM and task requirements.
	// +kcc:proto:field=google.cloud.batch.v1.TaskGroup.task_count_per_node
	TaskCountPerNode *int64 `json:"taskCountPerNode,omitempty"`

	// When true, Batch will populate a file with a list of all VMs assigned to
	//  the TaskGroup and set the BATCH_HOSTS_FILE environment variable to the path
	//  of that file. Defaults to false. The host file supports up to 1000 VMs.
	// +kcc:proto:field=google.cloud.batch.v1.TaskGroup.require_hosts_file
	RequireHostsFile *bool `json:"requireHostsFile,omitempty"`

	// When true, Batch will configure SSH to allow passwordless login between
	//  VMs running the Batch tasks in the same TaskGroup.
	// +kcc:proto:field=google.cloud.batch.v1.TaskGroup.permissive_ssh
	PermissiveSSH *bool `json:"permissiveSSH,omitempty"`

	// Optional. If not set or set to false, Batch uses the root user to execute
	//  runnables. If set to true, Batch runs the runnables using a non-root user.
	//  Currently, the non-root user Batch used is generated by OS Login. For more
	//  information, see [About OS
	//  Login](https://cloud.google.com/compute/docs/oslogin).
	// +kcc:proto:field=google.cloud.batch.v1.TaskGroup.run_as_non_root
	RunAsNonRoot *bool `json:"runAsNonRoot,omitempty"`
}

// +kcc:proto=google.cloud.batch.v1.TaskSpec
type TaskSpec struct {
	// Required. The sequence of one or more runnables (executable scripts,
	//  executable containers, and/or barriers) for each task in this task group to
	//  run. Each task runs this list of runnables in order. For a task to succeed,
	//  all of its script and container runnables each must meet at least one of
	//  the following conditions:
	//
	//  + The runnable exited with a zero status.
	//  + The runnable didn't finish, but you enabled its `background` subfield.
	//  + The runnable exited with a non-zero status, but you enabled its
	//    `ignore_exit_status` subfield.
	// +kcc:proto:field=google.cloud.batch.v1.TaskSpec.runnables
	Runnables []Runnable `json:"runnables,omitempty"`

	// ComputeResource requirements.
	// +kcc:proto:field=google.cloud.batch.v1.TaskSpec.compute_resource
	ComputeResource *ComputeResource `json:"computeResource,omitempty"`

	// Maximum duration the task should run before being automatically retried
	//  (if enabled) or automatically failed. Format the value of this field
	//  as a time limit in seconds followed by `s`&mdash;for example, `3600s`
	//  for 1 hour. The field accepts any value between 0 and the maximum listed
	//  for the `Duration` field type at
	//  https://protobuf.dev/reference/protobuf/google.protobuf/#duration; however,
	//  the actual maximum run time for a job will be limited to the maximum run
	//  time for a job listed at
	//  https://cloud.google.com/batch/quotas#max-job-duration.
	// +kcc:proto:field=google.cloud.batch.v1.TaskSpec.max_run_duration
	MaxRunDuration *string `json:"maxRunDuration,omitempty"`

	// Maximum number of retries on failures.
	//  The default, 0, which means never retry.
	//  The valid value range is [0, 10].
	// +kcc:proto:field=google.cloud.batch.v1.TaskSpec.max_retry_count
	MaxRetryCount *int32 `json:"maxRetryCount,omitempty"`

	// Lifecycle management schema when any task in a task group is failed.
	//  Currently we only support one lifecycle policy.
	//  When the lifecycle policy condition is met,
	//  the action in the policy will execute.
	//  If task execution result does not meet with the defined lifecycle
	//  policy, we consider it as the default policy.
	//  Default policy means if the exit code is 0, exit task.
	//  If task ends with non-zero exit code, retry the task with max_retry_count.
	// +kcc:proto:field=google.cloud.batch.v1.TaskSpec.lifecycle_policies
	LifecyclePolicies []LifecyclePolicy `json:"lifecyclePolicies,omitempty"`

	// Deprecated: please use environment(non-plural) instead.
	// +kcc:proto:field=google.cloud.batch.v1.TaskSpec.environments
	Environments map[string]string `json:"environments,omitempty"`

	// Volumes to mount before running Tasks using this TaskSpec.
	// +kcc:proto:field=google.cloud.batch.v1.TaskSpec.volumes
	Volumes []Volume `json:"volumes,omitempty"`

	// Environment variables to set before running the Task.
	// +kcc:proto:field=google.cloud.batch.v1.TaskSpec.environment
	Environment *Environment `json:"environment,omitempty"`
}

// +kcc:proto=google.cloud.batch.v1.Volume
type Volume struct {
	// A Network File System (NFS) volume. For example, a
	//  Filestore file share.
	// +kcc:proto:field=google.cloud.batch.v1.Volume.nfs
	Nfs *Nfs `json:"nfs,omitempty"`

	// A Google Cloud Storage (GCS) volume.
	// +kcc:proto:field=google.cloud.batch.v1.Volume.gcs
	Gcs *Gcs `json:"gcs,omitempty"`

	// Device name of an attached disk volume, which should align with a
	//  device_name specified by
	//  job.allocation_policy.instances[0].policy.disks[i].device_name or
	//  defined by the given instance template in
	//  job.allocation_policy.instances[0].instance_template.
	// +kcc:proto:field=google.cloud.batch.v1.Volume.device_name
	DeviceName *string `json:"deviceName,omitempty"`

	// The mount path for the volume, e.g. /mnt/disks/share.
	// +kcc:proto:field=google.cloud.batch.v1.Volume.mount_path
	MountPath *string `json:"mountPath,omitempty"`

	// Mount options vary based on the type of storage volume:
	//
	//  * For a Cloud Storage bucket, all the mount options provided
	//  by
	//    the [`gcsfuse` tool](https://cloud.google.com/storage/docs/gcsfuse-cli)
	//    are supported.
	//  * For an existing persistent disk, all mount options provided by the
	//    [`mount` command](https://man7.org/linux/man-pages/man8/mount.8.html)
	//    except writing are supported. This is due to restrictions of
	//    [multi-writer
	//    mode](https://cloud.google.com/compute/docs/disks/sharing-disks-between-vms).
	//  * For any other disk or a Network File System (NFS), all the
	//    mount options provided by the `mount` command are supported.
	// +kcc:proto:field=google.cloud.batch.v1.Volume.mount_options
	MountOptions []string `json:"mountOptions,omitempty"`
}

// +kcc:proto=google.cloud.batch.v1.TaskGroup
type TaskGroupObservedState struct {
	// Output only. TaskGroup name.
	//  The system generates this field based on parent Job name.
	//  For example:
	//  "projects/123456/locations/us-west1/jobs/job01/taskGroups/group01".
	// +kcc:proto:field=google.cloud.batch.v1.TaskGroup.name
	Name *string `json:"name,omitempty"`
}
