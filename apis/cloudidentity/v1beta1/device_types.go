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
// +kcc:spec:proto=google.apps.cloudidentity.v1beta1.Device
type CloudIdentityDeviceSpec struct {
	// Type of device.
	// +kcc:proto:field=google.apps.cloudidentity.v1beta1.Device.device_type
	DeviceType *string `json:"deviceType,omitempty"`

	// Asset tag of the device.
	// +kcc:proto:field=google.apps.cloudidentity.v1beta1.Device.asset_tag
	AssetTag *string `json:"assetTag,omitempty"`

	// Browser profiles on the device. This is a copy of the BrowserAttributes message defined in EndpointVerificationSpecificAttributes. We are replicating it here since EndpointVerification isn't the only client reporting browser profiles.
	// +kcc:proto:field=google.apps.cloudidentity.v1beta1.Device.browser_profiles
	BrowserProfiles []BrowserAttributes `json:"browserProfiles,omitempty"`

	// List of the clients the device is reporting to.
	// +kcc:proto:field=google.apps.cloudidentity.v1beta1.Device.client_types
	ClientTypes []string `json:"clientTypes,omitempty"`

	// Unique identifier for the device.
	// +kcc:proto:field=google.apps.cloudidentity.v1beta1.Device.device_id
	DeviceID *string `json:"deviceID,omitempty"`

	// Host name of the device.
	// +kcc:proto:field=google.apps.cloudidentity.v1beta1.Device.hostname
	Hostname *string `json:"hostname,omitempty"`

	// Most recent time when device synced with this service.
	// +kcc:proto:field=google.apps.cloudidentity.v1beta1.Device.last_sync_time
	LastSyncTime *string `json:"lastSyncTime,omitempty"`

	// Serial Number of device. Example: HT82V1A01076.
	// +kcc:proto:field=google.apps.cloudidentity.v1beta1.Device.serial_number
	SerialNumber *string `json:"serialNumber,omitempty"`

	// WiFi MAC addresses of device.
	// +kcc:proto:field=google.apps.cloudidentity.v1beta1.Device.wifi_mac_addresses
	WifiMacAddresses []string `json:"wifiMacAddresses,omitempty"`

	// The CloudIdentityDevice name. If not given, the metadata.name will be used.
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
// +kcc:observedstate:proto=google.apps.cloudidentity.v1beta1.Device
type CloudIdentityDeviceObservedState struct {

	// Output only. Attributes specific to Android devices.
	// +kcc:proto:field=google.apps.cloudidentity.v1beta1.Device.android_specific_attributes
	AndroidSpecificAttributes *AndroidAttributes `json:"androidSpecificAttributes,omitempty"`

	// Output only. Baseband version of the device.
	// +kcc:proto:field=google.apps.cloudidentity.v1beta1.Device.baseband_version
	BasebandVersion *string `json:"basebandVersion,omitempty"`

	// Output only. Device bootloader version. Example: 0.6.7.
	// +kcc:proto:field=google.apps.cloudidentity.v1beta1.Device.bootloader_version
	BootloaderVersion *string `json:"bootloaderVersion,omitempty"`

	// Output only. Device brand. Example: Samsung.
	// +kcc:proto:field=google.apps.cloudidentity.v1beta1.Device.brand
	Brand *string `json:"brand,omitempty"`

	// Output only. Build number of the device.
	// +kcc:proto:field=google.apps.cloudidentity.v1beta1.Device.build_number
	BuildNumber *string `json:"buildNumber,omitempty"`

	// Output only. Represents whether the Device is compromised.
	// +kcc:proto:field=google.apps.cloudidentity.v1beta1.Device.compromised_state
	CompromisedState *string `json:"compromisedState,omitempty"`

	// Output only. When the Company-Owned device was imported. This field is empty for BYOD devices.
	// +kcc:proto:field=google.apps.cloudidentity.v1beta1.Device.create_time
	CreateTime *string `json:"createTime,omitempty"`

	// Output only. Whether developer options is enabled on device.
	// +kcc:proto:field=google.apps.cloudidentity.v1beta1.Device.enabled_developer_options
	EnabledDeveloperOptions *bool `json:"enabledDeveloperOptions,omitempty"`

	// Output only. Whether USB debugging is enabled on device.
	// +kcc:proto:field=google.apps.cloudidentity.v1beta1.Device.enabled_usb_debugging
	EnabledUsbDebugging *bool `json:"enabledUsbDebugging,omitempty"`

	// Output only. Device encryption state.
	// +kcc:proto:field=google.apps.cloudidentity.v1beta1.Device.encryption_state
	EncryptionState *string `json:"encryptionState,omitempty"`

	// Output only. Attributes specific to [Endpoint Verification](https://cloud.google.com/endpoint-verification/docs/overview) devices.
	// +kcc:proto:field=google.apps.cloudidentity.v1beta1.Device.endpoint_verification_specific_attributes
	EndpointVerificationSpecificAttributes *EndpointVerificationSpecificAttributes `json:"endpointVerificationSpecificAttributes,omitempty"`

	// Output only. IMEI number of device if GSM device; empty otherwise.
	// +kcc:proto:field=google.apps.cloudidentity.v1beta1.Device.imei
	Imei *string `json:"imei,omitempty"`

	// Output only. Kernel version of the device.
	// +kcc:proto:field=google.apps.cloudidentity.v1beta1.Device.kernel_version
	KernelVersion *string `json:"kernelVersion,omitempty"`

	// Output only. Management state of the device
	// +kcc:proto:field=google.apps.cloudidentity.v1beta1.Device.management_state
	ManagementState *string `json:"managementState,omitempty"`

	// Output only. Device manufacturer. Example: Motorola.
	// +kcc:proto:field=google.apps.cloudidentity.v1beta1.Device.manufacturer
	Manufacturer *string `json:"manufacturer,omitempty"`

	// Output only. MEID number of device if CDMA device; empty otherwise.
	// +kcc:proto:field=google.apps.cloudidentity.v1beta1.Device.meid
	Meid *string `json:"meid,omitempty"`

	// Output only. Model name of device. Example: Pixel 3.
	// +kcc:proto:field=google.apps.cloudidentity.v1beta1.Device.model
	Model *string `json:"model,omitempty"`

	// Output only. [Resource name](https://cloud.google.com/apis/design/resource_names) of the Device in format: `devices/{device_id}`, where device_id is the unique id assigned to the Device.
	// +kcc:proto:field=google.apps.cloudidentity.v1beta1.Device.name
	Name *string `json:"name,omitempty"`

	// Output only. Mobile or network operator of device, if available.
	// +kcc:proto:field=google.apps.cloudidentity.v1beta1.Device.network_operator
	NetworkOperator *string `json:"networkOperator,omitempty"`

	// Output only. OS version of the device. Example: Android 8.1.0.
	// +kcc:proto:field=google.apps.cloudidentity.v1beta1.Device.os_version
	OSVersion *string `json:"osVersion,omitempty"`

	// Output only. Domain name for Google accounts on device. Type for other accounts on device. On Android, will only be populated if |ownership_privilege| is |PROFILE_OWNER| or |DEVICE_OWNER|. Does not include the account signed in to the device policy app if that account's domain has only one account. Examples: "com.example", "xyz.com".
	// +kcc:proto:field=google.apps.cloudidentity.v1beta1.Device.other_accounts
	OtherAccounts []string `json:"otherAccounts,omitempty"`

	// Output only. Whether the device is owned by the company or an individual
	// +kcc:proto:field=google.apps.cloudidentity.v1beta1.Device.owner_type
	OwnerType *string `json:"ownerType,omitempty"`

	// Output only. OS release version. Example: 6.0.
	// +kcc:proto:field=google.apps.cloudidentity.v1beta1.Device.release_version
	ReleaseVersion *string `json:"releaseVersion,omitempty"`

	// Output only. OS security patch update time on device.
	// +kcc:proto:field=google.apps.cloudidentity.v1beta1.Device.security_patch_time
	SecurityPatchTime *string `json:"securityPatchTime,omitempty"`

	// Output only. Unified device id of the device.
	// +kcc:proto:field=google.apps.cloudidentity.v1beta1.Device.unified_device_id
	UnifiedDeviceID *string `json:"unifiedDeviceID,omitempty"`
}

// +genclient
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:resource:categories=gcp,shortName=gcpcloudidentitydevice;gcpcloudidentitydevices
// +kubebuilder:subresource:status
// +kubebuilder:metadata:labels="cnrm.cloud.google.com/managed-by-kcc=true"
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
