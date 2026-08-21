// Copyright 2026 Google LLC
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//    http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package cloudtalentsolution

import (
	pb "cloud.google.com/go/talent/apiv4/talentpb"
	krm "github.com/GoogleCloudPlatform/k8s-config-connector/apis/cloudtalentsolution/v1alpha1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct"
	latlng "google.golang.org/genproto/googleapis/type/latlng"
	postaladdress "google.golang.org/genproto/googleapis/type/postaladdress"
)

func CloudTalentSolutionCompanySpec_FromProto(mapCtx *direct.MapContext, in *pb.Company) *krm.CloudTalentSolutionCompanySpec {
	if in == nil {
		return nil
	}
	out := &krm.CloudTalentSolutionCompanySpec{}
	out.DisplayName = in.GetDisplayName()
	out.ExternalID = in.GetExternalId()
	out.Size = direct.Enum_FromProto(mapCtx, in.GetSize())
	out.HeadquartersAddress = direct.LazyPtr(in.GetHeadquartersAddress())
	out.HiringAgency = direct.LazyPtr(in.GetHiringAgency())
	out.EeoText = direct.LazyPtr(in.GetEeoText())
	out.WebsiteURI = direct.LazyPtr(in.GetWebsiteUri())
	out.CareerSiteURI = direct.LazyPtr(in.GetCareerSiteUri())
	out.ImageURI = direct.LazyPtr(in.GetImageUri())
	out.KeywordSearchableJobCustomAttributes = in.KeywordSearchableJobCustomAttributes
	return out
}

func CloudTalentSolutionCompanySpec_ToProto(mapCtx *direct.MapContext, in *krm.CloudTalentSolutionCompanySpec) *pb.Company {
	if in == nil {
		return nil
	}
	out := &pb.Company{}
	out.DisplayName = in.DisplayName
	out.ExternalId = in.ExternalID
	out.Size = direct.Enum_ToProto[pb.CompanySize](mapCtx, in.Size)
	out.HeadquartersAddress = direct.ValueOf(in.HeadquartersAddress)
	out.HiringAgency = direct.ValueOf(in.HiringAgency)
	out.EeoText = direct.ValueOf(in.EeoText)
	out.WebsiteUri = direct.ValueOf(in.WebsiteURI)
	out.CareerSiteUri = direct.ValueOf(in.CareerSiteURI)
	out.ImageUri = direct.ValueOf(in.ImageURI)
	out.KeywordSearchableJobCustomAttributes = in.KeywordSearchableJobCustomAttributes
	return out
}

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
	out.AddressLines = in.GetAddressLines()
	out.Recipients = in.GetRecipients()
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

func LatLng_FromProto(mapCtx *direct.MapContext, in *latlng.LatLng) *krm.LatLng {
	if in == nil {
		return nil
	}
	out := &krm.LatLng{}
	out.Latitude = direct.LazyPtr(in.GetLatitude())
	out.Longitude = direct.LazyPtr(in.GetLongitude())
	return out
}

func LatLng_ToProto(mapCtx *direct.MapContext, in *krm.LatLng) *latlng.LatLng {
	if in == nil {
		return nil
	}
	out := &latlng.LatLng{}
	out.Latitude = direct.ValueOf(in.Latitude)
	out.Longitude = direct.ValueOf(in.Longitude)
	return out
}
