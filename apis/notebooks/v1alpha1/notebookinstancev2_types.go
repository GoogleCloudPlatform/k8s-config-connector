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
	computerefs "github.com/GoogleCloudPlatform/k8s-config-connector/apis/compute/refs"
	computev1beta1 "github.com/GoogleCloudPlatform/k8s-config-connector/apis/compute/v1beta1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/apis/refs/v1beta1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/apis/k8s/v1alpha1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

var NotebookInstanceV2GVK = GroupVersion.WithKind("NotebookInstanceV2")

// NotebookInstanceV2Spec defines the desired state of NotebookInstanceV2
// +kcc:spec:proto=google.cloud.notebooks.v2.Instance
type NotebookInstanceV2Spec struct {
	// The project that this resource belongs to.
	ProjectRef *v1beta1.ProjectRef `json:"projectRef"`

	// The location of this resource.
	Location *string `json:"location"`

	// The NotebookInstanceV2 name. If not given, the metadata.name will be used.
	ResourceID *string `json:"resourceID,omitempty"`

	// Optional. Compute Engine setup for the notebook. Uses notebook-defined fields.
	// +kcc:proto:field=google.cloud.notebooks.v2.Instance.gce_setup
	GCESetup *NotebookInstanceV2_GCESetup `json:"gceSetup,omitempty"`

	// Optional. Input only. The owner of this instance after creation. Format: `alias@example.com`
	// Currently supports one owner only. If not specified, all of the service
	// account users of your VM instance's service account can use the instance.
	// +kcc:proto:field=google.cloud.notebooks.v2.Instance.instance_owners
	InstanceOwners []string `json:"instanceOwners,omitempty"`

	// Optional. If true, the notebook instance will not register with the proxy.
	// +kcc:proto:field=google.cloud.notebooks.v2.Instance.disable_proxy_access
	DisableProxyAccess *bool `json:"disableProxyAccess,omitempty"`

	// Optional. Labels to apply to this instance.
	// These can be later modified by the UpdateInstance method.
	// +kcc:proto:field=google.cloud.notebooks.v2.Instance.labels
	Labels map[string]string `json:"labels,omitempty"`
}

// +kcc:proto=google.cloud.notebooks.v2.GceSetup
type NotebookInstanceV2_GCESetup struct {
	// Optional. The machine type of the VM instance.
	// https://cloud.google.com/compute/docs/machine-resource
	// +kcc:proto:field=google.cloud.notebooks.v2.GceSetup.machine_type
	MachineType *string `json:"machineType,omitempty"`

	// Optional. The hardware accelerators used on this instance. If you use
	// accelerators, make sure that your configuration has enough vCPUs and memory to support the `machine_type` you have selected.
	// Currently supports only one accelerator configuration.
	// +kcc:proto:field=google.cloud.notebooks.v2.GceSetup.accelerator_configs
	AcceleratorConfigs []NotebookInstanceV2_AcceleratorConfig `json:"acceleratorConfigs,omitempty"`

	// Optional. The service account that serves as an identity for the VM instance. Currently supports only one service account.
	// +kcc:proto:field=google.cloud.notebooks.v2.GceSetup.service_accounts
	ServiceAccounts []NotebookInstanceV2_ServiceAccount `json:"serviceAccounts,omitempty"`

	// Optional. Use a Compute Engine VM image to start the notebook instance.
	// +kcc:proto:field=google.cloud.notebooks.v2.GceSetup.vm_image
	VMImage *NotebookInstanceV2_VMImage `json:"vmImage,omitempty"`

	// Optional. Use a container image to start the notebook instance.
	// +kcc:proto:field=google.cloud.notebooks.v2.GceSetup.container_image
	ContainerImage *NotebookInstanceV2_ContainerImage `json:"containerImage,omitempty"`

	// Optional. The boot disk for the VM.
	// +kcc:proto:field=google.cloud.notebooks.v2.GceSetup.boot_disk
	BootDisk *NotebookInstanceV2_BootDisk `json:"bootDisk,omitempty"`

	// Optional. Data disks attached to the VM instance. Currently supports only one data disk.
	// +kcc:proto:field=google.cloud.notebooks.v2.GceSetup.data_disks
	DataDisks []NotebookInstanceV2_DataDisk `json:"dataDisks,omitempty"`

	// Optional. Shielded VM configuration.
	// [Images using supported Shielded VM features](https://cloud.google.com/compute/docs/instances/modifying-shielded-vm).
	// +kcc:proto:field=google.cloud.notebooks.v2.GceSetup.shielded_instance_config
	ShieldedInstanceConfig *NotebookInstanceV2_ShieldedInstanceConfig `json:"shieldedInstanceConfig,omitempty"`

	// Optional. The network interfaces for the VM. Supports only one interface.
	// +kcc:proto:field=google.cloud.notebooks.v2.GceSetup.network_interfaces
	NetworkInterfaces []NotebookInstanceV2_NetworkInterface `json:"networkInterfaces,omitempty"`

	// Optional. If true, no external IP will be assigned to this VM instance.
	// +kcc:proto:field=google.cloud.notebooks.v2.GceSetup.disable_public_ip
	DisablePublicIP *bool `json:"disablePublicIP,omitempty"`

	// Optional. The Compute Engine tags to add to runtime (see [Tagging instances](https://cloud.google.com/compute/docs/label-or-tag-resources#tags)).
	// +kcc:proto:field=google.cloud.notebooks.v2.GceSetup.tags
	Tags []string `json:"tags,omitempty"`

	// Optional. Custom metadata to apply to this instance.
	// +kcc:proto:field=google.cloud.notebooks.v2.GceSetup.metadata
	Metadata map[string]string `json:"metadata,omitempty"`

	// Optional. Flag to enable ip forwarding or not, default false/off.
	// https://cloud.google.com/vpc/docs/using-routes#canipforward
	// +kcc:proto:field=google.cloud.notebooks.v2.GceSetup.enable_ip_forwarding
	EnableIPForwarding *bool `json:"enableIPForwarding,omitempty"`

	// Optional. Configuration for GPU drivers.
	// +kcc:proto:field=google.cloud.notebooks.v2.GceSetup.gpu_driver_config
	GpuDriverConfig *NotebookInstanceV2_GPUDriverConfig `json:"gpuDriverConfig,omitempty"`
}

