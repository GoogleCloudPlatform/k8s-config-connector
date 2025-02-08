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
	krm "github.com/GoogleCloudPlatform/k8s-config-connector/apis/securitycenter/v1alpha1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct"
	refs "github.com/GoogleCloudPlatform/k8s-config-connector/apis/refs/v1beta1"
	pb "cloud.google.com/go/securitycenter/settings/apiv1beta1/settingspb"
)
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
func ComponentSettingsObservedState_FromProto(mapCtx *direct.MapContext, in *pb.ComponentSettings) *krm.ComponentSettingsObservedState {
	if in == nil {
		return nil
	}
	out := &krm.ComponentSettingsObservedState{}
	// MISSING: Name
	// MISSING: State
	out.ProjectServiceAccount = direct.LazyPtr(in.GetProjectServiceAccount())
	// MISSING: DetectorSettings
	out.Etag = direct.LazyPtr(in.GetEtag())
	out.UpdateTime = direct.StringTimestamp_FromProto(mapCtx, in.GetUpdateTime())
	// MISSING: ContainerThreatDetectionSettings
	// MISSING: EventThreatDetectionSettings
	// MISSING: SecurityHealthAnalyticsSettings
	// MISSING: WebSecurityScannerSettings
	return out
}
func ComponentSettingsObservedState_ToProto(mapCtx *direct.MapContext, in *krm.ComponentSettingsObservedState) *pb.ComponentSettings {
	if in == nil {
		return nil
	}
	out := &pb.ComponentSettings{}
	// MISSING: Name
	// MISSING: State
	out.ProjectServiceAccount = direct.ValueOf(in.ProjectServiceAccount)
	// MISSING: DetectorSettings
	out.Etag = direct.ValueOf(in.Etag)
	out.UpdateTime = direct.StringTimestamp_ToProto(mapCtx, in.UpdateTime)
	// MISSING: ContainerThreatDetectionSettings
	// MISSING: EventThreatDetectionSettings
	// MISSING: SecurityHealthAnalyticsSettings
	// MISSING: WebSecurityScannerSettings
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
func SecuritycenterComponentSettingsObservedState_FromProto(mapCtx *direct.MapContext, in *pb.ComponentSettings) *krm.SecuritycenterComponentSettingsObservedState {
	if in == nil {
		return nil
	}
	out := &krm.SecuritycenterComponentSettingsObservedState{}
	// MISSING: Name
	// MISSING: State
	// MISSING: ProjectServiceAccount
	// MISSING: DetectorSettings
	// MISSING: Etag
	// MISSING: UpdateTime
	// MISSING: ContainerThreatDetectionSettings
	// MISSING: EventThreatDetectionSettings
	// MISSING: SecurityHealthAnalyticsSettings
	// MISSING: WebSecurityScannerSettings
	return out
}
func SecuritycenterComponentSettingsObservedState_ToProto(mapCtx *direct.MapContext, in *krm.SecuritycenterComponentSettingsObservedState) *pb.ComponentSettings {
	if in == nil {
		return nil
	}
	out := &pb.ComponentSettings{}
	// MISSING: Name
	// MISSING: State
	// MISSING: ProjectServiceAccount
	// MISSING: DetectorSettings
	// MISSING: Etag
	// MISSING: UpdateTime
	// MISSING: ContainerThreatDetectionSettings
	// MISSING: EventThreatDetectionSettings
	// MISSING: SecurityHealthAnalyticsSettings
	// MISSING: WebSecurityScannerSettings
	return out
}
func SecuritycenterComponentSettingsSpec_FromProto(mapCtx *direct.MapContext, in *pb.ComponentSettings) *krm.SecuritycenterComponentSettingsSpec {
	if in == nil {
		return nil
	}
	out := &krm.SecuritycenterComponentSettingsSpec{}
	// MISSING: Name
	// MISSING: State
	// MISSING: ProjectServiceAccount
	// MISSING: DetectorSettings
	// MISSING: Etag
	// MISSING: UpdateTime
	// MISSING: ContainerThreatDetectionSettings
	// MISSING: EventThreatDetectionSettings
	// MISSING: SecurityHealthAnalyticsSettings
	// MISSING: WebSecurityScannerSettings
	return out
}
func SecuritycenterComponentSettingsSpec_ToProto(mapCtx *direct.MapContext, in *krm.SecuritycenterComponentSettingsSpec) *pb.ComponentSettings {
	if in == nil {
		return nil
	}
	out := &pb.ComponentSettings{}
	// MISSING: Name
	// MISSING: State
	// MISSING: ProjectServiceAccount
	// MISSING: DetectorSettings
	// MISSING: Etag
	// MISSING: UpdateTime
	// MISSING: ContainerThreatDetectionSettings
	// MISSING: EventThreatDetectionSettings
	// MISSING: SecurityHealthAnalyticsSettings
	// MISSING: WebSecurityScannerSettings
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
