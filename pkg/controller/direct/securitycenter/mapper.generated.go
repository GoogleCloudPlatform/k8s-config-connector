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
	pb "cloud.google.com/go/securitycenter/apiv1/securitycenterpb"
	krm "github.com/GoogleCloudPlatform/k8s-config-connector/apis/securitycenter/v1alpha1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct"
	refs "github.com/GoogleCloudPlatform/k8s-config-connector/apis/refs/v1beta1"
)
func CustomConfig_FromProto(mapCtx *direct.MapContext, in *pb.CustomConfig) *krm.CustomConfig {
	if in == nil {
		return nil
	}
	out := &krm.CustomConfig{}
	out.Predicate = Expr_FromProto(mapCtx, in.GetPredicate())
	out.CustomOutput = CustomConfig_CustomOutputSpec_FromProto(mapCtx, in.GetCustomOutput())
	out.ResourceSelector = CustomConfig_ResourceSelector_FromProto(mapCtx, in.GetResourceSelector())
	out.Severity = direct.Enum_FromProto(mapCtx, in.GetSeverity())
	out.Description = direct.LazyPtr(in.GetDescription())
	out.Recommendation = direct.LazyPtr(in.GetRecommendation())
	return out
}
func CustomConfig_ToProto(mapCtx *direct.MapContext, in *krm.CustomConfig) *pb.CustomConfig {
	if in == nil {
		return nil
	}
	out := &pb.CustomConfig{}
	out.Predicate = Expr_ToProto(mapCtx, in.Predicate)
	out.CustomOutput = CustomConfig_CustomOutputSpec_ToProto(mapCtx, in.CustomOutput)
	out.ResourceSelector = CustomConfig_ResourceSelector_ToProto(mapCtx, in.ResourceSelector)
	out.Severity = direct.Enum_ToProto[pb.CustomConfig_Severity](mapCtx, in.Severity)
	out.Description = direct.ValueOf(in.Description)
	out.Recommendation = direct.ValueOf(in.Recommendation)
	return out
}
func CustomConfig_CustomOutputSpec_FromProto(mapCtx *direct.MapContext, in *pb.CustomConfig_CustomOutputSpec) *krm.CustomConfig_CustomOutputSpec {
	if in == nil {
		return nil
	}
	out := &krm.CustomConfig_CustomOutputSpec{}
	out.Properties = direct.Slice_FromProto(mapCtx, in.Properties, CustomConfig_CustomOutputSpec_Property_FromProto)
	return out
}
func CustomConfig_CustomOutputSpec_ToProto(mapCtx *direct.MapContext, in *krm.CustomConfig_CustomOutputSpec) *pb.CustomConfig_CustomOutputSpec {
	if in == nil {
		return nil
	}
	out := &pb.CustomConfig_CustomOutputSpec{}
	out.Properties = direct.Slice_ToProto(mapCtx, in.Properties, CustomConfig_CustomOutputSpec_Property_ToProto)
	return out
}
func CustomConfig_CustomOutputSpec_Property_FromProto(mapCtx *direct.MapContext, in *pb.CustomConfig_CustomOutputSpec_Property) *krm.CustomConfig_CustomOutputSpec_Property {
	if in == nil {
		return nil
	}
	out := &krm.CustomConfig_CustomOutputSpec_Property{}
	out.Name = direct.LazyPtr(in.GetName())
	out.ValueExpression = Expr_FromProto(mapCtx, in.GetValueExpression())
	return out
}
func CustomConfig_CustomOutputSpec_Property_ToProto(mapCtx *direct.MapContext, in *krm.CustomConfig_CustomOutputSpec_Property) *pb.CustomConfig_CustomOutputSpec_Property {
	if in == nil {
		return nil
	}
	out := &pb.CustomConfig_CustomOutputSpec_Property{}
	out.Name = direct.ValueOf(in.Name)
	out.ValueExpression = Expr_ToProto(mapCtx, in.ValueExpression)
	return out
}
func CustomConfig_ResourceSelector_FromProto(mapCtx *direct.MapContext, in *pb.CustomConfig_ResourceSelector) *krm.CustomConfig_ResourceSelector {
	if in == nil {
		return nil
	}
	out := &krm.CustomConfig_ResourceSelector{}
	out.ResourceTypes = in.ResourceTypes
	return out
}
func CustomConfig_ResourceSelector_ToProto(mapCtx *direct.MapContext, in *krm.CustomConfig_ResourceSelector) *pb.CustomConfig_ResourceSelector {
	if in == nil {
		return nil
	}
	out := &pb.CustomConfig_ResourceSelector{}
	out.ResourceTypes = in.ResourceTypes
	return out
}
func SecurityHealthAnalyticsCustomModule_FromProto(mapCtx *direct.MapContext, in *pb.SecurityHealthAnalyticsCustomModule) *krm.SecurityHealthAnalyticsCustomModule {
	if in == nil {
		return nil
	}
	out := &krm.SecurityHealthAnalyticsCustomModule{}
	out.Name = direct.LazyPtr(in.GetName())
	out.DisplayName = direct.LazyPtr(in.GetDisplayName())
	out.EnablementState = direct.Enum_FromProto(mapCtx, in.GetEnablementState())
	// MISSING: UpdateTime
	// MISSING: LastEditor
	// MISSING: AncestorModule
	out.CustomConfig = CustomConfig_FromProto(mapCtx, in.GetCustomConfig())
	return out
}
func SecurityHealthAnalyticsCustomModule_ToProto(mapCtx *direct.MapContext, in *krm.SecurityHealthAnalyticsCustomModule) *pb.SecurityHealthAnalyticsCustomModule {
	if in == nil {
		return nil
	}
	out := &pb.SecurityHealthAnalyticsCustomModule{}
	out.Name = direct.ValueOf(in.Name)
	out.DisplayName = direct.ValueOf(in.DisplayName)
	out.EnablementState = direct.Enum_ToProto[pb.SecurityHealthAnalyticsCustomModule_EnablementState](mapCtx, in.EnablementState)
	// MISSING: UpdateTime
	// MISSING: LastEditor
	// MISSING: AncestorModule
	out.CustomConfig = CustomConfig_ToProto(mapCtx, in.CustomConfig)
	return out
}
func SecurityHealthAnalyticsCustomModuleObservedState_FromProto(mapCtx *direct.MapContext, in *pb.SecurityHealthAnalyticsCustomModule) *krm.SecurityHealthAnalyticsCustomModuleObservedState {
	if in == nil {
		return nil
	}
	out := &krm.SecurityHealthAnalyticsCustomModuleObservedState{}
	// MISSING: Name
	// MISSING: DisplayName
	// MISSING: EnablementState
	out.UpdateTime = direct.StringTimestamp_FromProto(mapCtx, in.GetUpdateTime())
	out.LastEditor = direct.LazyPtr(in.GetLastEditor())
	out.AncestorModule = direct.LazyPtr(in.GetAncestorModule())
	// MISSING: CustomConfig
	return out
}
func SecurityHealthAnalyticsCustomModuleObservedState_ToProto(mapCtx *direct.MapContext, in *krm.SecurityHealthAnalyticsCustomModuleObservedState) *pb.SecurityHealthAnalyticsCustomModule {
	if in == nil {
		return nil
	}
	out := &pb.SecurityHealthAnalyticsCustomModule{}
	// MISSING: Name
	// MISSING: DisplayName
	// MISSING: EnablementState
	out.UpdateTime = direct.StringTimestamp_ToProto(mapCtx, in.UpdateTime)
	out.LastEditor = direct.ValueOf(in.LastEditor)
	out.AncestorModule = direct.ValueOf(in.AncestorModule)
	// MISSING: CustomConfig
	return out
}
func SecuritycenterSecurityHealthAnalyticsCustomModuleObservedState_FromProto(mapCtx *direct.MapContext, in *pb.SecurityHealthAnalyticsCustomModule) *krm.SecuritycenterSecurityHealthAnalyticsCustomModuleObservedState {
	if in == nil {
		return nil
	}
	out := &krm.SecuritycenterSecurityHealthAnalyticsCustomModuleObservedState{}
	// MISSING: Name
	// MISSING: DisplayName
	// MISSING: EnablementState
	// MISSING: UpdateTime
	// MISSING: LastEditor
	// MISSING: AncestorModule
	// MISSING: CustomConfig
	return out
}
func SecuritycenterSecurityHealthAnalyticsCustomModuleObservedState_ToProto(mapCtx *direct.MapContext, in *krm.SecuritycenterSecurityHealthAnalyticsCustomModuleObservedState) *pb.SecurityHealthAnalyticsCustomModule {
	if in == nil {
		return nil
	}
	out := &pb.SecurityHealthAnalyticsCustomModule{}
	// MISSING: Name
	// MISSING: DisplayName
	// MISSING: EnablementState
	// MISSING: UpdateTime
	// MISSING: LastEditor
	// MISSING: AncestorModule
	// MISSING: CustomConfig
	return out
}
func SecuritycenterSecurityHealthAnalyticsCustomModuleSpec_FromProto(mapCtx *direct.MapContext, in *pb.SecurityHealthAnalyticsCustomModule) *krm.SecuritycenterSecurityHealthAnalyticsCustomModuleSpec {
	if in == nil {
		return nil
	}
	out := &krm.SecuritycenterSecurityHealthAnalyticsCustomModuleSpec{}
	// MISSING: Name
	// MISSING: DisplayName
	// MISSING: EnablementState
	// MISSING: UpdateTime
	// MISSING: LastEditor
	// MISSING: AncestorModule
	// MISSING: CustomConfig
	return out
}
func SecuritycenterSecurityHealthAnalyticsCustomModuleSpec_ToProto(mapCtx *direct.MapContext, in *krm.SecuritycenterSecurityHealthAnalyticsCustomModuleSpec) *pb.SecurityHealthAnalyticsCustomModule {
	if in == nil {
		return nil
	}
	out := &pb.SecurityHealthAnalyticsCustomModule{}
	// MISSING: Name
	// MISSING: DisplayName
	// MISSING: EnablementState
	// MISSING: UpdateTime
	// MISSING: LastEditor
	// MISSING: AncestorModule
	// MISSING: CustomConfig
	return out
}
