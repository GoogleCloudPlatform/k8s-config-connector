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
func Constraints_FromProto(mapCtx *direct.MapContext, in *pb.Constraints) *krm.Constraints {
	if in == nil {
		return nil
	}
	out := &krm.Constraints{}
	out.CustomerConstraints = CustomerConstraints_FromProto(mapCtx, in.GetCustomerConstraints())
	return out
}
func Constraints_ToProto(mapCtx *direct.MapContext, in *krm.Constraints) *pb.Constraints {
	if in == nil {
		return nil
	}
	out := &pb.Constraints{}
	out.CustomerConstraints = CustomerConstraints_ToProto(mapCtx, in.CustomerConstraints)
	return out
}
func CustomerConstraints_FromProto(mapCtx *direct.MapContext, in *pb.CustomerConstraints) *krm.CustomerConstraints {
	if in == nil {
		return nil
	}
	out := &krm.CustomerConstraints{}
	out.AllowedRegions = in.AllowedRegions
	out.AllowedCustomerTypes = direct.EnumSlice_FromProto(mapCtx, in.AllowedCustomerTypes)
	out.PromotionalOrderTypes = direct.EnumSlice_FromProto(mapCtx, in.PromotionalOrderTypes)
	return out
}
func CustomerConstraints_ToProto(mapCtx *direct.MapContext, in *krm.CustomerConstraints) *pb.CustomerConstraints {
	if in == nil {
		return nil
	}
	out := &pb.CustomerConstraints{}
	out.AllowedRegions = in.AllowedRegions
	out.AllowedCustomerTypes = direct.EnumSlice_ToProto[pb.CloudIdentityInfo_CustomerType](mapCtx, in.AllowedCustomerTypes)
	out.PromotionalOrderTypes = direct.EnumSlice_ToProto[pb.PromotionalOrderType](mapCtx, in.PromotionalOrderTypes)
	return out
}
func MarketingInfo_FromProto(mapCtx *direct.MapContext, in *pb.MarketingInfo) *krm.MarketingInfo {
	if in == nil {
		return nil
	}
	out := &krm.MarketingInfo{}
	out.DisplayName = direct.LazyPtr(in.GetDisplayName())
	out.Description = direct.LazyPtr(in.GetDescription())
	out.DefaultLogo = Media_FromProto(mapCtx, in.GetDefaultLogo())
	return out
}
func MarketingInfo_ToProto(mapCtx *direct.MapContext, in *krm.MarketingInfo) *pb.MarketingInfo {
	if in == nil {
		return nil
	}
	out := &pb.MarketingInfo{}
	out.DisplayName = direct.ValueOf(in.DisplayName)
	out.Description = direct.ValueOf(in.Description)
	out.DefaultLogo = Media_ToProto(mapCtx, in.DefaultLogo)
	return out
}
func Media_FromProto(mapCtx *direct.MapContext, in *pb.Media) *krm.Media {
	if in == nil {
		return nil
	}
	out := &krm.Media{}
	out.Title = direct.LazyPtr(in.GetTitle())
	out.Content = direct.LazyPtr(in.GetContent())
	out.Type = direct.Enum_FromProto(mapCtx, in.GetType())
	return out
}
func Media_ToProto(mapCtx *direct.MapContext, in *krm.Media) *pb.Media {
	if in == nil {
		return nil
	}
	out := &pb.Media{}
	out.Title = direct.ValueOf(in.Title)
	out.Content = direct.ValueOf(in.Content)
	out.Type = direct.Enum_ToProto[pb.MediaType](mapCtx, in.Type)
	return out
}
func Offer_FromProto(mapCtx *direct.MapContext, in *pb.Offer) *krm.Offer {
	if in == nil {
		return nil
	}
	out := &krm.Offer{}
	out.Name = direct.LazyPtr(in.GetName())
	out.MarketingInfo = MarketingInfo_FromProto(mapCtx, in.GetMarketingInfo())
	out.Sku = Sku_FromProto(mapCtx, in.GetSku())
	out.Plan = Plan_FromProto(mapCtx, in.GetPlan())
	out.Constraints = Constraints_FromProto(mapCtx, in.GetConstraints())
	out.PriceByResources = direct.Slice_FromProto(mapCtx, in.PriceByResources, PriceByResource_FromProto)
	out.StartTime = direct.StringTimestamp_FromProto(mapCtx, in.GetStartTime())
	// MISSING: EndTime
	out.ParameterDefinitions = direct.Slice_FromProto(mapCtx, in.ParameterDefinitions, ParameterDefinition_FromProto)
	out.DealCode = direct.LazyPtr(in.GetDealCode())
	return out
}
func Offer_ToProto(mapCtx *direct.MapContext, in *krm.Offer) *pb.Offer {
	if in == nil {
		return nil
	}
	out := &pb.Offer{}
	out.Name = direct.ValueOf(in.Name)
	out.MarketingInfo = MarketingInfo_ToProto(mapCtx, in.MarketingInfo)
	out.Sku = Sku_ToProto(mapCtx, in.Sku)
	out.Plan = Plan_ToProto(mapCtx, in.Plan)
	out.Constraints = Constraints_ToProto(mapCtx, in.Constraints)
	out.PriceByResources = direct.Slice_ToProto(mapCtx, in.PriceByResources, PriceByResource_ToProto)
	out.StartTime = direct.StringTimestamp_ToProto(mapCtx, in.StartTime)
	// MISSING: EndTime
	out.ParameterDefinitions = direct.Slice_ToProto(mapCtx, in.ParameterDefinitions, ParameterDefinition_ToProto)
	out.DealCode = direct.ValueOf(in.DealCode)
	return out
}
func OfferObservedState_FromProto(mapCtx *direct.MapContext, in *pb.Offer) *krm.OfferObservedState {
	if in == nil {
		return nil
	}
	out := &krm.OfferObservedState{}
	// MISSING: Name
	// MISSING: MarketingInfo
	// MISSING: Sku
	// MISSING: Plan
	// MISSING: Constraints
	// MISSING: PriceByResources
	// MISSING: StartTime
	out.EndTime = direct.StringTimestamp_FromProto(mapCtx, in.GetEndTime())
	// MISSING: ParameterDefinitions
	// MISSING: DealCode
	return out
}
func OfferObservedState_ToProto(mapCtx *direct.MapContext, in *krm.OfferObservedState) *pb.Offer {
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
	out.EndTime = direct.StringTimestamp_ToProto(mapCtx, in.EndTime)
	// MISSING: ParameterDefinitions
	// MISSING: DealCode
	return out
}
func ParameterDefinition_FromProto(mapCtx *direct.MapContext, in *pb.ParameterDefinition) *krm.ParameterDefinition {
	if in == nil {
		return nil
	}
	out := &krm.ParameterDefinition{}
	out.Name = direct.LazyPtr(in.GetName())
	out.ParameterType = direct.Enum_FromProto(mapCtx, in.GetParameterType())
	out.MinValue = Value_FromProto(mapCtx, in.GetMinValue())
	out.MaxValue = Value_FromProto(mapCtx, in.GetMaxValue())
	out.AllowedValues = direct.Slice_FromProto(mapCtx, in.AllowedValues, Value_FromProto)
	out.Optional = direct.LazyPtr(in.GetOptional())
	return out
}
func ParameterDefinition_ToProto(mapCtx *direct.MapContext, in *krm.ParameterDefinition) *pb.ParameterDefinition {
	if in == nil {
		return nil
	}
	out := &pb.ParameterDefinition{}
	out.Name = direct.ValueOf(in.Name)
	out.ParameterType = direct.Enum_ToProto[pb.ParameterDefinition_ParameterType](mapCtx, in.ParameterType)
	out.MinValue = Value_ToProto(mapCtx, in.MinValue)
	out.MaxValue = Value_ToProto(mapCtx, in.MaxValue)
	out.AllowedValues = direct.Slice_ToProto(mapCtx, in.AllowedValues, Value_ToProto)
	out.Optional = direct.ValueOf(in.Optional)
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
func Plan_FromProto(mapCtx *direct.MapContext, in *pb.Plan) *krm.Plan {
	if in == nil {
		return nil
	}
	out := &krm.Plan{}
	out.PaymentPlan = direct.Enum_FromProto(mapCtx, in.GetPaymentPlan())
	out.PaymentType = direct.Enum_FromProto(mapCtx, in.GetPaymentType())
	out.PaymentCycle = Period_FromProto(mapCtx, in.GetPaymentCycle())
	out.TrialPeriod = Period_FromProto(mapCtx, in.GetTrialPeriod())
	out.BillingAccount = direct.LazyPtr(in.GetBillingAccount())
	return out
}
func Plan_ToProto(mapCtx *direct.MapContext, in *krm.Plan) *pb.Plan {
	if in == nil {
		return nil
	}
	out := &pb.Plan{}
	out.PaymentPlan = direct.Enum_ToProto[pb.PaymentPlan](mapCtx, in.PaymentPlan)
	out.PaymentType = direct.Enum_ToProto[pb.PaymentType](mapCtx, in.PaymentType)
	out.PaymentCycle = Period_ToProto(mapCtx, in.PaymentCycle)
	out.TrialPeriod = Period_ToProto(mapCtx, in.TrialPeriod)
	out.BillingAccount = direct.ValueOf(in.BillingAccount)
	return out
}
func Price_FromProto(mapCtx *direct.MapContext, in *pb.Price) *krm.Price {
	if in == nil {
		return nil
	}
	out := &krm.Price{}
	out.BasePrice = Money_FromProto(mapCtx, in.GetBasePrice())
	out.Discount = direct.LazyPtr(in.GetDiscount())
	out.EffectivePrice = Money_FromProto(mapCtx, in.GetEffectivePrice())
	out.ExternalPriceURI = direct.LazyPtr(in.GetExternalPriceUri())
	return out
}
func Price_ToProto(mapCtx *direct.MapContext, in *krm.Price) *pb.Price {
	if in == nil {
		return nil
	}
	out := &pb.Price{}
	out.BasePrice = Money_ToProto(mapCtx, in.BasePrice)
	out.Discount = direct.ValueOf(in.Discount)
	out.EffectivePrice = Money_ToProto(mapCtx, in.EffectivePrice)
	out.ExternalPriceUri = direct.ValueOf(in.ExternalPriceURI)
	return out
}
func PriceByResource_FromProto(mapCtx *direct.MapContext, in *pb.PriceByResource) *krm.PriceByResource {
	if in == nil {
		return nil
	}
	out := &krm.PriceByResource{}
	out.ResourceType = direct.Enum_FromProto(mapCtx, in.GetResourceType())
	out.Price = Price_FromProto(mapCtx, in.GetPrice())
	out.PricePhases = direct.Slice_FromProto(mapCtx, in.PricePhases, PricePhase_FromProto)
	return out
}
func PriceByResource_ToProto(mapCtx *direct.MapContext, in *krm.PriceByResource) *pb.PriceByResource {
	if in == nil {
		return nil
	}
	out := &pb.PriceByResource{}
	out.ResourceType = direct.Enum_ToProto[pb.ResourceType](mapCtx, in.ResourceType)
	out.Price = Price_ToProto(mapCtx, in.Price)
	out.PricePhases = direct.Slice_ToProto(mapCtx, in.PricePhases, PricePhase_ToProto)
	return out
}
func PricePhase_FromProto(mapCtx *direct.MapContext, in *pb.PricePhase) *krm.PricePhase {
	if in == nil {
		return nil
	}
	out := &krm.PricePhase{}
	out.PeriodType = direct.Enum_FromProto(mapCtx, in.GetPeriodType())
	out.FirstPeriod = direct.LazyPtr(in.GetFirstPeriod())
	out.LastPeriod = direct.LazyPtr(in.GetLastPeriod())
	out.Price = Price_FromProto(mapCtx, in.GetPrice())
	out.PriceTiers = direct.Slice_FromProto(mapCtx, in.PriceTiers, PriceTier_FromProto)
	return out
}
func PricePhase_ToProto(mapCtx *direct.MapContext, in *krm.PricePhase) *pb.PricePhase {
	if in == nil {
		return nil
	}
	out := &pb.PricePhase{}
	out.PeriodType = direct.Enum_ToProto[pb.PeriodType](mapCtx, in.PeriodType)
	out.FirstPeriod = direct.ValueOf(in.FirstPeriod)
	out.LastPeriod = direct.ValueOf(in.LastPeriod)
	out.Price = Price_ToProto(mapCtx, in.Price)
	out.PriceTiers = direct.Slice_ToProto(mapCtx, in.PriceTiers, PriceTier_ToProto)
	return out
}
func PriceTier_FromProto(mapCtx *direct.MapContext, in *pb.PriceTier) *krm.PriceTier {
	if in == nil {
		return nil
	}
	out := &krm.PriceTier{}
	out.FirstResource = direct.LazyPtr(in.GetFirstResource())
	out.LastResource = direct.LazyPtr(in.GetLastResource())
	out.Price = Price_FromProto(mapCtx, in.GetPrice())
	return out
}
func PriceTier_ToProto(mapCtx *direct.MapContext, in *krm.PriceTier) *pb.PriceTier {
	if in == nil {
		return nil
	}
	out := &pb.PriceTier{}
	out.FirstResource = direct.ValueOf(in.FirstResource)
	out.LastResource = direct.ValueOf(in.LastResource)
	out.Price = Price_ToProto(mapCtx, in.Price)
	return out
}
func Product_FromProto(mapCtx *direct.MapContext, in *pb.Product) *krm.Product {
	if in == nil {
		return nil
	}
	out := &krm.Product{}
	out.Name = direct.LazyPtr(in.GetName())
	out.MarketingInfo = MarketingInfo_FromProto(mapCtx, in.GetMarketingInfo())
	return out
}
func Product_ToProto(mapCtx *direct.MapContext, in *krm.Product) *pb.Product {
	if in == nil {
		return nil
	}
	out := &pb.Product{}
	out.Name = direct.ValueOf(in.Name)
	out.MarketingInfo = MarketingInfo_ToProto(mapCtx, in.MarketingInfo)
	return out
}
func Sku_FromProto(mapCtx *direct.MapContext, in *pb.Sku) *krm.Sku {
	if in == nil {
		return nil
	}
	out := &krm.Sku{}
	out.Name = direct.LazyPtr(in.GetName())
	out.MarketingInfo = MarketingInfo_FromProto(mapCtx, in.GetMarketingInfo())
	out.Product = Product_FromProto(mapCtx, in.GetProduct())
	return out
}
func Sku_ToProto(mapCtx *direct.MapContext, in *krm.Sku) *pb.Sku {
	if in == nil {
		return nil
	}
	out := &pb.Sku{}
	out.Name = direct.ValueOf(in.Name)
	out.MarketingInfo = MarketingInfo_ToProto(mapCtx, in.MarketingInfo)
	out.Product = Product_ToProto(mapCtx, in.Product)
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