// +kcc:proto=google.cloud.notebooks.v2.AcceleratorConfig
type NotebookInstanceV2_AcceleratorConfig struct {
	// Optional. Type of this accelerator.
	// +kcc:proto:field=google.cloud.notebooks.v2.AcceleratorConfig.type
	Type *string `json:"type,omitempty"`

	// Optional. Count of cores of this accelerator.
	// +kcc:proto:field=google.cloud.notebooks.v2.AcceleratorConfig.core_count
	CoreCount *int64 `json:"coreCount,omitempty"`
}

// +kcc:proto=google.cloud.notebooks.v2.ServiceAccount
type NotebookInstanceV2_ServiceAccount struct {
	// Optional. Email address of the service account.
	// +kcc:proto:field=google.cloud.notebooks.v2.ServiceAccount.email
	ServiceAccountRef *v1beta1.IAMServiceAccountRef `json:"serviceAccountRef,omitempty"`
}

// +kcc:proto=google.cloud.notebooks.v2.VmImage
type NotebookInstanceV2_VMImage struct {
	// Required. The name of the Google Cloud project that this VM image belongs to. Format: `{project_id}`
	// +kcc:proto:field=google.cloud.notebooks.v2.VmImage.project
	Project *string `json:"project,omitempty"`

	// Optional. Use VM image name to find the image.
	// +kcc:proto:field=google.cloud.notebooks.v2.VmImage.name
	Name *string `json:"name,omitempty"`

	// Optional. Use this VM image family to find the image; the newest image in this family will be used.
	// +kcc:proto:field=google.cloud.notebooks.v2.VmImage.family
	Family *string `json:"family,omitempty"`
}

// +kcc:proto=google.cloud.notebooks.v2.ContainerImage
type NotebookInstanceV2_ContainerImage struct {
	// Required. The path to the container image repository. For example: `gcr.io/{project_id}/{image_name}`
	// +kcc:proto:field=google.cloud.notebooks.v2.ContainerImage.repository
	Repository *string `json:"repository,omitempty"`

	// Optional. The tag of the container image. If not specified, this defaults to the latest tag.
	// +kcc:proto:field=google.cloud.notebooks.v2.ContainerImage.tag
	Tag *string `json:"tag,omitempty"`
}

