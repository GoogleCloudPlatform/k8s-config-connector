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
	refs "github.com/GoogleCloudPlatform/k8s-config-connector/apis/refs/v1beta1"
	pb "cloud.google.com/go/securitycentermanagement/apiv1/securitycentermanagementpb"
	krm "github.com/GoogleCloudPlatform/k8s-config-connector/apis/securitycentermanagement/v1alpha1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct"
)
func EventThreatDetectionCustomModule_FromProto(mapCtx *direct.MapContext, in *pb.EventThreatDetectionCustomModule) *krm.EventThreatDetectionCustomModule {
	if in == nil {
		return nil
	}
	out := &krm.EventThreatDetectionCustomModule{}
	out.Name = direct.LazyPtr(in.GetName())
	out.Config = Config_FromProto(mapCtx, in.GetConfig())
	// MISSING: AncestorModule
	out.EnablementState = direct.Enum_FromProto(mapCtx, in.GetEnablementState())
	out.Type = direct.LazyPtr(in.GetType())
	out.DisplayName = direct.LazyPtr(in.GetDisplayName())
	out.Description = direct.LazyPtr(in.GetDescription())
	// MISSING: UpdateTime
	// MISSING: LastEditor
	return out
}
func EventThreatDetectionCustomModule_ToProto(mapCtx *direct.MapContext, in *krm.EventThreatDetectionCustomModule) *pb.EventThreatDetectionCustomModule {
	if in == nil {
		return nil
	}
	out := &pb.EventThreatDetectionCustomModule{}
	out.Name = direct.ValueOf(in.Name)
	out.Config = Config_ToProto(mapCtx, in.Config)
	// MISSING: AncestorModule
	out.EnablementState = direct.Enum_ToProto[pb.EventThreatDetectionCustomModule_EnablementState](mapCtx, in.EnablementState)
	out.Type = direct.ValueOf(in.Type)
	out.DisplayName = direct.ValueOf(in.DisplayName)
	out.Description = direct.ValueOf(in.Description)
	// MISSING: UpdateTime
	// MISSING: LastEditor
	return out
}
func EventThreatDetectionCustomModuleObservedState_FromProto(mapCtx *direct.MapContext, in *pb.EventThreatDetectionCustomModule) *krm.EventThreatDetectionCustomModuleObservedState {
	if in == nil {
		return nil
	}
	out := &krm.EventThreatDetectionCustomModuleObservedState{}
	// MISSING: Name
	// MISSING: Config
	out.AncestorModule = direct.LazyPtr(in.GetAncestorModule())
	// MISSING: EnablementState
	// MISSING: Type
	// MISSING: DisplayName
	// MISSING: Description
	out.UpdateTime = direct.StringTimestamp_FromProto(mapCtx, in.GetUpdateTime())
	out.LastEditor = direct.LazyPtr(in.GetLastEditor())
	return out
}
func EventThreatDetectionCustomModuleObservedState_ToProto(mapCtx *direct.MapContext, in *krm.EventThreatDetectionCustomModuleObservedState) *pb.EventThreatDetectionCustomModule {
	if in == nil {
		return nil
	}
	out := &pb.EventThreatDetectionCustomModule{}
	// MISSING: Name
	// MISSING: Config
	out.AncestorModule = direct.ValueOf(in.AncestorModule)
	// MISSING: EnablementState
	// MISSING: Type
	// MISSING: DisplayName
	// MISSING: Description
	out.UpdateTime = direct.StringTimestamp_ToProto(mapCtx, in.UpdateTime)
	out.LastEditor = direct.ValueOf(in.LastEditor)
	return out
}
func SecuritycentermanagementEventThreatDetectionCustomModuleObservedState_FromProto(mapCtx *direct.MapContext, in *pb.EventThreatDetectionCustomModule) *krm.SecuritycentermanagementEventThreatDetectionCustomModuleObservedState {
	if in == nil {
		return nil
	}
	out := &krm.SecuritycentermanagementEventThreatDetectionCustomModuleObservedState{}
	// MISSING: Name
	// MISSING: Config
	// MISSING: AncestorModule
	// MISSING: EnablementState
	// MISSING: Type
	// MISSING: DisplayName
	// MISSING: Description
	// MISSING: UpdateTime
	// MISSING: LastEditor
	return out
}
func SecuritycentermanagementEventThreatDetectionCustomModuleObservedState_ToProto(mapCtx *direct.MapContext, in *krm.SecuritycentermanagementEventThreatDetectionCustomModuleObservedState) *pb.EventThreatDetectionCustomModule {
	if in == nil {
		return nil
	}
	out := &pb.EventThreatDetectionCustomModule{}
	// MISSING: Name
	// MISSING: Config
	// MISSING: AncestorModule
	// MISSING: EnablementState
	// MISSING: Type
	// MISSING: DisplayName
	// MISSING: Description
	// MISSING: UpdateTime
	// MISSING: LastEditor
	return out
}
func SecuritycentermanagementEventThreatDetectionCustomModuleSpec_FromProto(mapCtx *direct.MapContext, in *pb.EventThreatDetectionCustomModule) *krm.SecuritycentermanagementEventThreatDetectionCustomModuleSpec {
	if in == nil {
		return nil
	}
	out := &krm.SecuritycentermanagementEventThreatDetectionCustomModuleSpec{}
	// MISSING: Name
	// MISSING: Config
	// MISSING: AncestorModule
	// MISSING: EnablementState
	// MISSING: Type
	// MISSING: DisplayName
	// MISSING: Description
	// MISSING: UpdateTime
	// MISSING: LastEditor
	return out
}
func SecuritycentermanagementEventThreatDetectionCustomModuleSpec_ToProto(mapCtx *direct.MapContext, in *krm.SecuritycentermanagementEventThreatDetectionCustomModuleSpec) *pb.EventThreatDetectionCustomModule {
	if in == nil {
		return nil
	}
	out := &pb.EventThreatDetectionCustomModule{}
	// MISSING: Name
	// MISSING: Config
	// MISSING: AncestorModule
	// MISSING: EnablementState
	// MISSING: Type
	// MISSING: DisplayName
	// MISSING: Description
	// MISSING: UpdateTime
	// MISSING: LastEditor
	return out
}
