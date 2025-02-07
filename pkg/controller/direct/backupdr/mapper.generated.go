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
func BackupPlanAssociation_FromProto(mapCtx *direct.MapContext, in *pb.BackupPlanAssociation) *krm.BackupPlanAssociation {
	if in == nil {
		return nil
	}
	out := &krm.BackupPlanAssociation{}
	// MISSING: Name
	out.ResourceType = direct.LazyPtr(in.GetResourceType())
	out.Resource = direct.LazyPtr(in.GetResource())
	out.BackupPlan = direct.LazyPtr(in.GetBackupPlan())
	// MISSING: CreateTime
	// MISSING: UpdateTime
	// MISSING: State
	// MISSING: RulesConfigInfo
	// MISSING: DataSource
	return out
}
func BackupPlanAssociation_ToProto(mapCtx *direct.MapContext, in *krm.BackupPlanAssociation) *pb.BackupPlanAssociation {
	if in == nil {
		return nil
	}
	out := &pb.BackupPlanAssociation{}
	// MISSING: Name
	out.ResourceType = direct.ValueOf(in.ResourceType)
	out.Resource = direct.ValueOf(in.Resource)
	out.BackupPlan = direct.ValueOf(in.BackupPlan)
	// MISSING: CreateTime
	// MISSING: UpdateTime
	// MISSING: State
	// MISSING: RulesConfigInfo
	// MISSING: DataSource
	return out
}
func BackupPlanAssociationObservedState_FromProto(mapCtx *direct.MapContext, in *pb.BackupPlanAssociation) *krm.BackupPlanAssociationObservedState {
	if in == nil {
		return nil
	}
	out := &krm.BackupPlanAssociationObservedState{}
	out.Name = direct.LazyPtr(in.GetName())
	// MISSING: ResourceType
	// MISSING: Resource
	// MISSING: BackupPlan
	out.CreateTime = direct.StringTimestamp_FromProto(mapCtx, in.GetCreateTime())
	out.UpdateTime = direct.StringTimestamp_FromProto(mapCtx, in.GetUpdateTime())
	out.State = direct.Enum_FromProto(mapCtx, in.GetState())
	out.RulesConfigInfo = direct.Slice_FromProto(mapCtx, in.RulesConfigInfo, RuleConfigInfo_FromProto)
	out.DataSource = direct.LazyPtr(in.GetDataSource())
	return out
}
func BackupPlanAssociationObservedState_ToProto(mapCtx *direct.MapContext, in *krm.BackupPlanAssociationObservedState) *pb.BackupPlanAssociation {
	if in == nil {
		return nil
	}
	out := &pb.BackupPlanAssociation{}
	out.Name = direct.ValueOf(in.Name)
	// MISSING: ResourceType
	// MISSING: Resource
	// MISSING: BackupPlan
	out.CreateTime = direct.StringTimestamp_ToProto(mapCtx, in.CreateTime)
	out.UpdateTime = direct.StringTimestamp_ToProto(mapCtx, in.UpdateTime)
	out.State = direct.Enum_ToProto[pb.BackupPlanAssociation_State](mapCtx, in.State)
	out.RulesConfigInfo = direct.Slice_ToProto(mapCtx, in.RulesConfigInfo, RuleConfigInfo_ToProto)
	out.DataSource = direct.ValueOf(in.DataSource)
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
func RuleConfigInfo_FromProto(mapCtx *direct.MapContext, in *pb.RuleConfigInfo) *krm.RuleConfigInfo {
	if in == nil {
		return nil
	}
	out := &krm.RuleConfigInfo{}
	// MISSING: RuleID
	// MISSING: LastBackupState
	// MISSING: LastBackupError
	// MISSING: LastSuccessfulBackupConsistencyTime
	return out
}
func RuleConfigInfo_ToProto(mapCtx *direct.MapContext, in *krm.RuleConfigInfo) *pb.RuleConfigInfo {
	if in == nil {
		return nil
	}
	out := &pb.RuleConfigInfo{}
	// MISSING: RuleID
	// MISSING: LastBackupState
	// MISSING: LastBackupError
	// MISSING: LastSuccessfulBackupConsistencyTime
	return out
}
func RuleConfigInfoObservedState_FromProto(mapCtx *direct.MapContext, in *pb.RuleConfigInfo) *krm.RuleConfigInfoObservedState {
	if in == nil {
		return nil
	}
	out := &krm.RuleConfigInfoObservedState{}
	out.RuleID = direct.LazyPtr(in.GetRuleId())
	out.LastBackupState = direct.Enum_FromProto(mapCtx, in.GetLastBackupState())
	out.LastBackupError = Status_FromProto(mapCtx, in.GetLastBackupError())
	out.LastSuccessfulBackupConsistencyTime = direct.StringTimestamp_FromProto(mapCtx, in.GetLastSuccessfulBackupConsistencyTime())
	return out
}
func RuleConfigInfoObservedState_ToProto(mapCtx *direct.MapContext, in *krm.RuleConfigInfoObservedState) *pb.RuleConfigInfo {
	if in == nil {
		return nil
	}
	out := &pb.RuleConfigInfo{}
	out.RuleId = direct.ValueOf(in.RuleID)
	out.LastBackupState = direct.Enum_ToProto[pb.RuleConfigInfo_LastBackupState](mapCtx, in.LastBackupState)
	out.LastBackupError = Status_ToProto(mapCtx, in.LastBackupError)
	out.LastSuccessfulBackupConsistencyTime = direct.StringTimestamp_ToProto(mapCtx, in.LastSuccessfulBackupConsistencyTime)
	return out
}
