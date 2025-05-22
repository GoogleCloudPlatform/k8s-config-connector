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
	kmsv1beta1 "github.com/GoogleCloudPlatform/k8s-config-connector/apis/kms/v1beta1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/apis/refs/v1beta1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/apis/k8s/v1alpha1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

var NotebookInstanceGVK = GroupVersion.WithKind("NotebookInstance")

// NotebookInstanceSpec defines the desired state of NotebookInstance
// +kcc:spec:proto=google.cloud.notebooks.v1.Instance
type NotebookInstanceSpec struct {

	// Use a Compute Engine VM image to start the notebook instance.
	// +kcc:proto:field=google.cloud.notebooks.v1.Instance.vm_image
	VMImage *VMImage `json:"vmImage,omitempty"`

	// Use a container image to start the notebook instance.
	// +kcc:proto:field=google.cloud.notebooks.v1.Instance.container_image
	ContainerImage *ContainerImage `json:"containerImage,omitempty"`

	// Path to a Bash script that automatically runs after a notebook instance
	//  fully boots up. The path must be a URL or
	//  Cloud Storage path (`gs://path-to-file/file-name`).
	// +kcc:proto:field=google.cloud.notebooks.v1.Instance.post_startup_script
	PostStartupScript *string `json:"postStartupScript,omitempty"`

	// Input only. The owner of this instance after creation. Format: `alias@example.com`
	//
	//  Currently supports one owner only. If not specified, all of the service
	//  account users of your VM instance's service account can use
	//  the instance.
	// +kcc:proto:field=google.cloud.notebooks.v1.Instance.instance_owners
	InstanceOwners []string `json:"instanceOwners,omitempty"`

	// The service account on this instance, giving access to other Google
	//  Cloud services.
	//  You can use any service account within the same project, but you
	//  must have the service account user permission to use the instance.
	//
	//  If not specified, the [Compute Engine default service
	//  account](https://cloud.google.com/compute/docs/access/service-accounts#default_service_account)
	//  is used.
	// +kcc:proto:field=google.cloud.notebooks.v1.Instance.service_account
	ServiceAccountRef *v1beta1.IAMServiceAccountRef `json:"serviceAccountRef,omitempty"`

	// Optional. The URIs of service account scopes to be included in
	//  Compute Engine instances.
	//
	//  If not specified, the following
	//  [scopes](https://cloud.google.com/compute/docs/access/service-accounts#accesscopesiam)
	//  are defined:
	//   - https://www.googleapis.com/auth/cloud-platform
	//   - https://www.googleapis.com/auth/userinfo.email
	//  If not using default scopes, you need at least:
	//     https://www.googleapis.com/auth/compute
	// +kcc:proto:field=google.cloud.notebooks.v1.Instance.service_account_scopes
	ServiceAccountScopes []string `json:"serviceAccountScopes,omitempty"`

	// Required. The [Compute Engine machine
	//  type](https://cloud.google.com/compute/docs/machine-types) of this
	//  instance.
	// +required.
	// +kcc:proto:field=google.cloud.notebooks.v1.Instance.machine_type
	MachineType *string `json:"machineType,omitempty"`

	// The hardware accelerator used on this instance. If you use
	//  accelerators, make sure that your configuration has
	//  [enough vCPUs and memory to support the `machine_type` you have
	//  selected](https://cloud.google.com/compute/docs/gpus/#gpus-list).
	// +kcc:proto:field=google.cloud.notebooks.v1.Instance.accelerator_config
	AcceleratorConfig *Instance_AcceleratorConfig `json:"acceleratorConfig,omitempty"`

	// Whether the end user authorizes Google Cloud to install GPU driver
	//  on this instance.
	//  If this field is empty or set to false, the GPU driver won't be installed.
	//  Only applicable to instances with GPUs.
	// +kcc:proto:field=google.cloud.notebooks.v1.Instance.install_gpu_driver
	InstallGpuDriver *bool `json:"installGpuDriver,omitempty"`

	// Specify a custom Cloud Storage path where the GPU driver is stored.
	//  If not specified, we'll automatically choose from official GPU drivers.
	// +kcc:proto:field=google.cloud.notebooks.v1.Instance.custom_gpu_driver_path
	CustomGpuDriverPath *string `json:"customGpuDriverPath,omitempty"`

	// Input only. The type of the boot disk attached to this instance, defaults to
	//  standard persistent disk (`PD_STANDARD`).
	// +kcc:proto:field=google.cloud.notebooks.v1.Instance.boot_disk_type
	BootDiskType *string `json:"bootDiskType,omitempty"`

	// Input only. The size of the boot disk in GB attached to this instance, up to a maximum
	//  of 64000 GB (64 TB). The minimum recommended value is 100 GB. If not
	//  specified, this defaults to 100.
	// +kcc:proto:field=google.cloud.notebooks.v1.Instance.boot_disk_size_gb
	BootDiskSizeGB *int64 `json:"bootDiskSizeGB,omitempty"`

	// Input only. The type of the data disk attached to this instance, defaults to
	//  standard persistent disk (`PD_STANDARD`).
	// +kcc:proto:field=google.cloud.notebooks.v1.Instance.data_disk_type
	DataDiskType *string `json:"dataDiskType,omitempty"`

	// Input only. The size of the data disk in GB attached to this instance, up to a maximum
	//  of 64000 GB (64 TB). You can choose the size of the data disk based on how
	//  big your notebooks and data are. If not specified, this defaults to 100.
	// +kcc:proto:field=google.cloud.notebooks.v1.Instance.data_disk_size_gb
	DataDiskSizeGB *int64 `json:"dataDiskSizeGB,omitempty"`

	// Input only. If true, the data disk will not be auto deleted when deleting the instance.
	// +kcc:proto:field=google.cloud.notebooks.v1.Instance.no_remove_data_disk
	NoRemoveDataDisk *bool `json:"noRemoveDataDisk,omitempty"`

	// Input only. Disk encryption method used on the boot and data disks, defaults to GMEK.
	// +kcc:proto:field=google.cloud.notebooks.v1.Instance.disk_encryption
	DiskEncryption *string `json:"diskEncryption,omitempty"`

	// Input only. The KMS key used to encrypt the disks, only applicable if disk_encryption
	//  is CMEK.
	//  Learn more about [using your own encryption keys](/kms/docs/quickstart).
	// +kcc:proto:field=google.cloud.notebooks.v1.Instance.kms_key
	KMSKeyRef *kmsv1beta1.KMSKeyRef_OneOf `json:"kmsKeyRef,omitempty"`

	// Optional. Shielded VM configuration.
	//  [Images using supported Shielded VM
	//  features](https://cloud.google.com/compute/docs/instances/modifying-shielded-vm).
	// +kcc:proto:field=google.cloud.notebooks.v1.Instance.shielded_instance_config
	ShieldedInstanceConfig *Instance_ShieldedInstanceConfig `json:"shieldedInstanceConfig,omitempty"`

	// If true, no public IP will be assigned to this instance.
	// +kcc:proto:field=google.cloud.notebooks.v1.Instance.no_public_ip
	NoPublicIP *bool `json:"noPublicIP,omitempty"`

	// If true, the notebook instance will not register with the proxy.
	// +kcc:proto:field=google.cloud.notebooks.v1.Instance.no_proxy_access
	NoProxyAccess *bool `json:"noProxyAccess,omitempty"`

	// The name of the VPC that this instance is in.
	// +kcc:proto:field=google.cloud.notebooks.v1.Instance.network
	NetworkRef *v1beta1.ComputeNetworkRef `json:"networkRef,omitempty"`

	// The name of the subnet that this instance is in.
	// +kcc:proto:field=google.cloud.notebooks.v1.Instance.subnet
	SubnetRef *v1beta1.ComputeSubnetworkRef `json:"subnetRef,omitempty"`

	// Labels to apply to this instance.
	//  These can be later modified by the setLabels method.
	// +kcc:proto:field=google.cloud.notebooks.v1.Instance.labels
	Labels map[string]string `json:"labels,omitempty"`

	// Custom metadata to apply to this instance.
	// +kcc:proto:field=google.cloud.notebooks.v1.Instance.metadata
	Metadata map[string]string `json:"metadata,omitempty"`

	// Optional. The Compute Engine tags to add to runtime (see [Tagging
	//  instances](https://cloud.google.com/compute/docs/label-or-tag-resources#tags)).
	// +kcc:proto:field=google.cloud.notebooks.v1.Instance.tags
	Tags []string `json:"tags,omitempty"`

	// The upgrade history of this instance.
	// +kcc:proto:field=google.cloud.notebooks.v1.Instance.upgrade_history
	UpgradeHistory []Instance_UpgradeHistoryEntry `json:"upgradeHistory,omitempty"`

	// Optional. The type of vNIC to be used on this interface. This may be gVNIC or
	//  VirtioNet.
	// +kcc:proto:field=google.cloud.notebooks.v1.Instance.nic_type
	NicType *string `json:"nicType,omitempty"`

	// Optional. The optional reservation affinity. Setting this field will apply
	//  the specified [Zonal Compute
	//  Reservation](https://cloud.google.com/compute/docs/instances/reserving-zonal-resources)
	//  to this notebook instance.
	// +kcc:proto:field=google.cloud.notebooks.v1.Instance.reservation_affinity
	ReservationAffinity *ReservationAffinity `json:"reservationAffinity,omitempty"`

	// Optional. Flag to enable ip forwarding or not, default false/off.
	//  https://cloud.google.com/vpc/docs/using-routes#canipforward
	// +kcc:proto:field=google.cloud.notebooks.v1.Instance.can_ip_forward
	CanIPForward *bool `json:"canIPForward,omitempty"`

	// Required. The parent resource name where the Job will be created. Pattern: "projects/{project}/locations/{location}"
	*Parent `json:",inline"`

	// The NotebookInstance name. If not given, the metadata.name will be used.
	ResourceID *string `json:"resourceID,omitempty"`
}

