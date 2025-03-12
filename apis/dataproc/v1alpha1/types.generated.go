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


// +kcc:proto=google.cloud.dataproc.v1.AcceleratorConfig
type AcceleratorConfig struct {
	// Full URL, partial URI, or short name of the accelerator type resource to
	//  expose to this instance. See
	//  [Compute Engine
	//  AcceleratorTypes](https://cloud.google.com/compute/docs/reference/v1/acceleratorTypes).
	//
	//  Examples:
	//
	//  * `https://www.googleapis.com/compute/v1/projects/[project_id]/zones/[zone]/acceleratorTypes/nvidia-tesla-t4`
	//  * `projects/[project_id]/zones/[zone]/acceleratorTypes/nvidia-tesla-t4`
	//  * `nvidia-tesla-t4`
	//
	//  **Auto Zone Exception**: If you are using the Dataproc
	//  [Auto Zone
	//  Placement](https://cloud.google.com/dataproc/docs/concepts/configuring-clusters/auto-zone#using_auto_zone_placement)
	//  feature, you must use the short name of the accelerator type
	//  resource, for example, `nvidia-tesla-t4`.
	// +kcc:proto:field=google.cloud.dataproc.v1.AcceleratorConfig.accelerator_type_uri
	AcceleratorTypeURI *string `json:"acceleratorTypeURI,omitempty"`

	// The number of the accelerator cards of this type exposed to this instance.
	// +kcc:proto:field=google.cloud.dataproc.v1.AcceleratorConfig.accelerator_count
	AcceleratorCount *int32 `json:"acceleratorCount,omitempty"`
}

// +kcc:proto=google.cloud.dataproc.v1.DiskConfig
type DiskConfig struct {
	// Optional. Type of the boot disk (default is "pd-standard").
	//  Valid values: "pd-balanced" (Persistent Disk Balanced Solid State Drive),
	//  "pd-ssd" (Persistent Disk Solid State Drive),
	//  or "pd-standard" (Persistent Disk Hard Disk Drive).
	//  See [Disk types](https://cloud.google.com/compute/docs/disks#disk-types).
	// +kcc:proto:field=google.cloud.dataproc.v1.DiskConfig.boot_disk_type
	BootDiskType *string `json:"bootDiskType,omitempty"`

	// Optional. Size in GB of the boot disk (default is 500GB).
	// +kcc:proto:field=google.cloud.dataproc.v1.DiskConfig.boot_disk_size_gb
	BootDiskSizeGB *int32 `json:"bootDiskSizeGB,omitempty"`

	// Optional. Number of attached SSDs, from 0 to 8 (default is 0).
	//  If SSDs are not attached, the boot disk is used to store runtime logs and
	//  [HDFS](https://hadoop.apache.org/docs/r1.2.1/hdfs_user_guide.html) data.
	//  If one or more SSDs are attached, this runtime bulk
	//  data is spread across them, and the boot disk contains only basic
	//  config and installed binaries.
	//
	//  Note: Local SSD options may vary by machine type and number of vCPUs
	//  selected.
	// +kcc:proto:field=google.cloud.dataproc.v1.DiskConfig.num_local_ssds
	NumLocalSsds *int32 `json:"numLocalSsds,omitempty"`

	// Optional. Interface type of local SSDs (default is "scsi").
	//  Valid values: "scsi" (Small Computer System Interface),
	//  "nvme" (Non-Volatile Memory Express).
	//  See [local SSD
	//  performance](https://cloud.google.com/compute/docs/disks/local-ssd#performance).
	// +kcc:proto:field=google.cloud.dataproc.v1.DiskConfig.local_ssd_interface
	LocalSsdInterface *string `json:"localSsdInterface,omitempty"`

	// Optional. Indicates how many IOPS to provision for the disk. This sets the
	//  number of I/O operations per second that the disk can handle. Note: This
	//  field is only supported if boot_disk_type is hyperdisk-balanced.
	// +kcc:proto:field=google.cloud.dataproc.v1.DiskConfig.boot_disk_provisioned_iops
	BootDiskProvisionedIops *int64 `json:"bootDiskProvisionedIops,omitempty"`

	// Optional. Indicates how much throughput to provision for the disk. This
	//  sets the number of throughput mb per second that the disk can handle.
	//  Values must be greater than or equal to 1. Note: This field is only
	//  supported if boot_disk_type is hyperdisk-balanced.
	// +kcc:proto:field=google.cloud.dataproc.v1.DiskConfig.boot_disk_provisioned_throughput
	BootDiskProvisionedThroughput *int64 `json:"bootDiskProvisionedThroughput,omitempty"`
}

