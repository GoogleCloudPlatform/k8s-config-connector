// Copyright 2026 Google LLC
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

package commerceproducer

import (
	pb "cloud.google.com/go/commerceproducer/apiv1beta/commerceproducerpb"
	krm "github.com/GoogleCloudPlatform/k8s-config-connector/apis/commerceproducer/v1alpha1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct"
	datetime "google.golang.org/genproto/googleapis/type/datetime"
	decimal "google.golang.org/genproto/googleapis/type/decimal"
	money "google.golang.org/genproto/googleapis/type/money"
)

func DateTime_FromProto(mapCtx *direct.MapContext, in *datetime.DateTime) *krm.DateTime {
	if in == nil {
		return nil
	}
	out := &krm.DateTime{}
	out.Year = direct.LazyPtr(in.GetYear())
	out.Month = direct.LazyPtr(in.GetMonth())
	out.Day = direct.LazyPtr(in.GetDay())
	out.Hours = direct.LazyPtr(in.GetHours())
	out.Minutes = direct.LazyPtr(in.GetMinutes())
	out.Seconds = direct.LazyPtr(in.GetSeconds())
	out.Nanos = direct.LazyPtr(in.GetNanos())
	if in.GetTimeZone() != nil {
		out.TimeZone = &krm.TimeZone{
			ID:      direct.LazyPtr(in.GetTimeZone().GetId()),
			Version: direct.LazyPtr(in.GetTimeZone().GetVersion()),
		}
	}
	if in.GetUtcOffset() != nil {
		out.UtcOffset = direct.StringDuration_FromProto(mapCtx, in.GetUtcOffset())
	}
	return out
}

func DateTime_ToProto(mapCtx *direct.MapContext, in *krm.DateTime) *datetime.DateTime {
	if in == nil {
		return nil
	}
	out := &datetime.DateTime{}
	out.Year = direct.ValueOf(in.Year)
	out.Month = direct.ValueOf(in.Month)
	out.Day = direct.ValueOf(in.Day)
	out.Hours = direct.ValueOf(in.Hours)
	out.Minutes = direct.ValueOf(in.Minutes)
	out.Seconds = direct.ValueOf(in.Seconds)
	out.Nanos = direct.ValueOf(in.Nanos)
	if in.TimeZone != nil {
		out.TimeOffset = &datetime.DateTime_TimeZone{
			TimeZone: &datetime.TimeZone{
				Id:      direct.ValueOf(in.TimeZone.ID),
				Version: direct.ValueOf(in.TimeZone.Version),
			},
		}
	} else if in.UtcOffset != nil {
		if dur := direct.StringDuration_ToProto(mapCtx, in.UtcOffset); dur != nil {
			out.TimeOffset = &datetime.DateTime_UtcOffset{
				UtcOffset: dur,
			}
		}
	}
	return out
}

func Decimal_FromProto(mapCtx *direct.MapContext, in *decimal.Decimal) *krm.Decimal {
	if in == nil {
		return nil
	}
	out := &krm.Decimal{}
	out.Value = direct.LazyPtr(in.GetValue())
	return out
}

func Decimal_ToProto(mapCtx *direct.MapContext, in *krm.Decimal) *decimal.Decimal {
	if in == nil {
		return nil
	}
	out := &decimal.Decimal{}
	out.Value = direct.ValueOf(in.Value)
	return out
}

func Money_FromProto(mapCtx *direct.MapContext, in *money.Money) *krm.Money {
	if in == nil {
		return nil
	}
	out := &krm.Money{}
	out.CurrencyCode = direct.LazyPtr(in.GetCurrencyCode())
	out.Units = direct.LazyPtr(in.GetUnits())
	out.Nanos = direct.LazyPtr(in.GetNanos())
	return out
}

func Money_ToProto(mapCtx *direct.MapContext, in *krm.Money) *money.Money {
	if in == nil {
		return nil
	}
	out := &money.Money{}
	out.CurrencyCode = direct.ValueOf(in.CurrencyCode)
	out.Units = direct.ValueOf(in.Units)
	out.Nanos = direct.ValueOf(in.Nanos)
	return out
}

