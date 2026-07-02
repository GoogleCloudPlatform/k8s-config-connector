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
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/apis/k8s/v1alpha1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

var CloudIdentityDeviceGVK = GroupVersion.WithKind("CloudIdentityDevice")

// CloudIdentityDeviceSpec defines the desired state of CloudIdentityDevice
// +kcc:spec:proto=google.apps.cloudidentity.v1beta1.GoogleAppsCloudidentityDevicesV1Device
type CloudIdentityDeviceSpec struct {
	// Immutable. Optional. The customer that this device belongs to.
	// Format: customers/{customer_id}
	Customer *string `json:"customer,omitempty"`

	// Immutable. Required. The device type.
	// See the [API reference](https://cloud.google.com/identity/docs/reference/rest/v1beta1/devices/create) for possible values.
	// +required
	DeviceType *string `json:"deviceType"`

	// Immutable. Required. The serial number of the device.
	// +required
	SerialNumber *string `json:"serialNumber"`

	// Immutable. Optional. Asset tag of the device.
	AssetTag *string `json:"assetTag,omitempty"`

	// Immutable. Optional. The service-generated name of the resource. Format: devices/{device} or {device}. Used for acquisition only. Leave unset to create a new resource.
	ResourceID *string `json:"resourceID,omitempty"`
}

// CloudIdentityDeviceStatus defines the config connector machine state of CloudIdentityDevice
type CloudIdentityDeviceStatus struct {
	/* Conditions represent the latest available observations of the
	   object's current state. */
	Conditions []v1alpha1.Condition `json:"conditions,omitempty"`

	// ObservedGeneration is the generation of the resource that was most recently observed by the Config Connector controller. If this is equal to metadata.generation, then that means that the current reported status reflects the most recent desired state of the resource.
	ObservedGeneration *int64 `json:"observedGeneration,omitempty"`

	// A unique specifier for the CloudIdentityDevice resource in GCP.
	ExternalRef *string `json:"externalRef,omitempty"`

	// ObservedState is the state of the resource as most recently observed in GCP.
	ObservedState *CloudIdentityDeviceObservedState `json:"observedState,omitempty"`
}

