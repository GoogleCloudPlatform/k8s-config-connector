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
func CompensationInfo_FromProto(mapCtx *direct.MapContext, in *pb.CompensationInfo) *krm.CompensationInfo {
	if in == nil {
		return nil
	}
	out := &krm.CompensationInfo{}
	out.Entries = direct.Slice_FromProto(mapCtx, in.Entries, CompensationInfo_CompensationEntry_FromProto)
	// MISSING: AnnualizedBaseCompensationRange
	// MISSING: AnnualizedTotalCompensationRange
	return out
}
func CompensationInfo_ToProto(mapCtx *direct.MapContext, in *krm.CompensationInfo) *pb.CompensationInfo {
	if in == nil {
		return nil
	}
	out := &pb.CompensationInfo{}
	out.Entries = direct.Slice_ToProto(mapCtx, in.Entries, CompensationInfo_CompensationEntry_ToProto)
	// MISSING: AnnualizedBaseCompensationRange
	// MISSING: AnnualizedTotalCompensationRange
	return out
}
func CompensationInfoObservedState_FromProto(mapCtx *direct.MapContext, in *pb.CompensationInfo) *krm.CompensationInfoObservedState {
	if in == nil {
		return nil
	}
	out := &krm.CompensationInfoObservedState{}
	// MISSING: Entries
	out.AnnualizedBaseCompensationRange = CompensationInfo_CompensationRange_FromProto(mapCtx, in.GetAnnualizedBaseCompensationRange())
	out.AnnualizedTotalCompensationRange = CompensationInfo_CompensationRange_FromProto(mapCtx, in.GetAnnualizedTotalCompensationRange())
	return out
}
func CompensationInfoObservedState_ToProto(mapCtx *direct.MapContext, in *krm.CompensationInfoObservedState) *pb.CompensationInfo {
	if in == nil {
		return nil
	}
	out := &pb.CompensationInfo{}
	// MISSING: Entries
	out.AnnualizedBaseCompensationRange = CompensationInfo_CompensationRange_ToProto(mapCtx, in.AnnualizedBaseCompensationRange)
	out.AnnualizedTotalCompensationRange = CompensationInfo_CompensationRange_ToProto(mapCtx, in.AnnualizedTotalCompensationRange)
	return out
}
func CompensationInfo_CompensationEntry_FromProto(mapCtx *direct.MapContext, in *pb.CompensationInfo_CompensationEntry) *krm.CompensationInfo_CompensationEntry {
	if in == nil {
		return nil
	}
	out := &krm.CompensationInfo_CompensationEntry{}
	out.Type = direct.Enum_FromProto(mapCtx, in.GetType())
	out.Unit = direct.Enum_FromProto(mapCtx, in.GetUnit())
	out.Amount = Money_FromProto(mapCtx, in.GetAmount())
	out.Range = CompensationInfo_CompensationRange_FromProto(mapCtx, in.GetRange())
	out.Description = direct.LazyPtr(in.GetDescription())
	out.ExpectedUnitsPerYear = DoubleValue_FromProto(mapCtx, in.GetExpectedUnitsPerYear())
	return out
}
func CompensationInfo_CompensationEntry_ToProto(mapCtx *direct.MapContext, in *krm.CompensationInfo_CompensationEntry) *pb.CompensationInfo_CompensationEntry {
	if in == nil {
		return nil
	}
	out := &pb.CompensationInfo_CompensationEntry{}
	out.Type = direct.Enum_ToProto[pb.CompensationInfo_CompensationType](mapCtx, in.Type)
	out.Unit = direct.Enum_ToProto[pb.CompensationInfo_CompensationUnit](mapCtx, in.Unit)
	if oneof := Money_ToProto(mapCtx, in.Amount); oneof != nil {
		out.CompensationAmount = &pb.CompensationInfo_CompensationEntry_Amount{Amount: oneof}
	}
	if oneof := CompensationInfo_CompensationRange_ToProto(mapCtx, in.Range); oneof != nil {
		out.CompensationAmount = &pb.CompensationInfo_CompensationEntry_Range{Range: oneof}
	}
	out.Description = direct.ValueOf(in.Description)
	out.ExpectedUnitsPerYear = DoubleValue_ToProto(mapCtx, in.ExpectedUnitsPerYear)
	return out
}
func CompensationInfo_CompensationRange_FromProto(mapCtx *direct.MapContext, in *pb.CompensationInfo_CompensationRange) *krm.CompensationInfo_CompensationRange {
	if in == nil {
		return nil
	}
	out := &krm.CompensationInfo_CompensationRange{}
	out.MaxCompensation = Money_FromProto(mapCtx, in.GetMaxCompensation())
	out.MinCompensation = Money_FromProto(mapCtx, in.GetMinCompensation())
	return out
}
func CompensationInfo_CompensationRange_ToProto(mapCtx *direct.MapContext, in *krm.CompensationInfo_CompensationRange) *pb.CompensationInfo_CompensationRange {
	if in == nil {
		return nil
	}
	out := &pb.CompensationInfo_CompensationRange{}
	out.MaxCompensation = Money_ToProto(mapCtx, in.MaxCompensation)
	out.MinCompensation = Money_ToProto(mapCtx, in.MinCompensation)
	return out
}
func CustomAttribute_FromProto(mapCtx *direct.MapContext, in *pb.CustomAttribute) *krm.CustomAttribute {
	if in == nil {
		return nil
	}
	out := &krm.CustomAttribute{}
	out.StringValues = in.StringValues
	out.LongValues = in.LongValues
	out.Filterable = direct.LazyPtr(in.GetFilterable())
	out.KeywordSearchable = direct.LazyPtr(in.GetKeywordSearchable())
	return out
}
func CustomAttribute_ToProto(mapCtx *direct.MapContext, in *krm.CustomAttribute) *pb.CustomAttribute {
	if in == nil {
		return nil
	}
	out := &pb.CustomAttribute{}
	out.StringValues = in.StringValues
	out.LongValues = in.LongValues
	out.Filterable = direct.ValueOf(in.Filterable)
	out.KeywordSearchable = direct.ValueOf(in.KeywordSearchable)
	return out
}
func Job_FromProto(mapCtx *direct.MapContext, in *pb.Job) *krm.Job {
	if in == nil {
		return nil
	}
	out := &krm.Job{}
	out.Name = direct.LazyPtr(in.GetName())
	out.Company = direct.LazyPtr(in.GetCompany())
	out.RequisitionID = direct.LazyPtr(in.GetRequisitionId())
	out.Title = direct.LazyPtr(in.GetTitle())
	out.Description = direct.LazyPtr(in.GetDescription())
	out.Addresses = in.Addresses
	out.ApplicationInfo = Job_ApplicationInfo_FromProto(mapCtx, in.GetApplicationInfo())
	out.JobBenefits = direct.EnumSlice_FromProto(mapCtx, in.JobBenefits)
	out.CompensationInfo = CompensationInfo_FromProto(mapCtx, in.GetCompensationInfo())
	// MISSING: CustomAttributes
	out.DegreeTypes = direct.EnumSlice_FromProto(mapCtx, in.DegreeTypes)
	out.Department = direct.LazyPtr(in.GetDepartment())
	out.EmploymentTypes = direct.EnumSlice_FromProto(mapCtx, in.EmploymentTypes)
	out.Incentives = direct.LazyPtr(in.GetIncentives())
	out.LanguageCode = direct.LazyPtr(in.GetLanguageCode())
	out.JobLevel = direct.Enum_FromProto(mapCtx, in.GetJobLevel())
	out.PromotionValue = direct.LazyPtr(in.GetPromotionValue())
	out.Qualifications = direct.LazyPtr(in.GetQualifications())
	out.Responsibilities = direct.LazyPtr(in.GetResponsibilities())
	out.PostingRegion = direct.Enum_FromProto(mapCtx, in.GetPostingRegion())
	out.Visibility = direct.Enum_FromProto(mapCtx, in.GetVisibility())
	out.JobStartTime = direct.StringTimestamp_FromProto(mapCtx, in.GetJobStartTime())
	out.JobEndTime = direct.StringTimestamp_FromProto(mapCtx, in.GetJobEndTime())
	out.PostingPublishTime = direct.StringTimestamp_FromProto(mapCtx, in.GetPostingPublishTime())
	out.PostingExpireTime = direct.StringTimestamp_FromProto(mapCtx, in.GetPostingExpireTime())
	// MISSING: PostingCreateTime
	// MISSING: PostingUpdateTime
	// MISSING: CompanyDisplayName
	// MISSING: DerivedInfo
	out.ProcessingOptions = Job_ProcessingOptions_FromProto(mapCtx, in.GetProcessingOptions())
	return out
}
func Job_ToProto(mapCtx *direct.MapContext, in *krm.Job) *pb.Job {
	if in == nil {
		return nil
	}
	out := &pb.Job{}
	out.Name = direct.ValueOf(in.Name)
	out.Company = direct.ValueOf(in.Company)
	out.RequisitionId = direct.ValueOf(in.RequisitionID)
	out.Title = direct.ValueOf(in.Title)
	out.Description = direct.ValueOf(in.Description)
	out.Addresses = in.Addresses
	out.ApplicationInfo = Job_ApplicationInfo_ToProto(mapCtx, in.ApplicationInfo)
	out.JobBenefits = direct.EnumSlice_ToProto[pb.JobBenefit](mapCtx, in.JobBenefits)
	out.CompensationInfo = CompensationInfo_ToProto(mapCtx, in.CompensationInfo)
	// MISSING: CustomAttributes
	out.DegreeTypes = direct.EnumSlice_ToProto[pb.DegreeType](mapCtx, in.DegreeTypes)
	out.Department = direct.ValueOf(in.Department)
	out.EmploymentTypes = direct.EnumSlice_ToProto[pb.EmploymentType](mapCtx, in.EmploymentTypes)
	out.Incentives = direct.ValueOf(in.Incentives)
	out.LanguageCode = direct.ValueOf(in.LanguageCode)
	out.JobLevel = direct.Enum_ToProto[pb.JobLevel](mapCtx, in.JobLevel)
	out.PromotionValue = direct.ValueOf(in.PromotionValue)
	out.Qualifications = direct.ValueOf(in.Qualifications)
	out.Responsibilities = direct.ValueOf(in.Responsibilities)
	out.PostingRegion = direct.Enum_ToProto[pb.PostingRegion](mapCtx, in.PostingRegion)
	out.Visibility = direct.Enum_ToProto[pb.Visibility](mapCtx, in.Visibility)
	out.JobStartTime = direct.StringTimestamp_ToProto(mapCtx, in.JobStartTime)
	out.JobEndTime = direct.StringTimestamp_ToProto(mapCtx, in.JobEndTime)
	out.PostingPublishTime = direct.StringTimestamp_ToProto(mapCtx, in.PostingPublishTime)
	out.PostingExpireTime = direct.StringTimestamp_ToProto(mapCtx, in.PostingExpireTime)
	// MISSING: PostingCreateTime
	// MISSING: PostingUpdateTime
	// MISSING: CompanyDisplayName
	// MISSING: DerivedInfo
	out.ProcessingOptions = Job_ProcessingOptions_ToProto(mapCtx, in.ProcessingOptions)
	return out
}
func JobObservedState_FromProto(mapCtx *direct.MapContext, in *pb.Job) *krm.JobObservedState {
	if in == nil {
		return nil
	}
	out := &krm.JobObservedState{}
	// MISSING: Name
	// MISSING: Company
	// MISSING: RequisitionID
	// MISSING: Title
	// MISSING: Description
	// MISSING: Addresses
	// MISSING: ApplicationInfo
	// MISSING: JobBenefits
	out.CompensationInfo = CompensationInfoObservedState_FromProto(mapCtx, in.GetCompensationInfo())
	// MISSING: CustomAttributes
	// MISSING: DegreeTypes
	// MISSING: Department
	// MISSING: EmploymentTypes
	// MISSING: Incentives
	// MISSING: LanguageCode
	// MISSING: JobLevel
	// MISSING: PromotionValue
	// MISSING: Qualifications
	// MISSING: Responsibilities
	// MISSING: PostingRegion
	// MISSING: Visibility
	// MISSING: JobStartTime
	// MISSING: JobEndTime
	// MISSING: PostingPublishTime
	// MISSING: PostingExpireTime
	out.PostingCreateTime = direct.StringTimestamp_FromProto(mapCtx, in.GetPostingCreateTime())
	out.PostingUpdateTime = direct.StringTimestamp_FromProto(mapCtx, in.GetPostingUpdateTime())
	out.CompanyDisplayName = direct.LazyPtr(in.GetCompanyDisplayName())
	out.DerivedInfo = Job_DerivedInfo_FromProto(mapCtx, in.GetDerivedInfo())
	// MISSING: ProcessingOptions
	return out
}
func JobObservedState_ToProto(mapCtx *direct.MapContext, in *krm.JobObservedState) *pb.Job {
	if in == nil {
		return nil
	}
	out := &pb.Job{}
	// MISSING: Name
	// MISSING: Company
	// MISSING: RequisitionID
	// MISSING: Title
	// MISSING: Description
	// MISSING: Addresses
	// MISSING: ApplicationInfo
	// MISSING: JobBenefits
	out.CompensationInfo = CompensationInfoObservedState_ToProto(mapCtx, in.CompensationInfo)
	// MISSING: CustomAttributes
	// MISSING: DegreeTypes
	// MISSING: Department
	// MISSING: EmploymentTypes
	// MISSING: Incentives
	// MISSING: LanguageCode
	// MISSING: JobLevel
	// MISSING: PromotionValue
	// MISSING: Qualifications
	// MISSING: Responsibilities
	// MISSING: PostingRegion
	// MISSING: Visibility
	// MISSING: JobStartTime
	// MISSING: JobEndTime
	// MISSING: PostingPublishTime
	// MISSING: PostingExpireTime
	out.PostingCreateTime = direct.StringTimestamp_ToProto(mapCtx, in.PostingCreateTime)
	out.PostingUpdateTime = direct.StringTimestamp_ToProto(mapCtx, in.PostingUpdateTime)
	out.CompanyDisplayName = direct.ValueOf(in.CompanyDisplayName)
	out.DerivedInfo = Job_DerivedInfo_ToProto(mapCtx, in.DerivedInfo)
	// MISSING: ProcessingOptions
	return out
}
func Job_ApplicationInfo_FromProto(mapCtx *direct.MapContext, in *pb.Job_ApplicationInfo) *krm.Job_ApplicationInfo {
	if in == nil {
		return nil
	}
	out := &krm.Job_ApplicationInfo{}
	out.Emails = in.Emails
	out.Instruction = direct.LazyPtr(in.GetInstruction())
	out.Uris = in.Uris
	return out
}
func Job_ApplicationInfo_ToProto(mapCtx *direct.MapContext, in *krm.Job_ApplicationInfo) *pb.Job_ApplicationInfo {
	if in == nil {
		return nil
	}
	out := &pb.Job_ApplicationInfo{}
	out.Emails = in.Emails
	out.Instruction = direct.ValueOf(in.Instruction)
	out.Uris = in.Uris
	return out
}
func Job_DerivedInfo_FromProto(mapCtx *direct.MapContext, in *pb.Job_DerivedInfo) *krm.Job_DerivedInfo {
	if in == nil {
		return nil
	}
	out := &krm.Job_DerivedInfo{}
	out.Locations = direct.Slice_FromProto(mapCtx, in.Locations, Location_FromProto)
	out.JobCategories = direct.EnumSlice_FromProto(mapCtx, in.JobCategories)
	return out
}
func Job_DerivedInfo_ToProto(mapCtx *direct.MapContext, in *krm.Job_DerivedInfo) *pb.Job_DerivedInfo {
	if in == nil {
		return nil
	}
	out := &pb.Job_DerivedInfo{}
	out.Locations = direct.Slice_ToProto(mapCtx, in.Locations, Location_ToProto)
	out.JobCategories = direct.EnumSlice_ToProto[pb.JobCategory](mapCtx, in.JobCategories)
	return out
}
func Job_ProcessingOptions_FromProto(mapCtx *direct.MapContext, in *pb.Job_ProcessingOptions) *krm.Job_ProcessingOptions {
	if in == nil {
		return nil
	}
	out := &krm.Job_ProcessingOptions{}
	out.DisableStreetAddressResolution = direct.LazyPtr(in.GetDisableStreetAddressResolution())
	out.HTMLSanitization = direct.Enum_FromProto(mapCtx, in.GetHtmlSanitization())
	return out
}
func Job_ProcessingOptions_ToProto(mapCtx *direct.MapContext, in *krm.Job_ProcessingOptions) *pb.Job_ProcessingOptions {
	if in == nil {
		return nil
	}
	out := &pb.Job_ProcessingOptions{}
	out.DisableStreetAddressResolution = direct.ValueOf(in.DisableStreetAddressResolution)
	out.HtmlSanitization = direct.Enum_ToProto[pb.HtmlSanitization](mapCtx, in.HTMLSanitization)
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
func TalentJobObservedState_FromProto(mapCtx *direct.MapContext, in *pb.Job) *krm.TalentJobObservedState {
	if in == nil {
		return nil
	}
	out := &krm.TalentJobObservedState{}
	// MISSING: Name
	// MISSING: Company
	// MISSING: RequisitionID
	// MISSING: Title
	// MISSING: Description
	// MISSING: Addresses
	// MISSING: ApplicationInfo
	// MISSING: JobBenefits
	// MISSING: CompensationInfo
	// MISSING: CustomAttributes
	// MISSING: DegreeTypes
	// MISSING: Department
	// MISSING: EmploymentTypes
	// MISSING: Incentives
	// MISSING: LanguageCode
	// MISSING: JobLevel
	// MISSING: PromotionValue
	// MISSING: Qualifications
	// MISSING: Responsibilities
	// MISSING: PostingRegion
	// MISSING: Visibility
	// MISSING: JobStartTime
	// MISSING: JobEndTime
	// MISSING: PostingPublishTime
	// MISSING: PostingExpireTime
	// MISSING: PostingCreateTime
	// MISSING: PostingUpdateTime
	// MISSING: CompanyDisplayName
	// MISSING: DerivedInfo
	// MISSING: ProcessingOptions
	return out
}
func TalentJobObservedState_ToProto(mapCtx *direct.MapContext, in *krm.TalentJobObservedState) *pb.Job {
	if in == nil {
		return nil
	}
	out := &pb.Job{}
	// MISSING: Name
	// MISSING: Company
	// MISSING: RequisitionID
	// MISSING: Title
	// MISSING: Description
	// MISSING: Addresses
	// MISSING: ApplicationInfo
	// MISSING: JobBenefits
	// MISSING: CompensationInfo
	// MISSING: CustomAttributes
	// MISSING: DegreeTypes
	// MISSING: Department
	// MISSING: EmploymentTypes
	// MISSING: Incentives
	// MISSING: LanguageCode
	// MISSING: JobLevel
	// MISSING: PromotionValue
	// MISSING: Qualifications
	// MISSING: Responsibilities
	// MISSING: PostingRegion
	// MISSING: Visibility
	// MISSING: JobStartTime
	// MISSING: JobEndTime
	// MISSING: PostingPublishTime
	// MISSING: PostingExpireTime
	// MISSING: PostingCreateTime
	// MISSING: PostingUpdateTime
	// MISSING: CompanyDisplayName
	// MISSING: DerivedInfo
	// MISSING: ProcessingOptions
	return out
}
func TalentJobSpec_FromProto(mapCtx *direct.MapContext, in *pb.Job) *krm.TalentJobSpec {
	if in == nil {
		return nil
	}
	out := &krm.TalentJobSpec{}
	// MISSING: Name
	// MISSING: Company
	// MISSING: RequisitionID
	// MISSING: Title
	// MISSING: Description
	// MISSING: Addresses
	// MISSING: ApplicationInfo
	// MISSING: JobBenefits
	// MISSING: CompensationInfo
	// MISSING: CustomAttributes
	// MISSING: DegreeTypes
	// MISSING: Department
	// MISSING: EmploymentTypes
	// MISSING: Incentives
	// MISSING: LanguageCode
	// MISSING: JobLevel
	// MISSING: PromotionValue
	// MISSING: Qualifications
	// MISSING: Responsibilities
	// MISSING: PostingRegion
	// MISSING: Visibility
	// MISSING: JobStartTime
	// MISSING: JobEndTime
	// MISSING: PostingPublishTime
	// MISSING: PostingExpireTime
	// MISSING: PostingCreateTime
	// MISSING: PostingUpdateTime
	// MISSING: CompanyDisplayName
	// MISSING: DerivedInfo
	// MISSING: ProcessingOptions
	return out
}
func TalentJobSpec_ToProto(mapCtx *direct.MapContext, in *krm.TalentJobSpec) *pb.Job {
	if in == nil {
		return nil
	}
	out := &pb.Job{}
	// MISSING: Name
	// MISSING: Company
	// MISSING: RequisitionID
	// MISSING: Title
	// MISSING: Description
	// MISSING: Addresses
	// MISSING: ApplicationInfo
	// MISSING: JobBenefits
	// MISSING: CompensationInfo
	// MISSING: CustomAttributes
	// MISSING: DegreeTypes
	// MISSING: Department
	// MISSING: EmploymentTypes
	// MISSING: Incentives
	// MISSING: LanguageCode
	// MISSING: JobLevel
	// MISSING: PromotionValue
	// MISSING: Qualifications
	// MISSING: Responsibilities
	// MISSING: PostingRegion
	// MISSING: Visibility
	// MISSING: JobStartTime
	// MISSING: JobEndTime
	// MISSING: PostingPublishTime
	// MISSING: PostingExpireTime
	// MISSING: PostingCreateTime
	// MISSING: PostingUpdateTime
	// MISSING: CompanyDisplayName
	// MISSING: DerivedInfo
	// MISSING: ProcessingOptions
	return out
}