// +kcc:proto=google.cloud.notebooks.v2.BootDisk
type NotebookInstanceV2_BootDisk struct {
	// Optional. The size of the boot disk in GB attached to this instance, up to a maximum of 64000 GB (64 TB). If not specified, this defaults to the recommended value of 150GB.
	// +kcc:proto:field=google.cloud.notebooks.v2.BootDisk.disk_size_gb
	DiskSizeGB *int64 `json:"diskSizeGB,omitempty"`

	// Optional. Indicates the type of the disk.
	// +kcc:proto:field=google.cloud.notebooks.v2.BootDisk.disk_type
	DiskType *string `json:"diskType,omitempty"`

	// Optional. Input only. Disk encryption method used on the boot and data disks, defaults to GMEK.
	// +kcc:proto:field=google.cloud.notebooks.v2.BootDisk.disk_encryption
	DiskEncryption *string `json:"diskEncryption,omitempty"`

	// Optional. Input only. The KMS key used to encrypt the disks, only applicable if disk_encryption is CMEK.
	// +kcc:proto:field=google.cloud.notebooks.v2.BootDisk.kms_key
	KMSKeyRef *v1beta1.KMSCryptoKeyRef `json:"kmsKeyRef,omitempty"`
}

// +kcc:proto=google.cloud.notebooks.v2.DataDisk
type NotebookInstanceV2_DataDisk struct {
	// Optional. The size of the disk in GB attached to this VM instance, up to a maximum of 64000 GB (64 TB). If not specified, this defaults to 100.
	// +kcc:proto:field=google.cloud.notebooks.v2.DataDisk.disk_size_gb
	DiskSizeGB *int64 `json:"diskSizeGB,omitempty"`

	// Optional. Input only. Indicates the type of the disk.
	// +kcc:proto:field=google.cloud.notebooks.v2.DataDisk.disk_type
	DiskType *string `json:"diskType,omitempty"`

	// Optional. Input only. Disk encryption method used on the boot and data disks, defaults to GMEK.
	// +kcc:proto:field=google.cloud.notebooks.v2.DataDisk.disk_encryption
	DiskEncryption *string `json:"diskEncryption,omitempty"`

	// Optional. Input only. The KMS key used to encrypt the disks, only applicable if disk_encryption is CMEK.
	// +kcc:proto:field=google.cloud.notebooks.v2.DataDisk.kms_key
	KMSKeyRef *v1beta1.KMSCryptoKeyRef `json:"kmsKeyRef,omitempty"`
}

// +kcc:proto=google.cloud.notebooks.v2.ShieldedInstanceConfig
type NotebookInstanceV2_ShieldedInstanceConfig struct {
	// Optional. Defines whether the VM instance has Secure Boot enabled. Secure Boot helps ensure that the system only runs authentic software by verifying the digital signature of all boot components, and halting the boot process if signature verification fails. Disabled by default.
	// +kcc:proto:field=google.cloud.notebooks.v2.ShieldedInstanceConfig.enable_secure_boot
	EnableSecureBoot *bool `json:"enableSecureBoot,omitempty"`

	// Optional. Defines whether the VM instance has the vTPM enabled. Enabled by default.
	// +kcc:proto:field=google.cloud.notebooks.v2.ShieldedInstanceConfig.enable_vtpm
	EnableVTPM *bool `json:"enableVTPM,omitempty"`

	// Optional. Defines whether the VM instance has integrity monitoring enabled.
	// Enables monitoring and attestation of the boot integrity of the VM instance. The attestation is performed against the integrity policy baseline. This baseline is initially derived from the implicitly trusted boot image when the VM instance is created. Enabled by default.
	// +kcc:proto:field=google.cloud.notebooks.v2.ShieldedInstanceConfig.enable_integrity_monitoring
	EnableIntegrityMonitoring *bool `json:"enableIntegrityMonitoring,omitempty"`
}

// +kcc:proto=google.cloud.notebooks.v2.NetworkInterface
type NotebookInstanceV2_NetworkInterface struct {
	// Optional. The name of the VPC that this VM instance is in.
	// +kcc:proto:field=google.cloud.notebooks.v2.NetworkInterface.network
	NetworkRef *computerefs.ComputeNetworkRef `json:"networkRef,omitempty"`

	// Optional. The name of the subnet that this VM instance is in.
	// +kcc:proto:field=google.cloud.notebooks.v2.NetworkInterface.subnet
	SubnetRef *computev1beta1.ComputeSubnetworkRef `json:"subnetRef,omitempty"`

	// Optional. The type of vNIC to be used on this interface. This may be gVNIC or VirtioNet.
	// +kcc:proto:field=google.cloud.notebooks.v2.NetworkInterface.nic_type
	NicType *string `json:"nicType,omitempty"`
}

