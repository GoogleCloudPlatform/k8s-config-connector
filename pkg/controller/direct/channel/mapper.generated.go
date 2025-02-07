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
func AssociationInfo_FromProto(mapCtx *direct.MapContext, in *pb.AssociationInfo) *krm.AssociationInfo {
	if in == nil {
		return nil
	}
	out := &krm.AssociationInfo{}
	out.BaseEntitlement = direct.LazyPtr(in.GetBaseEntitlement())
	return out
}
func AssociationInfo_ToProto(mapCtx *direct.MapContext, in *krm.AssociationInfo) *pb.AssociationInfo {
	if in == nil {
		return nil
	}
	out := &pb.AssociationInfo{}
	out.BaseEntitlement = direct.ValueOf(in.BaseEntitlement)
	return out
}
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
func ChannelEntitlementObservedState_FromProto(mapCtx *direct.MapContext, in *pb.Entitlement) *krm.ChannelEntitlementObservedState {
	if in == nil {
		return nil
	}
	out := &krm.ChannelEntitlementObservedState{}
	// MISSING: Name
	// MISSING: CreateTime
	// MISSING: UpdateTime
	// MISSING: Offer
	// MISSING: CommitmentSettings
	// MISSING: ProvisioningState
	// MISSING: ProvisionedService
	// MISSING: SuspensionReasons
	// MISSING: PurchaseOrderID
	// MISSING: TrialSettings
	// MISSING: AssociationInfo
	// MISSING: Parameters
	// MISSING: BillingAccount
	return out
}
func ChannelEntitlementObservedState_ToProto(mapCtx *direct.MapContext, in *krm.ChannelEntitlementObservedState) *pb.Entitlement {
	if in == nil {
		return nil
	}
	out := &pb.Entitlement{}
	// MISSING: Name
	// MISSING: CreateTime
	// MISSING: UpdateTime
	// MISSING: Offer
	// MISSING: CommitmentSettings
	// MISSING: ProvisioningState
	// MISSING: ProvisionedService
	// MISSING: SuspensionReasons
	// MISSING: PurchaseOrderID
	// MISSING: TrialSettings
	// MISSING: AssociationInfo
	// MISSING: Parameters
	// MISSING: BillingAccount
	return out
}
func ChannelEntitlementSpec_FromProto(mapCtx *direct.MapContext, in *pb.Entitlement) *krm.ChannelEntitlementSpec {
	if in == nil {
		return nil
	}
	out := &krm.ChannelEntitlementSpec{}
	// MISSING: Name
	// MISSING: CreateTime
	// MISSING: UpdateTime
	// MISSING: Offer
	// MISSING: CommitmentSettings
	// MISSING: ProvisioningState
	// MISSING: ProvisionedService
	// MISSING: SuspensionReasons
	// MISSING: PurchaseOrderID
	// MISSING: TrialSettings
	// MISSING: AssociationInfo
	// MISSING: Parameters
	// MISSING: BillingAccount
	return out
}
func ChannelEntitlementSpec_ToProto(mapCtx *direct.MapContext, in *krm.ChannelEntitlementSpec) *pb.Entitlement {
	if in == nil {
		return nil
	}
	out := &pb.Entitlement{}
	// MISSING: Name
	// MISSING: CreateTime
	// MISSING: UpdateTime
	// MISSING: Offer
	// MISSING: CommitmentSettings
	// MISSING: ProvisioningState
	// MISSING: ProvisionedService
	// MISSING: SuspensionReasons
	// MISSING: PurchaseOrderID
	// MISSING: TrialSettings
	// MISSING: AssociationInfo
	// MISSING: Parameters
	// MISSING: BillingAccount
	return out
}
func CommitmentSettings_FromProto(mapCtx *direct.MapContext, in *pb.CommitmentSettings) *krm.CommitmentSettings {
	if in == nil {
		return nil
	}
	out := &krm.CommitmentSettings{}
	// MISSING: StartTime
	// MISSING: EndTime
	out.RenewalSettings = RenewalSettings_FromProto(mapCtx, in.GetRenewalSettings())
	return out
}
func CommitmentSettings_ToProto(mapCtx *direct.MapContext, in *krm.CommitmentSettings) *pb.CommitmentSettings {
	if in == nil {
		return nil
	}
	out := &pb.CommitmentSettings{}
	// MISSING: StartTime
	// MISSING: EndTime
	out.RenewalSettings = RenewalSettings_ToProto(mapCtx, in.RenewalSettings)
	return out
}
func CommitmentSettingsObservedState_FromProto(mapCtx *direct.MapContext, in *pb.CommitmentSettings) *krm.CommitmentSettingsObservedState {
	if in == nil {
		return nil
	}
	out := &krm.CommitmentSettingsObservedState{}
	out.StartTime = direct.StringTimestamp_FromProto(mapCtx, in.GetStartTime())
	out.EndTime = direct.StringTimestamp_FromProto(mapCtx, in.GetEndTime())
	// MISSING: RenewalSettings
	return out
}
func CommitmentSettingsObservedState_ToProto(mapCtx *direct.MapContext, in *krm.CommitmentSettingsObservedState) *pb.CommitmentSettings {
	if in == nil {
		return nil
	}
	out := &pb.CommitmentSettings{}
	out.StartTime = direct.StringTimestamp_ToProto(mapCtx, in.StartTime)
	out.EndTime = direct.StringTimestamp_ToProto(mapCtx, in.EndTime)
	// MISSING: RenewalSettings
	return out
}
func Entitlement_FromProto(mapCtx *direct.MapContext, in *pb.Entitlement) *krm.Entitlement {
	if in == nil {
		return nil
	}
	out := &krm.Entitlement{}
	// MISSING: Name
	// MISSING: CreateTime
	// MISSING: UpdateTime
	out.Offer = direct.LazyPtr(in.GetOffer())
	out.CommitmentSettings = CommitmentSettings_FromProto(mapCtx, in.GetCommitmentSettings())
	// MISSING: ProvisioningState
	// MISSING: ProvisionedService
	// MISSING: SuspensionReasons
	out.PurchaseOrderID = direct.LazyPtr(in.GetPurchaseOrderId())
	// MISSING: TrialSettings
	out.AssociationInfo = AssociationInfo_FromProto(mapCtx, in.GetAssociationInfo())
	out.Parameters = direct.Slice_FromProto(mapCtx, in.Parameters, Parameter_FromProto)
	out.BillingAccount = direct.LazyPtr(in.GetBillingAccount())
	return out
}
func Entitlement_ToProto(mapCtx *direct.MapContext, in *krm.Entitlement) *pb.Entitlement {
	if in == nil {
		return nil
	}
	out := &pb.Entitlement{}
	// MISSING: Name
	// MISSING: CreateTime
	// MISSING: UpdateTime
	out.Offer = direct.ValueOf(in.Offer)
	out.CommitmentSettings = CommitmentSettings_ToProto(mapCtx, in.CommitmentSettings)
	// MISSING: ProvisioningState
	// MISSING: ProvisionedService
	// MISSING: SuspensionReasons
	out.PurchaseOrderId = direct.ValueOf(in.PurchaseOrderID)
	// MISSING: TrialSettings
	out.AssociationInfo = AssociationInfo_ToProto(mapCtx, in.AssociationInfo)
	out.Parameters = direct.Slice_ToProto(mapCtx, in.Parameters, Parameter_ToProto)
	out.BillingAccount = direct.ValueOf(in.BillingAccount)
	return out
}
func EntitlementObservedState_FromProto(mapCtx *direct.MapContext, in *pb.Entitlement) *krm.EntitlementObservedState {
	if in == nil {
		return nil
	}
	out := &krm.EntitlementObservedState{}
	out.Name = direct.LazyPtr(in.GetName())
	out.CreateTime = direct.StringTimestamp_FromProto(mapCtx, in.GetCreateTime())
	out.UpdateTime = direct.StringTimestamp_FromProto(mapCtx, in.GetUpdateTime())
	// MISSING: Offer
	out.CommitmentSettings = CommitmentSettingsObservedState_FromProto(mapCtx, in.GetCommitmentSettings())
	out.ProvisioningState = direct.Enum_FromProto(mapCtx, in.GetProvisioningState())
	out.ProvisionedService = ProvisionedService_FromProto(mapCtx, in.GetProvisionedService())
	out.SuspensionReasons = direct.EnumSlice_FromProto(mapCtx, in.SuspensionReasons)
	// MISSING: PurchaseOrderID
	out.TrialSettings = TrialSettings_FromProto(mapCtx, in.GetTrialSettings())
	// MISSING: AssociationInfo
	out.Parameters = direct.Slice_FromProto(mapCtx, in.Parameters, ParameterObservedState_FromProto)
	// MISSING: BillingAccount
	return out
}
func EntitlementObservedState_ToProto(mapCtx *direct.MapContext, in *krm.EntitlementObservedState) *pb.Entitlement {
	if in == nil {
		return nil
	}
	out := &pb.Entitlement{}
	out.Name = direct.ValueOf(in.Name)
	out.CreateTime = direct.StringTimestamp_ToProto(mapCtx, in.CreateTime)
	out.UpdateTime = direct.StringTimestamp_ToProto(mapCtx, in.UpdateTime)
	// MISSING: Offer
	out.CommitmentSettings = CommitmentSettingsObservedState_ToProto(mapCtx, in.CommitmentSettings)
	out.ProvisioningState = direct.Enum_ToProto[pb.Entitlement_ProvisioningState](mapCtx, in.ProvisioningState)
	out.ProvisionedService = ProvisionedService_ToProto(mapCtx, in.ProvisionedService)
	out.SuspensionReasons = direct.EnumSlice_ToProto[pb.Entitlement_SuspensionReason](mapCtx, in.SuspensionReasons)
	// MISSING: PurchaseOrderID
	out.TrialSettings = TrialSettings_ToProto(mapCtx, in.TrialSettings)
	// MISSING: AssociationInfo
	out.Parameters = direct.Slice_ToProto(mapCtx, in.Parameters, ParameterObservedState_ToProto)
	// MISSING: BillingAccount
	return out
}
func Parameter_FromProto(mapCtx *direct.MapContext, in *pb.Parameter) *krm.Parameter {
	if in == nil {
		return nil
	}
	out := &krm.Parameter{}
	out.Name = direct.LazyPtr(in.GetName())
	out.Value = Value_FromProto(mapCtx, in.GetValue())
	// MISSING: Editable
	return out
}
func Parameter_ToProto(mapCtx *direct.MapContext, in *krm.Parameter) *pb.Parameter {
	if in == nil {
		return nil
	}
	out := &pb.Parameter{}
	out.Name = direct.ValueOf(in.Name)
	out.Value = Value_ToProto(mapCtx, in.Value)
	// MISSING: Editable
	return out
}
func ParameterObservedState_FromProto(mapCtx *direct.MapContext, in *pb.Parameter) *krm.ParameterObservedState {
	if in == nil {
		return nil
	}
	out := &krm.ParameterObservedState{}
	// MISSING: Name
	// MISSING: Value
	out.Editable = direct.LazyPtr(in.GetEditable())
	return out
}
func ParameterObservedState_ToProto(mapCtx *direct.MapContext, in *krm.ParameterObservedState) *pb.Parameter {
	if in == nil {
		return nil
	}
	out := &pb.Parameter{}
	// MISSING: Name
	// MISSING: Value
	out.Editable = direct.ValueOf(in.Editable)
	return out
}
func Period_FromProto(mapCtx *direct.MapContext, in *pb.Period) *krm.Period {
	if in == nil {
		return nil
	}
	out := &krm.Period{}
	out.Duration = direct.LazyPtr(in.GetDuration())
	out.PeriodType = direct.Enum_FromProto(mapCtx, in.GetPeriodType())
	return out
}
func Period_ToProto(mapCtx *direct.MapContext, in *krm.Period) *pb.Period {
	if in == nil {
		return nil
	}
	out := &pb.Period{}
	out.Duration = direct.ValueOf(in.Duration)
	out.PeriodType = direct.Enum_ToProto[pb.PeriodType](mapCtx, in.PeriodType)
	return out
}
func ProvisionedService_FromProto(mapCtx *direct.MapContext, in *pb.ProvisionedService) *krm.ProvisionedService {
	if in == nil {
		return nil
	}
	out := &krm.ProvisionedService{}
	// MISSING: ProvisioningID
	// MISSING: ProductID
	// MISSING: SkuID
	return out
}
func ProvisionedService_ToProto(mapCtx *direct.MapContext, in *krm.ProvisionedService) *pb.ProvisionedService {
	if in == nil {
		return nil
	}
	out := &pb.ProvisionedService{}
	// MISSING: ProvisioningID
	// MISSING: ProductID
	// MISSING: SkuID
	return out
}
func ProvisionedServiceObservedState_FromProto(mapCtx *direct.MapContext, in *pb.ProvisionedService) *krm.ProvisionedServiceObservedState {
	if in == nil {
		return nil
	}
	out := &krm.ProvisionedServiceObservedState{}
	out.ProvisioningID = direct.LazyPtr(in.GetProvisioningId())
	out.ProductID = direct.LazyPtr(in.GetProductId())
	out.SkuID = direct.LazyPtr(in.GetSkuId())
	return out
}
func ProvisionedServiceObservedState_ToProto(mapCtx *direct.MapContext, in *krm.ProvisionedServiceObservedState) *pb.ProvisionedService {
	if in == nil {
		return nil
	}
	out := &pb.ProvisionedService{}
	out.ProvisioningId = direct.ValueOf(in.ProvisioningID)
	out.ProductId = direct.ValueOf(in.ProductID)
	out.SkuId = direct.ValueOf(in.SkuID)
	return out
}
func RenewalSettings_FromProto(mapCtx *direct.MapContext, in *pb.RenewalSettings) *krm.RenewalSettings {
	if in == nil {
		return nil
	}
	out := &krm.RenewalSettings{}
	out.EnableRenewal = direct.LazyPtr(in.GetEnableRenewal())
	out.ResizeUnitCount = direct.LazyPtr(in.GetResizeUnitCount())
	out.PaymentPlan = direct.Enum_FromProto(mapCtx, in.GetPaymentPlan())
	out.PaymentCycle = Period_FromProto(mapCtx, in.GetPaymentCycle())
	return out
}
func RenewalSettings_ToProto(mapCtx *direct.MapContext, in *krm.RenewalSettings) *pb.RenewalSettings {
	if in == nil {
		return nil
	}
	out := &pb.RenewalSettings{}
	out.EnableRenewal = direct.ValueOf(in.EnableRenewal)
	out.ResizeUnitCount = direct.ValueOf(in.ResizeUnitCount)
	out.PaymentPlan = direct.Enum_ToProto[pb.PaymentPlan](mapCtx, in.PaymentPlan)
	out.PaymentCycle = Period_ToProto(mapCtx, in.PaymentCycle)
	return out
}
func TrialSettings_FromProto(mapCtx *direct.MapContext, in *pb.TrialSettings) *krm.TrialSettings {
	if in == nil {
		return nil
	}
	out := &krm.TrialSettings{}
	out.Trial = direct.LazyPtr(in.GetTrial())
	out.EndTime = direct.StringTimestamp_FromProto(mapCtx, in.GetEndTime())
	return out
}
func TrialSettings_ToProto(mapCtx *direct.MapContext, in *krm.TrialSettings) *pb.TrialSettings {
	if in == nil {
		return nil
	}
	out := &pb.TrialSettings{}
	out.Trial = direct.ValueOf(in.Trial)
	out.EndTime = direct.StringTimestamp_ToProto(mapCtx, in.EndTime)
	return out
}
func Value_FromProto(mapCtx *direct.MapContext, in *pb.Value) *krm.Value {
	if in == nil {
		return nil
	}
	out := &krm.Value{}
	out.Int64Value = direct.LazyPtr(in.GetInt64Value())
	out.StringValue = direct.LazyPtr(in.GetStringValue())
	out.DoubleValue = direct.LazyPtr(in.GetDoubleValue())
	out.ProtoValue = Any_FromProto(mapCtx, in.GetProtoValue())
	out.BoolValue = direct.LazyPtr(in.GetBoolValue())
	return out
}
func Value_ToProto(mapCtx *direct.MapContext, in *krm.Value) *pb.Value {
	if in == nil {
		return nil
	}
	out := &pb.Value{}
	if oneof := Value_Int64Value_ToProto(mapCtx, in.Int64Value); oneof != nil {
		out.Kind = oneof
	}
	if oneof := Value_StringValue_ToProto(mapCtx, in.StringValue); oneof != nil {
		out.Kind = oneof
	}
	if oneof := Value_DoubleValue_ToProto(mapCtx, in.DoubleValue); oneof != nil {
		out.Kind = oneof
	}
	if oneof := Any_ToProto(mapCtx, in.ProtoValue); oneof != nil {
		out.Kind = &pb.Value_ProtoValue{ProtoValue: oneof}
	}
	if oneof := Value_BoolValue_ToProto(mapCtx, in.BoolValue); oneof != nil {
		out.Kind = oneof
	}
	return out
}
