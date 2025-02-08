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

package domains

import (
	refs "github.com/GoogleCloudPlatform/k8s-config-connector/apis/refs/v1beta1"
	pb "cloud.google.com/go/domains/apiv1beta1/domainspb"
	krm "github.com/GoogleCloudPlatform/k8s-config-connector/apis/domains/v1alpha1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct"
)
func ContactSettings_FromProto(mapCtx *direct.MapContext, in *pb.ContactSettings) *krm.ContactSettings {
	if in == nil {
		return nil
	}
	out := &krm.ContactSettings{}
	out.Privacy = direct.Enum_FromProto(mapCtx, in.GetPrivacy())
	out.RegistrantContact = ContactSettings_Contact_FromProto(mapCtx, in.GetRegistrantContact())
	out.AdminContact = ContactSettings_Contact_FromProto(mapCtx, in.GetAdminContact())
	out.TechnicalContact = ContactSettings_Contact_FromProto(mapCtx, in.GetTechnicalContact())
	return out
}
func ContactSettings_ToProto(mapCtx *direct.MapContext, in *krm.ContactSettings) *pb.ContactSettings {
	if in == nil {
		return nil
	}
	out := &pb.ContactSettings{}
	out.Privacy = direct.Enum_ToProto[pb.ContactPrivacy](mapCtx, in.Privacy)
	out.RegistrantContact = ContactSettings_Contact_ToProto(mapCtx, in.RegistrantContact)
	out.AdminContact = ContactSettings_Contact_ToProto(mapCtx, in.AdminContact)
	out.TechnicalContact = ContactSettings_Contact_ToProto(mapCtx, in.TechnicalContact)
	return out
}
func ContactSettings_Contact_FromProto(mapCtx *direct.MapContext, in *pb.ContactSettings_Contact) *krm.ContactSettings_Contact {
	if in == nil {
		return nil
	}
	out := &krm.ContactSettings_Contact{}
	out.PostalAddress = PostalAddress_FromProto(mapCtx, in.GetPostalAddress())
	out.Email = direct.LazyPtr(in.GetEmail())
	out.PhoneNumber = direct.LazyPtr(in.GetPhoneNumber())
	out.FaxNumber = direct.LazyPtr(in.GetFaxNumber())
	return out
}
func ContactSettings_Contact_ToProto(mapCtx *direct.MapContext, in *krm.ContactSettings_Contact) *pb.ContactSettings_Contact {
	if in == nil {
		return nil
	}
	out := &pb.ContactSettings_Contact{}
	out.PostalAddress = PostalAddress_ToProto(mapCtx, in.PostalAddress)
	out.Email = direct.ValueOf(in.Email)
	out.PhoneNumber = direct.ValueOf(in.PhoneNumber)
	out.FaxNumber = direct.ValueOf(in.FaxNumber)
	return out
}
func DnsSettings_FromProto(mapCtx *direct.MapContext, in *pb.DnsSettings) *krm.DnsSettings {
	if in == nil {
		return nil
	}
	out := &krm.DnsSettings{}
	out.CustomDns = DnsSettings_CustomDns_FromProto(mapCtx, in.GetCustomDns())
	out.GoogleDomainsDns = DnsSettings_GoogleDomainsDns_FromProto(mapCtx, in.GetGoogleDomainsDns())
	out.GlueRecords = direct.Slice_FromProto(mapCtx, in.GlueRecords, DnsSettings_GlueRecord_FromProto)
	return out
}
func DnsSettings_ToProto(mapCtx *direct.MapContext, in *krm.DnsSettings) *pb.DnsSettings {
	if in == nil {
		return nil
	}
	out := &pb.DnsSettings{}
	if oneof := DnsSettings_CustomDns_ToProto(mapCtx, in.CustomDns); oneof != nil {
		out.DnsProvider = &pb.DnsSettings_CustomDns_{CustomDns: oneof}
	}
	if oneof := DnsSettings_GoogleDomainsDns_ToProto(mapCtx, in.GoogleDomainsDns); oneof != nil {
		out.DnsProvider = &pb.DnsSettings_GoogleDomainsDns_{GoogleDomainsDns: oneof}
	}
	out.GlueRecords = direct.Slice_ToProto(mapCtx, in.GlueRecords, DnsSettings_GlueRecord_ToProto)
	return out
}
func DnsSettingsObservedState_FromProto(mapCtx *direct.MapContext, in *pb.DnsSettings) *krm.DnsSettingsObservedState {
	if in == nil {
		return nil
	}
	out := &krm.DnsSettingsObservedState{}
	// MISSING: CustomDns
	out.GoogleDomainsDns = DnsSettings_GoogleDomainsDnsObservedState_FromProto(mapCtx, in.GetGoogleDomainsDns())
	// MISSING: GlueRecords
	return out
}
func DnsSettingsObservedState_ToProto(mapCtx *direct.MapContext, in *krm.DnsSettingsObservedState) *pb.DnsSettings {
	if in == nil {
		return nil
	}
	out := &pb.DnsSettings{}
	// MISSING: CustomDns
	if oneof := DnsSettings_GoogleDomainsDnsObservedState_ToProto(mapCtx, in.GoogleDomainsDns); oneof != nil {
		out.DnsProvider = &pb.DnsSettings_GoogleDomainsDns_{GoogleDomainsDns: oneof}
	}
	// MISSING: GlueRecords
	return out
}
func DnsSettings_CustomDns_FromProto(mapCtx *direct.MapContext, in *pb.DnsSettings_CustomDns) *krm.DnsSettings_CustomDns {
	if in == nil {
		return nil
	}
	out := &krm.DnsSettings_CustomDns{}
	out.NameServers = in.NameServers
	out.DsRecords = direct.Slice_FromProto(mapCtx, in.DsRecords, DnsSettings_DsRecord_FromProto)
	return out
}
func DnsSettings_CustomDns_ToProto(mapCtx *direct.MapContext, in *krm.DnsSettings_CustomDns) *pb.DnsSettings_CustomDns {
	if in == nil {
		return nil
	}
	out := &pb.DnsSettings_CustomDns{}
	out.NameServers = in.NameServers
	out.DsRecords = direct.Slice_ToProto(mapCtx, in.DsRecords, DnsSettings_DsRecord_ToProto)
	return out
}
func DnsSettings_DsRecord_FromProto(mapCtx *direct.MapContext, in *pb.DnsSettings_DsRecord) *krm.DnsSettings_DsRecord {
	if in == nil {
		return nil
	}
	out := &krm.DnsSettings_DsRecord{}
	out.KeyTag = direct.LazyPtr(in.GetKeyTag())
	out.Algorithm = direct.Enum_FromProto(mapCtx, in.GetAlgorithm())
	out.DigestType = direct.Enum_FromProto(mapCtx, in.GetDigestType())
	out.Digest = direct.LazyPtr(in.GetDigest())
	return out
}
func DnsSettings_DsRecord_ToProto(mapCtx *direct.MapContext, in *krm.DnsSettings_DsRecord) *pb.DnsSettings_DsRecord {
	if in == nil {
		return nil
	}
	out := &pb.DnsSettings_DsRecord{}
	out.KeyTag = direct.ValueOf(in.KeyTag)
	out.Algorithm = direct.Enum_ToProto[pb.DnsSettings_DsRecord_Algorithm](mapCtx, in.Algorithm)
	out.DigestType = direct.Enum_ToProto[pb.DnsSettings_DsRecord_DigestType](mapCtx, in.DigestType)
	out.Digest = direct.ValueOf(in.Digest)
	return out
}
func DnsSettings_GlueRecord_FromProto(mapCtx *direct.MapContext, in *pb.DnsSettings_GlueRecord) *krm.DnsSettings_GlueRecord {
	if in == nil {
		return nil
	}
	out := &krm.DnsSettings_GlueRecord{}
	out.HostName = direct.LazyPtr(in.GetHostName())
	out.Ipv4Addresses = in.Ipv4Addresses
	out.Ipv6Addresses = in.Ipv6Addresses
	return out
}
func DnsSettings_GlueRecord_ToProto(mapCtx *direct.MapContext, in *krm.DnsSettings_GlueRecord) *pb.DnsSettings_GlueRecord {
	if in == nil {
		return nil
	}
	out := &pb.DnsSettings_GlueRecord{}
	out.HostName = direct.ValueOf(in.HostName)
	out.Ipv4Addresses = in.Ipv4Addresses
	out.Ipv6Addresses = in.Ipv6Addresses
	return out
}
func DnsSettings_GoogleDomainsDns_FromProto(mapCtx *direct.MapContext, in *pb.DnsSettings_GoogleDomainsDns) *krm.DnsSettings_GoogleDomainsDns {
	if in == nil {
		return nil
	}
	out := &krm.DnsSettings_GoogleDomainsDns{}
	// MISSING: NameServers
	out.DsState = direct.Enum_FromProto(mapCtx, in.GetDsState())
	// MISSING: DsRecords
	return out
}
func DnsSettings_GoogleDomainsDns_ToProto(mapCtx *direct.MapContext, in *krm.DnsSettings_GoogleDomainsDns) *pb.DnsSettings_GoogleDomainsDns {
	if in == nil {
		return nil
	}
	out := &pb.DnsSettings_GoogleDomainsDns{}
	// MISSING: NameServers
	out.DsState = direct.Enum_ToProto[pb.DnsSettings_DsState](mapCtx, in.DsState)
	// MISSING: DsRecords
	return out
}
func DnsSettings_GoogleDomainsDnsObservedState_FromProto(mapCtx *direct.MapContext, in *pb.DnsSettings_GoogleDomainsDns) *krm.DnsSettings_GoogleDomainsDnsObservedState {
	if in == nil {
		return nil
	}
	out := &krm.DnsSettings_GoogleDomainsDnsObservedState{}
	out.NameServers = in.NameServers
	// MISSING: DsState
	out.DsRecords = direct.Slice_FromProto(mapCtx, in.DsRecords, DnsSettings_DsRecord_FromProto)
	return out
}
func DnsSettings_GoogleDomainsDnsObservedState_ToProto(mapCtx *direct.MapContext, in *krm.DnsSettings_GoogleDomainsDnsObservedState) *pb.DnsSettings_GoogleDomainsDns {
	if in == nil {
		return nil
	}
	out := &pb.DnsSettings_GoogleDomainsDns{}
	out.NameServers = in.NameServers
	// MISSING: DsState
	out.DsRecords = direct.Slice_ToProto(mapCtx, in.DsRecords, DnsSettings_DsRecord_ToProto)
	return out
}
func ManagementSettings_FromProto(mapCtx *direct.MapContext, in *pb.ManagementSettings) *krm.ManagementSettings {
	if in == nil {
		return nil
	}
	out := &krm.ManagementSettings{}
	// MISSING: RenewalMethod
	out.TransferLockState = direct.Enum_FromProto(mapCtx, in.GetTransferLockState())
	return out
}
func ManagementSettings_ToProto(mapCtx *direct.MapContext, in *krm.ManagementSettings) *pb.ManagementSettings {
	if in == nil {
		return nil
	}
	out := &pb.ManagementSettings{}
	// MISSING: RenewalMethod
	out.TransferLockState = direct.Enum_ToProto[pb.TransferLockState](mapCtx, in.TransferLockState)
	return out
}
func ManagementSettingsObservedState_FromProto(mapCtx *direct.MapContext, in *pb.ManagementSettings) *krm.ManagementSettingsObservedState {
	if in == nil {
		return nil
	}
	out := &krm.ManagementSettingsObservedState{}
	out.RenewalMethod = direct.Enum_FromProto(mapCtx, in.GetRenewalMethod())
	// MISSING: TransferLockState
	return out
}
func ManagementSettingsObservedState_ToProto(mapCtx *direct.MapContext, in *krm.ManagementSettingsObservedState) *pb.ManagementSettings {
	if in == nil {
		return nil
	}
	out := &pb.ManagementSettings{}
	out.RenewalMethod = direct.Enum_ToProto[pb.ManagementSettings_RenewalMethod](mapCtx, in.RenewalMethod)
	// MISSING: TransferLockState
	return out
}
func Registration_FromProto(mapCtx *direct.MapContext, in *pb.Registration) *krm.Registration {
	if in == nil {
		return nil
	}
	out := &krm.Registration{}
	// MISSING: Name
	out.DomainName = direct.LazyPtr(in.GetDomainName())
	// MISSING: CreateTime
	// MISSING: ExpireTime
	// MISSING: State
	// MISSING: Issues
	out.Labels = in.Labels
	out.ManagementSettings = ManagementSettings_FromProto(mapCtx, in.GetManagementSettings())
	out.DnsSettings = DnsSettings_FromProto(mapCtx, in.GetDnsSettings())
	out.ContactSettings = ContactSettings_FromProto(mapCtx, in.GetContactSettings())
	// MISSING: PendingContactSettings
	// MISSING: SupportedPrivacy
	return out
}
func Registration_ToProto(mapCtx *direct.MapContext, in *krm.Registration) *pb.Registration {
	if in == nil {
		return nil
	}
	out := &pb.Registration{}
	// MISSING: Name
	out.DomainName = direct.ValueOf(in.DomainName)
	// MISSING: CreateTime
	// MISSING: ExpireTime
	// MISSING: State
	// MISSING: Issues
	out.Labels = in.Labels
	out.ManagementSettings = ManagementSettings_ToProto(mapCtx, in.ManagementSettings)
	out.DnsSettings = DnsSettings_ToProto(mapCtx, in.DnsSettings)
	out.ContactSettings = ContactSettings_ToProto(mapCtx, in.ContactSettings)
	// MISSING: PendingContactSettings
	// MISSING: SupportedPrivacy
	return out
}
func RegistrationObservedState_FromProto(mapCtx *direct.MapContext, in *pb.Registration) *krm.RegistrationObservedState {
	if in == nil {
		return nil
	}
	out := &krm.RegistrationObservedState{}
	out.Name = direct.LazyPtr(in.GetName())
	// MISSING: DomainName
	out.CreateTime = direct.StringTimestamp_FromProto(mapCtx, in.GetCreateTime())
	out.ExpireTime = direct.StringTimestamp_FromProto(mapCtx, in.GetExpireTime())
	out.State = direct.Enum_FromProto(mapCtx, in.GetState())
	out.Issues = direct.EnumSlice_FromProto(mapCtx, in.Issues)
	// MISSING: Labels
	out.ManagementSettings = ManagementSettingsObservedState_FromProto(mapCtx, in.GetManagementSettings())
	out.DnsSettings = DnsSettingsObservedState_FromProto(mapCtx, in.GetDnsSettings())
	// MISSING: ContactSettings
	out.PendingContactSettings = ContactSettings_FromProto(mapCtx, in.GetPendingContactSettings())
	out.SupportedPrivacy = direct.EnumSlice_FromProto(mapCtx, in.SupportedPrivacy)
	return out
}
func RegistrationObservedState_ToProto(mapCtx *direct.MapContext, in *krm.RegistrationObservedState) *pb.Registration {
	if in == nil {
		return nil
	}
	out := &pb.Registration{}
	out.Name = direct.ValueOf(in.Name)
	// MISSING: DomainName
	out.CreateTime = direct.StringTimestamp_ToProto(mapCtx, in.CreateTime)
	out.ExpireTime = direct.StringTimestamp_ToProto(mapCtx, in.ExpireTime)
	out.State = direct.Enum_ToProto[pb.Registration_State](mapCtx, in.State)
	out.Issues = direct.EnumSlice_ToProto[pb.Registration_Issue](mapCtx, in.Issues)
	// MISSING: Labels
	out.ManagementSettings = ManagementSettingsObservedState_ToProto(mapCtx, in.ManagementSettings)
	out.DnsSettings = DnsSettingsObservedState_ToProto(mapCtx, in.DnsSettings)
	// MISSING: ContactSettings
	out.PendingContactSettings = ContactSettings_ToProto(mapCtx, in.PendingContactSettings)
	out.SupportedPrivacy = direct.EnumSlice_ToProto[pb.ContactPrivacy](mapCtx, in.SupportedPrivacy)
	return out
}
