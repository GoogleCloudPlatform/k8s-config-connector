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

package monitoring

import (
	refs "github.com/GoogleCloudPlatform/k8s-config-connector/apis/refs/v1beta1"
	pb "cloud.google.com/go/monitoring/apiv3/v2/monitoringpb"
	krm "github.com/GoogleCloudPlatform/k8s-config-connector/apis/monitoring/v1alpha1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct"
)
func InternalChecker_FromProto(mapCtx *direct.MapContext, in *pb.InternalChecker) *krm.InternalChecker {
	if in == nil {
		return nil
	}
	out := &krm.InternalChecker{}
	out.Name = direct.LazyPtr(in.GetName())
	out.DisplayName = direct.LazyPtr(in.GetDisplayName())
	out.Network = direct.LazyPtr(in.GetNetwork())
	out.GcpZone = direct.LazyPtr(in.GetGcpZone())
	out.PeerProjectID = direct.LazyPtr(in.GetPeerProjectId())
	out.State = direct.Enum_FromProto(mapCtx, in.GetState())
	return out
}
func InternalChecker_ToProto(mapCtx *direct.MapContext, in *krm.InternalChecker) *pb.InternalChecker {
	if in == nil {
		return nil
	}
	out := &pb.InternalChecker{}
	out.Name = direct.ValueOf(in.Name)
	out.DisplayName = direct.ValueOf(in.DisplayName)
	out.Network = direct.ValueOf(in.Network)
	out.GcpZone = direct.ValueOf(in.GcpZone)
	out.PeerProjectId = direct.ValueOf(in.PeerProjectID)
	out.State = direct.Enum_ToProto[pb.InternalChecker_State](mapCtx, in.State)
	return out
}
func MonitoringUptimeCheckConfigObservedState_FromProto(mapCtx *direct.MapContext, in *pb.UptimeCheckConfig) *krm.MonitoringUptimeCheckConfigObservedState {
	if in == nil {
		return nil
	}
	out := &krm.MonitoringUptimeCheckConfigObservedState{}
	// MISSING: Name
	// MISSING: DisplayName
	// MISSING: MonitoredResource
	// MISSING: ResourceGroup
	// MISSING: SyntheticMonitor
	// MISSING: HTTPCheck
	// MISSING: TcpCheck
	// MISSING: Period
	// MISSING: Timeout
	// MISSING: ContentMatchers
	// MISSING: CheckerType
	// MISSING: SelectedRegions
	// MISSING: IsInternal
	// MISSING: InternalCheckers
	// MISSING: UserLabels
	return out
}
func MonitoringUptimeCheckConfigObservedState_ToProto(mapCtx *direct.MapContext, in *krm.MonitoringUptimeCheckConfigObservedState) *pb.UptimeCheckConfig {
	if in == nil {
		return nil
	}
	out := &pb.UptimeCheckConfig{}
	// MISSING: Name
	// MISSING: DisplayName
	// MISSING: MonitoredResource
	// MISSING: ResourceGroup
	// MISSING: SyntheticMonitor
	// MISSING: HTTPCheck
	// MISSING: TcpCheck
	// MISSING: Period
	// MISSING: Timeout
	// MISSING: ContentMatchers
	// MISSING: CheckerType
	// MISSING: SelectedRegions
	// MISSING: IsInternal
	// MISSING: InternalCheckers
	// MISSING: UserLabels
	return out
}
func MonitoringUptimeCheckConfigSpec_FromProto(mapCtx *direct.MapContext, in *pb.UptimeCheckConfig) *krm.MonitoringUptimeCheckConfigSpec {
	if in == nil {
		return nil
	}
	out := &krm.MonitoringUptimeCheckConfigSpec{}
	// MISSING: Name
	// MISSING: DisplayName
	// MISSING: MonitoredResource
	// MISSING: ResourceGroup
	// MISSING: SyntheticMonitor
	// MISSING: HTTPCheck
	// MISSING: TcpCheck
	// MISSING: Period
	// MISSING: Timeout
	// MISSING: ContentMatchers
	// MISSING: CheckerType
	// MISSING: SelectedRegions
	// MISSING: IsInternal
	// MISSING: InternalCheckers
	// MISSING: UserLabels
	return out
}
func MonitoringUptimeCheckConfigSpec_ToProto(mapCtx *direct.MapContext, in *krm.MonitoringUptimeCheckConfigSpec) *pb.UptimeCheckConfig {
	if in == nil {
		return nil
	}
	out := &pb.UptimeCheckConfig{}
	// MISSING: Name
	// MISSING: DisplayName
	// MISSING: MonitoredResource
	// MISSING: ResourceGroup
	// MISSING: SyntheticMonitor
	// MISSING: HTTPCheck
	// MISSING: TcpCheck
	// MISSING: Period
	// MISSING: Timeout
	// MISSING: ContentMatchers
	// MISSING: CheckerType
	// MISSING: SelectedRegions
	// MISSING: IsInternal
	// MISSING: InternalCheckers
	// MISSING: UserLabels
	return out
}
func SyntheticMonitorTarget_FromProto(mapCtx *direct.MapContext, in *pb.SyntheticMonitorTarget) *krm.SyntheticMonitorTarget {
	if in == nil {
		return nil
	}
	out := &krm.SyntheticMonitorTarget{}
	out.CloudFunctionV2 = SyntheticMonitorTarget_CloudFunctionV2Target_FromProto(mapCtx, in.GetCloudFunctionV2())
	return out
}
func SyntheticMonitorTarget_ToProto(mapCtx *direct.MapContext, in *krm.SyntheticMonitorTarget) *pb.SyntheticMonitorTarget {
	if in == nil {
		return nil
	}
	out := &pb.SyntheticMonitorTarget{}
	if oneof := SyntheticMonitorTarget_CloudFunctionV2Target_ToProto(mapCtx, in.CloudFunctionV2); oneof != nil {
		out.Target = &pb.SyntheticMonitorTarget_CloudFunctionV2{CloudFunctionV2: oneof}
	}
	return out
}
func SyntheticMonitorTargetObservedState_FromProto(mapCtx *direct.MapContext, in *pb.SyntheticMonitorTarget) *krm.SyntheticMonitorTargetObservedState {
	if in == nil {
		return nil
	}
	out := &krm.SyntheticMonitorTargetObservedState{}
	out.CloudFunctionV2 = SyntheticMonitorTarget_CloudFunctionV2TargetObservedState_FromProto(mapCtx, in.GetCloudFunctionV2())
	return out
}
func SyntheticMonitorTargetObservedState_ToProto(mapCtx *direct.MapContext, in *krm.SyntheticMonitorTargetObservedState) *pb.SyntheticMonitorTarget {
	if in == nil {
		return nil
	}
	out := &pb.SyntheticMonitorTarget{}
	if oneof := SyntheticMonitorTarget_CloudFunctionV2TargetObservedState_ToProto(mapCtx, in.CloudFunctionV2); oneof != nil {
		out.Target = &pb.SyntheticMonitorTarget_CloudFunctionV2{CloudFunctionV2: oneof}
	}
	return out
}
func SyntheticMonitorTarget_CloudFunctionV2Target_FromProto(mapCtx *direct.MapContext, in *pb.SyntheticMonitorTarget_CloudFunctionV2Target) *krm.SyntheticMonitorTarget_CloudFunctionV2Target {
	if in == nil {
		return nil
	}
	out := &krm.SyntheticMonitorTarget_CloudFunctionV2Target{}
	out.Name = direct.LazyPtr(in.GetName())
	// MISSING: CloudRunRevision
	return out
}
func SyntheticMonitorTarget_CloudFunctionV2Target_ToProto(mapCtx *direct.MapContext, in *krm.SyntheticMonitorTarget_CloudFunctionV2Target) *pb.SyntheticMonitorTarget_CloudFunctionV2Target {
	if in == nil {
		return nil
	}
	out := &pb.SyntheticMonitorTarget_CloudFunctionV2Target{}
	out.Name = direct.ValueOf(in.Name)
	// MISSING: CloudRunRevision
	return out
}
func SyntheticMonitorTarget_CloudFunctionV2TargetObservedState_FromProto(mapCtx *direct.MapContext, in *pb.SyntheticMonitorTarget_CloudFunctionV2Target) *krm.SyntheticMonitorTarget_CloudFunctionV2TargetObservedState {
	if in == nil {
		return nil
	}
	out := &krm.SyntheticMonitorTarget_CloudFunctionV2TargetObservedState{}
	// MISSING: Name
	out.CloudRunRevision = MonitoredResource_FromProto(mapCtx, in.GetCloudRunRevision())
	return out
}
func SyntheticMonitorTarget_CloudFunctionV2TargetObservedState_ToProto(mapCtx *direct.MapContext, in *krm.SyntheticMonitorTarget_CloudFunctionV2TargetObservedState) *pb.SyntheticMonitorTarget_CloudFunctionV2Target {
	if in == nil {
		return nil
	}
	out := &pb.SyntheticMonitorTarget_CloudFunctionV2Target{}
	// MISSING: Name
	out.CloudRunRevision = MonitoredResource_ToProto(mapCtx, in.CloudRunRevision)
	return out
}
func UptimeCheckConfig_FromProto(mapCtx *direct.MapContext, in *pb.UptimeCheckConfig) *krm.UptimeCheckConfig {
	if in == nil {
		return nil
	}
	out := &krm.UptimeCheckConfig{}
	out.Name = direct.LazyPtr(in.GetName())
	out.DisplayName = direct.LazyPtr(in.GetDisplayName())
	out.MonitoredResource = MonitoredResource_FromProto(mapCtx, in.GetMonitoredResource())
	out.ResourceGroup = UptimeCheckConfig_ResourceGroup_FromProto(mapCtx, in.GetResourceGroup())
	out.SyntheticMonitor = SyntheticMonitorTarget_FromProto(mapCtx, in.GetSyntheticMonitor())
	out.HTTPCheck = UptimeCheckConfig_HttpCheck_FromProto(mapCtx, in.GetHttpCheck())
	out.TcpCheck = UptimeCheckConfig_TcpCheck_FromProto(mapCtx, in.GetTcpCheck())
	out.Period = direct.StringDuration_FromProto(mapCtx, in.GetPeriod())
	out.Timeout = direct.StringDuration_FromProto(mapCtx, in.GetTimeout())
	out.ContentMatchers = direct.Slice_FromProto(mapCtx, in.ContentMatchers, UptimeCheckConfig_ContentMatcher_FromProto)
	out.CheckerType = direct.Enum_FromProto(mapCtx, in.GetCheckerType())
	out.SelectedRegions = direct.EnumSlice_FromProto(mapCtx, in.SelectedRegions)
	out.IsInternal = direct.LazyPtr(in.GetIsInternal())
	out.InternalCheckers = direct.Slice_FromProto(mapCtx, in.InternalCheckers, InternalChecker_FromProto)
	out.UserLabels = in.UserLabels
	return out
}
func UptimeCheckConfig_ToProto(mapCtx *direct.MapContext, in *krm.UptimeCheckConfig) *pb.UptimeCheckConfig {
	if in == nil {
		return nil
	}
	out := &pb.UptimeCheckConfig{}
	out.Name = direct.ValueOf(in.Name)
	out.DisplayName = direct.ValueOf(in.DisplayName)
	if oneof := MonitoredResource_ToProto(mapCtx, in.MonitoredResource); oneof != nil {
		out.Resource = &pb.UptimeCheckConfig_MonitoredResource{MonitoredResource: oneof}
	}
	if oneof := UptimeCheckConfig_ResourceGroup_ToProto(mapCtx, in.ResourceGroup); oneof != nil {
		out.Resource = &pb.UptimeCheckConfig_ResourceGroup_{ResourceGroup: oneof}
	}
	if oneof := SyntheticMonitorTarget_ToProto(mapCtx, in.SyntheticMonitor); oneof != nil {
		out.Resource = &pb.UptimeCheckConfig_SyntheticMonitor{SyntheticMonitor: oneof}
	}
	if oneof := UptimeCheckConfig_HttpCheck_ToProto(mapCtx, in.HTTPCheck); oneof != nil {
		out.CheckRequestType = &pb.UptimeCheckConfig_HttpCheck_{HttpCheck: oneof}
	}
	if oneof := UptimeCheckConfig_TcpCheck_ToProto(mapCtx, in.TcpCheck); oneof != nil {
		out.CheckRequestType = &pb.UptimeCheckConfig_TcpCheck_{TcpCheck: oneof}
	}
	out.Period = direct.StringDuration_ToProto(mapCtx, in.Period)
	out.Timeout = direct.StringDuration_ToProto(mapCtx, in.Timeout)
	out.ContentMatchers = direct.Slice_ToProto(mapCtx, in.ContentMatchers, UptimeCheckConfig_ContentMatcher_ToProto)
	out.CheckerType = direct.Enum_ToProto[pb.UptimeCheckConfig_CheckerType](mapCtx, in.CheckerType)
	out.SelectedRegions = direct.EnumSlice_ToProto[pb.UptimeCheckRegion](mapCtx, in.SelectedRegions)
	out.IsInternal = direct.ValueOf(in.IsInternal)
	out.InternalCheckers = direct.Slice_ToProto(mapCtx, in.InternalCheckers, InternalChecker_ToProto)
	out.UserLabels = in.UserLabels
	return out
}
func UptimeCheckConfigObservedState_FromProto(mapCtx *direct.MapContext, in *pb.UptimeCheckConfig) *krm.UptimeCheckConfigObservedState {
	if in == nil {
		return nil
	}
	out := &krm.UptimeCheckConfigObservedState{}
	// MISSING: Name
	// MISSING: DisplayName
	// MISSING: MonitoredResource
	// MISSING: ResourceGroup
	out.SyntheticMonitor = SyntheticMonitorTargetObservedState_FromProto(mapCtx, in.GetSyntheticMonitor())
	// MISSING: HTTPCheck
	// MISSING: TcpCheck
	// MISSING: Period
	// MISSING: Timeout
	// MISSING: ContentMatchers
	// MISSING: CheckerType
	// MISSING: SelectedRegions
	// MISSING: IsInternal
	// MISSING: InternalCheckers
	// MISSING: UserLabels
	return out
}
func UptimeCheckConfigObservedState_ToProto(mapCtx *direct.MapContext, in *krm.UptimeCheckConfigObservedState) *pb.UptimeCheckConfig {
	if in == nil {
		return nil
	}
	out := &pb.UptimeCheckConfig{}
	// MISSING: Name
	// MISSING: DisplayName
	// MISSING: MonitoredResource
	// MISSING: ResourceGroup
	if oneof := SyntheticMonitorTargetObservedState_ToProto(mapCtx, in.SyntheticMonitor); oneof != nil {
		out.Resource = &pb.UptimeCheckConfig_SyntheticMonitor{SyntheticMonitor: oneof}
	}
	// MISSING: HTTPCheck
	// MISSING: TcpCheck
	// MISSING: Period
	// MISSING: Timeout
	// MISSING: ContentMatchers
	// MISSING: CheckerType
	// MISSING: SelectedRegions
	// MISSING: IsInternal
	// MISSING: InternalCheckers
	// MISSING: UserLabels
	return out
}
func UptimeCheckConfig_ContentMatcher_FromProto(mapCtx *direct.MapContext, in *pb.UptimeCheckConfig_ContentMatcher) *krm.UptimeCheckConfig_ContentMatcher {
	if in == nil {
		return nil
	}
	out := &krm.UptimeCheckConfig_ContentMatcher{}
	out.Content = direct.LazyPtr(in.GetContent())
	out.Matcher = direct.Enum_FromProto(mapCtx, in.GetMatcher())
	out.JsonPathMatcher = UptimeCheckConfig_ContentMatcher_JsonPathMatcher_FromProto(mapCtx, in.GetJsonPathMatcher())
	return out
}
func UptimeCheckConfig_ContentMatcher_ToProto(mapCtx *direct.MapContext, in *krm.UptimeCheckConfig_ContentMatcher) *pb.UptimeCheckConfig_ContentMatcher {
	if in == nil {
		return nil
	}
	out := &pb.UptimeCheckConfig_ContentMatcher{}
	out.Content = direct.ValueOf(in.Content)
	out.Matcher = direct.Enum_ToProto[pb.UptimeCheckConfig_ContentMatcher_ContentMatcherOption](mapCtx, in.Matcher)
	if oneof := UptimeCheckConfig_ContentMatcher_JsonPathMatcher_ToProto(mapCtx, in.JsonPathMatcher); oneof != nil {
		out.AdditionalMatcherInfo = &pb.UptimeCheckConfig_ContentMatcher_JsonPathMatcher_{JsonPathMatcher: oneof}
	}
	return out
}
func UptimeCheckConfig_ContentMatcher_JsonPathMatcher_FromProto(mapCtx *direct.MapContext, in *pb.UptimeCheckConfig_ContentMatcher_JsonPathMatcher) *krm.UptimeCheckConfig_ContentMatcher_JsonPathMatcher {
	if in == nil {
		return nil
	}
	out := &krm.UptimeCheckConfig_ContentMatcher_JsonPathMatcher{}
	out.JsonPath = direct.LazyPtr(in.GetJsonPath())
	out.JsonMatcher = direct.Enum_FromProto(mapCtx, in.GetJsonMatcher())
	return out
}
func UptimeCheckConfig_ContentMatcher_JsonPathMatcher_ToProto(mapCtx *direct.MapContext, in *krm.UptimeCheckConfig_ContentMatcher_JsonPathMatcher) *pb.UptimeCheckConfig_ContentMatcher_JsonPathMatcher {
	if in == nil {
		return nil
	}
	out := &pb.UptimeCheckConfig_ContentMatcher_JsonPathMatcher{}
	out.JsonPath = direct.ValueOf(in.JsonPath)
	out.JsonMatcher = direct.Enum_ToProto[pb.UptimeCheckConfig_ContentMatcher_JsonPathMatcher_JsonPathMatcherOption](mapCtx, in.JsonMatcher)
	return out
}
func UptimeCheckConfig_HttpCheck_FromProto(mapCtx *direct.MapContext, in *pb.UptimeCheckConfig_HttpCheck) *krm.UptimeCheckConfig_HttpCheck {
	if in == nil {
		return nil
	}
	out := &krm.UptimeCheckConfig_HttpCheck{}
	out.RequestMethod = direct.Enum_FromProto(mapCtx, in.GetRequestMethod())
	out.UseSsl = direct.LazyPtr(in.GetUseSsl())
	out.Path = direct.LazyPtr(in.GetPath())
	out.Port = direct.LazyPtr(in.GetPort())
	out.AuthInfo = UptimeCheckConfig_HttpCheck_BasicAuthentication_FromProto(mapCtx, in.GetAuthInfo())
	out.MaskHeaders = direct.LazyPtr(in.GetMaskHeaders())
	out.Headers = in.Headers
	out.ContentType = direct.Enum_FromProto(mapCtx, in.GetContentType())
	out.CustomContentType = direct.LazyPtr(in.GetCustomContentType())
	out.ValidateSsl = direct.LazyPtr(in.GetValidateSsl())
	out.Body = in.GetBody()
	out.AcceptedResponseStatusCodes = direct.Slice_FromProto(mapCtx, in.AcceptedResponseStatusCodes, UptimeCheckConfig_HttpCheck_ResponseStatusCode_FromProto)
	out.PingConfig = UptimeCheckConfig_PingConfig_FromProto(mapCtx, in.GetPingConfig())
	out.ServiceAgentAuthentication = UptimeCheckConfig_HttpCheck_ServiceAgentAuthentication_FromProto(mapCtx, in.GetServiceAgentAuthentication())
	return out
}
func UptimeCheckConfig_HttpCheck_ToProto(mapCtx *direct.MapContext, in *krm.UptimeCheckConfig_HttpCheck) *pb.UptimeCheckConfig_HttpCheck {
	if in == nil {
		return nil
	}
	out := &pb.UptimeCheckConfig_HttpCheck{}
	out.RequestMethod = direct.Enum_ToProto[pb.UptimeCheckConfig_HttpCheck_RequestMethod](mapCtx, in.RequestMethod)
	out.UseSsl = direct.ValueOf(in.UseSsl)
	out.Path = direct.ValueOf(in.Path)
	out.Port = direct.ValueOf(in.Port)
	out.AuthInfo = UptimeCheckConfig_HttpCheck_BasicAuthentication_ToProto(mapCtx, in.AuthInfo)
	out.MaskHeaders = direct.ValueOf(in.MaskHeaders)
	out.Headers = in.Headers
	out.ContentType = direct.Enum_ToProto[pb.UptimeCheckConfig_HttpCheck_ContentType](mapCtx, in.ContentType)
	out.CustomContentType = direct.ValueOf(in.CustomContentType)
	out.ValidateSsl = direct.ValueOf(in.ValidateSsl)
	out.Body = in.Body
	out.AcceptedResponseStatusCodes = direct.Slice_ToProto(mapCtx, in.AcceptedResponseStatusCodes, UptimeCheckConfig_HttpCheck_ResponseStatusCode_ToProto)
	out.PingConfig = UptimeCheckConfig_PingConfig_ToProto(mapCtx, in.PingConfig)
	if oneof := UptimeCheckConfig_HttpCheck_ServiceAgentAuthentication_ToProto(mapCtx, in.ServiceAgentAuthentication); oneof != nil {
		out.AuthMethod = &pb.UptimeCheckConfig_HttpCheck_ServiceAgentAuthentication_{ServiceAgentAuthentication: oneof}
	}
	return out
}
func UptimeCheckConfig_HttpCheck_BasicAuthentication_FromProto(mapCtx *direct.MapContext, in *pb.UptimeCheckConfig_HttpCheck_BasicAuthentication) *krm.UptimeCheckConfig_HttpCheck_BasicAuthentication {
	if in == nil {
		return nil
	}
	out := &krm.UptimeCheckConfig_HttpCheck_BasicAuthentication{}
	out.Username = direct.LazyPtr(in.GetUsername())
	out.Password = direct.LazyPtr(in.GetPassword())
	return out
}
func UptimeCheckConfig_HttpCheck_BasicAuthentication_ToProto(mapCtx *direct.MapContext, in *krm.UptimeCheckConfig_HttpCheck_BasicAuthentication) *pb.UptimeCheckConfig_HttpCheck_BasicAuthentication {
	if in == nil {
		return nil
	}
	out := &pb.UptimeCheckConfig_HttpCheck_BasicAuthentication{}
	out.Username = direct.ValueOf(in.Username)
	out.Password = direct.ValueOf(in.Password)
	return out
}
func UptimeCheckConfig_HttpCheck_ResponseStatusCode_FromProto(mapCtx *direct.MapContext, in *pb.UptimeCheckConfig_HttpCheck_ResponseStatusCode) *krm.UptimeCheckConfig_HttpCheck_ResponseStatusCode {
	if in == nil {
		return nil
	}
	out := &krm.UptimeCheckConfig_HttpCheck_ResponseStatusCode{}
	out.StatusValue = direct.LazyPtr(in.GetStatusValue())
	out.StatusClass = direct.Enum_FromProto(mapCtx, in.GetStatusClass())
	return out
}
func UptimeCheckConfig_HttpCheck_ResponseStatusCode_ToProto(mapCtx *direct.MapContext, in *krm.UptimeCheckConfig_HttpCheck_ResponseStatusCode) *pb.UptimeCheckConfig_HttpCheck_ResponseStatusCode {
	if in == nil {
		return nil
	}
	out := &pb.UptimeCheckConfig_HttpCheck_ResponseStatusCode{}
	if oneof := UptimeCheckConfig_HttpCheck_ResponseStatusCode_StatusValue_ToProto(mapCtx, in.StatusValue); oneof != nil {
		out.StatusCode = oneof
	}
	if oneof := UptimeCheckConfig_HttpCheck_ResponseStatusCode_StatusClass_ToProto(mapCtx, in.StatusClass); oneof != nil {
		out.StatusCode = oneof
	}
	return out
}
func UptimeCheckConfig_HttpCheck_ServiceAgentAuthentication_FromProto(mapCtx *direct.MapContext, in *pb.UptimeCheckConfig_HttpCheck_ServiceAgentAuthentication) *krm.UptimeCheckConfig_HttpCheck_ServiceAgentAuthentication {
	if in == nil {
		return nil
	}
	out := &krm.UptimeCheckConfig_HttpCheck_ServiceAgentAuthentication{}
	out.Type = direct.Enum_FromProto(mapCtx, in.GetType())
	return out
}
func UptimeCheckConfig_HttpCheck_ServiceAgentAuthentication_ToProto(mapCtx *direct.MapContext, in *krm.UptimeCheckConfig_HttpCheck_ServiceAgentAuthentication) *pb.UptimeCheckConfig_HttpCheck_ServiceAgentAuthentication {
	if in == nil {
		return nil
	}
	out := &pb.UptimeCheckConfig_HttpCheck_ServiceAgentAuthentication{}
	out.Type = direct.Enum_ToProto[pb.UptimeCheckConfig_HttpCheck_ServiceAgentAuthentication_ServiceAgentAuthenticationType](mapCtx, in.Type)
	return out
}
func UptimeCheckConfig_PingConfig_FromProto(mapCtx *direct.MapContext, in *pb.UptimeCheckConfig_PingConfig) *krm.UptimeCheckConfig_PingConfig {
	if in == nil {
		return nil
	}
	out := &krm.UptimeCheckConfig_PingConfig{}
	out.PingsCount = direct.LazyPtr(in.GetPingsCount())
	return out
}
func UptimeCheckConfig_PingConfig_ToProto(mapCtx *direct.MapContext, in *krm.UptimeCheckConfig_PingConfig) *pb.UptimeCheckConfig_PingConfig {
	if in == nil {
		return nil
	}
	out := &pb.UptimeCheckConfig_PingConfig{}
	out.PingsCount = direct.ValueOf(in.PingsCount)
	return out
}
func UptimeCheckConfig_ResourceGroup_FromProto(mapCtx *direct.MapContext, in *pb.UptimeCheckConfig_ResourceGroup) *krm.UptimeCheckConfig_ResourceGroup {
	if in == nil {
		return nil
	}
	out := &krm.UptimeCheckConfig_ResourceGroup{}
	out.GroupID = direct.LazyPtr(in.GetGroupId())
	out.ResourceType = direct.Enum_FromProto(mapCtx, in.GetResourceType())
	return out
}
func UptimeCheckConfig_ResourceGroup_ToProto(mapCtx *direct.MapContext, in *krm.UptimeCheckConfig_ResourceGroup) *pb.UptimeCheckConfig_ResourceGroup {
	if in == nil {
		return nil
	}
	out := &pb.UptimeCheckConfig_ResourceGroup{}
	out.GroupId = direct.ValueOf(in.GroupID)
	out.ResourceType = direct.Enum_ToProto[pb.GroupResourceType](mapCtx, in.ResourceType)
	return out
}
func UptimeCheckConfig_TcpCheck_FromProto(mapCtx *direct.MapContext, in *pb.UptimeCheckConfig_TcpCheck) *krm.UptimeCheckConfig_TcpCheck {
	if in == nil {
		return nil
	}
	out := &krm.UptimeCheckConfig_TcpCheck{}
	out.Port = direct.LazyPtr(in.GetPort())
	out.PingConfig = UptimeCheckConfig_PingConfig_FromProto(mapCtx, in.GetPingConfig())
	return out
}
func UptimeCheckConfig_TcpCheck_ToProto(mapCtx *direct.MapContext, in *krm.UptimeCheckConfig_TcpCheck) *pb.UptimeCheckConfig_TcpCheck {
	if in == nil {
		return nil
	}
	out := &pb.UptimeCheckConfig_TcpCheck{}
	out.Port = direct.ValueOf(in.Port)
	out.PingConfig = UptimeCheckConfig_PingConfig_ToProto(mapCtx, in.PingConfig)
	return out
}
