// Copyright 2024 Google LLC
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

// +kcc:proto=google.cloud.workstations.v1.WorkstationConfig.Container
type WorkstationConfig_Container struct {
	// Optional. A Docker container image that defines a custom environment.
	//
	//  Cloud Workstations provides a number of
	//  [preconfigured
	//  images](https://cloud.google.com/workstations/docs/preconfigured-base-images),
	//  but you can create your own
	//  [custom container
	//  images](https://cloud.google.com/workstations/docs/custom-container-images).
	//  If using a private image, the `host.gceInstance.serviceAccount` field
	//  must be specified in the workstation configuration and must have
	//  permission to pull the specified image. Otherwise, the image must be
	//  publicly accessible.
	Image *string `json:"image,omitempty"`

	// Optional. If set, overrides the default ENTRYPOINT specified by the
	//  image.
	Command []string `json:"command,omitempty"`

	// Optional. Arguments passed to the entrypoint.
	Args []string `json:"args,omitempty"`

	// Optional. Environment variables passed to the container's entrypoint.
	Env map[string]string `json:"env,omitempty"`

	// Optional. If set, overrides the default DIR specified by the image.
	WorkingDir *string `json:"workingDir,omitempty"`

	// Optional. If set, overrides the USER specified in the image with the
	//  given uid.
	RunAsUser *int32 `json:"runAsUser,omitempty"`
}

// +kcc:proto=google.cloud.workstations.v1.WorkstationConfig.CustomerEncryptionKey
type WorkstationConfig_CustomerEncryptionKey struct {
	// Immutable. The name of the Google Cloud KMS encryption key. For example,
	//  `"projects/PROJECT_ID/locations/REGION/keyRings/KEY_RING/cryptoKeys/KEY_NAME"`.
	//  The key must be in the same region as the workstation configuration.
	KmsKey *string `json:"kmsKey,omitempty"`

	// Immutable. The service account to use with the specified
	//  KMS key. We recommend that you use a separate service account
	//  and follow KMS best practices. For more information, see
	//  [Separation of
	//  duties](https://cloud.google.com/kms/docs/separation-of-duties) and
	//  `gcloud kms keys add-iam-policy-binding`
	//  [`--member`](https://cloud.google.com/sdk/gcloud/reference/kms/keys/add-iam-policy-binding#--member).
	KmsKeyServiceAccount *string `json:"kmsKeyServiceAccount,omitempty"`
}

// +kcc:proto=google.cloud.workstations.v1.WorkstationConfig.Host
type WorkstationConfig_Host struct {
	// Specifies a Compute Engine instance as the host.
	GceInstance *WorkstationConfig_Host_GceInstance `json:"gceInstance,omitempty"`
}

