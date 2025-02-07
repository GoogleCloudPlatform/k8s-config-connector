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
	CpuMilli *int64 `json:"cpuMilli,omitempty"`

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
type GCS struct {
	// Remote path, either a bucket name or a subdirectory of a bucket, e.g.:
	//  bucket_name, bucket_name/subdirectory/
	// +kcc:proto:field=google.cloud.batch.v1.GCS.remote_path
	RemotePath *string `json:"remotePath,omitempty"`
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

// +kcc:proto=google.cloud.batch.v1.NFS
type NFS struct {
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
	Nfs *NFS `json:"nfs,omitempty"`

	// A Google Cloud Storage (GCS) volume.
	// +kcc:proto:field=google.cloud.batch.v1.Volume.gcs
	Gcs *GCS `json:"gcs,omitempty"`

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
