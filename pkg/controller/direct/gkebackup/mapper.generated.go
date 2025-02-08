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

package gkebackup

import (
	refs "github.com/GoogleCloudPlatform/k8s-config-connector/apis/refs/v1beta1"
	pb "cloud.google.com/go/gkebackup/apiv1/gkebackuppb"
	krm "github.com/GoogleCloudPlatform/k8s-config-connector/apis/gkebackup/v1alpha1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct"
)
func GkebackupRestorePlanObservedState_FromProto(mapCtx *direct.MapContext, in *pb.RestorePlan) *krm.GkebackupRestorePlanObservedState {
	if in == nil {
		return nil
	}
	out := &krm.GkebackupRestorePlanObservedState{}
	// MISSING: Name
	// MISSING: Uid
	// MISSING: CreateTime
	// MISSING: UpdateTime
	// MISSING: Description
	// MISSING: BackupPlan
	// MISSING: Cluster
	// MISSING: RestoreConfig
	// MISSING: Labels
	// MISSING: Etag
	// MISSING: State
	// MISSING: StateReason
	return out
}
func GkebackupRestorePlanObservedState_ToProto(mapCtx *direct.MapContext, in *krm.GkebackupRestorePlanObservedState) *pb.RestorePlan {
	if in == nil {
		return nil
	}
	out := &pb.RestorePlan{}
	// MISSING: Name
	// MISSING: Uid
	// MISSING: CreateTime
	// MISSING: UpdateTime
	// MISSING: Description
	// MISSING: BackupPlan
	// MISSING: Cluster
	// MISSING: RestoreConfig
	// MISSING: Labels
	// MISSING: Etag
	// MISSING: State
	// MISSING: StateReason
	return out
}
func GkebackupRestorePlanSpec_FromProto(mapCtx *direct.MapContext, in *pb.RestorePlan) *krm.GkebackupRestorePlanSpec {
	if in == nil {
		return nil
	}
	out := &krm.GkebackupRestorePlanSpec{}
	// MISSING: Name
	// MISSING: Uid
	// MISSING: CreateTime
	// MISSING: UpdateTime
	// MISSING: Description
	// MISSING: BackupPlan
	// MISSING: Cluster
	// MISSING: RestoreConfig
	// MISSING: Labels
	// MISSING: Etag
	// MISSING: State
	// MISSING: StateReason
	return out
}
func GkebackupRestorePlanSpec_ToProto(mapCtx *direct.MapContext, in *krm.GkebackupRestorePlanSpec) *pb.RestorePlan {
	if in == nil {
		return nil
	}
	out := &pb.RestorePlan{}
	// MISSING: Name
	// MISSING: Uid
	// MISSING: CreateTime
	// MISSING: UpdateTime
	// MISSING: Description
	// MISSING: BackupPlan
	// MISSING: Cluster
	// MISSING: RestoreConfig
	// MISSING: Labels
	// MISSING: Etag
	// MISSING: State
	// MISSING: StateReason
	return out
}
func NamespacedName_FromProto(mapCtx *direct.MapContext, in *pb.NamespacedName) *krm.NamespacedName {
	if in == nil {
		return nil
	}
	out := &krm.NamespacedName{}
	out.Namespace = direct.LazyPtr(in.GetNamespace())
	out.Name = direct.LazyPtr(in.GetName())
	return out
}
func NamespacedName_ToProto(mapCtx *direct.MapContext, in *krm.NamespacedName) *pb.NamespacedName {
	if in == nil {
		return nil
	}
	out := &pb.NamespacedName{}
	out.Namespace = direct.ValueOf(in.Namespace)
	out.Name = direct.ValueOf(in.Name)
	return out
}
func NamespacedNames_FromProto(mapCtx *direct.MapContext, in *pb.NamespacedNames) *krm.NamespacedNames {
	if in == nil {
		return nil
	}
	out := &krm.NamespacedNames{}
	out.NamespacedNames = direct.Slice_FromProto(mapCtx, in.NamespacedNames, NamespacedName_FromProto)
	return out
}
func NamespacedNames_ToProto(mapCtx *direct.MapContext, in *krm.NamespacedNames) *pb.NamespacedNames {
	if in == nil {
		return nil
	}
	out := &pb.NamespacedNames{}
	out.NamespacedNames = direct.Slice_ToProto(mapCtx, in.NamespacedNames, NamespacedName_ToProto)
	return out
}
func Namespaces_FromProto(mapCtx *direct.MapContext, in *pb.Namespaces) *krm.Namespaces {
	if in == nil {
		return nil
	}
	out := &krm.Namespaces{}
	out.Namespaces = in.Namespaces
	return out
}
func Namespaces_ToProto(mapCtx *direct.MapContext, in *krm.Namespaces) *pb.Namespaces {
	if in == nil {
		return nil
	}
	out := &pb.Namespaces{}
	out.Namespaces = in.Namespaces
	return out
}
func RestoreConfig_FromProto(mapCtx *direct.MapContext, in *pb.RestoreConfig) *krm.RestoreConfig {
	if in == nil {
		return nil
	}
	out := &krm.RestoreConfig{}
	out.VolumeDataRestorePolicy = direct.Enum_FromProto(mapCtx, in.GetVolumeDataRestorePolicy())
	out.ClusterResourceConflictPolicy = direct.Enum_FromProto(mapCtx, in.GetClusterResourceConflictPolicy())
	out.NamespacedResourceRestoreMode = direct.Enum_FromProto(mapCtx, in.GetNamespacedResourceRestoreMode())
	out.ClusterResourceRestoreScope = RestoreConfig_ClusterResourceRestoreScope_FromProto(mapCtx, in.GetClusterResourceRestoreScope())
	out.AllNamespaces = direct.LazyPtr(in.GetAllNamespaces())
	out.SelectedNamespaces = Namespaces_FromProto(mapCtx, in.GetSelectedNamespaces())
	out.SelectedApplications = NamespacedNames_FromProto(mapCtx, in.GetSelectedApplications())
	out.NoNamespaces = direct.LazyPtr(in.GetNoNamespaces())
	out.ExcludedNamespaces = Namespaces_FromProto(mapCtx, in.GetExcludedNamespaces())
	out.SubstitutionRules = direct.Slice_FromProto(mapCtx, in.SubstitutionRules, RestoreConfig_SubstitutionRule_FromProto)
	out.TransformationRules = direct.Slice_FromProto(mapCtx, in.TransformationRules, RestoreConfig_TransformationRule_FromProto)
	out.VolumeDataRestorePolicyBindings = direct.Slice_FromProto(mapCtx, in.VolumeDataRestorePolicyBindings, RestoreConfig_VolumeDataRestorePolicyBinding_FromProto)
	out.RestoreOrder = RestoreConfig_RestoreOrder_FromProto(mapCtx, in.GetRestoreOrder())
	return out
}
func RestoreConfig_ToProto(mapCtx *direct.MapContext, in *krm.RestoreConfig) *pb.RestoreConfig {
	if in == nil {
		return nil
	}
	out := &pb.RestoreConfig{}
	out.VolumeDataRestorePolicy = direct.Enum_ToProto[pb.RestoreConfig_VolumeDataRestorePolicy](mapCtx, in.VolumeDataRestorePolicy)
	out.ClusterResourceConflictPolicy = direct.Enum_ToProto[pb.RestoreConfig_ClusterResourceConflictPolicy](mapCtx, in.ClusterResourceConflictPolicy)
	out.NamespacedResourceRestoreMode = direct.Enum_ToProto[pb.RestoreConfig_NamespacedResourceRestoreMode](mapCtx, in.NamespacedResourceRestoreMode)
	out.ClusterResourceRestoreScope = RestoreConfig_ClusterResourceRestoreScope_ToProto(mapCtx, in.ClusterResourceRestoreScope)
	if oneof := RestoreConfig_AllNamespaces_ToProto(mapCtx, in.AllNamespaces); oneof != nil {
		out.NamespacedResourceRestoreScope = oneof
	}
	if oneof := Namespaces_ToProto(mapCtx, in.SelectedNamespaces); oneof != nil {
		out.NamespacedResourceRestoreScope = &pb.RestoreConfig_SelectedNamespaces{SelectedNamespaces: oneof}
	}
	if oneof := NamespacedNames_ToProto(mapCtx, in.SelectedApplications); oneof != nil {
		out.NamespacedResourceRestoreScope = &pb.RestoreConfig_SelectedApplications{SelectedApplications: oneof}
	}
	if oneof := RestoreConfig_NoNamespaces_ToProto(mapCtx, in.NoNamespaces); oneof != nil {
		out.NamespacedResourceRestoreScope = oneof
	}
	if oneof := Namespaces_ToProto(mapCtx, in.ExcludedNamespaces); oneof != nil {
		out.NamespacedResourceRestoreScope = &pb.RestoreConfig_ExcludedNamespaces{ExcludedNamespaces: oneof}
	}
	out.SubstitutionRules = direct.Slice_ToProto(mapCtx, in.SubstitutionRules, RestoreConfig_SubstitutionRule_ToProto)
	out.TransformationRules = direct.Slice_ToProto(mapCtx, in.TransformationRules, RestoreConfig_TransformationRule_ToProto)
	out.VolumeDataRestorePolicyBindings = direct.Slice_ToProto(mapCtx, in.VolumeDataRestorePolicyBindings, RestoreConfig_VolumeDataRestorePolicyBinding_ToProto)
	out.RestoreOrder = RestoreConfig_RestoreOrder_ToProto(mapCtx, in.RestoreOrder)
	return out
}
func RestoreConfig_ClusterResourceRestoreScope_FromProto(mapCtx *direct.MapContext, in *pb.RestoreConfig_ClusterResourceRestoreScope) *krm.RestoreConfig_ClusterResourceRestoreScope {
	if in == nil {
		return nil
	}
	out := &krm.RestoreConfig_ClusterResourceRestoreScope{}
	out.SelectedGroupKinds = direct.Slice_FromProto(mapCtx, in.SelectedGroupKinds, RestoreConfig_GroupKind_FromProto)
	out.ExcludedGroupKinds = direct.Slice_FromProto(mapCtx, in.ExcludedGroupKinds, RestoreConfig_GroupKind_FromProto)
	out.AllGroupKinds = direct.LazyPtr(in.GetAllGroupKinds())
	out.NoGroupKinds = direct.LazyPtr(in.GetNoGroupKinds())
	return out
}
func RestoreConfig_ClusterResourceRestoreScope_ToProto(mapCtx *direct.MapContext, in *krm.RestoreConfig_ClusterResourceRestoreScope) *pb.RestoreConfig_ClusterResourceRestoreScope {
	if in == nil {
		return nil
	}
	out := &pb.RestoreConfig_ClusterResourceRestoreScope{}
	out.SelectedGroupKinds = direct.Slice_ToProto(mapCtx, in.SelectedGroupKinds, RestoreConfig_GroupKind_ToProto)
	out.ExcludedGroupKinds = direct.Slice_ToProto(mapCtx, in.ExcludedGroupKinds, RestoreConfig_GroupKind_ToProto)
	out.AllGroupKinds = direct.ValueOf(in.AllGroupKinds)
	out.NoGroupKinds = direct.ValueOf(in.NoGroupKinds)
	return out
}
func RestoreConfig_GroupKind_FromProto(mapCtx *direct.MapContext, in *pb.RestoreConfig_GroupKind) *krm.RestoreConfig_GroupKind {
	if in == nil {
		return nil
	}
	out := &krm.RestoreConfig_GroupKind{}
	out.ResourceGroup = direct.LazyPtr(in.GetResourceGroup())
	out.ResourceKind = direct.LazyPtr(in.GetResourceKind())
	return out
}
func RestoreConfig_GroupKind_ToProto(mapCtx *direct.MapContext, in *krm.RestoreConfig_GroupKind) *pb.RestoreConfig_GroupKind {
	if in == nil {
		return nil
	}
	out := &pb.RestoreConfig_GroupKind{}
	out.ResourceGroup = direct.ValueOf(in.ResourceGroup)
	out.ResourceKind = direct.ValueOf(in.ResourceKind)
	return out
}
func RestoreConfig_ResourceFilter_FromProto(mapCtx *direct.MapContext, in *pb.RestoreConfig_ResourceFilter) *krm.RestoreConfig_ResourceFilter {
	if in == nil {
		return nil
	}
	out := &krm.RestoreConfig_ResourceFilter{}
	out.Namespaces = in.Namespaces
	out.GroupKinds = direct.Slice_FromProto(mapCtx, in.GroupKinds, RestoreConfig_GroupKind_FromProto)
	out.JsonPath = direct.LazyPtr(in.GetJsonPath())
	return out
}
func RestoreConfig_ResourceFilter_ToProto(mapCtx *direct.MapContext, in *krm.RestoreConfig_ResourceFilter) *pb.RestoreConfig_ResourceFilter {
	if in == nil {
		return nil
	}
	out := &pb.RestoreConfig_ResourceFilter{}
	out.Namespaces = in.Namespaces
	out.GroupKinds = direct.Slice_ToProto(mapCtx, in.GroupKinds, RestoreConfig_GroupKind_ToProto)
	out.JsonPath = direct.ValueOf(in.JsonPath)
	return out
}
func RestoreConfig_RestoreOrder_FromProto(mapCtx *direct.MapContext, in *pb.RestoreConfig_RestoreOrder) *krm.RestoreConfig_RestoreOrder {
	if in == nil {
		return nil
	}
	out := &krm.RestoreConfig_RestoreOrder{}
	out.GroupKindDependencies = direct.Slice_FromProto(mapCtx, in.GroupKindDependencies, RestoreConfig_RestoreOrder_GroupKindDependency_FromProto)
	return out
}
func RestoreConfig_RestoreOrder_ToProto(mapCtx *direct.MapContext, in *krm.RestoreConfig_RestoreOrder) *pb.RestoreConfig_RestoreOrder {
	if in == nil {
		return nil
	}
	out := &pb.RestoreConfig_RestoreOrder{}
	out.GroupKindDependencies = direct.Slice_ToProto(mapCtx, in.GroupKindDependencies, RestoreConfig_RestoreOrder_GroupKindDependency_ToProto)
	return out
}
func RestoreConfig_RestoreOrder_GroupKindDependency_FromProto(mapCtx *direct.MapContext, in *pb.RestoreConfig_RestoreOrder_GroupKindDependency) *krm.RestoreConfig_RestoreOrder_GroupKindDependency {
	if in == nil {
		return nil
	}
	out := &krm.RestoreConfig_RestoreOrder_GroupKindDependency{}
	out.Satisfying = RestoreConfig_GroupKind_FromProto(mapCtx, in.GetSatisfying())
	out.Requiring = RestoreConfig_GroupKind_FromProto(mapCtx, in.GetRequiring())
	return out
}
func RestoreConfig_RestoreOrder_GroupKindDependency_ToProto(mapCtx *direct.MapContext, in *krm.RestoreConfig_RestoreOrder_GroupKindDependency) *pb.RestoreConfig_RestoreOrder_GroupKindDependency {
	if in == nil {
		return nil
	}
	out := &pb.RestoreConfig_RestoreOrder_GroupKindDependency{}
	out.Satisfying = RestoreConfig_GroupKind_ToProto(mapCtx, in.Satisfying)
	out.Requiring = RestoreConfig_GroupKind_ToProto(mapCtx, in.Requiring)
	return out
}
func RestoreConfig_SubstitutionRule_FromProto(mapCtx *direct.MapContext, in *pb.RestoreConfig_SubstitutionRule) *krm.RestoreConfig_SubstitutionRule {
	if in == nil {
		return nil
	}
	out := &krm.RestoreConfig_SubstitutionRule{}
	out.TargetNamespaces = in.TargetNamespaces
	out.TargetGroupKinds = direct.Slice_FromProto(mapCtx, in.TargetGroupKinds, RestoreConfig_GroupKind_FromProto)
	out.TargetJsonPath = direct.LazyPtr(in.GetTargetJsonPath())
	out.OriginalValuePattern = direct.LazyPtr(in.GetOriginalValuePattern())
	out.NewValue = direct.LazyPtr(in.GetNewValue())
	return out
}
func RestoreConfig_SubstitutionRule_ToProto(mapCtx *direct.MapContext, in *krm.RestoreConfig_SubstitutionRule) *pb.RestoreConfig_SubstitutionRule {
	if in == nil {
		return nil
	}
	out := &pb.RestoreConfig_SubstitutionRule{}
	out.TargetNamespaces = in.TargetNamespaces
	out.TargetGroupKinds = direct.Slice_ToProto(mapCtx, in.TargetGroupKinds, RestoreConfig_GroupKind_ToProto)
	out.TargetJsonPath = direct.ValueOf(in.TargetJsonPath)
	out.OriginalValuePattern = direct.ValueOf(in.OriginalValuePattern)
	out.NewValue = direct.ValueOf(in.NewValue)
	return out
}
func RestoreConfig_TransformationRule_FromProto(mapCtx *direct.MapContext, in *pb.RestoreConfig_TransformationRule) *krm.RestoreConfig_TransformationRule {
	if in == nil {
		return nil
	}
	out := &krm.RestoreConfig_TransformationRule{}
	out.FieldActions = direct.Slice_FromProto(mapCtx, in.FieldActions, RestoreConfig_TransformationRuleAction_FromProto)
	out.ResourceFilter = RestoreConfig_ResourceFilter_FromProto(mapCtx, in.GetResourceFilter())
	out.Description = direct.LazyPtr(in.GetDescription())
	return out
}
func RestoreConfig_TransformationRule_ToProto(mapCtx *direct.MapContext, in *krm.RestoreConfig_TransformationRule) *pb.RestoreConfig_TransformationRule {
	if in == nil {
		return nil
	}
	out := &pb.RestoreConfig_TransformationRule{}
	out.FieldActions = direct.Slice_ToProto(mapCtx, in.FieldActions, RestoreConfig_TransformationRuleAction_ToProto)
	out.ResourceFilter = RestoreConfig_ResourceFilter_ToProto(mapCtx, in.ResourceFilter)
	out.Description = direct.ValueOf(in.Description)
	return out
}
func RestoreConfig_TransformationRuleAction_FromProto(mapCtx *direct.MapContext, in *pb.RestoreConfig_TransformationRuleAction) *krm.RestoreConfig_TransformationRuleAction {
	if in == nil {
		return nil
	}
	out := &krm.RestoreConfig_TransformationRuleAction{}
	out.Op = direct.Enum_FromProto(mapCtx, in.GetOp())
	out.FromPath = direct.LazyPtr(in.GetFromPath())
	out.Path = direct.LazyPtr(in.GetPath())
	out.Value = direct.LazyPtr(in.GetValue())
	return out
}
func RestoreConfig_TransformationRuleAction_ToProto(mapCtx *direct.MapContext, in *krm.RestoreConfig_TransformationRuleAction) *pb.RestoreConfig_TransformationRuleAction {
	if in == nil {
		return nil
	}
	out := &pb.RestoreConfig_TransformationRuleAction{}
	out.Op = direct.Enum_ToProto[pb.RestoreConfig_TransformationRuleAction_Op](mapCtx, in.Op)
	out.FromPath = direct.ValueOf(in.FromPath)
	out.Path = direct.ValueOf(in.Path)
	out.Value = direct.ValueOf(in.Value)
	return out
}
func RestoreConfig_VolumeDataRestorePolicyBinding_FromProto(mapCtx *direct.MapContext, in *pb.RestoreConfig_VolumeDataRestorePolicyBinding) *krm.RestoreConfig_VolumeDataRestorePolicyBinding {
	if in == nil {
		return nil
	}
	out := &krm.RestoreConfig_VolumeDataRestorePolicyBinding{}
	out.Policy = direct.Enum_FromProto(mapCtx, in.GetPolicy())
	out.VolumeType = direct.Enum_FromProto(mapCtx, in.GetVolumeType())
	return out
}
func RestoreConfig_VolumeDataRestorePolicyBinding_ToProto(mapCtx *direct.MapContext, in *krm.RestoreConfig_VolumeDataRestorePolicyBinding) *pb.RestoreConfig_VolumeDataRestorePolicyBinding {
	if in == nil {
		return nil
	}
	out := &pb.RestoreConfig_VolumeDataRestorePolicyBinding{}
	out.Policy = direct.Enum_ToProto[pb.RestoreConfig_VolumeDataRestorePolicy](mapCtx, in.Policy)
	if oneof := RestoreConfig_VolumeDataRestorePolicyBinding_VolumeType_ToProto(mapCtx, in.VolumeType); oneof != nil {
		out.Scope = oneof
	}
	return out
}
func RestorePlan_FromProto(mapCtx *direct.MapContext, in *pb.RestorePlan) *krm.RestorePlan {
	if in == nil {
		return nil
	}
	out := &krm.RestorePlan{}
	// MISSING: Name
	// MISSING: Uid
	// MISSING: CreateTime
	// MISSING: UpdateTime
	out.Description = direct.LazyPtr(in.GetDescription())
	out.BackupPlan = direct.LazyPtr(in.GetBackupPlan())
	out.Cluster = direct.LazyPtr(in.GetCluster())
	out.RestoreConfig = RestoreConfig_FromProto(mapCtx, in.GetRestoreConfig())
	out.Labels = in.Labels
	// MISSING: Etag
	// MISSING: State
	// MISSING: StateReason
	return out
}
func RestorePlan_ToProto(mapCtx *direct.MapContext, in *krm.RestorePlan) *pb.RestorePlan {
	if in == nil {
		return nil
	}
	out := &pb.RestorePlan{}
	// MISSING: Name
	// MISSING: Uid
	// MISSING: CreateTime
	// MISSING: UpdateTime
	out.Description = direct.ValueOf(in.Description)
	out.BackupPlan = direct.ValueOf(in.BackupPlan)
	out.Cluster = direct.ValueOf(in.Cluster)
	out.RestoreConfig = RestoreConfig_ToProto(mapCtx, in.RestoreConfig)
	out.Labels = in.Labels
	// MISSING: Etag
	// MISSING: State
	// MISSING: StateReason
	return out
}
func RestorePlanObservedState_FromProto(mapCtx *direct.MapContext, in *pb.RestorePlan) *krm.RestorePlanObservedState {
	if in == nil {
		return nil
	}
	out := &krm.RestorePlanObservedState{}
	out.Name = direct.LazyPtr(in.GetName())
	out.Uid = direct.LazyPtr(in.GetUid())
	out.CreateTime = direct.StringTimestamp_FromProto(mapCtx, in.GetCreateTime())
	out.UpdateTime = direct.StringTimestamp_FromProto(mapCtx, in.GetUpdateTime())
	// MISSING: Description
	// MISSING: BackupPlan
	// MISSING: Cluster
	// MISSING: RestoreConfig
	// MISSING: Labels
	out.Etag = direct.LazyPtr(in.GetEtag())
	out.State = direct.Enum_FromProto(mapCtx, in.GetState())
	out.StateReason = direct.LazyPtr(in.GetStateReason())
	return out
}
func RestorePlanObservedState_ToProto(mapCtx *direct.MapContext, in *krm.RestorePlanObservedState) *pb.RestorePlan {
	if in == nil {
		return nil
	}
	out := &pb.RestorePlan{}
	out.Name = direct.ValueOf(in.Name)
	out.Uid = direct.ValueOf(in.Uid)
	out.CreateTime = direct.StringTimestamp_ToProto(mapCtx, in.CreateTime)
	out.UpdateTime = direct.StringTimestamp_ToProto(mapCtx, in.UpdateTime)
	// MISSING: Description
	// MISSING: BackupPlan
	// MISSING: Cluster
	// MISSING: RestoreConfig
	// MISSING: Labels
	out.Etag = direct.ValueOf(in.Etag)
	out.State = direct.Enum_ToProto[pb.RestorePlan_State](mapCtx, in.State)
	out.StateReason = direct.ValueOf(in.StateReason)
	return out
}