// +kcc:proto=google.cloud.workstations.v1.WorkstationConfig.Host.GceInstance
type WorkstationConfig_Host_GceInstance struct {
	// Optional. The type of machine to use for VM instances—for example,
	//  `"e2-standard-4"`. For more information about machine types that
	//  Cloud Workstations supports, see the list of
	//  [available machine
	//  types](https://cloud.google.com/workstations/docs/available-machine-types).
	MachineType *string `json:"machineType,omitempty"`

	// Optional. The email address of the service account for Cloud
	//  Workstations VMs created with this configuration. When specified, be
	//  sure that the service account has `logginglogEntries.create` permission
	//  on the project so it can write logs out to Cloud Logging. If using a
	//  custom container image, the service account must have permissions to
	//  pull the specified image.
	//
	//  If you as the administrator want to be able to `ssh` into the
	//  underlying VM, you need to set this value to a service account
	//  for which you have the `iam.serviceAccounts.actAs` permission.
	//  Conversely, if you don't want anyone to be able to `ssh` into the
	//  underlying VM, use a service account where no one has that
	//  permission.
	//
	//  If not set, VMs run with a service account provided by the
	//  Cloud Workstations service, and the image must be publicly
	//  accessible.
	ServiceAccount *string `json:"serviceAccount,omitempty"`

	// Optional. Scopes to grant to the
	//  [service_account][google.cloud.workstations.v1.WorkstationConfig.Host.GceInstance.service_account].
	//  Various scopes are automatically added based on feature usage. When
	//  specified, users of workstations under this configuration must have
	//  `iam.serviceAccounts.actAs` on the service account.
	ServiceAccountScopes []string `json:"serviceAccountScopes,omitempty"`

	// Optional. Network tags to add to the Compute Engine VMs backing the
	//  workstations. This option applies
	//  [network
	//  tags](https://cloud.google.com/vpc/docs/add-remove-network-tags) to VMs
	//  created with this configuration. These network tags enable the creation
	//  of [firewall
	//  rules](https://cloud.google.com/workstations/docs/configure-firewall-rules).
	Tags []string `json:"tags,omitempty"`

	// Optional. The number of VMs that the system should keep idle so that
	//  new workstations can be started quickly for new users. Defaults to `0`
	//  in the API.
	PoolSize *int32 `json:"poolSize,omitempty"`

	// Output only. Number of instances currently available in the pool for
	//  faster workstation startup.
	PooledInstances *int32 `json:"pooledInstances,omitempty"`

	// Optional. When set to true, disables public IP addresses for VMs. If
	//  you disable public IP addresses, you must set up Private Google Access
	//  or Cloud NAT on your network. If you use Private Google Access and you
	//  use `private.googleapis.com` or `restricted.googleapis.com` for
	//  Container Registry and Artifact Registry, make sure that you set
	//  up DNS records for domains `*.gcr.io` and `*.pkg.dev`.
	//  Defaults to false (VMs have public IP addresses).
	DisablePublicIPAddresses *bool `json:"disablePublicIPAddresses,omitempty"`

	// Optional. Whether to enable nested virtualization on Cloud Workstations
	//  VMs created under this workstation configuration.
	//
	//  Nested virtualization lets you run virtual machine (VM) instances
	//  inside your workstation. Before enabling nested virtualization,
	//  consider the following important considerations. Cloud Workstations
	//  instances are subject to the [same restrictions as Compute Engine
	//  instances](https://cloud.google.com/compute/docs/instances/nested-virtualization/overview#restrictions):
	//
	//  * **Organization policy**: projects, folders, or
	//  organizations may be restricted from creating nested VMs if the
	//  **Disable VM nested virtualization** constraint is enforced in
	//  the organization policy. For more information, see the
	//  Compute Engine section,
	//  [Checking whether nested virtualization is
	//  allowed](https://cloud.google.com/compute/docs/instances/nested-virtualization/managing-constraint#checking_whether_nested_virtualization_is_allowed).
	//  * **Performance**: nested VMs might experience a 10% or greater
	//  decrease in performance for workloads that are CPU-bound and
	//  possibly greater than a 10% decrease for workloads that are
	//  input/output bound.
	//  * **Machine Type**: nested virtualization can only be enabled on
	//  workstation configurations that specify a
	//  [machine_type][google.cloud.workstations.v1.WorkstationConfig.Host.GceInstance.machine_type]
	//  in the N1 or N2 machine series.
	//  * **GPUs**: nested virtualization may not be enabled on workstation
	//  configurations with accelerators.
	//  * **Operating System**: Because
	//  [Container-Optimized
	//  OS](https://cloud.google.com/compute/docs/images/os-details#container-optimized_os_cos)
	//  does not support nested virtualization, when nested virtualization is
	//  enabled, the underlying Compute Engine VM instances boot from an
	//  [Ubuntu
	//  LTS](https://cloud.google.com/compute/docs/images/os-details#ubuntu_lts)
	//  image.
	EnableNestedVirtualization *bool `json:"enableNestedVirtualization,omitempty"`

	// Optional. A set of Compute Engine Shielded instance options.
	ShieldedInstanceConfig *WorkstationConfig_Host_GceInstance_GceShieldedInstanceConfig `json:"shieldedInstanceConfig,omitempty"`

	// Optional. A set of Compute Engine Confidential VM instance options.
	ConfidentialInstanceConfig *WorkstationConfig_Host_GceInstance_GceConfidentialInstanceConfig `json:"confidentialInstanceConfig,omitempty"`

	// Optional. The size of the boot disk for the VM in gigabytes (GB).
	//  The minimum boot disk size is `30` GB. Defaults to `50` GB.
	BootDiskSizeGb *int32 `json:"bootDiskSizeGb,omitempty"`
}

