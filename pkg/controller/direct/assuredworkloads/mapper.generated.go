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

package assuredworkloads

import (
	refs "github.com/GoogleCloudPlatform/k8s-config-connector/apis/refs/v1beta1"
	pb "cloud.google.com/go/assuredworkloads/apiv1beta1/assuredworkloadspb"
	krm "github.com/GoogleCloudPlatform/k8s-config-connector/apis/assuredworkloads/v1alpha1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct"
)
func Workload_FromProto(mapCtx *direct.MapContext, in *pb.Workload) *krm.Workload {
	if in == nil {
		return nil
	}
	out := &krm.Workload{}
	out.Name = direct.LazyPtr(in.GetName())
	out.DisplayName = direct.LazyPtr(in.GetDisplayName())
	// MISSING: Resources
	out.ComplianceRegime = direct.Enum_FromProto(mapCtx, in.GetComplianceRegime())
	// MISSING: CreateTime
	// MISSING: BillingAccount
	out.Il4Settings = Workload_IL4Settings_FromProto(mapCtx, in.GetIl4Settings())
	out.CjisSettings = Workload_CJISSettings_FromProto(mapCtx, in.GetCjisSettings())
	out.FedrampHighSettings = Workload_FedrampHighSettings_FromProto(mapCtx, in.GetFedrampHighSettings())
	out.FedrampModerateSettings = Workload_FedrampModerateSettings_FromProto(mapCtx, in.GetFedrampModerateSettings())
	out.Etag = direct.LazyPtr(in.GetEtag())
	out.Labels = in.Labels
	out.ProvisionedResourcesParent = direct.LazyPtr(in.GetProvisionedResourcesParent())
	out.KMSSettings = Workload_KMSSettings_FromProto(mapCtx, in.GetKmsSettings())
	out.ResourceSettings = direct.Slice_FromProto(mapCtx, in.ResourceSettings, Workload_ResourceSettings_FromProto)
	// MISSING: KajEnrollmentState
	out.EnableSovereignControls = direct.LazyPtr(in.GetEnableSovereignControls())
	// MISSING: SaaEnrollmentResponse
	// MISSING: CompliantButDisallowedServices
	return out
}
func Workload_ToProto(mapCtx *direct.MapContext, in *krm.Workload) *pb.Workload {
	if in == nil {
		return nil
	}
	out := &pb.Workload{}
	out.Name = direct.ValueOf(in.Name)
	out.DisplayName = direct.ValueOf(in.DisplayName)
	// MISSING: Resources
	out.ComplianceRegime = direct.Enum_ToProto[pb.Workload_ComplianceRegime](mapCtx, in.ComplianceRegime)
	// MISSING: CreateTime
	// MISSING: BillingAccount
	if oneof := Workload_IL4Settings_ToProto(mapCtx, in.Il4Settings); oneof != nil {
		out.ComplianceRegimeSettings = &pb.Workload_Il4Settings{Il4Settings: oneof}
	}
	if oneof := Workload_CJISSettings_ToProto(mapCtx, in.CjisSettings); oneof != nil {
		out.ComplianceRegimeSettings = &pb.Workload_CjisSettings{CjisSettings: oneof}
	}
	if oneof := Workload_FedrampHighSettings_ToProto(mapCtx, in.FedrampHighSettings); oneof != nil {
		out.ComplianceRegimeSettings = &pb.Workload_FedrampHighSettings_{FedrampHighSettings: oneof}
	}
	if oneof := Workload_FedrampModerateSettings_ToProto(mapCtx, in.FedrampModerateSettings); oneof != nil {
		out.ComplianceRegimeSettings = &pb.Workload_FedrampModerateSettings_{FedrampModerateSettings: oneof}
	}
	out.Etag = direct.ValueOf(in.Etag)
	out.Labels = in.Labels
	out.ProvisionedResourcesParent = direct.ValueOf(in.ProvisionedResourcesParent)
	out.KmsSettings = Workload_KMSSettings_ToProto(mapCtx, in.KMSSettings)
	out.ResourceSettings = direct.Slice_ToProto(mapCtx, in.ResourceSettings, Workload_ResourceSettings_ToProto)
	// MISSING: KajEnrollmentState
	out.EnableSovereignControls = direct.ValueOf(in.EnableSovereignControls)
	// MISSING: SaaEnrollmentResponse
	// MISSING: CompliantButDisallowedServices
	return out
}
func WorkloadObservedState_FromProto(mapCtx *direct.MapContext, in *pb.Workload) *krm.WorkloadObservedState {
	if in == nil {
		return nil
	}
	out := &krm.WorkloadObservedState{}
	// MISSING: Name
	// MISSING: DisplayName
	out.Resources = direct.Slice_FromProto(mapCtx, in.Resources, Workload_ResourceInfo_FromProto)
	// MISSING: ComplianceRegime
	out.CreateTime = direct.StringTimestamp_FromProto(mapCtx, in.GetCreateTime())
	out.BillingAccount = direct.LazyPtr(in.GetBillingAccount())
	// MISSING: Il4Settings
	// MISSING: CjisSettings
	// MISSING: FedrampHighSettings
	// MISSING: FedrampModerateSettings
	// MISSING: Etag
	// MISSING: Labels
	// MISSING: ProvisionedResourcesParent
	// MISSING: KMSSettings
	// MISSING: ResourceSettings
	out.KajEnrollmentState = direct.Enum_FromProto(mapCtx, in.GetKajEnrollmentState())
	// MISSING: EnableSovereignControls
	out.SaaEnrollmentResponse = Workload_SaaEnrollmentResponse_FromProto(mapCtx, in.GetSaaEnrollmentResponse())
	out.CompliantButDisallowedServices = in.CompliantButDisallowedServices
	return out
}
func WorkloadObservedState_ToProto(mapCtx *direct.MapContext, in *krm.WorkloadObservedState) *pb.Workload {
	if in == nil {
		return nil
	}
	out := &pb.Workload{}
	// MISSING: Name
	// MISSING: DisplayName
	out.Resources = direct.Slice_ToProto(mapCtx, in.Resources, Workload_ResourceInfo_ToProto)
	// MISSING: ComplianceRegime
	out.CreateTime = direct.StringTimestamp_ToProto(mapCtx, in.CreateTime)
	out.BillingAccount = direct.ValueOf(in.BillingAccount)
	// MISSING: Il4Settings
	// MISSING: CjisSettings
	// MISSING: FedrampHighSettings
	// MISSING: FedrampModerateSettings
	// MISSING: Etag
	// MISSING: Labels
	// MISSING: ProvisionedResourcesParent
	// MISSING: KMSSettings
	// MISSING: ResourceSettings
	out.KajEnrollmentState = direct.Enum_ToProto[pb.Workload_KajEnrollmentState](mapCtx, in.KajEnrollmentState)
	// MISSING: EnableSovereignControls
	out.SaaEnrollmentResponse = Workload_SaaEnrollmentResponse_ToProto(mapCtx, in.SaaEnrollmentResponse)
	out.CompliantButDisallowedServices = in.CompliantButDisallowedServices
	return out
}
func Workload_CJISSettings_FromProto(mapCtx *direct.MapContext, in *pb.Workload_CJISSettings) *krm.Workload_CJISSettings {
	if in == nil {
		return nil
	}
	out := &krm.Workload_CJISSettings{}
	out.KMSSettings = Workload_KMSSettings_FromProto(mapCtx, in.GetKmsSettings())
	return out
}
func Workload_CJISSettings_ToProto(mapCtx *direct.MapContext, in *krm.Workload_CJISSettings) *pb.Workload_CJISSettings {
	if in == nil {
		return nil
	}
	out := &pb.Workload_CJISSettings{}
	out.KmsSettings = Workload_KMSSettings_ToProto(mapCtx, in.KMSSettings)
	return out
}
func Workload_FedrampHighSettings_FromProto(mapCtx *direct.MapContext, in *pb.Workload_FedrampHighSettings) *krm.Workload_FedrampHighSettings {
	if in == nil {
		return nil
	}
	out := &krm.Workload_FedrampHighSettings{}
	out.KMSSettings = Workload_KMSSettings_FromProto(mapCtx, in.GetKmsSettings())
	return out
}
func Workload_FedrampHighSettings_ToProto(mapCtx *direct.MapContext, in *krm.Workload_FedrampHighSettings) *pb.Workload_FedrampHighSettings {
	if in == nil {
		return nil
	}
	out := &pb.Workload_FedrampHighSettings{}
	out.KmsSettings = Workload_KMSSettings_ToProto(mapCtx, in.KMSSettings)
	return out
}
func Workload_FedrampModerateSettings_FromProto(mapCtx *direct.MapContext, in *pb.Workload_FedrampModerateSettings) *krm.Workload_FedrampModerateSettings {
	if in == nil {
		return nil
	}
	out := &krm.Workload_FedrampModerateSettings{}
	out.KMSSettings = Workload_KMSSettings_FromProto(mapCtx, in.GetKmsSettings())
	return out
}
func Workload_FedrampModerateSettings_ToProto(mapCtx *direct.MapContext, in *krm.Workload_FedrampModerateSettings) *pb.Workload_FedrampModerateSettings {
	if in == nil {
		return nil
	}
	out := &pb.Workload_FedrampModerateSettings{}
	out.KmsSettings = Workload_KMSSettings_ToProto(mapCtx, in.KMSSettings)
	return out
}
func Workload_IL4Settings_FromProto(mapCtx *direct.MapContext, in *pb.Workload_IL4Settings) *krm.Workload_IL4Settings {
	if in == nil {
		return nil
	}
	out := &krm.Workload_IL4Settings{}
	out.KMSSettings = Workload_KMSSettings_FromProto(mapCtx, in.GetKmsSettings())
	return out
}
func Workload_IL4Settings_ToProto(mapCtx *direct.MapContext, in *krm.Workload_IL4Settings) *pb.Workload_IL4Settings {
	if in == nil {
		return nil
	}
	out := &pb.Workload_IL4Settings{}
	out.KmsSettings = Workload_KMSSettings_ToProto(mapCtx, in.KMSSettings)
	return out
}
func Workload_KMSSettings_FromProto(mapCtx *direct.MapContext, in *pb.Workload_KMSSettings) *krm.Workload_KMSSettings {
	if in == nil {
		return nil
	}
	out := &krm.Workload_KMSSettings{}
	out.NextRotationTime = direct.StringTimestamp_FromProto(mapCtx, in.GetNextRotationTime())
	out.RotationPeriod = direct.StringDuration_FromProto(mapCtx, in.GetRotationPeriod())
	return out
}
func Workload_KMSSettings_ToProto(mapCtx *direct.MapContext, in *krm.Workload_KMSSettings) *pb.Workload_KMSSettings {
	if in == nil {
		return nil
	}
	out := &pb.Workload_KMSSettings{}
	out.NextRotationTime = direct.StringTimestamp_ToProto(mapCtx, in.NextRotationTime)
	out.RotationPeriod = direct.StringDuration_ToProto(mapCtx, in.RotationPeriod)
	return out
}
func Workload_ResourceInfo_FromProto(mapCtx *direct.MapContext, in *pb.Workload_ResourceInfo) *krm.Workload_ResourceInfo {
	if in == nil {
		return nil
	}
	out := &krm.Workload_ResourceInfo{}
	out.ResourceID = direct.LazyPtr(in.GetResourceId())
	out.ResourceType = direct.Enum_FromProto(mapCtx, in.GetResourceType())
	return out
}
func Workload_ResourceInfo_ToProto(mapCtx *direct.MapContext, in *krm.Workload_ResourceInfo) *pb.Workload_ResourceInfo {
	if in == nil {
		return nil
	}
	out := &pb.Workload_ResourceInfo{}
	out.ResourceId = direct.ValueOf(in.ResourceID)
	out.ResourceType = direct.Enum_ToProto[pb.Workload_ResourceInfo_ResourceType](mapCtx, in.ResourceType)
	return out
}
func Workload_ResourceSettings_FromProto(mapCtx *direct.MapContext, in *pb.Workload_ResourceSettings) *krm.Workload_ResourceSettings {
	if in == nil {
		return nil
	}
	out := &krm.Workload_ResourceSettings{}
	out.ResourceID = direct.LazyPtr(in.GetResourceId())
	out.ResourceType = direct.Enum_FromProto(mapCtx, in.GetResourceType())
	out.DisplayName = direct.LazyPtr(in.GetDisplayName())
	return out
}
func Workload_ResourceSettings_ToProto(mapCtx *direct.MapContext, in *krm.Workload_ResourceSettings) *pb.Workload_ResourceSettings {
	if in == nil {
		return nil
	}
	out := &pb.Workload_ResourceSettings{}
	out.ResourceId = direct.ValueOf(in.ResourceID)
	out.ResourceType = direct.Enum_ToProto[pb.Workload_ResourceInfo_ResourceType](mapCtx, in.ResourceType)
	out.DisplayName = direct.ValueOf(in.DisplayName)
	return out
}