type Parent struct {
	// Immutable. The location where the notebook instance should reside.
	// +required
	Zone string `json:"zone,omitempty"`

	// The project that this resource belongs to.
	// +required
	ProjectRef *v1beta1.ProjectRef `json:"projectRef,omitempty"`
}

// NotebookInstanceStatus defines the config connector machine state of NotebookInstance
type NotebookInstanceStatus struct {
	/* Conditions represent the latest available observations of the
	   object's current state. */
	Conditions []v1alpha1.Condition `json:"conditions,omitempty"`

	// ObservedGeneration is the generation of the resource that was most recently observed by the Config Connector controller. If this is equal to metadata.generation, then that means that the current reported status reflects the most recent desired state of the resource.
	ObservedGeneration *int64 `json:"observedGeneration,omitempty"`

	// A unique specifier for the NotebookInstance resource in GCP.
	ExternalRef *string `json:"externalRef,omitempty"`

	// ObservedState is the state of the resource as most recently observed in GCP.
	ObservedState *NotebookInstanceObservedState `json:"observedState,omitempty"`
}

// NotebookInstanceObservedState is the state of the NotebookInstance resource as most recently observed in GCP.
// +kcc:observedstate:proto=google.cloud.notebooks.v1.Instance
type NotebookInstanceObservedState struct {
	// Output only. The name of this notebook instance. Format:
	//  `projects/{project_id}/locations/{location}/instances/{instance_id}`
	// +kcc:proto:field=google.cloud.notebooks.v1.Instance.name
	Name *string `json:"name,omitempty"`

	// Output only. The proxy endpoint that is used to access the Jupyter notebook.
	// +kcc:proto:field=google.cloud.notebooks.v1.Instance.proxy_uri
	ProxyURI *string `json:"proxyURI,omitempty"`

	// Output only. The state of this instance.
	// +kcc:proto:field=google.cloud.notebooks.v1.Instance.state
	State *string `json:"state,omitempty"`

	// Output only. Attached disks to notebook instance.
	// +kcc:proto:field=google.cloud.notebooks.v1.Instance.disks
	Disks []Instance_Disk `json:"disks,omitempty"`

	// Output only. Email address of entity that sent original CreateInstance request.
	// +kcc:proto:field=google.cloud.notebooks.v1.Instance.creator
	Creator *string `json:"creator,omitempty"`

	// Output only. Instance creation time.
	// +kcc:proto:field=google.cloud.notebooks.v1.Instance.create_time
	CreateTime *string `json:"createTime,omitempty"`

	// Output only. Instance update time.
	// +kcc:proto:field=google.cloud.notebooks.v1.Instance.update_time
	UpdateTime *string `json:"updateTime,omitempty"`
}

