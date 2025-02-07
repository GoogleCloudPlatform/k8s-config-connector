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

package billing

import (
	refs "github.com/GoogleCloudPlatform/k8s-config-connector/apis/refs/v1beta1"
	pb "cloud.google.com/go/billing/apiv1/billingpb"
	krm "github.com/GoogleCloudPlatform/k8s-config-connector/apis/billing/v1alpha1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct"
)
func AggregationInfo_FromProto(mapCtx *direct.MapContext, in *pb.AggregationInfo) *krm.AggregationInfo {
	if in == nil {
		return nil
	}
	out := &krm.AggregationInfo{}
	out.AggregationLevel = direct.Enum_FromProto(mapCtx, in.GetAggregationLevel())
	out.AggregationInterval = direct.Enum_FromProto(mapCtx, in.GetAggregationInterval())
	out.AggregationCount = direct.LazyPtr(in.GetAggregationCount())
	return out
}
func AggregationInfo_ToProto(mapCtx *direct.MapContext, in *krm.AggregationInfo) *pb.AggregationInfo {
	if in == nil {
		return nil
	}
	out := &pb.AggregationInfo{}
	out.AggregationLevel = direct.Enum_ToProto[pb.AggregationInfo_AggregationLevel](mapCtx, in.AggregationLevel)
	out.AggregationInterval = direct.Enum_ToProto[pb.AggregationInfo_AggregationInterval](mapCtx, in.AggregationInterval)
	out.AggregationCount = direct.ValueOf(in.AggregationCount)
	return out
}
func BillingBillingAccountObservedState_FromProto(mapCtx *direct.MapContext, in *pb.BillingAccount) *krm.BillingBillingAccountObservedState {
	if in == nil {
		return nil
	}
	out := &krm.BillingBillingAccountObservedState{}
	// MISSING: Name
	// MISSING: Open
	// MISSING: DisplayName
	// MISSING: MasterBillingAccount
	// MISSING: Parent
	// MISSING: CurrencyCode
	return out
}
func BillingBillingAccountObservedState_ToProto(mapCtx *direct.MapContext, in *krm.BillingBillingAccountObservedState) *pb.BillingAccount {
	if in == nil {
		return nil
	}
	out := &pb.BillingAccount{}
	// MISSING: Name
	// MISSING: Open
	// MISSING: DisplayName
	// MISSING: MasterBillingAccount
	// MISSING: Parent
	// MISSING: CurrencyCode
	return out
}
func BillingBillingAccountSpec_FromProto(mapCtx *direct.MapContext, in *pb.BillingAccount) *krm.BillingBillingAccountSpec {
	if in == nil {
		return nil
	}
	out := &krm.BillingBillingAccountSpec{}
	// MISSING: Name
	// MISSING: Open
	// MISSING: DisplayName
	// MISSING: MasterBillingAccount
	// MISSING: Parent
	// MISSING: CurrencyCode
	return out
}
func BillingBillingAccountSpec_ToProto(mapCtx *direct.MapContext, in *krm.BillingBillingAccountSpec) *pb.BillingAccount {
	if in == nil {
		return nil
	}
	out := &pb.BillingAccount{}
	// MISSING: Name
	// MISSING: Open
	// MISSING: DisplayName
	// MISSING: MasterBillingAccount
	// MISSING: Parent
	// MISSING: CurrencyCode
	return out
}
func BillingProjectBillingInfoObservedState_FromProto(mapCtx *direct.MapContext, in *pb.ProjectBillingInfo) *krm.BillingProjectBillingInfoObservedState {
	if in == nil {
		return nil
	}
	out := &krm.BillingProjectBillingInfoObservedState{}
	// MISSING: Name
	// MISSING: ProjectID
	// MISSING: BillingAccountName
	// MISSING: BillingEnabled
	return out
}
func BillingProjectBillingInfoObservedState_ToProto(mapCtx *direct.MapContext, in *krm.BillingProjectBillingInfoObservedState) *pb.ProjectBillingInfo {
	if in == nil {
		return nil
	}
	out := &pb.ProjectBillingInfo{}
	// MISSING: Name
	// MISSING: ProjectID
	// MISSING: BillingAccountName
	// MISSING: BillingEnabled
	return out
}
func BillingProjectBillingInfoSpec_FromProto(mapCtx *direct.MapContext, in *pb.ProjectBillingInfo) *krm.BillingProjectBillingInfoSpec {
	if in == nil {
		return nil
	}
	out := &krm.BillingProjectBillingInfoSpec{}
	// MISSING: Name
	// MISSING: ProjectID
	// MISSING: BillingAccountName
	// MISSING: BillingEnabled
	return out
}
func BillingProjectBillingInfoSpec_ToProto(mapCtx *direct.MapContext, in *krm.BillingProjectBillingInfoSpec) *pb.ProjectBillingInfo {
	if in == nil {
		return nil
	}
	out := &pb.ProjectBillingInfo{}
	// MISSING: Name
	// MISSING: ProjectID
	// MISSING: BillingAccountName
	// MISSING: BillingEnabled
	return out
}
func BillingServiceObservedState_FromProto(mapCtx *direct.MapContext, in *pb.Service) *krm.BillingServiceObservedState {
	if in == nil {
		return nil
	}
	out := &krm.BillingServiceObservedState{}
	// MISSING: Name
	// MISSING: ServiceID
	// MISSING: DisplayName
	// MISSING: BusinessEntityName
	return out
}
func BillingServiceObservedState_ToProto(mapCtx *direct.MapContext, in *krm.BillingServiceObservedState) *pb.Service {
	if in == nil {
		return nil
	}
	out := &pb.Service{}
	// MISSING: Name
	// MISSING: ServiceID
	// MISSING: DisplayName
	// MISSING: BusinessEntityName
	return out
}
func BillingServiceSpec_FromProto(mapCtx *direct.MapContext, in *pb.Service) *krm.BillingServiceSpec {
	if in == nil {
		return nil
	}
	out := &krm.BillingServiceSpec{}
	// MISSING: Name
	// MISSING: ServiceID
	// MISSING: DisplayName
	// MISSING: BusinessEntityName
	return out
}
func BillingServiceSpec_ToProto(mapCtx *direct.MapContext, in *krm.BillingServiceSpec) *pb.Service {
	if in == nil {
		return nil
	}
	out := &pb.Service{}
	// MISSING: Name
	// MISSING: ServiceID
	// MISSING: DisplayName
	// MISSING: BusinessEntityName
	return out
}
func BillingSkuObservedState_FromProto(mapCtx *direct.MapContext, in *pb.Sku) *krm.BillingSkuObservedState {
	if in == nil {
		return nil
	}
	out := &krm.BillingSkuObservedState{}
	// MISSING: Name
	// MISSING: SkuID
	// MISSING: Description
	// MISSING: Category
	// MISSING: ServiceRegions
	// MISSING: PricingInfo
	// MISSING: ServiceProviderName
	// MISSING: GeoTaxonomy
	return out
}
func BillingSkuObservedState_ToProto(mapCtx *direct.MapContext, in *krm.BillingSkuObservedState) *pb.Sku {
	if in == nil {
		return nil
	}
	out := &pb.Sku{}
	// MISSING: Name
	// MISSING: SkuID
	// MISSING: Description
	// MISSING: Category
	// MISSING: ServiceRegions
	// MISSING: PricingInfo
	// MISSING: ServiceProviderName
	// MISSING: GeoTaxonomy
	return out
}
func BillingSkuSpec_FromProto(mapCtx *direct.MapContext, in *pb.Sku) *krm.BillingSkuSpec {
	if in == nil {
		return nil
	}
	out := &krm.BillingSkuSpec{}
	// MISSING: Name
	// MISSING: SkuID
	// MISSING: Description
	// MISSING: Category
	// MISSING: ServiceRegions
	// MISSING: PricingInfo
	// MISSING: ServiceProviderName
	// MISSING: GeoTaxonomy
	return out
}
func BillingSkuSpec_ToProto(mapCtx *direct.MapContext, in *krm.BillingSkuSpec) *pb.Sku {
	if in == nil {
		return nil
	}
	out := &pb.Sku{}
	// MISSING: Name
	// MISSING: SkuID
	// MISSING: Description
	// MISSING: Category
	// MISSING: ServiceRegions
	// MISSING: PricingInfo
	// MISSING: ServiceProviderName
	// MISSING: GeoTaxonomy
	return out
}
func Category_FromProto(mapCtx *direct.MapContext, in *pb.Category) *krm.Category {
	if in == nil {
		return nil
	}
	out := &krm.Category{}
	out.ServiceDisplayName = direct.LazyPtr(in.GetServiceDisplayName())
	out.ResourceFamily = direct.LazyPtr(in.GetResourceFamily())
	out.ResourceGroup = direct.LazyPtr(in.GetResourceGroup())
	out.UsageType = direct.LazyPtr(in.GetUsageType())
	return out
}
func Category_ToProto(mapCtx *direct.MapContext, in *krm.Category) *pb.Category {
	if in == nil {
		return nil
	}
	out := &pb.Category{}
	out.ServiceDisplayName = direct.ValueOf(in.ServiceDisplayName)
	out.ResourceFamily = direct.ValueOf(in.ResourceFamily)
	out.ResourceGroup = direct.ValueOf(in.ResourceGroup)
	out.UsageType = direct.ValueOf(in.UsageType)
	return out
}
func GeoTaxonomy_FromProto(mapCtx *direct.MapContext, in *pb.GeoTaxonomy) *krm.GeoTaxonomy {
	if in == nil {
		return nil
	}
	out := &krm.GeoTaxonomy{}
	out.Type = direct.Enum_FromProto(mapCtx, in.GetType())
	out.Regions = in.Regions
	return out
}
func GeoTaxonomy_ToProto(mapCtx *direct.MapContext, in *krm.GeoTaxonomy) *pb.GeoTaxonomy {
	if in == nil {
		return nil
	}
	out := &pb.GeoTaxonomy{}
	out.Type = direct.Enum_ToProto[pb.GeoTaxonomy_Type](mapCtx, in.Type)
	out.Regions = in.Regions
	return out
}
func PricingExpression_FromProto(mapCtx *direct.MapContext, in *pb.PricingExpression) *krm.PricingExpression {
	if in == nil {
		return nil
	}
	out := &krm.PricingExpression{}
	out.UsageUnit = direct.LazyPtr(in.GetUsageUnit())
	out.DisplayQuantity = direct.LazyPtr(in.GetDisplayQuantity())
	out.TieredRates = direct.Slice_FromProto(mapCtx, in.TieredRates, PricingExpression_TierRate_FromProto)
	out.UsageUnitDescription = direct.LazyPtr(in.GetUsageUnitDescription())
	out.BaseUnit = direct.LazyPtr(in.GetBaseUnit())
	out.BaseUnitDescription = direct.LazyPtr(in.GetBaseUnitDescription())
	out.BaseUnitConversionFactor = direct.LazyPtr(in.GetBaseUnitConversionFactor())
	return out
}
func PricingExpression_ToProto(mapCtx *direct.MapContext, in *krm.PricingExpression) *pb.PricingExpression {
	if in == nil {
		return nil
	}
	out := &pb.PricingExpression{}
	out.UsageUnit = direct.ValueOf(in.UsageUnit)
	out.DisplayQuantity = direct.ValueOf(in.DisplayQuantity)
	out.TieredRates = direct.Slice_ToProto(mapCtx, in.TieredRates, PricingExpression_TierRate_ToProto)
	out.UsageUnitDescription = direct.ValueOf(in.UsageUnitDescription)
	out.BaseUnit = direct.ValueOf(in.BaseUnit)
	out.BaseUnitDescription = direct.ValueOf(in.BaseUnitDescription)
	out.BaseUnitConversionFactor = direct.ValueOf(in.BaseUnitConversionFactor)
	return out
}
func PricingExpression_TierRate_FromProto(mapCtx *direct.MapContext, in *pb.PricingExpression_TierRate) *krm.PricingExpression_TierRate {
	if in == nil {
		return nil
	}
	out := &krm.PricingExpression_TierRate{}
	out.StartUsageAmount = direct.LazyPtr(in.GetStartUsageAmount())
	out.UnitPrice = Money_FromProto(mapCtx, in.GetUnitPrice())
	return out
}
func PricingExpression_TierRate_ToProto(mapCtx *direct.MapContext, in *krm.PricingExpression_TierRate) *pb.PricingExpression_TierRate {
	if in == nil {
		return nil
	}
	out := &pb.PricingExpression_TierRate{}
	out.StartUsageAmount = direct.ValueOf(in.StartUsageAmount)
	out.UnitPrice = Money_ToProto(mapCtx, in.UnitPrice)
	return out
}
func PricingInfo_FromProto(mapCtx *direct.MapContext, in *pb.PricingInfo) *krm.PricingInfo {
	if in == nil {
		return nil
	}
	out := &krm.PricingInfo{}
	out.EffectiveTime = direct.StringTimestamp_FromProto(mapCtx, in.GetEffectiveTime())
	out.Summary = direct.LazyPtr(in.GetSummary())
	out.PricingExpression = PricingExpression_FromProto(mapCtx, in.GetPricingExpression())
	out.AggregationInfo = AggregationInfo_FromProto(mapCtx, in.GetAggregationInfo())
	out.CurrencyConversionRate = direct.LazyPtr(in.GetCurrencyConversionRate())
	return out
}
func PricingInfo_ToProto(mapCtx *direct.MapContext, in *krm.PricingInfo) *pb.PricingInfo {
	if in == nil {
		return nil
	}
	out := &pb.PricingInfo{}
	out.EffectiveTime = direct.StringTimestamp_ToProto(mapCtx, in.EffectiveTime)
	out.Summary = direct.ValueOf(in.Summary)
	out.PricingExpression = PricingExpression_ToProto(mapCtx, in.PricingExpression)
	out.AggregationInfo = AggregationInfo_ToProto(mapCtx, in.AggregationInfo)
	out.CurrencyConversionRate = direct.ValueOf(in.CurrencyConversionRate)
	return out
}
func Sku_FromProto(mapCtx *direct.MapContext, in *pb.Sku) *krm.Sku {
	if in == nil {
		return nil
	}
	out := &krm.Sku{}
	out.Name = direct.LazyPtr(in.GetName())
	out.SkuID = direct.LazyPtr(in.GetSkuId())
	out.Description = direct.LazyPtr(in.GetDescription())
	out.Category = Category_FromProto(mapCtx, in.GetCategory())
	out.ServiceRegions = in.ServiceRegions
	out.PricingInfo = direct.Slice_FromProto(mapCtx, in.PricingInfo, PricingInfo_FromProto)
	out.ServiceProviderName = direct.LazyPtr(in.GetServiceProviderName())
	out.GeoTaxonomy = GeoTaxonomy_FromProto(mapCtx, in.GetGeoTaxonomy())
	return out
}
func Sku_ToProto(mapCtx *direct.MapContext, in *krm.Sku) *pb.Sku {
	if in == nil {
		return nil
	}
	out := &pb.Sku{}
	out.Name = direct.ValueOf(in.Name)
	out.SkuId = direct.ValueOf(in.SkuID)
	out.Description = direct.ValueOf(in.Description)
	out.Category = Category_ToProto(mapCtx, in.Category)
	out.ServiceRegions = in.ServiceRegions
	out.PricingInfo = direct.Slice_ToProto(mapCtx, in.PricingInfo, PricingInfo_ToProto)
	out.ServiceProviderName = direct.ValueOf(in.ServiceProviderName)
	out.GeoTaxonomy = GeoTaxonomy_ToProto(mapCtx, in.GeoTaxonomy)
	return out
}
