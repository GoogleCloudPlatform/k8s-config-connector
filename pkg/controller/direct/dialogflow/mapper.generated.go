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

package dialogflow

import (
	refs "github.com/GoogleCloudPlatform/k8s-config-connector/apis/refs/v1beta1"
	pb "cloud.google.com/go/dialogflow/cx/apiv3/cxpb"
	krm "github.com/GoogleCloudPlatform/k8s-config-connector/apis/dialogflow/v1alpha1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct"
)
func DialogflowSecuritySettingsObservedState_FromProto(mapCtx *direct.MapContext, in *pb.SecuritySettings) *krm.DialogflowSecuritySettingsObservedState {
	if in == nil {
		return nil
	}
	out := &krm.DialogflowSecuritySettingsObservedState{}
	// MISSING: Name
	// MISSING: DisplayName
	// MISSING: RedactionStrategy
	// MISSING: RedactionScope
	// MISSING: InspectTemplate
	// MISSING: DeidentifyTemplate
	// MISSING: RetentionWindowDays
	// MISSING: RetentionStrategy
	// MISSING: PurgeDataTypes
	// MISSING: AudioExportSettings
	// MISSING: InsightsExportSettings
	return out
}
func DialogflowSecuritySettingsObservedState_ToProto(mapCtx *direct.MapContext, in *krm.DialogflowSecuritySettingsObservedState) *pb.SecuritySettings {
	if in == nil {
		return nil
	}
	out := &pb.SecuritySettings{}
	// MISSING: Name
	// MISSING: DisplayName
	// MISSING: RedactionStrategy
	// MISSING: RedactionScope
	// MISSING: InspectTemplate
	// MISSING: DeidentifyTemplate
	// MISSING: RetentionWindowDays
	// MISSING: RetentionStrategy
	// MISSING: PurgeDataTypes
	// MISSING: AudioExportSettings
	// MISSING: InsightsExportSettings
	return out
}
func DialogflowSecuritySettingsSpec_FromProto(mapCtx *direct.MapContext, in *pb.SecuritySettings) *krm.DialogflowSecuritySettingsSpec {
	if in == nil {
		return nil
	}
	out := &krm.DialogflowSecuritySettingsSpec{}
	// MISSING: Name
	// MISSING: DisplayName
	// MISSING: RedactionStrategy
	// MISSING: RedactionScope
	// MISSING: InspectTemplate
	// MISSING: DeidentifyTemplate
	// MISSING: RetentionWindowDays
	// MISSING: RetentionStrategy
	// MISSING: PurgeDataTypes
	// MISSING: AudioExportSettings
	// MISSING: InsightsExportSettings
	return out
}
func DialogflowSecuritySettingsSpec_ToProto(mapCtx *direct.MapContext, in *krm.DialogflowSecuritySettingsSpec) *pb.SecuritySettings {
	if in == nil {
		return nil
	}
	out := &pb.SecuritySettings{}
	// MISSING: Name
	// MISSING: DisplayName
	// MISSING: RedactionStrategy
	// MISSING: RedactionScope
	// MISSING: InspectTemplate
	// MISSING: DeidentifyTemplate
	// MISSING: RetentionWindowDays
	// MISSING: RetentionStrategy
	// MISSING: PurgeDataTypes
	// MISSING: AudioExportSettings
	// MISSING: InsightsExportSettings
	return out
}
func SecuritySettings_FromProto(mapCtx *direct.MapContext, in *pb.SecuritySettings) *krm.SecuritySettings {
	if in == nil {
		return nil
	}
	out := &krm.SecuritySettings{}
	out.Name = direct.LazyPtr(in.GetName())
	out.DisplayName = direct.LazyPtr(in.GetDisplayName())
	out.RedactionStrategy = direct.Enum_FromProto(mapCtx, in.GetRedactionStrategy())
	out.RedactionScope = direct.Enum_FromProto(mapCtx, in.GetRedactionScope())
	out.InspectTemplate = direct.LazyPtr(in.GetInspectTemplate())
	out.DeidentifyTemplate = direct.LazyPtr(in.GetDeidentifyTemplate())
	out.RetentionWindowDays = direct.LazyPtr(in.GetRetentionWindowDays())
	out.RetentionStrategy = direct.Enum_FromProto(mapCtx, in.GetRetentionStrategy())
	out.PurgeDataTypes = direct.EnumSlice_FromProto(mapCtx, in.PurgeDataTypes)
	out.AudioExportSettings = SecuritySettings_AudioExportSettings_FromProto(mapCtx, in.GetAudioExportSettings())
	out.InsightsExportSettings = SecuritySettings_InsightsExportSettings_FromProto(mapCtx, in.GetInsightsExportSettings())
	return out
}
func SecuritySettings_ToProto(mapCtx *direct.MapContext, in *krm.SecuritySettings) *pb.SecuritySettings {
	if in == nil {
		return nil
	}
	out := &pb.SecuritySettings{}
	out.Name = direct.ValueOf(in.Name)
	out.DisplayName = direct.ValueOf(in.DisplayName)
	out.RedactionStrategy = direct.Enum_ToProto[pb.SecuritySettings_RedactionStrategy](mapCtx, in.RedactionStrategy)
	out.RedactionScope = direct.Enum_ToProto[pb.SecuritySettings_RedactionScope](mapCtx, in.RedactionScope)
	out.InspectTemplate = direct.ValueOf(in.InspectTemplate)
	out.DeidentifyTemplate = direct.ValueOf(in.DeidentifyTemplate)
	if oneof := SecuritySettings_RetentionWindowDays_ToProto(mapCtx, in.RetentionWindowDays); oneof != nil {
		out.DataRetention = oneof
	}
	if oneof := SecuritySettings_RetentionStrategy_ToProto(mapCtx, in.RetentionStrategy); oneof != nil {
		out.DataRetention = oneof
	}
	out.PurgeDataTypes = direct.EnumSlice_ToProto[pb.SecuritySettings_PurgeDataType](mapCtx, in.PurgeDataTypes)
	out.AudioExportSettings = SecuritySettings_AudioExportSettings_ToProto(mapCtx, in.AudioExportSettings)
	out.InsightsExportSettings = SecuritySettings_InsightsExportSettings_ToProto(mapCtx, in.InsightsExportSettings)
	return out
}
func SecuritySettings_AudioExportSettings_FromProto(mapCtx *direct.MapContext, in *pb.SecuritySettings_AudioExportSettings) *krm.SecuritySettings_AudioExportSettings {
	if in == nil {
		return nil
	}
	out := &krm.SecuritySettings_AudioExportSettings{}
	out.GcsBucket = direct.LazyPtr(in.GetGcsBucket())
	out.AudioExportPattern = direct.LazyPtr(in.GetAudioExportPattern())
	out.EnableAudioRedaction = direct.LazyPtr(in.GetEnableAudioRedaction())
	out.AudioFormat = direct.Enum_FromProto(mapCtx, in.GetAudioFormat())
	out.StoreTtsAudio = direct.LazyPtr(in.GetStoreTtsAudio())
	return out
}
func SecuritySettings_AudioExportSettings_ToProto(mapCtx *direct.MapContext, in *krm.SecuritySettings_AudioExportSettings) *pb.SecuritySettings_AudioExportSettings {
	if in == nil {
		return nil
	}
	out := &pb.SecuritySettings_AudioExportSettings{}
	out.GcsBucket = direct.ValueOf(in.GcsBucket)
	out.AudioExportPattern = direct.ValueOf(in.AudioExportPattern)
	out.EnableAudioRedaction = direct.ValueOf(in.EnableAudioRedaction)
	out.AudioFormat = direct.Enum_ToProto[pb.SecuritySettings_AudioExportSettings_AudioFormat](mapCtx, in.AudioFormat)
	out.StoreTtsAudio = direct.ValueOf(in.StoreTtsAudio)
	return out
}
func SecuritySettings_InsightsExportSettings_FromProto(mapCtx *direct.MapContext, in *pb.SecuritySettings_InsightsExportSettings) *krm.SecuritySettings_InsightsExportSettings {
	if in == nil {
		return nil
	}
	out := &krm.SecuritySettings_InsightsExportSettings{}
	out.EnableInsightsExport = direct.LazyPtr(in.GetEnableInsightsExport())
	return out
}
func SecuritySettings_InsightsExportSettings_ToProto(mapCtx *direct.MapContext, in *krm.SecuritySettings_InsightsExportSettings) *pb.SecuritySettings_InsightsExportSettings {
	if in == nil {
		return nil
	}
	out := &pb.SecuritySettings_InsightsExportSettings{}
	out.EnableInsightsExport = direct.ValueOf(in.EnableInsightsExport)
	return out
}
