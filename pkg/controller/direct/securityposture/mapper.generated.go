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

package securityposture

import (
	refs "github.com/GoogleCloudPlatform/k8s-config-connector/apis/refs/v1beta1"
	pb "cloud.google.com/go/securityposture/apiv1/securityposturepb"
	krm "github.com/GoogleCloudPlatform/k8s-config-connector/apis/securityposture/v1alpha1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct"
)
func Constraint_FromProto(mapCtx *direct.MapContext, in *pb.Constraint) *krm.Constraint {
	if in == nil {
		return nil
	}
	out := &krm.Constraint{}
	out.SecurityHealthAnalyticsModule = SecurityHealthAnalyticsModule_FromProto(mapCtx, in.GetSecurityHealthAnalyticsModule())
	out.SecurityHealthAnalyticsCustomModule = SecurityHealthAnalyticsCustomModule_FromProto(mapCtx, in.GetSecurityHealthAnalyticsCustomModule())
	out.OrgPolicyConstraint = OrgPolicyConstraint_FromProto(mapCtx, in.GetOrgPolicyConstraint())
	out.OrgPolicyConstraintCustom = OrgPolicyConstraintCustom_FromProto(mapCtx, in.GetOrgPolicyConstraintCustom())
	return out
}
func Constraint_ToProto(mapCtx *direct.MapContext, in *krm.Constraint) *pb.Constraint {
	if in == nil {
		return nil
	}
	out := &pb.Constraint{}
	if oneof := SecurityHealthAnalyticsModule_ToProto(mapCtx, in.SecurityHealthAnalyticsModule); oneof != nil {
		out.Implementation = &pb.Constraint_SecurityHealthAnalyticsModule{SecurityHealthAnalyticsModule: oneof}
	}
	if oneof := SecurityHealthAnalyticsCustomModule_ToProto(mapCtx, in.SecurityHealthAnalyticsCustomModule); oneof != nil {
		out.Implementation = &pb.Constraint_SecurityHealthAnalyticsCustomModule{SecurityHealthAnalyticsCustomModule: oneof}
	}
	if oneof := OrgPolicyConstraint_ToProto(mapCtx, in.OrgPolicyConstraint); oneof != nil {
		out.Implementation = &pb.Constraint_OrgPolicyConstraint{OrgPolicyConstraint: oneof}
	}
	if oneof := OrgPolicyConstraintCustom_ToProto(mapCtx, in.OrgPolicyConstraintCustom); oneof != nil {
		out.Implementation = &pb.Constraint_OrgPolicyConstraintCustom{OrgPolicyConstraintCustom: oneof}
	}
	return out
}
func ConstraintObservedState_FromProto(mapCtx *direct.MapContext, in *pb.Constraint) *krm.ConstraintObservedState {
	if in == nil {
		return nil
	}
	out := &krm.ConstraintObservedState{}
	// MISSING: SecurityHealthAnalyticsModule
	out.SecurityHealthAnalyticsCustomModule = SecurityHealthAnalyticsCustomModuleObservedState_FromProto(mapCtx, in.GetSecurityHealthAnalyticsCustomModule())
	// MISSING: OrgPolicyConstraint
	out.OrgPolicyConstraintCustom = OrgPolicyConstraintCustomObservedState_FromProto(mapCtx, in.GetOrgPolicyConstraintCustom())
	return out
}
func ConstraintObservedState_ToProto(mapCtx *direct.MapContext, in *krm.ConstraintObservedState) *pb.Constraint {
	if in == nil {
		return nil
	}
	out := &pb.Constraint{}
	// MISSING: SecurityHealthAnalyticsModule
	if oneof := SecurityHealthAnalyticsCustomModuleObservedState_ToProto(mapCtx, in.SecurityHealthAnalyticsCustomModule); oneof != nil {
		out.Implementation = &pb.Constraint_SecurityHealthAnalyticsCustomModule{SecurityHealthAnalyticsCustomModule: oneof}
	}
	// MISSING: OrgPolicyConstraint
	if oneof := OrgPolicyConstraintCustomObservedState_ToProto(mapCtx, in.OrgPolicyConstraintCustom); oneof != nil {
		out.Implementation = &pb.Constraint_OrgPolicyConstraintCustom{OrgPolicyConstraintCustom: oneof}
	}
	return out
}
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
func CustomConstraint_FromProto(mapCtx *direct.MapContext, in *pb.CustomConstraint) *krm.CustomConstraint {
	if in == nil {
		return nil
	}
	out := &krm.CustomConstraint{}
	out.Name = direct.LazyPtr(in.GetName())
	out.ResourceTypes = in.ResourceTypes
	out.MethodTypes = direct.EnumSlice_FromProto(mapCtx, in.MethodTypes)
	out.Condition = direct.LazyPtr(in.GetCondition())
	out.ActionType = direct.Enum_FromProto(mapCtx, in.GetActionType())
	out.DisplayName = direct.LazyPtr(in.GetDisplayName())
	out.Description = direct.LazyPtr(in.GetDescription())
	// MISSING: UpdateTime
	return out
}
func CustomConstraint_ToProto(mapCtx *direct.MapContext, in *krm.CustomConstraint) *pb.CustomConstraint {
	if in == nil {
		return nil
	}
	out := &pb.CustomConstraint{}
	out.Name = direct.ValueOf(in.Name)
	out.ResourceTypes = in.ResourceTypes
	out.MethodTypes = direct.EnumSlice_ToProto[pb.CustomConstraint_MethodType](mapCtx, in.MethodTypes)
	out.Condition = direct.ValueOf(in.Condition)
	out.ActionType = direct.Enum_ToProto[pb.CustomConstraint_ActionType](mapCtx, in.ActionType)
	out.DisplayName = direct.ValueOf(in.DisplayName)
	out.Description = direct.ValueOf(in.Description)
	// MISSING: UpdateTime
	return out
}
func CustomConstraintObservedState_FromProto(mapCtx *direct.MapContext, in *pb.CustomConstraint) *krm.CustomConstraintObservedState {
	if in == nil {
		return nil
	}
	out := &krm.CustomConstraintObservedState{}
	// MISSING: Name
	// MISSING: ResourceTypes
	// MISSING: MethodTypes
	// MISSING: Condition
	// MISSING: ActionType
	// MISSING: DisplayName
	// MISSING: Description
	out.UpdateTime = direct.StringTimestamp_FromProto(mapCtx, in.GetUpdateTime())
	return out
}
func CustomConstraintObservedState_ToProto(mapCtx *direct.MapContext, in *krm.CustomConstraintObservedState) *pb.CustomConstraint {
	if in == nil {
		return nil
	}
	out := &pb.CustomConstraint{}
	// MISSING: Name
	// MISSING: ResourceTypes
	// MISSING: MethodTypes
	// MISSING: Condition
	// MISSING: ActionType
	// MISSING: DisplayName
	// MISSING: Description
	out.UpdateTime = direct.StringTimestamp_ToProto(mapCtx, in.UpdateTime)
	return out
}
func OrgPolicyConstraint_FromProto(mapCtx *direct.MapContext, in *pb.OrgPolicyConstraint) *krm.OrgPolicyConstraint {
	if in == nil {
		return nil
	}
	out := &krm.OrgPolicyConstraint{}
	out.CannedConstraintID = direct.LazyPtr(in.GetCannedConstraintId())
	out.PolicyRules = direct.Slice_FromProto(mapCtx, in.PolicyRules, PolicyRule_FromProto)
	return out
}
func OrgPolicyConstraint_ToProto(mapCtx *direct.MapContext, in *krm.OrgPolicyConstraint) *pb.OrgPolicyConstraint {
	if in == nil {
		return nil
	}
	out := &pb.OrgPolicyConstraint{}
	out.CannedConstraintId = direct.ValueOf(in.CannedConstraintID)
	out.PolicyRules = direct.Slice_ToProto(mapCtx, in.PolicyRules, PolicyRule_ToProto)
	return out
}
func OrgPolicyConstraintCustom_FromProto(mapCtx *direct.MapContext, in *pb.OrgPolicyConstraintCustom) *krm.OrgPolicyConstraintCustom {
	if in == nil {
		return nil
	}
	out := &krm.OrgPolicyConstraintCustom{}
	out.CustomConstraint = CustomConstraint_FromProto(mapCtx, in.GetCustomConstraint())
	out.PolicyRules = direct.Slice_FromProto(mapCtx, in.PolicyRules, PolicyRule_FromProto)
	return out
}
func OrgPolicyConstraintCustom_ToProto(mapCtx *direct.MapContext, in *krm.OrgPolicyConstraintCustom) *pb.OrgPolicyConstraintCustom {
	if in == nil {
		return nil
	}
	out := &pb.OrgPolicyConstraintCustom{}
	out.CustomConstraint = CustomConstraint_ToProto(mapCtx, in.CustomConstraint)
	out.PolicyRules = direct.Slice_ToProto(mapCtx, in.PolicyRules, PolicyRule_ToProto)
	return out
}
func OrgPolicyConstraintCustomObservedState_FromProto(mapCtx *direct.MapContext, in *pb.OrgPolicyConstraintCustom) *krm.OrgPolicyConstraintCustomObservedState {
	if in == nil {
		return nil
	}
	out := &krm.OrgPolicyConstraintCustomObservedState{}
	out.CustomConstraint = CustomConstraintObservedState_FromProto(mapCtx, in.GetCustomConstraint())
	// MISSING: PolicyRules
	return out
}
func OrgPolicyConstraintCustomObservedState_ToProto(mapCtx *direct.MapContext, in *krm.OrgPolicyConstraintCustomObservedState) *pb.OrgPolicyConstraintCustom {
	if in == nil {
		return nil
	}
	out := &pb.OrgPolicyConstraintCustom{}
	out.CustomConstraint = CustomConstraintObservedState_ToProto(mapCtx, in.CustomConstraint)
	// MISSING: PolicyRules
	return out
}
func Policy_FromProto(mapCtx *direct.MapContext, in *pb.Policy) *krm.Policy {
	if in == nil {
		return nil
	}
	out := &krm.Policy{}
	out.PolicyID = direct.LazyPtr(in.GetPolicyId())
	out.ComplianceStandards = direct.Slice_FromProto(mapCtx, in.ComplianceStandards, Policy_ComplianceStandard_FromProto)
	out.Constraint = Constraint_FromProto(mapCtx, in.GetConstraint())
	out.Description = direct.LazyPtr(in.GetDescription())
	return out
}
func Policy_ToProto(mapCtx *direct.MapContext, in *krm.Policy) *pb.Policy {
	if in == nil {
		return nil
	}
	out := &pb.Policy{}
	out.PolicyId = direct.ValueOf(in.PolicyID)
	out.ComplianceStandards = direct.Slice_ToProto(mapCtx, in.ComplianceStandards, Policy_ComplianceStandard_ToProto)
	out.Constraint = Constraint_ToProto(mapCtx, in.Constraint)
	out.Description = direct.ValueOf(in.Description)
	return out
}
func PolicyObservedState_FromProto(mapCtx *direct.MapContext, in *pb.Policy) *krm.PolicyObservedState {
	if in == nil {
		return nil
	}
	out := &krm.PolicyObservedState{}
	// MISSING: PolicyID
	// MISSING: ComplianceStandards
	out.Constraint = ConstraintObservedState_FromProto(mapCtx, in.GetConstraint())
	// MISSING: Description
	return out
}
func PolicyObservedState_ToProto(mapCtx *direct.MapContext, in *krm.PolicyObservedState) *pb.Policy {
	if in == nil {
		return nil
	}
	out := &pb.Policy{}
	// MISSING: PolicyID
	// MISSING: ComplianceStandards
	out.Constraint = ConstraintObservedState_ToProto(mapCtx, in.Constraint)
	// MISSING: Description
	return out
}
func PolicyRule_FromProto(mapCtx *direct.MapContext, in *pb.PolicyRule) *krm.PolicyRule {
	if in == nil {
		return nil
	}
	out := &krm.PolicyRule{}
	out.Values = PolicyRule_StringValues_FromProto(mapCtx, in.GetValues())
	out.AllowAll = direct.LazyPtr(in.GetAllowAll())
	out.DenyAll = direct.LazyPtr(in.GetDenyAll())
	out.Enforce = direct.LazyPtr(in.GetEnforce())
	out.Condition = Expr_FromProto(mapCtx, in.GetCondition())
	return out
}
func PolicyRule_ToProto(mapCtx *direct.MapContext, in *krm.PolicyRule) *pb.PolicyRule {
	if in == nil {
		return nil
	}
	out := &pb.PolicyRule{}
	if oneof := PolicyRule_StringValues_ToProto(mapCtx, in.Values); oneof != nil {
		out.Kind = &pb.PolicyRule_Values{Values: oneof}
	}
	if oneof := PolicyRule_AllowAll_ToProto(mapCtx, in.AllowAll); oneof != nil {
		out.Kind = oneof
	}
	if oneof := PolicyRule_DenyAll_ToProto(mapCtx, in.DenyAll); oneof != nil {
		out.Kind = oneof
	}
	if oneof := PolicyRule_Enforce_ToProto(mapCtx, in.Enforce); oneof != nil {
		out.Kind = oneof
	}
	out.Condition = Expr_ToProto(mapCtx, in.Condition)
	return out
}
func PolicyRule_StringValues_FromProto(mapCtx *direct.MapContext, in *pb.PolicyRule_StringValues) *krm.PolicyRule_StringValues {
	if in == nil {
		return nil
	}
	out := &krm.PolicyRule_StringValues{}
	out.AllowedValues = in.AllowedValues
	out.DeniedValues = in.DeniedValues
	return out
}
func PolicyRule_StringValues_ToProto(mapCtx *direct.MapContext, in *krm.PolicyRule_StringValues) *pb.PolicyRule_StringValues {
	if in == nil {
		return nil
	}
	out := &pb.PolicyRule_StringValues{}
	out.AllowedValues = in.AllowedValues
	out.DeniedValues = in.DeniedValues
	return out
}
func PolicySet_FromProto(mapCtx *direct.MapContext, in *pb.PolicySet) *krm.PolicySet {
	if in == nil {
		return nil
	}
	out := &krm.PolicySet{}
	out.PolicySetID = direct.LazyPtr(in.GetPolicySetId())
	out.Description = direct.LazyPtr(in.GetDescription())
	out.Policies = direct.Slice_FromProto(mapCtx, in.Policies, Policy_FromProto)
	return out
}
func PolicySet_ToProto(mapCtx *direct.MapContext, in *krm.PolicySet) *pb.PolicySet {
	if in == nil {
		return nil
	}
	out := &pb.PolicySet{}
	out.PolicySetId = direct.ValueOf(in.PolicySetID)
	out.Description = direct.ValueOf(in.Description)
	out.Policies = direct.Slice_ToProto(mapCtx, in.Policies, Policy_ToProto)
	return out
}
func PolicySetObservedState_FromProto(mapCtx *direct.MapContext, in *pb.PolicySet) *krm.PolicySetObservedState {
	if in == nil {
		return nil
	}
	out := &krm.PolicySetObservedState{}
	// MISSING: PolicySetID
	// MISSING: Description
	out.Policies = direct.Slice_FromProto(mapCtx, in.Policies, PolicyObservedState_FromProto)
	return out
}
func PolicySetObservedState_ToProto(mapCtx *direct.MapContext, in *krm.PolicySetObservedState) *pb.PolicySet {
	if in == nil {
		return nil
	}
	out := &pb.PolicySet{}
	// MISSING: PolicySetID
	// MISSING: Description
	out.Policies = direct.Slice_ToProto(mapCtx, in.Policies, PolicyObservedState_ToProto)
	return out
}
func Policy_ComplianceStandard_FromProto(mapCtx *direct.MapContext, in *pb.Policy_ComplianceStandard) *krm.Policy_ComplianceStandard {
	if in == nil {
		return nil
	}
	out := &krm.Policy_ComplianceStandard{}
	out.Standard = direct.LazyPtr(in.GetStandard())
	out.Control = direct.LazyPtr(in.GetControl())
	return out
}
func Policy_ComplianceStandard_ToProto(mapCtx *direct.MapContext, in *krm.Policy_ComplianceStandard) *pb.Policy_ComplianceStandard {
	if in == nil {
		return nil
	}
	out := &pb.Policy_ComplianceStandard{}
	out.Standard = direct.ValueOf(in.Standard)
	out.Control = direct.ValueOf(in.Control)
	return out
}
func Posture_FromProto(mapCtx *direct.MapContext, in *pb.Posture) *krm.Posture {
	if in == nil {
		return nil
	}
	out := &krm.Posture{}
	out.Name = direct.LazyPtr(in.GetName())
	out.State = direct.Enum_FromProto(mapCtx, in.GetState())
	// MISSING: RevisionID
	// MISSING: CreateTime
	// MISSING: UpdateTime
	out.Description = direct.LazyPtr(in.GetDescription())
	out.PolicySets = direct.Slice_FromProto(mapCtx, in.PolicySets, PolicySet_FromProto)
	out.Etag = direct.LazyPtr(in.GetEtag())
	out.Annotations = in.Annotations
	// MISSING: Reconciling
	return out
}
func Posture_ToProto(mapCtx *direct.MapContext, in *krm.Posture) *pb.Posture {
	if in == nil {
		return nil
	}
	out := &pb.Posture{}
	out.Name = direct.ValueOf(in.Name)
	out.State = direct.Enum_ToProto[pb.Posture_State](mapCtx, in.State)
	// MISSING: RevisionID
	// MISSING: CreateTime
	// MISSING: UpdateTime
	out.Description = direct.ValueOf(in.Description)
	out.PolicySets = direct.Slice_ToProto(mapCtx, in.PolicySets, PolicySet_ToProto)
	out.Etag = direct.ValueOf(in.Etag)
	out.Annotations = in.Annotations
	// MISSING: Reconciling
	return out
}
func PostureObservedState_FromProto(mapCtx *direct.MapContext, in *pb.Posture) *krm.PostureObservedState {
	if in == nil {
		return nil
	}
	out := &krm.PostureObservedState{}
	// MISSING: Name
	// MISSING: State
	out.RevisionID = direct.LazyPtr(in.GetRevisionId())
	out.CreateTime = direct.StringTimestamp_FromProto(mapCtx, in.GetCreateTime())
	out.UpdateTime = direct.StringTimestamp_FromProto(mapCtx, in.GetUpdateTime())
	// MISSING: Description
	out.PolicySets = direct.Slice_FromProto(mapCtx, in.PolicySets, PolicySetObservedState_FromProto)
	// MISSING: Etag
	// MISSING: Annotations
	out.Reconciling = direct.LazyPtr(in.GetReconciling())
	return out
}
func PostureObservedState_ToProto(mapCtx *direct.MapContext, in *krm.PostureObservedState) *pb.Posture {
	if in == nil {
		return nil
	}
	out := &pb.Posture{}
	// MISSING: Name
	// MISSING: State
	out.RevisionId = direct.ValueOf(in.RevisionID)
	out.CreateTime = direct.StringTimestamp_ToProto(mapCtx, in.CreateTime)
	out.UpdateTime = direct.StringTimestamp_ToProto(mapCtx, in.UpdateTime)
	// MISSING: Description
	out.PolicySets = direct.Slice_ToProto(mapCtx, in.PolicySets, PolicySetObservedState_ToProto)
	// MISSING: Etag
	// MISSING: Annotations
	out.Reconciling = direct.ValueOf(in.Reconciling)
	return out
}
func SecurityHealthAnalyticsCustomModule_FromProto(mapCtx *direct.MapContext, in *pb.SecurityHealthAnalyticsCustomModule) *krm.SecurityHealthAnalyticsCustomModule {
	if in == nil {
		return nil
	}
	out := &krm.SecurityHealthAnalyticsCustomModule{}
	// MISSING: ID
	out.DisplayName = direct.LazyPtr(in.GetDisplayName())
	out.Config = CustomConfig_FromProto(mapCtx, in.GetConfig())
	out.ModuleEnablementState = direct.Enum_FromProto(mapCtx, in.GetModuleEnablementState())
	return out
}
func SecurityHealthAnalyticsCustomModule_ToProto(mapCtx *direct.MapContext, in *krm.SecurityHealthAnalyticsCustomModule) *pb.SecurityHealthAnalyticsCustomModule {
	if in == nil {
		return nil
	}
	out := &pb.SecurityHealthAnalyticsCustomModule{}
	// MISSING: ID
	out.DisplayName = direct.ValueOf(in.DisplayName)
	out.Config = CustomConfig_ToProto(mapCtx, in.Config)
	out.ModuleEnablementState = direct.Enum_ToProto[pb.EnablementState](mapCtx, in.ModuleEnablementState)
	return out
}
func SecurityHealthAnalyticsCustomModuleObservedState_FromProto(mapCtx *direct.MapContext, in *pb.SecurityHealthAnalyticsCustomModule) *krm.SecurityHealthAnalyticsCustomModuleObservedState {
	if in == nil {
		return nil
	}
	out := &krm.SecurityHealthAnalyticsCustomModuleObservedState{}
	out.ID = direct.LazyPtr(in.GetId())
	// MISSING: DisplayName
	// MISSING: Config
	// MISSING: ModuleEnablementState
	return out
}
func SecurityHealthAnalyticsCustomModuleObservedState_ToProto(mapCtx *direct.MapContext, in *krm.SecurityHealthAnalyticsCustomModuleObservedState) *pb.SecurityHealthAnalyticsCustomModule {
	if in == nil {
		return nil
	}
	out := &pb.SecurityHealthAnalyticsCustomModule{}
	out.Id = direct.ValueOf(in.ID)
	// MISSING: DisplayName
	// MISSING: Config
	// MISSING: ModuleEnablementState
	return out
}
func SecurityHealthAnalyticsModule_FromProto(mapCtx *direct.MapContext, in *pb.SecurityHealthAnalyticsModule) *krm.SecurityHealthAnalyticsModule {
	if in == nil {
		return nil
	}
	out := &krm.SecurityHealthAnalyticsModule{}
	out.ModuleName = direct.LazyPtr(in.GetModuleName())
	out.ModuleEnablementState = direct.Enum_FromProto(mapCtx, in.GetModuleEnablementState())
	return out
}
func SecurityHealthAnalyticsModule_ToProto(mapCtx *direct.MapContext, in *krm.SecurityHealthAnalyticsModule) *pb.SecurityHealthAnalyticsModule {
	if in == nil {
		return nil
	}
	out := &pb.SecurityHealthAnalyticsModule{}
	out.ModuleName = direct.ValueOf(in.ModuleName)
	out.ModuleEnablementState = direct.Enum_ToProto[pb.EnablementState](mapCtx, in.ModuleEnablementState)
	return out
}
func SecurityposturePostureObservedState_FromProto(mapCtx *direct.MapContext, in *pb.Posture) *krm.SecurityposturePostureObservedState {
	if in == nil {
		return nil
	}
	out := &krm.SecurityposturePostureObservedState{}
	// MISSING: Name
	// MISSING: State
	// MISSING: RevisionID
	// MISSING: CreateTime
	// MISSING: UpdateTime
	// MISSING: Description
	// MISSING: PolicySets
	// MISSING: Etag
	// MISSING: Annotations
	// MISSING: Reconciling
	return out
}
func SecurityposturePostureObservedState_ToProto(mapCtx *direct.MapContext, in *krm.SecurityposturePostureObservedState) *pb.Posture {
	if in == nil {
		return nil
	}
	out := &pb.Posture{}
	// MISSING: Name
	// MISSING: State
	// MISSING: RevisionID
	// MISSING: CreateTime
	// MISSING: UpdateTime
	// MISSING: Description
	// MISSING: PolicySets
	// MISSING: Etag
	// MISSING: Annotations
	// MISSING: Reconciling
	return out
}
func SecurityposturePostureSpec_FromProto(mapCtx *direct.MapContext, in *pb.Posture) *krm.SecurityposturePostureSpec {
	if in == nil {
		return nil
	}
	out := &krm.SecurityposturePostureSpec{}
	// MISSING: Name
	// MISSING: State
	// MISSING: RevisionID
	// MISSING: CreateTime
	// MISSING: UpdateTime
	// MISSING: Description
	// MISSING: PolicySets
	// MISSING: Etag
	// MISSING: Annotations
	// MISSING: Reconciling
	return out
}
func SecurityposturePostureSpec_ToProto(mapCtx *direct.MapContext, in *krm.SecurityposturePostureSpec) *pb.Posture {
	if in == nil {
		return nil
	}
	out := &pb.Posture{}
	// MISSING: Name
	// MISSING: State
	// MISSING: RevisionID
	// MISSING: CreateTime
	// MISSING: UpdateTime
	// MISSING: Description
	// MISSING: PolicySets
	// MISSING: Etag
	// MISSING: Annotations
	// MISSING: Reconciling
	return out
}