// +genclient
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:resource:categories=gcp,shortName=gcpnotebookinstance;gcpnotebookinstances
// +kubebuilder:subresource:status
// +kubebuilder:metadata:labels="cnrm.cloud.google.com/managed-by-kcc=true";"cnrm.cloud.google.com/system=true"
// +kubebuilder:printcolumn:name="Age",JSONPath=".metadata.creationTimestamp",type="date"
// +kubebuilder:printcolumn:name="Ready",JSONPath=".status.conditions[?(@.type=='Ready')].status",type="string",description="When 'True', the most recent reconcile of the resource succeeded"
// +kubebuilder:printcolumn:name="Status",JSONPath=".status.conditions[?(@.type=='Ready')].reason",type="string",description="The reason for the value in 'Ready'"
// +kubebuilder:printcolumn:name="Status Age",JSONPath=".status.conditions[?(@.type=='Ready')].lastTransitionTime",type="date",description="The last transition time for the value in 'Status'"

// NotebookInstance is the Schema for the NotebookInstance API
// +k8s:openapi-gen=true
type NotebookInstance struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	// +required
	Spec   NotebookInstanceSpec   `json:"spec,omitempty"`
	Status NotebookInstanceStatus `json:"status,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// NotebookInstanceList contains a list of NotebookInstance
type NotebookInstanceList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []NotebookInstance `json:"items"`
}

func init() {
	SchemeBuilder.Register(&NotebookInstance{}, &NotebookInstanceList{})
}
