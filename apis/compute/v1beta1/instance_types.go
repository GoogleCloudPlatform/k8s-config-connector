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

package v1beta1

import (
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/apis/k8s/v1alpha1"
	v1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

var ComputeInstanceGVK = GroupVersion.WithKind("ComputeInstance")

// ComputeInstanceSpec defines the desired state of ComputeInstance
// +kcc:spec:proto=google.cloud.compute.v1.Instance
type ComputeInstanceSpec struct {
	// The ComputeInstance name. If not given, the metadata.name will be used.
	ResourceID *string `json:"resourceID,omitempty"`

	// Controls for advanced machine-related behavior features.
	// +kcc:proto:field=google.cloud.compute.v1.Instance.advanced_machine_features
	AdvancedMachineFeatures *ComputeInstanceAdvancedMachineFeatures `json:"advancedMachineFeatures,omitempty"`

	// List of disks attached to the instance.
	AttachedDisk []ComputeInstanceAttachedDisk `json:"attachedDisk,omitempty"`

	// Immutable. The boot disk for the instance.
	BootDisk *ComputeInstanceBootDisk `json:"bootDisk,omitempty"`

	// Whether sending and receiving of packets with non-matching source or destination IPs is allowed.
	// +kcc:proto:field=google.cloud.compute.v1.Instance.can_ip_forward
	CanIPForward *bool `json:"canIpForward,omitempty"`

	// Immutable. The Confidential VM config being used by the instance.  on_host_maintenance has to be set to TERMINATE or this will fail to create.
	// +kcc:proto:field=google.cloud.compute.v1.Instance.confidential_instance_config
	ConfidentialInstanceConfig *ComputeInstanceConfidentialInstanceConfig `json:"confidentialInstanceConfig,omitempty"`

	// Whether deletion protection is enabled on this instance.
	// +kcc:proto:field=google.cloud.compute.v1.Instance.deletion_protection
	DeletionProtection *bool `json:"deletionProtection,omitempty"`

	// Immutable. A brief description of the resource.
	// +kcc:proto:field=google.cloud.compute.v1.Instance.description
	Description *string `json:"description,omitempty"`

	// Desired status of the instance. Either "RUNNING" or "TERMINATED".
	DesiredStatus *string `json:"desiredStatus,omitempty"`

	// Whether the instance has virtual displays enabled.
	// +kcc:proto:field=google.cloud.compute.v1.Instance.display_device.enable_display
	EnableDisplay *bool `json:"enableDisplay,omitempty"`

	// Immutable. List of the type and count of accelerator cards attached to the instance.
	// +kcc:proto:field=google.cloud.compute.v1.Instance.guest_accelerators
	GuestAccelerator []ComputeInstanceGuestAccelerator `json:"guestAccelerator,omitempty"`

	// Immutable. A custom hostname for the instance. Must be a fully qualified DNS name and RFC-1035-valid. Valid format is a series of labels 1-63 characters long matching the regular expression [a-z]([-a-z0-9]*[a-z0-9]), concatenated with periods. The entire hostname must not exceed 253 characters. Changing this forces a new resource to be created.
	// +kcc:proto:field=google.cloud.compute.v1.Instance.hostname
	Hostname *string `json:"hostname,omitempty"`

	// +kcc:proto:field=google.cloud.compute.v1.Instance.source_instance_template
	InstanceTemplateRef *ComputeInstanceInstanceTemplateRef `json:"instanceTemplateRef,omitempty"`

	// The machine type to create.
	// +kcc:proto:field=google.cloud.compute.v1.Instance.machine_type
	MachineType *string `json:"machineType,omitempty"`

	// Metadata key/value pairs to make available from within the instance.
	Metadata []ComputeInstanceMetadata `json:"metadata,omitempty"`

	// Immutable. Metadata startup scripts made available within the instance.
	MetadataStartupScript *string `json:"metadataStartupScript,omitempty"`

	// The minimum CPU platform specified for the VM instance.
	// +kcc:proto:field=google.cloud.compute.v1.Instance.min_cpu_platform
	MinCPUPlatform *string `json:"minCpuPlatform,omitempty"`

	// Immutable. The networks attached to the instance.
	NetworkInterface []ComputeInstanceNetworkInterface `json:"networkInterface,omitempty"`

	// Immutable. Configures network performance settings for the instance. If not specified, the instance will be created with its default network performance configuration.
	// +kcc:proto:field=google.cloud.compute.v1.Instance.network_performance_config
	NetworkPerformanceConfig *NetworkPerformanceConfig `json:"networkPerformanceConfig,omitempty"`

	// Immutable. Stores additional params passed with the request, but not persisted as part of resource payload.
	// +kcc:proto:field=google.cloud.compute.v1.Instance.params
	Params *InstanceParams `json:"params,omitempty"`

	// Immutable. Specifies the reservations that this instance can consume from.
	// +kcc:proto:field=google.cloud.compute.v1.Instance.reservation_affinity
	ReservationAffinity *ReservationAffinity `json:"reservationAffinity,omitempty"`

	// Resource policies applied to this instance.
	// +kcc:proto:field=google.cloud.compute.v1.Instance.resource_policies
	ResourcePolicies []ComputeInstanceResourcePolicyRef `json:"resourcePolicies,omitempty"`

	// The scheduling strategy being used by the instance.
	// +kcc:proto:field=google.cloud.compute.v1.Instance.scheduling
	Scheduling *Scheduling `json:"scheduling,omitempty"`

	// Immutable. The scratch disks attached to the instance.
	ScratchDisk []ComputeInstanceScratchDisk `json:"scratchDisk,omitempty"`

	// The service account to attach to the instance.
	ServiceAccount *ComputeInstanceServiceAccount `json:"serviceAccount,omitempty"`

	// The shielded vm config being used by the instance.
	// +kcc:proto:field=google.cloud.compute.v1.Instance.shielded_instance_config
	ShieldedInstanceConfig *ComputeInstanceShieldedInstanceConfig `json:"shieldedInstanceConfig,omitempty"`

	// The list of tags attached to the instance.
	Tags []string `json:"tags,omitempty"`

	// Immutable. The zone of the instance. If self_link is provided, this value is ignored. If neither self_link nor zone are provided, the provider zone is used.
	// +kcc:proto:field=google.cloud.compute.v1.Instance.zone
	Zone *string `json:"zone,omitempty"`
}

// ComputeInstanceStatus defines the config connector machine state of ComputeInstance
type ComputeInstanceStatus struct {
	/* Conditions represent the latest available observations of the
	   object's current state. */
	Conditions []v1alpha1.Condition `json:"conditions,omitempty"`

	// ObservedGeneration is the generation of the resource that was most recently observed by the Config Connector controller. If this is equal to metadata.generation, then that means that the current reported status reflects the most recent desired state of the resource.
	ObservedGeneration *int64 `json:"observedGeneration,omitempty"`

	// A unique specifier for the ComputeInstance resource in GCP.
	ExternalRef *string `json:"externalRef,omitempty"`

	// ObservedState is the state of the resource as most recently observed in GCP.
	ObservedState *ComputeInstanceObservedState `json:"observedState,omitempty"`

	// The CPU platform used by this instance.
	// +kcc:proto:field=google.cloud.compute.v1.Instance.cpu_platform
	CpuPlatform *string `json:"cpuPlatform,omitempty"`

	// Current status of the instance.
	// This could be one of the following values: PROVISIONING, STAGING, RUNNING, STOPPING, SUSPENDING, SUSPENDED, REPAIRING, and TERMINATED.
	// For more information about the status of the instance, see [Instance life cycle](https://cloud.google.com/compute/docs/instances/instance-life-cycle).
	// +kcc:proto:field=google.cloud.compute.v1.Instance.status
	CurrentStatus *string `json:"currentStatus,omitempty"`

	// The server-assigned unique identifier of this instance.
	// +kcc:proto:field=google.cloud.compute.v1.Instance.id
	InstanceId *string `json:"instanceId,omitempty"`

	// The unique fingerprint of the labels.
	// +kcc:proto:field=google.cloud.compute.v1.Instance.label_fingerprint
	LabelFingerprint *string `json:"labelFingerprint,omitempty"`

	// The unique fingerprint of the metadata.
	// +kcc:proto:field=google.cloud.compute.v1.Instance.metadata.fingerprint
	MetadataFingerprint *string `json:"metadataFingerprint,omitempty"`

	// The URI of the created resource.
	// +kcc:proto:field=google.cloud.compute.v1.Instance.self_link
	SelfLink *string `json:"selfLink,omitempty"`

	// The unique fingerprint of the tags.
	// +kcc:proto:field=google.cloud.compute.v1.Instance.tags.fingerprint
	TagsFingerprint *string `json:"tagsFingerprint,omitempty"`
}

// ComputeInstanceObservedState is the state of the ComputeInstance resource as most recently observed in GCP.
// +kcc:observedstate:proto=google.cloud.compute.v1.Instance
type ComputeInstanceObservedState struct {
}

// +genclient
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:resource:categories=gcp,shortName=gcpcomputeinstance;gcpcomputeinstances
// +kubebuilder:subresource:status
// +kubebuilder:metadata:labels="cnrm.cloud.google.com/managed-by-kcc=true";"cnrm.cloud.google.com/system=true";"cnrm.cloud.google.com/stability-level=stable";"cnrm.cloud.google.com/tf2crd=true"
// +kubebuilder:printcolumn:name="Age",JSONPath=".metadata.creationTimestamp",type="date"
// +kubebuilder:printcolumn:name="Ready",JSONPath=".status.conditions[?(@.type=='Ready')].status",type="string",description="When 'True', the most recent reconcile of the resource succeeded"
// +kubebuilder:printcolumn:name="Status",JSONPath=".status.conditions[?(@.type=='Ready')].reason",type="string",description="The reason for the value in 'Ready'"
// +kubebuilder:printcolumn:name="Status Age",JSONPath=".status.conditions[?(@.type=='Ready')].lastTransitionTime",type="date",description="The last transition time for the value in 'Status'"

// ComputeInstance is the Schema for the ComputeInstance API
// +k8s:openapi-gen=true
type ComputeInstance struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	// +required
	Spec   ComputeInstanceSpec   `json:"spec,omitempty"`
	Status ComputeInstanceStatus `json:"status,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// ComputeInstanceList contains a list of ComputeInstance
type ComputeInstanceList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []ComputeInstance `json:"items"`
}

func init() {
	SchemeBuilder.Register(&ComputeInstance{}, &ComputeInstanceList{})
}

type ComputeInstanceAdvancedMachineFeatures struct {
	// Whether to enable nested virtualization or not.
	// +kcc:proto:field=google.cloud.compute.v1.AdvancedMachineFeatures.enable_nested_virtualization
	EnableNestedVirtualization *bool `json:"enableNestedVirtualization,omitempty"`

	// The number of threads per physical core. To disable simultaneous multithreading (SMT) set this to 1. If unset, the maximum number of threads supported per core by the underlying processor is assumed.
	// +kcc:proto:field=google.cloud.compute.v1.AdvancedMachineFeatures.threads_per_core
	ThreadsPerCore *int32 `json:"threadsPerCore,omitempty"`

	// The number of physical cores to expose to an instance. Multiply by the number of threads per core to compute the total number of virtual CPUs to expose to the instance. If unset, the number of cores is inferred from the instance's nominal CPU count and the underlying platform's SMT width.
	// +kcc:proto:field=google.cloud.compute.v1.AdvancedMachineFeatures.visible_core_count
	VisibleCoreCount *int32 `json:"visibleCoreCount,omitempty"`
}

type ComputeInstanceAttachedDisk struct {
	// Name with which the attached disk is accessible under /dev/disk/by-id/.
	DeviceName *string `json:"deviceName,omitempty"`

	// A 256-bit customer-supplied encryption key, encoded in RFC 4648 base64 to encrypt this disk. Only one of kms_key_self_link and disk_encryption_key_raw may be set.
	DiskEncryptionKeyRaw *ComputeInstanceDiskEncryptionKeyRaw `json:"diskEncryptionKeyRaw,omitempty"`

	// The RFC 4648 base64 encoded SHA-256 hash of the customer-supplied encryption key that protects this resource.
	DiskEncryptionKeySha256 *string `json:"diskEncryptionKeySha256,omitempty"`

	// The KMS key to use for the disk.
	KmsKeyRef *ComputeInstanceKmsKeyRef `json:"kmsKeyRef,omitempty"`

	// Read/write mode for the disk. One of "READ_ONLY" or "READ_WRITE".
	Mode *string `json:"mode,omitempty"`

	// The source disk.
	// +required
	SourceDiskRef *ComputeInstanceSourceDiskRef `json:"sourceDiskRef,omitempty"`
}

type ComputeInstanceDiskEncryptionKeyRaw struct {
	// Value of the field. Cannot be used if 'valueFrom' is specified.
	Value *string `json:"value,omitempty"`

	// Source for the field's value. Cannot be used if 'value' is specified.
	ValueFrom *ComputeInstanceValueFrom `json:"valueFrom,omitempty"`
}

type ComputeInstanceValueFrom struct {
	// Reference to a value with the given key in the given Secret in the resource's namespace.
	SecretKeyRef *v1.SecretKeySelector `json:"secretKeyRef,omitempty"`
}

type ComputeInstanceKmsKeyRef struct {
	// Allowed value: The `selfLink` field of a `KMSCryptoKey` resource.
	External *string `json:"external,omitempty"`

	// Name of the referent. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names
	Name *string `json:"name,omitempty"`

	// Namespace of the referent. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/namespaces/
	Namespace *string `json:"namespace,omitempty"`
}

type ComputeInstanceSourceDiskRef struct {
	// Allowed value: The `selfLink` field of a `ComputeDisk` resource.
	External *string `json:"external,omitempty"`

	// Name of the referent. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names
	Name *string `json:"name,omitempty"`

	// Namespace of the referent. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/namespaces/
	Namespace *string `json:"namespace,omitempty"`
}

type ComputeInstanceBootDisk struct {
	// Immutable. Whether the disk will be auto-deleted when the instance is deleted.
	AutoDelete *bool `json:"autoDelete,omitempty"`

	// Immutable. Name with which attached disk will be accessible under /dev/disk/by-id/.
	DeviceName *string `json:"deviceName,omitempty"`

	// Immutable. A 256-bit customer-supplied encryption key, encoded in RFC 4648 base64 to encrypt this disk. Only one of kms_key_self_link and disk_encryption_key_raw may be set.
	DiskEncryptionKeyRaw *ComputeInstanceDiskEncryptionKeyRaw `json:"diskEncryptionKeyRaw,omitempty"`

	// The RFC 4648 base64 encoded SHA-256 hash of the customer-supplied encryption key that protects this resource.
	DiskEncryptionKeySha256 *string `json:"diskEncryptionKeySha256,omitempty"`

	// Immutable. Parameters with which a disk was created alongside the instance.
	InitializeParams *ComputeInstanceInitializeParams `json:"initializeParams,omitempty"`

	// The KMS key to use for the disk.
	KmsKeyRef *ComputeInstanceKmsKeyRef `json:"kmsKeyRef,omitempty"`

	// Immutable. Read/write mode for the disk. One of "READ_ONLY" or "READ_WRITE".
	Mode *string `json:"mode,omitempty"`

	// Immutable. The source disk used to create this disk.
	SourceDiskRef *ComputeInstanceSourceDiskRef `json:"sourceDiskRef,omitempty"`
}

type ComputeInstanceInitializeParams struct {
	// Immutable. A set of key/value label pairs assigned to the disk.
	Labels map[string]string `json:"labels,omitempty"`

	// Immutable. A map of resource manager tags. Resource manager tag keys and values have the same definition as resource manager tags. Keys must be in the format tagKeys/{tag_key_id}, and values are in the format tagValues/456. The field is ignored (both PUT & PATCH) when empty.
	ResourceManagerTags map[string]string `json:"resourceManagerTags,omitempty"`

	// Immutable. The size of the image in gigabytes.
	Size *int64 `json:"size,omitempty"`

	// Immutable. The image from which to initialize this disk.
	SourceImageRef *ComputeInstanceSourceImageRef `json:"sourceImageRef,omitempty"`

	// Immutable. The Google Compute Engine disk type. Such as pd-standard, pd-ssd or pd-balanced.
	Type *string `json:"type,omitempty"`
}

type ComputeInstanceSourceImageRef struct {
	// Allowed value: The `selfLink` field of a `ComputeImage` resource.
	External *string `json:"external,omitempty"`

	// Name of the referent. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names
	Name *string `json:"name,omitempty"`

	// Namespace of the referent. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/namespaces/
	Namespace *string `json:"namespace,omitempty"`
}

type ComputeInstanceInstanceTemplateRef struct {
	// Allowed value: The `selfLink` field of a `ComputeInstanceTemplate` resource.
	External *string `json:"external,omitempty"`

	// Name of the referent. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names
	Name *string `json:"name,omitempty"`

	// Namespace of the referent. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/namespaces/
	Namespace *string `json:"namespace,omitempty"`
}

type ComputeInstanceMetadata struct {
	Key   *string `json:"key"`
	Value *string `json:"value"`
}

type ComputeInstanceNetworkInterface struct {
	// Access configurations, i.e. IPs via which this instance can be accessed via the Internet.
	AccessConfig []ComputeInstanceAccessConfig `json:"accessConfig,omitempty"`

	// An array of alias IP ranges for this network interface.
	AliasIpRange []ComputeInstanceAliasIpRange `json:"aliasIpRange,omitempty"`

	// The prefix length of the primary internal IPv6 range.
	InternalIpv6PrefixLength *int32 `json:"internalIpv6PrefixLength,omitempty"`

	// An array of IPv6 access configurations for this interface. Currently, only one IPv6 access config, DIRECT_IPV6, is supported. If there is no ipv6AccessConfig specified, then this instance will have no external IPv6 Internet access.
	Ipv6AccessConfig []ComputeInstanceIpv6AccessConfig `json:"ipv6AccessConfig,omitempty"`

	// One of EXTERNAL, INTERNAL to indicate whether the IP can be accessed from the Internet. This field is always inherited from its subnetwork.
	Ipv6AccessType *string `json:"ipv6AccessType,omitempty"`

	// An IPv6 internal network address for this network interface. If not specified, Google Cloud will automatically assign an internal IPv6 address from the instance's subnetwork.
	Ipv6Address *string `json:"ipv6Address,omitempty"`

	// The name of the interface.
	Name *string `json:"name,omitempty"`

	// DEPRECATED. Although this field is still available, there is limited support. We recommend that you use `spec.networkInterface.networkIpRef` instead.
	NetworkIp *string `json:"networkIp,omitempty"`

	NetworkIpRef *ComputeInstanceNetworkIpRef `json:"networkIpRef,omitempty"`

	NetworkRef *ComputeInstanceNetworkRef `json:"networkRef,omitempty"`

	// Immutable. The type of vNIC to be used on this interface. Possible values:GVNIC, VIRTIO_NET.
	NicType *string `json:"nicType,omitempty"`

	// Immutable. The networking queue count that's specified by users for the network interface. Both Rx and Tx queues will be set to this number. It will be empty if not specified.
	QueueCount *int32 `json:"queueCount,omitempty"`

	// The stack type for this network interface to identify whether the IPv6 feature is enabled or not. If not specified, IPV4_ONLY will be used.
	StackType *string `json:"stackType,omitempty"`

	// The project in which the subnetwork belongs.
	SubnetworkProject *string `json:"subnetworkProject,omitempty"`

	SubnetworkRef *ComputeInstanceSubnetworkRef `json:"subnetworkRef,omitempty"`
}

type ComputeInstanceAccessConfig struct {
	NatIpRef *ComputeInstanceNatIpRef `json:"natIpRef,omitempty"`

	// The networking tier used for configuring this instance. One of PREMIUM or STANDARD.
	NetworkTier *string `json:"networkTier,omitempty"`

	// The DNS domain name for the public PTR record.
	PublicPtrDomainName *string `json:"publicPtrDomainName,omitempty"`
}

type ComputeInstanceNatIpRef struct {
	// Allowed value: The `address` field of a `ComputeAddress` resource.
	External *string `json:"external,omitempty"`

	// Name of the referent. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names
	Name *string `json:"name,omitempty"`

	// Namespace of the referent. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/namespaces/
	Namespace *string `json:"namespace,omitempty"`
}

type ComputeInstanceAliasIpRange struct {
	// The IP CIDR range represented by this alias IP range.
	IpCidrRange *string `json:"ipCidrRange"`

	// The subnetwork secondary range name specifying the secondary range from which to allocate the IP CIDR range for this alias IP range.
	SubnetworkRangeName *string `json:"subnetworkRangeName,omitempty"`
}

type ComputeInstanceIpv6AccessConfig struct {
	// Immutable. The first IPv6 address of the external IPv6 range associated with this instance, prefix length is stored in externalIpv6PrefixLength in ipv6AccessConfig. To use a static external IP address, it must be unused and in the same region as the instance's zone. If not specified, Google Cloud will automatically assign an external IPv6 address from the instance's subnetwork.
	ExternalIpv6 *string `json:"externalIpv6,omitempty"`

	// Immutable. The prefix length of the external IPv6 range.
	ExternalIpv6PrefixLength *string `json:"externalIpv6PrefixLength,omitempty"`

	// Immutable. The name of this access configuration. In ipv6AccessConfigs, the recommended name is External IPv6.
	Name *string `json:"name,omitempty"`

	// The service-level to be provided for IPv6 traffic when the subnet has an external subnet. Only PREMIUM tier is valid for IPv6.
	NetworkTier *string `json:"networkTier"`

	// The domain name to be used when creating DNSv6 records for the external IPv6 ranges.
	PublicPtrDomainName *string `json:"publicPtrDomainName,omitempty"`
}

type ComputeInstanceNetworkIpRef struct {
	// Allowed value: The `address` field of a `ComputeAddress` resource.
	External *string `json:"external,omitempty"`

	// Kind of the referent. Allowed values: ComputeAddress
	Kind *string `json:"kind,omitempty"`

	// Name of the referent. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names
	Name *string `json:"name,omitempty"`

	// Namespace of the referent. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/namespaces/
	Namespace *string `json:"namespace,omitempty"`
}

type ComputeInstanceNetworkRef struct {
	// Allowed value: The `selfLink` field of a `ComputeNetwork` resource.
	External *string `json:"external,omitempty"`

	// Name of the referent. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names
	Name *string `json:"name,omitempty"`

	// Namespace of the referent. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/namespaces/
	Namespace *string `json:"namespace,omitempty"`
}

type ComputeInstanceSubnetworkRef struct {
	// Allowed value: The `selfLink` field of a `ComputeSubnetwork` resource.
	External *string `json:"external,omitempty"`

	// Name of the referent. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names
	Name *string `json:"name,omitempty"`

	// Namespace of the referent. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/namespaces/
	Namespace *string `json:"namespace,omitempty"`
}

type ComputeInstanceResourcePolicyRef struct {
	// Allowed value: The `selfLink` field of a `ComputeResourcePolicy` resource.
	External *string `json:"external,omitempty"`

	// Name of the referent. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names
	Name *string `json:"name,omitempty"`

	// Namespace of the referent. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/namespaces/
	Namespace *string `json:"namespace,omitempty"`
}

type ComputeInstanceScratchDisk struct {
	// The disk interface used for attaching this disk. One of SCSI or NVME.
	Interface *string `json:"interface"`

	// Immutable. The size of the disk in gigabytes. One of 375 or 3000.
	Size *int32 `json:"size,omitempty"`
}

type ComputeInstanceServiceAccount struct {
	// A list of service scopes.
	Scopes []string `json:"scopes"`

	ServiceAccountRef *ComputeInstanceServiceAccountRef `json:"serviceAccountRef,omitempty"`
}

type ComputeInstanceServiceAccountRef struct {
	// Allowed value: The `email` field of an `IAMServiceAccount` resource.
	External *string `json:"external,omitempty"`

	// Name of the referent. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names
	Name *string `json:"name,omitempty"`

	// Namespace of the referent. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/namespaces/
	Namespace *string `json:"namespace,omitempty"`
}

type ComputeInstanceConfidentialInstanceConfig struct {
	// Defines whether the instance should have confidential compute enabled.
	// +kcc:proto:field=google.cloud.compute.v1.ConfidentialInstanceConfig.enable_confidential_compute
	EnableConfidentialCompute *bool `json:"enableConfidentialCompute,omitempty"`
}

type ComputeInstanceShieldedInstanceConfig struct {
	// Whether integrity monitoring is enabled for the instance.
	// +kcc:proto:field=google.cloud.compute.v1.ShieldedInstanceConfig.enable_integrity_monitoring
	EnableIntegrityMonitoring *bool `json:"enableIntegrityMonitoring,omitempty"`

	// Whether secure boot is enabled for the instance.
	// +kcc:proto:field=google.cloud.compute.v1.ShieldedInstanceConfig.enable_secure_boot
	EnableSecureBoot *bool `json:"enableSecureBoot,omitempty"`

	// Whether the instance uses vTPM.
	// +kcc:proto:field=google.cloud.compute.v1.ShieldedInstanceConfig.enable_vtpm
	EnableVtpm *bool `json:"enableVtpm,omitempty"`
}

type ComputeInstanceGuestAccelerator struct {
	// Immutable. The number of the guest accelerator cards exposed to this instance.
	// +kcc:proto:field=google.cloud.compute.v1.AcceleratorConfig.accelerator_count
	Count *int32 `json:"count,omitempty"`

	// Immutable. The accelerator type resource exposed to this instance. E.g. nvidia-tesla-t4.
	// +kcc:proto:field=google.cloud.compute.v1.AcceleratorConfig.accelerator_type
	Type *string `json:"type,omitempty"`
}
