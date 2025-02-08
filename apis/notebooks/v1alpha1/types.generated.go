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


// +kcc:proto=google.cloud.notebooks.v1beta1.ContainerImage
type ContainerImage struct {
	// Required. The path to the container image repository. For example:
	//  `gcr.io/{project_id}/{image_name}`
	// +kcc:proto:field=google.cloud.notebooks.v1beta1.ContainerImage.repository
	Repository *string `json:"repository,omitempty"`

	// The tag of the container image. If not specified, this defaults
	//  to the latest tag.
	// +kcc:proto:field=google.cloud.notebooks.v1beta1.ContainerImage.tag
	Tag *string `json:"tag,omitempty"`
}

// +kcc:proto=google.cloud.notebooks.v1beta1.Instance
type Instance struct {

	// Use a Compute Engine VM image to start the notebook instance.
	// +kcc:proto:field=google.cloud.notebooks.v1beta1.Instance.vm_image
	VmImage *VmImage `json:"vmImage,omitempty"`

	// Use a container image to start the notebook instance.
	// +kcc:proto:field=google.cloud.notebooks.v1beta1.Instance.container_image
	ContainerImage *ContainerImage `json:"containerImage,omitempty"`

	// Path to a Bash script that automatically runs after a notebook instance
	//  fully boots up. The path must be a URL or
	//  Cloud Storage path (`gs://path-to-file/file-name`).
	// +kcc:proto:field=google.cloud.notebooks.v1beta1.Instance.post_startup_script
	PostStartupScript *string `json:"postStartupScript,omitempty"`

	// Input only. The owner of this instance after creation. Format: `alias@example.com`
	//
	//  Currently supports one owner only. If not specified, all of the service
	//  account users of your VM instance's service account can use
	//  the instance.
	// +kcc:proto:field=google.cloud.notebooks.v1beta1.Instance.instance_owners
	InstanceOwners []string `json:"instanceOwners,omitempty"`

	// The service account on this instance, giving access to other Google
	//  Cloud services.
	//  You can use any service account within the same project, but you
	//  must have the service account user permission to use the instance.
	//
	//  If not specified, the [Compute Engine default service
	//  account](https://cloud.google.com/compute/docs/access/service-accounts#default_service_account)
	//  is used.
	// +kcc:proto:field=google.cloud.notebooks.v1beta1.Instance.service_account
	ServiceAccount *string `json:"serviceAccount,omitempty"`

	// Required. The [Compute Engine machine
	//  type](https://cloud.google.com/compute/docs/machine-types) of this
	//  instance.
	// +kcc:proto:field=google.cloud.notebooks.v1beta1.Instance.machine_type
	MachineType *string `json:"machineType,omitempty"`

	// The hardware accelerator used on this instance. If you use
	//  accelerators, make sure that your configuration has
	//  [enough vCPUs and memory to support the `machine_type` you have
	//  selected](https://cloud.google.com/compute/docs/gpus/#gpus-list).
	// +kcc:proto:field=google.cloud.notebooks.v1beta1.Instance.accelerator_config
	AcceleratorConfig *Instance_AcceleratorConfig `json:"acceleratorConfig,omitempty"`

	// Whether the end user authorizes Google Cloud to install GPU driver
	//  on this instance.
	//  If this field is empty or set to false, the GPU driver won't be installed.
	//  Only applicable to instances with GPUs.
	// +kcc:proto:field=google.cloud.notebooks.v1beta1.Instance.install_gpu_driver
	InstallGpuDriver *bool `json:"installGpuDriver,omitempty"`

	// Specify a custom Cloud Storage path where the GPU driver is stored.
	//  If not specified, we'll automatically choose from official GPU drivers.
	// +kcc:proto:field=google.cloud.notebooks.v1beta1.Instance.custom_gpu_driver_path
	CustomGpuDriverPath *string `json:"customGpuDriverPath,omitempty"`

	// Input only. The type of the boot disk attached to this instance, defaults to
	//  standard persistent disk (`PD_STANDARD`).
	// +kcc:proto:field=google.cloud.notebooks.v1beta1.Instance.boot_disk_type
	BootDiskType *string `json:"bootDiskType,omitempty"`

	// Input only. The size of the boot disk in GB attached to this instance, up to a maximum
	//  of 64000 GB (64 TB). The minimum recommended value is 100 GB. If not
	//  specified, this defaults to 100.
	// +kcc:proto:field=google.cloud.notebooks.v1beta1.Instance.boot_disk_size_gb
	BootDiskSizeGB *int64 `json:"bootDiskSizeGB,omitempty"`

	// Input only. The type of the data disk attached to this instance, defaults to
	//  standard persistent disk (`PD_STANDARD`).
	// +kcc:proto:field=google.cloud.notebooks.v1beta1.Instance.data_disk_type
	DataDiskType *string `json:"dataDiskType,omitempty"`

	// Input only. The size of the data disk in GB attached to this instance, up to a maximum
	//  of 64000 GB (64 TB). You can choose the size of the data disk based on how
	//  big your notebooks and data are. If not specified, this defaults to 100.
	// +kcc:proto:field=google.cloud.notebooks.v1beta1.Instance.data_disk_size_gb
	DataDiskSizeGB *int64 `json:"dataDiskSizeGB,omitempty"`

	// Input only. If true, the data disk will not be auto deleted when deleting the instance.
	// +kcc:proto:field=google.cloud.notebooks.v1beta1.Instance.no_remove_data_disk
	NoRemoveDataDisk *bool `json:"noRemoveDataDisk,omitempty"`

	// Input only. Disk encryption method used on the boot and data disks, defaults to GMEK.
	// +kcc:proto:field=google.cloud.notebooks.v1beta1.Instance.disk_encryption
	DiskEncryption *string `json:"diskEncryption,omitempty"`

	// Input only. The KMS key used to encrypt the disks, only applicable if disk_encryption
	//  is CMEK.
	//  Format:
	//  `projects/{project_id}/locations/{location}/keyRings/{key_ring_id}/cryptoKeys/{key_id}`
	//
	//  Learn more about [using your own encryption
	//  keys](https://cloud.google.com/kms/docs/quickstart).
	// +kcc:proto:field=google.cloud.notebooks.v1beta1.Instance.kms_key
	KMSKey *string `json:"kmsKey,omitempty"`

	// If true, no public IP will be assigned to this instance.
	// +kcc:proto:field=google.cloud.notebooks.v1beta1.Instance.no_public_ip
	NoPublicIP *bool `json:"noPublicIP,omitempty"`

	// If true, the notebook instance will not register with the proxy.
	// +kcc:proto:field=google.cloud.notebooks.v1beta1.Instance.no_proxy_access
	NoProxyAccess *bool `json:"noProxyAccess,omitempty"`

	// The name of the VPC that this instance is in.
	//  Format:
	//  `projects/{project_id}/global/networks/{network_id}`
	// +kcc:proto:field=google.cloud.notebooks.v1beta1.Instance.network
	Network *string `json:"network,omitempty"`

	// The name of the subnet that this instance is in.
	//  Format:
	//  `projects/{project_id}/regions/{region}/subnetworks/{subnetwork_id}`
	// +kcc:proto:field=google.cloud.notebooks.v1beta1.Instance.subnet
	Subnet *string `json:"subnet,omitempty"`

	// Labels to apply to this instance.
	//  These can be later modified by the setLabels method.
	// +kcc:proto:field=google.cloud.notebooks.v1beta1.Instance.labels
	Labels map[string]string `json:"labels,omitempty"`

	// Custom metadata to apply to this instance.
	// +kcc:proto:field=google.cloud.notebooks.v1beta1.Instance.metadata
	Metadata map[string]string `json:"metadata,omitempty"`

	// Optional. The type of vNIC to be used on this interface. This may be gVNIC or
	//  VirtioNet.
	// +kcc:proto:field=google.cloud.notebooks.v1beta1.Instance.nic_type
	NicType *string `json:"nicType,omitempty"`

	// Optional. The optional reservation affinity. Setting this field will apply
	//  the specified [Zonal Compute
	//  Reservation](https://cloud.google.com/compute/docs/instances/reserving-zonal-resources)
	//  to this notebook instance.
	// +kcc:proto:field=google.cloud.notebooks.v1beta1.Instance.reservation_affinity
	ReservationAffinity *ReservationAffinity `json:"reservationAffinity,omitempty"`

	// Optional. Flag to enable ip forwarding or not, default false/off.
	//  https://cloud.google.com/vpc/docs/using-routes#canipforward
	// +kcc:proto:field=google.cloud.notebooks.v1beta1.Instance.can_ip_forward
	CanIPForward *bool `json:"canIPForward,omitempty"`
}

