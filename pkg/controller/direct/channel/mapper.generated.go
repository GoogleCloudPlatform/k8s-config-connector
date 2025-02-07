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
	pb "cloud.google.com/go/channel/apiv1/channelpb"
	krm "github.com/GoogleCloudPlatform/k8s-config-connector/apis/channel/v1alpha1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct"
	refs "github.com/GoogleCloudPlatform/k8s-config-connector/apis/refs/v1beta1"
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
func ReportJob_FromProto(mapCtx *direct.MapContext, in *pb.ReportJob) *krm.ReportJob {
	if in == nil {
		return nil
	}
	out := &krm.ReportJob{}
	out.Name = direct.LazyPtr(in.GetName())
	out.ReportStatus = ReportStatus_FromProto(mapCtx, in.GetReportStatus())
	return out
}
func ReportJob_ToProto(mapCtx *direct.MapContext, in *krm.ReportJob) *pb.ReportJob {
	if in == nil {
		return nil
	}
	out := &pb.ReportJob{}
	out.Name = direct.ValueOf(in.Name)
	out.ReportStatus = ReportStatus_ToProto(mapCtx, in.ReportStatus)
	return out
}
func ReportStatus_FromProto(mapCtx *direct.MapContext, in *pb.ReportStatus) *krm.ReportStatus {
	if in == nil {
		return nil
	}
	out := &krm.ReportStatus{}
	out.State = direct.Enum_FromProto(mapCtx, in.GetState())
	out.StartTime = direct.StringTimestamp_FromProto(mapCtx, in.GetStartTime())
	out.EndTime = direct.StringTimestamp_FromProto(mapCtx, in.GetEndTime())
	return out
}
func ReportStatus_ToProto(mapCtx *direct.MapContext, in *krm.ReportStatus) *pb.ReportStatus {
	if in == nil {
		return nil
	}
	out := &pb.ReportStatus{}
	out.State = direct.Enum_ToProto[pb.ReportStatus_State](mapCtx, in.State)
	out.StartTime = direct.StringTimestamp_ToProto(mapCtx, in.StartTime)
	out.EndTime = direct.StringTimestamp_ToProto(mapCtx, in.EndTime)
	return out
}
