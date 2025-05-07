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
	pb "cloud.google.com/go/gkebackup/apiv1/gkebackuppb"

	container "github.com/GoogleCloudPlatform/k8s-config-connector/apis/container/v1beta1"
	krm "github.com/GoogleCloudPlatform/k8s-config-connector/apis/gkebackup/v1alpha1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct"
)

func GKEBackupRestorePlanObservedState_FromProto(mapCtx *direct.MapContext, in *pb.RestorePlan) *krm.GKEBackupRestorePlanObservedState {
	if in == nil {
		return nil
	}
	out := &krm.GKEBackupRestorePlanObservedState{}
	// MISSING: Name
	out.UID = direct.LazyPtr(in.GetUid())
	out.CreateTime = direct.StringTimestamp_FromProto(mapCtx, in.GetCreateTime())
	out.UpdateTime = direct.StringTimestamp_FromProto(mapCtx, in.GetUpdateTime())
	out.Etag = direct.LazyPtr(in.GetEtag())
	out.State = direct.Enum_FromProto(mapCtx, in.GetState())
	out.StateReason = direct.LazyPtr(in.GetStateReason())
	return out
}
func GKEBackupRestorePlanObservedState_ToProto(mapCtx *direct.MapContext, in *krm.GKEBackupRestorePlanObservedState) *pb.RestorePlan {
	if in == nil {
		return nil
	}
	out := &pb.RestorePlan{}
	// MISSING: Name
	out.Uid = direct.ValueOf(in.UID)
	out.CreateTime = direct.StringTimestamp_ToProto(mapCtx, in.CreateTime)
	out.UpdateTime = direct.StringTimestamp_ToProto(mapCtx, in.UpdateTime)
	out.Etag = direct.ValueOf(in.Etag)
	out.State = direct.Enum_ToProto[pb.RestorePlan_State](mapCtx, in.State)
	out.StateReason = direct.ValueOf(in.StateReason)
	return out
}
func RestoreConfig_ResourceFilter_FromProto(mapCtx *direct.MapContext, in *pb.RestoreConfig_ResourceFilter) *krm.RestoreConfig_ResourceFilter {
	if in == nil {
		return nil
	}
	out := &krm.RestoreConfig_ResourceFilter{}
	out.Namespaces = in.Namespaces
	out.GroupKinds = direct.Slice_FromProto(mapCtx, in.GroupKinds, RestoreConfig_GroupKind_FromProto)
	out.JSONPath = direct.LazyPtr(in.GetJsonPath())
	return out
}
func RestoreConfig_ResourceFilter_ToProto(mapCtx *direct.MapContext, in *krm.RestoreConfig_ResourceFilter) *pb.RestoreConfig_ResourceFilter {
	if in == nil {
		return nil
	}
	out := &pb.RestoreConfig_ResourceFilter{}
	out.Namespaces = in.Namespaces
	out.GroupKinds = direct.Slice_ToProto(mapCtx, in.GroupKinds, RestoreConfig_GroupKind_ToProto)
	out.JsonPath = direct.ValueOf(in.JSONPath)
	return out
}
func RestoreConfig_SubstitutionRule_FromProto(mapCtx *direct.MapContext, in *pb.RestoreConfig_SubstitutionRule) *krm.RestoreConfig_SubstitutionRule {
	if in == nil {
		return nil
	}
	out := &krm.RestoreConfig_SubstitutionRule{}
	out.TargetNamespaces = in.TargetNamespaces
	out.TargetGroupKinds = direct.Slice_FromProto(mapCtx, in.TargetGroupKinds, RestoreConfig_GroupKind_FromProto)
	out.TargetJSONPath = direct.LazyPtr(in.GetTargetJsonPath())
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
	out.TargetJsonPath = direct.ValueOf(in.TargetJSONPath)
	out.OriginalValuePattern = direct.ValueOf(in.OriginalValuePattern)
	out.NewValue = direct.ValueOf(in.NewValue)
	return out
}
func GKEBackupRestorePlanSpec_FromProto(mapCtx *direct.MapContext, in *pb.RestorePlan) *krm.GKEBackupRestorePlanSpec {
	if in == nil {
		return nil
	}
	out := &krm.GKEBackupRestorePlanSpec{}
	// MISSING: Name
	// MISSING: Uid
	out.Description = direct.LazyPtr(in.GetDescription())
	if in.GetBackupPlan() != "" {
		out.BackupPlanRef = &krm.BackupPlanRef{External: in.GetBackupPlan()}
	}
	if in.GetCluster() != "" {
		out.ClusterRef = &container.ContainerClusterRef{External: in.GetCluster()}
	}
	out.RestoreConfig = RestoreConfig_FromProto(mapCtx, in.GetRestoreConfig())
	out.Labels = in.Labels
	return out
}
func RestoreConfig_AllNamespaces_ToProto(mapCtx *direct.MapContext, allNamespaces *bool) *pb.RestoreConfig_AllNamespaces {
	if allNamespaces == nil {
		return nil
	}
	out := &pb.RestoreConfig_AllNamespaces{
		AllNamespaces: direct.ValueOf(allNamespaces),
	}
	return out
}
func RestoreConfig_NoNamespaces_ToProto(mapCtx *direct.MapContext, noNamespaces *bool) *pb.RestoreConfig_NoNamespaces {
	if noNamespaces == nil {
		return nil
	}
	out := &pb.RestoreConfig_NoNamespaces{
		NoNamespaces: direct.ValueOf(noNamespaces),
	}
	return out
}
func RestoreConfig_VolumeDataRestorePolicyBinding_VolumeType_ToProto(mapCtx *direct.MapContext, volumeType *string) *pb.RestoreConfig_VolumeDataRestorePolicyBinding_VolumeType {
	if volumeType == nil {
		return nil
	}
	out := &pb.RestoreConfig_VolumeDataRestorePolicyBinding_VolumeType{}
	out.VolumeType = direct.Enum_ToProto[pb.VolumeTypeEnum_VolumeType](mapCtx, volumeType)
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
	if _, ok := in.NamespacedResourceRestoreScope.(*pb.RestoreConfig_AllNamespaces); ok {
		// special handling for oneof bool field to ensure it is round-trippable
		// LazyPtr will return nil if the value is the default value
		out.AllNamespaces = direct.PtrTo(in.GetAllNamespaces())
	}
	out.SelectedNamespaces = Namespaces_FromProto(mapCtx, in.GetSelectedNamespaces())
	out.SelectedApplications = NamespacedNames_FromProto(mapCtx, in.GetSelectedApplications())
	if _, ok := in.NamespacedResourceRestoreScope.(*pb.RestoreConfig_NoNamespaces); ok {
		// special handling for oneof bool field to ensure it is round-trippable
		// LazyPtr will return nil if the value is the default value
		out.NoNamespaces = direct.PtrTo(in.GetNoNamespaces())
	}
	out.ExcludedNamespaces = Namespaces_FromProto(mapCtx, in.GetExcludedNamespaces())
	out.SubstitutionRules = direct.Slice_FromProto(mapCtx, in.SubstitutionRules, RestoreConfig_SubstitutionRule_FromProto)
	out.TransformationRules = direct.Slice_FromProto(mapCtx, in.TransformationRules, RestoreConfig_TransformationRule_FromProto)
	out.VolumeDataRestorePolicyBindings = direct.Slice_FromProto(mapCtx, in.VolumeDataRestorePolicyBindings, RestoreConfig_VolumeDataRestorePolicyBinding_FromProto)
	out.RestoreOrder = RestoreConfig_RestoreOrder_FromProto(mapCtx, in.GetRestoreOrder())
	return out
}
func RestoreConfig_VolumeDataRestorePolicyBinding_FromProto(mapCtx *direct.MapContext, in *pb.RestoreConfig_VolumeDataRestorePolicyBinding) *krm.RestoreConfig_VolumeDataRestorePolicyBinding {
	if in == nil {
		return nil
	}
	out := &krm.RestoreConfig_VolumeDataRestorePolicyBinding{}
	out.Policy = direct.Enum_FromProto(mapCtx, in.GetPolicy())
	if in.Scope != nil {
		out.VolumeType = direct.Enum_FromProto(mapCtx, in.GetVolumeType())
	}
	return out
}