// +kcc:proto=google.cloud.dataproc.v1.InstanceFlexibilityPolicy
type InstanceFlexibilityPolicy struct {
	// Optional. Defines how the Group selects the provisioning model to ensure
	//  required reliability.
	// +kcc:proto:field=google.cloud.dataproc.v1.InstanceFlexibilityPolicy.provisioning_model_mix
	ProvisioningModelMix *InstanceFlexibilityPolicy_ProvisioningModelMix `json:"provisioningModelMix,omitempty"`

	// Optional. List of instance selection options that the group will use when
	//  creating new VMs.
	// +kcc:proto:field=google.cloud.dataproc.v1.InstanceFlexibilityPolicy.instance_selection_list
	InstanceSelectionList []InstanceFlexibilityPolicy_InstanceSelection `json:"instanceSelectionList,omitempty"`
}

// +kcc:proto=google.cloud.dataproc.v1.InstanceFlexibilityPolicy.InstanceSelection
type InstanceFlexibilityPolicy_InstanceSelection struct {
	// Optional. Full machine-type names, e.g. "n1-standard-16".
	// +kcc:proto:field=google.cloud.dataproc.v1.InstanceFlexibilityPolicy.InstanceSelection.machine_types
	MachineTypes []string `json:"machineTypes,omitempty"`

	// Optional. Preference of this instance selection. Lower number means
	//  higher preference. Dataproc will first try to create a VM based on the
	//  machine-type with priority rank and fallback to next rank based on
	//  availability. Machine types and instance selections with the same
	//  priority have the same preference.
	// +kcc:proto:field=google.cloud.dataproc.v1.InstanceFlexibilityPolicy.InstanceSelection.rank
	Rank *int32 `json:"rank,omitempty"`
}

// +kcc:proto=google.cloud.dataproc.v1.InstanceFlexibilityPolicy.InstanceSelectionResult
type InstanceFlexibilityPolicy_InstanceSelectionResult struct {
}

// +kcc:proto=google.cloud.dataproc.v1.InstanceFlexibilityPolicy.ProvisioningModelMix
type InstanceFlexibilityPolicy_ProvisioningModelMix struct {
	// Optional. The base capacity that will always use Standard VMs to avoid
	//  risk of more preemption than the minimum capacity you need. Dataproc will
	//  create only standard VMs until it reaches standard_capacity_base, then it
	//  will start using standard_capacity_percent_above_base to mix Spot with
	//  Standard VMs. eg. If 15 instances are requested and
	//  standard_capacity_base is 5, Dataproc will create 5 standard VMs and then
	//  start mixing spot and standard VMs for remaining 10 instances.
	// +kcc:proto:field=google.cloud.dataproc.v1.InstanceFlexibilityPolicy.ProvisioningModelMix.standard_capacity_base
	StandardCapacityBase *int32 `json:"standardCapacityBase,omitempty"`

	// Optional. The percentage of target capacity that should use Standard VM.
	//  The remaining percentage will use Spot VMs. The percentage applies only
	//  to the capacity above standard_capacity_base. eg. If 15 instances are
	//  requested and standard_capacity_base is 5 and
	//  standard_capacity_percent_above_base is 30, Dataproc will create 5
	//  standard VMs and then start mixing spot and standard VMs for remaining 10
	//  instances. The mix will be 30% standard and 70% spot.
	// +kcc:proto:field=google.cloud.dataproc.v1.InstanceFlexibilityPolicy.ProvisioningModelMix.standard_capacity_percent_above_base
	StandardCapacityPercentAboveBase *int32 `json:"standardCapacityPercentAboveBase,omitempty"`
}

