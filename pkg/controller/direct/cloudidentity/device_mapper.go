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
	krm "github.com/GoogleCloudPlatform/k8s-config-connector/apis/cloudidentity/v1alpha1"
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
	if in.DeviceType != nil {
		out.DeviceType = *in.DeviceType
	}
	if in.SerialNumber != nil {
		out.SerialNumber = *in.SerialNumber
	}
	if in.DeviceID != nil {
		out.DeviceId = *in.DeviceID
	}
	if in.Hostname != nil {
		out.Hostname = *in.Hostname
	}
	if in.LastSyncTime != nil {
		out.LastSyncTime = *in.LastSyncTime
	}
	if len(in.WifiMacAddresses) > 0 {
		out.WifiMacAddresses = make([]string, len(in.WifiMacAddresses))
		copy(out.WifiMacAddresses, in.WifiMacAddresses)
	}
	if len(in.ClientTypes) > 0 {
		out.ClientTypes = make([]string, len(in.ClientTypes))
		copy(out.ClientTypes, in.ClientTypes)
	}

	if len(in.BrowserProfiles) > 0 {
		out.BrowserProfiles = make([]*api.BrowserAttributes, len(in.BrowserProfiles))
		for i, bp := range in.BrowserProfiles {
			out.BrowserProfiles[i] = &api.BrowserAttributes{
				ChromeProfileId:     derefString(bp.ChromeProfileID),
				LastProfileSyncTime: derefString(bp.LastProfileSyncTime),
			}
			if bp.ChromeBrowserInfo != nil {
				cbi := bp.ChromeBrowserInfo
				out.BrowserProfiles[i].ChromeBrowserInfo = &api.BrowserInfo{
					BrowserVersion:                   derefString(cbi.BrowserVersion),
					IsBuiltInDnsClientEnabled:        derefBool(cbi.IsBuiltInDNSClientEnabled),
					IsBulkDataEntryAnalysisEnabled:   derefBool(cbi.IsBulkDataEntryAnalysisEnabled),
					IsChromeCleanupEnabled:           derefBool(cbi.IsChromeCleanupEnabled),
					IsChromeRemoteDesktopAppBlocked:  derefBool(cbi.IsChromeRemoteDesktopAppBlocked),
					IsFileDownloadAnalysisEnabled:    derefBool(cbi.IsFileDownloadAnalysisEnabled),
					IsFileUploadAnalysisEnabled:      derefBool(cbi.IsFileUploadAnalysisEnabled),
					IsRealtimeUrlCheckEnabled:        derefBool(cbi.IsRealtimeURLCheckEnabled),
					IsSecurityEventAnalysisEnabled:   derefBool(cbi.IsSecurityEventAnalysisEnabled),
					IsSiteIsolationEnabled:           derefBool(cbi.IsSiteIsolationEnabled),
					IsThirdPartyBlockingEnabled:      derefBool(cbi.IsThirdPartyBlockingEnabled),
					PasswordProtectionWarningTrigger: derefString(cbi.PasswordProtectionWarningTrigger),
					SafeBrowsingProtectionLevel:      derefString(cbi.SafeBrowsingProtectionLevel),
				}
			}
		}
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

	if len(in.OtherAccounts) > 0 {
		out.OtherAccounts = make([]string, len(in.OtherAccounts))
		copy(out.OtherAccounts, in.OtherAccounts)
	}

	if len(in.BrowserProfiles) > 0 {
		out.BrowserProfiles = make([]krm.BrowserAttributesObservedState, len(in.BrowserProfiles))
		for i, bp := range in.BrowserProfiles {
			out.BrowserProfiles[i] = krm.BrowserAttributesObservedState{}
			if bp.ChromeBrowserInfo != nil {
				cbi := bp.ChromeBrowserInfo
				out.BrowserProfiles[i].ChromeBrowserInfo = &krm.BrowserInfoObservedState{
					BrowserManagementState: stringPtr(cbi.BrowserManagementState),
				}
				if len(cbi.Policies) > 0 {
					out.BrowserProfiles[i].ChromeBrowserInfo.Policies = make([]krm.ChromePolicy, len(cbi.Policies))
					for j, p := range cbi.Policies {
						out.BrowserProfiles[i].ChromeBrowserInfo.Policies[j] = krm.ChromePolicy{
							Name:   stringPtr(p.Name),
							Scope:  stringPtr(p.Scope),
							Source: stringPtr(p.Source),
							Value:  stringPtr(p.Value),
						}
						if len(p.Conflicts) > 0 {
							out.BrowserProfiles[i].ChromeBrowserInfo.Policies[j].Conflicts = make([]krm.PolicyConflict, len(p.Conflicts))
							for k, c := range p.Conflicts {
								out.BrowserProfiles[i].ChromeBrowserInfo.Policies[j].Conflicts[k] = krm.PolicyConflict{
									Scope:  stringPtr(c.Scope),
									Source: stringPtr(c.Source),
									Value:  stringPtr(c.Value),
								}
							}
						}
					}
				}
			}
		}
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

	if len(in.BrowserProfiles) > 0 {
		out.BrowserProfiles = make([]krm.BrowserAttributes, len(in.BrowserProfiles))
		for i, bp := range in.BrowserProfiles {
			out.BrowserProfiles[i] = krm.BrowserAttributes{
				ChromeProfileID:     stringPtr(bp.ChromeProfileId),
				LastProfileSyncTime: stringPtr(bp.LastProfileSyncTime),
			}
			if bp.ChromeBrowserInfo != nil {
				cbi := bp.ChromeBrowserInfo
				out.BrowserProfiles[i].ChromeBrowserInfo = &krm.BrowserInfo{
					BrowserVersion:                   stringPtr(cbi.BrowserVersion),
					IsBuiltInDNSClientEnabled:        boolPtr(cbi.IsBuiltInDnsClientEnabled),
					IsBulkDataEntryAnalysisEnabled:   boolPtr(cbi.IsBulkDataEntryAnalysisEnabled),
					IsChromeCleanupEnabled:           boolPtr(cbi.IsChromeCleanupEnabled),
					IsChromeRemoteDesktopAppBlocked:  boolPtr(cbi.IsChromeRemoteDesktopAppBlocked),
					IsFileDownloadAnalysisEnabled:    boolPtr(cbi.IsFileDownloadAnalysisEnabled),
					IsFileUploadAnalysisEnabled:      boolPtr(cbi.IsFileUploadAnalysisEnabled),
					IsRealtimeURLCheckEnabled:        boolPtr(cbi.IsRealtimeUrlCheckEnabled),
					IsSecurityEventAnalysisEnabled:   boolPtr(cbi.IsSecurityEventAnalysisEnabled),
					IsSiteIsolationEnabled:           boolPtr(cbi.IsSiteIsolationEnabled),
					IsThirdPartyBlockingEnabled:      boolPtr(cbi.IsThirdPartyBlockingEnabled),
					PasswordProtectionWarningTrigger: stringPtr(cbi.PasswordProtectionWarningTrigger),
					SafeBrowsingProtectionLevel:      stringPtr(cbi.SafeBrowsingProtectionLevel),
				}
			}
		}
	}

	return out
}

func derefString(s *string) string {
	if s == nil {
		return ""
	}
	return *s
}

func derefBool(b *bool) bool {
	if b == nil {
		return false
	}
	return *b
}

func stringPtr(s string) *string {
	if s == "" {
		return nil
	}
	return &s
}

func boolPtr(b bool) *bool {
	return &b
}
