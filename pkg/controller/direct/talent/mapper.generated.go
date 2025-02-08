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

package talent

import (
	krm "github.com/GoogleCloudPlatform/k8s-config-connector/apis/talent/v1alpha1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct"
	refs "github.com/GoogleCloudPlatform/k8s-config-connector/apis/refs/v1beta1"
	pb "cloud.google.com/go/talent/apiv4/talentpb"
)
func Company_FromProto(mapCtx *direct.MapContext, in *pb.Company) *krm.Company {
	if in == nil {
		return nil
	}
	out := &krm.Company{}
	out.Name = direct.LazyPtr(in.GetName())
	out.DisplayName = direct.LazyPtr(in.GetDisplayName())
	out.ExternalID = direct.LazyPtr(in.GetExternalId())
	out.Size = direct.Enum_FromProto(mapCtx, in.GetSize())
	out.HeadquartersAddress = direct.LazyPtr(in.GetHeadquartersAddress())
	out.HiringAgency = direct.LazyPtr(in.GetHiringAgency())
	out.EeoText = direct.LazyPtr(in.GetEeoText())
	out.WebsiteURI = direct.LazyPtr(in.GetWebsiteUri())
	out.CareerSiteURI = direct.LazyPtr(in.GetCareerSiteUri())
	out.ImageURI = direct.LazyPtr(in.GetImageUri())
	out.KeywordSearchableJobCustomAttributes = in.KeywordSearchableJobCustomAttributes
	// MISSING: DerivedInfo
	// MISSING: Suspended
	return out
}
func Company_ToProto(mapCtx *direct.MapContext, in *krm.Company) *pb.Company {
	if in == nil {
		return nil
	}
	out := &pb.Company{}
	out.Name = direct.ValueOf(in.Name)
	out.DisplayName = direct.ValueOf(in.DisplayName)
	out.ExternalId = direct.ValueOf(in.ExternalID)
	out.Size = direct.Enum_ToProto[pb.CompanySize](mapCtx, in.Size)
	out.HeadquartersAddress = direct.ValueOf(in.HeadquartersAddress)
	out.HiringAgency = direct.ValueOf(in.HiringAgency)
	out.EeoText = direct.ValueOf(in.EeoText)
	out.WebsiteUri = direct.ValueOf(in.WebsiteURI)
	out.CareerSiteUri = direct.ValueOf(in.CareerSiteURI)
	out.ImageUri = direct.ValueOf(in.ImageURI)
	out.KeywordSearchableJobCustomAttributes = in.KeywordSearchableJobCustomAttributes
	// MISSING: DerivedInfo
	// MISSING: Suspended
	return out
}
func CompanyObservedState_FromProto(mapCtx *direct.MapContext, in *pb.Company) *krm.CompanyObservedState {
	if in == nil {
		return nil
	}
	out := &krm.CompanyObservedState{}
	// MISSING: Name
	// MISSING: DisplayName
	// MISSING: ExternalID
	// MISSING: Size
	// MISSING: HeadquartersAddress
	// MISSING: HiringAgency
	// MISSING: EeoText
	// MISSING: WebsiteURI
	// MISSING: CareerSiteURI
	// MISSING: ImageURI
	// MISSING: KeywordSearchableJobCustomAttributes
	out.DerivedInfo = Company_DerivedInfo_FromProto(mapCtx, in.GetDerivedInfo())
	out.Suspended = direct.LazyPtr(in.GetSuspended())
	return out
}
func CompanyObservedState_ToProto(mapCtx *direct.MapContext, in *krm.CompanyObservedState) *pb.Company {
	if in == nil {
		return nil
	}
	out := &pb.Company{}
	// MISSING: Name
	// MISSING: DisplayName
	// MISSING: ExternalID
	// MISSING: Size
	// MISSING: HeadquartersAddress
	// MISSING: HiringAgency
	// MISSING: EeoText
	// MISSING: WebsiteURI
	// MISSING: CareerSiteURI
	// MISSING: ImageURI
	// MISSING: KeywordSearchableJobCustomAttributes
	out.DerivedInfo = Company_DerivedInfo_ToProto(mapCtx, in.DerivedInfo)
	out.Suspended = direct.ValueOf(in.Suspended)
	return out
}
func Company_DerivedInfo_FromProto(mapCtx *direct.MapContext, in *pb.Company_DerivedInfo) *krm.Company_DerivedInfo {
	if in == nil {
		return nil
	}
	out := &krm.Company_DerivedInfo{}
	out.HeadquartersLocation = Location_FromProto(mapCtx, in.GetHeadquartersLocation())
	return out
}
func Company_DerivedInfo_ToProto(mapCtx *direct.MapContext, in *krm.Company_DerivedInfo) *pb.Company_DerivedInfo {
	if in == nil {
		return nil
	}
	out := &pb.Company_DerivedInfo{}
	out.HeadquartersLocation = Location_ToProto(mapCtx, in.HeadquartersLocation)
	return out
}
func Location_FromProto(mapCtx *direct.MapContext, in *pb.Location) *krm.Location {
	if in == nil {
		return nil
	}
	out := &krm.Location{}
	out.LocationType = direct.Enum_FromProto(mapCtx, in.GetLocationType())
	out.PostalAddress = PostalAddress_FromProto(mapCtx, in.GetPostalAddress())
	out.LatLng = LatLng_FromProto(mapCtx, in.GetLatLng())
	out.RadiusMiles = direct.LazyPtr(in.GetRadiusMiles())
	return out
}
func Location_ToProto(mapCtx *direct.MapContext, in *krm.Location) *pb.Location {
	if in == nil {
		return nil
	}
	out := &pb.Location{}
	out.LocationType = direct.Enum_ToProto[pb.Location_LocationType](mapCtx, in.LocationType)
	out.PostalAddress = PostalAddress_ToProto(mapCtx, in.PostalAddress)
	out.LatLng = LatLng_ToProto(mapCtx, in.LatLng)
	out.RadiusMiles = direct.ValueOf(in.RadiusMiles)
	return out
}
func TalentCompanyObservedState_FromProto(mapCtx *direct.MapContext, in *pb.Company) *krm.TalentCompanyObservedState {
	if in == nil {
		return nil
	}
	out := &krm.TalentCompanyObservedState{}
	// MISSING: Name
	// MISSING: DisplayName
	// MISSING: ExternalID
	// MISSING: Size
	// MISSING: HeadquartersAddress
	// MISSING: HiringAgency
	// MISSING: EeoText
	// MISSING: WebsiteURI
	// MISSING: CareerSiteURI
	// MISSING: ImageURI
	// MISSING: KeywordSearchableJobCustomAttributes
	// MISSING: DerivedInfo
	// MISSING: Suspended
	return out
}
func TalentCompanyObservedState_ToProto(mapCtx *direct.MapContext, in *krm.TalentCompanyObservedState) *pb.Company {
	if in == nil {
		return nil
	}
	out := &pb.Company{}
	// MISSING: Name
	// MISSING: DisplayName
	// MISSING: ExternalID
	// MISSING: Size
	// MISSING: HeadquartersAddress
	// MISSING: HiringAgency
	// MISSING: EeoText
	// MISSING: WebsiteURI
	// MISSING: CareerSiteURI
	// MISSING: ImageURI
	// MISSING: KeywordSearchableJobCustomAttributes
	// MISSING: DerivedInfo
	// MISSING: Suspended
	return out
}
func TalentCompanySpec_FromProto(mapCtx *direct.MapContext, in *pb.Company) *krm.TalentCompanySpec {
	if in == nil {
		return nil
	}
	out := &krm.TalentCompanySpec{}
	// MISSING: Name
	// MISSING: DisplayName
	// MISSING: ExternalID
	// MISSING: Size
	// MISSING: HeadquartersAddress
	// MISSING: HiringAgency
	// MISSING: EeoText
	// MISSING: WebsiteURI
	// MISSING: CareerSiteURI
	// MISSING: ImageURI
	// MISSING: KeywordSearchableJobCustomAttributes
	// MISSING: DerivedInfo
	// MISSING: Suspended
	return out
}
func TalentCompanySpec_ToProto(mapCtx *direct.MapContext, in *krm.TalentCompanySpec) *pb.Company {
	if in == nil {
		return nil
	}
	out := &pb.Company{}
	// MISSING: Name
	// MISSING: DisplayName
	// MISSING: ExternalID
	// MISSING: Size
	// MISSING: HeadquartersAddress
	// MISSING: HiringAgency
	// MISSING: EeoText
	// MISSING: WebsiteURI
	// MISSING: CareerSiteURI
	// MISSING: ImageURI
	// MISSING: KeywordSearchableJobCustomAttributes
	// MISSING: DerivedInfo
	// MISSING: Suspended
	return out
}