// +kcc:proto=google.cloud.dataproc.v1.InstanceGroupConfig
type InstanceGroupConfig struct {
	// Optional. The number of VM instances in the instance group.
	//  For [HA
	//  cluster](/dataproc/docs/concepts/configuring-clusters/high-availability)
	//  [master_config](#FIELDS.master_config) groups, **must be set to 3**.
	//  For standard cluster [master_config](#FIELDS.master_config) groups,
	//  **must be set to 1**.
	// +kcc:proto:field=google.cloud.dataproc.v1.InstanceGroupConfig.num_instances
	NumInstances *int32 `json:"numInstances,omitempty"`

	// Optional. The Compute Engine image resource used for cluster instances.
	//
	//  The URI can represent an image or image family.
	//
	//  Image examples:
	//
	//  * `https://www.googleapis.com/compute/v1/projects/[project_id]/global/images/[image-id]`
	//  * `projects/[project_id]/global/images/[image-id]`
	//  * `image-id`
	//
	//  Image family examples. Dataproc will use the most recent
	//  image from the family:
	//
	//  * `https://www.googleapis.com/compute/v1/projects/[project_id]/global/images/family/[custom-image-family-name]`
	//  * `projects/[project_id]/global/images/family/[custom-image-family-name]`
	//
	//  If the URI is unspecified, it will be inferred from
	//  `SoftwareConfig.image_version` or the system default.
	// +kcc:proto:field=google.cloud.dataproc.v1.InstanceGroupConfig.image_uri
	ImageURI *string `json:"imageURI,omitempty"`

	// Optional. The Compute Engine machine type used for cluster instances.
	//
	//  A full URL, partial URI, or short name are valid. Examples:
	//
	//  * `https://www.googleapis.com/compute/v1/projects/[project_id]/zones/[zone]/machineTypes/n1-standard-2`
	//  * `projects/[project_id]/zones/[zone]/machineTypes/n1-standard-2`
	//  * `n1-standard-2`
	//
	//  **Auto Zone Exception**: If you are using the Dataproc
	//  [Auto Zone
	//  Placement](https://cloud.google.com/dataproc/docs/concepts/configuring-clusters/auto-zone#using_auto_zone_placement)
	//  feature, you must use the short name of the machine type
	//  resource, for example, `n1-standard-2`.
	// +kcc:proto:field=google.cloud.dataproc.v1.InstanceGroupConfig.machine_type_uri
	MachineTypeURI *string `json:"machineTypeURI,omitempty"`

	// Optional. Disk option config settings.
	// +kcc:proto:field=google.cloud.dataproc.v1.InstanceGroupConfig.disk_config
	DiskConfig *DiskConfig `json:"diskConfig,omitempty"`

	// Optional. Specifies the preemptibility of the instance group.
	//
	//  The default value for master and worker groups is
	//  `NON_PREEMPTIBLE`. This default cannot be changed.
	//
	//  The default value for secondary instances is
	//  `PREEMPTIBLE`.
	// +kcc:proto:field=google.cloud.dataproc.v1.InstanceGroupConfig.preemptibility
	Preemptibility *string `json:"preemptibility,omitempty"`

	// Optional. The Compute Engine accelerator configuration for these
	//  instances.
	// +kcc:proto:field=google.cloud.dataproc.v1.InstanceGroupConfig.accelerators
	Accelerators []AcceleratorConfig `json:"accelerators,omitempty"`

	// Optional. Specifies the minimum cpu platform for the Instance Group.
	//  See [Dataproc -> Minimum CPU
	//  Platform](https://cloud.google.com/dataproc/docs/concepts/compute/dataproc-min-cpu).
	// +kcc:proto:field=google.cloud.dataproc.v1.InstanceGroupConfig.min_cpu_platform
	MinCPUPlatform *string `json:"minCPUPlatform,omitempty"`

	// Optional. The minimum number of primary worker instances to create.
	//  If `min_num_instances` is set, cluster creation will succeed if
	//  the number of primary workers created is at least equal to the
	//  `min_num_instances` number.
	//
	//  Example: Cluster creation request with `num_instances` = `5` and
	//  `min_num_instances` = `3`:
	//
	//  *  If 4 VMs are created and 1 instance fails,
	//     the failed VM is deleted. The cluster is
	//     resized to 4 instances and placed in a `RUNNING` state.
	//  *  If 2 instances are created and 3 instances fail,
	//     the cluster in placed in an `ERROR` state. The failed VMs
	//     are not deleted.
	// +kcc:proto:field=google.cloud.dataproc.v1.InstanceGroupConfig.min_num_instances
	MinNumInstances *int32 `json:"minNumInstances,omitempty"`

	// Optional. Instance flexibility Policy allowing a mixture of VM shapes and
	//  provisioning models.
	// +kcc:proto:field=google.cloud.dataproc.v1.InstanceGroupConfig.instance_flexibility_policy
	InstanceFlexibilityPolicy *InstanceFlexibilityPolicy `json:"instanceFlexibilityPolicy,omitempty"`

	// Optional. Configuration to handle the startup of instances during cluster
	//  create and update process.
	// +kcc:proto:field=google.cloud.dataproc.v1.InstanceGroupConfig.startup_config
	StartupConfig *StartupConfig `json:"startupConfig,omitempty"`
}

// +kcc:proto=google.cloud.dataproc.v1.InstanceReference
type InstanceReference struct {
	// The user-friendly name of the Compute Engine instance.
	// +kcc:proto:field=google.cloud.dataproc.v1.InstanceReference.instance_name
	InstanceName *string `json:"instanceName,omitempty"`

	// The unique identifier of the Compute Engine instance.
	// +kcc:proto:field=google.cloud.dataproc.v1.InstanceReference.instance_id
	InstanceID *string `json:"instanceID,omitempty"`

	// The public RSA key used for sharing data with this instance.
	// +kcc:proto:field=google.cloud.dataproc.v1.InstanceReference.public_key
	PublicKey *string `json:"publicKey,omitempty"`

	// The public ECIES key used for sharing data with this instance.
	// +kcc:proto:field=google.cloud.dataproc.v1.InstanceReference.public_ecies_key
	PublicEciesKey *string `json:"publicEciesKey,omitempty"`
}

