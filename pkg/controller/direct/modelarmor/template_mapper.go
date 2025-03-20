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

package modelarmor

import (
	pb "cloud.google.com/go/modelarmor/apiv1/modelarmorpb"
	krm "github.com/GoogleCloudPlatform/k8s-config-connector/apis/modelarmor/v1alpha1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct"
)

func FilterConfig_FromProto(mapCtx *direct.MapContext, in *pb.FilterConfig) *krm.FilterConfig {
	if in == nil {
		return nil
	}
	out := &krm.FilterConfig{}
	out.RaiSettings = RaiFilterSettings_FromProto(mapCtx, in.GetRaiSettings())
	out.SdpSettings = SdpFilterSettings_FromProto(mapCtx, in.GetSdpSettings())
	out.PiAndJailbreakFilterSettings = PiAndJailbreakFilterSettings_FromProto(mapCtx, in.GetPiAndJailbreakFilterSettings())
	out.MaliciousURIFilterSettings = MaliciousURIFilterSettings_FromProto(mapCtx, in.GetMaliciousUriFilterSettings())
	return out
}
func FilterConfig_ToProto(mapCtx *direct.MapContext, in *krm.FilterConfig) *pb.FilterConfig {
	if in == nil {
		return nil
	}
	out := &pb.FilterConfig{}
	out.RaiSettings = RaiFilterSettings_ToProto(mapCtx, in.RaiSettings)
	out.SdpSettings = SdpFilterSettings_ToProto(mapCtx, in.SdpSettings)
	out.PiAndJailbreakFilterSettings = PiAndJailbreakFilterSettings_ToProto(mapCtx, in.PiAndJailbreakFilterSettings)
	out.MaliciousUriFilterSettings = MaliciousURIFilterSettings_ToProto(mapCtx, in.MaliciousURIFilterSettings)
	return out
}
func MaliciousURIFilterSettings_FromProto(mapCtx *direct.MapContext, in *pb.MaliciousUriFilterSettings) *krm.MaliciousURIFilterSettings {
	if in == nil {
		return nil
	}
	out := &krm.MaliciousURIFilterSettings{}
	out.FilterEnforcement = direct.Enum_FromProto(mapCtx, in.GetFilterEnforcement())
	return out
}
func MaliciousURIFilterSettings_ToProto(mapCtx *direct.MapContext, in *krm.MaliciousURIFilterSettings) *pb.MaliciousUriFilterSettings {
	if in == nil {
		return nil
	}
	out := &pb.MaliciousUriFilterSettings{}
	out.FilterEnforcement = direct.Enum_ToProto[pb.MaliciousUriFilterSettings_MaliciousUriFilterEnforcement](mapCtx, in.FilterEnforcement)
	return out
}
func PiAndJailbreakFilterSettings_FromProto(mapCtx *direct.MapContext, in *pb.PiAndJailbreakFilterSettings) *krm.PiAndJailbreakFilterSettings {
	if in == nil {
		return nil
	}
	out := &krm.PiAndJailbreakFilterSettings{}
	out.FilterEnforcement = direct.Enum_FromProto(mapCtx, in.GetFilterEnforcement())
	out.ConfidenceLevel = direct.Enum_FromProto(mapCtx, in.GetConfidenceLevel())
	return out
}
func PiAndJailbreakFilterSettings_ToProto(mapCtx *direct.MapContext, in *krm.PiAndJailbreakFilterSettings) *pb.PiAndJailbreakFilterSettings {
	if in == nil {
		return nil
	}
	out := &pb.PiAndJailbreakFilterSettings{}
	out.FilterEnforcement = direct.Enum_ToProto[pb.PiAndJailbreakFilterSettings_PiAndJailbreakFilterEnforcement](mapCtx, in.FilterEnforcement)
	out.ConfidenceLevel = direct.Enum_ToProto[pb.DetectionConfidenceLevel](mapCtx, in.ConfidenceLevel)
	return out
}
func RaiFilterSettings_FromProto(mapCtx *direct.MapContext, in *pb.RaiFilterSettings) *krm.RaiFilterSettings {
	if in == nil {
		return nil
	}
	out := &krm.RaiFilterSettings{}
	out.RaiFilters = direct.Slice_FromProto(mapCtx, in.RaiFilters, RaiFilterSettings_RaiFilter_FromProto)
	return out
}
func RaiFilterSettings_ToProto(mapCtx *direct.MapContext, in *krm.RaiFilterSettings) *pb.RaiFilterSettings {
	if in == nil {
		return nil
	}
	out := &pb.RaiFilterSettings{}
	out.RaiFilters = direct.Slice_ToProto(mapCtx, in.RaiFilters, RaiFilterSettings_RaiFilter_ToProto)
	return out
}
func RaiFilterSettings_RaiFilter_FromProto(mapCtx *direct.MapContext, in *pb.RaiFilterSettings_RaiFilter) *krm.RaiFilterSettings_RaiFilter {
	if in == nil {
		return nil
	}
	out := &krm.RaiFilterSettings_RaiFilter{}
	out.FilterType = direct.Enum_FromProto(mapCtx, in.GetFilterType())
	out.ConfidenceLevel = direct.Enum_FromProto(mapCtx, in.GetConfidenceLevel())
	return out
}
func RaiFilterSettings_RaiFilter_ToProto(mapCtx *direct.MapContext, in *krm.RaiFilterSettings_RaiFilter) *pb.RaiFilterSettings_RaiFilter {
	if in == nil {
		return nil
	}
	out := &pb.RaiFilterSettings_RaiFilter{}
	out.FilterType = direct.Enum_ToProto[pb.RaiFilterType](mapCtx, in.FilterType)
	out.ConfidenceLevel = direct.Enum_ToProto[pb.DetectionConfidenceLevel](mapCtx, in.ConfidenceLevel)
	return out
}
func SdpAdvancedConfig_FromProto(mapCtx *direct.MapContext, in *pb.SdpAdvancedConfig) *krm.SdpAdvancedConfig {
	if in == nil {
		return nil
	}
	out := &krm.SdpAdvancedConfig{}
	out.InspectTemplate = direct.LazyPtr(in.GetInspectTemplate())
	out.DeidentifyTemplate = direct.LazyPtr(in.GetDeidentifyTemplate())
	return out
}
func SdpAdvancedConfig_ToProto(mapCtx *direct.MapContext, in *krm.SdpAdvancedConfig) *pb.SdpAdvancedConfig {
	if in == nil {
		return nil
	}
	out := &pb.SdpAdvancedConfig{}
	out.InspectTemplate = direct.ValueOf(in.InspectTemplate)
	out.DeidentifyTemplate = direct.ValueOf(in.DeidentifyTemplate)
	return out
}
func SdpBasicConfig_FromProto(mapCtx *direct.MapContext, in *pb.SdpBasicConfig) *krm.SdpBasicConfig {
	if in == nil {
		return nil
	}
	out := &krm.SdpBasicConfig{}
	out.FilterEnforcement = direct.Enum_FromProto(mapCtx, in.GetFilterEnforcement())
	return out
}
func SdpBasicConfig_ToProto(mapCtx *direct.MapContext, in *krm.SdpBasicConfig) *pb.SdpBasicConfig {
	if in == nil {
		return nil
	}
	out := &pb.SdpBasicConfig{}
	out.FilterEnforcement = direct.Enum_ToProto[pb.SdpBasicConfig_SdpBasicConfigEnforcement](mapCtx, in.FilterEnforcement)
	return out
}
func SdpFilterSettings_FromProto(mapCtx *direct.MapContext, in *pb.SdpFilterSettings) *krm.SdpFilterSettings {
	if in == nil {
		return nil
	}
	out := &krm.SdpFilterSettings{}
	out.BasicConfig = SdpBasicConfig_FromProto(mapCtx, in.GetBasicConfig())
	out.AdvancedConfig = SdpAdvancedConfig_FromProto(mapCtx, in.GetAdvancedConfig())
	return out
}
func SdpFilterSettings_ToProto(mapCtx *direct.MapContext, in *krm.SdpFilterSettings) *pb.SdpFilterSettings {
	if in == nil {
		return nil
	}
	out := &pb.SdpFilterSettings{}
	if oneof := SdpBasicConfig_ToProto(mapCtx, in.BasicConfig); oneof != nil {
		out.SdpConfiguration = &pb.SdpFilterSettings_BasicConfig{BasicConfig: oneof}
	}
	if oneof := SdpAdvancedConfig_ToProto(mapCtx, in.AdvancedConfig); oneof != nil {
		out.SdpConfiguration = &pb.SdpFilterSettings_AdvancedConfig{AdvancedConfig: oneof}
	}
	return out
}
func ModelArmorTemplateSpec_FromProto(mapCtx *direct.MapContext, in *pb.Template) *krm.ModelArmorTemplateSpec {
	if in == nil {
		return nil
	}
	out := &krm.Template{}
	out.Name = direct.LazyPtr(in.GetName())
	// MISSING: CreateTime
	// MISSING: UpdateTime
	out.Labels = in.Labels
	out.FilterConfig = FilterConfig_FromProto(mapCtx, in.GetFilterConfig())
	out.TemplateMetadata = Template_TemplateMetadata_FromProto(mapCtx, in.GetTemplateMetadata())
	return out
}
func ModelArmorTemplateSpec_ToProto(mapCtx *direct.MapContext, in *krm.ModelArmorTemplateSpec) *pb.Template {
	if in == nil {
		return nil
	}
	out := &pb.Template{}
	out.Name = direct.ValueOf(in.Name)
	// MISSING: CreateTime
	// MISSING: UpdateTime
	out.Labels = in.Labels
	out.FilterConfig = FilterConfig_ToProto(mapCtx, in.FilterConfig)
	out.TemplateMetadata = Template_TemplateMetadata_ToProto(mapCtx, in.TemplateMetadata)
	return out
}
func ModelArmorTemplateObservedState_FromProto(mapCtx *direct.MapContext, in *pb.Template) *krm.ModelArmorTemplateObservedState {
	if in == nil {
		return nil
	}
	out := &krm.ModelArmorTemplateObservedState{}
	// MISSING: Name
	out.CreateTime = direct.StringTimestamp_FromProto(mapCtx, in.GetCreateTime())
	out.UpdateTime = direct.StringTimestamp_FromProto(mapCtx, in.GetUpdateTime())
	// MISSING: Labels
	// MISSING: FilterConfig
	// MISSING: TemplateMetadata
	return out
}
func ModelArmorTemplateObservedState_ToProto(mapCtx *direct.MapContext, in *krm.ModelArmorTemplateObservedState) *pb.Template {
	if in == nil {
		return nil
	}
	out := &pb.Template{}
	// MISSING: Name
	out.CreateTime = direct.StringTimestamp_ToProto(mapCtx, in.CreateTime)
	out.UpdateTime = direct.StringTimestamp_ToProto(mapCtx, in.UpdateTime)
	// MISSING: Labels
	// MISSING: FilterConfig
	// MISSING: TemplateMetadata
	return out
}
