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

package resourcesettings

import (
	refs "github.com/GoogleCloudPlatform/k8s-config-connector/apis/refs/v1beta1"
	pb "cloud.google.com/go/resourcesettings/apiv1/resourcesettingspb"
	krm "github.com/GoogleCloudPlatform/k8s-config-connector/apis/resourcesettings/v1alpha1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct"
)
func ResourcesettingsSettingObservedState_FromProto(mapCtx *direct.MapContext, in *pb.Setting) *krm.ResourcesettingsSettingObservedState {
	if in == nil {
		return nil
	}
	out := &krm.ResourcesettingsSettingObservedState{}
	// MISSING: Name
	// MISSING: Metadata
	// MISSING: LocalValue
	// MISSING: EffectiveValue
	// MISSING: Etag
	return out
}
func ResourcesettingsSettingObservedState_ToProto(mapCtx *direct.MapContext, in *krm.ResourcesettingsSettingObservedState) *pb.Setting {
	if in == nil {
		return nil
	}
	out := &pb.Setting{}
	// MISSING: Name
	// MISSING: Metadata
	// MISSING: LocalValue
	// MISSING: EffectiveValue
	// MISSING: Etag
	return out
}
func ResourcesettingsSettingSpec_FromProto(mapCtx *direct.MapContext, in *pb.Setting) *krm.ResourcesettingsSettingSpec {
	if in == nil {
		return nil
	}
	out := &krm.ResourcesettingsSettingSpec{}
	// MISSING: Name
	// MISSING: Metadata
	// MISSING: LocalValue
	// MISSING: EffectiveValue
	// MISSING: Etag
	return out
}
func ResourcesettingsSettingSpec_ToProto(mapCtx *direct.MapContext, in *krm.ResourcesettingsSettingSpec) *pb.Setting {
	if in == nil {
		return nil
	}
	out := &pb.Setting{}
	// MISSING: Name
	// MISSING: Metadata
	// MISSING: LocalValue
	// MISSING: EffectiveValue
	// MISSING: Etag
	return out
}
func Setting_FromProto(mapCtx *direct.MapContext, in *pb.Setting) *krm.Setting {
	if in == nil {
		return nil
	}
	out := &krm.Setting{}
	out.Name = direct.LazyPtr(in.GetName())
	// MISSING: Metadata
	out.LocalValue = Value_FromProto(mapCtx, in.GetLocalValue())
	// MISSING: EffectiveValue
	out.Etag = direct.LazyPtr(in.GetEtag())
	return out
}
func Setting_ToProto(mapCtx *direct.MapContext, in *krm.Setting) *pb.Setting {
	if in == nil {
		return nil
	}
	out := &pb.Setting{}
	out.Name = direct.ValueOf(in.Name)
	// MISSING: Metadata
	out.LocalValue = Value_ToProto(mapCtx, in.LocalValue)
	// MISSING: EffectiveValue
	out.Etag = direct.ValueOf(in.Etag)
	return out
}
func SettingObservedState_FromProto(mapCtx *direct.MapContext, in *pb.Setting) *krm.SettingObservedState {
	if in == nil {
		return nil
	}
	out := &krm.SettingObservedState{}
	// MISSING: Name
	out.Metadata = SettingMetadata_FromProto(mapCtx, in.GetMetadata())
	// MISSING: LocalValue
	out.EffectiveValue = Value_FromProto(mapCtx, in.GetEffectiveValue())
	// MISSING: Etag
	return out
}
func SettingObservedState_ToProto(mapCtx *direct.MapContext, in *krm.SettingObservedState) *pb.Setting {
	if in == nil {
		return nil
	}
	out := &pb.Setting{}
	// MISSING: Name
	out.Metadata = SettingMetadata_ToProto(mapCtx, in.Metadata)
	// MISSING: LocalValue
	out.EffectiveValue = Value_ToProto(mapCtx, in.EffectiveValue)
	// MISSING: Etag
	return out
}
func Value_FromProto(mapCtx *direct.MapContext, in *pb.Value) *krm.Value {
	if in == nil {
		return nil
	}
	out := &krm.Value{}
	out.BooleanValue = direct.LazyPtr(in.GetBooleanValue())
	out.StringValue = direct.LazyPtr(in.GetStringValue())
	out.StringSetValue = Value_StringSet_FromProto(mapCtx, in.GetStringSetValue())
	out.EnumValue = Value_EnumValue_FromProto(mapCtx, in.GetEnumValue())
	return out
}
func Value_ToProto(mapCtx *direct.MapContext, in *krm.Value) *pb.Value {
	if in == nil {
		return nil
	}
	out := &pb.Value{}
	if oneof := Value_BooleanValue_ToProto(mapCtx, in.BooleanValue); oneof != nil {
		out.Value = oneof
	}
	if oneof := Value_StringValue_ToProto(mapCtx, in.StringValue); oneof != nil {
		out.Value = oneof
	}
	if oneof := Value_StringSet_ToProto(mapCtx, in.StringSetValue); oneof != nil {
		out.Value = &pb.Value_StringSetValue{StringSetValue: oneof}
	}
	if oneof := Value_EnumValue_ToProto(mapCtx, in.EnumValue); oneof != nil {
		out.Value = &pb.Value_EnumValue_{EnumValue: oneof}
	}
	return out
}
func Value_EnumValue_FromProto(mapCtx *direct.MapContext, in *pb.Value_EnumValue) *krm.Value_EnumValue {
	if in == nil {
		return nil
	}
	out := &krm.Value_EnumValue{}
	out.Value = direct.LazyPtr(in.GetValue())
	return out
}
func Value_EnumValue_ToProto(mapCtx *direct.MapContext, in *krm.Value_EnumValue) *pb.Value_EnumValue {
	if in == nil {
		return nil
	}
	out := &pb.Value_EnumValue{}
	out.Value = direct.ValueOf(in.Value)
	return out
}
func Value_StringSet_FromProto(mapCtx *direct.MapContext, in *pb.Value_StringSet) *krm.Value_StringSet {
	if in == nil {
		return nil
	}
	out := &krm.Value_StringSet{}
	out.Values = in.Values
	return out
}
func Value_StringSet_ToProto(mapCtx *direct.MapContext, in *krm.Value_StringSet) *pb.Value_StringSet {
	if in == nil {
		return nil
	}
	out := &pb.Value_StringSet{}
	out.Values = in.Values
	return out
}
