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

package retail

import (
	refs "github.com/GoogleCloudPlatform/k8s-config-connector/apis/refs/v1beta1"
	pb "cloud.google.com/go/retail/apiv2beta/retailpb"
	krm "github.com/GoogleCloudPlatform/k8s-config-connector/apis/retail/v1alpha1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct"
)
func Audience_FromProto(mapCtx *direct.MapContext, in *pb.Audience) *krm.Audience {
	if in == nil {
		return nil
	}
	out := &krm.Audience{}
	out.Genders = in.Genders
	out.AgeGroups = in.AgeGroups
	return out
}
func Audience_ToProto(mapCtx *direct.MapContext, in *krm.Audience) *pb.Audience {
	if in == nil {
		return nil
	}
	out := &pb.Audience{}
	out.Genders = in.Genders
	out.AgeGroups = in.AgeGroups
	return out
}
func ColorInfo_FromProto(mapCtx *direct.MapContext, in *pb.ColorInfo) *krm.ColorInfo {
	if in == nil {
		return nil
	}
	out := &krm.ColorInfo{}
	out.ColorFamilies = in.ColorFamilies
	out.Colors = in.Colors
	return out
}
func ColorInfo_ToProto(mapCtx *direct.MapContext, in *krm.ColorInfo) *pb.ColorInfo {
	if in == nil {
		return nil
	}
	out := &pb.ColorInfo{}
	out.ColorFamilies = in.ColorFamilies
	out.Colors = in.Colors
	return out
}
func CustomAttribute_FromProto(mapCtx *direct.MapContext, in *pb.CustomAttribute) *krm.CustomAttribute {
	if in == nil {
		return nil
	}
	out := &krm.CustomAttribute{}
	out.Text = in.Text
	out.Numbers = in.Numbers
	out.Searchable = in.Searchable
	out.Indexable = in.Indexable
	return out
}
func CustomAttribute_ToProto(mapCtx *direct.MapContext, in *krm.CustomAttribute) *pb.CustomAttribute {
	if in == nil {
		return nil
	}
	out := &pb.CustomAttribute{}
	out.Text = in.Text
	out.Numbers = in.Numbers
	out.Searchable = in.Searchable
	out.Indexable = in.Indexable
	return out
}
func FulfillmentInfo_FromProto(mapCtx *direct.MapContext, in *pb.FulfillmentInfo) *krm.FulfillmentInfo {
	if in == nil {
		return nil
	}
	out := &krm.FulfillmentInfo{}
	out.Type = direct.LazyPtr(in.GetType())
	out.PlaceIds = in.PlaceIds
	return out
}
func FulfillmentInfo_ToProto(mapCtx *direct.MapContext, in *krm.FulfillmentInfo) *pb.FulfillmentInfo {
	if in == nil {
		return nil
	}
	out := &pb.FulfillmentInfo{}
	out.Type = direct.ValueOf(in.Type)
	out.PlaceIds = in.PlaceIds
	return out
}
func Image_FromProto(mapCtx *direct.MapContext, in *pb.Image) *krm.Image {
	if in == nil {
		return nil
	}
	out := &krm.Image{}
	out.URI = direct.LazyPtr(in.GetUri())
	out.Height = direct.LazyPtr(in.GetHeight())
	out.Width = direct.LazyPtr(in.GetWidth())
	return out
}
func Image_ToProto(mapCtx *direct.MapContext, in *krm.Image) *pb.Image {
	if in == nil {
		return nil
	}
	out := &pb.Image{}
	out.Uri = direct.ValueOf(in.URI)
	out.Height = direct.ValueOf(in.Height)
	out.Width = direct.ValueOf(in.Width)
	return out
}
func Interval_FromProto(mapCtx *direct.MapContext, in *pb.Interval) *krm.Interval {
	if in == nil {
		return nil
	}
	out := &krm.Interval{}
	out.Minimum = direct.LazyPtr(in.GetMinimum())
	out.ExclusiveMinimum = direct.LazyPtr(in.GetExclusiveMinimum())
	out.Maximum = direct.LazyPtr(in.GetMaximum())
	out.ExclusiveMaximum = direct.LazyPtr(in.GetExclusiveMaximum())
	return out
}
func Interval_ToProto(mapCtx *direct.MapContext, in *krm.Interval) *pb.Interval {
	if in == nil {
		return nil
	}
	out := &pb.Interval{}
	if oneof := Interval_Minimum_ToProto(mapCtx, in.Minimum); oneof != nil {
		out.Min = oneof
	}
	if oneof := Interval_ExclusiveMinimum_ToProto(mapCtx, in.ExclusiveMinimum); oneof != nil {
		out.Min = oneof
	}
	if oneof := Interval_Maximum_ToProto(mapCtx, in.Maximum); oneof != nil {
		out.Max = oneof
	}
	if oneof := Interval_ExclusiveMaximum_ToProto(mapCtx, in.ExclusiveMaximum); oneof != nil {
		out.Max = oneof
	}
	return out
}
func LocalInventory_FromProto(mapCtx *direct.MapContext, in *pb.LocalInventory) *krm.LocalInventory {
	if in == nil {
		return nil
	}
	out := &krm.LocalInventory{}
	out.PlaceID = direct.LazyPtr(in.GetPlaceId())
	out.PriceInfo = PriceInfo_FromProto(mapCtx, in.GetPriceInfo())
	// MISSING: Attributes
	out.FulfillmentTypes = in.FulfillmentTypes
	return out
}
func LocalInventory_ToProto(mapCtx *direct.MapContext, in *krm.LocalInventory) *pb.LocalInventory {
	if in == nil {
		return nil
	}
	out := &pb.LocalInventory{}
	out.PlaceId = direct.ValueOf(in.PlaceID)
	out.PriceInfo = PriceInfo_ToProto(mapCtx, in.PriceInfo)
	// MISSING: Attributes
	out.FulfillmentTypes = in.FulfillmentTypes
	return out
}
func PriceInfo_FromProto(mapCtx *direct.MapContext, in *pb.PriceInfo) *krm.PriceInfo {
	if in == nil {
		return nil
	}
	out := &krm.PriceInfo{}
	out.CurrencyCode = direct.LazyPtr(in.GetCurrencyCode())
	out.Price = direct.LazyPtr(in.GetPrice())
	out.OriginalPrice = direct.LazyPtr(in.GetOriginalPrice())
	out.Cost = direct.LazyPtr(in.GetCost())
	out.PriceEffectiveTime = direct.StringTimestamp_FromProto(mapCtx, in.GetPriceEffectiveTime())
	out.PriceExpireTime = direct.StringTimestamp_FromProto(mapCtx, in.GetPriceExpireTime())
	// MISSING: PriceRange
	return out
}
func PriceInfo_ToProto(mapCtx *direct.MapContext, in *krm.PriceInfo) *pb.PriceInfo {
	if in == nil {
		return nil
	}
	out := &pb.PriceInfo{}
	out.CurrencyCode = direct.ValueOf(in.CurrencyCode)
	out.Price = direct.ValueOf(in.Price)
	out.OriginalPrice = direct.ValueOf(in.OriginalPrice)
	out.Cost = direct.ValueOf(in.Cost)
	out.PriceEffectiveTime = direct.StringTimestamp_ToProto(mapCtx, in.PriceEffectiveTime)
	out.PriceExpireTime = direct.StringTimestamp_ToProto(mapCtx, in.PriceExpireTime)
	// MISSING: PriceRange
	return out
}
func PriceInfoObservedState_FromProto(mapCtx *direct.MapContext, in *pb.PriceInfo) *krm.PriceInfoObservedState {
	if in == nil {
		return nil
	}
	out := &krm.PriceInfoObservedState{}
	// MISSING: CurrencyCode
	// MISSING: Price
	// MISSING: OriginalPrice
	// MISSING: Cost
	// MISSING: PriceEffectiveTime
	// MISSING: PriceExpireTime
	out.PriceRange = PriceInfo_PriceRange_FromProto(mapCtx, in.GetPriceRange())
	return out
}
func PriceInfoObservedState_ToProto(mapCtx *direct.MapContext, in *krm.PriceInfoObservedState) *pb.PriceInfo {
	if in == nil {
		return nil
	}
	out := &pb.PriceInfo{}
	// MISSING: CurrencyCode
	// MISSING: Price
	// MISSING: OriginalPrice
	// MISSING: Cost
	// MISSING: PriceEffectiveTime
	// MISSING: PriceExpireTime
	out.PriceRange = PriceInfo_PriceRange_ToProto(mapCtx, in.PriceRange)
	return out
}
func PriceInfo_PriceRange_FromProto(mapCtx *direct.MapContext, in *pb.PriceInfo_PriceRange) *krm.PriceInfo_PriceRange {
	if in == nil {
		return nil
	}
	out := &krm.PriceInfo_PriceRange{}
	out.Price = Interval_FromProto(mapCtx, in.GetPrice())
	out.OriginalPrice = Interval_FromProto(mapCtx, in.GetOriginalPrice())
	return out
}
func PriceInfo_PriceRange_ToProto(mapCtx *direct.MapContext, in *krm.PriceInfo_PriceRange) *pb.PriceInfo_PriceRange {
	if in == nil {
		return nil
	}
	out := &pb.PriceInfo_PriceRange{}
	out.Price = Interval_ToProto(mapCtx, in.Price)
	out.OriginalPrice = Interval_ToProto(mapCtx, in.OriginalPrice)
	return out
}
func Product_FromProto(mapCtx *direct.MapContext, in *pb.Product) *krm.Product {
	if in == nil {
		return nil
	}
	out := &krm.Product{}
	out.ExpireTime = direct.StringTimestamp_FromProto(mapCtx, in.GetExpireTime())
	out.Ttl = direct.StringDuration_FromProto(mapCtx, in.GetTtl())
	out.Name = direct.LazyPtr(in.GetName())
	out.ID = direct.LazyPtr(in.GetId())
	out.Type = direct.Enum_FromProto(mapCtx, in.GetType())
	out.PrimaryProductID = direct.LazyPtr(in.GetPrimaryProductId())
	out.CollectionMemberIds = in.CollectionMemberIds
	out.Gtin = direct.LazyPtr(in.GetGtin())
	out.Categories = in.Categories
	out.Title = direct.LazyPtr(in.GetTitle())
	out.Brands = in.Brands
	out.Description = direct.LazyPtr(in.GetDescription())
	out.LanguageCode = direct.LazyPtr(in.GetLanguageCode())
	// MISSING: Attributes
	out.Tags = in.Tags
	out.PriceInfo = PriceInfo_FromProto(mapCtx, in.GetPriceInfo())
	out.Rating = Rating_FromProto(mapCtx, in.GetRating())
	out.AvailableTime = direct.StringTimestamp_FromProto(mapCtx, in.GetAvailableTime())
	out.Availability = direct.Enum_FromProto(mapCtx, in.GetAvailability())
	out.AvailableQuantity = Int32Value_FromProto(mapCtx, in.GetAvailableQuantity())
	out.FulfillmentInfo = direct.Slice_FromProto(mapCtx, in.FulfillmentInfo, FulfillmentInfo_FromProto)
	out.URI = direct.LazyPtr(in.GetUri())
	out.Images = direct.Slice_FromProto(mapCtx, in.Images, Image_FromProto)
	out.Audience = Audience_FromProto(mapCtx, in.GetAudience())
	out.ColorInfo = ColorInfo_FromProto(mapCtx, in.GetColorInfo())
	out.Sizes = in.Sizes
	out.Materials = in.Materials
	out.Patterns = in.Patterns
	out.Conditions = in.Conditions
	out.Promotions = direct.Slice_FromProto(mapCtx, in.Promotions, Promotion_FromProto)
	out.PublishTime = direct.StringTimestamp_FromProto(mapCtx, in.GetPublishTime())
	out.RetrievableFields = FieldMask_FromProto(mapCtx, in.GetRetrievableFields())
	// MISSING: Variants
	// MISSING: LocalInventories
	return out
}
func Product_ToProto(mapCtx *direct.MapContext, in *krm.Product) *pb.Product {
	if in == nil {
		return nil
	}
	out := &pb.Product{}
	if oneof := direct.StringTimestamp_ToProto(mapCtx, in.ExpireTime); oneof != nil {
		out.Expiration = &pb.Product_ExpireTime{ExpireTime: oneof}
	}
	if oneof := direct.StringDuration_ToProto(mapCtx, in.Ttl); oneof != nil {
		out.Expiration = &pb.Product_Ttl{Ttl: oneof}
	}
	out.Name = direct.ValueOf(in.Name)
	out.Id = direct.ValueOf(in.ID)
	out.Type = direct.Enum_ToProto[pb.Product_Type](mapCtx, in.Type)
	out.PrimaryProductId = direct.ValueOf(in.PrimaryProductID)
	out.CollectionMemberIds = in.CollectionMemberIds
	out.Gtin = direct.ValueOf(in.Gtin)
	out.Categories = in.Categories
	out.Title = direct.ValueOf(in.Title)
	out.Brands = in.Brands
	out.Description = direct.ValueOf(in.Description)
	out.LanguageCode = direct.ValueOf(in.LanguageCode)
	// MISSING: Attributes
	out.Tags = in.Tags
	out.PriceInfo = PriceInfo_ToProto(mapCtx, in.PriceInfo)
	out.Rating = Rating_ToProto(mapCtx, in.Rating)
	out.AvailableTime = direct.StringTimestamp_ToProto(mapCtx, in.AvailableTime)
	out.Availability = direct.Enum_ToProto[pb.Product_Availability](mapCtx, in.Availability)
	out.AvailableQuantity = Int32Value_ToProto(mapCtx, in.AvailableQuantity)
	out.FulfillmentInfo = direct.Slice_ToProto(mapCtx, in.FulfillmentInfo, FulfillmentInfo_ToProto)
	out.Uri = direct.ValueOf(in.URI)
	out.Images = direct.Slice_ToProto(mapCtx, in.Images, Image_ToProto)
	out.Audience = Audience_ToProto(mapCtx, in.Audience)
	out.ColorInfo = ColorInfo_ToProto(mapCtx, in.ColorInfo)
	out.Sizes = in.Sizes
	out.Materials = in.Materials
	out.Patterns = in.Patterns
	out.Conditions = in.Conditions
	out.Promotions = direct.Slice_ToProto(mapCtx, in.Promotions, Promotion_ToProto)
	out.PublishTime = direct.StringTimestamp_ToProto(mapCtx, in.PublishTime)
	out.RetrievableFields = FieldMask_ToProto(mapCtx, in.RetrievableFields)
	// MISSING: Variants
	// MISSING: LocalInventories
	return out
}
func ProductObservedState_FromProto(mapCtx *direct.MapContext, in *pb.Product) *krm.ProductObservedState {
	if in == nil {
		return nil
	}
	out := &krm.ProductObservedState{}
	// MISSING: ExpireTime
	// MISSING: Ttl
	// MISSING: Name
	// MISSING: ID
	// MISSING: Type
	// MISSING: PrimaryProductID
	// MISSING: CollectionMemberIds
	// MISSING: Gtin
	// MISSING: Categories
	// MISSING: Title
	// MISSING: Brands
	// MISSING: Description
	// MISSING: LanguageCode
	// MISSING: Attributes
	// MISSING: Tags
	out.PriceInfo = PriceInfoObservedState_FromProto(mapCtx, in.GetPriceInfo())
	// MISSING: Rating
	// MISSING: AvailableTime
	// MISSING: Availability
	// MISSING: AvailableQuantity
	// MISSING: FulfillmentInfo
	// MISSING: URI
	// MISSING: Images
	// MISSING: Audience
	// MISSING: ColorInfo
	// MISSING: Sizes
	// MISSING: Materials
	// MISSING: Patterns
	// MISSING: Conditions
	// MISSING: Promotions
	// MISSING: PublishTime
	// MISSING: RetrievableFields
	out.Variants = direct.Slice_FromProto(mapCtx, in.Variants, Product_FromProto)
	out.LocalInventories = direct.Slice_FromProto(mapCtx, in.LocalInventories, LocalInventory_FromProto)
	return out
}
func ProductObservedState_ToProto(mapCtx *direct.MapContext, in *krm.ProductObservedState) *pb.Product {
	if in == nil {
		return nil
	}
	out := &pb.Product{}
	// MISSING: ExpireTime
	// MISSING: Ttl
	// MISSING: Name
	// MISSING: ID
	// MISSING: Type
	// MISSING: PrimaryProductID
	// MISSING: CollectionMemberIds
	// MISSING: Gtin
	// MISSING: Categories
	// MISSING: Title
	// MISSING: Brands
	// MISSING: Description
	// MISSING: LanguageCode
	// MISSING: Attributes
	// MISSING: Tags
	out.PriceInfo = PriceInfoObservedState_ToProto(mapCtx, in.PriceInfo)
	// MISSING: Rating
	// MISSING: AvailableTime
	// MISSING: Availability
	// MISSING: AvailableQuantity
	// MISSING: FulfillmentInfo
	// MISSING: URI
	// MISSING: Images
	// MISSING: Audience
	// MISSING: ColorInfo
	// MISSING: Sizes
	// MISSING: Materials
	// MISSING: Patterns
	// MISSING: Conditions
	// MISSING: Promotions
	// MISSING: PublishTime
	// MISSING: RetrievableFields
	out.Variants = direct.Slice_ToProto(mapCtx, in.Variants, Product_ToProto)
	out.LocalInventories = direct.Slice_ToProto(mapCtx, in.LocalInventories, LocalInventory_ToProto)
	return out
}
func Promotion_FromProto(mapCtx *direct.MapContext, in *pb.Promotion) *krm.Promotion {
	if in == nil {
		return nil
	}
	out := &krm.Promotion{}
	out.PromotionID = direct.LazyPtr(in.GetPromotionId())
	return out
}
func Promotion_ToProto(mapCtx *direct.MapContext, in *krm.Promotion) *pb.Promotion {
	if in == nil {
		return nil
	}
	out := &pb.Promotion{}
	out.PromotionId = direct.ValueOf(in.PromotionID)
	return out
}
func Rating_FromProto(mapCtx *direct.MapContext, in *pb.Rating) *krm.Rating {
	if in == nil {
		return nil
	}
	out := &krm.Rating{}
	out.RatingCount = direct.LazyPtr(in.GetRatingCount())
	out.AverageRating = direct.LazyPtr(in.GetAverageRating())
	out.RatingHistogram = in.RatingHistogram
	return out
}
func Rating_ToProto(mapCtx *direct.MapContext, in *krm.Rating) *pb.Rating {
	if in == nil {
		return nil
	}
	out := &pb.Rating{}
	out.RatingCount = direct.ValueOf(in.RatingCount)
	out.AverageRating = direct.ValueOf(in.AverageRating)
	out.RatingHistogram = in.RatingHistogram
	return out
}