// +kcc:proto=google.cloud.notebooks.v2.GPUDriverConfig
type NotebookInstanceV2_GPUDriverConfig struct {
	// Optional. Whether the end user authorizes Google Cloud to install GPU driver on this VM instance. If this field is empty or set to false, the GPU driver won't be installed. Only applicable to instances with GPUs.
	// +kcc:proto:field=google.cloud.notebooks.v2.GPUDriverConfig.enable_gpu_driver
	EnableGpuDriver *bool `json:"enableGpuDriver,omitempty"`

	// Optional. Specify a custom Cloud Storage path where the GPU driver is stored. If not specified, we'll automatically choose from official GPU drivers.
	// +kcc:proto:field=google.cloud.notebooks.v2.GPUDriverConfig.custom_gpu_driver_path
	CustomGpuDriverPath *string `json:"customGpuDriverPath,omitempty"`
}

// NotebookInstanceV2Status defines the config connector machine state of NotebookInstanceV2
type NotebookInstanceV2Status struct {
	/* Conditions represent the latest available observations of the
	   object's current state. */
	Conditions []v1alpha1.Condition `json:"conditions,omitempty"`

	// ObservedGeneration is the generation of the resource that was most recently observed by the Config Connector controller. If this is equal to metadata.generation, then that means that the current reported status reflects the most recent desired state of the resource.
	ObservedGeneration *int64 `json:"observedGeneration,omitempty"`

	// A unique specifier for the NotebookInstanceV2 resource in GCP.
	ExternalRef *string `json:"externalRef,omitempty"`

	// ObservedState is the state of the resource as most recently observed in GCP.
	ObservedState *NotebookInstanceV2ObservedState `json:"observedState,omitempty"`
}

// NotebookInstanceV2ObservedState is the state of the NotebookInstanceV2 resource as most recently observed in GCP.
// +kcc:observedstate:proto=google.cloud.notebooks.v2.Instance
type NotebookInstanceV2ObservedState struct {
	// Output only. The name of this notebook instance. Format: `projects/{project_id}/locations/{location}/instances/{instance_id}`
	// +kcc:proto:field=google.cloud.notebooks.v2.Instance.name
	Name *string `json:"name,omitempty"`

	// Optional. Compute Engine setup for the notebook. Uses notebook-defined fields.
	// +kcc:proto:field=google.cloud.notebooks.v2.Instance.gce_setup
	GCESetup *NotebookInstanceV2_GCESetupObservedState `json:"gceSetup,omitempty"`

	// Output only. The proxy endpoint that is used to access the Jupyter notebook.
	// +kcc:proto:field=google.cloud.notebooks.v2.Instance.proxy_uri
	ProxyURI *string `json:"proxyURI,omitempty"`

	// Output only. Email address of entity that sent original CreateInstance request.
	// +kcc:proto:field=google.cloud.notebooks.v2.Instance.creator
	Creator *string `json:"creator,omitempty"`

	// Output only. The state of this instance.
	// +kcc:proto:field=google.cloud.notebooks.v2.Instance.state
	State *string `json:"state,omitempty"`

	// Output only. The upgrade history of this instance.
	// +kcc:proto:field=google.cloud.notebooks.v2.Instance.upgrade_history
	UpgradeHistory []NotebookInstanceV2_UpgradeHistoryEntryObservedState `json:"upgradeHistory,omitempty"`

	// Output only. Unique ID of the resource.
	// +kcc:proto:field=google.cloud.notebooks.v2.Instance.id
	ID *string `json:"id,omitempty"`

	// Output only. Instance health_state.
	// +kcc:proto:field=google.cloud.notebooks.v2.Instance.health_state
	HealthState *string `json:"healthState,omitempty"`

	// Output only. Additional information about instance health.
	// +kcc:proto:field=google.cloud.notebooks.v2.Instance.health_info
	HealthInfo map[string]string `json:"healthInfo,omitempty"`

	// Output only. Instance creation time.
	// +kcc:proto:field=google.cloud.notebooks.v2.Instance.create_time
	CreateTime *string `json:"createTime,omitempty"`

	// Output only. Instance update time.
	// +kcc:proto:field=google.cloud.notebooks.v2.Instance.update_time
	UpdateTime *string `json:"updateTime,omitempty"`
}

// +kcc:observedstate:proto=google.cloud.notebooks.v2.GceSetup
type NotebookInstanceV2_GCESetupObservedState struct {
	// Optional. The service account that serves as an identity for the VM instance. Currently supports only one service account.
	// +kcc:proto:field=google.cloud.notebooks.v2.GceSetup.service_accounts
	ServiceAccounts []NotebookInstanceV2_ServiceAccountObservedState `json:"serviceAccounts,omitempty"`
}

