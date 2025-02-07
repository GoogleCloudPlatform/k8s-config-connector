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

package backupdr

import (
	refs "github.com/GoogleCloudPlatform/k8s-config-connector/apis/refs/v1beta1"
	pb "cloud.google.com/go/backupdr/apiv1/backupdrpb"
	krm "github.com/GoogleCloudPlatform/k8s-config-connector/apis/backupdr/v1alpha1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct"
)
func BackupVault_FromProto(mapCtx *direct.MapContext, in *pb.BackupVault) *krm.BackupVault {
	if in == nil {
		return nil
	}
	out := &krm.BackupVault{}
	// MISSING: Name
	out.Description = in.Description
	out.Labels = in.Labels
	// MISSING: CreateTime
	// MISSING: UpdateTime
	out.BackupMinimumEnforcedRetentionDuration = direct.StringDuration_FromProto(mapCtx, in.GetBackupMinimumEnforcedRetentionDuration())
	// MISSING: Deletable
	out.Etag = in.Etag
	// MISSING: State
	out.EffectiveTime = direct.StringTimestamp_FromProto(mapCtx, in.GetEffectiveTime())
	// MISSING: BackupCount
	// MISSING: ServiceAccount
	// MISSING: TotalStoredBytes
	// MISSING: Uid
	out.Annotations = in.Annotations
	out.AccessRestriction = direct.Enum_FromProto(mapCtx, in.GetAccessRestriction())
	return out
}
func BackupVault_ToProto(mapCtx *direct.MapContext, in *krm.BackupVault) *pb.BackupVault {
	if in == nil {
		return nil
	}
	out := &pb.BackupVault{}
	// MISSING: Name
	out.Description = in.Description
	out.Labels = in.Labels
	// MISSING: CreateTime
	// MISSING: UpdateTime
	if oneof := direct.StringDuration_ToProto(mapCtx, in.BackupMinimumEnforcedRetentionDuration); oneof != nil {
		out.BackupMinimumEnforcedRetentionDuration = &pb.BackupVault_BackupMinimumEnforcedRetentionDuration{BackupMinimumEnforcedRetentionDuration: oneof}
	}
	// MISSING: Deletable
	out.Etag = in.Etag
	// MISSING: State
	if oneof := direct.StringTimestamp_ToProto(mapCtx, in.EffectiveTime); oneof != nil {
		out.EffectiveTime = &pb.BackupVault_EffectiveTime{EffectiveTime: oneof}
	}
	// MISSING: BackupCount
	// MISSING: ServiceAccount
	// MISSING: TotalStoredBytes
	// MISSING: Uid
	out.Annotations = in.Annotations
	out.AccessRestriction = direct.Enum_ToProto[pb.BackupVault_AccessRestriction](mapCtx, in.AccessRestriction)
	return out
}
func BackupVaultObservedState_FromProto(mapCtx *direct.MapContext, in *pb.BackupVault) *krm.BackupVaultObservedState {
	if in == nil {
		return nil
	}
	out := &krm.BackupVaultObservedState{}
	out.Name = direct.LazyPtr(in.GetName())
	// MISSING: Description
	// MISSING: Labels
	out.CreateTime = direct.StringTimestamp_FromProto(mapCtx, in.GetCreateTime())
	out.UpdateTime = direct.StringTimestamp_FromProto(mapCtx, in.GetUpdateTime())
	// MISSING: BackupMinimumEnforcedRetentionDuration
	out.Deletable = in.Deletable
	// MISSING: Etag
	out.State = direct.Enum_FromProto(mapCtx, in.GetState())
	// MISSING: EffectiveTime
	out.BackupCount = direct.LazyPtr(in.GetBackupCount())
	out.ServiceAccount = direct.LazyPtr(in.GetServiceAccount())
	out.TotalStoredBytes = direct.LazyPtr(in.GetTotalStoredBytes())
	out.Uid = direct.LazyPtr(in.GetUid())
	// MISSING: Annotations
	// MISSING: AccessRestriction
	return out
}
func BackupVaultObservedState_ToProto(mapCtx *direct.MapContext, in *krm.BackupVaultObservedState) *pb.BackupVault {
	if in == nil {
		return nil
	}
	out := &pb.BackupVault{}
	out.Name = direct.ValueOf(in.Name)
	// MISSING: Description
	// MISSING: Labels
	if oneof := direct.StringTimestamp_ToProto(mapCtx, in.CreateTime); oneof != nil {
		out.CreateTime = &pb.BackupVault_CreateTime{CreateTime: oneof}
	}
	if oneof := direct.StringTimestamp_ToProto(mapCtx, in.UpdateTime); oneof != nil {
		out.UpdateTime = &pb.BackupVault_UpdateTime{UpdateTime: oneof}
	}
	// MISSING: BackupMinimumEnforcedRetentionDuration
	out.Deletable = in.Deletable
	// MISSING: Etag
	out.State = direct.Enum_ToProto[pb.BackupVault_State](mapCtx, in.State)
	// MISSING: EffectiveTime
	out.BackupCount = direct.ValueOf(in.BackupCount)
	out.ServiceAccount = direct.ValueOf(in.ServiceAccount)
	out.TotalStoredBytes = direct.ValueOf(in.TotalStoredBytes)
	out.Uid = direct.ValueOf(in.Uid)
	// MISSING: Annotations
	// MISSING: AccessRestriction
	return out
}
func BackupdrBackupPlanAssociationObservedState_FromProto(mapCtx *direct.MapContext, in *pb.BackupPlanAssociation) *krm.BackupdrBackupPlanAssociationObservedState {
	if in == nil {
		return nil
	}
	out := &krm.BackupdrBackupPlanAssociationObservedState{}
	// MISSING: Name
	// MISSING: ResourceType
	// MISSING: Resource
	// MISSING: BackupPlan
	// MISSING: CreateTime
	// MISSING: UpdateTime
	// MISSING: State
	// MISSING: RulesConfigInfo
	// MISSING: DataSource
	return out
}
func BackupdrBackupPlanAssociationObservedState_ToProto(mapCtx *direct.MapContext, in *krm.BackupdrBackupPlanAssociationObservedState) *pb.BackupPlanAssociation {
	if in == nil {
		return nil
	}
	out := &pb.BackupPlanAssociation{}
	// MISSING: Name
	// MISSING: ResourceType
	// MISSING: Resource
	// MISSING: BackupPlan
	// MISSING: CreateTime
	// MISSING: UpdateTime
	// MISSING: State
	// MISSING: RulesConfigInfo
	// MISSING: DataSource
	return out
}
func BackupdrBackupPlanAssociationSpec_FromProto(mapCtx *direct.MapContext, in *pb.BackupPlanAssociation) *krm.BackupdrBackupPlanAssociationSpec {
	if in == nil {
		return nil
	}
	out := &krm.BackupdrBackupPlanAssociationSpec{}
	// MISSING: Name
	// MISSING: ResourceType
	// MISSING: Resource
	// MISSING: BackupPlan
	// MISSING: CreateTime
	// MISSING: UpdateTime
	// MISSING: State
	// MISSING: RulesConfigInfo
	// MISSING: DataSource
	return out
}
func BackupdrBackupPlanAssociationSpec_ToProto(mapCtx *direct.MapContext, in *krm.BackupdrBackupPlanAssociationSpec) *pb.BackupPlanAssociation {
	if in == nil {
		return nil
	}
	out := &pb.BackupPlanAssociation{}
	// MISSING: Name
	// MISSING: ResourceType
	// MISSING: Resource
	// MISSING: BackupPlan
	// MISSING: CreateTime
	// MISSING: UpdateTime
	// MISSING: State
	// MISSING: RulesConfigInfo
	// MISSING: DataSource
	return out
}
func BackupdrBackupPlanObservedState_FromProto(mapCtx *direct.MapContext, in *pb.BackupPlan) *krm.BackupdrBackupPlanObservedState {
	if in == nil {
		return nil
	}
	out := &krm.BackupdrBackupPlanObservedState{}
	// MISSING: Name
	// MISSING: Description
	// MISSING: Labels
	// MISSING: CreateTime
	// MISSING: UpdateTime
	// MISSING: BackupRules
	// MISSING: State
	// MISSING: ResourceType
	// MISSING: Etag
	// MISSING: BackupVault
	// MISSING: BackupVaultServiceAccount
	return out
}
func BackupdrBackupPlanObservedState_ToProto(mapCtx *direct.MapContext, in *krm.BackupdrBackupPlanObservedState) *pb.BackupPlan {
	if in == nil {
		return nil
	}
	out := &pb.BackupPlan{}
	// MISSING: Name
	// MISSING: Description
	// MISSING: Labels
	// MISSING: CreateTime
	// MISSING: UpdateTime
	// MISSING: BackupRules
	// MISSING: State
	// MISSING: ResourceType
	// MISSING: Etag
	// MISSING: BackupVault
	// MISSING: BackupVaultServiceAccount
	return out
}
func BackupdrBackupPlanSpec_FromProto(mapCtx *direct.MapContext, in *pb.BackupPlan) *krm.BackupdrBackupPlanSpec {
	if in == nil {
		return nil
	}
	out := &krm.BackupdrBackupPlanSpec{}
	// MISSING: Name
	// MISSING: Description
	// MISSING: Labels
	// MISSING: CreateTime
	// MISSING: UpdateTime
	// MISSING: BackupRules
	// MISSING: State
	// MISSING: ResourceType
	// MISSING: Etag
	// MISSING: BackupVault
	// MISSING: BackupVaultServiceAccount
	return out
}
func BackupdrBackupPlanSpec_ToProto(mapCtx *direct.MapContext, in *krm.BackupdrBackupPlanSpec) *pb.BackupPlan {
	if in == nil {
		return nil
	}
	out := &pb.BackupPlan{}
	// MISSING: Name
	// MISSING: Description
	// MISSING: Labels
	// MISSING: CreateTime
	// MISSING: UpdateTime
	// MISSING: BackupRules
	// MISSING: State
	// MISSING: ResourceType
	// MISSING: Etag
	// MISSING: BackupVault
	// MISSING: BackupVaultServiceAccount
	return out
}
func BackupdrBackupVaultObservedState_FromProto(mapCtx *direct.MapContext, in *pb.BackupVault) *krm.BackupdrBackupVaultObservedState {
	if in == nil {
		return nil
	}
	out := &krm.BackupdrBackupVaultObservedState{}
	// MISSING: Name
	// MISSING: Description
	// MISSING: Labels
	// MISSING: CreateTime
	// MISSING: UpdateTime
	// MISSING: BackupMinimumEnforcedRetentionDuration
	// MISSING: Deletable
	// MISSING: Etag
	// MISSING: State
	// MISSING: EffectiveTime
	// MISSING: BackupCount
	// MISSING: ServiceAccount
	// MISSING: TotalStoredBytes
	// MISSING: Uid
	// MISSING: Annotations
	// MISSING: AccessRestriction
	return out
}
func BackupdrBackupVaultObservedState_ToProto(mapCtx *direct.MapContext, in *krm.BackupdrBackupVaultObservedState) *pb.BackupVault {
	if in == nil {
		return nil
	}
	out := &pb.BackupVault{}
	// MISSING: Name
	// MISSING: Description
	// MISSING: Labels
	// MISSING: CreateTime
	// MISSING: UpdateTime
	// MISSING: BackupMinimumEnforcedRetentionDuration
	// MISSING: Deletable
	// MISSING: Etag
	// MISSING: State
	// MISSING: EffectiveTime
	// MISSING: BackupCount
	// MISSING: ServiceAccount
	// MISSING: TotalStoredBytes
	// MISSING: Uid
	// MISSING: Annotations
	// MISSING: AccessRestriction
	return out
}
func BackupdrBackupVaultSpec_FromProto(mapCtx *direct.MapContext, in *pb.BackupVault) *krm.BackupdrBackupVaultSpec {
	if in == nil {
		return nil
	}
	out := &krm.BackupdrBackupVaultSpec{}
	// MISSING: Name
	// MISSING: Description
	// MISSING: Labels
	// MISSING: CreateTime
	// MISSING: UpdateTime
	// MISSING: BackupMinimumEnforcedRetentionDuration
	// MISSING: Deletable
	// MISSING: Etag
	// MISSING: State
	// MISSING: EffectiveTime
	// MISSING: BackupCount
	// MISSING: ServiceAccount
	// MISSING: TotalStoredBytes
	// MISSING: Uid
	// MISSING: Annotations
	// MISSING: AccessRestriction
	return out
}
func BackupdrBackupVaultSpec_ToProto(mapCtx *direct.MapContext, in *krm.BackupdrBackupVaultSpec) *pb.BackupVault {
	if in == nil {
		return nil
	}
	out := &pb.BackupVault{}
	// MISSING: Name
	// MISSING: Description
	// MISSING: Labels
	// MISSING: CreateTime
	// MISSING: UpdateTime
	// MISSING: BackupMinimumEnforcedRetentionDuration
	// MISSING: Deletable
	// MISSING: Etag
	// MISSING: State
	// MISSING: EffectiveTime
	// MISSING: BackupCount
	// MISSING: ServiceAccount
	// MISSING: TotalStoredBytes
	// MISSING: Uid
	// MISSING: Annotations
	// MISSING: AccessRestriction
	return out
}
func BackupdrManagementServerObservedState_FromProto(mapCtx *direct.MapContext, in *pb.ManagementServer) *krm.BackupdrManagementServerObservedState {
	if in == nil {
		return nil
	}
	out := &krm.BackupdrManagementServerObservedState{}
	// MISSING: Name
	// MISSING: Description
	// MISSING: Labels
	// MISSING: CreateTime
	// MISSING: UpdateTime
	// MISSING: Type
	// MISSING: ManagementURI
	// MISSING: WorkforceIdentityBasedManagementURI
	// MISSING: State
	// MISSING: Networks
	// MISSING: Etag
	// MISSING: Oauth2ClientID
	// MISSING: WorkforceIdentityBasedOauth2ClientID
	// MISSING: BaProxyURI
	// MISSING: SatisfiesPzs
	// MISSING: SatisfiesPzi
	return out
}
func BackupdrManagementServerObservedState_ToProto(mapCtx *direct.MapContext, in *krm.BackupdrManagementServerObservedState) *pb.ManagementServer {
	if in == nil {
		return nil
	}
	out := &pb.ManagementServer{}
	// MISSING: Name
	// MISSING: Description
	// MISSING: Labels
	// MISSING: CreateTime
	// MISSING: UpdateTime
	// MISSING: Type
	// MISSING: ManagementURI
	// MISSING: WorkforceIdentityBasedManagementURI
	// MISSING: State
	// MISSING: Networks
	// MISSING: Etag
	// MISSING: Oauth2ClientID
	// MISSING: WorkforceIdentityBasedOauth2ClientID
	// MISSING: BaProxyURI
	// MISSING: SatisfiesPzs
	// MISSING: SatisfiesPzi
	return out
}
func BackupdrManagementServerSpec_FromProto(mapCtx *direct.MapContext, in *pb.ManagementServer) *krm.BackupdrManagementServerSpec {
	if in == nil {
		return nil
	}
	out := &krm.BackupdrManagementServerSpec{}
	// MISSING: Name
	// MISSING: Description
	// MISSING: Labels
	// MISSING: CreateTime
	// MISSING: UpdateTime
	// MISSING: Type
	// MISSING: ManagementURI
	// MISSING: WorkforceIdentityBasedManagementURI
	// MISSING: State
	// MISSING: Networks
	// MISSING: Etag
	// MISSING: Oauth2ClientID
	// MISSING: WorkforceIdentityBasedOauth2ClientID
	// MISSING: BaProxyURI
	// MISSING: SatisfiesPzs
	// MISSING: SatisfiesPzi
	return out
}
func BackupdrManagementServerSpec_ToProto(mapCtx *direct.MapContext, in *krm.BackupdrManagementServerSpec) *pb.ManagementServer {
	if in == nil {
		return nil
	}
	out := &pb.ManagementServer{}
	// MISSING: Name
	// MISSING: Description
	// MISSING: Labels
	// MISSING: CreateTime
	// MISSING: UpdateTime
	// MISSING: Type
	// MISSING: ManagementURI
	// MISSING: WorkforceIdentityBasedManagementURI
	// MISSING: State
	// MISSING: Networks
	// MISSING: Etag
	// MISSING: Oauth2ClientID
	// MISSING: WorkforceIdentityBasedOauth2ClientID
	// MISSING: BaProxyURI
	// MISSING: SatisfiesPzs
	// MISSING: SatisfiesPzi
	return out
}
