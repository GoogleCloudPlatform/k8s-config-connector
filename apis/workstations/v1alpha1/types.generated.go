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

// +kcc:proto=google.cloud.workstations.v1.WorkstationConfig
type WorkstationConfig struct {
	// Full name of this workstation configuration.
	Name *string `json:"name,omitempty"`

	// Optional. Human-readable name for this workstation configuration.
	DisplayName *string `json:"displayName,omitempty"`

	// Output only. A system-assigned unique identifier for this workstation
	//  configuration.
	Uid *string `json:"uid,omitempty"`

	// Output only. Indicates whether this workstation configuration is currently
	//  being updated to match its intended state.
	Reconciling *bool `json:"reconciling,omitempty"`

	// Optional. Client-specified annotations.
	Annotations map[string]string `json:"annotations,omitempty"`

	// Optional.
	//  [Labels](https://cloud.google.com/workstations/docs/label-resources) that
	//  are applied to the workstation configuration and that are also propagated
	//  to the underlying Compute Engine resources.
	Labels map[string]string `json:"labels,omitempty"`

	// Output only. Time when this workstation configuration was created.
	CreateTime *string `json:"createTime,omitempty"`

	// Output only. Time when this workstation configuration was most recently
	//  updated.
	UpdateTime *string `json:"updateTime,omitempty"`

	// Output only. Time when this workstation configuration was soft-deleted.
	DeleteTime *string `json:"deleteTime,omitempty"`

	// Optional. Checksum computed by the server. May be sent on update and delete
	//  requests to make sure that the client has an up-to-date value before
	//  proceeding.
	Etag *string `json:"etag,omitempty"`

	// Optional. Number of seconds to wait before automatically stopping a
	//  workstation after it last received user traffic.
	//
	//  A value of `"0s"` indicates that Cloud Workstations VMs created with this
	//  configuration should never time out due to idleness.
	//  Provide
	//  [duration](https://developers.google.com/protocol-buffers/docs/reference/google.protobuf#duration)
	//  terminated by `s` for seconds—for example, `"7200s"` (2 hours).
	//  The default is `"1200s"` (20 minutes).
	IdleTimeout *string `json:"idleTimeout,omitempty"`

	// Optional. Number of seconds that a workstation can run until it is
	//  automatically shut down. We recommend that workstations be shut down daily
	//  to reduce costs and so that security updates can be applied upon restart.
	//  The
	//  [idle_timeout][google.cloud.workstations.v1.WorkstationConfig.idle_timeout]
	//  and
	//  [running_timeout][google.cloud.workstations.v1.WorkstationConfig.running_timeout]
	//  fields are independent of each other. Note that the
	//  [running_timeout][google.cloud.workstations.v1.WorkstationConfig.running_timeout]
	//  field shuts down VMs after the specified time, regardless of whether or not
	//  the VMs are idle.
	//
	//  Provide duration terminated by `s` for seconds—for example, `"54000s"`
	//  (15 hours). Defaults to `"43200s"` (12 hours). A value of `"0s"` indicates
	//  that workstations using this configuration should never time out. If
	//  [encryption_key][google.cloud.workstations.v1.WorkstationConfig.encryption_key]
	//  is set, it must be greater than `"0s"` and less than
	//  `"86400s"` (24 hours).
	//
	//  Warning: A value of `"0s"` indicates that Cloud Workstations VMs created
	//  with this configuration have no maximum running time. This is strongly
	//  discouraged because you incur costs and will not pick up security updates.
	RunningTimeout *string `json:"runningTimeout,omitempty"`

	// Optional. Runtime host for the workstation.
	Host *WorkstationConfig_Host `json:"host,omitempty"`

	// Optional. Directories to persist across workstation sessions.
	PersistentDirectories []WorkstationConfig_PersistentDirectory `json:"persistentDirectories,omitempty"`

	// Optional. Container that runs upon startup for each workstation using this
	//  workstation configuration.
	Container *WorkstationConfig_Container `json:"container,omitempty"`

	// Immutable. Encrypts resources of this workstation configuration using a
	//  customer-managed encryption key (CMEK).
	//
	//  If specified, the boot disk of the Compute Engine instance and the
	//  persistent disk are encrypted using this encryption key. If
	//  this field is not set, the disks are encrypted using a generated
	//  key. Customer-managed encryption keys do not protect disk metadata.
	//
	//  If the customer-managed encryption key is rotated, when the workstation
	//  instance is stopped, the system attempts to recreate the
	//  persistent disk with the new version of the key. Be sure to keep
	//  older versions of the key until the persistent disk is recreated.
	//  Otherwise, data on the persistent disk might be lost.
	//
	//  If the encryption key is revoked, the workstation session automatically
	//  stops within 7 hours.
	//
	//  Immutable after the workstation configuration is created.
	EncryptionKey *WorkstationConfig_CustomerEncryptionKey `json:"encryptionKey,omitempty"`

	// Optional. Readiness checks to perform when starting a workstation using
	//  this workstation configuration. Mark a workstation as running only after
	//  all specified readiness checks return 200 status codes.
	ReadinessChecks []WorkstationConfig_ReadinessCheck `json:"readinessChecks,omitempty"`

	// Optional. Immutable. Specifies the zones used to replicate the VM and disk
	//  resources within the region. If set, exactly two zones within the
	//  workstation cluster's region must be specified—for example,
	//  `['us-central1-a', 'us-central1-f']`. If this field is empty, two default
	//  zones within the region are used.
	//
	//  Immutable after the workstation configuration is created.
	ReplicaZones []string `json:"replicaZones,omitempty"`

	// Output only. Whether this resource is degraded, in which case it may
	//  require user action to restore full functionality. See also the
	//  [conditions][google.cloud.workstations.v1.WorkstationConfig.conditions]
	//  field.
	Degraded *bool `json:"degraded,omitempty"`

	// Output only. Status conditions describing the current resource state.
	Conditions []Status `json:"conditions,omitempty"`
}

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
	DisablePublicIpAddresses *bool `json:"disablePublicIpAddresses,omitempty"`

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

// +kcc:proto=google.protobuf.Any
type Any struct {
	// A URL/resource name that uniquely identifies the type of the serialized
	//  protocol buffer message. This string must contain at least
	//  one "/" character. The last segment of the URL's path must represent
	//  the fully qualified name of the type (as in
	//  `path/google.protobuf.Duration`). The name should be in a canonical form
	//  (e.g., leading "." is not accepted).
	//
	//  In practice, teams usually precompile into the binary all types that they
	//  expect it to use in the context of Any. However, for URLs which use the
	//  scheme `http`, `https`, or no scheme, one can optionally set up a type
	//  server that maps type URLs to message definitions as follows:
	//
	//  * If no scheme is provided, `https` is assumed.
	//  * An HTTP GET on the URL must yield a [google.protobuf.Type][]
	//    value in binary format, or produce an error.
	//  * Applications are allowed to cache lookup results based on the
	//    URL, or have them precompiled into a binary to avoid any
	//    lookup. Therefore, binary compatibility needs to be preserved
	//    on changes to types. (Use versioned type names to manage
	//    breaking changes.)
	//
	//  Note: this functionality is not currently available in the official
	//  protobuf release, and it is not used for type URLs beginning with
	//  type.googleapis.com.
	//
	//  Schemes other than `http`, `https` (or the empty scheme) might be
	//  used with implementation specific semantics.
	TypeURL *string `json:"typeURL,omitempty"`

	// Must be a valid serialized protocol buffer of the above specified type.
	Value []byte `json:"value,omitempty"`
}
