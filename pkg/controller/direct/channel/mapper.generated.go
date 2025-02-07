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
func ChannelPartnerLink_FromProto(mapCtx *direct.MapContext, in *pb.ChannelPartnerLink) *krm.ChannelPartnerLink {
	if in == nil {
		return nil
	}
	out := &krm.ChannelPartnerLink{}
	// MISSING: Name
	out.ResellerCloudIdentityID = direct.LazyPtr(in.GetResellerCloudIdentityId())
	out.LinkState = direct.Enum_FromProto(mapCtx, in.GetLinkState())
	// MISSING: InviteLinkURI
	// MISSING: CreateTime
	// MISSING: UpdateTime
	// MISSING: PublicID
	// MISSING: ChannelPartnerCloudIdentityInfo
	return out
}
func ChannelPartnerLink_ToProto(mapCtx *direct.MapContext, in *krm.ChannelPartnerLink) *pb.ChannelPartnerLink {
	if in == nil {
		return nil
	}
	out := &pb.ChannelPartnerLink{}
	// MISSING: Name
	out.ResellerCloudIdentityId = direct.ValueOf(in.ResellerCloudIdentityID)
	out.LinkState = direct.Enum_ToProto[pb.ChannelPartnerLinkState](mapCtx, in.LinkState)
	// MISSING: InviteLinkURI
	// MISSING: CreateTime
	// MISSING: UpdateTime
	// MISSING: PublicID
	// MISSING: ChannelPartnerCloudIdentityInfo
	return out
}
func ChannelPartnerLinkObservedState_FromProto(mapCtx *direct.MapContext, in *pb.ChannelPartnerLink) *krm.ChannelPartnerLinkObservedState {
	if in == nil {
		return nil
	}
	out := &krm.ChannelPartnerLinkObservedState{}
	out.Name = direct.LazyPtr(in.GetName())
	// MISSING: ResellerCloudIdentityID
	// MISSING: LinkState
	out.InviteLinkURI = direct.LazyPtr(in.GetInviteLinkUri())
	out.CreateTime = direct.StringTimestamp_FromProto(mapCtx, in.GetCreateTime())
	out.UpdateTime = direct.StringTimestamp_FromProto(mapCtx, in.GetUpdateTime())
	out.PublicID = direct.LazyPtr(in.GetPublicId())
	out.ChannelPartnerCloudIdentityInfo = CloudIdentityInfo_FromProto(mapCtx, in.GetChannelPartnerCloudIdentityInfo())
	return out
}
func ChannelPartnerLinkObservedState_ToProto(mapCtx *direct.MapContext, in *krm.ChannelPartnerLinkObservedState) *pb.ChannelPartnerLink {
	if in == nil {
		return nil
	}
	out := &pb.ChannelPartnerLink{}
	out.Name = direct.ValueOf(in.Name)
	// MISSING: ResellerCloudIdentityID
	// MISSING: LinkState
	out.InviteLinkUri = direct.ValueOf(in.InviteLinkURI)
	out.CreateTime = direct.StringTimestamp_ToProto(mapCtx, in.CreateTime)
	out.UpdateTime = direct.StringTimestamp_ToProto(mapCtx, in.UpdateTime)
	out.PublicId = direct.ValueOf(in.PublicID)
	out.ChannelPartnerCloudIdentityInfo = CloudIdentityInfo_ToProto(mapCtx, in.ChannelPartnerCloudIdentityInfo)
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
