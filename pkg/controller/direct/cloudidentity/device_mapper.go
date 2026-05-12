// Copyright 2026 Google LLC
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

package cloudidentity

import (
	krm "github.com/GoogleCloudPlatform/k8s-config-connector/apis/cloudidentity/v1beta1"
	api "google.golang.org/api/cloudidentity/v1beta1"
)

func CloudIdentityDeviceSpec_ToAPI(in *krm.CloudIdentityDeviceSpec) *api.Device {
	if in == nil {
		return nil
	}
	out := &api.Device{}
	if in.AssetTag != nil {
		out.AssetTag = *in.AssetTag
	}
	// Note: BrowserProfiles, ClientTypes, Hostname, LastSyncTime, SerialNumber, WifiMacAddresses
	// Are omitted from Create requests in the API, but we could map them.
	// We'll map them for completeness.
	if in.DeviceType != nil {
		out.DeviceType = *in.DeviceType
	}
	if in.SerialNumber != nil {
		out.SerialNumber = *in.SerialNumber
	}
	return out
}

func CloudIdentityDeviceObservedState_FromAPI(in *api.Device) *krm.CloudIdentityDeviceObservedState {
	if in == nil {
		return nil
	}
	out := &krm.CloudIdentityDeviceObservedState{}

	if in.BasebandVersion != "" {
		out.BasebandVersion = &in.BasebandVersion
	}
	if in.BootloaderVersion != "" {
		out.BootloaderVersion = &in.BootloaderVersion
	}
	if in.Brand != "" {
		out.Brand = &in.Brand
	}
	if in.BuildNumber != "" {
		out.BuildNumber = &in.BuildNumber
	}
	if in.CompromisedState != "" {
		out.CompromisedState = &in.CompromisedState
	}
	if in.CreateTime != "" {
		out.CreateTime = &in.CreateTime
	}
	if in.EncryptionState != "" {
		out.EncryptionState = &in.EncryptionState
	}
	if in.Imei != "" {
		out.Imei = &in.Imei
	}
	if in.KernelVersion != "" {
		out.KernelVersion = &in.KernelVersion
	}
	if in.ManagementState != "" {
		out.ManagementState = &in.ManagementState
	}
	if in.Manufacturer != "" {
		out.Manufacturer = &in.Manufacturer
	}
	if in.Meid != "" {
		out.Meid = &in.Meid
	}
	if in.Model != "" {
		out.Model = &in.Model
	}
	if in.Name != "" {
		out.Name = &in.Name
	}
	if in.NetworkOperator != "" {
		out.NetworkOperator = &in.NetworkOperator
	}
	if in.OsVersion != "" {
		out.OSVersion = &in.OsVersion
	}
	if in.OwnerType != "" {
		out.OwnerType = &in.OwnerType
	}
	if in.ReleaseVersion != "" {
		out.ReleaseVersion = &in.ReleaseVersion
	}
	if in.SecurityPatchTime != "" {
		out.SecurityPatchTime = &in.SecurityPatchTime
	}
	if in.UnifiedDeviceId != "" {
		out.UnifiedDeviceID = &in.UnifiedDeviceId
	}

	out.EnabledDeveloperOptions = &in.EnabledDeveloperOptions
	out.EnabledUsbDebugging = &in.EnabledUsbDebugging

	// For simple lists:
	if len(in.OtherAccounts) > 0 {
		out.OtherAccounts = make([]string, len(in.OtherAccounts))
		copy(out.OtherAccounts, in.OtherAccounts)
	}

	return out
}

func CloudIdentityDeviceSpec_FromAPI(in *api.Device) *krm.CloudIdentityDeviceSpec {
	if in == nil {
		return nil
	}
	out := &krm.CloudIdentityDeviceSpec{}

	if in.AssetTag != "" {
		out.AssetTag = &in.AssetTag
	}
	if in.DeviceType != "" {
		out.DeviceType = &in.DeviceType
	}
	if in.SerialNumber != "" {
		out.SerialNumber = &in.SerialNumber
	}
	if in.DeviceId != "" {
		out.DeviceID = &in.DeviceId
	}
	if in.Hostname != "" {
		out.Hostname = &in.Hostname
	}
	if in.LastSyncTime != "" {
		out.LastSyncTime = &in.LastSyncTime
	}
	if len(in.WifiMacAddresses) > 0 {
		out.WifiMacAddresses = make([]string, len(in.WifiMacAddresses))
		copy(out.WifiMacAddresses, in.WifiMacAddresses)
	}
	if len(in.ClientTypes) > 0 {
		out.ClientTypes = make([]string, len(in.ClientTypes))
		copy(out.ClientTypes, in.ClientTypes)
	}
	return out
}
