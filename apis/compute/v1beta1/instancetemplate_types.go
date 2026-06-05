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

package v1beta1

import (
	refsv1beta1 "github.com/GoogleCloudPlatform/k8s-config-connector/apis/refs/v1beta1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/apis/k8s/v1alpha1"
	apiextensionsv1 "k8s.io/apiextensions-apiserver/pkg/apis/apiextensions/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

var ComputeInstanceTemplateGVK = GroupVersion.WithKind("ComputeInstanceTemplate")

type InstanceTemplateAdvancedMachineFeatures struct {
	EnableNestedVirtualization *bool  `json:"enableNestedVirtualization,omitempty"`
	ThreadsPerCore             *int64 `json:"threadsPerCore,omitempty"`
	VisibleCoreCount           *int64 `json:"visibleCoreCount,omitempty"`
}

type InstanceTemplateConfidentialInstanceConfig struct {
	EnableConfidentialCompute bool `json:"enableConfidentialCompute"`
}

type ComputeDiskRef struct {
	External  string `json:"external,omitempty"`
	Name      string `json:"name,omitempty"`
	Namespace string `json:"namespace,omitempty"`
}

type ComputeSnapshotRef struct {
	External  string `json:"external,omitempty"`
	Name      string `json:"name,omitempty"`
	Namespace string `json:"namespace,omitempty"`
}

type InstanceTemplateDiskEncryptionKey struct {
	KmsKeyRef *refsv1beta1.KMSCryptoKeyRef `json:"kmsKeyRef"`
}

type InstanceTemplateSourceImageEncryptionKey struct {
	KmsKeySelfLinkRef       *refsv1beta1.KMSCryptoKeyRef      `json:"kmsKeySelfLinkRef"`
	KmsKeyServiceAccountRef *refsv1beta1.IAMServiceAccountRef `json:"kmsKeyServiceAccountRef,omitempty"`
}

type InstanceTemplateSourceSnapshotEncryptionKey struct {
	KmsKeySelfLinkRef       *refsv1beta1.KMSCryptoKeyRef      `json:"kmsKeySelfLinkRef"`
	KmsKeyServiceAccountRef *refsv1beta1.IAMServiceAccountRef `json:"kmsKeyServiceAccountRef,omitempty"`
}

type InstanceTemplateDisk struct {
	AutoDelete                  *bool                                        `json:"autoDelete,omitempty"`
	Boot                        *bool                                        `json:"boot,omitempty"`
	DeviceName                  *string                                      `json:"deviceName,omitempty"`
	DiskEncryptionKey           *InstanceTemplateDiskEncryptionKey           `json:"diskEncryptionKey,omitempty"`
	DiskName                    *string                                      `json:"diskName,omitempty"`
	DiskSizeGb                  *int64                                       `json:"diskSizeGb,omitempty"`
	DiskType                    *string                                      `json:"diskType,omitempty"`
	Interface                   *string                                      `json:"interface,omitempty"`
	Labels                      map[string]string                            `json:"labels,omitempty"`
	Mode                        *string                                      `json:"mode,omitempty"`
	ProvisionedIops             *int64                                       `json:"provisionedIops,omitempty"`
	ResourcePolicies            []*ComputeResourcePolicyRef                  `json:"resourcePolicies,omitempty"`
	SourceDiskRef               *ComputeDiskRef                              `json:"sourceDiskRef,omitempty"`
	SourceImageEncryptionKey    *InstanceTemplateSourceImageEncryptionKey    `json:"sourceImageEncryptionKey,omitempty"`
	SourceImageRef              *ComputeImageRef                             `json:"sourceImageRef,omitempty"`
	SourceSnapshotEncryptionKey *InstanceTemplateSourceSnapshotEncryptionKey `json:"sourceSnapshotEncryptionKey,omitempty"`
	SourceSnapshotRef           *ComputeSnapshotRef                          `json:"sourceSnapshotRef,omitempty"`
	Type                        *string                                      `json:"type,omitempty"`
}

type InstanceTemplateGuestAccelerator struct {
	Count int64  `json:"count"`
	Type  string `json:"type"`
}

type InstanceTemplateMetadata struct {
	Key   string `json:"key"`
	Value string `json:"value"`
}

type InstanceTemplateAccessConfig struct {
	NatIpRef            *refsv1beta1.ComputeAddressRef `json:"natIpRef,omitempty"`
	NetworkTier         *string                        `json:"networkTier,omitempty"`
	PublicPtrDomainName *string                        `json:"publicPtrDomainName,omitempty"`
}

type InstanceTemplateAliasIpRange struct {
	IpCidrRange         string  `json:"ipCidrRange"`
	SubnetworkRangeName *string `json:"subnetworkRangeName,omitempty"`
}

type InstanceTemplateIpv6AccessConfig struct {
	ExternalIpv6             *string `json:"externalIpv6,omitempty"`
	ExternalIpv6PrefixLength *string `json:"externalIpv6PrefixLength,omitempty"`
	Name                     *string `json:"name,omitempty"`
	NetworkTier              string  `json:"networkTier"`
	PublicPtrDomainName      *string `json:"publicPtrDomainName,omitempty"`
}

type InstanceTemplateNetworkInterface struct {
	AccessConfig             []InstanceTemplateAccessConfig     `json:"accessConfig,omitempty"`
	AliasIpRange             []InstanceTemplateAliasIpRange     `json:"aliasIpRange,omitempty"`
	InternalIpv6PrefixLength *int64                             `json:"internalIpv6PrefixLength,omitempty"`
	Ipv6AccessConfig         []InstanceTemplateIpv6AccessConfig `json:"ipv6AccessConfig,omitempty"`
	Ipv6AccessType           *string                            `json:"ipv6AccessType,omitempty"`
	Ipv6Address              *string                            `json:"ipv6Address,omitempty"`
	Name                     *string                            `json:"name,omitempty"`
	NetworkAttachment        *string                            `json:"networkAttachment,omitempty"`
	NetworkIp                *string                            `json:"networkIp,omitempty"`
	NetworkRef               *ComputeNetworkRef                 `json:"networkRef,omitempty"`
	NicType                  *string                            `json:"nicType,omitempty"`
	QueueCount               *int64                             `json:"queueCount,omitempty"`
	StackType                *string                            `json:"stackType,omitempty"`
	SubnetworkProject        *string                            `json:"subnetworkProject,omitempty"`
	SubnetworkRef            *refsv1beta1.ComputeSubnetworkRef  `json:"subnetworkRef,omitempty"`
}

type InstanceTemplateNetworkPerformanceConfig struct {
	TotalEgressBandwidthTier string `json:"totalEgressBandwidthTier"`
}

type InstanceTemplateSpecificReservation struct {
	Key    string   `json:"key"`
	Values []string `json:"values"`
}

type InstanceTemplateReservationAffinity struct {
	SpecificReservation *InstanceTemplateSpecificReservation `json:"specificReservation,omitempty"`
	Type                string                               `json:"type"`
}

type InstanceTemplateLocalSsdRecoveryTimeout struct {
	Nanos   *int64 `json:"nanos,omitempty"`
	Seconds int64  `json:"seconds"`
}

type InstanceTemplateMaxRunDuration struct {
	Nanos   *int64 `json:"nanos,omitempty"`
	Seconds int64  `json:"seconds"`
}

type InstanceTemplateNodeAffinities struct {
	// +kubebuilder:validation:XPreserveUnknownFields
	Value *apiextensionsv1.JSON `json:"value,omitempty"`
}

type InstanceTemplateScheduling struct {
	AutomaticRestart          *bool                                     `json:"automaticRestart,omitempty"`
	InstanceTerminationAction *string                                   `json:"instanceTerminationAction,omitempty"`
	LocalSsdRecoveryTimeout   []InstanceTemplateLocalSsdRecoveryTimeout `json:"localSsdRecoveryTimeout,omitempty"`
	MaintenanceInterval       *string                                   `json:"maintenanceInterval,omitempty"`
	MaxRunDuration            *InstanceTemplateMaxRunDuration           `json:"maxRunDuration,omitempty"`
	MinNodeCpus               *int64                                    `json:"minNodeCpus,omitempty"`
	NodeAffinities            []InstanceTemplateNodeAffinities          `json:"nodeAffinities,omitempty"`
	OnHostMaintenance         *string                                   `json:"onHostMaintenance,omitempty"`
	Preemptible               *bool                                     `json:"preemptible,omitempty"`
	ProvisioningModel         *string                                   `json:"provisioningModel,omitempty"`
}

type InstanceTemplateServiceAccount struct {
	Scopes            []string                          `json:"scopes"`
	ServiceAccountRef *refsv1beta1.IAMServiceAccountRef `json:"serviceAccountRef,omitempty"`
}

type InstanceTemplateShieldedInstanceConfig struct {
	EnableIntegrityMonitoring *bool `json:"enableIntegrityMonitoring,omitempty"`
	EnableSecureBoot          *bool `json:"enableSecureBoot,omitempty"`
	EnableVtpm                *bool `json:"enableVtpm,omitempty"`
}

// ComputeInstanceTemplateSpec defines the desired state of ComputeInstanceTemplate
type ComputeInstanceTemplateSpec struct {
	AdvancedMachineFeatures    *InstanceTemplateAdvancedMachineFeatures    `json:"advancedMachineFeatures,omitempty"`
	CanIpForward               *bool                                       `json:"canIpForward,omitempty"`
	ConfidentialInstanceConfig *InstanceTemplateConfidentialInstanceConfig `json:"confidentialInstanceConfig,omitempty"`
	Description                *string                                     `json:"description,omitempty"`
	Disk                       []InstanceTemplateDisk                      `json:"disk"`
	EnableDisplay              *bool                                       `json:"enableDisplay,omitempty"`
	GuestAccelerator           []InstanceTemplateGuestAccelerator          `json:"guestAccelerator,omitempty"`
	InstanceDescription        *string                                     `json:"instanceDescription,omitempty"`
	MachineType                string                                      `json:"machineType"`
	Metadata                   []InstanceTemplateMetadata                  `json:"metadata,omitempty"`
	MetadataStartupScript      *string                                     `json:"metadataStartupScript,omitempty"`
	MinCpuPlatform             *string                                     `json:"minCpuPlatform,omitempty"`
	NamePrefix                 *string                                     `json:"namePrefix,omitempty"`
	NetworkInterface           []InstanceTemplateNetworkInterface          `json:"networkInterface,omitempty"`
	NetworkPerformanceConfig   *InstanceTemplateNetworkPerformanceConfig   `json:"networkPerformanceConfig,omitempty"`
	Region                     *string                                     `json:"region,omitempty"`
	ReservationAffinity        *InstanceTemplateReservationAffinity        `json:"reservationAffinity,omitempty"`
	ResourceID                 *string                                     `json:"resourceID,omitempty"`
	ResourcePolicies           []*ComputeResourcePolicyRef                 `json:"resourcePolicies,omitempty"`
	Scheduling                 *InstanceTemplateScheduling                 `json:"scheduling,omitempty"`
	ServiceAccount             *InstanceTemplateServiceAccount             `json:"serviceAccount,omitempty"`
	ShieldedInstanceConfig     *InstanceTemplateShieldedInstanceConfig     `json:"shieldedInstanceConfig,omitempty"`
	Tags                       []string                                    `json:"tags,omitempty"`
}

// ComputeInstanceTemplateStatus defines the config connector machine state of ComputeInstanceTemplate
type ComputeInstanceTemplateStatus struct {
	/* Conditions represent the latest available observations of the
	   object's current state. */
	Conditions []v1alpha1.Condition `json:"conditions,omitempty"`

	// ObservedGeneration is the generation of the resource that was most recently observed by the Config Connector controller. If this is equal to metadata.generation, then that means that the current reported status reflects the most recent desired state of the resource.
	ObservedGeneration *int64 `json:"observedGeneration,omitempty"`

	MetadataFingerprint *string `json:"metadataFingerprint,omitempty"`

	SelfLink *string `json:"selfLink,omitempty"`

	SelfLinkUnique *string `json:"selfLinkUnique,omitempty"`

	TagsFingerprint *string `json:"tagsFingerprint,omitempty"`
}

// +genclient
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:resource:categories=gcp,shortName=gcpcomputeinstancetemplate;gcpcomputeinstancetemplates
// +kubebuilder:subresource:status
// +kubebuilder:metadata:labels="cnrm.cloud.google.com/managed-by-kcc=true"
// +kubebuilder:metadata:labels="cnrm.cloud.google.com/system=true"
// +kubebuilder:metadata:labels="cnrm.cloud.google.com/stability-level=stable"
// +kubebuilder:metadata:labels="cnrm.cloud.google.com/tf2crd=true"
// +kubebuilder:printcolumn:name="Age",JSONPath=".metadata.creationTimestamp",type="date"
// +kubebuilder:printcolumn:name="Ready",JSONPath=".status.conditions[?(@.type=='Ready')].status",type="string",description="When 'True', the most recent reconcile of the resource succeeded"
// +kubebuilder:printcolumn:name="Status",JSONPath=".status.conditions[?(@.type=='Ready')].reason",type="string",description="The reason for the value in 'Ready'"
// +kubebuilder:printcolumn:name="Status Age",JSONPath=".status.conditions[?(@.type=='Ready')].lastTransitionTime",type="date",description="The last transition time for the value in 'Status'"

// ComputeInstanceTemplate is the Schema for the ComputeInstanceTemplate API
// +k8s:openapi-gen=true
type ComputeInstanceTemplate struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	// +required
	Spec   ComputeInstanceTemplateSpec   `json:"spec,omitempty"`
	Status ComputeInstanceTemplateStatus `json:"status,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// ComputeInstanceTemplateList contains a list of ComputeInstanceTemplate
type ComputeInstanceTemplateList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []ComputeInstanceTemplate `json:"items"`
}

func init() {
	SchemeBuilder.Register(&ComputeInstanceTemplate{}, &ComputeInstanceTemplateList{})
}
