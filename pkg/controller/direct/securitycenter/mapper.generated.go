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
func EffectiveEventThreatDetectionCustomModule_FromProto(mapCtx *direct.MapContext, in *pb.EffectiveEventThreatDetectionCustomModule) *krm.EffectiveEventThreatDetectionCustomModule {
	if in == nil {
		return nil
	}
	out := &krm.EffectiveEventThreatDetectionCustomModule{}
	// MISSING: Name
	// MISSING: Config
	// MISSING: EnablementState
	// MISSING: Type
	// MISSING: DisplayName
	// MISSING: Description
	return out
}
func EffectiveEventThreatDetectionCustomModule_ToProto(mapCtx *direct.MapContext, in *krm.EffectiveEventThreatDetectionCustomModule) *pb.EffectiveEventThreatDetectionCustomModule {
	if in == nil {
		return nil
	}
	out := &pb.EffectiveEventThreatDetectionCustomModule{}
	// MISSING: Name
	// MISSING: Config
	// MISSING: EnablementState
	// MISSING: Type
	// MISSING: DisplayName
	// MISSING: Description
	return out
}
func EffectiveEventThreatDetectionCustomModuleObservedState_FromProto(mapCtx *direct.MapContext, in *pb.EffectiveEventThreatDetectionCustomModule) *krm.EffectiveEventThreatDetectionCustomModuleObservedState {
	if in == nil {
		return nil
	}
	out := &krm.EffectiveEventThreatDetectionCustomModuleObservedState{}
	out.Name = direct.LazyPtr(in.GetName())
	out.Config = Config_FromProto(mapCtx, in.GetConfig())
	out.EnablementState = direct.Enum_FromProto(mapCtx, in.GetEnablementState())
	out.Type = direct.LazyPtr(in.GetType())
	out.DisplayName = direct.LazyPtr(in.GetDisplayName())
	out.Description = direct.LazyPtr(in.GetDescription())
	return out
}
func EffectiveEventThreatDetectionCustomModuleObservedState_ToProto(mapCtx *direct.MapContext, in *krm.EffectiveEventThreatDetectionCustomModuleObservedState) *pb.EffectiveEventThreatDetectionCustomModule {
	if in == nil {
		return nil
	}
	out := &pb.EffectiveEventThreatDetectionCustomModule{}
	out.Name = direct.ValueOf(in.Name)
	out.Config = Config_ToProto(mapCtx, in.Config)
	out.EnablementState = direct.Enum_ToProto[pb.EffectiveEventThreatDetectionCustomModule_EnablementState](mapCtx, in.EnablementState)
	out.Type = direct.ValueOf(in.Type)
	out.DisplayName = direct.ValueOf(in.DisplayName)
	out.Description = direct.ValueOf(in.Description)
	return out
}
func SecuritycenterEffectiveEventThreatDetectionCustomModuleObservedState_FromProto(mapCtx *direct.MapContext, in *pb.EffectiveEventThreatDetectionCustomModule) *krm.SecuritycenterEffectiveEventThreatDetectionCustomModuleObservedState {
	if in == nil {
		return nil
	}
	out := &krm.SecuritycenterEffectiveEventThreatDetectionCustomModuleObservedState{}
	// MISSING: Name
	// MISSING: Config
	// MISSING: EnablementState
	// MISSING: Type
	// MISSING: DisplayName
	// MISSING: Description
	return out
}
func SecuritycenterEffectiveEventThreatDetectionCustomModuleObservedState_ToProto(mapCtx *direct.MapContext, in *krm.SecuritycenterEffectiveEventThreatDetectionCustomModuleObservedState) *pb.EffectiveEventThreatDetectionCustomModule {
	if in == nil {
		return nil
	}
	out := &pb.EffectiveEventThreatDetectionCustomModule{}
	// MISSING: Name
	// MISSING: Config
	// MISSING: EnablementState
	// MISSING: Type
	// MISSING: DisplayName
	// MISSING: Description
	return out
}
func SecuritycenterEffectiveEventThreatDetectionCustomModuleSpec_FromProto(mapCtx *direct.MapContext, in *pb.EffectiveEventThreatDetectionCustomModule) *krm.SecuritycenterEffectiveEventThreatDetectionCustomModuleSpec {
	if in == nil {
		return nil
	}
	out := &krm.SecuritycenterEffectiveEventThreatDetectionCustomModuleSpec{}
	// MISSING: Name
	// MISSING: Config
	// MISSING: EnablementState
	// MISSING: Type
	// MISSING: DisplayName
	// MISSING: Description
	return out
}
func SecuritycenterEffectiveEventThreatDetectionCustomModuleSpec_ToProto(mapCtx *direct.MapContext, in *krm.SecuritycenterEffectiveEventThreatDetectionCustomModuleSpec) *pb.EffectiveEventThreatDetectionCustomModule {
	if in == nil {
		return nil
	}
	out := &pb.EffectiveEventThreatDetectionCustomModule{}
	// MISSING: Name
	// MISSING: Config
	// MISSING: EnablementState
	// MISSING: Type
	// MISSING: DisplayName
	// MISSING: Description
	return out
}
