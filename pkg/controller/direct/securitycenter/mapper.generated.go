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

package securitycenter

import (
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct"
	refs "github.com/GoogleCloudPlatform/k8s-config-connector/apis/refs/v1beta1"
	pb "cloud.google.com/go/securitycenter/settings/apiv1beta1/settingspb"
	krm "github.com/GoogleCloudPlatform/k8s-config-connector/apis/securitycenter/v1alpha1"
)
func BillingSettings_FromProto(mapCtx *direct.MapContext, in *pb.BillingSettings) *krm.BillingSettings {
	if in == nil {
		return nil
	}
	out := &krm.BillingSettings{}
	// MISSING: BillingTier
	// MISSING: BillingType
	// MISSING: StartTime
	// MISSING: ExpireTime
	return out
}
func BillingSettings_ToProto(mapCtx *direct.MapContext, in *krm.BillingSettings) *pb.BillingSettings {
	if in == nil {
		return nil
	}
	out := &pb.BillingSettings{}
	// MISSING: BillingTier
	// MISSING: BillingType
	// MISSING: StartTime
	// MISSING: ExpireTime
	return out
}
func BillingSettingsObservedState_FromProto(mapCtx *direct.MapContext, in *pb.BillingSettings) *krm.BillingSettingsObservedState {
	if in == nil {
		return nil
	}
	out := &krm.BillingSettingsObservedState{}
	out.BillingTier = direct.Enum_FromProto(mapCtx, in.GetBillingTier())
	out.BillingType = direct.Enum_FromProto(mapCtx, in.GetBillingType())
	out.StartTime = direct.StringTimestamp_FromProto(mapCtx, in.GetStartTime())
	out.ExpireTime = direct.StringTimestamp_FromProto(mapCtx, in.GetExpireTime())
	return out
}
func BillingSettingsObservedState_ToProto(mapCtx *direct.MapContext, in *krm.BillingSettingsObservedState) *pb.BillingSettings {
	if in == nil {
		return nil
	}
	out := &pb.BillingSettings{}
	out.BillingTier = direct.Enum_ToProto[pb.BillingTier](mapCtx, in.BillingTier)
	out.BillingType = direct.Enum_ToProto[pb.BillingType](mapCtx, in.BillingType)
	out.StartTime = direct.StringTimestamp_ToProto(mapCtx, in.StartTime)
	out.ExpireTime = direct.StringTimestamp_ToProto(mapCtx, in.ExpireTime)
	return out
}
func ComponentSettings_FromProto(mapCtx *direct.MapContext, in *pb.ComponentSettings) *krm.ComponentSettings {
	if in == nil {
		return nil
	}
	out := &krm.ComponentSettings{}
	out.Name = direct.LazyPtr(in.GetName())
	out.State = direct.Enum_FromProto(mapCtx, in.GetState())
	// MISSING: ProjectServiceAccount
	// MISSING: DetectorSettings
	// MISSING: Etag
	// MISSING: UpdateTime
	out.ContainerThreatDetectionSettings = ContainerThreatDetectionSettings_FromProto(mapCtx, in.GetContainerThreatDetectionSettings())
	out.EventThreatDetectionSettings = EventThreatDetectionSettings_FromProto(mapCtx, in.GetEventThreatDetectionSettings())
	out.SecurityHealthAnalyticsSettings = SecurityHealthAnalyticsSettings_FromProto(mapCtx, in.GetSecurityHealthAnalyticsSettings())
	out.WebSecurityScannerSettings = WebSecurityScanner_FromProto(mapCtx, in.GetWebSecurityScannerSettings())
	return out
}
func ComponentSettings_ToProto(mapCtx *direct.MapContext, in *krm.ComponentSettings) *pb.ComponentSettings {
	if in == nil {
		return nil
	}
	out := &pb.ComponentSettings{}
	out.Name = direct.ValueOf(in.Name)
	out.State = direct.Enum_ToProto[pb.ComponentEnablementState](mapCtx, in.State)
	// MISSING: ProjectServiceAccount
	// MISSING: DetectorSettings
	// MISSING: Etag
	// MISSING: UpdateTime
	if oneof := ContainerThreatDetectionSettings_ToProto(mapCtx, in.ContainerThreatDetectionSettings); oneof != nil {
		out.SpecificSettings = &pb.ComponentSettings_ContainerThreatDetectionSettings{ContainerThreatDetectionSettings: oneof}
	}
	if oneof := EventThreatDetectionSettings_ToProto(mapCtx, in.EventThreatDetectionSettings); oneof != nil {
		out.SpecificSettings = &pb.ComponentSettings_EventThreatDetectionSettings{EventThreatDetectionSettings: oneof}
	}
	if oneof := SecurityHealthAnalyticsSettings_ToProto(mapCtx, in.SecurityHealthAnalyticsSettings); oneof != nil {
		out.SpecificSettings = &pb.ComponentSettings_SecurityHealthAnalyticsSettings{SecurityHealthAnalyticsSettings: oneof}
	}
	if oneof := WebSecurityScanner_ToProto(mapCtx, in.WebSecurityScannerSettings); oneof != nil {
		out.SpecificSettings = &pb.ComponentSettings_WebSecurityScannerSettings{WebSecurityScannerSettings: oneof}
	}
	return out
}
func ComponentSettings_DetectorSettings_FromProto(mapCtx *direct.MapContext, in *pb.ComponentSettings_DetectorSettings) *krm.ComponentSettings_DetectorSettings {
	if in == nil {
		return nil
	}
	out := &krm.ComponentSettings_DetectorSettings{}
	out.State = direct.Enum_FromProto(mapCtx, in.GetState())
	return out
}
func ComponentSettings_DetectorSettings_ToProto(mapCtx *direct.MapContext, in *krm.ComponentSettings_DetectorSettings) *pb.ComponentSettings_DetectorSettings {
	if in == nil {
		return nil
	}
	out := &pb.ComponentSettings_DetectorSettings{}
	out.State = direct.Enum_ToProto[pb.ComponentEnablementState](mapCtx, in.State)
	return out
}
func ContainerThreatDetectionSettings_FromProto(mapCtx *direct.MapContext, in *pb.ContainerThreatDetectionSettings) *krm.ContainerThreatDetectionSettings {
	if in == nil {
		return nil
	}
	out := &krm.ContainerThreatDetectionSettings{}
	return out
}
func ContainerThreatDetectionSettings_ToProto(mapCtx *direct.MapContext, in *krm.ContainerThreatDetectionSettings) *pb.ContainerThreatDetectionSettings {
	if in == nil {
		return nil
	}
	out := &pb.ContainerThreatDetectionSettings{}
	return out
}
func EventThreatDetectionSettings_FromProto(mapCtx *direct.MapContext, in *pb.EventThreatDetectionSettings) *krm.EventThreatDetectionSettings {
	if in == nil {
		return nil
	}
	out := &krm.EventThreatDetectionSettings{}
	return out
}
func EventThreatDetectionSettings_ToProto(mapCtx *direct.MapContext, in *krm.EventThreatDetectionSettings) *pb.EventThreatDetectionSettings {
	if in == nil {
		return nil
	}
	out := &pb.EventThreatDetectionSettings{}
	return out
}
func SecurityHealthAnalyticsSettings_FromProto(mapCtx *direct.MapContext, in *pb.SecurityHealthAnalyticsSettings) *krm.SecurityHealthAnalyticsSettings {
	if in == nil {
		return nil
	}
	out := &krm.SecurityHealthAnalyticsSettings{}
	out.NonOrgIamMemberSettings = SecurityHealthAnalyticsSettings_NonOrgIamMemberSettings_FromProto(mapCtx, in.GetNonOrgIamMemberSettings())
	out.AdminServiceAccountSettings = SecurityHealthAnalyticsSettings_AdminServiceAccountSettings_FromProto(mapCtx, in.GetAdminServiceAccountSettings())
	return out
}
func SecurityHealthAnalyticsSettings_ToProto(mapCtx *direct.MapContext, in *krm.SecurityHealthAnalyticsSettings) *pb.SecurityHealthAnalyticsSettings {
	if in == nil {
		return nil
	}
	out := &pb.SecurityHealthAnalyticsSettings{}
	out.NonOrgIamMemberSettings = SecurityHealthAnalyticsSettings_NonOrgIamMemberSettings_ToProto(mapCtx, in.NonOrgIamMemberSettings)
	out.AdminServiceAccountSettings = SecurityHealthAnalyticsSettings_AdminServiceAccountSettings_ToProto(mapCtx, in.AdminServiceAccountSettings)
	return out
}
func SecurityHealthAnalyticsSettings_AdminServiceAccountSettings_FromProto(mapCtx *direct.MapContext, in *pb.SecurityHealthAnalyticsSettings_AdminServiceAccountSettings) *krm.SecurityHealthAnalyticsSettings_AdminServiceAccountSettings {
	if in == nil {
		return nil
	}
	out := &krm.SecurityHealthAnalyticsSettings_AdminServiceAccountSettings{}
	out.ApprovedIdentities = in.ApprovedIdentities
	return out
}
func SecurityHealthAnalyticsSettings_AdminServiceAccountSettings_ToProto(mapCtx *direct.MapContext, in *krm.SecurityHealthAnalyticsSettings_AdminServiceAccountSettings) *pb.SecurityHealthAnalyticsSettings_AdminServiceAccountSettings {
	if in == nil {
		return nil
	}
	out := &pb.SecurityHealthAnalyticsSettings_AdminServiceAccountSettings{}
	out.ApprovedIdentities = in.ApprovedIdentities
	return out
}
func SecurityHealthAnalyticsSettings_NonOrgIamMemberSettings_FromProto(mapCtx *direct.MapContext, in *pb.SecurityHealthAnalyticsSettings_NonOrgIamMemberSettings) *krm.SecurityHealthAnalyticsSettings_NonOrgIamMemberSettings {
	if in == nil {
		return nil
	}
	out := &krm.SecurityHealthAnalyticsSettings_NonOrgIamMemberSettings{}
	out.ApprovedIdentities = in.ApprovedIdentities
	return out
}
func SecurityHealthAnalyticsSettings_NonOrgIamMemberSettings_ToProto(mapCtx *direct.MapContext, in *krm.SecurityHealthAnalyticsSettings_NonOrgIamMemberSettings) *pb.SecurityHealthAnalyticsSettings_NonOrgIamMemberSettings {
	if in == nil {
		return nil
	}
	out := &pb.SecurityHealthAnalyticsSettings_NonOrgIamMemberSettings{}
	out.ApprovedIdentities = in.ApprovedIdentities
	return out
}
func SecuritycenterSettingsObservedState_FromProto(mapCtx *direct.MapContext, in *pb.Settings) *krm.SecuritycenterSettingsObservedState {
	if in == nil {
		return nil
	}
	out := &krm.SecuritycenterSettingsObservedState{}
	// MISSING: Name
	// MISSING: BillingSettings
	// MISSING: State
	// MISSING: OrgServiceAccount
	// MISSING: SinkSettings
	// MISSING: ComponentSettings
	// MISSING: DetectorGroupSettings
	// MISSING: Etag
	// MISSING: UpdateTime
	return out
}
func SecuritycenterSettingsObservedState_ToProto(mapCtx *direct.MapContext, in *krm.SecuritycenterSettingsObservedState) *pb.Settings {
	if in == nil {
		return nil
	}
	out := &pb.Settings{}
	// MISSING: Name
	// MISSING: BillingSettings
	// MISSING: State
	// MISSING: OrgServiceAccount
	// MISSING: SinkSettings
	// MISSING: ComponentSettings
	// MISSING: DetectorGroupSettings
	// MISSING: Etag
	// MISSING: UpdateTime
	return out
}
func SecuritycenterSettingsSpec_FromProto(mapCtx *direct.MapContext, in *pb.Settings) *krm.SecuritycenterSettingsSpec {
	if in == nil {
		return nil
	}
	out := &krm.SecuritycenterSettingsSpec{}
	// MISSING: Name
	// MISSING: BillingSettings
	// MISSING: State
	// MISSING: OrgServiceAccount
	// MISSING: SinkSettings
	// MISSING: ComponentSettings
	// MISSING: DetectorGroupSettings
	// MISSING: Etag
	// MISSING: UpdateTime
	return out
}
func SecuritycenterSettingsSpec_ToProto(mapCtx *direct.MapContext, in *krm.SecuritycenterSettingsSpec) *pb.Settings {
	if in == nil {
		return nil
	}
	out := &pb.Settings{}
	// MISSING: Name
	// MISSING: BillingSettings
	// MISSING: State
	// MISSING: OrgServiceAccount
	// MISSING: SinkSettings
	// MISSING: ComponentSettings
	// MISSING: DetectorGroupSettings
	// MISSING: Etag
	// MISSING: UpdateTime
	return out
}
func Settings_FromProto(mapCtx *direct.MapContext, in *pb.Settings) *krm.Settings {
	if in == nil {
		return nil
	}
	out := &krm.Settings{}
	out.Name = direct.LazyPtr(in.GetName())
	out.BillingSettings = BillingSettings_FromProto(mapCtx, in.GetBillingSettings())
	out.State = direct.Enum_FromProto(mapCtx, in.GetState())
	// MISSING: OrgServiceAccount
	out.SinkSettings = SinkSettings_FromProto(mapCtx, in.GetSinkSettings())
	// MISSING: ComponentSettings
	// MISSING: DetectorGroupSettings
	out.Etag = direct.LazyPtr(in.GetEtag())
	// MISSING: UpdateTime
	return out
}
func Settings_ToProto(mapCtx *direct.MapContext, in *krm.Settings) *pb.Settings {
	if in == nil {
		return nil
	}
	out := &pb.Settings{}
	out.Name = direct.ValueOf(in.Name)
	out.BillingSettings = BillingSettings_ToProto(mapCtx, in.BillingSettings)
	out.State = direct.Enum_ToProto[pb.Settings_OnboardingState](mapCtx, in.State)
	// MISSING: OrgServiceAccount
	out.SinkSettings = SinkSettings_ToProto(mapCtx, in.SinkSettings)
	// MISSING: ComponentSettings
	// MISSING: DetectorGroupSettings
	out.Etag = direct.ValueOf(in.Etag)
	// MISSING: UpdateTime
	return out
}
func SettingsObservedState_FromProto(mapCtx *direct.MapContext, in *pb.Settings) *krm.SettingsObservedState {
	if in == nil {
		return nil
	}
	out := &krm.SettingsObservedState{}
	// MISSING: Name
	out.BillingSettings = BillingSettingsObservedState_FromProto(mapCtx, in.GetBillingSettings())
	// MISSING: State
	out.OrgServiceAccount = direct.LazyPtr(in.GetOrgServiceAccount())
	// MISSING: SinkSettings
	// MISSING: ComponentSettings
	// MISSING: DetectorGroupSettings
	// MISSING: Etag
	out.UpdateTime = direct.StringTimestamp_FromProto(mapCtx, in.GetUpdateTime())
	return out
}
func SettingsObservedState_ToProto(mapCtx *direct.MapContext, in *krm.SettingsObservedState) *pb.Settings {
	if in == nil {
		return nil
	}
	out := &pb.Settings{}
	// MISSING: Name
	out.BillingSettings = BillingSettingsObservedState_ToProto(mapCtx, in.BillingSettings)
	// MISSING: State
	out.OrgServiceAccount = direct.ValueOf(in.OrgServiceAccount)
	// MISSING: SinkSettings
	// MISSING: ComponentSettings
	// MISSING: DetectorGroupSettings
	// MISSING: Etag
	out.UpdateTime = direct.StringTimestamp_ToProto(mapCtx, in.UpdateTime)
	return out
}
func Settings_DetectorGroupSettings_FromProto(mapCtx *direct.MapContext, in *pb.Settings_DetectorGroupSettings) *krm.Settings_DetectorGroupSettings {
	if in == nil {
		return nil
	}
	out := &krm.Settings_DetectorGroupSettings{}
	out.State = direct.Enum_FromProto(mapCtx, in.GetState())
	return out
}
func Settings_DetectorGroupSettings_ToProto(mapCtx *direct.MapContext, in *krm.Settings_DetectorGroupSettings) *pb.Settings_DetectorGroupSettings {
	if in == nil {
		return nil
	}
	out := &pb.Settings_DetectorGroupSettings{}
	out.State = direct.Enum_ToProto[pb.ComponentEnablementState](mapCtx, in.State)
	return out
}
func SinkSettings_FromProto(mapCtx *direct.MapContext, in *pb.SinkSettings) *krm.SinkSettings {
	if in == nil {
		return nil
	}
	out := &krm.SinkSettings{}
	out.LoggingSinkProject = direct.LazyPtr(in.GetLoggingSinkProject())
	return out
}
func SinkSettings_ToProto(mapCtx *direct.MapContext, in *krm.SinkSettings) *pb.SinkSettings {
	if in == nil {
		return nil
	}
	out := &pb.SinkSettings{}
	out.LoggingSinkProject = direct.ValueOf(in.LoggingSinkProject)
	return out
}
func WebSecurityScanner_FromProto(mapCtx *direct.MapContext, in *pb.WebSecurityScanner) *krm.WebSecurityScanner {
	if in == nil {
		return nil
	}
	out := &krm.WebSecurityScanner{}
	return out
}
func WebSecurityScanner_ToProto(mapCtx *direct.MapContext, in *krm.WebSecurityScanner) *pb.WebSecurityScanner {
	if in == nil {
		return nil
	}
	out := &pb.WebSecurityScanner{}
	return out
}
