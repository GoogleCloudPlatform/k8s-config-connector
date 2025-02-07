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

package channel

import (
	refs "github.com/GoogleCloudPlatform/k8s-config-connector/apis/refs/v1beta1"
	pb "cloud.google.com/go/channel/apiv1/channelpb"
	krm "github.com/GoogleCloudPlatform/k8s-config-connector/apis/channel/v1alpha1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct"
)
func ChannelBillingAccountObservedState_FromProto(mapCtx *direct.MapContext, in *pb.BillingAccount) *krm.ChannelBillingAccountObservedState {
	if in == nil {
		return nil
	}
	out := &krm.ChannelBillingAccountObservedState{}
	// MISSING: Name
	// MISSING: DisplayName
	// MISSING: CreateTime
	// MISSING: CurrencyCode
	// MISSING: RegionCode
	return out
}
func ChannelBillingAccountObservedState_ToProto(mapCtx *direct.MapContext, in *krm.ChannelBillingAccountObservedState) *pb.BillingAccount {
	if in == nil {
		return nil
	}
	out := &pb.BillingAccount{}
	// MISSING: Name
	// MISSING: DisplayName
	// MISSING: CreateTime
	// MISSING: CurrencyCode
	// MISSING: RegionCode
	return out
}
func ChannelBillingAccountSpec_FromProto(mapCtx *direct.MapContext, in *pb.BillingAccount) *krm.ChannelBillingAccountSpec {
	if in == nil {
		return nil
	}
	out := &krm.ChannelBillingAccountSpec{}
	// MISSING: Name
	// MISSING: DisplayName
	// MISSING: CreateTime
	// MISSING: CurrencyCode
	// MISSING: RegionCode
	return out
}
func ChannelBillingAccountSpec_ToProto(mapCtx *direct.MapContext, in *krm.ChannelBillingAccountSpec) *pb.BillingAccount {
	if in == nil {
		return nil
	}
	out := &pb.BillingAccount{}
	// MISSING: Name
	// MISSING: DisplayName
	// MISSING: CreateTime
	// MISSING: CurrencyCode
	// MISSING: RegionCode
	return out
}
func ChannelChannelPartnerLinkObservedState_FromProto(mapCtx *direct.MapContext, in *pb.ChannelPartnerLink) *krm.ChannelChannelPartnerLinkObservedState {
	if in == nil {
		return nil
	}
	out := &krm.ChannelChannelPartnerLinkObservedState{}
	// MISSING: Name
	// MISSING: ResellerCloudIdentityID
	// MISSING: LinkState
	// MISSING: InviteLinkURI
	// MISSING: CreateTime
	// MISSING: UpdateTime
	// MISSING: PublicID
	// MISSING: ChannelPartnerCloudIdentityInfo
	return out
}
func ChannelChannelPartnerLinkObservedState_ToProto(mapCtx *direct.MapContext, in *krm.ChannelChannelPartnerLinkObservedState) *pb.ChannelPartnerLink {
	if in == nil {
		return nil
	}
	out := &pb.ChannelPartnerLink{}
	// MISSING: Name
	// MISSING: ResellerCloudIdentityID
	// MISSING: LinkState
	// MISSING: InviteLinkURI
	// MISSING: CreateTime
	// MISSING: UpdateTime
	// MISSING: PublicID
	// MISSING: ChannelPartnerCloudIdentityInfo
	return out
}
func ChannelChannelPartnerLinkSpec_FromProto(mapCtx *direct.MapContext, in *pb.ChannelPartnerLink) *krm.ChannelChannelPartnerLinkSpec {
	if in == nil {
		return nil
	}
	out := &krm.ChannelChannelPartnerLinkSpec{}
	// MISSING: Name
	// MISSING: ResellerCloudIdentityID
	// MISSING: LinkState
	// MISSING: InviteLinkURI
	// MISSING: CreateTime
	// MISSING: UpdateTime
	// MISSING: PublicID
	// MISSING: ChannelPartnerCloudIdentityInfo
	return out
}
func ChannelChannelPartnerLinkSpec_ToProto(mapCtx *direct.MapContext, in *krm.ChannelChannelPartnerLinkSpec) *pb.ChannelPartnerLink {
	if in == nil {
		return nil
	}
	out := &pb.ChannelPartnerLink{}
	// MISSING: Name
	// MISSING: ResellerCloudIdentityID
	// MISSING: LinkState
	// MISSING: InviteLinkURI
	// MISSING: CreateTime
	// MISSING: UpdateTime
	// MISSING: PublicID
	// MISSING: ChannelPartnerCloudIdentityInfo
	return out
}
func ChannelCustomerObservedState_FromProto(mapCtx *direct.MapContext, in *pb.Customer) *krm.ChannelCustomerObservedState {
	if in == nil {
		return nil
	}
	out := &krm.ChannelCustomerObservedState{}
	// MISSING: Name
	// MISSING: OrgDisplayName
	// MISSING: OrgPostalAddress
	// MISSING: PrimaryContactInfo
	// MISSING: AlternateEmail
	// MISSING: Domain
	// MISSING: CreateTime
	// MISSING: UpdateTime
	// MISSING: CloudIdentityID
	// MISSING: LanguageCode
	// MISSING: CloudIdentityInfo
	// MISSING: ChannelPartnerID
	// MISSING: CorrelationID
	return out
}
func ChannelCustomerObservedState_ToProto(mapCtx *direct.MapContext, in *krm.ChannelCustomerObservedState) *pb.Customer {
	if in == nil {
		return nil
	}
	out := &pb.Customer{}
	// MISSING: Name
	// MISSING: OrgDisplayName
	// MISSING: OrgPostalAddress
	// MISSING: PrimaryContactInfo
	// MISSING: AlternateEmail
	// MISSING: Domain
	// MISSING: CreateTime
	// MISSING: UpdateTime
	// MISSING: CloudIdentityID
	// MISSING: LanguageCode
	// MISSING: CloudIdentityInfo
	// MISSING: ChannelPartnerID
	// MISSING: CorrelationID
	return out
}
func ChannelCustomerSpec_FromProto(mapCtx *direct.MapContext, in *pb.Customer) *krm.ChannelCustomerSpec {
	if in == nil {
		return nil
	}
	out := &krm.ChannelCustomerSpec{}
	// MISSING: Name
	// MISSING: OrgDisplayName
	// MISSING: OrgPostalAddress
	// MISSING: PrimaryContactInfo
	// MISSING: AlternateEmail
	// MISSING: Domain
	// MISSING: CreateTime
	// MISSING: UpdateTime
	// MISSING: CloudIdentityID
	// MISSING: LanguageCode
	// MISSING: CloudIdentityInfo
	// MISSING: ChannelPartnerID
	// MISSING: CorrelationID
	return out
}
func ChannelCustomerSpec_ToProto(mapCtx *direct.MapContext, in *krm.ChannelCustomerSpec) *pb.Customer {
	if in == nil {
		return nil
	}
	out := &pb.Customer{}
	// MISSING: Name
	// MISSING: OrgDisplayName
	// MISSING: OrgPostalAddress
	// MISSING: PrimaryContactInfo
	// MISSING: AlternateEmail
	// MISSING: Domain
	// MISSING: CreateTime
	// MISSING: UpdateTime
	// MISSING: CloudIdentityID
	// MISSING: LanguageCode
	// MISSING: CloudIdentityInfo
	// MISSING: ChannelPartnerID
	// MISSING: CorrelationID
	return out
}
func CloudIdentityInfo_FromProto(mapCtx *direct.MapContext, in *pb.CloudIdentityInfo) *krm.CloudIdentityInfo {
	if in == nil {
		return nil
	}
	out := &krm.CloudIdentityInfo{}
	out.CustomerType = direct.Enum_FromProto(mapCtx, in.GetCustomerType())
	// MISSING: PrimaryDomain
	// MISSING: IsDomainVerified
	out.AlternateEmail = direct.LazyPtr(in.GetAlternateEmail())
	out.PhoneNumber = direct.LazyPtr(in.GetPhoneNumber())
	out.LanguageCode = direct.LazyPtr(in.GetLanguageCode())
	// MISSING: AdminConsoleURI
	out.EduData = EduData_FromProto(mapCtx, in.GetEduData())
	return out
}
func CloudIdentityInfo_ToProto(mapCtx *direct.MapContext, in *krm.CloudIdentityInfo) *pb.CloudIdentityInfo {
	if in == nil {
		return nil
	}
	out := &pb.CloudIdentityInfo{}
	out.CustomerType = direct.Enum_ToProto[pb.CloudIdentityInfo_CustomerType](mapCtx, in.CustomerType)
	// MISSING: PrimaryDomain
	// MISSING: IsDomainVerified
	out.AlternateEmail = direct.ValueOf(in.AlternateEmail)
	out.PhoneNumber = direct.ValueOf(in.PhoneNumber)
	out.LanguageCode = direct.ValueOf(in.LanguageCode)
	// MISSING: AdminConsoleURI
	out.EduData = EduData_ToProto(mapCtx, in.EduData)
	return out
}
func CloudIdentityInfoObservedState_FromProto(mapCtx *direct.MapContext, in *pb.CloudIdentityInfo) *krm.CloudIdentityInfoObservedState {
	if in == nil {
		return nil
	}
	out := &krm.CloudIdentityInfoObservedState{}
	// MISSING: CustomerType
	out.PrimaryDomain = direct.LazyPtr(in.GetPrimaryDomain())
	out.IsDomainVerified = direct.LazyPtr(in.GetIsDomainVerified())
	// MISSING: AlternateEmail
	// MISSING: PhoneNumber
	// MISSING: LanguageCode
	out.AdminConsoleURI = direct.LazyPtr(in.GetAdminConsoleUri())
	// MISSING: EduData
	return out
}
func CloudIdentityInfoObservedState_ToProto(mapCtx *direct.MapContext, in *krm.CloudIdentityInfoObservedState) *pb.CloudIdentityInfo {
	if in == nil {
		return nil
	}
	out := &pb.CloudIdentityInfo{}
	// MISSING: CustomerType
	out.PrimaryDomain = direct.ValueOf(in.PrimaryDomain)
	out.IsDomainVerified = direct.ValueOf(in.IsDomainVerified)
	// MISSING: AlternateEmail
	// MISSING: PhoneNumber
	// MISSING: LanguageCode
	out.AdminConsoleUri = direct.ValueOf(in.AdminConsoleURI)
	// MISSING: EduData
	return out
}
func ContactInfo_FromProto(mapCtx *direct.MapContext, in *pb.ContactInfo) *krm.ContactInfo {
	if in == nil {
		return nil
	}
	out := &krm.ContactInfo{}
	out.FirstName = direct.LazyPtr(in.GetFirstName())
	out.LastName = direct.LazyPtr(in.GetLastName())
	// MISSING: DisplayName
	out.Email = direct.LazyPtr(in.GetEmail())
	out.Title = direct.LazyPtr(in.GetTitle())
	out.Phone = direct.LazyPtr(in.GetPhone())
	return out
}
func ContactInfo_ToProto(mapCtx *direct.MapContext, in *krm.ContactInfo) *pb.ContactInfo {
	if in == nil {
		return nil
	}
	out := &pb.ContactInfo{}
	out.FirstName = direct.ValueOf(in.FirstName)
	out.LastName = direct.ValueOf(in.LastName)
	// MISSING: DisplayName
	out.Email = direct.ValueOf(in.Email)
	out.Title = direct.ValueOf(in.Title)
	out.Phone = direct.ValueOf(in.Phone)
	return out
}
func ContactInfoObservedState_FromProto(mapCtx *direct.MapContext, in *pb.ContactInfo) *krm.ContactInfoObservedState {
	if in == nil {
		return nil
	}
	out := &krm.ContactInfoObservedState{}
	// MISSING: FirstName
	// MISSING: LastName
	out.DisplayName = direct.LazyPtr(in.GetDisplayName())
	// MISSING: Email
	// MISSING: Title
	// MISSING: Phone
	return out
}
func ContactInfoObservedState_ToProto(mapCtx *direct.MapContext, in *krm.ContactInfoObservedState) *pb.ContactInfo {
	if in == nil {
		return nil
	}
	out := &pb.ContactInfo{}
	// MISSING: FirstName
	// MISSING: LastName
	out.DisplayName = direct.ValueOf(in.DisplayName)
	// MISSING: Email
	// MISSING: Title
	// MISSING: Phone
	return out
}
func Customer_FromProto(mapCtx *direct.MapContext, in *pb.Customer) *krm.Customer {
	if in == nil {
		return nil
	}
	out := &krm.Customer{}
	// MISSING: Name
	out.OrgDisplayName = direct.LazyPtr(in.GetOrgDisplayName())
	out.OrgPostalAddress = PostalAddress_FromProto(mapCtx, in.GetOrgPostalAddress())
	out.PrimaryContactInfo = ContactInfo_FromProto(mapCtx, in.GetPrimaryContactInfo())
	out.AlternateEmail = direct.LazyPtr(in.GetAlternateEmail())
	out.Domain = direct.LazyPtr(in.GetDomain())
	// MISSING: CreateTime
	// MISSING: UpdateTime
	// MISSING: CloudIdentityID
	out.LanguageCode = direct.LazyPtr(in.GetLanguageCode())
	// MISSING: CloudIdentityInfo
	out.ChannelPartnerID = direct.LazyPtr(in.GetChannelPartnerId())
	out.CorrelationID = direct.LazyPtr(in.GetCorrelationId())
	return out
}
func Customer_ToProto(mapCtx *direct.MapContext, in *krm.Customer) *pb.Customer {
	if in == nil {
		return nil
	}
	out := &pb.Customer{}
	// MISSING: Name
	out.OrgDisplayName = direct.ValueOf(in.OrgDisplayName)
	out.OrgPostalAddress = PostalAddress_ToProto(mapCtx, in.OrgPostalAddress)
	out.PrimaryContactInfo = ContactInfo_ToProto(mapCtx, in.PrimaryContactInfo)
	out.AlternateEmail = direct.ValueOf(in.AlternateEmail)
	out.Domain = direct.ValueOf(in.Domain)
	// MISSING: CreateTime
	// MISSING: UpdateTime
	// MISSING: CloudIdentityID
	out.LanguageCode = direct.ValueOf(in.LanguageCode)
	// MISSING: CloudIdentityInfo
	out.ChannelPartnerId = direct.ValueOf(in.ChannelPartnerID)
	out.CorrelationId = direct.ValueOf(in.CorrelationID)
	return out
}
func CustomerObservedState_FromProto(mapCtx *direct.MapContext, in *pb.Customer) *krm.CustomerObservedState {
	if in == nil {
		return nil
	}
	out := &krm.CustomerObservedState{}
	out.Name = direct.LazyPtr(in.GetName())
	// MISSING: OrgDisplayName
	// MISSING: OrgPostalAddress
	out.PrimaryContactInfo = ContactInfoObservedState_FromProto(mapCtx, in.GetPrimaryContactInfo())
	// MISSING: AlternateEmail
	// MISSING: Domain
	out.CreateTime = direct.StringTimestamp_FromProto(mapCtx, in.GetCreateTime())
	out.UpdateTime = direct.StringTimestamp_FromProto(mapCtx, in.GetUpdateTime())
	out.CloudIdentityID = direct.LazyPtr(in.GetCloudIdentityId())
	// MISSING: LanguageCode
	out.CloudIdentityInfo = CloudIdentityInfo_FromProto(mapCtx, in.GetCloudIdentityInfo())
	// MISSING: ChannelPartnerID
	// MISSING: CorrelationID
	return out
}
func CustomerObservedState_ToProto(mapCtx *direct.MapContext, in *krm.CustomerObservedState) *pb.Customer {
	if in == nil {
		return nil
	}
	out := &pb.Customer{}
	out.Name = direct.ValueOf(in.Name)
	// MISSING: OrgDisplayName
	// MISSING: OrgPostalAddress
	out.PrimaryContactInfo = ContactInfoObservedState_ToProto(mapCtx, in.PrimaryContactInfo)
	// MISSING: AlternateEmail
	// MISSING: Domain
	out.CreateTime = direct.StringTimestamp_ToProto(mapCtx, in.CreateTime)
	out.UpdateTime = direct.StringTimestamp_ToProto(mapCtx, in.UpdateTime)
	out.CloudIdentityId = direct.ValueOf(in.CloudIdentityID)
	// MISSING: LanguageCode
	out.CloudIdentityInfo = CloudIdentityInfo_ToProto(mapCtx, in.CloudIdentityInfo)
	// MISSING: ChannelPartnerID
	// MISSING: CorrelationID
	return out
}
func EduData_FromProto(mapCtx *direct.MapContext, in *pb.EduData) *krm.EduData {
	if in == nil {
		return nil
	}
	out := &krm.EduData{}
	out.InstituteType = direct.Enum_FromProto(mapCtx, in.GetInstituteType())
	out.InstituteSize = direct.Enum_FromProto(mapCtx, in.GetInstituteSize())
	out.Website = direct.LazyPtr(in.GetWebsite())
	return out
}
func EduData_ToProto(mapCtx *direct.MapContext, in *krm.EduData) *pb.EduData {
	if in == nil {
		return nil
	}
	out := &pb.EduData{}
	out.InstituteType = direct.Enum_ToProto[pb.EduData_InstituteType](mapCtx, in.InstituteType)
	out.InstituteSize = direct.Enum_ToProto[pb.EduData_InstituteSize](mapCtx, in.InstituteSize)
	out.Website = direct.ValueOf(in.Website)
	return out
}