// CloudIdentityDeviceObservedState is the state of the CloudIdentityDevice resource as most recently observed in GCP.
// +kcc:observedstate:proto=google.apps.cloudidentity.v1beta1.GoogleAppsCloudidentityDevicesV1Device
type CloudIdentityDeviceObservedState struct {
	// Output only. Attributes specific to Android devices.
	AndroidSpecificAttributes *GoogleAppsCloudidentityDevicesV1AndroidAttributes `json:"androidSpecificAttributes,omitempty"`

	// Output only. Baseband version of the device.
	BasebandVersion *string `json:"basebandVersion,omitempty"`

	// Output only. Device bootloader version. Example: 0.6.7.
	BootloaderVersion *string `json:"bootloaderVersion,omitempty"`

	// Output only. Device brand. Example: Samsung.
	Brand *string `json:"brand,omitempty"`

	// Output only. Build number of the device.
	BuildNumber *string `json:"buildNumber,omitempty"`

	// Output only. Represents whether the Device is compromised.
	CompromisedState *string `json:"compromisedState,omitempty"`

	// Output only. When the Company-Owned device was imported. This field is empty for BYOD devices.
	CreateTime *string `json:"createTime,omitempty"`

	// Output only. Unique identifier for the device.
	DeviceID *string `json:"deviceID,omitempty"`

	// Output only. Type of device.
	DeviceType *string `json:"deviceType,omitempty"`

	// Output only. Whether developer options is enabled on device.
	EnabledDeveloperOptions *bool `json:"enabledDeveloperOptions,omitempty"`

	// Output only. Whether USB debugging is enabled on device.
	EnabledUsbDebugging *bool `json:"enabledUsbDebugging,omitempty"`

	// Output only. Device encryption state.
	EncryptionState *string `json:"encryptionState,omitempty"`

	// Output only. Attributes specific to Endpoint Verification devices.
	EndpointVerificationSpecificAttributes *GoogleAppsCloudidentityDevicesV1EndpointVerificationSpecificAttributes `json:"endpointVerificationSpecificAttributes,omitempty"`

	// Output only. Host name of the device.
	Hostname *string `json:"hostname,omitempty"`

	// Output only. IMEI number of device if GSM device; empty otherwise.
	Imei *string `json:"imei,omitempty"`

	// Output only. Kernel version of the device.
	KernelVersion *string `json:"kernelVersion,omitempty"`

	// Output only. Most recent time when device synced with this service.
	LastSyncTime *string `json:"lastSyncTime,omitempty"`

	// Output only. Management state of the device
	ManagementState *string `json:"managementState,omitempty"`

	// Output only. Device manufacturer. Example: Motorola.
	Manufacturer *string `json:"manufacturer,omitempty"`

	// Output only. MEID number of device if CDMA device; empty otherwise.
	Meid *string `json:"meid,omitempty"`

	// Output only. Model name of device. Example: Pixel 3.
	Model *string `json:"model,omitempty"`

	// Output only. Mobile or network operator of device, if available.
	NetworkOperator *string `json:"networkOperator,omitempty"`

	// Output only. OS version of the device. Example: Android 8.1.0.
	OSVersion *string `json:"osVersion,omitempty"`

	// Output only. Domain name for Google accounts on device.
	OtherAccounts []string `json:"otherAccounts,omitempty"`

	// Output only. Whether the device is owned by the company or an individual
	OwnerType *string `json:"ownerType,omitempty"`

	// Output only. OS release version. Example: 6.0.
	ReleaseVersion *string `json:"releaseVersion,omitempty"`

	// Output only. OS security patch update time on device.
	SecurityPatchTime *string `json:"securityPatchTime,omitempty"`

	// Output only. Serial Number of device. Example: HT82V1A01076.
	SerialNumber *string `json:"serialNumber,omitempty"`

	// Output only. Unified device id of the device.
	UnifiedDeviceID *string `json:"unifiedDeviceID,omitempty"`

	// Output only. WiFi MAC addresses of device.
	WifiMacAddresses []string `json:"wifiMacAddresses,omitempty"`
}

// +genclient
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:resource:categories=gcp,shortName=gcpcloudidentitydevice;gcpcloudidentitydevices
// +kubebuilder:subresource:status
// +kubebuilder:metadata:labels="cnrm.cloud.google.com/managed-by-kcc=true"
// +kubebuilder:metadata:labels="cnrm.cloud.google.com/stability-level=alpha"
// +kubebuilder:metadata:labels="cnrm.cloud.google.com/system=true"
// +kubebuilder:printcolumn:name="Age",JSONPath=".metadata.creationTimestamp",type="date"
// +kubebuilder:printcolumn:name="Ready",JSONPath=".status.conditions[?(@.type=='Ready')].status",type="string",description="When 'True', the most recent reconcile of the resource succeeded"
// +kubebuilder:printcolumn:name="Status",JSONPath=".status.conditions[?(@.type=='Ready')].reason",type="string",description="The reason for the value in 'Ready'"
// +kubebuilder:printcolumn:name="Status Age",JSONPath=".status.conditions[?(@.type=='Ready')].lastTransitionTime",type="date",description="The last transition time for the value in 'Status'"

// CloudIdentityDevice is the Schema for the CloudIdentityDevice API
// +k8s:openapi-gen=true
type CloudIdentityDevice struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	// +required
	Spec   CloudIdentityDeviceSpec   `json:"spec,omitempty"`
	Status CloudIdentityDeviceStatus `json:"status,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// CloudIdentityDeviceList contains a list of CloudIdentityDevice
type CloudIdentityDeviceList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []CloudIdentityDevice `json:"items"`
}

func init() {
	SchemeBuilder.Register(&CloudIdentityDevice{}, &CloudIdentityDeviceList{})
}