// +kcc:observedstate:proto=google.cloud.notebooks.v2.ServiceAccount
type NotebookInstanceV2_ServiceAccountObservedState struct {
	// Output only. The list of scopes to be made available for this service account. Set by the CLH to https://www.googleapis.com/auth/cloud-platform
	// +kcc:proto:field=google.cloud.notebooks.v2.ServiceAccount.scopes
	Scopes []string `json:"scopes,omitempty"`
}

// +kcc:observedstate:proto=google.cloud.notebooks.v2.UpgradeHistoryEntry
type NotebookInstanceV2_UpgradeHistoryEntryObservedState struct {
	// Optional. The snapshot of the boot disk of this notebook instance before upgrade.
	// +kcc:proto:field=google.cloud.notebooks.v2.UpgradeHistoryEntry.snapshot
	Snapshot *string `json:"snapshot,omitempty"`

	// Optional. The VM image before this instance upgrade.
	// +kcc:proto:field=google.cloud.notebooks.v2.UpgradeHistoryEntry.vm_image
	VMImage *string `json:"vmImage,omitempty"`

	// Optional. The container image before this instance upgrade.
	// +kcc:proto:field=google.cloud.notebooks.v2.UpgradeHistoryEntry.container_image
	ContainerImage *string `json:"containerImage,omitempty"`

	// Optional. The framework of this notebook instance.
	// +kcc:proto:field=google.cloud.notebooks.v2.UpgradeHistoryEntry.framework
	Framework *string `json:"framework,omitempty"`

	// Optional. The version of the notebook instance before this upgrade.
	// +kcc:proto:field=google.cloud.notebooks.v2.UpgradeHistoryEntry.version
	Version *string `json:"version,omitempty"`

	// Output only. The state of this instance upgrade history entry.
	// +kcc:proto:field=google.cloud.notebooks.v2.UpgradeHistoryEntry.state
	State *string `json:"state,omitempty"`

	// Immutable. The time that this instance upgrade history entry is created.
	// +kcc:proto:field=google.cloud.notebooks.v2.UpgradeHistoryEntry.create_time
	CreateTime *string `json:"createTime,omitempty"`

	// Optional. Action. Rolloback or Upgrade.
	// +kcc:proto:field=google.cloud.notebooks.v2.UpgradeHistoryEntry.action
	Action *string `json:"action,omitempty"`

	// Optional. Target VM Version, like m63.
	// +kcc:proto:field=google.cloud.notebooks.v2.UpgradeHistoryEntry.target_version
	TargetVersion *string `json:"targetVersion,omitempty"`
}

// +genclient
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:resource:categories=gcp,shortName=gcpnotebookinstancev2;gcpnotebookinstancev2s
// +kubebuilder:subresource:status
// +kubebuilder:metadata:labels="cnrm.cloud.google.com/managed-by-kcc=true"
// +kubebuilder:metadata:labels="cnrm.cloud.google.com/system=true"
// +kubebuilder:metadata:labels="cnrm.cloud.google.com/stability-level=alpha"
// +kubebuilder:printcolumn:name="Age",JSONPath=".metadata.creationTimestamp",type="date"
// +kubebuilder:printcolumn:name="Ready",JSONPath=".status.conditions[?(@.type=='Ready')].status",type="string",description="When 'True', the most recent reconcile of the resource succeeded"
// +kubebuilder:printcolumn:name="Status",JSONPath=".status.conditions[?(@.type=='Ready')].reason",type="string",description="The reason for the value in 'Ready'"
// +kubebuilder:printcolumn:name="Status Age",JSONPath=".status.conditions[?(@.type=='Ready')].lastTransitionTime",type="date",description="The last transition time for the value in 'Status'"

// NotebookInstanceV2 is the Schema for the NotebookInstanceV2 API
// +k8s:openapi-gen=true
type NotebookInstanceV2 struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	// +required
	Spec   NotebookInstanceV2Spec   `json:"spec,omitempty"`
	Status NotebookInstanceV2Status `json:"status,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// NotebookInstanceV2List contains a list of NotebookInstanceV2
type NotebookInstanceV2List struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []NotebookInstanceV2 `json:"items"`
}

func init() {
	SchemeBuilder.Register(&NotebookInstanceV2{}, &NotebookInstanceV2List{})
}
