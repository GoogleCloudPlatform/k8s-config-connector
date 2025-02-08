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
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct"
	refs "github.com/GoogleCloudPlatform/k8s-config-connector/apis/refs/v1beta1"
	pb "cloud.google.com/go/securitycentermanagement/apiv1/securitycentermanagementpb"
	krm "github.com/GoogleCloudPlatform/k8s-config-connector/apis/securitycentermanagement/v1alpha1"
)
func SecurityCenterService_FromProto(mapCtx *direct.MapContext, in *pb.SecurityCenterService) *krm.SecurityCenterService {
	if in == nil {
		return nil
	}
	out := &krm.SecurityCenterService{}
	out.Name = direct.LazyPtr(in.GetName())
	out.IntendedEnablementState = direct.Enum_FromProto(mapCtx, in.GetIntendedEnablementState())
	// MISSING: EffectiveEnablementState
	// MISSING: Modules
	// MISSING: UpdateTime
	out.ServiceConfig = ServiceConfig_FromProto(mapCtx, in.GetServiceConfig())
	return out
}
func SecurityCenterService_ToProto(mapCtx *direct.MapContext, in *krm.SecurityCenterService) *pb.SecurityCenterService {
	if in == nil {
		return nil
	}
	out := &pb.SecurityCenterService{}
	out.Name = direct.ValueOf(in.Name)
	out.IntendedEnablementState = direct.Enum_ToProto[pb.SecurityCenterService_EnablementState](mapCtx, in.IntendedEnablementState)
	// MISSING: EffectiveEnablementState
	// MISSING: Modules
	// MISSING: UpdateTime
	out.ServiceConfig = ServiceConfig_ToProto(mapCtx, in.ServiceConfig)
	return out
}
func SecurityCenterServiceObservedState_FromProto(mapCtx *direct.MapContext, in *pb.SecurityCenterService) *krm.SecurityCenterServiceObservedState {
	if in == nil {
		return nil
	}
	out := &krm.SecurityCenterServiceObservedState{}
	// MISSING: Name
	// MISSING: IntendedEnablementState
	out.EffectiveEnablementState = direct.Enum_FromProto(mapCtx, in.GetEffectiveEnablementState())
	// MISSING: Modules
	out.UpdateTime = direct.StringTimestamp_FromProto(mapCtx, in.GetUpdateTime())
	// MISSING: ServiceConfig
	return out
}
func SecurityCenterServiceObservedState_ToProto(mapCtx *direct.MapContext, in *krm.SecurityCenterServiceObservedState) *pb.SecurityCenterService {
	if in == nil {
		return nil
	}
	out := &pb.SecurityCenterService{}
	// MISSING: Name
	// MISSING: IntendedEnablementState
	out.EffectiveEnablementState = direct.Enum_ToProto[pb.SecurityCenterService_EnablementState](mapCtx, in.EffectiveEnablementState)
	// MISSING: Modules
	out.UpdateTime = direct.StringTimestamp_ToProto(mapCtx, in.UpdateTime)
	// MISSING: ServiceConfig
	return out
}
func SecurityCenterService_ModuleSettings_FromProto(mapCtx *direct.MapContext, in *pb.SecurityCenterService_ModuleSettings) *krm.SecurityCenterService_ModuleSettings {
	if in == nil {
		return nil
	}
	out := &krm.SecurityCenterService_ModuleSettings{}
	out.IntendedEnablementState = direct.Enum_FromProto(mapCtx, in.GetIntendedEnablementState())
	// MISSING: EffectiveEnablementState
	return out
}
func SecurityCenterService_ModuleSettings_ToProto(mapCtx *direct.MapContext, in *krm.SecurityCenterService_ModuleSettings) *pb.SecurityCenterService_ModuleSettings {
	if in == nil {
		return nil
	}
	out := &pb.SecurityCenterService_ModuleSettings{}
	out.IntendedEnablementState = direct.Enum_ToProto[pb.SecurityCenterService_EnablementState](mapCtx, in.IntendedEnablementState)
	// MISSING: EffectiveEnablementState
	return out
}
func SecuritycentermanagementSecurityCenterServiceObservedState_FromProto(mapCtx *direct.MapContext, in *pb.SecurityCenterService) *krm.SecuritycentermanagementSecurityCenterServiceObservedState {
	if in == nil {
		return nil
	}
	out := &krm.SecuritycentermanagementSecurityCenterServiceObservedState{}
	// MISSING: Name
	// MISSING: IntendedEnablementState
	// MISSING: EffectiveEnablementState
	// MISSING: Modules
	// MISSING: UpdateTime
	// MISSING: ServiceConfig
	return out
}
func SecuritycentermanagementSecurityCenterServiceObservedState_ToProto(mapCtx *direct.MapContext, in *krm.SecuritycentermanagementSecurityCenterServiceObservedState) *pb.SecurityCenterService {
	if in == nil {
		return nil
	}
	out := &pb.SecurityCenterService{}
	// MISSING: Name
	// MISSING: IntendedEnablementState
	// MISSING: EffectiveEnablementState
	// MISSING: Modules
	// MISSING: UpdateTime
	// MISSING: ServiceConfig
	return out
}
func SecuritycentermanagementSecurityCenterServiceSpec_FromProto(mapCtx *direct.MapContext, in *pb.SecurityCenterService) *krm.SecuritycentermanagementSecurityCenterServiceSpec {
	if in == nil {
		return nil
	}
	out := &krm.SecuritycentermanagementSecurityCenterServiceSpec{}
	// MISSING: Name
	// MISSING: IntendedEnablementState
	// MISSING: EffectiveEnablementState
	// MISSING: Modules
	// MISSING: UpdateTime
	// MISSING: ServiceConfig
	return out
}
func SecuritycentermanagementSecurityCenterServiceSpec_ToProto(mapCtx *direct.MapContext, in *krm.SecuritycentermanagementSecurityCenterServiceSpec) *pb.SecurityCenterService {
	if in == nil {
		return nil
	}
	out := &pb.SecurityCenterService{}
	// MISSING: Name
	// MISSING: IntendedEnablementState
	// MISSING: EffectiveEnablementState
	// MISSING: Modules
	// MISSING: UpdateTime
	// MISSING: ServiceConfig
	return out
}
