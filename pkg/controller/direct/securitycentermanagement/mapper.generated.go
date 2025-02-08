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

package securitycentermanagement

import (
	pb "cloud.google.com/go/securitycentermanagement/apiv1/securitycentermanagementpb"
	krm "github.com/GoogleCloudPlatform/k8s-config-connector/apis/securitycentermanagement/v1alpha1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct"
	refs "github.com/GoogleCloudPlatform/k8s-config-connector/apis/refs/v1beta1"
)
func EffectiveEventThreatDetectionCustomModule_FromProto(mapCtx *direct.MapContext, in *pb.EffectiveEventThreatDetectionCustomModule) *krm.EffectiveEventThreatDetectionCustomModule {
	if in == nil {
		return nil
	}
	out := &krm.EffectiveEventThreatDetectionCustomModule{}
	out.Name = direct.LazyPtr(in.GetName())
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
	out.Name = direct.ValueOf(in.Name)
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
	// MISSING: Name
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
	// MISSING: Name
	out.Config = Config_ToProto(mapCtx, in.Config)
	out.EnablementState = direct.Enum_ToProto[pb.EffectiveEventThreatDetectionCustomModule_EnablementState](mapCtx, in.EnablementState)
	out.Type = direct.ValueOf(in.Type)
	out.DisplayName = direct.ValueOf(in.DisplayName)
	out.Description = direct.ValueOf(in.Description)
	return out
}
func SecuritycentermanagementEffectiveEventThreatDetectionCustomModuleObservedState_FromProto(mapCtx *direct.MapContext, in *pb.EffectiveEventThreatDetectionCustomModule) *krm.SecuritycentermanagementEffectiveEventThreatDetectionCustomModuleObservedState {
	if in == nil {
		return nil
	}
	out := &krm.SecuritycentermanagementEffectiveEventThreatDetectionCustomModuleObservedState{}
	// MISSING: Name
	// MISSING: Config
	// MISSING: EnablementState
	// MISSING: Type
	// MISSING: DisplayName
	// MISSING: Description
	return out
}
func SecuritycentermanagementEffectiveEventThreatDetectionCustomModuleObservedState_ToProto(mapCtx *direct.MapContext, in *krm.SecuritycentermanagementEffectiveEventThreatDetectionCustomModuleObservedState) *pb.EffectiveEventThreatDetectionCustomModule {
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
func SecuritycentermanagementEffectiveEventThreatDetectionCustomModuleSpec_FromProto(mapCtx *direct.MapContext, in *pb.EffectiveEventThreatDetectionCustomModule) *krm.SecuritycentermanagementEffectiveEventThreatDetectionCustomModuleSpec {
	if in == nil {
		return nil
	}
	out := &krm.SecuritycentermanagementEffectiveEventThreatDetectionCustomModuleSpec{}
	// MISSING: Name
	// MISSING: Config
	// MISSING: EnablementState
	// MISSING: Type
	// MISSING: DisplayName
	// MISSING: Description
	return out
}
func SecuritycentermanagementEffectiveEventThreatDetectionCustomModuleSpec_ToProto(mapCtx *direct.MapContext, in *krm.SecuritycentermanagementEffectiveEventThreatDetectionCustomModuleSpec) *pb.EffectiveEventThreatDetectionCustomModule {
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
