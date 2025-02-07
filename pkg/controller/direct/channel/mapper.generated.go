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
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct"
	refs "github.com/GoogleCloudPlatform/k8s-config-connector/apis/refs/v1beta1"
	pb "cloud.google.com/go/channel/apiv1/channelpb"
	krm "github.com/GoogleCloudPlatform/k8s-config-connector/apis/channel/v1alpha1"
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
func ChannelCustomerRepricingConfigObservedState_FromProto(mapCtx *direct.MapContext, in *pb.CustomerRepricingConfig) *krm.ChannelCustomerRepricingConfigObservedState {
	if in == nil {
		return nil
	}
	out := &krm.ChannelCustomerRepricingConfigObservedState{}
	// MISSING: Name
	// MISSING: RepricingConfig
	// MISSING: UpdateTime
	return out
}
func ChannelCustomerRepricingConfigObservedState_ToProto(mapCtx *direct.MapContext, in *krm.ChannelCustomerRepricingConfigObservedState) *pb.CustomerRepricingConfig {
	if in == nil {
		return nil
	}
	out := &pb.CustomerRepricingConfig{}
	// MISSING: Name
	// MISSING: RepricingConfig
	// MISSING: UpdateTime
	return out
}
func ChannelCustomerRepricingConfigSpec_FromProto(mapCtx *direct.MapContext, in *pb.CustomerRepricingConfig) *krm.ChannelCustomerRepricingConfigSpec {
	if in == nil {
		return nil
	}
	out := &krm.ChannelCustomerRepricingConfigSpec{}
	// MISSING: Name
	// MISSING: RepricingConfig
	// MISSING: UpdateTime
	return out
}
func ChannelCustomerRepricingConfigSpec_ToProto(mapCtx *direct.MapContext, in *krm.ChannelCustomerRepricingConfigSpec) *pb.CustomerRepricingConfig {
	if in == nil {
		return nil
	}
	out := &pb.CustomerRepricingConfig{}
	// MISSING: Name
	// MISSING: RepricingConfig
	// MISSING: UpdateTime
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
func ChannelOfferObservedState_FromProto(mapCtx *direct.MapContext, in *pb.Offer) *krm.ChannelOfferObservedState {
	if in == nil {
		return nil
	}
	out := &krm.ChannelOfferObservedState{}
	// MISSING: Name
	// MISSING: MarketingInfo
	// MISSING: Sku
	// MISSING: Plan
	// MISSING: Constraints
	// MISSING: PriceByResources
	// MISSING: StartTime
	// MISSING: EndTime
	// MISSING: ParameterDefinitions
	// MISSING: DealCode
	return out
}
func ChannelOfferObservedState_ToProto(mapCtx *direct.MapContext, in *krm.ChannelOfferObservedState) *pb.Offer {
	if in == nil {
		return nil
	}
	out := &pb.Offer{}
	// MISSING: Name
	// MISSING: MarketingInfo
	// MISSING: Sku
	// MISSING: Plan
	// MISSING: Constraints
	// MISSING: PriceByResources
	// MISSING: StartTime
	// MISSING: EndTime
	// MISSING: ParameterDefinitions
	// MISSING: DealCode
	return out
}
func ChannelOfferSpec_FromProto(mapCtx *direct.MapContext, in *pb.Offer) *krm.ChannelOfferSpec {
	if in == nil {
		return nil
	}
	out := &krm.ChannelOfferSpec{}
	// MISSING: Name
	// MISSING: MarketingInfo
	// MISSING: Sku
	// MISSING: Plan
	// MISSING: Constraints
	// MISSING: PriceByResources
	// MISSING: StartTime
	// MISSING: EndTime
	// MISSING: ParameterDefinitions
	// MISSING: DealCode
	return out
}
func ChannelOfferSpec_ToProto(mapCtx *direct.MapContext, in *krm.ChannelOfferSpec) *pb.Offer {
	if in == nil {
		return nil
	}
	out := &pb.Offer{}
	// MISSING: Name
	// MISSING: MarketingInfo
	// MISSING: Sku
	// MISSING: Plan
	// MISSING: Constraints
	// MISSING: PriceByResources
	// MISSING: StartTime
	// MISSING: EndTime
	// MISSING: ParameterDefinitions
	// MISSING: DealCode
	return out
}
func ChannelProductObservedState_FromProto(mapCtx *direct.MapContext, in *pb.Product) *krm.ChannelProductObservedState {
	if in == nil {
		return nil
	}
	out := &krm.ChannelProductObservedState{}
	// MISSING: Name
	// MISSING: MarketingInfo
	return out
}
func ChannelProductObservedState_ToProto(mapCtx *direct.MapContext, in *krm.ChannelProductObservedState) *pb.Product {
	if in == nil {
		return nil
	}
	out := &pb.Product{}
	// MISSING: Name
	// MISSING: MarketingInfo
	return out
}
func ChannelProductSpec_FromProto(mapCtx *direct.MapContext, in *pb.Product) *krm.ChannelProductSpec {
	if in == nil {
		return nil
	}
	out := &krm.ChannelProductSpec{}
	// MISSING: Name
	// MISSING: MarketingInfo
	return out
}
func ChannelProductSpec_ToProto(mapCtx *direct.MapContext, in *krm.ChannelProductSpec) *pb.Product {
	if in == nil {
		return nil
	}
	out := &pb.Product{}
	// MISSING: Name
	// MISSING: MarketingInfo
	return out
}
func ChannelReportJobObservedState_FromProto(mapCtx *direct.MapContext, in *pb.ReportJob) *krm.ChannelReportJobObservedState {
	if in == nil {
		return nil
	}
	out := &krm.ChannelReportJobObservedState{}
	// MISSING: Name
	// MISSING: ReportStatus
	return out
}
func ChannelReportJobObservedState_ToProto(mapCtx *direct.MapContext, in *krm.ChannelReportJobObservedState) *pb.ReportJob {
	if in == nil {
		return nil
	}
	out := &pb.ReportJob{}
	// MISSING: Name
	// MISSING: ReportStatus
	return out
}
func ChannelReportJobSpec_FromProto(mapCtx *direct.MapContext, in *pb.ReportJob) *krm.ChannelReportJobSpec {
	if in == nil {
		return nil
	}
	out := &krm.ChannelReportJobSpec{}
	// MISSING: Name
	// MISSING: ReportStatus
	return out
}
func ChannelReportJobSpec_ToProto(mapCtx *direct.MapContext, in *krm.ChannelReportJobSpec) *pb.ReportJob {
	if in == nil {
		return nil
	}
	out := &pb.ReportJob{}
	// MISSING: Name
	// MISSING: ReportStatus
	return out
}
func ChannelReportObservedState_FromProto(mapCtx *direct.MapContext, in *pb.Report) *krm.ChannelReportObservedState {
	if in == nil {
		return nil
	}
	out := &krm.ChannelReportObservedState{}
	// MISSING: Name
	// MISSING: DisplayName
	// MISSING: Columns
	// MISSING: Description
	return out
}
func ChannelReportObservedState_ToProto(mapCtx *direct.MapContext, in *krm.ChannelReportObservedState) *pb.Report {
	if in == nil {
		return nil
	}
	out := &pb.Report{}
	// MISSING: Name
	// MISSING: DisplayName
	// MISSING: Columns
	// MISSING: Description
	return out
}
func ChannelReportSpec_FromProto(mapCtx *direct.MapContext, in *pb.Report) *krm.ChannelReportSpec {
	if in == nil {
		return nil
	}
	out := &krm.ChannelReportSpec{}
	// MISSING: Name
	// MISSING: DisplayName
	// MISSING: Columns
	// MISSING: Description
	return out
}
func ChannelReportSpec_ToProto(mapCtx *direct.MapContext, in *krm.ChannelReportSpec) *pb.Report {
	if in == nil {
		return nil
	}
	out := &pb.Report{}
	// MISSING: Name
	// MISSING: DisplayName
	// MISSING: Columns
	// MISSING: Description
	return out
}
func ChannelSkuObservedState_FromProto(mapCtx *direct.MapContext, in *pb.Sku) *krm.ChannelSkuObservedState {
	if in == nil {
		return nil
	}
	out := &krm.ChannelSkuObservedState{}
	// MISSING: Name
	// MISSING: MarketingInfo
	// MISSING: Product
	return out
}
func ChannelSkuObservedState_ToProto(mapCtx *direct.MapContext, in *krm.ChannelSkuObservedState) *pb.Sku {
	if in == nil {
		return nil
	}
	out := &pb.Sku{}
	// MISSING: Name
	// MISSING: MarketingInfo
	// MISSING: Product
	return out
}
func ChannelSkuSpec_FromProto(mapCtx *direct.MapContext, in *pb.Sku) *krm.ChannelSkuSpec {
	if in == nil {
		return nil
	}
	out := &krm.ChannelSkuSpec{}
	// MISSING: Name
	// MISSING: MarketingInfo
	// MISSING: Product
	return out
}
func ChannelSkuSpec_ToProto(mapCtx *direct.MapContext, in *krm.ChannelSkuSpec) *pb.Sku {
	if in == nil {
		return nil
	}
	out := &pb.Sku{}
	// MISSING: Name
	// MISSING: MarketingInfo
	// MISSING: Product
	return out
}
func ConditionalOverride_FromProto(mapCtx *direct.MapContext, in *pb.ConditionalOverride) *krm.ConditionalOverride {
	if in == nil {
		return nil
	}
	out := &krm.ConditionalOverride{}
	out.Adjustment = RepricingAdjustment_FromProto(mapCtx, in.GetAdjustment())
	out.RebillingBasis = direct.Enum_FromProto(mapCtx, in.GetRebillingBasis())
	out.RepricingCondition = RepricingCondition_FromProto(mapCtx, in.GetRepricingCondition())
	return out
}
func ConditionalOverride_ToProto(mapCtx *direct.MapContext, in *krm.ConditionalOverride) *pb.ConditionalOverride {
	if in == nil {
		return nil
	}
	out := &pb.ConditionalOverride{}
	out.Adjustment = RepricingAdjustment_ToProto(mapCtx, in.Adjustment)
	out.RebillingBasis = direct.Enum_ToProto[pb.RebillingBasis](mapCtx, in.RebillingBasis)
	out.RepricingCondition = RepricingCondition_ToProto(mapCtx, in.RepricingCondition)
	return out
}
func CustomerRepricingConfig_FromProto(mapCtx *direct.MapContext, in *pb.CustomerRepricingConfig) *krm.CustomerRepricingConfig {
	if in == nil {
		return nil
	}
	out := &krm.CustomerRepricingConfig{}
	// MISSING: Name
	out.RepricingConfig = RepricingConfig_FromProto(mapCtx, in.GetRepricingConfig())
	// MISSING: UpdateTime
	return out
}
func CustomerRepricingConfig_ToProto(mapCtx *direct.MapContext, in *krm.CustomerRepricingConfig) *pb.CustomerRepricingConfig {
	if in == nil {
		return nil
	}
	out := &pb.CustomerRepricingConfig{}
	// MISSING: Name
	out.RepricingConfig = RepricingConfig_ToProto(mapCtx, in.RepricingConfig)
	// MISSING: UpdateTime
	return out
}
func CustomerRepricingConfigObservedState_FromProto(mapCtx *direct.MapContext, in *pb.CustomerRepricingConfig) *krm.CustomerRepricingConfigObservedState {
	if in == nil {
		return nil
	}
	out := &krm.CustomerRepricingConfigObservedState{}
	out.Name = direct.LazyPtr(in.GetName())
	// MISSING: RepricingConfig
	out.UpdateTime = direct.StringTimestamp_FromProto(mapCtx, in.GetUpdateTime())
	return out
}
func CustomerRepricingConfigObservedState_ToProto(mapCtx *direct.MapContext, in *krm.CustomerRepricingConfigObservedState) *pb.CustomerRepricingConfig {
	if in == nil {
		return nil
	}
	out := &pb.CustomerRepricingConfig{}
	out.Name = direct.ValueOf(in.Name)
	// MISSING: RepricingConfig
	out.UpdateTime = direct.StringTimestamp_ToProto(mapCtx, in.UpdateTime)
	return out
}
func PercentageAdjustment_FromProto(mapCtx *direct.MapContext, in *pb.PercentageAdjustment) *krm.PercentageAdjustment {
	if in == nil {
		return nil
	}
	out := &krm.PercentageAdjustment{}
	out.Percentage = Decimal_FromProto(mapCtx, in.GetPercentage())
	return out
}
func PercentageAdjustment_ToProto(mapCtx *direct.MapContext, in *krm.PercentageAdjustment) *pb.PercentageAdjustment {
	if in == nil {
		return nil
	}
	out := &pb.PercentageAdjustment{}
	out.Percentage = Decimal_ToProto(mapCtx, in.Percentage)
	return out
}
func RepricingAdjustment_FromProto(mapCtx *direct.MapContext, in *pb.RepricingAdjustment) *krm.RepricingAdjustment {
	if in == nil {
		return nil
	}
	out := &krm.RepricingAdjustment{}
	out.PercentageAdjustment = PercentageAdjustment_FromProto(mapCtx, in.GetPercentageAdjustment())
	return out
}
func RepricingAdjustment_ToProto(mapCtx *direct.MapContext, in *krm.RepricingAdjustment) *pb.RepricingAdjustment {
	if in == nil {
		return nil
	}
	out := &pb.RepricingAdjustment{}
	if oneof := PercentageAdjustment_ToProto(mapCtx, in.PercentageAdjustment); oneof != nil {
		out.Adjustment = &pb.RepricingAdjustment_PercentageAdjustment{PercentageAdjustment: oneof}
	}
	return out
}
func RepricingCondition_FromProto(mapCtx *direct.MapContext, in *pb.RepricingCondition) *krm.RepricingCondition {
	if in == nil {
		return nil
	}
	out := &krm.RepricingCondition{}
	out.SkuGroupCondition = SkuGroupCondition_FromProto(mapCtx, in.GetSkuGroupCondition())
	return out
}
func RepricingCondition_ToProto(mapCtx *direct.MapContext, in *krm.RepricingCondition) *pb.RepricingCondition {
	if in == nil {
		return nil
	}
	out := &pb.RepricingCondition{}
	if oneof := SkuGroupCondition_ToProto(mapCtx, in.SkuGroupCondition); oneof != nil {
		out.Condition = &pb.RepricingCondition_SkuGroupCondition{SkuGroupCondition: oneof}
	}
	return out
}
func RepricingConfig_FromProto(mapCtx *direct.MapContext, in *pb.RepricingConfig) *krm.RepricingConfig {
	if in == nil {
		return nil
	}
	out := &krm.RepricingConfig{}
	out.EntitlementGranularity = RepricingConfig_EntitlementGranularity_FromProto(mapCtx, in.GetEntitlementGranularity())
	out.ChannelPartnerGranularity = RepricingConfig_ChannelPartnerGranularity_FromProto(mapCtx, in.GetChannelPartnerGranularity())
	out.EffectiveInvoiceMonth = Date_FromProto(mapCtx, in.GetEffectiveInvoiceMonth())
	out.Adjustment = RepricingAdjustment_FromProto(mapCtx, in.GetAdjustment())
	out.RebillingBasis = direct.Enum_FromProto(mapCtx, in.GetRebillingBasis())
	out.ConditionalOverrides = direct.Slice_FromProto(mapCtx, in.ConditionalOverrides, ConditionalOverride_FromProto)
	return out
}
func RepricingConfig_ToProto(mapCtx *direct.MapContext, in *krm.RepricingConfig) *pb.RepricingConfig {
	if in == nil {
		return nil
	}
	out := &pb.RepricingConfig{}
	if oneof := RepricingConfig_EntitlementGranularity_ToProto(mapCtx, in.EntitlementGranularity); oneof != nil {
		out.Granularity = &pb.RepricingConfig_EntitlementGranularity_{EntitlementGranularity: oneof}
	}
	if oneof := RepricingConfig_ChannelPartnerGranularity_ToProto(mapCtx, in.ChannelPartnerGranularity); oneof != nil {
		out.Granularity = &pb.RepricingConfig_ChannelPartnerGranularity_{ChannelPartnerGranularity: oneof}
	}
	out.EffectiveInvoiceMonth = Date_ToProto(mapCtx, in.EffectiveInvoiceMonth)
	out.Adjustment = RepricingAdjustment_ToProto(mapCtx, in.Adjustment)
	out.RebillingBasis = direct.Enum_ToProto[pb.RebillingBasis](mapCtx, in.RebillingBasis)
	out.ConditionalOverrides = direct.Slice_ToProto(mapCtx, in.ConditionalOverrides, ConditionalOverride_ToProto)
	return out
}
func RepricingConfig_ChannelPartnerGranularity_FromProto(mapCtx *direct.MapContext, in *pb.RepricingConfig_ChannelPartnerGranularity) *krm.RepricingConfig_ChannelPartnerGranularity {
	if in == nil {
		return nil
	}
	out := &krm.RepricingConfig_ChannelPartnerGranularity{}
	return out
}
func RepricingConfig_ChannelPartnerGranularity_ToProto(mapCtx *direct.MapContext, in *krm.RepricingConfig_ChannelPartnerGranularity) *pb.RepricingConfig_ChannelPartnerGranularity {
	if in == nil {
		return nil
	}
	out := &pb.RepricingConfig_ChannelPartnerGranularity{}
	return out
}
func RepricingConfig_EntitlementGranularity_FromProto(mapCtx *direct.MapContext, in *pb.RepricingConfig_EntitlementGranularity) *krm.RepricingConfig_EntitlementGranularity {
	if in == nil {
		return nil
	}
	out := &krm.RepricingConfig_EntitlementGranularity{}
	out.Entitlement = direct.LazyPtr(in.GetEntitlement())
	return out
}
func RepricingConfig_EntitlementGranularity_ToProto(mapCtx *direct.MapContext, in *krm.RepricingConfig_EntitlementGranularity) *pb.RepricingConfig_EntitlementGranularity {
	if in == nil {
		return nil
	}
	out := &pb.RepricingConfig_EntitlementGranularity{}
	out.Entitlement = direct.ValueOf(in.Entitlement)
	return out
}
func SkuGroupCondition_FromProto(mapCtx *direct.MapContext, in *pb.SkuGroupCondition) *krm.SkuGroupCondition {
	if in == nil {
		return nil
	}
	out := &krm.SkuGroupCondition{}
	out.SkuGroup = direct.LazyPtr(in.GetSkuGroup())
	return out
}
func SkuGroupCondition_ToProto(mapCtx *direct.MapContext, in *krm.SkuGroupCondition) *pb.SkuGroupCondition {
	if in == nil {
		return nil
	}
	out := &pb.SkuGroupCondition{}
	out.SkuGroup = direct.ValueOf(in.SkuGroup)
	return out
}
