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
	refs "github.com/GoogleCloudPlatform/k8s-config-connector/apis/refs/v1beta1"
	pb "cloud.google.com/go/securitycenter/apiv1/securitycenterpb"
	krm "github.com/GoogleCloudPlatform/k8s-config-connector/apis/securitycenter/v1alpha1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct"
)
func ResourceValueConfig_FromProto(mapCtx *direct.MapContext, in *pb.ResourceValueConfig) *krm.ResourceValueConfig {
	if in == nil {
		return nil
	}
	out := &krm.ResourceValueConfig{}
	out.Name = direct.LazyPtr(in.GetName())
	out.ResourceValue = direct.Enum_FromProto(mapCtx, in.GetResourceValue())
	out.TagValues = in.TagValues
	out.ResourceType = direct.LazyPtr(in.GetResourceType())
	out.Scope = direct.LazyPtr(in.GetScope())
	out.ResourceLabelsSelector = in.ResourceLabelsSelector
	out.Description = direct.LazyPtr(in.GetDescription())
	// MISSING: CreateTime
	// MISSING: UpdateTime
	out.CloudProvider = direct.Enum_FromProto(mapCtx, in.GetCloudProvider())
	out.SensitiveDataProtectionMapping = ResourceValueConfig_SensitiveDataProtectionMapping_FromProto(mapCtx, in.GetSensitiveDataProtectionMapping())
	return out
}
func ResourceValueConfig_ToProto(mapCtx *direct.MapContext, in *krm.ResourceValueConfig) *pb.ResourceValueConfig {
	if in == nil {
		return nil
	}
	out := &pb.ResourceValueConfig{}
	out.Name = direct.ValueOf(in.Name)
	out.ResourceValue = direct.Enum_ToProto[pb.ResourceValue](mapCtx, in.ResourceValue)
	out.TagValues = in.TagValues
	out.ResourceType = direct.ValueOf(in.ResourceType)
	out.Scope = direct.ValueOf(in.Scope)
	out.ResourceLabelsSelector = in.ResourceLabelsSelector
	out.Description = direct.ValueOf(in.Description)
	// MISSING: CreateTime
	// MISSING: UpdateTime
	out.CloudProvider = direct.Enum_ToProto[pb.CloudProvider](mapCtx, in.CloudProvider)
	out.SensitiveDataProtectionMapping = ResourceValueConfig_SensitiveDataProtectionMapping_ToProto(mapCtx, in.SensitiveDataProtectionMapping)
	return out
}
func ResourceValueConfigObservedState_FromProto(mapCtx *direct.MapContext, in *pb.ResourceValueConfig) *krm.ResourceValueConfigObservedState {
	if in == nil {
		return nil
	}
	out := &krm.ResourceValueConfigObservedState{}
	// MISSING: Name
	// MISSING: ResourceValue
	// MISSING: TagValues
	// MISSING: ResourceType
	// MISSING: Scope
	// MISSING: ResourceLabelsSelector
	// MISSING: Description
	out.CreateTime = direct.StringTimestamp_FromProto(mapCtx, in.GetCreateTime())
	out.UpdateTime = direct.StringTimestamp_FromProto(mapCtx, in.GetUpdateTime())
	// MISSING: CloudProvider
	// MISSING: SensitiveDataProtectionMapping
	return out
}
func ResourceValueConfigObservedState_ToProto(mapCtx *direct.MapContext, in *krm.ResourceValueConfigObservedState) *pb.ResourceValueConfig {
	if in == nil {
		return nil
	}
	out := &pb.ResourceValueConfig{}
	// MISSING: Name
	// MISSING: ResourceValue
	// MISSING: TagValues
	// MISSING: ResourceType
	// MISSING: Scope
	// MISSING: ResourceLabelsSelector
	// MISSING: Description
	out.CreateTime = direct.StringTimestamp_ToProto(mapCtx, in.CreateTime)
	out.UpdateTime = direct.StringTimestamp_ToProto(mapCtx, in.UpdateTime)
	// MISSING: CloudProvider
	// MISSING: SensitiveDataProtectionMapping
	return out
}
func ResourceValueConfig_SensitiveDataProtectionMapping_FromProto(mapCtx *direct.MapContext, in *pb.ResourceValueConfig_SensitiveDataProtectionMapping) *krm.ResourceValueConfig_SensitiveDataProtectionMapping {
	if in == nil {
		return nil
	}
	out := &krm.ResourceValueConfig_SensitiveDataProtectionMapping{}
	out.HighSensitivityMapping = direct.Enum_FromProto(mapCtx, in.GetHighSensitivityMapping())
	out.MediumSensitivityMapping = direct.Enum_FromProto(mapCtx, in.GetMediumSensitivityMapping())
	return out
}
func ResourceValueConfig_SensitiveDataProtectionMapping_ToProto(mapCtx *direct.MapContext, in *krm.ResourceValueConfig_SensitiveDataProtectionMapping) *pb.ResourceValueConfig_SensitiveDataProtectionMapping {
	if in == nil {
		return nil
	}
	out := &pb.ResourceValueConfig_SensitiveDataProtectionMapping{}
	out.HighSensitivityMapping = direct.Enum_ToProto[pb.ResourceValue](mapCtx, in.HighSensitivityMapping)
	out.MediumSensitivityMapping = direct.Enum_ToProto[pb.ResourceValue](mapCtx, in.MediumSensitivityMapping)
	return out
}
func SecuritycenterResourceValueConfigObservedState_FromProto(mapCtx *direct.MapContext, in *pb.ResourceValueConfig) *krm.SecuritycenterResourceValueConfigObservedState {
	if in == nil {
		return nil
	}
	out := &krm.SecuritycenterResourceValueConfigObservedState{}
	// MISSING: Name
	// MISSING: ResourceValue
	// MISSING: TagValues
	// MISSING: ResourceType
	// MISSING: Scope
	// MISSING: ResourceLabelsSelector
	// MISSING: Description
	// MISSING: CreateTime
	// MISSING: UpdateTime
	// MISSING: CloudProvider
	// MISSING: SensitiveDataProtectionMapping
	return out
}
func SecuritycenterResourceValueConfigObservedState_ToProto(mapCtx *direct.MapContext, in *krm.SecuritycenterResourceValueConfigObservedState) *pb.ResourceValueConfig {
	if in == nil {
		return nil
	}
	out := &pb.ResourceValueConfig{}
	// MISSING: Name
	// MISSING: ResourceValue
	// MISSING: TagValues
	// MISSING: ResourceType
	// MISSING: Scope
	// MISSING: ResourceLabelsSelector
	// MISSING: Description
	// MISSING: CreateTime
	// MISSING: UpdateTime
	// MISSING: CloudProvider
	// MISSING: SensitiveDataProtectionMapping
	return out
}
func SecuritycenterResourceValueConfigSpec_FromProto(mapCtx *direct.MapContext, in *pb.ResourceValueConfig) *krm.SecuritycenterResourceValueConfigSpec {
	if in == nil {
		return nil
	}
	out := &krm.SecuritycenterResourceValueConfigSpec{}
	// MISSING: Name
	// MISSING: ResourceValue
	// MISSING: TagValues
	// MISSING: ResourceType
	// MISSING: Scope
	// MISSING: ResourceLabelsSelector
	// MISSING: Description
	// MISSING: CreateTime
	// MISSING: UpdateTime
	// MISSING: CloudProvider
	// MISSING: SensitiveDataProtectionMapping
	return out
}
func SecuritycenterResourceValueConfigSpec_ToProto(mapCtx *direct.MapContext, in *krm.SecuritycenterResourceValueConfigSpec) *pb.ResourceValueConfig {
	if in == nil {
		return nil
	}
	out := &pb.ResourceValueConfig{}
	// MISSING: Name
	// MISSING: ResourceValue
	// MISSING: TagValues
	// MISSING: ResourceType
	// MISSING: Scope
	// MISSING: ResourceLabelsSelector
	// MISSING: Description
	// MISSING: CreateTime
	// MISSING: UpdateTime
	// MISSING: CloudProvider
	// MISSING: SensitiveDataProtectionMapping
	return out
}