// +kcc:proto=google.cloud.dataproc.v1.ManagedGroupConfig
type ManagedGroupConfig struct {
}

// +kcc:proto=google.cloud.dataproc.v1.StartupConfig
type StartupConfig struct {
	// Optional. The config setting to enable cluster creation/ updation to be
	//  successful only after required_registration_fraction of instances are up
	//  and running. This configuration is applicable to only secondary workers for
	//  now. The cluster will fail if required_registration_fraction of instances
	//  are not available. This will include instance creation, agent registration,
	//  and service registration (if enabled).
	// +kcc:proto:field=google.cloud.dataproc.v1.StartupConfig.required_registration_fraction
	RequiredRegistrationFraction *float64 `json:"requiredRegistrationFraction,omitempty"`
}

// +kcc:proto=google.cloud.dataproc.v1.InstanceFlexibilityPolicy
type InstanceFlexibilityPolicyObservedState struct {
	// Output only. A list of instance selection results in the group.
	// +kcc:proto:field=google.cloud.dataproc.v1.InstanceFlexibilityPolicy.instance_selection_results
	InstanceSelectionResults []InstanceFlexibilityPolicy_InstanceSelectionResult `json:"instanceSelectionResults,omitempty"`
}

// +kcc:proto=google.cloud.dataproc.v1.InstanceFlexibilityPolicy.InstanceSelectionResult
type InstanceFlexibilityPolicy_InstanceSelectionResultObservedState struct {
	// Output only. Full machine-type names, e.g. "n1-standard-16".
	// +kcc:proto:field=google.cloud.dataproc.v1.InstanceFlexibilityPolicy.InstanceSelectionResult.machine_type
	MachineType *string `json:"machineType,omitempty"`

	// Output only. Number of VM provisioned with the machine_type.
	// +kcc:proto:field=google.cloud.dataproc.v1.InstanceFlexibilityPolicy.InstanceSelectionResult.vm_count
	VmCount *int32 `json:"vmCount,omitempty"`
}

// +kcc:proto=google.cloud.dataproc.v1.InstanceGroupConfig
type InstanceGroupConfigObservedState struct {
	// Output only. The list of instance names. Dataproc derives the names
	//  from `cluster_name`, `num_instances`, and the instance group.
	// +kcc:proto:field=google.cloud.dataproc.v1.InstanceGroupConfig.instance_names
	InstanceNames []string `json:"instanceNames,omitempty"`

	// Output only. List of references to Compute Engine instances.
	// +kcc:proto:field=google.cloud.dataproc.v1.InstanceGroupConfig.instance_references
	InstanceReferences []InstanceReference `json:"instanceReferences,omitempty"`

	// Output only. Specifies that this instance group contains preemptible
	//  instances.
	// +kcc:proto:field=google.cloud.dataproc.v1.InstanceGroupConfig.is_preemptible
	IsPreemptible *bool `json:"isPreemptible,omitempty"`

	// Output only. The config for Compute Engine Instance Group
	//  Manager that manages this group.
	//  This is only used for preemptible instance groups.
	// +kcc:proto:field=google.cloud.dataproc.v1.InstanceGroupConfig.managed_group_config
	ManagedGroupConfig *ManagedGroupConfig `json:"managedGroupConfig,omitempty"`

	// Optional. Instance flexibility Policy allowing a mixture of VM shapes and
	//  provisioning models.
	// +kcc:proto:field=google.cloud.dataproc.v1.InstanceGroupConfig.instance_flexibility_policy
	InstanceFlexibilityPolicy *InstanceFlexibilityPolicyObservedState `json:"instanceFlexibilityPolicy,omitempty"`
}

// +kcc:proto=google.cloud.dataproc.v1.ManagedGroupConfig
type ManagedGroupConfigObservedState struct {
	// Output only. The name of the Instance Template used for the Managed
	//  Instance Group.
	// +kcc:proto:field=google.cloud.dataproc.v1.ManagedGroupConfig.instance_template_name
	InstanceTemplateName *string `json:"instanceTemplateName,omitempty"`

	// Output only. The name of the Instance Group Manager for this group.
	// +kcc:proto:field=google.cloud.dataproc.v1.ManagedGroupConfig.instance_group_manager_name
	InstanceGroupManagerName *string `json:"instanceGroupManagerName,omitempty"`

	// Output only. The partial URI to the instance group manager for this group.
	//  E.g. projects/my-project/regions/us-central1/instanceGroupManagers/my-igm.
	// +kcc:proto:field=google.cloud.dataproc.v1.ManagedGroupConfig.instance_group_manager_uri
	InstanceGroupManagerURI *string `json:"instanceGroupManagerURI,omitempty"`
}