func PrivateOfferSingleProductOffer_FromProto(mapCtx *direct.MapContext, in *pb.PrivateOffer_SingleProductOffer) *krm.PrivateOfferSingleProductOffer {
	if in == nil {
		return nil
	}
	out := &krm.PrivateOfferSingleProductOffer{}
	if in.GetAmendedPrivateOffer() != "" {
		out.AmendedPrivateOfferRef = &krm.CommerceProducerPrivateOfferRef{External: in.GetAmendedPrivateOffer()}
	}
	if in.GetAmendedStandardOffer() != "" {
		out.AmendedStandardOfferRef = &krm.CommerceProducerStandardOfferRef{External: in.GetAmendedStandardOffer()}
	}
	out.StandardIntervalPrice = PrivateOfferSingleProductOfferStandardIntervalPrice_FromProto(mapCtx, in.GetStandardIntervalPrice())
	out.CustomIntervalPrice = PrivateOfferSingleProductOfferCustomIntervalPrice_FromProto(mapCtx, in.GetCustomIntervalPrice())
	if in.GetBaseStandardOffer() != "" {
		out.BaseStandardOfferRef = &krm.CommerceProducerStandardOfferRef{External: in.GetBaseStandardOffer()}
	}
	out.Features = direct.Slice_FromProto(mapCtx, in.Features, PrivateOfferSingleProductOfferFeature_FromProto)
	out.AdditionalContractValue = PrivateOfferSingleProductOfferAdditionalContractValue_FromProto(mapCtx, in.GetAdditionalContractValue())
	return out
}

func PrivateOfferSingleProductOffer_ToProto(mapCtx *direct.MapContext, in *krm.PrivateOfferSingleProductOffer) *pb.PrivateOffer_SingleProductOffer {
	if in == nil {
		return nil
	}
	out := &pb.PrivateOffer_SingleProductOffer{}
	if in.AmendedPrivateOfferRef != nil {
		out.AmendedOffer = &pb.PrivateOffer_SingleProductOffer_AmendedPrivateOffer{
			AmendedPrivateOffer: in.AmendedPrivateOfferRef.External,
		}
	}
	if in.AmendedStandardOfferRef != nil {
		out.AmendedOffer = &pb.PrivateOffer_SingleProductOffer_AmendedStandardOffer{
			AmendedStandardOffer: in.AmendedStandardOfferRef.External,
		}
	}
	if oneof := PrivateOfferSingleProductOfferStandardIntervalPrice_ToProto(mapCtx, in.StandardIntervalPrice); oneof != nil {
		out.Price = &pb.PrivateOffer_SingleProductOffer_StandardIntervalPrice_{StandardIntervalPrice: oneof}
	}
	if oneof := PrivateOfferSingleProductOfferCustomIntervalPrice_ToProto(mapCtx, in.CustomIntervalPrice); oneof != nil {
		out.Price = &pb.PrivateOffer_SingleProductOffer_CustomIntervalPrice_{CustomIntervalPrice: oneof}
	}
	if in.BaseStandardOfferRef != nil {
		out.BaseStandardOffer = in.BaseStandardOfferRef.External
	}
	out.Features = direct.Slice_ToProto(mapCtx, in.Features, PrivateOfferSingleProductOfferFeature_ToProto)
	out.AdditionalContractValue = PrivateOfferSingleProductOfferAdditionalContractValue_ToProto(mapCtx, in.AdditionalContractValue)
	return out
}

func PrivateOfferSingleProductOfferPriceModelSkuDiscount_ToProto(mapCtx *direct.MapContext, in *krm.PrivateOfferSingleProductOfferPriceModelSkuDiscount) *pb.PrivateOffer_SingleProductOffer_PriceModel_SkuDiscount {
	if in == nil {
		return nil
	}
	out := &pb.PrivateOffer_SingleProductOffer_PriceModel_SkuDiscount{}
	if in.SkuRef != nil {
		out.Target = &pb.PrivateOffer_SingleProductOffer_PriceModel_SkuDiscount_Sku{
			Sku: in.SkuRef.External,
		}
	}
	if oneof := PrivateOfferSingleProductOfferPriceModelSkuDiscount_CloudBillingSku_ToProto(mapCtx, in.CloudBillingSku); oneof != nil {
		out.Target = oneof
	}
	if oneof := Decimal_ToProto(mapCtx, in.DiscountPercent); oneof != nil {
		out.Discount = &pb.PrivateOffer_SingleProductOffer_PriceModel_SkuDiscount_DiscountPercent{DiscountPercent: oneof}
	}
	return out
}