// +kcc:proto=google.cloud.notebooks.v1beta1.Instance.AcceleratorConfig
type Instance_AcceleratorConfig struct {
	// Type of this accelerator.
	// +kcc:proto:field=google.cloud.notebooks.v1beta1.Instance.AcceleratorConfig.type
	Type *string `json:"type,omitempty"`

	// Count of cores of this accelerator.
	// +kcc:proto:field=google.cloud.notebooks.v1beta1.Instance.AcceleratorConfig.core_count
	CoreCount *int64 `json:"coreCount,omitempty"`
}

// +kcc:proto=google.cloud.notebooks.v1beta1.ReservationAffinity
type ReservationAffinity struct {
	// Optional. Type of reservation to consume
	// +kcc:proto:field=google.cloud.notebooks.v1beta1.ReservationAffinity.consume_reservation_type
	ConsumeReservationType *string `json:"consumeReservationType,omitempty"`

	// Optional. Corresponds to the label key of reservation resource.
	// +kcc:proto:field=google.cloud.notebooks.v1beta1.ReservationAffinity.key
	Key *string `json:"key,omitempty"`

	// Optional. Corresponds to the label values of reservation resource.
	// +kcc:proto:field=google.cloud.notebooks.v1beta1.ReservationAffinity.values
	Values []string `json:"values,omitempty"`
}

