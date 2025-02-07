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
	krm "github.com/GoogleCloudPlatform/k8s-config-connector/apis/channel/v1alpha1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct"
	refs "github.com/GoogleCloudPlatform/k8s-config-connector/apis/refs/v1beta1"
	pb "cloud.google.com/go/channel/apiv1/channelpb"
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