// +kcc:proto=google.cloud.workstations.v1.WorkstationConfig.Host.GceInstance.GceConfidentialInstanceConfig
type WorkstationConfig_Host_GceInstance_GceConfidentialInstanceConfig struct {
	// Optional. Whether the instance has confidential compute enabled.
	EnableConfidentialCompute *bool `json:"enableConfidentialCompute,omitempty"`
}

// +kcc:proto=google.cloud.workstations.v1.WorkstationConfig.Host.GceInstance.GceShieldedInstanceConfig
type WorkstationConfig_Host_GceInstance_GceShieldedInstanceConfig struct {
	// Optional. Whether the instance has Secure Boot enabled.
	EnableSecureBoot *bool `json:"enableSecureBoot,omitempty"`

	// Optional. Whether the instance has the vTPM enabled.
	EnableVtpm *bool `json:"enableVtpm,omitempty"`

	// Optional. Whether the instance has integrity monitoring enabled.
	EnableIntegrityMonitoring *bool `json:"enableIntegrityMonitoring,omitempty"`
}

// +kcc:proto=google.cloud.workstations.v1.WorkstationConfig.PersistentDirectory
type WorkstationConfig_PersistentDirectory struct {
	// A PersistentDirectory backed by a Compute Engine persistent disk.
	GcePd *WorkstationConfig_PersistentDirectory_GceRegionalPersistentDisk `json:"gcePd,omitempty"`

	// Optional. Location of this directory in the running workstation.
	MountPath *string `json:"mountPath,omitempty"`
}

// +kcc:proto=google.cloud.workstations.v1.WorkstationConfig.PersistentDirectory.GceRegionalPersistentDisk
type WorkstationConfig_PersistentDirectory_GceRegionalPersistentDisk struct {
	// Optional. The GB capacity of a persistent home directory for each
	//  workstation created with this configuration. Must be empty if
	//  [source_snapshot][google.cloud.workstations.v1.WorkstationConfig.PersistentDirectory.GceRegionalPersistentDisk.source_snapshot]
	//  is set.
	//
	//  Valid values are `10`, `50`, `100`, `200`, `500`, or `1000`.
	//  Defaults to `200`. If less than `200` GB, the
	//  [disk_type][google.cloud.workstations.v1.WorkstationConfig.PersistentDirectory.GceRegionalPersistentDisk.disk_type]
	//  must be
	//  `"pd-balanced"` or `"pd-ssd"`.
	SizeGb *int32 `json:"sizeGb,omitempty"`

	// Optional. Type of file system that the disk should be formatted with.
	//  The workstation image must support this file system type. Must be empty
	//  if
	//  [source_snapshot][google.cloud.workstations.v1.WorkstationConfig.PersistentDirectory.GceRegionalPersistentDisk.source_snapshot]
	//  is set. Defaults to `"ext4"`.
	FsType *string `json:"fsType,omitempty"`

	// Optional. The [type of the persistent
	//  disk](https://cloud.google.com/compute/docs/disks#disk-types) for the
	//  home directory. Defaults to `"pd-standard"`.
	DiskType *string `json:"diskType,omitempty"`

	// Optional. Name of the snapshot to use as the source for the disk. If
	//  set,
	//  [size_gb][google.cloud.workstations.v1.WorkstationConfig.PersistentDirectory.GceRegionalPersistentDisk.size_gb]
	//  and
	//  [fs_type][google.cloud.workstations.v1.WorkstationConfig.PersistentDirectory.GceRegionalPersistentDisk.fs_type]
	//  must be empty.
	SourceSnapshot *string `json:"sourceSnapshot,omitempty"`

	// Optional. Whether the persistent disk should be deleted when the
	//  workstation is deleted. Valid values are `DELETE` and `RETAIN`.
	//  Defaults to `DELETE`.
	ReclaimPolicy *string `json:"reclaimPolicy,omitempty"`
}

// +kcc:proto=google.cloud.workstations.v1.WorkstationConfig.ReadinessCheck
type WorkstationConfig_ReadinessCheck struct {
	// Optional. Path to which the request should be sent.
	Path *string `json:"path,omitempty"`

	// Optional. Port to which the request should be sent.
	Port *int32 `json:"port,omitempty"`
}