// +kcc:proto=google.cloud.notebooks.v1beta1.VmImage
type VmImage struct {
	// Required. The name of the Google Cloud project that this VM image belongs to.
	//  Format: `projects/{project_id}`
	// +kcc:proto:field=google.cloud.notebooks.v1beta1.VmImage.project
	Project *string `json:"project,omitempty"`

	// Use VM image name to find the image.
	// +kcc:proto:field=google.cloud.notebooks.v1beta1.VmImage.image_name
	ImageName *string `json:"imageName,omitempty"`

	// Use this VM image family to find the image; the newest image in this
	//  family will be used.
	// +kcc:proto:field=google.cloud.notebooks.v1beta1.VmImage.image_family
	ImageFamily *string `json:"imageFamily,omitempty"`
}

// +kcc:proto=google.cloud.notebooks.v1beta1.Instance
type InstanceObservedState struct {
	// Output only. The name of this notebook instance. Format:
	//  `projects/{project_id}/locations/{location}/instances/{instance_id}`
	// +kcc:proto:field=google.cloud.notebooks.v1beta1.Instance.name
	Name *string `json:"name,omitempty"`

	// Output only. The proxy endpoint that is used to access the Jupyter notebook.
	// +kcc:proto:field=google.cloud.notebooks.v1beta1.Instance.proxy_uri
	ProxyURI *string `json:"proxyURI,omitempty"`

	// Output only. The state of this instance.
	// +kcc:proto:field=google.cloud.notebooks.v1beta1.Instance.state
	State *string `json:"state,omitempty"`

	// Output only. Instance creation time.
	// +kcc:proto:field=google.cloud.notebooks.v1beta1.Instance.create_time
	CreateTime *string `json:"createTime,omitempty"`

	// Output only. Instance update time.
	// +kcc:proto:field=google.cloud.notebooks.v1beta1.Instance.update_time
	UpdateTime *string `json:"updateTime,omitempty"`
}
