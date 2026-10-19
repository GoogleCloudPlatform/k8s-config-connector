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

package cloudchannel

import (
	krm "github.com/GoogleCloudPlatform/k8s-config-connector/apis/cloudchannel/v1alpha1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct"
	"google.golang.org/genproto/googleapis/type/date"
	"google.golang.org/genproto/googleapis/type/decimal"
	"google.golang.org/genproto/googleapis/type/postaladdress"
)

func PostalAddress_FromProto(mapCtx *direct.MapContext, in *postaladdress.PostalAddress) *krm.PostalAddress {
	if in == nil {
		return nil
	}
	out := &krm.PostalAddress{}
	out.Revision = direct.LazyPtr(in.GetRevision())
	out.RegionCode = direct.LazyPtr(in.GetRegionCode())
	out.LanguageCode = direct.LazyPtr(in.GetLanguageCode())
	out.PostalCode = direct.LazyPtr(in.GetPostalCode())
	out.SortingCode = direct.LazyPtr(in.GetSortingCode())
	out.AdministrativeArea = direct.LazyPtr(in.GetAdministrativeArea())
	out.Locality = direct.LazyPtr(in.GetLocality())
	out.Sublocality = direct.LazyPtr(in.GetSublocality())
	out.AddressLines = in.AddressLines
	out.Recipients = in.Recipients
	out.Organization = direct.LazyPtr(in.GetOrganization())
	return out
}

func PostalAddress_ToProto(mapCtx *direct.MapContext, in *krm.PostalAddress) *postaladdress.PostalAddress {
	if in == nil {
		return nil
	}
	out := &postaladdress.PostalAddress{}
	out.Revision = direct.ValueOf(in.Revision)
	out.RegionCode = direct.ValueOf(in.RegionCode)
	out.LanguageCode = direct.ValueOf(in.LanguageCode)
	out.PostalCode = direct.ValueOf(in.PostalCode)
	out.SortingCode = direct.ValueOf(in.SortingCode)
	out.AdministrativeArea = direct.ValueOf(in.AdministrativeArea)
	out.Locality = direct.ValueOf(in.Locality)
	out.Sublocality = direct.ValueOf(in.Sublocality)
	out.AddressLines = in.AddressLines
	out.Recipients = in.Recipients
	out.Organization = direct.ValueOf(in.Organization)
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

func Date_FromProto(mapCtx *direct.MapContext, in *date.Date) *krm.Date {
	if in == nil {
		return nil
	}
	out := &krm.Date{}
	out.Year = direct.LazyPtr(in.GetYear())
	out.Month = direct.LazyPtr(in.GetMonth())
	out.Day = direct.LazyPtr(in.GetDay())
	return out
}

func Date_ToProto(mapCtx *direct.MapContext, in *krm.Date) *date.Date {
	if in == nil {
		return nil
	}
	out := &date.Date{}
	out.Year = direct.ValueOf(in.Year)
	out.Month = direct.ValueOf(in.Month)
	out.Day = direct.ValueOf(in.Day)
	return out
}
